# Aula 27: Logging em Go

Olá, futuro(a) Gopher!

Bem-vindo(a) à aula 27! Na aula anterior, exploramos ORMs e acesso a bancos de dados. Agora vamos mergulhar em um tópico fundamental para qualquer aplicação em produção: **Logging**.

Logging é essencial para monitoramento, depuração e manutenção de aplicações em produção. Sem logs adequados, você está "cego" quando algo dá errado. Nesta aula, vamos explorar desde o pacote padrão `log` até bibliotecas modernas e de alta performance como `slog`, `Zerolog` e `Zap`.

---

## Por que Logging é Importante?

Antes de mergulharmos no código, vamos entender por que logging é crucial:

### 1. **Depuração (Debugging)**
Quando algo dá errado em produção, os logs são sua primeira linha de defesa. Eles mostram o que aconteceu, quando aconteceu e em que contexto.

### 2. **Monitoramento**
Logs permitem monitorar o comportamento da aplicação em tempo real, identificar padrões e detectar problemas antes que se tornem críticos.

### 3. **Auditoria**
Em muitos sistemas, é necessário manter um registro de todas as operações importantes para fins de auditoria e conformidade.

### 4. **Análise de Performance**
Logs estruturados podem ser analisados para identificar gargalos, otimizar performance e entender o comportamento dos usuários.

### 5. **Rastreamento de Erros**
Com logs adequados, você pode rastrear erros desde sua origem até o impacto final, facilitando a correção.

---

## O Pacote `log` Padrão

Go vem com um pacote `log` na biblioteca padrão. É simples, mas funcional para casos básicos.

### Características do `log` Padrão

- **Simples**: API muito direta
- **Thread-safe**: Pode ser usado de forma concorrente
- **Sem níveis**: Não tem níveis de log (INFO, ERROR, etc.)
- **Formato fixo**: Formato de saída pré-definido
- **Sem estruturação**: Logs são apenas strings

### Exemplo Básico

```go
package main

import (
	"log"
	"os"
)

func main() {
	// Log padrão para stdout
	log.Println("Esta é uma mensagem de log")
	
	// Log para stderr
	log.SetOutput(os.Stderr)
	log.Println("Esta mensagem vai para stderr")
	
	// Log para arquivo
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Erro ao abrir arquivo de log:", err)
	}
	defer file.Close()
	
	log.SetOutput(file)
	log.Println("Esta mensagem vai para o arquivo")
}
```

### Limitações do `log` Padrão

O pacote `log` padrão tem várias limitações:

1. **Sem níveis de log**: Não diferencia entre INFO, WARN, ERROR, etc.
2. **Sem estruturação**: Não suporta campos estruturados (JSON)
3. **Formato fixo**: Não permite customização fácil do formato
4. **Sem contexto**: Difícil adicionar contexto estruturado
5. **Performance limitada**: Não otimizado para alta performance

Para aplicações simples, o `log` padrão pode ser suficiente. Mas para produção, precisamos de algo mais poderoso.

---

## O Pacote `slog` (Go 1.21+)

A partir do Go 1.21, a linguagem introduziu o pacote `slog` (structured logging) na biblioteca padrão. Ele resolve muitas limitações do `log` padrão.

### Características do `slog`

- **Estruturado**: Suporta logs estruturados (JSON, texto)
- **Níveis de log**: INFO, DEBUG, WARN, ERROR
- **Contexto**: Permite adicionar campos estruturados
- **Performance**: Otimizado para produção
- **Flexível**: Múltiplos handlers (texto, JSON, customizado)

### Instalação e Uso Básico

```go
package main

import (
	"log/slog"
	"os"
)

func main() {
	// Handler de texto padrão (stdout)
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	
	logger.Info("Aplicação iniciada",
		"versao", "1.0.0",
		"ambiente", "desenvolvimento")
	
	logger.Debug("Mensagem de debug",
		"usuario", "admin",
		"acao", "login")
	
	logger.Warn("Aviso importante",
		"tempo_restante", "5 minutos")
	
	logger.Error("Erro ocorreu",
		"erro", "conexão perdida",
		"tentativa", 3)
}
```

