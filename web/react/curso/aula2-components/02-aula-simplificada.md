# Aula 2 - Simplificada: Entendendo Components no React

## Components: Os Blocos de Construção da Interface

Imagine que você está construindo uma casa. Você não constrói tudo de uma vez - você usa tijolos, portas, janelas, telhas. Cada peça tem uma função específica e você combina várias peças para construir a casa completa.

**Components no React funcionam exatamente assim!** Cada component é como um "tijolo" ou "porta" - uma peça reutilizável que você combina com outras para construir sua aplicação.

### Analogia do Mundo Real

Pense em um carro:
- **Motor** = um component
- **Rodas** = outro component
- **Volante** = outro component
- **Bancos** = outro component

O carro inteiro é composto por esses components. E o melhor: você pode usar o mesmo tipo de roda em vários carros diferentes (reutilização)!

---

## Functional Components: Funções que Retornam Interface

### O Que São Functional Components?

Functional components são simplesmente **funções JavaScript que retornam JSX** (código que parece HTML).

**Analogia**: Pense em uma função de cozinha:
- Você dá ingredientes (props) para a função
- A função processa e retorna um prato pronto (JSX)
- Cada vez que você chama a função com ingredientes diferentes, você recebe um prato diferente

**Exemplo simples:**
```jsx
// Esta função é um component
function Saudacao({ nome }) {
  return <h1>Olá, {nome}!</h1>;
}

// Uso: como chamar uma função normal
<Saudacao nome="Maria" />
// Resultado: <h1>Olá, Maria!</h1>
```

### Por Que "Functional"?

Porque são **funções puras** na maioria dos casos:
- Dados os mesmos ingredientes (props), sempre produzem o mesmo resultado
- Não modificam coisas de fora (sem efeitos colaterais)
- São previsíveis e fáceis de testar

**Analogia**: Como uma calculadora - sempre que você digita `2 + 2`, o resultado é `4`. Não importa quantas vezes você faça, sempre dá o mesmo resultado.

---

## JSX: HTML Dentro do JavaScript

### O Que é JSX?

JSX é como escrever HTML, mas dentro do JavaScript. É uma forma de descrever como a interface deve parecer.

**Analogia**: Pense em JSX como um **molde de bolo**:
- O molde (JSX) descreve a forma do bolo
- O JavaScript preenche o molde com os ingredientes (dados)
- O resultado final é o bolo (interface renderizada)

**Exemplo:**
```jsx
// Isso parece HTML, mas é JavaScript!
const elemento = (
  <div>
    <h1>Título</h1>
    <p>Parágrafo</p>
  </div>
);
```

### Por Que JSX Existe?

**Sem JSX** (difícil de ler):
```jsx
React.createElement('div', null,
  React.createElement('h1', null, 'Título'),
  React.createElement('p', null, 'Parágrafo')
);
```

**Com JSX** (fácil de ler):
```jsx
<div>
  <h1>Título</h1>
  <p>Parágrafo</p>
</div>
```

JSX torna o código muito mais fácil de entender, como se você estivesse escrevendo HTML normal!

### Regras Importantes do JSX

#### 1. Um Único Elemento Raiz

**Analogia**: Como uma caixa - você pode colocar várias coisas dentro, mas precisa de uma caixa para guardar tudo.

```jsx
// ❌ Ruim: duas coisas soltas
<h1>Título</h1>
<p>Texto</p>

// ✅ Bom: tudo dentro de uma "caixa"
<div>
  <h1>Título</h1>
  <p>Texto</p>
</div>
```

#### 2. Atributos em camelCase

**Analogia**: JavaScript tem suas próprias "regras de gramática". Assim como você não escreve português em inglês, você não escreve HTML puro em JSX.

```jsx
// HTML normal
<div class="container" onclick="funcao()">

// JSX (JavaScript)
<div className="container" onClick={funcao}>
```

#### 3. Expressões JavaScript Dentro de Chaves

**Analogia**: As chaves `{}` são como "janelas" que permitem o JavaScript "olhar para fora" e usar variáveis e cálculos.

```jsx
const nome = "João";
const idade = 25;

<div>
  <h1>Olá, {nome}!</h1>
  <p>Você tem {idade} anos</p>
  <p>Você nasceu em {2024 - idade}</p>
</div>
```

Dentro das chaves, você pode usar qualquer expressão JavaScript válida!

---

## Props: Dados que Vêm de Fora

