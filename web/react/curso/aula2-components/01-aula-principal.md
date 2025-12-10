# Aula 2: Components - Os Blocos de Construção do React

## Introdução: O Que São Components?

**Components** (Componentes) são os blocos fundamentais de construção de aplicações React. Eles são unidades independentes e reutilizáveis de código que encapsulam tanto a lógica quanto a apresentação de uma parte da interface do usuário.

Pense em components como peças de Lego: cada peça tem uma função específica, e você combina várias peças para construir algo maior e mais complexo. Da mesma forma, em React, você constrói interfaces complexas combinando componentes menores e mais simples.

### Por Que Components São Importantes?

1. **Reutilização**: Escreva uma vez, use em qualquer lugar
2. **Manutenibilidade**: Código organizado e fácil de manter
3. **Testabilidade**: Componentes isolados são mais fáceis de testar
4. **Colaboração**: Diferentes desenvolvedores podem trabalhar em componentes diferentes
5. **Abstração**: Escondem complexidade, expondo apenas o necessário

---

## 1. Components: Conceito Fundamental

### 1.1 Definição Técnica

Um **component** em React é uma função JavaScript (ou classe, em versões antigas) que:
- Recebe dados opcionais através de **props** (propriedades)
- Retorna **JSX** (JavaScript XML) que descreve o que deve aparecer na tela
- Pode ter **estado interno** (state) que pode mudar ao longo do tempo
- Pode responder a eventos do usuário

### 1.2 Características dos Components

#### Encapsulamento

Um componente encapsula sua própria lógica e apresentação. Isso significa que tudo relacionado a uma funcionalidade específica está contido dentro do componente.

**Exemplo:**
```jsx
function Button({ text, onClick }) {
  const handleClick = () => {
    console.log('Botão clicado!');
    if (onClick) {
      onClick();
    }
  };

  return (
    <button onClick={handleClick}>
      {text}
    </button>
  );
}
```

Este componente `Button` encapsula:
- A lógica de clique
- A apresentação visual (o elemento `<button>`)
- O comportamento (o que acontece ao clicar)

#### Composição

Components podem ser compostos, ou seja, um componente pode conter outros componentes. Isso permite construir interfaces complexas a partir de componentes simples.

**Exemplo:**
```jsx
function Card({ title, children }) {
  return (
    <div className="card">
      <h2>{title}</h2>
      {children}
    </div>
  );
}

function Button({ text }) {
  return <button>{text}</button>;
}

function ProductCard({ productName, price }) {
  return (
    <Card title={productName}>
      <p>Preço: R$ {price}</p>
      <Button text="Comprar" />
    </Card>
  );
}
```

Aqui, `ProductCard` é composto por `Card` e `Button`, que são componentes menores.

#### Isolamento

Cada componente é isolado, o que significa que mudanças em um componente não afetam outros componentes diretamente (a menos que sejam intencionais através de props ou estado compartilhado).

### 1.3 Hierarquia de Components

Components formam uma árvore hierárquica, onde há componentes pais e filhos:

```
App (componente raiz)
├── Header
│   ├── Logo
│   └── Navigation
│       ├── NavItem
│       └── NavItem
├── Main
│   ├── ProductList
│   │   ├── ProductCard
│   │   │   ├── ProductImage
│   │   │   ├── ProductTitle
│   │   │   └── ProductButton
│   │   └── ProductCard
│   └── Sidebar
└── Footer
```

Cada componente tem uma responsabilidade específica e pode conter outros componentes.

---

## 2. Functional Components

### 2.1 O Que São Functional Components?

**Functional Components** (Componentes Funcionais) são componentes React definidos como funções JavaScript. Eles são a forma moderna e recomendada de criar componentes em React.

### 2.2 Sintaxe Básica

A forma mais simples de um functional component é uma função que retorna JSX:

```jsx
function Welcome() {
  return <h1>Bem-vindo ao React!</h1>;
}
```

Ou usando arrow function:

```jsx
const Welcome = () => {
  return <h1>Bem-vindo ao React!</h1>;
};
```

### 2.3 Functional Components com Props

Props são dados passados de um componente pai para um filho. Em functional components, props são recebidas como parâmetro da função:

