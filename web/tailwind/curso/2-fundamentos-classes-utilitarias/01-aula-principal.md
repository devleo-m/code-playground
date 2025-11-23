# Aula 2: Fundamentos do Sistema de Classes Utilitárias - Conteúdo Principal

## 📖 Introdução

Agora que você entende a filosofia do Tailwind CSS, vamos mergulhar no **sistema de classes utilitárias** que é o coração do framework. Cada classe Tailwind mapeia diretamente para propriedades CSS que você já conhece, mas com uma sintaxe consistente e padronizada.

Nesta aula, você aprenderá:
- Sistema de espaçamento (padding, margin, gap)
- Sistema de cores e backgrounds
- Tipografia e estilos de texto
- Bordas, arredondamento e sombras
- Opacidade e visibilidade

---

## 📏 Sistema de Espaçamento

O Tailwind usa um **sistema de espaçamento padronizado** baseado em múltiplos de `0.25rem` (4px). Isso garante consistência visual em todo o projeto.

### Escala de Espaçamento Padrão

| Valor Tailwind | Valor em rem | Valor em px | Equivalente CSS |
|----------------|--------------|-------------|-----------------|
| `0` | `0` | `0px` | `0` |
| `0.5` | `0.125rem` | `2px` | `0.125rem` |
| `1` | `0.25rem` | `4px` | `0.25rem` |
| `2` | `0.5rem` | `8px` | `0.5rem` |
| `3` | `0.75rem` | `12px` | `0.75rem` |
| `4` | `1rem` | `16px` | `1rem` |
| `5` | `1.25rem` | `20px` | `1.25rem` |
| `6` | `1.5rem` | `24px` | `1.5rem` |
| `8` | `2rem` | `32px` | `2rem` |
| `10` | `2.5rem` | `40px` | `2.5rem` |
| `12` | `3rem` | `48px` | `3rem` |
| `16` | `4rem` | `64px` | `4rem` |
| `20` | `5rem` | `80px` | `5rem` |
| `24` | `6rem` | `96px` | `6rem` |

### Padding (Preenchimento Interno)

**Sintaxe:** `p-{valor}` (todos os lados) ou direções específicas

#### Padding em Todos os Lados

```html
<div class="p-4">Padding de 1rem em todos os lados</div>
```

**Equivalente CSS:**
```css
div {
  padding: 1rem;
}
```

#### Padding por Direção

| Classe Tailwind | Propriedade CSS | Descrição |
|-----------------|-----------------|-----------|
| `p-{valor}` | `padding: {valor}` | Todos os lados |
| `px-{valor}` | `padding-left: {valor}; padding-right: {valor}` | Horizontal (esquerda e direita) |
| `py-{valor}` | `padding-top: {valor}; padding-bottom: {valor}` | Vertical (topo e fundo) |
| `pt-{valor}` | `padding-top: {valor}` | Topo |
| `pr-{valor}` | `padding-right: {valor}` | Direita |
| `pb-{valor}` | `padding-bottom: {valor}` | Fundo |
| `pl-{valor}` | `padding-left: {valor}` | Esquerda |
| `ps-{valor}` | `padding-inline-start: {valor}` | Início do texto (RTL support) |
| `pe-{valor}` | `padding-inline-end: {valor}` | Fim do texto (RTL support) |

**Exemplos práticos:**

```html
<!-- Padding horizontal de 1rem, vertical de 0.5rem -->
<div class="px-4 py-2">Conteúdo</div>
```

**Equivalente CSS:**
```css
div {
  padding-left: 1rem;
  padding-right: 1rem;
  padding-top: 0.5rem;
  padding-bottom: 0.5rem;
}
```

```html
<!-- Padding apenas no topo -->
<div class="pt-8">Conteúdo</div>
```

**Equivalente CSS:**
```css
div {
  padding-top: 2rem;
}
```

### Margin (Espaçamento Externo)

**Sintaxe:** `m-{valor}` (todos os lados) ou direções específicas

