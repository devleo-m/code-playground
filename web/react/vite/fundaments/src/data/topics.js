// Dados dos tópicos de estudo
export const topics = [
  {
    id: 1,
    title: "Componentes",
    description: "Aprenda sobre componentes React, a base de tudo!",
    content: `
      <h2>O que são Componentes?</h2>
      <p>Componentes são os blocos de construção de aplicações React. Eles nos permitem dividir a interface do usuário em partes independentes e reutilizáveis, e pensar sobre cada parte de forma isolada.</p>
      
      <p>Imagine construir uma casa: você não constrói tudo de uma vez. Você tem tijolos, portas, janelas, telhado - cada um é uma peça separada que pode ser reutilizada. Componentes React funcionam da mesma forma!</p>
      
      <h3>Por que usar Componentes?</h3>
      <ul>
        <li><strong>Reutilização:</strong> Escreva uma vez, use em qualquer lugar</li>
        <li><strong>Organização:</strong> Código mais limpo e fácil de entender</li>
        <li><strong>Manutenção:</strong> Corrija um bug em um lugar, não em vários</li>
        <li><strong>Colaboração:</strong> Diferentes desenvolvedores podem trabalhar em componentes diferentes</li>
        <li><strong>Testabilidade:</strong> Teste cada componente isoladamente</li>
      </ul>
      
      <h2>Componentes Funcionais (Functional Components)</h2>
      <p>Componentes funcionais são uma das formas mais comuns de criar componentes em React. Eles são simplesmente funções JavaScript que retornam JSX.</p>
      
      <h3>Como Funcionam?</h3>
      <p>Um componente funcional é uma função JavaScript que pode ou não receber dados como parâmetros (props). O valor de retorno dessa função é o código JSX que será renderizado na árvore DOM.</p>
      
      <h3>Exemplo Básico</h3>
      <pre><code>// Componente funcional simples
function MeuComponente() {
  return &lt;div&gt;Olá Mundo!&lt;/div&gt;;
}

// Usando arrow function (forma moderna)
const MeuComponente = () => {
  return &lt;div&gt;Olá Mundo!&lt;/div&gt;;
}</code></pre>
      
      <h3>Componente com Props</h3>
      <pre><code>function Saudacao({ nome, idade }) {
  return (
    &lt;div&gt;
      &lt;h1&gt;Olá, {nome}!&lt;/h1&gt;
      &lt;p&gt;Você tem {idade} anos.&lt;/p&gt;
    &lt;/div&gt;
  );
}

// Uso:
&lt;Saudacao nome="Maria" idade={25} /&gt;</code></pre>
      
      <h3>Estado em Componentes Funcionais</h3>
      <p>Componentes funcionais também podem ter estado, gerenciado usando React Hooks (como useState):</p>
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
      
      <h3>Vantagens dos Componentes Funcionais</h3>
      <ul>
        <li>Sintaxe mais simples e limpa</li>
        <li>Mais fáceis de testar</li>
        <li>Melhor performance (com React Hooks)</li>
        <li>São a forma recomendada pela equipe do React</li>
      </ul>
      
      <h2>JSX e TSX</h2>
      <p>JSX (JavaScript XML) e TSX (TypeScript XML) são extensões de sintaxe que permitem escrever HTML-like code dentro de JavaScript/TypeScript.</p>
      
      <h3>O que é JSX?</h3>
      <p>JSX é uma sintaxe especial que parece HTML, mas na verdade é JavaScript. Ele permite que você escreva a estrutura da interface de forma mais intuitiva.</p>
      
      <h3>Exemplo de JSX</h3>
      <pre><code>// Isso é JSX:
const elemento = &lt;h1&gt;Olá, Mundo!&lt;/h1&gt;;

// Por baixo dos panos, React transforma em:
const elemento = React.createElement('h1', null, 'Olá, Mundo!');</code></pre>
      
      <h3>Regras Importantes do JSX</h3>
      <ul>
        <li><strong>Retornar um único elemento raiz:</strong> Use uma &lt;div&gt; ou Fragment (&lt;&gt;&lt;/&gt;)</li>
        <li><strong>Fechar todas as tags:</strong> &lt;img /&gt; e &lt;br /&gt; devem ser auto-fechadas</li>
        <li><strong>Usar className ao invés de class:</strong> class é palavra reservada em JS</li>
        <li><strong>Usar camelCase para eventos:</strong> onClick, onChange, etc.</li>
        <li><strong>Interpolação com chaves:</strong> {variável} para inserir valores JavaScript</li>
      </ul>
      
      <h3>Exemplos Práticos</h3>
      <pre><code>// ✅ Correto - um elemento raiz
function Componente() {
  return (
    &lt;div&gt;
      &lt;h1&gt;Título&lt;/h1&gt;
      &lt;p&gt;Parágrafo&lt;/p&gt;
    &lt;/div&gt;
  );
}

// ✅ Correto - usando Fragment
function Componente() {
  return (
    &lt;&gt;
      &lt;h1&gt;Título&lt;/h1&gt;
      &lt;p&gt;Parágrafo&lt;/p&gt;
    &lt;/&gt;
  );
}

// ✅ Interpolação de variáveis
function Saudacao({ nome }) {
  const mensagem = "Bem-vindo";
  return &lt;h1&gt;{mensagem}, {nome}!&lt;/h1&gt;;
}

// ✅ Expressões JavaScript
function Calculadora() {
  const a = 5;
  const b = 3;
  return &lt;p&gt;{a} + {b} = {a + b}&lt;/p&gt;;
}</code></pre>
      
      <h3>O que é TSX?</h3>
      <p>TSX é JSX com TypeScript. Ele adiciona tipagem estática ao JSX, ajudando a prevenir erros e tornando o código mais seguro e fácil de manter.</p>
      
      <pre><code>// Exemplo TSX com tipagem
interface Props {
  nome: string;
  idade: number;
}

function Saudacao({ nome, idade }: Props) {
  return (
    &lt;div&gt;
      &lt;h1&gt;Olá, {nome}!&lt;/h1&gt;
      &lt;p&gt;Você tem {idade} anos.&lt;/p&gt;
    &lt;/div&gt;
  );
}</code></pre>
      
      <h3>Por que JSX/TSX é Importante?</h3>
      <ul>
        <li><strong>Legibilidade:</strong> Código mais fácil de ler e entender</li>
        <li><strong>Produtividade:</strong> Escreva menos código</li>
        <li><strong>Segurança:</strong> React previne ataques XSS automaticamente</li>
        <li><strong>Ferramentas:</strong> Melhor suporte de IDEs e ferramentas de desenvolvimento</li>
      </ul>
      
      <h2>Props vs State</h2>
      <p>Props (propriedades) e State (estado) são ambos objetos JavaScript simples. Ambos contêm informações que influenciam a renderização do componente, mas são diferentes em um aspecto importante.</p>
      
      <h3>Props (Propriedades)</h3>
      <p>Props são passadas para o componente (similar a parâmetros de função). Elas vêm de fora do componente e são <strong>somente leitura</strong>.</p>
      
      <pre><code>// Props são passadas de fora
function Saudacao({ nome }) {
  return &lt;h1&gt;Olá, {nome}!&lt;/h1&gt;;
}

// O componente pai passa as props
function App() {
  return &lt;Saudacao nome="Maria" /&gt;;
}</code></pre>
      
      <h3>State (Estado)</h3>
      <p>State é gerenciado dentro do componente (similar a variáveis declaradas dentro de uma função). Ele pode ser modificado e causa re-renderização quando muda.</p>
      
      <pre><code>import { useState } from 'react';

function Contador() {
  // State é gerenciado dentro do componente
  const [contador, setContador] = useState(0);
  
  return (
    &lt;div&gt;
      &lt;p&gt;Contador: {contador}&lt;/p&gt;
      &lt;button onClick={() =&gt; setContador(contador + 1)}&gt;
        Incrementar
      &lt;/button&gt;
    &lt;/div&gt;
  );
}</code></pre>
      
      <h3>Comparação: Props vs State</h3>
      <table style="width: 100%; border-collapse: collapse; margin: 20px 0;">
        <tr style="background-color: #f0f0f0;">
          <th style="border: 1px solid #ddd; padding: 8px; text-align: left;">Característica</th>
          <th style="border: 1px solid #ddd; padding: 8px; text-align: left;">Props</th>
          <th style="border: 1px solid #ddd; padding: 8px; text-align: left;">State</th>
        </tr>
        <tr>
          <td style="border: 1px solid #ddd; padding: 8px;"><strong>Origem</strong></td>
          <td style="border: 1px solid #ddd; padding: 8px;">Vem de fora (componente pai)</td>
          <td style="border: 1px solid #ddd; padding: 8px;">Gerenciado dentro do componente</td>
        </tr>
        <tr>
          <td style="border: 1px solid #ddd; padding: 8px;"><strong>Mutabilidade</strong></td>
          <td style="border: 1px solid #ddd; padding: 8px;">Somente leitura (imutável)</td>
          <td style="border: 1px solid #ddd; padding: 8px;">Pode ser modificado</td>
        </tr>
        <tr>
          <td style="border: 1px solid #ddd; padding: 8px;"><strong>Quando usar</strong></td>
          <td style="border: 1px solid #ddd; padding: 8px;">Dados que vêm de fora</td>
          <td style="border: 1px solid #ddd; padding: 8px;">Dados que mudam internamente</td>
        </tr>
        <tr>
          <td style="border: 1px solid #ddd; padding: 8px;"><strong>Exemplo</strong></td>
          <td style="border: 1px solid #ddd; padding: 8px;">Nome do usuário, configurações</td>
          <td style="border: 1px solid #ddd; padding: 8px;">Contador, formulário, toggle</td>
        </tr>
      </table>
      
      <h3>Exemplo Combinando Props e State</h3>
      <pre><code>function ContadorPersonalizado({ valorInicial, cor }) {
  // State interno
  const [contador, setContador] = useState(valorInicial);
  
  return (
    &lt;div style={{ color: cor }}&gt;
      &lt;p&gt;Contador: {contador}&lt;/p&gt;
      &lt;button onClick={() =&gt; setContador(contador + 1)}&gt;
        +1
      &lt;/button&gt;
    &lt;/div&gt;
  );
}

// Uso: props definem comportamento inicial, state controla mudanças
&lt;ContadorPersonalizado valorInicial={10} cor="blue" /&gt;</code></pre>
      
      <h2>Renderização Condicional (Conditional Rendering)</h2>
      <p>No React, você pode criar componentes distintos que encapsulam o comportamento que você precisa. Então, você pode renderizar apenas alguns deles, dependendo do estado da sua aplicação.</p>
      
      <p>A renderização condicional no React funciona da mesma forma que condições funcionam em JavaScript. Use operadores JavaScript como <code>if</code> ou o operador ternário para criar elementos representando o estado atual, e deixe o React atualizar a UI para corresponder a eles.</p>
      
      <h3>Método 1: Operador Ternário</h3>
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
      
      <h3>Método 2: Operador && (AND Lógico)</h3>
      <pre><code>function ListaTarefas({ tarefas }) {
  return (
    &lt;div&gt;
      &lt;h2&gt;Minhas Tarefas&lt;/h2&gt;
      {tarefas.length > 0 && (
        &lt;ul&gt;
          {tarefas.map(tarefa => &lt;li key={tarefa.id}&gt;{tarefa.texto}&lt;/li&gt;)}
        &lt;/ul&gt;
      )}
      {tarefas.length === 0 && (
        &lt;p&gt;Nenhuma tarefa ainda!&lt;/p&gt;
      )}
    &lt;/div&gt;
  );
}</code></pre>
      
      <h3>Método 3: If/Else Tradicional</h3>
      <pre><code>function Perfil({ usuario }) {
  if (!usuario) {
    return &lt;div&gt;Carregando...&lt;/div&gt;;
  }
  
  if (usuario.admin) {
    return (
      &lt;div&gt;
        &lt;h1&gt;Painel Administrativo&lt;/h1&gt;
        &lt;p&gt;Bem-vindo, {usuario.nome}!&lt;/p&gt;
      &lt;/div&gt;
    );
  }
  
  return (
    &lt;div&gt;
      &lt;h1&gt;Perfil do Usuário&lt;/h1&gt;
      &lt;p&gt;Olá, {usuario.nome}!&lt;/p&gt;
    &lt;/div&gt;
  );
}</code></pre>
      
      <h3>Método 4: Variáveis para Armazenar Elementos</h3>
      <pre><code>function LoginButton({ estaLogado, onLogin, onLogout }) {
  let botao;
  
  if (estaLogado) {
    botao = &lt;button onClick={onLogout}&gt;Sair&lt;/button&gt;;
  } else {
    botao = &lt;button onClick={onLogin}&gt;Entrar&lt;/button&gt;;
  }
  
  return &lt;div&gt;{botao}&lt;/div&gt;;
}</code></pre>
      
      <h3>Exemplo Prático Completo</h3>
      <pre><code>function Dashboard({ usuario, notificacoes }) {
  return (
    &lt;div&gt;
      {/* Renderização condicional com múltiplas condições */}
      {usuario ? (
        &lt;&gt;
          &lt;h1&gt;Bem-vindo, {usuario.nome}!&lt;/h1&gt;
          {usuario.premium && (
            &lt;div className="badge"&gt;⭐ Membro Premium&lt;/div&gt;
          )}
          {notificacoes.length > 0 ? (
            &lt;div&gt;
              &lt;h2&gt;Você tem {notificacoes.length} notificações&lt;/h2&gt;
              &lt;ul&gt;
                {notificacoes.map(notif => (
                  &lt;li key={notif.id}&gt;{notif.mensagem}&lt;/li&gt;
                ))}
              &lt;/ul&gt;
            &lt;/div&gt;
          ) : (
            &lt;p&gt;Nenhuma notificação nova&lt;/p&gt;
          )}
        &lt;/&gt;
      ) : (
        &lt;div&gt;Por favor, faça login para continuar&lt;/div&gt;
      )}
    &lt;/div&gt;
  );
}</code></pre>
      
      <h2>Composição vs Herança (Composition vs Inheritance)</h2>
      <p>React tem um poderoso modelo de composição, e é recomendado usar composição ao invés de herança para reutilizar código entre componentes.</p>
      
      <h3>Por que Composição e não Herança?</h3>
      <p>Em programação orientada a objetos tradicional, você pode criar uma classe e fazer outras classes herdarem dela. No React, isso não é necessário e geralmente não é recomendado. Em vez disso, você compõe componentes menores para criar componentes maiores.</p>
      
      <h3>O que é Composição?</h3>
      <p>Composição significa construir algo complexo combinando partes menores. É como construir com blocos de LEGO - você pega peças pequenas e as combina para fazer algo maior.</p>
      
      <h3>Exemplo: Composição com Children</h3>
      <pre><code>// Componente base que aceita qualquer conteúdo
function Card({ children, titulo }) {
  return (
    &lt;div className="card"&gt;
      {titulo && &lt;h2&gt;{titulo}&lt;/h2&gt;}
      &lt;div className="card-content"&gt;
        {children}
      &lt;/div&gt;
    &lt;/div&gt;
  );
}

// Usando composição
function App() {
  return (
    &lt;Card titulo="Meu Card"&gt;
      &lt;p&gt;Este é o conteúdo do card!&lt;/p&gt;
      &lt;button&gt;Clique aqui&lt;/button&gt;
    &lt;/Card&gt;
  );
}</code></pre>
      
      <h3>Exemplo: Composição com Props Específicas</h3>
      <pre><code>// Componentes pequenos e reutilizáveis
function Botao({ texto, onClick, variante = "primary" }) {
  return (
    &lt;button 
      className={"btn btn-" + variante}
      onClick={onClick}
    &gt;
      {texto}
    &lt;/button&gt;
  );
}

function Input({ label, value, onChange }) {
  return (
    &lt;div&gt;
      &lt;label&gt;{label}&lt;/label&gt;
      &lt;input value={value} onChange={onChange} /&gt;
    &lt;/div&gt;
  );
}

// Componente maior composto de componentes menores
function FormularioLogin() {
  const [email, setEmail] = useState('');
  const [senha, setSenha] = useState('');
  
  const handleSubmit = () => {
    console.log('Login:', email, senha);
  };
  
  return (
    &lt;Card titulo="Login"&gt;
      &lt;Input 
        label="Email" 
        value={email} 
        onChange={(e) =&gt; setEmail(e.target.value)} 
      /&gt;
      &lt;Input 
        label="Senha" 
        value={senha} 
        onChange={(e) =&gt; setSenha(e.target.value)} 
      /&gt;
      &lt;Botao 
        texto="Entrar" 
        onClick={handleSubmit}
        variante="primary"
      /&gt;
    &lt;/Card&gt;
  );
}</code></pre>
      
      <h3>Exemplo: Composição com Múltiplos Slots</h3>
      <pre><code>// Componente que aceita múltiplas seções
function Layout({ header, sidebar, main, footer }) {
  return (
    &lt;div className="layout"&gt;
      {header && &lt;header&gt;{header}&lt;/header&gt;}
      &lt;div className="body"&gt;
        {sidebar && &lt;aside&gt;{sidebar}&lt;/aside&gt;}
        &lt;main&gt;{main}&lt;/main&gt;
      &lt;/div&gt;
      {footer && &lt;footer&gt;{footer}&lt;/footer&gt;}
    &lt;/div&gt;
  );
}

// Usando o layout composto
function App() {
  return (
    &lt;Layout
      header={&lt;h1&gt;Minha Aplicação&lt;/h1&gt;}
      sidebar={&lt;nav&gt;Menu&lt;/nav&gt;}
      main={&lt;div&gt;Conteúdo principal&lt;/div&gt;}
      footer={&lt;p&gt;Rodapé&lt;/p&gt;}
    /&gt;
  );
}</code></pre>
      
      <h3>Exemplo: Composição com HOCs (Higher-Order Components)</h3>
      <pre><code>// Função que retorna um componente melhorado
function comAutenticacao(Componente) {
  return function ComponenteAutenticado(props) {
    const [usuario, setUsuario] = useState(null);
    
    useEffect(() => {
      // Lógica de autenticação
      setUsuario({ nome: "João" });
    }, []);
    
    if (!usuario) {
      return &lt;div&gt;Carregando...&lt;/div&gt;;
    }
    
    return &lt;Componente {...props} usuario={usuario} /&gt;;
  };
}

// Componente simples
function Perfil({ usuario }) {
  return &lt;h1&gt;Perfil de {usuario.nome}&lt;/h1&gt;;
}

// Compor com autenticação
const PerfilAutenticado = comAutenticacao(Perfil);</code></pre>
      
      <h3>Por que não Herança?</h3>
      <p>React não precisa de herança porque:</p>
      <ul>
        <li><strong>Flexibilidade:</strong> Composição é mais flexível que herança</li>
        <li><strong>Reutilização:</strong> É mais fácil reutilizar código com composição</li>
        <li><strong>Manutenção:</strong> Componentes compostos são mais fáceis de entender e manter</li>
        <li><strong>Testabilidade:</strong> Componentes menores são mais fáceis de testar</li>
      </ul>
      
      <h3>Comparação: Herança vs Composição</h3>
      <pre><code>// ❌ NÃO FAÇA ISSO (Herança - não recomendado)
class BotaoBase extends React.Component {
  render() {
    return &lt;button className="base"&gt;{this.props.children}&lt;/button&gt;;
  }
}

class BotaoPrimario extends BotaoBase {
  render() {
    return &lt;button className="primary"&gt;{super.render()}&lt;/button&gt;;
  }
}

// ✅ FAÇA ISSO (Composição - recomendado)
function Botao({ children, variante = "base" }) {
  return &lt;button className={variante}&gt;{children}&lt;/button&gt;;
}

function BotaoPrimario({ children }) {
  return &lt;Botao variante="primary"&gt;{children}&lt;/Botao&gt;;
}</code></pre>
      
      <h3>Padrões de Composição Comuns</h3>
      <ul>
        <li><strong>Containment:</strong> Usar <code>children</code> para passar conteúdo</li>
        <li><strong>Specialization:</strong> Componentes específicos que usam componentes genéricos</li>
        <li><strong>HOCs:</strong> Funções que retornam componentes melhorados</li>
        <li><strong>Render Props:</strong> Props que são funções que retornam JSX</li>
      </ul>
      
      <h2>Resumo: Fundamentos de Componentes</h2>
      <p>Nesta aula, você aprendeu que:</p>
      <ul>
        <li>✅ <strong>Componentes</strong> são blocos de construção reutilizáveis</li>
        <li>✅ <strong>Componentes Funcionais</strong> são funções JavaScript que retornam JSX</li>
        <li>✅ <strong>JSX/TSX</strong> é uma sintaxe que facilita escrever interfaces</li>
        <li>✅ <strong>Props</strong> vêm de fora (somente leitura)</li>
        <li>✅ <strong>State</strong> é gerenciado internamente (mutável)</li>
        <li>✅ <strong>Renderização Condicional</strong> permite mostrar conteúdo baseado em condições</li>
        <li>✅ <strong>Composição</strong> é preferível a herança para reutilizar código</li>
      </ul>
      
      <p>Com esses conceitos fundamentais, você está pronto para construir aplicações React incríveis! 🚀</p>
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

