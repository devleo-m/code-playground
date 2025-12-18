import React, { useRef, useState } from 'react';

export function Exemplo4Refs() {
  // 1. Ref para acessar elemento DOM
  const inputRef = useRef(null);
  
  // 2. Ref para persistir valor sem re-renderizar
  const renderCount = useRef(0);
  const [inputValue, setInputValue] = useState('');

  // Incrementa a cada render
  renderCount.current++;

  const focusInput = () => {
    if (inputRef.current) {
      inputRef.current.focus();
      inputRef.current.style.backgroundColor = '#e3f2fd';
      setTimeout(() => {
        if (inputRef.current) inputRef.current.style.backgroundColor = 'white';
      }, 1000);
    }
  };

  return (
    <div>
      <h3>Refs: DOM e Persistência</h3>
      
      <div style={{ marginBottom: '15px' }}>
        <input 
          ref={inputRef}
          type="text" 
          value={inputValue}
          onChange={(e) => setInputValue(e.target.value)}
          placeholder="Digite algo..."
          style={{ padding: '8px', marginRight: '8px' }}
        />
        <button onClick={focusInput}>Focar Input</button>
      </div>

      <div style={{ background: '#f5f5f5', padding: '10px', borderRadius: '4px' }}>
        <p>Você digitou: <strong>{inputValue}</strong></p>
        <p>O componente renderizou: <strong>{renderCount.current}</strong> vezes</p>
        <small>(Note que digitar atualiza o state e causa render, mas o contador usa ref para persistir)</small>
      </div>
    </div>
  );
}

