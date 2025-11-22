# Aula 7 - Exercícios e Reflexão

## 🎯 Exercícios Práticos

### Exercício 1: Criando uma Página com Marcação Semântica

Crie uma página HTML completa sobre "Tecnologias Web" usando elementos semânticos.

**Requisitos:**
- Use a estrutura completa de um documento HTML5
- Inclua um `<header>` com título principal e `<nav>` com menu de navegação
- Use `<main>` para o conteúdo principal
- Crie pelo menos 3 `<section>` com diferentes tópicos (ex: HTML, CSS, JavaScript)
- Cada seção deve ter um `<h2>` como título
- Inclua um `<aside>` com informações relacionadas
- Adicione um `<footer>` com informações de contato usando `<address>`
- Use pelo menos 2 elementos de citação (`<abbr>`, `<cite>`, `<q>` ou `<blockquote>`)

**Dica:** Use o arquivo `exemplo-04-layout-semantico.html` como referência.

---

### Exercício 2: Destacando Mudanças em Documentos

Crie uma página HTML que simule um documento com histórico de edições.

**Requisitos:**
- Crie um parágrafo sobre um preço de produto que foi alterado (use `<del>` e `<ins>`)
- Mostre uma lista de tarefas onde algumas foram concluídas (use `<s>` para tarefas não mais relevantes)
- Inclua um exemplo de texto que foi deletado e depois inserido com novos dados
- Use os atributos `datetime` e `cite` quando apropriado
- Adicione um parágrafo explicando as mudanças

**Exemplo de estrutura:**
```html
<p>
    Preço original: <del datetime="2024-01-01">R$ 150,00</del>
    Novo preço: <ins datetime="2024-01-15">R$ 120,00</ins>
</p>
```

---

### Exercício 3: Criando uma Página com Citações e Referências

Crie uma página HTML sobre um autor ou obra famosa, usando elementos de citação.

**Requisitos:**
- Use `<abbr>` para pelo menos 3 abreviações/acrônimos com atributo `title`
- Inclua pelo menos 2 citações usando `<blockquote>` com `<cite>` para a fonte
- Use `<q>` para uma citação curta inline
- Defina pelo menos 2 termos usando `<dfn>` na primeira ocorrência
- Adicione um `<address>` no rodapé com informações de contato do autor
- Use `<cite>` para referenciar obras mencionadas

**Exemplo de conteúdo:**
- Abreviações: HTML, CSS, JS, API, etc.
- Citações de autores famosos
- Definições de termos técnicos
- Informações de contato

---

### Exercício 4: Integrando CSS e JavaScript

Crie uma página HTML completa que demonstre os três métodos de adicionar CSS e JavaScript.

**Requisitos:**

**Parte 1 - CSS:**
- Use CSS inline para estilizar um elemento específico
- Use CSS interno (tag `<style>`) para estilizar elementos gerais
- Crie um arquivo CSS externo (`estilo.css`) e vincule-o
- Demonstre a ordem de precedência (inline > interno > externo)

**Parte 2 - JavaScript:**
- Use JavaScript inline em um botão (atributo `onclick`)
- Use JavaScript interno (tag `<script>`) para uma função
- Crie um arquivo JavaScript externo (`script.js`) e vincule-o
- Todos os scripts devem fazer algo visível (alerts, mudanças no DOM, etc.)

**Estrutura sugerida:**
```html
<!DOCTYPE html>
<html lang="pt-BR">
<head>
    <!-- CSS interno e externo aqui -->
</head>
<body>
    <!-- Conteúdo com CSS inline -->
    <!-- Botões para testar JavaScript -->
    <!-- Scripts aqui -->
</body>
</html>
```

---

## 🤔 Perguntas de Reflexão

### Reflexão 1: Semântica e Acessibilidade

**Cenário:** Você está criando um site de blog. Compare duas abordagens para estruturar a página:

**Abordagem A (não semântica):**
```html
<div class="cabecalho">
    <div class="titulo">Meu Blog</div>
    <div class="menu">...</div>
</div>
<div class="conteudo-principal">
    <div class="artigo">...</div>
</div>
<div class="barra-lateral">...</div>
<div class="rodape">...</div>
```

