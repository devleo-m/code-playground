# Aula 1: Introdução ao HTML - Conteúdo Principal

## 📖 O que é HTML?

**HTML**, ou **HyperText Markup Language** (Linguagem de Marcação de Hipertexto), é a linguagem padrão utilizada para criar e estruturar conteúdo na World Wide Web. É a base fundamental de todas as páginas web que você acessa diariamente.

### Características Principais

HTML utiliza um sistema de **tags** (etiquetas) para estruturar o conteúdo, permitindo que navegadores web interpretem e exibam corretamente elementos como:
- Títulos e parágrafos
- Imagens e vídeos
- Links e navegação
- Formulários e campos de entrada
- Listas e tabelas
- E muito mais

### Função do HTML na Web

O HTML funciona como o **esqueleto** de uma página web:
- Define a **estrutura** e **organização** do conteúdo
- Indica a **hierarquia** das informações (títulos principais, subtítulos, parágrafos)
- Estabelece a **semântica** (significado) dos elementos
- Permite que navegadores, leitores de tela e mecanismos de busca compreendam o conteúdo

---

## 🌐 Linguagens de Marcação (Markup Languages)

### O que são Linguagens de Marcação?

Linguagens de marcação são sistemas para **anotar texto** de forma que seja legível tanto por humanos quanto por computadores. Elas utilizam **tags** para definir elementos dentro de um documento, especificando como o texto deve ser:
- **Estruturado** (organização hierárquica)
- **Formatado** (aparência visual)
- **Exibido** (como será apresentado)

### Exemplos de Linguagens de Marcação

1. **HTML (HyperText Markup Language)**
   - Usado para estruturar conteúdo web
   - Tags como `<p>` para parágrafos, `<h1>` para títulos principais
   - Interpretado por navegadores web

2. **XML (eXtensible Markup Language)**
   - Usado para definir estruturas de dados
   - Permite criar tags personalizadas
   - Utilizado em APIs, configurações e transferência de dados

3. **Markdown**
   - Linguagem de marcação simplificada
   - Usada em documentação, READMEs e plataformas como GitHub
   - Converte texto simples em HTML

4. **LaTeX**
   - Usado para documentos acadêmicos e científicos
   - Foco em formatação tipográfica complexa

### Por que Usar Tags?

As tags funcionam como **instruções** para o navegador:
- `<h1>Meu Título</h1>` → Indica que "Meu Título" é um cabeçalho principal
- `<p>Meu parágrafo</p>` → Indica que "Meu parágrafo" é um parágrafo de texto
- `<img src="foto.jpg">` → Indica que deve exibir uma imagem

---

## 🎨 Desenvolvimento Frontend

### O que é Frontend Development?

**Frontend development** (desenvolvimento frontend) é a prática de criar a **interface do usuário** e a **experiência do usuário** de um website ou aplicação web. Foca nas partes do site com as quais os usuários **interagem diretamente**.

### Componentes do Frontend

O desenvolvimento frontend moderno é construído sobre três tecnologias fundamentais:

#### 1. **HTML (Estrutura)**
- Define o **conteúdo** e a **estrutura** da página
- É o esqueleto da página web
- Sem HTML, não há página para exibir

#### 2. **CSS (Estilo)**
- Controla a **aparência visual** da página
- Define cores, fontes, layouts, espaçamentos
- É a "roupa" que veste o HTML

#### 3. **JavaScript (Comportamento)**
- Adiciona **interatividade** e **dinamismo**
- Permite que a página responda a ações do usuário
- É o "cérebro" que torna a página reativa

### Como Funcionam Juntos?

```
HTML (Estrutura) + CSS (Estilo) + JavaScript (Comportamento) = Página Web Completa
```

**Exemplo Prático:**
- **HTML**: Define que existe um botão
- **CSS**: Estiliza o botão (cor, tamanho, posição)
- **JavaScript**: Faz o botão executar uma ação quando clicado

---

## 📄 HTML: A Linguagem de Marcação da Web

### Estrutura Básica de um Documento HTML

Todo documento HTML segue uma estrutura fundamental:

```html
<!DOCTYPE html>
<html lang="pt-BR">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Título da Página</title>
</head>
<body>
    <!-- Conteúdo da página aqui -->
</body>
</html>
```

### Componentes da Estrutura

1. **`<!DOCTYPE html>`**
   - Declaração do tipo de documento
   - Informa ao navegador que este é um documento HTML5
   - Deve ser a primeira linha do arquivo

2. **`<html>`**
   - Elemento raiz do documento
   - Contém todo o conteúdo HTML
   - Atributo `lang` indica o idioma (importante para acessibilidade e SEO)

