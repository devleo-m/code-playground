# Aula 6: SQL JOIN Queries (Consultas com JOIN)

## Introdução

Nesta aula, você aprenderá sobre **SQL JOIN Queries** (Consultas com JOIN), uma das funcionalidades mais importantes e poderosas do SQL. JOINs permitem combinar dados de duas ou mais tabelas em uma única query, estabelecendo relacionamentos entre tabelas e permitindo análises complexas que seriam impossíveis com uma única tabela.

SQL JOIN queries são essenciais para:
- Combinar dados relacionados de múltiplas tabelas
- Estabelecer relacionamentos entre entidades
- Realizar análises complexas que envolvem várias tabelas
- Recuperar informações completas que estão distribuídas em diferentes tabelas
- Criar relatórios e visualizações que dependem de dados relacionados

Dominar JOINs é fundamental para qualquer desenvolvedor ou analista de dados, pois a maioria das consultas reais em bancos de dados relacionais requer combinar dados de múltiplas tabelas.

---

## 1. O que são JOINs?

**JOIN** é uma operação SQL que combina linhas de duas ou mais tabelas baseado em uma condição de relacionamento entre elas. A condição geralmente compara uma coluna de uma tabela com uma coluna de outra tabela, frequentemente usando chaves primárias e estrangeiras.

### Por que Precisamos de JOINs?

Em bancos de dados relacionais, os dados são normalizados e distribuídos em múltiplas tabelas para evitar redundância e manter a integridade. No entanto, frequentemente precisamos ver dados combinados de várias tabelas.

**Exemplo do Problema:**

Imagine que você quer ver uma lista de livros com seus autores. Os dados estão em duas tabelas:

**Tabela `livros`:**
```
id | titulo              | autor_id
1  | Fundação            | 1
2  | 1984                | 2
3  | Dom Casmurro        | 7
```

**Tabela `autores`:**
```
id | nome
1  | Isaac Asimov
2  | George Orwell
7  | Machado de Assis
```

**Sem JOIN**, você teria que fazer duas queries separadas e combinar os resultados manualmente. **Com JOIN**, você pode obter tudo em uma única query:

```sql
SELECT l.titulo, a.nome AS autor
FROM livros l
JOIN autores a ON l.autor_id = a.id;
```

**Resultado:**
```
titulo       | autor
Fundação     | Isaac Asimov
1984         | George Orwell
Dom Casmurro | Machado de Assis
```

### Componentes de um JOIN

Um JOIN possui três componentes principais:

1. **Tabelas a serem unidas**: Duas ou mais tabelas
2. **Condição de JOIN**: A condição que determina como as tabelas são relacionadas
3. **Tipo de JOIN**: Determina quais linhas são incluídas no resultado

### Sintaxe Básica

```sql
SELECT colunas
FROM tabela1
[TIPO] JOIN tabela2 ON condição
[WHERE filtros]
[ORDER BY ordenação];
```

---

## 2. INNER JOIN

O **INNER JOIN** é o tipo de JOIN mais comum. Ele retorna apenas as linhas que têm correspondência em ambas as tabelas. Se não houver correspondência, a linha não aparece no resultado.

### Características do INNER JOIN

1. **Retorna apenas correspondências**: Apenas linhas com match em ambas as tabelas
2. **É o JOIN padrão**: Se você escrever apenas `JOIN`, o SQL assume `INNER JOIN`
3. **Exclui linhas sem correspondência**: Linhas sem match não aparecem no resultado
4. **Simétrico**: A ordem das tabelas não importa (mas pode afetar performance)

### Sintaxe

```sql
SELECT colunas
FROM tabela1
INNER JOIN tabela2 ON tabela1.coluna = tabela2.coluna;
```

Ou simplesmente:

```sql
SELECT colunas
FROM tabela1
JOIN tabela2 ON tabela1.coluna = tabela2.coluna;
```

### Exemplo Prático: Livros e Autores

```sql
-- Listar todos os livros com seus autores
SELECT 
    l.titulo,
    l.ano_publicacao,
    a.nome AS autor,
    a.nacionalidade
FROM livros l
INNER JOIN autores a ON l.autor_id = a.id;
```

