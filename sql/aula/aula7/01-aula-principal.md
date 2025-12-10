# Aula 7: Sub Queries (Subconsultas)

## Introdução

Nesta aula, você aprenderá sobre **Sub Queries** (Subconsultas), também conhecidas como queries aninhadas ou consultas internas. Subqueries são queries SQL embutidas dentro de outra query, permitindo realizar consultas complexas, criar critérios dinâmicos e resolver problemas que seriam difíceis ou impossíveis de resolver com uma única query.

Subqueries são essenciais para:
- Criar critérios de filtro dinâmicos baseados em outros dados
- Realizar cálculos e comparações complexas
- Resolver problemas que requerem múltiplas etapas de consulta
- Trabalhar com dados agregados em condições de filtro
- Criar consultas mais legíveis e organizadas
- Resolver problemas que são difíceis de resolver apenas com JOINs

Dominar subqueries é fundamental para qualquer desenvolvedor ou analista de dados, pois permite escrever queries mais poderosas e expressivas, especialmente em situações onde JOINs não são suficientes ou não são a melhor solução.

---

## 1. O que são Sub Queries?

**Sub Query** (Subconsulta) é uma query SQL que é executada dentro de outra query SQL. A subquery é escrita entre parênteses e pode ser usada em várias partes de uma query principal, como SELECT, FROM, WHERE e HAVING.

### Características das Subqueries

1. **Executadas Primeiro**: Geralmente, a subquery é executada antes da query principal
2. **Retornam Resultados**: Podem retornar um único valor, uma coluna, uma linha ou uma tabela completa
3. **Isoladas**: A subquery é independente e pode ser testada separadamente
4. **Dinâmicas**: Podem usar valores da query externa (correlated subqueries)
5. **Flexíveis**: Podem ser usadas em diferentes contextos (SELECT, FROM, WHERE, etc.)

### Estrutura Básica

```sql
SELECT colunas
FROM tabela
WHERE coluna OPERADOR (
    SELECT coluna
    FROM outra_tabela
    WHERE condição
);
```

### Exemplo Simples

**Problema**: Encontrar todos os livros que têm mais empréstimos que a média de empréstimos de todos os livros.

**Solução com Subquery**:

```sql
SELECT titulo
FROM livros
WHERE (
    SELECT COUNT(*) 
    FROM emprestimos 
    WHERE emprestimos.livro_id = livros.id
) > (
    SELECT AVG(total_emprestimos)
    FROM (
        SELECT livro_id, COUNT(*) AS total_emprestimos
        FROM emprestimos
        GROUP BY livro_id
    )
);
```

**O que acontece**:
1. A subquery interna calcula a média de empréstimos por livro
2. A subquery externa conta os empréstimos de cada livro
3. A query principal compara e retorna apenas livros acima da média

### Por que Usar Subqueries?

**Sem Subquery** (usando JOIN e GROUP BY):
```sql
SELECT l.titulo
FROM livros l
JOIN emprestimos e ON l.id = e.livro_id
GROUP BY l.id, l.titulo
HAVING COUNT(e.id) > (
    SELECT AVG(total) 
    FROM (
        SELECT COUNT(*) AS total 
        FROM emprestimos 
        GROUP BY livro_id
    )
);
```

**Com Subquery** (mais legível):
```sql
SELECT titulo
FROM livros
WHERE (
    SELECT COUNT(*) 
    FROM emprestimos 
    WHERE emprestimos.livro_id = livros.id
) > (
    SELECT AVG(total_emprestimos)
    FROM (
        SELECT COUNT(*) AS total_emprestimos
        FROM emprestimos
        GROUP BY livro_id
    )
);
```

---

## 2. Tipos de Sub Queries

Subqueries podem retornar diferentes tipos de resultados. O tipo de resultado determina como a subquery pode ser usada na query principal.

### 2.1 Scalar Subquery (Subconsulta Escalar)

Uma **Scalar Subquery** retorna exatamente **um único valor** (uma coluna e uma linha). É o tipo mais comum e pode ser usada em qualquer lugar onde um valor único é esperado.

#### Características

