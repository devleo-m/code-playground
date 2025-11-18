# Aula 4: Aggregate Queries (Consultas de Agregação)

## Introdução

Nesta aula, você aprenderá sobre **Aggregate Queries** (Consultas de Agregação), uma das funcionalidades mais poderosas do SQL. Enquanto as queries básicas que aprendemos até agora trabalham com registros individuais, as queries de agregação permitem realizar cálculos e análises sobre **conjuntos de dados**, retornando valores resumidos e estatísticas.

Aggregate queries são essenciais para:
- Gerar relatórios e estatísticas
- Analisar tendências e padrões nos dados
- Calcular totais, médias, contagens e extremos
- Agrupar dados por categorias
- Filtrar grupos baseado em condições agregadas

Dominar aggregate queries é fundamental para qualquer desenvolvedor ou analista de dados, pois permite transformar dados brutos em informações úteis e acionáveis.

---

## 1. O que são Aggregate Queries?

**Aggregate Queries** (Consultas de Agregação) são queries SQL que realizam cálculos sobre múltiplas linhas de dados, retornando um único valor resumido ou resultados agrupados. Elas utilizam **funções de agregação** para processar conjuntos de dados e produzir informações consolidadas.

### Características das Aggregate Queries

1. **Processam Múltiplas Linhas**: Ao invés de retornar cada registro individual, agregam valores de várias linhas
2. **Retornam Valores Resumidos**: Produzem totais, médias, contagens, máximos, mínimos
3. **Usam Funções de Agregação**: COUNT, SUM, AVG, MIN, MAX são as principais
4. **Podem Agrupar Dados**: GROUP BY permite agrupar resultados por categorias
5. **Filtram Grupos**: HAVING permite filtrar grupos baseado em condições agregadas

### Diferença entre Queries Normais e Aggregate Queries

**Query Normal (retorna múltiplas linhas):**
```sql
-- Retorna cada livro individualmente
SELECT titulo, quantidade_disponivel
FROM livros;
```

**Aggregate Query (retorna valor resumido):**
```sql
-- Retorna o total de livros disponíveis
SELECT SUM(quantidade_disponivel) AS total_livros
FROM livros;
```

---

## 2. Funções de Agregação Fundamentais

As funções de agregação são o coração das aggregate queries. Elas processam um conjunto de valores e retornam um único resultado.

### 2.1 COUNT() - Contar Registros

A função **COUNT()** retorna o número de linhas que correspondem a uma condição específica.

#### Sintaxe Básica

```sql
COUNT(*)
COUNT(coluna)
COUNT(DISTINCT coluna)
```

#### COUNT(*)

Conta **todas as linhas**, incluindo aquelas com valores NULL.

```sql
-- Contar todos os livros na tabela
SELECT COUNT(*) AS total_livros
FROM livros;
```

**Resultado**: Retorna o número total de registros na tabela `livros`.

#### COUNT(coluna)

Conta apenas linhas onde a coluna **não é NULL**.

```sql
-- Contar livros com ano de publicação informado
SELECT COUNT(ano_publicacao) AS livros_com_ano
FROM livros;
```

**Diferença importante**:
- `COUNT(*)` conta todas as linhas, mesmo com NULL
- `COUNT(coluna)` ignora linhas onde a coluna é NULL

#### COUNT(DISTINCT coluna)

Conta apenas **valores únicos** (distintos) na coluna.

```sql
-- Contar quantos autores únicos temos
SELECT COUNT(DISTINCT autor_id) AS total_autores
FROM livros;
```

#### Exemplos Práticos com COUNT

```sql
-- Total de livros na biblioteca
SELECT COUNT(*) AS total_livros FROM livros;

-- Livros com estoque disponível
SELECT COUNT(*) AS livros_disponiveis
FROM livros
WHERE quantidade_disponivel > 0;

-- Autores únicos que têm livros cadastrados
SELECT COUNT(DISTINCT autor_id) AS total_autores
FROM livros
WHERE autor_id IS NOT NULL;

-- Empréstimos ativos
SELECT COUNT(*) AS emprestimos_ativos
FROM emprestimos
WHERE status = 'ativo';
```

---

### 2.2 SUM() - Somar Valores

A função **SUM()** calcula a soma total de valores numéricos em uma coluna.

#### Sintaxe Básica

```sql
SUM(coluna)
```

#### Características

- Funciona apenas com valores numéricos (INTEGER, REAL, NUMERIC)
- Ignora valores NULL (não os inclui no cálculo)
- Retorna NULL se todas as linhas tiverem NULL na coluna
- Pode ser usado com expressões matemáticas

