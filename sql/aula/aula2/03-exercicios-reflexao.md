# Aula 2 - Exercícios e Reflexão

## Exercícios Práticos

### Exercício 1: Consultas SELECT Básicas

**Objetivo**: Praticar o comando SELECT com diferentes filtros e ordenações.

**Tarefas**:

1. Liste todos os títulos dos livros da biblioteca.

2. Liste apenas os títulos e anos de publicação dos livros publicados depois de 2000.

3. Liste os livros que têm quantidade disponível maior que 0, ordenados por título (A-Z).

4. Liste os 3 primeiros livros quando ordenados por ano de publicação (do mais antigo para o mais recente).

5. Liste todos os autores brasileiros, mostrando nome e nacionalidade.

**Soluções Esperadas**:

```sql
-- 1. Todos os títulos
SELECT titulo FROM livros;

-- 2. Livros recentes
SELECT titulo, ano_publicacao 
FROM livros 
WHERE ano_publicacao > 2000;

-- 3. Livros disponíveis ordenados
SELECT titulo, quantidade_disponivel 
FROM livros 
WHERE quantidade_disponivel > 0 
ORDER BY titulo;

-- 4. Primeiros 3 livros mais antigos
SELECT titulo, ano_publicacao 
FROM livros 
WHERE ano_publicacao IS NOT NULL
ORDER BY ano_publicacao ASC 
LIMIT 3;

-- 5. Autores brasileiros
SELECT nome, nacionalidade 
FROM autores 
WHERE nacionalidade = 'Brasileiro';
```

---

### Exercício 2: Operadores de Comparação e Lógicos

**Objetivo**: Praticar o uso de operadores para filtrar dados.

**Tarefas**:

1. Liste livros publicados entre 1990 e 2000 (inclusive).

2. Liste livros que têm quantidade disponível entre 1 e 5 (inclusive).

3. Liste autores que são brasileiros OU americanos.

4. Liste livros que NÃO são de ficção científica E têm estoque disponível.

5. Liste títulos de livros que começam com a letra "D".

**Questão de Reflexão**:
- Qual a diferença entre usar `BETWEEN 1990 AND 2000` e `ano >= 1990 AND ano <= 2000`? Há alguma diferença prática?

**Soluções Esperadas**:

```sql
-- 1. Livros entre 1990 e 2000
SELECT titulo, ano_publicacao 
FROM livros 
WHERE ano_publicacao BETWEEN 1990 AND 2000;

-- 2. Quantidade entre 1 e 5
SELECT titulo, quantidade_disponivel 
FROM livros 
WHERE quantidade_disponivel BETWEEN 1 AND 5;

-- 3. Autores brasileiros ou americanos
SELECT nome, nacionalidade 
FROM autores 
WHERE nacionalidade = 'Brasileiro' 
   OR nacionalidade = 'Americano';

-- 4. Livros não-ficção científica com estoque
SELECT l.titulo, l.quantidade_disponivel 
FROM livros l
JOIN categorias c ON l.categoria_id = c.id
WHERE NOT c.nome = 'Ficção Científica'
  AND l.quantidade_disponivel > 0;

-- 5. Títulos começando com "D"
SELECT titulo 
FROM livros 
WHERE titulo LIKE 'D%';
```

**Resposta Esperada para a Questão de Reflexão**:
- **Não há diferença prática**: `BETWEEN 1990 AND 2000` é equivalente a `>= 1990 AND <= 2000`
- **BETWEEN é inclusivo** nos dois extremos
- **Vantagem do BETWEEN**: Mais legível e conciso
- **Vantagem da forma explícita**: Mais claro sobre a intenção, especialmente quando você quer excluir um dos extremos

---

### Exercício 3: Operadores LIKE e Padrões

**Objetivo**: Praticar busca por padrões em strings.

**Tarefas**:

1. Liste títulos que contêm a palavra "Dom" (em qualquer posição).

2. Liste autores cujo nome contém "Silva".

3. Liste títulos que terminam com a letra "o".

4. Liste títulos que têm exatamente 10 caracteres (use `_` para representar cada caractere).

**Questão de Reflexão**:
- Se você precisasse buscar títulos que contêm o caractere especial `%` (porcentagem), como faria? Por que isso seria necessário?

**Soluções Esperadas**:

```sql
-- 1. Títulos com "Dom"
SELECT titulo 
FROM livros 
WHERE titulo LIKE '%Dom%';

-- 2. Autores com "Silva"
SELECT nome 
FROM autores 
WHERE nome LIKE '%Silva%';

-- 3. Títulos terminando em "o"
SELECT titulo 
FROM livros 
WHERE titulo LIKE '%o';

-- 4. Títulos com exatamente 10 caracteres
SELECT titulo, LENGTH(titulo) AS tamanho
FROM livros 
WHERE titulo LIKE '__________';  -- 10 underscores
-- ou
WHERE LENGTH(titulo) = 10;
```

**Resposta Esperada para a Questão de Reflexão**:
- **Problema**: `%` é um wildcard em LIKE, então `LIKE '%50%'` encontraria qualquer string com "50"
- **Solução**: Escapar o caractere especial usando `ESCAPE`:
  ```sql
  SELECT titulo FROM livros WHERE titulo LIKE '%\%%' ESCAPE '\';
  ```
- **Por que necessário**: Se você armazena dados que podem conter caracteres especiais (como `%`, `_`), precisa escapar para buscar literalmente esses caracteres

---

### Exercício 4: Valores NULL

**Objetivo**: Entender como trabalhar com valores nulos.

**Tarefas**:

1. Liste livros que NÃO têm ano de publicação informado (são NULL).

2. Liste livros que TÊM ano de publicação informado.

3. Liste autores que NÃO têm nacionalidade cadastrada.

4. Liste livros que têm editora informada.

**Questão de Reflexão**:
- Por que não podemos usar `WHERE ano_publicacao = NULL`? Qual a diferença entre NULL e uma string vazia (`''`)?

**Soluções Esperadas**:

```sql
-- 1. Livros sem ano
SELECT titulo, ano_publicacao 
FROM livros 
WHERE ano_publicacao IS NULL;

-- 2. Livros com ano
SELECT titulo, ano_publicacao 
FROM livros 
WHERE ano_publicacao IS NOT NULL;

-- 3. Autores sem nacionalidade
SELECT nome, nacionalidade 
FROM autores 
WHERE nacionalidade IS NULL;

-- 4. Livros com editora
SELECT titulo, editora 
FROM livros 
WHERE editora IS NOT NULL;
```

**Resposta Esperada para a Questão de Reflexão**:
- **Por que não usar `= NULL`**: NULL representa "ausência de valor", não é um valor em si. Qualquer comparação com NULL usando `=` ou `!=` resulta em NULL (não TRUE nem FALSE), então a condição nunca é satisfeita
- **Diferença entre NULL e string vazia**:
  - **NULL**: "Não sabemos" ou "Não foi informado" - ausência de informação
  - **String vazia (`''`)**: "Foi informado, mas está vazio" - é um valor, só que vazio
- **Exemplo prático**: 
  - `telefone IS NULL` = "Telefone não foi cadastrado"
  - `telefone = ''` = "Telefone foi cadastrado como vazio"

---

### Exercício 5: Inserção de Dados (INSERT)

**Objetivo**: Praticar a inserção de novos registros.

**Tarefas**:

1. Insira um novo autor: "Carlos Drummond de Andrade", nacionalidade "Brasileiro", nascido em "1902-10-31".

2. Insira uma nova categoria: "Poesia", com descrição "Livros de poemas e poesias".

3. Insira um novo livro: "A Rosa do Povo", do autor que você acabou de criar, na categoria "Poesia", publicado em 1945, com quantidade disponível 3.

4. Verifique se os dados foram inseridos corretamente fazendo uma consulta SELECT.

**⚠️ IMPORTANTE**: Anote os IDs retornados ou use subconsultas para encontrar os IDs automaticamente.

**Soluções Esperadas**:

```sql
-- 1. Inserir autor
INSERT INTO autores (nome, nacionalidade, data_nascimento)
VALUES ('Carlos Drummond de Andrade', 'Brasileiro', '1902-10-31');

-- 2. Inserir categoria
INSERT INTO categorias (nome, descricao)
VALUES ('Poesia', 'Livros de poemas e poesias');

-- 3. Inserir livro (usando subconsultas para encontrar IDs)
INSERT INTO livros (titulo, autor_id, categoria_id, ano_publicacao, quantidade_disponivel)
VALUES (
    'A Rosa do Povo',
    (SELECT id FROM autores WHERE nome = 'Carlos Drummond de Andrade'),
    (SELECT id FROM categorias WHERE nome = 'Poesia'),
    1945,
    3
);

-- 4. Verificar inserção
SELECT 
    l.titulo,
    a.nome AS autor,
    c.nome AS categoria,
    l.ano_publicacao,
    l.quantidade_disponivel
FROM livros l
JOIN autores a ON l.autor_id = a.id
JOIN categorias c ON l.categoria_id = c.id
WHERE l.titulo = 'A Rosa do Povo';
```

