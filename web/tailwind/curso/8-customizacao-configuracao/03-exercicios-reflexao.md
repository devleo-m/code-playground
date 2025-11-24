# Aula 8 - Exercícios e Reflexão: Customização e Configuração do Tailwind

## 🎯 Objetivos dos Exercícios

Ao completar estes exercícios, você será capaz de:
- Configurar o arquivo `tailwind.config.js` corretamente
- Customizar cores, espaçamento e tipografia do tema
- Adicionar breakpoints customizados
- Criar utilitários customizados
- Decidir quando usar `extend` vs substituir completamente
- Avaliar o impacto das customizações na manutenibilidade
- Pensar criticamente sobre quando customizar vs usar padrões

---

## 📝 Exercício 1: Configurando o Tema para um Projeto de E-commerce

### Tarefa

Você está criando um projeto de e-commerce e precisa configurar o Tailwind com as cores da marca, espaçamentos e tipografia específicos.

**Cores da marca:**
- Primária: `#2563eb` (azul)
- Secundária: `#f59e0b` (laranja)
- Sucesso: `#10b981` (verde)
- Erro: `#ef4444` (vermelho)
- Neutro: `#6b7280` (cinza)

**Requisitos:**
1. Adicione as cores da marca mantendo as cores padrão do Tailwind
2. Crie uma escala completa (50-950) para a cor primária
3. Adicione um espaçamento customizado `15` equivalente a `3.75rem`
4. Configure a fonte `Inter` como fonte sans padrão
5. Adicione um breakpoint `xs` em `475px`

### Código Base

```javascript
/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./src/**/*.{html,js}",
  ],
  theme: {
    // Seu código aqui
  },
  plugins: [],
}
```

### Teste

Após configurar, teste se funciona:

```html
<div class="bg-primary-500 text-white p-15">
  Teste de cor primária e espaçamento customizado
</div>
<div class="font-sans text-lg">
  Teste de fonte Inter
</div>
<div class="xs:text-xl">
  Teste de breakpoint xs
</div>
```

### Solução Esperada

<details>
<summary>Clique para ver a solução</summary>

```javascript
/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./src/**/*.{html,js}",
  ],
  theme: {
    extend: {
      colors: {
        'primary': {
          50: '#eff6ff',
          100: '#dbeafe',
          200: '#bfdbfe',
          300: '#93c5fd',
          400: '#60a5fa',
          500: '#2563eb', // Cor principal
          600: '#1d4ed8',
          700: '#1e40af',
          800: '#1e3a8a',
          900: '#1e3a8a',
          950: '#172554',
        },
        'secondary': '#f59e0b',
        'success': '#10b981',
        'error': '#ef4444',
        'neutral': '#6b7280',
      },
      spacing: {
        '15': '3.75rem',
      },
      fontFamily: {
        'sans': ['Inter', 'system-ui', 'sans-serif'],
      },
      screens: {
        'xs': '475px',
      },
    },
  },
  plugins: [],
}
```

</details>

---

## 📝 Exercício 2: Criando Utilitários Customizados

### Tarefa

Você precisa criar utilitários customizados para funcionalidades específicas do seu projeto:

1. **`.scrollbar-hide`**: Esconde a scrollbar mas mantém a funcionalidade de scroll
2. **`.text-balance`**: Aplica `text-wrap: balance` (útil para títulos)
3. **`.aspect-video-custom`**: Cria um aspect ratio de 21:9 (ultra-wide)

### Requisitos

- Use a função `addUtilities` em um plugin
- Todos os utilitários devem funcionar em diferentes navegadores
- Documente cada utilitário com comentários

### Código Base

```javascript
const plugin = require('tailwindcss/plugin')

module.exports = {
  plugins: [
    plugin(function({ addUtilities }) {
      // Seu código aqui
    })
  ]
}
```

### Teste

```html
<div class="scrollbar-hide overflow-auto h-64">
  Conteúdo com scroll escondido
</div>

<h1 class="text-balance text-2xl">
  Título que se ajusta automaticamente
</h1>

<div class="aspect-video-custom bg-gray-200">
  Container 21:9
</div>
```

### Solução Esperada

<details>
<summary>Clique para ver a solução</summary>

```javascript
const plugin = require('tailwindcss/plugin')

module.exports = {
  plugins: [
    plugin(function({ addUtilities }) {
      addUtilities({
        // Esconde scrollbar mas mantém funcionalidade
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
        // Balanceia texto (útil para títulos)
        '.text-balance': {
          'text-wrap': 'balance',
        },
        // Aspect ratio 21:9 (ultra-wide)
        '.aspect-video-custom': {
          'aspect-ratio': '21 / 9',
        },
      })
    })
  ]
}
```

