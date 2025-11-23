# Aula 6: Embedding Media e Formulários - Conteúdo Principal

## 📝 Revisão da Aula Anterior

Antes de começarmos, vamos relembrar os conceitos fundamentais que você já aprendeu:

- **Agrupamento de elementos** usando `<div>` e `<span>`
- **Atributos HTML** e sua importância para identificação e estilização
- **Listas** ordenadas, não ordenadas e de definição
- **Tabelas** para dados estruturados
- **Elementos semânticos** para melhor organização do conteúdo

Agora vamos aprender a incorporar mídia (imagens, áudio e vídeo) em páginas web e criar formulários interativos para coletar dados dos usuários!

---

## 🖼️ Embedding Media: Incorporando Mídia em HTML

### O que é Embedding Media?

**Embedding Media** (incorporação de mídia) é o processo de integrar conteúdo multimídia como imagens, áudio e vídeo diretamente em uma página web. Isso permite que os usuários visualizem ou ouçam esses elementos sem precisar baixá-los separadamente ou navegar para outro site.

### Por que Usar Embedding Media?

- **Experiência do usuário melhorada**: Todo o conteúdo fica em um único lugar
- **Performance**: Navegadores podem otimizar o carregamento de mídia
- **Acessibilidade**: Permite adicionar descrições e controles para leitores de tela
- **SEO**: Mídia incorporada pode melhorar o posicionamento em mecanismos de busca
- **Controle**: Você pode controlar como a mídia é exibida e reproduzida

---

## 🖼️ Imagens em HTML

### O Elemento `<img>`: Imagens Simples

O elemento `<img>` é usado para incorporar imagens em uma página HTML. É um elemento **vazio** (self-closing), o que significa que não possui tag de fechamento.

**Sintaxe básica:**
```html
<img src="caminho/para/imagem.jpg" alt="Descrição da imagem">
```

### Atributos Essenciais do `<img>`

#### `src` (source - obrigatório)
Especifica o caminho para o arquivo de imagem.

```html
<!-- Imagem local -->
<img src="imagens/foto.jpg" alt="Minha foto">

<!-- Imagem de URL externa -->
<img src="https://exemplo.com/imagem.png" alt="Imagem externa">
```

#### `alt` (alternative text - obrigatório para acessibilidade)
Fornece texto alternativo que descreve a imagem. É usado quando:
- A imagem não pode ser carregada
- Usuários com leitores de tela precisam entender o conteúdo
- A conexão está lenta e a imagem ainda não carregou

```html
<img src="gato.jpg" alt="Um gato laranja brincando com uma bola">
```

**Boas práticas para `alt`:**
- Seja descritivo, mas conciso
- Descreva o conteúdo e função da imagem
- Se a imagem é decorativa, use `alt=""` (vazio)
- Não comece com "Imagem de..." ou "Foto de..."

#### `width` e `height`
Especificam as dimensões da imagem em pixels.

```html
<img src="foto.jpg" alt="Minha foto" width="300" height="200">
```

**Importante:**
- Sempre especifique `width` e `height` para evitar **Cumulative Layout Shift (CLS)**
- Isso ajuda o navegador a reservar espaço antes da imagem carregar
- Melhora a performance e experiência do usuário

#### `title`
Fornece informações adicionais que aparecem como tooltip quando o usuário passa o mouse sobre a imagem.

```html
<img src="foto.jpg" alt="Minha foto" title="Foto tirada em 2024">
```

### Formatos de Imagem Comuns

1. **JPEG/JPG**: Melhor para fotografias, suporta milhões de cores
2. **PNG**: Melhor para imagens com transparência ou gráficos simples
3. **GIF**: Suporta animação, mas limitado a 256 cores
4. **WebP**: Formato moderno com melhor compressão (suporte limitado em navegadores antigos)
5. **SVG**: Gráficos vetoriais escaláveis, perfeito para ícones e ilustrações

### Exemplo Completo de Imagem

```html
<img 
    src="imagens/paisagem.jpg" 
    alt="Paisagem montanhosa ao pôr do sol" 
    width="800" 
    height="600"
    title="Foto de uma montanha ao entardecer"
    loading="lazy"
>
```

---

## 🎨 `<img>` vs. `<figure>`: Quando Usar Cada Um?

### O Elemento `<img>`: Imagem Simples

Use `<img>` quando:
- A imagem é parte do conteúdo do texto
- A imagem é decorativa ou ilustrativa simples
- Não precisa de legenda ou contexto adicional

```html
<p>
    Veja esta foto: <img src="foto.jpg" alt="Minha foto"> Ela foi tirada ontem.
</p>
```

### O Elemento `<figure>`: Conteúdo Autocontido

O elemento `<figure>` representa conteúdo autocontido, como uma imagem, ilustração, diagrama, código, etc., que é referenciado como uma unidade. Geralmente é usado com `<figcaption>` para fornecer uma legenda.

**Sintaxe:**
```html
<figure>
    <img src="diagrama.jpg" alt="Diagrama do sistema">
    <figcaption>Figura 1: Diagrama da arquitetura do sistema</figcaption>
</figure>
```

**Use `<figure>` quando:**
- A imagem precisa de uma legenda ou descrição adicional
- A imagem é uma figura, diagrama ou ilustração que faz referência ao texto
- Você quer agrupar a imagem com sua legenda semanticamente
- A imagem pode ser movida para outro lugar sem perder contexto

