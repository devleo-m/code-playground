# Aula 3: Data Definition Language (DDL)

Bem-vindo à terceira aula do curso de SQL! Esta aula apresenta os comandos DDL (Data Definition Language), que permitem criar, modificar e gerenciar a estrutura do banco de dados.

## 📚 Estrutura da Aula

Esta aula segue o ciclo de ensino de 5 etapas. Siga a ordem indicada:

### 1. Aula Principal
**Arquivo**: `01-aula-principal.md`

Conteúdo técnico completo sobre:
- O que é DDL e sua importância
- CREATE TABLE: Criar tabelas e definir estrutura
- ALTER TABLE: Modificar estrutura de tabelas existentes
- DROP TABLE: Remover tabelas completamente
- TRUNCATE TABLE: Limpar dados mantendo estrutura
- Constraints (restrições): PRIMARY KEY, FOREIGN KEY, UNIQUE, NOT NULL, CHECK
- Índices: CREATE INDEX e DROP INDEX

**Tempo estimado**: 60-90 minutos

---

### 2. Aula Simplificada
**Arquivo**: `02-aula-simplificada.md`

Mesmo conteúdo, mas explicado com:
- Analogias do dia a dia (tabelas como prateleiras, colunas como gavetas)
- Metáforas visuais
- Exemplos práticos
- Comparações com conceitos conhecidos

**Tempo estimado**: 45-60 minutos

---

### 3. Exercícios e Reflexão
**Arquivo**: `03-exercicios-reflexao.md`

Conjunto de exercícios práticos incluindo:
- Exercícios de criação de tabelas (CREATE TABLE)
- Exercícios de modificação de estrutura (ALTER TABLE)
- Exercícios de remoção de tabelas (DROP TABLE)
- Exercícios de limpeza de dados (TRUNCATE TABLE)
- Problemas que exigem raciocínio sobre estrutura de dados
- Perguntas de reflexão sobre eficiência, impacto e boas práticas

**Tempo estimado**: 90-120 minutos

**⚠️ IMPORTANTE**: Complete todos os exercícios e responda as perguntas de reflexão antes de prosseguir.

---

### 4. Performance, Boas Práticas e Otimização
**Arquivo**: `04-performance-boas-praticas.md`

Conteúdo sobre:
- Impacto de operações DDL no desempenho
- Quando e como usar ALTER TABLE em produção
- Estratégias para modificar tabelas grandes
- Boas práticas de nomenclatura de tabelas e colunas
- Normalização e desnormalização
- Índices: quando criar, quando não criar
- Constraints: importância para integridade de dados
- Backup antes de operações DDL destrutivas
- Versionamento de schema
- Migrations e controle de versão

**Tempo estimado**: 45-60 minutos

---

### 5. Análise e Feedback
**Status**: Aguardando resposta do aluno

Após completar os exercícios, envie suas respostas para análise. O tutor fornecerá feedback construtivo sobre:
- Correção sintática dos comandos DDL
- Eficiência e impacto das operações
- Compreensão dos conceitos de estrutura de dados
- Áreas que necessitam de melhoria

---

## 🎯 Objetivos de Aprendizado

Ao final desta aula, você deve ser capaz de:

- [ ] Entender o que é DDL e sua diferença em relação a DML
- [ ] Criar tabelas usando CREATE TABLE com todas as constraints necessárias
- [ ] Modificar estrutura de tabelas existentes usando ALTER TABLE
- [ ] Remover tabelas usando DROP TABLE
- [ ] Limpar dados de tabelas usando TRUNCATE TABLE
- [ ] Compreender e aplicar constraints (PRIMARY KEY, FOREIGN KEY, UNIQUE, NOT NULL, CHECK)
- [ ] Criar e remover índices para melhorar performance
- [ ] Entender o impacto de operações DDL no banco de dados
- [ ] Aplicar boas práticas ao criar e modificar estruturas de tabelas

## 📖 Pré-requisitos

- Conclusão da Aula 1 (Introdução ao SQL)
- Conclusão da Aula 2 (Sintaxe Básica de SQL)
- Banco de dados `biblioteca.db` criado e funcionando
- Acesso ao SQLite (via CLI ou ferramenta visual)
- Compreensão básica de tabelas, linhas, colunas e relacionamentos