#### Margin em Todos os Lados

```html
<div class="m-4">Margin de 1rem em todos os lados</div>
```

**Equivalente CSS:**
```css
div {
  margin: 1rem;
}
```

#### Margin por Direção

| Classe Tailwind | Propriedade CSS | Descrição |
|-----------------|-----------------|-----------|
| `m-{valor}` | `margin: {valor}` | Todos os lados |
| `mx-{valor}` | `margin-left: {valor}; margin-right: {valor}` | Horizontal |
| `my-{valor}` | `margin-top: {valor}; margin-bottom: {valor}` | Vertical |
| `mt-{valor}` | `margin-top: {valor}` | Topo |
| `mr-{valor}` | `margin-right: {valor}` | Direita |
| `mb-{valor}` | `margin-bottom: {valor}` | Fundo |
| `ml-{valor}` | `margin-left: {valor}` | Esquerda |
| `ms-{valor}` | `margin-inline-start: {valor}` | Início do texto |
| `me-{valor}` | `margin-inline-end: {valor}` | Fim do texto |

**Exemplos práticos:**

```html
<!-- Margin automático (centralização horizontal) -->
<div class="mx-auto">Centralizado</div>
```

**Equivalente CSS:**
```css
div {
  margin-left: auto;
  margin-right: auto;
}
```

```html
<!-- Margin negativo (sobreposição) -->
<div class="-mt-4">Sobreposto</div>
```

**Equivalente CSS:**
```css
div {
  margin-top: -1rem;
}
```

### Gap (Espaçamento em Flexbox e Grid)

**Sintaxe:** `gap-{valor}` para todos os lados, ou `gap-x-{valor}` e `gap-y-{valor}`

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

**Gap direcional:**

```html
<div class="grid gap-x-6 gap-y-4">
  <!-- Grid com gap horizontal maior que vertical -->
</div>
```

**Equivalente CSS:**
```css
div {
  display: grid;
  column-gap: 1.5rem;
  row-gap: 1rem;
}
```

---

## 🎨 Sistema de Cores

O Tailwind possui um **sistema de cores rico e consistente**, com cada cor tendo uma escala de 50 a 950, onde valores menores são mais claros e valores maiores são mais escuros.

### Estrutura de Cores

**Formato:** `{cor}-{intensidade}`

**Exemplo:** `blue-500` = azul na intensidade 500 (cor média)

### Escala de Intensidade

| Intensidade | Característica | Uso Comum |
|-------------|----------------|-----------|
| `50` | Muito claro | Backgrounds suaves, highlights |
| `100` | Claro | Backgrounds claros |
| `200` | Claro-médio | Bordas claras, backgrounds alternativos |
| `300` | Médio-claro | Hover states claros |
| `400` | Médio | Elementos secundários |
| `500` | Médio (cor base) | Cor principal da marca |
| `600` | Médio-escuro | Hover states, elementos ativos |
| `700` | Escuro | Texto, elementos importantes |
| `800` | Muito escuro | Texto escuro, elementos críticos |
| `900` | Extremamente escuro | Texto principal, máxima importância |
| `950` | Quase preto | Texto em máxima contraste |

### Cores Padrão do Tailwind

O Tailwind inclui as seguintes cores por padrão:

- **Slate** (cinza azulado)
- **Gray** (cinza neutro)
- **Zinc** (cinza metálico)
- **Neutral** (cinza neutro)
- **Stone** (cinza quente)
- **Red** (vermelho)
- **Orange** (laranja)
- **Amber** (âmbar)
- **Yellow** (amarelo)
- **Lime** (lima)
- **Green** (verde)
- **Emerald** (esmeralda)
- **Teal** (verde-azulado)
- **Cyan** (ciano)
- **Sky** (céu)
- **Blue** (azul)
- **Indigo** (anil)
- **Violet** (violeta)
- **Purple** (roxo)
- **Fuchsia** (fúcsia)
- **Pink** (rosa)
- **Rose** (rosa avermelhado)

