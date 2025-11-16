# Módulo 35: Deployment & Tooling em Go
## Aula 3: Exercícios e Reflexão - Praticando Building e Cross-compilation

Olá! Agora é hora de colocar a mão na massa! Vamos praticar tudo que aprendemos sobre building executáveis e cross-compilation através de exercícios práticos e reflexões.

---

## 🎯 Exercícios Práticos

### Exercício 1: Primeiro Build

**Objetivo**: Compilar seu primeiro executável Go.

**Tarefa:**
1. Crie um arquivo `hello.go` com o seguinte conteúdo:

```go
package main

import "fmt"

func main() {
    fmt.Println("Olá! Este é meu primeiro executável Go!")
    fmt.Println("Eu fui compilado com go build!")
}
```

2. Compile o programa:
```bash
go build hello.go
```

3. Execute o executável gerado:
```bash
# Linux/Mac
./hello

# Windows
hello.exe
```

4. Verifique o tamanho do arquivo:
```bash
ls -lh hello    # Linux/Mac
dir hello.exe   # Windows
```

**Perguntas para reflexão:**
- Qual o tamanho do executável gerado?
- O executável funciona sem ter Go instalado? (Teste em outro computador ou container)
- O que acontece se você deletar o arquivo `hello.go` e tentar executar `hello`?

---

### Exercício 2: Build com Nome Personalizado

**Objetivo**: Criar executável com nome específico.

**Tarefa:**
1. Crie um programa `calculadora.go`:

```go
package main

import (
    "fmt"
    "os"
    "strconv"
)

func main() {
    if len(os.Args) < 3 {
        fmt.Println("Uso: calculadora <número1> <número2>")
        os.Exit(1)
    }

    a, _ := strconv.Atoi(os.Args[1])
    b, _ := strconv.Atoi(os.Args[2])

    fmt.Printf("%d + %d = %d\n", a, b, a+b)
    fmt.Printf("%d - %d = %d\n", a, b, a-b)
    fmt.Printf("%d * %d = %d\n", a, b, a*b)
    if b != 0 {
        fmt.Printf("%d / %d = %.2f\n", a, b, float64(a)/float64(b))
    }
}
```

2. Compile com nome personalizado:
```bash
go build -o calc calculadora.go
```

3. Teste o executável:
```bash
./calc 10 5
# ou
calc.exe 10 5
```

**Perguntas para reflexão:**
- Por que é útil ter um nome personalizado para o executável?
- O que acontece se você não especificar `-o`?

---

### Exercício 3: Build com Informações de Versão

**Objetivo**: Injetar informações de build no executável.

**Tarefa:**
1. Crie um programa `app.go`:

```go
package main

import "fmt"

var (
    Version   = "dev"
    BuildTime = "unknown"
    GitCommit = "unknown"
)

func main() {
    fmt.Printf("Aplicação Versão: %s\n", Version)
    fmt.Printf("Build Time: %s\n", BuildTime)
    fmt.Printf("Git Commit: %s\n", GitCommit)
}
```

2. Compile sem informações:
```bash
go build -o app app.go
./app
```

3. Compile com informações:
```bash
# Linux/Mac
go build -ldflags "-X main.Version=1.0.0 -X main.BuildTime=$(date) -X main.GitCommit=abc123" -o app app.go

# Windows (PowerShell)
$env:VERSION="1.0.0"
$env:BUILDTIME=(Get-Date)
$env:COMMIT="abc123"
go build -ldflags "-X main.Version=$env:VERSION -X main.BuildTime=$env:BUILDTIME -X main.GitCommit=$env:COMMIT" -o app.exe app.go
```

4. Execute e compare as saídas.

**Perguntas para reflexão:**
- Por que é importante ter informações de versão no executável?
- Como isso ajuda em produção quando há problemas?
- Como você poderia automatizar isso em um script?

---

### Exercício 4: Otimização de Tamanho

**Objetivo**: Reduzir o tamanho do executável.

**Tarefa:**
1. Use o programa do exercício anterior (`app.go`).

2. Compile normalmente e meça o tamanho:
```bash
go build -o app-normal app.go
ls -lh app-normal
```

3. Compile otimizado e compare:
```bash
go build -ldflags "-s -w" -o app-otimizado app.go
ls -lh app-otimizado
```

4. Compare os tamanhos:
```bash
# Linux/Mac
echo "Normal: $(stat -f%z app-normal) bytes"
echo "Otimizado: $(stat -f%z app-otimizado) bytes"

# Ou simplesmente
ls -lh app-*
```

**Perguntas para reflexão:**
- Qual a diferença de tamanho?
- Quando você deveria usar `-s -w`?
- Quando você NÃO deveria usar `-s -w`?

---

### Exercício 5: Primeira Cross-compilation

**Objetivo**: Compilar para outra plataforma.

**Tarefa:**
1. Crie um programa simples `platform.go`:

```go
package main

import (
    "fmt"
    "runtime"
)

func main() {
    fmt.Printf("Sistema Operacional: %s\n", runtime.GOOS)
    fmt.Printf("Arquitetura: %s\n", runtime.GOARCH)
    fmt.Println("Este programa foi compilado para esta plataforma!")
}
```

