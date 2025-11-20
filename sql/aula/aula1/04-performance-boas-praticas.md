# Aula 1 - Performance, Boas Práticas e Otimização

## Introdução: Por que Performance Importa?

No contexto de uma introdução ao SQL, pode parecer prematuro falar sobre performance. No entanto, entender desde o início os princípios de boas práticas e otimização é fundamental para desenvolver hábitos corretos que evitarão problemas futuros.

**Regra de Ouro**: É muito mais fácil escrever código eficiente desde o início do que otimizar código ineficiente depois.

## 1. Performance: O Impacto das Queries no Desempenho

### Entendendo o Custo de uma Query

Toda query SQL tem um "custo" computacional. Vamos entender os fatores:

#### Fatores que Afetam a Performance

1. **Número de Registros Examinados**
   ```sql
   -- ❌ Ruim: Examina TODOS os registros
   SELECT * FROM livros;
   
   -- ✅ Melhor: Filtra antes de retornar
   SELECT titulo FROM livros WHERE ano_publicacao > 2000;
   ```

2. **Número de Colunas Retornadas**
   ```sql
   -- ❌ Ruim: Retorna todas as colunas (mesmo as que não precisa)
   SELECT * FROM livros WHERE id = 1;
   
   -- ✅ Melhor: Retorna apenas o necessário
   SELECT titulo, ano_publicacao FROM livros WHERE id = 1;
   ```

3. **Complexidade dos JOINs**
   ```sql
   -- Quanto mais tabelas no JOIN, mais complexo
   SELECT * 
   FROM livros l
   JOIN autores a ON l.autor_id = a.id
   JOIN categorias c ON l.categoria_id = c.id
   JOIN emprestimos e ON l.id = e.livro_id;
   -- 4 tabelas = mais processamento
   ```

4. **Operações de Ordenação e Agregação**
   ```sql
   -- ORDER BY e GROUP BY são operações custosas
   SELECT categoria_id, COUNT(*) 
   FROM livros 
   GROUP BY categoria_id 
   ORDER BY COUNT(*) DESC;
   ```

### O Problema do SELECT *

**NUNCA use `SELECT *` em produção**, exceto em casos muito específicos:

```sql
-- ❌ EVITE
SELECT * FROM livros;

-- ✅ PREFIRA
SELECT id, titulo, ano_publicacao FROM livros;
```

**Por quê?**
- Retorna dados desnecessários (mais tráfego de rede)
- Se a estrutura da tabela mudar, sua aplicação pode quebrar
- Mais difícil de entender a intenção da query
- Pior performance em tabelas com muitas colunas

**Quando é aceitável?**
- Durante desenvolvimento e exploração de dados
- Em queries administrativas onde você realmente precisa de tudo
- Em tabelas pequenas com poucas colunas (mas ainda não é ideal)

## 2. Índices: Os Atalhos do Banco de Dados

### O que são Índices?

Um **índice** é como o índice de um livro: em vez de ler o livro inteiro para encontrar um tópico, você vai direto ao índice e encontra a página.

### Como Funcionam os Índices

**Sem índice:**
```
Para encontrar "livros do autor_id = 5":
- Examina TODOS os 100.000 livros
- Verifica cada um: "Este tem autor_id = 5?"
- Retorna os que correspondem
- Tempo: O(n) - linear
```

**Com índice:**
```
Para encontrar "livros do autor_id = 5":
- Vai direto ao índice na seção "autor_id = 5"
- Encontra os IDs dos livros instantaneamente
- Retorna apenas esses
- Tempo: O(log n) - logarítmico (muito mais rápido!)
```

### Quando Usar Índices

**✅ USE ÍNDICES EM:**
- Colunas usadas frequentemente em `WHERE`
- Colunas usadas em `JOIN` (chaves estrangeiras)
- Colunas usadas em `ORDER BY` frequentes
- Colunas com alta seletividade (muitos valores únicos)

**❌ NÃO USE ÍNDICES EM:**
- Tabelas muito pequenas (< 1000 registros geralmente)
- Colunas raramente usadas em queries
- Colunas que mudam constantemente (INSERT/UPDATE frequentes)
- Colunas com baixa seletividade (poucos valores únicos, como "ativo/inativo")

### Exemplos Práticos