- Retorna um único valor
- Pode ser usada em SELECT, WHERE, HAVING
- Deve retornar exatamente uma linha
- Se retornar zero linhas, o resultado é NULL
- Se retornar múltiplas linhas, causa erro

#### Sintaxe

```sql
SELECT coluna, (SELECT valor FROM tabela WHERE condição) AS subquery_result
FROM tabela;
```

#### Exemplos Práticos

**Exemplo 1: Subquery no SELECT**

```sql
-- Listar livros com o total de empréstimos de cada um
SELECT 
    titulo,
    (SELECT COUNT(*) 
     FROM emprestimos 
     WHERE emprestimos.livro_id = livros.id) AS total_emprestimos
FROM livros;
```

**Resultado**:
```
titulo              | total_emprestimos
Fundaçao            | 3
1984                | 2
Dom Casmurro        | 1
...
```

**Exemplo 2: Subquery em WHERE com Operadores de Comparação**

```sql
-- Encontrar livros com estoque acima da média
SELECT titulo, quantidade_disponivel
FROM livros
WHERE quantidade_disponivel > (
    SELECT AVG(quantidade_disponivel) 
    FROM livros
);
```

**Exemplo 3: Subquery Escalar com Agregação**

```sql
-- Encontrar o livro mais emprestado
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

**Exemplo 4: Múltiplas Subqueries Escalares**

```sql
-- Comparar estoque de cada livro com média e máximo
SELECT 
    titulo,
    quantidade_disponivel,
    (SELECT AVG(quantidade_disponivel) FROM livros) AS media_estoque,
    (SELECT MAX(quantidade_disponivel) FROM livros) AS max_estoque
FROM livros;
```

#### Erros Comuns com Scalar Subqueries

**Erro 1: Subquery retorna múltiplas linhas**
```sql
-- ❌ ERRO: Subquery retorna múltiplos valores
SELECT titulo
FROM livros
WHERE quantidade_disponivel = (
    SELECT quantidade_disponivel FROM livros  -- Retorna várias linhas!
);
```

**Solução**: Use operadores adequados (IN, ANY, ALL) ou adicione LIMIT 1

**Erro 2: Subquery retorna zero linhas**
```sql
-- Retorna NULL se não houver empréstimos
SELECT titulo,
       (SELECT COUNT(*) 
        FROM emprestimos 
        WHERE emprestimos.livro_id = livros.id 
        AND status = 'inexistente') AS total  -- Pode retornar NULL
