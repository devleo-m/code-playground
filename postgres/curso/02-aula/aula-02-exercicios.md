# **Aula 2 - Exercícios e Reflexão**

## 📝 Instruções

Complete os exercícios abaixo usando **suas próprias palavras**. O objetivo é verificar se você compreendeu os conceitos fundamentais do modelo relacional.

**Lembre-se:** Não copie e cole. Demonstre seu entendimento real!

---

## 🎯 Exercício 1: Conceitos Fundamentais de E.F. Codd

### 1.1 - O Problema que Codd Resolveu

Antes de E.F. Codd criar o modelo relacional em 1970, os bancos de dados tinham muitos problemas.

**Explique com suas palavras:**

- Qual era o principal problema que Codd queria resolver?
- Como o modelo relacional (tabelas) resolveu esse problema?

**Sua resposta:**

```
[ESCREVA SUA RESPOSTA AQUI - MÍNIMO 4 LINHAS]
```

### 1.2 - Relação, Tupla, Atributo

Você tem a seguinte tabela:

```
LIVROS
┌────┬──────────────────────┬───────────────┬──────┬───────┐
│ ID │ TITULO               │ AUTOR         │ ANO  │ PRECO │
├────┼──────────────────────┼───────────────┼──────┼───────┤
│ 1  │ 1984                 │ George Orwell │ 1949 │ 45.90 │
│ 2  │ Dom Casmurro         │ Machado Assis │ 1899 │ 35.00 │
│ 3  │ O Cortiço            │ Aluísio Azevedo│ 1890 │ 28.50 │
└────┴──────────────────────┴───────────────┴──────┴───────┘
```

**Identifique:**

a) Qual é a **relação** (nome da tabela)?

```
[SUA RESPOSTA]
```

b) Dê um exemplo de uma **tupla** completa (uma linha):

```
[SUA RESPOSTA]
```

c) Liste todos os **atributos** (colunas):

```
[SUA RESPOSTA]
```

d) Qual seria o **domínio** adequado para o atributo ANO?

```
[SUA RESPOSTA]
```

e) Qual seria o **domínio** adequado para o atributo PRECO?

```
[SUA RESPOSTA]
```

---

## 🎯 Exercício 2: Chaves (Conceito Fundamental!)

### 2.1 - Identificando Chaves

Considere a tabela de FUNCIONÁRIOS de uma empresa:

```
FUNCIONARIOS
┌────┬──────────────┬──────────────┬────────────────────┬────────┬───────────┐
│ ID │ NOME         │ CPF          │ EMAIL              │ CARGO  │ SALARIO   │
├────┼──────────────┼──────────────┼────────────────────┼────────┼───────────┤
│ 1  │ João Silva   │ 111.111.111-11│ joao@empresa.com   │ Dev    │ 5000.00   │
│ 2  │ Maria Santos │ 222.222.222-22│ maria@empresa.com  │ Dev    │ 5000.00   │
│ 3  │ Pedro Costa  │ 333.333.333-33│ pedro@empresa.com  │ Gerente│ 8000.00   │
└────┴──────────────┴──────────────┴────────────────────┴────────┴───────────┘
```

**Responda:**

a) Quais atributos poderiam ser **chaves candidatas** (identificam unicamente um funcionário)?

```
[LISTE TODAS AS CHAVES CANDIDATAS POSSÍVEIS]
```

b) Por que NOME não pode ser chave candidata?

```
[SUA JUSTIFICATIVA]
```

c) Por que CARGO não pode ser chave candidata?

```
[SUA JUSTIFICATIVA]
```

d) Qual chave candidata você escolheria como **chave primária**? Por quê?

```
Chave escolhida:
Justificativa:
```

### 2.2 - Chaves Estrangeiras

Você precisa criar duas tabelas: DEPARTAMENTOS e FUNCIONARIOS, onde cada funcionário pertence a um departamento.

**Desenhe a estrutura das duas tabelas incluindo:**

