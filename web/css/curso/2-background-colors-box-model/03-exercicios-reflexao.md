# Aula 2: Exercícios e Reflexão

## 🎯 Objetivo dos Exercícios

Estes exercícios foram criados para ajudar você a consolidar o aprendizado sobre Background, Colors, Box Model e fundamentos de layout. Eles são progressivos - comece pelos mais simples e vá avançando.

---

## 📝 Exercícios Práticos

### Exercício 1: Explorando Cores

**Objetivo:** Familiarizar-se com diferentes formatos de cores.

**Tarefa:**
1. Crie um arquivo HTML simples com 5 divs
2. Estilize cada div com uma cor diferente usando formatos diferentes:
   - Div 1: cor nomeada (ex: `blue`)
   - Div 2: RGB (ex: `rgb(255, 0, 0)`)
   - Div 3: HEX (ex: `#00FF00`)
   - Div 4: HSL (ex: `hsl(240, 100%, 50%)`)
   - Div 5: RGBA com transparência (ex: `rgba(0, 0, 255, 0.5)`)

**Reflexão:**
- Qual formato você achou mais fácil de usar?
- Quando você usaria cada formato?
- Como a transparência (RGBA) muda a aparência?

---

### Exercício 2: Background Básico

**Objetivo:** Praticar propriedades de background.

**Tarefa:**
1. Crie um elemento com:
   - Cor de fundo azul
   - Padding de 20px
   - Texto branco dentro

2. Crie outro elemento com:
   - Imagem de fundo (use uma imagem qualquer)
   - Background-size: cover
   - Background-position: center
   - Texto sobre a imagem

**Reflexão:**
- Como a cor de fundo afeta a legibilidade do texto?
- O que acontece se você mudar o background-size para `contain`?
- Por que é importante ter uma cor de fundo mesmo quando há imagem?

---

### Exercício 3: Entendendo o Box Model

**Objetivo:** Visualizar os componentes do box model.

**Tarefa:**
1. Crie um elemento com:
   - Width: 200px
   - Height: 100px
   - Padding: 20px
   - Border: 5px solid black
   - Margin: 30px
   - Background-color: lightblue

2. Use as ferramentas de desenvolvedor do navegador (F12) para inspecionar o elemento e ver visualmente o box model.

**Reflexão:**
- Qual é o tamanho total do elemento? (width + padding + border)
- O que acontece se você mudar para `box-sizing: border-box`?
- Como o margin afeta o espaço total ocupado?

---

### Exercício 4: Padding vs Margin

**Objetivo:** Entender a diferença prática entre padding e margin.

**Tarefa:**
1. Crie dois botões lado a lado:
   - Botão 1: padding de 10px, margin de 5px, background azul
   - Botão 2: padding de 5px, margin de 10px, background vermelho

2. Observe a diferença visual entre eles.

**Reflexão:**
- Qual botão parece ter mais espaço "dentro"?
- Qual botão tem mais espaço "ao redor"?
- Se você mudar a cor de fundo, o que acontece com o padding? E com o margin?
- Qual você usaria para criar espaço clicável em um botão?

---

### Exercício 5: Bordas e Cantos Arredondados

**Objetivo:** Praticar diferentes estilos de borda.

**Tarefa:**
1. Crie 4 elementos, cada um com um estilo de borda diferente:
   - Elemento 1: `border: 2px solid black`
   - Elemento 2: `border: 2px dashed blue`
   - Elemento 3: `border: 2px dotted green`
   - Elemento 4: `border: 2px solid red` com `border-radius: 10px`

**Reflexão:**
- Qual estilo de borda você acha mais apropriado para diferentes situações?
- Como o border-radius muda a aparência?
- Quando você usaria cada estilo?

---

### Exercício 6: Box Shadow

**Objetivo:** Criar profundidade com sombras.

**Tarefa:**
1. Crie três cards:
   - Card 1: sem sombra
   - Card 2: sombra sutil (`box-shadow: 0 2px 4px rgba(0,0,0,0.1)`)
   - Card 3: sombra mais pronunciada (`box-shadow: 0 4px 8px rgba(0,0,0,0.2)`)