---

### Exercício 6: Atualização de Dados (UPDATE)

**Objetivo**: Praticar a modificação de registros existentes.

**⚠️ CRÍTICO**: Sempre teste com SELECT antes de fazer UPDATE!

**Tarefas**:

1. **PRIMEIRO**: Liste todos os livros que têm quantidade disponível igual a 0.

2. Atualize a quantidade disponível do livro "A Rosa do Povo" para 5.

3. Aumente a quantidade disponível de todos os livros de ficção científica em 2 unidades.

4. Atualize o ano de publicação de todos os livros que têm ano NULL para 2000 (apenas para teste - em produção, isso não faria sentido).

5. **VERIFICAÇÃO**: Liste novamente os livros atualizados para confirmar as mudanças.

**Questão de Reflexão**:
- Por que é crucial usar WHERE em UPDATE? O que aconteceria se você executasse `UPDATE livros SET quantidade_disponivel = 0;` sem WHERE?

**Soluções Esperadas**:

```sql
-- 1. Ver livros sem estoque (ANTES de atualizar)
SELECT titulo, quantidade_disponivel 
FROM livros 
WHERE quantidade_disponivel = 0;

-- 2. Atualizar livro específico
UPDATE livros
SET quantidade_disponivel = 5
WHERE titulo = 'A Rosa do Povo';

-- 3. Aumentar estoque de ficção científica
UPDATE livros
SET quantidade_disponivel = quantidade_disponivel + 2
WHERE categoria_id = (SELECT id FROM categorias WHERE nome = 'Ficção Científica');

-- 4. Atualizar anos NULL (apenas para exercício)
UPDATE livros
SET ano_publicacao = 2000
WHERE ano_publicacao IS NULL;

-- 5. Verificar mudanças
SELECT titulo, quantidade_disponivel, ano_publicacao 
FROM livros 
ORDER BY titulo;
```

**Resposta Esperada para a Questão de Reflexão**:
- **Sem WHERE**: `UPDATE livros SET quantidade_disponivel = 0;` atualizaria **TODOS** os registros da tabela, zerando o estoque de todos os livros
- **Impacto**: Catastrófico! Perda de dados, sistema inutilizável
- **Por que crucial**: WHERE é o "alvo" que especifica quais registros modificar
- **Boas práticas**:
  1. Sempre teste com SELECT primeiro: `SELECT * FROM livros WHERE ...`
  2. Use transações quando possível (para poder reverter)
  3. Faça backup antes de operações em massa
  4. Use WHERE mesmo quando "tem certeza" que quer todos

---

### Exercício 7: Remoção de Dados (DELETE)

**Objetivo**: Praticar a remoção de registros (com muito cuidado!).

**⚠️ CRÍTICO**: Sempre teste com SELECT antes de fazer DELETE!

**Tarefas**:

1. **PRIMEIRO**: Liste todos os empréstimos que têm status 'devolvido' e foram devolvidos antes de 2020.

2. Delete os empréstimos devolvidos antes de 2020 (apenas para limpeza de dados antigos - em produção, você normalmente manteria o histórico).

3. **PRIMEIRO**: Liste o livro que você inseriu no Exercício 5 ("A Rosa do Povo").

4. Delete o livro "A Rosa do Povo" que você inseriu anteriormente.

5. **VERIFICAÇÃO**: Confirme que o livro foi removido.

**Questão de Reflexão**:
- Se você tentasse deletar um autor que tem livros associados (com chave estrangeira), o que aconteceria? Por que isso é uma proteção importante?

**Soluções Esperadas**:

```sql
-- 1. Ver empréstimos antigos (ANTES de deletar)
SELECT * 
FROM emprestimos 
WHERE status = 'devolvido' 
  AND data_devolucao_real < '2020-01-01';

-- 2. Deletar empréstimos antigos
DELETE FROM emprestimos
WHERE status = 'devolvido' 
  AND data_devolucao_real < '2020-01-01';

-- 3. Ver livro antes de deletar
SELECT * FROM livros WHERE titulo = 'A Rosa do Povo';

-- 4. Deletar livro
DELETE FROM livros
WHERE titulo = 'A Rosa do Povo';

-- 5. Verificar remoção
SELECT * FROM livros WHERE titulo = 'A Rosa do Povo';
-- Deve retornar 0 linhas
```

