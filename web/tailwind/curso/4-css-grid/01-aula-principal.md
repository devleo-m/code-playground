# Aula 4: CSS Grid com Tailwind - Conteúdo Principal

## 📖 Introdução

Na aula anterior, você dominou Flexbox com Tailwind. Agora vamos explorar **CSS Grid**, o sistema de layout bidimensional mais poderoso do CSS. Grid permite criar layouts complexos em duas dimensões (linhas e colunas) simultaneamente, algo que Flexbox não faz nativamente.

Você já conhece CSS Grid em CSS puro:
- `display: grid`
- `grid-template-columns` e `grid-template-rows`
- `grid-column` e `grid-row`
- `gap` para espaçamento
- `grid-template-areas` para layouts nomeados

Agora vamos ver como o Tailwind mapeia essas propriedades em classes utilitárias que tornam o Grid ainda mais produtivo.

**Nesta aula você aprenderá:**
- Grid utilities básicas (`grid`, `grid-cols-*`, `grid-rows-*`)
- Gap no Grid (`gap`, `gap-x`, `gap-y`)
- Spanning (`col-span-*`, `row-span-*`)
- Grid template areas
- Auto-fit e auto-fill
- Alinhamento no Grid (`place-items`, `place-content`)
- Grid responsivo
- Comparação: Grid vs Flexbox no Tailwind

---

## 🎨 Grid Utilities Básicas

CSS Grid é um sistema de layout bidimensional que permite criar layouts complexos dividindo o espaço em linhas e colunas.

### Display Grid

**Classe Tailwind:** `grid`

```html
<div class="grid">
  <div>Item 1</div>
  <div>Item 2</div>
  <div>Item 3</div>
</div>
```

**Equivalente CSS:**
```css
div {
  display: grid;
}
```

**Comportamento:**
- Cria um container de grid
- Filhos diretos se tornam grid items
- Por padrão, cria uma coluna única

### Grid Columns

O Tailwind oferece utilities para definir o número de colunas do grid.

#### Grid Columns Fixas

**Classes Tailwind:** `grid-cols-1` até `grid-cols-12`

```html
<div class="grid grid-cols-3 gap-4">
  <div class="bg-blue-500 p-4">Item 1</div>
  <div class="bg-blue-500 p-4">Item 2</div>
  <div class="bg-blue-500 p-4">Item 3</div>
  <div class="bg-blue-500 p-4">Item 4</div>
  <div class="bg-blue-500 p-4">Item 5</div>
  <div class="bg-blue-500 p-4">Item 6</div>
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

**Valores disponíveis:**
- `grid-cols-1` = 1 coluna
- `grid-cols-2` = 2 colunas
- `grid-cols-3` = 3 colunas
- `grid-cols-4` = 4 colunas
- `grid-cols-5` = 5 colunas
- `grid-cols-6` = 6 colunas
- `grid-cols-7` = 7 colunas
- `grid-cols-8` = 8 colunas
- `grid-cols-9` = 9 colunas
- `grid-cols-10` = 10 colunas
- `grid-cols-11` = 11 colunas
- `grid-cols-12` = 12 colunas

**Exemplo prático:**
```html
<div class="grid grid-cols-4 gap-4">
  <div class="bg-red-500 p-4 text-white">1</div>
  <div class="bg-red-500 p-4 text-white">2</div>
  <div class="bg-red-500 p-4 text-white">3</div>
  <div class="bg-red-500 p-4 text-white">4</div>
</div>
```

### Grid Rows

Similar às colunas, você pode definir o número de linhas.

**Classes Tailwind:** `grid-rows-1` até `grid-rows-6`

```html
<div class="grid grid-cols-2 grid-rows-3 gap-4 h-96">
  <div class="bg-green-500 p-4">Item 1</div>
  <div class="bg-green-500 p-4">Item 2</div>
  <div class="bg-green-500 p-4">Item 3</div>
  <div class="bg-green-500 p-4">Item 4</div>
  <div class="bg-green-500 p-4">Item 5</div>
  <div class="bg-green-500 p-4">Item 6</div>
