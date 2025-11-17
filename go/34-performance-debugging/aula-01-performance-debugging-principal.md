# Módulo 34: Performance e Debugging em Go
## Aula 1: pprof, trace e Race Detector - Ferramentas Essenciais para Análise de Performance

Olá! Bem-vindo a este módulo crucial sobre **análise de performance e debugging** em Go. Até agora você aprendeu a escrever código Go funcional, mas em projetos profissionais, é essencial saber **identificar gargalos**, **otimizar performance** e **encontrar bugs de concorrência**.

Nesta aula, vamos mergulhar em três ferramentas poderosas do Go:
1. **`pprof`**: Para análise de performance (CPU, memória, goroutines)
2. **`trace`**: Para análise de execução e concorrência
3. **Race Detector**: Para detectar data races em programas concorrentes

Essas ferramentas são fundamentais para criar aplicações Go rápidas, eficientes e livres de bugs de concorrência.

---

## 1. pprof - Profiling de Performance

### O Que É pprof?

O **`pprof`** é uma ferramenta de profiling integrada ao Go que permite analisar o desempenho do seu programa em tempo de execução. Ele coleta dados sobre:

- **CPU**: Onde o programa gasta mais tempo processando
- **Memória**: Onde e como a memória está sendo alocada
- **Goroutines**: Quantas goroutines existem e onde estão bloqueadas
- **Blocking**: Operações que estão bloqueando a execução
- **Mutex**: Contenção em mutexes

### Por Que Usar pprof?

**Problemas que resolve:**
- ✅ **Identificar gargalos**: Descobrir onde o programa está lento
- ✅ **Otimizar memória**: Encontrar vazamentos e alocações desnecessárias
- ✅ **Analisar concorrência**: Entender comportamento de goroutines
- ✅ **Tomar decisões informadas**: Dados reais em vez de suposições
- ✅ **Medir melhorias**: Comparar performance antes e depois de otimizações

**Quando usar:**
- Quando o programa está mais lento do que esperado
- Quando há suspeita de vazamento de memória
- Quando há problemas de concorrência (muitas goroutines, deadlocks)
- Antes e depois de otimizações para medir impacto
- Em produção (com cuidado) para análise de problemas reais

### Como Funciona?

O `pprof` funciona coletando amostras do programa em execução. Existem duas formas principais de usar:

1. **Interface Web**: Importar `net/http/pprof` e acessar via HTTP
2. **Programático**: Usar `runtime/pprof` para salvar perfis em arquivos

### Método 1: Interface Web (Recomendado para Desenvolvimento)

A forma mais fácil de usar pprof é através de uma interface web. Basta importar o pacote `net/http/pprof`:

```go
package main

import (
    "fmt"
    "log"
    "net/http"
    _ "net/http/pprof" // Importa pprof automaticamente
    "time"
)

func main() {
    // Inicia servidor HTTP para pprof
    go func() {
        log.Println(http.ListenAndServe("localhost:6060", nil))
    }()

    // Seu código aqui
    for {
        doWork()
        time.Sleep(1 * time.Second)
    }
}

func doWork() {
    // Simula trabalho
    time.Sleep(100 * time.Millisecond)
}
```

**Como usar:**

1. Execute o programa
2. Acesse no navegador: `http://localhost:6060/debug/pprof/`
3. Você verá uma lista de perfis disponíveis:
   - `/debug/pprof/profile` - CPU profile
   - `/debug/pprof/heap` - Memory profile
   - `/debug/pprof/goroutine` - Goroutines
   - `/debug/pprof/block` - Blocking operations
   - `/debug/pprof/mutex` - Mutex contention

### Método 2: Programático (Para Produção)

Para produção, é melhor salvar perfis em arquivos:

```go
package main

import (
    "os"
    "runtime/pprof"
    "time"
)

func main() {
    // Criar arquivo para CPU profile
    f, err := os.Create("cpu.prof")
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()

    // Iniciar CPU profiling
    if err := pprof.StartCPUProfile(f); err != nil {
        log.Fatal(err)
    }
    defer pprof.StopCPUProfile()

    // Seu código aqui
    doWork()

    // Criar arquivo para memory profile
    memFile, err := os.Create("mem.prof")
    if err != nil {
        log.Fatal(err)
    }
    defer memFile.Close()

    // Forçar garbage collection
    runtime.GC()

    // Escrever memory profile
    if err := pprof.WriteHeapProfile(memFile); err != nil {
        log.Fatal(err)
    }
}

func doWork() {
    // Simula trabalho
    time.Sleep(100 * time.Millisecond)
}
```

### Analisando Perfis com go tool pprof

Depois de coletar um perfil, você pode analisá-lo usando `go tool pprof`:

```bash
# Analisar CPU profile
go tool pprof cpu.prof

# Analisar memory profile
go tool pprof mem.prof

# Analisar diretamente do servidor HTTP
go tool pprof http://localhost:6060/debug/pprof/profile?seconds=30
go tool pprof http://localhost:6060/debug/pprof/heap
```

**Comandos úteis no pprof interativo:**

```
top          # Mostra as funções que mais consomem recursos
top10        # Top 10 funções
list NomeFuncao  # Mostra código fonte da função
web          # Abre visualização gráfica no navegador
pdf          # Gera PDF com visualização
svg          # Gera SVG com visualização
```

### Exemplo Prático: Analisando CPU Profile

Vamos criar um programa com um gargalo intencional:

```go
package main

import (
    "fmt"
    "log"
    "net/http"
    _ "net/http/pprof"
    "time"
)

func main() {
    // Inicia servidor pprof
    go func() {
        log.Println(http.ListenAndServe("localhost:6060", nil))
    }()

    fmt.Println("Programa rodando. Acesse http://localhost:6060/debug/pprof/")
    fmt.Println("Para coletar CPU profile: go tool pprof http://localhost:6060/debug/pprof/profile?seconds=10")

    // Loop principal
    for {
        processData()
        time.Sleep(100 * time.Millisecond)
    }
}

func processData() {
    // Função rápida
    quickOperation()

    // Função lenta (gargalo)
    slowOperation()
}

func quickOperation() {
    time.Sleep(10 * time.Millisecond)
}

func slowOperation() {
    // Simula operação lenta
    sum := 0
    for i := 0; i < 1000000; i++ {
        sum += i
    }
    time.Sleep(50 * time.Millisecond)
}
```

**Como analisar:**

```bash
# 1. Execute o programa
go run main.go

# 2. Em outro terminal, colete CPU profile por 10 segundos
go tool pprof http://localhost:6060/debug/pprof/profile?seconds=10

# 3. No prompt do pprof, digite:
(pprof) top
(pprof) list slowOperation
(pprof) web
```

### Exemplo Prático: Analisando Memory Profile

Vamos criar um programa com vazamento de memória:

```go
package main

import (
    "fmt"
    "log"
    "net/http"
    _ "net/http/pprof"
    "runtime"
    "time"
)

var data []byte

func main() {
    // Inicia servidor pprof
    go func() {
        log.Println(http.ListenAndServe("localhost:6060", nil))
    }()

    fmt.Println("Programa rodando. Acesse http://localhost:6060/debug/pprof/heap")

    // Loop que causa vazamento de memória
    for {
        leakMemory()
        time.Sleep(1 * time.Second)
        
        // Mostrar estatísticas de memória
        var m runtime.MemStats
        runtime.ReadMemStats(&m)
        fmt.Printf("Alloc: %d KB, TotalAlloc: %d KB, Sys: %d KB\n",
            m.Alloc/1024, m.TotalAlloc/1024, m.Sys/1024)
    }
}

func leakMemory() {
    // Aloca memória mas nunca libera
    data = make([]byte, 1024*1024) // 1 MB
    // data deveria ser liberado, mas não é
}
```

