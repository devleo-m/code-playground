# Aula 4: Tags HTML Básicas - Conteúdo Principal

## 📝 Revisão da Aula Anterior

Antes de começarmos, vamos relembrar os conceitos fundamentais que você já aprendeu:

- **HTML** é uma linguagem de marcação que estrutura o conteúdo web
- Um documento HTML possui a estrutura básica: `<!DOCTYPE html>`, `<html>`, `<head>` e `<body>`
- **Tags** são elementos que definem a estrutura do conteúdo
- **Atributos** fornecem informações adicionais sobre os elementos
- HTML entities são usadas para caracteres especiais
- Comentários HTML ajudam na documentação do código

Agora vamos aprofundar nosso conhecimento sobre as tags fundamentais que formam a base de qualquer documento HTML!

---

## 🏗️ Estrutura Básica de um Documento HTML

### DOCTYPE: A Declaração do Tipo de Documento

A declaração `<!DOCTYPE html>` é a primeira linha de qualquer documento HTML5. Ela informa ao navegador qual versão do HTML está sendo usada e como o documento deve ser interpretado.

```html
<!DOCTYPE html>
```

**Características importantes:**
- Deve ser a primeira linha do arquivo (antes de qualquer outra coisa)
- Não é uma tag HTML, é uma declaração
- Em HTML5, é simplesmente `<!DOCTYPE html>` (muito mais simples que versões anteriores)
- Não possui tag de fechamento
- É case-insensitive, mas use sempre em maiúsculas por convenção

**Por que é importante?**
- Garante que o navegador renderize o documento no modo padrão (standards mode)
- Sem o DOCTYPE, o navegador pode entrar em "quirks mode", causando comportamentos inesperados
- É essencial para validação HTML

### O Elemento HTML: A Raiz do Documento

O elemento `<html>` é o elemento raiz de uma página HTML. Todos os outros elementos HTML (exceto `<!DOCTYPE>`) devem ser descendentes deste elemento.

```html
<!DOCTYPE html>
<html lang="pt-BR">
    <!-- Todo o conteúdo HTML aqui -->
</html>
```

**Atributos importantes:**
- **`lang`**: Define o idioma do documento (ex: `pt-BR`, `en-US`, `es-ES`)
  - Importante para acessibilidade (leitores de tela)
  - Ajuda mecanismos de busca a entender o conteúdo
  - Melhora a experiência do usuário

**Estrutura completa:**
```html
<!DOCTYPE html>
<html lang="pt-BR">
    <head>
        <!-- Metadados aqui -->
    </head>
    <body>
        <!-- Conteúdo visível aqui -->
    </body>
</html>
```

---

## 📋 A Seção HEAD: Metadados do Documento

A tag `<head>` contém metadados (dados sobre dados) sobre o documento HTML. Essas informações não são exibidas na página, mas são essenciais para navegadores, mecanismos de busca e outras ferramentas.

### Estrutura Básica do HEAD

```html
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Título da Página</title>
    <!-- Outros metadados aqui -->
</head>
```

### Meta Tags: Informações Essenciais

#### Meta Charset: Codificação de Caracteres

```html
<meta charset="UTF-8">
```

**O que faz:**
- Define a codificação de caracteres do documento
- UTF-8 suporta praticamente todos os caracteres do mundo
- Permite usar acentos, emojis e caracteres especiais diretamente

**Por que é importante:**
- Sem charset correto, acentos podem aparecer como caracteres estranhos (ex: "ã" vira "Ã£")
- Deve ser a primeira meta tag no `<head>`
- É obrigatória em HTML5

#### Meta Viewport: Responsividade em Dispositivos Móveis

```html
<meta name="viewport" content="width=device-width, initial-scale=1.0">
```

**O que faz:**
- Controla como a página é exibida em dispositivos móveis
- `width=device-width`: usa a largura do dispositivo
- `initial-scale=1.0`: define o zoom inicial

**Por que é importante:**
- Sem viewport, sites podem aparecer muito pequenos em celulares
- Essencial para design responsivo
- Melhora a experiência do usuário em dispositivos móveis

#### Meta Description: Descrição para Mecanismos de Busca

```html
<meta name="description" content="Descrição breve e relevante da página">
```

**O que faz:**
- Fornece uma descrição da página para mecanismos de busca
- Aparece nos resultados de busca (quando o Google decide usá-la)
- Deve ter entre 120-160 caracteres

