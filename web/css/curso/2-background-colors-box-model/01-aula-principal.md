# Aula 2: Background, Colors, Box Model e Fundamentos de Layout

## 🎨 Cores em CSS

### Por que Cores são Importantes?

As cores são fundamentais para criar interfaces visuais atraentes e comunicar informações. Em CSS, existem várias formas de definir cores, cada uma com suas vantagens e casos de uso específicos. Entender essas diferentes formas permite que você escolha a melhor opção para cada situação.

### O que são Cores Nomeadas (Named Colors)?

**Cores nomeadas** são palavras em inglês que representam cores específicas. São a forma mais simples e intuitiva de definir cores, especialmente para iniciantes.

#### Características:
- Fáceis de lembrar e usar
- Limitadas a um conjunto pré-definido de cores
- Não permitem variações ou transparência
- Úteis para cores básicas e prototipagem rápida

#### Exemplos de Cores Nomeadas:
- Cores básicas: `red`, `blue`, `green`, `yellow`, `black`, `white`
- Cores intermediárias: `orange`, `purple`, `pink`, `gray`
- Tons específicos: `crimson`, `navy`, `lime`, `teal`

#### Quando Usar:
- Prototipagem rápida
- Cores básicas que não precisam de precisão
- Quando você não precisa de transparência ou variações específicas

#### Limitações:
- Não há controle sobre transparência
- Conjunto limitado de cores disponíveis
- Diferentes navegadores podem interpretar algumas cores de forma ligeiramente diferente

---

### O que é RGB?

**RGB** significa **Red, Green, Blue** (Vermelho, Verde, Azul). É um modelo de cores baseado na luz, onde qualquer cor é criada combinando diferentes intensidades dessas três cores primárias.

#### Como Funciona:
- Cada cor primária pode ter valores de **0 a 255**
- `rgb(255, 0, 0)` = vermelho puro (máximo de vermelho, zero de verde e azul)
- `rgb(0, 255, 0)` = verde puro
- `rgb(0, 0, 255)` = azul puro
- `rgb(0, 0, 0)` = preto (ausência de todas as cores)
- `rgb(255, 255, 255)` = branco (máxima intensidade de todas as cores)

#### Vantagens:
- Controle preciso sobre cada componente de cor
- Permite criar milhões de cores diferentes
- Baseado em um sistema amplamente compreendido

#### Quando Usar:
- Quando você precisa de uma cor específica que não existe como cor nomeada
- Para criar variações de cores
- Quando trabalha com ferramentas de design que usam RGB

---

### O que é RGBA?

**RGBA** é uma extensão do RGB que adiciona um quarto valor: **Alpha** (transparência). O "A" em RGBA representa o canal de transparência.

#### Como Funciona:
- Os três primeiros valores são RGB (0-255)
- O quarto valor é a opacidade (0.0 a 1.0)
- `rgba(255, 0, 0, 1.0)` = vermelho totalmente opaco
- `rgba(255, 0, 0, 0.5)` = vermelho 50% transparente
- `rgba(255, 0, 0, 0)` = vermelho completamente transparente (invisível)

#### Vantagens:
- Permite criar cores com transparência
- Útil para sobreposições, overlays e efeitos visuais
- Mantém a cor mesmo quando transparente

#### Quando Usar:
- Quando você precisa de transparência em uma cor específica
- Para criar efeitos de sobreposição
- Quando quer que o fundo apareça através do elemento

#### Diferença entre RGBA e Opacity:
- `rgba()` afeta apenas a cor específica (fundo, texto, borda)
- `opacity` afeta o elemento inteiro, incluindo todo o conteúdo e elementos filhos

---

### O que é Hexadecimal (HEX)?

**Hexadecimal** é um sistema numérico de base 16 usado para representar cores de forma compacta. É amplamente usado em CSS e design web.

#### Como Funciona:
- Usa números de 0-9 e letras A-F
- Formato: `#RRGGBB` (6 dígitos)
- Cada par de dígitos representa uma cor primária (Red, Green, Blue)
- Valores vão de `00` (0) a `FF` (255)

#### Exemplos:
- `#FF0000` = vermelho (FF = 255 de vermelho, 00 de verde, 00 de azul)
- `#00FF00` = verde
- `#0000FF` = azul
- `#000000` = preto
- `#FFFFFF` = branco
- `#808080` = cinza médio

#### Formato Curto:
- `#F00` é equivalente a `#FF0000` (vermelho)
- `#0F0` é equivalente a `#00FF00` (verde)
- Funciona quando cada par de dígitos é igual

#### Vantagens:
- Formato compacto e eficiente
- Amplamente usado e reconhecido
- Fácil de copiar de ferramentas de design
- Padrão da indústria

#### Quando Usar:
- Formato mais comum em projetos profissionais
- Quando você copia cores de ferramentas de design (Figma, Photoshop, etc.)
- Para manter consistência com equipes de design

---

### O que é HSL?

**HSL** significa **Hue, Saturation, Lightness** (Matiz, Saturação, Luminosidade). É um modelo de cores baseado em como os humanos percebem as cores.

#### Componentes:
1. **Hue (Matiz)**: A cor em si, medida em graus (0-360)
   - 0° = vermelho
   - 120° = verde
   - 240° = azul
