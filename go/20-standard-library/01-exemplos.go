package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

// ============================================================================
// EXEMPLO 1: I/O e File Handling
// ============================================================================

func exemploIO() {
	fmt.Println("=== Exemplo 1: I/O e File Handling ===")

	// Criar e escrever arquivo
	arquivo, err := os.Create("exemplo_io.txt")
	if err != nil {
		fmt.Printf("Erro ao criar arquivo: %v\n", err)
		return
	}
	defer arquivo.Close()

	conteudo := "Linha 1\nLinha 2\nLinha 3"
	arquivo.WriteString(conteudo)
	fmt.Println("Arquivo criado e escrito com sucesso!")

	// Ler arquivo
	arquivoLido, err := os.Open("exemplo_io.txt")
	if err != nil {
		fmt.Printf("Erro ao abrir arquivo: %v\n", err)
		return
	}
	defer arquivoLido.Close()

	dados, err := io.ReadAll(arquivoLido)
	if err != nil {
		fmt.Printf("Erro ao ler: %v\n", err)
		return
	}
	fmt.Printf("Conteúdo lido: %s\n", string(dados))

	// Limpar
	os.Remove("exemplo_io.txt")
	fmt.Println()
}

// ============================================================================
// EXEMPLO 2: Flag - Argumentos de Linha de Comando
// ============================================================================

func exemploFlag() {
	fmt.Println("=== Exemplo 2: Flag - Argumentos de Linha de Comando ===")

	// Definir flags
	nome := flag.String("nome", "Visitante", "Seu nome")
	idade := flag.Int("idade", 0, "Sua idade")
	ativo := flag.Bool("ativo", false, "Está ativo?")

	// Para demonstração, vamos simular argumentos
	os.Args = []string{"programa", "-nome=João", "-idade=25", "-ativo"}
	flag.Parse()

	fmt.Printf("Nome: %s\n", *nome)
	fmt.Printf("Idade: %d\n", *idade)
	fmt.Printf("Ativo: %v\n", *ativo)
	fmt.Printf("Argumentos restantes: %v\n", flag.Args())
	fmt.Println()
}

// ============================================================================
// EXEMPLO 3: Time - Trabalhando com Tempo
// ============================================================================

func exemploTime() {
	fmt.Println("=== Exemplo 3: Time - Trabalhando com Tempo ===")

	// Time atual
	agora := time.Now()
	fmt.Printf("Agora: %v\n", agora)
	fmt.Printf("Formato RFC3339: %s\n", agora.Format(time.RFC3339))

	// Duration
	duracao := 2*time.Hour + 30*time.Minute + 45*time.Second
	fmt.Printf("Duração: %v\n", duracao)
	fmt.Printf("Em segundos: %.0f\n", duracao.Seconds())

	// Criar Time específico
	dataEspecifica := time.Date(2024, time.January, 15, 10, 30, 0, 0, time.UTC)
	fmt.Printf("Data específica: %v\n", dataEspecifica)

	// Parsing
	tempoStr := "2024-01-15T10:30:00Z"
	tempoParsed, err := time.Parse(time.RFC3339, tempoStr)
	if err != nil {
		fmt.Printf("Erro: %v\n", err)
	} else {
		fmt.Printf("Tempo parseado: %v\n", tempoParsed)
	}

	// Operações aritméticas
	futuro := agora.Add(24 * time.Hour)
	passado := agora.Add(-7 * 24 * time.Hour)
	fmt.Printf("Amanhã: %v\n", futuro)
	fmt.Printf("Há uma semana: %v\n", passado)

	// Formatação customizada
	layout := "02/01/2006 15:04:05"
	formatado := time.Now().Format(layout)
	fmt.Printf("Formatado: %s\n", formatado)
	fmt.Println()
}

// ============================================================================
// EXEMPLO 4: JSON - Marshal e Unmarshal
// ============================================================================

type Pessoa struct {
	Nome    string  `json:"nome"`
	Idade   int     `json:"idade"`
	Email   string  `json:"email,omitempty"`
	Ativo   bool    `json:"ativo"`
	Salario float64 `json:"-"`
}

func exemploJSON() {
	fmt.Println("=== Exemplo 4: JSON - Marshal e Unmarshal ===")

	// Criar pessoa
	pessoa := Pessoa{
		Nome:    "João Silva",
		Idade:   30,
		Email:   "joao@example.com",
		Ativo:   true,
		Salario: 5000.00,
	}

	// Marshal (struct → JSON)
	jsonBytes, err := json.Marshal(pessoa)
	if err != nil {
		fmt.Printf("Erro: %v\n", err)
		return
	}
	fmt.Printf("JSON: %s\n", string(jsonBytes))

	// Marshal com indentação
	jsonIndentado, err := json.MarshalIndent(pessoa, "", "  ")
	if err != nil {
		fmt.Printf("Erro: %v\n", err)
		return
	}
	fmt.Printf("JSON indentado:\n%s\n", string(jsonIndentado))

	// Unmarshal (JSON → struct)
	jsonStr := `{"nome":"Maria","idade":25,"email":"maria@example.com","ativo":true}`
	var pessoa2 Pessoa
	err = json.Unmarshal([]byte(jsonStr), &pessoa2)
	if err != nil {
		fmt.Printf("Erro: %v\n", err)
		return
	}
	fmt.Printf("Pessoa deserializada: %+v\n", pessoa2)
	fmt.Println()
}