### Aplicando Cores

#### Text Color (Cor do Texto)

**Sintaxe:** `text-{cor}-{intensidade}`

```html
<p class="text-blue-500">Texto azul médio</p>
<p class="text-gray-700">Texto cinza escuro</p>
<p class="text-red-600">Texto vermelho escuro</p>
```

**Equivalente CSS:**
```css
p {
  color: rgb(59 130 246); /* blue-500 */
}

p {
  color: rgb(55 65 81); /* gray-700 */
}

p {
  color: rgb(220 38 38); /* red-600 */
}
```

#### Background Color (Cor de Fundo)

**Sintaxe:** `bg-{cor}-{intensidade}`

```html
<div class="bg-blue-500">Fundo azul</div>
<div class="bg-gray-100">Fundo cinza claro</div>
```

**Equivalente CSS:**
```css
div {
  background-color: rgb(59 130 246); /* blue-500 */
}

div {
  background-color: rgb(243 244 246); /* gray-100 */
}
```

#### Border Color (Cor da Borda)

**Sintaxe:** `border-{cor}-{intensidade}`

```html
<div class="border-2 border-blue-500">Borda azul</div>
```

**Equivalente CSS:**
```css
div {
  border-width: 2px;
  border-color: rgb(59 130 246);
}
```

### Cores Especiais

#### Transparente

```html
<div class="bg-transparent">Fundo transparente</div>
```

**Equivalente CSS:**
```css
div {
  background-color: transparent;
}
```

#### Cor Atual (Current Color)

```html
<div class="text-blue-500 border-2 border-current">
  Borda usa a cor do texto
</div>
```

**Equivalente CSS:**
```css
div {
  color: rgb(59 130 246);
  border-width: 2px;
  border-color: currentColor;
}
```

---

## 🖼️ Backgrounds

Além de cores sólidas, o Tailwind oferece utilitários para backgrounds mais complexos.

### Background Images

**Sintaxe:** `bg-{tipo}`

#### Gradientes Lineares

```html
<div class="bg-gradient-to-r from-blue-500 to-purple-500">
  Gradiente horizontal
</div>
```

**Equivalente CSS:**
```css
div {
  background-image: linear-gradient(to right, rgb(59 130 246), rgb(168 85 247));
}
```

**Direções disponíveis:**
- `bg-gradient-to-t` → `linear-gradient(to top, ...)`
- `bg-gradient-to-tr` → `linear-gradient(to top right, ...)`
- `bg-gradient-to-r` → `linear-gradient(to right, ...)`
- `bg-gradient-to-br` → `linear-gradient(to bottom right, ...)`
- `bg-gradient-to-b` → `linear-gradient(to bottom, ...)`
- `bg-gradient-to-bl` → `linear-gradient(to bottom left, ...)`
- `bg-gradient-to-l` → `linear-gradient(to left, ...)`
- `bg-gradient-to-tl` → `linear-gradient(to top left, ...)`

**Gradiente com múltiplas cores:**

```html
<div class="bg-gradient-to-r from-blue-500 via-purple-500 to-pink-500">
  Gradiente com 3 cores
</div>
```

**Equivalente CSS:**
```css
div {
  background-image: linear-gradient(to right, 
    rgb(59 130 246), 
    rgb(168 85 247), 
    rgb(236 72 153)
  );
}
```

#### Background Size

```html
<div class="bg-cover">Cobre todo o elemento</div>
<div class="bg-contain">Contém a imagem</div>
<div class="bg-auto">Tamanho automático</div>
```

**Equivalente CSS:**
```css
div {
  background-size: cover;
}

div {
  background-size: contain;
}

div {
  background-size: auto;
}
```

#### Background Position

```html
<div class="bg-center">Centralizado</div>
<div class="bg-top">Topo</div>
<div class="bg-right">Direita</div>
<div class="bg-bottom">Fundo</div>
<div class="bg-left">Esquerda</div>
```

#### Background Repeat

