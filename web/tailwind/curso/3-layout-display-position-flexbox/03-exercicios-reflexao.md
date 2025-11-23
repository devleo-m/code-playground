# Aula 3 - Exercícios e Reflexão: Layout com Tailwind - Display, Position e Flexbox

## 📝 Exercícios Práticos

### Exercício 1: Criar um Header Responsivo

**Objetivo:** Criar um header de site usando Flexbox com Tailwind.

**Requisitos:**
- Logo à esquerda
- Menu de navegação centralizado (Home, Sobre, Contato)
- Botão "Login" à direita
- Padding de 1rem em todos os lados
- Background branco com sombra sutil
- Altura total de 4rem

**HTML base fornecido:**
```html
<header>
  <img src="logo.png" alt="Logo" class="h-8">
  <nav>
    <a href="#">Home</a>
    <a href="#">Sobre</a>
    <a href="#">Contato</a>
  </nav>
  <button>Login</button>
</header>
```

**Sua tarefa:** Adicione as classes Tailwind necessárias para criar o layout descrito.

---

### Exercício 2: Card de Produto com Layout Flex

**Objetivo:** Criar um card de produto usando `flex-col` e organizar os elementos.

**Requisitos:**
- Layout em coluna (vertical)
- Imagem no topo (largura total, altura de 200px)
- Título abaixo da imagem (texto grande, negrito)
- Descrição abaixo do título (texto cinza)
- Rodapé com preço à esquerda e botão "Comprar" à direita
- O rodapé deve ficar sempre na base do card
- Espaçamento de 1rem entre elementos
- Padding de 1.5rem
- Background branco, bordas arredondadas, sombra

**HTML base fornecido:**
```html
<div class="max-w-sm">
  <img src="produto.jpg" alt="Produto">
  <h3>Produto Incrível</h3>
  <p>Descrição detalhada do produto que você vai adorar.</p>
  <div>
    <span>R$ 99,90</span>
    <button>Comprar</button>
  </div>
</div>
```

**Sua tarefa:** Adicione as classes Tailwind necessárias.

---

### Exercício 3: Modal Centralizado

**Objetivo:** Criar um modal que aparece centralizado na tela com overlay de fundo.

**Requisitos:**
- Overlay de fundo que cobre toda a tela (preto com opacidade 50%)
- Modal centralizado (horizontal e verticalmente)
- Modal com largura de 400px
- Background branco, padding de 2rem, bordas arredondadas
- Botão de fechar no canto superior direito (posição absoluta)
- Z-index apropriado para aparecer sobre o overlay

**HTML base fornecido:**
```html
<!-- Overlay -->
<div>
  <!-- Modal -->
  <div>
    <button>×</button>
    <h2>Título do Modal</h2>
    <p>Conteúdo do modal aqui.</p>
  </div>
</div>
```

**Sua tarefa:** Adicione as classes Tailwind necessárias para posicionamento.

---

### Exercício 4: Sidebar com Layout

**Objetivo:** Criar um layout com sidebar fixa e conteúdo principal.

**Requisitos:**
- Sidebar à esquerda com largura fixa de 256px
- Sidebar com background escuro, texto branco
- Menu vertical na sidebar (Dashboard, Perfil, Configurações)
- Conteúdo principal ocupa o resto do espaço
- Layout deve ocupar altura total da tela
- Conteúdo principal com padding de 2rem

**HTML base fornecido:**
```html
<div>
  <aside>
    <h2>Menu</h2>
    <nav>
      <a href="#">Dashboard</a>
      <a href="#">Perfil</a>
      <a href="#">Configurações</a>
    </nav>
  </aside>
  <main>
    <h1>Conteúdo Principal</h1>
    <p>Este é o conteúdo da página...</p>
  </main>
</div>
```

**Sua tarefa:** Adicione as classes Tailwind necessárias.

---

### Exercício 5: Formulário com Layout Flex

**Objetivo:** Criar um formulário usando Flexbox.

**Requisitos:**
- Layout em coluna para os campos
- Cada campo tem label acima e input abaixo
- Espaçamento de 1rem entre campos
- Dois botões no final: "Enviar" e "Cancelar" lado a lado
- Botões devem ter largura igual (ocupar espaço igual)
- Largura máxima do formulário de 500px
- Centralizar o formulário na página

