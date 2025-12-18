import { useState } from 'react'

// Exemplo 3: Componente com State
// 💡 O que é State?
// State (estado) é a memória do componente. Diferente de variáveis comuns,
// quando o state muda, o React atualiza a tela automaticamente (re-render).

function Exemplo3State() {
  // useState(0) cria um estado com valor inicial 0.
  // Retorna um array com dois itens:
  // 1. count: o valor atual
  // 2. setCount: a função para mudar esse valor
  const [count, setCount] = useState(0)

  return (
    <div style={{ padding: '1rem', background: '#fff3e0', borderRadius: '8px', border: '2px solid #ff9800' }}>
      <h3 style={{ color: '#2c3e50', margin: '0 0 1rem 0' }}>Contador: {count}</h3>
      
      <div style={{ display: 'flex', gap: '0.5rem' }}>
        <button 
          onClick={() => setCount(count - 1)}
          style={buttonStyle('#f44336')} // Vermelho
        >
          - Decrementar
        </button>
        
        <button 
          onClick={() => setCount(0)}
          style={buttonStyle('#9e9e9e')} // Cinza
        >
          Resetar
        </button>
        
        <button 
          onClick={() => setCount(count + 1)}
          style={buttonStyle('#4caf50')} // Verde
        >
          + Incrementar
        </button>
      </div>

      <p style={{ color: '#666', margin: '1rem 0 0 0', fontSize: '0.9rem' }}>
        {/* Renderização Condicional: Mostra mensagens diferentes baseado no valor */}
        {count === 0 && 'O contador está zerado'}
        {count > 0 && `Você incrementou ${count} vez${count > 1 ? 'es' : ''}`}
        {count < 0 && `Você decrementou ${Math.abs(count)} vez${Math.abs(count) > 1 ? 'es' : ''}`}
      </p>
    </div>
  )
}

// Helper para estilos repetitivos (mantendo o código limpo)
const buttonStyle = (color) => ({
  padding: '0.5rem 1rem',
  background: color,
  color: 'white',
  border: 'none',
  borderRadius: '4px',
  cursor: 'pointer',
  fontSize: '1rem'
})

export default Exemplo3State
