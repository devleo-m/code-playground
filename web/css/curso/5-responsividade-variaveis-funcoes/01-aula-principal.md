# Aula 5: Responsividade, Variáveis e Funções em CSS

## 🎯 Introdução

Nesta aula, você aprenderá sobre as ferramentas essenciais para criar websites que funcionam perfeitamente em qualquer dispositivo, desde smartphones até monitores grandes. Além disso, descobrirá como usar variáveis CSS e funções para tornar seu código mais organizado, eficiente e fácil de manter.

---

## 📱 Media Queries

### O que são Media Queries?

**Media Queries** são uma funcionalidade poderosa do CSS que permite aplicar estilos diferentes baseados nas características do dispositivo ou tela que está visualizando a página. Elas são a base para criar designs **responsivos** - designs que se adaptam automaticamente a diferentes tamanhos de tela, resoluções e orientações.

### Por que Media Queries são Importantes?

Imagine que você criou um site que fica perfeito no seu computador, mas quando alguém acessa pelo celular, o texto fica muito pequeno, os botões ficam difíceis de clicar, e o layout fica quebrado. Media Queries resolvem exatamente esse problema, permitindo que você crie versões diferentes do seu design para diferentes dispositivos.

### Como Funcionam as Media Queries?

Media Queries funcionam como **condições** que o navegador verifica. Se a condição for verdadeira (por exemplo, "a tela tem menos de 768px de largura"), os estilos dentro daquela media query são aplicados. Se for falsa, eles são ignorados.

É como ter diferentes roupas para diferentes estações do ano. Você verifica a temperatura (a condição) e então escolhe a roupa apropriada (os estilos).

### Sintaxe Básica de Media Queries

A sintaxe básica de uma media query é:

```css
@media (condição) {
  /* estilos que serão aplicados se a condição for verdadeira */
}
```

A palavra `@media` indica que você está criando uma media query. Dentro dos parênteses, você define a condição. Dentro das chaves, você escreve os estilos que serão aplicados quando essa condição for verdadeira.

### Tipos de Condições em Media Queries

#### 1. Largura da Tela (Width)

A condição mais comum é verificar a largura da tela:

- **`max-width`**: Aplica estilos quando a tela tem **no máximo** uma certa largura (ou seja, quando a tela é menor ou igual)
- **`min-width`**: Aplica estilos quando a tela tem **no mínimo** uma certa largura (ou seja, quando a tela é maior ou igual)

**Exemplo conceitual:**
- `@media (max-width: 768px)` significa: "aplique esses estilos quando a tela tiver 768 pixels ou menos de largura"
- `@media (min-width: 769px)` significa: "aplique esses estilos quando a tela tiver 769 pixels ou mais de largura"

#### 2. Altura da Tela (Height)

Você também pode verificar a altura:

- **`max-height`**: Quando a altura é menor ou igual
- **`min-height`**: Quando a altura é maior ou igual

#### 3. Orientação (Orientation)

Verifica se o dispositivo está em modo retrato (portrait) ou paisagem (landscape):

- **`portrait`**: Dispositivo na vertical (mais alto que largo)
- **`landscape`**: Dispositivo na horizontal (mais largo que alto)

#### 4. Resolução (Resolution)

Verifica a resolução da tela (útil para telas de alta densidade como Retina):

- **`min-resolution`**: Resolução mínima
- **`max-resolution`**: Resolução máxima

### Breakpoints Comuns

**Breakpoints** são os pontos onde seu design muda. São valores de largura de tela que você usa como referência para aplicar estilos diferentes. Embora não existam breakpoints "oficiais", existem valores comuns baseados nos tamanhos típicos de dispositivos:

- **Mobile**: até 480px (smartphones pequenos)
- **Tablet**: 481px a 768px (tablets em modo retrato)
- **Desktop pequeno**: 769px a 1024px (laptops, tablets em modo paisagem)
- **Desktop médio**: 1025px a 1200px (desktops)
- **Desktop grande**: acima de 1200px (monitores grandes)

