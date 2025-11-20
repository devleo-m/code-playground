# Aula 26: ORMs & DB Access - Parte 2: pgx (Driver PostgreSQL Nativo)

Olá, futuro(a) Gopher!

Bem-vindo(a) à segunda parte da aula 26! Na parte anterior, aprendemos sobre `database/sql` e como trabalhar com bancos de dados usando a biblioteca padrão. Agora vamos conhecer o **pgx**, um driver PostgreSQL nativo escrito em Go puro que oferece recursos avançados e melhor performance.

---

## Revisão da Parte 1

Na primeira parte, aprendemos:
- Diferentes abordagens de acesso a banco (SQL puro, ORMs, Query Builders)
- Como usar `database/sql` da biblioteca padrão
- Queries, comandos, transações e tratamento de erros
- Quando usar cada abordagem

---

## O que é pgx?

**pgx** é um driver PostgreSQL escrito em Go puro (sem dependências C). Ele oferece duas interfaces:

1. **Compatibilidade com database/sql**: Pode ser usado como drop-in replacement
2. **API Nativa**: Interface própria com recursos avançados do PostgreSQL

### Por que pgx?

- **Performance**: Mais rápido que `lib/pq` (driver tradicional)
- **Recursos PostgreSQL**: Suporte a recursos específicos do PostgreSQL
- **Type Safety**: Melhor suporte a tipos do PostgreSQL
- **Arrays e JSON**: Suporte nativo a arrays e tipos JSON
- **LISTEN/NOTIFY**: Suporte a notificações do PostgreSQL
- **Copy Protocol**: Suporte ao protocolo COPY para importação/exportação rápida
- **Go Puro**: Sem dependências C, compila facilmente

---

## Instalação

```bash
go get github.com/jackc/pgx/v5
go get github.com/jackc/pgx/v5/pgxpool
go get github.com/jackc/pgx/v5/stdlib
```

**Pacotes principais:**
- `pgx/v5`: API nativa do pgx
- `pgx/v5/pgxpool`: Pool de conexões
- `pgx/v5/stdlib`: Compatibilidade com `database/sql`

---

## Modo 1: Compatibilidade com database/sql

pgx pode ser usado como substituto direto do `lib/pq`:

```go
package main

import (
    "database/sql"
    "fmt"
    "log"
    
    _ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
    dsn := "postgres://usuario:senha@localhost/meubanco?sslmode=disable"
    
    db, err := sql.Open("pgx", dsn)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()
    
    if err := db.Ping(); err != nil {
        log.Fatal(err)
    }
    
    fmt.Println("Conectado com pgx via database/sql!")
}
```

**Vantagem**: Você pode migrar de `lib/pq` para `pgx` apenas mudando o driver, sem alterar o código!

---

## Modo 2: API Nativa (Recomendado)

A API nativa do pgx oferece mais recursos e melhor performance:

### Conexão Simples

```go
package main

import (
    "context"
    "fmt"
    "log"
    
    "github.com/jackc/pgx/v5"
)

func main() {
    dsn := "postgres://usuario:senha@localhost/meubanco?sslmode=disable"
    
    conn, err := pgx.Connect(context.Background(), dsn)
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Close(context.Background())
    
    fmt.Println("Conectado com pgx nativo!")
}
```

### Pool de Conexões (Recomendado para Produção)

```go
package main

import (
    "context"
    "fmt"
    "log"
    "time"
    
    "github.com/jackc/pgx/v5/pgxpool"
)

func main() {
    dsn := "postgres://usuario:senha@localhost/meubanco?sslmode=disable"
    
    config, err := pgxpool.ParseConfig(dsn)
    if err != nil {
        log.Fatal(err)
    }
    
    // Configurar pool
    config.MaxConns = 25
    config.MinConns = 5
    config.MaxConnLifetime = time.Hour
    config.MaxConnIdleTime = time.Minute * 30
    config.HealthCheckPeriod = time.Minute
    
    pool, err := pgxpool.NewWithConfig(context.Background(), config)
    if err != nil {
        log.Fatal(err)
    }
    defer pool.Close()
    
    // Verificar conexão
    if err := pool.Ping(context.Background()); err != nil {
        log.Fatal(err)
    }
    
    fmt.Println("Pool de conexões criado!")
}
```

---

## Queries com API Nativa

### Query Simples

