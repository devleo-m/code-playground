# Aula 27 - Performance e Boas Práticas: Logging

Olá, futuro(a) Gopher!

Agora que você já entende os conceitos básicos de logging, vamos mergulhar nas **boas práticas** e **otimizações de performance**. Esta é a parte que separa código amador de código profissional em produção.

---

## Quando Usar Cada Biblioteca?

### Decisão Baseada em Necessidades

#### Use `log` Padrão Quando:
- ✅ Scripts simples e ferramentas de linha de comando
- ✅ Aplicações muito pequenas (< 100 linhas)
- ✅ Você quer **zero dependências**
- ✅ Logs são apenas para debug local
- ❌ **NÃO use** em aplicações que vão para produção
- ❌ **NÃO use** se você precisa de níveis de log
- ❌ **NÃO use** se você precisa de logs estruturados

**Exemplo de uso adequado:**
```go
// Script simples de backup
package main

import "log"

func main() {
    log.Println("Iniciando backup...")
    // ... fazer backup
    log.Println("Backup concluído!")
}
```

#### Use `slog` Quando:
- ✅ Você está usando Go 1.21 ou superior
- ✅ Quer logs estruturados **sem dependências externas**
- ✅ Aplicações de médio porte
- ✅ Performance é importante, mas não crítica
- ✅ Você quer uma solução moderna da biblioteca padrão
- ❌ **NÃO use** se precisa de máxima performance (use Zap/Zerolog)
- ❌ **NÃO use** se está em Go < 1.21

**Exemplo de uso adequado:**
```go
// API REST moderna
logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
logger.Info("Requisição processada",
    "method", "GET",
    "path", "/api/users",
    "status", 200)
```

#### Use Zerolog Quando:
- ✅ **Performance é crítica**
- ✅ Você precisa de **zero-allocation** logging
- ✅ Aplicações de alta carga
- ✅ Você prefere API fluente e simples
- ✅ Logs sempre em JSON
- ❌ **NÃO use** se precisa de sampling avançado
- ❌ **NÃO use** se precisa de múltiplos formatos de saída

**Exemplo de uso adequado:**
```go
// API de alta performance
log.Info().
    Str("method", "GET").
    Str("path", "/api/users").
    Int("status", 200).
    Msg("Requisição processada")
```

#### Use Zap Quando:
- ✅ **Máxima performance** é necessária
- ✅ Aplicações de **larga escala** (milhões de eventos/segundo)
- ✅ Você precisa de **sampling** para reduzir volume
- ✅ Precisa de flexibilidade máxima (múltiplos formatos)
- ✅ Aplicações críticas em produção
- ❌ **NÃO use** se a aplicação é simples (overkill)
- ❌ **NÃO use** se você não precisa de tanta performance

**Exemplo de uso adequado:**
```go
// Microsserviço crítico de alta carga
logger, _ := zap.NewProduction()
logger.Info("Evento processado",
    zap.String("event_id", eventID),
    zap.Duration("duration", elapsed))
```

---

## Performance: O Que Realmente Importa?

### Benchmarks Reais

Em aplicações de alta carga, a diferença de performance entre bibliotecas pode ser significativa:

```
Operações por segundo (mais é melhor):

log padrão:     ~500.000 ops/sec
slog:          ~2.000.000 ops/sec
Zerolog:       ~8.000.000 ops/sec
Zap (structured): ~10.000.000 ops/sec
```

**Mas atenção**: Para a maioria das aplicações, essa diferença **não importa**!

### Quando Performance Realmente Importa?

Performance de logging importa quando:

1. **Alta Frequência de Logs**
   - Você loga **milhões de eventos por segundo**
   - Cada requisição gera múltiplos logs
   - Logs são parte do "hot path" (caminho crítico)

2. **Zero-Allocation é Crítico**
   - Aplicações com restrições de memória
   - Sistemas embarcados
   - Aplicações que precisam evitar GC pauses

3. **Latência é Crítica**
   - Sistemas de tempo real
   - Trading systems
   - Jogos online

**Para 99% das aplicações**: `slog` ou Zerolog são mais que suficientes!

---

## Boas Práticas Essenciais

### 1. Use Níveis de Log Apropriados

#### ❌ ERRADO: Tudo como ERROR
```go
logger.Error("Usuário fez login")           // ❌ Isso não é um erro!
logger.Error("Requisição processada")        // ❌ Isso não é um erro!
logger.Error("Aplicação iniciada")          // ❌ Isso não é um erro!
```

#### ✅ CORRETO: Níveis Apropriados
```go
logger.Info("Usuário fez login", zap.String("usuario", user))
logger.Info("Requisição processada", zap.Int("status", 200))
logger.Info("Aplicação iniciada", zap.String("versao", version))
logger.Error("Falha ao processar pagamento", zap.Error(err))  // ✅ Agora sim é um erro!
```

