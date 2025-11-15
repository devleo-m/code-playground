# Módulo 17: Context Package em Go

## Aula 1: Context - Gerenciamento de Cancelamento e Valores de Requisição

Olá! Bem-vindo ao décimo sétimo módulo. Até agora você aprendeu sobre concorrência, goroutines, channels e como fazer múltiplas operações simultaneamente. Mas você já se perguntou: **"Como cancelar uma operação longa quando o usuário desiste? Como passar informações de uma requisição HTTP através de várias funções? Como garantir que operações de banco de dados não fiquem rodando para sempre?"**

O pacote `context` do Go foi criado exatamente para resolver esses problemas. Ele fornece uma forma padronizada de **cancelar operações**, **definir deadlines/timeouts** e **passar valores escopados a uma requisição** através de múltiplas camadas de chamadas de função.

Nesta aula, vamos mergulhar profundamente no pacote `context`: entender o que é um contexto, como criar e derivar contextos, como usar deadlines e timeouts, como cancelar operações, como passar valores através do contexto, e como usar context em aplicações reais como servidores HTTP e operações de banco de dados.

---

## 1. O que é Context?

### Definição

O **Context** (contexto) em Go é uma interface que carrega:
- **Deadlines**: Tempo limite absoluto para uma operação
- **Cancellation signals**: Sinais de cancelamento que podem ser propagados
- **Request-scoped values**: Valores específicos de uma requisição que precisam ser passados através de múltiplas funções

### Por que Context é Importante?

1. **Cancelamento de Operações**: Permite cancelar operações longas quando não são mais necessárias
2. **Timeouts**: Evita que operações fiquem rodando indefinidamente
3. **Propagação de Valores**: Passa informações (como user ID, request ID) através de múltiplas camadas sem poluir assinaturas de função
4. **Boas Práticas**: É o padrão da comunidade Go para gerenciar o ciclo de vida de operações

### A Interface Context

```go
type Context interface {
    Deadline() (deadline time.Time, ok bool)
    Done() <-chan struct{}
    Err() error
    Value(key interface{}) interface{}
}
```

**Métodos:**
- `Deadline()`: Retorna o tempo limite e se há um deadline definido
- `Done()`: Retorna um channel que é fechado quando o contexto é cancelado
- `Err()`: Retorna o erro que causou o cancelamento (nil se não foi cancelado)
- `Value(key)`: Retorna o valor associado à chave, ou nil se não existe

---

## 2. Criando Contextos

### Context.Background()

`context.Background()` retorna um contexto vazio que nunca é cancelado e não tem deadline ou valores. É usado como **contexto raiz** para toda a aplicação.

```go
ctx := context.Background()
```

**Quando usar:**
- Como ponto de partida para criar outros contextos
- Na função `main()`
- Em testes
- Como contexto raiz quando não há contexto pai disponível

### Context.TODO()

`context.TODO()` retorna um contexto vazio similar ao `Background()`, mas é usado quando você **não tem certeza** qual contexto usar ou quando o contexto ainda não está disponível.

```go
ctx := context.TODO()
```

**Quando usar:**
- Durante desenvolvimento quando você ainda não sabe qual contexto usar
- Quando refatorando código que ainda não aceita context
- Como placeholder temporário

**⚠️ Importante**: Em código de produção, prefira sempre `context.Background()` ou um contexto derivado.

---

## 3. Context com Timeout

### Context.WithTimeout()

`context.WithTimeout()` cria um contexto que é **automaticamente cancelado** após um período de tempo especificado.

```go
ctx, cancel := context.WithTimeout(parent context.Context, timeout time.Duration)
```

**Parâmetros:**
- `parent`: Contexto pai (geralmente `context.Background()`)
- `timeout`: Duração do timeout (ex: `5 * time.Second`)

**Retorna:**
- `ctx`: Novo contexto com timeout
- `cancel`: Função para cancelar manualmente (sempre chame com `defer cancel()`)

### Exemplo: Timeout em Operação Longa

```go
package main

import (
    "context"
    "fmt"
    "time"
)

func operacaoLonga(ctx context.Context) error {
    // Simular trabalho
    for i := 0; i < 10; i++ {
        // Verificar se contexto foi cancelado
        select {
        case <-ctx.Done():
            return ctx.Err() // Retornar erro de cancelamento
        default:
            fmt.Printf("Processando item %d...\n", i)
            time.Sleep(500 * time.Millisecond)
        }
    }
    return nil
}

func main() {
    // Criar contexto com timeout de 2 segundos
    ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
    defer cancel() // Sempre cancelar para liberar recursos

    err := operacaoLonga(ctx)
    if err != nil {
        fmt.Printf("Operação cancelada: %v\n", err)
    } else {
        fmt.Println("Operação concluída com sucesso")
    }
}
```

