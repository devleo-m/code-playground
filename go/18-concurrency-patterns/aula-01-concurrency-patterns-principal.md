# Módulo 18: Concurrency Patterns em Go

## Aula 1: Concurrency Patterns - Padrões de Design para Programação Concorrente

Olá! Bem-vindo ao décimo oitavo módulo. Até agora você aprendeu sobre goroutines, channels, select, context e os fundamentos da concorrência em Go. Mas você já se perguntou: **"Como estruturar programas concorrentes de forma eficiente e escalável? Como organizar goroutines e channels para resolver problemas complexos? Como evitar race conditions e deadlocks ao construir aplicações concorrentes?"**

Os **Concurrency Patterns** (Padrões de Concorrência) são abordagens estabelecidas e testadas para estruturar programas concorrentes usando goroutines e channels. Eles fornecem soluções reutilizáveis para problemas comuns em programação concorrente, ajudando você a construir aplicações eficientes, escaláveis e seguras.

Nesta aula, vamos mergulhar profundamente nos principais padrões de concorrência em Go: **Fan-in** (mesclar múltiplas entradas), **Fan-out** (distribuir trabalho), **Pipeline** (encadear operações), **Worker Pools** (piscina de trabalhadores) e **Pub-Sub** (publicar-assinar). Cada padrão será explicado com exemplos práticos e casos de uso reais.

---

## 1. O que são Concurrency Patterns?

### Definição

**Concurrency Patterns** são soluções de design reutilizáveis para problemas comuns em programação concorrente. Eles fornecem uma estrutura organizada para coordenar múltiplas goroutines e gerenciar a comunicação através de channels.

### Por que Usar Padrões de Concorrência?

1. **Organização**: Estrutura clara e previsível para código concorrente
2. **Reutilização**: Soluções testadas que podem ser aplicadas em diferentes contextos
3. **Segurança**: Ajudam a evitar race conditions e deadlocks
4. **Escalabilidade**: Permitem construir sistemas que podem crescer eficientemente
5. **Manutenibilidade**: Código mais fácil de entender e manter

### Filosofia Go: "Share Memory by Communicating"

Go segue a filosofia de que você deve **comunicar compartilhando memória através de channels**, não compartilhar memória diretamente. Os padrões de concorrência em Go refletem essa filosofia, usando channels como mecanismo principal de comunicação e sincronização.

---

## 2. Fan-In Pattern: Mesclando Múltiplas Entradas

### O que é Fan-In?

**Fan-in** é um padrão que **mescla múltiplos canais de entrada em um único canal de saída**. É útil quando você tem várias goroutines produzindo dados e precisa coletar todos os resultados em um único lugar.

### Analogia

Imagine várias pessoas (goroutines) escrevendo cartas (dados) e colocando-as em caixas de correio diferentes (channels). O padrão Fan-in é como um carteiro que coleta todas as cartas de todas as caixas e as coloca em uma única sacola (channel de saída).

### Implementação Básica com Select

A forma mais comum de implementar Fan-in é usando o statement `select` para ler de múltiplos channels simultaneamente:

```go
package main

import (
    "fmt"
    "time"
)

// Fan-in: mescla múltiplos channels em um único channel
func fanIn(input1, input2 <-chan string) <-chan string {
    output := make(chan string)
    
    go func() {
        defer close(output)
        for {
            select {
            case msg := <-input1:
                output <- msg
            case msg := <-input2:
                output <- msg
            }
        }
    }()
    
    return output
}

// Função que produz dados em um channel
func producer(name string, delay time.Duration) <-chan string {
    ch := make(chan string)
    go func() {
        defer close(ch)
        for i := 1; i <= 5; i++ {
            ch <- fmt.Sprintf("%s: mensagem %d", name, i)
            time.Sleep(delay)
        }
    }()
    return ch
}

func main() {
    // Criar dois producers
    producer1 := producer("Producer 1", 200*time.Millisecond)
    producer2 := producer("Producer 2", 300*time.Millisecond)
    
    // Mesclar os dois channels em um
    merged := fanIn(producer1, producer2)
    
    // Ler todas as mensagens mescladas
    for msg := range merged {
        fmt.Println(msg)
    }
    
    fmt.Println("Todas as mensagens foram processadas")
}
```

