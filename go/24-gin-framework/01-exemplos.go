package main

// Este arquivo contém exemplos práticos do framework Gin
// Execute cada exemplo separadamente comentando/descomentando as funções main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// ============================================================================
// EXEMPLO 1: Servidor Básico com Gin
// ============================================================================

func exemplo1ServidorBasico() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Olá, mundo com Gin!",
		})
	})

	r.Run(":8080")
}

// ============================================================================
// EXEMPLO 2: Rotas e Parâmetros
// ============================================================================

func exemplo2Rotas() {
	r := gin.Default()

	// Rota simples
	r.GET("/hello", func(c *gin.Context) {
		c.String(200, "Olá!")
	})

	// Parâmetro de rota
	r.GET("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(200, gin.H{"id": id})
	})

	// Query parameters
	r.GET("/search", func(c *gin.Context) {
		query := c.Query("q")
		page := c.DefaultQuery("page", "1")
		c.JSON(200, gin.H{
			"query": query,
			"page":  page,
		})
	})

	r.Run(":8080")
}

// ============================================================================
// EXEMPLO 3: JSON Binding e Validação
// ============================================================================

type User struct {
	Name  string `json:"name" binding:"required,min=3"`
	Email string `json:"email" binding:"required,email"`
	Age   int    `json:"age" binding:"required,gte=18"`
}

func exemplo3Binding() {
	r := gin.Default()

	r.POST("/users", func(c *gin.Context) {
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		c.JSON(201, user)
	})

	r.Run(":8080")
}

// ============================================================================
// EXEMPLO 4: Middleware
// ============================================================================

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		duration := time.Since(start)
		println(c.Request.Method, c.Request.URL.Path, duration)
	}
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(401, gin.H{"error": "Não autorizado"})
			c.Abort()
			return
		}
		c.Set("user_id", "123")
		c.Next()
	}
}

func exemplo4Middleware() {
	r := gin.Default()

	// Middleware global
	r.Use(LoggerMiddleware())

	// Rota pública
	r.GET("/public", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Público"})
	})

	// Grupo com middleware
	api := r.Group("/api")
	api.Use(AuthMiddleware())
	{
		api.GET("/users", func(c *gin.Context) {
			userID, _ := c.Get("user_id")
			c.JSON(200, gin.H{"user_id": userID})
		})
	}

	r.Run(":8080")
}

// ============================================================================
// EXEMPLO 5: API REST Completa
// ============================================================================

var users []User
var nextID = 1

func getUsers(c *gin.Context) {
	c.JSON(http.StatusOK, users)
}

func getUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	for _, user := range users {
		if user.Name == strconv.Itoa(id) { // Simplificado para exemplo
			c.JSON(http.StatusOK, user)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado"})
}

func createUser(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user.Name = strconv.Itoa(nextID)
	nextID++
	users = append(users, user)
	c.JSON(http.StatusCreated, user)
}

func exemplo5APIREST() {
	r := gin.Default()

	api := r.Group("/api/v1")
	{
		api.GET("/users", getUsers)
		api.GET("/users/:id", getUser)
		api.POST("/users", createUser)
	}

	r.Run(":8080")
}

// ============================================================================
// EXEMPLO 6: Upload de Arquivos
// ============================================================================

func exemplo6Upload() {
	r := gin.Default()

	r.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		dst := "./uploads/" + file.Filename
		if err := c.SaveUploadedFile(file, dst); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{
			"message": "Arquivo enviado",
			"file":    file.Filename,
		})
	})

	r.Run(":8080")
}

// ============================================================================
// EXEMPLO 7: Grupos de Rotas
// ============================================================================

func exemplo7Grupos() {
	r := gin.Default()

	// Grupo v1
	v1 := r.Group("/api/v1")
	{
		v1.GET("/users", getUsers)
	}

	// Grupo v2
	v2 := r.Group("/api/v2")
	{
		v2.GET("/users", getUsers)
	}

	r.Run(":8080")
}

// ============================================================================
// MAIN - Descomente o exemplo que deseja executar
// ============================================================================

func main() {
	// Descomente o exemplo que deseja executar:

	// exemplo1ServidorBasico()
	// exemplo2Rotas()
	// exemplo3Binding()
	// exemplo4Middleware()
	// exemplo5APIREST()
	// exemplo6Upload()
	// exemplo7Grupos()

	println("Nenhum exemplo selecionado. Descomente um exemplo no main()")
}

