import { useState } from 'react'

// Exemplo 2: Lista de Tarefas (Todo List)
// 💡 Gerenciando Listas
// Aqui aprendemos a:
// 1. Usar Arrays no State
// 2. Controlar Inputs (o React controla o que é digitado)
// 3. Renderizar Listas com .map()

function Exemplo2ListaTarefas() {
  const [tarefas, setTarefas] = useState([])
  const [input, setInput] = useState('')

  function adicionarTarefa() {
    if (input.trim() !== '') {
      // ⚠️ Imutabilidade: Nunca faça tarefas.push()
      // Crie um NOVO array com os itens antigos + o novo
      setTarefas([...tarefas, input])
      setInput('') // Limpa o input
    }
  }

  function removerTarefa(index) {
    // Filtramos o array para criar um novo sem o item removido
    setTarefas(tarefas.filter((_, i) => i !== index))
  }

  return (
    <div style={{ maxWidth: '500px', margin: '0 auto' }}>
      <h3>Minhas Tarefas</h3>
      
      <div style={{ display: 'flex', gap: '0.5rem', marginBottom: '1rem' }}>
        <input
          type="text"
          value={input}
          onChange={(e) => setInput(e.target.value)}
          placeholder="Digite uma tarefa..."
          style={{
            flex: 1,
            padding: '10px',
            borderRadius: '4px',
            border: '1px solid #ddd'
          }}
          onKeyPress={(e) => e.key === 'Enter' && adicionarTarefa()}
        />
        <button 
          onClick={adicionarTarefa}
          style={{
            padding: '10px 20px',
            background: '#2196f3',
            color: 'white',
            border: 'none',
            borderRadius: '4px',
            cursor: 'pointer'
          }}
        >
          Adicionar
        </button>
      </div>

      <ul style={{ listStyle: 'none', padding: 0 }}>
        {tarefas.map((tarefa, index) => (
          <li 
            key={index}
            style={{
              padding: '10px',
              borderBottom: '1px solid #eee',
              display: 'flex',
              justifyContent: 'space-between',
              alignItems: 'center'
            }}
          >
            <span>{tarefa}</span>
            <button
              onClick={() => removerTarefa(index)}
              style={{
                color: '#ff6b6b',
                background: 'none',
                border: 'none',
                cursor: 'pointer',
                fontWeight: 'bold'
              }}
            >
              Excluir
            </button>
          </li>
        ))}
        {tarefas.length === 0 && (
          <li style={{ color: '#999', textAlign: 'center', fontStyle: 'italic' }}>
            Nenhuma tarefa... que tal adicionar uma? 🤔
          </li>
        )}
      </ul>
    </div>
  )
}

export default Exemplo2ListaTarefas