**Importante:** Esses são valores comuns, mas você deve escolher breakpoints baseados no seu design específico, não apenas seguir valores genéricos.

### Abordagens: Mobile-First vs Desktop-First

#### Mobile-First (Móvel Primeiro)

**Mobile-First** significa que você começa criando o design para dispositivos móveis (telas pequenas) e depois adiciona estilos para telas maiores usando `min-width`.

**Vantagens:**
- Foca no essencial primeiro (mobile geralmente tem menos espaço)
- Melhor performance (menos CSS para carregar em mobile)
- Alinha com a maioria dos usuários (muitos acessam por mobile)
- Considerada a melhor prática moderna

**Como funciona:**
Você escreve os estilos base para mobile, e depois usa `@media (min-width: ...)` para adicionar estilos para telas maiores.

#### Desktop-First (Desktop Primeiro)

**Desktop-First** significa que você começa criando o design para desktops (telas grandes) e depois ajusta para telas menores usando `max-width`.

**Vantagens:**
- Pode ser mais intuitivo se você está acostumado a trabalhar em desktop
- Útil quando o design desktop é mais complexo

**Desvantagens:**
- Pode resultar em mais CSS sendo carregado em mobile
- Não é a abordagem recomendada atualmente

### Combinando Múltiplas Condições

Você pode combinar múltiplas condições usando `and`:

```css
@media (min-width: 768px) and (max-width: 1024px) {
  /* estilos para tablets */
}
```

Isso significa: "aplique esses estilos quando a tela tiver pelo menos 768px E no máximo 1024px".

### Quando Usar Media Queries?

- Quando você precisa ajustar o layout para diferentes tamanhos de tela
- Para esconder ou mostrar elementos em certos dispositivos
- Para ajustar tamanhos de fonte e espaçamentos
- Para mudar a direção de layouts (de horizontal para vertical)
- Para otimizar a experiência em diferentes dispositivos

### Limitações das Media Queries

Media Queries são baseadas no **viewport** (a área visível da janela do navegador), não no tamanho do elemento específico. Isso significa que se você tem um componente que pode aparecer em diferentes tamanhos dentro da mesma página, media queries não conseguem adaptar esse componente baseado no espaço que ele ocupa.

---

## 📦 Container Queries

### O que são Container Queries?

**Container Queries** são uma funcionalidade mais recente do CSS que permite aplicar estilos baseados no tamanho ou outras características do **elemento container** (o elemento pai), ao invés do tamanho da tela inteira (viewport). Isso permite criar componentes verdadeiramente reutilizáveis que se adaptam ao espaço disponível, não importa onde estejam na página.

### Por que Container Queries são Importantes?

Imagine que você tem um card de produto que pode aparecer em diferentes lugares: em uma sidebar estreita, em uma grade de produtos, ou em destaque na página principal. Com Media Queries, você só pode adaptar baseado no tamanho da tela inteira. Com Container Queries, o card se adapta baseado no espaço que ele realmente tem disponível, independente do tamanho da tela.

É a diferença entre escolher roupas baseado na temperatura do dia inteiro (Media Queries) versus escolher roupas baseado na temperatura do quarto onde você está (Container Queries).

### Como Funcionam as Container Queries?

Container Queries funcionam de forma similar às Media Queries, mas ao invés de verificar o viewport, elas verificam o tamanho do container. Para usar container queries, você precisa:

1. **Definir um container**: Usar `container-type` ou `container` no elemento que será o container
2. **Usar @container**: Escrever a query usando `@container` ao invés de `@media`

### Sintaxe de Container Queries

Primeiro, você define o container:

```css
.produto-container {
  container-type: inline-size;
}
```

Depois, você usa `@container` para aplicar estilos baseados no tamanho do container:

```css
@container (min-width: 300px) {
  .card-produto {
    /* estilos quando o container tem pelo menos 300px */
  }
}
```

### Tipos de Container

- **`inline-size`**: O container responde à largura (width)
- **`block-size`**: O container responde à altura (height)
- **`size`**: O container responde tanto à largura quanto à altura

