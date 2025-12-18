import { useState } from 'react'

// Exemplo 4: App Completo (Mini Loja)
// 💡 Juntando tudo!
// State, Props, Listas, Condicionais e Eventos.

function ProductCard({ product, inCart, onToggleCart }) {
  return (
    <div style={{
      border: '1px solid #eee',
      padding: '1rem',
      borderRadius: '8px',
      background: 'white',
      display: 'flex',
      flexDirection: 'column',
      justifyContent: 'space-between'
    }}>
      <div>
        <h4 style={{ margin: '0 0 0.5rem 0' }}>{product.name}</h4>
        <p style={{ color: '#4caf50', fontWeight: 'bold', fontSize: '1.2rem', margin: '0 0 0.5rem 0' }}>
          R$ {product.price}
        </p>
      </div>
      <button
        onClick={() => onToggleCart(product.id)}
        style={{
          width: '100%',
          padding: '8px',
          border: 'none',
          borderRadius: '4px',
          cursor: 'pointer',
          background: inCart ? '#ff5252' : '#2196f3',
          color: 'white'
        }}
      >
        {inCart ? 'Remover' : 'Comprar'}
      </button>
    </div>
  )
}

function Exemplo4AppCompleto() {
  const [products] = useState([
    { id: 1, name: 'Notebook Gamer', price: 3500 },
    { id: 2, name: 'Mouse Sem Fio', price: 150 },
    { id: 3, name: 'Teclado Mecânico', price: 450 },
    { id: 4, name: 'Monitor 4K', price: 2200 }
  ])

  const [cartIds, setCartIds] = useState([])

  const toggleCart = (id) => {
    if (cartIds.includes(id)) {
      setCartIds(cartIds.filter(itemId => itemId !== id))
    } else {
      setCartIds([...cartIds, id])
    }
  }

  const total = cartIds.reduce((sum, id) => {
    const product = products.find(p => p.id === id)
    return sum + (product ? product.price : 0)
  }, 0)

  return (
    <div>
      <div style={{ 
        background: '#e3f2fd', 
        padding: '1rem', 
        borderRadius: '8px', 
        marginBottom: '1rem',
        display: 'flex',
        justifyContent: 'space-between',
        alignItems: 'center'
      }}>
        <h3 style={{ margin: 0, color: '#1565c0' }}>🛍️ Mini Loja</h3>
        <div>
          <strong>Total: R$ {total.toFixed(2)}</strong> ({cartIds.length} itens)
        </div>
      </div>

      <div style={{ 
        display: 'grid', 
        gridTemplateColumns: 'repeat(auto-fill, minmax(200px, 1fr))', 
        gap: '1rem' 
      }}>
        {products.map(product => (
          <ProductCard
            key={product.id}
            product={product}
            inCart={cartIds.includes(product.id)}
            onToggleCart={toggleCart}
          />
        ))}
      </div>
    </div>
  )
}

export default Exemplo4AppCompleto

