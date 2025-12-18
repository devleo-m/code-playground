import React, { useState, useEffect } from 'react';

export function Exemplo1Lifecycle() {
  const [showTimer, setShowTimer] = useState(false);

  return (
    <div>
      <h3>Ciclo de Vida com useEffect</h3>
      <button onClick={() => setShowTimer(!showTimer)}>
        {showTimer ? 'Esconder Timer' : 'Mostrar Timer'}
      </button>
      
      {showTimer && <Timer />}
    </div>
  );
}

function Timer() {
  const [seconds, setSeconds] = useState(0);

  useEffect(() => {
    // Mounting
    console.log('Timer montado!');
    
    const intervalId = setInterval(() => {
      setSeconds(s => s + 1);
      console.log('Timer tick');
    }, 1000);

    // Unmounting (Cleanup)
    return () => {
      console.log('Timer desmontado! Limpando intervalo...');
      clearInterval(intervalId);
    };
  }, []); // Array vazio = apenas na montagem/desmontagem

  return (
    <div style={{ marginTop: '10px', padding: '10px', background: '#f0f0f0', borderRadius: '4px' }}>
      <p>Tempo: {seconds} segundos</p>
      <small>Abra o console para ver os logs!</small>
    </div>
  );
}