2. **Saturation (Saturação)**: Intensidade da cor (0% a 100%)
   - 0% = cinza (sem cor)
   - 100% = cor totalmente saturada (viva)
3. **Lightness (Luminosidade)**: Quanto de luz (0% a 100%)
   - 0% = preto
   - 50% = cor pura
   - 100% = branco

#### Exemplos:
- `hsl(0, 100%, 50%)` = vermelho puro
- `hsl(120, 100%, 50%)` = verde puro
- `hsl(0, 0%, 50%)` = cinza médio (sem saturação)
- `hsl(0, 100%, 25%)` = vermelho escuro
- `hsl(0, 100%, 75%)` = vermelho claro

#### Vantagens:
- Mais intuitivo para ajustar cores
- Fácil criar variações (mais claro/escuro, mais/menos saturado)
- Baseado em como vemos cores, não em valores técnicos

#### Quando Usar:
- Quando você precisa ajustar o brilho ou saturação de uma cor
- Para criar paletas de cores consistentes
- Quando trabalha com temas (claro/escuro)

---

### O que é HSLA?

**HSLA** adiciona transparência ao HSL, assim como RGBA faz com RGB.

#### Como Funciona:
- Os três primeiros valores são HSL
- O quarto valor é a opacidade (0.0 a 1.0)
- `hsla(0, 100%, 50%, 1.0)` = vermelho opaco
- `hsla(0, 100%, 50%, 0.5)` = vermelho 50% transparente

#### Vantagens:
- Combina a intuitividade do HSL com controle de transparência
- Útil para criar variações transparentes de cores

#### Quando Usar:
- Quando você precisa de transparência e quer ajustar facilmente a cor usando HSL

---

## 🖼️ Background (Fundo)

### O que é Background?

O **background** (fundo) em CSS refere-se a todas as propriedades que controlam a aparência visual da área de fundo de um elemento. Isso inclui cor de fundo, imagens, posicionamento, repetição e tamanho.

### Por que Background é Importante?

O fundo é uma das primeiras coisas que os usuários notam em uma página. Ele estabelece o tom visual, cria hierarquia e pode tornar o conteúdo mais legível ou atraente. Dominar as propriedades de background permite criar designs profissionais e visualmente interessantes.

---

### Background Color (Cor de Fundo)

**Background color** define a cor de fundo de um elemento. É a propriedade de background mais básica e mais usada.

#### Propriedade: `background-color`

#### Valores:
- Qualquer formato de cor: nomeada, RGB, RGBA, HEX, HSL, HSLA
- `transparent` - torna o fundo transparente

#### Características:
- A cor de fundo aparece atrás do conteúdo e padding
- Não afeta o conteúdo do elemento
- Pode ser combinada com imagens de fundo

#### Quando Usar:
- Para destacar seções da página
- Criar contraste para melhorar legibilidade
- Estabelecer identidade visual
- Separar visualmente diferentes áreas do layout

---

### Background Image (Imagem de Fundo)

**Background image** permite usar uma imagem como fundo de um elemento, em vez de apenas uma cor.

#### Propriedade: `background-image`

#### Valores:
- `url('caminho/para/imagem.jpg')` - caminho para a imagem
- `none` - remove imagem de fundo

#### Características:
- A imagem aparece atrás do conteúdo
- Pode ser combinada com cor de fundo (a cor aparece se a imagem não carregar ou for transparente)
- Por padrão, a imagem se repete para preencher o elemento

#### Quando Usar:
- Para adicionar texturas ou padrões
- Criar fundos visuais interessantes
- Adicionar imagens decorativas sem afetar o HTML
- Criar hero sections ou banners

#### Considerações:
- Imagens grandes podem afetar o desempenho
- Sempre tenha uma cor de fundo de fallback
- Considere a legibilidade do conteúdo sobre a imagem

---

### Background Position (Posicionamento do Fundo)

**Background position** controla onde a imagem de fundo é posicionada dentro do elemento.

#### Propriedade: `background-position`

#### Valores:
- Palavras-chave: `top`, `bottom`, `left`, `right`, `center`
- Combinações: `top left`, `center center`, `bottom right`
- Valores específicos: `50% 50%`, `10px 20px`
- Valores mistos: `center 20px`, `50% top`

#### Como Funciona:
- O primeiro valor controla a posição horizontal
- O segundo valor controla a posição vertical
- `center center` (ou apenas `center`) centraliza a imagem
- `top left` posiciona no canto superior esquerdo

#### Quando Usar:
- Para posicionar logos ou ícones específicos
- Criar efeitos de parallax ou movimento
- Controlar como imagens grandes são exibidas
- Alinhar padrões repetidos

---

### Background Repeat (Repetição do Fundo)

**Background repeat** controla se e como a imagem de fundo se repete para preencher o elemento.

#### Propriedade: `background-repeat`

#### Valores:
- `repeat` - repete em ambas as direções (padrão)
- `repeat-x` - repete apenas horizontalmente
- `repeat-y` - repete apenas verticalmente
- `no-repeat` - não repete (imagem aparece uma vez)
- `space` - repete com espaços entre as imagens
- `round` - repete e redimensiona para caber sem cortes

#### Quando Usar:
- `no-repeat` - para imagens grandes ou únicas
- `repeat` - para padrões e texturas
- `repeat-x` - para bordas horizontais
- `repeat-y` - para bordas verticais