### O Que São Props?

**Props** (propriedades) são como **argumentos de função**, mas para components. São dados que um component pai "passa" para um component filho.

**Analogia**: Pense em props como **ingredientes** que você dá para uma receita:
- A receita (component) é sempre a mesma
- Mas os ingredientes (props) mudam
- Com ingredientes diferentes, você obtém resultados diferentes

**Exemplo:**
```jsx
// Component que recebe props
function Card({ titulo, descricao }) {
  return (
    <div>
      <h2>{titulo}</h2>
      <p>{descricao}</p>
    </div>
  );
}

// Usando com props diferentes
<Card titulo="React" descricao="Biblioteca JavaScript" />
<Card titulo="Vite" descricao="Ferramenta de build" />
```

### Props São Read-Only (Somente Leitura)

**Analogia**: Props são como uma **carta registrada** - você pode ler, mas não pode modificar. Se você recebe uma carta, você não pode mudar o que está escrito nela.

```jsx
// ❌ ERRADO - tentar modificar props
function Componente({ nome }) {
  nome = "Novo Nome"; // NÃO FAÇA ISSO!
  return <div>{nome}</div>;
}

// ✅ CORRETO - apenas usar props
function Componente({ nome }) {
  return <div>{nome}</div>;
}
```

### Props.children: O Conteúdo Entre as Tags

**Analogia**: `children` é como uma **caixa surpresa** - você não sabe o que vai dentro até abrir, mas você pode colocar qualquer coisa.

```jsx
function Caixa({ children }) {
  return (
    <div className="caixa">
      {children}
    </div>
  );
}

// Uso: o que está entre as tags vira "children"
<Caixa>
  <h1>Título</h1>
  <p>Texto</p>
  <button>Clique</button>
</Caixa>
```

Aqui, `children` contém tudo que está entre `<Caixa>` e `</Caixa>`.

---

## State: A Memória do Component

### O Que é State?

**State** (estado) é a **memória interna** de um component. É como uma gaveta onde o component guarda informações que podem mudar.

**Analogia**: State é como a **memória de uma pessoa**:
- Uma pessoa lembra seu nome, idade, coisas que aconteceram
- Quando algo muda (você faz aniversário), a memória é atualizada
- A pessoa reage à mudança (comemora o aniversário!)

**Exemplo:**
```jsx
function Contador() {
  // useState cria uma "gaveta" chamada "count" com valor inicial 0
  const [count, setCount] = useState(0);

  return (
    <div>
      <p>Você clicou {count} vezes</p>
      {/* Quando clica, atualiza a "memória" (state) */}
      <button onClick={() => setCount(count + 1)}>
        Clique aqui
      </button>
    </div>
  );
}
```

### Como State Funciona?

1. **Inicialização**: `useState(0)` cria uma gaveta com valor `0`
2. **Leitura**: `count` lê o valor atual da gaveta
3. **Atualização**: `setCount(5)` atualiza o valor na gaveta para `5`
4. **Re-renderização**: Quando state muda, React re-desenha o component

**Analogia**: Como um **termômetro digital**:
- O termômetro tem uma memória (state) que guarda a temperatura
- Quando a temperatura muda, o display (interface) é atualizado automaticamente
- Você não precisa "dizer" para atualizar - acontece sozinho!

### State vs Props: Qual a Diferença?

**Analogia do Restaurante:**

- **Props** = **Ingredientes que o chef recebe**
  - Vêm de fora (do fornecedor/componente pai)
  - O chef não pode mudar os ingredientes
  - Diferentes ingredientes = pratos diferentes

- **State** = **Temperatura do fogão**
  - É controlado pelo próprio chef (componente)
  - O chef pode ajustar quando quiser
  - Quando muda, afeta como o prato é cozido

**Resumo:**
- **Props**: Dados que vêm de fora, não podem ser modificados
- **State**: Dados internos, podem ser modificados pelo próprio component

---

## Conditional Rendering: Mostrar Coisas Diferentes

### O Que é Conditional Rendering?

**Conditional Rendering** é mostrar coisas diferentes na tela dependendo de uma condição. É como um **semáforo** - mostra verde ou vermelho dependendo da situação.

**Analogia do Dia a Dia:**
- Se está chovendo → você pega o guarda-chuva
- Se não está chovendo → você não pega o guarda-chuva
- A ação depende da condição (está chovendo?)

