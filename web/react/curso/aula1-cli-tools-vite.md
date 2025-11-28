# 🚀 Aula 1: CLI Tools e Vite - Introdução ao React

## 📋 Sumário da Aula
- O que é React e por que ele existe?
- Como o React funciona (conceitos fundamentais)
- O que são CLI Tools e por que são essenciais
- Vite: a ferramenta moderna para desenvolvimento React
- Criando seu primeiro projeto React com Vite
- Estrutura de um projeto React
- Exemplos práticos e código funcionando

---

## 🎯 1. O que é React e Por Que Ele Existe?

### O Problema que o React Resolve

Imagine que você precisa construir uma página web que mostra uma lista de produtos. Cada produto tem:
- Nome
- Preço
- Botão "Adicionar ao carrinho"
- Contador de itens no carrinho

**Sem React (JavaScript puro):**
```javascript
// Você teria que manipular o DOM manualmente
const productList = document.getElementById('product-list');
const cartCount = document.getElementById('cart-count');
let count = 0;

function addToCart(productId) {
  count++;
  cartCount.textContent = count; // Atualizar manualmente
  // Criar elementos HTML manualmente
  const product = document.createElement('div');
  product.innerHTML = `<h3>${productName}</h3><p>${price}</p>`;
  productList.appendChild(product);
}
```

**Problemas dessa abordagem:**
- Código repetitivo e difícil de manter
- Fácil de introduzir bugs (esquecer de atualizar algum elemento)
- Difícil sincronizar estado entre diferentes partes da página
- Performance ruim quando há muitas atualizações

**Com React:**
```jsx
function Product({ name, price, onAddToCart }) {
  return (
    <div>
      <h3>{name}</h3>
      <p>{price}</p>
      <button onClick={onAddToCart}>Adicionar ao carrinho</button>
    </div>
  );
}

function App() {
  const [cartCount, setCartCount] = useState(0);
  
  return (
    <div>
      <CartCount count={cartCount} />
      <Product 
        name="Notebook" 
        price="R$ 2.500" 
        onAddToCart={() => setCartCount(cartCount + 1)} 
      />
    </div>
  );
}
```

**Vantagens:**
- Código declarativo (você descreve COMO deve ser, não COMO fazer)
- Estado gerenciado automaticamente
- Atualizações eficientes (React só atualiza o que mudou)
- Componentes reutilizáveis

### O Que é React, Tecnicamente?

**React** é uma biblioteca JavaScript criada pelo Facebook (Meta) em 2013 para construir interfaces de usuário (UI). Ele não é um framework completo - é focado apenas na camada de visualização.

**Características principais:**
1. **Baseado em Componentes**: Você divide a UI em pedaços reutilizáveis (componentes)
2. **Declarativo**: Você descreve como a UI deve ser, não como manipulá-la
3. **Virtual DOM**: React cria uma representação em memória do DOM e só atualiza o que realmente mudou
4. **Unidirecional**: Dados fluem de componentes pais para filhos (one-way data flow)

---

## 🧠 2. Como o React Funciona - Conceitos Fundamentais

### 2.1 Componentes: Os Blocos de Construção

Pense em componentes como **peças de Lego**. Cada peça tem uma função específica e você combina várias peças para construir algo maior.

**Analogia do Mundo Real:**
- Um carro tem componentes: motor, rodas, volante, bancos
- Cada componente tem uma responsabilidade
- Você pode reutilizar componentes (muitos carros usam o mesmo tipo de roda)

**No React:**
```jsx
// Componente simples - uma função que retorna JSX
function Button() {
  return <button>Clique aqui</button>;
}

// Componente com props (propriedades) - como argumentos de uma função
function Button({ text, color }) {
  return <button style={{ backgroundColor: color }}>{text}</button>;
}

// Usando o componente
function App() {
  return (
    <div>
      <Button text="Salvar" color="green" />
      <Button text="Cancelar" color="red" />
    </div>
  );
}
```

### 2.2 JSX: JavaScript + XML

**JSX** parece HTML, mas é JavaScript. Ele permite escrever código que parece HTML dentro do JavaScript.

```jsx
// Isso é JSX
const element = <h1>Olá, React!</h1>;

// Por baixo dos panos, React transforma em:
const element = React.createElement('h1', null, 'Olá, React!');
```

