# Aula 2: Como a Web Funciona - Conteúdo Principal

## 📖 Revisão da Aula Anterior

Na Aula 1, você aprendeu:
- **HTML** é a linguagem de marcação que estrutura páginas web
- **CSS** controla a aparência visual
- **JavaScript** adiciona interatividade
- A importância da semântica e estruturação correta

Agora, antes de mergulhar mais profundamente no HTML, é fundamental entender **como a web funciona por baixo dos panos**. Isso te dará uma base sólida para criar páginas web eficientes e entender o contexto do seu trabalho como desenvolvedor.

---

## 🌐 A Internet: A Rede Global de Computadores

### O que é a Internet?

A **Internet** é uma rede global de computadores e dispositivos interconectados que permite o compartilhamento de informações em escala mundial. É a infraestrutura física e lógica que torna possível acessar websites, enviar emails, assistir vídeos e se comunicar instantaneamente com pessoas ao redor do mundo.

### Características Principais

- **Rede Descentralizada**: Não há um único ponto de controle central
- **Protocolos Padronizados**: Usa protocolos comuns (TCP/IP, HTTP, etc.) para comunicação
- **Interconexão Global**: Milhões de dispositivos conectados simultaneamente
- **Redundância**: Múltiplos caminhos para transmitir dados, garantindo resiliência

### Como Funciona a Comunicação na Internet?

Quando você acessa um website, ocorre uma série de passos:

1. **Solicitação**: Seu navegador solicita uma página web
2. **Roteamento**: A solicitação é roteada através de múltiplos servidores
3. **Resposta**: O servidor web responde com os arquivos solicitados
4. **Renderização**: Seu navegador interpreta e exibe o conteúdo

### Componentes da Internet

- **Servidores**: Computadores que armazenam e servem conteúdo
- **Roteadores**: Dispositivos que direcionam o tráfego de dados
- **Cabo de Fibra Óptica**: Infraestrutura física que transmite dados
- **Provedores de Internet (ISPs)**: Empresas que fornecem acesso à internet
- **Data Centers**: Instalações que abrigam servidores e infraestrutura

---

## 🔌 HTTP: O Protocolo de Comunicação da Web

### O que é HTTP?

**HTTP** (Hypertext Transfer Protocol - Protocolo de Transferência de Hipertexto) é o protocolo de comunicação que permite a troca de informações entre navegadores (clientes) e servidores web. É a base de toda comunicação na World Wide Web.

### Características do HTTP

- **Stateless (Sem Estado)**: Cada requisição é independente; o servidor não "lembra" de requisições anteriores
- **Request/Response (Requisição/Resposta)**: Modelo baseado em solicitações e respostas
- **Text-Based (Baseado em Texto)**: Mensagens em formato de texto legível
- **Porta Padrão**: Usa a porta 80 (HTTP) ou 443 (HTTPS)

### Como Funciona o HTTP?

#### 1. Requisição HTTP (Request)

Quando você digita uma URL no navegador, ele cria uma **requisição HTTP**:

```
GET /index.html HTTP/1.1
Host: www.exemplo.com
User-Agent: Mozilla/5.0...
Accept: text/html
```

**Componentes de uma Requisição:**
- **Método HTTP**: GET, POST, PUT, DELETE, etc.
- **URL/Path**: O caminho do recurso solicitado
- **Headers**: Informações adicionais (cookies, tipo de conteúdo aceito, etc.)
- **Body**: Dados enviados (geralmente em requisições POST)

#### 2. Resposta HTTP (Response)

O servidor responde com uma **resposta HTTP**:

```
HTTP/1.1 200 OK
Content-Type: text/html
Content-Length: 1234

<!DOCTYPE html>
<html>...
```

**Componentes de uma Resposta:**
- **Status Code**: Código numérico indicando o resultado (200 = sucesso, 404 = não encontrado, etc.)
- **Headers**: Metadados sobre a resposta
- **Body**: O conteúdo real (HTML, JSON, imagem, etc.)

### Métodos HTTP Principais

