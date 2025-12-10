# Aula 7: Sub Queries (Subconsultas)

Bem-vindo à sétima aula do curso de SQL! Esta aula apresenta as **Sub Queries** (Subconsultas), também conhecidas como queries aninhadas ou consultas internas. Subqueries são queries SQL embutidas dentro de outra query, permitindo realizar consultas complexas e criar critérios dinâmicos de forma elegante e poderosa.

## 📚 Estrutura da Aula

Esta aula segue o ciclo de ensino de 5 etapas. Siga a ordem indicada:

### 1. Aula Principal
**Arquivo**: `01-aula-principal.md`

Conteúdo técnico completo sobre:
- O que são Sub Queries e por que são essenciais
- Tipos de Sub Queries: Scalar, Column, Row, Table
- Subqueries no SELECT, FROM, WHERE e HAVING
- Nested Subqueries (Subqueries aninhadas)
- Correlated Subqueries (Subqueries correlacionadas)
- Diferenças entre subqueries e JOINs
- Quando usar subqueries vs JOINs
- Sintaxe e exemplos práticos com o banco de dados da biblioteca
- Limitações e considerações de performance

**Tempo estimado**: 90-120 minutos

---

### 2. Aula Simplificada
**Arquivo**: `02-aula-simplificada.md`

Mesmo conteúdo, mas explicado com:
- Analogias do dia a dia (subqueries como perguntas dentro de perguntas)
- Metáforas visuais (consultas aninhadas como caixas dentro de caixas)
- Exemplos práticos da biblioteca
- Comparações com conceitos conhecidos (filtros dinâmicos, consultas auxiliares)

**Tempo estimado**: 60-75 minutos

---

### 3. Exercícios e Reflexão
**Arquivo**: `03-exercicios-reflexao.md`

Conjunto de exercícios práticos incluindo:
- Exercícios de escrita de subqueries de diferentes tipos
- Exercícios de análise de subqueries existentes
- Problemas que exigem raciocínio sobre quando usar subqueries
- Comparação entre subqueries e JOINs
- Perguntas de reflexão sobre eficiência e legibilidade
- Exercícios com nested e correlated subqueries

**Tempo estimado**: 120-150 minutos

**⚠️ IMPORTANTE**: Complete todos os exercícios e responda as perguntas de reflexão antes de prosseguir.

---

### 4. Performance, Boas Práticas e Otimização
**Arquivo**: `04-performance-boas-praticas.md`

Conteúdo sobre:
- Impacto de subqueries na performance
- Quando subqueries são mais eficientes que JOINs
- Quando JOINs são mais eficientes que subqueries
- Otimização de correlated subqueries
- Índices e subqueries
- Boas práticas de escrita de subqueries
- Evitando subqueries desnecessárias
- Troubleshooting de queries lentas com subqueries
- Estratégias de otimização e reescrita de queries

**Tempo estimado**: 60-90 minutos

---

### 5. Análise e Feedback
**Status**: Aguardando resposta do aluno

Após completar os exercícios, envie suas respostas para análise. O tutor fornecerá feedback construtivo sobre:
- Correção sintática das subqueries
- Escolha adequada entre subqueries e JOINs
- Eficiência e performance das queries
- Compreensão dos conceitos de subqueries
- Legibilidade e manutenibilidade do código
- Áreas que necessitam de melhoria

---

## 🎯 Objetivos de Aprendizado

Ao final desta aula, você deve ser capaz de:

- [ ] Entender o que são subqueries e sua importância
- [ ] Criar subqueries escalares (retornando um único valor)
- [ ] Criar subqueries que retornam uma coluna de valores
- [ ] Criar subqueries que retornam uma linha de valores
- [ ] Criar subqueries que retornam uma tabela completa
- [ ] Usar subqueries em SELECT, FROM, WHERE e HAVING
- [ ] Criar nested subqueries (subqueries aninhadas)
- [ ] Criar correlated subqueries (subqueries correlacionadas)
- [ ] Escolher entre subqueries e JOINs adequadamente
- [ ] Entender o impacto de subqueries na performance
- [ ] Aplicar boas práticas ao escrever subqueries
- [ ] Resolver problemas complexos usando subqueries

## 📖 Pré-requisitos

- Conclusão da Aula 1 (Introdução ao SQL)
- Conclusão da Aula 2 (Sintaxe Básica de SQL)
- Conclusão da Aula 3 (Data Definition Language - DDL)
- Conclusão da Aula 4 (Aggregate Queries)
- Conclusão da Aula 5 (Data Constraints)
- Conclusão da Aula 6 (SQL JOIN Queries)
- Banco de dados `biblioteca.db` criado e funcionando
- Acesso ao SQLite (via CLI ou ferramenta visual)
- Compreensão de SELECT, WHERE, JOINs e funções de agregação
- Compreensão de GROUP BY e HAVING

## 🔄 Revisão Rápida das Aulas Anteriores

Antes de começar, vamos revisar os conceitos-chave das aulas anteriores:

### Da Aula 1:
- **SQL** permite consultar e manipular dados relacionais
- **Bancos relacionais** organizam dados em tabelas com relacionamentos

### Da Aula 2:
- **SELECT** para recuperar dados
- **WHERE** para filtrar linhas
- **Relacionamentos** entre tabelas através de FOREIGN KEY

