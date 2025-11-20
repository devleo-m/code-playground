# Módulo 37: Escape Analysis em Detalhes
## Aula 4: Performance e Boas Práticas

Nesta aula, vamos focar em **otimizações práticas**, **boas práticas** e **armadilhas comuns** relacionadas a escape analysis em Go. Essas são lições aprendidas de projetos reais e podem fazer a diferença entre código eficiente e código com escapes desnecessários.

---

## 1. Boas Práticas de Escape Analysis

### ✅ Sempre Verifique Escape em Hot Paths

**Hot paths** são caminhos de código executados muito frequentemente. Sempre verifique escape nesses caminhos:

```bash
# Verificar escape
go build -gcflags="-m" main.go

# Filtrar apenas escapes
go build -gcflags="-m" main.go 2>&1 | grep "escape"
```

**Quando verificar:**
- ✅ Handlers HTTP
- ✅ Loops de processamento
- ✅ Funções chamadas milhões de vezes
- ✅ Código em bibliotecas de alto desempenho

### ✅ Retorne Valores Quando Possível

**❌ Ruim: Retornar Pointer Desnecessário**
```go
type Point struct {
    X, Y int
}

func getPoint() *Point {
    return &Point{X: 10, Y: 20}  // Escapa!
}
```

**✅ Bom: Retornar Valor**
```go
func getPoint() Point {
    return Point{X: 10, Y: 20}  // Não escapa!
}
```

**Regra**: Se a struct é pequena (< 100 bytes) e não precisa ser modificada, retorne valor.

### ✅ Use Tipos Concretos em Hot Paths

**❌ Ruim: Interface em Hot Path**
```go
func process(v interface{}) {  // v pode escapar
    // processar...
}
```

**✅ Bom: Tipo Concreto**
```go
func process(v int) {  // v não escapa
    // processar...
}
```

**Quando usar cada um:**
- ✅ **Tipo concreto**: Hot paths, performance crítica
- ⚠️ **Interface**: Quando realmente precisa de polimorfismo

### ✅ Pré-aloque Slices com Tamanho Conhecido

**❌ Ruim: Slice Sem Capacidade**
```go
func process(items []string) []string {
    var result []string  // Pode escapar
    for _, item := range items {
        result = append(result, process(item))
    }
    return result
}
```

**✅ Bom: Pré-alocar Capacidade**
```go
func process(items []string) []string {
    result := make([]string, 0, len(items))  // Menos provável de escapar
    for _, item := range items {
        result = append(result, process(item))
    }
    return result
}
```

**Benefício**: Reduz realocações e pode evitar escape.

---

## 2. Padrões de Otimização

### Padrão 1: Evitar Pointers em Structs Pequenas

**❌ Ruim:**
```go
type Config struct {
    Debug bool
    Port  int
}

func getConfig() *Config {
    return &Config{Debug: true, Port: 8080}  // Escapa
}
```

**✅ Bom:**
```go
func getConfig() Config {
    return Config{Debug: true, Port: 8080}  // Não escapa
}
```

**Quando usar pointer:**
- Struct é grande (> 100 bytes)
- Precisa modificar a struct
- Struct precisa ser compartilhada

### Padrão 2: Usar sync.Pool para Objetos Temporários

**❌ Ruim: Alocação Repetida**
```go
func handleRequest() {
    buf := make([]byte, 1024)  // Nova alocação, pode escapar
    // usar buf...
}
```

**✅ Bom: Pool de Buffers**
```go
var bufPool = sync.Pool{
    New: func() interface{} {
        return make([]byte, 0, 1024)
    },
}

func handleRequest() {
    buf := bufPool.Get().([]byte)
    defer bufPool.Put(buf[:0])
    // usar buf...
}
```

**Benefício**: Reutiliza buffers, reduz alocações e escapes.

### Padrão 3: Evitar Capturas Desnecessárias em Closures

**❌ Ruim: Captura de Variável Externa**
```go
func processItems(items []string) {
    for _, item := range items {
        go func() {
            fmt.Println(item)  // item escapa
        }()
    }
}
```

**✅ Bom: Cópia Local**
```go
func processItems(items []string) {
    for _, item := range items {
        item := item  // Cópia local
        go func() {
            fmt.Println(item)  // Cópia pode não escapar
        }()
    }
}
```

**Melhor ainda: Passar como Parâmetro**
```go
func processItems(items []string) {
    for _, item := range items {
        go func(it string) {
            fmt.Println(it)  // Parâmetro não escapa
        }(item)
    }
}
```

### Padrão 4: Usar strings.Builder ao Invés de Concatenação

**❌ Ruim: Concatenação de Strings**
```go
func buildMessage(parts []string) string {
    msg := ""
    for _, part := range parts {
        msg += part  // Múltiplas alocações, pode escapar
    }
    return msg
}
```

