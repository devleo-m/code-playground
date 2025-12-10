# Aula 5 - Exercícios e Reflexão: Responsividade, Variáveis e Funções

## 🎯 Objetivo dos Exercícios

Estes exercícios foram criados para consolidar seu aprendizado sobre Media Queries, Container Queries, Responsive Typography, CSS Variables e CSS Functions. Eles vão desde o básico até desafios que combinam múltiplos conceitos. Faça cada exercício com calma e pense sobre o que está fazendo.

---

## 📝 Exercício 1: Criando Media Queries Básicas

### Tarefa:
Crie um layout responsivo simples usando Media Queries. O objetivo é fazer com que:

1. **Mobile (até 768px):**
   - O texto tenha tamanho de 16px
   - Os elementos tenham padding de 10px
   - O fundo seja azul claro

2. **Tablet (769px a 1024px):**
   - O texto tenha tamanho de 18px
   - Os elementos tenham padding de 20px
   - O fundo seja verde claro

3. **Desktop (acima de 1024px):**
   - O texto tenha tamanho de 20px
   - Os elementos tenham padding de 30px
   - O fundo seja amarelo claro

### HTML de Referência:
```html
<!DOCTYPE html>
<html>
<head>
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <link rel="stylesheet" href="estilos.css">
</head>
<body>
  <div class="container">
    <h1>Título Responsivo</h1>
    <p>Este é um parágrafo que se adapta ao tamanho da tela.</p>
  </div>
</body>
</html>
```

### O que você deve escrever no `estilos.css`:
(Escreva suas regras CSS usando abordagem mobile-first)

---

## 📝 Exercício 2: Usando CSS Variables

### Tarefa:
Crie um sistema de cores usando CSS Variables. O objetivo é:

1. Definir variáveis para:
   - Cor primária: #3498db (azul)
   - Cor secundária: #2ecc71 (verde)
   - Cor de texto: #333333 (cinza escuro)
   - Espaçamento padrão: 16px
   - Tamanho de fonte base: 16px

2. Usar essas variáveis para estilizar:
   - Um botão com cor primária
   - Um link com cor secundária
   - Um parágrafo com cor de texto e espaçamento padrão
   - Um título com tamanho de fonte baseado na variável

### HTML de Referência:
```html
<!DOCTYPE html>
<html>
<head>
  <link rel="stylesheet" href="estilos.css">
</head>
<body>
  <h1>Título</h1>
  <p>Este é um parágrafo de exemplo.</p>
  <a href="#" class="link">Link de exemplo</a>
  <button class="botao">Botão</button>
</body>
</html>
```

### O que você deve escrever no `estilos.css`:
(Defina as variáveis em :root e use-as nos elementos)

---

## 📝 Exercício 3: Tipografia Responsiva com clamp()

### Tarefa:
Crie um sistema de tipografia responsiva usando a função `clamp()`. O objetivo é:

1. **Título h1:**
   - Tamanho mínimo: 24px
   - Tamanho preferido: 5vw (5% da largura da viewport)
   - Tamanho máximo: 48px

2. **Título h2:**
   - Tamanho mínimo: 20px
   - Tamanho preferido: 4vw
   - Tamanho máximo: 36px

3. **Parágrafo:**
   - Tamanho mínimo: 16px
   - Tamanho preferido: 2.5vw
   - Tamanho máximo: 18px
   - Line-height: 1.6

### HTML de Referência:
```html
<!DOCTYPE html>
<html>
<head>
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <link rel="stylesheet" href="estilos.css">
</head>
<body>
  <h1>Título Principal Responsivo</h1>
  <h2>Subtítulo Responsivo</h2>
  <p>Este é um parágrafo com tipografia responsiva que se ajusta automaticamente ao tamanho da tela, garantindo sempre legibilidade.</p>
</body>
</html>
```

### O que você deve escrever no `estilos.css`:
(Use clamp() para criar tamanhos de fonte fluidos)

---

## 📝 Exercício 4: Combinando calc() e CSS Variables

### Tarefa:
Crie um layout que usa tanto CSS Variables quanto a função `calc()`. O objetivo é:

1. Definir variáveis para:
   - Largura do container: 1200px
   - Padding lateral: 20px
   - Gap entre elementos: 16px

2. Criar um container que:
   - Tem largura máxima baseada na variável
   - Usa calc() para calcular a largura considerando o padding lateral
   - Centraliza na página

3. Criar cards dentro do container que:
   - Usam calc() para calcular largura considerando o gap
   - Têm espaçamento baseado nas variáveis

