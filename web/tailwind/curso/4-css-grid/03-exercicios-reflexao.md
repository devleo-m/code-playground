# Aula 4 - Exercícios e Reflexão: CSS Grid com Tailwind

## 🎯 Exercícios Práticos

### Exercício 1: Grid Básico

Crie um grid com 4 colunas e 6 itens. Cada item deve ter:
- Background azul (`bg-blue-500`)
- Padding de 1rem (`p-4`)
- Espaçamento de 1rem entre os itens (`gap-4`)

**Resultado esperado:** 6 itens distribuídos em 4 colunas (2 linhas completas).

---

### Exercício 2: Spanning de Colunas

Crie um grid com 4 colunas onde:
- O primeiro item ocupa 2 colunas (`col-span-2`)
- O segundo item ocupa 1 coluna (normal)
- O terceiro item ocupa 1 coluna (normal)
- O quarto item ocupa 3 colunas (`col-span-3`)
- O quinto item ocupa 1 coluna (normal)

Use cores diferentes para cada item para visualizar melhor.

**Desafio:** Faça o layout funcionar corretamente mesmo com os spans diferentes.

---

### Exercício 3: Layout de Blog

Crie um layout de blog responsivo com:
- **Header:** Ocupa toda a largura, background azul escuro, texto branco
- **Sidebar:** Ocupa 3 colunas em desktop, toda a largura em mobile
- **Conteúdo Principal:** Ocupa 9 colunas em desktop, toda a largura em mobile
- **Footer:** Ocupa toda a largura, background cinza escuro, texto branco

Use o sistema de 12 colunas do Tailwind.

**Breakpoints:**
- Mobile: 1 coluna (sidebar e conteúdo empilhados)
- Desktop (lg): 12 colunas (sidebar 3 + conteúdo 9)

---

### Exercício 4: Galeria de Imagens Responsiva

Crie uma galeria de imagens que:
- No mobile: 1 coluna
- No tablet (md): 2 colunas
- No desktop (lg): 3 colunas
- No large (xl): 4 colunas

Adicione 8 itens na galeria. Cada item deve ser um quadrado com background cinza.

**Dica:** Use `grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4`

---

### Exercício 5: Dashboard com Cards

Crie um dashboard com:
- **4 cards pequenos** na primeira linha (cada um ocupa 3 colunas em desktop)
- **1 card grande** na segunda linha (ocupa todas as 12 colunas)
- **2 cards médios** na terceira linha (cada um ocupa 6 colunas)

Use o sistema de 12 colunas. Em mobile, todos os cards devem ocupar toda a largura.

---

### Exercício 6: Grid com Gap Diferente

Crie um grid de 3 colunas onde:
- O espaçamento horizontal (`gap-x`) é maior que o vertical (`gap-y`)
- Use `gap-x-8` e `gap-y-4`

Observe como o espaçamento afeta a aparência do grid.

---

### Exercício 7: Alinhamento de Itens

Crie um grid de 3 colunas com 3 itens. Cada item deve ter:
- Largura fixa (`w-32`)
- Altura fixa (`h-32`)
- Backgrounds diferentes

Teste os diferentes alinhamentos:
1. `place-items-start` - itens no topo
2. `place-items-center` - itens no centro
3. `place-items-end` - itens na parte inferior
4. `place-items-stretch` - itens esticados (remova width e height)

---

### Exercício 8: Grid Aninhado

Crie um grid principal de 2 colunas. Dentro de cada coluna, crie um grid secundário de 2 colunas.

**Estrutura:**
```
Grid Principal (2 colunas)
  ├─ Coluna 1
  │   └─ Grid Secundário (2 colunas)
  │       ├─ Item 1.1
  │       └─ Item 1.2
  └─ Coluna 2
      └─ Grid Secundário (2 colunas)
          ├─ Item 2.1
          └─ Item 2.2
```

---

### Exercício 9: Layout Complexo com Spanning

Crie um layout que simula uma página de revista com:
- **Cabeçalho:** Ocupa 12 colunas, altura pequena
- **Artigo Principal:** Ocupa 8 colunas, altura grande
- **Sidebar:** Ocupa 4 colunas, altura grande
- **Galeria:** 3 imagens pequenas (4 colunas cada) abaixo do artigo
- **Rodapé:** Ocupa 12 colunas, altura pequena

Use cores diferentes para cada seção.

---

### Exercício 10: Grid Responsivo Avançado

Crie um grid que muda completamente em diferentes breakpoints:

**Mobile (padrão):**
- 1 coluna
- Todos os itens empilhados

