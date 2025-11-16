# Módulo 30: Code Quality and Analysis
## Aula 1: Code Quality and Analysis - Ferramentas Essenciais para Qualidade de Código

Olá! Bem-vindo a este módulo crucial sobre **qualidade de código**. Até agora você aprendeu a escrever código em Go, mas uma parte fundamental do desenvolvimento profissional é garantir que seu código não apenas funcione, mas também esteja **livre de bugs potenciais, bem formatado e seguindo as melhores práticas**.

Nesta aula, vamos mergulhar em duas ferramentas essenciais do ecossistema Go que todo desenvolvedor precisa dominar: **`go vet`** e **`goimports`**. Essas ferramentas são seus aliados na busca por código de alta qualidade e manutenibilidade.

---

## 1. `go vet` - O Detetive de Bugs

### O Que É?

O `go vet` é uma **ferramenta built-in** do Go que analisa o código-fonte em busca de construções suspeitas que provavelmente são bugs. É como ter um revisor de código experiente que examina seu código procurando por problemas comuns antes mesmo de você executá-lo.

### Características Principais

- ✅ **Built-in**: Já vem instalado com o Go, não precisa instalar nada
- ✅ **Análise estática**: Examina o código sem executá-lo
- ✅ **Foco em bugs reais**: Detecta problemas que realmente causam erros
- ✅ **Executado automaticamente**: Roda automaticamente com `go test`
- ✅ **Múltiplos verificadores**: Possui vários analisadores especializados

### O Que o `go vet` Detecta?

O `go vet` verifica uma ampla gama de problemas comuns:

#### 1. Código Inalcançável (Unreachable Code)

```go
func exemplo() {
    return
    fmt.Println("Isso nunca será executado!") // ❌ go vet detecta
}
```

#### 2. Formatos Incorretos em Printf

```go
func exemplo() {
    nome := "João"
    fmt.Printf("Olá, %d\n", nome) // ❌ %d espera int, mas recebe string
}
```

#### 3. Erros em Struct Tags

```go
type Usuario struct {
    Nome string `json:"nome" xml:"nome"` // ✅ Correto
    Idade int `json:"idade" xml:"idade"` // ✅ Correto
}

type UsuarioErrado struct {
    Nome string `json:"nome" xml:"nome"` // ❌ go vet detecta tags malformadas
    Idade int `json:"idade" xml:"idade"` // ❌ go vet detecta tags malformadas
}
```

#### 4. Possíveis Nil Pointer Dereferences

```go
func exemplo() {
    var ptr *int
    fmt.Println(*ptr) // ❌ go vet pode detectar alguns casos
}
```

#### 5. Problemas com Goroutines

```go
func exemplo() {
    go func() {
        // ❌ go vet detecta se você esqueceu de usar variáveis capturadas
    }()
}
```

#### 6. Problemas com Range Loops

```go
func exemplo() {
    slice := []int{1, 2, 3}
    for i := range slice {
        slice[i] = i // ✅ Correto
    }
    
    for i, v := range slice {
        v = i // ❌ go vet detecta: você está modificando a cópia, não o original
    }
}
```

### Sintaxe Básica

```bash
# Analisar o pacote atual
go vet

# Analisar um pacote específico
go vet ./pacote

# Analisar recursivamente
go vet ./...

# Analisar um arquivo específico
go vet arquivo.go

# Com mais verbosidade
go vet -v

# Mostrar todos os analisadores disponíveis
go vet -help
```

### Exemplos Práticos

**Analisar o pacote atual:**

```bash
go vet
```

**Analisar todo o projeto:**

```bash
go vet ./...
```

**Analisar um pacote específico:**

```bash
go vet ./cmd/server
```

**Verificar o que será analisado:**

```bash
go vet -v ./...
```

### Integração com `go test`

O `go vet` é **automaticamente executado** quando você roda `go test`. Isso significa que seus testes não apenas verificam se o código funciona, mas também se há problemas de qualidade.

```bash
# go test executa go vet automaticamente
go test ./...

# Se quiser desabilitar (não recomendado)
go test -vet=off ./...
```

### Analisadores Específicos

O `go vet` possui vários analisadores que você pode usar individualmente:

```bash
# Verificar apenas problemas de printf
go vet -printf ./...

# Verificar apenas problemas com range loops
go vet -rangeloops ./...

# Verificar apenas problemas com struct tags
go vet -structtag ./...

# Verificar apenas problemas com nil pointers
go vet -nilfunc ./...
```

### Exemplo Completo: Detectando Problemas

Vamos criar um exemplo com vários problemas que o `go vet` detecta:

```go
package main

import "fmt"

func main() {
    // Problema 1: Código inalcançável
    return
    fmt.Println("Isso nunca será executado")
    
    // Problema 2: Formato incorreto em printf
    nome := "João"
    idade := 25
    fmt.Printf("Nome: %d, Idade: %s\n", nome, idade) // Tipos trocados!
    
    // Problema 3: Range loop problemático
    numeros := []int{1, 2, 3}
    for _, num := range numeros {
        num = num * 2 // Modificando cópia, não o original
    }
}
```

