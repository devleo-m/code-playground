# Aula 8: Advanced SQL Functions (Funções Avançadas de SQL)

## Introdução

Nesta aula, você aprenderá sobre **Advanced SQL Functions** (Funções Avançadas de SQL), que são ferramentas poderosas para manipulação e análise de dados dentro de consultas SQL. Essas funções permitem realizar transformações complexas de dados, cálculos avançados e lógica condicional diretamente nas queries, tornando-as mais expressivas e eficientes.

Advanced SQL functions são essenciais para:
- Manipular e transformar dados de texto (strings)
- Trabalhar com datas e horas de forma eficiente
- Realizar cálculos matemáticos precisos
- Implementar lógica condicional em queries
- Formatar e apresentar dados de forma adequada
- Limpar e padronizar dados durante consultas

Dominar funções avançadas de SQL é fundamental para qualquer desenvolvedor ou analista de dados, pois permite criar queries mais poderosas e resolver problemas complexos de manipulação de dados sem precisar processar os dados na aplicação.

---

## 1. O que são Advanced SQL Functions?

**Advanced SQL Functions** são funções pré-definidas que realizam operações específicas sobre dados. Elas podem ser aplicadas a valores de colunas, expressões ou constantes, e retornam um resultado transformado ou calculado.

### Características das Funções SQL

1. **Pré-definidas**: Já estão disponíveis no SGBD
2. **Reutilizáveis**: Podem ser usadas em qualquer query
3. **Eficientes**: Otimizadas pelo banco de dados
4. **Tipadas**: Retornam tipos específicos de dados
5. **Combináveis**: Podem ser aninhadas e combinadas

### Categorias de Funções Avançadas

SQL oferece funções em várias categorias:

- **String Functions**: Manipulação de texto
- **Date & Time Functions**: Trabalho com datas e horas
- **Numeric Functions**: Cálculos matemáticos
- **Conditional Functions**: Lógica condicional

### Sintaxe Básica

```sql
SELECT FUNCAO(coluna) FROM tabela;
SELECT FUNCAO(coluna, parametro) FROM tabela;
SELECT FUNCAO(coluna1, coluna2) FROM tabela;
```

### Por que Usar Funções SQL?

**Sem Funções** (processamento na aplicação):
```python
# Na aplicação Python
resultados = db.execute("SELECT titulo, autor FROM livros")
for row in resultados:
    titulo_upper = row['titulo'].upper()  # Processamento na aplicação
    print(titulo_upper)
```

**Com Funções SQL** (processamento no banco):
```sql
-- Processamento no banco de dados (mais eficiente)
SELECT UPPER(titulo) AS titulo_maiusculo, autor
FROM livros;
```

**Vantagens**:
- Mais eficiente (processamento no banco)
- Menos código na aplicação
- Consistência de transformações
- Aproveitamento de índices e otimizações do banco

---

## 2. String Functions (Funções de String)

String functions permitem manipular e transformar dados de texto. Elas são essenciais para limpeza de dados, formatação, busca e transformação de strings.

### 2.1 CONCAT (Concatenar Strings)

A função **CONCAT** combina duas ou mais strings em uma única string.

#### Características

- Combina múltiplas strings
- Aceita dois ou mais argumentos
- Retorna uma nova string combinada
- Se qualquer argumento for NULL, o resultado pode ser NULL (depende do SGBD)

#### Sintaxe

```sql
CONCAT(string1, string2, ...)
```

#### Exemplos Práticos

**Exemplo 1: Combinar Nome e Sobrenome**

```sql
-- Combinar nome e email do usuário
SELECT 
    nome,
    email,
    CONCAT(nome, ' - ', email) AS nome_email
FROM usuarios;
```

**Resultado**:
```
nome          | email              | nome_email
João Silva    | joao@email.com     | João Silva - joao@email.com
Maria Santos  | maria@email.com    | Maria Santos - maria@email.com
```

**Exemplo 2: Criar Descrição Completa do Livro**

```sql
-- Combinar título, autor e ano
SELECT 
    l.titulo,
    a.nome AS autor,
    l.ano_publicacao,
    CONCAT(l.titulo, ' por ', a.nome, ' (', l.ano_publicacao, ')') AS descricao_completa
FROM livros l
JOIN autores a ON l.autor_id = a.id;
```

**Exemplo 3: Combinar Múltiplas Colunas**

```sql
-- Criar endereço completo do usuário
SELECT 
    nome,
    CONCAT('Email: ', email, ' | Telefone: ', COALESCE(telefone, 'Não informado')) AS contato
FROM usuarios;
```