```html
<div class="bg-no-repeat">Sem repetição</div>
<div class="bg-repeat">Repetir</div>
<div class="bg-repeat-x">Repetir horizontalmente</div>
<div class="bg-repeat-y">Repetir verticalmente</div>
```

---

## ✍️ Tipografia

O Tailwind oferece utilitários completos para tipografia, mapeando diretamente para propriedades CSS de texto.

### Font Size (Tamanho da Fonte)

**Sintaxe:** `text-{tamanho}`

| Classe Tailwind | Font Size (rem) | Font Size (px) | Line Height | Equivalente CSS |
|-----------------|-----------------|----------------|-------------|-----------------|
| `text-xs` | `0.75rem` | `12px` | `1rem` | `font-size: 0.75rem; line-height: 1rem;` |
| `text-sm` | `0.875rem` | `14px` | `1.25rem` | `font-size: 0.875rem; line-height: 1.25rem;` |
| `text-base` | `1rem` | `16px` | `1.5rem` | `font-size: 1rem; line-height: 1.5rem;` |
| `text-lg` | `1.125rem` | `18px` | `1.75rem` | `font-size: 1.125rem; line-height: 1.75rem;` |
| `text-xl` | `1.25rem` | `20px` | `1.75rem` | `font-size: 1.25rem; line-height: 1.75rem;` |
| `text-2xl` | `1.5rem` | `24px` | `2rem` | `font-size: 1.5rem; line-height: 2rem;` |
| `text-3xl` | `1.875rem` | `30px` | `2.25rem` | `font-size: 1.875rem; line-height: 2.25rem;` |
| `text-4xl` | `2.25rem` | `36px` | `2.5rem` | `font-size: 2.25rem; line-height: 2.5rem;` |
| `text-5xl` | `3rem` | `48px` | `3rem` | `font-size: 3rem; line-height: 3rem;` |
| `text-6xl` | `3.75rem` | `60px` | `3.75rem` | `font-size: 3.75rem; line-height: 3.75rem;` |
| `text-7xl` | `4.5rem` | `72px` | `4.5rem` | `font-size: 4.5rem; line-height: 4.5rem;` |
| `text-8xl` | `6rem` | `96px` | `6rem` | `font-size: 6rem; line-height: 6rem;` |
| `text-9xl` | `8rem` | `128px` | `8rem` | `font-size: 8rem; line-height: 8rem;` |

**Exemplo:**

```html
<h1 class="text-4xl">Título Grande</h1>
<p class="text-base">Texto normal</p>
```

**Equivalente CSS:**
```css
h1 {
  font-size: 2.25rem;
  line-height: 2.5rem;
}

p {
  font-size: 1rem;
  line-height: 1.5rem;
}
```

### Font Weight (Peso da Fonte)

**Sintaxe:** `font-{peso}`

| Classe Tailwind | Valor CSS | Descrição |
|-----------------|-----------|-----------|
| `font-thin` | `100` | Muito fino |
| `font-extralight` | `200` | Extra leve |
| `font-light` | `300` | Leve |
| `font-normal` | `400` | Normal (padrão) |
| `font-medium` | `500` | Médio |
| `font-semibold` | `600` | Semi-negrito |
| `font-bold` | `700` | Negrito |
| `font-extrabold` | `800` | Extra negrito |
| `font-black` | `900` | Muito negrito |

**Exemplo:**

```html
<p class="font-normal">Texto normal</p>
<p class="font-bold">Texto negrito</p>
```

**Equivalente CSS:**
```css
p {
  font-weight: 400;
}

p {
  font-weight: 700;
}
```

### Font Style (Estilo da Fonte)

```html
<p class="italic">Texto itálico</p>
<p class="not-italic">Texto não-itálico</p>
```

**Equivalente CSS:**
```css
p {
  font-style: italic;
}

p {
  font-style: normal;
}
```

### Line Height (Altura da Linha)

**Sintaxe:** `leading-{valor}`

