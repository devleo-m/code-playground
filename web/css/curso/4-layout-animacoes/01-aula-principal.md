# Aula 4: Layout e Animações em CSS

## 🎯 Introdução

Nesta aula, você aprenderá sobre os diferentes sistemas de layout disponíveis em CSS e como criar animações e transições que tornam suas páginas mais interativas e profissionais. Esses conceitos são fundamentais para criar layouts modernos e experiências de usuário agradáveis.

---

## 📄 Flow Layout (Layout Normal do Documento)

### O que é Flow Layout?

**Flow Layout** é o sistema de layout padrão do CSS. É como os elementos se comportam naturalmente quando você não aplica nenhum sistema de layout especial como Flexbox ou Grid. É o comportamento "normal" que os elementos têm por padrão.

### Por que Flow Layout é Importante?

Entender o flow layout é fundamental porque é a base de tudo. Antes de aprender sistemas mais complexos, você precisa entender como os elementos se comportam naturalmente. É como aprender a andar antes de correr.

### Como Funciona o Flow Layout?

No flow layout, os elementos são posicionados um após o outro, seguindo a ordem em que aparecem no HTML, como palavras em uma frase. O navegador lê o HTML de cima para baixo e posiciona os elementos nessa ordem.

#### Elementos Block no Flow Layout

Elementos **block** (como `<div>`, `<p>`, `<h1>`) se comportam assim:
- Ocupam **toda a largura disponível** do container pai
- Começam em uma **nova linha**
- Empilham verticalmente, um embaixo do outro
- Respeitam margin e padding normalmente

**Pense assim:** É como colocar caixas uma embaixo da outra. Cada caixa ocupa toda a largura da prateleira.

#### Elementos Inline no Flow Layout

Elementos **inline** (como `<span>`, `<a>`, `<strong>`) se comportam assim:
- Ocupam apenas o **espaço necessário** para seu conteúdo
- Ficam na **mesma linha** que outros elementos inline
- Flutuam horizontalmente, um ao lado do outro
- Não respeitam width, height, margin vertical completo

**Pense assim:** É como palavras em uma frase. Elas ficam uma ao lado da outra na mesma linha.

### Características do Flow Layout

1. **Ordem Natural**: Os elementos aparecem na ordem do HTML
2. **Empilhamento Vertical**: Elementos block se empilham verticalmente
3. **Fluxo Horizontal**: Elementos inline fluem horizontalmente
4. **Sem Controle Complexo**: Você não tem controle preciso sobre posicionamento complexo
5. **Simples e Previsível**: É o comportamento mais básico e fácil de entender

### Quando Usar Flow Layout?

- Para conteúdo simples que não precisa de layout complexo
- Quando a ordem natural do HTML é suficiente
- Para textos e parágrafos simples
- Como base antes de aplicar outros sistemas de layout

### Limitações do Flow Layout

- Não permite alinhamento fácil de elementos
- Dificulta criar layouts lado a lado complexos
- Não oferece controle sobre distribuição de espaço
- Limitado para designs modernos e responsivos

**Resumo:** Flow Layout é o comportamento padrão, simples e direto. É a base, mas para layouts modernos você precisará de Flexbox ou Grid.

---

## 🎯 Flexbox (Layout Flexível)

### O que é Flexbox?

**Flexbox** é um sistema de layout unidimensional que permite organizar elementos em uma linha ou coluna, com controle total sobre alinhamento, distribuição de espaço e ordem dos elementos. É chamado de "flexível" porque os elementos podem crescer, encolher e se adaptar ao espaço disponível.

### Por que Flexbox é Importante?

Flexbox revolucionou o CSS porque resolve problemas que eram muito difíceis antes, como centralizar elementos verticalmente, distribuir espaço igualmente, ou fazer elementos se adaptarem ao tamanho disponível. É uma das ferramentas mais usadas em CSS moderno.

### Como Funciona o Flexbox?

Flexbox funciona criando um **container flex** (o elemento pai) que controla como seus **itens flex** (os elementos filhos) se comportam. Quando você aplica `display: flex` em um elemento, ele se torna um container flex e seus filhos diretos se tornam itens flex.

#### Conceitos Fundamentais