### Da Aula 4:
- **Aggregate queries** para análise de dados
- **GROUP BY** e **HAVING** para agrupamento
- Funções de agregação (COUNT, SUM, AVG, etc.)

### Da Aula 6:
- **JOINs** para combinar dados de múltiplas tabelas
- **INNER JOIN**, **LEFT JOIN**, etc.
- Relacionamentos entre tabelas

Se você não se lembra desses conceitos, revise as aulas anteriores antes de prosseguir.

## 🚀 Como Usar Esta Aula

1. **Leia a Aula Principal** (`01-aula-principal.md`)
   - Não se preocupe em decorar tudo
   - Foque em entender quando e por que usar subqueries
   - Execute TODOS os exemplos no banco de dados
   - Anote suas dúvidas
   - Preste atenção especial às diferenças entre tipos de subqueries
   - Entenda a diferença entre subqueries e JOINs

2. **Leia a Aula Simplificada** (`02-aula-simplificada.md`)
   - Use as analogias para solidificar o entendimento
   - Relacione com situações do dia a dia
   - Visualize os conceitos através das metáforas
   - Compare com exemplos conhecidos

3. **Complete os Exercícios** (`03-exercicios-reflexao.md`)
   - Execute cada query no banco de dados
   - Teste diferentes variações
   - Compare resultados de subqueries vs JOINs
   - Responda TODAS as perguntas de reflexão
   - Não pule nenhum exercício
   - **⚠️ IMPORTANTE**: As perguntas de reflexão são cruciais!

4. **Estude Performance e Boas Práticas** (`04-performance-boas-praticas.md`)
   - Entenda os princípios desde o início
   - Pense sobre quando usar subqueries vs JOINs
   - Aprenda sobre otimização
   - Entenda o impacto na performance

5. **Envie suas Respostas**
   - Compartilhe suas respostas dos exercícios
   - Inclua suas respostas às perguntas de reflexão
   - Aguarde o feedback antes de prosseguir

## 💡 Dicas Importantes

- **Pratique muito**: Execute TODAS as queries no banco de dados real
- **Entenda o contexto**: Subqueries são poderosas, mas nem sempre são a melhor solução
- **Compare com JOINs**: Muitas vezes você pode resolver o mesmo problema com JOINs
- **Pense sobre performance**: Subqueries podem ser lentas se mal utilizadas
- **Anote dúvidas**: Escreva suas perguntas para discussão posterior
- **Revisite**: Não há problema em reler seções se necessário
- **Experimente erros**: Tente queries incorretas para entender os erros
- **Visualize mentalmente**: Pense em como a subquery é executada

## ⚠️ Avisos Importantes

### Erros Comuns

Alguns erros são muito comuns ao trabalhar com subqueries:

1. **Subquery retorna múltiplas linhas quando deveria retornar uma**: Use operadores adequados (IN, ANY, ALL)
2. **Subquery correlacionada mal escrita**: Pode causar performance muito ruim
3. **Subquery desnecessária**: Muitas vezes um JOIN resolve o problema de forma mais eficiente
4. **Esquecer de usar aliases**: Pode causar ambiguidade em subqueries correlacionadas
5. **Subquery muito complexa**: Pode ser difícil de entender e manter

### Performance

- Subqueries podem ser executadas múltiplas vezes (especialmente correlated subqueries)
- JOINs geralmente são mais eficientes para combinar dados de múltiplas tabelas
- Subqueries escalares podem ser lentas em tabelas grandes
- Índices são essenciais para performance de subqueries

### Recriar o Banco de Dados

Se você precisar recriar o banco de dados durante os exercícios:

```bash
go run init_database.go
```

Isso recriará o banco `biblioteca.db` do zero.

## 🔗 Recursos Adicionais

- [SQLite Subquery Documentation](https://www.sqlite.org/lang_select.html#subqueries)
- [SQL Subqueries Tutorial - W3Schools](https://www.w3schools.com/sql/sql_subqueries.asp)
- [SQL Subqueries Explained - SQLBolt](https://sqlbolt.com/lesson/select_queries_order_of_execution)

## 📊 Exemplos Rápidos

### Subquery Escalar
```sql
SELECT titulo, 
       (SELECT COUNT(*) FROM emprestimos WHERE livro_id = livros.id) AS total_emprestimos
FROM livros;
```

### Subquery em WHERE
```sql
SELECT * FROM livros
WHERE autor_id IN (SELECT id FROM autores WHERE nacionalidade = 'Brasileiro');
```

### Correlated Subquery
```sql
SELECT l.titulo
FROM livros l
WHERE l.quantidade_disponivel > (
    SELECT AVG(quantidade_disponivel) 
    FROM livros 
    WHERE categoria_id = l.categoria_id
);
```

## ❓ Próximos Passos

Após completar esta aula e receber feedback:

1. Revise os pontos destacados no feedback
2. Pratique mais se necessário
3. Quando estiver pronto, informe qual será o tópico da próxima aula

---

**Bons estudos! 🚀**

**Lembre-se**: Subqueries são uma ferramenta poderosa que permite resolver problemas complexos de forma elegante. Domine esses conceitos e você terá uma base sólida para consultas avançadas e análises de dados sofisticadas!