**Saída:**
```
Processando item 0...
Processando item 1...
Processando item 2...
Processando item 3...
Operação cancelada: context deadline exceeded
```

### Quando Usar Timeout?

- ✅ **Requisições HTTP**: Limitar tempo de resposta
- ✅ **Operações de Banco de Dados**: Evitar queries infinitas
- ✅ **Chamadas de API**: Garantir que não esperem para sempre
- ✅ **Operações de I/O**: Ler/escrever arquivos com limite de tempo

---

## 4. Context com Deadline

### Context.WithDeadline()

`context.WithDeadline()` cria um contexto que é cancelado em um **tempo absoluto específico** (não uma duração).

```go
ctx, cancel := context.WithDeadline(parent context.Context, deadline time.Time)
```

**Parâmetros:**
- `parent`: Contexto pai
- `deadline`: Tempo absoluto quando o contexto deve ser cancelado

### Exemplo: Deadline Absoluto

```go
package main

import (
    "context"
    "fmt"
    "time"
)

func main() {
    // Criar deadline para 5 segundos a partir de agora
    deadline := time.Now().Add(5 * time.Second)
    ctx, cancel := context.WithDeadline(context.Background(), deadline)
    defer cancel()

    // Verificar deadline
    d, ok := ctx.Deadline()
    if ok {
        fmt.Printf("Deadline: %v\n", d)
        fmt.Printf("Tempo restante: %v\n", time.Until(d))
    }

    // Simular trabalho
    for i := 0; i < 10; i++ {
        select {
        case <-ctx.Done():
            fmt.Printf("Deadline atingido! Erro: %v\n", ctx.Err())
            return
        case <-time.After(1 * time.Second):
            fmt.Printf("Trabalho %d concluído\n", i+1)
        }
    }
}
```

### Timeout vs Deadline

- **Timeout**: "Cancelar após X segundos/minutos"
- **Deadline**: "Cancelar em um momento específico (ex: 15:30:00)"

**Use Deadline quando:**
- Você tem um horário específico (ex: "antes das 15h")
- Você precisa sincronizar com eventos externos
- Você quer mais controle sobre o tempo exato

**Use Timeout quando:**
- Você quer "cancelar após X tempo"
- É mais simples e direto para a maioria dos casos

---

## 5. Context com Cancelamento Manual

### Context.WithCancel()

`context.WithCancel()` cria um contexto que pode ser **cancelado manualmente** chamando a função `cancel()`.

```go
ctx, cancel := context.WithCancel(parent context.Context)
```

### Exemplo: Cancelamento Manual

```go
package main

import (
    "context"
    "fmt"
    "time"
)

func worker(ctx context.Context, id int) {
    for {
        select {
        case <-ctx.Done():
            fmt.Printf("Worker %d: Cancelado - %v\n", id, ctx.Err())
            return
        default:
            fmt.Printf("Worker %d: Trabalhando...\n", id)
            time.Sleep(500 * time.Millisecond)
        }
    }
}

func main() {
    ctx, cancel := context.WithCancel(context.Background())

    // Iniciar múltiplos workers
    for i := 1; i <= 3; i++ {
        go worker(ctx, i)
    }

    // Trabalhar por 2 segundos
    time.Sleep(2 * time.Second)

    // Cancelar todos os workers
    fmt.Println("Cancelando todos os workers...")
    cancel()

    // Dar tempo para workers terminarem
    time.Sleep(500 * time.Millisecond)
}
```

**Saída:**
```
Worker 1: Trabalhando...
Worker 2: Trabalhando...
Worker 3: Trabalhando...
Worker 1: Trabalhando...
Worker 2: Trabalhando...
Worker 3: Trabalhando...
Worker 1: Trabalhando...
Worker 2: Trabalhando...
Worker 3: Trabalhando...
Worker 1: Trabalhando...
Cancelando todos os workers...
Worker 1: Cancelado - context canceled
Worker 2: Cancelado - context canceled
Worker 3: Cancelado - context canceled
```

### Quando Usar Cancelamento Manual?

- ✅ **Shutdown graceful**: Parar servidores de forma controlada
- ✅ **Cancelar operações**: Quando usuário cancela uma ação
- ✅ **Limpar recursos**: Garantir que goroutines sejam finalizadas
- ✅ **Coordenar múltiplas goroutines**: Cancelar todas de uma vez

