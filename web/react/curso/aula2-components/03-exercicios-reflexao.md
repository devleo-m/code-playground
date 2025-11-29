# Aula 2 - Exercícios e Reflexão: Components

## Exercícios Práticos

### Exercício 1: Criando Seu Primeiro Component

**Objetivo**: Familiarizar-se com a criação básica de components.

**Tarefas**:

1. Crie um component chamado `Saudacao` que recebe uma prop `nome` e exibe "Olá, [nome]!"

2. Crie um component chamado `Card` que recebe props `titulo` e `descricao` e exibe um card com essas informações.

3. Use ambos os components no component `App` passando diferentes valores.

**Solução Esperada**:

```jsx
// 1. Component Saudacao
function Saudacao({ nome }) {
  return <h1>Olá, {nome}!</h1>;
}

// 2. Component Card
function Card({ titulo, descricao }) {
  return (
    <div className="card">
      <h2>{titulo}</h2>
      <p>{descricao}</p>
    </div>
  );
}

// 3. Usando no App
function App() {
  return (
    <div>
      <Saudacao nome="Maria" />
      <Saudacao nome="João" />
      
      <Card 
        titulo="React" 
        descricao="Biblioteca JavaScript para construir UIs" 
      />
      <Card 
        titulo="Vite" 
        descricao="Ferramenta de build super rápida" 
      />
    </div>
  );
}
```

**Questão de Reflexão**: 
- Por que faz sentido criar components separados em vez de escrever tudo no component `App`?
- Quais são as vantagens de ter `Saudacao` e `Card` como components separados?

**Resposta Esperada**:
- **Reutilização**: Você pode usar `Saudacao` várias vezes com nomes diferentes sem repetir código
- **Manutenibilidade**: Se precisar mudar como um card aparece, você muda em um só lugar
- **Organização**: Código mais limpo e fácil de entender
- **Testabilidade**: Componentes pequenos são mais fáceis de testar isoladamente

---

### Exercício 2: Trabalhando com Props

**Objetivo**: Entender como passar e usar props corretamente.

**Tarefas**:

1. Crie um component `Botao` que recebe props `texto`, `cor` e `onClick`.

2. Crie um component `ListaItens` que recebe uma prop `itens` (array de strings) e renderiza uma lista.

3. Crie um component `PerfilUsuario` que recebe props `nome`, `idade`, `email` e `hobbies` (array) e exibe essas informações.

**Solução Esperada**:

```jsx
// 1. Component Botao
function Botao({ texto, cor, onClick }) {
  return (
    <button 
      style={{ backgroundColor: cor }}
      onClick={onClick}
    >
      {texto}
    </button>
  );
}

// 2. Component ListaItens
function ListaItens({ itens }) {
  return (
    <ul>
      {itens.map((item, index) => (
        <li key={index}>{item}</li>
      ))}
    </ul>
  );
}

// 3. Component PerfilUsuario
function PerfilUsuario({ nome, idade, email, hobbies }) {
  return (
    <div>
      <h2>{nome}</h2>
      <p>Idade: {idade} anos</p>
      <p>Email: {email}</p>
      <h3>Hobbies:</h3>
      <ul>
        {hobbies.map((hobby, index) => (
          <li key={index}>{hobby}</li>
        ))}
      </ul>
    </div>
  );
}

// Uso
function App() {
  const handleClick = () => {
    alert('Botão clicado!');
  };

  return (
    <div>
      <Botao 
        texto="Clique aqui" 
        cor="blue" 
        onClick={handleClick} 
      />
      
      <ListaItens itens={["Item 1", "Item 2", "Item 3"]} />
      
      <PerfilUsuario
        nome="João Silva"
        idade={25}
        email="joao@email.com"
        hobbies={["Leitura", "Programação", "Música"]}
      />
    </div>
  );
}
```