```jsx
function Greeting({ name, age }) {
  return (
    <div>
      <h1>Olá, {name}!</h1>
      <p>Você tem {age} anos.</p>
    </div>
  );
}

// Uso:
<Greeting name="João" age={25} />
```

**Observações importantes:**
- Props são **read-only** (somente leitura). Um componente não deve modificar suas props diretamente.
- Props podem ser qualquer tipo de dado: strings, números, objetos, arrays, funções, até mesmo outros componentes.

### 2.4 Functional Components com Estado

Antes do React Hooks (introduzidos no React 16.8), apenas class components podiam ter estado. Agora, functional components podem usar hooks como `useState` para gerenciar estado:

```jsx
import { useState } from 'react';

function Counter() {
  const [count, setCount] = useState(0);

  return (
    <div>
      <p>Você clicou {count} vezes</p>
      <button onClick={() => setCount(count + 1)}>
        Clique aqui
      </button>
    </div>
  );
}
```

**Como funciona:**
- `useState(0)` inicializa o estado com valor `0`
- Retorna um array: `[valorAtual, funçãoParaAtualizar]`
- Quando `setCount` é chamado, React re-renderiza o componente com o novo valor

### 2.5 Functional Components com Efeitos Colaterais

Para efeitos colaterais (como chamadas de API, subscriptions, manipulação do DOM), usamos o hook `useEffect`:

```jsx
import { useState, useEffect } from 'react';

function UserProfile({ userId }) {
  const [user, setUser] = useState(null);

  useEffect(() => {
    // Este código executa após o componente ser renderizado
    fetch(`/api/users/${userId}`)
      .then(response => response.json())
      .then(data => setUser(data));
  }, [userId]); // Re-executa se userId mudar

  if (!user) {
    return <div>Carregando...</div>;
  }

  return (
    <div>
      <h1>{user.name}</h1>
      <p>{user.email}</p>
    </div>
  );
}
```

### 2.6 Vantagens dos Functional Components

1. **Sintaxe mais simples**: Menos código boilerplate
2. **Mais fácil de entender**: Funções são mais diretas que classes
3. **Melhor performance**: React pode otimizar melhor functional components
4. **Hooks**: Permitem usar estado e efeitos sem classes
5. **Testabilidade**: Funções puras são mais fáceis de testar

### 2.7 Comparação: Functional vs Class Components (Histórico)

**Class Component (antigo, ainda funciona mas não é recomendado):**
```jsx
class Welcome extends React.Component {
  constructor(props) {
    super(props);
    this.state = { count: 0 };
  }

  handleClick = () => {
    this.setState({ count: this.state.count + 1 });
  }

  render() {
    return (
      <div>
        <p>Contador: {this.state.count}</p>
        <button onClick={this.handleClick}>Incrementar</button>
      </div>
    );
  }
}
```

**Functional Component (moderno, recomendado):**
```jsx
function Welcome() {
  const [count, setCount] = useState(0);

  return (
    <div>
      <p>Contador: {count}</p>
      <button onClick={() => setCount(count + 1)}>Incrementar</button>
    </div>
  );
}
```

O functional component é mais conciso e fácil de entender.

---

## 3. JSX (JavaScript XML) e TSX

### 3.1 O Que é JSX?

**JSX** (JavaScript XML) é uma extensão de sintaxe do JavaScript que permite escrever código que parece HTML dentro do JavaScript. JSX não é HTML - é uma forma de descrever a estrutura da UI de forma declarativa.

### 3.2 Por Que JSX Existe?

**Sem JSX (usando React.createElement):**
```jsx
const element = React.createElement(
  'div',
  { className: 'container' },
  React.createElement('h1', null, 'Título'),
  React.createElement('p', null, 'Parágrafo')
);
```

**Com JSX:**
```jsx
const element = (
  <div className="container">
    <h1>Título</h1>
    <p>Parágrafo</p>
  </div>
);
```

JSX torna o código muito mais legível e fácil de escrever.

### 3.3 Regras Fundamentais do JSX

#### 3.3.1 JSX Deve Retornar um Único Elemento Raiz

