# Aula 22: Ecosystem & Popular Libraries - Building CLIs

Olá, futuro(a) Gopher!

Bem-vindo(a) à aula 22! Agora que você já domina os fundamentos de Go, vamos explorar uma das áreas onde Go realmente brilha: **construção de ferramentas de linha de comando (CLIs)**.

Go é amplamente reconhecido por sua excelência em desenvolvimento de CLIs. Ferramentas como Docker, Kubernetes (kubectl), Terraform, GitHub CLI e muitas outras foram construídas com Go. Nesta aula, vamos aprender por que Go é tão bom para CLIs e como usar tanto a biblioteca padrão quanto frameworks populares para criar suas próprias ferramentas.

---

## Por que Go é Excelente para CLIs?

Antes de mergulharmos no código, vamos entender por que Go se destaca nessa área:

### 1. **Compilação Rápida**
Go compila extremamente rápido, permitindo um ciclo de desenvolvimento ágil. Você edita, compila e testa em segundos.

### 2. **Binário Único**
O compilador do Go gera um único arquivo executável, sem dependências externas. Isso significa distribuição simples: copie o binário e execute. Não precisa instalar runtime, bibliotecas ou gerenciar dependências.

### 3. **Cross-Compilation Nativa**
Go suporta compilação cruzada nativamente. Com um único comando, você pode gerar binários para Windows, Linux, macOS, ARM, e muitas outras plataformas.

```bash
# Compilar para diferentes plataformas
GOOS=linux GOARCH=amd64 go build
GOOS=windows GOARCH=amd64 go build
GOOS=darwin GOARCH=arm64 go build
```

### 4. **Performance**
Binários Go são rápidos e eficientes em uso de memória, ideais para ferramentas que precisam iniciar rapidamente.

### 5. **Ecossistema Rico**
A comunidade Go desenvolveu excelentes bibliotecas e frameworks para construção de CLIs, desde soluções simples até frameworks completos para interfaces terminais sofisticadas.

---

## O Pacote `flag` Padrão

Vamos começar com o básico: o pacote `flag` da biblioteca padrão do Go. Ele é perfeito para CLIs simples e não requer dependências externas.

### Conceitos Básicos

O pacote `flag` permite definir flags (opções) que podem ser passadas via linha de comando. Vamos ver um exemplo básico:

```go
package main

import (
	"flag"
	"fmt"
)

func main() {
	// Definindo flags
	nome := flag.String("nome", "Visitante", "Seu nome")
	idade := flag.Int("idade", 0, "Sua idade")
	ativo := flag.Bool("ativo", false, "Status ativo")
	
	// Parse dos argumentos (obrigatório!)
	flag.Parse()
	
	// Usando os valores (são ponteiros!)
	fmt.Printf("Olá, %s! Você tem %d anos. Status: %v\n", 
		*nome, *idade, *ativo)
}
```

**Executando:**
```bash
go run main.go --nome João --idade 25 --ativo
# Saída: Olá, João! Você tem 25 anos. Status: true
```

### Tipos de Flags

O pacote `flag` suporta vários tipos:

```go
// Strings
nome := flag.String("nome", "padrão", "Descrição")

// Inteiros
idade := flag.Int("idade", 0, "Descrição")
porta := flag.Int64("porta", 8080, "Descrição")

// Booleanos
verbose := flag.Bool("v", false, "Modo verboso")
ativo := flag.Bool("ativo", false, "Status")

// Duração (time.Duration)
timeout := flag.Duration("timeout", 5*time.Second, "Timeout")

// Variáveis customizadas (usando flag.Var)
var lista []string
flag.Var(&lista, "lista", "Lista de valores")
```

### Usando Variáveis Existentes

Em vez de criar novas variáveis, você pode usar variáveis existentes:

```go
var (
	nome  string
	idade int
	ativo bool
)

func init() {
	flag.StringVar(&nome, "nome", "Visitante", "Seu nome")
	flag.IntVar(&idade, "idade", 0, "Sua idade")
	flag.BoolVar(&ativo, "ativo", false, "Status ativo")
}

func main() {
	flag.Parse()
	fmt.Printf("Nome: %s, Idade: %d, Ativo: %v\n", nome, idade, ativo)
}
```

