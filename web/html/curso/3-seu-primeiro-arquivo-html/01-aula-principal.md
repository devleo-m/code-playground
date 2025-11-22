# Aula 3: Seu Primeiro Arquivo HTML - Conteúdo Principal

## 📝 Revisão da Aula Anterior

Antes de começarmos, vamos relembrar os conceitos fundamentais que você já aprendeu:

- **HTML** é uma linguagem de marcação que estrutura o conteúdo web
- Um documento HTML básico possui a estrutura: `<!DOCTYPE html>`, `<html>`, `<head>` e `<body>`
- As **tags** são elementos que definem a estrutura do conteúdo
- O HTML funciona como o esqueleto de uma página web

Agora vamos aprofundar nosso conhecimento sobre como criar e estruturar corretamente um arquivo HTML!

---

## 🚀 Criando Seu Primeiro Arquivo HTML

### Passo a Passo para Criar um Arquivo HTML

1. **Crie um arquivo de texto simples**
   - Use qualquer editor de texto (Notepad, VS Code, Sublime Text, etc.)
   - Salve o arquivo com a extensão `.html` (exemplo: `minha-pagina.html`)

2. **Adicione a estrutura básica**
   - Sempre comece com `<!DOCTYPE html>` para indicar que é um documento HTML5
   - Envolva todo o conteúdo na tag `<html>`
   - Adicione as seções `<head>` e `<body>`

3. **Salve e abra no navegador**
   - Clique duas vezes no arquivo ou arraste-o para o navegador
   - Você verá sua página renderizada!

### Exemplo de Arquivo HTML Mínimo

```html
<!DOCTYPE html>
<html lang="pt-BR">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Minha Primeira Página</title>
</head>
<body>
    <h1>Olá, Mundo!</h1>
    <p>Esta é minha primeira página HTML.</p>
</body>
</html>
```

### Explicação Linha por Linha

- **`<!DOCTYPE html>`**: Declara que este é um documento HTML5. Esta declaração deve ser a primeira linha do arquivo.
- **`<html lang="pt-BR">`**: Tag raiz do documento. O atributo `lang` indica o idioma (português do Brasil).
- **`<head>`**: Contém metadados sobre o documento (informações que não aparecem na página).
- **`<meta charset="UTF-8">`**: Define a codificação de caracteres (permite acentos e caracteres especiais).
- **`<meta name="viewport">`**: Configura a visualização em dispositivos móveis.
- **`<title>`**: Define o título que aparece na aba do navegador.
- **`<body>`**: Contém todo o conteúdo visível da página.

---

## 🏷️ Tags e Atributos: Os Blocos de Construção do HTML

### O que são Tags?

**Tags** são elementos HTML que definem a estrutura e o significado do conteúdo. Elas funcionam como "etiquetas" que informam ao navegador como interpretar e exibir cada parte do documento.

### Estrutura de uma Tag

As tags geralmente vêm em pares: uma tag de **abertura** e uma tag de **fechamento**.

```html
<tagname>Conteúdo aqui</tagname>
```

- **Tag de abertura**: `<tagname>` - indica o início do elemento
- **Conteúdo**: O texto ou outros elementos dentro da tag
- **Tag de fechamento**: `</tagname>` - indica o fim do elemento (note a barra `/`)

### Exemplos de Tags Comuns

```html
<!-- Título principal -->
<h1>Meu Título</h1>

<!-- Parágrafo -->
<p>Este é um parágrafo de texto.</p>

<!-- Link -->
<a href="https://www.exemplo.com">Clique aqui</a>

<!-- Imagem -->
<img src="foto.jpg" alt="Descrição da foto">

<!-- Lista não ordenada -->
<ul>
    <li>Item 1</li>
    <li>Item 2</li>
</ul>
```

### Tags Vazias (Self-Closing Tags)

Algumas tags não precisam de fechamento porque não contêm conteúdo. Elas são chamadas de **tags vazias** ou **self-closing tags**.

```html
<!-- Tag de imagem (não tem conteúdo interno) -->
<img src="foto.jpg" alt="Descrição">

<!-- Quebra de linha -->
<br>

<!-- Linha horizontal -->
<hr>

<!-- Input em formulário -->
<input type="text" name="usuario">
```

**Nota**: Em HTML5, você pode escrever tags vazias de duas formas:
- `<img src="foto.jpg">` (sem barra)
- `<img src="foto.jpg" />` (com barra - estilo XHTML)

