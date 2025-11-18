# Aula 2: Sintaxe Básica de SQL

## Introdução

Nesta aula, você aprenderá os fundamentos da sintaxe SQL. Esses são os blocos de construção essenciais que permitirão que você interaja com qualquer banco de dados relacional. Dominar esses conceitos é crucial para todo o aprendizado subsequente em SQL.

---

## 1. Sintaxe Básica de SQL

SQL (Structured Query Language) segue uma estrutura relativamente simples e consistente. A maioria dos comandos SQL segue um padrão geral:

```sql
COMANDO [cláusula] [condição] [ordenação] [limite]
```

### Características da Sintaxe SQL

1. **Case-Insensitive**: SQL não diferencia maiúsculas de minúsculas para palavras-chave e nomes de objetos (em muitos SGBDs). No entanto, é uma convenção escrever palavras-chave em MAIÚSCULAS.

2. **Terminadores**: Em SQLite, comandos são separados por ponto e vírgula (`;`), embora em modo interativo isso seja opcional.

3. **Comentários**: 
   - Comentários de linha única: `-- comentário`
   - Comentários de múltiplas linhas: `/* comentário */`

4. **Aspas**: Strings (texto) devem estar entre aspas simples (`'texto'`), não aspas duplas.

**Exemplo de sintaxe básica:**
```sql
-- Comentário explicativo
SELECT titulo, ano_publicacao 
FROM livros 
WHERE ano_publicacao > 2000 
ORDER BY ano_publicacao DESC 
LIMIT 10;
```

---

## 2. Palavras-Chave SQL (Keywords)

Palavras-chave são termos reservados que têm significado especial em SQL. Elas formam a estrutura dos comandos e não podem ser usadas como nomes de tabelas, colunas ou outros identificadores sem aspas.

### Categorias de Keywords

#### 2.1 Comandos Principais (DML - Data Manipulation Language)

- **SELECT**: Recupera dados de uma ou mais tabelas
- **INSERT**: Adiciona novos registros a uma tabela
- **UPDATE**: Modifica registros existentes
- **DELETE**: Remove registros de uma tabela

#### 2.2 Cláusulas de Consulta

- **FROM**: Especifica a tabela ou tabelas de onde os dados serão recuperados
- **WHERE**: Filtra registros baseado em condições
- **ORDER BY**: Ordena os resultados
- **GROUP BY**: Agrupa registros para funções de agregação
- **HAVING**: Filtra grupos (usado com GROUP BY)
- **LIMIT**: Limita o número de registros retornados
- **OFFSET**: Especifica quantos registros pular antes de começar a retornar

#### 2.3 Cláusulas de Modificação

- **SET**: Especifica quais colunas atualizar e seus novos valores (usado com UPDATE)
- **VALUES**: Especifica os valores a serem inseridos (usado com INSERT)

#### 2.4 Operadores e Modificadores

- **AS**: Cria um alias (apelido) para colunas ou tabelas
- **DISTINCT**: Remove duplicatas dos resultados
- **AND, OR, NOT**: Operadores lógicos
- **IN, NOT IN**: Verifica se um valor está em uma lista
- **BETWEEN**: Verifica se um valor está em um intervalo
- **LIKE**: Busca padrões em strings
- **IS NULL, IS NOT NULL**: Verifica valores nulos

#### 2.5 Joins e Relacionamentos

- **JOIN**: Combina dados de múltiplas tabelas
- **INNER JOIN**: Retorna apenas registros que têm correspondência em ambas as tabelas
- **LEFT JOIN**: Retorna todos os registros da tabela à esquerda e correspondências da direita
- **RIGHT JOIN**: Retorna todos os registros da tabela à direita e correspondências da esquerda
- **ON**: Especifica a condição de junção

#### 2.6 Funções de Agregação

- **COUNT**: Conta o número de registros
- **SUM**: Soma valores numéricos
- **AVG**: Calcula a média
- **MAX**: Retorna o valor máximo
- **MIN**: Retorna o valor mínimo

