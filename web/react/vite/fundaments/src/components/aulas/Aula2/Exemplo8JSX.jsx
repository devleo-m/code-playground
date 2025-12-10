// Exemplo 8: JSX - Expressões e Regras
function Exemplo8JSX() {
  const nome = 'React'
  const versao = 18
  const tecnologias = ['JSX', 'Hooks', 'Components']
  const estaAtivo = true
  const usuario = { nome: 'Maria', idade: 25 }

  return (
    <div style={{ padding: '1rem', background: '#f5f5f5', borderRadius: '8px' }}>
      <h3 style={{ color: '#2c3e50', margin: '0 0 1rem 0' }}>JSX em Ação</h3>
      
      <div style={{ marginBottom: '1rem', padding: '0.75rem', background: 'white', borderRadius: '4px' }}>
        <h4 style={{ margin: '0 0 0.5rem 0', fontSize: '1rem' }}>1. Variáveis em JSX</h4>
        <p style={{ margin: 0, color: '#555' }}>
          Olá, <strong>{nome}</strong>! Versão atual: <strong>{versao}</strong>
        </p>
      </div>

      <div style={{ marginBottom: '1rem', padding: '0.75rem', background: 'white', borderRadius: '4px' }}>
        <h4 style={{ margin: '0 0 0.5rem 0', fontSize: '1rem' }}>2. Expressões Matemáticas</h4>
        <p style={{ margin: 0, color: '#555' }}>
          10 + 5 = <strong>{10 + 5}</strong>
        </p>
        <p style={{ margin: '0.25rem 0 0 0', color: '#555' }}>
          {usuario.nome} tem {usuario.idade} anos e nasceu em <strong>{2024 - usuario.idade}</strong>
        </p>
      </div>

      <div style={{ marginBottom: '1rem', padding: '0.75rem', background: 'white', borderRadius: '4px' }}>
        <h4 style={{ margin: '0 0 0.5rem 0', fontSize: '1rem' }}>3. Operador Ternário</h4>
        <p style={{ margin: 0, color: '#555' }}>
          Status: {estaAtivo ? (
            <span style={{ color: '#4caf50', fontWeight: 'bold' }}>✅ Ativo</span>
          ) : (
            <span style={{ color: '#f44336', fontWeight: 'bold' }}>❌ Inativo</span>
          )}
        </p>
      </div>

      <div style={{ marginBottom: '1rem', padding: '0.75rem', background: 'white', borderRadius: '4px' }}>
        <h4 style={{ margin: '0 0 0.5rem 0', fontSize: '1rem' }}>4. Renderizando Arrays</h4>
        <p style={{ margin: '0 0 0.5rem 0', color: '#555' }}>Tecnologias do React:</p>
        <ul style={{ margin: 0, paddingLeft: '1.5rem', color: '#555' }}>
          {tecnologias.map((tech, index) => (
            <li key={index}>{tech}</li>
          ))}
        </ul>
      </div>

      <div style={{ marginBottom: '1rem', padding: '0.75rem', background: 'white', borderRadius: '4px' }}>
        <h4 style={{ margin: '0 0 0.5rem 0', fontSize: '1rem' }}>5. Atributos em camelCase</h4>
        <div style={{ color: '#555' }}>
          <p style={{ margin: 0 }}>Em JSX usamos <code style={{ background: '#f0f0f0', padding: '0.2rem 0.4rem', borderRadius: '3px' }}>className</code> ao invés de <code style={{ background: '#f0f0f0', padding: '0.2rem 0.4rem', borderRadius: '3px' }}>class</code></p>
          <p style={{ margin: '0.25rem 0 0 0' }}>E <code style={{ background: '#f0f0f0', padding: '0.2rem 0.4rem', borderRadius: '3px' }}>onClick</code> ao invés de <code style={{ background: '#f0f0f0', padding: '0.2rem 0.4rem', borderRadius: '3px' }}>onclick</code></p>
        </div>
      </div>

      <div style={{ 
        marginTop: '1rem', 
        padding: '0.75rem', 
        background: '#e3f2fd', 
        borderRadius: '4px',
        border: '1px solid #2196f3'
      }}>
        <p style={{ margin: 0, fontSize: '0.9rem', color: '#1976d2' }}>
          <strong>💡 Lembre-se:</strong> Dentro de <code>{'{}'}</code> você pode usar qualquer expressão JavaScript válida!
        </p>
      </div>
    </div>
  )
}

export default Exemplo8JSX



