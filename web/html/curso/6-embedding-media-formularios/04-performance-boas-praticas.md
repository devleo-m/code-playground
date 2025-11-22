# Aula 6 - Performance, Boas Práticas e Otimização

## 🚀 Performance: Otimização de Mídia

### Por que Otimizar Mídia é Crucial?

Mídia não otimizada é uma das principais causas de páginas lentas:

1. **Tamanho de Arquivo**: Imagens, vídeos e áudios podem ser muito grandes
2. **Tempo de Carregamento**: Arquivos grandes demoram para baixar
3. **Uso de Banda**: Consome dados do usuário (especialmente em mobile)
4. **Experiência do Usuário**: Páginas lentas frustram usuários
5. **SEO**: Google penaliza sites lentos

### Métricas de Performance Importantes

- **LCP (Largest Contentful Paint)**: Tempo até o maior elemento aparecer
- **CLS (Cumulative Layout Shift)**: Estabilidade visual da página
- **FCP (First Contentful Paint)**: Tempo até o primeiro conteúdo aparecer
- **TTI (Time to Interactive)**: Tempo até a página ficar interativa

---

## 🖼️ Performance de Imagens

### 1. Sempre Especifique Width e Height

**❌ Evite:**
```html
<img src="foto.jpg" alt="Foto">
```

**✅ Prefira:**
```html
<img src="foto.jpg" alt="Foto" width="800" height="600">
```

**Por quê?**
- Evita **Cumulative Layout Shift (CLS)**
- Navegador reserva espaço antes da imagem carregar
- Melhora a percepção de velocidade
- Essencial para Core Web Vitals

### 2. Use Lazy Loading para Imagens Abaixo da Dobra

**❌ Evite:**
```html
<img src="galeria-1.jpg" alt="Foto 1">
<img src="galeria-2.jpg" alt="Foto 2">
<!-- Todas carregam imediatamente -->
```

**✅ Prefira:**
```html
<!-- Hero image - carrega imediatamente -->
<img src="hero.jpg" alt="Banner" fetchpriority="high" width="1920" height="1080">

<!-- Galeria - carrega sob demanda -->
<img src="galeria-1.jpg" alt="Foto 1" loading="lazy" width="800" height="600">
<img src="galeria-2.jpg" alt="Foto 2" loading="lazy" width="800" height="600">
```

**Por quê?**
- Reduz o tempo de carregamento inicial
- Economiza banda do usuário
- Melhora LCP e FCP
- Navegador carrega apenas quando necessário

### 3. Use Priority Hints para Imagens Críticas

**✅ Boa Prática:**
```html
<!-- Imagem hero (above the fold) -->
<img 
    src="hero.jpg" 
    alt="Banner principal" 
    fetchpriority="high"
    width="1920"
    height="1080"
>

<!-- Imagens secundárias -->
<img 
    src="galeria.jpg" 
    alt="Galeria" 
    fetchpriority="low"
    loading="lazy"
    width="800"
    height="600"
>
```

**Por quê?**
- Navegador prioriza recursos importantes
- Melhora LCP (Largest Contentful Paint)
- Otimiza uso de banda

### 4. Escolha o Formato Correto

**Formatos e Quando Usar:**

- **JPEG**: Fotografias, imagens com muitas cores
- **PNG**: Imagens com transparência, gráficos simples
- **WebP**: Formato moderno, melhor compressão (use com fallback)
- **SVG**: Ícones, logos, gráficos vetoriais
- **AVIF**: Formato mais moderno, melhor compressão (suporte limitado)

**✅ Exemplo com Fallback:**
```html
<picture>
    <source srcset="imagem.avif" type="image/avif">
    <source srcset="imagem.webp" type="image/webp">
    <img src="imagem.jpg" alt="Imagem" width="800" height="600">
</picture>
```

### 5. Otimize Tamanho de Arquivo

**Ferramentas para Otimização:**
- **TinyPNG**: Compressão de PNG e JPEG
- **Squoosh**: Compressão avançada com preview
- **ImageOptim**: Ferramenta desktop
- **Sharp**: Biblioteca Node.js para otimização programática