3. **`<head>`**
   - Contém **metadados** (informações sobre o documento)
   - Não é exibido visualmente na página
   - Inclui: título, charset, links para CSS, scripts, etc.

4. **`<body>`**
   - Contém todo o **conteúdo visível** da página
   - Textos, imagens, links, formulários, etc.

### Tags, Elementos e Atributos

#### Tags
São as palavras-chave entre `<` e `>`:
- `<p>` - tag de abertura de parágrafo
- `</p>` - tag de fechamento de parágrafo

#### Elementos
É a combinação de tag de abertura + conteúdo + tag de fechamento:
```html
<p>Este é um parágrafo</p>
```

#### Atributos
Fornecem informações adicionais sobre elementos:
```html
<img src="imagem.jpg" alt="Descrição da imagem">
```
- `src` e `alt` são atributos do elemento `<img>`

### Elementos HTML Básicos

#### Títulos (Headings)
```html
<h1>Título Principal (maior importância)</h1>
<h2>Subtítulo</h2>
<h3>Sub-subtítulo</h3>
<h4>Título de nível 4</h4>
<h5>Título de nível 5</h5>
<h6>Título de nível 6 (menor importância)</h6>
```

#### Parágrafos
```html
<p>Este é um parágrafo de texto. Pode conter várias linhas, mas será exibido como um bloco contínuo.</p>
```

#### Links
```html
<a href="https://www.exemplo.com">Texto do Link</a>
```

#### Imagens
```html
<img src="caminho/para/imagem.jpg" alt="Descrição da imagem">
```

#### Listas

**Lista não ordenada:**
```html
<ul>
    <li>Item 1</li>
    <li>Item 2</li>
    <li>Item 3</li>
</ul>
```

**Lista ordenada:**
```html
<ol>
    <li>Primeiro item</li>
    <li>Segundo item</li>
    <li>Terceiro item</li>
</ol>
```

---

## 🎨 CSS: Folhas de Estilo em Cascata

### O que é CSS?

**CSS** (Cascading Style Sheets - Folhas de Estilo em Cascata) é uma linguagem usada para descrever a **aparência** e **formatação** de um documento escrito em HTML.

### Função do CSS

O CSS controla:
- **Cores** (texto, fundo, bordas)
- **Fontes** (tipo, tamanho, peso)
- **Layout** (posicionamento, espaçamento, alinhamento)
- **Responsividade** (adaptação a diferentes tamanhos de tela)
- **Animações** e **transições**

### Como CSS se Relaciona com HTML

O HTML fornece a **estrutura**, o CSS fornece a **estilo**:

```html
<!-- HTML -->
<h1>Meu Título</h1>
```

```css
/* CSS */
h1 {
    color: blue;
    font-size: 32px;
    text-align: center;
}
```

### Formas de Adicionar CSS

1. **CSS Inline** (dentro do elemento HTML)
```html
<h1 style="color: blue;">Título</h1>
```

2. **CSS Interno** (na seção `<head>`)
```html
<head>
    <style>
        h1 { color: blue; }
    </style>
</head>
```

3. **CSS Externo** (arquivo separado)
```html
<head>
    <link rel="stylesheet" href="estilo.css">
</head>
```

---

## ⚡ JavaScript: A Linguagem de Programação da Web

### O que é JavaScript?

**JavaScript** é uma linguagem de programação usada principalmente para criar **efeitos interativos** e **dinâmicos** dentro de navegadores web.

### Função do JavaScript

O JavaScript permite:
- **Atualizar conteúdo** dinamicamente sem recarregar a página
- **Controlar mídia** (vídeos, áudio)
- **Animar elementos** na página
- **Validar formulários** em tempo real
- **Responder a eventos** do usuário (cliques, teclas, movimentos do mouse)
- **Manipular o DOM** (Document Object Model)

### Como JavaScript se Relaciona com HTML

O HTML fornece a **estrutura**, o JavaScript fornece o **comportamento**:

```html
<!-- HTML -->
<button id="meuBotao">Clique em Mim</button>
<p id="mensagem"></p>
```

```javascript
// JavaScript
document.getElementById('meuBotao').addEventListener('click', function() {
    document.getElementById('mensagem').textContent = 'Botão clicado!';
});
```

### Formas de Adicionar JavaScript

1. **JavaScript Inline** (dentro do elemento HTML)
```html
<button onclick="alert('Olá!')">Clique</button>
```

2. **JavaScript Interno** (na seção `<head>` ou antes de `</body>`)
```html
<script>
    function minhaFuncao() {
        alert('Olá!');
    }
</script>
```

3. **JavaScript Externo** (arquivo separado)
```html
<script src="script.js"></script>
```

---

## 🔗 A Tríade Web: HTML, CSS e JavaScript

