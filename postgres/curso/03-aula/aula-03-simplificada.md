# **Aula 3 - Simplificada: Entendendo o Object Model do PostgreSQL**

## 🎯 Vamos simplificar os objetos do PostgreSQL!

---

## 🎭 PostgreSQL: O Super-Herói Híbrido

### ORDBMS = Super-Herói com Dois Poderes

Imagine que a maioria dos bancos de dados são como super-heróis com **um poder**:

**RDBMS** (Relacional) = Super-herói com poder de **organização perfeita** 📊

- Tudo em tabelas certinhas
- Regras rígidas
- Muito confiável

**OODBMS** (Orientado a Objetos) = Super-herói com poder de **flexibilidade** 🦸

- Pode criar formas customizadas
- Objetos herdam características
- Muito adaptável

**PostgreSQL (ORDBMS)** = Super-herói com **AMBOS os poderes!** 🦸‍♂️⚡

- Organização perfeita DE tabelas + Flexibilidade DE objetos

---

## 🏢 A Hierarquia: Do Prédio até a Gaveta

Pense no PostgreSQL como um **edifício de escritórios**:

```
🏢 EDIFÍCIO (Servidor PostgreSQL)
  │
  ├─ 🏬 ANDAR 1 (Database: "loja")
  │   │
  │   ├─ 🚪 SALA A (Schema: "public")
  │   │   │
  │   │   ├─ 🗄️ ARQUIVO 1 (Table: "clientes")
  │   │   │   │
  │   │   │   ├─ 📋 DIVISÓRIAS (Columns: id, nome, email)
  │   │   │   │
  │   │   │   └─ 📄 DOCUMENTOS (Rows: dados dos clientes)
  │   │   │
  │   │   └─ 🗄️ ARQUIVO 2 (Table: "pedidos")
  │   │
  │   └─ 🚪 SALA B (Schema: "vendas")
  │       └─ 🗄️ ARQUIVO 3 (Table: "comissoes")
  │
  └─ 🏬 ANDAR 2 (Database: "blog")
      └─ 🚪 SALA C (Schema: "public")
          ├─ 🗄️ ARQUIVO 4 (Table: "posts")
          └─ 🗄️ ARQUIVO 5 (Table: "comentarios")
```

**Regras do edifício:**

- **Você não pode pegar arquivo de outro andar** (databases são isolados)
- **Você pode pegar arquivo de outra sala do mesmo andar** (schemas no mesmo database)
- **Cada arquivo tem divisórias fixas** (colunas definidas)
- **Cada documento segue o formato do arquivo** (linhas seguem estrutura das colunas)

---

## 🗄️ 1. Database = Andar Inteiro do Prédio

Um **database** é como um **andar completo** no prédio:

```
🏬 ANDAR "LOJA" (Database)
├─ Sala Vendas
├─ Sala Estoque
└─ Sala Clientes

🏬 ANDAR "BLOG" (Database)
├─ Sala Posts
└─ Sala Comentarios
```

**Características:**

- ✅ Cada andar é **independente**
- ✅ Você precisa subir/descer para ir de um andar para outro
- ❌ Não pode passar documentos entre andares facilmente

**Quando ter múltiplos andares (databases)?**

- ✅ Aplicações completamente diferentes (loja, blog, CRM)
- ✅ Dados de clientes diferentes (cada cliente = um database)
- ❌ Apenas organização (use salas/schemas no mesmo andar)

---

## 📂 2. Schema = Sala no Andar

Um **schema** é como uma **sala** dentro do andar:

```
🏬 ANDAR "LOJA"
├─ 🚪 SALA "public" (padrão)
│   ├─ Arquivo: clientes
│   └─ Arquivo: produtos
│
├─ 🚪 SALA "vendas"
│   ├─ Arquivo: pedidos
│   └─ Arquivo: comissoes
│
└─ 🚪 SALA "estoque"
    ├─ Arquivo: movimentacoes
    └─ Arquivo: inventario
```

**Por que ter salas separadas?**

