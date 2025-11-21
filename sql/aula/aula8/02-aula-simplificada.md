# Aula 8 - Simplificada: Entendendo Advanced SQL Functions

## Introdução

Imagine que você está trabalhando em uma biblioteca e precisa:
- Formatar nomes de livros de forma consistente
- Calcular quantos dias um livro está emprestado
- Arredondar valores de estoque
- Criar mensagens personalizadas baseadas em condições

Essas tarefas são exatamente o que **Advanced SQL Functions** (Funções Avançadas de SQL) fazem: elas são como **ferramentas especiais** que transformam e manipulam dados diretamente no banco de dados, sem precisar processar na aplicação.

**Funções SQL são como uma caixa de ferramentas mágica**: você coloca dados de uma forma e elas transformam em outra forma útil!

---

## 1. Advanced SQL Functions: A Analogia da Fábrica de Transformação

### Pensando em Funções como Máquinas de Transformação

Imagine uma fábrica onde você tem diferentes máquinas especializadas:

**Máquina de Texto**: Transforma textos (maiúsculas, minúsculas, corta, cola)
**Máquina de Data**: Calcula datas, extrai anos, adiciona dias
**Máquina de Números**: Arredonda, calcula, transforma números
**Máquina de Decisão**: Toma decisões baseadas em condições

Cada máquina (função) faz uma transformação específica nos seus dados!

### Por que Precisamos de Funções?

**Sem Funções** (processamento manual):
```
Bibliotecário: "Preciso listar todos os títulos em maiúsculas..."
Ação: Abre cada livro, reescreve o título em maiúsculas manualmente ❌
Tempo: Muito lento e propenso a erros
```

**Com Funções** (processamento automático):
```
Bibliotecário: "Preciso listar todos os títulos em maiúsculas..."
Ação: SELECT UPPER(titulo) FROM livros; ✅
Tempo: Instantâneo e sempre correto
```

Funções são como **assistentes automáticos** que fazem o trabalho pesado por você!

---

## 2. String Functions: A Máquina de Transformação de Texto

### 2.1 CONCAT: A Cola de Textos

**Analogia**: Pense em CONCAT como uma **cola mágica** que junta pedaços de texto.

**Exemplo do Dia a Dia**:
```
Você tem: "João" e "Silva"
CONCAT junta: "João Silva"

Você tem: "Dom" e "Casmurro"
CONCAT junta: "Dom Casmurro"
```

**No Banco de Dados**:
```sql
-- Juntar nome e email
SELECT CONCAT(nome, ' - ', email) AS nome_email
FROM usuarios;

-- Resultado:
-- "João Silva - joao@email.com"
-- "Maria Santos - maria@email.com"
```

**Analogia Visual**: É como colar etiquetas:
```
[Etiqueta 1: "João"] + [Etiqueta 2: "Silva"] = [Etiqueta Final: "João Silva"]
```

**No SQLite** (usa `||` ao invés de CONCAT):
```sql
SELECT nome || ' - ' || email AS nome_email
FROM usuarios;
```

---

### 2.2 LENGTH: A Régua Digital

**Analogia**: Pense em LENGTH como uma **régua** que mede o tamanho de um texto.

**Exemplo do Dia a Dia**:
```
Texto: "Fundação"
LENGTH mede: 9 caracteres

Texto: "Dom Casmurro"
LENGTH mede: 12 caracteres (espaço conta!)
```

**No Banco de Dados**:
```sql
-- Medir tamanho dos títulos
SELECT 
    titulo,
    LENGTH(titulo) AS tamanho
FROM livros;

-- Resultado:
-- "Fundação" | 9
-- "Dom Casmurro" | 12
```

**Analogia Visual**: É como medir uma corda:
```
Corda: "Fundação"
       |--------| = 9 unidades
```

**Uso Prático**: Validar tamanhos
```sql
-- Encontrar títulos muito longos
SELECT titulo
FROM livros
WHERE LENGTH(titulo) > 50;
```

---

### 2.3 SUBSTRING: A Tesoura de Texto

**Analogia**: Pense em SUBSTRING como uma **tesoura** que corta pedaços específicos de um texto.

