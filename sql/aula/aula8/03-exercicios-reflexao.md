# Aula 8 - Exercícios e Reflexão: Advanced SQL Functions

## Exercícios Práticos

### Exercício 1: Trabalhando com String Functions

**Objetivo**: Praticar manipulação de strings usando funções CONCAT, LENGTH, SUBSTRING, REPLACE, UPPER e LOWER.

**Tarefas**:

1. Crie uma query que liste todos os livros com:
   - Título em maiúsculas
   - Primeiras 10 letras do título
   - Título sem espaços (substitua espaços por underscores)

2. Crie uma query que mostre:
   - Nome completo do usuário (nome + email) usando concatenação
   - Tamanho do email de cada usuário
   - Email em minúsculas

3. Crie uma query que encontre todos os títulos com mais de 30 caracteres.

4. Crie uma query que normalize os emails (converta para minúsculas) e mostre o domínio (parte após o @).

**Questão de Reflexão**:
- Quando você usaria funções de string no banco de dados ao invés de processar na aplicação? Quais são as vantagens e desvantagens de cada abordagem?

**Soluções Esperadas**:

```sql
-- 1. Manipular títulos
SELECT 
    titulo,
    UPPER(titulo) AS titulo_maiusculo,
    SUBSTR(titulo, 1, 10) AS primeiras_10_letras,
    REPLACE(titulo, ' ', '_') AS titulo_sem_espacos
FROM livros;

-- 2. Manipular dados de usuários
SELECT 
    nome || ' - ' || email AS nome_completo,
    LENGTH(email) AS tamanho_email,
    LOWER(email) AS email_minusculo
FROM usuarios;

-- 3. Encontrar títulos longos
SELECT 
    titulo,
    LENGTH(titulo) AS tamanho
FROM livros
WHERE LENGTH(titulo) > 30;

-- 4. Normalizar emails e extrair domínio
SELECT 
    nome,
    LOWER(email) AS email_normalizado,
    SUBSTR(email, INSTR(email, '@') + 1) AS dominio
FROM usuarios;
```

**Resposta Esperada para a Questão de Reflexão**:
- **Usar funções no banco quando**:
  - Precisa filtrar ou ordenar por valores transformados
  - Quer aproveitar índices em colunas transformadas
  - Precisa de consistência na transformação
  - Quer reduzir processamento na aplicação
- **Processar na aplicação quando**:
  - Transformação é muito complexa
  - Precisa de lógica de negócio específica da aplicação
  - Quer flexibilidade para mudar transformação sem alterar queries
- **Vantagens do banco**: Performance, consistência, aproveitamento de índices
- **Desvantagens do banco**: Menos flexível, pode ser mais difícil de manter

---

### Exercício 2: Trabalhando com Date & Time Functions

**Objetivo**: Praticar manipulação de datas usando funções DATE, strftime e cálculos de intervalo.

**Tarefas**:

1. Crie uma query que mostre:
   - Data de empréstimo
   - Ano do empréstimo
   - Mês do empréstimo
   - Dia da semana do empréstimo

2. Crie uma query que calcule:
   - Data de devolução prevista (15 dias após empréstimo)
   - Quantos dias se passaram desde o empréstimo
   - Quantos dias faltam para a devolução (se ainda não venceu)

3. Crie uma query que agrupe empréstimos por ano e mostre o total de cada ano.

4. Crie uma query que encontre todos os empréstimos que vencem nos próximos 7 dias.

**Questão de Reflexão**:
- Por que é importante trabalhar com datas no banco de dados? Quais problemas podem surgir ao processar datas na aplicação?

**Soluções Esperadas**:

```sql
-- 1. Extrair partes da data
SELECT 
    data_emprestimo,
    strftime('%Y', data_emprestimo) AS ano,
    strftime('%m', data_emprestimo) AS mes,
    CASE strftime('%w', data_emprestimo)
        WHEN '0' THEN 'Domingo'
        WHEN '1' THEN 'Segunda'
        WHEN '2' THEN 'Terça'
        WHEN '3' THEN 'Quarta'
        WHEN '4' THEN 'Quinta'
        WHEN '5' THEN 'Sexta'
        WHEN '6' THEN 'Sábado'
    END AS dia_semana
FROM emprestimos;

-- 2. Calcular datas e intervalos
SELECT 
    id,
    data_emprestimo,
    date(data_emprestimo, '+15 days') AS data_devolucao_calculada,
    ROUND(julianday('now') - julianday(data_emprestimo)) AS dias_decorridos,
    CASE
        WHEN data_devolucao_prevista IS NOT NULL THEN
            ROUND(julianday(data_devolucao_prevista) - julianday('now'))
        ELSE NULL
    END AS dias_restantes
FROM emprestimos
WHERE status = 'ativo';

-- 3. Agrupar por ano
SELECT 
    strftime('%Y', data_emprestimo) AS ano,
    COUNT(*) AS total_emprestimos
FROM emprestimos
GROUP BY ano
ORDER BY ano DESC;

-- 4. Empréstimos vencendo em 7 dias
SELECT 
    id,
    data_emprestimo,
    data_devolucao_prevista,
    ROUND(julianday(data_devolucao_prevista) - julianday('now')) AS dias_restantes
FROM emprestimos
WHERE status = 'ativo'
  AND data_devolucao_prevista IS NOT NULL
  AND julianday(data_devolucao_prevista) - julianday('now') <= 7
  AND julianday(data_devolucao_prevista) - julianday('now') >= 0;
```

**Resposta Esperada para a Questão de Reflexão**:
- **Importância de trabalhar com datas no banco**:
  - Consistência de timezone e formato
  - Performance em filtros e ordenações
  - Validação de datas (evita datas inválidas)
  - Cálculos precisos de intervalos
- **Problemas ao processar na aplicação**:
  - Diferenças de timezone entre servidor e cliente
  - Formato inconsistente
  - Cálculos imprecisos
  - Performance ruim em grandes volumes
  - Dificuldade em filtrar por período

---

### Exercício 3: Trabalhando com Numeric Functions

**Objetivo**: Praticar cálculos numéricos usando funções FLOOR, ABS, MOD, ROUND e CEILING.

**Tarefas**:

1. Crie uma query que mostre:
   - Quantidade disponível
   - Quantidade arredondada para inteiro
   - Quantidade arredondada para cima
   - Quantidade arredondada para baixo

2. Crie uma query que classifique livros como "Par" ou "Ímpar" baseado na quantidade disponível.

3. Crie uma query que calcule:
   - Média de estoque por categoria (arredondada para 2 casas decimais)
   - Diferença entre quantidade de cada livro e a média da categoria

4. Crie uma query que mostre a porcentagem que cada livro representa do estoque total (arredondada para 2 casas decimais).

**Questão de Reflexão**:
- Quando você usaria ROUND, FLOOR ou CEILING? Dê exemplos práticos de situações onde cada um seria mais apropriado.

**Soluções Esperadas**:

```sql
-- 1. Arredondamentos
SELECT 
    titulo,
    quantidade_disponivel,
    ROUND(quantidade_disponivel) AS arredondado,
    CEIL(quantidade_disponivel) AS arredondado_cima,
    FLOOR(quantidade_disponivel) AS arredondado_baixo
FROM livros;

-- 2. Classificar par/ímpar
SELECT 
    titulo,
    quantidade_disponivel,
    CASE 
        WHEN quantidade_disponivel % 2 = 0 THEN 'Par'
        ELSE 'Ímpar'
    END AS paridade
FROM livros;

-- 3. Média e diferença por categoria
SELECT 
    l.titulo,
    c.nome AS categoria,
    l.quantidade_disponivel,
    ROUND(AVG(l2.quantidade_disponivel) OVER (PARTITION BY c.id), 2) AS media_categoria,
    ROUND(l.quantidade_disponivel - AVG(l2.quantidade_disponivel) OVER (PARTITION BY c.id), 2) AS diferenca_media
FROM livros l
JOIN categorias c ON l.categoria_id = c.id
JOIN livros l2 ON c.id = l2.categoria_id
GROUP BY l.id, l.titulo, c.nome, l.quantidade_disponivel;

-- Alternativa mais simples (sem window functions):
SELECT 
    l.titulo,
    c.nome AS categoria,
    l.quantidade_disponivel,
    ROUND((SELECT AVG(quantidade_disponivel) FROM livros WHERE categoria_id = l.categoria_id), 2) AS media_categoria,
    ROUND(l.quantidade_disponivel - (SELECT AVG(quantidade_disponivel) FROM livros WHERE categoria_id = l.categoria_id), 2) AS diferenca_media
FROM livros l
JOIN categorias c ON l.categoria_id = c.id;

-- 4. Porcentagem do estoque total
SELECT 
    titulo,
    quantidade_disponivel,
    ROUND(
        (quantidade_disponivel * 100.0 / (SELECT SUM(quantidade_disponivel) FROM livros)), 
        2
    ) AS porcentagem_estoque
FROM livros
WHERE quantidade_disponivel > 0;
```

