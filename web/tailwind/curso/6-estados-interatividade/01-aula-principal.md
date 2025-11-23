# Aula 6: Estados e Interatividade com Tailwind CSS - Conteúdo Principal

## 📖 Revisão: Estados e Pseudo-classes em CSS

Antes de mergulharmos nos estados e interatividade com Tailwind, vamos relembrar os conceitos fundamentais que você já conhece de CSS puro:

**Pseudo-classes** são palavras-chave que permitem estilizar elementos baseados em seu estado ou posição no documento.

### Estados Básicos em CSS Puro

Você já aprendeu que em CSS puro usamos pseudo-classes para criar interatividade:

```css
/* CSS Puro - Estados básicos */
.button {
  background-color: blue;
  color: white;
}

.button:hover {
  background-color: darkblue; /* Quando o mouse passa por cima */
}

.button:focus {
  outline: 2px solid orange; /* Quando o elemento recebe foco */
}

.button:active {
  background-color: navy; /* Quando o elemento está sendo clicado */
}

.button:disabled {
  opacity: 0.5; /* Quando o elemento está desabilitado */
  cursor: not-allowed;
}
```

### Estados no Tailwind

O Tailwind CSS fornece **prefixos de estado** que você pode usar diretamente nas classes utilitárias, tornando a criação de interfaces interativas muito mais rápida e intuitiva.

---

## 🎯 Estados de Hover

### O que é Hover?

**Hover** é o estado que ocorre quando o usuário passa o cursor do mouse sobre um elemento, mas ainda não clicou nele.

### Hover em CSS Puro

```css
.button {
  background-color: rgb(59, 130, 246);
  transition: background-color 150ms;
}

.button:hover {
  background-color: rgb(37, 99, 235);
}
```

### Hover no Tailwind

No Tailwind, você usa o prefixo `hover:` para aplicar estilos quando o elemento está em estado hover:

```html
<!-- Botão com hover -->
<button class="bg-blue-500 hover:bg-blue-600 text-white px-4 py-2 rounded">
  Clique em mim
</button>
```

**Equivalente em CSS puro:**
```css
button {
  background-color: rgb(59, 130, 246);
  color: white;
  padding: 0.5rem 1rem;
  border-radius: 0.25rem;
}

button:hover {
  background-color: rgb(37, 99, 235);
}
```

### Exemplos Práticos de Hover

```html
<!-- Hover em diferentes propriedades -->
<div class="bg-gray-100 hover:bg-gray-200 p-4 rounded">
  Card com hover de background
</div>

<div class="text-gray-600 hover:text-blue-600">
  Link com hover de cor de texto
</div>

<div class="border-2 border-gray-300 hover:border-blue-500 p-4">
  Card com hover de borda
</div>

<div class="shadow-md hover:shadow-lg transition-shadow">
  Card com hover de sombra
</div>

<div class="scale-100 hover:scale-105 transition-transform">
  Elemento que cresce no hover
</div>
```

### Múltiplas Propriedades no Hover

Você pode aplicar múltiplas propriedades no hover:

```html
<button class="
  bg-blue-500 
  text-white 
  px-6 py-3 
  rounded-lg 
  shadow-md
  hover:bg-blue-600 
  hover:shadow-lg 
  hover:scale-105 
  hover:-translate-y-1
  transition-all
  duration-200
">
  Botão Interativo
</button>
```

**Equivalente em CSS puro:**
```css
button {
  background-color: rgb(59, 130, 246);
  color: white;
  padding: 0.75rem 1.5rem;
  border-radius: 0.5rem;
  box-shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.1);
  transition: all 200ms;
}

button:hover {
  background-color: rgb(37, 99, 235);
  box-shadow: 0 10px 15px -3px rgba(0, 0, 0, 0.1);
  transform: scale(1.05) translateY(-0.25rem);
}
```

---

## 🎯 Estados de Focus

### O que é Focus?

**Focus** é o estado que ocorre quando um elemento recebe foco, geralmente através de navegação por teclado (Tab) ou clicando no elemento (para elementos interativos como inputs e botões).

### Focus em CSS Puro

```css
.input {
  border: 2px solid gray;
  outline: none;
}

.input:focus {
  border-color: blue;
  outline: 2px solid blue;
  outline-offset: 2px;
}
```

### Focus no Tailwind

No Tailwind, você usa o prefixo `focus:` para aplicar estilos quando o elemento está em foco:

```html
<!-- Input com focus -->
<input 
  type="text" 
  class="border-2 border-gray-300 focus:border-blue-500 focus:outline-none focus:ring-2 focus:ring-blue-500 px-4 py-2 rounded"
  placeholder="Digite algo..."
/>
```

**Equivalente em CSS puro:**
```css
input {
  border: 2px solid rgb(209, 213, 219);
  padding: 0.5rem 1rem;
  border-radius: 0.25rem;
}

input:focus {
  border-color: rgb(59, 130, 246);
  outline: none;
  box-shadow: 0 0 0 2px rgba(59, 130, 246, 0.5);
}
```

### Variantes de Focus

O Tailwind oferece diferentes variantes de focus:

#### `focus:` - Focus geral
```html
<input class="focus:border-blue-500" />
```

#### `focus-visible:` - Focus apenas quando navegando por teclado
```html
<button class="focus-visible:ring-2 focus-visible:ring-blue-500">
  Botão com focus-visible
</button>
```

**Por que usar `focus-visible`?**
- Melhora a acessibilidade: mostra o foco apenas quando necessário (navegação por teclado)
- Evita o anel de foco desnecessário quando clicamos com o mouse

#### `focus-within:` - Focus em qualquer elemento filho
```html
<div class="focus-within:ring-2 focus-within:ring-blue-500">
  <input type="text" />
  <button>Enviar</button>
</div>
```

**Equivalente em CSS puro:**
```css
div:focus-within {
  box-shadow: 0 0 0 2px rgba(59, 130, 246, 0.5);
}
```

### Exemplos Práticos de Focus

```html
<!-- Input com focus ring -->
<input 
  type="email"
  class="
    w-full 
    px-4 py-2 
    border border-gray-300 
    rounded-lg
    focus:outline-none
    focus:ring-2 
    focus:ring-blue-500 
    focus:border-transparent
  "
  placeholder="seu@email.com"
/>

<!-- Textarea com focus -->
<textarea 
  class="
    w-full 
    px-4 py-2 
    border border-gray-300 
    rounded-lg
    focus:outline-none
    focus:ring-2 
    focus:ring-green-500 
    focus:border-green-500
  "
  rows="4"
></textarea>

<!-- Select com focus -->
<select 
  class="
    w-full 
    px-4 py-2 
    border border-gray-300 
    rounded-lg
    focus:outline-none
    focus:ring-2 
    focus:ring-purple-500
  "
>
  <option>Opção 1</option>
  <option>Opção 2</option>
</select>
```

---

## 🎯 Estados de Active

### O que é Active?

**Active** é o estado que ocorre quando o usuário está **clicando** ou **pressionando** um elemento (o momento entre o clique e o soltar).

### Active em CSS Puro

```css
.button {
  background-color: blue;
}

.button:active {
  background-color: darkblue;
  transform: scale(0.95);
}
```

### Active no Tailwind

No Tailwind, você usa o prefixo `active:` para aplicar estilos quando o elemento está sendo ativado:

```html
<button class="bg-blue-500 active:bg-blue-700 active:scale-95 text-white px-4 py-2 rounded">
  Clique e segure
</button>
```

**Equivalente em CSS puro:**
```css
button {
  background-color: rgb(59, 130, 246);
  color: white;
  padding: 0.5rem 1rem;
  border-radius: 0.25rem;
  transition: all 150ms;
}

button:active {
  background-color: rgb(29, 78, 216);
  transform: scale(0.95);
}
```

### Exemplos Práticos de Active

```html
<!-- Botão com feedback visual no clique -->
<button class="
  bg-green-500 
  hover:bg-green-600 
  active:bg-green-700 
  active:scale-95
  text-white 
  px-6 py-3 
  rounded-lg
  transition-all
  duration-150
">
  Enviar
</button>

<!-- Card clicável -->
<div class="
  bg-white 
  p-6 
  rounded-lg 
  shadow-md
  cursor-pointer
  hover:shadow-lg
  active:shadow-sm
  active:scale-98
  transition-all
">
  Card Clicável
</div>
```

---

## 🎯 Estados Disabled e Required

### Disabled

O estado **disabled** ocorre quando um elemento está desabilitado e não pode ser interagido.

```html
<!-- Botão desabilitado -->
<button 
  disabled
  class="
    bg-gray-400 
    text-white 
    px-4 py-2 
    rounded
    cursor-not-allowed
    disabled:opacity-50
    disabled:cursor-not-allowed
  "
>
  Desabilitado
</button>
```

