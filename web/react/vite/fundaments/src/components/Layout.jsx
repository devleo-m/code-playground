import Menu from './interface/Menu'
import Footer from './interface/Footer'
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