**Questão de Reflexão**:
- Por que é importante usar `key` quando renderizamos listas? O que aconteceria se não usássemos?
- Por que passamos `onClick` como prop em vez de definir a função diretamente no component `Botao`?

**Resposta Esperada**:
- **Key**: React usa `key` para identificar quais itens mudaram, foram adicionados ou removidos. Sem `key`, React pode renderizar incorretamente ou ter problemas de performance. A `key` deve ser única e estável.
- **onClick como prop**: Permite que o component pai controle o comportamento do botão. Diferentes instâncias do `Botao` podem ter comportamentos diferentes, tornando o component mais flexível e reutilizável.

---

### Exercício 3: Introduzindo State

**Objetivo**: Entender como usar `useState` para gerenciar estado interno.

**Tarefas**:

1. Crie um component `Contador` que mantém um contador e permite incrementar e decrementar.

2. Crie um component `InputControlado` que mantém o valor do input em state e exibe o valor atualizado abaixo.

3. Crie um component `ListaTarefas` que permite adicionar e remover tarefas de uma lista.

**Solução Esperada**:

```jsx
import { useState } from 'react';

// 1. Component Contador
function Contador() {
  const [count, setCount] = useState(0);

  return (
    <div>
      <p>Contador: {count}</p>
      <button onClick={() => setCount(count + 1)}>Incrementar</button>
      <button onClick={() => setCount(count - 1)}>Decrementar</button>
      <button onClick={() => setCount(0)}>Resetar</button>
    </div>
  );
}

// 2. Component InputControlado
function InputControlado() {
  const [valor, setValor] = useState('');

  return (
    <div>
      <input
        type="text"
        value={valor}
        onChange={(e) => setValor(e.target.value)}
        placeholder="Digite algo..."
      />
      <p>Você digitou: {valor}</p>
      <p>Número de caracteres: {valor.length}</p>
    </div>
  );
}

// 3. Component ListaTarefas
function ListaTarefas() {
  const [tarefas, setTarefas] = useState([]);
  const [input, setInput] = useState('');

  const adicionarTarefa = () => {
    if (input.trim() !== '') {
      setTarefas([...tarefas, { id: Date.now(), texto: input }]);
      setInput('');
    }
  };

  const removerTarefa = (id) => {
    setTarefas(tarefas.filter(tarefa => tarefa.id !== id));
  };

  return (
    <div>
      <div>
        <input
          type="text"
          value={input}
          onChange={(e) => setInput(e.target.value)}
          placeholder="Nova tarefa"
          onKeyPress={(e) => e.key === 'Enter' && adicionarTarefa()}
        />
        <button onClick={adicionarTarefa}>Adicionar</button>
      </div>
      
      <ul>
        {tarefas.map(tarefa => (
          <li key={tarefa.id}>
            {tarefa.texto}
            <button onClick={() => removerTarefa(tarefa.id)}>Remover</button>
          </li>
        ))}
      </ul>
      
      {tarefas.length === 0 && <p>Nenhuma tarefa adicionada ainda.</p>}
    </div>
  );
}
```

**Questão de Reflexão**:
- Por que usamos `[...tarefas, novaTarefa]` em vez de `tarefas.push(novaTarefa)`?
- O que aconteceria se tentássemos modificar o array `tarefas` diretamente?

**Resposta Esperada**:
- **Spread operator**: Criamos um novo array em vez de modificar o existente. React compara referências - se você modificar o array original, a referência não muda e React pode não detectar a mudança, não re-renderizando o component.
- **Mutação direta**: Se fizéssemos `tarefas.push(novaTarefa)`, estaríamos modificando o array original. React pode não detectar essa mudança porque a referência do array não mudou, resultando em bugs onde a interface não atualiza mesmo que os dados tenham mudado.

---

### Exercício 4: Conditional Rendering

**Objetivo**: Praticar renderização condicional em diferentes cenários.

**Tarefas**:

1. Crie um component `MensagemLogin` que mostra "Bem-vindo!" se o usuário estiver logado, ou "Por favor, faça login" se não estiver.