**Tablet (md):**
- 2 colunas
- Primeiro item ocupa 2 colunas (largura total)
- Demais itens em 2 colunas

**Desktop (lg):**
- 4 colunas
- Primeiro item ocupa 4 colunas (largura total)
- Demais itens em 4 colunas

**Large (xl):**
- 6 colunas
- Primeiro item ocupa 6 colunas (largura total)
- Demais itens em 6 colunas

---

## 🤔 Perguntas de Reflexão

### Reflexão 1: Grid vs Flexbox - Quando Usar Cada Um?

**Pergunta:** Em que situações você escolheria Grid sobre Flexbox? E vice-versa?

**Pense sobre:**
- Layouts unidimensionais vs bidimensionais
- Necessidade de controle preciso sobre linhas e colunas
- Complexidade do layout
- Performance e manutenibilidade

**Exemplo para analisar:**
```html
<!-- Este layout deveria usar Grid ou Flexbox? Por quê? -->
<div>
  <header>Header</header>
  <div>
    <aside>Sidebar</aside>
    <main>Main Content</main>
  </div>
  <footer>Footer</footer>
</div>
```

---

### Reflexão 2: Sistema de 12 Colunas - É Sempre a Melhor Escolha?

**Pergunta:** O sistema de 12 colunas é universalmente aplicável? Quando você consideraria usar um número diferente de colunas?

**Pense sobre:**
- Layouts que não se encaixam bem em 12 colunas
- Quando usar `grid-cols-5`, `grid-cols-7`, etc.
- Trade-offs entre flexibilidade e simplicidade
- Casos onde CSS customizado seria melhor

**Exemplo para analisar:**
```html
<!-- Este grid de 5 colunas é uma boa escolha? -->
<div class="grid grid-cols-5 gap-4">
  <!-- 5 itens iguais -->
</div>
```

---

### Reflexão 3: Spanning e Manutenibilidade

**Pergunta:** Como o uso excessivo de `col-span` e `row-span` pode afetar a manutenibilidade do código? Quando spanning se torna problemático?

**Pense sobre:**
- Legibilidade do código
- Facilidade de modificação futura
- Responsividade e breakpoints
- Quando criar componentes ao invés de usar spanning

**Exemplo para analisar:**
```html
<!-- Este código é fácil de manter? -->
<div class="grid grid-cols-12">
  <div class="col-span-3 md:col-span-4 lg:col-span-2 xl:col-span-3">Item 1</div>
  <div class="col-span-6 md:col-span-4 lg:col-span-5 xl:col-span-6">Item 2</div>
  <div class="col-span-3 md:col-span-4 lg:col-span-5 xl:col-span-3">Item 3</div>
</div>
```

---

### Reflexão 4: Performance de Grid vs Flexbox

**Pergunta:** Grid e Flexbox têm diferenças de performance? Em que cenários isso importa?

**Pense sobre:**
- Renderização do navegador
- Complexidade de cálculos
- Número de elementos
- Animações e transições
- Compatibilidade com navegadores antigos

**Exemplo para analisar:**
```html
<!-- Para 1000 itens, Grid ou Flexbox seria mais performático? -->
<div class="grid grid-cols-10 gap-2">
  <!-- 1000 divs aqui -->
</div>
```

---

### Reflexão 5: Grid Responsivo - Mobile-First ou Desktop-First?

**Pergunta:** Ao criar grids responsivos, qual abordagem é melhor: começar do mobile e expandir, ou começar do desktop e reduzir?

**Pense sobre:**
- Filosofia mobile-first do Tailwind
- Complexidade de breakpoints
- Manutenibilidade
- Performance
- Experiência do usuário

**Exemplo para analisar:**
```html
<!-- Abordagem 1: Mobile-first -->
<div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3">

<!-- Abordagem 2: Desktop-first (não recomendado) -->
<div class="grid grid-cols-3 md:grid-cols-2 lg:grid-cols-1">
```

---

### Reflexão 6: Gap e Espaçamento Consistente

**Pergunta:** Como você garante espaçamento consistente em um projeto grande usando Grid? Quando usar `gap` uniforme vs `gap-x` e `gap-y` diferentes?

**Pense sobre:**
- Design systems e consistência
- Escala de espaçamento do Tailwind
- Quando quebrar a consistência é aceitável
- Manutenibilidade em equipe

**Exemplo para analisar:**
```html
<!-- Qual abordagem é melhor para um design system? -->
<div class="grid grid-cols-3 gap-4">
  <!-- Abordagem 1: Gap uniforme -->
</div>

<div class="grid grid-cols-3 gap-x-8 gap-y-4">
  <!-- Abordagem 2: Gap diferente -->
</div>
```