### Handler JSON

Para produção, geralmente queremos logs em formato JSON:

```go
package main

import (
	"log/slog"
	"os"
)

func main() {
	// Handler JSON
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	
	logger.Info("Requisição processada",
		"method", "GET",
		"path", "/api/users",
		"status", 200,
		"duration_ms", 45)
}
```

**Saída JSON:**
```json
{"time":"2024-01-15T10:30:00Z","level":"INFO","msg":"Requisição processada","method":"GET","path":"/api/users","status":200,"duration_ms":45}
```

### Configurando Níveis de Log

```go
package main

import (
	"log/slog"
	"os"
)

func main() {
	// Criar handler com nível mínimo
	opts := &slog.HandlerOptions{
		Level: slog.LevelDebug, // Aceita DEBUG e acima
	}
	
	logger := slog.New(slog.NewJSONHandler(os.Stdout, opts))
	
	// Em produção, você pode usar:
	// opts.Level = slog.LevelInfo // Apenas INFO e acima
	
	logger.Debug("Esta mensagem será exibida")
	logger.Info("Esta também")
	logger.Warn("E esta")
	logger.Error("E esta também")
}
```

### Logs com Contexto

O `slog` permite criar loggers com contexto pré-definido:

```go
package main

import (
	"log/slog"
	"os"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	
	// Logger com contexto
	requestLogger := logger.With(
		"request_id", "abc123",
		"user_id", 42,
		"ip", "192.168.1.1",
	)
	
	requestLogger.Info("Requisição iniciada")
	requestLogger.Info("Processando dados")
	requestLogger.Info("Requisição concluída", "status", "sucesso")
}
```

**Saída:**
```json
{"time":"...","level":"INFO","msg":"Requisição iniciada","request_id":"abc123","user_id":42,"ip":"192.168.1.1"}
{"time":"...","level":"INFO","msg":"Processando dados","request_id":"abc123","user_id":42,"ip":"192.168.1.1"}
{"time":"...","level":"INFO","msg":"Requisição concluída","request_id":"abc123","user_id":42,"ip":"192.168.1.1","status":"sucesso"}
```

### Grupos de Campos

Você pode agrupar campos relacionados:

```go
package main

import (
	"log/slog"
	"os"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	
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
	)
}
```

**Saída:**
```json
{
  "time":"...",
  "level":"INFO",
  "msg":"Requisição processada",
  "http":{"method":"POST","path":"/api/users","status":201},
  "performance":{"duration_ms":120,"memory_mb":45}
}
```

---

## Zerolog: Zero-Allocation JSON Logger

Zerolog é uma biblioteca de logging focada em performance e simplicidade. Seu nome vem do objetivo de ter "zero alocações" durante operações de logging.

### Características do Zerolog

- **Zero-allocation**: Minimiza alocações de memória
- **Alta performance**: Extremamente rápido
- **API fluente**: Interface limpa e intuitiva
- **JSON estruturado**: Logs sempre em JSON
- **Níveis de log**: DEBUG, INFO, WARN, ERROR, FATAL
- **Contexto**: Suporta campos estruturados

### Instalação

```bash
go get github.com/rs/zerolog
```

### Uso Básico

```go
package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"time"
)

func main() {
	// Configurar formato de tempo
	zerolog.TimeFieldFormat = time.RFC3339
	
	// Logger básico
	log.Info().Msg("Aplicação iniciada")
	
	log.Debug().
		Str("usuario", "admin").
		Str("acao", "login").
		Msg("Usuário fez login")
	
	log.Warn().
		Int("tentativas", 3).
		Msg("Muitas tentativas de login")
	
	log.Error().
		Str("erro", "conexão perdida").
		Int("codigo", 500).
		Msg("Erro ao processar requisição")
}
```

