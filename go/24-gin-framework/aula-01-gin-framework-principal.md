# Aula 24: Frameworks Web - Gin

Olá, futuro(a) Gopher!

Bem-vindo(a) à aula 24! Na aula anterior, exploramos o pacote `net/http` da biblioteca padrão. Agora vamos conhecer o **Gin**, um framework HTTP web rápido e produtivo que facilita muito o desenvolvimento de APIs RESTful e aplicações web modernas.

Gin é um dos frameworks mais populares em Go, usado por muitas empresas e projetos de código aberto. Ele oferece uma API limpa e intuitiva, mantendo a performance alta que Go proporciona.

---

## Por que Usar um Framework?

Antes de mergulharmos no Gin, vamos entender por que usar um framework quando já temos `net/http`:

### Vantagens dos Frameworks

1. **Produtividade**: Menos código boilerplate, mais foco no negócio
2. **Validação Automática**: Validação de dados de entrada integrada
3. **Middleware Avançado**: Sistema de middleware mais flexível
4. **Routing Inteligente**: Roteamento mais poderoso e intuitivo
5. **JSON Binding**: Serialização/deserialização automática
6. **Estrutura Organizada**: Padrões e convenções estabelecidas
7. **Ecossistema**: Bibliotecas e plugins prontos

### Quando Usar Framework vs net/http?

**Use `net/http` quando:**
- Precisa de controle total
- Quer zero dependências
- Aplicação muito simples
- Performance máxima é crítica
- Quer entender tudo que acontece

**Use Gin (ou outro framework) quando:**
- Precisa de produtividade
- Quer validação automática
- Precisa de middleware complexo
- Está construindo APIs RESTful
- Quer conveniência sem perder muito performance

---

## Por que Gin?

Gin é uma excelente escolha por várias razões:

- **Performance**: Extremamente rápido, baseado em `httprouter`
- **Produtividade**: API limpa e intuitiva
- **Middleware**: Sistema de middleware poderoso e flexível
- **JSON Binding**: Validação e binding automático de JSON
- **Routing**: Sistema de roteamento eficiente com parâmetros
- **Documentação**: Bem documentado e com muitos exemplos
- **Comunidade**: Grande comunidade e ecossistema

### Benchmarks

Gin é um dos frameworks mais rápidos em Go, competindo diretamente com `net/http` em performance, mas oferecendo muito mais conveniência.

---

## Instalação

```bash
go get -u github.com/gin-gonic/gin
```

### Verificando a Instalação

```go
package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Gin instalado com sucesso!"})
	})
	r.Run(":8080")
}
```

---

## Conceitos Básicos

### Gin Context

O `gin.Context` é o coração do Gin. Ele encapsula a requisição HTTP e a resposta, fornecendo métodos convenientes para trabalhar com elas.

```go
func handler(c *gin.Context) {
	// c.Request - *http.Request original
	// c.Writer - http.ResponseWriter
	// Métodos convenientes do Gin
}
```

### Servidor Básico

```go
package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Olá, mundo!",
		})
	})
	
	r.Run(":8080")
}
```

**Executando:**
```bash
go run main.go
# Acesse http://localhost:8080
```

### gin.Default() vs gin.New()

```go
// gin.Default() - inclui middleware de logging e recovery
r := gin.Default()

// gin.New() - sem middlewares (mais controle)
r := gin.New()
r.Use(gin.Logger())
r.Use(gin.Recovery())
```

---

## Rotas

Gin oferece suporte completo a todos os métodos HTTP:

```go
func main() {
	r := gin.Default()
	
	r.GET("/users", getUsers)
	r.POST("/users", createUser)
	r.PUT("/users/:id", updateUser)
	r.DELETE("/users/:id", deleteUser)
	r.PATCH("/users/:id", patchUser)
	r.HEAD("/users", headUsers)
	r.OPTIONS("/users", optionsUsers)
	
	r.Run(":8080")
}
```

### Rotas com Múltiplos Métodos

```go
r.Any("/users", func(c *gin.Context) {
	// Responde a qualquer método HTTP
})

r.Handle("GET", "/users", getUsers)
r.Handle("POST", "/users", createUser)
```

### Parâmetros de Rota

#### Parâmetro Obrigatório

```go
r.GET("/users/:id", func(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{"id": id})
})
```

**Uso:** `GET /users/123` → `id = "123"`

#### Parâmetro Opcional

