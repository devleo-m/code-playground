# Aula 5 - Exercícios e Reflexão: Data Constraints

## Exercícios Práticos

### Exercício 1: Criando Tabelas com PRIMARY KEY

**Objetivo**: Praticar criação de tabelas com PRIMARY KEY e entender AUTOINCREMENT.

**Tarefas**:

1. Crie uma tabela chamada `editoras` com as seguintes colunas:
   - `id` (INTEGER, PRIMARY KEY com AUTOINCREMENT)
   - `nome` (TEXT, NOT NULL)
   - `cidade` (TEXT)
   - `ano_fundacao` (INTEGER)

2. Insira 3 editoras na tabela (sem especificar o ID, deixe o AUTOINCREMENT funcionar).

3. Verifique os IDs gerados automaticamente.

4. Tente inserir uma editora com um ID que já existe (deve dar erro).

**Questão de Reflexão**:
- Por que é importante que cada tabela tenha uma PRIMARY KEY? O que aconteceria se não tivéssemos PRIMARY KEY?

**Soluções Esperadas**:

```sql
-- 1. Criar tabela com PRIMARY KEY
CREATE TABLE editoras (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    nome TEXT NOT NULL,
    cidade TEXT,
    ano_fundacao INTEGER
);

-- 2. Inserir editoras (sem especificar ID)
INSERT INTO editoras (nome, cidade, ano_fundacao) 
VALUES ('Companhia das Letras', 'São Paulo', 1986);

INSERT INTO editoras (nome, cidade, ano_fundacao) 
VALUES ('Globo', 'Rio de Janeiro', 1925);

INSERT INTO editoras (nome, cidade, ano_fundacao) 
VALUES ('Rocco', 'Rio de Janeiro', 1975);

-- 3. Verificar IDs gerados
SELECT * FROM editoras;
-- Deve mostrar IDs 1, 2, 3 automaticamente

-- 4. Tentar inserir com ID duplicado (deve dar erro)
INSERT INTO editoras (id, nome) VALUES (1, 'Editora Teste');
-- Erro: UNIQUE constraint failed: editoras.id
```

**Resposta Esperada para a Questão de Reflexão**:
- **Importância da PRIMARY KEY**:
  - Identifica unicamente cada registro
  - Permite referências precisas (FOREIGN KEY)
  - Melhora performance com índice automático
  - Facilita JOINs e relacionamentos
- **Sem PRIMARY KEY**:
  - Não há como identificar registros únicos
  - Dificulta relacionamentos entre tabelas
  - Pode haver duplicatas
  - Queries podem ser mais lentas
  - Dificulta atualizações e deleções precisas

---

### Exercício 2: Trabalhando com FOREIGN KEY

**Objetivo**: Praticar criação e uso de FOREIGN KEY, incluindo habilitar verificação no SQLite.

**⚠️ IMPORTANTE**: Lembre-se de habilitar `PRAGMA foreign_keys = ON;` antes de trabalhar com FOREIGN KEY no SQLite!

**Tarefas**:

1. Habilite a verificação de FOREIGN KEY no SQLite.

2. Crie uma tabela `avaliacoes` com as seguintes colunas:
   - `id` (INTEGER, PRIMARY KEY AUTOINCREMENT)
   - `livro_id` (INTEGER, NOT NULL, FOREIGN KEY referenciando `livros(id)`)
   - `usuario_id` (INTEGER, NOT NULL, FOREIGN KEY referenciando `usuarios(id)`)
   - `nota` (INTEGER, entre 1 e 5)
   - `comentario` (TEXT)
   - `data_avaliacao` (DATE, DEFAULT CURRENT_DATE)

3. Insira 3 avaliações válidas (usando livros e usuários que existem).

4. Tente inserir uma avaliação com `livro_id` que não existe (deve dar erro).

5. Tente inserir uma avaliação com `usuario_id` que não existe (deve dar erro).

**Questão de Reflexão**:
- Qual a diferença entre usar `ON DELETE RESTRICT` e `ON DELETE CASCADE` em uma FOREIGN KEY? Em que situações cada um seria mais apropriado?

**Soluções Esperadas**:

```sql
-- 1. Habilitar FOREIGN KEY
PRAGMA foreign_keys = ON;

-- Verificar se está habilitado
PRAGMA foreign_keys;
-- Deve retornar 1

-- 2. Criar tabela com FOREIGN KEY
CREATE TABLE avaliacoes (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    livro_id INTEGER NOT NULL,
    usuario_id INTEGER NOT NULL,
    nota INTEGER CHECK (nota >= 1 AND nota <= 5),
    comentario TEXT,
    data_avaliacao DATE DEFAULT CURRENT_DATE,
    FOREIGN KEY (livro_id) REFERENCES livros(id) ON DELETE RESTRICT,
    FOREIGN KEY (usuario_id) REFERENCES usuarios(id) ON DELETE RESTRICT
);

-- 3. Inserir avaliações válidas
INSERT INTO avaliacoes (livro_id, usuario_id, nota, comentario)
VALUES (1, 1, 5, 'Excelente livro!');

INSERT INTO avaliacoes (livro_id, usuario_id, nota, comentario)
VALUES (2, 2, 4, 'Muito bom, recomendo.');

INSERT INTO avaliacoes (livro_id, usuario_id, nota, comentario)
VALUES (3, 3, 5, 'Obra-prima!');

-- 4. Tentar inserir com livro inexistente (deve dar erro)
INSERT INTO avaliacoes (livro_id, usuario_id, nota)
VALUES (999, 1, 5);
-- Erro: FOREIGN KEY constraint failed

-- 5. Tentar inserir com usuário inexistente (deve dar erro)
INSERT INTO avaliacoes (livro_id, usuario_id, nota)
VALUES (1, 999, 5);
-- Erro: FOREIGN KEY constraint failed
```