**Boas Práticas:**
- Comprima imagens antes de fazer upload
- Use dimensões apropriadas (não use imagem 4000px para exibir 400px)
- Considere usar CDN para servir imagens

---

## 🎵 Performance de Áudio

### 1. Use Preload Apropriado

**❌ Evite:**
```html
<audio src="podcast.mp3" controls autoplay></audio>
```

**✅ Prefira:**
```html
<audio src="podcast.mp3" controls preload="metadata"></audio>
```

**Valores de preload:**
- `none`: Não pré-carrega (economiza banda)
- `metadata`: Carrega apenas informações (duração, etc.)
- `auto`: Carrega tudo (padrão, pode ser lento)

**Por quê?**
- `metadata` permite mostrar duração sem carregar o arquivo inteiro
- Usuário decide se quer ouvir
- Economiza banda

### 2. Forneça Múltiplos Formatos

**✅ Boa Prática:**
```html
<audio controls preload="metadata">
    <source src="audio.mp3" type="audio/mpeg">
    <source src="audio.ogg" type="audio/ogg">
    <source src="audio.wav" type="audio/wav">
    <p>Seu navegador não suporta áudio HTML5.</p>
</audio>
```

**Por quê?**
- Compatibilidade com diferentes navegadores
- Navegador escolhe o formato que suporta
- Melhor experiência para todos

### 3. Compressão de Áudio

**Formatos e Qualidade:**
- **MP3**: 128kbps para fala, 192kbps+ para música
- **OGG**: Similar ao MP3, código aberto
- **Opus**: Melhor compressão, suporte limitado

**Ferramentas:**
- **Audacity**: Editor de áudio gratuito
- **FFmpeg**: Ferramenta de linha de comando
- **Online converters**: Vários disponíveis

---

## 🎬 Performance de Vídeo

### 1. Sempre Especifique Dimensões

**❌ Evite:**
```html
<video src="video.mp4" controls></video>
```

**✅ Prefira:**
```html
<video src="video.mp4" controls width="1280" height="720" poster="thumbnail.jpg"></video>
```

**Por quê?**
- Evita layout shift
- Navegador reserva espaço
- Melhora CLS

### 2. Use Poster Image

**✅ Boa Prática:**
```html
<video 
    src="tutorial.mp4" 
    controls 
    width="1280" 
    height="720"
    poster="thumbnail.jpg"
    preload="metadata"
></video>
```

**Por quê?**
- Dá contexto visual antes do play
- Melhora experiência do usuário
- Pode melhorar engajamento

### 3. Preload Apropriado

**✅ Boa Prática:**
```html
<!-- Vídeo principal (pode pré-carregar) -->
<video controls preload="metadata" poster="thumb.jpg">
    <source src="video.mp4" type="video/mp4">
</video>

<!-- Vídeos secundários (não pré-carregar) -->
<video controls preload="none" poster="thumb.jpg">
    <source src="video2.mp4" type="video/mp4">
</video>
```

### 4. Compressão e Formatos

**Formatos Recomendados:**
- **MP4 (H.264)**: Suporte universal
- **WebM (VP9)**: Melhor compressão, suporte limitado
- **MP4 (H.265/HEVC)**: Melhor compressão, suporte limitado

**Ferramentas:**
- **HandBrake**: Compressão de vídeo
- **FFmpeg**: Linha de comando
- **CloudConvert**: Conversão online

**Boas Práticas:**
- Comprima vídeos antes de fazer upload
- Considere múltiplas resoluções (480p, 720p, 1080p)
- Use streaming para vídeos longos

---

## 🖼️ Performance de iframe

### 1. Use Lazy Loading

**❌ Evite:**
```html
<iframe src="https://www.youtube.com/embed/VIDEO_ID"></iframe>
```

**✅ Prefira:**
```html
<iframe 
    src="https://www.youtube.com/embed/VIDEO_ID"
    loading="lazy"
    width="560"
    height="315"
    title="Vídeo do YouTube"
></iframe>
```

**Por quê?**
- iframes podem ser pesados
- Lazy loading melhora tempo inicial
- Carrega apenas quando visível

### 2. Limite Quantidade de iframes

