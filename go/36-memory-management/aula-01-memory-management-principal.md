# Módulo 36: Memory Management em Profundidade

## Aula 1: Memory Management in Depth - Gerenciamento Profundo de Memória

Olá! Bem-vindo ao módulo 36 sobre **Memory Management em Profundidade**. Até agora você aprendeu o básico sobre pointers e como Go gerencia memória automaticamente. Agora vamos mergulhar nos detalhes internos: como o Go decide onde alocar memória, como funciona o Garbage Collector, e como otimizar o uso de memória em aplicações de alta performance.

Nesta aula, vamos explorar:
1. **Stack vs Heap**: Onde as variáveis são alocadas
2. **Garbage Collection**: Como Go limpa a memória automaticamente
3. **Allocation Patterns**: Padrões de alocação e seus impactos
4. **Memory Pooling**: Técnicas para reduzir alocações
5. **Otimizações**: Como minimizar pressão no GC

Esses conhecimentos são essenciais para escrever código Go eficiente e entender o comportamento de suas aplicações em produção.

---

## 1. Stack vs Heap - Onde a Memória Vive?

### O Que É Stack?

**Stack** (pilha) é uma região de memória **rápida e pequena** que armazena variáveis locais de funções. É gerenciada automaticamente pelo sistema operacional e tem acesso muito rápido.

**Características do Stack:**
- ✅ **Muito rápido**: Alocação e desalocação são instantâneas
- ✅ **Automático**: Variáveis são criadas e destruídas automaticamente quando a função termina
- ✅ **Limitado**: Tamanho fixo (geralmente 1-8MB por goroutine)
- ✅ **LIFO**: Last In, First Out (último a entrar, primeiro a sair)
- ✅ **Sem GC**: Não precisa de garbage collection

### O Que É Heap?

**Heap** (monte) é uma região de memória **maior e mais lenta** que armazena dados que precisam persistir além do escopo de uma função ou que são compartilhados entre goroutines.

**Características do Heap:**
- ⚠️ **Mais lento**: Alocação e desalocação são mais custosas
- ⚠️ **Gerenciado pelo GC**: Requer garbage collection
- ✅ **Grande**: Pode crescer conforme necessário
- ✅ **Flexível**: Dados podem viver além do escopo da função
- ⚠️ **Overhead**: Tem custo de gerenciamento

### Como Go Decide: Stack ou Heap?

O compilador Go usa **escape analysis** (análise de escape) para decidir onde alocar cada variável. Uma variável "escapa" para o heap quando:

1. **Retornada de uma função**: Se você retorna um pointer ou valor que pode ser usado fora da função
2. **Armazenada em estrutura global**: Variáveis globais sempre vão para o heap
3. **Muito grande**: Estruturas muito grandes podem ir para o heap mesmo que não escapem
4. **Compartilhada entre goroutines**: Dados compartilhados vão para o heap
5. **Tamanho desconhecido em compile-time**: Slices e maps dinâmicos frequentemente vão para o heap

### Exemplos Práticos

#### Exemplo 1: Alocação no Stack

```go
func stackAllocation() int {
    x := 42  // Alocado no stack
    return x // Retorna o valor, não o pointer
}

func main() {
    result := stackAllocation()
    fmt.Println(result) // 42
}
```

**Por que no stack?**
- `x` é uma variável local
- Retornamos o **valor**, não o pointer
- `x` não precisa existir após a função terminar
- Compilador pode garantir que não há referências após o retorno

#### Exemplo 2: Alocação no Heap

```go
func heapAllocation() *int {
    x := 42  // Escapa para o heap!
    return &x // Retornamos o pointer
}

func main() {
    result := heapAllocation()
    fmt.Println(*result) // 42
}
```

**Por que no heap?**
- Retornamos um **pointer** para `x`
- `x` precisa existir após a função terminar
- O compilador não pode garantir quando `result` será usado pela última vez
- Portanto, `x` "escapa" para o heap

#### Exemplo 3: Estrutura Grande

```go
type LargeStruct struct {
    data [1000000]int // 4MB de dados
}

func createLarge() LargeStruct {
    var s LargeStruct // Pode ir para o heap mesmo sem escape
    return s
}
```

**Por que pode ir para o heap?**
- Estruturas muito grandes podem ir para o heap mesmo que não escapem
- O stack tem tamanho limitado (1-8MB por goroutine)
- Evita estouro de stack

---

## 2. Garbage Collection (GC) - O Coletor de Lixo

### O Que É Garbage Collection?

**Garbage Collection** é o processo automático de identificar e liberar memória que não está mais sendo usada. Go tem um GC **concorrente** e **não-compactante** que roda em background.

### Como Funciona o GC do Go?

O GC do Go usa um algoritmo chamado **tri-color mark-and-sweep**:

1. **Mark Phase** (Fase de Marcação):
   - Começa de objetos "roots" (variáveis globais, stack de goroutines)
   - Marca todos os objetos alcançáveis como "vivos"
   - Usa três cores: branco (não visitado), cinza (em processamento), preto (processado)

2. **Sweep Phase** (Fase de Varredura):
   - Libera memória de objetos não marcados (mortos)
   - Retorna memória para o pool de alocação