### Quando Usar Container Queries?

- Quando você tem componentes que aparecem em diferentes tamanhos na mesma página
- Para criar componentes verdadeiramente reutilizáveis
- Quando o layout do componente depende do espaço disponível, não do tamanho da tela
- Para designs mais granulares e contextuais

### Limitações das Container Queries

- Suporte de navegadores ainda está se expandindo (mais recente que Media Queries)
- Requer configuração do container antes de usar
- Não substituem Media Queries completamente - ambas têm seus usos

### Container Queries vs Media Queries

**Media Queries:** "Como está a tela inteira?"
**Container Queries:** "Como está o espaço disponível para este componente específico?"

Ambas são úteis e se complementam. Use Media Queries para ajustes gerais da página e Container Queries para componentes específicos.

---

## 📝 Responsive Typography (Tipografia Responsiva)

### O que é Responsive Typography?

**Responsive Typography** (Tipografia Responsiva) é sobre fazer o texto de uma página web parecer bom e ser facilmente legível em diferentes tamanhos de tela e dispositivos. Envolve ajustar tamanhos de fonte, alturas de linha, espaçamentos entre letras e outras propriedades de texto para garantir legibilidade ótima, seja em um monitor grande de desktop ou em um smartphone pequeno.

### Por que Tipografia Responsiva é Importante?

Texto que é perfeitamente legível em um desktop pode ficar muito pequeno em um celular, ou muito grande e quebrar o layout. Texto que funciona bem em um tablet pode não funcionar bem em nenhum dos dois. Tipografia responsiva garante que o texto seja sempre legível e agradável, independente do dispositivo.

### Como Funcionar a Tipografia Responsiva?

Tipografia responsiva funciona ajustando as propriedades de texto baseado no tamanho da tela ou do container. Você pode fazer isso de várias formas:

1. **Usando Media Queries**: Ajustar tamanhos de fonte em diferentes breakpoints
2. **Usando unidades relativas**: Usar `rem`, `em`, `vw`, `vh` ao invés de `px` fixos
3. **Usando funções CSS**: Usar `clamp()`, `min()`, `max()` para criar tamanhos fluidos
4. **Usando Container Queries**: Ajustar baseado no tamanho do container

### Unidades para Tipografia Responsiva

#### Unidades Relativas

- **`rem`**: Relativo ao tamanho da fonte do elemento raiz (`<html>`). Se o root tem 16px, 1rem = 16px, 2rem = 32px
- **`em`**: Relativo ao tamanho da fonte do elemento pai. Se o pai tem 18px, 1em = 18px
- **`%`**: Relativo ao tamanho da fonte do elemento pai (similar ao em)
- **`vw`**: 1vw = 1% da largura do viewport. 10vw = 10% da largura da tela
- **`vh`**: 1vh = 1% da altura do viewport. 10vh = 10% da altura da tela

#### Por que Usar Unidades Relativas?

Unidades relativas permitem que o texto escale proporcionalmente. Se o usuário aumentar o tamanho da fonte no navegador, o texto usando `rem` ou `em` vai aumentar também. Texto usando `px` fixo não vai aumentar.

### Técnicas de Tipografia Responsiva

#### 1. Media Queries para Tamanhos de Fonte

Ajustar tamanhos de fonte em diferentes breakpoints:

```css
h1 {
  font-size: 24px; /* mobile */
}

@media (min-width: 768px) {
  h1 {
    font-size: 32px; /* tablet */
  }
}

@media (min-width: 1024px) {
  h1 {
    font-size: 48px; /* desktop */
  }
}
```

#### 2. Tamanhos de Fonte Fluidos com clamp()

A função `clamp()` permite criar tamanhos que se ajustam automaticamente entre um mínimo e máximo:

```css
h1 {
  font-size: clamp(24px, 5vw, 48px);
}
```

Isso significa: "use 24px como mínimo, 5vw como ideal, e 48px como máximo. Ajuste automaticamente entre esses valores."

