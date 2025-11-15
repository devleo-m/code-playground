# Aula 24: Performance e Boas Práticas - Frameworks Web - Gin

Olá! Nesta parte da aula, vamos explorar boas práticas, otimizações de performance e padrões profissionais para desenvolvimento com Gin.

---

## Boas Práticas Gerais

### 1. Validação Robusta

**❌ Ruim:**
```go
r.POST("/users", func(c *gin.Context) {
	var user User
	c.ShouldBindJSON(&user)
	// Processar sem validar
})
```

**✅ Bom:**
```go
r.POST("/users", func(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{
			"error": "Dados inválidos",
			"details": formatValidationErrors(err),
		})
		return
	}
	
	if err := validateUser(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	
	// Processar
})
```

### 2. Tratamento de Erros Consistente

Crie um handler de erros centralizado:

```go
type ErrorResponse struct {
	Error   string            `json:"error"`
	Details map[string]string `json:"details,omitempty"`
}

func errorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		
		if len(c.Errors) > 0 {
			err := c.Errors.Last()
			
			switch err.Type {
			case gin.ErrorTypeBind:
				c.JSON(400, ErrorResponse{
					Error: "Erro de validação",
					Details: map[string]string{
						"message": err.Error(),
					},
				})
			case gin.ErrorTypePublic:
				c.JSON(500, ErrorResponse{
					Error: err.Error(),
				})
			default:
				c.JSON(500, ErrorResponse{
					Error: "Erro interno",
				})
			}
		}
	}
}

func main() {
	r := gin.Default()
	r.Use(errorHandler())
	// ... rotas
}
```

### 3. Middleware de Recuperação Customizado

Customize o middleware de recovery:

```go
func customRecovery() gin.HandlerFunc {
	return gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		log.Printf("Panic recuperado: %v", recovered)
		c.JSON(500, gin.H{
			"error": "Erro interno do servidor",
		})
		c.Abort()
	})
}

func main() {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(customRecovery())
	// ... rotas
}
```

### 4. Context para Timeouts

Use context para controlar timeouts:

```go
func timeoutMiddleware(timeout time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), timeout)
		defer cancel()
		
		c.Request = c.Request.WithContext(ctx)
		
		c.Next()
	}
}

func main() {
	r := gin.Default()
	r.Use(timeoutMiddleware(5 * time.Second))
	// ... rotas
}
```

---

## Performance

### 1. Modo Release

Use modo release em produção:

```go
func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	// ... rotas
}
```

### 2. Compressão de Respostas

Use middleware de compressão:

```go
import "github.com/gin-contrib/gzip"

func main() {
	r := gin.Default()
	r.Use(gzip.Gzip(gzip.DefaultCompression))
	// ... rotas
}
```

### 3. Cache de Respostas

Implemente cache:

```go
import "github.com/gin-contrib/cache"
import "time"

func main() {
	r := gin.Default()
	
	r.GET("/cacheable", cache.CachePage(store, time.Minute, func(c *gin.Context) {
		c.JSON(200, gin.H{"data": "cached"})
	}))
}
```

### 4. Connection Pooling

Configure o cliente HTTP adequadamente:

```go
import "net/http"

var httpClient = &http.Client{
	Timeout: 30 * time.Second,
	Transport: &http.Transport{
		MaxIdleConns:        100,
		MaxIdleConnsPerHost: 10,
		IdleConnTimeout:     90 * time.Second,
	},
}
```

### 5. Validação Customizada Eficiente

Registre validações customizadas uma vez:

```go
import "github.com/go-playground/validator/v10"

func setupValidator() *validator.Validate {
	validate := validator.New()
	
	validate.RegisterValidation("cpf", validateCPF)
	validate.RegisterValidation("cep", validateCEP)
	
	return validate
}

var validate = setupValidator()

type User struct {
	CPF string `json:"cpf" binding:"required,cpf"`
	CEP string `json:"cep" binding:"required,cep"`
}
```

---

## Segurança

### 1. CORS Configurado Corretamente

Configure CORS adequadamente:

```go
import "github.com/gin-contrib/cors"

func main() {
	r := gin.Default()
	
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://example.com"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
}
```

### 2. Rate Limiting

Implemente rate limiting:

```go
import "github.com/gin-contrib/limiter"

func main() {
	r := gin.Default()
	
	r.Use(limiter.Limit(limiter.Config{
		Max:        100,
		Duration:   time.Minute,
		Identifier: func(c *gin.Context) string {
			return c.ClientIP()
		},
	}))
}
```

### 3. Headers de Segurança

Adicione headers de segurança:

```go
func securityHeaders() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("X-Content-Type-Options", "nosniff")
		c.Header("X-Frame-Options", "DENY")
		c.Header("X-XSS-Protection", "1; mode=block")
		c.Header("Strict-Transport-Security", "max-age=31536000")
		c.Next()
	}
}
```