</div>
```

**Equivalente CSS:**
```css
div {
  display: grid;
  grid-template-columns: repeat(2, minmax(0, 1fr));
  grid-template-rows: repeat(3, minmax(0, 1fr));
  gap: 1rem;
  height: 24rem;
}
```

**Valores disponíveis:**
- `grid-rows-1` até `grid-rows-6`

**Nota:** Normalmente você não precisa definir `grid-rows` explicitamente, pois o Grid cria linhas automaticamente conforme necessário.

---

## 📏 Gap no Grid

O `gap` define o espaçamento entre as células do grid.

### Gap Uniforme

**Classe Tailwind:** `gap-{tamanho}`

```html
<div class="grid grid-cols-3 gap-4">
  <div class="bg-blue-500 p-4">Item 1</div>
  <div class="bg-blue-500 p-4">Item 2</div>
  <div class="bg-blue-500 p-4">Item 3</div>
</div>
```

**Equivalente CSS:**
```css
div {
  display: grid;
  grid-template-columns: repeat(3, minmax(0, 1fr));
  gap: 1rem; /* gap-row e gap-column */
}
```

**Valores disponíveis:**
- `gap-0` = 0
- `gap-1` = 0.25rem (4px)
- `gap-2` = 0.5rem (8px)
- `gap-3` = 0.75rem (12px)
- `gap-4` = 1rem (16px)
- `gap-5` = 1.25rem (20px)
- `gap-6` = 1.5rem (24px)
- `gap-8` = 2rem (32px)
- `gap-10` = 2.5rem (40px)
- `gap-12` = 3rem (48px)
- E assim por diante...

### Gap Horizontal (Column Gap)

**Classe Tailwind:** `gap-x-{tamanho}`

```html
<div class="grid grid-cols-3 gap-x-8">
  <div class="bg-blue-500 p-4">Item 1</div>
  <div class="bg-blue-500 p-4">Item 2</div>
  <div class="bg-blue-500 p-4">Item 3</div>
</div>
```

**Equivalente CSS:**
```css
div {
  display: grid;
  grid-template-columns: repeat(3, minmax(0, 1fr));
  column-gap: 2rem;
}
```

### Gap Vertical (Row Gap)

**Classe Tailwind:** `gap-y-{tamanho}`

```html
<div class="grid grid-cols-3 gap-y-6">
  <div class="bg-blue-500 p-4">Item 1</div>
  <div class="bg-blue-500 p-4">Item 2</div>
  <div class="bg-blue-500 p-4">Item 3</div>
</div>
```

**Equivalente CSS:**
```css
div {
  display: grid;
  grid-template-columns: repeat(3, minmax(0, 1fr));
  row-gap: 1.5rem;
}
```

**Exemplo combinado:**
```html
<div class="grid grid-cols-3 gap-x-8 gap-y-4">
  <!-- Espaçamento horizontal maior que vertical -->
</div>
```

---

## 🔀 Spanning (Estender Células)

Spanning permite que um item do grid ocupe múltiplas colunas ou linhas.

### Column Span

**Classes Tailwind:** `col-span-1` até `col-span-12`

```html
<div class="grid grid-cols-4 gap-4">
  <div class="col-span-2 bg-blue-500 p-4">Ocupa 2 colunas</div>
  <div class="bg-green-500 p-4">1 coluna</div>
  <div class="bg-green-500 p-4">1 coluna</div>
  <div class="col-span-3 bg-red-500 p-4">Ocupa 3 colunas</div>
  <div class="bg-green-500 p-4">1 coluna</div>
</div>
```

**Equivalente CSS:**
```css
.item-1 {
  grid-column: span 2 / span 2;
}

