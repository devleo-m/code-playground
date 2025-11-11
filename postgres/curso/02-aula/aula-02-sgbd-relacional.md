# **Aula 2: Conceitos de SGBD Relacional**

## 🎯 Objetivo da Aula

Compreender em profundidade o que são Sistemas de Gerenciamento de Banco de Dados Relacionais (SGBDR), entender o modelo relacional proposto por E.F. Codd e dominar os conceitos fundamentais que regem esse modelo.

---

## 📚 O que é um SGBD?

**SGBD** = **Sistema de Gerenciamento de Banco de Dados**

Um SGBD é um software que funciona como intermediário entre você (ou sua aplicação) e os dados armazenados. Pense nele como um bibliotecário digital que:

- **Armazena** dados de forma organizada
- **Protege** os dados contra perdas e acessos não autorizados
- **Gerencia** múltiplos usuários acessando dados simultaneamente
- **Garante** que os dados permaneçam consistentes
- **Facilita** consultas e manipulações complexas

### Sem SGBD vs Com SGBD

#### ❌ Sem SGBD (Arquivos simples)

```
app.py → clientes.txt
         pedidos.txt
         produtos.txt
```

**Problemas:**

- Como garantir que dois usuários não modificam o mesmo arquivo ao mesmo tempo?
- Como fazer uma busca complexa (ex: "clientes que compraram produto X e gastaram mais de R$ 1000")?
- Como garantir que um pedido não referencia um cliente inexistente?
- Como recuperar dados se o arquivo corromper?

#### ✅ Com SGBD (PostgreSQL)

```
app.py → PostgreSQL (SGBD) → Dados organizados
                            → Controle de acesso
                            → Transações ACID
                            → Consultas SQL
                            → Backups automáticos
```

**Benefícios:**

- Controle de concorrência automático
- Linguagem SQL para consultas complexas
- Integridade referencial (chaves estrangeiras)
- Recuperação de falhas
- Otimização automática de consultas

---

## 👨‍🔬 E.F. Codd e o Modelo Relacional (1970)

### Quem foi E.F. Codd?

**Edgar Frank Codd** (1923-2003) foi um cientista da computação britânico que trabalhava na IBM. Em **1970**, ele publicou um artigo revolucionário:

**"A Relational Model of Data for Large Shared Data Banks"**

Este artigo mudou completamente a forma como pensamos sobre bancos de dados.

### 🎯 Qual era o problema que Codd queria resolver?

Antes de 1970, os bancos de dados eram:

- **Hierárquicos**: Dados organizados em árvores (pais e filhos)
- **Em rede**: Dados conectados de forma complexa

**Problemas desses modelos:**

1. Dependência física: Mudar como os dados eram armazenados quebrava as aplicações
2. Consultas complexas: Era muito difícil fazer perguntas complexas aos dados
3. Redundância: Mesmos dados repetidos em vários lugares
4. Manutenção difícil: Adicionar novos tipos de relacionamentos era complicado

### 💡 A Solução de Codd: O Modelo Relacional

Codd propôs que:

1. Dados deveriam ser organizados em **relações** (tabelas)
2. Cada relação é um conjunto de **tuplas** (linhas)
3. Cada tupla contém **atributos** (colunas)
4. Relacionamentos são representados por **valores** (não por ponteiros físicos)

---

## 🏗️ Estrutura do Modelo Relacional

### 1. Relação (Tabela)

Uma **relação** é como uma tabela. Em termos matemáticos, é um conjunto de tuplas.

```
Relação: CLIENTES
┌────┬─────────────┬──────────────────┬────────┐
│ ID │ NOME        │ EMAIL            │ IDADE  │
├────┼─────────────┼──────────────────┼────────┤
│ 1  │ João Silva  │ joao@email.com   │ 25     │
│ 2  │ Maria Costa │ maria@email.com  │ 30     │
│ 3  │ Pedro Lima  │ pedro@email.com  │ 22     │
└────┴─────────────┴──────────────────┴────────┘
```

**Características de uma Relação:**

- Tem um nome único (ex: CLIENTES)
- É composta por tuplas (linhas)
- Cada tupla representa uma entidade (um cliente)
- Ordem das tuplas não importa (a linha 1 poderia ser a linha 3)

### 2. Tupla (Linha/Registro)

Uma **tupla** é uma linha individual na tabela. Representa uma instância única de dados.

