# Aula 7 - Exercícios e Reflexão

## Exercícios Práticos

### Exercício 1: Scalar Subqueries Básicas

**Objetivo**: Praticar o uso de subqueries escalares (que retornam um único valor).

**Contexto**: Você precisa fazer comparações e cálculos usando valores únicos calculados.

**Tarefas**:

1. Encontre todos os livros com estoque acima da média de estoque de todos os livros.

2. Encontre livros com quantidade disponível maior que o maior estoque de livros de ficção científica.

3. Liste todos os livros com o total de empréstimos de cada um usando uma subquery escalar no SELECT.

4. Crie uma query que mostre:
   - Título do livro
   - Quantidade disponível
   - Média geral de estoque (usando subquery)
   - Diferença entre o estoque do livro e a média

5. Encontre o livro com mais empréstimos usando uma subquery escalar.

**Soluções Esperadas**:

```sql
-- 1. Livros com estoque acima da média
SELECT titulo, quantidade_disponivel
FROM livros
WHERE quantidade_disponivel > (
    SELECT AVG(quantidade_disponivel) FROM livros
);

-- 2. Livros com estoque maior que o máximo de ficção científica
SELECT titulo, quantidade_disponivel
FROM livros
WHERE quantidade_disponivel > (
    SELECT MAX(quantidade_disponivel)
    FROM livros
    WHERE categoria_id = (
        SELECT id FROM categorias WHERE nome = 'Ficção Científica'
    )
);

-- 3. Livros com total de empréstimos
SELECT 
    titulo,
    (SELECT COUNT(*) 
     FROM emprestimos 
     WHERE emprestimos.livro_id = livros.id) AS total_emprestimos
FROM livros;

-- 4. Livros com comparação com média
SELECT 
    titulo,
    quantidade_disponivel,
    (SELECT AVG(quantidade_disponivel) FROM livros) AS media_estoque,
    quantidade_disponivel - (SELECT AVG(quantidade_disponivel) FROM livros) AS diferenca_media
FROM livros;

-- 5. Livro com mais empréstimos
SELECT titulo
FROM livros
WHERE id = (
    SELECT livro_id
    FROM emprestimos
    GROUP BY livro_id
    ORDER BY COUNT(*) DESC
    LIMIT 1
);
```

---

### Exercício 2: Column Subqueries com IN, NOT IN, EXISTS

**Objetivo**: Praticar o uso de subqueries que retornam uma coluna de valores com diferentes operadores.

**Contexto**: Você precisa filtrar dados baseado em listas de valores.

**Tarefas**:

1. Encontre todos os livros de autores brasileiros usando IN.

2. Encontre livros que nunca foram emprestados usando NOT IN.

3. Encontre categorias que têm livros emprestados usando EXISTS.

4. Encontre usuários que pegaram livros emprestados de autores brasileiros usando EXISTS.

5. Encontre livros que não são de ficção científica usando NOT IN.

6. Crie uma query que encontre autores que têm livros com estoque acima da média usando EXISTS.

**Soluções Esperadas**:

```sql
-- 1. Livros de autores brasileiros
SELECT titulo
FROM livros
WHERE autor_id IN (
    SELECT id FROM autores WHERE nacionalidade = 'Brasileiro'
);

-- 2. Livros nunca emprestados
SELECT titulo
FROM livros
WHERE id NOT IN (
    SELECT DISTINCT livro_id 
    FROM emprestimos
    WHERE livro_id IS NOT NULL
);

-- 3. Categorias com livros emprestados
SELECT nome
FROM categorias c
WHERE EXISTS (
    SELECT 1
    FROM livros l
    JOIN emprestimos e ON l.id = e.livro_id
    WHERE l.categoria_id = c.id
);

-- 4. Usuários com livros de autores brasileiros
SELECT DISTINCT u.nome
FROM usuarios u
WHERE EXISTS (
    SELECT 1
    FROM emprestimos e
    JOIN livros l ON e.livro_id = l.id
    JOIN autores a ON l.autor_id = a.id
    WHERE e.usuario_id = u.id
    AND a.nacionalidade = 'Brasileiro'
);

-- 5. Livros que não são de ficção científica
SELECT titulo
FROM livros
WHERE categoria_id NOT IN (
    SELECT id FROM categorias WHERE nome = 'Ficção Científica'
);

-- 6. Autores com livros acima da média
SELECT nome
FROM autores a
WHERE EXISTS (
    SELECT 1
    FROM livros l
    WHERE l.autor_id = a.id
    AND l.quantidade_disponivel > (
        SELECT AVG(quantidade_disponivel) FROM livros
    )
);
```

