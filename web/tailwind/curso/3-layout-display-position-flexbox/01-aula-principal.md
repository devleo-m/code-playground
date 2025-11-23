# Aula 3: Layout com Tailwind - Display, Position e Flexbox - Conteúdo Principal

## 📖 Introdução

Agora que você domina os fundamentos das classes utilitárias do Tailwind, vamos explorar uma das áreas mais importantes: **controle de layout**. Nesta aula, você aprenderá a controlar como os elementos são exibidos, posicionados e organizados usando as utilities do Tailwind.

Você já conhece essas propriedades em CSS puro:
- `display` (block, inline, flex, grid, none)
- `position` (static, relative, absolute, fixed, sticky)
- Flexbox completo

Agora vamos ver como o Tailwind mapeia essas propriedades em classes utilitárias simples e poderosas.

**Nesta aula você aprenderá:**
- Display utilities (block, inline, flex, grid, hidden)
- Position utilities (static, relative, absolute, fixed, sticky)
- Propriedades de posicionamento (top, right, bottom, left)
- Z-index utilities
- Flexbox completo com Tailwind
- Criar layouts comuns usando Flexbox

---

## 🎨 Display Utilities

O `display` é uma das propriedades CSS mais fundamentais. Ele determina como um elemento é renderizado e como interage com outros elementos no fluxo do documento.

### Display Básico

#### Block

**Classe Tailwind:** `block`

```html
<span class="block">Este span agora é um elemento block</span>
```

**Equivalente CSS:**
```css
span {
  display: block;
}
```

**Comportamento:**
- Ocupa toda a largura disponível
- Quebra linha antes e depois
- Aceita width, height, padding, margin em todos os lados

#### Inline

**Classe Tailwind:** `inline`

```html
<div class="inline">Este div agora é inline</div>
```

**Equivalente CSS:**
```css
div {
  display: inline;
}
```

**Comportamento:**
- Ocupa apenas o espaço necessário
- Não quebra linha
- Não aceita width, height, margin-top, margin-bottom

#### Inline-Block

**Classe Tailwind:** `inline-block`

```html
<div class="inline-block w-32 h-32 bg-blue-500"></div>
```

**Equivalente CSS:**
```css
div {
  display: inline-block;
  width: 8rem;
  height: 8rem;
  background-color: rgb(59 130 246);
}
```

**Comportamento:**
- Combina características de block e inline
- Não quebra linha (como inline)
- Aceita width e height (como block)

#### Flex

**Classe Tailwind:** `flex`

```html
<div class="flex">
  <div>Item 1</div>
  <div>Item 2</div>
</div>
```

**Equivalente CSS:**
```css
div {
  display: flex;
}
```

**Comportamento:**
- Cria um container flexível
- Filhos diretos se tornam flex items
- Permite controle de alinhamento e distribuição

#### Grid

**Classe Tailwind:** `grid`

```html
<div class="grid grid-cols-3 gap-4">
  <div>Item 1</div>
  <div>Item 2</div>
  <div>Item 3</div>
</div>
```

**Equivalente CSS:**
```css
div {
  display: grid;
  grid-template-columns: repeat(3, minmax(0, 1fr));
  gap: 1rem;
}
```

**Nota:** Grid será abordado em detalhes na próxima aula.

#### Hidden

**Classe Tailwind:** `hidden`

```html
<div class="hidden">Este elemento está oculto</div>
```

**Equivalente CSS:**
```css
div {
  display: none;
}
```

**Comportamento:**
- Remove o elemento completamente do fluxo do documento
- Não ocupa espaço
- Não é renderizado visualmente

**Diferença importante:**
- `hidden` = `display: none` (remove do fluxo)
- `invisible` = `visibility: hidden` (mantém espaço, mas invisível)
- `opacity-0` = `opacity: 0` (mantém espaço, totalmente transparente)

---

## 📍 Position Utilities

O `position` controla como um elemento é posicionado no documento. Você já conhece essas propriedades em CSS, agora vamos ver as classes Tailwind.

### Position Básico

#### Static (Padrão)

**Classe Tailwind:** `static`

```html
<div class="static">Posição estática (padrão)</div>
```

**Equivalente CSS:**
```css
div {
  position: static;
}
```

**Comportamento:**
- Posição normal no fluxo do documento
- `top`, `right`, `bottom`, `left` e `z-index` não têm efeito
- Geralmente não precisa ser especificado (é o padrão)