.item-4 {
  grid-column: span 3 / span 3;
}
```

**Valores disponíveis:**
- `col-span-1` até `col-span-12`
- `col-span-full` = ocupa todas as colunas disponíveis

**Exemplo prático - Layout de blog:**
```html
<div class="grid grid-cols-12 gap-4">
  <!-- Sidebar -->
  <aside class="col-span-3 bg-gray-200 p-4">
    Sidebar
  </aside>
  
  <!-- Conteúdo principal -->
  <main class="col-span-9 bg-white p-4">
    Conteúdo Principal
  </main>
</div>
```

### Row Span

**Classes Tailwind:** `row-span-1` até `row-span-6`

```html
<div class="grid grid-cols-3 grid-rows-3 gap-4 h-96">
  <div class="row-span-2 bg-blue-500 p-4">Ocupa 2 linhas</div>
  <div class="bg-green-500 p-4">1 linha</div>
  <div class="bg-green-500 p-4">1 linha</div>
  <div class="bg-green-500 p-4">1 linha</div>
  <div class="row-span-3 bg-red-500 p-4">Ocupa 3 linhas</div>
</div>
```

**Equivalente CSS:**
```css
.item-1 {
  grid-row: span 2 / span 2;
}

.item-5 {
  grid-row: span 3 / span 3;
}
```

**Valores disponíveis:**
- `row-span-1` até `row-span-6`
- `row-span-full` = ocupa todas as linhas disponíveis

### Spanning Combinado

Você pode combinar `col-span` e `row-span`:

```html
<div class="grid grid-cols-4 grid-rows-4 gap-4 h-96">
  <div class="col-span-2 row-span-2 bg-blue-500 p-4">
    Ocupa 2x2
  </div>
  <div class="bg-green-500 p-4">1x1</div>
  <div class="bg-green-500 p-4">1x1</div>
  <div class="col-span-3 bg-red-500 p-4">Ocupa 3 colunas</div>
</div>
```

**Equivalente CSS:**
```css
.item {
  grid-column: span 2 / span 2;
  grid-row: span 2 / span 2;
}
```

---

## 📐 Grid Template Areas

Grid template areas permite criar layouts nomeados, facilitando a organização visual do layout.

### Definindo Áreas

No Tailwind, você precisa usar CSS customizado para `grid-template-areas`, mas pode combinar com outras utilities:

```html
<div class="grid grid-cols-4 gap-4" style="grid-template-areas: 'header header header header' 'sidebar main main aside' 'footer footer footer footer';">
  <header class="bg-blue-500 p-4" style="grid-area: header;">Header</header>
  <aside class="bg-green-500 p-4" style="grid-area: sidebar;">Sidebar</aside>
  <main class="bg-yellow-500 p-4" style="grid-area: main;">Main</main>
  <aside class="bg-purple-500 p-4" style="grid-area: aside;">Aside</aside>
  <footer class="bg-red-500 p-4" style="grid-area: footer;">Footer</footer>
</div>
```

**Nota:** Para layouts complexos com `grid-template-areas`, considere usar `@apply` ou CSS customizado, pois o Tailwind não tem utilities diretas para isso.

**Alternativa com classes Tailwind:**
```html
<div class="grid grid-cols-4 gap-4">
  <header class="col-span-4 bg-blue-500 p-4">Header</header>
  <aside class="col-span-1 bg-green-500 p-4">Sidebar</aside>
  <main class="col-span-2 bg-yellow-500 p-4">Main</main>
  <aside class="col-span-1 bg-purple-500 p-4">Aside</aside>
  <footer class="col-span-4 bg-red-500 p-4">Footer</footer>
</div>
```

---

## 🔄 Auto-fit e Auto-fill

Para criar grids que se adaptam automaticamente ao número de itens, você pode usar CSS customizado com `repeat(auto-fit, minmax())` ou `repeat(auto-fill, minmax())`.

### Auto-fit com Tailwind

```html
<div class="grid gap-4" style="grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));">
  <div class="bg-blue-500 p-4">Item 1</div>
  <div class="bg-blue-500 p-4">Item 2</div>
  <div class="bg-blue-500 p-4">Item 3</div>
  <div class="bg-blue-500 p-4">Item 4</div>
