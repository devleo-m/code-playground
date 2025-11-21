# Aula 6 - Exercícios e Reflexão: SQL JOIN Queries

## Exercícios Práticos

### Exercício 1: INNER JOIN Básico

**Objetivo**: Praticar INNER JOIN para combinar dados de duas tabelas relacionadas.

**Tarefas**:

1. Escreva uma query usando INNER JOIN para listar todos os livros com seus respectivos autores. Mostre:
   - Título do livro
   - Nome do autor
   - Ano de publicação

2. Escreva uma query usando INNER JOIN para listar todos os empréstimos ativos com:
   - Nome do usuário
   - Título do livro
   - Data do empréstimo
   - Data de devolução prevista

3. Escreva uma query usando múltiplos INNER JOINs para listar livros com:
   - Título do livro
   - Nome do autor
   - Nome da categoria
   - Quantidade disponível

**Questão de Reflexão**:
- Por que INNER JOIN é o tipo de JOIN mais comum? Em que situações você usaria INNER JOIN ao invés de LEFT JOIN?

**Soluções Esperadas**:

```sql
-- 1. Livros com autores
SELECT 
    l.titulo,
    a.nome AS autor,
    l.ano_publicacao
FROM livros l
INNER JOIN autores a ON l.autor_id = a.id
ORDER BY l.titulo;

-- 2. Empréstimos ativos com informações
SELECT 
    u.nome AS usuario,
    l.titulo AS livro,
    e.data_emprestimo,
    e.data_devolucao_prevista
FROM emprestimos e
INNER JOIN usuarios u ON e.usuario_id = u.id
INNER JOIN livros l ON e.livro_id = l.id
WHERE e.status = 'ativo'
ORDER BY e.data_emprestimo DESC;

-- 3. Livros com autor e categoria
SELECT 
    l.titulo,
    a.nome AS autor,
    c.nome AS categoria,
    l.quantidade_disponivel
FROM livros l
INNER JOIN autores a ON l.autor_id = a.id
INNER JOIN categorias c ON l.categoria_id = c.id
ORDER BY c.nome, l.titulo;
```

**Resposta Esperada para a Questão de Reflexão**:
- **INNER JOIN é mais comum** porque:
  - A maioria das queries precisa apenas de dados relacionados
  - É mais eficiente (retorna menos dados)
  - É mais intuitivo (apenas o que combina)
  - É o comportamento padrão esperado
- **Use INNER JOIN quando**:
  - Você precisa apenas de registros com correspondência
  - Você quer excluir registros órfãos
  - Os relacionamentos são obrigatórios
- **Use LEFT JOIN quando**:
  - Você precisa incluir todos os registros de uma tabela
  - Você quer encontrar registros sem correspondência
  - Os relacionamentos são opcionais

---

### Exercício 2: LEFT JOIN - Incluindo Todos os Registros

**Objetivo**: Praticar LEFT JOIN para incluir todos os registros de uma tabela, mesmo sem correspondência.

**Tarefas**:

1. Escreva uma query usando LEFT JOIN para listar todas as categorias e quantos livros cada uma tem (incluindo categorias sem livros).

2. Escreva uma query usando LEFT JOIN para encontrar categorias que não têm nenhum livro cadastrado.

3. Escreva uma query usando LEFT JOIN para listar todos os autores e quantos livros cada um escreveu (incluindo autores sem livros).

4. Escreva uma query usando LEFT JOIN para encontrar autores que não têm nenhum livro cadastrado.

**Questão de Reflexão**:
- Qual a diferença entre usar `COUNT(*)` e `COUNT(coluna)` em um LEFT JOIN? Por que isso importa?

**Soluções Esperadas**:

```sql
-- 1. Todas as categorias com contagem de livros
SELECT 
    c.nome AS categoria,
    COUNT(l.id) AS total_livros
FROM categorias c
LEFT JOIN livros l ON c.id = l.categoria_id
GROUP BY c.id, c.nome
ORDER BY total_livros DESC;

-- 2. Categorias sem livros
SELECT 
    c.id,
    c.nome AS categoria
FROM categorias c
LEFT JOIN livros l ON c.id = l.categoria_id
WHERE l.id IS NULL
ORDER BY c.nome;

-- 3. Todos os autores com contagem de livros
SELECT 
    a.nome AS autor,
    a.nacionalidade,
    COUNT(l.id) AS total_livros
FROM autores a
LEFT JOIN livros l ON a.id = l.autor_id
GROUP BY a.id, a.nome, a.nacionalidade
ORDER BY total_livros DESC, a.nome;

-- 4. Autores sem livros
SELECT 
    a.id,
    a.nome AS autor,
    a.nacionalidade
FROM autores a
LEFT JOIN livros l ON a.id = l.autor_id
WHERE l.id IS NULL
ORDER BY a.nome;
```