---

## 6. Context com Valores (Request-Scoped Values)

### Context.WithValue()

`context.WithValue()` adiciona um par chave-valor ao contexto. Esses valores são **específicos da requisição** e devem ser usados apenas para dados que precisam ser passados através de múltiplas camadas.

```go
ctx := context.WithValue(parent context.Context, key, val interface{})
```

**⚠️ Importante**: 
- Use apenas para dados de requisição (user ID, request ID, trace ID)
- **NÃO** use para passar parâmetros opcionais de função
- **NÃO** use para passar dependências (use injeção de dependência)
- Chaves devem ser tipos específicos (não strings genéricas)

### Exemplo: Passando Valores através do Context

```go
package main

import (
    "context"
    "fmt"
)

// Tipos específicos para chaves (boa prática)
type userIDKey struct{}
type requestIDKey struct{}

func processarRequisicao(ctx context.Context) {
    // Recuperar valores do contexto
    userID := ctx.Value(userIDKey{}).(string)
    requestID := ctx.Value(requestIDKey{}).(string)

    fmt.Printf("Processando requisição %s para usuário %s\n", requestID, userID)
}

func autenticar(ctx context.Context) context.Context {
    // Simular autenticação e adicionar user ID
    userID := "12345"
    return context.WithValue(ctx, userIDKey{}, userID)
}

func adicionarRequestID(ctx context.Context) context.Context {
    // Adicionar request ID
    requestID := "req-abc-xyz"
    return context.WithValue(ctx, requestIDKey{}, requestID)
}

func main() {
    ctx := context.Background()
    ctx = autenticar(ctx)
    ctx = adicionarRequestID(ctx)

    processarRequisicao(ctx)
}
```

### Boas Práticas para Valores

**✅ FAÇA:**
```go
// Usar tipos específicos para chaves
type userIDKey struct{}

ctx := context.WithValue(ctx, userIDKey{}, "12345")
userID := ctx.Value(userIDKey{}).(string)
```

**❌ NÃO FAÇA:**
```go
// Usar strings como chaves (pode causar colisões)
ctx := context.WithValue(ctx, "userID", "12345")

// Passar dependências através do context
ctx := context.WithValue(ctx, "database", db)

// Passar parâmetros opcionais
ctx := context.WithValue(ctx, "timeout", 5*time.Second)
```

### Quando Usar Valores no Context?

- ✅ **Request ID**: Rastrear requisições através de logs
- ✅ **User ID**: Identificar usuário autenticado
- ✅ **Trace ID**: Rastreamento distribuído
- ✅ **Metadata de requisição**: Informações que precisam ser passadas através de múltiplas camadas

---

## 7. Verificando Status do Context

### Verificando se foi Cancelado

```go
// Verificar se contexto foi cancelado
if ctx.Err() != nil {
    // Context foi cancelado
    return ctx.Err()
}

// Ou usando select
select {
case <-ctx.Done():
    return ctx.Err()
default:
    // Context ainda está ativo
}
```

### Verificando Deadline

```go
deadline, ok := ctx.Deadline()
if ok {
    fmt.Printf("Deadline: %v\n", deadline)
    fmt.Printf("Tempo restante: %v\n", time.Until(deadline))
} else {
    fmt.Println("Sem deadline definido")
}
```

### Erros do Context

O método `Err()` retorna:
- `nil`: Context não foi cancelado
- `context.Canceled`: Context foi cancelado manualmente
- `context.DeadlineExceeded`: Deadline/timeout foi atingido

```go
switch ctx.Err() {
case nil:
    fmt.Println("Context ativo")
case context.Canceled:
    fmt.Println("Context cancelado manualmente")
case context.DeadlineExceeded:
    fmt.Println("Deadline excedido")
}
```

---

## 8. Context em Servidores HTTP

### Context Automático em HTTP Requests

O pacote `net/http` do Go automaticamente cria um contexto para cada requisição HTTP e o cancela quando a conexão é fechada.

```go
package main

import (
    "context"
    "fmt"
    "net/http"
    "time"
)

func handler(w http.ResponseWriter, r *http.Request) {
    // Context vem automaticamente do request
    ctx := r.Context()

    // Verificar se request foi cancelado
    select {
    case <-ctx.Done():
        fmt.Printf("Request cancelado: %v\n", ctx.Err())
        http.Error(w, "Request cancelado", http.StatusRequestTimeout)
        return
    default:
    }

    // Processar request
    fmt.Fprintf(w, "Processando requisição...\n")

    // Simular trabalho com timeout
    ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
    defer cancel()

    // Fazer operação longa
    select {
    case <-time.After(10 * time.Second):
        fmt.Fprintf(w, "Operação concluída\n")
    case <-ctx.Done():
        fmt.Fprintf(w, "Timeout: %v\n", ctx.Err())
    }
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
```