**❌ Evite:**
```html
<!-- Múltiplos iframes na mesma página -->
<iframe src="widget1.html"></iframe>
<iframe src="widget2.html"></iframe>
<iframe src="widget3.html"></iframe>
```

**✅ Prefira:**
```html
<!-- Carregue apenas o necessário -->
<iframe src="widget-principal.html" loading="lazy"></iframe>
<!-- Carregue outros sob demanda -->
```

**Por quê?**
- Cada iframe é uma página separada
- Múltiplos iframes = múltiplas requisições
- Impacta performance significativamente

### 3. Use Sandbox Apropriado

**✅ Boa Prática:**
```html
<iframe 
    src="conteudo-externo.html"
    sandbox="allow-scripts allow-same-origin"
    title="Conteúdo externo"
></iframe>
```

**Por quê?**
- Melhora segurança
- Reduz riscos de XSS
- Limita capacidades do iframe

---

## 📝 Performance de Formulários

### 1. Validação no Cliente (Mas Não Confie Nela)

**✅ Boa Prática:**
```html
<form action="/processar" method="post">
    <label for="email">Email:</label>
    <input 
        type="email" 
        id="email" 
        name="email" 
        required
        pattern="[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,}$"
    >
    <button type="submit">Enviar</button>
</form>
```

**Por quê?**
- Feedback imediato ao usuário
- Reduz requisições desnecessárias
- Melhora experiência do usuário
- **Mas sempre valide no servidor também!**

### 2. Evite Validação Excessiva no Cliente

**❌ Evite:**
```html
<!-- Validação muito complexa apenas no cliente -->
<input 
    type="text" 
    name="senha"
    pattern="(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[@$!%*?&])[A-Za-z\d@$!%*?&]{8,}"
    onblur="validarSenhaComplexa(this)"
    onkeyup="verificarForcaSenha(this)"
    onchange="checarHistoricoSenhas(this)"
>
```

**✅ Prefira:**
```html
<!-- Validação básica no cliente, completa no servidor -->
<input 
    type="password" 
    name="senha"
    required
    minlength="8"
    pattern="(?=.*[a-z])(?=.*[A-Z])(?=.*\d).{8,}"
>
```

**Por quê?**
- Validação complexa no cliente pode ser contornada
- Mantenha validação básica no cliente
- Validação completa no servidor

### 3. Otimize Upload de Arquivos

**✅ Boa Prática:**
```html
<form action="/upload" method="post" enctype="multipart/form-data">
    <label for="arquivo">Selecione arquivo (máx. 5MB):</label>
    <input 
        type="file" 
        id="arquivo" 
        name="arquivo" 
        accept="image/*"
        required
    >
    <button type="submit">Enviar</button>
</form>

<script>
// Validação de tamanho no cliente (antes do upload)
document.querySelector('form').addEventListener('submit', function(e) {
    const file = document.getElementById('arquivo').files[0];
    if (file && file.size > 5 * 1024 * 1024) { // 5MB
        e.preventDefault();
        alert('Arquivo muito grande! Máximo 5MB.');
        return false;
    }
});
</script>
```

**Por quê?**
- Validação de tamanho antes do upload economiza tempo
- Melhor experiência do usuário
- Reduz carga no servidor

---

## 🛠️ Boas Práticas de Código

### 1. Sempre Use Alt em Imagens

**❌ Evite:**
```html
<img src="foto.jpg">
<img src="decoracao.jpg" alt="decoracao">
```

**✅ Prefira:**
```html
<img src="foto.jpg" alt="Pessoa sorrindo em um parque">
<img src="decoracao.jpg" alt=""> <!-- Vazio se for puramente decorativa -->
```

**Regras para Alt:**
- Seja descritivo mas conciso
- Descreva conteúdo e função
- Se decorativa, use `alt=""`
- Não comece com "Imagem de..." ou "Foto de..."

### 2. Use Labels Corretamente

**❌ Evite:**
```html
Nome: <input type="text" name="nome">
<input type="checkbox" name="termos"> Aceito os termos
```

**✅ Prefira:**
```html
<label for="nome">Nome:</label>
<input type="text" id="nome" name="nome">

<label>
    <input type="checkbox" name="termos">
    Aceito os termos
</label>
```