**Saída possível:**
```
Producer 1: mensagem 1
Producer 2: mensagem 1
Producer 1: mensagem 2
Producer 1: mensagem 3
Producer 2: mensagem 2
Producer 1: mensagem 4
Producer 2: mensagem 3
Producer 1: mensagem 5
Producer 2: mensagem 4
Producer 2: mensagem 5
Todas as mensagens foram processadas
```

### Fan-In com N Channels Dinâmicos

Quando você tem um número variável de channels para mesclar, pode usar uma abordagem mais dinâmica:

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

// Fan-in dinâmico: mescla N channels em um único channel
func fanInDynamic(channels ...<-chan string) <-chan string {
    output := make(chan string)
    var wg sync.WaitGroup
    
    // Para cada channel de entrada, criar uma goroutine que lê e envia para output
    for _, ch := range channels {
        wg.Add(1)
        go func(c <-chan string) {
            defer wg.Done()
            for msg := range c {
                output <- msg
            }
        }(ch)
    }
    
    // Fechar output quando todos os channels de entrada forem fechados
    go func() {
        wg.Wait()
        close(output)
    }()
    
    return output
}

func producer(name string, count int, delay time.Duration) <-chan string {
    ch := make(chan string)
    go func() {
        defer close(ch)
        for i := 1; i <= count; i++ {
            ch <- fmt.Sprintf("%s: item %d", name, i)
            time.Sleep(delay)
        }
    }()
    return ch
}

func main() {
    // Criar múltiplos producers
    ch1 := producer("Worker 1", 3, 100*time.Millisecond)
    ch2 := producer("Worker 2", 3, 150*time.Millisecond)
    ch3 := producer("Worker 3", 3, 200*time.Millisecond)
    
    // Mesclar todos os channels
    merged := fanInDynamic(ch1, ch2, ch3)
    
    // Processar todas as mensagens
    for msg := range merged {
        fmt.Println("Recebido:", msg)
    }
    
    fmt.Println("Processamento concluído")
}
```

### Quando Usar Fan-In?

- ✅ **Agregação de Resultados**: Coletar resultados de múltiplas goroutines processando dados em paralelo
- ✅ **Múltiplas Fontes de Dados**: Mesclar dados de diferentes APIs, bancos de dados ou arquivos
- ✅ **Logging Centralizado**: Coletar logs de múltiplas goroutines em um único ponto
- ✅ **Monitoramento**: Agregar métricas de diferentes componentes do sistema

---

## 3. Fan-Out Pattern: Distribuindo Trabalho

### O que é Fan-Out?

**Fan-out** é um padrão que **distribui trabalho de um único canal de entrada para múltiplos workers (trabalhadores)**. Cada worker processa itens independentemente, permitindo paralelizar tarefas CPU-intensivas e aumentar o throughput.

### Analogia

Imagine uma linha de produção onde uma esteira (channel de entrada) transporta produtos. O padrão Fan-out é como ter vários trabalhadores (goroutines) pegando produtos da esteira e processando-os simultaneamente, cada um em sua própria estação de trabalho.

### Implementação Básica

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

// Worker processa itens do channel de entrada
func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
    defer wg.Done()
    for job := range jobs {
        // Simular processamento
        fmt.Printf("Worker %d: processando job %d\n", id, job)
        time.Sleep(500 * time.Millisecond) // Simular trabalho
        
        // Enviar resultado
        results <- job * 2 // Exemplo: dobrar o valor
    }
}

func main() {
    // Channels para jobs e resultados
    jobs := make(chan int, 10)
    results := make(chan int, 10)
    
    // Número de workers
    numWorkers := 3
    var wg sync.WaitGroup
    
    // Iniciar workers (Fan-out)
    for w := 1; w <= numWorkers; w++ {
        wg.Add(1)
        go worker(w, jobs, results, &wg)
    }
    
    // Enviar jobs
    go func() {
        defer close(jobs)
        for j := 1; j <= 9; j++ {
            jobs <- j
        }
    }()
    
    // Fechar results quando todos os workers terminarem
    go func() {
        wg.Wait()
        close(results)
    }()
    
    // Coletar resultados
    for result := range results {
        fmt.Printf("Resultado: %d\n", result)
    }
    
    fmt.Println("Todos os jobs foram processados")
}
```