**Exemplo:**
```html
<meta name="description" content="Aprenda HTML do zero com exemplos práticos e exercícios interativos. Curso completo de desenvolvimento web frontend.">
```

#### Meta Keywords (Obsoleto)

```html
<!-- NÃO USE MAIS! -->
<meta name="keywords" content="HTML, CSS, JavaScript">
```

**Por que não usar:**
- Não é mais usado pelos mecanismos de busca modernos
- Pode até ser considerado spam se usado excessivamente
- Foque em conteúdo de qualidade ao invés de keywords

### A Tag Title: Título do Documento

A tag `<title>` define o título do documento HTML. Este título aparece:
- Na aba do navegador
- Nos resultados de mecanismos de busca
- Quando a página é favoritada
- No histórico do navegador

```html
<title>Meu Site - Página Inicial</title>
```

**Boas práticas:**
- Use títulos descritivos e únicos para cada página
- Mantenha entre 50-60 caracteres
- Inclua palavras-chave relevantes no início
- Seja específico: "Curso de HTML - Aula 4" é melhor que apenas "Aula 4"

**Exemplos:**
```html
<!-- Bom -->
<title>Curso de HTML - Tags Básicas | Aprenda Desenvolvimento Web</title>

<!-- Ruim -->
<title>Página</title>
```

---

## 📄 A Seção BODY: Conteúdo Visível

A tag `<body>` contém todo o conteúdo visível da página web. É aqui que você coloca textos, imagens, links, formulários e todos os elementos que o usuário verá e interagirá.

```html
<body>
    <!-- Todo o conteúdo visível aqui -->
    <h1>Bem-vindo!</h1>
    <p>Este é o conteúdo da página.</p>
</body>
```

**Características:**
- Deve haver apenas um `<body>` por documento
- Contém todos os elementos visíveis
- Pode ter atributos globais (id, class, etc.)
- É onde a mágica acontece!

---

## 📝 Tags Textuais: Estruturando o Conteúdo

### Títulos: Hierarquia e Estrutura

HTML fornece seis níveis de títulos, de `<h1>` (mais importante) a `<h6>` (menos importante).

```html
<h1>Título Principal (Nível 1)</h1>
<h2>Subtítulo (Nível 2)</h2>
<h3>Subtítulo (Nível 3)</h3>
<h4>Subtítulo (Nível 4)</h4>
<h5>Subtítulo (Nível 5)</h5>
<h6>Subtítulo (Nível 6)</h6>
```

**Regras importantes:**
- **Use apenas um `<h1>` por página** - é o título principal
- Mantenha a hierarquia correta - não pule níveis (ex: não vá de h1 para h3)
- Use títulos para estruturar o conteúdo, não apenas para tamanho
- Títulos são essenciais para SEO e acessibilidade

**Exemplo de hierarquia correta:**
```html
<h1>Curso de HTML</h1>
    <h2>Introdução</h2>
        <h3>O que é HTML?</h3>
        <h3>História do HTML</h3>
    <h2>Tags Básicas</h2>
        <h3>Tags de Texto</h3>
        <h3>Tags de Estrutura</h3>
```

**Exemplo de hierarquia incorreta:**
```html
<h1>Curso de HTML</h1>
<h3>Introdução</h3>  <!-- ERRADO: pulou o h2 -->
<h2>Tags Básicas</h2>
```

### Parágrafos: Blocos de Texto

A tag `<p>` define um parágrafo. Navegadores adicionam automaticamente uma linha em branco antes e depois de cada parágrafo.

```html
<p>Este é um parágrafo de texto. Ele contém várias frases que formam uma unidade de pensamento.</p>

<p>Este é outro parágrafo. Ele é separado do anterior por uma linha em branco.</p>
```

**Características:**
- É um elemento de bloco (ocupa toda a largura disponível)
- Navegadores colapsam múltiplos espaços em branco
- Use para blocos de texto relacionados

**Exemplo:**
```html
<p>HTML é uma linguagem de marcação usada para estruturar conteúdo web.</p>
<p>CSS é usado para estilizar esse conteúdo.</p>
<p>JavaScript adiciona interatividade às páginas.</p>
```

### Quebras de Linha: BR

A tag `<br>` cria uma quebra de linha dentro de um bloco de texto. É uma tag vazia (self-closing).

```html
<p>
    Primeira linha<br>
    Segunda linha<br>
    Terceira linha
</p>
```

**Quando usar:**
- Endereços
- Poemas ou versos
- Quando a formatação de linha é importante

