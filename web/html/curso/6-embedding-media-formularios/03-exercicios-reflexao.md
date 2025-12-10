# Aula 6 - Exercícios e Reflexão

## 🎯 Exercícios Práticos

### Exercício 1: Criando uma Galeria de Imagens

Crie uma página HTML com uma galeria de imagens usando `<figure>` e `<figcaption>`.

**Requisitos:**
- Crie uma página HTML completa com estrutura básica
- Adicione um título principal: "Minha Galeria de Fotos"
- Crie 3 seções de imagens usando `<figure>`, cada uma com:
  - Uma imagem (use imagens de exemplo ou placeholders)
  - Um `<figcaption>` descritivo
- Use atributos apropriados: `alt`, `width`, `height`, `loading="lazy"`
- Adicione uma imagem hero no topo com `fetchpriority="high"`

**Dica:** Use placeholders como `https://via.placeholder.com/800x600` se não tiver imagens próprias.

---

### Exercício 2: Player de Áudio com Múltiplas Fontes

Crie um player de áudio que suporte múltiplos formatos para garantir compatibilidade.

**Requisitos:**
- Crie um elemento `<audio>` com controles
- Adicione 3 fontes diferentes usando `<source>`:
  - MP3 (type="audio/mpeg")
  - OGG (type="audio/ogg")
  - WAV (type="audio/wav")
- Configure `preload="metadata"`
- Adicione texto alternativo caso o navegador não suporte áudio
- Adicione um `<label>` descritivo antes do player

**Desafio extra:** Se você tiver arquivos de áudio, use-os. Caso contrário, explique como você estruturaria o código.

---

### Exercício 3: Vídeo com Legendas e Controles

Crie um player de vídeo completo com suporte a legendas.

**Requisitos:**
- Crie um elemento `<video>` com:
  - Controles visíveis
  - Largura de 1280px e altura de 720px
  - Uma imagem de poster
  - `preload="metadata"`
- Adicione múltiplas fontes de vídeo (MP4 e WebM)
- Adicione pelo menos uma faixa de legendas usando `<track>` (você pode criar um arquivo .vtt simples ou apenas estruturar o código)
- Adicione texto alternativo para navegadores que não suportam vídeo

**Nota:** Se você não tiver arquivos de vídeo reais, crie a estrutura completa do código explicando onde cada arquivo estaria.

---

### Exercício 4: Formulário de Cadastro Completo

Crie um formulário de cadastro completo com validação HTML5.

**Requisitos:**
- Crie um formulário com `action="/cadastrar"` e `method="post"`
- Adicione os seguintes campos com labels apropriados:
  1. **Nome completo** (text, required, minlength 3, maxlength 100)
  2. **Email** (email, required)
  3. **Senha** (password, required, minlength 8, pattern para senha forte)
  4. **Confirmar senha** (password, required)
  5. **Data de nascimento** (date, required, min="1900-01-01", max="2024-12-31")
  6. **Telefone** (tel, pattern para formato brasileiro)
  7. **Gênero** (radio buttons: Masculino, Feminino, Outro, Prefiro não informar)
  8. **Aceito os termos** (checkbox, required)
  9. **Newsletter** (checkbox opcional)
- Agrupe campos relacionados usando `<fieldset>` e `<legend>`
- Adicione botões de "Enviar" e "Limpar"
- Adicione placeholders apropriados onde necessário

**Desafio extra:** Adicione validação JavaScript para garantir que "Senha" e "Confirmar senha" sejam iguais.

---

### Exercício 5: Formulário de Upload com Preview

Crie um formulário de upload de imagem com preview da imagem selecionada.

**Requisitos:**
- Crie um formulário com `enctype="multipart/form-data"`
- Adicione um campo de upload de arquivo que aceite apenas imagens
- Adicione um elemento `<div>` para exibir o preview da imagem
- Use JavaScript para:
  - Detectar quando uma imagem é selecionada
  - Exibir um preview da imagem antes do envio
  - Validar que o arquivo é uma imagem
  - Mostrar mensagem de erro se não for imagem
- Adicione um botão de submit

