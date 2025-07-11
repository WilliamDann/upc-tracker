import { createRoot } from 'react-dom/client'
import { BrowserRouter, Routes, Route } from 'react-router'

import App from './App.jsx'
import NotFound from './pages/NotFound.jsx'

import 'bootstrap/dist/css/bootstrap.min.css';

import Search from './pages/Search.jsx';
import Scan from './pages/Scan.jsx'
import Home from './Home.jsx';

import Product from './resources/Product.jsx'
import Account from './resources/Account.jsx'
import Place from './resources/Place.jsx'

import Signin from './pages/Signin.jsx'

createRoot(document.getElementById('root')).render(
  <BrowserRouter>
    <Routes>
      <Route path="*" element={<NotFound />} />

      <Route path="/" element={<App />}>
        <Route index element={<Home />} />
        <Route path="search" element={<Search />} />
        <Route path="scan" element={<Scan />} />

        {new Account().Routes({})}
        {new Product().Routes({})}
        {new Place().Routes({})}
      </Route>

      <Route path="signin" element={<Signin />} />

    </Routes>
  </BrowserRouter>,
)
