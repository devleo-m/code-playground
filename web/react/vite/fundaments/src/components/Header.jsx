import Contador from './Contador'
import './Header.css'

function Header() {
  return (
    <>
      <header className="app-header">
        <h1>📚 Guia de Estudos React</h1>
        <p>Aprenda React de forma prática e organizada</p>
      </header>
      <>
        <Contador />
      </>
    </>
  )
}

export default Header

