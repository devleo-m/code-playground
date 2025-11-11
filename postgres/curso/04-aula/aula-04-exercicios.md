# **Aula 4 - Exercícios e Reflexão**

## 📝 Instruções

Esta é a **última aula antes da avaliação final**! Os exercícios consolidam TUDO que você aprendeu nas Aulas 1-4.

Complete todos os exercícios usando **suas próprias palavras** e código SQL correto.

**IMPORTANTE:** Esta aula testa conhecimento acumulado! Você precisa dominar:

- Aula 1: PostgreSQL, ACID, comparações
- Aula 2: Modelo relacional, chaves, relacionamentos, integridade
- Aula 3: Hierarquia, queries SQL, tipos de dados
- Aula 4: Domains, constraints avançadas, NULL

---

## 🎯 Exercício 1: Domains Customizados

### 1.1 - Criando Domains

Você está criando um sistema para uma imobiliária brasileira. Crie os seguintes domains:

a) **Domain para CPF brasileiro**

- Formato: XXX.XXX.XXX-XX
- Deve validar o formato com regex

```sql
[SEU CÓDIGO AQUI]
```

b) **Domain para CEP brasileiro**

- Formato: XXXXX-XXX
- Deve validar o formato

```sql
[SEU CÓDIGO AQUI]
```

c) **Domain para telefone brasileiro** (celular)

- Formato: (XX) 9XXXX-XXXX
- Deve validar o formato

```sql
[SEU CÓDIGO AQUI]
```

d) **Domain para valor monetário positivo**

- Tipo base: DECIMAL(12, 2)
- Deve ser maior que zero

```sql
[SEU CÓDIGO AQUI]
```

e) **Domain para percentual**

- Tipo base: DECIMAL(5, 2)
- Deve estar entre 0 e 100

```sql
[SEU CÓDIGO AQUI]
```

### 1.2 - Justificando o Uso de Domains

**Explique:**

a) Por que usar domain `cpf_brasileiro` em vez de `CHAR(14)` direto na tabela?

```
[SUA RESPOSTA - MÍNIMO 3 RAZÕES]
```

b) Se você tiver 20 tabelas usando `cpf_brasileiro` e precisar mudar a validação, quantos lugares precisa alterar? Como isso ajuda na manutenção?

```
[SUA RESPOSTA]
```

c) Qual a vantagem de criar um domain `valor_monetario_positivo` em vez de usar `CHECK (preco > 0)` em cada tabela?

```
[SUA RESPOSTA]
```

---

## 🎯 Exercício 2: Modelagem Completa com Domains e Constraints

Você precisa criar o banco de dados completo para a imobiliária. O sistema gerencia:

**Entidades:**

- **Proprietários**: CPF, nome, telefone, email
- **Imóveis**: Código único, endereço (logradouro, número, bairro, cidade, CEP), tipo (casa/apartamento/terreno), área (m²), quartos, valor
- **Clientes**: CPF, nome, telefone, email
- **Visitas**: Qual cliente visitou qual imóvel, data e hora da visita, feedback (texto)
- **Vendas**: Qual imóvel foi vendido para qual cliente, data da venda, valor de venda, comissão (percentual), status (pendente/concluída/cancelada)

### 2.1 - Criação Completa do Sistema

```sql
-- 1. Criar database
[SEU CÓDIGO AQUI]

-- 2. Conectar ao database
[SEU CÓDIGO AQUI]

-- 3. Criar todos os domains necessários (use os do Ex 1.1 + outros que precisar)
[SEU CÓDIGO AQUI]

-- 4. Criar ENUM para tipo de imóvel
[SEU CÓDIGO AQUI]

-- 5. Criar ENUM para status de venda
[SEU CÓDIGO AQUI]

-- 6. Criar tabela de proprietários
[SEU CÓDIGO AQUI]

-- 7. Criar tabela de imóveis
-- IMPORTANTE: Inclua:
-- - PRIMARY KEY
-- - FOREIGN KEY para proprietário
-- - CHECK para área (deve ser positiva)
-- - CHECK para quartos (0 ou positivo)
-- - CHECK para valor (positivo)
-- - NOT NULL nos campos obrigatórios
[SEU CÓDIGO AQUI]

-- 8. Criar tabela de clientes
[SEU CÓDIGO AQUI]

-- 9. Criar tabela de visitas
-- IMPORTANTE:
-- - FOREIGN KEY para cliente
-- - FOREIGN KEY para imóvel
-- - Feedback é opcional (pode ser NULL)
[SEU CÓDIGO AQUI]

-- 10. Criar tabela de vendas
-- IMPORTANTE:
-- - FOREIGN KEY para imóvel (com ON DELETE RESTRICT - não pode deletar imóvel vendido)
-- - FOREIGN KEY para cliente
-- - CHECK para comissão (use domain percentual)
-- - CHECK: data_venda não pode ser no futuro
-- - DEFAULT para status ('pendente')
[SEU CÓDIGO AQUI]
```

