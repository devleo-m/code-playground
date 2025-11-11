# **Aula 3: Object Model in PostgreSQL**

## 🎯 Objetivo da Aula

Compreender o modelo de objetos do PostgreSQL (ORDBMS), aprender sobre a hierarquia de objetos (databases → schemas → tables → columns → rows), dominar os tipos de dados disponíveis e entender como fazer queries básicas.

---

## 🎭 PostgreSQL: O Híbrido ORDBMS

### O que significa ORDBMS?

**ORDBMS** = **Object-Relational Database Management System**

PostgreSQL **não é apenas** um banco relacional (RDBMS). É um **híbrido** que combina:

```
┌────────────────────────────────────────────────┐
│  RDBMS (Relacional)        OODBMS (Orientado   │
│  - Tabelas                  a Objetos)         │
│  - SQL                      - Tipos customizados│
│  - Chaves/Relacionamentos   - Herança          │
│  - Integridade ACID         - Polimorfismo     │
└────────────────────────────────────────────────┘
              ↓ COMBINAÇÃO ↓
         ┌─────────────────────┐
         │   PostgreSQL        │
         │   (ORDBMS)          │
         └─────────────────────┘
```

### Recursos Orientados a Objetos no PostgreSQL

#### 1. **Tipos de Dados Customizados**

Você pode criar seus próprios tipos de dados!

```sql
-- Criar um tipo customizado para endereço
CREATE TYPE endereco AS (
    rua VARCHAR(100),
    numero INTEGER,
    cidade VARCHAR(50),
    estado CHAR(2),
    cep CHAR(9)
);

-- Usar o tipo customizado em uma tabela
CREATE TABLE clientes (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(100),
    endereco endereco  -- Tipo customizado!
);

-- Inserir dados
INSERT INTO clientes (nome, endereco)
VALUES ('João', ROW('Rua A', 123, 'São Paulo', 'SP', '01234-567'));

-- Consultar
SELECT nome, (endereco).cidade FROM clientes;
```

#### 2. **Herança de Tabelas**

Tabelas podem herdar estrutura de outras tabelas!

```sql
-- Tabela pai (superclasse)
CREATE TABLE pessoas (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(100),
    data_nascimento DATE
);

-- Tabela filha (subclasse) - herda tudo de pessoas
CREATE TABLE funcionarios (
    salario DECIMAL(10, 2),
    cargo VARCHAR(50)
) INHERITS (pessoas);

-- Funcionarios tem: id, nome, data_nascimento (herdados) + salario, cargo

-- Inserir funcionário
INSERT INTO funcionarios (nome, data_nascimento, salario, cargo)
VALUES ('Maria', '1990-05-15', 5000.00, 'Desenvolvedora');

-- Consultar só funcionários
SELECT * FROM funcionarios;

-- Consultar todas as pessoas (inclui funcionários!)
SELECT * FROM pessoas;
```

#### 3. **Polimorfismo**

Você pode consultar tabelas pai e automaticamente incluir dados das tabelas filhas.

```sql
-- Continua do exemplo anterior
CREATE TABLE clientes_externos (
    empresa VARCHAR(100)
) INHERITS (pessoas);

-- Inserir cliente externo
INSERT INTO clientes_externos (nome, data_nascimento, empresa)
VALUES ('Pedro', '1985-03-20', 'Empresa XYZ');

-- Consultar TODAS as pessoas (funcionarios + clientes_externos)
SELECT * FROM pessoas;

-- Resultado inclui registros de pessoas, funcionarios e clientes_externos!
```

---

## 🏗️ Hierarquia de Objetos no PostgreSQL

PostgreSQL organiza dados em uma hierarquia clara:

```
┌─────────────────────────────────────────────────┐
│  SERVIDOR PostgreSQL                            │
│  ├─ DATABASE 1 (Ex: loja)                       │
│  │  ├─ SCHEMA public                            │
│  │  │  ├─ TABLE clientes                        │
│  │  │  │  ├─ COLUMN id                          │
│  │  │  │  ├─ COLUMN nome                        │
│  │  │  │  └─ ROW (1, 'João')                    │
│  │  │  ├─ TABLE pedidos                         │
│  │  │  └─ VIEW relatorio_vendas                 │
│  │  └─ SCHEMA vendas                            │
│  │     └─ TABLE comissoes                       │
│  └─ DATABASE 2 (Ex: blog)                       │
│     └─ SCHEMA public                            │
│        ├─ TABLE posts                           │
│        └─ TABLE comentarios                     │
└─────────────────────────────────────────────────┘
```