#### GET
- **Uso**: Solicitar dados de um servidor
- **Características**: Não modifica dados no servidor
- **Exemplo**: Acessar uma página web, buscar informações

#### POST
- **Uso**: Enviar dados para o servidor
- **Características**: Pode modificar dados no servidor
- **Exemplo**: Enviar formulário, criar novo registro

#### PUT
- **Uso**: Atualizar um recurso existente
- **Características**: Substitui completamente o recurso

#### DELETE
- **Uso**: Remover um recurso
- **Características**: Deleta o recurso especificado

### Códigos de Status HTTP

#### Códigos 2xx (Sucesso)
- **200 OK**: Requisição bem-sucedida
- **201 Created**: Recurso criado com sucesso
- **204 No Content**: Sucesso sem conteúdo de retorno

#### Códigos 3xx (Redirecionamento)
- **301 Moved Permanently**: Recurso movido permanentemente
- **302 Found**: Redirecionamento temporário
- **304 Not Modified**: Recurso não modificado (cache)

#### Códigos 4xx (Erro do Cliente)
- **400 Bad Request**: Requisição inválida
- **401 Unauthorized**: Não autenticado
- **403 Forbidden**: Acesso negado
- **404 Not Found**: Recurso não encontrado
- **418 I'm a teapot**: (Código de brincadeira do HTTP)

#### Códigos 5xx (Erro do Servidor)
- **500 Internal Server Error**: Erro interno do servidor
- **502 Bad Gateway**: Erro no gateway
- **503 Service Unavailable**: Serviço indisponível

### HTTPS: HTTP Seguro

**HTTPS** (HTTP Secure) é HTTP com criptografia SSL/TLS:

- **Criptografia**: Dados criptografados durante a transmissão
- **Autenticação**: Verifica a identidade do servidor
- **Integridade**: Garante que os dados não foram alterados
- **Porta**: Usa a porta 443
- **Indicador Visual**: Cadeado no navegador

### Versões do HTTP

#### HTTP/1.1 (Atual Padrão)
- Conexões persistentes
- Compressão de conteúdo
- Cache melhorado

#### HTTP/2 (2015)
- **Multiplexing**: Múltiplas requisições em uma conexão
- **Server Push**: Servidor envia recursos antes de serem solicitados
- **Header Compression**: Compressão de headers
- **Melhor Performance**: Redução significativa de latência

#### HTTP/3 (2020)
- Baseado em **QUIC** (protocolo UDP)
- Conexões mais rápidas
- Melhor desempenho em conexões instáveis

---

## 🌍 Nomes de Domínio: Endereços Amigáveis da Web

### O que é um Nome de Domínio?

Um **nome de domínio** (domain name) é um endereço legível por humanos que substitui os endereços IP numéricos difíceis de lembrar. É como um "apelido" para o endereço IP do servidor.

### Estrutura de um Domínio

```
www.exemplo.com.br
│   │       │   │
│   │       │   └─ TLD (Top-Level Domain) de país (.br)
│   │       └───── TLD genérico (.com)
│   └───────────── Domínio de segundo nível (exemplo)
└───────────────── Subdomínio (www)
```

### Componentes de um Domínio

#### 1. TLD (Top-Level Domain)
- **TLDs Genéricos**: .com, .org, .net, .edu, .gov
- **TLDs de País**: .br (Brasil), .us (EUA), .uk (Reino Unido)
- **TLDs Modernos**: .tech, .dev, .app, .io

#### 2. Domínio de Segundo Nível (SLD)
- É a parte principal do domínio
- Exemplo: "exemplo" em "exemplo.com"
- Geralmente representa a marca ou organização

#### 3. Subdomínio
- Parte opcional antes do domínio principal
- Exemplos: www, blog, api, mail
- Permite organizar diferentes serviços

### Exemplos de Domínios

- `google.com` - Domínio principal do Google
- `www.google.com` - Subdomínio www do Google
- `mail.google.com` - Subdomínio de email do Google
- `exemplo.com.br` - Domínio brasileiro
- `blog.exemplo.com.br` - Subdomínio de blog