**No React:**
```jsx
function Saudacao({ estaLogado }) {
  // Se está logado, mostra uma coisa
  if (estaLogado) {
    return <h1>Bem-vindo de volta!</h1>;
  }
  
  // Se não está logado, mostra outra coisa
  return <h1>Por favor, faça login</h1>;
}
```

### Métodos de Conditional Rendering

#### 1. Operador Ternário (If/Else Rápido)

**Analogia**: Como escolher entre duas opções no cardápio:
- "Se estiver com fome, pegue o prato grande, senão pegue o pequeno"

```jsx
function Cardapio({ estaComFome }) {
  return (
    <div>
      {estaComFome ? (
        <PratoGrande />
      ) : (
        <PratoPequeno />
      )}
    </div>
  );
}
```

#### 2. Operador && (Mostrar ou Não Mostrar)

**Analogia**: Como uma luz que só acende se algo estiver ligado:
- "Se a lâmpada estiver ligada, mostre a luz"

```jsx
function Lampada({ estaLigada }) {
  return (
    <div>
      {estaLigada && <div className="luz">💡</div>}
    </div>
  );
}
```

**Tradução**: "Se `estaLigada` for verdadeiro, mostre a luz. Se for falso, não mostre nada."

#### 3. Early Return (Retorno Antecipado)

**Analogia**: Como uma receita que diz "se não tiver ovos, pare aqui":
- Você verifica a condição primeiro
- Se não passar, retorna imediatamente
- Se passar, continua com o resto

```jsx
function Bolo({ temOvos }) {
  // Se não tem ovos, para aqui
  if (!temOvos) {
    return <div>Você precisa de ovos para fazer o bolo!</div>;
  }

  // Se chegou aqui, tem ovos - continua a receita
  return (
    <div>
      <h1>Fazendo o bolo...</h1>
      <p>Adicione os ovos...</p>
    </div>
  );
}
```

---

## Composition: Construir com Peças Menores

### O Que é Composition?

**Composition** (composição) é construir coisas grandes usando coisas pequenas. É como construir com **Lego** - você pega peças pequenas e monta algo maior.

**Analogia do Lego:**
- Você tem peças pequenas (components pequenos)
- Você combina várias peças para fazer algo maior (component maior)
- Você pode reutilizar as mesmas peças em construções diferentes

**Exemplo:**
```jsx
// Peças pequenas (components pequenos)
function Botao({ texto }) {
  return <button>{texto}</button>;
}

function Titulo({ texto }) {
  return <h2>{texto}</h2>;
}

// Construção maior (usando as peças pequenas)
function Card({ titulo, botaoTexto }) {
  return (
    <div className="card">
      <Titulo texto={titulo} />
      <Botao texto={botaoTexto} />
    </div>
  );
}
```

### Por Que Composition é Melhor que Herança?

**Analogia da Cozinha:**

**Herança** (não recomendado no React):
- "Esta receita herda da receita base"
- Se você mudar a receita base, todas as receitas que herdam dela mudam
- Rígido e difícil de modificar

**Composition** (recomendado no React):
- "Esta receita usa ingredientes de outras receitas"
- Você combina receitas diferentes como quiser
- Flexível e fácil de modificar

**Exemplo de Composition:**
```jsx
// Componentes pequenos e reutilizáveis
function Container({ children }) {
  return <div className="container">{children}</div>;
}

function Titulo({ texto }) {
  return <h1>{texto}</h1>;
}

function Botao({ texto, onClick }) {
  return <button onClick={onClick}>{texto}</button>;
}

// Combinando para criar algo maior
function Pagina() {
  return (
    <Container>
      <Titulo texto="Minha Página" />
      <Botao texto="Clique aqui" onClick={() => alert('Oi!')} />
    </Container>
  );
}
```

### Composition com children

**Analogia da Caixa Mágica:**
- Você tem uma caixa (component Container)
- Você pode colocar qualquer coisa dentro da caixa (children)
- A caixa não precisa saber o que está dentro - só precisa mostrar

```jsx
function Caixa({ children, titulo }) {
  return (
    <div className="caixa">
      <h2>{titulo}</h2>
      <div className="conteudo">
        {children} {/* O que você colocar aqui aparece */}
      </div>
    </div>
  );
}

// Uso: você decide o que vai dentro
<Caixa titulo="Minha Caixa">
  <p>Qualquer coisa pode ir aqui!</p>
  <button>Botão</button>
  <img src="foto.jpg" alt="Foto" />
</Caixa>
```