- 🗂️ **Organização**: Arquivos relacionados ficam na mesma sala
- 🔐 **Segurança**: Pode trancar a sala e só dar chave para alguns
- 🏷️ **Evitar confusão**: Pode ter "pedidos" na sala vendas E na sala compras

**Sala "public":**

- Todo andar tem uma sala chamada "public" (padrão)
- Se você não disser qual sala, vai para "public"

---

## 📊 3. Table = Arquivo (Gaveteiro)

Uma **table** é como um **arquivo (gaveteiro)** na sala:

```
🗄️ ARQUIVO: "clientes"

┌────────────────────────────────────────┐
│ 📋 DIVISÓRIAS (COLUMNS):               │
│ [ID] [NOME] [EMAIL] [TELEFONE]        │
├────────────────────────────────────────┤
│ 📄 DOCUMENTO 1 (ROW):                  │
│ [1] [João] [joao@...] [11-9999-9999]  │
├────────────────────────────────────────┤
│ 📄 DOCUMENTO 2 (ROW):                  │
│ [2] [Maria] [maria@...] [11-8888-8888]│
└────────────────────────────────────────┘
```

**Regras do arquivo:**

- ✅ Todas as divisórias (colunas) são fixas
- ✅ Todo documento (linha) precisa seguir as divisórias
- ✅ Não pode ter documento sem ID (se ID é obrigatório)
- ✅ Não pode ter dois documentos com mesmo ID (se ID é único)

---

## 📋 4. Column = Divisória do Arquivo

Uma **column** é como uma **divisória** no arquivo:

```
ARQUIVO: clientes

Divisória 1: ID (número, único, obrigatório)
Divisória 2: NOME (texto até 100 letras, obrigatório)
Divisória 3: EMAIL (texto até 100 letras, único)
Divisória 4: TELEFONE (texto até 20 caracteres, opcional)
Divisória 5: ATIVO (sim/não, padrão: sim)
```

**Cada divisória define:**

- 📏 **Tipo**: Que tipo de coisa vai aqui? (número, texto, data...)
- ✅ **Regras**: É obrigatório? Tem que ser único? Tem valor padrão?
- 🏷️ **Nome**: Como vou chamar essa informação?

**Tipos comuns de divisórias:**

- 🔢 **Número**: ID, idade, quantidade, preço
- 📝 **Texto**: Nome, email, descrição
- 📅 **Data/Hora**: Data de nascimento, data de criação
- ☑️ **Sim/Não**: Ativo, promocao, disponível
- 📦 **JSON**: Informações flexíveis (especificações de produto)

---

## 📄 5. Row = Documento no Arquivo

Uma **row** é um **documento individual** dentro do arquivo:

```
Documento 1 do arquivo "clientes":
┌────┬──────┬─────────────────┬──────────────┬───────┐
│ 1  │ João │ joao@email.com  │ 11-9999-9999 │ TRUE  │
└────┴──────┴─────────────────┴──────────────┴───────┘
  ↑     ↑          ↑               ↑            ↑
  ID   Nome      Email         Telefone       Ativo
```

**Características:**

- Cada documento é **único** (mesmo que informações pareçam iguais)
- **Não há ordem** dos documentos (a menos que você peça ordem específica)
- Cada documento tem valor para **todas as divisórias** (pode ser vazio se permitido)

---

## 🔍 6. Query = Pedido ao Bibliotecário

**Query** é como você **pedir informações ao bibliotecário** do prédio:

### Pedidos Simples

```sql
-- "Me mostre todos os documentos do arquivo clientes"
SELECT * FROM clientes;

-- "Me mostre só o nome e email dos documentos"
SELECT nome, email FROM clientes;

-- "Me mostre clientes que têm email"
SELECT * FROM clientes WHERE email IS NOT NULL;
```

### Pedidos com Filtros (WHERE = "Só me mostre se...")

```sql
-- "Só clientes ativos"
SELECT * FROM clientes WHERE ativo = TRUE;

-- "Só produtos com preço maior que 100"
SELECT * FROM produtos WHERE preco > 100;

-- "Só produtos da categoria Livros OU Eletrônicos"
SELECT * FROM produtos
WHERE categoria IN ('Livros', 'Eletrônicos');

-- "Produtos que começam com 'Note'"
SELECT * FROM produtos WHERE nome LIKE 'Note%';
```