### Registro de Domínios

#### Como Funciona o Registro?

1. **Escolha do Domínio**: Você escolhe um nome disponível
2. **Registrador**: Compra através de um registrador (Registro.br, GoDaddy, etc.)
3. **Registro**: O domínio é registrado em um banco de dados global
4. **Renovação**: Domínios precisam ser renovados periodicamente

#### Informações Importantes

- **Propriedade**: Você "aluga" o domínio, não o possui permanentemente
- **Renovação**: Geralmente anual ou bianual
- **WHOIS**: Informações públicas sobre o proprietário do domínio
- **DNS**: Configuração de DNS aponta o domínio para um servidor

### Por que Domínios são Importantes?

- **Marca**: Representa sua identidade online
- **Memorização**: Mais fácil de lembrar que um IP
- **Profissionalismo**: Domínio próprio transmite credibilidade
- **SEO**: Domínios relevantes podem ajudar no SEO
- **Email**: Permite criar emails personalizados (contato@exemplo.com)

---

## 🖥️ Hospedagem Web: Onde os Sites Vivem

### O que é Hospedagem Web?

**Hospedagem web** (web hosting) é o serviço que armazena os arquivos de um website em servidores conectados à internet, tornando o site acessível 24/7 para visitantes ao redor do mundo.

### Tipos de Hospedagem

#### 1. Hospedagem Compartilhada (Shared Hosting)

**Características:**
- Múltiplos sites compartilham o mesmo servidor
- Recursos (CPU, memória, disco) compartilhados
- **Vantagens**: Custo baixo, fácil de gerenciar, ideal para iniciantes
- **Desvantagens**: Performance limitada, menos controle, impacto de outros sites

**Ideal para:**
- Sites pessoais
- Blogs
- Pequenos negócios
- Sites com tráfego baixo a médio

#### 2. Hospedagem Dedicada (Dedicated Hosting)

**Características:**
- Servidor exclusivo para seu site
- Controle total sobre recursos e configurações
- **Vantagens**: Performance máxima, controle total, segurança
- **Desvantagens**: Custo alto, requer conhecimento técnico

**Ideal para:**
- Grandes empresas
- Sites com alto tráfego
- Aplicações que requerem configurações específicas

#### 3. VPS (Virtual Private Server)

**Características:**
- Servidor virtual privado dentro de um servidor físico
- Recursos dedicados (mas compartilha hardware)
- **Vantagens**: Mais controle que compartilhado, melhor performance, custo intermediário
- **Desvantagens**: Requer conhecimento técnico, gerenciamento necessário

**Ideal para:**
- Desenvolvedores
- Sites em crescimento
- Aplicações que precisam de mais controle

#### 4. Cloud Hosting (Hospedagem em Nuvem)

**Características:**
- Recursos distribuídos em múltiplos servidores
- Escalabilidade automática
- **Vantagens**: Alta disponibilidade, escalável, paga pelo que usa
- **Desvantagens**: Custo pode variar, complexidade

**Ideal para:**
- Aplicações modernas
- Sites com tráfego variável
- Startups e empresas em crescimento

### Componentes da Hospedagem

#### Espaço em Disco
- Armazena arquivos do site (HTML, CSS, imagens, etc.)
- Medido em GB (Gigabytes) ou TB (Terabytes)

#### Largura de Banda (Bandwidth)
- Quantidade de dados que podem ser transferidos
- Importante para sites com muito tráfego

#### Banco de Dados
- Armazenamento para dados dinâmicos
- MySQL, PostgreSQL, MongoDB, etc.

#### Email
- Contas de email personalizadas (@seudominio.com)
- Geralmente incluído em planos de hospedagem

#### SSL/TLS
- Certificados de segurança (HTTPS)
- Muitas vezes incluído gratuitamente

### Serviços Adicionais de Hospedagem

- **CDN (Content Delivery Network)**: Distribui conteúdo globalmente
- **Backup Automático**: Cópias de segurança regulares
- **Suporte Técnico**: Assistência para problemas
- **Painel de Controle**: Interface para gerenciar o site (cPanel, Plesk)

