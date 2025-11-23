# Aula 1: CSS Basics - Conteúdo Principal

## 📖 O que é CSS?

**CSS**, ou **Cascading Style Sheets** (Folhas de Estilo em Cascata), é a linguagem responsável por definir a **aparência visual** das páginas web. Enquanto o HTML estrutura o conteúdo, o CSS determina como esse conteúdo será apresentado visualmente.

### Função do CSS na Web

O CSS funciona como a **maquiagem** ou **roupa** de uma página web:
- Define **cores** e **tipografia**
- Controla **espaçamentos** e **posicionamento**
- Estabelece **layouts** e **responsividade**
- Permite criar **animações** e **transições**
- Separa a **estrutura** (HTML) da **apresentação** (CSS)

### Por que CSS é Importante?

Sem CSS, todas as páginas web seriam apenas texto preto em fundo branco, sem formatação. O CSS transforma conteúdo simples em experiências visuais atraentes e funcionais.

---

## 🎨 CSS Basics: Fundamentos Essenciais

### O que são CSS Basics?

CSS Basics compreendem os **blocos fundamentais** para estilizar páginas web. Isso inclui:

1. **Seletores** - Como identificar quais elementos HTML você quer estilizar
2. **Propriedades** - As características visuais que você quer alterar (como cor, tamanho de fonte)
3. **Valores** - As configurações específicas para essas propriedades (como "vermelho" ou "16px")

Dominar esses fundamentos permite controlar completamente a apresentação do conteúdo do seu site.

---

## 📝 Métodos de Aplicação de CSS

Existem três formas principais de aplicar CSS em suas páginas web. Cada uma tem seu propósito e momento de uso.

### 1. Inline CSS

**Inline CSS** envolve embutir estilos diretamente dentro dos elementos HTML usando o atributo `style`.

#### Características:
- Aplicado diretamente no elemento HTML
- Tem a **maior prioridade** (sobrescreve outros estilos)
- Útil para estilos **únicos** e **específicos** de um elemento
- **Não recomendado** para projetos maiores devido a problemas de manutenção

#### Quando Usar:
- Estilos que aparecem apenas uma vez
- Testes rápidos durante o desenvolvimento
- Situações onde você precisa sobrescrever estilos externos temporariamente

#### Exemplo Conceitual:
Imagine que você tem um parágrafo e quer que apenas ele tenha cor vermelha. Você pode fazer isso diretamente no elemento:

```html
<p style="color: red;">Este texto será vermelho</p>
```

#### Por que Evitar em Projetos Grandes?
- **Dificulta manutenção**: Se você quiser mudar a cor de todos os parágrafos vermelhos, terá que editar cada um
- **Mistura estrutura com apresentação**: HTML deveria focar em estrutura, não em aparência
- **Impossibilita reutilização**: Não pode aplicar o mesmo estilo em múltiplos elementos facilmente

---

### 2. Internal CSS (CSS Interno)

**Internal CSS** envolve escrever regras CSS diretamente dentro do documento HTML usando a tag `<style>`, geralmente colocada dentro do `<head>`.

#### Características:
- CSS fica no mesmo arquivo HTML
- Estilos aplicados apenas àquele documento específico
- Útil para estilos **específicos de uma página**
- Melhor organização que inline, mas ainda mistura estrutura com apresentação

#### Quando Usar:
- Páginas únicas com estilos específicos
- Protótipos rápidos
- Quando você não quer criar arquivos CSS separados

#### Exemplo Conceitual:
Dentro do `<head>` do seu HTML, você pode escrever:

```html
<head>
  <style>
    p {
      color: blue;
      font-size: 16px;
    }
  </style>
</head>
```

Isso fará com que todos os parágrafos da página fiquem azuis com tamanho de 16px.

#### Vantagens sobre Inline:
- Pode estilizar múltiplos elementos de uma vez
- Mais fácil de manter que inline
- Ainda mantém tudo em um arquivo

