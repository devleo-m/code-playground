# Aula 4: Aggregate Queries (Consultas de Agregação)

Bem-vindo à quarta aula do curso de SQL! Esta aula apresenta as Aggregate Queries (Consultas de Agregação), que permitem realizar cálculos e análises sobre conjuntos de dados, retornando valores resumidos e estatísticas.

## 📚 Estrutura da Aula

Esta aula segue o ciclo de ensino de 5 etapas. Siga a ordem indicada:

### 1. Aula Principal
**Arquivo**: `01-aula-principal.md`

Conteúdo técnico completo sobre:
- O que são Aggregate Queries e por que são importantes
- Funções de agregação fundamentais: COUNT, SUM, AVG, MIN, MAX
- GROUP BY: Agrupando dados por categorias
- HAVING: Filtrando grupos baseado em condições agregadas
- Diferença entre WHERE e HAVING
- Tratamento de NULL em funções de agregação
- Ordem correta das cláusulas em aggregate queries
- Erros comuns e como evitá-los

**Tempo estimado**: 90-120 minutos

---

### 2. Aula Simplificada
**Arquivo**: `02-aula-simplificada.md`

Mesmo conteúdo, mas explicado com:
- Analogias do dia a dia (contador de pessoas, caixa registradora, média escolar)
- Metáforas visuais (organizar livros por gênero, filtrar caixas)
- Exemplos práticos da biblioteca
- Comparações com conceitos conhecidos

**Tempo estimado**: 60-75 minutos

---

### 3. Exercícios e Reflexão
**Arquivo**: `03-exercicios-reflexao.md`

Conjunto de exercícios práticos incluindo:
- Exercícios de funções de agregação básicas (COUNT, SUM, AVG, MIN, MAX)
- Exercícios de GROUP BY (agrupando por categoria, autor, etc.)
- Exercícios de HAVING (filtrando grupos)
- Exercícios combinando WHERE e HAVING
- Análise de empréstimos e dados temporais
- Problemas que exigem raciocínio sobre agregação de dados
- Perguntas de reflexão sobre eficiência, performance, escalabilidade e design

**Tempo estimado**: 120-150 minutos

**⚠️ IMPORTANTE**: Complete todos os exercícios e responda as perguntas de reflexão antes de prosseguir.

---

### 4. Performance, Boas Práticas e Otimização
**Arquivo**: `04-performance-boas-praticas.md`

Conteúdo sobre:
- Impacto de aggregate queries na performance
- Quando e como usar índices em aggregate queries
- Otimização de GROUP BY e funções de agregação
- Otimização de JOINs em aggregate queries
- Análise de performance com EXPLAIN QUERY PLAN
- Boas práticas de nomenclatura e tratamento de NULL
- Estratégias de cache e escalabilidade
- Segurança em aggregate queries
- Monitoramento e troubleshooting
- Normalização vs desnormalização para agregações

**Tempo estimado**: 60-90 minutos

---

### 5. Análise e Feedback
**Status**: Aguardando resposta do aluno

Após completar os exercícios, envie suas respostas para análise. O tutor fornecerá feedback construtivo sobre:
- Correção sintática das aggregate queries
- Eficiência e performance das queries
- Uso adequado de GROUP BY e HAVING
- Compreensão dos conceitos de agregação
- Áreas que necessitam de melhoria

---

## 🎯 Objetivos de Aprendizado

Ao final desta aula, você deve ser capaz de:

- [ ] Entender o que são Aggregate Queries e quando usá-las
- [ ] Usar funções de agregação: COUNT, SUM, AVG, MIN, MAX
- [ ] Compreender a diferença entre COUNT(*) e COUNT(coluna)
- [ ] Agrupar dados usando GROUP BY
- [ ] Filtrar grupos usando HAVING
- [ ] Diferenciar entre WHERE (filtra linhas) e HAVING (filtra grupos)
- [ ] Combinar múltiplas funções de agregação em uma única query
- [ ] Entender como NULL é tratado em funções de agregação
- [ ] Escrever aggregate queries complexas com JOINs
- [ ] Otimizar aggregate queries usando índices
- [ ] Analisar performance usando EXPLAIN QUERY PLAN
- [ ] Aplicar boas práticas ao escrever aggregate queries

## 📖 Pré-requisitos

- Conclusão da Aula 1 (Introdução ao SQL)
- Conclusão da Aula 2 (Sintaxe Básica de SQL)
- Conclusão da Aula 3 (Data Definition Language - DDL)
- Banco de dados `biblioteca.db` criado e funcionando
- Acesso ao SQLite (via CLI ou ferramenta visual)
- Compreensão de SELECT, WHERE, JOIN, e estrutura de tabelas

## 🔄 Revisão Rápida das Aulas Anteriores

Antes de começar, vamos revisar os conceitos-chave das aulas anteriores:

### Da Aula 1:
- **SQL** é composto por DDL, DML, DCL e Queries
- **Bancos relacionais** organizam dados em tabelas com relacionamentos
- **Tabelas** têm linhas (registros) e colunas (campos)