// ============================================================================
// EXEMPLO 5: OS - Variáveis de Ambiente e Sistema
// ============================================================================

func exemploOS() {
	fmt.Println("=== Exemplo 5: OS - Variáveis de Ambiente ===")

	// Obter variável de ambiente
	path := os.Getenv("PATH")
	if path != "" {
		fmt.Printf("PATH (primeiros 100 chars): %s...\n", path[:min(100, len(path))])
	}

	// Definir variável (apenas para este processo)
	os.Setenv("MINHA_VAR", "meu_valor")
	fmt.Printf("MINHA_VAR: %s\n", os.Getenv("MINHA_VAR"))

	// Hostname
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Printf("Erro ao obter hostname: %v\n", err)
	} else {
		fmt.Printf("Hostname: %s\n", hostname)
	}

	// Diretório atual
	dir, err := os.Getwd()
	if err != nil {
		fmt.Printf("Erro: %v\n", err)
	} else {
		fmt.Printf("Diretório atual: %s\n", dir)
	}

	// Informações do processo
	fmt.Printf("PID: %d\n", os.Getpid())
	fmt.Println()
}

// ============================================================================
// EXEMPLO 6: Bufio - I/O com Buffer
// ============================================================================

func exemploBufio() {
	fmt.Println("=== Exemplo 6: Bufio - I/O com Buffer ===")

	// Criar arquivo de exemplo
	arquivo, err := os.Create("exemplo_bufio.txt")
	if err != nil {
		fmt.Printf("Erro: %v\n", err)
		return
	}
	defer arquivo.Close()
	defer os.Remove("exemplo_bufio.txt")

	arquivo.WriteString("Linha 1\nLinha 2\nLinha 3\n")

	// Fechar e reabrir para leitura
	arquivo.Close()

	arquivoLido, err := os.Open("exemplo_bufio.txt")
	if err != nil {
		fmt.Printf("Erro: %v\n", err)
		return
	}
	defer arquivoLido.Close()

	// Scanner - ler linha por linha
	scanner := bufio.NewScanner(arquivoLido)
	linhaNum := 1
	for scanner.Scan() {
		linha := scanner.Text()
		fmt.Printf("Linha %d: %s\n", linhaNum, linha)
		linhaNum++
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Erro no scanner: %v\n", err)
	}

	// Reader com buffer
	texto := "Este é um texto de exemplo"
	reader := bufio.NewReader(strings.NewReader(texto))
	parte, err := reader.ReadString(' ')
	if err != nil && err != io.EOF {
		fmt.Printf("Erro: %v\n", err)
	} else {
		fmt.Printf("Primeira palavra: %s\n", strings.TrimSpace(parte))
	}
	fmt.Println()
}

// ============================================================================
// EXEMPLO 7: Slog - Logging Estruturado
// ============================================================================

func exemploSlog() {
	fmt.Println("=== Exemplo 7: Slog - Logging Estruturado ===")

	// Logger com texto simples
	fmt.Println("--- Logger Texto ---")
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	logger.Info("Mensagem de informação")
	logger.Warn("Aviso!")
	logger.Error("Erro ocorreu", "codigo", 500)

	// Logger com JSON
	fmt.Println("\n--- Logger JSON ---")
	jsonLogger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	jsonLogger.Info("Usuário logado",
		"usuario", "joao",
		"ip", "192.168.1.1",
		"timestamp", time.Now().Format(time.RFC3339))

	// Logger com níveis
	fmt.Println("\n--- Logger com Níveis ---")
	opts := &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}
	loggerNiveis := slog.New(slog.NewJSONHandler(os.Stdout, opts))
	loggerNiveis.Debug("Debug: informação detalhada")
	loggerNiveis.Info("Info: informação geral")
	loggerNiveis.Warn("Warn: aviso")
	loggerNiveis.Error("Error: erro ocorreu")
	fmt.Println()
}

// ============================================================================
// EXEMPLO 8: Regexp - Expressões Regulares
// ============================================================================

