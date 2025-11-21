# Aula 5 - Simplificada: Entendendo Data Constraints

## Introdução

Imagine que você está organizando uma biblioteca física. Você precisa de regras para garantir que:
- Cada livro tenha um número único de identificação
- Não haja livros duplicados com o mesmo número
- Cada empréstimo esteja vinculado a um livro que realmente existe
- Campos importantes (como título) sempre tenham valor
- Quantidades não sejam negativas

Essas regras são exatamente o que **Data Constraints** (Restrições de Dados) fazem no banco de dados: elas são como "leis" que garantem que os dados sempre estejam corretos e organizados.

**Constraints são como porteiros digitais**: eles verificam cada dado que entra no banco e só deixam passar o que está correto!

---

## 1. Data Constraints: A Analogia da Biblioteca Física

### Pensando em Constraints como Regras da Biblioteca

Imagine uma biblioteca física com regras rígidas:

**Regra 1**: Cada livro precisa ter um número único (não pode ter dois livros com o mesmo número)
**Regra 2**: Só pode emprestar livros que existem na biblioteca
**Regra 3**: Todo livro precisa ter um título (não pode ser vazio)
**Regra 4**: Não pode ter dois usuários com o mesmo email
**Regra 5**: Quantidade de livros não pode ser negativa

Essas regras são exatamente o que constraints fazem no banco de dados!

### Por que Precisamos de Constraints?

**Sem Constraints:**
```
Bibliotecário: "Vou cadastrar um livro sem título..."
Sistema: "Ok, cadastrado!" ❌

Bibliotecário: "Vou emprestar o livro número 999..."
Sistema: "Ok, emprestado!" ❌ (mas o livro 999 não existe!)

Bibliotecário: "Vou cadastrar -5 livros em estoque..."
Sistema: "Ok, cadastrado!" ❌ (quantidade negativa!)
```

**Com Constraints:**
```
Bibliotecário: "Vou cadastrar um livro sem título..."
Sistema: "ERRO! Título é obrigatório!" ✅

Bibliotecário: "Vou emprestar o livro número 999..."
Sistema: "ERRO! Livro 999 não existe!" ✅

Bibliotecário: "Vou cadastrar -5 livros em estoque..."
Sistema: "ERRO! Quantidade não pode ser negativa!" ✅
```

Constraints são como **porteiros inteligentes** que verificam tudo antes de deixar entrar!

---

## 2. PRIMARY KEY: O RG do Registro

### Analogia: RG (Registro Geral)

Pense na PRIMARY KEY como o **RG de cada registro**. Assim como cada pessoa tem um RG único, cada registro na tabela tem uma PRIMARY KEY única.

**Características do RG:**
- Cada pessoa tem um número único
- Não pode haver duas pessoas com o mesmo RG
- O RG identifica uma pessoa específica
- Você sempre precisa ter um RG

**Características da PRIMARY KEY:**
- Cada registro tem um valor único
- Não pode haver dois registros com a mesma PRIMARY KEY
- A PRIMARY KEY identifica um registro específico
- Todo registro precisa ter uma PRIMARY KEY

### Exemplo Prático

```sql
-- Criar tabela de livros com PRIMARY KEY
CREATE TABLE livros (
    id INTEGER PRIMARY KEY AUTOINCREMENT,  -- O "RG" do livro
    titulo TEXT NOT NULL
);
```

**O que acontece:**
```
Livro 1: id = 1, título = "Dom Casmurro"
Livro 2: id = 2, título = "1984"
Livro 3: id = 3, título = "Fundação"

-- Tentar criar outro livro com id = 1 (ERRO!)
Livro 4: id = 1, título = "Memórias Póstumas"
Sistema: "ERRO! Já existe um livro com id = 1!" ❌
```

### AUTOINCREMENT: O Gerador Automático de RGs

**AUTOINCREMENT** é como um sistema que gera automaticamente números de RG sequenciais:

```
Pessoa 1: RG = 1 (gerado automaticamente)
Pessoa 2: RG = 2 (gerado automaticamente)
Pessoa 3: RG = 3 (gerado automaticamente)
...
```

```sql
-- Você não precisa especificar o ID
INSERT INTO livros (titulo) VALUES ('Dom Casmurro');
-- Sistema automaticamente atribui id = 1

INSERT INTO livros (titulo) VALUES ('1984');
-- Sistema automaticamente atribui id = 2
```

É como um **contador automático** que sempre dá o próximo número disponível!

### Analogia Visual: Armários com Números

Pense em PRIMARY KEY como armários com números:

```
Armário 1: [Dom Casmurro]
Armário 2: [1984]
Armário 3: [Fundação]
Armário 4: [vazio]

-- Tentar colocar outro livro no Armário 1 (ERRO!)
Sistema: "Armário 1 já está ocupado! Use o Armário 4!"
```

Cada armário (registro) tem um número único (PRIMARY KEY) e não pode ser compartilhado!

---

## 3. FOREIGN KEY: O Link entre Tabelas

### Analogia: Referência entre Documentos

Pense em FOREIGN KEY como uma **referência entre documentos**. Por exemplo, em um empréstimo de livro, você precisa referenciar:
- Qual livro foi emprestado (referência à tabela de livros)
- Quem pegou emprestado (referência à tabela de usuários)

É como preencher um formulário onde você precisa colocar o número do livro e o número do usuário que existem em outros cadastros.

### Exemplo Prático: Empréstimo de Livro

```sql
-- Tabela de livros (tabela "pai")
CREATE TABLE livros (
    id INTEGER PRIMARY KEY,
    titulo TEXT NOT NULL
);

-- Tabela de empréstimos (tabela "filha")
CREATE TABLE emprestimos (
    id INTEGER PRIMARY KEY,
    livro_id INTEGER,  -- Referência ao livro
    FOREIGN KEY (livro_id) REFERENCES livros(id)
);
```

**O que acontece:**
```
Livros cadastrados:
- Livro 1: "Dom Casmurro"
- Livro 2: "1984"
- Livro 3: "Fundação"

Empréstimos:
- Empréstimo 1: livro_id = 1 ✅ (livro existe)
- Empréstimo 2: livro_id = 2 ✅ (livro existe)

-- Tentar emprestar livro 999 (ERRO!)
- Empréstimo 3: livro_id = 999 ❌
Sistema: "ERRO! Livro 999 não existe!"
```

### Analogia: Sistema de Biblioteca Física

Imagine uma biblioteca física:

**Sem FOREIGN KEY:**
```
Bibliotecário: "Vou registrar um empréstimo do livro número 999"
Sistema: "Ok, registrado!" ❌
(Mas o livro 999 não existe na biblioteca!)
```

**Com FOREIGN KEY:**
```
Bibliotecário: "Vou registrar um empréstimo do livro número 999"
Sistema: "ERRO! O livro 999 não está cadastrado na biblioteca!" ✅
```

FOREIGN KEY é como um **verificador** que garante que você só pode referenciar coisas que realmente existem!

### CASCADE: Ação em Cascata

**CASCADE** é como uma reação em cadeia:

```
Situação: Você deleta um livro da biblioteca

Sem CASCADE:
- Livro deletado ✅
- Empréstimos do livro ainda existem ❌ (órfãos!)

Com CASCADE:
- Livro deletado ✅
- Empréstimos do livro também são deletados automaticamente ✅
```

É como derrubar uma peça de dominó: quando você derruba uma, as outras caem automaticamente!

```sql
CREATE TABLE emprestimos (
    id INTEGER PRIMARY KEY,
    livro_id INTEGER,
    FOREIGN KEY (livro_id) REFERENCES livros(id) ON DELETE CASCADE
);

-- Deletar livro
DELETE FROM livros WHERE id = 1;
-- Empréstimos relacionados também são deletados automaticamente
```

---

## 4. UNIQUE: Valores Únicos como Email

