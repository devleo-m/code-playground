# Aula 7: Marcação Semântica e Estilização Básica - Conteúdo Principal

## 📖 O que é Marcação Semântica?

**Marcação semântica** (Semantic Markup) é o uso de tags HTML que transmitem o **significado** e a **estrutura** do conteúdo, não apenas sua aparência visual. Esta abordagem torna as páginas web mais acessíveis tanto para humanos quanto para máquinas, fornecendo contexto sobre as diferentes partes do conteúdo, como títulos, parágrafos, artigos e menus de navegação.

### Características Principais

A marcação semântica utiliza elementos HTML que descrevem **o que o conteúdo é**, não apenas **como ele parece**:

- **Significado claro**: Cada tag comunica o propósito do conteúdo
- **Estrutura organizada**: Hierarquia e relacionamento entre elementos
- **Acessibilidade**: Leitores de tela compreendem melhor a estrutura
- **SEO melhorado**: Mecanismos de busca indexam o conteúdo com mais precisão
- **Manutenibilidade**: Código mais fácil de entender e manter

### Por que é Importante?

1. **Acessibilidade**: Pessoas com deficiência visual usam leitores de tela que dependem da semântica
2. **SEO**: Mecanismos de busca usam elementos semânticos para entender o conteúdo
3. **Manutenção**: Código semântico é mais fácil de entender e modificar
4. **Padrões Web**: Segue as melhores práticas da web moderna
5. **Compatibilidade**: Funciona melhor com ferramentas e tecnologias assistivas

---

## ✏️ Destacar Mudanças em Documentos

### O Elemento `<del>`

O elemento `<del>` em HTML representa texto que foi **deletado** ou **removido** de um documento. Os navegadores geralmente renderizam o texto deletado com um **risco** (strikethrough), indicando visualmente que o conteúdo não é mais válido ou preciso.

#### Sintaxe

```html
<del>Texto deletado</del>
```

#### Exemplo

```html
<p>
    O preço era <del>R$ 100,00</del> agora é R$ 80,00.
</p>
```

#### Atributos Importantes

- `cite`: URL de um documento que explica por que o texto foi deletado
- `datetime`: Data e hora da deleção no formato ISO 8601

```html
<del cite="https://exemplo.com/atualizacao" datetime="2024-01-15T10:30:00Z">
    Informação antiga
</del>
```

### O Elemento `<s>`

O elemento `<s>` em HTML representa conteúdo que **não é mais preciso** ou **relevante**. Indica coisas que não são mais corretas, precisas ou relevantes. Os navegadores geralmente renderizam este elemento com um **risco**, indicando visualmente que o texto foi removido ou não é mais válido.

#### Sintaxe

```html
<s>Texto não mais relevante</s>
```

#### Exemplo

```html
<p>
    <s>Promoção válida até 31 de dezembro</s>
</p>
<p>
    Promoção estendida até 15 de janeiro!
</p>
```

#### Diferença entre `<del>` e `<s>`

- **`<del>`**: Usado para edições e revisões de documentos (indica remoção intencional)
- **`<s>`**: Usado para conteúdo que não é mais relevante (indica obsolescência)

### O Elemento `<ins>`

O elemento `<ins>` em HTML representa texto que foi **inserido** em um documento. É usado para indicar adições ou atualizações de conteúdo, frequentemente exibido com um **sublinhado** para distinguir visualmente do texto original.

#### Sintaxe

```html
<ins>Texto inserido</ins>
```

#### Exemplo

```html
<p>
    O preço era R$ 100,00, <ins>agora é R$ 80,00</ins>.
</p>
```

#### Atributos Importantes

- `cite`: URL de um documento que explica por que o texto foi inserido
- `datetime`: Data e hora da inserção no formato ISO 8601

```html
<ins cite="https://exemplo.com/atualizacao" datetime="2024-01-15T10:30:00Z">
    Nova informação adicionada
</ins>
```

#### Usando `<del>` e `<ins>` Juntos

