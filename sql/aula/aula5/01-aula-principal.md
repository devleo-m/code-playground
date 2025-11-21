# Aula 5: Data Constraints (Restrições de Dados)

## Introdução

Nesta aula, você aprenderá sobre **Data Constraints** (Restrições de Dados), que são regras aplicadas a colunas ou tabelas para garantir a integridade e consistência dos dados no banco de dados. Constraints são fundamentais para manter a qualidade dos dados e prevenir erros que podem comprometer a confiabilidade do sistema.

Data constraints são essenciais para:
- Garantir que os dados atendam a critérios específicos
- Manter relacionamentos válidos entre tabelas
- Prevenir inserção de dados inválidos ou inconsistentes
- Enforçar regras de negócio diretamente no banco de dados
- Melhorar a confiabilidade e manutenibilidade do sistema

Dominar constraints é fundamental para qualquer desenvolvedor ou administrador de banco de dados, pois permite criar sistemas robustos que protegem a integridade dos dados automaticamente.

---

## 1. O que são Data Constraints?

**Data Constraints** (Restrições de Dados) são regras aplicadas a colunas ou tabelas que definem limitações sobre os dados que podem ser inseridos, atualizados ou deletados. Elas garantem que os dados no banco atendam a critérios específicos e mantenham relacionamentos válidos.

### Características das Constraints

1. **Enforçadas pelo Banco de Dados**: As regras são verificadas automaticamente pelo SGBD
2. **Previnem Dados Inválidos**: Bloqueiam inserções/atualizações que violam as regras
3. **Mantêm Integridade**: Garantem consistência e relacionamentos válidos
4. **Criam Índices**: Algumas constraints criam índices automaticamente para performance
5. **Documentam Regras**: Servem como documentação das regras de negócio

### Tipos de Constraints

SQL suporta vários tipos de constraints:

- **PRIMARY KEY**: Identificador único para cada registro
- **FOREIGN KEY**: Relacionamento com outra tabela
- **UNIQUE**: Valores únicos em uma coluna ou conjunto de colunas
- **NOT NULL**: Campo obrigatório (não pode ser NULL)
- **CHECK**: Validação customizada de dados

### Por que Usar Constraints?

**Sem Constraints:**
```sql
-- Problema: Pode inserir dados inválidos
INSERT INTO livros (id, titulo) VALUES (1, NULL);  -- Título nulo?
INSERT INTO livros (id, titulo) VALUES (1, 'Dom Casmurro');  -- ID duplicado?
INSERT INTO emprestimos (livro_id) VALUES (999);  -- Livro que não existe?
```

**Com Constraints:**
```sql
-- Constraints previnem esses problemas automaticamente
-- O banco de dados rejeita dados inválidos
```

---

## 2. PRIMARY KEY (Chave Primária)

A **PRIMARY KEY** (Chave Primária) é um identificador único para cada registro em uma tabela. Ela garante que cada linha possa ser identificada de forma única e que não haja duplicatas.

### Características da PRIMARY KEY

1. **Única**: Não pode haver dois registros com a mesma chave primária
2. **Não Nula**: A PRIMARY KEY não pode conter valores NULL
3. **Imutável**: Idealmente, não deve ser alterada (embora tecnicamente possível)
4. **Cria Índice**: Automaticamente cria um índice para busca rápida
5. **Uma por Tabela**: Cada tabela pode ter apenas uma PRIMARY KEY

### Sintaxe Básica

#### PRIMARY KEY em uma Coluna

```sql
CREATE TABLE livros (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    titulo TEXT NOT NULL,
    ano_publicacao INTEGER
);
```

#### PRIMARY KEY Composta (Múltiplas Colunas)

```sql
CREATE TABLE emprestimos_livros (
    emprestimo_id INTEGER,
    livro_id INTEGER,
    quantidade INTEGER,
    PRIMARY KEY (emprestimo_id, livro_id)
);
```

### AUTOINCREMENT

O `AUTOINCREMENT` faz com que o banco de dados gere automaticamente valores únicos sequenciais para a coluna PRIMARY KEY.

