# Aula 5 - Exercícios e Reflexão

## 🎯 Exercícios Práticos

### Exercício 1: Agrupamento com `<div>` e `<span>`

Crie uma página HTML que demonstre o uso correto de `<div>` e `<span>`:

**Requisitos:**
1. Crie uma seção principal usando `<div>` que contenha:
   - Um título `<h2>`
   - Dois parágrafos
   - Uma imagem (use um placeholder se necessário)

2. Crie três `<div>` diferentes, cada uma representando um "card" de produto com:
   - Título do produto
   - Descrição
   - Preço

3. Dentro de um parágrafo, use `<span>` para destacar:
   - Uma palavra importante
   - Um número ou valor
   - Uma data

4. Crie uma `<div>` que agrupe informações de contato usando `<span>` para destacar partes específicas (ex: email, telefone)

**Desafio:** Identifique quando você deveria usar um elemento semântico (como `<section>`, `<article>`) ao invés de `<div>`. Reescreva pelo menos uma `<div>` usando um elemento semântico apropriado.

---

### Exercício 2: Atributos `id` e `class`

Crie uma página HTML demonstrando o uso correto de atributos `id` e `class`:

**Requisitos:**
1. Crie uma página com pelo menos 5 seções diferentes, cada uma com um `id` único:
   - `id="introducao"`
   - `id="sobre"`
   - `id="servicos"`
   - `id="contato"`
   - `id="rodape"`

2. Crie um menu de navegação no topo que use links âncora para navegar entre as seções (ex: `href="#sobre"`)

3. Crie pelo menos 3 elementos com a mesma `class` (ex: `class="destaque"`) e explique em comentários por que você usou `class` ao invés de `id`

4. Crie um elemento que tenha múltiplas classes (ex: `class="card destaque importante"`)

5. Use `id` para criar um link "Voltar ao topo" que leve ao início da página

**Validação:** Certifique-se de que cada `id` é único. Valide no W3C Validator.

---

### Exercício 3: Atributos `data-*` e `style`

Crie uma página HTML demonstrando o uso de atributos `data-*` e `style`:

**Requisitos:**
1. Crie uma lista de produtos usando `<div>` com atributos `data-*`:
   - `data-produto-id`
   - `data-categoria`
   - `data-preco`
   - `data-estoque`

2. Use o atributo `style` para:
   - Aplicar uma cor de fundo diferente em um elemento específico
   - Ajustar o tamanho da fonte de um parágrafo
   - Adicionar uma borda em um elemento

3. Em comentários HTML, explique:
   - Quando é apropriado usar `style` inline
   - Quando você deveria usar CSS externo ao invés de `style`
   - Para que servem os atributos `data-*` que você criou

**Desafio:** Crie um exemplo onde você usa `data-*` para armazenar informações que poderiam ser usadas por JavaScript (mesmo que você não escreva o JavaScript ainda).

---

### Exercício 4: Listas Ordenadas e Não Ordenadas

Crie uma página HTML com diferentes tipos de listas:

**Requisitos:**
1. Crie uma lista ordenada (`<ol>`) com instruções passo a passo para:
   - Fazer um café
   - Ou qualquer tarefa do dia a dia
   - Use pelo menos 5 passos

2. Crie uma lista não ordenada (`<ul>`) com:
   - Lista de compras (pelo menos 8 itens)
   - Lista de características de um produto
   - Menu de navegação com links

3. Experimente diferentes tipos de numeração na lista ordenada:
   - Números (padrão)
   - Letras maiúsculas (`type="A"`)
   - Algarismos romanos (`type="I"`)

4. Crie uma lista ordenada que comece em um número específico usando `start`

**Desafio:** Crie uma lista ordenada com instruções e dentro de cada item (`<li>`), adicione uma lista não ordenada com sub-itens relacionados.

---

### Exercício 5: Listas de Definição e Listas Aninhadas

Crie uma página HTML demonstrando listas de definição e listas aninhadas:

**Requisitos:**
1. Crie uma lista de definição (`<dl>`) com um glossário de termos técnicos:
   - Pelo menos 5 termos relacionados a HTML/web
   - Cada termo deve ter uma definição clara
   - Use `<dt>` para o termo e `<dd>` para a definição

2. Crie uma lista aninhada (lista dentro de lista) representando:
   - Um menu de restaurante com categorias e itens
   - Ou uma estrutura organizacional
   - Ou um índice de livro com capítulos e seções

