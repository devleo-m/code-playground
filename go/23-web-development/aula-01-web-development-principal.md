# Aula 23: Web Development com net/http

Olá, futuro(a) Gopher!

Bem-vindo(a) à aula 23! Agora que você já domina os fundamentos de Go e explorou diversas áreas, vamos mergulhar em uma das aplicações mais importantes da linguagem: **desenvolvimento web**.

Go é excepcional para desenvolvimento web. Sua performance, concorrência nativa, biblioteca padrão robusta e simplicidade de deploy fazem dela uma escolha ideal para construir APIs RESTful, microsserviços, servidores web escaláveis e muito mais. Nesta aula, vamos focar no pacote `net/http` da biblioteca padrão, que fornece tudo que você precisa para construir servidores e clientes HTTP sem dependências externas.

---

## Por que Go é Excelente para Web Development?

Antes de mergulharmos no código, vamos entender por que Go se destaca nessa área:

### 1. **Performance Excepcional**
Go compila para código de máquina nativo, resultando em aplicações web extremamente rápidas. Benchmarks mostram que servidores Go podem lidar com milhares de requisições por segundo com baixo uso de memória.

### 2. **Concorrência Nativa**
As `goroutines` permitem que você lide com milhares de conexões simultâneas sem a complexidade de threads tradicionais. Isso é perfeito para aplicações web que precisam escalar.

### 3. **Biblioteca Padrão Poderosa**
O pacote `net/http` da biblioteca padrão é completo e bem testado. Você pode construir servidores web robustos sem dependências externas.

### 4. **Deploy Simples**
Go compila para um único binário executável, sem dependências externas. Isso simplifica drasticamente o processo de deploy e distribuição.

### 5. **Type Safety**
Go é estaticamente tipada, o que ajuda a prevenir erros comuns em tempo de compilação, especialmente importante em APIs.

---

## O Pacote `net/http` Padrão

Vamos começar com o básico: o pacote `net/http` da biblioteca padrão. Ele fornece tudo que você precisa para construir servidores e clientes HTTP sem dependências externas.

### Servidor HTTP Básico

O exemplo mais simples de um servidor HTTP em Go:

```go
package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Olá, mundo!")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Servidor rodando em http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
```

**Executando:**
```bash
go run main.go
# Acesse http://localhost:8080 no navegador
```

### Entendendo `http.ResponseWriter` e `http.Request`

- **`http.ResponseWriter`**: Interface que permite escrever a resposta HTTP
- **`http.Request`**: Struct que contém todas as informações sobre a requisição recebida

```go
func handler(w http.ResponseWriter, r *http.Request) {
	// Informações da requisição
	fmt.Printf("Método: %s\n", r.Method)
	fmt.Printf("URL: %s\n", r.URL.Path)
	fmt.Printf("Host: %s\n", r.Host)
	fmt.Printf("RemoteAddr: %s\n", r.RemoteAddr)
	
	// Headers
	fmt.Printf("User-Agent: %s\n", r.Header.Get("User-Agent"))
	
	// Escrevendo resposta
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Requisição recebida!")
}
```

### Métodos HTTP

Você pode tratar diferentes métodos HTTP:

```go
func handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		fmt.Fprintf(w, "GET: Lendo dados")
	case http.MethodPost:
		fmt.Fprintf(w, "POST: Criando dados")
	case http.MethodPut:
		fmt.Fprintf(w, "PUT: Atualizando dados")
	case http.MethodDelete:
		fmt.Fprintf(w, "DELETE: Removendo dados")
	case http.MethodPatch:
		fmt.Fprintf(w, "PATCH: Atualização parcial")
	default:
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
	}
}
```

### Lendo Dados da Requisição

#### Query Parameters

```go
func handler(w http.ResponseWriter, r *http.Request) {
	// http://localhost:8080/?nome=João&idade=25
	nome := r.URL.Query().Get("nome")
	idade := r.URL.Query().Get("idade")
	
	// Ou obter todos os parâmetros
	queryParams := r.URL.Query()
	for key, values := range queryParams {
		fmt.Printf("%s: %v\n", key, values)
	}
	
	fmt.Fprintf(w, "Nome: %s, Idade: %s", nome, idade)
}
```

#### Form Data

```go
func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}
	
	// Parse do form
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Erro ao processar form", http.StatusBadRequest)
		return
	}
	
	nome := r.FormValue("nome")
	email := r.FormValue("email")
	
	// Ou obter todos os valores
	for key, values := range r.PostForm {
		fmt.Printf("%s: %v\n", key, values)
	}
	
	fmt.Fprintf(w, "Nome: %s, Email: %s", nome, email)
}
```