Vamos explorar cada nível!

---

## 🗄️ 1. Databases (Bancos de Dados)

Um **database** é uma coleção nomeada de objetos (tabelas, views, índices, etc.).

### Características

- Cada servidor PostgreSQL pode ter **múltiplos databases**
- Databases são **isolados** uns dos outros
- Você **não pode** fazer queries entre databases diferentes diretamente
- Cada database tem seu próprio conjunto de schemas

### Comandos Básicos

```sql
-- Listar todos os databases
\l
-- ou
SELECT datname FROM pg_database;

-- Criar um database
CREATE DATABASE loja;

CREATE DATABASE blog
    ENCODING 'UTF8'
    LC_COLLATE 'pt_BR.UTF-8'
    LC_CTYPE 'pt_BR.UTF-8';

-- Conectar a um database
\c loja
-- ou
\connect loja

-- Deletar um database (cuidado!)
DROP DATABASE blog;

-- Renomear database
ALTER DATABASE loja RENAME TO ecommerce;
```

### Quando Criar Múltiplos Databases?

✅ **Use databases separados quando:**

- Aplicações completamente diferentes (blog, loja, CRM)
- Dados de clientes diferentes (SaaS multi-tenant)
- Ambientes diferentes (produção, teste, desenvolvimento)

❌ **NÃO use databases separados quando:**

- Dados precisam se relacionar (use schemas no mesmo database)
- É apenas organização lógica (use schemas)

---

## 📂 2. Schemas

Um **schema** é um namespace dentro de um database que agrupa objetos relacionados.

### Por que usar Schemas?

```
DATABASE: loja
├─ SCHEMA: public (padrão)
│  ├─ clientes
│  └─ produtos
├─ SCHEMA: vendas
│  ├─ pedidos
│  └─ comissoes
└─ SCHEMA: estoque
   ├─ movimentacoes
   └─ inventario
```

**Benefícios:**

- **Organização**: Agrupa objetos relacionados
- **Namespacing**: Evita conflitos de nomes
- **Segurança**: Controle de acesso por schema
- **Multi-tenant**: Cada cliente em um schema diferente

### Comandos Básicos

```sql
-- Listar schemas
\dn
-- ou
SELECT schema_name FROM information_schema.schemata;

-- Criar schema
CREATE SCHEMA vendas;

CREATE SCHEMA estoque AUTHORIZATION usuario_estoque;

-- Definir schema padrão para a sessão
SET search_path TO vendas, public;

-- Criar tabela em schema específico
CREATE TABLE vendas.pedidos (
    id SERIAL PRIMARY KEY,
    total DECIMAL(10, 2)
);

-- Referenciar tabela de outro schema
SELECT * FROM vendas.pedidos;
SELECT * FROM estoque.produtos;

-- Deletar schema (vazio)
DROP SCHEMA vendas;

-- Deletar schema com todo conteúdo (cuidado!)
DROP SCHEMA vendas CASCADE;
```

### Schema Padrão: public

- Todo database tem um schema chamado `public` por padrão
- Se você não especificar schema, objetos são criados em `public`
- `search_path` define onde PostgreSQL procura objetos

```sql
-- Ver o search_path atual
SHOW search_path;
-- Resultado padrão: "$user", public

-- Criar tabela sem especificar schema (vai para 'public')
CREATE TABLE clientes (id SERIAL);
-- É o mesmo que:
CREATE TABLE public.clientes (id SERIAL);
```

---

## 📊 3. Tables (Tabelas)

Uma **table** é uma coleção de linhas organizadas em colunas.

### Anatomia de uma Tabela

```sql
CREATE TABLE produtos (
    -- Colunas (COLUMNS)
    id SERIAL PRIMARY KEY,
    nome VARCHAR(100) NOT NULL,
    preco DECIMAL(10, 2) CHECK (preco > 0),
    estoque INTEGER DEFAULT 0,
    categoria VARCHAR(50),
    criado_em TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

Cada tabela tem:

- **Nome**: Identificador único no schema
- **Colunas**: Definem estrutura (o que será armazenado)
- **Constraints**: Regras de integridade
- **Linhas**: Dados reais (registros)

### Comandos Básicos

```sql
-- Listar tabelas do schema atual
\dt
-- ou
SELECT table_name FROM information_schema.tables
WHERE table_schema = 'public';