**Resposta Esperada para a Questão de Reflexão**:
- **COUNT(*) em LEFT JOIN**:
  - Conta todas as linhas, incluindo NULLs
  - Para categorias sem livros, retorna 1 (não 0!)
  - Pode dar resultados incorretos
- **COUNT(coluna) em LEFT JOIN**:
  - Conta apenas valores não-NULL
  - Para categorias sem livros, retorna 0 (correto!)
  - É o comportamento desejado na maioria dos casos
- **Recomendação**: Sempre use `COUNT(coluna_da_tabela_direita)` em LEFT JOIN para contar apenas correspondências reais

---

### Exercício 3: Comparando INNER JOIN vs LEFT JOIN

**Objetivo**: Entender a diferença prática entre INNER JOIN e LEFT JOIN.

**Tarefas**:

1. Execute a mesma query usando INNER JOIN e depois LEFT JOIN para listar categorias com seus livros. Compare os resultados.

2. Conte quantas categorias aparecem em cada resultado.

3. Identifique quais categorias aparecem apenas no LEFT JOIN.

**Questão de Reflexão**:
- Quando você usaria INNER JOIN e quando usaria LEFT JOIN? Dê exemplos práticos de situações reais.

**Soluções Esperadas**:

```sql
-- INNER JOIN
SELECT 
    c.nome AS categoria,
    l.titulo AS livro
FROM categorias c
INNER JOIN livros l ON c.id = l.categoria_id
ORDER BY c.nome, l.titulo;

-- LEFT JOIN
SELECT 
    c.nome AS categoria,
    l.titulo AS livro
FROM categorias c
LEFT JOIN livros l ON c.id = l.categoria_id
ORDER BY c.nome, l.titulo;

-- Comparação: Contar categorias em cada resultado
-- INNER JOIN: apenas categorias com livros
SELECT COUNT(DISTINCT c.id) AS total_categorias
FROM categorias c
INNER JOIN livros l ON c.id = l.categoria_id;

-- LEFT JOIN: todas as categorias
SELECT COUNT(DISTINCT c.id) AS total_categorias
FROM categorias c
LEFT JOIN livros l ON c.id = l.categoria_id;
```

**Resposta Esperada para a Questão de Reflexão**:
- **Use INNER JOIN quando**:
  - "Mostre-me os pedidos com seus clientes" (apenas pedidos com cliente)
  - "Liste produtos que foram vendidos" (apenas produtos com vendas)
  - "Mostre funcionários com seus departamentos" (apenas funcionários alocados)
- **Use LEFT JOIN quando**:
  - "Mostre-me todos os produtos, mesmo os que nunca foram vendidos" (todos os produtos)
  - "Liste todos os clientes e seus pedidos" (incluindo clientes sem pedidos)
  - "Encontre categorias sem produtos" (precisa ver todas as categorias)
- **Regra geral**: INNER JOIN para dados obrigatórios, LEFT JOIN para dados opcionais ou quando você precisa ver tudo

---

### Exercício 4: SELF JOIN

**Objetivo**: Praticar SELF JOIN para comparar registros dentro da mesma tabela.

**Tarefas**:

1. Escreva uma query usando SELF JOIN para encontrar pares de autores da mesma nacionalidade (evite duplicatas e auto-comparação).

2. Escreva uma query usando SELF JOIN para encontrar livros do mesmo autor publicados em anos diferentes (mostre o livro mais antigo e o mais novo).

3. Escreva uma query usando SELF JOIN para encontrar livros da mesma categoria publicados no mesmo ano.

**Questão de Reflexão**:
- Por que aliases são obrigatórios em SELF JOIN? O que aconteceria se você não usasse aliases?

**Soluções Esperadas**:

```sql
-- 1. Autores da mesma nacionalidade
SELECT 
    a1.nome AS autor1,
    a2.nome AS autor2,
    a1.nacionalidade
FROM autores a1
INNER JOIN autores a2 ON a1.nacionalidade = a2.nacionalidade
WHERE a1.id < a2.id  -- Evita duplicatas e auto-comparação
ORDER BY a1.nacionalidade, a1.nome;

-- 2. Livros do mesmo autor em anos diferentes
SELECT 
    a.nome AS autor,
    l1.titulo AS livro_antigo,
    l1.ano_publicacao AS ano_antigo,
    l2.titulo AS livro_novo,
    l2.ano_publicacao AS ano_novo
FROM livros l1
INNER JOIN livros l2 ON l1.autor_id = l2.autor_id
INNER JOIN autores a ON l1.autor_id = a.id
WHERE l1.ano_publicacao < l2.ano_publicacao
ORDER BY a.nome, l1.ano_publicacao;

-- 3. Livros da mesma categoria no mesmo ano
SELECT 
    c.nome AS categoria,
    l1.titulo AS livro1,
    l2.titulo AS livro2,
    l1.ano_publicacao AS ano
FROM livros l1
INNER JOIN livros l2 ON l1.categoria_id = l2.categoria_id 
    AND l1.ano_publicacao = l2.ano_publicacao
INNER JOIN categorias c ON l1.categoria_id = c.id
WHERE l1.id < l2.id
ORDER BY c.nome, l1.ano_publicacao;
```

**Resposta Esperada para a Questão de Reflexão**:
- **Aliases são obrigatórios** porque:
  - A mesma tabela aparece duas vezes na query
  - Sem aliases, o SQL não saberia qual instância usar
  - É como ter dois irmãos gêmeos: você precisa dar nomes diferentes
- **Sem aliases**:
  - A query não funcionaria (erro de sintaxe)
  - Não haveria como distinguir as duas instâncias da tabela
  - As referências de colunas seriam ambíguas
- **Com aliases**:
  - Você pode referenciar `a1.nome` e `a2.nome` claramente
  - A query fica legível e sem ambiguidade
  - É possível comparar registros diferentes da mesma tabela

---

### Exercício 5: Múltiplos JOINs

**Objetivo**: Praticar combinação de múltiplos JOINs em uma única query.

**Tarefas**:

1. Escreva uma query usando múltiplos JOINs para criar um relatório completo de empréstimos mostrando:
   - ID do empréstimo
   - Nome do usuário
   - Email do usuário
   - Título do livro
   - Nome do autor
   - Nome da categoria
   - Data do empréstimo
   - Status

2. Escreva uma query usando múltiplos JOINs (incluindo LEFT JOIN) para listar todas as categorias com:
   - Nome da categoria
   - Total de livros
   - Total de autores únicos
   - Total de empréstimos de livros dessa categoria

3. Escreva uma query usando múltiplos JOINs para encontrar usuários que pegaram livros de autores brasileiros.

**Questão de Reflexão**:
- Como a ordem dos JOINs afeta a performance e a legibilidade da query? Existe uma ordem "correta"?

**Soluções Esperadas**:

```sql
-- 1. Relatório completo de empréstimos
SELECT 
    e.id AS emprestimo_id,
    u.nome AS usuario,
    u.email,
    l.titulo AS livro,
    a.nome AS autor,
    c.nome AS categoria,
    e.data_emprestimo,
    e.status
FROM emprestimos e
INNER JOIN usuarios u ON e.usuario_id = u.id
INNER JOIN livros l ON e.livro_id = l.id
INNER JOIN autores a ON l.autor_id = a.id
INNER JOIN categorias c ON l.categoria_id = c.id
ORDER BY e.data_emprestimo DESC;

-- 2. Estatísticas por categoria
SELECT 
    c.nome AS categoria,
    COUNT(DISTINCT l.id) AS total_livros,
    COUNT(DISTINCT l.autor_id) AS total_autores,
    COUNT(e.id) AS total_emprestimos
FROM categorias c
LEFT JOIN livros l ON c.id = l.categoria_id
LEFT JOIN emprestimos e ON l.id = e.livro_id
GROUP BY c.id, c.nome
ORDER BY total_livros DESC;

-- 3. Usuários com livros de autores brasileiros
SELECT DISTINCT
    u.nome AS usuario,
    u.email,
    a.nome AS autor_brasileiro
FROM usuarios u
INNER JOIN emprestimos e ON u.id = e.usuario_id
INNER JOIN livros l ON e.livro_id = l.id
INNER JOIN autores a ON l.autor_id = a.id
WHERE a.nacionalidade LIKE '%Brasileiro%'
ORDER BY u.nome;
```