```sql
-- Índice em chave estrangeira (muito comum)
CREATE INDEX idx_livros_autor ON livros(autor_id);

-- Índice em coluna usada para busca
CREATE INDEX idx_livros_ano ON livros(ano_publicacao);

-- Índice composto (múltiplas colunas)
CREATE INDEX idx_emprestimos_usuario_status 
ON emprestimos(usuario_id, status);
```

**No nosso banco de dados**, já temos alguns índices criados:
```sql
CREATE INDEX idx_livros_autor ON livros(autor_id);
CREATE INDEX idx_livros_categoria ON livros(categoria_id);
CREATE INDEX idx_emprestimos_livro ON emprestimos(livro_id);
CREATE INDEX idx_emprestimos_usuario ON emprestimos(usuario_id);
CREATE INDEX idx_emprestimos_status ON emprestimos(status);
```

### Verificando o Uso de Índices

No SQLite, você pode usar `EXPLAIN QUERY PLAN`:

```sql
EXPLAIN QUERY PLAN
SELECT * FROM livros WHERE autor_id = 1;
```

Isso mostra se o índice está sendo usado.

## 3. Boas Práticas: Nomenclatura e Organização

### Nomenclatura de Tabelas e Colunas

**✅ BOAS PRÁTICAS:**

1. **Use nomes descritivos e claros**
   ```sql
   -- ✅ Bom
   CREATE TABLE usuarios (...);
   CREATE TABLE emprestimos (...);
   
   -- ❌ Ruim
   CREATE TABLE usr (...);
   CREATE TABLE emp (...);
   ```

2. **Use singular ou plural consistentemente**
   ```sql
   -- ✅ Escolha um padrão e mantenha
   CREATE TABLE usuario (...);  -- singular
   CREATE TABLE livro (...);    -- singular
   
   -- OU
   CREATE TABLE usuarios (...); -- plural
   CREATE TABLE livros (...);    -- plural
   ```

3. **Use snake_case para nomes compostos**
   ```sql
   -- ✅ Bom
   data_emprestimo
   quantidade_disponivel
   data_devolucao_prevista
   
   -- ❌ Ruim (dependendo do SGBD)
   dataEmprestimo
   data-emprestimo
   ```

4. **Evite palavras reservadas do SQL**
   ```sql
   -- ❌ Ruim
   CREATE TABLE user (...);  -- "user" pode ser palavra reservada
   
   -- ✅ Bom
   CREATE TABLE usuarios (...);
   ```

### Organização de Queries

**✅ BOAS PRÁTICAS:**

1. **Formate queries para legibilidade**
   ```sql
   -- ✅ Bom: Fácil de ler
   SELECT 
       l.titulo,
       a.nome AS autor,
       c.nome AS categoria
   FROM livros l
   JOIN autores a ON l.autor_id = a.id
   JOIN categorias c ON l.categoria_id = c.id
   WHERE l.ano_publicacao > 2000
   ORDER BY l.titulo;
   
   -- ❌ Ruim: Difícil de ler
   SELECT l.titulo,a.nome AS autor,c.nome AS categoria FROM livros l JOIN autores a ON l.autor_id=a.id JOIN categorias c ON l.categoria_id=c.id WHERE l.ano_publicacao>2000 ORDER BY l.titulo;
   ```

2. **Use aliases consistentes**
   ```sql
   -- ✅ Bom: Aliases claros
   SELECT 
       l.titulo,
       a.nome AS autor
   FROM livros l
   JOIN autores a ON l.autor_id = a.id;
   
   -- ❌ Ruim: Aliases confusos
   SELECT 
       liv.titulo,
       aut.nome AS autor
   FROM livros liv
   JOIN autores aut ON liv.autor_id = aut.id;
   ```

3. **Comente queries complexas**
   ```sql
   -- Busca livros mais emprestados nos últimos 6 meses
   SELECT 
       l.titulo,
       COUNT(e.id) AS vezes_emprestado
   FROM livros l
   JOIN emprestimos e ON l.id = e.livro_id
   WHERE e.data_emprestimo >= date('now', '-6 months')
   GROUP BY l.id, l.titulo
   ORDER BY vezes_emprestado DESC
   LIMIT 10;
   ```

## 4. Segurança: SQL Injection e Permissões

### SQL Injection: O Perigo Real

**SQL Injection** é uma das vulnerabilidades mais comuns e perigosas em aplicações web.

#### O Problema

