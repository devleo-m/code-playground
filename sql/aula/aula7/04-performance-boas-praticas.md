# Aula 7 - Performance, Boas Práticas e Otimização: Sub Queries

## Introdução

Subqueries são poderosas e flexíveis, mas podem ter grande impacto na performance se não forem usadas corretamente. Nesta seção, você aprenderá como otimizar subqueries, quando usar cada tipo, e como evitar problemas comuns que tornam queries lentas.

**Regra de Ouro**: Subqueries bem otimizadas são eficientes. Subqueries mal otimizadas, especialmente correlated subqueries, podem ser extremamente lentas, especialmente em tabelas grandes.

---

## 1. Impacto de Subqueries na Performance

### Por que Subqueries Podem Ser Lentas?

Subqueries podem ser computacionalmente custosas por várias razões:

1. **Execução Múltipla**: Correlated subqueries são executadas uma vez para cada linha da query externa
2. **Sem Índices**: Subqueries sem índices nas colunas usadas são muito lentas
3. **Full Table Scans**: Subqueries podem forçar o banco a examinar todas as linhas
4. **Aninhamento Excessivo**: Múltiplos níveis de aninhamento aumentam a complexidade
5. **Resultados Grandes**: Subqueries que retornam muitos dados podem ser lentas

### Exemplo de Impacto

```sql
-- ❌ MUITO LENTO: Correlated subquery sem índice
SELECT titulo
FROM livros l1
WHERE (
    SELECT COUNT(*)
    FROM emprestimos e
    WHERE e.livro_id = l1.id  -- Executada para CADA livro!
) > 5;
-- Se livros tem 10.000 linhas:
-- A subquery é executada 10.000 vezes!
-- Se cada execução leva 10ms: 100 segundos total!

-- ✅ RÁPIDO: JOIN com agregação
SELECT l.titulo
FROM livros l
JOIN (
    SELECT livro_id, COUNT(*) AS total
    FROM emprestimos
    GROUP BY livro_id
    HAVING COUNT(*) > 5
) AS stats ON l.id = stats.livro_id;
-- Executa uma única vez: ~100ms total
```

**Diferença**: De minutos para milissegundos!

---

## 2. Correlated Subqueries: O Maior Problema de Performance

### Por que Correlated Subqueries São Lentas?

Correlated subqueries são executadas **uma vez para cada linha** da query externa. Isso pode resultar em milhares ou milhões de execuções!

```sql
-- ❌ PROBLEMA: Executada 10.000 vezes (uma para cada livro)
SELECT titulo
FROM livros l1
WHERE quantidade_disponivel > (
    SELECT AVG(quantidade_disponivel)
    FROM livros l2
    WHERE l2.categoria_id = l1.categoria_id  -- Correlação!
);
```

### Quando Correlated Subqueries São Aceitáveis?

Correlated subqueries são aceitáveis quando:

1. **Tabela Externa é Pequena**: Poucas linhas = poucas execuções
2. **Subquery é Rápida**: Usa índices e retorna rapidamente
3. **Não Há Alternativa Eficiente**: JOINs não resolvem o problema de forma melhor
4. **Performance Não é Crítica**: Em relatórios não-frequentes

### Alternativas para Correlated Subqueries

**Problema Original:**
```sql
-- Correlated subquery lenta
SELECT titulo
FROM livros l1
WHERE quantidade_disponivel > (
    SELECT AVG(quantidade_disponivel)
    FROM livros l2
    WHERE l2.categoria_id = l1.categoria_id
);
```

**Solução 1: JOIN com Tabela Derivada**
```sql
-- ✅ MUITO MAIS RÁPIDO
SELECT l1.titulo
FROM livros l1
JOIN (
    SELECT categoria_id, AVG(quantidade_disponivel) AS media
    FROM livros
    GROUP BY categoria_id
) AS medias ON l1.categoria_id = medias.categoria_id
WHERE l1.quantidade_disponivel > medias.media;
```