```go
r.GET("/users/:id/*action", func(c *gin.Context) {
	id := c.Param("id")
	action := c.Param("action")
	c.JSON(200, gin.H{
		"id":     id,
		"action": action,
	})
})
```

**Uso:** `GET /users/123/edit` → `id = "123"`, `action = "/edit"`

### Query Parameters

```go
r.GET("/search", func(c *gin.Context) {
	// Query obrigatório
	query := c.Query("q")
	
	// Query com valor padrão
	page := c.DefaultQuery("page", "1")
	
	// Query opcional (retorna string e bool)
	limit, exists := c.GetQuery("limit")
	if !exists {
		limit = "10"
	}
	
	c.JSON(200, gin.H{
		"query": query,
		"page":  page,
		"limit": limit,
	})
})
```

**Uso:** `GET /search?q=golang&page=2&limit=20`

### Query Array

```go
r.GET("/tags", func(c *gin.Context) {
	tags := c.QueryArray("tag")
	c.JSON(200, gin.H{"tags": tags})
})
```

**Uso:** `GET /tags?tag=go&tag=web&tag=api`

### Form Data

```go
r.POST("/form", func(c *gin.Context) {
	name := c.PostForm("name")
	email := c.DefaultPostForm("email", "default@example.com")
	
	c.JSON(200, gin.H{
		"name":  name,
		"email": email,
	})
})
```

### Multipart Form (Upload)

```go
r.POST("/upload", func(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	
	// Salvar arquivo
	dst := "./uploads/" + file.Filename
	if err := c.SaveUploadedFile(file, dst); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(200, gin.H{
		"message": "Arquivo enviado com sucesso",
		"file":    file.Filename,
		"size":    file.Size,
	})
})
```

### Múltiplos Arquivos

```go
r.POST("/upload-multiple", func(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	
	files := form.File["files"]
	var filenames []string
	
	for _, file := range files {
		dst := "./uploads/" + file.Filename
		if err := c.SaveUploadedFile(file, dst); err != nil {
			continue
		}
		filenames = append(filenames, file.Filename)
	}
	
	c.JSON(200, gin.H{"files": filenames})
})
```

---

## JSON Binding

Uma das maiores vantagens do Gin é o binding automático de JSON.

### Binding Básico

```go
type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
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

### Tipos de Binding

```go
// JSON
c.ShouldBindJSON(&user)

// Form
c.ShouldBind(&user)

// Query
c.ShouldBindQuery(&user)

// URI (parâmetros de rota)
c.ShouldBindUri(&user)

// Header
c.ShouldBindHeader(&user)

// YAML
c.ShouldBindYAML(&user)

// XML
c.ShouldBindXML(&user)
```

### Binding Automático

```go
r.POST("/users", func(c *gin.Context) {
	var user User
	// Tenta fazer bind automaticamente do JSON, Form, Query, etc.
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, user)
})
```

---

## Validação

Gin usa a biblioteca `validator` para validação automática através de tags.

### Tags de Validação Comuns

```go
type User struct {
	Name     string `json:"name" binding:"required,min=3,max=50"`
	Email    string `json:"email" binding:"required,email"`
	Age      int    `json:"age" binding:"required,gte=18,lte=100"`
	Password string `json:"password" binding:"required,min=8"`
	Website  string `json:"website" binding:"omitempty,url"`
}
```

### Validações Disponíveis

```go
// Strings
required          // Campo obrigatório
min=3             // Mínimo de caracteres
max=50            // Máximo de caracteres
len=10            // Exatamente N caracteres
email             // Formato de email válido
url               // URL válida
alphanum          // Apenas letras e números
alpha             // Apenas letras
numeric           // Apenas números

// Números
gte=18            // Maior ou igual a
lte=100           // Menor ou igual a
gt=0              // Maior que
lt=100            // Menor que
oneof=red green   // Um dos valores

// Outros
omitempty         // Valida apenas se não estiver vazio
```

### Validação Customizada

```go
import "github.com/go-playground/validator/v10"

var validate *validator.Validate

func validateEmail(fl validator.FieldLevel) bool {
	email := fl.Field().String()
	return strings.Contains(email, "@example.com")
}