func exemploRegexp() {
	fmt.Println("=== Exemplo 8: Regexp - Expressões Regulares ===")

	texto := "Meu email é joao@example.com e meu telefone é (11) 98765-4321"

	// Verificar correspondência
	emailPattern := `\b[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Z|a-z]{2,}\b`
	matched, err := regexp.MatchString(emailPattern, texto)
	if err != nil {
		fmt.Printf("Erro: %v\n", err)
		return
	}
	fmt.Printf("Tem email? %v\n", matched)

	// Compilar regex (mais eficiente)
	re := regexp.MustCompile(emailPattern)
	email := re.FindString(texto)
	fmt.Printf("Email encontrado: %s\n", email)

	// Encontrar todas as correspondências
	emails := re.FindAllString(texto, -1)
	fmt.Printf("Todos os emails: %v\n", emails)

	// Grupos de captura
	telefonePattern := `\((\d{2})\)\s(\d{5})-(\d{4})`
	reTel := regexp.MustCompile(telefonePattern)
	matches := reTel.FindStringSubmatch(texto)
	if len(matches) > 0 {
		fmt.Printf("Telefone completo: %s\n", matches[0])
		fmt.Printf("DDD: %s\n", matches[1])
		fmt.Printf("Número: %s-%s\n", matches[2], matches[3])
	}

	// Substituir
	replaced := reTel.ReplaceAllString(texto, "[TELEFONE OCULTO]")
	fmt.Printf("Texto substituído: %s\n", replaced)
	fmt.Println()
}

// ============================================================================
// EXEMPLO 9: Filepath - Trabalhando com Caminhos
// ============================================================================

func exemploFilepath() {
	fmt.Println("=== Exemplo 9: Filepath - Trabalhando com Caminhos ===")

	// Join de caminhos (cross-platform)
	caminho := filepath.Join("dir1", "dir2", "arquivo.txt")
	fmt.Printf("Caminho: %s\n", caminho)

	// Separador de diretório
	fmt.Printf("Separador: %s\n", string(filepath.Separator))

	// Base (nome do arquivo)
	base := filepath.Base("/home/user/arquivo.txt")
	fmt.Printf("Base: %s\n", base)

	// Dir (diretório)
	dir := filepath.Dir("/home/user/arquivo.txt")
	fmt.Printf("Dir: %s\n", dir)

	// Ext (extensão)
	ext := filepath.Ext("arquivo.txt")
	fmt.Printf("Extensão: %s\n", ext)

	// Abs (caminho absoluto)
	abs, err := filepath.Abs(".")
	if err != nil {
		fmt.Printf("Erro: %v\n", err)
	} else {
		fmt.Printf("Caminho absoluto: %s\n", abs)
	}
	fmt.Println()
}

// ============================================================================
// EXEMPLO 10: Combinando Múltiplos Pacotes
// ============================================================================

func exemploCombinado() {
	fmt.Println("=== Exemplo 10: Combinando Múltiplos Pacotes ===")

	// Criar estrutura de dados
	type LogEntry struct {
		Timestamp time.Time `json:"timestamp"`
		Level     string    `json:"level"`
		Message   string    `json:"message"`
		User      string    `json:"user,omitempty"`
	}

	// Criar entrada de log
	entry := LogEntry{
		Timestamp: time.Now(),
		Level:     "info",
		Message:   "Usuário realizou login",
		User:      "joao@example.com",
	}

	// Serializar para JSON
	jsonBytes, err := json.MarshalIndent(entry, "", "  ")
	if err != nil {
		fmt.Printf("Erro: %v\n", err)
		return
	}

	// Escrever em arquivo com buffer
	arquivo, err := os.Create("log_entry.json")
	if err != nil {
		fmt.Printf("Erro: %v\n", err)
		return
	}
	defer arquivo.Close()
	defer os.Remove("log_entry.json")

	writer := bufio.NewWriter(arquivo)
	writer.Write(jsonBytes)
	writer.WriteString("\n")
	writer.Flush()

	fmt.Println("Log entry criado e salvo!")

	// Ler e deserializar
	arquivoLido, err := os.Open("log_entry.json")
	if err != nil {
		fmt.Printf("Erro: %v\n", err)
		return
	}
	defer arquivoLido.Close()

	var entryLida LogEntry
	decoder := json.NewDecoder(arquivoLido)
	err = decoder.Decode(&entryLida)
	if err != nil {
		fmt.Printf("Erro: %v\n", err)
		return
	}

	fmt.Printf("Log entry lida: %+v\n", entryLida)
	fmt.Println()
}

// ============================================================================
// MAIN
// ============================================================================

func main() {
	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println("    EXEMPLOS DA STANDARD LIBRARY DO GO")
	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println()

	exemploIO()
	exemploTime()
	exemploJSON()
	exemploOS()
	exemploBufio()
	exemploSlog()
	exemploRegexp()
	exemploFilepath()
	exemploCombinado()

	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println("    TODOS OS EXEMPLOS FORAM EXECUTADOS!")
	fmt.Println("═══════════════════════════════════════════════════════════")
}