3. Crie uma lista não ordenada com pelo menos 3 níveis de aninhamento:
   - Nível 1: Categorias principais
   - Nível 2: Subcategorias
   - Nível 3: Itens específicos

**Desafio:** Crie uma lista de definição onde um termo tem múltiplas definições e onde múltiplos termos compartilham uma definição.

---

### Exercício 6: Tabelas HTML Básicas

Crie uma página HTML com tabelas bem estruturadas:

**Requisitos:**
1. Crie uma tabela simples com:
   - Cabeçalho (`<th>`) com pelo menos 3 colunas
   - Pelo menos 5 linhas de dados (`<tr>` com `<td>`)
   - Dados sobre qualquer tema (ex: alunos e notas, produtos e preços, horários)

2. Crie uma tabela usando estrutura semântica completa:
   - `<caption>` com título da tabela
   - `<thead>` com cabeçalhos
   - `<tbody>` com dados
   - `<tfoot>` com totais ou resumo

3. Crie uma tabela que demonstre:
   - `colspan` para mesclar células horizontalmente
   - `rowspan` para mesclar células verticalmente

4. Crie uma tabela de horário semanal (dias da semana × horários)

**Validação:** Valide sua tabela no W3C Validator. Certifique-se de que todas as linhas têm o mesmo número de células (ou use colspan/rowspan corretamente).

---

### Exercício 7: Projeto Integrado - Página Completa

Crie uma página HTML completa sobre um tema de sua escolha usando TODOS os conceitos aprendidos nesta aula:

**Requisitos:**
- Estrutura básica completa (DOCTYPE, html, head, body)
- Meta tags essenciais
- Use `<div>` para agrupar seções principais
- Use `<span>` para destacar partes do texto
- Aplique `id` em seções principais
- Use `class` para estilizar grupos de elementos
- Inclua pelo menos um atributo `data-*`
- Crie um menu de navegação usando lista não ordenada com links âncora
- Use lista ordenada para instruções ou passos
- Use lista não ordenada para características ou itens
- Crie uma lista de definição (glossário)
- Inclua pelo menos uma lista aninhada
- Crie uma tabela com dados relevantes
- Use estrutura semântica na tabela (thead, tbody, tfoot)

**Validação:** Valide no W3C Validator e corrija todos os erros.

---

### Exercício 8: Análise e Correção de Código

Analise o seguinte código HTML e identifique TODOS os problemas:

```html
<!DOCTYPE html>
<html lang="pt-BR">
<head>
    <meta charset="UTF-8">
    <title>Análise de Código</title>
</head>
<body>
    <div id="cabecalho" id="header">
        <h1>Minha Página</h1>
    </div>
    
    <div class="destaque">
        <p>Este é um parágrafo com <span class="destaque">texto destacado</span>.</p>
    </div>
    
    <div class="destaque">
        <p>Outro parágrafo.</p>
    </div>
    
    <ol>
        <li>Item 1</li>
        <li>Item 2
            <ul>
                <li>Subitem</li>
            </ul>
        </li>
    </ol>
    
    <ul>
        <li>Item A</li>
        <ol>
            <li>Subitem numerado</li>
        </ol>
    </ul>
    
    <table>
        <tr>
            <th>Nome</th>
            <th>Idade</th>
        </tr>
        <tr>
            <td>João</td>
        </tr>
        <tr>
            <td>Maria</td>
            <td>30</td>
            <td>Extra</td>
        </tr>
    </table>
    
    <dl>
        <dd>Termo sem definição</dd>
        <dt>Definição sem termo</dt>
    </dl>
    
    <div data-produto="123" data produto preco="50">
        Produto
    </div>
</body>
</html>
```

**Tarefas:**
1. Liste todos os problemas encontrados (sintaxe, semântica, estrutura)
2. Explique por que cada problema é um erro
3. Explique o impacto de cada erro em:
   - Validação HTML
   - Acessibilidade
   - Funcionalidade
4. Reescreva o código corrigindo todos os problemas
5. Valide o código corrigido no W3C Validator

---

## 🤔 Perguntas de Reflexão

### Reflexão 1: `<div>` vs Elementos Semânticos

**Pergunta:** Quando você deve usar `<div>` e quando deve usar elementos semânticos como `<section>`, `<article>`, `<aside>`, `<header>`, `<footer>`, `<nav>`?