**Equivalente em CSS puro:**
```css
button:disabled {
  opacity: 0.5;
  cursor: not-allowed;
}
```

### Required

O estado **required** é aplicado a elementos de formulário que são obrigatórios.

```html
<input 
  type="text"
  required
  class="
    border-2 
    border-gray-300
    required:border-red-500
    px-4 py-2
    rounded
  "
/>
```

---

## 🎯 Pseudo-classes de Posição

### First, Last, Odd, Even

O Tailwind fornece utilitários para estilizar elementos baseados em sua posição:

#### `first:` - Primeiro elemento
```html
<ul>
  <li class="first:pt-0">Primeiro item</li>
  <li>Segundo item</li>
  <li>Terceiro item</li>
</ul>
```

**Equivalente em CSS puro:**
```css
li:first-child {
  padding-top: 0;
}
```

#### `last:` - Último elemento
```html
<ul>
  <li>Primeiro item</li>
  <li>Segundo item</li>
  <li class="last:pb-0">Último item</li>
</ul>
```

#### `odd:` e `even:` - Elementos ímpares e pares
```html
<div class="space-y-2">
  <div class="odd:bg-gray-100 even:bg-white p-4">
    Item 1 (ímpar)
  </div>
  <div class="odd:bg-gray-100 even:bg-white p-4">
    Item 2 (par)
  </div>
  <div class="odd:bg-gray-100 even:bg-white p-4">
    Item 3 (ímpar)
  </div>
</div>
```

**Equivalente em CSS puro:**
```css
div:nth-child(odd) {
  background-color: rgb(243, 244, 246);
}

div:nth-child(even) {
  background-color: white;
}
```

### Exemplo Prático: Lista com Estilos Alternados

```html
<ul class="divide-y divide-gray-200">
  <li class="px-4 py-3 first:rounded-t-lg last:rounded-b-lg hover:bg-gray-50">
    Item 1
  </li>
  <li class="px-4 py-3 first:rounded-t-lg last:rounded-b-lg hover:bg-gray-50">
    Item 2
  </li>
  <li class="px-4 py-3 first:rounded-t-lg last:rounded-b-lg hover:bg-gray-50">
    Item 3
  </li>
</ul>
```

---

## 🎯 Group e Peer Utilities

### Group - Estados em Elementos Pais

O `group` permite aplicar estilos em elementos filhos quando o elemento pai está em um estado específico.

```html
<!-- Card com hover que afeta elementos filhos -->
<div class="group bg-white p-6 rounded-lg shadow-md hover:shadow-lg cursor-pointer">
  <h3 class="text-gray-800 group-hover:text-blue-600">Título</h3>
  <p class="text-gray-600 group-hover:text-gray-800">Descrição</p>
  <button class="mt-4 opacity-0 group-hover:opacity-100 transition-opacity">
    Ver mais
  </button>
</div>
```

**Equivalente em CSS puro:**
```css
.card {
  background-color: white;
  padding: 1.5rem;
  border-radius: 0.5rem;
  box-shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.1);
  cursor: pointer;
}

.card:hover {
  box-shadow: 0 10px 15px -3px rgba(0, 0, 0, 0.1);
}

.card:hover h3 {
  color: rgb(37, 99, 235);
}

.card:hover p {
  color: rgb(31, 41, 55);
}

.card button {
  opacity: 0;
  transition: opacity 150ms;
}

.card:hover button {
  opacity: 1;
}
```

### Peer - Estados em Elementos Irmãos

O `peer` permite aplicar estilos em um elemento baseado no estado de seu irmão adjacente.

```html
<!-- Input que afeta o label quando em focus -->
<div class="relative">
  <input 
    type="text"
    id="email"
    class="
      peer
      w-full 
      px-4 py-2 
      border-2 border-gray-300 
      rounded-lg
      focus:outline-none
      focus:border-blue-500
    "
    placeholder=" "
  />
  <label 
    for="email"
    class="
      absolute 
      left-4 
      top-2 
      text-gray-500
      transition-all
      peer-focus:text-blue-500
      peer-focus:top-0
      peer-focus:text-sm
      peer-placeholder-shown:top-2
      peer-placeholder-shown:text-base
    "
  >
    Email
  </label>
</div>
```