**Reflexão:**
- Qual card parece mais "elevado"?
- Como a sombra afeta a percepção de importância?
- Quando você usaria sombras mais fortes vs. mais sutis?

---

### Exercício 7: Unidades CSS

**Objetivo:** Experimentar diferentes unidades.

**Tarefa:**
1. Crie elementos com diferentes unidades:
   - Elemento 1: `width: 200px` (absoluta)
   - Elemento 2: `width: 50%` (relativa ao pai)
   - Elemento 3: `width: 50vw` (relativa à viewport)
   - Elemento 4: `font-size: 2rem` (relativa ao root)

2. Redimensione a janela do navegador e observe o que acontece.

**Reflexão:**
- Quais elementos mudam quando você redimensiona a janela?
- Por que alguns mudam e outros não?
- Quando você usaria cada tipo de unidade?

---

### Exercício 8: Função Calc()

**Objetivo:** Usar cálculos em CSS.

**Tarefa:**
1. Crie um container com `width: 100%`
2. Dentro, crie um elemento filho com `width: calc(100% - 40px)` e `margin: 20px`

**Reflexão:**
- Por que usar calc() em vez de apenas porcentagem?
- Em que situações calc() é útil?
- Que outros cálculos você poderia fazer?

---

### Exercício 9: Display Block vs Inline

**Objetivo:** Entender o comportamento de display.

**Tarefa:**
1. Crie 3 elementos `<span>` (normalmente inline):
   - Span 1: deixe como está (inline)
   - Span 2: `display: block`
   - Span 3: `display: inline-block` com `width: 100px`

2. Observe como cada um se comporta.

**Reflexão:**
- Por que o span 2 quebra a linha?
- Qual a diferença entre inline e inline-block?
- Quando você mudaria o display de um elemento?

---

### Exercício 10: Projeto Completo - Card de Produto

**Objetivo:** Combinar todos os conceitos aprendidos.

**Tarefa:**
Crie um card de produto que inclua:
- Cor de fundo ou imagem de fundo
- Padding adequado
- Border com cantos arredondados
- Box shadow para profundidade
- Margin para espaçamento
- Diferentes cores para texto e fundo
- Largura definida (pode usar calc, %, ou px)

**Reflexão:**
- Como cada propriedade contribui para o design final?
- O que você mudaria para melhorar?
- Quais conceitos foram mais desafiadores?

---

## 🤔 Perguntas de Reflexão

### Sobre Cores:

1. **Por que existem tantos formatos de cores?**
   - Pense nas vantagens de cada um. Quando você escolheria cada formato?

2. **Qual a diferença prática entre RGB e HSL?**
   - Em que situações HSL seria mais útil?

3. **Por que RGBA é diferente de usar opacity?**
   - Quando você usaria cada um?

### Sobre Background:

4. **Por que é importante ter uma cor de fundo mesmo quando há imagem?**
   - O que acontece se a imagem não carregar?

5. **Quando você usaria background-repeat vs no-repeat?**
   - Dê exemplos práticos de cada caso.

6. **Qual a diferença entre cover e contain?**
   - Em que situações cada um é mais apropriado?

### Sobre Box Model:

7. **Por que entender o box model é fundamental?**
   - Como isso afeta seu trabalho diário como desenvolvedor?

8. **Qual a diferença prática entre padding e margin?**
   - Dê exemplos de quando usar cada um.

9. **Por que box-sizing: border-box é considerado uma boa prática?**
   - Que problemas ele resolve?

### Sobre Unidades:

10. **Por que unidades relativas são importantes para acessibilidade?**
    - Como isso ajuda usuários com necessidades especiais?

11. **Quando você usaria vw/vh em vez de %?**
    - Qual a diferença prática?

12. **Por que clamp() é útil para tipografia?**
    - Que problemas ele resolve que outras abordagens não resolvem?

### Sobre Display:

13. **Por que alguns elementos são block por padrão e outros inline?**
    - Qual o propósito de cada um?

14. **Quando você mudaria o display de um elemento?**
    - Dê exemplos práticos.

15. **Qual a diferença entre inline e inline-block?**
    - Quando você usaria cada um?

---

## 🎓 Desafios Avançados

### Desafio 1: Layout Responsivo Básico