</div>
```

**Equivalente CSS:**
```css
div {
  display: grid;
  gap: 1rem;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
}
```

**Comportamento:**
- `auto-fit`: Cria o máximo de colunas que cabem, esticando os itens para preencher o espaço
- `auto-fill`: Cria o máximo de colunas que cabem, mantendo o tamanho mínimo

**Exemplo prático - Galeria responsiva:**
```html
<div class="grid gap-4 p-4" style="grid-template-columns: repeat(auto-fill, minmax(200px, 1fr));">
  <div class="bg-gray-300 h-48 rounded"></div>
  <div class="bg-gray-300 h-48 rounded"></div>
  <div class="bg-gray-300 h-48 rounded"></div>
  <div class="bg-gray-300 h-48 rounded"></div>
  <div class="bg-gray-300 h-48 rounded"></div>
  <div class="bg-gray-300 h-48 rounded"></div>
</div>
```

---

## 🎯 Alinhamento no Grid

O Tailwind oferece utilities para alinhar itens dentro das células do grid.

### Place Items (Alinhamento de Itens)

Controla o alinhamento de todos os itens dentro de suas células.

**Classes Tailwind:**
- `place-items-start`
- `place-items-end`
- `place-items-center`
- `place-items-stretch` (padrão)

```html
<div class="grid grid-cols-3 gap-4 place-items-center h-64">
  <div class="bg-blue-500 p-4">Centrado</div>
  <div class="bg-blue-500 p-4">Centrado</div>
  <div class="bg-blue-500 p-4">Centrado</div>
</div>
```

**Equivalente CSS:**
```css
div {
  display: grid;
  grid-template-columns: repeat(3, minmax(0, 1fr));
  gap: 1rem;
  place-items: center;
}
```

**Valores disponíveis:**
- `place-items-start` = `place-items: start`
- `place-items-end` = `place-items: end`
- `place-items-center` = `place-items: center`
- `place-items-stretch` = `place-items: stretch`

### Place Content (Alinhamento do Grid)

Controla o alinhamento do grid inteiro dentro do container.

**Classes Tailwind:**
- `place-content-center`
- `place-content-start`
- `place-content-end`
- `place-content-between`
- `place-content-around`
- `place-content-evenly`
- `place-content-stretch`

```html
<div class="grid grid-cols-3 gap-4 place-content-center h-96">
  <div class="bg-blue-500 p-4">Item 1</div>
  <div class="bg-blue-500 p-4">Item 2</div>
  <div class="bg-blue-500 p-4">Item 3</div>
</div>
```

**Equivalente CSS:**
```css
div {
  display: grid;
  grid-template-columns: repeat(3, minmax(0, 1fr));
  gap: 1rem;
  place-content: center;
}
```

### Alinhamento Individual

Você pode alinhar itens individuais usando `place-self-*`:

**Classes Tailwind:**
- `place-self-auto`
- `place-self-start`
- `place-self-end`
- `place-self-center`
- `place-self-stretch`

```html
<div class="grid grid-cols-3 gap-4 h-64">
  <div class="bg-blue-500 p-4 place-self-start">Start</div>
  <div class="bg-green-500 p-4 place-self-center">Center</div>
  <div class="bg-red-500 p-4 place-self-end">End</div>
</div>
```

---

## 📱 Grid Responsivo

O Tailwind permite criar grids responsivos usando breakpoints.

### Grid Responsivo Básico

```html
<div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-4">
  <div class="bg-blue-500 p-4">Item 1</div>
  <div class="bg-blue-500 p-4">Item 2</div>
  <div class="bg-blue-500 p-4">Item 3</div>
  <div class="bg-blue-500 p-4">Item 4</div>
