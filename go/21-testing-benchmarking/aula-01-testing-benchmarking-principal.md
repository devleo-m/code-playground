# Aula 21: Testing & Benchmarking - Testando e Medindo Performance

Olá! Bem-vindo(a) à aula 21!

Agora que você já domina os conceitos fundamentais de Go, é hora de aprender uma das habilidades mais importantes para qualquer desenvolvedor: **escrever testes confiáveis** e **medir a performance** do seu código.

Nesta aula, vamos explorar o pacote `testing` da biblioteca padrão do Go, aprender diferentes estratégias de teste, criar mocks e stubs, testar código HTTP, escrever benchmarks e medir cobertura de código.

---

## Por que Testar?

Antes de mergulharmos nos detalhes técnicos, vamos entender a importância dos testes:

1. **Confiança**: Testes garantem que seu código funciona como esperado
2. **Documentação**: Testes servem como documentação viva do comportamento do código
3. **Refatoração Segura**: Com testes, você pode refatorar código com confiança
4. **Detecção Precoce de Bugs**: Testes encontram problemas antes que cheguem à produção
5. **Design Melhor**: Escrever testes força você a pensar em design de código mais limpo

---

## Testing Package Basics

O pacote `testing` é parte da biblioteca padrão do Go e fornece suporte para testes automatizados, benchmarks e exemplos.

### Estrutura Básica de um Teste

Um teste em Go é uma função que:
- Está em um arquivo que termina com `_test.go`
- Tem o nome começando com `Test`
- Recebe um parâmetro `*testing.T`

```go
package main

import "testing"

func TestSoma(t *testing.T) {
    resultado := Soma(2, 3)
    esperado := 5
    
    if resultado != esperado {
        t.Errorf("Soma(2, 3) = %d; esperado %d", resultado, esperado)
    }
}
```

### Funções de Falha

O `*testing.T` fornece várias funções para reportar falhas:

- **`t.Error()` / `t.Errorf()`**: Reporta um erro mas continua a execução do teste
- **`t.Fatal()` / `t.Fatalf()`**: Reporta um erro e **para** a execução do teste imediatamente
- **`t.Skip()` / `t.Skipf()`**: Pula o teste (útil para testes condicionais)

```go
func TestDivisao(t *testing.T) {
    resultado := Dividir(10, 2)
    if resultado != 5 {
        t.Errorf("Esperado 5, obteve %f", resultado)
    }
    
    // Teste de divisão por zero
    resultado = Dividir(10, 0)
    if resultado != 0 {
        t.Fatalf("Divisão por zero deveria retornar 0, obteve %f", resultado)
    }
    
    // Este código não será executado se t.Fatalf() for chamado
    t.Log("Este código só executa se não houver falha fatal")
}
```

### Executando Testes

```bash
# Executar todos os testes no pacote atual
go test

# Executar testes com verbosidade
go test -v

# Executar um teste específico
go test -run TestSoma

# Executar testes com regex
go test -run "Test.*Soma"

# Executar testes de múltiplos pacotes
go test ./...
```

### Exemplo Completo

```go
// math.go
package main

func Soma(a, b int) int {
    return a + b
}

func Multiplicar(a, b int) int {
    return a * b
}

// math_test.go
package main

import "testing"

func TestSoma(t *testing.T) {
    casos := []struct {
        a, b, esperado int
    }{
        {2, 3, 5},
        {0, 0, 0},
        {-1, 1, 0},
        {10, -5, 5},
    }
    
    for _, caso := range casos {
        resultado := Soma(caso.a, caso.b)
        if resultado != caso.esperado {
            t.Errorf("Soma(%d, %d) = %d; esperado %d", 
                caso.a, caso.b, resultado, caso.esperado)
        }
    }
}

func TestMultiplicar(t *testing.T) {
    resultado := Multiplicar(3, 4)
    esperado := 12
    
    if resultado != esperado {
        t.Fatalf("Multiplicar(3, 4) = %d; esperado %d", resultado, esperado)
    }
}
```

---

## Table-driven Tests

Table-driven tests (testes orientados por tabela) são uma das práticas mais comuns e recomendadas em Go. Eles permitem testar múltiplos cenários com a mesma lógica de teste.

### Estrutura de um Table-driven Test

```go
func TestFuncao(t *testing.T) {
    casos := []struct {
        nome     string  // Nome do caso de teste
        entrada  int     // Entrada do teste
        esperado int     // Resultado esperado
    }{
        {
            nome:     "número positivo",
            entrada:  5,
            esperado: 25,
        },
        {
            nome:     "zero",
            entrada:  0,
            esperado: 0,
        },
        {
            nome:     "número negativo",
            entrada:  -3,
            esperado: 9,
        },
    }
    
    for _, caso := range casos {
        t.Run(caso.nome, func(t *testing.T) {
            resultado := Quadrado(caso.entrada)
            if resultado != caso.esperado {
                t.Errorf("Quadrado(%d) = %d; esperado %d", 
                    caso.entrada, resultado, caso.esperado)
            }
        })
    }
}
```

