# Aula 3: Data Definition Language (DDL)

## Introdução

Nesta aula, você aprenderá sobre **DDL (Data Definition Language)**, a parte do SQL responsável por criar, modificar e gerenciar a estrutura do banco de dados. Enquanto o DML (Data Manipulation Language) trabalha com os **dados** dentro das tabelas, o DDL trabalha com a **estrutura** das tabelas e outros objetos do banco.

Dominar DDL é essencial para qualquer desenvolvedor ou administrador de banco de dados, pois permite criar e adaptar a estrutura do banco conforme as necessidades da aplicação evoluem.

---

## 1. O que é Data Definition Language (DDL)?

**Data Definition Language (DDL)** é um subconjunto do SQL usado para definir e gerenciar a estrutura de objetos do banco de dados. Os comandos DDL incluem CREATE, ALTER, DROP e TRUNCATE, que são usados para criar, modificar, deletar e esvaziar estruturas de banco de dados como tabelas, índices, views e schemas.

### Características do DDL

1. **Operações Estruturais**: DDL trabalha com a estrutura, não com os dados em si
2. **Comandos Principais**: CREATE, ALTER, DROP, TRUNCATE
3. **Impacto Imediato**: Comandos DDL geralmente resultam em mudanças imediatas na estrutura do banco
4. **Transações**: Em muitos SGBDs, comandos DDL são auto-commit (não podem ser revertidos com ROLLBACK)
5. **Permissões**: Geralmente requerem privilégios administrativos

### DDL vs DML

| Aspecto | DDL (Data Definition Language) | DML (Data Manipulation Language) |
|--------|-------------------------------|----------------------------------|
| **Foco** | Estrutura do banco | Dados dentro das tabelas |
| **Comandos** | CREATE, ALTER, DROP, TRUNCATE | SELECT, INSERT, UPDATE, DELETE |
| **Frequência** | Menos frequente (mudanças de schema) | Muito frequente (operações diárias) |
| **Reversibilidade** | Difícil ou impossível | Pode ser revertido (em transações) |
| **Impacto** | Afeta toda a estrutura | Afeta registros específicos |

---

## 2. CREATE TABLE - Criando Tabelas

O comando **CREATE TABLE** é usado para criar uma nova tabela no banco de dados. Ele define a estrutura da tabela, incluindo nomes de colunas, tipos de dados e constraints (restrições).

### 2.1 Sintaxe Básica

```sql
CREATE TABLE nome_tabela (
    coluna1 tipo_dados [constraints],
    coluna2 tipo_dados [constraints],
    ...
);
```

### 2.2 Exemplo Simples

```sql
-- Criar uma tabela simples de categorias
CREATE TABLE categorias (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    nome TEXT NOT NULL UNIQUE,
    descricao TEXT
);
```

**Explicação linha por linha:**
- `CREATE TABLE categorias`: Cria uma tabela chamada "categorias"
- `id INTEGER PRIMARY KEY AUTOINCREMENT`: Coluna id do tipo INTEGER, chave primária que incrementa automaticamente
- `nome TEXT NOT NULL UNIQUE`: Coluna nome do tipo TEXT, não pode ser NULL e deve ser único
- `descricao TEXT`: Coluna descricao do tipo TEXT, pode ser NULL

### 2.3 CREATE TABLE Completo - Exemplo com o Banco Biblioteca

Vamos criar uma tabela completa usando o exemplo do nosso banco de biblioteca:

```sql
-- Criar tabela de livros com todas as constraints
CREATE TABLE livros (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    titulo TEXT NOT NULL,
    isbn TEXT UNIQUE,
    ano_publicacao INTEGER,
    editora TEXT,
    autor_id INTEGER,
    categoria_id INTEGER,
    quantidade_disponivel INTEGER DEFAULT 0,
    FOREIGN KEY (autor_id) REFERENCES autores(id),
    FOREIGN KEY (categoria_id) REFERENCES categorias(id)
);
```

**Análise detalhada:**

1. **`id INTEGER PRIMARY KEY AUTOINCREMENT`**
   - Tipo: INTEGER
   - PRIMARY KEY: Identificador único da tabela
   - AUTOINCREMENT: Valor incrementa automaticamente (SQLite)

2. **`titulo TEXT NOT NULL`**
   - Tipo: TEXT
   - NOT NULL: Campo obrigatório, não pode ser vazio

