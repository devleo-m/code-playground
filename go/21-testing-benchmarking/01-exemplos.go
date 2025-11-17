package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
)

// ============================================
// EXEMPLO 1: Função Simples com Teste Básico
// ============================================

func Soma(a, b int) int {
	return a + b
}

func exemplo1_TesteBasico() {
	fmt.Println("=== Exemplo 1: Teste Básico ===")
	fmt.Println("Execute: go test -v")
	fmt.Println()
	
	resultado := Soma(2, 3)
	fmt.Printf("Soma(2, 3) = %d\n", resultado)
	fmt.Println("✅ Crie um arquivo soma_test.go para testar esta função!")
}

// ============================================
// EXEMPLO 2: Table-driven Test
// ============================================

func ValidarEmail(email string) (bool, error) {
	if email == "" {
		return false, errors.New("email não pode ser vazio")
	}
	
	if !strings.Contains(email, "@") {
		return false, errors.New("email deve conter @")
	}
	
	partes := strings.Split(email, "@")
	if len(partes) != 2 {
		return false, errors.New("email deve ter exatamente um @")
	}
	
	if partes[0] == "" {
		return false, errors.New("email deve ter parte local antes do @")
	}
	
	if partes[1] == "" || !strings.Contains(partes[1], ".") {
		return false, errors.New("email deve ter domínio válido")
	}
	
	return true, nil
}

func exemplo2_TableDrivenTest() {
	fmt.Println("=== Exemplo 2: Table-driven Test ===")
	fmt.Println("Execute: go test -v -run TestValidarEmail")
	fmt.Println()
	
	casos := []struct {
		email    string
		esperado bool
	}{
		{"usuario@exemplo.com", true},
		{"invalido", false},
		{"", false},
		{"@exemplo.com", false},
		{"usuario@", false},
	}
	
	for _, caso := range casos {
		valido, err := ValidarEmail(caso.email)
		status := "✅"
		if valido != caso.esperado {
			status = "❌"
		}
		fmt.Printf("%s Email: %-20s | Válido: %v", status, caso.email, valido)
		if err != nil {
			fmt.Printf(" | Erro: %s", err.Error())
		}
		fmt.Println()
	}
	fmt.Println()
	fmt.Println("💡 Crie um arquivo email_test.go com table-driven tests!")
}

// ============================================
// EXEMPLO 3: Mock de Repositório
// ============================================

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type UserRepository interface {
	FindByID(id int) (*User, error)
	Save(user *User) error
}

// Implementação real (não usada em testes)
type InMemoryRepository struct {
	users map[int]*User
}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{
		users: make(map[int]*User),
	}
}

func (r *InMemoryRepository) FindByID(id int) (*User, error) {
	user, exists := r.users[id]
	if !exists {
		return nil, errors.New("usuário não encontrado")
	}
	return user, nil
}

func (r *InMemoryRepository) Save(user *User) error {
	if user.ID == 0 {
		return errors.New("ID não pode ser zero")
	}
	r.users[user.ID] = user
	return nil
}

// Mock para testes
type MockUserRepository struct {
	FindByIDCalls []int
	SaveCalls     []*User
	FindByIDFunc  func(int) (*User, error)
	SaveFunc      func(*User) error
}

func (m *MockUserRepository) FindByID(id int) (*User, error) {
	m.FindByIDCalls = append(m.FindByIDCalls, id)
	if m.FindByIDFunc != nil {
		return m.FindByIDFunc(id)
	}
	return &User{ID: id, Name: "Mock User"}, nil
}

func (m *MockUserRepository) Save(user *User) error {
	m.SaveCalls = append(m.SaveCalls, user)
	if m.SaveFunc != nil {
		return m.SaveFunc(user)
	}
	return nil
}

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetUser(id int) (*User, error) {
	return s.repo.FindByID(id)
}

func (s *UserService) CreateUser(name string) (*User, error) {
	user := &User{ID: len(s.repo.(*MockUserRepository).SaveCalls) + 1, Name: name}
	return user, s.repo.Save(user)
}

func exemplo3_Mock() {
	fmt.Println("=== Exemplo 3: Mock de Repositório ===")
	fmt.Println("Execute: go test -v -run TestUserService")
	fmt.Println()
	
	// Usando mock
	mockRepo := &MockUserRepository{
		FindByIDFunc: func(id int) (*User, error) {
			return &User{ID: id, Name: "Test User"}, nil
		},
	}
	
	service := NewUserService(mockRepo)
	user, err := service.GetUser(123)
	
	if err != nil {
		fmt.Printf("❌ Erro: %v\n", err)
	} else {
		fmt.Printf("✅ User obtido: ID=%d, Name=%s\n", user.ID, user.Name)
		fmt.Printf("✅ FindByID foi chamado %d vez(es)\n", len(mockRepo.FindByIDCalls))
		fmt.Printf("✅ ID passado: %d\n", mockRepo.FindByIDCalls[0])
	}
	fmt.Println()
	fmt.Println("💡 Crie um arquivo user_test.go para testar com mocks!")
}

// ============================================
// EXEMPLO 4: Teste HTTP com httptest
// ============================================

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}
	
	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{
		"message": fmt.Sprintf("Hello, %s!", name),
	}
	json.NewEncoder(w).Encode(response)
}