---

### Reflexão 7: Grid Aninhado - Quando é Demais?

**Pergunta:** Grids aninhados podem se tornar problemáticos? Qual é o limite prático de aninhamento?

**Pense sobre:**
- Complexidade mental
- Performance
- Manutenibilidade
- Alternativas (Flexbox interno, componentes)
- Quando simplificar a estrutura

**Exemplo para analisar:**
```html
<!-- Este nível de aninhamento é aceitável? -->
<div class="grid grid-cols-2">
  <div class="grid grid-cols-2">
    <div class="grid grid-cols-2">
      <!-- Mais aninhamento aqui? -->
    </div>
  </div>
</div>
```

---

### Reflexão 8: Acessibilidade e Grid

**Pergunta:** Como o uso de Grid afeta a acessibilidade? O que você precisa considerar?

**Pense sobre:**
- Ordem visual vs ordem do DOM
- Leitores de tela
- Navegação por teclado
- Contraste e espaçamento
- Responsividade para usuários com deficiência visual

**Exemplo para analisar:**
```html
<!-- Este grid mantém a ordem lógica para leitores de tela? -->
<div class="grid grid-cols-12">
  <aside class="col-span-3 order-2">Sidebar</aside>
  <main class="col-span-9 order-1">Main Content</main>
</div>
```

---

### Reflexão 9: Grid Template Areas - Quando Vale a Pena?

**Pergunta:** Quando você usaria `grid-template-areas` (CSS customizado) ao invés de classes Tailwind utilitárias?

**Pense sobre:**
- Complexidade do layout
- Legibilidade do código
- Manutenibilidade
- Trade-off entre Tailwind puro e CSS customizado
- Quando `@apply` seria útil

**Exemplo para analisar:**
```html
<!-- Abordagem 1: Tailwind puro -->
<div class="grid grid-cols-12">
  <header class="col-span-12">Header</header>
  <aside class="col-span-3">Sidebar</aside>
  <main class="col-span-9">Main</main>
</div>

<!-- Abordagem 2: Grid template areas (CSS customizado) -->
<div class="grid" style="grid-template-areas: 'header' 'sidebar main' 'footer';">
  <header style="grid-area: header;">Header</header>
  <aside style="grid-area: sidebar;">Sidebar</aside>
  <main style="grid-area: main;">Main</main>
</div>
```

---

### Reflexão 10: Tailwind Grid vs CSS Grid Puro

**Pergunta:** Quando você escreveria CSS Grid puro ao invés de usar as classes Tailwind? Quais são os limites das utilities do Tailwind?

**Pense sobre:**
- Casos complexos que Tailwind não cobre
- Auto-fit e auto-fill
- Grid template areas complexos
- Animações de grid
- Quando a produtividade do Tailwind não compensa

**Exemplo para analisar:**
```html
<!-- Tailwind consegue fazer isso facilmente? -->
<div class="grid" style="grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));">
  <!-- Itens que se adaptam automaticamente -->
</div>
```

---

## 📝 Instruções para Resolução

1. **Exercícios Práticos:**
   - Resolva cada exercício escrevendo o código HTML completo
   - Teste em diferentes tamanhos de tela
   - Use o DevTools para inspecionar o grid
   - Experimente diferentes valores para entender o comportamento

2. **Perguntas de Reflexão:**
   - Responda cada pergunta com base no que você aprendeu
   - Pense em exemplos práticos da sua experiência
   - Considere diferentes cenários e casos de uso
   - Não há resposta "certa" - o importante é o raciocínio

3. **Dicas:**
   - Use o Tailwind Play (https://play.tailwindcss.com/) para testar rapidamente
   - Inspecione elementos no DevTools para ver o CSS gerado
   - Experimente quebrar coisas para entender melhor
   - Compare Grid com Flexbox na prática

---

## ✅ Critérios de Sucesso

Você dominou esta aula quando conseguir:
- ✅ Criar grids básicos e complexos com Tailwind
- ✅ Usar spanning de forma eficiente
- ✅ Criar layouts responsivos com Grid
- ✅ Decidir quando usar Grid vs Flexbox
- ✅ Entender as limitações e quando usar CSS customizado
- ✅ Pensar criticamente sobre performance e manutenibilidade

---

**Bons estudos! 🚀**

Lembre-se: a prática é essencial. Não apenas leia os exercícios - escreva o código e veja funcionando!