2. Crie um component `ListaProdutos` que mostra uma lista de produtos ou uma mensagem "Nenhum produto disponível" se a lista estiver vazia.

3. Crie um component `Perfil` que mostra informações do usuário se existir, ou "Carregando..." se estiver carregando, ou "Usuário não encontrado" se não houver usuário.

**Solução Esperada**:

```jsx
// 1. Component MensagemLogin
function MensagemLogin({ estaLogado, nomeUsuario }) {
  if (estaLogado) {
    return <h1>Bem-vindo, {nomeUsuario}!</h1>;
  }
  
  return <h1>Por favor, faça login</h1>;
}

// 2. Component ListaProdutos
function ListaProdutos({ produtos }) {
  if (produtos.length === 0) {
    return <p>Nenhum produto disponível</p>;
  }

  return (
    <ul>
      {produtos.map(produto => (
        <li key={produto.id}>
          {produto.nome} - R$ {produto.preco}
        </li>
      ))}
    </ul>
  );
}

// 3. Component Perfil
function Perfil({ usuario, carregando }) {
  if (carregando) {
    return <div>Carregando...</div>;
  }

  if (!usuario) {
    return <div>Usuário não encontrado</div>;
  }

  return (
    <div>
      <h2>{usuario.nome}</h2>
      <p>Email: {usuario.email}</p>
      <p>Idade: {usuario.idade} anos</p>
    </div>
  );
}

// Uso
function App() {
  const [estaLogado, setEstaLogado] = useState(false);
  const [produtos, setProdutos] = useState([]);
  const [usuario, setUsuario] = useState(null);
  const [carregando, setCarregando] = useState(true);

  return (
    <div>
      <MensagemLogin 
        estaLogado={estaLogado} 
        nomeUsuario="João" 
      />
      
      <ListaProdutos produtos={produtos} />
      
      <Perfil usuario={usuario} carregando={carregando} />
    </div>
  );
}
```

**Questão de Reflexão**:
- Qual é a diferença entre usar `if/else` e usar operador ternário `? :` para conditional rendering? Quando usar cada um?
- Por que é importante verificar condições como `produtos.length === 0` antes de tentar renderizar a lista?

**Resposta Esperada**:
- **if/else vs ternário**: 
  - `if/else` é melhor para lógica mais complexa ou quando você quer retornar cedo (early return)
  - Operador ternário é melhor para condições simples inline no JSX
  - Use `if/else` quando a lógica é complexa ou quando você quer evitar aninhamento excessivo
- **Verificação de array vazio**: Sem verificar, você pode tentar fazer `map` em um array vazio, o que não quebra mas pode causar confusão. Além disso, você pode querer mostrar uma mensagem especial quando não há itens, melhorando a experiência do usuário.

---

### Exercício 5: Composition

**Objetivo**: Praticar composição de components.

**Tarefas**:

1. Crie components pequenos: `Botao`, `Titulo`, `Paragrafo`.

2. Use composition para criar um component `Card` que usa esses components menores.

3. Crie um component `Pagina` que usa múltiplos `Card` components.

**Solução Esperada**:

```jsx
// 1. Components pequenos
function Botao({ texto, onClick, cor = "blue" }) {
  return (
    <button 
      onClick={onClick}
      style={{ backgroundColor: cor }}
    >
      {texto}
    </button>
  );
}

function Titulo({ texto, nivel = 2 }) {
  const Tag = `h${nivel}`;
  return <Tag>{texto}</Tag>;
}

function Paragrafo({ children }) {
  return <p>{children}</p>;
}

// 2. Component Card usando composition
function Card({ titulo, conteudo, botaoTexto, onBotaoClick }) {
  return (
    <div className="card" style={{ 
      border: '1px solid #ccc', 
      padding: '20px', 
      margin: '10px',
      borderRadius: '8px'
    }}>
      <Titulo texto={titulo} nivel={2} />
      <Paragrafo>{conteudo}</Paragrafo>
      {botaoTexto && (
        <Botao 
          texto={botaoTexto} 
          onClick={onBotaoClick}
          cor="green"
        />
      )}
    </div>
  );
}

// 3. Component Pagina usando múltiplos Cards
function Pagina() {
  return (
    <div>
      <Titulo texto="Minha Página" nivel={1} />
      
      <Card
        titulo="React"
        conteudo="Biblioteca JavaScript para construir interfaces de usuário"
        botaoTexto="Aprender Mais"
        onBotaoClick={() => alert('Aprendendo React!')}
      />
      
      <Card
        titulo="Components"
        conteudo="Components são os blocos de construção do React"
        botaoTexto="Explorar"
        onBotaoClick={() => alert('Explorando Components!')}
      />
      
      <Card
        titulo="Composition"
        conteudo="Composition permite construir componentes maiores usando menores"
      />
    </div>
  );
}
```

**Questão de Reflexão**:
- Por que faz sentido criar components pequenos (`Botao`, `Titulo`, `Paragrafo`) em vez de escrever tudo diretamente no `Card`?
- Como a composition torna o código mais flexível e reutilizável?

**Resposta Esperada**:
- **Components pequenos**: 
  - Podem ser reutilizados em outros lugares além do `Card`
  - São mais fáceis de testar isoladamente
  - Facilitam manutenção - mudar o estilo de todos os botões em um só lugar
  - Tornam o código mais legível e organizado
- **Flexibilidade da composition**:
  - Você pode combinar components de diferentes formas
  - Fácil de modificar - trocar um `Botao` por outro tipo de botão sem mudar o `Card`
  - Permite criar variações facilmente - um `Card` com botão, outro sem
  - Segue o princípio de responsabilidade única - cada component faz uma coisa

---

### Exercício 6: Props.children

**Objetivo**: Entender e usar `props.children`.

**Tarefas**:

1. Crie um component `Container` que aceita `children` e aplica um estilo de container.

2. Crie um component `Modal` que tem um título fixo mas aceita conteúdo customizado via `children`.

3. Use esses components para criar diferentes layouts.

**Solução Esperada**:

```jsx
// 1. Component Container
function Container({ children, titulo }) {
  return (
    <div style={{ 
      maxWidth: '800px', 
      margin: '0 auto', 
      padding: '20px',
      border: '1px solid #ddd',
      borderRadius: '8px'
    }}>
      {titulo && <h2>{titulo}</h2>}
      {children}
    </div>
  );
}

// 2. Component Modal
function Modal({ children, onFechar, titulo }) {
  return (
    <div style={{
      position: 'fixed',
      top: '50%',
      left: '50%',
      transform: 'translate(-50%, -50%)',
      backgroundColor: 'white',
      padding: '20px',
      borderRadius: '8px',
      boxShadow: '0 4px 6px rgba(0,0,0,0.1)',
      zIndex: 1000
    }}>
      <div style={{ display: 'flex', justifyContent: 'space-between', marginBottom: '10px' }}>
        <h2>{titulo}</h2>
        <button onClick={onFechar}>X</button>
      </div>
      <div>
        {children}
      </div>
    </div>
  );
}

// 3. Uso
function App() {
  const [mostrarModal, setMostrarModal] = useState(false);

  return (
    <div>
      <Container titulo="Conteúdo Principal">
        <p>Este é o conteúdo dentro do container.</p>
        <button onClick={() => setMostrarModal(true)}>
          Abrir Modal
        </button>
      </Container>

      {mostrarModal && (
        <Modal 
          titulo="Meu Modal" 
          onFechar={() => setMostrarModal(false)}
        >
          <p>Este é o conteúdo customizado do modal.</p>
          <p>Você pode colocar qualquer coisa aqui!</p>
          <button onClick={() => setMostrarModal(false)}>
            Fechar
          </button>
        </Modal>
      )}
    </div>
  );
}
```

