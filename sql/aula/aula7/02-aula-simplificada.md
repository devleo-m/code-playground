# Aula 7 - Simplificada: Entendendo Sub Queries

## Introdução

Imagine que você está fazendo uma pergunta, mas para responder essa pergunta, você precisa fazer outra pergunta primeiro. Por exemplo:

**Pergunta Principal**: "Quais livros têm estoque acima da média?"

**Pergunta Auxiliar (que precisa ser respondida primeiro)**: "Qual é a média de estoque de todos os livros?"

Essa é exatamente a ideia por trás das **Sub Queries** (Subconsultas) em SQL: elas são como "perguntas dentro de perguntas" - você faz uma consulta para obter um resultado que será usado em outra consulta.

**Subqueries são como caixas dentro de caixas**: você abre uma caixa (query principal) e dentro dela encontra outra caixa (subquery) que precisa ser aberta primeiro para descobrir o conteúdo!

---

## 1. Sub Queries: A Analogia das Perguntas Aninhadas

### Pensando em Queries como Perguntas

Imagine que você é um bibliotecário e precisa responder perguntas complexas:

**Pergunta Simples (sem subquery):**
```
Você: "Quantos livros temos?"
Sistema: "15 livros"
```

**Pergunta Complexa (com subquery):**
```
Você: "Quais livros têm mais empréstimos que a média?"
Sistema: "Deixa eu ver... primeiro preciso saber qual é a média..."
       → Calcula média de empréstimos (subquery)
       → Compara cada livro com essa média (query principal)
       → "Fundação, 1984, Dom Casmurro"
```

### Por que Precisamos de Subqueries?

**Sem Subquery (trabalhoso):**
```
1. Você faz uma query: "Qual é a média de estoque?"
   → Resultado: 5 livros

2. Você faz outra query: "Quais livros têm mais de 5?"
   → Resultado: Lista de livros

3. Você combina os resultados manualmente ❌
```

**Com Subquery (automático):**
```
1. Você faz uma query com subquery:
   "Quais livros têm estoque > (média de estoque)?"
   
2. O sistema calcula tudo automaticamente ✅
```

**Subqueries permitem que você faça perguntas complexas de uma só vez!**

---

## 2. Scalar Subquery: Uma Única Resposta

### Analogia: Perguntar um Número Único

Pense em uma **Scalar Subquery** como fazer uma pergunta que tem uma **única resposta numérica**:

**Exemplo do dia a dia:**
```
Você: "Qual é a média de idade dos funcionários?"
Sistema: "32 anos" ← Uma única resposta (escalar)
```

**No SQL:**
```sql
-- Qual é a média de estoque?
SELECT AVG(quantidade_disponivel) FROM livros;
-- Resultado: 5.2 (um único número)
```

### Usando Scalar Subquery em uma Query Principal

**Analogia: Comparar com um Padrão**

Imagine que você quer saber quais funcionários são mais velhos que a média:

```
1. Pergunta auxiliar: "Qual é a média de idade?" → 32 anos
2. Pergunta principal: "Quais funcionários têm mais de 32 anos?"
```

**No SQL:**
```sql
-- Quais livros têm estoque acima da média?
SELECT titulo
FROM livros
WHERE quantidade_disponivel > (
    SELECT AVG(quantidade_disponivel) FROM livros  -- Subquery: retorna 5.2
);
```

**O que acontece:**
```
Livro 1: estoque = 3 → 3 > 5.2? ❌ Não
Livro 2: estoque = 7 → 7 > 5.2? ✅ Sim (aparece no resultado)
Livro 3: estoque = 6 → 6 > 5.2? ✅ Sim (aparece no resultado)
```

### Exemplo Prático: Contar Empréstimos

**Analogia: Adicionar uma Coluna Calculada**

Imagine que você tem uma lista de livros e quer adicionar uma coluna mostrando quantas vezes cada livro foi emprestado:

```sql
SELECT 
    titulo,
    (SELECT COUNT(*) 
     FROM emprestimos 
     WHERE emprestimos.livro_id = livros.id) AS total_emprestimos
FROM livros;
```

**O que acontece:**
```
Para cada livro:
1. A subquery conta quantos empréstimos esse livro tem
2. O resultado é adicionado como uma nova coluna
3. Você vê: "Fundação - 3 empréstimos"
```

**É como adicionar uma nota em cada item de uma lista!**

---

## 3. Column Subquery: Uma Lista de Respostas

### Analogia: Perguntar uma Lista

Pense em uma **Column Subquery** como fazer uma pergunta que retorna uma **lista de valores**:

**Exemplo do dia a dia:**
```
Você: "Quais são os IDs dos autores brasileiros?"
Sistema: "1, 7, 9" ← Uma lista de valores
```