```html
<p>
    Reunião marcada para <del datetime="2024-01-15">segunda-feira</del>
    <ins datetime="2024-01-16">terça-feira</ins>.
</p>
```

---

## 📚 Citações e Referências

### O Elemento `<abbr>`

A tag `<abbr>` em HTML representa uma **abreviação** ou **acrônimo**. É útil para fornecer uma descrição completa do termo abreviado quando o usuário passa o mouse sobre ele, melhorando a acessibilidade e clareza.

#### Sintaxe

```html
<abbr title="Texto completo da abreviação">Abreviação</abbr>
```

#### Exemplo

```html
<p>
    O <abbr title="HyperText Markup Language">HTML</abbr> é a linguagem
    de marcação da web.
</p>

<p>
    A <abbr title="Organização das Nações Unidas">ONU</abbr> foi fundada
    em 1945.
</p>
```

#### Boas Práticas

- Sempre use o atributo `title` para fornecer a forma expandida
- Use para abreviações e acrônimos que podem não ser familiares ao leitor
- Não use para siglas muito conhecidas (ex: HTML, CSS, JS em contexto técnico)

### O Elemento `<cite>`

O elemento `<cite>` em HTML é usado para definir o **título de uma obra criativa** (por exemplo, um livro, artigo, música, filme, pintura, escultura, etc.). É tipicamente usado para fornecer uma referência ou citação de uma fonte. O conteúdo dentro do elemento `<cite>` é frequentemente renderizado em **itálico** pelos navegadores, mas este estilo pode ser sobrescrito com CSS.

#### Sintaxe

```html
<cite>Título da Obra</cite>
```

#### Exemplo

```html
<p>
    Como disse em <cite>O Pequeno Príncipe</cite>:
    "O essencial é invisível aos olhos."
</p>

<blockquote>
    <p>
        A tecnologia é melhor quando ela traz as pessoas juntas.
    </p>
    <cite>— Matt Mullenweg</cite>
</blockquote>
```

#### Uso Correto

- Use para títulos de obras (livros, filmes, músicas, artigos)
- Não use para nomes de pessoas (use `<span>` ou texto simples)
- Pode ser usado dentro de `<blockquote>` para citar a fonte

### O Elemento `<dfn>`

O elemento `<dfn>` em HTML representa a **instância definidora de um termo**. É usado para indicar o local específico onde uma palavra ou frase está sendo definida pela primeira vez dentro de um documento. Tipicamente, o termo sendo definido é incluído dentro das tags `<dfn>`, e frequentemente uma definição ou explicação do termo é fornecida nas proximidades.

#### Sintaxe

```html
<dfn>Termo sendo definido</dfn>
```

#### Exemplo

```html
<p>
    O <dfn>HTML</dfn> (HyperText Markup Language) é uma linguagem
    de marcação usada para estruturar conteúdo na web.
</p>

<p>
    <dfn>CSS</dfn> significa Cascading Style Sheets, uma linguagem
    usada para estilizar documentos HTML.
</p>
```

#### Boas Práticas

- Use apenas na primeira ocorrência do termo no documento
- Forneça a definição completa próxima ao elemento
- Pode ser combinado com `<abbr>` quando o termo é uma abreviação

### O Elemento `<address>`

O elemento `<address>` em HTML representa **informações de contato** para o autor ou proprietário de um documento ou artigo. Isso pode incluir endereços físicos, endereços de email, números de telefone e links de redes sociais. É tipicamente usado dentro do `<footer>` de uma página ou seção para fornecer detalhes de contato.

#### Sintaxe

```html
<address>
    Informações de contato
</address>
```

#### Exemplo

```html
<footer>
    <address>
        <p>Escrito por João Silva</p>
        <p>
            Email: <a href="mailto:joao@exemplo.com">joao@exemplo.com</a>
        </p>
        <p>
            Telefone: <a href="tel:+5511999999999">(11) 99999-9999</a>
        </p>
        <p>
            Rua Exemplo, 123 - São Paulo, SP
        </p>
    </address>
</footer>
```

#### Boas Práticas

