package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

// ============================================
// EXEMPLO 1: Usando o pacote flag padrão
// ============================================

func exemploFlagPadrao() {
	// Definindo flags
	nome := flag.String("nome", "Visitante", "Seu nome")
	idade := flag.Int("idade", 0, "Sua idade")
	ativo := flag.Bool("ativo", false, "Status ativo")
	
	// Parse dos argumentos
	flag.Parse()
	
	fmt.Printf("Olá, %s! Você tem %d anos. Status: %v\n", *nome, *idade, *ativo)
}

// ============================================
// EXEMPLO 2: CLI simples com flag e subcomandos
// ============================================

func exemploCLISimples() {
	if len(os.Args) < 2 {
		fmt.Println("Uso: go run 01-exemplos.go <comando>")
		fmt.Println("Comandos: hello, calc")
		os.Exit(1)
	}
	
	comando := os.Args[1]
	
	switch comando {
	case "hello":
		nome := "Mundo"
		if len(os.Args) > 2 {
			nome = os.Args[2]
		}
		fmt.Printf("Olá, %s!\n", nome)
		
	case "calc":
		if len(os.Args) < 4 {
			fmt.Println("Uso: calc <operação> <num1> <num2>")
			os.Exit(1)
		}
		// Implementação simplificada
		fmt.Println("Calculadora simples - implemente a lógica aqui")
		
	default:
		fmt.Printf("Comando desconhecido: %s\n", comando)
		os.Exit(1)
	}
}

// ============================================
// EXEMPLO 3: Estrutura para CLI mais complexa
// ============================================

type CLI struct {
	comandos map[string]func([]string)
}

func NovoCLI() *CLI {
	return &CLI{
		comandos: make(map[string]func([]string)),
	}
}

func (c *CLI) RegistrarComando(nome string, handler func([]string)) {
	c.comandos[nome] = handler
}

func (c *CLI) Executar() {
	if len(os.Args) < 2 {
		c.mostrarAjuda()
		os.Exit(1)
	}
	
	comando := os.Args[1]
	args := os.Args[2:]
	
	handler, existe := c.comandos[comando]
	if !existe {
		fmt.Printf("Comando '%s' não encontrado\n", comando)
		c.mostrarAjuda()
		os.Exit(1)
	}
	
	handler(args)
}

func (c *CLI) mostrarAjuda() {
	fmt.Println("Comandos disponíveis:")
	for cmd := range c.comandos {
		fmt.Printf("  - %s\n", cmd)
	}
}

// ============================================
// EXEMPLO 4: Processamento de flags avançado
// ============================================

func exemploFlagsAvancado() {
	var (
		arquivo = flag.String("arquivo", "", "Caminho do arquivo (obrigatório)")
		modo    = flag.String("modo", "leitura", "Modo: leitura ou escrita")
		verbose = flag.Bool("v", false, "Modo verboso")
	)
	
	flag.Parse()
	
	if *arquivo == "" {
		fmt.Println("Erro: --arquivo é obrigatório")
		flag.Usage()
		os.Exit(1)
	}
	
	if *verbose {
		fmt.Printf("Processando arquivo: %s\n", *arquivo)
		fmt.Printf("Modo: %s\n", *modo)
	}
	
	fmt.Printf("Arquivo '%s' processado com sucesso no modo '%s'\n", *arquivo, *modo)
}

// ============================================
// EXEMPLO 5: Parsing de argumentos posicionais
// ============================================

func exemploArgumentosPosicionais() {
	flag.Parse()
	
	args := flag.Args() // Argumentos não-flags
	
	if len(args) == 0 {
		fmt.Println("Nenhum argumento fornecido")
		return
	}
	
	fmt.Println("Argumentos posicionais:")
	for i, arg := range args {
		fmt.Printf("  [%d] %s\n", i, arg)
	}
}

// ============================================
// EXEMPLO 6: Validação de entrada
// ============================================

func exemploValidacao() {
	porta := flag.Int("porta", 8080, "Porta do servidor (1-65535)")
	host := flag.String("host", "localhost", "Host do servidor")
	
	flag.Parse()
	
	// Validação
	if *porta < 1 || *porta > 65535 {
		fmt.Printf("Erro: porta inválida (%d). Deve estar entre 1 e 65535\n", *porta)
		os.Exit(1)
	}
	
	if *host == "" {
		fmt.Println("Erro: host não pode ser vazio")
		os.Exit(1)
	}
	
	fmt.Printf("Servidor configurado: %s:%d\n", *host, *porta)
}

