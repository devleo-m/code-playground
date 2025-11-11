# **Aula 5: Conceitos de Alto Nível de Banco de Dados**

## 🎯 Objetivo da Aula

Compreender os conceitos avançados que sustentam o funcionamento interno do PostgreSQL: ACID, MVCC, Transações, WAL e Processamento de Consultas. Estes conceitos são fundamentais para criar sistemas robustos, eficientes e escaláveis.

---

## 📚 Introdução: Conceitos de Alto Nível

**Conceitos de alto nível** são os princípios fundamentais que regem:

- **Design**: Como estruturar o banco de dados
- **Implementação**: Como o PostgreSQL funciona internamente
- **Gerenciamento**: Como manter o sistema funcionando corretamente

Estes conceitos garantem:

- ✅ **Confiabilidade**: Dados não são perdidos
- ✅ **Consistência**: Dados sempre corretos
- ✅ **Performance**: Sistema rápido mesmo com múltiplos usuários
- ✅ **Escalabilidade**: Cresce conforme necessidade

---

## 🔐 1. ACID: As 4 Propriedades Fundamentais

**ACID** é um acrônimo que representa as **4 propriedades** que garantem transações confiáveis em bancos de dados relacionais.

### A - Atomicidade (Atomicity)

**Definição:** Uma transação é uma unidade **indivisível** de trabalho. Ou **todas** as operações são executadas, ou **nenhuma** é.

**Analogia:** Como um átomo (que não pode ser dividido), uma transação não pode ser executada "pela metade".

#### Exemplo Prático: Transferência Bancária

```sql
-- Transferir R$ 100 da Conta A para Conta B
BEGIN;  -- Inicia transação

-- Operação 1: Debitar da Conta A
UPDATE contas SET saldo = saldo - 100 WHERE id = 'A';

-- Operação 2: Creditar na Conta B
UPDATE contas SET saldo = saldo + 100 WHERE id = 'B';

COMMIT;  -- Confirma transação
```

**Cenários:**

✅ **Sucesso:** Ambas as operações executam → Dinheiro transferido
❌ **Falha:** Se qualquer operação falhar → **ROLLBACK** automático → Nada acontece

```sql
BEGIN;

UPDATE contas SET saldo = saldo - 100 WHERE id = 'A';  -- ✅ Sucesso

-- 💥 Sistema cai aqui!

UPDATE contas SET saldo = saldo + 100 WHERE id = 'B';  -- Não executa

-- PostgreSQL detecta falha e faz ROLLBACK automático
-- Conta A volta ao estado original (dinheiro não some!)
```

---

### C - Consistência (Consistency)

**Definição:** Uma transação leva o banco de dados de um **estado consistente** para outro **estado consistente**, respeitando todas as regras (constraints, triggers, etc).

**Analogia:** Se há regra "saldo não pode ser negativo", o banco **nunca** vai permitir que isso aconteça.

#### Exemplo Prático: Regras de Negócio

```sql
CREATE TABLE contas (
    id CHAR(1) PRIMARY KEY,
    saldo DECIMAL(10, 2) CHECK (saldo >= 0)  -- Regra: saldo não pode ser negativo
);

INSERT INTO contas VALUES ('A', 500.00), ('B', 300.00);

-- Tentativa de transferência que violaria consistência
BEGIN;

UPDATE contas SET saldo = saldo - 600 WHERE id = 'A';  -- Saldo ficaria -100!
-- ❌ ERRO! Viola constraint CHECK (saldo >= 0)
-- Transação é abortada automaticamente

ROLLBACK;  -- Volta ao estado consistente
```

**Consistência garante:**

- ✅ Constraints são respeitadas (CHECK, NOT NULL, UNIQUE, FK)
- ✅ Triggers são executados
- ✅ Regras de negócio são aplicadas
- ✅ Banco nunca fica em estado inválido

---

### I - Isolamento (Isolation)

**Definição:** Transações concorrentes (executando ao mesmo tempo) são **isoladas** umas das outras. Cada transação opera como se fosse a única no sistema.

**Analogia:** Você está editando um documento. Outra pessoa também está editando. Vocês não veem as mudanças um do outro até que salvem (COMMIT).

#### Exemplo Prático: Duas Transações Simultâneas