#### Exemplos Práticos com SUM

```sql
-- Total de livros disponíveis em estoque
SELECT SUM(quantidade_disponivel) AS total_estoque
FROM livros;

-- Soma de livros disponíveis apenas de ficção científica
SELECT SUM(quantidade_disponivel) AS estoque_ficcao
FROM livros
WHERE categoria_id = 1;

-- Total de livros disponíveis por categoria
SELECT 
    c.nome AS categoria,
    SUM(l.quantidade_disponivel) AS total_estoque
FROM livros l
JOIN categorias c ON l.categoria_id = c.id
GROUP BY c.id, c.nome;
```

#### SUM com Expressões

```sql
-- Calcular valor total se cada livro tivesse um preço
-- (exemplo hipotético: preço = quantidade * 10)
SELECT SUM(quantidade_disponivel * 10) AS valor_total_hipotetico
FROM livros;
```

---

### 2.3 AVG() - Calcular Média

A função **AVG()** calcula a média aritmética dos valores em uma coluna numérica.

#### Sintaxe Básica

```sql
AVG(coluna)
```

#### Características

- Funciona apenas com valores numéricos
- Ignora valores NULL no cálculo
- Retorna NULL se todas as linhas tiverem NULL
- O resultado é sempre um número decimal (REAL), mesmo para colunas INTEGER

#### Como AVG Calcula

```
AVG = SUM(valores) / COUNT(valores não-nulos)
```

#### Exemplos Práticos com AVG

```sql
-- Média de livros disponíveis por livro
SELECT AVG(quantidade_disponivel) AS media_estoque
FROM livros;

-- Média de ano de publicação
SELECT AVG(ano_publicacao) AS ano_medio_publicacao
FROM livros
WHERE ano_publicacao IS NOT NULL;

-- Média de estoque por categoria
SELECT 
    c.nome AS categoria,
    AVG(l.quantidade_disponivel) AS media_estoque
FROM livros l
JOIN categorias c ON l.categoria_id = c.id
GROUP BY c.id, c.nome;
```

#### Arredondamento com AVG

```sql
-- Média arredondada para 2 casas decimais
SELECT ROUND(AVG(quantidade_disponivel), 2) AS media_estoque
FROM livros;
```

---

### 2.4 MIN() - Valor Mínimo

A função **MIN()** retorna o menor valor em uma coluna.

#### Sintaxe Básica

```sql
MIN(coluna)
```

#### Características

- Funciona com números, datas e strings
- Para números: retorna o menor valor numérico
- Para datas: retorna a data mais antiga
- Para strings: retorna o valor alfabeticamente primeiro
- Ignora valores NULL

#### Exemplos Práticos com MIN

```sql
-- Ano de publicação mais antigo
SELECT MIN(ano_publicacao) AS ano_mais_antigo
FROM livros
WHERE ano_publicacao IS NOT NULL;

-- Menor quantidade em estoque
SELECT MIN(quantidade_disponivel) AS menor_estoque
FROM livros;

-- Título alfabeticamente primeiro
SELECT MIN(titulo) AS primeiro_titulo
FROM livros;

-- Data de empréstimo mais antiga
SELECT MIN(data_emprestimo) AS primeiro_emprestimo
FROM emprestimos;
```

---

### 2.5 MAX() - Valor Máximo

A função **MAX()** retorna o maior valor em uma coluna.

#### Sintaxe Básica

```sql
MAX(coluna)
```

#### Características

- Funciona com números, datas e strings
- Para números: retorna o maior valor numérico
- Para datas: retorna a data mais recente
- Para strings: retorna o valor alfabeticamente último
- Ignora valores NULL

#### Exemplos Práticos com MAX

```sql
-- Ano de publicação mais recente
SELECT MAX(ano_publicacao) AS ano_mais_recente
FROM livros
WHERE ano_publicacao IS NOT NULL;

-- Maior quantidade em estoque
SELECT MAX(quantidade_disponivel) AS maior_estoque
FROM livros;

-- Título alfabeticamente último
SELECT MAX(titulo) AS ultimo_titulo
FROM livros;

-- Data de empréstimo mais recente
SELECT MAX(data_emprestimo) AS ultimo_emprestimo
FROM emprestimos;
```

---

### 2.6 Combinando Múltiplas Funções de Agregação

Você pode usar várias funções de agregação na mesma query:

```sql
-- Estatísticas completas sobre estoque
SELECT 
    COUNT(*) AS total_livros,
    SUM(quantidade_disponivel) AS total_estoque,
    AVG(quantidade_disponivel) AS media_estoque,
    MIN(quantidade_disponivel) AS menor_estoque,
    MAX(quantidade_disponivel) AS maior_estoque
FROM livros;
```

