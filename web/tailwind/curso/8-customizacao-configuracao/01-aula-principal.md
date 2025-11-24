# Aula 8: Customização e Configuração do Tailwind - Conteúdo Principal

## 📖 Introdução

Até agora, você usou o Tailwind com suas configurações padrão. Mas o verdadeiro poder do Tailwind está na sua **capacidade de customização**. Você pode adaptar cores, espaçamentos, breakpoints, tipografia e muito mais para se adequar perfeitamente ao seu projeto.

Nesta aula, você aprenderá:
- Estrutura do arquivo `tailwind.config.js`
- Customização do tema (cores, espaçamento, tipografia)
- Customização de breakpoints
- Adicionar utilitários customizados
- Usar `theme.extend` vs `theme` (substituir)
- Preservar valores padrão enquanto adiciona novos

---

## ⚙️ O Arquivo tailwind.config.js

O arquivo `tailwind.config.js` é o **coração da customização** do Tailwind. É aqui que você define como o Tailwind deve gerar suas classes utilitárias.

### Estrutura Básica

```javascript
/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./src/**/*.{html,js}",
  ],
  theme: {
    extend: {},
  },
  plugins: [],
}
```

### Componentes Principais

1. **`content`**: Define onde o Tailwind deve procurar classes usadas (para purge/tree-shaking)
2. **`theme`**: Define o sistema de design (cores, espaçamento, etc.)
3. **`plugins`**: Adiciona funcionalidades extras

### Conexão com CSS

O `tailwind.config.js` funciona como um **gerador de variáveis CSS**. Quando você customiza o tema, o Tailwind gera classes baseadas nessas configurações, que se traduzem em CSS puro.

**Exemplo:**
```javascript
// tailwind.config.js
theme: {
  extend: {
    colors: {
      'minha-cor': '#ff6b6b',
    }
  }
}
```

**Resultado em CSS:**
```css
.bg-minha-cor {
  background-color: #ff6b6b;
}
```

---

## 🎨 Customizando Cores

O Tailwind vem com uma paleta de cores padrão (slate, gray, zinc, neutral, stone, red, orange, yellow, green, emerald, teal, cyan, sky, blue, indigo, violet, purple, fuchsia, pink, rose), cada uma com tons de 50 a 950.

### Adicionar Cores Personalizadas com `extend`

**Use `extend` quando quiser ADICIONAR cores sem perder as padrão:**

```javascript
/** @type {import('tailwindcss').Config} */
module.exports = {
  theme: {
    extend: {
      colors: {
        'brand': {
          50: '#f0f9ff',
          100: '#e0f2fe',
          200: '#bae6fd',
          300: '#7dd3fc',
          400: '#38bdf8',
          500: '#0ea5e9', // Cor principal
          600: '#0284c7',
          700: '#0369a1',
          800: '#075985',
          900: '#0c4a6e',
          950: '#082f49',
        },
        'accent': '#ff6b6b',
        'primary': '#4ecdc4',
      }
    }
  }
}
```

**Uso:**
```html
<div class="bg-brand-500 text-white p-4">
  Usando cor customizada
</div>
<div class="bg-accent text-white p-4">
  Cor simples customizada
</div>
```

**Equivalente CSS:**
```css
.bg-brand-500 {
  background-color: #0ea5e9;
}

.bg-accent {
  background-color: #ff6b6b;
}
```

### Substituir Cores Padrão com `theme` (sem extend)

**Use `theme` diretamente quando quiser SUBSTITUIR cores padrão:**

```javascript
/** @type {import('tailwindcss').Config} */
module.exports = {
  theme: {
    colors: {
      // Isso SUBSTITUI todas as cores padrão
      'blue': '#1e40af',
      'red': '#dc2626',
      'white': '#ffffff',
      'black': '#000000',
    }
  }
}
```

⚠️ **Atenção**: Quando você usa `theme` sem `extend`, você **substitui** completamente o sistema de cores padrão. As cores padrão (slate, gray, etc.) não estarão mais disponíveis.

### Cores com Opacidade

O Tailwind permite usar cores com opacidade usando a sintaxe `/opacidade`:

```html
<div class="bg-brand-500/50">50% de opacidade</div>
<div class="bg-brand-500/75">75% de opacidade</div>
```

**Equivalente CSS:**
```css
.bg-brand-500\/50 {
  background-color: rgb(14 165 233 / 0.5);
}
```

### Cores CSS Custom Properties

Você também pode usar variáveis CSS:

```javascript
theme: {
  extend: {
    colors: {
      'primary': 'var(--color-primary)',
      'secondary': 'var(--color-secondary)',
    }
  }
}
```

