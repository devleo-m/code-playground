# Aula 1: Introdução ao SQL e Bancos de Dados Relacionais

## O que é SQL?

SQL, que significa **Structured Query Language** (Linguagem de Consulta Estruturada), é uma linguagem de programação utilizada para comunicar-se e gerenciar bancos de dados. SQL é uma linguagem padrão para manipular dados armazenados em sistemas de gerenciamento de bancos de dados relacionais (RDBMS - Relational Database Management System) ou para processamento de fluxo de dados em sistemas de gerenciamento de fluxo de dados relacionais (RDSMS).

### Histórico

SQL foi desenvolvido pela primeira vez na década de 1970 pela IBM, como parte do projeto System R. Desde então, tornou-se o padrão de fato para trabalhar com bancos de dados relacionais, sendo adotado por praticamente todos os principais sistemas de banco de dados do mercado.

## Componentes do SQL

SQL é composto por vários componentes, cada um servindo a um propósito único na comunicação com bancos de dados:

### 1. Queries (Consultas)

Este é o componente que permite recuperar dados de um banco de dados. A instrução **SELECT** é a mais comumente usada para esse propósito.

**Exemplo básico:**
```sql
SELECT * FROM livros;
```

Esta query recupera todos os registros da tabela `livros`.

### 2. Data Definition Language (DDL) - Linguagem de Definição de Dados

O DDL permite criar, alterar ou excluir bancos de dados e seus objetos relacionados, como tabelas, views, índices, etc. Os comandos principais incluem:

- **CREATE**: Cria novos objetos no banco de dados
- **ALTER**: Modifica a estrutura de objetos existentes
- **DROP**: Remove objetos do banco de dados
- **TRUNCATE**: Remove todos os dados de uma tabela, mantendo sua estrutura

**Exemplo:**
```sql
CREATE TABLE livros (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    titulo TEXT NOT NULL,
    isbn TEXT UNIQUE,
    ano_publicacao INTEGER
);
```

### 3. Data Manipulation Language (DML) - Linguagem de Manipulação de Dados

O DML permite gerenciar dados dentro dos objetos do banco de dados. Esses comandos incluem:

- **SELECT**: Recupera dados (também considerado parte do DML)
- **INSERT**: Insere novos registros
- **UPDATE**: Atualiza registros existentes
- **DELETE**: Remove registros

**Exemplos:**
```sql
-- Inserir um novo livro
INSERT INTO livros (titulo, isbn, ano_publicacao) 
VALUES ('Dom Casmurro', '978-8535914841', 1899);

-- Atualizar um livro existente
UPDATE livros 
SET quantidade_disponivel = 10 
WHERE id = 1;

-- Deletar um livro
DELETE FROM livros 
WHERE id = 15;
```

### 4. Data Control Language (DCL) - Linguagem de Controle de Dados

O DCL inclui comandos como **GRANT** e **REVOKE**, que lidam principalmente com direitos, permissões e outras tarefas de gerenciamento de controle do sistema de banco de dados.

**Exemplo:**
```sql
-- Conceder permissão de leitura
GRANT SELECT ON livros TO usuario_leitura;

-- Revogar permissão
REVOKE SELECT ON livros FROM usuario_leitura;
```

## O que são Bancos de Dados Relacionais?

Bancos de dados relacionais são um tipo de sistema de gerenciamento de banco de dados (SGBD) que armazena e fornece acesso a pontos de dados que estão relacionados entre si. Baseados no modelo relacional introduzido por E.F. Codd em 1970, eles usam uma estrutura que permite que os dados sejam organizados em tabelas com linhas e colunas.

### Características Principais

1. **Uso de SQL**: Utilizam SQL (Structured Query Language) para consultar e gerenciar dados

2. **Suporte a Transações ACID**: 
   - **Atomicity (Atomicidade)**: Uma transação é uma unidade indivisível - ou todas as operações são executadas ou nenhuma
   - **Consistency (Consistência)**: O banco sempre permanece em um estado válido
   - **Isolation (Isolamento)**: Transações concorrentes não interferem umas nas outras
   - **Durability (Durabilidade)**: Mudanças confirmadas são permanentes

3. **Integridade de Dados**: Através de constraints (restrições) como:
   - Chaves primárias (PRIMARY KEY)
   - Chaves estrangeiras (FOREIGN KEY)
   - Valores únicos (UNIQUE)
   - Validações (CHECK)

4. **Relacionamentos entre Tabelas**: Permitem estabelecer relacionamentos entre tabelas, possibilitando consultas complexas e recuperação de dados relacionais

5. **Escalabilidade e Multi-usuário**: Suportam ambientes com múltiplos usuários simultâneos

### Exemplos de Sistemas RDBMS Populares

- **MySQL**: Um dos mais populares, open-source
- **PostgreSQL**: Open-source, muito robusto e com recursos avançados
- **Oracle Database**: Solução empresarial robusta
- **Microsoft SQL Server**: Solução da Microsoft
- **SQLite**: Banco de dados leve, embutido (usado no nosso projeto)

### Estrutura de um Banco Relacional

No nosso banco de dados de biblioteca, temos:

```
categorias (1) ────< (N) livros
autores (1) ────< (N) livros
usuarios (1) ────< (N) emprestimos
livros (1) ────< (N) emprestimos
```

Isso significa:
- Um autor pode ter vários livros (relacionamento 1 para N)
- Uma categoria pode ter vários livros (relacionamento 1 para N)
- Um usuário pode ter vários empréstimos (relacionamento 1 para N)
- Um livro pode ser emprestado várias vezes (relacionamento 1 para N)

## Benefícios de um RDBMS

