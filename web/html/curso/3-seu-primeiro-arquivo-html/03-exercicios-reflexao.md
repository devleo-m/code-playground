# Aula 3 - Exercícios e Reflexão

## 🎯 Exercícios Práticos

### Exercício 1: Criando Seu Primeiro Arquivo HTML Completo

Crie um arquivo HTML completo com a seguinte estrutura e conteúdo:

**Requisitos:**
- Use a estrutura básica de um documento HTML5
- Adicione o atributo `lang="pt-BR"` na tag `<html>`
- No `<head>`, adicione:
  - Meta tag `charset="UTF-8"`
  - Meta tag `viewport` para dispositivos móveis
  - Título: "Meu Primeiro Arquivo HTML"
- No `<body>`, adicione:
  - Um título principal (`<h1>`) com seu nome
  - Um parágrafo (`<p>`) de apresentação
  - Uma imagem (`<img>`) com atributos `src`, `alt`, `width` e `height`
  - Um link (`<a>`) com atributos `href` e `target="_blank"`
  - Use pelo menos 3 atributos globais diferentes (`id`, `class`, `title`)

**Dica:** Use o arquivo `exemplo-01-primeiro-arquivo.html` como referência.

---

### Exercício 2: Trabalhando com Tags e Atributos

Crie uma página HTML que demonstre o uso correto de tags e atributos:

**Requisitos:**
1. Crie uma seção com um título usando `<h2>` com `id="sobre-mim"`
2. Adicione um parágrafo com a classe `destaque`
3. Crie uma lista não ordenada (`<ul>`) com 3 itens, cada um com um atributo `title` explicativo
4. Adicione um link externo com:
   - `href` apontando para um site real
   - `target="_blank"` para abrir em nova aba
   - `title` com uma descrição
5. Adicione uma imagem com todos os atributos necessários (`src`, `alt`, `width`, `height`)
6. Crie um formulário simples com:
   - Um input de texto com `type="text"`, `name="usuario"`, `placeholder="Digite seu nome"` e `required`
   - Um botão com `type="submit"` e texto "Enviar"

**Tarefa adicional:** Identifique quais atributos são globais e quais são específicos de cada tag.

---

### Exercício 3: Usando HTML Entities

Crie uma página HTML que demonstre o uso de HTML entities:

**Requisitos:**
1. Crie um parágrafo que explique como escrever tags HTML usando entities:
   - Mostre `<p>` e `</p>` usando entities
   - Mostre `<div>` e `</div>` usando entities
2. Adicione símbolos usando entities:
   - Copyright (©)
   - Marca registrada (®)
   - Euro (€)
   - Seta para direita (→)
   - Sinal de multiplicação (×)
3. Crie uma seção sobre preços que use:
   - Símbolo de moeda (R$ ou €)
   - Sinal de mais ou menos (±) para variação de preço
4. Adicione um parágrafo com caracteres especiais em português usando entities:
   - Pelo menos 5 acentos diferentes (á, é, í, ó, ú, ç)

**Desafio:** Compare escrever os caracteres diretamente (com UTF-8) vs. usar entities. Quando cada abordagem é mais apropriada?

---

### Exercício 4: Organizando Código com Comentários

Analise o seguinte código HTML e adicione comentários apropriados para melhorar sua organização e legibilidade:

```html
<!DOCTYPE html>
<html lang="pt-BR">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Site de Receitas</title>
</head>
<body>
    <header>
        <h1>Receitas Deliciosas</h1>
        <nav>
            <ul>
                <li><a href="#doces">Doces</a></li>
                <li><a href="#salgados">Salgados</a></li>
                <li><a href="#bebidas">Bebidas</a></li>
            </ul>
        </nav>
    </header>
    <main>
        <section>
            <h2>Bolo de Chocolate</h2>
            <h3>Ingredientes</h3>
            <ul>
                <li>2 xícaras de farinha</li>
                <li>1 xícara de açúcar</li>
                <li>3 ovos</li>
                <li>1 xícara de leite</li>
            </ul>
            <h3>Modo de Preparo</h3>
            <ol>
                <li>Misture os ingredientes secos</li>
                <li>Adicione os líquidos</li>
                <li>Asse por 40 minutos</li>
            </ol>
        </section>
    </main>
    <footer>
        <p>Copyright 2024</p>
    </footer>
</body>
</html>
```

