# Aula 23: Performance e Boas Práticas - Web Development com net/http

Olá! Nesta parte da aula, vamos explorar boas práticas, otimizações de performance e padrões profissionais para desenvolvimento web com `net/http`.

---

## Boas Práticas Gerais

### 1. Tratamento de Erros Robusto

**❌ Ruim:**
```go
func handler(w http.ResponseWriter, r *http.Request) {
	user, err := getUser(r.URL.Query().Get("id"))
	fmt.Fprintf(w, "%v", user)
}
```

**✅ Bom:**
```go
func handler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID é obrigatório", http.StatusBadRequest)
		return
	}
	
	user, err := getUser(id)
	if err != nil {
		log.Printf("Erro ao buscar usuário %s: %v", id, err)
		http.Error(w, "Erro ao processar requisição", http.StatusInternalServerError)
		return
	}
	
	if user == nil {
		http.Error(w, "Usuário não encontrado", http.StatusNotFound)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
```

**Princípios:**
- Sempre valide entrada
- Trate todos os erros
- Use códigos de status HTTP apropriados
- Log erros internos, mas não exponha detalhes ao cliente

### 2. Códigos de Status HTTP Apropriados

Use códigos de status corretos:

```go
// Sucesso
w.WriteHeader(http.StatusOK)        // 200
w.WriteHeader(http.StatusCreated)    // 201
w.WriteHeader(http.StatusNoContent)  // 204

// Erro do Cliente
w.WriteHeader(http.StatusBadRequest)      // 400
w.WriteHeader(http.StatusUnauthorized)    // 401
w.WriteHeader(http.StatusForbidden)      // 403
w.WriteHeader(http.StatusNotFound)       // 404
w.WriteHeader(http.StatusMethodNotAllowed) // 405

// Erro do Servidor
w.WriteHeader(http.StatusInternalServerError) // 500
w.WriteHeader(http.StatusBadGateway)           // 502
w.WriteHeader(http.StatusServiceUnavailable)   // 503
```

### 3. Validação de Entrada

Sempre valide dados de entrada:

```go
func validateUser(user *User) error {
	if user.Name == "" {
		return fmt.Errorf("nome é obrigatório")
	}
	if len(user.Name) < 3 {
		return fmt.Errorf("nome deve ter no mínimo 3 caracteres")
	}
	if !isValidEmail(user.Email) {
		return fmt.Errorf("email inválido: %s", user.Email)
	}
	if user.Age < 18 || user.Age > 120 {
		return fmt.Errorf("idade deve estar entre 18 e 120")
	}
	return nil
}

func isValidEmail(email string) bool {
	return strings.Contains(email, "@") && strings.Contains(email, ".")
}
```

### 4. Context para Timeouts e Cancelamento

Use context para controlar timeouts:

```go
func handler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()
	
	// Passar context para operações
	result, err := processWithContext(ctx)
	if err != nil {
		if err == context.DeadlineExceeded {
			http.Error(w, "Timeout", http.StatusRequestTimeout)
			return
		}
		http.Error(w, "Erro interno", http.StatusInternalServerError)
		return
	}
	
	json.NewEncoder(w).Encode(result)
}
```

### 5. Logging Estruturado

Use logging estruturado:

```go
import (
	"log/slog"
	"os"
)

func setupLogger() *slog.Logger {
	return slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))
}

func handler(w http.ResponseWriter, r *http.Request) {
	logger := r.Context().Value("logger").(*slog.Logger)
	
	logger.Info("Request recebida",
		"method", r.Method,
		"path", r.URL.Path,
		"ip", r.RemoteAddr,
	)
	
	// Processar...
	
	logger.Info("Request processada",
		"method", r.Method,
		"path", r.URL.Path,
		"status", 200,
		"duration_ms", duration.Milliseconds(),
	)
}
```

---

## Performance

### 1. Connection Pooling

Configure o cliente HTTP adequadamente:

```go
client := &http.Client{
	Timeout: 30 * time.Second,
	Transport: &http.Transport{
		MaxIdleConns:        100,
		MaxIdleConnsPerHost: 10,
		IdleConnTimeout:     90 * time.Second,
		DisableCompression:  false,
	},
}
```

