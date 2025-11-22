# Aula 2 - Performance, Boas Práticas e Otimização

## 🚀 Performance Web: Por Que É Importante?

### Impacto da Performance

A velocidade de um site impacta diretamente:
- **Experiência do Usuário**: Sites lentos frustram usuários
- **Conversão**: Cada segundo de atraso pode reduzir conversões em 7%
- **SEO**: Google prioriza sites rápidos nos resultados de busca
- **Custos**: Sites lentos consomem mais recursos do servidor
- **Acessibilidade**: Usuários com conexões lentas são mais afetados

### Métricas de Performance

#### Core Web Vitals (Google)

**LCP (Largest Contentful Paint)**
- **Meta**: < 2.5 segundos
- **Mede**: Tempo para carregar o conteúdo principal
- **Impacto**: Usuário vê quando o conteúdo principal aparece

**FID (First Input Delay)**
- **Meta**: < 100 milissegundos
- **Mede**: Tempo de resposta à primeira interação
- **Impacto**: Quão responsivo o site é às ações do usuário

**CLS (Cumulative Layout Shift)**
- **Meta**: < 0.1
- **Mede**: Estabilidade visual durante o carregamento
- **Impacto**: Evita que elementos "pulem" na tela

---

## ⚡ Otimização de DNS

### O Problema do DNS

A resolução DNS pode adicionar **50-200ms** ao tempo de carregamento. Em conexões lentas ou com DNS lento, isso pode ser ainda pior.

### Boas Práticas de DNS

#### 1. DNS Prefetch

Informe ao navegador para resolver DNS antecipadamente:

```html
<head>
    <!-- Para recursos externos que você sabe que vai usar -->
    <link rel="dns-prefetch" href="https://fonts.googleapis.com">
    <link rel="dns-prefetch" href="https://cdn.exemplo.com">
</head>
```

**Quando usar:**
- Recursos de CDN
- APIs externas
- Fontes web (Google Fonts, etc.)
- Analytics e tracking

#### 2. Minimizar Domínios Externos

Cada domínio diferente requer uma nova resolução DNS:

```html
❌ Ruim: Muitos domínios diferentes
<link href="https://cdn1.com/style.css">
<script src="https://cdn2.com/script.js">
<img src="https://cdn3.com/image.jpg">
```

```html
✅ Bom: Consolidar em menos domínios
<link href="https://cdn.exemplo.com/style.css">
<script src="https://cdn.exemplo.com/script.js">
<img src="https://cdn.exemplo.com/image.jpg">
```

#### 3. Usar CDN (Content Delivery Network)

CDNs distribuem conteúdo globalmente, reduzindo latência:
- Servidores próximos aos usuários
- Cache distribuído
- Redução de latência DNS e de rede

#### 4. TTL (Time To Live) Apropriado

Configure TTL adequado nos registros DNS:
- **Alto TTL** (ex: 86400 = 24h): Para IPs estáveis, reduz requisições DNS
- **Baixo TTL** (ex: 300 = 5min): Para mudanças frequentes, mas aumenta requisições

---

## 🌐 Otimização de HTTP

### HTTP/2 e HTTP/3

#### HTTP/2

**Vantagens:**
- **Multiplexing**: Múltiplas requisições em uma conexão
- **Header Compression**: Compressão de headers
- **Server Push**: Servidor envia recursos antes de serem solicitados
- **Priorização**: Prioriza recursos importantes

**Como Habilitar:**
- Configure seu servidor para suportar HTTP/2
- Use HTTPS (HTTP/2 requer SSL/TLS)
- Verifique com DevTools → Network → Protocol

#### HTTP/3

**Vantagens:**
- Baseado em QUIC (UDP)
- Conexões mais rápidas
- Melhor em conexões instáveis
- Redução de latência

**Status:**
- Ainda em adoção gradual
- Suportado por navegadores modernos
- Requer suporte do servidor

### Compressão

#### Gzip/Brotli

Comprima recursos textuais (HTML, CSS, JS):

```html
<!-- Servidor deve configurar compressão -->
Content-Encoding: gzip
Content-Encoding: br  <!-- Brotli é melhor que Gzip -->
```

**Benefícios:**
- Redução de 70-90% no tamanho de arquivos de texto
- Menor tempo de download
- Menor uso de banda

