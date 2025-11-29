# Aula 3: Rendering e Conceitos Avançados do React

## Introdução

Esta aula aborda conceitos fundamentais e avançados sobre como o React funciona internamente, desde o processo de renderização até padrões avançados de reutilização de código. Compreender esses conceitos é essencial para escrever código React eficiente, performático e manutenível.

---

## 1. Rendering (Renderização)

### 1.1 Abordagem Declarativa vs Imperativa

React segue uma **abordagem declarativa** para renderização de componentes. Isso significa que você, como desenvolvedor, especifica **o que** um componente deve parecer, e o React se encarrega de **como** renderizar esse componente na tela.

#### Abordagem Imperativa (JavaScript Puro)

Na abordagem imperativa, você escreve código que descreve **exatamente como** manipular o DOM:

```javascript
// Abordagem imperativa - você diz COMO fazer
const container = document.getElementById('app');
const button = document.createElement('button');
button.textContent = 'Clique aqui';
button.className = 'btn-primary';
button.onclick = function() {
  const counter = document.getElementById('counter');
  const currentValue = parseInt(counter.textContent);
  counter.textContent = currentValue + 1;
};
container.appendChild(button);

const counterElement = document.createElement('div');
counterElement.id = 'counter';
counterElement.textContent = '0';
container.appendChild(counterElement);
```

**Características da abordagem imperativa:**
- Você controla cada passo da manipulação do DOM
- Precisa gerenciar manualmente quando atualizar elementos
- Código verboso e propenso a erros
- Difícil sincronizar múltiplos elementos

#### Abordagem Declarativa (React)

Na abordagem declarativa, você descreve **o que** você quer ver:

```jsx
// Abordagem declarativa - você diz O QUE quer ver
function Counter() {
  const [count, setCount] = useState(0);
  
  return (
    <div>
      <button className="btn-primary" onClick={() => setCount(count + 1)}>
        Clique aqui
      </button>
      <div id="counter">{count}</div>
    </div>
  );
}
```

**Características da abordagem declarativa:**
- Você descreve o estado desejado da UI
- React decide como atualizar o DOM
- Código mais limpo e fácil de entender
- React sincroniza automaticamente todas as atualizações

**Por que a abordagem declarativa é melhor?**
1. **Menos erros**: Você não precisa se preocupar com detalhes de manipulação do DOM
2. **Código mais legível**: Fica claro qual é o estado desejado da interface
3. **Manutenção mais fácil**: Mudanças são feitas no estado, não na manipulação direta do DOM
4. **Performance otimizada**: React otimiza as atualizações automaticamente

### 1.2 O Processo de Renderização no React

O processo de renderização no React pode ser dividido em etapas claras:

#### Etapa 1: Componentes Retornam JSX

Quando você escreve um componente, ele retorna JSX que descreve a estrutura desejada:

```jsx
function ProductCard({ name, price }) {
  return (
    <div className="product-card">
      <h3>{name}</h3>
      <p>R$ {price}</p>
      <button>Comprar</button>
    </div>
  );
}
```

Este JSX é uma **descrição** do que deve aparecer, não o DOM real ainda.

#### Etapa 2: React Cria o Virtual DOM

React transforma o JSX em uma representação JavaScript chamada **Virtual DOM (VDOM)**. O Virtual DOM é uma estrutura de dados leve que representa a estrutura do DOM real.

```jsx
// O JSX acima é transformado em algo como:
{
  type: 'div',
  props: {
    className: 'product-card',
    children: [
      {
        type: 'h3',
        props: { children: name }
      },
      {
        type: 'p',
        props: { children: `R$ ${price}` }
      },
      {
        type: 'button',
        props: { children: 'Comprar' }
      }
    ]
  }
}
```

**Por que Virtual DOM?**
- É muito mais rápido criar e comparar objetos JavaScript do que manipular o DOM real
- Permite que React calcule as mudanças mínimas necessárias
- Facilita a renderização em diferentes ambientes (navegador, mobile, etc.)

#### Etapa 3: Comparação (Diffing/Reconciliation)

Quando o estado ou props de um componente mudam, React cria um **novo** Virtual DOM. Em seguida, React compara (faz "diff") o novo Virtual DOM com o anterior para identificar o que mudou.

```jsx
function Counter() {
  const [count, setCount] = useState(0);
  
  // Primeira renderização: count = 0
  // Virtual DOM: <div>Contador: 0</div>
  
  // Após setCount(1): count = 1
  // Novo Virtual DOM: <div>Contador: 1</div>
  // React compara: apenas o texto mudou, não precisa recriar o <div>
}
```

**Algoritmo de Diffing:**
React usa heurísticas para fazer a comparação de forma eficiente:
1. **Elementos de tipos diferentes**: React destrói a árvore antiga e constrói uma nova
2. **Elementos do mesmo tipo**: React atualiza apenas as propriedades que mudaram
3. **Componentes do mesmo tipo**: React mantém a instância e atualiza apenas as props
4. **Keys em listas**: React usa keys para identificar quais itens mudaram, foram adicionados ou removidos

#### Etapa 4: Atualização do DOM Real (Commit)

Após identificar as mudanças, React atualiza o DOM real com o **mínimo de operações necessárias**. Isso é chamado de "commit phase".

```jsx
// React não faz isso (ineficiente):
document.getElementById('counter').innerHTML = '1'; // Recria tudo

// React faz isso (eficiente):
// Atualiza apenas o texto do nó específico
textNode.nodeValue = '1';
```

**Exemplo prático de otimização:**

```jsx
function TodoList({ todos }) {
  return (
    <ul>
      {todos.map(todo => (
        <li key={todo.id}>{todo.text}</li>
      ))}
    </ul>
  );
}

// Se adicionarmos um novo todo no final:
// React NÃO recria todos os <li> existentes
// React apenas adiciona o novo <li> no final

// Se reordenarmos os todos:
// React usa as keys para identificar quais itens mudaram de posição
// React move apenas os elementos necessários
```

### 1.3 Virtual DOM em Detalhes

#### O Que é o Virtual DOM?

O **Virtual DOM (VDOM)** é uma representação em memória (JavaScript) da estrutura do DOM real. É uma "cópia leve" que React mantém para poder fazer comparações rápidas.

**Estrutura do Virtual DOM:**
```javascript
// Cada elemento no Virtual DOM é um objeto JavaScript
{
  type: 'div',                    // Tipo do elemento
  props: {                        // Propriedades
    className: 'container',
    onClick: handleClick,
    children: [...]               // Filhos (outros elementos VDOM)
  },
  key: null,                      // Key (para listas)
  ref: null                       // Ref (para acesso direto)
}
```

#### Por Que Virtual DOM é Rápido?

1. **Comparação em memória**: Comparar objetos JavaScript é muito mais rápido que manipular o DOM
2. **Batch updates**: React agrupa múltiplas atualizações e aplica todas de uma vez
3. **Cálculo mínimo**: React calcula exatamente o que precisa mudar, não recria tudo

**Exemplo de performance:**

```jsx
function App() {
  const [items, setItems] = useState([1, 2, 3, 4, 5]);
  
  // Se mudarmos apenas o primeiro item:
  setItems([10, 2, 3, 4, 5]);
  
  // React:
  // 1. Cria novo Virtual DOM
  // 2. Compara com o anterior
  // 3. Identifica: apenas o primeiro <li> mudou
  // 4. Atualiza APENAS o primeiro <li> no DOM real
  // 5. Os outros 4 <li> permanecem intocados
  
  return (
    <ul>
      {items.map(item => (
        <li key={item}>{item}</li>
      ))}
    </ul>
  );
}
```

#### Virtual DOM vs DOM Real

| Aspecto | DOM Real | Virtual DOM |
|---------|----------|-------------|
| **Localização** | Navegador (C++) | Memória JavaScript |
| **Manipulação** | Lenta (reflow/repaint) | Rápida (objetos JS) |
| **Custo de criação** | Alto | Baixo |
| **Custo de comparação** | Muito alto | Baixo |
| **Atualização** | Síncrona e cara | Assíncrona e barata |

### 1.4 Reconciliation (Reconciliação)

**Reconciliation** é o processo pelo qual React compara o Virtual DOM anterior com o novo e determina quais mudanças precisam ser aplicadas ao DOM real.

#### Como Funciona a Reconciliation

```jsx
// Estado inicial
function App() {
  const [count, setCount] = useState(0);
  return <div>Contador: {count}</div>;
}

// Virtual DOM inicial:
// { type: 'div', props: { children: 'Contador: 0' } }

// Após setCount(1):
// Novo Virtual DOM: { type: 'div', props: { children: 'Contador: 1' } }

// React compara:
// - type: 'div' === 'div' ✅ (mesmo tipo, mantém elemento)
// - props.children: 'Contador: 0' !== 'Contador: 1' ❌ (mudou)
// - Ação: Atualizar apenas o texto do nó de texto
```

#### Regras de Reconciliation

1. **Elementos de tipos diferentes:**
```jsx
// Antes:
<div>
  <Counter />
</div>

// Depois:
<span>
  <Counter />
</span>

// React destrói <div> e <Counter>, cria <span> e nova instância de <Counter>
```

2. **Elementos do mesmo tipo:**
```jsx
// Antes:
<div className="old">Conteúdo</div>

// Depois:
<div className="new">Conteúdo</div>

// React mantém o <div>, atualiza apenas className
```

3. **Componentes do mesmo tipo:**
```jsx
// Antes:
<Counter count={0} />

// Depois:
<Counter count={1} />

// React mantém a instância do componente, atualiza apenas as props
// O componente re-renderiza com novas props
```

#### Importância das Keys na Reconciliation

Keys ajudam React a identificar quais itens em uma lista mudaram:

```jsx
// ❌ SEM keys (React não sabe qual item mudou)
{todos.map(todo => (
  <TodoItem todo={todo} />
))}

// ✅ COM keys (React identifica cada item)
{todos.map(todo => (
  <TodoItem key={todo.id} todo={todo} />
))}
```

**Exemplo prático:**

```jsx
// Lista inicial: [A, B, C]
// Virtual DOM: [
//   { type: TodoItem, key: 'A', props: { todo: A } },
//   { type: TodoItem, key: 'B', props: { todo: B } },
//   { type: TodoItem, key: 'C', props: { todo: C } }
// ]

// Nova lista: [A, C, D] (B removido, D adicionado, C movido)
// Novo Virtual DOM: [
//   { type: TodoItem, key: 'A', props: { todo: A } }, // ✅ Mesmo, mantém
//   { type: TodoItem, key: 'C', props: { todo: C } }, // ✅ Movido, move no DOM
//   { type: TodoItem, key: 'D', props: { todo: D } }  // ❌ Novo, cria
// ]

// React:
// - Mantém instância de A (mesma posição)
// - Move instância de C (nova posição)
// - Remove instância de B (não existe mais)
// - Cria nova instância de D
```

### 1.5 Renderização de Componentes