#### Relative

**Classe Tailwind:** `relative`

```html
<div class="relative top-4 left-4">
  Posicionado relativamente à sua posição original
</div>
```

**Equivalente CSS:**
```css
div {
  position: relative;
  top: 1rem;
  left: 1rem;
}
```

**Comportamento:**
- Mantém seu espaço no fluxo normal
- Pode ser deslocado usando `top`, `right`, `bottom`, `left`
- Serve como referência para elementos `absolute` filhos

#### Absolute

**Classe Tailwind:** `absolute`

```html
<div class="relative">
  <div class="absolute top-0 right-0">Posicionado absolutamente</div>
</div>
```

**Equivalente CSS:**
```css
div {
  position: relative;
}

div > div {
  position: absolute;
  top: 0;
  right: 0;
}
```

**Comportamento:**
- Removido do fluxo normal do documento
- Posicionado relativamente ao ancestral mais próximo com `position` diferente de `static`
- Não ocupa espaço no layout
- Pode sobrepor outros elementos

#### Fixed

**Classe Tailwind:** `fixed`

```html
<div class="fixed top-0 left-0 w-full bg-white shadow-lg">
  Header fixo no topo
</div>
```

**Equivalente CSS:**
```css
div {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  background-color: white;
  box-shadow: 0 10px 15px -3px rgb(0 0 0 / 0.1);
}
```

**Comportamento:**
- Posicionado relativamente à viewport (janela do navegador)
- Permanece fixo mesmo ao rolar a página
- Não ocupa espaço no fluxo do documento
- Comum para headers, sidebars, modais

#### Sticky

**Classe Tailwind:** `sticky`

```html
<div class="sticky top-0 bg-white z-10">
  Este elemento "gruda" quando você rola a página
</div>
```

**Equivalente CSS:**
```css
div {
  position: sticky;
  top: 0;
  background-color: white;
  z-index: 10;
}
```

**Comportamento:**
- Combina `relative` e `fixed`
- Mantém seu espaço no fluxo até atingir o threshold (ex: `top-0`)
- Quando atinge o threshold, "gruda" como `fixed`
- Útil para headers de tabelas, navegação, etc.

---

## 🎯 Propriedades de Posicionamento

Quando você usa `relative`, `absolute`, `fixed` ou `sticky`, pode controlar a posição com as seguintes utilities:

### Top, Right, Bottom, Left

**Sintaxe:** `{propriedade}-{valor}`

```html
<!-- Posicionamento absoluto no canto superior direito -->
<div class="relative h-64">
  <div class="absolute top-4 right-4 bg-blue-500 p-4 rounded">
    Canto superior direito
  </div>
</div>
```

**Equivalente CSS:**
```css
div {
  position: relative;
  height: 16rem;
}

div > div {
  position: absolute;
  top: 1rem;
  right: 1rem;
  background-color: rgb(59 130 246);
  padding: 1rem;
  border-radius: 0.25rem;
}
```

**Valores disponíveis:**
- `0` = `0`
- `1` = `0.25rem` (4px)
- `2` = `0.5rem` (8px)
- `4` = `1rem` (16px)
- `8` = `2rem` (32px)
- E assim por diante...

**Valores negativos:**
- `-top-4` = `top: -1rem`
- `-right-2` = `right: -0.5rem`

**Valores percentuais:**
- `top-1/2` = `top: 50%`
- `left-1/3` = `left: 33.333333%`

**Valores completos:**
- `inset-0` = `top: 0; right: 0; bottom: 0; left: 0;`
- `inset-x-0` = `left: 0; right: 0;`
- `inset-y-0` = `top: 0; bottom: 0;`

### Exemplo Prático: Overlay Modal

```html
<!-- Overlay de fundo -->
<div class="fixed inset-0 bg-black bg-opacity-50 z-40">
  <!-- Modal centralizado -->
  <div class="absolute top-1/2 left-1/2 transform -translate-x-1/2 -translate-y-1/2 
              bg-white rounded-lg p-6 w-96 z-50">
    <h2 class="text-2xl font-bold mb-4">Modal</h2>
    <p>Conteúdo do modal</p>
  </div>
</div>
```

---

## 📚 Z-Index Utilities

O `z-index` controla a ordem de empilhamento de elementos posicionados. Valores maiores aparecem na frente.

**Sintaxe:** `z-{valor}`