```sql
CREATE TABLE autores (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    nome TEXT NOT NULL
);

-- Ao inserir, não precisa especificar o ID
INSERT INTO autores (nome) VALUES ('Machado de Assis');
-- O banco automaticamente atribui id = 1

INSERT INTO autores (nome) VALUES ('Clarice Lispector');
-- O banco automaticamente atribui id = 2
```

**Importante**: Em SQLite, `AUTOINCREMENT` é opcional. Sem ele, o banco ainda gera IDs sequenciais, mas com comportamento ligeiramente diferente ao reutilizar IDs deletados.

### Exemplos Práticos

#### Criar Tabela com PRIMARY KEY

```sql
CREATE TABLE categorias (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    nome TEXT NOT NULL UNIQUE,
    descricao TEXT
);
```

#### Verificar PRIMARY KEY Existente

```sql
-- Ver estrutura da tabela
.schema livros

-- Ou consultar informações do SQLite
SELECT sql FROM sqlite_master WHERE type='table' AND name='livros';
```

#### Tentar Violar PRIMARY KEY

```sql
-- Inserir primeiro registro
INSERT INTO livros (id, titulo) VALUES (1, 'Dom Casmurro');

-- Tentar inserir com mesmo ID (ERRO!)
INSERT INTO livros (id, titulo) VALUES (1, 'Memórias Póstumas');
-- Erro: UNIQUE constraint failed: livros.id
```

### Modificar PRIMARY KEY

**⚠️ ATENÇÃO**: Modificar PRIMARY KEY é uma operação complexa e raramente necessária. Em SQLite, você precisa recriar a tabela.

```sql
-- Não é possível ALTER TABLE para modificar PRIMARY KEY diretamente
-- Solução: Recriar a tabela
```

---

## 3. FOREIGN KEY (Chave Estrangeira)

A **FOREIGN KEY** (Chave Estrangeira) é uma coluna ou conjunto de colunas em uma tabela que referencia a PRIMARY KEY de outra tabela. Ela estabelece um relacionamento entre duas tabelas e garante integridade referencial.

### Características da FOREIGN KEY

1. **Referencia PRIMARY KEY**: Aponta para a chave primária de outra tabela
2. **Mantém Integridade Referencial**: Garante que valores referenciados existam
3. **Previne Orfãos**: Evita registros "órfãos" (referências a registros inexistentes)
4. **Suporta CASCADE**: Pode deletar/atualizar registros relacionados automaticamente
5. **Cria Índice**: Automaticamente cria um índice para performance

### Sintaxe Básica

```sql
CREATE TABLE emprestimos (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    livro_id INTEGER NOT NULL,
    usuario_id INTEGER NOT NULL,
    data_emprestimo DATE DEFAULT CURRENT_DATE,
    FOREIGN KEY (livro_id) REFERENCES livros(id),
    FOREIGN KEY (usuario_id) REFERENCES usuarios(id)
);
```

### Habilitar FOREIGN KEY no SQLite

**⚠️ IMPORTANTE**: Por padrão, SQLite **desabilita** verificação de FOREIGN KEY por compatibilidade. Você precisa habilitar:

```sql
-- Habilitar verificação de FOREIGN KEY
PRAGMA foreign_keys = ON;

-- Verificar se está habilitado
PRAGMA foreign_keys;
-- Retorna 1 se habilitado, 0 se desabilitado
```

**Nota**: Esta configuração é por conexão. Você precisa executar `PRAGMA foreign_keys = ON;` toda vez que abrir uma nova conexão.

### Comportamento em DELETE e UPDATE

FOREIGN KEY pode ter diferentes comportamentos quando o registro referenciado é deletado ou atualizado:

#### RESTRICT (Padrão)
Bloqueia a operação se houver referências.

```sql
CREATE TABLE emprestimos (
    id INTEGER PRIMARY KEY,
    livro_id INTEGER,
    FOREIGN KEY (livro_id) REFERENCES livros(id) ON DELETE RESTRICT
);

-- Tentar deletar livro que tem empréstimos (ERRO!)
DELETE FROM livros WHERE id = 1;
-- Erro: FOREIGN KEY constraint failed
```

#### CASCADE
Deleta/atualiza registros relacionados automaticamente.