**Recursos que devem ser comprimidos:**
- HTML
- CSS
- JavaScript
- JSON
- XML
- SVG

**Recursos que NÃO devem ser comprimidos:**
- Imagens (já comprimidas)
- Vídeos (já comprimidos)
- Arquivos binários

### Cache HTTP

#### Headers de Cache

```html
<!-- Cache por 1 ano (para recursos estáticos) -->
Cache-Control: max-age=31536000, immutable

<!-- Revalidação (para HTML dinâmico) -->
Cache-Control: max-age=3600, must-revalidate

<!-- Sem cache (para conteúdo sensível) -->
Cache-Control: no-cache, no-store, must-revalidate
```

#### Estratégias de Cache

**Recursos Estáticos (CSS, JS, Imagens):**
```html
Cache-Control: public, max-age=31536000, immutable
```
- Cache longo (1 ano)
- Nomeie arquivos com hash/versão para invalidação

**HTML Dinâmico:**
```html
Cache-Control: no-cache, must-revalidate
```
- Sempre revalidar
- Permite atualizações rápidas

**APIs:**
```html
Cache-Control: private, max-age=300
```
- Cache curto (5 minutos)
- Dados podem mudar frequentemente

### Redução de Requisições HTTP

#### 1. Combinar Arquivos

```html
❌ Ruim: Múltiplas requisições
<link rel="stylesheet" href="style1.css">
<link rel="stylesheet" href="style2.css">
<link rel="stylesheet" href="style3.css">
```

```html
✅ Bom: Um arquivo combinado (em produção)
<link rel="stylesheet" href="styles.min.css">
```

#### 2. Inline CSS/JS Crítico

Para CSS/JS crítico (above-the-fold):

```html
<head>
    <style>
        /* CSS crítico inline */
        body { margin: 0; }
        .header { ... }
    </style>
</head>
```

**Benefício:** Reduz render-blocking resources

#### 3. Lazy Loading

Carregue recursos apenas quando necessário:

```html
<!-- Imagens -->
<img src="imagem.jpg" loading="lazy" alt="Descrição">

<!-- Iframes -->
<iframe src="video.html" loading="lazy"></iframe>
```

---

## 🏠 Otimização de Hospedagem

### Escolha do Provedor

#### Fatores de Performance

**Localização dos Servidores:**
- Escolha servidores próximos ao seu público-alvo
- Use CDN para distribuição global

**Uptime (Tempo Online):**
- Procure 99.9%+ de uptime
- SLA (Service Level Agreement) claro

**Performance do Servidor:**
- SSD ao invés de HDD
- RAM adequada
- CPU suficiente

**Suporte a HTTP/2 e HTTP/3:**
- Verifique se o provedor suporta versões modernas

### Configuração do Servidor

#### 1. Compressão

Configure Gzip/Brotli no servidor:
- **Apache**: mod_deflate
- **Nginx**: gzip module
- **Cloudflare**: Automático

#### 2. Cache Headers

Configure headers de cache apropriados:
- Recursos estáticos: Cache longo
- HTML: Cache curto ou sem cache

#### 3. Minificação

Minifique HTML, CSS e JS:
- Remove espaços, comentários, quebras de linha
- Reduz tamanho de arquivo
- Use ferramentas de build (Webpack, Vite, etc.)

---

## 🌐 Otimização de Navegadores

### Renderização Otimizada

#### 1. Evitar Render-Blocking Resources

```html
❌ Ruim: CSS bloqueia renderização
<head>
    <link rel="stylesheet" href="styles.css">
</head>
```

```html
✅ Bom: CSS crítico inline, resto assíncrono
<head>
    <style>/* CSS crítico */</style>
    <link rel="preload" href="styles.css" as="style" onload="this.onload=null;this.rel='stylesheet'">
</head>
```

#### 2. Defer/Async em Scripts

```html
<!-- Defer: Executa após HTML parseado -->
<script src="script.js" defer></script>

<!-- Async: Executa assim que disponível (não bloqueia) -->
<script src="analytics.js" async></script>
```

**Quando usar:**
- **Defer**: Scripts que dependem do DOM
- **Async**: Scripts independentes (analytics, ads)

#### 3. Preload e Prefetch