func main() {
	validate = validator.New()
	validate.RegisterValidation("example_email", validateEmail)
	
	// Usar no struct
	type User struct {
		Email string `json:"email" binding:"required,example_email"`
	}
}
```

### Mensagens de Erro Customizadas

```go
r.POST("/users", func(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		var errors []string
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			for _, fieldError := range validationErrors {
				errors = append(errors, fmt.Sprintf(
					"Campo %s: %s", 
					fieldError.Field(), 
					fieldError.Tag(),
				))
			}
		}
		c.JSON(400, gin.H{"errors": errors})
		return
	}
	c.JSON(201, user)
})
```

---

## Respostas

Gin oferece vários métodos convenientes para enviar respostas:

### JSON

```go
c.JSON(200, gin.H{"message": "Sucesso"})
c.JSON(http.StatusOK, user)
```

### XML

```go
c.XML(200, gin.H{"message": "Sucesso"})
```

### YAML

```go
c.YAML(200, gin.H{"message": "Sucesso"})
```

### String

```go
c.String(200, "Texto simples")
```

### HTML

```go
c.HTML(200, "index.html", gin.H{
	"title": "Página Inicial",
})
```

### Arquivo

```go
c.File("./static/file.pdf")
```

### Download de Arquivo

```go
c.FileAttachment("./files/report.pdf", "report.pdf")
```

### Dados Binários

```go
c.Data(200, "application/octet-stream", data)
```

### Redirecionamento

```go
c.Redirect(301, "https://example.com")
```

### Headers Customizados

```go
c.Header("X-Custom-Header", "valor")
c.JSON(200, gin.H{"message": "Sucesso"})
```

---

## Middleware

Gin tem um sistema de middleware muito poderoso e flexível.

### Middleware Básico

```go
func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		
		// Processar request
		c.Next()
		
		// Log após processar
		duration := time.Since(start)
		fmt.Printf("%s %s - %v\n", 
			c.Request.Method, 
			c.Request.URL.Path, 
			duration)
	}
}

func main() {
	r := gin.Default()
	r.Use(LoggerMiddleware())
	r.GET("/", handler)
	r.Run(":8080")
}
```

### Middleware de Autenticação

```go
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(401, gin.H{"error": "Não autorizado"})
			c.Abort()
			return
		}
		
		// Validar token...
		c.Set("user_id", "123") // Armazenar no context
		c.Next()
	}
}
```

### Middleware de CORS

```go
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		
		c.Next()
	}
}
```

### Usando Valores do Context

```go
func handler(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(401, gin.H{"error": "Não autorizado"})
		return
	}
	
	c.JSON(200, gin.H{"user_id": userID})
}
```

### Abortando a Cadeia

```go
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !isAuthenticated(c) {
			c.AbortWithStatusJSON(401, gin.H{"error": "Não autorizado"})
			return
		}
		c.Next()
	}
}
```

### Grupos de Rotas com Middleware

```go
func main() {
	r := gin.Default()
	
	// Rotas públicas
	r.GET("/public", publicHandler)
	
	// Grupo de rotas com middleware
	api := r.Group("/api")
	api.Use(AuthMiddleware())
	{
		api.GET("/users", getUsers)
		api.POST("/users", createUser)
	}
	
	r.Run(":8080")
}
```

---

## Grupos de Rotas

Grupos permitem organizar rotas e aplicar middlewares específicos:

```go
func main() {
	r := gin.Default()
	
	// Grupo v1
	v1 := r.Group("/api/v1")
	{
		v1.GET("/users", getUsersV1)
		v1.POST("/users", createUserV1)
	}
	
	// Grupo v2 com middleware
	v2 := r.Group("/api/v2")
	v2.Use(AuthMiddleware())
	{
		v2.GET("/users", getUsersV2)
		v2.POST("/users", createUserV2)
	}
	
	// Grupo aninhado
	admin := r.Group("/admin")
	admin.Use(AdminMiddleware())
	{
		users := admin.Group("/users")
		{
			users.GET("", adminGetUsers)
			users.DELETE("/:id", adminDeleteUser)
		}
	}
	
	r.Run(":8080")
}
```

---

## Servindo Arquivos Estáticos

```go
func main() {
	r := gin.Default()
	
	// Servir arquivos estáticos de um diretório
	r.Static("/static", "./static")
	
	// Servir um arquivo específico
	r.StaticFile("/favicon.ico", "./favicon.ico")
	
	// Servir arquivos com prefixo
	r.StaticFS("/files", http.Dir("./uploads"))
	
	r.Run(":8080")
}
```

---

## Templates HTML

Gin suporta templates HTML:

```go
func main() {
	r := gin.Default()
	
	// Carregar templates
	r.LoadHTMLGlob("templates/*")
	// ou
	r.LoadHTMLFiles("templates/index.html", "templates/about.html")
	
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{
			"title": "Página Inicial",
		})
	})
	
	r.Run(":8080")
}
```

### Template com Layout

```go
r.LoadHTMLGlob("templates/**/*")