**Abordagem B (semântica):**
```html
<header>
    <h1>Meu Blog</h1>
    <nav>...</nav>
</header>
<main>
    <article>...</article>
</main>
<aside>...</aside>
<footer>...</footer>
```

**Perguntas:**
1. Como um leitor de tela interpretaria cada abordagem? Qual seria mais clara para uma pessoa cega?
2. Qual abordagem ajuda mais os mecanismos de busca a entender a estrutura da página? Por quê?
3. Se você precisasse modificar o layout da página daqui a 6 meses, qual código seria mais fácil de entender?
4. Em termos de manutenibilidade, qual abordagem facilita a identificação de problemas?
5. Como a semântica afeta a experiência de usuários que navegam apenas com teclado?

**Resposta esperada:** Explique as vantagens da abordagem semântica considerando acessibilidade, SEO e manutenibilidade.

---

### Reflexão 2: Elementos de Mudança e Contexto

**Cenário:** Você está criando um sistema de versionamento de documentos onde é importante mostrar o histórico de mudanças.

**Perguntas:**
1. Qual é a diferença prática entre `<del>` e `<s>`? Em que situações você usaria cada um?
2. Por que é importante usar os atributos `datetime` e `cite` nos elementos `<del>` e `<ins>`?
3. Como esses elementos ajudam usuários com deficiência visual a entender mudanças em documentos?
4. Qual é o impacto de usar esses elementos semânticos vs. apenas CSS para riscar texto?
5. Como você garantiria que as mudanças sejam claras tanto visualmente quanto semanticamente?

**Resposta esperada:** Explique a importância semântica desses elementos e como eles melhoram a acessibilidade.

---

### Reflexão 3: Citações e Referências

**Cenário:** Você está criando uma página acadêmica que precisa de várias citações e referências.

**Perguntas:**
1. Quando você usaria `<q>` vs. `<blockquote>`? Dê exemplos práticos.
2. Por que é importante usar `<cite>` para referenciar fontes? Qual é o impacto no SEO?
3. Como o elemento `<abbr>` melhora a acessibilidade? Dê um exemplo de uso eficaz.
4. Qual é a diferença entre usar `<dfn>` e apenas colocar um termo em negrito? Por que isso importa?
5. Como você organizaria uma página com muitas citações para manter a clareza e acessibilidade?

**Resposta esperada:** Explique como cada elemento de citação serve a um propósito específico e melhora a compreensão do conteúdo.

---

### Reflexão 4: CSS: Inline, Interno ou Externo?

**Cenário:** Você está desenvolvendo um site com 10 páginas. Analise quando usar cada método de CSS:

**Perguntas:**
1. Em que situações o CSS inline seria apropriado? Dê exemplos práticos.
2. Quando faz sentido usar CSS interno vs. externo? Quais são os trade-offs?
3. Se você tem estilos que são usados em todas as 10 páginas, qual método você escolheria? Por quê?
4. Como a escolha entre CSS interno e externo afeta a performance do site?
5. Qual é o impacto na manutenibilidade de usar CSS inline em muitos elementos?
6. Como você organizaria CSS em um projeto grande? (Dica: pense em múltiplos arquivos CSS externos)

**Resposta esperada:** Explique quando usar cada método considerando reutilização, manutenibilidade e performance.

---

### Reflexão 5: JavaScript e Separação de Responsabilidades

**Cenário:** Você está criando uma aplicação web interativa com várias funcionalidades.

**Perguntas:**
1. Por que é recomendado colocar scripts antes de `</body>` em vez de no `<head>`?
2. Qual é a diferença entre os atributos `defer` e `async` na tag `<script>`? Quando usar cada um?
3. Como a separação de JavaScript em arquivos externos melhora a organização do código?
4. Em termos de performance, qual é o impacto de ter múltiplos arquivos JavaScript vs. um único arquivo?
5. Como você garantiria que o JavaScript funcione mesmo se o usuário tiver JavaScript desabilitado? (Graceful degradation)
6. Qual é a importância de separar HTML (estrutura), CSS (estilo) e JavaScript (comportamento)? O que acontece quando misturamos tudo?

