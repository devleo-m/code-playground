# Aula 21 - Performance e Boas Práticas: Testing & Benchmarking

Olá! Agora que você entende os conceitos de Testing & Benchmarking, é crucial aprender **quando e como** usá-los de forma eficiente e correta. Nesta aula, vamos explorar aspectos de performance, boas práticas, anti-padrões e os erros comuns que você deve evitar.

---

## 🚀 Performance: Otimizando Testes

### Testes Rápidos vs. Testes Completos

**Fato importante:** Testes devem ser rápidos o suficiente para serem executados frequentemente.

**Custos típicos de testes lentos:**
- Desenvolvedores evitam executar testes
- Feedback lento atrasa o desenvolvimento
- CI/CD demora muito para completar
- Desenvolvimento menos ágil

**Estratégias para testes rápidos:**

```go
// ❌ LENTO: Teste que faz I/O real
func TestDownloadFile(t *testing.T) {
    data, err := http.Get("https://exemplo.com/arquivo.txt")
    // ... lento e depende de rede
}

// ✅ RÁPIDO: Teste com httptest
func TestDownloadFile(t *testing.T) {
    server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("conteúdo"))
    }))
    defer server.Close()
    
    data, err := http.Get(server.URL + "/arquivo.txt")
    // ... rápido e não depende de rede
}
```

**Regra geral:**
- ✅ **Sempre** use mocks/stubs para dependências externas
- ✅ **Sempre** use `httptest` para testes HTTP
- ✅ **Sempre** evite I/O real (arquivos, banco de dados, rede)
- ⚠️ **Considere** testes de integração separados para I/O real

---

## ⚡ Performance: Benchmarks Eficientes

### Evitando Armadilhas Comuns em Benchmarks

**1. Setup não contado**

```go
// ❌ ERRADO: Setup conta no tempo do benchmark
func BenchmarkOperacao(b *testing.B) {
    dados := gerarDadosGrandes(100000)  // Setup lento!
    
    for i := 0; i < b.N; i++ {
        Operacao(dados)
    }
}

// ✅ CORRETO: Reset timer após setup
func BenchmarkOperacao(b *testing.B) {
    dados := gerarDadosGrandes(100000)
    b.ResetTimer()  // Reset para não contar o setup
    
    for i := 0; i < b.N; i++ {
        Operacao(dados)
    }
}
```

**2. Alocação de memória não medida**

```go
// ❌ ERRADO: Não mede alocação
func BenchmarkSoma(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Soma(1, 2)
    }
}
// go test -bench=.  (sem -benchmem)

// ✅ CORRETO: Mede alocação também
// go test -bench=. -benchmem
// Output inclui: 0 B/op, 0 allocs/op
```

**3. Otimizações do compilador**

```go
// ❌ ERRADO: Compilador pode otimizar demais
func BenchmarkSoma(b *testing.B) {
    for i := 0; i < b.N; i++ {
        resultado := Soma(1, 2)
        _ = resultado  // Compilador pode remover tudo!
    }
}

// ✅ CORRETO: Usar resultado de forma que não seja otimizado
func BenchmarkSoma(b *testing.B) {
    var resultado int
    for i := 0; i < b.N; i++ {
        resultado = Soma(1, 2)
    }
    _ = resultado  // Previne otimização
}
```

**4. Benchmarks paralelos**

```go
// ✅ ÚTIL: Para testar código concorrente
func BenchmarkOperacao_Paralelo(b *testing.B) {
    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            Operacao()
        }
    })
}

// Use quando:
// - Código é thread-safe
// - Quer testar performance sob concorrência
// - Quer encontrar race conditions
```

---

## 📊 Performance: Cobertura Eficiente

### Balanceando Cobertura e Tempo

**Fato importante:** 100% de cobertura não é sempre necessário ou viável.

**Estratégias de cobertura:**

```go
// ✅ FOCAR: Código crítico primeiro
func TestTransferenciaBancaria(t *testing.T) {
    // Teste crítico - deve ter cobertura alta
    casos := []struct {
        // ... muitos casos
    }{
        // Teste todos os caminhos críticos
    }
}

// ⚠️ ACEITÁVEL: Código simples pode ter cobertura menor
func TestHelperFunction(t *testing.T) {
    // Função helper simples - cobertura básica é suficiente
    resultado := Helper("test")
    if resultado != "test" {
        t.Error("falhou")
    }
}
```