**Saída possível:**
```
Worker 1: processando job 1
Worker 2: processando job 2
Worker 3: processando job 3
Worker 1: processando job 4
Resultado: 2
Worker 2: processando job 5
Resultado: 4
Worker 3: processando job 6
Resultado: 6
Worker 1: processando job 7
Resultado: 8
Worker 2: processando job 8
Resultado: 10
Worker 3: processando job 9
Resultado: 12
Resultado: 14
Resultado: 16
Resultado: 18
Todos os jobs foram processados
```

### Fan-Out com Rate Limiting

Você pode controlar quantos workers processam simultaneamente:

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

func worker(id int, jobs <-chan string, results chan<- string, wg *sync.WaitGroup) {
    defer wg.Done()
    for job := range jobs {
        fmt.Printf("Worker %d: processando %s\n", id, job)
        time.Sleep(1 * time.Second)
        results <- fmt.Sprintf("%s processado por worker %d", job, id)
    }
}

func main() {
    jobs := make(chan string, 100)
    results := make(chan string, 100)
    
    // Limitar número de workers simultâneos
    maxWorkers := 5
    var wg sync.WaitGroup
    
    // Iniciar workers limitados
    for w := 1; w <= maxWorkers; w++ {
        wg.Add(1)
        go worker(w, jobs, results, &wg)
    }
    
    // Enviar jobs
    go func() {
        defer close(jobs)
        tasks := []string{"task1", "task2", "task3", "task4", "task5", "task6", "task7", "task8"}
        for _, task := range tasks {
            jobs <- task
        }
    }()
    
    // Fechar results quando workers terminarem
    go func() {
        wg.Wait()
        close(results)
    }()
    
    // Coletar resultados
    for result := range results {
        fmt.Println("Resultado:", result)
    }
}
```

### Quando Usar Fan-Out?

- ✅ **Processamento Paralelo**: Distribuir tarefas CPU-intensivas entre múltiplos workers
- ✅ **Aumentar Throughput**: Processar mais itens por segundo usando paralelismo
- ✅ **Balanceamento de Carga**: Distribuir requisições entre múltiplos servidores/workers
- ✅ **Processamento de Dados**: Processar grandes volumes de dados em paralelo

---

## 4. Pipeline Pattern: Encadeando Estágios de Processamento

### O que é Pipeline?

**Pipeline** é um padrão que **encadeia múltiplos estágios de processamento**, onde a saída de um estágio se torna a entrada do próximo. Cada estágio roda em uma goroutine separada, permitindo processamento paralelo e separação de responsabilidades.

### Analogia

Imagine uma linha de montagem de carros: primeiro estágio monta o chassi, segundo pinta, terceiro instala o motor, quarto adiciona os bancos. Cada estágio recebe o trabalho do estágio anterior, faz sua parte e passa para o próximo.

### Implementação Básica

```go
package main

import (
    "fmt"
    "strings"
    "time"
)

// Estágio 1: Gerar números
func generateNumbers(count int) <-chan int {
    out := make(chan int)
    go func() {
        defer close(out)
        for i := 1; i <= count; i++ {
            out <- i
            time.Sleep(100 * time.Millisecond)
        }
    }()
    return out
}

// Estágio 2: Elevar ao quadrado
func square(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        defer close(out)
        for n := range in {
            squared := n * n
            fmt.Printf("Estágio 2: %d² = %d\n", n, squared)
            out <- squared
            time.Sleep(100 * time.Millisecond)
        }
    }()
    return out
}

