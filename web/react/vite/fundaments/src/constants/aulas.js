// Dados estáticos das aulas
import {
  Exemplo1Contador,
  Exemplo2ListaTarefas,
  Exemplo3ComponentesProps,
  Exemplo4AppCompleto
} from '../components/aulas/Aula1'

import {
  Exemplo1ComponenteBasico,
  Exemplo2Props,
  Exemplo3State,
  Exemplo4PropsVsState,
  Exemplo5ConditionalRendering,
  Exemplo6Composition,
  Exemplo7Children,
  Exemplo8JSX
} from '../components/aulas/Aula2'

import {
  Exemplo1Lifecycle,
  Exemplo2ListsKeys,
  Exemplo3RenderProps,
  Exemplo4Refs,
  Exemplo5Events
} from '../components/aulas/Aula3'

export const AULAS = [
  {
    id: '1',
    titulo: 'Aula 1: CLI Tools e Vite - Introdução ao React',
    descricao: 'Domine a base do React: Virtual DOM, Vite e a estrutura inicial do projeto.',
    teoria: `
        <div style="line-height: 1.8; color: #333;">
            <h2 style="color: #2c3e50; border-bottom: 2px solid #3498db; padding-bottom: 0.5rem; margin-top: 2rem; margin-bottom: 1rem;">1. A Revolução do React</h2>
            
            <p><strong>React</strong> não é apenas uma biblioteca; é uma mudança de paradigma. Criado pelo Facebook em 2013, ele resolveu um problema gigante: <strong>como atualizar interfaces complexas de forma eficiente e sem dor de cabeça?</strong></p>
            
            <div style="background: #e3f2fd; padding: 1.5rem; border-radius: 8px; border-left: 5px solid #1565c0; margin: 1.5rem 0;">
                <h3 style="margin-top: 0; color: #1565c0;">Imperativo vs Declarativo: A Grande Sacada</h3>
                <p>Para entender o React, você precisa entender essa diferença. Imagine que você quer um táxi.</p>
                
                <div style="display: grid; grid-template-columns: 1fr 1fr; gap: 1rem; margin-top: 1rem;">
                    <div style="background: #fff; padding: 1rem; border-radius: 4px; border: 1px solid #bbdefb;">
                        <strong style="color: #d32f2f;">Abordagem Imperativa (jQuery/JS Puro)</strong>
                        <p style="font-size: 0.9rem; margin-bottom: 0.5rem;">Você diz COMO fazer passo-a-passo:</p>
                        <ol style="font-size: 0.85rem; padding-left: 1.2rem;">
                            <li>Saia de casa.</li>
                            <li>Vá até a esquina.</li>
                            <li>Levante a mão.</li>
                            <li>Espere um táxi parar.</li>
                            <li>Entre no táxi.</li>
                        </ol>
                        <code style="display: block; background: #f5f5f5; padding: 0.5rem; margin-top: 0.5rem; font-size: 0.8rem;">
                            const btn = document.createElement('button');<br>
                            btn.innerText = 'Táxi';<br>
                            btn.className = 'blue';<br>
                            parent.appendChild(btn);
                        </code>
                    </div>
                    
                    <div style="background: #fff; padding: 1rem; border-radius: 4px; border: 1px solid #c8e6c9;">
                        <strong style="color: #2e7d32;">Abordagem Declarativa (React)</strong>
                        <p style="font-size: 0.9rem; margin-bottom: 0.5rem;">Você diz O QUE você quer:</p>
                        <ul style="font-size: 0.85rem; padding-left: 1.2rem;">
                            <li>"Quero um táxi aqui."</li>
                            <li>(O aplicativo resolve como o táxi chega até você)</li>
                        </ul>
                        <code style="display: block; background: #f5f5f5; padding: 0.5rem; margin-top: 0.5rem; font-size: 0.8rem;">
                            return &lt;Button color="blue"&gt;Táxi&lt;/Button&gt;
                        </code>
                    </div>
                </div>
            </div>

            <h2 style="color: #2c3e50; border-bottom: 2px solid #3498db; padding-bottom: 0.5rem; margin-top: 2rem; margin-bottom: 1rem;">2. O Motor do React: Virtual DOM</h2>

            <p>O <strong>DOM (Document Object Model)</strong> é a árvore de elementos do seu navegador. Ele é lento para atualizar. Se você mudar uma cor de fundo 1000 vezes, o navegador recalcula o layout 1000 vezes.</p>

            <div style="background: #fff3e0; padding: 1rem; border-radius: 8px; border-left: 4px solid #ff9800; margin: 1rem 0;">
                <h3 style="margin-top: 0; color: #e65100;">A Analogia do Garçom 🍽️</h3>
                <p>Imagine o DOM como a cozinha de um restaurante e você é o cliente (React).</p>
                <ul>
                    <li><strong>Sem Virtual DOM:</strong> Você grita cada pedido para a cozinha. "Quero água!", a cozinha para e pega água. "Quero pão!", a cozinha para e pega pão. É ineficiente.</li>
                    <li><strong>Com Virtual DOM:</strong> Você fala com o garçom (Virtual DOM). Você faz vários pedidos ("água, pão, sopa"). O garçom anota tudo, otimiza a ordem e entrega um pedido único para a cozinha.</li>
                </ul>
            </div>

            <h3 style="color: #2c3e50;">Como funciona tecnicamente (Reconciliação):</h3>
            <ol>
                <li><strong>Render:</strong> Quando o estado muda, o React cria uma nova árvore Virtual DOM.</li>
                <li><strong>Diffing:</strong> Ele compara essa nova árvore com a anterior. Ele vê: "Ah, só o texto do botão mudou de 'Entrar' para 'Sair'".</li>
                <li><strong>Commit:</strong> Ele vai no DOM real e muda <em>apenas</em> aquele texto. O resto da página nem pisca.</li>
            </ol>

            <h2 style="color: #2c3e50; border-bottom: 2px solid #3498db; padding-bottom: 0.5rem; margin-top: 2rem; margin-bottom: 1rem;">3. Vite: Velocidade Supersônica</h2>

            <p>Por que usamos Vite e não o antigo <code>create-react-app</code>? A diferença é arquitetural.</p>

            <div style="display: flex; gap: 1rem; flex-wrap: wrap;">
                <div style="flex: 1; min-width: 250px; background: #f5f5f5; padding: 1rem; border-radius: 8px;">
                    <h4 style="margin-top: 0;">🐢 Webpack (Antigo)</h4>
                    <p>Precisava ler <strong>TODOS</strong> os seus arquivos, empacotar tudo em um arquivo gigante (bundle.js) antes de iniciar o servidor. Em projetos grandes, isso levava minutos.</p>
                </div>
                <div style="flex: 1; min-width: 250px; background: #e8f5e9; padding: 1rem; border-radius: 8px; border: 1px solid #4caf50;">
                    <h4 style="margin-top: 0;">⚡ Vite (Novo)</h4>
                    <p>Usa <strong>ES Modules</strong> nativos do navegador. Ele não empacota nada para iniciar. Ele serve os arquivos conforme o navegador pede. O início é instantâneo (ms).</p>
                </div>
            </div>

            <h3 style="color: #2c3e50; margin-top: 1rem;">HMR (Hot Module Replacement)</h3>
            <p>O Vite tem um HMR incrível. Se você editar um componente <code>Botao.jsx</code>, ele troca apenas esse arquivo no navegador rodando, mantendo o estado da aplicação (ex: o texto que você digitou num formulário não some).</p>

            <h2 style="color: #2c3e50; border-bottom: 2px solid #3498db; padding-bottom: 0.5rem; margin-top: 2rem; margin-bottom: 1rem;">4. Anatomia de um Componente React</h2>

            <p>Vamos dissecar um componente simples para entender cada parte:</p>

            <pre style="background: #282c34; color: #abb2bf; padding: 1.5rem; border-radius: 8px; overflow-x: auto; font-family: 'Fira Code', monospace;">
<span style="color: #c678dd;">import</span> { useState } <span style="color: #c678dd;">from</span> <span style="color: #98c379;">'react'</span> <span style="color: #5c6370;">// 1. Importação de Hooks</span>

<span style="color: #c678dd;">function</span> <span style="color: #e5c07b;">MeuBotao</span>() {              <span style="color: #5c6370;">// 2. Declaração (PascalCase)</span>
  <span style="color: #5c6370;">// 3. Lógica e Estado (JavaScript puro)</span>
  <span style="color: #c678dd;">const</span> [ativo, setAtivo] = <span style="color: #61afef;">useState</span>(<span style="color: #d19a66;">false</span>)

  <span style="color: #c678dd;">return</span> (
    <span style="color: #5c6370;">// 4. O Retorno (JSX - O que aparece na tela)</span>
    &lt;<span style="color: #e06c75;">button</span> 
      <span style="color: #d19a66;">onClick</span>={() => <span style="color: #61afef;">setAtivo</span>(!ativo)}
      <span style="color: #d19a66;">className</span>={ativo ? <span style="color: #98c379;">"ligado"</span> : <span style="color: #98c379;">"desligado"</span>}
    &gt;
      {ativo ? <span style="color: #98c379;">"LIGADO"</span> : <span style="color: #98c379;">"DESLIGADO"</span>}
    &lt;/<span style="color: #e06c75;">button</span>&gt;
  )
}
            </pre>

            <ul>
                <li><strong>PascalCase:</strong> Componentes devem sempre começar com letra maiúscula (ex: <code>MeuComponente</code>), senão o React acha que é uma tag HTML normal.</li>
                <li><strong>Hooks:</strong> Funções que começam com <code>use</code> (como <code>useState</code>). Elas "gancham" funcionalidades do React na sua função.</li>
                <li><strong>JSX:</strong> Parece HTML, mas aceita lógica JavaScript dentro de <code>{}</code>.</li>
            </ul>

            <h2 style="color: #2c3e50; border-bottom: 2px solid #3498db; padding-bottom: 0.5rem; margin-top: 2rem; margin-bottom: 1rem;">5. Dicas de Ouro para Iniciantes 💎</h2>
            
            <ul style="list-style: none; padding: 0;">
                <li style="margin-bottom: 1rem; display: flex; align-items: start;">
                    <span style="font-size: 1.5rem; margin-right: 0.5rem;">🚫</span>
                    <div>
                        <strong>Não manipule o DOM diretamente:</strong> Esqueça <code>document.getElementById</code> ou <code>classList.add</code>. Se algo tem que mudar na tela, mude o <strong>Estado</strong> e deixe o React reagir.
                    </div>
                </li>
                <li style="margin-bottom: 1rem; display: flex; align-items: start;">
                    <span style="font-size: 1.5rem; margin-right: 0.5rem;">🔄</span>
                    <div>
                        <strong>Fluxo Unidirecional:</strong> Dados descem (Pai para Filho via Props). Eventos sobem (Filho chama função do Pai). Nunca tente passar dados "de lado" sem usar gerenciamento de estado global.
                    </div>
                </li>
                <li style="margin-bottom: 1rem; display: flex; align-items: start;">
                    <span style="font-size: 1.5rem; margin-right: 0.5rem;">🧩</span>
                    <div>
                        <strong>Pense em Componentes:</strong> Olhe para uma interface (ex: Instagram). O que você vê? Navbar, Story, FeedItem, LikeButton. Quebre tudo em pedaços pequenos.
                    </div>
                </li>
            </ul>
        </div>
        `,
    exemplos: [
      {
        title: '1. Contador Simples',
        description: 'Aprenda o básico de State com useState. Veja como o número muda quando você clica.',
        code: `// O useState cria uma variável que o React "observa"
const [count, setCount] = useState(0)

// Para mudar o valor, usamos a função setCount
<button onClick={() => setCount(count + 1)}>
  Incrementar
</button>`,
        ExampleComponent: Exemplo1Contador
      },
      {
        title: '2. Lista de Tarefas',
        description: 'Gerenciando listas e inputs. Aprenda a adicionar e remover itens de um array no State.',
        code: `// Adicionando item (Imutabilidade)
setTarefas([...tarefas, novaTarefa])

// Removendo item (Filter)
setTarefas(tarefas.filter((_, i) => i !== index))`,
        ExampleComponent: Exemplo2ListaTarefas
      },
      {
        title: '3. Componentes e Props',
        description: 'Como criar componentes reutilizáveis e passar dados para eles via Props.',
        code: `function Card({ titulo, cor }) {
  return <div style={{ borderColor: cor }}>{titulo}</div>
}

// Reutilizando:
<Card titulo="React" cor="blue" />
<Card titulo="Vite" cor="purple" />`,
        ExampleComponent: Exemplo3ComponentesProps
      },
      {
        title: '4. Mini Loja (App Completo)',
        description: 'Juntando tudo: State, Props, Renderização Condicional e Listas em um app de verdade.',
        code: `// Renderização Condicional
{inCart ? (
  <button>Remover</button>
) : (
  <button>Comprar</button>
)}`,
        ExampleComponent: Exemplo4AppCompleto
      }
    ]
  },
  {
    id: '2',
    titulo: 'Aula 2: Components - Os Blocos de Construção do React',
    descricao: 'Aprofunde-se em Components, JSX, Props, State e Composição.',
    teoria: `
      <div style="line-height: 1.8; color: #333;">
        <h2 style="color: #2c3e50; border-bottom: 2px solid #3498db; padding-bottom: 0.5rem; margin-top: 2rem; margin-bottom: 1rem;">O Que São Components?</h2>
        
        <p><strong>Components</strong> são como peças de LEGO. Em vez de construir o site inteiro de uma vez, você constrói pequenas peças (Botão, Cabeçalho, Card) e as junta.</p>
        
        <div style="background: #e8f5e9; padding: 1rem; border-radius: 8px; border-left: 4px solid #4caf50; margin: 1rem 0;">
          <h3 style="margin-top: 0; color: #2e7d32;">Pilares dos Componentes</h3>
          <ul style="margin-bottom: 0;">
            <li><strong>Reutilização:</strong> Crie uma vez, use quantas vezes quiser.</li>
            <li><strong>Isolamento:</strong> Se um botão quebra, o resto do site continua funcionando.</li>
            <li><strong>Composição:</strong> Componentes podem conter outros componentes.</li>
          </ul>
        </div>

        <h2 style="color: #2c3e50; border-bottom: 2px solid #3498db; padding-bottom: 0.5rem; margin-top: 2rem; margin-bottom: 1rem;">Functional Components e JSX</h2>
        
        <p>Hoje em dia, usamos <strong>Functional Components</strong>. São apenas funções JavaScript que retornam <strong>JSX</strong>.</p>
        <p>JSX parece HTML, mas é JavaScript disfarçado. Ele permite misturar lógica (JS) com marcação (HTML).</p>
        
        <div style="background: #fff3e0; padding: 1rem; border-radius: 8px; border-left: 4px solid #ff9800; margin: 1rem 0;">
          <h3 style="margin-top: 0; color: #e65100;">Regras de Ouro do JSX:</h3>
          <ul>
            <li>Retorne sempre <strong>um único elemento pai</strong> (ou use <code>&lt;&gt;...&lt;/&gt;</code>).</li>
            <li>Use <code>className</code> em vez de <code>class</code>.</li>
            <li>Abra chaves <code>{}</code> para escrever JavaScript dentro do HTML.</li>
            <li>Feche todas as tags (mesmo <code>&lt;br /&gt;</code>).</li>
          </ul>
        </div>

        <h2 style="color: #2c3e50; border-bottom: 2px solid #3498db; padding-bottom: 0.5rem; margin-top: 2rem; margin-bottom: 1rem;">A Batalha: Props vs State</h2>
        
        <div style="display: grid; grid-template-columns: 1fr 1fr; gap: 1rem; margin: 1rem 0;">
          <div style="background: #e3f2fd; padding: 1rem; border-radius: 8px; border: 2px solid #2196f3;">
            <h3 style="margin-top: 0; color: #1976d2;">Props (Externo)</h3>
            <ul>
                <li>Vêm do Pai para o Filho.</li>
                <li>São <strong>Imutáveis</strong> (Read-Only).</li>
                <li>São como argumentos de função.</li>
                <li>Ex: Cor de um botão, Título de um card.</li>
            </ul>
          </div>
          
          <div style="background: #f3e5f5; padding: 1rem; border-radius: 8px; border: 2px solid #9c27b0;">
            <h3 style="margin-top: 0; color: #7b1fa2;">State (Interno)</h3>
            <ul>
                <li>Nasce e morre dentro do componente.</li>
                <li>É <strong>Mutável</strong> (via setState).</li>
                <li>É como a memória local.</li>
                <li>Ex: O texto digitado em um input, se um modal está aberto.</li>
            </ul>
          </div>
        </div>

        <h2 style="color: #2c3e50; border-bottom: 2px solid #3498db; padding-bottom: 0.5rem; margin-top: 2rem; margin-bottom: 1rem;">Renderização Condicional</h2>
        <p>Não existe <code>v-if</code> ou <code>ng-if</code> no React. Usamos JavaScript puro!</p>
        <ul>
            <li><strong>Ternário (<code>? :</code>):</strong> Para "Se isso, então aquilo, senão aquilo outro".</li>
            <li><strong>AND (<code>&&</code>):</strong> Para "Se isso for verdade, mostre isso".</li>
        </ul>

        <h2 style="color: #2c3e50; border-bottom: 2px solid #3498db; padding-bottom: 0.5rem; margin-top: 2rem; margin-bottom: 1rem;">Composição vs Herança</h2>
        <p>No React, preferimos <strong>Composição</strong>. Em vez de estender classes, passamos componentes como props (geralmente via <code>children</code>).</p>
        <p>A prop <code>children</code> é mágica: ela pega tudo o que você coloca dentro das tags de abertura e fechamento do seu componente.</p>
      </div>
    `,
    exemplos: [
      {
        title: '1. Componente Básico',
        description: 'O exemplo mais simples de um componente React. Uma função que retorna JSX.',
        code: `function Exemplo1ComponenteBasico() {
  return (
    <div>
      <h3>Olá, React!</h3>
      <p>Este é meu primeiro componente funcional!</p>
    </div>
  )
}`,
        ExampleComponent: Exemplo1ComponenteBasico
      },
      {
        title: '2. Componente com Props',
        description: 'Props são dados passados de um componente pai para um filho. Veja como o mesmo componente pode ser usado com dados diferentes.',
        code: `function Saudacao({ nome, idade }) {
  return (
    <div>
      <h3>Olá, {nome}!</h3>
      <p>Você tem {idade} anos.</p>
      <p>Você nasceu em {new Date().getFullYear() - idade}.</p>
    </div>
  )
}

// Uso:
<Saudacao nome="Maria" idade={25} />
<Saudacao nome="João" idade={30} />`,
        ExampleComponent: Exemplo2Props
      },
      {
        title: '3. Componente com State',
        description: 'State é a memória interna do componente. Quando o state muda, o componente re-renderiza automaticamente. Use useState para criar estado.',
        code: `import { useState } from 'react'

function Exemplo3State() {
  const [count, setCount] = useState(0)

  return (
    <div>
      <h3>Contador: {count}</h3>
      <button onClick={() => setCount(count - 1)}>
        - Decrementar
      </button>
      <button onClick={() => setCount(0)}>Resetar</button>
      <button onClick={() => setCount(count + 1)}>
        + Incrementar
      </button>
    </div>
  )
}`,
        ExampleComponent: Exemplo3State
      },
      {
        title: '4. Props vs State - Diferenças',
        description: 'Veja a diferença prática entre Props (read-only, vêm de fora) e State (mutável, interno ao componente).',
        code: `// Props: Read-only, vêm do componente pai
function DisplayProps({ mensagem }) {
  return <p>Mensagem: {mensagem}</p>
  // Não pode modificar mensagem aqui!
}

// State: Mutável, controlado pelo próprio componente
function DisplayState() {
  const [mensagem, setMensagem] = useState('Estado inicial')
  
  return (
    <div>
      <p>Mensagem: {mensagem}</p>
      <button onClick={() => setMensagem('Nova mensagem!')}>
        Mudar Mensagem
      </button>
    </div>
  )
}`,
        ExampleComponent: Exemplo4PropsVsState
      },
      {
        title: '5. Conditional Rendering',
        description: 'Diferentes formas de renderizar condicionalmente: operador ternário, operador &&, e múltiplas condições.',
        code: `// Operador Ternário
{estaLogado ? (
  <p>Bem-vindo de volta!</p>
) : (
  <p>Por favor, faça login</p>
)}

// Operador &&
{mostrarNotificacao && (
  <div>Notificação importante!</div>
)}

// Múltiplas Condições
{status === 'loading' && <p>Carregando...</p>}
{status === 'success' && <p>Sucesso!</p>}
{status === 'error' && <p>Erro!</p>}`,
        ExampleComponent: Exemplo5ConditionalRendering
      },
      {
        title: '6. Composition (Composição)',
        description: 'Construa componentes maiores combinando componentes menores. Veja como Botao, Card e Container são combinados para criar uma interface complexa.',
        code: `// Componentes pequenos
function Botao({ children, onClick, variant }) {
  return <button onClick={onClick}>{children}</button>
}

function Card({ title, children, footer }) {
  return (
    <div>
      {title && <h3>{title}</h3>}
      <div>{children}</div>
      {footer && <div>{footer}</div>}
    </div>
  )
}

// Composição: combinando componentes
function ProductCard({ product }) {
  return (
    <Card
      title={product.name}
      footer={
        <Botao onClick={() => addToCart(product.id)}>
          Adicionar ao Carrinho
        </Botao>
      }
    >
      <p>Preço: R$ {product.price}</p>
    </Card>
  )
}`,
        ExampleComponent: Exemplo6Composition
      },
      {
        title: '7. Props.children',
        description: 'children é uma prop especial que contém tudo que você coloca entre as tags de abertura e fechamento do componente.',
        code: `function Caixa({ children, titulo }) {
  return (
    <div>
      {titulo && <h4>{titulo}</h4>}
      <div>{children}</div>
    </div>
  )
}

// Uso: o que está entre as tags vira children
<Caixa titulo="Minha Caixa">
  <p>Este parágrafo é children</p>
  <button>Este botão também é children</button>
</Caixa>`,
        ExampleComponent: Exemplo7Children
      },
      {
        title: '8. JSX - Expressões e Regras',
        description: 'Veja como usar expressões JavaScript dentro do JSX: variáveis, cálculos, operadores ternários, e renderização de arrays.',
        code: `function Exemplo8JSX() {
  const nome = 'React'
  const versao = 18
  const tecnologias = ['JSX', 'Hooks', 'Components']
  const estaAtivo = true

  return (
    <div>
      {/* Variáveis */}
      <p>Olá, {nome}! Versão: {versao}</p>
      
      {/* Expressões matemáticas */}
      <p>10 + 5 = {10 + 5}</p>
      
      {/* Operador ternário */}
      <p>Status: {estaAtivo ? 'Ativo' : 'Inativo'}</p>
      
      {/* Renderizando arrays */}
      <ul>
        {tecnologias.map((tech, index) => (
          <li key={index}>{tech}</li>
        ))}
      </ul>
    </div>
  )
}`,
        ExampleComponent: Exemplo8JSX
      }
    ]
  },
  {
    id: '3',
    titulo: 'Aula 3: Rendering e Ciclo de Vida',
    descricao: 'Entenda como o React renderiza, o Virtual DOM, Ciclo de Vida, Listas e Hooks avançados.',
    teoria: `
      <div style="line-height: 1.8; color: #333;">
        <h2 style="color: #2c3e50; border-bottom: 2px solid #3498db; padding-bottom: 0.5rem; margin-top: 2rem; margin-bottom: 1rem;">O Motor do React: Rendering</h2>
        
        <p>Como o React sabe o que mudar na tela? Ele usa uma <strong>Abordagem Declarativa</strong>. Você diz "Quero um botão azul" e ele se vira para atualizar o DOM.</p>

        <div style="background: #e3f2fd; padding: 1rem; border-radius: 8px; border-left: 4px solid #2196f3; margin: 1rem 0;">
          <h3 style="margin-top: 0; color: #1565c0;">O Processo de Reconciliação (Reconciliation)</h3>
          <ol>
            <li><strong>Render Phase:</strong> O React chama seus componentes e cria um novo <strong>Virtual DOM</strong>.</li>
            <li><strong>Diffing:</strong> Ele compara esse novo Virtual DOM com o anterior.</li>
            <li><strong>Commit Phase:</strong> Ele aplica apenas as mudanças necessárias no <strong>DOM Real</strong>.</li>
          </ol>
        </div>

        <h2 style="color: #2c3e50; border-bottom: 2px solid #3498db; padding-bottom: 0.5rem; margin-top: 2rem; margin-bottom: 1rem;">Ciclo de Vida (Life Cycle)</h2>
        <p>Todo componente nasce (Mount), vive (Update) e morre (Unmount). Com Hooks, usamos o <code>useEffect</code> para controlar isso.</p>

        <ul style="background: #f5f5f5; padding: 1.5rem; border-radius: 8px;">
          <li><strong>Mounting (Nascer):</strong> <code>useEffect(() => {}, [])</code> - Executa uma vez ao aparecer.</li>
          <li><strong>Updating (Viver):</strong> <code>useEffect(() => {}, [dep])</code> - Executa quando 'dep' muda.</li>
          <li><strong>Unmounting (Morrer):</strong> A função de retorno do <code>useEffect</code> (Cleanup).</li>
        </ul>

        <h2 style="color: #2c3e50; border-bottom: 2px solid #3498db; padding-bottom: 0.5rem; margin-top: 2rem; margin-bottom: 1rem;">Listas e Keys 🔑</h2>
        <p>Ao renderizar listas (arrays), o React precisa de uma <strong>Key</strong> única para cada item. Isso ajuda ele a saber qual item mudou, foi adicionado ou removido.</p>
        <p style="color: red;"><strong>Nunca use o índice do array como key se a lista puder mudar de ordem!</strong></p>

        <h2 style="color: #2c3e50; border-bottom: 2px solid #3498db; padding-bottom: 0.5rem; margin-top: 2rem; margin-bottom: 1rem;">Render Props</h2>
        <p>É uma técnica onde você passa uma função como prop para um componente, e ele usa essa função para saber o que renderizar. É ótimo para compartilhar lógica (ex: rastrear mouse, scroll).</p>

        <h2 style="color: #2c3e50; border-bottom: 2px solid #3498db; padding-bottom: 0.5rem; margin-top: 2rem; margin-bottom: 1rem;">Refs (Referências)</h2>
        <p>As vezes você precisa "sair" do React e tocar no DOM diretamente (focar um input, medir um elemento). Para isso usamos <code>useRef</code>.</p>
        <p>O <code>useRef</code> também serve para guardar valores que persistem entre renders mas NÃO causam re-render.</p>
      </div>
    `,
    exemplos: [
      {
        title: '1. Ciclo de Vida (useEffect)',
        description: 'Entenda Mounting, Updating e Unmounting com um Timer.',
        code: `useEffect(() => {
  // Mount
  const timer = setInterval(() => console.log('Tick'), 1000);

  // Unmount (Cleanup)
  return () => clearInterval(timer);
}, []); // [] = Apenas no mount`,
        ExampleComponent: Exemplo1Lifecycle
      },
      {
        title: '2. Listas e Keys',
        description: 'A importância de usar keys únicas para performance e evitar bugs.',
        code: `// Errado (Index)
{items.map((item, index) => <li key={index}>{item}</li>)}

// Correto (ID único)
{items.map(item => <li key={item.id}>{item.text}</li>)}`,
        ExampleComponent: Exemplo2ListsKeys
      },
      {
        title: '3. Render Props',
        description: 'Compartilhando lógica de posição do mouse via props.',
        code: `<MouseTracker render={({ x, y }) => (
  <h1>O mouse está em {x}, {y}</h1>
)} />`,
        ExampleComponent: Exemplo3RenderProps
      },
      {
        title: '4. Refs e DOM',
        description: 'Acessando elementos do DOM imperativamente com useRef.',
        code: `const inputRef = useRef(null);

function focar() {
  inputRef.current.focus();
}

<input ref={inputRef} />`,
        ExampleComponent: Exemplo4Refs
      },
      {
        title: '5. Eventos (SyntheticEvent)',
        description: 'Como o React lida com eventos de forma consistente entre navegadores.',
        code: `function handleClick(e) {
  e.preventDefault(); // Funciona igual em todo lugar
  console.log(e.target);
}

<button onClick={handleClick}>Clique</button>`,
        ExampleComponent: Exemplo5Events
      }
    ]
  }
]
