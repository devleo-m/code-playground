# Aula 16 - Performance e Boas Práticas: Concorrência

Olá! Agora que você entende os conceitos de concorrência, é crucial aprender **quando e como** usá-los de forma eficiente. Nesta aula, vamos explorar aspectos de performance, boas práticas, e os erros comuns que você deve evitar.

---

## 🚀 Performance: Como Goroutines Funcionam

### Goroutines são Leves, mas Não Infinitas

**Fato importante:** Goroutines são muito mais leves que threads do sistema operacional, mas ainda têm limites.

**Números típicos:**
- **Thread do OS**: ~1-2 MB de stack por thread
- **Goroutine**: ~2-8 KB de stack inicial (cresce conforme necessário)
- **Capacidade**: Você pode ter milhões de goroutines, mas há limites práticos

**Limites práticos:**
```go
// ✅ OK: Milhares de goroutines
for i := 0; i < 10000; i++ {
    go processar()
}

// ⚠️ CUIDADO: Milhões podem ser demais
for i := 0; i < 1000000; i++ {
    go processar() // Pode esgotar recursos
}
```

**Quando usar muitas goroutines:**
- ✅ Operações I/O (leitura de arquivos, requisições HTTP)
- ✅ Operações que bloqueiam frequentemente
- ✅ Quando você precisa de alta concorrência

**Quando limitar goroutines:**
- ⚠️ Operações CPU-bound (cálculos pesados)
- ⚠️ Quando há recursos limitados (conexões de banco, file descriptors)
- ⚠️ Quando precisa de controle sobre uso de recursos

---

## ⚡ Performance: Channels vs Mutexes

### Channels Têm Overhead

**Importante:** Channels não são "grátis". Eles têm overhead de sincronização.

**Comparação de performance:**

```go
// Mutex (mais rápido para sincronização simples)
var mu sync.Mutex
mu.Lock()
contador++
mu.Unlock()

// Channel (mais lento, mas mais seguro para comunicação)
ch <- valor
valor := <-ch
```

**Benchmark típico:**
- **Mutex**: ~10-50 nanosegundos por operação
- **Channel**: ~100-500 nanosegundos por operação

**Quando usar cada um:**

✅ **Use Mutex quando:**
- Apenas precisa proteger acesso a dados compartilhados
- Não precisa de comunicação complexa
- Performance é crítica
- Operações são rápidas

✅ **Use Channel quando:**
- Precisa de comunicação entre goroutines
- Precisa de sincronização com contexto
- Quer seguir a filosofia "share by communicating"
- Operações são mais complexas

**Exemplo prático:**
```go
// ❌ EVITE: Channel para simples incremento
ch <- 1
contador += <-ch

// ✅ PREFIRA: Mutex para simples incremento
mu.Lock()
contador++
mu.Unlock()

// ✅ BOM: Channel para comunicação real
resultado := processar(dados)
ch <- resultado
```

---

## 🎯 Boas Práticas: Quando Usar Cada Ferramenta

### ✅ USE Goroutines Quando:

#### 1. Operações I/O (Leitura/Escrita)

```go
// ✅ EXCELENTE uso
func buscarURLs(urls []string) {
    for _, url := range urls {
        go func(u string) {
            resp, _ := http.Get(u)
            processar(resp)
        }(url)
    }
}
```

#### 2. Processamento Paralelo de Dados

```go
// ✅ BOM uso
func processarItens(items []Item) {
    var wg sync.WaitGroup
    for _, item := range items {
        wg.Add(1)
        go func(i Item) {
            defer wg.Done()
            processar(i)
        }(item)
    }
    wg.Wait()
}
```

#### 3. Background Tasks

```go
// ✅ BOM uso
go func() {
    for {
        time.Sleep(1 * time.Minute)
        fazerBackup()
    }
}()
```

### ❌ EVITE Goroutines Quando:

#### 1. Operações Simples e Rápidas

```go
// ❌ DESNECESSÁRIO
go func() {
    fmt.Println("Olá")
}()

// ✅ SIMPLES
fmt.Println("Olá")
```

#### 2. Sem Controle de Sincronização

```go
// ❌ PERIGO: Goroutine pode não terminar
go processar()
// main termina antes da goroutine!

// ✅ CORRETO: Usar WaitGroup ou channel
var wg sync.WaitGroup
wg.Add(1)
go func() {
    defer wg.Done()
    processar()
}()
wg.Wait()
```

---

## 🔒 Boas Práticas: Mutexes

### ✅ Sempre Use Defer para Unlock

```go
// ❌ PERIGO: Pode esquecer de fazer unlock
mu.Lock()
if condicao {
    return // Esqueceu de fazer unlock!
}
mu.Unlock()

// ✅ CORRETO: Defer garante unlock
mu.Lock()
defer mu.Unlock()
if condicao {
    return // Unlock acontece automaticamente
}
```

### ✅ Proteja Apenas o Necessário

```go
// ❌ RUIM: Lock muito longo
mu.Lock()
dados := buscarDados() // Operação lenta
processar(dados)
salvar(dados)
mu.Unlock()

// ✅ BOM: Lock apenas o necessário
dados := buscarDados() // Fora do lock
mu.Lock()
processar(dados)
salvar(dados)
mu.Unlock()
```