### Vantagens dos Table-driven Tests

1. **Fácil de adicionar novos casos**: Basta adicionar uma entrada na tabela
2. **Código DRY**: Evita duplicação de código de teste
3. **Fácil de ler**: Todos os casos estão em um lugar
4. **Execução isolada**: Cada caso roda como um subteste (usando `t.Run()`)

### Exemplo Avançado

```go
func TestValidarEmail(t *testing.T) {
    casos := []struct {
        nome     string
        email    string
        esperado bool
        erro     string
    }{
        {
            nome:     "email válido",
            email:    "usuario@exemplo.com",
            esperado: true,
            erro:     "",
        },
        {
            nome:     "email sem @",
            email:    "usuarioexemplo.com",
            esperado: false,
            erro:     "email deve conter @",
        },
        {
            nome:     "email vazio",
            email:    "",
            esperado: false,
            erro:     "email não pode ser vazio",
        },
        {
            nome:     "email sem domínio",
            email:    "usuario@",
            esperado: false,
            erro:     "email deve ter domínio válido",
        },
    }
    
    for _, caso := range casos {
        t.Run(caso.nome, func(t *testing.T) {
            valido, err := ValidarEmail(caso.email)
            
            if valido != caso.esperado {
                t.Errorf("ValidarEmail(%q) = %v; esperado %v", 
                    caso.email, valido, caso.esperado)
            }
            
            if caso.esperado == false {
                if err == nil {
                    t.Errorf("Esperado erro, mas não houve erro")
                } else if err.Error() != caso.erro {
                    t.Errorf("Erro = %q; esperado %q", err.Error(), caso.erro)
                }
            }
        })
    }
}
```

---

## Mocks e Stubs

Mocks e stubs são técnicas para isolar código durante testes, substituindo dependências por implementações controladas.

### O que são Stubs?

**Stubs** são implementações simples que retornam valores pré-definidos. Eles não verificam como foram chamados.

```go
// Interface original
type Database interface {
    GetUser(id int) (*User, error)
}

// Stub para testes
type StubDatabase struct{}

func (s *StubDatabase) GetUser(id int) (*User, error) {
    // Retorna sempre o mesmo usuário, independente do ID
    return &User{ID: 1, Name: "Test User"}, nil
}
```

### O que são Mocks?

**Mocks** são implementações que verificam como foram chamadas. Eles registram chamadas e permitem fazer asserções sobre elas.

```go
// Mock para testes
type MockDatabase struct {
    GetUserCalls []int  // Registra todas as chamadas
    GetUserFunc  func(int) (*User, error)  // Função customizável
}

func (m *MockDatabase) GetUser(id int) (*User, error) {
    m.GetUserCalls = append(m.GetUserCalls, id)
    
    if m.GetUserFunc != nil {
        return m.GetUserFunc(id)
    }
    
    // Comportamento padrão
    return &User{ID: id, Name: "Mock User"}, nil
}

// No teste
func TestService(t *testing.T) {
    mockDB := &MockDatabase{}
    service := NewService(mockDB)
    
    user, err := service.GetUser(123)
    
    // Verificar que GetUser foi chamado
    if len(mockDB.GetUserCalls) != 1 {
        t.Errorf("GetUser deveria ser chamado 1 vez, foi chamado %d vezes", 
            len(mockDB.GetUserCalls))
    }
    
    // Verificar o ID passado
    if mockDB.GetUserCalls[0] != 123 {
        t.Errorf("GetUser foi chamado com ID %d, esperado 123", 
            mockDB.GetUserCalls[0])
    }
}
```

### Exemplo Prático: Serviço com Dependência

```go
// user_service.go
package main

type UserRepository interface {
    FindByID(id int) (*User, error)
    Save(user *User) error
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
    user := &User{Name: name}
    return user, s.repo.Save(user)
}

// user_service_test.go
package main

import "testing"

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
    return nil, nil
}

func (m *MockUserRepository) Save(user *User) error {
    m.SaveCalls = append(m.SaveCalls, user)
    if m.SaveFunc != nil {
        return m.SaveFunc(user)
    }
    return nil
}

func TestUserService_GetUser(t *testing.T) {
    mockRepo := &MockUserRepository{
        FindByIDFunc: func(id int) (*User, error) {
            return &User{ID: id, Name: "Test User"}, nil
        },
    }
    
    service := NewUserService(mockRepo)
    user, err := service.GetUser(123)
    
    if err != nil {
        t.Fatalf("Erro inesperado: %v", err)
    }
    
    if user.ID != 123 {
        t.Errorf("User.ID = %d; esperado 123", user.ID)
    }
    
    if len(mockRepo.FindByIDCalls) != 1 {
        t.Errorf("FindByID deveria ser chamado 1 vez")
    }
}
```

