# Aula 2 - Exercícios e Reflexão

## 🎯 Exercícios Práticos

### Exercício 1: Analisando Requisições HTTP

Abra o DevTools do seu navegador (F12) e vá para a aba **Network** (Rede).

**Tarefa:**
1. Acesse um site qualquer (ex: google.com)
2. Observe as requisições HTTP que aparecem
3. Identifique:
   - Quantas requisições foram feitas?
   - Quais são os métodos HTTP usados (GET, POST, etc.)?
   - Quais são os códigos de status (200, 304, etc.)?
   - Quais tipos de arquivos foram solicitados (HTML, CSS, JS, imagens)?

**Desafio Extra:**
- Clique em uma requisição específica
- Analise os **Headers** (cabeçalhos) da requisição e resposta
- Identifique o **Content-Type** de diferentes recursos

**Dica:** Filtre por tipo de recurso (HTML, CSS, JS, Img) para ver melhor a organização.

---

### Exercício 2: Investigando DNS

Use ferramentas online ou linha de comando para investigar DNS.

**Opção A - Online:**
1. Acesse [whatsmydns.net](https://www.whatsmydns.net) ou [dnschecker.org](https://dnschecker.org)
2. Digite um domínio (ex: google.com)
3. Observe os diferentes endereços IP retornados

**Opção B - Linha de Comando:**
```bash
# No terminal (Mac/Linux) ou PowerShell (Windows)
nslookup google.com
# ou
dig google.com
```

**Tarefa:**
1. Verifique o IP de 3 domínios diferentes
2. Compare os IPs - eles são diferentes?
3. Tente acessar um site diretamente pelo IP (ex: digite o IP no navegador)
4. O que acontece? Por que isso pode não funcionar?

**Perguntas para Reflexão:**
- Por que um mesmo domínio pode ter IPs diferentes em locais diferentes?
- O que acontece se você acessar um site pelo IP ao invés do domínio?

---

### Exercício 3: Analisando Meta Tags de SEO

Escolha 3 sites diferentes (pode ser sites de notícias, blogs, e-commerce) e analise suas meta tags.

**Tarefa:**
1. Abra cada site no navegador
2. Clique com botão direito → "Inspecionar" ou pressione F12
3. Vá para a aba **Elements** (Elementos)
4. Expanda a tag `<head>`
5. Identifique e anote:
   - A tag `<title>`
   - A meta tag `description`
   - A meta tag `viewport`
   - Outras meta tags que encontrar

**Crie uma Tabela Comparativa:**

| Site | Title | Description | Viewport | Outras Meta Tags |
|------|-------|-------------|----------|------------------|
| Site 1 | ... | ... | ... | ... |
| Site 2 | ... | ... | ... | ... |
| Site 3 | ... | ... | ... | ... |

**Análise:**
- Qual site tem o título mais descritivo?
- Qual descrição é mais atrativa?
- Todos têm a meta tag viewport? Por que isso é importante?

---

### Exercício 4: Criando uma Página HTML com Boas Práticas de SEO

Crie um arquivo HTML para uma página sobre "Receitas Veganas" seguindo boas práticas de SEO.

**Requisitos:**
1. Estrutura HTML5 completa e válida
2. Meta tags essenciais:
   - `<title>` otimizado (50-60 caracteres)
   - `<meta name="description">` (150-160 caracteres)
   - `<meta name="viewport">`
   - `<meta charset="UTF-8">`
3. Estrutura semântica:
   - Um `<h1>` principal
   - `<h2>` para seções
   - Hierarquia correta de headings
4. Imagens com atributo `alt` descritivo
5. Links com texto descritivo (não "clique aqui")

**Conteúdo Sugerido:**
- Título: "Receitas Veganas Deliciosas e Fáceis"
- Seções: Introdução, Receita 1, Receita 2, Dicas
- Pelo menos 2 imagens com alt text apropriado
- Links internos entre seções

**Dica:** Use o [W3C Validator](https://validator.w3.org/) para validar seu código.

---

### Exercício 5: Comparando Tipos de Hospedagem

Pesquise sobre diferentes tipos de hospedagem web e crie uma comparação.

**Tarefa:**
1. Pesquise sobre:
   - Hospedagem Compartilhada
   - VPS (Virtual Private Server)
   - Servidor Dedicado
   - Cloud Hosting
2. Para cada tipo, identifique:
   - Custo aproximado (faixa de preço)
   - Performance esperada
   - Controle e flexibilidade
   - Ideal para que tipo de site/projeto
3. Crie uma tabela comparativa

**Tabela Sugerida:**

| Tipo | Custo | Performance | Controle | Ideal Para |
|------|-------|-------------|----------|------------|
| Compartilhada | ... | ... | ... | ... |
| VPS | ... | ... | ... | ... |
| Dedicado | ... | ... | ... | ... |
| Cloud | ... | ... | ... | ... |

**Pergunta de Reflexão:**
- Se você fosse criar um blog pessoal, qual tipo escolheria? Por quê?
- E se fosse criar um e-commerce grande? Por quê?

---

## 🤔 Perguntas de Reflexão

### Reflexão 1: HTTP e a Importância dos Códigos de Status

**Cenário:** Você está desenvolvendo um site e um usuário relata que não consegue acessar uma página específica.

**Situação A:**
O servidor retorna o código **404 Not Found**.

**Situação B:**
O servidor retorna o código **500 Internal Server Error**.

**Perguntas:**
1. Qual é a diferença entre esses dois códigos de status?
2. Como cada código ajuda você a identificar o problema?
3. O que você faria diferente para resolver cada situação?
4. Por que é importante que desenvolvedores entendam códigos de status HTTP?
5. Como você pode usar o DevTools para diagnosticar problemas de requisições HTTP?

**Resposta esperada:** Explique a diferença entre erros do cliente (4xx) e erros do servidor (5xx), e como isso impacta o desenvolvimento e a experiência do usuário.

---

### Reflexão 2: DNS e Performance Web

**Cenário:** Você criou um site e está preocupado com a velocidade de carregamento. Você descobre que o tempo de resolução DNS está demorando muito.

**Perguntas:**
1. Como o DNS impacta a velocidade de carregamento de um site?
2. O que é cache DNS e como ele ajuda na performance?
3. Se você mudar de provedor DNS (ex: usar Google DNS 8.8.8.8), isso pode melhorar a velocidade? Por quê?
4. Por que um site pode carregar mais rápido na segunda visita do que na primeira?
5. Como desenvolvedores podem otimizar o uso de DNS em seus sites? (Dica: pense em DNS prefetch, subdomínios, CDN)

**Resposta esperada:** Explique a relação entre DNS e performance, e estratégias para otimização.

---

### Reflexão 3: Domínios, Marca e Identidade Online

**Cenário:** Você está criando um site para sua empresa e precisa escolher um domínio.

**Opção A:** `minhaempresa123.com`  
**Opção B:** `minha-empresa.com`  
**Opção C:** `minhaempresa.com.br`

**Perguntas:**
1. Qual opção transmite mais profissionalismo? Por quê?
2. Como a escolha do TLD (.com vs .com.br) pode impactar seu público-alvo?
3. Por que é importante ter um domínio próprio ao invés de usar subdomínios gratuitos (ex: meusite.wordpress.com)?
4. Como um domínio pode afetar o SEO e a confiança dos usuários?
5. Quais fatores você consideraria ao escolher um domínio para um negócio?

**Resposta esperada:** Analise a importância de domínios para marca, SEO e credibilidade online.

---

### Reflexão 4: Hospedagem e Escalabilidade

**Cenário:** Você criou um blog que começou pequeno, mas agora está recebendo muito tráfego e o site está lento ou caindo.

**Perguntas:**
1. Por que um site pode ficar lento quando recebe mais visitantes?
2. Quais são os sinais de que você precisa mudar de tipo de hospedagem?
3. Como a escolha inicial de hospedagem pode limitar o crescimento futuro?
4. O que é escalabilidade e por que é importante pensar nisso desde o início?
5. Quais são as vantagens e desvantagens de começar com hospedagem compartilhada e depois migrar?

**Resposta esperada:** Discuta a relação entre hospedagem, tráfego e escalabilidade, e a importância de planejamento.

---

### Reflexão 5: Navegadores e Compatibilidade

**Cenário:** Você desenvolveu um site que funciona perfeitamente no Chrome, mas quando testa no Firefox, alguns elementos aparecem diferentes ou quebrados.

**Perguntas:**
1. Por que o mesmo código HTML/CSS pode renderizar diferente em navegadores diferentes?
2. Qual é a importância de testar em múltiplos navegadores?
3. Como você pode usar DevTools para identificar e resolver problemas de compatibilidade?
4. O que são "prefixos CSS" (como -webkit-, -moz-) e por que são necessários?
5. Como você pode garantir que seu site funcione bem em todos os navegadores modernos?

**Resposta esperada:** Explique a importância da compatibilidade entre navegadores e estratégias para garantir que sites funcionem em todos eles.

---

### Reflexão 6: SEO e Acessibilidade

**Cenário:** Você está otimizando seu site para SEO e descobre que muitas das práticas de SEO também melhoram a acessibilidade.

**Perguntas:**
1. Como o uso correto de headings (H1-H6) beneficia tanto SEO quanto acessibilidade?
2. Por que o atributo `alt` em imagens é importante para ambos SEO e acessibilidade?
3. Como a estrutura semântica HTML5 ajuda tanto mecanismos de busca quanto leitores de tela?
4. Quais outras práticas de SEO também melhoram a acessibilidade?
5. Por que é importante pensar em SEO e acessibilidade juntos desde o início do desenvolvimento?

**Resposta esperada:** Demonstre como SEO e acessibilidade estão interconectados e como boas práticas beneficiam ambos.

---

### Reflexão 7: O Fluxo Completo: Da URL à Página Renderizada

**Cenário:** Você digita `exemplo.com` no navegador e a página aparece. Mas o que aconteceu por trás dos panos?

**Perguntas:**
1. Descreva passo a passo o que acontece desde que você digita a URL até ver a página renderizada.
2. Em qual etapa ocorre a resolução DNS? Por que ela precisa acontecer antes da requisição HTTP?
3. O que acontece se o DNS falhar? E se o servidor HTTP não responder?
4. Como o navegador decide quais recursos adicionais solicitar (CSS, JS, imagens)?
5. Por que entender esse fluxo completo é importante para um desenvolvedor web?

**Resposta esperada:** Explique o fluxo completo de acesso a um site, conectando todos os conceitos aprendidos (DNS, HTTP, navegadores, etc.).

---

### Reflexão 8: HTTPS e Segurança na Web

**Cenário:** Você está criando um site que vai coletar informações de usuários (formulários, dados pessoais).

**Perguntas:**
1. Por que é essencial usar HTTPS ao invés de HTTP para sites que coletam dados?
2. O que acontece com os dados quando são transmitidos via HTTP? E via HTTPS?
3. Como os usuários podem identificar se um site é seguro (HTTPS)?
4. Qual é o impacto de não usar HTTPS no SEO e na confiança dos usuários?
5. Como você pode obter um certificado SSL/TLS para seu site? (Pesquise sobre Let's Encrypt)

**Resposta esperada:** Explique a importância de HTTPS para segurança, privacidade e SEO.

---

## 📋 Checklist de Aprendizado

Após completar os exercícios, verifique se você consegue:

- [ ] Explicar o que é HTTP e como funciona o modelo request/response
- [ ] Identificar diferentes métodos HTTP (GET, POST, PUT, DELETE)
- [ ] Reconhecer códigos de status HTTP comuns (200, 404, 500, etc.)
- [ ] Explicar o que é um domínio e sua estrutura
- [ ] Entender como DNS traduz domínios em IPs
- [ ] Diferenciar tipos de hospedagem web
- [ ] Explicar como navegadores processam e renderizam páginas
- [ ] Criar meta tags essenciais para SEO
- [ ] Entender a relação entre SEO e acessibilidade
- [ ] Descrever o fluxo completo de acesso a um site
- [ ] Usar DevTools para analisar requisições HTTP
- [ ] Reconhecer a importância de HTTPS para segurança

---

## 🎓 Dicas para Resolução

### Dica 1: DevTools é Seu Amigo
Aprenda a usar o DevTools do navegador - ele é uma das ferramentas mais importantes para desenvolvedores web. Pratique inspecionar requisições, analisar código e debugar problemas.

### Dica 2: Teste em Múltiplos Navegadores
Sempre teste seu site em diferentes navegadores (Chrome, Firefox, Safari, Edge) para garantir compatibilidade.

### Dica 3: Valide Seu HTML
Use o [W3C Validator](https://validator.w3.org/) regularmente para garantir que seu código HTML está correto e semântico.

### Dica 4: Monitore Performance
Use ferramentas como [PageSpeed Insights](https://pagespeed.web.dev/) ou [GTmetrix](https://gtmetrix.com/) para analisar a performance do seu site.

### Dica 5: Aprenda Sobre SEO Continuamente
SEO muda constantemente. Mantenha-se atualizado com as melhores práticas e algoritmos dos mecanismos de busca.

---

## 📝 Instruções para Entrega

1. Crie uma pasta chamada `exercicios-aula-2` dentro da pasta da aula
2. Salve cada exercício em arquivos separados:
   - `exercicio-1-analise-http.md` (anotações sobre requisições HTTP)
   - `exercicio-2-dns.md` (resultados da investigação DNS)
   - `exercicio-3-meta-tags.md` (tabela comparativa de meta tags)
   - `exercicio-4-pagina-seo.html` (página HTML com boas práticas)
   - `exercicio-5-hospedagem.md` (tabela comparativa de hospedagem)
3. Crie um arquivo `reflexoes.md` com suas respostas às perguntas de reflexão
4. Revise suas respostas antes de considerar concluído

**Boa sorte! Lembre-se: entender como a web funciona é fundamental para ser um bom desenvolvedor!** 🚀

---

## 🔗 Recursos Adicionais

### Ferramentas Úteis

- **W3C Validator**: [validator.w3.org](https://validator.w3.org/)
- **PageSpeed Insights**: [pagespeed.web.dev](https://pagespeed.web.dev/)
- **DNS Checker**: [dnschecker.org](https://dnschecker.org)
- **SSL Test**: [ssllabs.com/ssltest](https://www.ssllabs.com/ssltest/)

### Documentação

- **MDN Web Docs - HTTP**: [developer.mozilla.org/en-US/docs/Web/HTTP](https://developer.mozilla.org/en-US/docs/Web/HTTP)
- **Google Search Central**: [developers.google.com/search](https://developers.google.com/search)
- **Web.dev - Performance**: [web.dev/performance](https://web.dev/performance)