**O que acontece:**
- Para cada livro na tabela `livros`, o SQL procura o autor correspondente na tabela `autores`
- Se encontrar (match), a linha é incluída no resultado
- Se não encontrar (sem match), a linha é excluída

### Exemplo: Livros, Autores e Categorias

```sql
-- Listar livros com autor e categoria
SELECT 
    l.titulo,
    a.nome AS autor,
    c.nome AS categoria
FROM livros l
INNER JOIN autores a ON l.autor_id = a.id
INNER JOIN categorias c ON l.categoria_id = c.id;
```

**Observações:**
- Podemos fazer múltiplos JOINs em uma única query
- Cada JOIN adiciona uma nova tabela ao resultado
- A condição de cada JOIN é independente

### Exemplo: Empréstimos com Informações Completas

```sql
-- Listar empréstimos ativos com informações do livro e usuário
SELECT 
    e.id AS emprestimo_id,
    u.nome AS usuario,
    l.titulo AS livro,
    e.data_emprestimo,
    e.data_devolucao_prevista
FROM emprestimos e
INNER JOIN usuarios u ON e.usuario_id = u.id
INNER JOIN livros l ON e.livro_id = l.id
WHERE e.status = 'ativo'
ORDER BY e.data_emprestimo DESC;
```

### Quando Usar INNER JOIN

Use INNER JOIN quando:
- Você precisa apenas de registros que têm correspondência em ambas as tabelas
- Você quer excluir registros órfãos (sem relacionamento)
- Você está trabalhando com dados obrigatórios (todos os livros têm autor)
- É o caso mais comum na maioria das situações

---

## 3. LEFT JOIN (LEFT OUTER JOIN)

O **LEFT JOIN** (ou **LEFT OUTER JOIN**) retorna **todas as linhas da tabela esquerda** (primeira tabela) e as linhas correspondentes da tabela direita (segunda tabela). Se não houver correspondência na tabela direita, as colunas da tabela direita aparecem como `NULL`.

### Características do LEFT JOIN

1. **Inclui todas as linhas da tabela esquerda**: Mesmo sem correspondência
2. **Colunas da tabela direita podem ser NULL**: Quando não há match
3. **A ordem importa**: LEFT JOIN não é simétrico
4. **Útil para encontrar registros órfãos**: Registros sem relacionamento

### Sintaxe

```sql
SELECT colunas
FROM tabela1
LEFT JOIN tabela2 ON tabela1.coluna = tabela2.coluna;
```

### Exemplo Prático: Categorias com Livros

```sql
-- Listar todas as categorias e seus livros (mesmo categorias sem livros)
SELECT 
    c.nome AS categoria,
    l.titulo AS livro
FROM categorias c
LEFT JOIN livros l ON c.id = l.categoria_id;
```

**Resultado possível:**
```
categoria          | livro
Ficção Científica  | Fundação
Ficção Científica  | Eu, Robô
Romance            | 1984
Romance            | Dom Casmurro
Filosofia          | NULL  ← Categoria sem livros!
```

### Exemplo: Encontrar Categorias sem Livros

```sql
-- Encontrar categorias que não têm livros cadastrados
SELECT 
    c.id,
    c.nome AS categoria
FROM categorias c
LEFT JOIN livros l ON c.id = l.categoria_id
WHERE l.id IS NULL;
```

**Como funciona:**
- LEFT JOIN inclui todas as categorias
- Para categorias sem livros, `l.id` será `NULL`
- `WHERE l.id IS NULL` filtra apenas essas categorias

### Exemplo: Autores com Contagem de Livros

```sql
-- Listar todos os autores e quantos livros cada um tem
SELECT 
    a.nome AS autor,
    COUNT(l.id) AS total_livros
FROM autores a
LEFT JOIN livros l ON a.id = l.autor_id
GROUP BY a.id, a.nome
ORDER BY total_livros DESC;
```

**Observação importante:**
- Use `COUNT(l.id)` ao invés de `COUNT(*)` para contar apenas livros reais
- `COUNT(*)` contaria 1 mesmo para autores sem livros
- `COUNT(l.id)` conta apenas quando `l.id` não é NULL

### Quando Usar LEFT JOIN

Use LEFT JOIN quando:
- Você precisa de todos os registros da tabela esquerda
- Você quer incluir registros que podem não ter correspondência
- Você quer encontrar registros órfãos (sem relacionamento)
- Você está criando relatórios que devem incluir todas as opções