**Regras importantes do JSX:**
- Deve retornar um único elemento raiz (ou usar Fragment)
- Atributos usam camelCase (`className` em vez de `class`)
- Expressões JavaScript dentro de `{}`
- `if/else` não funciona diretamente, use operador ternário ou `&&`

```jsx
function Greeting({ name, isLoggedIn }) {
  return (
    <div>
      {isLoggedIn ? (
        <h1>Bem-vindo, {name}!</h1>
      ) : (
        <h1>Por favor, faça login</h1>
      )}
      {name && <p>Seu nome tem {name.length} letras</p>}
    </div>
  );
}
```

### 2.3 Estado (State): A Memória dos Componentes

**Estado** é como a "memória" de um componente. É como uma gaveta que guarda informações que podem mudar.

**Analogia:**
- Uma lâmpada tem estado: ligada ou desligada
- Quando você clica no interruptor, o estado muda
- A lâmpada reage ao novo estado (acende ou apaga)

**No React:**
```jsx
import { useState } from 'react';

function Counter() {
  // useState retorna [valor, função para atualizar]
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

**Por que isso é importante?**
- Quando `setCount` é chamado, React **re-renderiza** o componente
- React compara o novo estado com o anterior
- Só atualiza o que mudou no DOM real (eficiente!)

### 2.4 Props: Passando Dados Entre Componentes

**Props** (propriedades) são como argumentos de função. Você passa dados de um componente pai para um filho.

```jsx
// Componente Pai
function App() {
  const userName = "João";
  const userAge = 25;
  
  return <UserProfile name={userName} age={userAge} />;
}

// Componente Filho recebe props
function UserProfile({ name, age }) {
  return (
    <div>
      <h2>{name}</h2>
      <p>Idade: {age} anos</p>
    </div>
  );
}
```

**Fluxo de dados:**
- **Unidirecional**: Pai → Filho (dados fluem para baixo)
- Filho não pode modificar props diretamente
- Se precisar mudar, o pai passa uma função

### 2.5 Virtual DOM: A Mágica da Performance

**DOM (Document Object Model)** é a representação da página no navegador. Manipular o DOM é lento.

**Virtual DOM** é uma cópia do DOM em memória (JavaScript). React usa isso para:
1. Criar uma nova versão do Virtual DOM quando algo muda
2. Comparar (diff) a versão antiga com a nova
3. Atualizar apenas as partes diferentes no DOM real

**Analogia:**
- DOM Real = uma casa de verdade
- Virtual DOM = uma planta/maquete da casa
- Quando você quer mudar algo, você:
  1. Desenha uma nova planta (novo Virtual DOM)
  2. Compara com a planta antiga (diff)
  3. Só constrói/muda o que é diferente na casa real

**Por que isso é rápido?**
- Comparar objetos JavaScript é muito mais rápido que manipular o DOM
- React agrupa várias mudanças e aplica de uma vez
- Evita re-renderizações desnecessárias

---

## 🛠️ 3. O Que São CLI Tools e Por Que São Essenciais?

### 3.1 O Que é uma CLI Tool?

**CLI (Command Line Interface)** é uma ferramenta que você usa no terminal/comando para automatizar tarefas.

**Exemplos do dia a dia:**
- `git` - controla versões do código
- `npm` - gerencia pacotes JavaScript
- `ls` (Linux/Mac) ou `dir` (Windows) - lista arquivos

**No contexto React:**
- Ferramentas que criam projetos
- Configuram o ambiente de desenvolvimento
- Compilam e otimizam código
- Iniciam servidores de desenvolvimento

### 3.2 Por Que Usar CLI Tools?

**Sem CLI Tool (fazer tudo manualmente):**
1. Criar estrutura de pastas
2. Instalar dependências manualmente
3. Configurar Babel (converte JSX para JavaScript)
4. Configurar Webpack (empacota o código)
5. Configurar servidor de desenvolvimento
6. Configurar hot reload (atualização automática)
7. Configurar build para produção

**Tempo estimado:** 2-4 horas (e você pode errar!)

**Com CLI Tool:**
```bash
npm create vite@latest meu-projeto -- --template react
cd meu-projeto
npm install
npm run dev
```

**Tempo estimado:** 2 minutos

**Benefícios:**
- ✅ Padrão da indústria (todos usam a mesma estrutura)
- ✅ Configuração otimizada e testada
- ✅ Foco no código, não na configuração
- ✅ Atualizações e melhorias automáticas

---

## ⚡ 4. Vite: A Ferramenta Moderna para React

### 4.1 O Que é Vite?

**Vite** (pronuncia-se "veet", francês para "rápido") é um build tool criado por Evan You (criador do Vue.js) em 2020.

**Problema que resolve:**
- Ferramentas antigas (como Create React App) eram lentas
- Esperavam empacotar TODO o código antes de iniciar
- Em projetos grandes, podia levar minutos para iniciar

### 4.2 Como o Vite Funciona?

**Desenvolvimento (dev mode):**
1. Usa **ESM (ES Modules)** nativo do navegador
2. Serve arquivos diretamente, sem empacotar
3. Compila sob demanda (só o que você está usando)
4. Hot Module Replacement (HMR) instantâneo

**Produção (build):**
1. Usa **Rollup** (empacotador rápido)
2. Otimiza e minifica o código
3. Code splitting automático
4. Tree shaking (remove código não usado)

**Comparação de velocidade:**

| Operação | Create React App | Vite |
|----------|------------------|------|
| Iniciar servidor | 10-30s | <1s |
| Hot Reload | 1-3s | <50ms |
| Build produção | 30-60s | 5-15s |

### 4.3 Por Que Vite é Melhor?

1. **Velocidade**: Início instantâneo, mesmo em projetos grandes
2. **DX (Developer Experience)**: Feedback imediato ao salvar
3. **Moderno**: Usa recursos nativos do navegador
4. **Simples**: Menos configuração, funciona out-of-the-box
5. **Flexível**: Fácil de customizar quando necessário

---

## 🏗️ 5. Criando Seu Primeiro Projeto React com Vite

### 5.1 Pré-requisitos

Você precisa ter instalado:
- **Node.js** (versão 18 ou superior)
- **npm** (vem com Node.js) ou **yarn** ou **pnpm**

**Verificar instalação:**
```bash
node --version  # Deve mostrar v18.x.x ou superior
npm --version   # Deve mostrar 9.x.x ou superior
```

### 5.2 Criando o Projeto

```bash
# Criar projeto React com Vite
npm create vite@latest meu-primeiro-react -- --template react

