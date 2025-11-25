package handlers

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"sync"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"

	"github.com/lucaspanzera1/chat/internal/auth"
	"github.com/lucaspanzera1/chat/internal/models"
	"github.com/lucaspanzera1/chat/internal/repository"
)

type OAuthHandler struct {
	userRepo   *repository.UserRepository
	config     *oauth2.Config
	stateStore map[string]time.Time
	stateMutex sync.RWMutex
}

type GoogleUserInfo struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Name          string `json:"name"`
	Picture       string `json:"picture"`
}

func NewOAuthHandler(userRepo *repository.UserRepository) *OAuthHandler {
	config := &oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		RedirectURL:  os.Getenv("GOOGLE_REDIRECT_URL"),
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}

	handler := &OAuthHandler{
		userRepo:   userRepo,
		config:     config,
		stateStore: make(map[string]time.Time),
	}

	go handler.cleanupStates()

	return handler
}

func (h *OAuthHandler) generateState() string {
	b := make([]byte, 32)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)

	h.stateMutex.Lock()
	h.stateStore[state] = time.Now().Add(10 * time.Minute)
	h.stateMutex.Unlock()

	return state
}

func (h *OAuthHandler) validateState(state string) bool {
	h.stateMutex.RLock()
	expiry, exists := h.stateStore[state]
	h.stateMutex.RUnlock()

	if !exists || time.Now().After(expiry) {
		return false
	}

	h.stateMutex.Lock()
	delete(h.stateStore, state)
	h.stateMutex.Unlock()

	return true
}

func (h *OAuthHandler) cleanupStates() {
	ticker := time.NewTicker(5 * time.Minute)
	for range ticker.C {
		h.stateMutex.Lock()
		now := time.Now()
		for state, expiry := range h.stateStore {
			if now.After(expiry) {
				delete(h.stateStore, state)
			}
		}
		h.stateMutex.Unlock()
	}
}

func (h *OAuthHandler) GoogleLogin(w http.ResponseWriter, r *http.Request) {
	if h.config.ClientID == "" {
		http.Error(w, "Google OAuth não configurado", http.StatusInternalServerError)
		return
	}

	state := h.generateState()
	url := h.config.AuthCodeURL(state, oauth2.AccessTypeOffline)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func (h *OAuthHandler) GoogleCallback(w http.ResponseWriter, r *http.Request) {
	state := r.URL.Query().Get("state")
	if !h.validateState(state) {
		http.Redirect(w, r, "/auth.html?error=invalid_state", http.StatusTemporaryRedirect)
		return
	}

	code := r.URL.Query().Get("code")
	if code == "" {
		http.Redirect(w, r, "/auth.html?error=no_code", http.StatusTemporaryRedirect)
		return
	}

	token, err := h.config.Exchange(context.Background(), code)
	if err != nil {
		log.Printf("Erro ao trocar código por token: %v", err)
		http.Redirect(w, r, "/auth.html?error=token_exchange", http.StatusTemporaryRedirect)
		return
	}

	client := h.config.Client(context.Background(), token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		log.Printf("Erro ao buscar info do usuário: %v", err)
		http.Redirect(w, r, "/auth.html?error=user_info", http.StatusTemporaryRedirect)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Erro ao ler resposta: %v", err)
		http.Redirect(w, r, "/auth.html?error=read_response", http.StatusTemporaryRedirect)
		return
	}

	var googleUser GoogleUserInfo
	if err := json.Unmarshal(body, &googleUser); err != nil {
		log.Printf("Erro ao parsear usuário: %v", err)
		http.Redirect(w, r, "/auth.html?error=parse_user", http.StatusTemporaryRedirect)
		return
	}

	user, isNewUser, err := h.findOrCreateUser(r.Context(), googleUser)
	if err != nil {
		log.Printf("Erro ao criar/buscar usuário: %v", err)
		http.Redirect(w, r, "/auth.html?error=create_user", http.StatusTemporaryRedirect)
		return
	}

	jwtToken, err := auth.GenerateToken(user)
	if err != nil {
		log.Printf("Erro ao gerar token: %v", err)
		http.Redirect(w, r, "/auth.html?error=generate_token", http.StatusTemporaryRedirect)
		return
	}

	if isNewUser || user.Username == "" {
		redirectURL := fmt.Sprintf("/setup.html?token=%s&email=%s&avatar=%s",
			jwtToken,
			url.QueryEscape(user.Email),
			url.QueryEscape(googleUser.Picture))
		http.Redirect(w, r, redirectURL, http.StatusTemporaryRedirect)
		return
	}

	redirectURL := fmt.Sprintf("/auth.html?token=%s&user=%s", jwtToken, user.Username)
	http.Redirect(w, r, redirectURL, http.StatusTemporaryRedirect)
}

func (h *OAuthHandler) findOrCreateUser(ctx context.Context, googleUser GoogleUserInfo) (*models.User, bool, error) {

	user, err := h.userRepo.GetByGoogleID(ctx, googleUser.ID)
	if err != nil {
		return nil, false, err
	}
	if user != nil {
		return user, false, nil
	}

	user, err = h.userRepo.GetByEmail(ctx, googleUser.Email)
	if err != nil {
		return nil, false, err
	}
	if user != nil {
		if err := h.userRepo.LinkGoogleAccount(ctx, user.ID, googleUser.ID, googleUser.Picture); err != nil {
			return nil, false, err
		}
		user.GoogleID = &googleUser.ID
		if user.AvatarURL == nil {
			user.AvatarURL = &googleUser.Picture
		}
		return user, false, nil
	}

	user, err = h.userRepo.CreateWithGoogle(ctx, googleUser.Email, googleUser.ID, googleUser.Picture)
	if err != nil {
		return nil, false, err
	}
	return user, true, nil
}
