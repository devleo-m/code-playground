# Aula 4 - Exercícios e Reflexão

## Exercícios Práticos

### Exercício 1: Funções de Agregação Básicas

**Objetivo**: Praticar o uso das funções de agregação fundamentais (COUNT, SUM, AVG, MIN, MAX).

**Contexto**: Você precisa gerar estatísticas gerais sobre a biblioteca.

**Tarefas**:

1. Conte o total de livros na biblioteca usando `COUNT(*)`.

2. Conte quantos livros têm ano de publicação informado usando `COUNT(ano_publicacao)`.

3. Calcule o total de livros disponíveis em estoque usando `SUM(quantidade_disponivel)`.

4. Calcule a média de livros disponíveis por título usando `AVG(quantidade_disponivel)`.

5. Encontre o menor estoque disponível usando `MIN(quantidade_disponivel)`.

6. Encontre o maior estoque disponível usando `MAX(quantidade_disponivel)`.

7. Crie uma query que retorne todas essas estatísticas em uma única consulta.

**Soluções Esperadas**:

```sql
-- 1. Total de livros
SELECT COUNT(*) AS total_livros FROM livros;

-- 2. Livros com ano informado
SELECT COUNT(ano_publicacao) AS livros_com_ano FROM livros;

-- 3. Total em estoque
SELECT SUM(quantidade_disponivel) AS total_estoque FROM livros;

-- 4. Média de estoque
SELECT AVG(quantidade_disponivel) AS media_estoque FROM livros;

-- 5. Menor estoque
SELECT MIN(quantidade_disponivel) AS menor_estoque FROM livros;

-- 6. Maior estoque
SELECT MAX(quantidade_disponivel) AS maior_estoque FROM livros;

-- 7. Todas as estatísticas juntas
SELECT 
    COUNT(*) AS total_livros,
    COUNT(ano_publicacao) AS livros_com_ano,
    SUM(quantidade_disponivel) AS total_estoque,
    AVG(quantidade_disponivel) AS media_estoque,
    MIN(quantidade_disponivel) AS menor_estoque,
    MAX(quantidade_disponivel) AS maior_estoque
FROM livros;
```

---

### Exercício 2: GROUP BY - Agrupando por Categoria

**Objetivo**: Praticar o uso de GROUP BY para agrupar dados por categoria.

**Contexto**: Você precisa gerar um relatório mostrando estatísticas por categoria de livros.

**Tarefas**:

1. Conte quantos livros existem em cada categoria. Use JOIN para mostrar o nome da categoria.

2. Calcule o total de estoque (soma de `quantidade_disponivel`) por categoria.

3. Calcule a média de estoque por categoria.

4. Crie uma query que mostre, para cada categoria:
   - Nome da categoria
   - Total de livros
   - Total de estoque
   - Média de estoque
   - Menor estoque
   - Maior estoque

5. Ordene os resultados por total de livros (do maior para o menor).

**Soluções Esperadas**:

```sql
-- 1. Contar livros por categoria
SELECT 
    c.nome AS categoria,
    COUNT(*) AS total_livros
FROM livros l
JOIN categorias c ON l.categoria_id = c.id
GROUP BY c.id, c.nome;

-- 2. Total de estoque por categoria
SELECT 
    c.nome AS categoria,
    SUM(l.quantidade_disponivel) AS total_estoque
FROM livros l
JOIN categorias c ON l.categoria_id = c.id
GROUP BY c.id, c.nome;

-- 3. Média de estoque por categoria
SELECT 
    c.nome AS categoria,
    AVG(l.quantidade_disponivel) AS media_estoque
FROM livros l
JOIN categorias c ON l.categoria_id = c.id
GROUP BY c.id, c.nome;

-- 4. Estatísticas completas por categoria
SELECT 
    c.nome AS categoria,
    COUNT(*) AS total_livros,
    SUM(l.quantidade_disponivel) AS total_estoque,
    AVG(l.quantidade_disponivel) AS media_estoque,
    MIN(l.quantidade_disponivel) AS menor_estoque,
    MAX(l.quantidade_disponivel) AS maior_estoque
FROM livros l
JOIN categorias c ON l.categoria_id = c.id
GROUP BY c.id, c.nome
ORDER BY total_livros DESC;
```