**Nota sobre SQLite**: SQLite não tem função CONCAT, mas usa o operador `||` para concatenar:

```sql
-- SQLite: usar operador ||
SELECT nome || ' - ' || email AS nome_email
FROM usuarios;
```

---

### 2.2 LENGTH (Comprimento da String)

A função **LENGTH** retorna o número de caracteres em uma string.

#### Características

- Retorna o comprimento em caracteres
- Espaços contam como caracteres
- NULL retorna NULL
- Útil para validação e formatação

#### Sintaxe

```sql
LENGTH(string)
```

#### Exemplos Práticos

**Exemplo 1: Verificar Tamanho de Títulos**

```sql
-- Listar títulos e seus comprimentos
SELECT 
    titulo,
    LENGTH(titulo) AS tamanho_titulo
FROM livros
ORDER BY tamanho_titulo DESC;
```

**Exemplo 2: Filtrar Títulos Longos**

```sql
-- Encontrar títulos com mais de 30 caracteres
SELECT 
    titulo,
    LENGTH(titulo) AS tamanho
FROM livros
WHERE LENGTH(titulo) > 30;
```

**Exemplo 3: Validar Tamanho de Email**

```sql
-- Verificar emails muito curtos ou muito longos
SELECT 
    nome,
    email,
    LENGTH(email) AS tamanho_email
FROM usuarios
WHERE LENGTH(email) < 10 OR LENGTH(email) > 50;
```

**Exemplo 4: Estatísticas de Comprimento**

```sql
-- Calcular média de comprimento de títulos por categoria
SELECT 
    c.nome AS categoria,
    AVG(LENGTH(l.titulo)) AS media_comprimento_titulo,
    MAX(LENGTH(l.titulo)) AS maior_titulo,
    MIN(LENGTH(l.titulo)) AS menor_titulo
FROM categorias c
JOIN livros l ON c.id = l.categoria_id
GROUP BY c.id, c.nome;
```

---

### 2.3 SUBSTRING (Extrair Substring)

A função **SUBSTRING** extrai uma porção de uma string, começando em uma posição específica e com um comprimento opcional.

#### Características

- Extrai parte de uma string
- Requer posição inicial (1-based)
- Comprimento opcional (se omitido, extrai até o fim)
- Útil para parsing e formatação

#### Sintaxe

```sql
SUBSTRING(string, inicio, comprimento)
-- ou
SUBSTR(string, inicio, comprimento)  -- SQLite usa SUBSTR
```

#### Exemplos Práticos

**Exemplo 1: Extrair Primeiros Caracteres**

```sql
-- Extrair primeiras 10 letras do título
SELECT 
    titulo,
    SUBSTR(titulo, 1, 10) AS primeiras_10_letras
FROM livros;
```

**Exemplo 2: Extrair Ano do ISBN**

```sql
-- Se ISBN tem formato específico, extrair parte
SELECT 
    titulo,
    isbn,
    SUBSTR(isbn, 1, 3) AS prefixo_isbn
FROM livros
WHERE isbn IS NOT NULL;
```

**Exemplo 3: Extrair Domínio do Email**

```sql
-- Extrair parte após @ do email
SELECT 
    nome,
    email,
    SUBSTR(email, INSTR(email, '@') + 1) AS dominio
FROM usuarios;
```

**Exemplo 4: Criar Abreviação**

```sql
-- Criar abreviação do título (primeiras 3 letras)
SELECT 
    titulo,
    SUBSTR(titulo, 1, 3) AS abreviacao
FROM livros;
```

---

### 2.4 REPLACE (Substituir Texto)

A função **REPLACE** substitui todas as ocorrências de uma substring por outra substring.

#### Características

- Substitui todas as ocorrências
- Case-sensitive (diferencia maiúsculas/minúsculas)
- Retorna nova string modificada
- Útil para limpeza e padronização de dados

#### Sintaxe

```sql
REPLACE(string, substring_antiga, substring_nova)
```

#### Exemplos Práticos

**Exemplo 1: Padronizar Formato**

```sql
-- Substituir espaços por underscores
SELECT 
    titulo,
    REPLACE(titulo, ' ', '_') AS titulo_sem_espacos
FROM livros;
```

**Exemplo 2: Corrigir Erros Comuns**

```sql
-- Corrigir erro comum de digitação
SELECT 
    titulo,
    REPLACE(titulo, 'Fundação', 'Fundacao') AS titulo_corrigido
FROM livros;
```

**Exemplo 3: Remover Caracteres Especiais**

```sql
-- Remover hífens de ISBN
SELECT 
    titulo,
    isbn,
    REPLACE(isbn, '-', '') AS isbn_sem_hifen
FROM livros
WHERE isbn IS NOT NULL;
```

