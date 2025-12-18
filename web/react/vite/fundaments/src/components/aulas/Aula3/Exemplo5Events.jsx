import React, { useState } from 'react';

export function Exemplo5Events() {
  const [lastEvent, setLastEvent] = useState('Nenhum evento ainda');

  const handleClick = (e) => {
    // e é um SyntheticEvent
    setLastEvent(`Click em: ${e.target.tagName}`);
    console.log('Event:', e);
  };

  const handleMouseEnter = () => {
    setLastEvent('Mouse entrou na área!');
  };

  const handleInputChange = (e) => {
    setLastEvent(`Digitando: ${e.target.value}`);
  };

  return (
    <div>
      <h3>Eventos (SyntheticEvent)</h3>
      
      <div 
        style={{ 
          padding: '20px', 
          background: '#e0f7fa', 
          marginBottom: '15px', 
          borderRadius: '4px',
          cursor: 'pointer'
        }}
        onClick={handleClick}
        onMouseEnter={handleMouseEnter}
      >
        <p>Interaja comigo! (Clique ou passe o mouse)</p>
      </div>

      <input 
        type="text" 
        placeholder="Digite para ver o evento change"
        onChange={handleInputChange}
        style={{ width: '100%', padding: '8px', marginBottom: '10px' }}
      />

      <div style={{ borderTop: '1px solid #ddd', paddingTop: '10px' }}>
        <strong>Último evento:</strong> <span style={{ color: '#00796b' }}>{lastEvent}</span>
      </div>
    </div>
  );
}