#### Componentes como Funções

Componentes React são essencialmente funções que retornam uma descrição (JSX) do que deve aparecer:

```jsx
// Componente como função
function Welcome({ name }) {
  return <h1>Bem-vindo, {name}!</h1>;
}

// Quando você usa <Welcome name="João" />, React:
// 1. Chama a função Welcome({ name: "João" })
// 2. Recebe o JSX retornado
// 3. Transforma em Virtual DOM
// 4. Renderiza no DOM real
```

#### Render Method

Em componentes funcionais, o "render method" é simplesmente o retorno da função:

```jsx
function Component() {
  // Qualquer lógica aqui
  const value = calculateSomething();
  
  // O "render" é o return
  return (
    <div>
      <p>{value}</p>
    </div>
  );
}
```

Em componentes de classe (legado), havia um método `render()` explícito:

```jsx
// Componente de classe (não recomendado, apenas para entender)
class Component extends React.Component {
  render() {
    return <div>Conteúdo</div>;
  }
}
```

#### Quando React Renderiza?

React renderiza um componente nas seguintes situações:

1. **Renderização inicial (Mount):**
```jsx
// Quando o componente aparece pela primeira vez
function App() {
  return <Counter />; // Counter é renderizado pela primeira vez
}
```

2. **Quando estado muda:**
```jsx
function Counter() {
  const [count, setCount] = useState(0);
  
  // Cada vez que setCount é chamado, React re-renderiza Counter
  return <div>{count}</div>;
}
```

3. **Quando props mudam:**
```jsx
function Parent() {
  const [name, setName] = useState('João');
  
  // Quando name muda, Child é re-renderizado com nova prop
  return <Child name={name} />;
}
```

4. **Quando componente pai re-renderiza:**
```jsx
function Parent() {
  const [count, setCount] = useState(0);
  
  // Quando Parent re-renderiza, Child também re-renderiza
  // (mesmo que as props não tenham mudado)
  return (
    <div>
      <p>Count: {count}</p>
      <Child name="João" /> {/* Re-renderiza mesmo com props iguais */}
    </div>
  );
}
```

**Importante:** React é inteligente e pode otimizar re-renderizações desnecessárias usando `React.memo`, `useMemo`, e `useCallback` (veremos em aulas futuras).

### 1.6 Exemplo Prático Completo

Vamos ver um exemplo completo que demonstra todo o processo de renderização:

```jsx
import { useState } from 'react';

function ProductList() {
  const [products, setProducts] = useState([
    { id: 1, name: 'Notebook', price: 2500 },
    { id: 2, name: 'Mouse', price: 50 },
    { id: 3, name: 'Teclado', price: 150 }
  ]);
  
  const [filter, setFilter] = useState('');
  
  // Filtrar produtos
  const filteredProducts = products.filter(product =>
    product.name.toLowerCase().includes(filter.toLowerCase())
  );
  
  return (
    <div>
      <input
        type="text"
        placeholder="Filtrar produtos..."
        value={filter}
        onChange={(e) => setFilter(e.target.value)}
      />
      
      <ul>
        {filteredProducts.map(product => (
          <ProductItem key={product.id} product={product} />
        ))}
      </ul>
    </div>
  );
}

function ProductItem({ product }) {
  return (
    <li>
      <strong>{product.name}</strong> - R$ {product.price}
    </li>
  );
}
```

**O que acontece quando o usuário digita no input:**

1. **Evento onChange dispara** → `setFilter` é chamado
2. **Estado `filter` muda** → React marca `ProductList` para re-renderização
3. **React chama ProductList novamente** → Função executa, calcula `filteredProducts`
4. **React cria novo Virtual DOM** → Compara com o Virtual DOM anterior
5. **React identifica mudanças:**
   - O valor do `<input>` mudou → Atualiza o atributo `value`
   - A lista `filteredProducts` pode ter mudado → Compara usando `key={product.id}`
   - Se um produto foi removido da lista filtrada → Remove o `<li>` correspondente
   - Se um produto foi adicionado → Adiciona novo `<li>`
6. **React atualiza o DOM real** → Apenas os elementos que mudaram

**Por que isso é eficiente:**
- React não recria toda a lista
- React usa as keys para identificar quais itens mudaram
- Apenas os elementos necessários são atualizados no DOM real

---

## 2. Component Life Cycle (Ciclo de Vida de Componentes)

### 2.1 Introdução ao Ciclo de Vida

Componentes React têm um **ciclo de vida** que consiste em três fases principais:

1. **Mounting (Montagem)**: Quando o componente é criado e inserido no DOM pela primeira vez
2. **Updating (Atualização)**: Quando o componente é re-renderizado devido a mudanças em props ou estado
3. **Unmounting (Desmontagem)**: Quando o componente é removido do DOM

Cada fase tem **lifecycle methods** (métodos de ciclo de vida) que você pode usar para executar código em momentos específicos.

**Importante:** Em componentes funcionais modernos, usamos **hooks** (especialmente `useEffect`) em vez de lifecycle methods. Lifecycle methods existem apenas em componentes de classe, que não são mais recomendados.

### 2.2 Fase 1: Mounting (Montagem)

A fase de **mounting** ocorre quando um componente é criado e inserido no DOM pela primeira vez.

#### O Que Acontece Durante o Mounting?

```jsx
// Quando você renderiza um componente pela primeira vez:
<Counter initialCount={0} />

// React:
// 1. Cria a instância do componente
// 2. Inicializa o estado (se houver)
// 3. Executa a primeira renderização
// 4. Insere o componente no DOM
// 5. Executa efeitos (useEffect com dependências vazias)
```

#### Lifecycle Methods de Mounting (Componentes de Classe)

Em componentes de classe, havia métodos específicos para cada etapa:

```jsx
// ❌ Componente de classe (não recomendado, apenas para entender)
class Counter extends React.Component {
  constructor(props) {
    super(props);
    // 1. CONSTRUCTOR: Inicializa estado e bind de métodos
    this.state = { count: 0 };
    console.log('1. Constructor executado');
  }
  
  static getDerivedStateFromProps(props, state) {
    // 2. GET_DERIVED_STATE_FROM_PROPS: Atualiza estado baseado em props
    console.log('2. getDerivedStateFromProps executado');
    return null;
  }
  
  componentDidMount() {
    // 3. COMPONENT_DID_MOUNT: Após componente ser inserido no DOM
    console.log('3. componentDidMount executado');
    // Ideal para: chamadas de API, subscriptions, manipulação de DOM
  }
  
  render() {
    // RENDER: Cria o Virtual DOM
    console.log('Render executado');
    return <div>{this.state.count}</div>;
  }
}
```

**Ordem de execução durante mounting:**
1. `constructor()`
2. `static getDerivedStateFromProps()`
3. `render()`
4. `componentDidMount()`

#### Equivalente com Hooks (Componentes Funcionais)

Em componentes funcionais modernos, usamos `useEffect` para replicar o comportamento de `componentDidMount`:

```jsx
// ✅ Componente funcional moderno (recomendado)
import { useState, useEffect } from 'react';

function Counter({ initialCount }) {
  // Equivalente ao constructor + useState
  const [count, setCount] = useState(initialCount);
  
  // Equivalente ao componentDidMount
  useEffect(() => {
    console.log('Componente montado!');
    // Ideal para: chamadas de API, subscriptions, manipulação de DOM
    
    // Cleanup (equivalente ao componentWillUnmount)
    return () => {
      console.log('Componente será desmontado');
      // Limpar subscriptions, timers, etc.
    };
  }, []); // Array vazio = executa apenas no mount
  
  return <div>{count}</div>;
}
```

**Por que hooks são melhores:**
- Código mais simples e legível
- Lógica relacionada fica junta
- Mais fácil de testar
- Evita problemas com `this`

#### Exemplo Prático: Buscar Dados na Montagem

```jsx
import { useState, useEffect } from 'react';

function UserProfile({ userId }) {
  const [user, setUser] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  
  // Executa apenas quando componente é montado (e quando userId muda)
  useEffect(() => {
    // Função assíncrona para buscar dados
    async function fetchUser() {
      try {
        setLoading(true);
        const response = await fetch(`/api/users/${userId}`);
        const userData = await response.json();
        setUser(userData);
        setError(null);
      } catch (err) {
        setError(err.message);
      } finally {
        setLoading(false);
      }
    }
    
    fetchUser();
  }, [userId]); // Dependência: re-executa se userId mudar
  
  if (loading) return <div>Carregando...</div>;
  if (error) return <div>Erro: {error}</div>;
  if (!user) return <div>Usuário não encontrado</div>;
  
  return (
    <div>
      <h1>{user.name}</h1>
      <p>{user.email}</p>
    </div>
  );
}
```

**O que acontece:**
1. Componente é montado → `useEffect` executa
2. `fetchUser()` é chamado → Busca dados da API
3. Estado é atualizado → Componente re-renderiza
4. UI mostra os dados do usuário

### 2.3 Fase 2: Updating (Atualização)

A fase de **updating** ocorre quando um componente é re-renderizado devido a mudanças em props ou estado.

#### O Que Acontece Durante o Updating?

```jsx
// Componente já montado, estado muda:
function Counter() {
  const [count, setCount] = useState(0);
  
  // Quando setCount(1) é chamado:
  // 1. Estado muda
  // 2. React marca componente para re-renderização
  // 3. React chama o componente novamente
  // 4. React cria novo Virtual DOM
  // 5. React compara com Virtual DOM anterior
  // 6. React atualiza DOM real (se necessário)
  // 7. React executa useEffect (se dependências mudaram)
}
```

#### Lifecycle Methods de Updating (Componentes de Classe)

```jsx
// ❌ Componente de classe (não recomendado)
class Counter extends React.Component {
  static getDerivedStateFromProps(props, state) {
    // 1. Executado ANTES de cada render
    // Atualiza estado baseado em mudanças de props
    return null;
  }
  
  shouldComponentUpdate(nextProps, nextState) {
    // 2. Executado ANTES do render
    // Retorna true/false para decidir se deve re-renderizar
    // Útil para otimização
    return true; // ou false para evitar re-render
  }
  
  render() {
    // 3. RENDER: Cria novo Virtual DOM
    return <div>{this.state.count}</div>;
  }
  
  getSnapshotBeforeUpdate(prevProps, prevState) {
    // 4. Executado ANTES da atualização do DOM
    // Captura informações do DOM antes de mudar
    return null;
  }
  
  componentDidUpdate(prevProps, prevState, snapshot) {
    // 5. Executado APÓS atualização do DOM
    // Ideal para: manipulação de DOM, chamadas de API baseadas em mudanças
    console.log('Componente atualizado!');
  }
}
```

**Ordem de execução durante updating:**
1. `static getDerivedStateFromProps()`
2. `shouldComponentUpdate()`
3. `render()`
4. `getSnapshotBeforeUpdate()`
5. `componentDidUpdate()`

#### Equivalente com Hooks (Componentes Funcionais)