#### Desvantagens:
- Não pode reutilizar em outras páginas
- Mistura CSS com HTML no mesmo arquivo
- Pode tornar o arquivo HTML muito grande

---

### 3. External CSS (CSS Externo)

**External CSS** envolve escrever regras CSS em um arquivo separado (com extensão `.css`) e então vincular esse arquivo ao seu documento HTML.

#### Características:
- CSS em arquivo separado (ex: `estilos.css`)
- Vinculado ao HTML através da tag `<link>` no `<head>`
- Promove **código limpo** e **fácil manutenção**
- Permite **reutilização** de estilos em múltiplas páginas
- **Melhor prática** para projetos profissionais

#### Quando Usar:
- **Sempre que possível** em projetos reais
- Quando você tem múltiplas páginas que compartilham estilos
- Para manter código organizado e profissional

#### Exemplo Conceitual:

**Arquivo: estilos.css**
```css
p {
  color: green;
  font-size: 18px;
}
```

**Arquivo: index.html**
```html
<head>
  <link rel="stylesheet" href="estilos.css">
</head>
```

#### Vantagens:
- **Separação de responsabilidades**: HTML para estrutura, CSS para apresentação
- **Reutilização**: Um arquivo CSS pode estilizar múltiplas páginas
- **Manutenção fácil**: Mude o CSS uma vez, afeta todas as páginas
- **Performance**: Navegador pode fazer cache do arquivo CSS
- **Organização**: Código mais limpo e profissional

#### Por que é a Melhor Prática?
- Facilita trabalho em equipe (designer trabalha no CSS, desenvolvedor no HTML)
- Permite criar temas diferentes facilmente
- Código mais testável e debuggável
- Segue princípios de arquitetura de software

---

## 🔄 Cascading Order (Ordem de Cascata)

### O que é Cascata?

**Cascata** em CSS determina **quais estilos são aplicados** a um elemento quando múltiplas regras conflitantes o atingem. É um conjunto de regras que os navegadores seguem para resolver esses conflitos.

### Por que é Importante?

Imagine que você tem:
- Um arquivo CSS externo que diz: "todos os parágrafos são azuis"
- Um CSS interno que diz: "todos os parágrafos são vermelhos"
- Um estilo inline que diz: "este parágrafo é verde"

Qual cor será aplicada? A cascata resolve isso!

### Ordem de Prioridade (do menor para o maior):

1. **Estilos do navegador** (padrão)
2. **CSS Externo** (arquivo `.css` vinculado)
3. **CSS Interno** (tag `<style>`)
4. **CSS Inline** (atributo `style`)
5. **!important** (força uma regra, use com cuidado)

### Regra Geral:

Quando há conflito, o estilo com **maior especificidade** ou que foi **definido por último** (na mesma especificidade) vence.

### Exemplo Conceitual:

Se você tem:
- CSS externo: `p { color: blue; }`
- CSS interno: `p { color: red; }`
- Inline: `<p style="color: green;">Texto</p>`

O parágrafo será **verde** porque inline tem maior prioridade.

### Especificidade

Além da ordem, a **especificidade** também importa:
- Seletores mais específicos têm prioridade
- ID (`#id`) tem mais especificidade que classe (`.classe`)
- Classe tem mais especificidade que elemento (`p`)

---

## 📐 Estrutura de uma Regra CSS

### CSS Rules (Regras CSS)

**Regras CSS** são os blocos fundamentais de uma folha de estilo. Cada regra especifica como elementos HTML específicos devem ser estilizados.

### Componentes de uma Regra:

Uma regra CSS possui três partes principais:

```
seletor {
  propriedade: valor;
}
```

#### 1. Selector (Seletor)

O **seletor** identifica **quais elementos HTML** a regra deve estilizar. É como um "endereço" que aponta para os elementos.

