# Aula 3 - Performance, Boas Práticas e Otimização: Layout com Tailwind

## 🚀 Performance

### Impacto do Display e Position no Rendering

#### Reflow e Repaint

Quando você muda `display` ou `position`, o navegador precisa recalcular o layout (reflow) e repintar (repaint) a página.

**Classes que causam reflow:**
- `block`, `inline`, `flex`, `grid` - Mudam o contexto de formatação
- `relative`, `absolute`, `fixed` - Mudam o posicionamento no fluxo

**Impacto na performance:**
```html
<!-- ❌ Múltiplas mudanças de display causam reflows -->
<div class="block inline flex">Mudanças constantes</div>

<!-- ✅ Defina display uma vez -->
<div class="flex">Display definido</div>
```

**Boas práticas:**
- Evite mudar `display` dinamicamente com JavaScript quando possível
- Use `transform` e `opacity` para animações (não causam reflow)
- Prefira `position: fixed` para elementos que precisam ficar fixos (melhor que mudar `position` dinamicamente)

---

### Performance do Flexbox

#### GPU Acceleration

Flexbox é otimizado pelo navegador, mas algumas propriedades são mais custosas:

**Propriedades custosas:**
- `justify-content: space-between` - Requer cálculo de espaçamento
- `flex-wrap: wrap` - Pode causar múltiplos reflows ao quebrar linhas
- `align-items: stretch` - Requer cálculo de altura

**Otimizações:**
```html
<!-- ✅ Use gap em vez de margin para evitar cálculos complexos -->
<div class="flex gap-4">
  <div>Item 1</div>
  <div>Item 2</div>
</div>

<!-- ❌ Evite margin complexo -->
<div class="flex">
  <div class="mr-4">Item 1</div>
  <div>Item 2</div>
</div>
```

#### Evitando Layout Thrashing

Layout thrashing acontece quando você força múltiplos reflows em sequência.

```javascript
// ❌ Ruim: Múltiplos reflows
element.classList.add('flex');
element.classList.add('items-center');
element.classList.add('justify-center');

// ✅ Bom: Uma única mudança
element.className = 'flex items-center justify-center';
```

---

### Position e Performance

#### Fixed vs Absolute

**`position: fixed`:**
- Renderizado em camada separada (compositor)
- Não causa reflow quando a página rola
- Melhor performance para elementos que ficam fixos

**`position: absolute`:**
- Pode causar reflow se o elemento pai mudar
- Melhor para elementos relativos a um container específico

```html
<!-- ✅ Fixed para headers que ficam sempre visíveis -->
<header class="fixed top-0 left-0 w-full z-50">
  Header fixo
</header>

<!-- ✅ Absolute para elementos dentro de containers -->
<div class="relative">
  <div class="absolute top-0 right-0">
    Badge
  </div>
</div>
```

#### Sticky e Performance

`position: sticky` pode ser custoso em listas longas:

```html
<!-- ⚠️ Cuidado com sticky em listas muito longas -->
<div class="space-y-4">
  <div class="sticky top-0">Header sticky</div>
  <!-- Muitos itens aqui podem causar problemas -->
</div>

<!-- ✅ Limite o uso de sticky -->
<div class="sticky top-0 z-10">
  Apenas um elemento sticky por vez
</div>
```

---

### Z-Index e Stacking Context

#### Gerenciando Z-Index

Z-index cria novos stacking contexts, o que pode afetar performance:

```html
<!-- ✅ Use valores padronizados do Tailwind -->
<div class="z-10">Camada 10</div>
<div class="z-20">Camada 20</div>
<div class="z-30">Camada 30</div>

<!-- ❌ Evite valores arbitrários muito altos -->
<div class="z-[9999]">Pode causar problemas</div>
```

**Hierarquia recomendada:**
- `z-10`: Dropdowns, tooltips
- `z-20`: Sticky headers
- `z-30`: Modals, overlays
- `z-40`: Popovers
- `z-50`: Notificações, toasts

---

## 🎯 Boas Práticas

### Organização de Classes

#### Ordem Recomendada

Organize classes Tailwind em uma ordem lógica para melhor legibilidade:

```html
<!-- Ordem sugerida: Layout → Espaçamento → Aparência → Interatividade -->

<!-- ✅ Boa organização -->
<div class="
  flex items-center justify-between
  p-6 gap-4
  bg-white rounded-lg shadow-md
  hover:shadow-lg transition-shadow
">
  Conteúdo
</div>

<!-- ❌ Desorganizado -->
<div class="bg-white flex p-6 hover:shadow-lg items-center rounded-lg justify-between shadow-md gap-4 transition-shadow">
  Conteúdo
</div>
```

