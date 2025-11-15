# Aula 19: Race Detection - Detectando Condições de Corrida

Olá! Bem-vindo(a) à aula 19! 

Agora que você já aprendeu sobre concorrência, goroutines, channels e padrões de programação concorrente, é crucial entender como **detectar e prevenir bugs** que podem surgir quando múltiplas goroutines acessam dados compartilhados de forma não sincronizada.

Nesta aula, vamos explorar o **Race Detector** do Go, uma ferramenta poderosa e nativa que ajuda a identificar condições de corrida (race conditions) em seus programas.

---

## O que é uma Condição de Corrida?

Antes de mergulharmos na ferramenta, vamos entender o problema que ela resolve.

Uma **condição de corrida (race condition)** ocorre quando duas ou mais goroutines acessam a mesma variável compartilhada **sem sincronização adequada**, e pelo menos uma dessas operações é uma **escrita**. O resultado do programa se torna **imprevisível** e depende da ordem de execução das goroutines, que pode variar a cada execução.

### Exemplo de Race Condition

```go
package main

import (
    "fmt"
    "sync"
)

var contador int

func incrementar() {
    contador++ // Operação não atômica: READ + MODIFY + WRITE
}

func main() {
    var wg sync.WaitGroup
    
    // Criar 1000 goroutines que incrementam o contador
    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            incrementar()
        }()
    }
    
    wg.Wait()
    fmt.Println("Contador final:", contador) // Pode ser qualquer valor <= 1000
}
```

**O problema:** A operação `contador++` não é atômica. Ela envolve três etapas:
1. **READ**: Ler o valor atual de `contador`
2. **MODIFY**: Incrementar o valor
3. **WRITE**: Escrever o novo valor de volta

Quando múltiplas goroutines executam isso simultaneamente, podem ler o mesmo valor, incrementar e escrever, resultando em perda de incrementos.

**Resultado esperado:** 1000  
**Resultado real (com race condition):** Pode ser 987, 923, 1000, ou qualquer outro valor ≤ 1000

---

## O que é o Race Detector?

O **Race Detector** é uma ferramenta **built-in** (integrada) do Go que detecta condições de corrida em tempo de execução. Ela foi desenvolvida para ajudar desenvolvedores a encontrar bugs de concorrência que são extremamente difíceis de detectar manualmente.

### Características Principais

1. **Ferramenta Nativa**: Não precisa instalar nada adicional, já vem com o Go
2. **Detecção em Tempo de Execução**: Analisa o comportamento do programa durante a execução
3. **Detecção de Acesso Não Sincronizado**: Identifica quando múltiplas goroutines acessam a mesma variável sem sincronização adequada
4. **Relatórios Detalhados**: Fornece informações precisas sobre onde e quando a race condition ocorreu

---

## Como Usar o Race Detector?

O Race Detector é ativado usando a flag `-race` durante a compilação, teste ou execução do programa.

### 1. Compilando com Race Detector

```bash
go build -race
```

Isso compila seu programa com instrumentação adicional para detectar race conditions.

### 2. Executando com Race Detector

```bash
go run -race main.go
```

Executa o programa com o race detector ativo.

### 3. Testando com Race Detector

```bash
go test -race
```

Executa os testes com detecção de race conditions. **Essa é a forma mais comum de uso!**

### 4. Instalando com Race Detector

```bash
go install -race
```

Instala o programa com race detector (útil para ferramentas de linha de comando).

---

## Exemplo Prático: Detectando uma Race Condition

Vamos ver o Race Detector em ação com nosso exemplo anterior:

### Código com Race Condition

```go
package main

import (
    "fmt"
    "sync"
)

var contador int

func incrementar() {
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
    fmt.Println("Contador final:", contador)
}
```

### Executando com Race Detector

```bash
go run -race main.go
```

### Saída do Race Detector

```
==================
WARNING: DATA RACE
Read at 0x0000012345678900 by goroutine 7:
  main.incrementar()
      /caminho/para/main.go:11 +0x3a
  main.main.func1()
      /caminho/para/main.go:20 +0x44

Previous write at 0x0000012345678900 by goroutine 6:
  main.incrementar()
      /caminho/para/main.go:11 +0x47
  main.main.func1()
      /caminho/para/main.go:20 +0x44

Goroutine 7 (running) created at:
  main.main()
      /caminho/para/main.go:18 +0x7a

Goroutine 6 (finished) created at:
  main.main()
      /caminho/para/main.go:18 +0x7a
==================
Contador final: 987
Found 1 data race(s)
exit status 66
```

### Entendendo o Relatório

O Race Detector fornece informações detalhadas:

1. **WARNING: DATA RACE**: Indica que uma race condition foi detectada
2. **Read at**: Mostra onde uma goroutine **leu** a variável
3. **Previous write at**: Mostra onde outra goroutine **escreveu** na mesma variável anteriormente
4. **Goroutine X created at**: Mostra onde cada goroutine foi criada
5. **Found N data race(s)**: Número total de race conditions encontradas

---

## Corrigindo a Race Condition

Agora que detectamos o problema, vamos corrigi-lo usando sincronização adequada:

### Solução 1: Usando Mutex

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
    mu.Lock()
    defer mu.Unlock()
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
    fmt.Println("Contador final:", contador) // Sempre 1000
}
```

### Solução 2: Usando atomic

```go
package main

