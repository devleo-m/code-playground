package main

import (
	"errors"
	"log"
	"log/slog"
	"os"
	"sync"
	"time"

	zlog "github.com/rs/zerolog"
	zerologlog "github.com/rs/zerolog/log"
	"go.uber.org/zap"
)

// ============================================================================
// Exemplo 1: Log Padrão (log)
// ============================================================================

func exemploLogPadrao() {
	log.Println("=== Exemplo: log padrão ===")
	
	// Log básico
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

// ============================================================================
// Exemplo 2: slog (Structured Logging - Go 1.21+)
// ============================================================================

func exemploSlog() {
	slog.Info("=== Exemplo: slog ===")
	
	// Handler de texto padrão
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
	
	// Handler JSON
	jsonLogger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	
	jsonLogger.Info("Requisição processada",
		"method", "GET",
		"path", "/api/users",
		"status", 200,
		"duration_ms", 45)
	
	// Logger com contexto
	requestLogger := logger.With(
		"request_id", "abc123",
		"user_id", 42,
		"ip", "192.168.1.1",
	)
	
	requestLogger.Info("Requisição iniciada")
	requestLogger.Info("Processando dados")
	requestLogger.Info("Requisição concluída", "status", "sucesso")
	
	// Grupos de campos
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

// ============================================================================
// Exemplo 3: Zerolog
// ============================================================================

func exemploZerolog() {
	// Configurar formato de tempo
	zlog.TimeFieldFormat = time.RFC3339
	
	// Logger básico
	zerologlog.Info().Msg("=== Exemplo: Zerolog ===")
	
	zerologlog.Info().Msg("Aplicação iniciada")
	
	zerologlog.Debug().
		Str("usuario", "admin").
		Str("acao", "login").
		Msg("Usuário fez login")
	
	zerologlog.Warn().
		Int("tentativas", 3).
		Msg("Muitas tentativas de login")
	
	zerologlog.Error().
		Str("erro", "conexão perdida").
		Int("codigo", 500).
		Msg("Erro ao processar requisição")
	
	// Configurar nível de log
	zlog.SetGlobalLevel(zlog.InfoLevel)
	
	// Logger com contexto
	requestLogger := zerologlog.With().
		Str("request_id", "abc123").
		Int("user_id", 42).
		Logger()
	
	requestLogger.Info().Msg("Requisição iniciada")
	requestLogger.Info().Msg("Processando dados")
	requestLogger.Info().Str("status", "sucesso").Msg("Requisição concluída")
	
	// Sub-loggers por componente
	dbLogger := zerologlog.With().Str("component", "database").Logger()
	apiLogger := zerologlog.With().Str("component", "api").Logger()
	authLogger := zerologlog.With().Str("component", "auth").Logger()
	
	dbLogger.Info().Msg("Conexão com banco estabelecida")
	apiLogger.Info().Str("endpoint", "/users").Msg("Endpoint registrado")
	authLogger.Warn().Msg("Token expirando em breve")
	
	// Pretty console para desenvolvimento
	zerologlog.Logger = zerologlog.Output(zlog.ConsoleWriter{Out: os.Stderr})
	
	zerologlog.Info().
		Str("method", "GET").
		Str("path", "/api/users").
		Int("status", 200).
		Msg("Requisição processada")
}

// ============================================================================
// Exemplo 4: Zap
// ============================================================================

func exemploZap() {
	// Logger de produção (JSON)
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()
	
	logger.Info("=== Exemplo: Zap ===")
	
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
	
	// API Sugared (mais conveniente)
	devLogger, _ := zap.NewDevelopment()
	defer devLogger.Sync()
	
	sugar := devLogger.Sugar()
	
	sugar.Infof("Aplicação iniciada: versão %s", "1.0.0")
	sugar.Debugw("Debug message",
		"usuario", "admin",
		"acao", "login")
	sugar.Warnw("Aviso",
		"tentativas", 3)
	sugar.Errorw("Erro",
		"erro", "conexão perdida",
		"codigo", 500)
	
	// Logger com contexto
	requestLogger := logger.With(
		zap.String("request_id", "abc123"),
		zap.Int("user_id", 42),
	)
	
	requestLogger.Info("Requisição iniciada")
	requestLogger.Info("Processando dados")
	requestLogger.Info("Requisição concluída", zap.String("status", "sucesso"))
	
	// Sub-loggers por componente
	baseLogger, _ := zap.NewProduction()
	defer baseLogger.Sync()
	
	dbLogger := baseLogger.With(zap.String("component", "database"))
	apiLogger := baseLogger.With(zap.String("component", "api"))
	
	dbLogger.Info("Conexão estabelecida")
	apiLogger.Info("Endpoint registrado", zap.String("path", "/users"))
}

// ============================================================================
// Exemplo 5: Logger com Contexto (Padrão Avançado)
// ============================================================================

func exemploLoggerComContexto() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	
	// Simular contexto de requisição HTTP
	type requestContext struct {
		requestID string
		userID    int
		ip        string
	}
	
	processRequest := func(ctx requestContext) {
		requestLogger := logger.With(
			zap.String("request_id", ctx.requestID),
			zap.Int("user_id", ctx.userID),
			zap.String("ip", ctx.ip),
		)
		
		requestLogger.Info("Requisição iniciada")
		
		// Simular processamento
		time.Sleep(10 * time.Millisecond)
		
		requestLogger.Info("Processando dados")
		
		// Simular conclusão
		requestLogger.Info("Requisição concluída",
			zap.String("status", "sucesso"),
			zap.Duration("duration", 10*time.Millisecond))
	}
	
	ctx := requestContext{
		requestID: "req-123",
		userID:    42,
		ip:        "192.168.1.1",
	}
	
	processRequest(ctx)
}

// ============================================================================
// Exemplo 6: Configuração por Ambiente
// ============================================================================

func setupLoggerPorAmbiente(env string) *zap.Logger {
	if env == "production" {
		config := zap.NewProductionConfig()
		config.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
		logger, _ := config.Build()
		return logger
	}
	
	// Desenvolvimento
	config := zap.NewDevelopmentConfig()
	config.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	logger, _ := config.Build()
	return logger
}

// ============================================================================
// Exemplo 7: Tratamento de Erros com Logging
// ============================================================================

func exemploTratamentoErros() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	
	processarPagamento := func(pedidoID string, valor float64) error {
		// Simular erro
		err := errors.New("cartão recusado")
		
		if err != nil {
			logger.Error("Falha ao processar pagamento",
				zap.String("pedido_id", pedidoID),
				zap.Float64("valor", valor),
				zap.Error(err),
				zap.String("gateway", "stripe"))
			return err
		}
		
		logger.Info("Pagamento processado com sucesso",
			zap.String("pedido_id", pedidoID),
			zap.Float64("valor", valor))
		
		return nil
	}
	
	processarPagamento("ped-123", 99.99)
}

