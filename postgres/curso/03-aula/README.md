# 🏗️ Aula 3: Object Model in PostgreSQL

Bem-vindo à Aula 3! Esta é a **última aula teórica antes da avaliação**. Aqui você vai aprender sobre o modelo de objetos do PostgreSQL (ORDBMS), a hierarquia completa de objetos, tipos de dados e como fazer queries práticas.

---

## 📚 Estrutura da Aula 3

### **Tópico: Object Model in PostgreSQL**

#### 📖 Materiais Disponíveis:

1. **[aula-03-object-model.md](./aula-03-object-model.md)** - Aula principal completa

   - ORDBMS: Combinação de relacional + orientado a objetos
   - Recursos OO: Tipos customizados, herança, polimorfismo
   - Hierarquia completa: Server → Database → Schema → Table → Column → Row
   - Databases: Criação, conexão, gerenciamento
   - Schemas: Organização e namespacing
   - Tables: Criação, modificação, tipos
   - Columns: Definição, constraints, modificação
   - Rows: Inserção, atualização, deleção
   - Queries: SELECT, WHERE, ORDER BY, LIMIT, agregações, GROUP BY, HAVING, JOINs, subqueries
   - Tipos de dados: Numéricos, texto, data/hora, boolean, enum, arrays, JSON/JSONB, geométricos, UUID

2. **[aula-03-simplificada.md](./aula-03-simplificada.md)** - Versão com analogias

   - PostgreSQL = Super-herói híbrido com dois poderes
   - Hierarquia = Edifício de escritórios
   - Database = Andar do prédio
   - Schema = Sala no andar
   - Table = Arquivo (gaveteiro)
   - Column = Divisória do arquivo
   - Row = Documento no arquivo
   - Query = Pedido ao bibliotecário
   - Tipos de dados = Tipos de divisórias

3. **[aula-03-exercicios.md](./aula-03-exercicios.md)** - Exercícios práticos extensos
   - 6 exercícios práticos (modelagem completa de clínica médica)
   - Escrita de código SQL real
   - 3 reflexões profundas (databases vs schemas, VARCHAR vs TEXT, ENUM vs tabela)
   - Tempo estimado: 3-4 horas

---

## 🎯 Como Estudar Esta Aula

### **Passo 1: Leia a Aula Principal** (90-120 minutos)

📖 Abra: `aula-03-object-model.md`

**Esta aula é DENSA e PRÁTICA!** Foque em:

#### Parte 1: Conceitos ORDBMS

- Por que PostgreSQL é "Object-Relational"
- Recursos orientados a objetos (tipos custom, herança, polimorfismo)

#### Parte 2: Hierarquia de Objetos

- Server → Database → Schema → Table → Column → Row
- Como cada nível se relaciona com o próximo
- Quando usar múltiplos databases vs múltiplos schemas

#### Parte 3: Queries

- SELECT básico e avançado
- WHERE (filtros)
- ORDER BY (ordenação)
- LIMIT/OFFSET (paginação)
- Funções de agregação (COUNT, SUM, AVG, MIN, MAX)
- GROUP BY e HAVING
- JOINs (INNER, LEFT, RIGHT, FULL)
- Subqueries

#### Parte 4: Tipos de Dados

- Numéricos (INTEGER, BIGINT, DECIMAL, SERIAL)
- Texto (VARCHAR, TEXT, CHAR)
- Data/Hora (DATE, TIME, TIMESTAMP)
- Boolean
- ENUM
- Arrays
- JSON/JSONB (importante!)
- UUID

**Dica:** Esta aula tem MUITO código SQL. Não apenas leia - tente entender O QUE cada comando faz!

---

### **Passo 2: Leia a Versão Simplificada** (40-60 minutos)

🎯 Abra: `aula-03-simplificada.md`

**A analogia do edifício de escritórios vai clarear tudo!**

Você vai entender:

- Por que databases são isolados (andares diferentes)
- Por que schemas facilitam organização (salas no mesmo andar)
- Por que tabelas têm estrutura fixa (divisórias do arquivo)
- Como queries funcionam (pedidos ao bibliotecário)

**Benefício:** Conceitos abstratos viram imagens mentais claras!

---

### **Passo 3: Faça os Exercícios** (3-4 horas) ⚠️ IMPORTANTE!

✏️ Abra: `aula-03-exercicios.md`

**ATENÇÃO:** Esta é a aula mais prática do curso até agora!

**Estrutura dos exercícios:**

#### **Exercício 1:** Conceitos ORDBMS e Hierarquia

- O que é ORDBMS?
- Ordenar hierarquia
- Database vs Schema

#### **Exercício 2:** Criando Databases e Schemas

- Modelar sistema de universidade
- Decidir estrutura (databases/schemas)
- Escrever comandos SQL

#### **Exercício 3:** Tipos de Dados

- Escolher tipos apropriados para diferentes campos
- Identificar erros de tipo
- Justificar escolhas

#### **Exercício 4:** Modelagem Completa ⭐ FUNDAMENTAL

- Sistema completo de clínica médica
- Criar database, schemas, tables
- Definir colunas com tipos corretos
- Estabelecer relacionamentos

#### **Exercício 5:** Queries Práticas ⭐ FUNDAMENTAL

- Inserir dados
- Consultas básicas (SELECT, WHERE)
- Agregações (COUNT, GROUP BY)
- JOINs entre tabelas
- Queries complexas

#### **Exercício 6:** Trabalhando com JSON

- Adicionar colunas JSONB
- Inserir dados JSON
- Consultar dentro de JSON

#### **Reflexões:**

1. Database isolado vs Schema compartilhado (multi-tenancy)
2. VARCHAR(n) vs TEXT (limites vs flexibilidade)
3. ENUM vs VARCHAR vs Tabela separada (modelagem de status)

---

### **Passo 4: Envie para Avaliação Final** 📤

**Esta é a última aula antes da prova!** Suas respostas serão avaliadas para determinar se você:

- ✅ Dominou a hierarquia de objetos
- ✅ Sabe escolher tipos de dados apropriados
- ✅ Consegue modelar sistemas completos
- ✅ Escreve SQL corretamente
- ✅ Pensa criticamente sobre decisões de design

---

## ⚠️ Por Que Esta Aula é Crucial?

Esta aula é a **ponte entre teoria e prática**:

```
┌──────────────────────────────────────────┐
│  AULA 1: O que é PostgreSQL              │
│  (Visão geral)                           │
├──────────────────────────────────────────┤
│  AULA 2: Modelo Relacional               │
│  (Base teórica - conceitos abstratos)    │
├──────────────────────────────────────────┤
│  AULA 3: Object Model ← VOCÊ ESTÁ AQUI   │
│  (Como PostgreSQL implementa na prática) │
├──────────────────────────────────────────┤
│  ↓ Próximo: PROVA                        │
│  (Avaliação do aprendizado)              │
└──────────────────────────────────────────┘
```

**Se você não dominar esta aula:**

- ❌ Não saberá escrever SQL
- ❌ Não saberá criar tabelas
- ❌ Não saberá escolher tipos de dados
- ❌ Não estará pronto para projetos reais

**Se você dominar esta aula:**

- ✅ Saberá escrever SQL do básico ao intermediário
- ✅ Poderá modelar sistemas reais
- ✅ Entenderá mensagens de erro do PostgreSQL
- ✅ Estará pronto para avançar!

---

## 📊 Conceitos-Chave Desta Aula

Ao final, você deve dominar:

### Hierarquia

- [ ] Servidor → Database → Schema → Table → Column → Row
- [ ] Quando usar múltiplos databases
- [ ] Quando usar múltiplos schemas
- [ ] Schema padrão "public"

### Objetos