**Como analisar:**

```bash
# 1. Execute o programa
go run main.go

# 2. Em outro terminal, colete memory profile
go tool pprof http://localhost:6060/debug/pprof/heap

# 3. No prompt do pprof:
(pprof) top
(pprof) list leakMemory
(pprof) web
```

### Tipos de Perfis Disponíveis

#### 1. CPU Profile (`/debug/pprof/profile`)

Mostra onde o programa gasta tempo de CPU:

```bash
go tool pprof http://localhost:6060/debug/pprof/profile?seconds=30
```

**Útil para:**
- Encontrar funções que consomem mais CPU
- Identificar loops lentos
- Otimizar algoritmos

#### 2. Heap Profile (`/debug/pprof/heap`)

Mostra alocações de memória no heap:

```bash
go tool pprof http://localhost:6060/debug/pprof/heap
```

**Útil para:**
- Encontrar vazamentos de memória
- Identificar alocações desnecessárias
- Otimizar uso de memória

#### 3. Goroutine Profile (`/debug/pprof/goroutine`)

Mostra todas as goroutines e onde estão:

```bash
go tool pprof http://localhost:6060/debug/pprof/goroutine
```

**Útil para:**
- Encontrar goroutines vazadas
- Entender estrutura de concorrência
- Detectar deadlocks potenciais

#### 4. Block Profile (`/debug/pprof/block`)

Mostra operações que estão bloqueando:

```bash
# Primeiro, habilite block profiling
runtime.SetBlockProfileRate(1)

# Depois colete o perfil
go tool pprof http://localhost:6060/debug/pprof/block
```

**Útil para:**
- Encontrar onde goroutines estão bloqueadas
- Otimizar I/O
- Melhorar concorrência

#### 5. Mutex Profile (`/debug/pprof/mutex`)

Mostra contenção em mutexes:

```bash
# Primeiro, habilite mutex profiling
runtime.SetMutexProfileFraction(1)

# Depois colete o perfil
go tool pprof http://localhost:6060/debug/pprof/mutex
```

**Útil para:**
- Encontrar mutexes com muita contenção
- Otimizar sincronização
- Reduzir lock contention

### Boas Práticas com pprof

#### 1. Coletar Perfis em Produção (Com Cuidado)

```go
// Adicione endpoint protegido para coletar perfis em produção
func setupPprof(mux *http.ServeMux) {
    if !isProduction() {
        // Em desenvolvimento, sempre disponível
        _ = net/http/pprof
        return
    }

    // Em produção, proteger com autenticação
    mux.HandleFunc("/debug/pprof/", func(w http.ResponseWriter, r *http.Request) {
        if !isAuthorized(r) {
            http.Error(w, "Forbidden", http.StatusForbidden)
            return
        }
        pprof.Index(w, r)
    })
}
```

#### 2. Coletar Perfis por Período Específico

```go
// Coletar CPU profile apenas quando necessário
func collectCPUProfile(duration time.Duration) error {
    f, err := os.Create(fmt.Sprintf("cpu-%d.prof", time.Now().Unix()))
    if err != nil {
        return err
    }
    defer f.Close()

    pprof.StartCPUProfile(f)
    time.Sleep(duration)
    pprof.StopCPUProfile()
    return nil
}
```

#### 3. Comparar Perfis Antes e Depois

```bash
# Antes da otimização
go tool pprof -base before.prof after.prof
```

---

## 2. trace - Análise de Execução e Concorrência

### O Que É trace?

O **`trace`** é uma ferramenta que captura uma **execução completa** do programa, mostrando:

- **Goroutines**: Quando são criadas, quando executam, quando bloqueiam
- **System calls**: Chamadas ao sistema operacional
- **Garbage Collection**: Quando o GC roda e quanto tempo leva
- **Scheduling**: Como o scheduler do Go distribui goroutines
- **Network I/O**: Operações de rede
- **Synchronization**: Mutexes, channels, etc.