### 2.2 - Justifique suas Escolhas

a) Por que você usou `ON DELETE RESTRICT` na venda de imóveis?

```
[SUA RESPOSTA]
```

b) Que campos você marcou como `NOT NULL`? Por quê esses e não outros?

```
[SUA RESPOSTA]
```

c) Por que usar ENUM para `tipo_imovel` e `status_venda`?

```
[SUA RESPOSTA]
```

d) Identifique todos os relacionamentos e seus tipos (1:1, 1:N, N:M):

```
Proprietários ↔ Imóveis: [TIPO E JUSTIFICATIVA]
Clientes ↔ Visitas: [TIPO E JUSTIFICATIVA]
Imóveis ↔ Visitas: [TIPO E JUSTIFICATIVA]
Clientes ↔ Vendas: [TIPO E JUSTIFICATIVA]
Imóveis ↔ Vendas: [TIPO E JUSTIFICATIVA]
```

---

## 🎯 Exercício 3: Populando o Banco de Dados

Use o banco de dados do Exercício 2.

### 3.1 - Inserir Dados

```sql
-- Inserir 5 proprietários
[SEU CÓDIGO AQUI]

-- Inserir 10 imóveis (variando tipos, proprietários)
[SEU CÓDIGO AQUI]

-- Inserir 8 clientes
[SEU CÓDIGO AQUI]

-- Inserir 15 visitas (diferentes combinações de clientes e imóveis)
[SEU CÓDIGO AQUI]

-- Inserir 5 vendas (status variados)
[SEU CÓDIGO AQUI]
```

### 3.2 - Testando Constraints

Tente fazer as seguintes operações e **preveja** o resultado (ERRO ou SUCESSO):

a) Inserir proprietário com CPF inválido (sem pontos e traço):

```sql
INSERT INTO proprietarios (cpf, nome) VALUES ('12345678900', 'Teste');

Resultado esperado: [ERRO ou SUCESSO? POR QUÊ?]
```

b) Inserir imóvel com área negativa:

```sql
INSERT INTO imoveis (codigo, tipo, area, valor, proprietario_id)
VALUES ('IM001', 'casa', -50, 200000, 1);

Resultado esperado: [ERRO ou SUCESSO? POR QUÊ?]
```

c) Inserir venda com comissão de 150%:

```sql
INSERT INTO vendas (imovel_id, cliente_id, valor_venda, comissao)
VALUES (1, 1, 300000, 150.00);

Resultado esperado: [ERRO ou SUCESSO? POR QUÊ?]
```

d) Deletar cliente que já fez visitas:

```sql
DELETE FROM clientes WHERE id = 1;

Resultado esperado: [ERRO ou SUCESSO? POR QUÊ? Depende de que constraint?]
```

e) Inserir visita para imóvel que não existe:

```sql
INSERT INTO visitas (cliente_id, imovel_id, data_visita)
VALUES (1, 9999, '2024-12-01');

Resultado esperado: [ERRO ou SUCESSO? POR QUÊ?]
```

---

## 🎯 Exercício 4: Queries Complexas

Use o banco do Exercício 2 (já populado no Ex 3).

### 4.1 - Consultas Básicas

a) Listar todos os imóveis do tipo 'apartamento' com mais de 2 quartos, ordenados por valor (do menor ao maior):

```sql
[SEU CÓDIGO AQUI]
```

b) Listar todos os proprietários que têm pelo menos um imóvel cadastrado:

```sql
[SEU CÓDIGO AQUI]
```

c) Listar clientes que nunca fizeram visitas:

```sql
[SEU CÓDIGO AQUI]
```

d) Listar imóveis que nunca receberam visitas:

```sql
[SEU CÓDIGO AQUI]
```

### 4.2 - Consultas com Agregação

a) Contar quantos imóveis cada proprietário tem:

```sql
[SEU CÓDIGO AQUI]
```

b) Calcular o valor médio dos imóveis por tipo (casa, apartamento, terreno):

```sql
[SEU CÓDIGO AQUI]
```

c) Listar os 3 imóveis mais visitados (quantidade de visitas):

```sql
[SEU CÓDIGO AQUI]
```

d) Calcular o valor total de vendas concluídas:

```sql
[SEU CÓDIGO AQUI]
```

e) Calcular a comissão total que a imobiliária vai receber (soma de todas as comissões):

```sql
-- Lembre-se: comissão é percentual, valor_venda é o valor
-- Comissão em reais = valor_venda * (comissao / 100)
[SEU CÓDIGO AQUI]
```

### 4.3 - Consultas com JOIN

a) Listar todas as visitas mostrando nome do cliente, código do imóvel e data:

```sql
[SEU CÓDIGO AQUI]
```

b) Listar todas as vendas mostrando: código do imóvel, nome do cliente, nome do proprietário, valor de venda, status:

```sql
[SEU CÓDIGO AQUI]
```

c) Listar clientes que visitaram imóveis mas ainda não compraram nenhum:

```sql
[SEU CÓDIGO AQUI]
```

d) Listar imóveis que receberam visitas mas ainda não foram vendidos:

```sql
[SEU CÓDIGO AQUI]
```

---

## 🎯 Exercício 5: Trabalhando com NULL

### 5.1 - Conceitos de NULL

a) Qual a diferença entre estas três situações?

```
Cliente A: telefone = '(11) 98765-4321'
Cliente B: telefone = ''  (string vazia)
Cliente C: telefone = NULL

Explique a diferença prática:
[SUA RESPOSTA]
```

b) Por que esta query NÃO funciona?

```sql
SELECT * FROM clientes WHERE telefone = NULL;

Explique o problema e mostre a forma correta:
[SUA RESPOSTA E CÓDIGO CORRETO]
```

c) O que acontece nestas operações?

```sql
SELECT 100 + NULL;
SELECT 100 > NULL;
SELECT NULL = NULL;

Resultado de cada:
[SUA RESPOSTA]
```

### 5.2 - Usando COALESCE

a) Exibir contato de cada cliente (use telefone se tiver, senão email, senão 'Sem contato'):

```sql
[SEU CÓDIGO AQUI]
```

b) Listar imóveis com feedback das visitas, mas se não houver feedback, mostrar 'Sem feedback':

```sql
[SEU CÓDIGO AQUI]
```

### 5.3 - Usando NULLIF

a) Calcular área média dos imóveis por tipo, mas evitar divisão por zero usando NULLIF:

```sql
-- Se count for 0, NULLIF retorna NULL evitando divisão por zero
[SEU CÓDIGO AQUI]
```

b) Criar query que mostra valor do imóvel, mas se valor for igual ao valor médio, mostra NULL:

```sql
[SEU CÓDIGO AQUI]
```

---

## 🎯 Exercício 6: Constraints Avançadas

### 6.1 - Adicionando EXCLUSION Constraint

Você quer garantir que um cliente não possa ter duas visitas agendadas para o mesmo imóvel no mesmo dia.

a) Crie a constraint EXCLUSION necessária:

```sql
-- Dica: Use btree_gist extension
-- Exclude: mesmo cliente + mesmo imóvel + mesmo dia
[SEU CÓDIGO AQUI]
```

b) Teste a constraint tentando inserir conflito:

```sql
[SEU CÓDIGO DE TESTE]
```

### 6.2 - CHECK Constraints Complexas

a) Adicione constraint: imóvel tipo 'terreno' não pode ter quartos (quartos deve ser 0):

```sql
[SEU CÓDIGO AQUI]
```

b) Adicione constraint: data de visita não pode ser no futuro:

```sql
[SEU CÓDIGO AQUI]
```

c) Adicione constraint: em vendas, valor_venda deve ser pelo menos 80% do valor original do imóvel:

```sql
-- Dica: Precisa fazer JOIN na CHECK constraint
[SEU CÓDIGO AQUI OU EXPLIQUE POR QUE É DIFÍCIL/IMPOSSÍVEL]
```

---

## 🧠 Perguntas de Reflexão

### Reflexão 1: Domains vs Constraints Inline

Você viu que pode criar domains ou usar constraints diretamente nas tabelas.

