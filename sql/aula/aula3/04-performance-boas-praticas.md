# Aula 3 - Performance, Boas Práticas e Otimização

## Introdução: DDL em Produção

Operações DDL (Data Definition Language) são fundamentais para gerenciar a estrutura do banco de dados, mas podem ter impactos significativos em sistemas em produção. Entender como executar essas operações de forma eficiente e segura é crucial para qualquer desenvolvedor ou administrador de banco de dados.

**Regra de Ouro**: DDL em produção requer planejamento, testes e backup. Nunca execute comandos DDL destrutivos sem essas precauções.

---

## 1. Performance: Impacto de Operações DDL

### 1.1 CREATE TABLE - Impacto e Considerações

#### Performance do CREATE TABLE

Criar uma tabela é geralmente uma operação rápida, mas há considerações importantes:

```sql
-- Operação rápida (cria apenas estrutura)
CREATE TABLE nova_tabela (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    nome TEXT
);
```

**Fatores que afetam a performance:**

1. **Número de colunas**: Mais colunas = mais tempo (mínimo)
2. **Constraints complexas**: CHECK, FOREIGN KEY podem adicionar validação
3. **Índices criados automaticamente**: PRIMARY KEY e UNIQUE criam índices automaticamente

#### Boas Práticas para CREATE TABLE

```sql
-- ✅ BOM: Tabela bem planejada com constraints necessárias
CREATE TABLE livros (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    titulo TEXT NOT NULL,
    isbn TEXT UNIQUE,
    autor_id INTEGER NOT NULL,
    FOREIGN KEY (autor_id) REFERENCES autores(id)
);

-- ❌ EVITE: Tabela sem constraints (permite dados inválidos)
CREATE TABLE livros (
    id INTEGER,
    titulo TEXT,
    autor_id INTEGER
);
```

**Por quê usar constraints?**
- Previnem dados inválidos desde o início
- Melhoram performance (índices automáticos)
- Facilitam manutenção
- Documentam a intenção do schema

---

### 1.2 ALTER TABLE - O Maior Impacto em Produção

ALTER TABLE é uma das operações DDL mais perigosas em produção, especialmente em tabelas grandes.

#### Impacto de ADD COLUMN

```sql
-- Adicionar coluna simples
ALTER TABLE livros
ADD COLUMN preco REAL;
```

**No SQLite:**
- Operação relativamente rápida
- Não modifica registros existentes fisicamente
- Apenas atualiza o schema
- Registros antigos têm NULL (ou DEFAULT) para a nova coluna

**Em outros SGBDs (PostgreSQL, MySQL):**
- Pode ser lenta em tabelas grandes
- Pode bloquear a tabela durante a operação
- Pode causar downtime em sistemas críticos

#### Estratégias para ALTER TABLE em Tabelas Grandes

**Problema**: Tabela com 10 milhões de registros, precisa adicionar coluna NOT NULL.

**❌ Abordagem Ruim:**
```sql
-- Isso pode falhar ou ser muito lento
ALTER TABLE livros
ADD COLUMN nova_coluna TEXT NOT NULL;
```

**✅ Abordagem Segura:**

```sql
-- Passo 1: Adicionar coluna como NULL
ALTER TABLE livros
ADD COLUMN nova_coluna TEXT;

-- Passo 2: Popular dados existentes (em lotes se necessário)
UPDATE livros
SET nova_coluna = 'valor_padrao'
WHERE nova_coluna IS NULL;

-- Passo 3: Se necessário, alterar para NOT NULL (depende do SGBD)
-- No SQLite, isso pode não ser possível diretamente
```

#### Impacto de DROP COLUMN

**⚠️ CRÍTICO**: DROP COLUMN é uma operação **destrutiva e irreversível**!

```sql
-- Remove coluna e TODOS os seus dados
ALTER TABLE livros
DROP COLUMN editora;
```

**Considerações:**
- Dados são perdidos permanentemente
- Em SQLite (versões antigas), requer recriar a tabela
- Pode ser lenta em tabelas grandes
- Pode quebrar aplicações que dependem da coluna

**Estratégia segura:**
1. Fazer backup completo
2. Verificar dependências (aplicações, views, triggers)
3. Testar em ambiente de desenvolvimento
4. Executar em janela de manutenção
5. Monitorar aplicações após a mudança

---

