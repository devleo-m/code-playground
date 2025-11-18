# Aula 2 - Performance, Boas Práticas e Otimização

## Introdução: Construindo Fundamentos Sólidos

Nesta aula, você aprendeu os comandos básicos de SQL (SELECT, INSERT, UPDATE, DELETE). Agora vamos aprender como usá-los de forma eficiente, segura e profissional desde o início. Esses hábitos são fundamentais e evitarão problemas graves no futuro.

**Regra de Ouro**: É muito mais fácil escrever código eficiente e seguro desde o início do que corrigir problemas depois.

---

## 1. Performance: Otimizando Queries SELECT

### 1.1 O Problema do SELECT *

**NUNCA use `SELECT *` em produção**, exceto em casos muito específicos:

```sql
-- ❌ EVITE em produção
SELECT * FROM livros WHERE id = 1;

-- ✅ PREFIRA
SELECT titulo, ano_publicacao, quantidade_disponivel 
FROM livros 
WHERE id = 1;
```

**Por quê evitar SELECT ***:

1. **Transferência de Dados Desnecessária**:
   - Retorna todas as colunas, mesmo as que você não precisa
   - Mais tráfego de rede
   - Mais uso de memória
   - Mais tempo de processamento

2. **Fragilidade do Código**:
   - Se a estrutura da tabela mudar (nova coluna adicionada), sua aplicação pode quebrar
   - Se uma coluna for removida, sua aplicação quebra
   - Dificulta manutenção

3. **Clareza de Intenção**:
   - `SELECT *` não deixa claro quais dados você realmente precisa
   - Outros desenvolvedores não sabem o que a query retorna

4. **Performance em Tabelas Grandes**:
   - Tabelas com muitas colunas ou colunas BLOB são muito lentas com `SELECT *`

**Quando SELECT * é aceitável**:
- Durante desenvolvimento e exploração de dados
- Em queries administrativas onde você realmente precisa de tudo
- Em tabelas pequenas com poucas colunas (mas ainda não é ideal)

### 1.2 Especificando Colunas: Boa Prática

```sql
-- ✅ Sempre especifique as colunas necessárias
SELECT 
    titulo,
    ano_publicacao,
    quantidade_disponivel
FROM livros
WHERE categoria_id = 1;
```

**Benefícios**:
- Mais rápido (menos dados transferidos)
- Mais claro (sabe exatamente o que retorna)
- Mais seguro (não quebra se estrutura mudar)
- Melhor para documentação

### 1.3 Usando WHERE Eficientemente

WHERE é crucial para performance. Sem WHERE, o banco precisa examinar TODOS os registros:

```sql
-- ❌ Ruim: Examina TODOS os livros
SELECT titulo FROM livros;

-- ✅ Melhor: Filtra antes de retornar
SELECT titulo 
FROM livros 
WHERE ano_publicacao > 2000;
```

**Regra**: Sempre use WHERE quando possível. Quanto mais específico o filtro, melhor.

### 1.4 LIMIT: Controlando o Volume de Dados

Sempre use LIMIT quando você não precisa de todos os resultados:

```sql
-- ❌ Ruim: Pode retornar milhares de registros
SELECT titulo FROM livros ORDER BY titulo;

-- ✅ Melhor: Limita os resultados
SELECT titulo 
FROM livros 
ORDER BY titulo 
LIMIT 10;
```

**Benefícios do LIMIT**:
- Reduz uso de memória
- Reduz tempo de processamento
- Reduz tráfego de rede
- Melhor experiência do usuário (resultados aparecem mais rápido)

**Uso em Paginação**:
```sql
-- Primeira página (10 primeiros)
SELECT titulo FROM livros LIMIT 10 OFFSET 0;

-- Segunda página (próximos 10)
SELECT titulo FROM livros LIMIT 10 OFFSET 10;

-- Terceira página
SELECT titulo FROM livros LIMIT 10 OFFSET 20;
```

