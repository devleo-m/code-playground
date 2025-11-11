# 🐘 Curso de PostgreSQL - Do Iniciante ao Avançado

Bem-vindo ao curso completo de PostgreSQL! Este curso foi estruturado seguindo uma metodologia rigorosa de ensino, focada em compreensão profunda, não em memorização superficial.

---

## 🎯 Sobre Este Curso

### Metodologia de Ensino

Este curso segue o **Ciclo de 4 Etapas**:

```
┌─────────────────────────────────────────────────────────┐
│  1. AULA PRINCIPAL                                      │
│     └─ Conteúdo técnico completo e detalhado           │
├─────────────────────────────────────────────────────────┤
│  2. AULA SIMPLIFICADA                                   │
│     └─ Mesmos conceitos com analogias do cotidiano     │
├─────────────────────────────────────────────────────────┤
│  3. EXERCÍCIOS E REFLEXÃO                               │
│     └─ Práticas + perguntas que exigem pensamento      │
├─────────────────────────────────────────────────────────┤
│  4. ANÁLISE E FEEDBACK                                  │
│     └─ Avaliação crítica do seu desempenho             │
│     └─ Identificação de pontos fortes e fracos         │
│     └─ Recomendações personalizadas                    │
└─────────────────────────────────────────────────────────┘
```

### Perfil do Aluno

Este curso é desenhado para **iniciantes absolutos** em programação e banco de dados. Cada conceito é explicado de forma clara, com exemplos práticos e analogias do dia a dia.

---

## 📚 Estrutura do Curso

### 🟢 **Aula 1: Introdução ao PostgreSQL**

📂 Pasta: [`01-aula/`](./01-aula/)

**Tópicos cobertos:**

- O que é PostgreSQL e sua história
- Bancos de dados relacionais
- Propriedades ACID
- Benefícios e limitações de RDBMS
- PostgreSQL vs NoSQL
- PostgreSQL vs outros bancos relacionais (MySQL, Oracle, SQL Server)
- Recursos especiais do PostgreSQL

**Tempo estimado:** 2-3 horas

**Status:** ✅ Completa

---

### 🟢 **Aula 2: Conceitos de SGBD Relacional**

📂 Pasta: [`02-aula/`](./02-aula/)

**Tópicos cobertos:**

- O que é um SGBD (Sistema de Gerenciamento de Banco de Dados)
- História de E.F. Codd e o modelo relacional (1970)
- Estrutura do modelo relacional:
  - Relação (tabela)
  - Tupla (linha)
  - Atributo (coluna)
  - Domínio
- Chaves:
  - Superchave
  - Chave candidata
  - Chave primária (PK)
  - Chave estrangeira (FK)
- Tipos de relacionamentos:
  - Um para Um (1:1)
  - Um para Muitos (1:N)
  - Muitos para Muitos (N:M)
- Integridade de dados:
  - Integridade de Entidade
  - Integridade Referencial
  - Integridade de Domínio
  - Integridade de Negócio
- Independência de dados
- Operações relacionais

**Tempo estimado:** 3-4 horas

**Status:** ✅ Completa

**⚠️ Aula Fundamental:** O modelo relacional é a base de tudo em PostgreSQL!

---

### 🟢 **Aula 3: Object Model in PostgreSQL**

📂 Pasta: [`03-aula/`](./03-aula/)

**Tópicos cobertos:**

- ORDBMS: PostgreSQL como híbrido (relacional + orientado a objetos)
- Recursos orientados a objetos:
  - Tipos de dados customizados
  - Herança de tabelas
  - Polimorfismo
- Hierarquia de objetos completa:
  - Server → Database → Schema → Table → Column → Row
- **Databases**: Criação, conexão, gerenciamento, isolamento
- **Schemas**: Organização, namespacing, multi-tenancy
- **Tables**: Criação, modificação, tipos (permanentes, temporárias, unlogged)
- **Columns**: Definição, tipos de dados, constraints, modificação
- **Rows**: Inserção, atualização, deleção
- **Queries no PostgreSQL**:
  - SELECT (básico e avançado)
  - WHERE (filtros e operadores)
  - ORDER BY (ordenação)
  - LIMIT e OFFSET (paginação)
  - Funções de agregação (COUNT, SUM, AVG, MIN, MAX)
  - GROUP BY (agrupamento)
  - HAVING (filtro pós-agregação)
  - JOINs (INNER, LEFT, RIGHT, FULL OUTER)
  - Subconsultas (subqueries)