### Bibliotecas de Mock (Opcional)

Embora você possa criar mocks manualmente, existem bibliotecas que facilitam:

- **testify/mock**: Biblioteca popular para criar mocks
- **gomock**: Gerador de mocks baseado em interfaces

---

## httptest para HTTP Tests

O pacote `net/http/httptest` fornece utilitários para testar servidores e clientes HTTP sem precisar de conexões de rede reais.

### Testando Handlers HTTP

```go
package main

import (
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestHandler(t *testing.T) {
    // Criar um request
    req, err := http.NewRequest("GET", "/hello", nil)
    if err != nil {
        t.Fatal(err)
    }
    
    // Criar um ResponseRecorder para capturar a resposta
    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(HelloHandler)
    
    // Executar o handler
    handler.ServeHTTP(rr, req)
    
    // Verificar status code
    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler retornou status %v; esperado %v", 
            status, http.StatusOK)
    }
    
    // Verificar body
    esperado := "Hello, World!"
    if rr.Body.String() != esperado {
        t.Errorf("handler retornou %q; esperado %q", 
            rr.Body.String(), esperado)
    }
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Hello, World!"))
}
```

### Criando um Servidor de Teste

```go
func TestAPIServer(t *testing.T) {
    // Criar um servidor HTTP de teste
    server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        switch r.URL.Path {
        case "/api/users":
            w.WriteHeader(http.StatusOK)
            w.Write([]byte(`[{"id": 1, "name": "Alice"}]`))
        case "/api/users/1":
            w.WriteHeader(http.StatusOK)
            w.Write([]byte(`{"id": 1, "name": "Alice"}`))
        default:
            w.WriteHeader(http.StatusNotFound)
        }
    }))
    defer server.Close()  // Importante: fechar o servidor
    
    // Fazer requisições para o servidor de teste
    resp, err := http.Get(server.URL + "/api/users")
    if err != nil {
        t.Fatalf("Erro ao fazer requisição: %v", err)
    }
    defer resp.Body.Close()
    
    if resp.StatusCode != http.StatusOK {
        t.Errorf("Status = %d; esperado %d", resp.StatusCode, http.StatusOK)
    }
}
```

### Testando Middleware

```go
func TestAuthMiddleware(t *testing.T) {
    casos := []struct {
        nome           string
        token          string
        esperadoStatus int
    }{
        {
            nome:           "token válido",
            token:          "valid-token",
            esperadoStatus: http.StatusOK,
        },
        {
            nome:           "token inválido",
            token:          "invalid-token",
            esperadoStatus: http.StatusUnauthorized,
        },
        {
            nome:           "sem token",
            token:          "",
            esperadoStatus: http.StatusUnauthorized,
        },
    }
    
    for _, caso := range casos {
        t.Run(caso.nome, func(t *testing.T) {
            req := httptest.NewRequest("GET", "/protected", nil)
            if caso.token != "" {
                req.Header.Set("Authorization", "Bearer "+caso.token)
            }
            
            rr := httptest.NewRecorder()
            handler := AuthMiddleware(ProtectedHandler)
            handler.ServeHTTP(rr, req)
            
            if rr.Code != caso.esperadoStatus {
                t.Errorf("Status = %d; esperado %d", 
                    rr.Code, caso.esperadoStatus)
            }
        })
    }
}
```

---

## Benchmarks

Benchmarks medem a performance do código executando-o múltiplas vezes e calculando estatísticas sobre o tempo de execução.

### Estrutura de um Benchmark

Um benchmark é uma função que:
- Está em um arquivo `_test.go`
- Tem o nome começando com `Benchmark`
- Recebe um parâmetro `*testing.B`

```go
func BenchmarkSoma(b *testing.B) {
    // b.N é o número de iterações que o Go vai executar
    for i := 0; i < b.N; i++ {
        Soma(2, 3)
    }
}
```

### Executando Benchmarks

```bash
# Executar todos os benchmarks
go test -bench=.

# Executar benchmark específico
go test -bench=BenchmarkSoma

# Executar benchmarks com regex
go test -bench="Benchmark.*Soma"

# Mostrar alocação de memória
go test -bench=. -benchmem

# Comparar benchmarks (útil para otimizações)
go test -bench=. -benchmem > antes.txt
# ... fazer mudanças no código ...
go test -bench=. -benchmem > depois.txt
# Comparar os arquivos
```