```sql
-- Transação 1 (Usuário A)
BEGIN;
SELECT saldo FROM contas WHERE id = 'A';  -- Vê: 500.00
UPDATE contas SET saldo = saldo - 100 WHERE id = 'A';
-- Ainda não fez COMMIT!

-- Transação 2 (Usuário B) - executando ao mesmo tempo
BEGIN;
SELECT saldo FROM contas WHERE id = 'A';  -- Vê: 500.00 (não vê mudança de A!)
UPDATE contas SET saldo = saldo - 50 WHERE id = 'A';
COMMIT;  -- Usuário B confirma primeiro

-- Transação 1 continua
COMMIT;  -- Usuário A confirma depois
```

**Níveis de Isolamento** (do mais fraco ao mais forte):

1. **READ UNCOMMITTED** (não suportado no PostgreSQL)

   - Lê dados não confirmados (dirty reads)

2. **READ COMMITTED** (padrão no PostgreSQL)

   - Lê apenas dados confirmados
   - Cada query vê snapshot no momento da execução

3. **REPEATABLE READ**

   - Vê snapshot no início da transação
   - Previne leituras não repetíveis

4. **SERIALIZABLE** (mais forte)
   - Transações executam como se fossem seriais (uma após a outra)
   - Previne anomalias de serialização

```sql
-- Definir nível de isolamento
BEGIN TRANSACTION ISOLATION LEVEL REPEATABLE READ;
-- Suas queries aqui
COMMIT;

-- Ou para toda sessão
SET SESSION CHARACTERISTICS AS TRANSACTION ISOLATION LEVEL SERIALIZABLE;
```

---

### D - Durabilidade (Durability)

**Definição:** Uma vez que uma transação é **confirmada (COMMIT)**, as mudanças são **permanentes**, mesmo que o sistema caia imediatamente depois.

**Analogia:** Quando você salva um documento e vê "Salvo com sucesso", mesmo que o computador desligue, o documento está salvo.

#### Como PostgreSQL Garante Durabilidade?

```sql
BEGIN;

UPDATE contas SET saldo = saldo + 1000 WHERE id = 'A';

COMMIT;  -- Neste momento, PostgreSQL:
         -- 1. Escreve mudanças no WAL (Write-Ahead Log)
         -- 2. Força escrita no disco (fsync)
         -- 3. Só então retorna "COMMIT" para você

-- 💥 Sistema pode cair AGORA!
-- ✅ Dados estão salvos! WAL garante recuperação
```

**Mecanismos de Durabilidade:**

- **WAL (Write-Ahead Log)**: Registra mudanças antes de aplicar
- **fsync**: Força escrita física no disco
- **Checkpoints**: Pontos de recuperação
- **Replicação**: Cópias em outros servidores

---

## 🔄 2. MVCC: Controle de Concorrência Multiversão

**MVCC** (Multi-Version Concurrency Control) é a técnica que o PostgreSQL usa para permitir **múltiplas transações simultâneas** sem bloqueios excessivos.

### Como Funciona?

**Ideia central:** Em vez de bloquear dados, o PostgreSQL mantém **múltiplas versões** de cada linha.

```
Linha Original:
┌────┬──────┬────────┐
│ ID │ Nome │ Saldo  │
├────┼──────┼────────┤
│ 1  │ João │ 500.00 │
└────┴──────┴────────┘

Transação A altera (mas não fez COMMIT):
┌────┬──────┬────────┬─────────────┐
│ ID │ Nome │ Saldo  │ Visível para│
├────┼──────┼────────┼─────────────┤
│ 1  │ João │ 500.00 │ Trans antigas│ ← Versão antiga
│ 1  │ João │ 400.00 │ Trans A     │ ← Versão nova (só Trans A vê)
└────┴──────┴────────┴─────────────┘

Outras transações continuam vendo versão antiga!
```

### Vantagens do MVCC

#### 1. Leitores Não Bloqueiam Escritores

```sql
-- Transação 1: Lendo dados
BEGIN;
SELECT * FROM produtos WHERE categoria = 'Eletrônicos';
-- Pode demorar 10 segundos...

-- Transação 2: Escrevendo dados (ao mesmo tempo!)
BEGIN;
UPDATE produtos SET preco = preco * 1.1 WHERE categoria = 'Eletrônicos';
COMMIT;  -- ✅ Não precisa esperar Transação 1!

-- Transação 1 continua vendo versão antiga
-- Não há bloqueio!
```

#### 2. Escritores Não Bloqueiam Leitores

```sql
-- Transação 1: Atualizando
BEGIN;
UPDATE produtos SET estoque = estoque - 1 WHERE id = 1;
-- Ainda não fez COMMIT

-- Transação 2: Lendo (ao mesmo tempo!)
BEGIN;
SELECT * FROM produtos WHERE id = 1;
-- ✅ Vê versão antiga, não precisa esperar!
COMMIT;
```

