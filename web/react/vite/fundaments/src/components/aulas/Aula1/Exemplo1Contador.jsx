import { useState } from 'react'

// Exemplo 1: Contador Simples
// 💡 O básico do State
// O useState é como a memória do componente.
// Quando usamos setCount, o React sabe que precisa atualizar a tela!

function Exemplo1Contador() {
  // useState retorna: [valorAtual, funcaoParaAtualizar]
  const [count, setCount] = useState(0)

  return (
    <div style={{ padding: '1rem', textAlign: 'center' }}>
      <h2>Contador: {count}</h2>
      
      <div style={{ marginTop: '1rem', display: 'flex', gap: '0.5rem', justifyContent: 'center' }}>
        <button 
          onClick={() => setCount(count + 1)}
          style={buttonStyle('#4caf50')}
        >
          Incrementar (+1)
        </button>
        
        <button 
          onClick={() => setCount(count - 1)}
          style={buttonStyle('#f44336')}
        >
          Decrementar (-1)
        </button>
        
        <button 
          onClick={() => setCount(0)}
          style={buttonStyle('#9e9e9e')}
        >
          Resetar
        </button>
      </div>

      <p style={{ marginTop: '1rem', color: '#666' }}>
        Clique nos botões e veja o número mudar magicamente! ✨
      </p>
    </div>
  )
}

const buttonStyle = (color) => ({
  margin: '5px', 
  padding: '10px 20px',
  fontSize: '1rem',
  cursor: 'pointer',
  backgroundColor: color,
  color: 'white',
  border: 'none',
  borderRadius: '4px'
})

export default Exemplo1Contador