**Exemplo 4: Normalizar Espaços**

```sql
-- Substituir múltiplos espaços por um único espaço
SELECT 
    titulo,
    REPLACE(REPLACE(titulo, '  ', ' '), '  ', ' ') AS titulo_normalizado
FROM livros;
```

---

### 2.5 UPPER (Converter para Maiúsculas)

A função **UPPER** converte todos os caracteres de uma string para maiúsculas.

#### Características

- Converte para maiúsculas
- Não afeta caracteres não-alfabéticos
- Útil para normalização e comparações case-insensitive
- Retorna nova string

#### Sintaxe

```sql
UPPER(string)
```

#### Exemplos Práticos

**Exemplo 1: Normalizar Títulos**

```sql
-- Converter todos os títulos para maiúsculas
SELECT 
    titulo,
    UPPER(titulo) AS titulo_maiusculo
FROM livros;
```

**Exemplo 2: Busca Case-Insensitive**

```sql
-- Buscar livros independente de maiúsculas/minúsculas
SELECT titulo
FROM livros
WHERE UPPER(titulo) LIKE UPPER('%fundação%');
```

**Exemplo 3: Comparar Emails Normalizados**

```sql
-- Comparar emails sem considerar maiúsculas/minúsculas
SELECT 
    nome,
    email,
    UPPER(email) AS email_normalizado
FROM usuarios
ORDER BY email_normalizado;
```

**Exemplo 4: Criar Códigos de Referência**

```sql
-- Criar código de referência em maiúsculas
SELECT 
    titulo,
    UPPER(SUBSTR(titulo, 1, 3)) || id AS codigo_referencia
FROM livros;
```

---

### 2.6 LOWER (Converter para Minúsculas)

A função **LOWER** converte todos os caracteres de uma string para minúsculas.

#### Características

- Converte para minúsculas
- Não afeta caracteres não-alfabéticos
- Útil para normalização e comparações case-insensitive
- Retorna nova string

#### Sintaxe

```sql
LOWER(string)
```

#### Exemplos Práticos

**Exemplo 1: Normalizar Dados**

```sql
-- Converter emails para minúsculas
SELECT 
    nome,
    email,
    LOWER(email) AS email_minusculo
FROM usuarios;
```

**Exemplo 2: Busca Case-Insensitive**

```sql
-- Buscar autores independente de maiúsculas/minúsculas
SELECT nome
FROM autores
WHERE LOWER(nome) LIKE LOWER('%machado%');
```

**Exemplo 3: Padronizar Nacionalidades**

```sql
-- Padronizar nacionalidades para minúsculas
SELECT 
    nome,
    nacionalidade,
    LOWER(nacionalidade) AS nacionalidade_normalizada
FROM autores;
```

**Exemplo 4: Comparação Case-Insensitive**

```sql
-- Comparar strings sem considerar maiúsculas/minúsculas
SELECT 
    l.titulo,
    c.nome AS categoria
FROM livros l
JOIN categorias c ON l.categoria_id = c.id
WHERE LOWER(c.nome) = 'ficção científica';
```

---

## 3. Date & Time Functions (Funções de Data e Hora)

Date & Time functions permitem trabalhar com dados temporais, realizar cálculos com datas e formatar valores de data e hora.

### 3.1 DATE (Tipo e Função de Data)

O tipo **DATE** armazena apenas a data (sem hora), e a função **DATE** pode extrair ou converter valores para data.

#### Características

- Armazena data no formato YYYY-MM-DD
- Não inclui informação de hora
- Permite cálculos e comparações de datas
- Útil para datas de nascimento, eventos, etc.

#### Sintaxe

```sql
DATE(string)
DATE('YYYY-MM-DD')
```

#### Exemplos Práticos

**Exemplo 1: Extrair Data de String**

```sql
-- Converter string para data
SELECT 
    nome,
    data_nascimento,
    DATE(data_nascimento) AS data_formatada
FROM autores
WHERE data_nascimento IS NOT NULL;
```

**Exemplo 2: Filtrar por Data**

```sql
-- Encontrar autores nascidos após 1950
SELECT 
    nome,
    data_nascimento
FROM autores
WHERE DATE(data_nascimento) > DATE('1950-01-01');
```

**Exemplo 3: Calcular Idade Aproximada**

```sql
-- Calcular idade baseada na data de nascimento
SELECT 
    nome,
    data_nascimento,
    (julianday('now') - julianday(data_nascimento)) / 365.25 AS idade_aproximada
FROM autores
WHERE data_nascimento IS NOT NULL;
```

