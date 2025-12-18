// Exemplo 8: JSX - Expressões e Regras
// 💡 JSX = JavaScript + XML
// Parece HTML, mas é JavaScript poderoso!
// A regra de ouro é: use chaves {} para voltar para o "mundo JavaScript".

function Exemplo8JSX() {
  // Variáveis JavaScript normais
  const nome = 'Aluno React'
  const nota = 10
  const tecnologias = ['React', 'Vite', 'JavaScript']
  const estaLogado = true
  
  const estiloCard = {
    marginBottom: '1rem', 
    padding: '0.75rem', 
    background: 'white', 
    borderRadius: '4px',
    boxShadow: '0 2px 4px rgba(0,0,0,0.05)'
  }

  return (
    <div style={{ padding: '1rem', background: '#f5f5f5', borderRadius: '8px' }}>
      <h3 style={{ color: '#2c3e50', margin: '0 0 1rem 0' }}>O Poder do JSX</h3>
      
      {/* 1. Variáveis */}
      <div style={estiloCard}>
        <h4 style={{ margin: '0 0 0.5rem 0' }}>1. Variáveis</h4>
        <p style={{ margin: 0 }}>
          Olá, <strong>{nome}</strong>! Sua nota é <strong>{nota}</strong>.
        </p>
      </div>

      {/* 2. Expressões Matemáticas */}
      <div style={estiloCard}>
        <h4 style={{ margin: '0 0 0.5rem 0' }}>2. Matemática</h4>
        <p style={{ margin: 0 }}>
          Quanto é 5 + 5? Resposta: <strong>{5 + 5}</strong>
        </p>
      </div>

      {/* 3. Lógica Condicional (Ternário) */}
      <div style={estiloCard}>
        <h4 style={{ margin: '0 0 0.5rem 0' }}>3. Lógica</h4>
        <p style={{ margin: 0 }}>
          Status: {estaLogado ? <span style={{color: 'green'}}>🟢 Online</span> : '🔴 Offline'}
        </p>
      </div>

      {/* 4. Loops (Map) */}
      <div style={estiloCard}>
        <h4 style={{ margin: '0 0 0.5rem 0' }}>4. Listas (Map)</h4>
        <p style={{ margin: '0 0 0.5rem 0' }}>Tecnologias que estamos aprendendo:</p>
        <ul style={{ margin: 0, paddingLeft: '1.5rem' }}>
          {/* Transformamos cada string do array em um elemento <li> */}
          {tecnologias.map((tech, index) => (
            <li key={index}>{tech}</li>
          ))}
        </ul>
      </div>

      <div style={{ 
        marginTop: '1rem', 
        padding: '0.75rem', 
        background: '#e3f2fd', 
        borderRadius: '4px', 
        border: '1px solid #2196f3',
        fontSize: '0.9rem',
        color: '#1565c0'
      }}>
        <strong>💡 Dica Pro:</strong> Tudo dentro de <code>{'{ }'}</code> é executado como JavaScript.
        Você pode chamar funções, fazer contas, acessar objetos... o céu é o limite!
      </div>
    </div>
  )
}

export default Exemplo8JSX