**Exemplo completo:**
```html
<article>
    <h2>Como Funciona a Fotosíntese</h2>
    <p>O processo de fotossíntese é complexo...</p>
    
    <figure>
        <img 
            src="diagramas/fotossintese.png" 
            alt="Diagrama mostrando o processo de fotossíntese nas plantas"
            width="600"
            height="400"
        >
        <figcaption>
            Figura 1: Processo completo de fotossíntese, mostrando a absorção 
            de luz solar, água e CO₂, e a produção de glicose e oxigênio.
        </figcaption>
    </figure>
    
    <p>Como podemos ver na figura acima...</p>
</article>
```

**Múltiplas imagens em uma figura:**
```html
<figure>
    <img src="antes.jpg" alt="Estado antes da reforma">
    <img src="depois.jpg" alt="Estado depois da reforma">
    <figcaption>Comparação antes e depois da reforma da cozinha</figcaption>
</figure>
```

---

## ⚡ Priority Hints: Priorizando Recursos

### O que são Priority Hints?

**Priority Hints** (dicas de prioridade) permitem que desenvolvedores indiquem a prioridade relativa de carregamento de recursos como imagens. Isso ajuda o navegador a decidir quais imagens carregar primeiro, potencialmente melhorando os tempos de carregamento da página e a experiência do usuário.

### O Atributo `fetchpriority`

O atributo `fetchpriority` indica ao navegador a importância de um recurso em relação a outros recursos do mesmo tipo.

**Valores possíveis:**
- `high`: Recurso de alta prioridade (ex: imagem hero acima da dobra)
- `low`: Recurso de baixa prioridade (ex: imagens abaixo da dobra)
- `auto`: Prioridade padrão (comportamento normal do navegador)

**Sintaxe:**
```html
<!-- Imagem importante (hero image) -->
<img 
    src="hero.jpg" 
    alt="Produto principal" 
    fetchpriority="high"
>

<!-- Imagem menos importante (abaixo da dobra) -->
<img 
    src="galeria-1.jpg" 
    alt="Galeria de produtos" 
    fetchpriority="low"
    loading="lazy"
>
```

**Quando usar:**
- **`high`**: Para imagens críticas que aparecem acima da dobra (above the fold)
- **`low`**: Para imagens que aparecem abaixo da dobra ou em galerias
- **`auto`**: Deixe o navegador decidir (padrão)

**Exemplo prático:**
```html
<header>
    <!-- Imagem hero - alta prioridade -->
    <img 
        src="hero-banner.jpg" 
        alt="Banner principal do site" 
        fetchpriority="high"
        width="1920"
        height="1080"
    >
</header>

<main>
    <section>
        <h1>Produtos em Destaque</h1>
        <!-- Imagens de produtos - prioridade normal/baixa -->
        <img 
            src="produto-1.jpg" 
            alt="Produto 1" 
            fetchpriority="low"
            loading="lazy"
        >
        <img 
            src="produto-2.jpg" 
            alt="Produto 2" 
            fetchpriority="low"
            loading="lazy"
        >
    </section>
</main>
```

**Combinando com `loading="lazy"`:**
```html
<!-- Carregamento imediato (alta prioridade) -->
<img src="importante.jpg" alt="Imagem importante" fetchpriority="high">

<!-- Carregamento preguiçoso (baixa prioridade) -->
<img 
    src="galeria.jpg" 
    alt="Imagem da galeria" 
    fetchpriority="low"
    loading="lazy"
>
```

---

## 🎵 Áudio em HTML

### O Elemento `<audio>`

O elemento `<audio>` é usado para incorporar conteúdo de áudio em uma página HTML. Permite que os usuários ouçam música, podcasts ou outros arquivos de áudio diretamente na página.

**Sintaxe básica:**
```html
<audio src="audio/musica.mp3" controls></audio>
```

### Atributos do `<audio>`

#### `src`
Especifica o caminho para o arquivo de áudio.

```html
<audio src="audio/musica.mp3" controls></audio>
```

#### `controls`
Exibe controles de reprodução (play, pause, volume, etc.).

```html
<audio src="audio/musica.mp3" controls></audio>
```

#### `autoplay`
Inicia a reprodução automaticamente quando a página carrega.

```html
<audio src="audio/musica.mp3" controls autoplay></audio>
```

**⚠️ Atenção:** Muitos navegadores bloqueiam autoplay de áudio por padrão para melhorar a experiência do usuário.

#### `loop`
Reproduz o áudio em loop (repetição contínua).

```html
<audio src="audio/musica.mp3" controls loop></audio>
```

#### `muted`
Inicia o áudio mutado.

```html
<audio src="audio/musica.mp3" controls muted></audio>
```

#### `preload`
Especifica como o navegador deve carregar o áudio:
- `none`: Não pré-carrega o áudio
- `metadata`: Carrega apenas metadados (duração, dimensões, etc.)
- `auto`: Carrega o áudio inteiro (padrão)

```html
<audio src="audio/musica.mp3" controls preload="metadata"></audio>
```

### Múltiplas Fontes com `<source>`

Para garantir compatibilidade com diferentes navegadores, você pode fornecer múltiplos formatos de áudio usando o elemento `<source>` dentro de `<audio>`.

**Formatos de áudio comuns:**
- **MP3**: Suporte universal, boa compressão
- **OGG**: Código aberto, boa qualidade
- **WAV**: Alta qualidade, arquivos grandes