```go
type User struct {
    ID    int32
    Name  string
    Email string
}

func getUsers(pool *pgxpool.Pool) ([]User, error) {
    rows, err := pool.Query(context.Background(),
        "SELECT id, name, email FROM users")
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    
    var users []User
    for rows.Next() {
        var u User
        if err := rows.Scan(&u.ID, &u.Name, &u.Email); err != nil {
            return nil, err
        }
        users = append(users, u)
    }
    
    return users, rows.Err()
}
```

### QueryRow

```go
func getUserByID(pool *pgxpool.Pool, id int32) (*User, error) {
    var u User
    err := pool.QueryRow(context.Background(),
        "SELECT id, name, email FROM users WHERE id = $1",
        id,
    ).Scan(&u.ID, &u.Name, &u.Email)
    
    if err == pgx.ErrNoRows {
        return nil, fmt.Errorf("usuário %d não encontrado", id)
    }
    if err != nil {
        return nil, err
    }
    
    return &u, nil
}
```

### Exec (INSERT, UPDATE, DELETE)

```go
func createUser(pool *pgxpool.Pool, name, email string) (int32, error) {
    var id int32
    err := pool.QueryRow(context.Background(),
        `INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id`,
        name, email,
    ).Scan(&id)
    
    return id, err
}

func updateUser(pool *pgxpool.Pool, id int32, name, email string) error {
    tag, err := pool.Exec(context.Background(),
        "UPDATE users SET name = $1, email = $2 WHERE id = $3",
        name, email, id,
    )
    if err != nil {
        return err
    }
    
    if tag.RowsAffected() == 0 {
        return fmt.Errorf("usuário %d não encontrado", id)
    }
    
    return nil
}
```

---

## Tipos PostgreSQL Especiais

### Arrays

pgx tem suporte nativo a arrays do PostgreSQL:

```go
import "github.com/jackc/pgx/v5/pgtype"

func createPostWithTags(pool *pgxpool.Pool, title string, tags []string) error {
    // Converter []string para pgtype.TextArray
    var tagArray pgtype.TextArray
    if err := tagArray.Set(tags); err != nil {
        return err
    }
    
    _, err := pool.Exec(context.Background(),
        "INSERT INTO posts (title, tags) VALUES ($1, $2)",
        title, tagArray,
    )
    return err
}

func getPostWithTags(pool *pgxpool.Pool, id int32) (string, []string, error) {
    var title string
    var tagArray pgtype.TextArray
    
    err := pool.QueryRow(context.Background(),
        "SELECT title, tags FROM posts WHERE id = $1",
        id,
    ).Scan(&title, &tagArray)
    
    if err != nil {
        return "", nil, err
    }
    
    var tags []string
    if err := tagArray.AssignTo(&tags); err != nil {
        return "", nil, err
    }
    
    return title, tags, nil
}
```

### JSON/JSONB

```go
import (
    "encoding/json"
    "github.com/jackc/pgx/v5/pgtype"
)

type Metadata struct {
    Author string `json:"author"`
    Views  int    `json:"views"`
}

func createPostWithMetadata(pool *pgxpool.Pool, title string, meta Metadata) error {
    metaJSON, err := json.Marshal(meta)
    if err != nil {
        return err
    }
    
    _, err = pool.Exec(context.Background(),
        "INSERT INTO posts (title, metadata) VALUES ($1, $2::jsonb)",
        title, metaJSON,
    )
    return err
}

func getPostMetadata(pool *pgxpool.Pool, id int32) (*Metadata, error) {
    var metaJSON []byte
    err := pool.QueryRow(context.Background(),
        "SELECT metadata FROM posts WHERE id = $1",
        id,
    ).Scan(&metaJSON)
    
    if err != nil {
        return nil, err
    }
    
    var meta Metadata
    if err := json.Unmarshal(metaJSON, &meta); err != nil {
        return nil, err
    }
    
    return &meta, nil
}
```

### Tipos Numéricos

```go
import "github.com/jackc/pgx/v5/pgtype"

func handleDecimal(pool *pgxpool.Pool) error {
    var price pgtype.Numeric
    
    err := pool.QueryRow(context.Background(),
        "SELECT price FROM products WHERE id = $1",
        1,
    ).Scan(&price)
    
    if err != nil {
        return err
    }
    
    // Converter para float64
    var priceFloat float64
    if err := price.AssignTo(&priceFloat); err != nil {
        return err
    }
    
    fmt.Printf("Preço: %.2f\n", priceFloat)
    return nil
}
```