**Solução 2: Window Functions (se disponível)**
```sql
-- ✅ AINDA MAIS EFICIENTE (se o SGBD suportar)
SELECT titulo
FROM (
    SELECT 
        titulo,
        quantidade_disponivel,
        AVG(quantidade_disponivel) OVER (PARTITION BY categoria_id) AS media_categoria
    FROM livros
) AS ranked
WHERE quantidade_disponivel > media_categoria;
```

### Regra de Ouro para Correlated Subqueries

**SEMPRE** tente reescrever correlated subqueries como JOINs ou tabelas derivadas. Apenas use correlated subqueries se:
- A tabela externa é muito pequena (< 100 linhas)
- Não há alternativa eficiente
- Performance não é crítica

---

## 3. Índices e Subqueries: A Chave para Performance

### Por que Índices São Essenciais?

Índices são **fundamentais** para performance de subqueries. Sem índices, o banco precisa fazer "full table scan" (examinar todas as linhas) em cada execução da subquery.

### Índices em Colunas de Subqueries

**Regra de Ouro**: Sempre tenha índices nas colunas usadas em condições de subqueries, especialmente em correlated subqueries.

```sql
-- ✅ BOM: Coluna usada na subquery tem índice
CREATE INDEX idx_emprestimos_livro ON emprestimos(livro_id);
SELECT titulo
FROM livros l
WHERE EXISTS (
    SELECT 1
    FROM emprestimos e
    WHERE e.livro_id = l.id  -- Usa índice!
    AND e.status = 'ativo'
);

-- ❌ RUIM: Coluna sem índice
SELECT titulo
FROM livros l
WHERE EXISTS (
    SELECT 1
    FROM emprestimos e
    WHERE e.livro_id = l.id  -- Full table scan em cada execução!
    AND e.status = 'ativo'
);
```

### Verificando Índices Existentes

```sql
-- Ver todos os índices
SELECT * FROM sqlite_master WHERE type='index';

-- Ver índices de uma tabela específica
SELECT * FROM sqlite_master 
WHERE type='index' AND tbl_name='emprestimos';

-- Ver estrutura de uma tabela
.schema emprestimos
```

### Criando Índices para Subqueries

```sql
-- Índices essenciais para subqueries comuns
CREATE INDEX IF NOT EXISTS idx_emprestimos_livro ON emprestimos(livro_id);
CREATE INDEX IF NOT EXISTS idx_emprestimos_usuario ON emprestimos(usuario_id);
CREATE INDEX IF NOT EXISTS idx_emprestimos_status ON emprestimos(status);
CREATE INDEX IF NOT EXISTS idx_livros_autor ON livros(autor_id);
CREATE INDEX IF NOT EXISTS idx_livros_categoria ON livros(categoria_id);

-- Índices compostos para subqueries com múltiplas condições
CREATE INDEX IF NOT EXISTS idx_emprestimos_livro_status 
ON emprestimos(livro_id, status);

-- Verificar se índices estão sendo usados
EXPLAIN QUERY PLAN
SELECT titulo
FROM livros l
WHERE EXISTS (
    SELECT 1
    FROM emprestimos e
    WHERE e.livro_id = l.id
    AND e.status = 'ativo'
);
```

---

## 4. Subqueries vs JOINs: Escolhendo a Abordagem Mais Eficiente

### Quando Subqueries São Mais Eficientes?

Subqueries são mais eficientes quando:

1. **Verificação de Existência**: EXISTS é geralmente mais eficiente que JOIN para verificar existência
2. **Valores Únicos**: Quando você precisa de um único valor calculado
3. **Filtros Complexos**: Quando o filtro é complexo e difícil de expressar com JOIN
4. **Tabela Externa Pequena**: Quando a query externa retorna poucas linhas

