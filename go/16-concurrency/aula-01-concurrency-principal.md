# Módulo 16: Concurrency em Go

## Aula 1: Concurrency - Programação Concorrente e Paralela

Olá! Bem-vindo ao décimo sexto módulo. Até agora você aprendeu sobre funções, métodos, interfaces, generics, tratamento de erros, organização de código e muitos outros conceitos fundamentais de Go. Mas você já se perguntou: **"Como meu programa pode fazer várias coisas ao mesmo tempo? Como processar múltiplas tarefas simultaneamente de forma eficiente?"**

Go foi projetada desde o início com **concorrência** em mente. Enquanto outras linguagens adicionaram concorrência como uma funcionalidade posterior, Go a tem como um cidadão de primeira classe. A concorrência permite que seu programa execute múltiplas operações simultaneamente, aproveitando melhor os recursos do sistema, especialmente em sistemas com múltiplos núcleos de CPU.

Nesta aula, vamos mergulhar profundamente na concorrência em Go: entender goroutines (threads leves), channels (canais de comunicação), o statement `select`, a diferença entre channels buffered e unbuffered, worker pools, o pacote `sync` e seus primitivos de sincronização como Mutexes e WaitGroups.

---

## 1. O que é Concorrência?

### Concorrência vs Paralelismo

Antes de começarmos, é importante entender a diferença entre **concorrência** e **paralelismo**:

- **Concorrência**: É sobre **estruturar** um programa para que ele possa lidar com múltiplas tarefas ao mesmo tempo. Não necessariamente significa que as tarefas estão sendo executadas simultaneamente.
- **Paralelismo**: É sobre **executar** múltiplas tarefas simultaneamente, geralmente em múltiplos núcleos de CPU.

**Analogia**: Imagine um garçom em um restaurante:
- **Concorrência**: O garçom atende várias mesas, alternando entre elas (estrutura para lidar com múltiplas tarefas)
- **Paralelismo**: Vários garçons atendendo mesas ao mesmo tempo (execução simultânea)

Go suporta ambos! Você pode escrever código concorrente que o runtime do Go pode executar em paralelo quando há múltiplos núcleos disponíveis.

### Por que Concorrência é Importante?

1. **Aproveitamento de Recursos**: Em vez de esperar uma operação I/O (como ler um arquivo ou fazer uma requisição HTTP) terminar antes de fazer outra, você pode fazer várias ao mesmo tempo
2. **Responsividade**: Interfaces de usuário podem permanecer responsivas enquanto processam dados em background
3. **Performance**: Em sistemas com múltiplos núcleos, você pode distribuir o trabalho entre eles
4. **Escalabilidade**: Servidores podem lidar com milhares de requisições simultâneas

### A Filosofia do Go: "Don't communicate by sharing memory; share memory by communicating"

Go adota uma filosofia única: em vez de compartilhar memória entre threads (o que pode causar race conditions e bugs difíceis de encontrar), Go usa **channels** para comunicação entre goroutines. Isso torna o código mais seguro e mais fácil de entender.

---

## 2. Goroutines: Threads Leves

### O que são Goroutines?

**Goroutines** são funções ou métodos que são executados de forma independente e simultânea com outras goroutines. Elas são muito mais leves que threads tradicionais do sistema operacional.

### Características das Goroutines

- ✅ **Leves**: Uma goroutine ocupa apenas alguns KB de memória (threads tradicionais podem ocupar MB)
- ✅ **Baratas**: Você pode criar milhares ou até milhões de goroutines
- ✅ **Gerenciadas pelo Runtime**: O runtime do Go gerencia o agendamento das goroutines
- ✅ **Multiplexadas**: Múltiplas goroutines são multiplexadas em um número menor de threads do OS

### Criando uma Goroutine

Para executar uma função como goroutine, simplesmente adicione a palavra-chave `go` antes da chamada da função:

```go
package main

import (
    "fmt"
    "time"
)

func dizerOla() {
    fmt.Println("Olá do goroutine!")
}

func main() {
    // Executa dizerOla() como goroutine
    go dizerOla()
    
    // A função main continua executando
    fmt.Println("Olá do main!")
    
    // Damos um tempo para a goroutine executar
    // (em produção, você usaria channels ou WaitGroups)
    time.Sleep(1 * time.Second)
}
```

**Saída possível:**
```
Olá do main!
Olá do goroutine!
```

**Importante**: A função `main` não espera que as goroutines terminem. Se `main` terminar, todas as goroutines são encerradas.

### Exemplo: Múltiplas Goroutines

```go
package main

import (
    "fmt"
    "time"
)

func contar(nome string, vezes int) {
    for i := 1; i <= vezes; i++ {
        fmt.Printf("%s: %d\n", nome, i)
        time.Sleep(100 * time.Millisecond)
    }
}

func main() {
    // Criar múltiplas goroutines
    go contar("Goroutine 1", 5)
    go contar("Goroutine 2", 5)
    go contar("Goroutine 3", 5)
    
    // Esperar um pouco para ver as goroutines executando
    time.Sleep(2 * time.Second)
    
    fmt.Println("Main terminou")
}
```

**Saída possível** (a ordem pode variar):
```
Goroutine 2: 1
Goroutine 1: 1
Goroutine 3: 1
Goroutine 2: 2
Goroutine 1: 2
Goroutine 3: 2
...
```

### Goroutines com Funções Anônimas

Você também pode usar goroutines com funções anônimas:

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    // Goroutine com função anônima
    go func() {
        for i := 0; i < 5; i++ {
            fmt.Printf("Anônima: %d\n", i)
            time.Sleep(100 * time.Millisecond)
        }
    }()
    
    time.Sleep(1 * time.Second)
}
```

### Quando Usar Goroutines?

- ✅ **Operações I/O**: Leitura de arquivos, requisições HTTP, operações de banco de dados
- ✅ **Processamento paralelo**: Processar múltiplos itens de uma lista simultaneamente
- ✅ **Background tasks**: Tarefas que devem rodar em segundo plano
- ✅ **Servidores**: Lidar com múltiplas requisições simultaneamente

---

## 3. Channels: Comunicação entre Goroutines

### O que são Channels?

**Channels** são os mecanismos primários de comunicação entre goroutines em Go. Eles permitem que goroutines enviem e recebam valores entre si, seguindo o princípio: **"Don't communicate by sharing memory; share memory by communicating"**.

### Criando um Channel

Channels são criados usando a função `make()`:

```go
// Channel não-buffered (síncrono)
ch := make(chan int)

// Channel buffered (assíncrono até a capacidade)
ch := make(chan int, 10) // capacidade de 10
```

### Enviando e Recebendo Valores

```go
package main

import "fmt"

func main() {
    // Criar um channel
    ch := make(chan string)
    
    // Goroutine que envia um valor
    go func() {
        ch <- "Olá do channel!" // Enviar valor
    }()
    
    // Receber valor do channel
    mensagem := <-ch
    fmt.Println(mensagem)
}
```

**Operadores:**
- `ch <- valor`: Envia `valor` para o channel `ch`
- `valor := <-ch`: Recebe um valor do channel `ch`

### Channels são Tipados

Channels são fortemente tipados. Um channel de `int` só pode enviar/receber `int`:

```go
chInt := make(chan int)
chString := make(chan string)

chInt <- 42        // ✅ OK
chString <- "olá"  // ✅ OK
chInt <- "erro"    // ❌ Erro de compilação
```

### Exemplo: Worker Enviando Resultado

```go
package main

import (
    "fmt"
    "time"
)

func calcularQuadrado(n int, ch chan int) {
    resultado := n * n
    time.Sleep(100 * time.Millisecond) // Simular trabalho
    ch <- resultado // Enviar resultado
}