- Chave primária de cada tabela
- Chave estrangeira que conecta as tabelas
- Pelo menos 3 outros atributos em cada tabela

**Sua resposta:**

```
Tabela: DEPARTAMENTOS
┌────┬...
[COMPLETE AQUI]


Tabela: FUNCIONARIOS
┌────┬...
[COMPLETE AQUI]


Explicação do relacionamento:
[EXPLIQUE COMO AS TABELAS SE CONECTAM]
```

---

## 🎯 Exercício 3: Tipos de Relacionamentos

Para cada cenário abaixo, identifique o tipo de relacionamento (1:1, 1:N ou N:M) e **justifique** sua resposta.

### Cenário A: Biblioteca

- Livros podem ser escritos por vários autores
- Autores podem escrever vários livros

**Tipo de relacionamento:**

```
[1:1, 1:N ou N:M?]
```

**Justificativa:**

```
[POR QUE?]
```

**Como você implementaria no banco? (desenhe as tabelas necessárias)**

```
[DESENHE AS TABELAS AQUI]
```

---

### Cenário B: Hospital

- Cada paciente tem um único prontuário médico
- Cada prontuário pertence a um único paciente

**Tipo de relacionamento:**

```
[1:1, 1:N ou N:M?]
```

**Justificativa:**

```
[POR QUE?]
```

**Como você implementaria no banco? (desenhe as tabelas necessárias)**

```
[DESENHE AS TABELAS AQUI]
```

---

### Cenário C: E-commerce

- Clientes fazem pedidos
- Um cliente pode fazer vários pedidos
- Cada pedido é de um único cliente

**Tipo de relacionamento:**

```
[1:1, 1:N ou N:M?]
```

**Justificativa:**

```
[POR QUE?]
```

**Como você implementaria no banco? (desenhe as tabelas necessárias)**

```
[DESENHE AS TABELAS AQUI]
```

---

## 🎯 Exercício 4: Integridade de Dados

Considere o seguinte banco de dados de uma escola:

```sql
CREATE TABLE alunos (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(100) NOT NULL,
    data_nascimento DATE CHECK (data_nascimento < CURRENT_DATE),
    email VARCHAR(100) UNIQUE
);

CREATE TABLE turmas (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(50) NOT NULL,
    capacidade_maxima INTEGER CHECK (capacidade_maxima > 0)
);

CREATE TABLE matriculas (
    id SERIAL PRIMARY KEY,
    aluno_id INTEGER NOT NULL,
    turma_id INTEGER NOT NULL,
    data_matricula DATE NOT NULL,
    FOREIGN KEY (aluno_id) REFERENCES alunos(id),
    FOREIGN KEY (turma_id) REFERENCES turmas(id),
    UNIQUE(aluno_id, turma_id)
);
```

### 4.1 - Identifique os Tipos de Integridade

Para cada regra abaixo, identifique qual tipo de integridade está sendo aplicada:

- **Integridade de Entidade**
- **Integridade Referencial**
- **Integridade de Domínio**
- **Integridade de Negócio**

**a) `id SERIAL PRIMARY KEY` na tabela alunos**

```
Tipo: [QUAL TIPO?]
Explicação: [POR QUE?]
```

**b) `FOREIGN KEY (aluno_id) REFERENCES alunos(id)` na tabela matriculas**

```
Tipo: [QUAL TIPO?]
Explicação: [POR QUE?]
```

**c) `CHECK (capacidade_maxima > 0)` na tabela turmas**

```
Tipo: [QUAL TIPO?]
Explicação: [POR QUE?]
```

**d) `CHECK (data_nascimento < CURRENT_DATE)` na tabela alunos**

```
Tipo: [QUAL TIPO?]
Explicação: [POR QUE?]
```

**e) `UNIQUE(aluno_id, turma_id)` na tabela matriculas**

```
Tipo: [QUAL TIPO?]
Explicação: [POR QUE? O QUE ISSO IMPEDE?]
```

### 4.2 - O que Aconteceria?