FROM livros;
```

---

### 2.2 Column Subquery (Subconsulta de Coluna)

Uma **Column Subquery** retorna **uma coluna de valores** (múltiplas linhas, uma coluna). É usada principalmente com operadores como IN, NOT IN, ANY, ALL.

#### Características

- Retorna uma coluna com múltiplas linhas
- Usada principalmente em WHERE com operadores especiais
- Pode retornar zero, uma ou múltiplas linhas
- Não pode ser usada diretamente em SELECT (a menos que seja escalar)

#### Sintaxe

```sql
SELECT colunas
FROM tabela
WHERE coluna IN (SELECT coluna FROM outra_tabela WHERE condição);
```

#### Operadores para Column Subqueries

- **IN**: O valor está na lista retornada pela subquery
- **NOT IN**: O valor não está na lista retornada pela subquery
- **ANY/SOME**: O valor satisfaz a condição para pelo menos um valor da subquery
- **ALL**: O valor satisfaz a condição para todos os valores da subquery
- **EXISTS**: Verifica se a subquery retorna pelo menos uma linha

#### Exemplos Práticos

**Exemplo 1: IN - Encontrar livros de autores brasileiros**

```sql
SELECT titulo, autor_id
FROM livros
WHERE autor_id IN (
    SELECT id 
    FROM autores 
    WHERE nacionalidade = 'Brasileiro'
);
```

**Exemplo 2: NOT IN - Encontrar livros que nunca foram emprestados**

```sql
SELECT titulo
FROM livros
WHERE id NOT IN (
    SELECT DISTINCT livro_id 
    FROM emprestimos
    WHERE livro_id IS NOT NULL
);
```

**Exemplo 3: ANY - Encontrar livros com estoque maior que qualquer livro de ficção científica**

```sql
SELECT titulo, quantidade_disponivel
FROM livros
WHERE quantidade_disponivel > ANY (
    SELECT quantidade_disponivel
    FROM livros
    WHERE categoria_id = (
        SELECT id FROM categorias WHERE nome = 'Ficção Científica'
    )
);
```

**Exemplo 4: ALL - Encontrar livros com estoque maior que todos os livros de uma categoria**

```sql
SELECT titulo, quantidade_disponivel
FROM livros
WHERE quantidade_disponivel > ALL (
    SELECT quantidade_disponivel
    FROM livros
    WHERE categoria_id = (
        SELECT id FROM categorias WHERE nome = 'Romance'
    )
);
```

**Exemplo 5: EXISTS - Verificar se há empréstimos ativos**

```sql
SELECT titulo
FROM livros
WHERE EXISTS (
    SELECT 1
    FROM emprestimos
    WHERE emprestimos.livro_id = livros.id
    AND emprestimos.status = 'ativo'
);
```

#### Diferença entre IN e EXISTS

**IN**: Retorna todos os valores e depois verifica se estão na lista
```sql
SELECT * FROM livros
WHERE autor_id IN (SELECT id FROM autores WHERE nacionalidade = 'Brasileiro');
```

**EXISTS**: Para na primeira correspondência (pode ser mais eficiente)
```sql
SELECT * FROM livros l
WHERE EXISTS (
    SELECT 1 FROM autores a 
    WHERE a.id = l.autor_id 
    AND a.nacionalidade = 'Brasileiro'
);
```

---

### 2.3 Row Subquery (Subconsulta de Linha)

Uma **Row Subquery** retorna **uma linha com múltiplas colunas**. É usada com operadores de comparação de linhas.

#### Características

- Retorna uma linha com múltiplas colunas
- Usada com operadores de comparação de tuplas
- Menos comum que scalar e column subqueries
- Útil para comparações complexas

#### Sintaxe

```sql
SELECT colunas
FROM tabela
WHERE (coluna1, coluna2) = (SELECT coluna1, coluna2 FROM outra_tabela WHERE condição);
```

#### Exemplos Práticos

**Exemplo 1: Comparar múltiplas colunas**

```sql
-- Encontrar livros com mesmo autor e categoria que um livro específico
SELECT titulo
FROM livros
WHERE (autor_id, categoria_id) = (
    SELECT autor_id, categoria_id
    FROM livros
    WHERE titulo = 'Fundação'
);
```

**Exemplo 2: Comparar com múltiplas condições**

```sql
-- Encontrar livros publicados no mesmo ano e pela mesma editora
SELECT titulo
FROM livros
WHERE (ano_publicacao, editora) IN (
    SELECT ano_publicacao, editora
    FROM livros
    WHERE ano_publicacao IS NOT NULL
    AND editora IS NOT NULL
    GROUP BY ano_publicacao, editora
    HAVING COUNT(*) > 1
);
```

---

### 2.4 Table Subquery (Subconsulta de Tabela)

Uma **Table Subquery** retorna **uma tabela completa** (múltiplas linhas e múltiplas colunas). É usada principalmente na cláusula FROM como uma tabela virtual (derived table).

#### Características

- Retorna uma tabela completa
- Usada principalmente em FROM
- Pode ser usada como uma tabela temporária
- Útil para simplificar queries complexas
- Deve ter um alias

#### Sintaxe

```sql
SELECT colunas
FROM (
    SELECT colunas
    FROM tabela
    WHERE condição
) AS alias;
```

#### Exemplos Práticos

**Exemplo 1: Subquery em FROM - Estatísticas por categoria**

```sql
-- Calcular estatísticas de empréstimos por categoria
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
```

**Exemplo 2: Subquery em FROM - Top N por categoria**

```sql
-- Encontrar os 3 livros mais emprestados de cada categoria
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
```

**Exemplo 3: Subquery em FROM - Agregações aninhadas**

```sql
-- Calcular média de empréstimos por autor
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
```

---

## 3. Subqueries em Diferentes Cláusulas

Subqueries podem ser usadas em várias partes de uma query SQL. Cada contexto tem suas próprias regras e limitações.

### 3.1 Subqueries em SELECT

Subqueries no SELECT devem retornar valores escalares (um único valor por linha).

#### Características

- Devem retornar um único valor (escalar)
- Executadas para cada linha do resultado
- Podem ser correlated (usar valores da query externa)
- Podem impactar performance se mal utilizadas

#### Exemplos

**Exemplo 1: Contar empréstimos por livro**

```sql
SELECT 
    titulo,
    (SELECT COUNT(*) 
     FROM emprestimos 
     WHERE emprestimos.livro_id = livros.id) AS total_emprestimos