**Regras de cobertura:**
- ✅ **Sempre** teste lógica de negócio crítica (100% ideal)
- ✅ **Sempre** teste caminhos de erro importantes
- ⚠️ **Considere** cobertura menor para código simples/helper
- ❌ **Nunca** obceque por 100% se não for crítico

---

## 🎯 Boas Práticas: Estrutura de Testes

### Organizando Testes de Forma Clara

**1. Nomes descritivos**

```go
// ❌ RUIM: Nome vago
func Test1(t *testing.T) { }

// ✅ BOM: Nome descritivo
func TestCalcularAreaRetangulo_ComNumerosPositivos(t *testing.T) { }

// ✅ MELHOR: Nome que descreve comportamento
func TestCalcularAreaRetangulo_RetornaAreaCorreta_QuandoRecebeLarguraEAltura(t *testing.T) { }
```

**2. Um conceito por teste**

```go
// ❌ RUIM: Múltiplos conceitos
func TestUsuario(t *testing.T) {
    // Testa criação
    // Testa atualização
    // Testa deleção
    // Testa validação
}

// ✅ BOM: Um conceito por teste
func TestUsuario_Criar_RetornaUsuarioValido(t *testing.T) { }
func TestUsuario_Atualizar_AtualizaCamposCorretos(t *testing.T) { }
func TestUsuario_Deletar_RemoveDoBanco(t *testing.T) { }
func TestUsuario_Validar_RetornaErroParaEmailInvalido(t *testing.T) { }
```

**3. Arrange-Act-Assert pattern**

```go
// ✅ BOM: Estrutura clara
func TestSoma(t *testing.T) {
    // Arrange: Preparar dados
    a, b := 2, 3
    esperado := 5
    
    // Act: Executar ação
    resultado := Soma(a, b)
    
    // Assert: Verificar resultado
    if resultado != esperado {
        t.Errorf("Soma(%d, %d) = %d; esperado %d", 
            a, b, resultado, esperado)
    }
}
```

---

## 🎯 Boas Práticas: Table-driven Tests

### Quando e Como Usar

**✅ Use table-driven tests quando:**
- Múltiplos casos de teste com mesma lógica
- Variações de entrada/saída esperada
- Testes de validação com diferentes inputs
- Testes de formatação/parsing

**✅ Estrutura recomendada:**

```go
func TestValidarEmail(t *testing.T) {
    casos := []struct {
        nome     string  // Nome descritivo do caso
        entrada  string  // Input do teste
        esperado bool    // Output esperado
        erro     string  // Erro esperado (se houver)
    }{
        {
            nome:     "email válido com domínio comum",
            entrada:  "usuario@exemplo.com",
            esperado: true,
            erro:     "",
        },
        {
            nome:     "email sem @ deve falhar",
            entrada:  "usuarioexemplo.com",
            esperado: false,
            erro:     "email deve conter @",
        },
        // ... mais casos
    }
    
    for _, caso := range casos {
        t.Run(caso.nome, func(t *testing.T) {
            // Teste isolado para cada caso
            resultado, err := ValidarEmail(caso.entrada)
            
            if resultado != caso.esperado {
                t.Errorf("ValidarEmail(%q) = %v; esperado %v", 
                    caso.entrada, resultado, caso.esperado)
            }
            
            if caso.esperado == false && err == nil {
                t.Error("Esperado erro, mas não houve erro")
            }
        })
    }
}
```

**❌ Evite table-driven tests quando:**
- Cada caso precisa de setup muito diferente
- Lógica de teste é muito complexa
- Apenas 1-2 casos de teste simples

---

## 🎯 Boas Práticas: Mocks e Stubs

### Criando Mocks Eficientes

**1. Mocks simples e focados**

```go
// ❌ RUIM: Mock muito complexo
type MockComplexo struct {
    // 50 campos diferentes
    // Lógica complexa de verificação
}

// ✅ BOM: Mock focado no que precisa
type MockUserRepository struct {
    FindByIDCalls []int
    FindByIDFunc  func(int) (*User, error)
}

func (m *MockUserRepository) FindByID(id int) (*User, error) {
    m.FindByIDCalls = append(m.FindByIDCalls, id)
    if m.FindByIDFunc != nil {
        return m.FindByIDFunc(id)
    }
    return &User{ID: id}, nil
}
```

