// Exemplo 7: Props.children
// 💡 O que é 'children'?
// É uma prop mágica! ✨
// Tudo que você coloca DENTRO de um componente (<Comp>AQUI</Comp>)
// aparece automaticamente na prop 'children'.

function Caixa({ children, titulo, cor = '#e3f2fd' }) {
  return (
    <div style={{
      padding: '1rem',
      background: cor,
      borderRadius: '8px',
      border: '2px solid #2196f3',
      marginBottom: '1rem'
    }}>
      {/* Se tiver título, mostra o título */}
      {titulo && (
        <h4 style={{ margin: '0 0 0.5rem 0', color: '#1976d2' }}>{titulo}</h4>
      )}
      
      {/* Aqui é onde o conteúdo "filho" será renderizado */}
      <div style={{ color: '#555' }}>
        {children}
      </div>
    </div>
  )
}

function Exemplo7Children() {
  return (
    <div>
      {/* 1. Passando texto simples como children */}
      <Caixa titulo="Texto Simples" cor="#e8f5e9">
        <p style={{ margin: 0 }}>Eu sou um parágrafo passado como children!</p>
      </Caixa>

      {/* 2. Passando HTML complexo como children */}
      <Caixa titulo="HTML Complexo" cor="#fff3e0">
        <p style={{ margin: '0 0 0.5rem 0' }}>Olha o que dá pra fazer:</p>
        <ul style={{ margin: 0, paddingLeft: '1.5rem' }}>
          <li>Listas</li>
          <li><strong>Negrito</strong></li>
          <li><em>Itálico</em></li>
        </ul>
      </Caixa>

      {/* 3. Passando Outros Componentes como children (Composição!) */}
      <Caixa titulo="Interatividade" cor="#fce4ec">
        <p style={{ margin: '0 0 0.5rem 0' }}>Botões funcionam aqui dentro:</p>
        <button 
          onClick={() => alert('Clicou!')}
          style={{
            padding: '0.5rem 1rem',
            background: '#e91e63',
            color: 'white',
            border: 'none',
            borderRadius: '4px',
            cursor: 'pointer'
          }}
        >
          Clique em mim
        </button>
      </Caixa>
    </div>
  )
}

export default Exemplo7Children
