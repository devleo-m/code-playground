// Exemplo 2: Componente com Props
// 💡 O que são Props?
// Props (propriedades) são como argumentos de uma função. 
// Elas permitem passar dados de um componente pai para um filho.

// Aqui usamos desestruturação { nome, idade } para pegar as propriedades diretamente
function Saudacao({ nome, idade }) {
  // Calculando idade de nascimento dinamicamente
  const anoAtual = new Date().getFullYear()
  const anoNascimento = anoAtual - idade

  return (
    <div style={{ padding: '1rem', background: '#e8f5e9', borderRadius: '8px', border: '2px solid #4caf50' }}>
      {/* Usamos chaves {} para inserir valores JavaScript no JSX */}
      <h3 style={{ color: '#2c3e50', margin: '0 0 0.5rem 0' }}>Olá, {nome}!</h3>
      
      <p style={{ color: '#555', margin: 0 }}>Você tem {idade} anos.</p>
      
      <p style={{ color: '#555', margin: '0.5rem 0 0 0', fontSize: '0.9rem' }}>
        Você nasceu por volta de {anoNascimento}.
      </p>
    </div>
  )
}

function Exemplo2Props() {
  return (
    <div>
      {/* Passando strings com aspas e números com chaves {} */}
      <Saudacao nome="Maria" idade={25} />
      
      <div style={{ marginTop: '1rem' }}>
        <Saudacao nome="João" idade={30} />
      </div>
    </div>
  )
}

export default Exemplo2Props