### 1.3 DROP TABLE - Operação Destrutiva

DROP TABLE remove completamente a tabela e todos os seus dados.

```sql
DROP TABLE tabela_temporaria;
```

**Impacto:**
- **Imediato**: Tabela deixa de existir
- **Irreversível**: Sem backup, dados são perdidos permanentemente
- **Dependências**: Pode quebrar FOREIGN KEYs, views, triggers que referenciam a tabela

**Checklist antes de DROP TABLE:**
- [ ] Backup completo do banco de dados
- [ ] Verificar dependências (FOREIGN KEYs, views, triggers)
- [ ] Notificar equipe sobre a remoção
- [ ] Testar em ambiente de desenvolvimento
- [ ] Executar em janela de manutenção
- [ ] Documentar a remoção

---

### 1.4 TRUNCATE vs DELETE - Performance

**No SQLite**, use `DELETE FROM` como equivalente a TRUNCATE:

```sql
-- Limpar todos os dados
DELETE FROM livros_temporarios;
```

**Comparação de Performance:**

| Operação | Velocidade | Logs | Reversível |
|----------|------------|------|------------|
| **DELETE FROM** | Mais lenta | Logs de cada linha | Sim (transação) |
| **TRUNCATE** | Mais rápida | Mínimos logs | Não (auto-commit) |

**Quando usar cada um:**

- **DELETE FROM**: Quando precisa de controle transacional ou quer deletar com WHERE
- **TRUNCATE**: Quando quer limpar tudo rapidamente e não precisa reverter

---

## 2. Índices: Performance e Trade-offs

### 2.1 Quando Criar Índices

Índices melhoram drasticamente a performance de consultas, mas têm custos.

#### Benefícios dos Índices

```sql
-- Sem índice: Busca linear (lenta)
SELECT * FROM livros WHERE autor_id = 5;
-- Examina TODOS os registros

-- Com índice: Busca indexada (rápida)
CREATE INDEX idx_livros_autor ON livros(autor_id);
SELECT * FROM livros WHERE autor_id = 5;
-- Vai direto aos registros com autor_id = 5
```

**Crie índices em:**
- Colunas usadas frequentemente em WHERE
- Chaves estrangeiras (usadas em JOINs)
- Colunas usadas em ORDER BY
- Colunas usadas em GROUP BY
- Colunas com alta seletividade (muitos valores únicos)

#### Custos dos Índices

**Desvantagens:**
- Ocupam espaço em disco
- Atrasam INSERT (precisa atualizar índice)
- Atrasam UPDATE (se coluna indexada mudar)
- Atrasam DELETE (precisa remover do índice)

**Exemplo do impacto:**

```sql
-- Tabela sem índices: INSERT rápido
INSERT INTO livros (titulo, autor_id) VALUES ('Novo Livro', 1);
-- Tempo: 1ms

-- Tabela com 5 índices: INSERT mais lento
INSERT INTO livros (titulo, autor_id) VALUES ('Novo Livro', 1);
-- Tempo: 5ms (atualiza 5 índices)
```

### 2.2 Índices Compostos vs Separados

#### Índice Composto

```sql
-- Índice composto (múltiplas colunas)
CREATE INDEX idx_emprestimos_usuario_status 
ON emprestimos(usuario_id, status);
```

**Quando usar:**
- Você sempre consulta as colunas juntas
- Exemplo: `WHERE usuario_id = 5 AND status = 'ativo'`

**Vantagem**: Mais eficiente para consultas que usam ambas as colunas

**Desvantagem**: Não ajuda em consultas que usam apenas uma das colunas

#### Índices Separados

```sql
-- Índices separados
CREATE INDEX idx_emprestimos_usuario ON emprestimos(usuario_id);
CREATE INDEX idx_emprestimos_status ON emprestimos(status);
```

**Quando usar:**
- Você consulta as colunas independentemente
- Exemplo: `WHERE usuario_id = 5` OU `WHERE status = 'ativo'`

**Vantagem**: Funciona bem para ambas as consultas

**Desvantagem**: Ocupa mais espaço e atrasa mais os INSERTs

### 2.3 Monitoramento de Índices

**Verificar índices existentes:**

```sql
-- SQLite
SELECT name, tbl_name, sql 
FROM sqlite_master 
WHERE type = 'index';

-- Verificar se índice está sendo usado (em outros SGBDs)
-- PostgreSQL: EXPLAIN ANALYZE
-- MySQL: EXPLAIN
```