### 1.5 Índices e Performance em SELECT

Índices são cruciais para performance, especialmente em WHERE e ORDER BY:

```sql
-- Se você busca frequentemente por autor_id, crie um índice:
CREATE INDEX idx_livros_autor ON livros(autor_id);

-- Agora esta query será muito mais rápida:
SELECT titulo 
FROM livros 
WHERE autor_id = 5;
```

**Quando criar índices**:
- Colunas usadas frequentemente em WHERE
- Colunas usadas em JOINs
- Colunas usadas em ORDER BY

**Quando NÃO criar índices**:
- Tabelas muito pequenas (o custo de manter o índice pode ser maior que o benefício)
- Colunas raramente usadas em filtros
- Colunas que mudam frequentemente (UPDATE frequente)

---

## 2. Performance: Otimizando INSERT

### 2.1 INSERT Múltiplo vs Múltiplos INSERTs

Sempre que possível, use INSERT múltiplo em vez de vários INSERTs individuais:

```sql
-- ❌ Ruim: 3 operações separadas
INSERT INTO autores (nome, nacionalidade) VALUES ('Autor 1', 'Brasileiro');
INSERT INTO autores (nome, nacionalidade) VALUES ('Autor 2', 'Brasileiro');
INSERT INTO autores (nome, nacionalidade) VALUES ('Autor 3', 'Brasileiro');

-- ✅ Melhor: 1 operação
INSERT INTO autores (nome, nacionalidade)
VALUES 
    ('Autor 1', 'Brasileiro'),
    ('Autor 2', 'Brasileiro'),
    ('Autor 3', 'Brasileiro');
```

**Por quê**:
- Menos overhead de transação
- Mais rápido (uma única operação)
- Mais eficiente em termos de I/O

**Limite prático**: Em SQLite, você pode inserir centenas ou milhares de registros de uma vez. Em outros SGBDs, o limite pode variar.

### 2.2 Especificando Colunas em INSERT

Sempre especifique as colunas em INSERT:

```sql
-- ❌ Ruim: Depende da ordem das colunas
INSERT INTO livros VALUES (NULL, 'Título', 'ISBN', 2000, 'Editora', 1, 1, 5);

-- ✅ Melhor: Explícito e claro
INSERT INTO livros (titulo, isbn, ano_publicacao, editora, autor_id, categoria_id, quantidade_disponivel)
VALUES ('Título', 'ISBN', 2000, 'Editora', 1, 1, 5);
```

**Benefícios**:
- Não quebra se a ordem das colunas mudar
- Mais legível
- Mais fácil de manter
- Evita erros (especialmente com valores NULL)

### 2.3 Valores Padrão e INSERT

Use valores padrão quando apropriado:

```sql
-- Se quantidade_disponivel tem DEFAULT 0, você pode omitir:
INSERT INTO livros (titulo, autor_id, categoria_id)
VALUES ('Novo Livro', 1, 1);
-- quantidade_disponivel será automaticamente 0
```

**Benefício**: Menos código, menos chance de erro.

---

## 3. Performance e Segurança: UPDATE

### 3.1 WHERE é OBRIGATÓRIO em UPDATE

**NUNCA execute UPDATE sem WHERE**, a menos que você realmente queira atualizar TODOS os registros:

```sql
-- ⚠️ CATASTROFE: Atualiza TODOS os livros!
UPDATE livros SET quantidade_disponivel = 0;
-- Isso zeraria o estoque de TODOS os livros no banco!
```

**Sempre use WHERE**:
```sql
-- ✅ Seguro: Atualiza apenas um livro
UPDATE livros 
SET quantidade_disponivel = 0 
WHERE id = 1;
```

### 3.2 Testando UPDATE com SELECT Primeiro

**Sempre teste com SELECT antes de fazer UPDATE**:

```sql
-- 1. PRIMEIRO: Veja o que será afetado
SELECT * FROM livros WHERE categoria_id = 1;

-- 2. Verifique se está correto

-- 3. DEPOIS: Execute o UPDATE
UPDATE livros 
SET quantidade_disponivel = quantidade_disponivel + 1 
WHERE categoria_id = 1;
```

**Benefícios**:
- Evita atualizações acidentais
- Permite revisar os dados antes de modificar
- Facilita debugging

### 3.3 UPDATE com Expressões

Você pode usar expressões em UPDATE, mas cuidado com a performance:

```sql
-- ✅ Bom: Atualização simples
UPDATE livros 
SET quantidade_disponivel = quantidade_disponivel + 1 
WHERE id = 1;

-- ⚠️ Cuidado: Atualização em massa pode ser lenta
UPDATE livros 
SET quantidade_disponivel = quantidade_disponivel * 2 
WHERE categoria_id = 1;
-- Se houver 100.000 livros nesta categoria, pode ser lento
```

**Dica**: Para atualizações em massa, considere fazer em lotes ou usar transações.

### 3.4 Índices e Performance em UPDATE

Índices melhoram WHERE, mas podem tornar UPDATE mais lento:

```sql
-- Índice ajuda na busca:
CREATE INDEX idx_livros_categoria ON livros(categoria_id);

-- Mas quando você atualiza, o índice também precisa ser atualizado
UPDATE livros 
SET quantidade_disponivel = 0 
WHERE categoria_id = 1;
-- O banco atualiza a tabela E o índice
```

**Trade-off**: Índices aceleram SELECT, mas podem desacelerar UPDATE/INSERT/DELETE. Use com sabedoria.

---

## 4. Performance e Segurança: DELETE

### 4.1 WHERE é AINDA MAIS CRUCIAL em DELETE

**DELETE sem WHERE é CATASTROFICO**:

```sql
-- ⚠️ CATASTROFE TOTAL: Deleta TODOS os livros!
DELETE FROM livros;
-- Isso apagaria TODOS os registros da tabela!
```

**Sempre use WHERE**:
```sql
-- ✅ Seguro: Deleta apenas um livro
DELETE FROM livros WHERE id = 15;
```

### 4.2 Testando DELETE com SELECT Primeiro

**SEMPRE teste com SELECT antes de DELETE**:

```sql
-- 1. PRIMEIRO: Veja o que será deletado
SELECT * FROM emprestimos 
WHERE status = 'devolvido' 
  AND data_devolucao_real < '2020-01-01';

-- 2. Verifique se está correto
-- 3. Conte quantos registros serão afetados
SELECT COUNT(*) FROM emprestimos 
WHERE status = 'devolvido' 
  AND data_devolucao_real < '2020-01-01';

-- 4. DEPOIS: Execute o DELETE
DELETE FROM emprestimos 
WHERE status = 'devolvido' 
  AND data_devolucao_real < '2020-01-01';
```

### 4.3 Soft Delete vs Hard Delete

Em muitos casos, é melhor fazer "soft delete" (marcar como deletado) em vez de deletar realmente:

```sql
-- Em vez de deletar:
DELETE FROM livros WHERE id = 15;

-- Considere marcar como inativo:
ALTER TABLE livros ADD COLUMN ativo INTEGER DEFAULT 1;

UPDATE livros 
SET ativo = 0 
WHERE id = 15;

-- Depois, filtre os ativos:
SELECT * FROM livros WHERE ativo = 1;
```

**Vantagens do Soft Delete**:
- Pode recuperar dados se necessário
- Mantém histórico
- Permite auditoria
- Evita problemas com chaves estrangeiras

**Desvantagens**:
- Tabela cresce (mas pode arquivar depois)
- Precisa filtrar em todas as queries

### 4.4 DELETE e Integridade Referencial

Cuidado ao deletar registros que são referenciados por outros:

```sql
-- Isso pode falhar se houver empréstimos deste livro:
DELETE FROM livros WHERE id = 1;
-- Erro: FOREIGN KEY constraint failed
```

**Soluções**:
1. Deletar registros relacionados primeiro
2. Usar CASCADE (deleta automaticamente os relacionados)
3. Fazer soft delete em vez de hard delete

---

## 5. Operadores e Performance

### 5.1 LIKE e Performance

LIKE pode ser lento, especialmente com `%` no início:

```sql
-- ❌ Lento: Não pode usar índice eficientemente
SELECT * FROM livros WHERE titulo LIKE '%Dom%';

-- ✅ Mais rápido: Pode usar índice (busca por prefixo)
SELECT * FROM livros WHERE titulo LIKE 'Dom%';

-- ✅ Ainda melhor: Se possível, use busca exata
SELECT * FROM livros WHERE titulo = 'Dom Casmurro';
```

**Regra**: Se possível, evite `LIKE '%texto%'`. Prefira `LIKE 'texto%'` ou busca exata.

### 5.2 Operadores de Comparação

Operadores simples (`=`, `>`, `<`) são geralmente mais rápidos que operadores complexos:

```sql
-- ✅ Rápido: Comparação simples
SELECT * FROM livros WHERE ano_publicacao = 2000;

-- ✅ Rápido: Comparação de intervalo
SELECT * FROM livros WHERE ano_publicacao BETWEEN 1990 AND 2000;

-- ⚠️ Pode ser mais lento: Múltiplas condições OR
SELECT * FROM livros 
WHERE ano_publicacao = 1990 
   OR ano_publicacao = 2000 
   OR ano_publicacao = 2010;
```

### 5.3 IN vs Múltiplos OR

IN é geralmente mais eficiente que múltiplos OR:

```sql
-- ❌ Menos eficiente
SELECT * FROM livros 
WHERE categoria_id = 1 
   OR categoria_id = 2 
   OR categoria_id = 3;

-- ✅ Mais eficiente
SELECT * FROM livros 
WHERE categoria_id IN (1, 2, 3);
```

**Benefício**: Mais legível e geralmente mais rápido.

### 5.4 NULL e Performance

Verificações de NULL são eficientes, mas cuidado com lógica complexa:

```sql
-- ✅ Eficiente
SELECT * FROM livros WHERE ano_publicacao IS NULL;

-- ⚠️ Pode ser lento se houver muitos NULLs
SELECT * FROM livros 
WHERE ano_publicacao IS NULL 
   OR ano_publicacao > 2000;
```

---

## 6. Boas Práticas de Nomenclatura e Organização

### 6.1 Nomes de Tabelas e Colunas

Use nomes claros e descritivos:

```sql
-- ❌ Ruim: Abreviações obscuras
CREATE TABLE usr (usr_id INT, usr_nm TEXT);

-- ✅ Bom: Nomes claros
CREATE TABLE usuarios (id INTEGER, nome TEXT);
```

**Convenções comuns**:
- Tabelas: plural (`usuarios`, `livros`)
- Colunas: singular (`nome`, `email`)
- IDs: `id` ou `tabela_id` (ex: `usuario_id`)
- Chaves estrangeiras: `tabela_referenciada_id` (ex: `autor_id`)

### 6.2 Organização de Queries

Organize queries complexas de forma legível:

```sql
-- ❌ Ruim: Tudo em uma linha
SELECT l.titulo, a.nome, c.nome FROM livros l JOIN autores a ON l.autor_id = a.id JOIN categorias c ON l.categoria_id = c.id WHERE l.quantidade_disponivel > 0 ORDER BY l.titulo;

-- ✅ Bom: Organizado e legível
SELECT 
    l.titulo,
    a.nome AS autor,
    c.nome AS categoria
FROM livros l
JOIN autores a ON l.autor_id = a.id
JOIN categorias c ON l.categoria_id = c.id
WHERE l.quantidade_disponivel > 0
ORDER BY l.titulo;
```