### Analogia: Email Único

Pense em UNIQUE como a regra de que **cada pessoa precisa ter um email único**. Não pode haver duas pessoas com o mesmo email!

**Características do Email:**
- Cada pessoa tem um email único
- Não pode haver duas pessoas com o mesmo email
- Mas pode haver pessoas sem email (NULL)

**Características da UNIQUE:**
- Cada registro tem um valor único na coluna
- Não pode haver dois registros com o mesmo valor
- Mas pode haver registros com NULL (dependendo da implementação)

### Exemplo Prático

```sql
CREATE TABLE usuarios (
    id INTEGER PRIMARY KEY,
    nome TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL  -- Email deve ser único
);
```

**O que acontece:**
```
Usuários cadastrados:
- Usuário 1: email = "joao@email.com" ✅
- Usuário 2: email = "maria@email.com" ✅

-- Tentar cadastrar outro usuário com mesmo email (ERRO!)
- Usuário 3: email = "joao@email.com" ❌
Sistema: "ERRO! Este email já está cadastrado!"
```

### Analogia: Números de Telefone

Pense em UNIQUE como números de telefone:

```
Telefone 1: (11) 98765-4321 → João
Telefone 2: (11) 97654-3210 → Maria

-- Tentar cadastrar outro telefone igual (ERRO!)
Telefone 3: (11) 98765-4321 → Pedro ❌
Sistema: "ERRO! Este número já está em uso!"
```

### UNIQUE vs PRIMARY KEY

**PRIMARY KEY** = RG (sempre único, sempre existe, identifica o registro)
**UNIQUE** = Email (único, mas não necessariamente identifica o registro)

```
PRIMARY KEY (id):
- Sempre existe ✅
- Sempre único ✅
- Identifica o registro ✅

UNIQUE (email):
- Pode ser NULL (depende) ⚠️
- Sempre único ✅
- Não identifica necessariamente o registro ❌
```

---

## 5. NOT NULL: Campo Obrigatório

### Analogia: Formulário com Campos Obrigatórios

Pense em NOT NULL como campos **obrigatórios em um formulário**. Você não pode deixar em branco!

**Formulário de Cadastro:**
```
Nome: [_____________] ← OBRIGATÓRIO (NOT NULL)
Email: [_____________] ← OBRIGATÓRIO (NOT NULL)
Telefone: [_____________] ← OPCIONAL (pode ser NULL)
```

Se você tentar enviar o formulário sem preencher os campos obrigatórios, o sistema não aceita!

### Exemplo Prático

```sql
CREATE TABLE livros (
    id INTEGER PRIMARY KEY,
    titulo TEXT NOT NULL,  -- OBRIGATÓRIO
    isbn TEXT,             -- OPCIONAL
    descricao TEXT         -- OPCIONAL
);
```

**O que acontece:**
```
-- Tentar cadastrar livro sem título (ERRO!)
INSERT INTO livros (isbn) VALUES ('123');
Sistema: "ERRO! Título é obrigatório!" ❌

-- Correto: fornecer título
INSERT INTO livros (titulo, isbn) VALUES ('Dom Casmurro', '123');
Sistema: "Cadastrado com sucesso!" ✅
```

### Analogia: Checklist de Viagem

Pense em NOT NULL como itens **obrigatórios em uma checklist de viagem**:

```
Checklist de Viagem:
☑ Passaporte (OBRIGATÓRIO - NOT NULL)
☑ Passagem (OBRIGATÓRIO - NOT NULL)
☐ Seguro viagem (OPCIONAL - pode ser NULL)
☐ Guia turístico (OPCIONAL - pode ser NULL)

-- Tentar viajar sem passaporte (ERRO!)
Sistema: "ERRO! Você não pode viajar sem passaporte!" ❌
```

### NULL vs String Vazia

**NULL** = Ausência de valor (não preenchido)
**String vazia (`''`)** = Valor presente, mas vazio