---

### Background Size (Tamanho do Fundo)

**Background size** controla o tamanho da imagem de fundo.

#### Propriedade: `background-size`

#### Valores:
- `auto` - tamanho original da imagem
- `cover` - cobre todo o elemento, pode cortar partes da imagem
- `contain` - mostra a imagem inteira, pode deixar espaços vazios
- Valores específicos: `100px 200px`, `50% 80%`

#### Como Funciona:
- `cover` - a imagem preenche todo o espaço, mantendo proporção, pode cortar
- `contain` - a imagem inteira é visível, mantendo proporção, pode deixar espaços
- Valores específicos permitem controle total sobre largura e altura

#### Quando Usar:
- `cover` - para hero sections ou fundos que devem preencher completamente
- `contain` - quando você precisa ver a imagem inteira
- Valores específicos - para controle preciso do tamanho

---

### Background Shorthand (Atalho)

**Background shorthand** permite definir múltiplas propriedades de background em uma única declaração.

#### Propriedade: `background`

#### Ordem Recomendada:
```
background: color image position / size repeat attachment origin clip;
```

#### Exemplo:
```css
background: #333 url('imagem.jpg') center center / cover no-repeat;
```

#### Vantagens:
- Código mais conciso
- Define múltiplas propriedades de uma vez

#### Quando Usar:
- Para definir todas as propriedades de background de uma vez
- Quando você quer código mais limpo

---

## 📦 Box Model (Modelo de Caixa)

### O que é o Box Model?

O **Box Model** (Modelo de Caixa) é um conceito fundamental em CSS que descreve como cada elemento HTML é estruturado como uma caixa retangular. Entender o box model é essencial para controlar o tamanho, espaçamento e posicionamento dos elementos.

### Por que o Box Model é Importante?

O box model determina como os elementos ocupam espaço na página. Sem entender o box model, é impossível criar layouts precisos e controlar adequadamente o espaçamento entre elementos. É a base de todo layout CSS.

### Componentes do Box Model

Cada elemento é composto por quatro áreas, de dentro para fora:

1. **Content (Conteúdo)** - O conteúdo real do elemento (texto, imagens, etc.)
2. **Padding (Preenchimento)** - Espaço entre o conteúdo e a borda
3. **Border (Borda)** - Linha ao redor do padding e conteúdo
4. **Margin (Margem)** - Espaço fora da borda, entre elementos

### Visualização do Box Model

```
┌─────────────────────────────────┐ ← Margin (fora)
│  ┌───────────────────────────┐  │
│  │ ┌────────────────────────┐ │  │ ← Border (borda)
│  │ │ ┌────────────────────┐ │ │  │
│  │ │ │                    │ │ │  │ ← Padding (preenchimento)
│  │ │ │     Content        │ │ │  │ ← Content (conteúdo)
│  │ │ │                    │ │ │  │
│  │ │ └────────────────────┘ │ │  │
│  │ └────────────────────────┘ │  │
│  └───────────────────────────┘  │
└─────────────────────────────────┘
```

### Box Sizing

A propriedade `box-sizing` controla como a largura e altura são calculadas.

#### Valores:
- `content-box` (padrão) - width/height incluem apenas o conteúdo
- `border-box` - width/height incluem conteúdo, padding e border

#### Por que é Importante?
- Com `content-box`, se você define `width: 200px`, o elemento pode ser maior que 200px devido ao padding e border
- Com `border-box`, `width: 200px` significa que o elemento inteiro (incluindo padding e border) terá 200px

#### Boa Prática:
Sempre use `box-sizing: border-box` para facilitar cálculos e evitar surpresas com tamanhos.

---

## 📏 Padding (Preenchimento)

### O que é Padding?

**Padding** é o espaço entre o conteúdo de um elemento e sua borda. É como uma "almofada" interna que cria espaço dentro do elemento.

### Por que Padding é Importante?

O padding melhora a legibilidade, cria respiração visual e ajuda a organizar o conteúdo. Sem padding adequado, o texto pode ficar colado nas bordas, tornando a leitura difícil e o design pouco profissional.

### Características do Padding:
- Cria espaço **dentro** do elemento
- Afeta a área clicável de elementos interativos
- Pode ser definido individualmente para cada lado
- Não pode ter valores negativos
- É parte do elemento (afeta a cor de fundo)

### Propriedades de Padding:

#### Padding Individual:
- `padding-top` - padding superior
- `padding-right` - padding direito
- `padding-bottom` - padding inferior
- `padding-left` - padding esquerdo

#### Padding Shorthand:
- `padding: 10px;` - aplica 10px em todos os lados
- `padding: 10px 20px;` - 10px em cima/baixo, 20px em esquerda/direita
- `padding: 10px 20px 30px;` - 10px em cima, 20px esquerda/direita, 30px embaixo
- `padding: 10px 20px 30px 40px;` - cima, direita, baixo, esquerda (sentido horário)

### Quando Usar Padding:
- Para criar espaço interno em botões e cards
- Melhorar legibilidade do texto
- Criar separação visual entre conteúdo e borda
- Adicionar área clicável em elementos interativos

---

## 📐 Margin (Margem)

### O que é Margin?