**Resposta Esperada para a Questão de Reflexão**:
- **ROUND**: Quando precisa do valor mais próximo
  - Exemplo: Preços, médias, porcentagens
  - "R$ 19.99 arredondado = R$ 20"
- **FLOOR**: Quando precisa garantir que não ultrapasse um limite
  - Exemplo: Capacidade (não pode ter mais que X), divisões que não podem exceder
  - "23 livros, 10 por prateleira = 2 prateleiras completas (FLOOR)"
- **CEILING**: Quando precisa garantir que tenha espaço suficiente
  - Exemplo: Recursos necessários, capacidade mínima
  - "23 livros, 10 por prateleira = 3 prateleiras necessárias (CEIL)"

---

### Exercício 4: Trabalhando com Conditional Functions

**Objetivo**: Praticar lógica condicional usando CASE, NULLIF e COALESCE.

**Tarefas**:

1. Crie uma query que classifique livros por nível de estoque:
   - "Esgotado" se quantidade = 0
   - "Estoque Baixo" se quantidade < 5
   - "Estoque Médio" se quantidade < 10
   - "Estoque Alto" se quantidade >= 10

2. Crie uma query que classifique empréstimos:
   - "Atrasado" se status = 'ativo' e data de devolução já passou
   - "Vencendo" se status = 'ativo' e vence em até 3 dias
   - "No Prazo" se status = 'ativo' e ainda tem mais de 3 dias
   - "Devolvido" se status = 'devolvido'

3. Crie uma query que use COALESCE para mostrar:
   - Telefone do usuário ou "Não informado" se NULL
   - Data de devolução real ou data prevista ou "Não definida"

4. Crie uma query que use NULLIF para tratar valores específicos como NULL:
   - Editora "Desconhecida" deve ser tratada como NULL
   - Telefone vazio ('') deve ser tratado como NULL

**Questão de Reflexão**:
- Qual a diferença entre CASE, NULLIF e COALESCE? Quando você usaria cada um? Dê exemplos práticos.

**Soluções Esperadas**:

```sql
-- 1. Classificar estoque
SELECT 
    titulo,
    quantidade_disponivel,
    CASE
        WHEN quantidade_disponivel = 0 THEN 'Esgotado'
        WHEN quantidade_disponivel < 5 THEN 'Estoque Baixo'
        WHEN quantidade_disponivel < 10 THEN 'Estoque Médio'
        ELSE 'Estoque Alto'
    END AS nivel_estoque
FROM livros
ORDER BY quantidade_disponivel DESC;

-- 2. Classificar empréstimos
SELECT 
    id,
    data_emprestimo,
    data_devolucao_prevista,
    status,
    CASE
        WHEN status = 'ativo' AND julianday('now') > julianday(data_devolucao_prevista) THEN 'Atrasado'
        WHEN status = 'ativo' AND julianday(data_devolucao_prevista) - julianday('now') <= 3 THEN 'Vencendo'
        WHEN status = 'ativo' THEN 'No Prazo'
        WHEN status = 'devolvido' THEN 'Devolvido'
        ELSE 'Desconhecido'
    END AS status_detalhado
FROM emprestimos;

-- 3. Usar COALESCE para valores padrão
SELECT 
    nome,
    COALESCE(telefone, 'Não informado') AS telefone_display,
    COALESCE(data_devolucao_real, data_devolucao_prevista, 'Não definida') AS data_final
FROM usuarios u
LEFT JOIN emprestimos e ON u.id = e.usuario_id;

-- 4. Usar NULLIF para tratar valores específicos
SELECT 
    titulo,
    NULLIF(editora, 'Desconhecida') AS editora_limpa,
    NULLIF(telefone, '') AS telefone_ou_null
FROM livros l
CROSS JOIN usuarios u
WHERE u.telefone IS NOT NULL;
```

