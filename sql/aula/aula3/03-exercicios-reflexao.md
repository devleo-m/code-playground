# Aula 3 - Exercícios e Reflexão

## Exercícios Práticos

### Exercício 1: Criando Tabelas com CREATE TABLE

**Objetivo**: Praticar a criação de tabelas com diferentes tipos de constraints.

**Contexto**: Vamos expandir o sistema da biblioteca criando uma nova funcionalidade de avaliações de livros.

**Tarefas**:

1. Crie uma tabela chamada `avaliacoes` com as seguintes características:
   - `id`: INTEGER, PRIMARY KEY, AUTOINCREMENT
   - `livro_id`: INTEGER, NOT NULL, FOREIGN KEY referenciando `livros(id)`
   - `usuario_id`: INTEGER, NOT NULL, FOREIGN KEY referenciando `usuarios(id)`
   - `nota`: INTEGER, NOT NULL, com CHECK garantindo que a nota está entre 1 e 5
   - `comentario`: TEXT (pode ser NULL)
   - `data_avaliacao`: DATE, DEFAULT CURRENT_DATE, NOT NULL
   - Constraint UNIQUE garantindo que um usuário só pode avaliar um livro uma vez (combinação única de `livro_id` e `usuario_id`)

2. Crie uma tabela chamada `reservas` com:
   - `id`: INTEGER, PRIMARY KEY, AUTOINCREMENT
   - `livro_id`: INTEGER, NOT NULL, FOREIGN KEY referenciando `livros(id)`
   - `usuario_id`: INTEGER, NOT NULL, FOREIGN KEY referenciando `usuarios(id)`
   - `data_reserva`: DATE, DEFAULT CURRENT_DATE, NOT NULL
   - `data_limite`: DATE (pode ser NULL)
   - `status`: TEXT, DEFAULT 'ativa', com CHECK permitindo apenas 'ativa', 'cancelada' ou 'concluida'

3. Verifique se as tabelas foram criadas corretamente usando:
   ```sql
   .schema avaliacoes
   .schema reservas
   ```

**Soluções Esperadas**:

```sql
-- 1. Tabela de avaliações
CREATE TABLE avaliacoes (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    livro_id INTEGER NOT NULL,
    usuario_id INTEGER NOT NULL,
    nota INTEGER NOT NULL CHECK (nota >= 1 AND nota <= 5),
    comentario TEXT,
    data_avaliacao DATE DEFAULT CURRENT_DATE NOT NULL,
    FOREIGN KEY (livro_id) REFERENCES livros(id),
    FOREIGN KEY (usuario_id) REFERENCES usuarios(id),
    UNIQUE(livro_id, usuario_id)
);

-- 2. Tabela de reservas
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

-- 3. Verificar estrutura
.schema avaliacoes
.schema reservas
```

---

### Exercício 2: Modificando Tabelas com ALTER TABLE

**Objetivo**: Praticar modificações de estrutura usando ALTER TABLE.

**Tarefas**:

1. Adicione uma coluna `preco` do tipo REAL à tabela `livros`.

2. Adicione uma coluna `data_cadastro` do tipo DATE com DEFAULT CURRENT_DATE à tabela `livros`.

3. Adicione uma coluna `ativo` do tipo INTEGER com DEFAULT 1 à tabela `livros` (1 = ativo, 0 = inativo).

4. Verifique as mudanças usando:
   ```sql
   PRAGMA table_info(livros);
   ```

5. (Opcional - SQLite 3.25.0+) Renomeie a coluna `quantidade_disponivel` para `estoque` na tabela `livros`.

**Soluções Esperadas**:

```sql
-- 1. Adicionar coluna preco
ALTER TABLE livros
ADD COLUMN preco REAL;

-- 2. Adicionar coluna data_cadastro
ALTER TABLE livros
ADD COLUMN data_cadastro DATE DEFAULT CURRENT_DATE;

-- 3. Adicionar coluna ativo
ALTER TABLE livros
ADD COLUMN ativo INTEGER DEFAULT 1;

-- 4. Verificar estrutura
PRAGMA table_info(livros);

-- 5. Renomear coluna (SQLite 3.25.0+)
ALTER TABLE livros
RENAME COLUMN quantidade_disponivel TO estoque;
```

