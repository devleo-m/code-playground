# **Aula 3 - Exercícios e Reflexão**

## 📝 Instruções

Complete os exercícios abaixo usando **suas próprias palavras** e, quando solicitado, escreva código SQL. Esta aula é prática - você vai aplicar os conceitos do modelo de objetos do PostgreSQL!

**Importante:** Não copie e cole. Demonstre que você entendeu a hierarquia e como trabalhar com objetos!

---

## 🎯 Exercício 1: Entendendo a Hierarquia ORDBMS

### 1.1 - O que é ORDBMS?

PostgreSQL é um ORDBMS (Object-Relational Database Management System), não apenas um RDBMS.

**Explique:**

- O que significa "Object-Relational"?
- Quais são as duas características que PostgreSQL combina?
- Dê um exemplo prático de um recurso "orientado a objetos" que PostgreSQL oferece.

**Sua resposta:**

```
[ESCREVA SUA RESPOSTA AQUI - MÍNIMO 5 LINHAS]
```

### 1.2 - Hierarquia de Objetos

Coloque os seguintes elementos na ordem hierárquica correta (do mais alto/abrangente ao mais específico/detalhado):

**Elementos:** Row, Database, Column, Schema, Table, Server

**Sua resposta (ordem correta):**

```
1. [...]
2. [...]
3. [...]
4. [...]
5. [...]
6. [...]
```

### 1.3 - Database vs Schema

**Explique a diferença entre Database e Schema:**

- Quando você usaria múltiplos databases?
- Quando você usaria múltiplos schemas no mesmo database?
- Por que você NÃO pode fazer queries entre databases diferentes facilmente?

**Sua resposta:**

```
[ESCREVA SUA RESPOSTA AQUI - MÍNIMO 5 LINHAS]
```

---

## 🎯 Exercício 2: Trabalhando com Databases e Schemas

### 2.1 - Criando Estrutura

Você foi contratado para criar o banco de dados de uma **universidade**. A universidade tem:

- Dados acadêmicos (alunos, professores, disciplinas)
- Dados administrativos (funcionários, departamentos, salários)
- Dados da biblioteca (livros, empréstimos)

**Sua tarefa:**

a) Você criaria 1 database ou 3 databases? Por quê?

```
[SUA DECISÃO E JUSTIFICATIVA]
```

b) Escreva os comandos SQL para criar a estrutura (database + schemas) que você propôs:

```sql
[ESCREVA OS COMANDOS SQL AQUI]
```

c) Como você se conectaria ao database criado?

```sql
[COMANDO AQUI]
```

d) Como você criaria uma tabela no schema de biblioteca?

```sql
[EXEMPLO DE COMANDO AQUI]
```

---

## 🎯 Exercício 3: Tipos de Dados

### 3.1 - Escolhendo Tipos Corretos

Para cada campo abaixo, escolha o tipo de dado mais apropriado e **justifique** sua escolha:

a) **CPF de uma pessoa**

```
Tipo escolhido:
Justificativa:
```

b) **Preço de um produto** (ex: R$ 199,90)

```
Tipo escolhido:
Justificativa:
```

c) **Descrição longa de um artigo de blog**

```
Tipo escolhido:
Justificativa:
```

d) **Status de um pedido** (pode ser: pendente, processando, enviado, entregue, cancelado)

```
Tipo escolhido:
Justificativa:
```

e) **Especificações técnicas de produtos** (cada produto tem atributos diferentes)

```
Tipo escolhido:
Justificativa:
Exemplo de dado:
```

f) **Quantidade de produtos em estoque**

```
Tipo escolhido:
Justificativa:
```

g) **Data e hora em que um registro foi criado**

```
Tipo escolhido:
Justificativa:
```

h) **Tags de um post** (ex: ['postgresql', 'database', 'sql'])

```
Tipo escolhido:
Justificativa:
```

### 3.2 - Identificando Erros de Tipo

Para cada situação, identifique se daria **ERRO** ou seria **PERMITIDO**, e explique por quê:

```sql
CREATE TABLE produtos (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(50),
    preco DECIMAL(10, 2),
    estoque INTEGER,
    ativo BOOLEAN
);
```

a) `INSERT INTO produtos (nome, preco, estoque, ativo) VALUES ('Mouse', 50.00, 100, TRUE);`

```
Resultado: [ERRO ou PERMITIDO?]
Por quê:
```

b) `INSERT INTO produtos (nome, preco, estoque, ativo) VALUES ('Teclado', '150 reais', 50, TRUE);`

```
Resultado: [ERRO ou PERMITIDO?]
Por quê:
```

c) `INSERT INTO produtos (nome, preco, estoque, ativo) VALUES ('Monitor', 800.999, 25, TRUE);`

```
Resultado: [ERRO ou PERMITIDO?]
Por quê:
```

d) `INSERT INTO produtos (nome, preco, estoque) VALUES ('Webcam', 200.00, 30);`

```
Resultado: [ERRO ou PERMITIDO?]
Por quê:
O que acontece com a coluna 'ativo'?
```

e) `INSERT INTO produtos (preco, estoque, ativo) VALUES (100.00, 15, FALSE);`

```
Resultado: [ERRO ou PERMITIDO?]
Por quê:
```

---

## 🎯 Exercício 4: Modelagem Completa

Você precisa criar o banco de dados para um **sistema de uma clínica médica**. O sistema precisa guardar:

**Requisitos:**

- **Médicos**: CRM, nome, especialidade, telefone
- **Pacientes**: CPF, nome, data de nascimento, endereço, telefone, email
- **Consultas**: qual médico atendeu qual paciente, data e hora, diagnóstico, observações
- **Status das consultas**: pode ser "agendada", "realizada", "cancelada"

### 4.1 - Desenhe a Estrutura Completa

```sql
-- 1. Criar database
[SEU CÓDIGO AQUI]

-- 2. Conectar ao database
[SEU CÓDIGO AQUI]

-- 3. Criar schemas (se necessário)
[SEU CÓDIGO AQUI OU EXPLIQUE POR QUE NÃO PRECISA]

-- 4. Criar tipo ENUM para status
[SEU CÓDIGO AQUI]

-- 5. Criar tabela de médicos
[SEU CÓDIGO AQUI]

-- 6. Criar tabela de pacientes
[SEU CÓDIGO AQUI]

-- 7. Criar tabela de consultas
[SEU CÓDIGO AQUI]
```

### 4.2 - Justifique suas Escolhas

a) Que tipo você escolheu para o CRM do médico? Por quê?

```
[SUA RESPOSTA]
```

b) Que tipo você escolheu para data de nascimento? Por quê?

```
[SUA RESPOSTA]
```

c) Que tipo você escolheu para diagnóstico e observações? Por quê?

```
[SUA RESPOSTA]
```

d) Como você garantiu que uma consulta sempre está associada a um médico e um paciente existentes?

```
[SUA RESPOSTA]
```

---

## 🎯 Exercício 5: Queries Práticas

Considerando as tabelas criadas no Exercício 4, escreva queries SQL para:

### 5.1 - Inserir Dados

a) Inserir 3 médicos:

```sql
[SEU CÓDIGO AQUI]
```

b) Inserir 5 pacientes:

```sql
[SEU CÓDIGO AQUI]
```

c) Inserir 10 consultas (diferentes combinações de médicos e pacientes):

```sql
[SEU CÓDIGO AQUI]
```

### 5.2 - Consultas Básicas

a) Listar todos os médicos ordenados por nome:

```sql
[SEU CÓDIGO AQUI]
```

b) Listar todos os pacientes que têm email cadastrado:

```sql
[SEU CÓDIGO AQUI]
```

c) Listar consultas com status "realizada":

```sql
[SEU CÓDIGO AQUI]
```

d) Contar quantos pacientes existem no sistema:

```sql
[SEU CÓDIGO AQUI]
```

### 5.3 - Consultas com WHERE

a) Buscar médicos da especialidade "Cardiologia":