```sql
CREATE TABLE emprestimos (
    id INTEGER PRIMARY KEY,
    livro_id INTEGER,
    FOREIGN KEY (livro_id) REFERENCES livros(id) ON DELETE CASCADE
);

-- Deletar livro também deleta seus empréstimos
DELETE FROM livros WHERE id = 1;
-- Empréstimos relacionados também são deletados
```

#### SET NULL
Define a FOREIGN KEY como NULL quando o registro referenciado é deletado.

```sql
CREATE TABLE livros (
    id INTEGER PRIMARY KEY,
    autor_id INTEGER,
    FOREIGN KEY (autor_id) REFERENCES autores(id) ON DELETE SET NULL
);

-- Deletar autor define autor_id como NULL nos livros
DELETE FROM autores WHERE id = 1;
-- Livros relacionados têm autor_id = NULL
```

#### SET DEFAULT
Define a FOREIGN KEY como valor padrão quando o registro referenciado é deletado.

```sql
CREATE TABLE livros (
    id INTEGER PRIMARY KEY,
    categoria_id INTEGER DEFAULT 1,
    FOREIGN KEY (categoria_id) REFERENCES categorias(id) ON DELETE SET DEFAULT
);
```

### Exemplos Práticos

#### Criar Tabela com FOREIGN KEY

```sql
-- Habilitar FOREIGN KEY
PRAGMA foreign_keys = ON;

CREATE TABLE emprestimos (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    livro_id INTEGER NOT NULL,
    usuario_id INTEGER NOT NULL,
    data_emprestimo DATE DEFAULT CURRENT_DATE,
    FOREIGN KEY (livro_id) REFERENCES livros(id) ON DELETE RESTRICT,
    FOREIGN KEY (usuario_id) REFERENCES usuarios(id) ON DELETE RESTRICT
);
```

#### Tentar Violar FOREIGN KEY

```sql
-- Tentar inserir empréstimo com livro inexistente (ERRO!)
INSERT INTO emprestimos (livro_id, usuario_id) VALUES (999, 1);
-- Erro: FOREIGN KEY constraint failed

-- Tentar deletar livro que tem empréstimos (ERRO com RESTRICT!)
DELETE FROM livros WHERE id = 1;
-- Erro: FOREIGN KEY constraint failed
```

#### Verificar FOREIGN KEY

```sql
-- Ver estrutura da tabela
.schema emprestimos

-- Consultar constraints
SELECT sql FROM sqlite_master 
WHERE type='table' AND name='emprestimos';
```

### Adicionar FOREIGN KEY a Tabela Existente

```sql
-- Adicionar FOREIGN KEY a tabela existente
ALTER TABLE livros 
ADD COLUMN autor_id INTEGER;

-- SQLite não suporta ADD CONSTRAINT diretamente
-- Solução: Recriar a tabela ou usar trigger
```

**Nota**: SQLite tem limitações ao adicionar constraints a tabelas existentes. A melhor prática é definir constraints na criação da tabela.

---

## 4. UNIQUE (Valores Únicos)

A constraint **UNIQUE** garante que todos os valores em uma coluna ou conjunto de colunas sejam distintos (únicos). Diferente da PRIMARY KEY, uma tabela pode ter múltiplas constraints UNIQUE, e colunas UNIQUE podem conter NULL (a menos que também sejam NOT NULL).

### Características da UNIQUE

1. **Valores Únicos**: Não permite valores duplicados
2. **Múltiplas por Tabela**: Uma tabela pode ter várias constraints UNIQUE
3. **Pode Ser NULL**: Colunas UNIQUE podem ter NULL (mas apenas um NULL é permitido em algumas implementações)
4. **Cria Índice**: Automaticamente cria um índice para busca rápida
5. **Não é Identificador**: Diferente de PRIMARY KEY, não identifica necessariamente o registro

### Sintaxe Básica

#### UNIQUE em uma Coluna

```sql
CREATE TABLE usuarios (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    nome TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL,
    telefone TEXT
);
```

#### UNIQUE em Múltiplas Colunas

```sql
CREATE TABLE livros_autores (
    livro_id INTEGER,
    autor_id INTEGER,
    ordem INTEGER,
    UNIQUE (livro_id, autor_id)
);
```