// Estágio 3: Formatar como string
func format(in <-chan int) <-chan string {
    out := make(chan string)
    go func() {
        defer close(out)
        for n := range in {
            formatted := fmt.Sprintf("Resultado: %d", n)
            fmt.Printf("Estágio 3: %s\n", formatted)
            out <- formatted
        }
    }()
    return out
}

func main() {
    // Criar pipeline: generate -> square -> format
    numbers := generateNumbers(5)
    squared := square(numbers)
    formatted := format(squared)
    
    // Consumir resultado final
    for result := range formatted {
        fmt.Printf("Saída final: %s\n", result)
        fmt.Println(strings.Repeat("-", 30))
    }
}
```

**Saída:**
```
Estágio 2: 1² = 1
Estágio 3: Resultado: 1
Saída final: Resultado: 1
------------------------------
Estágio 2: 2² = 4
Estágio 3: Resultado: 4
Saída final: Resultado: 4
------------------------------
Estágio 2: 3² = 9
Estágio 3: Resultado: 9
Saída final: Resultado: 9
------------------------------
Estágio 2: 4² = 16
Estágio 3: Resultado: 16
Saída final: Resultado: 16
------------------------------
Estágio 2: 5² = 25
Estágio 3: Resultado: 25
Saída final: Resultado: 25
------------------------------
```

### Pipeline com Múltiplos Estágios

Você pode criar pipelines mais complexos com vários estágios:

```go
package main

import (
    "fmt"
    "strings"
    "time"
)

// Pipeline para processar texto
func readLines(lines []string) <-chan string {
    out := make(chan string)
    go func() {
        defer close(out)
        for _, line := range lines {
            out <- line
        }
    }()
    return out
}

func toUpperCase(in <-chan string) <-chan string {
    out := make(chan string)
    go func() {
        defer close(out)
        for line := range in {
            out <- strings.ToUpper(line)
            time.Sleep(50 * time.Millisecond)
        }
    }()
    return out
}

func addPrefix(prefix string) func(<-chan string) <-chan string {
    return func(in <-chan string) <-chan string {
        out := make(chan string)
        go func() {
            defer close(out)
            for line := range in {
                out <- prefix + line
                time.Sleep(50 * time.Millisecond)
            }
        }()
        return out
    }
}

func filterEmpty(in <-chan string) <-chan string {
    out := make(chan string)
    go func() {
        defer close(out)
        for line := range in {
            if strings.TrimSpace(line) != "" {
                out <- line
            }
        }
    }()
    return out
}

