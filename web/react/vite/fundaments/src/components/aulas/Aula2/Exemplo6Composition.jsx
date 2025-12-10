// Exemplo 6: Composition (Composição)
function Botao({ children, onClick, variant = 'primary' }) {
  const styles = {
    primary: { background: '#2196f3', color: 'white' },
    secondary: { background: '#757575', color: 'white' },
    danger: { background: '#f44336', color: 'white' }
  }

  return (
    <button
      onClick={onClick}
      style={{
        padding: '0.5rem 1rem',
        border: 'none',
        borderRadius: '4px',
        cursor: 'pointer',
        fontSize: '0.9rem',
        ...styles[variant]
      }}
    >
      {children}
    </button>
  )
}

function Card({ title, children, footer }) {
  return (
    <div style={{
      border: '2px solid #e0e0e0',
      borderRadius: '8px',
      overflow: 'hidden',
      marginBottom: '1rem'
    }}>
      {title && (
        <div style={{
          padding: '0.75rem',
          background: '#f5f5f5',
          borderBottom: '1px solid #e0e0e0',
          fontWeight: 'bold',
          color: '#2c3e50'
        }}>
          {title}
        </div>
      )}
      <div style={{ padding: '1rem' }}>
        {children}
      </div>
      {footer && (
        <div style={{
          padding: '0.75rem',
          background: '#fafafa',
          borderTop: '1px solid #e0e0e0'
        }}>
          {footer}
        </div>
      )}
    </div>
  )
}

function Container({ children, title }) {
  return (
    <div style={{
      padding: '1.5rem',
      background: '#f8f9fa',
      borderRadius: '8px',
      border: '1px solid #dee2e6'
    }}>
      {title && <h3 style={{ margin: '0 0 1rem 0', color: '#2c3e50' }}>{title}</h3>}
      {children}
    </div>
  )
}

function Exemplo6Composition() {
  return (
    <div>
      <Container title="Exemplo de Composição">
        <Card 
          title="Produto: Notebook"
          footer={
            <div style={{ display: 'flex', gap: '0.5rem', justifyContent: 'flex-end' }}>
              <Botao variant="secondary">Cancelar</Botao>
              <Botao variant="primary">Adicionar ao Carrinho</Botao>
            </div>
          }
        >
          <p style={{ margin: '0 0 0.5rem 0' }}>Preço: R$ 2.500,00</p>
          <p style={{ margin: 0, color: '#666', fontSize: '0.9rem' }}>
            Notebook de alta performance com 16GB RAM e SSD 512GB
          </p>
        </Card>

        <Card 
          title="Produto: Mouse"
          footer={
            <div style={{ display: 'flex', gap: '0.5rem', justifyContent: 'flex-end' }}>
              <Botao variant="secondary">Cancelar</Botao>
              <Botao variant="danger">Remover</Botao>
            </div>
          }
        >
          <p style={{ margin: '0 0 0.5rem 0' }}>Preço: R$ 50,00</p>
          <p style={{ margin: 0, color: '#666', fontSize: '0.9rem' }}>
            Mouse óptico sem fio com design ergonômico
          </p>
        </Card>
      </Container>
      
      <div style={{ marginTop: '1rem', padding: '1rem', background: '#e8f5e9', borderRadius: '8px' }}>
        <p style={{ margin: 0, fontSize: '0.9rem', color: '#2e7d32' }}>
          <strong>💡 Observe:</strong> Este exemplo mostra como componentes pequenos 
          (<code>Botao</code>, <code>Card</code>, <code>Container</code>) são combinados 
          para criar interfaces mais complexas. Isso é <strong>composição</strong>!
        </p>
      </div>
    </div>
  )
}

export default Exemplo6Composition