func main() {
    ch := make(chan int)
    
    // Iniciar goroutine
    go calcularQuadrado(5, ch)
    
    // Receber resultado
    resultado := <-ch
    fmt.Printf("Quadrado de 5: %d\n", resultado)
}
```

### Múltiplos Valores

```go
package main

import "fmt"

func gerarNumeros(ch chan int) {
    for i := 1; i <= 5; i++ {
        ch <- i
    }
    close(ch) // Fechar channel quando terminar
}

func main() {
    ch := make(chan int)
    
    go gerarNumeros(ch)
    
    // Receber valores até o channel ser fechado
    for valor := range ch {
        fmt.Println(valor)
    }
}
```

### Fechando Channels

Quando você termina de enviar valores, pode fechar o channel com `close()`:

```go
close(ch)
```

**Regras:**
- Apenas o sender (quem envia) deve fechar o channel
- Enviar em um channel fechado causa panic
- Receber de um channel fechado retorna o zero value e `false` como segundo valor

```go
valor, ok := <-ch
if !ok {
    fmt.Println("Channel foi fechado")
}
```

---

## 4. Buffered vs Unbuffered Channels

### Unbuffered Channels (Canais Não-Bufferizados)

**Unbuffered channels** são síncronos. O sender bloqueia até que um receiver esteja pronto para receber o valor, e vice-versa.

```go
ch := make(chan int) // Unbuffered
```

**Características:**
- ✅ **Síncrono**: Garante que o valor foi recebido antes de continuar
- ✅ **Coordenação**: Perfeito para sincronizar goroutines
- ⚠️ **Bloqueante**: Sender e receiver devem estar prontos simultaneamente

**Exemplo:**

```go
package main

import "fmt"

func main() {
    ch := make(chan int) // Unbuffered
    
    go func() {
        fmt.Println("Enviando 42...")
        ch <- 42 // Bloqueia até alguém receber
        fmt.Println("42 foi enviado!")
    }()
    
    fmt.Println("Recebendo...")
    valor := <-ch // Bloqueia até receber
    fmt.Printf("Recebido: %d\n", valor)
}
```

### Buffered Channels (Canais Bufferizados)

**Buffered channels** têm uma capacidade. O sender só bloqueia quando o buffer está cheio, e o receiver só bloqueia quando o buffer está vazio.

```go
ch := make(chan int, 3) // Buffered com capacidade 3
```

**Características:**
- ✅ **Assíncrono**: Permite enviar múltiplos valores sem bloquear imediatamente
- ✅ **Performance**: Reduz bloqueios desnecessários
- ✅ **Decoupling**: Desacopla sender e receiver temporariamente
- ⚠️ **Cuidado**: Pode mascarar problemas de sincronização

**Exemplo:**

```go
package main

import "fmt"

func main() {
    ch := make(chan int, 3) // Buffered com capacidade 3
    
    // Enviar 3 valores sem bloquear
    ch <- 1
    ch <- 2
    ch <- 3
    fmt.Println("3 valores enviados")
    
    // Agora o buffer está cheio, próxima operação bloquearia
    // ch <- 4 // Isso bloquearia
    
    // Receber valores
    fmt.Println(<-ch) // 1
    fmt.Println(<-ch) // 2
    fmt.Println(<-ch) // 3
}
```

### Quando Usar Cada Um?

**Use Unbuffered quando:**
- ✅ Precisa garantir que o valor foi recebido antes de continuar
- ✅ Precisa sincronizar goroutines
- ✅ O receiver deve processar imediatamente

**Use Buffered quando:**
- ✅ Sender e receiver operam em velocidades diferentes
- ✅ Precisa de throughput maior
- ✅ Quer desacoplar temporariamente sender e receiver
- ✅ Processando batches de dados

**Exemplo Prático: Unbuffered para Sincronização**

```go
package main

import (
    "fmt"
    "time"
)

func worker(done chan bool) {
    fmt.Println("Trabalhando...")
    time.Sleep(2 * time.Second)
    fmt.Println("Terminei!")
    done <- true // Sinalizar que terminou
}