---

### Exercício 3: Table Subqueries (Subqueries em FROM)

**Objetivo**: Praticar o uso de subqueries que retornam tabelas completas na cláusula FROM.

**Contexto**: Você precisa criar tabelas temporárias para fazer análises mais complexas.

**Tarefas**:

1. Crie uma query que mostre estatísticas de empréstimos por categoria usando uma subquery em FROM.

2. Crie uma query que mostre os 3 livros mais emprestados de cada categoria usando subquery em FROM.

3. Crie uma query que calcule a média de empréstimos por autor usando uma tabela derivada.

4. Crie uma query que mostre categorias com suas estatísticas (total de livros, total de empréstimos, média de empréstimos por livro) usando subquery em FROM.

**Soluções Esperadas**:

```sql
-- 1. Estatísticas de empréstimos por categoria
SELECT 
    categoria,
    total_emprestimos,
    total_livros,
    ROUND(CAST(total_emprestimos AS REAL) / total_livros, 2) AS media_por_livro
FROM (
    SELECT 
        c.nome AS categoria,
        COUNT(DISTINCT e.id) AS total_emprestimos,
        COUNT(DISTINCT l.id) AS total_livros
    FROM categorias c
    LEFT JOIN livros l ON c.id = l.categoria_id
    LEFT JOIN emprestimos e ON l.id = e.livro_id
    GROUP BY c.id, c.nome
) AS estatisticas
ORDER BY total_emprestimos DESC;

-- 2. Top 3 livros mais emprestados por categoria
SELECT 
    categoria,
    titulo,
    total_emprestimos
FROM (
    SELECT 
        c.nome AS categoria,
        l.titulo,
        COUNT(e.id) AS total_emprestimos,
        ROW_NUMBER() OVER (PARTITION BY c.id ORDER BY COUNT(e.id) DESC) AS rank
    FROM categorias c
    JOIN livros l ON c.id = l.categoria_id
    LEFT JOIN emprestimos e ON l.id = e.livro_id
    GROUP BY c.id, c.nome, l.id, l.titulo
) AS ranked
WHERE rank <= 3
ORDER BY categoria, total_emprestimos DESC;

-- 3. Média de empréstimos por autor
SELECT 
    autor,
    total_livros,
    total_emprestimos,
    ROUND(CAST(total_emprestimos AS REAL) / total_livros, 2) AS media_emprestimos_por_livro
FROM (
    SELECT 
        a.nome AS autor,
        COUNT(DISTINCT l.id) AS total_livros,
        COUNT(e.id) AS total_emprestimos
    FROM autores a
    LEFT JOIN livros l ON a.id = l.autor_id
    LEFT JOIN emprestimos e ON l.id = e.livro_id
    GROUP BY a.id, a.nome
) AS autor_stats
WHERE total_livros > 0
ORDER BY media_emprestimos_por_livro DESC;

-- 4. Estatísticas completas por categoria
SELECT 
    categoria,
    total_livros,
    total_estoque,
    total_emprestimos,
    ROUND(CAST(total_emprestimos AS REAL) / total_livros, 2) AS media_emprestimos_por_livro
FROM (
    SELECT 
        c.nome AS categoria,
        COUNT(DISTINCT l.id) AS total_livros,
        SUM(l.quantidade_disponivel) AS total_estoque,
        COUNT(e.id) AS total_emprestimos
    FROM categorias c
    LEFT JOIN livros l ON c.id = l.categoria_id
    LEFT JOIN emprestimos e ON l.id = e.livro_id
    GROUP BY c.id, c.nome
) AS categoria_stats
ORDER BY total_emprestimos DESC;
```

---

### Exercício 4: Correlated Subqueries

