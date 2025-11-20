# 📚 Projeto SQL - Banco de Dados Biblioteca

Este projeto contém um banco de dados SQLite de exemplo para prática e aprendizado de SQL.

## 🗄️ Estrutura do Banco de Dados

O banco de dados simula um sistema de biblioteca com as seguintes tabelas:

### Tabelas

1. **categorias** - Categorias dos livros (Ficção Científica, Romance, Técnico, etc.)
2. **autores** - Informações sobre os autores dos livros
3. **livros** - Catálogo de livros da biblioteca
4. **usuarios** - Usuários cadastrados na biblioteca
5. **emprestimos** - Registro de empréstimos de livros

### Relacionamentos

```
autores (1) ────< (N) livros
categorias (1) ────< (N) livros
usuarios (1) ────< (N) emprestimos
livros (1) ────< (N) emprestimos
```

## 🚀 Como Usar

### Pré-requisitos

- Go 1.21 ou superior
- Git (para baixar dependências)

### Inicialização

1. **Instalar dependências:**
   ```bash
   go mod download
   ```

2. **Criar o banco de dados:**
   ```bash
   go run init_database.go
   ```

   Isso criará o arquivo `biblioteca.db` com todas as tabelas e dados de exemplo.

### Usando o Banco de Dados

#### Opção 1: SQLite CLI

```bash
sqlite3 biblioteca.db
```

Depois você pode executar queries SQL diretamente:
```sql
SELECT * FROM livros;
SELECT * FROM autores;
```

#### Opção 2: Ferramentas Visuais

Você pode usar ferramentas como:
- **DB Browser for SQLite** (https://sqlitebrowser.org/)
- **DBeaver** (https://dbeaver.io/)
- **TablePlus** (https://tableplus.com/)

Basta abrir o arquivo `biblioteca.db` com uma dessas ferramentas.

#### Opção 3: Via Go

Você pode criar scripts Go para interagir com o banco:

```go
package main

import (
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
)

func main() {
    db, _ := sql.Open("sqlite3", "biblioteca.db")
    defer db.Close()
    
    // Suas queries aqui
}
```

## 📊 Dados de Exemplo

O banco já vem populado com:

- **6 categorias** (Ficção Científica, Romance, Técnico, História, Filosofia, Mistério)
- **10 autores** (incluindo brasileiros e internacionais)
- **15 livros** (com diferentes categorias e autores)
- **8 usuários** cadastrados
- **10 empréstimos** (alguns ativos, alguns devolvidos)

## 💡 Exemplos de Queries para Praticar

### Consultas Básicas

```sql
-- Listar todos os livros
SELECT * FROM livros;

-- Listar livros com seus autores
SELECT l.titulo, a.nome AS autor
FROM livros l
JOIN autores a ON l.autor_id = a.id;

-- Contar livros por categoria
SELECT c.nome, COUNT(l.id) AS total_livros
FROM categorias c
LEFT JOIN livros l ON c.id = l.categoria_id
GROUP BY c.id, c.nome;
```

### Consultas Intermediárias

```sql
-- Empréstimos ativos com informações do livro e usuário
SELECT 
    u.nome AS usuario,
    l.titulo AS livro,
    e.data_emprestimo,
    e.data_devolucao_prevista
FROM emprestimos e
JOIN usuarios u ON e.usuario_id = u.id
JOIN livros l ON e.livro_id = l.id
WHERE e.status = 'ativo';

-- Autores com mais livros
SELECT 
    a.nome,
    COUNT(l.id) AS total_livros
FROM autores a
LEFT JOIN livros l ON a.id = l.autor_id
GROUP BY a.id, a.nome
ORDER BY total_livros DESC;
```

### Consultas Avançadas

```sql
-- Usuários que nunca pegaram livros emprestados
SELECT u.nome, u.email
FROM usuarios u
LEFT JOIN emprestimos e ON u.id = e.usuario_id
WHERE e.id IS NULL;

-- Livros mais emprestados
SELECT 
    l.titulo,
    COUNT(e.id) AS vezes_emprestado
FROM livros l
LEFT JOIN emprestimos e ON l.id = e.livro_id
GROUP BY l.id, l.titulo
ORDER BY vezes_emprestado DESC;
```

## 🔄 Recriar o Banco de Dados

Se quiser recriar o banco do zero (apagando dados existentes):

```bash
go run init_database.go
```

O script automaticamente remove o banco existente antes de criar um novo.

## 📝 Notas

- O banco de dados é SQLite, então é um arquivo único (`biblioteca.db`)
- Todos os dados são fictícios e criados apenas para fins educacionais
- Sinta-se livre para modificar, adicionar ou remover dados conforme necessário para seus estudos

## 🎓 Próximos Passos

Agora você pode:
1. Explorar as tabelas e seus relacionamentos
2. Praticar queries SELECT básicas
3. Experimentar JOINs entre tabelas
4. Testar GROUP BY e funções de agregação
5. Criar suas próprias queries e análises

Bons estudos! 🚀


