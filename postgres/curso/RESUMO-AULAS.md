# 📚 Resumo Rápido das Aulas do Curso PostgreSQL

Este arquivo contém um resumo executivo de cada aula para consulta rápida.

---

## 🟢 **Aula 1: Introdução ao PostgreSQL**

### 📂 Localização: `01-aula/`

### 🎯 Objetivo

Entender o que é PostgreSQL, quando usar, e suas vantagens comparado a outros bancos de dados.

### 📋 Tópicos Principais

1. **O que é PostgreSQL**

   - ORDBMS (Object-Relational DBMS)
   - Open-source, gratuito, poderoso
   - História: Universidade da Califórnia, Berkeley (década de 1980)

2. **Bancos de Dados Relacionais**

   - Organização em tabelas (linhas e colunas)
   - Relacionamentos entre tabelas
   - Diferença entre dados estruturados e não-estruturados

3. **ACID**

   - **A**tomicidade: Tudo ou nada
   - **C**onsistência: Regras sempre respeitadas
   - **I**solamento: Transações não interferem
   - **D**urabilidade: Dados persistem após confirmação

4. **RDBMS: Benefícios**

   - Integridade de dados (ACID)
   - SQL poderoso
   - Relacionamentos fortes (foreign keys)
   - Escalabilidade vertical

5. **RDBMS: Limitações**

   - Escalabilidade horizontal difícil
   - Rigidez de schema
   - Não ideal para dados não-estruturados

6. **PostgreSQL vs NoSQL**

   - PostgreSQL: Estruturado, ACID, relacional
   - NoSQL: Flexível, escalável horizontalmente, eventual consistency

7. **PostgreSQL vs Outros Bancos**

   - vs MySQL: Mais features, melhor conformidade SQL
   - vs Oracle: Gratuito vs caro, ambos poderosos
   - vs SQL Server: Multi-plataforma vs Windows-centric

8. **Recursos Especiais do PostgreSQL**
   - Extensões (PostGIS, pg_trgm, uuid-ossp)
   - JSON/JSONB
   - Tipos avançados (arrays, hstore)
   - Índices avançados (GiST, GIN, BRIN)
   - Full-text search nativo

### ⏱️ Tempo: 2-3 horas

---

## 🟢 **Aula 2: Conceitos de SGBD Relacional**

### 📂 Localização: `02-aula/`

### 🎯 Objetivo

Dominar os conceitos fundamentais do modelo relacional proposto por E.F. Codd (1970).

### 📋 Tópicos Principais

1. **O que é SGBD**

   - Software intermediário entre usuário e dados
   - Gerencia armazenamento, acesso, concorrência, integridade

2. **E.F. Codd e o Modelo Relacional (1970)**

   - Problema que resolveu: bancos hierárquicos/rede eram complexos
   - Solução: Organizar dados em tabelas (relações)
   - Separação lógica vs física

3. **Estrutura do Modelo**

   - **Relação**: Tabela completa
   - **Tupla**: Linha/registro individual
   - **Atributo**: Coluna/campo
   - **Domínio**: Conjunto de valores válidos

4. **Chaves**

   - **Superchave**: Qualquer conjunto que identifica unicamente
   - **Chave Candidata**: Superchave mínima
   - **Chave Primária (PK)**: Candidata escolhida como identificador oficial
   - **Chave Estrangeira (FK)**: Referência a PK de outra tabela

5. **Tipos de Relacionamentos**

   - **1:1** (Um para Um): Ex: Pessoa ↔ Passaporte
   - **1:N** (Um para Muitos): Ex: Cliente → Pedidos
   - **N:M** (Muitos para Muitos): Ex: Alunos ↔ Turmas (requer tabela associativa)

6. **Integridade de Dados**

   - **Integridade de Entidade**: PK não pode ser NULL e deve ser única
   - **Integridade Referencial**: FK deve referenciar PK existente
   - **Integridade de Domínio**: Valores dentro do domínio definido
   - **Integridade de Negócio**: Regras customizadas da aplicação

7. **Independência de Dados**

   - Nível lógico separado do nível físico
   - Mudanças físicas não afetam aplicações

8. **Operações Relacionais**
   - Seleção (σ): Filtrar linhas
   - Projeção (π): Selecionar colunas
   - Junção (⋈): Combinar tabelas
   - União (∪): Combinar resultados
   - Diferença (−): Subtrair resultados

### ⏱️ Tempo: 3-4 horas

---

## 🟢 **Aula 3: Object Model in PostgreSQL**

### 📂 Localização: `03-aula/`

### 🎯 Objetivo

Entender como PostgreSQL implementa o modelo relacional na prática, incluindo recursos orientados a objetos.

### 📋 Tópicos Principais

1. **ORDBMS: Híbrido**

   - Combina relacional (RDBMS) + orientado a objetos (OODBMS)
   - Tipos customizados, herança de tabelas, polimorfismo

2. **Hierarquia de Objetos**

   ```
   Server → Database → Schema → Table → Column → Row
   ```

3. **Databases**

   - Coleção isolada de schemas
   - Múltiplos databases por servidor
   - Comandos: `CREATE DATABASE`, `DROP DATABASE`, `\c database_name`

