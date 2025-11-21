# Aula 5: Data Constraints (Restrições de Dados)

Bem-vindo à quinta aula do curso de SQL! Esta aula apresenta as **Data Constraints** (Restrições de Dados), que são regras aplicadas a colunas ou tabelas para garantir a integridade e consistência dos dados no banco de dados.

## 📚 Estrutura da Aula

Esta aula segue o ciclo de ensino de 5 etapas. Siga a ordem indicada:

### 1. Aula Principal
**Arquivo**: `01-aula-principal.md`

Conteúdo técnico completo sobre:
- O que são Data Constraints e por que são essenciais
- PRIMARY KEY: Identificadores únicos e integridade de entidade
- FOREIGN KEY: Relacionamentos e integridade referencial
- UNIQUE: Garantindo valores únicos
- NOT NULL: Campos obrigatórios
- CHECK: Validação de dados customizada
- Como criar, modificar e remover constraints
- Comportamento de constraints em operações de INSERT, UPDATE e DELETE
- Erros comuns e como evitá-los

**Tempo estimado**: 90-120 minutos

---

### 2. Aula Simplificada
**Arquivo**: `02-aula-simplificada.md`

Mesmo conteúdo, mas explicado com:
- Analogias do dia a dia (RG, carteira de identidade, validações)
- Metáforas visuais (chaves, cadeados, porteiros)
- Exemplos práticos da biblioteca
- Comparações com conceitos conhecidos (formulários, validações)

**Tempo estimado**: 60-75 minutos

---

### 3. Exercícios e Reflexão
**Arquivo**: `03-exercicios-reflexao.md`

Conjunto de exercícios práticos incluindo:
- Exercícios de criação de constraints (PRIMARY KEY, FOREIGN KEY, UNIQUE, NOT NULL, CHECK)
- Exercícios de modificação de tabelas existentes
- Exercícios de teste de constraints (tentando violar regras)
- Análise de integridade referencial
- Problemas que exigem raciocínio sobre design de banco de dados
- Perguntas de reflexão sobre integridade, performance e escalabilidade

**Tempo estimado**: 120-150 minutos

**⚠️ IMPORTANTE**: Complete todos os exercícios e responda as perguntas de reflexão antes de prosseguir.

---

### 4. Performance, Boas Práticas e Otimização
**Arquivo**: `04-performance-boas-praticas.md`

Conteúdo sobre:
- Impacto de constraints na performance
- Quando usar cada tipo de constraint
- Índices automáticos criados por constraints
- Estratégias de validação: banco vs aplicação
- Boas práticas de design de constraints
- Segurança e integridade de dados
- Monitoramento e troubleshooting
- Trade-offs entre constraints e flexibilidade

**Tempo estimado**: 60-90 minutos

---

### 5. Análise e Feedback
**Status**: Aguardando resposta do aluno

Após completar os exercícios, envie suas respostas para análise. O tutor fornecerá feedback construtivo sobre:
- Correção sintática das constraints criadas
- Design adequado de constraints
- Compreensão dos conceitos de integridade
- Áreas que necessitam de melhoria

---

## 🎯 Objetivos de Aprendizado

Ao final desta aula, você deve ser capaz de:

- [ ] Entender o que são Data Constraints e sua importância
- [ ] Criar e gerenciar PRIMARY KEY constraints
- [ ] Criar e gerenciar FOREIGN KEY constraints
- [ ] Criar e gerenciar UNIQUE constraints
- [ ] Criar e gerenciar NOT NULL constraints
- [ ] Criar e gerenciar CHECK constraints
- [ ] Entender integridade referencial e cascata
- [ ] Modificar constraints em tabelas existentes
- [ ] Testar e validar constraints
- [ ] Entender o impacto de constraints na performance
- [ ] Aplicar boas práticas ao projetar constraints
- [ ] Decidir quando usar constraints vs validação na aplicação

## 📖 Pré-requisitos

- Conclusão da Aula 1 (Introdução ao SQL)
- Conclusão da Aula 2 (Sintaxe Básica de SQL)
- Conclusão da Aula 3 (Data Definition Language - DDL)
- Conclusão da Aula 4 (Aggregate Queries)
- Banco de dados `biblioteca.db` criado e funcionando
- Acesso ao SQLite (via CLI ou ferramenta visual)
- Compreensão de CREATE TABLE, ALTER TABLE e estrutura de tabelas

## 🔄 Revisão Rápida das Aulas Anteriores

Antes de começar, vamos revisar os conceitos-chave das aulas anteriores:

### Da Aula 1:
- **SQL** é composto por DDL, DML, DCL e Queries
- **Bancos relacionais** organizam dados em tabelas com relacionamentos
- **Integridade de dados** é fundamental para bancos relacionais

### Da Aula 2:
- **SELECT, INSERT, UPDATE, DELETE** para manipular dados
- **WHERE** para filtrar linhas
- **JOIN** para combinar dados de múltiplas tabelas