```
Uma tupla da relação CLIENTES:
┌────┬─────────────┬──────────────────┬────────┐
│ 2  │ Maria Costa │ maria@email.com  │ 30     │
└────┴─────────────┴──────────────────┴────────┘
```

### 3. Atributo (Coluna/Campo)

Um **atributo** é uma propriedade ou característica da entidade.

```
Atributos da relação CLIENTES:
- ID (identificador único)
- NOME (nome completo)
- EMAIL (endereço de email)
- IDADE (idade em anos)
```

**Cada atributo tem:**

- **Nome**: ID, NOME, EMAIL, IDADE
- **Domínio**: Conjunto de valores válidos
  - ID: números inteiros positivos (1, 2, 3, ...)
  - NOME: texto até 100 caracteres
  - EMAIL: texto no formato email
  - IDADE: números inteiros de 0 a 120

### 4. Domínio

O **domínio** é o conjunto de valores permitidos para um atributo.

```
Exemplos de domínios:

Atributo: IDADE
Domínio: {0, 1, 2, 3, ..., 120}

Atributo: SEXO
Domínio: {'M', 'F', 'Outro', 'Não informado'}

Atributo: STATUS_PEDIDO
Domínio: {'Pendente', 'Processando', 'Enviado', 'Entregue', 'Cancelado'}

Atributo: EMAIL
Domínio: Strings no formato "usuario@dominio.extensao"
```

---

## 🔑 Conceitos Fundamentais: Chaves

### 1. Superchave (Superkey)

Uma **superchave** é qualquer conjunto de atributos que identifica unicamente uma tupla.

```
Relação: CLIENTES
┌────┬─────────────┬──────────────────┬────────┬──────────────┐
│ ID │ NOME        │ EMAIL            │ IDADE  │ CPF          │
├────┼─────────────┼──────────────────┼────────┼──────────────┤
│ 1  │ João Silva  │ joao@email.com   │ 25     │ 111.111.111-11 │
│ 2  │ Maria Costa │ maria@email.com  │ 30     │ 222.222.222-22 │
└────┴─────────────┴──────────────────┴────────┴──────────────┘
```

**Exemplos de superchaves:**

- `{ID}` ✅ - Identifica unicamente
- `{CPF}` ✅ - Identifica unicamente
- `{EMAIL}` ✅ - Identifica unicamente
- `{ID, NOME}` ✅ - Identifica unicamente (mas tem informação redundante)
- `{ID, EMAIL, CPF}` ✅ - Identifica unicamente (mas tem muita redundância)
- `{NOME}` ❌ - Pode haver dois clientes com mesmo nome
- `{IDADE}` ❌ - Muitos clientes podem ter a mesma idade

### 2. Chave Candidata (Candidate Key)

Uma **chave candidata** é uma superchave **mínima** - não tem atributos desnecessários.

```
Chaves candidatas da relação CLIENTES:
- {ID}     ✅ Chave candidata
- {CPF}    ✅ Chave candidata
- {EMAIL}  ✅ Chave candidata

Não são chaves candidatas:
- {ID, NOME}  ❌ Superchave mas não mínima (NOME é redundante)
- {CPF, ID}   ❌ Superchave mas não mínima (pode usar só um deles)
```

### 3. Chave Primária (Primary Key - PK)

A **chave primária** é a chave candidata escolhida para ser o identificador oficial da relação.

```
Relação: CLIENTES
┌────┬─────────────┬──────────────────┬──────────────┐
│ ID │ NOME        │ EMAIL            │ CPF          │  ← ID escolhido como PK
└────┴─────────────┴──────────────────┴──────────────┘
  ↑
  PK
```

**Regras da chave primária:**

- ✅ Deve ser única (não pode haver duplicatas)
- ✅ Não pode ser NULL (deve sempre ter valor)
- ✅ Deve ser imutável (não deve mudar com o tempo)
- ✅ Preferencialmente simples (um único atributo se possível)

**Exemplo em PostgreSQL:**

```sql
CREATE TABLE clientes (
    id SERIAL PRIMARY KEY,        -- Chave primária
    nome VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE,    -- Também único, mas não é a PK
    cpf CHAR(14) UNIQUE,          -- Também único, mas não é a PK
    idade INTEGER
);
```

### 4. Chave Estrangeira (Foreign Key - FK)

