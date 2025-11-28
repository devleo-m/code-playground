import { useState } from 'react'

/**
 * Exemplo 2: Lista de Tarefas Básica
 * 
 * Este exemplo demonstra:
 * - Estado com arrays
 * - Input controlado (controlled input)
 * - Spread operator para criar novos arrays
 * - Renderização de listas com .map()
 * - Key prop para identificação de elementos
 */
function TodoList() {
  // Estado para armazenar array de tarefas
  const [todos, setTodos] = useState([])
  
  // Estado para o input (input controlado)
  const [input, setInput] = useState('')

  /**
   * Função para adicionar nova tarefa
   * Usa spread operator para criar novo array (imutabilidade!)
   */
  function addTodo() {
    // Verifica se input não está vazio (trim remove espaços)
    if (input.trim() !== '') {
      // Cria NOVO array com todas as tarefas + nova tarefa
      setTodos([...todos, input])
      // Limpa o input
      setInput('')
    }
  }

  /**
   * Função para remover tarefa pelo índice
   */
  function removeTodo(index) {
    // Filtra o array, removendo o item no índice especificado
    setTodos(todos.filter((_, i) => i !== index))
  }

  /**
   * Função para lidar com Enter no input
   */
  function handleKeyPress(e) {
    if (e.key === 'Enter') {
      addTodo()
    }
  }

  return (
    <div style={{ 
      padding: '20px', 
      maxWidth: '500px',
      margin: '0 auto',
      fontFamily: 'Arial, sans-serif'
    }}>
      <h2>Minhas Tarefas</h2>
      
      {/* Input controlado */}
      <div style={{ marginBottom: '20px' }}>
        <input
          type="text"
          value={input}
          onChange={(e) => setInput(e.target.value)}
          onKeyPress={handleKeyPress}
          placeholder="Digite uma nova tarefa..."
          style={{
            padding: '10px',
            fontSize: '16px',
            width: '70%',
            marginRight: '10px',
            border: '1px solid #ddd',
            borderRadius: '4px'
          }}
        />
        <button 
          onClick={addTodo}
          style={{
            padding: '10px 20px',
            fontSize: '16px',
            backgroundColor: '#4CAF50',
            color: 'white',
            border: 'none',
            borderRadius: '4px',
            cursor: 'pointer'
          }}
        >
          Adicionar
        </button>
      </div>

      {/* Lista de tarefas */}
      {todos.length === 0 ? (
        <p style={{ color: '#999', fontStyle: 'italic' }}>
          Nenhuma tarefa ainda. Adicione uma acima!
        </p>
      ) : (
        <ul style={{ listStyle: 'none', padding: 0 }}>
          {todos.map((todo, index) => (
            <li 
              key={index}
              style={{
                padding: '10px',
                margin: '5px 0',
                backgroundColor: '#f5f5f5',
                borderRadius: '4px',
                display: 'flex',
                justifyContent: 'space-between',
                alignItems: 'center'
              }}
            >
              <span>{todo}</span>
              <button
                onClick={() => removeTodo(index)}
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
            </li>
          ))}
        </ul>
      )}

      <p style={{ marginTop: '20px', color: '#666', fontSize: '14px' }}>
        Total de tarefas: {todos.length}
      </p>
    </div>
  )
}

export default TodoList

