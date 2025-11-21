# Aula 8 - Performance, Boas Práticas e Otimização: Advanced SQL Functions

## Introdução

Funções SQL são poderosas, mas podem ter impacto significativo na performance se não forem usadas adequadamente. Nesta seção, você aprenderá como usar funções de forma eficiente, quando cada função é apropriada, e como balancear funcionalidade com performance.

---

## 1. Impacto de Funções na Performance

### Por que Funções Afetam Performance?

Funções SQL são executadas para cada linha processada, o que significa:

1. **Processamento por Linha**: Cada função é executada uma vez por linha
2. **Impedem Uso de Índices**: Funções em WHERE podem impedir uso de índices
3. **Cálculos Repetidos**: Funções complexas são recalculadas a cada execução
4. **Overhead de Transformação**: Transformar dados tem custo computacional

### Impacto por Tipo de Função

#### String Functions

**Impacto**: **Variável** (depende da função e tamanho da string)

```sql
-- Função simples (rápida)
SELECT UPPER(titulo) FROM livros;  -- Rápido

-- Função complexa (mais lenta)
SELECT REPLACE(REPLACE(REPLACE(titulo, ' ', '_'), '-', ''), '.', '') FROM livros;  -- Mais lento
```

**Fatores que Afetam Performance**:
- Tamanho da string (strings maiores = mais lento)
- Número de operações (múltiplas funções = mais lento)
- Complexidade da função (REPLACE múltiplo = mais lento)

#### Date/Time Functions

**Impacto**: **Moderado** (cálculos de data podem ser custosos)

```sql
-- Função simples (rápida)
SELECT strftime('%Y', data_emprestimo) FROM emprestimos;  -- Rápido

-- Cálculo complexo (mais lento)
SELECT julianday(data_devolucao) - julianday(data_emprestimo) FROM emprestimos;  -- Mais lento
```

**Fatores que Afetam Performance**:
- Conversões de formato (strftime pode ser lento)
- Cálculos de intervalo (julianday é mais custoso)
- Agrupamentos por data (pode ser lento em grandes volumes)

#### Numeric Functions

**Impacto**: **Baixo a Moderado** (cálculos matemáticos são geralmente rápidos)

```sql
-- Função simples (muito rápida)
SELECT ROUND(quantidade_disponivel, 2) FROM livros;  -- Muito rápido

-- Função com agregação (pode ser mais lento)
SELECT ROUND(AVG(quantidade_disponivel), 2) FROM livros GROUP BY categoria_id;  -- Mais lento
```

**Fatores que Afetam Performance**:
- Complexidade do cálculo (FLOOR/CEIL são rápidos, cálculos complexos são mais lentos)
- Volume de dados (muitas linhas = mais tempo)

#### Conditional Functions

**Impacto**: **Variável** (depende da complexidade da lógica)

```sql
-- CASE simples (rápido)
SELECT CASE WHEN quantidade > 0 THEN 'Disponível' ELSE 'Esgotado' END FROM livros;  -- Rápido

-- CASE complexo (mais lento)
SELECT CASE 
    WHEN quantidade = 0 THEN 'Esgotado'
    WHEN quantidade < 5 THEN 'Baixo'
    WHEN quantidade < 10 THEN 'Médio'
    WHEN quantidade < 20 THEN 'Alto'
    ELSE 'Muito Alto'
END FROM livros;  -- Mais lento (mas ainda rápido)
```

**Fatores que Afetam Performance**:
- Número de condições (mais condições = mais lento)
- Complexidade das condições (subqueries em condições = muito lento)

---

## 2. Funções em WHERE: O Problema dos Índices

### Por que Funções em WHERE São Problemáticas?

Funções em WHERE impedem o uso de índices porque o banco precisa transformar cada valor antes de comparar.

**Exemplo Problemático**:
```sql
-- ❌ LENTO: função impede uso de índice
SELECT * FROM livros 
WHERE UPPER(titulo) = 'FUNDAÇÃO';
-- Banco precisa: 1) Ler todas as linhas, 2) Aplicar UPPER, 3) Comparar
-- Não pode usar índice em 'titulo'
```