1. **Flex Container**: O elemento pai que tem `display: flex`
2. **Flex Items**: Os elementos filhos diretos do container
3. **Main Axis (Eixo Principal)**: A direção principal (linha ou coluna)
4. **Cross Axis (Eixo Transversal)**: A direção perpendicular ao eixo principal

### Direção do Flexbox

A propriedade `flex-direction` controla a direção dos itens:

- **`row`** (padrão): Itens em linha, da esquerda para direita
- **`row-reverse`**: Itens em linha, da direita para esquerda
- **`column`**: Itens em coluna, de cima para baixo
- **`column-reverse`**: Itens em coluna, de baixo para cima

**Pense assim:** É como organizar livros em uma prateleira. Você pode colocá-los horizontalmente (row) ou verticalmente (column), e pode começar da esquerda ou direita.

### Alinhamento no Flexbox

Flexbox oferece controle poderoso sobre alinhamento:

#### Justify-Content (Alinhamento no Eixo Principal)

Controla como os itens são distribuídos ao longo do eixo principal:
- **`flex-start`**: Alinha no início (esquerda em row, topo em column)
- **`flex-end`**: Alinha no final (direita em row, baixo em column)
- **`center`**: Centraliza os itens
- **`space-between`**: Espaço igual entre itens, sem espaço nas extremidades
- **`space-around`**: Espaço igual ao redor de cada item
- **`space-evenly`**: Espaço igual entre todos os itens e extremidades

#### Align-Items (Alinhamento no Eixo Transversal)

Controla como os itens são alinhados perpendicularmente:
- **`flex-start`**: Alinha no início do eixo transversal
- **`flex-end`**: Alinha no final do eixo transversal
- **`center`**: Centraliza no eixo transversal
- **`stretch`**: Estica os itens para preencher o container
- **`baseline`**: Alinha pela linha de base do texto

**Pense assim:** `justify-content` é como você distribui pessoas em uma fila (horizontalmente), e `align-items` é como você alinha a altura delas (verticalmente).

### Flexibilidade dos Itens

Os itens flex podem crescer e encolher:

- **`flex-grow`**: Define se o item pode crescer para preencher espaço extra
- **`flex-shrink`**: Define se o item pode encolher quando há pouco espaço
- **`flex-basis`**: Define o tamanho inicial do item antes de crescer/encolher

**Pense assim:** É como pessoas em uma mesa. Algumas podem esticar os braços mais (grow), outras podem encolher (shrink), e cada uma tem um tamanho inicial (basis).

### Quando Usar Flexbox?

- Para centralizar elementos (horizontal e verticalmente)
- Para criar navegações horizontais
- Para distribuir espaço igualmente entre elementos
- Para layouts em uma dimensão (linha ou coluna)
- Para fazer elementos se adaptarem ao espaço disponível
- Para criar cards e botões que se ajustam

### Vantagens do Flexbox

- Fácil centralização (problema que era muito difícil antes)
- Controle sobre distribuição de espaço
- Elementos flexíveis que se adaptam
- Alinhamento poderoso
- Ordem dos elementos pode ser mudada com CSS

**Resumo:** Flexbox é perfeito para layouts em uma dimensão, centralização e distribuição de espaço. É uma das ferramentas mais úteis do CSS moderno.

---

## 📊 CSS Grid (Layout em Grade)

### O que é CSS Grid?

**CSS Grid** é um sistema de layout bidimensional que permite criar layouts complexos usando linhas e colunas. Diferente do Flexbox que trabalha em uma dimensão, o Grid trabalha em duas dimensões simultaneamente, oferecendo controle preciso sobre onde cada elemento deve estar posicionado.

### Por que CSS Grid é Importante?

Grid resolve problemas que Flexbox não consegue resolver sozinho. Enquanto Flexbox é ótimo para uma dimensão, Grid é perfeito para layouts complexos que precisam de controle em linhas E colunas ao mesmo tempo, como layouts de página completos, galerias de imagens ou dashboards.

### Como Funciona o CSS Grid?

Grid funciona criando um **grid container** (o elemento pai com `display: grid`) que divide o espaço em uma grade de linhas e colunas. Os elementos filhos se posicionam nessa grade, ocupando uma ou mais células.

#### Conceitos Fundamentais