**Questão de Reflexão**:
- Por que `children` é uma prop especial? Qual é sua vantagem sobre passar conteúdo como prop normal?
- Em que situações `children` é mais útil que props específicas?

**Resposta Esperada**:
- **children como prop especial**: 
  - Permite passar JSX diretamente entre as tags do component
  - Sintaxe mais natural e legível
  - Não precisa criar props para cada tipo de conteúdo possível
  - Flexível - pode passar qualquer JSX como children
- **Quando usar children**:
  - Quando o conteúdo é variável e pode ser qualquer coisa
  - Quando você quer criar "wrappers" ou "containers" genéricos
  - Quando a estrutura é fixa mas o conteúdo interno muda
  - Componentes de layout (Container, Card, Modal, etc.)

---

## Exercícios de Reflexão Profunda

### Reflexão 1: Props vs State - Quando Usar Cada Um?

**Cenário**: Você está criando um component de formulário de cadastro de usuário.

**Pergunta**: 
- O nome do usuário deve ser uma prop ou state?
- A lista de erros de validação deve ser prop ou state?
- A função de submit deve ser prop ou state?

**Pense sobre**:
- De onde vêm esses dados?
- Quem controla esses dados?
- Esses dados mudam durante a vida do component?

**Sua resposta**:
(Escreva sua reflexão aqui)

**Resposta Esperada**:
- **Nome do usuário**: Depende do contexto. Se vem de um component pai (ex: editando usuário existente), é prop. Se o usuário está digitando no formulário, é state.
- **Lista de erros**: Geralmente é state, pois é gerada internamente pelo component durante a validação.
- **Função de submit**: É prop (callback), pois o component pai precisa saber quando o formulário foi submetido para processar os dados.

---

### Reflexão 2: Component Size - Quando Dividir um Component?

**Cenário**: Você tem um component `PerfilUsuario` que mostra:
- Foto do usuário
- Nome e informações básicas
- Lista de posts
- Lista de amigos
- Formulário para adicionar novo post

**Pergunta**: 
- Este component está muito grande? Deveria ser dividido?
- Se sim, como você dividiria? Quais seriam os novos components?
- Quais critérios você usaria para decidir?

**Pense sobre**:
- Responsabilidade única
- Reutilização
- Manutenibilidade
- Complexidade

**Sua resposta**:
(Escreva sua reflexão aqui)

**Resposta Esperada**:
- **Sim, está muito grande**: O component tem muitas responsabilidades diferentes.
- **Divisão sugerida**:
  - `FotoPerfil` - mostra a foto
  - `InfoUsuario` - nome e informações básicas
  - `ListaPosts` - lista de posts (pode ser reutilizado)
  - `ListaAmigos` - lista de amigos (pode ser reutilizado)
  - `FormularioPost` - formulário para novo post
  - `PerfilUsuario` - component principal que compõe todos os outros
- **Critérios**:
  - Cada component deve ter uma responsabilidade clara
  - Components que podem ser reutilizados devem ser separados
  - Se um component tem mais de ~100 linhas, considere dividir
  - Se você consegue dar um nome claro e específico ao component, provavelmente está no tamanho certo

---

### Reflexão 3: Performance e Re-renderização

**Cenário**: Você tem um component `ListaProdutos` que recebe uma lista de 1000 produtos e renderiza cada um.

**Pergunta**: 
- O que acontece quando o state de um produto individual muda?
- Todos os 1000 produtos re-renderizam?
- Isso é eficiente?
- Como você otimizaria isso?

**Pense sobre**:
- Como React decide o que re-renderizar
- Memoização
- Keys em listas
- Componentização

**Sua resposta**:
(Escreva sua reflexão aqui)