**Equivalente em CSS puro:**
```css
input:focus + label {
  color: rgb(59, 130, 246);
  top: 0;
  font-size: 0.875rem;
}

input:placeholder-shown + label {
  top: 0.5rem;
  font-size: 1rem;
}
```

### Exemplo Prático: Toggle Switch

```html
<!-- Toggle switch com peer -->
<label class="relative inline-flex items-center cursor-pointer">
  <input type="checkbox" class="sr-only peer" />
  <div class="
    w-11 h-6 
    bg-gray-200 
    peer-focus:outline-none 
    peer-focus:ring-4 
    peer-focus:ring-blue-300 
    rounded-full 
    peer 
    peer-checked:after:translate-x-full 
    peer-checked:after:border-white 
    after:content-[''] 
    after:absolute 
    after:top-[2px] 
    after:left-[2px] 
    after:bg-white 
    after:border-gray-300 
    after:border 
    after:rounded-full 
    after:h-5 
    after:w-5 
    after:transition-all 
    peer-checked:bg-blue-600
  "></div>
  <span class="ml-3 text-sm font-medium text-gray-900">Toggle me</span>
</label>
```

---

## 🎯 Transições e Animações

### Transitions em CSS Puro

```css
.button {
  background-color: blue;
  transition: background-color 300ms ease-in-out;
}

.button:hover {
  background-color: darkblue;
}
```

### Transitions no Tailwind

O Tailwind fornece utilitários para criar transições suaves:

#### `transition` - Propriedade geral
```html
<div class="bg-blue-500 hover:bg-blue-600 transition">
  Transição suave
</div>
```

#### `transition-all` - Todas as propriedades
```html
<div class="bg-blue-500 hover:bg-blue-600 hover:scale-105 transition-all">
  Transição em todas as propriedades
</div>
```

#### `transition-colors` - Apenas cores
```html
<div class="bg-blue-500 hover:bg-blue-600 transition-colors">
  Transição apenas de cores
</div>
```

#### `transition-transform` - Apenas transformações
```html
<div class="scale-100 hover:scale-110 transition-transform">
  Transição apenas de transform
</div>
```

### Duração das Transições

```html
<!-- Durações disponíveis -->
<div class="transition duration-75">75ms</div>
<div class="transition duration-100">100ms</div>
<div class="transition duration-150">150ms (padrão)</div>
<div class="transition duration-200">200ms</div>
<div class="transition duration-300">300ms</div>
<div class="transition duration-500">500ms</div>
<div class="transition duration-700">700ms</div>
<div class="transition duration-1000">1000ms</div>
```

### Timing Functions (Curvas de Animação)

```html
<!-- Timing functions disponíveis -->
<div class="transition ease-linear">Linear</div>
<div class="transition ease-in">Ease In</div>
<div class="transition ease-out">Ease Out</div>
<div class="transition ease-in-out">Ease In Out (padrão)</div>
```

**Equivalente em CSS puro:**
```css
.ease-linear {
  transition-timing-function: linear;
}

.ease-in {
  transition-timing-function: cubic-bezier(0.4, 0, 1, 1);
}

.ease-out {
  transition-timing-function: cubic-bezier(0, 0, 0.2, 1);
}

.ease-in-out {
  transition-timing-function: cubic-bezier(0.4, 0, 0.2, 1);
}
```

### Delay nas Transições

```html
<div class="transition delay-75">Delay de 75ms</div>
<div class="transition delay-100">Delay de 100ms</div>
<div class="transition delay-150">Delay de 150ms</div>
<div class="transition delay-200">Delay de 200ms</div>
<div class="transition delay-300">Delay de 300ms</div>
<div class="transition delay-500">Delay de 500ms</div>
<div class="transition delay-700">Delay de 700ms</div>
<div class="transition delay-1000">Delay de 1000ms</div>
```

### Exemplo Completo de Transição

```html
<button class="
  bg-blue-500 
  hover:bg-blue-600 
  text-white 
  px-6 py-3 
  rounded-lg
  shadow-md
  hover:shadow-lg
  hover:scale-105
  hover:-translate-y-1
  active:scale-95
  transition-all
  duration-200
  ease-in-out
">
  Botão com Transição Completa
</button>
```