### Por Que Usar trace?

**Problemas que resolve:**
- ✅ **Análise de concorrência**: Ver como goroutines interagem
- ✅ **Problemas de scheduling**: Entender por que goroutines não executam
- ✅ **GC pauses**: Identificar pausas do garbage collector
- ✅ **Latência**: Encontrar onde há atrasos na execução
- ✅ **Visualização**: Interface visual da execução

**Quando usar:**
- Quando há problemas de concorrência difíceis de debugar
- Quando o programa está lento mas pprof não mostra o problema
- Quando há suspeita de problemas com GC
- Para entender comportamento de goroutines
- Para otimizar scheduling

### Como Funciona?

O trace funciona capturando eventos durante a execução do programa. Você precisa:

1. **Gerar o trace**: Usar `runtime/trace` para criar arquivo de trace
2. **Analisar o trace**: Usar `go tool trace` para visualizar

### Gerando um Trace

```go
package main

import (
    "os"
    "runtime/trace"
    "time"
)

func main() {
    // Criar arquivo de trace
    f, err := os.Create("trace.out")
    if err != nil {
        panic(err)
    }
    defer f.Close()

    // Iniciar trace
    if err := trace.Start(f); err != nil {
        panic(err)
    }
    defer trace.Stop()

    // Seu código aqui
    doWork()
}

func doWork() {
    go func() {
        // Goroutine 1
        time.Sleep(100 * time.Millisecond)
    }()

    go func() {
        // Goroutine 2
        time.Sleep(200 * time.Millisecond)
    }()

    time.Sleep(300 * time.Millisecond)
}
```

### Analisando um Trace

```bash
# Gerar o trace
go run main.go

# Abrir o trace no navegador
go tool trace trace.out
```

Isso abrirá uma interface web no navegador com várias visualizações:

- **View trace**: Timeline visual de toda a execução
- **Goroutine analysis**: Análise de cada goroutine
- **Network blocking profile**: Bloqueios de rede
- **Synchronization blocking profile**: Bloqueios de sincronização
- **Syscall blocking profile**: Bloqueios de system calls
- **Scheduler latency profile**: Latência do scheduler

### Exemplo Prático: Analisando Concorrência

Vamos criar um programa com problema de concorrência:

```go
package main

import (
    "fmt"
    "os"
    "runtime/trace"
    "sync"
    "time"
)

func main() {
    // Criar arquivo de trace
    f, err := os.Create("trace.out")
    if err != nil {
        panic(err)
    }
    defer f.Close()

    // Iniciar trace
    if err := trace.Start(f); err != nil {
        panic(err)
    }
    defer trace.Stop()

    fmt.Println("Executando programa...")
    fmt.Println("Após terminar, execute: go tool trace trace.out")

    // Problema: muitas goroutines competindo por um mutex
    var mu sync.Mutex
    var wg sync.WaitGroup

    for i := 0; i < 100; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            
            mu.Lock()
            // Simula trabalho
            time.Sleep(10 * time.Millisecond)
            mu.Unlock()
        }(i)
    }

    wg.Wait()
    fmt.Println("Programa terminado!")
}
```

**Como analisar:**

```bash
# 1. Execute o programa
go run main.go

# 2. Abra o trace
go tool trace trace.out

# 3. No navegador, clique em "View trace"
# 4. Você verá visualmente como as goroutines competem pelo mutex
```

### Exemplo Prático: Analisando GC

Vamos criar um programa que causa muitas pausas do GC:

```go
package main

import (
    "os"
    "runtime/trace"
    "time"
)

func main() {
    f, err := os.Create("trace.out")
    if err != nil {
        panic(err)
    }
    defer f.Close()

    trace.Start(f)
    defer trace.Stop()

    // Aloca muita memória rapidamente
    for i := 0; i < 1000; i++ {
        data := make([]byte, 1024*1024) // 1 MB
        _ = data
        time.Sleep(1 * time.Millisecond)
    }
}
```