func main() {
    done := make(chan bool) // Unbuffered para sincronização
    
    go worker(done)
    
    <-done // Esperar worker terminar
    fmt.Println("Main pode continuar")
}
```

**Exemplo Prático: Buffered para Throughput**

```go
package main

import "fmt"

func main() {
    // Buffered channel para permitir múltiplos valores
    ch := make(chan int, 10)
    
    // Enviar múltiplos valores rapidamente
    for i := 0; i < 10; i++ {
        ch <- i
    }
    
    // Processar depois
    for i := 0; i < 10; i++ {
        fmt.Println(<-ch)
    }
}
```

---

## 5. Select Statement: Multiplexando Channels

### O que é o Select?

O `select` statement permite que uma goroutine aguarde múltiplas operações de channel simultaneamente. Ele é similar ao `switch`, mas para channels.

### Sintaxe Básica

```go
select {
case valor := <-ch1:
    // Processar valor de ch1
case valor := <-ch2:
    // Processar valor de ch2
case ch3 <- valor:
    // Enviar valor para ch3
default:
    // Executar se nenhum case estiver pronto
}
```

### Exemplo: Múltiplos Channels

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    ch1 := make(chan string)
    ch2 := make(chan string)
    
    go func() {
        time.Sleep(1 * time.Second)
        ch1 <- "do channel 1"
    }()
    
    go func() {
        time.Sleep(2 * time.Second)
        ch2 <- "do channel 2"
    }()
    
    // Select aguarda o primeiro channel que estiver pronto
    select {
    case msg1 := <-ch1:
        fmt.Println("Recebido", msg1)
    case msg2 := <-ch2:
        fmt.Println("Recebido", msg2)
    }
}
```

### Select com Loop

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    ch1 := make(chan int)
    ch2 := make(chan int)
    
    go func() {
        for i := 0; i < 5; i++ {
            ch1 <- i
            time.Sleep(100 * time.Millisecond)
        }
        close(ch1)
    }()
    
    go func() {
        for i := 10; i < 15; i++ {
            ch2 <- i
            time.Sleep(150 * time.Millisecond)
        }
        close(ch2)
    }()
    
    ch1Fechado := false
    ch2Fechado := false
    
    for {
        select {
        case valor, ok := <-ch1:
            if !ok {
                ch1Fechado = true
            } else {
                fmt.Printf("Ch1: %d\n", valor)
            }
        case valor, ok := <-ch2:
            if !ok {
                ch2Fechado = true
            } else {
                fmt.Printf("Ch2: %d\n", valor)
            }
        }
        
        if ch1Fechado && ch2Fechado {
            break
        }
    }
}
```

### Select com Default (Non-Blocking)

O `default` case permite que o `select` não bloqueie:

```go
package main

import "fmt"

func main() {
    ch := make(chan int)
    
    select {
    case valor := <-ch:
        fmt.Println("Recebido:", valor)
    default:
        fmt.Println("Nenhum valor disponível, não bloqueando")
    }
    
    fmt.Println("Continuando...")
}
```

### Select para Timeouts

Um padrão muito comum é usar `select` com `time.After` para implementar timeouts:

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    ch := make(chan string)
    
    go func() {
        time.Sleep(3 * time.Second)
        ch <- "resultado"
    }()
    
    select {
    case resultado := <-ch:
        fmt.Println("Recebido:", resultado)
    case <-time.After(2 * time.Second):
        fmt.Println("Timeout! Operação demorou muito")
    }
}
```

### Select para Cancelamento

```go
package main

import (
    "fmt"
    "time"
)

func worker(done <-chan bool) {
    for {
        select {
        case <-done:
            fmt.Println("Worker cancelado")
            return
        default:
            fmt.Println("Trabalhando...")
            time.Sleep(500 * time.Millisecond)
        }
    }
}

func main() {
    done := make(chan bool)
    
    go worker(done)
    
    time.Sleep(2 * time.Second)
    close(done) // Cancelar worker
    
    time.Sleep(500 * time.Millisecond)
}
```