## 🔄 Revisão Rápida das Aulas Anteriores

Antes de começar, vamos revisar os conceitos-chave das aulas anteriores:

### Da Aula 1:
- **SQL** é composto por DDL, DML, DCL e Queries
- **DDL** (Data Definition Language) cria e modifica estruturas
- **Bancos relacionais** organizam dados em tabelas com relacionamentos

### Da Aula 2:
- **SELECT, INSERT, UPDATE, DELETE** são comandos DML para manipular dados
- **Tipos de dados**: INTEGER, TEXT, DATE, etc.
- **Operadores**: comparação, lógicos, aritméticos
- **WHERE** é crucial para filtrar dados

Se você não se lembra desses conceitos, revise as aulas anteriores antes de prosseguir.

## 🚀 Como Usar Esta Aula

1. **Leia a Aula Principal** (`01-aula-principal.md`)
   - Não se preocupe em decorar tudo
   - Foque em entender os conceitos de estrutura de dados
   - Execute os exemplos no banco de dados
   - Anote suas dúvidas
   - **⚠️ CUIDADO**: Alguns comandos DDL são destrutivos. Use com precaução!

2. **Leia a Aula Simplificada** (`02-aula-simplificada.md`)
   - Use as analogias para solidificar o entendimento
   - Relacione com situações do dia a dia
   - Visualize os conceitos através das metáforas

3. **Complete os Exercícios** (`03-exercicios-reflexao.md`)
   - Execute cada comando no banco de dados
   - Teste diferentes variações
   - Responda todas as perguntas de reflexão
   - Não pule nenhum exercício
   - **⚠️ IMPORTANTE**: Faça backup ou recrie o banco se necessário

4. **Estude Performance e Boas Práticas** (`04-performance-boas-praticas.md`)
   - Entenda os princípios desde o início
   - Use o checklist ao criar/modificar tabelas
   - Pense sobre integridade e performance

5. **Envie suas Respostas**
   - Compartilhe suas respostas dos exercícios
   - Aguarde o feedback antes de prosseguir

## 💡 Dicas Importantes

- **Cuidado com DDL**: Comandos DDL são poderosos e podem ser destrutivos
- **Backup**: Sempre faça backup antes de operações DDL em produção
- **Teste primeiro**: Teste comandos DDL em ambiente de desenvolvimento
- **Pense na estrutura**: Planeje a estrutura antes de criar tabelas
- **Constraints são importantes**: Use constraints para garantir integridade
- **Índices com moderação**: Índices melhoram leitura, mas podem atrasar escrita
- **Documente**: Documente suas mudanças de schema
- **Anote dúvidas**: Escreva suas perguntas para discussão posterior
- **Revisite**: Não há problema em reler seções se necessário

## ⚠️ Avisos Importantes

### Operações Destrutivas

Alguns comandos DDL são **irreversíveis** ou difíceis de reverter:

- **DROP TABLE**: Remove a tabela e todos os dados permanentemente
- **TRUNCATE TABLE**: Remove todos os dados (no SQLite, use DELETE)
- **ALTER TABLE DROP COLUMN**: Remove uma coluna e seus dados

**Sempre**:
- Faça backup antes de executar comandos destrutivos
- Teste em ambiente de desenvolvimento primeiro
- Leia o comando duas vezes antes de executar

### Recriar o Banco de Dados

Se você precisar recriar o banco de dados durante os exercícios:

```bash
go run init_database.go
```

Isso recriará o banco `biblioteca.db` do zero.

## 🔗 Recursos Adicionais

- [SQLite CREATE TABLE](https://www.sqlite.org/lang_createtable.html)
- [SQLite ALTER TABLE](https://www.sqlite.org/lang_altertable.html)
- [SQLite Data Types](https://www.sqlite.org/datatype3.html)
- [Database Normalization](https://www.studytonight.com/dbms/database-normalization.php)

## ❓ Próximos Passos

Após completar esta aula e receber feedback:

1. Revise os pontos destacados no feedback
2. Pratique mais se necessário
3. Quando estiver pronto, informe qual será o tópico da próxima aula

---

**Bons estudos! 🚀**


