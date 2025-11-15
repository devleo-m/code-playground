# Aula 17 - Performance e Boas Práticas: Context Package

Olá! Agora que você entende os conceitos de Context, é crucial aprender **quando e como** usá-los de forma eficiente e correta. Nesta aula, vamos explorar aspectos de performance, boas práticas, anti-padrões e os erros comuns que você deve evitar.

---

## 🚀 Performance: Overhead do Context

### Context é Leve, mas Não é Grátis

**Fato importante:** Context tem um custo de performance, mas é mínimo e geralmente aceitável.

**Custos típicos:**
- **Criar context**: ~10-50 nanosegundos
- **Verificar `ctx.Done()`**: ~1-5 nanosegundos
- **Cancelar context**: ~10-100 nanosegundos
- **Acessar valores**: ~5-20 nanosegundos

**Quando o overhead importa:**
```go
// ⚠️ CUIDADO: Verificar em loop muito apertado
for i := 0; i < 1000000000; i++ {
    select {
    case <-ctx.Done(): // Overhead acumula
        return
    default:
        // Trabalho muito rápido
    }
}

// ✅ MELHOR: Verificar periodicamente
for i := 0; i < 1000000000; i++ {
    if i%10000 == 0 { // Verificar a cada 10k iterações
        select {
        case <-ctx.Done():
            return
        default:
        }
    }
    // Trabalho rápido
}
```

**Regra geral:**
- ✅ **Sempre** use context em operações I/O (HTTP, DB, arquivos)
- ✅ **Sempre** use context em loops longos
- ⚠️ **Considere** verificar periodicamente em loops muito apertados
- ❌ **Nunca** remova context por performance sem medir primeiro

---

## ⚡ Performance: Verificando Cancelamento

### Verificação Eficiente

**Padrão recomendado:**
```go
// ✅ EFICIENTE: Verificar antes de operações caras
func processar(ctx context.Context, dados []Item) error {
    for _, item := range dados {
        // Verificar ANTES de processar
        select {
        case <-ctx.Done():
            return ctx.Err()
        default:
        }
        
        // Operação que pode demorar
        if err := processarItem(ctx, item); err != nil {
            return err
        }
    }
    return nil
}
```

**Evite verificar desnecessariamente:**
```go
// ❌ INEFICIENTE: Verificar em operações muito rápidas
func incrementar(ctx context.Context, n int) int {
    for i := 0; i < n; i++ {
        select { // Desnecessário para operação tão rápida
        case <-ctx.Done():
            return 0
        default:
        }
        contador++
    }
    return contador
}

// ✅ MELHOR: Verificar apenas se operação pode demorar
func incrementar(ctx context.Context, n int) int {
    // Operação rápida, não precisa verificar
    for i := 0; i < n; i++ {
        contador++
    }
    return contador
}
```

---

## 🎯 Boas Práticas: Quando Usar Context

### ✅ USE Context Quando:

#### 1. Operações I/O (HTTP, Banco de Dados, Arquivos)

```go
// ✅ EXCELENTE uso
func buscarUsuario(ctx context.Context, userID string) (*User, error) {
    ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
    defer cancel()
    
    resp, err := http.Get(fmt.Sprintf("https://api.com/users/%s", userID))
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()
    
    // Processar resposta...
    return &User{}, nil
}
```

#### 2. Loops Longos ou Processamento em Lote

```go
// ✅ BOM uso
func processarLote(ctx context.Context, items []Item) error {
    for _, item := range items {
        select {
        case <-ctx.Done():
            return ctx.Err()
        default:
        }
        
        if err := processarItem(ctx, item); err != nil {
            return err
        }
    }
    return nil
}
```

#### 3. Operações que Podem Ser Canceladas pelo Usuário

```go
// ✅ BOM uso
func downloadArquivo(ctx context.Context, url string) error {
    req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
    if err != nil {
        return err
    }
    
    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        return err
    }
    defer resp.Body.Close()
    
    // Download...
    return nil
}
```

#### 4. Passando Valores de Requisição

```go
// ✅ BOM uso
type userIDKey struct{}

func autenticar(ctx context.Context, token string) (context.Context, error) {
    userID := validarToken(token)
    return context.WithValue(ctx, userIDKey{}, userID), nil
}

func processarRequisicao(ctx context.Context) {
    userID := ctx.Value(userIDKey{}).(string)
    // Usar userID...
}
```

### ❌ EVITE Context Quando:

#### 1. Operações Puramente Locais e Rápidas

```go
// ❌ DESNECESSÁRIO
func somar(ctx context.Context, a, b int) int {
    return a + b // Operação muito rápida, não precisa de context
}

// ✅ MELHOR
func somar(a, b int) int {
    return a + b
}
```

#### 2. Para Passar Parâmetros Opcionais

```go
// ❌ ERRADO
func buscar(ctx context.Context, query string) []Result {
    incluirMetadata := ctx.Value("incluirMetadata").(bool) // ERRADO!
    // ...
}

// ✅ CORRETO
func buscar(ctx context.Context, query string, incluirMetadata bool) []Result {
    // incluirMetadata como parâmetro normal
    // ...
}
```

#### 3. Para Passar Dependências

```go
// ❌ ERRADO
type dbKey struct{}
ctx := context.WithValue(ctx, dbKey{}, db)

func processar(ctx context.Context) {
    db := ctx.Value(dbKey{}).(*sql.DB) // ERRADO!
    // ...
}

// ✅ CORRETO
func processar(ctx context.Context, db *sql.DB) {
    // db como parâmetro normal
    // ...
}
```

#### 4. Para Configurações Globais

```go
// ❌ ERRADO
ctx := context.WithValue(ctx, "timeout", 5*time.Second)
ctx = context.WithValue(ctx, "retries", 3)

// ✅ CORRETO
type Config struct {
    Timeout time.Duration
    Retries int
}

func processar(ctx context.Context, config Config) {
    ctx, cancel := context.WithTimeout(ctx, config.Timeout)
    defer cancel()
    // ...
}
```

---

## 🛡️ Boas Práticas: Padrões de Uso

### Padrão 1: Sempre Context como Primeiro Parâmetro

```go
// ✅ CORRETO
func processar(ctx context.Context, dados []string) error
func buscar(ctx context.Context, id string) (*Item, error)
func salvar(ctx context.Context, item *Item) error

// ❌ ERRADO
func processar(dados []string, ctx context.Context) error
func buscar(id string, ctx context.Context) (*Item, error)
```

**Por quê?**
- Convenção universal em Go
- Facilita leitura e manutenção
- Permite composição de funções

### Padrão 2: Sempre Usar defer cancel()

```go
// ✅ CORRETO
func operacao(ctx context.Context) error {
    ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
    defer cancel() // SEMPRE cancelar
    
    // Operação...
    return nil
}

// ❌ ERRADO
func operacao(ctx context.Context) error {
    ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
    // Esqueceu de cancelar! Vazamento de recursos!
    
    // Operação...
    return nil
}
```

**Por quê?**
- Libera recursos imediatamente
- Previne vazamentos de memória
- Garante limpeza mesmo em caso de erro

### Padrão 3: Verificar Cancelamento em Loops

```go
// ✅ CORRETO
func processarLista(ctx context.Context, items []Item) error {
    for _, item := range items {
        // Verificar ANTES de processar
        select {
        case <-ctx.Done():
            return ctx.Err()
        default:
        }
        
        if err := processarItem(ctx, item); err != nil {
            return err
        }
    }
    return nil
}

// ❌ ERRADO
func processarLista(ctx context.Context, items []Item) error {
    for _, item := range items {
        // Não verifica cancelamento!
        if err := processarItem(ctx, item); err != nil {
            return err
        }
    }
    return nil
}
```

### Padrão 4: Tipos Específicos para Chaves

```go
// ✅ CORRETO
type userIDKey struct{}
type requestIDKey struct{}

ctx := context.WithValue(ctx, userIDKey{}, "12345")
ctx = context.WithValue(ctx, requestIDKey{}, "req-abc")

// ❌ ERRADO
ctx := context.WithValue(ctx, "userID", "12345") // Pode causar colisões!
ctx = context.WithValue(ctx, "requestID", "req-abc")
```

**Por quê?**
- Evita colisões de chaves
- Type-safe
- Mais fácil de depurar

### Padrão 5: Não Armazenar Context em Structs

```go
// ❌ ERRADO
type Service struct {
    ctx context.Context // ERRADO!
    db  *sql.DB
}

func (s *Service) Processar() error {
    // Usar s.ctx...
}

// ✅ CORRETO
type Service struct {
    db *sql.DB
}

func (s *Service) Processar(ctx context.Context) error {
    // Context vem como parâmetro
    // ...
}
```

