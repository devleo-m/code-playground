# Aula 6 - Performance, Boas Práticas e Otimização: Estados e Interatividade

## 🚀 Performance de Estados e Transições

### O Problema: Reflow e Repaint

Quando você aplica estados e transições, o navegador precisa recalcular o layout (reflow) e redesenhar elementos (repaint). Isso pode impactar significativamente a performance, especialmente em dispositivos móveis.

#### Propriedades que Causam Reflow

**Reflow** ocorre quando o navegador precisa recalcular o layout da página:

```css
/* ❌ RUIM - Causa reflow */
.element:hover {
  width: 200px;        /* Reflow */
  height: 100px;       /* Reflow */
  padding: 20px;       /* Reflow */
  margin: 10px;         /* Reflow */
  border: 2px solid;   /* Reflow */
}
```

**Equivalente Tailwind (evitar):**
```html
<div class="hover:w-48 hover:h-24 hover:p-5 hover:m-2 hover:border-2">
  Elemento que causa reflow
</div>
```

#### Propriedades que Causam Apenas Repaint

**Repaint** ocorre quando apenas a aparência visual muda, sem recalcular layout:

```css
/* ✅ BOM - Apenas repaint */
.element:hover {
  background-color: blue;  /* Repaint apenas */
  color: white;            /* Repaint apenas */
  opacity: 0.8;            /* Repaint apenas */
  box-shadow: 0 4px 8px;   /* Repaint apenas */
}
```

**Equivalente Tailwind (preferir):**
```html
<div class="hover:bg-blue-500 hover:text-white hover:opacity-80 hover:shadow-lg">
  Elemento otimizado
</div>
```

#### Propriedades que Usam GPU (Melhor Performance)

**Transform** e **opacity** podem ser acelerados pela GPU:

```css
/* ✅ EXCELENTE - Usa GPU */
.element:hover {
  transform: scale(1.1);     /* GPU acceleration */
  opacity: 0.9;              /* GPU acceleration */
}
```

**Equivalente Tailwind:**
```html
<div class="hover:scale-110 hover:opacity-90">
  Elemento com GPU acceleration
</div>
```

### Otimizando Transições

#### ❌ Evite: `transition-all`

```html
<!-- ❌ RUIM - Anima todas as propriedades -->
<div class="transition-all hover:bg-blue-500 hover:scale-110 hover:opacity-80">
  Anima tudo (ineficiente)
</div>
```

**Problemas:**
- Anima propriedades que não precisam ser animadas
- Pode causar reflow desnecessário
- Maior uso de CPU/GPU

#### ✅ Prefira: Transições Específicas

```html
<!-- ✅ BOM - Anima apenas o necessário -->
<div class="
  transition-colors        <!-- Apenas cores -->
  hover:bg-blue-500
  transition-transform     <!-- Apenas transform -->
  hover:scale-110
">
  Anima apenas o necessário
</div>
```

**Vantagens:**
- Melhor performance
- Mais controle sobre o que anima
- Menor uso de recursos

#### ✅ Melhor: Combine Propriedades Compatíveis

```html
<!-- ✅ EXCELENTE - Propriedades que usam GPU -->
<div class="
  transition-transform     <!-- Transform usa GPU -->
  transition-opacity      <!-- Opacity usa GPU -->
  hover:scale-110
  hover:opacity-80
">
  Performance otimizada
</div>
```

### Duração das Transições

#### Escolhendo a Duração Correta

```html
<!-- ✅ BOM - Durações apropriadas -->
<button class="
  transition-colors duration-150    <!-- Rápido para cores -->
  hover:bg-blue-600
">
  Botão
</button>

<div class="
  transition-transform duration-300  <!-- Médio para transform -->
  hover:scale-105
">
  Card
</div>

<div class="
  transition-all duration-500        <!-- Lento para mudanças complexas -->
  hover:shadow-xl
  hover:-translate-y-4
">
  Elemento complexo
</div>
```