**Por quê?**
- Acessibilidade (leitores de tela)
- Usabilidade (clicar no label foca o campo)
- SEO e semântica

### 3. Agrupe Campos Relacionados

**❌ Evite:**
```html
<label>Nome:</label>
<input type="text" name="nome">
<label>Email:</label>
<input type="email" name="email">
<label>Telefone:</label>
<input type="tel" name="telefone">
```

**✅ Prefira:**
```html
<fieldset>
    <legend>Informações de Contato</legend>
    <label for="nome">Nome:</label>
    <input type="text" id="nome" name="nome">
    
    <label for="email">Email:</label>
    <input type="email" id="email" name="email">
    
    <label for="telefone">Telefone:</label>
    <input type="tel" id="telefone" name="telefone">
</fieldset>
```

**Por quê?**
- Melhor organização visual
- Melhor acessibilidade
- Agrupamento semântico

### 4. Use Tipos de Input Apropriados

**❌ Evite:**
```html
<input type="text" name="email">
<input type="text" name="telefone">
<input type="text" name="data">
```

**✅ Prefira:**
```html
<input type="email" name="email">
<input type="tel" name="telefone">
<input type="date" name="data">
```

**Por quê?**
- Validação nativa do navegador
- Teclado apropriado em mobile
- Melhor experiência do usuário
- Acessibilidade

---

## 🔒 Segurança

### 1. Content Security Policy (CSP)

**✅ Boa Prática:**
```html
<head>
    <meta 
        http-equiv="Content-Security-Policy" 
        content="
            default-src 'self';
            img-src 'self' https:;
            script-src 'self';
            style-src 'self' 'unsafe-inline';
            frame-src 'self' https://www.youtube.com;
        "
    >
</head>
```

**Por quê?**
- Previne XSS (Cross-Site Scripting)
- Controla de onde recursos podem ser carregados
- Melhora segurança geral

### 2. Validação no Servidor (Sempre!)

**⚠️ Crítico:**
```html
<!-- Validação no cliente é apenas UX -->
<form action="/processar" method="post">
    <input type="email" name="email" required>
    <button type="submit">Enviar</button>
</form>
```

**No servidor (exemplo conceitual):**
```javascript
// SEMPRE valide no servidor também!
if (!isValidEmail(req.body.email)) {
    return res.status(400).json({ error: 'Email inválido' });
}
```

**Por quê?**
- Cliente pode ser manipulado
- Ataques podem contornar validação do cliente
- Segurança real vem do servidor

### 3. Sanitização de Dados

**⚠️ Importante:**
- Sempre sanitize dados de entrada
- Escape HTML para prevenir XSS
- Valide tipos de arquivo em uploads
- Limite tamanho de arquivos

### 4. HTTPS para Formulários

**✅ Sempre:**
- Use HTTPS para formulários
- Especialmente para dados sensíveis
- Previne interceptação de dados

---

## ♿ Acessibilidade

### 1. Imagens Acessíveis

**✅ Boa Prática:**
```html
<!-- Imagem informativa -->
<img src="grafico.png" alt="Gráfico mostrando crescimento de 25% em vendas">

<!-- Imagem decorativa -->
<img src="decoracao.jpg" alt="">

<!-- Imagem com legenda -->
<figure>
    <img src="diagrama.png" alt="Diagrama do processo">
    <figcaption>Figura 1: Fluxo do processo de produção</figcaption>
</figure>
```

### 2. Formulários Acessíveis

**✅ Boa Prática:**
```html
<form>
    <fieldset>
        <legend>Informações Pessoais</legend>
        
        <label for="nome">Nome completo: <span aria-label="obrigatório">*</span></label>
        <input 
            type="text" 
            id="nome" 
            name="nome" 
            required
            aria-required="true"
            aria-describedby="nome-help"
        >
        <small id="nome-help">Digite seu nome completo</small>
        
        <label for="email">Email: <span aria-label="obrigatório">*</span></label>
        <input 
            type="email" 
            id="email" 
            name="email" 
            required
            aria-required="true"
        >
    </fieldset>
    
    <button type="submit">Enviar formulário</button>
</form>
```

