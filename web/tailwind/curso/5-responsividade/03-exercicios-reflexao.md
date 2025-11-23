# Aula 5 - Exercícios e Reflexão: Responsividade com Tailwind

## 🎯 Exercícios Práticos

### Exercício 1: Card de Produto Responsivo

**Objetivo:** Criar um card de produto que se adapte perfeitamente a diferentes tamanhos de tela.

**Requisitos:**
1. No mobile: card ocupa toda a largura, padding pequeno, texto pequeno
2. No tablet (md): padding médio, texto médio, imagem maior
3. No desktop (lg): padding grande, texto grande, layout mais espaçado
4. A imagem deve crescer proporcionalmente
5. O botão e preço devem ficar em coluna no mobile e em linha no tablet+

**HTML base (complete com classes Tailwind):**
```html
<div class="[SUAS CLASSES AQUI]">
  <img 
    src="https://via.placeholder.com/400x300" 
    alt="Produto"
    class="[SUAS CLASSES AQUI]"
  >
  <h3 class="[SUAS CLASSES AQUI]">
    Smartphone XYZ Pro
  </h3>
  <p class="[SUAS CLASSES AQUI]">
    O smartphone mais avançado do mercado com câmera de 108MP e tela AMOLED.
  </p>
  <div class="[SUAS CLASSES AQUI]">
    <span class="[SUAS CLASSES AQUI]">R$ 2.999,90</span>
    <button class="[SUAS CLASSES AQUI]">
      Comprar Agora
    </button>
  </div>
</div>
```

**Dicas:**
- Use `bg-white`, `rounded-lg`, `shadow-md` para o card
- Use `p-4 md:p-6 lg:p-8` para padding responsivo
- Use `text-lg md:text-xl lg:text-2xl` para o título
- Use `h-32 md:h-48 lg:h-64` para a imagem
- Use `flex flex-col sm:flex-row` para o container de preço/botão

---

### Exercício 2: Grid de Cards Responsivo

**Objetivo:** Criar um grid de cards que muda o número de colunas conforme o tamanho da tela.

**Requisitos:**
1. Mobile: 1 coluna
2. Tablet (sm): 2 colunas
3. Desktop (lg): 3 colunas
4. Desktop grande (xl): 4 colunas
5. Gap deve aumentar em telas maiores
6. Cada card deve ter padding responsivo

**HTML base:**
```html
<div class="container mx-auto px-4 py-8">
  <div class="[SUAS CLASSES AQUI]">
    <div class="[SUAS CLASSES AQUI]">
      <h3>Card 1</h3>
      <p>Conteúdo do card</p>
    </div>
    <div class="[SUAS CLASSES AQUI]">
      <h3>Card 2</h3>
      <p>Conteúdo do card</p>
    </div>
    <div class="[SUAS CLASSES AQUI]">
      <h3>Card 3</h3>
      <p>Conteúdo do card</p>
    </div>
    <div class="[SUAS CLASSES AQUI]">
      <h3>Card 4</h3>
      <p>Conteúdo do card</p>
    </div>
  </div>
</div>
```

**Dicas:**
- Use `grid` com `grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4`
- Use `gap-4 md:gap-6 lg:gap-8` para gap responsivo
- Use `bg-white`, `p-4 md:p-6`, `rounded-lg`, `shadow` para os cards

---

### Exercício 3: Navegação Responsiva

**Objetivo:** Criar uma barra de navegação que muda completamente entre mobile e desktop.

**Requisitos:**
1. No mobile:
   - Logo à esquerda
   - Botão hamburger (☰) à direita
   - Menu de links escondido
2. No desktop (md+):
   - Logo à esquerda
   - Menu de links horizontal à direita
   - Botão hamburger escondido
3. Background azul escuro, texto branco
4. Links devem ter hover effect

**HTML base:**
```html
<nav class="[SUAS CLASSES AQUI]">
  <div class="container mx-auto px-4">
    <div class="[SUAS CLASSES AQUI]">
      <!-- Logo -->
      <div class="[SUAS CLASSES AQUI]">
        MeuSite
      </div>
      
      <!-- Botão Hamburger (só mobile) -->
      <button class="[SUAS CLASSES AQUI]">
        <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h16"></path>
        </svg>
      </button>
      
      <!-- Menu (só desktop) -->
      <div class="[SUAS CLASSES AQUI]">
        <a href="#" class="[SUAS CLASSES AQUI]">Home</a>
        <a href="#" class="[SUAS CLASSES AQUI]">Sobre</a>
        <a href="#" class="[SUAS CLASSES AQUI]">Serviços</a>
        <a href="#" class="[SUAS CLASSES AQUI]">Contato</a>
      </div>
    </div>
  </div>
</nav>
```