- **Tipos de Dados**:
  - Numéricos (SMALLINT, INTEGER, BIGINT, DECIMAL, SERIAL)
  - Caracteres (CHAR, VARCHAR, TEXT)
  - Data e Hora (DATE, TIME, TIMESTAMP, INTERVAL)
  - Boolean
  - Enum (tipos enumerados)
  - Arrays
  - JSON e JSONB (dados semi-estruturados)
  - Geométricos
  - UUID

**Tempo estimado:** 5-7 horas

**Status:** ✅ Completa

**⚠️ Última Aula Antes da Avaliação:** Prepare-se para aplicar TUDO que aprendeu!

---

### 🔵 **Próximas Aulas (Em Desenvolvimento)**

#### Sugestões de Tópicos:

**Módulo 1: Fundamentos**

- [ ] Aula 4: Instalação e Configuração do PostgreSQL
- [ ] Aula 5: SQL Avançado - DDL (Data Definition Language)
- [ ] Aula 6: SQL Avançado - DML (Data Manipulation Language)

**Módulo 2: Design e Modelagem**

- [ ] Aula 7: Normalização de Banco de Dados (1FN, 2FN, 3FN, BCNF)
- [ ] Aula 8: Modelagem de Dados Avançada
- [ ] Aula 9: Constraints e Validações

**Módulo 3: Consultas**

- [ ] Aula 10: SELECT Avançado
- [ ] Aula 11: JOINs (INNER, LEFT, RIGHT, FULL, CROSS)
- [ ] Aula 12: Subconsultas e CTEs (Common Table Expressions)
- [ ] Aula 13: Funções de Agregação e GROUP BY

**Módulo 4: Performance**

- [ ] Aula 14: Índices (B-tree, Hash, GiST, GIN, BRIN)
- [ ] Aula 15: EXPLAIN e Análise de Performance
- [ ] Aula 16: Otimização de Consultas

**Módulo 5: Recursos Avançados**

- [ ] Aula 17: Transações e Controle de Concorrência
- [ ] Aula 18: Views, Views Materializadas
- [ ] Aula 19: Triggers e Stored Procedures
- [ ] Aula 20: Trabalhando com JSON/JSONB
- [ ] Aula 21: Full-Text Search
- [ ] Aula 22: PostGIS - Dados Geográficos

**Módulo 6: Administração**

- [ ] Aula 23: Backup e Recuperação
- [ ] Aula 24: Segurança e Controle de Acesso
- [ ] Aula 25: Replicação e Alta Disponibilidade

---

## 🎓 Como Usar Este Curso

### Para Cada Aula:

#### **Passo 1: Leia a Aula Principal** 📖

Cada aula tem um arquivo principal (ex: `aula-01-introducao-postgresql.md`) com conteúdo técnico completo.

**Dica:** Faça anotações dos conceitos-chave!

---

#### **Passo 2: Leia a Versão Simplificada** 🎯

Reforce o aprendizado com a versão simplificada (ex: `aula-01-simplificada.md`) que usa analogias e exemplos do cotidiano.

**Benefício:** Consolida conceitos abstratos de forma intuitiva.

---

#### **Passo 3: Faça os Exercícios** ✏️

Complete todos os exercícios (ex: `aula-01-exercicios.md`) **sem consultar** as aulas.

**Importante:** Use suas próprias palavras! Não copie e cole.

---

#### **Passo 4: Envie para Análise** 📤

Envie suas respostas completas para receber feedback personalizado.

---

## ⚠️ Regras de Ouro do Curso

### ❌ NÃO Faça:

- Pular exercícios
- Copiar e colar respostas
- Avançar sem dominar a aula anterior
- Decorar sem entender
- Apressar o processo

### ✅ FAÇA:

- Dedique tempo de qualidade
- Use suas próprias palavras
- Reflita profundamente sobre as perguntas
- Seja honesto sobre dúvidas
- Conecte conceitos com experiências reais
- Desenhe diagramas e tabelas
- Explique conceitos em voz alta

---

## 📊 Sobre o Feedback

O feedback que você receberá será:

### 🎯 **Focado em Melhoria**

- Destaque nos **pontos fracos**, não nos fortes
- Crítica construtiva e honesta
- Zero elogios vazios