- Use apenas para informações de contato reais
- Não use para endereços fictícios ou de exemplo
- Pode conter links, mas não outros elementos semânticos complexos

### O Elemento `<blockquote>`

O elemento `blockquote` em HTML representa uma seção de texto que é **citada de outra fonte**. É usado para indicar que o conteúdo incluído é uma citação estendida, frequentemente exibida com indentação ou outros indicadores visuais para distinguir do texto circundante. O atributo `cite` pode ser usado para especificar a URL do documento ou mensagem de origem.

#### Sintaxe

```html
<blockquote cite="URL da fonte">
    Texto da citação
</blockquote>
```

#### Exemplo

```html
<blockquote cite="https://www.exemplo.com/artigo">
    <p>
        A única forma de fazer um excelente trabalho é amar o que você faz.
    </p>
    <cite>— Steve Jobs</cite>
</blockquote>
```

#### Atributos

- `cite`: URL do documento ou mensagem de origem da citação

### O Elemento `<q>`

O elemento `<q>` em HTML representa uma **citação curta e inline**. Os navegadores tipicamente renderizam este elemento com **aspas** ao redor do conteúdo que ele contém. É projetado para citações breves que cabem dentro de um parágrafo, em oposição a citações mais longas em nível de bloco que usariam o elemento `<blockquote>`.

#### Sintaxe

```html
<q cite="URL da fonte">Citação curta</q>
```

#### Exemplo

```html
<p>
    Como disse Einstein: <q cite="https://exemplo.com">
    A imaginação é mais importante que o conhecimento
    </q>.
</p>
```

#### Diferença entre `<q>` e `<blockquote>`

- **`<q>`**: Para citações curtas inline (dentro de parágrafos)
- **`<blockquote>`**: Para citações longas em bloco (separadas do texto)

---

## 🏗️ Tags de Layout Semântico

### O Elemento `<header>`

O elemento `<header>` representa **conteúdo introdutório**, tipicamente contendo um grupo de auxílios introdutórios ou de navegação. Pode conter um título, logo, formulário de busca ou outro conteúdo relevante. É usado para definir a seção superior de um documento, artigo ou seção.

#### Sintaxe

```html
<header>
    Conteúdo do cabeçalho
</header>
```

#### Exemplo

```html
<header>
    <h1>Meu Site</h1>
    <nav>
        <ul>
            <li><a href="#inicio">Início</a></li>
            <li><a href="#sobre">Sobre</a></li>
            <li><a href="#contato">Contato</a></li>
        </ul>
    </nav>
</header>
```

#### Uso Correto

- Pode aparecer múltiplas vezes na página (um por seção)
- Não deve ser usado dentro de `<footer>`, `<address>` ou outro `<header>`
- Geralmente contém o título principal da página ou seção

### O Elemento `<nav>`

O elemento `<nav>` em HTML é usado para definir uma seção de uma página que contém **links de navegação**. É destinado para blocos de navegação principais, como o menu de um site, um índice de conteúdo ou um conjunto de breadcrumbs. Usar `<nav>` ajuda a estruturar seu conteúdo e torna-o mais acessível para leitores de tela e mecanismos de busca, identificando claramente seções de navegação.

#### Sintaxe

```html
<nav>
    Links de navegação
</nav>
```

#### Exemplo

```html
<nav>
    <ul>
        <li><a href="/">Início</a></li>
        <li><a href="/sobre">Sobre</a></li>
        <li><a href="/produtos">Produtos</a></li>
        <li><a href="/contato">Contato</a></li>
    </ul>
</nav>
```

#### Boas Práticas

- Use apenas para navegação principal (não para todos os links)
- Pode aparecer múltiplas vezes (menu principal, breadcrumbs, etc.)
- Geralmente contém uma lista de links

### O Elemento `<main>`

O elemento `<main>` em HTML define o **conteúdo principal** do `<body>` de um documento. Deve conter o tópico central da página, excluindo qualquer conteúdo que seja repetido em múltiplas páginas, como navegação, cabeçalhos ou rodapés. Usar `<main>` ajuda a melhorar a acessibilidade e fornece uma estrutura clara para mecanismos de busca e tecnologias assistivas entenderem o propósito da página.