FROM livros;
```

**Exemplo 2: Comparar com média**

```sql
SELECT 
    titulo,
    quantidade_disponivel,
    (SELECT AVG(quantidade_disponivel) FROM livros) AS media_estoque,
    CASE 
        WHEN quantidade_disponivel > (SELECT AVG(quantidade_disponivel) FROM livros) 
        THEN 'Acima da média'
        ELSE 'Abaixo da média'
    END AS status_estoque
FROM livros;
```

---

### 3.2 Subqueries em WHERE

Subqueries em WHERE são muito comuns e podem retornar valores escalares, colunas ou linhas, dependendo do operador usado.

#### Características

- Podem retornar escalar, coluna ou linha
- Usadas para criar filtros dinâmicos
- Podem ser correlated ou não-correlated
- Muito flexíveis e poderosas

#### Exemplos

**Exemplo 1: Comparação com valor escalar**

```sql
-- Livros com estoque acima da média
SELECT titulo, quantidade_disponivel
FROM livros
WHERE quantidade_disponivel > (
    SELECT AVG(quantidade_disponivel) FROM livros
);
```

**Exemplo 2: Usando IN**

```sql
-- Livros de autores brasileiros
SELECT titulo
FROM livros
WHERE autor_id IN (
    SELECT id FROM autores WHERE nacionalidade = 'Brasileiro'
);
```

**Exemplo 3: Usando EXISTS**

```sql
-- Livros que têm empréstimos ativos
SELECT titulo
FROM livros
WHERE EXISTS (
    SELECT 1
    FROM emprestimos
    WHERE emprestimos.livro_id = livros.id
    AND emprestimos.status = 'ativo'
);
```

---

### 3.3 Subqueries em FROM

Subqueries em FROM criam tabelas derivadas (derived tables) que podem ser usadas como tabelas normais.

#### Características

- Devem retornar uma tabela completa
- Devem ter um alias
- Podem ser usadas em JOINs
- Úteis para simplificar queries complexas

#### Exemplos

**Exemplo 1: Tabela derivada simples**

```sql
SELECT 
    categoria,
    COUNT(*) AS total_livros
FROM (
    SELECT 
        c.nome AS categoria,
        l.titulo
    FROM categorias c
    JOIN livros l ON c.id = l.categoria_id
) AS livros_por_categoria
GROUP BY categoria;
```

**Exemplo 2: JOIN com tabela derivada**

```sql
SELECT 
    l.titulo,
    stats.total_emprestimos
FROM livros l
JOIN (
    SELECT 
        livro_id,
        COUNT(*) AS total_emprestimos
    FROM emprestimos
    GROUP BY livro_id
) AS stats ON l.id = stats.livro_id;
```

---

### 3.4 Subqueries em HAVING

Subqueries em HAVING são usadas para filtrar grupos baseado em condições que envolvem outras consultas.

#### Características

- Usadas após GROUP BY
- Podem retornar escalar, coluna ou linha
- Úteis para comparações com agregados de outras consultas
- Menos comuns que WHERE

#### Exemplos

**Exemplo 1: Comparar grupo com média geral**

```sql
-- Categorias com média de estoque acima da média geral
SELECT 
    c.nome AS categoria,
    AVG(l.quantidade_disponivel) AS media_estoque