### Flags Curtos e Longos

Você pode definir flags com nomes curtos e longos:

```go
verbose := flag.Bool("v", false, "Modo verboso")
flag.BoolVar(verbose, "verbose", *verbose, "Modo verboso (longa)")
```

### Argumentos Posicionais

Após o `flag.Parse()`, você pode acessar argumentos que não são flags usando `flag.Args()`:

```go
flag.Parse()
args := flag.Args() // Argumentos não-flags

for i, arg := range args {
	fmt.Printf("Argumento [%d]: %s\n", i, arg)
}
```

**Uso:**
```bash
go run main.go --nome João arquivo1.txt arquivo2.txt
# args = ["arquivo1.txt", "arquivo2.txt"]
```

### Help Customizado

Você pode personalizar a mensagem de ajuda:

```go
flag.Usage = func() {
	fmt.Fprintf(os.Stderr, "Uso: %s [opções] [comando]\n\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "Opções:\n")
	flag.PrintDefaults()
	fmt.Fprintf(os.Stderr, "\nExemplos:\n")
	fmt.Fprintf(os.Stderr, "  %s --nome João --idade 25\n", os.Args[0])
}
```

### Validação de Entrada

Sempre valide os valores recebidos:

```go
porta := flag.Int("porta", 8080, "Porta do servidor")

flag.Parse()

if *porta < 1 || *porta > 65535 {
	fmt.Printf("Erro: porta inválida (%d). Deve estar entre 1 e 65535\n", *porta)
	os.Exit(1)
}
```

### Exemplo Completo com flag

```go
package main

import (
	"flag"
	"fmt"
	"os"
)

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
	
	// Validação
	if cfg.Nome == "" {
		fmt.Println("Erro: nome não pode ser vazio")
		flag.Usage()
		os.Exit(1)
	}
	
	return cfg
}

func main() {
	cfg := parseConfig()
	
	if cfg.Verbose {
		fmt.Println("=== Modo Verboso Ativado ===")
	}
	
	fmt.Printf("Configuração:\n")
	fmt.Printf("  Nome: %s\n", cfg.Nome)
	fmt.Printf("  Idade: %d\n", cfg.Idade)
	fmt.Printf("  Ativo: %v\n", cfg.Ativo)
}
```

---

## Cobra: Framework Poderoso para CLIs Modernas

Quando você precisa de algo mais poderoso que o pacote `flag`, o **Cobra** é uma excelente escolha. É usado por ferramentas como kubectl, Hugo, GitHub CLI e muitas outras.

### Por que Cobra?

- **Subcomandos aninhados**: Organize comandos hierarquicamente
- **Flags inteligentes**: Suporte completo a flags locais e globais
- **Sugestões automáticas**: Sugere comandos similares quando você digita errado
- **Geração automática de help**: Help bem formatado automaticamente
- **Shell completion**: Suporte nativo para bash, zsh, fish, PowerShell
- **Gerador de comandos**: Ferramenta CLI para criar estrutura de comandos rapidamente

### Instalação

```bash
go get -u github.com/spf13/cobra/cobra
```

### Estrutura Básica

Vamos criar um CLI simples com Cobra:

```go
package main

import (
	"fmt"
	"os"
	
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "meucli",
	Short: "Uma ferramenta CLI de exemplo",
	Long:  `Uma ferramenta CLI completa construída com Cobra para demonstrar suas capacidades.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Olá do CLI!")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Erro: %v\n", err)
		os.Exit(1)
	}
}

func main() {
	Execute()
}
```

### Adicionando Flags

```go
var (
	nome  string
	idade int
	ativo bool
)

var rootCmd = &cobra.Command{
	Use:   "meucli",
	Short: "Uma ferramenta CLI de exemplo",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Nome: %s, Idade: %d, Ativo: %v\n", nome, idade, ativo)
	},
}