**Resposta Esperada para a Questão de Reflexão**:
- **CASE**: Lógica condicional complexa com múltiplas condições
  - Exemplo: Classificar estoque em várias categorias
  - "Se X então A, se Y então B, senão C"
- **NULLIF**: Transformar valores específicos em NULL
  - Exemplo: Tratar "Desconhecida" como NULL
  - "Se valor = X então NULL, senão valor"
- **COALESCE**: Escolher primeiro valor não-NULL de uma lista
  - Exemplo: Usar valor padrão se NULL
  - "Primeiro não-NULL de [valor1, valor2, padrão]"
- **Quando usar cada um**:
  - CASE: Múltiplas condições, classificações complexas
  - NULLIF: Limpeza de dados, normalização
  - COALESCE: Valores padrão, priorização de valores

---

### Exercício 5: Combinando Funções

**Objetivo**: Praticar combinação de diferentes tipos de funções em queries complexas.

**Tarefas**:

1. Crie uma query que gere um código de referência para cada livro no formato:
   - Primeiras 3 letras do título em maiúsculas
   - Hífen
   - ID do livro
   - Exemplo: "FUN-1" para "Fundação" com ID 1

2. Crie uma query que mostre uma descrição completa do empréstimo:
   - Nome do usuário em maiúsculas
   - Título do livro
   - Status formatado (usando CASE)
   - Dias de atraso (se atrasado) ou dias restantes (se no prazo)

3. Crie uma query que analise categorias:
   - Código da categoria (primeiras 3 letras em maiúsculas)
   - Total de livros (arredondado)
   - Média de estoque (arredondada para 2 casas)
   - Classificação de tamanho (Grande/Média/Pequena baseado no total)

4. Crie uma query que normalize e limpe dados:
   - Títulos sem espaços extras (substituir múltiplos espaços por um)
   - Emails em minúsculas
   - Nacionalidades padronizadas (primeira letra maiúscula, resto minúscula)

**Questão de Reflexão**:
- Quais são os desafios de combinar múltiplas funções em uma query? Como você pode tornar queries complexas mais legíveis e manuteníveis?

**Soluções Esperadas**:

```sql
-- 1. Código de referência
SELECT 
    titulo,
    id,
    UPPER(SUBSTR(titulo, 1, 3)) || '-' || id AS codigo_referencia
FROM livros;

-- 2. Descrição completa de empréstimo
SELECT 
    UPPER(u.nome) AS usuario_maiusculo,
    l.titulo AS livro,
    CASE
        WHEN e.status = 'ativo' AND julianday('now') > julianday(e.data_devolucao_prevista) THEN 
            'Atrasado - ' || ROUND(julianday('now') - julianday(e.data_devolucao_prevista)) || ' dias'
        WHEN e.status = 'ativo' THEN 
            'No Prazo - ' || ROUND(julianday(e.data_devolucao_prevista) - julianday('now')) || ' dias restantes'
        ELSE 'Devolvido'
    END AS status_formatado
FROM emprestimos e
JOIN usuarios u ON e.usuario_id = u.id
JOIN livros l ON e.livro_id = l.id;

-- 3. Análise de categorias
SELECT 
    UPPER(SUBSTR(c.nome, 1, 3)) AS codigo_categoria,
    ROUND(COUNT(l.id)) AS total_livros,
    ROUND(AVG(l.quantidade_disponivel), 2) AS media_estoque,
    CASE
        WHEN COUNT(l.id) > 5 THEN 'Grande'
        WHEN COUNT(l.id) > 2 THEN 'Média'
        ELSE 'Pequena'
    END AS tamanho_categoria
FROM categorias c
LEFT JOIN livros l ON c.id = l.categoria_id
GROUP BY c.id, c.nome;

-- 4. Normalizar e limpar dados
SELECT 
    REPLACE(REPLACE(titulo, '  ', ' '), '  ', ' ') AS titulo_limpo,
    LOWER(email) AS email_normalizado,
    UPPER(SUBSTR(nacionalidade, 1, 1)) || LOWER(SUBSTR(nacionalidade, 2)) AS nacionalidade_formatada
FROM livros l
CROSS JOIN usuarios u
CROSS JOIN autores a
WHERE a.nacionalidade IS NOT NULL
LIMIT 10;  -- Limitar para exemplo
```

