# Aula 22: Performance e Boas Práticas - Building CLIs

Olá! Nesta parte da aula, vamos explorar boas práticas, otimizações de performance e padrões profissionais para desenvolvimento de CLIs em Go.

---

## Boas Práticas Gerais

### 1. Mensagens de Erro Claras e Úteis

**❌ Ruim:**
```go
if err != nil {
	fmt.Println("Erro")
	os.Exit(1)
}
```

**✅ Bom:**
```go
if err != nil {
	fmt.Fprintf(os.Stderr, "Erro ao ler arquivo '%s': %v\n", arquivo, err)
	fmt.Fprintf(os.Stderr, "Dica: Verifique se o arquivo existe e você tem permissão de leitura.\n")
	os.Exit(1)
}
```

**Princípios:**
- Use `os.Stderr` para erros (não `os.Stdout`)
- Seja específico sobre o que deu errado
- Forneça contexto (arquivo, linha, etc.)
- Dê dicas de como resolver o problema

### 2. Códigos de Saída Apropriados

Use códigos de saída padrão do Unix:

```go
const (
	ExitSuccess = 0  // Sucesso
	ExitError   = 1  // Erro geral
	ExitUsage   = 2  // Uso incorreto (argumentos inválidos)
	ExitDataErr = 65 // Erro nos dados de entrada
	ExitNoInput = 66 // Arquivo de entrada não encontrado
	ExitSoftware = 70 // Erro interno de software
)
```

**Exemplo:**
```go
func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Erro: comando obrigatório não fornecido\n")
		os.Exit(ExitUsage)
	}
	
	if err := processar(); err != nil {
		fmt.Fprintf(os.Stderr, "Erro: %v\n", err)
		os.Exit(ExitError)
	}
	
	os.Exit(ExitSuccess)
}
```

### 3. Validação de Entrada Robusta

Sempre valide entrada antes de processar:

```go
func validarPorta(porta int) error {
	if porta < 1 || porta > 65535 {
		return fmt.Errorf("porta deve estar entre 1 e 65535, recebido: %d", porta)
	}
	return nil
}

func validarEmail(email string) error {
	if email == "" {
		return fmt.Errorf("email não pode ser vazio")
	}
	if !strings.Contains(email, "@") {
		return fmt.Errorf("email inválido: %s", email)
	}
	return nil
}

// Uso
porta := flag.Int("porta", 8080, "Porta do servidor")
flag.Parse()

if err := validarPorta(*porta); err != nil {
	fmt.Fprintf(os.Stderr, "Erro: %v\n", err)
	os.Exit(ExitUsage)
}
```

### 4. Help Útil e Documentação

Forneça ajuda clara e exemplos:

```go
flag.Usage = func() {
	fmt.Fprintf(os.Stderr, "Uso: %s [opções] <comando> [argumentos]\n\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "Comandos:\n")
	fmt.Fprintf(os.Stderr, "  greet <nome>    Cumprimenta uma pessoa\n")
	fmt.Fprintf(os.Stderr, "  calc <expr>     Calcula uma expressão\n\n")
	fmt.Fprintf(os.Stderr, "Opções:\n")
	flag.PrintDefaults()
	fmt.Fprintf(os.Stderr, "\nExemplos:\n")
	fmt.Fprintf(os.Stderr, "  %s greet João\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "  %s --verbose calc '2+2'\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "\nPara mais informações, visite: https://exemplo.com/docs\n")
}
```

### 5. Logging Estruturado

Para CLIs complexas, use logging estruturado:

```go
import (
	"log/slog"
	"os"
)

func setupLogger(verbose bool) *slog.Logger {
	opts := &slog.HandlerOptions{}
	if verbose {
		opts.Level = slog.LevelDebug
	}
	
	handler := slog.NewTextHandler(os.Stderr, opts)
	return slog.New(handler)
}

func main() {
	logger := setupLogger(verbose)
	
	logger.Info("Iniciando aplicação",
		"versao", "1.0.0",
		"porta", porta)
	
	logger.Debug("Configuração carregada",
		"config", config)
	
	if err := processar(); err != nil {
		logger.Error("Erro ao processar",
			"erro", err,
			"arquivo", nomeArquivo)
		os.Exit(1)
	}
}
```

---