#### Sintaxe

```html
<main>
    Conteúdo principal
</main>
```

#### Exemplo

```html
<body>
    <header>
        <h1>Meu Site</h1>
    </header>
    
    <nav>...</nav>
    
    <main>
        <article>
            <h2>Título do Artigo</h2>
            <p>Conteúdo principal do artigo...</p>
        </article>
    </main>
    
    <footer>...</footer>
</body>
```

#### Regras Importantes

- Deve haver apenas **um** elemento `<main>` por página
- Não deve ser descendente de `<article>`, `<aside>`, `<footer>`, `<header>` ou `<nav>`
- Deve conter conteúdo único à página

### O Elemento `<section>`

A tag `<section>` em HTML é usada para definir **agrupamentos temáticos de conteúdo** dentro de um documento. É uma forma de organizar conteúdo relacionado, como capítulos, introduções ou itens de notícias. Uma seção tipicamente tem um título e pode conter outros elementos HTML para estruturar o conteúdo dentro dela.

#### Sintaxe

```html
<section>
    <h2>Título da Seção</h2>
    Conteúdo da seção
</section>
```

#### Exemplo

```html
<main>
    <section>
        <h2>Introdução</h2>
        <p>Conteúdo introdutório...</p>
    </section>
    
    <section>
        <h2>Desenvolvimento</h2>
        <p>Conteúdo de desenvolvimento...</p>
    </section>
    
    <section>
        <h2>Conclusão</h2>
        <p>Conteúdo de conclusão...</p>
    </section>
</main>
```

#### Quando Usar

- Para agrupar conteúdo temático relacionado
- Quando o conteúdo precisa de um título próprio
- Para criar seções distintas dentro de um documento

### O Elemento `<article>`

O elemento `<article>` em HTML representa uma **composição autocontida** em um documento, página, aplicação ou site. É destinado a ser independentemente distribuível ou reutilizável, por exemplo, em sindicação. Isso poderia ser uma postagem de fórum, um artigo de revista ou jornal, uma entrada de blog, um comentário enviado pelo usuário ou qualquer outro item de conteúdo independente.

#### Sintaxe

```html
<article>
    Conteúdo do artigo
</article>
</html>
```

#### Exemplo

```html
<main>
    <article>
        <header>
            <h2>Título do Artigo</h2>
            <p>Publicado em <time datetime="2024-01-15">15 de janeiro de 2024</time></p>
        </header>
        <p>Conteúdo do artigo...</p>
        <footer>
            <p>Autor: João Silva</p>
        </footer>
    </article>
    
    <article>
        <header>
            <h2>Outro Artigo</h2>
        </header>
        <p>Conteúdo de outro artigo...</p>
    </article>
</main>
```

#### Características

- Deve fazer sentido independentemente do resto da página
- Pode ser aninhado (artigo dentro de artigo)
- Geralmente contém `<header>` e `<footer>` próprios

### O Elemento `<aside>`

O elemento `<aside>` em HTML representa uma seção de uma página que é **tangencialmente relacionada** ao conteúdo principal. É frequentemente usado para barras laterais, citações destacadas ou outro conteúdo que fornece informações ou contexto adicionais, mas não é essencial para entender o conteúdo principal. Pense nisso como conteúdo que pode ser removido sem impactar significativamente a compreensão do usuário sobre o tópico principal.

#### Sintaxe

```html
<aside>
    Conteúdo relacionado
</aside>
```

#### Exemplo

```html
<main>
    <article>
        <h2>Título do Artigo</h2>
        <p>Conteúdo principal do artigo...</p>
    </article>
    
    <aside>
        <h3>Artigos Relacionados</h3>
        <ul>
            <li><a href="#">Artigo 1</a></li>
            <li><a href="#">Artigo 2</a></li>
        </ul>
    </aside>
</main>
```

#### Uso Comum

- Barras laterais com conteúdo relacionado
- Citações destacadas
- Informações complementares
- Anúncios (quando relevantes ao conteúdo)