# Entrar na pasta
cd meu-primeiro-react

# Instalar dependências
npm install

# Iniciar servidor de desenvolvimento
npm run dev
```

**O que acontece:**
1. Vite cria a estrutura de pastas
2. Instala React e dependências
3. Configura tudo automaticamente
4. Abre em `http://localhost:5173` (porta padrão do Vite)

### 5.3 Estrutura do Projeto Criado

```
meu-primeiro-react/
├── node_modules/          # Dependências instaladas (não edite)
├── public/               # Arquivos estáticos (imagens, etc)
│   └── vite.svg
├── src/                  # Seu código fonte
│   ├── App.jsx          # Componente principal
│   ├── App.css          # Estilos do App
│   ├── index.css        # Estilos globais
│   └── main.jsx         # Ponto de entrada
├── .gitignore           # Arquivos ignorados pelo Git
├── index.html           # HTML principal
├── package.json         # Dependências e scripts
├── vite.config.js       # Configuração do Vite
└── README.md            # Documentação
```

### 5.4 Entendendo os Arquivos Principais

#### `index.html`
```html
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <link rel="icon" type="image/svg+xml" href="/vite.svg" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>Vite + React</title>
  </head>
  <body>
    <div id="root"></div>
    <script type="module" src="/src/main.jsx"></script>
  </body>
</html>
```

**Pontos importantes:**
- `<div id="root"></div>` - Onde React vai "montar" a aplicação
- `<script type="module">` - Permite usar ESM nativo
- `/src/main.jsx` - Ponto de entrada do JavaScript

#### `src/main.jsx`
```jsx
import React from 'react'
import ReactDOM from 'react-dom/client'
import App from './App.jsx'
import './index.css'

ReactDOM.createRoot(document.getElementById('root')).render(
  <React.StrictMode>
    <App />
  </React.StrictMode>,
)
```

