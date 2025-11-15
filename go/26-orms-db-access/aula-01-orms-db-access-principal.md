# Aula 26: ORMs & DB Access - Parte 1: Fundamentos e database/sql

Olá, futuro(a) Gopher!

Bem-vindo(a) à aula 26! Na aula anterior, exploramos gRPC e Protocol Buffers para comunicação entre microsserviços. Agora vamos mergulhar no mundo de **acesso a bancos de dados** em Go, um tópico fundamental para qualquer aplicação real.

Go oferece múltiplas abordagens para trabalhar com bancos de dados, desde SQL puro até ORMs completos. Nesta primeira parte, vamos entender os fundamentos e a biblioteca padrão `database/sql`.

---

## Revisão da Aula Anterior

Na aula 25, aprendemos sobre:
- **gRPC**: Framework RPC de alto desempenho
- **Protocol Buffers**: Serialização eficiente de dados
- Comunicação entre microsserviços
- Streaming bidirecional

Esses conceitos são importantes porque muitas aplicações que usam bancos de dados também precisam se comunicar com outros serviços!

---

## Por que Aprender Acesso a Bancos de Dados?

Quase todas as aplicações reais precisam armazenar e recuperar dados persistentes. Go oferece várias abordagens, cada uma com seus prós e contras:

1. **database/sql** (Biblioteca Padrão): SQL puro, máximo controle
2. **pgx**: Driver PostgreSQL nativo com recursos avançados
3. **GORM**: ORM completo para desenvolvimento rápido
4. **Query Builders**: Meio termo entre SQL puro e ORM

---

## Abordagens de Acesso a Banco de Dados

### 1. SQL Puro (database/sql)

**Características:**
- Controle total sobre as queries
- Performance máxima
- Sem abstrações, você escreve SQL diretamente
- Ideal para queries complexas e otimizadas

**Quando usar:**
- Performance é crítica
- Queries muito complexas
- Você quer entender exatamente o que está acontecendo
- Aplicações que precisam de controle fino

### 2. ORMs (Object-Relational Mapping)

**Características:**
- Abstração do banco de dados
- Mapeamento automático de structs para tabelas
- Migrations automáticas
- Menos código boilerplate

**Quando usar:**
- Desenvolvimento rápido (RAD)
- Aplicações com muitas operações CRUD simples
- Equipes que preferem trabalhar com objetos ao invés de SQL
- Prototipagem rápida

### 3. Query Builders

**Características:**
- Meio termo entre SQL puro e ORM
- Type-safe queries
- Mais legível que SQL puro
- Mais controle que ORMs

**Quando usar:**
- Quer type-safety sem perder controle
- Queries dinâmicas complexas
- Não quer escrever SQL manualmente, mas quer controle

---

## database/sql - A Biblioteca Padrão

O pacote `database/sql` é a interface padrão do Go para trabalhar com bancos de dados SQL. Ele fornece uma API genérica que funciona com qualquer driver SQL.

### Características Principais

1. **Interface Genérica**: Funciona com qualquer banco SQL (PostgreSQL, MySQL, SQLite, etc.)
2. **Connection Pooling**: Gerencia pool de conexões automaticamente
3. **Prepared Statements**: Suporte a prepared statements para segurança e performance
4. **Transactions**: Suporte completo a transações
5. **Context Support**: Integração com `context.Context` para timeouts e cancelamento

---

## Instalação e Setup

### Instalando um Driver

O `database/sql` é apenas uma interface. Você precisa instalar um driver específico para o banco que vai usar:

```bash
# PostgreSQL (lib/pq - driver tradicional)
go get github.com/lib/pq

# PostgreSQL (pgx - driver nativo, veremos na próxima parte)
go get github.com/jackc/pgx/v5

# MySQL
go get github.com/go-sql-driver/mysql

# SQLite
go get github.com/mattn/go-sqlite3
```

### Importando

```go
import (
    "database/sql"
    _ "github.com/lib/pq" // Importação anônima para registrar o driver
)
```

**Por que importação anônima?** O driver se registra automaticamente no `database/sql` quando importado. A importação anônima (`_`) é usada porque não usamos o pacote diretamente, apenas para seu efeito colateral (registro do driver).