### O Elemento `<footer>`

O elemento `<footer>` em HTML representa um **container para conteúdo introdutório ou links de navegação** que tipicamente aparecem no final de uma seção ou documento. Geralmente contém informações sobre o autor, dados de copyright, termos de uso, informações de contato, documentos relacionados e às vezes navegação. Um rodapé não está necessariamente restrito ao fundo da página; pode ser usado dentro de seções para denotar o fim dessa área de conteúdo específica.

#### Sintaxe

```html
<footer>
    Conteúdo do rodapé
</footer>
```

#### Exemplo

```html
<footer>
    <p>&copy; 2024 Meu Site. Todos os direitos reservados.</p>
    <address>
        <p>Contato: <a href="mailto:contato@exemplo.com">contato@exemplo.com</a></p>
    </address>
    <nav>
        <ul>
            <li><a href="/privacidade">Privacidade</a></li>
            <li><a href="/termos">Termos de Uso</a></li>
        </ul>
    </nav>
</footer>
```

#### Uso Correto

- Pode aparecer múltiplas vezes (um por seção ou artigo)
- Geralmente contém informações de copyright, contato e links relacionados
- Não deve conter `<header>` ou outro `<footer>`

---

## 🎨 Fundamentos de Estilização

### O que é CSS?

**CSS** (Cascading Style Sheets - Folhas de Estilo em Cascata) é uma linguagem usada para descrever a **aparência** e **formatação** de um documento escrito em HTML. O CSS controla propriedades como cores, fontes, layout e responsividade, permitindo personalizar a aparência e a experiência do usuário de suas páginas web além dos estilos padrão do navegador.

### CSS Inline

**CSS inline** envolve aplicar estilos diretamente a elementos HTML individuais usando o atributo `style`. Este método permite definir propriedades visuais específicas para um único elemento, sobrescrevendo quaisquer estilos definidos em folhas de estilo externas ou blocos de estilo internos.

#### Sintaxe

```html
<elemento style="propriedade: valor;">
```

#### Exemplo

```html
<h1 style="color: blue; font-size: 32px;">Título Azul</h1>
<p style="background-color: yellow; padding: 10px;">
    Parágrafo com fundo amarelo
</p>
```

#### Vantagens

- Aplicação rápida e direta
- Sobrescreve estilos externos e internos
- Útil para estilos únicos e específicos

#### Desvantagens

- Difícil de manter em projetos grandes
- Não promove reutilização
- Mistura estrutura (HTML) com apresentação (CSS)
- Aumenta o tamanho do arquivo HTML

### CSS Interno

**CSS interno** permite adicionar estilos diretamente dentro de um documento HTML usando a tag `<style>`. Esta tag é tipicamente colocada dentro da seção `<head>` do seu arquivo HTML. Dentro da tag `<style>`, você pode definir regras CSS que se aplicam a elementos HTML específicos naquela página, controlando sua aparência e layout.

#### Sintaxe

```html
<head>
    <style>
        seletor {
            propriedade: valor;
        }
    </style>
</head>
```

#### Exemplo

```html
<!DOCTYPE html>
<html lang="pt-BR">
<head>
    <meta charset="UTF-8">
    <title>CSS Interno</title>
    <style>
        h1 {
            color: blue;
            font-size: 32px;
            text-align: center;
        }
        
        p {
            color: #333;
            line-height: 1.6;
            margin: 10px 0;
        }
        
        .destaque {
            background-color: yellow;
            padding: 10px;
        }
    </style>
</head>
<body>
    <h1>Título Estilizado</h1>
    <p>Parágrafo normal</p>
    <p class="destaque">Parágrafo destacado</p>
</body>
</html>
```

#### Vantagens

- Mantém estilos organizados em um local
- Aplicável a toda a página
- Mais fácil de manter que CSS inline
- Permite uso de seletores CSS completos

#### Desvantagens

- Não pode ser reutilizado em múltiplas páginas
- Aumenta o tamanho do arquivo HTML
- Mistura estrutura com apresentação no mesmo arquivo