**Considere:**
- Qual é o impacto na acessibilidade quando você usa `<div>` ao invés de elementos semânticos?
- Como leitores de tela interpretam `<div>` vs `<section>`?
- Qual é o impacto no SEO quando você usa elementos semânticos?
- Em que situações `<div>` é realmente a melhor escolha?
- Como você decide qual elemento semântico usar quando há múltiplas opções?

**Tarefa:** Crie um exemplo onde você inicialmente usaria `<div>` mas deveria usar um elemento semântico. Explique sua escolha.

---

### Reflexão 2: `id` vs `class` - Quando Usar Cada Um?

**Pergunta:** Qual é a diferença fundamental entre `id` e `class`? Quando você deve usar cada um?

**Considere:**
- Por que `id` deve ser único enquanto `class` pode ser repetida?
- Qual é o impacto de ter múltiplos elementos com o mesmo `id`?
- Como navegadores e JavaScript lidam com `id` duplicados?
- Quando você precisa de um identificador único (`id`) vs quando precisa agrupar elementos (`class`)?
- Qual é o impacto na performance quando você usa muitos `id` vs muitas `class`?
- Como `id` e `class` se relacionam com CSS e JavaScript?

**Tarefa:** Crie exemplos práticos onde:
1. Você DEVE usar `id` (e explicar por quê)
2. Você DEVE usar `class` (e explicar por quê)
3. Você poderia usar ambos (e explicar quando cada um é apropriado)

---

### Reflexão 3: Atributos `data-*` e Separação de Responsabilidades

**Pergunta:** Qual é o propósito dos atributos `data-*`? Como eles se relacionam com a separação de responsabilidades entre HTML, CSS e JavaScript?

**Considere:**
- Por que armazenar dados em atributos `data-*` ao invés de hardcodar no JavaScript?
- Qual é a diferença entre usar `data-*` e usar `id` ou `class` para identificar elementos?
- Como atributos `data-*` melhoram a manutenibilidade do código?
- Qual é o impacto na acessibilidade quando você usa `data-*`?
- Como mecanismos de busca interpretam atributos `data-*`?
- Quando você NÃO deveria usar `data-*`?

**Tarefa:** Crie um exemplo onde você usa `data-*` para armazenar informações que seriam úteis para JavaScript, mas que não devem aparecer visualmente na página.

---

### Reflexão 4: Listas - Semântica e Acessibilidade

**Pergunta:** Qual é a importância semântica de usar os tipos corretos de listas (`<ol>`, `<ul>`, `<dl>`)? Como isso afeta acessibilidade e SEO?

**Considere:**
- Como leitores de tela anunciam diferentes tipos de listas?
- Qual é o impacto de usar `<ul>` quando você deveria usar `<ol>` (e vice-versa)?
- Por que a ordem importa em listas ordenadas para acessibilidade?
- Como mecanismos de busca interpretam listas ordenadas vs não ordenadas?
- Qual é a importância de usar `<dl>` para glossários ao invés de parágrafos?
- Como listas aninhadas afetam a navegação por leitores de tela?
- Qual é o impacto de usar muitos níveis de aninhamento?

**Tarefa:** Crie exemplos de conteúdo que poderiam ser apresentados como lista, mas que estão incorretamente marcados. Identifique o problema e corrija usando o tipo de lista apropriado.

---

### Reflexão 5: Tabelas - Quando e Como Usar

**Pergunta:** Quando você deve usar tabelas HTML? Qual é a diferença entre dados tabulares e layout visual?

**Considere:**
- Por que tabelas não devem ser usadas para layout de página?
- Qual é o impacto na acessibilidade quando você usa tabelas para layout?
- Como leitores de tela navegam por tabelas?
- Qual é a importância de usar `<thead>`, `<tbody>`, `<tfoot>` e `<caption>`?
- Como `colspan` e `rowspan` afetam a acessibilidade?
- Qual é o impacto no SEO quando você usa tabelas apropriadamente?
- Como dispositivos móveis lidam com tabelas grandes?
- Quando você deveria considerar alternativas a tabelas (ex: cards, listas)?

**Tarefa:** 
1. Crie um exemplo de dados que DEVERIAM ser apresentados em tabela
2. Crie um exemplo de layout que NÃO deveria usar tabela (e sugira alternativas)
3. Explique a diferença e por que cada abordagem é apropriada ou não