**Regra de Ouro:**
- **DEBUG**: Informações detalhadas para desenvolvimento
- **INFO**: Eventos normais e importantes da aplicação
- **WARN**: Situações que merecem atenção, mas não são erros
- **ERROR**: Erros que impedem uma operação específica
- **FATAL**: Erros críticos que param a aplicação

### 2. Sempre Adicione Contexto

#### ❌ ERRADO: Logs Sem Contexto
```go
logger.Info("Erro ocorreu")
logger.Info("Requisição processada")
logger.Error("Falha")
```

**Problema**: Como você vai encontrar o problema? Quando aconteceu? Para quem?

#### ✅ CORRETO: Logs com Contexto
```go
logger.Info("Requisição processada",
    zap.String("method", "POST"),
    zap.String("path", "/api/users"),
    zap.String("user_id", "123"),
    zap.Int("status", 201),
    zap.Duration("duration", elapsed))

logger.Error("Falha ao processar pagamento",
    zap.String("pedido_id", "456"),
    zap.String("usuario_id", "123"),
    zap.Error(err),
    zap.String("gateway", "stripe"))
```

**Regra**: Se você precisar investigar um problema, quais informações você precisaria? Adicione essas informações ao log!

### 3. Não Logue Informações Sensíveis

#### ❌ ERRADO: Informações Sensíveis
```go
logger.Info("Login realizado",
    zap.String("usuario", user),
    zap.String("senha", password),        // ❌ NUNCA!
    zap.String("token", jwtToken),        // ❌ NUNCA!
    zap.String("cartao_credito", card))   // ❌ NUNCA!
```

#### ✅ CORRETO: Informações Seguras
```go
logger.Info("Login realizado",
    zap.String("usuario", user),
    zap.String("ip", ip),
    zap.Time("timestamp", time.Now()))

// Se precisar logar token, use apenas hash ou últimos 4 caracteres
logger.Debug("Token gerado",
    zap.String("token_hash", hashToken(token)),
    zap.String("usuario", user))
```

**Nunca logue:**
- Senhas
- Tokens de autenticação completos
- Números de cartão de crédito
- Dados pessoais sensíveis (CPF, etc.)
- Chaves de API

### 4. Use Estruturação, Não Strings Formatadas

#### ❌ ERRADO: Strings Formatadas
```go
logger.Info(fmt.Sprintf("Usuário %s fez login do IP %s às %s", user, ip, time.Now()))
```

**Problemas:**
- Difícil de processar automaticamente
- Difícil de buscar
- Difícil de analisar

#### ✅ CORRETO: Campos Estruturados
```go
logger.Info("Login realizado",
    zap.String("usuario", user),
    zap.String("ip", ip),
    zap.Time("timestamp", time.Now()))
```

**Vantagens:**
- Fácil de processar (JSON)
- Fácil de buscar ("mostre todos os logins do usuário X")
- Fácil de analisar ("quantos logins por IP?")

### 5. Configure Níveis por Ambiente

#### ❌ ERRADO: Mesmo Nível em Todos os Ambientes
```go
// Sempre DEBUG, mesmo em produção
zerolog.SetGlobalLevel(zerolog.DebugLevel)
```

**Problema**: Produção vai gerar milhões de logs desnecessários!

#### ✅ CORRETO: Níveis por Ambiente
```go
func setupLogger(env string) *zap.Logger {
    if env == "production" {
        config := zap.NewProductionConfig()
        config.Level = zap.NewAtomicLevelAt(zapcore.InfoLevel) // Apenas INFO+
        logger, _ := config.Build()
        return logger
    }
    
    // Desenvolvimento: DEBUG
    config := zap.NewDevelopmentConfig()
    config.Level = zap.NewAtomicLevelAt(zapcore.DebugLevel)
    logger, _ := config.Build()
    return logger
}
```

**Regra:**
- **Desenvolvimento**: DEBUG (veja tudo)
- **Staging**: INFO (veja eventos importantes)
- **Produção**: WARN ou INFO (apenas o essencial)

### 6. Use Sampling em Alta Carga

Quando você tem milhões de eventos, nem todos precisam ser logados:

#### ✅ CORRETO: Sampling com Zap
```go
config := zap.NewProductionConfig()

// Logar as primeiras 100 mensagens, depois 1 a cada 100
config.Sampling = &zap.SamplingConfig{
    Initial:    100,
    Thereafter: 100,
}

logger, _ := config.Build()
```

**Quando usar:**
- Alta frequência de logs (milhares por segundo)
- Logs não críticos (ex: cada requisição HTTP)
- Você quer reduzir volume sem perder informações importantes

**Quando NÃO usar:**
- Logs de erro (sempre logue erros!)
- Logs críticos de negócio
- Baixa frequência de logs