---

### 3.2 TIME (Tipo e Função de Hora)

O tipo **TIME** armazena apenas a hora (sem data), no formato HH:MM:SS.

#### Características

- Armazena hora no formato HH:MM:SS
- Não inclui informação de data
- Permite cálculos e comparações de horas
- Útil para horários de eventos, duração, etc.

#### Sintaxe

```sql
TIME(string)
TIME('HH:MM:SS')
```

#### Exemplos Práticos

**Exemplo 1: Extrair Hora de Timestamp**

```sql
-- Se tivéssemos timestamp, extrair apenas hora
-- Exemplo conceitual
SELECT TIME('2024-01-15 14:30:00') AS hora;
-- Retorna: 14:30:00
```

**Exemplo 2: Comparar Horas**

```sql
-- Exemplo conceitual de comparação de horas
-- SELECT * FROM eventos WHERE TIME(hora_evento) > TIME('12:00:00');
```

---

### 3.3 TIMESTAMP (Tipo e Função de Data/Hora)

O tipo **TIMESTAMP** armazena tanto data quanto hora, no formato YYYY-MM-DD HH:MM:SS.

#### Características

- Armazena data e hora completas
- Formato: YYYY-MM-DD HH:MM:SS
- Permite cálculos precisos de tempo
- Útil para logs, auditoria, timestamps de eventos

#### Sintaxe

```sql
TIMESTAMP(string)
DATETIME(string)  -- SQLite usa DATETIME
```

#### Exemplos Práticos

**Exemplo 1: Criar Timestamp Atual**

```sql
-- Obter data e hora atual
SELECT DATETIME('now') AS timestamp_atual;
```

**Exemplo 2: Adicionar Timestamp a Registro**

```sql
-- Se tivéssemos coluna de timestamp
-- UPDATE emprestimos 
-- SET data_hora_emprestimo = DATETIME('now')
-- WHERE id = 1;
```

**Exemplo 3: Calcular Diferença de Tempo**

```sql
-- Calcular dias entre duas datas
SELECT 
    data_emprestimo,
    data_devolucao_prevista,
    julianday(data_devolucao_prevista) - julianday(data_emprestimo) AS dias_emprestimo
FROM emprestimos
WHERE data_devolucao_prevista IS NOT NULL;
```

---

### 3.4 DATEPART / strftime (Extrair Parte da Data)

A função **DATEPART** (ou **strftime** no SQLite) extrai uma parte específica de uma data ou hora.

#### Características

- Extrai ano, mês, dia, hora, minuto, segundo
- Retorna valor numérico
- Útil para agrupamentos e filtros por período
- SQLite usa strftime com formato diferente

#### Sintaxe SQLite

```sql
strftime('%Y', data)  -- Ano
strftime('%m', data)  -- Mês (01-12)
strftime('%d', data)  -- Dia (01-31)
strftime('%w', data)  -- Dia da semana (0=Domingo)
```

#### Exemplos Práticos

**Exemplo 1: Extrair Ano**

```sql
-- Extrair ano de publicação
SELECT 
    titulo,
    ano_publicacao,
    strftime('%Y', data_nascimento) AS ano_nascimento
FROM autores
WHERE data_nascimento IS NOT NULL;
```

**Exemplo 2: Agrupar por Ano**

```sql
-- Contar empréstimos por ano
SELECT 
    strftime('%Y', data_emprestimo) AS ano,
    COUNT(*) AS total_emprestimos
FROM emprestimos
GROUP BY ano
ORDER BY ano DESC;
```

**Exemplo 3: Filtrar por Mês**

```sql
-- Encontrar empréstimos de janeiro
SELECT 
    id,
    data_emprestimo,
    strftime('%m', data_emprestimo) AS mes
FROM emprestimos
WHERE strftime('%m', data_emprestimo) = '01';
```

**Exemplo 4: Extrair Dia da Semana**

```sql
-- Verificar em que dia da semana foram feitos os empréstimos
SELECT 
    data_emprestimo,
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
```

---

### 3.5 DATEADD / date (Adicionar Intervalo à Data)

A função **DATEADD** (ou funções de data do SQLite) adiciona ou subtrai um intervalo de tempo a uma data.

#### Características

- Adiciona/subtrai dias, meses, anos
- Retorna nova data calculada
- Útil para calcular datas futuras ou passadas
- SQLite usa julianday e date para cálculos

#### Sintaxe SQLite

```sql
date(data, '+N days')   -- Adicionar dias
date(data, '+N months') -- Adicionar meses
date(data, '+N years')  -- Adicionar anos
```