</div>
```

**Comportamento:**
- Mobile: 1 coluna
- Tablet (md): 2 colunas
- Desktop (lg): 3 colunas
- Large (xl): 4 colunas

**Equivalente CSS:**
```css
.grid {
  display: grid;
  grid-template-columns: repeat(1, minmax(0, 1fr));
  gap: 1rem;
}

@media (min-width: 768px) {
  .grid {
    grid-template-columns: repeat(2, minmax(0, 1fr));
  }
}

@media (min-width: 1024px) {
  .grid {
    grid-template-columns: repeat(3, minmax(0, 1fr));
  }
}

@media (min-width: 1280px) {
  .grid {
    grid-template-columns: repeat(4, minmax(0, 1fr));
  }
}
```

### Spanning Responsivo

Você pode fazer spanning responsivo:

```html
<div class="grid grid-cols-4 gap-4">
  <div class="col-span-4 md:col-span-2 lg:col-span-1 bg-blue-500 p-4">
    Responsivo
  </div>
  <div class="col-span-4 md:col-span-2 lg:col-span-3 bg-green-500 p-4">
    Responsivo
  </div>
</div>
```

**Comportamento:**
- Mobile: ambos ocupam 4 colunas (100%)
- Tablet: cada um ocupa 2 colunas (50%)
- Desktop: primeiro ocupa 1 coluna, segundo ocupa 3 colunas

### Exemplo Prático - Layout de Dashboard

```html
<div class="grid grid-cols-1 lg:grid-cols-12 gap-4">
  <!-- Sidebar -->
  <aside class="lg:col-span-3 bg-gray-200 p-4">
    Sidebar
  </aside>
  
  <!-- Conteúdo -->
  <main class="lg:col-span-9 grid grid-cols-1 md:grid-cols-2 gap-4">
    <div class="bg-white p-4">Card 1</div>
    <div class="bg-white p-4">Card 2</div>
    <div class="bg-white p-4 md:col-span-2">Card 3 (largura total)</div>
  </main>
</div>
```

---

## ⚖️ Grid vs Flexbox no Tailwind

### Quando usar Grid

**Use Grid quando:**
- Você precisa de layout bidimensional (linhas E colunas)
- Você quer criar layouts complexos com áreas nomeadas
- Você precisa de controle preciso sobre posicionamento
- Você quer criar galerias ou grids de cards
- Você precisa de layouts que se adaptam ao espaço disponível

**Exemplos:**
- Layouts de página completos (header, sidebar, main, footer)
- Galerias de imagens
- Dashboards com múltiplas seções
- Formulários complexos com múltiplas colunas

### Quando usar Flexbox

**Use Flexbox quando:**
- Você precisa de layout unidimensional (linha OU coluna)
- Você quer alinhar itens em uma direção
- Você precisa de distribuição de espaço flexível
- Você quer criar componentes como navbars, cards, botões

**Exemplos:**
- Navbars e menus
- Cards com conteúdo vertical
- Botões com ícones
- Listas de itens em uma direção

### Combinando Grid e Flexbox

Você pode e deve combinar ambos:

```html
<div class="grid grid-cols-3 gap-4">
  <!-- Grid para layout principal -->
  <div class="flex flex-col gap-2">
    <!-- Flexbox para conteúdo interno -->
    <h3 class="font-bold">Título</h3>
    <p>Descrição</p>
    <button class="mt-auto">Botão</button>
  </div>
</div>
```

---

## 🎨 Exemplos Práticos

### Exemplo 1: Layout de Blog

```html
<div class="grid grid-cols-1 lg:grid-cols-12 gap-6">
  <!-- Sidebar -->
  <aside class="lg:col-span-3 bg-gray-100 p-6 rounded">
    <h2 class="font-bold text-xl mb-4">Sidebar</h2>
    <nav class="flex flex-col gap-2">
      <a href="#" class="text-blue-600 hover:underline">Link 1</a>
      <a href="#" class="text-blue-600 hover:underline">Link 2</a>
      <a href="#" class="text-blue-600 hover:underline">Link 3</a>
    </nav>
  </aside>
  
  <!-- Conteúdo Principal -->
  <main class="lg:col-span-9">
    <article class="bg-white p-6 rounded shadow">
      <h1 class="text-3xl font-bold mb-4">Título do Artigo</h1>
      <p class="text-gray-700 mb-4">Conteúdo do artigo...</p>
    </article>
  </main>