### 3. Mídia Acessível

**✅ Boa Prática:**
```html
<!-- Vídeo com legendas -->
<video controls>
    <source src="video.mp4" type="video/mp4">
    <track 
        kind="subtitles" 
        src="legendas-pt.vtt" 
        srclang="pt" 
        label="Português"
        default
    >
</video>

<!-- Áudio com transcrição -->
<audio controls>
    <source src="podcast.mp3" type="audio/mpeg">
</audio>
<p>
    <a href="transcricao.txt">Ler transcrição do áudio</a>
</p>
```

---

## ❌ O Que NÃO Deve Ser Utilizado

### 1. Atributos Obsoletos

**❌ Não use:**
```html
<img src="foto.jpg" border="1">
<img src="foto.jpg" align="left">
<iframe src="page.html" frameborder="0"></iframe>
```

**✅ Use CSS:**
```html
<img src="foto.jpg" style="border: 1px solid #000;">
<img src="foto.jpg" style="float: left;">
<iframe src="page.html" style="border: none;"></iframe>
```

### 2. Autoplay de Mídia com Áudio

**❌ Evite:**
```html
<video src="video.mp4" controls autoplay></video>
<audio src="audio.mp3" controls autoplay></audio>
```

**✅ Prefira:**
```html
<video src="video.mp4" controls></video>
<audio src="audio.mp3" controls></audio>
<!-- Ou use muted com autoplay se necessário -->
<video src="video.mp4" controls autoplay muted></video>
```

**Por quê?**
- Muitos navegadores bloqueiam autoplay com áudio
- Pode ser irritante para usuários
- Pode consumir dados desnecessariamente

### 3. iframe sem Restrições

**❌ Evite:**
```html
<iframe src="https://site-desconhecido.com"></iframe>
```

**✅ Prefira:**
```html
<iframe 
    src="https://site-confiavel.com" 
    sandbox="allow-scripts allow-same-origin"
    title="Conteúdo externo"
></iframe>
```

### 4. Validação Apenas no Cliente

**❌ Nunca confie apenas em:**
```html
<input type="email" name="email" required>
<!-- Sem validação no servidor -->
```

**✅ Sempre valide também no servidor:**
```html
<!-- Cliente: UX -->
<input type="email" name="email" required>

<!-- Servidor: Segurança -->
<!-- Validação no backend -->
```

---

## 📊 Ferramentas de Análise

### 1. Google PageSpeed Insights
- Analisa performance
- Sugere melhorias
- Mede Core Web Vitals

### 2. Lighthouse (Chrome DevTools)
- Performance
- Acessibilidade
- SEO
- Boas práticas

### 3. WebPageTest
- Análise detalhada de performance
- Testa em diferentes conexões
- Waterfall de recursos

### 4. W3C Validator
- Valida HTML
- Encontra erros
- Sugere correções

---

## 🎯 Resumo das Boas Práticas

### Imagens
- ✅ Sempre especifique `width` e `height`
- ✅ Use `alt` descritivo
- ✅ Use `loading="lazy"` para imagens abaixo da dobra
- ✅ Use `fetchpriority="high"` para imagens críticas
- ✅ Otimize tamanho de arquivo
- ✅ Escolha formato apropriado

### Áudio e Vídeo
- ✅ Use `preload="metadata"` quando apropriado
- ✅ Forneça múltiplos formatos
- ✅ Especifique dimensões em vídeos
- ✅ Use poster em vídeos
- ✅ Comprima arquivos

### Formulários
- ✅ Sempre use `<label>`
- ✅ Use tipos de input apropriados
- ✅ Valide no cliente (UX) e servidor (segurança)
- ✅ Agrupe campos relacionados
- ✅ Forneça feedback claro

### Segurança
- ✅ Use CSP
- ✅ Valide no servidor sempre
- ✅ Sanitize dados
- ✅ Use HTTPS

### Acessibilidade
- ✅ Alt descritivo em imagens
- ✅ Labels conectados
- ✅ Legendas em vídeos
- ✅ Transcrições de áudio

---

**Lembre-se:** Performance e acessibilidade não são opcionais - são essenciais para uma boa experiência do usuário! 🚀

