import { useState } from 'react'

export function Exemplo1UseState() {
  const [count, setCount] = useState(0)

  function incrementNormal() {
    setCount(count + 1)
    setCount(count + 1)
    setCount(count + 1)
    // Resultado: Aumenta apenas 1, pois 'count' é o mesmo valor neste ciclo.
  }

  function incrementFunctional() {
    setCount(c => c + 1)
    setCount(c => c + 1)
    setCount(c => c + 1)
    // Resultado: Aumenta 3, pois usa o valor mais recente do estado anterior.
  }

  return (
    <div style={{ padding: '20px', border: '1px solid #ccc', borderRadius: '8px' }}>
      <h3>Contador: {count}</h3>
      <div style={{ display: 'flex', gap: '10px' }}>
        <button onClick={incrementNormal}>+3 (Normal - Falha)</button>
        <button onClick={incrementFunctional}>+3 (Functional - Correto)</button>
        <button onClick={() => setCount(0)}>Reset</button>
      </div>
      <p style={{ fontSize: '0.8rem', color: '#666' }}>
        Abra o console ou veja o código para entender a diferença.
      </p>
    </div>
  )
}