---

## 6. Worker Pools: Padrão de Concorrência

### O que são Worker Pools?

**Worker Pools** são um padrão de concorrência onde um número fixo de goroutines (workers) processa tarefas de uma fila compartilhada. Isso permite controlar o uso de recursos enquanto mantém paralelismo.

### Por que Usar Worker Pools?

- ✅ **Controle de Recursos**: Limita o número de goroutines simultâneas
- ✅ **Rate Limiting**: Controla a taxa de processamento
- ✅ **Eficiência**: Evita criar goroutines demais
- ✅ **Previsibilidade**: Comportamento previsível sob carga

### Implementação Básica

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
    for job := range jobs {
        fmt.Printf("Worker %d processando job %d\n", id, job)
        time.Sleep(1 * time.Second) // Simular trabalho
        results <- job * 2 // Enviar resultado
    }
}

func main() {
    const numWorkers = 3
    const numJobs = 5
    
    jobs := make(chan int, numJobs)
    results := make(chan int, numJobs)
    
    // Criar workers
    for w := 1; w <= numWorkers; w++ {
        go worker(w, jobs, results)
    }
    
    // Enviar jobs
    for j := 1; j <= numJobs; j++ {
        jobs <- j
    }
    close(jobs)
    
    // Coletar resultados
    for r := 1; r <= numJobs; r++ {
        fmt.Printf("Resultado: %d\n", <-results)
    }
}
```

### Worker Pool com WaitGroup

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
    defer wg.Done()
    
    for job := range jobs {
        fmt.Printf("Worker %d processando job %d\n", id, job)
        time.Sleep(500 * time.Millisecond)
        fmt.Printf("Worker %d completou job %d\n", id, job)
    }
}

func main() {
    const numWorkers = 3
    const numJobs = 10
    
    jobs := make(chan int, numJobs)
    var wg sync.WaitGroup
    
    // Criar workers
    for w := 1; w <= numWorkers; w++ {
        wg.Add(1)
        go worker(w, jobs, &wg)
    }
    
    // Enviar jobs
    for j := 1; j <= numJobs; j++ {
        jobs <- j
    }
    close(jobs)
    
    // Esperar todos os workers terminarem
    wg.Wait()
    fmt.Println("Todos os jobs foram processados")
}
```

### Exemplo: Processar URLs

```go
package main

import (
    "fmt"
    "net/http"
    "sync"
    "time"
)

func fetchURL(url string) {
    start := time.Now()
    resp, err := http.Get(url)
    duration := time.Since(start)
    
    if err != nil {
        fmt.Printf("Erro ao buscar %s: %v\n", url, err)
        return
    }
    defer resp.Body.Close()
    
    fmt.Printf("%s: %d (tempo: %v)\n", url, resp.StatusCode, duration)
}

func main() {
    urls := []string{
        "https://www.google.com",
        "https://www.github.com",
        "https://www.stackoverflow.com",
        "https://www.reddit.com",
        "https://www.twitter.com",
    }
    
    const numWorkers = 3
    jobs := make(chan string, len(urls))
    var wg sync.WaitGroup
    
    // Criar workers
    for w := 1; w <= numWorkers; w++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            for url := range jobs {
                fetchURL(url)
            }
        }()
    }
    
    // Enviar URLs
    for _, url := range urls {
        jobs <- url
    }
    close(jobs)
    
    wg.Wait()
    fmt.Println("Todas as requisições foram processadas")
}
```

---

## 7. O Pacote `sync`: Sincronização

### O que é o Pacote `sync`?

O pacote `sync` fornece primitivos de sincronização de baixo nível para coordenar goroutines e proteger acesso a dados compartilhados. Embora Go prefira comunicação via channels, às vezes você precisa de sincronização mais direta.

### Componentes Principais