## Performance

### 1. Inicialização Rápida

CLIs devem iniciar rapidamente. Evite inicializações pesadas:

**❌ Ruim:**
```go
func init() {
	// Carrega configuração pesada na inicialização
	config = carregarConfiguracaoCompleta()
	bancoDados = conectarBancoDados()
}
```

**✅ Bom:**
```go
// Carregue apenas quando necessário (lazy loading)
func getConfig() *Config {
	if config == nil {
		config = carregarConfiguracao()
	}
	return config
}
```

### 2. Processamento Assíncrono

Para operações que podem ser paralelas:

```go
func processarArquivos(arquivos []string) error {
	erros := make(chan error, len(arquivos))
	
	for _, arquivo := range arquivos {
		go func(f string) {
			if err := processarArquivo(f); err != nil {
				erros <- err
			} else {
				erros <- nil
			}
		}(arquivo)
	}
	
	// Coletar resultados
	for i := 0; i < len(arquivos); i++ {
		if err := <-erros; err != nil {
			return err
		}
	}
	
	return nil
}
```

### 3. Buffering para I/O

Use buffering para operações de I/O:

```go
import (
	"bufio"
	"os"
)

func processarArquivoGrande(nomeArquivo string) error {
	arquivo, err := os.Open(nomeArquivo)
	if err != nil {
		return err
	}
	defer arquivo.Close()
	
	// Buffer de leitura
	scanner := bufio.NewScanner(arquivo)
	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, 1024*1024) // Buffer de 1MB
	
	for scanner.Scan() {
		linha := scanner.Text()
		processarLinha(linha)
	}
	
	return scanner.Err()
}
```

### 4. Evite Alocações Desnecessárias

Reutilize buffers quando possível:

```go
// ❌ Ruim: aloca novo slice a cada iteração
for _, item := range items {
	resultado := make([]byte, 0)
	resultado = append(resultado, processar(item)...)
}

// ✅ Bom: reutiliza buffer
buffer := make([]byte, 0, 1024)
for _, item := range items {
	buffer = buffer[:0] // Reset sem realocar
	buffer = append(buffer, processar(item)...)
}
```

---

## Padrões de Design

### 1. Estrutura de Comandos (Cobra)

Organize comandos de forma hierárquica:

```go
// Estrutura sugerida:
// meucli/
//   ├── cmd/
//   │   ├── root.go      // Comando raiz
//   │   ├── greet.go     // Subcomando
//   │   └── calc.go      // Subcomando
//   ├── internal/
//   │   ├── config/       // Configuração
//   │   └── utils/        // Utilitários
//   └── main.go

// cmd/root.go
var rootCmd = &cobra.Command{
	Use:   "meucli",
	Short: "CLI de exemplo",
	Long:  "Descrição longa...",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Erro: %v\n", err)
		os.Exit(1)
	}
}

// main.go
func main() {
	Execute()
}
```

### 2. Injeção de Dependências

Facilite testes com injeção de dependências:

```go
type App struct {
	config  *Config
	logger  *slog.Logger
	output  io.Writer
	erro    io.Writer
}

func NewApp(config *Config, logger *slog.Logger) *App {
	return &App{
		config: config,
		logger: logger,
		output: os.Stdout,
		erro:   os.Stderr,
	}
}

// Facilita testes
func (a *App) SetOutput(w io.Writer) {
	a.output = w
}

func (a *App) Run() error {
	fmt.Fprintf(a.output, "Executando...\n")
	// ...
	return nil
}
```

### 3. Configuração Centralizada

Centralize configuração:

```go
type Config struct {
	Nome    string
	Idade   int
	Verbose bool
	Porta   int
}

func LoadConfig() (*Config, error) {
	cfg := &Config{
		Nome:    "padrão",
		Idade:   0,
		Verbose: false,
		Porta:   8080,
	}
	
	// Carregar de flags
	flag.StringVar(&cfg.Nome, "nome", cfg.Nome, "Nome")
	flag.IntVar(&cfg.Idade, "idade", cfg.Idade, "Idade")
	flag.BoolVar(&cfg.Verbose, "verbose", cfg.Verbose, "Verbose")
	flag.IntVar(&cfg.Porta, "porta", cfg.Porta, "Porta")
	
	// Carregar de variáveis de ambiente
	if nome := os.Getenv("NOME"); nome != "" {
		cfg.Nome = nome
	}
	if porta := os.Getenv("PORTA"); porta != "" {
		if p, err := strconv.Atoi(porta); err == nil {
			cfg.Porta = p
		}
	}
	
	flag.Parse()
	
	// Validação
	if err := cfg.Validate(); err != nil {
		return nil, err
	}
	
	return cfg, nil
}

func (c *Config) Validate() error {
	if c.Porta < 1 || c.Porta > 65535 {
		return fmt.Errorf("porta inválida: %d", c.Porta)
	}
	return nil
}
```