**HTML base fornecido:**
```html
<form>
  <div>
    <label>Nome</label>
    <input type="text">
  </div>
  <div>
    <label>Email</label>
    <input type="email">
  </div>
  <div>
    <button type="submit">Enviar</button>
    <button type="button">Cancelar</button>
  </div>
</form>
```

**Sua tarefa:** Adicione as classes Tailwind necessárias.

---

### Exercício 6: Sticky Header

**Objetivo:** Criar um header que "gruda" no topo ao rolar a página.

**Requisitos:**
- Header com position sticky
- Deve "grudar" quando chegar no topo (top-0)
- Background branco com sombra quando sticky
- Z-index alto para ficar sobre outros elementos
- Conteúdo do header: logo à esquerda, menu centralizado, botão à direita
- Adicione conteúdo suficiente abaixo para poder rolar a página

**HTML base fornecido:**
```html
<header>
  <img src="logo.png" alt="Logo" class="h-8">
  <nav>
    <a href="#">Home</a>
    <a href="#">Sobre</a>
  </nav>
  <button>Login</button>
</header>
<main>
  <!-- Adicione conteúdo suficiente para rolar -->
  <div style="height: 200vh;">
    <h1>Conteúdo da Página</h1>
    <p>Role a página para ver o header grudar no topo...</p>
  </div>
</main>
```

**Sua tarefa:** Adicione as classes Tailwind necessárias.

---

### Exercício 7: Grid de Cards com Flexbox

**Objetivo:** Criar um grid de cards usando Flexbox com wrap.

**Requisitos:**
- Container flex com wrap
- Cada card tem largura de 300px
- Espaçamento de 1.5rem entre cards
- Cards devem quebrar linha quando necessário
- Cards centralizados quando não preenchem a linha completa
- Cada card tem: imagem, título, descrição, botão

**HTML base fornecido:**
```html
<div>
  <div>
    <img src="card1.jpg" alt="Card 1">
    <h3>Card 1</h3>
    <p>Descrição do card 1</p>
    <button>Ver mais</button>
  </div>
  <div>
    <img src="card2.jpg" alt="Card 2">
    <h3>Card 2</h3>
    <p>Descrição do card 2</p>
    <button>Ver mais</button>
  </div>
  <div>
    <img src="card3.jpg" alt="Card 3">
    <h3>Card 3</h3>
    <p>Descrição do card 3</p>
    <button>Ver mais</button>
  </div>
</div>
```

**Sua tarefa:** Adicione as classes Tailwind necessárias.

---

### Exercício 8: Análise de Código - Identificar Problemas

**Objetivo:** Analise o código abaixo e identifique problemas de layout.

**Código fornecido:**
```html
<div class="flex">
  <div class="w-64">Sidebar</div>
  <div>Conteúdo Principal</div>
</div>

<div class="flex justify-center">
  <div class="w-32 h-32 bg-blue-500"></div>
  <div class="w-32 h-32 bg-red-500"></div>
</div>

<div class="flex items-center">
  <div>Item 1</div>
  <div>Item 2</div>
</div>
```

**Perguntas:**
1. No primeiro exemplo, o conteúdo principal não está ocupando o espaço disponível. Como corrigir?
2. No segundo exemplo, os quadrados não estão centralizados verticalmente. O que falta?
3. No terceiro exemplo, os itens não estão centralizados horizontalmente. O que falta?

**Sua tarefa:** Identifique os problemas e forneça as correções.

---

## 🤔 Perguntas de Reflexão

### Reflexão 1: Display vs Position

**Pergunta:** Qual é a diferença fundamental entre `display: flex` e `position: absolute`? Quando você usaria cada um?

**Contexto para pensar:**
- `display: flex` muda como o elemento e seus filhos se comportam no layout
- `position: absolute` muda onde o elemento é posicionado no documento