**Remover índices não utilizados:**

```sql
-- Se um índice não está sendo usado, remova-o
DROP INDEX idx_livros_nao_usado;
```

**Regra**: Monitore o uso de índices e remova os que não estão sendo utilizados.

---

## 3. Constraints: Integridade vs Performance

### 3.1 Impacto de Constraints na Performance

Constraints garantem integridade, mas têm custos de performance.

#### PRIMARY KEY

```sql
CREATE TABLE livros (
    id INTEGER PRIMARY KEY AUTOINCREMENT
);
```

**Impacto:**
- ✅ Cria índice automaticamente (melhora buscas)
- ⚠️ Valida unicidade (pequeno custo em INSERT)
- ✅ Geralmente melhora performance geral

#### FOREIGN KEY

```sql
CREATE TABLE livros (
    autor_id INTEGER,
    FOREIGN KEY (autor_id) REFERENCES autores(id)
);
```

**Impacto:**
- ✅ Garante integridade referencial
- ⚠️ Valida existência em INSERT/UPDATE (pequeno custo)
- ⚠️ Pode atrasar DELETE (verifica dependências)

**Recomendação**: Sempre use FOREIGN KEY. O pequeno custo de performance é compensado pela garantia de integridade.

#### CHECK

```sql
CREATE TABLE avaliacoes (
    nota INTEGER CHECK (nota >= 1 AND nota <= 5)
);
```

**Impacto:**
- ✅ Previne dados inválidos
- ⚠️ Valida em cada INSERT/UPDATE (pequeno custo)
- ✅ Custo geralmente baixo

**Recomendação**: Use CHECK para validações importantes. O custo é mínimo comparado ao benefício.

### 3.2 Quando Não Usar Constraints

**Evite constraints excessivas:**

```sql
-- ❌ EVITE: Constraint muito complexa que raramente é violada
CREATE TABLE livros (
    titulo TEXT CHECK (
        LENGTH(titulo) > 5 AND 
        LENGTH(titulo) < 200 AND
        titulo LIKE '% %'  -- Deve ter espaço
    )
);
```

**Problema**: Validação complexa em cada INSERT pode ser custosa.

**Solução**: Valide na aplicação para casos simples, use CHECK apenas para regras críticas.

---

## 4. Boas Práticas de Nomenclatura

### 4.1 Convenções de Nomenclatura

#### Tabelas

```sql
-- ✅ BOM: Nomes no plural, descritivos, em minúsculas
CREATE TABLE livros (...);
CREATE TABLE usuarios (...);
CREATE TABLE emprestimos (...);

-- ❌ EVITE: Nomes confusos
CREATE TABLE tbl1 (...);
CREATE TABLE data (...);  -- palavra reservada
CREATE TABLE Livros (...);  -- case inconsistente
```

#### Colunas

```sql
-- ✅ BOM: Nomes descritivos, em minúsculas, snake_case
CREATE TABLE livros (
    id INTEGER,
    titulo TEXT,
    ano_publicacao INTEGER,
    data_cadastro DATE
);

-- ❌ EVITE: Nomes ambíguos
CREATE TABLE livros (
    id INTEGER,
    nome TEXT,  -- "nome" do quê? título? autor?
    ano INTEGER,  -- muito genérico
    dt DATE  -- abreviação confusa
);
```

#### Índices

```sql
-- ✅ BOM: Prefixo "idx_" + nome descritivo
CREATE INDEX idx_livros_autor ON livros(autor_id);
CREATE INDEX idx_emprestimos_usuario_status ON emprestimos(usuario_id, status);

-- ❌ EVITE: Nomes genéricos
CREATE INDEX index1 ON livros(autor_id);
CREATE INDEX idx1 ON livros(autor_id);
```

### 4.2 Documentação de Schema

**Comente constraints complexas:**

```sql
CREATE TABLE emprestimos (
    id INTEGER PRIMARY KEY,
    -- Status pode ser: 'ativo', 'devolvido', 'atrasado'
    status TEXT CHECK (status IN ('ativo', 'devolvido', 'atrasado')),
    -- Data limite deve ser >= data_emprestimo
    data_limite DATE,
    CHECK (data_limite IS NULL OR data_limite >= data_emprestimo)
);
```

---

