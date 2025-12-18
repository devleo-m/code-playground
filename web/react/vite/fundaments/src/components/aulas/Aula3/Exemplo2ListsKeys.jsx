import React, { useState } from 'react';

export function Exemplo2ListsKeys() {
  const [items, setItems] = useState([
    { id: 1, text: 'Aprender React' },
    { id: 2, text: 'Praticar useEffect' },
    { id: 3, text: 'Entender Keys' }
  ]);

  const removeItem = (id) => {
    setItems(items.filter(item => item.id !== id));
  };

  const addItem = () => {
    const newItem = {
      id: Date.now(),
      text: `Nova Tarefa ${items.length + 1}`
    };
    setItems([...items, newItem]);
  };

  return (
    <div>
      <h3>Listas e Keys</h3>
      <button onClick={addItem} style={{ marginBottom: '10px' }}>Adicionar Item</button>
      
      <ul style={{ listStyle: 'none', padding: 0 }}>
        {items.map((item) => (
          <li 
            key={item.id} 
            style={{ 
              display: 'flex', 
              justifyContent: 'space-between',
              padding: '8px',
              borderBottom: '1px solid #eee',
              alignItems: 'center'
            }}
          >
            <span>{item.text}</span>
            <button 
              onClick={() => removeItem(item.id)}
              style={{ background: '#ff4444', color: 'white', border: 'none', padding: '4px 8px', borderRadius: '4px', cursor: 'pointer' }}
            >
              Remover
            </button>
          </li>
        ))}
      </ul>
      <p style={{ fontSize: '0.8rem', color: '#666' }}>
        Observe que usamos <code>item.id</code> como key, não o index!
      </p>
    </div>
  );
}