```sql
[SEU CÓDIGO AQUI]
```

b) Buscar pacientes nascidos após 1990:

```sql
[SEU CÓDIGO AQUI]
```

c) Buscar consultas realizadas em dezembro de 2024:

```sql
[SEU CÓDIGO AQUI]
```

d) Buscar consultas que ainda não foram realizadas (agendadas ou canceladas):

```sql
[SEU CÓDIGO AQUI]
```

### 5.4 - Consultas com Agregação

a) Contar quantas consultas cada médico realizou:

```sql
[SEU CÓDIGO AQUI]
```

b) Listar especialidades e quantos médicos existem em cada uma:

```sql
[SEU CÓDIGO AQUI]
```

c) Contar quantas consultas foram realizadas por mês:

```sql
[SEU CÓDIGO AQUI]
```

### 5.5 - Consultas com JOIN

a) Listar todas as consultas mostrando o nome do médico e do paciente:

```sql
[SEU CÓDIGO AQUI]
```

b) Listar pacientes que já tiveram consulta com médico de "Cardiologia":

```sql
[SEU CÓDIGO AQUI]
```

c) Listar médicos que nunca realizaram consultas:

```sql
[SEU CÓDIGO AQUI]
```

---

## 🎯 Exercício 6: Trabalhando com JSON

### 6.1 - Modelagem com JSONB

Você precisa adicionar informações flexíveis ao sistema da clínica:

- Médicos podem ter certificações variadas
- Pacientes podem ter alergias e condições pré-existentes variadas

a) Modifique a tabela de médicos para incluir uma coluna `certificacoes` (JSONB):

```sql
[SEU CÓDIGO AQUI]
```

b) Modifique a tabela de pacientes para incluir uma coluna `info_medica` (JSONB):

```sql
[SEU CÓDIGO AQUI]
```

c) Insira um médico com certificações:

```sql
-- Exemplo de JSON para certificações:
-- {"especialidades": ["Cardiologia", "Clínica Geral"], "anos_experiencia": 15, "idiomas": ["português", "inglês"]}

[SEU CÓDIGO AQUI]
```

d) Insira um paciente com informações médicas:

```sql
-- Exemplo de JSON para info médica:
-- {"alergias": ["penicilina", "látex"], "condicoes": ["diabetes tipo 2"], "tipo_sanguineo": "O+"}

[SEU CÓDIGO AQUI]
```

e) Busque médicos que falam inglês:

```sql
[SEU CÓDIGO AQUI]
```

f) Busque pacientes com alergia a penicilina:

```sql
[SEU CÓDIGO AQUI]
```

---

## 🧠 Perguntas de Reflexão

### Reflexão 1: Database Isolado vs Schema Compartilhado

Imagine que você está desenvolvendo um sistema SaaS (Software as a Service) onde múltiplos clientes (empresas diferentes) usarão o mesmo software, mas cada um precisa ter seus dados isolados.

**Duas abordagens possíveis:**

**Abordagem A:** Um database por cliente

```
- database_cliente_1
- database_cliente_2
- database_cliente_3
```

**Abordagem B:** Um database, um schema por cliente

```
- database_sistema
  - schema_cliente_1
  - schema_cliente_2
  - schema_cliente_3
```

**Reflita:**

- Quais são as vantagens e desvantagens de cada abordagem?
- Qual seria mais fácil de fazer backup?
- Qual seria mais fácil de atualizar (adicionar nova coluna em todas as tabelas)?
- Se você tivesse 1000 clientes, qual abordagem escolheria? Por quê?

**Sua reflexão:**

```
[ESCREVA SEUS PENSAMENTOS AQUI - MÍNIMO 6 LINHAS]
```

---

### Reflexão 2: VARCHAR(n) vs TEXT

PostgreSQL tem dois tipos principais para texto: `VARCHAR(n)` (com limite) e `TEXT` (sem limite).

Alguns desenvolvedores sempre usam `VARCHAR` com limites (ex: `VARCHAR(100)` para nomes).
Outros sempre usam `TEXT` sem limites.