**Resposta Esperada para a Questão de Reflexão**:
- **Desafios de combinar funções**:
  - Queries podem ficar difíceis de ler
  - Difícil de debugar quando há erro
  - Performance pode ser afetada
  - Manutenção se torna complexa
- **Como tornar mais legível**:
  - Use aliases descritivos
  - Quebre em múltiplas queries se necessário
  - Adicione comentários explicativos
  - Teste cada função separadamente antes de combinar
  - Considere usar CTEs (Common Table Expressions) para queries muito complexas
  - Documente a lógica de negócio

---

### Exercício 6: Análise Prática Completa

**Objetivo**: Criar uma análise completa usando todas as categorias de funções aprendidas.

**Tarefas**:

Crie um relatório completo de empréstimos que mostre:

1. **Informações do Empréstimo**:
   - ID do empréstimo
   - Data de empréstimo formatada (DD/MM/YYYY)
   - Data de devolução prevista formatada

2. **Informações do Usuário**:
   - Nome completo (nome + email entre parênteses)
   - Telefone ou "Não informado"

3. **Informações do Livro**:
   - Título em maiúsculas
   - Categoria
   - Status de estoque (Esgotado/Baixo/Médio/Alto)

4. **Análise Temporal**:
   - Ano do empréstimo
   - Mês do empréstimo (nome do mês)
   - Dias decorridos desde o empréstimo
   - Status de prazo (Atrasado/Vencendo/No Prazo/Devolvido)

5. **Estatísticas**:
   - Total de empréstimos por usuário
   - Média de dias de empréstimo por categoria

**Questão de Reflexão**:
- Como funções SQL ajudam a criar relatórios e análises? Quais são as vantagens de fazer transformações no banco ao invés de na aplicação para relatórios?

**Solução Esperada**:

```sql
-- Relatório completo de empréstimos
SELECT 
    -- Informações do Empréstimo
    e.id AS emprestimo_id,
    strftime('%d/%m/%Y', e.data_emprestimo) AS data_emprestimo_formatada,
    CASE 
        WHEN e.data_devolucao_prevista IS NOT NULL 
        THEN strftime('%d/%m/%Y', e.data_devolucao_prevista)
        ELSE 'Não definida'
    END AS data_devolucao_formatada,
    
    -- Informações do Usuário
    u.nome || ' (' || u.email || ')' AS usuario_completo,
    COALESCE(u.telefone, 'Não informado') AS telefone,
    
    -- Informações do Livro
    UPPER(l.titulo) AS titulo_maiusculo,
    c.nome AS categoria,
    CASE
        WHEN l.quantidade_disponivel = 0 THEN 'Esgotado'
        WHEN l.quantidade_disponivel < 5 THEN 'Baixo'
        WHEN l.quantidade_disponivel < 10 THEN 'Médio'
        ELSE 'Alto'
    END AS status_estoque,
    
    -- Análise Temporal
    strftime('%Y', e.data_emprestimo) AS ano_emprestimo,
    CASE strftime('%m', e.data_emprestimo)
        WHEN '01' THEN 'Janeiro'
        WHEN '02' THEN 'Fevereiro'
        WHEN '03' THEN 'Março'
        WHEN '04' THEN 'Abril'
        WHEN '05' THEN 'Maio'
        WHEN '06' THEN 'Junho'
        WHEN '07' THEN 'Julho'
        WHEN '08' THEN 'Agosto'
        WHEN '09' THEN 'Setembro'
        WHEN '10' THEN 'Outubro'
        WHEN '11' THEN 'Novembro'
        WHEN '12' THEN 'Dezembro'
    END AS mes_emprestimo,
    ROUND(julianday('now') - julianday(e.data_emprestimo)) AS dias_decorridos,
    CASE
        WHEN e.status = 'devolvido' THEN 'Devolvido'
        WHEN e.status = 'ativo' AND julianday('now') > julianday(e.data_devolucao_prevista) THEN 'Atrasado'
        WHEN e.status = 'ativo' AND julianday(e.data_devolucao_prevista) - julianday('now') <= 3 THEN 'Vencendo'
        WHEN e.status = 'ativo' THEN 'No Prazo'
        ELSE 'Desconhecido'
    END AS status_prazo
    
FROM emprestimos e
JOIN usuarios u ON e.usuario_id = u.id
JOIN livros l ON e.livro_id = l.id
JOIN categorias c ON l.categoria_id = c.id
ORDER BY e.data_emprestimo DESC;

-- Estatísticas adicionais
SELECT 
    u.nome,
    COUNT(e.id) AS total_emprestimos
FROM usuarios u
LEFT JOIN emprestimos e ON u.id = e.usuario_id
GROUP BY u.id, u.nome
ORDER BY total_emprestimos DESC;

-- Média de dias por categoria
SELECT 
    c.nome AS categoria,
    ROUND(AVG(julianday(COALESCE(e.data_devolucao_real, 'now')) - julianday(e.data_emprestimo)), 2) AS media_dias_emprestimo
FROM categorias c
JOIN livros l ON c.id = l.categoria_id
JOIN emprestimos e ON l.id = e.livro_id
GROUP BY c.id, c.nome;
```