**Exemplo: EXISTS vs JOIN**
```sql
-- ✅ EXISTS: Para apenas quando encontra (pode ser mais rápido)
SELECT titulo
FROM livros l
WHERE EXISTS (
    SELECT 1
    FROM emprestimos e
    WHERE e.livro_id = l.id
    AND e.status = 'ativo'
);

-- JOIN: Combina todos os dados (pode ser mais lento)
SELECT DISTINCT l.titulo
FROM livros l
JOIN emprestimos e ON l.id = e.livro_id
WHERE e.status = 'ativo';
```

### Quando JOINs São Mais Eficientes?

JOINs são mais eficientes quando:

1. **Combinar Dados**: Quando você precisa de dados de múltiplas tabelas
2. **Agregações**: Quando você precisa agregar dados de múltiplas tabelas
3. **Múltiplas Colunas**: Quando você precisa de várias colunas relacionadas
4. **Tabelas Grandes**: Quando ambas as tabelas são grandes

**Exemplo: JOIN vs Subquery**
```sql
-- ❌ Subquery: Executada para cada linha
SELECT 
    titulo,
    (SELECT COUNT(*) FROM emprestimos WHERE livro_id = livros.id) AS total
FROM livros;

-- ✅ JOIN: Executada uma única vez
SELECT 
    l.titulo,
    COALESCE(stats.total, 0) AS total
FROM livros l
LEFT JOIN (
    SELECT livro_id, COUNT(*) AS total
    FROM emprestimos
    GROUP BY livro_id
) AS stats ON l.id = stats.livro_id;
```

### Regra de Ouro

- **Subqueries**: Para filtros, comparações e valores únicos
- **JOINs**: Para combinar e relacionar dados de múltiplas tabelas
- **Teste ambos**: Quando em dúvida, teste ambas as abordagens e compare performance

---

## 5. Otimizando Subqueries Escalares

### Subqueries Escalares no SELECT

Subqueries escalares no SELECT são executadas **uma vez para cada linha** do resultado. Isso pode ser lento se a subquery for complexa.

```sql
-- ⚠️ ATENÇÃO: Executada para cada livro
SELECT 
    titulo,
    (SELECT COUNT(*) 
     FROM emprestimos 
     WHERE emprestimos.livro_id = livros.id) AS total_emprestimos
FROM livros;
```

**Otimização:**
```sql
-- ✅ MELHOR: JOIN com agregação (executada uma vez)
SELECT 
    l.titulo,
    COALESCE(stats.total_emprestimos, 0) AS total_emprestimos
FROM livros l
LEFT JOIN (
    SELECT livro_id, COUNT(*) AS total_emprestimos
    FROM emprestimos
    GROUP BY livro_id
) AS stats ON l.id = stats.livro_id;
```

### Subqueries Escalares em WHERE

Subqueries escalares em WHERE são executadas **uma vez** (não para cada linha), então são geralmente mais eficientes.

```sql
-- ✅ EFICIENTE: Executada uma única vez
SELECT titulo
FROM livros
WHERE quantidade_disponivel > (
    SELECT AVG(quantidade_disponivel) FROM livros
);
```

**Dica**: Se a subquery escalar for complexa, considere calcular o valor uma vez e armazená-lo em uma variável (se o SGBD suportar) ou usar uma tabela derivada.

---

## 6. Otimizando Subqueries com IN, NOT IN, EXISTS

### IN vs EXISTS: Qual é Mais Eficiente?

**IN**: Retorna todos os valores e depois verifica se estão na lista
```sql
SELECT titulo
FROM livros
WHERE autor_id IN (
    SELECT id FROM autores WHERE nacionalidade = 'Brasileiro'
);
```

**EXISTS**: Para na primeira correspondência (geralmente mais eficiente)
```sql
SELECT l.titulo
FROM livros l
WHERE EXISTS (
    SELECT 1 FROM autores a 
    WHERE a.id = l.autor_id 
    AND a.nacionalidade = 'Brasileiro'
);
```

