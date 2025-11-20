# Aula 26: ORMs & DB Access - Parte 3: GORM (ORM Completo)

Olá, futuro(a) Gopher!

Bem-vindo(a) à terceira e última parte da aula 26! Nas partes anteriores, aprendemos sobre `database/sql` e `pgx`. Agora vamos conhecer o **GORM**, um ORM (Object-Relational Mapping) completo e popular em Go que acelera o desenvolvimento através de abstrações poderosas.

---

## Revisão das Partes Anteriores

**Parte 1**: Aprendemos sobre `database/sql`, SQL puro, queries, transações e quando usar cada abordagem.

**Parte 2**: Exploramos o `pgx`, driver PostgreSQL nativo com recursos avançados como arrays, JSON, LISTEN/NOTIFY e COPY protocol.

---

## O que é GORM?

**GORM** (Go Object-Relational Mapping) é a biblioteca ORM mais popular em Go. Ele fornece:

- **Abstração de Banco**: Trabalhe com structs ao invés de SQL
- **Migrations Automáticas**: Cria e atualiza tabelas automaticamente
- **Associações**: Relacionamentos entre modelos (has one, has many, belongs to, many to many)
- **Hooks**: Callbacks antes/depois de operações
- **Query Builder**: API fluente para construir queries
- **Suporte Múltiplos Bancos**: PostgreSQL, MySQL, SQLite, SQL Server
- **Connection Pooling**: Gerenciamento automático de conexões

### Por que GORM?

- **Produtividade**: Desenvolvimento muito mais rápido
- **Menos Código**: Menos boilerplate que SQL puro
- **Type Safety**: Trabalha com structs tipadas
- **Convenções**: Segue convenções sensatas (mas customizáveis)
- **Ecossistema**: Muitos plugins e extensões

---

## Instalação

```bash
# GORM core
go get -u gorm.io/gorm

# Driver PostgreSQL
go get -u gorm.io/driver/postgres

# Driver MySQL
go get -u gorm.io/driver/mysql

# Driver SQLite
go get -u gorm.io/driver/sqlite
```

---

## Conectando ao Banco de Dados

### PostgreSQL

```go
package main

import (
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "log"
)

func main() {
    dsn := "host=localhost user=usuario password=senha dbname=meubanco port=5432 sslmode=disable"
    
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Erro ao conectar:", err)
    }
    
    sqlDB, err := db.DB()
    if err != nil {
        log.Fatal(err)
    }
    defer sqlDB.Close()
    
    log.Println("Conectado com sucesso!")
}
```

### MySQL

```go
import "gorm.io/driver/mysql"

dsn := "usuario:senha@tcp(localhost:3306)/meubanco?charset=utf8mb4&parseTime=True&loc=Local"
db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
```

### SQLite

```go
import "gorm.io/driver/sqlite"

db, err := gorm.Open(sqlite.Open("banco.db"), &gorm.Config{})
```

### Configurando o Pool

```go
sqlDB, err := db.DB()
if err != nil {
    log.Fatal(err)
}

// Configurar pool
sqlDB.SetMaxOpenConns(25)
sqlDB.SetMaxIdleConns(5)
sqlDB.SetConnMaxLifetime(time.Hour)
```

---

## Definindo Modelos

GORM usa structs para representar tabelas. Por convenção:
- Nome da struct = Nome da tabela (pluralizado, snake_case)
- Campos com `ID` ou `Id` = Primary Key
- Campos com `CreatedAt`, `UpdatedAt`, `DeletedAt` = Timestamps automáticos

### Modelo Básico

```go
type User struct {
    ID        uint      `gorm:"primaryKey"`
    Name      string    `gorm:"not null"`
    Email     string    `gorm:"uniqueIndex;not null"`
    Age       int
    CreatedAt time.Time
    UpdatedAt time.Time
}

// Nome da tabela (opcional, se quiser customizar)
func (User) TableName() string {
    return "users"
}
```

### Tags GORM Comuns

```go
type Product struct {
    ID          uint    `gorm:"primaryKey"`
    Name        string  `gorm:"type:varchar(100);not null"`
    Price       float64 `gorm:"type:decimal(10,2);not null"`
    Description string  `gorm:"type:text"`
    SKU         string  `gorm:"uniqueIndex;size:50"`
    Stock       int     `gorm:"default:0"`
    Active      bool    `gorm:"default:true"`
    CreatedAt   time.Time
    UpdatedAt   time.Time
}
```