**Resposta Esperada para a Questão de Reflexão**:
- **ON DELETE RESTRICT**:
  - Bloqueia a deleção se houver referências
  - Protege dados importantes de serem deletados acidentalmente
  - Mais seguro, mas exige deleção manual de dependentes
  - **Uso apropriado**: Quando dados relacionados são importantes e não devem ser deletados automaticamente (ex: empréstimos, histórico)
- **ON DELETE CASCADE**:
  - Deleta automaticamente registros relacionados
  - Mais conveniente, mas pode deletar dados importantes
  - Pode causar perda de dados se usado incorretamente
  - **Uso apropriado**: Quando dados relacionados não são importantes ou devem ser deletados junto (ex: comentários de um post, itens de um pedido)
- **Recomendação**: Use RESTRICT por padrão, CASCADE apenas quando realmente necessário e seguro

---

### Exercício 3: Trabalhando com UNIQUE

**Objetivo**: Praticar criação e uso de constraint UNIQUE.

**Tarefas**:

1. Verifique a estrutura atual da tabela `usuarios` e veja se já tem constraint UNIQUE no email.

2. Se não tiver, adicione uma constraint UNIQUE ao email (crie um índice único).

3. Tente inserir dois usuários com o mesmo email (o segundo deve dar erro).

4. Tente inserir dois usuários com email NULL (ambos devem ser permitidos, se a implementação permitir múltiplos NULL).

5. Crie uma tabela `isbns` com:
   - `id` (INTEGER, PRIMARY KEY AUTOINCREMENT)
   - `isbn` (TEXT, UNIQUE NOT NULL)
   - `livro_id` (INTEGER, FOREIGN KEY referenciando `livros(id)`)
   - `data_registro` (DATE, DEFAULT CURRENT_DATE)

**Questão de Reflexão**:
- Qual a diferença entre PRIMARY KEY e UNIQUE? Quando você usaria cada um?

**Soluções Esperadas**:

```sql
-- 1. Verificar estrutura atual
.schema usuarios
-- Ou
SELECT sql FROM sqlite_master WHERE type='table' AND name='usuarios';

-- 2. Adicionar UNIQUE ao email (se não existir)
CREATE UNIQUE INDEX idx_usuarios_email ON usuarios(email);

-- 3. Tentar inserir usuários com mesmo email (deve dar erro no segundo)
INSERT INTO usuarios (nome, email) VALUES ('João Silva', 'joao@email.com');
-- Sucesso

INSERT INTO usuarios (nome, email) VALUES ('João Santos', 'joao@email.com');
-- Erro: UNIQUE constraint failed: usuarios.email

-- 4. Tentar inserir dois usuários com email NULL
INSERT INTO usuarios (nome, email) VALUES ('Usuário 1', NULL);
-- Pode ser permitido dependendo da implementação

INSERT INTO usuarios (nome, email) VALUES ('Usuário 2', NULL);
-- Pode ser permitido dependendo da implementação

-- 5. Criar tabela com UNIQUE
CREATE TABLE isbns (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    isbn TEXT UNIQUE NOT NULL,
    livro_id INTEGER,
    data_registro DATE DEFAULT CURRENT_DATE,
    FOREIGN KEY (livro_id) REFERENCES livros(id)
);
```

**Resposta Esperada para a Questão de Reflexão**:
- **PRIMARY KEY**:
  - Identifica unicamente cada registro (é o "RG" do registro)
  - Não pode ser NULL
  - Uma por tabela
  - Cria índice automaticamente
  - Usado para identificar e referenciar registros
- **UNIQUE**:
  - Garante valores únicos, mas não necessariamente identifica o registro
  - Pode ser NULL (dependendo da implementação)
  - Múltiplas por tabela
  - Cria índice automaticamente
  - Usado para campos que devem ser únicos mas não são identificadores (email, CPF, ISBN)
- **Quando usar cada um**:
  - **PRIMARY KEY**: Sempre! Toda tabela precisa de uma para identificar registros
  - **UNIQUE**: Para campos que devem ser únicos mas não são o identificador principal (email, CPF, matrícula, código de produto)

---

### Exercício 4: Trabalhando com NOT NULL

**Objetivo**: Praticar criação e uso de constraint NOT NULL.

**Tarefas**:

1. Crie uma tabela `categorias_livros` com:
   - `id` (INTEGER, PRIMARY KEY AUTOINCREMENT)
   - `nome` (TEXT, NOT NULL)
   - `descricao` (TEXT) - pode ser NULL
   - `data_criacao` (DATE, NOT NULL, DEFAULT CURRENT_DATE)