**Resposta Esperada para a Questão de Reflexão**:
- **Como funções ajudam em relatórios**:
  - Formatação consistente de dados
  - Cálculos complexos diretamente no banco
  - Agregações e análises temporais
  - Classificações e categorizações
- **Vantagens de fazer no banco**:
  - Performance: processamento próximo aos dados
  - Consistência: mesma lógica para todos os relatórios
  - Eficiência: menos dados transferidos para aplicação
  - Reutilização: queries podem ser usadas por diferentes aplicações
  - Manutenção: mudanças em um lugar afetam todos os relatórios

---

## Exercícios de Reflexão Adicional

### Reflexão 1: Performance vs Legibilidade

**Pergunta**: Quando você priorizaria performance sobre legibilidade ao usar funções SQL? Dê exemplos.

**Resposta Esperada**:
- **Priorizar performance quando**:
  - Tabelas muito grandes (milhões de registros)
  - Queries executadas frequentemente
  - Tempo de resposta crítico
  - Exemplo: Evitar funções em WHERE de colunas indexadas
- **Priorizar legibilidade quando**:
  - Queries complexas que precisam ser mantidas
  - Desenvolvimento inicial
  - Queries executadas raramente
  - Exemplo: Usar CASE ao invés de múltiplas queries separadas

---

### Reflexão 2: Manutenibilidade

**Pergunta**: Como você garantiria que queries com muitas funções sejam fáceis de manter e modificar?

**Resposta Esperada**:
- **Estratégias**:
  - Usar aliases descritivos
  - Adicionar comentários explicativos
  - Quebrar queries complexas em partes menores
  - Documentar lógica de negócio
  - Usar CTEs para queries muito complexas
  - Testar cada função separadamente
  - Versionar queries importantes
  - Revisar periodicamente queries antigas

---

## Conclusão dos Exercícios

Após completar estes exercícios, você deve ser capaz de:

- ✅ Manipular strings usando todas as funções aprendidas
- ✅ Trabalhar com datas e realizar cálculos temporais
- ✅ Realizar cálculos numéricos e arredondamentos
- ✅ Implementar lógica condicional em queries
- ✅ Combinar múltiplas funções em queries complexas
- ✅ Criar relatórios e análises usando funções SQL
- ✅ Entender quando usar funções no banco vs na aplicação
- ✅ Escrever queries legíveis e manuteníveis

**Próximo Passo**: Leia o arquivo de Performance e Boas Práticas para entender como otimizar o uso de funções SQL!

---

**Bons estudos! 🚀**