FROM categorias c
JOIN livros l ON c.id = l.categoria_id
GROUP BY c.id, c.nome
HAVING AVG(l.quantidade_disponivel) > (
    SELECT AVG(quantidade_disponivel) FROM livros
);
```

**Exemplo 2: Comparar com outro grupo**

```sql
-- Autores com mais livros que a média de livros por autor
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
```

---

## 4. Nested Subqueries (Subqueries Aninhadas)

**Nested Subqueries** são subqueries que contêm outras subqueries dentro delas. Elas permitem resolver problemas muito complexos, mas podem se tornar difíceis de ler e manter.

### Características

- Subqueries dentro de subqueries
- Podem ter múltiplos níveis de aninhamento
- Podem ser difíceis de entender e manter
- Às vezes podem ser reescritas como JOINs ou CTEs (Common Table Expressions)

### Exemplos Práticos

**Exemplo 1: Dois Níveis de Aninhamento**

```sql
-- Encontrar livros de autores que têm mais livros que a média
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
```

**Exemplo 2: Múltiplos Níveis com Diferentes Tipos**

```sql
-- Encontrar categorias onde a média de empréstimos por livro é maior que a média geral
SELECT nome
FROM categorias
WHERE id IN (
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
```

### Quando Usar Nested Subqueries

- Quando a lógica é complexa e requer múltiplas etapas
- Quando você precisa de resultados intermediários
- Quando JOINs não são suficientes ou são muito complexos

### Alternativas

Muitas vezes, nested subqueries podem ser reescritas usando:
- **JOINs**: Mais eficientes e legíveis
- **CTEs (Common Table Expressions)**: Mais legíveis (disponível em alguns SGBDs)
- **Views**: Para reutilização

---

## 5. Correlated Subqueries (Subqueries Correlacionadas)

Uma **Correlated Subquery** é uma subquery que referencia colunas da query externa (outer query). Diferente de subqueries normais, correlated subqueries são executadas uma vez para cada linha processada pela query externa.

### Características

- Referencia colunas da query externa
- Executada uma vez para cada linha da query externa
- Pode ser muito lenta em tabelas grandes
- Útil quando você precisa comparar cada linha com resultados agregados

### Diferença entre Subquery Normal e Correlated Subquery

**Subquery Normal (Não-Correlacionada)**:
```sql
-- Executada UMA vez
SELECT titulo
FROM livros
WHERE quantidade_disponivel > (
    SELECT AVG(quantidade_disponivel) FROM livros  -- Não usa coluna externa
);
```

**Correlated Subquery**:
```sql
-- Executada para CADA linha
SELECT titulo
FROM livros l1
WHERE quantidade_disponivel > (
    SELECT AVG(quantidade_disponivel) 
    FROM livros l2
    WHERE l2.categoria_id = l1.categoria_id  -- Usa coluna da query externa (l1)
);
```

### Exemplos Práticos

**Exemplo 1: Comparar com Média do Grupo**

```sql
-- Encontrar livros com estoque acima da média de sua categoria
SELECT 
    l1.titulo,
    l1.quantidade_disponivel,
    c.nome AS categoria
FROM livros l1
JOIN categorias c ON l1.categoria_id = c.id
WHERE l1.quantidade_disponivel > (
    SELECT AVG(l2.quantidade_disponivel)
    FROM livros l2
    WHERE l2.categoria_id = l1.categoria_id  -- Correlação: usa l1.categoria_id
);
```

**Exemplo 2: Contar Relacionados**

```sql
-- Listar livros com o número de empréstimos de cada um
SELECT 
    titulo,
    (SELECT COUNT(*)
     FROM emprestimos e
     WHERE e.livro_id = livros.id) AS total_emprestimos
FROM livros;
```

**Exemplo 3: Verificar Existência**

```sql
-- Encontrar autores que têm livros emprestados
SELECT nome
FROM autores a
WHERE EXISTS (
    SELECT 1
    FROM livros l
    JOIN emprestimos e ON l.id = e.livro_id
    WHERE l.autor_id = a.id  -- Correlação: usa a.id
    AND e.status = 'ativo'
);
```

**Exemplo 4: Comparar com Máximo do Grupo**

```sql
-- Encontrar o livro mais emprestado de cada categoria
SELECT 
    l1.titulo,
    c.nome AS categoria,
    (SELECT COUNT(*)
     FROM emprestimos e
     WHERE e.livro_id = l1.id) AS total_emprestimos
FROM livros l1
JOIN categorias c ON l1.categoria_id = c.id
WHERE (SELECT COUNT(*)
       FROM emprestimos e
       WHERE e.livro_id = l1.id) = (
    SELECT MAX(total)
    FROM (
        SELECT COUNT(*) AS total
        FROM emprestimos e2
        JOIN livros l2 ON e2.livro_id = l2.id
        WHERE l2.categoria_id = l1.categoria_id  -- Correlação
        GROUP BY l2.id
    )
);
```

### Performance de Correlated Subqueries

**⚠️ ATENÇÃO**: Correlated subqueries podem ser muito lentas porque são executadas uma vez para cada linha da query externa.

**Exemplo Lento**:
```sql
-- Executa a subquery para CADA livro (pode ser muito lento)
SELECT titulo
FROM livros l1
WHERE (
    SELECT COUNT(*)
    FROM emprestimos e
    WHERE e.livro_id = l1.id
) > 5;
```

**Alternativa Mais Rápida (usando JOIN)**:
```sql
-- Executa uma única vez (muito mais rápido)
SELECT l.titulo
FROM livros l
JOIN (
    SELECT livro_id, COUNT(*) AS total
    FROM emprestimos
    GROUP BY livro_id
    HAVING COUNT(*) > 5
) AS stats ON l.id = stats.livro_id;
```

### Quando Usar Correlated Subqueries

Use correlated subqueries quando:
- Você precisa comparar cada linha com resultados específicos do seu grupo
- A lógica é mais clara com correlated subquery
- A tabela é pequena ou a performance não é crítica
- Não há uma alternativa eficiente com JOIN

Evite correlated subqueries quando:
- A tabela é muito grande
- Performance é crítica
- Você pode resolver com JOIN de forma mais eficiente

---

## 6. Subqueries vs JOINs

Muitas vezes, o mesmo problema pode ser resolvido tanto com subqueries quanto com JOINs. É importante entender quando usar cada abordagem.

### Quando Usar Subqueries

1. **Filtros Dinâmicos Complexos**: Quando você precisa filtrar baseado em cálculos ou agregações
2. **Comparações com Agregados**: Quando você compara valores com médias, máximos, etc.
3. **Existência**: Quando você só precisa verificar se algo existe (EXISTS)
4. **Legibilidade**: Quando a subquery torna o código mais legível
5. **Valores Únicos**: Quando você precisa de um único valor calculado

**Exemplo - Melhor com Subquery**:
```sql
-- Livros com estoque acima da média
SELECT titulo
FROM livros
WHERE quantidade_disponivel > (SELECT AVG(quantidade_disponivel) FROM livros);
```

### Quando Usar JOINs

1. **Combinar Dados**: Quando você precisa combinar dados de múltiplas tabelas
2. **Performance**: Quando JOINs são mais eficientes
3. **Múltiplas Colunas**: Quando você precisa de várias colunas de tabelas relacionadas
4. **Agrupamentos**: Quando você precisa agrupar dados de múltiplas tabelas
5. **Relacionamentos Diretos**: Quando há relacionamentos claros entre tabelas

**Exemplo - Melhor com JOIN**:
```sql
-- Listar livros com seus autores
SELECT l.titulo, a.nome AS autor
FROM livros l
JOIN autores a ON l.autor_id = a.id;
```

### Comparação: Mesmo Problema, Diferentes Abordagens

**Problema**: Encontrar livros que têm empréstimos ativos

**Solução 1: Subquery com EXISTS**
```sql
SELECT titulo
FROM livros
WHERE EXISTS (
    SELECT 1
    FROM emprestimos
    WHERE emprestimos.livro_id = livros.id
    AND emprestimos.status = 'ativo'
);
```

**Solução 2: JOIN**
```sql
SELECT DISTINCT l.titulo
FROM livros l
JOIN emprestimos e ON l.id = e.livro_id
WHERE e.status = 'ativo';
```

**Qual usar?**
- **EXISTS**: Geralmente mais eficiente, para apenas verificar existência
- **JOIN**: Mais flexível, se você precisar de dados do empréstimo também

### Regra Geral

- **Subqueries**: Para filtros, comparações e valores únicos
- **JOINs**: Para combinar e relacionar dados de múltiplas tabelas
- **Ambos**: Muitas vezes você pode usar ambos, escolha baseado em legibilidade e performance

---

## 7. Exemplos Práticos Avançados

Vamos ver alguns exemplos práticos que combinam diferentes conceitos de subqueries.

### Exemplo 1: Análise Complexa de Empréstimos

```sql
-- Encontrar usuários que pegaram emprestado mais livros que a média
SELECT 
    u.nome,
    COUNT(e.id) AS total_emprestimos
FROM usuarios u
JOIN emprestimos e ON u.id = e.usuario_id
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

### Exemplo 2: Comparação com Múltiplos Níveis

```sql
-- Encontrar categorias onde todos os livros têm estoque acima da média geral
SELECT c.nome AS categoria
FROM categorias c
WHERE NOT EXISTS (
    SELECT 1
    FROM livros l
    WHERE l.categoria_id = c.id
    AND l.quantidade_disponivel <= (
        SELECT AVG(quantidade_disponivel) FROM livros
    )
);
```

### Exemplo 3: Ranking com Subqueries

```sql
-- Listar livros com seu ranking de empréstimos dentro da categoria
SELECT 
    l.titulo,
    c.nome AS categoria,
    (SELECT COUNT(*)
     FROM emprestimos e
     WHERE e.livro_id = l.id) AS total_emprestimos,
    (SELECT COUNT(*)
     FROM livros l2
     JOIN emprestimos e2 ON l2.id = e2.livro_id
     WHERE l2.categoria_id = l.categoria_id
     AND (SELECT COUNT(*) FROM emprestimos WHERE livro_id = l2.id) > 
         (SELECT COUNT(*) FROM emprestimos WHERE livro_id = l.id)
    ) + 1 AS ranking_categoria
FROM livros l
JOIN categorias c ON l.categoria_id = c.id
ORDER BY c.nome, total_emprestimos DESC;
```

---

## 8. Limitações e Considerações

### Limitações de Subqueries

1. **Performance**: Subqueries podem ser lentas, especialmente correlated subqueries
2. **Legibilidade**: Subqueries aninhadas podem ser difíceis de ler
3. **Manutenibilidade**: Queries complexas com muitas subqueries são difíceis de manter
4. **Limitações do SQLite**: Algumas funcionalidades avançadas podem não estar disponíveis

### Boas Práticas

1. **Teste Subqueries Separadamente**: Sempre teste a subquery sozinha primeiro
2. **Use Aliases Claros**: Dê nomes descritivos às subqueries e aliases
3. **Evite Aninhamento Excessivo**: Se possível, reescreva com JOINs ou CTEs
4. **Considere Performance**: Para tabelas grandes, prefira JOINs quando possível
5. **Documente Lógica Complexa**: Adicione comentários para subqueries complexas

### Quando Revisar Subqueries

Revisite suas subqueries se:
- A query está lenta
- O código é difícil de entender
- Você precisa adicionar mais funcionalidades
- Há erros ou resultados inesperados

---

## Conclusão

Subqueries são uma ferramenta poderosa do SQL que permite resolver problemas complexos de forma elegante. Elas são essenciais para:

- Criar filtros dinâmicos
- Realizar comparações com agregados
- Verificar existência de dados
- Simplificar queries complexas
- Trabalhar com dados relacionados de forma flexível

Dominar subqueries, junto com JOINs e aggregate queries, permite criar consultas sofisticadas e resolver problemas de análise de dados complexos.

**Próximos Passos**:
1. Pratique os exemplos desta aula no banco de dados
2. Complete os exercícios práticos
3. Estude as boas práticas de performance
4. Experimente reescrever subqueries como JOINs e vice-versa

**Lembre-se**: Subqueries são poderosas, mas nem sempre são a melhor solução. Considere sempre legibilidade, performance e manutenibilidade ao escolher entre subqueries e JOINs.