### Da Aula 3:
- **DDL** cria e modifica estruturas (CREATE, ALTER, DROP)
- **CREATE TABLE** define estrutura de tabelas
- **ALTER TABLE** modifica tabelas existentes

### Da Aula 4:
- **Aggregate queries** para análise de dados
- **GROUP BY** e **HAVING** para agrupamento

Se você não se lembra desses conceitos, revise as aulas anteriores antes de prosseguir.

## 🚀 Como Usar Esta Aula

1. **Leia a Aula Principal** (`01-aula-principal.md`)
   - Não se preocupe em decorar tudo
   - Foque em entender o propósito de cada constraint
   - Execute TODOS os exemplos no banco de dados
   - Anote suas dúvidas
   - Preste atenção especial à integridade referencial

2. **Leia a Aula Simplificada** (`02-aula-simplificada.md`)
   - Use as analogias para solidificar o entendimento
   - Relacione com situações do dia a dia
   - Visualize os conceitos através das metáforas
   - Compare com exemplos conhecidos (RG, validações de formulário)

3. **Complete os Exercícios** (`03-exercicios-reflexao.md`)
   - Execute cada query no banco de dados
   - Teste diferentes variações
   - Tente violar constraints para entender o comportamento
   - Responda TODAS as perguntas de reflexão
   - Não pule nenhum exercício
   - **⚠️ IMPORTANTE**: As perguntas de reflexão são cruciais!

4. **Estude Performance e Boas Práticas** (`04-performance-boas-praticas.md`)
   - Entenda os princípios desde o início
   - Pense sobre quando usar cada constraint
   - Aprenda sobre trade-offs e decisões de design

5. **Envie suas Respostas**
   - Compartilhe suas respostas dos exercícios
   - Inclua suas respostas às perguntas de reflexão
   - Aguarde o feedback antes de prosseguir

## 💡 Dicas Importantes

- **Pratique muito**: Execute TODAS as queries no banco de dados real
- **Teste violações**: Tente inserir dados inválidos para ver como constraints funcionam
- **Entenda integridade referencial**: Esta é uma das partes mais importantes!
- **Pense sobre design**: Constraints afetam como você estrutura seu banco
- **Anote dúvidas**: Escreva suas perguntas para discussão posterior
- **Revisite**: Não há problema em reler seções se necessário
- **Experimente erros**: Tente queries incorretas para entender os erros

## ⚠️ Avisos Importantes

### Erros Comuns

Alguns erros são muito comuns ao trabalhar com constraints:

1. **Esquecer FOREIGN KEY**: Criar relacionamentos sem constraints
2. **Violar integridade referencial**: Tentar deletar registros referenciados
3. **Duplicar PRIMARY KEY**: Tentar inserir IDs duplicados
4. **Violar UNIQUE**: Tentar inserir valores duplicados em colunas UNIQUE
5. **Inserir NULL em NOT NULL**: Tentar inserir valores nulos em campos obrigatórios

### Performance

- Constraints criam índices automáticos (PRIMARY KEY, UNIQUE, FOREIGN KEY)
- CHECK constraints podem impactar performance em INSERT/UPDATE
- FOREIGN KEY com CASCADE pode ter impacto em DELETE

### Recriar o Banco de Dados

Se você precisar recriar o banco de dados durante os exercícios:

```bash
go run init_database.go
```

Isso recriará o banco `biblioteca.db` do zero.

## 🔗 Recursos Adicionais

- [SQLite CREATE TABLE](https://www.sqlite.org/lang_createtable.html)
- [SQLite Foreign Keys](https://www.sqlite.org/foreignkeys.html)
- [SQLite Constraints](https://www.sqlite.org/lang_createtable.html#constraints)
- [Database Constraints Tutorial](https://www.w3schools.com/sql/sql_constraints.asp)

## 📊 Exemplos Rápidos

### PRIMARY KEY
```sql
CREATE TABLE livros (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    titulo TEXT NOT NULL
);
```

### FOREIGN KEY
```sql
CREATE TABLE emprestimos (
    id INTEGER PRIMARY KEY,
    livro_id INTEGER,
    FOREIGN KEY (livro_id) REFERENCES livros(id)
);
```

### UNIQUE
```sql
CREATE TABLE usuarios (
    id INTEGER PRIMARY KEY,
    email TEXT UNIQUE NOT NULL
);
```

### NOT NULL
```sql
CREATE TABLE autores (
    id INTEGER PRIMARY KEY,
    nome TEXT NOT NULL
);
```

### CHECK
```sql
CREATE TABLE livros (
    id INTEGER PRIMARY KEY,
    quantidade_disponivel INTEGER CHECK (quantidade_disponivel >= 0)
);
```

## ❓ Próximos Passos

Após completar esta aula e receber feedback:

1. Revise os pontos destacados no feedback
2. Pratique mais se necessário
3. Quando estiver pronto, informe qual será o tópico da próxima aula

---

**Bons estudos! 🚀**

**Lembre-se**: Constraints são fundamentais para garantir a integridade e qualidade dos dados. Domine esses conceitos e você terá uma base sólida para projetar bancos de dados confiáveis!
