# Aula 5 - Performance, Boas Práticas e Otimização: Data Constraints

## Introdução

Constraints são fundamentais para integridade de dados, mas também têm impacto na performance e no design do banco de dados. Nesta seção, você aprenderá como usar constraints de forma eficiente, quando cada tipo é apropriado, e como balancear segurança com performance.

---

## 1. Impacto de Constraints na Performance

### Por que Constraints Afetam Performance?

Constraints verificam dados a cada INSERT, UPDATE e DELETE. Isso significa:

1. **Verificações Adicionais**: Cada constraint adiciona uma verificação
2. **Índices Automáticos**: PRIMARY KEY, UNIQUE e FOREIGN KEY criam índices
3. **Overhead de Validação**: CHECK constraints executam expressões SQL
4. **Integridade Referencial**: FOREIGN KEY verifica outras tabelas

### Impacto por Tipo de Constraint

#### PRIMARY KEY e UNIQUE

**Impacto**: **Positivo** (criam índices que aceleram queries)

```sql
-- PRIMARY KEY cria índice automaticamente
CREATE TABLE livros (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    titulo TEXT NOT NULL
);
-- Índice em 'id' é criado automaticamente

-- Queries são rápidas
SELECT * FROM livros WHERE id = 1;  -- Usa índice, muito rápido
```

**Benefícios**:
- Busca por PRIMARY KEY é extremamente rápida (O(log n))
- JOINs são otimizados automaticamente
- Ordenação por PRIMARY KEY é rápida

**Custo**:
- INSERT/UPDATE podem ser ligeiramente mais lentos (precisa atualizar índice)
- Mais espaço em disco (índices ocupam espaço)

#### FOREIGN KEY

**Impacto**: **Neutro a Negativo** (verifica outras tabelas)

```sql
PRAGMA foreign_keys = ON;

CREATE TABLE emprestimos (
    id INTEGER PRIMARY KEY,
    livro_id INTEGER,
    FOREIGN KEY (livro_id) REFERENCES livros(id)
);
```

**Benefícios**:
- Cria índice automaticamente (acelera JOINs)
- Garante integridade referencial

**Custo**:
- INSERT/UPDATE verificam tabela referenciada (pode ser lento se tabela for grande)
- DELETE pode ser bloqueado ou executar CASCADE (pode ser lento)

**Otimização**:
```sql
-- Índice na FOREIGN KEY já existe automaticamente
-- Mas você pode criar índices adicionais se necessário
CREATE INDEX idx_emprestimos_livro ON emprestimos(livro_id);
```

#### NOT NULL

**Impacto**: **Mínimo** (verificação muito simples)

```sql
CREATE TABLE livros (
    id INTEGER PRIMARY KEY,
    titulo TEXT NOT NULL  -- Verificação muito rápida
);
```

**Benefícios**:
- Verificação extremamente rápida (apenas verifica se é NULL)
- Previne problemas de dados incompletos

**Custo**:
- Praticamente zero (verificação trivial)

#### CHECK

**Impacto**: **Variável** (depende da complexidade da expressão)

```sql
-- CHECK simples (rápido)
CREATE TABLE livros (
    quantidade_disponivel INTEGER CHECK (quantidade_disponivel >= 0)
);

-- CHECK complexo (mais lento)
CREATE TABLE emprestimos (
    data_emprestimo DATE,
    data_devolucao DATE,
    CHECK (data_devolucao >= data_emprestimo AND 
           data_devolucao <= DATE('now', '+30 days'))
);
```

**Benefícios**:
- Valida dados antes de inserir
- Previne dados inválidos

**Custo**:
- Expressões simples: impacto mínimo
- Expressões complexas: podem ser lentas em INSERT/UPDATE frequentes
- Subqueries não são permitidas (limitação do SQLite)

### Exemplo de Impacto

```sql
-- Sem constraints (mais rápido para INSERT, mas dados podem ser inválidos)
INSERT INTO livros (titulo, quantidade_disponivel) 
VALUES ('Dom Casmurro', -5);  -- Aceita (mas é inválido!)

-- Com constraints (mais lento para INSERT, mas dados sempre válidos)
INSERT INTO livros (titulo, quantidade_disponivel) 
VALUES ('Dom Casmurro', -5);  -- Erro: CHECK constraint failed
```