2. Insira 3 categorias, algumas com descrição e algumas sem.

3. Tente inserir uma categoria sem nome (deve dar erro).

4. Tente inserir uma categoria sem data_criacao (deve funcionar, usando o DEFAULT).

5. Verifique quais campos são NOT NULL na tabela `livros` existente.

**Questão de Reflexão**:
- Por que é importante usar NOT NULL em campos essenciais? Quais problemas podem ocorrer se campos importantes forem NULL?

**Soluções Esperadas**:

```sql
-- 1. Criar tabela com NOT NULL
CREATE TABLE categorias_livros (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    nome TEXT NOT NULL,
    descricao TEXT,  -- Pode ser NULL
    data_criacao DATE NOT NULL DEFAULT CURRENT_DATE
);

-- 2. Inserir categorias
INSERT INTO categorias_livros (nome, descricao) 
VALUES ('Ficção', 'Livros de ficção científica e fantasia');

INSERT INTO categorias_livros (nome) 
VALUES ('Romance');  -- descricao será NULL

INSERT INTO categorias_livros (nome, descricao) 
VALUES ('Técnico', 'Livros técnicos de programação');

-- 3. Tentar inserir sem nome (deve dar erro)
INSERT INTO categorias_livros (descricao) 
VALUES ('Categoria sem nome');
-- Erro: NOT NULL constraint failed: categorias_livros.nome

-- 4. Inserir sem data_criacao (deve usar DEFAULT)
INSERT INTO categorias_livros (nome) 
VALUES ('História');
-- data_criacao será automaticamente CURRENT_DATE

-- Verificar
SELECT * FROM categorias_livros;

-- 5. Verificar NOT NULL na tabela livros
.schema livros
-- Ou
SELECT sql FROM sqlite_master WHERE type='table' AND name='livros';
```

**Resposta Esperada para a Questão de Reflexão**:
- **Importância do NOT NULL**:
  - Garante que dados essenciais sempre existam
  - Previne dados incompletos ou inválidos
  - Facilita queries (não precisa verificar NULL)
  - Melhora integridade dos dados
- **Problemas sem NOT NULL**:
  - Dados incompletos podem quebrar aplicações
  - Queries podem retornar resultados inesperados
  - Lógica de negócio pode falhar
  - Relatórios podem ter dados faltantes
  - JOINs podem excluir registros importantes
- **Exemplo**: Se `titulo` de um livro for NULL, como você exibe o livro em uma lista? Como você busca por título?

---

### Exercício 5: Trabalhando com CHECK

**Objetivo**: Praticar criação e uso de constraint CHECK.

**Tarefas**:

1. Crie uma tabela `reservas` com:
   - `id` (INTEGER, PRIMARY KEY AUTOINCREMENT)
   - `livro_id` (INTEGER, NOT NULL, FOREIGN KEY)
   - `usuario_id` (INTEGER, NOT NULL, FOREIGN KEY)
   - `data_reserva` (DATE, NOT NULL, DEFAULT CURRENT_DATE)
   - `data_limite` (DATE, NOT NULL)
   - `status` (TEXT, NOT NULL, DEFAULT 'pendente', CHECK: 'pendente', 'confirmada', 'cancelada')
   - CHECK: `data_limite` deve ser maior ou igual a `data_reserva`

2. Insira 3 reservas válidas.

3. Tente inserir uma reserva com status inválido (deve dar erro).

4. Tente inserir uma reserva com `data_limite` anterior a `data_reserva` (deve dar erro).

5. Crie uma tabela `descontos` com:
   - `id` (INTEGER, PRIMARY KEY AUTOINCREMENT)
   - `porcentagem` (INTEGER, CHECK entre 0 e 100)
   - `data_inicio` (DATE, NOT NULL)
   - `data_fim` (DATE, NOT NULL)
   - CHECK: `data_fim` deve ser maior ou igual a `data_inicio`

**Questão de Reflexão**:
- Quais são as limitações do CHECK? Quando você deveria usar CHECK vs validação na aplicação?

**Soluções Esperadas**:

```sql
-- 1. Criar tabela com CHECK
PRAGMA foreign_keys = ON;

CREATE TABLE reservas (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    livro_id INTEGER NOT NULL,
    usuario_id INTEGER NOT NULL,
    data_reserva DATE NOT NULL DEFAULT CURRENT_DATE,
    data_limite DATE NOT NULL,
    status TEXT NOT NULL DEFAULT 'pendente' 
        CHECK (status IN ('pendente', 'confirmada', 'cancelada')),
    FOREIGN KEY (livro_id) REFERENCES livros(id),
    FOREIGN KEY (usuario_id) REFERENCES usuarios(id),
    CHECK (data_limite >= data_reserva)
);

-- 2. Inserir reservas válidas
INSERT INTO reservas (livro_id, usuario_id, data_limite)
VALUES (1, 1, '2024-12-31');

INSERT INTO reservas (livro_id, usuario_id, data_reserva, data_limite)
VALUES (2, 2, '2024-01-15', '2024-02-15');

INSERT INTO reservas (livro_id, usuario_id, data_limite, status)
VALUES (3, 3, '2024-12-31', 'confirmada');

-- 3. Tentar inserir com status inválido (deve dar erro)
INSERT INTO reservas (livro_id, usuario_id, data_limite, status)
VALUES (1, 1, '2024-12-31', 'expirada');
-- Erro: CHECK constraint failed: reservas

-- 4. Tentar inserir com data_limite inválida (deve dar erro)
INSERT INTO reservas (livro_id, usuario_id, data_reserva, data_limite)
VALUES (1, 1, '2024-12-31', '2024-01-01');
-- Erro: CHECK constraint failed: reservas

-- 5. Criar tabela descontos com CHECK
CREATE TABLE descontos (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    porcentagem INTEGER CHECK (porcentagem >= 0 AND porcentagem <= 100),
    data_inicio DATE NOT NULL,
    data_fim DATE NOT NULL,
    CHECK (data_fim >= data_inicio)
);
```