Ambas funcionam, mas a primeira é mais comum em HTML5.

---

### O que são Atributos?

**Atributos** fornecem informações adicionais sobre um elemento HTML. Eles modificam o comportamento ou a aparência da tag e são sempre especificados na tag de abertura.

### Sintaxe de Atributos

```html
<tagname atributo="valor">Conteúdo</tagname>
```

### Tipos de Atributos

#### 1. Atributos Globais (Disponíveis em Todas as Tags)

```html
<!-- id: Identificador único -->
<div id="meu-elemento">Conteúdo</div>

<!-- class: Classe para estilização -->
<p class="destaque">Texto destacado</p>

<!-- style: Estilo inline (use com moderação) -->
<h1 style="color: blue;">Título Azul</h1>

<!-- title: Tooltip ao passar o mouse -->
<a href="#" title="Clique para mais informações">Link</a>

<!-- lang: Idioma do elemento -->
<p lang="en">This is English text</p>
```

#### 2. Atributos Específicos de Tags

```html
<!-- Atributo href na tag <a> -->
<a href="https://www.exemplo.com">Link</a>

<!-- Atributos src e alt na tag <img> -->
<img src="imagem.jpg" alt="Descrição da imagem" width="300" height="200">

<!-- Atributo type na tag <input> -->
<input type="email" name="email" required>

<!-- Atributo target na tag <a> -->
<a href="https://www.exemplo.com" target="_blank">Abrir em nova aba</a>
```

### Atributos Booleanos

Alguns atributos não precisam de valor - sua presença já indica que a funcionalidade está ativada.

```html
<!-- Atributo disabled -->
<input type="text" disabled>

<!-- Atributo required -->
<input type="email" required>

<!-- Atributo checked -->
<input type="checkbox" checked>

<!-- Atributo readonly -->
<input type="text" readonly>
```

**Nota**: Em HTML5, você pode escrever atributos booleanos de duas formas:
- `<input required>` (sem valor)
- `<input required="required">` (com valor)

Ambas funcionam, mas a primeira é mais comum.

### Exemplo Prático: Tags e Atributos Trabalhando Juntos

```html
<!DOCTYPE html>
<html lang="pt-BR">
<head>
    <meta charset="UTF-8">
    <title>Exemplo de Tags e Atributos</title>
</head>
<body>
    <h1 id="titulo-principal" class="destaque">Bem-vindo</h1>
    
    <p class="introducao">
        Este é um parágrafo com uma 
        <a href="https://www.exemplo.com" target="_blank" title="Visite nosso site">
            link externo
        </a>.
    </p>
    
    <img src="logo.png" 
         alt="Logo da empresa" 
         width="200" 
         height="100"
         class="logo">
    
    <form>
        <input type="email" 
               name="email" 
               placeholder="Digite seu email"
               required>
        
        <button type="submit" disabled>Enviar</button>
    </form>
</body>
</html>
```

---

## 🔤 Case Insensitivity: HTML Não É Sensível a Maiúsculas e Minúsculas

### O que é Case Insensitivity?

**Case insensitivity** significa que o HTML não diferencia entre letras maiúsculas e minúsculas. Isso significa que você pode escrever tags e atributos usando qualquer combinação de maiúsculas e minúsculas, e o navegador entenderá da mesma forma.

### Exemplos de Case Insensitivity

Todas essas formas funcionam identicamente:

```html
<!-- Todas essas tags funcionam da mesma forma -->
<HTML>
<Html>
<html>
<HTml>
<hTmL>
```

```html
<!-- Atributos também são case-insensitive -->
<IMG SRC="foto.jpg" ALT="Descrição">
<img src="foto.jpg" alt="Descrição">
<Img Src="foto.jpg" Alt="Descrição">
```

### Por que Usar Minúsculas?

Embora o HTML aceite qualquer combinação, **a convenção e boa prática é usar sempre minúsculas**:

1. **Legibilidade**: Código em minúsculas é mais fácil de ler
2. **Consistência**: Mantém o código uniforme e profissional
3. **Compatibilidade**: XHTML (versão mais rigorosa) exige minúsculas
4. **Padrão da Indústria**: Todos os desenvolvedores profissionais usam minúsculas
5. **Ferramentas**: Muitas ferramentas e validadores esperam minúsculas