**Tarefa:**
1. Adicione comentários explicando cada seção principal
2. Adicione comentários para marcar o início e fim de blocos importantes
3. Adicione pelo menos um comentário TODO ou FIXME
4. Use comentários para organizar visualmente o código (com linhas separadoras)

---

### Exercício 5: Entendendo Whitespaces

Crie três versões da mesma página HTML para entender como whitespaces funcionam:

**Versão 1: Sem Formatação (Tudo em Uma Linha)**
```html
<!DOCTYPE html><html><head><title>Teste</title></head><body><h1>Título</h1><p>Parágrafo com     múltiplos     espaços</p></body></html>
```

**Versão 2: Com Formatação Normal**
```html
<!DOCTYPE html>
<html>
<head>
    <title>Teste</title>
</head>
<body>
    <h1>Título</h1>
    <p>Parágrafo com     múltiplos     espaços</p>
</body>
</html>
```

**Versão 3: Com Tag <pre>**
```html
<!DOCTYPE html>
<html>
<head>
    <title>Teste</title>
</head>
<body>
    <h1>Título</h1>
    <pre>Parágrafo com     múltiplos     espaços
    e quebras
    de linha</pre>
</body>
</html>
```

**Tarefa:**
1. Crie os três arquivos e abra cada um no navegador
2. Compare como cada versão é renderizada
3. Explique as diferenças que você observou
4. Crie um parágrafo que use `&nbsp;` para espaços não separáveis (ex: "Dr.&nbsp;Silva", "R$&nbsp;100,00")

---

### Exercício 6: Corrigindo Problemas de Case Sensitivity

O seguinte código HTML tem problemas de formatação e uso incorreto de maiúsculas/minúsculas. Corrija todos os problemas:

```html
<!DOCTYPE HTML>
<HTML LANG="pt-BR">
<HEAD>
    <META CHARSET="UTF-8">
    <TITLE>Minha Página</TITLE>
</HEAD>
<BODY>
    <H1 ID="titulo">Bem-vindo</H1>
    <P CLASS="destaque">Este é um parágrafo.</P>
    <IMG SRC="foto.jpg" ALT="Foto">
    <A HREF="https://www.exemplo.com" TARGET="_BLANK">Link</A>
</BODY>
</HTML>
```

**Tarefa:**
1. Converta todas as tags e atributos para minúsculas
2. Corrija a estrutura e indentação
3. Adicione comentários explicativos
4. Valide o código no W3C Validator

---

## 🤔 Perguntas de Reflexão

### Reflexão 1: Tags e Atributos - Semântica e Acessibilidade

**Cenário:** Você está criando um formulário de contato e precisa adicionar um campo de email. Você tem duas opções:

**Opção A:**
```html
<input type="text" name="email">
```

**Opção B:**
```html
<label for="email-usuario">Email:</label>
<input type="email" id="email-usuario" name="email" required>
```

**Perguntas:**
1. Qual opção é mais semântica e por quê?
2. Como cada opção afeta usuários que usam leitores de tela?
3. Qual é a diferença entre `type="text"` e `type="email"`? Por que isso importa?
4. Por que o atributo `for` no `<label>` é importante? Como ele se relaciona com o atributo `id`?
5. Qual é o impacto do atributo `required` na experiência do usuário e na validação do formulário?
6. Se você fosse um desenvolvedor que precisa manter esse código daqui a 6 meses, qual opção seria mais fácil de entender e modificar?

**Resposta esperada:** Explique a importância de usar atributos apropriados para melhorar semântica, acessibilidade e experiência do usuário.

---

### Reflexão 2: Case Insensitivity - Padrões e Consistência

**Cenário:** Você está trabalhando em um projeto com uma equipe de desenvolvedores. Um desenvolvedor escreve código assim:

```html
<HTML>
<HEAD>
<TITLE>Página</TITLE>
</HEAD>
<BODY>
<DIV CLASS="container">
<H1>Título</H1>
</DIV>
</BODY>
</HTML>
```

Outro desenvolvedor escreve assim:

```html
<html>
<head>
<title>Página</title>
</head>
<body>
<div class="container">
<h1>Título</h1>
</div>
</body>
</html>
```

