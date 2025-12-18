// Exemplo 6: Composition (Composição)
// 💡 O poder da Composição
// Em vez de herança (como em OOP clássica), o React usa composição.
// Criamos componentes pequenos e genéricos e os "encaixamos" para criar componentes maiores.

// 1. Botão Genérico
// Recebe children (texto/ícone) e variant (cor)
function Botao({ children, onClick, variant = 'primary' }) {
  const colors = {
    primary: '#2196f3',
    secondary: '#757575',
    danger: '#f44336'
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
        background: colors[variant],
        color: 'white'
      }}
    >
      {children}
    </button>
  )
}

// 2. Card Genérico
// Define a ESTRUTURA, mas não o CONTEÚDO específico.
// Aceita `title` (opcional), `children` (corpo) e `footer` (rodapé opcional)
function Card({ title, children, footer }) {
  return (
    <div style={{
      border: '2px solid #e0e0e0',
      borderRadius: '8px',
      overflow: 'hidden',
      marginBottom: '1rem',
      background: 'white'
    }}>
      {/* Renderiza título apenas se existir */}
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
      
      {/* Área de conteúdo flexível */}
      <div style={{ padding: '1rem' }}>
        {children}
      </div>

      {/* Rodapé flexível */}
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

// 3. Container
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

// Juntando tudo!
function Exemplo6Composition() {
  return (
    <div>
      <Container title="Loja de Eletrônicos (Exemplo de Composição)">
        
        {/* Usando Card para um Notebook */}
        <Card 
          title="💻 Notebook Pro"
          footer={
            <div style={{ display: 'flex', gap: '0.5rem', justifyContent: 'flex-end' }}>
              <Botao variant="secondary">Detalhes</Botao>
              <Botao variant="primary">Comprar</Botao>
            </div>
          }
        >
          <p style={{ margin: '0 0 0.5rem 0', fontWeight: 'bold' }}>R$ 4.500,00</p>
          <p style={{ margin: 0, color: '#666', fontSize: '0.9rem' }}>
            Processador i7, 16GB RAM, SSD 512GB. Ideal para desenvolvedores.
          </p>
        </Card>

        {/* Usando o MESMO componente Card para um Mouse, mas com conteúdo diferente */}
        <Card 
          title="🖱️ Mouse Gamer"
          footer={
            <div style={{ display: 'flex', gap: '0.5rem', justifyContent: 'flex-end' }}>
              <Botao variant="danger">Remover</Botao>
            </div>
          }
        >
          <p style={{ margin: '0 0 0.5rem 0', fontWeight: 'bold' }}>R$ 150,00</p>
          <p style={{ margin: 0, color: '#666', fontSize: '0.9rem' }}>
            Mouse RGB com 6 botões programáveis.
          </p>
        </Card>

      </Container>
      
      <div style={{ marginTop: '1rem', padding: '1rem', background: '#e8f5e9', borderRadius: '8px' }}>
        <p style={{ margin: 0, fontSize: '0.9rem', color: '#2e7d32' }}>
          <strong>💡 Dica:</strong> Veja como reutilizamos <code>Card</code> e <code>Botao</code> 
          para criar interfaces diferentes apenas passando props e children diferentes!
        </p>
      </div>
    </div>
  )
}

export default Exemplo6Composition
