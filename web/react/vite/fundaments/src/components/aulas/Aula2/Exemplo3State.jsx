import { useState } from 'react'

// Exemplo 3: Componente com State
function Exemplo3State() {
  const [count, setCount] = useState(0)

  return (
    <div style={{ padding: '1rem', background: '#fff3e0', borderRadius: '8px', border: '2px solid #ff9800' }}>
      <h3 style={{ color: '#2c3e50', margin: '0 0 1rem 0' }}>Contador: {count}</h3>
      <div style={{ display: 'flex', gap: '0.5rem' }}>
        <button 
          onClick={() => setCount(count - 1)}
          style={{
            padding: '0.5rem 1rem',
            background: '#f44336',
            color: 'white',
            border: 'none',
            borderRadius: '4px',
            cursor: 'pointer',
            fontSize: '1rem'
          }}
        >
          - Decrementar
        </button>
        <button 
          onClick={() => setCount(0)}
          style={{
            padding: '0.5rem 1rem',
            background: '#9e9e9e',
            color: 'white',
            border: 'none',
            borderRadius: '4px',
            cursor: 'pointer',
            fontSize: '1rem'
          }}
        >
          Resetar
        </button>
        <button 
          onClick={() => setCount(count + 1)}
          style={{
            padding: '0.5rem 1rem',
            background: '#4caf50',
            color: 'white',
            border: 'none',
            borderRadius: '4px',
            cursor: 'pointer',
            fontSize: '1rem'
          }}
        >
          + Incrementar
        </button>
      </div>
      <p style={{ color: '#666', margin: '1rem 0 0 0', fontSize: '0.9rem' }}>
        {count === 0 && 'O contador está zerado'}
        {count > 0 && `Você incrementou ${count} vez${count > 1 ? 'es' : ''}`}
        {count < 0 && `Você decrementou ${Math.abs(count)} vez${Math.abs(count) > 1 ? 'es' : ''}`}
      </p>
    </div>
  )
}

export default Exemplo3State