3. **`isbn TEXT UNIQUE`**
   - Tipo: TEXT
   - UNIQUE: Cada valor deve ser único (mas pode ser NULL)

4. **`ano_publicacao INTEGER`**
   - Tipo: INTEGER
   - Pode ser NULL (não especificado)

5. **`quantidade_disponivel INTEGER DEFAULT 0`**
   - Tipo: INTEGER
   - DEFAULT 0: Se não especificado, assume valor 0

6. **`FOREIGN KEY (autor_id) REFERENCES autores(id)`**
   - Chave estrangeira: autor_id deve existir na tabela autores
   - Garante integridade referencial

### 2.4 Tipos de Constraints em CREATE TABLE

#### PRIMARY KEY
Identifica unicamente cada registro na tabela.

```sql
CREATE TABLE usuarios (
    id INTEGER PRIMARY KEY,
    nome TEXT NOT NULL
);
```

#### FOREIGN KEY
Estabelece relacionamento com outra tabela.

```sql
CREATE TABLE emprestimos (
    id INTEGER PRIMARY KEY,
    livro_id INTEGER NOT NULL,
    usuario_id INTEGER NOT NULL,
    FOREIGN KEY (livro_id) REFERENCES livros(id),
    FOREIGN KEY (usuario_id) REFERENCES usuarios(id)
);
```

#### UNIQUE
Garante que valores na coluna sejam únicos.

```sql
CREATE TABLE usuarios (
    id INTEGER PRIMARY KEY,
    email TEXT UNIQUE NOT NULL
);
```

#### NOT NULL
Garante que a coluna não aceite valores NULL.

```sql
CREATE TABLE livros (
    id INTEGER PRIMARY KEY,
    titulo TEXT NOT NULL
);
```

#### CHECK
Valida valores baseado em uma condição.

```sql
CREATE TABLE livros (
    id INTEGER PRIMARY KEY,
    quantidade_disponivel INTEGER CHECK (quantidade_disponivel >= 0)
);
```

#### DEFAULT
Define um valor padrão quando nenhum valor é especificado.

```sql
CREATE TABLE emprestimos (
    id INTEGER PRIMARY KEY,
    data_emprestimo DATE DEFAULT CURRENT_DATE,
    status TEXT DEFAULT 'ativo'
);
```

### 2.5 Criando Tabela com Múltiplas Constraints

```sql
-- Tabela completa com todas as constraints
CREATE TABLE emprestimos (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    livro_id INTEGER NOT NULL,
    usuario_id INTEGER NOT NULL,
    data_emprestimo DATE DEFAULT CURRENT_DATE NOT NULL,
    data_devolucao_prevista DATE,
    data_devolucao_real DATE,
    status TEXT DEFAULT 'ativo' CHECK (status IN ('ativo', 'devolvido', 'atrasado')),
    FOREIGN KEY (livro_id) REFERENCES livros(id),
    FOREIGN KEY (usuario_id) REFERENCES usuarios(id),
    CHECK (data_devolucao_prevista IS NULL OR data_devolucao_prevista >= data_emprestimo)
);
```

---

## 3. ALTER TABLE - Modificando Tabelas

O comando **ALTER TABLE** permite modificar a estrutura de uma tabela existente sem perder os dados. É uma das operações DDL mais importantes, pois permite adaptar o banco conforme as necessidades mudam.

### 3.1 Sintaxe Básica

```sql
ALTER TABLE nome_tabela
[operação];
```

### 3.2 Adicionar Coluna (ADD COLUMN)

Adiciona uma nova coluna à tabela existente.

```sql
-- Adicionar coluna de preço à tabela livros
ALTER TABLE livros
ADD COLUMN preco REAL;

-- Adicionar coluna com constraints
ALTER TABLE livros
ADD COLUMN edicao INTEGER DEFAULT 1 NOT NULL;
```

**⚠️ Importante no SQLite:**
- SQLite tem limitações: você pode apenas adicionar colunas ao final da tabela
- Não pode adicionar colunas com NOT NULL sem DEFAULT em tabelas que já têm dados
- Não pode adicionar colunas com UNIQUE ou PRIMARY KEY diretamente

### 3.3 Renomear Coluna (RENAME COLUMN)

Renomeia uma coluna existente.

```sql
-- Renomear coluna (SQLite 3.25.0+)
ALTER TABLE livros
RENAME COLUMN quantidade_disponivel TO estoque;
```