**Exemplo do Dia a Dia**:
```
Texto completo: "Fundação"
Cortar primeiras 3 letras: "Fun"
Cortar do meio: "dação"
```

**No Banco de Dados**:
```sql
-- Cortar primeiras 10 letras do título
SELECT 
    titulo,
    SUBSTR(titulo, 1, 10) AS primeiras_10
FROM livros;

-- Resultado:
-- "Fundação" | "Fundação"
-- "Dom Casmurro" | "Dom Casmur"
```

**Analogia Visual**: É como cortar um bolo:
```
Bolo: "Fundação"
      |---| (primeiras 3 fatias = "Fun")
```

**Uso Prático**: Criar abreviações
```sql
-- Criar código de 3 letras
SELECT SUBSTR(titulo, 1, 3) AS codigo
FROM livros;
```

---

### 2.4 REPLACE: A Máquina de Substituição

**Analogia**: Pense em REPLACE como uma **máquina find-and-replace** que substitui todas as ocorrências.

**Exemplo do Dia a Dia**:
```
Texto: "Fundação"
Substituir "ã" por "a": "Fundacao"
Substituir espaços por "_": "Fundação" → "Fundação" (sem espaços para substituir)
```

**No Banco de Dados**:
```sql
-- Remover hífens de ISBN
SELECT 
    isbn,
    REPLACE(isbn, '-', '') AS isbn_sem_hifen
FROM livros;

-- Resultado:
-- "978-85-359-1484-1" | "9788535914841"
```

**Analogia Visual**: É como trocar todas as peças de uma cor por outra:
```
Antes: [Vermelho] [Azul] [Vermelho] [Verde]
Trocar Vermelho por Amarelo:
Depois: [Amarelo] [Azul] [Amarelo] [Verde]
```

**Uso Prático**: Limpeza de dados
```sql
-- Padronizar espaços
SELECT REPLACE(titulo, '  ', ' ') AS titulo_limpo
FROM livros;
```

---

### 2.5 UPPER: A Máquina de Maiúsculas

**Analogia**: Pense em UPPER como uma **máquina** que transforma tudo em maiúsculas.

**Exemplo do Dia a Dia**:
```
Entrada: "fundação"
UPPER transforma: "FUNDAÇÃO"

Entrada: "Dom Casmurro"
UPPER transforma: "DOM CASMURRO"
```

**No Banco de Dados**:
```sql
-- Converter títulos para maiúsculas
SELECT 
    titulo,
    UPPER(titulo) AS titulo_maiusculo
FROM livros;

-- Resultado:
-- "Fundação" | "FUNDAÇÃO"
-- "Dom Casmurro" | "DOM CASMURRO"
```

**Analogia Visual**: É como passar tudo por uma máquina que aumenta as letras:
```
Entrada: "fundação"
         ↓ [Máquina UPPER]
Saída:   "FUNDAÇÃO"
```

**Uso Prático**: Normalização e busca
```sql
-- Buscar sem considerar maiúsculas/minúsculas
SELECT * FROM livros
WHERE UPPER(titulo) LIKE UPPER('%fundação%');
```

---

### 2.6 LOWER: A Máquina de Minúsculas

**Analogia**: Pense em LOWER como uma **máquina** que transforma tudo em minúsculas.

**Exemplo do Dia a Dia**:
```
Entrada: "FUNDAÇÃO"
LOWER transforma: "fundação"

Entrada: "Dom Casmurro"
LOWER transforma: "dom casmurro"
```

**No Banco de Dados**:
```sql
-- Converter emails para minúsculas
SELECT 
    email,
    LOWER(email) AS email_normalizado
FROM usuarios;

-- Resultado:
-- "Joao@Email.com" | "joao@email.com"
```

**Analogia Visual**: É como passar tudo por uma máquina que diminui as letras:
```
Entrada: "FUNDAÇÃO"
         ↓ [Máquina LOWER]
Saída:   "fundação"
```

**Uso Prático**: Padronização
```sql
-- Padronizar nacionalidades
SELECT LOWER(nacionalidade) AS nacionalidade_padrao
FROM autores;
```

---

## 3. Date & Time Functions: A Máquina do Tempo

