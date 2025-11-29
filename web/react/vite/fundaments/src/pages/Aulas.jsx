import { Link } from 'react-router-dom'
import { AULAS } from '../constants/aulas'
import './Aulas.css'

function Aulas() {
  const aulas = AULAS

  return (
    <div className="aulas-page">
      <div className="aulas-header">
        <h1>📚 Lista de Aulas</h1>
        <p>Selecione uma aula para começar a estudar</p>
      </div>

      {aulas.length === 0 ? (
        <div className="aulas-empty">
          <p>As aulas serão adicionadas em breve!</p>
          <p className="aulas-empty-sub">Aguarde enquanto preparamos o conteúdo.</p>
        </div>
      ) : (
        <div className="aulas-list">
          {aulas.map((aula) => (
            <Link 
              key={aula.id} 
              to={`/aula/${aula.id}`}
              className="aula-card"
            >
              <h3>{aula.titulo}</h3>
              <p>{aula.descricao}</p>
            </Link>
          ))}
        </div>
      )}
    </div>
  )
}

export default Aulas