### 2. Reutilização de Buffers

Reutilize buffers para reduzir alocações:

```go
var bufferPool = sync.Pool{
	New: func() interface{} {
		return make([]byte, 0, 4096)
	},
}

func handler(w http.ResponseWriter, r *http.Request) {
	buf := bufferPool.Get().([]byte)
	defer bufferPool.Put(buf[:0])
	
	// Usar buffer
	buf = append(buf, "dados"...)
	w.Write(buf)
}
```

### 3. Compressão de Respostas

Use compressão para reduzir tamanho de respostas:

```go
import "compress/gzip"

func gzipMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			next(w, r)
			return
		}
		
		w.Header().Set("Content-Encoding", "gzip")
		gz := gzip.NewWriter(w)
		defer gz.Close()
		
		gzw := gzipResponseWriter{Writer: gz, ResponseWriter: w}
		next(gzw, r)
	}
}
```

### 4. Cache de Respostas

Implemente cache para respostas estáticas:

```go
type cacheEntry struct {
	data      []byte
	timestamp time.Time
	ttl       time.Duration
}

var cache = sync.Map{}

func cachedHandler(key string, ttl time.Duration, handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			handler(w, r)
			return
		}
		
		if entry, ok := cache.Load(key); ok {
			e := entry.(cacheEntry)
			if time.Since(e.timestamp) < e.ttl {
				w.Write(e.data)
				return
			}
		}
		
		rec := &responseRecorder{ResponseWriter: w, body: &bytes.Buffer{}}
		handler(rec, r)
		
		cache.Store(key, cacheEntry{
			data:      rec.body.Bytes(),
			timestamp: time.Now(),
			ttl:       ttl,
		})
	}
}
```

### 5. Configuração do Servidor

Configure o servidor adequadamente:

```go
server := &http.Server{
	Addr:         ":8080",
	Handler:      mux,
	ReadTimeout:  15 * time.Second,
	WriteTimeout: 15 * time.Second,
	IdleTimeout:  60 * time.Second,
	MaxHeaderBytes: 1 << 20, // 1MB
}
```

---

## Segurança

### 1. Validação e Sanitização

Sempre valide e sanitize entrada:

```go
import "html"

func sanitizeInput(input string) string {
	// Remover HTML tags
	input = html.EscapeString(input)
	// Remover caracteres perigosos
	input = strings.TrimSpace(input)
	return input
}
```

### 2. Proteção CSRF

Implemente proteção CSRF:

```go
func csrfMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost || r.Method == http.MethodPut || r.Method == http.MethodDelete {
			token := r.Header.Get("X-CSRF-Token")
			if !isValidCSRFToken(token) {
				http.Error(w, "CSRF token inválido", http.StatusForbidden)
				return
			}
		}
		next(w, r)
	}
}
```

### 3. Rate Limiting

Implemente rate limiting:

```go
type rateLimiter struct {
	requests map[string][]time.Time
	mu       sync.Mutex
	maxReqs  int
	window   time.Duration
}

func (rl *rateLimiter) allow(ip string) bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()
	
	now := time.Now()
	cutoff := now.Add(-rl.window)
	
	// Limpar requisições antigas
	reqs := rl.requests[ip]
	validReqs := []time.Time{}
	for _, req := range reqs {
		if req.After(cutoff) {
			validReqs = append(validReqs, req)
		}
	}
	
	if len(validReqs) >= rl.maxReqs {
		return false
	}
	
	validReqs = append(validReqs, now)
	rl.requests[ip] = validReqs
	return true
}

func rateLimitMiddleware(next http.HandlerFunc) http.HandlerFunc {
	limiter := &rateLimiter{
		requests: make(map[string][]time.Time),
		maxReqs:  100,
		window:   time.Minute,
	}
	
	return func(w http.ResponseWriter, r *http.Request) {
		ip := r.RemoteAddr
		if !limiter.allow(ip) {
			http.Error(w, "Rate limit excedido", http.StatusTooManyRequests)
			return
		}
		next(w, r)
	}
}
```