-- Ver estrutura de uma tabela
\d produtos
-- ou
\d+ produtos  -- mais detalhes

-- Criar tabela
CREATE TABLE clientes (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE,
    ativo BOOLEAN DEFAULT TRUE
);

-- Criar tabela a partir de outra (copia estrutura + dados)
CREATE TABLE clientes_backup AS
SELECT * FROM clientes;

-- Criar tabela com estrutura mas sem dados
CREATE TABLE clientes_novo (LIKE clientes INCLUDING ALL);

-- Modificar tabela (adicionar coluna)
ALTER TABLE clientes ADD COLUMN telefone VARCHAR(20);

-- Modificar tabela (remover coluna)
ALTER TABLE clientes DROP COLUMN telefone;

-- Renomear tabela
ALTER TABLE clientes RENAME TO customers;

-- Deletar tabela
DROP TABLE customers;

-- Deletar se existir (não dá erro se não existir)
DROP TABLE IF EXISTS customers;

-- Deletar com cascade (deleta objetos dependentes)
DROP TABLE clientes CASCADE;
```

### Tipos de Tabelas no PostgreSQL

#### Tabelas Permanentes (Padrão)

```sql
CREATE TABLE produtos (...);
```

#### Tabelas Temporárias

```sql
-- Existe apenas durante a sessão
CREATE TEMP TABLE temp_calculo (
    id INTEGER,
    valor DECIMAL
);
-- Deletada automaticamente ao desconectar
```

#### Tabelas Não Logadas (UNLOGGED)

```sql
-- Mais rápida, mas não recupera após crash
CREATE UNLOGGED TABLE logs (
    id SERIAL,
    mensagem TEXT,
    momento TIMESTAMP
);
```

---

## 📋 4. Columns (Colunas)

Uma **column** define um atributo que todos os registros da tabela terão.

### Definindo Colunas

```sql
CREATE TABLE funcionarios (
    -- Sintaxe: nome_coluna TIPO [CONSTRAINT] [DEFAULT]

    id SERIAL PRIMARY KEY,                    -- Auto-incremento, chave primária
    nome VARCHAR(100) NOT NULL,               -- Texto até 100 chars, obrigatório
    email VARCHAR(100) UNIQUE,                -- Texto único
    salario DECIMAL(10, 2) CHECK (salario > 0), -- Numérico, positivo
    ativo BOOLEAN DEFAULT TRUE,               -- Booleano, padrão TRUE
    data_admissao DATE NOT NULL,              -- Data obrigatória
    criado_em TIMESTAMP DEFAULT NOW()         -- Timestamp auto
);
```

### Modificando Colunas

```sql
-- Adicionar coluna
ALTER TABLE funcionarios ADD COLUMN departamento VARCHAR(50);

-- Remover coluna
ALTER TABLE funcionarios DROP COLUMN departamento;

-- Renomear coluna
ALTER TABLE funcionarios RENAME COLUMN nome TO nome_completo;

-- Mudar tipo de dados
ALTER TABLE funcionarios ALTER COLUMN salario TYPE NUMERIC(12, 2);

-- Adicionar NOT NULL
ALTER TABLE funcionarios ALTER COLUMN email SET NOT NULL;

-- Remover NOT NULL
ALTER TABLE funcionarios ALTER COLUMN email DROP NOT NULL;

-- Definir valor padrão
ALTER TABLE funcionarios ALTER COLUMN ativo SET DEFAULT TRUE;

-- Remover valor padrão
ALTER TABLE funcionarios ALTER COLUMN ativo DROP DEFAULT;

-- Adicionar constraint
ALTER TABLE funcionarios ADD CONSTRAINT salario_positivo
CHECK (salario > 0);
```

---

## 📝 5. Rows (Linhas/Registros)

Uma **row** é um registro individual na tabela, contendo valores para cada coluna.

### Características

- Cada linha é **única** (mesmo que valores sejam iguais, internamente tem identificador)
- Linhas não têm ordem garantida (a menos que use `ORDER BY`)
- Uma linha contém um valor para **cada coluna** (pode ser NULL se permitido)

### Manipulando Linhas

```sql
-- INSERIR linhas
INSERT INTO produtos (nome, preco, estoque)
VALUES ('Notebook', 3000.00, 10);