1. **Grid Container**: O elemento pai com `display: grid`
2. **Grid Items**: Os elementos filhos que se posicionam na grade
3. **Grid Lines**: As linhas que dividem a grade (horizontais e verticais)
4. **Grid Tracks**: Os espaços entre duas linhas (linhas ou colunas)
5. **Grid Cells**: A interseção de uma linha e uma coluna
6. **Grid Areas**: Grupos de células nomeadas

### Definindo Colunas e Linhas

Você define a estrutura da grade com:

- **`grid-template-columns`**: Define quantas colunas e seus tamanhos
- **`grid-template-rows`**: Define quantas linhas e seus tamanhos

**Exemplo conceitual:**
- `grid-template-columns: 200px 200px 200px` cria 3 colunas de 200px cada
- `grid-template-columns: 1fr 1fr 1fr` cria 3 colunas de tamanho igual (fr = fração)
- `grid-template-columns: repeat(3, 1fr)` faz a mesma coisa de forma mais concisa

**Pense assim:** É como criar uma tabela. Você define quantas colunas quer e qual o tamanho de cada uma.

### Posicionando Itens na Grade

Você pode posicionar itens de várias formas:

- **`grid-column`**: Define em quais colunas o item vai ocupar
- **`grid-row`**: Define em quais linhas o item vai ocupar
- **`grid-area`**: Define uma área nomeada que o item vai ocupar

**Pense assim:** É como dizer "este elemento vai da coluna 1 até a 3, e da linha 2 até a 4".

### Gap (Espaçamento)

A propriedade `gap` cria espaço entre as células da grade:
- **`gap`**: Espaço igual entre linhas e colunas
- **`row-gap`**: Espaço entre linhas
- **`column-gap`**: Espaço entre colunas

**Pense assim:** É como o espaçamento entre azulejos. Você define quanto espaço quer entre cada azulejo.

### Alinhamento no Grid

Grid também oferece controle sobre alinhamento:

- **`justify-items`**: Alinha itens horizontalmente dentro de suas células
- **`align-items`**: Alinha itens verticalmente dentro de suas células
- **`justify-content`**: Alinha a grade inteira horizontalmente no container
- **`align-content`**: Alinha a grade inteira verticalmente no container

### Quando Usar CSS Grid?

- Para layouts de página completos (header, sidebar, main, footer)
- Para criar galerias de imagens
- Para dashboards e painéis complexos
- Para layouts que precisam de controle em duas dimensões
- Para criar designs com elementos de tamanhos diferentes
- Para layouts que precisam se reorganizar em diferentes breakpoints

### Vantagens do CSS Grid

- Controle bidimensional preciso
- Layouts complexos ficam mais simples
- Alinhamento poderoso
- Responsividade mais fácil
- Menos código necessário para layouts complexos

### Grid vs Flexbox: Quando Usar Cada Um?

- **Use Flexbox** quando:
  - Você precisa de layout em uma dimensão (linha OU coluna)
  - Precisa centralizar elementos
  - Precisa distribuir espaço entre elementos
  - Trabalhando com componentes pequenos (botões, cards, navegação)

- **Use Grid** quando:
  - Você precisa de layout em duas dimensões (linhas E colunas)
  - Criando layouts de página completos
  - Precisa de controle preciso sobre posicionamento
  - Trabalhando com layouts complexos

**Resumo:** Grid é perfeito para layouts bidimensionais complexos. É como ter controle total sobre uma grade, posicionando elementos exatamente onde você quer.

---

## 📰 Multicolumn Layout (Layout em Múltiplas Colunas)

### O que é Multicolumn Layout?

**Multicolumn Layout** permite dividir o conteúdo de um elemento em múltiplas colunas, similar a um jornal ou revista. O texto flui automaticamente de uma coluna para a próxima, criando um layout de múltiplas colunas sem precisar estruturar o HTML de forma especial.

### Por que Multicolumn Layout é Importante?

É útil para melhorar a legibilidade de textos longos, especialmente em telas largas. Em vez de ter uma linha muito longa de texto (difícil de ler), você pode dividir o conteúdo em colunas mais estreitas, que são mais fáceis de ler.

### Como Funciona o Multicolumn Layout?

Você aplica propriedades de coluna em um elemento, e o navegador automaticamente divide o conteúdo desse elemento em colunas. O texto flui de uma coluna para a próxima automaticamente.

#### Propriedades Principais

- **`column-count`**: Define quantas colunas você quer
- **`column-width`**: Define a largura mínima de cada coluna
- **`column-gap`**: Define o espaço entre as colunas
- **`column-rule`**: Define uma linha divisória entre colunas (similar a border)