func init() {
	rootCmd.Flags().StringVarP(&nome, "nome", "n", "Visitante", "Seu nome")
	rootCmd.Flags().IntVarP(&idade, "idade", "i", 0, "Sua idade")
	rootCmd.Flags().BoolVarP(&ativo, "ativo", "a", false, "Status ativo")
}
```

**Uso:**
```bash
go run main.go --nome João --idade 25 --ativo
# ou com flags curtos:
go run main.go -n João -i 25 -a
```

### Criando Subcomandos

O poder real do Cobra está nos subcomandos:

```go
var greetCmd = &cobra.Command{
	Use:   "greet [nome]",
	Short: "Cumprimenta alguém",
	Long:  "Cumprimenta uma pessoa pelo nome",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		nome := args[0]
		fmt.Printf("Olá, %s!\n", nome)
	},
}

var calcCmd = &cobra.Command{
	Use:   "calc",
	Short: "Calculadora simples",
	Run: func(cmd *cobra.Command, args []string) {
		// Implementação da calculadora
		fmt.Println("Calculadora")
	},
}

func init() {
	rootCmd.AddCommand(greetCmd)
	rootCmd.AddCommand(calcCmd)
}
```

**Uso:**
```bash
go run main.go greet João
go run main.go calc
```

### Flags em Subcomandos

Você pode ter flags locais (apenas no subcomando) e globais (em todos os comandos):

```go
var verbose bool

var rootCmd = &cobra.Command{
	Use:   "meucli",
	Short: "CLI de exemplo",
}

var greetCmd = &cobra.Command{
	Use:   "greet [nome]",
	Short: "Cumprimenta alguém",
	Run: func(cmd *cobra.Command, args []string) {
		if verbose {
			fmt.Println("[DEBUG] Executando greet...")
		}
		fmt.Printf("Olá, %s!\n", args[0])
	},
}

func init() {
	// Flag global (disponível em todos os comandos)
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Modo verboso")
	
	rootCmd.AddCommand(greetCmd)
}
```

### Validação de Argumentos

Cobra fornece validadores de argumentos:

```go
var greetCmd = &cobra.Command{
	Use:   "greet [nome]",
	Short: "Cumprimenta alguém",
	Args:  cobra.ExactArgs(1), // Exatamente 1 argumento
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Olá, %s!\n", args[0])
	},
}

// Outros validadores:
// cobra.MinimumNArgs(1)  - pelo menos 1 argumento
// cobra.MaximumNArgs(5)  - no máximo 5 argumentos
// cobra.RangeArgs(1, 5)  - entre 1 e 5 argumentos
// cobra.NoArgs           - nenhum argumento
```

### Gerador de Comandos

Cobra vem com um gerador CLI para criar a estrutura rapidamente:

```bash
# Instalar o gerador
go install github.com/spf13/cobra-cli@latest

# Criar um novo comando
cobra-cli add greet
cobra-cli add calc
```

### Exemplo Completo com Cobra

```go
package main

import (
	"fmt"
	"os"
	"strings"
	
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "meucli",
	Short: "CLI de exemplo com Cobra",
	Long:  `Uma ferramenta CLI completa demonstrando as capacidades do Cobra.`,
}

var greetCmd = &cobra.Command{
	Use:   "greet [nome]",
	Short: "Cumprimenta alguém",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		nome := strings.Join(args, " ")
		fmt.Printf("Olá, %s!\n", nome)
	},
}

var reverseCmd = &cobra.Command{
	Use:   "reverse [texto]",
	Short: "Inverte um texto",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		texto := strings.Join(args, " ")
		reverso := ""
		for i := len(texto) - 1; i >= 0; i-- {
			reverso += string(texto[i])
		}
		fmt.Println(reverso)
	},
}

var verbose bool