### Pedidos com Ordem (ORDER BY = "Organize assim...")

```sql
-- "Me mostre produtos do mais barato ao mais caro"
SELECT * FROM produtos ORDER BY preco ASC;

-- "Me mostre produtos do mais caro ao mais barato"
SELECT * FROM produtos ORDER BY preco DESC;

-- "Organize por categoria, e dentro de cada categoria, por preço"
SELECT * FROM produtos ORDER BY categoria, preco;
```

### Pedidos Contando/Somando (Funções de Agregação)

```sql
-- "Quantos clientes eu tenho?"
SELECT COUNT(*) FROM clientes;

-- "Qual o preço médio dos produtos?"
SELECT AVG(preco) FROM produtos;

-- "Qual o produto mais caro?"
SELECT MAX(preco) FROM produtos;

-- "Quanto vale todo meu estoque?"
SELECT SUM(preco * quantidade) FROM produtos;
```

### Pedidos Agrupados (GROUP BY = "Agrupe e conte...")

```sql
-- "Quantos produtos tenho em cada categoria?"
SELECT categoria, COUNT(*) AS total
FROM produtos
GROUP BY categoria;

-- "Qual o preço médio por categoria?"
SELECT categoria, AVG(preco) AS preco_medio
FROM produtos
GROUP BY categoria;
```

### Pedidos Combinando Arquivos (JOIN = "Junte informações de...")

```sql
-- "Me mostre pedidos COM o nome do cliente"
SELECT
    pedidos.id,
    clientes.nome,
    pedidos.produto,
    pedidos.valor
FROM pedidos
JOIN clientes ON pedidos.cliente_id = clientes.id;
```

---

## 📦 7. Tipos de Dados = Tipos de Divisórias

Cada **divisória (coluna)** só aceita um **tipo** de informação:

### 🔢 Divisória para Números

```sql
SMALLINT    -- Números pequenos (-32.768 a 32.767)
INTEGER     -- Números normais (-2 bilhões a 2 bilhões)
BIGINT      -- Números gigantes
DECIMAL     -- Dinheiro (preciso)
SERIAL      -- Número que aumenta sozinho (1, 2, 3, ...)
```

**Exemplo do mundo real:**

- **Idade**: SMALLINT (ninguém tem 1.000 anos)
- **População de cidade**: INTEGER
- **PIB de país**: BIGINT
- **Preço de produto**: DECIMAL(10, 2) - ex: 199.99

### 📝 Divisória para Texto

```sql
VARCHAR(100)  -- Texto até 100 letras
TEXT          -- Texto sem limite
CHAR(2)       -- Texto SEMPRE com 2 letras (ex: 'SP', 'RJ')
```

**Exemplo do mundo real:**

- **Estado**: CHAR(2) - 'SP', 'RJ', 'MG'
- **Nome**: VARCHAR(100)
- **Descrição longa**: TEXT

### 📅 Divisória para Data/Hora

```sql
DATE          -- Só a data: 2024-12-01
TIME          -- Só a hora: 14:30:00
TIMESTAMP     -- Data e hora juntas: 2024-12-01 14:30:00
```

**Exemplo do mundo real:**

- **Data de nascimento**: DATE
- **Horário de abertura**: TIME
- **Momento de criação**: TIMESTAMP

### ☑️ Divisória para Sim/Não

```sql
BOOLEAN       -- TRUE (sim) ou FALSE (não)
```

**Exemplo do mundo real:**

- **Cliente ativo?**: BOOLEAN
- **Produto em promoção?**: BOOLEAN
- **Email verificado?**: BOOLEAN

### 📦 Divisória para Informações Flexíveis

```sql
JSONB         -- Como uma caixinha com coisas variadas
```

**Exemplo do mundo real:**

```sql
-- Produtos têm especificações diferentes:
Notebook: {marca: "Dell", ram: "16GB", tela: 15.6}
Mouse: {marca: "Logitech", dpi: 1600, sem_fio: true}
Livro: {autor: "João", páginas: 300, idioma: "PT"}
```