-- Inserir múltiplas linhas
INSERT INTO produtos (nome, preco, estoque)
VALUES
    ('Mouse', 50.00, 100),
    ('Teclado', 150.00, 50),
    ('Monitor', 800.00, 25);

-- Inserir e retornar a linha criada
INSERT INTO produtos (nome, preco, estoque)
VALUES ('Webcam', 200.00, 30)
RETURNING *;

-- ATUALIZAR linhas
UPDATE produtos
SET preco = 45.00
WHERE nome = 'Mouse';

-- Atualizar múltiplas colunas
UPDATE produtos
SET preco = preco * 1.1,  -- Aumenta 10%
    estoque = estoque - 1
WHERE id = 1;

-- DELETAR linhas
DELETE FROM produtos
WHERE estoque = 0;

-- Deletar todas as linhas (cuidado!)
DELETE FROM produtos;

-- Truncar (mais rápido que DELETE, reseta sequências)
TRUNCATE TABLE produtos;
```

---

## 🎯 6. Queries no PostgreSQL

**Queries** são a forma principal de interagir com o banco de dados.

### Estrutura Básica do SELECT

```sql
SELECT [colunas]
FROM [tabela]
WHERE [condições]
GROUP BY [agrupamento]
HAVING [condições de grupo]
ORDER BY [ordenação]
LIMIT [quantidade]
OFFSET [pular];
```

### SELECT Básico

```sql
-- Selecionar todas as colunas
SELECT * FROM produtos;

-- Selecionar colunas específicas
SELECT nome, preco FROM produtos;

-- Renomear colunas no resultado (alias)
SELECT
    nome AS produto,
    preco AS valor,
    estoque AS quantidade
FROM produtos;

-- Cálculos nas colunas
SELECT
    nome,
    preco,
    preco * 1.1 AS preco_com_imposto,
    preco * estoque AS valor_total_estoque
FROM produtos;

-- Valores distintos
SELECT DISTINCT categoria FROM produtos;
```

### WHERE - Filtrando Dados

```sql
-- Comparações
SELECT * FROM produtos WHERE preco > 100;
SELECT * FROM produtos WHERE estoque <= 10;
SELECT * FROM produtos WHERE categoria = 'Eletrônicos';
SELECT * FROM produtos WHERE categoria != 'Livros';

-- Operadores lógicos
SELECT * FROM produtos
WHERE preco > 50 AND estoque > 0;

SELECT * FROM produtos
WHERE categoria = 'Livros' OR categoria = 'Eletrônicos';

SELECT * FROM produtos
WHERE NOT (estoque = 0);

-- BETWEEN
SELECT * FROM produtos
WHERE preco BETWEEN 100 AND 500;

-- IN
SELECT * FROM produtos
WHERE categoria IN ('Livros', 'Eletrônicos', 'Roupas');

-- LIKE (padrões de texto)
SELECT * FROM produtos WHERE nome LIKE 'Note%';     -- Começa com "Note"
SELECT * FROM produtos WHERE nome LIKE '%book%';    -- Contém "book"
SELECT * FROM produtos WHERE nome LIKE '_ouse';     -- M_ouse (um caractere)

-- ILIKE (case insensitive)
SELECT * FROM produtos WHERE nome ILIKE '%MOUSE%';  -- Ignora maiúsculas

-- IS NULL / IS NOT NULL
SELECT * FROM produtos WHERE descricao IS NULL;
SELECT * FROM produtos WHERE descricao IS NOT NULL;
```

### ORDER BY - Ordenação

```sql
-- Ordem crescente (padrão)
SELECT * FROM produtos ORDER BY preco;
SELECT * FROM produtos ORDER BY preco ASC;

-- Ordem decrescente
SELECT * FROM produtos ORDER BY preco DESC;

-- Múltiplas colunas
SELECT * FROM produtos
ORDER BY categoria ASC, preco DESC;

-- Ordenar por expressão
SELECT nome, preco, estoque, (preco * estoque) AS total
FROM produtos
ORDER BY total DESC;

-- NULLS FIRST / NULLS LAST
SELECT * FROM produtos ORDER BY descricao NULLS LAST;
```

### LIMIT e OFFSET - Paginação

```sql
-- Primeiros 10 registros
SELECT * FROM produtos LIMIT 10;

-- Pular 20, pegar próximos 10 (página 3)
SELECT * FROM produtos LIMIT 10 OFFSET 20;