**Solução 1: Usar Valor Direto**
```sql
-- ✅ RÁPIDO: índice pode ser usado
SELECT * FROM livros 
WHERE titulo = 'Fundação';
-- Banco pode usar índice diretamente
```

**Solução 2: Indexar Coluna Transformada**
```sql
-- Criar índice na coluna transformada
CREATE INDEX idx_livros_titulo_upper ON livros(UPPER(titulo));

-- Agora a query é rápida
SELECT * FROM livros 
WHERE UPPER(titulo) = 'FUNDAÇÃO';
-- Banco pode usar índice em UPPER(titulo)
```

**Solução 3: Normalizar Dados na Inserção**
```sql
-- Adicionar coluna normalizada
ALTER TABLE livros ADD COLUMN titulo_normalizado TEXT;

-- Atualizar dados existentes
UPDATE livros SET titulo_normalizado = UPPER(titulo);

-- Criar índice na coluna normalizada
CREATE INDEX idx_livros_titulo_normalizado ON livros(titulo_normalizado);

-- Query rápida
SELECT * FROM livros 
WHERE titulo_normalizado = 'FUNDAÇÃO';
```

### Outros Exemplos Problemáticos

```sql
-- ❌ LENTO: função em WHERE
SELECT * FROM emprestimos 
WHERE strftime('%Y', data_emprestimo) = '2024';

-- ✅ MELHOR: usar range de datas
SELECT * FROM emprestimos 
WHERE data_emprestimo >= '2024-01-01' 
  AND data_emprestimo < '2025-01-01';

-- ❌ LENTO: função em WHERE
SELECT * FROM livros 
WHERE LENGTH(titulo) > 30;

-- ✅ MELHOR: criar coluna calculada (se necessário frequentemente)
-- Ou aceitar que precisa processar todas as linhas
```

---

## 3. Funções Aninhadas e Complexas

### Impacto de Funções Aninhadas

Funções aninhadas são executadas de dentro para fora, multiplicando o custo:

```sql
-- Função aninhada: cada nível adiciona overhead
SELECT REPLACE(UPPER(SUBSTR(titulo, 1, 10)), ' ', '_') FROM livros;

-- Execução:
-- 1. SUBSTR(titulo, 1, 10) → resultado1
-- 2. UPPER(resultado1) → resultado2
-- 3. REPLACE(resultado2, ' ', '_') → resultado final
-- Custo: 3 operações por linha
```

### Quando Simplificar

**Antes (Complexo)**:
```sql
SELECT 
    REPLACE(
        UPPER(
            SUBSTR(
                REPLACE(titulo, ' ', '_'),
                1,
                10
            )
        ),
        '_',
        '-'
    ) AS codigo
FROM livros;
```

**Depois (Simplificado)**:
```sql
-- Se possível, fazer em etapas ou criar coluna calculada
ALTER TABLE livros ADD COLUMN codigo TEXT;

UPDATE livros SET codigo = UPPER(SUBSTR(REPLACE(titulo, ' ', '_'), 1, 10));

SELECT codigo FROM livros;
```

### Boa Prática: Cache de Resultados

Para cálculos complexos que são usados frequentemente:

```sql
-- Criar coluna calculada
ALTER TABLE livros ADD COLUMN titulo_normalizado TEXT;
ALTER TABLE livros ADD COLUMN tamanho_titulo INTEGER;

-- Atualizar uma vez
UPDATE livros SET 
    titulo_normalizado = UPPER(titulo),
    tamanho_titulo = LENGTH(titulo);

-- Criar índices
CREATE INDEX idx_livros_titulo_norm ON livros(titulo_normalizado);
CREATE INDEX idx_livros_tamanho ON livros(tamanho_titulo);

-- Queries rápidas
SELECT * FROM livros WHERE titulo_normalizado = 'FUNDAÇÃO';
SELECT * FROM livros WHERE tamanho_titulo > 30;
```

---

## 4. Funções em SELECT vs WHERE vs GROUP BY