### 3.1 DATE: O Calendário Digital

**Analogia**: Pense em DATE como um **calendário** que trabalha apenas com datas (sem horas).

**Exemplo do Dia a Dia**:
```
Data completa: "2024-01-15 14:30:00"
DATE extrai: "2024-01-15" (só a data, sem hora)
```

**No Banco de Dados**:
```sql
-- Trabalhar apenas com datas
SELECT 
    nome,
    DATE(data_nascimento) AS data_nasc
FROM autores;
```

**Analogia Visual**: É como rasgar apenas a parte do calendário:
```
Calendário completo: "15 de Janeiro de 2024, 14:30"
                     ↓ [Extrair apenas data]
Data:                "15 de Janeiro de 2024"
```

---

### 3.2 strftime: O Extrator de Partes da Data

**Analogia**: Pense em strftime como uma **lupa especial** que vê apenas partes específicas de uma data.

**Exemplo do Dia a Dia**:
```
Data completa: "2024-01-15"
strftime('%Y') extrai: "2024" (ano)
strftime('%m') extrai: "01" (mês)
strftime('%d') extrai: "15" (dia)
```

**No Banco de Dados**:
```sql
-- Extrair ano de nascimento
SELECT 
    nome,
    strftime('%Y', data_nascimento) AS ano_nascimento
FROM autores;

-- Resultado:
-- "Machado de Assis" | "1839"
```

**Analogia Visual**: É como olhar apenas uma parte do calendário:
```
Calendário: [2024] [Janeiro] [15]
            ↓ [Lupa do Ano]
Ano:        [2024]
```

**Uso Prático**: Agrupar por período
```sql
-- Contar empréstimos por ano
SELECT 
    strftime('%Y', data_emprestimo) AS ano,
    COUNT(*) AS total
FROM emprestimos
GROUP BY ano;
```

---

### 3.3 date com Intervalos: A Calculadora de Datas

**Analogia**: Pense em `date(..., '+N days')` como uma **calculadora de datas** que adiciona ou subtrai tempo.

**Exemplo do Dia a Dia**:
```
Data atual: "15 de Janeiro"
Adicionar 15 dias: "30 de Janeiro"
Adicionar 1 mês: "15 de Fevereiro"
```

**No Banco de Dados**:
```sql
-- Calcular data de devolução (15 dias após empréstimo)
SELECT 
    data_emprestimo,
    date(data_emprestimo, '+15 days') AS data_devolucao
FROM emprestimos;

-- Resultado:
-- "2024-01-15" | "2024-01-30"
```

**Analogia Visual**: É como avançar no calendário:
```
Hoje: [15 Jan]
      ↓ [Adicionar 15 dias]
Futuro: [30 Jan]
```

**Uso Prático**: Calcular prazos
```sql
-- Empréstimos que vencem em 7 dias
SELECT *
FROM emprestimos
WHERE julianday(data_devolucao_prevista) - julianday('now') <= 7;
```

---

## 4. Numeric Functions: A Calculadora Avançada

### 4.1 FLOOR: O Arredondador para Baixo

**Analogia**: Pense em FLOOR como **cortar a parte decimal** e manter apenas o número inteiro menor.

**Exemplo do Dia a Dia**:
```
Número: 4.7
FLOOR: 4 (corta tudo depois da vírgula, sempre para baixo)

Número: 4.2
FLOOR: 4 (mesmo sendo próximo de 4, vai para 4)
```

**No Banco de Dados**:
```sql
-- Arredondar quantidade para baixo
SELECT 
    titulo,
    quantidade_disponivel,
    FLOOR(quantidade_disponivel / 2.0) AS metade_arredondada
FROM livros;

-- Resultado:
-- quantidade = 7 → metade = 3.5 → FLOOR = 3
```

**Analogia Visual**: É como cortar um bolo e pegar apenas a parte inteira:
```
Bolo: [████████] (8 pedaços)
Metade: [████] (4 pedaços inteiros)
FLOOR de 7.5: [████] (4 pedaços, descarta os 3.5 restantes)
```

---

### 4.2 ABS: O Removedor de Sinal Negativo