**Margin** é o espaço **fora** da borda de um elemento, criando distância entre o elemento e outros elementos ao redor.

### Por que Margin é Importante?

O margin controla o espaçamento entre elementos, criando hierarquia visual e organização no layout. É essencial para criar designs limpos e bem espaçados.

### Características do Margin:
- Cria espaço **fora** do elemento
- Não afeta a cor de fundo (é transparente)
- Pode ter valores negativos (permite sobreposição)
- Pode ser definido individualmente para cada lado
- Margens verticais podem colapsar (margin collapse)

### Propriedades de Margin:

#### Margin Individual:
- `margin-top` - margem superior
- `margin-right` - margem direito
- `margin-bottom` - margem inferior
- `margin-left` - margem esquerdo

#### Margin Shorthand:
- `margin: 10px;` - aplica 10px em todos os lados
- `margin: 10px 20px;` - 10px em cima/baixo, 20px em esquerda/direita
- `margin: 10px 20px 30px;` - 10px em cima, 20px esquerda/direita, 30px embaixo
- `margin: 10px 20px 30px 40px;` - cima, direita, baixo, esquerda (sentido horário)

#### Valores Especiais:
- `auto` - centraliza elementos horizontalmente (quando width está definido)
- Valores negativos - permitem sobreposição de elementos

### Margin Collapse (Colapso de Margem):

Quando dois elementos têm margens verticais adjacentes, elas colapsam (se combinam), não se somam. Apenas a maior margem é aplicada.

#### Exemplo:
- Elemento A tem `margin-bottom: 20px`
- Elemento B tem `margin-top: 30px`
- O espaço entre eles será 30px (não 50px)

### Quando Usar Margin:
- Para criar espaço entre elementos diferentes
- Separar seções da página
- Centralizar elementos horizontalmente
- Criar hierarquia visual através do espaçamento

---

## 📐 Width e Height (Largura e Altura)

### O que são Width e Height?

**Width** (largura) e **Height** (altura) definem o tamanho da área de conteúdo de um elemento. Eles controlam quantos pixels (ou outras unidades) o elemento ocupa horizontalmente e verticalmente.

### Por que Width e Height são Importantes?

Controlar o tamanho dos elementos é fundamental para criar layouts estruturados. Sem controle sobre dimensões, os elementos ocupam todo o espaço disponível ou apenas o necessário para o conteúdo, o que pode não ser o desejado.

### Propriedades:

#### Width (Largura):
- `width: 200px;` - largura fixa em pixels
- `width: 50%;` - largura relativa (50% do elemento pai)
- `width: auto;` - largura automática (padrão)
- `width: 100%;` - ocupa toda a largura disponível

#### Height (Altura):
- `height: 300px;` - altura fixa em pixels
- `height: 50%;` - altura relativa (50% do elemento pai)
- `height: auto;` - altura automática (padrão)
- `height: 100vh;` - altura da viewport (tela visível)

### Características:
- Por padrão, elementos têm `width: auto` e `height: auto`
- `width: auto` faz o elemento ocupar toda a largura disponível (elementos block)
- `height: auto` ajusta a altura ao conteúdo
- Valores percentuais são relativos ao elemento pai

### Quando Usar:
- Para controlar o tamanho de containers
- Criar layouts com larguras específicas
- Definir alturas para seções específicas
- Criar designs responsivos com unidades relativas

### Considerações:
- Altura fixa pode causar problemas se o conteúdo for maior
- Largura fixa pode quebrar em telas pequenas
- Prefira unidades relativas para responsividade

---

## 🔲 Border (Borda)

### O que é Border?

**Border** (borda) é a linha que circunda o conteúdo e padding de um elemento. Ela separa visualmente o elemento de outros elementos e pode ser usada para destacar ou organizar conteúdo.

### Por que Border é Importante?

Bordas ajudam a definir limites, criar separação visual, destacar elementos importantes e organizar o layout. São uma ferramenta poderosa para criar hierarquia visual.

### Características do Border:
- Aparece entre padding e margin
- Pode ter diferentes estilos, larguras e cores
- Afeta o tamanho total do elemento (a menos que use `box-sizing: border-box`)
- Pode ser definida individualmente para cada lado

### Propriedades de Border:

#### Border Individual por Lado:
- `border-top` - borda superior
- `border-right` - borda direita
- `border-bottom` - borda inferior
- `border-left` - borda esquerda

#### Propriedades Específicas:
- `border-width` - espessura da borda
- `border-style` - estilo da borda
- `border-color` - cor da borda

#### Border Shorthand:
```css
border: width style color;
```
Exemplo: `border: 2px solid black;`

### Border Style (Estilos de Borda):

- `solid` - linha sólida contínua
- `dashed` - linha tracejada
- `dotted` - linha pontilhada
- `double` - linha dupla
- `groove` - borda 3D com efeito de sulco
- `ridge` - borda 3D com efeito de relevo
- `inset` - borda 3D com efeito interno
- `outset` - borda 3D com efeito externo
- `none` - sem borda
- `hidden` - borda oculta (mas ocupa espaço)

### Border Width (Largura):
- Valores: `thin`, `medium`, `thick`, ou valores específicos como `2px`, `0.5em`
- Pode ser definida individualmente: `border-top-width`, `border-right-width`, etc.