### 1. Dados Estruturados

RDBMS permite armazenar dados de forma estruturada, usando linhas e colunas em tabelas. Isso facilita a manipulação dos dados usando SQL, garantindo uso eficiente e flexível.

### 2. Propriedades ACID

As propriedades ACID garantem manipulação de dados confiável e segura em um RDBMS, tornando-o adequado para aplicações críticas onde a integridade dos dados é essencial.

### 3. Normalização

RDBMS suporta normalização de dados, um processo que organiza os dados de forma a reduzir redundância e melhorar a integridade dos dados.

**Exemplo de normalização:**
Em vez de armazenar o nome do autor repetidamente em cada livro:
```
❌ Não normalizado:
livros: [id, titulo, autor_nome, autor_nacionalidade, ...]
```

Fazemos:
```
✅ Normalizado:
autores: [id, nome, nacionalidade]
livros: [id, titulo, autor_id, ...]
```

### 4. Escalabilidade

RDBMSs geralmente fornecem boas opções de escalabilidade, permitindo a adição de mais armazenamento ou recursos computacionais conforme os dados e a carga de trabalho crescem.

### 5. Integridade de Dados

RDBMS fornece mecanismos como constraints, chaves primárias e chaves estrangeiras para impor integridade e consistência dos dados, garantindo que os dados sejam precisos e confiáveis.

**Exemplo:**
```sql
CREATE TABLE livros (
    id INTEGER PRIMARY KEY,
    autor_id INTEGER,
    FOREIGN KEY (autor_id) REFERENCES autores(id)
);
```

Isso garante que não possamos criar um livro com um `autor_id` que não existe na tabela `autores`.

### 6. Segurança

RDBMSs oferecem vários recursos de segurança, como autenticação de usuários, controle de acesso e criptografia de dados para proteger informações sensíveis.

## Limitações de um RDBMS

### 1. Complexidade

Configurar e gerenciar um RDBMS pode ser complexo, especialmente para aplicações grandes. Requer conhecimento técnico e habilidades para gerenciar, ajustar e otimizar o banco de dados.

### 2. Custo

RDBMSs podem ser caros, tanto em termos de taxas de licenciamento quanto dos recursos computacionais e de armazenamento que requerem. (Embora existam opções open-source gratuitas como MySQL e PostgreSQL)

### 3. Schema Fixo

RDBMS segue um esquema rígido para organização de dados, o que significa que quaisquer mudanças no esquema podem ser demoradas e complicadas.

**Exemplo:**
Se precisarmos adicionar um novo campo a uma tabela com milhões de registros, isso pode ser uma operação custosa:
```sql
ALTER TABLE livros ADD COLUMN novo_campo TEXT;
```

### 4. Dados Não Estruturados

RDBMSs não são adequados para lidar com dados não estruturados como arquivos multimídia, posts de redes sociais e dados de sensores, pois sua estrutura relacional é otimizada para dados estruturados.

### 5. Escalabilidade Horizontal

RDBMSs não são tão facilmente escaláveis horizontalmente quanto bancos de dados NoSQL. Escalar horizontalmente, que envolve adicionar mais máquinas ao sistema, pode ser desafiador em termos de custo e complexidade.

## SQL vs NoSQL

SQL (relacional) e NoSQL (não-relacional) representam duas abordagens diferentes para armazenamento e recuperação de dados.

### Diferenças Principais

| Aspecto | SQL (Relacional) | NoSQL (Não-Relacional) |
|---------|------------------|------------------------|
| **Estrutura** | Tabelas com linhas e colunas fixas | Documentos, chave-valor, grafos, colunas |
| **Schema** | Fixo e rígido | Flexível e dinâmico |
| **Escalabilidade** | Vertical (adicionar mais poder à máquina) | Horizontal (adicionar mais máquinas) |
| **Consistência** | Forte (ACID) | Variável (BASE - Basically Available, Soft state, Eventual consistency) |
| **Consultas** | SQL padronizado | APIs específicas por tipo |
| **Transações** | Suporte completo a transações complexas | Limitado ou inexistente |
| **Uso Ideal** | Dados estruturados, integridade crítica | Dados não estruturados, alta escala, flexibilidade |

### Quando Usar SQL?

- Dados estruturados e bem definidos
- Integridade de dados é crítica
- Necessidade de transações complexas
- Relacionamentos complexos entre dados
- Consultas ad-hoc complexas

### Quando Usar NoSQL?

- Dados não estruturados ou semi-estruturados
- Necessidade de escalabilidade horizontal massiva
- Schema em constante mudança
- Alta velocidade de escrita
- Dados que não precisam de relacionamentos complexos

### Exemplos de Bancos NoSQL

- **MongoDB**: Banco de documentos
- **Redis**: Banco chave-valor em memória
- **Cassandra**: Banco de colunas
- **Neo4j**: Banco de grafos

## Conclusão

SQL e bancos de dados relacionais são fundamentais para o desenvolvimento de software moderno. Eles fornecem uma base sólida, confiável e padronizada para armazenar e gerenciar dados estruturados. Embora existam alternativas NoSQL para casos específicos, o conhecimento de SQL continua sendo essencial para qualquer desenvolvedor ou analista de dados.

No nosso projeto, utilizaremos SQLite, que é um RDBMS leve e embutido, perfeito para aprendizado e aplicações pequenas a médias. Todas as técnicas que aprenderemos aqui são transferíveis para outros sistemas RDBMS como MySQL, PostgreSQL, etc.

---

**Próximo Passo**: Na próxima seção, vamos simplificar esses conceitos usando analogias do dia a dia para facilitar o entendimento.