Uma **chave estrangeira** é um atributo (ou conjunto de atributos) que cria um relacionamento com outra tabela, referenciando a chave primária dessa tabela.

```
Relação: PEDIDOS
┌────┬────────────┬─────────────┬────────┐
│ ID │ CLIENTE_ID │ PRODUTO     │ VALOR  │
├────┼────────────┼─────────────┼────────┤
│ 1  │ 1          │ Notebook    │ 3000   │  ← CLIENTE_ID=1 referencia
│ 2  │ 2          │ Mouse       │ 50     │    o cliente com ID=1
│ 3  │ 1          │ Teclado     │ 200    │
└────┴────────────┴─────────────┴────────┘
         ↑
         FK (referencia CLIENTES.ID)

Relação: CLIENTES
┌────┬─────────────┬──────────────────┐
│ ID │ NOME        │ EMAIL            │
├────┼─────────────┼──────────────────┤
│ 1  │ João Silva  │ joao@email.com   │  ← Referenciado pelos pedidos 1 e 3
│ 2  │ Maria Costa │ maria@email.com  │  ← Referenciado pelo pedido 2
└────┴─────────────┴──────────────────┘
  ↑
  PK (referenciado por PEDIDOS.CLIENTE_ID)
```

**O que a chave estrangeira garante:**

```sql
-- Criar tabela com chave estrangeira
CREATE TABLE pedidos (
    id SERIAL PRIMARY KEY,
    cliente_id INTEGER NOT NULL,
    produto VARCHAR(100),
    valor DECIMAL(10, 2),
    FOREIGN KEY (cliente_id) REFERENCES clientes(id)
);

-- ✅ PERMITIDO: Cliente 1 existe
INSERT INTO pedidos (cliente_id, produto, valor)
VALUES (1, 'Mouse', 50.00);

-- ❌ ERRO: Cliente 999 não existe!
INSERT INTO pedidos (cliente_id, produto, valor)
VALUES (999, 'Teclado', 100.00);
-- ERRO: insert or update on table "pedidos" violates foreign key constraint

-- ❌ ERRO: Não pode deletar cliente que tem pedidos!
DELETE FROM clientes WHERE id = 1;
-- ERRO: update or delete on table "clientes" violates foreign key constraint
```

---

## 🎭 Tipos de Relacionamentos

### 1. Um para Um (1:1)

Um registro em uma tabela está relacionado a **no máximo um** registro em outra tabela.

```
PESSOAS                    PASSAPORTES
┌────┬──────────┐         ┌────┬────────────┬───────────┐
│ ID │ NOME     │         │ ID │ NUMERO     │ PESSOA_ID │
├────┼──────────┤         ├────┼────────────┼───────────┤
│ 1  │ João     │ ←──────→│ 1  │ BR123456   │ 1         │
│ 2  │ Maria    │ ←──────→│ 2  │ BR789012   │ 2         │
└────┴──────────┘         └────┴────────────┴───────────┘

Cada pessoa tem no máximo um passaporte
Cada passaporte pertence a exatamente uma pessoa
```

**Implementação:**

```sql
CREATE TABLE pessoas (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(100)
);

CREATE TABLE passaportes (
    id SERIAL PRIMARY KEY,
    numero VARCHAR(20) UNIQUE,
    pessoa_id INTEGER UNIQUE,  -- UNIQUE garante 1:1
    FOREIGN KEY (pessoa_id) REFERENCES pessoas(id)
);
```

### 2. Um para Muitos (1:N)

Um registro em uma tabela pode estar relacionado a **vários** registros em outra tabela.

```
CLIENTES                   PEDIDOS
┌────┬──────────┐         ┌────┬────────────┬────────┐
│ ID │ NOME     │         │ ID │ PRODUTO    │ CLI_ID │
├────┼──────────┤         ├────┼────────────┼────────┤
│ 1  │ João     │ ←───────│ 1  │ Notebook   │ 1      │
│    │          │    ┌────│ 2  │ Mouse      │ 1      │
│    │          │    │ ┌──│ 3  │ Teclado    │ 1      │
│ 2  │ Maria    │ ←──┴─┴──│ 4  │ Monitor    │ 2      │
└────┴──────────┘         └────┴────────────┴────────┘

Um cliente pode ter vários pedidos
Cada pedido pertence a exatamente um cliente
```

**Implementação:**

