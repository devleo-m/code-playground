package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ============================================================================
// EXEMPLO 1: database/sql - CRUD Básico
// ============================================================================

type UserSQL struct {
	ID        int
	Name      string
	Email     string
	CreatedAt time.Time
}

func exemploDatabaseSQL() {
	dsn := "postgres://usuario:senha@localhost/meubanco?sslmode=disable"
	
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	
	// Criar tabela
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users_sql (
			id SERIAL PRIMARY KEY,
			name VARCHAR(100) NOT NULL,
			email VARCHAR(100) UNIQUE NOT NULL,
			created_at TIMESTAMP DEFAULT NOW()
		)
	`)
	if err != nil {
		log.Fatal(err)
	}
	
	// CREATE
	user := UserSQL{
		Name:  "João Silva",
		Email: "joao@example.com",
	}
	
	err = db.QueryRow(
		"INSERT INTO users_sql (name, email) VALUES ($1, $2) RETURNING id, created_at",
		user.Name, user.Email,
	).Scan(&user.ID, &user.CreatedAt)
	
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Usuário criado: %+v\n", user)
	
	// READ
	var found UserSQL
	err = db.QueryRow(
		"SELECT id, name, email, created_at FROM users_sql WHERE id = $1",
		user.ID,
	).Scan(&found.ID, &found.Name, &found.Email, &found.CreatedAt)
	
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Usuário encontrado: %+v\n", found)
	
	// UPDATE
	_, err = db.Exec(
		"UPDATE users_sql SET name = $1 WHERE id = $2",
		"João Santos", user.ID,
	)
	if err != nil {
		log.Fatal(err)
	}
	
	// DELETE
	_, err = db.Exec("DELETE FROM users_sql WHERE id = $1", user.ID)
	if err != nil {
		log.Fatal(err)
	}
}

// ============================================================================
// EXEMPLO 2: database/sql - Transações
// ============================================================================

type Account struct {
	ID      int
	Balance float64
}

func transferMoney(db *sql.DB, fromID, toID int, amount float64) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	
	// Debita
	_, err = tx.Exec(
		"UPDATE accounts SET balance = balance - $1 WHERE id = $2",
		amount, fromID,
	)
	if err != nil {
		return err
	}
	
	// Credita
	_, err = tx.Exec(
		"UPDATE accounts SET balance = balance + $1 WHERE id = $2",
		amount, toID,
	)
	if err != nil {
		return err
	}
	
	return tx.Commit()
}

// ============================================================================
// EXEMPLO 3: pgx - API Nativa
// ============================================================================

type UserPgx struct {
	ID    int32
	Name  string
	Email string
	Tags  []string
}

func exemploPgx() {
	dsn := "postgres://usuario:senha@localhost/meubanco?sslmode=disable"
	
	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		log.Fatal(err)
	}
	
	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()
	
	// Criar tabela com array
	_, err = pool.Exec(context.Background(), `
		CREATE TABLE IF NOT EXISTS users_pgx (
			id SERIAL PRIMARY KEY,
			name VARCHAR(100) NOT NULL,
			email VARCHAR(100) UNIQUE NOT NULL,
			tags TEXT[]
		)
	`)
	if err != nil {
		log.Fatal(err)
	}
	
	// INSERT com array
	tags := []string{"golang", "backend", "developer"}
	var tagArray pgtype.TextArray
	if err := tagArray.Set(tags); err != nil {
		log.Fatal(err)
	}
	
	var id int32
	err = pool.QueryRow(context.Background(),
		"INSERT INTO users_pgx (name, email, tags) VALUES ($1, $2, $3) RETURNING id",
		"Maria Silva", "maria@example.com", tagArray,
	).Scan(&id)
	
	if err != nil {
		log.Fatal(err)
	}
	
	// SELECT com array
	var user UserPgx
	var tagArrayOut pgtype.TextArray
	err = pool.QueryRow(context.Background(),
		"SELECT id, name, email, tags FROM users_pgx WHERE id = $1",
		id,
	).Scan(&user.ID, &user.Name, &user.Email, &tagArrayOut)
	
	if err != nil {
		log.Fatal(err)
	}
	
	if err := tagArrayOut.AssignTo(&user.Tags); err != nil {
		log.Fatal(err)
	}
	
	fmt.Printf("Usuário com tags: %+v\n", user)
}

// ============================================================================
// EXEMPLO 4: pgx - JSON
// ============================================================================

type Metadata struct {
	Author string `json:"author"`
	Views  int    `json:"views"`
}

type Post struct {
	ID       int32
	Title    string
	Metadata Metadata
}

func exemploPgxJSON() {
	dsn := "postgres://usuario:senha@localhost/meubanco?sslmode=disable"
	
	pool, err := pgxpool.NewWithConfig(context.Background(), &pgxpool.Config{
		ConnConfig: pgx.ConnConfig{
			Config: pgx.Config{
				Host:     "localhost",
				User:     "usuario",
				Password: "senha",
				Database: "meubanco",
			},
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()
	
	// Criar tabela com JSONB
	_, err = pool.Exec(context.Background(), `
		CREATE TABLE IF NOT EXISTS posts (
			id SERIAL PRIMARY KEY,
			title VARCHAR(200) NOT NULL,
			metadata JSONB
		)
	`)
	if err != nil {
		log.Fatal(err)
	}
	
	// INSERT com JSON
	meta := Metadata{Author: "João", Views: 100}
	metaJSON, err := json.Marshal(meta)
	if err != nil {
		log.Fatal(err)
	}
	
	var id int32
	err = pool.QueryRow(context.Background(),
		"INSERT INTO posts (title, metadata) VALUES ($1, $2::jsonb) RETURNING id",
		"Meu Post", metaJSON,
	).Scan(&id)
	
	if err != nil {
		log.Fatal(err)
	}
	
	// SELECT com JSON
	var post Post
	var metaJSONOut []byte
	err = pool.QueryRow(context.Background(),
		"SELECT id, title, metadata FROM posts WHERE id = $1",
		id,
	).Scan(&post.ID, &post.Title, &metaJSONOut)
	
	if err != nil {
		log.Fatal(err)
	}
	
	if err := json.Unmarshal(metaJSONOut, &post.Metadata); err != nil {
		log.Fatal(err)
	}
	
	fmt.Printf("Post com metadata: %+v\n", post)
}

// ============================================================================
// EXEMPLO 5: GORM - Modelo Básico
// ============================================================================

type UserGORM struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"not null" json:"name"`
	Email     string    `gorm:"uniqueIndex;not null" json:"email"`
	Age       int       `json:"age"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func exemploGORM() {
	dsn := "host=localhost user=usuario password=senha dbname=meubanco port=5432 sslmode=disable"
	
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer sqlDB.Close()
	
	// Migrar
	if err := db.AutoMigrate(&UserGORM{}); err != nil {
		log.Fatal(err)
	}
	
	// CREATE
	user := UserGORM{
		Name:  "Pedro Santos",
		Email: "pedro@example.com",
		Age:   28,
	}
	
	if err := db.Create(&user).Error; err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Usuário criado: %+v\n", user)
	
	// READ
	var found UserGORM
	if err := db.First(&found, user.ID).Error; err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Usuário encontrado: %+v\n", found)
	
	// UPDATE
	found.Name = "Pedro Silva"
	if err := db.Save(&found).Error; err != nil {
		log.Fatal(err)
	}
	
	// DELETE
	if err := db.Delete(&found).Error; err != nil {
		log.Fatal(err)
	}
}

// ============================================================================
// EXEMPLO 6: GORM - Associações
// ============================================================================

type PostGORM struct {
	ID     uint   `gorm:"primaryKey"`
	Title  string `gorm:"not null"`
	Body   string
	UserID uint
	User   UserGORM `gorm:"foreignKey:UserID"`
}

func exemploGORMAssociacoes() {
	dsn := "host=localhost user=usuario password=senha dbname=meubanco port=5432 sslmode=disable"
	
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer sqlDB.Close()
	
	// Migrar
	if err := db.AutoMigrate(&UserGORM{}, &PostGORM{}); err != nil {
		log.Fatal(err)
	}
	
	// Criar usuário com posts
	user := UserGORM{
		Name:  "Ana Costa",
		Email: "ana@example.com",
		Age:   25,
		Posts: []PostGORM{
			{Title: "Post 1", Body: "Conteúdo do post 1"},
			{Title: "Post 2", Body: "Conteúdo do post 2"},
		},
	}
	
	if err := db.Create(&user).Error; err != nil {
		log.Fatal(err)
	}
	
	// Buscar com associação
	var userWithPosts UserGORM
	if err := db.Preload("Posts").First(&userWithPosts, user.ID).Error; err != nil {
		log.Fatal(err)
	}
	
	fmt.Printf("Usuário com posts: %+v\n", userWithPosts)
}

// ============================================================================
// EXEMPLO 7: GORM - Hooks
// ============================================================================

type UserWithHooks struct {
	ID        uint
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// BeforeCreate - Hash da senha antes de criar
func (u *UserWithHooks) BeforeCreate(tx *gorm.DB) error {
	// Simulação de hash (use bcrypt em produção)
	u.Password = "hashed_" + u.Password
	return nil
}

// AfterCreate - Log após criar
func (u *UserWithHooks) AfterCreate(tx *gorm.DB) error {
	fmt.Printf("Usuário %s criado com sucesso!\n", u.Name)
	return nil
}

// AfterFind - Ocultar senha após buscar
func (u *UserWithHooks) AfterFind(tx *gorm.DB) error {
	u.Password = ""
	return nil
}

func exemploGORMHooks() {
	dsn := "host=localhost user=usuario password=senha dbname=meubanco port=5432 sslmode=disable"
	
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer sqlDB.Close()
	
	if err := db.AutoMigrate(&UserWithHooks{}); err != nil {
		log.Fatal(err)
	}
	
	user := UserWithHooks{
		Name:     "Carlos",
		Email:    "carlos@example.com",
		Password: "senha123",
	}
	
	if err := db.Create(&user).Error; err != nil {
		log.Fatal(err)
	}
	
	// Buscar (senha será ocultada pelo AfterFind)
	var found UserWithHooks
	if err := db.First(&found, user.ID).Error; err != nil {
		log.Fatal(err)
	}
	
	fmt.Printf("Usuário (senha oculta): %+v\n", found)
}

// ============================================================================
// EXEMPLO 8: GORM - Validação
// ============================================================================

type UserValidated struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `gorm:"not null" validate:"required,min=3,max=50"`
	Email    string `gorm:"uniqueIndex" validate:"required,email"`
	Age      int    `validate:"gte=18,lte=100"`
}

func (u *UserValidated) BeforeCreate(tx *gorm.DB) error {
	validate := validator.New()
	if err := validate.Struct(u); err != nil {
		return err
	}
	return nil
}

func exemploGORMValidacao() {
	dsn := "host=localhost user=usuario password=senha dbname=meubanco port=5432 sslmode=disable"
	
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer sqlDB.Close()
	
	if err := db.AutoMigrate(&UserValidated{}); err != nil {
		log.Fatal(err)
	}
	
	// Tentar criar com dados inválidos
	user := UserValidated{
		Name:  "AB", // Muito curto
		Email: "email-invalido", // Email inválido
		Age:   15, // Menor que 18
	}
	
	if err := db.Create(&user).Error; err != nil {
		fmt.Printf("Erro de validação (esperado): %v\n", err)
	}
	
	// Criar com dados válidos
	user = UserValidated{
		Name:  "João Silva",
		Email: "joao@example.com",
		Age:   30,
	}
	
	if err := db.Create(&user).Error; err != nil {
		log.Fatal(err)
	}
	
	fmt.Printf("Usuário válido criado: %+v\n", user)
}

// ============================================================================
// EXEMPLO 9: GORM - Transações
// ============================================================================

func exemploGORMTransacoes(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		user1 := UserGORM{
			Name:  "User 1",
			Email: "user1@example.com",
			Age:   25,
		}
		
		if err := tx.Create(&user1).Error; err != nil {
			return err // Rollback automático
		}
		
		user2 := UserGORM{
			Name:  "User 2",
			Email: "user2@example.com",
			Age:   30,
		}
		
		if err := tx.Create(&user2).Error; err != nil {
			return err // Rollback automático
		}
		
		return nil // Commit automático
	})
}

// ============================================================================
// EXEMPLO 10: GORM - Query Builder Avançado
// ============================================================================

func exemploGORMQueryBuilder(db *gorm.DB) {
	var users []UserGORM
	
	// Where com múltiplas condições
	db.Where("age > ?", 25).
		Where("name LIKE ?", "%Silva%").
		Find(&users)
	
	// Ordenação e limite
	db.Order("age desc").
		Limit(10).
		Find(&users)
	
	// Paginação
	page := 1
	pageSize := 10
	db.Offset((page - 1) * pageSize).
		Limit(pageSize).
		Find(&users)
	
	// Agregações
	var count int64
	db.Model(&UserGORM{}).Count(&count)
	
	var avgAge float64
	db.Model(&UserGORM{}).Select("avg(age)").Scan(&avgAge)
}

// ============================================================================
// MAIN
// ============================================================================

func main() {
	fmt.Println("=== Exemplos de ORMs & DB Access em Go ===\n")
	
	// Descomente o exemplo que deseja executar:
	
	// exemploDatabaseSQL()
	// exemploPgx()
	// exemploPgxJSON()
	// exemploGORM()
	// exemploGORMAssociacoes()
	// exemploGORMHooks()
	// exemploGORMValidacao()
	
	fmt.Println("\n=== Fim dos exemplos ===")
}