**Resultado esperado:**
```
total_livros | total_estoque | media_estoque | menor_estoque | maior_estoque
-------------|---------------|---------------|---------------|---------------
15           | 90            | 6.0           | 3             | 10
```

---

## 3. GROUP BY - Agrupando Dados

A cláusula **GROUP BY** agrupa linhas que compartilham os mesmos valores em colunas especificadas, permitindo aplicar funções de agregação a cada grupo separadamente.

### 3.1 Conceito de GROUP BY

**GROUP BY** organiza os dados em grupos baseado nos valores de uma ou mais colunas. Cada grupo contém todas as linhas que têm os mesmos valores nessas colunas.

### 3.2 Sintaxe Básica

```sql
SELECT coluna1, coluna2, FUNCAO_AGREGACAO(coluna3)
FROM tabela
GROUP BY coluna1, coluna2;
```

### 3.3 Regra Importante: Colunas no SELECT

**⚠️ REGRA CRÍTICA**: Todas as colunas no SELECT que **não são funções de agregação** devem aparecer no GROUP BY.

**✅ Correto:**
```sql
SELECT categoria_id, COUNT(*) AS total
FROM livros
GROUP BY categoria_id;
```

**❌ Errado:**
```sql
-- ERRO: titulo não está no GROUP BY
SELECT titulo, COUNT(*) AS total
FROM livros
GROUP BY categoria_id;
```

### 3.4 Exemplos Práticos com GROUP BY

#### Agrupar por Categoria

```sql
-- Contar livros por categoria
SELECT 
    c.nome AS categoria,
    COUNT(*) AS total_livros
FROM livros l
JOIN categorias c ON l.categoria_id = c.id
GROUP BY c.id, c.nome;
```

**Resultado esperado:**
```
categoria           | total_livros
--------------------|-------------
Ficção Científica   | 4
Romance             | 5
Técnico             | 2
História            | 2
Filosofia           | 1
Mistério            | 1
```

#### Agrupar por Autor

```sql
-- Total de livros e estoque por autor
SELECT 
    a.nome AS autor,
    COUNT(*) AS total_livros,
    SUM(l.quantidade_disponivel) AS total_estoque
FROM livros l
JOIN autores a ON l.autor_id = a.id
GROUP BY a.id, a.nome
ORDER BY total_livros DESC;
```

#### Agrupar por Múltiplas Colunas

```sql
-- Livros por categoria e autor
SELECT 
    c.nome AS categoria,
    a.nome AS autor,
    COUNT(*) AS total_livros
FROM livros l
JOIN categorias c ON l.categoria_id = c.id
JOIN autores a ON l.autor_id = a.id
GROUP BY c.id, c.nome, a.id, a.nome;
```

#### Agrupar por Ano

```sql
-- Estatísticas de livros por década de publicação
SELECT 
    (ano_publicacao / 10) * 10 AS decada,
    COUNT(*) AS total_livros,
    AVG(quantidade_disponivel) AS media_estoque
FROM livros
WHERE ano_publicacao IS NOT NULL
GROUP BY (ano_publicacao / 10) * 10
ORDER BY decada;
```

### 3.5 GROUP BY com WHERE

**WHERE** filtra linhas **antes** do agrupamento:

```sql
-- Contar livros disponíveis por categoria (apenas com estoque > 0)
SELECT 
    c.nome AS categoria,
    COUNT(*) AS total_livros_disponiveis
FROM livros l
JOIN categorias c ON l.categoria_id = c.id
WHERE l.quantidade_disponivel > 0
GROUP BY c.id, c.nome;
```

**Ordem de execução:**
1. FROM: Seleciona as tabelas
2. JOIN: Faz os relacionamentos
3. WHERE: Filtra as linhas
4. GROUP BY: Agrupa as linhas filtradas
5. SELECT: Aplica funções de agregação

---

## 4. HAVING - Filtrando Grupos

A cláusula **HAVING** filtra os resultados de **GROUP BY** baseado em condições aplicadas às funções de agregação.

### 4.1 Diferença entre WHERE e HAVING

| Aspecto | WHERE | HAVING |
|---------|-------|--------|
| **Quando aplica** | Antes do agrupamento | Depois do agrupamento |
| **O que filtra** | Linhas individuais | Grupos completos |
| **Pode usar funções de agregação?** | ❌ Não | ✅ Sim |
| **Pode usar colunas normais?** | ✅ Sim | ✅ Sim (mas geralmente usa agregações) |