```sql
CREATE TABLE clientes (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(100)
);

CREATE TABLE pedidos (
    id SERIAL PRIMARY KEY,
    produto VARCHAR(100),
    cliente_id INTEGER,  -- Sem UNIQUE, permite múltiplos pedidos
    FOREIGN KEY (cliente_id) REFERENCES clientes(id)
);
```

### 3. Muitos para Muitos (N:M)

Vários registros em uma tabela podem estar relacionados a **vários** registros em outra tabela.

```
ALUNOS                     TURMAS
┌────┬──────────┐         ┌────┬───────────┐
│ ID │ NOME     │         │ ID │ NOME      │
├────┼──────────┤         ├────┼───────────┤
│ 1  │ João     │ ←───┬───│ 1  │ Matemática│
│ 2  │ Maria    │ ←─┐ │┌──│ 2  │ Português │
│ 3  │ Pedro    │ ← │ ││  └────┴───────────┘
└────┴──────────┘   │ ││
                    │ ││  MATRICULAS (Tabela associativa)
                    │ ││  ┌────┬──────────┬──────────┐
                    │ ││  │ ID │ ALUNO_ID │ TURMA_ID │
                    │ ││  ├────┼──────────┼──────────┤
                    └─┼┼──│ 1  │ 1        │ 1        │
                      ││  │ 2  │ 1        │ 2        │
                      │└──│ 3  │ 2        │ 1        │
                      └───│ 4  │ 2        │ 2        │
                          │ 5  │ 3        │ 1        │
                          └────┴──────────┴──────────┘

Um aluno pode estar em várias turmas
Uma turma pode ter vários alunos
```

**Implementação:**

```sql
CREATE TABLE alunos (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(100)
);

CREATE TABLE turmas (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(100)
);

-- Tabela associativa (junction table)
CREATE TABLE matriculas (
    id SERIAL PRIMARY KEY,
    aluno_id INTEGER,
    turma_id INTEGER,
    FOREIGN KEY (aluno_id) REFERENCES alunos(id),
    FOREIGN KEY (turma_id) REFERENCES turmas(id),
    UNIQUE(aluno_id, turma_id)  -- Evita matrícula duplicada
);
```

---

## 🎯 Integridade de Dados

### 1. Integridade de Entidade

**Regra:** A chave primária não pode ser NULL e deve ser única.

```sql
-- ✅ VÁLIDO
INSERT INTO clientes (id, nome) VALUES (1, 'João');

-- ❌ ERRO: Chave primária NULL
INSERT INTO clientes (id, nome) VALUES (NULL, 'Maria');

-- ❌ ERRO: Chave primária duplicada
INSERT INTO clientes (id, nome) VALUES (1, 'Pedro');
```

### 2. Integridade Referencial

**Regra:** Uma chave estrangeira deve referenciar uma chave primária existente ou ser NULL (se permitido).

```sql
-- ✅ VÁLIDO: Cliente 1 existe
INSERT INTO pedidos (cliente_id, produto) VALUES (1, 'Mouse');

-- ❌ ERRO: Cliente 999 não existe
INSERT INTO pedidos (cliente_id, produto) VALUES (999, 'Teclado');

-- ✅ VÁLIDO: NULL permitido se definido assim
INSERT INTO pedidos (cliente_id, produto) VALUES (NULL, 'Brinde');
```

### 3. Integridade de Domínio

**Regra:** Os valores devem estar dentro do domínio definido para o atributo.

```sql
CREATE TABLE produtos (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(100),
    preco DECIMAL(10, 2) CHECK (preco > 0),  -- Domínio: preço positivo
    categoria VARCHAR(20) CHECK (categoria IN ('Eletrônico', 'Livro', 'Roupa')),
    estoque INTEGER CHECK (estoque >= 0)
);

-- ✅ VÁLIDO
INSERT INTO produtos (nome, preco, categoria, estoque)
VALUES ('Mouse', 50.00, 'Eletrônico', 10);

-- ❌ ERRO: Preço negativo
INSERT INTO produtos (nome, preco, categoria, estoque)
VALUES ('Teclado', -100.00, 'Eletrônico', 5);

-- ❌ ERRO: Categoria inválida
INSERT INTO produtos (nome, preco, categoria, estoque)
VALUES ('Caderno', 10.00, 'Papelaria', 20);
```

### 4. Integridade Definida pelo Usuário

Regras personalizadas específicas do negócio.