### Border Color (Cor):
- Qualquer formato de cor: nomeada, RGB, HEX, HSL, etc.
- Pode ser definida individualmente: `border-top-color`, etc.

### Border Radius (Bordas Arredondadas):

A propriedade `border-radius` arredonda os cantos da borda.

#### Valores:
- `border-radius: 5px;` - arredonda todos os cantos
- `border-radius: 10px 20px;` - canto superior esquerdo/inferior direito, depois outros
- `border-radius: 10px 20px 30px 40px;` - cada canto individualmente

#### Quando Usar:
- Para criar designs modernos e suaves
- Botões e cards com cantos arredondados
- Criar formas circulares (`border-radius: 50%`)

### Quando Usar Border:
- Para destacar elementos importantes
- Criar separação visual entre seções
- Definir limites de formulários e inputs
- Criar cards e containers visuais

---

## 🔳 Outline (Contorno)

### O que é Outline?

**Outline** é uma linha desenhada ao redor de um elemento, **fora** da borda. Diferente da borda, o outline não afeta o tamanho ou posição do elemento no layout.

### Por que Outline é Importante?

O outline é principalmente usado para acessibilidade, especialmente para indicar quando um elemento está em foco (foco do teclado). É essencial para usuários que navegam com teclado.

### Características do Outline:
- Aparece **fora** da borda
- **Não afeta** o tamanho do elemento (não ocupa espaço no layout)
- Não pode ter cantos arredondados
- Pode sobrepor outros elementos
- Principalmente usado para indicar foco

### Propriedades de Outline:

#### Outline Shorthand:
```css
outline: width style color;
```
Exemplo: `outline: 2px solid blue;`

#### Propriedades Individuais:
- `outline-width` - largura do contorno
- `outline-style` - estilo (solid, dashed, dotted, etc.)
- `outline-color` - cor do contorno
- `outline-offset` - distância entre outline e borda

### Outline Style:
- Similar aos estilos de border: `solid`, `dashed`, `dotted`, `none`, etc.

### Quando Usar Outline:
- Para indicar foco em elementos interativos (acessibilidade)
- Destacar elementos temporariamente durante desenvolvimento
- Criar indicadores visuais sem afetar o layout

### ⚠️ Importante sobre Outline:
- **Nunca remova o outline** de elementos focáveis sem fornecer uma alternativa
- Outline é crucial para acessibilidade
- Se remover, sempre adicione outro indicador visual de foco

---

## 🌑 Box Shadow (Sombra de Caixa)

### O que é Box Shadow?

**Box shadow** cria uma sombra ao redor do elemento, adicionando profundidade e dimensão visual. É uma das propriedades mais usadas para criar designs modernos e interfaces com profundidade.

### Por que Box Shadow é Importante?

Sombras ajudam a criar hierarquia visual, destacar elementos, criar sensação de profundidade e tornar interfaces mais modernas e profissionais. É uma ferramenta poderosa para design.

### Características do Box Shadow:
- Não afeta o tamanho do elemento
- Pode ter múltiplas sombras
- Pode criar efeitos de elevação e profundidade
- Muito usado em cards, botões e modais

### Propriedade: `box-shadow`

#### Sintaxe:
```css
box-shadow: offset-x offset-y blur-radius spread-radius color;
```

#### Componentes:
1. **offset-x** - deslocamento horizontal da sombra (positivo = direita, negativo = esquerda)
2. **offset-y** - deslocamento vertical da sombra (positivo = baixo, negativo = cima)
3. **blur-radius** - quanto a sombra é desfocada (0 = sombra nítida, valores maiores = mais desfocada)
4. **spread-radius** - quanto a sombra se espalha (opcional, padrão 0)
5. **color** - cor da sombra (pode incluir transparência)

#### Exemplos:
- `box-shadow: 2px 2px 4px rgba(0,0,0,0.2);` - sombra simples e suave
- `box-shadow: 0 4px 6px rgba(0,0,0,0.1);` - sombra abaixo do elemento
- `box-shadow: 0 0 10px rgba(0,0,0,0.5);` - sombra desfocada ao redor
- `box-shadow: inset 0 2px 4px rgba(0,0,0,0.1);` - sombra interna

#### Inset (Sombra Interna):
- Adiciona `inset` no início para criar sombra dentro do elemento
- Útil para criar efeitos de profundidade interna

#### Múltiplas Sombras:
Você pode adicionar múltiplas sombras separadas por vírgula:
```css
box-shadow: 
  0 2px 4px rgba(0,0,0,0.1),
  0 4px 8px rgba(0,0,0,0.05);
```

### Quando Usar Box Shadow:
- Para destacar cards e containers
- Criar efeitos de elevação (elementos "flutuando")
- Adicionar profundidade a botões
- Criar modais e popups com destaque
- Melhorar a hierarquia visual

### Boas Práticas:
- Use sombras sutis para não sobrecarregar o design
- Considere a direção da luz (geralmente de cima)
- Use transparência (rgba) para sombras mais naturais
- Evite sombras muito escuras ou muito grandes

---

## 📏 Unidades CSS

### O que são Unidades CSS?

**Unidades CSS** são valores que definem tamanhos, distâncias e medidas em CSS. Elas especificam como o navegador deve interpretar valores numéricos para propriedades como width, height, margin, padding, font-size, etc.