</div>
```

### Exemplo 2: Galeria de Imagens

```html
<div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-4">
  <div class="bg-gray-300 h-48 rounded-lg"></div>
  <div class="bg-gray-300 h-48 rounded-lg"></div>
  <div class="bg-gray-300 h-48 rounded-lg"></div>
  <div class="bg-gray-300 h-48 rounded-lg"></div>
  <div class="bg-gray-300 h-48 rounded-lg col-span-1 sm:col-span-2 lg:col-span-1"></div>
  <div class="bg-gray-300 h-48 rounded-lg"></div>
</div>
```

### Exemplo 3: Dashboard com Cards

```html
<div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
  <!-- Card de Estatística -->
  <div class="bg-white p-6 rounded-lg shadow">
    <h3 class="text-gray-500 text-sm mb-2">Total de Vendas</h3>
    <p class="text-3xl font-bold">R$ 12.345</p>
  </div>
  
  <div class="bg-white p-6 rounded-lg shadow">
    <h3 class="text-gray-500 text-sm mb-2">Usuários</h3>
    <p class="text-3xl font-bold">1.234</p>
  </div>
  
  <div class="bg-white p-6 rounded-lg shadow">
    <h3 class="text-gray-500 text-sm mb-2">Pedidos</h3>
    <p class="text-3xl font-bold">567</p>
  </div>
  
  <div class="bg-white p-6 rounded-lg shadow">
    <h3 class="text-gray-500 text-sm mb-2">Taxa de Conversão</h3>
    <p class="text-3xl font-bold">12.5%</p>
  </div>
  
  <!-- Gráfico (ocupa 2 colunas) -->
  <div class="md:col-span-2 lg:col-span-4 bg-white p-6 rounded-lg shadow">
    <h3 class="text-xl font-bold mb-4">Gráfico de Vendas</h3>
    <div class="bg-gray-100 h-64 rounded"></div>
  </div>
</div>
```

---

## 📝 Resumo das Classes Principais

### Grid Container
- `grid` - Cria container grid
- `grid-cols-{n}` - Define número de colunas (1-12)
- `grid-rows-{n}` - Define número de linhas (1-6)

### Gap
- `gap-{tamanho}` - Espaçamento uniforme
- `gap-x-{tamanho}` - Espaçamento horizontal
- `gap-y-{tamanho}` - Espaçamento vertical

### Spanning
- `col-span-{n}` - Ocupa n colunas (1-12, full)
- `row-span-{n}` - Ocupa n linhas (1-6, full)

### Alinhamento
- `place-items-{valor}` - Alinha todos os itens
- `place-content-{valor}` - Alinha o grid
- `place-self-{valor}` - Alinha item individual

### Responsivo
- `{breakpoint}:grid-cols-{n}` - Grid responsivo
- `{breakpoint}:col-span-{n}` - Spanning responsivo

---

## 🎯 Conclusão

CSS Grid com Tailwind oferece uma forma poderosa e produtiva de criar layouts bidimensionais complexos. As classes utilitárias do Tailwind tornam o Grid ainda mais acessível, permitindo criar layouts responsivos e flexíveis com poucas classes.

**Principais pontos:**
- Grid é ideal para layouts bidimensionais
- Use `grid-cols-*` para definir colunas
- Use `col-span-*` e `row-span-*` para spanning
- Combine Grid com Flexbox quando necessário
- Use breakpoints para grids responsivos
- Grid e Flexbox se complementam, não competem

Na próxima aula, exploraremos **Responsividade com Tailwind** em profundidade, incluindo breakpoints customizados e estratégias mobile-first.

