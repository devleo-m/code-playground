# Aula 2 - Performance, Boas Práticas e Otimização de Components

## Introdução: Por Que Performance e Boas Práticas Importam?

No desenvolvimento React, escrever código que funciona é apenas o primeiro passo. Escrever código que é **performático, manutenível e escalável** é o que separa desenvolvedores iniciantes de profissionais.

**Regra de Ouro**: É muito mais fácil escrever código eficiente desde o início do que otimizar código ineficiente depois.

---

## 1. Performance: Entendendo Re-renderizações

### 1.1 Quando React Re-renderiza?

React re-renderiza um component quando:
1. **State muda**: `setState` ou hooks de state são chamados
2. **Props mudam**: Componente pai passa novas props
3. **Componente pai re-renderiza**: Mesmo que props não mudem, se o pai re-renderiza, filhos também re-renderizam
4. **Context muda**: Se o component consome um Context que mudou

### 1.2 O Problema das Re-renderizações Desnecessárias

**Exemplo Problemático**:

```jsx
// ❌ Component re-renderiza mesmo quando não precisa
function ListaProdutos({ produtos }) {
  console.log('ListaProdutos renderizou'); // Executa toda vez!
  
  return (
    <div>
      {produtos.map(produto => (
        <ProdutoItem 
          key={produto.id}
          produto={produto}
          onComprar={() => console.log('Comprar')} // Nova função a cada render!
        />
      ))}
    </div>
  );
}

function ProdutoItem({ produto, onComprar }) {
  console.log('ProdutoItem renderizou', produto.id); // Re-renderiza sempre!
  return (
    <div>
      <h3>{produto.nome}</h3>
      <button onClick={onComprar}>Comprar</button>
    </div>
  );
}
```

**Problemas**:
- `onComprar` é uma nova função a cada render, causando re-renderização de todos os `ProdutoItem`
- Se `ListaProdutos` re-renderiza, todos os `ProdutoItem` re-renderizam, mesmo que seus props não tenham mudado

### 1.3 Soluções: Memoização e Callbacks Estáveis

#### React.memo: Memorizando Components

```jsx
import { memo } from 'react';

// ✅ Component só re-renderiza se props mudarem
const ProdutoItem = memo(function ProdutoItem({ produto, onComprar }) {
  console.log('ProdutoItem renderizou', produto.id);
  return (
    <div>
      <h3>{produto.nome}</h3>
      <button onClick={onComprar}>Comprar</button>
    </div>
  );
});

// Agora ProdutoItem só re-renderiza se produto ou onComprar mudarem
```

**Quando usar `React.memo`**:
- ✅ Componentes que re-renderizam frequentemente com as mesmas props
- ✅ Componentes "caros" de renderizar (com cálculos complexos)
- ✅ Componentes em listas grandes
- ❌ Componentes simples que sempre recebem props diferentes
- ❌ Componentes que sempre re-renderizam de qualquer forma

#### useMemo: Memorizando Valores Calculados

```jsx
import { useMemo } from 'react';

function ListaProdutos({ produtos, filtro }) {
  // ✅ Só recalcula se produtos ou filtro mudarem
  const produtosFiltrados = useMemo(() => {
    console.log('Filtrando produtos...'); // Só executa quando necessário
    return produtos.filter(produto => 
      produto.nome.toLowerCase().includes(filtro.toLowerCase())
    );
  }, [produtos, filtro]);

  return (
    <div>
      {produtosFiltrados.map(produto => (
        <ProdutoItem key={produto.id} produto={produto} />
      ))}
    </div>
  );
}
```

**Quando usar `useMemo`**:
- ✅ Cálculos caros (filtros, ordenações, transformações)
- ✅ Valores derivados que são usados em múltiplos lugares
- ✅ Quando o cálculo é mais caro que a memorização
- ❌ Valores simples que são baratos de calcular
- ❌ Quando as dependências mudam frequentemente