// ============================================================================
// Exemplo 8: Logging em Funções Concorrentes
// ============================================================================

func exemploLoggingConcorrente() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	
	processarItem := func(id int, wg *sync.WaitGroup) {
		defer wg.Done()
		
		itemLogger := logger.With(zap.Int("item_id", id))
		
		itemLogger.Info("Processando item")
		
		// Simular processamento
		time.Sleep(50 * time.Millisecond)
		
		itemLogger.Info("Item processado",
			zap.Duration("duration", 50*time.Millisecond))
	}
	
	var wg sync.WaitGroup
	
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go processarItem(i, &wg)
	}
	
	wg.Wait()
	logger.Info("Todos os itens processados")
}

// ============================================================================
// Função Principal
// ============================================================================

func main() {
	// Descomente o exemplo que deseja executar:
	
	// exemploLogPadrao()
	// exemploSlog()
	// exemploZerolog()
	// exemploZap()
	// exemploLoggerComContexto()
	// exemploTratamentoErros()
	// exemploLoggingConcorrente()
	
	// Exemplo de uso do logger por ambiente
	logger := setupLoggerPorAmbiente("development")
	defer logger.Sync()
	
	logger.Info("Aplicação iniciada",
		zap.String("ambiente", "development"),
		zap.String("versao", "1.0.0"))
}