### Funções em SELECT

**Impacto**: **Baixo** (apenas transforma resultado, não afeta filtros)

```sql
-- Funções em SELECT são geralmente OK
SELECT 
    UPPER(titulo) AS titulo_maiusculo,
    LENGTH(titulo) AS tamanho
FROM livros;
-- Performance: Boa (apenas transforma resultados)
```

**Quando é Problemático**:
- Muitas funções aninhadas
- Funções muito complexas em muitas linhas
- Funções que fazem subqueries

### Funções em WHERE

**Impacto**: **Alto** (impede índices, processa todas as linhas)

```sql
-- ❌ PROBLEMÁTICO
SELECT * FROM livros 
WHERE UPPER(titulo) LIKE '%FUNDAÇÃO%';
-- Processa TODAS as linhas antes de filtrar
```

**Soluções**:
- Evitar funções em WHERE quando possível
- Usar valores diretos
- Criar índices em colunas transformadas
- Normalizar dados na inserção

### Funções em GROUP BY

**Impacto**: **Moderado a Alto** (agrupa após transformar)

```sql
-- Agrupar por ano (função em GROUP BY)
SELECT 
    strftime('%Y', data_emprestimo) AS ano,
    COUNT(*) AS total
FROM emprestimos
GROUP BY strftime('%Y', data_emprestimo);
-- Performance: Moderada (precisa transformar antes de agrupar)
```

**Otimização**:
```sql
-- Criar coluna calculada
ALTER TABLE emprestimos ADD COLUMN ano_emprestimo INTEGER;

UPDATE emprestimos SET ano_emprestimo = CAST(strftime('%Y', data_emprestimo) AS INTEGER);

CREATE INDEX idx_emprestimos_ano ON emprestimos(ano_emprestimo);

-- Query otimizada
SELECT ano_emprestimo, COUNT(*) AS total
FROM emprestimos
GROUP BY ano_emprestimo;
```

---

## 5. Estratégias de Otimização

### Estratégia 1: Evitar Funções Quando Possível

```sql
-- ❌ EVITE
SELECT * FROM livros WHERE UPPER(titulo) = 'FUNDAÇÃO';

-- ✅ PREFIRA
SELECT * FROM livros WHERE titulo = 'Fundação';
```

### Estratégia 2: Normalizar Dados na Inserção

```sql
-- Criar colunas normalizadas
CREATE TABLE livros (
    id INTEGER PRIMARY KEY,
    titulo TEXT,
    titulo_normalizado TEXT,  -- Já em maiúsculas
    tamanho_titulo INTEGER     -- Já calculado
);

-- Inserir com valores normalizados
INSERT INTO livros (titulo, titulo_normalizado, tamanho_titulo)
VALUES ('Fundação', 'FUNDAÇÃO', 9);

-- Queries rápidas
SELECT * FROM livros WHERE titulo_normalizado = 'FUNDAÇÃO';
SELECT * FROM livros WHERE tamanho_titulo > 30;
```

### Estratégia 3: Usar Índices em Colunas Transformadas

```sql
-- Criar índice em coluna transformada
CREATE INDEX idx_livros_titulo_upper ON livros(UPPER(titulo));

-- Query pode usar índice
SELECT * FROM livros WHERE UPPER(titulo) = 'FUNDAÇÃO';
```

**Limitação**: SQLite não suporta índices em expressões diretamente. Use colunas calculadas.

### Estratégia 4: Cache de Resultados Complexos

```sql
-- Para cálculos muito complexos, criar coluna calculada
ALTER TABLE emprestimos ADD COLUMN dias_emprestimo INTEGER;

UPDATE emprestimos SET 
    dias_emprestimo = julianday(COALESCE(data_devolucao_real, 'now')) - julianday(data_emprestimo);

-- Query rápida
SELECT * FROM emprestimos WHERE dias_emprestimo > 30;
```

### Estratégia 5: Usar Views para Queries Complexas