### Timeout em Handlers HTTP

```go
func handlerComTimeout(w http.ResponseWriter, r *http.Request) {
    ctx := r.Context()

    // Adicionar timeout de 3 segundos
    ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
    defer cancel()

    // Operação que pode demorar
    resultado, err := operacaoLonga(ctx)
    if err != nil {
        http.Error(w, err.Error(), http.StatusRequestTimeout)
        return
    }

    fmt.Fprintf(w, "Resultado: %s\n", resultado)
}
```

---

## 9. Context em Operações de Banco de Dados

### Context em Queries de Banco

A maioria dos drivers de banco de dados em Go aceita context como primeiro parâmetro:

```go
package main

import (
    "context"
    "database/sql"
    "fmt"
    "time"
)

func consultarUsuarios(ctx context.Context, db *sql.DB) ([]User, error) {
    // Adicionar timeout de 5 segundos
    ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
    defer cancel()

    // Query com context
    rows, err := db.QueryContext(ctx, "SELECT id, name FROM users")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var users []User
    for rows.Next() {
        var user User
        if err := rows.Scan(&user.ID, &user.Name); err != nil {
            return nil, err
        }
        users = append(users, user)
    }

    return users, rows.Err()
}

func main() {
    // db seria inicializado anteriormente
    ctx := context.Background()
    users, err := consultarUsuarios(ctx, db)
    if err != nil {
        if err == context.DeadlineExceeded {
            fmt.Println("Query demorou muito tempo")
        }
        return
    }
    fmt.Printf("Encontrados %d usuários\n", len(users))
}
```

### Context em Transações

```go
func executarTransacao(ctx context.Context, db *sql.DB) error {
    ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()

    tx, err := db.BeginTx(ctx, nil)
    if err != nil {
        return err
    }
    defer tx.Rollback()

    // Executar operações na transação
    _, err = tx.ExecContext(ctx, "INSERT INTO users ...")
    if err != nil {
        return err
    }

    return tx.Commit()
}
```

---

## 10. Context Aninhado (Context Chain)

### Derivando Contextos

Você pode derivar múltiplos contextos de um contexto pai. Quando o contexto pai é cancelado, **todos os contextos derivados também são cancelados**.

```go
package main

import (
    "context"
    "fmt"
    "time"
)

func main() {
    // Context raiz
    ctx := context.Background()

    // Adicionar timeout de 10 segundos
    ctx, cancel1 := context.WithTimeout(ctx, 10*time.Second)
    defer cancel1()

    // Adicionar valores
    ctx = context.WithValue(ctx, "userID", "12345")

    // Adicionar timeout mais restritivo (2 segundos)
    // Este timeout sobrescreve o anterior se for menor
    ctx, cancel2 := context.WithTimeout(ctx, 2*time.Second)
    defer cancel2()

    // Agora temos:
    // - Timeout de 2 segundos (o mais restritivo)
    // - userID no contexto
    // - Se cancel1() for chamado, tudo é cancelado
    // - Se cancel2() for chamado, apenas este contexto é cancelado

    select {
    case <-time.After(3 * time.Second):
        fmt.Println("Operação concluída")
    case <-ctx.Done():
        fmt.Printf("Cancelado: %v\n", ctx.Err())
    }
}
```

### Hierarquia de Cancelamento

```
Background (raiz)
  └── WithTimeout(10s)
       └── WithValue(userID)
            └── WithTimeout(2s)  ← Mais restritivo
```

Quando qualquer contexto na cadeia é cancelado, todos os contextos filhos também são cancelados.

---

## 11. Padrões de Uso do Context

### Padrão: Context como Primeiro Parâmetro

**Sempre** passe context como o **primeiro parâmetro** de funções que fazem operações que podem ser canceladas:

```go
// ✅ CORRETO
func processarDados(ctx context.Context, dados []string) error {
    // ...
}

// ❌ ERRADO
func processarDados(dados []string, ctx context.Context) error {
    // ...
}
```

### Padrão: Verificar Cancelamento em Loops

```go
func processarItems(ctx context.Context, items []Item) error {
    for _, item := range items {
        // Verificar cancelamento antes de cada iteração
        select {
        case <-ctx.Done():
            return ctx.Err()
        default:
        }

        // Processar item
        if err := processarItem(ctx, item); err != nil {
            return err
        }
    }
    return nil
}
```