**Benefícios**:
- Mais fácil de ler
- Mais fácil de manter
- Mais fácil de debugar
- Mais fácil de revisar

### 6.3 Comentários

Use comentários para explicar queries complexas:

```sql
-- Buscar livros disponíveis de autores brasileiros
-- para exibir na página inicial
SELECT 
    l.titulo,
    a.nome AS autor
FROM livros l
JOIN autores a ON l.autor_id = a.id
WHERE a.nacionalidade = 'Brasileiro'
  AND l.quantidade_disponivel > 0
ORDER BY l.titulo
LIMIT 10;
```

**Quando comentar**:
- Lógica complexa ou não óbvia
- Decisões de negócio importantes
- Workarounds ou soluções temporárias
- Queries que serão executadas manualmente

---

## 7. Segurança: Proteção Contra SQL Injection

### 7.1 O Problema do SQL Injection

SQL Injection é uma vulnerabilidade crítica onde atacantes podem executar código SQL malicioso:

```sql
-- ⚠️ VULNERÁVEL: Se o usuário digitar código SQL
SELECT * FROM usuarios WHERE email = 'admin@test.com' OR '1'='1';
-- Isso retornaria TODOS os usuários!
```

### 7.2 Usando Parâmetros (Prepared Statements)

**SEMPRE use parâmetros** em vez de concatenar strings:

```sql
-- ❌ VULNERÁVEL: Concatenando strings
-- (em código de aplicação)
query = "SELECT * FROM usuarios WHERE email = '" + userInput + "'";

-- ✅ SEGURO: Usando parâmetros
-- (em código de aplicação)
query = "SELECT * FROM usuarios WHERE email = ?";
-- Depois, passe userInput como parâmetro
```

**No SQLite CLI**, você não precisa se preocupar tanto, mas em aplicações, **SEMPRE use prepared statements**.

### 7.3 Validação de Entrada

Sempre valide e sanitize dados de entrada:

```sql
-- Antes de inserir, valide:
-- - Email tem formato válido?
-- - Nome não contém caracteres especiais perigosos?
-- - Números estão no range esperado?
```

---

## 8. Transações e Integridade de Dados

### 8.1 O que são Transações?

Transações garantem que operações múltiplas sejam executadas como uma unidade atômica:

```sql
-- Iniciar transação
BEGIN TRANSACTION;

-- Operações
INSERT INTO livros (titulo, autor_id) VALUES ('Livro 1', 1);
UPDATE autores SET nome = 'Novo Nome' WHERE id = 1;

-- Se tudo estiver OK, confirmar
COMMIT;

-- Se houver erro, reverter
ROLLBACK;
```

**Propriedades ACID**:
- **Atomicidade**: Tudo ou nada
- **Consistência**: Dados sempre válidos
- **Isolamento**: Transações não interferem
- **Durabilidade**: Mudanças são permanentes

### 8.2 Quando Usar Transações

Use transações quando:
- Múltiplas operações devem ser atômicas
- Operações críticas (não pode ter estado intermediário)
- Operações que podem falhar

**Exemplo prático**:
```sql
-- Empréstimo de livro: deve atualizar livro E criar empréstimo
BEGIN TRANSACTION;

UPDATE livros 
SET quantidade_disponivel = quantidade_disponivel - 1 
WHERE id = 1;

INSERT INTO emprestimos (livro_id, usuario_id, data_devolucao_prevista)
VALUES (1, 1, date('now', '+14 days'));

-- Se qualquer operação falhar, ambas são revertidas
COMMIT;
```

---

## 9. Checklist de Boas Práticas

### Antes de Escrever uma Query

- [ ] Entendi o que a query deve fazer?
- [ ] Preciso de todas as colunas ou posso especificar?
- [ ] Preciso de todos os registros ou posso usar WHERE/LIMIT?
- [ ] A query será executada frequentemente? (precisa de índice?)