**Dica:** Use `FileReader` API do JavaScript para ler e exibir a imagem.

---

### Exercício 6: Incorporando Conteúdo com iframe

Crie uma página que demonstre diferentes usos de iframe.

**Requisitos:**
- Crie uma página HTML completa
- Adicione 3 exemplos de iframe:
  1. Um iframe incorporando um vídeo do YouTube (use um ID de vídeo real ou explique a estrutura)
  2. Um iframe incorporando um mapa do Google Maps (use coordenadas ou explique a estrutura)
  3. Um iframe incorporando outra página HTML local (crie uma página simples para isso)
- Para cada iframe:
  - Adicione `title` apropriado para acessibilidade
  - Defina `width` e `height`
  - Use `sandbox` onde apropriado para segurança
- Adicione uma seção explicando quando usar cada tipo de iframe

---

## 🤔 Perguntas de Reflexão

### Reflexão 1: Acessibilidade em Imagens

**Cenário:** Você está criando um site de e-commerce e precisa adicionar imagens de produtos. Analise as seguintes abordagens:

**Abordagem A:**
```html
<img src="produto.jpg" alt="Produto">
```

**Abordagem B:**
```html
<img src="produto.jpg" alt="Camiseta de algodão 100% orgânico, cor azul marinho, tamanho M, marca EcoWear, preço R$ 89,90">
```

**Abordagem C:**
```html
<figure>
    <img src="produto.jpg" alt="Camiseta EcoWear azul marinho tamanho M">
    <figcaption>Camiseta de algodão orgânico - R$ 89,90</figcaption>
</figure>
```

**Perguntas:**
1. Qual abordagem é mais apropriada para um site de e-commerce? Por quê?
2. Como cada abordagem afeta usuários com leitores de tela?
3. Qual é o impacto de cada abordagem no SEO?
4. Como você equilibraria informação útil com concisão no atributo `alt`?
5. Quando seria apropriado usar `alt=""` (vazio) em uma imagem?
6. Qual é a diferença prática entre usar `alt` descritivo e `figcaption`?

**Resposta esperada:** Explique sua escolha considerando acessibilidade, SEO, experiência do usuário e boas práticas.

---

### Reflexão 2: Performance e Carregamento de Mídia

**Cenário:** Você está criando uma página de portfólio de fotografia com 50 imagens de alta resolução. Analise as seguintes estratégias:

**Estratégia A:**
```html
<img src="foto1.jpg" alt="Foto 1">
<img src="foto2.jpg" alt="Foto 2">
<!-- ... 48 imagens mais ... -->
```

**Estratégia B:**
```html
<img src="foto-hero.jpg" alt="Foto principal" fetchpriority="high" width="1920" height="1080">
<img src="foto1.jpg" alt="Foto 1" loading="lazy" width="800" height="600">
<img src="foto2.jpg" alt="Foto 2" loading="lazy" width="800" height="600">
<!-- ... -->
```

**Estratégia C:**
```html
<img src="foto-hero.jpg" alt="Foto principal" fetchpriority="high" width="1920" height="1080">
<!-- Usar JavaScript para carregar imagens sob demanda quando usuário rola a página -->
```

**Perguntas:**
1. Qual estratégia oferece melhor performance inicial? Por quê?
2. Como `loading="lazy"` melhora a experiência do usuário?
3. Por que especificar `width` e `height` é importante para performance?
4. Qual é o impacto de cada estratégia no Cumulative Layout Shift (CLS)?
5. Como você priorizaria o carregamento de imagens em uma página com muitas imagens?
6. Quando seria apropriado usar JavaScript para carregamento dinâmico vs. atributos HTML nativos?
7. Qual é o trade-off entre performance e experiência do usuário em cada estratégia?

**Resposta esperada:** Analise cada estratégia considerando métricas de performance (LCP, CLS, FCP) e experiência do usuário.

---

### Reflexão 3: Validação de Formulários - Cliente vs. Servidor

**Cenário:** Você está criando um formulário de cadastro de usuário. Analise as seguintes abordagens de validação:

**Abordagem A: Apenas Validação HTML5**
```html
<input type="email" name="email" required>
<input type="password" name="senha" required minlength="8">
```