- **Mutex**: Exclusão mútua para proteger dados compartilhados
- **RWMutex**: Mutex de leitura/escrita (múltiplos leitores, um escritor)
- **WaitGroup**: Aguardar múltiplas goroutines terminarem
- **Once**: Executar algo apenas uma vez
- **Cond**: Variável de condição para sincronização
- **Map**: Map thread-safe

---

## 8. Mutexes: Proteção de Dados Compartilhados

### O que são Mutexes?

**Mutex** (Mutual Exclusion) é um lock que garante que apenas uma goroutine acesse um recurso compartilhado por vez. Isso previne **race conditions**.

### Race Condition

Uma race condition ocorre quando múltiplas goroutines acessam dados compartilhados simultaneamente sem sincronização:

```go
package main

import (
    "fmt"
    "sync"
)

var contador int

func incrementar() {
    contador++ // ❌ Race condition!
}

func main() {
    var wg sync.WaitGroup
    
    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            incrementar()
        }()
    }
    
    wg.Wait()
    fmt.Println("Contador:", contador) // Pode não ser 1000!
}
```

### Usando Mutex

```go
package main

import (
    "fmt"
    "sync"
)

var (
    contador int
    mu       sync.Mutex
)

func incrementar() {
    mu.Lock()         // Bloquear
    defer mu.Unlock() // Desbloquear (sempre usar defer!)
    contador++
}

func main() {
    var wg sync.WaitGroup
    
    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            incrementar()
        }()
    }
    
    wg.Wait()
    fmt.Println("Contador:", contador) // Sempre 1000!
}
```

### Métodos do Mutex

- `Lock()`: Bloqueia o mutex. Se já estiver bloqueado, espera.
- `Unlock()`: Desbloqueia o mutex.
- **Sempre use `defer mu.Unlock()`** para garantir que o unlock aconteça mesmo em caso de panic.

### Exemplo: Contador Thread-Safe

```go
package main

import (
    "fmt"
    "sync"
)

type ContadorSeguro struct {
    mu    sync.Mutex
    valor int
}

func (c *ContadorSeguro) Incrementar() {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.valor++
}

func (c *ContadorSeguro) Valor() int {
    c.mu.Lock()
    defer c.mu.Unlock()
    return c.valor
}

func main() {
    contador := &ContadorSeguro{}
    var wg sync.WaitGroup
    
    for i := 0; i < 100; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            contador.Incrementar()
        }()
    }
    
    wg.Wait()
    fmt.Println("Valor final:", contador.Valor())
}
```

### RWMutex: Múltiplos Leitores

`RWMutex` permite múltiplas leituras simultâneas ou uma escrita exclusiva:

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

type Cache struct {
    mu   sync.RWMutex
    data map[string]string
}

func NewCache() *Cache {
    return &Cache{
        data: make(map[string]string),
    }
}

func (c *Cache) Get(key string) string {
    c.mu.RLock()         // Lock para leitura
    defer c.mu.RUnlock() // Unlock para leitura
    return c.data[key]
}

func (c *Cache) Set(key, value string) {
    c.mu.Lock()         // Lock para escrita (exclusivo)
    defer c.mu.Unlock() // Unlock para escrita
    c.data[key] = value
}