```jsx
// ❌ ERRADO - múltiplos elementos raiz
function Component() {
  return (
    <h1>Título</h1>
    <p>Parágrafo</p>
  );
}

// ✅ CORRETO - um único elemento raiz
function Component() {
  return (
    <div>
      <h1>Título</h1>
      <p>Parágrafo</p>
    </div>
  );
}

// ✅ CORRETO - usando Fragment (não cria elemento no DOM)
function Component() {
  return (
    <>
      <h1>Título</h1>
      <p>Parágrafo</p>
    </>
  );
}
```

#### 3.3.2 Atributos Usam camelCase

```jsx
// ❌ ERRADO - HTML usa kebab-case
<div class="container" onclick="handleClick()">

// ✅ CORRETO - JSX usa camelCase
<div className="container" onClick={handleClick}>
```

**Exceções importantes:**
- `class` → `className` (porque `class` é palavra reservada em JavaScript)
- `for` → `htmlFor` (porque `for` é palavra reservada)
- Atributos de dados: `data-testid` permanece como está

#### 3.3.3 Expressões JavaScript Dentro de Chaves

```jsx
function Greeting({ name, age }) {
  const message = `Olá, ${name}!`;
  const isAdult = age >= 18;

  return (
    <div>
      <h1>{message}</h1>
      <p>Idade: {age}</p>
      <p>{isAdult ? 'Você é maior de idade' : 'Você é menor de idade'}</p>
      <p>Ano de nascimento: {2024 - age}</p>
    </div>
  );
}
```

**O que pode ir dentro de `{}`:**
- Variáveis
- Expressões matemáticas
- Chamadas de função
- Operadores ternários
- Arrays (cada elemento vira um elemento JSX)
- Outros componentes

**O que NÃO pode ir dentro de `{}`:**
- Objetos JavaScript diretamente (exceto em atributos de estilo)
- `if/else` statements (use operador ternário ou `&&`)
- `for` loops (use `map` para arrays)

#### 3.3.4 JSX Previne XSS Automaticamente

JSX escapa automaticamente valores para prevenir ataques XSS (Cross-Site Scripting):

```jsx
const userInput = '<script>alert("XSS")</script>';

// JSX escapa automaticamente, mostrando o texto literal
<div>{userInput}</div>
// Renderiza: <div>&lt;script&gt;alert("XSS")&lt;/script&gt;</div>
```

### 3.4 JSX é Transformado em JavaScript

Por baixo dos panos, ferramentas como Babel transformam JSX em chamadas `React.createElement`:

**JSX:**
```jsx
<div className="container">
  <h1>Título</h1>
</div>
```

**Transformado em:**
```jsx
React.createElement(
  'div',
  { className: 'container' },
  React.createElement('h1', null, 'Título')
);
```

### 3.5 TSX: JSX com TypeScript

**TSX** é JSX usado em arquivos TypeScript (`.tsx`). A diferença é que você tem tipagem estática:

```tsx
// Componente com props tipadas
interface ButtonProps {
  text: string;
  onClick: () => void;
  disabled?: boolean; // opcional
}

function Button({ text, onClick, disabled = false }: ButtonProps) {
  return (
    <button onClick={onClick} disabled={disabled}>
      {text}
    </button>
  );
}
```

**Vantagens do TSX:**
- Autocomplete melhor no editor
- Erros detectados em tempo de desenvolvimento
- Documentação implícita através de tipos
- Refatoração mais segura

**Quando usar TSX:**
- Projetos grandes onde tipagem ajuda
- Equipes que valorizam segurança de tipos
- Quando você quer prevenir erros comuns

**Quando JSX é suficiente:**
- Projetos pequenos
- Prototipagem rápida
- Quando a equipe prefere JavaScript puro

---

## 4. Props vs State

### 4.1 Props (Properties)

**Props** são dados passados de um componente pai para um filho. Elas são **imutáveis** (read-only) do ponto de vista do componente filho.

#### Características das Props

1. **Unidirecionais**: Fluem apenas de pai para filho
2. **Read-only**: Componente filho não pode modificar props diretamente
3. **Podem ser qualquer tipo**: strings, números, objetos, arrays, funções, componentes
4. **Podem ter valores padrão**

#### Exemplo de Props