```html
<div class="relative z-10">Camada 10</div>
<div class="relative z-20">Camada 20 (aparece na frente)</div>
<div class="relative z-30">Camada 30 (aparece na frente de todos)</div>
```

**Equivalente CSS:**
```css
div:nth-child(1) {
  position: relative;
  z-index: 10;
}

div:nth-child(2) {
  position: relative;
  z-index: 20;
}

div:nth-child(3) {
  position: relative;
  z-index: 30;
}
```

**Valores padrão do Tailwind:**
- `z-0` = `z-index: 0`
- `z-10` = `z-index: 10`
- `z-20` = `z-index: 20`
- `z-30` = `z-index: 30`
- `z-40` = `z-index: 40`
- `z-50` = `z-index: 50`
- `z-auto` = `z-index: auto`

**Valores negativos:**
- `z-[-1]` = `z-index: -1` (arbitrary value)

**Casos de uso comuns:**
- `z-10`: Dropdowns, tooltips
- `z-20`: Sticky headers
- `z-30`: Modals, overlays
- `z-40`: Popovers
- `z-50`: Notificações, toasts

---

## 🎪 Flexbox com Tailwind

Flexbox é uma das ferramentas mais poderosas para layout. O Tailwind fornece utilities completas para todas as propriedades do Flexbox.

### Container Flex

**Classe:** `flex`

```html
<div class="flex">
  <div>Item 1</div>
  <div>Item 2</div>
  <div>Item 3</div>
</div>
```

**Equivalente CSS:**
```css
div {
  display: flex;
}
```

### Direção do Flex

#### Row (Padrão)

**Classe:** `flex-row`

```html
<div class="flex flex-row">
  <div>Item 1</div>
  <div>Item 2</div>
</div>
```

**Equivalente CSS:**
```css
div {
  display: flex;
  flex-direction: row;
}
```

**Comportamento:**
- Itens dispostos horizontalmente (da esquerda para direita)
- É o padrão quando você usa apenas `flex`

#### Row Reverse

**Classe:** `flex-row-reverse`

```html
<div class="flex flex-row-reverse">
  <div>Item 1</div>
  <div>Item 2</div>
</div>
```

**Equivalente CSS:**
```css
div {
  display: flex;
  flex-direction: row-reverse;
}
```

**Comportamento:**
- Itens dispostos horizontalmente (da direita para esquerda)
- Ordem visual invertida

#### Column

**Classe:** `flex-col`

```html
<div class="flex flex-col">
  <div>Item 1</div>
  <div>Item 2</div>
</div>
```

**Equivalente CSS:**
```css
div {
  display: flex;
  flex-direction: column;
}
```

**Comportamento:**
- Itens dispostos verticalmente (de cima para baixo)
- Útil para layouts verticais, cards, formulários

#### Column Reverse

**Classe:** `flex-col-reverse`

```html
<div class="flex flex-col-reverse">
  <div>Item 1</div>
  <div>Item 2</div>
</div>
```

**Equivalente CSS:**
```css
div {
  display: flex;
  flex-direction: column-reverse;
}
```

**Comportamento:**
- Itens dispostos verticalmente (de baixo para cima)

---

### Alinhamento no Eixo Principal (Justify Content)

Controla como os itens são distribuídos ao longo do **eixo principal** (horizontal em `flex-row`, vertical em `flex-col`).

#### Flex Start (Padrão)

**Classe:** `justify-start`

```html
<div class="flex justify-start">
  <div>Item 1</div>
  <div>Item 2</div>
</div>
```

**Equivalente CSS:**
```css
div {
  display: flex;
  justify-content: flex-start;
}
```

**Comportamento:**
- Itens alinhados no início do eixo principal

#### Flex End

**Classe:** `justify-end`

```html
<div class="flex justify-end">
  <div>Item 1</div>
  <div>Item 2</div>
</div>
```

**Equivalente CSS:**
```css
div {
  display: flex;
  justify-content: flex-end;
}
```

#### Center

**Classe:** `justify-center`

```html
<div class="flex justify-center">
  <div>Item 1</div>
  <div>Item 2</div>
</div>
```

**Equivalente CSS:**
```css
div {
  display: flex;
  justify-content: center;
}
```

**Comportamento:**
- Itens centralizados no eixo principal
- Muito usado para centralizar conteúdo horizontalmente

#### Space Between

**Classe:** `justify-between`

```html
<div class="flex justify-between">
  <div>Item 1</div>
  <div>Item 2</div>
  <div>Item 3</div>
</div>
```

