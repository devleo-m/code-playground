package main

import (
	"fmt"
	"os"
	"runtime"
)

// Variáveis de build que podem ser injetadas via -ldflags
var (
	Version   = "dev"
	BuildTime = "unknown"
	GitCommit = "unknown"
)

// Exemplo 1: Programa simples para build básico
// Execute: go build -o exemplo1 01-exemplos.go
// Depois: ./exemplo1
func exemplo1_build_basico() {
	fmt.Println("=== Exemplo 1: Build Básico ===")
	fmt.Println("Este é um executável Go compilado!")
	fmt.Println("Execute com: go build -o exemplo1 01-exemplos.go")
	fmt.Println("Depois execute: ./exemplo1")
}

// Exemplo 2: Programa com informações de versão
// Execute: go build -ldflags "-X main.Version=1.0.0 -X main.BuildTime=$(date) -X main.GitCommit=abc123" -o exemplo2 01-exemplos.go
// Depois: ./exemplo2 version
func exemplo2_com_versao() {
	fmt.Println("=== Exemplo 2: Build com Versão ===")
	
	if len(os.Args) > 1 && os.Args[1] == "version" {
		fmt.Printf("Versão: %s\n", Version)
		fmt.Printf("Build Time: %s\n", BuildTime)
		fmt.Printf("Git Commit: %s\n", GitCommit)
		return
	}
	
	fmt.Println("Use 'version' para ver informações de build")
	fmt.Println("\nPara compilar com versão:")
	fmt.Println(`go build -ldflags "-X main.Version=1.0.0 -X main.BuildTime=$(date) -X main.GitCommit=abc123" -o exemplo2 01-exemplos.go`)
}

// Exemplo 3: Programa que mostra informações da plataforma
// Execute: go build -o exemplo3 01-exemplos.go
// Depois: ./exemplo3
// Para cross-compilation: GOOS=linux GOARCH=amd64 go build -o exemplo3-linux 01-exemplos.go
func exemplo3_info_plataforma() {
	fmt.Println("=== Exemplo 3: Informações de Plataforma ===")
	fmt.Printf("Sistema Operacional: %s\n", runtime.GOOS)
	fmt.Printf("Arquitetura: %s\n", runtime.GOARCH)
	fmt.Printf("Versão do Go: %s\n", runtime.Version())
	fmt.Printf("Número de CPUs: %d\n", runtime.NumCPU())
	
	fmt.Println("\nPara cross-compilation:")
	fmt.Println("  GOOS=linux GOARCH=amd64 go build -o exemplo3-linux 01-exemplos.go")
	fmt.Println("  GOOS=windows GOARCH=amd64 go build -o exemplo3-windows.exe 01-exemplos.go")
	fmt.Println("  GOOS=darwin GOARCH=arm64 go build -o exemplo3-mac-arm 01-exemplos.go")
}

// Exemplo 4: Comparação de tamanhos de binário
// Execute: go build -o exemplo4-normal 01-exemplos.go
// Depois: go build -ldflags "-s -w" -o exemplo4-otimizado 01-exemplos.go
// Compare os tamanhos: ls -lh exemplo4-*
func exemplo4_comparacao_tamanho() {
	fmt.Println("=== Exemplo 4: Comparação de Tamanho ===")
	fmt.Println("Este exemplo demonstra a diferença de tamanho entre builds")
	fmt.Println("\nBuild normal:")
	fmt.Println("  go build -o exemplo4-normal 01-exemplos.go")
	fmt.Println("\nBuild otimizado (menor):")
	fmt.Println("  go build -ldflags \"-s -w\" -o exemplo4-otimizado 01-exemplos.go")
	fmt.Println("\nCompare os tamanhos:")
	fmt.Println("  ls -lh exemplo4-*")
	fmt.Println("\nO build otimizado remove símbolos de debug (-s) e informações DWARF (-w)")
}