**2. Verificações claras**

```go
// ✅ BOM: Verificações explícitas
func TestService_GetUser(t *testing.T) {
    mock := &MockUserRepository{}
    service := NewService(mock)
    
    user, err := service.GetUser(123)
    
    // Verificações claras
    if err != nil {
        t.Fatalf("Erro inesperado: %v", err)
    }
    
    if len(mock.FindByIDCalls) != 1 {
        t.Errorf("FindByID deveria ser chamado 1 vez, foi %d", 
            len(mock.FindByIDCalls))
    }
    
    if mock.FindByIDCalls[0] != 123 {
        t.Errorf("FindByID chamado com %d, esperado 123", 
            mock.FindByIDCalls[0])
    }
}
```

**3. Quando não usar mocks**

```go
// ⚠️ CONSIDERE: Testar com implementação real se for simples
type InMemoryRepository struct {
    users map[int]*User
}

func (r *InMemoryRepository) FindByID(id int) (*User, error) {
    return r.users[id], nil
}

// Às vezes é melhor usar implementação real simples
// do que criar um mock complexo
```

---

## 🎯 Boas Práticas: Testes HTTP

### Testando Handlers Eficientemente

**1. Teste comportamento, não implementação**

```go
// ❌ RUIM: Testa detalhes de implementação
func TestHandler(t *testing.T) {
    // Verifica se usa json.Marshal especificamente
    // Verifica ordem dos campos
    // Muito acoplado à implementação
}

// ✅ BOM: Testa comportamento
func TestHandler(t *testing.T) {
    req := httptest.NewRequest("GET", "/users/1", nil)
    rr := httptest.NewRecorder()
    
    handler.ServeHTTP(rr, req)
    
    // Verifica o que importa: status e conteúdo
    if rr.Code != http.StatusOK {
        t.Errorf("Status = %d; esperado %d", rr.Code, http.StatusOK)
    }
    
    var user User
    json.Unmarshal(rr.Body.Bytes(), &user)
    if user.ID != 1 {
        t.Errorf("User.ID = %d; esperado 1", user.ID)
    }
}
```

**2. Teste casos de erro**

```go
// ✅ BOM: Testa tanto sucesso quanto erro
func TestGetUser(t *testing.T) {
    casos := []struct {
        nome           string
        userID         string
        esperadoStatus int
    }{
        {
            nome:           "usuário existe",
            userID:         "1",
            esperadoStatus: http.StatusOK,
        },
        {
            nome:           "usuário não existe",
            userID:         "999",
            esperadoStatus: http.StatusNotFound,
        },
        {
            nome:           "ID inválido",
            userID:         "abc",
            esperadoStatus: http.StatusBadRequest,
        },
    }
    
    for _, caso := range casos {
        t.Run(caso.nome, func(t *testing.T) {
            // ... teste
        })
    }
}
```

---

## 🎯 Boas Práticas: Benchmarks

### Escrevendo Benchmarks Úteis

**1. Compare implementações**

```go
// ✅ BOM: Compara diferentes abordagens
func BenchmarkSoma_Loop(b *testing.B) {
    slice := gerarSlice(1000)
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        SomaLoop(slice)
    }
}

func BenchmarkSoma_Range(b *testing.B) {
    slice := gerarSlice(1000)
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        SomaRange(slice)
    }
}

// Execute: go test -bench=. -benchmem
// Compare os resultados
```

**2. Teste diferentes tamanhos**

```go
// ✅ BOM: Sub-benchmarks para diferentes tamanhos
func BenchmarkProcessar(b *testing.B) {
    tamanhos := []int{10, 100, 1000, 10000}
    
    for _, tamanho := range tamanhos {
        b.Run(fmt.Sprintf("tamanho-%d", tamanho), func(b *testing.B) {
            dados := gerarDados(tamanho)
            b.ResetTimer()
            
            for i := 0; i < b.N; i++ {
                Processar(dados)
            }
        })
    }
}
```

**3. Use resultados para decisões**

```go
// ✅ BOM: Benchmark informa decisão de design
// Se BenchmarkA é 2x mais rápido que BenchmarkB,
// considere usar implementação A em código crítico
```

---

## ❌ Anti-padrões Comuns

### Erros que Você Deve Evitar