**Perguntas:**
1. Ambas as abordagens funcionam no navegador. Por que então é importante padronizar?
2. Qual é o impacto de ter código inconsistente em um projeto grande com múltiplos desenvolvedores?
3. Como a inconsistência de case pode afetar ferramentas de validação e linting?
4. Por que XHTML (versão mais rigorosa do HTML) exige minúsculas? Qual é a relação com XML?
5. Se você fosse um desenvolvedor júnior lendo código pela primeira vez, qual versão seria mais fácil de entender?
6. Como você convenceria um desenvolvedor que prefere maiúsculas a adotar o padrão de minúsculas?

**Resposta esperada:** Explique a importância de seguir padrões de código e como isso afeta manutenibilidade, colaboração e qualidade do código.

---

### Reflexão 3: HTML Entities - Quando e Por Que Usar

**Cenário:** Você está criando uma página de e-commerce que exibe preços de produtos. Você precisa mostrar:

```html
<p>Preço: R$ 50,00</p>
<p>Desconto: 10%</p>
<p>Copyright © 2024</p>
<p>Para mais informações, visite: www.exemplo.com →</p>
```

**Perguntas:**
1. Você precisa usar entities para esses caracteres? Por quê?
2. Em que situações você **deve** usar entities? (Dica: pense em caracteres com significado especial no HTML)
3. Qual é a diferença entre usar UTF-8 diretamente vs. usar entities? Quando cada abordagem é mais apropriada?
4. Se você estiver trabalhando com conteúdo dinâmico (gerado por JavaScript), qual abordagem é mais prática?
5. Como entities podem ajudar em situações onde a codificação de caracteres não está configurada corretamente?
6. Qual é o impacto no tamanho do arquivo ao usar entities vs. caracteres UTF-8 diretamente?

**Resposta esperada:** Explique quando usar entities é necessário vs. opcional, e os trade-offs de cada abordagem.

---

### Reflexão 4: HTML Comments - Documentação e Manutenção

**Cenário:** Você está revisando código HTML de um projeto antigo. Você encontra três situações:

**Situação A: Código sem comentários**
```html
<div class="container">
    <div class="header">
        <h1>Título</h1>
    </div>
    <div class="content">
        <p>Texto</p>
    </div>
</div>
```

**Situação B: Código com comentários excessivos**
```html
<!-- Div container -->
<div class="container">
    <!-- Div header -->
    <div class="header">
        <!-- Título -->
        <h1>Título</h1>
    </div>
    <!-- Div content -->
    <div class="content">
        <!-- Parágrafo -->
        <p>Texto</p>
    </div>
</div>
```

**Situação C: Código com comentários apropriados**
```html
<!-- Container principal da página -->
<div class="container">
    <!-- Cabeçalho com logo e navegação -->
    <div class="header">
        <h1>Título</h1>
    </div>
    
    <!-- Conteúdo principal carregado dinamicamente via AJAX -->
    <div class="content">
        <p>Texto</p>
    </div>
</div>
```

**Perguntas:**
1. Qual abordagem é mais útil para um desenvolvedor que está vendo o código pela primeira vez?
2. Como comentários podem ajudar na manutenção de código a longo prazo?
3. Quais são os riscos de ter comentários desatualizados ou incorretos?
4. Por que comentários excessivos podem ser prejudiciais? (Dica: pense em ruído visual)
5. Como você decidiria o que merece um comentário e o que não merece?
6. Comentários HTML são visíveis no código-fonte da página. Que implicações isso tem para segurança e privacidade?

**Resposta esperada:** Explique o equilíbrio entre documentação útil e comentários desnecessários, e as melhores práticas para comentários HTML.

---

### Reflexão 5: Whitespaces - Formatação e Performance

**Cenário:** Você está debatendo com um colega sobre como formatar código HTML. Ele argumenta que código minificado (sem espaços) é melhor porque é menor e carrega mais rápido. Você argumenta que código formatado é melhor porque é mais fácil de manter.

**Perguntas:**
1. Qual é o impacto real de whitespaces no tamanho do arquivo HTML? (Dica: calcule a diferença percentual)
2. Em um projeto moderno, quando faz sentido minificar HTML? Quando não faz sentido?
3. Como ferramentas de build (como webpack, gulp) podem resolver esse conflito entre legibilidade e performance?
4. Por que a indentação é importante para a legibilidade do código? Como ela ajuda a visualizar a hierarquia?
5. Qual é o impacto de código mal formatado na produtividade de uma equipe de desenvolvimento?
6. Como você explicaria a um desenvolvedor júnior a importância de manter código bem formatado, mesmo que "funcione" de qualquer forma?