**Abordagem B: Validação HTML5 + JavaScript**
```html
<input type="email" name="email" required>
<input type="password" name="senha" required minlength="8">
<script>
// Validação adicional com JavaScript
</script>
```

**Abordagem C: Validação Completa (HTML5 + JavaScript + Servidor)**
```html
<!-- HTML5 e JavaScript no cliente -->
<!-- + Validação no servidor (PHP, Node.js, etc.) -->
```

**Perguntas:**
1. Por que validação apenas no cliente (HTML5/JavaScript) não é suficiente para segurança?
2. Qual é o papel de cada tipo de validação (HTML5, JavaScript, Servidor)?
3. Como um atacante poderia contornar validação apenas no cliente?
4. Qual é o impacto de cada abordagem na experiência do usuário?
5. Como você equilibraria feedback imediato (validação no cliente) com segurança (validação no servidor)?
6. Quando seria apropriado usar `novalidate` no formulário?
7. Qual é a importância de mensagens de erro claras e úteis?

**Resposta esperada:** Explique a importância de validação em múltiplas camadas e como cada camada serve a um propósito diferente.

---

### Reflexão 4: Segurança com iframe e Content Security Policy

**Cenário:** Você está criando um site que precisa incorporar conteúdo de fontes externas (vídeos do YouTube, mapas do Google, widgets de terceiros). Analise as considerações de segurança:

**Situação A: iframe sem restrições**
```html
<iframe src="https://site-externo.com/widget"></iframe>
```

**Situação B: iframe com sandbox**
```html
<iframe src="https://site-externo.com/widget" sandbox="allow-scripts"></iframe>
```

**Situação C: iframe + CSP**
```html
<meta http-equiv="Content-Security-Policy" content="frame-src 'self' https://www.youtube.com">
<iframe src="https://www.youtube.com/embed/VIDEO_ID"></iframe>
```

**Perguntas:**
1. Quais são os riscos de segurança ao incorporar conteúdo de fontes não confiáveis?
2. Como o atributo `sandbox` protege sua página?
3. Qual é o papel do Content Security Policy (CSP) na segurança?
4. Como você decidiria quais fontes confiar para incorporação?
5. Qual é o impacto de CSP muito restritivo vs. muito permissivo?
6. Como você testaria se suas políticas de segurança estão funcionando?
7. Qual é a importância de manter CSP atualizado quando você adiciona novos recursos externos?

**Resposta esperada:** Explique os riscos de segurança, como CSP e sandbox protegem contra ataques, e como equilibrar funcionalidade com segurança.

---

### Reflexão 5: Acessibilidade em Formulários

**Cenário:** Você está criando um formulário de contato. Analise as seguintes implementações:

**Implementação A:**
```html
<input type="text" name="nome" placeholder="Nome">
<input type="email" name="email" placeholder="Email">
<button type="submit">Enviar</button>
```

**Implementação B:**
```html
<label for="nome">Nome completo:</label>
<input type="text" id="nome" name="nome" required>
<label for="email">Email:</label>
<input type="email" id="email" name="email" required>
<button type="submit">Enviar</button>
```

**Implementação C:**
```html
<fieldset>
    <legend>Informações de Contato</legend>
    <label for="nome">Nome completo: <span aria-label="obrigatório">*</span></label>
    <input type="text" id="nome" name="nome" required aria-required="true">
    <label for="email">Email: <span aria-label="obrigatório">*</span></label>
    <input type="email" id="email" name="email" required aria-required="true">
</fieldset>
<button type="submit">Enviar formulário</button>
```

**Perguntas:**
1. Como cada implementação afeta usuários com leitores de tela?
2. Por que `placeholder` não substitui `<label>`?
3. Qual é a importância de agrupar campos relacionados com `<fieldset>`?
4. Como `aria-required` complementa o atributo `required`?
5. Qual é o impacto de cada implementação na usabilidade geral?
6. Como você indicaria campos obrigatórios de forma acessível?
7. Qual é a importância de mensagens de erro acessíveis em formulários?

