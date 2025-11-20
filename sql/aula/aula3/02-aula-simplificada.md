# Aula 3 - Simplificada: Entendendo DDL (Data Definition Language)

## DDL: Construindo e Reformando a Biblioteca

Imagine que você é o arquiteto de uma biblioteca. Enquanto os bibliotecários (DML) trabalham com os **livros** (dados) - organizando, emprestando, devolvendo - você trabalha com a **estrutura** da biblioteca - criando prateleiras, adicionando seções, reformando espaços.

**DDL é como ser o arquiteto do banco de dados**: você cria e modifica a estrutura onde os dados vão viver.

---

## CREATE TABLE: Construindo uma Nova Prateleira

### A Analogia da Prateleira

Pense em uma tabela como uma **prateleira organizada** na biblioteca. Antes de colocar livros nela, você precisa:

1. **Decidir o nome da prateleira** (nome da tabela)
2. **Definir os espaços** (colunas) que ela terá
3. **Decidir o tipo de coisa** que cada espaço guarda (tipo de dados)
4. **Criar regras** (constraints) para manter organização

### Exemplo Prático: Criando a Prateleira de Livros

```sql
CREATE TABLE livros (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    titulo TEXT NOT NULL,
    ano_publicacao INTEGER
);
```

**Pensando em português simples:**

"Vamos criar uma prateleira chamada 'livros' com três espaços:
- Um espaço para **id** (número único que identifica cada livro, que aumenta sozinho)
- Um espaço para **título** (texto, obrigatório - não pode ficar vazio)
- Um espaço para **ano de publicação** (número, opcional)"

### Constraints: As Regras da Prateleira

Constraints são como **regras de organização** que você coloca na prateleira:

#### PRIMARY KEY: O Identificador Único

Pense em PRIMARY KEY como um **número de registro único** que cada livro recebe. É como o código de barras - não pode haver dois livros com o mesmo código.

```
Prateleira de Livros:
┌────┬──────────────────┐
│ id │ titulo           │  ← id é a PRIMARY KEY
├────┼──────────────────┤
│ 1  │ Fundação         │  ← Livro #1
│ 2  │ Dom Casmurro     │  ← Livro #2
│ 3  │ 1984             │  ← Livro #3
└────┴──────────────────┘
```

**Analogia**: É como o número da sua identidade - único para cada pessoa.

#### NOT NULL: Campo Obrigatório

NOT NULL significa "este espaço **não pode ficar vazio**".

```sql
titulo TEXT NOT NULL
```

**Analogia**: É como um formulário onde alguns campos têm um asterisco (*) - você **precisa** preencher.

- ✅ **Com NOT NULL**: "Todo livro DEVE ter um título"
- ❌ **Sem NOT NULL**: "Um livro pode não ter título" (não faz sentido, certo?)

#### UNIQUE: Valor Único

UNIQUE significa "este valor não pode se repetir".

```sql
isbn TEXT UNIQUE
```

**Analogia**: É como o CPF - cada pessoa tem um único. Dois livros não podem ter o mesmo ISBN.

#### FOREIGN KEY: O "Apontador" para Outra Prateleira

FOREIGN KEY é como um **cartão de referência** que aponta para outra prateleira.

```sql
CREATE TABLE livros (
    id INTEGER PRIMARY KEY,
    autor_id INTEGER,
    FOREIGN KEY (autor_id) REFERENCES autores(id)
);
```

**Analogia Visual:**

```
Prateleira AUTORES:          Prateleira LIVROS:
┌────┬──────────────┐        ┌────┬──────────┬──────────┐
│ id │ nome        │        │ id │ titulo   │ autor_id │ ← aponta para
├────┼──────────────┤        ├────┼──────────┼──────────┤    autores.id
│ 1  │ Isaac Asimov│◄───────┤ 1  │ Fundação │ 1        │
│ 2  │ George Orwell│        │ 2  │ 1984     │ 2        │◄─── aponta
└────┴──────────────┘        └────┴──────────┴──────────┘
```

