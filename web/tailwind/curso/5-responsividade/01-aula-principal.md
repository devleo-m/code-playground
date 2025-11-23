# Aula 5: Responsividade com Tailwind CSS - Conteúdo Principal

## 📖 Revisão: O que é Responsividade?

Antes de mergulharmos na responsividade com Tailwind, vamos relembrar o conceito fundamental que você já conhece de CSS puro:

**Responsividade** é a capacidade de um site ou aplicação se adaptar a diferentes tamanhos de tela e dispositivos, proporcionando uma experiência otimizada em desktops, tablets e smartphones.

### Responsividade em CSS Puro

Você já aprendeu que em CSS puro usamos **media queries** para criar designs responsivos:

```css
/* CSS Puro - Desktop First */
.container {
  width: 1200px;
  padding: 2rem;
}

@media (max-width: 768px) {
  .container {
    width: 100%;
    padding: 1rem;
  }
}

/* CSS Puro - Mobile First */
.container {
  width: 100%;
  padding: 1rem;
}

@media (min-width: 768px) {
  .container {
    width: 1200px;
    padding: 2rem;
  }
}
```

### Responsividade no Tailwind

O Tailwind CSS adota uma abordagem **mobile-first** e fornece um sistema de **breakpoints** pré-configurados que você pode usar diretamente nas classes utilitárias através de **prefixos responsivos**.

---

## 🎯 Filosofia Mobile-First do Tailwind

### O que é Mobile-First?

**Mobile-first** é uma abordagem de design onde você começa estilizando para dispositivos móveis (telas pequenas) e depois adiciona estilos para telas maiores usando media queries com `min-width`.

### Por que Mobile-First?

1. **Performance**: Dispositivos móveis geralmente têm menos poder de processamento
2. **Priorização**: A maioria dos usuários acessa sites via mobile
3. **Progressive Enhancement**: Adiciona recursos conforme o dispositivo suporta
4. **Simplicidade**: Começar simples e adicionar complexidade é mais fácil

### Como Funciona no Tailwind

No Tailwind, **por padrão, todas as classes aplicam estilos para mobile** (telas pequenas). Para aplicar estilos em telas maiores, você usa **prefixos responsivos**.

**Exemplo:**
```html
<!-- Padding de 1rem em mobile, 2rem em tablets, 3rem em desktop -->
<div class="p-4 md:p-8 lg:p-12">
  Conteúdo
</div>
```

**Equivalente em CSS puro:**
```css
div {
  padding: 1rem; /* Mobile - padrão */
}

@media (min-width: 768px) {
  div {
    padding: 2rem; /* Tablet */
  }
}

@media (min-width: 1024px) {
  div {
    padding: 3rem; /* Desktop */
  }
}
```

---

## 📱 Breakpoints Padrão do Tailwind

O Tailwind vem com 5 breakpoints pré-configurados:

| Prefixo | Breakpoint | Descrição |
|---------|-----------|-----------|
| (sem prefixo) | < 640px | Mobile (padrão) |
| `sm:` | ≥ 640px | Smartphones grandes |
| `md:` | ≥ 768px | Tablets |
| `lg:` | ≥ 1024px | Desktops |
| `xl:` | ≥ 1280px | Desktops grandes |
| `2xl:` | ≥ 1536px | Telas muito grandes |

### Valores Exatos dos Breakpoints

```javascript
// tailwind.config.js (valores padrão)
screens: {
  'sm': '640px',   // min-width: 640px
  'md': '768px',   // min-width: 768px
  'lg': '1024px',  // min-width: 1024px
  'xl': '1280px',  // min-width: 1280px
  '2xl': '1536px', // min-width: 1536px
}
```

### Como Usar Prefixos Responsivos

A sintaxe é simples: **`[breakpoint]:[classe]`**

```html
<!-- Exemplo básico -->
<div class="text-sm md:text-base lg:text-lg xl:text-xl">
  Texto responsivo
</div>
```

**Tradução para CSS:**
```css
div {
  font-size: 0.875rem; /* text-sm - mobile */
}

@media (min-width: 768px) {
  div {
    font-size: 1rem; /* text-base - tablet */
  }
}

@media (min-width: 1024px) {
  div {
    font-size: 1.125rem; /* text-lg - desktop */
  }
}

@media (min-width: 1280px) {
  div {
    font-size: 1.25rem; /* text-xl - desktop grande */
  }
}
```

---

## 🎨 Responsividade em Diferentes Propriedades

### 1. Espaçamento Responsivo

#### Padding Responsivo

```html
<!-- Padding que aumenta conforme a tela cresce -->
<div class="p-2 sm:p-4 md:p-6 lg:p-8">
  Conteúdo com padding responsivo
</div>
```