- [ ] Como criar e gerenciar databases
- [ ] Como criar e usar schemas
- [ ] Como criar, modificar e deletar tables
- [ ] Como definir columns com tipos e constraints
- [ ] Como inserir, atualizar e deletar rows

### Queries

- [ ] SELECT básico (colunas, alias)
- [ ] WHERE (filtros: =, >, <, LIKE, IN, IS NULL)
- [ ] ORDER BY (ASC, DESC)
- [ ] LIMIT e OFFSET (paginação)
- [ ] Funções de agregação (COUNT, SUM, AVG, MIN, MAX)
- [ ] GROUP BY (agrupamento)
- [ ] HAVING (filtro pós-agregação)
- [ ] JOINs (INNER, LEFT, RIGHT, FULL)
- [ ] Subqueries

### Tipos de Dados

- [ ] Numéricos: INTEGER, BIGINT, DECIMAL, SERIAL
- [ ] Texto: VARCHAR(n), TEXT
- [ ] Data/Hora: DATE, TIME, TIMESTAMP
- [ ] Boolean: TRUE/FALSE
- [ ] ENUM: Valores pré-definidos
- [ ] Arrays: Listas de valores
- [ ] JSONB: Dados semi-estruturados
- [ ] UUID: Identificadores únicos

### Recursos ORDBMS

- [ ] Tipos de dados customizados
- [ ] Herança de tabelas
- [ ] Polimorfismo em queries

---

## 🎯 Checklist de Prontidão

Antes de fazer os exercícios, você deve conseguir:

- [ ] Explicar a diferença entre database e schema
- [ ] Desenhar a hierarquia completa de objetos
- [ ] Criar uma tabela com 5+ colunas de tipos diferentes
- [ ] Escrever SELECT com WHERE e ORDER BY
- [ ] Usar COUNT, AVG, SUM
- [ ] Fazer GROUP BY
- [ ] Fazer INNER JOIN entre duas tabelas
- [ ] Escolher o tipo correto para: ID, nome, preço, data, status
- [ ] Explicar quando usar JSONB

**Menos de 7 marcados:** Releia a aula principal.
**7-9 marcados:** Você está quase pronto!
**Todos marcados:** Vá para os exercícios!

---

## 💡 Dicas de Estudo

### 1. Não Apenas Leia o SQL - Execute! 💻

Se você tem PostgreSQL instalado (ou pode usar online):

- **Execute** os exemplos de código
- **Modifique** e veja o que acontece
- **Experimente** variações

### 2. Desenhe a Hierarquia ✏️

Pegue papel e caneta e desenhe:

```
Seu Edifício PostgreSQL:
- Que databases você teria?
- Que schemas em cada database?
- Que tabelas em cada schema?
- Que colunas em cada tabela?
```

### 3. Pratique Escolha de Tipos 🎯

Para cada sistema que você usa (Instagram, YouTube, Banco):

- Que tipos de dados eles usam?
- Como modelariam usuários, posts, comentários?

### 4. Escreva SQL no Papel Primeiro 📝

Antes de olhar a resposta:

- Escreva a query no papel
- Pense: "O que eu quero selecionar? De onde? Com que filtros?"
- Depois verifique se está correto

### 5. Foque nas Reflexões 🤔

As reflexões não têm resposta "certa". Elas fazem você:

- Pensar em trade-offs
- Considerar cenários reais
- Tomar decisões fundamentadas

---

## 🔄 Conectando com Aulas Anteriores

### Aula 1 → Aula 3

**Aula 1** disse: "PostgreSQL é poderoso e flexível"
**Aula 3** mostra: "JSONB, herança, tipos custom = flexibilidade"

### Aula 2 → Aula 3

**Aula 2** explicou: "Relação, tupla, atributo, domínio"
**Aula 3** traduz: "Table, row, column, data type"

```
TEORIA (Aula 2)          PRÁTICA (Aula 3)
─────────────────────────────────────────
Relação             →    Table
Tupla               →    Row
Atributo            →    Column
Domínio             →    Data Type
Chave Primária      →    PRIMARY KEY
Chave Estrangeira   →    FOREIGN KEY REFERENCES
```