### Desvantagens do MVCC

#### 1. Espaço em Disco (Bloat)

Versões antigas ocupam espaço até serem limpas.

```sql
-- Limpeza manual
VACUUM produtos;

-- Limpeza completa (mais agressiva)
VACUUM FULL produtos;

-- Autovacuum (automático)
-- PostgreSQL tem processo que limpa automaticamente
```

#### 2. Transaction ID Wraparound

PostgreSQL usa IDs de transação (32 bits). Após ~2 bilhões de transações, precisa fazer wraparound.

```sql
-- Verificar idade das transações
SELECT datname, age(datfrozenxid)
FROM pg_database
ORDER BY age(datfrozenxid) DESC;

-- Prevenir wraparound
VACUUM FREEZE;
```

---

## 💼 3. Transações

**Transação** é uma sequência de operações tratadas como uma **unidade única** de trabalho.

### Comandos Básicos

```sql
-- Iniciar transação
BEGIN;
-- ou
START TRANSACTION;

-- Executar operações
INSERT INTO ...;
UPDATE ...;
DELETE FROM ...;

-- Confirmar (tornar permanente)
COMMIT;

-- ou Cancelar (desfazer tudo)
ROLLBACK;
```

### Exemplo Completo: Sistema de Pedidos

```sql
BEGIN;

-- 1. Criar pedido
INSERT INTO pedidos (cliente_id, total)
VALUES (1, 150.00)
RETURNING id INTO pedido_id;

-- 2. Adicionar itens do pedido
INSERT INTO itens_pedido (pedido_id, produto_id, quantidade, preco)
VALUES
    (pedido_id, 10, 2, 50.00),
    (pedido_id, 20, 1, 50.00);

-- 3. Atualizar estoque
UPDATE produtos SET estoque = estoque - 2 WHERE id = 10;
UPDATE produtos SET estoque = estoque - 1 WHERE id = 20;

-- 4. Registrar pagamento
INSERT INTO pagamentos (pedido_id, valor, metodo)
VALUES (pedido_id, 150.00, 'cartao');

-- Se tudo deu certo:
COMMIT;

-- Se algo deu errado em qualquer etapa:
-- ROLLBACK;  -- Desfaz TUDO!
```

### Savepoints (Pontos de Salvamento)

Permitem **desfazer parcialmente** uma transação.

```sql
BEGIN;

-- Operação 1
INSERT INTO logs (mensagem) VALUES ('Início');

SAVEPOINT sp1;  -- Ponto de salvamento

-- Operação 2
UPDATE produtos SET preco = preco * 1.1;

SAVEPOINT sp2;  -- Outro ponto

-- Operação 3
DELETE FROM produtos WHERE estoque = 0;

-- Ops, não quero deletar!
ROLLBACK TO sp2;  -- Volta para sp2 (desfaz DELETE, mantém UPDATE)

-- Ops, não quero UPDATE também!
ROLLBACK TO sp1;  -- Volta para sp1 (desfaz UPDATE, mantém INSERT)

COMMIT;  -- Confirma apenas INSERT
```

### Transações Implícitas vs Explícitas

```sql
-- Implícita (autocommit - padrão)
UPDATE produtos SET preco = 100 WHERE id = 1;
-- Automaticamente faz COMMIT após executar

-- Explícita
BEGIN;
UPDATE produtos SET preco = 100 WHERE id = 1;
UPDATE produtos SET preco = 200 WHERE id = 2;
COMMIT;  -- Ambas confirmadas juntas
```

### Bloqueios (Locks)

Transações podem adquirir bloqueios para prevenir conflitos.

```sql
-- Bloqueio de leitura (outros podem ler, mas não alterar)
BEGIN;
SELECT * FROM produtos WHERE id = 1 FOR SHARE;
-- Outros podem fazer SELECT, mas não UPDATE/DELETE
COMMIT;

-- Bloqueio de escrita (ninguém pode ler ou alterar)
BEGIN;
SELECT * FROM produtos WHERE id = 1 FOR UPDATE;
-- Outros não podem SELECT FOR UPDATE nem UPDATE/DELETE
UPDATE produtos SET estoque = estoque - 1 WHERE id = 1;
COMMIT;
```

---

## 📝 4. WAL: Write-Ahead Log