</details>

---

## 📝 Exercício 3: Customização Avançada - Sistema de Design Completo

### Tarefa

Crie um sistema de design completo customizando múltiplos aspectos do Tailwind:

**Requisitos:**
1. Cores: Adicione uma paleta completa de cores da marca (escala 50-950)
2. Tipografia: 
   - Fonte sans: `'Poppins', sans-serif`
   - Fonte display: `'Montserrat', sans-serif`
   - Tamanho customizado `hero`: `5rem` com line-height `1.1`
3. Espaçamento: Adicione valores `13`, `17`, `21` (baseados na escala)
4. Breakpoints: Adicione `xs: 475px` e `3xl: 1920px`
5. Sombras: Adicione uma sombra customizada `glow-blue` com brilho azul
6. Bordas: Adicione `rounded-extra: 2rem`

### Código Base

```javascript
/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./src/**/*.{html,js}"],
  theme: {
    extend: {
      // Seu código aqui
    },
  },
  plugins: [],
}
```

### Solução Esperada

<details>
<summary>Clique para ver a solução</summary>

```javascript
/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./src/**/*.{html,js}"],
  theme: {
    extend: {
      colors: {
        'brand': {
          50: '#f0f9ff',
          100: '#e0f2fe',
          200: '#bae6fd',
          300: '#7dd3fc',
          400: '#38bdf8',
          500: '#0ea5e9',
          600: '#0284c7',
          700: '#0369a1',
          800: '#075985',
          900: '#0c4a6e',
          950: '#082f49',
        },
      },
      fontFamily: {
        'sans': ['Poppins', 'sans-serif'],
        'display': ['Montserrat', 'sans-serif'],
      },
      fontSize: {
        'hero': ['5rem', { lineHeight: '1.1' }],
      },
      spacing: {
        '13': '3.25rem',
        '17': '4.25rem',
        '21': '5.25rem',
      },
      screens: {
        'xs': '475px',
        '3xl': '1920px',
      },
      boxShadow: {
        'glow-blue': '0 0 20px rgba(59, 130, 246, 0.5)',
      },
      borderRadius: {
        'extra': '2rem',
      },
    },
  },
  plugins: [],
}
```

</details>

---

## 📝 Exercício 4: Análise de Código - Identificando Problemas

### Tarefa

Analise o seguinte `tailwind.config.js` e identifique os problemas:

```javascript
/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./src/**/*.{html,js}"],
  theme: {
    colors: {
      'primary': '#2563eb',
      'secondary': '#f59e0b',
    },
    spacing: {
      '4': '1rem',
      '8': '2rem',
    },
    screens: {
      'md': '768px',
      'lg': '1024px',
    },
  },
  plugins: [],
}
```

### Perguntas

1. Qual é o problema principal desta configuração?
2. O que acontecerá quando você tentar usar `bg-blue-500`?
3. O que acontecerá quando você tentar usar `p-2`?
4. Como você corrigiria esta configuração?

### Solução Esperada

<details>
<summary>Clique para ver a solução</summary>

**Problemas identificados:**

1. **Uso de `theme` direto sem `extend`**: Isso **substitui** completamente os valores padrão do Tailwind, removendo todas as cores, espaçamentos e breakpoints padrão.

2. **Cores padrão perdidas**: `bg-blue-500`, `bg-red-500`, `text-gray-500`, etc. não funcionarão mais.

3. **Espaçamento limitado**: Apenas `p-4` e `p-8` funcionarão. `p-2`, `p-6`, `p-12`, etc. não existirão mais.

4. **Breakpoints limitados**: Apenas `md` e `lg` funcionarão. `sm`, `xl`, `2xl` não existirão mais.

**Correção:**

```javascript
/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./src/**/*.{html,js}"],
  theme: {
    extend: { // ✅ Usar extend para adicionar sem perder padrões
      colors: {
        'primary': '#2563eb',
        'secondary': '#f59e0b',
      },
      // Não precisa redefinir spacing e screens se usar extend
    },
  },
  plugins: [],
}
```

</details>

---

## 🤔 Perguntas de Reflexão

### Reflexão 1: Quando Customizar vs Usar Padrões

**Pergunta:** Em um projeto novo, você deve começar customizando o Tailwind imediatamente ou usar os padrões primeiro? Justifique sua resposta.

**Pontos para considerar:**
- Velocidade de desenvolvimento inicial
- Consistência do design system
- Manutenibilidade a longo prazo
- Facilidade de onboarding de novos desenvolvedores
- Tamanho do bundle CSS