**⚠️ Nota**: Versões antigas do SQLite não suportam RENAME COLUMN. Nesses casos, é necessário recriar a tabela.

### 3.4 Remover Coluna (DROP COLUMN)

Remove uma coluna da tabela.

```sql
-- Remover coluna (SQLite 3.35.0+)
ALTER TABLE livros
DROP COLUMN editora;
```

**⚠️ CUIDADO**: Esta operação é **destrutiva** e remove a coluna e todos os seus dados permanentemente!

**⚠️ Limitações no SQLite:**
- SQLite não suporta DROP COLUMN diretamente em versões antigas
- Para remover coluna em versões antigas, é necessário:
  1. Criar nova tabela sem a coluna
  2. Copiar dados
  3. Deletar tabela antiga
  4. Renomear nova tabela

### 3.5 Adicionar Constraint

Adiciona uma constraint a uma tabela existente.

```sql
-- Adicionar constraint UNIQUE
CREATE UNIQUE INDEX idx_livros_isbn ON livros(isbn);

-- Adicionar constraint CHECK (SQLite não suporta diretamente, use trigger)
-- Em outros SGBDs:
-- ALTER TABLE livros ADD CONSTRAINT check_quantidade CHECK (quantidade_disponivel >= 0);
```

### 3.6 Modificar Tipo de Dados

**⚠️ SQLite não suporta alterar tipo de dados diretamente.**

Em outros SGBDs (PostgreSQL, MySQL):

```sql
-- PostgreSQL
ALTER TABLE livros
ALTER COLUMN ano_publicacao TYPE VARCHAR(4);

-- MySQL
ALTER TABLE livros
MODIFY COLUMN ano_publicacao VARCHAR(4);
```

**No SQLite**: Para alterar tipo de dados, é necessário recriar a tabela.

### 3.7 Exemplos Práticos de ALTER TABLE

#### Exemplo 1: Adicionar Campo de Data de Cadastro

```sql
-- Adicionar data de cadastro aos livros
ALTER TABLE livros
ADD COLUMN data_cadastro DATE DEFAULT CURRENT_DATE;
```

#### Exemplo 2: Adicionar Campo de Observações

```sql
-- Adicionar campo de observações
ALTER TABLE livros
ADD COLUMN observacoes TEXT;
```

#### Exemplo 3: Adicionar Campo com Valor Padrão

```sql
-- Adicionar campo de ativo/inativo
ALTER TABLE livros
ADD COLUMN ativo INTEGER DEFAULT 1;
```

### 3.8 Limitações do ALTER TABLE no SQLite

SQLite tem limitações significativas em relação a ALTER TABLE:

1. **Pode fazer:**
   - ADD COLUMN (ao final da tabela)
   - RENAME TABLE
   - RENAME COLUMN (versões 3.25.0+)
   - DROP COLUMN (versões 3.35.0+)

2. **Não pode fazer diretamente:**
   - Modificar tipo de dados
   - Adicionar constraints complexas
   - Remover constraints
   - Reordenar colunas

3. **Solução para limitações:**
   - Recriar a tabela com a estrutura desejada
   - Copiar dados da tabela antiga
   - Deletar tabela antiga
   - Renomear nova tabela

---

## 4. DROP TABLE - Removendo Tabelas

O comando **DROP TABLE** remove completamente uma tabela do banco de dados, incluindo sua estrutura e todos os dados.

### 4.1 Sintaxe Básica

```sql
DROP TABLE [IF EXISTS] nome_tabela;
```

### 4.2 Exemplo Simples

```sql
-- Remover uma tabela
DROP TABLE tabela_temporaria;

-- Remover com verificação de existência (recomendado)
DROP TABLE IF EXISTS tabela_temporaria;
```

### 4.3 Comportamento do DROP TABLE

1. **Remove tudo**: Estrutura, dados, índices, triggers associados
2. **Irreversível**: Não pode ser desfeito (a menos que tenha backup)
3. **Dependências**: Pode falhar se houver chaves estrangeiras referenciando a tabela

### 4.4 Exemplo Prático

```sql
-- Criar tabela temporária para testes
CREATE TABLE livros_temporarios (
    id INTEGER PRIMARY KEY,
    titulo TEXT
);

-- Inserir alguns dados
INSERT INTO livros_temporarios (titulo)
VALUES ('Livro Teste 1'), ('Livro Teste 2');

-- Remover a tabela (remove estrutura E dados)
DROP TABLE livros_temporarios;
```