Ao executar `go vet` neste código:

```bash
go vet main.go
```

Você receberá avisos como:

```
main.go:8:2: unreachable code
main.go:12:2: fmt.Printf format %d has arg nome of wrong type string
main.go:12:2: fmt.Printf format %s has arg idade of wrong type int
main.go:17:4: range variable num is never modified; consider using &num
```

### Quando Usar?

- ✅ **Sempre antes de commitar código**
- ✅ Como parte do pipeline CI/CD
- ✅ Durante o desenvolvimento (antes de testar)
- ✅ Para revisar código legado
- ✅ Integrado com editores (VS Code, GoLand)

### Boas Práticas

1. **Execute regularmente**: Faça do `go vet` parte do seu workflow
2. **Integre com CI/CD**: Configure para rodar automaticamente
3. **Não ignore avisos**: Eles geralmente indicam bugs reais
4. **Use com `go test`**: Deixe que rode automaticamente nos testes

---

## 2. `goimports` - O Organizador de Imports

### O Que É?

O `goimports` é uma ferramenta que **gerencia automaticamente** as declarações de import no código Go. Ela adiciona imports faltantes, remove imports não utilizados e formata os imports de acordo com as convenções do Go.

### Características Principais

- ✅ **Automático**: Adiciona e remove imports automaticamente
- ✅ **Formatado**: Organiza imports em grupos (stdlib, terceiros, locais)
- ✅ **Inteligente**: Detecta quais imports são necessários
- ✅ **Integração com editores**: Funciona perfeitamente com VS Code, GoLand, etc.
- ✅ **Mais conveniente que `gofmt`**: `goimports` faz tudo que `gofmt` faz + gerencia imports

### Diferença: `gofmt` vs `goimports`

| `gofmt` | `goimports` |
|---------|-------------|
| Formata código | Formata código **+** gerencia imports |
| Não adiciona imports | Adiciona imports faltantes |
| Não remove imports | Remove imports não utilizados |
| Built-in | Precisa instalar |

**Resumo**: `goimports` é um `gofmt` melhorado que também gerencia imports!

### Instalação

O `goimports` não vem com o Go, mas é fácil de instalar:

```bash
# Instalar goimports
go install golang.org/x/tools/cmd/goimports@latest

# Verificar instalação
goimports -version
```

### Sintaxe Básica

```bash
# Formatar arquivo (modifica in-place)
goimports -w arquivo.go

# Formatar e mostrar diff (não modifica)
goimports -d arquivo.go

# Formatar múltiplos arquivos
goimports -w *.go

# Formatar recursivamente
goimports -w ./...

# Verificar se precisa formatação (útil para CI)
goimports -l .
```

### Flags Importantes

```bash
# -w: Escrever resultado no arquivo (modifica o arquivo)
goimports -w main.go

# -d: Mostrar diff (não modifica)
goimports -d main.go

# -l: Listar arquivos que precisam formatação
goimports -l .

# -local: Prefixo para imports locais (organização)
goimports -local "github.com/meu-usuario" main.go

# -format-only: Apenas formatar, não adicionar/remover imports
goimports -format-only main.go
```

### Como Funciona?

O `goimports` analisa seu código e:

1. **Detecta imports faltantes**: Se você usa `fmt.Println` mas não importou `fmt`, ele adiciona
2. **Remove imports não utilizados**: Se você importou algo mas não usa, ele remove
3. **Organiza imports em grupos**:
   - Grupo 1: Biblioteca padrão (stdlib)
   - Grupo 2: Imports de terceiros
   - Grupo 3: Imports locais (do seu projeto)

### Exemplo Prático: Antes e Depois

**Antes (código com problemas):**

```go
package main

import (
    "fmt"
    "os"  // ❌ Não está sendo usado
    "time"
)

// ❌ Falta importar "strings"
func main() {
    resultado := strings.ToUpper("hello")
    fmt.Println(resultado)
    time.Sleep(1 * time.Second)
}
```

**Depois de `goimports -w main.go`:**

```go
package main

import (
    "fmt"
    "strings"  // ✅ Adicionado automaticamente
    "time"
)

// ✅ "os" foi removido automaticamente (não estava sendo usado)
func main() {
    resultado := strings.ToUpper("hello")
    fmt.Println(resultado)
    time.Sleep(1 * time.Second)
}
```

### Organização de Imports

O `goimports` organiza os imports em grupos separados por linha em branco:

```go
package main

import (
    // Grupo 1: Biblioteca padrão (stdlib)
    "fmt"
    "os"
    "strings"
    "time"

    // Grupo 2: Imports de terceiros
    "github.com/gin-gonic/gin"
    "github.com/lib/pq"

    // Grupo 3: Imports locais (seu projeto)
    "meu-projeto/internal/models"
    "meu-projeto/pkg/utils"
)
```