**Ordem sugerida:**
1. Display e Layout (`flex`, `grid`, `block`)
2. Position (`relative`, `absolute`, `fixed`)
3. Flexbox/Grid (`flex-row`, `justify-center`, `items-center`)
4. Espaçamento (`p-*`, `m-*`, `gap-*`)
5. Dimensões (`w-*`, `h-*`)
6. Tipografia (`text-*`, `font-*`)
7. Cores e Backgrounds (`bg-*`, `text-*`)
8. Bordas e Sombras (`border-*`, `rounded-*`, `shadow-*`)
9. Estados (`hover:`, `focus:`, `active:`)
10. Transições (`transition-*`)

#### Quebra de Linhas

Para classes longas, quebre em múltiplas linhas:

```html
<!-- ✅ Legível -->
<button class="
  flex items-center justify-center
  px-6 py-3
  bg-blue-500 text-white
  rounded-lg shadow-md
  hover:bg-blue-600 hover:shadow-lg
  transition-all duration-200
">
  Clique aqui
</button>

<!-- ❌ Difícil de ler -->
<button class="flex items-center justify-center px-6 py-3 bg-blue-500 text-white rounded-lg shadow-md hover:bg-blue-600 hover:shadow-lg transition-all duration-200">
  Clique aqui
</button>
```

---

### Nomenclatura e Estrutura

#### Quando Criar Componentes vs Usar Utilities

**Use utilities diretamente quando:**
- O componente é usado apenas uma vez
- O layout é simples e não se repete
- Não há variações complexas

```html
<!-- ✅ Utilities diretas para uso único -->
<div class="flex items-center gap-4 p-4">
  Conteúdo único
</div>
```

**Crie componentes (com @apply) quando:**
- O padrão se repete muitas vezes
- Há variações que precisam ser gerenciadas
- O código fica muito verboso

```css
/* ✅ Componente para padrão repetido */
.btn-primary {
  @apply flex items-center justify-center
         px-6 py-3
         bg-blue-500 text-white
         rounded-lg shadow-md
         hover:bg-blue-600
         transition-all;
}
```

---

### Responsividade e Mobile-First

#### Abordagem Mobile-First

Sempre comece com o layout mobile e adicione breakpoints para telas maiores:

```html
<!-- ✅ Mobile-first: começa vertical, vira horizontal em md -->
<div class="flex flex-col md:flex-row gap-4">
  <div>Item 1</div>
  <div>Item 2</div>
</div>

<!-- ❌ Desktop-first: pode quebrar em mobile -->
<div class="flex flex-row flex-col md:flex-row gap-4">
  <div>Item 1</div>
  <div>Item 2</div>
</div>
```

#### Breakpoints Consistentes

Use os breakpoints padrão do Tailwind de forma consistente:

```html
<!-- ✅ Breakpoints consistentes -->
<div class="
  flex flex-col
  md:flex-row md:justify-between
  lg:gap-8
">
  Conteúdo
</div>
```

**Breakpoints do Tailwind:**
- `sm:` - 640px
- `md:` - 768px
- `lg:` - 1024px
- `xl:` - 1280px
- `2xl:` - 1536px

---

### Acessibilidade

#### Ordem de Foco

A ordem visual pode diferir da ordem do DOM. Garanta que a ordem de foco faça sentido:

```html
<!-- ⚠️ Cuidado: ordem visual vs ordem de foco -->
<div class="flex flex-row-reverse">
  <button>Botão 1</button>
  <button>Botão 2</button>
</div>
<!-- Foco ainda vai na ordem do HTML, não visual -->

<!-- ✅ Se necessário, reordene no HTML -->
<div class="flex">
  <button>Botão 2</button>
  <button>Botão 1</button>
</div>
```

#### Contraste e Visibilidade

Garanta que elementos posicionados não quebrem o contraste:

```html
<!-- ✅ Overlay com contraste adequado -->
<div class="fixed inset-0 bg-black bg-opacity-75 z-40">
  <!-- Conteúdo com bom contraste -->
</div>

<!-- ❌ Overlay muito transparente -->
<div class="fixed inset-0 bg-black bg-opacity-10 z-40">
  <!-- Pode não ter contraste suficiente -->
</div>
```

#### Foco Visível

Elementos focáveis devem ter estados de foco visíveis:

```html
<!-- ✅ Foco visível -->
<button class="
  px-4 py-2 bg-blue-500 text-white rounded
  focus:outline-none focus:ring-2 focus:ring-blue-300
">
  Botão
</button>

<!-- ❌ Sem foco visível -->
<button class="px-4 py-2 bg-blue-500 text-white rounded">
  Botão
</button>
```