**Resposta esperada:** Explique as melhores práticas de inclusão de JavaScript e a importância da separação de responsabilidades.

---

### Reflexão 6: Layout Semântico e Estrutura

**Cenário:** Você precisa criar uma página de blog com múltiplos artigos, cada um com comentários.

**Perguntas:**
1. Como você estruturaria a página usando `<header>`, `<nav>`, `<main>`, `<article>`, `<section>`, `<aside>` e `<footer>`?
2. Por que só deve haver um elemento `<main>` por página? O que acontece se houver múltiplos?
3. Qual é a diferença entre `<section>` e `<article>`? Quando usar cada um?
4. Como você organizaria comentários dentro de um artigo? Eles seriam `<section>`, `<article>` ou outro elemento?
5. Se você tivesse uma barra lateral com widgets (pesquisa, categorias, anúncios), qual elemento semântico usaria?
6. Como a estrutura semântica facilita a criação de layouts responsivos com CSS?

**Resposta esperada:** Explique como estruturar uma página complexa usando elementos semânticos de forma apropriada.

---

## 📋 Checklist de Aprendizado

Após completar os exercícios, verifique se você consegue:

- [ ] Usar elementos semânticos de layout (`<header>`, `<nav>`, `<main>`, `<section>`, `<article>`, `<aside>`, `<footer>`)
- [ ] Aplicar elementos para destacar mudanças (`<del>`, `<s>`, `<ins>`) com atributos apropriados
- [ ] Usar elementos de citação e referência (`<abbr>`, `<cite>`, `<dfn>`, `<address>`, `<blockquote>`, `<q>`)
- [ ] Adicionar CSS inline, interno e externo corretamente
- [ ] Incluir JavaScript inline, interno e externo
- [ ] Entender a ordem de precedência de CSS
- [ ] Estruturar páginas web de forma semântica e acessível
- [ ] Separar responsabilidades entre HTML, CSS e JavaScript
- [ ] Compreender quando usar cada método de adicionar CSS e JavaScript

---

## 🎓 Dicas para Resolução

### Dica 1: Validação Semântica
Use o [W3C Validator](https://validator.w3.org/) para verificar se seu código HTML está correto e semântico.

### Dica 2: Teste de Acessibilidade
Use extensões de navegador como "WAVE" ou "axe DevTools" para verificar a acessibilidade de suas páginas.

### Dica 3: Inspeção de Elementos
Use as DevTools do navegador (F12) para inspecionar como os elementos semânticos são interpretados.

### Dica 4: Teste com Leitores de Tela
Experimente usar leitores de tela (como NVDA ou VoiceOver) para entender como a semântica afeta a experiência.

### Dica 5: Organização de Arquivos
Mantenha seus arquivos CSS e JavaScript organizados em pastas separadas (ex: `css/` e `js/`).

---

## 📝 Instruções para Entrega

1. Crie uma pasta chamada `exercicios-aula-7` dentro da pasta da aula
2. Salve cada exercício em um arquivo separado:
   - `exercicio-1-marcacao-semantica.html`
   - `exercicio-2-destacar-mudancas.html`
   - `exercicio-3-citacoes-referencias.html`
   - `exercicio-4-css-javascript.html`
   - `estilo.css` (para o exercício 4)
   - `script.js` (para o exercício 4)
3. Crie um arquivo `reflexoes.md` com suas respostas às perguntas de reflexão
4. Revise seu código antes de considerar concluído
5. Valide todos os arquivos HTML usando o W3C Validator

**Boa sorte! Lembre-se: a semântica é fundamental para criar páginas web acessíveis e bem estruturadas!** 🚀

---

## 💡 Desafio Extra (Opcional)

Crie uma página completa de portfólio pessoal que demonstre todos os conceitos aprendidos:

- Estrutura semântica completa
- Uso de elementos de citação e referência
- CSS externo bem organizado
- JavaScript externo para interatividade
- Acessibilidade completa
- Responsividade básica

Este desafio consolidará todo o conhecimento da aula!