-- Sintaxe alternativa (LIMIT offset, count)
SELECT * FROM produtos OFFSET 20 LIMIT 10;

-- Top 5 mais caros
SELECT * FROM produtos ORDER BY preco DESC LIMIT 5;
```

### Funções de Agregação

```sql
-- COUNT - contar registros
SELECT COUNT(*) FROM produtos;
SELECT COUNT(*) FROM produtos WHERE estoque > 0;
SELECT COUNT(DISTINCT categoria) FROM produtos;

-- SUM - somar valores
SELECT SUM(estoque) FROM produtos;
SELECT SUM(preco * estoque) AS valor_total_inventario FROM produtos;

-- AVG - média
SELECT AVG(preco) FROM produtos;
SELECT AVG(preco) FROM produtos WHERE categoria = 'Livros';

-- MIN e MAX
SELECT MIN(preco) FROM produtos;
SELECT MAX(preco) FROM produtos;
SELECT MIN(preco) AS mais_barato, MAX(preco) AS mais_caro FROM produtos;
```

### GROUP BY - Agrupamento

```sql
-- Contar produtos por categoria
SELECT categoria, COUNT(*) AS total
FROM produtos
GROUP BY categoria;

-- Valor médio por categoria
SELECT
    categoria,
    COUNT(*) AS quantidade,
    AVG(preco) AS preco_medio,
    SUM(estoque) AS estoque_total
FROM produtos
GROUP BY categoria;

-- Múltiplas colunas no GROUP BY
SELECT
    categoria,
    ativo,
    COUNT(*) AS total
FROM produtos
GROUP BY categoria, ativo;
```

### HAVING - Filtro Após Agregação

```sql
-- Categorias com mais de 10 produtos
SELECT categoria, COUNT(*) AS total
FROM produtos
GROUP BY categoria
HAVING COUNT(*) > 10;

-- Categorias com preço médio acima de 100
SELECT categoria, AVG(preco) AS preco_medio
FROM produtos
GROUP BY categoria
HAVING AVG(preco) > 100;

-- WHERE (antes da agregação) vs HAVING (depois)
SELECT categoria, AVG(preco) AS preco_medio
FROM produtos
WHERE estoque > 0           -- Filtra linhas ANTES de agrupar
GROUP BY categoria
HAVING AVG(preco) > 50;     -- Filtra grupos DEPOIS de agregar
```

### JOINs - Combinando Tabelas

```sql
-- Exemplo: duas tabelas relacionadas
CREATE TABLE categorias (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(50)
);

CREATE TABLE produtos (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(100),
    categoria_id INTEGER REFERENCES categorias(id)
);

-- INNER JOIN (apenas registros com correspondência)
SELECT p.nome, c.nome AS categoria
FROM produtos p
INNER JOIN categorias c ON p.categoria_id = c.id;

-- LEFT JOIN (todos de produtos, mesmo sem categoria)
SELECT p.nome, c.nome AS categoria
FROM produtos p
LEFT JOIN categorias c ON p.categoria_id = c.id;

-- RIGHT JOIN (todas categorias, mesmo sem produtos)
SELECT p.nome, c.nome AS categoria
FROM produtos p
RIGHT JOIN categorias c ON p.categoria_id = c.id;

-- FULL OUTER JOIN (todos de ambos)
SELECT p.nome, c.nome AS categoria
FROM produtos p
FULL OUTER JOIN categorias c ON p.categoria_id = c.id;
```

### Subconsultas (Subqueries)

```sql
-- Produtos com preço acima da média
SELECT nome, preco
FROM produtos
WHERE preco > (SELECT AVG(preco) FROM produtos);

-- IN com subconsulta
SELECT nome
FROM produtos
WHERE categoria_id IN (
    SELECT id FROM categorias WHERE nome IN ('Livros', 'Eletrônicos')
);

-- EXISTS
SELECT c.nome
FROM categorias c
WHERE EXISTS (
    SELECT 1 FROM produtos p WHERE p.categoria_id = c.id
);
```

---

## 📦 7. Data Types (Tipos de Dados)

PostgreSQL oferece uma vasta gama de tipos de dados.

### Tipos Numéricos

```sql
-- Inteiros
SMALLINT        -- -32,768 a 32,767 (2 bytes)
INTEGER         -- -2 bilhões a 2 bilhões (4 bytes)
BIGINT          -- -9 quintilhões a 9 quintilhões (8 bytes)