**Regra de Ouro:**
- **Cores e opacidade**: 150-200ms (rápido)
- **Transform**: 200-300ms (médio)
- **Múltiplas propriedades**: 300-500ms (médio-lento)

### Timing Functions

#### Escolhendo a Curva Correta

```html
<!-- ✅ BOM - Timing functions apropriadas -->
<button class="
  transition-colors 
  ease-in-out      <!-- Suave para início e fim -->
  duration-200
  hover:bg-blue-600
">
  Botão suave
</button>

<div class="
  transition-transform
  ease-out         <!-- Rápido no início, lento no fim -->
  duration-300
  hover:scale-110
">
  Elemento com "bounce"
</div>
```

---

## 🎯 Boas Práticas de Estados

### 1. Sempre Forneça Estados de Focus

#### ❌ Ruim: Sem Focus Visível

```html
<button class="bg-blue-500 hover:bg-blue-600 text-white px-4 py-2">
  Botão sem focus
</button>
```

**Problemas:**
- Usuários de teclado não sabem onde estão
- Não atende diretrizes WCAG
- Pobre experiência de acessibilidade

#### ✅ Bom: Focus Visível e Claro

```html
<button class="
  bg-blue-500 
  hover:bg-blue-600 
  focus:outline-none
  focus:ring-2
  focus:ring-blue-500
  focus:ring-offset-2
  text-white 
  px-4 py-2
">
  Botão acessível
</button>
```

**Vantagens:**
- Acessível para navegação por teclado
- Atende diretrizes WCAG 2.1
- Melhor experiência para todos os usuários

### 2. Use `focus-visible:` Quando Apropriado

#### Quando Usar `focus-visible:`

```html
<!-- ✅ BOM - Focus apenas quando navegação por teclado -->
<button class="
  bg-blue-500
  hover:bg-blue-600
  focus-visible:ring-2      <!-- Apenas com teclado -->
  focus-visible:ring-blue-500
  text-white
  px-4 py-2
">
  Botão inteligente
</button>
```

**Por que usar:**
- Remove o anel de foco quando clicamos com mouse (melhor UX)
- Mantém o anel quando navegamos com teclado (acessibilidade)
- Melhor dos dois mundos

### 3. Estados Disabled Devem Ser Claros

#### ❌ Ruim: Disabled Pouco Visível

```html
<button 
  disabled
  class="bg-blue-500 text-white px-4 py-2"
>
  Botão desabilitado
</button>
```

**Problemas:**
- Usuário não sabe que está desabilitado
- Pode tentar clicar várias vezes
- Frustração

#### ✅ Bom: Disabled Claro e Óbvio

```html
<button 
  disabled
  class="
    bg-gray-400
    text-white
    px-4 py-2
    opacity-50
    cursor-not-allowed
    disabled:opacity-50
    disabled:cursor-not-allowed
  "
>
  Botão desabilitado
</button>
```

**Vantagens:**
- Visualmente claro que está desabilitado
- Cursor indica que não pode clicar
- Melhor feedback visual

### 4. Evite Hover em Elementos Não Interativos

#### ❌ Ruim: Hover em Texto Normal

```html
<p class="hover:text-blue-500">
  Este texto não é clicável, mas parece ser
</p>
```

**Problemas:**
- Confunde usuários (parece clicável)
- Expectativa não atendida
- Pobre UX

#### ✅ Bom: Hover Apenas em Elementos Interativos

```html
<a href="#" class="hover:text-blue-500 underline">
  Link clicável
</a>

<button class="hover:bg-blue-600">
  Botão clicável
</button>
```

**Vantagens:**
- Comportamento previsível
- Melhor UX
- Segue convenções da web

### 5. Consistência de Estados

#### ❌ Ruim: Estados Inconsistentes

```html
<!-- Botão 1 -->
<button class="bg-blue-500 hover:bg-blue-600">Botão 1</button>

<!-- Botão 2 -->
<button class="bg-green-500 hover:bg-green-700">Botão 2</button>

<!-- Botão 3 -->
<button class="bg-red-500 hover:bg-red-800">Botão 3</button>
```