### Escolhendo um Provedor de Hospedagem

**Fatores a Considerar:**
- **Custo**: Preço mensal/anual
- **Performance**: Velocidade e uptime (tempo online)
- **Suporte**: Qualidade e disponibilidade do suporte
- **Recursos**: Espaço, banda, banco de dados
- **Facilidade de Uso**: Interface amigável
- **Localização**: Proximidade dos servidores aos usuários

---

## 🔍 DNS: O Sistema de Nomes de Domínio

### O que é DNS?

**DNS** (Domain Name System - Sistema de Nomes de Domínio) é um sistema distribuído que traduz nomes de domínio legíveis por humanos (como `google.com`) em endereços IP numéricos (como `142.250.191.46`) que os computadores usam para se comunicar.

### Por que DNS é Necessário?

**Sem DNS:**
```
Você digitaria: 142.250.191.46
Difícil de lembrar, não é?
```

**Com DNS:**
```
Você digita: google.com
DNS traduz para: 142.250.191.46
Muito mais fácil!
```

### Como Funciona o DNS?

#### 1. Requisição DNS

Quando você digita `exemplo.com` no navegador:

1. **Navegador**: "Preciso do IP de exemplo.com"
2. **Resolver DNS Local**: Verifica cache local
3. **Servidor DNS Raiz**: Direciona para servidor TLD (.com)
4. **Servidor DNS TLD**: Direciona para servidor autoritativo
5. **Servidor DNS Autoritativo**: Retorna o IP correto
6. **Navegador**: Recebe o IP e faz a requisição HTTP

#### 2. Cache DNS

- **Cache Local**: Seu computador armazena traduções recentes
- **Cache do Roteador**: Roteador também pode ter cache
- **TTL (Time To Live)**: Tempo que uma entrada fica no cache
- **Benefício**: Reduz tempo de resolução para acessos repetidos

### Tipos de Registros DNS

#### A Record (Address Record)
- Mapeia domínio para endereço IPv4
- Exemplo: `exemplo.com` → `192.168.1.1`

#### AAAA Record
- Mapeia domínio para endereço IPv6
- Exemplo: `exemplo.com` → `2001:0db8::1`

#### CNAME Record (Canonical Name)
- Cria um alias (apelido) para outro domínio
- Exemplo: `www.exemplo.com` → `exemplo.com`

#### MX Record (Mail Exchange)
- Especifica servidores de email
- Exemplo: `exemplo.com` → `mail.exemplo.com`

#### TXT Record
- Armazena informações textuais
- Usado para verificação, SPF, DKIM, etc.

### Servidores DNS Públicos

#### Google DNS
- **IPv4**: 8.8.8.8 e 8.8.4.4
- **IPv6**: 2001:4860:4860::8888

#### Cloudflare DNS
- **IPv4**: 1.1.1.1 e 1.0.0.1
- **IPv6**: 2606:4700:4700::1111

#### OpenDNS
- **IPv4**: 208.67.222.222 e 208.67.220.220

### DNS e Performance

- **Latência DNS**: Tempo para resolver um domínio
- **Cache**: Reduz requisições DNS repetidas
- **CDN**: Distribui conteúdo próximo aos usuários
- **DNS Prefetch**: Navegador resolve domínios antecipadamente

---

## 🌐 Navegadores Web: Interpretadores de HTML

### O que é um Navegador?

Um **navegador web** (web browser) é um software que solicita, recebe, interpreta e exibe páginas web. É a interface entre você e a World Wide Web.

### Funções Principais de um Navegador

1. **Solicitar Recursos**: Faz requisições HTTP/HTTPS
2. **Interpretar HTML/CSS/JavaScript**: Processa e renderiza o código
3. **Gerenciar Cache**: Armazena recursos para acesso rápido
4. **Gerenciar Cookies**: Armazena dados de sessão
5. **Segurança**: Protege contra sites maliciosos
6. **Extensões**: Permite adicionar funcionalidades

### Componentes de um Navegador