```html
<!-- Preload: Recurso crítico -->
<link rel="preload" href="font.woff2" as="font" type="font/woff2" crossorigin>

<!-- Prefetch: Recurso futuro -->
<link rel="prefetch" href="next-page.html">
```

### Otimização de Imagens

#### 1. Formatos Modernos

```html
<!-- Use formatos modernos quando possível -->
<picture>
    <source srcset="imagem.avif" type="image/avif">
    <source srcset="imagem.webp" type="image/webp">
    <img src="imagem.jpg" alt="Descrição">
</picture>
```

**Formatos:**
- **AVIF**: Melhor compressão, suporte crescente
- **WebP**: Boa compressão, amplo suporte
- **JPEG/PNG**: Fallback para navegadores antigos

#### 2. Tamanhos Responsivos

```html
<img srcset="
    imagem-400w.jpg 400w,
    imagem-800w.jpg 800w,
    imagem-1200w.jpg 1200w
" sizes="(max-width: 600px) 400px, (max-width: 1200px) 800px, 1200px"
     src="imagem-800w.jpg" alt="Descrição">
```

#### 3. Lazy Loading

```html
<img src="imagem.jpg" loading="lazy" alt="Descrição">
```

**Benefício:** Carrega imagens apenas quando visíveis

#### 4. Otimização de Tamanho

- Comprima imagens antes de fazer upload
- Use ferramentas: TinyPNG, ImageOptim, Squoosh
- Remova metadados desnecessários

---

## 🔍 Otimização de SEO e Performance

### SEO Técnico e Performance

#### 1. Core Web Vitals

Otimize as métricas Core Web Vitals:
- **LCP**: Otimize imagens, use CDN, minimize render-blocking
- **FID**: Minimize JavaScript, use code splitting
- **CLS**: Defina dimensões de imagens, evite conteúdo dinâmico acima do fold

#### 2. Mobile-First

```html
<meta name="viewport" content="width=device-width, initial-scale=1.0">
```

- Design mobile-first
- Teste em dispositivos reais
- Google prioriza sites mobile-friendly

#### 3. Estrutura Semântica

```html
✅ Bom: Estrutura semântica clara
<header>
    <h1>Título Principal</h1>
</header>
<main>
    <article>
        <h2>Subtítulo</h2>
        <p>Conteúdo...</p>
    </article>
</main>
```

**Benefícios:**
- Melhor indexação
- Melhor acessibilidade
- Melhor performance de renderização

#### 4. URLs Amigáveis

```
❌ Ruim: exemplo.com/pagina.php?id=123&cat=abc
✅ Bom: exemplo.com/produtos/categoria/nome-produto
```

**Benefícios:**
- Melhor SEO
- Mais fácil de compartilhar
- Mais fácil de lembrar

### Schema Markup

Adicione dados estruturados:

```html
<script type="application/ld+json">
{
  "@context": "https://schema.org",
  "@type": "Article",
  "headline": "Título do Artigo",
  "author": {
    "@type": "Person",
    "name": "Nome do Autor"
  }
}
</script>
```

**Benefícios:**
- Rich snippets nos resultados de busca
- Melhor compreensão do conteúdo
- Potencial para melhor ranking

---

## 🛡️ Segurança e Performance

### HTTPS

**Por que HTTPS é importante:**
- **Segurança**: Criptografia de dados
- **SEO**: Google prioriza sites HTTPS
- **Performance**: HTTP/2 requer HTTPS
- **Confiança**: Usuários confiam mais

**Como obter:**
- **Let's Encrypt**: Certificado SSL gratuito
- **Cloudflare**: SSL automático
- **Provedor de Hospedagem**: Muitos oferecem SSL gratuito

### Headers de Segurança

```html
<!-- Content Security Policy -->
Content-Security-Policy: default-src 'self'

<!-- X-Frame-Options -->
X-Frame-Options: DENY

<!-- X-Content-Type-Options -->
X-Content-Type-Options: nosniff
```

**Benefício:** Proteção contra ataques, melhor segurança

---

## 📊 Ferramentas de Análise

### Ferramentas de Performance