**Trade-off**: Performance vs Integridade

---

## 2. Índices Criados Automaticamente

### Constraints que Criam Índices

Algumas constraints criam índices automaticamente:

1. **PRIMARY KEY**: Cria índice único automaticamente
2. **UNIQUE**: Cria índice único automaticamente
3. **FOREIGN KEY**: Cria índice (em algumas implementações)

### Verificando Índices

```sql
-- Ver todos os índices
SELECT * FROM sqlite_master WHERE type='index';

-- Ver índices de uma tabela específica
SELECT * FROM sqlite_master 
WHERE type='index' AND tbl_name='livros';
```

### Quando Criar Índices Adicionais

Mesmo com índices automáticos, você pode precisar de índices adicionais:

```sql
-- FOREIGN KEY cria índice, mas você pode querer índice composto
CREATE TABLE emprestimos (
    id INTEGER PRIMARY KEY,
    livro_id INTEGER,
    usuario_id INTEGER,
    data_emprestimo DATE,
    FOREIGN KEY (livro_id) REFERENCES livros(id),
    FOREIGN KEY (usuario_id) REFERENCES usuarios(id)
);

-- Índice composto para query comum
CREATE INDEX idx_emprestimos_usuario_data 
ON emprestimos(usuario_id, data_emprestimo);

-- Query que se beneficia:
SELECT * FROM emprestimos 
WHERE usuario_id = 1 AND data_emprestimo > '2024-01-01';
```

---

## 3. Estratégias de Validação: Banco vs Aplicação

### Quando Usar Constraints no Banco

**Use constraints no banco para**:

1. **Regras Fundamentais**: Regras que nunca mudam
   ```sql
   -- Quantidade não pode ser negativa (sempre verdadeiro)
   quantidade_disponivel INTEGER CHECK (quantidade_disponivel >= 0)
   ```

2. **Integridade Referencial**: Relacionamentos entre tabelas
   ```sql
   -- Sempre deve referenciar livro existente
   FOREIGN KEY (livro_id) REFERENCES livros(id)
   ```

3. **Identificadores Únicos**: PRIMARY KEY, UNIQUE
   ```sql
   -- Email sempre deve ser único
   email TEXT UNIQUE NOT NULL
   ```

4. **Campos Obrigatórios**: NOT NULL
   ```sql
   -- Título sempre é obrigatório
   titulo TEXT NOT NULL
   ```

### Quando Usar Validação na Aplicação

**Use validação na aplicação para**:

1. **Regras que Mudam Frequentemente**
   ```sql
   -- ❌ EVITE no banco: Regra que pode mudar
   -- CHECK (desconto <= 50)  -- E se mudar para 60?
   
   -- ✅ MELHOR na aplicação: Validação flexível
   ```

2. **Validações Complexas que Dependem de Outras Tabelas**
   ```sql
   -- ❌ EVITE no banco: CHECK não pode usar subqueries
   -- CHECK (quantidade <= (SELECT estoque FROM produtos WHERE id = produto_id))
   
   -- ✅ MELHOR na aplicação: Validação complexa
   ```

3. **Validações que Precisam de Contexto da Aplicação**
   ```sql
   -- ❌ EVITE no banco: Precisa de contexto
   -- CHECK (usuario_pode_emprestar(usuario_id))
   
   -- ✅ MELHOR na aplicação: Lógica de negócio
   ```

4. **Mensagens de Erro Personalizadas**
   ```sql
   -- ❌ EVITE no banco: Mensagem genérica
   -- Erro: CHECK constraint failed
   
   -- ✅ MELHOR na aplicação: Mensagem personalizada
   -- "Quantidade não pode ser negativa. Você tentou inserir -5."
   ```

### Abordagem Híbrida (Recomendada)

**Melhor prática**: Use ambos!

```sql
-- Banco: Regras fundamentais e integridade
CREATE TABLE livros (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    titulo TEXT NOT NULL,  -- Banco: Obrigatório
    quantidade_disponivel INTEGER CHECK (quantidade_disponivel >= 0),  -- Banco: Regra fundamental
    isbn TEXT UNIQUE  -- Banco: Único
);

-- Aplicação: Validações complexas e regras de negócio
-- - Validar formato de ISBN
-- - Verificar se ISBN já existe em outro sistema
-- - Validar quantidade baseado em regras de negócio complexas
```