**O que faz:**
1. Importa React e ReactDOM
2. Importa o componente App
3. Cria uma "raiz" React no elemento `#root`
4. Renderiza o componente `<App />` dentro dela
5. `<React.StrictMode>` - Modo estrito (ajuda a encontrar problemas)

#### `src/App.jsx`
```jsx
import { useState } from 'react'
import './App.css'

function App() {
  const [count, setCount] = useState(0)

  return (
    <div className="App">
      <h1>Vite + React</h1>
      <div className="card">
        <button onClick={() => setCount((count) => count + 1)}>
          count is {count}
        </button>
        <p>
          Edit <code>src/App.jsx</code> and save to test HMR
        </p>
      </div>
    </div>
  )
}

export default App
```

**Análise linha por linha:**
- `import { useState } from 'react'` - Importa o hook useState
- `const [count, setCount] = useState(0)` - Cria estado inicializado com 0
- `onClick={() => setCount((count) => count + 1)}` - Atualiza estado ao clicar
- `export default App` - Exporta o componente para ser usado em outros lugares

---

## 💡 6. Exemplos Práticos

### Exemplo 1: Contador Simples

```jsx
import { useState } from 'react'

function Counter() {
  const [count, setCount] = useState(0)

  return (
    <div>
      <h2>Contador: {count}</h2>
      <button onClick={() => setCount(count + 1)}>
        Incrementar
      </button>
      <button onClick={() => setCount(count - 1)}>
        Decrementar
      </button>
      <button onClick={() => setCount(0)}>
        Resetar
      </button>
    </div>
  )
}

export default Counter
```

**O que acontece:**
1. Estado inicial: `count = 0`
2. Ao clicar "Incrementar": `count` vira `1`, React re-renderiza
3. Ao clicar "Decrementar": `count` vira `0`, React re-renderiza
4. Ao clicar "Resetar": `count` vira `0`, React re-renderiza

### Exemplo 2: Lista de Tarefas Básica

```jsx
import { useState } from 'react'

function TodoList() {
  const [todos, setTodos] = useState([])
  const [input, setInput] = useState('')

  function addTodo() {
    if (input.trim() !== '') {
      setTodos([...todos, input])
      setInput('')
    }
  }

  return (
    <div>
      <h2>Minhas Tarefas</h2>
      <div>
        <input
          value={input}
          onChange={(e) => setInput(e.target.value)}
          placeholder="Nova tarefa"
        />
        <button onClick={addTodo}>Adicionar</button>
      </div>
      <ul>
        {todos.map((todo, index) => (
          <li key={index}>{todo}</li>
        ))}
      </ul>
    </div>
  )
}

export default TodoList
```

**Conceitos importantes:**
- **Estado de array**: `useState([])`
- **Spread operator**: `[...todos, input]` - cria novo array com todos + novo item
- **Input controlado**: `value={input}` + `onChange` - React controla o valor
- **Renderização de lista**: `todos.map()` - cria elemento para cada item
- **Key prop**: `key={index}` - ajuda React a identificar cada item

### Exemplo 3: Componente com Props

```jsx
// Componente Card reutilizável
function Card({ title, description, color }) {
  return (
    <div style={{ 
      border: `2px solid ${color}`, 
      padding: '20px', 
      margin: '10px',
      borderRadius: '8px'
    }}>
      <h3>{title}</h3>
      <p>{description}</p>
    </div>
  )
}

// Usando o componente
function App() {
  return (
    <div>
      <Card 
        title="React" 
        description="Biblioteca para construir UIs" 
        color="blue" 
      />
      <Card 
        title="Vite" 
        description="Build tool super rápido" 
        color="green" 
      />
      <Card 
        title="JavaScript" 
        description="Linguagem de programação" 
        color="yellow" 
      />
    </div>
  )
}

export default App
```

**Por que isso é poderoso:**
- Escreve o componente uma vez
- Reutiliza com dados diferentes
- Fácil de manter (mudança em um lugar afeta todos)

---

## 🎓 7. Conceitos-Chave para Entender

### 7.1 Renderização

**Renderizar** = React cria elementos e os coloca no DOM.