```python
# ❌ VULNERÁVEL A SQL INJECTION
query = f"SELECT * FROM usuarios WHERE email = '{email}' AND senha = '{senha}'"

# Se o usuário digitar:
# email = "admin@email.com' OR '1'='1"
# A query vira:
# SELECT * FROM usuarios WHERE email = 'admin@email.com' OR '1'='1' AND senha = ''
# Isso retorna TODOS os usuários!
```

#### A Solução: Prepared Statements / Parameterized Queries

```python
# ✅ SEGURO
query = "SELECT * FROM usuarios WHERE email = ? AND senha = ?"
db.execute(query, (email, senha))
```

**No SQL direto**, sempre valide e sanitize inputs antes de usar em queries.

### Permissões e Controle de Acesso

**Princípio do Menor Privilégio**: Dê apenas as permissões necessárias.

```sql
-- ✅ Bom: Usuário de leitura só pode ler
GRANT SELECT ON livros TO usuario_leitura;
GRANT SELECT ON autores TO usuario_leitura;

-- ✅ Bom: Usuário de escrita pode inserir/atualizar
GRANT SELECT, INSERT, UPDATE ON livros TO usuario_escrita;
GRANT SELECT ON autores TO usuario_escrita;

-- ❌ Ruim: Dar todas as permissões desnecessariamente
GRANT ALL ON * TO usuario_leitura;  -- Por que precisa de DELETE?
```

## 5. Normalização vs Desnormalização

### Normalização: Organizar para Evitar Redundância

**Normalização** é o processo de organizar dados para reduzir redundância e melhorar integridade.

**Exemplo Normalizado (nosso banco):**
```
autores: [id, nome, nacionalidade]
livros: [id, titulo, autor_id]
```

**Vantagens:**
- Sem redundância (nome do autor armazenado uma vez)
- Fácil manutenção (mudar nome do autor em um lugar)
- Consistência garantida
- Menos espaço de armazenamento

**Desvantagens:**
- Mais JOINs necessários
- Queries podem ser mais complexas

### Desnormalização: Quando Aceitar Redundância

Às vezes, **desnormalizar** (aceitar alguma redundância) pode melhorar a performance.

**Exemplo Desnormalizado:**
```
livros: [id, titulo, autor_id, autor_nome]
```

**Quando fazer:**
- Tabelas muito grandes onde JOINs são lentos
- Dados que raramente mudam (como nome de autor)
- Relatórios que precisam de performance extrema
- Dados históricos que não mudam

**⚠️ CUIDADO**: Desnormalização deve ser uma decisão consciente e documentada. Pode causar problemas de consistência se não for gerenciada corretamente.

## 6. Transações e Integridade de Dados

### O que são Transações?

Uma **transação** é uma sequência de operações que devem ser executadas como uma unidade indivisível.

### Propriedades ACID (Revisão)

- **Atomicity**: Tudo ou nada
- **Consistency**: Sempre em estado válido
- **Isolation**: Transações não interferem
- **Durability**: Mudanças são permanentes

### Exemplo Prático

```sql
-- Empréstimo de livro: várias operações devem acontecer juntas
BEGIN TRANSACTION;

-- 1. Verificar se o livro está disponível
SELECT quantidade_disponivel FROM livros WHERE id = 1;

-- 2. Reduzir quantidade disponível
UPDATE livros SET quantidade_disponivel = quantidade_disponivel - 1 WHERE id = 1;

-- 3. Criar registro de empréstimo
INSERT INTO emprestimos (livro_id, usuario_id, data_emprestimo, status) 
VALUES (1, 1, CURRENT_DATE, 'ativo');

-- Se tudo deu certo:
COMMIT;

-- Se algo deu errado:
ROLLBACK;  -- Desfaz tudo
```

**Por que usar transações?**
- Garante que todas as operações aconteçam ou nenhuma
- Evita estados inconsistentes (ex: empréstimo criado mas quantidade não atualizada)
- Permite reverter em caso de erro

## 7. Otimização de Queries: Como Identificar Problemas

### Sinais de Query Lenta

1. **Query demora mais de alguns segundos** (depende do contexto)
2. **Uso excessivo de CPU ou memória**
3. **Bloqueios de tabela durante execução**
4. **Timeouts frequentes**

### Ferramentas de Análise

**SQLite:**
```sql
-- Ver plano de execução
EXPLAIN QUERY PLAN
SELECT * FROM livros WHERE autor_id = 1;

-- Habilitar timer
.timer on
SELECT * FROM livros WHERE autor_id = 1;
```