### UNIQUE vs PRIMARY KEY

| Característica | PRIMARY KEY | UNIQUE |
|---------------|-------------|--------|
| Valores únicos | Sim | Sim |
| Não pode ser NULL | Sim | Pode (depende) |
| Quantidade por tabela | Uma | Múltiplas |
| Cria índice | Sim | Sim |
| Identifica registro | Sim | Não necessariamente |

### Exemplos Práticos

#### Criar Tabela com UNIQUE

```sql
CREATE TABLE usuarios (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    nome TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL,
    cpf TEXT UNIQUE
);
```

#### Tentar Violar UNIQUE

```sql
-- Inserir primeiro usuário
INSERT INTO usuarios (nome, email) VALUES ('João', 'joao@email.com');

-- Tentar inserir com mesmo email (ERRO!)
INSERT INTO usuarios (nome, email) VALUES ('Maria', 'joao@email.com');
-- Erro: UNIQUE constraint failed: usuarios.email
```

#### UNIQUE com NULL

```sql
-- Em SQLite, múltiplos NULL são permitidos em colunas UNIQUE
INSERT INTO usuarios (nome, email, cpf) VALUES ('João', 'joao@email.com', NULL);
INSERT INTO usuarios (nome, email, cpf) VALUES ('Maria', 'maria@email.com', NULL);
-- Ambos são permitidos (cpf é NULL em ambos)
```

### Adicionar UNIQUE a Tabela Existente

```sql
-- Criar índice único (equivalente a UNIQUE)
CREATE UNIQUE INDEX idx_usuarios_email ON usuarios(email);
```

---

## 5. NOT NULL (Não Nulo)

A constraint **NOT NULL** garante que uma coluna não possa conter valores NULL. Ela torna o campo obrigatório, exigindo que um valor seja fornecido ao inserir ou atualizar um registro.

### Características da NOT NULL

1. **Campo Obrigatório**: Exige que um valor seja fornecido
2. **Previne NULL**: Bloqueia inserção/atualização com NULL
3. **Múltiplas por Tabela**: Pode ser aplicada a várias colunas
4. **Não Cria Índice**: Não cria índice automaticamente
5. **Simples mas Essencial**: Uma das constraints mais importantes

### Sintaxe Básica

```sql
CREATE TABLE livros (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    titulo TEXT NOT NULL,
    isbn TEXT,
    ano_publicacao INTEGER
);
```

### NULL vs NOT NULL

**NULL** significa "ausência de valor" ou "desconhecido". É diferente de:
- String vazia (`''`)
- Zero (`0`)
- Espaço em branco (`' '`)

```sql
-- NULL: ausência de valor
INSERT INTO livros (titulo, isbn) VALUES ('Dom Casmurro', NULL);

-- String vazia: valor presente (mas vazio)
INSERT INTO livros (titulo, isbn) VALUES ('Dom Casmurro', '');

-- NOT NULL previne NULL, mas permite string vazia
CREATE TABLE livros (
    titulo TEXT NOT NULL  -- Pode ser '', mas não pode ser NULL
);
```

### Exemplos Práticos

#### Criar Tabela com NOT NULL

```sql
CREATE TABLE autores (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    nome TEXT NOT NULL,
    nacionalidade TEXT,  -- Pode ser NULL
    data_nascimento DATE  -- Pode ser NULL
);
```

#### Tentar Violar NOT NULL

```sql
-- Tentar inserir sem valor obrigatório (ERRO!)
INSERT INTO autores (nacionalidade) VALUES ('Brasileiro');
-- Erro: NOT NULL constraint failed: autores.nome

-- Correto: fornecer valor para campo NOT NULL
INSERT INTO autores (nome, nacionalidade) VALUES ('Machado de Assis', 'Brasileiro');
```

#### Combinar NOT NULL com DEFAULT

```sql
CREATE TABLE emprestimos (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    livro_id INTEGER NOT NULL,
    usuario_id INTEGER NOT NULL,
    data_emprestimo DATE NOT NULL DEFAULT CURRENT_DATE,
    status TEXT NOT NULL DEFAULT 'ativo'
);
```

### Adicionar NOT NULL a Tabela Existente