### 7. Sempre Faça Sync Antes de Sair

#### ❌ ERRADO: Não Fazer Sync
```go
func main() {
    logger, _ := zap.NewProduction()
    // ... código ...
    // Aplicação termina sem garantir que logs foram escritos!
}
```

**Problema**: Logs podem ser perdidos se a aplicação terminar abruptamente!

#### ✅ CORRETO: Sempre Sync
```go
func main() {
    logger, _ := zap.NewProduction()
    defer logger.Sync() // Garante que logs sejam escritos
    
    // ... código ...
}
```

**Regra**: Sempre use `defer logger.Sync()` para garantir que logs sejam escritos antes da aplicação terminar.

### 8. Use Loggers com Contexto para Rastreamento

#### ❌ ERRADO: Logger Global Sem Contexto
```go
func handleRequest(w http.ResponseWriter, r *http.Request) {
    logger.Info("Processando requisição")
    // ... código ...
    logger.Info("Requisição concluída")
}
```

**Problema**: Como rastrear logs da mesma requisição?

#### ✅ CORRETO: Logger com Contexto
```go
func handleRequest(w http.ResponseWriter, r *http.Request) {
    requestID := generateRequestID()
    requestLogger := logger.With(zap.String("request_id", requestID))
    
    requestLogger.Info("Requisição iniciada")
    // ... código ...
    requestLogger.Info("Requisição concluída")
}
```

**Vantagem**: Todos os logs da mesma requisição terão o mesmo `request_id`, facilitando rastreamento!

---

## Anti-Padrões Comuns (O Que NÃO Fazer)

### 1. ❌ Logar Demais (Log Spam)

```go
// ❌ ERRADO: Logar cada iteração de um loop
for i := 0; i < 1000000; i++ {
    logger.Debug(fmt.Sprintf("Processando item %d", i)) // Milhões de logs!
}
```

**Solução**: Logue apenas resumos ou use sampling.

### 2. ❌ Logar Muito Pouco

```go
// ❌ ERRADO: Apenas um log no início
func main() {
    logger.Info("Aplicação iniciada")
    // ... 1000 linhas de código sem logs ...
}
```

**Solução**: Adicione logs em pontos críticos (início/fim de operações importantes, erros, etc.).

### 3. ❌ Logs Inconsistentes

```go
// ❌ ERRADO: Formato diferente em cada lugar
logger.Info("User login: " + user)
logger.Info(fmt.Sprintf("Login realizado por %s", user))
logger.Info("Login", zap.String("user", user))
```

**Solução**: Use um padrão consistente em toda a aplicação.

### 4. ❌ Não Usar Níveis Apropriados

```go
// ❌ ERRADO: Tudo como INFO
logger.Info("Erro crítico: banco de dados desconectado")
logger.Info("Debug: verificando conexão")
logger.Info("Aviso: memória em 90%")
```

**Solução**: Use níveis apropriados (ERROR, DEBUG, WARN).

### 5. ❌ Logs em Hot Path Sem Otimização

```go
// ❌ ERRADO: Logar em loop crítico sem verificar nível
func processMillions(items []Item) {
    for _, item := range items {
        logger.Debug("Processando item", zap.String("id", item.ID)) // Sempre executa!
    }
}
```

**Solução**: Verifique nível antes ou use logging condicional:
```go
if logger.Core().Enabled(zapcore.DebugLevel) {
    logger.Debug("Processando item", zap.String("id", item.ID))
}
```

---

## Padrões Avançados

### 1. Logger por Componente

Crie loggers específicos para cada componente da aplicação:

```go
var (
    dbLogger    = logger.With(zap.String("component", "database"))
    apiLogger   = logger.With(zap.String("component", "api"))
    authLogger  = logger.With(zap.String("component", "auth"))
    cacheLogger = logger.With(zap.String("component", "cache"))
)

func main() {
    dbLogger.Info("Conexão estabelecida")
    apiLogger.Info("Servidor iniciado")
    authLogger.Info("Sistema de autenticação carregado")
}
```

**Vantagem**: Fácil filtrar logs por componente!

### 2. Logger com Request ID

Para rastrear requisições HTTP:

```go
func loggingMiddleware(logger *zap.Logger) func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            requestID := r.Header.Get("X-Request-ID")
            if requestID == "" {
                requestID = generateRequestID()
            }
            
            // Adicionar ao contexto
            ctx := context.WithValue(r.Context(), "request_id", requestID)
            r = r.WithContext(ctx)
            
            // Criar logger com request_id
            requestLogger := logger.With(zap.String("request_id", requestID))
            
            // Adicionar ao contexto da requisição
            ctx = context.WithValue(ctx, "logger", requestLogger)
            r = r.WithContext(ctx)
            
            next.ServeHTTP(w, r)
        })
    }
}

// Usar no handler
func handler(w http.ResponseWriter, r *http.Request) {
    logger := r.Context().Value("logger").(*zap.Logger)
    logger.Info("Processando requisição")
}
```

