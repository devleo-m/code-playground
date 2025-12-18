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

export const AULAS = [
  {
    id: '1',
    titulo: 'Aula 1: CLI Tools e Vite - Introdução ao React',
    descricao: 'Aprenda os conceitos básicos do React: State, Props e como o Vite funciona.',
    teoria: `
        <div style="line-height: 1.8; color: #333;">
            <h2 style="color: #2c3e50; border-bottom: 2px solid #3498db; padding-bottom: 0.5rem; margin-top: 2rem; margin-bottom: 1rem;">O Que é React?</h2>
            <p><strong>React</strong> é uma biblioteca JavaScript para construir interfaces de usuário (UI). Ele é baseado em <strong>Componentes</strong>, que são como peças de LEGO que você junta para criar seu site.</p>
            
            <h2 style="color: #2c3e50; border-bottom: 2px solid #3498db; padding-bottom: 0.5rem; margin-top: 2rem; margin-bottom: 1rem;">Conceitos Fundamentais</h2>
            
            <div style="display: grid; grid-template-columns: 1fr 1fr; gap: 1rem; margin: 1rem 0;">
                <div style="background: #e3f2fd; padding: 1rem; border-radius: 8px; border-left: 4px solid #2196f3;">
                    <h3 style="margin-top: 0; color: #1565c0;">Props (Propriedades)</h3>
                    <p>São dados que passamos para os componentes. Como argumentos de uma função. São <strong>leitura (read-only)</strong>.</p>
                </div>
                <div style="background: #e8f5e9; padding: 1rem; border-radius: 8px; border-left: 4px solid #4caf50;">
                    <h3 style="margin-top: 0; color: #2e7d32;">State (Estado)</h3>
                    <p>É a memória do componente. Dados que mudam com o tempo (como um contador). Quando muda, o React atualiza a tela.</p>
                </div>
            </div>

            <h2 style="color: #2c3e50; border-bottom: 2px solid #3498db; padding-bottom: 0.5rem; margin-top: 2rem; margin-bottom: 1rem;">Por que Vite?</h2>
            <p><strong>Vite</strong> é a ferramenta que usamos para criar e rodar o projeto. Ele é extremamente rápido e moderno.</p>
            <ul style="background: #f5f5f5; padding: 1rem; border-radius: 8px;">
                <li>🚀 Inicia o servidor instantaneamente</li>
                <li>⚡ Atualiza a tela muito rápido (HMR)</li>
                <li>📦 Já vem configurado para React</li>
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
    descricao: 'Aprenda a criar e usar componentes React, entender props, state, JSX e composição',
    teoria: `
      <div style="line-height: 1.8; color: #333;">
        <h2 style="color: #2c3e50; border-bottom: 2px solid #3498db; padding-bottom: 0.5rem; margin-top: 2rem; margin-bottom: 1rem;">O Que São Components?</h2>
        
        <p><strong>Components</strong> (Componentes) são os blocos fundamentais de construção de aplicações React. Eles são unidades independentes e reutilizáveis de código que encapsulam tanto a lógica quanto a apresentação de uma parte da interface do usuário.</p>
        
        <p>Pense em components como peças de Lego: cada peça tem uma função específica, e você combina várias peças para construir algo maior e mais complexo. Da mesma forma, em React, você constrói interfaces complexas combinando componentes menores e mais simples.</p>
        
        <div style="background: #e8f5e9; padding: 1rem; border-radius: 8px; border-left: 4px solid #4caf50; margin: 1rem 0;">
          <h3 style="margin-top: 0; color: #2e7d32;">Por Que Components São Importantes?</h3>
          <ul style="margin-bottom: 0;">
            <li><strong>Reutilização:</strong> Escreva uma vez, use em qualquer lugar</li>
            <li><strong>Manutenibilidade:</strong> Código organizado e fácil de manter</li>
            <li><strong>Testabilidade:</strong> Componentes isolados são mais fáceis de testar</li>
            <li><strong>Colaboração:</strong> Diferentes desenvolvedores podem trabalhar em componentes diferentes</li>
            <li><strong>Abstração:</strong> Escondem complexidade, expondo apenas o necessário</li>
          </ul>
        </div>

        <h2 style="color: #2c3e50; border-bottom: 2px solid #3498db; padding-bottom: 0.5rem; margin-top: 2rem; margin-bottom: 1rem;">Functional Components</h2>
        
        <p><strong>Functional Components</strong> são componentes React definidos como funções JavaScript. Eles são a forma moderna e recomendada de criar componentes em React.</p>
        
        <p>A forma mais simples de um functional component é uma função que retorna JSX:</p>
        
        <div style="background: #f5f5f5; padding: 1rem; border-radius: 8px; margin: 1rem 0;">
          <pre style="margin: 0; overflow-x: auto;"><code>function Welcome() {
  return &lt;h1&gt;Bem-vindo ao React!&lt;/h1&gt;;
}</code></pre>
        </div>

        <h2 style="color: #2c3e50; border-bottom: 2px solid #3498db; padding-bottom: 0.5rem; margin-top: 2rem; margin-bottom: 1rem;">JSX (JavaScript XML)</h2>
        
        <p><strong>JSX</strong> é uma extensão de sintaxe do JavaScript que permite escrever código que parece HTML dentro do JavaScript. JSX não é HTML - é uma forma de descrever a estrutura da UI de forma declarativa.</p>
        
        <div style="background: #fff3e0; padding: 1rem; border-radius: 8px; border-left: 4px solid #ff9800; margin: 1rem 0;">
          <h3 style="margin-top: 0; color: #e65100;">Regras Importantes do JSX:</h3>
          <ul style="margin-bottom: 0;">
            <li><strong>Um único elemento raiz:</strong> JSX deve retornar um único elemento (ou usar Fragment <code>&lt;&gt;&lt;/&gt;</code>)</li>
            <li><strong>Atributos em camelCase:</strong> <code>class</code> vira <code>className</code>, <code>onclick</code> vira <code>onClick</code></li>
            <li><strong>Expressões JavaScript:</strong> Use <code>{}</code> para inserir variáveis e expressões</li>
            <li><strong>Segurança:</strong> JSX previne XSS automaticamente escapando valores</li>
          </ul>
        </div>

        <h2 style="color: #2c3e50; border-bottom: 2px solid #3498db; padding-bottom: 0.5rem; margin-top: 2rem; margin-bottom: 1rem;">Props vs State</h2>
        
        <div style="display: grid; grid-template-columns: 1fr 1fr; gap: 1rem; margin: 1rem 0;">
          <div style="background: #e3f2fd; padding: 1rem; border-radius: 8px; border: 2px solid #2196f3;">
            <h3 style="margin-top: 0; color: #1976d2;">Props (Propriedades)</h3>
            <ul style="margin-bottom: 0;">
              <li>Dados passados do componente pai</li>
              <li><strong>Read-only</strong> (somente leitura)</li>
              <li>Unidirecionais (pai → filho)</li>
              <li>Não podem ser modificadas</li>
            </ul>
          </div>
          
          <div style="background: #f3e5f5; padding: 1rem; border-radius: 8px; border: 2px solid #9c27b0;">
            <h3 style="margin-top: 0; color: #7b1fa2;">State (Estado)</h3>
            <ul style="margin-bottom: 0;">
              <li>Memória interna do componente</li>
              <li><strong>Mutável</strong> (pode ser atualizado)</li>
              <li>Local ao componente</li>
              <li>Causa re-renderização quando muda</li>
            </ul>
          </div>
        </div>

        <div style="background: #fff9c4; padding: 1rem; border-radius: 8px; border-left: 4px solid #fbc02d; margin: 1rem 0;">
          <p style="margin: 0;"><strong>💡 Dica:</strong> Use <strong>Props</strong> quando os dados vêm de fora. Use <strong>State</strong> quando os dados são internos e podem mudar.</p>
        </div>

        <h2 style="color: #2c3e50; border-bottom: 2px solid #3498db; padding-bottom: 0.5rem; margin-top: 2rem; margin-bottom: 1rem;">Conditional Rendering (Renderização Condicional)</h2>
        
        <p><strong>Conditional Rendering</strong> é a técnica de renderizar diferentes elementos ou componentes baseado em condições. Em React, isso funciona da mesma forma que condições em JavaScript.</p>
        
        <div style="background: #f5f5f5; padding: 1rem; border-radius: 8px; margin: 1rem 0;">
          <h3 style="margin-top: 0;">Métodos Comuns:</h3>
          <ul>
            <li><strong>Operador Ternário:</strong> <code>{'condicao ? <ComponenteA /> : <ComponenteB />'}</code></li>
            <li><strong>Operador &&:</strong> <code>{'condicao && <Componente />'}</code></li>
            <li><strong>Early Return:</strong> Retornar cedo se a condição não for satisfeita</li>
            <li><strong>Múltiplas Condições:</strong> Usar <code>if/else</code> ou <code>switch</code></li>
          </ul>
        </div>

        <h2 style="color: #2c3e50; border-bottom: 2px solid #3498db; padding-bottom: 0.5rem; margin-top: 2rem; margin-bottom: 1rem;">Composition (Composição)</h2>
        
        <p>React tem um <strong>modelo de composição poderoso</strong> e recomenda usar composição em vez de herança para reutilizar código entre componentes.</p>
        
        <p><strong>Composição</strong> significa construir componentes maiores combinando componentes menores. É como construir com blocos de Lego.</p>
        
        <div style="background: #e8f5e9; padding: 1rem; border-radius: 8px; border-left: 4px solid #4caf50; margin: 1rem 0;">
          <h3 style="margin-top: 0; color: #2e7d32;">Vantagens da Composição:</h3>
          <ul style="margin-bottom: 0;">
            <li><strong>Flexibilidade:</strong> Fácil de modificar e estender</li>
            <li><strong>Reutilização:</strong> Componentes pequenos podem ser combinados de várias formas</li>
            <li><strong>Testabilidade:</strong> Componentes pequenos são mais fáceis de testar</li>
            <li><strong>Manutenibilidade:</strong> Mudanças em um componente não afetam outros</li>
          </ul>
        </div>

        <h2 style="color: #2c3e50; border-bottom: 2px solid #3498db; padding-bottom: 0.5rem; margin-top: 2rem; margin-bottom: 1rem;">Props.children</h2>
        
        <p><strong>children</strong> é uma prop especial que contém o conteúdo entre as tags de abertura e fechamento de um componente. É muito útil para criar componentes genéricos e reutilizáveis.</p>
        
        <div style="background: #f5f5f5; padding: 1rem; border-radius: 8px; margin: 1rem 0;">
          <pre style="margin: 0; overflow-x: auto;"><code>function Card({ children }) {
  return (
    &lt;div className="card"&gt;
      {children}
    &lt;/div&gt;
  );
}

// Uso:
&lt;Card&gt;
  &lt;p&gt;Conteúdo aqui&lt;/p&gt;
  &lt;button&gt;Clique&lt;/button&gt;
&lt;/Card&gt;</code></pre>
        </div>

        <h2 style="color: #2c3e50; border-bottom: 2px solid #3498db; padding-bottom: 0.5rem; margin-top: 2rem; margin-bottom: 1rem;">Conceitos-Chave</h2>
        
        <div style="background: #f0f0f0; padding: 1.5rem; border-radius: 8px; margin: 1rem 0;">
          <ul style="margin: 0; line-height: 2;">
            <li><strong>Components:</strong> Blocos de construção reutilizáveis</li>
            <li><strong>Functional Components:</strong> Forma moderna de criar componentes</li>
            <li><strong>JSX:</strong> Sintaxe que permite HTML-like em JavaScript</li>
            <li><strong>Props:</strong> Dados passados de pai para filho (read-only)</li>
            <li><strong>State:</strong> Memória interna do componente (mutável)</li>
            <li><strong>Conditional Rendering:</strong> Renderizar baseado em condições</li>
            <li><strong>Composition:</strong> Combinar componentes menores para criar maiores</li>
          </ul>
        </div>

        <div style="background: #e3f2fd; padding: 1.5rem; border-radius: 8px; border: 2px solid #2196f3; margin: 2rem 0;">
          <h3 style="margin-top: 0; color: #1976d2;">🎯 Próximos Passos</h3>
          <p style="margin-bottom: 0;">Agora que você entende components, você pode:</p>
          <ul style="margin-top: 0.5rem; margin-bottom: 0;">
            <li>Criar componentes reutilizáveis</li>
            <li>Gerenciar estado e props</li>
            <li>Renderizar condicionalmente</li>
            <li>Compor interfaces complexas</li>
          </ul>
        </div>
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
  }
]