**CSS:**
```css
:root {
  --color-primary: #4ecdc4;
  --color-secondary: #ff6b6b;
}
```

---

## 📏 Customizando Espaçamento

O sistema de espaçamento do Tailwind é usado para `padding`, `margin`, `gap`, `width`, `height`, etc.

### Adicionar Valores de Espaçamento

```javascript
theme: {
  extend: {
    spacing: {
      '18': '4.5rem',   // 72px
      '88': '22rem',    // 352px
      '128': '32rem',   // 512px
      '1/2': '50%',
      '1/3': '33.333333%',
      '2/3': '66.666667%',
      '1/4': '25%',
      '3/4': '75%',
    }
  }
}
```

**Uso:**
```html
<div class="p-18">Padding de 4.5rem</div>
<div class="w-1/2">Largura de 50%</div>
<div class="gap-128">Gap de 32rem</div>
```

**Equivalente CSS:**
```css
.p-18 {
  padding: 4.5rem;
}

.w-1\/2 {
  width: 50%;
}

.gap-128 {
  gap: 32rem;
}
```

### Substituir Espaçamento Padrão

```javascript
theme: {
  spacing: {
    // Isso substitui TODA a escala padrão
    '0': '0',
    '1': '0.25rem',
    '2': '0.5rem',
    '4': '1rem',
    '8': '2rem',
  }
}
```

---

## 🔤 Customizando Tipografia

### Font Families (Famílias de Fonte)

```javascript
theme: {
  extend: {
    fontFamily: {
      'sans': ['Inter', 'system-ui', 'sans-serif'],
      'serif': ['Merriweather', 'serif'],
      'mono': ['Fira Code', 'monospace'],
      'display': ['Poppins', 'sans-serif'],
    }
  }
}
```

**Uso:**
```html
<p class="font-sans">Fonte sans customizada</p>
<p class="font-display">Fonte display</p>
```

**Equivalente CSS:**
```css
.font-sans {
  font-family: Inter, system-ui, sans-serif;
}

.font-display {
  font-family: Poppins, sans-serif;
}
```

### Font Sizes (Tamanhos de Fonte)

O Tailwind permite customizar tamanhos de fonte com `fontSize`, que pode incluir `line-height` e `letter-spacing`:

```javascript
theme: {
  extend: {
    fontSize: {
      'xs': ['0.75rem', { lineHeight: '1rem' }],
      'sm': ['0.875rem', { lineHeight: '1.25rem' }],
      'base': ['1rem', { lineHeight: '1.5rem' }],
      'lg': ['1.125rem', { lineHeight: '1.75rem' }],
      'xl': ['1.25rem', { lineHeight: '1.75rem' }],
      '2xl': ['1.5rem', { lineHeight: '2rem' }],
      '3xl': ['1.875rem', { lineHeight: '2.25rem' }],
      '4xl': ['2.25rem', { lineHeight: '2.5rem' }],
      '5xl': ['3rem', { lineHeight: '1' }],
      '6xl': ['3.75rem', { lineHeight: '1' }],
      '7xl': ['4.5rem', { lineHeight: '1' }],
      '8xl': ['6rem', { lineHeight: '1' }],
      '9xl': ['8rem', { lineHeight: '1' }],
      // Tamanho customizado
      'hero': ['4rem', { lineHeight: '1.1', letterSpacing: '-0.02em' }],
    }
  }
}
```

**Uso:**
```html
<h1 class="text-hero font-bold">Título Hero</h1>
```

**Equivalente CSS:**
```css
.text-hero {
  font-size: 4rem;
  line-height: 1.1;
  letter-spacing: -0.02em;
}
```

### Font Weight (Peso da Fonte)

```javascript
theme: {
  extend: {
    fontWeight: {
      'thin': '100',
      'extralight': '200',
      'light': '300',
      'normal': '400',
      'medium': '500',
      'semibold': '600',
      'bold': '700',
      'extrabold': '800',
      'black': '900',
      // Customizado
      'custom': '350',
    }
  }
}
```

---

## 📱 Customizando Breakpoints

Os breakpoints padrão do Tailwind são:
- `sm`: 640px
- `md`: 768px
- `lg`: 1024px
- `xl`: 1280px
- `2xl`: 1536px

### Adicionar Breakpoints Customizados

```javascript
theme: {
  extend: {
    screens: {
      'xs': '475px',
      '3xl': '1920px',
      '4xl': '2560px',
      // Breakpoint customizado com min e max
      'tablet': {'min': '640px', 'max': '1023px'},
      // Breakpoint apenas com max
      'mobile': {'max': '639px'},
    }
  }
}
```