**Resposta Esperada para a Questão de Reflexão**:
- **O que aconteceria**: O banco de dados retornaria um erro de **FOREIGN KEY constraint failed**
- **Por que**: Existem livros que referenciam esse autor através de `autor_id`. Se você deletasse o autor, os livros ficariam "órfãos" (referenciando um autor que não existe mais)
- **Por que é proteção importante**:
  - **Integridade referencial**: Garante que os dados estão consistentes
  - **Previne dados órfãos**: Evita registros que referenciam dados inexistentes
  - **Proteção acidental**: Evita deletar dados importantes por engano
- **Soluções possíveis**:
  1. Deletar os livros primeiro, depois o autor
  2. Usar CASCADE (deletar livros automaticamente quando deletar autor)
  3. Marcar como "inativo" em vez de deletar (soft delete)

---

### Exercício 8: Análise de Query Existente

**Objetivo**: Analisar queries e pensar sobre eficiência e correção.

Considere a seguinte query:

```sql
SELECT * 
FROM livros 
WHERE titulo LIKE '%Dom%' 
ORDER BY titulo;
```

**Perguntas de Reflexão**:

1. O que esta query faz? Descreva em português.

2. Esta query é eficiente? Por quê?

3. Se a tabela `livros` tivesse 1 milhão de registros, o que aconteceria com esta query?

4. Como você melhoraria esta query? (Pense em índices, especificação de colunas, etc.)

5. Qual a diferença entre `LIKE '%Dom%'` e `LIKE 'Dom%'`? Qual seria mais eficiente e por quê?

**Respostas Esperadas**:

1. **Descrição**: A query retorna todos os livros cujo título contém a palavra "Dom" em qualquer posição, ordenados alfabeticamente pelo título, mostrando todas as colunas.

2. **Eficiência**: 
   - **Pequenos volumes**: Eficiente o suficiente
   - **Grandes volumes**: Pode ser lenta porque:
     - `LIKE '%Dom%'` (com % no início) não pode usar índices eficientemente
     - Precisa verificar cada registro
     - `SELECT *` retorna todas as colunas desnecessariamente

3. **Com 1 milhão de registros**:
   - Seria muito lenta (pode levar vários segundos ou minutos)
   - Consumiria muita memória
   - Poderia travar o sistema se muitos usuários executassem simultaneamente
   - O banco precisaria fazer "full table scan" (verificar todos os registros)

4. **Melhorias**:
   ```sql
   -- Especificar colunas necessárias
   SELECT titulo, ano_publicacao 
   FROM livros 
   WHERE titulo LIKE '%Dom%' 
   ORDER BY titulo;
   
   -- Se possível, mudar para busca que começa com o padrão
   SELECT titulo, ano_publicacao 
   FROM livros 
   WHERE titulo LIKE 'Dom%' 
   ORDER BY titulo;
   
   -- Criar índice para busca de títulos (se busca por prefixo)
   CREATE INDEX idx_livros_titulo ON livros(titulo);
   ```

5. **Diferença**:
   - `LIKE '%Dom%'`: Busca "Dom" em qualquer posição (início, meio ou fim)
   - `LIKE 'Dom%'`: Busca apenas títulos que começam com "Dom"
   - **Mais eficiente**: `LIKE 'Dom%'` porque:
     - Pode usar índices (busca por prefixo)
     - O banco pode parar de buscar quando encontra algo que não começa com "Dom"
     - Muito mais rápido em grandes volumes

---

### Exercício 9: Pensando em Integridade de Dados

**Objetivo**: Entender a importância de manter dados consistentes.

**Cenário**: Você precisa atualizar o nome de um autor de "Machado de Assis" para "Machado de Assis (Joaquim Maria)".

**Perguntas de Reflexão**:

1. Se você simplesmente atualizar o nome na tabela `autores`, o que acontece com os livros que referenciam esse autor?

2. Por que isso é uma vantagem do modelo relacional?

3. Se você tivesse o nome do autor armazenado diretamente na tabela `livros` (sem tabela separada), quantas atualizações seriam necessárias?

4. Qual abordagem é mais eficiente e por quê?

**Respostas Esperadas**:

1. **O que acontece**: Os livros que referenciam esse autor (através de `autor_id`) automaticamente "herdam" a mudança, porque eles referenciam o ID, não o nome diretamente. Quando você consulta `livros JOIN autores`, o novo nome aparece automaticamente.

2. **Vantagem do modelo relacional**:
   - **Atualização única**: Você atualiza em um só lugar
   - **Consistência automática**: Todos os livros mostram o nome atualizado
   - **Sem redundância**: Não precisa atualizar em múltiplos lugares
   - **Menos erros**: Impossível ter nomes diferentes para o mesmo autor