### Configurando Nível de Log

```go
package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

func main() {
	// Definir nível global
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	
	// Em desenvolvimento
	// zerolog.SetGlobalLevel(zerolog.DebugLevel)
	
	// Em produção
	// zerolog.SetGlobalLevel(zerolog.WarnLevel)
	
	log.Debug().Msg("Esta mensagem não será exibida")
	log.Info().Msg("Esta será exibida")
	log.Warn().Msg("Esta também")
}
```

### Logger com Contexto

```go
package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"context"
)

func processRequest(ctx context.Context, requestID string, userID int) {
	// Criar logger com contexto
	logger := log.With().
		Str("request_id", requestID).
		Int("user_id", userID).
		Logger()
	
	logger.Info().Msg("Requisição iniciada")
	logger.Info().Msg("Processando dados")
	logger.Info().Str("status", "sucesso").Msg("Requisição concluída")
}
```

### Pretty Console (Desenvolvimento)

Para desenvolvimento, você pode usar saída formatada no console:

```go
package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

func main() {
	// Pretty console para desenvolvimento
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	
	log.Info().
		Str("method", "GET").
		Str("path", "/api/users").
		Int("status", 200).
		Msg("Requisição processada")
}
```

### Sub-loggers

Você pode criar sub-loggers para diferentes partes da aplicação:

```go
package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var (
	dbLogger    = log.With().Str("component", "database").Logger()
	apiLogger   = log.With().Str("component", "api").Logger()
	authLogger  = log.With().Str("component", "auth").Logger()
)

func main() {
	dbLogger.Info().Msg("Conexão com banco estabelecida")
	apiLogger.Info().Str("endpoint", "/users").Msg("Endpoint registrado")
	authLogger.Warn().Msg("Token expirando em breve")
}
```

### Hooks (Ganchos)

Zerolog suporta hooks para adicionar campos automaticamente:

```go
package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

func main() {
	// Hook que adiciona timestamp em Unix
	log.Logger = log.Hook(zerolog.HookFunc(func(e *zerolog.Event, level zerolog.Level, msg string) {
		e.Int64("unix_time", time.Now().Unix())
	}))
	
	log.Info().Msg("Mensagem com timestamp Unix")
}
```

---

## Zap: High-Performance Logger by Uber

Zap é uma biblioteca de logging de alta performance desenvolvida pelo Uber. Oferece duas APIs: uma estruturada (rápida) e uma estilo printf (conveniente).

### Características do Zap

- **Alta performance**: Uma das mais rápidas bibliotecas de logging
- **Duas APIs**: Structured (rápida) e Sugared (conveniente)
- **Zero-allocation**: Em modo structured, zero alocações
- **Níveis de log**: DEBUG, INFO, WARN, ERROR, DPANIC, PANIC, FATAL
- **Sampling**: Suporta sampling para reduzir volume de logs
- **Desenvolvimento e Produção**: Diferentes configurações

### Instalação

```bash
go get go.uber.org/zap
```

### API Structured (Alta Performance)

A API structured é a mais rápida, mas menos conveniente:

```go
package main

import (
	"go.uber.org/zap"
)

func main() {
	// Logger de produção (JSON)
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()
	
	logger.Info("Aplicação iniciada",
		zap.String("versao", "1.0.0"),
		zap.String("ambiente", "producao"))
	
	logger.Debug("Mensagem de debug",
		zap.String("usuario", "admin"))
	
	logger.Warn("Aviso",
		zap.Int("tentativas", 3))
	
	logger.Error("Erro ocorreu",
		zap.String("erro", "conexão perdida"),
		zap.Int("codigo", 500))
}
```

### API Sugared (Conveniente)

A API sugared é mais conveniente, mas um pouco mais lenta:

```go
package main

import (
	"go.uber.org/zap"
)

func main() {
	// Logger de desenvolvimento (console formatado)
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()
	
	sugar := logger.Sugar()
	
	sugar.Infof("Aplicação iniciada: versão %s", "1.0.0")
	sugar.Debugw("Debug message",
		"usuario", "admin",
		"acao", "login")
	sugar.Warnw("Aviso",
		"tentativas", 3)
	sugar.Errorw("Erro",
		"erro", "conexão perdida",
		"codigo", 500)
}
```

### Configuração Customizada

```go
package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func main() {
	// Configuração customizada
	config := zap.NewProductionConfig()
	
	// Mudar nível de log
	config.Level = zap.NewAtomicLevelAt(zapcore.DebugLevel)
	
	// Mudar encoder para console (desenvolvimento)
	config.Encoding = "console"
	config.EncoderConfig = zap.NewDevelopmentEncoderConfig()
	
	logger, err := config.Build()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()
	
	logger.Debug("Debug habilitado")
	logger.Info("Info message")
}
```

### Logger com Contexto

```go
package main

import (
	"go.uber.org/zap"
)

func processRequest(logger *zap.Logger, requestID string, userID int) {
	// Criar logger com contexto
	requestLogger := logger.With(
		zap.String("request_id", requestID),
		zap.Int("user_id", userID),
	)
	
	requestLogger.Info("Requisição iniciada")
	requestLogger.Info("Processando dados")
	requestLogger.Info("Requisição concluída", zap.String("status", "sucesso"))
}
```

### Sampling (Amostragem)

Para reduzir o volume de logs em alta carga:

```go
package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

func main() {
	config := zap.NewProductionConfig()
	
	// Sampling: logar 1 a cada 100 mensagens após as primeiras 100
	config.Sampling = &zap.SamplingConfig{
		Initial:    100,
		Thereafter: 100,
	}
	
	logger, _ := config.Build()
	defer logger.Sync()
	
	// Em alta carga, apenas uma fração será logada
	for i := 0; i < 1000; i++ {
		logger.Info("Mensagem de alta frequência", zap.Int("iteracao", i))
	}
}
```

### Sub-loggers por Componente

```go
package main

import (
	"go.uber.org/zap"
)

var (
	dbLogger  *zap.Logger
	apiLogger *zap.Logger
)

func init() {
	baseLogger, _ := zap.NewProduction()
	
	dbLogger = baseLogger.With(zap.String("component", "database"))
	apiLogger = baseLogger.With(zap.String("component", "api"))
}

func main() {
	dbLogger.Info("Conexão estabelecida")
	apiLogger.Info("Endpoint registrado", zap.String("path", "/users"))
}
```

### Comparação de Performance

Zap é otimizado para performance. A API structured é mais rápida que a sugared:

```go
// Structured API (mais rápida, zero-allocation em muitos casos)
logger.Info("message", zap.String("key", "value"))

// Sugared API (mais conveniente, mas mais lenta)
sugar.Infow("message", "key", "value")
```

---

## Comparação: Quando Usar Cada Um?

### `log` Padrão
- ✅ Aplicações muito simples
- ✅ Scripts e ferramentas pequenas
- ✅ Quando você quer zero dependências
- ❌ Não use em produção para aplicações complexas

### `slog` (Go 1.21+)
- ✅ Você quer usar apenas biblioteca padrão
- ✅ Aplicações modernas em Go 1.21+
- ✅ Logs estruturados sem dependências externas
- ✅ Boa performance e flexibilidade
- ❌ Não disponível em versões antigas do Go

### Zerolog
- ✅ Alta performance é crítica
- ✅ Você quer zero-allocation
- ✅ API fluente e simples
- ✅ Aplicações de alta carga
- ❌ Menos features que Zap

### Zap
- ✅ Máxima performance necessária
- ✅ Você precisa de sampling
- ✅ Aplicações de larga escala (como Uber)
- ✅ Flexibilidade máxima
- ❌ API um pouco mais complexa