### 4.5 Cuidados com DROP TABLE

**⚠️ CRÍTICO**: DROP TABLE é uma operação **destrutiva e irreversível**!

**Sempre:**
- Faça backup antes de executar DROP TABLE
- Verifique se não há dependências (chaves estrangeiras)
- Use `IF EXISTS` para evitar erros se a tabela não existir
- Teste em ambiente de desenvolvimento primeiro

**Exemplo seguro:**

```sql
-- Verificar se há dependências primeiro
SELECT name FROM sqlite_master 
WHERE type='table' AND sql LIKE '%livros%';

-- Depois, se seguro, remover
DROP TABLE IF EXISTS livros_temporarios;
```

### 4.6 DROP TABLE vs DELETE FROM

| Operação | O que remove | Estrutura | Reversível |
|----------|--------------|-----------|------------|
| **DELETE FROM** | Apenas dados (registros) | Mantém estrutura | Sim (em transações) |
| **DROP TABLE** | Dados + estrutura | Remove tudo | Não (sem backup) |

```sql
-- DELETE: Remove apenas os dados
DELETE FROM livros WHERE id = 1;
-- Tabela 'livros' ainda existe, apenas sem o registro id=1

-- DROP TABLE: Remove tabela inteira
DROP TABLE livros;
-- Tabela 'livros' não existe mais
```

---

## 5. TRUNCATE TABLE - Limpando Dados

O comando **TRUNCATE TABLE** remove todos os dados de uma tabela, mas mantém a estrutura (colunas, constraints, índices).

### 5.1 Sintaxe Básica

```sql
TRUNCATE TABLE nome_tabela;
```

### 5.2 Características do TRUNCATE

1. **Remove todos os dados**: Esvazia a tabela completamente
2. **Mantém estrutura**: Colunas, constraints e índices permanecem
3. **Mais rápido que DELETE**: Não gera logs individuais de cada linha
4. **Não pode ser revertido**: Geralmente não pode ser revertido em transações
5. **Reinicia contadores**: AUTO_INCREMENT/IDENTITY volta ao valor inicial

### 5.3 TRUNCATE no SQLite

**⚠️ IMPORTANTE**: SQLite **não possui** comando TRUNCATE TABLE nativo!

**Alternativas no SQLite:**

#### Opção 1: DELETE FROM (mais comum)
```sql
-- Remove todos os dados, mantém estrutura
DELETE FROM livros_temporarios;
```

#### Opção 2: DROP + CREATE (mais rápido para grandes volumes)
```sql
-- Salvar estrutura
CREATE TABLE livros_backup AS SELECT * FROM livros_temporarios WHERE 1=0;

-- Remover tabela
DROP TABLE livros_temporarios;

-- Recriar (estrutura vazia)
CREATE TABLE livros_temporarios AS SELECT * FROM livros_backup;

-- Limpar backup
DROP TABLE livros_backup;
```

### 5.4 TRUNCATE vs DELETE

| Aspecto | TRUNCATE TABLE | DELETE FROM |
|---------|----------------|-------------|
| **Dados removidos** | Todos (não pode usar WHERE) | Todos ou filtrados (pode usar WHERE) |
| **Estrutura** | Mantida | Mantida |
| **Performance** | Mais rápido | Mais lento |
| **Logs** | Mínimos | Logs de cada linha |
| **Transações** | Geralmente auto-commit | Pode ser revertido |
| **Triggers** | Não executa | Executa |
| **Contadores** | Reinicia | Mantém sequência |

### 5.5 Exemplo Prático

```sql
-- Criar tabela de teste
CREATE TABLE teste_truncate (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    nome TEXT
);

-- Inserir dados
INSERT INTO teste_truncate (nome) VALUES 
    ('Item 1'), ('Item 2'), ('Item 3');

-- Ver dados
SELECT * FROM teste_truncate;
-- Resultado: 3 registros

-- No SQLite, usar DELETE (equivalente a TRUNCATE)
DELETE FROM teste_truncate;

-- Ver dados novamente
SELECT * FROM teste_truncate;
-- Resultado: 0 registros (tabela vazia, mas estrutura mantida)

-- Verificar estrutura ainda existe
SELECT sql FROM sqlite_master WHERE name = 'teste_truncate';
-- Resultado: Estrutura ainda existe

-- Limpar tabela de teste
DROP TABLE teste_truncate;
```