```sql
-- Criar view com transformações
CREATE VIEW livros_formatados AS
SELECT 
    id,
    UPPER(titulo) AS titulo_maiusculo,
    LENGTH(titulo) AS tamanho_titulo,
    CASE 
        WHEN quantidade_disponivel = 0 THEN 'Esgotado'
        ELSE 'Disponível'
    END AS status
FROM livros;

-- Usar view (transformações são calculadas, mas query é mais limpa)
SELECT * FROM livros_formatados WHERE status = 'Esgotado';
```

---

## 6. Medindo Performance

### Como Medir Impacto de Funções

```sql
-- Habilitar timer no SQLite
.timer ON

-- Teste 1: Sem função
SELECT * FROM livros WHERE titulo = 'Fundação';

-- Teste 2: Com função
SELECT * FROM livros WHERE UPPER(titulo) = 'FUNDAÇÃO';

-- Comparar tempos
```

### EXPLAIN QUERY PLAN

```sql
-- Ver plano de execução
EXPLAIN QUERY PLAN
SELECT * FROM livros WHERE UPPER(titulo) = 'FUNDAÇÃO';

-- Verificar se usa índice ou faz scan completo
```

**Interpretação**:
- `SEARCH` ou `SCAN TABLE USING INDEX`: Bom (usa índice)
- `SCAN TABLE`: Ruim (scan completo, lento)

---

## 7. Boas Práticas por Tipo de Função

### String Functions

**Boas Práticas**:
1. **Evite funções em WHERE**: Use valores diretos quando possível
2. **Normalize na inserção**: Armazene dados já normalizados
3. **Limite aninhamento**: Evite muitas funções aninhadas
4. **Use índices**: Crie índices em colunas normalizadas

**Exemplo**:
```sql
-- ✅ BOM: Normalizar na inserção
INSERT INTO livros (titulo, titulo_normalizado) 
VALUES ('Fundação', UPPER('Fundação'));

-- ✅ BOM: Query rápida
SELECT * FROM livros WHERE titulo_normalizado = 'FUNDAÇÃO';
```

### Date/Time Functions

**Boas Práticas**:
1. **Use ranges ao invés de funções**: `data >= '2024-01-01' AND data < '2025-01-01'`
2. **Extraia partes na inserção**: Armazene ano, mês separadamente se necessário
3. **Evite cálculos complexos em WHERE**: Pre-calcule intervalos

**Exemplo**:
```sql
-- ❌ EVITE
SELECT * FROM emprestimos WHERE strftime('%Y', data_emprestimo) = '2024';

-- ✅ PREFIRA
SELECT * FROM emprestimos 
WHERE data_emprestimo >= '2024-01-01' 
  AND data_emprestimo < '2025-01-01';
```

### Numeric Functions

**Boas Práticas**:
1. **Funções simples são rápidas**: ROUND, FLOOR, CEIL são geralmente OK
2. **Evite em WHERE se possível**: Use valores diretos
3. **Pre-calcule se usado frequentemente**: Armazene valores arredondados

**Exemplo**:
```sql
-- ✅ BOM: ROUND em SELECT é geralmente OK
SELECT ROUND(AVG(quantidade_disponivel), 2) FROM livros;

-- ❌ EVITE: Função em WHERE
SELECT * FROM livros WHERE ROUND(quantidade_disponivel / 2.0) = 5;

-- ✅ PREFIRA: Calcular antes
SELECT * FROM livros WHERE quantidade_disponivel BETWEEN 9 AND 11;
```

### Conditional Functions

**Boas Práticas**:
1. **CASE simples é rápido**: Use sem preocupação em SELECT
2. **Evite subqueries em CASE**: Pode ser muito lento
3. **Considere coluna calculada**: Para classificações complexas usadas frequentemente

**Exemplo**:
```sql
-- ✅ BOM: CASE simples
SELECT 
    CASE 
        WHEN quantidade > 0 THEN 'Disponível'
        ELSE 'Esgotado'
    END AS status
FROM livros;

-- ❌ EVITE: Subquery em CASE (muito lento)
SELECT 
    CASE 
        WHEN quantidade > (SELECT AVG(quantidade) FROM livros) THEN 'Acima da média'
        ELSE 'Abaixo da média'
    END AS status
FROM livros;
```

