# Aula 2: Performance e Boas Práticas

## 🎯 Objetivo

Este documento apresenta boas práticas, otimizações e dicas de performance relacionadas aos conceitos aprendidos nesta aula. Seguir essas práticas desde o início ajudará você a criar código mais eficiente, manutenível e profissional.

---

## 🎨 Boas Práticas: Cores

### 1. Escolha o Formato Apropriado

**Use HEX para cores sólidas:**
- Mais compacto que RGB
- Padrão da indústria
- Fácil de copiar de ferramentas de design

**Use RGB/RGBA quando precisar de transparência:**
- RGBA é mais legível que HEX com alpha
- Mais fácil de ajustar valores

**Use HSL quando precisar ajustar cores:**
- Mais intuitivo para criar variações
- Fácil criar paletas consistentes

**Evite cores nomeadas em produção:**
- Limitadas e podem ter interpretações diferentes
- Use apenas para prototipagem rápida

### 2. Organize Cores com Variáveis CSS

```css
:root {
  --cor-primaria: #3498db;
  --cor-secundaria: #2ecc71;
  --cor-texto: #333333;
  --cor-fundo: #ffffff;
}

.elemento {
  color: var(--cor-primaria);
}
```

**Vantagens:**
- Fácil manutenção (mude uma vez, afeta tudo)
- Consistência em todo o projeto
- Facilita criação de temas

### 3. Use Cores Acessíveis

- Garanta contraste adequado entre texto e fundo
- Teste com ferramentas de acessibilidade
- Considere usuários com daltonismo

---

## 🖼️ Boas Práticas: Background

### 1. Sempre Forneça Fallback

```css
.elemento {
  background-color: #333; /* Fallback */
  background-image: url('imagem.jpg');
}
```

**Por quê?**
- Se a imagem não carregar, a cor aparece
- Melhora a experiência do usuário
- Melhora performance percebida

### 2. Otimize Imagens de Fundo

**Problemas comuns:**
- Imagens muito grandes (afetam carregamento)
- Formato inadequado (use WebP quando possível)
- Múltiplas imagens desnecessárias

**Soluções:**
- Comprima imagens antes de usar
- Use formatos modernos (WebP, AVIF)
- Considere usar gradientes CSS em vez de imagens quando possível
- Use `background-size: cover` para evitar imagens muito grandes

### 3. Use Background Shorthand com Cuidado

**Bom:**
```css
/* Quando você precisa de todas as propriedades */
background: #333 url('img.jpg') center/cover no-repeat;
```

**Evite:**
```css
/* Quando você só precisa de uma propriedade */
background: url('img.jpg'); /* Use background-image em vez disso */
```

**Por quê?**
- Shorthand reseta propriedades não especificadas
- Pode sobrescrever estilos anteriores acidentalmente

### 4. Considere Performance

**Evite:**
- Múltiplas imagens de fundo desnecessárias
- Imagens muito grandes para elementos pequenos
- Animações complexas em background

**Prefira:**
- Gradientes CSS em vez de imagens quando possível
- Imagens otimizadas
- Lazy loading para imagens abaixo da dobra

---

## 📦 Boas Práticas: Box Model

### 1. Sempre Use box-sizing: border-box

```css
*,
*::before,
*::after {
  box-sizing: border-box;
}
```

**Por quê?**
- Facilita cálculos (width inclui padding e border)
- Mais previsível e intuitivo
- Padrão moderno da indústria

**Onde colocar:**
- No início do seu CSS
- Em um reset CSS
- Como padrão global

### 2. Entenda o Tamanho Real dos Elementos

**Lembre-se:**
- Com `content-box`: tamanho total = width + padding + border
- Com `border-box`: tamanho total = width

**Use ferramentas de desenvolvedor:**
- Inspecione elementos para ver o box model visualmente
- Verifique se os tamanhos estão como esperado

### 3. Evite Valores Negativos em Padding

**Problema:**
- Padding negativo não é permitido
- Se precisar de espaço negativo, use margin negativo (com cuidado)

---

## 📏 Boas Práticas: Padding e Margin

### 1. Use Shorthand com Consistência

**Bom:**
```css
/* Todos os lados iguais */
padding: 20px;

/* Vertical e horizontal */
padding: 10px 20px;

/* Todos os lados diferentes */
padding: 10px 20px 15px 25px;
```

**Evite:**
```css
/* Verboso desnecessariamente */
padding-top: 10px;
padding-right: 10px;
padding-bottom: 10px;
padding-left: 10px; /* Use padding: 10px; */
```

### 2. Crie Sistema de Espaçamento

**Use valores consistentes:**
```css
:root {
  --espaco-xs: 4px;
  --espaco-sm: 8px;
  --espaco-md: 16px;
  --espaco-lg: 24px;
  --espaco-xl: 32px;
}

.elemento {
  padding: var(--espaco-md);
  margin-bottom: var(--espaco-lg);
}
```

**Vantagens:**
- Consistência visual
- Fácil manutenção
- Design mais profissional

### 3. Entenda Margin Collapse