#### Exemplos Práticos

**Exemplo 1: Calcular Data de Devolução**

```sql
-- Calcular data de devolução (15 dias após empréstimo)
SELECT 
    id,
    data_emprestimo,
    date(data_emprestimo, '+15 days') AS data_devolucao_calculada
FROM emprestimos;
```

**Exemplo 2: Encontrar Empréstimos Próximos do Vencimento**

```sql
-- Empréstimos que vencem nos próximos 7 dias
SELECT 
    id,
    data_devolucao_prevista,
    julianday(data_devolucao_prevista) - julianday('now') AS dias_restantes
FROM emprestimos
WHERE status = 'ativo'
  AND data_devolucao_prevista IS NOT NULL
  AND julianday(data_devolucao_prevista) - julianday('now') <= 7;
```

**Exemplo 3: Calcular Data de Aniversário**

```sql
-- Calcular próximo aniversário
SELECT 
    nome,
    data_nascimento,
    date(data_nascimento, '+' || (strftime('%Y', 'now') - strftime('%Y', data_nascimento) + 1) || ' years') AS proximo_aniversario
FROM autores
WHERE data_nascimento IS NOT NULL;
```

**Exemplo 4: Adicionar Meses**

```sql
-- Calcular data após 3 meses
SELECT 
    data_emprestimo,
    date(data_emprestimo, '+3 months') AS data_apos_3_meses
FROM emprestimos;
```

---

## 4. Numeric Functions (Funções Numéricas)

Numeric functions realizam cálculos matemáticos e transformações em valores numéricos.

### 4.1 FLOOR (Arredondar para Baixo)

A função **FLOOR** arredonda um número para baixo até o inteiro mais próximo.

#### Características

- Arredonda para baixo
- Sempre retorna inteiro menor ou igual
- Útil para divisões e cálculos de quantidade

#### Sintaxe

```sql
FLOOR(número)
```

#### Exemplos Práticos

**Exemplo 1: Calcular Páginas Necessárias**

```sql
-- Se tivéssemos número de páginas, calcular páginas por livro
-- Exemplo conceitual
SELECT FLOOR(1000 / 25) AS livros_por_pagina;
-- Retorna: 40
```

**Exemplo 2: Arredondar Quantidades**

```sql
-- Arredondar quantidade disponível para baixo
SELECT 
    titulo,
    quantidade_disponivel,
    FLOOR(quantidade_disponivel / 2.0) AS metade_arredondada
FROM livros;
```

---

### 4.2 ABS (Valor Absoluto)

A função **ABS** retorna o valor absoluto de um número (remove o sinal negativo).

#### Características

- Remove sinal negativo
- Retorna sempre positivo ou zero
- Útil para cálculos de diferença e distância

#### Sintaxe

```sql
ABS(número)
```

#### Exemplos Práticos

**Exemplo 1: Calcular Diferença Absoluta**

```sql
-- Calcular diferença absoluta entre datas
SELECT 
    data_emprestimo,
    data_devolucao_prevista,
    ABS(julianday(data_devolucao_prevista) - julianday(data_emprestimo)) AS dias_absolutos
FROM emprestimos
WHERE data_devolucao_prevista IS NOT NULL;
```

**Exemplo 2: Normalizar Valores**

```sql
-- Garantir que diferença seja sempre positiva
SELECT 
    id,
    quantidade_disponivel,
    ABS(quantidade_disponivel) AS quantidade_positiva
FROM livros;
```

---

### 4.3 MOD (Módulo/Resto da Divisão)

A função **MOD** retorna o resto da divisão de um número por outro.

#### Características

- Retorna resto da divisão
- Útil para verificar paridade, ciclos, distribuição
- SQLite usa operador %

#### Sintaxe

```sql
MOD(dividendo, divisor)
-- ou em SQLite
dividendo % divisor
```

#### Exemplos Práticos

**Exemplo 1: Verificar Paridade**

```sql
-- Verificar se quantidade é par ou ímpar
SELECT 
    titulo,
    quantidade_disponivel,
    CASE 
        WHEN quantidade_disponivel % 2 = 0 THEN 'Par'
        ELSE 'Ímpar'
    END AS paridade
FROM livros;
```

**Exemplo 2: Agrupar em Categorias**

```sql
-- Agrupar livros em categorias baseado no ID
SELECT 
    titulo,
    id,
    (id % 3) AS grupo
FROM livros;
```

**Exemplo 3: Verificar Múltiplos**

```sql
-- Verificar se quantidade é múltiplo de 5
SELECT 
    titulo,
    quantidade_disponivel,
    CASE 
        WHEN quantidade_disponivel % 5 = 0 THEN 'Múltiplo de 5'
        ELSE 'Não é múltiplo de 5'
    END AS status
FROM livros;
```

