# Módulo 18: Concurrency Patterns em Go

## Aula 3: Exercícios Práticos e Reflexão

Olá! Agora é hora de colocar a mão na massa! Vamos praticar os padrões de concorrência que aprendemos através de exercícios práticos e questões de reflexão que vão fazer você pensar sobre o "porquê" e não apenas o "como".

---

## 📝 Exercício 1: Implementar Fan-In Básico

### Objetivo
Criar uma função que mescla três channels de strings em um único channel.

### Tarefa
Implemente a função `fanIn` que recebe três channels (`ch1`, `ch2`, `ch3`) e retorna um channel que contém todas as mensagens dos três channels mescladas.

### Esqueleto do Código

```go
package main

import (
    "fmt"
    "time"
)

func producer(name string, delay time.Duration, count int) <-chan string {
    ch := make(chan string)
    go func() {
        defer close(ch)
        for i := 1; i <= count; i++ {
            ch <- fmt.Sprintf("%s: mensagem %d", name, i)
            time.Sleep(delay)
        }
    }()
    return ch
}

// TODO: Implemente a função fanIn
func fanIn(ch1, ch2, ch3 <-chan string) <-chan string {
    // Sua implementação aqui
}

func main() {
    ch1 := producer("Producer 1", 200*time.Millisecond, 3)
    ch2 := producer("Producer 2", 300*time.Millisecond, 3)
    ch3 := producer("Producer 3", 400*time.Millisecond, 3)
    
    merged := fanIn(ch1, ch2, ch3)
    
    for msg := range merged {
        fmt.Println(msg)
    }
}
```

### Resultado Esperado
Você deve ver mensagens dos três producers mescladas, algo como:
```
Producer 1: mensagem 1
Producer 2: mensagem 1
Producer 3: mensagem 1
Producer 1: mensagem 2
...
```

### Dica
Use `select` para ler de múltiplos channels simultaneamente. Lembre-se de fechar o channel de saída quando todos os channels de entrada forem fechados.

---

## 📝 Exercício 2: Criar um Pipeline de Processamento

### Objetivo
Criar um pipeline que processa números: gera → filtra pares → eleva ao quadrado → formata como string.

### Tarefa
Implemente as funções do pipeline:
1. `generateNumbers(max int)`: Gera números de 1 até `max`
2. `filterEven(in <-chan int)`: Filtra apenas números pares
3. `square(in <-chan int)`: Eleva cada número ao quadrado
4. `format(in <-chan int)`: Formata como "Resultado: X"

### Esqueleto do Código

```go
package main

import (
    "fmt"
)

// TODO: Implemente generateNumbers
func generateNumbers(max int) <-chan int {
    // Sua implementação aqui
}

// TODO: Implemente filterEven
func filterEven(in <-chan int) <-chan int {
    // Sua implementação aqui
}

// TODO: Implemente square
func square(in <-chan int) <-chan int {
    // Sua implementação aqui
}

// TODO: Implemente format
func format(in <-chan int) <-chan string {
    // Sua implementação aqui
}

func main() {
    // Criar pipeline
    numbers := generateNumbers(10)
    evens := filterEven(numbers)
    squared := square(evens)
    formatted := format(squared)
    
    // Consumir resultado
    for result := range formatted {
        fmt.Println(result)
    }
}
```

### Resultado Esperado
```
Resultado: 4
Resultado: 16
Resultado: 36
Resultado: 64
Resultado: 100
```

### Dica
Cada função deve criar uma goroutine, ler do channel de entrada, processar e enviar para o channel de saída. Não esqueça de fechar os channels!

---

## 📝 Exercício 3: Worker Pool Simples

### Objetivo
Criar um worker pool que processa tarefas com um número fixo de workers.

### Tarefa
Implemente uma estrutura `WorkerPool` com:
- Método `NewWorkerPool(workers int)`: Cria um novo pool
- Método `Start()`: Inicia os workers
- Método `Submit(task Task)`: Submete uma tarefa para processamento
- Método `Stop()`: Para o pool e aguarda workers terminarem

### Esqueleto do Código

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

type Task struct {
    ID   int
    Data string
}

