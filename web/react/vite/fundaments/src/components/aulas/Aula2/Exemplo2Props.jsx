// Exemplo 2: Componente com Props
function Saudacao({ nome, idade }) {
  return (
    <div style={{ padding: '1rem', background: '#e8f5e9', borderRadius: '8px', border: '2px solid #4caf50' }}>
      <h3 style={{ color: '#2c3e50', margin: '0 0 0.5rem 0' }}>Olá, {nome}!</h3>
      <p style={{ color: '#555', margin: 0 }}>Você tem {idade} anos.</p>
      <p style={{ color: '#555', margin: '0.5rem 0 0 0', fontSize: '0.9rem' }}>
        Você nasceu em {new Date().getFullYear() - idade}.
      </p>
    </div>
  )
}

function Exemplo2Props() {
  return (
    <div>
      <Saudacao nome="Maria" idade={25} />
      <div style={{ marginTop: '1rem' }}>
        <Saudacao nome="João" idade={30} />
      </div>
    </div>
  )
}

export default Exemplo2Props