**Tags importantes:**
- `primaryKey`: Chave primária
- `not null`: Campo obrigatório
- `uniqueIndex`: Índice único
- `index`: Índice normal
- `type:TIPO`: Tipo específico do banco
- `size:N`: Tamanho (para strings)
- `default:VALOR`: Valor padrão
- `autoIncrement`: Auto incremento

---

## Migrations Automáticas

GORM pode criar/atualizar tabelas automaticamente:

```go
func main() {
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal(err)
    }
    
    // Migrar (criar/atualizar tabelas)
    err = db.AutoMigrate(&User{}, &Product{})
    if err != nil {
        log.Fatal(err)
    }
    
    log.Println("Migração concluída!")
}
```

**O que `AutoMigrate` faz:**
- Cria tabelas que não existem
- Adiciona colunas faltantes
- Cria índices
- **NÃO** remove colunas ou índices não usados
- **NÃO** altera tipos de colunas existentes

**Para produção**, considere usar ferramentas de migration como `golang-migrate` ou `goose`.

---

## CRUD Básico

### Create (Criar)

```go
// Criar um registro
user := User{
    Name:  "João Silva",
    Email: "joao@example.com",
    Age:   30,
}

result := db.Create(&user)
if result.Error != nil {
    log.Fatal(result.Error)
}

fmt.Printf("Usuário criado com ID: %d\n", user.ID)

// Criar múltiplos
users := []User{
    {Name: "Maria", Email: "maria@example.com", Age: 25},
    {Name: "Pedro", Email: "pedro@example.com", Age: 35},
}

result = db.Create(&users)
if result.Error != nil {
    log.Fatal(result.Error)
}

// Criar com campos específicos
db.Select("Name", "Email").Create(&user)

// Criar ignorando campos específicos
db.Omit("Age").Create(&user)
```

### Read (Ler)

```go
// Buscar primeiro registro
var user User
result := db.First(&user)
if result.Error != nil {
    if errors.Is(result.Error, gorm.ErrRecordNotFound) {
        fmt.Println("Nenhum registro encontrado")
    } else {
        log.Fatal(result.Error)
    }
}

// Buscar por ID
var user User
db.First(&user, 1) // ID = 1

// Buscar último
db.Last(&user)

// Buscar um (sem erro se não encontrar)
db.Take(&user)

// Buscar todos
var users []User
db.Find(&users)

// Buscar com condições
var users []User
db.Where("age > ?", 25).Find(&users)

// Buscar com múltiplas condições
db.Where("name = ? AND age > ?", "João", 25).Find(&users)

// Buscar com IN
db.Where("name IN ?", []string{"João", "Maria"}).Find(&users)

// Buscar com LIKE
db.Where("email LIKE ?", "%@example.com").Find(&users)
```

### Update (Atualizar)

```go
// Atualizar um registro
user := User{ID: 1}
db.Model(&user).Update("name", "João Santos")

// Atualizar múltiplos campos
db.Model(&user).Updates(User{Name: "João Santos", Age: 31})

// Atualizar com map
db.Model(&user).Updates(map[string]interface{}{
    "name": "João Santos",
    "age":  31,
})

// Atualizar todos os registros que atendem condição
db.Model(&User{}).Where("age < ?", 18).Update("active", false)

// Atualizar campo específico
db.Model(&user).Select("name").Updates(User{Name: "João", Age: 31}) // Só atualiza name

// Atualizar ignorando campo
db.Model(&user).Omit("age").Updates(User{Name: "João", Age: 31}) // Atualiza tudo exceto age

// Save (atualiza todos os campos)
user.Name = "João Santos"
user.Age = 31
db.Save(&user)
```

### Delete (Deletar)

```go
// Deletar um registro
user := User{ID: 1}
db.Delete(&user)

// Deletar por ID
db.Delete(&User{}, 1)

// Deletar múltiplos
db.Delete(&User{}, []int{1, 2, 3})

// Deletar com condição
db.Where("age < ?", 18).Delete(&User{})

// Deletar permanente (ignora soft delete)
db.Unscoped().Delete(&User{}, 1)
```

---

## Soft Delete

GORM suporta soft delete (não remove fisicamente, apenas marca como deletado):