**Vantagens**:
- Banco garante integridade básica (mesmo se aplicação tiver bugs)
- Aplicação fornece validações complexas e mensagens claras
- Melhor experiência do usuário
- Mais flexível para mudanças

---

## 4. Boas Práticas de Design de Constraints

### 1. Sempre Use PRIMARY KEY

**Regra de ouro**: Toda tabela deve ter uma PRIMARY KEY.

```sql
-- ✅ SEMPRE
CREATE TABLE livros (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    titulo TEXT NOT NULL
);

-- ❌ NUNCA
CREATE TABLE livros (
    titulo TEXT NOT NULL  -- Sem PRIMARY KEY!
);
```

**Por quê?**
- Identifica registros únicos
- Facilita JOINs e relacionamentos
- Melhora performance com índice automático
- Essencial para integridade referencial

### 2. Use FOREIGN KEY para Relacionamentos

**Sempre defina FOREIGN KEY para relacionamentos**.

```sql
-- ✅ SEMPRE
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

**Por quê?**
- Mantém integridade referencial
- Previne dados órfãos
- Documenta relacionamentos
- Cria índice automaticamente

### 3. Use NOT NULL para Campos Essenciais

**Aplique NOT NULL a campos que são sempre necessários**.

```sql
-- ✅ BOM
CREATE TABLE usuarios (
    id INTEGER PRIMARY KEY,
    nome TEXT NOT NULL,  -- Sempre necessário
    email TEXT NOT NULL,  -- Sempre necessário
    telefone TEXT  -- Pode ser opcional
);

-- ❌ EVITE
CREATE TABLE usuarios (
    id INTEGER PRIMARY KEY,
    nome TEXT,  -- Pode ser NULL? Problema!
    email TEXT  -- Pode ser NULL? Problema!
);
```

**Critérios para NOT NULL**:
- Campo é sempre necessário para o registro fazer sentido?
- Campo é usado em queries importantes?
- Campo é referenciado em relacionamentos?

### 4. Use UNIQUE para Valores que Devem Ser Únicos

**Aplique UNIQUE a campos que devem ser únicos mas não são identificadores**.

```sql
-- ✅ BOM
CREATE TABLE usuarios (
    id INTEGER PRIMARY KEY,
    email TEXT UNIQUE NOT NULL,  -- Deve ser único
    cpf TEXT UNIQUE  -- Deve ser único
);

-- ❌ EVITE
CREATE TABLE usuarios (
    id INTEGER PRIMARY KEY,
    email TEXT NOT NULL  -- Pode ter duplicatas? Problema!
);
```

**Quando usar UNIQUE**:
- Email, CPF, matrícula, código de produto
- Qualquer campo que deve ser único por regra de negócio

### 5. Use CHECK para Validações Simples

**Use CHECK apenas para validações simples e estáveis**.

```sql
-- ✅ BOM: Validação simples e estável
CREATE TABLE livros (
    quantidade_disponivel INTEGER CHECK (quantidade_disponivel >= 0),
    ano_publicacao INTEGER CHECK (ano_publicacao > 0 AND ano_publicacao <= 2100)
);

-- ❌ EVITE: Validação complexa ou que pode mudar
CREATE TABLE livros (
    desconto INTEGER CHECK (desconto <= 50)  -- E se mudar para 60?
);
```

**Critérios para CHECK**:
- Validação é simples (não depende de outras tabelas)?
- Validação é estável (não muda frequentemente)?
- Validação é fundamental (sempre verdadeira)?

### 6. Defina Constraints na Criação

**Sempre defina constraints na criação da tabela**.

```sql
-- ✅ BOM: Tudo definido na criação
CREATE TABLE livros (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    titulo TEXT NOT NULL,
    isbn TEXT UNIQUE,
    quantidade_disponivel INTEGER CHECK (quantidade_disponivel >= 0)
);

