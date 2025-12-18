import { useState } from 'react'

// Exemplo 3: Componentes e Props
// 💡 Reutilização é a chave!
// Em vez de copiar e colar código, criamos componentes.
// Props são como "argumentos" para nossos componentes.

// 1. Componente Card Reutilizável
function Card({ titulo, descricao, cor, onClick }) {
  return (
    <div 
      onClick={onClick}
      style={{ 
        border: `2px solid ${cor}`, 
        padding: '1rem', 
        borderRadius: '8px',
        cursor: 'pointer',
        transition: 'transform 0.2s',
        background: 'white'
      }}
      onMouseEnter={(e) => e.currentTarget.style.transform = 'scale(1.02)'}
      onMouseLeave={(e) => e.currentTarget.style.transform = 'scale(1)'}
    >
      <h4 style={{ color: cor, margin: '0 0 0.5rem 0' }}>{titulo}</h4>
      <p style={{ margin: 0, color: '#555', fontSize: '0.9rem' }}>{descricao}</p>
    </div>
  )
}

// 2. Componente Badge (Etiqueta)
function Badge({ label, count }) {
  return (
    <span style={{
      background: '#eee',
      padding: '5px 10px',
      borderRadius: '20px',
      fontSize: '0.8rem',
      margin: '0 5px'
    }}>
      {label}: <strong>{count}</strong>
    </span>
  )
}

function Exemplo3ComponentesProps() {
  const [clicks, setClicks] = useState({ react: 0, vite: 0, js: 0 })

  const addClick = (key) => {
    setClicks(prev => ({ ...prev, [key]: prev[key] + 1 }))
  }

  return (
    <div>
      <div style={{ textAlign: 'center', marginBottom: '1rem' }}>
        <Badge label="React" count={clicks.react} />
        <Badge label="Vite" count={clicks.vite} />
        <Badge label="JS" count={clicks.js} />
      </div>

      <div style={{ display: 'grid', gridTemplateColumns: '1fr 1fr 1fr', gap: '1rem' }}>
        <Card 
          titulo="React" 
          descricao="Biblioteca para interfaces."
          cor="#61DAFB"
          onClick={() => addClick('react')}
        />
        <Card 
          titulo="Vite" 
          descricao="Build tool super rápida."
          cor="#646CFF"
          onClick={() => addClick('vite')}
        />
        <Card 
          titulo="JavaScript" 
          descricao="A linguagem da web."
          cor="#F7DF1E"
          onClick={() => addClick('js')}
        />
      </div>
      
      <p style={{ textAlign: 'center', marginTop: '1rem', color: '#666', fontSize: '0.9rem' }}>
        👆 Clique nos cards para votar na sua tecnologia favorita!
      </p>
    </div>
  )
}

export default Exemplo3ComponentesProps