**Pense assim:** É como um jornal. Você define quantas colunas quer, e o texto se distribui automaticamente entre elas.

### Quando Usar Multicolumn Layout?

- Para artigos longos e textos extensos
- Para melhorar legibilidade em telas largas
- Para criar layouts estilo jornal ou revista
- Para listas que ficam muito longas horizontalmente
- Para conteúdo que se beneficia de colunas estreitas

### Limitações do Multicolumn Layout

- Não oferece controle preciso sobre qual conteúdo vai em qual coluna
- O fluxo é automático, você não controla a ordem
- Pode não funcionar bem em telas muito pequenas
- Não é ideal para layouts complexos

**Resumo:** Multicolumn Layout é perfeito para textos longos que você quer dividir em colunas para melhorar a legibilidade, como em jornais.

---

## 🎈 Floating Elements (Elementos Flutuantes)

### O que são Floating Elements?

**Floating** é uma técnica antiga do CSS que permite "flutuar" um elemento para a esquerda ou direita, fazendo com que outros elementos fluam ao redor dele. Era muito usado antes do Flexbox e Grid existirem.

### Por que Entender Floating?

Embora não seja mais a forma recomendada para criar layouts principais, o float ainda é útil em situações específicas, como fazer texto fluir ao redor de imagens. Além disso, entender float ajuda a entender a evolução do CSS e a resolver problemas em código legado.

### Como Funciona o Float?

Quando você aplica `float: left` ou `float: right` em um elemento, ele é removido do fluxo normal e flutua para o lado especificado. Os elementos ao redor fluem ao redor do elemento flutuante.

#### Valores do Float

- **`left`**: Flutua para a esquerda
- **`right`**: Flutua para a direita
- **`none`**: Não flutua (comportamento normal)

**Pense assim:** É como uma imagem em um artigo de jornal. A imagem fica de um lado, e o texto flui ao redor dela.

### Quando Usar Float?

- Para fazer texto fluir ao redor de imagens (caso de uso clássico)
- Em código legado que já usa float
- Para pequenos ajustes de posicionamento

### Quando NÃO Usar Float?

- Para criar layouts principais (use Flexbox ou Grid)
- Para centralizar elementos (use Flexbox)
- Para criar navegações (use Flexbox)
- Para layouts modernos (use Flexbox ou Grid)

### Limitações do Float

- Não foi criado para layouts, apenas para texto ao redor de imagens
- Pode causar problemas de layout difíceis de resolver
- Requer "limpar" o float para evitar problemas
- Não oferece controle preciso sobre posicionamento

**Resumo:** Float é uma técnica antiga, ainda útil para texto ao redor de imagens, mas não deve ser usado para layouts principais. Prefira Flexbox ou Grid para layouts modernos.

---

## ✨ Transitions (Transições)

### O que são Transitions?

**Transitions** permitem que mudanças de propriedades CSS aconteçam de forma suave e gradual, em vez de instantânea. Quando você muda uma propriedade (como cor, tamanho, posição), a transição cria uma animação suave entre o estado antigo e o novo.

### Por que Transitions são Importantes?

Transições tornam as interações muito mais agradáveis. Em vez de mudanças bruscas que podem parecer "quebradas", as transições criam uma sensação de fluidez e profissionalismo. Elas melhoram significativamente a experiência do usuário.

### Como Funcionam as Transitions?

Você define quais propriedades devem ter transição, quanto tempo a transição deve durar, e como a velocidade deve variar durante a transição. O navegador então anima automaticamente a mudança.

#### Propriedades de Transition

- **`transition-property`**: Quais propriedades devem ter transição (ex: `color`, `width`, `all`)
- **`transition-duration`**: Quanto tempo a transição leva (ex: `0.3s`, `500ms`)
- **`transition-timing-function`**: Como a velocidade varia (ex: `ease`, `linear`, `ease-in-out`)
- **`transition-delay`**: Atraso antes da transição começar

#### Shorthand (Forma Abreviada)

Você pode usar `transition` para definir tudo de uma vez:
- `transition: property duration timing-function delay`

**Pense assim:** É como uma porta que abre suavemente em vez de bater. Você define quanto tempo leva para abrir e como a velocidade muda durante a abertura.