**Analogia**: Pense em ABS como uma **máquina** que sempre transforma números negativos em positivos.

**Exemplo do Dia a Dia**:
```
Número: -5
ABS: 5 (remove o sinal negativo)

Número: 5
ABS: 5 (já é positivo, não muda)
```

**No Banco de Dados**:
```sql
-- Calcular diferença absoluta
SELECT 
    ABS(-10) AS resultado;
-- Retorna: 10
```

**Analogia Visual**: É como espelhar números negativos:
```
Antes: -5
       ↓ [Máquina ABS]
Depois: 5
```

---

### 4.3 MOD: O Calculador de Resto

**Analogia**: Pense em MOD como uma **calculadora de resto** que mostra o que sobra de uma divisão.

**Exemplo do Dia a Dia**:
```
10 ÷ 3 = 3 com resto 1
MOD(10, 3) = 1 (o resto)

8 ÷ 2 = 4 com resto 0
MOD(8, 2) = 0 (divisão exata)
```

**No Banco de Dados**:
```sql
-- Verificar se quantidade é par ou ímpar
SELECT 
    titulo,
    quantidade_disponivel,
    quantidade_disponivel % 2 AS resto
FROM livros;

-- Se resto = 0 → par
-- Se resto = 1 → ímpar
```

**Analogia Visual**: É como dividir balas entre pessoas e ver o que sobra:
```
10 balas ÷ 3 pessoas = 3 balas cada, sobra 1 bala
MOD(10, 3) = 1 (a bala que sobrou)
```

---

### 4.4 ROUND: O Arredondador Inteligente

**Analogia**: Pense em ROUND como um **arredondador** que vai para o número mais próximo.

**Exemplo do Dia a Dia**:
```
Número: 4.7
ROUND: 5 (mais próximo de 5)

Número: 4.2
ROUND: 4 (mais próximo de 4)

Número: 4.5
ROUND: 5 (meio termo vai para cima)
```

**No Banco de Dados**:
```sql
-- Arredondar média de estoque
SELECT 
    ROUND(AVG(quantidade_disponivel), 2) AS media_arredondada
FROM livros;

-- Resultado: 12.345 → 12.35 (2 casas decimais)
```

**Analogia Visual**: É como escolher o número mais próximo:
```
4.2 → [4] ← mais próximo
4.7 → [5] ← mais próximo
```

---

### 4.5 CEILING: O Arredondador para Cima

**Analogia**: Pense em CEILING como um **arredondador** que sempre vai para cima, mesmo que seja pouco.

**Exemplo do Dia a Dia**:
```
Número: 4.1
CEILING: 5 (sempre para cima, mesmo sendo pouco)

Número: 4.9
CEILING: 5 (também para cima)
```

**No Banco de Dados**:
```sql
-- Calcular prateleiras necessárias
SELECT CEIL(23 / 10.0) AS prateleiras;
-- 23 livros, 10 por prateleira = 3 prateleiras (sempre arredonda para cima)
```

**Analogia Visual**: É como sempre pegar o próximo número inteiro maior:
```
4.1 → [5] (sempre para cima)
4.9 → [5] (sempre para cima)
```

---

## 5. Conditional Functions: A Máquina de Decisão

### 5.1 CASE: O Tomador de Decisões

**Analogia**: Pense em CASE como um **fluxograma** que toma decisões baseadas em condições.

**Exemplo do Dia a Dia**:
```
Se estoque = 0 → "Esgotado"
Se estoque < 5 → "Estoque Baixo"
Se estoque < 10 → "Estoque Médio"
Senão → "Estoque Alto"
```

**No Banco de Dados**:
```sql
-- Classificar estoque
SELECT 
    titulo,
    quantidade_disponivel,
    CASE
        WHEN quantidade_disponivel = 0 THEN 'Esgotado'
        WHEN quantidade_disponivel < 5 THEN 'Estoque Baixo'
        WHEN quantidade_disponivel < 10 THEN 'Estoque Médio'
        ELSE 'Estoque Alto'
    END AS status
FROM livros;
```

**Analogia Visual**: É como um fluxograma:
```
Quantidade?
├─ = 0 → "Esgotado"
├─ < 5 → "Estoque Baixo"
├─ < 10 → "Estoque Médio"
└─ Outro → "Estoque Alto"
```