#### 3. Line Height Responsivo

Altura de linha também deve ser ajustada para diferentes tamanhos de tela. Geralmente, textos menores precisam de line-height maior para legibilidade.

#### 4. Letter Spacing e Word Spacing

Espaçamentos entre letras e palavras podem ser ajustados para melhorar legibilidade em diferentes tamanhos.

### Boas Práticas de Tipografia Responsiva

- **Tamanhos mínimos**: Nunca use font-size menor que 16px para texto do corpo (legibilidade)
- **Contraste**: Garanta contraste suficiente entre texto e fundo
- **Line-height**: Use line-height entre 1.4 e 1.6 para texto do corpo
- **Hierarquia**: Mantenha hierarquia visual clara entre títulos e texto
- **Teste**: Sempre teste em dispositivos reais, não apenas redimensionando a janela

### Quando Usar Tipografia Responsiva?

- Sempre! Tipografia responsiva deve ser parte de todo projeto web moderno
- Especialmente importante para sites com muito conteúdo textual
- Crítico para acessibilidade e experiência do usuário

---

## 🎨 CSS Variables (Custom Properties)

### O que são CSS Variables?

**CSS Variables**, também conhecidas como **Custom Properties** (Propriedades Customizadas), são entidades definidas por você que contêm valores específicos para serem reutilizados em toda a folha de estilo. Elas permitem armazenar um valor em um lugar e referenciá-lo em múltiplos outros lugares, tornando seu código CSS mais fácil de manter e atualizar.

### Por que CSS Variables são Importantes?

Antes das CSS Variables, se você quisesse usar a mesma cor em 50 lugares diferentes, você teria que escrever o valor da cor 50 vezes. Se depois você quisesse mudar essa cor, teria que encontrar e substituir em 50 lugares. Com CSS Variables, você define a cor uma vez e a usa em todos os lugares. Se precisar mudar, muda em um lugar só.

É como ter um dicionário de cores e valores que você pode consultar sempre que precisar, ao invés de ter que lembrar ou copiar valores toda vez.

### Como Funcionam as CSS Variables?

CSS Variables funcionam em duas etapas:

1. **Definir a variável**: Você cria a variável e atribui um valor a ela
2. **Usar a variável**: Você referencia a variável usando a função `var()`

### Sintaxe de CSS Variables

#### Definindo Variáveis

Variáveis CSS são definidas usando dois hífens (`--`) seguidos do nome da variável:

```css
:root {
  --cor-primaria: #3498db;
  --cor-secundaria: #2ecc71;
  --espacamento-padrao: 16px;
}
```

O `:root` é um seletor especial que representa o elemento raiz do documento (geralmente `<html>`). Definir variáveis em `:root` as torna disponíveis em todo o documento.

#### Usando Variáveis

Para usar uma variável, você usa a função `var()`:

```css
.botao {
  background-color: var(--cor-primaria);
  padding: var(--espacamento-padrao);
}
```

### Escopo de Variáveis

Variáveis CSS têm **escopo**, o que significa que elas podem ser definidas em diferentes níveis:

- **`:root`**: Disponível em todo o documento (escopo global)
- **Elemento específico**: Disponível apenas para aquele elemento e seus filhos (escopo local)

**Exemplo de escopo local:**

```css
.card {
  --cor-fundo: #ffffff;
  background-color: var(--cor-fundo);
}

.outro-elemento {
  /* --cor-fundo não está disponível aqui */
}
```

### Fallback Values (Valores de Fallback)

Você pode fornecer um valor de fallback (reserva) caso a variável não esteja definida:

```css
.botao {
  color: var(--cor-texto, #000000);
}
```

Isso significa: "use `--cor-texto` se estiver disponível, caso contrário use `#000000`".

### Vantagens das CSS Variables

1. **Manutenibilidade**: Mude um valor em um lugar, afeta todos os lugares que usam
2. **Consistência**: Garante que valores relacionados sejam sempre os mesmos
3. **Organização**: Centraliza valores importantes em um lugar
4. **Flexibilidade**: Pode ser alterada via JavaScript
5. **Temas**: Facilita criação de temas (modo claro/escuro)