```jsx
// ✅ Componente funcional moderno
import { useState, useEffect, useMemo } from 'react';

function Counter({ initialCount }) {
  const [count, setCount] = useState(initialCount);
  
  // Equivalente ao componentDidUpdate
  useEffect(() => {
    console.log('Count mudou para:', count);
    // Executa após cada render onde count mudou
  }, [count]); // Re-executa quando count muda
  
  // Equivalente ao shouldComponentUpdate
  const expensiveValue = useMemo(() => {
    // Só recalcula se count mudar
    return expensiveCalculation(count);
  }, [count]);
  
  return <div>{count}</div>;
}
```

#### Exemplo Prático: Atualizar Dados Quando Props Mudam

```jsx
import { useState, useEffect } from 'react';

function ProductDetails({ productId }) {
  const [product, setProduct] = useState(null);
  const [loading, setLoading] = useState(true);
  
  // Re-executa sempre que productId mudar
  useEffect(() => {
    async function fetchProduct() {
      setLoading(true);
      const response = await fetch(`/api/products/${productId}`);
      const data = await response.json();
      setProduct(data);
      setLoading(false);
    }
    
    fetchProduct();
  }, [productId]); // Dependência: re-executa se productId mudar
  
  if (loading) return <div>Carregando produto...</div>;
  if (!product) return <div>Produto não encontrado</div>;
  
  return (
    <div>
      <h1>{product.name}</h1>
      <p>{product.description}</p>
      <p>R$ {product.price}</p>
    </div>
  );
}

// Uso:
function App() {
  const [selectedId, setSelectedId] = useState(1);
  
  return (
    <div>
      <button onClick={() => setSelectedId(1)}>Produto 1</button>
      <button onClick={() => setSelectedId(2)}>Produto 2</button>
      <ProductDetails productId={selectedId} />
      {/* Quando selectedId muda, ProductDetails re-executa useEffect */}
    </div>
  );
}
```

### 2.4 Fase 3: Unmounting (Desmontagem)

A fase de **unmounting** ocorre quando um componente é removido do DOM.

#### O Que Acontece Durante o Unmounting?

```jsx
// Componente está montado:
<Counter />

// Quando você remove o componente:
{showCounter && <Counter />}
// Se showCounter vira false, Counter é desmontado

// React:
// 1. Executa cleanup de useEffect (se houver)
// 2. Remove componente do DOM
// 3. Destrói a instância do componente
```

#### Lifecycle Method de Unmounting (Componentes de Classe)

```jsx
// ❌ Componente de classe
class Counter extends React.Component {
  componentWillUnmount() {
    // Executado ANTES do componente ser removido
    // Ideal para: limpar subscriptions, timers, cancelar requisições
    console.log('Componente será desmontado');
    // Limpar recursos aqui
  }
  
  render() {
    return <div>Counter</div>;
  }
}
```

#### Equivalente com Hooks (Componentes Funcionais)

```jsx
// ✅ Componente funcional moderno
import { useEffect } from 'react';

function Counter() {
  useEffect(() => {
    // Código que executa no mount/update
    const timer = setInterval(() => {
      console.log('Timer executando...');
    }, 1000);
    
    // Cleanup: executa no unmount (e antes de re-executar o effect)
    return () => {
      console.log('Limpando timer...');
      clearInterval(timer);
    };
  }, []);
  
  return <div>Counter</div>;
}
```

**A função de cleanup:**
- É retornada pelo `useEffect`
- Executa quando o componente é desmontado
- Também executa antes de re-executar o effect (se dependências mudarem)

#### Exemplo Prático: Limpar Subscriptions

```jsx
import { useState, useEffect } from 'react';

function ChatRoom({ roomId }) {
  const [messages, setMessages] = useState([]);
  
  useEffect(() => {
    // Simular subscription a um chat
    const subscription = {
      subscribe: (callback) => {
        // Simular recebimento de mensagens
        const interval = setInterval(() => {
          callback({ text: 'Nova mensagem', timestamp: Date.now() });
        }, 2000);
        return () => clearInterval(interval);
      }
    };
    
    const unsubscribe = subscription.subscribe((message) => {
      setMessages(prev => [...prev, message]);
    });
    
    // Cleanup: desinscrever quando componente desmontar
    return () => {
      console.log('Desinscrevendo do chat...');
      unsubscribe();
    };
  }, [roomId]);
  
  return (
    <div>
      <h2>Sala: {roomId}</h2>
      <ul>
        {messages.map((msg, idx) => (
          <li key={idx}>{msg.text}</li>
        ))}
      </ul>
    </div>
  );
}

// Uso:
function App() {
  const [showChat, setShowChat] = useState(true);
  const [roomId, setRoomId] = useState('general');
  
  return (
    <div>
      <button onClick={() => setShowChat(!showChat)}>
        {showChat ? 'Fechar Chat' : 'Abrir Chat'}
      </button>
      {showChat && <ChatRoom roomId={roomId} />}
      {/* Quando showChat vira false, ChatRoom é desmontado e cleanup executa */}
    </div>
  );
}
```

### 2.5 Resumo do Ciclo de Vida

#### Fluxo Completo

```
MOUNTING (Primeira vez)
├── useState inicializa estado
├── Componente renderiza (primeira vez)
├── useEffect executa (se dependências permitirem)
└── Componente está no DOM

UPDATING (Re-renderizações)
├── Estado ou props mudam
├── Componente re-renderiza
├── useEffect re-executa (se dependências mudaram)
└── DOM é atualizado (se necessário)

UNMOUNTING (Remoção)
├── Cleanup de useEffect executa
├── Componente é removido do DOM
└── Instância é destruída
```

#### Tabela Comparativa: Lifecycle Methods vs Hooks

| Lifecycle Method (Classe) | Hook Equivalente (Funcional) |
|---------------------------|------------------------------|
| `constructor` | `useState` (inicialização) |
| `componentDidMount` | `useEffect(() => {...}, [])` |
| `componentDidUpdate` | `useEffect(() => {...}, [deps])` |
| `componentWillUnmount` | `useEffect(() => { return cleanup }, [])` |
| `shouldComponentUpdate` | `React.memo` ou `useMemo` |
| `getDerivedStateFromProps` | `useState` com lógica no corpo |

### 2.6 Boas Práticas com useEffect

#### 1. Sempre Inclua Dependências

```jsx
// ❌ ERRADO: falta dependência
function Counter({ step }) {
  const [count, setCount] = useState(0);
  
  useEffect(() => {
    // Usa 'step' mas não está nas dependências
    setCount(count + step);
  }, []); // Array vazio = executa apenas no mount
}

// ✅ CORRETO: todas as dependências incluídas
function Counter({ step }) {
  const [count, setCount] = useState(0);
  
  useEffect(() => {
    setCount(prev => prev + step);
  }, [step]); // step está nas dependências
}
```

#### 2. Use Cleanup para Recursos

```jsx
// ✅ Sempre limpe timers, subscriptions, etc.
useEffect(() => {
  const timer = setInterval(() => {
    // ...
  }, 1000);
  
  return () => clearInterval(timer); // Cleanup
}, []);
```

#### 3. Separe Efeitos por Responsabilidade

```jsx
// ❌ ERRADO: múltiplas responsabilidades em um effect
useEffect(() => {
  fetchData();
  setupSubscription();
  updateDocumentTitle();
}, []);

// ✅ CORRETO: um effect por responsabilidade
useEffect(() => {
  fetchData();
}, []);

useEffect(() => {
  setupSubscription();
  return () => cleanupSubscription();
}, []);

useEffect(() => {
  updateDocumentTitle();
}, [title]);
```

---

## 📝 Resumo Parcial da Aula

Até agora, cobrimos:

### ✅ Rendering (Renderização)
- Abordagem declarativa vs imperativa
- Processo de renderização no React
- Virtual DOM e sua importância
- Reconciliation (reconciliação)
- Quando React renderiza componentes

### ✅ Component Life Cycle (Ciclo de Vida)
- Fase de Mounting (montagem)
- Fase de Updating (atualização)
- Fase de Unmounting (desmontagem)
- Lifecycle methods vs Hooks modernos
- Uso correto de `useEffect`

---

## 3. Lists and Keys (Listas e Chaves)

### 3.1 Por Que Keys São Essenciais?

Quando você renderiza uma lista de elementos em React, cada elemento precisa de uma **key** única. As keys ajudam React a identificar quais itens mudaram, foram adicionados ou removidos.

#### O Problema Sem Keys

```jsx
// ❌ SEM keys - React não consegue identificar itens
function TodoList({ todos }) {
  return (
    <ul>
      {todos.map(todo => (
        <li>{todo.text}</li>
      ))}
    </ul>
  );
}
```

**O que acontece sem keys:**
- React não consegue rastrear qual item é qual
- Quando a lista muda, React pode recriar todos os elementos
- Estado interno de componentes filhos pode ser perdido
- Performance ruim em listas grandes

#### A Solução: Keys

```jsx
// ✅ COM keys - React identifica cada item
function TodoList({ todos }) {
  return (
    <ul>
      {todos.map(todo => (
        <li key={todo.id}>{todo.text}</li>
      ))}
    </ul>
  );
}
```

**O que acontece com keys:**
- React identifica cada item pela key
- Quando a lista muda, React atualiza apenas o necessário
- Estado interno é preservado corretamente
- Performance otimizada

### 3.2 Como React Usa Keys

React usa keys para fazer a reconciliação (diffing) de listas:

```jsx
// Lista inicial
const todos = [
  { id: 1, text: 'Aprender React' },
  { id: 2, text: 'Praticar hooks' },
  { id: 3, text: 'Construir app' }
];

// Virtual DOM inicial:
// [
//   { type: 'li', key: '1', props: { children: 'Aprender React' } },
//   { type: 'li', key: '2', props: { children: 'Praticar hooks' } },
//   { type: 'li', key: '3', props: { children: 'Construir app' } }
// ]

// Nova lista (removido item 2, adicionado item 4)
const newTodos = [
  { id: 1, text: 'Aprender React' },
  { id: 3, text: 'Construir app' },
  { id: 4, text: 'Deploy' }
];

// Novo Virtual DOM:
// [
//   { type: 'li', key: '1', props: { children: 'Aprender React' } }, // ✅ Mesmo, mantém
//   { type: 'li', key: '3', props: { children: 'Construir app' } },   // ✅ Movido, move no DOM
//   { type: 'li', key: '4', props: { children: 'Deploy' } }          // ❌ Novo, cria
// ]

// React:
// - Mantém instância do item com key='1'
// - Move instância do item com key='3' para nova posição
// - Remove instância do item com key='2' (não existe mais)
// - Cria nova instância do item com key='4'
```

### 3.3 Como Escolher Boas Keys

#### Regras para Keys

1. **Keys devem ser únicas** entre irmãos (não globalmente)
2. **Keys devem ser estáveis** (não mudar entre renders)
3. **Keys devem ser previsíveis** (não aleatórias)

