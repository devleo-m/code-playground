# Aula 5: Agrupamento de Texto, Atributos e Listas

## 📚 Sobre Esta Aula

Esta é a quinta aula do curso de HTML, onde você aprenderá sobre agrupamento de elementos (`<div>` e `<span>`), atributos padrão (`id`, `class`, `data-*`, `style`), diferentes tipos de listas (ordenadas, não ordenadas, de definição e aninhadas) e como criar tabelas HTML bem estruturadas.

## 📖 Conteúdo da Aula

### 1. Aula Principal (`01-aula-principal.md`)
Conteúdo técnico completo sobre:
- Agrupamento de elementos com `<div>` (block-level) e `<span>` (inline)
- Diferenças entre div e span e quando usar cada um
- Atributos padrão: `id`, `class`, `data-*`, `style`
- Listas ordenadas (`<ol>`) e seus atributos
- Listas não ordenadas (`<ul>`) e seus usos
- Listas de definição (`<dl>`, `<dt>`, `<dd>`)
- Listas aninhadas e hierarquia
- Tabelas HTML: estrutura básica e semântica
- Elementos de tabela: `<table>`, `<tr>`, `<td>`, `<th>`
- Estrutura semântica: `<thead>`, `<tbody>`, `<tfoot>`, `<caption>`
- Atributos `colspan` e `rowspan` para mesclagem de células

### 2. Aula Simplificada (`02-aula-simplificada.md`)
Mesmo conteúdo explicado de forma mais acessível com:
- Analogias do dia a dia (div como caixa, span como marcador)
- Metáforas visuais (atributos como etiquetas)
- Exemplos práticos passo a passo
- Comparações com conceitos conhecidos (listas como receitas, tabelas como planilhas)

### 3. Exercícios e Reflexão (`03-exercicios-reflexao.md`)
Inclui:
- 8 exercícios práticos de HTML
- 6 perguntas de reflexão sobre semântica, acessibilidade, SEO e estruturação
- Checklist de aprendizado
- Instruções detalhadas para entrega

### 4. Performance e Boas Práticas (`04-performance-boas-praticas.md`)
Cobre:
- Quando usar `<div>` vs elementos semânticos
- Boas práticas de atributos (`id`, `class`, `data-*`, `style`)
- Escolha do tipo correto de lista
- Estruturação correta de listas e tabelas
- Acessibilidade em listas e tabelas
- Tabelas responsivas
- Impacto na performance
- Checklist completo de boas práticas

## 📁 Arquivos de Exemplo

### `exemplo-01-div-span.html`
Demonstra o uso de:
- `<div>` como container de seção
- Múltiplas divs para cards
- `<span>` para destacar texto inline
- Diferenças visuais entre div (block-level) e span (inline)
- Quando usar e quando não usar div/span

### `exemplo-02-atributos.html`
Exemplos práticos de:
- Atributo `id` para identificação única
- Atributo `class` para agrupamento
- Múltiplas classes em um elemento
- Atributos `data-*` para dados customizados
- Atributo `style` inline (uso limitado)
- Navegação usando IDs
- Comparação entre ID e Class

### `exemplo-03-listas.html`
Demonstra:
- Listas ordenadas (`<ol>`) com diferentes tipos de numeração
- Listas não ordenadas (`<ul>`) para menus e características
- Listas de definição (`<dl>`) para glossários e FAQs
- Listas aninhadas (hierarquia de informações)
- Listas mistas (ordenadas e não ordenadas)
- Listas com elementos complexos

### `exemplo-04-tabelas.html`
Exemplos de:
- Tabelas básicas com cabeçalhos e dados
- Estrutura semântica completa (thead, tbody, tfoot, caption)
- Mesclagem de células com `colspan` e `rowspan`
- Atributo `scope` para acessibilidade
- Tabelas de comparação
- Tabelas responsivas com scroll horizontal

## 🎯 Objetivos de Aprendizado

Ao final desta aula, você será capaz de:

- ✅ Entender a diferença entre `<div>` (block-level) e `<span>` (inline)
- ✅ Saber quando usar `<div>` vs elementos semânticos
- ✅ Usar atributos `id` e `class` corretamente
- ✅ Aplicar atributos `data-*` e `style` quando apropriado
- ✅ Criar listas ordenadas, não ordenadas e de definição
- ✅ Criar listas aninhadas mantendo estrutura correta
- ✅ Criar tabelas HTML bem estruturadas e semânticas
- ✅ Usar `colspan` e `rowspan` corretamente
- ✅ Entender quando usar cada tipo de elemento
- ✅ Considerar acessibilidade e semântica em suas escolhas
- ✅ Validar código HTML

## 🚀 Como Usar Esta Aula

1. **Leia a Aula Principal** (`01-aula-principal.md`)
   - Leia com atenção, fazendo anotações
   - Experimente os exemplos de código
   - Crie seus próprios arquivos HTML

2. **Revise com a Aula Simplificada** (`02-aula-simplificada.md`)
   - Use para reforçar conceitos que não ficaram claros
   - As analogias ajudam a fixar o aprendizado
   - Perfeito para revisão antes dos exercícios

3. **Abra os Exemplos HTML**
   - Abra cada arquivo de exemplo no navegador
   - Inspecione o código usando DevTools (F12)
   - Modifique os exemplos para experimentar
   - Compare como diferentes abordagens são renderizadas