**Sintaxe:**
```html
<audio controls>
    <source src="audio/musica.mp3" type="audio/mpeg">
    <source src="audio/musica.ogg" type="audio/ogg">
    <source src="audio/musica.wav" type="audio/wav">
    Seu navegador não suporta o elemento de áudio.
</audio>
```

**Exemplo completo:**
```html
<audio controls preload="metadata">
    <source src="podcast-episodio-1.mp3" type="audio/mpeg">
    <source src="podcast-episodio-1.ogg" type="audio/ogg">
    <p>
        Seu navegador não suporta áudio HTML5. 
        <a href="podcast-episodio-1.mp3">Baixe o arquivo</a>.
    </p>
</audio>
```

### Texto Alternativo para Áudio

O texto dentro do elemento `<audio>` (quando não há suporte) serve como fallback e também pode ser usado para acessibilidade.

```html
<audio controls>
    <source src="narracao.mp3" type="audio/mpeg">
    <p>
        Áudio: Narração sobre a história do HTML.
        <a href="narracao.mp3">Baixar áudio</a> ou 
        <a href="transcricao.txt">Ler transcrição</a>.
    </p>
</audio>
```

---

## 🎬 Vídeo em HTML

### O Elemento `<video>`

O elemento `<video>` é usado para incorporar conteúdo de vídeo diretamente em uma página HTML. Permite exibir arquivos de vídeo, controlar a reprodução e fornecer opções para os usuários interagirem com o vídeo.

**Sintaxe básica:**
```html
<video src="video/filme.mp4" controls width="640" height="360"></video>
```

### Atributos do `<video>`

#### `src`
Especifica o caminho para o arquivo de vídeo.

```html
<video src="video/filme.mp4" controls></video>
```

#### `controls`
Exibe controles de reprodução (play, pause, volume, tela cheia, etc.).

```html
<video src="video/filme.mp4" controls></video>
```

#### `width` e `height`
Especificam as dimensões do player de vídeo.

```html
<video src="video/filme.mp4" controls width="1280" height="720"></video>
```

**Importante:** Sempre especifique `width` e `height` para evitar layout shift.

#### `autoplay`
Inicia a reprodução automaticamente.

```html
<video src="video/filme.mp4" controls autoplay></video>
```

**⚠️ Atenção:** Navegadores modernos geralmente bloqueiam autoplay com áudio. Use `muted` junto com `autoplay` para garantir que funcione.

```html
<video src="video/filme.mp4" controls autoplay muted></video>
```

#### `loop`
Reproduz o vídeo em loop.

```html
<video src="video/filme.mp4" controls loop></video>
```

#### `muted`
Inicia o vídeo mutado.

```html
<video src="video/filme.mp4" controls muted></video>
```

#### `poster`
Especifica uma imagem a ser exibida antes do vídeo começar a reproduzir.

```html
<video 
    src="video/filme.mp4" 
    controls 
    poster="imagens/thumbnail.jpg"
></video>
```

#### `preload`
Especifica como o navegador deve carregar o vídeo:
- `none`: Não pré-carrega
- `metadata`: Carrega apenas metadados
- `auto`: Carrega o vídeo inteiro (padrão)

```html
<video src="video/filme.mp4" controls preload="metadata"></video>
```

### Múltiplas Fontes com `<source>`

Assim como com áudio, você pode fornecer múltiplos formatos de vídeo para compatibilidade.

**Formatos de vídeo comuns:**
- **MP4**: Suporte universal, boa compressão
- **WebM**: Código aberto, boa qualidade
- **OGG**: Alternativa de código aberto

**Sintaxe:**
```html
<video controls width="1280" height="720" poster="thumbnail.jpg">
    <source src="video/filme.mp4" type="video/mp4">
    <source src="video/filme.webm" type="video/webm">
    <source src="video/filme.ogv" type="video/ogg">
    Seu navegador não suporta o elemento de vídeo.
</video>
```

**Exemplo completo:**
```html
<video 
    controls 
    width="1280" 
    height="720" 
    poster="imagens/thumbnail-video.jpg"
    preload="metadata"
>
    <source src="tutorial-html.mp4" type="video/mp4">
    <source src="tutorial-html.webm" type="video/webm">
    <p>
        Seu navegador não suporta vídeo HTML5.
        <a href="tutorial-html.mp4">Baixe o vídeo</a> ou
        <a href="transcricao-video.txt">Leia a transcrição</a>.
    </p>
</video>
```

### Legendas e Legendas com `<track>`

O elemento `<track>` permite adicionar legendas, legendas descritivas ou capítulos ao vídeo.

**Sintaxe:**
```html
<video controls>
    <source src="video.mp4" type="video/mp4">
    <track 
        kind="subtitles" 
        src="legendas-pt.vtt" 
        srclang="pt" 
        label="Português"
        default
    >
    <track 
        kind="subtitles" 
        src="legendas-en.vtt" 
        srclang="en" 
        label="English"
    >
</video>
```

**Tipos de `kind`:**
- `subtitles`: Legendas (tradução do diálogo)
- `captions`: Legendas descritivas (inclui sons)
- `descriptions`: Descrições de áudio para leitores de tela
- `chapters`: Capítulos para navegação

---

## 🖼️ iframe: Incorporando Conteúdo Externo

### O que é um iframe?