4. **Schemas**

   - Namespace dentro de database
   - Organização lógica de objetos
   - Schema padrão: `public`
   - Comandos: `CREATE SCHEMA`, `DROP SCHEMA`, `SET search_path`

5. **Tables**

   - Coleção de linhas com estrutura definida
   - Tipos: permanentes, temporárias, unlogged
   - Comandos: `CREATE TABLE`, `ALTER TABLE`, `DROP TABLE`

6. **Columns**

   - Definem atributos (tipo, constraints)
   - Modificação: `ALTER TABLE ... ADD/DROP/RENAME COLUMN`

7. **Rows**

   - Registros individuais
   - Operações: `INSERT`, `UPDATE`, `DELETE`

8. **Queries no PostgreSQL**

   - **SELECT**: Básico e avançado
   - **WHERE**: Filtros (=, >, <, LIKE, IN, IS NULL, BETWEEN)
   - **ORDER BY**: Ordenação (ASC, DESC)
   - **LIMIT/OFFSET**: Paginação
   - **Agregações**: COUNT, SUM, AVG, MIN, MAX
   - **GROUP BY**: Agrupamento
   - **HAVING**: Filtro pós-agregação
   - **JOINs**: INNER, LEFT, RIGHT, FULL OUTER
   - **Subqueries**: Consultas aninhadas

9. **Tipos de Dados**
   - **Numéricos**: SMALLINT, INTEGER, BIGINT, DECIMAL, SERIAL
   - **Texto**: CHAR(n), VARCHAR(n), TEXT
   - **Data/Hora**: DATE, TIME, TIMESTAMP, INTERVAL
   - **Boolean**: TRUE/FALSE/NULL
   - **ENUM**: Valores pré-definidos
   - **Arrays**: Listas de valores
   - **JSON/JSONB**: Dados semi-estruturados
   - **Geométricos**: POINT, LINE, CIRCLE, etc.
   - **UUID**: Identificadores únicos universais

### ⏱️ Tempo: 5-7 horas

---

## 🟢 **Aula 4: Relational Model (Aprofundamento)**

### 📂 Localização: `04-aula/`

### 🎯 Objetivo

Aprofundar no modelo relacional com foco em domains, constraints, null values e implementação prática no PostgreSQL.

### 📋 Tópicos Principais

1. **Revisão do Modelo Relacional**

   - Fundamentos de E.F. Codd (1970)
   - Organização em relações (tabelas)

2. **Domains (Domínios Customizados)**

   - Tipos de dados customizados
   - Constraints e validações reutilizáveis
   - `CREATE DOMAIN`, `ALTER DOMAIN`, `DROP DOMAIN`

3. **Attributes (Atributos)**

   - Colunas que definem propriedades
   - Domínio de cada atributo
   - Papel na integridade de dados

4. **Tuples (Tuplas)**

   - Registros individuais
   - Conjunto ordenado de valores
   - Operações em tuplas

5. **Relations (Relações)**

   - Estrutura: schema + tuplas
   - Integrity constraints
   - Operações relacionais

6. **Constraints (Restrições)**

   - **PRIMARY KEY**: Identificador único e obrigatório
   - **FOREIGN KEY**: Integridade referencial
   - **UNIQUE**: Valores únicos
   - **CHECK**: Validações customizadas
   - **NOT NULL**: Valor obrigatório
   - **EXCLUSION**: Previne conflitos entre linhas

7. **Null Values (Valores Nulos)**
   - Representa ausência de valor
   - Diferente de zero, string vazia, false
   - Operações com NULL retornam NULL
   - Comparações com NULL: `IS NULL`, `IS NOT NULL`
   - Função `COALESCE`: Primeiro valor não-nulo

### ⏱️ Tempo: 4-6 horas

---

## 📊 Resumo Total do Curso (Aulas 1-4)

| Aula      | Foco                        | Carga Horária |
| --------- | --------------------------- | ------------- |
| 1         | Visão geral PostgreSQL      | 2-3h          |
| 2         | Teoria do modelo relacional | 3-4h          |
| 3         | Implementação prática       | 5-7h          |
| 4         | Aprofundamento relacional   | 4-6h          |
| **TOTAL** | **Fundamentos Completos**   | **14-20h**    |

---

## 🎯 Sequência de Aprendizado

```
Aula 1: O QUE é PostgreSQL?
   ↓
Aula 2: COMO funciona o modelo relacional? (teoria)
   ↓
Aula 3: COMO usar PostgreSQL? (prática: queries, tipos)
   ↓
Aula 4: COMO garantir qualidade? (domains, constraints, nulls)
```

---

## 📝 Próximas Aulas Sugeridas

- Aula 5: Normalização de Banco de Dados (1FN, 2FN, 3FN, BCNF)
- Aula 6: Índices e Performance
- Aula 7: Transações e Controle de Concorrência
- Aula 8: Views, Functions e Triggers
- Aula 9: JSON/JSONB Avançado
- Aula 10: PostGIS - Dados Geográficos

---

**Última atualização:** Aula 4 completa ✅

**Use este arquivo para:** Revisão rápida antes de exercícios/provas ou para lembrar o que cada aula aborda!
