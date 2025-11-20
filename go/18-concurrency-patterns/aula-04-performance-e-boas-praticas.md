# Módulo 18: Concurrency Patterns em Go

## Aula 4: Performance, Boas Práticas e Armadilhas Comuns

Olá! Agora que você já entende os padrões de concorrência, é hora de aprender como usá-los de forma eficiente, segura e profissional. Nesta aula, vamos focar em performance, boas práticas, o que fazer e o que NÃO fazer, e como evitar as armadilhas mais comuns.

---

## 🚀 1. Performance: Otimizando Padrões de Concorrência

### 1.1 Buffer Size em Channels

**Problema**: Channels unbuffered podem causar bloqueios desnecessários.

```go
// ❌ RUIM: Channel unbuffered pode bloquear
output := make(chan string)
go producer(output) // Pode bloquear se consumidor estiver lento

// ✅ BOM: Buffer apropriado reduz bloqueios
output := make(chan string, 100) // Buffer para 100 itens
go producer(output) // Não bloqueia até buffer encher
```

**Quando usar buffer:**
- ✅ Produtor é mais rápido que consumidor
- ✅ Você quer processar em lotes
- ✅ Evitar bloqueios em pipelines

**Quando NÃO usar buffer:**
- ✅ Sincronização precisa (quando precisa garantir que consumidor recebeu)
- ✅ Backpressure é desejado (quando quer que produtor espere)

### 1.2 Número de Workers em Worker Pool

**Problema**: Muitos workers podem causar overhead, poucos podem subutilizar CPU.

```go
// ❌ RUIM: Workers demais causam overhead
pool := NewWorkerPool(10000) // Muito overhead de contexto switching

// ❌ RUIM: Workers de menos subutilizam CPU
pool := NewWorkerPool(1) // Não aproveita múltiplos cores

// ✅ BOM: Número baseado em CPU disponível
numCPU := runtime.NumCPU()
pool := NewWorkerPool(numCPU * 2) // 2x número de CPUs é um bom ponto de partida
```

**Regra de ouro:**
- Para **CPU-bound tasks**: Use `runtime.NumCPU()` ou `NumCPU() * 2`
- Para **I/O-bound tasks**: Use mais workers (10-100x número de CPUs)
- **Meça e ajuste**: Não adivinhe, use profiling!

### 1.3 Pipeline: Evitar Estágios Desnecessários

**Problema**: Muitos estágios adicionam latência e overhead.

```go
// ❌ RUIM: Estágios demais para tarefa simples
numbers := generate()
formatted1 := format1(numbers)
formatted2 := format2(formatted1)
formatted3 := format3(formatted2)
final := format4(formatted3)

// ✅ BOM: Combinar estágios quando faz sentido
numbers := generate()
final := format(numbers) // Uma função faz tudo que precisa
```

**Quando usar múltiplos estágios:**
- ✅ Cada estágio tem responsabilidade clara e diferente
- ✅ Estágios podem rodar em paralelo com dados diferentes
- ✅ Você quer reutilizar estágios em diferentes pipelines

### 1.4 Fan-In: Evitar Goroutines Desnecessárias

**Problema**: Criar goroutine para cada channel pode ser excessivo.

```go
// ❌ RUIM: Goroutine para cada channel (pode ser excessivo para muitos channels)
func fanIn(channels ...<-chan string) <-chan string {
    out := make(chan string)
    for _, ch := range channels { // 1000 channels = 1000 goroutines!
        go func(c <-chan string) {
            for msg := range c {
                out <- msg
            }
        }(ch)
    }
    return out
}

// ✅ BOM: Usar select para poucos channels, ou limitar goroutines
func fanInSelect(ch1, ch2, ch3 <-chan string) <-chan string {
    out := make(chan string)
    go func() {
        defer close(out)
        for {
            select {
            case msg := <-ch1:
                out <- msg
            case msg := <-ch2:
                out <- msg
            case msg := <-ch3:
                out <- msg
            }
        }
    }()
    return out
}
```

---

## ✅ 2. Boas Práticas Essenciais

### 2.1 Sempre Feche Channels

**CRÍTICO**: Channels não fechados causam goroutines vazando e deadlocks.

```go
// ✅ SEMPRE faça isso
go func() {
    defer close(output) // Sempre feche!
    for item := range input {
        output <- process(item)
    }
}()
```

**Regra**: Quem cria o channel é responsável por fechá-lo.

### 2.2 Use Context para Cancelamento

**Sempre** passe context para operações que podem ser canceladas.