**Objetivo**: Praticar o uso de subqueries correlacionadas que usam valores da query externa.

**Contexto**: Você precisa fazer comparações que dependem de cada registro individual.

**Tarefas**:

1. Encontre livros com estoque acima da média de sua categoria usando correlated subquery.

2. Liste todos os livros com o número de empréstimos de cada um usando correlated subquery no SELECT.

3. Encontre autores que têm pelo menos um livro com estoque acima da média geral usando EXISTS.

4. Encontre categorias onde todos os livros têm estoque acima da média geral usando NOT EXISTS.

5. Crie uma query que mostre, para cada livro, quantos outros livros da mesma categoria têm mais empréstimos.

**Soluções Esperadas**:

```sql
-- 1. Livros com estoque acima da média da categoria
SELECT 
    l1.titulo,
    l1.quantidade_disponivel,
    c.nome AS categoria
FROM livros l1
JOIN categorias c ON l1.categoria_id = c.id
WHERE l1.quantidade_disponivel > (
    SELECT AVG(l2.quantidade_disponivel)
    FROM livros l2
    WHERE l2.categoria_id = l1.categoria_id
);

-- 2. Livros com total de empréstimos
SELECT 
    titulo,
    (SELECT COUNT(*)
     FROM emprestimos e
     WHERE e.livro_id = livros.id) AS total_emprestimos
FROM livros;

-- 3. Autores com livros acima da média
SELECT DISTINCT a.nome
FROM autores a
WHERE EXISTS (
    SELECT 1
    FROM livros l
    WHERE l.autor_id = a.id
    AND l.quantidade_disponivel > (
        SELECT AVG(quantidade_disponivel) FROM livros
    )
);

-- 4. Categorias onde todos os livros estão acima da média
SELECT c.nome
FROM categorias c
WHERE NOT EXISTS (
    SELECT 1
    FROM livros l
    WHERE l.categoria_id = c.id
    AND l.quantidade_disponivel <= (
        SELECT AVG(quantidade_disponivel) FROM livros
    )
);

-- 5. Ranking de empréstimos dentro da categoria
SELECT 
    l1.titulo,
    c.nome AS categoria,
    (SELECT COUNT(*) FROM emprestimos WHERE livro_id = l1.id) AS total_emprestimos,
    (SELECT COUNT(*)
     FROM livros l2
     WHERE l2.categoria_id = l1.categoria_id
     AND (SELECT COUNT(*) FROM emprestimos WHERE livro_id = l2.id) > 
         (SELECT COUNT(*) FROM emprestimos WHERE livro_id = l1.id)
    ) + 1 AS ranking_categoria
FROM livros l1
JOIN categorias c ON l1.categoria_id = c.id
ORDER BY c.nome, total_emprestimos DESC;
```

---

### Exercício 5: Nested Subqueries (Subqueries Aninhadas)

**Objetivo**: Praticar o uso de subqueries dentro de outras subqueries.

**Contexto**: Você precisa resolver problemas complexos que requerem múltiplas etapas.

**Tarefas**:

1. Encontre livros de autores que têm mais livros que a média de livros por autor.

2. Encontre categorias onde a média de empréstimos por livro é maior que a média geral de empréstimos por livro.

3. Encontre usuários que pegaram emprestado livros de categorias que têm mais livros que a média.

4. Crie uma query complexa que encontre o livro mais emprestado de cada categoria, mostrando também quantos empréstimos ele tem.

**Soluções Esperadas**:

```sql
-- 1. Livros de autores prolíficos
SELECT titulo
FROM livros
WHERE autor_id IN (
    SELECT id
    FROM autores
    WHERE id IN (
        SELECT autor_id
        FROM livros
        GROUP BY autor_id
        HAVING COUNT(*) > (
            SELECT AVG(total)
            FROM (
                SELECT COUNT(*) AS total
                FROM livros
                GROUP BY autor_id
            )
        )
    )
);

-- 2. Categorias com alta média de empréstimos
SELECT c.nome
FROM categorias c
WHERE c.id IN (
    SELECT categoria_id
    FROM livros
    WHERE id IN (
        SELECT livro_id
        FROM emprestimos
        GROUP BY livro_id
        HAVING COUNT(*) > (
            SELECT AVG(total_emprestimos)
            FROM (
                SELECT COUNT(*) AS total_emprestimos
                FROM emprestimos
                GROUP BY livro_id
            )
        )
    )
    GROUP BY categoria_id
    HAVING AVG((
        SELECT COUNT(*)
        FROM emprestimos
        WHERE emprestimos.livro_id = livros.id
    )) > (
        SELECT AVG(total_emprestimos)
        FROM (
            SELECT COUNT(*) AS total_emprestimos
            FROM emprestimos
            GROUP BY livro_id
        )
    )
);

-- 3. Usuários com livros de categorias grandes
SELECT DISTINCT u.nome
FROM usuarios u
WHERE u.id IN (
    SELECT usuario_id
    FROM emprestimos
    WHERE livro_id IN (
        SELECT id
        FROM livros
        WHERE categoria_id IN (
            SELECT id
            FROM categorias
            WHERE id IN (
                SELECT categoria_id
                FROM livros
                GROUP BY categoria_id
                HAVING COUNT(*) > (
                    SELECT AVG(total)
                    FROM (
                        SELECT COUNT(*) AS total
                        FROM livros
                        GROUP BY categoria_id
                    )
                )
            )
        )
    )
);

-- 4. Livro mais emprestado de cada categoria
SELECT 
    categoria,
    titulo,
    total_emprestimos
FROM (
    SELECT 
        c.nome AS categoria,
        l.titulo,
        (SELECT COUNT(*) FROM emprestimos WHERE livro_id = l.id) AS total_emprestimos,
        (SELECT COUNT(*)
         FROM livros l2
         WHERE l2.categoria_id = l.categoria_id
         AND (SELECT COUNT(*) FROM emprestimos WHERE livro_id = l2.id) > 
             (SELECT COUNT(*) FROM emprestimos WHERE livro_id = l.id)
        ) AS ranking
    FROM livros l
    JOIN categorias c ON l.categoria_id = c.id
) AS ranked
WHERE ranking = 0
ORDER BY categoria, total_emprestimos DESC;
```

---

### Exercício 6: Subqueries em HAVING

**Objetivo**: Praticar o uso de subqueries na cláusula HAVING para filtrar grupos.

**Contexto**: Você precisa filtrar grupos baseado em comparações com outras consultas.

**Tarefas**:

1. Encontre categorias com média de estoque acima da média geral de estoque.

2. Encontre autores com mais livros que a média de livros por autor.

3. Encontre categorias onde o total de empréstimos é maior que a média de empréstimos por categoria.

4. Crie uma query que mostre usuários que pegaram mais livros emprestados que a média.

**Soluções Esperadas**:

```sql
-- 1. Categorias com média de estoque acima da média geral
SELECT 
    c.nome AS categoria,
    AVG(l.quantidade_disponivel) AS media_estoque
FROM categorias c
JOIN livros l ON c.id = l.categoria_id
GROUP BY c.id, c.nome
HAVING AVG(l.quantidade_disponivel) > (
    SELECT AVG(quantidade_disponivel) FROM livros
);

-- 2. Autores com mais livros que a média
SELECT 
    a.nome AS autor,
    COUNT(l.id) AS total_livros
FROM autores a
LEFT JOIN livros l ON a.id = l.autor_id
GROUP BY a.id, a.nome
HAVING COUNT(l.id) > (
    SELECT AVG(total)
    FROM (
        SELECT COUNT(*) AS total
        FROM livros
        GROUP BY autor_id
    )
);

-- 3. Categorias com total de empréstimos acima da média
SELECT 
    c.nome AS categoria,
    COUNT(e.id) AS total_emprestimos
FROM categorias c
LEFT JOIN livros l ON c.id = l.categoria_id
LEFT JOIN emprestimos e ON l.id = e.livro_id
GROUP BY c.id, c.nome
HAVING COUNT(e.id) > (
    SELECT AVG(total)
    FROM (
        SELECT COUNT(*) AS total
        FROM emprestimos e2
        JOIN livros l2 ON e2.livro_id = l2.id
        GROUP BY l2.categoria_id
    )
);

-- 4. Usuários com mais empréstimos que a média
SELECT 
    u.nome,
    COUNT(e.id) AS total_emprestimos
FROM usuarios u
LEFT JOIN emprestimos e ON u.id = e.usuario_id
GROUP BY u.id, u.nome
HAVING COUNT(e.id) > (
    SELECT AVG(total)
    FROM (
        SELECT COUNT(*) AS total
        FROM emprestimos
        GROUP BY usuario_id
    )
)
ORDER BY total_emprestimos DESC;
```

