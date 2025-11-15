package main

// Este arquivo contém exemplos práticos de Web Development em Go
// Execute cada exemplo separadamente comentando/descomentando as funções main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

// ============================================================================
// EXEMPLO 1: Servidor HTTP Básico com net/http
// ============================================================================

func exemplo1ServidorBasico() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Olá, mundo!")
	})

	fmt.Println("Servidor rodando em http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

// ============================================================================
// EXEMPLO 2: API REST Simples com net/http
// ============================================================================

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

	user, found := store.GetUser(id)
	if !found {
		http.Error(w, "Usuário não encontrado", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func exemplo2APIREST() {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/users", usersHandler)
	mux.HandleFunc("/api/users/", userHandler)

	fmt.Println("API REST rodando em http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

// ============================================================================
// EXEMPLO 3: Middleware de Logging
// ============================================================================

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

func exemplo3Middleware() {
	http.HandleFunc("/", loggingMiddleware(handler))
	fmt.Println("Servidor com middleware rodando em http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

// ============================================================================
// EXEMPLO 4: Cliente HTTP
// ============================================================================

func exemplo4ClienteHTTP() {
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

	fmt.Println("Resposta:")
	fmt.Println(string(body))
}

// ============================================================================
// EXEMPLO 5: POST Request com JSON
// ============================================================================

func exemplo5POSTRequest() {
	user := User{
		Name:  "João",
		Email: "joao@example.com",
	}

	jsonData, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest("POST", "https://httpbin.org/post",
		strings.NewReader(string(jsonData)))
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

	body, _ := io.ReadAll(resp.Body)
	fmt.Println("Resposta:")
	fmt.Println(string(body))
}

// ============================================================================
// EXEMPLO 6: Query Parameters e Form Data
// ============================================================================

func queryHandler(w http.ResponseWriter, r *http.Request) {
	// Query parameters: ?nome=João&idade=25
	nome := r.URL.Query().Get("nome")
	idade := r.URL.Query().Get("idade")

	fmt.Fprintf(w, "Nome: %s, Idade: %s\n", nome, idade)
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Erro ao processar form", http.StatusBadRequest)
		return
	}

	nome := r.FormValue("nome")
	email := r.FormValue("email")

	fmt.Fprintf(w, "Nome: %s, Email: %s\n", nome, email)
}

func exemplo6QueryForm() {
	mux := http.NewServeMux()
	mux.HandleFunc("/query", queryHandler)
	mux.HandleFunc("/form", formHandler)

	fmt.Println("Servidor rodando em http://localhost:8080")
	http.ListenAndServe(":8080", mux)
}

// ============================================================================
// EXEMPLO 7: Servindo Arquivos Estáticos
// ============================================================================

func exemplo7ArquivosEstaticos() {
	// Servir arquivos estáticos de um diretório
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Servir um único arquivo
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./index.html")
	})

	fmt.Println("Servidor de arquivos estáticos rodando em http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

// ============================================================================
// MAIN - Descomente o exemplo que deseja executar
// ============================================================================

func main() {
	// Descomente o exemplo que deseja executar:

	// exemplo1ServidorBasico()
	// exemplo2APIREST()
	// exemplo3Middleware()
	// exemplo4ClienteHTTP()
	// exemplo5POSTRequest()
	// exemplo6QueryForm()
	// exemplo7ArquivosEstaticos()

	fmt.Println("Nenhum exemplo selecionado. Descomente um exemplo no main()")
}