#### Multipart Form (Upload de Arquivos)

```go
func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}
	
	// Parse do multipart form (limite de 10MB)
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "Erro ao processar form", http.StatusBadRequest)
		return
	}
	
	// Obter arquivo
	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Erro ao obter arquivo", http.StatusBadRequest)
		return
	}
	defer file.Close()
	
	fmt.Printf("Arquivo recebido: %s, Tamanho: %d bytes\n", 
		handler.Filename, handler.Size)
	
	// Salvar arquivo (exemplo)
	// dst, _ := os.Create("./uploads/" + handler.Filename)
	// defer dst.Close()
	// io.Copy(dst, file)
	
	fmt.Fprintf(w, "Arquivo %s recebido com sucesso!", handler.Filename)
}
```

#### JSON Body

```go
import (
	"encoding/json"
	"net/http"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}
	
	// Verificar Content-Type
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Content-Type deve ser application/json", http.StatusBadRequest)
		return
	}
	
	var user User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)
	if err != nil {
		http.Error(w, "Erro ao decodificar JSON", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	
	// Processar user...
	fmt.Fprintf(w, "Usuário recebido: %+v", user)
}
```

### Enviando Respostas

#### Resposta de Texto Simples

```go
func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Olá, mundo!")
}
```

#### Resposta JSON

```go
func handler(w http.ResponseWriter, r *http.Request) {
	user := User{
		Name:  "João",
		Email: "joao@example.com",
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	
	encoder := json.NewEncoder(w)
	if err := encoder.Encode(user); err != nil {
		http.Error(w, "Erro ao codificar JSON", http.StatusInternalServerError)
		return
	}
}
```

#### Resposta HTML

```go
func handler(w http.ResponseWriter, r *http.Request) {
	html := `
	<!DOCTYPE html>
	<html>
	<head>
		<title>Página Go</title>
	</head>
	<body>
		<h1>Olá do Go!</h1>
	</body>
	</html>
	`
	
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, html)
}
```

#### Códigos de Status HTTP

```go
// Sucesso
w.WriteHeader(http.StatusOK)        // 200
w.WriteHeader(http.StatusCreated)     // 201
w.WriteHeader(http.StatusNoContent)   // 204

// Redirecionamento
w.WriteHeader(http.StatusMovedPermanently)  // 301
w.WriteHeader(http.StatusFound)              // 302

// Erro do Cliente
w.WriteHeader(http.StatusBadRequest)         // 400
w.WriteHeader(http.StatusUnauthorized)        // 401
w.WriteHeader(http.StatusForbidden)          // 403
w.WriteHeader(http.StatusNotFound)           // 404
w.WriteHeader(http.StatusMethodNotAllowed)   // 405

// Erro do Servidor
w.WriteHeader(http.StatusInternalServerError) // 500
w.WriteHeader(http.StatusBadGateway)           // 502
w.WriteHeader(http.StatusServiceUnavailable)  // 503
```

### Routing Manual

Você pode criar rotas manualmente usando `http.HandleFunc`:

```go
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Página inicial")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Sobre nós")
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	// Extrair ID da URL: /user/123
	path := r.URL.Path
	id := strings.TrimPrefix(path, "/user/")
	fmt.Fprintf(w, "Usuário ID: %s", id)
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/user/", userHandler)
	
	http.ListenAndServe(":8080", nil)
}
```

### Usando `http.ServeMux` para Routing Avançado

`ServeMux` é um roteador HTTP que permite maior controle:

```go
func main() {
	mux := http.NewServeMux()
	
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/api/users", usersHandler)
	mux.HandleFunc("/api/users/", userDetailHandler)
	
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
	
	fmt.Println("Servidor rodando em http://localhost:8080")
	server.ListenAndServe()
}
```

### Middleware

Middleware são funções que processam requisições antes que elas cheguem ao handler final. Vamos criar um middleware simples de logging:

```go
func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		
		// Chamar o próximo handler
		next(w, r)
		
		// Log após processar
		duration := time.Since(start)
		fmt.Printf("%s %s - %v\n", r.Method, r.URL.Path, duration)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Olá, mundo!")
}

func main() {
	http.HandleFunc("/", loggingMiddleware(handler))
	http.ListenAndServe(":8080", nil)
}
```

### Middleware com `http.Handler`

Para maior flexibilidade, você pode usar `http.Handler`:

```go
type loggingMiddleware struct {
	next http.Handler
}

func (m *loggingMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	m.next.ServeHTTP(w, r)
	duration := time.Since(start)
	fmt.Printf("%s %s - %v\n", r.Method, r.URL.Path, duration)
}

func NewLoggingMiddleware(next http.Handler) http.Handler {
	return &loggingMiddleware{next: next}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)
	
	loggingMux := NewLoggingMiddleware(mux)
	http.ListenAndServe(":8080", loggingMux)
}
```

### Middleware Chain

Você pode encadear múltiplos middlewares:

```go
type Middleware func(http.HandlerFunc) http.HandlerFunc

func chainMiddleware(middlewares ...Middleware) Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		for i := len(middlewares) - 1; i >= 0; i-- {
			next = middlewares[i](next)
		}
		return next
	}
}

func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Request: %s %s\n", r.Method, r.URL.Path)
		next(w, r)
	}
}

func authMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			http.Error(w, "Não autorizado", http.StatusUnauthorized)
			return
		}
		next(w, r)
	}
}

func corsMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		
		next(w, r)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Dados protegidos")
}

func main() {
	middleware := chainMiddleware(
		loggingMiddleware,
		corsMiddleware,
		authMiddleware,
	)
	http.HandleFunc("/protected", middleware(handler))
	http.ListenAndServe(":8080", nil)
}
```

### Servindo Arquivos Estáticos

```go
func main() {
	// Servir arquivos estáticos de um diretório
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	
	// Servir um único arquivo
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./index.html")
	})
	
	http.ListenAndServe(":8080", nil)
}
```

### Cliente HTTP

O pacote `net/http` também fornece um cliente HTTP poderoso:

```go
func main() {
	// GET request simples
	resp, err := http.Get("https://httpbin.org/get")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println(string(body))
}
```

### Cliente HTTP com Mais Controle

```go
func main() {
	// Criar request customizado
	req, err := http.NewRequest("GET", "https://api.example.com/users", nil)
	if err != nil {
		log.Fatal(err)
	}
	
	// Adicionar headers
	req.Header.Set("Authorization", "Bearer token123")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "MyApp/1.0")
	
	// Criar cliente com timeout e configurações
	client := &http.Client{
		Timeout: 10 * time.Second,
		Transport: &http.Transport{
			MaxIdleConns:       100,
			IdleConnTimeout:    90 * time.Second,
			DisableCompression: false,
		},
	}
	
	// Fazer request
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	
	// Verificar status
	if resp.StatusCode != http.StatusOK {
		log.Printf("Status inesperado: %s", resp.Status)
	}
	
	// Processar resposta
	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))
}
```

### POST Request com JSON

```go
func main() {
	user := User{
		Name:  "João",
		Email: "joao@example.com",
	}
	
	jsonData, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
	}
	
	req, err := http.NewRequest("POST", "https://api.example.com/users", 
		bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatal(err)
	}
	
	req.Header.Set("Content-Type", "application/json")
	
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	
	fmt.Printf("Status: %s\n", resp.Status)
	
	// Ler resposta
	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))
}
```

### Usando Context para Timeouts

```go
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	
	req, err := http.NewRequestWithContext(ctx, "GET", 
		"https://api.example.com/users", nil)
	if err != nil {
		log.Fatal(err)
	}
	
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		if err == context.DeadlineExceeded {
			log.Println("Timeout excedido")
		}
		log.Fatal(err)
	}
	defer resp.Body.Close()
	
	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))
}
```

### Exemplo Completo: API REST Simples

```go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserStore struct {
	mu    sync.RWMutex
	users []User
	nextID int
}

func NewUserStore() *UserStore {
	return &UserStore{
		users:  make([]User, 0),
		nextID: 1,
	}
}

func (s *UserStore) CreateUser(name, email string) User {
	s.mu.Lock()
	defer s.mu.Unlock()

	user := User{
		ID:    s.nextID,
		Name:  name,
		Email: email,
	}
	s.users = append(s.users, user)
	s.nextID++
	return user
}

func (s *UserStore) GetUser(id int) (User, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	for _, user := range s.users {
		if user.ID == id {
			return user, true
		}
	}
	return User{}, false
}

func (s *UserStore) GetAllUsers() []User {
	s.mu.RLock()
	defer s.mu.RUnlock()

	users := make([]User, len(s.users))
	copy(users, s.users)
	return users
}

func (s *UserStore) UpdateUser(id int, name, email string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	for i, user := range s.users {
		if user.ID == id {
			s.users[i].Name = name
			s.users[i].Email = email
			return true
		}
	}
	return false
}

func (s *UserStore) DeleteUser(id int) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	for i, user := range s.users {
		if user.ID == id {
			s.users = append(s.users[:i], s.users[i+1:]...)
			return true
		}
	}
	return false
}

var store = NewUserStore()

func usersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		users := store.GetAllUsers()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(users)

	case http.MethodPost:
		var user User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, "Erro ao decodificar JSON", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		newUser := store.CreateUser(user.Name, user.Email)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newUser)

	default:
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
	}
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	idStr := strings.TrimPrefix(path, "/api/users/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		user, found := store.GetUser(id)
		if !found {
			http.Error(w, "Usuário não encontrado", http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)

	case http.MethodPut:
		var user User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, "Erro ao decodificar JSON", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		if !store.UpdateUser(id, user.Name, user.Email) {
			http.Error(w, "Usuário não encontrado", http.StatusNotFound)
			return
		}
		user.ID = id
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)

	case http.MethodDelete:
		if !store.DeleteUser(id) {
			http.Error(w, "Usuário não encontrado", http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusNoContent)

	default:
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
	}
}

func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next(w, r)
		duration := time.Since(start)
		log.Printf("%s %s - %v", r.Method, r.URL.Path, duration)
	}
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/users", loggingMiddleware(usersHandler))
	mux.HandleFunc("/api/users/", loggingMiddleware(userHandler))

	server := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	fmt.Println("API REST rodando em http://localhost:8080")
	log.Fatal(server.ListenAndServe())
}
```