### Ao Escrever SELECT

- [ ] Especifiquei as colunas necessárias (não use `SELECT *`)?
- [ ] Usei WHERE para filtrar quando possível?
- [ ] Usei LIMIT quando não preciso de todos os resultados?
- [ ] A query está organizada e legível?

### Ao Escrever INSERT

- [ ] Especifiquei as colunas explicitamente?
- [ ] Posso usar INSERT múltiplo em vez de vários INSERTs?
- [ ] Validei os dados antes de inserir?

### Ao Escrever UPDATE

- [ ] **USEI WHERE?** (verifique duas vezes!)
- [ ] Testei com SELECT primeiro?
- [ ] A atualização afeta apenas os registros corretos?
- [ ] Considerei usar transação?

### Ao Escrever DELETE

- [ ] **USEI WHERE?** (verifique três vezes!)
- [ ] Testei com SELECT primeiro?
- [ ] Contei quantos registros serão afetados?
- [ ] Considerei soft delete em vez de hard delete?
- [ ] Verifiquei se há chaves estrangeiras que podem impedir?

### Após Escrever a Query

- [ ] Testei a query?
- [ ] Verifiquei os resultados?
- [ ] A query é eficiente para o volume de dados esperado?
- [ ] Documentei queries complexas?

---

## 10. Identificando Queries Lentas

### 10.1 Sinais de Query Lenta

- Query demora mais de alguns segundos
- Sistema fica lento quando a query executa
- Uso alto de CPU ou memória
- Timeout de conexão

### 10.2 Como Investigar

**No SQLite**, você pode usar `EXPLAIN QUERY PLAN`:

```sql
EXPLAIN QUERY PLAN
SELECT titulo 
FROM livros 
WHERE autor_id = 5;
```

Isso mostra como o banco planeja executar a query.

### 10.3 Soluções Comuns

1. **Adicionar índice**:
   ```sql
   CREATE INDEX idx_livros_autor ON livros(autor_id);
   ```

2. **Refinar WHERE**:
   ```sql
   -- Mais específico = mais rápido
   WHERE autor_id = 5 AND categoria_id = 1;
   ```

3. **Limitar resultados**:
   ```sql
   LIMIT 10;
   ```

4. **Especificar colunas**:
   ```sql
   SELECT titulo, ano_publicacao;  -- Em vez de SELECT *
   ```

---

## 11. Resumo: O que Fazer e o que NÃO Fazer

### ✅ FAÇA

- Especifique colunas em SELECT
- Use WHERE em UPDATE e DELETE (sempre!)
- Teste com SELECT antes de modificar dados
- Use LIMIT quando apropriado
- Organize queries de forma legível
- Use transações para operações críticas
- Valide dados de entrada
- Use índices em colunas frequentemente consultadas
- Comente queries complexas

### ❌ NÃO FAÇA

- Não use `SELECT *` em produção
- Não execute UPDATE sem WHERE
- Não execute DELETE sem WHERE
- Não concatene strings em queries (SQL Injection)
- Não ignore erros de integridade referencial
- Não crie índices desnecessários
- Não faça operações em massa sem testar primeiro
- Não confie apenas na aplicação para validação

---

## 12. Próximos Passos

Agora que você entende as boas práticas básicas:

1. **Pratique**: Escreva queries seguindo estas práticas
2. **Revise**: Olhe suas queries antigas e melhore-as
3. **Pense**: Sempre considere performance e segurança
4. **Aprenda**: Continue estudando índices, JOINs e otimização avançada

**Lembrete**: Essas práticas são fundamentais. Quanto mais cedo você adotá-las, melhor desenvolvedor/analista de dados você será.

---

**Próximo Passo**: Após completar os exercícios e estudar esta seção, envie suas respostas para análise e feedback.

