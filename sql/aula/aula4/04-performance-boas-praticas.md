# Aula 4 - Performance, Boas Práticas e Otimização

## Introdução

Aggregate queries são poderosas, mas podem ser lentas se não forem otimizadas corretamente. Nesta seção, você aprenderá como escrever aggregate queries eficientes, quando usar índices, e como pensar sobre performance desde o início.

---

## 1. Impacto de Aggregate Queries na Performance

### Por que Aggregate Queries Podem Ser Lentas?

Aggregate queries processam **múltiplas linhas** para produzir resultados resumidos. Isso significa:

1. **Leitura de Muitos Dados**: Precisam ler muitas linhas da tabela
2. **Processamento Intensivo**: Aplicam funções de agregação em grandes volumes
3. **Agrupamento Custo**: GROUP BY precisa ordenar e agrupar dados
4. **JOINs Adicionais**: Frequentemente combinam múltiplas tabelas

### Exemplo de Impacto

```sql
-- Query simples (rápida)
SELECT titulo FROM livros WHERE id = 1;
-- Lê 1 linha, retorna imediatamente

-- Aggregate query (pode ser lenta)
SELECT categoria_id, COUNT(*), SUM(quantidade_disponivel)
FROM livros
GROUP BY categoria_id;
-- Lê TODAS as linhas, agrupa, calcula agregações
```

**Com 1.000 livros**: Rápido (milissegundos)
**Com 1.000.000 livros**: Pode levar segundos ou minutos sem otimização

---

## 2. Índices e Aggregate Queries

### Quando Índices Ajudam em Aggregate Queries

Índices são cruciais para otimizar aggregate queries, especialmente em:

#### 2.1 Índices para GROUP BY

Se você agrupa por uma coluna, um índice nessa coluna pode acelerar significativamente:

```sql
-- Query que se beneficia de índice
SELECT categoria_id, COUNT(*)
FROM livros
GROUP BY categoria_id;

-- Índice recomendado:
CREATE INDEX idx_livros_categoria ON livros(categoria_id);
```

**Por quê?** O índice já organiza os dados por `categoria_id`, facilitando o agrupamento.

#### 2.2 Índices para WHERE + GROUP BY

Quando você combina WHERE com GROUP BY, índices compostos podem ajudar:

```sql
-- Query com WHERE e GROUP BY
SELECT categoria_id, COUNT(*)
FROM livros
WHERE quantidade_disponivel > 0
GROUP BY categoria_id;

-- Índice composto recomendado:
CREATE INDEX idx_livros_categoria_estoque 
ON livros(categoria_id, quantidade_disponivel);
```

**Ordem do índice**: Coloque a coluna do GROUP BY primeiro, depois a do WHERE.

#### 2.3 Índices para JOINs em Aggregate Queries

JOINs em aggregate queries se beneficiam de índices nas chaves estrangeiras:

```sql
-- Query com JOIN
SELECT c.nome, COUNT(*)
FROM livros l
JOIN categorias c ON l.categoria_id = c.id
GROUP BY c.id, c.nome;

-- Índices recomendados:
CREATE INDEX idx_livros_categoria ON livros(categoria_id);
-- (c.id já é PRIMARY KEY, então já tem índice)
```

### Quando Índices NÃO Ajudam

Índices **não ajudam** (e podem até atrasar) quando:

1. **Agregação em Todas as Linhas**: Se você precisa processar todas as linhas de qualquer forma
2. **Funções de Agregação Complexas**: AVG, SUM precisam ler todos os valores
3. **Múltiplas Agregações**: Muitas funções diferentes na mesma query

```sql
-- Índice não ajuda muito aqui (precisa ler todas as linhas)
SELECT 
    SUM(quantidade_disponivel),
    AVG(quantidade_disponivel),
    COUNT(*)
FROM livros;
```

---

## 3. Otimização de GROUP BY

### 3.1 Filtrar Antes de Agrupar (WHERE vs HAVING)

