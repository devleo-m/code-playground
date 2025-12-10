import { useParams } from 'react-router-dom'
import { CodeExample } from '../components/ui'
import { AULAS } from '../constants/aulas'
import './Aula.css'

function Aula() {
  const { aulaId } = useParams()
  const aula = AULAS.find(a => a.id === aulaId)

  if (!aula) {
    return (
      <div className="aula-page">
        <div className="aula-header">
          <h1>Aula não encontrada</h1>
          <p className="aula-subtitle">A aula com ID "{aulaId}" não existe.</p>
        </div>
      </div>
    )
  }

  return (
    <div className="aula-page">
      <div className="aula-header">
        <h1>{aula.titulo}</h1>
        {aula.descricao && <p className="aula-subtitle">{aula.descricao}</p>}
      </div>
      
      <div className="aula-content">
        <div className="aula-section">
          <h2>📖 Teoria</h2>
          <div 
            className="aula-theory"
            dangerouslySetInnerHTML={{ __html: aula.teoria }}
          />
        </div>

        <div className="aula-section">
          <h2>💻 Prática</h2>
          <div className="aula-practice">
            {aula.exemplos?.map((exemplo, index) => (
              <CodeExample
                key={index}
                title={exemplo.title}
                description={exemplo.description}
                code={exemplo.code}
                ExampleComponent={exemplo.ExampleComponent}
              />
            ))}
          </div>
        </div>
      </div>
    </div>
  )
}

export default Aula