**Uso Prático**: Classificações e categorizações
```sql
-- Classificar empréstimos
SELECT 
    id,
    CASE
        WHEN status = 'ativo' AND data_devolucao_prevista < DATE('now') THEN 'Atrasado'
        WHEN status = 'ativo' THEN 'Em Andamento'
        ELSE 'Devolvido'
    END AS status_detalhado
FROM emprestimos;
```

---

### 5.2 NULLIF: O Comparador Especial

**Analogia**: Pense em NULLIF como um **detector** que transforma valores específicos em NULL.

**Exemplo do Dia a Dia**:
```
Valor: "Desconhecida"
NULLIF compara: "Desconhecida" = "Desconhecida" → NULL

Valor: "Editora X"
NULLIF compara: "Editora X" ≠ "Desconhecida" → "Editora X" (mantém)
```

**No Banco de Dados**:
```sql
-- Tratar "Desconhecida" como NULL
SELECT 
    titulo,
    NULLIF(editora, 'Desconhecida') AS editora_limpa
FROM livros;

-- Se editora = "Desconhecida" → NULL
-- Se editora = "Editora X" → "Editora X"
```

**Analogia Visual**: É como um filtro que remove valores específicos:
```
Entrada: "Desconhecida"
         ↓ [NULLIF remove]
Saída:   NULL
```

**Uso Prático**: Limpeza de dados
```sql
-- Tratar string vazia como NULL
SELECT NULLIF(telefone, '') AS telefone_ou_null
FROM usuarios;
```

---

### 5.3 COALESCE: O Escolhedor do Primeiro Disponível

**Analogia**: Pense em COALESCE como um **escolhedor** que pega o primeiro valor que não seja NULL.

**Exemplo do Dia a Dia**:
```
Opção 1: NULL
Opção 2: "Não informado"
Opção 3: "João"

COALESCE escolhe: "Não informado" (primeiro não-NULL)
```

**No Banco de Dados**:
```sql
-- Usar valor padrão se NULL
SELECT 
    nome,
    COALESCE(telefone, 'Não informado') AS telefone_display
FROM usuarios;

-- Se telefone = NULL → "Não informado"
-- Se telefone = "123" → "123"
```

**Analogia Visual**: É como escolher o primeiro item disponível:
```
Opções: [NULL] [Não informado] [João]
        ↓ [COALESCE pega primeiro não-NULL]
Escolhido: "Não informado"
```

**Uso Prático**: Valores padrão
```sql
-- Priorizar valores
SELECT 
    COALESCE(data_devolucao_real, data_devolucao_prevista, 'Não definida') AS data_final
FROM emprestimos;
```

---

## 6. Combinando Funções: A Fábrica Completa

### Analogia: Linha de Produção

Pense em combinar funções como uma **linha de produção** onde cada máquina (função) faz sua parte:

```
Matéria-prima: "Dom Casmurro"
    ↓
[Máquina UPPER] → "DOM CASMURRO"
    ↓
[Máquina SUBSTR] → "DOM"
    ↓
[Máquina CONCAT] → "DOM - Disponível"
    ↓
Produto final: "DOM - Disponível"
```

**No Banco de Dados**:
```sql
-- Combinar múltiplas funções
SELECT 
    CONCAT(
        UPPER(SUBSTR(titulo, 1, 3)),
        ' - ',
        CASE 
            WHEN quantidade_disponivel > 0 THEN 'Disponível'
            ELSE 'Esgotado'
        END
    ) AS codigo_status
FROM livros;
```

**Exemplo Completo**:
```sql
-- Análise complexa combinando várias funções
SELECT 
    UPPER(SUBSTR(c.nome, 1, 3)) AS codigo,
    COUNT(l.id) AS total,
    ROUND(AVG(l.quantidade_disponivel), 2) AS media,
    CASE
        WHEN COUNT(l.id) > 5 THEN 'Grande'
        ELSE 'Pequena'
    END AS tamanho
FROM categorias c
LEFT JOIN livros l ON c.id = l.categoria_id
GROUP BY c.id, c.nome;
```