-- Auto-incremento
SERIAL          -- INTEGER auto-incremento
BIGSERIAL       -- BIGINT auto-incremento

-- Decimais exatos
DECIMAL(10, 2)  -- 10 dígitos, 2 após vírgula
NUMERIC(10, 2)  -- Sinônimo de DECIMAL

-- Ponto flutuante
REAL            -- 6 dígitos de precisão (4 bytes)
DOUBLE PRECISION -- 15 dígitos de precisão (8 bytes)

-- Exemplos
CREATE TABLE numeros (
    id SERIAL,
    idade SMALLINT,
    populacao INTEGER,
    PIB BIGINT,
    preco DECIMAL(10, 2),
    taxa DOUBLE PRECISION
);
```

### Tipos de Caractere

```sql
-- Texto
CHAR(n)         -- Tamanho fixo n, preenche com espaços
VARCHAR(n)      -- Tamanho variável até n
TEXT            -- Tamanho ilimitado

-- Exemplos
CREATE TABLE textos (
    sigla CHAR(2),          -- 'SP', 'RJ' (sempre 2 chars)
    nome VARCHAR(100),      -- Até 100 caracteres
    descricao TEXT          -- Sem limite
);

-- Recomendação: Use TEXT ou VARCHAR(n), evite CHAR
```

### Tipos de Data e Hora

```sql
DATE            -- Data (2024-01-15)
TIME            -- Hora (14:30:00)
TIME WITH TIME ZONE  -- Hora com timezone
TIMESTAMP       -- Data e hora (2024-01-15 14:30:00)
TIMESTAMP WITH TIME ZONE  -- Com timezone
INTERVAL        -- Intervalo de tempo (3 days, 2 hours)

-- Exemplos
CREATE TABLE eventos (
    id SERIAL PRIMARY KEY,
    titulo VARCHAR(100),
    data DATE,
    hora_inicio TIME,
    criado_em TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    duracao INTERVAL
);

-- Inserir
INSERT INTO eventos (titulo, data, hora_inicio, duracao)
VALUES ('Reunião', '2024-12-01', '14:00:00', '2 hours');

-- Operações
SELECT CURRENT_DATE;                    -- Data hoje
SELECT CURRENT_TIME;                    -- Hora agora
SELECT CURRENT_TIMESTAMP;               -- Data e hora agora
SELECT NOW();                           -- Igual a CURRENT_TIMESTAMP

SELECT data + INTERVAL '7 days' FROM eventos;  -- Adicionar 7 dias
SELECT AGE('2024-01-01', '2000-01-01');        -- Calcular idade/diferença
```

### Boolean

```sql
BOOLEAN  -- TRUE, FALSE, NULL

CREATE TABLE produtos (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(100),
    ativo BOOLEAN DEFAULT TRUE,
    promocao BOOLEAN
);

-- Inserir
INSERT INTO produtos (nome, ativo, promocao)
VALUES
    ('Produto 1', TRUE, FALSE),
    ('Produto 2', 't', 'f'),        -- Aceita 't'/'f'
    ('Produto 3', 'yes', 'no'),     -- Aceita 'yes'/'no'
    ('Produto 4', 1, 0);            -- Aceita 1/0

-- Consultar
SELECT * FROM produtos WHERE ativo;              -- TRUE
SELECT * FROM produtos WHERE NOT ativo;          -- FALSE
SELECT * FROM produtos WHERE ativo IS NULL;      -- NULL
```

### Enum (Enumeração)

```sql
-- Criar tipo ENUM
CREATE TYPE status_pedido AS ENUM ('pendente', 'processando', 'enviado', 'entregue', 'cancelado');

-- Usar em tabela
CREATE TABLE pedidos (
    id SERIAL PRIMARY KEY,
    status status_pedido DEFAULT 'pendente'
);

-- Inserir
INSERT INTO pedidos (status) VALUES ('processando');

-- Erro se usar valor não definido
INSERT INTO pedidos (status) VALUES ('desconhecido');  -- ERRO!

-- Ordenação segue ordem definida no ENUM
SELECT * FROM pedidos ORDER BY status;
```

### Arrays

```sql
-- Definir coluna array
CREATE TABLE artigos (
    id SERIAL PRIMARY KEY,
    titulo VARCHAR(200),
    tags TEXT[]          -- Array de texto
);