---

## 4. RIGHT JOIN (RIGHT OUTER JOIN)

O **RIGHT JOIN** (ou **RIGHT OUTER JOIN**) é o oposto do LEFT JOIN. Ele retorna **todas as linhas da tabela direita** (segunda tabela) e as linhas correspondentes da tabela esquerda (primeira tabela). Se não houver correspondência na tabela esquerda, as colunas da tabela esquerda aparecem como `NULL`.

### Características do RIGHT JOIN

1. **Inclui todas as linhas da tabela direita**: Mesmo sem correspondência
2. **Colunas da tabela esquerda podem ser NULL**: Quando não há match
3. **A ordem importa**: RIGHT JOIN não é simétrico
4. **Menos comum que LEFT JOIN**: Pode ser substituído por LEFT JOIN invertido

### Sintaxe

```sql
SELECT colunas
FROM tabela1
RIGHT JOIN tabela2 ON tabela1.coluna = tabela2.coluna;
```

### Exemplo Prático

```sql
-- Listar todos os livros e suas categorias (mesmo livros sem categoria)
SELECT 
    l.titulo AS livro,
    c.nome AS categoria
FROM categorias c
RIGHT JOIN livros l ON c.id = l.categoria_id;
```

**Nota:** Este exemplo é equivalente a:

```sql
-- Mesmo resultado usando LEFT JOIN
SELECT 
    l.titulo AS livro,
    c.nome AS categoria
FROM livros l
LEFT JOIN categorias c ON l.categoria_id = c.id;
```

### Quando Usar RIGHT JOIN

Use RIGHT JOIN quando:
- Você precisa de todos os registros da tabela direita
- A ordem das tabelas na query é importante para legibilidade
- Você prefere a sintaxe RIGHT JOIN para clareza

**Nota:** Muitos desenvolvedores preferem usar LEFT JOIN invertido ao invés de RIGHT JOIN, pois é mais comum e fácil de entender. SQLite **não suporta RIGHT JOIN**, então você sempre precisará usar LEFT JOIN invertido.

---

## 5. FULL OUTER JOIN (FULL JOIN)

O **FULL OUTER JOIN** (ou **FULL JOIN**) combina os resultados de LEFT JOIN e RIGHT JOIN. Ele retorna **todas as linhas de ambas as tabelas**, combinando-as quando há correspondência e preenchendo com `NULL` quando não há match.

### Características do FULL OUTER JOIN

1. **Inclui todas as linhas de ambas as tabelas**: Com ou sem correspondência
2. **Colunas podem ser NULL**: Quando não há match em uma das tabelas
3. **Útil para reconciliação de dados**: Ver todos os dados de ambas as tabelas
4. **Menos comum**: Geralmente usado em casos específicos

### Sintaxe

```sql
SELECT colunas
FROM tabela1
FULL OUTER JOIN tabela2 ON tabela1.coluna = tabela2.coluna;
```

### Exemplo Prático

```sql
-- Listar todos os autores e todos os livros (mesmo sem correspondência)
SELECT 
    a.nome AS autor,
    l.titulo AS livro
FROM autores a
FULL OUTER JOIN livros l ON a.id = l.autor_id;
```

**Resultado possível:**
```
autor              | livro
Isaac Asimov       | Fundação
Isaac Asimov       | Eu, Robô
George Orwell      | 1984
NULL               | Livro sem autor cadastrado
Autor sem livros   | NULL
```

### Simulando FULL OUTER JOIN no SQLite

**Nota importante:** SQLite **não suporta FULL OUTER JOIN**. Para obter o mesmo resultado, você pode usar UNION:

```sql
-- Simular FULL OUTER JOIN usando UNION
SELECT 
    a.nome AS autor,
    l.titulo AS livro
FROM autores a
LEFT JOIN livros l ON a.id = l.autor_id

UNION

SELECT 
    a.nome AS autor,
    l.titulo AS livro
FROM livros l
LEFT JOIN autores a ON l.autor_id = a.id
WHERE a.id IS NULL;
```

### Quando Usar FULL OUTER JOIN