### Casos de Uso Comuns

- **Cores**: Definir paleta de cores centralizada
- **Espaçamentos**: Valores de padding e margin consistentes
- **Tamanhos de fonte**: Sistema tipográfico centralizado
- **Breakpoints**: Valores de media queries reutilizáveis
- **Temas**: Alternar entre temas claro/escuro facilmente

### CSS Variables vs Valores Fixos

**Antes (sem variáveis):**
```css
.botao {
  background-color: #3498db;
}

.link {
  color: #3498db;
}

.titulo {
  border-bottom: 2px solid #3498db;
}
```

**Depois (com variáveis):**
```css
:root {
  --cor-primaria: #3498db;
}

.botao {
  background-color: var(--cor-primaria);
}

.link {
  color: var(--cor-primaria);
}

.titulo {
  border-bottom: 2px solid var(--cor-primaria);
}
```

Se você quiser mudar a cor primária, muda apenas em `:root`!

### Quando Usar CSS Variables?

- Quando você tem valores que são usados em múltiplos lugares
- Para criar sistemas de design consistentes
- Para facilitar manutenção e atualizações
- Para criar temas dinâmicos
- Para valores que podem mudar (cores, espaçamentos, tamanhos)

---

## ⚙️ CSS Functions (Funções CSS)

### O que são CSS Functions?

**CSS Functions** são operações pré-definidas que realizam tarefas específicas dentro do seu código CSS. Elas permitem manipular valores, realizar cálculos e gerar resultados dinâmicos, tornando suas folhas de estilo mais flexíveis e poderosas.

### Por que CSS Functions são Importantes?

Funções CSS permitem que você faça coisas que valores estáticos não conseguem. Elas permitem cálculos, transformações, e lógica dentro do CSS, tornando possível criar designs mais dinâmicos e responsivos sem precisar de JavaScript.

### Como Funcionam as CSS Functions?

Funções CSS são escritas com o nome da função seguido de parênteses, dentro dos quais você passa os argumentos (valores) necessários. A função processa esses valores e retorna um resultado que é usado como valor da propriedade CSS.

### Sintaxe de Funções CSS

A sintaxe geral é:

```css
propriedade: nome-da-funcao(argumento1, argumento2, ...);
```

### Funções CSS Comuns

#### 1. calc() - Cálculos

A função `calc()` permite realizar cálculos matemáticos:

```css
.largura {
  width: calc(100% - 40px);
}
```

Isso significa: "a largura é 100% do container menos 40 pixels". Você pode usar `+`, `-`, `*`, `/` dentro de `calc()`.

**Por que é útil?** Permite combinar unidades diferentes (como % e px) que normalmente não podem ser combinadas.

#### 2. var() - Variáveis CSS

A função `var()` é usada para acessar variáveis CSS (já vimos isso na seção anterior):

```css
.cor {
  color: var(--minha-variavel);
}
```

#### 3. clamp() - Valores Limitados

A função `clamp()` força um valor a ficar entre um mínimo e máximo:

```css
.tamanho {
  font-size: clamp(16px, 4vw, 24px);
}
```

Isso significa: "use 16px como mínimo, 4vw como valor preferido, e 24px como máximo. Escolha o valor apropriado entre esses limites."

**Por que é útil?** Cria valores responsivos que se ajustam automaticamente mas nunca ficam muito pequenos ou muito grandes.

#### 4. min() - Valor Mínimo

A função `min()` retorna o menor valor entre os fornecidos:

```css
.largura {
  width: min(100%, 500px);
}
```

Isso significa: "use 100% ou 500px, o que for menor".

#### 5. max() - Valor Máximo

A função `max()` retorna o maior valor entre os fornecidos:

```css
.largura {
  width: max(300px, 50%);
}
```

Isso significa: "use 300px ou 50%, o que for maior".

#### 6. rgb() e rgba() - Cores

Funções para definir cores usando valores RGB:

```css
.cor {
  background-color: rgb(52, 152, 219);
  color: rgba(0, 0, 0, 0.8); /* com transparência */
}
```

#### 7. hsl() e hsla() - Cores HSL

Funções para definir cores usando HSL (Hue, Saturation, Lightness):

```css
.cor {
  background-color: hsl(200, 70%, 50%);
  color: hsla(0, 0%, 0%, 0.5);
}
```

#### 8. linear-gradient() - Gradientes Lineares

Cria gradientes lineares:

```css
.fundo {
  background: linear-gradient(to right, #3498db, #2ecc71);
}
```

#### 9. url() - URLs

Usada para referenciar arquivos externos:

```css
.imagem {
  background-image: url('caminho/para/imagem.jpg');
}
```

### Combinando Funções

Você pode combinar múltiplas funções:

```css
.tamanho {
  font-size: clamp(16px, calc(1rem + 1vw), 24px);
}
```

### Quando Usar CSS Functions?

- **calc()**: Quando você precisa combinar unidades diferentes ou fazer cálculos
- **clamp()**: Para valores responsivos que precisam de limites
- **min()/max()**: Para garantir que valores não fiquem muito pequenos ou grandes
- **var()**: Sempre que usar CSS Variables
- **Gradientes**: Para criar efeitos visuais complexos
- **Cores**: Para definir cores de forma programática

### Vantagens das CSS Functions

1. **Flexibilidade**: Permite lógica e cálculos no CSS
2. **Responsividade**: Cria valores que se adaptam automaticamente
3. **Manutenibilidade**: Facilita ajustes e cálculos complexos
4. **Performance**: Cálculos são feitos pelo navegador, não por JavaScript

### Limitações

- Nem todas as funções são suportadas em navegadores antigos
- Algumas funções podem ser complexas de entender inicialmente
- Requer entendimento dos conceitos matemáticos básicos

---

## 🔗 Como Tudo se Conecta

Media Queries, Container Queries, Responsive Typography, CSS Variables e CSS Functions trabalham juntos para criar designs modernos, responsivos e manuteníveis:

- **Media Queries** ajustam o layout geral baseado no tamanho da tela
- **Container Queries** ajustam componentes baseado no espaço disponível
- **Responsive Typography** garante que o texto seja sempre legível
- **CSS Variables** centralizam valores para fácil manutenção
- **CSS Functions** permitem cálculos e valores dinâmicos

Juntas, essas ferramentas permitem criar websites que são ao mesmo tempo bonitos, funcionais e fáceis de manter.

---

## 📚 Resumo dos Conceitos Principais

### Media Queries
- Aplicam estilos baseados no tamanho da tela/viewport
- Sintaxe: `@media (condição) { estilos }`
- Usam breakpoints para diferentes dispositivos
- Abordagem mobile-first é recomendada

### Container Queries
- Aplicam estilos baseados no tamanho do container
- Sintaxe: `@container (condição) { estilos }`
- Requerem definição do container primeiro
- Úteis para componentes reutilizáveis

### Responsive Typography
- Garante legibilidade em todos os dispositivos
- Usa unidades relativas (rem, em, vw, vh)
- Pode usar funções como clamp() para valores fluidos
- Deve sempre considerar acessibilidade

### CSS Variables
- Armazenam valores para reutilização
- Sintaxe: `--nome-variavel: valor;` e `var(--nome-variavel)`
- Têm escopo (global ou local)
- Facilitam manutenção e criação de temas

### CSS Functions
- Realizam operações e cálculos
- Exemplos: calc(), clamp(), min(), max(), var()
- Permitem valores dinâmicos e responsivos
- Tornam CSS mais poderoso e flexível

---

## 🎓 Próximos Passos

Agora que você entendeu os conceitos fundamentais de responsividade, variáveis e funções em CSS, você está pronto para praticar. Na próxima etapa, você verá uma versão simplificada desses conceitos com analogias do dia a dia, seguida de exercícios práticos para consolidar seu aprendizado.