**Cenário:** Sistema com 50 tabelas, 30 delas têm campos de email.

**Opção A:** Criar domain `email_type` e usar nas 30 tabelas
**Opção B:** Colocar `CHECK (email ~* 'regex')` nas 30 tabelas

**Reflita:**

- Qual opção é mais fácil de manter?
- Qual opção documenta melhor a intenção?
- Qual opção você escolheria? Por quê?
- Em que situação você NÃO usaria domain?
- Se você usou domain e depois quer remover (voltar para VARCHAR simples), como faria?

**Sua reflexão:**

```
[ESCREVA SEUS PENSAMENTOS AQUI - MÍNIMO 7 LINHAS]
```

---

### Reflexão 2: NULL vs Valor Padrão

```sql
-- Opção A: Permitir NULL
CREATE TABLE produtos (
    estoque INTEGER  -- Pode ser NULL
);

-- Opção B: Usar valor padrão
CREATE TABLE produtos (
    estoque INTEGER DEFAULT 0  -- Nunca NULL, sempre 0 se não informado
);

-- Opção C: Proibir NULL sem padrão
CREATE TABLE produtos (
    estoque INTEGER NOT NULL  -- Obrigatório informar
);
```

**Reflita:**

- Qual a diferença entre `estoque = 0` e `estoque = NULL`?
- Em que situação NULL faz mais sentido que zero?
- Em que situação zero faz mais sentido que NULL?
- Como você decidiria entre permitir NULL, usar DEFAULT, ou exigir NOT NULL?
- E se o campo for `data_nascimento`? NULL faz sentido? Default faz sentido?

**Sua reflexão:**

```
[ESCREVA SEUS PENSAMENTOS AQUI - MÍNIMO 7 LINHAS]
```

---

### Reflexão 3: Integridade vs Performance

Constraints garantem integridade mas têm custo de performance:

- `CHECK` constraints são verificadas a cada INSERT/UPDATE
- `FOREIGN KEY` constraints exigem consulta na tabela referenciada
- `UNIQUE` constraints mantêm índices adicionais

**Cenário:** Sistema de logs que recebe 10.000 inserções por segundo.

**Reflita:**

- Vale a pena ter FOREIGN KEY se vai consultar tabela relacionada 10.000 vezes/segundo?
- Em que situações você abriria mão de constraints por performance?
- Como você garantiria integridade se não usar constraints do banco?
- Qual o risco de não ter constraints?
- Se você tivesse que escolher APENAS UMA constraint para manter, qual seria? Por quê?

**Sua reflexão:**

```
[ESCREVA SEUS PENSAMENTOS AQUI - MÍNIMO 7 LINHAS]
```

---

## 📤 Como Enviar Suas Respostas

1. Copie este arquivo
2. Preencha **TODOS** os exercícios e reflexões
3. Teste seu código SQL (se possível)
4. Envie para avaliação final

**Critérios de avaliação:**

- ✅ Domínio de domains e quando usar
- ✅ Uso correto de constraints (PK, FK, UNIQUE, CHECK, NOT NULL, EXCLUSION)
- ✅ Modelagem completa de sistema real
- ✅ Qualidade do código SQL
- ✅ Entendimento profundo de NULL
- ✅ Uso correto de COALESCE e NULLIF
- ✅ Queries complexas (JOINs, agregações, subqueries)
- ✅ Profundidade nas reflexões
- ✅ Pensamento crítico sobre trade-offs

---

## ⏱️ Tempo Estimado

- Exercício 1 (Domains): 30-40 min
- Exercício 2 (Modelagem): 60-80 min
- Exercício 3 (Populando): 40-50 min
- Exercício 4 (Queries): 50-70 min
- Exercício 5 (NULL): 30-40 min
- Exercício 6 (Constraints avançadas): 40-50 min
- Reflexões: 40-50 min
- **Total: 4,5-6 horas**

**IMPORTANTE:** Este é o exercício mais completo do curso! Reserve tempo adequado!

---

## 🎯 Esta É a Última Aula Antes da Avaliação!

Após completar estes exercícios, você terá demonstrado:

- ✅ Domínio completo do modelo relacional
- ✅ Capacidade de modelar sistemas complexos
- ✅ Habilidade avançada em SQL
- ✅ Compreensão profunda de integridade de dados
- ✅ Pensamento crítico sobre decisões de design

**Você está pronto para a avaliação final!** 🚀

Boa sorte! 💪