## 5. Segurança: Backup e Versionamento

### 5.1 Backup Antes de DDL

**Sempre faça backup antes de operações DDL destrutivas:**

```bash
# SQLite: Copiar arquivo do banco
cp biblioteca.db biblioteca.db.backup

# Outros SGBDs: Usar ferramentas de backup
# PostgreSQL: pg_dump
# MySQL: mysqldump
```

**Checklist de backup:**
- [ ] Backup completo do banco
- [ ] Backup do schema (estrutura)
- [ ] Verificar integridade do backup
- [ ] Testar restauração em ambiente de teste

### 5.2 Versionamento de Schema

**Mantenha histórico de mudanças:**

```sql
-- migrations/001_create_livros.sql
CREATE TABLE livros (
    id INTEGER PRIMARY KEY,
    titulo TEXT NOT NULL
);

-- migrations/002_add_preco_to_livros.sql
ALTER TABLE livros
ADD COLUMN preco REAL;

-- migrations/003_create_avaliacoes.sql
CREATE TABLE avaliacoes (
    id INTEGER PRIMARY KEY,
    livro_id INTEGER,
    nota INTEGER,
    FOREIGN KEY (livro_id) REFERENCES livros(id)
);
```

**Ferramentas de versionamento:**
- **SQLite**: Arquivos SQL de migração
- **PostgreSQL**: Ferramentas como Flyway, Liquibase
- **MySQL**: Ferramentas como Flyway, Liquibase

### 5.3 Testes em Ambiente de Desenvolvimento

**Nunca execute DDL em produção sem testar primeiro:**

1. **Desenvolvimento**: Teste a operação
2. **Staging**: Teste em ambiente similar à produção
3. **Produção**: Execute com cuidado e monitoramento

---

## 6. Estratégias de Migração

### 6.1 Migração de Schema em Tabelas Grandes

**Problema**: Adicionar coluna NOT NULL em tabela com milhões de registros.

**Estratégia em Etapas:**

```sql
-- Etapa 1: Adicionar coluna como NULL
ALTER TABLE livros
ADD COLUMN nova_coluna TEXT;

-- Etapa 2: Popular dados em lotes (para não travar)
UPDATE livros
SET nova_coluna = 'valor_padrao'
WHERE id BETWEEN 1 AND 10000;

UPDATE livros
SET nova_coluna = 'valor_padrao'
WHERE id BETWEEN 10001 AND 20000;
-- ... continuar em lotes

-- Etapa 3: Verificar que todos têm valor
SELECT COUNT(*) FROM livros WHERE nova_coluna IS NULL;
-- Deve retornar 0

-- Etapa 4: Em outros SGBDs, alterar para NOT NULL
-- (SQLite pode não suportar isso diretamente)
```

### 6.2 Recriar Tabela (Quando ALTER TABLE Não Funciona)

**Quando ALTER TABLE é limitado (SQLite antigo), recrie a tabela:**

```sql
-- Passo 1: Criar nova tabela com estrutura desejada
CREATE TABLE livros_nova (
    id INTEGER PRIMARY KEY,
    titulo TEXT NOT NULL,
    preco REAL  -- nova coluna
);

-- Passo 2: Copiar dados
INSERT INTO livros_nova (id, titulo, preco)
SELECT id, titulo, NULL FROM livros;

-- Passo 3: Popular nova coluna
UPDATE livros_nova SET preco = 0 WHERE preco IS NULL;

-- Passo 4: Remover tabela antiga
DROP TABLE livros;

-- Passo 5: Renomear nova tabela
ALTER TABLE livros_nova RENAME TO livros;

-- Passo 6: Recriar índices
CREATE INDEX idx_livros_autor ON livros(autor_id);
```

**⚠️ CUIDADO**: Esta estratégia requer downtime e pode quebrar aplicações durante a migração.

---

## 7. Checklist de Boas Práticas DDL

### Antes de Executar DDL

- [ ] **Backup completo** do banco de dados
- [ ] **Testado em desenvolvimento** primeiro
- [ ] **Documentado** a mudança e o motivo
- [ ] **Notificada a equipe** sobre a mudança
- [ ] **Verificadas dependências** (aplicações, views, triggers)
- [ ] **Planejada janela de manutenção** se necessário

### Durante a Execução

- [ ] **Monitorar** o progresso da operação
- [ ] **Verificar logs** para erros
- [ ] **Não interromper** operações em andamento

