# Aula 5: Agrupamento de Texto, Atributos e Listas - Conteúdo Principal

## 📝 Revisão da Aula Anterior

Antes de começarmos, vamos relembrar os conceitos fundamentais que você já aprendeu:

- **Estrutura básica** de um documento HTML (DOCTYPE, html, head, body)
- **Meta tags** e sua importância para SEO e acessibilidade
- **Títulos** (h1-h6) e hierarquia semântica
- **Parágrafos** e formatação de texto
- **Links** e navegação entre páginas
- **Elementos semânticos** básicos (header, nav, main, section, footer)

Agora vamos aprender a agrupar elementos, usar atributos para identificar e estilizar conteúdo, e criar listas e tabelas estruturadas!

---

## 🧩 Agrupamento de Texto: `<div>` e `<span>`

### O Elemento `<div>`: Container de Bloco

O elemento `<div>` (divisão) é um container genérico de nível de bloco usado para agrupar outros elementos HTML. Ele não possui significado semântico próprio, mas é extremamente útil para estruturação e estilização.

**Características principais:**
- É um elemento **block-level** (nível de bloco)
- Cria uma quebra de linha antes e depois
- Ocupa toda a largura disponível por padrão
- Não possui significado semântico inerente
- É usado principalmente para agrupamento estrutural e estilização com CSS

**Sintaxe básica:**
```html
<div>
    <!-- Conteúdo agrupado aqui -->
</div>
```

**Exemplo prático:**
```html
<div>
    <h2>Título da Seção</h2>
    <p>Conteúdo da seção agrupado em uma div.</p>
    <p>Outro parágrafo na mesma div.</p>
</div>
```

**Quando usar `<div>`:**
- Para agrupar elementos relacionados visualmente
- Quando você precisa de um container para aplicar CSS
- Para criar layouts e estruturas complexas
- Quando não há um elemento semântico mais apropriado

**Quando NÃO usar `<div>`:**
- Se existe um elemento semântico mais apropriado (section, article, aside, etc.)
- Para agrupar texto inline (use `<span>`)
- Quando o agrupamento tem significado semântico (prefira elementos semânticos)

### O Elemento `<span>`: Container Inline

O elemento `<span>` é um container genérico de nível inline usado para agrupar texto ou elementos inline para fins de estilização ou manipulação.

**Características principais:**
- É um elemento **inline** (em linha)
- Não cria quebra de linha
- Ocupa apenas o espaço necessário ao conteúdo
- Flui naturalmente com o texto ao redor
- Não possui significado semântico inerente

**Sintaxe básica:**
```html
<span>Texto ou elementos inline</span>
```

**Exemplo prático:**
```html
<p>Este é um parágrafo com uma <span>palavra destacada</span> no meio do texto.</p>
```

**Quando usar `<span>`:**
- Para estilizar uma parte específica de um texto
- Para agrupar elementos inline relacionados
- Quando você precisa aplicar CSS ou JavaScript a uma porção de texto
- Para marcar texto sem alterar o fluxo do documento

**Quando NÃO usar `<span>`:**
- Para agrupar elementos block-level (use `<div>`)
- Quando existe um elemento semântico mais apropriado (strong, em, mark, etc.)
- Para criar estrutura de página (use elementos semânticos)

### Diferenças entre `<div>` e `<span>`

| Característica | `<div>` | `<span>` |
|----------------|---------|----------|
| Tipo | Block-level | Inline |
| Quebra de linha | Sim | Não |
| Largura | 100% (por padrão) | Conteúdo |
| Uso principal | Estrutura e layout | Texto e elementos inline |
| Exemplo | Seções, containers | Palavras, frases |

**Exemplo comparativo:**
```html
<!-- div: cria um bloco separado -->
<div>
    <p>Parágrafo 1</p>
    <p>Parágrafo 2</p>
</div>

<!-- span: dentro do fluxo do texto -->
<p>Este é um parágrafo com <span>texto destacado</span> inline.</p>
```

---

## 🏷️ Atributos Padrão em HTML

Atributos padrão são propriedades que podem ser aplicadas a quase todos os elementos HTML para fornecer informações adicionais, identificação, estilização ou funcionalidade.

### O Atributo `id`: Identificador Único

O atributo `id` fornece um identificador único para um elemento dentro de um documento HTML. Cada `id` deve ser único em todo o documento.

**Características:**
- Deve ser único em todo o documento
- Não pode conter espaços
- É case-sensitive (diferencia maiúsculas de minúsculas)
- Deve começar com uma letra (não número)
- Pode conter letras, números, hífens e underscores