Um **iframe** (Inline Frame) é um elemento HTML que permite incorporar outro documento HTML dentro da página atual. É como criar uma "janela" dentro da sua página onde você pode exibir conteúdo de outra fonte.

### O Elemento `<iframe>`

**Sintaxe básica:**
```html
<iframe src="https://exemplo.com/pagina.html"></iframe>
```

### Atributos do `<iframe>`

#### `src`
Especifica a URL do documento a ser incorporado.

```html
<iframe src="https://www.youtube.com/embed/dQw4w9WgXcQ"></iframe>
```

#### `width` e `height`
Especificam as dimensões do iframe.

```html
<iframe 
    src="https://exemplo.com" 
    width="800" 
    height="600"
></iframe>
```

#### `title`
Fornece um título descritivo para acessibilidade.

```html
<iframe 
    src="https://exemplo.com" 
    width="800" 
    height="600"
    title="Mapa interativo da cidade"
></iframe>
```

#### `frameborder` (obsoleto em HTML5)
Controla a borda do iframe. Use CSS `border` em vez disso.

```html
<!-- Não recomendado -->
<iframe src="exemplo.com" frameborder="0"></iframe>

<!-- Recomendado -->
<iframe src="exemplo.com" style="border: none;"></iframe>
```

#### `sandbox`
Aplica restrições de segurança ao conteúdo do iframe.

```html
<iframe 
    src="conteudo-externo.html" 
    sandbox="allow-scripts allow-same-origin"
></iframe>
```

**Valores de `sandbox`:**
- `allow-scripts`: Permite JavaScript
- `allow-same-origin`: Permite acesso à mesma origem
- `allow-forms`: Permite formulários
- `allow-popups`: Permite pop-ups
- `allow-top-navigation`: Permite navegação no topo

### Exemplos Comuns de iframe

#### Incorporar Vídeo do YouTube

```html
<iframe 
    width="560" 
    height="315" 
    src="https://www.youtube.com/embed/VIDEO_ID" 
    title="Vídeo do YouTube"
    frameborder="0" 
    allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture" 
    allowfullscreen
></iframe>
```

#### Incorporar Mapa do Google Maps

```html
<iframe 
    src="https://www.google.com/maps/embed?pb=..." 
    width="600" 
    height="450" 
    style="border:0;" 
    allowfullscreen="" 
    loading="lazy" 
    referrerpolicy="no-referrer-when-downgrade"
    title="Localização no mapa"
></iframe>
```

#### Incorporar PDF

```html
<iframe 
    src="documento.pdf" 
    width="100%" 
    height="600"
    title="Documento PDF"
></iframe>
```

### Segurança com iframe

**⚠️ Importante:** Sempre considere a segurança ao usar iframes:

1. **Use `sandbox`** para conteúdo não confiável
2. **Valide URLs** antes de incorporar
3. **Use HTTPS** para conteúdo externo
4. **Considere CSP** (Content Security Policy) para restringir fontes

---

## 🔒 Content Security Policy (CSP)

### O que é CSP?

**Content Security Policy (CSP)** é um padrão de segurança introduzido para prevenir ataques de cross-site scripting (XSS), clickjacking e outras injeções de código. Funciona permitindo que você defina uma whitelist de fontes das quais o navegador tem permissão para carregar recursos.

### Como Funciona o CSP?

O CSP funciona restringindo de onde recursos podem ser carregados:
- Scripts (JavaScript)
- Estilos (CSS)
- Imagens
- Fontes
- iframes
- Mídia (áudio, vídeo)
- E outros recursos

### Implementando CSP

#### Via Meta Tag (Recomendado para testes)

```html
<head>
    <meta 
        http-equiv="Content-Security-Policy" 
        content="default-src 'self'; img-src 'self' https:; script-src 'self'"
    >
</head>
```

#### Via Header HTTP (Recomendado para produção)

Configure no servidor web (Apache, Nginx, etc.):

```
Content-Security-Policy: default-src 'self'; img-src 'self' https:; script-src 'self'
```

### Diretivas CSP Comuns

#### `default-src`
Define a política padrão para todos os tipos de recursos.

```html
<meta 
    http-equiv="Content-Security-Policy" 
    content="default-src 'self'"
>
```

#### `script-src`
Controla de onde scripts podem ser carregados.

```html
<!-- Permite apenas scripts do mesmo domínio -->
<meta 
    http-equiv="Content-Security-Policy" 
    content="script-src 'self'"
>

<!-- Permite scripts do mesmo domínio e de CDNs específicos -->
<meta 
    http-equiv="Content-Security-Policy" 
    content="script-src 'self' https://cdn.exemplo.com"
>
```

#### `style-src`
Controla de onde estilos podem ser carregados.

```html
<meta 
    http-equiv="Content-Security-Policy" 
    content="style-src 'self' 'unsafe-inline'"
>
```

#### `img-src`
Controla de onde imagens podem ser carregadas.

```html
<!-- Permite imagens do mesmo domínio e qualquer HTTPS -->
<meta 
    http-equiv="Content-Security-Policy" 
    content="img-src 'self' https:"
>
```

#### `frame-src` ou `child-src`
Controla de onde iframes podem ser carregados.

```html
<meta 
    http-equiv="Content-Security-Policy" 
    content="frame-src 'self' https://www.youtube.com"
>
```

#### `media-src`
Controla de onde mídia (áudio, vídeo) pode ser carregada.