---

### Exercício 3: GROUP BY - Agrupando por Autor

**Objetivo**: Praticar GROUP BY com diferentes colunas de agrupamento.

**Contexto**: Você precisa analisar a produção literária dos autores.

**Tarefas**:

1. Conte quantos livros cada autor tem cadastrado. Mostre o nome do autor.

2. Calcule o total de estoque de livros por autor.

3. Encontre quantos autores únicos têm livros cadastrados usando `COUNT(DISTINCT autor_id)`.

4. Crie uma query que mostre os top 5 autores com mais livros, ordenados do maior para o menor.

5. Crie uma query que mostre, para cada autor:
   - Nome do autor
   - Nacionalidade
   - Total de livros
   - Total de estoque
   - Média de estoque

**Soluções Esperadas**:

```sql
-- 1. Contar livros por autor
SELECT 
    a.nome AS autor,
    COUNT(*) AS total_livros
FROM livros l
JOIN autores a ON l.autor_id = a.id
GROUP BY a.id, a.nome;

-- 2. Total de estoque por autor
SELECT 
    a.nome AS autor,
    SUM(l.quantidade_disponivel) AS total_estoque
FROM livros l
JOIN autores a ON l.autor_id = a.id
GROUP BY a.id, a.nome;

-- 3. Autores únicos
SELECT COUNT(DISTINCT autor_id) AS total_autores
FROM livros
WHERE autor_id IS NOT NULL;

-- 4. Top 5 autores
SELECT 
    a.nome AS autor,
    COUNT(*) AS total_livros
FROM livros l
JOIN autores a ON l.autor_id = a.id
GROUP BY a.id, a.nome
ORDER BY total_livros DESC
LIMIT 5;

-- 5. Estatísticas por autor
SELECT 
    a.nome AS autor,
    a.nacionalidade,
    COUNT(*) AS total_livros,
    SUM(l.quantidade_disponivel) AS total_estoque,
    AVG(l.quantidade_disponivel) AS media_estoque
FROM livros l
JOIN autores a ON l.autor_id = a.id
GROUP BY a.id, a.nome, a.nacionalidade
ORDER BY total_livros DESC;
```

---

### Exercício 4: HAVING - Filtrando Grupos

**Objetivo**: Praticar o uso de HAVING para filtrar grupos baseado em condições agregadas.

**Contexto**: Você precisa identificar categorias e autores que atendem a critérios específicos.

**Tarefas**:

1. Encontre categorias que têm mais de 2 livros cadastrados.

2. Encontre categorias com total de estoque maior que 10 livros.

3. Encontre categorias com média de estoque maior que 5.

4. Encontre autores que têm mais de 1 livro cadastrado.

5. Encontre autores com total de estoque maior que 15 livros.

6. Crie uma query que mostre categorias com mais de 2 livros E média de estoque maior que 5.

**Soluções Esperadas**:

```sql
-- 1. Categorias com mais de 2 livros
SELECT 
    c.nome AS categoria,
    COUNT(*) AS total_livros
FROM livros l
JOIN categorias c ON l.categoria_id = c.id
GROUP BY c.id, c.nome
HAVING COUNT(*) > 2;

-- 2. Categorias com estoque > 10
SELECT 
    c.nome AS categoria,
    SUM(l.quantidade_disponivel) AS total_estoque
FROM livros l
JOIN categorias c ON l.categoria_id = c.id
GROUP BY c.id, c.nome
HAVING SUM(l.quantidade_disponivel) > 10;

-- 3. Categorias com média de estoque > 5
SELECT 
    c.nome AS categoria,
    AVG(l.quantidade_disponivel) AS media_estoque
FROM livros l
JOIN categorias c ON l.categoria_id = c.id
GROUP BY c.id, c.nome
HAVING AVG(l.quantidade_disponivel) > 5;

-- 4. Autores com mais de 1 livro
SELECT 
    a.nome AS autor,
    COUNT(*) AS total_livros
FROM livros l
JOIN autores a ON l.autor_id = a.id
GROUP BY a.id, a.nome
HAVING COUNT(*) > 1;

-- 5. Autores com estoque > 15
SELECT 
    a.nome AS autor,
    SUM(l.quantidade_disponivel) AS total_estoque
FROM livros l
JOIN autores a ON l.autor_id = a.id
GROUP BY a.id, a.nome
HAVING SUM(l.quantidade_disponivel) > 15;

-- 6. Categorias com múltiplas condições
SELECT 
    c.nome AS categoria,
    COUNT(*) AS total_livros,
    AVG(l.quantidade_disponivel) AS media_estoque
FROM livros l
JOIN categorias c ON l.categoria_id = c.id
GROUP BY c.id, c.nome
HAVING COUNT(*) > 2 
   AND AVG(l.quantidade_disponivel) > 5;
```