### 3. Structured Logging com Grupos

Organize campos relacionados em grupos:

```go
logger.Info("Requisição processada",
    slog.Group("http",
        "method", "POST",
        "path", "/api/users",
        "status", 201,
    ),
    slog.Group("performance",
        "duration_ms", 120,
        "memory_mb", 45,
    ),
    slog.Group("user",
        "id", 123,
        "email", "user@example.com",
    ),
)
```

**Vantagem**: Logs mais organizados e fáceis de analisar!

---

## Integração com Ferramentas de Monitoramento

### 1. Enviando para Elasticsearch

```go
import "github.com/olivere/elastic/v7"

func setupElasticsearchLogger() *zap.Logger {
    config := zap.NewProductionConfig()
    
    // Customizar encoder para compatibilidade com Elasticsearch
    config.EncoderConfig.TimeKey = "@timestamp"
    config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
    
    logger, _ := config.Build()
    return logger
}
```

### 2. Enviando para CloudWatch (AWS)

```go
import "github.com/aws/aws-sdk-go/service/cloudwatchlogs"

// Usar handler customizado que envia para CloudWatch
// (implementação específica depende da biblioteca)
```

### 3. Enviando para Datadog, New Relic, etc.

Muitas ferramentas de monitoramento têm bibliotecas Go que integram com Zap/Zerolog. Sempre verifique a documentação específica.

---

## Performance: Otimizações Avançadas

### 1. Evite Alocações Desnecessárias

#### ❌ ERRADO: Alocações em Loop
```go
for i := 0; i < 1000000; i++ {
    logger.Info("Item processado", zap.String("id", fmt.Sprintf("%d", i))) // Aloca string!
}
```

#### ✅ CORRETO: Reutilizar ou Usar Tipos Primitivos
```go
for i := 0; i < 1000000; i++ {
    logger.Info("Item processado", zap.Int("id", i)) // Sem alocação!
}
```

### 2. Use API Structured do Zap em Hot Paths

```go
// ✅ Rápido (zero-allocation em muitos casos)
logger.Info("Evento", zap.String("key", "value"))

// ⚠️ Mais lento (alocações)
sugar.Infow("Evento", "key", "value")
```

### 3. Desabilite Caller Information em Produção

```go
config := zap.NewProductionConfig()
config.DisableCaller = true  // Mais rápido!
config.DisableStacktrace = true  // Mais rápido ainda!
```

**Trade-off**: Você perde informação de onde o log foi chamado, mas ganha performance.

---

## Checklist de Boas Práticas

Antes de colocar sua aplicação em produção, verifique:

- [ ] Estou usando níveis de log apropriados?
- [ ] Todos os logs têm contexto relevante?
- [ ] Não estou logando informações sensíveis?
- [ ] Estou usando logging estruturado (não strings formatadas)?
- [ ] Configurei níveis diferentes por ambiente?
- [ ] Estou fazendo sync antes de sair?
- [ ] Estou usando loggers com contexto para rastreamento?
- [ ] Não estou gerando log spam?
- [ ] Escolhi a biblioteca certa para minha necessidade?
- [ ] Testei a performance do logging em carga?

---

## Resumo: O Que Você Deve Lembrar

### Escolha da Biblioteca
- **Simples**: `log` padrão ou `slog`
- **Média carga**: `slog` ou Zerolog
- **Alta carga**: Zerolog ou Zap
- **Máxima performance**: Zap (structured API)

### Boas Práticas Essenciais
1. ✅ Use níveis apropriados
2. ✅ Sempre adicione contexto
3. ✅ Não logue informações sensíveis
4. ✅ Use estruturação, não strings
5. ✅ Configure níveis por ambiente
6. ✅ Use sampling em alta carga
7. ✅ Sempre faça sync
8. ✅ Use loggers com contexto

### Anti-Padrões a Evitar
1. ❌ Log spam
2. ❌ Logs sem contexto
3. ❌ Logs inconsistentes
4. ❌ Níveis incorretos
5. ❌ Logs em hot path sem otimização

---

## Conclusão

Logging é uma arte que separa desenvolvedores juniores de seniores. Logs bem estruturados e pensados são **essenciais** para:

- 🔍 Depuração rápida de problemas
- 📊 Análise de comportamento da aplicação
- 🚨 Detecção proativa de problemas
- 📈 Melhoria contínua da aplicação

**Lembre-se**: Logs são para **produção**, não apenas para desenvolvimento. Invista tempo em fazer logging direito desde o início!

---

E assim terminamos nossa aula sobre Performance e Boas Práticas de Logging! 

Agora você está pronto para implementar logging profissional em suas aplicações Go! 🚀