**Exemplo usando várias keywords:**
```sql
SELECT 
    a.nome AS autor,
    COUNT(l.id) AS total_livros
FROM autores a
LEFT JOIN livros l ON a.id = l.autor_id
WHERE a.nacionalidade = 'Brasileiro'
GROUP BY a.id, a.nome
HAVING COUNT(l.id) > 0
ORDER BY total_livros DESC
LIMIT 5;
```

---

## 3. Tipos de Dados SQL

Tipos de dados definem o tipo de valor que pode ser armazenado em uma coluna. Cada SGBD tem suas variações, mas os conceitos são similares. No SQLite, os tipos são mais flexíveis.

### 3.1 Tipos Numéricos

#### INTEGER
Armazena números inteiros (positivos e negativos).

```sql
-- Exemplos de valores INTEGER
ano_publicacao INTEGER  -- 1899, 2000, 2023
quantidade_disponivel INTEGER  -- 0, 5, 100, -10
```

**Características:**
- Tamanho: 1, 2, 3, 4, 6 ou 8 bytes (dependendo do valor)
- Intervalo: -9,223,372,036,854,775,808 a 9,223,372,036,854,775,807 (8 bytes)
- Uso: IDs, contadores, anos, quantidades

#### REAL
Armazena números de ponto flutuante (decimais).

```sql
-- Exemplo de valores REAL
preco REAL  -- 29.99, 150.50, 0.99
```

**Características:**
- Precisão: Aproximadamente 15 dígitos significativos
- Uso: Preços, medidas, porcentagens

#### NUMERIC / DECIMAL
Armazena números decimais com precisão fixa.

```sql
-- Exemplo (sintaxe varia por SGBD)
preco DECIMAL(10, 2)  -- 10 dígitos totais, 2 após a vírgula
-- Valores: 12345678.90, 99.99, 0.50
```

**Características:**
- Precisão exata (não arredonda como REAL)
- Uso: Valores monetários, medidas precisas

### 3.2 Tipos de Texto

#### TEXT
Armazena strings de tamanho variável.

```sql
-- Exemplos de valores TEXT
titulo TEXT  -- 'Dom Casmurro', 'Fundação'
nome TEXT  -- 'Machado de Assis'
descricao TEXT  -- 'Romance brasileiro...'
```

**Características:**
- Tamanho: Variável, limitado pela capacidade do banco
- No SQLite: Pode armazenar até 1 bilhão de caracteres
- Uso: Nomes, descrições, conteúdo textual

#### VARCHAR(n) / CHAR(n)
Em outros SGBDs, você pode especificar tamanho máximo.

```sql
-- Exemplo (PostgreSQL, MySQL)
nome VARCHAR(100)  -- Máximo 100 caracteres
codigo CHAR(10)    -- Exatamente 10 caracteres (preenche com espaços)
```

**Diferença:**
- **VARCHAR**: Tamanho variável até o máximo
- **CHAR**: Tamanho fixo (preenche com espaços se necessário)

### 3.3 Tipos de Data e Hora

#### DATE
Armazena apenas a data (sem hora).

```sql
-- Exemplos de valores DATE
data_nascimento DATE  -- '1899-06-21', '1942-01-02'
data_cadastro DATE    -- '2023-01-15'
```

**Formato padrão:** `YYYY-MM-DD` (ISO 8601)

#### DATETIME / TIMESTAMP
Armazena data e hora.

```sql
-- Exemplo
data_emprestimo TIMESTAMP  -- '2023-01-15 14:30:00'
```

**Formato padrão:** `YYYY-MM-DD HH:MM:SS`

**No SQLite:**
- DATE, DATETIME e TIMESTAMP são armazenados como TEXT, INTEGER ou REAL
- Funções de data trabalham com strings no formato ISO 8601

### 3.4 Tipos Booleanos

#### BOOLEAN / BOOL
Armazena valores verdadeiro/falso.