---

### Exercício 5: WHERE e HAVING Combinados

**Objetivo**: Entender a diferença e combinação entre WHERE e HAVING.

**Contexto**: Você precisa fazer análises mais complexas combinando filtros de linhas e grupos.

**Tarefas**:

1. Encontre categorias que têm livros publicados depois de 2000, mostrando quantos livros cada categoria tem nesse período.

2. Encontre autores que têm mais de 1 livro publicado depois de 2000.

3. Encontre categorias com média de estoque maior que 5, considerando apenas livros com estoque maior que 0.

4. Crie uma query que mostre:
   - Categorias
   - Total de livros publicados depois de 2000
   - Média de estoque desses livros
   - Apenas categorias com mais de 1 livro nesse período

**Soluções Esperadas**:

```sql
-- 1. Categorias com livros depois de 2000
SELECT 
    c.nome AS categoria,
    COUNT(*) AS total_livros
FROM livros l
JOIN categorias c ON l.categoria_id = c.id
WHERE l.ano_publicacao > 2000
GROUP BY c.id, c.nome;

-- 2. Autores com mais de 1 livro depois de 2000
SELECT 
    a.nome AS autor,
    COUNT(*) AS total_livros
FROM livros l
JOIN autores a ON l.autor_id = a.id
WHERE l.ano_publicacao > 2000
GROUP BY a.id, a.nome
HAVING COUNT(*) > 1;

-- 3. Categorias com média de estoque > 5 (apenas livros com estoque > 0)
SELECT 
    c.nome AS categoria,
    AVG(l.quantidade_disponivel) AS media_estoque
FROM livros l
JOIN categorias c ON l.categoria_id = c.id
WHERE l.quantidade_disponivel > 0
GROUP BY c.id, c.nome
HAVING AVG(l.quantidade_disponivel) > 5;

-- 4. Query completa combinando WHERE e HAVING
SELECT 
    c.nome AS categoria,
    COUNT(*) AS total_livros,
    AVG(l.quantidade_disponivel) AS media_estoque
FROM livros l
JOIN categorias c ON l.categoria_id = c.id
WHERE l.ano_publicacao > 2000
GROUP BY c.id, c.nome
HAVING COUNT(*) > 1
ORDER BY total_livros DESC;
```

---

### Exercício 6: Análise de Empréstimos

**Objetivo**: Aplicar aggregate queries em uma tabela diferente (empréstimos).

**Contexto**: Você precisa analisar os padrões de empréstimos da biblioteca.

**Tarefas**:

1. Conte quantos empréstimos existem no total.

2. Conte quantos empréstimos estão ativos (status = 'ativo').

3. Conte quantos empréstimos cada usuário tem (ativos e totais).

4. Encontre usuários com mais de 1 empréstimo ativo.

5. Conte quantos empréstimos foram feitos por livro.

6. Encontre livros que foram emprestados mais de 1 vez.

**Soluções Esperadas**:

```sql
-- 1. Total de empréstimos
SELECT COUNT(*) AS total_emprestimos FROM emprestimos;

-- 2. Empréstimos ativos
SELECT COUNT(*) AS emprestimos_ativos 
FROM emprestimos 
WHERE status = 'ativo';

-- 3. Empréstimos por usuário
SELECT 
    u.nome AS usuario,
    COUNT(*) AS total_emprestimos,
    SUM(CASE WHEN e.status = 'ativo' THEN 1 ELSE 0 END) AS emprestimos_ativos
FROM emprestimos e
JOIN usuarios u ON e.usuario_id = u.id
GROUP BY u.id, u.nome;

-- 4. Usuários com mais de 1 empréstimo ativo
SELECT 
    u.nome AS usuario,
    COUNT(*) AS emprestimos_ativos
FROM emprestimos e
JOIN usuarios u ON e.usuario_id = u.id
WHERE e.status = 'ativo'
GROUP BY u.id, u.nome
HAVING COUNT(*) > 1;

-- 5. Empréstimos por livro
SELECT 
    l.titulo AS livro,
    COUNT(*) AS total_emprestimos
FROM emprestimos e
JOIN livros l ON e.livro_id = l.id
GROUP BY l.id, l.titulo;

-- 6. Livros emprestados mais de 1 vez
SELECT 
    l.titulo AS livro,
    COUNT(*) AS total_emprestimos
FROM emprestimos e
JOIN livros l ON e.livro_id = l.id
GROUP BY l.id, l.titulo
HAVING COUNT(*) > 1;
```

---

### Exercício 7: Análise Temporal

**Objetivo**: Praticar agregações com dados de data.

**Contexto**: Você precisa analisar padrões temporais nos empréstimos.

**Tarefas**:

1. Conte quantos empréstimos foram feitos por mês (use `strftime('%Y-%m', data_emprestimo)`).

2. Encontre o mês com mais empréstimos.

3. Conte quantos empréstimos foram feitos por ano de publicação do livro.

4. Encontre a década de publicação com mais livros cadastrados.

**Soluções Esperadas**:

```sql
-- 1. Empréstimos por mês
SELECT 
    strftime('%Y-%m', data_emprestimo) AS mes,
    COUNT(*) AS total_emprestimos
FROM emprestimos
GROUP BY strftime('%Y-%m', data_emprestimo)
ORDER BY mes;

-- 2. Mês com mais empréstimos
SELECT 
    strftime('%Y-%m', data_emprestimo) AS mes,
    COUNT(*) AS total_emprestimos
FROM emprestimos
GROUP BY strftime('%Y-%m', data_emprestimo)
ORDER BY total_emprestimos DESC
LIMIT 1;

-- 3. Empréstimos por ano de publicação
SELECT 
    l.ano_publicacao AS ano,
    COUNT(*) AS total_emprestimos
FROM emprestimos e
JOIN livros l ON e.livro_id = l.id
WHERE l.ano_publicacao IS NOT NULL
GROUP BY l.ano_publicacao
ORDER BY ano;

-- 4. Década com mais livros
SELECT 
    (ano_publicacao / 10) * 10 AS decada,
    COUNT(*) AS total_livros
FROM livros
WHERE ano_publicacao IS NOT NULL
GROUP BY (ano_publicacao / 10) * 10
ORDER BY total_livros DESC
LIMIT 1;
```

---

## Perguntas de Reflexão

### Reflexão 1: Performance e Eficiência

**Pergunta**: Considere a seguinte query:

```sql
SELECT 
    c.nome AS categoria,
    COUNT(*) AS total_livros,
    SUM(l.quantidade_disponivel) AS total_estoque
FROM livros l
JOIN categorias c ON l.categoria_id = c.id
GROUP BY c.id, c.nome
HAVING COUNT(*) > 2;
```

1. **O que acontece se a tabela `livros` tiver 1 milhão de registros?** A query ainda será eficiente? Por quê?

2. **Quais índices seriam úteis para otimizar essa query?** Pense em quais colunas são usadas em JOINs, WHERE, GROUP BY e HAVING.

3. **A ordem das operações (WHERE antes de GROUP BY, HAVING depois) afeta a performance?** Explique.