---

## 🔧 Otimização

### PurgeCSS e Tree-Shaking

#### Configuração de Content

Garanta que o Tailwind encontre todas as classes usadas:

```javascript
// tailwind.config.js
module.exports = {
  content: [
    './src/**/*.{html,js,jsx,ts,tsx}',
    './public/**/*.html',
  ],
  // ...
}
```

**Por que importa:**
- Classes não encontradas são removidas do CSS final
- Reduz drasticamente o tamanho do arquivo CSS
- Melhora performance de carregamento

#### Classes Dinâmicas

Cuidado com classes geradas dinamicamente:

```javascript
// ⚠️ Tailwind pode não detectar classes dinâmicas
const color = 'blue';
const className = `bg-${color}-500`; // Não funciona!

// ✅ Use classes completas ou safelist
const className = color === 'blue' ? 'bg-blue-500' : 'bg-red-500';

// Ou configure safelist
module.exports = {
  safelist: [
    'bg-blue-500',
    'bg-red-500',
    // ...
  ]
}
```

---

### CSS Crítico

#### Identificando CSS Crítico

Para layouts acima da dobra (above the fold), extraia CSS crítico:

```html
<!-- ✅ CSS crítico inline -->
<style>
  .header { display: flex; align-items: center; }
  .hero { position: relative; }
</style>

<!-- CSS completo carregado depois -->
<link rel="stylesheet" href="styles.css">
```

**Ferramentas:**
- `critical` (npm package)
- `purgecss` com opção de extrair crítico
- Build tools como Vite, Webpack com plugins

---

### Minimizando Reflows

#### Usando Transform em vez de Position

Para animações, prefira `transform`:

```html
<!-- ✅ Transform não causa reflow -->
<div class="transform translate-x-4 transition-transform">
  Elemento animado
</div>

<!-- ❌ Mudar left/right causa reflow -->
<div class="relative left-4 transition-all">
  Elemento animado
</div>
```

#### Will-Change

Para elementos que vão animar, use `will-change`:

```css
/* Adicione via @apply ou CSS customizado */
.animated-element {
  @apply transform transition-transform;
  will-change: transform;
}
```

---

## 🎨 Padrões Comuns

### Layout de Header

```html
<!-- Padrão recomendado -->
<header class="
  sticky top-0 z-20
  flex items-center justify-between
  px-4 md:px-6 lg:px-8
  h-16 md:h-20
  bg-white shadow-sm
  backdrop-blur-sm
">
  <div class="flex items-center gap-4">
    <img src="logo.png" alt="Logo" class="h-8 md:h-10">
    <nav class="hidden md:flex gap-6">
      <!-- Menu -->
    </nav>
  </div>
  <button class="px-4 py-2 bg-blue-500 text-white rounded">
    Ação
  </button>
</header>
```

**Características:**
- Sticky para ficar visível ao rolar
- Responsivo com breakpoints
- Backdrop blur para efeito moderno
- Z-index apropriado

---

### Card com Layout Flex

```html
<!-- Padrão recomendado -->
<div class="
  flex flex-col
  gap-4 p-6
  bg-white rounded-lg shadow-md
  hover:shadow-lg transition-shadow
  max-w-sm
">
  <img src="image.jpg" alt="Imagem" class="w-full rounded">
  <div class="flex flex-col gap-2 flex-1">
    <h3 class="text-xl font-bold">Título</h3>
    <p class="text-gray-600">Descrição</p>
  </div>
  <div class="flex justify-between items-center mt-auto">
    <span class="text-2xl font-bold">R$ 99,90</span>
    <button class="px-4 py-2 bg-blue-500 text-white rounded">
      Comprar
    </button>
  </div>
</div>
```

**Características:**
- `flex-col` para layout vertical
- `flex-1` no conteúdo para empurrar rodapé
- `mt-auto` no rodapé para fixar na base
- Transições suaves

---

### Modal Centralizado

```html
<!-- Padrão recomendado -->
<!-- Overlay -->
<div class="
  fixed inset-0
  bg-black bg-opacity-50
  z-40
  flex items-center justify-center
  p-4
">
  <!-- Modal -->
  <div class="
    relative
    w-full max-w-md
    bg-white rounded-lg shadow-xl
    p-6
    z-50
  ">
    <!-- Botão fechar -->
    <button class="
      absolute top-4 right-4
      text-gray-400 hover:text-gray-600
    ">
      ×
    </button>
    <!-- Conteúdo -->
    <div>Conteúdo do modal</div>
  </div>
</div>
```