```go
// ✅ BOM: Context para cancelamento
func worker(ctx context.Context, jobs <-chan Job) {
    for {
        select {
        case <-ctx.Done():
            return // Cancelar quando contexto for cancelado
        case job := <-jobs:
            process(job)
        }
    }
}
```

### 2.3 Use WaitGroup Corretamente

**Sempre** use `defer wg.Done()` para garantir que é chamado mesmo em caso de panic.

```go
// ✅ BOM: defer garante Done() sempre é chamado
func worker(wg *sync.WaitGroup, jobs <-chan Job) {
    defer wg.Done() // Sempre chama, mesmo com panic
    for job := range jobs {
        process(job)
    }
}
```

### 2.4 Proteja Dados Compartilhados com Mutex

**Se** você realmente precisa compartilhar memória (não apenas channels), use mutex.

```go
// ✅ BOM: Mutex para proteger dados compartilhados
type Counter struct {
    mu    sync.RWMutex
    value int
}

func (c *Counter) Increment() {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.value++
}

func (c *Counter) Value() int {
    c.mu.RLock()
    defer c.mu.RUnlock()
    return c.value
}
```

**Mas lembre-se**: Prefira channels quando possível!

### 2.5 Use RWMutex para Leitura Concorrente

**Quando** você tem muitas leituras e poucas escritas, use `RWMutex`.

```go
// ✅ BOM: RWMutex permite múltiplas leituras simultâneas
type Cache struct {
    mu    sync.RWMutex
    data  map[string]string
}

func (c *Cache) Get(key string) string {
    c.mu.RLock() // Lock para leitura (não bloqueia outras leituras)
    defer c.mu.RUnlock()
    return c.data[key]
}

func (c *Cache) Set(key, value string) {
    c.mu.Lock() // Lock exclusivo para escrita
    defer c.mu.Unlock()
    c.data[key] = value
}
```

---

## ❌ 3. O que NÃO Fazer: Armadilhas Comuns

### 3.1 NÃO Compartilhe Memória sem Proteção

```go
// ❌ ERRADO: Race condition!
var counter int

func increment() {
    counter++ // Múltiplas goroutines = race condition!
}

// ✅ CORRETO: Use channels ou mutex
var counter int
var mu sync.Mutex

func increment() {
    mu.Lock()
    defer mu.Unlock()
    counter++
}
```

### 3.2 NÃO Esqueça de Fechar Channels

```go
// ❌ ERRADO: Channel nunca é fechado, goroutine vaza
func producer() <-chan string {
    ch := make(chan string)
    go func() {
        for i := 0; i < 10; i++ {
            ch <- fmt.Sprintf("item %d", i)
        }
        // Esqueceu de fechar!
    }()
    return ch
}

// ✅ CORRETO: Sempre feche
func producer() <-chan string {
    ch := make(chan string)
    go func() {
        defer close(ch) // Sempre fecha!
        for i := 0; i < 10; i++ {
            ch <- fmt.Sprintf("item %d", i)
        }
    }()
    return ch
}
```

### 3.3 NÃO Crie Goroutines Demais

```go
// ❌ ERRADO: Pode criar milhões de goroutines
for i := 0; i < 1000000; i++ {
    go process(i) // Muito overhead!
}

// ✅ CORRETO: Use worker pool
pool := NewWorkerPool(100) // Limita número de goroutines
for i := 0; i < 1000000; i++ {
    pool.Submit(Task{ID: i})
}
```

### 3.4 NÃO Use Channels para Dados que Não Precisam de Sincronização

```go
// ❌ ERRADO: Channel desnecessário
func getConfig() <-chan Config {
    ch := make(chan Config, 1)
    ch <- loadConfig()
    return ch
}

// ✅ CORRETO: Retornar valor diretamente
func getConfig() Config {
    return loadConfig()
}
```

**Use channels quando:**
- ✅ Precisa de sincronização
- ✅ Dados chegam assincronamente
- ✅ Precisa comunicar entre goroutines

### 3.5 NÃO Bloqueie em Channels sem Timeout

```go
// ❌ ERRADO: Pode bloquear para sempre
msg := <-ch // E se ninguém enviar?

// ✅ CORRETO: Use select com timeout
select {
case msg := <-ch:
    process(msg)
case <-time.After(5 * time.Second):
    fmt.Println("Timeout!")
}
```

### 3.6 NÃO Feche Channel Duas Vezes

```go
// ❌ ERRADO: Fechar channel duas vezes causa panic
close(ch)
close(ch) // PANIC!

// ✅ CORRETO: Use sync.Once ou verifique se já foi fechado
var once sync.Once
once.Do(func() {
    close(ch)
})
```

### 3.7 NÃO Envie para Channel Fechado

