import { useState, useEffect } from 'react'

export function Exemplo2UseEffect() {
  const [count, setCount] = useState(0)
  const [text, setText] = useState('')

  // 1. Executa a cada render (Sem array)
  useEffect(() => {
    console.log('🔄 Efeito Global (Todo Render)')
  })

  // 2. Executa apenas no Mount (Array Vazio)
  useEffect(() => {
    console.log('✅ Efeito Mount (Apenas uma vez)')
    
    // Cleanup no Unmount
    return () => console.log('❌ Componente Desmontado')
  }, [])

  // 3. Executa quando 'count' muda
  useEffect(() => {
    console.log(`🔢 Count mudou para: ${count}`)
  }, [count])

  return (
    <div style={{ padding: '20px', border: '1px solid #ccc', borderRadius: '8px' }}>
      <h3>Abra o Console (F12)</h3>
      <div style={{ marginBottom: '10px' }}>
        <button onClick={() => setCount(c => c + 1)}>Count: {count}</button>
      </div>
      <div>
        <input 
          value={text} 
          onChange={e => setText(e.target.value)} 
          placeholder="Digite algo..."
        />
        <p>Texto: {text}</p>
      </div>
    </div>
  )
}