#### useCallback: Memorizando Funções

```jsx
import { useCallback } from 'react';

function ListaProdutos({ produtos }) {
  // ✅ Função estável - mesma referência entre renders
  const handleComprar = useCallback((produtoId) => {
    console.log('Comprar produto', produtoId);
    // Lógica de compra
  }, []); // Dependências vazias = função nunca muda

  return (
    <div>
      {produtos.map(produto => (
        <ProdutoItem 
          key={produto.id}
          produto={produto}
          onComprar={() => handleComprar(produto.id)}
        />
      ))}
    </div>
  );
}
```

**Quando usar `useCallback`**:
- ✅ Funções passadas como props para components memorizados
- ✅ Funções usadas em dependências de outros hooks
- ✅ Funções que são criadas dentro de loops
- ❌ Funções simples que não são passadas como props
- ❌ Quando não há ganho de performance real

### 1.4 Evitando Criação de Objetos/Arrays em Render

```jsx
// ❌ Ruim: Novo objeto a cada render
function Component({ nome, idade }) {
  return <UserCard user={{ nome, idade }} />; // Novo objeto sempre!
}

// ✅ Bom: Objeto estável ou criar fora do render
function Component({ nome, idade }) {
  const user = useMemo(() => ({ nome, idade }), [nome, idade]);
  return <UserCard user={user} />;
}

// Ou melhor ainda: passar props individuais
function Component({ nome, idade }) {
  return <UserCard nome={nome} idade={idade} />;
}
```

---

## 2. Boas Práticas: Estrutura e Organização

### 2.1 Nomenclatura de Components

**✅ BOAS PRÁTICAS**:

```jsx
// ✅ PascalCase para components
function UserProfile() { }
function ProductCard() { }
function NavigationMenu() { }

// ✅ Nomes descritivos e específicos
function UserAvatar() { } // Não: Avatar (muito genérico)
function ShoppingCartItem() { } // Não: Item (muito genérico)

// ✅ Nomes que descrevem o propósito
function LoginForm() { } // Claro: é um formulário de login
function ErrorMessage() { } // Claro: é uma mensagem de erro
```

**❌ EVITE**:

```jsx
// ❌ camelCase (para variáveis, não components)
function userProfile() { }

// ❌ Nomes genéricos
function Component() { }
function Item() { }
function Box() { }

// ❌ Abreviações confusas
function UsrProf() { }
function ProdCard() { }
```

### 2.2 Organização de Arquivos

**Estrutura Recomendada**:

```
src/
├── components/
│   ├── Button/
│   │   ├── Button.jsx
│   │   ├── Button.css
│   │   └── index.js
│   ├── Card/
│   │   ├── Card.jsx
│   │   ├── CardHeader.jsx
│   │   ├── CardBody.jsx
│   │   ├── Card.css
│   │   └── index.js
│   └── shared/
│       ├── LoadingSpinner.jsx
│       └── ErrorMessage.jsx
├── pages/
│   ├── Home.jsx
│   └── About.jsx
└── App.jsx
```

**Padrão de Exportação**:

```jsx
// Button/Button.jsx
export function Button({ children, onClick }) {
  return <button onClick={onClick}>{children}</button>;
}

// Button/index.js
export { Button } from './Button';

// Uso
import { Button } from './components/Button';
```

### 2.3 Tamanho e Responsabilidade dos Components

**Regra Geral**: Um component deve fazer **uma coisa bem feita**.

**✅ BOM - Component Pequeno e Focado**:

```jsx
function UserAvatar({ user, size = 'medium' }) {
  return (
    <img 
      src={user.avatarUrl} 
      alt={user.name}
      className={`avatar avatar-${size}`}
    />
  );
}
```

**❌ RUIM - Component Fazendo Muitas Coisas**:

```jsx
function UserProfile({ user }) {
  // ❌ Muitas responsabilidades:
  // - Mostrar avatar
  // - Mostrar informações
  // - Gerenciar formulário
  // - Fazer chamada de API
  // - Gerenciar estado de edição
  
  const [isEditing, setIsEditing] = useState(false);
  const [formData, setFormData] = useState(user);
  
  useEffect(() => {
    fetchUserData(user.id);
  }, [user.id]);
  
  const handleSubmit = async () => {
    await updateUser(formData);
  };
  
  return (
    <div>
      {/* 100+ linhas de JSX... */}
    </div>
  );
}
```

**✅ BOM - Dividido em Components Menores**:

```jsx
function UserProfile({ user }) {
  return (
    <div>
      <UserAvatar user={user} />
      <UserInfo user={user} />
      <UserActions userId={user.id} />
    </div>
  );
}

function UserAvatar({ user }) { /* ... */ }
function UserInfo({ user }) { /* ... */ }
function UserActions({ userId }) { /* ... */ }
```

### 2.4 Props: Validação e Documentação

#### PropTypes (Desenvolvimento)

```jsx
import PropTypes from 'prop-types';

function UserCard({ nome, idade, email, ativo }) {
  return (
    <div>
      <h2>{nome}</h2>
      <p>Idade: {idade}</p>
      <p>Email: {email}</p>
      <p>Status: {ativo ? 'Ativo' : 'Inativo'}</p>
    </div>
  );
}

UserCard.propTypes = {
  nome: PropTypes.string.isRequired,
  idade: PropTypes.number.isRequired,
  email: PropTypes.string.isRequired,
  ativo: PropTypes.bool
};

UserCard.defaultProps = {
  ativo: false
};
```

#### TypeScript (Produção)

```tsx
interface UserCardProps {
  nome: string;
  idade: number;
  email: string;
  ativo?: boolean;
}

function UserCard({ nome, idade, email, ativo = false }: UserCardProps) {
  // ...
}
```

---

## 3. Padrões de Components: Quando Usar Cada Um

### 3.1 Presentational vs Container Components

**Presentational Components** (Componentes de Apresentação):
- Focam em **como** as coisas aparecem
- Recebem dados via props
- Não gerenciam estado de negócio
- São altamente reutilizáveis

```jsx
// ✅ Presentational
function ProductCard({ produto, onComprar }) {
  return (
    <div className="product-card">
      <img src={produto.imagem} alt={produto.nome} />
      <h3>{produto.nome}</h3>
      <p>R$ {produto.preco}</p>
      <button onClick={() => onComprar(produto.id)}>
        Comprar
      </button>
    </div>
  );
}
```

**Container Components** (Componentes de Container):
- Focam em **como** as coisas funcionam
- Gerenciam estado e lógica de negócio
- Fazem chamadas de API
- Passam dados para presentational components

```jsx
// ✅ Container
function ProductListContainer() {
  const [produtos, setProdutos] = useState([]);
  const [loading, setLoading] = useState(true);
  const [erro, setErro] = useState(null);

  useEffect(() => {
    fetch('/api/produtos')
      .then(res => res.json())
      .then(data => {
        setProdutos(data);
        setLoading(false);
      })
      .catch(err => {
        setErro(err.message);
        setLoading(false);
      });
  }, []);

  const handleComprar = (produtoId) => {
    // Lógica de compra
    console.log('Comprar produto', produtoId);
  };

  if (loading) return <LoadingSpinner />;
  if (erro) return <ErrorMessage message={erro} />;

  return (
    <div>
      {produtos.map(produto => (
        <ProductCard
          key={produto.id}
          produto={produto}
          onComprar={handleComprar}
        />
      ))}
    </div>
  );
}
```

**Nota**: Com hooks, essa distinção é menos rígida, mas ainda é um padrão útil para organização.

### 3.2 Controlled vs Uncontrolled Components

**Controlled Components** (Recomendado na maioria dos casos):
- React controla o valor através de state
- Valor sempre sincronizado com state
- Fácil de validar e manipular