Para cada operação abaixo, diga se seria **PERMITIDA** ou **BLOQUEADA**, e explique por quê:

**a) Inserir aluno sem nome:**

```sql
INSERT INTO alunos (data_nascimento, email)
VALUES ('2005-05-15', 'joao@email.com');
```

```
Resultado: [PERMITIDA ou BLOQUEADA?]
Por quê:
```

**b) Inserir aluno com data de nascimento no futuro:**

```sql
INSERT INTO alunos (nome, data_nascimento, email)
VALUES ('João', '2030-01-01', 'joao@email.com');
```

```
Resultado: [PERMITIDA ou BLOQUEADA?]
Por quê:
```

**c) Criar matrícula para aluno inexistente:**

```sql
INSERT INTO matriculas (aluno_id, turma_id, data_matricula)
VALUES (999, 1, '2024-01-15');
```

```
Resultado: [PERMITIDA ou BLOQUEADA?]
Por quê:
```

**d) Matricular o mesmo aluno na mesma turma duas vezes:**

```sql
-- Primeira vez
INSERT INTO matriculas (aluno_id, turma_id, data_matricula)
VALUES (1, 1, '2024-01-15');

-- Segunda vez
INSERT INTO matriculas (aluno_id, turma_id, data_matricula)
VALUES (1, 1, '2024-01-20');
```

```
Resultado: [PERMITIDA ou BLOQUEADA?]
Por quê:
```

---

## 🎯 Exercício 5: Modelagem Completa

Você foi contratado para criar o banco de dados de uma **locadora de filmes**. O sistema precisa guardar:

- **Filmes**: título, diretor, ano, gênero, duração
- **Clientes**: nome, CPF, email, telefone, endereço
- **Locações**: qual cliente alugou qual filme, data de retirada, data de devolução

### Sua Tarefa:

**a) Desenhe as 3 tabelas com:**

- Todos os atributos necessários
- Chaves primárias
- Chaves estrangeiras
- Pelo menos 2 restrições de domínio (CHECK)

```
[DESENHE AS TABELAS AQUI]
```

**b) Identifique os relacionamentos:**

```
Relacionamento entre CLIENTES e LOCACOES:
Tipo: [1:1, 1:N ou N:M?]
Justificativa:

Relacionamento entre FILMES e LOCACOES:
Tipo: [1:1, 1:N ou N:M?]
Justificativa:
```

**c) Liste 3 regras de integridade importantes para este sistema:**

```
1. [REGRA 1 E TIPO DE INTEGRIDADE]

2. [REGRA 2 E TIPO DE INTEGRIDADE]

3. [REGRA 3 E TIPO DE INTEGRIDADE]
```

---

## 🧠 Perguntas de Reflexão

### Reflexão 1: Independência de Dados

O modelo relacional promove a "independência de dados" - a separação entre a visão lógica (como você consulta) e a implementação física (como está armazenado).

**Reflita:**

- Por que essa separação é importante para aplicações de longo prazo?
- Imagine que você criou um sistema que usa PostgreSQL. Depois de 5 anos, você quer otimizar o banco adicionando índices, mudando particionamento, etc. Como a independência de dados ajuda nessa situação?
- O que aconteceria se não houvesse essa separação? (dica: pense nos bancos de dados antes de 1970)

**Sua reflexão:**

```
[ESCREVA SEUS PENSAMENTOS AQUI - MÍNIMO 5 LINHAS]
```

---

### Reflexão 2: Normalização vs. Desempenho

O modelo relacional incentiva que você não repita dados (normalização). Por exemplo:

**Forma não normalizada (dados repetidos):**

```
PEDIDOS
┌────┬──────────────┬─────────────────┬──────────┐
│ ID │ CLIENTE_NOME │ CLIENTE_EMAIL   │ PRODUTO  │
├────┼──────────────┼─────────────────┼──────────┤
│ 1  │ João         │ joao@email.com  │ Notebook │
│ 2  │ João         │ joao@email.com  │ Mouse    │  ← Nome e email repetidos!
│ 3  │ João         │ joao@email.com  │ Teclado  │
└────┴──────────────┴─────────────────┴──────────┘
```