---

## Conectando ao Banco de Dados

### String de Conexão

Cada banco tem seu formato de string de conexão:

```go
// PostgreSQL
dsn := "postgres://usuario:senha@localhost/nome_banco?sslmode=disable"

// MySQL
dsn := "usuario:senha@tcp(localhost:3306)/nome_banco"

// SQLite
dsn := "./banco.db"
```

### Abrindo Conexão

```go
package main

import (
    "database/sql"
    "fmt"
    "log"
    _ "github.com/lib/pq"
)

func main() {
    // String de conexão PostgreSQL
    dsn := "postgres://usuario:senha@localhost/meubanco?sslmode=disable"
    
    // Abre conexão (na verdade, abre um pool de conexões)
    db, err := sql.Open("postgres", dsn)
    if err != nil {
        log.Fatal("Erro ao abrir conexão:", err)
    }
    defer db.Close() // Importante: sempre feche!
    
    // Verifica se a conexão está funcionando
    if err := db.Ping(); err != nil {
        log.Fatal("Erro ao conectar:", err)
    }
    
    fmt.Println("Conectado com sucesso!")
}
```

**Importante:**
- `sql.Open()` não cria uma conexão real, apenas prepara o pool
- Use `db.Ping()` para verificar se a conexão funciona
- Sempre feche com `defer db.Close()`

---

## Configurando o Pool de Conexões

O `database/sql` gerencia um pool de conexões automaticamente. Você pode configurá-lo:

```go
db, err := sql.Open("postgres", dsn)
if err != nil {
    log.Fatal(err)
}
defer db.Close()

// Configurar pool
db.SetMaxOpenConns(25)        // Máximo de conexões abertas
db.SetMaxIdleConns(5)         // Máximo de conexões ociosas
db.SetConnMaxLifetime(5 * time.Minute) // Tempo máximo de vida de uma conexão
db.SetConnMaxIdleTime(10 * time.Minute) // Tempo máximo que uma conexão pode ficar ociosa
```

**Por que configurar?**
- **MaxOpenConns**: Evita abrir muitas conexões e sobrecarregar o banco
- **MaxIdleConns**: Mantém algumas conexões prontas para uso rápido
- **ConnMaxLifetime**: Evita usar conexões antigas que podem estar "quebradas"
- **ConnMaxIdleTime**: Libera conexões que não estão sendo usadas

---

## Executando Queries

### Query Simples (SELECT)

```go
type User struct {
    ID    int
    Name  string
    Email string
}

func getUsers(db *sql.DB) ([]User, error) {
    rows, err := db.Query("SELECT id, name, email FROM users")
    if err != nil {
        return nil, err
    }
    defer rows.Close() // Sempre feche!
    
    var users []User
    for rows.Next() {
        var u User
        if err := rows.Scan(&u.ID, &u.Name, &u.Email); err != nil {
            return nil, err
        }
        users = append(users, u)
    }
    
    // Verifica erros que podem ter ocorrido durante a iteração
    if err := rows.Err(); err != nil {
        return nil, err
    }
    
    return users, nil
}
```

**Pontos importantes:**
- Sempre use `defer rows.Close()`
- Use `rows.Next()` para iterar
- Use `rows.Scan()` para ler os valores
- Verifique `rows.Err()` após o loop

### Query com Parâmetros (Prepared Statement)

```go
func getUserByID(db *sql.DB, id int) (*User, error) {
    var u User
    err := db.QueryRow("SELECT id, name, email FROM users WHERE id = $1", id).
        Scan(&u.ID, &u.Name, &u.Email)
    
    if err == sql.ErrNoRows {
        return nil, fmt.Errorf("usuário %d não encontrado", id)
    }
    if err != nil {
        return nil, err
    }
    
    return &u, nil
}
```

**Por que usar parâmetros?**
- **Segurança**: Previne SQL Injection
- **Performance**: Prepared statements são mais rápidos
- **Legibilidade**: Código mais limpo

**Nota:** PostgreSQL usa `$1, $2, $3...`, MySQL usa `?`, SQLite aceita ambos.

