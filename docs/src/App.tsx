import { BrowserRouter, Routes, Route, Navigate } from 'react-router-dom';
import { Layout } from './components/Layout';
import { Introduction } from './pages/Introduction';
import { ApiReference } from './pages/ApiReference';
import { Architecture } from './pages/Architecture';

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Layout />}>
          <Route index element={<Introduction />} />
          <Route path="architecture" element={<Architecture />} />
          <Route path="api" element={<ApiReference />} />
          {/* Redirects for specific sections to the main API page with hash (handled by browser) or just to the page */}
          <Route path="api/*" element={<ApiReference />} />
          <Route path="*" element={<Navigate to="/" replace />} />
        </Route>
      </Routes>
    </BrowserRouter>
  );
}

export default App;
