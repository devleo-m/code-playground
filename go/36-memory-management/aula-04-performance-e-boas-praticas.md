# Módulo 36: Memory Management em Profundidade
## Aula 4: Performance e Boas Práticas

Nesta aula, vamos focar em **otimizações práticas**, **boas práticas** e **armadilhas comuns** relacionadas a gerenciamento de memória em Go. Essas são lições aprendidas de projetos reais e podem fazer a diferença entre uma aplicação lenta e uma aplicação eficiente.

---

## 1. Boas Práticas de Alocação

### ✅ Sempre Pré-aloque Slices Quando Possível

**❌ Ruim:**
```go
func processItems(items []string) []string {
    var result []string  // Capacidade zero
    for _, item := range items {
        result = append(result, process(item))  // Múltiplas realocações!
    }
    return result
}
```

**Problema**: `append` pode realocar o slice múltiplas vezes, causando:
- Múltiplas alocações
- Cópias de dados
- Pressão no GC

**✅ Bom:**
```go
func processItems(items []string) []string {
    result := make([]string, 0, len(items))  // Pré-aloca capacidade
    for _, item := range items {
        result = append(result, process(item))  // Sem realocações!
    }
    return result
}
```

**Benefício**: Uma única alocação, sem cópias desnecessárias.

### ✅ Use `strings.Builder` para Múltiplas Concatenações

**❌ Ruim:**
```go
func buildMessage(parts []string) string {
    msg := ""
    for _, part := range parts {
        msg += part  // Nova string a cada concatenação!
    }
    return msg
}
```

**Problema**: Strings são imutáveis em Go. Cada `+=` cria uma nova string, copiando tudo.

**✅ Bom:**
```go
func buildMessage(parts []string) string {
    var builder strings.Builder
    builder.Grow(len(parts) * 10)  // Pré-aloca espaço estimado
    for _, part := range parts {
        builder.WriteString(part)
    }
    return builder.String()
}
```

**Benefício**: Uma única alocação final, muito mais eficiente.

### ✅ Reutilize Slices em Loops

**❌ Ruim:**
```go
func processBatch(items []Item) {
    for _, item := range items {
        buffer := make([]byte, 1024)  // Nova alocação a cada iteração
        process(item, buffer)
    }
}
```

**✅ Bom:**
```go
func processBatch(items []Item) {
    buffer := make([]byte, 0, 1024)  // Uma única alocação
    for _, item := range items {
        buffer = buffer[:0]  // Resetar sem realocar
        process(item, buffer)
    }
}
```

**Benefício**: Reduz alocações drasticamente em loops.

---

## 2. Quando Usar sync.Pool

### ✅ Use sync.Pool Para Objetos Temporários e Caros

**Cenários ideais:**
- Buffers (`*bytes.Buffer`, `[]byte`)
- Parsers temporários
- Objetos que são criados frequentemente e descartados rapidamente
- Objetos que são caros de criar

**Exemplo prático:**
```go
var bufferPool = sync.Pool{
    New: func() interface{} {
        return &bytes.Buffer{}
    },
}

func processRequest(data string) string {
    buf := bufferPool.Get().(*bytes.Buffer)
    defer bufferPool.Put(buf)
    
    buf.Reset()  // CRÍTICO: Resetar antes de usar
    buf.WriteString("Response: ")
    buf.WriteString(data)
    
    return buf.String()
}
```

### ❌ NÃO Use sync.Pool Para:

- Objetos com estado que precisa persistir
- Objetos que são usados por muito tempo
- Objetos muito baratos de criar
- Quando o overhead de gerenciamento é maior que o benefício

**Exemplo de uso errado:**
```go
// ❌ ERRADO: int é muito barato, não vale o overhead
var intPool = sync.Pool{
    New: func() interface{} { return new(int) },
}
```

---

## 3. Otimização de Estruturas

### ✅ Organize Campos por Tamanho (Alinhamento)

**❌ Ruim: Alinhamento Ineficiente**
```go
type Inefficient struct {
    a bool    // 1 byte + 7 bytes padding
    b int64   // 8 bytes
    c bool    // 1 byte + 7 bytes padding
    d int32   // 4 bytes + 4 bytes padding
}
// Total: 32 bytes (12 bytes desperdiçados!)
```

**✅ Bom: Campos Organizados**
```go
type Efficient struct {
    b int64   // 8 bytes
    d int32   // 4 bytes
    a bool    // 1 byte
    c bool    // 1 byte + 2 bytes padding
}
// Total: 16 bytes (50% menor!)
```