| Classe Tailwind | Valor CSS | Descrição |
|-----------------|-----------|-----------|
| `leading-none` | `1` | Sem espaçamento extra |
| `leading-tight` | `1.25` | Apertado |
| `leading-snug` | `1.375` | Confortável |
| `leading-normal` | `1.5` | Normal |
| `leading-relaxed` | `1.625` | Relaxado |
| `leading-loose` | `2` | Solto |

**Exemplo:**

```html
<p class="leading-tight">Texto com line-height apertado</p>
<p class="leading-relaxed">Texto com line-height relaxado</p>
```

**Equivalente CSS:**
```css
p {
  line-height: 1.25;
}

p {
  line-height: 1.625;
}
```

### Letter Spacing (Espaçamento entre Letras)

**Sintaxe:** `tracking-{valor}`

| Classe Tailwind | Valor CSS | Descrição |
|-----------------|-----------|-----------|
| `tracking-tighter` | `-0.05em` | Mais apertado |
| `tracking-tight` | `-0.025em` | Apertado |
| `tracking-normal` | `0em` | Normal |
| `tracking-wide` | `0.025em` | Amplo |
| `tracking-wider` | `0.05em` | Mais amplo |
| `tracking-widest` | `0.1em` | Mais amplo ainda |

**Exemplo:**

```html
<h1 class="tracking-wide">Título com espaçamento amplo</h1>
```

**Equivalente CSS:**
```css
h1 {
  letter-spacing: 0.025em;
}
```

### Text Alignment (Alinhamento de Texto)

```html
<p class="text-left">Alinhado à esquerda</p>
<p class="text-center">Centralizado</p>
<p class="text-right">Alinhado à direita</p>
<p class="text-justify">Justificado</p>
```

**Equivalente CSS:**
```css
p {
  text-align: left;
}

p {
  text-align: center;
}

p {
  text-align: right;
}

p {
  text-align: justify;
}
```

### Text Decoration (Decoração de Texto)

```html
<p class="underline">Sublinhado</p>
<p class="overline">Sobrelinha</p>
<p class="line-through">Riscado</p>
<p class="no-underline">Sem decoração</p>
```

**Equivalente CSS:**
```css
p {
  text-decoration-line: underline;
}

p {
  text-decoration-line: overline;
}

p {
  text-decoration-line: line-through;
}

p {
  text-decoration-line: none;
}
```

### Text Transform (Transformação de Texto)

```html
<p class="uppercase">TEXTO EM MAIÚSCULAS</p>
<p class="lowercase">texto em minúsculas</p>
<p class="capitalize">Primeira Letra De Cada Palavra</p>
<p class="normal-case">Texto Normal</p>
```

**Equivalente CSS:**
```css
p {
  text-transform: uppercase;
}

p {
  text-transform: lowercase;
}

p {
  text-transform: capitalize;
}

p {
  text-transform: none;
}
```

### Text Overflow (Truncamento de Texto)

```html
<p class="truncate">Texto muito longo que será truncado...</p>
```

**Equivalente CSS:**
```css
p {
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}
```

---

## 🔲 Bordas

O Tailwind oferece utilitários completos para trabalhar com bordas.

### Border Width (Largura da Borda)

**Sintaxe:** `border` ou `border-{direção}-{valor}`

| Classe Tailwind | Propriedade CSS | Descrição |
|-----------------|-----------------|-----------|
| `border` | `border-width: 1px` | Borda de 1px em todos os lados |
| `border-0` | `border-width: 0px` | Sem borda |
| `border-2` | `border-width: 2px` | Borda de 2px |
| `border-4` | `border-width: 4px` | Borda de 4px |
| `border-8` | `border-width: 8px` | Borda de 8px |
| `border-t` | `border-top-width: 1px` | Borda apenas no topo |
| `border-r` | `border-right-width: 1px` | Borda apenas à direita |
| `border-b` | `border-bottom-width: 1px` | Borda apenas no fundo |
| `border-l` | `border-left-width: 1px` | Borda apenas à esquerda |