**Resposta esperada:** Explique como cada elemento melhora a acessibilidade e por que formulários acessíveis beneficiam todos os usuários, não apenas aqueles com deficiências.

---

### Reflexão 6: Otimização de Mídia para Diferentes Dispositivos

**Cenário:** Você está criando um site que será acessado de diferentes dispositivos (desktop, tablet, mobile) com diferentes velocidades de conexão. Analise as estratégias:

**Estratégia A: Uma imagem para todos**
```html
<img src="imagem-grande.jpg" alt="Imagem">
```

**Estratégia B: Responsive images com srcset**
```html
<img 
    srcset="imagem-pequena.jpg 400w,
            imagem-media.jpg 800w,
            imagem-grande.jpg 1200w"
    sizes="(max-width: 600px) 400px,
           (max-width: 1200px) 800px,
           1200px"
    src="imagem-grande.jpg"
    alt="Imagem responsiva"
>
```

**Estratégia C: Picture element com diferentes formatos**
```html
<picture>
    <source media="(max-width: 600px)" srcset="imagem-mobile.webp">
    <source media="(max-width: 1200px)" srcset="imagem-tablet.webp">
    <img src="imagem-desktop.jpg" alt="Imagem">
</picture>
```

**Perguntas:**
1. Qual é o impacto de cada estratégia no tempo de carregamento em dispositivos móveis?
2. Como `srcset` e `sizes` melhoram a performance?
3. Qual é a importância de oferecer diferentes formatos de imagem (WebP vs. JPEG)?
4. Como você decidiria quais tamanhos de imagem criar?
5. Qual é o trade-off entre qualidade de imagem e tamanho de arquivo?
6. Como você testaria se suas otimizações estão funcionando em diferentes dispositivos?
7. Qual é a importância de considerar conexões lentas ao otimizar mídia?

**Resposta esperada:** Explique como otimizar mídia para diferentes contextos e por que isso é crucial para experiência do usuário e performance.

---

## ✅ Checklist de Aprendizado

Use este checklist para verificar seu entendimento dos conceitos desta aula:

### Embedding Media
- [ ] Entendo o que é embedding media e por que é importante
- [ ] Sei usar o elemento `<img>` com atributos apropriados
- [ ] Compreendo a diferença entre `<img>` e `<figure>`
- [ ] Sei usar `fetchpriority` para otimizar carregamento
- [ ] Posso criar players de áudio com múltiplas fontes
- [ ] Posso criar players de vídeo com controles e legendas
- [ ] Entendo como usar iframe de forma segura
- [ ] Compreendo o básico de Content Security Policy

### Formulários
- [ ] Sei criar formulários HTML básicos
- [ ] Entendo a importância de usar `<label>` corretamente
- [ ] Conheço os diferentes tipos de input e quando usar cada um
- [ ] Posso criar campos de upload de arquivos
- [ ] Entendo validação HTML5 nativa
- [ ] Sei aplicar restrições apropriadas aos campos
- [ ] Compreendo a importância de validação no servidor
- [ ] Posso criar formulários acessíveis

### Boas Práticas
- [ ] Sempre adiciono `alt` descritivo em imagens
- [ ] Especifico `width` e `height` em imagens e vídeos
- [ ] Uso `loading="lazy"` para imagens abaixo da dobra
- [ ] Valido formulários tanto no cliente quanto no servidor
- [ ] Uso labels apropriados para acessibilidade
- [ ] Considero segurança ao usar iframe
- [ ] Otimizo mídia para diferentes dispositivos

---

## 🎯 Próximos Passos

Após completar os exercícios e refletir sobre as perguntas:

1. **Revise suas respostas** - Compare com as melhores práticas
2. **Experimente** - Crie seus próprios exemplos
3. **Teste em diferentes navegadores** - Verifique compatibilidade
4. **Valide seu código** - Use o W3C Validator
5. **Pense em acessibilidade** - Teste com leitores de tela se possível
6. **Considere performance** - Use DevTools para medir tempos de carregamento

---

**Lembre-se:** A prática é essencial! Crie seus próprios projetos aplicando esses conceitos. 🚀





