# Aula 19 - Exercícios e Reflexão: Race Detection

Olá! Agora é hora de colocar em prática o que você aprendeu sobre Race Detection. Vamos fazer alguns exercícios práticos e depois refletir sobre os conceitos!

---

## 📝 Exercícios Práticos

### Exercício 1: Detectando a Race Condition

Analise o código abaixo e responda:

1. Este código tem uma race condition? Por quê?
2. Execute o código com `go run -race` e veja o que acontece
3. Corrija o código usando `sync.Mutex`

```go
package main

import (
    "fmt"
    "sync"
)

var total int

func adicionar(valor int) {
    total += valor
}

func main() {
    var wg sync.WaitGroup
    
    // Criar 100 goroutines que adicionam valores
    for i := 0; i < 100; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            for j := 0; j < 100; j++ {
                adicionar(1)
            }
        }(i)
    }
    
    wg.Wait()
    fmt.Printf("Total esperado: 10000\n")
    fmt.Printf("Total obtido: %d\n", total)
}
```

**Tarefas:**
- [ ] Identifique a race condition
- [ ] Execute com `-race` e copie o output
- [ ] Corrija usando `sync.Mutex`
- [ ] Execute novamente com `-race` e verifique que não há mais warnings

---

### Exercício 2: Cache Thread-Safe

Crie uma estrutura de cache simples que seja **thread-safe** (segura para uso concorrente).

**Requisitos:**
1. A estrutura deve ter métodos `Get(key string)` e `Set(key string, value int)`
2. Deve usar `sync.RWMutex` para permitir múltiplas leituras simultâneas
3. Deve passar no race detector (`go test -race`)
4. Crie testes que executam múltiplas goroutines lendo e escrevendo simultaneamente

**Template inicial:**

```go
package main

import (
    "fmt"
    "sync"
)

type Cache struct {
    // TODO: Adicione os campos necessários
}

func NewCache() *Cache {
    // TODO: Implemente
    return nil
}

func (c *Cache) Get(key string) (int, bool) {
    // TODO: Implemente com RLock
    return 0, false
}

func (c *Cache) Set(key string, value int) {
    // TODO: Implemente com Lock
}

func main() {
    cache := NewCache()
    var wg sync.WaitGroup
    
    // TODO: Crie goroutines que fazem Set e Get simultaneamente
    // Use o race detector para verificar que está correto
    
    wg.Wait()
    fmt.Println("Cache operations completed")
}
```

**Tarefas:**
- [ ] Implemente a estrutura Cache
- [ ] Implemente os métodos Get e Set com sincronização adequada
- [ ] Crie goroutines que testam leitura e escrita simultânea
- [ ] Execute com `go run -race` e verifique que não há race conditions

---

### Exercício 3: Contador Atômico vs Mutex

Compare duas implementações de um contador thread-safe:

**Implementação A: Usando Mutex**

```go
type ContadorMutex struct {
    valor int
    mu    sync.Mutex
}

func (c *ContadorMutex) Incrementar() {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.valor++
}

func (c *ContadorMutex) Valor() int {
    c.mu.Lock()
    defer c.mu.Unlock()
    return c.valor
}
```

**Implementação B: Usando atomic**

```go
import "sync/atomic"

type ContadorAtomic struct {
    valor int64
}

func (c *ContadorAtomic) Incrementar() {
    atomic.AddInt64(&c.valor, 1)
}

func (c *ContadorAtomic) Valor() int64 {
    return atomic.LoadInt64(&c.valor)
}
```

**Tarefas:**
1. Crie um programa que testa ambas as implementações
2. Execute com `go run -race` para ambas e verifique que não há race conditions
3. Crie um benchmark (`go test -bench=.`) para comparar a performance
4. Responda: Qual é mais rápida? Por quê?

---

### Exercício 4: Encontrando o Bug

O código abaixo tem uma race condition sutil. Encontre-a e corrija:

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

type Contador struct {
    valor int
    mu    sync.Mutex
}

func (c *Contador) Incrementar() {
    c.mu.Lock()
    c.valor++
    c.mu.Unlock()
}

func (c *Contador) Valor() int {
    // Ops! Esqueci de usar o mutex aqui!
    return c.valor
}