---

### 4.4 ROUND (Arredondar)

A função **ROUND** arredonda um número para um número específico de casas decimais.

#### Características

- Arredonda para casas decimais especificadas
- Se não especificar casas, arredonda para inteiro
- Útil para formatação e apresentação

#### Sintaxe

```sql
ROUND(número, casas_decimais)
```

#### Exemplos Práticos

**Exemplo 1: Arredondar Médias**

```sql
-- Calcular média de estoque arredondada
SELECT 
    c.nome AS categoria,
    ROUND(AVG(l.quantidade_disponivel), 2) AS media_estoque
FROM categorias c
JOIN livros l ON c.id = l.categoria_id
GROUP BY c.id, c.nome;
```

**Exemplo 2: Arredondar para Inteiro**

```sql
-- Arredondar quantidade para inteiro
SELECT 
    titulo,
    quantidade_disponivel,
    ROUND(quantidade_disponivel) AS quantidade_arredondada
FROM livros;
```

**Exemplo 3: Arredondar Cálculos**

```sql
-- Calcular porcentagem arredondada
SELECT 
    titulo,
    quantidade_disponivel,
    ROUND((quantidade_disponivel * 100.0 / (SELECT SUM(quantidade_disponivel) FROM livros)), 2) AS porcentagem_estoque
FROM livros;
```

---

### 4.5 CEILING (Arredondar para Cima)

A função **CEILING** (ou **CEIL**) arredonda um número para cima até o inteiro mais próximo.

#### Características

- Arredonda para cima
- Sempre retorna inteiro maior ou igual
- Útil para cálculos de capacidade e recursos

#### Sintaxe

```sql
CEILING(número)
-- ou
CEIL(número)
```

#### Exemplos Práticos

**Exemplo 1: Calcular Capacidade Necessária**

```sql
-- Se tivéssemos capacidade por prateleira
-- Calcular prateleiras necessárias
-- Exemplo: 23 livros, 10 por prateleira = 3 prateleiras
SELECT CEIL(23 / 10.0) AS prateleiras_necessarias;
-- Retorna: 3
```

**Exemplo 2: Arredondar para Cima**

```sql
-- Arredondar quantidade para cima
SELECT 
    titulo,
    quantidade_disponivel,
    CEIL(quantidade_disponivel / 2.0) AS metade_arredondada_cima
FROM livros;
```

---

## 5. Conditional Functions (Funções Condicionais)

Conditional functions permitem implementar lógica condicional diretamente em queries SQL, tornando-as mais expressivas e flexíveis.

### 5.1 CASE (Estrutura Condicional)

A estrutura **CASE** permite criar lógica condicional tipo if-else dentro de queries SQL.

#### Características

- Implementa lógica condicional
- Pode ter múltiplas condições (WHEN)
- Pode ter valor padrão (ELSE)
- Retorna valor baseado em condições
- Muito flexível e poderosa

#### Sintaxe

```sql
CASE
    WHEN condição1 THEN valor1
    WHEN condição2 THEN valor2
    ELSE valor_padrao
END
```

#### Exemplos Práticos

**Exemplo 1: Classificar Estoque**

```sql
-- Classificar livros por quantidade de estoque
SELECT 
    titulo,
    quantidade_disponivel,
    CASE
        WHEN quantidade_disponivel = 0 THEN 'Esgotado'
        WHEN quantidade_disponivel < 5 THEN 'Estoque Baixo'
        WHEN quantidade_disponivel < 10 THEN 'Estoque Médio'
        ELSE 'Estoque Alto'
    END AS status_estoque
FROM livros;
```

**Exemplo 2: Classificar Empréstimos**

```sql
-- Classificar empréstimos por status
SELECT 
    id,
    data_emprestimo,
    data_devolucao_prevista,
    CASE
        WHEN status = 'ativo' AND julianday('now') > julianday(data_devolucao_prevista) THEN 'Atrasado'
        WHEN status = 'ativo' THEN 'Em Andamento'
        WHEN status = 'devolvido' THEN 'Devolvido'
        ELSE 'Desconhecido'
    END AS status_detalhado
FROM emprestimos;
```

**Exemplo 3: CASE em SELECT**