### 4. Validação de Upload

Valide uploads adequadamente:

```go
func uploadHandler(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(400, gin.H{"error": "Arquivo não fornecido"})
		return
	}
	
	// Validar tamanho
	if file.Size > 10<<20 { // 10MB
		c.JSON(400, gin.H{"error": "Arquivo muito grande"})
		return
	}
	
	// Validar tipo
	ext := filepath.Ext(file.Filename)
	allowed := []string{".jpg", ".png", ".pdf"}
	if !contains(allowed, ext) {
		c.JSON(400, gin.H{"error": "Tipo de arquivo não permitido"})
		return
	}
	
	// Processar
}
```

---

## Organização de Código

### 1. Estrutura de Projeto

Organize o projeto adequadamente:

```
project/
├── cmd/
│   └── server/
│       └── main.go
├── internal/
│   ├── handlers/
│   │   ├── user.go
│   │   └── product.go
│   ├── models/
│   │   └── user.go
│   ├── middleware/
│   │   ├── auth.go
│   │   └── logging.go
│   └── routes/
│       └── routes.go
├── pkg/
│   └── utils/
└── go.mod
```

### 2. Handlers Organizados

Separe handlers em arquivos:

```go
// handlers/user.go
type UserHandler struct {
	service UserService
}

func NewUserHandler(service UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) GetUser(c *gin.Context) {
	id := c.Param("id")
	user, err := h.service.GetUser(id)
	if err != nil {
		c.JSON(404, gin.H{"error": "Usuário não encontrado"})
		return
	}
	c.JSON(200, user)
}
```

### 3. Rotas Organizadas

Organize rotas em grupos:

```go
// routes/routes.go
func SetupRoutes(r *gin.Engine, handlers *Handlers) {
	api := r.Group("/api/v1")
	{
		users := api.Group("/users")
		{
			users.GET("", handlers.User.List)
			users.GET("/:id", handlers.User.Get)
			users.POST("", handlers.User.Create)
			users.PUT("/:id", handlers.User.Update)
			users.DELETE("/:id", handlers.User.Delete)
		}
	}
}
```

### 4. Injeção de Dependências

Use injeção de dependências:

```go
type App struct {
	router  *gin.Engine
	handlers *Handlers
	services *Services
}

func NewApp() *App {
	services := NewServices()
	handlers := NewHandlers(services)
	router := gin.Default()
	
	SetupRoutes(router, handlers)
	
	return &App{
		router:   router,
		handlers: handlers,
		services: services,
	}
}
```

---

## Testes

### 1. Testes de Handlers

Teste handlers com Gin:

```go
func TestGetUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	
	c.Request = httptest.NewRequest("GET", "/users/1", nil)
	c.Params = gin.Params{gin.Param{Key: "id", Value: "1"}}
	
	handler := &UserHandler{service: mockService}
	handler.GetUser(c)
	
	if w.Code != http.StatusOK {
		t.Errorf("Esperado %d, obteve %d", http.StatusOK, w.Code)
	}
}
```

### 2. Testes de Rotas

Teste rotas completas:

```go
func TestUserRoutes(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	SetupRoutes(r, mockHandlers)
	
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/api/v1/users/1", nil)
	r.ServeHTTP(w, req)
	
	if w.Code != http.StatusOK {
		t.Errorf("Esperado %d, obteve %d", http.StatusOK, w.Code)
	}
}
```

---

## Monitoramento

### 1. Métricas com Prometheus

Integre métricas:

```go
import "github.com/prometheus/client_golang/prometheus/promhttp"

func main() {
	r := gin.Default()
	
	// Endpoint de métricas
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))
	
	// ... outras rotas
}
```

### 2. Health Check

Implemente health check:

```go
func healthCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "healthy",
		"timestamp": time.Now().Unix(),
	})
}

func main() {
	r := gin.Default()
	r.GET("/health", healthCheck)
}
```

---

## Graceful Shutdown

Implemente shutdown graceful:

```go
func main() {
	router := gin.Default()
	// ... rotas
	
	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()
	
	// Capturar sinais
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	<-sigChan
	
	log.Println("Iniciando shutdown graceful...")
	
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Erro no shutdown:", err)
	}
	
	log.Println("Servidor encerrado")
}
```

---

## Conclusão

Seguindo essas boas práticas, você criará aplicações robustas, seguras e performáticas com Gin. Lembre-se:

- Valide sempre entrada
- Trate erros consistentemente
- Use middlewares adequadamente
- Organize o código
- Teste adequadamente
- Monitore em produção
- Configure segurança
- Otimize performance

Pratique essas técnicas e você estará pronto para criar APIs profissionais com Gin!