```html
<meta 
    http-equiv="Content-Security-Policy" 
    content="media-src 'self' https://cdn.exemplo.com"
>
```

### Valores Especiais

- **`'self'`**: Permite recursos do mesmo domínio
- **`'unsafe-inline'`**: Permite JavaScript/CSS inline (não recomendado)
- **`'unsafe-eval'`**: Permite `eval()` (não recomendado)
- **`https:`**: Permite qualquer origem HTTPS
- **`'none'`**: Bloqueia todos os recursos desse tipo

### Exemplo Completo de CSP

```html
<head>
    <meta 
        http-equiv="Content-Security-Policy" 
        content="
            default-src 'self';
            script-src 'self' https://cdn.exemplo.com;
            style-src 'self' 'unsafe-inline';
            img-src 'self' https: data:;
            font-src 'self' https://fonts.googleapis.com;
            frame-src 'self' https://www.youtube.com;
            media-src 'self';
            connect-src 'self' https://api.exemplo.com;
        "
    >
</head>
```

### Report-Only Mode

Para testar CSP sem bloquear recursos, use `Content-Security-Policy-Report-Only`:

```html
<meta 
    http-equiv="Content-Security-Policy-Report-Only" 
    content="default-src 'self'"
>
```

Isso reporta violações sem bloqueá-las, permitindo ajustar a política antes de ativá-la.

---

## 📝 Formulários em HTML

### O que são Formulários?

**Formulários** em HTML são usados para coletar dados dos usuários. Eles fornecem uma maneira para os usuários inserirem informações como texto, senhas, seleções e enviar esses dados para um servidor para processamento.

### O Elemento `<form>`

O elemento `<form>` é o container que agrupa todos os campos de entrada e define como os dados serão enviados.

**Sintaxe básica:**
```html
<form action="/processar" method="post">
    <!-- Campos do formulário aqui -->
</form>
```

### Atributos do `<form>`

#### `action`
Especifica para onde os dados do formulário serão enviados (URL do servidor).

```html
<form action="/api/cadastro" method="post">
    <!-- campos -->
</form>
```

#### `method`
Especifica o método HTTP usado para enviar os dados:
- **`GET`**: Dados enviados na URL (visíveis, limitados)
- **`POST`**: Dados enviados no corpo da requisição (mais seguro, sem limite)

```html
<!-- GET: para buscas, filtros -->
<form action="/buscar" method="get">
    <input type="text" name="q" placeholder="Buscar...">
    <button type="submit">Buscar</button>
</form>

<!-- POST: para cadastros, envio de dados -->
<form action="/cadastrar" method="post">
    <!-- campos -->
</form>
```

#### `enctype`
Especifica como os dados são codificados ao enviar:
- `application/x-www-form-urlencoded`: Padrão (texto simples)
- `multipart/form-data`: Necessário para upload de arquivos
- `text/plain`: Apenas texto (não recomendado)

```html
<form action="/upload" method="post" enctype="multipart/form-data">
    <input type="file" name="arquivo">
    <button type="submit">Enviar</button>
</form>
```

#### `target`
Especifica onde abrir a resposta do formulário:
- `_self`: Na mesma janela (padrão)
- `_blank`: Em nova janela/aba
- `_parent`: No frame pai
- `_top`: Na janela principal

#### `novalidate`
Desabilita a validação HTML5 do formulário (útil quando você usa validação JavaScript customizada).

```html
<form action="/processar" method="post" novalidate>
    <!-- campos -->
</form>
```

---

## 🏷️ Labels e Inputs: Fundamentos dos Formulários

### Labels e Inputs: Trabalhando Juntos

**Labels** (rótulos) e **inputs** (campos de entrada) são os blocos fundamentais para criar formulários em HTML. Labels fornecem texto descritivo que informa aos usuários qual informação é esperada em um campo correspondente. Inputs são os controles interativos onde os usuários podem inserir dados.

### O Elemento `<label>`

O elemento `<label>` associa texto descritivo a um campo de entrada, melhorando acessibilidade e usabilidade.

**Duas formas de usar:**

#### 1. Label Envolvendo o Input

```html
<label>
    Nome completo:
    <input type="text" name="nome">
</label>
```

#### 2. Label com `for` e `id`

```html
<label for="nome">Nome completo:</label>
<input type="text" id="nome" name="nome">
```

**Por que usar labels?**
- **Acessibilidade**: Leitores de tela podem associar o texto ao campo
- **Usabilidade**: Clicar no label foca o campo correspondente
- **SEO**: Melhora a semântica do formulário

### O Elemento `<input>`

O elemento `<input>` é usado para criar vários tipos de campos de entrada, dependendo do valor do atributo `type`.

**Sintaxe básica:**
```html
<input type="text" name="campo" id="campo">
```

### Tipos de Input Comuns

#### `type="text"` - Texto Simples

```html
<label for="nome">Nome:</label>
<input type="text" id="nome" name="nome" placeholder="Digite seu nome">
```

#### `type="email"` - Email

```html
<label for="email">Email:</label>
<input type="email" id="email" name="email" placeholder="seu@email.com">
```

#### `type="password"` - Senha

```html
<label for="senha">Senha:</label>
<input type="password" id="senha" name="senha" placeholder="Digite sua senha">
```

#### `type="number"` - Número

```html
<label for="idade">Idade:</label>
<input type="number" id="idade" name="idade" min="0" max="120" step="1">
```