### HTTPS e TLS

Go suporta HTTPS nativamente:

```go
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "HTTPS funcionando!")
	})

	server := &http.Server{
		Addr:    ":443",
		Handler: mux,
	}

	// Certificado e chave privada
	log.Fatal(server.ListenAndServeTLS("cert.pem", "key.pem"))
}
```

### Redirecionamento HTTP para HTTPS

```go
func redirectToHTTPS(w http.ResponseWriter, r *http.Request) {
	httpsURL := "https://" + r.Host + r.RequestURI
	http.Redirect(w, r, httpsURL, http.StatusMovedPermanently)
}

func main() {
	// Servidor HTTP que redireciona para HTTPS
	go func() {
		http.ListenAndServe(":80", http.HandlerFunc(redirectToHTTPS))
	}()

	// Servidor HTTPS
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "HTTPS seguro!")
	})

	server := &http.Server{
		Addr:    ":443",
		Handler: mux,
	}

	log.Fatal(server.ListenAndServeTLS("cert.pem", "key.pem"))
}
```

---

## Boas Práticas

### 1. **Tratamento de Erros**
Sempre trate erros adequadamente:

```go
if err != nil {
	log.Printf("Erro: %v", err)
	http.Error(w, "Erro interno", http.StatusInternalServerError)
	return
}
```

### 2. **Context para Timeouts**
Use context para controlar timeouts:

```go
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

req, _ := http.NewRequestWithContext(ctx, "GET", url, nil)
```

### 3. **Validação de Entrada**
Sempre valide dados de entrada:

```go
if user.Name == "" {
	http.Error(w, "Nome é obrigatório", http.StatusBadRequest)
	return
}
```

### 4. **Logging Estruturado**
Use logging estruturado:

```go
import "log/slog"

logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
logger.Info("Request processado", 
	"method", r.Method, 
	"path", r.URL.Path,
	"status", statusCode)
```

### 5. **Segurança**
- Use HTTPS em produção
- Valide e sanitize todas as entradas
- Use middleware de autenticação
- Implemente rate limiting
- Proteja contra CSRF e XSS
- Não exponha informações sensíveis em erros

### 6. **Testes**
Teste seus handlers:

```go
func TestHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	handler(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Esperado %d, obteve %d", http.StatusOK, w.Code)
	}
}
```

### 7. **Configuração do Servidor**
Configure timeouts adequados:

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

## Conclusão

Nesta aula, exploramos o desenvolvimento web em Go usando o pacote `net/http` da biblioteca padrão:

1. **Servidor HTTP**: Como criar servidores HTTP básicos e avançados
2. **Routing**: Como criar e gerenciar rotas
3. **Middleware**: Como criar e encadear middlewares
4. **Cliente HTTP**: Como fazer requisições HTTP
5. **API REST**: Como construir uma API REST completa
6. **HTTPS/TLS**: Como configurar servidores seguros

O pacote `net/http` é poderoso e suficiente para a maioria das aplicações web. Ele oferece controle total e zero dependências externas, tornando-o ideal para entender como tudo funciona e para aplicações que precisam de máxima performance.

Na próxima aula, vamos explorar frameworks web como Gin, que oferecem mais conveniência e produtividade para desenvolvimento de APIs modernas!

Até lá, experimente criar sua própria API REST usando apenas `net/http`. Pratique criando diferentes endpoints e middlewares!