```sql
-- Exemplo (em alguns SGBDs)
ativo BOOLEAN  -- TRUE, FALSE
```

**No SQLite:**
- Não há tipo BOOLEAN nativo
- Usa-se INTEGER: 0 (false) ou 1 (true)
- Ou TEXT: 'true'/'false', 'yes'/'no'

### 3.5 Tipos Binários

#### BLOB (Binary Large Object)
Armazena dados binários (imagens, arquivos, etc.).

```sql
-- Exemplo
foto BLOB  -- Dados binários de uma imagem
```

**Uso:** Imagens, documentos, arquivos binários

### 3.6 Valores Nulos (NULL)

**NULL** representa a ausência de valor. É diferente de zero ou string vazia.

```sql
-- Uma coluna pode aceitar NULL
telefone TEXT  -- Pode ser NULL se o usuário não informou
```

**Características:**
- NULL não é igual a NULL (use `IS NULL` ou `IS NOT NULL`)
- Qualquer operação com NULL resulta em NULL
- Colunas podem ser definidas como `NOT NULL` para não aceitar valores nulos

### Exemplo Prático com Tipos de Dados

Olhando nosso banco de dados:

```sql
-- Estrutura da tabela livros
CREATE TABLE livros (
    id INTEGER PRIMARY KEY,              -- INTEGER
    titulo TEXT NOT NULL,                 -- TEXT
    isbn TEXT UNIQUE,                     -- TEXT (pode ser NULL)
    ano_publicacao INTEGER,               -- INTEGER (pode ser NULL)
    editora TEXT,                         -- TEXT (pode ser NULL)
    quantidade_disponivel INTEGER DEFAULT 0  -- INTEGER com valor padrão
);
```

---

## 4. Operadores SQL

Operadores são símbolos ou palavras-chave usadas para realizar operações em dados.

### 4.1 Operadores Aritméticos

Usados para realizar cálculos matemáticos.

| Operador | Descrição | Exemplo |
|----------|-----------|---------|
| `+` | Adição | `quantidade + 1` |
| `-` | Subtração | `preco - desconto` |
| `*` | Multiplicação | `quantidade * preco` |
| `/` | Divisão | `total / quantidade` |
| `%` | Módulo (resto da divisão) | `id % 2` |

**Exemplo:**
```sql
-- Calcular o total de livros disponíveis
SELECT 
    titulo,
    quantidade_disponivel,
    quantidade_disponivel * 2 AS duplicado
FROM livros
WHERE quantidade_disponivel > 0;
```

### 4.2 Operadores de Comparação

Usados para comparar valores em condições WHERE.

| Operador | Descrição | Exemplo |
|----------|-----------|---------|
| `=` | Igual a | `ano_publicacao = 2000` |
| `!=` ou `<>` | Diferente de | `status != 'ativo'` |
| `<` | Menor que | `ano_publicacao < 2000` |
| `>` | Maior que | `ano_publicacao > 2000` |
| `<=` | Menor ou igual a | `quantidade <= 10` |
| `>=` | Maior ou igual a | `quantidade >= 5` |

**Exemplo:**
```sql
-- Livros publicados depois de 2000
SELECT titulo, ano_publicacao
FROM livros
WHERE ano_publicacao > 2000;

-- Livros com quantidade disponível entre 1 e 10
SELECT titulo, quantidade_disponivel
FROM livros
WHERE quantidade_disponivel >= 1 
  AND quantidade_disponivel <= 10;
```

### 4.3 Operadores Lógicos

Combinam múltiplas condições.

| Operador | Descrição | Exemplo |
|----------|-----------|---------|
| `AND` | Ambas condições devem ser verdadeiras | `ano > 2000 AND quantidade > 0` |
| `OR` | Pelo menos uma condição deve ser verdadeira | `status = 'ativo' OR status = 'pendente'` |
| `NOT` | Inverte o resultado da condição | `NOT status = 'ativo'` |