```sql
-- Criar descrição baseada em múltiplas condições
SELECT 
    l.titulo,
    a.nome AS autor,
    c.nome AS categoria,
    CASE
        WHEN l.quantidade_disponivel > 10 AND l.ano_publicacao > 2000 THEN 'Novo e Disponível'
        WHEN l.quantidade_disponivel > 10 THEN 'Disponível'
        WHEN l.ano_publicacao > 2000 THEN 'Novo'
        ELSE 'Verificar'
    END AS classificacao
FROM livros l
JOIN autores a ON l.autor_id = a.id
JOIN categorias c ON l.categoria_id = c.id;
```

**Exemplo 4: CASE com Agregação**

```sql
-- Contar por categoria de estoque
SELECT 
    CASE
        WHEN quantidade_disponivel = 0 THEN 'Esgotado'
        WHEN quantidade_disponivel < 5 THEN 'Baixo'
        ELSE 'Normal'
    END AS nivel_estoque,
    COUNT(*) AS total_livros
FROM livros
GROUP BY nivel_estoque;
```

---

### 5.2 NULLIF (Retornar NULL se Igual)

A função **NULLIF** compara dois valores e retorna NULL se forem iguais, caso contrário retorna o primeiro valor.

#### Características

- Retorna NULL se valores forem iguais
- Retorna primeiro valor se diferentes
- Útil para evitar divisão por zero
- Útil para tratar valores padrão como NULL

#### Sintaxe

```sql
NULLIF(valor1, valor2)
```

#### Exemplos Práticos

**Exemplo 1: Evitar Divisão por Zero**

```sql
-- Calcular média evitando divisão por zero
SELECT 
    titulo,
    quantidade_disponivel,
    CASE 
        WHEN NULLIF(quantidade_disponivel, 0) IS NULL THEN 0
        ELSE 100 / quantidade_disponivel
    END AS calculo_seguro
FROM livros;
```

**Exemplo 2: Tratar Valores Padrão**

```sql
-- Tratar string vazia como NULL
SELECT 
    nome,
    NULLIF(telefone, '') AS telefone_ou_null
FROM usuarios;
```

**Exemplo 3: Normalizar Valores**

```sql
-- Tratar valores específicos como NULL
SELECT 
    titulo,
    NULLIF(editora, 'Desconhecida') AS editora_normalizada
FROM livros;
```

---

### 5.3 COALESCE (Primeiro Valor Não-NULL)

A função **COALESCE** retorna o primeiro valor não-NULL de uma lista de valores.

#### Características

- Retorna primeiro valor não-NULL
- Aceita múltiplos argumentos
- Útil para fornecer valores padrão
- Útil para combinar valores de múltiplas colunas

#### Sintaxe

```sql
COALESCE(valor1, valor2, valor3, ...)
```

#### Exemplos Práticos

**Exemplo 1: Fornecer Valor Padrão**

```sql
-- Usar 'Não informado' se telefone for NULL
SELECT 
    nome,
    email,
    COALESCE(telefone, 'Não informado') AS telefone_display
FROM usuarios;
```

**Exemplo 2: Combinar Múltiplas Colunas**

```sql
-- Usar primeira coluna não-NULL disponível
SELECT 
    nome,
    COALESCE(telefone, email, 'Sem contato') AS contato_principal
FROM usuarios;
```

**Exemplo 3: Calcular com Valores Padrão**

```sql
-- Calcular com valor padrão se NULL
SELECT 
    titulo,
    quantidade_disponivel,
    COALESCE(quantidade_disponivel, 0) AS estoque_seguro
FROM livros;
```

**Exemplo 4: Priorizar Valores**

```sql
-- Priorizar data_devolucao_real sobre data_devolucao_prevista
SELECT 
    id,
    data_emprestimo,
    COALESCE(data_devolucao_real, data_devolucao_prevista, 'Não definida') AS data_final
FROM emprestimos;
```

---

## 6. Combinando Funções

Funções SQL podem ser combinadas e aninhadas para criar transformações complexas.

### Exemplos de Combinação

**Exemplo 1: String + Conditional**

```sql
-- Criar descrição formatada com lógica condicional
SELECT 
    titulo,
    CONCAT(
        UPPER(SUBSTR(titulo, 1, 1)),
        LOWER(SUBSTR(titulo, 2)),
        CASE 
            WHEN quantidade_disponivel > 0 THEN ' - Disponível'
            ELSE ' - Esgotado'
        END
    ) AS titulo_formatado
FROM livros;
```

**Exemplo 2: Date + Numeric + Conditional**

```sql
-- Calcular status de empréstimo com múltiplas funções
SELECT 
    id,
    data_emprestimo,
    data_devolucao_prevista,
    ROUND(julianday('now') - julianday(data_emprestimo)) AS dias_decorridos,
    CASE
        WHEN julianday('now') > julianday(data_devolucao_prevista) THEN 'Atrasado'
        WHEN ROUND(julianday(data_devolucao_prevista) - julianday('now')) <= 3 THEN 'Vencendo'
        ELSE 'No Prazo'
    END AS status
FROM emprestimos
WHERE status = 'ativo';
```