**Resposta esperada:** Explique o equilíbrio entre formatação para legibilidade e otimização para performance, e quando cada abordagem é apropriada.

---

### Reflexão 6: Integração de Conceitos - Criando Código Profissional

**Cenário:** Você precisa criar uma página HTML para um cliente. A página deve ser:
- Bem estruturada e semântica
- Fácil de manter
- Acessível
- Otimizada
- Seguir padrões da indústria

**Perguntas:**
1. Como você combinaria todos os conceitos aprendidos (tags, atributos, entities, comentários, whitespaces) para criar código profissional?
2. Qual é a ordem de prioridade ao escrever código HTML? (O que é mais importante: funcionar, ser legível, ser acessível, ser otimizado?)
3. Como você garantiria que seu código seja consistente em um projeto grande?
4. Que ferramentas ou processos você usaria para manter a qualidade do código HTML?
5. Como você ensinaria esses conceitos para um desenvolvedor júnior que está começando?
6. Qual é o impacto de código HTML bem escrito vs. mal escrito na experiência do usuário final?

**Resposta esperada:** Sintetize todos os conceitos aprendidos e explique como aplicá-los na prática para criar código HTML profissional e de alta qualidade.

---

## 📋 Checklist de Aprendizado

Após completar os exercícios, verifique se você consegue:

- [ ] Criar um arquivo HTML completo com estrutura básica correta
- [ ] Usar tags de abertura e fechamento corretamente
- [ ] Adicionar atributos globais e específicos às tags
- [ ] Entender a diferença entre atributos booleanos e atributos com valores
- [ ] Escrever código HTML sempre em minúsculas (convenção)
- [ ] Usar HTML entities para caracteres especiais quando necessário
- [ ] Adicionar comentários HTML para documentação e organização
- [ ] Entender como navegadores tratam whitespaces
- [ ] Usar indentação para melhorar legibilidade do código
- [ ] Usar a tag `<pre>` quando precisar preservar whitespaces
- [ ] Usar `&nbsp;` para espaços não separáveis quando apropriado
- [ ] Validar código HTML usando ferramentas apropriadas

---

## 🎓 Dicas para Resolução

### Dica 1: Validação
Use o [W3C Validator](https://validator.w3.org/) para verificar se seu código HTML está correto após cada exercício. Isso ajudará a identificar problemas de sintaxe e estrutura.

### Dica 2: Teste no Navegador
Sempre abra seus arquivos HTML no navegador para ver como ficam visualmente. Compare diferentes navegadores (Chrome, Firefox, Safari) para entender diferenças de renderização.

### Dica 3: Inspeção de Elementos
Use as DevTools do navegador (F12) para:
- Inspecionar elementos e ver seus atributos
- Ver como whitespaces são renderizados
- Entender a estrutura do DOM

### Dica 4: Experimente Diferentes Abordagens
Para os exercícios de entities e whitespaces, crie versões diferentes do mesmo código e compare os resultados. Isso ajudará a entender quando usar cada abordagem.

### Dica 5: Documentação
Consulte a [MDN Web Docs](https://developer.mozilla.org/pt-BR/docs/Web/HTML) quando tiver dúvidas sobre tags, atributos ou entities específicas.

---

## 📝 Instruções para Entrega

1. Crie uma pasta chamada `exercicios-aula-3` dentro da pasta da aula
2. Salve cada exercício em um arquivo separado:
   - `exercicio-1-primeiro-arquivo.html`
   - `exercicio-2-tags-atributos.html`
   - `exercicio-3-entities.html`
   - `exercicio-4-comentarios.html`
   - `exercicio-5-whitespaces.html` (crie 3 arquivos: versao-1.html, versao-2.html, versao-3.html)
   - `exercicio-6-correcao-case.html`
3. Crie um arquivo `reflexoes.md` com suas respostas às perguntas de reflexão
4. Revise seu código antes de considerar concluído
5. Valide todos os arquivos HTML no W3C Validator

**Boa sorte! Lembre-se: dominar esses fundamentos é essencial para escrever HTML profissional!** 🚀