### 4. HTTPS em Produção

Sempre use HTTPS em produção:

```go
func main() {
	mux := http.NewServeMux()
	// ... rotas
	
	server := &http.Server{
		Addr:    ":443",
		Handler: mux,
	}
	
	// Redirecionar HTTP para HTTPS
	go func() {
		http.ListenAndServe(":80", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, "https://"+r.Host+r.RequestURI, http.StatusMovedPermanently)
		}))
	}()
	
	log.Fatal(server.ListenAndServeTLS("cert.pem", "key.pem"))
}
```

### 5. Headers de Segurança

Adicione headers de segurança:

```go
func securityHeadersMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("X-Frame-Options", "DENY")
		w.Header().Set("X-XSS-Protection", "1; mode=block")
		w.Header().Set("Strict-Transport-Security", "max-age=31536000; includeSubDomains")
		w.Header().Set("Content-Security-Policy", "default-src 'self'")
		next(w, r)
	}
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
│   └── store/
│       └── user.go
├── pkg/
│   └── utils/
└── go.mod
```

### 2. Separação de Responsabilidades

Separe handlers, models e lógica de negócio:

```go
// models/user.go
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// store/user.go
type UserStore interface {
	GetUser(id int) (*User, error)
	CreateUser(user *User) error
}

// handlers/user.go
type UserHandler struct {
	store UserStore
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	// Lógica do handler
}
```

### 3. Injeção de Dependências

Use injeção de dependências:

```go
type App struct {
	store   UserStore
	logger  *slog.Logger
	handler *UserHandler
}

func NewApp(store UserStore, logger *slog.Logger) *App {
	return &App{
		store:   store,
		logger:  logger,
		handler: NewUserHandler(store, logger),
	}
}

func (a *App) SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/users", a.handler.GetUser)
	return mux
}
```

---

## Graceful Shutdown

Implemente shutdown graceful:

```go
func main() {
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
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

## Testes

### 1. Testes de Handlers

Teste seus handlers:

```go
func TestGetUser(t *testing.T) {
	req := httptest.NewRequest("GET", "/users/1", nil)
	w := httptest.NewRecorder()
	
	handler := &UserHandler{store: mockStore}
	handler.GetUser(w, req)
	
	if w.Code != http.StatusOK {
		t.Errorf("Esperado %d, obteve %d", http.StatusOK, w.Code)
	}
	
	var user User
	json.NewDecoder(w.Body).Decode(&user)
	if user.ID != 1 {
		t.Errorf("Esperado ID 1, obteve %d", user.ID)
	}
}
```

### 2. Testes de Middleware

Teste middlewares:

```go
func TestAuthMiddleware(t *testing.T) {
	req := httptest.NewRequest("GET", "/protected", nil)
	w := httptest.NewRecorder()
	
	handler := authMiddleware(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	
	handler(w, req)
	
	if w.Code != http.StatusUnauthorized {
		t.Errorf("Esperado 401, obteve %d", w.Code)
	}
}
```

---

## Monitoramento

### 1. Métricas Básicas

Colete métricas básicas:

```go
type metrics struct {
	requests    int64
	errors      int64
	duration    time.Duration
	mu          sync.Mutex
}

func (m *metrics) recordRequest(duration time.Duration, err error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.requests++
	if err != nil {
		m.errors++
	}
	m.duration += duration
}

func metricsMiddleware(next http.HandlerFunc) http.HandlerFunc {
	m := &metrics{}
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next(w, r)
		duration := time.Since(start)
		m.recordRequest(duration, nil)
	}
}
```

---

## Conclusão

Seguindo essas boas práticas, você criará aplicações web robustas, seguras e performáticas com `net/http`. Lembre-se:

- Sempre valide entrada
- Trate todos os erros
- Use códigos de status apropriados
- Configure timeouts
- Implemente segurança
- Organize o código
- Teste adequadamente
- Monitore em produção

Pratique essas técnicas e você estará pronto para criar aplicações web profissionais!