**Regra Geral**: 
- **EXISTS** é geralmente mais eficiente, especialmente quando há muitas correspondências
- **IN** pode ser mais eficiente quando há poucas correspondências e a subquery é simples

### NOT IN vs NOT EXISTS: Cuidado com NULLs!

**NOT IN**: Pode ter comportamento inesperado com NULLs
```sql
-- ⚠️ PROBLEMA: Se a subquery retornar NULL, NOT IN pode não funcionar como esperado
SELECT titulo
FROM livros
WHERE autor_id NOT IN (
    SELECT id FROM autores WHERE nacionalidade = 'Americano'
    -- Se algum id for NULL, o resultado pode ser vazio!
);
```

**NOT EXISTS**: Mais seguro com NULLs
```sql
-- ✅ SEGURO: Funciona corretamente mesmo com NULLs
SELECT l.titulo
FROM livros l
WHERE NOT EXISTS (
    SELECT 1 FROM autores a 
    WHERE a.id = l.autor_id 
    AND a.nacionalidade = 'Americano'
);
```

**Regra de Ouro**: Use **NOT EXISTS** ao invés de **NOT IN** para evitar problemas com NULLs.

---

## 7. Otimizando Nested Subqueries

### O Problema do Aninhamento Excessivo

Nested subqueries (subqueries aninhadas) podem ser difíceis de otimizar e entender:

```sql
-- ❌ DIFÍCIL DE OTIMIZAR: Múltiplos níveis
SELECT titulo
FROM livros
WHERE autor_id IN (
    SELECT id FROM autores
    WHERE id IN (
        SELECT autor_id FROM livros
        GROUP BY autor_id
        HAVING COUNT(*) > (
            SELECT AVG(total) FROM (
                SELECT COUNT(*) AS total
                FROM livros
                GROUP BY autor_id
            )
        )
    )
);
```

### Estratégias de Otimização

**Estratégia 1: Quebrar em Múltiplas Queries**
```sql
-- Calcular média primeiro
-- (Em aplicação, armazenar resultado)
SELECT AVG(total) AS media
FROM (
    SELECT COUNT(*) AS total
    FROM livros
    GROUP BY autor_id
);

-- Usar resultado na query principal
SELECT titulo
FROM livros
WHERE autor_id IN (
    SELECT autor_id
    FROM livros
    GROUP BY autor_id
    HAVING COUNT(*) > 2.5  -- Valor calculado anteriormente
);
```

**Estratégia 2: Usar JOINs e Tabelas Derivadas**
```sql
-- ✅ MAIS EFICIENTE: JOINs e tabelas derivadas
SELECT l.titulo
FROM livros l
JOIN (
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
) AS autores_prolificos ON l.autor_id = autores_prolificos.autor_id;
```

**Estratégia 3: Usar CTEs (Common Table Expressions) - se disponível**
```sql
-- ✅ MAIS LEGÍVEL: CTEs (se o SGBD suportar)
WITH media_por_autor AS (
    SELECT autor_id, COUNT(*) AS total
    FROM livros
    GROUP BY autor_id
),
media_geral AS (
    SELECT AVG(total) AS media
    FROM media_por_autor
),
autores_prolificos AS (
    SELECT autor_id
    FROM media_por_autor
    CROSS JOIN media_geral
    WHERE total > media
)
SELECT l.titulo
FROM livros l
JOIN autores_prolificos ap ON l.autor_id = ap.autor_id;
```

---

## 8. Boas Práticas de Escrita de Subqueries

### 1. Teste Subqueries Separadamente

**SEMPRE** teste a subquery sozinha primeiro para garantir que ela retorna o que você espera:

```sql
-- ✅ BOM: Teste a subquery primeiro
-- Primeiro, teste:
SELECT AVG(quantidade_disponivel) FROM livros;
-- Resultado: 5.2 ✅

-- Depois use na query principal:
SELECT titulo FROM livros WHERE quantidade_disponivel > 5.2;
```

### 2. Use Aliases Claros