### Integração com Editores

A maioria dos editores modernos suporta `goimports` automaticamente:

#### VS Code

Adicione ao `settings.json`:

```json
{
    "go.formatTool": "goimports",
    "editor.formatOnSave": true
}
```

#### GoLand

Vá em: `Settings > Tools > Actions on Save > Run goimports`

#### Vim/Neovim

Use plugins como `vim-go` que integram `goimports` automaticamente.

### Exemplo Completo: Workflow com goimports

```bash
# 1. Escrever código (sem se preocupar com imports)
# main.go
package main

func main() {
    fmt.Println("Hello")
    time.Sleep(1 * time.Second)
}

# 2. Executar goimports
goimports -w main.go

# 3. Código formatado e imports corrigidos automaticamente
# main.go (após goimports)
package main

import (
    "fmt"
    "time"
)

func main() {
    fmt.Println("Hello")
    time.Sleep(1 * time.Second)
}
```

### Quando Usar?

- ✅ **Sempre antes de commitar código**
- ✅ **Configurar no editor para executar ao salvar**
- ✅ Como parte do pipeline CI/CD
- ✅ Para limpar código legado
- ✅ Durante o desenvolvimento (deixe o editor fazer automaticamente)

### Boas Práticas

1. **Configure no editor**: Deixe executar automaticamente ao salvar
2. **Use `-w` com cuidado**: Sempre revise antes de commitar
3. **Combine com `go vet`**: Use ambas as ferramentas juntas
4. **Use em CI/CD**: Configure para verificar formatação

---

## 3. Trabalhando Juntas: `go vet` + `goimports`

### Workflow Recomendado

Um workflow profissional combina ambas as ferramentas:

```bash
# 1. Formatar código e organizar imports
goimports -w ./...

# 2. Verificar problemas de qualidade
go vet ./...

# 3. Executar testes (go vet roda automaticamente aqui também)
go test ./...

# 4. Se tudo passar, commitar
git add .
git commit -m "feat: nova funcionalidade"
```

### Script de Pre-commit

Você pode criar um script para garantir que o código está sempre formatado:

```bash
#!/bin/sh
# .git/hooks/pre-commit

# Formatar código
goimports -w .

# Verificar problemas
go vet ./...

# Se go vet encontrar problemas, abortar commit
if [ $? -ne 0 ]; then
    echo "❌ go vet encontrou problemas. Corrija antes de commitar."
    exit 1
fi

# Adicionar arquivos formatados
git add -u
```

### Integração com Makefile

```makefile
.PHONY: format vet test

# Formatar código
format:
	goimports -w .

# Verificar qualidade
vet:
	go vet ./...

# Testar
test:
	go test ./...

# Tudo junto
all: format vet test
```

### Exemplo Completo: Projeto Real

Vamos ver um exemplo completo de como usar essas ferramentas em um projeto:

```bash
# Estrutura do projeto
projeto/
├── main.go
├── utils.go
└── utils_test.go

# 1. Desenvolver código
# ... escrever código ...

# 2. Formatar e organizar imports
goimports -w .

# 3. Verificar qualidade
go vet ./...

# 4. Executar testes (go vet roda automaticamente)
go test -v ./...

# 5. Se tudo passar, build
go build -o app .
```

---

## 4. Comparação: Ferramentas de Qualidade

### `go vet` vs Linters Externos

| Ferramenta | Foco | Quando Usar |
|----------------|------|-------------|
| `go vet` | Bugs reais | Sempre, parte do workflow básico |
| `golangci-lint` | Estilo + bugs | Projetos maiores, equipes |
| `staticcheck` | Análise avançada | Projetos profissionais |
| `goimports` | Formatação + imports | Sempre, integrado no editor |

### Hierarquia de Ferramentas

```
Nível 1 (Essencial - Sempre usar):
├── go vet          ✅ Built-in, sempre disponível
└── goimports       ✅ Fácil de instalar, essencial

Nível 2 (Recomendado - Projetos maiores):
├── golangci-lint   ⭐ Análise completa
└── staticcheck     ⭐ Análise avançada
```

---

## Resumo dos Conceitos

| Ferramenta | Propósito | Quando Usar |
|------------|-----------|-------------|
| `go vet` | Detectar bugs potenciais | Sempre, antes de commitar |
| `goimports` | Gerenciar imports automaticamente | Sempre, integrado no editor |

---

## Conclusão

Dominar `go vet` e `goimports` é fundamental para:

1. **Qualidade**: Código livre de bugs comuns
2. **Consistência**: Imports sempre organizados e código formatado
3. **Produtividade**: Menos tempo corrigindo problemas, mais tempo desenvolvendo
4. **Profissionalismo**: Código que segue as melhores práticas da comunidade Go

Essas ferramentas são a base de um workflow profissional em Go. Na próxima parte desta aula, vamos simplificar esses conceitos com analogias do dia a dia para fixar melhor o aprendizado!