func main() {
    lines := []string{"hello", "world", "", "go", "programming", ""}
    
    // Pipeline: read -> filter -> uppercase -> addPrefix
    pipeline := addPrefix("[PROCESSED] ")(
        toUpperCase(
            filterEmpty(
                readLines(lines),
            ),
        ),
    )
    
    for result := range pipeline {
        fmt.Println(result)
    }
}
```

**Saída:**
```
[PROCESSED] HELLO
[PROCESSED] WORLD
[PROCESSED] GO
[PROCESSED] PROGRAMMING
```

### Quando Usar Pipeline?

- ✅ **Processamento de Dados**: Transformações sequenciais em grandes volumes de dados
- ✅ **ETL (Extract, Transform, Load)**: Extrair, transformar e carregar dados
- ✅ **Streaming**: Processar dados em tempo real conforme chegam
- ✅ **Separação de Responsabilidades**: Cada estágio tem uma responsabilidade única

---

## 5. Worker Pool Pattern: Piscina de Trabalhadores

### O que é Worker Pool?

**Worker Pool** (também conhecido como Thread Pool) é um padrão que **mantém um número fixo de goroutines trabalhando em uma fila de tarefas**. É uma combinação de Fan-out com controle do número de workers simultâneos, evitando criar goroutines demais e esgotar recursos.

### Analogia

Imagine uma empresa com um número fixo de funcionários (workers). Quando chegam tarefas (jobs), elas vão para uma fila. Os funcionários pegam tarefas da fila, processam e pegam a próxima. Se todos estão ocupados, as tarefas esperam na fila.

### Implementação Completa

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

type Job struct {
    ID     int
    Data   string
    Result chan string
}

type WorkerPool struct {
    workers    int
    jobQueue   chan Job
    wg         sync.WaitGroup
}

func NewWorkerPool(workers int, queueSize int) *WorkerPool {
    return &WorkerPool{
        workers:  workers,
        jobQueue: make(chan Job, queueSize),
    }
}

func (wp *WorkerPool) Start() {
    for i := 1; i <= wp.workers; i++ {
        wp.wg.Add(1)
        go wp.worker(i)
    }
}

func (wp *WorkerPool) worker(id int) {
    defer wp.wg.Done()
    for job := range wp.jobQueue {
        // Processar job
        fmt.Printf("Worker %d: processando job %d (%s)\n", id, job.ID, job.Data)
        time.Sleep(1 * time.Second) // Simular trabalho
        
        // Enviar resultado
        result := fmt.Sprintf("Job %d processado por worker %d", job.ID, id)
        job.Result <- result
    }
}

func (wp *WorkerPool) Submit(job Job) {
    wp.jobQueue <- job
}

func (wp *WorkerPool) Stop() {
    close(wp.jobQueue)
    wp.wg.Wait()
}

func main() {
    // Criar worker pool com 3 workers e fila de 10 jobs
    pool := NewWorkerPool(3, 10)
    pool.Start()
    
    // Submeter jobs
    numJobs := 9
    results := make([]chan string, numJobs)
    
    for i := 0; i < numJobs; i++ {
        resultChan := make(chan string, 1)
        results[i] = resultChan
        
        job := Job{
            ID:     i + 1,
            Data:   fmt.Sprintf("dados do job %d", i+1),
            Result: resultChan,
        }
        
        pool.Submit(job)
        fmt.Printf("Job %d submetido\n", i+1)
    }
    
    // Coletar resultados
    for i, resultChan := range results {
        result := <-resultChan
        fmt.Printf("Resultado do job %d: %s\n", i+1, result)
    }
    
    // Parar o pool
    pool.Stop()
    fmt.Println("Worker pool encerrado")
}
```

### Worker Pool com Context para Cancelamento

```go
package main

import (
    "context"
    "fmt"
    "sync"
    "time"
)

type Task struct {
    ID   int
    Data string
}

type WorkerPoolWithContext struct {
    workers  int
    taskChan chan Task
    wg       sync.WaitGroup
}

func NewWorkerPoolWithContext(workers int) *WorkerPoolWithContext {
    return &WorkerPoolWithContext{
        workers:  workers,
        taskChan: make(chan Task),
    }
}

func (wp *WorkerPoolWithContext) Start(ctx context.Context) {
    for i := 1; i <= wp.workers; i++ {
        wp.wg.Add(1)
        go wp.worker(ctx, i)
    }
}

func (wp *WorkerPoolWithContext) worker(ctx context.Context, id int) {
    defer wp.wg.Done()
    for {
        select {
        case <-ctx.Done():
            fmt.Printf("Worker %d: cancelado\n", id)
            return
        case task, ok := <-wp.taskChan:
            if !ok {
                fmt.Printf("Worker %d: channel fechado\n", id)
                return
            }
            fmt.Printf("Worker %d: processando task %d\n", id, task.ID)
            time.Sleep(500 * time.Millisecond)
        }
    }
}

func (wp *WorkerPoolWithContext) Submit(task Task) {
    wp.taskChan <- task
}

func (wp *WorkerPoolWithContext) Stop() {
    close(wp.taskChan)
    wp.wg.Wait()
}

func main() {
    ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
    defer cancel()
    
    pool := NewWorkerPoolWithContext(3)
    pool.Start(ctx)
    
    // Submeter tarefas
    go func() {
        for i := 1; i <= 10; i++ {
            pool.Submit(Task{ID: i, Data: fmt.Sprintf("data %d", i)})
            time.Sleep(200 * time.Millisecond)
        }
    }()
    
    // Aguardar cancelamento ou conclusão
    <-ctx.Done()
    pool.Stop()
    fmt.Println("Pool encerrado")
}
```

### Quando Usar Worker Pool?