**Equivalente em CSS puro:**
```css
button {
  background-color: rgb(59, 130, 246);
  color: white;
  padding: 0.75rem 1.5rem;
  border-radius: 0.5rem;
  box-shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.1);
  transition: all 200ms cubic-bezier(0.4, 0, 0.2, 1);
}

button:hover {
  background-color: rgb(37, 99, 235);
  box-shadow: 0 10px 15px -3px rgba(0, 0, 0, 0.1);
  transform: scale(1.05) translateY(-0.25rem);
}

button:active {
  transform: scale(0.95);
}
```

---

## 🎯 Transform Utilities

### Transform em CSS Puro

```css
.element {
  transform: translateX(10px) rotate(45deg) scale(1.2);
}
```

### Transform no Tailwind

O Tailwind fornece utilitários para transformações comuns:

#### Scale (Escala)
```html
<div class="scale-100 hover:scale-110">Cresce no hover</div>
<div class="scale-50">50% do tamanho</div>
<div class="scale-150">150% do tamanho</div>
```

#### Rotate (Rotação)
```html
<div class="rotate-0 hover:rotate-45">Rotaciona no hover</div>
<div class="rotate-90">90 graus</div>
<div class="rotate-180">180 graus</div>
<div class="-rotate-45">-45 graus (sentido anti-horário)</div>
```

#### Translate (Translação)
```html
<div class="translate-x-4">Move 1rem para direita</div>
<div class="-translate-x-4">Move 1rem para esquerda</div>
<div class="translate-y-4">Move 1rem para baixo</div>
<div class="-translate-y-4">Move 1rem para cima</div>
```

#### Skew (Inclinação)
```html
<div class="skew-x-12">Inclina no eixo X</div>
<div class="skew-y-12">Inclina no eixo Y</div>
<div class="-skew-x-12">Inclina no sentido oposto</div>
```

### Combinando Transforms

```html
<div class="
  hover:scale-110 
  hover:rotate-3 
  hover:-translate-y-2
  transition-transform
  duration-300
">
  Múltiplas transformações
</div>
```

**Equivalente em CSS puro:**
```css
div:hover {
  transform: scale(1.1) rotate(3deg) translateY(-0.5rem);
  transition: transform 300ms;
}
```

---

## 🎯 Animações Pré-definidas

O Tailwind vem com algumas animações pré-definidas:

### `animate-spin` - Rotação contínua
```html
<div class="animate-spin">⚙️</div>
```

**Equivalente em CSS puro:**
```css
@keyframes spin {
  to {
    transform: rotate(360deg);
  }
}

.animate-spin {
  animation: spin 1s linear infinite;
}
```

### `animate-ping` - Efeito de ping
```html
<div class="relative">
  <div class="animate-ping absolute">🔔</div>
  <div class="relative">Notificação</div>
</div>
```

### `animate-pulse` - Pulsação
```html
<div class="animate-pulse bg-gray-200 h-4 w-32"></div>
```

### `animate-bounce` - Bounce
```html
<div class="animate-bounce">⚽</div>
```

---

## 🎯 Exemplos Práticos Completos

### Exemplo 1: Card Interativo

```html
<div class="
  group
  bg-white 
  rounded-lg 
  shadow-md 
  p-6
  cursor-pointer
  hover:shadow-xl
  hover:-translate-y-2
  transition-all
  duration-300
">
  <div class="w-12 h-12 bg-blue-500 rounded-full group-hover:bg-blue-600 transition-colors"></div>
  <h3 class="mt-4 text-xl font-bold text-gray-800 group-hover:text-blue-600 transition-colors">
    Título do Card
  </h3>
  <p class="mt-2 text-gray-600 group-hover:text-gray-800 transition-colors">
    Descrição do card que muda de cor no hover
  </p>
  <button class="
    mt-4 
    opacity-0 
    group-hover:opacity-100
    bg-blue-500 
    hover:bg-blue-600
    text-white 
    px-4 py-2 
    rounded
    transition-all
    duration-300
  ">
    Ver mais
  </button>
</div>
```

### Exemplo 2: Input com Label Flutuante

```html
<div class="relative">
  <input 
    type="text"
    id="username"
    class="
      peer
      w-full 
      px-4 
      pt-6 
      pb-2
      border-2 
      border-gray-300 
      rounded-lg
      focus:outline-none
      focus:border-blue-500
      focus:ring-2
      focus:ring-blue-200
      transition-all
    "
    placeholder=" "
  />
  <label 
    for="username"
    class="
      absolute 
      left-4 
      top-4
      text-gray-500
      transition-all
      peer-focus:text-blue-500
      peer-focus:top-2
      peer-focus:text-xs
      peer-placeholder-shown:top-4
      peer-placeholder-shown:text-base
      pointer-events-none
    "
  >
    Nome de usuário
  </label>
</div>
```