func main() {
    cache := NewCache()
    var wg sync.WaitGroup
    
    // Múltiplos leitores
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            for j := 0; j < 5; j++ {
                valor := cache.Get("chave")
                fmt.Printf("Leitor %d: %s\n", id, valor)
                time.Sleep(10 * time.Millisecond)
            }
        }(i)
    }
    
    // Um escritor
    wg.Add(1)
    go func() {
        defer wg.Done()
        for i := 0; i < 5; i++ {
            cache.Set("chave", fmt.Sprintf("valor-%d", i))
            time.Sleep(50 * time.Millisecond)
        }
    }()
    
    wg.Wait()
}
```

**Quando usar RWMutex:**
- ✅ Muitas leituras e poucas escritas
- ✅ Leituras são operações caras
- ⚠️ Se leituras são rápidas, Mutex simples pode ser mais eficiente

---

## 9. WaitGroups: Aguardar Goroutines

### O que são WaitGroups?

**WaitGroup** é uma primitiva de sincronização que permite aguardar um grupo de goroutines terminar antes de continuar.

### Métodos do WaitGroup

- `Add(delta int)`: Incrementa o contador em `delta`
- `Done()`: Decrementa o contador em 1 (equivalente a `Add(-1)`)
- `Wait()`: Bloqueia até o contador chegar a zero

### Exemplo Básico

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

func worker(id int, wg *sync.WaitGroup) {
    defer wg.Done() // Decrementar contador quando terminar
    
    fmt.Printf("Worker %d começando\n", id)
    time.Sleep(time.Second) // Simular trabalho
    fmt.Printf("Worker %d terminando\n", id)
}

func main() {
    var wg sync.WaitGroup
    
    // Iniciar 3 workers
    for i := 1; i <= 3; i++ {
        wg.Add(1) // Incrementar contador
        go worker(i, &wg)
    }
    
    wg.Wait() // Esperar todos terminarem
    fmt.Println("Todos os workers terminaram")
}
```

### Exemplo: Processar Lista em Paralelo

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

func processarItem(item int, wg *sync.WaitGroup) {
    defer wg.Done()
    
    fmt.Printf("Processando item %d\n", item)
    time.Sleep(100 * time.Millisecond) // Simular processamento
    fmt.Printf("Item %d processado\n", item)
}

func main() {
    items := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    var wg sync.WaitGroup
    
    for _, item := range items {
        wg.Add(1)
        go processarItem(item, &wg)
    }
    
    wg.Wait()
    fmt.Println("Todos os itens foram processados")
}
```

### WaitGroup com Resultados

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

func calcularQuadrado(indice int, numero int, resultados []int, wg *sync.WaitGroup) {
    defer wg.Done()
    
    time.Sleep(100 * time.Millisecond) // Simular trabalho
    resultados[indice] = numero * numero
}

func main() {
    numeros := []int{1, 2, 3, 4, 5}
    resultados := make([]int, len(numeros))
    var wg sync.WaitGroup
    
    for i, num := range numeros {
        wg.Add(1)
        go calcularQuadrado(i, num, resultados, &wg)
    }
    
    wg.Wait()
    
    for i, num := range numeros {
        fmt.Printf("%d² = %d\n", num, resultados[i])
    }
}
```

### Importante: Passar WaitGroup por Referência

Sempre passe `WaitGroup` por **referência** (`*sync.WaitGroup`), não por valor:

```go
// ❌ ERRADO
func worker(wg sync.WaitGroup) { ... }

// ✅ CORRETO
func worker(wg *sync.WaitGroup) { ... }
```

---

## 10. Padrões Comuns de Concorrência

### Padrão: Fan-Out / Fan-In

**Fan-Out**: Distribuir trabalho para múltiplos workers
**Fan-In**: Combinar resultados de múltiplos workers

```go
package main

import (
    "fmt"
    "sync"
)

func producer(nums []int) <-chan int {
    out := make(chan int)
    go func() {
        defer close(out)
        for _, n := range nums {
            out <- n
        }
    }()
    return out
}

func square(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        defer close(out)
        for n := range in {
            out <- n * n
        }
    }()
    return out
}

func merge(channels ...<-chan int) <-chan int {
    var wg sync.WaitGroup
    out := make(chan int)
    
    // Fan-in: receber de todos os channels
    output := func(c <-chan int) {
        defer wg.Done()
        for n := range c {
            out <- n
        }
    }
    
    wg.Add(len(channels))
    for _, c := range channels {
        go output(c)
    }
    
    go func() {
        wg.Wait()
        close(out)
    }()
    
    return out
}

func main() {
    nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    
    // Fan-out: distribuir trabalho para múltiplos workers
    in := producer(nums)
    
    // Criar 3 workers que processam do mesmo channel
    // Os valores serão distribuídos automaticamente entre os workers
    c1 := square(in)
    c2 := square(in)
    c3 := square(in)
    
    // Fan-in: combinar resultados de todos os workers em um único channel
    for n := range merge(c1, c2, c3) {
        fmt.Println(n)
    }
}
```