3. **Se nome estivesse em `livros`**:
   - Precisaria atualizar **cada registro** de livro desse autor
   - Se houver 10 livros, seriam 10 UPDATEs
   - Risco de esquecer algum livro
   - Risco de inconsistência (alguns com nome antigo, outros com nome novo)

4. **Abordagem relacional é mais eficiente**:
   - **1 UPDATE** vs **N UPDATEs**
   - **Garantia de consistência** (impossível ter dados inconsistentes)
   - **Menos trabalho manual**
   - **Melhor performance** (uma operação vs múltiplas)

---

### Exercício 10: Cenário Prático Completo

**Objetivo**: Aplicar todos os conceitos aprendidos em um cenário real.

**Cenário**: A biblioteca recebeu uma doação de 3 novos livros do mesmo autor. Você precisa:

1. Verificar se o autor já existe no banco. Se não existir, cadastrá-lo.
2. Cadastrar os 3 novos livros.
3. Atualizar a quantidade disponível de cada livro para 5.
4. Listar todos os livros deste autor para verificação.

**Dados**:
- Autor: "J.K. Rowling", nacionalidade "Britânica"
- Livros:
  - "Harry Potter e a Pedra Filosofal" (1997, categoria: Ficção Científica ou similar)
  - "Harry Potter e a Câmara Secreta" (1998, mesma categoria)
  - "Harry Potter e o Prisioneiro de Azkaban" (1999, mesma categoria)

**Tarefa**: Escreva as queries SQL necessárias para completar este cenário.

**Solução Esperada**:

```sql
-- 1. Verificar se autor existe
SELECT id, nome FROM autores WHERE nome = 'J.K. Rowling';

-- 2. Se não existir, cadastrar (ajuste o ID da categoria conforme seu banco)
INSERT INTO autores (nome, nacionalidade)
VALUES ('J.K. Rowling', 'Britânica');

-- 3. Cadastrar os 3 livros
INSERT INTO livros (titulo, autor_id, categoria_id, ano_publicacao, quantidade_disponivel)
VALUES 
    ('Harry Potter e a Pedra Filosofal', 
     (SELECT id FROM autores WHERE nome = 'J.K. Rowling'),
     (SELECT id FROM categorias WHERE nome = 'Ficção Científica'),
     1997,
     5),
    ('Harry Potter e a Câmara Secreta',
     (SELECT id FROM autores WHERE nome = 'J.K. Rowling'),
     (SELECT id FROM categorias WHERE nome = 'Ficção Científica'),
     1998,
     5),
    ('Harry Potter e o Prisioneiro de Azkaban',
     (SELECT id FROM autores WHERE nome = 'J.K. Rowling'),
     (SELECT id FROM categorias WHERE nome = 'Ficção Científica'),
     1999,
     5);

-- 4. Verificar livros do autor
SELECT 
    l.titulo,
    l.ano_publicacao,
    l.quantidade_disponivel,
    a.nome AS autor
FROM livros l
JOIN autores a ON l.autor_id = a.id
WHERE a.nome = 'J.K. Rowling'
ORDER BY l.ano_publicacao;
```

---

## Instruções para Resolução

1. **Conecte-se ao banco de dados**:
   ```bash
   sqlite3 biblioteca.db
   ```

2. **Execute cada query** e observe os resultados cuidadosamente.

3. **Para exercícios de INSERT/UPDATE/DELETE**:
   - **SEMPRE** teste com SELECT primeiro
   - Verifique os dados antes de modificar
   - Confirme as mudanças após executar

4. **Responda as perguntas de reflexão** por escrito, pensando criticamente sobre cada questão.

5. **Anote suas dúvidas** para discussão posterior.

6. **Se cometer erros**: Não se preocupe! Erros fazem parte do aprendizado. Se necessário, recrie o banco de dados executando `go run init_database.go`.

---

## Dicas Importantes

- **Sempre use WHERE** em UPDATE e DELETE
- **Sempre teste com SELECT** antes de modificar dados
- **Pense sobre performance**: "Esta query seria eficiente com milhões de registros?"
- **Especifique colunas**: Evite `SELECT *` quando possível
- **Entenda NULL**: Use `IS NULL` e `IS NOT NULL`, nunca `= NULL`
- **Cuidado com LIKE**: `LIKE '%texto%'` é mais lento que `LIKE 'texto%'`

---

**Próximo Passo**: Após completar todos os exercícios e responder as perguntas de reflexão, aguarde o feedback e análise do seu desempenho antes de prosseguir para a seção de Performance e Boas Práticas.