**Equivalente CSS:**
```css
div {
  display: flex;
  justify-content: space-between;
}
```

**Comportamento:**
- Primeiro item no início, último no fim
- Espaço igual entre os itens do meio
- Útil para headers com logo à esquerda e menu à direita

#### Space Around

**Classe:** `justify-around`

```html
<div class="flex justify-around">
  <div>Item 1</div>
  <div>Item 2</div>
  <div>Item 3</div>
</div>
```

**Equivalente CSS:**
```css
div {
  display: flex;
  justify-content: space-around;
}
```

**Comportamento:**
- Espaço igual ao redor de cada item
- Metade do espaço nas extremidades

#### Space Evenly

**Classe:** `justify-evenly`

```html
<div class="flex justify-evenly">
  <div>Item 1</div>
  <div>Item 2</div>
  <div>Item 3</div>
</div>
```

**Equivalente CSS:**
```css
div {
  display: flex;
  justify-content: space-evenly;
}
```

**Comportamento:**
- Espaço igual entre todos os itens, incluindo extremidades

---

### Alinhamento no Eixo Cruzado (Align Items)

Controla como os itens são alinhados no **eixo cruzado** (vertical em `flex-row`, horizontal em `flex-col`).

#### Stretch (Padrão)

**Classe:** `items-stretch`

```html
<div class="flex items-stretch h-32">
  <div class="bg-blue-500">Item 1</div>
  <div class="bg-red-500">Item 2</div>
</div>
```

**Equivalente CSS:**
```css
div {
  display: flex;
  align-items: stretch;
  height: 8rem;
}
```

**Comportamento:**
- Itens esticam para preencher a altura do container

#### Flex Start

**Classe:** `items-start`

```html
<div class="flex items-start h-32">
  <div class="bg-blue-500">Item 1</div>
  <div class="bg-red-500">Item 2</div>
</div>
```

**Equivalente CSS:**
```css
div {
  display: flex;
  align-items: flex-start;
  height: 8rem;
}
```

**Comportamento:**
- Itens alinhados no início do eixo cruzado

#### Flex End

**Classe:** `items-end`

```html
<div class="flex items-end h-32">
  <div class="bg-blue-500">Item 1</div>
  <div class="bg-red-500">Item 2</div>
</div>
```

**Equivalente CSS:**
```css
div {
  display: flex;
  align-items: flex-end;
  height: 8rem;
}
```

#### Center

**Classe:** `items-center`

```html
<div class="flex items-center h-32">
  <div class="bg-blue-500">Item 1</div>
  <div class="bg-red-500">Item 2</div>
</div>
```

**Equivalente CSS:**
```css
div {
  display: flex;
  align-items: center;
  height: 8rem;
}
```

**Comportamento:**
- Itens centralizados no eixo cruzado
- Muito usado para centralizar verticalmente

#### Baseline

**Classe:** `items-baseline`

```html
<div class="flex items-baseline">
  <div class="text-4xl">Texto Grande</div>
  <div class="text-sm">Texto Pequeno</div>
</div>
```

**Equivalente CSS:**
```css
div {
  display: flex;
  align-items: baseline;
}
```

**Comportamento:**
- Itens alinhados pela linha de base do texto
- Útil para alinhar textos de tamanhos diferentes

---

### Centralização Completa

A combinação mais comum: centralizar horizontal e verticalmente.

```html
<div class="flex items-center justify-center h-screen">
  <div>Conteúdo centralizado</div>
</div>
```

**Equivalente CSS:**
```css
div {
  display: flex;
  align-items: center;
  justify-content: center;
  height: 100vh;
}
```

---

### Wrap (Quebra de Linha)

Controla se os itens podem quebrar linha quando não cabem no container.

#### No Wrap (Padrão)

**Classe:** `flex-nowrap`

```html
<div class="flex flex-nowrap">
  <div class="w-64">Item 1</div>
  <div class="w-64">Item 2</div>
  <div class="w-64">Item 3</div>
</div>
```

**Equivalente CSS:**
```css
div {
  display: flex;
  flex-wrap: nowrap;
}
```

**Comportamento:**
- Itens não quebram linha
- Podem ultrapassar o container

#### Wrap

**Classe:** `flex-wrap`

```html
<div class="flex flex-wrap">
  <div class="w-64">Item 1</div>
  <div class="w-64">Item 2</div>
  <div class="w-64">Item 3</div>
</div>
```