**Como analisar:**

```bash
# Execute e abra o trace
go run main.go
go tool trace trace.out

# No navegador, procure por barras vermelhas (GC events)
# Veja quanto tempo o GC está pausando a execução
```

### Interface Web do Trace

Quando você executa `go tool trace trace.out`, ele abre uma página web com:

1. **View trace**: Timeline interativa
   - Clique e arraste para zoom
   - Veja goroutines, GC, syscalls em cores diferentes
   - Clique em eventos para ver detalhes

2. **Goroutine analysis**: Lista de todas as goroutines
   - Veja onde cada goroutine passa tempo
   - Identifique goroutines bloqueadas

3. **Network blocking**: Bloqueios de rede
4. **Synchronization blocking**: Bloqueios de sincronização
5. **Syscall blocking**: Bloqueios de system calls

### Boas Práticas com trace

#### 1. Coletar Traces em Períodos Específicos

```go
func collectTrace(duration time.Duration) error {
    f, err := os.Create(fmt.Sprintf("trace-%d.out", time.Now().Unix()))
    if err != nil {
        return err
    }
    defer f.Close()

    trace.Start(f)
    time.Sleep(duration)
    trace.Stop()
    return nil
}
```

#### 2. Usar trace em Produção (Com Cuidado)

```go
// Endpoint para coletar trace em produção
func handleTrace(w http.ResponseWriter, r *http.Request) {
    if !isAuthorized(r) {
        http.Error(w, "Forbidden", http.StatusForbidden)
        return
    }

    duration, _ := time.ParseDuration(r.URL.Query().Get("seconds"))
    if duration == 0 {
        duration = 5 * time.Second
    }

    f, err := os.Create(fmt.Sprintf("trace-%d.out", time.Now().Unix()))
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer f.Close()

    trace.Start(f)
    time.Sleep(duration)
    trace.Stop()

    // Enviar arquivo
    http.ServeFile(w, r, f.Name())
}
```

---

## 3. Race Detector - Detectando Data Races

### O Que É Race Detector?

O **Race Detector** é uma ferramenta de runtime do Go que detecta **data races** em programas concorrentes. Um data race ocorre quando:

- Duas ou mais goroutines acessam a mesma variável
- Pelo menos uma das acessos é uma escrita
- Os acessos não são sincronizados

### Por Que Usar Race Detector?

**Problemas que resolve:**
- ✅ **Detectar data races**: Encontrar bugs de concorrência
- ✅ **Prevenir bugs sutis**: Races podem causar comportamento imprevisível
- ✅ **Garantir correção**: Verificar que código concorrente está correto
- ✅ **Debugging**: Encontrar a causa de bugs misteriosos

**Quando usar:**
- Sempre durante desenvolvimento de código concorrente
- Em testes de código que usa goroutines
- Quando há suspeita de data race
- Antes de fazer deploy de código concorrente
- Em CI/CD para detectar races automaticamente

### Como Funciona?

O Race Detector funciona rastreando acessos à memória durante a execução. Ele detecta quando há acessos não sincronizados à mesma memória.

**Importante**: O Race Detector tem um overhead significativo (2-10x mais lento), então use apenas durante desenvolvimento e testes, não em produção.

### Usando o Race Detector

Para usar o Race Detector, basta compilar e executar com a flag `-race`:

```bash
# Compilar com race detector
go build -race

# Executar com race detector
go run -race main.go

# Testar com race detector
go test -race ./...
```

### Exemplo 1: Data Race Simples

Vamos criar um programa com data race intencional:

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

var counter int