**Como funciona:**
- **Fan-Out**: O channel `in` é lido por 3 workers diferentes (`c1`, `c2`, `c3`). Cada valor enviado para `in` será processado por apenas um dos workers (distribuição automática do Go runtime).
- **Fan-In**: A função `merge` combina os resultados de todos os workers em um único channel de saída, permitindo processar todos os resultados de forma unificada.

### Padrão: Pipeline

```go
package main

import "fmt"

func gerar(nums ...int) <-chan int {
    out := make(chan int)
    go func() {
        for _, n := range nums {
            out <- n
        }
        close(out)
    }()
    return out
}

func quadrado(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        for n := range in {
            out <- n * n
        }
        close(out)
    }()
    return out
}

func main() {
    // Pipeline: gerar -> quadrado -> imprimir
    for n := range quadrado(gerar(1, 2, 3, 4, 5)) {
        fmt.Println(n)
    }
}
```

---

## 11. Boas Práticas e Armadilhas Comuns

### ✅ Boas Práticas

1. **Sempre feche channels quando terminar de enviar**
   ```go
   defer close(ch)
   ```

2. **Use defer para Unlock de Mutex**
   ```go
   mu.Lock()
   defer mu.Unlock()
   ```

3. **Passe WaitGroup por referência**
   ```go
   func worker(wg *sync.WaitGroup) { ... }
   ```

4. **Prefira channels para comunicação, Mutex para sincronização simples**
   - Channels: Comunicação entre goroutines
   - Mutex: Proteção de dados compartilhados

5. **Use buffered channels quando apropriado**
   - Para desacoplar sender e receiver
   - Para melhorar throughput

6. **Sempre verifique se channel está fechado**
   ```go
   valor, ok := <-ch
   if !ok {
       // channel fechado
   }
   ```

### ⚠️ Armadilhas Comuns

1. **Goroutine leak**: Não esquecer de fechar channels ou aguardar goroutines
   ```go
   // ❌ Leak potencial
   go func() {
       ch <- valor // Se ninguém receber, goroutine fica bloqueada
   }()
   ```

2. **Race conditions**: Sempre proteger dados compartilhados
   ```go
   // ❌ Race condition
   contador++
   
   // ✅ Protegido
   mu.Lock()
   contador++
   mu.Unlock()
   ```

3. **Deadlock**: Cuidado com dependências circulares
   ```go
   // ❌ Deadlock potencial
   ch1 := make(chan int)
   ch2 := make(chan int)
   go func() {
       ch1 <- <-ch2 // Espera ch2, mas ch2 espera ch1
   }()
   ch2 <- <-ch1
   ```

4. **Não usar time.Sleep em produção**: Use channels, WaitGroups ou context

---

## 12. Resumo

Nesta aula, você aprendeu:

- ✅ **Goroutines**: Threads leves para execução concorrente
- ✅ **Channels**: Comunicação entre goroutines
- ✅ **Select**: Multiplexar múltiplos channels
- ✅ **Buffered vs Unbuffered**: Síncrono vs assíncrono
- ✅ **Worker Pools**: Padrão para controlar concorrência
- ✅ **sync Package**: Primitivos de sincronização
- ✅ **Mutexes**: Proteção de dados compartilhados
- ✅ **WaitGroups**: Aguardar múltiplas goroutines

A concorrência é uma das características mais poderosas do Go. Com goroutines e channels, você pode escrever programas eficientes e escaláveis que aproveitam ao máximo os recursos do sistema.

Na próxima aula, vamos simplificar esses conceitos com analogias e exemplos do cotidiano para fixar ainda mais o aprendizado!