**Exemplo:**
```sql
-- Livros brasileiros OU publicados depois de 2000
SELECT l.titulo, a.nome, l.ano_publicacao
FROM livros l
JOIN autores a ON l.autor_id = a.id
WHERE a.nacionalidade = 'Brasileiro' 
   OR l.ano_publicacao > 2000;

-- Livros que NÃO são de ficção científica E têm estoque
SELECT titulo, quantidade_disponivel
FROM livros l
JOIN categorias c ON l.categoria_id = c.id
WHERE NOT c.nome = 'Ficção Científica'
  AND l.quantidade_disponivel > 0;
```

### 4.4 Operadores de Padrão (Pattern Matching)

#### LIKE
Busca padrões em strings usando wildcards.

| Wildcard | Descrição | Exemplo |
|----------|-----------|---------|
| `%` | Qualquer sequência de caracteres (0 ou mais) | `'Dom%'` encontra 'Dom Casmurro', 'Dom Quixote' |
| `_` | Qualquer caractere único | `'Dom _asmurro'` encontra 'Dom Casmurro' |

**Exemplo:**
```sql
-- Livros cujo título começa com "Dom"
SELECT titulo
FROM livros
WHERE titulo LIKE 'Dom%';

-- Autores cujo nome contém "Silva"
SELECT nome
FROM autores
WHERE nome LIKE '%Silva%';

-- Títulos com exatamente 5 caracteres
SELECT titulo
FROM livros
WHERE titulo LIKE '_____';  -- 5 underscores
```

**NOT LIKE**: Inverte a condição LIKE.

### 4.5 Operadores de Conjunto

#### IN / NOT IN
Verifica se um valor está em uma lista.

```sql
-- Livros de categorias específicas
SELECT titulo
FROM livros
WHERE categoria_id IN (1, 2, 3);

-- Equivale a:
SELECT titulo
FROM livros
WHERE categoria_id = 1 
   OR categoria_id = 2 
   OR categoria_id = 3;

-- Autores que NÃO são brasileiros ou americanos
SELECT nome, nacionalidade
FROM autores
WHERE nacionalidade NOT IN ('Brasileiro', 'Americano');
```

#### BETWEEN
Verifica se um valor está em um intervalo (inclusivo).

```sql
-- Livros publicados entre 1990 e 2000
SELECT titulo, ano_publicacao
FROM livros
WHERE ano_publicacao BETWEEN 1990 AND 2000;

-- Equivale a:
SELECT titulo, ano_publicacao
FROM livros
WHERE ano_publicacao >= 1990 
  AND ano_publicacao <= 2000;

-- NOT BETWEEN: Fora do intervalo
SELECT titulo, ano_publicacao
FROM livros
WHERE ano_publicacao NOT BETWEEN 1990 AND 2000;
```

### 4.6 Operadores de Nulos

#### IS NULL / IS NOT NULL
Verifica valores nulos (NULL não pode ser comparado com `=` ou `!=`).

```sql
-- Livros sem ano de publicação
SELECT titulo
FROM livros
WHERE ano_publicacao IS NULL;

-- Livros com ano de publicação informado
SELECT titulo, ano_publicacao
FROM livros
WHERE ano_publicacao IS NOT NULL;
```

**⚠️ IMPORTANTE**: Nunca use `= NULL` ou `!= NULL`. Sempre use `IS NULL` ou `IS NOT NULL`.

---

## 5. Comando SELECT

O comando SELECT é o mais fundamental em SQL. Ele permite recuperar dados de uma ou mais tabelas.

### 5.1 Sintaxe Básica

```sql
SELECT coluna1, coluna2, ...
FROM nome_tabela
[WHERE condição]
[ORDER BY coluna [ASC|DESC]]
[LIMIT número];
```

### 5.2 SELECT com Todas as Colunas

Use `*` para selecionar todas as colunas:

```sql
-- Seleciona todas as colunas de livros
SELECT * FROM livros;
```