---

## Segurança

### 1. Validação de Entrada

Nunca confie na entrada do usuário:

```go
func validarCaminhoArquivo(caminho string) error {
	// Prevenir path traversal
	if strings.Contains(caminho, "..") {
		return fmt.Errorf("caminho inválido: %s", caminho)
	}
	
	// Verificar se é um caminho absoluto ou relativo válido
	if !filepath.IsAbs(caminho) && !strings.HasPrefix(caminho, "./") {
		caminho = "./" + caminho
	}
	
	// Verificar se arquivo existe
	if _, err := os.Stat(caminho); os.IsNotExist(err) {
		return fmt.Errorf("arquivo não encontrado: %s", caminho)
	}
	
	return nil
}
```

### 2. Sanitização de Output

Sanitize output para prevenir injection:

```go
import "html"

func exibirTexto(texto string) {
	// Escapar HTML se necessário
	textoEscapado := html.EscapeString(texto)
	fmt.Println(textoEscapado)
}
```

### 3. Tratamento Seguro de Senhas

Nunca exiba senhas em logs ou output:

```go
func lerSenha() (string, error) {
	// Use bibliotecas como golang.org/x/term
	fd := int(os.Stdin.Fd())
	senha, err := term.ReadPassword(fd)
	if err != nil {
		return "", err
	}
	return string(senha), nil
}
```

---

## Testes

### 1. Testes Unitários

Teste funções isoladamente:

```go
// main.go
func calcularSoma(a, b int) int {
	return a + b
}

// main_test.go
func TestCalcularSoma(t *testing.T) {
	tests := []struct {
		a, b     int
		esperado int
	}{
		{2, 3, 5},
		{0, 0, 0},
		{-1, 1, 0},
	}
	
	for _, tt := range tests {
		resultado := calcularSoma(tt.a, tt.b)
		if resultado != tt.esperado {
			t.Errorf("calcularSoma(%d, %d) = %d; esperado %d",
				tt.a, tt.b, resultado, tt.esperado)
		}
	}
}
```

### 2. Testes de Integração

Teste o CLI completo:

```go
func TestCLICompleto(t *testing.T) {
	// Salvar args originais
	argsOriginais := os.Args
	
	// Testar comando
	os.Args = []string{"meucli", "greet", "João"}
	
	// Capturar output
	var buf bytes.Buffer
	os.Stdout = &buf
	
	main()
	
	// Verificar output
	esperado := "Olá, João!\n"
	if buf.String() != esperado {
		t.Errorf("Output esperado '%s', obteve '%s'", esperado, buf.String())
	}
	
	// Restaurar
	os.Args = argsOriginais
	os.Stdout = os.NewFile(1, "/dev/stdout")
}
```

### 3. Testes com Table-Driven

Use table-driven tests:

```go
func TestValidarPorta(t *testing.T) {
	tests := []struct {
		nome     string
		porta    int
		esperado error
	}{
		{"porta válida", 8080, nil},
		{"porta mínima", 1, nil},
		{"porta máxima", 65535, nil},
		{"porta zero", 0, errPortaInvalida},
		{"porta negativa", -1, errPortaInvalida},
		{"porta muito grande", 65536, errPortaInvalida},
	}
	
	for _, tt := range tests {
		t.Run(tt.nome, func(t *testing.T) {
			err := validarPorta(tt.porta)
			if (err != nil) != (tt.esperado != nil) {
				t.Errorf("validarPorta(%d) = %v; esperado %v",
					tt.porta, err, tt.esperado)
			}
		})
	}
}
```

---

## Distribuição e Versionamento

### 1. Versionamento Semântico