### Por que Unidades são Importantes?

Escolher a unidade correta é crucial para criar layouts que funcionem bem em diferentes dispositivos e tamanhos de tela. Unidades diferentes se comportam de formas diferentes, e entender essas diferenças permite criar designs responsivos e acessíveis.

### Categorias de Unidades:

#### 1. Unidades Absolutas
Tamanhos fixos que não mudam independente do contexto.

#### 2. Unidades Relativas
Tamanhos que são calculados baseados em outros valores (elemento pai, viewport, etc.).

---

## 📐 Unidades Absolutas

### O que são Unidades Absolutas?

**Unidades absolutas** representam medidas fixas que sempre terão o mesmo tamanho físico, independente do dispositivo, tamanho de tela ou outros fatores.

### Características:
- Tamanho fixo e consistente
- Não se adaptam ao contexto
- Úteis para impressão e medidas precisas
- Podem causar problemas em telas pequenas

### Unidades Absolutas Comuns:

#### Pixels (px)
- **Mais comum** em desenvolvimento web
- `1px` = 1 pixel na tela
- Tamanho fixo, não escala
- **Quando usar**: Para elementos que precisam de tamanho preciso, bordas, sombras

#### Pontos (pt)
- Usado principalmente para impressão
- `1pt` = 1/72 de polegada
- **Quando usar**: Estilos de impressão (media queries para print)

#### Centímetros (cm) e Milímetros (mm)
- Unidades físicas reais
- **Quando usar**: Estilos de impressão, quando você precisa de medidas físicas precisas

#### Polegadas (in)
- `1in` = 2.54cm = 96px
- **Quando usar**: Estilos de impressão

### Vantagens das Unidades Absolutas:
- Precisão e consistência
- Controle total sobre tamanhos
- Previsibilidade

### Desvantagens:
- Não se adaptam a diferentes tamanhos de tela
- Podem causar problemas em dispositivos móveis
- Não são ideais para responsividade

---

## 📐 Unidades Relativas

### O que são Unidades Relativas?

**Unidades relativas** são calculadas baseadas em outros valores, como o tamanho da fonte do elemento pai, o tamanho da viewport, ou o tamanho do elemento raiz. Elas se adaptam ao contexto.

### Características:
- Se adaptam ao contexto
- Melhor para responsividade
- Mais flexíveis e acessíveis
- Podem ser mais complexas de entender inicialmente

### Unidades Relativas Comuns:

#### Em (em)
- Relativo ao **tamanho da fonte do elemento pai**
- `1em` = tamanho da fonte do elemento pai
- Se o pai tem `font-size: 16px`, então `1em = 16px`
- **Quando usar**: Para criar escalas proporcionais baseadas no contexto do elemento

#### Rem (rem)
- Relativo ao **tamanho da fonte do elemento raiz** (geralmente `<html>`)
- `1rem` = tamanho da fonte do elemento raiz
- Se o root tem `font-size: 16px`, então `1rem = 16px`
- **Quando usar**: Para criar escalas consistentes em toda a página (mais previsível que em)

#### Porcentagem (%)
- Relativo ao **elemento pai**
- `50%` = metade do tamanho do elemento pai
- **Quando usar**: Para criar layouts flexíveis e responsivos

#### Viewport Width (vw)
- Relativo à **largura da viewport** (janela do navegador)
- `1vw` = 1% da largura da viewport
- `100vw` = largura total da tela
- **Quando usar**: Para criar elementos que se adaptam à largura da tela

#### Viewport Height (vh)
- Relativo à **altura da viewport**
- `1vh` = 1% da altura da viewport
- `100vh` = altura total da tela
- **Quando usar**: Para criar elementos que ocupam a altura total da tela

#### Viewport Minimum (vmin)
- Relativo ao **menor lado** da viewport (largura ou altura, o que for menor)
- `1vmin` = 1% do menor lado
- **Quando usar**: Para garantir que elementos se adaptem ao menor lado da tela

#### Viewport Maximum (vmax)
- Relativo ao **maior lado** da viewport
- `1vmax` = 1% do maior lado
- **Quando usar**: Para elementos que devem se adaptar ao maior lado

### Vantagens das Unidades Relativas:
- Responsividade automática
- Melhor acessibilidade (respeitam preferências do usuário)
- Escalabilidade
- Mais flexíveis

### Desvantagens:
- Podem ser mais difíceis de prever
- Podem criar efeitos cascata inesperados (especialmente com em)
- Requerem mais entendimento do contexto

---

## 🔄 Unidades Absolutas vs. Relativas

### Quando Usar Cada Uma?

#### Use Unidades Absolutas (px) quando:
- Você precisa de precisão absoluta
- Trabalhando com bordas, sombras, outlines
- Elementos que não devem escalar
- Quando o tamanho exato é crítico

#### Use Unidades Relativas (rem, em, %, vw, vh) quando:
- Criando layouts responsivos
- Trabalhando com tipografia
- Quer que elementos se adaptem ao contexto
- Criando designs acessíveis
- Trabalhando com espaçamento que deve escalar

### Boa Prática Moderna:
- Use **rem** para font-size, margin, padding (consistência)
- Use **px** para bordas, sombras, outlines (precisão)
- Use **%** ou **vw/vh** para layouts e containers (responsividade)
- Use **em** quando você precisa de escalas proporcionais ao contexto local