### 5.6 Quando Usar TRUNCATE

**Use TRUNCATE quando:**
- Precisa remover **todos** os dados de uma tabela
- Quer manter a estrutura
- Performance é importante (tabelas grandes)
- Não precisa de logs detalhados

**Não use TRUNCATE quando:**
- Precisa remover apenas alguns registros (use DELETE com WHERE)
- Precisa que triggers sejam executados
- Precisa reverter a operação em uma transação
- Trabalha com SQLite (use DELETE FROM)

---

## 6. Índices - CREATE INDEX e DROP INDEX

Índices são estruturas que melhoram a performance de consultas, permitindo acesso mais rápido aos dados.

### 6.1 CREATE INDEX

Cria um índice em uma ou mais colunas.

#### Sintaxe Básica

```sql
CREATE INDEX nome_indice
ON nome_tabela (coluna1, coluna2, ...);
```

#### Exemplo Simples

```sql
-- Criar índice na coluna autor_id para melhorar JOINs
CREATE INDEX idx_livros_autor ON livros(autor_id);

-- Criar índice composto (múltiplas colunas)
CREATE INDEX idx_emprestimos_usuario_status 
ON emprestimos(usuario_id, status);
```

#### Índice Único

```sql
-- Criar índice único (garante unicidade)
CREATE UNIQUE INDEX idx_livros_isbn ON livros(isbn);
```

### 6.2 Quando Criar Índices

**Crie índices em:**
- Colunas usadas frequentemente em WHERE
- Colunas usadas em JOINs (chaves estrangeiras)
- Colunas usadas em ORDER BY
- Colunas usadas em GROUP BY
- Colunas com muitos valores únicos

**Exemplo do banco biblioteca:**

```sql
-- Índices úteis para o banco biblioteca
CREATE INDEX idx_livros_autor ON livros(autor_id);
CREATE INDEX idx_livros_categoria ON livros(categoria_id);
CREATE INDEX idx_emprestimos_livro ON emprestimos(livro_id);
CREATE INDEX idx_emprestimos_usuario ON emprestimos(usuario_id);
CREATE INDEX idx_emprestimos_status ON emprestimos(status);
```

### 6.3 DROP INDEX

Remove um índice do banco de dados.

#### Sintaxe

```sql
DROP INDEX [IF EXISTS] nome_indice;
```

#### Exemplo

```sql
-- Remover índice
DROP INDEX idx_livros_autor;

-- Remover com verificação
DROP INDEX IF EXISTS idx_livros_autor;
```

### 6.4 Considerações sobre Índices

**Vantagens:**
- Melhoram performance de consultas (SELECT)
- Aceleram JOINs
- Aceleram ORDER BY e GROUP BY

**Desvantagens:**
- Ocupam espaço em disco
- Podem atrasar INSERT, UPDATE, DELETE
- Requerem manutenção quando dados mudam

**Regra geral:**
- Índices são bons para leitura (SELECT)
- Índices podem atrasar escrita (INSERT, UPDATE, DELETE)
- Crie índices baseado em padrões de consulta reais

---

## 7. Exemplos Práticos Completos

### Exemplo 1: Criar Sistema de Avaliações

```sql
-- Criar tabela de avaliações de livros
CREATE TABLE avaliacoes (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    livro_id INTEGER NOT NULL,
    usuario_id INTEGER NOT NULL,
    nota INTEGER NOT NULL CHECK (nota >= 1 AND nota <= 5),
    comentario TEXT,
    data_avaliacao DATE DEFAULT CURRENT_DATE,
    FOREIGN KEY (livro_id) REFERENCES livros(id),
    FOREIGN KEY (usuario_id) REFERENCES usuarios(id),
    UNIQUE(livro_id, usuario_id)  -- Um usuário só pode avaliar um livro uma vez
);

-- Criar índice para melhorar consultas
CREATE INDEX idx_avaliacoes_livro ON avaliacoes(livro_id);
CREATE INDEX idx_avaliacoes_usuario ON avaliacoes(usuario_id);
```

### Exemplo 2: Modificar Tabela Existente