```go
type User struct {
    ID        uint
    Name      string
    Email     string
    DeletedAt gorm.DeletedAt `gorm:"index"` // Adiciona soft delete
}

// Deletar (soft delete)
db.Delete(&user) // Adiciona timestamp em DeletedAt

// Buscar incluindo deletados
db.Unscoped().Find(&users)

// Buscar apenas deletados
db.Unscoped().Where("deleted_at IS NOT NULL").Find(&users)

// Deletar permanentemente
db.Unscoped().Delete(&user)
```

---

## Query Builder Avançado

### Where Clauses

```go
// Where básico
db.Where("name = ?", "João").Find(&users)

// Where com struct
db.Where(&User{Name: "João", Age: 30}).Find(&users)

// Where com map
db.Where(map[string]interface{}{"name": "João", "age": 30}).Find(&users)

// Where com slice (IN)
db.Where("name IN ?", []string{"João", "Maria"}).Find(&users)

// Where com NOT
db.Not("name = ?", "João").Find(&users)

// Where com OR
db.Where("name = ?", "João").Or("name = ?", "Maria").Find(&users)

// Where com AND
db.Where("name = ?", "João").Where("age > ?", 25).Find(&users)
```

### Seleção de Campos

```go
// Selecionar campos específicos
db.Select("name", "email").Find(&users)

// Excluir campos
db.Omit("age", "created_at").Find(&users)
```

### Ordenação

```go
// Ordenar
db.Order("age desc").Find(&users)
db.Order("created_at asc").Find(&users)

// Múltiplas ordenações
db.Order("age desc").Order("name asc").Find(&users)
```

### Limite e Offset

```go
// Limite
db.Limit(10).Find(&users)

// Offset (paginação)
db.Offset(10).Limit(10).Find(&users) // Página 2

// Paginação simples
var users []User
page := 1
pageSize := 10
db.Offset((page - 1) * pageSize).Limit(pageSize).Find(&users)
```

### Agregações

```go
// Count
var count int64
db.Model(&User{}).Count(&count)

// Count com condição
db.Model(&User{}).Where("age > ?", 25).Count(&count)

// Sum
var total float64
db.Model(&Product{}).Select("sum(price)").Scan(&total)

// Avg, Max, Min
var avgAge float64
db.Model(&User{}).Select("avg(age)").Scan(&avgAge)

var maxAge int
db.Model(&User{}).Select("max(age)").Scan(&maxAge)
```

### Group By e Having

```go
type Result struct {
    Age  int
    Count int64
}

var results []Result
db.Model(&User{}).
    Select("age, count(*) as count").
    Group("age").
    Having("count(*) > ?", 5).
    Scan(&results)
```

### Joins

```go
type User struct {
    ID    uint
    Name  string
    Posts []Post
}

type Post struct {
    ID     uint
    Title  string
    UserID uint
    User   User
}

// Inner Join
var users []User
db.Joins("Posts").Find(&users)

// Left Join
db.Joins("LEFT JOIN posts ON posts.user_id = users.id").Find(&users)

// Join com condições
db.Joins("Posts", db.Where("posts.active = ?", true)).Find(&users)
```

---

## Associações (Relacionamentos)

### Has One (Tem Um)

```go
type User struct {
    ID    uint
    Name  string
    Profile Profile `gorm:"foreignKey:UserID"`
}

type Profile struct {
    ID     uint
    UserID uint
    Bio    string
    User   User
}

// Criar com associação
user := User{
    Name: "João",
    Profile: Profile{Bio: "Desenvolvedor Go"},
}
db.Create(&user)

// Buscar com associação
var user User
db.Preload("Profile").First(&user, 1)
```

### Has Many (Tem Muitos)

```go
type User struct {
    ID    uint
    Name  string
    Posts []Post `gorm:"foreignKey:UserID"`
}

type Post struct {
    ID     uint
    Title  string
    UserID uint
    User   User
}

// Criar com associação
user := User{
    Name: "João",
    Posts: []Post{
        {Title: "Post 1"},
        {Title: "Post 2"},
    },
}
db.Create(&user)

// Buscar com associação
var user User
db.Preload("Posts").First(&user, 1)
```

### Belongs To (Pertence A)