### Exemplo: Comparação de Estilos

**❌ Não Recomendado (Mistura de Maiúsculas):**
```html
<HTML>
<HEAD>
<TITLE>Minha Página</TITLE>
</HEAD>
<BODY>
<H1>Bem-vindo</H1>
<P>Este é um parágrafo.</P>
</BODY>
</HTML>
```

**✅ Recomendado (Tudo em Minúsculas):**
```html
<html>
<head>
<title>Minha Página</title>
</head>
<body>
<h1>Bem-vindo</h1>
<p>Este é um parágrafo.</p>
</body>
</html>
```

### Exceção: Valores de Atributos

**Importante**: Embora os **nomes** de tags e atributos sejam case-insensitive, os **valores** de alguns atributos podem ser case-sensitive:

```html
<!-- O valor do atributo type é case-sensitive em alguns casos -->
<input type="email">  <!-- Correto -->
<input type="EMAIL">  <!-- Pode não funcionar como esperado -->

<!-- URLs são case-sensitive -->
<a href="Pagina.html">Link</a>  <!-- Diferente de pagina.html -->
```

---

## 🔣 HTML Entities: Caracteres Especiais

### O que são HTML Entities?

**HTML Entities** (entidades HTML) são códigos especiais usados para representar caracteres que têm significado especial no HTML ou que são difíceis de digitar diretamente.

### Por que Usar Entities?

Alguns caracteres têm significados especiais no HTML:
- `<` e `>` são usados para tags
- `&` é usado para iniciar entities
- `"` e `'` são usados para valores de atributos

Se você tentar usar esses caracteres diretamente, o navegador pode interpretá-los incorretamente.

### Sintaxe de Entities

As entities HTML começam com `&` e terminam com `;`:

```html
&nome;     <!-- Entity nomeada -->
&#número;  <!-- Entity numérica -->
```

### Entities Mais Comuns

#### Caracteres Especiais

```html
<!-- Espaço não separável -->
&nbsp;     <!-- Espaço que não quebra linha -->

<!-- Aspas -->
&quot;     <!-- Aspas duplas (") -->
&apos;     <!-- Aspa simples (') -->

<!-- Símbolos matemáticos -->
&lt;        <!-- Menor que (<) -->
&gt;        <!-- Maior que (>) -->
&amp;       <!-- E comercial (&) -->
&times;    <!-- Sinal de multiplicação (×) -->
&divide;   <!-- Sinal de divisão (÷) -->
&plusmn;   <!-- Mais ou menos (±) -->
```

#### Símbolos de Moeda

```html
&euro;     <!-- Euro (€) -->
&pound;    <!-- Libra (£) -->
&yen;      <!-- Iene (¥) -->
&cent;     <!-- Centavo (¢) -->
```

#### Símbolos de Direitos Autorais

```html
&copy;     <!-- Copyright (©) -->
&reg;      <!-- Marca registrada (®) -->
&trade;    <!-- Marca comercial (™) -->
```

#### Símbolos de Setas

```html
&larr;     <!-- Seta esquerda (←) -->
&rarr;     <!-- Seta direita (→) -->
&uarr;     <!-- Seta para cima (↑) -->
&darr;     <!-- Seta para baixo (↓) -->
```

#### Acentos e Caracteres Especiais (Português)

```html
&aacute;   <!-- á -->
&Aacute;   <!-- Á -->
&agrave;   <!-- à -->
&Agrave;   <!-- À -->
&atilde;   <!-- ã -->
&Atilde;   <!-- Ã -->
&acirc;    <!-- â -->
&Acirc;    <!-- Â -->
&eacute;   <!-- é -->
&Eacute;   <!-- É -->
&ecirc;    <!-- ê -->
&Ecirc;    <!-- Ê -->
&iacute;   <!-- í -->
&Iacute;   <!-- Í -->
&oacute;   <!-- ó -->
&Oacute;   <!-- Ó -->
&otilde;   <!-- õ -->
&Otilde;   <!-- Õ -->
&uacute;   <!-- ú -->
&Uacute;   <!-- Ú -->
&ccedil;   <!-- ç -->
&Ccedil;   <!-- Ç -->
```

### Entities Numéricas

Você também pode usar números decimais ou hexadecimais:

```html
<!-- Decimal -->
&#169;     <!-- © (copyright) -->
&#174;     <!-- ® (registered) -->

<!-- Hexadecimal -->
&#xA9;     <!-- © (copyright) -->
&#xAE;     <!-- ® (registered) -->
```

### Exemplos Práticos

```html
<!DOCTYPE html>
<html lang="pt-BR">
<head>
    <meta charset="UTF-8">
    <title>Exemplo de Entities</title>
</head>
<body>
    <h1>Usando HTML Entities</h1>
    
    <p>
        Para escrever &lt;div&gt; em HTML, use entities: 
        &amp;lt;div&amp;gt;
    </p>
    
    <p>
        Preço: &euro;50,00 ou R$ 100,00
    </p>
    
    <p>
        Copyright &copy; 2024 Minha Empresa&trade;
    </p>
    
    <p>
        A fórmula matemática: 5 &times; 3 = 15
    </p>
    
    <p>
        Caracteres especiais: &aacute; &eacute; &iacute; &oacute; &uacute; &ccedil;
    </p>
    
    <p>
        Setas de navegação: &larr; Anterior | Próximo &rarr;
    </p>
</body>
</html>
```

### Quando Usar Entities?

**Use entities quando:**
- Você precisa exibir caracteres que têm significado especial no HTML (`<`, `>`, `&`)
- Você quer garantir que caracteres especiais sejam exibidos corretamente em qualquer navegador
- Você está trabalhando com conteúdo que pode não ter codificação UTF-8 adequada

**Não precisa usar entities quando:**
- Você está usando UTF-8 (que suporta a maioria dos caracteres diretamente)
- Você está escrevendo texto normal em português (UTF-8 já suporta acentos)

**Exemplo:**
```html
<!-- Com UTF-8, você pode escrever diretamente -->
<p>Café & Pão</p>

<!-- Ou usar entities (funciona igual) -->
<p>Caf&eacute; &amp; P&atilde;o</p>
```

---

## 💬 HTML Comments: Comentários no Código

### O que são Comentários HTML?

**Comentários HTML** são notas que você adiciona ao código para:
- Explicar o que o código faz
- Fazer lembretes para você mesmo
- Documentar decisões de desenvolvimento
- Temporariamente desabilitar partes do código
- Comunicar com outros desenvolvedores

### Importante: Comentários Não Aparecem no Navegador

Os comentários são **visíveis apenas no código-fonte**. Eles não são exibidos na página renderizada pelo navegador, mas podem ser vistos por qualquer pessoa que inspecione o código-fonte da página.

### Sintaxe de Comentários

```html
<!-- Este é um comentário HTML -->
```

Os comentários começam com `<!--` e terminam com `-->`. Tudo entre esses marcadores é ignorado pelo navegador.

### Exemplos de Uso

#### 1. Comentários Simples

```html
<!DOCTYPE html>
<html>
<head>
    <title>Minha Página</title>
</head>
<body>
    <!-- Este é o cabeçalho principal -->
    <h1>Bem-vindo</h1>
    
    <!-- Este parágrafo contém informações importantes -->
    <p>Conteúdo da página</p>
</body>
</html>
```

#### 2. Comentários Multilinha

```html
<!--
    Este é um comentário
    que ocupa múltiplas
    linhas de código.
    
    Útil para documentação
    extensa ou explicações
    detalhadas.
-->
<h1>Título</h1>
```

#### 3. Comentários para Desabilitar Código

```html
<!-- Temporariamente desabilitado
<h1>Título Antigo</h1>
-->

<h1>Novo Título</h1>
```

#### 4. Comentários para Organização

```html
<!DOCTYPE html>
<html>
<head>
    <!-- ============================================ -->
    <!-- METADADOS E CONFIGURAÇÕES -->
    <!-- ============================================ -->
    <meta charset="UTF-8">
    <title>Minha Página</title>
</head>
<body>
    <!-- ============================================ -->
    <!-- CABEÇALHO -->
    <!-- ============================================ -->
    <header>
        <h1>Meu Site</h1>
    </header>
    
    <!-- ============================================ -->
    <!-- CONTEÚDO PRINCIPAL -->
    <!-- ============================================ -->
    <main>
        <p>Conteúdo aqui</p>
    </main>
    
    <!-- ============================================ -->
    <!-- RODAPÉ -->
    <!-- ============================================ -->
    <footer>
        <p>Copyright 2024</p>
    </footer>
</body>
</html>
```