**Sintaxe:**
```html
<elemento id="identificador-unico">Conteúdo</elemento>
```

**Exemplos práticos:**
```html
<div id="cabecalho">
    <h1>Título Principal</h1>
</div>

<section id="sobre-nos">
    <h2>Sobre Nós</h2>
    <p>Conteúdo da seção...</p>
</section>

<p id="paragrafo-destaque">Este parágrafo tem um ID único.</p>
```

**Usos do `id`:**
- **CSS**: Selecionar e estilizar elementos específicos
- **JavaScript**: Manipular elementos específicos
- **Navegação**: Criar links âncora (ex: `#secao1`)
- **Acessibilidade**: Associar labels a inputs em formulários

**Exemplo com CSS:**
```html
<style>
    #cabecalho {
        background-color: #333;
        color: white;
        padding: 20px;
    }
</style>
<div id="cabecalho">Cabeçalho estilizado</div>
```

**Exemplo com JavaScript:**
```html
<div id="mensagem">Olá!</div>
<script>
    document.getElementById('mensagem').textContent = 'Olá, Mundo!';
</script>
```

**Exemplo com navegação:**
```html
<a href="#sobre-nos">Ir para Sobre Nós</a>
<!-- ... -->
<section id="sobre-nos">
    <h2>Sobre Nós</h2>
</section>
```

### O Atributo `class`: Agrupamento por Classes

O atributo `class` permite especificar uma ou mais classes para um elemento HTML. Diferente do `id`, múltiplos elementos podem compartilhar a mesma classe.

**Características:**
- Múltiplos elementos podem ter a mesma classe
- Um elemento pode ter múltiplas classes (separadas por espaço)
- Não precisa ser único no documento
- É case-sensitive
- Pode conter letras, números, hífens e underscores

**Sintaxe:**
```html
<elemento class="nome-classe">Conteúdo</elemento>
<elemento class="classe1 classe2 classe3">Conteúdo</elemento>
```

**Exemplos práticos:**
```html
<p class="destaque">Parágrafo destacado</p>
<p class="destaque">Outro parágrafo destacado</p>

<div class="card">
    <h2 class="titulo-card">Título</h2>
    <p class="texto-card">Conteúdo</p>
</div>

<button class="btn btn-primary btn-large">Clique Aqui</button>
```

**Usos do `class`:**
- **CSS**: Aplicar estilos a grupos de elementos
- **JavaScript**: Selecionar e manipular grupos de elementos
- **Organização**: Categorizar elementos por função ou estilo
- **Frameworks**: Usar classes de frameworks CSS (Bootstrap, Tailwind, etc.)

**Exemplo com CSS:**
```html
<style>
    .destaque {
        background-color: yellow;
        font-weight: bold;
    }
    
    .card {
        border: 1px solid #ccc;
        padding: 20px;
        margin: 10px;
    }
</style>

<p class="destaque">Texto destacado</p>
<div class="card">Conteúdo do card</div>
```

**Múltiplas classes:**
```html
<div class="container card destaque">
    <!-- Este elemento tem três classes -->
</div>
```

### Atributos `data-*`: Dados Customizados

Atributos `data-*` permitem armazenar informações customizadas diretamente em elementos HTML. Esses dados são privados à página e não são processados por navegadores ou mecanismos de busca.

**Características:**
- Sempre começam com `data-`
- Seguidos por um nome em minúsculas (pode usar hífens)
- Podem conter qualquer valor
- São acessíveis via JavaScript
- Não afetam a renderização visual

**Sintaxe:**
```html
<elemento data-nome="valor">Conteúdo</elemento>
<elemento data-usuario-id="123" data-status="ativo">Conteúdo</elemento>
```

**Exemplos práticos:**
```html
<div data-produto-id="456" data-categoria="eletronicos" data-preco="299.99">
    Produto: Smartphone
</div>

<button data-acao="salvar" data-form-id="formulario-contato">
    Salvar
</button>

<span data-tooltip="Informação adicional">Passe o mouse aqui</span>
```

**Usos dos atributos `data-*`:**
- **JavaScript**: Armazenar dados para manipulação via JS
- **Frameworks**: Integração com frameworks JavaScript
- **Testes**: Identificadores para testes automatizados
- **Configuração**: Armazenar configurações específicas de elementos

**Exemplo com JavaScript:**
```html
<div data-usuario="joao" data-idade="25">João Silva</div>
<script>
    const div = document.querySelector('[data-usuario="joao"]');
    const idade = div.dataset.idade; // "25"
    console.log(idade);
</script>
```