**✅ SEMPRE use WHERE quando possível** - filtra linhas antes do agrupamento:

```sql
-- ✅ EFICIENTE: Filtra antes de agrupar
SELECT categoria_id, COUNT(*)
FROM livros
WHERE quantidade_disponivel > 0  -- Filtra primeiro
GROUP BY categoria_id;

-- ❌ MENOS EFICIENTE: Agrupa tudo, depois filtra grupos
SELECT categoria_id, COUNT(*)
FROM livros
GROUP BY categoria_id
HAVING COUNT(*) > 5;  -- Filtra depois (mas ainda precisa agrupar tudo)
```

**Regra de ouro**: Use WHERE para filtrar linhas, HAVING apenas para filtrar grupos baseado em agregações.

### 3.2 Limitar Colunas no GROUP BY

Agrupe apenas pelas colunas necessárias:

```sql
-- ✅ BOM: Agrupa apenas pelo necessário
SELECT categoria_id, COUNT(*)
FROM livros
GROUP BY categoria_id;

-- ❌ EVITE: Agrupa por colunas desnecessárias
SELECT categoria_id, COUNT(*)
FROM livros
GROUP BY categoria_id, autor_id, editora;  -- Desnecessário se não usar essas colunas
```

### 3.3 Evitar GROUP BY Desnecessário

Não use GROUP BY se você não precisa agrupar:

```sql
-- ❌ ERRADO: GROUP BY desnecessário
SELECT COUNT(*) FROM livros GROUP BY 1;

-- ✅ CORRETO: Sem GROUP BY
SELECT COUNT(*) FROM livros;
```

---

## 4. Otimização de Funções de Agregação

### 4.1 Evitar Múltiplas Passagens pelos Dados

Se possível, calcule todas as agregações em uma única query:

```sql
-- ✅ BOM: Uma única passagem
SELECT 
    COUNT(*) AS total,
    SUM(quantidade_disponivel) AS total_estoque,
    AVG(quantidade_disponivel) AS media_estoque
FROM livros;

-- ❌ EVITE: Múltiplas queries (múltiplas passagens)
SELECT COUNT(*) FROM livros;
SELECT SUM(quantidade_disponivel) FROM livros;
SELECT AVG(quantidade_disponivel) FROM livros;
```

### 4.2 Usar COUNT(*) vs COUNT(coluna) Apropriadamente

- **COUNT(*)**: Mais rápido, conta todas as linhas
- **COUNT(coluna)**: Mais lento, precisa verificar NULL

```sql
-- ✅ Use COUNT(*) se não precisa verificar NULL
SELECT COUNT(*) FROM livros;

-- ✅ Use COUNT(coluna) apenas se precisa ignorar NULL
SELECT COUNT(ano_publicacao) FROM livros;  -- Ignora livros sem ano
```

### 4.3 Evitar Agregações em Subqueries Desnecessárias

```sql
-- ❌ EVITE: Subquery desnecessária
SELECT 
    categoria_id,
    (SELECT COUNT(*) FROM livros l2 WHERE l2.categoria_id = l1.categoria_id) AS total
FROM livros l1
GROUP BY categoria_id;

-- ✅ MELHOR: Agregação direta
SELECT 
    categoria_id,
    COUNT(*) AS total
FROM livros
GROUP BY categoria_id;
```

---

## 5. Otimização de JOINs em Aggregate Queries

### 5.1 JOIN Apenas o Necessário

Faça JOIN apenas com tabelas que você realmente usa:

```sql
-- ❌ EVITE: JOIN desnecessário
SELECT 
    c.nome,
    COUNT(*)
FROM livros l
JOIN categorias c ON l.categoria_id = c.id
JOIN autores a ON l.autor_id = a.id  -- Não usado na query!
GROUP BY c.id, c.nome;

-- ✅ MELHOR: Apenas JOIN necessário
SELECT 
    c.nome,
    COUNT(*)
FROM livros l
JOIN categorias c ON l.categoria_id = c.id
GROUP BY c.id, c.nome;
```