**Regra**: Coloque campos maiores primeiro para minimizar padding.

### ✅ Use Tipos Menores Quando Possível

**❌ Ruim:**
```go
type User struct {
    Age int64  // 8 bytes para idade (0-150)
}
```

**✅ Bom:**
```go
type User struct {
    Age uint8  // 1 byte é suficiente para idade
}
```

**Benefício**: Menos memória, melhor cache locality.

---

## 4. Reduzindo Pressão no GC

### ✅ Reduza Alocações em Hot Paths

**Hot path**: Caminho de código executado muito frequentemente (ex: dentro de loops, handlers HTTP).

**Estratégias:**
1. Pré-alocar estruturas
2. Reutilizar buffers
3. Usar `sync.Pool` para objetos temporários
4. Evitar alocações desnecessárias em condições

**Exemplo:**
```go
// ❌ Ruim: Alocação no hot path
func handleRequest(w http.ResponseWriter, r *http.Request) {
    data := make([]byte, 1024)  // Alocação a cada request!
    // ...
}

// ✅ Bom: Pool de buffers
var bufferPool = sync.Pool{
    New: func() interface{} { return make([]byte, 0, 1024) },
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
    buf := bufferPool.Get().([]byte)
    defer bufferPool.Put(buf[:0])
    // ...
}
```

### ✅ Evite Retornar Pointers Desnecessários

**❌ Ruim:**
```go
func getValue() *int {
    x := 42
    return &x  // Escapa para heap desnecessariamente
}
```

**✅ Bom:**
```go
func getValue() int {
    x := 42
    return x  // Fica no stack
}
```

**Regra**: Só retorne pointer se realmente precisar compartilhar ou modificar.

---

## 5. Monitoramento e Profiling

### ✅ Use pprof para Identificar Problemas

**Habilitar pprof:**
```go
import _ "net/http/pprof"

func main() {
    go func() {
        log.Println(http.ListenAndServe("localhost:6060", nil))
    }()
    // seu código...
}
```

**Acessar perfis:**
- Heap: `http://localhost:6060/debug/pprof/heap`
- Allocs: `http://localhost:6060/debug/pprof/allocs`
- Goroutine: `http://localhost:6060/debug/pprof/goroutine`

**Analisar:**
```bash
go tool pprof http://localhost:6060/debug/pprof/heap
```

### ✅ Use Benchmarks com -benchmem

```go
func BenchmarkExample(b *testing.B) {
    for i := 0; i < b.N; i++ {
        // seu código...
    }
}
```

```bash
go test -bench=. -benchmem
```

**Output mostra:**
- Tempo de execução
- Número de alocações
- Bytes alocados

### ✅ Monitore GC em Produção

```go
import (
    "runtime"
    "time"
)

func monitorGC() {
    var m runtime.MemStats
    ticker := time.NewTicker(5 * time.Second)
    defer ticker.Stop()
    
    for range ticker.C {
        runtime.ReadMemStats(&m)
        log.Printf("GC runs: %d, Pause total: %d ms",
            m.NumGC,
            m.PauseTotalNs/1e6)
    }
}
```

---

## 6. Armadilhas Comuns

### ❌ Armadilha 1: Esquecer de Resetar em sync.Pool

```go
// ❌ ERRADO
buf := pool.Get().(*bytes.Buffer)
defer pool.Put(buf)
// Esqueceu de Reset! Dados antigos podem vazar!

// ✅ CORRETO
buf := pool.Get().(*bytes.Buffer)
defer pool.Put(buf)
buf.Reset()  // Sempre resetar!
```

### ❌ Armadilha 2: Assumir que sync.Pool Mantém Objetos

```go
// ❌ ERRADO: Assumir que objeto persiste
obj := pool.Get()
pool.Put(obj)
// GC pode limpar obj do pool!
obj2 := pool.Get()  // Pode ser nil ou diferente!
```

**Solução**: Sempre verifique se objeto do pool é válido ou use `New`.

### ❌ Armadilha 3: Otimização Prematura

```go
// ❌ ERRADO: Pool para algo muito simples
var stringPool = sync.Pool{
    New: func() interface{} { return "" },
}
// Overhead maior que benefício!
```

**Regra**: Meça antes de otimizar. Use profiling para identificar problemas reais.

### ❌ Armadilha 4: Ignorar Escape Analysis

```go
// ❌ Pode não perceber que escapa
func create() *LargeStruct {
    return &LargeStruct{}  // Escapa para heap!
}
```