**⚠️ Cuidado**: Em produção, evite `SELECT *` quando possível. Especifique as colunas necessárias para melhor performance.

### 5.3 SELECT com Colunas Específicas

```sql
-- Seleciona apenas título e ano
SELECT titulo, ano_publicacao
FROM livros;
```

### 5.4 SELECT com Aliases (AS)

Aliases dão nomes temporários a colunas ou tabelas:

```sql
-- Alias para coluna
SELECT 
    titulo AS nome_do_livro,
    ano_publicacao AS ano
FROM livros;

-- Alias para tabela (útil em JOINs)
SELECT l.titulo, a.nome
FROM livros l
JOIN autores a ON l.autor_id = a.id;
```

### 5.5 SELECT com DISTINCT

Remove duplicatas dos resultados:

```sql
-- Lista nacionalidades únicas dos autores
SELECT DISTINCT nacionalidade
FROM autores
WHERE nacionalidade IS NOT NULL;
```

### 5.6 SELECT com WHERE

Filtra registros baseado em condições:

```sql
-- Livros publicados depois de 2000
SELECT titulo, ano_publicacao
FROM livros
WHERE ano_publicacao > 2000;

-- Múltiplas condições
SELECT titulo, quantidade_disponivel
FROM livros
WHERE quantidade_disponivel > 0 
  AND ano_publicacao >= 2000;
```

### 5.7 SELECT com ORDER BY

Ordena os resultados:

```sql
-- Ordenar por título (crescente - padrão)
SELECT titulo, ano_publicacao
FROM livros
ORDER BY titulo;

-- Ordenar por ano (decrescente)
SELECT titulo, ano_publicacao
FROM livros
ORDER BY ano_publicacao DESC;

-- Ordenar por múltiplas colunas
SELECT titulo, ano_publicacao, quantidade_disponivel
FROM livros
ORDER BY ano_publicacao DESC, titulo ASC;
```

### 5.8 SELECT com LIMIT

Limita o número de registros retornados:

```sql
-- Primeiros 5 livros
SELECT titulo
FROM livros
LIMIT 5;

-- Pular 10 registros e pegar os próximos 5
SELECT titulo
FROM livros
ORDER BY titulo
LIMIT 5 OFFSET 10;
```

### 5.9 SELECT com Expressões

Você pode usar expressões matemáticas e funções:

```sql
-- Calcular idade dos livros (assumindo ano atual 2024)
SELECT 
    titulo,
    ano_publicacao,
    2024 - ano_publicacao AS idade_anos
FROM livros
WHERE ano_publicacao IS NOT NULL;
```

### 5.10 SELECT com Funções de String

```sql
-- Converter título para maiúsculas
SELECT UPPER(titulo) AS titulo_maiusculo
FROM livros;

-- Tamanho do título
SELECT 
    titulo,
    LENGTH(titulo) AS tamanho
FROM livros;
```

---

## 6. Comando INSERT

O comando INSERT adiciona novos registros a uma tabela.

### 6.1 Sintaxe Básica

```sql
INSERT INTO nome_tabela (coluna1, coluna2, ...)
VALUES (valor1, valor2, ...);
```

### 6.2 INSERT com Todas as Colunas

Se você especificar valores para todas as colunas na ordem correta, pode omitir os nomes das colunas:

```sql
-- Inserir em categorias (todas as colunas)
INSERT INTO categorias
VALUES (NULL, 'Biografia', 'Livros sobre a vida de pessoas');
```

**⚠️ Cuidado**: Esta forma é menos legível e pode quebrar se a estrutura da tabela mudar.

### 6.3 INSERT com Colunas Específicas (Recomendado)

```sql
-- Inserir novo autor
INSERT INTO autores (nome, nacionalidade, data_nascimento)
VALUES ('Clarice Lispector', 'Brasileira', '1920-12-10');

-- Inserir livro
INSERT INTO livros (titulo, isbn, ano_publicacao, autor_id, categoria_id, quantidade_disponivel)
VALUES ('A Hora da Estrela', '978-8535914842', 1977, 11, 2, 3);
```