**⚠️ ATENÇÃO**: Adicionar NOT NULL a uma tabela existente pode falhar se houver registros com NULL na coluna.

```sql
-- Primeiro, atualizar valores NULL existentes
UPDATE livros SET titulo = 'Sem Título' WHERE titulo IS NULL;

-- Depois, adicionar NOT NULL (SQLite não suporta diretamente)
-- Solução: Recriar a tabela
```

---

## 6. CHECK (Validação Customizada)

A constraint **CHECK** permite definir uma condição customizada que deve ser verdadeira para cada linha na tabela. Ela valida dados baseado em uma expressão SQL que você define.

### Características da CHECK

1. **Validação Customizada**: Permite definir regras específicas
2. **Expressão SQL**: Usa expressões SQL para validação
3. **Avaliada por Linha**: Verificada para cada INSERT/UPDATE
4. **Não Cria Índice**: Não cria índice automaticamente
5. **Flexível**: Pode validar múltiplas colunas e condições complexas

### Sintaxe Básica

```sql
CREATE TABLE livros (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    titulo TEXT NOT NULL,
    quantidade_disponivel INTEGER CHECK (quantidade_disponivel >= 0),
    ano_publicacao INTEGER CHECK (ano_publicacao > 0 AND ano_publicacao <= 2100)
);
```

### Exemplos de CHECK

#### Validação de Intervalo

```sql
CREATE TABLE livros (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    titulo TEXT NOT NULL,
    quantidade_disponivel INTEGER CHECK (quantidade_disponivel >= 0 AND quantidade_disponivel <= 1000),
    ano_publicacao INTEGER CHECK (ano_publicacao BETWEEN 1000 AND 2100)
);
```

#### Validação de Lista de Valores

```sql
CREATE TABLE emprestimos (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    livro_id INTEGER NOT NULL,
    usuario_id INTEGER NOT NULL,
    status TEXT NOT NULL CHECK (status IN ('ativo', 'devolvido', 'atrasado', 'cancelado'))
);
```

#### Validação de Múltiplas Colunas

```sql
CREATE TABLE emprestimos (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    livro_id INTEGER NOT NULL,
    usuario_id INTEGER NOT NULL,
    data_emprestimo DATE NOT NULL,
    data_devolucao_prevista DATE,
    CHECK (data_devolucao_prevista IS NULL OR data_devolucao_prevista >= data_emprestimo)
);
```

#### Validação de Formato (Básico)

```sql
CREATE TABLE usuarios (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    nome TEXT NOT NULL,
    email TEXT NOT NULL CHECK (email LIKE '%@%.%'),
    telefone TEXT CHECK (telefone IS NULL OR LENGTH(telefone) >= 10)
);
```

### Exemplos Práticos

#### Criar Tabela com CHECK

```sql
CREATE TABLE livros (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    titulo TEXT NOT NULL,
    quantidade_disponivel INTEGER DEFAULT 0 CHECK (quantidade_disponivel >= 0),
    ano_publicacao INTEGER CHECK (ano_publicacao IS NULL OR (ano_publicacao > 0 AND ano_publicacao <= 2100))
);
```

#### Tentar Violar CHECK

```sql
-- Tentar inserir quantidade negativa (ERRO!)
INSERT INTO livros (titulo, quantidade_disponivel) 
VALUES ('Dom Casmurro', -5);
-- Erro: CHECK constraint failed: livros

-- Tentar inserir ano inválido (ERRO!)
INSERT INTO livros (titulo, ano_publicacao) 
VALUES ('Dom Casmurro', 3000);
-- Erro: CHECK constraint failed: livros
```

### Limitações do CHECK

**⚠️ IMPORTANTE**: CHECK tem algumas limitações:

1. **Não pode usar subqueries**: Não pode referenciar outras tabelas
2. **Não pode usar funções não-determinísticas**: Algumas funções não são permitidas
3. **Performance**: Pode impactar performance em INSERT/UPDATE frequentes
4. **Validação Básica**: Para validações complexas, considere usar triggers ou validação na aplicação

### Quando Usar CHECK

**Use CHECK para:**
- Validação de intervalos numéricos
- Validação de listas de valores permitidos
- Validação de formatos básicos
- Regras de negócio simples