---

## 🧮 Unidades com Funções

### O que são Funções CSS?

**Funções CSS** permitem calcular valores dinamicamente usando operações matemáticas. Elas tornam os designs mais flexíveis e adaptáveis.

### Por que Funções são Importantes?

Funções permitem criar layouts que se adaptam automaticamente a diferentes contextos, evitando valores fixos e criando designs mais inteligentes e responsivos.

### Funções Principais:

#### Calc() - Cálculo Matemático

**`calc()`** permite realizar cálculos matemáticos diretamente no CSS.

##### Sintaxe:
```css
width: calc(100% - 20px);
```

##### Operações Suportadas:
- Adição: `+`
- Subtração: `-`
- Multiplicação: `*`
- Divisão: `/`

##### Características:
- Pode misturar diferentes unidades
- Espaços ao redor dos operadores são obrigatórios
- Muito útil para layouts responsivos

##### Exemplos:
- `calc(100% - 40px)` - largura total menos 40px
- `calc(50% + 10px)` - metade mais 10px
- `calc(100vw - 100px)` - largura da tela menos 100px
- `calc((100% - 20px) / 2)` - cálculo complexo

##### Quando Usar:
- Para criar layouts com espaçamento fixo e conteúdo flexível
- Quando você precisa subtrair padding/margin de width
- Para criar grids e layouts complexos
- Quando valores absolutos e relativos precisam ser combinados

---

#### Min() - Valor Mínimo

**`min()`** retorna o menor valor entre os fornecidos.

##### Sintaxe:
```css
width: min(100%, 500px);
```

##### Como Funciona:
- Compara os valores e usa o menor
- `min(100%, 500px)` = usa 100% se for menor que 500px, senão usa 500px

##### Exemplos:
- `min(100%, 1200px)` - nunca maior que 1200px
- `min(50vw, 400px)` - menor entre metade da tela e 400px
- `font-size: min(5vw, 24px)` - fonte responsiva com limite máximo

##### Quando Usar:
- Para criar elementos que não ultrapassem um tamanho máximo
- Layouts responsivos com limites
- Tipografia responsiva com tamanhos máximos

---

#### Max() - Valor Máximo

**`max()`** retorna o maior valor entre os fornecidos.

##### Sintaxe:
```css
width: max(300px, 50%);
```

##### Como Funciona:
- Compara os valores e usa o maior
- `max(300px, 50%)` = usa 300px se for maior que 50%, senão usa 50%

##### Exemplos:
- `max(300px, 50%)` - nunca menor que 300px
- `min-height: max(100vh, 600px)` - altura mínima garantida
- `font-size: max(16px, 1.2rem)` - tamanho mínimo de fonte

##### Quando Usar:
- Para garantir tamanhos mínimos
- Elementos que devem ter um tamanho mínimo independente da tela
- Acessibilidade (tamanhos mínimos legíveis)

---

#### Clamp() - Valor Limitado

**`clamp()`** define um valor que está entre um mínimo e máximo, com um valor preferencial.

##### Sintaxe:
```css
font-size: clamp(min, preferred, max);
```

##### Como Funciona:
- Se o valor preferencial for menor que o mínimo, usa o mínimo
- Se o valor preferencial for maior que o máximo, usa o máximo
- Caso contrário, usa o valor preferencial

##### Exemplos:
- `clamp(16px, 5vw, 24px)` - entre 16px e 24px, preferindo 5vw
- `width: clamp(300px, 50%, 800px)` - largura entre 300px e 800px
- `font-size: clamp(1rem, 2.5vw, 2rem)` - tipografia fluida

##### Quando Usar:
- **Tipografia fluida** (fluid typography) - texto que escala suavemente
- Elementos que devem ter limites mínimo e máximo
- Designs totalmente responsivos
- Uma das funções mais úteis para responsividade moderna

---

### Vantagens das Funções:
- Código mais inteligente e adaptável
- Menos media queries necessárias
- Layouts mais flexíveis
- Melhor responsividade

### Quando Usar Funções:
- **calc()** - quando você precisa fazer cálculos
- **min()** - para estabelecer limites máximos
- **max()** - para estabelecer limites mínimos
- **clamp()** - para tipografia fluida e valores com limites

---

## 🎭 Display (Exibição)

### O que é Display?

A propriedade **`display`** é uma das mais importantes do CSS. Ela controla como um elemento é renderizado na página e como ele interage com outros elementos ao redor.

### Por que Display é Importante?

O display determina o comportamento fundamental de um elemento: se ele ocupa toda a largura disponível, se fica na mesma linha que outros elementos, se cria um novo contexto de formatação, etc. Sem entender display, é impossível criar layouts complexos.

### Valores Principais de Display:

#### Block (Bloco)

**Características:**
- Ocupa **toda a largura disponível**
- Começa em uma **nova linha**
- Pode ter width e height definidos
- Margin e padding funcionam em todos os lados
- Elementos block empilham verticalmente

**Elementos Block por Padrão:**
- `<div>`, `<p>`, `<h1>` até `<h6>`, `<section>`, `<article>`, `<header>`, `<footer>`, etc.

