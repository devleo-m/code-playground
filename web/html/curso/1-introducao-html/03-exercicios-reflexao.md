# Aula 1 - Exercícios e Reflexão

## 🎯 Exercícios Práticos

### Exercício 1: Criando sua Primeira Página HTML

Crie um arquivo HTML básico com a seguinte estrutura:

**Requisitos:**
- Use a estrutura básica de um documento HTML5
- Adicione um título na tag `<title>`: "Minha Primeira Página"
- No `<body>`, adicione:
  - Um título principal (`<h1>`) com seu nome
  - Um parágrafo (`<p>`) apresentando-se brevemente
  - Uma lista não ordenada (`<ul>`) com 3 coisas que você gosta de fazer
  - Um link (`<a>`) para um site que você gosta

**Dica:** Use o arquivo `exemplo-01-estrutura-basica.html` como referência.

---

### Exercício 2: Analisando Código HTML

Analise o seguinte código HTML e identifique **pelo menos 5 problemas** ou melhorias que podem ser feitas:

```html
<!DOCTYPE html>
<html>
<head>
<title>Minha Página
</head>
<body>
<h1>Bem-vindo</h1>
<p>Esta é minha página web.
<h2>Sobre Mim</h2>
<p>Eu gosto de programar e estudar HTML.
<img src="foto.jpg">
<a href="https://www.google.com">Clique aqui</a>
<div>Contato: email@exemplo.com</div>
</body>
</html>
```

**Tarefa:**
1. Liste os problemas encontrados
2. Reescreva o código corrigindo todos os problemas identificados
3. Explique por que cada correção é importante

---

### Exercício 3: Estruturando Conteúdo

Você precisa criar uma página HTML para uma receita de bolo. Organize o seguinte conteúdo usando as tags HTML apropriadas:

**Conteúdo:**
- Título: "Bolo de Chocolate"
- Subtítulo: "Ingredientes"
- Lista: 2 xícaras de farinha, 1 xícara de açúcar, 3 ovos, 1 xícara de leite, 1/2 xícara de óleo, 1 colher de sopa de fermento
- Subtítulo: "Modo de Preparo"
- Parágrafo: "Misture todos os ingredientes secos em uma tigela. Em outra tigela, bata os ovos e adicione o leite e o óleo. Combine as misturas e adicione o fermento. Asse em forno pré-aquecido a 180°C por 40 minutos."
- Rodapé: "Receita da vovó - 2024"

**Requisitos:**
- Use a estrutura completa de um documento HTML5
- Use tags semânticas quando apropriado
- Organize o conteúdo de forma hierárquica e lógica

---

### Exercício 4: Criando uma Página de Apresentação

Crie uma página HTML completa que sirva como sua **carta de apresentação pessoal**.

**Estrutura necessária:**
1. Cabeçalho com seu nome como título principal
2. Seção "Sobre Mim" com um parágrafo de apresentação
3. Seção "Habilidades" com uma lista ordenada das suas principais habilidades
4. Seção "Interesses" com uma lista não ordenada dos seus interesses
5. Seção "Contato" com links para suas redes sociais ou email
6. Rodapé com informação de copyright

**Requisitos adicionais:**
- Use tags semânticas HTML5 (`<header>`, `<main>`, `<section>`, `<footer>`)
- Adicione o atributo `lang="pt-BR"` na tag `<html>`
- Inclua meta tags apropriadas no `<head>`
- Todos os links devem ter o atributo `target="_blank"` para abrir em nova aba

---

## 🤔 Perguntas de Reflexão

### Reflexão 1: Semântica e Acessibilidade

**Cenário:** Você está criando uma página web para um site de notícias. Você tem duas opções para criar o título principal:

**Opção A:**
```html
<div class="titulo-principal">Últimas Notícias</div>
```

**Opção B:**
```html
<h1>Últimas Notícias</h1>
```

**Perguntas:**
1. Qual opção é mais semântica e por quê?
2. Como cada opção afeta leitores de tela (ferramentas de acessibilidade)?
3. Qual é o impacto de cada opção no SEO (mecanismos de busca)?
4. Se você fosse um desenvolvedor que precisa manter esse código daqui a 6 meses, qual opção seria mais fácil de entender?

**Resposta esperada:** Explique sua escolha considerando acessibilidade, SEO e manutenibilidade do código.

---

### Reflexão 2: Estrutura Hierárquica e SEO

**Cenário:** Você está criando uma página de blog sobre "Receitas Veganas". Analise a seguinte estrutura de títulos:

```html
<h1>Receitas Veganas</h1>
<h3>Bolo de Chocolate Vegano</h3>
<h2>Ingredientes</h2>
<h4>Dicas de Preparo</h4>
<h2>Modo de Preparo</h2>
```

**Perguntas:**
1. Identifique os problemas na hierarquia dos títulos. O que está errado?
2. Como essa hierarquia incorreta pode afetar a experiência do usuário?
3. Qual é o impacto dessa estrutura no SEO? Por que os mecanismos de busca se importam com a hierarquia?
4. Como você reorganizaria essa estrutura para seguir as melhores práticas?
5. Por que é importante manter uma hierarquia lógica de títulos (h1 → h2 → h3, sem pular níveis)?

