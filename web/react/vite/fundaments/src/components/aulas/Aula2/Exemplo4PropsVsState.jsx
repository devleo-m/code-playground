import { useState } from 'react'

// Exemplo 4: Props vs State
function DisplayProps({ mensagem }) {
  return (
    <div style={{ 
      padding: '1rem', 
      background: '#e3f2fd', 
      borderRadius: '8px', 
      border: '2px solid #2196f3',
      marginBottom: '1rem'
    }}>
      <h4 style={{ color: '#1976d2', margin: '0 0 0.5rem 0' }}>Props (Read-Only)</h4>
      <p style={{ color: '#555', margin: 0 }}>Mensagem recebida: <strong>{mensagem}</strong></p>
      <p style={{ color: '#666', margin: '0.5rem 0 0 0', fontSize: '0.85rem' }}>
        Esta mensagem vem do componente pai e não pode ser modificada aqui.
      </p>
    </div>
  )
}

function DisplayState() {
  const [mensagem, setMensagem] = useState('Estado inicial')

  return (
    <div style={{ 
      padding: '1rem', 
      background: '#f3e5f5', 
      borderRadius: '8px', 
      border: '2px solid #9c27b0'
    }}>
      <h4 style={{ color: '#7b1fa2', margin: '0 0 0.5rem 0' }}>State (Mutável)</h4>
      <p style={{ color: '#555', margin: '0 0 1rem 0' }}>
        Mensagem atual: <strong>{mensagem}</strong>
      </p>
      <div style={{ display: 'flex', gap: '0.5rem', flexWrap: 'wrap' }}>
        <button 
          onClick={() => setMensagem('Olá do State!')}
          style={{
            padding: '0.4rem 0.8rem',
            background: '#9c27b0',
            color: 'white',
            border: 'none',
            borderRadius: '4px',
            cursor: 'pointer',
            fontSize: '0.9rem'
          }}
        >
          Mudar para "Olá do State!"
        </button>
        <button 
          onClick={() => setMensagem('Estado foi atualizado!')}
          style={{
            padding: '0.4rem 0.8rem',
            background: '#9c27b0',
            color: 'white',
            border: 'none',
            borderRadius: '4px',
            cursor: 'pointer',
            fontSize: '0.9rem'
          }}
        >
          Mudar para "Estado foi atualizado!"
        </button>
        <button 
          onClick={() => setMensagem('Estado inicial')}
          style={{
            padding: '0.4rem 0.8rem',
            background: '#757575',
            color: 'white',
            border: 'none',
            borderRadius: '4px',
            cursor: 'pointer',
            fontSize: '0.9rem'
          }}
        >
          Resetar
        </button>
      </div>
      <p style={{ color: '#666', margin: '1rem 0 0 0', fontSize: '0.85rem' }}>
        Este estado pode ser modificado pelo próprio componente.
      </p>
    </div>
  )
}

function Exemplo4PropsVsState() {
  return (
    <div>
      <DisplayProps mensagem="Esta mensagem vem via props e não pode ser modificada" />
      <DisplayState />
    </div>
  )
}

export default Exemplo4PropsVsState