```
NULL: "Este campo não foi preenchido"
String vazia: "Este campo foi preenchido, mas está vazio"

NOT NULL previne NULL, mas permite string vazia:
- NULL ❌ (bloqueado por NOT NULL)
- '' ✅ (permitido, é um valor)
```

---

## 6. CHECK: O Validador Personalizado

### Analogia: Validação de Formulário

Pense em CHECK como **validações customizadas em um formulário**. Por exemplo:
- Idade deve ser entre 18 e 100 anos
- Quantidade deve ser maior ou igual a zero
- Status deve ser "ativo", "inativo" ou "pendente"

É como ter um **validador inteligente** que verifica se os dados fazem sentido antes de aceitar!

### Exemplo Prático: Validação de Quantidade

```sql
CREATE TABLE livros (
    id INTEGER PRIMARY KEY,
    titulo TEXT NOT NULL,
    quantidade_disponivel INTEGER CHECK (quantidade_disponivel >= 0)
);
```

**O que acontece:**
```
-- Tentar cadastrar quantidade negativa (ERRO!)
INSERT INTO livros (titulo, quantidade_disponivel) 
VALUES ('Dom Casmurro', -5);
Sistema: "ERRO! Quantidade não pode ser negativa!" ❌

-- Correto: quantidade válida
INSERT INTO livros (titulo, quantidade_disponivel) 
VALUES ('Dom Casmurro', 10);
Sistema: "Cadastrado com sucesso!" ✅
```

### Analogia: Validação de Idade

Pense em CHECK como validação de idade em um formulário:

```
Formulário de Cadastro:
Nome: [_____________]
Idade: [____] ← Deve ser entre 18 e 100

-- Tentar cadastrar idade 15 (ERRO!)
Sistema: "ERRO! Idade deve ser entre 18 e 100 anos!" ❌

-- Tentar cadastrar idade 150 (ERRO!)
Sistema: "ERRO! Idade deve ser entre 18 e 100 anos!" ❌

-- Correto: idade válida
Idade: 25 ✅
Sistema: "Cadastrado com sucesso!" ✅
```

### Exemplo: Validação de Status

```sql
CREATE TABLE emprestimos (
    id INTEGER PRIMARY KEY,
    status TEXT CHECK (status IN ('ativo', 'devolvido', 'atrasado'))
);
```

**O que acontece:**
```
-- Tentar cadastrar status inválido (ERRO!)
INSERT INTO emprestimos (status) VALUES ('cancelado');
Sistema: "ERRO! Status deve ser 'ativo', 'devolvido' ou 'atrasado'!" ❌

-- Correto: status válido
INSERT INTO emprestimos (status) VALUES ('ativo');
Sistema: "Cadastrado com sucesso!" ✅
```

### Analogia: Porteiro Inteligente

Pense em CHECK como um **porteiro inteligente** que verifica regras específicas:

```
Porteiro: "Quantos livros você quer cadastrar?"
Você: "-5 livros"
Porteiro: "ERRO! Não pode ter quantidade negativa!" ❌

Porteiro: "Qual o status do empréstimo?"
Você: "cancelado"
Porteiro: "ERRO! Status deve ser 'ativo', 'devolvido' ou 'atrasado'!" ❌
```

CHECK é como ter um **validador personalizado** que conhece as regras específicas do seu negócio!

---

## 7. Combinando Constraints: O Sistema Completo

### Analogia: Sistema de Biblioteca Completo

Imagine um sistema completo de biblioteca com todas as regras:

```sql
CREATE TABLE emprestimos (
    -- PRIMARY KEY: Cada empréstimo tem um número único (como RG)
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    
    -- FOREIGN KEY: Só pode referenciar livros que existem
    livro_id INTEGER NOT NULL,
    usuario_id INTEGER NOT NULL,
    
    -- NOT NULL: Data de empréstimo é obrigatória
    data_emprestimo DATE NOT NULL DEFAULT CURRENT_DATE,
    
    -- CHECK: Status deve ser um dos valores permitidos
    status TEXT NOT NULL DEFAULT 'ativo' 
        CHECK (status IN ('ativo', 'devolvido', 'atrasado', 'cancelado')),
    
    -- FOREIGN KEY: Garante integridade referencial
    FOREIGN KEY (livro_id) REFERENCES livros(id),
    FOREIGN KEY (usuario_id) REFERENCES usuarios(id)
);
```