-- Inserir
INSERT INTO artigos (titulo, tags)
VALUES
    ('PostgreSQL', ARRAY['banco', 'sql', 'database']),
    ('Python', '{programação, python, dev}');  -- Sintaxe alternativa

-- Consultar
SELECT * FROM artigos WHERE 'sql' = ANY(tags);      -- Contém 'sql'
SELECT * FROM artigos WHERE tags @> ARRAY['sql'];   -- Operador de array
SELECT * FROM artigos WHERE tags && ARRAY['python', 'java'];  -- Overlap

-- Acessar elementos (índice começa em 1!)
SELECT titulo, tags[1], tags[2] FROM artigos;
```

### JSON e JSONB

```sql
JSON   -- Armazena texto JSON (mais lento, mantém formato original)
JSONB  -- Armazena binário (mais rápido, indexável, recomendado)

CREATE TABLE produtos (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(100),
    especificacoes JSONB
);

-- Inserir
INSERT INTO produtos (nome, especificacoes)
VALUES (
    'Notebook',
    '{"marca": "Dell", "ram": "16GB", "ssd": "512GB", "tela": 15.6}'::jsonb
);

-- Consultar
SELECT * FROM produtos WHERE especificacoes->>'marca' = 'Dell';
SELECT nome, especificacoes->'ram' AS memoria FROM produtos;
SELECT * FROM produtos WHERE especificacoes @> '{"marca": "Dell"}';

-- Atualizar JSON
UPDATE produtos
SET especificacoes = especificacoes || '{"cor": "prata"}'::jsonb
WHERE id = 1;
```

### Tipos Geométricos

```sql
POINT       -- Ponto (x, y)
LINE        -- Linha infinita
LSEG        -- Segmento de linha
BOX         -- Retângulo
PATH        -- Caminho (aberto ou fechado)
POLYGON     -- Polígono
CIRCLE      -- Círculo

CREATE TABLE locais (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(100),
    localizacao POINT,
    area CIRCLE
);

INSERT INTO locais (nome, localizacao, area)
VALUES ('Loja Centro', POINT(10.5, 20.3), CIRCLE(POINT(10.5, 20.3), 5.0));
```

### UUID (Identificador Único Universal)

```sql
-- Habilitar extensão (uma vez)
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE usuarios (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    nome VARCHAR(100),
    email VARCHAR(100)
);

INSERT INTO usuarios (nome, email)
VALUES ('João', 'joao@email.com')
RETURNING id;  -- Retorna UUID gerado: ex: 550e8400-e29b-41d4-a716-446655440000
```

---

## 📊 Resumo da Hierarquia

```
┌─────────────────────────────────────────────────────────┐
│  SERVIDOR                                               │
│  └─ DATABASE (loja)              ← Isolamento completo  │
│     ├─ SCHEMA (public)           ← Namespace/organização│
│     │  ├─ TABLE (clientes)       ← Estrutura de dados   │
│     │  │  ├─ COLUMN (id)         ← Atributo/tipo        │
│     │  │  ├─ COLUMN (nome)                              │
│     │  │  └─ ROW (1, 'João')     ← Dados reais          │
│     │  └─ TABLE (pedidos)                               │
│     └─ SCHEMA (vendas)                                  │
└─────────────────────────────────────────────────────────┘
```

---

## 🎓 Conclusão

Nesta aula você aprendeu:

1. **ORDBMS**: PostgreSQL combina modelo relacional com recursos orientados a objetos
2. **Hierarquia**: Database → Schema → Table → Column → Row
3. **Queries**: Como buscar e manipular dados com SELECT, WHERE, JOIN, etc.
4. **Tipos de Dados**: Numéricos, texto, data/hora, boolean, JSON, arrays e mais

O modelo de objetos do PostgreSQL oferece flexibilidade sem perder o rigor do modelo relacional!

---

## 🔑 Conceitos para Memorizar

- **ORDBMS**: Object-Relational DBMS (híbrido)
- **Database**: Coleção isolada de schemas
- **Schema**: Namespace que organiza objetos
- **Table**: Coleção de linhas com estrutura definida
- **Column**: Define tipo e constraints de um atributo
- **Row**: Registro individual com dados reais
- **Query**: Comando para buscar/manipular dados
- **JSONB**: Tipo para dados semi-estruturados (recomendado sobre JSON)
- **SERIAL**: Auto-incremento para chaves primárias