### Timing Functions (Funções de Tempo)

As timing functions controlam como a velocidade muda durante a transição:

- **`ease`** (padrão): Começa devagar, acelera, depois desacelera
- **`linear`**: Velocidade constante durante toda a transição
- **`ease-in`**: Começa devagar e acelera
- **`ease-out`**: Começa rápido e desacelera
- **`ease-in-out`**: Começa devagar, acelera no meio, desacelera no final

**Pense assim:** É como diferentes formas de acelerar um carro. `ease-in` é como acelerar gradualmente, `ease-out` é como frear gradualmente.

### Quando Usar Transitions?

- Para mudanças de cor em hover (botões, links)
- Para mudanças de tamanho suaves
- Para mostrar/ocultar elementos suavemente
- Para qualquer mudança de propriedade que você quer que seja suave
- Para melhorar feedback visual em interações

### Exemplos Comuns

- Botão que muda de cor ao passar o mouse
- Card que cresce ligeiramente ao passar o mouse
- Menu que aparece suavemente
- Elemento que muda de posição suavemente

**Resumo:** Transitions tornam mudanças suaves e profissionais. São essenciais para criar interfaces que se sentem polidas e responsivas.

---

## 🎬 Keyframe Animations (Animações com Keyframes)

### O que são Keyframe Animations?

**Keyframe Animations** permitem criar animações complexas definindo múltiplos pontos (keyframes) ao longo da animação. Diferente das transições que apenas animam entre dois estados, as animações com keyframes permitem definir o que acontece em vários momentos da animação.

### Por que Keyframe Animations são Importantes?

Enquanto transições são ótimas para mudanças simples entre estados, keyframe animations permitem criar animações mais complexas e controladas. Você pode criar animações que se repetem, que vão e voltam, ou que têm múltiplas etapas.

### Como Funcionam as Keyframe Animations?

Você define uma sequência de **keyframes** (pontos-chave) usando `@keyframes`, especificando como o elemento deve estar em cada ponto (0%, 25%, 50%, 100%, etc.). Depois, você aplica essa animação a um elemento usando a propriedade `animation`.

#### Criando Keyframes

Você define keyframes assim:
```css
@keyframes nomeDaAnimacao {
  0% { /* estado inicial */ }
  50% { /* estado no meio */ }
  100% { /* estado final */ }
}
```

**Pense assim:** É como um filme. Você define como o elemento deve estar em diferentes "frames" da animação, e o navegador preenche os frames entre eles.

#### Propriedades de Animation

- **`animation-name`**: Nome da animação definida em @keyframes
- **`animation-duration`**: Quanto tempo a animação leva
- **`animation-timing-function`**: Como a velocidade varia
- **`animation-delay`**: Atraso antes de começar
- **`animation-iteration-count`**: Quantas vezes repetir (número ou `infinite`)
- **`animation-direction`**: Direção (`normal`, `reverse`, `alternate`)
- **`animation-fill-mode`**: Como o elemento fica antes/depois da animação

#### Shorthand (Forma Abreviada)

Você pode usar `animation` para definir tudo:
- `animation: name duration timing-function delay iteration-count direction fill-mode`

### Quando Usar Keyframe Animations?

- Para animações que se repetem (loading spinners, pulos)
- Para animações complexas com múltiplas etapas
- Para animações que vão e voltam
- Para criar efeitos visuais elaborados
- Quando transições não são suficientes

### Diferença entre Transitions e Animations

- **Transitions**: Animam mudanças entre estados (hover, focus, etc.)
- **Animations**: Criam animações independentes que podem se repetir

**Pense assim:** Transitions são como uma porta que abre quando você passa o mouse. Animations são como um ventilador que fica girando continuamente.

**Resumo:** Keyframe Animations permitem criar animações complexas e controladas. São perfeitas para animações que se repetem ou têm múltiplas etapas.

---

## 🔄 Transforms (Transformações)

### O que são Transforms?

**Transforms** permitem modificar visualmente elementos usando operações como rotação, escala, translação (movimento) e inclinação (skew). Diferente de outras propriedades, transforms não afetam o fluxo do documento - outros elementos não se movem para dar espaço.

### Por que Transforms são Importantes?

Transforms permitem criar efeitos visuais interessantes sem afetar o layout. Você pode rotacionar, aumentar, mover elementos sem que isso empurre outros elementos. É muito útil para criar interações visuais e animações.