func exemplo4_HTTPTest() {
	fmt.Println("=== Exemplo 4: Teste HTTP com httptest ===")
	fmt.Println("Execute: go test -v -run TestHelloHandler")
	fmt.Println()
	
	// Simular requisição
	req := httptest.NewRequest("GET", "/hello?name=Go", nil)
	rr := httptest.NewRecorder()
	
	HelloHandler(rr, req)
	
	fmt.Printf("✅ Status Code: %d\n", rr.Code)
	fmt.Printf("✅ Response Body: %s\n", rr.Body.String())
	fmt.Println()
	
	// Teste sem parâmetro
	req2 := httptest.NewRequest("GET", "/hello", nil)
	rr2 := httptest.NewRecorder()
	
	HelloHandler(rr2, req2)
	
	fmt.Printf("✅ Status Code (sem name): %d\n", rr2.Code)
	fmt.Printf("✅ Response Body: %s\n", rr2.Body.String())
	fmt.Println()
	fmt.Println("💡 Crie um arquivo handler_test.go para testar handlers HTTP!")
}

// ============================================
// EXEMPLO 5: Benchmark Simples
// ============================================

func SomaLoop(slice []int) int {
	soma := 0
	for i := 0; i < len(slice); i++ {
		soma += slice[i]
	}
	return soma
}

func SomaRange(slice []int) int {
	soma := 0
	for _, valor := range slice {
		soma += valor
	}
	return soma
}

func exemplo5_Benchmark() {
	fmt.Println("=== Exemplo 5: Benchmark ===")
	fmt.Println("Execute: go test -bench=. -benchmem")
	fmt.Println()
	
	slice := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		slice[i] = i
	}
	
	// Executar algumas vezes para "aquecer"
	for i := 0; i < 1000; i++ {
		SomaLoop(slice)
		SomaRange(slice)
	}
	
	fmt.Println("✅ Funções implementadas!")
	fmt.Println("💡 Crie um arquivo soma_test.go com benchmarks:")
	fmt.Println("   func BenchmarkSomaLoop(b *testing.B) { ... }")
	fmt.Println("   func BenchmarkSomaRange(b *testing.B) { ... }")
	fmt.Println()
}

// ============================================
// EXEMPLO 6: Cobertura de Testes
// ============================================

func Dividir(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("divisão por zero")
	}
	return a / b, nil
}

func exemplo6_Coverage() {
	fmt.Println("=== Exemplo 6: Cobertura de Testes ===")
	fmt.Println("Execute: go test -cover")
	fmt.Println("Execute: go test -coverprofile=coverage.out")
	fmt.Println("Execute: go tool cover -html=coverage.out")
	fmt.Println()
	
	resultado, err := Dividir(10, 2)
	if err != nil {
		fmt.Printf("❌ Erro: %v\n", err)
	} else {
		fmt.Printf("✅ Dividir(10, 2) = %.2f\n", resultado)
	}
	
	_, err = Dividir(10, 0)
	if err != nil {
		fmt.Printf("✅ Erro esperado capturado: %v\n", err)
	}
	fmt.Println()
	fmt.Println("💡 Crie testes para ambos os casos (sucesso e erro)!")
	fmt.Println("💡 Visualize a cobertura em HTML para ver linhas não testadas!")
}

// ============================================
// EXEMPLO 7: Servidor HTTP de Teste
// ============================================

func exemplo7_HTTPServer() {
	fmt.Println("=== Exemplo 7: Servidor HTTP de Teste ===")
	fmt.Println("Execute: go test -v -run TestAPIServer")
	fmt.Println()
	
	// Criar servidor de teste
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/api/users":
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode([]User{
				{ID: 1, Name: "Alice"},
				{ID: 2, Name: "Bob"},
			})
		case "/api/users/1":
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(User{ID: 1, Name: "Alice"})
		default:
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{"error": "not found"})
		}
	}))
	defer server.Close()
	
	// Fazer requisição
	resp, err := http.Get(server.URL + "/api/users")
	if err != nil {
		fmt.Printf("❌ Erro: %v\n", err)
		return
	}
	defer resp.Body.Close()
	
	fmt.Printf("✅ Status Code: %d\n", resp.StatusCode)
	fmt.Printf("✅ URL do servidor de teste: %s\n", server.URL)
	fmt.Println()
	fmt.Println("💡 Use httptest.NewServer para testar clientes HTTP!")
}

// ============================================
// MAIN: Execute os exemplos
// ============================================

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Uso: go run 01-exemplos.go [exemplo1|exemplo2|exemplo3|exemplo4|exemplo5|exemplo6|exemplo7]")
		fmt.Println()
		fmt.Println("Exemplos disponíveis:")
		fmt.Println("  exemplo1 - Teste Básico")
		fmt.Println("  exemplo2 - Table-driven Test")
		fmt.Println("  exemplo3 - Mock de Repositório")
		fmt.Println("  exemplo4 - Teste HTTP com httptest")
		fmt.Println("  exemplo5 - Benchmark")
		fmt.Println("  exemplo6 - Cobertura de Testes")
		fmt.Println("  exemplo7 - Servidor HTTP de Teste")
		fmt.Println()
		fmt.Println("💡 Dica: Execute os testes correspondentes:")
		fmt.Println("   go test -v")
		fmt.Println("   go test -bench=. -benchmem")
		fmt.Println("   go test -cover")
		return
	}
	
	exemplo := os.Args[1]
	
	switch exemplo {
	case "exemplo1":
		exemplo1_TesteBasico()
	case "exemplo2":
		exemplo2_TableDrivenTest()
	case "exemplo3":
		exemplo3_Mock()
	case "exemplo4":
		exemplo4_HTTPTest()
	case "exemplo5":
		exemplo5_Benchmark()
	case "exemplo6":
		exemplo6_Coverage()
	case "exemplo7":
		exemplo7_HTTPServer()
	default:
		fmt.Printf("Exemplo '%s' não encontrado!\n", exemplo)
		fmt.Println("Use: exemplo1, exemplo2, exemplo3, exemplo4, exemplo5, exemplo6 ou exemplo7")
	}
}