**Forma normalizada (sem repetição):**

```
CLIENTES                  PEDIDOS
┌────┬──────┬──────────┐  ┌────┬────────────┬──────────┐
│ ID │ NOME │ EMAIL    │  │ ID │ CLIENTE_ID │ PRODUTO  │
├────┼──────┼──────────┤  ├────┼────────────┼──────────┤
│ 1  │ João │ joao@... │  │ 1  │ 1          │ Notebook │
└────┴──────┴──────────┘  │ 2  │ 1          │ Mouse    │
                          │ 3  │ 1          │ Teclado  │
                          └────┴────────────┴──────────┘
```

**Mas:** A forma normalizada requer JOIN para ver nome do cliente + produto, o que pode ser mais lento.

**Reflita:**

- Por que repetir dados (forma não normalizada) pode ser problemático?
- Em que situações você consideraria repetir dados propositalmente para ganhar velocidade?
- Como você balancearia integridade de dados vs. desempenho em um sistema real?

**Sua reflexão:**

```
[ESCREVA SEUS PENSAMENTOS AQUI - MÍNIMO 5 LINHAS]
```

---

### Reflexão 3: Chaves Naturais vs. Chaves Artificiais

Existem dois tipos de chaves primárias:

**Chave Natural:** Usa dado real (CPF, email, ISBN de livro)
**Chave Artificial:** Usa ID inventado (1, 2, 3, ...)

Exemplo:

```
Opção A - Chave Natural (CPF):
CLIENTES
┌──────────────┬──────┬──────────┐
│ CPF (PK)     │ NOME │ EMAIL    │
├──────────────┼──────┼──────────┤
│ 111.111.111-11│ João │ joao@... │
└──────────────┴──────┴──────────┘

Opção B - Chave Artificial (ID):
CLIENTES
┌────┬──────────────┬──────┬──────────┐
│ ID │ CPF          │ NOME │ EMAIL    │
├────┼──────────────┼──────┼──────────┤
│ 1  │ 111.111.111-11│ João │ joao@... │
└────┴──────────────┴──────┴──────────┘
```

**Reflita:**

- Quais são as vantagens de usar uma chave natural (CPF)?
- Quais são as vantagens de usar uma chave artificial (ID numérico)?
- O que aconteceria se uma pessoa precisasse mudar o CPF (casos raros mas existem)? Qual abordagem seria mais fácil de lidar com isso?
- Qual você preferiria usar em seus projetos? Por quê?

**Sua reflexão:**

```
[ESCREVA SEUS PENSAMENTOS AQUI - MÍNIMO 5 LINHAS]
```

---

## 📤 Como Enviar Suas Respostas

1. Copie este arquivo
2. Preencha **todas** as seções
3. Envie para análise

**Critérios de avaliação:**

- ✅ Compreensão dos conceitos fundamentais (relação, tupla, atributo, domínio)
- ✅ Domínio de chaves (PK, FK, candidatas)
- ✅ Identificação correta de tipos de relacionamentos
- ✅ Entendimento de integridade de dados
- ✅ Profundidade nas reflexões
- ✅ Capacidade de aplicar conceitos em situações práticas

---

## ⏱️ Tempo Estimado

- Exercícios 1-5: 60-80 minutos
- Reflexões: 30-40 minutos
- **Total: 90-120 minutos**

Este é um conteúdo denso e fundamental. Dedique tempo de qualidade!

---

## 🎯 Próximos Passos

Após enviar suas respostas, você receberá:

- 📊 Análise detalhada do seu desempenho
- ⚠️ Conceitos que precisam ser reforçados
- ✅ Pontos fortes identificados
- 🎯 Recomendações personalizadas para próxima aula

O modelo relacional é a BASE de tudo em PostgreSQL. Dominar esses conceitos é essencial!

Boa sorte! 🚀