// Exemplo 5: Programa que funciona diferente em cada plataforma
// Use build tags para código específico de plataforma
func exemplo5_plataforma_especifica() {
	fmt.Println("=== Exemplo 5: Código Específico de Plataforma ===")
	
	switch runtime.GOOS {
	case "linux":
		fmt.Println("Rodando em Linux!")
		fmt.Println("Este binário foi compilado para Linux")
	case "darwin":
		fmt.Println("Rodando em macOS!")
		fmt.Println("Este binário foi compilado para macOS")
	case "windows":
		fmt.Println("Rodando em Windows!")
		fmt.Println("Este binário foi compilado para Windows")
	default:
		fmt.Printf("Rodando em %s!\n", runtime.GOOS)
	}
	
	fmt.Println("\nPara cross-compilation:")
	fmt.Println("  GOOS=linux GOARCH=amd64 go build -o exemplo5-linux 01-exemplos.go")
	fmt.Println("  GOOS=windows GOARCH=amd64 go build -o exemplo5-windows.exe 01-exemplos.go")
}

// Exemplo 6: Programa que verifica se foi compilado com CGO
func exemplo6_verificar_cgo() {
	fmt.Println("=== Exemplo 6: Verificação de CGO ===")
	
	// CGO_ENABLED é definido em tempo de compilação
	// Se CGO estiver desabilitado, algumas funcionalidades podem não estar disponíveis
	fmt.Println("Informações de build:")
	fmt.Printf("  GOOS: %s\n", runtime.GOOS)
	fmt.Printf("  GOARCH: %s\n", runtime.GOARCH)
	fmt.Printf("  Compilador: %s\n", runtime.Compiler)
	
	fmt.Println("\nPara compilar sem CGO:")
	fmt.Println("  CGO_ENABLED=0 go build -o exemplo6-no-cgo 01-exemplos.go")
	fmt.Println("\nIsso torna o binário mais portável e menor")
}

// Exemplo 7: Programa que demonstra build reproduzível
func exemplo7_build_reproduzivel() {
	fmt.Println("=== Exemplo 7: Build Reproduzível ===")
	fmt.Println("Build reproduzível gera binários idênticos")
	fmt.Println("quando compilados com as mesmas fontes")
	
	fmt.Println("\nPara build reproduzível:")
	fmt.Println(`  go build -trimpath -ldflags "-s -w -buildid=" -o exemplo7-repro 01-exemplos.go`)
	
	fmt.Println("\nBenefícios:")
	fmt.Println("  - Segurança (verificação de integridade)")
	fmt.Println("  - Debugging (mesmo binário = mesmo comportamento)")
	fmt.Println("  - CI/CD (detectar mudanças reais)")
}

// Função auxiliar para mostrar menu
func mostrarMenu() {
	fmt.Println("\n=== Exemplos de Deployment & Tooling ===")
	fmt.Println("Descomente a função desejada no main() para executar")
	fmt.Println("\nExemplos disponíveis:")
	fmt.Println("  1. exemplo1_build_basico()           - Build básico")
	fmt.Println("  2. exemplo2_com_versao()             - Build com informações de versão")
	fmt.Println("  3. exemplo3_info_plataforma()        - Informações de plataforma")
	fmt.Println("  4. exemplo4_comparacao_tamanho()     - Comparação de tamanhos")
	fmt.Println("  5. exemplo5_plataforma_especifica()   - Código específico de plataforma")
	fmt.Println("  6. exemplo6_verificar_cgo()          - Verificação de CGO")
	fmt.Println("  7. exemplo7_build_reproduzivel()      - Build reproduzível")
	fmt.Println("\nPara executar um exemplo específico:")
	fmt.Println("  go run 01-exemplos.go exemplo1")
	fmt.Println("  go run 01-exemplos.go exemplo2 version")
}

func main() {
	// Descomente o exemplo que deseja executar:
	
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "exemplo1", "1":
			exemplo1_build_basico()
		case "exemplo2", "2":
			exemplo2_com_versao()
		case "exemplo3", "3":
			exemplo3_info_plataforma()
		case "exemplo4", "4":
			exemplo4_comparacao_tamanho()
		case "exemplo5", "5":
			exemplo5_plataforma_especifica()
		case "exemplo6", "6":
			exemplo6_verificar_cgo()
		case "exemplo7", "7":
			exemplo7_build_reproduzivel()
		default:
			mostrarMenu()
		}
		return
	}
	
	// Se nenhum argumento, mostrar menu
	mostrarMenu()
	
	// Exemplos comentados - descomente para executar:
	// exemplo1_build_basico()
	// exemplo2_com_versao()
	// exemplo3_info_plataforma()
	// exemplo4_comparacao_tamanho()
	// exemplo5_plataforma_especifica()
	// exemplo6_verificar_cgo()
	// exemplo7_build_reproduzivel()
}