### Trabalhando em Conjunto

Estas três tecnologias trabalham juntas para criar experiências web completas:

```
┌─────────────────────────────────────────┐
│         PÁGINA WEB COMPLETA             │
├─────────────────────────────────────────┤
│  HTML → Estrutura e Conteúdo           │
│  CSS  → Aparência e Estilo             │
│  JS   → Interatividade e Comportamento │
└─────────────────────────────────────────┘
```

### Analogia da Casa

- **HTML** = A estrutura da casa (paredes, portas, janelas)
- **CSS** = A decoração e pintura da casa (cores, móveis, estilo)
- **JavaScript** = A eletricidade e automação (luzes, portas automáticas, sistemas)

### Exemplo Completo

```html
<!DOCTYPE html>
<html lang="pt-BR">
<head>
    <meta charset="UTF-8">
    <title>Exemplo Completo</title>
    <style>
        /* CSS */
        button {
            background-color: blue;
            color: white;
            padding: 10px 20px;
            border: none;
            cursor: pointer;
        }
        button:hover {
            background-color: darkblue;
        }
    </style>
</head>
<body>
    <!-- HTML -->
    <h1>Minha Página</h1>
    <button id="botao">Clique Aqui</button>
    <p id="resultado"></p>
    
    <script>
        // JavaScript
        document.getElementById('botao').addEventListener('click', function() {
            document.getElementById('resultado').textContent = 'Botão foi clicado!';
        });
    </script>
</body>
</html>
```

---

## 📚 Versões do HTML

### Evolução do HTML

- **HTML 1.0** (1993) - Primeira versão
- **HTML 2.0** (1995) - Padronização
- **HTML 3.2** (1997) - Suporte a tabelas e formulários
- **HTML 4.01** (1999) - Separação de estrutura e apresentação
- **XHTML** (2000) - HTML baseado em XML
- **HTML5** (2014) - Versão atual, com novos elementos semânticos

### HTML5: A Versão Moderna

HTML5 introduziu:
- Novos elementos semânticos (`<header>`, `<nav>`, `<main>`, `<article>`, `<section>`, `<footer>`)
- Suporte nativo a áudio e vídeo (`<audio>`, `<video>`)
- Canvas para gráficos (`<canvas>`)
- Melhorias em formulários (novos tipos de input)
- Melhor suporte para dispositivos móveis

---

## 🎯 Importância da Semântica

### O que é HTML Semântico?

HTML semântico significa usar as **tags corretas** para o **propósito correto**:
- `<header>` para cabeçalho, não `<div class="header">`
- `<nav>` para navegação, não `<div class="nav">`
- `<article>` para artigos, não `<div class="article">`

### Por que é Importante?

1. **Acessibilidade**: Leitores de tela compreendem melhor a estrutura
2. **SEO**: Mecanismos de busca indexam melhor o conteúdo
3. **Manutenção**: Código mais fácil de entender e manter
4. **Padrões**: Segue as melhores práticas da web moderna

---

## 🛠️ Ferramentas Essenciais

### Editores de Código

- **Visual Studio Code** (recomendado)
- **Sublime Text**
- **Atom**
- **Notepad++** (Windows)

### Navegadores

- **Chrome** / **Edge** (DevTools excelentes)
- **Firefox** (DevTools avançados)
- **Safari** (macOS/iOS)

### Ferramentas de Validação

- **W3C Validator** - Valida código HTML
- **DevTools do Navegador** - Inspeciona e depura código

---

## 📝 Resumo da Aula

Nesta aula, você aprendeu:

✅ **HTML** é a linguagem de marcação que estrutura páginas web  
✅ **Linguagens de marcação** usam tags para anotar e estruturar texto  
✅ **Frontend development** combina HTML, CSS e JavaScript  
✅ **HTML** fornece a estrutura, **CSS** o estilo, **JavaScript** a interatividade  
✅ **HTML semântico** usa tags apropriadas para melhor acessibilidade e SEO  
✅ **HTML5** é a versão moderna com novos elementos semânticos  

### Próximos Passos

Na próxima aula, você aprenderá sobre:
- Estrutura detalhada de um documento HTML
- Tags básicas e seus usos
- Atributos e seus valores
- Como criar seu primeiro arquivo HTML

---

## 🔍 Conceitos-Chave para Revisão

- **HTML**: Linguagem de marcação para estruturar conteúdo web
- **Tags**: Instruções entre `<` e `>` que definem elementos
- **Elementos**: Combinação de tag de abertura, conteúdo e tag de fechamento
- **Atributos**: Informações adicionais sobre elementos
- **Semântica**: Uso de tags apropriadas para o propósito correto
- **Frontend**: HTML + CSS + JavaScript trabalhando juntos