```go
// ❌ ERRADO: Enviar para channel fechado causa panic
close(ch)
ch <- "message" // PANIC!

// ✅ CORRETO: Verifique se channel está fechado antes de enviar
select {
case ch <- "message":
    // Enviado com sucesso
case <-ctx.Done():
    // Context cancelado, não enviar
default:
    // Channel cheio ou fechado, tratar apropriadamente
}
```

---

## 🎯 4. Padrões Específicos: Boas Práticas

### 4.1 Fan-In: Detectar Quando Todos os Channels Fecharam

```go
// ✅ BOM: Detectar quando todos os channels fecharam
func fanIn(channels ...<-chan string) <-chan string {
    out := make(chan string)
    var wg sync.WaitGroup
    
    for _, ch := range channels {
        wg.Add(1)
        go func(c <-chan string) {
            defer wg.Done()
            for msg := range c {
                out <- msg
            }
        }(ch)
    }
    
    go func() {
        wg.Wait()
        close(out) // Fechar quando todos terminarem
    }()
    
    return out
}
```

### 4.2 Fan-Out: Controlar Número de Workers

```go
// ✅ BOM: Limitar número de workers
func fanOut(jobs <-chan Job, numWorkers int) {
    var wg sync.WaitGroup
    
    for i := 0; i < numWorkers; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            for job := range jobs {
                process(job)
            }
        }(i)
    }
    
    wg.Wait()
}
```

### 4.3 Pipeline: Verificar Cancelamento em Cada Estágio

```go
// ✅ BOM: Cada estágio verifica cancelamento
func stage(ctx context.Context, in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        defer close(out)
        for {
            select {
            case <-ctx.Done():
                return
            case val, ok := <-in:
                if !ok {
                    return
                }
                select {
                case out <- process(val):
                case <-ctx.Done():
                    return
                }
            }
        }
    }()
    return out
}
```

### 4.4 Worker Pool: Graceful Shutdown

```go
// ✅ BOM: Shutdown graceful com context
type WorkerPool struct {
    workers  int
    jobs     chan Job
    wg       sync.WaitGroup
    shutdown chan struct{}
}

func (wp *WorkerPool) Start(ctx context.Context) {
    for i := 0; i < wp.workers; i++ {
        wp.wg.Add(1)
        go wp.worker(ctx, i)
    }
}

func (wp *WorkerPool) worker(ctx context.Context, id int) {
    defer wp.wg.Done()
    for {
        select {
        case <-ctx.Done():
            return
        case <-wp.shutdown:
            return
        case job, ok := <-wp.jobs:
            if !ok {
                return
            }
            process(job)
        }
    }
}

func (wp *WorkerPool) Stop() {
    close(wp.shutdown)
    close(wp.jobs)
    wp.wg.Wait()
}
```

### 4.5 Pub-Sub: Evitar Bloqueio do Publisher

```go
// ✅ BOM: Buffer e select para evitar bloqueio
func (ps *PubSub) Publish(topic string, content string) {
    ps.mutex.RLock()
    defer ps.mutex.RUnlock()
    
    msg := Message{Topic: topic, Content: content}
    for _, ch := range ps.subscribers[topic] {
        select {
        case ch <- msg: // Não bloqueia se channel cheio
        default:
            // Subscriber lento, pular (ou logar)
        }
    }
}
```

---

## 🔍 5. Debugging e Profiling

### 5.1 Detectar Race Conditions

Use a flag `-race` ao compilar e executar:

```bash
go run -race main.go
go test -race ./...
```

**Importante**: `-race` adiciona overhead, use apenas em desenvolvimento e testes.

### 5.2 Profiling de Goroutines

```go
import (
    "net/http"
    _ "net/http/pprof"
    "runtime"
)

func main() {
    // Iniciar servidor de profiling
    go func() {
        http.ListenAndServe("localhost:6060", nil)
    }()
    
    // Seu código aqui
}
```

Acesse `http://localhost:6060/debug/pprof/goroutine?debug=1` para ver goroutines.

### 5.3 Contar Goroutines

```go
import "runtime"

func logGoroutines() {
    num := runtime.NumGoroutine()
    fmt.Printf("Goroutines ativas: %d\n", num)
}
```

### 5.4 Detectar Vazamentos de Goroutines

```go
func TestGoroutineLeak(t *testing.T) {
    before := runtime.NumGoroutine()
    
    // Seu código que pode vazar goroutines
    
    runtime.GC() // Forçar garbage collection
    time.Sleep(100 * time.Millisecond)
    
    after := runtime.NumGoroutine()
    if after > before {
        t.Errorf("Possível vazamento: %d -> %d goroutines", before, after)
    }
}
```