func init() {
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Modo verboso")
	rootCmd.AddCommand(greetCmd)
	rootCmd.AddCommand(reverseCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Erro: %v\n", err)
		os.Exit(1)
	}
}
```

---

## urfave/cli: Simplicidade e Intuitividade

O **urfave/cli** (anteriormente codegangsta/cli) é uma biblioteca simples e intuitiva para construir CLIs. É perfeita quando você quer algo mais poderoso que `flag`, mas não precisa de toda a complexidade do Cobra.

### Por que urfave/cli?

- **API intuitiva**: Fácil de aprender e usar
- **Help automático**: Geração automática de ajuda
- **Bash completion**: Suporte para autocompletar
- **Subcomandos aninhados**: Suporte a comandos hierárquicos
- **Integração com variáveis de ambiente**: Lê valores de env vars automaticamente
- **Leve**: Menos dependências que Cobra

### Instalação

```bash
go get github.com/urfave/cli/v2
```

### Estrutura Básica

```go
package main

import (
	"fmt"
	"log"
	"os"
	
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "meucli",
		Usage: "Uma ferramenta CLI de exemplo",
		Action: func(c *cli.Context) error {
			fmt.Println("Olá do CLI!")
			return nil
		},
	}
	
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
```

### Adicionando Flags

```go
app := &cli.App{
	Name:  "meucli",
	Usage: "CLI de exemplo",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "nome",
			Aliases: []string{"n"},
			Usage:   "Seu nome",
			Value:   "Visitante",
		},
		&cli.IntFlag{
			Name:    "idade",
			Aliases: []string{"i"},
			Usage:   "Sua idade",
			Value:   0,
		},
		&cli.BoolFlag{
			Name:    "ativo",
			Aliases: []string{"a"},
			Usage:   "Status ativo",
		},
	},
	Action: func(c *cli.Context) error {
		fmt.Printf("Nome: %s, Idade: %d, Ativo: %v\n",
			c.String("nome"),
			c.Int("idade"),
			c.Bool("ativo"))
		return nil
	},
}
```

### Criando Comandos

```go
app := &cli.App{
	Name:  "meucli",
	Usage: "CLI de exemplo",
	Commands: []*cli.Command{
		{
			Name:    "greet",
			Aliases: []string{"g"},
			Usage:   "Cumprimenta alguém",
			Action: func(c *cli.Context) error {
				nome := c.Args().First()
				if nome == "" {
					return fmt.Errorf("nome é obrigatório")
				}
				fmt.Printf("Olá, %s!\n", nome)
				return nil
			},
		},
		{
			Name:  "calc",
			Usage: "Calculadora",
			Subcommands: []*cli.Command{
				{
					Name:  "add",
					Usage: "Soma dois números",
					Action: func(c *cli.Context) error {
						// Implementação
						return nil
					},
				},
			},
		},
	},
}
```

### Variáveis de Ambiente

urfave/cli suporta variáveis de ambiente automaticamente:

```go
&cli.StringFlag{
	Name:    "porta",
	Aliases: []string{"p"},
	Usage:   "Porta do servidor",
	Value:   "8080",
	EnvVars: []string{"PORT", "SERVER_PORT"},
}
```

O valor será lido primeiro da flag, depois da variável de ambiente.

### Exemplo Completo com urfave/cli

```go
package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "meucli",
		Usage: "CLI de exemplo com urfave/cli",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "verbose",
				Aliases: []string{"v"},
				Usage:   "Modo verboso",
			},
		},
		Commands: []*cli.Command{
			{
				Name:    "greet",
				Aliases: []string{"g"},
				Usage:   "Cumprimenta alguém",
				Action: func(c *cli.Context) error {
					nome := c.Args().First()
					if nome == "" {
						return fmt.Errorf("nome é obrigatório")
					}
					if c.Bool("verbose") {
						fmt.Println("[DEBUG] Executando greet...")
					}
					fmt.Printf("Olá, %s!\n", nome)
					return nil
				},
			},
			{
				Name:  "reverse",
				Usage: "Inverte um texto",
				Action: func(c *cli.Context) error {
					texto := strings.Join(c.Args().Slice(), " ")
					if texto == "" {
						return fmt.Errorf("texto é obrigatório")
					}
					reverso := ""
					for i := len(texto) - 1; i >= 0; i-- {
						reverso += string(texto[i])
					}
					fmt.Println(reverso)
					return nil
				},
			},
		},
	}
	
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
```

---

## Bubble Tea: Interfaces Terminais Interativas

Quando você precisa criar interfaces terminais mais sofisticadas, com interação em tempo real, o **Bubble Tea** é a escolha perfeita. Ele é baseado na Arquitetura Elm e usa o padrão model-update-view.

### Por que Bubble Tea?

- **Interfaces interativas**: Suporte a entrada de teclado, mouse, e atualizações em tempo real
- **Arquitetura Elm**: Padrão model-update-view bem definido
- **Componentes reutilizáveis**: Biblioteca de componentes prontos
- **Estilização**: Suporte a cores, estilos e layouts
- **Composição**: Combine componentes para criar interfaces complexas

### Instalação

```bash
go get github.com/charmbracelet/bubbletea
```

### Conceitos Básicos

Bubble Tea usa três conceitos principais:

1. **Model**: O estado da aplicação
2. **Update**: Função que processa mensagens e atualiza o modelo
3. **View**: Função que renderiza o modelo

### Exemplo Básico

```go
package main