Crie um layout que:
- Use unidades relativas (rem, %, vw)
- Se adapte a diferentes tamanhos de tela
- Use calc() para espaçamentos
- Tenha padding e margin proporcionais

### Desafio 2: Paleta de Cores Consistente

Crie uma paleta de 5 cores usando HSL:
- Uma cor principal
- Uma cor secundária
- Uma cor de destaque
- Uma cor de texto
- Uma cor de fundo

Use apenas HSL para criar variações (mais claro, mais escuro, mais saturado).

### Desafio 3: Card com Múltiplas Camadas

Crie um card que demonstre:
- Background com gradiente ou imagem
- Padding interno
- Border com estilo interessante
- Box shadow com múltiplas camadas
- Outline quando em foco (para acessibilidade)

### Desafio 4: Tipografia Fluida

Use clamp() para criar:
- Títulos que escalam entre 24px e 48px baseado na largura da tela
- Texto de parágrafo que escala entre 16px e 20px
- Espaçamento que se adapta proporcionalmente

---

## ✅ Checklist de Aprendizado

Marque os conceitos que você domina:

### Cores
- [ ] Entendo a diferença entre cores nomeadas, RGB, HEX, HSL
- [ ] Sei quando usar RGBA vs opacity
- [ ] Consigo criar cores com transparência
- [ ] Entendo como ajustar cores usando HSL

### Background
- [ ] Sei aplicar cor de fundo
- [ ] Consigo usar imagens de fundo
- [ ] Entendo background-position e background-size
- [ ] Sei controlar background-repeat
- [ ] Consigo usar background shorthand

### Box Model
- [ ] Entendo os 4 componentes (content, padding, border, margin)
- [ ] Sei a diferença entre padding e margin
- [ ] Entendo box-sizing: border-box
- [ ] Consigo calcular o tamanho total de um elemento

### Padding e Margin
- [ ] Sei usar padding individual e shorthand
- [ ] Sei usar margin individual e shorthand
- [ ] Entendo margin collapse
- [ ] Consigo centralizar elementos com margin: auto

### Width e Height
- [ ] Entendo a diferença entre valores fixos e relativos
- [ ] Sei quando usar auto
- [ ] Consigo criar layouts responsivos com unidades relativas

### Border
- [ ] Sei aplicar bordas com diferentes estilos
- [ ] Consigo usar border-radius
- [ ] Entendo como border afeta o tamanho do elemento

### Outline
- [ ] Entendo a diferença entre border e outline
- [ ] Sei a importância do outline para acessibilidade
- [ ] Não removo outline sem alternativa

### Box Shadow
- [ ] Consigo criar sombras básicas
- [ ] Entendo os componentes da sombra (offset, blur, spread, color)
- [ ] Sei criar sombras internas (inset)
- [ ] Consigo usar múltiplas sombras

### Unidades
- [ ] Entendo a diferença entre unidades absolutas e relativas
- [ ] Sei quando usar px, rem, em, %, vw, vh
- [ ] Consigo escolher a unidade apropriada para cada situação

### Funções
- [ ] Sei usar calc() para cálculos
- [ ] Entendo min() e max()
- [ ] Consigo usar clamp() para tipografia fluida

### Display
- [ ] Entendo block, inline, inline-block
- [ ] Sei quando mudar o display de um elemento
- [ ] Entendo como display afeta o layout

---

## 💡 Dicas para os Exercícios

1. **Use as ferramentas de desenvolvedor:** O F12 do navegador mostra visualmente o box model
2. **Experimente valores:** Mude valores e veja o que acontece
3. **Compare lado a lado:** Crie versões diferentes e compare
4. **Pense visualmente:** Use as analogias da aula simplificada
5. **Não tenha pressa:** Domine cada conceito antes de avançar

---

## 🎯 Próximos Passos

Depois de completar estes exercícios, você deve:
- Sentir confiança ao trabalhar com cores e backgrounds
- Entender claramente o box model
- Saber quando usar padding vs margin
- Conseguir escolher unidades apropriadas
- Compreender como display afeta o layout

Se algum conceito ainda estiver confuso, revise a aula principal e simplificada antes de avançar para a próxima aula.




