import { useState } from 'react'

/**
 * Exemplo 1: Contador Simples
 * 
 * Este exemplo demonstra:
 * - Como usar useState para gerenciar estado
 * - Como atualizar estado com setState
 * - Como eventos onClick funcionam
 * - Re-renderização automática quando estado muda
 */
function Counter() {
  // useState retorna [valor, função para atualizar]
  // Inicializamos count com 0
  const [count, setCount] = useState(0)

  return (
    <div style={{ 
      padding: '20px', 
      textAlign: 'center',
      fontFamily: 'Arial, sans-serif'
    }}>
      <h2>Contador: {count}</h2>
      
      <div style={{ marginTop: '20px' }}>
        <button 
          onClick={() => setCount(count + 1)}
          style={{ 
            margin: '5px', 
            padding: '10px 20px',
            fontSize: '16px',
            cursor: 'pointer'
          }}
        >
          Incrementar (+1)
        </button>
        
        <button 
          onClick={() => setCount(count - 1)}
          style={{ 
            margin: '5px', 
            padding: '10px 20px',
            fontSize: '16px',
            cursor: 'pointer'
          }}
        >
          Decrementar (-1)
        </button>
        
        <button 
          onClick={() => setCount(0)}
          style={{ 
            margin: '5px', 
            padding: '10px 20px',
            fontSize: '16px',
            cursor: 'pointer',
            backgroundColor: '#ff6b6b',
            color: 'white',
            border: 'none',
            borderRadius: '4px'
          }}
        >
          Resetar
        </button>
      </div>

      <p style={{ marginTop: '20px', color: '#666' }}>
        Clique nos botões para ver o estado mudar!
      </p>
    </div>
  )
}

export default Counter