### 5.2 Usar INNER JOIN vs LEFT JOIN Apropriadamente

- **INNER JOIN**: Mais rápido, retorna apenas correspondências
- **LEFT JOIN**: Mais lento, precisa processar linhas sem correspondência

```sql
-- ✅ Use INNER JOIN se não precisa de NULLs
SELECT c.nome, COUNT(*)
FROM livros l
INNER JOIN categorias c ON l.categoria_id = c.id
GROUP BY c.id, c.nome;

-- ✅ Use LEFT JOIN apenas se precisa incluir livros sem categoria
SELECT 
    COALESCE(c.nome, 'Sem categoria') AS categoria,
    COUNT(*)
FROM livros l
LEFT JOIN categorias c ON l.categoria_id = c.id
GROUP BY c.id, c.nome;
```

---

## 6. Análise de Performance com EXPLAIN

### 6.1 Usando EXPLAIN QUERY PLAN

SQLite oferece `EXPLAIN QUERY PLAN` para analisar como a query será executada:

```sql
EXPLAIN QUERY PLAN
SELECT categoria_id, COUNT(*)
FROM livros
GROUP BY categoria_id;
```

**O que procurar:**
- **SCAN TABLE**: Lê toda a tabela (pode ser lento)
- **SEARCH TABLE USING INDEX**: Usa índice (geralmente rápido)
- **USE TEMP B-TREE FOR GROUP BY**: Cria estrutura temporária para agrupar

### 6.2 Interpretando Resultados

```
QUERY PLAN
|--SCAN TABLE livros          ← Lê toda a tabela
`--USE TEMP B-TREE FOR GROUP BY  ← Cria estrutura temporária
```

**Com índice:**
```
QUERY PLAN
|--SEARCH TABLE livros USING INDEX idx_livros_categoria  ← Usa índice!
`--USE TEMP B-TREE FOR GROUP BY
```

### 6.3 Exemplo Prático

```sql
-- 1. Ver plano sem índice
EXPLAIN QUERY PLAN
SELECT categoria_id, COUNT(*)
FROM livros
GROUP BY categoria_id;

-- 2. Criar índice
CREATE INDEX idx_livros_categoria ON livros(categoria_id);

-- 3. Ver plano com índice
EXPLAIN QUERY PLAN
SELECT categoria_id, COUNT(*)
FROM livros
GROUP BY categoria_id;
```

---

## 7. Boas Práticas de Nomenclatura

### 7.1 Aliases Descritivos

Use aliases claros e descritivos:

```sql
-- ✅ BOM: Aliases descritivos
SELECT 
    c.nome AS categoria,
    COUNT(*) AS total_livros,
    SUM(l.quantidade_disponivel) AS total_estoque,
    AVG(l.quantidade_disponivel) AS media_estoque
FROM livros l
JOIN categorias c ON l.categoria_id = c.id
GROUP BY c.id, c.nome;

-- ❌ EVITE: Aliases confusos
SELECT 
    c.nome AS c1,
    COUNT(*) AS c2,
    SUM(l.quantidade_disponivel) AS s1
FROM livros l
JOIN categorias c ON l.categoria_id = c.id
GROUP BY c.id, c.nome;
```

### 7.2 Consistência em Nomes

Mantenha padrão consistente:
- `total_*` para contagens e somas
- `media_*` ou `avg_*` para médias
- `min_*` e `max_*` para extremos

---

## 8. Tratamento de NULL em Produção

### 8.1 Sempre Considere NULL

Em produção, sempre considere como NULL será tratado:

```sql
-- ✅ BOM: Trata NULL explicitamente
SELECT 
    categoria_id,
    COUNT(*) AS total_livros,
    COUNT(ano_publicacao) AS livros_com_ano,
    AVG(CASE WHEN ano_publicacao IS NOT NULL THEN ano_publicacao END) AS media_ano
FROM livros
GROUP BY categoria_id;

-- ❌ EVITE: Ignorar NULL pode causar resultados inesperados
SELECT 
    categoria_id,
    AVG(ano_publicacao) AS media_ano  -- Pode retornar NULL se todos forem NULL
FROM livros
GROUP BY categoria_id;
```

### 8.2 Usar COALESCE para Valores Padrão

```sql
-- Fornece valor padrão quando agregação retorna NULL
SELECT 
    categoria_id,
    COALESCE(AVG(quantidade_disponivel), 0) AS media_estoque
FROM livros
GROUP BY categoria_id;
```

---

## 9. Escalabilidade e Cache

### 9.1 Quando Cachear Resultados

Considere cachear resultados de aggregate queries quando:

1. **Dados mudam pouco**: Estatísticas que não precisam ser em tempo real
2. **Query é executada frequentemente**: Centenas ou milhares de vezes por dia
3. **Query é custosa**: Leva vários segundos para executar

### 9.2 Estratégias de Cache

#### Opção 1: Tabela de Estatísticas

```sql
-- Criar tabela de cache
CREATE TABLE estatisticas_categorias (
    categoria_id INTEGER PRIMARY KEY,
    total_livros INTEGER,
    total_estoque INTEGER,
    media_estoque REAL,
    ultima_atualizacao TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Atualizar periodicamente (ex: a cada hora)
INSERT OR REPLACE INTO estatisticas_categorias
SELECT 
    categoria_id,
    COUNT(*) AS total_livros,
    SUM(quantidade_disponivel) AS total_estoque,
    AVG(quantidade_disponivel) AS media_estoque,
    CURRENT_TIMESTAMP
FROM livros
GROUP BY categoria_id;
```

#### Opção 2: Views Materializadas (PostgreSQL, outros SGBDs)

Alguns SGBDs suportam views materializadas que são atualizadas automaticamente.

### 9.3 Quando NÃO Cachear

Não cacheie se:
- Dados mudam frequentemente
- Precisa de resultados em tempo real
- Query é rápida o suficiente (< 100ms)

---

## 10. Segurança em Aggregate Queries

### 10.1 Proteção contra SQL Injection

Mesmo em aggregate queries, sempre use parâmetros:

```sql
-- ✅ SEGURO: Usa parâmetro
SELECT categoria_id, COUNT(*)
FROM livros
WHERE categoria_id = ?  -- Parâmetro
GROUP BY categoria_id;

-- ❌ INSEGURO: Concatenação de strings
SELECT categoria_id, COUNT(*)
FROM livros
WHERE categoria_id = $variavel  -- Perigoso!
GROUP BY categoria_id;
```

### 10.2 Limitar Resultados

Use LIMIT para evitar retornar milhões de linhas:

```sql
-- ✅ BOM: Limita resultados
SELECT categoria_id, COUNT(*)
FROM livros
GROUP BY categoria_id
ORDER BY COUNT(*) DESC
LIMIT 10;

-- ❌ EVITE: Pode retornar muitos grupos
SELECT categoria_id, COUNT(*)
FROM livros
GROUP BY categoria_id;  -- E se houver 1000 categorias?
```

---

## 11. Monitoramento e Troubleshooting

### 11.1 Identificar Queries Lentas

Monitore queries que:
- Levam mais de 1 segundo
- Processam mais de 100.000 linhas
- Fazem múltiplos JOINs com agregações

### 11.2 Métricas Importantes

Acompanhe:
- **Tempo de execução**: Quanto tempo a query leva
- **Linhas processadas**: Quantas linhas foram lidas
- **Uso de índices**: Se índices estão sendo usados
- **Uso de memória**: Se cria estruturas temporárias grandes

### 11.3 Quando Revisar Queries

Revise e otimize quando:
- Query leva mais de 1 segundo regularmente
- Usuários reclamam de lentidão
- Sistema fica lento em horários de pico
- Dados cresceram significativamente