**No SQL:**
```sql
-- Quais são os IDs dos autores brasileiros?
SELECT id FROM autores WHERE nacionalidade = 'Brasileiro';
-- Resultado: 1, 7, 9 (uma lista)
```

### Usando Column Subquery com IN

**Analogia: Verificar se Está na Lista**

Imagine que você tem uma lista de convidados VIP e quer verificar quem está na lista:

```
Lista VIP: [João, Maria, Pedro]

Pessoa 1: "João" → Está na lista? ✅ Sim
Pessoa 2: "Ana" → Está na lista? ❌ Não
Pessoa 3: "Maria" → Está na lista? ✅ Sim
```

**No SQL:**
```sql
-- Quais livros são de autores brasileiros?
SELECT titulo
FROM livros
WHERE autor_id IN (
    SELECT id FROM autores WHERE nacionalidade = 'Brasileiro'
    -- Subquery retorna: 1, 7, 9
);
```

**O que acontece:**
```
Livro 1: autor_id = 1 → 1 está em [1, 7, 9]? ✅ Sim
Livro 2: autor_id = 5 → 5 está em [1, 7, 9]? ❌ Não
Livro 3: autor_id = 7 → 7 está em [1, 7, 9]? ✅ Sim
```

### Operadores para Listas

**IN - "Está na lista?"**
```sql
WHERE autor_id IN (SELECT id FROM autores WHERE nacionalidade = 'Brasileiro')
-- "O autor está na lista de brasileiros?"
```

**NOT IN - "NÃO está na lista?"**
```sql
WHERE id NOT IN (SELECT DISTINCT livro_id FROM emprestimos)
-- "O livro NÃO está na lista de livros emprestados?"
```

**EXISTS - "A lista tem pelo menos um item?"**
```sql
WHERE EXISTS (SELECT 1 FROM emprestimos WHERE livro_id = livros.id)
-- "Existe pelo menos um empréstimo para este livro?"
```

---

## 4. Table Subquery: Uma Tabela Completa

### Analogia: Criar uma Tabela Temporária

Pense em uma **Table Subquery** como criar uma **tabela temporária** que você usa como se fosse uma tabela normal:

**Exemplo do dia a dia:**
```
Você: "Crie uma lista de estatísticas por categoria"
Sistema: Cria uma tabela temporária:
         Categoria | Total Livros | Total Empréstimos
         Ficção    | 5            | 12
         Romance   | 3            | 8
```

**No SQL:**
```sql
-- Usar uma tabela temporária (subquery) como se fosse uma tabela normal
SELECT categoria, total_livros
FROM (
    SELECT 
        c.nome AS categoria,
        COUNT(l.id) AS total_livros
    FROM categorias c
    LEFT JOIN livros l ON c.id = l.categoria_id
    GROUP BY c.id, c.nome
) AS estatisticas  -- Esta é a "tabela temporária"
WHERE total_livros > 2;
```

**O que acontece:**
```
1. A subquery cria uma tabela temporária com estatísticas
2. A query principal usa essa tabela como se fosse uma tabela normal
3. Você pode filtrar, ordenar, etc. nessa tabela temporária
```

**É como criar uma planilha intermediária e depois usar ela para fazer cálculos!**

---

## 5. Nested Subqueries: Perguntas Dentro de Perguntas Dentro de Perguntas

### Analogia: Matryoshka (Bonecas Russas)

Pense em **Nested Subqueries** como **bonecas russas** (matryoshka): uma boneca dentro de outra, dentro de outra:

```
Boneca Externa (Query Principal)
  └─ Boneca Média (Subquery 1)
      └─ Boneca Pequena (Subquery 2)
          └─ Boneca Mínima (Subquery 3)
```

**Exemplo do dia a dia:**
```
Pergunta Principal: "Quais livros são de autores que têm mais livros que a média?"

Para responder, preciso:
1. Pergunta 1: "Qual é a média de livros por autor?"
   → Para responder isso, preciso:
      2. Pergunta 2: "Quantos livros cada autor tem?"
          → Para responder isso, preciso:
              3. Pergunta 3: "Liste todos os autores e conte seus livros"
```

**No SQL:**
```sql
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

**⚠️ CUIDADO**: Quanto mais aninhadas, mais difícil de entender! Às vezes é melhor usar JOINs.

---

## 6. Correlated Subqueries: Perguntas que Dependem da Resposta Anterior

### Analogia: Comparar Cada Item com Seu Grupo

Pense em uma **Correlated Subquery** como fazer uma pergunta que **muda para cada item**:

**Exemplo do dia a dia:**
```
Você está olhando para cada funcionário e perguntando:
"Este funcionário ganha mais que a média do seu departamento?"