---

### Reflexão 6: Estruturação e Organização de Código

**Pergunta:** Como a escolha entre `<div>`, elementos semânticos, listas e tabelas afeta a organização e manutenibilidade do código HTML?

**Considere:**
- Como código bem estruturado facilita a manutenção?
- Qual é o impacto de usar elementos semânticos na legibilidade do código?
- Como outros desenvolvedores interpretam seu código quando você usa elementos apropriados?
- Qual é a relação entre estrutura HTML e CSS?
- Como estrutura bem organizada facilita a integração com JavaScript?
- Qual é o impacto na performance quando você usa estrutura apropriada?
- Como você decide entre múltiplas formas de estruturar o mesmo conteúdo?

**Tarefa:** Crie a mesma página de duas formas diferentes:
1. Usando apenas `<div>` e estrutura básica
2. Usando elementos semânticos, listas e tabelas apropriadas

Compare e explique:
- Qual é mais legível?
- Qual é mais acessível?
- Qual é mais fácil de manter?
- Qual é melhor para SEO?

---

## ✅ Checklist de Aprendizado

Antes de considerar esta aula completa, certifique-se de que você:

- [ ] Entende a diferença entre `<div>` (block-level) e `<span>` (inline)
- [ ] Sabe quando usar `<div>` vs elementos semânticos
- [ ] Compreende que `id` deve ser único e `class` pode ser repetida
- [ ] Sabe quando usar `id` vs `class`
- [ ] Entende o propósito dos atributos `data-*`
- [ ] Sabe quando usar `style` inline (e quando não usar)
- [ ] Consegue criar listas ordenadas (`<ol>`) corretamente
- [ ] Consegue criar listas não ordenadas (`<ul>`) corretamente
- [ ] Consegue criar listas de definição (`<dl>`) corretamente
- [ ] Sabe criar listas aninhadas mantendo estrutura correta
- [ ] Entende quando usar cada tipo de lista
- [ ] Consegue criar tabelas básicas com `<table>`, `<tr>`, `<td>`, `<th>`
- [ ] Sabe usar estrutura semântica de tabelas (`<thead>`, `<tbody>`, `<tfoot>`, `<caption>`)
- [ ] Consegue usar `colspan` e `rowspan` corretamente
- [ ] Entende quando usar tabelas (dados tabulares) e quando não usar (layout)
- [ ] Valida seu código HTML no W3C Validator
- [ ] Considera acessibilidade ao criar listas e tabelas
- [ ] Pensa em semântica ao escolher elementos HTML

---

## 📝 Instruções para Entrega

1. **Crie uma pasta** chamada `exercicios-aula-5` dentro do diretório desta aula

2. **Salve cada exercício** em um arquivo separado:
   - `exercicio-01-div-span.html`
   - `exercicio-02-id-class.html`
   - `exercicio-03-data-style.html`
   - `exercicio-04-listas-ordenadas.html`
   - `exercicio-05-listas-definicao-aninhadas.html`
   - `exercicio-06-tabelas.html`
   - `exercicio-07-projeto-integrado.html`
   - `exercicio-08-analise-correcao.html`

3. **Para as perguntas de reflexão**, crie um arquivo `reflexoes.md` onde você responde cada pergunta de forma detalhada

4. **Valide todos os arquivos HTML** no [W3C Validator](https://validator.w3.org/)

5. **Revise seu código** usando o checklist acima

---

## 🎯 Objetivos de Aprendizado

Ao completar estes exercícios e reflexões, você deve ser capaz de:

✅ Agrupar elementos usando `<div>` e `<span>` apropriadamente  
✅ Usar atributos `id` e `class` corretamente  
✅ Aplicar atributos `data-*` e `style` quando apropriado  
✅ Criar listas ordenadas, não ordenadas e de definição  
✅ Criar listas aninhadas mantendo estrutura correta  
✅ Criar tabelas HTML bem estruturadas e semânticas  
✅ Entender quando usar cada tipo de elemento  
✅ Considerar acessibilidade e semântica em suas escolhas  
✅ Validar código HTML  
✅ Pensar criticamente sobre estruturação de conteúdo  

---

**Boa sorte com os exercícios! Lembre-se: a prática é essencial para dominar HTML!** 🚀