- ✅ **Limitar Recursos**: Controlar número máximo de goroutines simultâneas
- ✅ **Tarefas Repetitivas**: Processar muitos jobs similares de forma eficiente
- ✅ **Controle de Throughput**: Limitar taxa de processamento
- ✅ **Aplicações Web**: Processar requisições HTTP com pool de workers

---

## 6. Pub-Sub Pattern: Publicar-Assinar

### O que é Pub-Sub?

**Pub-Sub** (Publish-Subscribe) é um padrão onde **publicadores enviam mensagens sem saber quem são os assinantes**, e **assinantes recebem mensagens de tópicos que lhes interessam**. É útil para desacoplar componentes e permitir comunicação assíncrona.

### Analogia

Imagine um jornal (publisher) que publica notícias. Pessoas (subscribers) se inscrevem para receber notícias sobre tópicos específicos (esportes, tecnologia, etc.). O jornal não sabe quem são os assinantes, apenas publica. Os assinantes recebem apenas o que lhes interessa.

### Implementação Básica

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

type Message struct {
    Topic   string
    Content string
}

type PubSub struct {
    subscribers map[string][]chan Message
    mutex       sync.RWMutex
}

func NewPubSub() *PubSub {
    return &PubSub{
        subscribers: make(map[string][]chan Message),
    }
}

// Subscribe: inscrever-se em um tópico
func (ps *PubSub) Subscribe(topic string) <-chan Message {
    ps.mutex.Lock()
    defer ps.mutex.Unlock()
    
    ch := make(chan Message, 10) // Buffer para não bloquear
    ps.subscribers[topic] = append(ps.subscribers[topic], ch)
    
    return ch
}

// Publish: publicar mensagem em um tópico
func (ps *PubSub) Publish(topic string, content string) {
    ps.mutex.RLock()
    defer ps.mutex.RUnlock()
    
    msg := Message{Topic: topic, Content: content}
    
    // Enviar para todos os assinantes do tópico
    for _, ch := range ps.subscribers[topic] {
        select {
        case ch <- msg:
        default:
            // Channel cheio, pular (evitar bloqueio)
        }
    }
}

// Unsubscribe: remover assinante (simplificado)
func (ps *PubSub) Close() {
    ps.mutex.Lock()
    defer ps.mutex.Unlock()
    
    for _, channels := range ps.subscribers {
        for _, ch := range channels {
            close(ch)
        }
    }
    ps.subscribers = make(map[string][]chan Message)
}

func main() {
    pubsub := NewPubSub()
    
    // Assinante 1: interessado em "tech"
    sub1 := pubsub.Subscribe("tech")
    go func() {
        for msg := range sub1 {
            fmt.Printf("Assinante 1 [tech]: %s\n", msg.Content)
        }
    }()
    
    // Assinante 2: interessado em "tech" e "sports"
    sub2Tech := pubsub.Subscribe("tech")
    sub2Sports := pubsub.Subscribe("sports")
    
    go func() {
        for {
            select {
            case msg := <-sub2Tech:
                fmt.Printf("Assinante 2 [tech]: %s\n", msg.Content)
            case msg := <-sub2Sports:
                fmt.Printf("Assinante 2 [sports]: %s\n", msg.Content)
            }
        }
    }()
    
    // Assinante 3: interessado em "sports"
    sub3 := pubsub.Subscribe("sports")
    go func() {
        for msg := range sub3 {
            fmt.Printf("Assinante 3 [sports]: %s\n", msg.Content)
        }
    }()
    
    // Publicar mensagens
    time.Sleep(100 * time.Millisecond) // Dar tempo para assinantes se inscreverem
    
    pubsub.Publish("tech", "Nova versão do Go lançada!")
    time.Sleep(100 * time.Millisecond)
    
    pubsub.Publish("sports", "Brasil vence a copa!")
    time.Sleep(100 * time.Millisecond)
    
    pubsub.Publish("tech", "Concorrência em Go é incrível!")
    time.Sleep(100 * time.Millisecond)
    
    pubsub.Publish("sports", "Jogo de futebol hoje às 20h")
    time.Sleep(100 * time.Millisecond)
    
    // Dar tempo para processar
    time.Sleep(1 * time.Second)
    
    pubsub.Close()
    fmt.Println("Pub-Sub encerrado")
}
```

**Saída possível:**
```
Assinante 1 [tech]: Nova versão do Go lançada!
Assinante 2 [tech]: Nova versão do Go lançada!
Assinante 2 [sports]: Brasil vence a copa!
Assinante 3 [sports]: Brasil vence a copa!
Assinante 1 [tech]: Concorrência em Go é incrível!
Assinante 2 [tech]: Concorrência em Go é incrível!
Assinante 2 [sports]: Jogo de futebol hoje às 20h
Assinante 3 [sports]: Jogo de futebol hoje às 20h
Pub-Sub encerrado
```

### Pub-Sub com Context

```go
package main