### HTML de Referência:
```html
<!DOCTYPE html>
<html>
<head>
  <link rel="stylesheet" href="estilos.css">
</head>
<body>
  <div class="container">
    <div class="card">Card 1</div>
    <div class="card">Card 2</div>
    <div class="card">Card 3</div>
  </div>
</body>
</html>
```

### O que você deve escrever no `estilos.css`:
(Combine variáveis CSS com calc() para criar um layout flexível)

---

## 📝 Exercício 5: Media Queries e Responsive Typography Juntos

### Tarefa:
Crie um card de produto que se adapta a diferentes tamanhos de tela, combinando Media Queries com tipografia responsiva. O objetivo é:

1. **Mobile (até 768px):**
   - Card ocupa 100% da largura
   - Título: 20px
   - Preço: 18px
   - Descrição: 14px
   - Botão: padding 10px 20px

2. **Tablet (769px a 1024px):**
   - Card ocupa 48% da largura (2 colunas)
   - Título: 24px
   - Preço: 20px
   - Descrição: 16px
   - Botão: padding 12px 24px

3. **Desktop (acima de 1024px):**
   - Card ocupa 30% da largura (3 colunas)
   - Título: 28px
   - Preço: 22px
   - Descrição: 16px
   - Botão: padding 14px 28px

### HTML de Referência:
```html
<!DOCTYPE html>
<html>
<head>
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <link rel="stylesheet" href="estilos.css">
</head>
<body>
  <div class="produtos">
    <div class="card-produto">
      <h2 class="titulo-produto">Produto Exemplo</h2>
      <p class="preco">R$ 99,90</p>
      <p class="descricao">Esta é uma descrição do produto que se adapta ao tamanho da tela.</p>
      <button class="botao-comprar">Comprar</button>
    </div>
    <div class="card-produto">
      <h2 class="titulo-produto">Outro Produto</h2>
      <p class="preco">R$ 149,90</p>
      <p class="descricao">Outra descrição de produto responsiva.</p>
      <button class="botao-comprar">Comprar</button>
    </div>
    <div class="card-produto">
      <h2 class="titulo-produto">Terceiro Produto</h2>
      <p class="preco">R$ 199,90</p>
      <p class="descricao">Mais uma descrição responsiva.</p>
      <button class="botao-comprar">Comprar</button>
    </div>
  </div>
</body>
</html>
```

### O que você deve escrever no `estilos.css`:
(Use Media Queries para diferentes breakpoints e ajuste tipografia e layout)

---

## 📝 Exercício 6: Sistema de Tema com CSS Variables

### Tarefa:
Crie um sistema simples de tema claro/escuro usando CSS Variables. O objetivo é:

1. Definir variáveis para tema claro:
   - Cor de fundo: #ffffff
   - Cor de texto: #333333
   - Cor primária: #3498db

2. Definir variáveis para tema escuro:
   - Cor de fundo: #1a1a1a
   - Cor de texto: #ffffff
   - Cor primária: #5dade2

3. Criar uma classe `.tema-escuro` que sobrescreve as variáveis
4. Aplicar as variáveis em elementos da página

### HTML de Referência:
```html
<!DOCTYPE html>
<html>
<head>
  <link rel="stylesheet" href="estilos.css">
</head>
<body class="tema-claro">
  <div class="container">
    <h1>Título da Página</h1>
    <p>Este é um parágrafo de exemplo.</p>
    <button class="botao">Botão</button>
    <button onclick="document.body.classList.toggle('tema-escuro')">Alternar Tema</button>
  </div>
</body>
</html>
```

### O que você deve escrever no `estilos.css`:
(Crie variáveis para ambos os temas e use escopo para alternar)

---

## 🤔 Perguntas de Reflexão

### Reflexão 1: Mobile-First vs Desktop-First

**Pergunta:** Você está começando um novo projeto web. Qual abordagem você escolheria: Mobile-First ou Desktop-First? Por quê?

**Pense sobre:**
- Qual abordagem é mais eficiente em termos de código?
- Qual abordagem alinha melhor com a maioria dos usuários?
- Qual abordagem facilita a manutenção do código?
- Quais são as implicações de performance de cada abordagem?
- Como cada abordagem afeta o processo de desenvolvimento?

**Sua resposta:**

---

### Reflexão 2: Media Queries vs Container Queries

**Cenário:** Você está criando um componente de card de produto que pode aparecer em diferentes contextos:
- Em uma sidebar estreita (300px de largura)
- Em uma grade de produtos na página principal (cada card tem 400px)
- Em destaque na página (card tem 800px de largura)