### Exemplo 3: Botão com Múltiplos Estados

```html
<button class="
  relative
  bg-gradient-to-r from-blue-500 to-purple-500
  hover:from-blue-600 hover:to-purple-600
  active:from-blue-700 active:to-purple-700
  text-white
  font-semibold
  px-8 py-3
  rounded-lg
  shadow-lg
  hover:shadow-xl
  hover:scale-105
  active:scale-95
  focus:outline-none
  focus:ring-4
  focus:ring-blue-300
  disabled:opacity-50
  disabled:cursor-not-allowed
  disabled:hover:scale-100
  transition-all
  duration-200
  ease-in-out
  overflow-hidden
">
  <span class="relative z-10">Clique em mim</span>
  <span class="
    absolute 
    inset-0 
    bg-white 
    opacity-0 
    hover:opacity-20
    transition-opacity
  "></span>
</button>
```

### Exemplo 4: Menu Dropdown com Animações

```html
<div class="relative group">
  <button class="
    px-4 py-2 
    bg-gray-100 
    rounded-lg
    hover:bg-gray-200
    focus:outline-none
    focus:ring-2
    focus:ring-blue-500
  ">
    Menu
  </button>
  <div class="
    absolute 
    top-full 
    left-0 
    mt-2
    w-48
    bg-white
    rounded-lg
    shadow-lg
    opacity-0
    invisible
    group-hover:opacity-100
    group-hover:visible
    group-hover:translate-y-0
    -translate-y-2
    transition-all
    duration-200
  ">
    <a href="#" class="block px-4 py-2 hover:bg-gray-100 first:rounded-t-lg last:rounded-b-lg">
      Item 1
    </a>
    <a href="#" class="block px-4 py-2 hover:bg-gray-100 first:rounded-t-lg last:rounded-b-lg">
      Item 2
    </a>
    <a href="#" class="block px-4 py-2 hover:bg-gray-100 first:rounded-t-lg last:rounded-b-lg">
      Item 3
    </a>
  </div>
</div>
```

---

## 📚 Resumo: Mapeamento Tailwind → CSS

| Tailwind | CSS Puro |
|----------|----------|
| `hover:bg-blue-600` | `:hover { background-color: rgb(37, 99, 235); }` |
| `focus:ring-2 focus:ring-blue-500` | `:focus { box-shadow: 0 0 0 2px rgba(59, 130, 246, 0.5); }` |
| `active:scale-95` | `:active { transform: scale(0.95); }` |
| `disabled:opacity-50` | `:disabled { opacity: 0.5; }` |
| `first:pt-0` | `:first-child { padding-top: 0; }` |
| `last:pb-0` | `:last-child { padding-bottom: 0; }` |
| `odd:bg-gray-100` | `:nth-child(odd) { background-color: rgb(243, 244, 246); }` |
| `group-hover:text-blue-600` | `.group:hover .element { color: rgb(37, 99, 235); }` |
| `peer-focus:border-blue-500` | `.peer:focus ~ .element { border-color: rgb(59, 130, 246); }` |
| `transition-all duration-200` | `transition: all 200ms;` |
| `hover:scale-110` | `:hover { transform: scale(1.1); }` |
| `animate-spin` | `animation: spin 1s linear infinite;` |

---

## 🎯 Pontos Importantes

1. **Mobile-First para Estados**: Lembre-se que os estados também seguem a filosofia mobile-first. Use breakpoints quando necessário: `md:hover:bg-blue-600`

2. **Performance**: Use `transition-colors` ou `transition-transform` ao invés de `transition-all` quando possível para melhor performance

3. **Acessibilidade**: Sempre forneça estados de focus visíveis para navegação por teclado. Use `focus-visible:` quando apropriado

4. **Group e Peer**: Essas utilities são poderosas para criar interações complexas sem JavaScript

5. **Transições**: Sempre especifique quais propriedades devem transicionar para evitar animações desnecessárias

6. **Estados Combinados**: Você pode combinar múltiplos estados: `hover:bg-blue-600 focus:ring-2 active:scale-95`

---

**Próximo Passo**: Na próxima aula, vamos aprender sobre **Componentes e Reutilização com @apply**, onde você aprenderá a criar componentes reutilizáveis combinando múltiplas classes utilitárias.