### ✅ Use RWMutex para Muitas Leituras

```go
// ✅ BOM: Múltiplos leitores simultâneos
type Cache struct {
    mu   sync.RWMutex
    data map[string]string
}

func (c *Cache) Get(key string) string {
    c.mu.RLock()         // Múltiplos leitores OK
    defer c.mu.RUnlock()
    return c.data[key]
}

func (c *Cache) Set(key, value string) {
    c.mu.Lock()          // Apenas 1 escritor
    defer c.mu.Unlock()
    c.data[key] = value
}
```

---

## 📞 Boas Práticas: Channels

### ✅ Sempre Feche Channels

```go
// ❌ PERIGO: Channel nunca fechado
func gerar() <-chan int {
    ch := make(chan int)
    go func() {
        for i := 0; i < 10; i++ {
            ch <- i
        }
        // Esqueceu de fechar!
    }()
    return ch
}

// ✅ CORRETO: Sempre fechar
func gerar() <-chan int {
    ch := make(chan int)
    go func() {
        defer close(ch) // Garante fechamento
        for i := 0; i < 10; i++ {
            ch <- i
        }
    }()
    return ch
}
```

### ✅ Verifique se Channel Está Fechado

```go
// ✅ BOM: Verificar se fechado
for {
    valor, ok := <-ch
    if !ok {
        fmt.Println("Channel fechado")
        break
    }
    processar(valor)
}

// ✅ MELHOR: Usar range (fecha automaticamente)
for valor := range ch {
    processar(valor)
}
```

### ✅ Use Buffered Channels Apropriadamente

```go
// ❌ EVITE: Buffer muito grande sem necessidade
ch := make(chan int, 1000000) // Desperdício de memória

// ✅ BOM: Buffer apropriado para o caso
ch := make(chan int, 10) // Suficiente para desacoplar

// ✅ BOM: Unbuffered quando precisa sincronização
ch := make(chan int) // Garante sincronização
```

---

## 👷 Boas Práticas: Worker Pools

### ✅ Número de Workers Baseado em Recursos

```go
// ❌ RUIM: Número arbitrário
const numWorkers = 100 // Por quê 100?

// ✅ BOM: Baseado em recursos
const numWorkers = runtime.NumCPU() // Um por núcleo para CPU-bound
const numWorkers = 10 // Fixo para I/O-bound (conexões limitadas)
```

**Regra geral:**
- **CPU-bound**: `runtime.NumCPU()` ou `runtime.NumCPU() * 2`
- **I/O-bound**: Baseado em recursos limitantes (conexões, file descriptors)
- **Misto**: Experimente e meça!

### ✅ Use Buffered Channels para Worker Pools

```go
// ✅ BOM: Buffer permite enviar jobs sem bloquear
jobs := make(chan Job, numWorkers*2) // Buffer suficiente

// Workers processam
for w := 0; w < numWorkers; w++ {
    go worker(jobs, results)
}

// Enviar jobs (não bloqueia se buffer tiver espaço)
for _, job := range jobList {
    jobs <- job
}
close(jobs)
```

---

## ⚠️ Armadilhas Comuns e Como Evitá-las

### 1. Goroutine Leak

**Problema:** Goroutine nunca termina, consumindo recursos.

```go
// ❌ LEAK: Goroutine bloqueada para sempre
func processar() {
    ch := make(chan int)
    go func() {
        ch <- 42 // Bloqueia se ninguém receber
    }()
    // Esqueceu de receber!
}

// ✅ CORRETO: Garantir que alguém recebe
func processar() {
    ch := make(chan int)
    go func() {
        ch <- 42
    }()
    valor := <-ch // Recebe o valor
}
```

**Solução:**
- Sempre garanta que channels sejam lidos
- Use context para cancelamento
- Use select com timeout

### 2. Race Condition

**Problema:** Múltiplas goroutines acessam dados sem proteção.

```go
// ❌ RACE CONDITION
var contador int

func incrementar() {
    contador++ // Não protegido!
}

// ✅ CORRETO: Proteger com Mutex
var (
    contador int
    mu       sync.Mutex
)

func incrementar() {
    mu.Lock()
    defer mu.Unlock()
    contador++
}
```

**Ferramenta:** Use `go run -race` para detectar race conditions!

### 3. Deadlock

**Problema:** Goroutines bloqueadas esperando umas pelas outras.

```go
// ❌ DEADLOCK: Dependência circular
ch1 := make(chan int)
ch2 := make(chan int)

go func() {
    ch1 <- <-ch2 // Espera ch2, mas ch2 espera ch1
}()

ch2 <- <-ch1 // Espera ch1, mas ch1 espera ch2
```

**Solução:**
- Evite dependências circulares
- Use buffered channels quando apropriado
- Use select com default para não bloquear

### 4. Usar time.Sleep em Produção

**Problema:** `time.Sleep` não é confiável para sincronização.