import (
    "context"
    "fmt"
    "sync"
    "time"
)

type PubSubWithContext struct {
    subscribers map[string][]chan Message
    mutex       sync.RWMutex
}

func NewPubSubWithContext() *PubSubWithContext {
    return &PubSubWithContext{
        subscribers: make(map[string][]chan Message),
    }
}

func (ps *PubSubWithContext) Subscribe(ctx context.Context, topic string) <-chan Message {
    ps.mutex.Lock()
    ch := make(chan Message, 10)
    ps.subscribers[topic] = append(ps.subscribers[topic], ch)
    ps.mutex.Unlock()
    
    // Goroutine para fechar channel quando contexto for cancelado
    go func() {
        <-ctx.Done()
        ps.mutex.Lock()
        // Remover channel da lista (simplificado)
        ps.mutex.Unlock()
        close(ch)
    }()
    
    return ch
}

func (ps *PubSubWithContext) Publish(topic string, content string) {
    ps.mutex.RLock()
    defer ps.mutex.RUnlock()
    
    msg := Message{Topic: topic, Content: content}
    for _, ch := range ps.subscribers[topic] {
        select {
        case ch <- msg:
        default:
        }
    }
}

func main() {
    pubsub := NewPubSubWithContext()
    
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    
    // Assinante com timeout
    sub := pubsub.Subscribe(ctx, "news")
    go func() {
        for msg := range sub {
            fmt.Printf("Recebido: %s\n", msg.Content)
        }
        fmt.Println("Assinatura expirada")
    }()
    
    // Publicar mensagens
    for i := 1; i <= 10; i++ {
        pubsub.Publish("news", fmt.Sprintf("Notícia %d", i))
        time.Sleep(500 * time.Millisecond)
    }
    
    time.Sleep(1 * time.Second)
}
```

### Quando Usar Pub-Sub?

- ✅ **Desacoplamento**: Componentes não precisam conhecer uns aos outros
- ✅ **Eventos**: Sistema de eventos onde múltiplos componentes reagem
- ✅ **Notificações**: Sistema de notificações para múltiplos clientes
- ✅ **Microserviços**: Comunicação assíncrona entre serviços

---

## 7. Combinando Padrões: Exemplo Prático

Vamos criar um exemplo que combina vários padrões:

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

// Sistema de processamento de dados que usa múltiplos padrões
func main() {
    // 1. Pipeline: gerar -> processar -> formatar
    dataChan := generateData(10)
    processedChan := processData(dataChan)
    formattedChan := formatData(processedChan)
    
    // 2. Fan-out: distribuir formatação para múltiplos workers
    output1 := worker("Worker 1", formattedChan)
    output2 := worker("Worker 2", formattedChan)
    output3 := worker("Worker 3", formattedChan)
    
    // 3. Fan-in: mesclar resultados dos workers
    finalOutput := fanIn(output1, output2, output3)
    
    // Consumir resultado final
    for result := range finalOutput {
        fmt.Println(result)
    }
}

func generateData(count int) <-chan int {
    out := make(chan int)
    go func() {
        defer close(out)
        for i := 1; i <= count; i++ {
            out <- i
            time.Sleep(100 * time.Millisecond)
        }
    }()
    return out
}

func processData(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        defer close(out)
        for n := range in {
            // Processar: elevar ao quadrado
            out <- n * n
            time.Sleep(50 * time.Millisecond)
        }
    }()
    return out
}

func formatData(in <-chan int) <-chan string {
    out := make(chan string)
    go func() {
        defer close(out)
        for n := range in {
            out <- fmt.Sprintf("Valor processado: %d", n)
        }
    }()
    return out
}

func worker(name string, in <-chan string) <-chan string {
    out := make(chan string)
    go func() {
        defer close(out)
        for msg := range in {
            result := fmt.Sprintf("[%s] %s", name, msg)
            out <- result
            time.Sleep(100 * time.Millisecond)
        }
    }()
    return out
}

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
        close(out)
    }()
    
    return out
}
```