```jsx
// ✅ Controlled
function LoginForm() {
  const [email, setEmail] = useState('');
  const [senha, setSenha] = useState('');

  const handleSubmit = (e) => {
    e.preventDefault();
    // email e senha sempre atualizados
    console.log({ email, senha });
  };

  return (
    <form onSubmit={handleSubmit}>
      <input
        type="email"
        value={email}
        onChange={(e) => setEmail(e.target.value)}
      />
      <input
        type="password"
        value={senha}
        onChange={(e) => setSenha(e.target.value)}
      />
      <button type="submit">Entrar</button>
    </form>
  );
}
```

**Uncontrolled Components**:
- DOM controla o valor
- Acessado via refs
- Útil para integração com bibliotecas não-React

```jsx
// ⚠️ Uncontrolled (use apenas quando necessário)
function LoginForm() {
  const emailRef = useRef(null);
  const senhaRef = useRef(null);

  const handleSubmit = (e) => {
    e.preventDefault();
    // Acessa valores via refs
    console.log({
      email: emailRef.current.value,
      senha: senhaRef.current.value
    });
  };

  return (
    <form onSubmit={handleSubmit}>
      <input type="email" ref={emailRef} />
      <input type="password" ref={senhaRef} />
      <button type="submit">Entrar</button>
    </form>
  );
}
```

**Quando usar Uncontrolled**:
- Integração com bibliotecas de terceiros
- Campos de arquivo (`<input type="file">`)
- Quando performance é crítica e você não precisa do valor em tempo real

---

## 4. Anti-padrões Comuns e Como Evitá-los

### 4.1 Modificar Props Diretamente

```jsx
// ❌ ERRADO - Nunca modifique props
function UserCard({ user }) {
  user.nome = "Novo Nome"; // NÃO FAÇA ISSO!
  return <div>{user.nome}</div>;
}

// ✅ CORRETO - Props são read-only
function UserCard({ user }) {
  return <div>{user.nome}</div>;
}
```

### 4.2 Usar Índice como Key em Listas

```jsx
// ❌ ERRADO - Índice como key
function Lista({ itens }) {
  return (
    <ul>
      {itens.map((item, index) => (
        <li key={index}>{item.nome}</li>
      ))}
    </ul>
  );
}

// ✅ CORRETO - ID único como key
function Lista({ itens }) {
  return (
    <ul>
      {itens.map((item) => (
        <li key={item.id}>{item.nome}</li>
      ))}
    </ul>
  );
}
```

**Por quê?**: Se a ordem dos itens mudar, React pode renderizar incorretamente ou ter problemas de performance.

### 4.3 Criar Funções Dentro do Render

```jsx
// ❌ ERRADO - Nova função a cada render
function Lista({ itens }) {
  return (
    <ul>
      {itens.map(item => (
        <Item 
          key={item.id}
          item={item}
          onClick={() => handleClick(item.id)} // Nova função sempre!
        />
      ))}
    </ul>
  );
}

// ✅ CORRETO - useCallback ou função estável
function Lista({ itens }) {
  const handleClick = useCallback((id) => {
    console.log('Clicou', id);
  }, []);

  return (
    <ul>
      {itens.map(item => (
        <Item 
          key={item.id}
          item={item}
          onClick={() => handleClick(item.id)}
        />
      ))}
    </ul>
  );
}
```

### 4.4 State em Componentes que Não Precisam

```jsx
// ❌ ERRADO - State desnecessário
function DisplayNome({ nome }) {
  const [displayNome, setDisplayNome] = useState(nome);
  
  return <div>{displayNome}</div>;
}

// ✅ CORRETO - Use props diretamente
function DisplayNome({ nome }) {
  return <div>{nome}</div>;
}
```

### 4.5 Efeitos Colaterais no Render