func main() {
    var wg sync.WaitGroup

    // Múltiplas goroutines escrevendo na mesma variável
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            for j := 0; j < 1000; j++ {
                counter++ // DATA RACE!
            }
        }()
    }

    wg.Wait()
    fmt.Println("Counter:", counter)
}
```

**Executando com race detector:**

```bash
go run -race main.go
```

**Saída esperada:**

```
==================
WARNING: DATA RACE
Read at 0x0000012345678900 by goroutine 7:
  main.main.func1()
      /path/to/main.go:20: +0x123

Previous write at 0x0000012345678900 by goroutine 6:
  main.main.func1()
      /path/to/main.go:20: +0x456

Goroutine 7 (running) created at:
  main.main()
      /path/to/main.go:15: +0x789

Goroutine 6 (finished) created at:
  main.main()
      /path/to/main.go:15: +0xabc
==================
Counter: 8234
Found 1 data race(s)
exit status 66
```

### Exemplo 2: Corrigindo o Data Race

Vamos corrigir o exemplo anterior usando mutex:

```go
package main

import (
    "fmt"
    "sync"
)

var (
    counter int
    mu      sync.Mutex
)

func main() {
    var wg sync.WaitGroup

    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            for j := 0; j < 1000; j++ {
                mu.Lock()
                counter++ // Agora protegido!
                mu.Unlock()
            }
        }()
    }

    wg.Wait()
    fmt.Println("Counter:", counter) // Sempre 10000
}
```

**Executando com race detector:**

```bash
go run -race main.go
# Sem warnings! Sem data races!
```

### Exemplo 3: Data Race em Slice

Data races também ocorrem em estruturas de dados:

```go
package main

import (
    "fmt"
    "sync"
)

var data []int

func main() {
    var wg sync.WaitGroup

    // Goroutine 1: adiciona elementos
    wg.Add(1)
    go func() {
        defer wg.Done()
        for i := 0; i < 1000; i++ {
            data = append(data, i) // DATA RACE!
        }
    }()

    // Goroutine 2: lê elementos
    wg.Add(1)
    go func() {
        defer wg.Done()
        for i := 0; i < 1000; i++ {
            if len(data) > i {
                _ = data[i] // DATA RACE!
            }
        }
    }()

    wg.Wait()
    fmt.Println("Done")
}
```

**Correção:**

```go
package main

import (
    "fmt"
    "sync"
)

var (
    data []int
    mu   sync.RWMutex
)

func main() {
    var wg sync.WaitGroup

    // Goroutine 1: adiciona elementos
    wg.Add(1)
    go func() {
        defer wg.Done()
        for i := 0; i < 1000; i++ {
            mu.Lock()
            data = append(data, i)
            mu.Unlock()
        }
    }()

    // Goroutine 2: lê elementos
    wg.Add(1)
    go func() {
        defer wg.Done()
        for i := 0; i < 1000; i++ {
            mu.RLock()
            if len(data) > i {
                _ = data[i]
            }
            mu.RUnlock()
        }
    }()

    wg.Wait()
    fmt.Println("Done")
}
```

### Exemplo 4: Data Race em Map

Maps em Go não são thread-safe:

```go
package main

import (
    "sync"
)

var m = make(map[string]int)

func main() {
    var wg sync.WaitGroup

    // DATA RACE: múltiplas goroutines escrevendo no map
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            m[fmt.Sprintf("key%d", id)] = id // DATA RACE!
        }(i)
    }

    wg.Wait()
}
```

**Correção:**

```go
package main

import (
    "fmt"
    "sync"
)

var (
    m  = make(map[string]int)
    mu sync.Mutex
)

func main() {
    var wg sync.WaitGroup

    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            mu.Lock()
            m[fmt.Sprintf("key%d", id)] = id
            mu.Unlock()
        }(i)
    }

    wg.Wait()
}
```

### Boas Práticas com Race Detector

#### 1. Sempre Usar em Testes

```bash
# Adicione ao seu Makefile
test:
	go test -race ./...