#### 1. Google PageSpeed Insights
- [pagespeed.web.dev](https://pagespeed.web.dev/)
- Analisa Core Web Vitals
- Sugestões de otimização

#### 2. GTmetrix
- [gtmetrix.com](https://gtmetrix.com/)
- Análise detalhada de performance
- Waterfall de requisições

#### 3. WebPageTest
- [webpagetest.org](https://www.webpagetest.org/)
- Teste de múltiplas localizações
- Análise profunda

#### 4. Chrome DevTools
- Network tab: Análise de requisições
- Performance tab: Profiling de performance
- Lighthouse: Auditoria automatizada

### Métricas para Monitorar

- **Time to First Byte (TTFB)**: < 200ms
- **First Contentful Paint (FCP)**: < 1.8s
- **LCP**: < 2.5s
- **FID**: < 100ms
- **CLS**: < 0.1
- **Total Blocking Time (TBT)**: < 200ms

---

## ✅ Checklist de Otimização

### DNS
- [ ] DNS Prefetch para recursos externos
- [ ] Minimizar número de domínios
- [ ] Usar CDN quando apropriado
- [ ] TTL configurado adequadamente

### HTTP
- [ ] HTTP/2 ou HTTP/3 habilitado
- [ ] Compressão Gzip/Brotli configurada
- [ ] Headers de cache apropriados
- [ ] Redução de requisições HTTP

### Hospedagem
- [ ] Servidor próximo ao público-alvo
- [ ] Uptime alto (99.9%+)
- [ ] SSD e recursos adequados
- [ ] Suporte a HTTP/2/3

### Navegador
- [ ] CSS crítico inline
- [ ] Scripts com defer/async
- [ ] Lazy loading de imagens
- [ ] Preload de recursos críticos

### Imagens
- [ ] Formatos modernos (WebP, AVIF)
- [ ] Tamanhos responsivos (srcset)
- [ ] Compressão adequada
- [ ] Lazy loading

### SEO
- [ ] Core Web Vitals otimizados
- [ ] Mobile-first design
- [ ] Estrutura semântica
- [ ] URLs amigáveis
- [ ] Schema markup

### Segurança
- [ ] HTTPS habilitado
- [ ] Headers de segurança configurados
- [ ] Certificado SSL válido

---

## 🎯 Priorização de Otimizações

### Quick Wins (Ganhos Rápidos)

1. **Habilitar Compressão**: 5 minutos, grande impacto
2. **Configurar Cache**: 10 minutos, impacto significativo
3. **Otimizar Imagens**: 30 minutos, impacto visual
4. **Lazy Loading**: 5 minutos, melhora percepção

### Médio Prazo

1. **Implementar HTTP/2**: Requer configuração do servidor
2. **CDN**: Requer configuração e possível migração
3. **Minificação**: Requer processo de build
4. **Code Splitting**: Requer refatoração

### Longo Prazo

1. **Arquitetura de Performance**: Requer planejamento
2. **Monitoramento Contínuo**: Requer ferramentas
3. **Otimização Contínua**: Processo iterativo

---

## 📝 Resumo

### Princípios Fundamentais

1. **Medir Primeiro**: Use ferramentas para identificar problemas
2. **Otimizar o Crítico**: Foque no que mais impacta
3. **Testar Continuamente**: Performance muda com o tempo
4. **Monitorar**: Configure alertas e monitoramento

### Impacto Esperado

Com essas otimizações, você pode esperar:
- **Redução de 30-50%** no tempo de carregamento
- **Melhoria significativa** nos Core Web Vitals
- **Melhor ranking** nos mecanismos de busca
- **Maior satisfação** dos usuários

### Próximos Passos

1. Execute uma auditoria de performance no seu site
2. Identifique os maiores problemas
3. Implemente otimizações priorizadas
4. Meça o impacto
5. Itere e melhore continuamente

---

## 🔗 Recursos Adicionais

- **Web.dev Performance**: [web.dev/performance](https://web.dev/performance)
- **Google PageSpeed Insights**: [pagespeed.web.dev](https://pagespeed.web.dev/)
- **MDN - Web Performance**: [developer.mozilla.org/en-US/docs/Web/Performance](https://developer.mozilla.org/en-US/docs/Web/Performance)
- **HTTP Archive**: [httparchive.org](https://httparchive.org/)

**Lembre-se: Performance não é um destino, é uma jornada contínua!** 🚀