---

## 7. Quando Usar Cada Tipo de Função

### Guia Rápido de Decisão

**Use String Functions quando:**
- Precisa formatar ou limpar texto
- Precisa buscar sem considerar maiúsculas/minúsculas
- Precisa extrair partes de texto
- Precisa combinar textos

**Use Date/Time Functions quando:**
- Precisa calcular datas futuras ou passadas
- Precisa agrupar por período (ano, mês)
- Precisa calcular diferenças de tempo
- Precisa extrair partes de datas

**Use Numeric Functions quando:**
- Precisa arredondar números
- Precisa calcular resto de divisão
- Precisa garantir valores positivos
- Precisa fazer cálculos matemáticos

**Use Conditional Functions quando:**
- Precisa classificar ou categorizar dados
- Precisa fornecer valores padrão
- Precisa tomar decisões baseadas em condições
- Precisa tratar valores NULL

---

## 8. Resumo Visual

### String Functions
```
Texto → [CONCAT] → Texto Combinado
Texto → [LENGTH] → Número
Texto → [SUBSTR] → Parte do Texto
Texto → [REPLACE] → Texto Substituído
Texto → [UPPER] → Texto Maiúsculo
Texto → [LOWER] → Texto Minúsculo
```

### Date/Time Functions
```
Data → [DATE] → Apenas Data
Data → [strftime] → Parte da Data
Data → [date(..., '+N days')] → Data Futura
```

### Numeric Functions
```
Número → [FLOOR] → Inteiro para Baixo
Número → [ABS] → Sempre Positivo
Número → [MOD] → Resto da Divisão
Número → [ROUND] → Arredondado
Número → [CEIL] → Inteiro para Cima
```

### Conditional Functions
```
Valor → [CASE] → Decisão Baseada em Condição
Valor → [NULLIF] → NULL se Igual
Valores → [COALESCE] → Primeiro Não-NULL
```

---

## 9. Dicas Práticas

### Dica 1: Teste Funções Separadamente
```sql
-- Teste cada função sozinha primeiro
SELECT UPPER('teste');  -- Ver resultado
SELECT LENGTH('teste'); -- Ver resultado
-- Depois combine
```

### Dica 2: Use Aliases para Legibilidade
```sql
-- ✅ BOM: alias claro
SELECT 
    UPPER(titulo) AS titulo_maiusculo,
    LENGTH(titulo) AS tamanho_titulo
FROM livros;

-- ❌ EVITE: sem alias
SELECT UPPER(titulo), LENGTH(titulo) FROM livros;
```

### Dica 3: Aninhe com Cuidado
```sql
-- ✅ BOM: legível
SELECT 
    UPPER(SUBSTR(titulo, 1, 3)) AS codigo
FROM livros;

-- ❌ EVITE: muito aninhado e confuso
SELECT UPPER(REPLACE(SUBSTR(CONCAT(titulo, ' - ', autor), 1, 10), ' ', '_')) FROM livros;
```

### Dica 4: Documente Lógica Complexa
```sql
-- Adicione comentários para lógica complexa
SELECT 
    -- Criar código de 3 letras em maiúsculas
    UPPER(SUBSTR(titulo, 1, 3)) AS codigo,
    -- Classificar por quantidade
    CASE
        WHEN quantidade_disponivel = 0 THEN 'Esgotado'
        ELSE 'Disponível'
    END AS status
FROM livros;
```

---

## 10. Conclusão

Funções SQL são como **ferramentas mágicas** que transformam seus dados:

- **String Functions**: Transformam textos
- **Date/Time Functions**: Trabalham com tempo
- **Numeric Functions**: Calculam números
- **Conditional Functions**: Tomam decisões

**Lembre-se**: 
- Pratique cada função separadamente
- Combine funções para criar transformações complexas
- Use aliases para tornar queries legíveis
- Teste sempre no banco de dados real

**Próximo Passo**: Complete os exercícios práticos para dominar essas funções!

---

**Bons estudos! 🚀**

**Lembre-se**: Funções SQL são poderosas e permitem fazer muito trabalho diretamente no banco de dados. Quanto mais você praticar, mais natural será usar essas ferramentas!

