# Aula 4 - Exercícios e Reflexão

## 🎯 Exercícios Práticos

### Exercício 1: Estrutura Básica Completa

Crie um arquivo HTML completo seguindo todas as boas práticas aprendidas:

**Requisitos:**
- Use `<!DOCTYPE html>` na primeira linha
- Tag `<html>` com atributo `lang="pt-BR"`
- Seção `<head>` com:
  - Meta tag `charset="UTF-8"` (primeira meta tag)
  - Meta tag `viewport` para dispositivos móveis
  - Meta tag `description` com descrição relevante (120-160 caracteres)
  - Tag `<title>` descritiva e única
- Seção `<body>` com:
  - Um `<h1>` como título principal
  - Pelo menos 3 parágrafos (`<p>`)
  - Uma hierarquia de títulos correta (h1 → h2 → h3, sem pular níveis)

**Validação:** Valide seu código no [W3C Validator](https://validator.w3.org/) e corrija todos os erros.

---

### Exercício 2: Hierarquia de Títulos Correta

Crie uma página sobre "Receitas de Culinária" demonstrando hierarquia correta de títulos:

**Requisitos:**
1. Use apenas **um** `<h1>` com o título "Receitas de Culinária"
2. Crie pelo menos 3 seções principais usando `<h2>`:
   - "Receitas Doces"
   - "Receitas Salgadas"
   - "Receitas Vegetarianas"
3. Dentro de cada seção `<h2>`, adicione pelo menos 2 receitas usando `<h3>`
4. Dentro de cada receita (`<h3>`), adicione subtítulos usando `<h4>`:
   - "Ingredientes"
   - "Modo de Preparo"
   - "Dicas"

**Desafio:** Identifique o que aconteceria se você pulasse níveis (ex: h1 → h3). Por que isso é problemático?

---

### Exercício 3: Formatação de Texto Semântica

Crie uma página que demonstre o uso correto de todas as tags de formatação aprendidas:

**Requisitos:**
1. Crie um parágrafo usando `<strong>` para destacar importância
2. Crie um parágrafo usando `<em>` para dar ênfase
3. Use `<mark>` para destacar texto relevante
4. Crie uma seção sobre fórmulas químicas usando `<sub>`:
   - H₂O (água)
   - CO₂ (dióxido de carbono)
   - C₆H₁₂O₆ (glicose)
5. Crie uma seção sobre matemática usando `<sup>`:
   - 2³ = 8
   - x² + y² = z²
   - 1º lugar, 2º lugar, 3º lugar
6. Use `<pre>` para mostrar um trecho de código JavaScript
7. Demonstre a diferença entre `<b>` e `<strong>`, `<i>` e `<em>` com exemplos comentados

**Tarefa adicional:** Explique em comentários HTML quando usar cada tag e por quê.

---

### Exercício 4: Links e Navegação

Crie uma página com diferentes tipos de links:

**Requisitos:**
1. Crie um menu de navegação usando links internos (âncoras) para diferentes seções da página
2. Adicione pelo menos 3 links externos com:
   - `target="_blank"`
   - `rel="noopener noreferrer"`
   - `title` descritivo
3. Crie um link de email com `mailto:` incluindo assunto e corpo
4. Crie um link de telefone com `tel:`
5. Crie uma seção "Voltar ao topo" usando link âncora
6. Adicione um link de download (use um arquivo fictício)

**Desafio:** Crie uma página longa (com scroll) e implemente um menu fixo no topo que permite navegar entre as seções.

---

### Exercício 5: Estruturação Completa de Conteúdo

Crie uma página completa sobre um tema de sua escolha (ex: seu hobby, um lugar que visitou, um livro favorito) usando TODAS as tags aprendidas:

**Requisitos:**
- Estrutura básica completa (DOCTYPE, html, head, body)
- Meta tags essenciais
- Título apropriado
- Hierarquia de títulos correta (h1, h2, h3)
- Parágrafos bem estruturados
- Uso de `<br>` quando apropriado (ex: endereço)
- Uso de `<hr>` para separar seções
- Formatação de texto variada (strong, em, mark, sub, sup)
- Links internos e externos
- Texto pré-formatado (`<pre>`) se relevante

**Validação:** Valide no W3C Validator e corrija todos os erros.

---

### Exercício 6: Análise e Correção de Código

Analise o seguinte código HTML e identifique TODOS os problemas:

```html
<!DOCTYPE HTML>
<HTML LANG="pt-br">
<HEAD>
    <meta charset="utf-8">
    <title>Minha Página</title>
</HEAD>
<BODY>
    <h1>Título Principal</h1>
    <h1>Outro Título Principal</h1>
    <h3>Subtítulo</h3>
    <p>Este é um parágrafo.
    Continua aqui.</p>
    <p>Outro parágrafo<br><br><br>com muitas quebras</p>
    <strong><b>Texto importante</b></strong>
    <i><em>Texto com ênfase</em></i>
    <a href="https://www.exemplo.com" target="_blank">Link</a>
    <p>Fórmula: H2O</p>
    <p>Potência: 2^3</p>
</BODY>
</HTML>
```

**Tarefas:**
1. Liste todos os problemas encontrados
2. Explique por que cada problema é um erro
3. Reescreva o código corrigindo todos os problemas
4. Valide o código corrigido no W3C Validator

---

## 🤔 Perguntas de Reflexão

### Reflexão 1: Semântica vs. Visual

**Pergunta:** Qual é a diferença fundamental entre usar `<strong>` e `<b>`, ou `<em>` e `<i>`? Por que essa diferença é importante para:
- Acessibilidade (leitores de tela)?
- SEO (mecanismos de busca)?
- Manutenção do código?

**Instruções:** Escreva um parágrafo explicando sua compreensão e dê exemplos práticos de quando cada tag seria mais apropriada.

---

### Reflexão 2: Hierarquia de Títulos e Estruturação

**Pergunta:** Por que é importante manter a hierarquia correta de títulos (não pular de h1 para h3)? Qual o impacto de:
- Pular níveis de títulos?
- Usar múltiplos `<h1>` na mesma página?
- Usar títulos apenas para tamanho visual (sem considerar semântica)?

**Instruções:** Crie um exemplo visual mostrando uma estrutura correta vs. incorreta e explique as consequências de cada abordagem.

---

### Reflexão 3: Meta Tags e SEO

**Pergunta:** 
1. Por que a meta tag `charset="UTF-8"` deve ser a primeira meta tag no `<head>`?
2. Qual a importância da meta tag `description` para SEO?
3. Por que a meta tag `keywords` não é mais recomendada?
4. Como a tag `<title>` afeta a experiência do usuário e SEO?

**Instruções:** Pesquise e explique como cada meta tag afeta a indexação e exibição nos resultados de busca. Dê exemplos de boas e más práticas.

---

### Reflexão 4: Acessibilidade em Links

**Pergunta:**
1. Por que é importante usar `rel="noopener noreferrer"` com `target="_blank"`?
2. Qual o impacto de links sem texto descritivo (ex: "clique aqui") para acessibilidade?
3. Como leitores de tela interpretam links? O que acontece quando um link só tem "clique aqui"?
4. Qual a melhor prática para links que abrem em nova aba?

**Instruções:** 
- Teste uma página com leitor de tela (ou extensão de navegador que simula)
- Compare links descritivos vs. "clique aqui"
- Explique por que links descritivos são essenciais

---

### Reflexão 5: Estruturação Semântica do Conteúdo

**Pergunta:** 
1. Por que usar `<p>` para parágrafos ao invés de apenas `<br>` repetidas vezes?
2. Quando é apropriado usar `<br>` vs. criar um novo parágrafo?
3. Qual o impacto de usar `<hr>` para separação visual vs. separação semântica?
4. Como a estruturação semântica afeta leitores de tela e mecanismos de busca?

**Instruções:** Crie dois exemplos: um com estruturação semântica correta e outro semânticamente incorreto. Explique as diferenças e impactos.

---

### Reflexão 6: Performance e Boas Práticas

**Pergunta:**
1. Qual o impacto de ter múltiplas meta tags desnecessárias no `<head>`?
2. Por que a ordem das meta tags importa (charset primeiro)?
3. Como tags mal estruturadas podem afetar o tempo de renderização?
4. Qual a importância de validar HTML no W3C Validator?

**Instruções:** 
- Valide uma página no W3C Validator
- Analise os erros e avisos
- Explique como cada erro pode afetar performance, acessibilidade ou SEO
- Crie um checklist pessoal de validação antes de publicar código

---

## 📋 Checklist de Aprendizado

Antes de considerar esta aula completa, verifique se você:

### Conhecimento Técnico
- [ ] Entendo a estrutura completa de um documento HTML
- [ ] Sei quando e como usar cada meta tag essencial
- [ ] Compreendo a hierarquia de títulos e por que não devo pular níveis
- [ ] Sei a diferença entre tags semânticas e visuais
- [ ] Entendo quando usar cada tag de formatação de texto
- [ ] Sei criar links seguros e acessíveis
- [ ] Compreendo a diferença entre `<br>` e novos parágrafos

### Aplicação Prática
- [ ] Criei pelo menos 3 arquivos HTML completos
- [ ] Validei meu código no W3C Validator
- [ ] Testei meus links em diferentes navegadores
- [ ] Verifiquei a hierarquia de títulos em minhas páginas
- [ ] Usei tags semânticas apropriadas

### Reflexão e Compreensão
- [ ] Entendo por que semântica é importante
- [ ] Compreendo o impacto na acessibilidade
- [ ] Sei como minha estruturação afeta SEO
- [ ] Entendo boas práticas de segurança em links
- [ ] Posso explicar a diferença entre tags semânticas e visuais

---

## 📤 Instruções para Entrega

### O que entregar:

1. **Arquivos HTML criados:**
   - Exercício 1: `exercicio-01-estrutura.html`
   - Exercício 2: `exercicio-02-hierarquia.html`
   - Exercício 3: `exercicio-03-formatacao.html`
   - Exercício 4: `exercicio-04-links.html`
   - Exercício 5: `exercicio-05-conteudo-completo.html`
   - Exercício 6: `exercicio-06-codigo-corrigido.html`

2. **Documento de reflexão:**
   - Crie um arquivo `reflexoes.md` com suas respostas às 6 perguntas de reflexão
   - Seja detalhado e use exemplos práticos
   - Inclua screenshots ou exemplos de código quando relevante

3. **Validação:**
   - Screenshots ou links dos resultados do W3C Validator para cada arquivo
   - Todos os arquivos devem estar sem erros (avisos são aceitáveis)

### Critérios de Avaliação:

- **Correção técnica:** Código válido e sem erros
- **Semântica:** Uso apropriado de tags semânticas
- **Estruturação:** Hierarquia correta e organização lógica
- **Acessibilidade:** Links descritivos, estrutura acessível
- **Reflexão:** Respostas bem fundamentadas e com exemplos

---

## 💡 Dicas para os Exercícios

1. **Sempre valide:** Use o W3C Validator após cada exercício
2. **Teste no navegador:** Veja como seu código é renderizado
3. **Use DevTools:** Inspecione elementos para entender a estrutura
4. **Pense em semântica:** Sempre escolha a tag mais apropriada
5. **Seja consistente:** Mantenha padrões de formatação e nomenclatura
6. **Documente:** Adicione comentários quando necessário

---

## 🚀 Próximos Passos

Após completar todos os exercícios e reflexões:

1. Revise suas respostas e compare com as boas práticas
2. Estude o arquivo `04-performance-boas-praticas.md`
3. Crie um projeto pessoal aplicando tudo que aprendeu
4. Prepare-se para a próxima aula sobre listas e estruturas

---

**Boa sorte com os exercícios!** 🎯

Lembre-se: a prática é essencial. Quanto mais você codificar, mais natural se tornará o uso correto das tags HTML!