### Padrão: Timeout em Operações Externas

```go
func chamarAPI(ctx context.Context, url string) (*Response, error) {
    // Adicionar timeout específico para esta operação
    ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
    defer cancel()

    req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
    if err != nil {
        return nil, err
    }

    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    // Processar resposta...
    return &Response{}, nil
}
```

### Padrão: Graceful Shutdown

```go
func main() {
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()

    // Iniciar servidor
    server := &http.Server{Addr: ":8080"}

    // Goroutine para shutdown graceful
    go func() {
        <-ctx.Done()
        fmt.Println("Encerrando servidor...")
        server.Shutdown(context.Background())
    }()

    // Aguardar sinal de interrupção
    sigChan := make(chan os.Signal, 1)
    signal.Notify(sigChan, os.Interrupt)

    go func() {
        <-sigChan
        fmt.Println("Recebido sinal de interrupção")
        cancel() // Cancela contexto, que encerra servidor
    }()

    server.ListenAndServe()
}
```

---

## 12. Erros Comuns e Como Evitá-los

### Erro 1: Não Cancelar Context

```go
// ❌ ERRADO
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// Esqueceu de chamar cancel()!

// ✅ CORRETO
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel() // Sempre cancelar
```

**Por quê?** Se você não cancelar, recursos podem vazar. Sempre use `defer cancel()`.

### Erro 2: Passar Nil Context

```go
// ❌ ERRADO
func processar(ctx context.Context) {
    if ctx == nil {
        ctx = context.Background()
    }
    // ...
}

// ✅ CORRETO
// Sempre aceite context não-nil
func processar(ctx context.Context) {
    // ctx nunca deve ser nil
}
```

**Por quê?** Context nunca deve ser nil. Se uma função não precisa de context, não o inclua na assinatura.

### Erro 3: Usar Context para Parâmetros Opcionais

```go
// ❌ ERRADO
ctx := context.WithValue(context.Background(), "timeout", 5*time.Second)

// ✅ CORRETO
func processar(ctx context.Context, timeout time.Duration) {
    ctx, cancel := context.WithTimeout(ctx, timeout)
    defer cancel()
    // ...
}
```

**Por quê?** Context é para cancelamento e valores de requisição, não para parâmetros de função.

### Erro 4: Não Verificar Cancelamento

```go
// ❌ ERRADO
func operacaoLonga(ctx context.Context) {
    for i := 0; i < 1000000; i++ {
        // Não verifica se foi cancelado
        processar(i)
    }
}

// ✅ CORRETO
func operacaoLonga(ctx context.Context) error {
    for i := 0; i < 1000000; i++ {
        select {
        case <-ctx.Done():
            return ctx.Err()
        default:
        }
        processar(i)
    }
    return nil
}
```

---

## 13. Resumo dos Conceitos

### Tipos de Context

1. **Background**: Contexto raiz, nunca cancelado
2. **TODO**: Placeholder quando não sabe qual usar
3. **WithTimeout**: Cancelado após duração específica
4. **WithDeadline**: Cancelado em tempo absoluto
5. **WithCancel**: Cancelado manualmente
6. **WithValue**: Adiciona valores ao contexto

### Quando Usar Cada Um

- **Background/TODO**: Ponto de partida
- **WithTimeout**: Maioria dos casos (requisições HTTP, queries)
- **WithDeadline**: Quando precisa de tempo absoluto
- **WithCancel**: Shutdown graceful, cancelamento manual
- **WithValue**: Request ID, User ID, Trace ID

### Boas Práticas

1. ✅ Sempre passe context como primeiro parâmetro
2. ✅ Sempre use `defer cancel()` após criar context com cancel
3. ✅ Verifique `ctx.Done()` em loops longos
4. ✅ Use tipos específicos para chaves de valores
5. ✅ Não use context para parâmetros opcionais
6. ✅ Não armazene context em structs (passe como parâmetro)

---

E assim terminamos nossa primeira aula sobre Context! Você agora entende:
- O que é context e por que é importante
- Como criar diferentes tipos de contextos
- Como usar timeouts e deadlines
- Como cancelar operações
- Como passar valores através do contexto
- Como usar context em aplicações reais

Na próxima aula, vamos simplificar esses conceitos com analogias do dia a dia para fixar ainda mais o aprendizado!

Sinta-se à vontade para reler este material e experimentar com os exemplos. Se tiver qualquer dúvida, pode perguntar!