**Por quê?**
- Context é específico de requisição
- Structs geralmente têm vida útil longa
- Context deve ser passado, não armazenado

---

## 🚫 Anti-Padrões: O Que NÃO Fazer

### Anti-Padrão 1: Context Nil

```go
// ❌ ERRADO
func processar(ctx context.Context) {
    if ctx == nil {
        ctx = context.Background()
    }
    // ...
}

// ✅ CORRETO
// Context nunca deve ser nil
// Se função não precisa de context, não inclua na assinatura
func processar() {
    // ...
}
```

### Anti-Padrão 2: Passar Nil Explicitamente

```go
// ❌ ERRADO
processar(nil) // Nunca faça isso!

// ✅ CORRETO
processar(context.Background())
// Ou remover context da assinatura se não é necessário
```

### Anti-Padrão 3: Context para Parâmetros Opcionais

```go
// ❌ ERRADO
func buscar(ctx context.Context, query string) []Result {
    incluirMetadata := false
    if val := ctx.Value("incluirMetadata"); val != nil {
        incluirMetadata = val.(bool)
    }
    // ...
}

// ✅ CORRETO
func buscar(ctx context.Context, query string, incluirMetadata bool) []Result {
    // Parâmetro normal
    // ...
}

// Ou use struct de opções
type BuscarOptions struct {
    IncluirMetadata bool
}

func buscar(ctx context.Context, query string, opts BuscarOptions) []Result {
    // ...
}
```

### Anti-Padrão 4: Context para Dependências

```go
// ❌ ERRADO
type dbKey struct{}
ctx := context.WithValue(ctx, dbKey{}, db)

func processar(ctx context.Context) {
    db := ctx.Value(dbKey{}).(*sql.DB)
    // ...
}

// ✅ CORRETO
func processar(ctx context.Context, db *sql.DB) {
    // Dependência como parâmetro
    // ...
}

// Ou use injeção de dependência
type Service struct {
    db *sql.DB
}

func (s *Service) Processar(ctx context.Context) {
    // db vem do struct
    // ...
}
```

### Anti-Padrão 5: Não Verificar Cancelamento

```go
// ❌ ERRADO
func operacaoLonga(ctx context.Context) {
    for i := 0; i < 1000000; i++ {
        // Não verifica se foi cancelado!
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

### Anti-Padrão 6: Múltiplos Timeouts Conflitantes

```go
// ❌ CONFUSO
func processar(ctx context.Context) error {
    ctx1, cancel1 := context.WithTimeout(ctx, 10*time.Second)
    defer cancel1()
    
    ctx2, cancel2 := context.WithTimeout(ctx, 5*time.Second) // Qual prevalece?
    defer cancel2()
    
    // Usar ctx1 ou ctx2? Confuso!
}

// ✅ CLARO
func processar(ctx context.Context) error {
    // Timeout claro e único
    ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
    defer cancel()
    
    // Usar ctx...
}
```

---

## ⚡ Performance: Timeouts Apropriados

### Escolhendo Timeouts Corretos

**Timeouts muito curtos:**
```go
// ❌ MUITO CURTO
ctx, cancel := context.WithTimeout(ctx, 100*time.Millisecond)
// Pode cancelar operações válidas muito rápido
```

**Timeouts muito longos:**
```go
// ❌ MUITO LONGO
ctx, cancel := context.WithTimeout(ctx, 1*time.Hour)
// Operações podem travar por muito tempo
```

**Timeouts apropriados:**
```go
// ✅ APROPRIADO
// HTTP request: 5-30 segundos
ctx, cancel := context.WithTimeout(ctx, 10*time.Second)

// Database query: 5-60 segundos
ctx, cancel := context.WithTimeout(ctx, 30*time.Second)

// File I/O: 1-10 segundos
ctx, cancel := context.WithTimeout(ctx, 5*time.Second)