### 📈 **Analítico**

- Identificação de lacunas conceituais
- Avaliação da profundidade de compreensão
- Análise da capacidade de aplicação prática

### 🎓 **Orientado a Ação**

- Recomendações específicas do que estudar
- Sugestões de revisão
- Indicação de conceitos que precisam ser reforçados

---

## 🗺️ Trilha de Aprendizado

```
┌─────────────────────────────────────────────────────────┐
│  FUNDAMENTOS                                            │
│  ├─ Aula 1: Introdução ao PostgreSQL         ✅         │
│  └─ Aula 2: Modelo Relacional                ✅         │
├─────────────────────────────────────────────────────────┤
│  PRÁTICA BÁSICA                                         │
│  ├─ Instalação e Configuração                          │
│  ├─ Tipos de Dados                                     │
│  └─ SQL Básico (DDL + DML)                             │
├─────────────────────────────────────────────────────────┤
│  DESIGN                                                 │
│  ├─ Normalização                                       │
│  ├─ Modelagem Avançada                                 │
│  └─ Constraints                                        │
├─────────────────────────────────────────────────────────┤
│  CONSULTAS AVANÇADAS                                    │
│  ├─ SELECT Complexo                                    │
│  ├─ JOINs                                              │
│  └─ Subconsultas e CTEs                                │
├─────────────────────────────────────────────────────────┤
│  PERFORMANCE                                            │
│  ├─ Índices                                            │
│  ├─ EXPLAIN                                            │
│  └─ Otimização                                         │
├─────────────────────────────────────────────────────────┤
│  RECURSOS AVANÇADOS                                     │
│  ├─ Transações                                         │
│  ├─ Views e Triggers                                   │
│  ├─ JSON/JSONB                                         │
│  └─ PostGIS                                            │
└─────────────────────────────────────────────────────────┘
```

---

## 💡 Por Que Este Curso é Diferente?

### 1. 🎯 **Foco em Entendimento Profundo**

Não apenas "como fazer", mas "por que funciona assim".

### 2. 🧠 **Perguntas de Reflexão**

Você é forçado a **pensar**, não apenas memorizar.

### 3. 📊 **Feedback Rigoroso**

Análise honesta do seu desempenho, sem elogios vazios.

### 4. 🏗️ **Base Sólida**

Começamos pelos fundamentos teóricos (modelo relacional) antes de praticar.

### 5. 🔄 **Metodologia Comprovada**

Aula técnica → Simplificada → Exercícios → Feedback

---

## 📈 Progresso Recomendado

### Iniciante Absoluto (Você está aqui!)

- **Meta:** Dominar Aulas 1-6
- **Tempo:** 2-3 meses (3-4 horas/semana)
- **Resultado:** Base sólida em PostgreSQL

### Intermediário

- **Meta:** Completar Aulas 7-16
- **Tempo:** 2-3 meses
- **Resultado:** Capaz de desenvolver sistemas reais

### Avançado

- **Meta:** Completar Aulas 17-25
- **Tempo:** 3-4 meses
- **Resultado:** Expertise em PostgreSQL

---

## ⏱️ Estimativas de Tempo

| Aula      | Tópico                   | Leitura      | Exercícios  | Total      |
| --------- | ------------------------ | ------------ | ----------- | ---------- |
| 1         | Introdução ao PostgreSQL | 60-90 min    | 50-70 min   | 2-3h       |
| 2         | Modelo Relacional        | 90-120 min   | 90-120 min  | 3-4h       |
| 3         | Object Model             | 130-180 min  | 180-240 min | 5-7h       |
| **TOTAL** | **3 Aulas Fundamentais** | **4,5-6,5h** | **5,5-7h**  | **10-14h** |

**Dica:** Não tente fazer tudo de uma vez. Divida em múltiplas sessões de estudo!

**Recomendação para Aula 3:** Divida em 3-4 sessões (é a mais extensa e prática)!

---

## 🎯 Objetivo Final do Curso

Ao completar este curso, você será capaz de:

- ✅ Entender profundamente como bancos de dados relacionais funcionam
- ✅ Modelar bancos de dados robustos e bem estruturados
- ✅ Escrever SQL complexo e eficiente
- ✅ Otimizar performance de consultas
- ✅ Usar recursos avançados do PostgreSQL (JSON, GIS, full-text search)
- ✅ Administrar bancos de dados em produção
- ✅ Tomar decisões arquiteturais fundamentadas