**CSS equivalente:**
```css
div {
  padding: 0.5rem; /* p-2 */
}

@media (min-width: 640px) {
  div {
    padding: 1rem; /* p-4 */
  }
}

@media (min-width: 768px) {
  div {
    padding: 1.5rem; /* p-6 */
  }
}

@media (min-width: 1024px) {
  div {
    padding: 2rem; /* p-8 */
  }
}
```

#### Margin Responsivo

```html
<!-- Margin que muda em diferentes breakpoints -->
<div class="m-2 md:m-4 lg:m-8">
  Conteúdo
</div>
```

#### Gap Responsivo (Flexbox/Grid)

```html
<div class="flex gap-2 md:gap-4 lg:gap-8">
  <div>Item 1</div>
  <div>Item 2</div>
  <div>Item 3</div>
</div>
```

### 2. Tipografia Responsiva

#### Tamanho de Fonte

```html
<!-- Texto que cresce em telas maiores -->
<h1 class="text-2xl sm:text-3xl md:text-4xl lg:text-5xl xl:text-6xl">
  Título Responsivo
</h1>
```

**CSS equivalente:**
```css
h1 {
  font-size: 1.5rem; /* text-2xl */
}

@media (min-width: 640px) {
  h1 {
    font-size: 1.875rem; /* text-3xl */
  }
}

@media (min-width: 768px) {
  h1 {
    font-size: 2.25rem; /* text-4xl */
  }
}

@media (min-width: 1024px) {
  h1 {
    font-size: 3rem; /* text-5xl */
  }
}

@media (min-width: 1280px) {
  h1 {
    font-size: 3.75rem; /* text-6xl */
  }
}
```

#### Peso da Fonte

```html
<p class="font-normal md:font-semibold lg:font-bold">
  Texto com peso responsivo
</p>
```

#### Line Height

```html
<p class="leading-tight md:leading-normal lg:leading-relaxed">
  Texto com line-height responsivo
</p>
```

### 3. Layout Responsivo

#### Display Responsivo

```html
<!-- Stack vertical em mobile, horizontal em desktop -->
<div class="flex flex-col md:flex-row">
  <div>Item 1</div>
  <div>Item 2</div>
</div>
```

**CSS equivalente:**
```css
div {
  display: flex;
  flex-direction: column; /* flex-col */
}

@media (min-width: 768px) {
  div {
    flex-direction: row; /* flex-row */
  }
}
```

#### Grid Responsivo

```html
<!-- 1 coluna em mobile, 2 em tablet, 3 em desktop -->
<div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
  <div>Item 1</div>
  <div>Item 2</div>
  <div>Item 3</div>
</div>
```

**CSS equivalente:**
```css
div {
  display: grid;
  grid-template-columns: repeat(1, minmax(0, 1fr)); /* grid-cols-1 */
  gap: 1rem;
}

@media (min-width: 768px) {
  div {
    grid-template-columns: repeat(2, minmax(0, 1fr)); /* grid-cols-2 */
  }
}

@media (min-width: 1024px) {
  div {
    grid-template-columns: repeat(3, minmax(0, 1fr)); /* grid-cols-3 */
  }
}
```

#### Width e Height Responsivos

```html
<!-- Largura que se adapta -->
<div class="w-full md:w-1/2 lg:w-1/3">
  Conteúdo
</div>
```

```html
<!-- Altura responsiva -->
<div class="h-32 md:h-48 lg:h-64">
  Conteúdo
</div>
```

### 4. Cores e Backgrounds Responsivos

```html
<!-- Background que muda de cor em diferentes telas -->
<div class="bg-blue-500 md:bg-green-500 lg:bg-purple-500">
  Background responsivo
</div>
```

**CSS equivalente:**
```css
div {
  background-color: rgb(59 130 246); /* bg-blue-500 */
}

@media (min-width: 768px) {
  div {
    background-color: rgb(34 197 94); /* bg-green-500 */
  }
}

@media (min-width: 1024px) {
  div {
    background-color: rgb(168 85 247); /* bg-purple-500 */
  }
}
```

### 5. Visibilidade Responsiva

```html
<!-- Esconder em mobile, mostrar em desktop -->
<div class="hidden md:block">
  Visível apenas em tablets e desktops
</div>

<!-- Mostrar em mobile, esconder em desktop -->
<div class="block md:hidden">
  Visível apenas em mobile
</div>
```

**CSS equivalente:**
```css
/* Esconder em mobile, mostrar em desktop */
.hidden-md-up {
  display: none;
}

@media (min-width: 768px) {
  .hidden-md-up {
    display: block;
  }
}

/* Mostrar em mobile, esconder em desktop */
.visible-mobile {
  display: block;
}

@media (min-width: 768px) {
  .visible-mobile {
    display: none;
  }
}
```

---

## 🔄 Múltiplos Breakpoints na Mesma Propriedade