---

## 8. Resumo dos Padrões

### Fan-In
- **Propósito**: Mesclar múltiplos channels em um único channel
- **Uso**: Agregar resultados de múltiplas goroutines
- **Implementação**: `select` ou múltiplas goroutines lendo e escrevendo

### Fan-Out
- **Propósito**: Distribuir trabalho de um channel para múltiplos workers
- **Uso**: Paralelizar processamento de tarefas
- **Implementação**: Múltiplas goroutines lendo do mesmo channel

### Pipeline
- **Propósito**: Encadear estágios de processamento
- **Uso**: Transformações sequenciais de dados
- **Implementação**: Cada estágio é uma goroutine com channel de entrada e saída

### Worker Pool
- **Propósito**: Manter número fixo de workers processando jobs
- **Uso**: Controlar recursos e throughput
- **Implementação**: Pool de goroutines + fila de jobs

### Pub-Sub
- **Propósito**: Comunicação desacoplada via tópicos
- **Uso**: Eventos, notificações, microserviços
- **Implementação**: Map de tópicos para slices de channels

---

## 9. Boas Práticas

### 1. Sempre Feche Channels

```go
// ✅ CORRETO
go func() {
    defer close(output)
    for item := range input {
        output <- process(item)
    }
}()

// ❌ ERRADO
go func() {
    for item := range input {
        output <- process(item) // Channel nunca é fechado!
    }
}()
```

### 2. Use Buffered Channels Quando Apropriado

```go
// ✅ Para evitar bloqueio quando produtor é mais rápido que consumidor
output := make(chan string, 100)

// ✅ Para Pub-Sub, evitar bloquear publisher
ch := make(chan Message, 10)
```

### 3. Use Context para Cancelamento

```go
// ✅ Sempre passe context para operações que podem ser canceladas
func worker(ctx context.Context, jobs <-chan Job) {
    for {
        select {
        case <-ctx.Done():
            return
        case job := <-jobs:
            process(job)
        }
    }
}
```

### 4. Evite Vazamentos de Goroutines

```go
// ✅ Use WaitGroup para garantir que goroutines terminem
var wg sync.WaitGroup
for i := 0; i < 10; i++ {
    wg.Add(1)
    go func() {
        defer wg.Done()
        // trabalho
    }()
}
wg.Wait()
```

### 5. Prefira Channels sobre Shared Memory

```go
// ✅ CORRETO: Comunicação via channels
result := make(chan int)
go func() {
    result <- calculate()
}()

// ❌ ERRADO: Compartilhar memória diretamente
var result int
go func() {
    result = calculate() // Race condition!
}()
```

---

E assim terminamos nossa primeira aula sobre Concurrency Patterns! Você agora entende:
- O que são padrões de concorrência e por que são importantes
- Como implementar Fan-in para mesclar múltiplas entradas
- Como implementar Fan-out para distribuir trabalho
- Como criar Pipelines para encadear processamento
- Como usar Worker Pools para controlar recursos
- Como implementar Pub-Sub para comunicação desacoplada
- Como combinar padrões para resolver problemas complexos

Na próxima aula, vamos simplificar esses conceitos com analogias do dia a dia para fixar ainda mais o aprendizado!

Sinta-se à vontade para reler este material e experimentar com os exemplos. Se tiver qualquer dúvida, pode perguntar!