**Exemplo:**

```html
<div class="border-2 border-blue-500">Borda azul de 2px</div>
```

**Equivalente CSS:**
```css
div {
  border-width: 2px;
  border-color: rgb(59 130 246);
}
```

### Border Style (Estilo da Borda)

```html
<div class="border-solid">Borda sólida</div>
<div class="border-dashed">Borda tracejada</div>
<div class="border-dotted">Borda pontilhada</div>
<div class="border-double">Borda dupla</div>
<div class="border-none">Sem borda</div>
```

**Equivalente CSS:**
```css
div {
  border-style: solid;
}

div {
  border-style: dashed;
}

div {
  border-style: dotted;
}

div {
  border-style: double;
}

div {
  border-style: none;
}
```

### Border Radius (Arredondamento)

**Sintaxe:** `rounded-{valor}` ou direções específicas

| Classe Tailwind | Valor CSS | Descrição |
|-----------------|-----------|-----------|
| `rounded-none` | `0` | Sem arredondamento |
| `rounded-sm` | `0.125rem` | Pequeno (2px) |
| `rounded` | `0.25rem` | Padrão (4px) |
| `rounded-md` | `0.375rem` | Médio (6px) |
| `rounded-lg` | `0.5rem` | Grande (8px) |
| `rounded-xl` | `0.75rem` | Extra grande (12px) |
| `rounded-2xl` | `1rem` | 2x grande (16px) |
| `rounded-3xl` | `1.5rem` | 3x grande (24px) |
| `rounded-full` | `9999px` | Círculo completo |

**Arredondamento por canto:**

| Classe Tailwind | Propriedade CSS |
|-----------------|-----------------|
| `rounded-t-{valor}` | `border-top-left-radius` e `border-top-right-radius` |
| `rounded-r-{valor}` | `border-top-right-radius` e `border-bottom-right-radius` |
| `rounded-b-{valor}` | `border-bottom-right-radius` e `border-bottom-left-radius` |
| `rounded-l-{valor}` | `border-top-left-radius` e `border-bottom-left-radius` |
| `rounded-tl-{valor}` | `border-top-left-radius` |
| `rounded-tr-{valor}` | `border-top-right-radius` |
| `rounded-br-{valor}` | `border-bottom-right-radius` |
| `rounded-bl-{valor}` | `border-bottom-left-radius` |

**Exemplo:**

```html
<div class="rounded-lg">Cantos arredondados</div>
<div class="rounded-full w-16 h-16">Círculo</div>
```

**Equivalente CSS:**
```css
div {
  border-radius: 0.5rem;
}

div {
  border-radius: 9999px;
  width: 4rem;
  height: 4rem;
}
```

---

## 🌑 Sombras

O Tailwind oferece um sistema de sombras consistente.

### Box Shadow

**Sintaxe:** `shadow-{tamanho}`

| Classe Tailwind | Valor CSS | Descrição |
|-----------------|-----------|-----------|
| `shadow-sm` | `0 1px 2px 0 rgba(0, 0, 0, 0.05)` | Sombra pequena |
| `shadow` | `0 1px 3px 0 rgba(0, 0, 0, 0.1), 0 1px 2px -1px rgba(0, 0, 0, 0.1)` | Sombra padrão |
| `shadow-md` | `0 4px 6px -1px rgba(0, 0, 0, 0.1), 0 2px 4px -2px rgba(0, 0, 0, 0.1)` | Sombra média |
| `shadow-lg` | `0 10px 15px -3px rgba(0, 0, 0, 0.1), 0 4px 6px -4px rgba(0, 0, 0, 0.1)` | Sombra grande |
| `shadow-xl` | `0 20px 25px -5px rgba(0, 0, 0, 0.1), 0 8px 10px -6px rgba(0, 0, 0, 0.1)` | Sombra extra grande |
| `shadow-2xl` | `0 25px 50px -12px rgba(0, 0, 0, 0.25)` | Sombra 2x grande |
| `shadow-inner` | `inset 0 2px 4px 0 rgba(0, 0, 0, 0.05)` | Sombra interna |
| `shadow-none` | `none` | Sem sombra |