Exemplos de seletores:
- `p` - seleciona todos os parágrafos
- `.destaque` - seleciona elementos com classe "destaque"
- `#cabecalho` - seleciona elemento com ID "cabecalho"

#### 2. Declaration Block (Bloco de Declaração)

O **bloco de declaração** é tudo que está entre as chaves `{ }`. Contém uma ou mais declarações.

#### 3. Declaration (Declaração)

Uma **declaração** é uma instrução individual que especifica uma propriedade e seu valor. Formato: `propriedade: valor;`

Cada declaração termina com ponto e vírgula (`;`).

### Exemplo Completo:

```css
p {
  color: blue;
  font-size: 16px;
  margin: 10px;
}
```

**Análise:**
- `p` = seletor (todos os parágrafos)
- `{ }` = bloco de declaração
- `color: blue;` = declaração (propriedade `color` com valor `blue`)
- `font-size: 16px;` = declaração (propriedade `font-size` com valor `16px`)
- `margin: 10px;` = declaração (propriedade `margin` com valor `10px`)

---

## 🎯 Seletores CSS

### O que são Seletores?

**Seletores** são padrões que identificam quais elementos HTML devem receber os estilos definidos. Eles são a forma de "falar" com o navegador: "aplique esses estilos a esses elementos".

### Tipos de Seletores

#### 1. Element Selectors (Seletores de Elemento)

Selecionam elementos diretamente pelo **nome da tag**.

**Características:**
- Mais básico e direto
- Aplica estilo a **todos** os elementos daquele tipo
- Sintaxe: apenas o nome da tag

**Exemplo:**
```css
h1 {
  color: red;
}
```
Isso estiliza **todos** os `<h1>` da página.

**Quando Usar:**
- Quando você quer estilizar todos os elementos de um tipo
- Para definir estilos base/padrão

---

#### 2. Class Selectors (Seletores de Classe)

Selecionam elementos baseados no **atributo `class`**.

**Características:**
- Permite estilizar múltiplos elementos que compartilham a mesma classe
- Mais flexível que seletores de elemento
- Sintaxe: ponto (`.`) seguido do nome da classe

**Exemplo:**
```css
.destaque {
  background-color: yellow;
}
```

Isso estiliza todos os elementos com `class="destaque"`.

**Quando Usar:**
- Quando você quer aplicar o mesmo estilo a múltiplos elementos diferentes
- Para criar estilos reutilizáveis

---

#### 3. ID Selectors (Seletores de ID)

Selecionam um **único elemento** baseado no atributo `id`.

**Características:**
- ID deve ser **único** na página (não repetir)
- Usado para estilizar um elemento específico
- Sintaxe: hash (`#`) seguido do ID

**Exemplo:**
```css
#cabecalho {
  background-color: blue;
}
```

Isso estiliza apenas o elemento com `id="cabecalho"`.

**Quando Usar:**
- Para elementos únicos na página (cabeçalho, rodapé, menu)
- Quando você precisa de alta especificidade

**⚠️ Importante:** IDs devem ser únicos. Não use o mesmo ID em múltiplos elementos.

---

#### 4. Universal Selector (Seletor Universal)

O seletor universal (`*`) seleciona **todos os elementos** da página.

**Características:**
- Representado por asterisco (`*`)
- Aplica estilo a **qualquer elemento**
- Usado frequentemente para "reset" de estilos

**Exemplo:**
```css
* {
  margin: 0;
  padding: 0;
}
```

Isso remove margens e paddings de **todos** os elementos.

**Quando Usar:**
- Para resetar estilos padrão do navegador
- Com muito cuidado, pois afeta tudo

---

#### 5. Grouping Selectors (Agrupamento de Seletores)

Permite aplicar os **mesmos estilos** a múltiplos seletores diferentes.

**Características:**
- Lista seletores separados por vírgula
- Aplica o mesmo estilo a todos
- Torna código mais conciso

**Exemplo:**
```css
h1, h2, h3 {
  color: blue;
  font-weight: bold;
}
```