4. **Se você precisasse executar essa query frequentemente (centenas de vezes por dia), o que você faria para otimizá-la?**

**Respostas Esperadas (guia de pensamento)**:

1. Com 1 milhão de registros, a query precisará:
   - Fazer JOIN entre duas tabelas grandes
   - Agrupar milhões de linhas
   - Aplicar funções de agregação
   - Filtrar grupos
   Isso pode ser lento sem índices adequados.

2. Índices úteis:
   - `idx_livros_categoria` em `livros(categoria_id)` para o JOIN
   - Índice em `categorias(id)` (geralmente já existe como PRIMARY KEY)
   - Índice composto pode ajudar se houver WHERE adicional

3. Sim, a ordem importa:
   - WHERE filtra antes de agrupar, reduzindo o número de linhas processadas
   - HAVING filtra depois, mas ainda precisa processar os grupos
   - Filtrar com WHERE quando possível é mais eficiente

4. Otimizações possíveis:
   - Criar índices apropriados
   - Considerar materializar resultados em uma view ou tabela de cache
   - Usar EXPLAIN QUERY PLAN para analisar o plano de execução
   - Considerar desnormalização se a query for crítica

---

### Reflexão 2: Compreensão de NULL

**Pergunta**: Considere as seguintes queries:

```sql
-- Query 1
SELECT COUNT(*) FROM livros;

-- Query 2
SELECT COUNT(ano_publicacao) FROM livros;

-- Query 3
SELECT AVG(ano_publicacao) FROM livros;
```

1. **Se a tabela `livros` tiver 100 registros, mas apenas 80 têm `ano_publicacao` informado, quais serão os resultados de cada query?**

2. **Por que `COUNT(*)` e `COUNT(ano_publicacao)` podem retornar valores diferentes?**

3. **Se todos os registros tiverem `ano_publicacao = NULL`, qual será o resultado de `AVG(ano_publicacao)`? Por quê?**

4. **Como você garantiria que `AVG(ano_publicacao)` calcule apenas livros com ano informado?**

**Respostas Esperadas (guia de pensamento)**:

1. 
   - Query 1: 100 (conta todas as linhas)
   - Query 2: 80 (conta apenas linhas com ano não-nulo)
   - Query 3: Média dos 80 valores não-nulos

2. `COUNT(*)` conta linhas, `COUNT(coluna)` conta valores não-nulos na coluna.

3. `AVG(ano_publicacao)` retornaria NULL, pois não há valores para calcular a média.

4. Usar `WHERE ano_publicacao IS NOT NULL` ou garantir que a coluna tenha valores.

---

### Reflexão 3: GROUP BY e Integridade de Dados

**Pergunta**: Considere a seguinte situação:

Você tem uma query que agrupa livros por categoria e calcula estatísticas. Mas alguns livros têm `categoria_id = NULL`.

```sql
SELECT 
    c.nome AS categoria,
    COUNT(*) AS total_livros
FROM livros l
LEFT JOIN categorias c ON l.categoria_id = c.id
GROUP BY c.id, c.nome;
```

1. **O que acontece com livros que têm `categoria_id = NULL`?** Eles aparecerão nos resultados?

2. **Se você usar `INNER JOIN` ao invés de `LEFT JOIN`, o que muda?**

3. **Como você contaria livros sem categoria separadamente?**

4. **Em um sistema real, é melhor permitir `categoria_id = NULL` ou forçar que todos os livros tenham categoria?** Por quê?

**Respostas Esperadas (guia de pensamento)**:

1. Com LEFT JOIN, livros sem categoria aparecerão com `categoria = NULL`.

2. Com INNER JOIN, livros sem categoria seriam excluídos completamente dos resultados.

3. Usar `COALESCE(c.nome, 'Sem categoria')` ou `CASE WHEN` para agrupar NULLs.

4. Depende do negócio, mas geralmente é melhor forçar categoria (NOT NULL) para manter integridade dos dados.

---

### Reflexão 4: HAVING vs WHERE

**Pergunta**: Analise as seguintes queries:

```sql
-- Query A
SELECT categoria_id, COUNT(*)
FROM livros
WHERE quantidade_disponivel > 0
GROUP BY categoria_id;

-- Query B
SELECT categoria_id, COUNT(*)
FROM livros
GROUP BY categoria_id
HAVING COUNT(*) > 2;

-- Query C
SELECT categoria_id, COUNT(*)
FROM livros
WHERE quantidade_disponivel > 0
GROUP BY categoria_id
HAVING COUNT(*) > 2;
```

1. **Qual é a diferença prática entre Query A e Query B?** O que cada uma faz?

2. **Query C combina WHERE e HAVING. Em que ordem as operações acontecem?**

3. **É possível reescrever Query B usando apenas WHERE?** Por quê?

4. **Em termos de performance, qual é mais eficiente: filtrar com WHERE ou com HAVING?** Por quê?

**Respostas Esperadas (guia de pensamento)**:

1. 
   - Query A: Filtra livros com estoque > 0, depois agrupa
   - Query B: Agrupa todos os livros, depois mostra apenas grupos com mais de 2 livros

2. Ordem: WHERE (filtra linhas) → GROUP BY (agrupa) → HAVING (filtra grupos) → SELECT

3. Não, porque COUNT(*) é uma agregação que só existe depois do GROUP BY.

4. WHERE é geralmente mais eficiente porque filtra antes de agrupar, reduzindo dados processados.

---

### Reflexão 5: Escalabilidade e Design

**Pergunta**: Imagine que você precisa criar um dashboard que mostra:

- Total de livros por categoria
- Total de estoque por categoria
- Média de estoque por categoria
- Top 5 categorias com mais livros

Essas informações precisam ser atualizadas em tempo real sempre que alguém acessa o dashboard.

1. **Se o banco tiver 10 milhões de livros e 1000 categorias, essa query será rápida?** Quais fatores afetam a performance?

2. **Quais estratégias você usaria para otimizar esse dashboard?**

3. **Seria melhor calcular essas estatísticas de forma diferente?** (ex: armazenar em tabela separada, usar views materializadas, cache)

4. **Como você monitoraria a performance dessa query em produção?**

**Respostas Esperadas (guia de pensamento)**:

1. Pode ser lenta sem otimização:
   - JOIN em tabelas grandes
   - Agrupamento de milhões de linhas
   - Múltiplas agregações

2. Estratégias:
   - Índices apropriados
   - Views materializadas
   - Cache de resultados
   - Tabela de estatísticas atualizada periodicamente

3. Para dados que mudam pouco, pode ser melhor:
   - Calcular estatísticas em background
   - Armazenar em tabela de cache
   - Atualizar incrementalmente

4. Monitoramento:
   - Usar EXPLAIN QUERY PLAN
   - Logs de queries lentas
   - Métricas de tempo de resposta
   - Alertas para queries acima de threshold

---

## Checklist de Aprendizado

Antes de prosseguir, certifique-se de que você:

- [ ] Consegue usar COUNT, SUM, AVG, MIN, MAX corretamente
- [ ] Entende a diferença entre COUNT(*) e COUNT(coluna)
- [ ] Sabe quando e como usar GROUP BY
- [ ] Compreende a diferença entre WHERE e HAVING
- [ ] Consegue combinar múltiplas funções de agregação
- [ ] Entende como NULL é tratado em agregações
- [ ] Consegue criar queries complexas com JOIN, WHERE, GROUP BY e HAVING
- [ ] Sabe identificar erros comuns em aggregate queries
- [ ] Pensa sobre performance ao escrever queries de agregação

---

## Próximos Passos

Após completar todos os exercícios e responder as perguntas de reflexão:

1. Revise suas respostas
2. Execute todas as queries no banco de dados real
3. Experimente variações dos exercícios
4. Quando estiver confiante, envie suas respostas para análise

**⚠️ IMPORTANTE**: Não pule as perguntas de reflexão! Elas são cruciais para desenvolver pensamento crítico sobre performance, escalabilidade e design de queries.

---

**Bons estudos! 🚀**