#### 1. Interface do Usuário
- Barra de endereço (URL bar)
- Botões de navegação (voltar, avançar, recarregar)
- Abas (tabs)
- Menu e configurações

#### 2. Motor de Renderização (Rendering Engine)
- **Blink**: Chrome, Edge, Opera
- **Gecko**: Firefox
- **WebKit**: Safari
- **Trident/EdgeHTML**: Internet Explorer (descontinuado)

**Função**: Interpreta HTML e CSS e renderiza na tela

#### 3. Motor JavaScript
- **V8**: Chrome, Edge, Opera
- **SpiderMonkey**: Firefox
- **JavaScriptCore**: Safari

**Função**: Executa código JavaScript

#### 4. Camada de Rede
- Faz requisições HTTP/HTTPS
- Gerencia conexões
- Implementa cache

#### 5. Backend de UI
- Desenha widgets da interface
- Gerencia janelas e diálogos

### Navegadores Principais

#### Google Chrome
- **Motor**: Blink (V8)
- **Market Share**: ~65%
- **Características**: Rápido, extensões, DevTools excelentes

#### Mozilla Firefox
- **Motor**: Gecko (SpiderMonkey)
- **Market Share**: ~3%
- **Características**: Foco em privacidade, open source

#### Microsoft Edge
- **Motor**: Blink (V8)
- **Market Share**: ~5%
- **Características**: Integração com Windows, performance

#### Safari
- **Motor**: WebKit (JavaScriptCore)
- **Market Share**: ~19%
- **Características**: Otimizado para macOS/iOS, eficiência energética

#### Opera
- **Motor**: Blink (V8)
- **Market Share**: ~2%
- **Características**: VPN integrado, bloqueador de anúncios

### Como Navegadores Processam uma Página Web

#### 1. Parsing (Análise)
- Navegador recebe HTML
- Analisa a estrutura (parsing)
- Cria árvore DOM (Document Object Model)

#### 2. Renderização
- Aplica CSS (cria árvore de renderização)
- Calcula layout (onde cada elemento fica)
- Pinta pixels na tela

#### 3. Execução JavaScript
- Executa scripts JavaScript
- Pode modificar DOM dinamicamente
- Pode fazer requisições adicionais

#### 4. Otimizações
- Lazy loading de imagens
- Cache de recursos
- Compressão de dados

### DevTools: Ferramentas do Desenvolvedor

Navegadores modernos incluem **DevTools** (F12):

- **Inspector**: Inspeciona HTML e CSS
- **Console**: Executa JavaScript e mostra erros
- **Network**: Monitora requisições HTTP
- **Performance**: Analisa performance da página
- **Application**: Gerencia cache, cookies, storage

### Compatibilidade entre Navegadores

- **Diferenças**: Navegadores podem renderizar HTML/CSS de forma ligeiramente diferente
- **Prefixos CSS**: Algumas propriedades precisam de prefixos (-webkit-, -moz-)
- **Testes**: Sempre teste em múltiplos navegadores
- **Polyfills**: Código que adiciona funcionalidades faltantes

---

## 🔎 SEO: Otimização para Mecanismos de Busca

### O que é SEO?

**SEO** (Search Engine Optimization - Otimização para Mecanismos de Busca) é a prática de melhorar seu website para aumentar sua visibilidade quando pessoas pesquisam produtos, serviços ou informações relacionadas ao seu negócio em mecanismos de busca como Google, Bing, etc.

### Por que SEO é Importante?

- **Visibilidade**: Quanto melhor sua posição nos resultados, mais pessoas veem seu site
- **Tráfego Orgânico**: Visitantes que chegam sem pagar por anúncios
- **Credibilidade**: Sites nas primeiras posições são vistos como mais confiáveis
- **ROI**: Investimento que continua gerando resultados ao longo do tempo

### Como Funcionam os Mecanismos de Busca?

#### 1. Rastreamento (Crawling)
- **Bots** (robôs) visitam páginas web
- Seguem links entre páginas
- Coletam informações sobre o conteúdo