```go
// ❌ RUIM: Não confiável
go processar()
time.Sleep(1 * time.Second) // E se demorar mais?

// ✅ CORRETO: Usar WaitGroup ou channel
var wg sync.WaitGroup
wg.Add(1)
go func() {
    defer wg.Done()
    processar()
}()
wg.Wait() // Espera realmente terminar
```

---

## 🎯 Padrões Recomendados

### Padrão 1: Worker Pool para I/O

```go
func processarURLs(urls []string) {
    const numWorkers = 10
    jobs := make(chan string, len(urls))
    var wg sync.WaitGroup
    
    // Criar workers
    for i := 0; i < numWorkers; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            for url := range jobs {
                processarURL(url)
            }
        }()
    }
    
    // Enviar trabalho
    for _, url := range urls {
        jobs <- url
    }
    close(jobs)
    
    wg.Wait()
}
```

### Padrão 2: Pipeline com Channels

```go
func pipeline(dados []int) {
    // Etapa 1: Gerar
    nums := gerar(dados)
    
    // Etapa 2: Processar
    processados := processar(nums)
    
    // Etapa 3: Salvar
    salvar(processados)
}
```

### Padrão 3: Select com Timeout

```go
func comTimeout() {
    ch := make(chan string)
    
    go func() {
        time.Sleep(5 * time.Second)
        ch <- "resultado"
    }()
    
    select {
    case resultado := <-ch:
        fmt.Println(resultado)
    case <-time.After(2 * time.Second):
        fmt.Println("Timeout!")
    }
}
```

---

## 📊 Quando Usar Cada Ferramenta: Guia Rápido

| Situação | Ferramenta | Por quê? |
|----------|------------|----------|
| Comunicação entre goroutines | Channel | Filosofia do Go |
| Proteger dados compartilhados | Mutex | Simples e rápido |
| Muitas leituras, poucas escritas | RWMutex | Permite leituras paralelas |
| Esperar múltiplas goroutines | WaitGroup | Simples e eficiente |
| Multiplexar channels | Select | Essencial para coordenação |
| Controlar concorrência | Worker Pool | Limita uso de recursos |
| Timeout/Cancelamento | Select + Context | Padrão idiomático |
| Operações I/O | Goroutines | Aproveita tempo de espera |
| Operações CPU-bound | Worker Pool limitado | Evita sobrecarga |

---

## 🚨 Erros Críticos a Evitar

### ❌ NUNCA faça:

1. **Acessar dados compartilhados sem proteção**
   ```go
   // ❌ NUNCA
   contador++ // Sem mutex!
   ```

2. **Esquecer de fechar channels**
   ```go
   // ❌ NUNCA
   go func() {
       for i := 0; i < 10; i++ {
           ch <- i
       }
       // Esqueceu close(ch)!
   }()
   ```

3. **Passar WaitGroup por valor**
   ```go
   // ❌ NUNCA
   func worker(wg sync.WaitGroup) { ... }
   
   // ✅ SEMPRE
   func worker(wg *sync.WaitGroup) { ... }
   ```

4. **Usar mutex sem defer**
   ```go
   // ❌ NUNCA
   mu.Lock()
   // código...
   mu.Unlock() // Pode esquecer em caso de return/panic
   
   // ✅ SEMPRE
   mu.Lock()
   defer mu.Unlock()
   ```

5. **Criar goroutines sem controle**
   ```go
   // ❌ NUNCA
   for i := 0; i < 1000000; i++ {
       go processar() // Pode esgotar recursos
   }
   ```

---

## ✅ Checklist de Boas Práticas

Antes de finalizar seu código concorrente, verifique:

- [ ] Todas as goroutines têm uma forma de terminar?
- [ ] Channels são fechados quando não há mais dados?
- [ ] Dados compartilhados estão protegidos com Mutex?
- [ ] WaitGroups são passados por referência?
- [ ] Mutexes usam `defer Unlock()`?
- [ ] Worker pools têm número apropriado de workers?
- [ ] Há tratamento de erros nas goroutines?
- [ ] Não há race conditions (teste com `-race`)?
- [ ] Não há deadlocks potenciais?
- [ ] Timeouts estão implementados onde necessário?

---

## 🎓 Resumo Final

**Princípios fundamentais:**

1. **"Don't communicate by sharing memory; share memory by communicating"**
   - Prefira channels sobre mutexes quando possível
   - Use mutexes apenas quando necessário

2. **Goroutines são leves, mas não infinitas**
   - Use worker pools para controlar concorrência
   - Limite baseado em recursos disponíveis

3. **Sempre proteja dados compartilhados**
   - Use mutexes ou channels
   - Teste com `-race` flag

4. **Sempre limpe recursos**
   - Feche channels
   - Use WaitGroups para esperar goroutines
   - Use defer para garantir limpeza

5. **Meça e otimize**
   - Não otimize prematuramente
   - Use benchmarks para medir performance
   - Perfil seu código para encontrar gargalos

**Lembre-se:** Concorrência é poderosa, mas requer cuidado. Pratique, teste com `-race`, e aprenda com os erros!

Boa sorte com seus programas concorrentes! 🚀