---

### Exercício 7: Subqueries vs JOINs - Escolhendo a Abordagem Correta

**Objetivo**: Praticar a escolha entre subqueries e JOINs para resolver o mesmo problema.

**Contexto**: Você precisa resolver problemas de diferentes formas e entender quando usar cada abordagem.

**Tarefas**:

1. Resolva o problema "encontrar livros que têm empréstimos ativos" usando:
   - a) Subquery com EXISTS
   - b) JOIN
   - Compare os resultados e explique qual é mais eficiente.

2. Resolva o problema "encontrar livros de autores brasileiros" usando:
   - a) Subquery com IN
   - b) JOIN
   - Compare e explique quando cada abordagem é melhor.

3. Resolva o problema "encontrar livros com estoque acima da média" usando:
   - a) Subquery escalar
   - b) JOIN com tabela derivada
   - Compare a legibilidade de cada abordagem.

4. Para cada problema acima, escreva uma reflexão sobre:
   - Qual abordagem é mais legível?
   - Qual é mais eficiente?
   - Quando você escolheria cada uma?

**Soluções Esperadas**:

```sql
-- 1a. Livros com empréstimos ativos - EXISTS
SELECT titulo
FROM livros
WHERE EXISTS (
    SELECT 1
    FROM emprestimos
    WHERE emprestimos.livro_id = livros.id
    AND emprestimos.status = 'ativo'
);

-- 1b. Livros com empréstimos ativos - JOIN
SELECT DISTINCT l.titulo
FROM livros l
JOIN emprestimos e ON l.id = e.livro_id
WHERE e.status = 'ativo';

-- 2a. Livros de autores brasileiros - IN
SELECT titulo
FROM livros
WHERE autor_id IN (
    SELECT id FROM autores WHERE nacionalidade = 'Brasileiro'
);

-- 2b. Livros de autores brasileiros - JOIN
SELECT l.titulo, a.nome AS autor
FROM livros l
JOIN autores a ON l.autor_id = a.id
WHERE a.nacionalidade = 'Brasileiro';

-- 3a. Livros com estoque acima da média - Subquery
SELECT titulo, quantidade_disponivel
FROM livros
WHERE quantidade_disponivel > (
    SELECT AVG(quantidade_disponivel) FROM livros
);

-- 3b. Livros com estoque acima da média - JOIN com tabela derivada
SELECT l.titulo, l.quantidade_disponivel
FROM livros l
CROSS JOIN (
    SELECT AVG(quantidade_disponivel) AS media FROM livros
) AS stats
WHERE l.quantidade_disponivel > stats.media;
```

---

## Perguntas de Reflexão

### Reflexão 1: Performance e Eficiência

1. **Correlated Subqueries**: Por que correlated subqueries podem ser mais lentas que JOINs? Dê um exemplo prático.

2. **Subquery vs JOIN**: Em que situações uma subquery seria mais eficiente que um JOIN? E vice-versa?

3. **Nested Subqueries**: Quais são os riscos de performance ao usar muitas subqueries aninhadas? Como você pode mitigar esses riscos?

4. **Índices**: Como índices podem melhorar a performance de subqueries? Dê exemplos específicos.

### Reflexão 2: Legibilidade e Manutenibilidade

1. **Complexidade**: Quando uma subquery se torna muito complexa? Qual é o limite de aninhamento que você consideraria aceitável?

2. **Documentação**: Como você documentaria uma subquery complexa para que outros desenvolvedores possam entendê-la?

3. **Refatoração**: Quando você consideraria reescrever uma subquery como JOIN ou CTE? Quais critérios você usaria?

