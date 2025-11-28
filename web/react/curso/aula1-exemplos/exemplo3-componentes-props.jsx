import { useState } from 'react'

/**
 * Exemplo 3: Componentes com Props
 * 
 * Este exemplo demonstra:
 * - Como criar componentes reutilizáveis
 * - Como passar props (propriedades)
 * - Como usar props para customizar componentes
 * - Composição de componentes
 */

/**
 * Componente Card reutilizável
 * Recebe props: title, description, color
 */
function Card({ title, description, color, onAction }) {
  return (
    <div style={{ 
      border: `3px solid ${color}`, 
      padding: '20px', 
      margin: '15px',
      borderRadius: '8px',
      backgroundColor: '#fff',
      boxShadow: '0 2px 4px rgba(0,0,0,0.1)',
      transition: 'transform 0.2s',
      cursor: onAction ? 'pointer' : 'default'
    }}
    onClick={onAction}
    onMouseEnter={(e) => {
      if (onAction) {
        e.currentTarget.style.transform = 'scale(1.05)'
      }
    }}
    onMouseLeave={(e) => {
      e.currentTarget.style.transform = 'scale(1)'
    }}
    >
      <h3 style={{ 
        margin: '0 0 10px 0',
        color: color 
      }}>
        {title}
      </h3>
      <p style={{ 
        margin: 0,
        color: '#666',
        lineHeight: '1.5'
      }}>
        {description}
      </p>
    </div>
  )
}

/**
 * Componente Badge para mostrar contador
 */
function Badge({ count, label }) {
  return (
    <div style={{
      display: 'inline-block',
      backgroundColor: '#4CAF50',
      color: 'white',
      padding: '5px 10px',
      borderRadius: '20px',
      fontSize: '14px',
      margin: '5px'
    }}>
      {label}: {count}
    </div>
  )
}

/**
 * Componente principal que usa os componentes acima
 */
function App() {
  // Estado para contar cliques em cada card
  const [clicks, setClicks] = useState({
    react: 0,
    vite: 0,
    javascript: 0
  })

  // Função para incrementar contador
  function handleCardClick(cardName) {
    setClicks(prev => ({
      ...prev,
      [cardName]: prev[cardName] + 1
    }))
  }

  return (
    <div style={{ 
      padding: '20px',
      fontFamily: 'Arial, sans-serif',
      maxWidth: '800px',
      margin: '0 auto'
    }}>
      <h1 style={{ textAlign: 'center' }}>
        Tecnologias Web Modernas
      </h1>

      <div style={{ 
        textAlign: 'center',
        marginBottom: '20px'
      }}>
        <Badge count={clicks.react} label="React" />
        <Badge count={clicks.vite} label="Vite" />
        <Badge count={clicks.javascript} label="JavaScript" />
      </div>

      {/* Usando o mesmo componente Card com props diferentes */}
      <Card 
        title="⚛️ React" 
        description="Biblioteca JavaScript para construir interfaces de usuário. Baseada em componentes reutilizáveis e Virtual DOM para performance." 
        color="#61DAFB"
        onAction={() => handleCardClick('react')}
      />
      
      <Card 
        title="⚡ Vite" 
        description="Build tool moderna que oferece desenvolvimento extremamente rápido. Usa ESM nativo e compilação sob demanda." 
        color="#646CFF"
        onAction={() => handleCardClick('vite')}
      />
      
      <Card 
        title="🟨 JavaScript" 
        description="Linguagem de programação que roda no navegador. É a base de React e todas as tecnologias web modernas." 
        color="#F7DF1E"
        onAction={() => handleCardClick('javascript')}
      />

      <div style={{
        marginTop: '30px',
        padding: '15px',
        backgroundColor: '#f0f0f0',
        borderRadius: '8px',
        fontSize: '14px',
        color: '#666'
      }}>
        <strong>💡 Dica:</strong> Clique nos cards acima para ver os contadores aumentarem!
        Isso demonstra como componentes podem ter comportamento interativo através de props.
      </div>
    </div>
  )
}

export default App