**Resposta Esperada para a Questão de Reflexão**:
- **Ordem dos JOINs pode afetar**:
  - **Performance**: O otimizador pode escolher diferentes planos de execução
  - **Legibilidade**: Começar com a tabela "principal" é mais intuitivo
  - **Resultados**: Em alguns casos raros, a ordem pode importar
- **Ordem recomendada**:
  - Comece com a tabela "principal" (a que você quer no resultado)
  - Adicione JOINs em ordem lógica de relacionamento
  - Use LEFT JOIN quando precisar incluir todos os registros
  - Agrupe JOINs relacionados logicamente
- **Exemplo de ordem lógica**:
  - `emprestimos` → `usuarios` → `livros` → `autores` → `categorias`
  - Segue o fluxo natural: empréstimo tem usuário e livro, livro tem autor e categoria

---

### Exercício 6: Encontrar Registros Órfãos

**Objetivo**: Usar LEFT JOIN para encontrar registros sem correspondência.

**Tarefas**:

1. Escreva uma query para encontrar usuários que nunca pegaram livros emprestados.

2. Escreva uma query para encontrar livros que nunca foram emprestados.

3. Escreva uma query para encontrar categorias que não têm livros cadastrados.

4. Escreva uma query para encontrar autores que não têm livros cadastrados.

**Questão de Reflexão**:
- Por que é útil encontrar registros órfãos? Dê exemplos de situações onde isso seria importante em um sistema real.

**Soluções Esperadas**:

```sql
-- 1. Usuários sem empréstimos
SELECT 
    u.id,
    u.nome,
    u.email,
    u.data_cadastro
FROM usuarios u
LEFT JOIN emprestimos e ON u.id = e.usuario_id
WHERE e.id IS NULL
ORDER BY u.data_cadastro DESC;

-- 2. Livros nunca emprestados
SELECT 
    l.id,
    l.titulo,
    a.nome AS autor,
    l.quantidade_disponivel
FROM livros l
LEFT JOIN emprestimos e ON l.id = e.livro_id
WHERE e.id IS NULL
ORDER BY l.titulo;

-- 3. Categorias sem livros
SELECT 
    c.id,
    c.nome AS categoria,
    c.descricao
FROM categorias c
LEFT JOIN livros l ON c.id = l.categoria_id
WHERE l.id IS NULL
ORDER BY c.nome;

-- 4. Autores sem livros
SELECT 
    a.id,
    a.nome AS autor,
    a.nacionalidade
FROM autores a
LEFT JOIN livros l ON a.id = l.autor_id
WHERE l.id IS NULL
ORDER BY a.nome;
```

**Resposta Esperada para a Questão de Reflexão**:
- **É útil encontrar registros órfãos para**:
  - **Limpeza de dados**: Identificar dados não utilizados
  - **Análise de negócio**: Entender por que alguns registros não são usados
  - **Manutenção**: Remover dados obsoletos ou incorretos
  - **Relatórios**: Mostrar o que está faltando ou não está sendo usado
- **Exemplos práticos**:
  - **E-commerce**: Produtos que nunca foram vendidos → análise de catálogo
  - **Sistema de pedidos**: Clientes que nunca fizeram pedidos → campanhas de marketing
  - **Biblioteca**: Livros nunca emprestados → decisão sobre manter ou remover
  - **RH**: Funcionários sem departamento → identificar problemas de alocação

---

### Exercício 7: Condições ON vs WHERE

**Objetivo**: Entender a diferença entre colocar condições no ON e no WHERE.

**Tarefas**:

1. Escreva uma query usando LEFT JOIN com condição no ON para listar todas as categorias e apenas livros publicados após 2000.

2. Escreva a mesma query, mas com a condição no WHERE ao invés do ON.

3. Compare os resultados e explique a diferença.

**Questão de Reflexão**:
- Quando você colocaria uma condição no ON e quando no WHERE? Qual é a regra geral?