**Problemas:**
- Cada botão se comporta diferente
- Usuário não sabe o que esperar
- Interface inconsistente

#### ✅ Bom: Estados Consistentes

```html
<!-- Todos os botões seguem o mesmo padrão -->
<button class="
  bg-blue-500
  hover:bg-blue-600
  active:bg-blue-700
  focus:ring-2
  focus:ring-blue-500
  transition-colors
  duration-200
">
  Botão Primário
</button>

<button class="
  bg-gray-500
  hover:bg-gray-600
  active:bg-gray-700
  focus:ring-2
  focus:ring-gray-500
  transition-colors
  duration-200
">
  Botão Secundário
</button>
```

**Vantagens:**
- Comportamento previsível
- Interface consistente
- Melhor aprendizado do usuário

---

## 🎨 Organização e Manutenibilidade

### 1. Agrupe Classes por Categoria

#### ❌ Ruim: Classes Desorganizadas

```html
<button class="
  hover:bg-blue-600
  bg-blue-500
  text-white
  focus:ring-2
  px-4
  active:bg-blue-700
  py-2
  rounded-lg
  transition-colors
  focus:ring-blue-500
">
  Botão
</button>
```

**Problemas:**
- Difícil de ler
- Difícil de manter
- Fácil de cometer erros

#### ✅ Bom: Classes Organizadas

```html
<button class="
  /* Layout */
  px-4 py-2 rounded-lg
  
  /* Cores */
  bg-blue-500 text-white
  
  /* Estados */
  hover:bg-blue-600
  active:bg-blue-700
  focus:ring-2 focus:ring-blue-500
  
  /* Transições */
  transition-colors duration-200
">
  Botão
</button>
```

**Vantagens:**
- Fácil de ler
- Fácil de manter
- Padrão claro

### 2. Use Componentes para Padrões Repetidos

#### ❌ Ruim: Repetir Classes

```html
<button class="bg-blue-500 hover:bg-blue-600 ...">Botão 1</button>
<button class="bg-blue-500 hover:bg-blue-600 ...">Botão 2</button>
<button class="bg-blue-500 hover:bg-blue-600 ...">Botão 3</button>
```

**Problemas:**
- Código repetitivo
- Difícil de manter
- Mudanças requerem atualizar múltiplos lugares

#### ✅ Bom: Usar @apply ou Componentes

```css
/* Usando @apply */
.btn-primary {
  @apply bg-blue-500 hover:bg-blue-600 active:bg-blue-700 
         focus:ring-2 focus:ring-blue-500 text-white 
         px-4 py-2 rounded-lg transition-colors duration-200;
}
```

```html
<button class="btn-primary">Botão 1</button>
<button class="btn-primary">Botão 2</button>
<button class="btn-primary">Botão 3</button>
```