**Não use CHECK para:**
- Validações que dependem de outras tabelas
- Validações muito complexas
- Validações que mudam frequentemente
- Validações que precisam de lógica de negócio complexa

---

## 7. Combinando Constraints

É comum combinar múltiplas constraints em uma única tabela para garantir integridade completa dos dados.

### Exemplo Completo

```sql
-- Habilitar FOREIGN KEY
PRAGMA foreign_keys = ON;

CREATE TABLE emprestimos (
    -- PRIMARY KEY
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    
    -- FOREIGN KEY com NOT NULL
    livro_id INTEGER NOT NULL,
    usuario_id INTEGER NOT NULL,
    
    -- NOT NULL com DEFAULT
    data_emprestimo DATE NOT NULL DEFAULT CURRENT_DATE,
    data_devolucao_prevista DATE,
    data_devolucao_real DATE,
    
    -- CHECK com lista de valores
    status TEXT NOT NULL DEFAULT 'ativo' 
        CHECK (status IN ('ativo', 'devolvido', 'atrasado', 'cancelado')),
    
    -- FOREIGN KEY constraints
    FOREIGN KEY (livro_id) REFERENCES livros(id) ON DELETE RESTRICT,
    FOREIGN KEY (usuario_id) REFERENCES usuarios(id) ON DELETE RESTRICT,
    
    -- CHECK com múltiplas colunas
    CHECK (data_devolucao_prevista IS NULL OR data_devolucao_prevista >= data_emprestimo),
    CHECK (data_devolucao_real IS NULL OR data_devolucao_real >= data_emprestimo)
);
```

### Ordem de Verificação

Quando você insere ou atualiza dados, as constraints são verificadas nesta ordem:

1. **NOT NULL**: Verifica se campos obrigatórios têm valores
2. **CHECK**: Verifica condições customizadas
3. **UNIQUE**: Verifica valores únicos
4. **PRIMARY KEY**: Verifica chave primária (também é UNIQUE)
5. **FOREIGN KEY**: Verifica integridade referencial

---

## 8. Modificando Constraints

### Adicionar Constraints a Tabelas Existentes

SQLite tem limitações ao adicionar constraints a tabelas existentes. Algumas opções:

#### Adicionar NOT NULL (com cuidado)

```sql
-- Primeiro, garantir que não há NULLs
UPDATE livros SET titulo = 'Sem Título' WHERE titulo IS NULL;

-- SQLite não suporta ALTER TABLE ... ADD CONSTRAINT NOT NULL diretamente
-- Solução: Recriar a tabela
```

#### Adicionar UNIQUE

```sql
-- Criar índice único (equivalente a UNIQUE)
CREATE UNIQUE INDEX idx_livros_isbn ON livros(isbn);
```

#### Adicionar CHECK

```sql
-- SQLite não suporta ALTER TABLE ... ADD CONSTRAINT CHECK diretamente
-- Solução: Recriar a tabela ou usar triggers
```

#### Adicionar FOREIGN KEY

```sql
-- SQLite não suporta ALTER TABLE ... ADD CONSTRAINT FOREIGN KEY diretamente
-- Solução: Recriar a tabela
```

### Remover Constraints

#### Remover UNIQUE

```sql
-- Remover índice único
DROP INDEX idx_livros_isbn;
```

#### Remover CHECK

```sql
-- SQLite não suporta remover CHECK diretamente
-- Solução: Recriar a tabela sem a constraint
```

#### Remover FOREIGN KEY

```sql
-- SQLite não suporta remover FOREIGN KEY diretamente
-- Solução: Recriar a tabela sem a constraint
```

### Recriar Tabela (Método Completo)

Quando você precisa modificar constraints significativamente, a melhor abordagem é recriar a tabela:

```sql
-- 1. Criar nova tabela com constraints desejadas
CREATE TABLE livros_novo (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    titulo TEXT NOT NULL,
    isbn TEXT UNIQUE,
    quantidade_disponivel INTEGER CHECK (quantidade_disponivel >= 0)
);

-- 2. Copiar dados
INSERT INTO livros_novo SELECT * FROM livros;

-- 3. Deletar tabela antiga
DROP TABLE livros;

-- 4. Renomear nova tabela
ALTER TABLE livros_novo RENAME TO livros;
```