### 4.2 Sintaxe Básica

```sql
SELECT coluna1, FUNCAO_AGREGACAO(coluna2)
FROM tabela
GROUP BY coluna1
HAVING FUNCAO_AGREGACAO(coluna2) > valor;
```

### 4.3 Exemplos Práticos com HAVING

#### Filtrar Grupos por Contagem

```sql
-- Categorias com mais de 2 livros
SELECT 
    c.nome AS categoria,
    COUNT(*) AS total_livros
FROM livros l
JOIN categorias c ON l.categoria_id = c.id
GROUP BY c.id, c.nome
HAVING COUNT(*) > 2;
```

#### Filtrar Grupos por Soma

```sql
-- Autores com mais de 10 livros em estoque
SELECT 
    a.nome AS autor,
    SUM(l.quantidade_disponivel) AS total_estoque
FROM livros l
JOIN autores a ON l.autor_id = a.id
GROUP BY a.id, a.nome
HAVING SUM(l.quantidade_disponivel) > 10;
```

#### Filtrar Grupos por Média

```sql
-- Categorias com média de estoque acima de 5
SELECT 
    c.nome AS categoria,
    AVG(l.quantidade_disponivel) AS media_estoque
FROM livros l
JOIN categorias c ON l.categoria_id = c.id
GROUP BY c.id, c.nome
HAVING AVG(l.quantidade_disponivel) > 5;
```

#### Combinar WHERE e HAVING

```sql
-- Autores com mais de 1 livro publicado depois de 2000
SELECT 
    a.nome AS autor,
    COUNT(*) AS total_livros
FROM livros l
JOIN autores a ON l.autor_id = a.id
WHERE l.ano_publicacao > 2000
GROUP BY a.id, a.nome
HAVING COUNT(*) > 1;
```

**Ordem de execução:**
1. FROM / JOIN
2. WHERE (filtra linhas)
3. GROUP BY (agrupa linhas filtradas)
4. HAVING (filtra grupos)
5. SELECT (aplica funções e retorna)

### 4.4 HAVING com Múltiplas Condições

```sql
-- Categorias com mais de 2 livros E média de estoque acima de 5
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

## 5. Exemplos Práticos Combinados

### Exemplo 1: Relatório Completo de Categorias

```sql
-- Estatísticas completas por categoria
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

### Exemplo 2: Análise de Empréstimos por Usuário

```sql
-- Usuários com mais empréstimos ativos
SELECT 
    u.nome AS usuario,
    COUNT(*) AS total_emprestimos_ativos
FROM emprestimos e
JOIN usuarios u ON e.usuario_id = u.id
WHERE e.status = 'ativo'
GROUP BY u.id, u.nome
ORDER BY total_emprestimos_ativos DESC;
```

### Exemplo 3: Autores Mais Produtivos

```sql
-- Top 5 autores com mais livros
SELECT 
    a.nome AS autor,
    COUNT(*) AS total_livros,
    SUM(l.quantidade_disponivel) AS total_estoque
FROM livros l
JOIN autores a ON l.autor_id = a.id
GROUP BY a.id, a.nome
ORDER BY total_livros DESC
LIMIT 5;
```

### Exemplo 4: Análise Temporal de Empréstimos

```sql
-- Empréstimos por mês (apenas ativos)
SELECT 
    strftime('%Y-%m', data_emprestimo) AS mes,
    COUNT(*) AS total_emprestimos
FROM emprestimos
WHERE status = 'ativo'
GROUP BY strftime('%Y-%m', data_emprestimo)
ORDER BY mes;
```

### Exemplo 5: Categorias com Baixo Estoque

```sql
-- Categorias que precisam de reposição (média de estoque < 5)
SELECT 
    c.nome AS categoria,
    COUNT(*) AS total_livros,
    AVG(l.quantidade_disponivel) AS media_estoque
FROM livros l
JOIN categorias c ON l.categoria_id = c.id
GROUP BY c.id, c.nome
HAVING AVG(l.quantidade_disponivel) < 5
ORDER BY media_estoque;
```

---

## 6. Aliases em Aggregate Queries

Aliases tornam os resultados mais legíveis e podem ser usados em ORDER BY e HAVING (em alguns SGBDs).

### Exemplos com Aliases