**Quando Usar:**
- Para containers e seções
- Quando você quer que o elemento ocupe toda a largura
- Para criar estrutura de layout

**Exemplo Visual:**
```
[============ Elemento Block ============]
[============ Outro Elemento Block ============]
```

---

#### Inline (Em Linha)

**Características:**
- Ocupa apenas o **espaço necessário** para o conteúdo
- Fica na **mesma linha** que outros elementos inline
- **Não pode** ter width e height definidos
- Margin e padding funcionam apenas horizontalmente (esquerda/direita)
- Elementos inline ficam lado a lado

**Elementos Inline por Padrão:**
- `<span>`, `<a>`, `<strong>`, `<em>`, `<img>`, etc.

**Quando Usar:**
- Para estilizar partes de texto
- Quando você quer elementos na mesma linha
- Para elementos que não devem quebrar o fluxo

**Exemplo Visual:**
```
[Texto] [inline] [elementos] [na mesma] [linha]
```

---

#### Inline-Block (Em Linha com Características de Bloco)

**Características:**
- Fica na **mesma linha** que outros elementos (como inline)
- **Pode ter** width e height definidos (como block)
- Margin e padding funcionam em **todos os lados** (como block)
- Combina vantagens de inline e block

**Quando Usar:**
- Para criar elementos que ficam lado a lado mas têm tamanho controlado
- Botões e cards em linha
- Navegação horizontal
- Quando você precisa de elementos inline com dimensões

**Exemplo Visual:**
```
[Elemento 1] [Elemento 2] [Elemento 3]
(todos na mesma linha, mas com width/height)
```

---

#### None (Nenhum)

**Características:**
- Elemento é **completamente removido** do layout
- Não ocupa espaço
- Não é renderizado
- Diferente de `visibility: hidden` (que ocupa espaço)

**Quando Usar:**
- Para esconder elementos completamente
- Em menus que aparecem/desaparecem
- Para elementos condicionais

---

#### Flex (Flexbox)

**Características:**
- Cria um **container flexível**
- Permite alinhamento e distribuição de espaço
- Elementos filhos se tornam flex items
- Sistema de layout moderno e poderoso

**Quando Usar:**
- Para criar layouts flexíveis
- Alinhar elementos facilmente
- Distribuir espaço entre elementos
- Criar navegações e menus

**Nota:** Flexbox é um tópico avançado que será coberto em aulas futuras.

---

#### Grid (Grade)

**Características:**
- Cria um **sistema de grade** bidimensional
- Permite layouts complexos em linhas e colunas
- Sistema de layout muito poderoso

**Quando Usar:**
- Para layouts complexos em duas dimensões
- Grids e estruturas de página
- Layouts profissionais

**Nota:** CSS Grid é um tópico avançado que será coberto em aulas futuras.

---

### Mudando o Display:

Você pode mudar o display de qualquer elemento:

```css
span {
  display: block; /* span normalmente é inline, agora é block */
}

div {
  display: inline; /* div normalmente é block, agora é inline */
}
```

### Por que Mudar o Display?

- Para controlar como elementos se comportam
- Para criar layouts específicos
- Para resolver problemas de alinhamento
- Para fazer elementos se comportarem de forma diferente do padrão

---

## 📚 Resumo dos Conceitos Principais

### Cores:
- **Nomeadas**: simples, limitadas
- **RGB/RGBA**: controle preciso, com/sem transparência
- **HEX**: formato compacto, padrão da indústria
- **HSL/HSLA**: intuitivo, fácil de ajustar

### Background:
- **background-color**: cor de fundo
- **background-image**: imagem de fundo
- **background-position**: onde a imagem aparece
- **background-repeat**: como a imagem se repete
- **background-size**: tamanho da imagem

### Box Model:
- **Content**: conteúdo do elemento
- **Padding**: espaço interno
- **Border**: borda ao redor
- **Margin**: espaço externo

### Espaçamento:
- **Padding**: espaço dentro (afeta área clicável)
- **Margin**: espaço fora (cria distância entre elementos)

### Dimensões:
- **Width/Height**: tamanho do conteúdo
- **Box-sizing**: como width/height são calculados

### Bordas e Contornos:
- **Border**: linha que afeta o tamanho
- **Outline**: linha que não afeta o tamanho (acessibilidade)
- **Border-radius**: cantos arredondados

### Efeitos:
- **Box-shadow**: sombras para profundidade

### Unidades:
- **Absolutas (px)**: tamanhos fixos
- **Relativas (rem, em, %, vw, vh)**: tamanhos adaptáveis

### Funções:
- **calc()**: cálculos matemáticos
- **min()**: valor mínimo
- **max()**: valor máximo
- **clamp()**: valor com limites

### Display:
- **block**: ocupa toda largura, nova linha
- **inline**: apenas espaço necessário, mesma linha
- **inline-block**: mesma linha mas com dimensões
- **none**: removido do layout

---

## 🎯 Próximos Passos

Agora que você entendeu os fundamentos de cores, backgrounds, box model e display, você está pronto para:
- Criar layouts básicos estruturados
- Controlar espaçamento e dimensões
- Aplicar cores e fundos de forma eficiente
- Entender como elementos ocupam espaço na página

Na próxima aula, você aprenderá sobre posicionamento, layouts mais complexos e como combinar esses conceitos para criar designs profissionais.