**WAL** (Write-Ahead Log) é o mecanismo que garante **durabilidade** e permite **recuperação** após falhas.

### Como Funciona?

**Princípio:** Escrever mudanças no **log** antes de escrever nos **arquivos de dados**.

```
Fluxo sem WAL (perigoso):
┌─────────────────────────────────────┐
│ 1. Aplicação faz UPDATE             │
│ 2. PostgreSQL altera dados na RAM   │
│ 3. Escreve no disco                 │
│    💥 Sistema cai aqui!              │
│    ❌ Dados perdidos!                │
└─────────────────────────────────────┘

Fluxo com WAL (seguro):
┌─────────────────────────────────────┐
│ 1. Aplicação faz UPDATE             │
│ 2. PostgreSQL altera dados na RAM   │
│ 3. Escreve mudança no WAL (log)    │ ← PRIMEIRO!
│ 4. Força WAL para disco (fsync)    │
│ 5. Retorna "COMMIT" para aplicação │
│    💥 Sistema pode cair aqui!        │
│    ✅ Dados estão no WAL!            │
│ 6. Mais tarde, escreve no arquivo  │
└─────────────────────────────────────┘
```

### Estrutura do WAL

```
WAL é sequência de registros:
┌──────────────────────────────────────┐
│ Registro 1: INSERT INTO produtos ... │
│ Registro 2: UPDATE contas SET ...    │
│ Registro 3: DELETE FROM logs ...     │
│ Registro 4: COMMIT                   │
│ ...                                  │
└──────────────────────────────────────┘

Arquivos WAL:
/var/lib/postgresql/data/pg_wal/
├─ 000000010000000000000001
├─ 000000010000000000000002
├─ 000000010000000000000003
└─ ...
```

### Configurações do WAL

```sql
-- Ver configurações atuais
SHOW wal_level;          -- minimal, replica, logical
SHOW fsync;              -- on (garante escrita no disco)
SHOW synchronous_commit; -- on, off, local, remote_write, remote_apply

-- Configurar (em postgresql.conf)
wal_level = replica              -- Nível de informação no WAL
max_wal_size = 1GB              -- Tamanho máximo antes de checkpoint
min_wal_size = 80MB             -- Tamanho mínimo a manter
wal_compression = on            -- Comprimir WAL
```

### Recuperação com WAL

Quando PostgreSQL inicia após crash:

```
1. Lê último checkpoint
2. Aplica registros do WAL desde o checkpoint
3. Refaz (REDO) transações confirmadas
4. Desfaz (UNDO) transações não confirmadas
5. Banco volta ao estado consistente!
```

### Arquivamento do WAL

```sql
-- Configurar arquivamento (em postgresql.conf)
archive_mode = on
archive_command = 'cp %p /archive/%f'

-- Arquivos WAL são copiados para /archive/
-- Permite Point-In-Time Recovery (PITR)
```

---

## 🔍 5. Processamento de Consultas

**Processamento de consultas** é como o PostgreSQL transforma seu SQL em resultados.

### Etapas do Processamento

```
SQL Query
    ↓
┌─────────────────────────────────────┐
│ 1. PARSING (Análise Sintática)     │
│    - Verifica sintaxe               │
│    - Cria árvore de análise         │
└─────────────────────────────────────┘
    ↓
┌─────────────────────────────────────┐
│ 2. REWRITE (Reescrita)              │
│    - Aplica views                   │
│    - Aplica rules                   │
└─────────────────────────────────────┘
    ↓
┌─────────────────────────────────────┐
│ 3. PLANNING (Planejamento)          │
│    - Gera planos de execução        │
│    - Estima custos                  │
│    - Escolhe melhor plano           │
└─────────────────────────────────────┘
    ↓
┌─────────────────────────────────────┐
│ 4. EXECUTION (Execução)             │
│    - Executa plano escolhido        │
│    - Retorna resultados             │
└─────────────────────────────────────┘
```

### 1. Parsing (Análise)

```sql
-- Query
SELECT nome, preco
FROM produtos
WHERE categoria = 'Eletrônicos';

-- PostgreSQL verifica:
-- ✅ Sintaxe correta?
-- ✅ Tabela 'produtos' existe?
-- ✅ Colunas 'nome', 'preco', 'categoria' existem?
-- ✅ Você tem permissão?
```

### 2. Rewrite (Reescrita)

