import React, { useState } from 'react';

// Componente que implementa a lógica (Render Prop Pattern)
function MouseTracker({ render }) {
  const [position, setPosition] = useState({ x: 0, y: 0 });

  const handleMouseMove = (event) => {
    // Pegando posição relativa ao elemento pai para o exemplo não quebrar o layout
    const bounds = event.currentTarget.getBoundingClientRect();
    setPosition({
      x: event.clientX - bounds.left,
      y: event.clientY - bounds.top
    });
  };

  return (
    <div 
      style={{ height: '200px', border: '2px dashed #ccc', position: 'relative', cursor: 'crosshair' }} 
      onMouseMove={handleMouseMove}
    >
      {render(position)}
    </div>
  );
}

export function Exemplo3RenderProps() {
  return (
    <div>
      <h3>Render Props</h3>
      <p>Mova o mouse dentro da área pontilhada:</p>
      
      <MouseTracker render={({ x, y }) => (
        <div style={{ 
          position: 'absolute', 
          left: x, 
          top: y, 
          pointerEvents: 'none',
          background: 'rgba(0,0,0,0.8)',
          color: 'white',
          padding: '4px 8px',
          borderRadius: '4px',
          transform: 'translate(10px, 10px)'
        }}>
          x: {Math.round(x)}, y: {Math.round(y)}
        </div>
      )} />
    </div>
  );
}