---

## Boas Práticas de Logging

### 1. Use Níveis Apropriados

```go
// ❌ ERRADO
logger.Info("Erro crítico: conexão perdida")

// ✅ CORRETO
logger.Error("Conexão perdida", zap.String("component", "database"))
```

### 2. Adicione Contexto

```go
// ❌ ERRADO
logger.Info("Requisição processada")

// ✅ CORRETO
logger.Info("Requisição processada",
	zap.String("method", "GET"),
	zap.String("path", "/api/users"),
	zap.Int("status", 200),
	zap.Duration("duration", elapsed))
```

### 3. Não Logue Informações Sensíveis

```go
// ❌ ERRADO
logger.Info("Login realizado", zap.String("senha", senha))

// ✅ CORRETO
logger.Info("Login realizado", zap.String("usuario", usuario))
```

### 4. Use Estruturação

```go
// ❌ ERRADO
logger.Info(fmt.Sprintf("Usuário %s fez login do IP %s", user, ip))

// ✅ CORRETO
logger.Info("Login realizado",
	zap.String("usuario", user),
	zap.String("ip", ip))
```

### 5. Configure Níveis por Ambiente

```go
func setupLogger(env string) *zap.Logger {
	if env == "production" {
		logger, _ := zap.NewProduction()
		return logger
	}
	logger, _ := zap.NewDevelopment()
	return logger
}
```

### 6. Use Sampling em Alta Carga

```go
config.Sampling = &zap.SamplingConfig{
	Initial:    100,
	Thereafter: 100,
}
```

### 7. Sempre Faça Sync Antes de Sair

```go
defer logger.Sync() // Garante que logs sejam escritos
```

---

## Exemplo Prático: Logger em uma API

Vamos ver um exemplo completo de como usar logging em uma API:

```go
package main

import (
	"go.uber.org/zap"
	"net/http"
	"time"
)

var logger *zap.Logger

func init() {
	var err error
	logger, err = zap.NewProduction()
	if err != nil {
		panic(err)
	}
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		
		// Criar logger com contexto da requisição
		requestLogger := logger.With(
			zap.String("method", r.Method),
			zap.String("path", r.URL.Path),
			zap.String("remote_addr", r.RemoteAddr),
		)
		
		// Criar ResponseWriter customizado para capturar status
		lw := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}
		
		next.ServeHTTP(lw, r)
		
		duration := time.Since(start)
		
		requestLogger.Info("Requisição processada",
			zap.Int("status", lw.statusCode),
			zap.Duration("duration", duration),
		)
	})
}

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/users", handleUsers)
	
	handler := loggingMiddleware(mux)
	
	logger.Info("Servidor iniciado", zap.String("port", "8080"))
	http.ListenAndServe(":8080", handler)
}

func handleUsers(w http.ResponseWriter, r *http.Request) {
	logger.Info("Processando requisição de usuários")
	// ... lógica da API
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"users":[]}`))
}
```

---

## Resumo

Nesta aula, exploramos:

1. **`log` padrão**: Simples, mas limitado
2. **`slog`**: Logging estruturado na biblioteca padrão (Go 1.21+)
3. **Zerolog**: Zero-allocation, alta performance, API fluente
4. **Zap**: Máxima performance, duas APIs, sampling

### Escolha Baseada em Necessidades:

- **Simplicidade**: `log` padrão ou `slog`
- **Performance**: Zerolog ou Zap
- **Biblioteca padrão**: `slog`
- **Máxima flexibilidade**: Zap

### Próximos Passos:

Na próxima aula, vamos criar uma versão simplificada deste conteúdo com analogias do dia a dia para fixar os conceitos!

---

E assim terminamos nossa aula sobre Logging! Espero que você tenha entendido a importância do logging e as diferentes opções disponíveis em Go.

Sinta-se à vontade para reler este material. Se tiver qualquer dúvida, pode perguntar!