**Resposta Esperada para a Questão de Reflexão**:
- **Limitações do CHECK**:
  - Não pode usar subqueries (não pode referenciar outras tabelas)
  - Não pode usar funções não-determinísticas
  - Pode impactar performance em INSERT/UPDATE frequentes
  - Validação básica apenas
  - Não pode ter lógica de negócio complexa
- **Quando usar CHECK**:
  - Validações simples de intervalo (idade entre 18-100)
  - Validações de lista de valores (status IN ('ativo', 'inativo'))
  - Validações de formato básico (email LIKE '%@%.%')
  - Validações que não dependem de outras tabelas
- **Quando usar validação na aplicação**:
  - Validações complexas que dependem de outras tabelas
  - Validações que mudam frequentemente
  - Validações que precisam de lógica de negócio complexa
  - Validações que precisam de mensagens de erro personalizadas
  - Validações que precisam de contexto da aplicação
- **Recomendação**: Use CHECK para validações simples e estáveis, validação na aplicação para o resto

---

### Exercício 6: Combinando Múltiplas Constraints

**Objetivo**: Praticar criação de tabelas com múltiplas constraints combinadas.

**Tarefas**:

1. Crie uma tabela `multas` com todas as constraints aprendidas:
   - `id` (PRIMARY KEY AUTOINCREMENT)
   - `emprestimo_id` (FOREIGN KEY, NOT NULL)
   - `valor` (REAL, NOT NULL, CHECK >= 0)
   - `data_multa` (DATE, NOT NULL, DEFAULT CURRENT_DATE)
   - `data_vencimento` (DATE, NOT NULL)
   - `status` (TEXT, NOT NULL, DEFAULT 'pendente', CHECK: 'pendente', 'paga', 'cancelada')
   - `observacao` (TEXT) - pode ser NULL
   - CHECK: `data_vencimento` >= `data_multa`

2. Insira 2 multas válidas.

3. Tente violar cada tipo de constraint (PRIMARY KEY, FOREIGN KEY, NOT NULL, CHECK) e documente os erros.

4. Crie uma tabela `eventos_biblioteca` com:
   - `id` (PRIMARY KEY AUTOINCREMENT)
   - `titulo` (TEXT, UNIQUE NOT NULL)
   - `data_evento` (DATE, NOT NULL)
   - `capacidade` (INTEGER, CHECK >= 1 AND <= 500)
   - `inscritos` (INTEGER, DEFAULT 0, CHECK >= 0)
   - CHECK: `inscritos` <= `capacidade`

**Questão de Reflexão**:
- Qual a ordem de verificação das constraints quando você insere um registro? Por que essa ordem importa?

**Soluções Esperadas**:

```sql
-- 1. Criar tabela com múltiplas constraints
PRAGMA foreign_keys = ON;

CREATE TABLE multas (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    emprestimo_id INTEGER NOT NULL,
    valor REAL NOT NULL CHECK (valor >= 0),
    data_multa DATE NOT NULL DEFAULT CURRENT_DATE,
    data_vencimento DATE NOT NULL,
    status TEXT NOT NULL DEFAULT 'pendente' 
        CHECK (status IN ('pendente', 'paga', 'cancelada')),
    observacao TEXT,
    FOREIGN KEY (emprestimo_id) REFERENCES emprestimos(id) ON DELETE RESTRICT,
    CHECK (data_vencimento >= data_multa)
);

-- 2. Inserir multas válidas
INSERT INTO multas (emprestimo_id, valor, data_vencimento)
VALUES (1, 5.50, '2024-12-31');

INSERT INTO multas (emprestimo_id, valor, data_multa, data_vencimento, status)
VALUES (2, 10.00, '2024-01-15', '2024-02-15', 'paga');

-- 3. Tentar violar cada constraint
-- PRIMARY KEY: Tentar inserir com ID duplicado
INSERT INTO multas (id, emprestimo_id, valor, data_vencimento)
VALUES (1, 1, 5.50, '2024-12-31');
-- Erro: UNIQUE constraint failed: multas.id

-- FOREIGN KEY: Tentar inserir com emprestimo inexistente
INSERT INTO multas (emprestimo_id, valor, data_vencimento)
VALUES (999, 5.50, '2024-12-31');
-- Erro: FOREIGN KEY constraint failed

-- NOT NULL: Tentar inserir sem valor obrigatório
INSERT INTO multas (valor, data_vencimento)
VALUES (5.50, '2024-12-31');
-- Erro: NOT NULL constraint failed: multas.emprestimo_id

-- CHECK: Tentar inserir valor negativo
INSERT INTO multas (emprestimo_id, valor, data_vencimento)
VALUES (1, -5.50, '2024-12-31');
-- Erro: CHECK constraint failed: multas

-- CHECK: Tentar inserir status inválido
INSERT INTO multas (emprestimo_id, valor, data_vencimento, status)
VALUES (1, 5.50, '2024-12-31', 'expirada');
-- Erro: CHECK constraint failed: multas

-- CHECK: Tentar inserir data_vencimento inválida
INSERT INTO multas (emprestimo_id, valor, data_multa, data_vencimento)
VALUES (1, 5.50, '2024-12-31', '2024-01-01');
-- Erro: CHECK constraint failed: multas

-- 4. Criar tabela eventos_biblioteca
CREATE TABLE eventos_biblioteca (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    titulo TEXT UNIQUE NOT NULL,
    data_evento DATE NOT NULL,
    capacidade INTEGER CHECK (capacidade >= 1 AND capacidade <= 500),
    inscritos INTEGER DEFAULT 0 CHECK (inscritos >= 0),
    CHECK (inscritos <= capacidade)
);
```