r.GET("/", func(c *gin.Context) {
	c.HTML(200, "pages/index.html", gin.H{
		"title": "Página Inicial",
	})
})
```

---

## Exemplo Completo: API REST com Gin

```go
package main

import (
	"net/http"
	"strconv"
	
	"github.com/gin-gonic/gin"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name" binding:"required,min=3"`
	Email string `json:"email" binding:"required,email"`
}

var users []User
var nextID = 1

func getUsers(c *gin.Context) {
	c.JSON(http.StatusOK, users)
}

func getUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	
	for _, user := range users {
		if user.ID == id {
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
	
	user.ID = nextID
	nextID++
	users = append(users, user)
	
	c.JSON(http.StatusCreated, user)
}

func updateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	
	var updatedUser User
	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	for i, user := range users {
		if user.ID == id {
			updatedUser.ID = id
			users[i] = updatedUser
			c.JSON(http.StatusOK, updatedUser)
			return
		}
	}
	
	c.JSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado"})
}

func deleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	
	for i, user := range users {
		if user.ID == id {
			users = append(users[:i], users[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Usuário deletado"})
			return
		}
	}
	
	c.JSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado"})
}

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		duration := time.Since(start)
		fmt.Printf("%s %s - %v\n", c.Request.Method, c.Request.URL.Path, duration)
	}
}

func main() {
	r := gin.Default()
	
	// Middleware global
	r.Use(LoggerMiddleware())
	
	// CORS
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Next()
	})
	
	// API v1
	api := r.Group("/api/v1")
	{
		api.GET("/users", getUsers)
		api.GET("/users/:id", getUser)
		api.POST("/users", createUser)
		api.PUT("/users/:id", updateUser)
		api.DELETE("/users/:id", deleteUser)
	}
	
	r.Run(":8080")
}
```

---

## Boas Práticas

### 1. **Organização de Código**
Separe handlers, models e middleware em arquivos diferentes:

```
project/
├── main.go
├── handlers/
│   ├── user.go
│   └── auth.go
├── models/
│   └── user.go
├── middleware/
│   ├── auth.go
│   └── logger.go
└── routes/
    └── routes.go
```

### 2. **Validação Sempre**
Sempre valide dados de entrada:

```go
if err := c.ShouldBindJSON(&user); err != nil {
	c.JSON(400, gin.H{"error": err.Error()})
	return
}
```

### 3. **Tratamento de Erros**
Trate erros adequadamente:

```go
if err != nil {
	c.JSON(500, gin.H{"error": "Erro interno"})
	return
}
```

### 4. **Status Codes Apropriados**
Use códigos de status HTTP apropriados:

```go
c.JSON(http.StatusCreated, user)      // 201
c.JSON(http.StatusOK, user)           // 200
c.JSON(http.StatusNotFound, ...)      // 404
c.JSON(http.StatusBadRequest, ...)     // 400
```

### 5. **Middleware Reutilizável**
Crie middlewares reutilizáveis:

```go
func AuthMiddleware() gin.HandlerFunc {
	// Implementação
}
```

### 6. **Logging**
Use logging estruturado:

```go
import "log/slog"

logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
logger.Info("Request", "method", c.Request.Method, "path", c.Request.URL.Path)
```

---

## Conclusão

Nesta aula, exploramos o framework Gin:

1. **Instalação e Conceitos Básicos**: Como começar com Gin
2. **Rotas e Parâmetros**: Sistema de roteamento poderoso
3. **JSON Binding**: Serialização/deserialização automática
4. **Validação**: Validação automática de dados
5. **Middleware**: Sistema flexível de middleware
6. **Grupos de Rotas**: Organização de rotas
7. **API REST Completa**: Exemplo prático

Gin oferece um equilíbrio perfeito entre produtividade e performance. Ele mantém a velocidade do Go enquanto oferece conveniências modernas para desenvolvimento web.

Na próxima aula, vamos explorar gRPC e Protocol Buffers para comunicação entre microsserviços!

Até lá, experimente criar sua própria API REST com Gin. Pratique usando validação, middleware e grupos de rotas!