**Exemplo:**

```html
<div class="shadow-lg">Card com sombra grande</div>
```

**Equivalente CSS:**
```css
div {
  box-shadow: 0 10px 15px -3px rgba(0, 0, 0, 0.1), 
              0 4px 6px -4px rgba(0, 0, 0, 0.1);
}
```

### Shadow Color (Cor da Sombra)

**Sintaxe:** `shadow-{cor}-{intensidade}`

```html
<div class="shadow-lg shadow-blue-500/50">
  Sombra azul com opacidade
</div>
```

---

## 👻 Opacidade e Visibilidade

### Opacity (Opacidade)

**Sintaxe:** `opacity-{valor}`

| Classe Tailwind | Valor CSS | Descrição |
|-----------------|-----------|-----------|
| `opacity-0` | `0` | Totalmente transparente |
| `opacity-5` | `0.05` | 5% opaco |
| `opacity-10` | `0.1` | 10% opaco |
| `opacity-20` | `0.2` | 20% opaco |
| `opacity-25` | `0.25` | 25% opaco |
| `opacity-30` | `0.3` | 30% opaco |
| `opacity-40` | `0.4` | 40% opaco |
| `opacity-50` | `0.5` | 50% opaco (semi-transparente) |
| `opacity-60` | `0.6` | 60% opaco |
| `opacity-70` | `0.7` | 70% opaco |
| `opacity-75` | `0.75` | 75% opaco |
| `opacity-80` | `0.8` | 80% opaco |
| `opacity-90` | `0.9` | 90% opaco |
| `opacity-95` | `0.95` | 95% opaco |
| `opacity-100` | `1` | Totalmente opaco (padrão) |

**Exemplo:**

```html
<div class="opacity-50">Elemento semi-transparente</div>
```

**Equivalente CSS:**
```css
div {
  opacity: 0.5;
}
```

### Opacidade em Cores (Sintaxe Moderna)

Tailwind também permite definir opacidade diretamente nas cores:

```html
<div class="bg-blue-500/50">Fundo azul com 50% de opacidade</div>
<div class="text-red-600/75">Texto vermelho com 75% de opacidade</div>
```

**Equivalente CSS:**
```css
div {
  background-color: rgb(59 130 246 / 0.5);
}

div {
  color: rgb(220 38 38 / 0.75);
}
```

### Visibility (Visibilidade)

```html
<div class="visible">Visível</div>
<div class="invisible">Invisível (mas ocupa espaço)</div>
```

**Equivalente CSS:**
```css
div {
  visibility: visible;
}

div {
  visibility: hidden;
}
```

**Diferença importante:**
- `invisible` → elemento ocupa espaço no layout, mas não é visível
- `hidden` (display) → elemento não ocupa espaço no layout

---

## 🎯 Combinando Classes: Exemplos Práticos

Agora que você conhece os fundamentos, vamos ver como combinar classes para criar componentes:

### Exemplo 1: Card Simples

```html
<div class="p-6 bg-white rounded-lg shadow-md border border-gray-200">
  <h2 class="text-2xl font-bold text-gray-800 mb-2">Título do Card</h2>
  <p class="text-gray-600 leading-relaxed">Descrição do conteúdo do card.</p>
</div>
```

**Equivalente CSS:**
```css
div {
  padding: 1.5rem;
  background-color: white;
  border-radius: 0.5rem;
  box-shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.1), 
              0 2px 4px -2px rgba(0, 0, 0, 0.1);
  border-width: 1px;
  border-color: rgb(229 231 235);
}

h2 {
  font-size: 1.5rem;
  font-weight: 700;
  color: rgb(31 41 55);
  margin-bottom: 0.5rem;
}

p {
  color: rgb(75 85 99);
  line-height: 1.625;
}
```