Isso estiliza todos os `h1`, `h2` e `h3` da mesma forma.

**Quando Usar:**
- Quando múltiplos elementos precisam dos mesmos estilos
- Para evitar repetição de código

---

### Combinator Selectors (Seletores Combinadores)

Combinadores definem **relacionamentos** entre elementos baseados na estrutura do HTML.

#### 1. Descendant Combinator (Combinador Descendente)

Seleciona elementos que são **descendentes** (filhos, netos, etc.) de outro elemento.

**Sintaxe:** espaço (` `) entre seletores

**Exemplo:**
```css
div p {
  color: red;
}
```

Isso seleciona todos os `<p>` que estão **dentro** de um `<div>` (em qualquer nível).

**Quando Usar:**
- Para estilizar elementos dentro de containers específicos
- Quando você quer atingir elementos aninhados

---

#### 2. Child Combinator (Combinador Filho)

Seleciona elementos que são **filhos diretos** de outro elemento.

**Sintaxe:** sinal de maior (`>`)

**Exemplo:**
```css
div > p {
  color: blue;
}
```

Isso seleciona apenas `<p>` que são **filhos diretos** de `<div>`, ignorando netos.

**Diferença do Descendente:**
- Descendente: qualquer nível dentro
- Filho: apenas nível imediato

**Quando Usar:**
- Quando você precisa de precisão (apenas filhos diretos)
- Para evitar estilizar elementos muito aninhados

---

#### 3. Next Sibling Combinator (Combinador Irmão Adjacente)

Seleciona elemento que é o **próximo irmão imediato** de outro elemento.

**Sintaxe:** sinal de mais (`+`)

**Exemplo:**
```css
h1 + p {
  margin-top: 0;
}
```

Isso seleciona o `<p>` que vem **imediatamente depois** de um `<h1>`.

**Quando Usar:**
- Para estilizar elementos que aparecem logo após outros
- Para remover espaçamento entre elementos específicos

---

#### 4. Subsequent Sibling Combinator (Combinador Irmão Geral)

Seleciona **todos os irmãos** que vêm depois de um elemento.

**Sintaxe:** til (`~`)

**Exemplo:**
```css
h1 ~ p {
  color: green;
}
```

Isso seleciona **todos** os `<p>` que são irmãos de `<h1>` e aparecem depois dele.

**Diferença do Adjacente:**
- Adjacente (`+`): apenas o próximo irmão
- Geral (`~`): todos os irmãos seguintes

**Quando Usar:**
- Para estilizar múltiplos irmãos de uma vez
- Quando você não sabe quantos irmãos existem

---

#### 5. Attribute Selectors (Seletores de Atributo)

Selecionam elementos baseados na **presença ou valor** de atributos HTML.

**Sintaxe básica:** `[atributo]` ou `[atributo="valor"]`

**Exemplos:**
```css
/* Elementos com atributo href */
a[href] {
  color: blue;
}

/* Links que apontam para URLs específicas */
a[href="https://exemplo.com"] {
  color: green;
}

/* Elementos com atributo class contendo "destaque" */
[class*="destaque"] {
  background: yellow;
}
```

**Quando Usar:**
- Para estilizar elementos baseados em atributos específicos
- Para criar estilos mais precisos e condicionais

---

## 💬 Comentários em CSS

### O que são Comentários?

**Comentários** são notas que você adiciona ao código para explicar o que ele faz ou para desabilitar temporariamente partes do código. Navegadores **ignoram** comentários, então eles não afetam a aparência do site.

### Sintaxe:

```css
/* Este é um comentário */
```

Comentários começam com `/*` e terminam com `*/`.

### Exemplos de Uso:

```css
/* Estilos para o cabeçalho */
h1 {
  color: blue;
}

/* 
  Comentário de múltiplas linhas
  útil para explicações longas
*/
p {
  font-size: 16px;
}
```

### Por que Usar Comentários?