**1. Testes que dependem de ordem**

```go
// ❌ RUIM: Testes dependem de ordem de execução
var contadorGlobal int

func TestA(t *testing.T) {
    contadorGlobal = 10
    // ...
}

func TestB(t *testing.T) {
    // Depende de TestA ter executado antes!
    if contadorGlobal != 10 {
        t.Error("falhou")
    }
}

// ✅ BOM: Cada teste é independente
func TestA(t *testing.T) {
    contador := 10
    // Usa variável local
}

func TestB(t *testing.T) {
    contador := 10
    // Independente de TestA
}
```

**2. Testes que não testam nada**

```go
// ❌ RUIM: Teste que sempre passa
func TestSoma(t *testing.T) {
    Soma(2, 3)
    // Sem verificação!
}

// ✅ BOM: Teste verifica resultado
func TestSoma(t *testing.T) {
    resultado := Soma(2, 3)
    if resultado != 5 {
        t.Errorf("Esperado 5, obteve %d", resultado)
    }
}
```

**3. Testes muito complexos**

```go
// ❌ RUIM: Teste difícil de entender
func TestComplexo(t *testing.T) {
    // 200 linhas de código
    // Múltiplas responsabilidades
    // Difícil de debugar quando falha
}

// ✅ BOM: Testes simples e focados
func TestCaso1(t *testing.T) {
    // Teste focado em um caso
}

func TestCaso2(t *testing.T) {
    // Teste focado em outro caso
}
```

**4. Ignorar erros em testes**

```go
// ❌ RUIM: Ignora erros
func TestOperacao(t *testing.T) {
    resultado, err := Operacao()
    _ = err  // Ignora erro!
    // ...
}

// ✅ BOM: Verifica erros
func TestOperacao(t *testing.T) {
    resultado, err := Operacao()
    if err != nil {
        t.Fatalf("Erro inesperado: %v", err)
    }
    // ...
}
```

---

## 📋 Checklist de Boas Práticas

Antes de considerar seus testes "prontos", verifique:

### Estrutura
- [ ] Nomes de testes são descritivos
- [ ] Cada teste tem uma responsabilidade clara
- [ ] Testes são independentes (não dependem de ordem)
- [ ] Usa Arrange-Act-Assert quando apropriado

### Table-driven Tests
- [ ] Usa table-driven tests quando há múltiplos casos similares
- [ ] Cada caso tem nome descritivo
- [ ] Usa `t.Run()` para isolar casos

### Mocks e Stubs
- [ ] Mocks são simples e focados
- [ ] Verificações são claras e explícitas
- [ ] Usa mocks apenas quando necessário

### Testes HTTP
- [ ] Usa `httptest` em vez de servidores reais
- [ ] Testa casos de sucesso e erro
- [ ] Verifica status code e conteúdo

### Benchmarks
- [ ] Setup não é contado no tempo (usa `ResetTimer`)
- [ ] Executa com `-benchmem` para medir alocação
- [ ] Compara implementações quando apropriado

### Cobertura
- [ ] Mede cobertura regularmente
- [ ] Foca em código crítico primeiro
- [ ] Não obceca por 100% se não for necessário

### Performance
- [ ] Testes são rápidos o suficiente
- [ ] Usa mocks para dependências lentas
- [ ] Evita I/O real em testes unitários

---

## 🎓 Resumo

Nesta aula, você aprendeu:

1. ✅ **Performance de Testes**: Como manter testes rápidos
2. ✅ **Benchmarks Eficientes**: Como escrever benchmarks úteis
3. ✅ **Cobertura Balanceada**: Quando focar em cobertura
4. ✅ **Estrutura de Testes**: Como organizar testes claramente
5. ✅ **Table-driven Tests**: Quando e como usar
6. ✅ **Mocks Eficientes**: Como criar mocks simples e úteis
7. ✅ **Testes HTTP**: Como testar handlers eficientemente
8. ✅ **Anti-padrões**: Erros comuns a evitar

---

## 💡 Dica Final

Lembre-se: **bons testes são uma forma de documentação viva**. Eles devem:
- Ser fáceis de ler e entender
- Executar rapidamente
- Ser confiáveis (não falham aleatoriamente)
- Testar comportamento, não implementação
- Ser mantidos junto com o código

Testes são um investimento que paga dividendos ao longo do tempo! 🚀