```jsx
// Componente Pai
function App() {
  const userName = "João";
  const userAge = 25;
  const userHobbies = ["Leitura", "Programação"];

  return (
    <UserProfile 
      name={userName}
      age={userAge}
      hobbies={userHobbies}
      onUpdateName={(newName) => console.log(newName)}
    />
  );
}

// Componente Filho
function UserProfile({ name, age, hobbies, onUpdateName }) {
  return (
    <div>
      <h1>{name}</h1>
      <p>Idade: {age}</p>
      <ul>
        {hobbies.map((hobby, index) => (
          <li key={index}>{hobby}</li>
        ))}
      </ul>
      <button onClick={() => onUpdateName("Novo Nome")}>
        Atualizar Nome
      </button>
    </div>
  );
}
```

#### Props com Valores Padrão

```jsx
function Button({ text, color = "blue", size = "medium" }) {
  return (
    <button 
      style={{ 
        backgroundColor: color,
        fontSize: size === "large" ? "18px" : "14px"
      }}
    >
      {text}
    </button>
  );
}

// Uso sem especificar todas as props
<Button text="Clique aqui" />
// color será "blue" e size será "medium" por padrão
```

#### Props.children

`children` é uma prop especial que contém o conteúdo entre as tags de abertura e fechamento:

```jsx
function Card({ title, children }) {
  return (
    <div className="card">
      <h2>{title}</h2>
      <div className="card-content">
        {children}
      </div>
    </div>
  );
}

// Uso
<Card title="Meu Card">
  <p>Este é o conteúdo do card</p>
  <button>Clique aqui</button>
</Card>
```

Aqui, `children` contém `<p>` e `<button>`.

### 4.2 State (Estado)

**State** é a memória interna de um componente. É um objeto JavaScript que armazena dados que podem mudar ao longo do tempo e que, quando mudam, causam uma re-renderização do componente.

#### Características do State

1. **Local ao componente**: Cada instância de componente tem seu próprio estado
2. **Mutável**: Pode ser atualizado usando funções setter (como `setState` ou hooks)
3. **Causa re-renderização**: Quando state muda, React re-renderiza o componente
4. **Isolado**: Estado de um componente não afeta diretamente outros componentes

#### Exemplo de State com useState

```jsx
import { useState } from 'react';

function Counter() {
  // useState retorna [valor, funçãoParaAtualizar]
  const [count, setCount] = useState(0);
  const [isActive, setIsActive] = useState(false);

  return (
    <div>
      <p>Contador: {count}</p>
      <p>Status: {isActive ? 'Ativo' : 'Inativo'}</p>
      
      <button onClick={() => setCount(count + 1)}>
        Incrementar
      </button>
      
      <button onClick={() => setIsActive(!isActive)}>
        {isActive ? 'Desativar' : 'Ativar'}
      </button>
    </div>
  );
}
```

#### State com Objetos e Arrays

```jsx
import { useState } from 'react';

function TodoList() {
  const [todos, setTodos] = useState([]);
  const [input, setInput] = useState('');

  const addTodo = () => {
    // IMPORTANTE: Criar novo array, não modificar o existente
    setTodos([...todos, { id: Date.now(), text: input }]);
    setInput('');
  };

  const removeTodo = (id) => {
    // Filtrar para criar novo array sem o item removido
    setTodos(todos.filter(todo => todo.id !== id));
  };

  return (
    <div>
      <input 
        value={input}
        onChange={(e) => setInput(e.target.value)}
        placeholder="Nova tarefa"
      />
      <button onClick={addTodo}>Adicionar</button>
      
      <ul>
        {todos.map(todo => (
          <li key={todo.id}>
            {todo.text}
            <button onClick={() => removeTodo(todo.id)}>Remover</button>
          </li>
        ))}
      </ul>
    </div>
  );
}
```

**Regra de Ouro**: Nunca modifique state diretamente. Sempre crie novos objetos/arrays.

```jsx
// ❌ ERRADO - mutação direta
const [items, setItems] = useState([1, 2, 3]);
items.push(4); // NÃO FAÇA ISSO!

// ✅ CORRETO - criar novo array
const [items, setItems] = useState([1, 2, 3]);
setItems([...items, 4]); // Cria novo array
```

### 4.3 Diferenças Fundamentais: Props vs State

| Aspecto | Props | State |
|---------|-------|-------|
| **Origem** | Passado do componente pai | Gerenciado dentro do componente |
| **Mutabilidade** | Read-only (imutável) | Mutável (através de setters) |
| **Fluxo** | Unidirecional (pai → filho) | Local ao componente |
| **Quando usar** | Dados que vêm de fora | Dados que mudam internamente |
| **Exemplo** | Nome do usuário passado como prop | Contador interno do componente |