-- ❌ EVITE: Adicionar depois (mais difícil)
CREATE TABLE livros (
    id INTEGER PRIMARY KEY,
    titulo TEXT,
    isbn TEXT
);
-- Depois precisa recriar tabela para adicionar constraints
```

**Por quê?**
- Mais fácil e direto
- Evita problemas com dados existentes
- Garante integridade desde o início

### 7. Documente Constraints Complexas

**Comente constraints complexas para facilitar manutenção**.

```sql
-- ✅ BOM: Documentado
CREATE TABLE emprestimos (
    id INTEGER PRIMARY KEY,
    data_emprestimo DATE NOT NULL,
    data_devolucao_prevista DATE NOT NULL,
    -- Garante que data de devolução seja após empréstimo
    CHECK (data_devolucao_prevista >= data_emprestimo)
);
```

---

## 5. Performance: Otimizando Constraints

### 1. Índices Automáticos

**Aproveite índices criados automaticamente**:

```sql
-- PRIMARY KEY cria índice automaticamente
SELECT * FROM livros WHERE id = 1;  -- Muito rápido (usa índice)

-- UNIQUE cria índice automaticamente
SELECT * FROM usuarios WHERE email = 'joao@email.com';  -- Muito rápido (usa índice)
```

### 2. Índices Compostos para FOREIGN KEY

**Crie índices compostos quando necessário**:

```sql
-- FOREIGN KEY cria índice em livro_id
CREATE TABLE emprestimos (
    id INTEGER PRIMARY KEY,
    livro_id INTEGER,
    usuario_id INTEGER,
    data_emprestimo DATE,
    FOREIGN KEY (livro_id) REFERENCES livros(id)
);

-- Mas se você busca por usuario_id + data_emprestimo frequentemente:
CREATE INDEX idx_emprestimos_usuario_data 
ON emprestimos(usuario_id, data_emprestimo);
```

### 3. Otimizar CHECK Constraints

**Mantenha CHECK constraints simples**:

```sql
-- ✅ RÁPIDO: Expressão simples
CREATE TABLE livros (
    quantidade_disponivel INTEGER CHECK (quantidade_disponivel >= 0)
);

-- ❌ LENTO: Expressão complexa
CREATE TABLE livros (
    quantidade_disponivel INTEGER CHECK (
        quantidade_disponivel >= 0 AND 
        quantidade_disponivel <= 1000 AND
        quantidade_disponivel % 1 = 0 AND
        quantidade_disponivel IS NOT NULL
    )
);
```

### 4. Evitar CHECK Desnecessários

**Não use CHECK para validações que já são garantidas por outras constraints**:

```sql
-- ❌ REDUNDANTE: NOT NULL já garante que não é NULL
CREATE TABLE livros (
    titulo TEXT NOT NULL CHECK (titulo IS NOT NULL)
);

-- ✅ CORRETO: Apenas NOT NULL
CREATE TABLE livros (
    titulo TEXT NOT NULL
);
```

---

## 6. Segurança e Integridade

### 1. Constraints Previnem Dados Inválidos

**Constraints protegem mesmo se a aplicação tiver bugs**:

```sql
-- Aplicação pode ter bug e tentar inserir dados inválidos
-- Mas constraints no banco previnem:

INSERT INTO livros (titulo, quantidade_disponivel) 
VALUES (NULL, -5);
-- Erro: NOT NULL constraint failed (título)
-- Erro: CHECK constraint failed (quantidade)
```

### 2. FOREIGN KEY Previne Dados Órfãos

**FOREIGN KEY garante que referências sempre sejam válidas**:

```sql
-- Sem FOREIGN KEY: Pode criar dados órfãos
INSERT INTO emprestimos (livro_id) VALUES (999);  -- Livro 999 não existe!

-- Com FOREIGN KEY: Previne dados órfãos
INSERT INTO emprestimos (livro_id) VALUES (999);
-- Erro: FOREIGN KEY constraint failed
```

### 3. Constraints como Documentação

**Constraints documentam regras de negócio**:

```sql
-- Constraints mostram claramente as regras:
CREATE TABLE usuarios (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    nome TEXT NOT NULL,  -- Nome é obrigatório
    email TEXT UNIQUE NOT NULL,  -- Email é único e obrigatório
    idade INTEGER CHECK (idade >= 18 AND idade <= 100)  -- Idade entre 18-100
);
-- Qualquer desenvolvedor vê imediatamente as regras!
```

---

## 7. Monitoramento e Troubleshooting

### 1. Verificar Constraints

```sql
-- Ver estrutura de tabela (mostra constraints)
.schema livros