**Equivalente CSS:**
```css
div {
  display: flex;
  flex-wrap: wrap;
}
```

**Comportamento:**
- Itens quebram linha quando necessário
- Útil para grids responsivos, cards

#### Wrap Reverse

**Classe:** `flex-wrap-reverse`

```html
<div class="flex flex-wrap-reverse">
  <div class="w-64">Item 1</div>
  <div class="w-64">Item 2</div>
  <div class="w-64">Item 3</div>
</div>
```

**Equivalente CSS:**
```css
div {
  display: flex;
  flex-wrap: wrap-reverse;
}
```

**Comportamento:**
- Itens quebram linha, mas em ordem reversa

---

### Propriedades de Flex Items

#### Flex Grow

Controla quanto um item pode crescer em relação aos outros.

**Classe:** `flex-grow` ou `grow`

```html
<div class="flex">
  <div class="flex-grow">Cresce para preencher espaço</div>
  <div class="w-32">Largura fixa</div>
</div>
```

**Equivalente CSS:**
```css
div {
  display: flex;
}

div > div:first-child {
  flex-grow: 1;
}

div > div:last-child {
  width: 8rem;
}
```

**Valores:**
- `grow` = `flex-grow: 1`
- `grow-0` = `flex-grow: 0` (padrão)

#### Flex Shrink

Controla quanto um item pode encolher.

**Classe:** `flex-shrink` ou `shrink`

```html
<div class="flex">
  <div class="w-64 shrink-0">Não encolhe</div>
  <div class="flex-grow">Pode encolher</div>
</div>
```

**Equivalente CSS:**
```css
div {
  display: flex;
}

div > div:first-child {
  width: 16rem;
  flex-shrink: 0;
}

div > div:last-child {
  flex-grow: 1;
}
```

**Valores:**
- `shrink` = `flex-shrink: 1` (padrão)
- `shrink-0` = `flex-shrink: 0` (não encolhe)

#### Flex Basis

Define o tamanho inicial de um item antes do espaço ser distribuído.

**Classe:** `flex-basis-{valor}` ou usar `w-{valor}`

```html
<div class="flex">
  <div class="basis-1/4">25% da largura</div>
  <div class="basis-3/4">75% da largura</div>
</div>
```

**Equivalente CSS:**
```css
div {
  display: flex;
}

div > div:first-child {
  flex-basis: 25%;
}

div > div:last-child {
  flex-basis: 75%;
}
```

#### Flex (Shorthand)

Combina `flex-grow`, `flex-shrink` e `flex-basis`.

**Classes:**
- `flex-1` = `flex: 1 1 0%` (cresce e encolhe, base 0)
- `flex-auto` = `flex: 1 1 auto` (cresce e encolhe, base auto)
- `flex-initial` = `flex: 0 1 auto` (não cresce, encolhe, base auto)
- `flex-none` = `flex: none` (não cresce nem encolhe)

```html
<div class="flex">
  <div class="flex-1">Ocupa espaço disponível</div>
  <div class="flex-none w-32">Largura fixa</div>
</div>
```

---

### Align Self

Controla o alinhamento de um item individual no eixo cruzado.

**Classes:**
- `self-auto` = `align-self: auto`
- `self-start` = `align-self: flex-start`
- `self-end` = `align-self: flex-end`
- `self-center` = `align-self: center`
- `self-stretch` = `align-self: stretch`
- `self-baseline` = `align-self: baseline`

```html
<div class="flex items-start h-32">
  <div class="self-center">Centralizado individualmente</div>
  <div>Alinhado no início</div>
</div>
```

---

### Gap (Espaçamento entre Itens)

Controla o espaço entre itens flex.

**Sintaxe:** `gap-{valor}`

```html
<div class="flex gap-4">
  <div>Item 1</div>
  <div>Item 2</div>
  <div>Item 3</div>
</div>
```

**Equivalente CSS:**
```css
div {
  display: flex;
  gap: 1rem;
}
```

**Variações:**
- `gap-x-4` = `column-gap: 1rem` (espaço horizontal)
- `gap-y-4` = `row-gap: 1rem` (espaço vertical)

**Vantagem:**
- Não precisa usar `margin` nos itens
- Mais limpo e sem colapso de margens

---

## 🎨 Exemplos Práticos de Layouts

### Layout de Header