**Outros SGBDs:**
- MySQL: `EXPLAIN` ou `EXPLAIN ANALYZE`
- PostgreSQL: `EXPLAIN ANALYZE`
- SQL Server: `SET STATISTICS TIME ON`

### Estratégias de Otimização

1. **Adicione índices apropriados**
2. **Evite SELECT ***
3. **Use WHERE para filtrar cedo**
4. **Limite resultados com LIMIT quando apropriado**
5. **Evite funções em colunas do WHERE** (impedem uso de índices)
   ```sql
   -- ❌ Ruim: Função impede uso de índice
   SELECT * FROM livros WHERE YEAR(data_publicacao) = 2020;
   
   -- ✅ Melhor: Comparação direta
   SELECT * FROM livros 
   WHERE data_publicacao >= '2020-01-01' 
   AND data_publicacao < '2021-01-01';
   ```

## 8. Aplicação Prática: Resolvendo Problemas Reais

### Cenário 1: Query Lenta de Relatório

**Problema**: Relatório de "livros mais emprestados" demora 30 segundos.

**Análise:**
```sql
-- Query atual (lenta)
SELECT 
    l.titulo,
    COUNT(e.id) AS vezes_emprestado
FROM livros l
LEFT JOIN emprestimos e ON l.id = e.livro_id
GROUP BY l.id, l.titulo
ORDER BY vezes_emprestado DESC;
```

**Soluções:**
1. Adicionar índice: `CREATE INDEX idx_emprestimos_livro ON emprestimos(livro_id);`
2. Se já existe, verificar se está sendo usado: `EXPLAIN QUERY PLAN`
3. Considerar cachear resultados se os dados não mudam frequentemente
4. Para dados históricos, considerar tabela de agregação pré-calculada

### Cenário 2: Múltiplos Acessos Simultâneos

**Problema**: Sistema trava quando muitos usuários fazem empréstimos ao mesmo tempo.

**Soluções:**
1. Usar transações apropriadas
2. Implementar locks otimistas ou pessimistas
3. Considerar fila de processamento para operações críticas
4. Monitorar deadlocks e ajustar timeouts

### Cenário 3: Crescimento de Dados

**Problema**: Banco está ficando lento conforme cresce.

**Soluções:**
1. Revisar e adicionar índices necessários
2. Arquivar dados antigos (mover empréstimos antigos para tabela de histórico)
3. Considerar particionamento de tabelas grandes
4. Revisar queries e otimizar as mais usadas
5. Monitorar e analisar planos de execução regularmente

## 9. Checklist de Boas Práticas

Use este checklist ao escrever queries:

- [ ] Especifique colunas em vez de `SELECT *`
- [ ] Use `WHERE` para filtrar dados quando possível
- [ ] Use `LIMIT` ao explorar dados grandes
- [ ] Verifique se índices existem para colunas usadas em `WHERE` e `JOIN`
- [ ] Formate queries para legibilidade
- [ ] Use transações para operações que devem ser atômicas
- [ ] Valide e sanitize inputs para prevenir SQL Injection
- [ ] Teste queries em dados de teste antes de produção
- [ ] Documente queries complexas
- [ ] Monitore performance de queries críticas

## 10. Conclusão: Desenvolvendo como Profissional

### Mentalidade Correta

1. **Pense em Escala**: "Esta query funcionaria com 1 milhão de registros?"
2. **Meça, Não Adivinhe**: Use ferramentas de análise antes de otimizar
3. **Documente Decisões**: Por que você desnormalizou? Por que criou esse índice?
4. **Teste Sempre**: Nunca assuma que uma query funciona sem testar
5. **Aprenda Continuamente**: Performance e otimização são áreas em constante evolução

### Próximos Passos

Agora que você entende os fundamentos de SQL e as boas práticas, você está pronto para:
- Aprender comandos SQL específicos (SELECT, JOIN, etc.)
- Praticar com queries mais complexas
- Aplicar esses conceitos em projetos reais

**Lembre-se**: Aprender SQL é uma jornada. Comece com o básico, pratique muito, e sempre pense sobre performance e boas práticas desde o início.

---

**Próximo Passo**: Complete os exercícios da seção anterior e aguarde o feedback e análise do seu desempenho antes de prosseguir para a próxima aula.


