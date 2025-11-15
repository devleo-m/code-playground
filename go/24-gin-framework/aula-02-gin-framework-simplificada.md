# Aula 24: Frameworks Web - Gin (Versão Simplificada)

Olá! Esta é uma versão resumida da aula 24 sobre o framework Gin.

---

## Por que Gin?

- ✅ **Performance** (baseado em httprouter)
- ✅ **Produtividade** (API limpa)
- ✅ **Validação automática**
- ✅ **Middleware poderoso**
- ✅ **JSON binding** automático

---

## 1. Instalação e Básico

```bash
go get -u github.com/gin-gonic/gin
```

```go
package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Olá!"})
	})
	r.Run(":8080")
}
```

---

## 2. Rotas

```go
r.GET("/users", getUsers)
r.POST("/users", createUser)
r.PUT("/users/:id", updateUser)
r.DELETE("/users/:id", deleteUser)
```

---

## 3. Parâmetros

### Rota

```go
r.GET("/users/:id", func(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{"id": id})
})
```

### Query

```go
r.GET("/search", func(c *gin.Context) {
	query := c.Query("q")
	page := c.DefaultQuery("page", "1")
	c.JSON(200, gin.H{"query": query, "page": page})
})
```

---

## 4. JSON Binding

```go
type User struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
}

r.POST("/users", func(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, user)
})
```

---

## 5. Validação

```go
type User struct {
	Name     string `json:"name" binding:"required,min=3,max=50"`
	Email    string `json:"email" binding:"required,email"`
	Age      int    `json:"age" binding:"required,gte=18"`
	Password string `json:"password" binding:"required,min=8"`
}
```

**Tags comuns:**
- `required` - Campo obrigatório
- `min=3` - Mínimo de caracteres
- `max=50` - Máximo de caracteres
- `email` - Formato de email
- `gte=18` - Maior ou igual a

---

## 6. Middleware

```go
func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		fmt.Printf("%s %s - %v\n", c.Request.Method, c.Request.URL.Path, time.Since(start))
	}
}

func main() {
	r := gin.Default()
	r.Use(LoggerMiddleware())
	r.GET("/", handler)
}
```

---

## 7. Grupos de Rotas

```go
api := r.Group("/api")
api.Use(AuthMiddleware())
{
	api.GET("/users", getUsers)
	api.POST("/users", createUser)
}
```

---

## 8. Upload de Arquivos

```go
r.POST("/upload", func(c *gin.Context) {
	file, _ := c.FormFile("file")
	c.SaveUploadedFile(file, "./uploads/"+file.Filename)
	c.JSON(200, gin.H{"message": "Arquivo enviado"})
})
```

---

## 9. Respostas

```go
c.JSON(200, gin.H{"message": "Sucesso"})
c.String(200, "Texto")
c.HTML(200, "index.html", gin.H{"title": "Página"})
c.File("./file.pdf")
c.Redirect(301, "https://example.com")
```

---

## 10. Arquivos Estáticos

```go
r.Static("/static", "./static")
r.StaticFile("/favicon.ico", "./favicon.ico")
```

---

## Resumo Rápido

| Tarefa | Código |
|--------|--------|
| Criar router | `r := gin.Default()` |
| Rota GET | `r.GET("/", handler)` |
| Parâmetro rota | `c.Param("id")` |
| Query param | `c.Query("q")` |
| JSON binding | `c.ShouldBindJSON(&user)` |
| Resposta JSON | `c.JSON(200, data)` |
| Middleware | `r.Use(middleware)` |
| Grupo | `r.Group("/api")` |

---

Pronto! Agora você tem o básico do Gin. Pratique criando uma API REST completa!