Você pode aplicar diferentes valores em múltiplos breakpoints:

```html
<!-- Padding que muda em cada breakpoint -->
<div class="p-2 sm:p-4 md:p-6 lg:p-8 xl:p-10">
  Conteúdo
</div>
```

**Ordem Importante:**
- Sempre comece com o valor mobile (sem prefixo)
- Depois adicione breakpoints em ordem crescente
- O Tailwind aplica estilos em cascata (mobile → sm → md → lg → xl → 2xl)

---

## 📐 Exemplos Práticos de Layouts Responsivos

### Exemplo 1: Card de Produto Responsivo

```html
<div class="bg-white rounded-lg shadow-md p-4 md:p-6 lg:p-8">
  <img 
    src="produto.jpg" 
    alt="Produto"
    class="w-full h-48 md:h-64 lg:h-80 object-cover rounded-md mb-4"
  >
  <h3 class="text-lg md:text-xl lg:text-2xl font-bold mb-2">
    Nome do Produto
  </h3>
  <p class="text-sm md:text-base text-gray-600 mb-4">
    Descrição do produto que se adapta ao tamanho da tela.
  </p>
  <div class="flex flex-col sm:flex-row gap-2 sm:gap-4">
    <span class="text-xl md:text-2xl font-bold">R$ 99,90</span>
    <button class="px-4 py-2 bg-blue-500 text-white rounded hover:bg-blue-600">
      Comprar
    </button>
  </div>
</div>
```

**Características:**
- Padding aumenta em telas maiores
- Imagem cresce proporcionalmente
- Título aumenta de tamanho
- Botão e preço mudam de layout (coluna → linha)

### Exemplo 2: Grid de Cards Responsivo

```html
<div class="container mx-auto px-4 py-8">
  <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-4 md:gap-6">
    <div class="bg-white p-4 rounded shadow">Card 1</div>
    <div class="bg-white p-4 rounded shadow">Card 2</div>
    <div class="bg-white p-4 rounded shadow">Card 3</div>
    <div class="bg-white p-4 rounded shadow">Card 4</div>
  </div>
</div>
```

**Comportamento:**
- Mobile: 1 coluna
- Tablet (sm): 2 colunas
- Desktop (lg): 3 colunas
- Desktop grande (xl): 4 colunas
- Gap aumenta em telas maiores

### Exemplo 3: Navegação Responsiva

```html
<nav class="bg-blue-600 text-white">
  <div class="container mx-auto px-4 py-4">
    <div class="flex flex-col md:flex-row justify-between items-center gap-4">
      <!-- Logo -->
      <div class="text-xl font-bold">Logo</div>
      
      <!-- Menu Mobile (hamburger) - visível apenas em mobile -->
      <button class="md:hidden">
        <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h16"></path>
        </svg>
      </button>
      
      <!-- Menu Desktop - visível apenas em desktop -->
      <div class="hidden md:flex gap-6">
        <a href="#" class="hover:text-blue-200">Home</a>
        <a href="#" class="hover:text-blue-200">Sobre</a>
        <a href="#" class="hover:text-blue-200">Contato</a>
      </div>
    </div>
  </div>
</nav>
```

**Características:**
- Layout muda de coluna (mobile) para linha (desktop)
- Menu hamburger aparece apenas em mobile
- Menu horizontal aparece apenas em desktop

### Exemplo 4: Container Responsivo

```html
<div class="container mx-auto px-4 sm:px-6 md:px-8 lg:px-12">
  <div class="max-w-7xl mx-auto">
    <h1 class="text-3xl md:text-4xl lg:text-5xl mb-8">
      Título Principal
    </h1>
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
      <!-- Conteúdo -->
    </div>
  </div>
</div>
```

**Características:**
- Padding do container aumenta progressivamente
- Título cresce em cada breakpoint
- Grid se adapta de 1 → 2 → 3 colunas

---

## 🎛️ Breakpoints Customizados

Você pode criar breakpoints customizados no `tailwind.config.js`:

```javascript
// tailwind.config.js
module.exports = {
  theme: {
    extend: {
      screens: {
        'xs': '475px',      // Breakpoint extra pequeno
        '3xl': '1920px',    // Breakpoint extra grande
        'tablet': '768px',  // Breakpoint com nome customizado
        'desktop': '1024px',
      },
    },
  },
}
```

**Uso:**
```html
<div class="p-4 xs:p-6 tablet:p-8 desktop:p-10 3xl:p-12">
  Conteúdo
</div>
```

### Substituindo Breakpoints Padrão

Se você quiser **substituir** os breakpoints padrão (não estender):

```javascript
// tailwind.config.js
module.exports = {
  theme: {
    screens: {
      'sm': '640px',
      'md': '768px',
      'lg': '1024px',
      'xl': '1280px',
      '2xl': '1536px',
    },
  },
}
```