**Dicas:**
- Use `bg-blue-900 text-white` para o nav
- Use `flex justify-between items-center` para o container
- Use `hidden md:flex` para o menu desktop
- Use `block md:hidden` para o botão hamburger
- Use `hover:text-blue-200` para hover nos links

---

### Exercício 4: Hero Section Responsiva

**Objetivo:** Criar uma seção hero (banner principal) totalmente responsiva.

**Requisitos:**
1. Título deve crescer: `text-3xl md:text-5xl lg:text-7xl`
2. Subtítulo deve crescer: `text-base md:text-lg lg:text-xl`
3. Padding vertical deve aumentar: `py-8 md:py-16 lg:py-24`
4. Botão CTA deve ter tamanho responsivo
5. Background pode mudar de cor (opcional)

**HTML base:**
```html
<section class="[SUAS CLASSES AQUI]">
  <div class="container mx-auto px-4 text-center">
    <h1 class="[SUAS CLASSES AQUI] mb-4">
      Bem-vindo ao Nosso Site
    </h1>
    <p class="[SUAS CLASSES AQUI] mb-8 max-w-2xl mx-auto">
      Descubra as melhores soluções para o seu negócio com nossa plataforma inovadora.
    </p>
    <button class="[SUAS CLASSES AQUI]">
      Começar Agora
    </button>
  </div>
</section>
```

**Dicas:**
- Use `bg-gradient-to-r from-blue-500 to-purple-600 text-white`
- Use padding vertical responsivo
- Use `px-4 py-2 md:px-6 md:py-3 lg:px-8 lg:py-4` para o botão

---

## 🔍 Exercícios de Análise de Código

### Exercício 5: Identificar Problemas

Analise o código abaixo e identifique os problemas de responsividade:

```html
<div class="p-8 lg:p-4 md:p-2">
  <h1 class="text-6xl md:text-2xl lg:text-4xl">
    Título
  </h1>
  <div class="grid grid-cols-4 md:grid-cols-2 lg:grid-cols-1">
    <div>Card 1</div>
    <div>Card 2</div>
    <div>Card 3</div>
    <div>Card 4</div>
  </div>
</div>
```

**Perguntas:**
1. Qual é o problema com a ordem dos breakpoints no padding?
2. Por que a ordem dos breakpoints no título está errada?
3. O que está errado com o grid? Por que começar com 4 colunas?
4. Como você corrigiria esse código?

---

### Exercício 6: Converter CSS para Tailwind

Converta o seguinte CSS responsivo para classes Tailwind:

```css
.container {
  padding: 1rem;
  max-width: 100%;
}

@media (min-width: 768px) {
  .container {
    padding: 2rem;
    max-width: 768px;
  }
}

@media (min-width: 1024px) {
  .container {
    padding: 3rem;
    max-width: 1024px;
  }
}

.card {
  display: flex;
  flex-direction: column;
  gap: 1rem;
}

@media (min-width: 768px) {
  .card {
    flex-direction: row;
    gap: 2rem;
  }
}

.title {
  font-size: 1.5rem;
  margin-bottom: 1rem;
}

@media (min-width: 768px) {
  .title {
    font-size: 2.25rem;
    margin-bottom: 1.5rem;
  }
}

@media (min-width: 1024px) {
  .title {
    font-size: 3rem;
    margin-bottom: 2rem;
  }
}
```

**Sua resposta (HTML com classes Tailwind):**
```html
<!-- Container -->
<div class="[SUAS CLASSES]">
  
  <!-- Card -->
  <div class="[SUAS CLASSES]">
    <div>Conteúdo 1</div>
    <div>Conteúdo 2</div>
  </div>
  
  <!-- Title -->
  <h1 class="[SUAS CLASSES]">Título</h1>
  
</div>
```

---

## 💭 Perguntas de Reflexão

### Reflexão 1: Mobile-First vs Desktop-First

**Pergunta:** Por que o Tailwind adota a abordagem mobile-first? Quais são as vantagens e desvantagens dessa abordagem comparada com desktop-first?

**Considere:**
- Performance em dispositivos móveis
- Priorização de usuários
- Complexidade de desenvolvimento
- Manutenibilidade do código
- Quando desktop-first poderia ser mais apropriado

**Sua resposta:**

---

### Reflexão 2: Breakpoints e Decisões de Design