**Argumentos para VARCHAR(n):**

- "Limita o tamanho, evita dados absurdos (ninguém tem nome com 10.000 caracteres)"
- "Documenta expectativa (VARCHAR(100) diz que nome deve ter até 100 chars)"
- "Previne ataques (usuário malicioso não pode enviar 1GB de texto)"

**Argumentos para TEXT:**

- "Performance é igual no PostgreSQL (VARCHAR e TEXT são implementados da mesma forma)"
- "Flexibilidade (se precisar de nome maior, não precisa ALTER TABLE)"
- "Simplicidade (não precisa ficar decidindo 50, 100, 200 chars?)"

**Reflita:**

- Qual abordagem você prefere? Por quê?
- Em que situações você SEMPRE usaria VARCHAR(n) com limite?
- Em que situações você SEMPRE usaria TEXT sem limite?
- Como você lidaria com validação de tamanho se usar TEXT?

**Sua reflexão:**

```
[ESCREVA SEUS PENSAMENTOS AQUI - MÍNIMO 6 LINHAS]
```

---

### Reflexão 3: ENUM vs VARCHAR vs Tabela Separada

Para guardar "status" (ex: pendente, processando, enviado, entregue), existem 3 abordagens:

**Abordagem A: VARCHAR**

```sql
CREATE TABLE pedidos (
    id SERIAL PRIMARY KEY,
    status VARCHAR(20) CHECK (status IN ('pendente', 'processando', 'enviado', 'entregue'))
);
```

**Abordagem B: ENUM**

```sql
CREATE TYPE status_pedido AS ENUM ('pendente', 'processando', 'enviado', 'entregue');
CREATE TABLE pedidos (
    id SERIAL PRIMARY KEY,
    status status_pedido
);
```

**Abordagem C: Tabela separada**

```sql
CREATE TABLE status (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(20) UNIQUE
);
CREATE TABLE pedidos (
    id SERIAL PRIMARY KEY,
    status_id INTEGER REFERENCES status(id)
);
```

**Reflita:**

- Qual é a mais simples? Qual é a mais complexa?
- O que acontece se você precisar adicionar um novo status ("cancelado") em cada abordagem?
- O que acontece se você precisar renomear um status ("enviado" → "em_transito")?
- Se os status mudam raramente (anos), qual você escolheria?
- Se os status mudam frequentemente (semanalmente), qual você escolheria?
- E se status tivessem informações adicionais (ex: cor para exibir na UI, descrição)?

**Sua reflexão:**

```
[ESCREVA SEUS PENSAMENTOS AQUI - MÍNIMO 7 LINHAS]
```

---

## 📤 Como Enviar Suas Respostas

1. Copie este arquivo
2. Preencha **todos** os exercícios e reflexões
3. Envie para análise

**Critérios de avaliação:**

- ✅ Compreensão da hierarquia de objetos
- ✅ Escolha correta de tipos de dados
- ✅ Qualidade do código SQL
- ✅ Capacidade de modelagem completa
- ✅ Profundidade nas reflexões
- ✅ Aplicação prática dos conceitos

---

## ⏱️ Tempo Estimado

- Exercícios 1-3: 40-50 minutos
- Exercício 4: 30-40 minutos
- Exercício 5: 50-60 minutos
- Exercício 6: 30-40 minutos
- Reflexões: 30-40 minutos
- **Total: 3-4 horas**

Este é um exercício PRÁTICO e extenso. Reserve tempo adequado!

---

## 🎯 Próximos Passos

Após enviar suas respostas, você receberá:

- 📊 Análise do seu código SQL
- ⚠️ Erros conceituais ou de sintaxe
- ✅ Pontos onde você demonstrou domínio
- 🎯 Áreas que precisam de mais prática
- 📝 **Avaliação final das 3 aulas** para determinar se você está pronto para avançar

Esta é a **última aula antes da prova**! Demonstre todo o conhecimento adquirido! 💪

Boa sorte! 🚀