---

### Exercício 3: Criando Índices para Performance

**Objetivo**: Criar índices para melhorar a performance de consultas.

**Tarefas**:

1. Crie um índice chamado `idx_livros_autor` na coluna `autor_id` da tabela `livros`.

2. Crie um índice chamado `idx_livros_categoria` na coluna `categoria_id` da tabela `livros`.

3. Crie um índice composto chamado `idx_emprestimos_usuario_status` nas colunas `usuario_id` e `status` da tabela `emprestimos`.

4. Crie um índice chamado `idx_avaliacoes_livro` na coluna `livro_id` da tabela `avaliacoes` (se você criou essa tabela no Exercício 1).

5. Liste todos os índices criados usando:
   ```sql
   SELECT name, tbl_name FROM sqlite_master WHERE type = 'index';
   ```

**Soluções Esperadas**:

```sql
-- 1. Índice em autor_id
CREATE INDEX idx_livros_autor ON livros(autor_id);

-- 2. Índice em categoria_id
CREATE INDEX idx_livros_categoria ON livros(categoria_id);

-- 3. Índice composto
CREATE INDEX idx_emprestimos_usuario_status ON emprestimos(usuario_id, status);

-- 4. Índice em avaliacoes
CREATE INDEX idx_avaliacoes_livro ON avaliacoes(livro_id);

-- 5. Listar índices
SELECT name, tbl_name FROM sqlite_master WHERE type = 'index';
```

---

### Exercício 4: Trabalhando com Tabelas Temporárias

**Objetivo**: Praticar criação, uso e remoção de tabelas temporárias.

**Tarefas**:

1. Crie uma tabela temporária chamada `livros_backup` que tenha a mesma estrutura da tabela `livros`, mas sem dados.

2. Insira alguns dados de teste na tabela `livros_backup` (pelo menos 2 registros).

3. Verifique os dados inseridos:
   ```sql
   SELECT * FROM livros_backup;
   ```

4. Limpe todos os dados da tabela `livros_backup` usando DELETE (equivalente a TRUNCATE no SQLite).

5. Verifique que a tabela está vazia, mas a estrutura ainda existe.

6. Remova a tabela `livros_backup` completamente usando DROP TABLE.

7. Tente consultar a tabela novamente. O que acontece?

**Soluções Esperadas**:

```sql
-- 1. Criar tabela temporária com mesma estrutura
CREATE TABLE livros_backup AS 
SELECT * FROM livros WHERE 1=0;  -- WHERE 1=0 garante que não copia dados

-- 2. Inserir dados de teste
INSERT INTO livros_backup (titulo, autor_id, categoria_id, quantidade_disponivel)
VALUES 
    ('Livro Teste 1', 1, 1, 5),
    ('Livro Teste 2', 2, 2, 3);

-- 3. Verificar dados
SELECT * FROM livros_backup;

-- 4. Limpar dados (equivalente a TRUNCATE)
DELETE FROM livros_backup;

-- 5. Verificar estrutura ainda existe
PRAGMA table_info(livros_backup);
-- Deve mostrar a estrutura, mas sem dados

-- 6. Remover tabela
DROP TABLE livros_backup;

-- 7. Tentar consultar (deve dar erro)
SELECT * FROM livros_backup;
-- Erro: no such table: livros_backup
```

---

### Exercício 5: Análise de Constraints

**Objetivo**: Entender o comportamento de diferentes constraints através de testes práticos.

**Tarefas**:

1. Tente inserir um registro na tabela `avaliacoes` (criada no Exercício 1) com `nota = 6`. O que acontece? Por quê?

2. Tente inserir dois registros na tabela `avaliacoes` com o mesmo `livro_id` e `usuario_id`, mas notas diferentes. O que acontece? Por quê?

3. Tente inserir um registro na tabela `avaliacoes` com `livro_id = 999` (assumindo que não existe um livro com id 999). O que acontece? Por quê?

4. Tente inserir um registro na tabela `avaliacoes` sem especificar a `nota`. O que acontece? Por quê?

5. Tente inserir um registro na tabela `avaliacoes` sem especificar a `data_avaliacao`. O que acontece? Por quê?

**Soluções Esperadas**:

```sql
-- 1. Tentar inserir nota = 6 (deve falhar)
INSERT INTO avaliacoes (livro_id, usuario_id, nota)
VALUES (1, 1, 6);
-- Erro: CHECK constraint failed: nota >= 1 AND nota <= 5
-- Porque: A constraint CHECK valida que a nota deve estar entre 1 e 5

-- 2. Tentar inserir avaliação duplicada (deve falhar)
INSERT INTO avaliacoes (livro_id, usuario_id, nota)
VALUES (1, 1, 5);
INSERT INTO avaliacoes (livro_id, usuario_id, nota)
VALUES (1, 1, 4);  -- Mesmo livro_id e usuario_id
-- Erro: UNIQUE constraint failed: avaliacoes.livro_id, avaliacoes.usuario_id
-- Porque: A constraint UNIQUE garante que um usuário só pode avaliar um livro uma vez

-- 3. Tentar inserir com livro_id inexistente (deve falhar)
INSERT INTO avaliacoes (livro_id, usuario_id, nota)
VALUES (999, 1, 5);
-- Erro: FOREIGN KEY constraint failed
-- Porque: A FOREIGN KEY garante que livro_id deve existir na tabela livros

-- 4. Tentar inserir sem nota (deve falhar)
INSERT INTO avaliacoes (livro_id, usuario_id)
VALUES (1, 1);
-- Erro: NOT NULL constraint failed: avaliacoes.nota
-- Porque: A constraint NOT NULL exige que o campo seja preenchido

-- 5. Tentar inserir sem data_avaliacao (deve funcionar)
INSERT INTO avaliacoes (livro_id, usuario_id, nota)
VALUES (1, 1, 5);
-- Deve funcionar porque data_avaliacao tem DEFAULT CURRENT_DATE
-- O valor será preenchido automaticamente com a data atual
```

---

## Perguntas de Reflexão

### Reflexão 1: Impacto de ALTER TABLE em Produção

**Situação**: Você precisa adicionar uma nova coluna `preco` à tabela `livros` em um banco de dados de produção que tem 1 milhão de registros.

**Perguntas**:

1. Qual é o impacto de executar `ALTER TABLE livros ADD COLUMN preco REAL;` em uma tabela com 1 milhão de registros?

2. O que acontece com os registros existentes quando você adiciona uma nova coluna? Eles são modificados?

3. Se a nova coluna tiver `NOT NULL` sem `DEFAULT`, o que acontece? Por quê isso pode ser problemático?

4. Qual seria uma estratégia segura para adicionar uma coluna `NOT NULL` com valor padrão em uma tabela grande?

**Respostas Esperadas**:

1. **Impacto**: Em SQLite, adicionar uma coluna é relativamente rápido porque SQLite não precisa modificar os registros existentes - ele apenas atualiza o schema. Em outros SGBDs (PostgreSQL, MySQL), pode ser mais lento e bloquear a tabela durante a operação.

2. **Registros existentes**: Os registros existentes não são modificados fisicamente. A nova coluna é adicionada ao schema, e para registros antigos, o valor será NULL (ou o DEFAULT, se especificado). O banco de dados "sabe" que esses registros não têm valor para a nova coluna.

3. **NOT NULL sem DEFAULT**: Isso causaria um erro, porque os 1 milhão de registros existentes não teriam valor para a nova coluna, violando a constraint NOT NULL. É por isso que SQLite não permite adicionar colunas NOT NULL sem DEFAULT em tabelas que já têm dados.

4. **Estratégia segura**:
   - Adicionar a coluna como NULL primeiro
   - Atualizar todos os registros existentes com o valor padrão
   - Depois, se necessário, alterar para NOT NULL (isso pode não ser possível no SQLite, mas funciona em outros SGBDs)

---

### Reflexão 2: Quando Usar DROP TABLE vs DELETE FROM

**Situação**: Você tem uma tabela `logs_temporarios` que armazena logs de sistema. Periodicamente, você precisa limpar os dados antigos.

**Perguntas**:

1. Qual é a diferença prática entre `DROP TABLE logs_temporarios` e `DELETE FROM logs_temporarios`?

2. Se você quer manter a estrutura da tabela mas remover todos os dados, qual comando você deve usar? Por quê?

3. Se você quer remover a tabela completamente porque não será mais usada, qual comando você deve usar?