type WorkerPool struct {
    // TODO: Adicione os campos necessários
}

// TODO: Implemente NewWorkerPool
func NewWorkerPool(workers int) *WorkerPool {
    // Sua implementação aqui
}

// TODO: Implemente Start
func (wp *WorkerPool) Start() {
    // Sua implementação aqui
}

// TODO: Implemente worker (método auxiliar)
func (wp *WorkerPool) worker(id int) {
    // Sua implementação aqui
}

// TODO: Implemente Submit
func (wp *WorkerPool) Submit(task Task) {
    // Sua implementação aqui
}

// TODO: Implemente Stop
func (wp *WorkerPool) Stop() {
    // Sua implementação aqui
}

func main() {
    pool := NewWorkerPool(3)
    pool.Start()
    
    // Submeter tarefas
    for i := 1; i <= 9; i++ {
        pool.Submit(Task{
            ID:   i,
            Data: fmt.Sprintf("dados da tarefa %d", i),
        })
    }
    
    time.Sleep(5 * time.Second) // Dar tempo para processar
    pool.Stop()
    fmt.Println("Pool encerrado")
}
```

### Dica
Use um channel para a fila de tarefas, `sync.WaitGroup` para aguardar workers terminarem, e feche o channel para sinalizar que não há mais tarefas.

---

## 📝 Exercício 4: Pub-Sub Básico

### Objetivo
Implementar um sistema Pub-Sub simples onde publicadores enviam mensagens para tópicos e assinantes recebem mensagens dos tópicos que se inscreveram.

### Tarefa
Implemente a estrutura `PubSub` com:
- Método `Subscribe(topic string)`: Inscreve-se em um tópico e retorna um channel
- Método `Publish(topic string, message string)`: Publica uma mensagem em um tópico
- Método `Unsubscribe(topic string, ch <-chan Message)`: Remove uma assinatura (opcional, pode simplificar)

### Esqueleto do Código

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
    // TODO: Adicione os campos necessários
}

// TODO: Implemente NewPubSub
func NewPubSub() *PubSub {
    // Sua implementação aqui
}

// TODO: Implemente Subscribe
func (ps *PubSub) Subscribe(topic string) <-chan Message {
    // Sua implementação aqui
}

// TODO: Implemente Publish
func (ps *PubSub) Publish(topic string, content string) {
    // Sua implementação aqui
}

func main() {
    pubsub := NewPubSub()
    
    // Assinante 1: apenas tecnologia
    sub1 := pubsub.Subscribe("tecnologia")
    go func() {
        for msg := range sub1 {
            fmt.Printf("Assinante 1 recebeu [%s]: %s\n", msg.Topic, msg.Content)
        }
    }()
    
    // Assinante 2: tecnologia e esportes
    sub2Tech := pubsub.Subscribe("tecnologia")
    sub2Sports := pubsub.Subscribe("esportes")
    go func() {
        for {
            select {
            case msg := <-sub2Tech:
                fmt.Printf("Assinante 2 recebeu [%s]: %s\n", msg.Topic, msg.Content)
            case msg := <-sub2Sports:
                fmt.Printf("Assinante 2 recebeu [%s]: %s\n", msg.Topic, msg.Content)
            }
        }
    }()
    
    // Publicar mensagens
    time.Sleep(100 * time.Millisecond)
    pubsub.Publish("tecnologia", "Go 1.21 lançado!")
    time.Sleep(100 * time.Millisecond)
    pubsub.Publish("esportes", "Brasil vence a copa!")
    time.Sleep(100 * time.Millisecond)
    pubsub.Publish("tecnologia", "Concorrência é poderosa!")
    
    time.Sleep(1 * time.Second)
}
```

### Dica
Use um `map[string][]chan Message` para armazenar assinantes por tópico. Use `sync.RWMutex` para proteger o map de race conditions.

---

## 🤔 Questões de Reflexão

### Reflexão 1: Escolhendo o Padrão Certo

**Pergunta**: Você precisa processar 10.000 imagens. Cada imagem precisa ser: baixada, redimensionada, ter um filtro aplicado e ser enviada para um servidor. 