---

## 9. Verificando Constraints

### Ver Estrutura de Tabela

```sql
-- Ver schema completo
.schema livros

-- Ou consultar sqlite_master
SELECT sql FROM sqlite_master 
WHERE type='table' AND name='livros';
```

### Ver Constraints Específicas

```sql
-- Ver índices (incluindo PRIMARY KEY, UNIQUE, FOREIGN KEY)
SELECT * FROM sqlite_master 
WHERE type='index' AND tbl_name='livros';

-- Ver todas as tabelas e suas constraints
SELECT name, sql FROM sqlite_master 
WHERE type='table';
```

### Verificar Integridade Referencial

```sql
-- Verificar se há registros órfãos
SELECT e.* 
FROM emprestimos e
LEFT JOIN livros l ON e.livro_id = l.id
WHERE l.id IS NULL;
-- Se retornar linhas, há registros órfãos (violação de integridade)
```

### Testar Constraints

```sql
-- Testar PRIMARY KEY
INSERT INTO livros (id, titulo) VALUES (1, 'Teste');
INSERT INTO livros (id, titulo) VALUES (1, 'Teste 2');  -- Deve falhar

-- Testar UNIQUE
INSERT INTO usuarios (nome, email) VALUES ('João', 'joao@email.com');
INSERT INTO usuarios (nome, email) VALUES ('Maria', 'joao@email.com');  -- Deve falhar

-- Testar NOT NULL
INSERT INTO livros (isbn) VALUES ('123');  -- Deve falhar se titulo é NOT NULL

-- Testar CHECK
INSERT INTO livros (titulo, quantidade_disponivel) 
VALUES ('Teste', -1);  -- Deve falhar se CHECK >= 0

-- Testar FOREIGN KEY
PRAGMA foreign_keys = ON;
INSERT INTO emprestimos (livro_id, usuario_id) 
VALUES (999, 1);  -- Deve falhar se livro 999 não existe
```

---

## 10. Erros Comuns e Soluções

### Erro: UNIQUE constraint failed

**Causa**: Tentativa de inserir valor duplicado em coluna UNIQUE ou PRIMARY KEY.

**Solução**:
```sql
-- Verificar valores existentes
SELECT email, COUNT(*) FROM usuarios GROUP BY email HAVING COUNT(*) > 1;

-- Usar INSERT OR IGNORE ou INSERT OR REPLACE
INSERT OR IGNORE INTO usuarios (nome, email) VALUES ('João', 'joao@email.com');
```

### Erro: NOT NULL constraint failed

**Causa**: Tentativa de inserir NULL em coluna NOT NULL.

**Solução**:
```sql
-- Fornecer valor para campo obrigatório
INSERT INTO livros (titulo, isbn) VALUES ('Dom Casmurro', '123');
-- Ou usar DEFAULT
INSERT INTO livros (titulo) VALUES ('Dom Casmurro');
```

### Erro: FOREIGN KEY constraint failed

**Causa**: Tentativa de inserir/deletar registro que viola integridade referencial.

**Solução**:
```sql
-- Verificar se registro referenciado existe
SELECT * FROM livros WHERE id = 1;

-- Ou usar CASCADE para deletar automaticamente
-- Ou deletar registros dependentes primeiro
```

### Erro: CHECK constraint failed

**Causa**: Valor não atende à condição CHECK.

**Solução**:
```sql
-- Verificar a constraint
.schema tabela

-- Corrigir o valor para atender à condição
INSERT INTO livros (titulo, quantidade_disponivel) 
VALUES ('Dom Casmurro', 0);  -- Usar valor válido
```

### FOREIGN KEY não está funcionando

**Causa**: FOREIGN KEY está desabilitado no SQLite.

**Solução**:
```sql
-- Habilitar FOREIGN KEY
PRAGMA foreign_keys = ON;

-- Verificar
PRAGMA foreign_keys;  -- Deve retornar 1
```

---

## 11. Boas Práticas

### 1. Sempre Use PRIMARY KEY