**Quando NÃO usar:**
- Para criar espaçamento entre parágrafos (use `<p>`)
- Para layout visual (use CSS)
- Para separar seções (use títulos ou `<hr>`)

**Exemplo de uso correto:**
```html
<address>
    Rua das Flores, 123<br>
    Bairro Centro<br>
    São Paulo - SP<br>
    CEP: 01234-567
</address>
```

### Regra Horizontal: HR

A tag `<hr>` cria uma quebra temática horizontal. É visualmente exibida como uma linha horizontal.

```html
<section>
    <h2>Primeira Seção</h2>
    <p>Conteúdo da primeira seção...</p>
</section>

<hr>

<section>
    <h2>Segunda Seção</h2>
    <p>Conteúdo da segunda seção...</p>
</section>
```

**Características:**
- É uma tag vazia (self-closing)
- Indica uma mudança de tema ou seção
- Visualmente separa conteúdo
- Em HTML5, tem significado semântico (mudança temática)

**Quando usar:**
- Para separar seções de conteúdo relacionado
- Para indicar mudança de assunto
- Em documentos longos para melhorar legibilidade

---

## ✏️ Formatação de Texto

### Strong: Importância Forte

A tag `<strong>` indica que o texto tem **forte importância**. Visualmente, geralmente aparece em negrito.

```html
<p>Este é um texto <strong>muito importante</strong> que precisa ser destacado.</p>
```

**Características:**
- Semântica: indica importância
- Visual: geralmente negrito
- Importante para acessibilidade (leitores de tela enfatizam)

### B: Negrito Visual

A tag `<b>` torna o texto **visualmente em negrito**, mas sem significado semântico.

```html
<p>Este texto está em <b>negrito</b> apenas visualmente.</p>
```

**Diferença entre `<strong>` e `<b>`:**
- `<strong>`: tem significado (importância)
- `<b>`: apenas visual (estilo)

**Quando usar cada um:**
- Use `<strong>` quando o texto for realmente importante
- Use `<b>` apenas para destaque visual sem importância semântica
- **Prefira `<strong>` na maioria dos casos**

### Em: Ênfase

A tag `<em>` indica **ênfase** no texto. Visualmente, geralmente aparece em itálico.

```html
<p>Este é um texto com <em>ênfase</em> em uma palavra.</p>
```

**Características:**
- Semântica: indica ênfase
- Visual: geralmente itálico
- Importante para acessibilidade

### I: Itálico Visual

A tag `<i>` torna o texto **visualmente em itálico**, mas sem significado semântico.

```html
<p>Este texto está em <i>itálico</i> apenas visualmente.</p>
```

**Quando usar:**
- Termos técnicos
- Nomes científicos
- Pensamentos ou citações
- Palavras estrangeiras

**Exemplo:**
```html
<p>O <i>Homo sapiens</i> é a espécie humana moderna.</p>
<p>Ele pensou: <i>Será que isso vai funcionar?</i></p>
```

### Mark: Texto Marcado

A tag `<mark>` marca ou destaca partes do texto, geralmente com fundo amarelo.

```html
<p>Este é um texto com uma <mark>palavra marcada</mark> para destaque.</p>
```

**Quando usar:**
- Destacar texto relevante para o usuário
- Resultados de busca
- Texto que precisa atenção especial

### Sub: Subscrito

A tag `<sub>` cria texto **subscrito** (abaixo da linha normal).

```html
<p>A fórmula da água é H<sub>2</sub>O.</p>
<p>O número 2 está na posição 10<sub>base</sub>.</p>
```

**Quando usar:**
- Fórmulas químicas (H₂O, CO₂)
- Notas de rodapé
- Expressões matemáticas

### Sup: Sobrescrito

A tag `<sup>` cria texto **sobrescrito** (acima da linha normal).

```html
<p>O resultado é 2<sup>3</sup> = 8.</p>
<p>Este é o 1<sup>o</sup> lugar.</p>
<p>Referência<sup>1</sup> ao final do texto.</p>
```

**Quando usar:**
- Exponenciação matemática (2³, x²)
- Ordinais (1º, 2º, 3º)
- Notas de rodapé
- Referências

### Pre: Texto Pré-formatado

A tag `<pre>` preserva espaços em branco e quebras de linha. O texto é exibido em fonte monoespaçada.

```html
<pre>
    Este texto
    mantém     os espaços
    e quebras de linha
    exatamente como estão.
</pre>
```