**Pensando em português**: "O campo `autor_id` no livro aponta para um autor que DEVE existir na prateleira de autores. Não posso criar um livro com `autor_id = 99` se não existir um autor com `id = 99`."

#### DEFAULT: Valor Padrão

DEFAULT é como ter um **valor pré-preenchido** que aparece automaticamente.

```sql
quantidade_disponivel INTEGER DEFAULT 0
```

**Analogia**: É como um formulário online que já vem com "Brasil" selecionado no campo "País" - se você não mudar, fica Brasil. Se não especificar quantidade, fica 0.

#### CHECK: Validação de Regras

CHECK é como ter um **fiscal** que verifica se o valor está dentro das regras.

```sql
nota INTEGER CHECK (nota >= 1 AND nota <= 5)
```

**Analogia**: É como uma máquina de venda que só aceita moedas de R$ 0,50, R$ 1,00 e R$ 2,00. Se você tentar colocar R$ 0,25, ela rejeita.

---

## ALTER TABLE: Reformando a Prateleira

### A Analogia da Reforma

ALTER TABLE é como **reformar uma prateleira existente** sem precisar jogar todos os livros fora.

### Adicionar Coluna: Adicionar um Novo Espaço

```sql
ALTER TABLE livros
ADD COLUMN preco REAL;
```

**Analogia**: É como adicionar uma nova gaveta na sua mesa. A mesa já existe, os objetos já estão lá, você só está adicionando um novo espaço.

**Antes:**
```
Prateleira LIVROS:
┌────┬──────────┬──────────┐
│ id │ titulo   │ autor_id │
├────┼──────────┼──────────┤
│ 1  │ Fundação │ 1        │
└────┴──────────┴──────────┘
```

**Depois:**
```
Prateleira LIVROS:
┌────┬──────────┬──────────┬────────┐
│ id │ titulo   │ autor_id │ preco  │ ← NOVO!
├────┼──────────┼──────────┼────────┤
│ 1  │ Fundação │ 1        │ NULL   │ ← vazio por enquanto
└────┴──────────┴──────────┴────────┘
```

### Renomear Coluna: Mudar o Nome do Espaço

```sql
ALTER TABLE livros
RENAME COLUMN quantidade_disponivel TO estoque;
```

**Analogia**: É como trocar a etiqueta de uma gaveta. O conteúdo não muda, só o nome.

**Antes**: Gaveta chamada "Quantidade Disponível"  
**Depois**: Gaveta chamada "Estoque"  
**Conteúdo**: Continua o mesmo!

### Remover Coluna: Retirar um Espaço

```sql
ALTER TABLE livros
DROP COLUMN editora;
```

**Analogia**: É como remover uma gaveta da mesa. **CUIDADO**: Tudo que estava nela será perdido!

**⚠️ Atenção**: Esta é uma operação **destrutiva**. É como jogar uma gaveta no lixo - você não consegue recuperar o que estava dentro depois.

---

## DROP TABLE: Demolindo a Prateleira

### A Analogia da Demolição

DROP TABLE é como **demolir completamente uma prateleira** - estrutura e tudo que está dentro.

```sql
DROP TABLE livros_temporarios;
```

**Analogia Visual:**

**Antes:**
```
Prateleira TEMPORÁRIA:
┌────┬──────────┐
│ id │ nome     │
├────┼──────────┤
│ 1  │ Livro A  │
│ 2  │ Livro B  │
└────┴──────────┘
```

**Depois:**
```
(nada - a prateleira não existe mais)
```

**⚠️ CRÍTICO**: É como demolir um prédio com tudo dentro. Não tem como desfazer (a menos que você tenha um backup)!

### DROP TABLE vs DELETE FROM

**DELETE FROM**: Remove apenas os **livros** (dados), mas a **prateleira** (estrutura) continua.

```sql
DELETE FROM livros WHERE id = 1;
```

**Analogia**: É como jogar um livro específico no lixo. A prateleira continua lá, vazia naquele espaço.