#### 5. Comentários Condicionais (IE - Legado)

```html
<!--[if IE]>
    <p>Você está usando Internet Explorer</p>
<![endif]-->
```

**Nota**: Comentários condicionais eram usados para versões antigas do Internet Explorer. Não são mais necessários hoje em dia.

### Boas Práticas com Comentários

1. **Seja Claro e Conciso**
   ```html
   <!-- ✅ Bom -->
   <!-- Menu de navegação principal -->
   
   <!-- ❌ Ruim -->
   <!-- menu -->
   ```

2. **Documente Decisões Importantes**
   ```html
   <!-- Usamos div aqui porque precisamos de um container
        sem semântica específica para o layout flexbox -->
   <div class="container">
   ```

3. **Mantenha Comentários Atualizados**
   - Se você mudar o código, atualize os comentários também
   - Remova comentários obsoletos

4. **Use Comentários para Debugging**
   ```html
   <!-- TODO: Adicionar validação de formulário -->
   <!-- FIXME: Corrigir problema de layout no mobile -->
   <!-- NOTE: Esta seção será refatorada na próxima versão -->
   ```

5. **Não Comente Código Óbvio**
   ```html
   <!-- ❌ Desnecessário -->
   <!-- Parágrafo -->
   <p>Texto</p>
   
   <!-- ✅ Útil -->
   <!-- Este parágrafo contém a descrição do produto,
        que é carregada dinamicamente via JavaScript -->
   <p id="descricao-produto"></p>
   ```

### Comentários Aninhados (Não Funcionam!)

**Importante**: Você **não pode** aninhar comentários HTML. O primeiro `-->` fecha o comentário, mesmo que esteja dentro de outro.

```html
<!-- ❌ Isso NÃO funciona -->
<!--
    Comentário externo
    <!-- Comentário interno -->
    Mais texto
-->

<!-- ✅ Solução: Use comentários separados -->
<!-- Comentário externo -->
<!-- Comentário interno -->
```

---

## ⚪ Whitespaces: Espaços em Branco no HTML

### O que são Whitespaces?

**Whitespaces** (espaços em branco) são caracteres invisíveis que incluem:
- **Espaços** (barra de espaço)
- **Tabs** (tabulação)
- **Quebras de linha** (Enter)
- **Retornos de carro**

### Como Navegadores Tratam Whitespaces

Os navegadores têm uma regra importante: **múltiplos whitespaces consecutivos são colapsados em um único espaço**.

### Exemplos de Colapso de Whitespaces

```html
<!-- Código HTML -->
<p>Olá     Mundo</p>

<!-- Renderizado no navegador -->
Olá Mundo
```

```html
<!-- Código HTML com múltiplas quebras de linha -->
<p>Linha 1


Linha 2</p>

<!-- Renderizado no navegador -->
Linha 1 Linha 2
```

### Espaços em Branco Preservados

Algumas tags preservam whitespaces exatamente como estão no código:

#### 1. Tag `<pre>` (Texto Pré-formatado)

```html
<pre>
    Este texto
    mantém todos os espaços
    e quebras de linha
    exatamente como estão.
</pre>
```

#### 2. Tag `<code>` (dentro de `<pre>`)

```html
<pre><code>
function exemplo() {
    console.log("Olá");
}
</code></pre>
```

#### 3. Atributo `white-space: pre` (CSS)

```html
<p style="white-space: pre;">
    Este parágrafo
    também preserva
    os espaços.
</p>
```

### Espaços em Branco em Atributos

Os espaços no início e fim dos valores de atributos são geralmente ignorados:

```html
<!-- Estes são equivalentes -->
<img src="foto.jpg" alt="Descrição">
<img src=" foto.jpg " alt=" Descrição ">
```

Mas espaços **dentro** do valor são preservados:

```html
<!-- Estes são diferentes -->
<img alt="Descrição da foto">
<img alt="Descrição  da  foto">  <!-- Dois espaços -->
```

### Whitespaces para Formatação do Código

Embora os whitespaces não afetem a renderização (na maioria dos casos), eles são **essenciais para a legibilidade do código**:

```html
<!-- ❌ Difícil de ler -->
<html><head><title>Página</title></head><body><h1>Título</h1><p>Texto</p></body></html>

<!-- ✅ Fácil de ler -->
<html>
    <head>
        <title>Página</title>
    </head>
    <body>
        <h1>Título</h1>
        <p>Texto</p>
    </body>
</html>
```