2. Compile para sua plataforma atual:
```bash
go build -o platform-local platform.go
./platform-local
```

3. Compile para Linux (se você estiver no Mac/Windows):
```bash
GOOS=linux GOARCH=amd64 go build -o platform-linux platform.go
```

4. Compile para Windows (se você estiver no Mac/Linux):
```bash
GOOS=windows GOARCH=amd64 go build -o platform-windows.exe platform.go
```

5. Verifique os arquivos gerados:
```bash
ls -lh platform-*
```

**Perguntas para reflexão:**
- Você consegue executar o binário Linux no seu sistema? (Provavelmente não, a menos que use WSL ou container)
- Como você testaria se o binário cross-compilado funciona corretamente?
- Por que é útil poder compilar para outras plataformas?

---

### Exercício 6: Build Multiplataforma com Script

**Objetivo**: Criar script que compila para todas as plataformas.

**Tarefa:**
1. Crie um programa `greeter.go`:

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    name := "Mundo"
    if len(os.Args) > 1 {
        name = os.Args[1]
    }
    fmt.Printf("Olá, %s!\n", name)
}
```

2. Crie um script `build-all.sh` (Linux/Mac) ou `build-all.ps1` (Windows):

**Linux/Mac (`build-all.sh`):**
```bash
#!/bin/bash

APP_NAME="greeter"
mkdir -p dist

echo "Building for Linux amd64..."
GOOS=linux GOARCH=amd64 go build -o dist/${APP_NAME}-linux-amd64 greeter.go

echo "Building for macOS amd64..."
GOOS=darwin GOARCH=amd64 go build -o dist/${APP_NAME}-darwin-amd64 greeter.go

echo "Building for macOS arm64..."
GOOS=darwin GOARCH=arm64 go build -o dist/${APP_NAME}-darwin-arm64 greeter.go

echo "Building for Windows amd64..."
GOOS=windows GOARCH=amd64 go build -o dist/${APP_NAME}-windows-amd64.exe greeter.go

echo "Done! Binaries in ./dist/"
ls -lh dist/
```

**Windows (`build-all.ps1`):**
```powershell
$APP_NAME = "greeter"
New-Item -ItemType Directory -Force -Path dist | Out-Null

Write-Host "Building for Linux amd64..."
$env:GOOS = "linux"
$env:GOARCH = "amd64"
go build -o "dist/${APP_NAME}-linux-amd64" greeter.go

Write-Host "Building for macOS amd64..."
$env:GOOS = "darwin"
$env:GOARCH = "amd64"
go build -o "dist/${APP_NAME}-darwin-amd64" greeter.go

Write-Host "Building for macOS arm64..."
$env:GOOS = "darwin"
$env:GOARCH = "arm64"
go build -o "dist/${APP_NAME}-darwin-arm64" greeter.go

Write-Host "Building for Windows amd64..."
$env:GOOS = "windows"
$env:GOARCH = "amd64"
go build -o "dist/${APP_NAME}-windows-amd64.exe" greeter.go

Write-Host "Done! Binaries in ./dist/"
Get-ChildItem dist/
```

3. Execute o script:
```bash
# Linux/Mac
chmod +x build-all.sh
./build-all.sh

# Windows
.\build-all.ps1
```

4. Verifique os arquivos gerados em `dist/`.

**Perguntas para reflexão:**
- Quantos executáveis foram criados?
- Como você distribuiria esses executáveis para usuários?
- Como você poderia melhorar este script?

---

### Exercício 7: Makefile para Builds

**Objetivo**: Criar Makefile para automatizar builds.

**Tarefa:**
1. Crie um `Makefile`:

```makefile
APP_NAME := myapp
VERSION := $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
BUILD_TIME := $(shell date -u +%Y-%m-%dT%H:%M:%SZ)
GIT_COMMIT := $(shell git rev-parse --short HEAD 2>/dev/null || echo "unknown")

LDFLAGS := -ldflags "-X main.Version=$(VERSION) -X main.BuildTime=$(BUILD_TIME) -X main.GitCommit=$(GIT_COMMIT) -s -w"

.PHONY: build build-all clean test run help

help:
	@echo "Available targets:"
	@echo "  build      - Build for current platform"
	@echo "  build-all  - Build for all platforms"
	@echo "  clean      - Remove build artifacts"
	@echo "  test       - Run tests"
	@echo "  run        - Run the application"

build:
	@echo "Building $(APP_NAME)..."
	go build $(LDFLAGS) -o bin/$(APP_NAME) .

build-all: clean
	@echo "Building for all platforms..."
	@mkdir -p dist
	@echo "  Linux amd64..."
	@GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o dist/$(APP_NAME)-linux-amd64 .
	@echo "  Linux arm64..."
	@GOOS=linux GOARCH=arm64 go build $(LDFLAGS) -o dist/$(APP_NAME)-linux-arm64 .
	@echo "  macOS amd64..."
	@GOOS=darwin GOARCH=amd64 go build $(LDFLAGS) -o dist/$(APP_NAME)-darwin-amd64 .
	@echo "  macOS arm64..."
	@GOOS=darwin GOARCH=arm64 go build $(LDFLAGS) -o dist/$(APP_NAME)-darwin-arm64 .
	@echo "  Windows amd64..."
	@GOOS=windows GOARCH=amd64 go build $(LDFLAGS) -o dist/$(APP_NAME)-windows-amd64.exe .
	@echo "Done! Files in ./dist/"
	@ls -lh dist/