Use FULL OUTER JOIN quando:
- Você precisa ver todos os dados de ambas as tabelas
- Você está fazendo reconciliação de dados
- Você quer identificar registros órfãos em ambas as tabelas
- Você está comparando duas fontes de dados

---

## 6. SELF JOIN

Um **SELF JOIN** é uma operação onde uma tabela é unida a si mesma. Isso é útil quando você precisa comparar linhas dentro da mesma tabela ou trabalhar com relacionamentos hierárquicos.

### Características do SELF JOIN

1. **Mesma tabela duas vezes**: A tabela aparece como "esquerda" e "direita"
2. **Requer aliases**: Você DEVE usar aliases diferentes para distinguir as instâncias
3. **Útil para comparações**: Comparar registros dentro da mesma tabela
4. **Relacionamentos hierárquicos**: Trabalhar com estruturas de árvore

### Sintaxe

```sql
SELECT colunas
FROM tabela alias1
JOIN tabela alias2 ON condição
WHERE filtros;
```

### Exemplo Prático: Autores da Mesma Nacionalidade

```sql
-- Encontrar pares de autores da mesma nacionalidade
SELECT 
    a1.nome AS autor1,
    a2.nome AS autor2,
    a1.nacionalidade
FROM autores a1
INNER JOIN autores a2 ON a1.nacionalidade = a2.nacionalidade
WHERE a1.id < a2.id  -- Evita duplicatas e auto-comparação
ORDER BY a1.nacionalidade, a1.nome;
```

**Como funciona:**
- `a1` e `a2` são duas "instâncias" da mesma tabela `autores`
- A condição `a1.nacionalidade = a2.nacionalidade` encontra autores da mesma nacionalidade
- `WHERE a1.id < a2.id` evita:
  - Duplicatas (autor1-autor2 e autor2-autor1)
  - Auto-comparação (autor1-autor1)

### Exemplo: Livros do Mesmo Autor

```sql
-- Listar todos os livros de Isaac Asimov com seus anos de publicação
SELECT 
    l1.titulo AS livro1,
    l1.ano_publicacao AS ano1,
    l2.titulo AS livro2,
    l2.ano_publicacao AS ano2
FROM livros l1
INNER JOIN livros l2 ON l1.autor_id = l2.autor_id
WHERE l1.autor_id = 1  -- Isaac Asimov
  AND l1.id < l2.id
ORDER BY l1.ano_publicacao, l2.ano_publicacao;
```

### Exemplo: Comparar Livros da Mesma Categoria

```sql
-- Encontrar livros da mesma categoria publicados no mesmo ano
SELECT 
    l1.titulo AS livro1,
    l2.titulo AS livro2,
    l1.categoria_id,
    l1.ano_publicacao
FROM livros l1
INNER JOIN livros l2 ON l1.categoria_id = l2.categoria_id 
    AND l1.ano_publicacao = l2.ano_publicacao
WHERE l1.id < l2.id
ORDER BY l1.categoria_id, l1.ano_publicacao;
```

### Quando Usar SELF JOIN

Use SELF JOIN quando:
- Você precisa comparar registros dentro da mesma tabela
- Você trabalha com estruturas hierárquicas (árvores)
- Você quer encontrar relacionamentos entre registros da mesma entidade
- Você precisa agrupar ou comparar registros similares

---

## 7. CROSS JOIN

O **CROSS JOIN** (também chamado de **produto cartesiano**) combina **cada linha da primeira tabela com cada linha da segunda tabela**, sem nenhuma condição de relacionamento. O resultado é o produto cartesiano das duas tabelas.

### Características do CROSS JOIN

1. **Sem condição de JOIN**: Não usa `ON`
2. **Produto cartesiano**: Cada linha da tabela1 com cada linha da tabela2
3. **Pode gerar muitos resultados**: Se tabela1 tem N linhas e tabela2 tem M linhas, o resultado tem N×M linhas
4. **Use com cuidado**: Pode ser muito custoso computacionalmente

### Sintaxe

```sql
SELECT colunas
FROM tabela1
CROSS JOIN tabela2;
```

Ou implicitamente:

```sql
SELECT colunas
FROM tabela1, tabela2;
```

### Exemplo Prático: Todas as Combinações

```sql
-- Gerar todas as combinações de categorias e autores
SELECT 
    c.nome AS categoria,
    a.nome AS autor
FROM categorias c
CROSS JOIN autores a;
```