func main() {
    contador := &Contador{}
    var wg sync.WaitGroup
    
    // 10 goroutines incrementando
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            for j := 0; j < 1000; j++ {
                contador.Incrementar()
            }
        }()
    }
    
    // 1 goroutine lendo
    wg.Add(1)
    go func() {
        defer wg.Done()
        for {
            v := contador.Valor()
            fmt.Printf("Valor atual: %d\n", v)
            time.Sleep(10 * time.Millisecond)
            if v >= 10000 {
                break
            }
        }
    }()
    
    wg.Wait()
}
```

**Tarefas:**
- [ ] Execute com `go run -race` e identifique a race condition
- [ ] Corrija o método `Valor()` para usar o mutex adequadamente
- [ ] Execute novamente e verifique que está correto

**Dica:** Leitura sem sincronização enquanto há escrita também causa race condition!

---

## 🤔 Perguntas de Reflexão

### Reflexão 1: Por que Race Conditions são Perigosas?

Pense em um sistema real (por exemplo, um sistema bancário, um e-commerce, ou um jogo multiplayer).

**Perguntas:**
1. O que poderia acontecer se houver uma race condition em um sistema de transferência bancária?
2. Por que race conditions são consideradas um dos bugs mais difíceis de debugar?
3. Dê um exemplo de como uma race condition poderia causar perda de dados em um sistema real.

**Sua resposta deve ter pelo menos 3-4 parágrafos explicando:**
- O impacto em sistemas reais
- Por que são difíceis de detectar sem ferramentas
- Consequências possíveis (perda de dados, corrupção, comportamento inesperado)

---

### Reflexão 2: Quando Usar Race Detector?

O Race Detector tem overhead significativo (2-10x mais lento, 5-10x mais memória).

**Perguntas:**
1. Em quais situações você **sempre** deve usar o race detector?
2. Em quais situações você **nunca** deve usar o race detector?
3. Como você integraria o race detector no seu fluxo de trabalho de desenvolvimento?
4. Por que é importante executar testes com race detector mesmo em código que "parece" correto?

**Sua resposta deve incluir:**
- Uma estratégia clara de quando usar
- Exemplos de integração em CI/CD
- Explicação de por que código "correto" ainda pode ter race conditions

---

### Reflexão 3: Sincronização: Mutex vs Channels vs Atomic

Go oferece várias formas de sincronização: `sync.Mutex`, `channels`, e `sync/atomic`.

**Perguntas:**
1. Quando você escolheria usar `sync.Mutex` em vez de `channels`?
2. Quando você escolheria usar `sync/atomic` em vez de `sync.Mutex`?
3. Qual é a filosofia do Go sobre sincronização? (Dica: "Don't communicate by sharing memory; share memory by communicating")
4. Dê um exemplo prático de quando cada abordagem seria mais apropriada.

**Sua resposta deve:**
- Comparar as três abordagens
- Explicar as vantagens e desvantagens de cada uma
- Dar exemplos práticos de uso
- Explicar a filosofia do Go sobre comunicação entre goroutines

---

### Reflexão 4: O Custo da Segurança

Sincronização adiciona overhead (locks, unlocks, canais, etc.). Race Detector também adiciona overhead significativo.

**Perguntas:**
1. Por que é importante aceitar esse overhead durante desenvolvimento?
2. Como você balancearia performance e segurança em um sistema de produção?
3. Existe alguma situação onde você consideraria remover sincronização para ganhar performance? Quando isso seria aceitável?
4. Como você mediria o impacto da sincronização em seu código?

**Sua resposta deve:**
- Explicar o trade-off entre performance e segurança
- Dar exemplos de quando overhead é aceitável vs inaceitável
- Discutir estratégias de otimização sem comprometer segurança
- Explicar a importância de medir antes de otimizar

---

## ✅ Checklist de Entrega

Antes de enviar suas respostas, verifique:

- [ ] Todos os exercícios práticos foram implementados e testados
- [ ] Todos os códigos foram executados com `go run -race` ou `go test -race`
- [ ] Nenhum código tem race conditions (sem warnings do race detector)
- [ ] Todas as perguntas de reflexão foram respondidas com profundidade
- [ ] As respostas de reflexão têm pelo menos 3-4 parágrafos cada
- [ ] Você incluiu exemplos práticos nas respostas de reflexão

---

## 📚 Recursos Adicionais (Opcional)

Se quiser se aprofundar mais:

1. Leia a documentação oficial: `go doc -race`
2. Explore o código-fonte do race detector (se tiver curiosidade)
3. Experimente criar race conditions intencionais para ver como o detector as encontra
4. Teste diferentes padrões de sincronização e compare performance

---

Boa sorte com os exercícios! Lembre-se: a prática é essencial para dominar concorrência em Go. Race conditions são sutis e difíceis de encontrar, mas com o Race Detector, você tem uma ferramenta poderosa ao seu lado! 🚀

Envie suas respostas quando estiver pronto, e eu farei uma análise detalhada do seu desempenho!