#### ✅ Boas Keys

```jsx
// 1. IDs únicos do banco de dados (MELHOR OPÇÃO)
const users = [
  { id: 101, name: 'João' },
  { id: 102, name: 'Maria' }
];
{users.map(user => <User key={user.id} user={user} />)}

// 2. IDs gerados de forma estável
const items = todos.map((todo, index) => ({
  ...todo,
  stableId: `todo-${todo.id}-${todo.createdAt}`
}));
{items.map(item => <TodoItem key={item.stableId} item={item} />)}

// 3. Combinação de propriedades únicas
const products = [
  { category: 'electronics', sku: 'ELC-001' },
  { category: 'books', isbn: '978-1234567890' }
];
{products.map(product => (
  <Product key={product.sku || product.isbn} product={product} />
))}
```

#### ❌ Keys Ruins

```jsx
// ❌ Índice como key (só se lista nunca reordena/remove)
{todos.map((todo, index) => (
  <TodoItem key={index} todo={todo} />
))}
// Problema: Se você remover o primeiro item, todos os índices mudam
// React pensa que todos os itens mudaram, não apenas o removido

// ❌ Keys aleatórias
{todos.map(todo => (
  <TodoItem key={Math.random()} todo={todo} />
))}
// Problema: Key muda a cada render, React recria todos os elementos

// ❌ Keys que mudam
{todos.map(todo => (
  <TodoItem key={todo.text} todo={todo} />
))}
// Problema: Se o texto mudar, a key muda, React recria o elemento
```

### 3.4 Exemplo Prático: Problema com Índice como Key

```jsx
// ❌ PROBLEMA: Usando índice como key
function TodoList() {
  const [todos, setTodos] = useState([
    { id: 1, text: 'Aprender React', completed: false },
    { id: 2, text: 'Praticar hooks', completed: false },
    { id: 3, text: 'Construir app', completed: false }
  ]);
  
  const removeTodo = (index) => {
    setTodos(todos.filter((_, i) => i !== index));
  };
  
  return (
    <ul>
      {todos.map((todo, index) => (
        <TodoItem 
          key={index} // ❌ PROBLEMA: índice como key
          todo={todo}
          onRemove={() => removeTodo(index)}
        />
      ))}
    </ul>
  );
}

function TodoItem({ todo, onRemove }) {
  const [isEditing, setIsEditing] = useState(false);
  
  return (
    <li>
      {isEditing ? (
        <input defaultValue={todo.text} />
      ) : (
        <span>{todo.text}</span>
      )}
      <button onClick={onRemove}>Remover</button>
      <button onClick={() => setIsEditing(!isEditing)}>
        {isEditing ? 'Salvar' : 'Editar'}
      </button>
    </li>
  );
}
```

**O que acontece quando você remove o primeiro item:**
1. Lista muda de `[A, B, C]` para `[B, C]`
2. Índices mudam: `[0, 1, 2]` para `[0, 1]`
3. React pensa:
   - Item com key=0: era A, agora é B → **atualiza** (errado!)
   - Item com key=1: era B, agora é C → **atualiza** (errado!)
   - Item com key=2: era C, agora não existe → **remove** (correto)
4. Estado interno (`isEditing`) é perdido ou atribuído ao item errado

**Solução: usar ID único**

```jsx
// ✅ SOLUÇÃO: Usando ID único como key
function TodoList() {
  const [todos, setTodos] = useState([
    { id: 1, text: 'Aprender React', completed: false },
    { id: 2, text: 'Praticar hooks', completed: false },
    { id: 3, text: 'Construir app', completed: false }
  ]);
  
  const removeTodo = (id) => {
    setTodos(todos.filter(todo => todo.id !== id));
  };
  
  return (
    <ul>
      {todos.map(todo => (
        <TodoItem 
          key={todo.id} // ✅ SOLUÇÃO: ID único como key
          todo={todo}
          onRemove={() => removeTodo(todo.id)}
        />
      ))}
    </ul>
  );
}
```

**Agora quando você remove o primeiro item:**
1. Lista muda de `[A(id:1), B(id:2), C(id:3)]` para `[B(id:2), C(id:3)]`
2. React pensa:
   - Item com key=1: não existe mais → **remove** (correto!)
   - Item com key=2: ainda é B → **mantém** (correto!)
   - Item com key=3: ainda é C → **mantém** (correto!)
3. Estado interno é preservado corretamente

### 3.5 Quando Índice Pode Ser Usado como Key?

Índice como key é aceitável **apenas** quando:
- A lista **nunca** será reordenada
- Itens **nunca** serão removidos ou adicionados no meio
- Não há estado interno nos componentes filhos
- Performance não é crítica

```jsx
// ✅ Aceitável: Lista estática que nunca muda
const staticMenuItems = ['Home', 'About', 'Contact'];
{staticMenuItems.map((item, index) => (
  <MenuItem key={index} label={item} />
))}
```

**Mas mesmo assim, é melhor usar uma key mais estável:**

```jsx
// ✅ Melhor: Key baseada no conteúdo
{staticMenuItems.map(item => (
  <MenuItem key={item} label={item} />
))}
```

### 3.6 Renderização Eficiente de Listas

#### Otimização com React.memo

Para listas grandes, você pode otimizar componentes filhos:

```jsx
import { memo } from 'react';

// Componente memoizado: só re-renderiza se props mudarem
const TodoItem = memo(function TodoItem({ todo, onToggle }) {
  return (
    <li>
      <input
        type="checkbox"
        checked={todo.completed}
        onChange={() => onToggle(todo.id)}
      />
      <span>{todo.text}</span>
    </li>
  );
});

function TodoList({ todos, onToggle }) {
  return (
    <ul>
      {todos.map(todo => (
        <TodoItem 
          key={todo.id}
          todo={todo}
          onToggle={onToggle}
        />
      ))}
    </ul>
  );
}
```

**Por que isso ajuda:**
- Se `todos` não mudou, `TodoItem` não re-renderiza
- Apenas itens que mudaram são re-renderizados
- Performance muito melhor em listas grandes

#### Virtualização para Listas Muito Grandes

Para listas com milhares de itens, considere virtualização:

```jsx
// Usando react-window ou react-virtualized
import { FixedSizeList } from 'react-window';

function VirtualizedList({ items }) {
  const Row = ({ index, style }) => (
    <div style={style}>
      <Item item={items[index]} />
    </div>
  );
  
  return (
    <FixedSizeList
      height={600}
      itemCount={items.length}
      itemSize={50}
      width="100%"
    >
      {Row}
    </FixedSizeList>
  );
}
```

**Virtualização:**
- Renderiza apenas itens visíveis na tela
- Melhora drasticamente performance
- Útil para listas com 1000+ itens

### 3.7 Exemplo Completo: Lista de Produtos

```jsx
import { useState, memo } from 'react';

// Componente memoizado
const ProductCard = memo(function ProductCard({ product, onAddToCart }) {
  return (
    <div className="product-card">
      <img src={product.image} alt={product.name} />
      <h3>{product.name}</h3>
      <p>R$ {product.price}</p>
      <button onClick={() => onAddToCart(product.id)}>
        Adicionar ao Carrinho
      </button>
    </div>
  );
});

function ProductList({ products, onAddToCart }) {
  const [filter, setFilter] = useState('');
  const [sortBy, setSortBy] = useState('name');
  
  // Filtrar e ordenar produtos
  const filteredProducts = products
    .filter(product => 
      product.name.toLowerCase().includes(filter.toLowerCase())
    )
    .sort((a, b) => {
      if (sortBy === 'name') return a.name.localeCompare(b.name);
      if (sortBy === 'price') return a.price - b.price;
      return 0;
    });
  
  return (
    <div>
      <input
        type="text"
        placeholder="Filtrar produtos..."
        value={filter}
        onChange={(e) => setFilter(e.target.value)}
      />
      <select value={sortBy} onChange={(e) => setSortBy(e.target.value)}>
        <option value="name">Nome</option>
        <option value="price">Preço</option>
      </select>
      
      <div className="product-grid">
        {filteredProducts.map(product => (
          <ProductCard
            key={product.id} // ✅ ID único do banco de dados
            product={product}
            onAddToCart={onAddToCart}
          />
        ))}
      </div>
    </div>
  );
}
```

**Por que este código é eficiente:**
- Keys únicas e estáveis (`product.id`)
- Componentes memoizados (`React.memo`)
- React só atualiza produtos que mudaram
- Estado interno é preservado corretamente

---

## 4. Render Props

### 4.1 O Que São Render Props?

**Render Props** é uma técnica para compartilhar código entre componentes React usando uma prop cujo valor é uma **função** que retorna JSX.

O termo "render prop" refere-se a uma técnica onde um componente recebe uma função como prop e chama essa função em vez de implementar sua própria lógica de renderização.

#### Estrutura Básica

```jsx
// Componente com render prop
function DataFetcher({ url, render }) {
  const [data, setData] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  
  useEffect(() => {
    fetch(url)
      .then(res => res.json())
      .then(data => {
        setData(data);
        setLoading(false);
      })
      .catch(err => {
        setError(err);
        setLoading(false);
      });
  }, [url]);
  
  // Chama a função render com os dados
  return render({ data, loading, error });
}

// Uso: passa uma função como prop
function App() {
  return (
    <DataFetcher
      url="/api/users"
      render={({ data, loading, error }) => {
        if (loading) return <div>Carregando...</div>;
        if (error) return <div>Erro: {error.message}</div>;
        return (
          <ul>
            {data.map(user => (
              <li key={user.id}>{user.name}</li>
            ))}
          </ul>
        );
      }}
    />
  );
}
```

### 4.2 Por Que Usar Render Props?

Render props permitem:
1. **Compartilhar lógica** entre componentes
2. **Flexibilidade** na renderização
3. **Reutilização** de código
4. **Separação de responsabilidades**

#### Exemplo: Compartilhar Lógica de Mouse Position

```jsx
// Componente que gerencia posição do mouse
function MouseTracker({ render }) {
  const [position, setPosition] = useState({ x: 0, y: 0 });
  
  useEffect(() => {
    const handleMouseMove = (e) => {
      setPosition({ x: e.clientX, y: e.clientY });
    };
    
    window.addEventListener('mousemove', handleMouseMove);
    return () => window.removeEventListener('mousemove', handleMouseMove);
  }, []);
  
  return render(position);
}

// Uso 1: Mostrar coordenadas
function App() {
  return (
    <MouseTracker
      render={({ x, y }) => (
        <div>
          Mouse está em: ({x}, {y})
        </div>
      )}
    />
  );
}

// Uso 2: Imagem que segue o mouse
function CatApp() {
  return (
    <MouseTracker
      render={({ x, y }) => (
        <img
          src="/cat.png"
          style={{
            position: 'absolute',
            left: x,
            top: y,
            pointerEvents: 'none'
          }}
          alt="Cat"
        />
      )}
    />
  );
}
```