4. **Complete os Exercícios** (`03-exercicios-reflexao.md`)
   - Faça todos os 8 exercícios práticos
   - Reflita sobre as 6 perguntas de reflexão
   - Use o checklist para verificar seu aprendizado
   - Valide seu código no W3C Validator

5. **Estude as Boas Práticas** (`04-performance-boas-praticas.md`)
   - Entenda por que cada prática é importante
   - Use como referência ao criar seus próprios projetos
   - Siga o checklist de boas práticas

## 🛠️ Ferramentas Recomendadas

- **Editor de Código**: Visual Studio Code (recomendado)
  - Extensões úteis: HTML CSS Support, Prettier, HTMLHint
- **Navegador**: Chrome, Firefox ou Edge (com DevTools)
- **Validador**: [W3C Markup Validator](https://validator.w3.org/)
- **Linter**: [HTMLHint](https://htmlhint.com/)

## 📝 Conceitos-Chave

### Agrupamento
- `<div>`: Container block-level para estruturação
- `<span>`: Container inline para texto e elementos inline
- Prefira elementos semânticos quando apropriado

### Atributos Padrão
- `id`: Identificador único (um por documento)
- `class`: Agrupamento por classes (múltiplos elementos)
- `data-*`: Dados customizados para JavaScript
- `style`: Estilização inline (uso limitado)

### Listas
- `<ol>`: Listas ordenadas (ordem importa - 1, 2, 3...)
- `<ul>`: Listas não ordenadas (ordem não importa - • • •)
- `<dl>`: Listas de definição (termo → definição)
- Listas aninhadas: Hierarquia de informações

### Tabelas
- `<table>`: Container da tabela
- `<tr>`: Linhas
- `<td>`: Células de dados
- `<th>`: Células de cabeçalho
- `<thead>`, `<tbody>`, `<tfoot>`: Estrutura semântica
- `<caption>`: Título da tabela
- `colspan` e `rowspan`: Mesclagem de células

## 💡 Dicas de Estudo

1. **Pratique regularmente**: Crie seus próprios arquivos HTML
2. **Experimente**: Modifique os exemplos fornecidos
3. **Valide seu código**: Use o W3C Validator sempre
4. **Use DevTools**: Aprenda a inspecionar elementos
5. **Pense em semântica**: Escolha elementos que tenham significado
6. **Teste acessibilidade**: Use leitores de tela ou extensões
7. **Considere responsividade**: Teste em diferentes tamanhos de tela
8. **Organize seu código**: Use indentação e comentários

## ❓ Dúvidas Comuns

### Quando devo usar div ao invés de elementos semânticos?
Use `<div>` apenas quando não há um elemento semântico apropriado. Prefira `<section>`, `<article>`, `<header>`, `<footer>`, `<nav>`, etc. quando fizer sentido semântico.

### Qual a diferença entre id e class?
`id` deve ser único (como um CPF), enquanto `class` pode ser repetida (como um uniforme). Use `id` para elementos únicos e `class` para agrupar elementos relacionados.

### Posso ter múltiplos elementos com o mesmo id?
Não! IDs devem ser únicos. Ter IDs duplicados quebra funcionalidade de CSS, JavaScript e links âncora.

### Quando devo usar listas ordenadas vs não ordenadas?
Use `<ol>` quando a ordem importa (receitas, instruções, rankings). Use `<ul>` quando a ordem não importa (características, menus, listas de compras).

### Posso usar tabelas para fazer layout?
Não! Tabelas devem ser usadas apenas para dados tabulares. Para layout, use CSS Grid ou Flexbox.

### Quantos níveis de aninhamento posso ter em listas?
Recomenda-se limitar a 2-3 níveis. Aninhamento excessivo dificulta navegação e acessibilidade.

## 📚 Próximos Passos

Após completar esta aula, você estará pronto para:
- Aula 6: Formulários HTML
- Aula 7: Imagens e Mídia
- Aula 8: Elementos Semânticos Avançados

## 🔗 Recursos Adicionais

- [MDN Web Docs - HTML Elements](https://developer.mozilla.org/pt-BR/docs/Web/HTML/Element)
- [MDN Web Docs - Lists](https://developer.mozilla.org/pt-BR/docs/Learn/HTML/Introduction_to_HTML/HTML_text_fundamentals#lists)
- [MDN Web Docs - Tables](https://developer.mozilla.org/pt-BR/docs/Learn/HTML/Tables)
- [W3C HTML Validator](https://validator.w3.org/)
- [HTML Semantic Elements](https://www.w3schools.com/html/html5_semantic_elements.asp)
- [Accessible Rich Internet Applications (ARIA)](https://www.w3.org/WAI/ARIA/apg/)

---

## 📋 Checklist de Conclusão

Antes de avançar para a próxima aula, certifique-se de que você:

- [ ] Leu a aula principal completa
- [ ] Revisou a aula simplificada
- [ ] Abriu e experimentou todos os exemplos HTML
- [ ] Completou todos os 8 exercícios práticos
- [ ] Refletiu sobre as 6 perguntas de reflexão
- [ ] Estudou as boas práticas
- [ ] Validou seu código no W3C Validator
- [ ] Entendeu todos os conceitos-chave
- [ ] Criou pelo menos um arquivo HTML próprio usando todos os conceitos aprendidos
- [ ] Compreende quando usar cada tipo de elemento
- [ ] Considera acessibilidade ao criar listas e tabelas

---

**Boa sorte na sua jornada de aprendizado HTML!** 🚀

Lembre-se: dominar agrupamento, atributos, listas e tabelas é essencial para escrever HTML profissional e de alta qualidade!

