import Menu from '../Menu/Menu'
import Footer from '../Footer/Footer'
import './Layout.css'

function Layout({ children }) {
  return (
    <div className="layout">
      <Menu />
      <main className="layout-main">
        {children}
      </main>
      <Footer />
    </div>
  )
}

export default Layout

