import { useState } from 'react'

/**
 * Exemplo 4: Aplicação Completa - Gerenciador de Produtos
 * 
 * Este exemplo combina todos os conceitos aprendidos:
 * - Múltiplos estados
 * - Componentes com props
 * - Listas e renderização condicional
 * - Eventos e formulários
 * - Imutabilidade
 */

/**
 * Componente para exibir um produto
 */
function ProductCard({ product, onAddToCart, onRemoveFromCart, inCart }) {
  return (
    <div style={{
      border: '1px solid #ddd',
      borderRadius: '8px',
      padding: '15px',
      margin: '10px',
      backgroundColor: '#fff',
      boxShadow: '0 2px 4px rgba(0,0,0,0.1)',
      width: '250px',
      display: 'inline-block'
    }}>
      <h3 style={{ margin: '0 0 10px 0', color: '#333' }}>
        {product.name}
      </h3>
      <p style={{ 
        fontSize: '24px', 
        fontWeight: 'bold',
        color: '#4CAF50',
        margin: '10px 0'
      }}>
        R$ {product.price.toFixed(2)}
      </p>
      <p style={{ 
        color: '#666', 
        fontSize: '14px',
        margin: '10px 0'
      }}>
        {product.description}
      </p>
      
      {inCart ? (
        <button
          onClick={() => onRemoveFromCart(product.id)}
          style={{
            width: '100%',
            padding: '10px',
            backgroundColor: '#ff6b6b',
            color: 'white',
            border: 'none',
            borderRadius: '4px',
            cursor: 'pointer',
            fontSize: '14px',
            fontWeight: 'bold'
          }}
        >
          Remover do Carrinho
        </button>
      ) : (
        <button
          onClick={() => onAddToCart(product.id)}
          style={{
            width: '100%',
            padding: '10px',
            backgroundColor: '#4CAF50',
            color: 'white',
            border: 'none',
            borderRadius: '4px',
            cursor: 'pointer',
            fontSize: '14px',
            fontWeight: 'bold'
          }}
        >
          Adicionar ao Carrinho
        </button>
      )}
    </div>
  )
}

/**
 * Componente para exibir o carrinho
 */
function Cart({ cart, products, onRemoveFromCart, onClearCart }) {
  // Calcula o total
  const total = cart.reduce((sum, itemId) => {
    const product = products.find(p => p.id === itemId)
    return sum + (product ? product.price : 0)
  }, 0)

  // Agrupa produtos por ID e conta quantidades
  const cartItems = cart.reduce((acc, itemId) => {
    acc[itemId] = (acc[itemId] || 0) + 1
    return acc
  }, {})

  return (
    <div style={{
      border: '2px solid #4CAF50',
      borderRadius: '8px',
      padding: '20px',
      margin: '20px 0',
      backgroundColor: '#f9f9f9'
    }}>
      <h2 style={{ marginTop: 0 }}>
        🛒 Carrinho de Compras
      </h2>
      
      {cart.length === 0 ? (
        <p style={{ color: '#999', fontStyle: 'italic' }}>
          Seu carrinho está vazio
        </p>
      ) : (
        <>
          <ul style={{ listStyle: 'none', padding: 0 }}>
            {Object.entries(cartItems).map(([productId, quantity]) => {
              const product = products.find(p => p.id === parseInt(productId))
              if (!product) return null
              
              return (
                <li 
                  key={productId}
                  style={{
                    padding: '10px',
                    margin: '5px 0',
                    backgroundColor: '#fff',
                    borderRadius: '4px',
                    display: 'flex',
                    justifyContent: 'space-between',
                    alignItems: 'center'
                  }}
                >
                  <span>
                    {product.name} x{quantity}
                  </span>
                  <div style={{ display: 'flex', alignItems: 'center', gap: '10px' }}>
                    <span style={{ fontWeight: 'bold', color: '#4CAF50' }}>
                      R$ {(product.price * quantity).toFixed(2)}
                    </span>
                    <button
                      onClick={() => onRemoveFromCart(parseInt(productId))}
                      style={{
                        padding: '5px 10px',
                        backgroundColor: '#ff6b6b',
                        color: 'white',
                        border: 'none',
                        borderRadius: '4px',
                        cursor: 'pointer',
                        fontSize: '12px'
                      }}
                    >
                      Remover
                    </button>
                  </div>
                </li>
              )
            })}
          </ul>
          
          <div style={{
            marginTop: '20px',
            paddingTop: '20px',
            borderTop: '2px solid #ddd',
            display: 'flex',
            justifyContent: 'space-between',
            alignItems: 'center'
          }}>
            <strong style={{ fontSize: '20px' }}>
              Total: R$ {total.toFixed(2)}
            </strong>
            <button
              onClick={onClearCart}
              style={{
                padding: '10px 20px',
                backgroundColor: '#ff6b6b',
                color: 'white',
                border: 'none',
                borderRadius: '4px',
                cursor: 'pointer',
                fontSize: '14px',
                fontWeight: 'bold'
              }}
            >
              Limpar Carrinho
            </button>
          </div>
        </>
      )}
    </div>
  )
}