---

## 12. Checklist de Otimização

Ao escrever aggregate queries, verifique:

- [ ] **Índices apropriados** criados nas colunas de GROUP BY e JOIN
- [ ] **WHERE usado** para filtrar antes de agrupar (quando possível)
- [ ] **HAVING usado** apenas para filtrar grupos baseado em agregações
- [ ] **JOINs mínimos** - apenas tabelas necessárias
- [ ] **Múltiplas agregações** calculadas em uma única query
- [ ] **NULL tratado** apropriadamente
- [ ] **LIMIT usado** quando apropriado
- [ ] **EXPLAIN QUERY PLAN** analisado para entender performance
- [ ] **Aliases descritivos** usados
- [ ] **Parâmetros** usados (não concatenação de strings)

---

## 13. Exemplos de Otimização

### Exemplo 1: Query Lenta → Otimizada

**❌ Query Lenta:**
```sql
SELECT 
    c.nome,
    COUNT(*) AS total
FROM livros l
JOIN categorias c ON l.categoria_id = c.id
GROUP BY c.id, c.nome
HAVING COUNT(*) > 2;
```

**✅ Otimizações:**
```sql
-- 1. Criar índice
CREATE INDEX idx_livros_categoria ON livros(categoria_id);

-- 2. Query otimizada (se possível, adicionar WHERE)
SELECT 
    c.nome,
    COUNT(*) AS total
FROM livros l
JOIN categorias c ON l.categoria_id = c.id
WHERE l.quantidade_disponivel > 0  -- Filtra antes se possível
GROUP BY c.id, c.nome
HAVING COUNT(*) > 2;
```

### Exemplo 2: Múltiplas Queries → Uma Query

**❌ Múltiplas queries:**
```sql
SELECT COUNT(*) FROM livros;
SELECT SUM(quantidade_disponivel) FROM livros;
SELECT AVG(quantidade_disponivel) FROM livros;
```

**✅ Uma query:**
```sql
SELECT 
    COUNT(*) AS total,
    SUM(quantidade_disponivel) AS total_estoque,
    AVG(quantidade_disponivel) AS media_estoque
FROM livros;
```

---

## 14. Normalização vs Desnormalização para Agregações

### Quando Normalização é Boa

Normalização (dados separados em tabelas) é boa quando:
- Dados mudam frequentemente
- Precisamos de integridade referencial
- Storage é limitado

### Quando Desnormalização Pode Ajudar

Para aggregate queries frequentes, às vezes vale desnormalizar:

```sql
-- Tabela desnormalizada para estatísticas rápidas
CREATE TABLE livros_com_estatisticas (
    id INTEGER PRIMARY KEY,
    titulo TEXT,
    categoria_id INTEGER,
    categoria_nome TEXT,  -- Desnormalizado!
    total_estoque INTEGER,  -- Desnormalizado!
    ...
);
```

**Trade-off**: Mais rápido para ler, mais lento para atualizar.

---

## 15. Conclusão

Aggregate queries são poderosas, mas requerem atenção à performance:

1. **Use índices** nas colunas de GROUP BY e JOIN
2. **Filtre com WHERE** antes de agrupar quando possível
3. **Use HAVING** apenas para filtrar grupos
4. **Monitore performance** com EXPLAIN QUERY PLAN
5. **Considere cache** para queries frequentes e custosas
6. **Trate NULL** apropriadamente
7. **Use LIMIT** para evitar resultados enormes

Lembre-se: **Otimização prematura é ruim, mas ignorar performance também é ruim**. Encontre o equilíbrio baseado nas necessidades reais do seu sistema.

---

**Próximo Passo**: Após completar os exercícios, envie suas respostas para análise e feedback!

---

**💡 Dica Final**: Sempre teste queries em dados de tamanho similar à produção. Uma query rápida com 100 registros pode ser lenta com 1 milhão!