-- Ver todas as constraints
SELECT sql FROM sqlite_master WHERE type='table';

-- Ver índices (incluindo os criados por constraints)
SELECT * FROM sqlite_master WHERE type='index';
```

### 2. Testar Constraints

```sql
-- Testar cada constraint tentando violá-la
-- PRIMARY KEY
INSERT INTO livros (id, titulo) VALUES (1, 'Teste');
INSERT INTO livros (id, titulo) VALUES (1, 'Teste 2');  -- Deve falhar

-- UNIQUE
INSERT INTO usuarios (nome, email) VALUES ('João', 'joao@email.com');
INSERT INTO usuarios (nome, email) VALUES ('Maria', 'joao@email.com');  -- Deve falhar

-- NOT NULL
INSERT INTO livros (isbn) VALUES ('123');  -- Deve falhar se titulo é NOT NULL

-- CHECK
INSERT INTO livros (titulo, quantidade_disponivel) 
VALUES ('Teste', -1);  -- Deve falhar

-- FOREIGN KEY
PRAGMA foreign_keys = ON;
INSERT INTO emprestimos (livro_id, usuario_id) 
VALUES (999, 1);  -- Deve falhar se livro 999 não existe
```

### 3. Identificar Problemas de Performance

```sql
-- Usar EXPLAIN QUERY PLAN para ver se índices estão sendo usados
EXPLAIN QUERY PLAN
SELECT * FROM livros WHERE id = 1;
-- Deve mostrar uso de índice (se PRIMARY KEY)

-- Verificar se FOREIGN KEY está habilitado
PRAGMA foreign_keys;
-- Deve retornar 1 se habilitado
```

### 4. Resolver Problemas Comuns

**Problema**: FOREIGN KEY não está funcionando

```sql
-- Solução: Habilitar FOREIGN KEY
PRAGMA foreign_keys = ON;
```

**Problema**: Não consigo deletar registro (FOREIGN KEY RESTRICT)

```sql
-- Opção 1: Deletar dependentes primeiro
DELETE FROM emprestimos WHERE livro_id = 1;
DELETE FROM livros WHERE id = 1;

-- Opção 2: Usar CASCADE (se apropriado)
-- Precisa recriar tabela com ON DELETE CASCADE
```

**Problema**: CHECK constraint muito lenta

```sql
-- Solução: Simplificar expressão CHECK
-- Ou mover validação para aplicação
```

---

## 8. Trade-offs e Decisões de Design

### 1. Performance vs Integridade

**Trade-off**: Constraints garantem integridade mas podem impactar performance.

```sql
-- Mais constraints = Mais segurança, mas pode ser mais lento
CREATE TABLE livros (
    id INTEGER PRIMARY KEY,
    titulo TEXT NOT NULL,
    isbn TEXT UNIQUE,
    quantidade_disponivel INTEGER CHECK (quantidade_disponivel >= 0),
    ano_publicacao INTEGER CHECK (ano_publicacao > 0 AND ano_publicacao <= 2100),
    editora TEXT NOT NULL,
    -- Muitas constraints = Mais verificações = Mais lento em INSERT/UPDATE
);
```

**Recomendação**: Use constraints essenciais, validações complexas na aplicação.

### 2. Flexibilidade vs Segurança

**Trade-off**: Mais constraints = Mais segurança, mas menos flexibilidade.

```sql
-- Muitas constraints = Menos flexível para mudanças
CREATE TABLE livros (
    desconto INTEGER CHECK (desconto <= 50)  -- E se mudar para 60?
);