```jsx
// ❌ ERRADO - Efeito colateral no render
function Component({ dados }) {
  // NUNCA faça isso no render!
  localStorage.setItem('dados', JSON.stringify(dados));
  fetch('/api/dados', { method: 'POST', body: dados });
  
  return <div>{dados}</div>;
}

// ✅ CORRETO - Use useEffect
function Component({ dados }) {
  useEffect(() => {
    localStorage.setItem('dados', JSON.stringify(dados));
    fetch('/api/dados', { method: 'POST', body: dados });
  }, [dados]);
  
  return <div>{dados}</div>;
}
```

---

## 5. Otimização de Performance: Técnicas Avançadas

### 5.1 Code Splitting e Lazy Loading

```jsx
import { lazy, Suspense } from 'react';

// ✅ Lazy loading - carrega apenas quando necessário
const ComponentePesado = lazy(() => import('./ComponentePesado'));

function App() {
  return (
    <Suspense fallback={<div>Carregando...</div>}>
      <ComponentePesado />
    </Suspense>
  );
}
```

**Quando usar**:
- Componentes grandes que não são sempre visíveis
- Rotas de aplicação
- Bibliotecas pesadas

### 5.2 Virtualização de Listas

Para listas muito grandes (1000+ itens):

```jsx
import { FixedSizeList } from 'react-window';

function ListaGrande({ itens }) {
  return (
    <FixedSizeList
      height={600}
      itemCount={itens.length}
      itemSize={50}
      width="100%"
    >
      {({ index, style }) => (
        <div style={style}>
          {itens[index].nome}
        </div>
      )}
    </FixedSizeList>
  );
}
```

**Benefício**: Renderiza apenas os itens visíveis, melhorando drasticamente a performance.

### 5.3 Debounce e Throttle

```jsx
import { useState, useEffect } from 'react';

function Busca() {
  const [termo, setTermo] = useState('');
  const [resultados, setResultados] = useState([]);

  useEffect(() => {
    // Debounce: espera 500ms após parar de digitar
    const timer = setTimeout(() => {
      if (termo) {
        buscarResultados(termo).then(setResultados);
      }
    }, 500);

    return () => clearTimeout(timer);
  }, [termo]);

  return (
    <div>
      <input
        value={termo}
        onChange={(e) => setTermo(e.target.value)}
        placeholder="Buscar..."
      />
      {/* Resultados */}
    </div>
  );
}
```

---

## 6. Acessibilidade em Components

### 6.1 Semântica HTML

```jsx
// ❌ Ruim - divs genéricas
function Botao({ children, onClick }) {
  return <div onClick={onClick}>{children}</div>;
}

// ✅ Bom - elemento semântico
function Botao({ children, onClick }) {
  return <button onClick={onClick}>{children}</button>;
}
```

### 6.2 Atributos ARIA

```jsx
function Modal({ aberto, onFechar, children }) {
  return (
    <>
      {aberto && (
        <div
          role="dialog"
          aria-modal="true"
          aria-labelledby="modal-titulo"
        >
          <h2 id="modal-titulo">Título do Modal</h2>
          {children}
          <button onClick={onFechar} aria-label="Fechar modal">
            X
          </button>
        </div>
      )}
    </>
  );
}
```

### 6.3 Navegação por Teclado

```jsx
function Botao({ children, onClick }) {
  const handleKeyDown = (e) => {
    if (e.key === 'Enter' || e.key === ' ') {
      onClick();
    }
  };

  return (
    <button
      onClick={onClick}
      onKeyDown={handleKeyDown}
    >
      {children}
    </button>
  );
}
```

---

## 7. Testabilidade de Components

### 7.1 Components Testáveis

**✅ BOM - Fácil de Testar**:

```jsx
function Contador({ valorInicial = 0 }) {
  const [count, setCount] = useState(valorInicial);
  
  return (
    <div>
      <p data-testid="contador">{count}</p>
      <button 
        data-testid="incrementar"
        onClick={() => setCount(count + 1)}
      >
        Incrementar
      </button>
    </div>
  );
}
```