**Resposta Esperada para a Questão de Reflexão**:
- **Ordem de verificação**:
  1. **NOT NULL**: Verifica se campos obrigatórios têm valores
  2. **CHECK**: Verifica condições customizadas
  3. **UNIQUE**: Verifica valores únicos
  4. **PRIMARY KEY**: Verifica chave primária (também é UNIQUE)
  5. **FOREIGN KEY**: Verifica integridade referencial
- **Por que essa ordem importa**:
  - **Eficiência**: Verifica erros simples primeiro (mais rápido)
  - **Lógica**: Validações básicas antes de validações complexas
  - **Clareza**: Erros mais específicos aparecem primeiro
  - **Performance**: Evita verificações desnecessárias se já falhou em algo básico
- **Exemplo**: Se um campo NOT NULL está faltando, não faz sentido verificar FOREIGN KEY ou UNIQUE

---

### Exercício 7: Analisando Constraints Existentes

**Objetivo**: Analisar e entender constraints em tabelas existentes.

**Tarefas**:

1. Liste todas as tabelas do banco de dados.

2. Para cada tabela, identifique:
   - Quais colunas têm PRIMARY KEY
   - Quais colunas têm FOREIGN KEY
   - Quais colunas têm UNIQUE
   - Quais colunas têm NOT NULL
   - Quais colunas têm CHECK

3. Tente identificar relacionamentos entre tabelas através das FOREIGN KEY.

4. Documente suas descobertas em um diagrama simples.

**Questão de Reflexão**:
- Como as constraints ajudam a documentar a estrutura e relacionamentos do banco de dados? Por que isso é importante?

**Soluções Esperadas**:

```sql
-- 1. Listar todas as tabelas
SELECT name FROM sqlite_master WHERE type='table';

-- 2. Analisar cada tabela
-- Tabela: categorias
.schema categorias
-- PRIMARY KEY: id
-- UNIQUE: nome
-- NOT NULL: nome

-- Tabela: autores
.schema autores
-- PRIMARY KEY: id
-- NOT NULL: nome

-- Tabela: livros
.schema livros
-- PRIMARY KEY: id
-- FOREIGN KEY: autor_id -> autores(id), categoria_id -> categorias(id)
-- UNIQUE: isbn
-- NOT NULL: titulo

-- Tabela: usuarios
.schema usuarios
-- PRIMARY KEY: id
-- UNIQUE: email
-- NOT NULL: nome, email

-- Tabela: emprestimos
.schema emprestimos
-- PRIMARY KEY: id
-- FOREIGN KEY: livro_id -> livros(id), usuario_id -> usuarios(id)
-- NOT NULL: livro_id, usuario_id

-- 3. Identificar relacionamentos
-- categorias (1) -> (N) livros
-- autores (1) -> (N) livros
-- livros (1) -> (N) emprestimos
-- usuarios (1) -> (N) emprestimos

-- 4. Diagrama simples
/*
categorias
  └─< livros (categoria_id)
  
autores
  └─< livros (autor_id)
  
livros
  └─< emprestimos (livro_id)
  
usuarios
  └─< emprestimos (usuario_id)
*/
```

**Resposta Esperada para a Questão de Reflexão**:
- **Como constraints documentam**:
  - **PRIMARY KEY**: Mostra como identificar registros únicos
  - **FOREIGN KEY**: Mostra relacionamentos entre tabelas
  - **UNIQUE**: Mostra campos que devem ser únicos
  - **NOT NULL**: Mostra campos obrigatórios
  - **CHECK**: Mostra regras de validação
- **Por que é importante**:
  - **Manutenção**: Facilita entender a estrutura sem documentação externa
  - **Onboarding**: Novos desenvolvedores entendem rapidamente
  - **Debugging**: Ajuda a identificar problemas de integridade
  - **Design**: Mostra a intenção do design do banco
  - **Refatoração**: Facilita mudanças futuras