import (
	"fmt"
	"os"
	
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	choices  []string
	cursor   int
	selected map[int]struct{}
}

func initialModel() model {
	return model{
		choices: []string{"Opção 1", "Opção 2", "Opção 3"},
		selected: make(map[int]struct{}),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "enter", " ":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		}
	}
	return m, nil
}

func (m model) View() string {
	s := "Selecione uma opção:\n\n"
	
	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}
		
		checked := " "
		if _, ok := m.selected[i]; ok {
			checked = "x"
		}
		
		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}
	
	s += "\nPressione q para sair.\n"
	return s
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Erro: %v", err)
		os.Exit(1)
	}
}
```

### Componentes do Bubble Tea

Bubble Tea vem com vários componentes prontos através do pacote `bubbles`:

```bash
go get github.com/charmbracelet/bubbles
```

Exemplos de componentes:
- **Textinput**: Campo de entrada de texto
- **List**: Lista selecionável
- **Spinner**: Indicador de carregamento
- **Progress**: Barra de progresso
- **Table**: Tabela de dados

### Exemplo com Textinput

```go
import (
	"fmt"
	"os"
	
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	textInput textinput.Model
	err       error
}

func initialModel() model {
	ti := textinput.New()
	ti.Placeholder = "Digite seu nome..."
	ti.Focus()
	ti.CharLimit = 50
	ti.Width = 20
	
	return model{
		textInput: ti,
		err:       nil,
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			fmt.Printf("Você digitou: %s\n", m.textInput.Value())
			return m, tea.Quit
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		}
	}
	
	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return fmt.Sprintf(
		"Qual é o seu nome?\n\n%s\n\n(pressione Enter para confirmar, Esc para sair)",
		m.textInput.View(),
	)
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Erro: %v", err)
		os.Exit(1)
	}
}
```

---

## Comparação: Quando Usar Cada Ferramenta?

### Pacote `flag` (Padrão)
**Use quando:**
- CLI muito simples
- Não precisa de subcomandos complexos
- Quer zero dependências externas
- É uma ferramenta interna ou script

**Exemplo:** Script de automação, ferramenta de desenvolvimento pessoal

### Cobra
**Use quando:**
- Precisa de subcomandos aninhados complexos
- Quer shell completion automático
- Está construindo uma ferramenta pública
- Precisa de uma estrutura bem organizada

**Exemplo:** kubectl, GitHub CLI, ferramentas de infraestrutura

### urfave/cli
**Use quando:**
- Quer algo mais simples que Cobra
- Precisa de integração com variáveis de ambiente
- Quer uma API intuitiva e direta
- Não precisa de subcomandos muito complexos

**Exemplo:** Ferramentas de desenvolvimento, CLIs de aplicações web

### Bubble Tea
**Use quando:**
- Precisa de interface interativa
- Quer dashboards no terminal
- Precisa de entrada de teclado em tempo real
- Está construindo uma ferramenta com UX rica

**Exemplo:** Ferramentas de monitoramento, editores de configuração interativos, dashboards

---

## Cross-Compilation: Compilando para Múltiplas Plataformas

Uma das grandes vantagens do Go é a compilação cruzada. Você pode compilar seu CLI para diferentes plataformas a partir de uma única máquina.

### Compilação Básica

```bash
# Linux 64-bit
GOOS=linux GOARCH=amd64 go build -o meucli-linux