// ============================================
// EXEMPLO 7: CLI com múltiplos comandos
// ============================================

func exemploMultiplosComandos() {
	cli := NovoCLI()
	
	// Comando: greet
	cli.RegistrarComando("greet", func(args []string) {
		nome := "Mundo"
		if len(args) > 0 {
			nome = args[0]
		}
		fmt.Printf("Olá, %s!\n", nome)
	})
	
	// Comando: reverse
	cli.RegistrarComando("reverse", func(args []string) {
		if len(args) == 0 {
			fmt.Println("Erro: forneça uma string para inverter")
			return
		}
		texto := strings.Join(args, " ")
		reverso := ""
		for i := len(texto) - 1; i >= 0; i-- {
			reverso += string(texto[i])
		}
		fmt.Println(reverso)
	})
	
	// Comando: count
	cli.RegistrarComando("count", func(args []string) {
		fmt.Printf("Número de argumentos: %d\n", len(args))
		for i, arg := range args {
			fmt.Printf("  [%d] %s (tamanho: %d)\n", i, arg, len(arg))
		}
	})
	
	cli.Executar()
}

// ============================================
// EXEMPLO 8: Help customizado
// ============================================

func exemploHelpCustomizado() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Uso: %s [opções] [comando]\n\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Opções:\n")
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\nExemplos:\n")
		fmt.Fprintf(os.Stderr, "  %s --nome João --idade 25\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  %s --help\n", os.Args[0])
	}
	
	nome := flag.String("nome", "", "Seu nome")
	idade := flag.Int("idade", 0, "Sua idade")
	
	flag.Parse()
	
	if *nome == "" {
		flag.Usage()
		os.Exit(1)
	}
	
	fmt.Printf("Bem-vindo, %s (%d anos)!\n", *nome, *idade)
}

// ============================================
// EXEMPLO 9: Variáveis de ambiente
// ============================================

func exemploVariaveisAmbiente() {
	// Lê variável de ambiente ou usa valor padrão
	porta := os.Getenv("PORT")
	if porta == "" {
		porta = "8080"
	}
	
	host := os.Getenv("HOST")
	if host == "" {
		host = "localhost"
	}
	
	fmt.Printf("Servidor: %s:%s\n", host, porta)
	fmt.Println("Dica: defina PORT e HOST como variáveis de ambiente")
}

// ============================================
// EXEMPLO 10: Estrutura completa de CLI
// ============================================

type Config struct {
	Nome    string
	Idade   int
	Ativo   bool
	Verbose bool
}

func parseConfig() *Config {
	cfg := &Config{}
	
	flag.StringVar(&cfg.Nome, "nome", "Visitante", "Nome do usuário")
	flag.IntVar(&cfg.Idade, "idade", 0, "Idade do usuário")
	flag.BoolVar(&cfg.Ativo, "ativo", false, "Status ativo")
	flag.BoolVar(&cfg.Verbose, "v", false, "Modo verboso")
	flag.BoolVar(&cfg.Verbose, "verbose", false, "Modo verboso (longa)")
	
	flag.Parse()
	
	return cfg
}

func exemploEstruturaCompleta() {
	cfg := parseConfig()
	
	if cfg.Verbose {
		fmt.Println("=== Modo Verboso Ativado ===")
	}
	
	fmt.Printf("Configuração:\n")
	fmt.Printf("  Nome: %s\n", cfg.Nome)
	fmt.Printf("  Idade: %d\n", cfg.Idade)
	fmt.Printf("  Ativo: %v\n", cfg.Ativo)
}

// ============================================
// MAIN - Descomente o exemplo que deseja testar
// ============================================

func main() {
	// Descomente o exemplo que deseja executar:
	
	// exemploFlagPadrao()
	// exemploCLISimples()
	// exemploFlagsAvancado()
	// exemploArgumentosPosicionais()
	// exemploValidacao()
	// exemploMultiplosComandos()
	// exemploHelpCustomizado()
	// exemploVariaveisAmbiente()
	// exemploEstruturaCompleta()
	
	// Para testar exemploMultiplosComandos, use:
	// go run 01-exemplos.go greet João
	// go run 01-exemplos.go reverse Olá Mundo
	// go run 01-exemplos.go count um dois três
	
	fmt.Println("Descomente um exemplo no main() para testar!")
	fmt.Println("\nExemplos de uso:")
	fmt.Println("  go run 01-exemplos.go --nome João --idade 25")
	fmt.Println("  go run 01-exemplos.go greet João")
	fmt.Println("  go run 01-exemplos.go reverse Olá Mundo")
}