**Conversão de nomes:**
- HTML: `data-usuario-id` → JavaScript: `dataset.usuarioId` (camelCase)
- HTML: `data-produto-preco` → JavaScript: `dataset.produtoPreco`

### O Atributo `style`: Estilização Inline

O atributo `style` permite aplicar CSS diretamente a um elemento HTML. Embora funcional, seu uso deve ser limitado.

**Características:**
- Aplica estilos CSS diretamente no elemento
- Sobrescreve estilos externos e internos
- Deve conter CSS válido
- Não é recomendado para uso extensivo (prefira CSS externo)

**Sintaxe:**
```html
<elemento style="propriedade: valor; propriedade2: valor2;">
    Conteúdo
</elemento>
```

**Exemplos práticos:**
```html
<p style="color: blue; font-size: 18px;">Texto azul e grande</p>

<div style="background-color: #f0f0f0; padding: 20px; border: 1px solid #ccc;">
    Container estilizado
</div>

<span style="color: red; font-weight: bold;">Texto vermelho em negrito</span>
```

**Quando usar `style`:**
- Estilos únicos e específicos de um elemento
- Prototipagem rápida
- Estilos dinâmicos gerados por JavaScript
- Override temporário de estilos

**Quando NÃO usar `style`:**
- Para estilos reutilizáveis (use classes CSS)
- Para estilos de múltiplos elementos (use CSS externo)
- Para manter separação de responsabilidades
- Em produção (prefira CSS organizado)

**Boas práticas:**
```html
<!-- ❌ Ruim: estilo inline extenso -->
<div style="background-color: #fff; padding: 20px; margin: 10px; border: 1px solid #000; border-radius: 5px;">
    Conteúdo
</div>

<!-- ✅ Bom: usar classe CSS -->
<div class="card">
    Conteúdo
</div>
```

---

## 📋 Listas em HTML

Listas são elementos HTML usados para apresentar informações de forma estruturada e organizada. HTML oferece três tipos principais de listas.

### Listas Ordenadas (`<ol>`): Itens Numerados

Listas ordenadas são usadas quando a ordem dos itens é importante. Os itens são numerados automaticamente.

**Estrutura:**
```html
<ol>
    <li>Item 1</li>
    <li>Item 2</li>
    <li>Item 3</li>
</ol>
```

**Elementos:**
- `<ol>`: Container da lista ordenada (ordered list)
- `<li>`: Item da lista (list item)

**Atributos do `<ol>`:**
- `type`: Tipo de numeração (`1`, `A`, `a`, `I`, `i`)
- `start`: Número inicial da lista
- `reversed`: Inverte a ordem da numeração

**Exemplos práticos:**
```html
<!-- Lista numerada padrão -->
<ol>
    <li>Primeiro passo</li>
    <li>Segundo passo</li>
    <li>Terceiro passo</li>
</ol>

<!-- Lista com letras maiúsculas -->
<ol type="A">
    <li>Opção A</li>
    <li>Opção B</li>
    <li>Opção C</li>
</ol>

<!-- Lista com algarismos romanos -->
<ol type="I">
    <li>Capítulo I</li>
    <li>Capítulo II</li>
    <li>Capítulo III</li>
</ol>

<!-- Lista começando em número específico -->
<ol start="5">
    <li>Item 5</li>
    <li>Item 6</li>
    <li>Item 7</li>
</ol>
```

**Quando usar listas ordenadas:**
- Instruções passo a passo
- Rankings e classificações
- Sequências cronológicas
- Qualquer conteúdo onde a ordem importa

### Listas Não Ordenadas (`<ul>`): Itens com Marcadores

Listas não ordenadas são usadas quando a ordem dos itens não é importante. Os itens são marcados com bullets (pontos, círculos, quadrados).

**Estrutura:**
```html
<ul>
    <li>Item 1</li>
    <li>Item 2</li>
    <li>Item 3</li>
</ul>
```

**Elementos:**
- `<ul>`: Container da lista não ordenada (unordered list)
- `<li>`: Item da lista (list item)

**Atributos do `<ul>`:**
- `type`: Tipo de marcador (`disc`, `circle`, `square`) - obsoleto em HTML5, use CSS

**Exemplos práticos:**
```html
<!-- Lista com bullets padrão -->
<ul>
    <li>Maçã</li>
    <li>Banana</li>
    <li>Laranja</li>
</ul>

<!-- Lista de navegação -->
<ul>
    <li><a href="#home">Home</a></li>
    <li><a href="#sobre">Sobre</a></li>
    <li><a href="#contato">Contato</a></li>
</ul>

<!-- Lista com elementos complexos -->
<ul>
    <li>
        <h3>Título do Item</h3>
        <p>Descrição do item</p>
    </li>
    <li>
        <h3>Outro Título</h3>
        <p>Outra descrição</p>
    </li>
</ul>
```