**Características de components testáveis**:
- Props claras e bem definidas
- Comportamento previsível
- Fácil de isolar
- `data-testid` para seletores estáveis

### 7.2 Evitando Lógica Complexa em Components

```jsx
// ❌ Ruim - lógica complexa misturada
function Component({ dados }) {
  const processados = dados
    .filter(d => d.ativo)
    .map(d => ({ ...d, nome: d.nome.toUpperCase() }))
    .sort((a, b) => a.nome.localeCompare(b.nome));
  
  return <div>{/* ... */}</div>;
}

// ✅ Bom - lógica separada, component focado
function processarDados(dados) {
  return dados
    .filter(d => d.ativo)
    .map(d => ({ ...d, nome: d.nome.toUpperCase() }))
    .sort((a, b) => a.nome.localeCompare(b.nome));
}

function Component({ dados }) {
  const processados = useMemo(() => processarDados(dados), [dados]);
  return <div>{/* ... */}</div>;
}
```

---

## 8. Checklist de Boas Práticas

Use este checklist ao criar components:

### Estrutura
- [ ] Component tem uma responsabilidade clara
- [ ] Nome do component é descritivo e em PascalCase
- [ ] Component não é muito grande (< 200 linhas geralmente)
- [ ] Lógica complexa está separada do component

### Props e State
- [ ] Props são validadas (PropTypes ou TypeScript)
- [ ] Props têm valores padrão quando apropriado
- [ ] State é usado apenas quando necessário
- [ ] Props não são modificadas diretamente

### Performance
- [ ] `React.memo` usado quando apropriado
- [ ] `useMemo` para cálculos caros
- [ ] `useCallback` para funções passadas como props
- [ ] Keys únicas e estáveis em listas
- [ ] Evita criar objetos/arrays novos em render

### Acessibilidade
- [ ] Elementos semânticos HTML usados
- [ ] Atributos ARIA quando necessário
- [ ] Navegação por teclado funcional
- [ ] Contraste de cores adequado

### Código Limpo
- [ ] Código é legível e bem formatado
- [ ] Comentários explicam "por quê", não "o quê"
- [ ] Funções têm nomes descritivos
- [ ] Sem código duplicado

---

## 9. Ferramentas de Análise

### 9.1 React DevTools Profiler

Use o Profiler para identificar components que re-renderizam desnecessariamente:

1. Abra React DevTools
2. Vá para a aba "Profiler"
3. Clique em "Record"
4. Interaja com sua aplicação
5. Pare a gravação
6. Analise quais components renderizaram e por quê

### 9.2 Console Warnings

React avisa sobre problemas comuns:
- Missing keys em listas
- Props modificadas
- Hooks usados incorretamente

**Sempre corrija os warnings!**

---

## 10. Conclusão: Desenvolvendo como Profissional

### Mentalidade Correta

1. **Pense em Performance desde o Início**: Não espere problemas para otimizar
2. **Componentes Pequenos e Focados**: Mais fácil de entender, testar e manter
3. **Reutilização Inteligente**: Não force reutilização onde não faz sentido
4. **Teste Seus Components**: Garanta que funcionam corretamente
5. **Acessibilidade é Essencial**: Não é opcional, é obrigatório

### Próximos Passos

Agora que você entende components, props, state e boas práticas, você está pronto para:
- Aprender hooks avançados (`useEffect`, `useContext`, etc.)
- Gerenciar estado global (Context API, Redux, etc.)
- Trabalhar com rotas e navegação
- Integrar com APIs
- Construir aplicações React completas

**Lembre-se**: Boas práticas não são regras rígidas. Use seu julgamento e adapte às necessidades do seu projeto. O importante é entender os princípios e aplicá-los de forma inteligente.

---

**Próximo Passo**: Complete os exercícios da seção anterior e aguarde o feedback e análise do seu desempenho antes de prosseguir para a próxima aula.