### 4.3 Convenções de Nomeação

Embora o nome da prop possa ser qualquer coisa, é comum usar:
- `render` - mais comum
- `children` - quando usado como children (mais idiomático)
- Nomes específicos como `renderHeader`, `renderBody`

#### Usando `children` como Render Prop

```jsx
// Componente que usa children como função
function MouseTracker({ children }) {
  const [position, setPosition] = useState({ x: 0, y: 0 });
  
  useEffect(() => {
    const handleMouseMove = (e) => {
      setPosition({ x: e.clientX, y: e.clientY });
    };
    
    window.addEventListener('mousemove', handleMouseMove);
    return () => window.removeEventListener('mousemove', handleMouseMove);
  }, []);
  
  // children é uma função, não JSX
  return children(position);
}

// Uso mais idiomático
function App() {
  return (
    <MouseTracker>
      {({ x, y }) => (
        <div>
          Mouse: ({x}, {y})
        </div>
      )}
    </MouseTracker>
  );
}
```

### 4.4 Exemplo Prático: Toggle Component

```jsx
// Componente genérico de toggle
function Toggle({ children, onToggle }) {
  const [isOn, setIsOn] = useState(false);
  
  const toggle = () => {
    setIsOn(!isOn);
    if (onToggle) {
      onToggle(!isOn);
    }
  };
  
  return children({ isOn, toggle });
}

// Uso 1: Switch simples
function SwitchExample() {
  return (
    <Toggle>
      {({ isOn, toggle }) => (
        <div>
          <button onClick={toggle}>
            {isOn ? 'ON' : 'OFF'}
          </button>
        </div>
      )}
    </Toggle>
  );
}

// Uso 2: Checkbox customizado
function CustomCheckbox({ label }) {
  return (
    <Toggle>
      {({ isOn, toggle }) => (
        <div onClick={toggle} style={{ cursor: 'pointer' }}>
          <input type="checkbox" checked={isOn} readOnly />
          <label>{label}</label>
        </div>
      )}
    </Toggle>
  );
}

// Uso 3: Accordion
function Accordion({ title, content }) {
  return (
    <Toggle>
      {({ isOn, toggle }) => (
        <div>
          <button onClick={toggle}>{title}</button>
          {isOn && <div>{content}</div>}
        </div>
      )}
    </Toggle>
  );
}
```

### 4.5 Render Props vs Outros Padrões

#### Render Props vs Hooks

```jsx
// ✅ Render Props
function MouseTracker({ children }) {
  const [position, setPosition] = useState({ x: 0, y: 0 });
  // ... lógica
  return children(position);
}

// ✅ Hook customizado (geralmente preferido)
function useMousePosition() {
  const [position, setPosition] = useState({ x: 0, y: 0 });
  // ... lógica
  return position;
}

// Uso do hook
function App() {
  const position = useMousePosition();
  return <div>Mouse: ({position.x}, {position.y})</div>;
}
```

**Quando usar cada um:**
- **Hooks**: Para lógica reutilizável simples
- **Render Props**: Quando você precisa de flexibilidade máxima na renderização

#### Render Props vs Higher-Order Components

```jsx
// Render Props
function DataFetcher({ url, children }) {
  const [data, setData] = useState(null);
  // ... lógica
  return children({ data });
}

// HOC equivalente
function withData(url) {
  return function(Component) {
    return function(props) {
      const [data, setData] = useState(null);
      // ... lógica
      return <Component {...props} data={data} />;
    };
  };
}
```

**Render Props são geralmente preferidos** porque:
- Mais flexíveis
- Não criam camadas extras de componentes
- Mais fáceis de debugar

### 4.6 Exemplo Completo: Form com Validação

```jsx
// Componente de formulário com render prop
function Form({ initialValues, onSubmit, children }) {
  const [values, setValues] = useState(initialValues || {});
  const [errors, setErrors] = useState({});
  const [touched, setTouched] = useState({});
  
  const handleChange = (name, value) => {
    setValues(prev => ({ ...prev, [name]: value }));
    // Limpar erro quando usuário começa a digitar
    if (errors[name]) {
      setErrors(prev => {
        const newErrors = { ...prev };
        delete newErrors[name];
        return newErrors;
      });
    }
  };
  
  const handleBlur = (name) => {
    setTouched(prev => ({ ...prev, [name]: true }));
  };
  
  const handleSubmit = (e) => {
    e.preventDefault();
    // Validação básica
    const newErrors = {};
    Object.keys(values).forEach(key => {
      if (!values[key]) {
        newErrors[key] = 'Campo obrigatório';
      }
    });
    
    if (Object.keys(newErrors).length === 0) {
      onSubmit(values);
    } else {
      setErrors(newErrors);
      setTouched(
        Object.keys(values).reduce((acc, key) => ({ ...acc, [key]: true }), {})
      );
    }
  };
  
  return children({
    values,
    errors,
    touched,
    handleChange,
    handleBlur,
    handleSubmit
  });
}

// Uso
function LoginForm() {
  return (
    <Form
      initialValues={{ email: '', password: '' }}
      onSubmit={(values) => {
        console.log('Submetendo:', values);
      }}
    >
      {({ values, errors, touched, handleChange, handleBlur, handleSubmit }) => (
        <form onSubmit={handleSubmit}>
          <div>
            <input
              type="email"
              name="email"
              value={values.email}
              onChange={(e) => handleChange('email', e.target.value)}
              onBlur={() => handleBlur('email')}
              placeholder="Email"
            />
            {touched.email && errors.email && (
              <span style={{ color: 'red' }}>{errors.email}</span>
            )}
          </div>
          
          <div>
            <input
              type="password"
              name="password"
              value={values.password}
              onChange={(e) => handleChange('password', e.target.value)}
              onBlur={() => handleBlur('password')}
              placeholder="Senha"
            />
            {touched.password && errors.password && (
              <span style={{ color: 'red' }}>{errors.password}</span>
            )}
          </div>
          
          <button type="submit">Entrar</button>
        </form>
      )}
    </Form>
  );
}
```

---

## 5. Refs

### 5.1 O Que São Refs?

**Refs** (referências) fornecem uma forma de acessar diretamente elementos DOM ou instâncias de componentes React criados no método render.

#### Quando Usar Refs?

Refs devem ser usadas quando você precisa:
1. **Acessar elementos DOM diretamente** (foco, scroll, medições)
2. **Acessar métodos de componentes filhos** (imperativo)
3. **Armazenar valores mutáveis** que não causam re-render

**Importante:** Refs são uma "escape hatch" - use apenas quando necessário. A maioria das coisas pode ser feita de forma declarativa.

### 5.2 useRef Hook

O hook `useRef` retorna um objeto mutável com uma propriedade `.current` que persiste durante todo o ciclo de vida do componente.

#### Sintaxe Básica

```jsx
import { useRef } from 'react';

function MyComponent() {
  const myRef = useRef(initialValue);
  
  // Acessar valor: myRef.current
  // Atualizar valor: myRef.current = newValue
  
  return <div ref={myRef}>Conteúdo</div>;
}
```

**Características importantes:**
- `useRef` retorna o mesmo objeto em cada render
- Mudanças em `ref.current` **não causam re-render**
- `ref.current` persiste entre renders

### 5.3 Refs para Elementos DOM

#### Exemplo 1: Focar Input

```jsx
import { useRef } from 'react';

function SearchBox() {
  const inputRef = useRef(null);
  
  const handleFocus = () => {
    // Acessar elemento DOM diretamente
    inputRef.current.focus();
  };
  
  return (
    <div>
      <input
        ref={inputRef}
        type="text"
        placeholder="Buscar..."
      />
      <button onClick={handleFocus}>Focar Input</button>
    </div>
  );
}
```

#### Exemplo 2: Medir Dimensões

```jsx
import { useRef, useState, useEffect } from 'react';

function MeasurableBox() {
  const boxRef = useRef(null);
  const [dimensions, setDimensions] = useState({ width: 0, height: 0 });
  
  useEffect(() => {
    if (boxRef.current) {
      const rect = boxRef.current.getBoundingClientRect();
      setDimensions({
        width: rect.width,
        height: rect.height
      });
    }
  }, []);
  
  return (
    <div>
      <div
        ref={boxRef}
        style={{ width: '200px', height: '100px', background: 'lightblue' }}
      >
        Caixa
      </div>
      <p>
        Largura: {dimensions.width}px, Altura: {dimensions.height}px
      </p>
    </div>
  );
}
```

#### Exemplo 3: Scroll para Elemento

```jsx
import { useRef } from 'react';

function ScrollableList() {
  const itemsRef = useRef([]);
  
  const scrollToItem = (index) => {
    if (itemsRef.current[index]) {
      itemsRef.current[index].scrollIntoView({
        behavior: 'smooth',
        block: 'center'
      });
    }
  };
  
  const items = ['Item 1', 'Item 2', 'Item 3', 'Item 4', 'Item 5'];
  
  return (
    <div>
      <div>
        {items.map((item, index) => (
          <button key={index} onClick={() => scrollToItem(index)}>
            Ir para {item}
          </button>
        ))}
      </div>
      
      <div style={{ height: '200px', overflow: 'auto' }}>
        {items.map((item, index) => (
          <div
            key={index}
            ref={(el) => (itemsRef.current[index] = el)}
            style={{ height: '100px', margin: '10px', background: 'lightgray' }}
          >
            {item}
          </div>
        ))}
      </div>
    </div>
  );
}
```

### 5.4 Refs para Armazenar Valores Mutáveis

Refs podem armazenar valores que não causam re-render quando mudam:

```jsx
import { useRef, useState } from 'react';

function Timer() {
  const [count, setCount] = useState(0);
  const intervalRef = useRef(null);
  
  const startTimer = () => {
    if (intervalRef.current === null) {
      intervalRef.current = setInterval(() => {
        setCount(prev => prev + 1);
      }, 1000);
    }
  };
  
  const stopTimer = () => {
    if (intervalRef.current !== null) {
      clearInterval(intervalRef.current);
      intervalRef.current = null;
    }
  };
  
  return (
    <div>
      <p>Contador: {count}</p>
      <button onClick={startTimer}>Iniciar</button>
      <button onClick={stopTimer}>Parar</button>
    </div>
  );
}
```

**Por que usar ref aqui?**
- `intervalRef.current` armazena o ID do intervalo
- Mudanças em `intervalRef.current` não causam re-render
- O valor persiste entre renders
- Mais eficiente que usar state

### 5.5 Refs para Componentes (ImperativeHandle)

Para acessar métodos de componentes filhos, use `useImperativeHandle`:

```jsx
import { forwardRef, useImperativeHandle, useRef } from 'react';

// Componente filho que expõe métodos
const Input = forwardRef(function Input(props, ref) {
  const inputRef = useRef(null);
  
  // Expõe métodos para o componente pai
  useImperativeHandle(ref, () => ({
    focus: () => {
      inputRef.current.focus();
    },
    clear: () => {
      inputRef.current.value = '';
    },
    getValue: () => {
      return inputRef.current.value;
    }
  }));
  
  return <input ref={inputRef} {...props} />;
});

// Componente pai que usa os métodos
function Form() {
  const inputRef = useRef(null);
  
  const handleFocus = () => {
    inputRef.current.focus();
  };
  
  const handleClear = () => {
    inputRef.current.clear();
  };
  
  const handleGetValue = () => {
    console.log('Valor:', inputRef.current.getValue());
  };
  
  return (
    <div>
      <Input ref={inputRef} placeholder="Digite algo..." />
      <button onClick={handleFocus}>Focar</button>
      <button onClick={handleClear}>Limpar</button>
      <button onClick={handleGetValue}>Obter Valor</button>
    </div>
  );
}
```

### 5.6 Callback Refs

Além de `useRef`, você pode usar **callback refs**:

```jsx
function MeasureExample() {
  const [height, setHeight] = useState(0);
  
  // Callback ref: função que recebe o elemento DOM
  const measuredRef = useCallback((node) => {
    if (node !== null) {
      setHeight(node.getBoundingClientRect().height);
    }
  }, []);
  
  return (
    <div>
      <div ref={measuredRef}>
        <h1>Conteúdo que será medido</h1>
        <p>Algum texto aqui...</p>
      </div>
      <p>A altura acima é {Math.round(height)}px</p>
    </div>
  );
}
```

**Quando usar callback refs:**
- Quando você precisa executar código quando o ref é atribuído
- Quando você precisa medir elementos dinamicamente
- Quando você precisa limpar recursos quando o elemento é desmontado

### 5.7 Exemplo Completo: Formulário com Refs

```jsx
import { useRef, useState } from 'react';

function ContactForm() {
  const nameRef = useRef(null);
  const emailRef = useRef(null);
  const messageRef = useRef(null);
  const [errors, setErrors] = useState({});
  
  const validate = () => {
    const newErrors = {};
    
    if (!nameRef.current.value.trim()) {
      newErrors.name = 'Nome é obrigatório';
      nameRef.current.focus();
    }
    
    if (!emailRef.current.value.includes('@')) {
      newErrors.email = 'Email inválido';
      if (!newErrors.name) emailRef.current.focus();
    }
    
    if (!messageRef.current.value.trim()) {
      newErrors.message = 'Mensagem é obrigatória';
      if (!newErrors.name && !newErrors.email) messageRef.current.focus();
    }
    
    setErrors(newErrors);
    return Object.keys(newErrors).length === 0;
  };
  
  const handleSubmit = (e) => {
    e.preventDefault();
    
    if (validate()) {
      const formData = {
        name: nameRef.current.value,
        email: emailRef.current.value,
        message: messageRef.current.value
      };
      
      console.log('Enviando:', formData);
      // Enviar dados...
    }
  };
  
  return (
    <form onSubmit={handleSubmit}>
      <div>
        <label>Nome:</label>
        <input ref={nameRef} type="text" />
        {errors.name && <span style={{ color: 'red' }}>{errors.name}</span>}
      </div>
      
      <div>
        <label>Email:</label>
        <input ref={emailRef} type="email" />
        {errors.email && <span style={{ color: 'red' }}>{errors.email}</span>}
      </div>
      
      <div>
        <label>Mensagem:</label>
        <textarea ref={messageRef} rows="4" />
        {errors.message && <span style={{ color: 'red' }}>{errors.message}</span>}
      </div>
      
      <button type="submit">Enviar</button>
    </form>
  );
}
```

**Nota:** Este exemplo usa refs, mas em muitos casos, **controlled components** (com state) são preferidos. Use refs quando precisar de acesso imperativo ao DOM.

### 5.8 Refs vs State

| Aspecto | Refs | State |
|---------|------|-------|
| **Causa re-render?** | Não | Sim |
| **Mutável?** | Sim (`.current`) | Não (imutável) |
| **Quando usar** | Acesso DOM, valores que não precisam de render | Dados que afetam UI |
| **Persistência** | Entre renders | Entre renders |
| **Atualização** | `ref.current = value` | `setState(value)` |

**Regra geral:**
- Use **state** para dados que afetam o que é renderizado
- Use **refs** para valores que não precisam causar re-render ou para acesso DOM

---

## 📝 Resumo Parcial da Aula

Até agora, cobrimos:

### ✅ Rendering (Renderização)
- Abordagem declarativa vs imperativa
- Processo de renderização no React
- Virtual DOM e sua importância
- Reconciliation (reconciliação)

### ✅ Component Life Cycle (Ciclo de Vida)
- Fase de Mounting, Updating, Unmounting
- Lifecycle methods vs Hooks modernos
- Uso correto de `useEffect`

### ✅ Lists and Keys
- Por que keys são essenciais
- Como escolher boas keys
- Problemas comuns e soluções
- Renderização eficiente de listas

### ✅ Render Props
- Conceito de render props
- Quando usar render props
- Padrões comuns
- Comparação com outros padrões

### ✅ Refs
- O que são refs e quando usar
- `useRef` hook
- Refs para elementos DOM
- Refs para componentes
- Callback refs

---

## 6. Events (Eventos)

### 6.1 Sistema de Eventos do React

Manipular eventos com elementos React é muito similar a manipular eventos em elementos DOM, mas há algumas diferenças importantes de sintaxe e comportamento.

#### Diferenças Principais

1. **Nomenclatura em camelCase**: Em vez de `onclick`, React usa `onClick`
2. **Função como handler**: Em vez de string, você passa uma função
3. **SyntheticEvent**: React usa um wrapper chamado SyntheticEvent
4. **Event pooling**: Em versões antigas, eventos eram reutilizados (não mais no React 17+)

#### Comparação: HTML vs React

```html
<!-- HTML tradicional -->
<button onclick="handleClick()">Clique aqui</button>
```

```jsx
// React
<button onClick={handleClick}>Clique aqui</button>
```

### 6.2 Sintaxe de Event Handlers

#### Handler Inline

```jsx
function Button() {
  return (
    <button onClick={() => console.log('Clicado!')}>
      Clique aqui
    </button>
  );
}
```

#### Handler como Função Nomeada

```jsx
function Button() {
  const handleClick = () => {
    console.log('Clicado!');
  };
  
  return <button onClick={handleClick}>Clique aqui</button>;
}
```

#### Handler com Parâmetros

```jsx
function TodoList({ todos, onToggle }) {
  return (
    <ul>
      {todos.map(todo => (
        <li key={todo.id}>
          <button onClick={() => onToggle(todo.id)}>
            {todo.completed ? 'Desmarcar' : 'Marcar'}
          </button>
          {todo.text}
        </li>
      ))}
    </ul>
  );
}
```

### 6.3 SyntheticEvent

React envolve eventos nativos do navegador em um objeto chamado **SyntheticEvent**. Isso garante que eventos funcionem de forma consistente em todos os navegadores.

#### Propriedades do SyntheticEvent

```jsx
function EventExample() {
  const handleClick = (e) => {
    // e é um SyntheticEvent
    console.log(e.type);           // "click"
    console.log(e.target);          // Elemento que disparou o evento
    console.log(e.currentTarget);   // Elemento onde o handler está anexado
    console.log(e.nativeEvent);     // Evento nativo do navegador
  };
  
  return <button onClick={handleClick}>Clique</button>;
}
```

#### Propriedades Comuns

```jsx
function FormExample() {
  const handleSubmit = (e) => {
    e.preventDefault();        // Previne comportamento padrão
    e.stopPropagation();       // Para propagação do evento
    console.log(e.target);     // Elemento que disparou
    console.log(e.type);       // Tipo do evento
  };
  
  return (
    <form onSubmit={handleSubmit}>
      <input type="text" />
      <button type="submit">Enviar</button>
    </form>
  );
}
```

### 6.4 Tipos de Eventos Comuns

#### Eventos de Mouse

```jsx
function MouseEvents() {
  const handleClick = (e) => console.log('Click', e.clientX, e.clientY);
  const handleDoubleClick = () => console.log('Double click');
  const handleMouseEnter = () => console.log('Mouse entrou');
  const handleMouseLeave = () => console.log('Mouse saiu');
  const handleMouseOver = () => console.log('Mouse sobre');
  const handleMouseMove = (e) => console.log('Mouse moveu', e.clientX, e.clientY);
  
  return (
    <div
      onClick={handleClick}
      onDoubleClick={handleDoubleClick}
      onMouseEnter={handleMouseEnter}
      onMouseLeave={handleMouseLeave}
      onMouseOver={handleMouseOver}
      onMouseMove={handleMouseMove}
      style={{ padding: '20px', background: 'lightblue' }}
    >
      Passe o mouse aqui
    </div>
  );
}
```

#### Eventos de Teclado

```jsx
function KeyboardEvents() {
  const handleKeyDown = (e) => {
    console.log('Tecla pressionada:', e.key);
    if (e.key === 'Enter') {
      console.log('Enter pressionado!');
    }
    if (e.ctrlKey && e.key === 's') {
      e.preventDefault();
      console.log('Salvar (Ctrl+S)');
    }
  };
  
  const handleKeyUp = (e) => {
    console.log('Tecla solta:', e.key);
  };
  
  return (
    <input
      type="text"
      onKeyDown={handleKeyDown}
      onKeyUp={handleKeyUp}
      placeholder="Digite algo..."
    />
  );
}
```

#### Eventos de Formulário

```jsx
function FormEvents() {
  const [value, setValue] = useState('');
  
  const handleChange = (e) => {
    setValue(e.target.value);
    console.log('Valor mudou:', e.target.value);
  };
  
  const handleFocus = () => {
    console.log('Input recebeu foco');
  };
  
  const handleBlur = () => {
    console.log('Input perdeu foco');
  };
  
  const handleSubmit = (e) => {
    e.preventDefault();
    console.log('Formulário submetido:', value);
  };
  
  return (
    <form onSubmit={handleSubmit}>
      <input
        type="text"
        value={value}
        onChange={handleChange}
        onFocus={handleFocus}
        onBlur={handleBlur}
        placeholder="Digite algo..."
      />
      <button type="submit">Enviar</button>
    </form>
  );
}
```

### 6.5 Prevenção de Comportamento Padrão

#### preventDefault()

Previne o comportamento padrão do navegador:

```jsx
function LinkExample() {
  const handleClick = (e) => {
    e.preventDefault();
    console.log('Link clicado, mas navegação prevenida');
    // Fazer algo customizado em vez de navegar
  };
  
  return (
    <a href="/page" onClick={handleClick}>
      Clique aqui (não navega)
    </a>
  );
}
```

#### stopPropagation()

Para a propagação do evento (bubbling):