// Background job: minutos a horas
ctx, cancel := context.WithTimeout(ctx, 1*time.Hour)
```

### Hierarquia de Timeouts

```go
// ✅ BOM: Timeout global mais longo, timeouts específicos mais curtos
func processarRequisicao(ctx context.Context) error {
    // Timeout global da requisição: 30 segundos
    ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
    defer cancel()
    
    // Timeout específico para HTTP: 5 segundos
    ctxHTTP, cancelHTTP := context.WithTimeout(ctx, 5*time.Second)
    defer cancelHTTP()
    fazerHTTPRequest(ctxHTTP)
    
    // Timeout específico para DB: 10 segundos
    ctxDB, cancelDB := context.WithTimeout(ctx, 10*time.Second)
    defer cancelDB()
    fazerQuery(ctxDB)
    
    return nil
}
```

---

## 🔒 Segurança: Valores no Context

### O Que Pode Ir no Context

**✅ SEGURO:**
- Request ID (para logging/tracing)
- User ID (após autenticação)
- Trace ID (para distributed tracing)
- Correlation ID
- Request metadata (não sensível)

**❌ NUNCA:**
- Senhas ou tokens
- Dados sensíveis
- Chaves de API
- Informações de autenticação completas

```go
// ❌ PERIGOSO
ctx := context.WithValue(ctx, "password", senha) // NUNCA!
ctx = context.WithValue(ctx, "apiKey", key) // NUNCA!

// ✅ SEGURO
ctx := context.WithValue(ctx, userIDKey{}, userID) // OK
ctx = context.WithValue(ctx, requestIDKey{}, reqID) // OK
```

---

## 📊 Resumo: Checklist de Boas Práticas

### ✅ Sempre Faça:

- [ ] Passe context como primeiro parâmetro
- [ ] Use `defer cancel()` após criar context com cancelamento
- [ ] Verifique `ctx.Done()` em loops longos
- [ ] Use tipos específicos para chaves de valores
- [ ] Use context em operações I/O
- [ ] Use timeouts apropriados para cada tipo de operação
- [ ] Propague context através de todas as camadas

### ❌ Nunca Faça:

- [ ] Passar context nil
- [ ] Esquecer de cancelar context
- [ ] Usar context para parâmetros opcionais
- [ ] Usar context para dependências
- [ ] Armazenar context em structs
- [ ] Usar strings como chaves de valores
- [ ] Armazenar dados sensíveis no context
- [ ] Ignorar cancelamento em loops

---

## 🎯 Decisões de Design

### Quando Adicionar Context a uma Função?

**Adicione context se:**
- ✅ Função faz I/O (HTTP, DB, arquivos)
- ✅ Função pode demorar
- ✅ Função pode ser cancelada
- ✅ Função precisa de valores de requisição

**Não adicione context se:**
- ❌ Função é puramente computacional e rápida
- ❌ Função não faz I/O
- ❌ Função não precisa ser cancelada
- ❌ Função não precisa de valores de requisição

### Exemplo: Evolução de uma Função

```go
// Versão 1: Sem context (OK para função simples)
func calcularTotal(items []Item) float64 {
    total := 0.0
    for _, item := range items {
        total += item.Preco
    }
    return total
}

// Versão 2: Com context (necessário se busca dados)
func calcularTotal(ctx context.Context, userID string) (float64, error) {
    // Buscar items do banco (precisa de context!)
    items, err := buscarItems(ctx, userID)
    if err != nil {
        return 0, err
    }
    
    total := 0.0
    for _, item := range items {
        total += item.Preco
    }
    return total, nil
}
```

---

## 🚀 Performance: Dicas Finais

### 1. Meça, Não Adivinhe

```go
// Sempre meça performance antes de otimizar
func BenchmarkComContext(b *testing.B) {
    ctx := context.Background()
    for i := 0; i < b.N; i++ {
        processar(ctx)
    }
}
```

### 2. Use Timeouts Apropriados

- **HTTP**: 5-30 segundos
- **Database**: 5-60 segundos
- **File I/O**: 1-10 segundos
- **Background jobs**: minutos a horas

### 3. Verifique Periodicamente em Loops Apertados

```go
// Em loops muito rápidos, verifique periodicamente
for i := 0; i < 1000000000; i++ {
    if i%10000 == 0 {
        select {
        case <-ctx.Done():
            return ctx.Err()
        default:
        }
    }
    // Trabalho rápido
}
```

---

E assim terminamos nossa aula sobre performance e boas práticas! Você agora sabe:

- ✅ Quando usar context e quando não usar
- ✅ Como usar context de forma eficiente
- ✅ Quais são os anti-padrões comuns
- ✅ Como escolher timeouts apropriados
- ✅ Como garantir segurança com valores no context

Lembre-se: **Context é uma ferramenta poderosa, mas deve ser usada corretamente**. Siga as boas práticas e evite os anti-padrões para escrever código Go robusto e eficiente!

Na próxima etapa, você fará os exercícios e reflexões. Boa sorte! 🚀