### CSS Externo

**CSS externo** envolve criar arquivos separados (com extensão `.css`) para manter todas as regras de estilo de um documento HTML. Esses arquivos CSS são então vinculados ao documento HTML usando a tag `<link>`, permitindo aplicar os mesmos estilos em múltiplas páginas e manter seu código HTML limpo e organizado. Esta abordagem promove reutilização e manutenibilidade do design do seu site.

#### Sintaxe

**Arquivo HTML:**
```html
<head>
    <link rel="stylesheet" href="caminho/para/estilo.css">
</head>
```

**Arquivo CSS (estilo.css):**
```css
h1 {
    color: blue;
    font-size: 32px;
}

p {
    color: #333;
    line-height: 1.6;
}
```

#### Exemplo Completo

**index.html:**
```html
<!DOCTYPE html>
<html lang="pt-BR">
<head>
    <meta charset="UTF-8">
    <title>CSS Externo</title>
    <link rel="stylesheet" href="estilo.css">
</head>
<body>
    <h1>Título Estilizado</h1>
    <p>Parágrafo estilizado</p>
</body>
</html>
```

**estilo.css:**
```css
/* Estilos globais */
body {
    font-family: Arial, sans-serif;
    margin: 0;
    padding: 20px;
    background-color: #f5f5f5;
}

h1 {
    color: #333;
    font-size: 32px;
    text-align: center;
    margin-bottom: 20px;
}

p {
    color: #666;
    line-height: 1.6;
    margin: 10px 0;
}
```

#### Vantagens

- **Reutilização**: Um arquivo CSS pode ser usado em múltiplas páginas
- **Manutenibilidade**: Fácil de atualizar e manter
- **Separação de responsabilidades**: HTML para estrutura, CSS para estilo
- **Performance**: Arquivo CSS pode ser armazenado em cache pelo navegador
- **Organização**: Código mais limpo e organizado

#### Desvantagens

- Requer uma requisição HTTP adicional (mas pode ser cacheado)
- Pode haver um pequeno atraso no carregamento inicial

### Comparação dos Métodos

| Método | Uso Ideal | Manutenibilidade | Reutilização |
|--------|-----------|------------------|--------------|
| **Inline** | Estilos únicos e específicos | Baixa | Nenhuma |
| **Interno** | Estilos específicos de uma página | Média | Nenhuma |
| **Externo** | Estilos globais e reutilizáveis | Alta | Total |

### Ordem de Precedência

Quando múltiplos estilos se aplicam ao mesmo elemento, a ordem de precedência é:

1. **CSS Inline** (maior prioridade)
2. **CSS Interno**
3. **CSS Externo** (menor prioridade)

---

## ⚡ Incluindo JavaScript

JavaScript pode ser adicionado a documentos HTML para tornar páginas web interativas. Isso é feito incorporando o código JavaScript diretamente dentro do HTML ou vinculando a arquivos JavaScript externos. A incorporação usa a tag `<script>`, enquanto a vinculação usa a tag `<script>` com o atributo `src` apontando para o arquivo JavaScript.

### JavaScript Inline

JavaScript inline é colocado diretamente dentro do elemento HTML usando o atributo de evento ou dentro da tag `<script>` no HTML.

#### Usando Atributos de Evento

```html
<button onclick="alert('Olá, mundo!')">Clique Aqui</button>
```

#### Usando Tag `<script>` no Body

```html
<body>
    <button id="meuBotao">Clique Aqui</button>
    
    <script>
        document.getElementById('meuBotao').addEventListener('click', function() {
            alert('Olá, mundo!');
        });
    </script>
</body>
```

### JavaScript Interno

JavaScript interno é colocado dentro da tag `<script>` na seção `<head>` ou antes do fechamento de `</body>`.

#### No `<head>`

```html
<head>
    <meta charset="UTF-8">
    <title>JavaScript Interno</title>
    <script>
        function saudacao() {
            alert('Bem-vindo!');
        }
    </script>
</head>
<body>
    <button onclick="saudacao()">Clique Aqui</button>
</body>
```