### Exemplo 2: Botão Estilizado

```html
<button class="px-6 py-3 bg-blue-500 text-white font-semibold rounded-lg shadow-md hover:bg-blue-600 transition-colors">
  Clique Aqui
</button>
```

**Equivalente CSS:**
```css
button {
  padding-left: 1.5rem;
  padding-right: 1.5rem;
  padding-top: 0.75rem;
  padding-bottom: 0.75rem;
  background-color: rgb(59 130 246);
  color: white;
  font-weight: 600;
  border-radius: 0.5rem;
  box-shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.1), 
              0 2px 4px -2px rgba(0, 0, 0, 0.1);
  transition-property: color, background-color, border-color, text-decoration-color, fill, stroke;
  transition-timing-function: cubic-bezier(0.4, 0, 0.2, 1);
  transition-duration: 150ms;
}

button:hover {
  background-color: rgb(37 99 235);
}
```

### Exemplo 3: Badge/Etiqueta

```html
<span class="inline-block px-3 py-1 bg-green-100 text-green-800 text-sm font-medium rounded-full">
  Ativo
</span>
```

**Equivalente CSS:**
```css
span {
  display: inline-block;
  padding-left: 0.75rem;
  padding-right: 0.75rem;
  padding-top: 0.25rem;
  padding-bottom: 0.25rem;
  background-color: rgb(220 252 231);
  color: rgb(22 101 52);
  font-size: 0.875rem;
  font-weight: 500;
  border-radius: 9999px;
}
```

---

## 📝 Resumo dos Conceitos Principais

### Sistema de Espaçamento
- **Padding:** `p-{valor}`, `px-{valor}`, `py-{valor}`, `pt-{valor}`, etc.
- **Margin:** `m-{valor}`, `mx-{valor}`, `my-{valor}`, `mt-{valor}`, etc.
- **Gap:** `gap-{valor}`, `gap-x-{valor}`, `gap-y-{valor}`
- Escala baseada em múltiplos de `0.25rem` (4px)

### Sistema de Cores
- **Formato:** `{cor}-{intensidade}` (ex: `blue-500`)
- **Intensidades:** 50 (mais claro) a 950 (mais escuro)
- **Aplicação:** `text-{cor}`, `bg-{cor}`, `border-{cor}`
- **Opacidade:** `{cor}-{intensidade}/50` (50% de opacidade)

### Tipografia
- **Tamanho:** `text-xs` a `text-9xl`
- **Peso:** `font-thin` (100) a `font-black` (900)
- **Alinhamento:** `text-left`, `text-center`, `text-right`
- **Decoração:** `underline`, `line-through`, `no-underline`
- **Transformação:** `uppercase`, `lowercase`, `capitalize`

### Bordas e Arredondamento
- **Largura:** `border`, `border-2`, `border-4`
- **Cor:** `border-{cor}-{intensidade}`
- **Arredondamento:** `rounded-sm` a `rounded-full`
- **Direções:** `rounded-t`, `rounded-r`, `rounded-b`, `rounded-l`

### Sombras
- **Tamanhos:** `shadow-sm` a `shadow-2xl`
- **Especial:** `shadow-inner` (sombra interna)
- **Cor:** `shadow-{cor}-{intensidade}`

### Opacidade e Visibilidade
- **Opacidade:** `opacity-0` a `opacity-100`
- **Visibilidade:** `visible`, `invisible`
- **Opacidade em cores:** `bg-blue-500/50`

---

## 🎓 Próximos Passos

Agora que você domina os fundamentos das classes utilitárias, você está pronto para:
- Trabalhar com layouts (Flexbox e Grid) - **Aula 3**
- Criar designs responsivos - **Aula 5**
- Adicionar interatividade e estados - **Aula 6**

Continue praticando combinando essas classes para criar componentes visuais. Quanto mais você praticar, mais natural se tornará o mapeamento mental entre classes Tailwind e propriedades CSS.

