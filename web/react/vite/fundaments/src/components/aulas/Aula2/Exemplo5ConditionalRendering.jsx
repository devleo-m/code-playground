import { useState } from 'react'

// Exemplo 5: Conditional Rendering
function Exemplo5ConditionalRendering() {
  const [estaLogado, setEstaLogado] = useState(false)
  const [mostrarNotificacao, setMostrarNotificacao] = useState(true)
  const [status, setStatus] = useState('loading')

  return (
    <div style={{ display: 'flex', flexDirection: 'column', gap: '1.5rem' }}>
      {/* Exemplo 1: Operador Ternário */}
      <div style={{ padding: '1rem', background: '#fff9c4', borderRadius: '8px', border: '2px solid #fbc02d' }}>
        <h4 style={{ color: '#f57f17', margin: '0 0 0.5rem 0' }}>1. Operador Ternário</h4>
        <div style={{ marginBottom: '1rem' }}>
          {estaLogado ? (
            <div style={{ padding: '0.5rem', background: '#c8e6c9', borderRadius: '4px' }}>
              ✅ <strong>Bem-vindo de volta!</strong> Você está logado.
            </div>
          ) : (
            <div style={{ padding: '0.5rem', background: '#ffcdd2', borderRadius: '4px' }}>
              ⚠️ <strong>Por favor, faça login</strong> para continuar.
            </div>
          )}
        </div>
        <button 
          onClick={() => setEstaLogado(!estaLogado)}
          style={{
            padding: '0.5rem 1rem',
            background: estaLogado ? '#f44336' : '#4caf50',
            color: 'white',
            border: 'none',
            borderRadius: '4px',
            cursor: 'pointer'
          }}
        >
          {estaLogado ? 'Fazer Logout' : 'Fazer Login'}
        </button>
      </div>

      {/* Exemplo 2: Operador && */}
      <div style={{ padding: '1rem', background: '#e1f5fe', borderRadius: '8px', border: '2px solid #0288d1' }}>
        <h4 style={{ color: '#01579b', margin: '0 0 0.5rem 0' }}>2. Operador && (AND)</h4>
        <div style={{ marginBottom: '1rem' }}>
          {mostrarNotificacao && (
            <div style={{ 
              padding: '0.75rem', 
              background: '#0288d1', 
              color: 'white', 
              borderRadius: '4px',
              display: 'flex',
              justifyContent: 'space-between',
              alignItems: 'center'
            }}>
              <span>🔔 Esta é uma notificação importante!</span>
              <button
                onClick={() => setMostrarNotificacao(false)}
                style={{
                  background: 'rgba(255,255,255,0.3)',
                  border: 'none',
                  color: 'white',
                  padding: '0.25rem 0.5rem',
                  borderRadius: '4px',
                  cursor: 'pointer'
                }}
              >
                ✕
              </button>
            </div>
          )}
        </div>
        <button 
          onClick={() => setMostrarNotificacao(!mostrarNotificacao)}
          style={{
            padding: '0.5rem 1rem',
            background: mostrarNotificacao ? '#0288d1' : '#9e9e9e',
            color: 'white',
            border: 'none',
            borderRadius: '4px',
            cursor: 'pointer'
          }}
        >
          {mostrarNotificacao ? 'Ocultar Notificação' : 'Mostrar Notificação'}
        </button>
      </div>

      {/* Exemplo 3: Múltiplas Condições */}
      <div style={{ padding: '1rem', background: '#fce4ec', borderRadius: '8px', border: '2px solid #e91e63' }}>
        <h4 style={{ color: '#c2185b', margin: '0 0 0.5rem 0' }}>3. Múltiplas Condições</h4>
        <div style={{ marginBottom: '1rem', minHeight: '60px' }}>
          {status === 'loading' && (
            <div style={{ padding: '0.75rem', background: '#fff9c4', borderRadius: '4px' }}>
              ⏳ <strong>Carregando...</strong> Aguarde um momento.
            </div>
          )}
          {status === 'success' && (
            <div style={{ padding: '0.75rem', background: '#c8e6c9', borderRadius: '4px' }}>
              ✅ <strong>Sucesso!</strong> Dados carregados com sucesso.
            </div>
          )}
          {status === 'error' && (
            <div style={{ padding: '0.75rem', background: '#ffcdd2', borderRadius: '4px' }}>
              ❌ <strong>Erro!</strong> Não foi possível carregar os dados.
            </div>
          )}
        </div>
        <div style={{ display: 'flex', gap: '0.5rem', flexWrap: 'wrap' }}>
          <button 
            onClick={() => setStatus('loading')}
            style={{
              padding: '0.4rem 0.8rem',
              background: status === 'loading' ? '#fbc02d' : '#fff9c4',
              color: status === 'loading' ? 'white' : '#333',
              border: 'none',
              borderRadius: '4px',
              cursor: 'pointer'
            }}
          >
            Loading
          </button>
          <button 
            onClick={() => setStatus('success')}
            style={{
              padding: '0.4rem 0.8rem',
              background: status === 'success' ? '#4caf50' : '#c8e6c9',
              color: status === 'success' ? 'white' : '#333',
              border: 'none',
              borderRadius: '4px',
              cursor: 'pointer'
            }}
          >
            Success
          </button>
          <button 
            onClick={() => setStatus('error')}
            style={{
              padding: '0.4rem 0.8rem',
              background: status === 'error' ? '#f44336' : '#ffcdd2',
              color: status === 'error' ? 'white' : '#333',
              border: 'none',
              borderRadius: '4px',
              cursor: 'pointer'
            }}
          >
            Error
          </button>
        </div>
      </div>
    </div>
  )
}

export default Exemplo5ConditionalRendering