- **Exemplo**: Ver FOREIGN KEY mostra imediatamente que `emprestimos` depende de `livros` e `usuarios`

---

### Exercício 8: Testando Integridade Referencial

**Objetivo**: Testar e entender como FOREIGN KEY mantém integridade referencial.

**Tarefas**:

1. Habilite FOREIGN KEY se ainda não estiver habilitado.

2. Tente deletar um autor que tem livros associados (deve dar erro com RESTRICT).

3. Liste todos os livros de um autor específico.

4. Tente deletar um livro que tem empréstimos associados (deve dar erro).

5. Crie uma tabela de teste com CASCADE e teste o comportamento.

**Questão de Reflexão**:
- Qual é a diferença prática entre RESTRICT e CASCADE? Em que situações cada um seria mais apropriado? Quais são os riscos de usar CASCADE?

**Soluções Esperadas**:

```sql
-- 1. Habilitar FOREIGN KEY
PRAGMA foreign_keys = ON;

-- 2. Tentar deletar autor com livros (deve dar erro)
-- Primeiro, ver quais autores têm livros
SELECT a.id, a.nome, COUNT(l.id) AS total_livros
FROM autores a
LEFT JOIN livros l ON a.id = l.autor_id
GROUP BY a.id, a.nome
HAVING COUNT(l.id) > 0;

-- Tentar deletar um autor que tem livros
DELETE FROM autores WHERE id = 1;
-- Erro: FOREIGN KEY constraint failed (se RESTRICT)

-- 3. Listar livros de um autor
SELECT l.titulo, a.nome AS autor
FROM livros l
JOIN autores a ON l.autor_id = a.id
WHERE a.id = 1;

-- 4. Tentar deletar livro com empréstimos (deve dar erro)
-- Primeiro, ver quais livros têm empréstimos
SELECT l.id, l.titulo, COUNT(e.id) AS total_emprestimos
FROM livros l
LEFT JOIN emprestimos e ON l.id = e.livro_id
GROUP BY l.id, l.titulo
HAVING COUNT(e.id) > 0;

-- Tentar deletar um livro que tem empréstimos
DELETE FROM livros WHERE id = 1;
-- Erro: FOREIGN KEY constraint failed (se RESTRICT)

-- 5. Criar tabela de teste com CASCADE
CREATE TABLE autores_teste (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    nome TEXT NOT NULL
);

CREATE TABLE livros_teste (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    titulo TEXT NOT NULL,
    autor_id INTEGER,
    FOREIGN KEY (autor_id) REFERENCES autores_teste(id) ON DELETE CASCADE
);

-- Inserir dados de teste
INSERT INTO autores_teste (nome) VALUES ('Autor Teste');
INSERT INTO livros_teste (titulo, autor_id) VALUES ('Livro 1', 1);
INSERT INTO livros_teste (titulo, autor_id) VALUES ('Livro 2', 1);

-- Deletar autor (deve deletar livros também)
DELETE FROM autores_teste WHERE id = 1;

-- Verificar que livros foram deletados
SELECT * FROM livros_teste;
-- Deve retornar 0 linhas (livros foram deletados em cascata)
```

**Resposta Esperada para a Questão de Reflexão**:
- **RESTRICT**:
  - Bloqueia deleção se houver referências
  - Mais seguro, protege dados importantes
  - Exige deleção manual de dependentes
  - **Uso apropriado**: Dados importantes que não devem ser deletados acidentalmente (histórico, transações, registros críticos)
- **CASCADE**:
  - Deleta automaticamente registros relacionados
  - Mais conveniente, mas perigoso
  - Pode causar perda de dados se usado incorretamente
  - **Uso apropriado**: Dados temporários ou que devem ser deletados junto (comentários de post, itens de pedido)
- **Riscos de CASCADE**:
  - **Perda acidental de dados**: Deletar um registro pode deletar muitos outros
  - **Dificuldade de recuperação**: Dados deletados podem ser difíceis de recuperar
  - **Comportamento inesperado**: Pode deletar mais do que o esperado
  - **Dependências ocultas**: Pode não ser óbvio quais dados serão deletados
- **Recomendação**: Use RESTRICT por padrão, CASCADE apenas quando realmente necessário e com muito cuidado

---

### Exercício 9: Modificando Constraints

**Objetivo**: Entender limitações e métodos para modificar constraints.

**Tarefas**:

1. Tente adicionar uma constraint UNIQUE a uma coluna existente na tabela `livros` (se não existir).

2. Tente adicionar uma constraint NOT NULL a uma coluna existente (deve falhar se houver NULLs).

3. Documente as limitações do SQLite para modificar constraints.

4. Crie um script que recria uma tabela com novas constraints (método completo).

**Questão de Reflexão**:
- Por que é importante planejar constraints desde o início? Quais são os desafios de adicionar constraints depois?

**Soluções Esperadas**:

```sql
-- 1. Adicionar UNIQUE a coluna existente
-- Verificar se isbn já tem UNIQUE
.schema livros

-- Se não tiver, adicionar índice único
CREATE UNIQUE INDEX idx_livros_isbn ON livros(isbn);

-- 2. Tentar adicionar NOT NULL (SQLite não suporta diretamente)
-- Primeiro, verificar se há NULLs
SELECT COUNT(*) FROM livros WHERE editora IS NULL;

-- Se houver NULLs, atualizar primeiro
UPDATE livros SET editora = 'Desconhecida' WHERE editora IS NULL;

-- SQLite não suporta ALTER TABLE ... ADD CONSTRAINT NOT NULL diretamente
-- Solução: Recriar tabela

-- 3. Limitações do SQLite
/*
SQLite tem limitações ao modificar constraints:
- Não suporta ALTER TABLE ... ADD CONSTRAINT diretamente
- Não suporta ALTER TABLE ... DROP CONSTRAINT diretamente
- Soluções:
  1. Recriar tabela (método completo)
  2. Criar índices para UNIQUE
  3. Usar triggers para validações complexas
*/

-- 4. Script para recriar tabela com novas constraints
-- Passo 1: Criar nova tabela
CREATE TABLE livros_novo (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    titulo TEXT NOT NULL,
    isbn TEXT UNIQUE,
    ano_publicacao INTEGER CHECK (ano_publicacao > 0),
    editora TEXT NOT NULL,  -- Nova constraint NOT NULL
    autor_id INTEGER,
    categoria_id INTEGER,
    quantidade_disponivel INTEGER DEFAULT 0 CHECK (quantidade_disponivel >= 0),
    FOREIGN KEY (autor_id) REFERENCES autores(id),
    FOREIGN KEY (categoria_id) REFERENCES categorias(id)
);

-- Passo 2: Copiar dados (atualizando NULLs primeiro)
UPDATE livros SET editora = 'Desconhecida' WHERE editora IS NULL;

INSERT INTO livros_novo 
SELECT * FROM livros;

-- Passo 3: Deletar tabela antiga
DROP TABLE livros;

-- Passo 4: Renomear nova tabela
ALTER TABLE livros_novo RENAME TO livros;
```

**Resposta Esperada para a Questão de Reflexão**:
- **Importância de planejar desde o início**:
  - Mais fácil definir constraints na criação
  - Evita problemas de dados existentes
  - Previne necessidade de migrações complexas
  - Garante integridade desde o início
- **Desafios de adicionar depois**:
  - **Dados existentes**: Pode haver dados que violam a nova constraint
  - **Migração complexa**: Pode precisar recriar tabelas
  - **Downtime**: Pode exigir parar o sistema
  - **Risco de perda de dados**: Migrações podem causar perda de dados
  - **Complexidade**: Mais difícil e propenso a erros
- **Recomendação**: Sempre planeje constraints desde o início, pensando em todas as regras de negócio antes de criar tabelas

---

### Exercício 10: Análise de Design

**Objetivo**: Analisar e melhorar design de constraints em um cenário real.

**Cenário**: Você precisa criar um sistema de biblioteca com as seguintes regras de negócio:
- Cada livro tem um ISBN único
- Cada usuário tem um email único
- Quantidade de livros não pode ser negativa
- Data de devolução deve ser após data de empréstimo
- Status de empréstimo deve ser 'ativo', 'devolvido' ou 'atrasado'
- Usuário e livro são obrigatórios em empréstimos

**Tarefas**:

1. Crie as tabelas necessárias com todas as constraints apropriadas.

2. Documente cada constraint e explique por que ela é necessária.

3. Teste cada constraint tentando violá-la.

4. Identifique possíveis melhorias ou constraints adicionais que poderiam ser úteis.

**Questão de Reflexão**:
- Como você balanceia entre ter muitas constraints (mais segurança, menos flexibilidade) vs poucas constraints (mais flexibilidade, menos segurança)? Qual é o ponto ideal?

**Soluções Esperadas**:

```sql
-- 1. Criar tabelas com constraints
PRAGMA foreign_keys = ON;

-- Tabela usuarios
CREATE TABLE usuarios (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    nome TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL,  -- Email único e obrigatório
    telefone TEXT,
    data_cadastro DATE NOT NULL DEFAULT CURRENT_DATE
);

-- Tabela livros
CREATE TABLE livros (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    titulo TEXT NOT NULL,
    isbn TEXT UNIQUE,  -- ISBN único
    ano_publicacao INTEGER CHECK (ano_publicacao > 0 AND ano_publicacao <= 2100),
    editora TEXT,
    autor_id INTEGER,
    categoria_id INTEGER,
    quantidade_disponivel INTEGER DEFAULT 0 CHECK (quantidade_disponivel >= 0),  -- Não pode ser negativo
    FOREIGN KEY (autor_id) REFERENCES autores(id),
    FOREIGN KEY (categoria_id) REFERENCES categorias(id)
);

-- Tabela emprestimos
CREATE TABLE emprestimos (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    livro_id INTEGER NOT NULL,  -- Obrigatório
    usuario_id INTEGER NOT NULL,  -- Obrigatório
    data_emprestimo DATE NOT NULL DEFAULT CURRENT_DATE,
    data_devolucao_prevista DATE NOT NULL,
    data_devolucao_real DATE,
    status TEXT NOT NULL DEFAULT 'ativo' 
        CHECK (status IN ('ativo', 'devolvido', 'atrasado')),  -- Status válido
    FOREIGN KEY (livro_id) REFERENCES livros(id) ON DELETE RESTRICT,
    FOREIGN KEY (usuario_id) REFERENCES usuarios(id) ON DELETE RESTRICT,
    CHECK (data_devolucao_prevista >= data_emprestimo),  -- Data válida
    CHECK (data_devolucao_real IS NULL OR data_devolucao_real >= data_emprestimo)  -- Data válida
);

-- 2. Documentação de constraints
/*
USUARIOS:
- PRIMARY KEY (id): Identifica unicamente cada usuário
- UNIQUE (email): Garante que cada email seja único
- NOT NULL (nome, email): Campos obrigatórios
- DEFAULT (data_cadastro): Data automática

LIVROS:
- PRIMARY KEY (id): Identifica unicamente cada livro
- UNIQUE (isbn): Garante que cada ISBN seja único
- NOT NULL (titulo): Título obrigatório
- CHECK (quantidade_disponivel >= 0): Previne quantidades negativas
- CHECK (ano_publicacao): Valida ano razoável
- FOREIGN KEY (autor_id, categoria_id): Mantém integridade referencial

EMPRESTIMOS:
- PRIMARY KEY (id): Identifica unicamente cada empréstimo
- NOT NULL (livro_id, usuario_id): Referências obrigatórias
- FOREIGN KEY: Mantém integridade referencial
- CHECK (status): Valida status permitido
- CHECK (datas): Valida que datas fazem sentido
*/

-- 3. Testar cada constraint
-- Testar UNIQUE email
INSERT INTO usuarios (nome, email) VALUES ('João', 'joao@email.com');
INSERT INTO usuarios (nome, email) VALUES ('Maria', 'joao@email.com');  -- Erro

-- Testar UNIQUE isbn
INSERT INTO livros (titulo, isbn) VALUES ('Livro 1', '123');
INSERT INTO livros (titulo, isbn) VALUES ('Livro 2', '123');  -- Erro

-- Testar CHECK quantidade
INSERT INTO livros (titulo, quantidade_disponivel) VALUES ('Livro', -5);  -- Erro

-- Testar NOT NULL
INSERT INTO emprestimos (livro_id) VALUES (1);  -- Erro (falta usuario_id)

-- Testar CHECK status
INSERT INTO emprestimos (livro_id, usuario_id, status) 
VALUES (1, 1, 'cancelado');  -- Erro

-- Testar CHECK data
INSERT INTO emprestimos (livro_id, usuario_id, data_emprestimo, data_devolucao_prevista)
VALUES (1, 1, '2024-12-31', '2024-01-01');  -- Erro

-- 4. Possíveis melhorias
/*
- Adicionar CHECK para validar formato de email
- Adicionar CHECK para validar formato de ISBN
- Adicionar CHECK para validar telefone
- Adicionar constraint para garantir que quantidade_disponivel <= quantidade_total
- Adicionar constraint para limitar número de empréstimos por usuário
- Adicionar constraint para validar datas futuras
*/
```

**Resposta Esperada para a Questão de Reflexão**:
- **Muitas constraints (mais segurança, menos flexibilidade)**:
  - **Vantagens**: Dados sempre válidos, menos erros, mais confiável
  - **Desvantagens**: Menos flexível, mudanças difíceis, pode bloquear casos legítimos
- **Poucas constraints (mais flexibilidade, menos segurança)**:
  - **Vantagens**: Mais flexível, mudanças fáceis, aceita mais casos
  - **Desvantagens**: Dados podem ser inválidos, mais erros, menos confiável
- **Ponto ideal**:
  - **Constraints essenciais**: Sempre use (PRIMARY KEY, FOREIGN KEY para relacionamentos críticos, NOT NULL para campos obrigatórios)
  - **Constraints importantes**: Use para regras de negócio fundamentais (UNIQUE para identificadores, CHECK para validações críticas)
  - **Constraints opcionais**: Use com cuidado (CHECK para validações que podem mudar)
  - **Validação na aplicação**: Para regras complexas ou que mudam frequentemente
- **Recomendação**: Use constraints para regras fundamentais e estáveis, validação na aplicação para o resto. É melhor ter mais constraints do que menos, mas não exagere com validações que podem mudar.

---

## Perguntas de Reflexão Finais

### 1. Integridade de Dados

- Por que a integridade de dados é fundamental em um banco de dados? Quais problemas podem ocorrer sem constraints adequadas?
- Como você garante que os dados em seu banco sempre estejam consistentes e válidos?

### 2. Design de Constraints

- Quando você deve usar constraints vs validação na aplicação? Qual é o melhor lugar para cada tipo de validação?
- Como você decide quais campos devem ser NOT NULL? Quais critérios você usa?

### 3. Performance e Manutenção

- Qual o impacto de constraints na performance? Quando constraints podem ser um problema?
- Como você lida com a necessidade de modificar constraints em um banco de dados em produção?

### 4. Boas Práticas

- Quais são as regras de ouro para usar constraints? O que você sempre deve fazer e o que deve evitar?
- Como você documenta constraints para outros desenvolvedores?

---

**Próximo Passo**: Estude **Performance, Boas Práticas e Otimização** (`04-performance-boas-praticas.md`) para aprender como usar constraints de forma eficiente!