**O que cada constraint faz:**

1. **PRIMARY KEY (id)**: Garante que cada empréstimo tenha um número único
   - Como: RG do empréstimo

2. **FOREIGN KEY (livro_id)**: Garante que só pode referenciar livros existentes
   - Como: Verificador que o livro existe na biblioteca

3. **NOT NULL (data_emprestimo)**: Garante que sempre tenha data
   - Como: Campo obrigatório no formulário

4. **CHECK (status)**: Garante que status seja um dos valores permitidos
   - Como: Validador que verifica se o status é válido

5. **DEFAULT**: Fornece valores padrão quando não especificados
   - Como: Preenchimento automático do formulário

### Exemplo de Uso

```sql
-- Tentar criar empréstimo inválido (múltiplos erros!)
INSERT INTO emprestimos (livro_id, usuario_id, status) 
VALUES (999, 1, 'cancelado');
-- ERRO 1: Livro 999 não existe (FOREIGN KEY)
-- ERRO 2: Status 'cancelado' não é permitido (CHECK)

-- Correto: empréstimo válido
INSERT INTO emprestimos (livro_id, usuario_id) 
VALUES (1, 1);
-- Sistema automaticamente:
-- - Atribui id único (PRIMARY KEY)
-- - Verifica que livro 1 existe (FOREIGN KEY)
-- - Verifica que usuário 1 existe (FOREIGN KEY)
-- - Define data_emprestimo = hoje (DEFAULT)
-- - Define status = 'ativo' (DEFAULT)
```

---

## 8. Ordem de Verificação: O Processo de Validação

### Analogia: Processo de Entrada em um Evento

Pense nas constraints como **checkpoints em um evento**:

```
Checkpoint 1: Verificar se tem ingresso (NOT NULL)
Checkpoint 2: Verificar se o ingresso é válido (CHECK)
Checkpoint 3: Verificar se o ingresso não foi usado (UNIQUE)
Checkpoint 4: Verificar se a pessoa está na lista (FOREIGN KEY)
Checkpoint 5: Verificar se não é duplicata (PRIMARY KEY)
```

**Ordem de verificação das constraints:**

1. **NOT NULL**: "Este campo foi preenchido?"
2. **CHECK**: "O valor atende às condições?"
3. **UNIQUE**: "Este valor já existe?"
4. **PRIMARY KEY**: "Este ID já existe?"
5. **FOREIGN KEY**: "O registro referenciado existe?"

### Exemplo Prático

```sql
-- Tentar inserir registro inválido
INSERT INTO emprestimos (livro_id, usuario_id, status) 
VALUES (NULL, 1, 'ativo');

-- Ordem de verificação:
1. NOT NULL: livro_id é NULL? ❌ ERRO! (para aqui)
2. CHECK: (não chega aqui)
3. UNIQUE: (não chega aqui)
4. PRIMARY KEY: (não chega aqui)
5. FOREIGN KEY: (não chega aqui)
```

---

## 9. Erros Comuns: O que Acontece Quando Você Viola Constraints

### Erro: UNIQUE constraint failed

**Analogia**: Tentar cadastrar o mesmo email duas vezes

```
Situação: Tentar cadastrar "joao@email.com" novamente
Sistema: "ERRO! Este email já está cadastrado!" ❌

Solução: Usar um email diferente ou atualizar o existente
```

### Erro: NOT NULL constraint failed

**Analogia**: Tentar enviar formulário sem preencher campo obrigatório

```
Situação: Tentar cadastrar livro sem título
Sistema: "ERRO! Título é obrigatório!" ❌

Solução: Preencher o campo obrigatório
```