**Quando usar listas não ordenadas:**
- Listas de características
- Menus de navegação
- Listas de itens sem ordem específica
- Qualquer conteúdo onde a ordem não importa

### Listas de Definição (`<dl>`): Termos e Definições

Listas de definição são usadas para apresentar termos e suas definições correspondentes, como em um glossário ou dicionário.

**Estrutura:**
```html
<dl>
    <dt>Termo</dt>
    <dd>Definição do termo</dd>
    <dt>Outro Termo</dt>
    <dd>Definição do outro termo</dd>
</dl>
```

**Elementos:**
- `<dl>`: Container da lista de definição (definition list)
- `<dt>`: Termo a ser definido (definition term)
- `<dd>`: Definição do termo (definition description)

**Exemplos práticos:**
```html
<!-- Glossário básico -->
<dl>
    <dt>HTML</dt>
    <dd>HyperText Markup Language - Linguagem de marcação para web</dd>
    
    <dt>CSS</dt>
    <dd>Cascading Style Sheets - Linguagem de estilização</dd>
    
    <dt>JavaScript</dt>
    <dd>Linguagem de programação para web</dd>
</dl>

<!-- Múltiplas definições para um termo -->
<dl>
    <dt>Navegador</dt>
    <dd>Software para acessar a internet</dd>
    <dd>Programa que interpreta HTML e CSS</dd>
</dl>

<!-- Múltiplos termos para uma definição -->
<dl>
    <dt>HTML</dt>
    <dt>HyperText Markup Language</dt>
    <dd>Linguagem de marcação para estruturar conteúdo web</dd>
</dl>
```

**Quando usar listas de definição:**
- Glossários e dicionários
- Listas de FAQ (perguntas e respostas)
- Metadados e informações estruturadas
- Qualquer conteúdo termo-definição

### Listas Aninhadas: Listas Dentro de Listas

Listas aninhadas permitem criar estruturas hierárquicas, colocando uma lista dentro de outra lista.

**Estrutura:**
```html
<ul>
    <li>Item 1
        <ul>
            <li>Subitem 1.1</li>
            <li>Subitem 1.2</li>
        </ul>
    </li>
    <li>Item 2</li>
</ul>
```

**Exemplos práticos:**
```html
<!-- Lista não ordenada aninhada -->
<ul>
    <li>Frutas
        <ul>
            <li>Maçã</li>
            <li>Banana</li>
            <li>Laranja</li>
        </ul>
    </li>
    <li>Vegetais
        <ul>
            <li>Cenoura</li>
            <li>Brócolis</li>
        </ul>
    </li>
</ul>

<!-- Lista ordenada aninhada -->
<ol>
    <li>Preparação
        <ol>
            <li>Lavar os ingredientes</li>
            <li>Cortar em pedaços</li>
        </ol>
    </li>
    <li>Cozinhar
        <ol>
            <li>Aquecer a panela</li>
            <li>Adicionar ingredientes</li>
        </ol>
    </li>
</ol>

<!-- Lista mista (ordenada e não ordenada) -->
<ol>
    <li>Capítulo 1
        <ul>
            <li>Seção 1.1</li>
            <li>Seção 1.2</li>
        </ul>
    </li>
    <li>Capítulo 2
        <ul>
            <li>Seção 2.1</li>
        </ul>
    </li>
</ol>
```

**Boas práticas para listas aninhadas:**
- Mantenha a indentação clara
- Não aninhe mais de 3-4 níveis
- Use tipos de lista apropriados para cada nível
- Considere acessibilidade (leitores de tela)

---

## 📊 Tabelas em HTML

Tabelas HTML são usadas para apresentar dados tabulares de forma estruturada em linhas e colunas.

### Estrutura Básica de Tabelas

**Elementos fundamentais:**
- `<table>`: Container principal da tabela
- `<tr>`: Linha da tabela (table row)
- `<td>`: Célula de dados (table data)
- `<th>`: Célula de cabeçalho (table header)

**Estrutura mínima:**
```html
<table>
    <tr>
        <th>Cabeçalho 1</th>
        <th>Cabeçalho 2</th>
    </tr>
    <tr>
        <td>Dado 1</td>
        <td>Dado 2</td>
    </tr>
</table>
```