---

## 📊 6. Métricas e Monitoramento

### 6.1 Medir Throughput

```go
func measureThroughput(process func()) {
    start := time.Now()
    count := 0
    
    for i := 0; i < 1000; i++ {
        process()
        count++
    }
    
    duration := time.Since(start)
    throughput := float64(count) / duration.Seconds()
    fmt.Printf("Throughput: %.2f operações/segundo\n", throughput)
}
```

### 6.2 Medir Latência

```go
func measureLatency(process func()) time.Duration {
    start := time.Now()
    process()
    return time.Since(start)
}
```

### 6.3 Monitorar Tamanho de Channels

```go
type MonitoredChannel struct {
    ch     chan string
    size   int64
    mu     sync.RWMutex
}

func (mc *MonitoredChannel) Send(msg string) {
    mc.ch <- msg
    mc.mu.Lock()
    mc.size++
    mc.mu.Unlock()
}

func (mc *MonitoredChannel) Size() int64 {
    mc.mu.RLock()
    defer mc.mu.RUnlock()
    return mc.size
}
```

---

## 🎓 7. Resumo: Checklist de Boas Práticas

### Sempre Faça:
- [ ] Feche channels com `defer close()`
- [ ] Use context para cancelamento
- [ ] Use `defer wg.Done()` em WaitGroups
- [ ] Proteja dados compartilhados com mutex (se necessário)
- [ ] Use buffer apropriado em channels
- [ ] Limite número de workers em pools
- [ ] Teste com `-race` flag
- [ ] Implemente graceful shutdown

### Nunca Faça:
- [ ] Compartilhar memória sem proteção
- [ ] Esquecer de fechar channels
- [ ] Criar goroutines ilimitadas
- [ ] Usar channels quando não precisa
- [ ] Bloquear sem timeout
- [ ] Fechar channel duas vezes
- [ ] Enviar para channel fechado
- [ ] Ignorar race conditions

---

## 🚀 8. Dicas Avançadas

### 8.1 Usar sync.Pool para Reduzir Allocations

```go
var bufferPool = sync.Pool{
    New: func() interface{} {
        return make([]byte, 0, 1024)
    },
}

func process() {
    buf := bufferPool.Get().([]byte)
    defer bufferPool.Put(buf[:0]) // Reset e retorna ao pool
    
    // Usar buf
}
```

### 8.2 Usar atomic para Contadores Simples

```go
import "sync/atomic"

var counter int64

func increment() {
    atomic.AddInt64(&counter, 1) // Mais rápido que mutex para contadores
}

func value() int64 {
    return atomic.LoadInt64(&counter)
}
```

### 8.3 Usar errgroup para Coordenar Goroutines com Erros

```go
import "golang.org/x/sync/errgroup"

func processAll(items []Item) error {
    var g errgroup.Group
    
    for _, item := range items {
        item := item // Importante: capturar variável
        g.Go(func() error {
            return process(item)
        })
    }
    
    return g.Wait() // Retorna primeiro erro, se houver
}
```

---

## 🎯 9. Quando Usar Cada Padrão: Guia Rápido

| Padrão | Quando Usar | Performance | Complexidade |
|--------|-------------|------------|--------------|
| **Fan-In** | Agregar resultados | Alta | Baixa |
| **Fan-Out** | Paralelizar trabalho | Muito Alta | Média |
| **Pipeline** | Transformações sequenciais | Alta | Média |
| **Worker Pool** | Controlar recursos | Alta | Média-Alta |
| **Pub-Sub** | Eventos desacoplados | Média | Alta |

---

## 📚 10. Recursos para Aprofundar

1. **Effective Go**: https://go.dev/doc/effective_go
2. **Go Concurrency Patterns**: https://go.dev/blog/pipelines
3. **Advanced Go Concurrency Patterns**: https://go.dev/blog/advanced-go-concurrency-patterns
4. **Race Detector**: https://go.dev/blog/race-detector
5. **Profiling Go Programs**: https://go.dev/blog/pprof

---

E assim terminamos nossa quarta aula sobre Concurrency Patterns! Você agora sabe:
- Como otimizar performance dos padrões
- Boas práticas essenciais
- O que NÃO fazer (armadilhas comuns)
- Como debugar e fazer profiling
- Dicas avançadas para código profissional

Lembre-se: **escrever código concorrente correto é difícil**. Sempre teste com `-race`, meça performance, e prefira simplicidade sobre otimização prematura.

Na próxima etapa, você fará os exercícios e responderá as questões de reflexão. Depois, analisarei seu desempenho e darei feedback construtivo!

Boa sorte! 🚀