### Indentação e Boas Práticas

A **indentação** (espaçamento no início das linhas) ajuda a visualizar a hierarquia do código:

```html
<!DOCTYPE html>
<html>
    <head>
        <meta charset="UTF-8">
        <title>Minha Página</title>
    </head>
    <body>
        <header>
            <h1>Título</h1>
            <nav>
                <ul>
                    <li>Item 1</li>
                    <li>Item 2</li>
                </ul>
            </nav>
        </header>
        <main>
            <article>
                <h2>Subtítulo</h2>
                <p>Conteúdo do artigo</p>
            </article>
        </main>
    </body>
</html>
```

**Dica**: Use 2 ou 4 espaços para indentação (escolha um e seja consistente). Tabs também funcionam, mas espaços são mais universais.

### Entity para Espaço Não Separável

Se você precisar de um espaço que **não pode ser quebrado** (útil para evitar que palavras sejam separadas), use a entity `&nbsp;`:

```html
<p>Dr.&nbsp;Silva</p>  <!-- "Dr." e "Silva" não serão separados -->
<p>R$&nbsp;100,00</p>  <!-- "R$" e "100,00" não serão separados -->
```

### Exemplo Prático: Whitespaces em Ação

```html
<!DOCTYPE html>
<html lang="pt-BR">
<head>
    <meta charset="UTF-8">
    <title>Exemplo de Whitespaces</title>
</head>
<body>
    <h1>Demonstração de Whitespaces</h1>
    
    <!-- Múltiplos espaços são colapsados -->
    <p>Olá     Mundo</p>
    <!-- Renderiza: "Olá Mundo" -->
    
    <!-- Quebras de linha são colapsadas -->
    <p>Linha 1
    
    
    Linha 2</p>
    <!-- Renderiza: "Linha 1 Linha 2" -->
    
    <!-- Tag <pre> preserva whitespaces -->
    <pre>
        Este texto
        mantém todos os
        espaços e quebras.
    </pre>
    
    <!-- Espaço não separável -->
    <p>Preço: R$&nbsp;50,00</p>
    
    <!-- Código formatado -->
    <pre><code>
function exemplo() {
    return "Olá";
}
    </code></pre>
</body>
</html>
```

---

## 📋 Resumo dos Conceitos

### Tags e Atributos
- **Tags** definem a estrutura do conteúdo
- **Atributos** fornecem informações adicionais sobre elementos
- Tags podem ser de abertura/fechamento ou self-closing
- Atributos podem ser globais ou específicos de tags

### Case Insensitivity
- HTML não diferencia maiúsculas de minúsculas
- **Mas use sempre minúsculas** por convenção e boas práticas
- Valores de atributos podem ser case-sensitive

### HTML Entities
- Códigos especiais para caracteres com significado especial
- Sintaxe: `&nome;` ou `&#número;`
- Úteis para `<`, `>`, `&`, símbolos e caracteres especiais
- Com UTF-8, muitos caracteres podem ser escritos diretamente

### HTML Comments
- Sintaxe: `<!-- comentário -->`
- Não aparecem no navegador, apenas no código-fonte
- Use para documentação, organização e debugging
- Não podem ser aninhados

### Whitespaces
- Espaços, tabs e quebras de linha
- Múltiplos whitespaces são colapsados em um único espaço
- Use para formatar e indentar código (legibilidade)
- Tag `<pre>` preserva whitespaces exatamente como estão

---

## 🎯 Próximos Passos

Agora que você entende os fundamentos de tags, atributos, entities, comentários e whitespaces, você está pronto para:

1. Criar arquivos HTML mais complexos
2. Estruturar conteúdo de forma semântica
3. Adicionar metadados e configurações adequadas
4. Escrever código limpo e bem documentado

Na próxima aula, vamos explorar elementos de texto e formatação!

---

## 📚 Recursos Adicionais

- [MDN Web Docs - HTML Elements](https://developer.mozilla.org/pt-BR/docs/Web/HTML/Element)
- [W3C HTML Validator](https://validator.w3.org/)
- [HTML Entities Reference](https://www.w3schools.com/html/html_entities.asp)
- [HTML Comments Best Practices](https://developer.mozilla.org/pt-BR/docs/Learn/HTML/Introduction_to_HTML/Getting_started#comentários_html)

