# Aula 23: Web Development com net/http (Versão Simplificada)

Olá! Esta é uma versão resumida da aula 23 sobre desenvolvimento web com `net/http`.

---

## Por que Go é Excelente para Web?

- ✅ **Performance excepcional** (código nativo)
- ✅ **Concorrência nativa** (goroutines)
- ✅ **Biblioteca padrão poderosa** (net/http)
- ✅ **Deploy simples** (binário único)
- ✅ **Type safety** (menos erros)

---

## 1. Servidor HTTP Básico

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
	http.ListenAndServe(":8080", nil)
}
```

---

## 2. Métodos HTTP

```go
func handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		fmt.Fprintf(w, "GET")
	case http.MethodPost:
		fmt.Fprintf(w, "POST")
	default:
		http.Error(w, "Método não permitido", 405)
	}
}
```

---

## 3. Query Parameters

```go
func handler(w http.ResponseWriter, r *http.Request) {
	nome := r.URL.Query().Get("nome")
	idade := r.URL.Query().Get("idade")
	fmt.Fprintf(w, "Nome: %s, Idade: %s", nome, idade)
}
```

**Uso:** `GET /?nome=João&idade=25`

---

## 4. Form Data

```go
func handler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	nome := r.FormValue("nome")
	email := r.FormValue("email")
	fmt.Fprintf(w, "Nome: %s, Email: %s", nome, email)
}
```

---

## 5. JSON

### Receber JSON

```go
type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	var user User
	json.NewDecoder(r.Body).Decode(&user)
	defer r.Body.Close()
	// Processar user...
}
```

### Enviar JSON

```go
func handler(w http.ResponseWriter, r *http.Request) {
	user := User{Name: "João", Email: "joao@example.com"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
```

---

## 6. Routing com ServeMux

```go
mux := http.NewServeMux()
mux.HandleFunc("/", homeHandler)
mux.HandleFunc("/api/users", usersHandler)

server := &http.Server{
	Addr:    ":8080",
	Handler: mux,
}
server.ListenAndServe()
```

---

## 7. Middleware

```go
func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next(w, r)
		fmt.Printf("%s %s - %v\n", r.Method, r.URL.Path, time.Since(start))
	}
}

func main() {
	http.HandleFunc("/", loggingMiddleware(handler))
	http.ListenAndServe(":8080", nil)
}
```

---

## 8. Cliente HTTP

### GET Simples

```go
resp, err := http.Get("https://api.example.com/users")
defer resp.Body.Close()
body, _ := io.ReadAll(resp.Body)
```

### POST com JSON

```go
jsonData, _ := json.Marshal(user)
req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
req.Header.Set("Content-Type", "application/json")

client := &http.Client{Timeout: 10 * time.Second}
resp, _ := client.Do(req)
```

---

## 9. Arquivos Estáticos

```go
fs := http.FileServer(http.Dir("./static"))
http.Handle("/static/", http.StripPrefix("/static/", fs))
```

---

## 10. HTTPS

```go
server.ListenAndServeTLS("cert.pem", "key.pem")
```

---

## Resumo Rápido

| Tarefa | Código |
|--------|--------|
| Servidor básico | `http.HandleFunc("/", handler)` |
| Query params | `r.URL.Query().Get("key")` |
| Form data | `r.FormValue("key")` |
| JSON | `json.NewDecoder(r.Body).Decode(&obj)` |
| Middleware | `func(next http.HandlerFunc) http.HandlerFunc` |
| Cliente GET | `http.Get(url)` |
| Cliente POST | `http.NewRequest("POST", url, body)` |

---

Pronto! Agora você tem o básico de `net/http`. Pratique criando uma API REST simples!