### Como Funcionam os Transforms?

Você aplica funções de transform em um elemento. O elemento é transformado visualmente, mas continua ocupando seu espaço original no layout. Outros elementos não são afetados.

#### Funções de Transform Principais

- **`translate()`**: Move o elemento (horizontal e/ou verticalmente)
- **`rotate()`**: Rotaciona o elemento (em graus)
- **`scale()`**: Aumenta ou diminui o tamanho do elemento
- **`skew()`**: Inclina o elemento (distorce)

**Pense assim:** É como pegar uma foto e movê-la, girá-la, aumentar ou distorcer, mas sem afetar o espaço que ela ocupa na parede.

### Translate (Translação/Movimento)

Move o elemento sem afetar o layout:
- `translateX(20px)` - move 20px para a direita
- `translateY(-10px)` - move 10px para cima
- `translate(20px, -10px)` - move horizontal e verticalmente

**Pense assim:** É como deslizar um objeto sem empurrar outros objetos.

### Rotate (Rotação)

Rotaciona o elemento:
- `rotate(45deg)` - rotaciona 45 graus no sentido horário
- `rotate(-90deg)` - rotaciona 90 graus no sentido anti-horário

**Pense assim:** É como girar um objeto no lugar.

### Scale (Escala)

Aumenta ou diminui o tamanho:
- `scale(1.5)` - aumenta 50%
- `scale(0.5)` - diminui 50%
- `scale(2, 0.5)` - aumenta horizontalmente e diminui verticalmente

**Pense assim:** É como usar zoom em uma imagem, mas sem afetar o espaço que ela ocupa.

### Skew (Inclinação)

Inclina/distorce o elemento:
- `skewX(20deg)` - inclina horizontalmente
- `skewY(10deg)` - inclina verticalmente

**Pense assim:** É como inclinar um objeto, criando um efeito de perspectiva.

### Combinando Transforms

Você pode combinar múltiplas transformações:
- `transform: translateX(20px) rotate(45deg) scale(1.2);`

**Pense assim:** É como fazer várias operações em sequência: mover, depois girar, depois aumentar.

### Quando Usar Transforms?

- Para criar efeitos de hover interessantes
- Para animações de movimento e rotação
- Para criar elementos que "flutuam" ou se movem
- Para efeitos visuais sem afetar o layout
- Combinado com transitions para animações suaves

### Vantagens dos Transforms

- Não afetam o layout (outros elementos não se movem)
- Performance melhor que mudar position
- Permitem criar efeitos visuais interessantes
- Podem ser animados suavemente com transitions

**Resumo:** Transforms permitem modificar elementos visualmente sem afetar o layout. São perfeitos para criar efeitos visuais e animações que não empurram outros elementos.

---

## 🎯 Resumo dos Conceitos Principais

### Sistemas de Layout:

- **Flow Layout**: Comportamento padrão, elementos seguem a ordem do HTML
- **Flexbox**: Layout unidimensional, perfeito para centralização e distribuição de espaço
- **CSS Grid**: Layout bidimensional, perfeito para layouts complexos
- **Multicolumn**: Divide conteúdo em colunas, como jornal
- **Float**: Técnica antiga, útil apenas para texto ao redor de imagens

### Animações e Efeitos:

- **Transitions**: Mudanças suaves entre estados
- **Keyframe Animations**: Animações complexas com múltiplas etapas
- **Transforms**: Modificações visuais sem afetar o layout

### Quando Usar Cada Sistema:

- **Layout simples**: Flow Layout
- **Centralizar, distribuir espaço**: Flexbox
- **Layout complexo bidimensional**: CSS Grid
- **Texto em colunas**: Multicolumn
- **Animações simples**: Transitions
- **Animações complexas**: Keyframe Animations
- **Efeitos visuais**: Transforms

---

## 🚀 Próximos Passos

Agora que você entendeu os sistemas de layout e animações, você está pronto para:
- Criar layouts modernos e responsivos
- Centralizar elementos facilmente
- Criar animações suaves e profissionais
- Combinar diferentes sistemas de layout
- Criar interfaces interativas e polidas

Na próxima aula, você aprenderá mais sobre responsividade, media queries e como fazer seus layouts funcionarem em diferentes dispositivos.


