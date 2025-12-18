import { useState, useEffect } from 'react'

export function Exemplo4DataFetch() {
  const [id, setId] = useState(1)
  const [user, setUser] = useState(null)
  const [loading, setLoading] = useState(false)

  useEffect(() => {
    setLoading(true)
    setUser(null)

    fetch(`https://jsonplaceholder.typicode.com/users/${id}`)
      .then(res => res.json())
      .then(data => {
        setUser(data)
        setLoading(false)
      })
      .catch(err => {
        console.error(err)
        setLoading(false)
      })

    // Em app real, poderíamos cancelar o fetch aqui se o id mudasse rápido
  }, [id])

  return (
    <div style={{ padding: '20px', border: '1px solid #ccc', borderRadius: '8px' }}>
      <div style={{ marginBottom: '15px' }}>
        <label>User ID: </label>
        <button disabled={id <= 1} onClick={() => setId(id - 1)}>-</button>
        <span style={{ margin: '0 10px', fontWeight: 'bold' }}>{id}</span>
        <button disabled={id >= 10} onClick={() => setId(id + 1)}>+</button>
      </div>

      {loading && <p>Carregando...</p>}
      
      {user && !loading && (
        <div style={{ background: '#f5f5f5', padding: '10px', borderRadius: '4px' }}>
          <strong>Nome:</strong> {user.name}<br/>
          <strong>Email:</strong> {user.email}<br/>
          <strong>Empresa:</strong> {user.company?.name}
        </div>
      )}
    </div>
  )
}