Dê nomes descritivos para facilitar a leitura e manutenção:

```sql
-- ✅ BOM: Aliases claros
SELECT l1.titulo
FROM livros l1
WHERE l1.quantidade_disponivel > (
    SELECT AVG(l2.quantidade_disponivel)
    FROM livros l2
    WHERE l2.categoria_id = l1.categoria_id
);

-- ❌ RUIM: Sem aliases ou confusos
SELECT titulo
FROM livros
WHERE quantidade_disponivel > (
    SELECT AVG(quantidade_disponivel)
    FROM livros
    WHERE categoria_id = livros.categoria_id
);
```

### 3. Evite Aninhamento Excessivo

Se você tem mais de 3 níveis de aninhamento, considere reescrever:

```sql
-- ❌ MUITO ANINHADO (difícil de entender e manter)
SELECT ... FROM ... WHERE ... IN (
    SELECT ... FROM ... WHERE ... IN (
        SELECT ... FROM ... WHERE ... IN (
            SELECT ... FROM ...
        )
    )
);

-- ✅ MELHOR: Use JOINs ou quebre em múltiplas queries
```

### 4. Documente Subqueries Complexas

Adicione comentários para subqueries complexas:

```sql
-- Encontrar livros com estoque acima da média de sua categoria
SELECT l1.titulo
FROM livros l1
WHERE l1.quantidade_disponivel > (
    -- Subquery: Calcula média de estoque por categoria
    SELECT AVG(l2.quantidade_disponivel)
    FROM livros l2
    WHERE l2.categoria_id = l1.categoria_id
);
```

### 5. Considere Performance desde o Início

Pense em performance ao escrever subqueries:

- Use índices nas colunas de subqueries
- Evite correlated subqueries quando possível
- Prefira JOINs para combinar dados
- Teste performance de queries complexas

---

## 9. Troubleshooting de Subqueries Lentas

### Processo de Debugging

1. **Identifique a Subquery Lenta**
   ```sql
   -- Use EXPLAIN QUERY PLAN para ver o plano de execução
   EXPLAIN QUERY PLAN
   SELECT titulo
   FROM livros l
   WHERE EXISTS (
       SELECT 1
       FROM emprestimos e
       WHERE e.livro_id = l.id
   );
   ```

2. **Teste a Subquery Separadamente**
   ```sql
   -- Teste a subquery sozinha
   SELECT 1
   FROM emprestimos e
   WHERE e.livro_id = 1;  -- Teste com um ID específico
   ```

3. **Verifique Índices**
   ```sql
   -- Verifique se há índices nas colunas usadas
   SELECT * FROM sqlite_master 
   WHERE type='index' AND tbl_name='emprestimos';
   ```

4. **Considere Alternativas**
   - Reescreva como JOIN
   - Use tabela derivada
   - Quebre em múltiplas queries

### Exemplo de Troubleshooting

**Problema**: Query muito lenta
```sql
SELECT titulo
FROM livros l
WHERE (
    SELECT COUNT(*)
    FROM emprestimos e
    WHERE e.livro_id = l.id
) > 5;
```

**Passo 1: Verificar Plano de Execução**
```sql
EXPLAIN QUERY PLAN
SELECT titulo
FROM livros l
WHERE (
    SELECT COUNT(*)
    FROM emprestimos e
    WHERE e.livro_id = l.id
) > 5;
-- Resultado: "SCAN TABLE emprestimos" (sem índice!)
```

**Passo 2: Criar Índice**
```sql
CREATE INDEX idx_emprestimos_livro ON emprestimos(livro_id);
```

**Passo 3: Reescrever como JOIN (ainda melhor)**
```sql
SELECT l.titulo
FROM livros l
JOIN (
    SELECT livro_id, COUNT(*) AS total
    FROM emprestimos
    GROUP BY livro_id
    HAVING COUNT(*) > 5
) AS stats ON l.id = stats.livro_id;
```

---

## 10. Checklist de Otimização