**Soluções Esperadas**:

```sql
-- 1. Condição no ON
SELECT 
    c.nome AS categoria,
    l.titulo AS livro,
    l.ano_publicacao
FROM categorias c
LEFT JOIN livros l ON c.id = l.categoria_id 
    AND l.ano_publicacao > 2000
ORDER BY c.nome, l.titulo;

-- 2. Condição no WHERE
SELECT 
    c.nome AS categoria,
    l.titulo AS livro,
    l.ano_publicacao
FROM categorias c
LEFT JOIN livros l ON c.id = l.categoria_id
WHERE l.ano_publicacao > 2000
ORDER BY c.nome, l.titulo;

-- 3. Comparação: Contar categorias em cada resultado
-- Com condição no ON: todas as categorias aparecem
SELECT COUNT(DISTINCT c.id) AS total_categorias
FROM categorias c
LEFT JOIN livros l ON c.id = l.categoria_id 
    AND l.ano_publicacao > 2000;

-- Com condição no WHERE: apenas categorias com livros após 2000
SELECT COUNT(DISTINCT c.id) AS total_categorias
FROM categorias c
LEFT JOIN livros l ON c.id = l.categoria_id
WHERE l.ano_publicacao > 2000;
```

**Resposta Esperada para a Questão de Reflexão**:
- **Condição no ON**:
  - Aplica o filtro **durante** o JOIN
  - Para LEFT JOIN, inclui todas as linhas da esquerda, mas filtra a direita
  - Útil quando você quer incluir todos os registros da esquerda, mas filtrar a direita
  - Exemplo: "Todas as categorias, mas apenas livros após 2000"
- **Condição no WHERE**:
  - Aplica o filtro **depois** do JOIN
  - Para LEFT JOIN, pode excluir linhas que você queria incluir
  - Útil quando você quer filtrar o resultado final
  - Exemplo: "Apenas categorias que têm livros após 2000"
- **Regra geral**:
  - **Condições de relacionamento** → `ON` (como as tabelas se relacionam)
  - **Filtros de resultado** → `WHERE` (o que você quer no resultado final)
  - **Para LEFT JOIN**: Use `ON` quando quiser incluir todos da esquerda, use `WHERE` quando quiser filtrar o resultado

---

### Exercício 8: Queries Complexas com JOINs

**Objetivo**: Combinar JOINs com funções de agregação e outras operações.

**Tarefas**:

1. Escreva uma query para encontrar os 5 livros mais emprestados, mostrando:
   - Título do livro
   - Nome do autor
   - Número de vezes emprestado

2. Escreva uma query para encontrar autores com mais de 2 livros cadastrados, mostrando:
   - Nome do autor
   - Nacionalidade
   - Total de livros

3. Escreva uma query para encontrar categorias com a maior quantidade total de livros disponíveis em estoque.

4. Escreva uma query para encontrar usuários que pegaram livros de mais de uma categoria diferente.

**Questão de Reflexão**:
- Como JOINs se combinam com GROUP BY e HAVING? Qual a ordem lógica de execução?

**Soluções Esperadas**:

```sql
-- 1. Top 5 livros mais emprestados
SELECT 
    l.titulo,
    a.nome AS autor,
    COUNT(e.id) AS vezes_emprestado
FROM livros l
INNER JOIN autores a ON l.autor_id = a.id
LEFT JOIN emprestimos e ON l.id = e.livro_id
GROUP BY l.id, l.titulo, a.nome
ORDER BY vezes_emprestado DESC
LIMIT 5;

-- 2. Autores com mais de 2 livros
SELECT 
    a.nome AS autor,
    a.nacionalidade,
    COUNT(l.id) AS total_livros
FROM autores a
INNER JOIN livros l ON a.id = l.autor_id
GROUP BY a.id, a.nome, a.nacionalidade
HAVING COUNT(l.id) > 2
ORDER BY total_livros DESC;

-- 3. Categoria com maior estoque
SELECT 
    c.nome AS categoria,
    SUM(l.quantidade_disponivel) AS total_estoque
FROM categorias c
LEFT JOIN livros l ON c.id = l.categoria_id
GROUP BY c.id, c.nome
ORDER BY total_estoque DESC
LIMIT 1;

-- 4. Usuários com livros de múltiplas categorias
SELECT 
    u.nome AS usuario,
    u.email,
    COUNT(DISTINCT l.categoria_id) AS categorias_diferentes
FROM usuarios u
INNER JOIN emprestimos e ON u.id = e.usuario_id
INNER JOIN livros l ON e.livro_id = l.id
GROUP BY u.id, u.nome, u.email
HAVING COUNT(DISTINCT l.categoria_id) > 1
ORDER BY categorias_diferentes DESC;
```