### Após a Execução

- [ ] **Verificar** que a operação foi bem-sucedida
- [ ] **Testar** aplicações que dependem da estrutura
- [ ] **Documentar** a mudança no changelog
- [ ] **Monitorar** performance após a mudança

---

## 8. Erros Comuns e Como Evitá-los

### 8.1 Esquecer WHERE em DELETE/UPDATE

**❌ Erro comum:**
```sql
-- Esqueceu WHERE - deleta TUDO!
DELETE FROM livros;
```

**✅ Correto:**
```sql
DELETE FROM livros WHERE id = 1;
```

### 8.2 DROP TABLE Sem Verificar Dependências

**❌ Erro comum:**
```sql
-- Remove tabela que é referenciada por outras
DROP TABLE autores;
-- Erro: FOREIGN KEY constraint failed
```

**✅ Correto:**
```sql
-- Verificar dependências primeiro
SELECT name FROM sqlite_master 
WHERE sql LIKE '%autores%';

-- Depois, se seguro, remover
DROP TABLE IF EXISTS autores;
```

### 8.3 Adicionar Coluna NOT NULL Sem DEFAULT

**❌ Erro comum:**
```sql
-- Falha se tabela já tem dados
ALTER TABLE livros
ADD COLUMN nova_coluna TEXT NOT NULL;
```

**✅ Correto:**
```sql
-- Adicionar como NULL primeiro
ALTER TABLE livros
ADD COLUMN nova_coluna TEXT;

-- Popular dados
UPDATE livros SET nova_coluna = 'padrao';

-- Depois, se necessário, alterar para NOT NULL
```

### 8.4 Criar Índices Demais

**❌ Erro comum:**
```sql
-- Índice em cada coluna (desnecessário)
CREATE INDEX idx_livros_titulo ON livros(titulo);
CREATE INDEX idx_livros_ano ON livros(ano_publicacao);
CREATE INDEX idx_livros_editora ON livros(editora);
-- ... muitos outros
```

**✅ Correto:**
```sql
-- Apenas índices realmente necessários
CREATE INDEX idx_livros_autor ON livros(autor_id);  -- Usado em JOINs
CREATE INDEX idx_livros_categoria ON livros(categoria_id);  -- Usado em JOINs
-- Não criar índice em colunas raramente consultadas
```

---

## 9. Performance: Métricas e Monitoramento

### 9.1 Medir Tempo de Operações DDL

**No SQLite:**
```sql
-- Habilitar timer
.timer on

-- Executar operação
ALTER TABLE livros ADD COLUMN preco REAL;

-- Ver tempo de execução
```

**Em outros SGBDs:**
- Use ferramentas de monitoramento
- Analise logs de performance
- Use EXPLAIN para entender o plano de execução

### 9.2 Identificar Operações Lentas

**Sinais de problemas:**
- ALTER TABLE demora muito (> 1 minuto em tabelas pequenas)
- DROP TABLE trava o banco
- CREATE INDEX demora muito

**Soluções:**
- Executar em janela de manutenção
- Dividir em operações menores
- Usar estratégias de migração alternativas

---

## 10. Conclusão: DDL Eficiente e Seguro

### Princípios Fundamentais

1. **Planejamento**: Pense antes de executar
2. **Backup**: Sempre faça backup antes de operações destrutivas
3. **Testes**: Teste em desenvolvimento primeiro
4. **Documentação**: Documente todas as mudanças
5. **Monitoramento**: Monitore impacto após mudanças

### Trade-offs a Considerar

- **Constraints**: Integridade vs Performance (geralmente vale a pena)
- **Índices**: Leitura rápida vs Escrita lenta (balanceie conforme necessidade)
- **ALTER TABLE**: Flexibilidade vs Risco (planeje cuidadosamente)

### Próximos Passos

Agora que você entende DDL, performance e boas práticas:

1. Pratique criando e modificando tabelas
2. Experimente diferentes estratégias de migração
3. Monitore o impacto de suas mudanças
4. Desenvolva seu próprio conjunto de boas práticas

---

**⚠️ Lembrete Final**: DDL é poderoso, mas pode ser destrutivo. Sempre priorize segurança sobre velocidade. É melhor ser cuidadoso e lento do que rápido e destrutivo.

**Boa sorte com suas operações DDL! 🚀**