Funcionário 1 (TI): "Ganho R$ 5000. A média de TI é R$ 4500?" → ✅ Sim
Funcionário 2 (TI): "Ganho R$ 4000. A média de TI é R$ 4500?" → ❌ Não
Funcionário 3 (RH): "Ganho R$ 3000. A média de RH é R$ 2500?" → ✅ Sim
```

**A pergunta muda dependendo do funcionário!**

### Diferença: Subquery Normal vs Correlated

**Subquery Normal (não-correlacionada):**
```sql
-- Executada UMA vez
SELECT titulo
FROM livros
WHERE quantidade_disponivel > (
    SELECT AVG(quantidade_disponivel) FROM livros
    -- Esta média é a mesma para TODOS os livros
);
```

**O que acontece:**
```
1. Calcula média geral: 5.2
2. Compara TODOS os livros com 5.2
3. Retorna os que têm mais que 5.2
```

**Correlated Subquery:**
```sql
-- Executada para CADA livro
SELECT titulo
FROM livros l1
WHERE quantidade_disponivel > (
    SELECT AVG(quantidade_disponivel)
    FROM livros l2
    WHERE l2.categoria_id = l1.categoria_id
    -- Esta média muda para CADA categoria!
);
```

**O que acontece:**
```
Para Livro 1 (Ficção):
1. Calcula média de Ficção: 6.0
2. Compara Livro 1 com 6.0

Para Livro 2 (Romance):
1. Calcula média de Romance: 4.0
2. Compara Livro 2 com 4.0

Para Livro 3 (Ficção):
1. Calcula média de Ficção: 6.0 (já calculada antes, mas recalcula)
2. Compara Livro 3 com 6.0
```

**⚠️ ATENÇÃO**: Correlated subqueries podem ser **muito lentas** porque são executadas muitas vezes!

---

## 7. Subqueries vs JOINs: Quando Usar Cada Um?

### Analogia: Diferentes Ferramentas para Diferentes Trabalhos

Pense em **Subqueries** e **JOINs** como **ferramentas diferentes**:

- **JOINs**: Como uma **chave de fenda** - perfeita para juntar peças
- **Subqueries**: Como uma **chave inglesa** - perfeita para ajustes e comparações

### Quando Usar Subqueries

**Cenário 1: Comparar com um Valor Calculado**
```
Pergunta: "Quais livros têm estoque acima da média?"
→ Você precisa calcular a média primeiro
→ Subquery é perfeita para isso!
```

```sql
SELECT titulo
FROM livros
WHERE quantidade_disponivel > (
    SELECT AVG(quantidade_disponivel) FROM livros
);
```

**Cenário 2: Verificar Existência**
```
Pergunta: "Quais livros têm empréstimos ativos?"
→ Você só precisa saber se existe (não precisa dos dados do empréstimo)
→ EXISTS é perfeito!
```

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

### Quando Usar JOINs

**Cenário 1: Combinar Dados de Múltiplas Tabelas**
```
Pergunta: "Mostre livros com seus autores"
→ Você precisa de dados de ambas as tabelas
→ JOIN é perfeito!
```

```sql
SELECT l.titulo, a.nome AS autor
FROM livros l
JOIN autores a ON l.autor_id = a.id;
```

**Cenário 2: Performance**
```
Quando você tem tabelas grandes e precisa de performance
→ JOINs geralmente são mais rápidos
→ Use JOINs!
```

### Regra de Ouro

- **Subqueries**: Para **filtros**, **comparações** e **valores únicos**
- **JOINs**: Para **combinar** e **relacionar** dados de múltiplas tabelas
- **Ambos funcionam**: Escolha baseado em **legibilidade** e **performance**

---

## 8. Exemplos Práticos do Dia a Dia

### Exemplo 1: Encontrar "Outliers" (Valores Extremos)

**Analogia**: Encontrar alunos que tiraram nota muito acima da média da turma

```sql
-- Livros com estoque muito acima da média
SELECT titulo, quantidade_disponivel
FROM livros
WHERE quantidade_disponivel > (
    SELECT AVG(quantidade_disponivel) * 1.5 FROM livros
    -- 1.5 vezes a média = muito acima
);
```

**O que faz**: Encontra livros que têm 50% mais estoque que a média (valores extremos)

### Exemplo 2: Encontrar "Melhores do Grupo"

**Analogia**: Encontrar o melhor aluno de cada turma

```sql
-- Livro mais emprestado de cada categoria
SELECT 
    l1.titulo,
    c.nome AS categoria