4. Em termos de performance, qual é mais rápido: DROP TABLE ou DELETE FROM? Por quê?

5. Qual comando pode ser revertido em uma transação? Por quê isso é importante?

**Respostas Esperadas**:

1. **Diferença**:
   - `DROP TABLE`: Remove a tabela inteira (estrutura + dados). A tabela deixa de existir.
   - `DELETE FROM`: Remove apenas os dados. A estrutura da tabela permanece.

2. **Manter estrutura**: Use `DELETE FROM logs_temporarios`. Isso remove todos os dados mas mantém a estrutura, permitindo que novos dados sejam inseridos depois.

3. **Remover completamente**: Use `DROP TABLE logs_temporarios`. Isso remove tudo, indicando que a tabela não será mais usada.

4. **Performance**: `DROP TABLE` é geralmente mais rápido porque apenas remove metadados e marca o espaço como disponível. `DELETE FROM` precisa processar cada registro individualmente, gerando logs e verificando constraints.

5. **Reversibilidade**: `DELETE FROM` pode ser revertido em uma transação (ROLLBACK). `DROP TABLE` geralmente é auto-commit e não pode ser revertido facilmente. Isso é importante para segurança - se você cometer um erro, pode desfazer com DELETE, mas não com DROP.

---

### Reflexão 3: Estratégia de Índices

**Situação**: Você está criando um sistema de biblioteca e precisa decidir onde criar índices.

**Perguntas**:

1. Por que faz sentido criar um índice na coluna `autor_id` da tabela `livros`?

2. Se você criar índices em TODAS as colunas da tabela `livros`, isso seria uma boa ideia? Por quê?

3. Quando você faz um `INSERT` em uma tabela que tem muitos índices, o que acontece? Isso é bom ou ruim?

4. Como você decide quais colunas devem ter índices? Quais critérios você usa?

5. Um índice composto (múltiplas colunas) é sempre melhor que índices separados? Quando usar cada um?

**Respostas Esperadas**:

1. **Índice em autor_id**: Faz sentido porque `autor_id` é frequentemente usado em JOINs (para buscar livros de um autor) e em WHERE (para filtrar por autor). O índice acelera essas operações significativamente.

2. **Índices em todas as colunas**: Não é uma boa ideia porque:
   - Cada INSERT/UPDATE/DELETE precisa atualizar todos os índices, tornando essas operações muito lentas
   - Índices ocupam espaço em disco
   - Nem todas as colunas são usadas em buscas frequentes

3. **INSERT com muitos índices**: Cada índice precisa ser atualizado quando um novo registro é inserido. Com muitos índices, isso pode tornar INSERTs muito lentos. É um trade-off: índices melhoram leitura, mas podem atrasar escrita.

4. **Critérios para criar índices**:
   - Colunas usadas frequentemente em WHERE
   - Colunas usadas em JOINs (especialmente chaves estrangeiras)
   - Colunas usadas em ORDER BY
   - Colunas com muitos valores únicos (alta seletividade)
   - Colunas consultadas frequentemente juntas (índice composto)

5. **Índice composto vs separados**:
   - **Índice composto**: Use quando você sempre consulta as colunas juntas (ex: `usuario_id` E `status` juntos)
   - **Índices separados**: Use quando você consulta as colunas independentemente
   - Índice composto `(usuario_id, status)` é útil para `WHERE usuario_id = X AND status = Y`, mas não ajuda muito em `WHERE status = Y` sozinho

---

### Reflexão 4: Constraints e Integridade de Dados

**Situação**: Você está projetando o schema de um banco de dados para um sistema de biblioteca.

**Perguntas**:

1. Por que é importante ter uma PRIMARY KEY em cada tabela?

2. Qual é a diferença entre uma FOREIGN KEY e uma referência simples (apenas colocar o ID sem constraint)?

3. Se você não usar a constraint CHECK na coluna `nota` da tabela `avaliacoes`, o que pode acontecer?

4. Por que a constraint UNIQUE é importante para a combinação `(livro_id, usuario_id)` na tabela `avaliacoes`?

5. Constraints melhoram ou pioram a performance? Por quê?

**Respostas Esperadas**:

1. **PRIMARY KEY importante porque**:
   - Garante unicidade (cada registro é único)
   - Cria índice automaticamente (melhora performance)
   - Permite referências de outras tabelas (FOREIGN KEY)
   - É a forma padrão de identificar registros

2. **FOREIGN KEY vs referência simples**:
   - **FOREIGN KEY**: O banco garante que o valor existe na tabela referenciada. Se você tentar inserir um `livro_id = 999` que não existe, o banco rejeita.
   - **Referência simples**: Apenas um número. O banco não valida se existe. Você pode ter dados inconsistentes (órfãos).

3. **Sem CHECK na nota**: Sem a constraint, você poderia inserir valores inválidos como `nota = 0`, `nota = 10`, `nota = -5`, etc. A constraint CHECK garante que apenas valores válidos (1-5) sejam aceitos.

4. **UNIQUE em (livro_id, usuario_id)**: Garante que um usuário só pode avaliar um livro uma vez. Sem isso, o mesmo usuário poderia avaliar o mesmo livro múltiplas vezes, o que geralmente não faz sentido no contexto de um sistema de avaliações.

5. **Performance de constraints**:
   - **Melhoram**: PRIMARY KEY e UNIQUE criam índices automaticamente, melhorando buscas
   - **Podem atrasar**: FOREIGN KEY e CHECK precisam validar dados, o que pode atrasar INSERT/UPDATE
   - **Trade-off**: A pequena perda de performance em escrita é compensada pela garantia de integridade e prevenção de bugs

---

### Reflexão 5: Planejamento de Schema

**Situação**: Você precisa criar um novo sistema e está planejando o schema do banco de dados.

**Perguntas**:

1. Por que é importante planejar o schema antes de começar a criar tabelas?

2. Se você esquecer de adicionar uma coluna importante e já tiver dados na tabela, qual é a melhor estratégia?

3. Como você decide se uma informação deve ser uma coluna em uma tabela existente ou uma nova tabela?

4. Qual é o impacto de fazer muitas mudanças de schema (ALTER TABLE) em um sistema em produção?

5. Como você documenta mudanças de schema para que outros desenvolvedores saibam o que mudou?

**Respostas Esperadas**:

1. **Planejamento importante porque**:
   - Evita refatorações custosas depois
   - Garante que relacionamentos estão corretos desde o início
   - Previne problemas de integridade de dados
   - Facilita manutenção futura

2. **Estratégia para adicionar coluna esquecida**:
   - Adicionar a coluna com ALTER TABLE
   - Se for NOT NULL, adicionar como NULL primeiro, popular dados, depois alterar
   - Se necessário, criar script de migração para popular dados existentes
   - Testar em ambiente de desenvolvimento primeiro

3. **Decisão coluna vs nova tabela**:
   - **Coluna na tabela existente**: Quando a informação é diretamente relacionada e sempre presente (ex: `preco` em `livros`)
   - **Nova tabela**: Quando há relacionamento 1-N ou N-N (ex: múltiplas avaliações por livro), ou quando a informação é opcional e complexa

4. **Impacto de muitas mudanças**:
   - Pode causar downtime se a tabela for bloqueada
   - Pode ser lento em tabelas grandes
   - Pode quebrar aplicações que dependem da estrutura antiga
   - Requer coordenação e testes cuidadosos

5. **Documentação de mudanças**:
   - Usar sistema de versionamento (migrations)
   - Manter changelog de schema
   - Documentar razão da mudança
   - Incluir scripts de rollback quando possível
   - Comunicar mudanças à equipe

---

## Conclusão dos Exercícios

Após completar estes exercícios e refletir sobre as perguntas, você deve ter:

- ✅ Praticado criação de tabelas com diferentes constraints
- ✅ Entendido como modificar estruturas existentes
- ✅ Compreendido quando e como criar índices
- ✅ Experimentado operações destrutivas (DROP, DELETE)
- ✅ Refletido sobre impacto e performance de operações DDL
- ✅ Pensado criticamente sobre design de schema

**Próximo Passo**: Estude a seção de Performance, Boas Práticas e Otimização para aprofundar seu entendimento sobre como usar DDL de forma eficiente e segura.

---

**⚠️ Lembrete**: Sempre faça backup antes de executar comandos DDL destrutivos, especialmente DROP TABLE!