---

## Resumindo: Componentes São Como...

### Components = Tijolos de Lego
- Cada peça tem uma função
- Você combina peças para fazer algo maior
- Você pode reutilizar as mesmas peças

### Props = Ingredientes
- Vêm de fora (do componente pai)
- Não podem ser modificados
- Diferentes ingredientes = resultado diferente

### State = Memória/Gaveta
- Guarda informações que podem mudar
- Quando muda, a interface atualiza automaticamente
- Cada component tem sua própria "gaveta"

### JSX = Molde
- Descreve como a interface deve parecer
- Parece HTML, mas é JavaScript
- Permite usar variáveis e expressões

### Conditional Rendering = Semáforo
- Mostra coisas diferentes dependendo da condição
- Se/então - se verdadeiro mostra uma coisa, se falso mostra outra

### Composition = Construção com Lego
- Peças pequenas formam coisas maiores
- Flexível e reutilizável
- Melhor que herança

---

## Exemplo Completo: Construindo uma Interface

Vamos ver como tudo se encaixa em um exemplo prático:

```jsx
// 1. Component pequeno: Botão
function Botao({ texto, onClick }) {
  return <button onClick={onClick}>{texto}</button>;
}

// 2. Component pequeno: Card
function Card({ titulo, children }) {
  return (
    <div className="card">
      <h2>{titulo}</h2>
      {children}
    </div>
  );
}

// 3. Component maior: Lista de Produtos (usa composition)
function ListaProdutos({ produtos, mostrarVazia }) {
  // Conditional rendering
  if (produtos.length === 0) {
    return mostrarVazia ? <p>Nenhum produto</p> : null;
  }

  return (
    <div>
      {produtos.map(produto => (
        <Card key={produto.id} titulo={produto.nome}>
          <p>Preço: R$ {produto.preco}</p>
          <Botao 
            texto="Comprar" 
            onClick={() => console.log('Comprar', produto.nome)} 
          />
        </Card>
      ))}
    </div>
  );
}

// 4. Component principal com state
function App() {
  // State: memória do component
  const [produtos, setProdutos] = useState([
    { id: 1, nome: "Notebook", preco: 2500 },
    { id: 2, nome: "Mouse", preco: 50 }
  ]);

  return (
    <div>
      <h1>Minha Loja</h1>
      {/* Props: passando dados para o component filho */}
      <ListaProdutos 
        produtos={produtos} 
        mostrarVazia={true} 
      />
    </div>
  );
}
```

**O que está acontecendo aqui:**
1. **Components pequenos** (`Botao`, `Card`) são reutilizáveis
2. **Composition**: `ListaProdutos` usa `Card` e `Botao`
3. **Props**: Dados são passados de `App` para `ListaProdutos`
4. **State**: `App` gerencia a lista de produtos
5. **Conditional Rendering**: Mostra mensagem se não houver produtos
6. **JSX**: Tudo é escrito em JSX

---

## Dicas Importantes

### 1. Comece Pequeno
- Crie components pequenos e simples primeiro
- Combine-os para fazer coisas maiores
- Não tente fazer tudo de uma vez

### 2. Um Component, Uma Responsabilidade
- Cada component deve fazer uma coisa bem
- Se um component está fazendo muitas coisas, divida-o

### 3. Props vs State
- Se os dados vêm de fora → use **Props**
- Se os dados são internos e mudam → use **State**

### 4. Composition é Seu Amigo
- Sempre prefira composition sobre herança
- Components pequenos são mais fáceis de entender e testar

### 5. Pratique
- A melhor forma de aprender é praticando
- Crie components, experimente, quebre coisas, aprenda!

---

## Conclusão Simplificada

**Components** são como peças de Lego - você constrói coisas grandes usando peças pequenas.

**Props** são ingredientes que vêm de fora - você não pode mudá-los, apenas usá-los.

**State** é a memória do component - pode mudar e quando muda, a interface atualiza.

**JSX** é como escrever HTML, mas dentro do JavaScript.

**Conditional Rendering** é mostrar coisas diferentes dependendo de condições.

**Composition** é construir coisas grandes combinando coisas pequenas.

**Lembre-se**: Não precisa decorar tudo de uma vez. O importante é entender os **conceitos**. A sintaxe você aprende praticando!

**Próximo Passo**: Vamos praticar com exercícios para fixar esses conceitos!

