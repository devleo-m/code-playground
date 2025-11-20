# Aula 1 - Exercícios e Reflexão

## Exercícios Práticos

### Exercício 1: Exploração Básica do Banco de Dados

**Objetivo**: Familiarizar-se com a estrutura do banco de dados da biblioteca.

**Tarefas**:

1. Liste todas as tabelas do banco de dados. (Dica: No SQLite, use `.tables` no modo interativo ou consulte a documentação)

2. Examine a estrutura da tabela `livros`. Quantas colunas ela possui? Quais são seus nomes?

3. Conte quantos livros existem no banco de dados.

4. Liste os primeiros 5 livros da tabela, mostrando apenas o título e o ano de publicação.

**Soluções Esperadas**:

```sql
-- 1. Listar tabelas (no SQLite CLI)
.tables

-- 2. Ver estrutura da tabela livros
.schema livros
-- ou
PRAGMA table_info(livros);

-- 3. Contar livros
SELECT COUNT(*) AS total_livros FROM livros;

-- 4. Primeiros 5 livros
SELECT titulo, ano_publicacao 
FROM livros 
LIMIT 5;
```

---

### Exercício 2: Entendendo Relacionamentos

**Objetivo**: Compreender como as tabelas se relacionam.

**Tarefas**:

1. Quantos autores existem no banco de dados?

2. Quantas categorias diferentes temos?

3. Liste todos os autores com suas nacionalidades.

4. Liste todas as categorias com suas descrições.

**Questão de Análise**: 
- Por que faz sentido ter uma tabela separada para autores e outra para livros, em vez de colocar o nome do autor diretamente na tabela de livros?

**Soluções Esperadas**:

```sql
-- 1. Contar autores
SELECT COUNT(*) AS total_autores FROM autores;

-- 2. Contar categorias
SELECT COUNT(*) AS total_categorias FROM categorias;

-- 3. Listar autores
SELECT nome, nacionalidade FROM autores;

-- 4. Listar categorias
SELECT nome, descricao FROM categorias;
```

**Resposta Esperada para a Questão de Análise**:
- Separar em tabelas evita redundância (não repetimos o nome do autor em cada livro)
- Facilita manutenção (se mudarmos o nome de um autor, mudamos em um só lugar)
- Garante consistência (não podemos ter "Isaac Asimov" e "I. Asimov" como autores diferentes)
- Permite adicionar mais informações sobre autores sem modificar a tabela de livros

---

### Exercício 3: Análise de Dados Básica

**Objetivo**: Praticar consultas simples e pensar sobre os dados.

**Tarefas**:

1. Quantos usuários estão cadastrados na biblioteca?

2. Quantos empréstimos existem no total?

3. Quantos empréstimos estão ativos (status = 'ativo')?

4. Liste os nomes e emails de todos os usuários cadastrados.

**Questão de Reflexão**:
- Se a biblioteca tivesse 1 milhão de usuários e você executasse `SELECT * FROM usuarios`, o que aconteceria? Por que isso seria um problema?

**Soluções Esperadas**:

```sql
-- 1. Contar usuários
SELECT COUNT(*) AS total_usuarios FROM usuarios;

-- 2. Contar empréstimos
SELECT COUNT(*) AS total_emprestimos FROM emprestimos;

-- 3. Empréstimos ativos
SELECT COUNT(*) AS emprestimos_ativos 
FROM emprestimos 
WHERE status = 'ativo';

-- 4. Listar usuários
SELECT nome, email FROM usuarios;
```

**Resposta Esperada para a Questão de Reflexão**:
- A query retornaria todos os 1 milhão de registros, o que:
  - Consumiria muita memória
  - Seria muito lento
  - Poderia travar o sistema
  - Na maioria dos casos, você não precisa de TODOS os dados de uma vez
- **Solução**: Sempre usar `LIMIT` ou `WHERE` para filtrar, ou paginar os resultados

---

### Exercício 4: Pensando em Estrutura de Dados

**Objetivo**: Analisar a estrutura do banco e pensar em melhorias.

**Tarefa**:

Examine a estrutura da tabela `emprestimos`:

```sql
PRAGMA table_info(emprestimos);
```

**Perguntas de Reflexão**:

1. Por que a tabela `emprestimos` tem campos `data_devolucao_prevista` e `data_devolucao_real`? Qual a diferença prática entre eles?

2. O campo `status` pode ter apenas dois valores: 'ativo' ou 'devolvido'. Seria melhor usar um campo booleano (TRUE/FALSE) ou manter como texto? Por quê?

3. Se um livro for emprestado 100 vezes, quantos registros serão criados na tabela `emprestimos`? Isso é eficiente?

4. Como você garantiria que um usuário não possa pegar mais de 3 livros emprestados ao mesmo tempo? (Pense na estrutura, não precisa escrever código ainda)

**Respostas Esperadas**:

1. **data_devolucao_prevista**: Quando o livro DEVERIA ser devolvido (definido no momento do empréstimo, geralmente 14 dias depois)
   **data_devolucao_real**: Quando o livro FOI realmente devolvido (pode ser antes, no prazo, ou atrasado)
   Isso permite calcular multas, estatísticas de pontualidade, etc.

2. **Depende do contexto**:
   - Texto ('ativo', 'devolvido'): Mais flexível se no futuro precisarmos de outros status como 'atrasado', 'perdido', 'renovado'
   - Booleano: Mais simples, mas menos flexível
   - **Recomendação**: Para este caso, texto é melhor pela flexibilidade futura

3. **100 registros** - Sim, isso é eficiente e correto! Cada empréstimo é um evento histórico que deve ser registrado. Isso permite:
   - Histórico completo
   - Estatísticas (qual livro é mais emprestado?)
   - Auditoria
   - Análise de padrões de uso