### Erro: FOREIGN KEY constraint failed

**Analogia**: Tentar referenciar algo que não existe

```
Situação: Tentar emprestar livro 999 (que não existe)
Sistema: "ERRO! Livro 999 não está cadastrado!" ❌

Solução: Usar um livro que realmente existe
```

### Erro: CHECK constraint failed

**Analogia**: Tentar inserir valor que não atende à validação

```
Situação: Tentar cadastrar quantidade -5
Sistema: "ERRO! Quantidade não pode ser negativa!" ❌

Solução: Usar um valor válido (>= 0)
```

---

## 10. Boas Práticas: Regras de Ouro

### 1. Sempre Use PRIMARY KEY

**Analogia**: Todo documento precisa de um número de identificação

```
✅ BOM: Toda tabela tem PRIMARY KEY
❌ EVITE: Tabela sem PRIMARY KEY (como documento sem RG)
```

### 2. Use FOREIGN KEY para Relacionamentos

**Analogia**: Sempre verificar se o que você referencia existe

```
✅ BOM: Empréstimo referencia livro que existe
❌ EVITE: Empréstimo referencia livro que não existe
```

### 3. Use NOT NULL para Campos Essenciais

**Analogia**: Campos obrigatórios em formulários

```
✅ BOM: Título do livro é obrigatório
❌ EVITE: Título pode ser vazio (dados incompletos)
```

### 4. Use UNIQUE para Valores que Devem Ser Únicos

**Analogia**: Email, CPF, ISBN devem ser únicos

```
✅ BOM: Email de usuário é único
❌ EVITE: Múltiplos usuários com mesmo email
```

### 5. Use CHECK para Validações Simples

**Analogia**: Validações que não dependem de outras tabelas

```
✅ BOM: Quantidade >= 0
❌ EVITE: Validações complexas que dependem de outras tabelas
```

---

## 11. Resumo Visual

### Constraints como Porteiros

```
┌─────────────────────────────────────┐
│     BANCO DE DADOS (Biblioteca)     │
│                                     │
│  ┌───────────────────────────────┐ │
│  │  PRIMARY KEY (Porteiro 1)     │ │
│  │  "Você tem um ID único?"     │ │
│  └───────────────────────────────┘ │
│              ↓                      │
│  ┌───────────────────────────────┐ │
│  │  NOT NULL (Porteiro 2)        │ │
│  │  "Você preencheu tudo?"       │ │
│  └───────────────────────────────┘ │
│              ↓                      │
│  ┌───────────────────────────────┐ │
│  │  CHECK (Porteiro 3)           │ │
│  │  "Os valores são válidos?"    │ │
│  └───────────────────────────────┘ │
│              ↓                      │
│  ┌───────────────────────────────┐ │
│  │  UNIQUE (Porteiro 4)         │ │
│  │  "Este valor já existe?"     │ │
│  └───────────────────────────────┘ │
│              ↓                      │
│  ┌───────────────────────────────┐ │
│  │  FOREIGN KEY (Porteiro 5)    │ │
│  │  "O que você referencia      │ │
│  │   realmente existe?"         │ │
│  └───────────────────────────────┘ │
│              ↓                      │
│     ✅ DADOS VÁLIDOS ENTRAM!        │
└─────────────────────────────────────┘
```

### Tabela de Comparação

| Constraint | Analogia | Função |
|------------|----------|--------|
| **PRIMARY KEY** | RG do registro | Identifica unicamente |
| **FOREIGN KEY** | Referência entre documentos | Mantém relacionamentos |
| **UNIQUE** | Email único | Garante valores únicos |
| **NOT NULL** | Campo obrigatório | Previne valores vazios |
| **CHECK** | Validador personalizado | Valida regras customizadas |

---

**Próximo Passo**: Complete os **Exercícios e Reflexão** (`03-exercicios-reflexao.md`) para praticar criando e testando constraints!