#### `type="tel"` - Telefone

```html
<label for="telefone">Telefone:</label>
<input type="tel" id="telefone" name="telefone" placeholder="(00) 00000-0000">
```

#### `type="url"` - URL

```html
<label for="website">Website:</label>
<input type="url" id="website" name="website" placeholder="https://exemplo.com">
```

#### `type="date"` - Data

```html
<label for="nascimento">Data de nascimento:</label>
<input type="date" id="nascimento" name="nascimento">
```

#### `type="time"` - Hora

```html
<label for="horario">Horário:</label>
<input type="time" id="horario" name="horario">
```

#### `type="datetime-local"` - Data e Hora Local

```html
<label for="datahora">Data e hora:</label>
<input type="datetime-local" id="datahora" name="datahora">
```

#### `type="checkbox"` - Caixa de Seleção

```html
<label>
    <input type="checkbox" name="termos" value="aceito">
    Aceito os termos e condições
</label>
```

#### `type="radio"` - Botão de Opção

```html
<fieldset>
    <legend>Gênero:</legend>
    <label>
        <input type="radio" name="genero" value="masculino">
        Masculino
    </label>
    <label>
        <input type="radio" name="genero" value="feminino">
        Feminino
    </label>
    <label>
        <input type="radio" name="genero" value="outro">
        Outro
    </label>
</fieldset>
```

#### `type="range"` - Controle Deslizante

```html
<label for="volume">Volume:</label>
<input type="range" id="volume" name="volume" min="0" max="100" value="50">
<span id="volume-valor">50</span>
```

#### `type="color"` - Seletor de Cor

```html
<label for="cor">Escolha uma cor:</label>
<input type="color" id="cor" name="cor" value="#ff0000">
```

#### `type="search"` - Campo de Busca

```html
<label for="busca">Buscar:</label>
<input type="search" id="busca" name="busca" placeholder="Digite sua busca">
```

#### `type="hidden"` - Campo Oculto

```html
<input type="hidden" name="token" value="abc123xyz">
```

### Atributos Comuns de Input

#### `name`
Identifica o campo quando o formulário é enviado (obrigatório para envio).

```html
<input type="text" name="nome">
```

#### `id`
Identificador único do elemento (usado com `<label for="">`).

```html
<label for="email">Email:</label>
<input type="email" id="email" name="email">
```

#### `value`
Valor padrão do campo.

```html
<input type="text" name="nome" value="João Silva">
```

#### `placeholder`
Texto de exemplo que aparece quando o campo está vazio.

```html
<input type="text" name="nome" placeholder="Digite seu nome">
```

#### `required`
Torna o campo obrigatório.

```html
<input type="email" name="email" required>
```

#### `readonly`
Torna o campo somente leitura (pode ser enviado).

```html
<input type="text" name="id" value="123" readonly>
```

#### `disabled`
Desabilita o campo (não é enviado no formulário).

```html
<input type="text" name="campo" disabled>
```

#### `maxlength`
Limita o número máximo de caracteres.

```html
<input type="text" name="nome" maxlength="50">
```

#### `min` e `max`
Define valores mínimo e máximo (para números, datas, etc.).

```html
<input type="number" name="idade" min="18" max="100">
```

#### `step`
Define o incremento para campos numéricos.

```html
<input type="number" name="quantidade" min="0" max="100" step="5">
```

#### `pattern`
Define um padrão regex para validação.

```html
<input type="text" name="cep" pattern="[0-9]{5}-[0-9]{3}" placeholder="12345-678">
```

#### `autocomplete`
Controla o preenchimento automático do navegador.

```html
<input type="email" name="email" autocomplete="email">
<input type="password" name="senha" autocomplete="current-password">
```

---

## 📤 Upload de Arquivos

### O que é Upload de Arquivos?

**Upload de arquivos** permite que usuários enviem arquivos de seus computadores para um servidor web. Isso é tipicamente alcançado através de um formulário HTML que inclui um elemento `<input>` com o atributo `type` definido como `"file"`.

### Input `type="file"`

**Sintaxe básica:**
```html
<input type="file" name="arquivo">
```

### Atributos do Input File

#### `accept`
Especifica quais tipos de arquivo são aceitos.

```html
<!-- Apenas imagens -->
<input type="file" name="foto" accept="image/*">

<!-- Apenas PDFs -->
<input type="file" name="documento" accept=".pdf">

<!-- Múltiplos tipos -->
<input type="file" name="arquivo" accept=".pdf,.doc,.docx,image/*">
```

#### `multiple`
Permite selecionar múltiplos arquivos.

```html
<input type="file" name="fotos" accept="image/*" multiple>
```

#### `capture`
Em dispositivos móveis, especifica qual câmera usar.

```html
<!-- Usar câmera traseira -->
<input type="file" name="foto" accept="image/*" capture="environment">

<!-- Usar câmera frontal -->
<input type="file" name="foto" accept="image/*" capture="user">
```

### Formulário Completo para Upload

```html
<form action="/upload" method="post" enctype="multipart/form-data">
    <label for="arquivo">Selecione um arquivo:</label>
    <input 
        type="file" 
        id="arquivo" 
        name="arquivo" 
        accept="image/*"
        required
    >
    
    <button type="submit">Enviar Arquivo</button>
</form>
```

**⚠️ Importante:** Sempre use `enctype="multipart/form-data"` no formulário quando houver upload de arquivos.