**Resposta esperada:** Explique a hierarquia correta e o impacto de uma estrutura bem organizada.

---

### Reflexão 3: Acessibilidade e Inclusão Digital

**Cenário:** Você criou uma página com várias imagens, mas esqueceu de adicionar o atributo `alt` em todas elas:

```html
<img src="produto1.jpg">
<img src="produto2.jpg" alt="Produto 2">
<img src="produto3.jpg">
```

**Perguntas:**
1. O que acontece quando uma pessoa cega acessa essa página usando um leitor de tela?
2. Por que o atributo `alt` é considerado essencial para acessibilidade web?
3. Além de pessoas cegas, que outros grupos de usuários se beneficiam do atributo `alt`?
4. Qual é a diferença entre um `alt` vazio (`alt=""`) e a ausência do atributo `alt`?
5. Como você escreveria bons textos alternativos? Dê exemplos de `alt` bons e ruins para uma imagem de um produto.

**Resposta esperada:** Explique a importância do `alt` e como escrever descrições eficazes.

---

### Reflexão 4: Estrutura Semântica vs. Divs Genéricas

**Cenário:** Dois desenvolvedores criaram a mesma página, mas com abordagens diferentes:

**Desenvolvedor A:**
```html
<div class="header">
    <div class="titulo">Meu Site</div>
</div>
<div class="conteudo">
    <div class="artigo">Artigo sobre HTML...</div>
</div>
<div class="rodape">Copyright 2024</div>
```

**Desenvolvedor B:**
```html
<header>
    <h1>Meu Site</h1>
</header>
<main>
    <article>Artigo sobre HTML...</article>
</main>
<footer>Copyright 2024</footer>
```

**Perguntas:**
1. Qual abordagem é mais semântica? Por quê?
2. Como cada abordagem afeta a navegação por leitores de tela?
3. Qual abordagem é mais fácil de manter e entender para outros desenvolvedores?
4. Os mecanismos de busca (Google, Bing) conseguem entender melhor qual estrutura? Por quê?
5. Em termos de performance, há alguma diferença? (Dica: pense em como os navegadores processam o código)
6. Se você precisasse adicionar navegação por teclado (acessibilidade), qual estrutura facilitaria isso?

**Resposta esperada:** Compare as duas abordagens considerando semântica, acessibilidade, SEO e manutenibilidade.

---

### Reflexão 5: Responsividade e Dispositivos Móveis

**Cenário:** Você criou uma página HTML que funciona perfeitamente no seu computador, mas quando você acessa no celular, o texto fica muito pequeno e difícil de ler.

**Perguntas:**
1. Qual meta tag é essencial para que uma página funcione bem em dispositivos móveis?
2. Por que é importante pensar em dispositivos móveis desde o início do desenvolvimento?
3. Como a estrutura HTML pode influenciar a responsividade de uma página? (Dica: pense em como elementos semânticos podem ser reorganizados)
4. Além do HTML, que outras tecnologias são necessárias para criar uma experiência mobile adequada?
5. Qual é o impacto de uma página não responsiva na experiência do usuário e no SEO?

**Resposta esperada:** Explique a importância da meta tag viewport e como HTML se relaciona com design responsivo.

---

## 📋 Checklist de Aprendizado

Após completar os exercícios, verifique se você consegue:

- [ ] Criar a estrutura básica de um documento HTML5
- [ ] Usar tags básicas: `<h1>` a `<h6>`, `<p>`, `<ul>`, `<ol>`, `<li>`, `<a>`, `<img>`
- [ ] Entender a diferença entre tags de abertura e fechamento
- [ ] Adicionar atributos aos elementos HTML
- [ ] Criar links internos e externos
- [ ] Inserir imagens com atributo `alt`
- [ ] Organizar conteúdo de forma hierárquica
- [ ] Usar tags semânticas quando apropriado
- [ ] Entender a importância da acessibilidade
- [ ] Compreender o impacto do HTML no SEO

---

## 🎓 Dicas para Resolução

### Dica 1: Validação
Use o [W3C Validator](https://validator.w3.org/) para verificar se seu código HTML está correto após cada exercício.

### Dica 2: Teste no Navegador
Sempre abra seus arquivos HTML no navegador para ver como ficam visualmente.

### Dica 3: Inspeção de Elementos
Use as DevTools do navegador (F12) para inspecionar elementos e entender a estrutura.

### Dica 4: Acessibilidade
Teste suas páginas com leitores de tela ou extensões de acessibilidade para entender o impacto das suas escolhas.

---

## 📝 Instruções para Entrega

1. Crie uma pasta chamada `exercicios-aula-1` dentro da pasta da aula
2. Salve cada exercício em um arquivo separado:
   - `exercicio-1-minha-primeira-pagina.html`
   - `exercicio-2-analise-codigo.html` (código corrigido)
   - `exercicio-3-receita-bolo.html`
   - `exercicio-4-apresentacao-pessoal.html`
3. Crie um arquivo `reflexoes.md` com suas respostas às perguntas de reflexão
4. Revise seu código antes de considerar concluído

**Boa sorte! Lembre-se: a prática é essencial para aprender HTML!** 🚀