### Da Aula 2:
- **SELECT** para consultar dados
- **WHERE** para filtrar linhas individuais
- **JOIN** para combinar dados de múltiplas tabelas
- **Tipos de dados**: INTEGER, TEXT, DATE, etc.
- **Operadores**: comparação, lógicos, aritméticos

### Da Aula 3:
- **DDL** cria e modifica estruturas (CREATE, ALTER, DROP)
- **Constraints**: PRIMARY KEY, FOREIGN KEY, UNIQUE, NOT NULL
- **Índices** melhoram performance de consultas

Se você não se lembra desses conceitos, revise as aulas anteriores antes de prosseguir.

## 🚀 Como Usar Esta Aula

1. **Leia a Aula Principal** (`01-aula-principal.md`)
   - Não se preocupe em decorar tudo
   - Foque em entender os conceitos de agregação
   - Execute TODOS os exemplos no banco de dados
   - Anote suas dúvidas
   - Preste atenção especial à diferença entre WHERE e HAVING

2. **Leia a Aula Simplificada** (`02-aula-simplificada.md`)
   - Use as analogias para solidificar o entendimento
   - Relacione com situações do dia a dia
   - Visualize os conceitos através das metáforas
   - Compare com exemplos conhecidos (média escolar, contador de pessoas)

3. **Complete os Exercícios** (`03-exercicios-reflexao.md`)
   - Execute cada query no banco de dados
   - Teste diferentes variações
   - Responda TODAS as perguntas de reflexão
   - Não pule nenhum exercício
   - **⚠️ IMPORTANTE**: As perguntas de reflexão são cruciais!

4. **Estude Performance e Boas Práticas** (`04-performance-boas-praticas.md`)
   - Entenda os princípios desde o início
   - Use EXPLAIN QUERY PLAN para analisar suas queries
   - Pense sobre índices e otimização
   - Aprenda sobre cache e escalabilidade

5. **Envie suas Respostas**
   - Compartilhe suas respostas dos exercícios
   - Inclua suas respostas às perguntas de reflexão
   - Aguarde o feedback antes de prosseguir

## 💡 Dicas Importantes

- **Pratique muito**: Execute TODAS as queries no banco de dados real
- **Teste variações**: Modifique os exemplos para ver o que acontece
- **Entenda WHERE vs HAVING**: Esta é uma das diferenças mais importantes!
- **Pense sobre performance**: Aggregate queries podem ser lentas sem otimização
- **Use EXPLAIN**: Aprenda a analisar o plano de execução
- **Anote dúvidas**: Escreva suas perguntas para discussão posterior
- **Revisite**: Não há problema em reler seções se necessário
- **Experimente erros**: Tente queries incorretas para entender os erros

## ⚠️ Avisos Importantes

### Erros Comuns

Alguns erros são muito comuns em aggregate queries:

1. **Esquecer GROUP BY**: Ao usar funções de agregação com colunas não agregadas
2. **Usar função de agregação no WHERE**: Use HAVING ao invés de WHERE
3. **Coluna não agrupada no SELECT**: Todas as colunas não agregadas devem estar no GROUP BY

### Performance

- Aggregate queries podem ser lentas com muitos dados
- Sempre use índices nas colunas de GROUP BY e JOIN
- Filtre com WHERE antes de agrupar quando possível
- Monitore performance com EXPLAIN QUERY PLAN

### Recriar o Banco de Dados

Se você precisar recriar o banco de dados durante os exercícios:

```bash
go run init_database.go
```

Isso recriará o banco `biblioteca.db` do zero.

## 🔗 Recursos Adicionais

- [SQLite Aggregate Functions](https://www.sqlite.org/lang_aggfunc.html)
- [SQLite GROUP BY](https://www.sqlite.org/lang_select.html#resultset)
- [SQLite EXPLAIN QUERY PLAN](https://www.sqlite.org/eqp.html)
- [SQL Aggregate Functions Tutorial](https://www.w3schools.com/sql/sql_count_avg_sum.asp)

## 📊 Exemplos Rápidos

### Contar Registros
```sql
SELECT COUNT(*) FROM livros;
```

### Somar Valores
```sql
SELECT SUM(quantidade_disponivel) FROM livros;
```

### Agrupar por Categoria
```sql
SELECT categoria_id, COUNT(*) 
FROM livros 
GROUP BY categoria_id;
```

### Filtrar Grupos
```sql
SELECT categoria_id, COUNT(*) 
FROM livros 
GROUP BY categoria_id
HAVING COUNT(*) > 2;
```

## ❓ Próximos Passos

Após completar esta aula e receber feedback:

1. Revise os pontos destacados no feedback
2. Pratique mais se necessário
3. Quando estiver pronto, informe qual será o tópico da próxima aula

---

**Bons estudos! 🚀**

**Lembre-se**: Aggregate queries são fundamentais para análise de dados e geração de relatórios. Domine esses conceitos e você terá uma base sólida para trabalhar com dados!