**Resposta Esperada**:
- **Re-renderização**: Se o state está no component pai (`ListaProdutos`), toda a lista re-renderiza quando qualquer produto muda.
- **Não é eficiente**: Re-renderizar 1000 componentes quando apenas 1 mudou é desperdício.
- **Otimizações**:
  - Mover state para o component individual (`ProdutoItem`)
  - Usar `React.memo` para evitar re-renderizações desnecessárias
  - Usar keys estáveis e únicas
  - Considerar virtualização para listas muito grandes
  - Separar components que mudam frequentemente dos que não mudam

---

### Reflexão 4: Composition vs Props Específicas

**Cenário**: Você está criando um component `Card` que pode ter:
- Título
- Conteúdo
- Botão de ação
- Imagem
- Badge de status

**Pergunta**: 
- Você criaria props específicas para cada parte (`titulo`, `conteudo`, `botao`, etc.)?
- Ou usaria composition com `children`?
- Ou uma combinação?
- Qual abordagem é mais flexível?

**Pense sobre**:
- Flexibilidade
- Facilidade de uso
- Manutenibilidade
- Casos de uso diferentes

**Sua resposta**:
(Escreva sua reflexão aqui)

**Resposta Esperada**:
- **Abordagem híbrida é geralmente melhor**:
  - Props específicas para partes comuns e bem definidas (`titulo`, `imagem`)
  - `children` para conteúdo variável e flexível
  - Props opcionais para funcionalidades extras (`botao`, `badge`)
- **Exemplo**:
  ```jsx
  <Card 
    titulo="Produto"
    imagem="/produto.jpg"
    badge="Novo"
  >
    <p>Descrição customizada aqui</p>
    <CustomContent />
  </Card>
  ```
- **Por quê**: Combina a simplicidade de props específicas com a flexibilidade de `children`.

---

### Reflexão 5: State Management - Quando State Vira Prop?

**Cenário**: Você tem um component `Contador` com seu próprio state. Agora você precisa de dois contadores que devem ser sincronizados (quando um muda, o outro também muda).

**Pergunta**: 
- Como você resolveria isso?
- O state deveria ser movido para o component pai?
- Quando faz sentido "elevar" state (lifting state up)?

**Pense sobre**:
- Compartilhamento de estado
- Single source of truth
- Fluxo de dados
- Complexidade

**Sua resposta**:
(Escreva sua reflexão aqui)

**Resposta Esperada**:
- **Lifting State Up**: Mover o state para o component pai comum e passar como props para os filhos.
- **Quando fazer**:
  - Quando múltiplos components precisam do mesmo estado
  - Quando você precisa sincronizar estado entre components
  - Quando o estado precisa ser acessado por um ancestor comum
- **Solução**:
  ```jsx
  function App() {
    const [count, setCount] = useState(0);
    
    return (
      <div>
        <Contador count={count} onIncrement={() => setCount(count + 1)} />
        <Contador count={count} onIncrement={() => setCount(count + 1)} />
      </div>
    );
  }
  ```
- **Alternativa**: Para casos mais complexos, considerar Context API ou bibliotecas de state management.

---

## Exercício Final: Construindo uma Aplicação Completa

### Tarefa:

Crie uma aplicação de "Lista de Tarefas" (Todo List) que demonstre todos os conceitos aprendidos:

**Requisitos**:

1. **Component `TodoItem`**: 
   - Recebe uma tarefa como prop
   - Mostra o texto da tarefa
   - Permite marcar como concluída
   - Permite remover

2. **Component `TodoList`**:
   - Gerencia a lista de tarefas (state)
   - Permite adicionar novas tarefas
   - Mostra lista de tarefas ou mensagem se vazia (conditional rendering)
   - Usa composition com `TodoItem`

3. **Component `App`**:
   - Component principal
   - Usa `TodoList`
   - Demonstra props, state, conditional rendering e composition

**Funcionalidades Extras** (opcional):
- Filtrar tarefas (todas, ativas, concluídas)
- Contador de tarefas
- Persistência no localStorage

**Solução Esperada**:

```jsx
import { useState } from 'react';

// Component TodoItem
function TodoItem({ tarefa, onToggle, onRemove }) {
  return (
    <li style={{ 
      textDecoration: tarefa.concluida ? 'line-through' : 'none',
      opacity: tarefa.concluida ? 0.6 : 1
    }}>
      <input
        type="checkbox"
        checked={tarefa.concluida}
        onChange={() => onToggle(tarefa.id)}
      />
      <span style={{ marginLeft: '10px' }}>{tarefa.texto}</span>
      <button 
        onClick={() => onRemove(tarefa.id)}
        style={{ marginLeft: '10px' }}
      >
        Remover
      </button>
    </li>
  );
}

// Component TodoList
function TodoList() {
  const [tarefas, setTarefas] = useState([]);
  const [input, setInput] = useState('');

  const adicionarTarefa = () => {
    if (input.trim() !== '') {
      setTarefas([...tarefas, {
        id: Date.now(),
        texto: input,
        concluida: false
      }]);
      setInput('');
    }
  };

  const toggleTarefa = (id) => {
    setTarefas(tarefas.map(tarefa =>
      tarefa.id === id
        ? { ...tarefa, concluida: !tarefa.concluida }
        : tarefa
    ));
  };

  const removerTarefa = (id) => {
    setTarefas(tarefas.filter(tarefa => tarefa.id !== id));
  };

  const tarefasAtivas = tarefas.filter(t => !t.concluida).length;
  const tarefasConcluidas = tarefas.filter(t => t.concluida).length;

  return (
    <div>
      <h1>Minha Lista de Tarefas</h1>
      
      <div>
        <input
          type="text"
          value={input}
          onChange={(e) => setInput(e.target.value)}
          placeholder="Nova tarefa"
          onKeyPress={(e) => e.key === 'Enter' && adicionarTarefa()}
        />
        <button onClick={adicionarTarefa}>Adicionar</button>
      </div>

      <div style={{ marginTop: '20px' }}>
        <p>Total: {tarefas.length} | Ativas: {tarefasAtivas} | Concluídas: {tarefasConcluidas}</p>
      </div>

      {tarefas.length === 0 ? (
        <p>Nenhuma tarefa adicionada ainda. Adicione uma tarefa acima!</p>
      ) : (
        <ul>
          {tarefas.map(tarefa => (
            <TodoItem
              key={tarefa.id}
              tarefa={tarefa}
              onToggle={toggleTarefa}
              onRemove={removerTarefa}
            />
          ))}
        </ul>
      )}
    </div>
  );
}

// Component App
function App() {
  return (
    <div style={{ maxWidth: '600px', margin: '50px auto', padding: '20px' }}>
      <TodoList />
    </div>
  );
}

export default App;
```

---

## Instruções para Resolução

1. **Crie um novo projeto Vite** (se ainda não tiver):
   ```bash
   npm create vite@latest meu-projeto -- --template react
   cd meu-projeto
   npm install
   ```

2. **Implemente cada exercício** em arquivos separados ou no mesmo arquivo

3. **Teste cada solução** executando `npm run dev`

4. **Responda as perguntas de reflexão** por escrito, pensando criticamente

5. **Experimente variações**: Modifique os exercícios, tente abordagens diferentes

6. **Anote suas dúvidas** para discussão posterior

---

## Dicas Importantes

- **Não tenha pressa**: Entenda cada conceito antes de avançar
- **Experimente**: Modifique o código, quebre coisas, aprenda com erros
- **Pratique naming**: Dê nomes claros e descritivos aos components
- **Pense em reutilização**: Sempre pergunte "Este component pode ser reutilizado?"
- **Mantenha components pequenos**: Se um component está fazendo muitas coisas, divida-o
- **Use props e state corretamente**: Props vêm de fora, state é interno

---

**Próximo Passo**: Após completar os exercícios e reflexões, aguarde o feedback e análise do seu desempenho antes de prosseguir para a seção de Performance e Boas Práticas.

