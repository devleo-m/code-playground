# Aula 22: Ecosystem & Popular Libraries - Building CLIs (Versão Simplificada)

Olá! Esta é uma versão resumida da aula 22 sobre construção de CLIs em Go.

---

## Por que Go é Excelente para CLIs?

- ✅ **Compilação rápida**
- ✅ **Binário único** (sem dependências)
- ✅ **Cross-compilation** nativa
- ✅ **Performance** excelente
- ✅ **Ecossistema rico**

---

## 1. Pacote `flag` (Padrão)

O mais simples. Perfeito para CLIs básicas.

### Exemplo Básico

```go
package main

import (
	"flag"
	"fmt"
)

func main() {
	nome := flag.String("nome", "Visitante", "Seu nome")
	idade := flag.Int("idade", 0, "Sua idade")
	
	flag.Parse()
	
	fmt.Printf("Olá, %s! Você tem %d anos\n", *nome, *idade)
}
```

**Uso:**
```bash
go run main.go --nome João --idade 25
```

### Tipos de Flags

```go
flag.String("nome", "padrão", "Descrição")
flag.Int("idade", 0, "Descrição")
flag.Bool("ativo", false, "Descrição")
flag.Duration("timeout", 5*time.Second, "Descrição")
```

### Usando Variáveis Existentes

```go
var nome string
var idade int

func init() {
	flag.StringVar(&nome, "nome", "Visitante", "Seu nome")
	flag.IntVar(&idade, "idade", 0, "Sua idade")
}

func main() {
	flag.Parse()
	fmt.Printf("Nome: %s, Idade: %d\n", nome, idade)
}
```

### Argumentos Posicionais

```go
flag.Parse()
args := flag.Args() // Argumentos que não são flags

for _, arg := range args {
	fmt.Println(arg)
}
```

---

## 2. Cobra

Framework poderoso usado por kubectl, GitHub CLI, etc.

### Instalação

```bash
go get -u github.com/spf13/cobra/cobra
```

### Estrutura Básica

```go
package main

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "meucli",
	Short: "CLI de exemplo",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Olá do CLI!")
	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Erro: %v\n", err)
		os.Exit(1)
	}
}
```

### Flags

```go
var nome string
var idade int

func init() {
	rootCmd.Flags().StringVarP(&nome, "nome", "n", "Visitante", "Seu nome")
	rootCmd.Flags().IntVarP(&idade, "idade", "i", 0, "Sua idade")
}
```

### Subcomandos

```go
var greetCmd = &cobra.Command{
	Use:   "greet [nome]",
	Short: "Cumprimenta alguém",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Olá, %s!\n", args[0])
	},
}

func init() {
	rootCmd.AddCommand(greetCmd)
}
```

**Uso:**
```bash
go run main.go greet João
```

### Flags Globais vs Locais

```go
var verbose bool

func init() {
	// Flag global (todos os comandos)
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Modo verboso")
	
	// Flag local (apenas este comando)
	greetCmd.Flags().String("titulo", "Sr.", "Título")
}
```

---

## 3. urfave/cli

Simples e intuitivo. Perfeito para CLIs intermediárias.

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
		Usage: "CLI de exemplo",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "nome",
				Aliases: []string{"n"},
				Usage:   "Seu nome",
				Value:   "Visitante",
			},
		},
		Action: func(c *cli.Context) error {
			fmt.Printf("Olá, %s!\n", c.String("nome"))
			return nil
		},
	}
	
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
```

### Comandos

```go
app := &cli.App{
	Commands: []*cli.Command{
		{
			Name:  "greet",
			Usage: "Cumprimenta alguém",
			Action: func(c *cli.Context) error {
				nome := c.Args().First()
				fmt.Printf("Olá, %s!\n", nome)
				return nil
			},
		},
	},
}
```

### Variáveis de Ambiente

```go
&cli.StringFlag{
	Name:    "porta",
	Usage:   "Porta do servidor",
	Value:   "8080",
	EnvVars: []string{"PORT", "SERVER_PORT"},
}
```

---

## 4. Bubble Tea

Para interfaces terminais interativas.

### Instalação

```bash
go get github.com/charmbracelet/bubbletea
```

### Estrutura Básica

Bubble Tea usa três componentes:
- **Model**: Estado da aplicação
- **Update**: Processa mensagens
- **View**: Renderiza a interface

### Exemplo Mínimo

```go
package main

import (
	"fmt"
	"os"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	cursor int
	choices []string
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		}
	}
	return m, nil
}

func (m model) View() string {
	s := "Selecione:\n\n"
	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}
		s += fmt.Sprintf("%s %s\n", cursor, choice)
	}
	return s
}

func main() {
	p := tea.NewProgram(model{
		choices: []string{"Opção 1", "Opção 2", "Opção 3"},
	})
	if _, err := p.Run(); err != nil {
		fmt.Printf("Erro: %v", err)
		os.Exit(1)
	}
}
```

---

## Comparação Rápida

| Ferramenta | Complexidade | Quando Usar |
|------------|--------------|-------------|
| `flag` | ⭐ Simples | CLIs muito básicas, zero dependências |
| **Cobra** | ⭐⭐⭐ Avançado | Subcomandos complexos, ferramentas públicas |
| **urfave/cli** | ⭐⭐ Intermediário | Equilíbrio entre simplicidade e poder |
| **Bubble Tea** | ⭐⭐⭐ Avançado | Interfaces interativas, dashboards |

---

## Cross-Compilation

Compile para múltiplas plataformas:

```bash
# Linux
GOOS=linux GOARCH=amd64 go build -o meucli-linux

# Windows
GOOS=windows GOARCH=amd64 go build -o meucli-windows.exe

# macOS
GOOS=darwin GOARCH=arm64 go build -o meucli-darwin
```

---

## Resumo

1. **flag**: Use para CLIs simples
2. **Cobra**: Use para CLIs complexas com subcomandos
3. **urfave/cli**: Use para um equilíbrio entre simplicidade e poder
4. **Bubble Tea**: Use para interfaces interativas

Escolha a ferramenta certa para o seu projeto!