clean:
	@echo "Cleaning..."
	rm -rf bin/ dist/

test:
	go test ./...

run:
	go run .
```

2. Crie um programa `main.go`:

```go
package main

import "fmt"

var (
    Version   = "dev"
    BuildTime = "unknown"
    GitCommit = "unknown"
)

func main() {
    fmt.Printf("App: %s\n", "myapp")
    fmt.Printf("Version: %s\n", Version)
    fmt.Printf("Build Time: %s\n", BuildTime)
    fmt.Printf("Git Commit: %s\n", GitCommit)
}
```

3. Use o Makefile:
```bash
make build      # Build local
make build-all  # Build para todas as plataformas
make clean      # Limpar
make help       # Ver ajuda
```

**Perguntas para reflexão:**
- Quais são as vantagens de usar Makefile?
- Como você adicionaria mais targets ao Makefile?
- Como isso se integra com CI/CD?

---

## 🤔 Questões para Reflexão

### 1. Por Que Binários Standalone São Importantes?

Pense em cenários onde você precisa:
- Distribuir uma ferramenta para clientes
- Fazer deploy em servidores sem Go instalado
- Criar aplicações que rodam em containers minimalistas

**Reflita:**
- Como isso simplifica o deployment?
- Quais problemas isso resolve?
- Quando você preferiria uma linguagem interpretada?

---

### 2. Cross-compilation vs Máquinas Virtuais

**Cenário**: Você precisa criar binários para Linux, Windows e Mac.

**Opções:**
- A) Usar cross-compilation (uma máquina, múltiplos binários)
- B) Usar VMs/containers (uma máquina por plataforma)

**Reflita:**
- Quais são os prós e contras de cada abordagem?
- Quando cada uma é melhor?
- Como isso afeta CI/CD?

---

### 3. Tamanho dos Binários Go

Binários Go são relativamente grandes (comparados a C, mas menores que Java/.NET).

**Reflita:**
- Quando o tamanho do binário importa?
- Quais técnicas você usaria para reduzir o tamanho?
- Quando você aceitaria um binário maior em troca de outras vantagens?

---

### 4. Build Tags e Condicional Compilation

Go suporta build tags para compilar código condicionalmente:

```go
// +build linux

package main
```

**Reflita:**
- Quando você usaria build tags?
- Como isso se relaciona com cross-compilation?
- Quais são os trade-offs?

---

### 5. CI/CD e Builds Automatizados

**Cenário**: Toda vez que você faz push, quer gerar binários para todas as plataformas.

**Reflita:**
- Como você configuraria isso no GitHub Actions?
- Como você versionaria os binários?
- Como você distribuiria os binários (GitHub Releases, S3, etc.)?

---

## 🎓 Desafios Avançados

### Desafio 1: Build com Múltiplos Comandos

Crie uma estrutura de projeto com múltiplos comandos:

```
projeto/
├── cmd/
│   ├── server/
│   │   └── main.go
│   └── cli/
│       └── main.go
├── internal/
│   └── shared/
│       └── utils.go
└── go.mod
```

Crie um Makefile que:
- Compila ambos os comandos
- Cria binários para todas as plataformas
- Inclui informações de versão

---

### Desafio 2: Script de Release Automatizado

Crie um script que:
1. Executa testes
2. Gera binários para todas as plataformas
3. Cria checksums (SHA256) de cada binário
4. Cria um arquivo de release notes
5. (Opcional) Faz upload para GitHub Releases

---

### Desafio 3: Build Otimizado com Análise

Crie um script que:
1. Compila com diferentes otimizações
2. Compara tamanhos dos binários
3. Executa benchmarks de performance
4. Gera relatório comparativo

---

## 📝 Checklist de Aprendizado

Marque o que você conseguiu fazer:

- [ ] Compilei meu primeiro executável Go
- [ ] Criei executável com nome personalizado
- [ ] Incluí informações de versão no executável
- [ ] Reduzi o tamanho do binário com `-s -w`
- [ ] Fiz cross-compilation para outra plataforma
- [ ] Criei script que compila para múltiplas plataformas
- [ ] Criei e usei um Makefile
- [ ] Entendi quando usar cada técnica
- [ ] Testei binários cross-compilados

---

## 🎯 Próximos Passos

Depois de completar os exercícios:

1. ✅ Experimente criar um projeto real e compilá-lo
2. ✅ Configure um pipeline CI/CD que faz builds automáticos
3. ✅ Crie uma ferramenta CLI e distribua para diferentes plataformas
4. ✅ Explore outras flags do `go build` (use `go help build`)

Na próxima parte, vamos ver boas práticas e considerações de performance!