Antes de considerar uma query com subquery otimizada, verifique:

- [ ] **Índices**: Há índices nas colunas usadas nas subqueries?
- [ ] **Correlated Subqueries**: Posso reescrever como JOIN?
- [ ] **Aninhamento**: Posso reduzir o nível de aninhamento?
- [ ] **EXISTS vs IN**: Estou usando o operador mais eficiente?
- [ ] **NULLs**: Estou lidando corretamente com NULLs (NOT EXISTS vs NOT IN)?
- [ ] **Testabilidade**: A subquery pode ser testada separadamente?
- [ ] **Legibilidade**: A query é fácil de entender e manter?
- [ ] **Performance**: Testei a performance e está aceitável?
- [ ] **Alternativas**: Considerei JOINs, tabelas derivadas ou CTEs?

---

## 11. Casos Especiais e Armadilhas

### Armadilha 1: Subquery Retorna Múltiplas Linhas

**Problema**: Subquery escalar retorna múltiplas linhas
```sql
-- ❌ ERRO: "Subquery returns more than one row"
SELECT titulo
FROM livros
WHERE quantidade_disponivel = (
    SELECT quantidade_disponivel FROM livros  -- Retorna várias linhas!
);
```

**Solução**: Use operadores adequados ou adicione LIMIT
```sql
-- ✅ CORRETO: Use IN ou adicione LIMIT
SELECT titulo
FROM livros
WHERE quantidade_disponivel IN (
    SELECT quantidade_disponivel FROM livros WHERE categoria_id = 1
);
```

### Armadilha 2: NULL em NOT IN

**Problema**: NOT IN com NULLs pode retornar resultados vazios
```sql
-- ⚠️ PROBLEMA: Se subquery retornar NULL, resultado pode ser vazio
SELECT titulo
FROM livros
WHERE autor_id NOT IN (
    SELECT id FROM autores WHERE nacionalidade = 'Americano'
    -- Se algum id for NULL, resultado vazio!
);
```

**Solução**: Use NOT EXISTS ou filtre NULLs
```sql
-- ✅ CORRETO: Use NOT EXISTS
SELECT l.titulo
FROM livros l
WHERE NOT EXISTS (
    SELECT 1 FROM autores a 
    WHERE a.id = l.autor_id 
    AND a.nacionalidade = 'Americano'
);
```

### Armadilha 3: Performance de Correlated Subqueries

**Problema**: Correlated subquery executada muitas vezes
```sql
-- ❌ Pode ser muito lento
SELECT titulo
FROM livros l1
WHERE (
    SELECT COUNT(*)
    FROM emprestimos e
    WHERE e.livro_id = l1.id
) > 5;
```

**Solução**: Reescreva como JOIN
```sql
-- ✅ Geralmente mais rápido
SELECT l.titulo
FROM livros l
JOIN (
    SELECT livro_id, COUNT(*) AS total
    FROM emprestimos
    GROUP BY livro_id
    HAVING COUNT(*) > 5
) AS stats ON l.id = stats.livro_id;
```

---

## Conclusão

Subqueries são poderosas, mas requerem cuidado com performance:

- ✅ **Use índices** nas colunas de subqueries
- ✅ **Evite correlated subqueries** quando possível
- ✅ **Prefira JOINs** para combinar dados
- ✅ **Teste performance** de queries complexas
- ✅ **Documente** subqueries complexas
- ✅ **Considere alternativas** (JOINs, CTEs, tabelas derivadas)

**Lembre-se**: Performance não é apenas sobre velocidade - é sobre usar a ferramenta certa para cada trabalho. Subqueries são ótimas para alguns problemas, mas JOINs são melhores para outros. Escolha sabiamente!

**Próximos Passos**:
1. Revise suas queries com subqueries
2. Identifique oportunidades de otimização
3. Teste diferentes abordagens
4. Aplique as boas práticas aprendidas

**Bons estudos! 🚀**