**Uso:**
```html
<div class="text-sm mobile:text-base tablet:text-lg desktop:text-xl">
  Texto responsivo
</div>
```

**Equivalente CSS:**
```css
@media (max-width: 639px) {
  .mobile\:text-base {
    font-size: 1rem;
  }
}

@media (min-width: 640px) and (max-width: 1023px) {
  .tablet\:text-lg {
    font-size: 1.125rem;
  }
}
```

### Substituir Breakpoints Padrão

```javascript
theme: {
  screens: {
    // Isso substitui TODOS os breakpoints padrão
    'sm': '600px',
    'md': '900px',
    'lg': '1200px',
  }
}
```

---

## 🎭 Customizando Bordas

### Border Radius (Raio da Borda)

```javascript
theme: {
  extend: {
    borderRadius: {
      'none': '0',
      'sm': '0.125rem',
      'DEFAULT': '0.25rem',
      'md': '0.375rem',
      'lg': '0.5rem',
      'xl': '0.75rem',
      '2xl': '1rem',
      '3xl': '1.5rem',
      'full': '9999px',
      // Customizado
      'extra': '2rem',
    }
  }
}
```

### Border Width (Largura da Borda)

```javascript
theme: {
  extend: {
    borderWidth: {
      DEFAULT: '1px',
      '0': '0',
      '2': '2px',
      '4': '4px',
      '8': '8px',
      // Customizado
      '3': '3px',
      '5': '5px',
    }
  }
}
```

---

## 🌑 Customizando Sombras

```javascript
theme: {
  extend: {
    boxShadow: {
      'sm': '0 1px 2px 0 rgb(0 0 0 / 0.05)',
      'DEFAULT': '0 1px 3px 0 rgb(0 0 0 / 0.1), 0 1px 2px -1px rgb(0 0 0 / 0.1)',
      'md': '0 4px 6px -1px rgb(0 0 0 / 0.1), 0 2px 4px -2px rgb(0 0 0 / 0.1)',
      'lg': '0 10px 15px -3px rgb(0 0 0 / 0.1), 0 4px 6px -4px rgb(0 0 0 / 0.1)',
      'xl': '0 20px 25px -5px rgb(0 0 0 / 0.1), 0 8px 10px -6px rgb(0 0 0 / 0.1)',
      '2xl': '0 25px 50px -12px rgb(0 0 0 / 0.25)',
      'inner': 'inset 0 2px 4px 0 rgb(0 0 0 / 0.05)',
      'none': 'none',
      // Sombra customizada
      'glow': '0 0 20px rgba(59, 130, 246, 0.5)',
      'glow-lg': '0 0 40px rgba(59, 130, 246, 0.6)',
    }
  }
}
```

**Uso:**
```html
<div class="shadow-glow">Elemento com brilho</div>
```

---

## 🎯 Adicionando Utilitários Customizados

Você pode adicionar utilitários completamente novos usando plugins ou diretamente no config.

### Método 1: Usando `addUtilities` em um Plugin

```javascript
// tailwind.config.js
const plugin = require('tailwindcss/plugin')

module.exports = {
  plugins: [
    plugin(function({ addUtilities }) {
      addUtilities({
        '.scrollbar-hide': {
          /* IE and Edge */
          '-ms-overflow-style': 'none',
          /* Firefox */
          'scrollbar-width': 'none',
          /* Safari and Chrome */
          '&::-webkit-scrollbar': {
            display: 'none'
          }
        },
        '.text-balance': {
          'text-wrap': 'balance',
        },
      })
    })
  ]
}
```

**Uso:**
```html
<div class="scrollbar-hide overflow-auto">
  Conteúdo com scrollbar escondida
</div>
```

### Método 2: Usando CSS Customizado com @apply

Você também pode criar utilitários usando CSS:

```css
/* styles.css */
@tailwind base;
@tailwind components;
@tailwind utilities;

@layer utilities {
  .scrollbar-hide {
    -ms-overflow-style: none;
    scrollbar-width: none;
  }
  
  .scrollbar-hide::-webkit-scrollbar {
    display: none;
  }
}
```

---

## 🔄 theme.extend vs theme (Substituir)

### `theme.extend` - Adicionar sem Perder Padrões

```javascript
theme: {
  extend: {
    colors: {
      'brand': '#ff6b6b',
    }
  }
}
```

**Resultado:**
- ✅ Mantém todas as cores padrão (blue, red, green, etc.)
- ✅ Adiciona `brand` às cores disponíveis
- ✅ Você pode usar `bg-blue-500` E `bg-brand`