### 4.4 Quando Usar Props vs State?

**Use Props quando:**
- Dados vêm do componente pai
- Dados não mudam dentro do componente
- Você quer passar configurações ou callbacks
- Dados são compartilhados entre múltiplos componentes

**Use State quando:**
- Dados são específicos do componente
- Dados mudam ao longo do tempo
- Dados são resultado de interações do usuário
- Dados precisam causar re-renderização quando mudam

### 4.5 Lifting State Up (Elevando Estado)

Às vezes, múltiplos componentes precisam compartilhar o mesmo estado. Nesse caso, "elevamos" o estado para o componente pai comum:

```jsx
// ❌ Estado duplicado (ruim)
function Counter1() {
  const [count, setCount] = useState(0);
  return <div>Contador 1: {count}</div>;
}

function Counter2() {
  const [count, setCount] = useState(0);
  return <div>Contador 2: {count}</div>;
}

// ✅ Estado compartilhado (bom)
function App() {
  const [count, setCount] = useState(0);
  
  return (
    <div>
      <Counter1 count={count} />
      <Counter2 count={count} />
      <button onClick={() => setCount(count + 1)}>
        Incrementar ambos
      </button>
    </div>
  );
}

function Counter1({ count }) {
  return <div>Contador 1: {count}</div>;
}

function Counter2({ count }) {
  return <div>Contador 2: {count}</div>;
}
```

---

## 5. Conditional Rendering (Renderização Condicional)

### 5.1 O Que é Conditional Rendering?

**Conditional Rendering** é a técnica de renderizar diferentes elementos ou componentes baseado em condições. Em React, isso funciona da mesma forma que condições em JavaScript.

### 5.2 Métodos de Conditional Rendering

#### 5.2.1 Operador Ternário

```jsx
function Greeting({ isLoggedIn }) {
  return (
    <div>
      {isLoggedIn ? (
        <h1>Bem-vindo de volta!</h1>
      ) : (
        <h1>Por favor, faça login</h1>
      )}
    </div>
  );
}
```

#### 5.2.2 Operador Lógico && (AND)

```jsx
function Notification({ message, show }) {
  return (
    <div>
      {show && <div className="notification">{message}</div>}
    </div>
  );
}

// Se show for false, nada é renderizado
// Se show for true, a notificação é renderizada
```

**Cuidado**: Se a condição for `0`, React renderizará `0`:
```jsx
// Se count for 0, isso renderiza "0" na tela
{count && <div>Você tem {count} itens</div>}

// Solução: converter para booleano
{count > 0 && <div>Você tem {count} itens</div>}
```

#### 5.2.3 Múltiplas Condições

```jsx
function StatusMessage({ status }) {
  if (status === 'loading') {
    return <div>Carregando...</div>;
  }
  
  if (status === 'error') {
    return <div>Erro ao carregar</div>;
  }
  
  if (status === 'success') {
    return <div>Sucesso!</div>;
  }
  
  return <div>Status desconhecido</div>;
}
```

#### 5.2.4 Early Return

```jsx
function UserProfile({ user }) {
  // Early return se não houver usuário
  if (!user) {
    return <div>Usuário não encontrado</div>;
  }

  // Resto do componente só executa se user existir
  return (
    <div>
      <h1>{user.name}</h1>
      <p>{user.email}</p>
    </div>
  );
}
```

#### 5.2.5 Switch Case (dentro de função)

```jsx
function getStatusComponent(status) {
  switch (status) {
    case 'loading':
      return <Spinner />;
    case 'success':
      return <SuccessMessage />;
    case 'error':
      return <ErrorMessage />;
    default:
      return <UnknownStatus />;
  }
}

function App({ status }) {
  return <div>{getStatusComponent(status)}</div>;
}
```

### 5.3 Conditional Rendering com Arrays

```jsx
function ProductList({ products, showEmptyMessage }) {
  if (products.length === 0) {
    return showEmptyMessage ? (
      <p>Nenhum produto encontrado</p>
    ) : null;
  }

  return (
    <ul>
      {products.map(product => (
        <li key={product.id}>{product.name}</li>
      ))}
    </ul>
  );
}
```