### Upload de Múltiplos Arquivos

```html
<form action="/upload" method="post" enctype="multipart/form-data">
    <label for="fotos">Selecione uma ou mais fotos:</label>
    <input 
        type="file" 
        id="fotos" 
        name="fotos" 
        accept="image/*"
        multiple
        required
    >
    
    <button type="submit">Enviar Fotos</button>
</form>
```

### Exemplo com Preview de Imagem

```html
<form action="/upload" method="post" enctype="multipart/form-data">
    <label for="foto">Selecione uma foto:</label>
    <input 
        type="file" 
        id="foto" 
        name="foto" 
        accept="image/*"
        onchange="previewImage(this)"
    >
    
    <div id="preview"></div>
    
    <button type="submit">Enviar</button>
</form>

<script>
function previewImage(input) {
    const preview = document.getElementById('preview');
    preview.innerHTML = '';
    
    if (input.files && input.files[0]) {
        const reader = new FileReader();
        
        reader.onload = function(e) {
            const img = document.createElement('img');
            img.src = e.target.result;
            img.style.maxWidth = '300px';
            preview.appendChild(img);
        };
        
        reader.readAsDataURL(input.files[0]);
    }
}
</script>
```

---

## ✅ Validação de Formulários

### O que é Validação de Formulários?

**Validação de formulários** é o processo de verificar se as informações fornecidas pelo usuário em um formulário estão corretas e completas antes de serem enviadas. Isso garante que os dados recebidos sejam precisos e atendam ao formato necessário, prevenindo erros e melhorando a qualidade dos dados.

### Validação HTML5 Nativa

HTML5 fornece validação nativa através de atributos nos elementos de formulário.

#### `required` - Campo Obrigatório

```html
<label for="nome">Nome:</label>
<input type="text" id="nome" name="nome" required>
```

#### `min` e `max` - Valores Mínimo e Máximo

```html
<label for="idade">Idade (18-100):</label>
<input type="number" id="idade" name="idade" min="18" max="100" required>
```

#### `minlength` e `maxlength` - Comprimento do Texto

```html
<label for="senha">Senha (mínimo 8 caracteres):</label>
<input type="password" id="senha" name="senha" minlength="8" required>
```

#### `pattern` - Padrão Regex

```html
<label for="cep">CEP:</label>
<input 
    type="text" 
    id="cep" 
    name="cep" 
    pattern="[0-9]{5}-[0-9]{3}"
    placeholder="12345-678"
    required
>
```

#### `type="email"` - Validação de Email

```html
<label for="email">Email:</label>
<input type="email" id="email" name="email" required>
```

#### `type="url"` - Validação de URL

```html
<label for="website">Website:</label>
<input type="url" id="website" name="website" required>
```

### Mensagens de Validação Customizadas

Você pode personalizar mensagens de validação usando JavaScript:

```html
<form id="meuFormulario">
    <label for="email">Email:</label>
    <input 
        type="email" 
        id="email" 
        name="email" 
        required
        oninvalid="this.setCustomValidity('Por favor, insira um email válido')"
        oninput="this.setCustomValidity('')"
    >
    
    <button type="submit">Enviar</button>
</form>
```

### Validação com JavaScript

```html
<form id="formulario">
    <label for="senha">Senha:</label>
    <input type="password" id="senha" name="senha" required>
    <span id="erro-senha" style="color: red;"></span>
    
    <label for="confirmar-senha">Confirmar Senha:</label>
    <input type="password" id="confirmar-senha" name="confirmar-senha" required>
    <span id="erro-confirmar" style="color: red;"></span>
    
    <button type="submit">Cadastrar</button>
</form>

<script>
document.getElementById('formulario').addEventListener('submit', function(e) {
    const senha = document.getElementById('senha').value;
    const confirmar = document.getElementById('confirmar-senha').value;
    
    if (senha !== confirmar) {
        e.preventDefault();
        document.getElementById('erro-confirmar').textContent = 
            'As senhas não coincidem';
        return false;
    }
    
    if (senha.length < 8) {
        e.preventDefault();
        document.getElementById('erro-senha').textContent = 
            'A senha deve ter pelo menos 8 caracteres';
        return false;
    }
});
</script>
```

### Estados de Validação CSS

Você pode estilizar campos válidos e inválidos usando pseudo-classes CSS:

```html
<style>
input:valid {
    border: 2px solid green;
}

input:invalid {
    border: 2px solid red;
}

input:focus:invalid {
    outline: none;
    border-color: red;
    box-shadow: 0 0 5px red;
}
</style>
```

---

## 🎯 Restrições de Formulários HTML

### O que são Restrições de Formulários?

**Restrições de formulários HTML** são regras que você define nos campos do formulário para controlar que tipo de dados os usuários podem inserir. Essas restrições ajudam a garantir que as informações enviadas sejam válidas e atendam aos seus requisitos.

### Tipos de Restrições

#### 1. Campos Obrigatórios (`required`)

```html
<input type="text" name="nome" required>
```

#### 2. Comprimento Mínimo e Máximo

```html
<!-- Texto -->
<input type="text" name="usuario" minlength="3" maxlength="20">

<!-- Números -->
<input type="number" name="quantidade" min="1" max="100">
```

#### 3. Padrões de Texto (`pattern`)