import (
    "fmt"
    "sync"
    "sync/atomic"
)

var contador int64

func incrementar() {
    atomic.AddInt64(&contador, 1)
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
    fmt.Println("Contador final:", contador) // Sempre 1000
}
```

Agora, ao executar com `go run -race`, não haverá warnings!

---

## Quando o Race Detector Detecta Problemas?

O Race Detector detecta race conditions quando:

1. **Acesso não sincronizado**: Duas ou mais goroutines acessam a mesma variável
2. **Pelo menos uma escrita**: Pelo menos uma das operações é uma escrita
3. **Sem sincronização**: Não há uso adequado de mutex, channels, atomic, ou outras primitivas de sincronização

### Exemplos de Acesso que Causam Race Conditions

```go
// ❌ Race condition: escrita sem sincronização
var x int
go func() { x = 1 }()
go func() { x = 2 }()

// ❌ Race condition: leitura e escrita simultâneas
var y int
go func() { y++ }()
go func() { fmt.Println(y) }()

// ✅ Seguro: apenas leituras
var z int = 10
go func() { fmt.Println(z) }()
go func() { fmt.Println(z) }()

// ✅ Seguro: sincronizado com mutex
var w int
var mu sync.Mutex
go func() {
    mu.Lock()
    defer mu.Unlock()
    w = 1
}()
go func() {
    mu.Lock()
    defer mu.Unlock()
    w = 2
}()
```

---

## Limitações e Considerações

### 1. Overhead de Performance

O Race Detector adiciona **overhead significativo** ao programa:

- **Tempo de execução**: 2-10x mais lento
- **Uso de memória**: 5-10x mais memória
- **Tamanho do binário**: Aumenta consideravelmente

**Por isso, NUNCA use `-race` em produção!** Use apenas durante desenvolvimento e testes.

### 2. Detecção de Race Conditions Existentes

O Race Detector só detecta race conditions que **realmente ocorrem durante a execução**. Se uma race condition existe no código mas não é executada durante os testes, ela não será detectada.

**Solução**: Execute testes com diferentes cenários e cargas de trabalho.

### 3. Falsos Positivos

Em casos muito raros, o Race Detector pode reportar falsos positivos, especialmente em código que usa `unsafe` ou interage com C.

### 4. Não Previne, Apenas Detecta

O Race Detector **não previne** race conditions, apenas as **detecta**. Você ainda precisa corrigir o código manualmente.

---

## Boas Práticas com Race Detector

### 1. Use em Todos os Testes

```bash
# Adicione ao seu Makefile ou script de CI
test:
    go test -race ./...
```

### 2. Execute com Diferentes Cargas

```go
// Teste com diferentes números de goroutines
func TestRaceCondition(t *testing.T) {
    for numGoroutines := 1; numGoroutines <= 1000; numGoroutines *= 10 {
        t.Run(fmt.Sprintf("%d goroutines", numGoroutines), func(t *testing.T) {
            // Seu teste aqui
        })
    }
}
```

### 3. Integre no CI/CD

Configure seu pipeline de CI para executar testes com `-race`:

```yaml
# Exemplo para GitHub Actions
- name: Run tests with race detector
  run: go test -race ./...
```

### 4. Teste Código de Produção Regularmente

Execute periodicamente seus testes com race detector, mesmo em código que parece estável.

---

## Exemplo Completo: Sistema de Cache com Race Detection

Vamos ver um exemplo mais realista:

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

type Cache struct {
    data map[string]int
    mu   sync.RWMutex
}

func NewCache() *Cache {
    return &Cache{
        data: make(map[string]int),
    }
}

func (c *Cache) Get(key string) int {
    c.mu.RLock()
    defer c.mu.RUnlock()
    return c.data[key]
}

func (c *Cache) Set(key string, value int) {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.data[key] = value
}

func main() {
    cache := NewCache()
    var wg sync.WaitGroup
    
    // Múltiplas goroutines escrevendo
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            for j := 0; j < 100; j++ {
                cache.Set(fmt.Sprintf("key-%d", id), j)
            }
        }(i)
    }
    
    // Múltiplas goroutines lendo
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            for j := 0; j < 100; j++ {
                _ = cache.Get(fmt.Sprintf("key-%d", id))
            }
        }(i)
    }
    
    wg.Wait()
    fmt.Println("Cache operations completed")
}
```

Este código está **correto** e não terá race conditions porque usa `sync.RWMutex` para sincronização.

---

## Resumo

Nesta aula, você aprendeu:

1. ✅ **O que são race conditions**: Acesso não sincronizado a dados compartilhados
2. ✅ **O que é o Race Detector**: Ferramenta nativa do Go para detectar race conditions
3. ✅ **Como usar**: Flag `-race` em build, run, test ou install
4. ✅ **Como interpretar relatórios**: Entender os warnings do Race Detector
5. ✅ **Como corrigir**: Usar mutex, atomic, ou channels para sincronização
6. ✅ **Limitações**: Overhead de performance, não use em produção
7. ✅ **Boas práticas**: Use em testes, CI/CD, e com diferentes cargas

---

## Próximos Passos

Na próxima aula, vamos simplificar esses conceitos com analogias do dia a dia para fixar melhor o aprendizado!

Se tiver dúvidas, não hesite em perguntar. Race conditions são um dos bugs mais difíceis de debugar, mas com o Race Detector, você tem uma ferramenta poderosa ao seu lado!