```jsx
function PropagationExample() {
  const handleParentClick = () => {
    console.log('Parent clicado');
  };
  
  const handleChildClick = (e) => {
    e.stopPropagation(); // Para aqui, não propaga para o parent
    console.log('Child clicado');
  };
  
  return (
    <div onClick={handleParentClick} style={{ padding: '20px', background: 'lightblue' }}>
      <p>Parent</p>
      <button onClick={handleChildClick}>
        Child (não propaga)
      </button>
    </div>
  );
}
```

#### stopPropagation() + preventDefault()

```jsx
function CombinedExample() {
  const handleClick = (e) => {
    e.preventDefault();
    e.stopPropagation();
    console.log('Comportamento padrão e propagação prevenidos');
  };
  
  return (
    <form>
      <button type="submit" onClick={handleClick}>
        Enviar (não submete formulário)
      </button>
    </form>
  );
}
```

### 6.6 Event Bubbling e Capturing

React usa **event delegation** - todos os eventos são delegados ao elemento raiz e depois propagam (bubbling).

```jsx
function BubblingExample() {
  const handleGrandparent = () => console.log('Grandparent');
  const handleParent = () => console.log('Parent');
  const handleChild = () => console.log('Child');
  
  return (
    <div onClick={handleGrandparent} style={{ padding: '30px', background: 'lightgray' }}>
      <div onClick={handleParent} style={{ padding: '20px', background: 'lightblue' }}>
        <div onClick={handleChild} style={{ padding: '10px', background: 'lightgreen' }}>
          Clique aqui
        </div>
      </div>
    </div>
  );
}

// Ao clicar no elemento mais interno:
// Output: "Child", "Parent", "Grandparent"
// (bubbling: do mais interno para o mais externo)
```

### 6.7 Passando Argumentos para Event Handlers

#### Usando Arrow Functions

```jsx
function TodoList({ todos, onDelete }) {
  return (
    <ul>
      {todos.map(todo => (
        <li key={todo.id}>
          {todo.text}
          <button onClick={() => onDelete(todo.id)}>
            Deletar
          </button>
        </li>
      ))}
    </ul>
  );
}
```

#### Usando bind

```jsx
function TodoList({ todos, onDelete }) {
  return (
    <ul>
      {todos.map(todo => (
        <li key={todo.id}>
          {todo.text}
          <button onClick={onDelete.bind(null, todo.id)}>
            Deletar
          </button>
        </li>
      ))}
    </ul>
  );
}
```

#### Handler que Retorna Função

```jsx
function TodoList({ todos, onDelete }) {
  const handleDelete = (id) => () => {
    onDelete(id);
  };
  
  return (
    <ul>
      {todos.map(todo => (
        <li key={todo.id}>
          {todo.text}
          <button onClick={handleDelete(todo.id)}>
            Deletar
          </button>
        </li>
      ))}
    </ul>
  );
}
```

### 6.8 Eventos Customizados

Você pode criar eventos customizados usando `CustomEvent`:

```jsx
function CustomEventExample() {
  const handleCustomEvent = (e) => {
    console.log('Evento customizado:', e.detail);
  };
  
  useEffect(() => {
    const event = new CustomEvent('myCustomEvent', {
      detail: { message: 'Hello from custom event!' }
    });
    
    window.addEventListener('myCustomEvent', handleCustomEvent);
    
    // Disparar evento após 2 segundos
    setTimeout(() => {
      window.dispatchEvent(event);
    }, 2000);
    
    return () => {
      window.removeEventListener('myCustomEvent', handleCustomEvent);
    };
  }, []);
  
  return <div>Esperando evento customizado...</div>;
}
```

### 6.9 Exemplo Completo: Formulário com Validação

```jsx
import { useState } from 'react';

function ContactForm() {
  const [formData, setFormData] = useState({
    name: '',
    email: '',
    message: ''
  });
  const [errors, setErrors] = useState({});
  const [touched, setTouched] = useState({});
  
  const handleChange = (e) => {
    const { name, value } = e.target;
    setFormData(prev => ({ ...prev, [name]: value }));
    
    // Limpar erro quando usuário começa a digitar
    if (errors[name]) {
      setErrors(prev => {
        const newErrors = { ...prev };
        delete newErrors[name];
        return newErrors;
      });
    }
  };
  
  const handleBlur = (e) => {
    const { name } = e.target;
    setTouched(prev => ({ ...prev, [name]: true }));
    validateField(name, formData[name]);
  };
  
  const validateField = (name, value) => {
    const newErrors = { ...errors };
    
    if (name === 'name' && !value.trim()) {
      newErrors.name = 'Nome é obrigatório';
    }
    
    if (name === 'email') {
      if (!value.trim()) {
        newErrors.email = 'Email é obrigatório';
      } else if (!value.includes('@')) {
        newErrors.email = 'Email inválido';
      }
    }
    
    if (name === 'message' && !value.trim()) {
      newErrors.message = 'Mensagem é obrigatória';
    }
    
    setErrors(newErrors);
  };
  
  const handleSubmit = (e) => {
    e.preventDefault();
    
    // Validar todos os campos
    Object.keys(formData).forEach(key => {
      validateField(key, formData[key]);
      setTouched(prev => ({ ...prev, [key]: true }));
    });
    
    // Se não há erros, submeter
    if (Object.keys(errors).length === 0) {
      console.log('Enviando:', formData);
      // Enviar dados...
    }
  };
  
  return (
    <form onSubmit={handleSubmit}>
      <div>
        <label>Nome:</label>
        <input
          type="text"
          name="name"
          value={formData.name}
          onChange={handleChange}
          onBlur={handleBlur}
        />
        {touched.name && errors.name && (
          <span style={{ color: 'red' }}>{errors.name}</span>
        )}
      </div>
      
      <div>
        <label>Email:</label>
        <input
          type="email"
          name="email"
          value={formData.email}
          onChange={handleChange}
          onBlur={handleBlur}
        />
        {touched.email && errors.email && (
          <span style={{ color: 'red' }}>{errors.email}</span>
        )}
      </div>
      
      <div>
        <label>Mensagem:</label>
        <textarea
          name="message"
          value={formData.message}
          onChange={handleChange}
          onBlur={handleBlur}
          rows="4"
        />
        {touched.message && errors.message && (
          <span style={{ color: 'red' }}>{errors.message}</span>
        )}
      </div>
      
      <button type="submit">Enviar</button>
    </form>
  );
}
```

### 6.10 Boas Práticas com Eventos

#### 1. Não Criar Handlers Dentro do Render (sem otimização)

```jsx
// ❌ ERRADO: Nova função criada a cada render
function Component({ items }) {
  return (
    <ul>
      {items.map(item => (
        <li key={item.id}>
          <button onClick={() => handleClick(item.id)}>
            {item.name}
          </button>
        </li>
      ))}
    </ul>
  );
}

// ✅ CORRETO: Handler estável (se não precisar de otimização, arrow function é OK)
function Component({ items, onItemClick }) {
  return (
    <ul>
      {items.map(item => (
        <li key={item.id}>
          <button onClick={() => onItemClick(item.id)}>
            {item.name}
          </button>
        </li>
      ))}
    </ul>
  );
}
```

#### 2. Usar useCallback para Handlers Complexos

```jsx
import { useCallback } from 'react';

function Component({ items, onItemClick }) {
  const handleClick = useCallback((id) => {
    // Lógica complexa aqui
    onItemClick(id);
  }, [onItemClick]);
  
  return (
    <ul>
      {items.map(item => (
        <li key={item.id}>
          <button onClick={() => handleClick(item.id)}>
            {item.name}
          </button>
        </li>
      ))}
    </ul>
  );
}
```

#### 3. Sempre Validar Dados de Eventos

```jsx
function SafeComponent() {
  const handleSubmit = (e) => {
    e.preventDefault();
    
    const formData = new FormData(e.target);
    const email = formData.get('email');
    
    // Validar antes de usar
    if (!email || !email.includes('@')) {
      alert('Email inválido');
      return;
    }
    
    // Usar dados validados
    console.log('Email válido:', email);
  };
  
  return (
    <form onSubmit={handleSubmit}>
      <input type="email" name="email" />
      <button type="submit">Enviar</button>
    </form>
  );
}
```

---

## 7. Higher-Order Components (HOCs)

### 7.1 O Que São HOCs?

Um **Higher-Order Component (HOC)** é uma função que recebe um componente e retorna um novo componente com funcionalidades adicionais. HOCs são um padrão avançado de reutilização de lógica em React.

**Definição formal:**
> Um Higher-Order Component é uma função que recebe um componente e retorna um novo componente.

#### Estrutura Básica

```jsx
// HOC básico
function withSomething(Component) {
  return function EnhancedComponent(props) {
    // Lógica adicional aqui
    return <Component {...props} />;
  };
}

// Uso
const EnhancedButton = withSomething(Button);
```

### 7.2 Por Que Usar HOCs?

HOCs permitem:
1. **Reutilizar lógica** entre componentes
2. **Adicionar funcionalidades** sem modificar o componente original
3. **Separar preocupações** (lógica vs apresentação)
4. **Compartilhar código** de forma composável

**Nota importante:** Com a introdução de Hooks, HOCs são menos comuns. Hooks geralmente são preferidos para reutilização de lógica. Mas entender HOCs é importante para código legado e alguns casos específicos.

### 7.3 Criando um HOC Simples

#### Exemplo 1: HOC que Adiciona Loading

```jsx
// HOC que adiciona funcionalidade de loading
function withLoading(Component) {
  return function WithLoadingComponent({ isLoading, ...props }) {
    if (isLoading) {
      return <div>Carregando...</div>;
    }
    return <Component {...props} />;
  };
}

// Componente original
function UserList({ users }) {
  return (
    <ul>
      {users.map(user => (
        <li key={user.id}>{user.name}</li>
      ))}
    </ul>
  );
}

// Componente aprimorado com HOC
const UserListWithLoading = withLoading(UserList);

// Uso
function App() {
  const [isLoading, setIsLoading] = useState(true);
  const [users, setUsers] = useState([]);
  
  return (
    <UserListWithLoading isLoading={isLoading} users={users} />
  );
}
```

#### Exemplo 2: HOC que Adiciona Autenticação

```jsx
// HOC que verifica autenticação
function withAuth(Component) {
  return function AuthenticatedComponent(props) {
    const [isAuthenticated, setIsAuthenticated] = useState(false);
    const [isChecking, setIsChecking] = useState(true);
    
    useEffect(() => {
      // Verificar autenticação
      checkAuth().then(auth => {
        setIsAuthenticated(auth);
        setIsChecking(false);
      });
    }, []);
    
    if (isChecking) {
      return <div>Verificando autenticação...</div>;
    }
    
    if (!isAuthenticated) {
      return <div>Por favor, faça login</div>;
    }
    
    return <Component {...props} />;
  };
}

// Componente protegido
function Dashboard() {
  return <div>Conteúdo do Dashboard</div>;
}

// Componente com autenticação
const ProtectedDashboard = withAuth(Dashboard);
```

### 7.4 HOCs com Props Adicionais

HOCs podem adicionar props ao componente:

```jsx
// HOC que adiciona dados de usuário
function withUserData(Component) {
  return function WithUserDataComponent(props) {
    const [user, setUser] = useState(null);
    
    useEffect(() => {
      fetchUser().then(setUser);
    }, []);
    
    return <Component {...props} user={user} />;
  };
}

// Componente que recebe user como prop
function Profile({ user }) {
  if (!user) return <div>Carregando usuário...</div>;
  
  return (
    <div>
      <h1>{user.name}</h1>
      <p>{user.email}</p>
    </div>
  );
}

// Componente com dados de usuário
const ProfileWithUser = withUserData(Profile);
```

### 7.5 HOCs que Modificam Props

```jsx
// HOC que transforma props
function withUpperCase(Component) {
  return function WithUpperCaseComponent({ text, ...props }) {
    const upperText = text ? text.toUpperCase() : '';
    return <Component {...props} text={upperText} />;
  };
}

// Componente original
function DisplayText({ text }) {
  return <p>{text}</p>;
}

// Componente com texto em maiúsculas
const DisplayTextUpper = withUpperCase(DisplayText);

// Uso
<DisplayTextUpper text="hello world" />
// Renderiza: <p>HELLO WORLD</p>
```

### 7.6 HOCs com Nomes de Exibição

Para facilitar debugging, é bom dar nomes aos componentes retornados:

```jsx
function withLoading(Component) {
  function WithLoadingComponent({ isLoading, ...props }) {
    if (isLoading) {
      return <div>Carregando...</div>;
    }
    return <Component {...props} />;
  }
  
  // Dar nome ao componente para debugging
  WithLoadingComponent.displayName = `withLoading(${Component.displayName || Component.name || 'Component'})`;
  
  return WithLoadingComponent;
}
```

### 7.7 HOCs Compostos (Composição de HOCs)

Você pode compor múltiplos HOCs:

```jsx
// HOC 1: Loading
function withLoading(Component) {
  return function WithLoading({ isLoading, ...props }) {
    if (isLoading) return <div>Carregando...</div>;
    return <Component {...props} />;
  };
}

// HOC 2: Error handling
function withError(Component) {
  return function WithError({ error, ...props }) {
    if (error) return <div>Erro: {error.message}</div>;
    return <Component {...props} />;
  };
}

// HOC 3: User data
function withUser(Component) {
  return function WithUser(props) {
    const [user, setUser] = useState(null);
    useEffect(() => {
      fetchUser().then(setUser);
    }, []);
    return <Component {...props} user={user} />;
  };
}

// Composição de HOCs
const EnhancedComponent = withLoading(
  withError(
    withUser(UserProfile)
  )
);

// Ou usando uma função auxiliar
function compose(...hocs) {
  return (Component) => hocs.reduceRight((acc, hoc) => hoc(acc), Component);
}

const EnhancedComponent = compose(
  withLoading,
  withError,
  withUser
)(UserProfile);
```

### 7.8 HOCs vs Hooks

#### Com HOC (Antigo)

```jsx
// HOC
function withWindowSize(Component) {
  return function WithWindowSize(props) {
    const [size, setSize] = useState({ width: window.innerWidth, height: window.innerHeight });
    
    useEffect(() => {
      const handleResize = () => {
        setSize({ width: window.innerWidth, height: window.innerHeight });
      };
      window.addEventListener('resize', handleResize);
      return () => window.removeEventListener('resize', handleResize);
    }, []);
    
    return <Component {...props} windowSize={size} />;
  };
}

// Uso
const ComponentWithSize = withWindowSize(MyComponent);
```

#### Com Hook (Moderno - Preferido)

```jsx
// Hook customizado
function useWindowSize() {
  const [size, setSize] = useState({ width: window.innerWidth, height: window.innerHeight });
  
  useEffect(() => {
    const handleResize = () => {
      setSize({ width: window.innerWidth, height: window.innerHeight });
    };
    window.addEventListener('resize', handleResize);
    return () => window.removeEventListener('resize', handleResize);
  }, []);
  
  return size;
}

// Uso
function MyComponent() {
  const windowSize = useWindowSize();
  return <div>Tamanho: {windowSize.width}x{windowSize.height}</div>;
}
```

**Por que hooks são preferidos:**
- Mais simples e diretos
- Não criam camadas extras de componentes
- Mais fáceis de testar
- Melhor para debugging
- Mais flexíveis

### 7.9 Quando Usar HOCs?

Use HOCs quando:
- Você precisa adicionar funcionalidade a múltiplos componentes
- Você está trabalhando com código legado que usa HOCs
- Você precisa de um padrão específico que HOCs facilitam

**Mas considere usar Hooks primeiro:**
- Hooks são geralmente mais simples
- Hooks são mais modernos e recomendados
- Hooks são mais fáceis de entender e manter

### 7.10 Exemplo Completo: HOC de Logging

```jsx
// HOC que adiciona logging
function withLogging(Component, componentName) {
  return function WithLoggingComponent(props) {
    useEffect(() => {
      console.log(`${componentName} montado`);
      return () => {
        console.log(`${componentName} desmontado`);
      };
    }, []);
    
    useEffect(() => {
      console.log(`${componentName} props atualizadas:`, props);
    });
    
    const handleClick = (...args) => {
      console.log(`${componentName} clicado:`, args);
      if (props.onClick) {
        props.onClick(...args);
      }
    };
    
    return <Component {...props} onClick={handleClick} />;
  };
}

// Componente original
function Button({ children, onClick }) {
  return <button onClick={onClick}>{children}</button>;
}

// Componente com logging
const LoggedButton = withLogging(Button, 'Button');

// Uso
function App() {
  return (
    <LoggedButton onClick={() => console.log('Button clicado!')}>
      Clique aqui
    </LoggedButton>
  );
}
```

### 7.11 Padrões Comuns de HOCs

#### 1. HOC de Autenticação

```jsx
function withAuth(Component) {
  return function AuthenticatedComponent(props) {
    const [user, setUser] = useState(null);
    const [loading, setLoading] = useState(true);
    
    useEffect(() => {
      checkAuth().then(userData => {
        setUser(userData);
        setLoading(false);
      });
    }, []);
    
    if (loading) return <div>Verificando...</div>;
    if (!user) return <div>Faça login</div>;
    
    return <Component {...props} user={user} />;
  };
}
```

#### 2. HOC de Dados (Data Fetching)

```jsx
function withData(url) {
  return function(Component) {
    return function WithDataComponent(props) {
      const [data, setData] = useState(null);
      const [loading, setLoading] = useState(true);
      const [error, setError] = useState(null);
      
      useEffect(() => {
        fetch(url)
          .then(res => res.json())
          .then(data => {
            setData(data);
            setLoading(false);
          })
          .catch(err => {
            setError(err);
            setLoading(false);
          });
      }, [url]);
      
      return (
        <Component
          {...props}
          data={data}
          loading={loading}
          error={error}
        />
      );
    };
  };
}

// Uso
const UserListWithData = withData('/api/users')(UserList);
```

#### 3. HOC de Estilização

```jsx
function withStyles(styles) {
  return function(Component) {
    return function StyledComponent(props) {
      return (
        <div style={styles.container}>
          <Component {...props} style={styles.content} />
        </div>
      );
    };
  };
}

// Uso
const StyledCard = withStyles({
  container: { padding: '20px', border: '1px solid #ccc' },
  content: { background: '#f5f5f5' }
})(Card);
```

### 7.12 Problemas Comuns com HOCs

#### 1. Props Colisão

```jsx
// ❌ PROBLEMA: HOC pode sobrescrever props
function withUser(Component) {
  return function WithUser({ user, ...props }) {
    const fetchedUser = useFetchUser();
    return <Component {...props} user={fetchedUser} />;
  };
}

// Se o componente já recebe 'user' como prop, há conflito
```

**Solução:** Usar nomes específicos ou mesclar props:

```jsx
function withUser(Component) {
  return function WithUser(props) {
    const fetchedUser = useFetchUser();
    return <Component {...props} fetchedUser={fetchedUser} />;
  };
}
```

#### 2. Refs Não São Passadas

```jsx
// ❌ PROBLEMA: Refs não são passadas automaticamente
function withSomething(Component) {
  return function WithSomething(props) {
    return <Component {...props} />; // ref não é passada
  };
}
```

**Solução:** Usar `forwardRef`:

```jsx
import { forwardRef } from 'react';

function withSomething(Component) {
  return forwardRef(function WithSomething(props, ref) {
    return <Component {...props} ref={ref} />;
  });
}
```

### 7.13 Resumo: HOCs vs Hooks

| Aspecto | HOCs | Hooks |
|---------|------|-------|
| **Sintaxe** | Função que retorna componente | Função que retorna valor |
| **Complexidade** | Mais complexo | Mais simples |
| **Debugging** | Mais difícil (camadas extras) | Mais fácil |
| **Composição** | Composição de funções | Composição natural |
| **Performance** | Pode adicionar camadas | Sem overhead |
| **Recomendação** | Código legado | Novo código |

**Regra geral:** Use Hooks para novo código. Use HOCs apenas se necessário para código legado ou casos específicos.

---

## 📝 Resumo Completo da Aula

Cobrimos todos os tópicos fundamentais sobre Rendering e Conceitos Avançados do React:

### ✅ Rendering (Renderização)
- Abordagem declarativa vs imperativa
- Processo de renderização no React
- Virtual DOM e sua importância
- Reconciliation (reconciliação)
- Quando React renderiza componentes

### ✅ Component Life Cycle (Ciclo de Vida)
- Fase de Mounting (montagem)
- Fase de Updating (atualização)
- Fase de Unmounting (desmontagem)
- Lifecycle methods vs Hooks modernos
- Uso correto de `useEffect`

### ✅ Lists and Keys
- Por que keys são essenciais
- Como escolher boas keys
- Problemas comuns e soluções
- Renderização eficiente de listas

### ✅ Render Props
- Conceito de render props
- Quando usar render props
- Padrões comuns
- Comparação com outros padrões

### ✅ Refs
- O que são refs e quando usar
- `useRef` hook
- Refs para elementos DOM
- Refs para componentes
- Callback refs

### ✅ Events
- Sistema de eventos do React
- SyntheticEvent
- Manipulação de eventos
- Event handlers
- Prevenção de comportamento padrão
- Event bubbling e capturing

### ✅ Higher-Order Components (HOCs)
- Conceito de HOCs
- Como criar HOCs
- Quando usar HOCs
- HOCs vs Hooks
- Padrões comuns

---

## 🎯 Próximos Passos

Agora que você entendeu esses conceitos fundamentais:

1. **Pratique** cada conceito isoladamente
2. **Experimente** combinar diferentes padrões
3. **Leia a Aula Simplificada** para reforçar com analogias
4. **Complete os Exercícios** para consolidar o conhecimento
5. **Estude Performance e Boas Práticas** para escrever código profissional

---

**Parabéns por completar a Aula 3! 🎉**