Qual padrão (ou combinação de padrões) você usaria e por quê? Justifique sua escolha considerando:
- Performance
- Uso de recursos
- Complexidade do código
- Facilidade de manutenção

**Escreva sua resposta aqui:**
```
[Seu espaço para resposta]
```

---

### Reflexão 2: Fan-In vs Fan-Out - Quando Usar Cada Um?

**Pergunta**: 
1. Em que situação você escolheria **Fan-In** ao invés de **Fan-Out**?
2. É possível usar ambos no mesmo programa? Dê um exemplo prático de quando isso faria sentido.
3. Qual é a principal diferença conceitual entre os dois padrões?

**Escreva sua resposta aqui:**
```
[Seu espaço para resposta]
```

---

### Reflexão 3: Pipeline e Separação de Responsabilidades

**Pergunta**: 
1. Por que usar Pipeline pode tornar o código mais fácil de manter e testar?
2. Dê um exemplo de um problema real (não relacionado a programação) que se beneficiaria de um pipeline. Explique como cada estágio do pipeline funcionaria.
3. Quais são as desvantagens de usar muitos estágios em um pipeline?

**Escreva sua resposta aqui:**
```
[Seu espaço para resposta]
```

---

### Reflexão 4: Worker Pool e Controle de Recursos

**Pergunta**: 
1. Por que é importante limitar o número de workers em um Worker Pool? O que pode acontecer se você criar goroutines demais?
2. Como você decidiria quantos workers usar em um Worker Pool? Quais fatores você consideraria?
3. Em um sistema que processa requisições HTTP, qual seria a diferença prática entre usar um Worker Pool e criar uma nova goroutine para cada requisição?

**Escreva sua resposta aqui:**
```
[Seu espaço para resposta]
```

---

### Reflexão 5: Pub-Sub e Desacoplamento

**Pergunta**: 
1. Por que Pub-Sub é útil em arquiteturas de microserviços? Dê um exemplo concreto.
2. Quais são as vantagens e desvantagens de usar Pub-Sub comparado a comunicação direta entre componentes?
3. Em um sistema de e-commerce, quais eventos você publicaria usando Pub-Sub? Quem seriam os assinantes de cada evento?

**Escreva sua resposta aqui:**
```
[Seu espaço para resposta]
```

---

### Reflexão 6: Combinando Padrões

**Pergunta**: 
Imagine um sistema de processamento de logs:
- Logs chegam de múltiplas fontes (servidores diferentes)
- Cada log precisa ser: validado, enriquecido com metadados, filtrado por nível de severidade, e armazenado
- O sistema precisa processar até 1000 logs por segundo
- Você quer notificar diferentes sistemas quando logs críticos são encontrados

Descreva como você combinaria os padrões aprendidos para resolver esse problema. Explique:
- Quais padrões usaria
- Como eles se conectariam
- Por que essa arquitetura faz sentido

**Escreva sua resposta aqui:**
```
[Seu espaço para resposta]
```

---

## ✅ Checklist de Aprendizado

Antes de considerar que você dominou este conteúdo, verifique se você consegue:

- [ ] Implementar Fan-In para mesclar múltiplos channels
- [ ] Implementar Fan-Out para distribuir trabalho entre workers
- [ ] Criar um Pipeline com múltiplos estágios
- [ ] Construir um Worker Pool funcional
- [ ] Implementar um sistema Pub-Sub básico
- [ ] Explicar quando usar cada padrão
- [ ] Combinar múltiplos padrões em uma solução
- [ ] Identificar race conditions potenciais e evitá-las
- [ ] Fechar channels corretamente para evitar vazamentos
- [ ] Usar context para cancelamento em padrões concorrentes

---

## 🎯 Próximos Passos

Depois de completar os exercícios e responder as questões de reflexão:

1. **Revise suas respostas**: Certifique-se de que você entendeu os conceitos, não apenas memorizou
2. **Experimente variações**: Tente modificar os exercícios (ex: adicionar timeout, usar context, etc.)
3. **Pense em aplicações reais**: Onde você poderia usar esses padrões em projetos reais?

Na próxima aula, vamos ver performance, boas práticas e o que fazer e não fazer ao usar esses padrões!

Boa sorte com os exercícios! 🚀