**Questões para considerar:**
- Um elemento pode ter `display: flex` e `position: absolute` ao mesmo tempo? Por quê?
- Se você quer centralizar um elemento na tela, qual abordagem é melhor: `flex` com `justify-center` e `items-center`, ou `position: absolute` com `top-1/2 left-1/2`? Por quê?
- Qual tem mais impacto no layout dos elementos ao redor?

---

### Reflexão 2: Flexbox vs CSS Grid

**Pergunta:** Você já conhece CSS Grid. Quando seria melhor usar Flexbox (`flex`) e quando seria melhor usar Grid (`grid`)?

**Contexto para pensar:**
- Flexbox é unidimensional (linha OU coluna)
- Grid é bidimensional (linhas E colunas simultaneamente)

**Questões para considerar:**
- Para um menu horizontal, você usaria Flexbox ou Grid? Por quê?
- Para um layout de blog com sidebar e conteúdo principal, qual seria melhor?
- Para um grid de produtos (3 colunas, múltiplas linhas), qual seria melhor?
- Existe algum caso onde ambos funcionariam igualmente bem?

---

### Reflexão 3: Performance e Especificidade

**Pergunta:** Usar muitas classes Tailwind (como `flex items-center justify-between gap-4 p-6 bg-white rounded-lg shadow-md`) tem algum impacto negativo? Como isso se compara com escrever CSS customizado?

**Contexto para pensar:**
- Tailwind gera CSS baseado nas classes usadas
- CSS customizado é escrito diretamente
- Especificidade e tamanho do arquivo CSS final

**Questões para considerar:**
- Qual abordagem gera mais CSS no arquivo final?
- Qual é mais fácil de manter e modificar?
- Qual tem melhor performance de renderização?
- Em um projeto grande, qual seria mais escalável?
- Como o PurgeCSS do Tailwind afeta essa comparação?

---

### Reflexão 4: Responsividade com Flexbox

**Pergunta:** Como você faria um layout que é horizontal em telas grandes e vertical em telas pequenas usando apenas Flexbox?

**Contexto para pensar:**
- Mobile-first: começar com layout mobile e adaptar para desktop
- Breakpoints do Tailwind: `sm:`, `md:`, `lg:`, `xl:`, `2xl:`

**Questões para considerar:**
- Você começaria com `flex-col` ou `flex-row`? Por quê?
- Como usar breakpoints do Tailwind para mudar a direção?
- Existe alguma limitação do Flexbox para layouts muito complexos?
- Quando seria melhor usar Grid para responsividade?

**Exemplo prático:**
Crie um header que:
- Em mobile: logo e menu em coluna (vertical)
- Em desktop: logo e menu em linha (horizontal)

---

### Reflexão 5: Acessibilidade e Layout

**Pergunta:** Como o uso de Flexbox e Position afeta a acessibilidade de uma página? O que você precisa considerar?

**Contexto para pensar:**
- Leitores de tela navegam pelo DOM na ordem do HTML
- `position: absolute` e `flex-direction: row-reverse` podem mudar a ordem visual, mas não a ordem do DOM
- Foco do teclado segue a ordem do DOM

**Questões para considerar:**
- Se você usa `flex-row-reverse` para inverter a ordem visual, isso afeta leitores de tela?
- Um elemento com `position: absolute` que está visualmente no topo, mas no HTML está no final, como leitores de tela o leem?
- Como garantir que a ordem de foco do teclado faça sentido mesmo com layouts complexos?
- Qual é a melhor prática: mudar a ordem no HTML ou usar CSS para ordem visual?

---

### Reflexão 6: Manutenibilidade e Escalabilidade

**Pergunta:** Em um projeto grande com múltiplos desenvolvedores, quais são as vantagens e desvantagens de usar Tailwind para layout vs CSS customizado?

**Contexto para pensar:**
- Tailwind: classes utilitárias padronizadas
- CSS customizado: classes semânticas específicas do projeto
- Trabalho em equipe: consistência e comunicação

**Questões para considerar:**
- É mais fácil para um novo desenvolvedor entender `flex items-center justify-between` ou uma classe customizada `.header-container`?
- Como garantir consistência de espaçamento em um projeto grande?
- Qual abordagem é mais fácil de refatorar quando o design muda?
- Como lidar com casos muito específicos que não se encaixam nas utilities do Tailwind?
- Qual tem melhor suporte para design systems e componentes reutilizáveis?