```sql
-- Exemplo: Data de nascimento deve ser no passado
CREATE TABLE funcionarios (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(100),
    data_nascimento DATE CHECK (data_nascimento < CURRENT_DATE),
    salario DECIMAL(10, 2) CHECK (salario >= 1320.00)  -- Salário mínimo
);

-- Exemplo: Data de entrega deve ser após data do pedido
CREATE TABLE entregas (
    id SERIAL PRIMARY KEY,
    pedido_id INTEGER,
    data_pedido DATE,
    data_entrega DATE,
    CHECK (data_entrega >= data_pedido)
);
```

---

## 📐 Modelo Relacional: Principais Conceitos

### 1. Independência de Dados

O modelo relacional separa:

- **Nível lógico**: Como os dados são organizados (tabelas, colunas)
- **Nível físico**: Como os dados são armazenados no disco

**Benefício:** Você pode mudar a forma de armazenamento sem mudar as aplicações!

```sql
-- A aplicação sempre usa a mesma query
SELECT * FROM clientes WHERE idade > 25;

-- Mas o PostgreSQL pode:
-- - Adicionar índices
-- - Mudar particionamento
-- - Reorganizar armazenamento
-- Sem afetar a aplicação!
```

### 2. Operações Relacionais

O modelo relacional define operações para manipular dados:

#### **Seleção (σ)** - Filtra linhas

```sql
SELECT * FROM clientes WHERE idade > 25;
```

#### **Projeção (π)** - Seleciona colunas

```sql
SELECT nome, email FROM clientes;
```

#### **Junção (⋈)** - Combina tabelas

```sql
SELECT c.nome, p.produto
FROM clientes c
JOIN pedidos p ON c.id = p.cliente_id;
```

#### **União (∪)** - Combina resultados

```sql
SELECT nome FROM clientes_sp
UNION
SELECT nome FROM clientes_rj;
```

#### **Diferença (−)** - Subtrai resultados

```sql
SELECT nome FROM todos_clientes
EXCEPT
SELECT nome FROM clientes_inativos;
```

---

## 🏆 Vantagens do Modelo Relacional

### 1. ✅ Simplicidade

Tabelas são intuitivas e fáceis de entender

### 2. ✅ Flexibilidade

Consultas complexas sem precisar conhecer a estrutura física

### 3. ✅ Integridade

Restrições garantem dados consistentes

### 4. ✅ Independência

Mudanças físicas não afetam aplicações

### 5. ✅ Matemática Sólida

Baseado em teoria de conjuntos e álgebra relacional

### 6. ✅ SQL

Linguagem padrão poderosa e declarativa

---

## 📊 Resumo dos Conceitos-Chave

| Conceito              | Definição                   | Exemplo                         |
| --------------------- | --------------------------- | ------------------------------- |
| **Relação**           | Tabela com linhas e colunas | CLIENTES                        |
| **Tupla**             | Linha/registro individual   | `(1, 'João', 'joao@email.com')` |
| **Atributo**          | Coluna/campo                | NOME, EMAIL, IDADE              |
| **Domínio**           | Valores permitidos          | IDADE: 0-120                    |
| **Chave Primária**    | Identificador único         | ID                              |
| **Chave Estrangeira** | Referência a outra tabela   | CLIENTE_ID                      |
| **Integridade**       | Regras de consistência      | PK não NULL, FK válida          |
| **Relacionamento**    | Conexão entre tabelas       | 1:1, 1:N, N:M                   |

---

## 🎓 Conclusão

O modelo relacional de E.F. Codd revolucionou o mundo dos bancos de dados ao:

1. **Simplificar** a representação de dados
2. **Garantir** integridade através de regras claras
3. **Separar** lógica de implementação física
4. **Fornecer** base matemática sólida
5. **Permitir** consultas flexíveis e poderosas

PostgreSQL implementa fielmente esse modelo, respeitando os princípios definidos por Codd e adicionando recursos modernos.

---

## 🔑 Conceitos para Memorizar

- **SGBD**: Software que gerencia banco de dados
- **E.F. Codd**: Criador do modelo relacional (1970)
- **Relação**: Tabela (estrutura matemática, não apenas visual)
- **Chave Primária**: Identificador único e obrigatório
- **Chave Estrangeira**: Cria relacionamentos entre tabelas
- **Integridade**: Garantia de que dados seguem regras definidas
- **Independência de dados**: Lógica separada da implementação física