```html
<!-- CEP brasileiro -->
<input 
    type="text" 
    name="cep" 
    pattern="[0-9]{5}-[0-9]{3}"
    placeholder="12345-678"
>

<!-- Telefone -->
<input 
    type="tel" 
    name="telefone" 
    pattern="[0-9]{2} [0-9]{5}-[0-9]{4}"
    placeholder="11 99999-9999"
>
```

#### 4. Tipos de Dados Específicos

```html
<!-- Email -->
<input type="email" name="email">

<!-- URL -->
<input type="url" name="website">

<!-- Data -->
<input type="date" name="nascimento" min="1900-01-01" max="2024-12-31">

<!-- Hora -->
<input type="time" name="horario" min="09:00" max="18:00">
```

#### 5. Incrementos (`step`)

```html
<!-- Números inteiros -->
<input type="number" name="quantidade" step="1" min="0">

<!-- Decimais (0.5) -->
<input type="number" name="peso" step="0.5" min="0">

<!-- Decimais (0.01) -->
<input type="number" name="preco" step="0.01" min="0">
```

### Exemplo Completo de Formulário com Restrições

```html
<form action="/cadastrar" method="post">
    <fieldset>
        <legend>Dados Pessoais</legend>
        
        <label for="nome">Nome completo:</label>
        <input 
            type="text" 
            id="nome" 
            name="nome" 
            required
            minlength="3"
            maxlength="100"
            pattern="[A-Za-zÀ-ÿ\s]+"
            title="Apenas letras e espaços"
        >
        
        <label for="email">Email:</label>
        <input 
            type="email" 
            id="email" 
            name="email" 
            required
            autocomplete="email"
        >
        
        <label for="idade">Idade:</label>
        <input 
            type="number" 
            id="idade" 
            name="idade" 
            required
            min="18"
            max="100"
            step="1"
        >
        
        <label for="senha">Senha:</label>
        <input 
            type="password" 
            id="senha" 
            name="senha" 
            required
            minlength="8"
            pattern="(?=.*[a-z])(?=.*[A-Z])(?=.*\d).{8,}"
            title="Mínimo 8 caracteres, incluindo maiúscula, minúscula e número"
        >
    </fieldset>
    
    <button type="submit">Cadastrar</button>
</form>
```

### Elementos Adicionais de Formulário

#### `<textarea>` - Área de Texto

```html
<label for="mensagem">Mensagem:</label>
<textarea 
    id="mensagem" 
    name="mensagem" 
    rows="5" 
    cols="50"
    required
    minlength="10"
    maxlength="500"
    placeholder="Digite sua mensagem aqui..."
></textarea>
```

#### `<select>` - Lista Suspensa

```html
<label for="pais">País:</label>
<select id="pais" name="pais" required>
    <option value="">Selecione um país</option>
    <option value="br">Brasil</option>
    <option value="us">Estados Unidos</option>
    <option value="pt">Portugal</option>
</select>
```

#### `<select>` Múltipla Escolha

```html
<label for="interesses">Interesses (selecione múltiplos):</label>
<select id="interesses" name="interesses" multiple size="5">
    <option value="tecnologia">Tecnologia</option>
    <option value="esportes">Esportes</option>
    <option value="musica">Música</option>
    <option value="viagens">Viagens</option>
</select>
```

#### `<fieldset>` e `<legend>` - Agrupamento

```html
<form>
    <fieldset>
        <legend>Informações de Contato</legend>
        <label for="email">Email:</label>
        <input type="email" id="email" name="email">
        
        <label for="telefone">Telefone:</label>
        <input type="tel" id="telefone" name="telefone">
    </fieldset>
    
    <fieldset>
        <legend>Preferências</legend>
        <label>
            <input type="checkbox" name="newsletter" value="sim">
            Receber newsletter
        </label>
    </fieldset>
</form>
```

#### `<button>` - Botões

```html
<!-- Botão de submit -->
<button type="submit">Enviar</button>

<!-- Botão de reset -->
<button type="reset">Limpar</button>

<!-- Botão genérico (para JavaScript) -->
<button type="button" onclick="minhaFuncao()">Clique Aqui</button>
```

---

## 📋 Resumo dos Conceitos

### Embedding Media
- **Imagens**: Use `<img>` para imagens simples, `<figure>` para imagens com legenda
- **Áudio**: Use `<audio>` com múltiplas fontes para compatibilidade
- **Vídeo**: Use `<video>` com controles e legendas para acessibilidade
- **iframe**: Use com cuidado e sempre considere segurança

### Imagens
- Sempre use `alt` descritivo
- Especifique `width` e `height` para evitar layout shift
- Use `fetchpriority` para otimizar carregamento
- Use `loading="lazy"` para imagens abaixo da dobra

### Formulários
- Use `<label>` sempre para acessibilidade
- Escolha o `type` correto para cada campo
- Use validação HTML5 nativa quando possível
- Sempre valide no servidor também (segurança)

### Segurança
- Use CSP para prevenir XSS
- Valide dados no servidor (nunca confie apenas no cliente)
- Use HTTPS para formulários
- Sanitize dados de entrada

---

## 🎯 Próximos Passos

Agora que você aprendeu sobre embedding media e formulários, você está pronto para:
- Criar páginas web ricas em conteúdo multimídia
- Coletar dados dos usuários de forma segura
- Validar informações antes do envio
- Melhorar a experiência do usuário com mídia otimizada

Na próxima aula, continuaremos explorando recursos avançados do HTML!