- **Documentação**: Explica o propósito do código
- **Organização**: Divide seções do CSS
- **Debugging**: Permite desabilitar código temporariamente
- **Colaboração**: Ajuda outros desenvolvedores a entender seu código

---

## ✍️ Propriedades e Valores

### O que são Propriedades?

**Propriedades** são instruções que dizem ao navegador **como estilizar** um elemento HTML. Cada propriedade controla uma característica visual específica.

### O que são Valores?

**Valores** especificam a **configuração exata** que você quer aplicar à propriedade.

### Estrutura:

```
propriedade: valor;
```

### Exemplos:

- `color: red;` - propriedade `color` com valor `red` (cor do texto)
- `font-size: 16px;` - propriedade `font-size` com valor `16px` (tamanho da fonte)
- `margin: 10px;` - propriedade `margin` com valor `10px` (margem)

### Tipos de Valores Comuns:

- **Cores**: `red`, `#FF0000`, `rgb(255, 0, 0)`
- **Tamanhos**: `16px`, `2em`, `50%`
- **Palavras-chave**: `bold`, `center`, `none`
- **Números**: `1`, `0.5`, `100`

---

## 🎨 Estilização de Texto e Tipografia

### Font Family (Família de Fonte)

Define qual **fonte** será usada para o texto.

**Propriedade:** `font-family`

**Valores:** Nome da fonte (ex: `Arial`, `Times New Roman`, `Georgia`)

**Exemplo:**
```css
p {
  font-family: Arial, sans-serif;
}
```

**Fallbacks:** Você pode listar múltiplas fontes. Se a primeira não estiver disponível, o navegador tenta a próxima.

---

### Font Size (Tamanho da Fonte)

Controla o **tamanho** do texto.

**Propriedade:** `font-size`

**Valores:** Tamanhos como `16px`, `1.2em`, `100%`

**Exemplo:**
```css
p {
  font-size: 16px;
}
```

**Unidades Comuns:**
- `px` - pixels (tamanho fixo)
- `em` - relativo ao elemento pai
- `rem` - relativo ao elemento raiz
- `%` - porcentagem

---

### Font Style (Estilo da Fonte)

Controla se o texto é **normal**, **itálico** ou **oblíquo**.

**Propriedade:** `font-style`

**Valores:** `normal`, `italic`, `oblique`

**Exemplo:**
```css
em {
  font-style: italic;
}
```

---

### Font Weight (Peso da Fonte)

Controla a **espessura** do texto (negrito).

**Propriedade:** `font-weight`

**Valores:** `normal`, `bold`, `100`, `200`, `300`, `400`, `500`, `600`, `700`, `800`, `900`

**Exemplo:**
```css
strong {
  font-weight: bold;
}
```

---

### Font Variant (Variante da Fonte)

Controla variações especiais da fonte, como **small-caps** (maiúsculas pequenas).

**Propriedade:** `font-variant`

**Valores:** `normal`, `small-caps`

---

### Outras Propriedades de Fonte:

- **`line-height`** - Espaçamento entre linhas
- **`letter-spacing`** - Espaçamento entre letras
- **`word-spacing`** - Espaçamento entre palavras

---

## 🎨 Propriedades de Texto

### Color (Cor)

Define a **cor do texto**.

**Propriedade:** `color`

**Valores:** Nomes de cores (`red`, `blue`), hex (`#FF0000`), rgb (`rgb(255, 0, 0)`)

**Exemplo:**
```css
p {
  color: blue;
}
```

---

### Text Alignment (Alinhamento de Texto)

Controla como o texto é **posicionado horizontalmente** dentro do elemento.

**Propriedade:** `text-align`

**Valores:** `left`, `right`, `center`, `justify`

**Exemplo:**
```css
p {
  text-align: center;
}
```

**Quando Usar:**
- `left` - padrão, texto à esquerda
- `right` - texto à direita
- `center` - texto centralizado
- `justify` - texto justificado (alinhado nas duas bordas)