#### Antes de `</body>` (Recomendado)

```html
<body>
    <button id="meuBotao">Clique Aqui</button>
    
    <script>
        document.getElementById('meuBotao').addEventListener('click', function() {
            alert('Botão clicado!');
        });
    </script>
</body>
```

**Por que antes de `</body>`?**
- Garante que os elementos HTML já foram carregados
- Evita erros de elementos não encontrados
- Melhora a performance (não bloqueia o carregamento da página)

### JavaScript Externo

JavaScript externo envolve criar arquivos separados (com extensão `.js`) e vinculá-los ao HTML usando a tag `<script>` com o atributo `src`.

#### Sintaxe

**Arquivo HTML:**
```html
<script src="caminho/para/script.js"></script>
```

#### Exemplo Completo

**index.html:**
```html
<!DOCTYPE html>
<html lang="pt-BR">
<head>
    <meta charset="UTF-8">
    <title>JavaScript Externo</title>
</head>
<body>
    <button id="meuBotao">Clique Aqui</button>
    <p id="mensagem"></p>
    
    <script src="script.js"></script>
</body>
</html>
```

**script.js:**
```javascript
document.getElementById('meuBotao').addEventListener('click', function() {
    document.getElementById('mensagem').textContent = 'Botão clicado!';
});
```

#### Vantagens do JavaScript Externo

- **Reutilização**: Um arquivo pode ser usado em múltiplas páginas
- **Manutenibilidade**: Fácil de atualizar e manter
- **Separação de responsabilidades**: HTML para estrutura, JS para comportamento
- **Performance**: Arquivo JS pode ser armazenado em cache
- **Organização**: Código mais limpo e organizado

### Atributos da Tag `<script>`

#### `src`
Especifica a URL do arquivo JavaScript externo.

```html
<script src="script.js"></script>
```

#### `defer`
O script é executado após o documento ter sido parseado.

```html
<script src="script.js" defer></script>
```

#### `async`
O script é executado assincronamente (não bloqueia o parsing).

```html
<script src="script.js" async></script>
```

#### `type`
Especifica o tipo MIME do script (geralmente não necessário em HTML5).

```html
<script type="text/javascript" src="script.js"></script>
```

### Boas Práticas

1. **Coloque scripts antes de `</body>`** para melhor performance
2. **Use JavaScript externo** para código reutilizável
3. **Evite JavaScript inline** em atributos de evento quando possível
4. **Use `defer` ou `async`** para scripts que não bloqueiam
5. **Organize seu código** em funções e módulos

---

## 📝 Resumo da Aula

Nesta aula, você aprendeu:

✅ **Marcação semântica** usa tags HTML que transmitem significado e estrutura  
✅ **Elementos de mudança** (`<del>`, `<s>`, `<ins>`) destacam edições em documentos  
✅ **Elementos de citação** (`<abbr>`, `<cite>`, `<dfn>`, `<address>`, `<blockquote>`, `<q>`) fornecem contexto e referências  
✅ **Tags de layout semântico** (`<header>`, `<nav>`, `<main>`, `<section>`, `<article>`, `<aside>`, `<footer>`) estruturam páginas web  
✅ **CSS inline** aplica estilos diretamente em elementos  
✅ **CSS interno** usa a tag `<style>` no `<head>`  
✅ **CSS externo** vincula arquivos `.css` separados  
✅ **JavaScript** pode ser inline, interno ou externo usando a tag `<script>`  

### Próximos Passos

Na próxima aula, você aprenderá sobre:
- Formulários avançados
- Validação de formulários
- Elementos de entrada modernos
- Técnicas de acessibilidade em formulários

---

## 🔍 Conceitos-Chave para Revisão

- **Semântica**: Uso de tags apropriadas para transmitir significado
- **Acessibilidade**: Tornar conteúdo acessível para todos os usuários
- **SEO**: Otimização para mecanismos de busca
- **Separação de responsabilidades**: HTML (estrutura), CSS (estilo), JS (comportamento)
- **Manutenibilidade**: Código fácil de entender e modificar

