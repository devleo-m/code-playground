import './Home.css'

function Home() {
  return (
    <div className="home-page">
      <div className="home-hero">
        <h1>📚 Fundamentos de React</h1>
        <p>Aprenda React de forma prática e organizada</p>
        <p className="home-description">
          Este guia contém aulas práticas sobre React. Cada aula explica a teoria
          e mostra exemplos práticos que você pode ver em ação.
        </p>
      </div>
      
      <div className="home-content">
        <h2>Como usar este guia</h2>
        <ul className="home-instructions">
          <li>📖 Leia a teoria em cada aula</li>
          <li>💻 Veja os exemplos práticos em código</li>
          <li>👀 Visualize os componentes funcionando</li>
          <li>🔍 Entenda como tudo se conecta</li>
        </ul>
      </div>
    </div>
  )
}

export default Home