```go
type Post struct {
    ID     uint
    Title  string
    UserID uint
    User   User `gorm:"foreignKey:UserID"`
}

// Buscar com belongs to
var post Post
db.Preload("User").First(&post, 1)
```

### Many to Many (Muitos para Muitos)

```go
type User struct {
    ID    uint
    Name  string
    Roles []Role `gorm:"many2many:user_roles;"`
}

type Role struct {
    ID    uint
    Name  string
    Users []User `gorm:"many2many:user_roles;"`
}

// Criar relacionamento
user := User{Name: "João"}
role := Role{Name: "Admin"}
db.Create(&user)
db.Create(&role)

// Adicionar relacionamento
db.Model(&user).Association("Roles").Append(&role)

// Buscar com many to many
var user User
db.Preload("Roles").First(&user, 1)
```

---

## Hooks (Callbacks)

GORM permite executar código antes/depois de operações:

```go
type User struct {
    ID        uint
    Name      string
    Email     string
    Password  string
}

// BeforeCreate - Antes de criar
func (u *User) BeforeCreate(tx *gorm.DB) error {
    // Hash da senha antes de salvar
    hashedPassword, err := hashPassword(u.Password)
    if err != nil {
        return err
    }
    u.Password = hashedPassword
    return nil
}

// AfterCreate - Depois de criar
func (u *User) AfterCreate(tx *gorm.DB) error {
    // Enviar email de boas-vindas
    sendWelcomeEmail(u.Email)
    return nil
}

// BeforeUpdate - Antes de atualizar
func (u *User) BeforeUpdate(tx *gorm.DB) error {
    // Validar dados antes de atualizar
    if u.Email == "" {
        return errors.New("email é obrigatório")
    }
    return nil
}

// AfterFind - Depois de buscar
func (u *User) AfterFind(tx *gorm.DB) error {
    // Ocultar senha
    u.Password = ""
    return nil
}
```

**Hooks disponíveis:**
- `BeforeCreate`, `AfterCreate`
- `BeforeUpdate`, `AfterUpdate`
- `BeforeDelete`, `AfterDelete`
- `BeforeSave`, `AfterSave`
- `AfterFind`

---

## Transações

```go
// Transação simples
err := db.Transaction(func(tx *gorm.DB) error {
    if err := tx.Create(&user1).Error; err != nil {
        return err // Rollback automático
    }
    
    if err := tx.Create(&user2).Error; err != nil {
        return err // Rollback automático
    }
    
    return nil // Commit automático
})

// Transação manual
tx := db.Begin()
defer func() {
    if r := recover(); r != nil {
        tx.Rollback()
    }
}()

if err := tx.Create(&user1).Error; err != nil {
    tx.Rollback()
    return err
}

if err := tx.Create(&user2).Error; err != nil {
    tx.Rollback()
    return err
}

return tx.Commit().Error
```

---

## Scopes (Reutilização de Queries)

Scopes permitem reutilizar lógica de query:

```go
// Definir scope
func ActiveUsers(db *gorm.DB) *gorm.DB {
    return db.Where("active = ?", true)
}

func OlderThan(age int) func(db *gorm.DB) *gorm.DB {
    return func(db *gorm.DB) *gorm.DB {
        return db.Where("age > ?", age)
    }
}

// Usar scope
var users []User
db.Scopes(ActiveUsers, OlderThan(25)).Find(&users)
```

---

## Validação

GORM não tem validação built-in, mas funciona bem com `validator`:

```go
import "github.com/go-playground/validator/v10"

type User struct {
    ID       uint   `gorm:"primaryKey"`
    Name     string `gorm:"not null" validate:"required,min=3,max=50"`
    Email    string `gorm:"uniqueIndex" validate:"required,email"`
    Age      int    `validate:"gte=18,lte=100"`
}

// Validar antes de salvar
func (u *User) BeforeCreate(tx *gorm.DB) error {
    validate := validator.New()
    return validate.Struct(u)
}
```

---

## Exemplo Completo: API REST com GORM

```go
package main

import (
    "net/http"
    "strconv"
    
    "github.com/gin-gonic/gin"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

type User struct {
    ID        uint   `gorm:"primaryKey" json:"id"`
    Name      string `gorm:"not null" json:"name"`
    Email     string `gorm:"uniqueIndex;not null" json:"email"`
    Age       int    `json:"age"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

type UserService struct {
    db *gorm.DB
}