**Lembre-se:**
- Margens verticais colapsam (não se somam)
- Margens horizontais não colapsam
- Isso pode causar surpresas se não for entendido

**Solução:**
- Use padding quando possível para evitar colapso
- Ou use apenas uma direção de margin (ex: só margin-bottom)

### 4. Use Margin: Auto para Centralizar

```css
.container {
  width: 800px;
  margin: 0 auto; /* Centraliza horizontalmente */
}
```

**Importante:**
- Só funciona com elementos block
- Requer width definido
- Não funciona verticalmente (use flexbox/grid para isso)

---

## 📐 Boas Práticas: Width e Height

### 1. Evite Altura Fixa quando Possível

**Problema:**
```css
.container {
  height: 300px; /* E se o conteúdo for maior? */
}
```

**Solução:**
```css
.container {
  min-height: 300px; /* Permite crescimento */
  /* ou */
  height: auto; /* Ajusta ao conteúdo */
}
```

**Por quê?**
- Conteúdo pode crescer
- Melhor para responsividade
- Evita overflow indesejado

### 2. Use Unidades Relativas para Responsividade

**Bom:**
```css
.container {
  width: 90%;
  max-width: 1200px;
}
```

**Evite (quando possível):**
```css
.container {
  width: 1200px; /* Quebra em telas pequenas */
}
```

### 3. Combine Unidades para Controle

```css
.container {
  width: min(100%, 1200px); /* Nunca maior que 1200px */
  width: max(300px, 50%); /* Nunca menor que 300px */
  width: clamp(300px, 50%, 1200px); /* Entre limites */
}
```

---

## 🔲 Boas Práticas: Border

### 1. Use Border Shorthand

**Bom:**
```css
border: 2px solid #333;
```

**Evite:**
```css
border-width: 2px;
border-style: solid;
border-color: #333; /* Use shorthand */
```

### 2. Considere Border no Cálculo de Tamanho

**Lembre-se:**
- Border afeta o tamanho total (a menos que use border-box)
- Considere isso ao calcular layouts
- Use border-box para evitar surpresas

### 3. Use Border-Radius Consistentemente

**Crie um sistema:**
```css
:root {
  --raio-sm: 4px;
  --raio-md: 8px;
  --raio-lg: 16px;
  --raio-full: 50%;
}
```

---

## 🔳 Boas Práticas: Outline

### 1. NUNCA Remova Outline sem Alternativa

**❌ Ruim:**
```css
* {
  outline: none; /* Remove acessibilidade! */
}
```

**✅ Bom:**
```css
/* Mantenha outline padrão */
/* Ou forneça alternativa visual clara */
.botao:focus {
  outline: 2px solid blue;
  outline-offset: 2px;
}
```

**Por quê?**
- Outline é crucial para navegação por teclado
- Remover sem alternativa torna o site inacessível
- Pode violar leis de acessibilidade

### 2. Melhore o Outline Visualmente

```css
.botao:focus {
  outline: 2px solid #4A90E2;
  outline-offset: 4px;
  border-radius: 4px;
}
```

---

## 🌑 Boas Práticas: Box Shadow

### 1. Use Sombras Sutis

**Bom:**
```css
box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
```

**Evite:**
```css
box-shadow: 0 10px 20px rgba(0, 0, 0, 0.8); /* Muito forte */
```

**Por quê?**
- Sombras muito fortes distraem
- Design mais profissional com sombras sutis
- Melhor hierarquia visual

### 2. Use RGBA para Transparência

**Bom:**
```css
box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
```

**Evite:**
```css
box-shadow: 0 4px 6px #000000; /* Sem transparência, muito duro */
```

### 3. Crie Sistema de Elevação

```css
:root {
  --sombra-sm: 0 1px 2px rgba(0, 0, 0, 0.05);
  --sombra-md: 0 4px 6px rgba(0, 0, 0, 0.1);
  --sombra-lg: 0 10px 15px rgba(0, 0, 0, 0.1);
  --sombra-xl: 0 20px 25px rgba(0, 0, 0, 0.15);
}
```

**Vantagens:**
- Consistência visual
- Hierarquia clara
- Fácil manutenção

---

## 📏 Boas Práticas: Unidades

### 1. Use Rem para Tipografia e Espaçamento

**Bom:**
```css
body {
  font-size: 1rem; /* 16px padrão */
}

.titulo {
  font-size: 2rem; /* 32px se root for 16px */
  margin-bottom: 1.5rem; /* 24px */
}
```

**Por quê?**
- Respeita preferências do usuário
- Escala proporcionalmente
- Mais acessível

### 2. Use Px para Bordas e Sombras

**Bom:**
```css
.botao {
  border: 1px solid #333; /* Precisão */
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}
```

**Por quê?**
- Bordas e sombras precisam de precisão
- Não precisam escalar com fonte
- Mais previsível

### 3. Use % e vw/vh para Layouts

**Bom:**
```css
.container {
  width: 90%;
  max-width: 1200px;
}

.hero {
  height: 100vh; /* Altura da tela */
}
```

### 4. Evite Unidades Absolutas para Tipografia