**Pergunta:** Os breakpoints padrão do Tailwind (640px, 768px, 1024px, etc.) são adequados para todos os projetos? Em que situações você deveria criar breakpoints customizados?

**Considere:**
- Diferentes dispositivos no mercado
- Projetos com requisitos específicos
- Design systems corporativos
- Acessibilidade e usabilidade
- Trade-off entre customização e padrão

**Sua resposta:**

---

### Reflexão 3: Performance e Bundle Size

**Pergunta:** Como a responsividade com Tailwind afeta o tamanho do CSS final? Quais estratégias você usaria para otimizar um site que usa muitos breakpoints?

**Considere:**
- Quantas classes responsivas são geradas
- Impacto do PurgeCSS/JIT
- Quando usar breakpoints vs CSS customizado
- Balanceamento entre utilitários e componentes
- Análise de bundle size

**Sua resposta:**

---

### Reflexão 4: Manutenibilidade e Legibilidade

**Pergunta:** Quando você tem muitas classes responsivas em um elemento (ex: `class="p-2 sm:p-4 md:p-6 lg:p-8 xl:p-10 text-sm md:text-base lg:text-lg bg-blue-500 md:bg-green-500 lg:bg-purple-500"`), o código pode ficar difícil de ler e manter. Quais estratégias você usaria para melhorar isso?

**Considere:**
- Quando usar `@apply` para criar componentes
- Organização de classes (ferramentas como Headwind)
- Quando escrever CSS customizado
- Trabalho em equipe e code review
- Balanceamento entre utilitários e abstrações

**Sua resposta:**

---

### Reflexão 5: Acessibilidade e Responsividade

**Pergunta:** Como a responsividade se relaciona com acessibilidade? Quais problemas de acessibilidade podem surgir quando você cria designs responsivos?

**Considere:**
- Tamanho de fonte e legibilidade
- Espaçamento e área de toque
- Navegação em diferentes dispositivos
- Contraste e visibilidade
- Leitores de tela e responsividade
- Zoom do navegador

**Sua resposta:**

---

### Reflexão 6: Testes e Debugging

**Pergunta:** Como você testaria a responsividade de um site feito com Tailwind? Quais ferramentas e estratégias são mais eficazes?

**Considere:**
- DevTools do navegador
- Dispositivos reais vs emuladores
- Testes automatizados
- Breakpoints customizados
- Edge cases (telas muito pequenas/grandes)
- Performance em diferentes dispositivos

**Sua resposta:**

---

## 🎯 Desafio Final

### Desafio: Landing Page Responsiva Completa

Crie uma landing page completa e totalmente responsiva usando apenas Tailwind CSS. A página deve incluir:

1. **Header/Navegação:**
   - Logo à esquerda
   - Menu hamburger no mobile, menu horizontal no desktop
   - Background que muda de cor em diferentes breakpoints

2. **Hero Section:**
   - Título responsivo (cresce em telas maiores)
   - Subtítulo responsivo
   - Botão CTA com tamanho responsivo
   - Padding vertical que aumenta

3. **Seção de Features (3 cards):**
   - 1 coluna no mobile
   - 2 colunas no tablet
   - 3 colunas no desktop
   - Cards com padding e texto responsivos

4. **Seção de Testimonials:**
   - Grid responsivo
   - Texto que se adapta
   - Imagens que crescem

5. **Footer:**
   - Layout que muda de coluna (mobile) para linha (desktop)
   - Links organizados responsivamente

**Requisitos Técnicos:**
- Use apenas classes Tailwind (sem CSS customizado)
- Todos os elementos devem ser responsivos
- Use breakpoints de forma consistente
- Código deve ser legível e bem organizado

**Entrega:**
Crie um arquivo HTML completo com sua landing page.

---

## 📝 Checklist de Aprendizado

Após completar os exercícios, verifique se você consegue:

- [ ] Criar layouts responsivos usando prefixos do Tailwind
- [ ] Entender e aplicar a filosofia mobile-first
- [ ] Usar breakpoints padrão e customizados
- [ ] Aplicar responsividade em diferentes propriedades (espaçamento, tipografia, layout)
- [ ] Mostrar/esconder elementos em diferentes breakpoints
- [ ] Criar grids responsivos que mudam o número de colunas
- [ ] Converter CSS responsivo para classes Tailwind
- [ ] Identificar e corrigir problemas de responsividade
- [ ] Pensar criticamente sobre decisões de design responsivo
- [ ] Considerar performance e acessibilidade em designs responsivos

---

**Bons exercícios! Lembre-se: a prática é essencial para dominar responsividade com Tailwind! 🚀**