**Solução**: Use `go build -gcflags="-m"` regularmente.

---

## 7. Checklist de Boas Práticas

### Alocação
- [ ] Pré-aloco slices quando sei o tamanho aproximado
- [ ] Uso `strings.Builder` para múltiplas concatenações
- [ ] Reutilizo slices em loops quando possível
- [ ] Evito retornar pointers desnecessários

### sync.Pool
- [ ] Uso pool apenas para objetos temporários e caros
- [ ] Sempre reseto objetos antes de devolver ao pool
- [ ] Não assumo que objetos persistem no pool
- [ ] Meço o benefício antes de usar pool

### Estruturas
- [ ] Organizo campos por tamanho (maiores primeiro)
- [ ] Uso tipos menores quando possível
- [ ] Minimizo padding desnecessário

### Monitoramento
- [ ] Uso `-benchmem` em benchmarks
- [ ] Uso pprof para identificar problemas
- [ ] Monitorei GC em produção
- [ ] Verifico escape analysis regularmente

### Performance
- [ ] Reduzi alocações em hot paths
- [ ] Evitei otimização prematura
- [ ] Meço antes de otimizar
- [ ] Documentei otimizações complexas

---

## 8. Quando NÃO Otimizar

### ❌ Não Otimize Se:

1. **Não há problema de performance**: Se a aplicação está rápida o suficiente, não otimize.
2. **Não mediu**: Sem dados, você está chutando. Use profiling.
3. **Código fica ilegível**: Legibilidade > micro-otimizações.
4. **Overhead maior que benefício**: Algumas otimizações têm custo maior que ganho.

### ✅ Otimize Quando:

1. **Profiling mostra problema**: Você identificou gargalo real.
2. **Aplicação é crítica**: Latência ou throughput são importantes.
3. **Escala alta**: Milhões de operações por segundo.
4. **Recursos limitados**: Memória ou CPU são restritos.

---

## 9. Exemplo Completo: Handler HTTP Otimizado

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
    Status string `json:"status"`
    Data   interface{} `json:"data"`
}

func handleAPI(w http.ResponseWriter, r *http.Request) {
    // Obter buffer do pool
    buf := bufferPool.Get().(*bytes.Buffer)
    defer bufferPool.Put(buf)
    buf.Reset()  // CRÍTICO: Resetar
    
    // Obter encoder do pool
    enc := encoderPool.Get().(*json.Encoder)
    defer encoderPool.Put(enc)
    enc.SetOutput(buf)
    
    // Criar resposta
    resp := Response{
        Status: "success",
        Data:   processRequest(r),
    }
    
    // Encodar
    if err := enc.Encode(resp); err != nil {
        http.Error(w, err.Error(), 500)
        return
    }
    
    // Escrever resposta
    w.Header().Set("Content-Type", "application/json")
    w.Write(buf.Bytes())
}

func processRequest(r *http.Request) interface{} {
    // Processar request...
    return map[string]string{"message": "ok"}
}
```

**Otimizações aplicadas:**
- ✅ Pool de buffers (reutilização)
- ✅ Pool de encoders (reutilização)
- ✅ Reset antes de usar
- ✅ Evita alocações em hot path

---

## 10. Recursos Adicionais

### Ferramentas
- **pprof**: `go tool pprof`
- **trace**: `go tool trace`
- **escape analysis**: `go build -gcflags="-m"`

### Documentação
- [Go Memory Model](https://go.dev/ref/mem)
- [runtime package](https://pkg.go.dev/runtime)
- [sync.Pool](https://pkg.go.dev/sync#Pool)

### Artigos Recomendados
- "Understanding Go's Memory Allocator" (blog oficial)
- "Go GC: Prioritizing low latency" (blog oficial)
- "Escape Analysis in Go" (vários artigos)

---

## Resumo Final

**Princípios fundamentais:**
1. **Meça antes de otimizar**: Use profiling para identificar problemas reais
2. **Reduza alocações**: Pré-aloque, reutilize, use pools quando apropriado
3. **Monitore GC**: Entenda o comportamento do GC na sua aplicação
4. **Evite otimização prematura**: Legibilidade e simplicidade primeiro
5. **Use ferramentas**: pprof, benchmarks, escape analysis

**Lembre-se**: A melhor otimização é a que você não precisa fazer. Escreva código simples e claro primeiro. Otimize apenas quando necessário e baseado em dados reais.

---

**Bons estudos e happy optimizing! 🚀**