/**
 * Componente principal da aplicação
 */
function App() {
  // Lista de produtos disponíveis
  const [products] = useState([
    {
      id: 1,
      name: 'Notebook Gamer',
      price: 3500.00,
      description: 'Notebook potente para jogos e trabalho'
    },
    {
      id: 2,
      name: 'Mouse Wireless',
      price: 150.00,
      description: 'Mouse ergonômico sem fio'
    },
    {
      id: 3,
      name: 'Teclado Mecânico',
      price: 450.00,
      description: 'Teclado com switches mecânicos'
    },
    {
      id: 4,
      name: 'Monitor 4K',
      price: 2200.00,
      description: 'Monitor ultra HD de 27 polegadas'
    }
  ])

  // Estado do carrinho (array de IDs de produtos)
  const [cart, setCart] = useState([])

  /**
   * Adiciona produto ao carrinho
   * Cria novo array (imutabilidade!)
   */
  function handleAddToCart(productId) {
    setCart([...cart, productId])
  }

  /**
   * Remove uma unidade do produto do carrinho
   */
  function handleRemoveFromCart(productId) {
    const index = cart.indexOf(productId)
    if (index > -1) {
      // Cria novo array sem o item no índice encontrado
      setCart([...cart.slice(0, index), ...cart.slice(index + 1)])
    }
  }

  /**
   * Limpa todo o carrinho
   */
  function handleClearCart() {
    setCart([])
  }

  return (
    <div style={{
      padding: '20px',
      fontFamily: 'Arial, sans-serif',
      maxWidth: '1200px',
      margin: '0 auto'
    }}>
      <h1 style={{ textAlign: 'center', color: '#333' }}>
        🛍️ Loja de Eletrônicos
      </h1>

      {/* Carrinho */}
      <Cart
        cart={cart}
        products={products}
        onRemoveFromCart={handleRemoveFromCart}
        onClearCart={handleClearCart}
      />

      {/* Lista de produtos */}
      <div>
        <h2>Produtos Disponíveis</h2>
        <div style={{ display: 'flex', flexWrap: 'wrap', justifyContent: 'center' }}>
          {products.map(product => (
            <ProductCard
              key={product.id}
              product={product}
              onAddToCart={handleAddToCart}
              onRemoveFromCart={handleRemoveFromCart}
              inCart={cart.includes(product.id)}
            />
          ))}
        </div>
      </div>

      {/* Informações */}
      <div style={{
        marginTop: '40px',
        padding: '15px',
        backgroundColor: '#e3f2fd',
        borderRadius: '8px',
        fontSize: '14px'
      }}>
        <strong>📚 Conceitos demonstrados neste exemplo:</strong>
        <ul>
          <li>Múltiplos estados gerenciados com useState</li>
          <li>Componentes reutilizáveis (ProductCard, Cart)</li>
          <li>Props para passar dados e funções</li>
          <li>Renderização condicional (botão muda baseado em inCart)</li>
          <li>Renderização de listas com .map()</li>
          <li>Imutabilidade (sempre criando novos arrays)</li>
          <li>Composição de componentes</li>
        </ul>
      </div>
    </div>
  )
}

export default App