func NewUserService(dsn string) (*UserService, error) {
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        return nil, err
    }
    
    // Migrar
    if err := db.AutoMigrate(&User{}); err != nil {
        return nil, err
    }
    
    return &UserService{db: db}, nil
}

func (s *UserService) Create(c *gin.Context) {
    var user User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    
    if err := s.db.Create(&user).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    
    c.JSON(http.StatusCreated, user)
}

func (s *UserService) GetByID(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    
    var user User
    if err := s.db.First(&user, id).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            c.JSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado"})
            return
        }
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    
    c.JSON(http.StatusOK, user)
}

func (s *UserService) GetAll(c *gin.Context) {
    var users []User
    
    // Paginação
    page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
    pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
    offset := (page - 1) * pageSize
    
    // Filtros
    query := s.db.Model(&User{})
    
    if name := c.Query("name"); name != "" {
        query = query.Where("name LIKE ?", "%"+name+"%")
    }
    
    if age := c.Query("age"); age != "" {
        query = query.Where("age = ?", age)
    }
    
    // Buscar
    if err := query.Offset(offset).Limit(pageSize).Find(&users).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    
    c.JSON(http.StatusOK, gin.H{
        "data": users,
        "page": page,
        "page_size": pageSize,
    })
}

func (s *UserService) Update(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    
    var user User
    if err := s.db.First(&user, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado"})
        return
    }
    
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    
    if err := s.db.Save(&user).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    
    c.JSON(http.StatusOK, user)
}

func (s *UserService) Delete(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    
    if err := s.db.Delete(&User{}, id).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    
    c.JSON(http.StatusOK, gin.H{"message": "Usuário deletado"})
}

func main() {
    dsn := "host=localhost user=usuario password=senha dbname=meubanco port=5432 sslmode=disable"
    
    service, err := NewUserService(dsn)
    if err != nil {
        log.Fatal(err)
    }
    
    r := gin.Default()
    
    r.POST("/users", service.Create)
    r.GET("/users/:id", service.GetByID)
    r.GET("/users", service.GetAll)
    r.PUT("/users/:id", service.Update)
    r.DELETE("/users/:id", service.Delete)
    
    r.Run(":8080")
}
```

---

## Quando Usar GORM?

### ✅ Use GORM quando:

- Desenvolvimento rápido é prioridade
- Muitas operações CRUD simples
- Precisa de migrations automáticas
- Quer trabalhar com objetos ao invés de SQL
- Precisa de associações complexas
- Equipe prefere abstração

### ❌ Evite GORM quando:

- Performance é crítica (SQL puro é mais rápido)
- Queries são muito complexas e específicas
- Precisa de controle total sobre SQL
- Queries precisam ser otimizadas manualmente
- Aplicação tem poucas operações de banco

---

## Comparação: database/sql vs GORM

| Característica | database/sql | GORM |
|---------------|--------------|------|
| Performance | Máxima | Boa (overhead) |
| Produtividade | Baixa | Alta |
| Controle | Total | Limitado |
| Migrations | Manual | Automática |
| Associações | Manual | Automática |
| Código | Mais verboso | Menos código |
| Curva de Aprendizado | Baixa | Média |

---

## Próximos Passos

Nesta parte, aprendemos:
- O que é GORM e por que usá-lo
- Como definir modelos e fazer migrations
- Operações CRUD completas
- Associações e relacionamentos
- Hooks, transações e scopes
- Exemplo completo de API REST

---

## Conclusão da Aula 26

Ao longo desta aula, exploramos três abordagens principais para acesso a bancos de dados em Go:

1. **database/sql**: SQL puro, máximo controle e performance
2. **pgx**: Driver PostgreSQL nativo com recursos avançados
3. **GORM**: ORM completo para desenvolvimento rápido

Cada abordagem tem seu lugar:
- Use **database/sql** quando performance e controle são críticos
- Use **pgx** quando trabalha com PostgreSQL e precisa de recursos avançados
- Use **GORM** quando produtividade e desenvolvimento rápido são prioridades

A escolha depende do seu projeto, equipe e requisitos. Muitas vezes, você pode combinar abordagens: GORM para CRUD simples e SQL puro para queries complexas!

Na próxima aula, vamos explorar outros tópicos avançados de Go. Até lá, pratique criando aplicações com diferentes abordagens de acesso a banco de dados!



