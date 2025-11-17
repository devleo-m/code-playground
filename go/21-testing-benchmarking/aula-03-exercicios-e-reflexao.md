# Aula 21 - Exercícios e Reflexão: Testing & Benchmarking

Olá! Agora é hora de colocar em prática o que você aprendeu sobre Testing & Benchmarking. Vamos fazer alguns exercícios práticos e depois refletir sobre os conceitos!

---

## 📝 Exercícios Práticos

### Exercício 1: Escrevendo Seu Primeiro Teste

Crie uma função `CalcularAreaRetangulo` e escreva testes para ela usando table-driven tests.

**Requisitos:**
1. A função deve receber `largura` e `altura` (ambos `float64`) e retornar a área
2. Escreva testes para pelo menos 5 casos diferentes:
   - Números positivos normais
   - Números decimais
   - Zero (deve retornar 0)
   - Números muito grandes
   - Números muito pequenos

**Template inicial:**

```go
// math.go
package main

func CalcularAreaRetangulo(largura, altura float64) float64 {
    // TODO: Implemente
    return 0
}

// math_test.go
package main

import "testing"

func TestCalcularAreaRetangulo(t *testing.T) {
    casos := []struct {
        nome           string
        largura        float64
        altura         float64
        esperado       float64
    }{
        // TODO: Adicione seus casos de teste aqui
    }
    
    for _, caso := range casos {
        t.Run(caso.nome, func(t *testing.T) {
            // TODO: Implemente o teste
        })
    }
}
```

**Tarefas:**
- [ ] Implemente a função `CalcularAreaRetangulo`
- [ ] Crie pelo menos 5 casos de teste diferentes
- [ ] Execute `go test -v` e verifique que todos os testes passam
- [ ] Execute `go test -cover` e verifique a cobertura

---

### Exercício 2: Testando Validação de Email

Crie uma função `ValidarEmail` que valida se um email é válido e escreva testes abrangentes.

**Requisitos:**
1. A função deve retornar `(bool, error)`
2. Regras de validação:
   - Deve conter exatamente um `@`
   - Deve ter parte antes do `@` (local part)
   - Deve ter parte depois do `@` (domain)
   - O domain deve conter pelo menos um ponto
   - Não pode ser string vazia
3. Use table-driven tests com pelo menos 10 casos diferentes

**Template inicial:**

```go
// email.go
package main

import "errors"

func ValidarEmail(email string) (bool, error) {
    // TODO: Implemente a validação
    return false, errors.New("não implementado")
}

// email_test.go
package main

import "testing"

func TestValidarEmail(t *testing.T) {
    casos := []struct {
        nome     string
        email    string
        esperado bool
        temErro  bool
    }{
        // TODO: Adicione casos de teste
    }
    
    for _, caso := range casos {
        t.Run(caso.nome, func(t *testing.T) {
            // TODO: Implemente o teste
        })
    }
}
```

**Tarefas:**
- [ ] Implemente a função `ValidarEmail` com todas as regras
- [ ] Crie pelo menos 10 casos de teste (válidos e inválidos)
- [ ] Teste casos extremos (strings vazias, emails muito longos, etc.)
- [ ] Execute `go test -v` e verifique que todos os testes passam

---

### Exercício 3: Criando um Mock

Crie um serviço que depende de um repositório de usuários e escreva testes usando mocks.

**Requisitos:**
1. Crie uma interface `UserRepository` com métodos:
   - `FindByID(id int) (*User, error)`
   - `Save(user *User) error`
2. Crie um `UserService` que usa o repositório
3. Crie um mock do repositório
4. Escreva testes que verificam:
   - Que os métodos do repositório são chamados corretamente
   - Que os parâmetros passados estão corretos
   - Que os erros são tratados adequadamente

**Template inicial:**

```go
// user.go
package main

type User struct {
    ID   int
    Name string
}

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
    // TODO: Implemente
    return nil, nil
}

func (s *UserService) CreateUser(name string) (*User, error) {
    // TODO: Implemente
    return nil, nil
}

// user_test.go
package main

import "testing"

type MockUserRepository struct {
    // TODO: Adicione campos para rastrear chamadas
}

func (m *MockUserRepository) FindByID(id int) (*User, error) {
    // TODO: Implemente o mock
    return nil, nil
}

func (m *MockUserRepository) Save(user *User) error {
    // TODO: Implemente o mock
    return nil
}

func TestUserService_GetUser(t *testing.T) {
    // TODO: Implemente o teste
}

func TestUserService_CreateUser(t *testing.T) {
    // TODO: Implemente o teste
}
```

**Tarefas:**
- [ ] Implemente o `UserService` com os métodos `GetUser` e `CreateUser`
- [ ] Crie um `MockUserRepository` completo que rastreia chamadas
- [ ] Escreva testes que verificam as chamadas aos métodos do mock
- [ ] Teste casos de sucesso e erro

---

### Exercício 4: Testando Handlers HTTP

