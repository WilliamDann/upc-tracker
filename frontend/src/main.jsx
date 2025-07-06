import { createRoot } from 'react-dom/client'
import { BrowserRouter, Routes, Route } from 'react-router'

import App from './App.jsx'
import NotFound from './NotFound.jsx'

import 'bootstrap/dist/css/bootstrap.min.css';
import Search from './Search.jsx';
import Scan from './Scan.jsx'
import Home from './Home.jsx';

import List from './components/Products/List.jsx'
import View from './components/Products/View.jsx'
import Edit from './components/Products/Edit.jsx';

createRoot(document.getElementById('root')).render(
  <BrowserRouter>
    <Routes>
      <Route path="*" element={<NotFound />} />

      <Route path="/" element={<App />}>
        <Route index element={<Home />} />
        <Route path="search" element={<Search />} />
        <Route path="scan" element={<Scan />} />

        <Route path='/products'>
          <Route index element={<List />} />
          <Route path='/products/view/:id' element={<View />} />
          <Route path='/products/edit/:id' element={<Edit />} />
        </Route>
      </Route>

    </Routes>
  </BrowserRouter>,
)