### 6.4 INSERT Múltiplo

Alguns SGBDs permitem inserir múltiplos registros de uma vez:

```sql
-- Inserir múltiplos autores (SQLite suporta)
INSERT INTO autores (nome, nacionalidade)
VALUES 
    ('Jorge Amado', 'Brasileiro'),
    ('Érico Veríssimo', 'Brasileiro'),
    ('Graciliano Ramos', 'Brasileiro');
```

### 6.5 INSERT com Valores Padrão

Se uma coluna tem valor padrão ou aceita NULL, você pode omiti-la:

```sql
-- quantidade_disponivel tem DEFAULT 0, então pode ser omitida
INSERT INTO livros (titulo, autor_id, categoria_id)
VALUES ('Novo Livro', 1, 1);
-- quantidade_disponivel será automaticamente 0
```

### 6.6 INSERT com SELECT (Inserir de Outra Tabela)

Você pode inserir dados de outra tabela:

```sql
-- Criar uma tabela temporária com autores brasileiros
CREATE TABLE autores_brasileiros AS
SELECT id, nome, nacionalidade
FROM autores
WHERE nacionalidade = 'Brasileiro';
```

---

## 7. Comando UPDATE

O comando UPDATE modifica registros existentes em uma tabela.

### 7.1 Sintaxe Básica

```sql
UPDATE nome_tabela
SET coluna1 = valor1, coluna2 = valor2, ...
WHERE condição;
```

### 7.2 UPDATE Simples

```sql
-- Atualizar quantidade disponível de um livro específico
UPDATE livros
SET quantidade_disponivel = 10
WHERE id = 1;
```

### 7.3 UPDATE Múltiplas Colunas

```sql
-- Atualizar várias colunas de uma vez
UPDATE livros
SET 
    quantidade_disponivel = 5,
    editora = 'Editora Nova'
WHERE id = 2;
```

### 7.4 UPDATE com Expressões

Você pode usar expressões para calcular novos valores:

```sql
-- Aumentar quantidade disponível em 1
UPDATE livros
SET quantidade_disponivel = quantidade_disponivel + 1
WHERE id = 3;

-- Dobrar a quantidade disponível
UPDATE livros
SET quantidade_disponivel = quantidade_disponivel * 2
WHERE quantidade_disponivel > 0;
```

### 7.5 UPDATE com Múltiplas Condições

```sql
-- Atualizar livros de uma categoria específica
UPDATE livros
SET quantidade_disponivel = 0
WHERE categoria_id = 1 
  AND quantidade_disponivel > 0;
```

### 7.6 UPDATE com Subconsultas

```sql
-- Atualizar livros de um autor específico
UPDATE livros
SET quantidade_disponivel = 10
WHERE autor_id IN (
    SELECT id 
    FROM autores 
    WHERE nacionalidade = 'Brasileiro'
);
```

**⚠️ CRÍTICO**: Sempre use WHERE em UPDATE! Sem WHERE, TODOS os registros serão atualizados:

```sql
-- ⚠️ PERIGOSO - Atualiza TODOS os livros!
UPDATE livros
SET quantidade_disponivel = 0;
-- Isso zeraria o estoque de TODOS os livros!
```

---

## 8. Comando DELETE

O comando DELETE remove registros de uma tabela.

### 8.1 Sintaxe Básica

```sql
DELETE FROM nome_tabela
WHERE condição;
```

### 8.2 DELETE com Condição

```sql
-- Deletar um livro específico
DELETE FROM livros
WHERE id = 15;

-- Deletar livros sem estoque
DELETE FROM livros
WHERE quantidade_disponivel = 0;
```

### 8.3 DELETE com Múltiplas Condições

```sql
-- Deletar livros antigos sem estoque
DELETE FROM livros
WHERE ano_publicacao < 1900 
  AND quantidade_disponivel = 0;
```