**Exemplo prático básico:**
```html
<table>
    <tr>
        <th>Nome</th>
        <th>Idade</th>
        <th>Cidade</th>
    </tr>
    <tr>
        <td>João</td>
        <td>25</td>
        <td>São Paulo</td>
    </tr>
    <tr>
        <td>Maria</td>
        <td>30</td>
        <td>Rio de Janeiro</td>
    </tr>
</table>
```

### Estrutura Semântica de Tabelas

Para melhor organização e acessibilidade, use elementos semânticos:

**Elementos semânticos:**
- `<thead>`: Cabeçalho da tabela (table head)
- `<tbody>`: Corpo da tabela (table body)
- `<tfoot>`: Rodapé da tabela (table foot)
- `<caption>`: Legenda/título da tabela

**Estrutura completa:**
```html
<table>
    <caption>Título da Tabela</caption>
    <thead>
        <tr>
            <th>Coluna 1</th>
            <th>Coluna 2</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>Dado 1</td>
            <td>Dado 2</td>
        </tr>
    </tbody>
    <tfoot>
        <tr>
            <td>Total 1</td>
            <td>Total 2</td>
        </tr>
    </tfoot>
</table>
```

**Exemplo prático completo:**
```html
<table>
    <caption>Vendas do Mês</caption>
    <thead>
        <tr>
            <th>Produto</th>
            <th>Quantidade</th>
            <th>Preço Unitário</th>
            <th>Total</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>Notebook</td>
            <td>5</td>
            <td>R$ 2.500,00</td>
            <td>R$ 12.500,00</td>
        </tr>
        <tr>
            <td>Mouse</td>
            <td>10</td>
            <td>R$ 50,00</td>
            <td>R$ 500,00</td>
        </tr>
    </tbody>
    <tfoot>
        <tr>
            <td colspan="3"><strong>Total Geral</strong></td>
            <td><strong>R$ 13.000,00</strong></td>
        </tr>
    </tfoot>
</table>
```

### Atributos de Tabela

**`colspan`**: Mescla células horizontalmente (colunas)
```html
<td colspan="2">Esta célula ocupa 2 colunas</td>
```

**`rowspan`**: Mescla células verticalmente (linhas)
```html
<td rowspan="2">Esta célula ocupa 2 linhas</td>
```

**Exemplo com colspan e rowspan:**
```html
<table>
    <tr>
        <th rowspan="2">Nome</th>
        <th colspan="2">Notas</th>
    </tr>
    <tr>
        <th>Prova 1</th>
        <th>Prova 2</th>
    </tr>
    <tr>
        <td>João</td>
        <td>8.5</td>
        <td>9.0</td>
    </tr>
    <tr>
        <td>Maria</td>
        <td>7.5</td>
        <td>8.5</td>
    </tr>
</table>
```

### Quando Usar Tabelas

**✅ Use tabelas para:**
- Dados tabulares (informações em linhas e colunas)
- Comparações de dados
- Dados estruturados que fazem sentido em formato tabular
- Informações que precisam de alinhamento em colunas

**❌ NÃO use tabelas para:**
- Layout de página (use CSS Grid ou Flexbox)
- Estruturação de conteúdo não tabular
- Design visual (use CSS)
- Organização de elementos da interface

---

## 🎯 Resumo dos Conceitos

### Agrupamento
- **`<div>`**: Container block-level para estruturação
- **`<span>`**: Container inline para texto e elementos inline

### Atributos Padrão
- **`id`**: Identificador único (um por documento)
- **`class`**: Agrupamento por classes (múltiplos elementos)
- **`data-*`**: Dados customizados para JavaScript
- **`style`**: Estilização inline (uso limitado)

### Listas
- **`<ol>`**: Listas ordenadas (numeradas)
- **`<ul>`**: Listas não ordenadas (com bullets)
- **`<dl>`**: Listas de definição (termos e definições)
- **Listas aninhadas**: Hierarquia de informações

### Tabelas
- **`<table>`**: Container da tabela
- **`<tr>`**: Linhas
- **`<td>`**: Células de dados
- **`<th>`**: Células de cabeçalho
- **`<thead>`, `<tbody>`, `<tfoot>`**: Estrutura semântica
- **`colspan` e `rowspan`**: Mesclagem de células

---

## 📚 Próximos Passos

Agora que você aprendeu sobre agrupamento, atributos e listas, você está pronto para:
- Aplicar estilos com CSS usando classes e IDs
- Criar estruturas complexas com divs e spans
- Organizar informações com listas
- Apresentar dados tabulares com tabelas
- Integrar JavaScript usando atributos data-*

**Lembre-se**: Sempre prefira elementos semânticos quando apropriado, e use div/span apenas quando necessário para estruturação ou estilização!