**Resultado:**
Se você tem 6 categorias e 10 autores, o resultado terá 60 linhas (6 × 10).

### Quando CROSS JOIN é Útil

CROSS JOIN pode ser útil para:
- **Gerar dados de teste**: Criar todas as combinações possíveis
- **Tabelas de referência**: Combinar listas de valores
- **Análises combinatórias**: Quando você precisa de todas as combinações
- **Casos específicos**: Situações onde você realmente precisa do produto cartesiano

### Quando Evitar CROSS JOIN

**CUIDADO:** Na maioria dos casos, CROSS JOIN é um **erro acidental**:

```sql
-- ❌ ERRO COMUM: Esqueceu a condição ON
SELECT l.titulo, a.nome
FROM livros l
JOIN autores a;  -- Faltou ON! Isso vira CROSS JOIN!

-- ✅ CORRETO
SELECT l.titulo, a.nome
FROM livros l
JOIN autores a ON l.autor_id = a.id;
```

**Sempre verifique:**
- Se você realmente precisa de todas as combinações
- Se você esqueceu a condição `ON`
- Se o número de resultados faz sentido

---

## 8. Múltiplos JOINs

Você pode combinar múltiplos JOINs em uma única query para unir várias tabelas.

### Exemplo: Múltiplos INNER JOINs

```sql
-- Listar empréstimos com todas as informações relacionadas
SELECT 
    e.id AS emprestimo_id,
    u.nome AS usuario,
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
WHERE e.status = 'ativo'
ORDER BY e.data_emprestimo DESC;
```

### Exemplo: Misturando Tipos de JOIN

```sql
-- Listar todas as categorias com seus livros e autores (mesmo categorias sem livros)
SELECT 
    c.nome AS categoria,
    l.titulo AS livro,
    a.nome AS autor
FROM categorias c
LEFT JOIN livros l ON c.id = l.categoria_id
LEFT JOIN autores a ON l.autor_id = a.id
ORDER BY c.nome, l.titulo;
```

**Observação:**
- O primeiro JOIN é LEFT (inclui todas as categorias)
- O segundo JOIN também é LEFT (inclui categorias sem livros)
- Se uma categoria não tem livros, `l.titulo` e `a.nome` serão NULL

### Ordem dos JOINs

A ordem dos JOINs pode afetar:
1. **Performance**: Dependendo do otimizador do banco
2. **Legibilidade**: Quais relacionamentos são mais importantes
3. **Resultados**: Em alguns casos, a ordem pode importar

**Dica:** Comece com a tabela "principal" e vá adicionando JOINs conforme necessário.

---

## 9. Condições de JOIN vs WHERE

É importante entender a diferença entre colocar condições no `ON` (condição de JOIN) e no `WHERE` (filtro de resultado).

### Condição no ON

```sql
-- Filtra ANTES de fazer o JOIN
SELECT l.titulo, a.nome
FROM livros l
LEFT JOIN autores a ON l.autor_id = a.id AND a.nacionalidade = 'Brasileiro';
```

**Comportamento:**
- A condição `a.nacionalidade = 'Brasileiro'` é aplicada durante o JOIN
- Para LEFT JOIN, isso pode afetar quais linhas são incluídas

### Condição no WHERE

```sql
-- Filtra DEPOIS de fazer o JOIN
SELECT l.titulo, a.nome
FROM livros l
LEFT JOIN autores a ON l.autor_id = a.id
WHERE a.nacionalidade = 'Brasileiro';
```

**Comportamento:**
- O JOIN é feito primeiro
- Depois, o WHERE filtra os resultados
- Para LEFT JOIN, isso pode excluir linhas que você queria incluir

### Diferença Importante

```sql
-- LEFT JOIN com condição no ON
SELECT c.nome, COUNT(l.id) AS total
FROM categorias c
LEFT JOIN livros l ON c.id = l.categoria_id AND l.ano_publicacao > 2000
GROUP BY c.id, c.nome;
-- Inclui todas as categorias, mas conta apenas livros após 2000

-- LEFT JOIN com condição no WHERE
SELECT c.nome, COUNT(l.id) AS total
FROM categorias c
LEFT JOIN livros l ON c.id = l.categoria_id
WHERE l.ano_publicacao > 2000
GROUP BY c.id, c.nome;
-- Exclui categorias que não têm livros após 2000 (comporta como INNER JOIN)
```

