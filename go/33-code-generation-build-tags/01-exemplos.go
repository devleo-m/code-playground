package main

import "fmt"

// ============================================
// EXEMPLO 1: go generate com stringer
// ============================================

//go:generate stringer -type=Status
// Para usar este exemplo:
// 1. Instale: go install golang.org/x/tools/cmd/stringer@latest
// 2. Execute: go generate ./...
// 3. Isso gerará status_string.go automaticamente

type Status int

const (
	StatusPending Status = iota
	StatusInProgress
	StatusCompleted
	StatusCancelled
)

func exemploStringer() {
	status := StatusInProgress
	fmt.Printf("Status: %s\n", status.String())
	// Output: Status: StatusInProgress
}

// ============================================
// EXEMPLO 2: Build Tags para diferentes OS
// ============================================

// Este arquivo funciona em todas as plataformas
// Para ver exemplos específicos de plataforma, veja:
// - config_linux.go (com //go:build linux)
// - config_windows.go (com //go:build windows)

func exemploBuildTags() {
	path := getConfigPath()
	fmt.Printf("Config path: %s\n", path)
}

// getConfigPath será implementado em arquivos com build tags
// config_linux.go: //go:build linux
// config_windows.go: //go:build windows
func getConfigPath() string {
	// Implementação default (será sobrescrita por arquivos com build tags)
	return "./config.json"
}

// ============================================
// EXEMPLO 3: Feature Flags com Build Tags
// ============================================

// Este exemplo mostra como usar feature flags
// Veja features_dev.go e features_prod.go para implementações

func exemploFeatureFlags() {
	if EnableDebugLogging {
		fmt.Println("Debug logging is enabled")
	}
	
	fmt.Printf("Log level: %s\n", LogLevel)
	fmt.Printf("Max connections: %d\n", MaxConnections)
}

// Essas constantes serão definidas em arquivos com build tags:
// - features_dev.go (//go:build dev)
// - features_prod.go (//go:build !dev)
var (
	EnableDebugLogging bool
	LogLevel           string
	MaxConnections     int
)

// ============================================
// EXEMPLO 4: Combinando go generate e Build Tags
// ============================================

//go:generate stringer -type=LogLevel
// Este comando gera o método String() para LogLevel
// Execute: go generate ./...

type LogLevel int

const (
	LogLevelDebug LogLevel = iota
	LogLevelInfo
	LogLevelWarning
	LogLevelError
)

func exemploCombinado() {
	level := LogLevelWarning
	fmt.Printf("Log level: %s\n", level.String())
	
	// DefaultLogLevel será definido por build tags
	fmt.Printf("Default log level: %s\n", DefaultLogLevel.String())
}

// DefaultLogLevel será definido em:
// - config_dev.go (//go:build dev) -> LogLevelDebug
// - config_prod.go (//go:build !dev) -> LogLevelInfo
var DefaultLogLevel LogLevel

func main() {
	fmt.Println("=== Exemplo 1: go generate ===")
	exemploStringer()
	
	fmt.Println("\n=== Exemplo 2: Build Tags ===")
	exemploBuildTags()
	
	fmt.Println("\n=== Exemplo 3: Feature Flags ===")
	exemploFeatureFlags()
	
	fmt.Println("\n=== Exemplo 4: Combinado ===")
	exemploCombinado()
}