Toda tabela deve ter uma PRIMARY KEY para identificar registros únicos.

```sql
-- ✅ BOM
CREATE TABLE livros (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    titulo TEXT NOT NULL
);

-- ❌ EVITE
CREATE TABLE livros (
    titulo TEXT NOT NULL  -- Sem PRIMARY KEY!
);
```

### 2. Use FOREIGN KEY para Relacionamentos

Sempre defina FOREIGN KEY para manter integridade referencial.

```sql
-- ✅ BOM
CREATE TABLE emprestimos (
    id INTEGER PRIMARY KEY,
    livro_id INTEGER,
    FOREIGN KEY (livro_id) REFERENCES livros(id)
);

-- ❌ EVITE
CREATE TABLE emprestimos (
    id INTEGER PRIMARY KEY,
    livro_id INTEGER  -- Sem FOREIGN KEY!
);
```

### 3. Use NOT NULL para Campos Essenciais

Aplique NOT NULL a campos que são sempre necessários.

```sql
-- ✅ BOM
CREATE TABLE usuarios (
    id INTEGER PRIMARY KEY,
    nome TEXT NOT NULL,
    email TEXT NOT NULL,
    telefone TEXT  -- Pode ser NULL
);
```

### 4. Use UNIQUE para Valores que Devem Ser Únicos

Aplique UNIQUE a campos como email, CPF, ISBN que devem ser únicos.

```sql
-- ✅ BOM
CREATE TABLE usuarios (
    id INTEGER PRIMARY KEY,
    email TEXT UNIQUE NOT NULL
);
```

### 5. Use CHECK para Validações Simples

Use CHECK para validações que não dependem de outras tabelas.

```sql
-- ✅ BOM
CREATE TABLE livros (
    id INTEGER PRIMARY KEY,
    quantidade_disponivel INTEGER CHECK (quantidade_disponivel >= 0)
);
```

### 6. Defina Constraints na Criação

É mais fácil definir constraints na criação da tabela do que adicionar depois.

```sql
-- ✅ BOM: Definir tudo na criação
CREATE TABLE livros (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    titulo TEXT NOT NULL,
    isbn TEXT UNIQUE,
    quantidade_disponivel INTEGER CHECK (quantidade_disponivel >= 0)
);
```

### 7. Documente Constraints Complexas

Comente constraints complexas para facilitar manutenção.

```sql
CREATE TABLE emprestimos (
    id INTEGER PRIMARY KEY,
    -- Garante que data de devolução seja após empréstimo
    CHECK (data_devolucao_prevista IS NULL OR 
           data_devolucao_prevista >= data_emprestimo)
);
```

### 8. Teste Constraints

Sempre teste suas constraints para garantir que funcionam corretamente.

```sql
-- Testar cada constraint
-- Tentar violar e verificar se é bloqueado
```

---

## 12. Resumo

### Constraints Fundamentais

| Constraint | Propósito | Características |
|------------|-----------|-----------------|
| **PRIMARY KEY** | Identificador único | Única, não nula, cria índice |
| **FOREIGN KEY** | Relacionamento | Mantém integridade referencial |
| **UNIQUE** | Valores únicos | Múltiplas por tabela, cria índice |
| **NOT NULL** | Campo obrigatório | Previne NULL |
| **CHECK** | Validação customizada | Expressão SQL personalizada |

### Ordem de Aplicação

1. Defina PRIMARY KEY para cada tabela
2. Defina FOREIGN KEY para relacionamentos
3. Adicione NOT NULL para campos essenciais
4. Adicione UNIQUE para valores que devem ser únicos
5. Adicione CHECK para validações customizadas

### Comandos Importantes

```sql
-- Habilitar FOREIGN KEY (SQLite)
PRAGMA foreign_keys = ON;

-- Ver estrutura
.schema tabela

-- Ver constraints
SELECT sql FROM sqlite_master WHERE type='table' AND name='tabela';

-- Testar constraint
INSERT INTO tabela (...) VALUES (...);  -- Tentar violar
```

---

**Próximo Passo**: Leia a **Aula Simplificada** (`02-aula-simplificada.md`) para ver esses conceitos explicados com analogias do dia a dia!