Cada produto tem informações diferentes, mas tudo cabe na divisória JSONB!

---

## 🎯 O Poder das Queries: Analogia do Google

Fazer **queries** no PostgreSQL é como fazer buscas no Google:

```
Google (simples):
"receita de bolo"

PostgreSQL (simples):
SELECT * FROM receitas WHERE tipo = 'bolo';

---

Google (complexo):
"receita de bolo de chocolate publicada em 2024 ordenar por avaliação"

PostgreSQL (complexo):
SELECT * FROM receitas
WHERE tipo = 'bolo'
  AND ingrediente = 'chocolate'
  AND ano = 2024
ORDER BY avaliacao DESC;
```

**Diferença:** Google busca texto, PostgreSQL busca dados estruturados!

---

## 🏗️ Construindo do Zero: Passo a Passo

Imagine que você vai abrir uma loja online. Vamos construir:

### Passo 1: Construir o Edifício (Servidor já existe)

```sql
-- Servidor PostgreSQL já está rodando
```

### Passo 2: Criar um Andar (Database)

```sql
CREATE DATABASE loja;
```

### Passo 3: Entrar no Andar

```sql
\c loja  -- Conectar ao database
```

### Passo 4: Criar Sala (Schema) - Opcional

```sql
CREATE SCHEMA vendas;
-- Mas pode usar a sala 'public' que já existe!
```

### Passo 5: Criar Arquivo (Table)

```sql
CREATE TABLE clientes (
    id SERIAL PRIMARY KEY,      -- Divisória 1: número automático
    nome VARCHAR(100) NOT NULL, -- Divisória 2: nome obrigatório
    email VARCHAR(100) UNIQUE,  -- Divisória 3: email único
    ativo BOOLEAN DEFAULT TRUE  -- Divisória 4: ativo (padrão: sim)
);
```

### Passo 6: Adicionar Documentos (Rows)

```sql
INSERT INTO clientes (nome, email)
VALUES ('João', 'joao@email.com');

INSERT INTO clientes (nome, email)
VALUES ('Maria', 'maria@email.com');
```

### Passo 7: Buscar Documentos (Query)

```sql
SELECT * FROM clientes;
SELECT * FROM clientes WHERE ativo = TRUE;
```

---

## 🎓 Resumo Ultra-Simplificado

| Conceito PostgreSQL | Analogia do Mundo Real      |
| ------------------- | --------------------------- |
| **Servidor**        | Edifício inteiro            |
| **Database**        | Andar do edifício           |
| **Schema**          | Sala no andar               |
| **Table**           | Arquivo (gaveteiro)         |
| **Column**          | Divisória do arquivo        |
| **Row**             | Documento no arquivo        |
| **Query**           | Pedido ao bibliotecário     |
| **Data Type**       | Tipo de divisória           |
| **Primary Key**     | Número do documento (único) |
| **Foreign Key**     | Referência a outro arquivo  |

---

## 💡 A Grande Lição

**PostgreSQL é organizado como um edifício de escritórios:**

```
🏢 Um EDIFÍCIO (Servidor)
   └─ Vários ANDARES (Databases)
      └─ Cada andar tem SALAS (Schemas)
         └─ Cada sala tem ARQUIVOS (Tables)
            └─ Cada arquivo tem DIVISÓRIAS fixas (Columns)
               └─ Cada arquivo guarda DOCUMENTOS (Rows)
```

**Para trabalhar com dados:**

1. Entre no andar certo (conecte ao database)
2. Vá na sala certa (use o schema, ou deixe padrão 'public')
3. Abra o arquivo certo (selecione a table)
4. Busque os documentos que precisa (faça a query)

**E a mágica:** O PostgreSQL organiza tudo isso automaticamente! Você só precisa saber pedir o que quer! 🎯

---

## 🎯 Próximo Passo

Agora você vai fazer exercícios práticos para:

- Criar databases, schemas e tables
- Definir colunas com tipos corretos
- Inserir e buscar dados
- Fazer queries complexas

Prepare-se para colocar a mão na massa! 💪

