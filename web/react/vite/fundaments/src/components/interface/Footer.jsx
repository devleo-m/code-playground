import './Footer.css'

function Footer() {
  const currentYear = new Date().getFullYear()

  return (
    <footer className="app-footer">
      <div className="footer-container">
        <p className="footer-text">
          © {currentYear} React Fundamentos - Guia de Estudos Prático
        </p>
        <p className="footer-subtext">
          Aprenda React de forma clara e prática
        </p>
      </div>
    </footer>
  )
}

export default Footer