---

### Reflexão 2: Impacto de Customizações Excessivas

**Pergunta:** Quais são os riscos de customizar demais o Tailwind? Quando as customizações se tornam um problema?

**Pontos para considerar:**
- Complexidade de manutenção
- Curva de aprendizado para novos desenvolvedores
- Tamanho do bundle CSS gerado
- Compatibilidade com plugins do Tailwind
- Facilidade de atualizar o Tailwind
- Consistência visual do projeto

---

### Reflexão 3: theme.extend vs theme Direto

**Pergunta:** Em que cenários específicos faria sentido usar `theme` diretamente (sem `extend`) ao invés de `extend`? Dê exemplos práticos.

**Pontos para considerar:**
- Projetos com design system muito específico
- Projetos que precisam de controle total
- Migração de outros frameworks
- Projetos com requisitos de acessibilidade específicos
- Trade-offs entre flexibilidade e controle

---

### Reflexão 4: Performance e Bundle Size

**Pergunta:** Como as customizações do Tailwind afetam o tamanho do bundle CSS final? Quais estratégias você usaria para minimizar o impacto?

**Pontos para considerar:**
- PurgeCSS e tree-shaking
- Quantidade de valores customizados
- Uso de variáveis CSS vs valores diretos
- JIT mode e geração sob demanda
- Análise de bundle size
- Estratégias de otimização

---

### Reflexão 5: Manutenibilidade e Trabalho em Equipe

**Pergunta:** Como você organizaria e documentaria as customizações do Tailwind em um projeto grande com múltiplos desenvolvedores?

**Pontos para considerar:**
- Estrutura do arquivo de configuração
- Documentação de customizações
- Convenções de nomenclatura
- Versionamento e controle de mudanças
- Code review de customizações
- Onboarding de novos desenvolvedores

---

### Reflexão 6: Customizações vs CSS Customizado

**Pergunta:** Quando é melhor criar um utilitário customizado no Tailwind vs escrever CSS customizado tradicional? Dê exemplos de cada caso.

**Pontos para considerar:**
- Complexidade da funcionalidade
- Reutilização do código
- Manutenibilidade
- Performance
- Legibilidade do código
- Integração com o sistema do Tailwind

---

## 🎯 Desafio Final: Criando um Design System Completo

### Tarefa

Crie um design system completo para um projeto fictício de uma **plataforma de educação online**. O design system deve incluir:

**Requisitos:**

1. **Cores:**
   - Primária: `#6366f1` (índigo)
   - Secundária: `#8b5cf6` (roxo)
   - Sucesso: `#10b981` (verde)
   - Aviso: `#f59e0b` (laranja)
   - Erro: `#ef4444` (vermelho)
   - Escala completa (50-950) para primária e secundária

2. **Tipografia:**
   - Sans: `'Inter', sans-serif`
   - Display: `'Poppins', sans-serif`
   - Mono: `'Fira Code', monospace`
   - Tamanhos customizados: `hero` (6rem), `subhero` (3rem)

3. **Espaçamento:**
   - Valores customizados: `13`, `17`, `21`, `25`

4. **Breakpoints:**
   - `xs: 475px`
   - `3xl: 1920px`
   - `tablet: 640px - 1023px` (apenas tablets)

5. **Utilitários customizados:**
   - `.scrollbar-thin`: Scrollbar fina customizada
   - `.text-gradient`: Texto com gradiente

6. **Sombras:**
   - `glow-indigo`: Brilho índigo
   - `glow-purple`: Brilho roxo

### Entregáveis

1. Arquivo `tailwind.config.js` completo e funcional
2. Documentação explicando cada customização e o motivo
3. Exemplos de uso em HTML para cada customização

### Critérios de Avaliação

- ✅ Todas as customizações funcionam corretamente
- ✅ Uso correto de `extend` (não substituir padrões desnecessariamente)
- ✅ Código bem organizado e comentado
- ✅ Documentação clara e útil
- ✅ Exemplos práticos de uso

---

## 📚 Recursos para Aprofundamento

- [Documentação Oficial - Theme Configuration](https://tailwindcss.com/docs/theme)
- [Documentação Oficial - Customizing Colors](https://tailwindcss.com/docs/customizing-colors)
- [Documentação Oficial - Plugins](https://tailwindcss.com/docs/plugins)
- [Tailwind Play](https://play.tailwindcss.com/) - Teste suas configurações online

---

**Bons exercícios! Pratique bastante e reflita sobre as decisões de design e arquitetura! 🚀**

