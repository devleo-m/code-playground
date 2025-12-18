import { useState } from 'react'

// Exemplo 5: Conditional Rendering (Renderização Condicional)
// 💡 Como mostrar/esconder coisas no React?
// Usamos JavaScript puro! Não tem "v-if" ou "ng-if" aqui.
// As formas mais comuns são: Operador Ternário e Operador && (AND).

function Exemplo5ConditionalRendering() {
  const [estaLogado, setEstaLogado] = useState(false)
  const [mostrarNotificacao, setMostrarNotificacao] = useState(true)
  const [status, setStatus] = useState('loading')

  return (
    <div style={{ display: 'flex', flexDirection: 'column', gap: '1.5rem' }}>
      
      {/* 1. Operador Ternário (condicao ? verdadeiro : falso) */}
      <div style={boxStyle('#fff9c4', '#fbc02d')}>
        <h4 style={{ color: '#f57f17', margin: '0 0 0.5rem 0' }}>1. Operador Ternário</h4>
        <div style={{ marginBottom: '1rem' }}>
          {estaLogado ? (
            // Bloco SE VERDADEIRO
            <div style={alertStyle('#c8e6c9')}>
              ✅ <strong>Bem-vindo de volta!</strong> Você está logado.
            </div>
          ) : (
            // Bloco SE FALSO
            <div style={alertStyle('#ffcdd2')}>
              ⚠️ <strong>Por favor, faça login</strong> para continuar.
            </div>
          )}
        </div>
        
        <button 
          onClick={() => setEstaLogado(!estaLogado)}
          style={buttonStyle(estaLogado ? '#f44336' : '#4caf50')}
        >
          {estaLogado ? 'Fazer Logout' : 'Fazer Login'}
        </button>
      </div>

      {/* 2. Operador && (condicao && renderiza) */}
      {/* Útil quando você quer mostrar algo se for true, ou nada se for false */}
      <div style={boxStyle('#e1f5fe', '#0288d1')}>
        <h4 style={{ color: '#01579b', margin: '0 0 0.5rem 0' }}>2. Operador && (AND)</h4>
        <div style={{ marginBottom: '1rem' }}>
          {mostrarNotificacao && (
            <div style={{ ...alertStyle('#0288d1'), color: 'white', display: 'flex', justifyContent: 'space-between' }}>
              <span>🔔 Notificação importante!</span>
              <button
                onClick={() => setMostrarNotificacao(false)}
                style={{ background: 'transparent', border: 'none', color: 'white', cursor: 'pointer' }}
              >
                ✕
              </button>
            </div>
          )}
        </div>
        
        <button 
          onClick={() => setMostrarNotificacao(!mostrarNotificacao)}
          style={buttonStyle(mostrarNotificacao ? '#0288d1' : '#9e9e9e')}
        >
          {mostrarNotificacao ? 'Ocultar' : 'Mostrar'}
        </button>
      </div>

      {/* 3. Múltiplas Condições (comum para status de loading/erro) */}
      <div style={boxStyle('#fce4ec', '#e91e63')}>
        <h4 style={{ color: '#c2185b', margin: '0 0 0.5rem 0' }}>3. Estados de Carregamento</h4>
        <div style={{ marginBottom: '1rem', minHeight: '60px' }}>
          {status === 'loading' && (
            <div style={alertStyle('#fff9c4')}>⏳ <strong>Carregando...</strong></div>
          )}
          {status === 'success' && (
            <div style={alertStyle('#c8e6c9')}>✅ <strong>Sucesso!</strong> Dados carregados.</div>
          )}
          {status === 'error' && (
            <div style={alertStyle('#ffcdd2')}>❌ <strong>Erro!</strong> Tente novamente.</div>
          )}
        </div>
        
        <div style={{ display: 'flex', gap: '0.5rem' }}>
          <button onClick={() => setStatus('loading')} style={buttonStyle('#fbc02d')}>Loading</button>
          <button onClick={() => setStatus('success')} style={buttonStyle('#4caf50')}>Success</button>
          <button onClick={() => setStatus('error')} style={buttonStyle('#f44336')}>Error</button>
        </div>
      </div>
    </div>
  )
}

// Helpers de estilo para manter o código limpo
const boxStyle = (bg, border) => ({
  padding: '1rem',
  background: bg,
  borderRadius: '8px',
  border: `2px solid ${border}`
})

const alertStyle = (bg) => ({
  padding: '0.75rem',
  background: bg,
  borderRadius: '4px'
})

const buttonStyle = (bg) => ({
  padding: '0.5rem 1rem',
  background: bg,
  color: 'white',
  border: 'none',
  borderRadius: '4px',
  cursor: 'pointer'
})

export default Exemplo5ConditionalRendering