**DROP TABLE**: Remove a **prateleira inteira** (estrutura + dados).

```sql
DROP TABLE livros;
```

**Analogia**: É como demolir a prateleira inteira com todos os livros dentro.

---

## TRUNCATE TABLE: Esvaziando a Prateleira

### A Analogia da Limpeza Completa

TRUNCATE TABLE é como **esvaziar completamente uma prateleira**, mas **mantendo a prateleira** (estrutura).

**⚠️ No SQLite**: Não existe TRUNCATE. Use `DELETE FROM tabela;` que faz a mesma coisa.

```sql
DELETE FROM livros_temporarios;
```

**Analogia Visual:**

**Antes:**
```
Prateleira TEMPORÁRIA:
┌────┬──────────┐
│ id │ nome     │
├────┼──────────┤
│ 1  │ Livro A  │
│ 2  │ Livro B  │
│ 3  │ Livro C  │
└────┴──────────┘
```

**Depois:**
```
Prateleira TEMPORÁRIA:
┌────┬──────────┐
│ id │ nome     │  ← Estrutura mantida
├────┼──────────┤  ← Mas vazia!
│    │          │
└────┴──────────┘
```

**Pensando em português**: "Remova todos os livros, mas mantenha a prateleira pronta para receber novos livros."

### TRUNCATE vs DELETE: Qual a Diferença?

| Operação | O que remove | Estrutura | Velocidade |
|----------|--------------|-----------|------------|
| **DELETE FROM** | Apenas os livros | Mantém prateleira | Mais lento (remove um por um) |
| **TRUNCATE** | Todos os livros de uma vez | Mantém prateleira | Mais rápido (remove tudo de uma vez) |

**Analogia**: 
- **DELETE**: É como remover livro por livro da prateleira, um de cada vez
- **TRUNCATE**: É como virar a prateleira de cabeça para baixo e deixar todos os livros caírem de uma vez

---

## Índices: O Sistema de Busca Rápida

### A Analogia do Índice de Livro

Um **índice** no banco de dados é como o **índice no final de um livro** - ele te ajuda a encontrar informações rapidamente sem precisar ler tudo.

### Sem Índice: Busca Lenta

Imagine procurar a palavra "SQL" em um livro de 500 páginas **sem índice**:
- Você precisa ler página por página
- Pode levar muito tempo
- É trabalhoso

### Com Índice: Busca Rápida

Com um **índice alfabético** no final do livro:
- Você vai direto para a letra "S"
- Encontra "SQL" na página 245
- Vai direto lá!

### Criando um Índice

```sql
CREATE INDEX idx_livros_autor ON livros(autor_id);
```

**Analogia**: É como criar um "índice alfabético" na prateleira de livros, organizado por autor. Quando você quer encontrar todos os livros de "Isaac Asimov", o banco vai direto para eles, sem precisar verificar todos os livros.

### Quando Usar Índices?

**Use índices quando:**
- Você busca frequentemente por uma coluna (ex: buscar livros por autor)
- Você faz JOINs usando essa coluna
- Você ordena por essa coluna frequentemente

**Analogia**: Se você sempre busca livros por autor, faz sentido ter um "índice por autor" na prateleira.

**Não use índices demais quando:**
- Você adiciona/atualiza dados frequentemente
- A coluna tem poucos valores únicos

**Analogia**: Ter muitos índices é como ter muitos sistemas de organização diferentes na mesma prateleira - pode confundir e atrasar quando você quer adicionar novos livros.

---

## Exemplo Completo: Construindo uma Biblioteca do Zero

Vamos pensar em construir uma biblioteca completa usando analogias:

### Passo 1: Criar a Prateleira de Autores

```sql
CREATE TABLE autores (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    nome TEXT NOT NULL,
    nacionalidade TEXT
);
```

**Pensando em português**: "Vamos criar uma prateleira chamada 'autores' com espaços para id (número único), nome (obrigatório) e nacionalidade (opcional)."

### Passo 2: Criar a Prateleira de Livros