Adicione versão ao seu CLI:

```go
var (
	Version   = "dev"
	BuildTime = "unknown"
	GitCommit = "unknown"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Exibe a versão",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Versão: %s\n", Version)
		fmt.Printf("Build Time: %s\n", BuildTime)
		fmt.Printf("Git Commit: %s\n", GitCommit)
	},
}
```

**Build com versão:**
```bash
go build -ldflags "-X main.Version=1.0.0 -X main.BuildTime=$(date -u +%Y-%m-%dT%H:%M:%SZ) -X main.GitCommit=$(git rev-parse HEAD)"
```

### 2. Releases com goreleaser

Use goreleaser para releases profissionais:

```yaml
# .goreleaser.yml
project_name: meucli
builds:
  - env:
      - CGO_ENABLED=0
    goos:
      - linux
      - darwin
      - windows
    goarch:
      - amd64
      - arm64
    ldflags:
      - -s -w
      - -X main.Version={{.Version}}
      - -X main.BuildTime={{.Date}}

archives:
  - format: tar.gz
    name_template: "{{ .ProjectName }}_{{ .Version }}_{{ .Os }}_{{ .Arch }}"

checksum:
  name_template: "{{ .ProjectName }}_{{ .Version }}_checksums.txt"

snapshot:
  name_template: "{{ .Tag }}-next"
```

### 3. Shell Completion

Adicione suporte a shell completion (Cobra):

```go
var completionCmd = &cobra.Command{
	Use:   "completion [bash|zsh|fish|powershell]",
	Short: "Gera script de completion",
	Long:  "Gera script de completion para o shell especificado",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return fmt.Errorf("especifique o shell: bash, zsh, fish ou powershell")
		}
		
		switch args[0] {
		case "bash":
			return cmd.Root().GenBashCompletion(os.Stdout)
		case "zsh":
			return cmd.Root().GenZshCompletion(os.Stdout)
		case "fish":
			return cmd.Root().GenFishCompletion(os.Stdout, true)
		case "powershell":
			return cmd.Root().GenPowerShellCompletion(os.Stdout)
		default:
			return fmt.Errorf("shell não suportado: %s", args[0])
		}
	},
}
```

**Uso:**
```bash
meucli completion bash > /usr/local/etc/bash_completion.d/meucli
source /usr/local/etc/bash_completion.d/meucli
```

---

## Checklist de Boas Práticas

Antes de publicar seu CLI, verifique:

- [ ] Mensagens de erro claras e úteis
- [ ] Códigos de saída apropriados
- [ ] Validação de entrada robusta
- [ ] Help útil com exemplos
- [ ] Logging estruturado (se necessário)
- [ ] Testes unitários e de integração
- [ ] Versionamento semântico
- [ ] Cross-compilation testada
- [ ] Shell completion (se usando Cobra)
- [ ] Documentação completa
- [ ] Tratamento seguro de dados sensíveis
- [ ] Performance otimizada (inicialização rápida)
- [ ] Estrutura de código organizada

---

## Recursos Adicionais

### Bibliotecas Úteis

- **spf13/cobra**: Framework CLI poderoso
- **urfave/cli**: CLI simples e intuitivo
- **charmbracelet/bubbletea**: Interfaces terminais interativas
- **charmbracelet/bubbles**: Componentes para Bubble Tea
- **goreleaser**: Releases automatizados
- **spf13/viper**: Gerenciamento de configuração
- **fatih/color**: Cores no terminal
- **olekukonko/tablewriter**: Tabelas formatadas

### Ferramentas

- **goreleaser**: Releases e distribuição
- **golangci-lint**: Linter para Go
- **gofumpt**: Formatador de código
- **richgo**: Testes com cores

---

## Conclusão

Desenvolver CLIs profissionais em Go requer atenção a:

1. **Usabilidade**: Mensagens claras, help útil, validação robusta
2. **Performance**: Inicialização rápida, processamento eficiente
3. **Segurança**: Validação de entrada, tratamento seguro de dados
4. **Testabilidade**: Código testável, testes abrangentes
5. **Distribuição**: Versionamento, cross-compilation, releases

Seguindo essas práticas, você criará CLIs robustas, profissionais e fáceis de usar!

Boa sorte com seus projetos! 🚀