**Resposta Esperada para a Questão de Reflexão**:
- **Ordem lógica de execução**:
  1. **FROM**: Seleciona as tabelas
  2. **JOIN**: Combina as tabelas
  3. **WHERE**: Filtra as linhas
  4. **GROUP BY**: Agrupa as linhas
  5. **HAVING**: Filtra os grupos
  6. **SELECT**: Seleciona as colunas
  7. **ORDER BY**: Ordena os resultados
- **JOINs com GROUP BY**:
  - JOINs combinam dados de múltiplas tabelas
  - GROUP BY agrupa os resultados combinados
  - HAVING filtra os grupos resultantes
- **Exemplo prático**:
  - JOIN combina livros com autores
  - GROUP BY agrupa por autor
  - COUNT conta livros por autor
  - HAVING filtra autores com mais de 2 livros

---

### Exercício 9: CROSS JOIN (Cuidado!)

**Objetivo**: Entender CROSS JOIN e quando evitá-lo.

**Tarefas**:

1. Execute um CROSS JOIN entre `categorias` e `autores`. Quantas linhas foram retornadas?

2. Compare com um INNER JOIN entre as mesmas tabelas usando uma condição apropriada (se houver relacionamento) ou explique por que não faz sentido.

3. Escreva uma query que acidentalmente esquece a condição ON (resultando em CROSS JOIN) e depois corrija.

**Questão de Reflexão**:
- Por que CROSS JOIN geralmente é um erro? Quando seria apropriado usar CROSS JOIN?

**Soluções Esperadas**:

```sql
-- 1. CROSS JOIN (cuidado!)
SELECT 
    c.nome AS categoria,
    a.nome AS autor
FROM categorias c
CROSS JOIN autores a;
-- Resultado: 6 categorias × 10 autores = 60 linhas!

-- 2. Comparação: Não faz sentido fazer INNER JOIN direto
-- porque não há relacionamento direto entre categorias e autores
-- (eles se relacionam através de livros)

-- 3. Erro comum: esquecer ON
-- ❌ ERRADO (CROSS JOIN acidental)
SELECT l.titulo, a.nome
FROM livros l
JOIN autores a;  -- Faltou ON!

-- ✅ CORRETO
SELECT l.titulo, a.nome
FROM livros l
JOIN autores a ON l.autor_id = a.id;
```

**Resposta Esperada para a Questão de Reflexão**:
- **CROSS JOIN geralmente é um erro** porque:
  - Gera muitas linhas (produto cartesiano)
  - Na maioria dos casos, você esqueceu a condição `ON`
  - Resultados geralmente não fazem sentido
  - Pode ser muito custoso computacionalmente
- **CROSS JOIN seria apropriado quando**:
  - Você realmente precisa de todas as combinações
  - Gerar dados de teste com todas as combinações
  - Tabelas de referência que precisam ser combinadas
  - Casos muito específicos de análise combinatória
- **Recomendação**: Sempre verifique se você realmente precisa de CROSS JOIN. Na maioria dos casos, é um erro!

---

### Exercício 10: Desafio Final - Query Completa

**Objetivo**: Combinar todos os conceitos aprendidos em uma query complexa.

**Tarefas**:

Escreva uma query completa que mostre um relatório de empréstimos incluindo:
- Nome do usuário
- Email do usuário
- Título do livro
- Nome do autor
- Nacionalidade do autor
- Nome da categoria
- Data do empréstimo
- Status do empréstimo
- Apenas empréstimos ativos
- Ordenado por data de empréstimo (mais recente primeiro)

Depois, modifique a query para incluir também usuários que nunca pegaram livros (mostrando NULL para informações do livro).

**Questão de Reflexão**:
- Como você decidiria qual tipo de JOIN usar em uma query complexa? Qual seu processo de raciocínio?

**Soluções Esperadas**:

```sql
-- Query 1: Apenas empréstimos ativos
SELECT 
    u.nome AS usuario,
    u.email,
    l.titulo AS livro,
    a.nome AS autor,
    a.nacionalidade,
    c.nome AS categoria,
    e.data_emprestimo,
    e.status
FROM emprestimos e
INNER JOIN usuarios u ON e.usuario_id = u.id
INNER JOIN livros l ON e.livro_id = l.id
INNER JOIN autores a ON l.autor_id = a.id
INNER JOIN categorias c ON l.categoria_id = c.id
WHERE e.status = 'ativo'
ORDER BY e.data_emprestimo DESC;

-- Query 2: Incluindo usuários sem empréstimos
SELECT 
    u.nome AS usuario,
    u.email,
    l.titulo AS livro,
    a.nome AS autor,
    a.nacionalidade,
    c.nome AS categoria,
    e.data_emprestimo,
    e.status
FROM usuarios u
LEFT JOIN emprestimos e ON u.id = e.usuario_id AND e.status = 'ativo'
LEFT JOIN livros l ON e.livro_id = l.id
LEFT JOIN autores a ON l.autor_id = a.id
LEFT JOIN categorias c ON l.categoria_id = c.id
ORDER BY e.data_emprestimo DESC NULLS LAST, u.nome;
```

**Resposta Esperada para a Questão de Reflexão**:
- **Processo de decisão**:
  1. **Identifique a tabela principal**: Qual tabela você quer ver todos os registros?
  2. **Identifique relacionamentos**: Quais tabelas se relacionam e como?
  3. **Determine o resultado desejado**: Você quer apenas correspondências ou todos os registros?
  4. **Escolha o JOIN apropriado**:
     - Todos da principal? → LEFT JOIN
     - Apenas correspondências? → INNER JOIN
     - Comparar dentro da mesma tabela? → SELF JOIN
  5. **Teste e ajuste**: Execute a query e verifique se o resultado está correto
- **Exemplo prático**:
  - "Mostre todos os usuários e seus empréstimos"
  - Tabela principal: `usuarios`
  - Quer todos os usuários: LEFT JOIN
  - Relacionamento: `usuarios` → `emprestimos`
  - Resultado: LEFT JOIN de `usuarios` com `emprestimos`

---

## Perguntas de Reflexão Gerais

### 1. Performance e Otimização

- Como JOINs afetam a performance de uma query? Quais fatores influenciam a velocidade de execução?
- Quando você usaria índices para otimizar JOINs? Quais colunas devem ter índices?
- Como a ordem dos JOINs pode afetar a performance? Existe uma ordem "melhor"?

### 2. Design de Banco de Dados

- Como o design do banco de dados (normalização) afeta o uso de JOINs?
- Em que situações você consideraria desnormalizar para reduzir o número de JOINs?
- Como FOREIGN KEYs facilitam o uso de JOINs?

### 3. Boas Práticas

- Quais são as boas práticas ao escrever queries com JOINs?
- Como você tornaria uma query com múltiplos JOINs mais legível?
- Quando você usaria subqueries ao invés de JOINs?

### 4. Casos de Uso

- Dê exemplos de situações reais onde cada tipo de JOIN seria apropriado.
- Como você explicaria JOINs para alguém que não conhece SQL?
- Qual tipo de JOIN você usa com mais frequência e por quê?

---

## Checklist de Aprendizado

Antes de prosseguir, certifique-se de que você:

- [ ] Entende a diferença entre INNER JOIN, LEFT JOIN, RIGHT JOIN e FULL JOIN
- [ ] Sabe quando usar cada tipo de JOIN
- [ ] Consegue escrever queries com múltiplos JOINs
- [ ] Entende SELF JOIN e quando usá-lo
- [ ] Sabe evitar CROSS JOINs acidentais
- [ ] Entende a diferença entre condições no ON e no WHERE
- [ ] Consegue encontrar registros órfãos usando LEFT JOIN
- [ ] Sabe combinar JOINs com GROUP BY e HAVING
- [ ] Consegue escrever queries complexas combinando vários conceitos
- [ ] Entende o impacto de JOINs na performance

---

**Bons estudos! 🚀**

**Lembre-se**: Pratique muito! Execute todas as queries no banco de dados real e experimente variações. A prática é essencial para dominar JOINs!