**Quando usar:**
- Código de programação
- Poemas com formatação específica
- ASCII art
- Qualquer texto onde a formatação é importante

**Exemplo com código:**
```html
<pre>
function exemplo() {
    console.log("Olá, mundo!");
}
</pre>
```

---

## 🔗 Links: Navegação e Hiperlinks

A tag `<a>` (âncora) cria um hiperlink, permitindo navegação entre páginas ou seções.

### Link Básico

```html
<a href="https://www.exemplo.com">Visite nosso site</a>
```

### Atributos Importantes

#### href: O Destino do Link

```html
<!-- Link externo -->
<a href="https://www.google.com">Google</a>

<!-- Link interno (mesmo site) -->
<a href="sobre.html">Sobre Nós</a>

<!-- Link para seção da mesma página (âncora) -->
<a href="#secao1">Ir para Seção 1</a>

<!-- Link de email -->
<a href="mailto:contato@exemplo.com">Enviar Email</a>

<!-- Link de telefone -->
<a href="tel:+5511999999999">Ligar Agora</a>
```

#### target: Onde Abrir o Link

```html
<!-- Abre na mesma aba (padrão) -->
<a href="pagina.html">Link Normal</a>

<!-- Abre em nova aba -->
<a href="https://www.exemplo.com" target="_blank">Link Externo</a>
```

**⚠️ IMPORTANTE:** Sempre use `rel="noopener noreferrer"` com `target="_blank"` por segurança:

```html
<a href="https://www.exemplo.com" target="_blank" rel="noopener noreferrer">
    Link Seguro
</a>
```

**Por quê?**
- Previne vulnerabilidades de segurança
- Melhora performance
- Boa prática de segurança web

#### title: Título do Link

```html
<a href="pagina.html" title="Clique para ver mais informações">
    Saiba Mais
</a>
```

O `title` aparece quando o usuário passa o mouse sobre o link.

### Exemplos Práticos de Links

**Navegação entre seções:**
```html
<nav>
    <a href="#inicio">Início</a>
    <a href="#sobre">Sobre</a>
    <a href="#contato">Contato</a>
</nav>

<section id="inicio">
    <h2>Início</h2>
    <p>Conteúdo da seção inicial...</p>
</section>
```

**Links externos seguros:**
```html
<p>
    Visite o 
    <a href="https://www.mozilla.org" target="_blank" rel="noopener noreferrer">
        site da Mozilla
    </a>
    para mais informações.
</p>
```

**Link de download:**
```html
<a href="documento.pdf" download>Baixar PDF</a>
```

**Link de email:**
```html
<p>
    Entre em contato: 
    <a href="mailto:contato@exemplo.com?subject=Contato&body=Olá!">
        contato@exemplo.com
    </a>
</p>
```

---

## 🎯 Resumo das Tags Aprendidas

### Estrutura Básica
- `<!DOCTYPE html>` - Declaração do tipo de documento
- `<html>` - Elemento raiz
- `<head>` - Metadados
- `<body>` - Conteúdo visível

### Metadados (HEAD)
- `<meta charset="UTF-8">` - Codificação de caracteres
- `<meta name="viewport">` - Configuração mobile
- `<meta name="description">` - Descrição para SEO
- `<title>` - Título do documento

### Texto e Estrutura
- `<h1>` a `<h6>` - Títulos (hierarquia)
- `<p>` - Parágrafos
- `<br>` - Quebra de linha
- `<hr>` - Regra horizontal

### Formatação
- `<strong>` - Importância forte (semântico)
- `<b>` - Negrito (visual)
- `<em>` - Ênfase (semântico)
- `<i>` - Itálico (visual)
- `<mark>` - Texto marcado
- `<sub>` - Subscrito
- `<sup>` - Sobrescrito
- `<pre>` - Texto pré-formatado

### Navegação
- `<a>` - Links e hiperlinks

---

## ✅ Checklist de Aprendizado

Antes de prosseguir, certifique-se de entender:

- [ ] A estrutura completa de um documento HTML
- [ ] A função de cada seção (head e body)
- [ ] Como usar meta tags corretamente
- [ ] A hierarquia de títulos (h1-h6)
- [ ] Quando usar cada tag de formatação
- [ ] A diferença entre tags semânticas e visuais
- [ ] Como criar links funcionais e seguros
- [ ] Boas práticas de estruturação HTML

---

**Próximo passo:** Revise a Aula Simplificada para reforçar os conceitos com analogias e exemplos práticos do dia a dia!