---

## 8. Quando Processar na Aplicação vs no Banco

### Processar no Banco Quando:

1. **Filtros e Ordenações**: Precisa filtrar/ordenar por valor transformado
2. **Agregações**: Precisa agrupar por valor transformado
3. **Consistência**: Quer garantir mesma transformação para todos
4. **Performance**: Processamento no banco é mais eficiente
5. **Índices**: Pode criar índices em colunas transformadas

### Processar na Aplicação Quando:

1. **Lógica Complexa**: Transformação muito complexa ou específica da aplicação
2. **Flexibilidade**: Precisa mudar transformação frequentemente
3. **Formatação de Apresentação**: Apenas para exibição (não para filtros)
4. **Dados Pequenos**: Volume pequeno, processamento na aplicação é aceitável
5. **Lógica de Negócio**: Regras específicas da aplicação que mudam frequentemente

### Exemplo de Decisão

```sql
-- Processar no banco (para filtro)
SELECT * FROM livros 
WHERE titulo_normalizado LIKE '%FUNDAÇÃO%';
-- Precisa estar no banco para usar índice

-- Processar na aplicação (apenas formatação)
SELECT titulo FROM livros;
-- Na aplicação: display_titulo = titulo.toUpperCase()
-- Não precisa estar no banco se não filtra/ordena
```

---

## 9. Troubleshooting de Queries Lentas

### Identificar Problemas

1. **Usar EXPLAIN QUERY PLAN**:
```sql
EXPLAIN QUERY PLAN
SELECT * FROM livros WHERE UPPER(titulo) = 'FUNDAÇÃO';
```

2. **Medir Tempo**:
```sql
.timer ON
SELECT * FROM livros WHERE UPPER(titulo) = 'FUNDAÇÃO';
```

3. **Verificar Índices**:
```sql
SELECT * FROM sqlite_master WHERE type='index' AND tbl_name='livros';
```

### Soluções Comuns

**Problema**: Query lenta com função em WHERE
**Solução**: Remover função ou criar coluna normalizada

**Problema**: Muitas funções aninhadas
**Solução**: Simplificar ou criar coluna calculada

**Problema**: Função em GROUP BY lenta
**Solução**: Criar coluna calculada e indexar

**Problema**: Cálculos complexos repetidos
**Solução**: Cache resultado em coluna calculada

---

## 10. Resumo de Boas Práticas

### Regras Gerais

1. ✅ **Evite funções em WHERE** quando possível
2. ✅ **Normalize dados na inserção** para queries frequentes
3. ✅ **Use índices** em colunas transformadas
4. ✅ **Pre-calcule valores** usados frequentemente
5. ✅ **Simplifique funções aninhadas** quando possível
6. ✅ **Meça performance** antes e depois de otimizações
7. ✅ **Use EXPLAIN QUERY PLAN** para entender execução
8. ✅ **Considere views** para queries complexas frequentes

### Checklist de Performance

Antes de usar uma função, pergunte:

- [ ] Esta função está em WHERE? (Se sim, pode ser lenta)
- [ ] Esta função é usada frequentemente? (Se sim, considere normalizar)
- [ ] Esta função está aninhada muitas vezes? (Se sim, simplifique)
- [ ] Posso usar valor direto ao invés de função? (Se sim, prefira)
- [ ] Posso criar coluna calculada? (Se sim, considere)
- [ ] Esta função impede uso de índice? (Se sim, otimize)

---

## 11. Conclusão

Funções SQL são poderosas, mas precisam ser usadas com sabedoria:

- **Use funções** para transformações necessárias
- **Evite funções em WHERE** quando possível
- **Normalize dados** para queries frequentes
- **Meça performance** antes e depois
- **Otimize** quando necessário

**Lembre-se**: Performance é um equilíbrio entre funcionalidade e velocidade. Nem sempre a solução mais rápida é a melhor - considere também legibilidade e manutenibilidade.

---

**Bons estudos! 🚀**

