import { Link, useLocation } from 'react-router-dom'
import './Menu.css'

function Menu() {
  const location = useLocation()

  return (
    <nav className="menu-nav">
      <div className="menu-container">
        <div className="menu-logo">
          <Link to="/">
            <span className="logo-icon">⚛️</span>
            <span className="logo-text">React Fundamentos</span>
          </Link>
        </div>
        
        <ul className="menu-links">
          <li>
            <Link 
              to="/" 
              className={location.pathname === '/' ? 'active' : ''}
            >
              Home
            </Link>
          </li>
          <li>
            <Link 
              to="/aulas" 
              className={location.pathname.startsWith('/aulas') ? 'active' : ''}
            >
              Aulas
            </Link>
          </li>
        </ul>
      </div>
    </nav>
  )
}

export default Menu