**Exemplo 3: Múltiplas Funções Aninhadas**

```sql
-- Análise complexa combinando várias funções
SELECT 
    UPPER(SUBSTR(c.nome, 1, 3)) AS codigo_categoria,
    COUNT(l.id) AS total_livros,
    ROUND(AVG(l.quantidade_disponivel), 2) AS media_estoque,
    COALESCE(MAX(l.ano_publicacao), 0) AS ultimo_ano,
    CASE
        WHEN COUNT(l.id) > 5 THEN 'Grande'
        WHEN COUNT(l.id) > 2 THEN 'Média'
        ELSE 'Pequena'
    END AS tamanho_categoria
FROM categorias c
LEFT JOIN livros l ON c.id = l.categoria_id
GROUP BY c.id, c.nome;
```

---

## 7. Considerações de Performance

### Impacto das Funções na Performance

1. **Funções em WHERE**: Podem impedir uso de índices
   ```sql
   -- ❌ Lento: função impede uso de índice
   SELECT * FROM livros WHERE UPPER(titulo) = 'FUNDAÇÃO';
   
   -- ✅ Melhor: índice pode ser usado
   SELECT * FROM livros WHERE titulo = 'Fundação';
   ```

2. **Funções Aninhadas**: Podem ser custosas
   ```sql
   -- Múltiplas funções podem ser lentas
   SELECT REPLACE(UPPER(SUBSTR(titulo, 1, 10)), ' ', '_') FROM livros;
   ```

3. **Funções em Agregações**: Podem impactar GROUP BY
   ```sql
   -- Funções em GROUP BY podem ser lentas
   SELECT strftime('%Y', data_emprestimo), COUNT(*) 
   FROM emprestimos 
   GROUP BY strftime('%Y', data_emprestimo);
   ```

### Boas Práticas

1. **Evite Funções em WHERE quando possível**: Use valores diretos
2. **Indexe Colunas Transformadas**: Se precisar de função em WHERE frequentemente
3. **Cache Resultados**: Para cálculos complexos repetidos
4. **Use Funções Apropriadas**: Escolha a função mais eficiente
5. **Teste Performance**: Meça impacto de funções em queries grandes

---

## 8. Resumo

### Funções por Categoria

| Categoria | Funções Principais | Uso Comum |
|-----------|-------------------|-----------|
| **String** | CONCAT, LENGTH, SUBSTRING, REPLACE, UPPER, LOWER | Limpeza, formatação, busca |
| **Date/Time** | DATE, TIME, TIMESTAMP, DATEPART, DATEADD | Cálculos temporais, agrupamentos |
| **Numeric** | FLOOR, ABS, MOD, ROUND, CEILING | Cálculos matemáticos, arredondamentos |
| **Conditional** | CASE, NULLIF, COALESCE | Lógica condicional, valores padrão |

### Quando Usar Cada Tipo

- **String Functions**: Quando precisa manipular texto
- **Date/Time Functions**: Quando trabalha com datas e horas
- **Numeric Functions**: Quando precisa calcular ou arredondar números
- **Conditional Functions**: Quando precisa de lógica if-else em queries

### Comandos Importantes

```sql
-- String
CONCAT(str1, str2)  -- ou || no SQLite
LENGTH(string)
SUBSTR(string, start, length)
REPLACE(string, old, new)
UPPER(string)
LOWER(string)

-- Date/Time (SQLite)
DATE('now')
DATETIME('now')
strftime('%Y', date)
date(date, '+N days')

-- Numeric
FLOOR(number)
ABS(number)
number % divisor  -- MOD
ROUND(number, decimals)
CEIL(number)

-- Conditional
CASE WHEN ... THEN ... ELSE ... END
NULLIF(val1, val2)
COALESCE(val1, val2, ...)
```

---

## 9. Próximos Passos

Agora que você entende funções avançadas de SQL:

1. **Execute todos os exemplos** no banco de dados `biblioteca.db`
2. **Experimente combinações** de diferentes funções
3. **Crie suas próprias queries** usando funções avançadas
4. **Leia a aula simplificada** para reforçar o entendimento
5. **Complete os exercícios** para praticar

---

**Bons estudos! 🚀**

**Lembre-se**: Funções SQL são poderosas e permitem realizar transformações complexas diretamente no banco de dados. Pratique muito e você dominará essas ferramentas essenciais!