Crie um handler HTTP simples e escreva testes usando `httptest`.

**Requisitos:**
1. Crie um handler `HelloHandler` que retorna "Hello, {name}!" onde `{name}` vem de um query parameter
2. Se não houver query parameter, retorna "Hello, World!"
3. Escreva testes para:
   - Handler com nome fornecido
   - Handler sem nome (deve usar "World")
   - Handler com nome vazio
   - Verificar status code correto
   - Verificar conteúdo da resposta

**Template inicial:**

```go
// handler.go
package main

import (
    "fmt"
    "net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
    // TODO: Implemente o handler
}

// handler_test.go
package main

import (
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestHelloHandler(t *testing.T) {
    casos := []struct {
        nome           string
        queryParam     string
        esperadoStatus int
        esperadoBody   string
    }{
        // TODO: Adicione casos de teste
    }
    
    for _, caso := range casos {
        t.Run(caso.nome, func(t *testing.T) {
            // TODO: Implemente o teste
        })
    }
}
```

**Tarefas:**
- [ ] Implemente o `HelloHandler` com toda a lógica
- [ ] Crie testes usando `httptest.NewRequest` e `httptest.NewRecorder`
- [ ] Verifique status code e body da resposta
- [ ] Execute `go test -v` e verifique que todos os testes passam

---

### Exercício 5: Escrevendo Benchmarks

Crie duas implementações diferentes de uma função que calcula a soma de um slice e compare-as usando benchmarks.

**Requisitos:**
1. Implementação A: Loop simples
2. Implementação B: Usando `range` com acumulador
3. Crie benchmarks para ambas
4. Teste com slices de diferentes tamanhos (10, 100, 1000, 10000 elementos)
5. Use sub-benchmarks para organizar

**Template inicial:**

```go
// soma.go
package main

func SomaLoop(slice []int) int {
    // TODO: Implemente com loop tradicional
    return 0
}

func SomaRange(slice []int) int {
    // TODO: Implemente com range
    return 0
}

// soma_test.go
package main

import "testing"

func gerarSlice(tamanho int) []int {
    slice := make([]int, tamanho)
    for i := 0; i < tamanho; i++ {
        slice[i] = i
    }
    return slice
}

func BenchmarkSomaLoop(b *testing.B) {
    tamanhos := []int{10, 100, 1000, 10000}
    
    for _, tamanho := range tamanhos {
        b.Run(fmt.Sprintf("tamanho-%d", tamanho), func(b *testing.B) {
            // TODO: Implemente o benchmark
        })
    }
}

func BenchmarkSomaRange(b *testing.B) {
    tamanhos := []int{10, 100, 1000, 10000}
    
    for _, tamanho := range tamanhos {
        b.Run(fmt.Sprintf("tamanho-%d", tamanho), func(b *testing.B) {
            // TODO: Implemente o benchmark
        })
    }
}
```

**Tarefas:**
- [ ] Implemente ambas as funções
- [ ] Crie benchmarks com sub-benchmarks para diferentes tamanhos
- [ ] Execute `go test -bench=. -benchmem`
- [ ] Compare os resultados e responda: qual é mais rápida? Por quê?

---

### Exercício 6: Medindo Cobertura

Pegue um dos exercícios anteriores e meça a cobertura de testes.

**Requisitos:**
1. Escolha uma função que você implementou (por exemplo, `ValidarEmail`)
2. Execute `go test -coverprofile=coverage.out`
3. Visualize a cobertura com `go tool cover -html=coverage.out`
4. Identifique linhas não cobertas
5. Escreva testes adicionais para aumentar a cobertura

**Tarefas:**
- [ ] Execute `go test -cover` e anote a cobertura inicial
- [ ] Gere o perfil de cobertura com `-coverprofile`
- [ ] Visualize em HTML e identifique código não testado
- [ ] Escreva testes adicionais para cobrir o código não testado
- [ ] Execute novamente e verifique o aumento da cobertura

---

## 🤔 Perguntas de Reflexão

### Reflexão 1: Por que Testes são Importantes?

Pense em um projeto real que você já trabalhou ou que gostaria de trabalhar.

**Perguntas:**
1. Qual é o custo de encontrar um bug em produção vs. encontrar durante testes?
2. Como testes ajudam na refatoração de código?
3. Testes servem como documentação? Dê exemplos.
4. Qual é a relação entre testes e confiança no código?

**Sua resposta deve ter pelo menos 3-4 parágrafos explicando:**
- O valor dos testes em projetos reais
- Como testes economizam tempo e dinheiro
- A relação entre testes e qualidade de código
- Exemplos práticos de como testes ajudam no dia a dia

---

### Reflexão 2: Table-driven Tests vs. Testes Individuais

Table-driven tests são uma prática muito comum em Go, mas nem sempre são a melhor escolha.

**Perguntas:**
1. Quando você escolheria table-driven tests em vez de testes individuais?
2. Quando testes individuais seriam mais apropriados?
3. Quais são as vantagens e desvantagens de cada abordagem?
4. Dê exemplos práticos de quando usar cada uma.