3. **Concorrência**:
   - A maior parte do trabalho acontece em paralelo com a execução do programa
   - Pausas (STW - Stop The World) são muito curtas (< 1ms na maioria dos casos)

### Características do GC do Go

**Vantagens:**
- ✅ **Automático**: Você não precisa gerenciar memória manualmente
- ✅ **Concorrente**: Roda em paralelo com seu código
- ✅ **Baixa latência**: Pausas muito curtas
- ✅ **Eficiente**: Otimizado para baixa latência, não máxima throughput

**Desvantagens:**
- ⚠️ **Overhead**: Consome CPU e memória
- ⚠️ **Não-determinístico**: Você não controla quando roda
- ⚠️ **Pode causar latência**: Em aplicações muito sensíveis a latência

### Quando o GC Roda?

O GC é **automático** e roda quando:
- A memória heap atinge um threshold (limite)
- Você chama `runtime.GC()` manualmente (não recomendado em produção)
- O sistema detecta que há muita memória não utilizada

### Monitorando o GC

Você pode ver estatísticas do GC usando `runtime.ReadMemStats()`:

```go
package main

import (
    "fmt"
    "runtime"
)

func printGCStats() {
    var m runtime.MemStats
    runtime.ReadMemStats(&m)
    
    fmt.Printf("Alloc: %d KB\n", m.Alloc/1024)
    fmt.Printf("TotalAlloc: %d KB\n", m.TotalAlloc/1024)
    fmt.Printf("Sys: %d KB\n", m.Sys/1024)
    fmt.Printf("NumGC: %d\n", m.NumGC) // Número de GCs executados
    fmt.Printf("PauseTotalNs: %d ns\n", m.PauseTotalNs)
}

func main() {
    printGCStats()
    
    // Criar alguns objetos
    for i := 0; i < 1000; i++ {
        _ = make([]byte, 1024)
    }
    
    runtime.GC() // Forçar GC
    printGCStats()
}
```

### Variáveis de Ambiente do GC

Go permite controlar alguns aspectos do GC via variáveis de ambiente:

```bash
# Controlar o percentual de crescimento do heap antes do GC
export GOGC=100  # Padrão: 100 (GC quando heap dobra)
export GOGC=50   # GC mais frequente (quando heap cresce 50%)
export GOGC=200  # GC menos frequente (quando heap triplica)

# Desabilitar GC completamente (NÃO RECOMENDADO!)
export GOGC=off
```

**⚠️ Cuidado**: Alterar `GOGC` pode melhorar throughput mas aumentar latência.

---

## 3. Allocation Patterns - Padrões de Alocação

### O Que São Allocation Patterns?

**Allocation patterns** são padrões de código que determinam **quando** e **como** a memória é alocada. Entender esses padrões ajuda a escrever código mais eficiente.

### Padrão 1: Alocação em Loop

#### ❌ Ruim: Alocação Repetida

```go
func processItems(items []string) []string {
    var result []string
    for _, item := range items {
        result = append(result, strings.ToUpper(item))
    }
    return result
}
```

**Problema**: `append` pode realocar o slice múltiplas vezes, causando muitas alocações.

#### ✅ Bom: Pré-alocação

```go
func processItems(items []string) []string {
    result := make([]string, 0, len(items)) // Pré-aloca capacidade
    for _, item := range items {
        result = append(result, strings.ToUpper(item))
    }
    return result
}
```

**Vantagem**: Uma única alocação, sem realocações.

### Padrão 2: Pointer vs Value

#### ❌ Ruim: Retornar Struct Grande por Valor

```go
type LargeStruct struct {
    data [1000]int
}

func createLarge() LargeStruct {
    return LargeStruct{} // Cópia grande!
}
```

#### ✅ Bom: Retornar Pointer

```go
func createLarge() *LargeStruct {
    return &LargeStruct{} // Apenas pointer (8 bytes)
}
```

**Vantagem**: Evita copiar estruturas grandes.

### Padrão 3: Reutilização de Slices

#### ❌ Ruim: Criar Novo Slice Sempre

```go
func processBatch() {
    for i := 0; i < 1000; i++ {
        data := make([]byte, 1024) // Nova alocação a cada iteração
        // usar data...
    }
}
```

#### ✅ Bom: Reutilizar Slice

```go
func processBatch() {
    data := make([]byte, 1024) // Uma única alocação
    for i := 0; i < 1000; i++ {
        // reutilizar data...
        data = data[:0] // Resetar sem realocar
    }
}
```

**Vantagem**: Reduz alocações drasticamente.

---

## 4. Memory Pooling - Pool de Memória

### O Que É Memory Pooling?

**Memory pooling** é uma técnica onde você **reutiliza** objetos alocados ao invés de criar novos. Isso reduz a pressão no GC e melhora performance.

### sync.Pool - Pool Nativo do Go

Go fornece `sync.Pool` para implementar memory pooling:

```go
package main

import (
    "fmt"
    "sync"
)

var bufferPool = sync.Pool{
    New: func() interface{} {
        return make([]byte, 0, 1024) // Cria novo buffer se pool estiver vazio
    },
}

func getBuffer() []byte {
    return bufferPool.Get().([]byte)
}

func putBuffer(buf []byte) {
    buf = buf[:0] // Resetar antes de devolver
    bufferPool.Put(buf)
}

func main() {
    // Usar buffer do pool
    buf := getBuffer()
    buf = append(buf, "Hello"...)
    fmt.Println(string(buf))
    
    // Devolver para o pool
    putBuffer(buf)
    
    // Próximo uso pode reutilizar o mesmo buffer
    buf2 := getBuffer()
    fmt.Println(len(buf2)) // 0 (foi resetado)
}
```

### Como Funciona sync.Pool?

1. **Get()**: Tenta obter um objeto do pool. Se vazio, chama `New()`.
2. **Put()**: Devolve objeto ao pool para reutilização.
3. **GC**: Objetos no pool podem ser coletados pelo GC se não forem usados.

**⚠️ Importante**: `sync.Pool` não garante que você receberá o mesmo objeto. O GC pode limpar objetos não usados.

### Quando Usar sync.Pool?

**Use quando:**
- ✅ Você aloca muitos objetos do mesmo tipo
- ✅ Objetos são caros de criar
- ✅ Objetos têm vida curta
- ✅ Você quer reduzir pressão no GC

**Não use quando:**
- ❌ Objetos têm vida longa
- ❌ Você precisa garantir que objetos persistem
- ❌ Overhead de gerenciamento é maior que benefício

### Exemplo Prático: Pool de Buffers

```go
package main

import (
    "bytes"
    "fmt"
    "sync"
)

var bufPool = sync.Pool{
    New: func() interface{} {
        return &bytes.Buffer{}
    },
}

func processRequest(data string) string {
    buf := bufPool.Get().(*bytes.Buffer)
    defer bufPool.Put(buf)
    
    buf.Reset() // Limpar antes de usar
    buf.WriteString("Processed: ")
    buf.WriteString(data)
    
    return buf.String()
}

func main() {
    for i := 0; i < 10; i++ {
        result := processRequest(fmt.Sprintf("request-%d", i))
        fmt.Println(result)
    }
}
```

---

## 5. Otimizações de Memória

### Técnica 1: Reduzir Alocações

**Estratégias:**
- Pré-alocar slices com `make([]T, 0, capacity)`
- Reutilizar buffers e estruturas
- Evitar alocações em hot paths (caminhos críticos)
- Usar `sync.Pool` para objetos temporários

### Técnica 2: Reduzir Tamanho de Estruturas

```go
// ❌ Ruim: Estrutura mal alinhada
type BadStruct struct {
    a bool    // 1 byte + 7 bytes padding
    b int64   // 8 bytes
    c bool    // 1 byte + 7 bytes padding
}
// Total: 24 bytes (muito padding!)

// ✅ Bom: Estrutura bem alinhada
type GoodStruct struct {
    b int64   // 8 bytes
    a bool    // 1 byte
    c bool    // 1 byte + 6 bytes padding
}
// Total: 16 bytes (menos padding)
```

**Dica**: Coloque campos maiores primeiro para melhor alinhamento.

### Técnica 3: Usar Stack Quando Possível

- Evite retornar pointers desnecessariamente
- Use valores pequenos ao invés de pointers
- Mantenha variáveis locais quando possível

### Técnica 4: Monitorar e Profilar

Use ferramentas de profiling:

```bash
# Profiling de memória
go test -memprofile=mem.prof -bench=.

# Analisar perfil
go tool pprof mem.prof
```

---

## 6. Ferramentas de Análise

### go build -gcflags="-m"

Para ver decisões de escape analysis:

```bash
go build -gcflags="-m" main.go
```

**Output exemplo:**
```
./main.go:10:6: can inline createStack
./main.go:15:6: can inline createHeap
./main.go:15:9: &x escapes to heap
./main.go:20:6: can inline main
```

### runtime.MemStats

Já vimos como usar para estatísticas do GC.

### pprof

Ferramenta poderosa para profiling de memória:

```go
import _ "net/http/pprof"

func main() {
    go func() {
        log.Println(http.ListenAndServe("localhost:6060", nil))
    }()
    // seu código...
}
```

Acesse `http://localhost:6060/debug/pprof/heap` para ver perfil de memória.

---

## Resumo

Nesta aula aprendemos:

1. **Stack vs Heap**: Stack é rápido e automático, heap é mais lento mas flexível
2. **Garbage Collection**: GC do Go é concorrente e eficiente, mas tem overhead
3. **Allocation Patterns**: Padrões de código afetam onde e como memória é alocada
4. **Memory Pooling**: Reutilizar objetos reduz pressão no GC
5. **Otimizações**: Pré-alocação, alinhamento de estruturas, e profiling são essenciais

Na próxima aula, vamos ver **Escape Analysis** em detalhes para entender exatamente como o compilador decide onde alocar cada variável.

---

**Referências:**
- [Go Memory Model](https://go.dev/ref/mem)
- [runtime package](https://pkg.go.dev/runtime)
- [sync.Pool documentation](https://pkg.go.dev/sync#Pool)