```sql
CREATE TABLE livros (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    titulo TEXT NOT NULL,
    autor_id INTEGER,
    FOREIGN KEY (autor_id) REFERENCES autores(id)
);
```

**Pensando em português**: "Vamos criar uma prateleira chamada 'livros' que 'aponta' para a prateleira de autores. Cada livro tem um autor_id que DEVE existir na prateleira de autores."

### Passo 3: Adicionar Sistema de Busca (Índices)

```sql
CREATE INDEX idx_livros_autor ON livros(autor_id);
```

**Pensando em português**: "Vamos criar um sistema de busca rápida para encontrar livros por autor."

### Passo 4: Reformar a Prateleira (Adicionar Campo)

```sql
ALTER TABLE livros
ADD COLUMN quantidade_disponivel INTEGER DEFAULT 0;
```

**Pensando em português**: "Vamos adicionar um novo espaço na prateleira de livros para guardar quantos exemplares temos disponíveis. Se não especificarmos, assume 0."

### Passo 5: Limpar Dados de Teste

```sql
DELETE FROM livros_temporarios;
```

**Pensando em português**: "Vamos esvaziar a prateleira temporária, mas mantê-la pronta para uso futuro."

---

## Resumo com Analogias

| Comando DDL | Analogia | O que faz |
|-------------|----------|-----------|
| **CREATE TABLE** | Construir nova prateleira | Cria estrutura nova |
| **ALTER TABLE ADD COLUMN** | Adicionar gaveta na mesa | Adiciona novo espaço |
| **ALTER TABLE RENAME COLUMN** | Trocar etiqueta da gaveta | Muda nome do espaço |
| **ALTER TABLE DROP COLUMN** | Remover gaveta | Remove espaço (destrutivo) |
| **DROP TABLE** | Demolir prateleira | Remove tudo (muito destrutivo) |
| **TRUNCATE/DELETE FROM** | Esvaziar prateleira | Remove dados, mantém estrutura |
| **CREATE INDEX** | Criar índice do livro | Sistema de busca rápida |

---

## Dicas Finais com Analogias

### 1. Pense Antes de Criar

**Antes de criar uma tabela**, pense: "Que informações preciso guardar? Como elas se relacionam?"

**Analogia**: Antes de construir uma prateleira, você planeja: "Quantas gavetas preciso? Que tamanho? Para que vou usar?"

### 2. Constraints São Seus Amigos

**Use constraints** para garantir que os dados estão corretos.

**Analogia**: É como ter regras na biblioteca: "Não pode ter dois livros com o mesmo ISBN" (UNIQUE), "Todo livro precisa de título" (NOT NULL).

### 3. Cuidado com Operações Destrutivas

**DROP TABLE e DROP COLUMN** são como demolir - não tem como desfazer facilmente.

**Analogia**: É como quebrar um prato - você pode colar, mas nunca será igual. Sempre faça backup!

### 4. Índices: Nem Muitos, Nem Poucos

**Crie índices** nas colunas que você busca frequentemente, mas não exagere.

**Analogia**: É como ter um sistema de organização na biblioteca - útil, mas ter muitos sistemas diferentes pode confundir.

---

## Conclusão

DDL é como ser o **arquiteto da biblioteca**:

- **CREATE TABLE**: Você constrói novas prateleiras (tabelas)
- **ALTER TABLE**: Você reforma prateleiras existentes
- **DROP TABLE**: Você remove prateleiras (cuidado!)
- **TRUNCATE/DELETE**: Você limpa prateleiras
- **CREATE INDEX**: Você cria sistemas de busca rápida

Lembre-se: você trabalha com a **estrutura**, enquanto DML trabalha com os **dados** dentro da estrutura.

**Próximo Passo**: Agora que você entendeu os conceitos de forma simplificada, vamos praticar com exercícios reais!

---

**💡 Dica**: Sempre pense nas analogias quando estiver criando ou modificando tabelas. "Estou construindo uma prateleira? Reformando? Ou demolindo?" Isso ajuda a entender o impacto de cada comando!