**✅ Bom: strings.Builder**
```go
func buildMessage(parts []string) string {
    var builder strings.Builder
    builder.Grow(len(parts) * 10)  // Pré-aloca
    for _, part := range parts {
        builder.WriteString(part)
    }
    return builder.String()
}
```

**Benefício**: Uma única alocação final.

---

## 3. Armadilhas Comuns

### ❌ Armadilha 1: Assumir que fmt.Println Não Escapa

```go
// ⚠️ Pode escapar em versões antigas do Go
func printValue() {
    x := 42
    fmt.Println(x)  // x pode escapar
}
```

**Solução**: Versões recentes do Go otimizaram `fmt.Println`, mas ainda pode escapar em alguns casos. Use tipos concretos quando possível.

### ❌ Armadilha 2: Retornar Slice de Array Local

```go
// ❌ Ruim: Slice pode escapar
func getSlice() []int {
    arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    return arr[:]  // Slice pode escapar
}
```

**✅ Bom: Criar Slice Diretamente**
```go
func getSlice() []int {
    return []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}  // Pode não escapar
}
```

### ❌ Armadilha 3: Interface{} em Funções Públicas

```go
// ⚠️ Pode causar escapes
func PublicFunction(v interface{}) {
    // v pode escapar
}
```

**Solução**: Use tipos concretos quando possível, ou aceite o escape se interface for necessária para flexibilidade.

### ❌ Armadilha 4: Assumir Comportamento Entre Versões

```go
// Comportamento pode mudar entre versões do Go
func example() {
    x := 42
    fmt.Println(x)  // Pode ou não escapar dependendo da versão
}
```

**Solução**: Sempre verifique com `go build -gcflags="-m"` na versão que você está usando.

---

## 4. Workflow de Otimização

### Passo 1: Identificar Hot Paths

Use profiling para identificar onde o código passa mais tempo:

```bash
go test -bench=. -cpuprofile=cpu.prof
go tool pprof cpu.prof
```

### Passo 2: Analisar Escape

```bash
go build -gcflags="-m -m" main.go 2>&1 | grep "escape"
```

### Passo 3: Priorizar Otimizações

Priorize otimizações que:
- ✅ Estão em hot paths
- ✅ Causam muitos escapes
- ✅ São fáceis de implementar
- ✅ Não prejudicam legibilidade

### Passo 4: Implementar Otimizações

Aplique as técnicas aprendidas:
- Retornar valores ao invés de pointers
- Usar tipos concretos
- Pré-alocar slices
- Usar sync.Pool

### Passo 5: Verificar Melhoria

```bash
# Antes
go test -bench=. -benchmem > before.txt

# Depois
go test -bench=. -benchmem > after.txt

# Comparar
diff before.txt after.txt
```

### Passo 6: Validar com Escape Analysis

```bash
go build -gcflags="-m" main.go 2>&1 | grep "escape" | wc -l
```

Compare o número de escapes antes e depois.

---

## 5. Exemplos Práticos de Otimização

### Exemplo 1: Handler HTTP Otimizado

**Antes:**
```go
func handler1(w http.ResponseWriter, r *http.Request) {
    data := make([]byte, 1024)  // Escapa
    // processar...
    w.Write(data)
}
```

**Depois:**
```go
var dataPool = sync.Pool{
    New: func() interface{} {
        return make([]byte, 0, 1024)
    },
}

func handler2(w http.ResponseWriter, r *http.Request) {
    data := dataPool.Get().([]byte)
    defer dataPool.Put(data[:0])
    // processar...
    w.Write(data)
}
```

**Benefício**: Reutiliza buffers, reduz alocações e escapes.

### Exemplo 2: Parser Otimizado

**Antes:**
```go
func parse1(jsonStr string) (map[string]interface{}, error) {
    var result map[string]interface{}  // Escapa
    err := json.Unmarshal([]byte(jsonStr), &result)
    return result, err
}
```

**Depois:**
```go
type MyStruct struct {
    Field1 string
    Field2 int
}

func parse2(jsonStr string) (MyStruct, error) {
    var result MyStruct  // Pode não escapar se struct for pequena
    err := json.Unmarshal([]byte(jsonStr), &result)
    return result, err
}
```

**Benefício**: Struct concreta pode evitar escape se for pequena.

### Exemplo 3: Builder Otimizado

**Antes:**
```go
func build1(parts []string) string {
    msg := ""
    for _, part := range parts {
        msg += part  // Múltiplas alocações
    }
    return msg
}
```

**Depois:**
```go
func build2(parts []string) string {
    var builder strings.Builder
    builder.Grow(len(parts) * 10)
    for _, part := range parts {
        builder.WriteString(part)
    }
    return builder.String()
}
```

**Benefício**: Uma única alocação final.

---

## 6. Quando NÃO Otimizar Escape

### ❌ Não Otimize Se:

1. **Não está em hot path**: Se código não é executado frequentemente, otimização não vale a pena
2. **Legibilidade prejudicada**: Se otimização torna código ilegível, não faça
3. **Sem dados**: Sem profiling mostrando problema, você está chutando
4. **API pública**: Mudar API para otimizar escape pode quebrar compatibilidade

### ✅ Otimize Quando:

1. **Profiling mostra problema**: Você identificou escape como gargalo
2. **Hot path identificado**: Código executado milhões de vezes
3. **Latência crítica**: Aplicações que precisam de baixa latência
4. **Fácil de implementar**: Otimização simples que não prejudica código

---

## 7. Ferramentas e Comandos Úteis

### Análise de Escape

```bash
# Básico
go build -gcflags="-m" main.go

# Detalhado
go build -gcflags="-m -m" main.go

# Máximo
go build -gcflags="-m -m -m" main.go

# Filtrar apenas escapes
go build -gcflags="-m" main.go 2>&1 | grep "escape"

# Contar escapes
go build -gcflags="-m" main.go 2>&1 | grep "escape" | wc -l
```

### Benchmark com Memória

```bash
go test -bench=. -benchmem

# Comparar antes/depois
go test -bench=. -benchmem > before.txt
# fazer mudanças
go test -bench=. -benchmem > after.txt
diff before.txt after.txt
```

### Profiling

```bash
# CPU profiling
go test -bench=. -cpuprofile=cpu.prof
go tool pprof cpu.prof

# Memory profiling
go test -bench=. -memprofile=mem.prof
go tool pprof mem.prof
```

---

## 8. Checklist de Boas Práticas

### Análise
- [ ] Identifiquei hot paths com profiling
- [ ] Analisei escape com `go build -gcflags="-m"`
- [ ] Documentei escapes encontrados
- [ ] Priorizei otimizações baseado em impacto

### Otimização
- [ ] Retorno valores ao invés de pointers quando apropriado
- [ ] Uso tipos concretos em hot paths
- [ ] Pré-aloco slices com tamanho conhecido
- [ ] Uso sync.Pool para objetos temporários
- [ ] Evito capturas desnecessárias em closures

### Validação
- [ ] Verifiquei escape após otimizações
- [ ] Comparei performance com benchmarks
- [ ] Validei que legibilidade não foi prejudicada
- [ ] Documentei otimizações feitas

### Manutenção
- [ ] Código está documentado
- [ ] Otimizações são justificadas
- [ ] Não há otimização prematura
- [ ] Código é mantível

---

## 9. Exemplo Completo: Otimização de Handler

```go
package main

import (
    "bytes"
    "encoding/json"
    "net/http"
    "sync"
)

var (
    bufferPool = sync.Pool{
        New: func() interface{} {
            return &bytes.Buffer{}
        },
    }
    
    encoderPool = sync.Pool{
        New: func() interface{} {
            return json.NewEncoder(nil)
        },
    }
)

type Response struct {
    Status string      `json:"status"`
    Data   interface{} `json:"data"`
}

// Versão otimizada: minimiza escapes
func handleAPI(w http.ResponseWriter, r *http.Request) {
    // Pool de buffers (reutilização)
    buf := bufferPool.Get().(*bytes.Buffer)
    defer bufferPool.Put(buf)
    buf.Reset()
    
    // Pool de encoders (reutilização)
    enc := encoderPool.Get().(*json.Encoder)
    defer encoderPool.Put(enc)
    enc.SetOutput(buf)
    
    // Struct pequena, retorna valor (não escapa)
    resp := Response{
        Status: "success",
        Data:   processRequest(r),
    }
    
    if err := enc.Encode(resp); err != nil {
        http.Error(w, err.Error(), 500)
        return
    }
    
    w.Header().Set("Content-Type", "application/json")
    w.Write(buf.Bytes())
}

func processRequest(r *http.Request) map[string]string {
    // Retorna valor (não pointer)
    return map[string]string{"message": "ok"}
}
```

**Otimizações aplicadas:**
- ✅ Pool de buffers (reutilização)
- ✅ Pool de encoders (reutilização)
- ✅ Retorna valores ao invés de pointers
- ✅ Minimiza escapes em hot path

---

## 10. Resumo Final

**Princípios fundamentais:**
1. **Meça antes de otimizar**: Use profiling e escape analysis para identificar problemas reais
2. **Priorize hot paths**: Foque em código executado frequentemente
3. **Retorne valores quando possível**: Evite pointers desnecessários
4. **Use tipos concretos**: Interfaces podem causar escapes
5. **Reutilize objetos**: sync.Pool para objetos temporários
6. **Valide otimizações**: Sempre compare antes/depois

**Lembre-se**: Escape analysis é uma ferramenta poderosa, mas não é uma panaceia. Escreva código limpo primeiro, otimize apenas quando necessário e baseado em dados reais.

---

**Bons estudos e happy optimizing! 🚀**