#### 2. Indexação
- Mecanismos de busca organizam o conteúdo coletado
- Criam um índice (catálogo) de todas as páginas
- Armazenam informações sobre palavras-chave, títulos, etc.

#### 3. Classificação (Ranking)
- Quando alguém pesquisa, o mecanismo busca no índice
- Classifica resultados por relevância e qualidade
- Exibe os resultados mais relevantes primeiro

### Fatores de SEO

#### SEO On-Page (Na Página)

**HTML e Estrutura:**
- **Títulos (H1-H6)**: Hierarquia clara e relevante
- **Meta Tags**: Description, keywords, title
- **URLs Amigáveis**: URLs descritivas e limpas
- **Estrutura Semântica**: Uso correto de elementos HTML5
- **Alt Text em Imagens**: Descrições para imagens
- **Links Internos**: Conexões entre páginas do site

**Conteúdo:**
- **Qualidade**: Conteúdo original, útil e relevante
- **Palavras-chave**: Uso natural de termos de busca
- **Atualização**: Conteúdo fresco e atualizado
- **Comprimento**: Conteúdo completo e detalhado

#### SEO Técnico

**Performance:**
- **Velocidade de Carregamento**: Páginas rápidas ranqueiam melhor
- **Mobile-Friendly**: Sites responsivos são priorizados
- **HTTPS**: Sites seguros têm vantagem
- **Core Web Vitals**: Métricas de experiência do usuário

**Estrutura:**
- **Sitemap XML**: Mapa do site para mecanismos de busca
- **Robots.txt**: Instruções para bots
- **Schema Markup**: Dados estruturados (JSON-LD)

#### SEO Off-Page (Fora da Página)

- **Backlinks**: Links de outros sites apontando para o seu
- **Redes Sociais**: Compartilhamentos e menções
- **Autoridade de Domínio**: Reputação do domínio
- **Local SEO**: Otimização para buscas locais

### Meta Tags Essenciais para SEO

```html
<head>
    <!-- Título da página (aparece na aba do navegador e nos resultados de busca) -->
    <title>Título Otimizado da Página - Palavra-chave Principal</title>
    
    <!-- Descrição (aparece nos resultados de busca) -->
    <meta name="description" content="Descrição clara e atrativa de 150-160 caracteres com palavras-chave relevantes">
    
    <!-- Palavras-chave (menos importante hoje, mas ainda usado) -->
    <meta name="keywords" content="palavra-chave1, palavra-chave2, palavra-chave3">
    
    <!-- Viewport (essencial para mobile) -->
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    
    <!-- Charset (codificação de caracteres) -->
    <meta charset="UTF-8">
    
    <!-- Open Graph (para redes sociais) -->
    <meta property="og:title" content="Título para Compartilhamento">
    <meta property="og:description" content="Descrição para Compartilhamento">
    <meta property="og:image" content="https://exemplo.com/imagem.jpg">
</head>
```

### Boas Práticas de SEO em HTML

#### 1. Estrutura de Títulos Hierárquica
```html
<h1>Título Principal (um por página)</h1>
<h2>Subtítulo Principal</h2>
<h3>Subtítulo Secundário</h3>
```

#### 2. URLs Amigáveis
```
❌ Ruim: exemplo.com/pagina.php?id=123&cat=abc
✅ Bom: exemplo.com/produtos/categoria/nome-produto
```

#### 3. Alt Text em Imagens
```html
<img src="produto.jpg" alt="Produto X - Descrição detalhada">
```

#### 4. Links com Texto Descritivo
```html
❌ Ruim: <a href="/sobre">Clique aqui</a>
✅ Bom: <a href="/sobre">Saiba mais sobre nossa empresa</a>
```

### Core Web Vitals

Métricas importantes do Google:

- **LCP (Largest Contentful Paint)**: Tempo para carregar o conteúdo principal (< 2.5s)
- **FID (First Input Delay)**: Tempo de resposta à primeira interação (< 100ms)
- **CLS (Cumulative Layout Shift)**: Estabilidade visual (< 0.1)