---

### Text Decoration (Decoração de Texto)

Adiciona **linhas decorativas** ao texto (sublinhado, riscado, etc.).

**Propriedade:** `text-decoration`

**Valores:** `none`, `underline`, `overline`, `line-through`

**Exemplo:**
```css
a {
  text-decoration: underline;
}
```

**Uso Comum:**
- Remover sublinhado de links: `text-decoration: none;`
- Adicionar sublinhado: `text-decoration: underline;`
- Riscar texto: `text-decoration: line-through;`

---

### Text Transform (Transformação de Texto)

Controla a **capitalização** do texto, independente de como está escrito no HTML.

**Propriedade:** `text-transform`

**Valores:** `none`, `uppercase`, `lowercase`, `capitalize`

**Exemplo:**
```css
h1 {
  text-transform: uppercase;
}
```

**Efeitos:**
- `uppercase` - TUDO EM MAIÚSCULAS
- `lowercase` - tudo em minúsculas
- `capitalize` - Primeira Letra De Cada Palavra
- `none` - mantém como está no HTML

---

### Direction (Direção)

Define a **direção do texto** (esquerda para direita ou direita para esquerda).

**Propriedade:** `direction`

**Valores:** `ltr` (left-to-right), `rtl` (right-to-left)

**Quando Usar:**
- `ltr` - padrão para português, inglês
- `rtl` - para árabe, hebraico

---

## 🎭 Opacity (Opacidade)

### O que é Opacidade?

**Opacidade** controla a **transparência** de um elemento. Determina quanto do fundo atrás do elemento é visível.

**Propriedade:** `opacity`

**Valores:** Número entre `0` e `1`

### Escala de Valores:

- `1` - Elemento **totalmente opaco** (não transparente, completamente visível)
- `0.5` - Elemento **semi-transparente** (50% visível, 50% transparente)
- `0` - Elemento **completamente transparente** (invisível)

### Exemplo:

```css
.imagem {
  opacity: 0.7;
}
```

Isso torna o elemento 70% opaco (30% transparente).

### Quando Usar:

- Criar efeitos de sobreposição
- Destaques sutis
- Transições suaves
- Elementos que devem ser menos proeminentes

### ⚠️ Diferença entre `opacity` e `rgba`:

- `opacity` afeta o **elemento inteiro** (incluindo conteúdo e filhos)
- `rgba` afeta apenas a **cor de fundo**, não o conteúdo

---

## 📚 Resumo dos Conceitos Principais

### Estrutura de uma Regra CSS:
```
seletor {
  propriedade: valor;
}
```

### Três Métodos de Aplicação:
1. **Inline** - no atributo `style` (menos recomendado)
2. **Internal** - na tag `<style>` (para páginas únicas)
3. **External** - em arquivo `.css` separado (melhor prática)

### Ordem de Prioridade (Cascata):
1. Navegador (padrão)
2. CSS Externo
3. CSS Interno
4. CSS Inline
5. !important

### Seletores Principais:
- **Elemento**: `p`, `h1`, `div`
- **Classe**: `.destaque`
- **ID**: `#cabecalho`
- **Universal**: `*`
- **Agrupamento**: `h1, h2, h3`

### Propriedades de Texto Essenciais:
- `color` - cor do texto
- `font-family` - fonte
- `font-size` - tamanho
- `font-weight` - negrito
- `text-align` - alinhamento
- `text-decoration` - decoração
- `text-transform` - capitalização

---

## 🎯 Próximos Passos

Agora que você entendeu os fundamentos do CSS, você está pronto para:
- Aplicar estilos básicos em suas páginas
- Entender como o CSS se relaciona com HTML
- Começar a criar páginas visualmente atraentes

Na próxima aula, você aprenderá sobre cores, backgrounds e o modelo de caixa (box model), que são fundamentais para criar layouts profissionais.

