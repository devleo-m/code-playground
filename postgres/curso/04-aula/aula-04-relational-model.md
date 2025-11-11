# **Aula 4: Relational Model - Aprofundamento**

## 🎯 Objetivo da Aula

Aprofundar os conceitos do modelo relacional com foco em domínios customizados, constraints avançadas, tratamento de null values e implementação prática no PostgreSQL.

---

## 📚 Revisão: O Modelo Relacional

### A Revolução de E.F. Codd (1970)

O modelo relacional transformou bancos de dados ao introduzir:

- Organização de dados em **tabelas** (relações)
- Separação entre **representação lógica** e **armazenamento físico**
- Base matemática sólida (teoria de conjuntos e álgebra relacional)

**Fundamentos que já vimos:**

- Relação (tabela)
- Tupla (linha/registro)
- Atributo (coluna/campo)
- Domínio (valores permitidos)

Agora vamos **aprofundar** esses conceitos e ver como PostgreSQL os implementa!

---

## 🎨 1. Domains (Domínios Customizados)

### O que são Domains?

**Domains** no PostgreSQL são **tipos de dados customizados** que você pode criar para reutilizar constraints e validações.

### Por que usar Domains?

Imagine que você tem 10 tabelas com campos de email:

```sql
-- SEM domains (repetindo constraint em cada tabela)
CREATE TABLE usuarios (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) CHECK (email ~* '^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$')
);

CREATE TABLE clientes (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) CHECK (email ~* '^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$')
);

-- ... e assim por diante (repetindo 10 vezes!) ❌
```

**Problema:** Código duplicado, difícil de manter. Se quiser mudar a validação, precisa alterar 10 lugares!

**Solução com Domains:**

```sql
-- COM domain (define uma vez, usa em qualquer lugar)
CREATE DOMAIN email_type AS VARCHAR(255)
    CHECK (VALUE ~* '^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$');

-- Usar em todas as tabelas
CREATE TABLE usuarios (
    id SERIAL PRIMARY KEY,
    email email_type  -- ✅ Usa o domain!
);

CREATE TABLE clientes (
    id SERIAL PRIMARY KEY,
    email email_type  -- ✅ Usa o domain!
);
```

**Benefícios:**