### QueryRow para Uma Linha

```go
func getUserCount(db *sql.DB) (int, error) {
    var count int
    err := db.QueryRow("SELECT COUNT(*) FROM users").Scan(&count)
    if err != nil {
        return 0, err
    }
    return count, nil
}
```

`QueryRow` retorna apenas uma linha. Se não houver resultados, retorna `sql.ErrNoRows`.

---

## Executando Comandos (INSERT, UPDATE, DELETE)

### INSERT

```go
func createUser(db *sql.DB, name, email string) (int, error) {
    var id int
    err := db.QueryRow(
        "INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id",
        name, email,
    ).Scan(&id)
    
    if err != nil {
        return 0, err
    }
    
    return id, nil
}
```

**RETURNING**: PostgreSQL permite retornar valores após INSERT/UPDATE. MySQL usa `LastInsertId()`.

### UPDATE

```go
func updateUser(db *sql.DB, id int, name, email string) error {
    result, err := db.Exec(
        "UPDATE users SET name = $1, email = $2 WHERE id = $3",
        name, email, id,
    )
    if err != nil {
        return err
    }
    
    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return err
    }
    
    if rowsAffected == 0 {
        return fmt.Errorf("nenhum usuário encontrado com id %d", id)
    }
    
    return nil
}
```

### DELETE

```go
func deleteUser(db *sql.DB, id int) error {
    result, err := db.Exec("DELETE FROM users WHERE id = $1", id)
    if err != nil {
        return err
    }
    
    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return err
    }
    
    if rowsAffected == 0 {
        return fmt.Errorf("usuário %d não encontrado", id)
    }
    
    return nil
}
```

---

## Prepared Statements Explícitos

Para queries que serão executadas múltiplas vezes, você pode preparar o statement uma vez:

```go
func insertManyUsers(db *sql.DB, users []User) error {
    stmt, err := db.Prepare("INSERT INTO users (name, email) VALUES ($1, $2)")
    if err != nil {
        return err
    }
    defer stmt.Close()
    
    for _, u := range users {
        _, err := stmt.Exec(u.Name, u.Email)
        if err != nil {
            return err
        }
    }
    
    return nil
}
```

**Vantagens:**
- Mais eficiente para múltiplas execuções
- Validação da query acontece uma vez
- Melhor performance

---

## Transações

Transações garantem que múltiplas operações sejam executadas atomicamente (tudo ou nada).

### Transação Simples

```go
func transferMoney(db *sql.DB, fromID, toID int, amount float64) error {
    tx, err := db.Begin()
    if err != nil {
        return err
    }
    
    // Se algo der errado, faz rollback
    defer tx.Rollback()
    
    // Debita da conta origem
    _, err = tx.Exec(
        "UPDATE accounts SET balance = balance - $1 WHERE id = $2",
        amount, fromID,
    )
    if err != nil {
        return err
    }
    
    // Credita na conta destino
    _, err = tx.Exec(
        "UPDATE accounts SET balance = balance + $1 WHERE id = $2",
        amount, toID,
    )
    if err != nil {
        return err
    }
    
    // Se tudo deu certo, faz commit
    return tx.Commit()
}
```

### Transação com Context

```go
func transferMoneyWithContext(ctx context.Context, db *sql.DB, fromID, toID int, amount float64) error {
    tx, err := db.BeginTx(ctx, nil)
    if err != nil {
        return err
    }
    defer tx.Rollback()
    
    // Operações com contexto
    _, err = tx.ExecContext(ctx, "UPDATE accounts SET balance = balance - $1 WHERE id = $2", amount, fromID)
    if err != nil {
        return err
    }
    
    _, err = tx.ExecContext(ctx, "UPDATE accounts SET balance = balance + $1 WHERE id = $2", amount, toID)
    if err != nil {
        return err
    }
    
    return tx.Commit()
}
```

**Boas práticas:**
- Sempre use `defer tx.Rollback()` (é seguro chamar mesmo após commit)
- Use `BeginTx` com contexto para timeouts
- Operações dentro da transação devem usar `tx.Exec`, não `db.Exec`