**Características:**
- Overlay com `fixed inset-0`
- Modal centralizado com `flex items-center justify-center`
- Z-index apropriado
- Responsivo com `max-w-md` e `p-4`

---

### Sidebar com Layout

```html
<!-- Padrão recomendado -->
<div class="flex h-screen">
  <!-- Sidebar -->
  <aside class="
    hidden md:flex flex-col
    w-64
    bg-gray-800 text-white
    p-4
  ">
    <h2 class="text-xl font-bold mb-4">Menu</h2>
    <nav class="flex flex-col gap-2">
      <!-- Itens do menu -->
    </nav>
  </aside>
  
  <!-- Conteúdo -->
  <main class="
    flex-1
    p-4 md:p-8
    bg-gray-50
    overflow-y-auto
  ">
    Conteúdo principal
  </main>
</div>
```

**Características:**
- Sidebar oculta em mobile (`hidden md:flex`)
- Conteúdo principal com `flex-1` para ocupar espaço
- `overflow-y-auto` para scroll quando necessário
- Altura total da tela (`h-screen`)

---

## ⚠️ O que NÃO Fazer

### Evite Múltiplas Mudanças de Display

```html
<!-- ❌ Ruim: Múltiplas mudanças -->
<div class="block inline flex grid">
  Conteúdo
</div>

<!-- ✅ Bom: Um display por vez -->
<div class="flex">
  Conteúdo
</div>
```

### Evite Z-Index Muito Altos

```html
<!-- ❌ Ruim: Z-index excessivo -->
<div class="z-[9999]">Elemento</div>

<!-- ✅ Bom: Use valores padronizados -->
<div class="z-50">Elemento</div>
```

### Evite Position sem Contexto

```html
<!-- ❌ Ruim: Absolute sem relative no pai -->
<div>
  <div class="absolute top-0">Pode não funcionar como esperado</div>
</div>

<!-- ✅ Bom: Container relative -->
<div class="relative">
  <div class="absolute top-0">Funciona corretamente</div>
</div>
```

### Evite Flexbox Desnecessário

```html
<!-- ❌ Ruim: Flex para elemento único -->
<div class="flex">
  <p>Um único parágrafo</p>
</div>

<!-- ✅ Bom: Block é suficiente -->
<div>
  <p>Um único parágrafo</p>
</div>
```

---

## 📊 Métricas de Performance

### Como Medir

**Ferramentas:**
- Chrome DevTools - Performance tab
- Lighthouse - Layout shifts, render blocking
- WebPageTest - Análise detalhada

**Métricas importantes:**
- **First Contentful Paint (FCP)**: Quando o primeiro conteúdo aparece
- **Largest Contentful Paint (LCP)**: Quando o maior elemento aparece
- **Cumulative Layout Shift (CLS)**: Estabilidade visual
- **Time to Interactive (TTI)**: Quando a página fica interativa

### Otimizações Específicas

**Para Layout:**
- Minimize reflows usando `transform` em vez de `position`
- Use `will-change` para elementos que vão animar
- Evite mudanças dinâmicas de `display`
- Prefira `gap` em vez de `margin` para espaçamento

**Para Tailwind:**
- Configure PurgeCSS corretamente
- Use JIT mode para desenvolvimento
- Monitore o tamanho do CSS final
- Extraia CSS crítico quando possível

---

## 🎯 Resumo de Boas Práticas

### Layout
- ✅ Use mobile-first approach
- ✅ Organize classes em ordem lógica
- ✅ Quebre linhas para classes longas
- ✅ Prefira `gap` em vez de `margin` para espaçamento

### Performance
- ✅ Minimize reflows usando `transform`
- ✅ Configure PurgeCSS corretamente
- ✅ Use valores padronizados de z-index
- ✅ Evite mudanças dinâmicas de display

### Acessibilidade
- ✅ Garanta ordem de foco lógica
- ✅ Adicione estados de foco visíveis
- ✅ Mantenha contraste adequado
- ✅ Teste com leitores de tela

### Manutenibilidade
- ✅ Use componentes para padrões repetidos
- ✅ Mantenha consistência de breakpoints
- ✅ Documente decisões de layout complexas
- ✅ Revise e refatore regularmente

---

## 🚀 Próximos Passos

Agora que você domina as boas práticas de layout com Tailwind:

1. **Aplique em projetos reais** - Use os padrões aprendidos
2. **Meça performance** - Use as ferramentas mencionadas
3. **Refatore código existente** - Aplique as otimizações
4. **Mantenha consistência** - Siga os padrões estabelecidos

Na próxima aula, exploraremos **CSS Grid com Tailwind**, que oferece controle bidimensional ainda mais poderoso!