FROM livros l1
JOIN categorias c ON l1.categoria_id = c.id
WHERE (
    SELECT COUNT(*)
    FROM emprestimos e
    WHERE e.livro_id = l1.id
) = (
    SELECT MAX(total)
    FROM (
        SELECT COUNT(*) AS total
        FROM emprestimos e2
        JOIN livros l2 ON e2.livro_id = l2.id
        WHERE l2.categoria_id = l1.categoria_id
        GROUP BY l2.id
    )
);
```

**O que faz**: Para cada categoria, encontra o livro com mais empréstimos

### Exemplo 3: Verificar "Completude"

**Analogia**: Verificar se todos os itens de uma lista têm algo

```sql
-- Categorias onde TODOS os livros têm estoque acima da média geral
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
```

**O que faz**: Encontra categorias onde nenhum livro tem estoque abaixo da média (todos estão acima)

---

## 9. Armadilhas Comuns e Como Evitá-las

### Armadilha 1: Subquery Retorna Múltiplas Linhas

**Problema**: Você espera um único valor, mas a subquery retorna vários

```sql
-- ❌ ERRO: Subquery retorna vários valores
SELECT titulo
FROM livros
WHERE quantidade_disponivel = (
    SELECT quantidade_disponivel FROM livros  -- Retorna várias linhas!
);
```

**Solução**: Use operadores adequados (IN, ANY, ALL) ou adicione LIMIT 1

```sql
-- ✅ CORRETO
SELECT titulo
FROM livros
WHERE quantidade_disponivel IN (
    SELECT quantidade_disponivel FROM livros WHERE categoria_id = 1
);
```

### Armadilha 2: Correlated Subquery Muito Lenta

**Problema**: Correlated subquery executada muitas vezes

```sql
-- ❌ Pode ser muito lento em tabelas grandes
SELECT titulo
FROM livros l1
WHERE (
    SELECT COUNT(*)
    FROM emprestimos e
    WHERE e.livro_id = l1.id
) > 5;
```

**Solução**: Reescreva com JOIN quando possível

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

### Armadilha 3: Subquery Desnecessária

**Problema**: Usar subquery quando um JOIN resolve melhor

```sql
-- ❌ Subquery desnecessária
SELECT titulo
FROM livros
WHERE autor_id IN (
    SELECT id FROM autores WHERE nacionalidade = 'Brasileiro'
);
```

**Solução**: Use JOIN quando você precisa de dados de ambas as tabelas

```sql
-- ✅ Mais eficiente e flexível
SELECT l.titulo, a.nome
FROM livros l
JOIN autores a ON l.autor_id = a.id
WHERE a.nacionalidade = 'Brasileiro';
```

---

## 10. Dicas Finais

### Dica 1: Teste a Subquery Separadamente

Sempre teste a subquery sozinha primeiro para garantir que ela retorna o que você espera:

```sql
-- Teste primeiro:
SELECT AVG(quantidade_disponivel) FROM livros;
-- Resultado: 5.2 ✅

-- Depois use na query principal:
SELECT titulo FROM livros WHERE quantidade_disponivel > 5.2;
```

### Dica 2: Use Aliases Claros

Dê nomes descritivos para facilitar a leitura:

```sql
-- ✅ BOM: Aliases claros
SELECT l.titulo
FROM livros l
WHERE l.quantidade_disponivel > (
    SELECT AVG(l2.quantidade_disponivel) 
    FROM livros l2
);

-- ❌ RUIM: Sem aliases ou confusos
SELECT titulo
FROM livros
WHERE quantidade_disponivel > (
    SELECT AVG(quantidade_disponivel) FROM livros
);
```

### Dica 3: Evite Aninhamento Excessivo

Se você tem mais de 3 níveis de aninhamento, considere reescrever:

```sql
-- ❌ Muito aninhado (difícil de entender)
SELECT ... FROM ... WHERE ... IN (
    SELECT ... FROM ... WHERE ... IN (
        SELECT ... FROM ... WHERE ... IN (
            SELECT ... FROM ...
        )
    )
);

-- ✅ Melhor: Use JOINs ou quebre em múltiplas queries
```

### Dica 4: Pense em Performance

Para tabelas grandes:
- Prefira JOINs quando possível
- Evite correlated subqueries se houver alternativa
- Use índices nas colunas usadas nas subqueries

---

## Conclusão

Subqueries são como **perguntas auxiliares** que ajudam a responder perguntas principais mais complexas. Elas são poderosas e flexíveis, mas devem ser usadas com sabedoria:

- ✅ **Use subqueries** para filtros dinâmicos, comparações e valores únicos
- ✅ **Use JOINs** para combinar dados de múltiplas tabelas
- ✅ **Teste sempre** suas subqueries separadamente
- ✅ **Pense em performance** especialmente com tabelas grandes
- ✅ **Mantenha simples** - evite aninhamento excessivo

**Lembre-se**: Subqueries são uma ferramenta poderosa, mas nem sempre são a melhor solução. Escolha a ferramenta certa para cada trabalho!

**Próximos Passos**:
1. Pratique os exemplos desta aula
2. Complete os exercícios práticos
3. Experimente reescrever subqueries como JOINs e vice-versa
4. Estude as boas práticas de performance

**Bons estudos! 🚀**