---

## 📞 Dúvidas e Suporte

### Tem dúvida sobre algum conceito?

- Releia a seção específica
- Consulte a versão simplificada
- Tente explicar com suas próprias palavras
- **Pergunte!** Dúvidas demonstram pensamento crítico

### Não entendeu o feedback?

- Peça esclarecimentos específicos
- Solicite exemplos adicionais
- Pergunte "por que" sua resposta estava incorreta

---

## 🚀 Vamos Começar!

**Pronto para começar sua jornada em PostgreSQL?**

👉 Comece pela **[Aula 1: Introdução ao PostgreSQL](./01-aula/)**

Lembre-se: **Qualidade > Velocidade**

Cada hora investida agora em fundamentos sólidos vai economizar dezenas de horas de confusão no futuro.

---

## 🏆 Filosofia do Curso

> "Ensinar não é transferir conhecimento, mas criar as possibilidades para a sua própria produção ou a sua construção."  
> — Paulo Freire

Este curso não vai te dar respostas prontas. Vai te ensinar a **pensar** como um profissional de banco de dados.

---

---

## 📝 Preparação para Avaliação

Você completou as **3 aulas fundamentais** do curso! Agora é hora de consolidar seu conhecimento.

### ✅ O que você aprendeu até agora:

```
┌──────────────────────────────────────────────────────────┐
│  AULA 1: Visão Geral do PostgreSQL                      │
│  - O que é PostgreSQL (ORDBMS)                          │
│  - ACID e integridade                                   │
│  - Quando usar PostgreSQL vs NoSQL                      │
│  - Comparação com outros bancos                         │
└──────────────────────────────────────────────────────────┘
                         ↓
┌──────────────────────────────────────────────────────────┐
│  AULA 2: Modelo Relacional (E.F. Codd)                  │
│  - Relação, tupla, atributo, domínio                    │
│  - Chaves (primárias, estrangeiras, candidatas)         │
│  - Relacionamentos (1:1, 1:N, N:M)                      │
│  - Integridade de dados                                 │
└──────────────────────────────────────────────────────────┘
                         ↓
┌──────────────────────────────────────────────────────────┐
│  AULA 3: Implementação Prática                          │
│  - Hierarquia: Database → Schema → Table → Column → Row│
│  - Tipos de dados (numéricos, texto, JSON, etc)        │
│  - Queries SQL (SELECT, JOIN, GROUP BY, etc)           │
│  - Modelagem de sistemas reais                          │
└──────────────────────────────────────────────────────────┘
```

### 🎯 Próximo Passo: Avaliação

Antes de avançar para aulas mais avançadas, você será **avaliado** sobre:

1. **Compreensão Conceitual** (Aulas 1-2)

   - Entendimento do modelo relacional
   - Conhecimento de chaves e relacionamentos
   - Compreensão de integridade

2. **Aplicação Prática** (Aula 3)

   - Modelagem de sistema completo
   - Escolha correta de tipos de dados
   - Escrita de SQL (SELECT, INSERT, UPDATE, JOIN)

3. **Pensamento Crítico**
   - Justificativa de decisões de design
   - Trade-offs (normalização vs performance, etc)
   - Resolução de problemas reais

### 📋 Checklist Final

Antes de submeter os exercícios da Aula 3, verifique:

**Conceitos (Aula 1-2):**

- [ ] Sei explicar ACID
- [ ] Entendo diferença entre PK e FK
- [ ] Sei identificar tipos de relacionamentos (1:1, 1:N, N:M)
- [ ] Compreendo os 4 tipos de integridade

**Prática (Aula 3):**

- [ ] Sei criar databases e schemas
- [ ] Consigo criar tabelas com tipos corretos
- [ ] Escrevo SELECT com WHERE e ORDER BY
- [ ] Faço JOINs entre tabelas
- [ ] Uso GROUP BY e funções de agregação
- [ ] Sei quando usar JSONB

**Se marcou menos de 10:** Revise antes de enviar!
**Se marcou todos:** Você está pronto! 🚀

---

**Última atualização:** Aula 3 completa ✅ (Aulas fundamentais concluídas!)

**Próximo passo:** Complete os exercícios da Aula 3 e envie para avaliação!

Bons estudos! 💪🐘
