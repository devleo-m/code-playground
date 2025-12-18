// Exemplo 1: Componente Básico
// 💡 O que é um Componente Funcional?
// É apenas uma função JavaScript que retorna JSX (uma mistura de HTML com JavaScript).
// Pense neles como blocos de construção da sua interface (como peças de LEGO).

function Exemplo1ComponenteBasico() {
  // O return contém o que vai aparecer na tela
  return (
    // 🎨 Estilos Inline:
    // Em React, podemos escrever CSS direto no JavaScript usando um objeto {{ chave: valor }}
    <div style={{ padding: '1rem', background: '#f0f0f0', borderRadius: '8px' }}>
      <h3 style={{ color: '#2c3e50', margin: '0 0 0.5rem 0' }}>Olá, React!</h3>
      <p style={{ color: '#555', margin: 0 }}>Este é meu primeiro componente funcional!</p>
    </div>
  )
}

// Exportamos o componente para que ele possa ser usado em outros arquivos
export default Exemplo1ComponenteBasico