---

## 📊 Exercício de Análise: Comparar Abordagens

**Objetivo:** Compare duas implementações do mesmo layout e analise as diferenças.

### Implementação A: Tailwind Utilities

```html
<div class="flex items-center justify-between p-6 bg-white rounded-lg shadow-md">
  <div class="flex items-center gap-4">
    <img src="avatar.jpg" alt="Avatar" class="w-12 h-12 rounded-full">
    <div class="flex flex-col">
      <h3 class="text-lg font-bold">João Silva</h3>
      <p class="text-sm text-gray-600">Desenvolvedor</p>
    </div>
  </div>
  <button class="px-4 py-2 bg-blue-500 text-white rounded hover:bg-blue-600">
    Seguir
  </button>
</div>
```

### Implementação B: CSS Customizado

```html
<div class="user-card">
  <div class="user-info">
    <img src="avatar.jpg" alt="Avatar" class="avatar">
    <div class="user-details">
      <h3>João Silva</h3>
      <p>Desenvolvedor</p>
    </div>
  </div>
  <button class="follow-button">Seguir</button>
</div>
```

```css
.user-card {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 1.5rem;
  background-color: white;
  border-radius: 0.5rem;
  box-shadow: 0 4px 6px -1px rgb(0 0 0 / 0.1);
}

.user-info {
  display: flex;
  align-items: center;
  gap: 1rem;
}

.user-details {
  display: flex;
  flex-direction: column;
}

.avatar {
  width: 3rem;
  height: 3rem;
  border-radius: 9999px;
}

.follow-button {
  padding: 0.5rem 1rem;
  background-color: rgb(59 130 246);
  color: white;
  border-radius: 0.25rem;
}

.follow-button:hover {
  background-color: rgb(37 99 235);
}
```

**Perguntas para análise:**

1. **Legibilidade:** Qual implementação é mais fácil de ler e entender à primeira vista?

2. **Manutenibilidade:** Se você precisar mudar o espaçamento entre o avatar e o nome, qual é mais fácil de modificar?

3. **Reutilização:** Se você precisar criar um card similar mas com layout ligeiramente diferente, qual abordagem facilita mais?

4. **Performance:** Qual gera menos CSS no arquivo final? (Considere que o Tailwind usa PurgeCSS)

5. **Consistência:** Em um projeto grande, qual abordagem garante mais consistência visual?

6. **Aprendizado:** Para um desenvolvedor iniciante, qual é mais fácil de aprender e usar?

7. **Flexibilidade:** Qual permite mais flexibilidade para casos muito específicos?

**Sua tarefa:** Responda cada pergunta com justificativa baseada nos conceitos aprendidos.

---

## ✅ Checklist de Aprendizado

Antes de avançar para a próxima aula, certifique-se de que você consegue:

- [ ] Criar layouts usando `flex`, `flex-row`, `flex-col`
- [ ] Centralizar elementos horizontalmente e verticalmente
- [ ] Usar `justify-content` e `align-items` corretamente
- [ ] Posicionar elementos com `relative`, `absolute`, `fixed`, `sticky`
- [ ] Usar `gap` para espaçamento entre itens flex
- [ ] Criar headers, cards, formulários e sidebars com Flexbox
- [ ] Entender quando usar `flex-1`, `grow`, `shrink`
- [ ] Mapear classes Tailwind para propriedades CSS equivalentes
- [ ] Decidir quando usar Flexbox vs outras abordagens de layout
- [ ] Considerar acessibilidade ao criar layouts

---

## 🎯 Próximos Passos

Após completar os exercícios e reflexões:

1. **Revise suas respostas** - Certifique-se de entender os conceitos, não apenas memorizar
2. **Pratique criando layouts reais** - Tente recriar layouts de sites que você conhece
3. **Experimente combinações** - Misture diferentes utilities para ver o que funciona
4. **Compare com CSS puro** - Sempre relacione as classes Tailwind com o CSS equivalente

Na próxima aula, exploraremos **CSS Grid com Tailwind**, que oferece controle bidimensional ainda mais poderoso!