**Regra geral:**
- **Condições de relacionamento** → `ON`
- **Filtros de resultado** → `WHERE`

---

## 10. Aliases de Tabelas

**Aliases** (apelidos) são essenciais ao trabalhar com JOINs. Eles tornam as queries mais legíveis e são obrigatórios em SELF JOINs.

### Sintaxe de Aliases

```sql
-- Alias curto
SELECT l.titulo, a.nome
FROM livros l
JOIN autores a ON l.autor_id = a.id;

-- Alias com AS (opcional, mas mais claro)
SELECT l.titulo, a.nome
FROM livros AS l
JOIN autores AS a ON l.autor_id = a.id;
```

### Boas Práticas de Aliases

1. **Use aliases curtos mas descritivos**: `l` para livros, `a` para autores
2. **Seja consistente**: Use os mesmos aliases em todo o projeto
3. **Evite aliases confusos**: Não use `a1`, `a2` a menos que seja SELF JOIN
4. **Prefira clareza**: Às vezes, nomes mais longos são melhores

### Exemplo com Aliases

```sql
-- Sem aliases (verboso)
SELECT 
    livros.titulo,
    autores.nome AS autor,
    categorias.nome AS categoria
FROM livros
INNER JOIN autores ON livros.autor_id = autores.id
INNER JOIN categorias ON livros.categoria_id = categorias.id;

-- Com aliases (mais limpo)
SELECT 
    l.titulo,
    a.nome AS autor,
    c.nome AS categoria
FROM livros l
INNER JOIN autores a ON l.autor_id = a.id
INNER JOIN categorias c ON l.categoria_id = c.id;
```

---

## 11. Resumo dos Tipos de JOIN

| Tipo de JOIN | Descrição | Quando Usar |
|--------------|-----------|-------------|
| **INNER JOIN** | Apenas linhas com correspondência em ambas as tabelas | Caso mais comum, quando você precisa apenas de dados relacionados |
| **LEFT JOIN** | Todas as linhas da tabela esquerda + correspondências da direita | Quando você quer incluir todos os registros da tabela principal |
| **RIGHT JOIN** | Todas as linhas da tabela direita + correspondências da esquerda | Menos comum, pode ser substituído por LEFT JOIN invertido |
| **FULL OUTER JOIN** | Todas as linhas de ambas as tabelas | Quando você precisa ver todos os dados de ambas as tabelas |
| **SELF JOIN** | Tabela unida a si mesma | Para comparar ou relacionar registros da mesma tabela |
| **CROSS JOIN** | Produto cartesiano (todas as combinações) | Raramente necessário, geralmente é um erro |

### Diagrama Visual (Venn)

```
INNER JOIN:     [A ∩ B]        (apenas interseção)
LEFT JOIN:      [A]            (tudo de A, interseção com B)
RIGHT JOIN:     [B]            (tudo de B, interseção com A)
FULL JOIN:      [A ∪ B]        (tudo de A e B)
CROSS JOIN:     [A × B]        (produto cartesiano)
```

---

## 12. Exemplos Práticos Completos

### Exemplo 1: Relatório de Empréstimos

```sql
-- Relatório completo de empréstimos
SELECT 
    e.id AS emprestimo_id,
    u.nome AS usuario,
    u.email,
    l.titulo AS livro,
    a.nome AS autor,
    c.nome AS categoria,
    e.data_emprestimo,
    e.data_devolucao_prevista,
    CASE 
        WHEN e.data_devolucao_real IS NULL THEN 'Em andamento'
        ELSE 'Devolvido'
    END AS status_emprestimo
FROM emprestimos e
INNER JOIN usuarios u ON e.usuario_id = u.id
INNER JOIN livros l ON e.livro_id = l.id
INNER JOIN autores a ON l.autor_id = a.id
INNER JOIN categorias c ON l.categoria_id = c.id
ORDER BY e.data_emprestimo DESC;
```

### Exemplo 2: Categorias com Estatísticas