---

## ⏱️ Tempo Total Estimado

| Atividade               | Tempo         |
| ----------------------- | ------------- |
| Leitura principal       | 90-120 min    |
| Leitura simplificada    | 40-60 min     |
| Exercícios 1-3          | 70-90 min     |
| Exercício 4 (modelagem) | 30-40 min     |
| Exercício 5 (queries)   | 50-60 min     |
| Exercício 6 (JSON)      | 30-40 min     |
| Reflexões               | 30-40 min     |
| **Total**               | **5-7 horas** |

**Recomendação:** Divida em 3-4 sessões de estudo!

- **Sessão 1:** Leitura principal + simplificada (2-3h)
- **Sessão 2:** Exercícios 1-3 (1-2h)
- **Sessão 3:** Exercícios 4-5 (1,5-2h)
- **Sessão 4:** Exercício 6 + Reflexões (1-1,5h)

---

## 🎓 Preparação para a Prova

Após completar esta aula, você terá visto:

- ✅ Fundamentos teóricos (Aula 1 e 2)
- ✅ Implementação prática (Aula 3)
- ✅ Modelagem de sistemas reais
- ✅ Escrita de SQL

**A prova vai testar:**

- Compreensão dos conceitos das 3 aulas
- Capacidade de modelar um sistema completo
- Habilidade de escrever SQL correto
- Pensamento crítico sobre decisões de design

---

## 🚀 Mensagem Importante

Esta é a aula mais **prática** e **extensa** até agora. Não tenha pressa!

**É melhor:**

- ✅ Fazer devagar e entender profundamente
- ✅ Praticar cada tipo de query várias vezes
- ✅ Revisar conceitos que não ficaram claros

**Do que:**

- ❌ Correr e responder superficialmente
- ❌ Copiar exemplos sem entender
- ❌ Pular partes "difíceis"

**Lembre-se:** Esta é a última aula antes da avaliação. O tempo que você investir agora vai determinar seu sucesso na prova! 💪

---

## 📞 Dúvidas Frequentes

### "Preciso ter PostgreSQL instalado?"

- **Ideal:** Sim, para praticar os comandos
- **Alternativa:** Pode fazer exercícios teóricos e enviar para correção
- **Nota:** Na prova, você precisará demonstrar conhecimento prático

### "Não entendi JOINs, o que faço?"

- Releia a seção de JOINs na aula principal
- Veja a analogia na versão simplificada
- Desenhe tabelas no papel e como elas se conectam
- Pergunte! JOINs são fundamentais

### "JSON/JSONB é muito importante?"

- **Sim!** É um diferencial do PostgreSQL
- Permite flexibilidade sem perder estrutura
- Muito usado em sistemas modernos
- Dedique tempo a entender

### "Quantas vezes devo revisar?"

- Aula principal: **1 vez** (concentrado)
- Simplificada: **1-2 vezes** (reforço)
- Exercícios: **Fazer todos** (prática essencial)
- Código SQL: **Praticar até fluir naturalmente**

---

## 🎯 Próximo Passo

**Você tem duas opções:**

### Opção A: Começar Agora! 🚀

1. Abra `aula-03-object-model.md`
2. Leia com atenção (90-120 min)
3. Leia a versão simplificada
4. Faça os exercícios

### Opção B: Revisar Aulas Anteriores Primeiro 🔄

Se sente que não dominou Aulas 1-2:

1. Revise os conceitos não claros
2. Refaça exercícios anteriores
3. **Depois** vá para Aula 3

**A escolha é sua!** Mas lembre-se: base sólida é essencial!

---

Boa sorte! Esta é a aula mais desafiadora até agora, mas também a mais recompensadora! 🌟

Você está a um passo de dominar os fundamentos do PostgreSQL! 💪🐘