### 5.4 Evitando Re-renderizações Desnecessárias

```jsx
// ❌ Componente sempre re-renderiza, mesmo quando não precisa
function ExpensiveComponent({ data, show }) {
  const processedData = expensiveProcessing(data);
  
  return show ? <div>{processedData}</div> : null;
}

// ✅ Processamento só acontece quando necessário
function ExpensiveComponent({ data, show }) {
  if (!show) {
    return null;
  }
  
  const processedData = expensiveProcessing(data);
  return <div>{processedData}</div>;
}
```

---

## 6. Composition vs Inheritance

### 6.1 O Princípio da Composição

React tem um **modelo de composição poderoso** e recomenda usar composição em vez de herança para reutilizar código entre componentes.

### 6.2 Por Que Não Herança?

Em programação orientada a objetos tradicional, você pode criar uma classe base e outras classes herdam dela. React não usa esse modelo porque:

1. **Herança cria acoplamento forte**: Mudanças na classe base afetam todas as classes filhas
2. **Herança é rígida**: Difícil de modificar comportamento sem quebrar outras partes
3. **Herança não escala bem**: Hierarquias profundas são difíceis de entender e manter
4. **JavaScript não tem herança múltipla**: Limita flexibilidade

### 6.3 Composição: A Abordagem React

Composição significa construir componentes maiores combinando componentes menores. É como construir com blocos de Lego.

#### 6.3.1 Composição com children

```jsx
// Componente genérico e reutilizável
function Container({ children, title }) {
  return (
    <div className="container">
      {title && <h2>{title}</h2>}
      <div className="content">
        {children}
      </div>
    </div>
  );
}

// Usando composição
function App() {
  return (
    <Container title="Meu Conteúdo">
      <p>Este é o conteúdo dentro do container</p>
      <button>Clique aqui</button>
    </Container>
  );
}
```

#### 6.3.2 Composição com Props Específicas

```jsx
// Componentes pequenos e focados
function Button({ children, onClick, variant = "primary" }) {
  return (
    <button 
      className={`btn btn-${variant}`}
      onClick={onClick}
    >
      {children}
    </button>
  );
}

function Card({ children, title, footer }) {
  return (
    <div className="card">
      {title && <div className="card-header">{title}</div>}
      <div className="card-body">{children}</div>
      {footer && <div className="card-footer">{footer}</div>}
    </div>
  );
}

// Composição: combinando componentes
function ProductCard({ product }) {
  return (
    <Card
      title={product.name}
      footer={
        <Button onClick={() => addToCart(product.id)}>
          Adicionar ao Carrinho
        </Button>
      }
    >
      <p>Preço: R$ {product.price}</p>
      <p>{product.description}</p>
    </Card>
  );
}
```

#### 6.3.3 Composição com Múltiplos "Slots"

```jsx
function Dialog({ header, body, footer }) {
  return (
    <div className="dialog">
      <div className="dialog-header">{header}</div>
      <div className="dialog-body">{body}</div>
      <div className="dialog-footer">{footer}</div>
    </div>
  );
}

function App() {
  return (
    <Dialog
      header={<h2>Confirmar Ação</h2>}
      body={<p>Tem certeza que deseja continuar?</p>}
      footer={
        <>
          <Button>Cancelar</Button>
          <Button variant="primary">Confirmar</Button>
        </>
      }
    />
  );
}
```

#### 6.3.4 Composição com HOCs (Higher-Order Components)

HOCs são funções que recebem um componente e retornam um novo componente com funcionalidades adicionais:

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

// Componente com funcionalidade de loading
const UserListWithLoading = withLoading(UserList);

// Uso
function App() {
  const [isLoading, setIsLoading] = useState(true);
  const [users, setUsers] = useState([]);

  return (
    <UserListWithLoading 
      isLoading={isLoading} 
      users={users} 
    />
  );
}
```

**Nota**: Hooks são geralmente preferidos sobre HOCs na maioria dos casos modernos.

#### 6.3.5 Composição com Render Props

Render props é um padrão onde um componente recebe uma função como prop que retorna JSX:

```jsx
function DataFetcher({ url, render }) {
  const [data, setData] = useState(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    fetch(url)
      .then(res => res.json())
      .then(data => {
        setData(data);
        setLoading(false);
      });
  }, [url]);

  return render({ data, loading });
}