```sql
-- Estatísticas por categoria
SELECT 
    c.nome AS categoria,
    COUNT(l.id) AS total_livros,
    SUM(l.quantidade_disponivel) AS total_estoque,
    COUNT(DISTINCT l.autor_id) AS total_autores
FROM categorias c
LEFT JOIN livros l ON c.id = l.categoria_id
GROUP BY c.id, c.nome
ORDER BY total_livros DESC;
```

### Exemplo 3: Usuários que Nunca Pegaram Livros

```sql
-- Usuários sem empréstimos
SELECT 
    u.id,
    u.nome,
    u.email,
    u.data_cadastro
FROM usuarios u
LEFT JOIN emprestimos e ON u.id = e.usuario_id
WHERE e.id IS NULL
ORDER BY u.data_cadastro DESC;
```

### Exemplo 4: Livros Mais Emprestados

```sql
-- Top 5 livros mais emprestados
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
```

### Exemplo 5: Autores com Todos os Seus Livros

```sql
-- Listar autores com todos os seus livros
SELECT 
    a.nome AS autor,
    a.nacionalidade,
    l.titulo AS livro,
    l.ano_publicacao
FROM autores a
LEFT JOIN livros l ON a.id = l.autor_id
ORDER BY a.nome, l.ano_publicacao;
```

---

## 13. Erros Comuns e Como Evitá-los

### Erro 1: Esquecer a Condição ON

```sql
-- ❌ ERRADO: CROSS JOIN acidental
SELECT l.titulo, a.nome
FROM livros l
JOIN autores a;  -- Faltou ON!

-- ✅ CORRETO
SELECT l.titulo, a.nome
FROM livros l
JOIN autores a ON l.autor_id = a.id;
```

### Erro 2: Usar WHERE ao Invés de ON para Relacionamento

```sql
-- ❌ Funciona, mas não é ideal
SELECT l.titulo, a.nome
FROM livros l, autores a
WHERE l.autor_id = a.id;

-- ✅ Melhor: usar JOIN explícito
SELECT l.titulo, a.nome
FROM livros l
JOIN autores a ON l.autor_id = a.id;
```

### Erro 3: Confundir LEFT JOIN com INNER JOIN

```sql
-- ❌ Se você quer categorias sem livros, não use INNER JOIN
SELECT c.nome, COUNT(l.id)
FROM categorias c
INNER JOIN livros l ON c.id = l.categoria_id
GROUP BY c.id, c.nome;
-- Isso exclui categorias sem livros!

-- ✅ Use LEFT JOIN
SELECT c.nome, COUNT(l.id)
FROM categorias c
LEFT JOIN livros l ON c.id = l.categoria_id
GROUP BY c.id, c.nome;
```

### Erro 4: Não Usar Aliases em SELF JOIN

```sql
-- ❌ ERRADO: Tabela sem alias
SELECT nome
FROM autores
JOIN autores ON nacionalidade = nacionalidade;

-- ✅ CORRETO: Aliases obrigatórios
SELECT a1.nome, a2.nome
FROM autores a1
JOIN autores a2 ON a1.nacionalidade = a2.nacionalidade
WHERE a1.id < a2.id;
```

### Erro 5: COUNT(*) em LEFT JOIN

```sql
-- ❌ Pode contar incorretamente
SELECT c.nome, COUNT(*) AS total
FROM categorias c
LEFT JOIN livros l ON c.id = l.categoria_id
GROUP BY c.id, c.nome;
-- COUNT(*) conta 1 mesmo para categorias sem livros

-- ✅ Conte a coluna da tabela direita
SELECT c.nome, COUNT(l.id) AS total
FROM categorias c
LEFT JOIN livros l ON c.id = l.categoria_id
GROUP BY c.id, c.nome;
```

---

## 14. Próximos Passos

Agora que você entende os diferentes tipos de JOIN, pratique:

1. **Execute todos os exemplos** no banco de dados `biblioteca.db`
2. **Experimente variações** dos exemplos
3. **Crie suas próprias queries** combinando diferentes tipos de JOIN
4. **Compare resultados** de diferentes tipos de JOIN
5. **Leia a aula simplificada** para reforçar o entendimento
6. **Complete os exercícios** para praticar

---

**Bons estudos! 🚀**

**Lembre-se**: JOINs são fundamentais para trabalhar com bancos de dados relacionais. Pratique muito e você dominará essa habilidade essencial!