**Sua resposta deve incluir:**
- Comparação clara entre as duas abordagens
- Exemplos de quando cada uma é mais apropriada
- Vantagens e desvantagens de cada abordagem
- Sua opinião sobre quando usar cada uma

---

### Reflexão 3: Mocks e Testabilidade

Mocks permitem testar código isoladamente, mas também podem tornar testes mais complexos.

**Perguntas:**
1. Quando você deve usar mocks? Quando deve evitar?
2. Qual é a diferença entre mocks e stubs? Quando usar cada um?
3. Como mocks afetam a manutenibilidade dos testes?
4. Existe alguma situação onde testar com dependências reais é melhor?

**Sua resposta deve:**
- Explicar quando mocks são úteis vs. quando são prejudiciais
- Comparar mocks e stubs com exemplos práticos
- Discutir o trade-off entre isolamento e complexidade
- Dar exemplos de quando dependências reais são melhores

---

### Reflexão 4: Benchmarks e Otimização Prematura

Benchmarks são ferramentas poderosas, mas podem levar a otimização prematura.

**Perguntas:**
1. Quando você deve escrever benchmarks?
2. Qual é o perigo da otimização prematura?
3. Como você decide se uma otimização vale a pena?
4. Qual é a relação entre benchmarks e testes de performance reais?

**Sua resposta deve:**
- Explicar quando benchmarks são apropriados
- Discutir os perigos da otimização prematura
- Dar critérios para decidir se uma otimização vale a pena
- Comparar benchmarks com testes de performance reais

---

### Reflexão 5: Cobertura de Testes: Quanto é Suficiente?

Cobertura de testes é uma métrica importante, mas 100% de cobertura não garante código sem bugs.

**Perguntas:**
1. Qual é o nível de cobertura ideal? Por quê?
2. Por que 100% de cobertura não garante código sem bugs?
3. Que tipos de código são mais importantes de testar?
4. Como você balancearia cobertura com tempo de desenvolvimento?

**Sua resposta deve:**
- Dar uma opinião fundamentada sobre cobertura ideal
- Explicar por que 100% não é sempre necessário ou possível
- Discutir quais partes do código são mais críticas
- Dar estratégias para balancear cobertura e produtividade

---

### Reflexão 6: Testes e Cultura de Desenvolvimento

Testes não são apenas uma ferramenta técnica, mas parte de uma cultura de desenvolvimento.

**Perguntas:**
1. Como você convenceria um time a escrever mais testes?
2. Qual é o papel dos testes em um processo de desenvolvimento ágil?
3. Como testes se relacionam com CI/CD?
4. Qual é a importância de testes em projetos open source?

**Sua resposta deve:**
- Discutir a importância cultural dos testes
- Explicar como testes se integram em processos ágeis
- Descrever a relação entre testes e CI/CD
- Dar exemplos de como testes melhoram colaboração

---

## ✅ Checklist de Entrega

Antes de enviar suas respostas, verifique:

- [ ] Todos os exercícios práticos foram implementados e testados
- [ ] Todos os códigos foram executados com `go test` ou `go test -v`
- [ ] Todos os testes passam sem erros
- [ ] Benchmarks foram executados e resultados foram analisados
- [ ] Cobertura foi medida e visualizada
- [ ] Todas as perguntas de reflexão foram respondidas com profundidade
- [ ] As respostas de reflexão têm pelo menos 3-4 parágrafos cada
- [ ] Você incluiu exemplos práticos nas respostas de reflexão

---

## 📚 Recursos Adicionais (Opcional)

Se quiser se aprofundar mais:

1. Leia a documentação oficial: `go doc testing`
2. Explore o código-fonte do pacote `testing` (se tiver curiosidade)
3. Experimente bibliotecas de mock como `testify/mock` ou `gomock`
4. Pratique escrevendo testes para código existente
5. Explore ferramentas de cobertura mais avançadas

---

## 💡 Dicas Finais

1. **Comece simples**: Não tente testar tudo de uma vez. Comece com funções simples e vá evoluindo.

2. **Teste comportamento, não implementação**: Foque no que a função faz, não em como ela faz.

3. **Mantenha testes simples**: Testes devem ser fáceis de ler e entender. Se um teste é complexo, talvez o código que está testando também seja.

4. **Use table-driven tests**: Eles são uma das práticas mais recomendadas em Go.

5. **Não obceque por 100% de cobertura**: Foque em testar código crítico e lógica de negócio.

6. **Benchmarks são ferramentas de medição**: Use-os para tomar decisões baseadas em dados, não em suposições.

---

Boa sorte com os exercícios! Lembre-se: a prática é essencial para dominar testes em Go. Testes são uma habilidade que melhora com o tempo e a experiência! 🚀

Envie suas respostas quando estiver pronto, e eu farei uma análise detalhada do seu desempenho!