# Windows 64-bit
GOOS=windows GOARCH=amd64 go build -o meucli-windows.exe

# macOS Intel
GOOS=darwin GOARCH=amd64 go build -o meucli-darwin-amd64

# macOS Apple Silicon
GOOS=darwin GOARCH=arm64 go build -o meucli-darwin-arm64

# Linux ARM
GOOS=linux GOARCH=arm64 go build -o meucli-linux-arm64
```

### Script de Build para Múltiplas Plataformas

```bash
#!/bin/bash
# build.sh

APP_NAME="meucli"
VERSION="1.0.0"

PLATFORMS=(
	"linux/amd64"
	"linux/arm64"
	"darwin/amd64"
	"darwin/arm64"
	"windows/amd64"
)

for platform in "${PLATFORMS[@]}"; do
	GOOS=${platform%/*}
	GOARCH=${platform#*/}
	
	output_name="${APP_NAME}-${VERSION}-${GOOS}-${GOARCH}"
	if [ $GOOS = "windows" ]; then
		output_name+='.exe'
	fi
	
	echo "Compilando para ${GOOS}/${GOARCH}..."
	GOOS=$GOOS GOARCH=$GOARCH go build -o "dist/${output_name}"
done

echo "Compilação concluída!"
```

### Usando goreleaser

Para releases profissionais, considere usar [goreleaser](https://goreleaser.com/):

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
```

---

## Boas Práticas para CLIs

### 1. **Mensagens de Erro Claras**
```go
if err != nil {
	fmt.Fprintf(os.Stderr, "Erro ao processar arquivo: %v\n", err)
	os.Exit(1)
}
```

### 2. **Códigos de Saída Apropriados**
```go
os.Exit(0)  // Sucesso
os.Exit(1)  // Erro geral
os.Exit(2)  // Uso incorreto
```

### 3. **Validação de Entrada**
Sempre valide os dados de entrada antes de processar.

### 4. **Help Útil**
Forneça mensagens de ajuda claras e exemplos de uso.

### 5. **Logging Estruturado**
Para ferramentas complexas, considere usar logging estruturado:
```go
import "log/slog"

logger := slog.New(slog.NewTextHandler(os.Stderr, nil))
logger.Info("Processando arquivo", "arquivo", nomeArquivo)
```

### 6. **Testes**
Teste seus CLIs como qualquer outro código Go:
```go
func TestCLI(t *testing.T) {
	os.Args = []string{"meucli", "greet", "João"}
	// Testar comportamento
}
```

---

## Conclusão

Nesta aula, exploramos o ecossistema de desenvolvimento de CLIs em Go:

1. **Pacote flag**: Perfeito para CLIs simples, zero dependências
2. **Cobra**: Framework poderoso para CLIs complexas com subcomandos
3. **urfave/cli**: Solução intermediária, intuitiva e fácil de usar
4. **Bubble Tea**: Para interfaces terminais interativas e sofisticadas

Cada ferramenta tem seu lugar. Escolha baseado nas necessidades do seu projeto:
- Simplicidade → `flag`
- Poder e estrutura → Cobra
- Equilíbrio → urfave/cli
- Interatividade → Bubble Tea

Na próxima aula, vamos ver uma versão simplificada deste conteúdo e depois praticar com exercícios!

Até lá, experimente criar seu próprio CLI. Comece simples e vá evoluindo conforme necessário!