### `theme` (sem extend) - Substituir Completamente

```javascript
theme: {
  colors: {
    'brand': '#ff6b6b',
  }
}
```

**Resultado:**
- ❌ Remove TODAS as cores padrão
- ✅ Apenas `brand` estará disponível
- ❌ `bg-blue-500` não funcionará mais

**Quando usar cada um:**

- **Use `extend`**: Quando quiser adicionar valores mantendo os padrão (99% dos casos)
- **Use `theme` direto**: Quando quiser criar um sistema de design completamente novo do zero

---

## 📐 Exemplo Completo de Configuração

Aqui está um exemplo completo de `tailwind.config.js` com várias customizações:

```javascript
/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./src/**/*.{html,js,jsx,ts,tsx}",
    "./public/index.html",
  ],
  theme: {
    extend: {
      // Cores personalizadas
      colors: {
        'brand': {
          50: '#f0f9ff',
          100: '#e0f2fe',
          500: '#0ea5e9',
          900: '#0c4a6e',
        },
        'accent': '#ff6b6b',
      },
      // Espaçamento customizado
      spacing: {
        '18': '4.5rem',
        '88': '22rem',
      },
      // Tipografia
      fontFamily: {
        'sans': ['Inter', 'system-ui', 'sans-serif'],
        'display': ['Poppins', 'sans-serif'],
      },
      fontSize: {
        'hero': ['4rem', { lineHeight: '1.1', letterSpacing: '-0.02em' }],
      },
      // Breakpoints
      screens: {
        'xs': '475px',
        '3xl': '1920px',
      },
      // Bordas
      borderRadius: {
        'extra': '2rem',
      },
      // Sombras
      boxShadow: {
        'glow': '0 0 20px rgba(59, 130, 246, 0.5)',
      },
    },
  },
  plugins: [],
}
```

---

## 🎨 Usando Variáveis CSS no Tema

Você pode combinar Tailwind com CSS Custom Properties para máxima flexibilidade:

```javascript
// tailwind.config.js
module.exports = {
  theme: {
    extend: {
      colors: {
        'primary': 'var(--color-primary)',
        'secondary': 'var(--color-secondary)',
      }
    }
  }
}
```

```css
/* styles.css */
:root {
  --color-primary: #4ecdc4;
  --color-secondary: #ff6b6b;
}

[data-theme="dark"] {
  --color-primary: #6ee7d8;
  --color-secondary: #ff8787;
}
```

Isso permite **temas dinâmicos** que mudam em runtime!

---

## 🔍 Acessando Valores do Tema em CSS Customizado

Você pode acessar valores do tema usando a função `theme()`:

```css
@layer components {
  .custom-component {
    padding: theme('spacing.4');
    background-color: theme('colors.brand.500');
    border-radius: theme('borderRadius.lg');
  }
}
```

**Equivalente:**
```css
.custom-component {
  padding: 1rem;
  background-color: #0ea5e9;
  border-radius: 0.5rem;
}
```

---

## 📝 Resumo dos Conceitos Principais

### 1. Estrutura do Config
- `content`: Onde o Tailwind procura classes
- `theme`: Sistema de design
- `plugins`: Funcionalidades extras

### 2. Customização de Cores
- Use `extend` para adicionar sem perder padrões
- Use `theme` direto para substituir completamente
- Suporte a opacidade com `/valor`

### 3. Customização de Espaçamento
- Valores numéricos ou percentuais
- Usado em padding, margin, gap, width, height

### 4. Customização de Tipografia
- `fontFamily`: Famílias de fonte
- `fontSize`: Tamanhos com line-height e letter-spacing
- `fontWeight`: Pesos da fonte

### 5. Customização de Breakpoints
- Adicionar novos breakpoints
- Criar breakpoints com min/max
- Substituir breakpoints padrão

### 6. Utilitários Customizados
- Usar plugins com `addUtilities`
- Usar CSS com `@layer utilities`

---

## 🎯 Próximos Passos

Agora que você entende como customizar o Tailwind, você pode:
- Adaptar o Tailwind ao design system do seu projeto
- Criar utilitários específicos para suas necessidades
- Combinar Tailwind com CSS Custom Properties para temas dinâmicos

Na próxima aula, você aprenderá sobre **Plugins e Extensões do Tailwind**, expandindo ainda mais as capacidades do framework!

---

**Lembre-se**: Customização é poderosa, mas use com moderação. Muitas customizações podem tornar o projeto difícil de manter. Customize quando realmente necessário para seu design system!