**Pergunta:** Você usaria Media Queries ou Container Queries para fazer esse card se adaptar? Por quê? Em que situações cada uma seria mais apropriada?

**Pense sobre:**
- Qual ferramenta se adapta melhor ao contexto do componente?
- Como cada abordagem afeta a reutilização do componente?
- Qual oferece mais flexibilidade para o futuro?
- Quais são as limitações de cada abordagem?
- Como você combinaria ambas se necessário?

**Sua resposta:**

---

### Reflexão 3: CSS Variables e Manutenibilidade

**Cenário:** Você criou um site com 50 páginas, todas usando a mesma paleta de cores. Você definiu as cores diretamente em cada lugar onde são usadas (sem variáveis). Agora você precisa mudar a cor primária de azul para verde em todo o site.

**Pergunta:** 
1. Quais seriam os desafios de fazer essa mudança sem CSS Variables?
2. Como CSS Variables resolveriam esse problema?
3. Além de cores, que outros valores se beneficiariam de serem variáveis?
4. Quando NÃO faz sentido usar CSS Variables?

**Pense sobre:**
- O impacto na manutenibilidade do código
- O tempo necessário para fazer mudanças
- A probabilidade de erros ao fazer mudanças manuais
- A organização e estrutura do código
- Quando variáveis podem tornar o código mais complexo ao invés de mais simples

**Sua resposta:**

---

### Reflexão 4: Responsive Typography e Acessibilidade

**Cenário:** Você criou um site com tipografia responsiva usando `clamp()`. O texto se ajusta automaticamente entre 16px e 24px baseado no tamanho da tela. Um usuário com deficiência visual aumenta o tamanho da fonte no navegador, mas o texto não muda.

**Pergunta:** 
1. Por que o texto não mudou quando o usuário aumentou o tamanho da fonte no navegador?
2. Como você garantiria que o texto respeite as preferências do usuário?
3. Qual é a relação entre unidades relativas (rem, em) e acessibilidade?
4. Como você equilibraria design responsivo com acessibilidade?

**Pense sobre:**
- A diferença entre unidades absolutas (px) e relativas (rem, em)
- Como as preferências do usuário devem ser respeitadas
- O impacto na experiência de usuários com deficiências
- Como testar acessibilidade em diferentes cenários
- As diretrizes WCAG sobre tamanho de texto

**Sua resposta:**

---

### Reflexão 5: CSS Functions e Performance

**Pergunta:** Você está usando várias funções CSS (`calc()`, `clamp()`, `min()`, `max()`) em seu projeto. Isso afeta a performance do site? Quando o uso excessivo de funções CSS pode se tornar um problema?

**Pense sobre:**
- Como o navegador processa funções CSS
- O impacto no tempo de renderização
- Quando funções CSS são mais eficientes que valores fixos
- Quando valores fixos podem ser mais apropriados
- Como balancear flexibilidade com performance
- O impacto em dispositivos com menos poder de processamento

**Sua resposta:**

---

### Reflexão 6: Breakpoints e Design Systems

**Pergunta:** Você está criando um design system para uma empresa. Como você decidiria quais breakpoints usar? Você seguiria valores genéricos (como 768px, 1024px) ou criaria breakpoints customizados baseados no seu design específico?

**Pense sobre:**
- Como breakpoints genéricos podem não se adequar ao seu design
- O processo de identificar onde o design "quebra"
- A importância de testar em dispositivos reais
- Como documentar breakpoints em um design system
- A manutenibilidade de breakpoints customizados vs genéricos
- Como breakpoints afetam a experiência do usuário

**Sua resposta:**

---

## 💡 Dicas para os Exercícios

1. **Comece simples:** Não tente fazer tudo de uma vez. Comece com um conceito e vá adicionando outros gradualmente.

2. **Teste em diferentes tamanhos:** Use as ferramentas de desenvolvedor do navegador para testar em diferentes tamanhos de tela.

3. **Use nomes descritivos:** Quando criar variáveis CSS, use nomes que deixem claro o propósito (ex: `--cor-primaria` ao invés de `--cor1`).

4. **Pense em acessibilidade:** Sempre considere como suas escolhas afetam usuários com diferentes necessidades.

5. **Documente suas decisões:** Se você escolher breakpoints customizados ou valores específicos, anote o porquê para referência futura.

---

## 🎓 Próximos Passos

Após completar estes exercícios e reflexões, você terá uma base sólida em responsividade, variáveis e funções CSS. Na próxima etapa, você aprenderá sobre performance, boas práticas e otimização para aplicar esses conceitos de forma profissional.