**⚠️ Atenção:** Isso remove os breakpoints padrão. Use `extend` para adicionar sem remover.

---

## 🔀 Breakpoints com Valores Customizados

Você também pode usar breakpoints com valores específicos diretamente:

```html
<!-- Usando valores arbitrários (Tailwind 3.0+) -->
<div class="p-4 [@media(min-width:500px)]:p-6 [@media(min-width:900px)]:p-8">
  Conteúdo
</div>
```

**Nota:** Isso é menos comum e geralmente não é recomendado. Prefira usar breakpoints configurados.

---

## 📱 Container Responsivo

O Tailwind fornece uma classe `.container` que se adapta automaticamente:

```html
<div class="container mx-auto">
  Conteúdo centralizado e responsivo
</div>
```

**Comportamento:**
- Mobile: largura total (100%)
- sm (640px): max-width: 640px
- md (768px): max-width: 768px
- lg (1024px): max-width: 1024px
- xl (1280px): max-width: 1280px
- 2xl (1536px): max-width: 1536px

### Customizando o Container

```javascript
// tailwind.config.js
module.exports = {
  theme: {
    container: {
      center: true,        // mx-auto automático
      padding: '2rem',     // padding padrão
      screens: {
        'sm': '640px',
        'md': '768px',
        'lg': '1024px',
        'xl': '1280px',
      },
    },
  },
}
```

---

## 🎨 Responsividade com Variantes

Você pode combinar breakpoints com outras variantes do Tailwind:

### Hover Responsivo

```html
<button class="bg-blue-500 hover:bg-blue-600 md:hover:bg-blue-700">
  Botão
</button>
```

### Focus Responsivo

```html
<input class="border focus:border-blue-500 md:focus:ring-2">
```

### Group Hover Responsivo

```html
<div class="group">
  <div class="opacity-50 group-hover:opacity-100 md:group-hover:scale-105">
    Conteúdo
  </div>
</div>
```

---

## 🔍 Debugging Responsividade

### Usando DevTools

1. **Abra o DevTools** (F12)
2. **Ative o modo responsivo** (Ctrl+Shift+M / Cmd+Shift+M)
3. **Teste diferentes breakpoints**
4. **Inspecione as classes aplicadas**

### Verificando Media Queries

No DevTools, você pode ver as media queries geradas:

```css
/* CSS gerado pelo Tailwind */
@media (min-width: 768px) {
  .md\:p-8 {
    padding: 2rem;
  }
}
```

---

## 📊 Tabela de Referência Rápida

### Breakpoints

| Prefixo | Tamanho | Uso Comum |
|---------|---------|-----------|
| (padrão) | < 640px | Mobile |
| `sm:` | ≥ 640px | Smartphones grandes |
| `md:` | ≥ 768px | Tablets |
| `lg:` | ≥ 1024px | Desktops |
| `xl:` | ≥ 1280px | Desktops grandes |
| `2xl:` | ≥ 1536px | Telas muito grandes |

### Exemplos de Uso

```html
<!-- Espaçamento -->
<div class="p-2 md:p-4 lg:p-8"></div>

<!-- Tipografia -->
<h1 class="text-2xl md:text-4xl lg:text-6xl"></h1>

<!-- Layout -->
<div class="flex-col md:flex-row"></div>

<!-- Grid -->
<div class="grid-cols-1 md:grid-cols-2 lg:grid-cols-3"></div>

<!-- Visibilidade -->
<div class="hidden md:block"></div>
```

---

## 🎯 Resumo da Aula

### Conceitos Principais

1. **Mobile-First**: Tailwind usa abordagem mobile-first por padrão
2. **Prefixos Responsivos**: Use `sm:`, `md:`, `lg:`, `xl:`, `2xl:` para aplicar estilos em breakpoints específicos
3. **Breakpoints Padrão**: 5 breakpoints pré-configurados (640px, 768px, 1024px, 1280px, 1536px)
4. **Sintaxe**: `[breakpoint]:[classe]` (ex: `md:p-8`)
5. **Ordem**: Sempre comece com mobile (sem prefixo) e adicione breakpoints em ordem crescente

### Mapeamento CSS → Tailwind

| CSS Puro | Tailwind |
|----------|----------|
| `@media (min-width: 768px) { padding: 2rem; }` | `md:p-8` |
| `@media (min-width: 1024px) { font-size: 1.5rem; }` | `lg:text-2xl` |
| `@media (min-width: 768px) { display: flex; }` | `md:flex` |
| `@media (min-width: 1024px) { grid-template-columns: repeat(3, 1fr); }` | `lg:grid-cols-3` |

### Próximos Passos

Na próxima aula, você aprenderá sobre **Estados e Interatividade** no Tailwind, incluindo hover, focus, active e transições responsivas.

---

**Bons estudos! 🚀**