### 8.4 DELETE com Subconsultas

```sql
-- Deletar empréstimos de um usuário específico
DELETE FROM emprestimos
WHERE usuario_id IN (
    SELECT id 
    FROM usuarios 
    WHERE email = 'usuario@exemplo.com'
);
```

**⚠️ CRÍTICO**: Sempre use WHERE em DELETE! Sem WHERE, TODOS os registros serão deletados:

```sql
-- ⚠️ PERIGOSO - Deleta TODOS os livros!
DELETE FROM livros;
-- Isso apagaria TODOS os registros da tabela!
```

### 8.5 DELETE vs TRUNCATE

- **DELETE**: Remove registros linha por linha, pode ser revertido (em transações), mais lento
- **TRUNCATE**: Remove todos os registros de uma vez, mais rápido, não pode ser revertido facilmente

**No SQLite**: Não há comando TRUNCATE. Use `DELETE FROM tabela;` ou recrie a tabela.

### 8.6 Considerações sobre DELETE

1. **Integridade Referencial**: Se houver chaves estrangeiras, você pode não conseguir deletar um registro que é referenciado por outros:

```sql
-- Isso pode falhar se houver empréstimos deste livro
DELETE FROM livros
WHERE id = 1;
-- Erro: FOREIGN KEY constraint failed
```

2. **Cascata**: Alguns SGBDs permitem DELETE CASCADE para deletar registros relacionados automaticamente.

---

## Exemplos Práticos Combinados

### Exemplo 1: Consulta Completa

```sql
-- Buscar livros disponíveis de autores brasileiros, ordenados por título
SELECT 
    l.titulo AS livro,
    a.nome AS autor,
    l.ano_publicacao AS ano,
    l.quantidade_disponivel AS estoque
FROM livros l
JOIN autores a ON l.autor_id = a.id
WHERE a.nacionalidade = 'Brasileiro'
  AND l.quantidade_disponivel > 0
ORDER BY l.titulo;
```

### Exemplo 2: Inserção e Atualização

```sql
-- 1. Inserir novo autor
INSERT INTO autores (nome, nacionalidade)
VALUES ('Monteiro Lobato', 'Brasileiro');

-- 2. Inserir livro deste autor
INSERT INTO livros (titulo, autor_id, categoria_id, quantidade_disponivel)
VALUES ('O Sítio do Pica-Pau Amarelo', 
        (SELECT id FROM autores WHERE nome = 'Monteiro Lobato'),
        1,
        5);

-- 3. Atualizar quantidade após empréstimo
UPDATE livros
SET quantidade_disponivel = quantidade_disponivel - 1
WHERE titulo = 'O Sítio do Pica-Pau Amarelo';
```

### Exemplo 3: Limpeza de Dados

```sql
-- Deletar empréstimos antigos já devolvidos
DELETE FROM emprestimos
WHERE status = 'devolvido'
  AND data_devolucao_real < '2020-01-01';
```

---

## Conclusão

Nesta aula, você aprendeu:

1. **Sintaxe básica de SQL**: Estrutura geral dos comandos
2. **Palavras-chave**: Comandos, cláusulas e operadores principais
3. **Tipos de dados**: INTEGER, TEXT, DATE, etc.
4. **Operadores**: Aritméticos, comparação, lógicos, padrão
5. **SELECT**: Consultar dados com filtros, ordenação e limites
6. **INSERT**: Adicionar novos registros
7. **UPDATE**: Modificar registros existentes
8. **DELETE**: Remover registros

Esses são os fundamentos que você usará em todas as queries SQL. Pratique muito executando esses comandos no banco de dados da biblioteca.

**Próximo Passo**: Na próxima seção, vamos simplificar esses conceitos usando analogias do dia a dia para facilitar o entendimento.

---

**⚠️ Lembrete Importante**: Sempre use WHERE em UPDATE e DELETE, ou você pode modificar/deletar dados incorretos!