```sql
-- Se você tem uma view:
CREATE VIEW produtos_eletronicos AS
SELECT * FROM produtos WHERE categoria = 'Eletrônicos';

-- Query original:
SELECT nome FROM produtos_eletronicos;

-- PostgreSQL reescreve para:
SELECT nome FROM produtos WHERE categoria = 'Eletrônicos';
```

### 3. Planning (Planejamento)

PostgreSQL gera **múltiplos planos** e escolhe o mais eficiente.

```sql
-- Query
SELECT * FROM pedidos WHERE cliente_id = 100;

-- Plano A: Sequential Scan (varrer toda tabela)
-- Custo: 1000 unidades

-- Plano B: Index Scan (usar índice em cliente_id)
-- Custo: 10 unidades

-- PostgreSQL escolhe Plano B! ✅
```

#### Ver Plano de Execução

```sql
-- Ver plano (sem executar)
EXPLAIN
SELECT * FROM produtos WHERE preco > 100;

-- Ver plano E executar (mostra tempo real)
EXPLAIN ANALYZE
SELECT * FROM produtos WHERE preco > 100;

-- Resultado:
/*
Seq Scan on produtos  (cost=0.00..35.50 rows=10 width=100) (actual time=0.012..0.234 rows=8 loops=1)
  Filter: (preco > 100)
  Rows Removed by Filter: 992
Planning Time: 0.123 ms
Execution Time: 0.456 ms
*/
```

#### Tipos de Scan

**Sequential Scan** (Varredura Sequencial):

```sql
-- Lê tabela inteira, linha por linha
EXPLAIN SELECT * FROM produtos;
-- Seq Scan on produtos
```

**Index Scan** (Varredura de Índice):

```sql
-- Usa índice para encontrar linhas
CREATE INDEX idx_preco ON produtos(preco);
EXPLAIN SELECT * FROM produtos WHERE preco = 100;
-- Index Scan using idx_preco on produtos
```

**Index Only Scan** (Só Índice):

```sql
-- Dados estão no índice, não precisa acessar tabela
CREATE INDEX idx_preco ON produtos(preco);
EXPLAIN SELECT preco FROM produtos WHERE preco > 100;
-- Index Only Scan using idx_preco on produtos
```

**Bitmap Scan** (Varredura de Bitmap):

```sql
-- Combina múltiplos índices
EXPLAIN SELECT * FROM produtos
WHERE preco > 100 AND categoria = 'Eletrônicos';
-- Bitmap Heap Scan on produtos
```

### 4. Execution (Execução)

PostgreSQL executa o plano escolhido e retorna resultados.

```sql
-- Executor percorre plano de baixo para cima
-- Cada nó processa dados e passa para nó acima
```

---

## 📊 Resumo dos Conceitos

| Conceito             | Definição                               | Benefício                         |
| -------------------- | --------------------------------------- | --------------------------------- |
| **ACID**             | 4 propriedades de transações confiáveis | Dados consistentes e duráveis     |
| **Atomicidade**      | Tudo ou nada                            | Previne estados parciais          |
| **Consistência**     | Regras sempre respeitadas               | Dados sempre válidos              |
| **Isolamento**       | Transações não interferem               | Concorrência segura               |
| **Durabilidade**     | Mudanças permanentes após COMMIT        | Dados não são perdidos            |
| **MVCC**             | Múltiplas versões de dados              | Leitores não bloqueiam escritores |
| **Transações**       | Unidade de trabalho                     | Agrupa operações                  |
| **WAL**              | Log de mudanças                         | Recuperação após falhas           |
| **Query Processing** | Transformar SQL em resultados           | Performance otimizada             |

---

## 🎓 Conclusão

Nesta aula você aprendeu os conceitos avançados que fazem o PostgreSQL funcionar:

1. **ACID**: Garantias fundamentais de transações
2. **MVCC**: Concorrência sem bloqueios excessivos
3. **Transações**: Como agrupar operações
4. **WAL**: Durabilidade e recuperação
5. **Processamento de Consultas**: Como SQL vira resultados

Estes conceitos são a base para entender performance, tuning e troubleshooting!

---

## 🔑 Conceitos para Memorizar

- **ACID**: Atomicidade, Consistência, Isolamento, Durabilidade
- **MVCC**: Múltiplas versões = Leitores não bloqueiam escritores
- **BEGIN/COMMIT/ROLLBACK**: Controle de transações
- **SAVEPOINT**: Desfazer parcialmente
- **WAL**: Escreve log antes dos dados (durabilidade)
- **EXPLAIN**: Ver plano de execução
- **Sequential Scan vs Index Scan**: Varredura completa vs usar índice