- ✅ **DRY** (Don't Repeat Yourself): Define uma vez, usa em qualquer lugar
- ✅ **Manutenção fácil**: Alterar domain afeta todas as tabelas
- ✅ **Documentação**: Nome do domain documenta intenção (`email_type`, `cep_brasileiro`)
- ✅ **Consistência**: Mesma validação em todo sistema

---

### Criando Domains

#### Sintaxe Básica

```sql
CREATE DOMAIN nome_domain AS tipo_base
    [DEFAULT expressao]
    [CONSTRAINT nome_constraint CHECK (condicao)]
    [NOT NULL];
```

#### Exemplos Práticos

##### 1. Domain para Email

```sql
CREATE DOMAIN email_type AS VARCHAR(255)
    CHECK (VALUE ~* '^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$');

-- Usar
CREATE TABLE contatos (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(100),
    email email_type
);

-- Teste
INSERT INTO contatos (nome, email) VALUES ('João', 'joao@example.com');  -- ✅
INSERT INTO contatos (nome, email) VALUES ('Maria', 'email-invalido');    -- ❌ ERRO!
```

##### 2. Domain para CPF Brasileiro

```sql
CREATE DOMAIN cpf_brasileiro AS CHAR(14)
    CHECK (VALUE ~ '^\d{3}\.\d{3}\.\d{3}-\d{2}$');

-- Usar
CREATE TABLE pessoas (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(100),
    cpf cpf_brasileiro
);

-- Teste
INSERT INTO pessoas (nome, cpf) VALUES ('João', '123.456.789-00');   -- ✅
INSERT INTO pessoas (nome, cpf) VALUES ('Maria', '12345678900');      -- ❌ Formato errado!
```

##### 3. Domain para CEP Brasileiro

```sql
CREATE DOMAIN cep_brasileiro AS CHAR(9)
    CHECK (VALUE ~ '^\d{5}-\d{3}$');

CREATE TABLE enderecos (
    id SERIAL PRIMARY KEY,
    logradouro VARCHAR(200),
    cep cep_brasileiro
);
```

##### 4. Domain para Valores Positivos

```sql
CREATE DOMAIN valor_positivo AS DECIMAL(10, 2)
    CHECK (VALUE > 0);

CREATE TABLE produtos (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(100),
    preco valor_positivo,      -- Sempre positivo!
    desconto valor_positivo    -- Desconto também sempre positivo!
);
```

##### 5. Domain para Notas (0 a 10)

```sql
CREATE DOMAIN nota_escolar AS DECIMAL(4, 2)
    CHECK (VALUE >= 0 AND VALUE <= 10);

CREATE TABLE avaliacoes (
    id SERIAL PRIMARY KEY,
    aluno_id INTEGER,
    disciplina VARCHAR(50),
    nota nota_escolar          -- Entre 0 e 10
);
```

##### 6. Domain com Valor Padrão

```sql
CREATE DOMAIN status_ativo AS BOOLEAN
    DEFAULT TRUE
    NOT NULL;

CREATE TABLE usuarios (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(100),
    ativo status_ativo         -- Padrão TRUE, nunca NULL
);
```

##### 7. Domain para Ano Válido

```sql
CREATE DOMAIN ano_valido AS INTEGER
    CHECK (VALUE >= 1900 AND VALUE <= 2100);

CREATE TABLE filmes (
    id SERIAL PRIMARY KEY,
    titulo VARCHAR(200),
    ano_lancamento ano_valido
);
```

---

### Modificando Domains

```sql
-- Adicionar constraint
ALTER DOMAIN email_type ADD CONSTRAINT email_nao_vazio CHECK (length(VALUE) > 0);

-- Remover constraint
ALTER DOMAIN email_type DROP CONSTRAINT email_nao_vazio;

-- Definir valor padrão
ALTER DOMAIN email_type SET DEFAULT 'sem-email@example.com';

-- Remover valor padrão
ALTER DOMAIN email_type DROP DEFAULT;

-- Adicionar NOT NULL
ALTER DOMAIN email_type SET NOT NULL;

-- Remover NOT NULL
ALTER DOMAIN email_type DROP NOT NULL;

-- Renomear domain
ALTER DOMAIN email_type RENAME TO tipo_email;

-- Renomear constraint
ALTER DOMAIN email_type RENAME CONSTRAINT email_check TO validacao_email;
```

---

### Deletando Domains

```sql
-- Deletar domain (erro se estiver em uso)
DROP DOMAIN email_type;

-- Deletar se existir (não dá erro se não existir)
DROP DOMAIN IF EXISTS email_type;

-- Deletar e atualizar tabelas que usam (remove domain, mantém tipo base)
DROP DOMAIN email_type CASCADE;
```

---

### Consultando Domains

```sql
-- Listar todos os domains
SELECT domain_name, data_type
FROM information_schema.domains
WHERE domain_schema = 'public';

-- Ver detalhes de um domain
\dD+ email_type

-- Ver constraints de um domain
SELECT constraint_name, check_clause
FROM information_schema.domain_constraints
WHERE domain_name = 'email_type';
```

---

### Domains vs CHECK Constraints

| Aspecto           | Domain                       | CHECK Constraint               |
| ----------------- | ---------------------------- | ------------------------------ |
| **Reutilização**  | ✅ Sim, em múltiplas tabelas | ❌ Não, específico da tabela   |
| **Manutenção**    | ✅ Alterar em um lugar       | ❌ Alterar em cada tabela      |
| **Documentação**  | ✅ Nome documenta intenção   | ⚠️ Constraint pode ser verbosa |
| **Flexibilidade** | ⚠️ Menos flexível            | ✅ Mais flexível por tabela    |

**Use Domains quando:**

- ✅ Mesma validação em múltiplas tabelas
- ✅ Quer documentar tipo de negócio (email, cpf, cep)
- ✅ Quer manter consistência no sistema

**Use CHECK diretamente quando:**

- ✅ Validação específica de uma tabela
- ✅ Validação envolve múltiplas colunas
- ✅ Regra muito particular

---

## 🏷️ 2. Attributes (Atributos)

### O que são Atributos?

**Atributos** são as **colunas** de uma relação (tabela), representando propriedades ou características da entidade.

### Características de Atributos

```sql
CREATE TABLE funcionarios (
    -- Cada linha abaixo define um ATRIBUTO

    id SERIAL PRIMARY KEY,              -- Atributo: id, Domínio: INTEGER (auto)
    nome VARCHAR(100) NOT NULL,         -- Atributo: nome, Domínio: VARCHAR(100)
    email email_type UNIQUE,            -- Atributo: email, Domínio: email_type (custom!)
    salario DECIMAL(10, 2) CHECK (salario > 0),  -- Atributo: salario, Domínio: DECIMAL(10,2)
    data_admissao DATE DEFAULT CURRENT_DATE,     -- Atributo: data_admissao, Domínio: DATE
    ativo BOOLEAN DEFAULT TRUE          -- Atributo: ativo, Domínio: BOOLEAN
);
```

**Cada atributo tem:**

1. **Nome**: Identificador do atributo (ex: `nome`, `email`, `salario`)
2. **Domínio**: Tipo de dados + restrições (ex: `VARCHAR(100)`, `email_type`)
3. **Constraints**: Regras adicionais (ex: `NOT NULL`, `UNIQUE`, `CHECK`)
4. **Valor Padrão**: Opcional (ex: `DEFAULT CURRENT_DATE`)

---

### Papel dos Atributos

#### 1. Definir Estrutura

```sql
-- Esquema define que atributos uma entidade tem
CREATE TABLE livros (
    isbn VARCHAR(17),      -- Atributo que identifica livro
    titulo VARCHAR(200),   -- Atributo descritivo
    autor VARCHAR(100),    -- Atributo descritivo
    ano INTEGER,           -- Atributo numérico
    preco DECIMAL(10, 2)   -- Atributo monetário
);
```

#### 2. Armazenar Dados

```sql
-- Cada tupla atribui valores aos atributos
INSERT INTO livros (isbn, titulo, autor, ano, preco)
VALUES ('978-3-16-148410-0', '1984', 'George Orwell', 1949, 35.90);
```

#### 3. Manter Integridade

```sql
CREATE TABLE produtos (
    id SERIAL PRIMARY KEY,
    codigo VARCHAR(20) UNIQUE NOT NULL,     -- Atributo com UNIQUE
    nome VARCHAR(100) NOT NULL,             -- Atributo obrigatório
    preco DECIMAL(10, 2) CHECK (preco > 0), -- Atributo com validação
    estoque INTEGER CHECK (estoque >= 0)    -- Atributo não-negativo
);
```

#### 4. Habilitar Operações

```sql
-- Atributos permitem projeção (selecionar colunas)
SELECT nome, preco FROM produtos;

-- Atributos permitem seleção (filtrar)
SELECT * FROM produtos WHERE preco > 100;

-- Atributos permitem ordenação
SELECT * FROM produtos ORDER BY nome;

-- Atributos permitem agregação
SELECT AVG(preco), MAX(preco) FROM produtos;
```

---

### Atributos Atômicos vs Compostos

#### Atômicos (Indivisíveis) ✅

```sql
CREATE TABLE clientes (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(100),
    email VARCHAR(100),
    telefone VARCHAR(20)
);
```

#### Compostos (Divisíveis) - Não Normalizado ❌

```sql
-- Evite!
CREATE TABLE clientes_mal_modelados (
    id SERIAL PRIMARY KEY,
    nome_completo VARCHAR(200),  -- João Silva Santos
    endereco_completo TEXT       -- Rua A, 123, São Paulo, SP, 01234-567
);

-- Problema: Difícil buscar por sobrenome ou cidade
SELECT * FROM clientes_mal_modelados WHERE nome_completo LIKE '% Silva%';  -- Impreciso!
```

#### Bem Modelizado ✅

```sql
CREATE TABLE clientes_bem_modelados (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(50),
    sobrenome VARCHAR(50),
    logradouro VARCHAR(100),
    numero VARCHAR(10),
    cidade VARCHAR(50),
    estado CHAR(2),
    cep CHAR(9)
);

-- Agora posso buscar facilmente
SELECT * FROM clientes_bem_modelados WHERE sobrenome = 'Silva';
SELECT * FROM clientes_bem_modelados WHERE cidade = 'São Paulo';
```

---

## 📄 3. Tuples (Tuplas)

### O que são Tuplas?

**Tupla** é um **conjunto ordenado de valores** que corresponde aos atributos de uma relação. Cada tupla representa uma entidade específica.

```sql
CREATE TABLE usuarios (
    id INTEGER,
    nome VARCHAR(100),
    email VARCHAR(100)
);

-- Inserir tuplas
INSERT INTO usuarios VALUES (1, 'João', 'joao@example.com');     -- Tupla 1
INSERT INTO usuarios VALUES (2, 'Maria', 'maria@example.com');   -- Tupla 2
INSERT INTO usuarios VALUES (3, 'Pedro', 'pedro@example.com');   -- Tupla 3

-- Cada linha é uma tupla:
-- Tupla 1: (1, 'João', 'joao@example.com')
-- Tupla 2: (2, 'Maria', 'maria@example.com')
-- Tupla 3: (3, 'Pedro', 'pedro@example.com')
```

---

### Características de Tuplas

#### 1. **Ordenação de Valores**

Valores em uma tupla correspondem à ordem dos atributos:

```sql
-- Tabela
CREATE TABLE produtos (id, nome, preco);

-- Tupla: (1, 'Mouse', 50.00)
--         ↑    ↑       ↑
--         id   nome   preco
```

#### 2. **Unicidade**

Mesmo que valores sejam iguais, tuplas são distintas (internamente, PostgreSQL mantém identificadores):

```sql
INSERT INTO produtos VALUES (1, 'Mouse', 50.00);
INSERT INTO produtos VALUES (2, 'Mouse', 50.00);  -- Tupla diferente (ID diferente)
```

#### 3. **Valores podem ser NULL**

```sql
INSERT INTO produtos (id, nome) VALUES (3, 'Webcam');  -- preco = NULL
-- Tupla: (3, 'Webcam', NULL)
```

---

### Operações em Tuplas

#### Inserção

```sql
-- Inserir uma tupla
INSERT INTO usuarios (id, nome, email)
VALUES (4, 'Ana', 'ana@example.com');

-- Inserir múltiplas tuplas
INSERT INTO usuarios (id, nome, email)
VALUES
    (5, 'Carlos', 'carlos@example.com'),
    (6, 'Beatriz', 'beatriz@example.com');
```

#### Atualização

```sql
-- Atualizar tupla específica
UPDATE usuarios
SET email = 'joao.novo@example.com'
WHERE id = 1;

-- Tupla (1, 'João', 'joao@example.com')
-- vira (1, 'João', 'joao.novo@example.com')
```

#### Deleção

```sql
-- Deletar tupla específica
DELETE FROM usuarios WHERE id = 2;

-- Tupla (2, 'Maria', 'maria@example.com') é removida
```

#### Seleção

```sql
-- Buscar tuplas que atendem condição
SELECT * FROM usuarios WHERE nome LIKE 'J%';

-- Retorna tuplas cujo nome começa com 'J'
```

---

### Tipo Composto (Composite Type)

PostgreSQL permite criar tipos compostos que representam tuplas:

```sql
-- Criar tipo composto
CREATE TYPE endereco_type AS (
    logradouro VARCHAR(100),
    numero VARCHAR(10),
    cidade VARCHAR(50),
    estado CHAR(2),
    cep CHAR(9)
);

-- Usar em tabela
CREATE TABLE clientes (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(100),
    endereco endereco_type  -- Atributo do tipo composto!
);

-- Inserir
INSERT INTO clientes (nome, endereco)
VALUES (
    'João',
    ROW('Rua A', '123', 'São Paulo', 'SP', '01234-567')
);

-- Consultar atributos da tupla
SELECT
    nome,
    (endereco).cidade,      -- Acessar atributo da tupla
    (endereco).estado
FROM clientes;
```

---

## 🗂️ 4. Relations (Relações)

### O que são Relações?

Uma **relação** é uma **tabela** composta por:

- **Schema**: Estrutura (nome da relação + nomes e tipos dos atributos)
- **Instance**: Conjunto de tuplas (dados reais)

```sql
-- SCHEMA (estrutura)
CREATE TABLE produtos (
    id INTEGER PRIMARY KEY,
    nome VARCHAR(100),
    preco DECIMAL(10, 2)
);

-- INSTANCE (dados)
INSERT INTO produtos VALUES (1, 'Mouse', 50.00);
INSERT INTO produtos VALUES (2, 'Teclado', 150.00);
```

---

### Propriedades de Relações

#### 1. **Sem Tuplas Duplicadas**

Relações são **conjuntos**, portanto não há tuplas duplicadas (garantido por chave primária):

```sql
CREATE TABLE categorias (
    id INTEGER PRIMARY KEY,
    nome VARCHAR(50) UNIQUE
);

INSERT INTO categorias VALUES (1, 'Eletrônicos');
INSERT INTO categorias VALUES (2, 'Eletrônicos');  -- ❌ ERRO! Nome duplicado (UNIQUE)
```

#### 2. **Ordem das Tuplas Não Importa**

Tuplas não têm ordem fixa (a menos que use `ORDER BY`):

```sql
-- Essas duas queries podem retornar tuplas em qualquer ordem
SELECT * FROM produtos;
SELECT * FROM produtos;

-- Para garantir ordem, use ORDER BY
SELECT * FROM produtos ORDER BY id;
```

#### 3. **Ordem dos Atributos Importa (internamente)**

Atributos têm posição definida:

```sql
-- Inserir valores na ordem dos atributos
INSERT INTO produtos VALUES (1, 'Mouse', 50.00);  -- id, nome, preco

-- Ou especificar ordem diferente
INSERT INTO produtos (nome, preco, id) VALUES ('Teclado', 150.00, 2);
```

---

### Grau e Cardinalidade

#### Grau (Degree)

Número de atributos:

```sql
CREATE TABLE usuarios (
    id INTEGER,
    nome VARCHAR(100),
    email VARCHAR(100)
);
-- Grau = 3 (três atributos)
```

#### Cardinalidade (Cardinality)

Número de tuplas:

```sql
INSERT INTO usuarios VALUES (1, 'João', 'joao@example.com');
INSERT INTO usuarios VALUES (2, 'Maria', 'maria@example.com');
-- Cardinalidade = 2 (duas tuplas)
```

---

### Operações Relacionais

#### Seleção (σ) - Filtrar Tuplas

```sql
SELECT * FROM produtos WHERE preco > 100;
-- σ(preco > 100)(produtos)
```

#### Projeção (π) - Selecionar Atributos

```sql
SELECT nome, preco FROM produtos;
-- π(nome, preco)(produtos)
```

#### União (∪) - Combinar Relações

```sql
SELECT nome FROM clientes_sp
UNION
SELECT nome FROM clientes_rj;
```

#### Interseção (∩)

```sql
SELECT nome FROM clientes_sp
INTERSECT
SELECT nome FROM clientes_premium;
```

#### Diferença (−)

```sql
SELECT nome FROM todos_clientes
EXCEPT
SELECT nome FROM clientes_inativos;
```

#### Produto Cartesiano (×)

```sql
SELECT * FROM tabela1 CROSS JOIN tabela2;
```

#### Junção (⋈)

```sql
SELECT *
FROM pedidos p
INNER JOIN clientes c ON p.cliente_id = c.id;
```

---

## 🔒 5. Constraints (Restrições de Integridade)

Constraints garantem que os dados sigam regras específicas.

### 1. PRIMARY KEY (Chave Primária)

Identifica unicamente cada tupla. **Unique + Not NULL**.

```sql
CREATE TABLE alunos (
    id SERIAL PRIMARY KEY,      -- PK em uma coluna
    nome VARCHAR(100),
    email VARCHAR(100)
);

-- PK composta (múltiplas colunas)
CREATE TABLE matriculas (
    aluno_id INTEGER,
    turma_id INTEGER,
    data_matricula DATE,
    PRIMARY KEY (aluno_id, turma_id)  -- PK composta
);
```

**Características:**

- ✅ Valores únicos
- ✅ Não pode ser NULL
- ✅ Cria índice automaticamente
- ✅ Uma por tabela

---

### 2. FOREIGN KEY (Chave Estrangeira)

Mantém integridade referencial entre tabelas.

```sql
CREATE TABLE categorias (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(50)
);

CREATE TABLE produtos (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(100),
    categoria_id INTEGER,
    FOREIGN KEY (categoria_id) REFERENCES categorias(id)
);

-- Teste
INSERT INTO categorias (id, nome) VALUES (1, 'Eletrônicos');
INSERT INTO produtos (nome, categoria_id) VALUES ('Mouse', 1);     -- ✅
INSERT INTO produtos (nome, categoria_id) VALUES ('Teclado', 999); -- ❌ ERRO! Categoria 999 não existe
```

#### Ações em CASCADE

```sql
CREATE TABLE produtos (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(100),
    categoria_id INTEGER,
    FOREIGN KEY (categoria_id) REFERENCES categorias(id)
        ON DELETE CASCADE          -- Deletar categoria deleta produtos
        ON UPDATE CASCADE          -- Atualizar id da categoria atualiza produtos
);

-- Outras opções:
-- ON DELETE SET NULL          -- Seta NULL nos produtos
-- ON DELETE SET DEFAULT       -- Seta valor padrão
-- ON DELETE RESTRICT (padrão) -- Não permite deletar se houver produtos
```

---

### 3. UNIQUE (Unicidade)

Garante que valores sejam únicos (mas permite NULL).

```sql
CREATE TABLE usuarios (
    id SERIAL PRIMARY KEY,
    email VARCHAR(100) UNIQUE,     -- Email único
    cpf CHAR(14) UNIQUE            -- CPF único
);

-- Teste
INSERT INTO usuarios (email, cpf) VALUES ('joao@example.com', '111.111.111-11');  -- ✅
INSERT INTO usuarios (email, cpf) VALUES ('joao@example.com', '222.222.222-22');  -- ❌ Email duplicado!
INSERT INTO usuarios (email, cpf) VALUES (NULL, '333.333.333-33');                -- ✅ NULL é permitido
INSERT INTO usuarios (email, cpf) VALUES (NULL, '444.444.444-44');                -- ✅ Múltiplos NULL OK
```

#### UNIQUE Composto

```sql
CREATE TABLE assentos_cinema (
    sala INTEGER,
    fileira CHAR(1),
    numero INTEGER,
    UNIQUE (sala, fileira, numero)  -- Combinação única
);
```

---

### 4. CHECK (Validação)

Valida condições personalizadas.

```sql
CREATE TABLE funcionarios (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(100),
    salario DECIMAL(10, 2) CHECK (salario > 0),           -- Salário positivo
    idade INTEGER CHECK (idade >= 18 AND idade <= 100),   -- Idade válida
    email VARCHAR(100) CHECK (email LIKE '%@%')           -- Email tem @
);

-- Teste
INSERT INTO funcionarios (nome, salario, idade, email)
VALUES ('João', 5000.00, 25, 'joao@example.com');  -- ✅

INSERT INTO funcionarios (nome, salario, idade, email)
VALUES ('Maria', -1000.00, 25, 'maria@example.com');  -- ❌ Salário negativo!

INSERT INTO funcionarios (nome, salario, idade, email)
VALUES ('Pedro', 3000.00, 15, 'pedro@example.com');  -- ❌ Menor de idade!
```

#### CHECK com Nome

```sql
CREATE TABLE produtos (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(100),
    preco DECIMAL(10, 2),
    desconto DECIMAL(10, 2),
    CONSTRAINT preco_positivo CHECK (preco > 0),
    CONSTRAINT desconto_valido CHECK (desconto >= 0 AND desconto <= preco)
);
```

#### CHECK entre Múltiplas Colunas

```sql
CREATE TABLE eventos (
    id SERIAL PRIMARY KEY,
    titulo VARCHAR(100),
    data_inicio DATE,
    data_fim DATE,
    CHECK (data_fim >= data_inicio)  -- Fim não pode ser antes do início
);
```

---

### 5. NOT NULL (Valor Obrigatório)

Garante que coluna não pode ser NULL.

```sql
CREATE TABLE clientes (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(100) NOT NULL,      -- Obrigatório
    email VARCHAR(100) NOT NULL,     -- Obrigatório
    telefone VARCHAR(20)             -- Opcional (pode ser NULL)
);

-- Teste
INSERT INTO clientes (nome, email) VALUES ('João', 'joao@example.com');  -- ✅
INSERT INTO clientes (email) VALUES ('maria@example.com');                -- ❌ Nome é obrigatório!
```

---

### 6. EXCLUSION (Exclusão)

Previne conflitos entre tuplas (constraint avançada).

```sql
-- Habilitar extensão (uma vez)
CREATE EXTENSION IF NOT EXISTS btree_gist;

-- Tabela de reservas de salas
CREATE TABLE reservas (
    id SERIAL PRIMARY KEY,
    sala INTEGER,
    periodo TSRANGE,  -- Range de timestamp
    EXCLUDE USING GIST (
        sala WITH =,           -- Mesma sala
        periodo WITH &&        -- Períodos sobrepostos
    )
);

-- Teste
INSERT INTO reservas (sala, periodo)
VALUES (1, '[2024-12-01 14:00, 2024-12-01 16:00)');  -- ✅

INSERT INTO reservas (sala, periodo)
VALUES (1, '[2024-12-01 15:00, 2024-12-01 17:00)');  -- ❌ Conflito! Sobrepõe período
```

---

## ❓ 6. Null Values (Valores Nulos)

### O que é NULL?

**NULL** representa **ausência de valor** ou **valor desconhecido**. Não é:

- ❌ Zero (0)
- ❌ String vazia ('')
- ❌ False
- ❌ Espaço em branco (' ')

NULL é um marcador especial que significa "não há valor aqui".

---

### Comparações com NULL

```sql
-- Criar tabela de teste
CREATE TABLE teste_null (
    id SERIAL PRIMARY KEY,
    valor INTEGER
);

INSERT INTO teste_null (valor) VALUES (10), (NULL), (20);

-- Comparações retornam NULL (não TRUE nem FALSE)
SELECT * FROM teste_null WHERE valor = NULL;      -- ❌ Retorna nada! (NULL = NULL é NULL)
SELECT * FROM teste_null WHERE valor != NULL;     -- ❌ Retorna nada!

-- Forma correta: IS NULL / IS NOT NULL
SELECT * FROM teste_null WHERE valor IS NULL;     -- ✅ Retorna linha com NULL
SELECT * FROM teste_null WHERE valor IS NOT NULL; -- ✅ Retorna linhas com 10 e 20
```

---

### Operações Aritméticas com NULL

Qualquer operação com NULL resulta em NULL:

```sql
SELECT 10 + NULL;        -- NULL
SELECT 10 * NULL;        -- NULL
SELECT NULL / 5;         -- NULL
SELECT 10 > NULL;        -- NULL (não TRUE nem FALSE!)
```

---

### Lógica Booleana com NULL (Três Valores)

PostgreSQL usa lógica de **três valores**: TRUE, FALSE, **NULL** (desconhecido)

```sql
-- AND
SELECT TRUE AND NULL;    -- NULL (pode ser TRUE ou FALSE)
SELECT FALSE AND NULL;   -- FALSE (já sabemos que é FALSE)

-- OR
SELECT TRUE OR NULL;     -- TRUE (já sabemos que é TRUE)
SELECT FALSE OR NULL;    -- NULL (pode ser TRUE ou FALSE)

-- NOT
SELECT NOT NULL;         -- NULL (não sabemos o valor)
```

---

### Funções para Trabalhar com NULL

#### IS NULL / IS NOT NULL

```sql
SELECT * FROM usuarios WHERE email IS NULL;
SELECT * FROM usuarios WHERE email IS NOT NULL;
```

#### COALESCE - Retorna Primeiro Valor Não-Nulo

```sql
SELECT
    nome,
    COALESCE(telefone, email, 'Sem contato') AS contato
FROM clientes;

-- Se telefone não é NULL, usa telefone
-- Se telefone é NULL mas email não é, usa email
-- Se ambos são NULL, usa 'Sem contato'
```

#### NULLIF - Retorna NULL se Valores São Iguais

```sql
SELECT NULLIF(valor1, valor2);

-- Se valor1 = valor2, retorna NULL
-- Senão, retorna valor1

-- Exemplo: Evitar divisão por zero
SELECT
    total / NULLIF(quantidade, 0) AS media
FROM estatisticas;
-- Se quantidade é 0, NULLIF retorna NULL, evitando divisão por zero
```

#### COALESCE vs NULLIF Juntos

```sql
-- Média segura: Se quantidade é 0, retorna 0 em vez de NULL
SELECT
    COALESCE(total / NULLIF(quantidade, 0), 0) AS media_segura
FROM estatisticas;
```

---

### NULL em Agregações

```sql
CREATE TABLE vendas (
    id SERIAL PRIMARY KEY,
    valor DECIMAL(10, 2)
);

INSERT INTO vendas (valor) VALUES (100), (200), (NULL), (300);

-- Agregações ignoram NULL
SELECT COUNT(*) FROM vendas;           -- 4 (conta todas as linhas)
SELECT COUNT(valor) FROM vendas;       -- 3 (ignora NULL)
SELECT SUM(valor) FROM vendas;         -- 600 (ignora NULL)
SELECT AVG(valor) FROM vendas;         -- 200 (média de 100, 200, 300)
```

---

### NULL em ORDER BY

```sql
-- NULL vai para o final (padrão)
SELECT * FROM usuarios ORDER BY email;

-- NULL vai para o início
SELECT * FROM usuarios ORDER BY email NULLS FIRST;

-- NULL vai para o final (explícito)
SELECT * FROM usuarios ORDER BY email NULLS LAST;
```

---

### NULL vs Valor Padrão

```sql
-- Com DEFAULT: Valor é preenchido automaticamente
CREATE TABLE produtos (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(100),
    estoque INTEGER DEFAULT 0  -- Se não informado, usa 0
);

INSERT INTO produtos (nome) VALUES ('Mouse');
SELECT * FROM produtos;  -- estoque = 0 (não NULL)

-- Sem DEFAULT: Valor fica NULL
CREATE TABLE produtos2 (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(100),
    estoque INTEGER  -- Sem DEFAULT
);

INSERT INTO produtos2 (nome) VALUES ('Mouse');
SELECT * FROM produtos2;  -- estoque = NULL
```

---

## 📊 Resumo dos Conceitos

| Conceito        | Definição                         | Exemplo                         |
| --------------- | --------------------------------- | ------------------------------- |
| **Domain**      | Tipo customizado com constraints  | `email_type`, `cpf_brasileiro`  |
| **Attribute**   | Coluna da tabela                  | `nome`, `email`, `preco`        |
| **Tuple**       | Linha/registro                    | `(1, 'João', 'joao@email.com')` |
| **Relation**    | Tabela (schema + tuplas)          | Tabela `usuarios`               |
| **PRIMARY KEY** | Identificador único e obrigatório | `id SERIAL PRIMARY KEY`         |
| **FOREIGN KEY** | Referência a outra tabela         | `REFERENCES categorias(id)`     |
| **UNIQUE**      | Valores únicos (permite NULL)     | `email VARCHAR(100) UNIQUE`     |
| **CHECK**       | Validação customizada             | `CHECK (idade >= 18)`           |
| **NOT NULL**    | Valor obrigatório                 | `nome VARCHAR(100) NOT NULL`    |
| **EXCLUSION**   | Previne conflitos                 | Reservas sem sobreposição       |
| **NULL**        | Ausência de valor                 | Diferente de 0, '', false       |

---

## 🎓 Conclusão

Nesta aula você aprendeu:

1. **Domains**: Tipos customizados reutilizáveis com constraints
2. **Attributes**: Colunas que definem propriedades das entidades
3. **Tuples**: Registros individuais (linhas)
4. **Relations**: Tabelas com schema e dados
5. **Constraints**: PRIMARY KEY, FOREIGN KEY, UNIQUE, CHECK, NOT NULL, EXCLUSION
6. **NULL**: Ausência de valor e como trabalhar com ele

Esses conceitos formam a base sólida do modelo relacional no PostgreSQL!

---

## 🔑 Conceitos para Memorizar

- **Domain**: Tipo customizado = DRY + Manutenibilidade
- **Attribute**: Coluna = Propriedade + Domínio
- **Tuple**: Linha = Conjunto ordenado de valores
- **Relation**: Tabela = Schema + Instância (dados)
- **PRIMARY KEY**: Único + Not NULL + Índice automático
- **FOREIGN KEY**: Integridade referencial + Cascades
- **NULL**: ≠ 0, ≠ '', ≠ false (use IS NULL para comparar)
- **COALESCE**: Primeiro valor não-nulo
- **NULLIF**: Retorna NULL se iguais (útil para evitar divisão por zero)