4. **Testabilidade**: Como você testaria uma subquery complexa? Qual é a melhor estratégia?

### Reflexão 3: Escolha de Abordagem

1. **Subquery vs JOIN**: Crie um critério de decisão (checklist) para escolher entre subquery e JOIN. Quando você escolheria cada um?

2. **EXISTS vs IN**: Qual é a diferença prática entre EXISTS e IN? Quando cada um é mais apropriado?

3. **Correlated vs Non-Correlated**: Quando você escolheria uma correlated subquery ao invés de uma não-correlacionada? E vice-versa?

4. **Table Subquery vs View**: Quando você criaria uma view ao invés de usar uma table subquery? Quais são as vantagens e desvantagens?

### Reflexão 4: Casos de Uso Reais

1. **Análise de Dados**: Dê um exemplo de um problema de análise de dados que seria melhor resolvido com subqueries do que com JOINs.

2. **Relatórios**: Como você usaria subqueries para criar relatórios complexos? Dê um exemplo prático.

3. **Validação de Dados**: Como subqueries podem ser usadas para validar integridade de dados? Dê exemplos.

4. **Business Logic**: Como subqueries podem implementar regras de negócio complexas diretamente no banco de dados? Isso é uma boa prática?

### Reflexão 5: Erros Comuns e Soluções

1. **Erro "Subquery returns more than one row"**: Quando esse erro ocorre? Como você o resolveria? Dê exemplos.

2. **Performance Degradada**: Você escreveu uma query com subquery que está muito lenta. Qual é o seu processo de debugging? Quais são as primeiras coisas que você verifica?

3. **Resultados Inesperados**: Sua subquery retorna resultados diferentes do esperado. Qual é a sua estratégia para debugar? Como você isola o problema?

4. **NULL Handling**: Como NULLs são tratados em subqueries? Quais são as armadilhas comuns? Dê exemplos.

---

## Desafios Avançados

### Desafio 1: Análise Complexa de Performance

Crie uma query que identifique:
- Livros que têm mais empréstimos que a média de empréstimos de livros da mesma categoria
- E que também têm estoque acima da média geral
- E que são de autores que têm mais livros que a média de livros por autor

Use subqueries para resolver este problema e depois tente reescrever usando JOINs. Compare as duas abordagens.

### Desafio 2: Ranking e Percentis

Crie uma query que mostre:
- Para cada categoria, o livro que está no percentil 75 de empréstimos (ou seja, tem mais empréstimos que 75% dos livros da categoria)
- Use subqueries correlacionadas para calcular o ranking

### Desafio 3: Análise Temporal

Crie uma query que identifique:
- Usuários que pegaram livros emprestados em meses diferentes
- E que pegaram mais livros que a média de empréstimos por usuário
- Use subqueries aninhadas para resolver

### Desafio 4: Otimização

Pegue uma das queries complexas que você criou nos exercícios anteriores e:
1. Analise sua performance (tempo de execução)
2. Identifique gargalos
3. Reescreva usando diferentes abordagens (subqueries, JOINs, tabelas derivadas)
4. Compare os resultados e performance
5. Documente qual abordagem é melhor e por quê

---

## Checklist de Aprendizado

Antes de prosseguir para a próxima aula, certifique-se de que você:

- [ ] Consegue criar subqueries escalares e usá-las em SELECT e WHERE
- [ ] Entende a diferença entre IN, NOT IN, EXISTS e quando usar cada um
- [ ] Consegue criar table subqueries (subqueries em FROM)
- [ ] Entende o que são correlated subqueries e quando usá-las
- [ ] Consegue criar nested subqueries (subqueries aninhadas)
- [ ] Sabe quando usar subqueries vs JOINs
- [ ] Entende o impacto de subqueries na performance
- [ ] Consegue debugar e otimizar subqueries
- [ ] Consegue resolver problemas complexos usando subqueries
- [ ] Entende as limitações e armadilhas comuns de subqueries

---

**⚠️ IMPORTANTE**: Complete todos os exercícios e responda todas as perguntas de reflexão antes de prosseguir para a próxima seção (Performance e Boas Práticas).

**Bons estudos! 🚀**





