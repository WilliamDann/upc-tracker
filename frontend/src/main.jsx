import { createRoot } from 'react-dom/client'
import { BrowserRouter, Routes, Route } from 'react-router'

import App from './App.jsx'
import NotFound from './NotFound.jsx'

import 'bootstrap/dist/css/bootstrap.min.css';
import Search from './Search.jsx';
import Scan from './Scan.jsx'
import Home from './Home.jsx';

createRoot(document.getElementById('root')).render(
  <BrowserRouter>
    <Routes>
      <Route path="*" element={<NotFound />} />

      <Route path="/" element={<App />}>
        <Route index element={<Home />} />
        <Route path="search" element={<Search />} />
        <Route path="scan" element={<Scan />} />
      </Route>

    </Routes>
  </BrowserRouter>,
)
