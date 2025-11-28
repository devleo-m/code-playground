// Dados dos tópicos de estudo
export const topics = [
  {
    id: 1,
    title: "Componentes",
    description: "Aprenda sobre componentes React, a base de tudo!",
    content: `
      <h2>O que são Componentes?</h2>
      <p>Componentes são como blocos de construção para sua aplicação React. Eles permitem que você divida a interface em partes reutilizáveis.</p>
      
      <h3>Componente Funcional</h3>
      <p>É a forma mais simples de criar um componente:</p>
      <pre><code>function MeuComponente() {
  return &lt;div&gt;Olá Mundo!&lt;/div&gt;;
}</code></pre>
      
      <h3>Por que usar Componentes?</h3>
      <ul>
        <li>Reutilização de código</li>
        <li>Organização melhor do projeto</li>
        <li>Manutenção mais fácil</li>
      </ul>
    `
  },
  {
    id: 2,
    title: "Props",
    description: "Entenda como passar dados entre componentes",
    content: `
      <h2>O que são Props?</h2>
      <p>Props (propriedades) são como argumentos que você passa para um componente, permitindo que ele receba dados do componente pai.</p>
      
      <h3>Exemplo Básico</h3>
      <pre><code>function Saudacao({ nome }) {
  return &lt;h1&gt;Olá, {nome}!&lt;/h1&gt;;
}

// Uso:
&lt;Saudacao nome="Maria" /&gt;</code></pre>
      
      <h3>Dicas Importantes</h3>
      <ul>
        <li>Props são somente leitura (não podem ser alteradas)</li>
        <li>Você pode passar qualquer tipo de dado</li>
        <li>Use props para tornar componentes reutilizáveis</li>
      </ul>
    `
  },
  {
    id: 3,
    title: "Estado (State)",
    description: "Aprenda a gerenciar dados que mudam com useState",
    content: `
      <h2>O que é Estado?</h2>
      <p>Estado permite que componentes "lembrem" de informações e reajam a mudanças. Quando o estado muda, o componente é renderizado novamente.</p>
      
      <h3>useState Hook</h3>
      <pre><code>import { useState } from 'react';

function Contador() {
  const [contador, setContador] = useState(0);
  
  return (
    &lt;div&gt;
      &lt;p&gt;Você clicou {contador} vezes&lt;/p&gt;
      &lt;button onClick={() =&gt; setContador(contador + 1)}&gt;
        Clique aqui
      &lt;/button&gt;
    &lt;/div&gt;
  );
}</code></pre>
      
      <h3>Como Funciona?</h3>
      <ul>
        <li>useState retorna um array com [valor, função para atualizar]</li>
        <li>Quando você atualiza o estado, o componente re-renderiza</li>
        <li>O estado é privado para cada componente</li>
      </ul>
    `
  },
  {
    id: 4,
    title: "Eventos",
    description: "Aprenda a lidar com cliques, formulários e mais",
    content: `
      <h2>Eventos em React</h2>
      <p>React tem seu próprio sistema de eventos chamado SyntheticEvent. É muito similar aos eventos do DOM, mas funciona em todos os navegadores.</p>
      
      <h3>Exemplo de Clique</h3>
      <pre><code>function Botao() {
  const handleClick = () => {
    alert('Botão clicado!');
  };
  
  return &lt;button onClick={handleClick}&gt;Clique aqui&lt;/button&gt;;
}</code></pre>
      
      <h3>Eventos Comuns</h3>
      <ul>
        <li>onClick - quando clica</li>
        <li>onChange - quando o valor muda (inputs)</li>
        <li>onSubmit - quando submete um formulário</li>
        <li>onMouseOver - quando passa o mouse</li>
      </ul>
    `
  },
  {
    id: 5,
    title: "Renderização Condicional",
    description: "Mostre conteúdo diferente baseado em condições",
    content: `
      <h2>Renderização Condicional</h2>
      <p>Você pode mostrar diferentes elementos baseado em condições, usando if/else, operador ternário ou operador &&.</p>
      
      <h3>Operador Ternário</h3>
      <pre><code>function Saudacao({ estaLogado }) {
  return (
    &lt;div&gt;
      {estaLogado ? (
        &lt;h1&gt;Bem-vindo de volta!&lt;/h1&gt;
      ) : (
        &lt;h1&gt;Por favor, faça login&lt;/h1&gt;
      )}
    &lt;/div&gt;
  );
}</code></pre>
      
      <h3>Operador &&</h3>
      <pre><code>{contador > 0 && &lt;p&gt;Você tem {contador} itens&lt;/p&gt;}</code></pre>
    `
  },
  {
    id: 6,
    title: "Listas e Keys",
    description: "Como renderizar listas de dados",
    content: `
      <h2>Renderizando Listas</h2>
      <p>Você pode renderizar múltiplos componentes usando o método map().</p>
      
      <h3>Exemplo Básico</h3>
      <pre><code>function ListaTarefas({ tarefas }) {
  return (
    &lt;ul&gt;
      {tarefas.map((tarefa) => (
        &lt;li key={tarefa.id}&gt;{tarefa.nome}&lt;/li&gt;
      ))}
    &lt;/ul&gt;
  );
}</code></pre>
      
      <h3>Por que Keys são Importantes?</h3>
      <ul>
        <li>Keys ajudam React a identificar quais itens mudaram</li>
        <li>Devem ser únicas entre irmãos</li>
        <li>Melhoram a performance</li>
      </ul>
    `
  }
];