**Vantagens:**
- DRY (Don't Repeat Yourself)
- Fácil de manter
- Mudanças em um lugar

---

## ♿ Acessibilidade

### 1. Contraste de Cores em Estados

#### ❌ Ruim: Contraste Insuficiente

```html
<button class="
  bg-gray-200
  hover:bg-gray-300    <!-- Contraste muito baixo -->
  text-gray-400
">
  Botão com baixo contraste
</button>
```

**Problemas:**
- Não atende WCAG AA (4.5:1)
- Difícil de ler
- Pobre acessibilidade

#### ✅ Bom: Contraste Adequado

```html
<button class="
  bg-blue-500
  hover:bg-blue-600
  text-white            <!-- Contraste alto (4.5:1+) -->
">
  Botão acessível
</button>
```

**Vantagens:**
- Atende diretrizes WCAG
- Legível para todos
- Melhor acessibilidade

### 2. Tamanho de Área de Toque

#### ❌ Ruim: Área de Toque Pequena

```html
<button class="px-2 py-1 text-sm">
  Clique
</button>
```

**Problemas:**
- Difícil de clicar em mobile
- Não atende diretrizes (mínimo 44x44px)
- Pobre UX em touch devices

#### ✅ Bom: Área de Toque Adequada

```html
<button class="
  px-4 py-3           <!-- Área mínima 44x44px -->
  min-h-[44px]        <!-- Garante altura mínima -->
  min-w-[44px]        <!-- Garante largura mínima -->
">
  Clique
</button>
```

**Vantagens:**
- Fácil de clicar
- Atende diretrizes
- Melhor UX mobile

### 3. Redução de Movimento

#### ✅ Bom: Respeitar `prefers-reduced-motion`

```css
/* CSS puro */
@media (prefers-reduced-motion: reduce) {
  * {
    animation-duration: 0.01ms !important;
    animation-iteration-count: 1 !important;
    transition-duration: 0.01ms !important;
  }
}
```

**No Tailwind:**
```html
<!-- Use motion-reduce: quando disponível -->
<div class="
  transition-transform duration-300
  motion-reduce:transition-none
  hover:scale-110
">
  Elemento que respeita preferências
</div>
```

**Por que é importante:**
- Alguns usuários têm sensibilidade a movimento
- Melhor acessibilidade
- Respeita preferências do usuário

---

## 🔍 Debugging e Troubleshooting

### 1. Estados Não Funcionando

#### Problema: Hover Não Funciona em Mobile

**Causa:** Hover não existe em dispositivos touch

**Solução:**
```html
<!-- Use active: para touch devices -->
<button class="
  hover:bg-blue-600
  active:bg-blue-600      <!-- Funciona em touch -->
">
  Botão
</button>
```

#### Problema: Focus Não Aparece

**Causa:** Outline padrão foi removido sem substituição

**Solução:**
```html
<!-- Sempre substitua outline removido -->
<button class="
  focus:outline-none      <!-- Remove outline padrão -->
  focus:ring-2            <!-- Adiciona ring customizado -->
  focus:ring-blue-500
">
  Botão
</button>
```

### 2. Transições Não Suaves

#### Problema: Transição "Travando"

**Causas possíveis:**
- Animar propriedades que causam reflow
- Muitas propriedades animando simultaneamente
- Duração muito curta ou muito longa

**Solução:**
```html
<!-- Anime apenas propriedades otimizadas -->
<div class="
  transition-transform    <!-- Apenas transform -->
  transition-opacity      <!-- Apenas opacity -->
  duration-300            <!-- Duração adequada -->
  hover:scale-110
  hover:opacity-80
">
  Elemento otimizado
</div>
```

### 3. Group e Peer Não Funcionando

#### Problema: `group-hover:` Não Funciona

**Causa:** Esqueceu de adicionar `group` no elemento pai

**Solução:**
```html
<!-- ✅ CORRETO -->
<div class="group">
  <div class="group-hover:bg-blue-500">Filho</div>
</div>
```

#### Problema: `peer-focus:` Não Funciona

**Causa:** Elemento peer não está adjacente ou não tem `peer` class

**Solução:**
```html
<!-- ✅ CORRETO -->
<input type="checkbox" class="peer" />
<div class="peer-checked:bg-blue-500">Irmão</div>
```

---

## 📊 Métricas de Performance

### Como Medir Performance de Estados

#### 1. Chrome DevTools - Performance Tab

1. Abra DevTools (F12)
2. Vá para a aba "Performance"
3. Clique em "Record"
4. Interaja com elementos (hover, click, etc.)
5. Pare a gravação
6. Analise:
   - Tempo de repaint
   - FPS (frames per second)
   - Uso de CPU/GPU

#### 2. Lighthouse - Performance Score

1. Abra DevTools
2. Vá para a aba "Lighthouse"
3. Execute análise
4. Verifique:
   - Performance score
   - First Contentful Paint
   - Time to Interactive

#### 3. CSS Triggers

Use [csstriggers.com](https://csstriggers.com) para verificar quais propriedades causam reflow/repaint.

---

## 🎯 Checklist de Performance

Use este checklist ao criar estados e interações:

### Transições
- [ ] Uso `transition-colors` ou `transition-transform` ao invés de `transition-all` quando possível
- [ ] Duração das transições é apropriada (150-300ms para a maioria)
- [ ] Timing function é adequada para o tipo de interação

### Propriedades Animadas
- [ ] Prefiro `transform` e `opacity` (GPU acceleration)
- [ ] Evito animar `width`, `height`, `padding`, `margin` (causam reflow)
- [ ] Limito o número de propriedades animadas simultaneamente

### Estados
- [ ] Todos os elementos interativos têm estados de focus visíveis
- [ ] Estados disabled são claros e óbvios
- [ ] Hover é usado apenas em elementos interativos
- [ ] Estados são consistentes em toda a aplicação

### Acessibilidade
- [ ] Contraste de cores atende WCAG AA (4.5:1)
- [ ] Áreas de toque têm no mínimo 44x44px
- [ ] Respeito `prefers-reduced-motion` quando possível
- [ ] Navegação por teclado funciona corretamente

### Organização
- [ ] Classes estão organizadas por categoria
- [ ] Padrões repetidos usam componentes ou @apply
- [ ] Código é legível e manutenível

---

## 🚀 Otimizações Avançadas

### 1. Will-Change para Animações Complexas

```html
<!-- Use will-change apenas quando necessário -->
<div class="
  will-change-transform    <!-- Avisa o navegador -->
  transition-transform
  hover:scale-110
">
  Elemento com animação complexa
</div>
```

**⚠️ Atenção:** Use `will-change` com moderação. Pode causar overhead se usado excessivamente.

### 2. Contain para Isolar Animações

```css
/* CSS puro - isola reflow/repaint */
.animated-container {
  contain: layout style paint;
}
```

**Vantagens:**
- Isola reflow/repaint
- Melhor performance
- Animações mais suaves

### 3. GPU Acceleration Explícita

```html
<!-- Force GPU acceleration -->
<div class="
  transform-gpu           <!-- Se disponível no Tailwind -->
  transition-transform
  hover:scale-110
">
  Elemento com GPU
</div>
```

---

## 📚 Recursos Adicionais

### Ferramentas
- [Chrome DevTools](https://developer.chrome.com/docs/devtools/)
- [CSS Triggers](https://csstriggers.com)
- [WebPageTest](https://www.webpagetest.org/)
- [Lighthouse](https://developers.google.com/web/tools/lighthouse)

### Documentação
- [WCAG 2.1 Guidelines](https://www.w3.org/WAI/WCAG21/quickref/)
- [MDN: CSS Transitions](https://developer.mozilla.org/en-US/docs/Web/CSS/CSS_Transitions)
- [Tailwind: Transitions](https://tailwindcss.com/docs/transition-property)

### Artigos
- [High Performance Animations](https://web.dev/animations/)
- [Accessible Focus Indicators](https://www.a11yproject.com/posts/never-remove-css-outlines/)
- [CSS Performance Optimization](https://developer.mozilla.org/en-US/docs/Learn/Performance/CSS)

---

## 🎯 Resumo: Melhores Práticas

### ✅ Faça
- Use `transition-colors` ou `transition-transform` ao invés de `transition-all`
- Sempre forneça estados de focus visíveis
- Use `transform` e `opacity` para animações (GPU acceleration)
- Mantenha duração de transições entre 150-300ms
- Organize classes por categoria
- Teste em diferentes dispositivos (mouse e touch)
- Respeite `prefers-reduced-motion`

### ❌ Evite
- Animar propriedades que causam reflow (`width`, `height`, `padding`, `margin`)
- Usar `transition-all` desnecessariamente
- Remover outline sem substituir por ring
- Hover em elementos não interativos
- Estados inconsistentes
- Contraste de cores insuficiente
- Áreas de toque muito pequenas

---

**Lembre-se**: Performance e acessibilidade não são opcionais. São fundamentais para criar uma experiência web de qualidade para todos os usuários.