---

## Trabalhando com NULL

Em Go, valores NULL do banco precisam ser tratados com tipos especiais:

```go
import "database/sql"

type User struct {
    ID       int
    Name     string
    Email    string
    Bio      sql.NullString  // String que pode ser NULL
    Age      sql.NullInt64   // Int que pode ser NULL
    Active   sql.NullBool    // Bool que pode ser NULL
}

func getUserWithNulls(db *sql.DB, id int) (*User, error) {
    var u User
    err := db.QueryRow(
        "SELECT id, name, email, bio, age, active FROM users WHERE id = $1",
        id,
    ).Scan(&u.ID, &u.Name, &u.Email, &u.Bio, &u.Age, &u.Active)
    
    if err != nil {
        return nil, err
    }
    
    return &u, nil
}

// Usando valores NULL
func updateUserBio(db *sql.DB, id int, bio string) error {
    var bioValue sql.NullString
    if bio != "" {
        bioValue = sql.NullString{String: bio, Valid: true}
    } else {
        bioValue = sql.NullString{Valid: false}
    }
    
    _, err := db.Exec(
        "UPDATE users SET bio = $1 WHERE id = $2",
        bioValue, id,
    )
    return err
}
```

**Tipos NULL:**
- `sql.NullString`
- `sql.NullInt64`
- `sql.NullBool`
- `sql.NullFloat64`
- `sql.NullTime`

---

## Exemplo Completo: CRUD de Usuários

```go
package main

import (
    "database/sql"
    "fmt"
    "log"
    "time"
    
    _ "github.com/lib/pq"
)

type User struct {
    ID        int
    Name      string
    Email     string
    CreatedAt time.Time
}

type UserService struct {
    db *sql.DB
}

func NewUserService(dsn string) (*UserService, error) {
    db, err := sql.Open("postgres", dsn)
    if err != nil {
        return nil, err
    }
    
    if err := db.Ping(); err != nil {
        return nil, err
    }
    
    // Configurar pool
    db.SetMaxOpenConns(25)
    db.SetMaxIdleConns(5)
    db.SetConnMaxLifetime(5 * time.Minute)
    
    return &UserService{db: db}, nil
}

func (s *UserService) Close() error {
    return s.db.Close()
}

// Create
func (s *UserService) Create(name, email string) (*User, error) {
    var user User
    err := s.db.QueryRow(
        `INSERT INTO users (name, email, created_at) 
         VALUES ($1, $2, $3) 
         RETURNING id, name, email, created_at`,
        name, email, time.Now(),
    ).Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt)
    
    if err != nil {
        return nil, fmt.Errorf("erro ao criar usuário: %w", err)
    }
    
    return &user, nil
}

// Read
func (s *UserService) GetByID(id int) (*User, error) {
    var user User
    err := s.db.QueryRow(
        "SELECT id, name, email, created_at FROM users WHERE id = $1",
        id,
    ).Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt)
    
    if err == sql.ErrNoRows {
        return nil, fmt.Errorf("usuário %d não encontrado", id)
    }
    if err != nil {
        return nil, fmt.Errorf("erro ao buscar usuário: %w", err)
    }
    
    return &user, nil
}

func (s *UserService) GetAll() ([]User, error) {
    rows, err := s.db.Query("SELECT id, name, email, created_at FROM users ORDER BY id")
    if err != nil {
        return nil, fmt.Errorf("erro ao buscar usuários: %w", err)
    }
    defer rows.Close()
    
    var users []User
    for rows.Next() {
        var u User
        if err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.CreatedAt); err != nil {
            return nil, fmt.Errorf("erro ao escanear usuário: %w", err)
        }
        users = append(users, u)
    }
    
    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("erro durante iteração: %w", err)
    }
    
    return users, nil
}

// Update
func (s *UserService) Update(id int, name, email string) error {
    result, err := s.db.Exec(
        "UPDATE users SET name = $1, email = $2 WHERE id = $3",
        name, email, id,
    )
    if err != nil {
        return fmt.Errorf("erro ao atualizar: %w", err)
    }
    
    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return fmt.Errorf("erro ao verificar linhas afetadas: %w", err)
    }
    
    if rowsAffected == 0 {
        return fmt.Errorf("usuário %d não encontrado", id)
    }
    
    return nil
}

// Delete
func (s *UserService) Delete(id int) error {
    result, err := s.db.Exec("DELETE FROM users WHERE id = $1", id)
    if err != nil {
        return fmt.Errorf("erro ao deletar: %w", err)
    }
    
    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return fmt.Errorf("erro ao verificar linhas afetadas: %w", err)
    }
    
    if rowsAffected == 0 {
        return fmt.Errorf("usuário %d não encontrado", id)
    }
    
    return nil
}

func main() {
    dsn := "postgres://usuario:senha@localhost/meubanco?sslmode=disable"
    
    service, err := NewUserService(dsn)
    if err != nil {
        log.Fatal(err)
    }
    defer service.Close()
    
    // Criar usuário
    user, err := service.Create("João Silva", "joao@example.com")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Usuário criado: %+v\n", user)
    
    // Buscar por ID
    found, err := service.GetByID(user.ID)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Usuário encontrado: %+v\n", found)
    
    // Listar todos
    users, err := service.GetAll()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Total de usuários: %d\n", len(users))
}
```