4. **Soluções possíveis**:
   - Criar uma constraint/trigger que conta empréstimos ativos antes de permitir novo empréstimo
   - Validar na aplicação antes de inserir
   - Usar uma view ou função que verifica o limite
   - **Importante**: A validação deve acontecer no banco OU na aplicação, mas de forma consistente

---

## Exercício 5: Análise de Query Existente

**Objetivo**: Analisar uma query e pensar sobre sua eficiência.

Considere a seguinte query:

```sql
SELECT * 
FROM livros 
WHERE ano_publicacao > 2000 
ORDER BY titulo;
```

**Perguntas de Reflexão**:

1. O que esta query faz? Descreva em português.

2. Esta query é eficiente para um banco com 100 livros? E para um banco com 10 milhões de livros?

3. O que aconteceria se não existissem livros publicados depois de 2000? A query daria erro?

4. Se você quisesse apenas os títulos (sem todas as outras colunas), como modificaria a query? Por que isso seria melhor?

**Respostas Esperadas**:

1. **Descrição**: A query retorna todos os livros publicados depois do ano 2000, ordenados alfabeticamente pelo título, mostrando todas as colunas de cada livro.

2. **Eficiência**:
   - **100 livros**: Muito eficiente, executa instantaneamente
   - **10 milhões de livros**: Pode ser lenta porque:
     - Precisa verificar cada registro para o filtro `ano_publicacao > 2000`
     - Precisa ordenar todos os resultados
     - Retorna muitas colunas desnecessárias (`SELECT *`)
   - **Solução**: Usar índices na coluna `ano_publicacao` e especificar apenas as colunas necessárias

3. **Não daria erro**: A query retornaria 0 linhas (resultado vazio), mas executaria normalmente. SQL não dá erro quando não encontra resultados, apenas retorna conjunto vazio.

4. **Query melhorada**:
   ```sql
   SELECT titulo 
   FROM livros 
   WHERE ano_publicacao > 2000 
   ORDER BY titulo;
   ```
   **Por que é melhor**:
   - Menos dados transferidos (apenas título, não todas as colunas)
   - Mais rápido
   - Mais claro sobre a intenção da query
   - Menos uso de memória e rede

---

## Exercício 6: Pensando em Escalabilidade

**Objetivo**: Entender os limites e desafios de bancos relacionais.

**Cenário**: A biblioteca está crescendo rapidamente. Atualmente temos:
- 15 livros
- 10 autores
- 8 usuários
- 10 empréstimos

Mas em 5 anos, podemos ter:
- 100.000 livros
- 5.000 autores
- 50.000 usuários
- 1.000.000 de empréstimos históricos

**Perguntas de Reflexão**:

1. A estrutura atual do banco (as tabelas e relacionamentos) continuaria funcionando bem com esses números? Por quê?

2. Quais operações podem ficar lentas com esse volume de dados? (Pense em consultas comuns como "listar todos os livros")

3. Se você precisasse encontrar rapidamente "todos os livros de um autor específico", o que poderia ser feito para melhorar a performance? (Dica: pense em índices)

4. No contexto de SQL vs NoSQL: este banco de biblioteca seria melhor em SQL ou NoSQL? Justifique.

**Respostas Esperadas**:

1. **Sim, a estrutura continuaria funcionando bem** porque:
   - A normalização (separar autores, categorias, livros) evita redundância
   - Os relacionamentos (chaves estrangeiras) mantêm a integridade
   - A estrutura relacional escala bem verticalmente (mais poder de processamento)
   - SQL é otimizado para esse tipo de consulta estruturada

2. **Operações que podem ficar lentas**:
   - `SELECT * FROM livros` (sem filtros) - precisa ler 100.000 registros
   - Buscar empréstimos antigos sem índices adequados
   - JOINs complexos entre múltiplas tabelas grandes
   - Agregações (COUNT, SUM) em tabelas muito grandes sem índices

3. **Solução: Índices**
   ```sql
   CREATE INDEX idx_livros_autor ON livros(autor_id);
   ```
   Isso cria um "atalho" que permite encontrar rapidamente todos os livros de um autor sem precisar verificar cada linha da tabela.

4. **SQL é melhor para este caso** porque:
   - Dados são estruturados e bem definidos
   - Relacionamentos são importantes (autor-livro, livro-empréstimo)
   - Integridade é crítica (não pode ter empréstimo de livro inexistente)
   - Consultas complexas são necessárias (relatórios, estatísticas)
   - **NoSQL seria melhor se**: precisássemos armazenar dados não estruturados, ou se a escala horizontal fosse o foco principal

---

## Instruções para Resolução

1. **Conecte-se ao banco de dados**:
   ```bash
   sqlite3 biblioteca.db
   ```

2. **Execute cada query** e observe os resultados

3. **Responda as perguntas de reflexão** por escrito, pensando criticamente sobre cada questão

4. **Compare suas respostas** com as soluções esperadas, mas não se preocupe se não acertar tudo - o importante é o processo de pensamento!

5. **Anote suas dúvidas** para discussão posterior

---

## Dicas Importantes

- Sempre teste suas queries antes de executá-las em produção
- Use `LIMIT` quando estiver explorando dados grandes
- Pense sempre: "Esta query seria eficiente com milhões de registros?"
- Entenda o que a query faz antes de executá-la
- Quando possível, especifique colunas em vez de usar `SELECT *`

---

**Próximo Passo**: Após completar os exercícios e reflexões, aguarde o feedback e análise do seu desempenho antes de prosseguir para a seção de Performance e Boas Práticas.