### SEO Local

Para negócios físicos:
- **Google My Business**: Perfil no Google
- **Informações de Contato**: Endereço, telefone claramente visíveis
- **Schema Markup**: Dados estruturados para localização

### GEO: Generative Engine Optimization

**GEO** (Generative Engine Optimization) é uma evolução do SEO focada em otimizar conteúdo para experiências de busca alimentadas por IA, como ChatGPT, Bing Chat, etc.

**Diferenças:**
- **SEO Tradicional**: Otimiza para listas de resultados
- **GEO**: Otimiza para respostas geradas por IA

**Estratégias GEO:**
- Conteúdo autoritativo e completo
- Estrutura clara e bem organizada
- Dados estruturados (Schema)
- Informações factuais e verificáveis

---

## 🔗 Como Tudo se Conecta: O Fluxo Completo

### O Que Acontece Quando Você Acessa um Site?

Vamos acompanhar o que acontece quando você digita `exemplo.com` no navegador:

#### Passo 1: Digitação da URL
```
Você digita: exemplo.com
```

#### Passo 2: Resolução DNS
```
Navegador → DNS: "Qual é o IP de exemplo.com?"
DNS → Navegador: "192.168.1.100"
```

#### Passo 3: Requisição HTTP
```
Navegador → Servidor (192.168.1.100): 
GET /index.html HTTP/1.1
Host: exemplo.com
```

#### Passo 4: Resposta do Servidor
```
Servidor → Navegador:
HTTP/1.1 200 OK
Content-Type: text/html

<!DOCTYPE html>
<html>...
```

#### Passo 5: Parsing e Renderização
```
Navegador:
1. Analisa HTML (parsing)
2. Aplica CSS
3. Executa JavaScript
4. Renderiza na tela
```

#### Passo 6: Requisições Adicionais
```
Navegador pode solicitar:
- CSS externo
- JavaScript externo
- Imagens
- Fontes
- Outros recursos
```

### Tempo Total

- **DNS Lookup**: ~20-120ms
- **Conexão TCP**: ~100-200ms
- **Requisição HTTP**: ~50-200ms
- **Download**: Depende do tamanho
- **Parsing/Renderização**: ~50-300ms

**Total**: Geralmente 1-3 segundos para uma página simples

---

## 📝 Resumo da Aula

Nesta aula, você aprendeu:

✅ **Internet** é a rede global que conecta computadores  
✅ **HTTP** é o protocolo de comunicação entre navegador e servidor  
✅ **Domínios** são endereços amigáveis que substituem IPs  
✅ **Hospedagem** armazena arquivos do site em servidores  
✅ **DNS** traduz nomes de domínio em endereços IP  
✅ **Navegadores** interpretam HTML, CSS e JavaScript  
✅ **SEO** otimiza sites para mecanismos de busca  

### Conceitos-Chave

- **Request/Response**: Modelo de comunicação HTTP
- **DNS Resolution**: Tradução de domínio para IP
- **Rendering Engine**: Motor que renderiza páginas
- **On-Page SEO**: Otimizações no código HTML
- **Core Web Vitals**: Métricas de performance do Google

### Próximos Passos

Agora que você entende como a web funciona, na próxima aula você aprenderá:
- Estrutura detalhada de documentos HTML
- Elementos semânticos HTML5
- Como criar páginas bem estruturadas
- Boas práticas de marcação

---

## 🔍 Glossário

- **ISP (Internet Service Provider)**: Provedor de serviços de internet
- **TCP/IP**: Protocolo de comunicação da internet
- **SSL/TLS**: Protocolos de criptografia para HTTPS
- **CDN (Content Delivery Network)**: Rede de distribuição de conteúdo
- **Uptime**: Tempo que um servidor fica online
- **Bandwidth**: Largura de banda, capacidade de transferência
- **Cache**: Armazenamento temporário para acesso rápido
- **Backlink**: Link de outro site apontando para o seu
- **Sitemap**: Mapa do site para mecanismos de busca
- **Schema Markup**: Dados estruturados em formato JSON-LD