---

## Context e Timeouts

Sempre use `context.Context` para controlar timeouts e cancelamento:

```go
import (
    "context"
    "time"
)

func getUserWithTimeout(db *sql.DB, id int) (*User, error) {
    // Timeout de 5 segundos
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    
    var user User
    err := db.QueryRowContext(ctx,
        "SELECT id, name, email FROM users WHERE id = $1",
        id,
    ).Scan(&user.ID, &user.Name, &user.Email)
    
    if err != nil {
        return nil, err
    }
    
    return &user, nil
}
```

**Por que usar context?**
- **Timeouts**: Evita queries que demoram muito
- **Cancelamento**: Permite cancelar operações longas
- **Tracing**: Integração com sistemas de observabilidade
- **Boas práticas**: Padrão em Go moderno

---

## Tratamento de Erros

### Erros Comuns

```go
func handleDBError(err error) {
    if err == sql.ErrNoRows {
        // Nenhuma linha encontrada
        fmt.Println("Registro não encontrado")
        return
    }
    
    if err == sql.ErrConnDone {
        // Conexão foi fechada
        fmt.Println("Conexão fechada")
        return
    }
    
    // Outros erros
    fmt.Printf("Erro de banco: %v\n", err)
}
```

### Wrapping de Erros

```go
import "fmt"

func getUser(db *sql.DB, id int) (*User, error) {
    var user User
    err := db.QueryRow("SELECT id, name FROM users WHERE id = $1", id).
        Scan(&user.ID, &user.Name)
    
    if err == sql.ErrNoRows {
        return nil, fmt.Errorf("usuário %d não encontrado: %w", id, err)
    }
    if err != nil {
        return nil, fmt.Errorf("erro ao buscar usuário %d: %w", id, err)
    }
    
    return &user, nil
}
```

---

## Quando Usar database/sql?

### ✅ Use database/sql quando:

- Performance é crítica
- Você precisa de controle total sobre as queries
- Queries são complexas e específicas
- Você quer entender exatamente o que está acontecendo
- Aplicação tem poucas operações de banco
- Você prefere SQL puro

### ❌ Evite database/sql quando:

- Muitas operações CRUD simples e repetitivas
- Você quer desenvolvimento rápido (RAD)
- Equipe prefere trabalhar com objetos
- Precisa de migrations automáticas
- Quer menos código boilerplate

---

## Próximos Passos

Nesta primeira parte, aprendemos:
- Diferentes abordagens de acesso a banco de dados
- Como usar `database/sql` da biblioteca padrão
- Queries, comandos e transações
- Tratamento de erros e context

Na **Parte 2**, vamos explorar o **pgx**, um driver PostgreSQL nativo que oferece recursos avançados e melhor performance.

Na **Parte 3**, vamos conhecer o **GORM**, um ORM completo que acelera o desenvolvimento.

Até lá, pratique criando suas próprias queries com `database/sql`!