-- Menos constraints = Mais flexível, mas menos seguro
CREATE TABLE livros (
    desconto INTEGER  -- Pode ser qualquer valor (inclusive inválido)
);
```

**Recomendação**: Use constraints para regras fundamentais e estáveis.

### 3. Banco vs Aplicação

**Trade-off**: Onde colocar validações?

**Banco (Constraints)**:
- ✅ Regras fundamentais
- ✅ Integridade referencial
- ✅ Proteção mesmo com bugs na aplicação
- ❌ Menos flexível
- ❌ Mensagens de erro genéricas

**Aplicação**:
- ✅ Validações complexas
- ✅ Mensagens personalizadas
- ✅ Mais flexível
- ❌ Pode ter bugs
- ❌ Não protege se aplicação for contornada

**Recomendação**: Use ambos! Banco para fundamentos, aplicação para complexidade.

---

## 9. Resumo de Boas Práticas

### Checklist de Constraints

Ao criar uma tabela, verifique:

- [ ] Tem PRIMARY KEY?
- [ ] FOREIGN KEY definidos para relacionamentos?
- [ ] NOT NULL em campos essenciais?
- [ ] UNIQUE em campos que devem ser únicos?
- [ ] CHECK apenas para validações simples e estáveis?
- [ ] Constraints documentadas (se complexas)?
- [ ] Índices adicionais criados quando necessário?

### Regras de Ouro

1. **Sempre use PRIMARY KEY** em toda tabela
2. **Sempre use FOREIGN KEY** para relacionamentos
3. **Use NOT NULL** para campos essenciais
4. **Use UNIQUE** para valores que devem ser únicos
5. **Use CHECK** apenas para validações simples
6. **Defina constraints na criação** da tabela
7. **Documente constraints complexas**
8. **Teste todas as constraints** antes de produção
9. **Use validação na aplicação** para complexidade
10. **Monitore performance** e ajuste quando necessário

---

## 10. Exemplos Práticos

### Exemplo 1: Tabela Bem Projetada

```sql
PRAGMA foreign_keys = ON;

CREATE TABLE emprestimos (
    -- PRIMARY KEY: Identificador único
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    
    -- FOREIGN KEY: Relacionamentos
    livro_id INTEGER NOT NULL,
    usuario_id INTEGER NOT NULL,
    
    -- NOT NULL: Campos obrigatórios
    data_emprestimo DATE NOT NULL DEFAULT CURRENT_DATE,
    data_devolucao_prevista DATE NOT NULL,
    data_devolucao_real DATE,
    
    -- CHECK: Validações simples
    status TEXT NOT NULL DEFAULT 'ativo' 
        CHECK (status IN ('ativo', 'devolvido', 'atrasado')),
    
    -- FOREIGN KEY constraints
    FOREIGN KEY (livro_id) REFERENCES livros(id) ON DELETE RESTRICT,
    FOREIGN KEY (usuario_id) REFERENCES usuarios(id) ON DELETE RESTRICT,
    
    -- CHECK: Validação de datas
    CHECK (data_devolucao_prevista >= data_emprestimo),
    CHECK (data_devolucao_real IS NULL OR data_devolucao_real >= data_emprestimo)
);

-- Índice adicional para query comum
CREATE INDEX idx_emprestimos_usuario_data 
ON emprestimos(usuario_id, data_emprestimo);
```

### Exemplo 2: Otimização de Performance

```sql
-- Tabela com muitas constraints (pode ser lenta)
CREATE TABLE livros (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    titulo TEXT NOT NULL,
    isbn TEXT UNIQUE,
    quantidade_disponivel INTEGER CHECK (quantidade_disponivel >= 0),
    ano_publicacao INTEGER CHECK (ano_publicacao > 0 AND ano_publicacao <= 2100),
    preco REAL CHECK (preco > 0),
    desconto REAL CHECK (desconto >= 0 AND desconto <= 100),
    -- Muitas verificações CHECK podem ser lentas
);

-- Otimização: Mover algumas validações para aplicação
CREATE TABLE livros (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    titulo TEXT NOT NULL,
    isbn TEXT UNIQUE,
    quantidade_disponivel INTEGER CHECK (quantidade_disponivel >= 0),  -- Fundamental
    ano_publicacao INTEGER,  -- Validação na aplicação
    preco REAL CHECK (preco > 0),  -- Fundamental
    desconto REAL  -- Validação na aplicação (pode mudar)
);
```

---

**Próximo Passo**: Revise todos os conceitos e pratique criando tabelas com constraints apropriadas. Lembre-se: constraints são fundamentais para integridade de dados, mas use com sabedoria!