# Ou no CI/CD
- name: Run tests with race detector
  run: go test -race ./...
```

#### 2. Usar em Desenvolvimento

```bash
# Desenvolvimento local
go run -race main.go

# Build para testes
go build -race -o myapp
```

#### 3. NÃO Usar em Produção

**Importante**: Nunca use `-race` em produção porque:
- Overhead de 2-10x mais lento
- Consumo de memória muito maior
- Apenas para detectar bugs, não para executar

#### 4. Testar Código Concorrente Extensivamente

```go
func TestConcurrentAccess(t *testing.T) {
    var counter int
    var mu sync.Mutex
    var wg sync.WaitGroup

    // Múltiplas goroutines acessando
    for i := 0; i < 100; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            mu.Lock()
            counter++
            mu.Unlock()
        }()
    }

    wg.Wait()
    
    if counter != 100 {
        t.Errorf("Expected 100, got %d", counter)
    }
}
```

Execute com: `go test -race -v`

---

## 4. Combinando as Ferramentas

### Workflow Completo de Análise

1. **Race Detector**: Primeiro, garanta que não há data races
   ```bash
   go test -race ./...
   ```

2. **pprof**: Depois, analise performance
   ```bash
   go tool pprof http://localhost:6060/debug/pprof/profile
   ```

3. **trace**: Se necessário, analise execução detalhada
   ```bash
   go tool trace trace.out
   ```

### Exemplo Completo

```go
package main

import (
    "log"
    "net/http"
    _ "net/http/pprof"
    "os"
    "runtime/trace"
    "sync"
    "time"
)

var (
    data  = make(map[string]int)
    mu    sync.Mutex
    wg    sync.WaitGroup
)

func main() {
    // Iniciar pprof
    go func() {
        log.Println(http.ListenAndServe("localhost:6060", nil))
    }()

    // Iniciar trace (opcional)
    if os.Getenv("TRACE") == "1" {
        f, _ := os.Create("trace.out")
        trace.Start(f)
        defer trace.Stop()
    }

    // Seu código aqui
    processData()

    wg.Wait()
    log.Println("Done")
}

func processData() {
    for i := 0; i < 100; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            
            mu.Lock()
            data[fmt.Sprintf("key%d", id)] = id
            mu.Unlock()
            
            time.Sleep(10 * time.Millisecond)
        }(i)
    }
}
```

**Como usar:**

```bash
# 1. Testar com race detector
go test -race ./...

# 2. Executar e analisar com pprof
go run main.go
go tool pprof http://localhost:6060/debug/pprof/profile

# 3. Executar com trace
TRACE=1 go run main.go
go tool trace trace.out
```

---

## Resumo dos Conceitos

| Ferramenta | O Que Faz | Quando Usar | Overhead |
|------------|-----------|-------------|----------|
| **pprof** | Analisa performance (CPU, memória, goroutines) | Quando há problemas de performance | Baixo (pode coletar em produção) |
| **trace** | Captura execução completa (goroutines, GC, scheduling) | Quando há problemas de concorrência ou GC | Médio (use com cuidado) |
| **Race Detector** | Detecta data races | Sempre em desenvolvimento/testes | Alto (2-10x mais lento) |

---

## Conclusão

Dominar `pprof`, `trace` e Race Detector é essencial para:

1. **Criar código performático**: Identificar e corrigir gargalos
2. **Debugging eficiente**: Encontrar problemas rapidamente
3. **Garantir correção**: Detectar bugs de concorrência
4. **Tomar decisões informadas**: Dados reais em vez de suposições
5. **Otimizar sistematicamente**: Medir antes e depois de mudanças

Essas ferramentas são fundamentais para qualquer programador Go profissional. Integre-as no seu workflow e veja a qualidade do seu código melhorar!

Na próxima parte desta aula, vamos simplificar esses conceitos com analogias do dia a dia para fixar melhor o aprendizado!