### Exemplo: Comparando Implementações

```go
func BenchmarkSomaSimples(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Soma(1, 2)
    }
}

func BenchmarkSomaComLoop(b *testing.B) {
    numeros := []int{1, 2, 3, 4, 5}
    for i := 0; i < b.N; i++ {
        soma := 0
        for _, n := range numeros {
            soma += n
        }
    }
}

// Output esperado:
// BenchmarkSomaSimples-8          1000000000    0.234 ns/op
// BenchmarkSomaComLoop-8           50000000     25.4 ns/op
```

### Benchmarks com Setup

```go
func BenchmarkProcessarDados(b *testing.B) {
    // Setup: executado uma vez antes do benchmark
    dados := gerarDadosGrandes(10000)
    
    // Reset do timer para não contar o setup
    b.ResetTimer()
    
    // Código que será medido
    for i := 0; i < b.N; i++ {
        ProcessarDados(dados)
    }
}

func BenchmarkProcessarDados_Paralelo(b *testing.B) {
    dados := gerarDadosGrandes(10000)
    b.ResetTimer()
    
    // Executar em paralelo
    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            ProcessarDados(dados)
        }
    })
}
```

### Sub-benchmarks

```go
func BenchmarkOperacao(b *testing.B) {
    tamanhos := []int{10, 100, 1000, 10000}
    
    for _, tamanho := range tamanhos {
        b.Run(fmt.Sprintf("tamanho-%d", tamanho), func(b *testing.B) {
            dados := gerarDados(tamanho)
            b.ResetTimer()
            
            for i := 0; i < b.N; i++ {
                Operacao(dados)
            }
        })
    }
}
```

---

## Coverage (Cobertura de Testes)

Cobertura de testes mede quanto do seu código é executado durante os testes. Isso ajuda a identificar código não testado.

### Medindo Cobertura

```bash
# Cobertura básica
go test -cover

# Cobertura detalhada com perfil
go test -coverprofile=coverage.out

# Visualizar cobertura em HTML
go tool cover -html=coverage.out

# Cobertura de múltiplos pacotes
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

### Exemplo de Output

```bash
$ go test -cover
PASS
coverage: 85.7% of statements
ok      exemplo    0.234s
```

### Interpretando Cobertura

- **0-50%**: Cobertura baixa, muitos caminhos não testados
- **50-80%**: Cobertura razoável, mas ainda há espaço para melhorias
- **80-95%**: Cobertura boa, a maioria do código está testado
- **95-100%**: Cobertura excelente, mas cuidado com obsessão por 100%

### Exemplo Prático

```go
// math.go
package main

func Soma(a, b int) int {
    return a + b
}

func Dividir(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("divisão por zero")
    }
    return a / b, nil
}

// math_test.go
package main

import "testing"

func TestSoma(t *testing.T) {
    resultado := Soma(2, 3)
    if resultado != 5 {
        t.Errorf("Esperado 5, obteve %d", resultado)
    }
}

// TestDividir não testa o caso de divisão por zero!
// Isso resultará em cobertura parcial da função Dividir
```

Após executar `go test -coverprofile=coverage.out` e `go tool cover -html=coverage.out`, você verá:
- Linhas verdes: código executado pelos testes
- Linhas vermelhas: código não executado (como o caso de erro em `Dividir`)

### Boas Práticas de Cobertura

1. **Não obceque por 100%**: Algum código pode ser difícil ou desnecessário de testar
2. **Foque em caminhos críticos**: Teste lógica de negócio importante
3. **Use cobertura como guia**: Identifique áreas que precisam de mais testes
4. **Integre no CI/CD**: Configure cobertura mínima no pipeline

---

## Resumo

Nesta aula, você aprendeu:

1. ✅ **Testing Package Basics**: Como escrever e executar testes básicos
2. ✅ **Table-driven Tests**: Estratégia poderosa para testar múltiplos cenários
3. ✅ **Mocks e Stubs**: Como isolar código substituindo dependências
4. ✅ **httptest**: Como testar código HTTP sem rede real
5. ✅ **Benchmarks**: Como medir e comparar performance
6. ✅ **Coverage**: Como medir e visualizar cobertura de testes

---

## Próximos Passos

Na próxima aula, vamos simplificar esses conceitos com analogias do dia a dia para fixar melhor o aprendizado!

Se tiver dúvidas, não hesite em perguntar. Testes são fundamentais para código de qualidade!