// Uso
function App() {
  return (
    <DataFetcher
      url="/api/users"
      render={({ data, loading }) => {
        if (loading) return <div>Carregando...</div>;
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

**Nota**: Hooks customizados são geralmente preferidos sobre render props.

### 6.4 Quando Usar Composição vs Herança?

**Sempre use Composição quando:**
- Você quer reutilizar código entre componentes
- Você quer criar componentes flexíveis
- Você quer evitar acoplamento forte
- Você está trabalhando com React

**Nunca use Herança para:**
- Reutilizar código entre componentes React
- Estender funcionalidades de componentes

**Herança pode ser usada para:**
- Classes JavaScript regulares (não componentes React)
- Estruturas de dados
- Mas mesmo assim, composição geralmente é preferível

### 6.5 Vantagens da Composição

1. **Flexibilidade**: Fácil de modificar e estender
2. **Reutilização**: Componentes pequenos podem ser combinados de várias formas
3. **Testabilidade**: Componentes pequenos são mais fáceis de testar
4. **Manutenibilidade**: Mudanças em um componente não afetam outros
5. **Clareza**: Fica claro o que cada componente faz

---

## 7. Padrões Comuns de Components

### 7.1 Presentational vs Container Components

**Presentational Components** (Componentes de Apresentação):
- Focam em como as coisas aparecem
- Recebem dados via props
- Não gerenciam estado (ou apenas estado de UI)
- São reutilizáveis

```jsx
// Presentational Component
function Button({ text, onClick, disabled }) {
  return (
    <button onClick={onClick} disabled={disabled}>
      {text}
    </button>
  );
}
```

**Container Components** (Componentes de Container):
- Focam em como as coisas funcionam
- Gerenciam estado e lógica de negócio
- Fazem chamadas de API
- Passam dados para presentational components

```jsx
// Container Component
function UserListContainer() {
  const [users, setUsers] = useState([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    fetch('/api/users')
      .then(res => res.json())
      .then(data => {
        setUsers(data);
        setLoading(false);
      });
  }, []);

  if (loading) {
    return <div>Carregando...</div>;
  }

  return <UserList users={users} />; // Presentational
}
```

**Nota**: Com hooks, essa distinção é menos rígida, mas ainda é um padrão útil.

### 7.2 Controlled vs Uncontrolled Components

**Controlled Components**: React controla o valor do input através de state:

```jsx
function ControlledInput() {
  const [value, setValue] = useState('');

  return (
    <input
      value={value}
      onChange={(e) => setValue(e.target.value)}
    />
  );
}
```

**Uncontrolled Components**: O DOM controla o valor (usando refs):

```jsx
function UncontrolledInput() {
  const inputRef = useRef(null);

  const handleSubmit = () => {
    console.log(inputRef.current.value);
  };

  return (
    <div>
      <input ref={inputRef} />
      <button onClick={handleSubmit}>Enviar</button>
    </div>
  );
}
```

Geralmente, controlled components são preferidos.

---

## 8. Conclusão

Components são o coração do React. Eles permitem:

- **Organizar código** de forma modular
- **Reutilizar** lógica e apresentação
- **Manter** aplicações grandes de forma gerenciável
- **Testar** funcionalidades isoladamente
- **Colaborar** em equipes grandes

### Conceitos-Chave Revisados

1. **Components**: Blocos de construção reutilizáveis
2. **Functional Components**: Forma moderna de criar componentes
3. **JSX/TSX**: Sintaxe que permite HTML-like em JavaScript
4. **Props**: Dados passados de pai para filho (read-only)
5. **State**: Memória interna do componente (mutável)
6. **Conditional Rendering**: Renderizar baseado em condições
7. **Composition**: Combinar componentes menores para criar maiores

### Próximos Passos

Agora que você entende components, você pode:
- Criar componentes reutilizáveis
- Gerenciar estado e props
- Renderizar condicionalmente
- Compor interfaces complexas

Na próxima aula, vamos aprofundar em hooks, eventos, e padrões avançados de components.

---

**Lembre-se**: A melhor forma de aprender components é praticando. Crie componentes pequenos, combine-os, experimente diferentes padrões. A prática é essencial para internalizar esses conceitos!



