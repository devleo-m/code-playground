// Exemplo 7: Props.children
function Caixa({ children, titulo, cor = '#e3f2fd' }) {
  return (
    <div style={{
      padding: '1rem',
      background: cor,
      borderRadius: '8px',
      border: '2px solid #2196f3',
      marginBottom: '1rem'
    }}>
      {titulo && (
        <h4 style={{ margin: '0 0 0.5rem 0', color: '#1976d2' }}>{titulo}</h4>
      )}
      <div style={{ color: '#555' }}>
        {children}
      </div>
    </div>
  )
}

function Exemplo7Children() {
  return (
    <div>
      <Caixa titulo="Caixa com Título e Conteúdo Simples" cor="#e8f5e9">
        <p style={{ margin: 0 }}>Este é um parágrafo dentro da caixa usando children.</p>
      </Caixa>

      <Caixa titulo="Caixa com Múltiplos Elementos" cor="#fff3e0">
        <p style={{ margin: '0 0 0.5rem 0' }}>Você pode colocar qualquer coisa dentro:</p>
        <ul style={{ margin: 0, paddingLeft: '1.5rem' }}>
          <li>Lista de itens</li>
          <li>Mais elementos</li>
          <li>Qualquer JSX válido</li>
        </ul>
      </Caixa>

      <Caixa titulo="Caixa com Botões e Texto" cor="#fce4ec">
        <p style={{ margin: '0 0 0.5rem 0' }}>Até mesmo botões podem ser children:</p>
        <button 
          style={{
            padding: '0.5rem 1rem',
            background: '#e91e63',
            color: 'white',
            border: 'none',
            borderRadius: '4px',
            cursor: 'pointer'
          }}
        >
          Botão dentro da caixa!
        </button>
      </Caixa>

      <div style={{ 
        marginTop: '1rem', 
        padding: '1rem', 
        background: '#fff9c4', 
        borderRadius: '8px',
        border: '1px solid #fbc02d'
      }}>
        <p style={{ margin: 0, fontSize: '0.9rem', color: '#f57f17' }}>
          <strong>💡 Dica:</strong> <code>children</code> é uma prop especial que contém 
          tudo que você coloca entre as tags de abertura e fechamento do componente. 
          É muito útil para criar componentes genéricos e reutilizáveis!
        </p>
      </div>
    </div>
  )
}

export default Exemplo7Children