**Evite:**
```css
.texto {
  font-size: 14px; /* Não escala */
}
```

**Prefira:**
```css
.texto {
  font-size: 0.875rem; /* 14px mas escalável */
}
```

---

## 🧮 Boas Práticas: Funções

### 1. Use Calc() para Layouts Flexíveis

**Bom:**
```css
.sidebar {
  width: 300px;
}

.conteudo {
  width: calc(100% - 300px); /* Resto do espaço */
}
```

**Por quê?**
- Evita valores fixos
- Mais flexível
- Funciona em diferentes tamanhos de tela

### 2. Use Clamp() para Tipografia Fluida

**Bom:**
```css
.titulo {
  font-size: clamp(1.5rem, 5vw, 3rem);
}
```

**Por quê?**
- Escala suavemente
- Sempre dentro de limites
- Menos media queries necessárias

### 3. Combine Funções

```css
.container {
  width: min(100%, calc(1200px - 2rem));
  padding: clamp(1rem, 5vw, 2rem);
}
```

---

## 🎭 Boas Práticas: Display

### 1. Entenda o Display Padrão

**Lembre-se:**
- Alguns elementos são block por padrão
- Outros são inline por padrão
- Mude apenas quando necessário

### 2. Use Inline-Block com Cuidado

**Problema:**
- Espaços em branco no HTML criam gaps visuais
- Pode causar problemas de alinhamento

**Solução:**
```css
/* Remova espaços ou use font-size: 0 no pai */
.container {
  font-size: 0;
}
.item {
  display: inline-block;
  font-size: 1rem; /* Restaura */
}
```

### 3. Prefira Flexbox/Grid para Layouts

**Em vez de:**
```css
.item {
  display: inline-block;
  width: 33.333%;
}
```

**Prefira (em aulas futuras):**
```css
.container {
  display: flex;
}
```

---

## ⚡ Performance

### 1. Minimize Propriedades de Background

**Evite:**
```css
/* Múltiplas imagens desnecessárias */
background-image: url('img1.jpg'), url('img2.jpg'), url('img3.jpg');
```

**Prefira:**
- Uma imagem quando possível
- Gradientes CSS em vez de imagens
- Sprites para ícones pequenos

### 2. Use Will-Change com Cuidado

```css
/* Apenas para elementos que realmente animam */
.animado {
  will-change: transform;
}
```

**Por quê?**
- Pode melhorar performance de animações
- Mas use apenas quando necessário
- Pode consumir mais memória

### 3. Evite Reflows Desnecessários

**Problema:**
- Mudanças em width, height, margin causam reflow
- Múltiplos reflows são custosos

**Solução:**
- Agrupe mudanças quando possível
- Use transform em vez de position quando animando
- Evite ler e escrever propriedades de layout em loop

---

## 🎯 Checklist de Boas Práticas

### Cores
- [ ] Uso variáveis CSS para cores
- [ ] Escolho formato apropriado (HEX, RGB, HSL)
- [ ] Garanto contraste adequado
- [ ] Organizo cores de forma consistente

### Background
- [ ] Sempre forneço cor de fallback
- [ ] Otimizo imagens antes de usar
- [ ] Uso formatos modernos quando possível
- [ ] Evito múltiplas imagens desnecessárias

### Box Model
- [ ] Uso box-sizing: border-box globalmente
- [ ] Entendo o tamanho real dos elementos
- [ ] Uso ferramentas de desenvolvedor para verificar

### Padding e Margin
- [ ] Uso shorthand quando apropriado
- [ ] Criei sistema de espaçamento consistente
- [ ] Entendo margin collapse
- [ ] Uso margin: auto para centralizar

### Width e Height
- [ ] Evito altura fixa quando possível
- [ ] Uso unidades relativas para responsividade
- [ ] Combino unidades para controle

### Border
- [ ] Uso border shorthand
- [ ] Considero border no cálculo de tamanho
- [ ] Uso border-radius consistentemente

### Outline
- [ ] Nunca removo outline sem alternativa
- [ ] Melhoro outline visualmente quando necessário
- [ ] Garanto acessibilidade

### Box Shadow
- [ ] Uso sombras sutis
- [ ] Uso RGBA para transparência
- [ ] Criei sistema de elevação

### Unidades
- [ ] Uso rem para tipografia
- [ ] Uso px para bordas e sombras
- [ ] Uso % e vw/vh para layouts
- [ ] Evito unidades absolutas para tipografia

### Funções
- [ ] Uso calc() para layouts flexíveis
- [ ] Uso clamp() para tipografia fluida
- [ ] Combino funções quando apropriado

### Display
- [ ] Entendo display padrão dos elementos
- [ ] Mudo display apenas quando necessário
- [ ] Prefiro flexbox/grid para layouts complexos

---

## 🚀 Próximos Passos

Agora que você conhece as boas práticas:
- Aplique-as nos seus projetos
- Revise código antigo com essas práticas em mente
- Continue aprendendo e adaptando

Lembre-se: boas práticas evoluem. O importante é entender o "porquê" por trás de cada uma, não apenas decorá-las.