**Quando React renderiza:**
1. Primeira vez (mount) - quando componente aparece
2. Quando estado muda (`setState` é chamado)
3. Quando props mudam (componente pai passa novos dados)
4. Quando componente pai re-renderiza

**Importante:** React é inteligente - só re-renderiza o que mudou!

### 7.2 Imutabilidade

**Regra de ouro:** Nunca modifique estado diretamente!

```jsx
// ❌ ERRADO - mutação direta
const [items, setItems] = useState([1, 2, 3])
items.push(4) // NÃO FAÇA ISSO!

// ✅ CORRETO - criar novo array
const [items, setItems] = useState([1, 2, 3])
setItems([...items, 4]) // Cria novo array
```

**Por quê?**
- React compara referências para saber se mudou
- Se você modifica o mesmo objeto, React não detecta mudança
- Pode causar bugs e performance ruim

### 7.3 Composição de Componentes

**Composição** = construir componentes maiores usando menores.

```jsx
// Componentes pequenos e focados
function Button({ children, onClick }) {
  return <button onClick={onClick}>{children}</button>
}

function Card({ children }) {
  return <div className="card">{children}</div>
}

// Composição - juntando componentes
function ProductCard({ name, price }) {
  return (
    <Card>
      <h3>{name}</h3>
      <p>{price}</p>
      <Button onClick={() => alert('Adicionado!')}>
        Comprar
      </Button>
    </Card>
  )
}
```

**Benefício:** Componentes pequenos são mais fáceis de testar e reutilizar.

---

## 📊 8. Impacto na Performance e Experiência do Usuário

### Por Que React é Rápido?

1. **Virtual DOM**: Compara mudanças antes de atualizar DOM real
2. **Reconciliação**: Algoritmo inteligente que minimiza operações no DOM
3. **Batching**: Agrupa várias atualizações de estado em uma renderização
4. **Code Splitting**: Carrega apenas o código necessário

### Impacto na Manutenção

1. **Componentes reutilizáveis**: Menos código duplicado
2. **Separação de responsabilidades**: Cada componente tem uma função
3. **Fluxo de dados previsível**: Fácil de debugar
4. **Ecosystem**: Muitas bibliotecas e ferramentas

---

## ✅ 9. Checklist de Compreensão

Antes de avançar, certifique-se de entender:

- [ ] O que é React e por que foi criado
- [ ] O que são componentes e como funcionam
- [ ] O que é JSX e suas regras básicas
- [ ] Como funciona estado (`useState`)
- [ ] Como passar dados com props
- [ ] O que é Virtual DOM e por que é importante
- [ ] Por que usar CLI Tools
- [ ] O que é Vite e por que é melhor que alternativas
- [ ] Como criar um projeto React com Vite
- [ ] Estrutura básica de um projeto React
- [ ] Por que imutabilidade é importante

---

## 🎯 10. Próximos Passos

Na próxima aula, vamos aprofundar em:
- Componentes funcionais vs classes (histórico)
- Hooks mais avançados (`useEffect`, `useContext`)
- Eventos e formulários
- Estilização (CSS Modules, Styled Components)
- Debugging com React DevTools

---

## 📝 Resumo

**React** é uma biblioteca para construir interfaces de usuário baseada em componentes. Ele resolve problemas de sincronização de estado e performance usando Virtual DOM.

**Vite** é uma ferramenta moderna que torna o desenvolvimento React muito mais rápido, usando ESM nativo e compilação sob demanda.

**Componentes** são funções que retornam JSX, podem receber props e gerenciar estado.

**Estado** é a memória do componente, atualizado com `useState`, e mudanças causam re-renderização.

**Props** são dados passados de componente pai para filho, em fluxo unidirecional.

**Virtual DOM** é uma representação em memória que React usa para atualizar o DOM real de forma eficiente.

---

## 🔗 Recursos Adicionais

- [Documentação oficial do React](https://react.dev)
- [Documentação do Vite](https://vitejs.dev)
- [React DevTools](https://react.dev/learn/react-developer-tools) - Extensão do navegador para debugar

---

**Lembre-se:** React é sobre pensar em componentes, estado e como os dados fluem. Pratique criando pequenos componentes e vendo como eles se conectam. A prática é essencial para internalizar esses conceitos!

