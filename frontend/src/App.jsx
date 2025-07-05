import { Link, Outlet } from 'react-router';
import SideBar from './components/SideBar';

function App() {
  return (
    <div className="d-flex vh-100">
      <div className="bg-dark text-white" style={{ width: '215px' }}>
        <SideBar />
      </div>

      {/* Flexible main content */}
      <div className="flex-grow-1 p-4">
        <Outlet />
      </div>
    </div>
  )
}

export default App