```sql
-- Usando aliases para melhorar legibilidade
SELECT 
    c.nome AS categoria,
    COUNT(*) AS total,
    SUM(l.quantidade_disponivel) AS estoque_total,
    AVG(l.quantidade_disponivel) AS estoque_medio
FROM livros l
JOIN categorias c ON l.categoria_id = c.id
GROUP BY c.id, c.nome
HAVING estoque_total > 10  -- Usando alias no HAVING (SQLite suporta)
ORDER BY estoque_medio DESC;  -- Usando alias no ORDER BY
```

---

## 7. Tratamento de NULL em Funções de Agregação

### Comportamento com NULL

Todas as funções de agregação **ignoram valores NULL**, exceto COUNT(*):

```sql
-- COUNT(*) conta todas as linhas, incluindo NULL
SELECT COUNT(*) FROM livros;  -- Conta todos os livros

-- COUNT(coluna) ignora NULL
SELECT COUNT(ano_publicacao) FROM livros;  -- Conta apenas livros com ano

-- SUM, AVG, MIN, MAX ignoram NULL
SELECT 
    AVG(ano_publicacao) AS media_ano,
    MIN(ano_publicacao) AS ano_minimo,
    MAX(ano_publicacao) AS ano_maximo
FROM livros;
```

### Filtrar NULL antes de Agregar

```sql
-- Calcular média apenas de livros com ano informado
SELECT AVG(ano_publicacao) AS media_ano
FROM livros
WHERE ano_publicacao IS NOT NULL;
```

---

## 8. Ordem das Cláusulas em Aggregate Queries

A ordem correta das cláusulas em uma query com agregação é:

```sql
SELECT colunas, funcoes_agregacao
FROM tabela
[JOIN outras_tabelas]
[WHERE condicoes_linhas]
GROUP BY colunas
[HAVING condicoes_grupos]
[ORDER BY colunas]
[LIMIT numero];
```

**Ordem de execução:**
1. **FROM**: Seleciona as tabelas
2. **JOIN**: Faz os relacionamentos
3. **WHERE**: Filtra linhas individuais
4. **GROUP BY**: Agrupa as linhas
5. **HAVING**: Filtra os grupos
6. **SELECT**: Aplica funções e projeta colunas
7. **ORDER BY**: Ordena os resultados
8. **LIMIT**: Limita o número de resultados

---

## 9. Erros Comuns e Como Evitá-los

### Erro 1: Coluna não agrupada no SELECT

**❌ Errado:**
```sql
SELECT titulo, COUNT(*) 
FROM livros 
GROUP BY categoria_id;
```

**✅ Correto:**
```sql
SELECT categoria_id, COUNT(*) 
FROM livros 
GROUP BY categoria_id;
```

### Erro 2: Usar função de agregação no WHERE

**❌ Errado:**
```sql
SELECT categoria_id, COUNT(*) 
FROM livros 
WHERE COUNT(*) > 5
GROUP BY categoria_id;
```

**✅ Correto:**
```sql
SELECT categoria_id, COUNT(*) 
FROM livros 
GROUP BY categoria_id
HAVING COUNT(*) > 5;
```

### Erro 3: Esquecer GROUP BY ao usar agregação

**❌ Errado:**
```sql
SELECT categoria_id, COUNT(*) 
FROM livros;
```

**✅ Correto:**
```sql
SELECT categoria_id, COUNT(*) 
FROM livros
GROUP BY categoria_id;
```

---

## Conclusão

Nesta aula, você aprendeu:

1. **O que são Aggregate Queries**: Queries que processam múltiplas linhas e retornam valores resumidos
2. **Funções de Agregação Fundamentais**:
   - **COUNT()**: Contar registros
   - **SUM()**: Somar valores numéricos
   - **AVG()**: Calcular média
   - **MIN()**: Encontrar valor mínimo
   - **MAX()**: Encontrar valor máximo
3. **GROUP BY**: Agrupar dados por categorias para análise segmentada
4. **HAVING**: Filtrar grupos baseado em condições agregadas
5. **Tratamento de NULL**: Como as funções lidam com valores nulos
6. **Ordem das cláusulas**: Sequência correta em queries com agregação
7. **Erros comuns**: Como evitar os erros mais frequentes

Aggregate queries são fundamentais para análise de dados, geração de relatórios e tomada de decisões baseada em dados. Pratique muito executando essas queries no banco de dados da biblioteca.

**Próximo Passo**: Na próxima seção, vamos simplificar esses conceitos usando analogias do dia a dia para facilitar o entendimento.

---

**⚠️ Lembrete Importante**: 
- Sempre use GROUP BY quando usar funções de agregação com colunas não agregadas
- Use WHERE para filtrar linhas, HAVING para filtrar grupos
- Lembre-se que funções de agregação ignoram NULL (exceto COUNT(*))