```sql
-- Adicionar campo de avaliação média aos livros
ALTER TABLE livros
ADD COLUMN avaliacao_media REAL;

-- Adicionar campo de número de avaliações
ALTER TABLE livros
ADD COLUMN total_avaliacoes INTEGER DEFAULT 0;
```

### Exemplo 3: Criar Tabela de Reservas

```sql
-- Criar tabela de reservas
CREATE TABLE reservas (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    livro_id INTEGER NOT NULL,
    usuario_id INTEGER NOT NULL,
    data_reserva DATE DEFAULT CURRENT_DATE NOT NULL,
    data_limite DATE,
    status TEXT DEFAULT 'ativa' CHECK (status IN ('ativa', 'cancelada', 'concluida')),
    FOREIGN KEY (livro_id) REFERENCES livros(id),
    FOREIGN KEY (usuario_id) REFERENCES usuarios(id)
);

-- Criar índices
CREATE INDEX idx_reservas_livro ON reservas(livro_id);
CREATE INDEX idx_reservas_usuario ON reservas(usuario_id);
CREATE INDEX idx_reservas_status ON reservas(status);
```

### Exemplo 4: Limpar e Recriar Estrutura

```sql
-- Remover tabela de teste se existir
DROP TABLE IF EXISTS teste_estrutura;

-- Criar nova estrutura
CREATE TABLE teste_estrutura (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    nome TEXT NOT NULL,
    valor INTEGER DEFAULT 0
);

-- Adicionar coluna depois
ALTER TABLE teste_estrutura
ADD COLUMN descricao TEXT;

-- Limpar dados (equivalente a TRUNCATE no SQLite)
DELETE FROM teste_estrutura;

-- Remover tabela
DROP TABLE teste_estrutura;
```

---

## 8. Boas Práticas em DDL

### 8.1 Nomenclatura

- **Tabelas**: Nomes no plural, descritivos (ex: `livros`, `usuarios`)
- **Colunas**: Nomes no singular, descritivos (ex: `titulo`, `data_cadastro`)
- **Índices**: Prefixo `idx_` seguido do nome (ex: `idx_livros_autor`)
- **Constraints**: Nomes descritivos quando nomeados explicitamente

### 8.2 Planejamento

- **Pense antes de criar**: Planeje a estrutura antes de executar CREATE TABLE
- **Considere relacionamentos**: Pense em FOREIGN KEYs desde o início
- **Antecipe necessidades**: Considere campos que podem ser necessários no futuro

### 8.3 Constraints

- **Use constraints**: PRIMARY KEY, FOREIGN KEY, NOT NULL, UNIQUE, CHECK
- **Garanta integridade**: Constraints previnem dados inválidos
- **Documente**: Comente constraints complexas

### 8.4 Índices

- **Crie com moderação**: Muitos índices podem atrasar escrita
- **Monitore uso**: Remova índices não utilizados
- **Índices compostos**: Use quando múltiplas colunas são consultadas juntas

### 8.5 Segurança

- **Backup antes de DDL**: Sempre faça backup antes de operações destrutivas
- **Teste primeiro**: Teste em ambiente de desenvolvimento
- **Versionamento**: Mantenha histórico de mudanças de schema

---

## Conclusão

Nesta aula, você aprendeu:

1. **O que é DDL**: Linguagem para definir estrutura do banco
2. **CREATE TABLE**: Criar tabelas com colunas, tipos e constraints
3. **ALTER TABLE**: Modificar estrutura de tabelas existentes
4. **DROP TABLE**: Remover tabelas completamente
5. **TRUNCATE TABLE**: Limpar dados mantendo estrutura (no SQLite, use DELETE)
6. **Índices**: CREATE INDEX e DROP INDEX para melhorar performance
7. **Constraints**: PRIMARY KEY, FOREIGN KEY, UNIQUE, NOT NULL, CHECK, DEFAULT
8. **Boas práticas**: Nomenclatura, planejamento, segurança

Esses comandos DDL são fundamentais para criar e gerenciar a estrutura do banco de dados. Pratique criando tabelas, modificando estruturas e entendendo o impacto de cada operação.

**Próximo Passo**: Na próxima seção, vamos simplificar esses conceitos usando analogias do dia a dia para facilitar o entendimento.

---

**⚠️ Lembrete Importante**: Comandos DDL podem ser destrutivos. Sempre faça backup e teste em ambiente de desenvolvimento antes de executar em produção!