---

## Transações

### Transação Simples

```go
func transferMoney(pool *pgxpool.Pool, fromID, toID int32, amount float64) error {
    tx, err := pool.Begin(context.Background())
    if err != nil {
        return err
    }
    defer tx.Rollback(context.Background())
    
    // Debita
    _, err = tx.Exec(context.Background(),
        "UPDATE accounts SET balance = balance - $1 WHERE id = $2",
        amount, fromID,
    )
    if err != nil {
        return err
    }
    
    // Credita
    _, err = tx.Exec(context.Background(),
        "UPDATE accounts SET balance = balance + $1 WHERE id = $2",
        amount, toID,
    )
    if err != nil {
        return err
    }
    
    return tx.Commit(context.Background())
}
```

### Transação com Savepoints

```go
func complexOperation(pool *pgxpool.Pool) error {
    tx, err := pool.Begin(context.Background())
    if err != nil {
        return err
    }
    defer tx.Rollback(context.Background())
    
    // Operação 1
    _, err = tx.Exec(context.Background(), "INSERT INTO logs (message) VALUES ($1)", "Step 1")
    if err != nil {
        return err
    }
    
    // Criar savepoint
    _, err = tx.Exec(context.Background(), "SAVEPOINT sp1")
    if err != nil {
        return err
    }
    
    // Operação que pode falhar
    _, err = tx.Exec(context.Background(), "INSERT INTO logs (message) VALUES ($1)", "Step 2")
    if err != nil {
        // Rollback para savepoint
        tx.Exec(context.Background(), "ROLLBACK TO SAVEPOINT sp1")
        // Continua a transação
    }
    
    return tx.Commit(context.Background())
}
```

---

## LISTEN/NOTIFY

PostgreSQL suporta notificações pub/sub. pgx tem suporte nativo:

### Listener

```go
import "github.com/jackc/pgx/v5/pgxpool"

func setupListener(pool *pgxpool.Pool) error {
    conn, err := pool.Acquire(context.Background())
    if err != nil {
        return err
    }
    defer conn.Release()
    
    // Escutar canal
    _, err = conn.Exec(context.Background(), "LISTEN user_updates")
    if err != nil {
        return err
    }
    
    // Processar notificações
    go func() {
        for {
            notification, err := conn.Conn().WaitForNotification(context.Background())
            if err != nil {
                log.Printf("Erro ao receber notificação: %v", err)
                break
            }
            
            fmt.Printf("Notificação recebida: %s - %s\n", 
                notification.Channel, notification.Payload)
        }
    }()
    
    return nil
}
```

### Notify

```go
func notifyUserUpdate(pool *pgxpool.Pool, userID int32) error {
    _, err := pool.Exec(context.Background(),
        "NOTIFY user_updates, $1",
        fmt.Sprintf("User %d updated", userID),
    )
    return err
}
```

---

## COPY Protocol

O protocolo COPY do PostgreSQL permite importação/exportação muito rápida:

### Importação (COPY FROM)

```go
func importUsers(pool *pgxpool.Pool, users []User) error {
    conn, err := pool.Acquire(context.Background())
    if err != nil {
        return err
    }
    defer conn.Release()
    
    _, err = conn.CopyFrom(
        context.Background(),
        pgx.Identifier{"users"},
        []string{"name", "email"},
        pgx.CopyFromSlice(len(users), func(i int) ([]any, error) {
            return []any{users[i].Name, users[i].Email}, nil
        }),
    )
    
    return err
}
```

### Exportação (COPY TO)

```go
func exportUsers(pool *pgxpool.Pool) ([]User, error) {
    conn, err := pool.Acquire(context.Background())
    if err != nil {
        return nil, err
    }
    defer conn.Release()
    
    var users []User
    _, err = conn.CopyFrom(
        context.Background(),
        pgx.Identifier{"users"},
        []string{"id", "name", "email"},
        pgx.CopyToSlice(func(rows [][]any) error {
            for _, row := range rows {
                users = append(users, User{
                    ID:    row[0].(int32),
                    Name:  row[1].(string),
                    Email: row[2].(string),
                })
            }
            return nil
        }),
    )
    
    return users, err
}
```

---

## Batch Operations

pgx suporta batch operations para executar múltiplas queries de uma vez:

```go
func batchInsert(pool *pgxpool.Pool, users []User) error {
    batch := &pgx.Batch{}
    
    for _, u := range users {
        batch.Queue(
            "INSERT INTO users (name, email) VALUES ($1, $2)",
            u.Name, u.Email,
        )
    }
    
    results := pool.SendBatch(context.Background(), batch)
    defer results.Close()
    
    for i := 0; i < len(users); i++ {
        _, err := results.Exec()
        if err != nil {
            return err
        }
    }
    
    return results.Close()
}
```

---

## Exemplo Completo: Serviço com pgx

```go
package main

import (
    "context"
    "fmt"
    "log"
    "time"
    
    "github.com/jackc/pgx/v5/pgxpool"
)

type User struct {
    ID        int32
    Name      string
    Email     string
    CreatedAt time.Time
}

type UserService struct {
    pool *pgxpool.Pool
}

func NewUserService(dsn string) (*UserService, error) {
    config, err := pgxpool.ParseConfig(dsn)
    if err != nil {
        return nil, err
    }
    
    config.MaxConns = 25
    config.MinConns = 5
    
    pool, err := pgxpool.NewWithConfig(context.Background(), config)
    if err != nil {
        return nil, err
    }
    
    if err := pool.Ping(context.Background()); err != nil {
        return nil, err
    }
    
    return &UserService{pool: pool}, nil
}

func (s *UserService) Close() {
    s.pool.Close()
}

func (s *UserService) Create(ctx context.Context, name, email string) (*User, error) {
    var user User
    err := s.pool.QueryRow(ctx,
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

func (s *UserService) GetByID(ctx context.Context, id int32) (*User, error) {
    var user User
    err := s.pool.QueryRow(ctx,
        "SELECT id, name, email, created_at FROM users WHERE id = $1",
        id,
    ).Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt)
    
    if err == pgx.ErrNoRows {
        return nil, fmt.Errorf("usuário %d não encontrado", id)
    }
    if err != nil {
        return nil, fmt.Errorf("erro ao buscar usuário: %w", err)
    }
    
    return &user, nil
}

func (s *UserService) GetAll(ctx context.Context) ([]User, error) {
    rows, err := s.pool.Query(ctx,
        "SELECT id, name, email, created_at FROM users ORDER BY id")
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
    
    return users, rows.Err()
}

func main() {
    dsn := "postgres://usuario:senha@localhost/meubanco?sslmode=disable"
    
    service, err := NewUserService(dsn)
    if err != nil {
        log.Fatal(err)
    }
    defer service.Close()
    
    ctx := context.Background()
    
    // Criar usuário
    user, err := service.Create(ctx, "João Silva", "joao@example.com")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Usuário criado: %+v\n", user)
    
    // Buscar por ID
    found, err := service.GetByID(ctx, user.ID)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Usuário encontrado: %+v\n", found)
}
```

---

## Quando Usar pgx?

### ✅ Use pgx quando:

- Você usa PostgreSQL
- Precisa de melhor performance que `lib/pq`
- Quer usar recursos específicos do PostgreSQL (arrays, JSON, LISTEN/NOTIFY)
- Precisa de COPY protocol para importação/exportação rápida
- Quer type safety melhor
- Prefere Go puro (sem dependências C)

### ❌ Evite pgx quando:

- Precisa de compatibilidade com múltiplos bancos
- Já está usando `database/sql` e não precisa de recursos avançados
- Quer abstração completa (use ORM)

---

## Comparação: lib/pq vs pgx

| Característica | lib/pq | pgx |
|---------------|--------|-----|
| Performance | Boa | Melhor |
| Recursos PostgreSQL | Básico | Avançado |
| Arrays | Limitado | Nativo |
| JSON/JSONB | Básico | Nativo |
| LISTEN/NOTIFY | Sim | Sim (melhor) |
| COPY Protocol | Não | Sim |
| Dependências C | Sim | Não |
| API Nativa | Não | Sim |

---

## Próximos Passos

Nesta parte, aprendemos:
- O que é pgx e por que usá-lo
- Como usar pgx com `database/sql` e API nativa
- Recursos avançados: arrays, JSON, LISTEN/NOTIFY, COPY
- Batch operations e transações

Na **Parte 3**, vamos explorar o **GORM**, um ORM completo que acelera o desenvolvimento com abstrações poderosas.

Até lá, experimente usar pgx em seus projetos PostgreSQL!



