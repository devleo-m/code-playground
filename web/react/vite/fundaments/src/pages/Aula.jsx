import { useParams } from 'react-router-dom'
import CodeExample from '../components/CodeExample'
import './Aula.css'

function Aula() {
  const { aulaId } = useParams()

  // Estrutura base para as aulas
  // Cada aula terá sua teoria e exemplos práticos
  // Exemplo de como usar:
  /*
  const teoria = `
    <h3>O que é React?</h3>
    <p>React é uma biblioteca JavaScript para construir interfaces de usuário...</p>
  `
  
  const exemploCodigo = `
    function MeuComponente() {
      return <h1>Olá, React!</h1>
    }
  `
  
  function ExemploComponente() {
    return <h1>Olá, React!</h1>
  }
  */

  return (
    <div className="aula-page">
      <div className="aula-header">
        <h1>Aula {aulaId}</h1>
        <p className="aula-subtitle">Conteúdo da aula será adicionado aqui</p>
      </div>
      
      <div className="aula-content">
        <div className="aula-section">
          <h2>📖 Teoria</h2>
          <div className="aula-theory">
            <p>O conteúdo teórico será exibido aqui.</p>
            <p>Você pode usar HTML dentro da seção de teoria para formatar o texto.</p>
          </div>
        </div>

        <div className="aula-section">
          <h2>💻 Prática</h2>
          <div className="aula-practice">
            <p>Os exemplos práticos serão exibidos aqui usando o componente CodeExample.</p>
            <p>Exemplo de uso do CodeExample:</p>
            {/* 
              <CodeExample
                title="Título do Exemplo"
                description="Descrição do que o exemplo faz"
                code={exemploCodigo}
                ExampleComponent={ExemploComponente}
              />
            */}
          </div>
        </div>
      </div>
    </div>
  )
}

export default Aula