```html
<header class="flex items-center justify-between p-4 bg-white shadow-md">
  <div class="flex items-center gap-4">
    <img src="logo.png" alt="Logo" class="h-8">
    <nav class="flex gap-4">
      <a href="#" class="text-blue-600">Home</a>
      <a href="#" class="text-gray-600">Sobre</a>
      <a href="#" class="text-gray-600">Contato</a>
    </nav>
  </div>
  <button class="px-4 py-2 bg-blue-500 text-white rounded">
    Login
  </button>
</header>
```

### Card com Layout Flex

```html
<div class="flex flex-col gap-4 p-6 bg-white rounded-lg shadow-lg max-w-md">
  <img src="image.jpg" alt="Imagem" class="w-full rounded">
  <div class="flex flex-col gap-2">
    <h3 class="text-xl font-bold">Título do Card</h3>
    <p class="text-gray-600">Descrição do card com algum texto explicativo.</p>
  </div>
  <div class="flex justify-between items-center mt-auto">
    <span class="text-2xl font-bold text-blue-600">R$ 99,90</span>
    <button class="px-4 py-2 bg-blue-500 text-white rounded">
      Comprar
    </button>
  </div>
</div>
```

### Sidebar com Layout

```html
<div class="flex h-screen">
  <!-- Sidebar -->
  <aside class="flex flex-col w-64 bg-gray-800 text-white p-4">
    <h2 class="text-xl font-bold mb-4">Menu</h2>
    <nav class="flex flex-col gap-2">
      <a href="#" class="p-2 rounded hover:bg-gray-700">Dashboard</a>
      <a href="#" class="p-2 rounded hover:bg-gray-700">Perfil</a>
      <a href="#" class="p-2 rounded hover:bg-gray-700">Configurações</a>
    </nav>
  </aside>
  
  <!-- Conteúdo Principal -->
  <main class="flex-1 p-8 bg-gray-100">
    <h1 class="text-3xl font-bold mb-4">Conteúdo Principal</h1>
    <p>Conteúdo da página...</p>
  </main>
</div>
```

### Formulário Vertical

```html
<form class="flex flex-col gap-4 max-w-md">
  <div class="flex flex-col gap-2">
    <label class="text-sm font-medium">Nome</label>
    <input type="text" class="px-4 py-2 border rounded">
  </div>
  <div class="flex flex-col gap-2">
    <label class="text-sm font-medium">Email</label>
    <input type="email" class="px-4 py-2 border rounded">
  </div>
  <div class="flex gap-4">
    <button type="submit" class="flex-1 px-4 py-2 bg-blue-500 text-white rounded">
      Enviar
    </button>
    <button type="reset" class="flex-1 px-4 py-2 bg-gray-300 rounded">
      Limpar
    </button>
  </div>
</form>
```

---

## 🔄 Resumo das Classes

### Display
- `block`, `inline`, `inline-block`
- `flex`, `grid`
- `hidden`

### Position
- `static`, `relative`, `absolute`, `fixed`, `sticky`
- `top-{valor}`, `right-{valor}`, `bottom-{valor}`, `left-{valor}`
- `inset-{valor}`, `inset-x-{valor}`, `inset-y-{valor}`
- `z-{valor}`

### Flexbox Container
- `flex`, `flex-row`, `flex-col`, `flex-row-reverse`, `flex-col-reverse`
- `flex-wrap`, `flex-nowrap`, `flex-wrap-reverse`
- `justify-start`, `justify-end`, `justify-center`
- `justify-between`, `justify-around`, `justify-evenly`
- `items-start`, `items-end`, `items-center`, `items-stretch`, `items-baseline`
- `gap-{valor}`, `gap-x-{valor}`, `gap-y-{valor}`

### Flexbox Items
- `flex-1`, `flex-auto`, `flex-initial`, `flex-none`
- `grow`, `grow-0`
- `shrink`, `shrink-0`
- `basis-{valor}`
- `self-auto`, `self-start`, `self-end`, `self-center`, `self-stretch`, `self-baseline`

---

## 🎯 Próximos Passos

Agora que você domina Display, Position e Flexbox com Tailwind, na próxima aula exploraremos **CSS Grid com Tailwind**, que oferece controle bidimensional ainda mais poderoso para layouts complexos.

**Conceitos importantes para revisar:**
- Como cada classe Tailwind mapeia para CSS puro
- Quando usar Flexbox vs Grid (será abordado na próxima aula)
- Combinações comuns de classes para layouts comuns

