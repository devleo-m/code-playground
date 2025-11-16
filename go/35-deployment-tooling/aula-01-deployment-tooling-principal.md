# Módulo 35: Deployment & Tooling em Go
## Aula 1: Building Executables e Cross-compilation - Criando Binários para Produção

Olá! Bem-vindo a este módulo essencial sobre **deployment e tooling** em Go. Até agora você aprendeu a escrever código Go funcional, mas em projetos profissionais, é crucial saber **compilar executáveis**, **gerar binários para diferentes plataformas** e **preparar aplicações para produção**.

Nesta aula, vamos mergulhar em dois tópicos fundamentais:
1. **Building Executables**: Compilar código Go em binários nativos standalone
2. **Cross-compilation**: Criar executáveis para diferentes sistemas operacionais e arquiteturas

Essas habilidades são essenciais para criar aplicações Go que podem ser facilmente distribuídas e implantadas em qualquer ambiente.

---

## 1. Building Executables - Compilando Código Go

### O Que É `go build`?

O comando **`go build`** é a ferramenta principal do Go para compilar código fonte em **executáveis nativos**. Diferente de linguagens interpretadas (como Python ou JavaScript), Go compila seu código em um **binário standalone** que contém tudo necessário para executar, incluindo todas as dependências.

### Por Que Usar `go build`?

**Vantagens:**
- ✅ **Binário standalone**: Um único arquivo executável, sem dependências externas
- ✅ **Sem necessidade de Go instalado**: O binário roda em qualquer sistema compatível, sem precisar do Go instalado
- ✅ **Linking estático**: Todas as dependências são incluídas no binário
- ✅ **Performance**: Código compilado é muito mais rápido que código interpretado
- ✅ **Distribuição simples**: Basta copiar o arquivo e executar
- ✅ **Segurança**: Não precisa de runtime externo ou dependências que podem ter vulnerabilidades

**Quando usar:**
- Para criar aplicações de produção
- Para distribuir ferramentas CLI
- Para criar serviços que rodam em servidores
- Para criar aplicações desktop
- Para criar binários para diferentes plataformas

### Como Funciona?

O `go build` funciona em várias etapas:

1. **Análise**: Lê todos os arquivos `.go` no pacote
2. **Compilação**: Compila o código Go para código de máquina
3. **Linking**: Liga todas as dependências (incluindo biblioteca padrão)
4. **Geração**: Cria um arquivo executável binário

### Uso Básico do `go build`

#### Compilar um Programa Simples

```go
// main.go
package main

import "fmt"

func main() {
    fmt.Println("Olá, Mundo!")
}
```

```bash
# Compilar
go build main.go

# Isso cria um executável chamado 'main' (Linux/Mac) ou 'main.exe' (Windows)
# Execute:
./main        # Linux/Mac
# ou
main.exe      # Windows
```

#### Compilar um Módulo Completo

```bash
# Se você está no diretório raiz do módulo
go build

# Isso cria um executável com o nome do diretório
# Exemplo: se o diretório é 'myapp', cria 'myapp' ou 'myapp.exe'
```

#### Especificar Nome do Executável

```bash
# Compilar e especificar o nome do executável
go build -o minha-aplicacao main.go

# Linux/Mac: cria 'minha-aplicacao'
# Windows: cria 'minha-aplicacao.exe'
```

### Flags Úteis do `go build`

#### 1. `-o` - Especificar Nome do Output

```bash
go build -o bin/myapp main.go
```

#### 2. `-v` - Modo Verboso

Mostra os pacotes sendo compilados:

```bash
go build -v
```

**Saída:**
```
main
fmt
runtime
...
```

#### 3. `-x` - Mostrar Comandos Executados

Mostra todos os comandos que o build executa:

```bash
go build -x
```

**Saída:**
```
WORK=/tmp/go-build123456789
mkdir -p $WORK/b001/
cd /path/to/project
/usr/local/go/pkg/tool/linux_amd64/compile ...
...
```

#### 4. `-a` - Recompilar Tudo

Força recompilação de todos os pacotes, mesmo os que não mudaram:

```bash
go build -a
```

#### 5. `-i` - Instalar Dependências

Instala dependências no cache antes de compilar:

```bash
go build -i
```

#### 6. `-race` - Habilitar Race Detector

Compila com race detector habilitado (útil para debugging):

```bash
go build -race
```

**Nota**: Binários com `-race` são muito mais lentos e não devem ser usados em produção.

#### 7. `-tags` - Build Tags

Compila com tags específicas:

```bash
go build -tags dev
```

#### 8. `-ldflags` - Flags do Linker

Passa flags para o linker. Útil para injetar informações de versão:

```bash
go build -ldflags "-X main.Version=1.0.0 -X main.BuildTime=$(date)"
```

### Exemplo Prático: Aplicação com Versão

Vamos criar uma aplicação que mostra sua versão:

```go
// main.go
package main

import (
    "fmt"
    "os"
)

var (
    Version   = "dev"
    BuildTime = "unknown"
    GitCommit = "unknown"
)

func main() {
    if len(os.Args) > 1 && os.Args[1] == "version" {
        fmt.Printf("Versão: %s\n", Version)
        fmt.Printf("Build Time: %s\n", BuildTime)
        fmt.Printf("Git Commit: %s\n", GitCommit)
        return
    }

    fmt.Println("Minha Aplicação")
    fmt.Println("Use 'version' para ver informações de build")
}
```

**Compilar com informações de versão:**

```bash
go build -ldflags "-X main.Version=1.0.0 -X main.BuildTime=$(date -u +%Y-%m-%dT%H:%M:%SZ) -X main.GitCommit=$(git rev-parse --short HEAD)" -o myapp
```

**Windows (PowerShell):**
```powershell
$env:VERSION="1.0.0"
$env:BUILDTIME=(Get-Date -Format "yyyy-MM-ddTHH:mm:ssZ")
$env:COMMIT=(git rev-parse --short HEAD)
go build -ldflags "-X main.Version=$env:VERSION -X main.BuildTime=$env:BUILDTIME -X main.GitCommit=$env:COMMIT" -o myapp.exe
```

### Exemplo Prático: Build com Múltiplos Arquivos

```go
// main.go
package main

import "fmt"

func main() {
    fmt.Println("Aplicação principal")
    helper()
}
```

```go
// helper.go
package main

import "fmt"

func helper() {
    fmt.Println("Função helper")
}
```

```bash
# Compilar todos os arquivos .go no diretório
go build

# Ou especificar arquivos
go build main.go helper.go
```

### Exemplo Prático: Build de Módulo com Dependências

```go
// go.mod
module myapp

go 1.21

require (
    github.com/gorilla/mux v1.8.1
)
```

```go
// main.go
package main

import (
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
)

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Olá!")
    })
    
    fmt.Println("Servidor rodando em :8080")
    http.ListenAndServe(":8080", r)
}
```

```bash
# Baixar dependências
go mod download

# Compilar (dependências são incluídas automaticamente)
go build -o myapp

# O binário inclui todas as dependências!
```

### Otimizações de Build

#### 1. Remover Informações de Debug

```bash
# Remove símbolos de debug (reduz tamanho do binário)
go build -ldflags "-s -w" -o myapp
```

- `-s`: Remove tabela de símbolos e informações de debug
- `-w`: Remove informações de debug DWARF

**Resultado**: Binário menor, mas sem informações de debug.

#### 2. Build Otimizado para Produção

```bash
# Build otimizado (padrão já é otimizado, mas pode forçar)
go build -ldflags "-s -w" -trimpath -o myapp
```

- `-trimpath`: Remove caminhos do sistema de arquivos do binário

#### 3. Build com CGO Desabilitado

Se você não precisa de CGO, desabilite para builds mais rápidos e binários menores:

```bash
CGO_ENABLED=0 go build -o myapp
```

### Tamanho dos Binários

Go gera binários relativamente grandes porque inclui o runtime e garbage collector. Para reduzir:

```bash
# 1. Remover símbolos de debug
go build -ldflags "-s -w" -o myapp

# 2. Desabilitar CGO (se não precisar)
CGO_ENABLED=0 go build -o myapp

# 3. Usar UPX para comprimir (opcional, pode causar problemas)
upx myapp
```

### Build para Testes

```bash
# Compilar sem executar
go build

# Compilar e executar testes
go test ./...

# Compilar e executar
go run main.go
```

### Estrutura de Diretórios para Builds

```
projeto/
├── cmd/
│   ├── server/
│   │   └── main.go
│   └── cli/
│       └── main.go
├── internal/
│   └── ...
├── go.mod
└── Makefile
```

**Compilar cada comando:**

```bash
# Compilar servidor
go build -o bin/server ./cmd/server

# Compilar CLI
go build -o bin/cli ./cmd/cli
```

### Makefile para Builds

```makefile
# Makefile
.PHONY: build clean test

VERSION := $(shell git describe --tags --always --dirty)
BUILD_TIME := $(shell date -u +%Y-%m-%dT%H:%M:%SZ)
GIT_COMMIT := $(shell git rev-parse --short HEAD)

LDFLAGS := -ldflags "-X main.Version=$(VERSION) -X main.BuildTime=$(BUILD_TIME) -X main.GitCommit=$(GIT_COMMIT) -s -w"

build:
	go build $(LDFLAGS) -o bin/myapp ./cmd/server

build-all:
	GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o bin/myapp-linux-amd64 ./cmd/server
	GOOS=darwin GOARCH=amd64 go build $(LDFLAGS) -o bin/myapp-darwin-amd64 ./cmd/server
	GOOS=windows GOARCH=amd64 go build $(LDFLAGS) -o bin/myapp-windows-amd64.exe ./cmd/server

clean:
	rm -rf bin/

test:
	go test ./...

run:
	go run ./cmd/server
```

**Uso:**
```bash
make build        # Build local
make build-all    # Build para todas as plataformas
make clean        # Limpar binários
make test         # Executar testes
```

---

## 2. Cross-compilation - Compilando para Outras Plataformas

### O Que É Cross-compilation?

**Cross-compilation** é a capacidade de compilar código para uma plataforma diferente daquela em que você está desenvolvendo. Por exemplo, compilar um binário Linux enquanto você está em um Mac, ou compilar para Windows enquanto você está em Linux.

### Por Que Usar Cross-compilation?

**Vantagens:**
- ✅ **Desenvolvimento único**: Desenvolva em uma plataforma, compile para todas
- ✅ **CI/CD simplificado**: Um único servidor de build pode gerar binários para todas as plataformas
- ✅ **Sem necessidade de VMs**: Não precisa de máquinas virtuais ou containers para cada plataforma
- ✅ **Distribuição fácil**: Gere binários para todos os seus usuários de uma vez
- ✅ **Testes**: Teste em diferentes plataformas sem mudar de máquina

**Quando usar:**
- Quando você precisa distribuir para múltiplas plataformas
- Em pipelines CI/CD
- Quando desenvolvedores usam diferentes sistemas operacionais
- Para criar ferramentas CLI multiplataforma
- Para criar aplicações que rodam em servidores Linux mas você desenvolve em Mac/Windows

### Como Funciona?

Go suporta cross-compilation nativamente através das variáveis de ambiente `GOOS` e `GOARCH`:

- **GOOS**: Sistema operacional de destino (linux, darwin, windows, freebsd, etc.)
- **GOARCH**: Arquitetura de destino (amd64, arm64, 386, arm, etc.)

### Plataformas Suportadas

#### Sistemas Operacionais (GOOS)

- `linux` - Linux
- `darwin` - macOS
- `windows` - Windows
- `freebsd` - FreeBSD
- `openbsd` - OpenBSD
- `netbsd` - NetBSD
- `android` - Android
- `ios` - iOS
- `js` - JavaScript (WebAssembly)
- `plan9` - Plan 9

#### Arquiteturas (GOARCH)

- `amd64` - x86-64 (64-bit)
- `386` - x86 (32-bit)
- `arm` - ARM (32-bit)
- `arm64` - ARM64 (64-bit)
- `mips` - MIPS
- `mips64` - MIPS64
- `ppc64` - PowerPC 64-bit
- `ppc64le` - PowerPC 64-bit little-endian
- `riscv64` - RISC-V 64-bit
- `s390x` - IBM z/Architecture
- `wasm` - WebAssembly

### Uso Básico de Cross-compilation

#### Compilar para Linux (de Mac/Windows)

```bash
# Compilar para Linux amd64
GOOS=linux GOARCH=amd64 go build -o myapp-linux main.go

# Execute no Linux:
./myapp-linux
```

#### Compilar para Windows (de Mac/Linux)

```bash
# Compilar para Windows amd64
GOOS=windows GOARCH=amd64 go build -o myapp-windows.exe main.go

# Execute no Windows:
myapp-windows.exe
```

#### Compilar para macOS (de Linux/Windows)

```bash
# Compilar para macOS amd64
GOOS=darwin GOARCH=amd64 go build -o myapp-darwin main.go

# Compilar para macOS ARM (Apple Silicon)
GOOS=darwin GOARCH=arm64 go build -o myapp-darwin-arm64 main.go
```

### Exemplo Prático: Script de Build Multiplataforma

```bash
#!/bin/bash
# build-all.sh

APP_NAME="myapp"
VERSION=$(git describe --tags --always --dirty)
BUILD_TIME=$(date -u +%Y-%m-%dT%H:%M:%SZ)
GIT_COMMIT=$(git rev-parse --short HEAD)

LDFLAGS="-X main.Version=${VERSION} -X main.BuildTime=${BUILD_TIME} -X main.GitCommit=${GIT_COMMIT} -s -w"

# Criar diretório de output
mkdir -p dist

# Linux
echo "Building for Linux amd64..."
GOOS=linux GOARCH=amd64 go build -ldflags "${LDFLAGS}" -o dist/${APP_NAME}-linux-amd64

# Linux ARM64
echo "Building for Linux arm64..."
GOOS=linux GOARCH=arm64 go build -ldflags "${LDFLAGS}" -o dist/${APP_NAME}-linux-arm64

# macOS Intel
echo "Building for macOS amd64..."
GOOS=darwin GOARCH=amd64 go build -ldflags "${LDFLAGS}" -o dist/${APP_NAME}-darwin-amd64

# macOS Apple Silicon
echo "Building for macOS arm64..."
GOOS=darwin GOARCH=arm64 go build -ldflags "${LDFLAGS}" -o dist/${APP_NAME}-darwin-arm64

# Windows
echo "Building for Windows amd64..."
GOOS=windows GOARCH=amd64 go build -ldflags "${LDFLAGS}" -o dist/${APP_NAME}-windows-amd64.exe

echo "Builds completos! Arquivos em ./dist/"
ls -lh dist/
```

**Tornar executável e usar:**
```bash
chmod +x build-all.sh
./build-all.sh
```

### Exemplo Prático: Build com Makefile

```makefile
# Makefile
APP_NAME := myapp
VERSION := $(shell git describe --tags --always --dirty)
BUILD_TIME := $(shell date -u +%Y-%m-%dT%H:%M:%SZ)
GIT_COMMIT := $(shell git rev-parse --short HEAD)

LDFLAGS := -ldflags "-X main.Version=$(VERSION) -X main.BuildTime=$(BUILD_TIME) -X main.GitCommit=$(GIT_COMMIT) -s -w"

.PHONY: build build-all clean

# Build local
build:
	go build $(LDFLAGS) -o bin/$(APP_NAME)

# Build para todas as plataformas
build-all: clean
	@echo "Building for Linux amd64..."
	@GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o dist/$(APP_NAME)-linux-amd64
	@echo "Building for Linux arm64..."
	@GOOS=linux GOARCH=arm64 go build $(LDFLAGS) -o dist/$(APP_NAME)-linux-arm64
	@echo "Building for macOS amd64..."
	@GOOS=darwin GOARCH=amd64 go build $(LDFLAGS) -o dist/$(APP_NAME)-darwin-amd64
	@echo "Building for macOS arm64..."
	@GOOS=darwin GOARCH=arm64 go build $(LDFLAGS) -o dist/$(APP_NAME)-darwin-arm64
	@echo "Building for Windows amd64..."
	@GOOS=windows GOARCH=amd64 go build $(LDFLAGS) -o dist/$(APP_NAME)-windows-amd64.exe
	@echo "Done! Files in ./dist/"
	@ls -lh dist/

clean:
	rm -rf bin/ dist/
```

**Uso:**
```bash
make build      # Build local
make build-all  # Build para todas as plataformas
make clean      # Limpar
```

### Exemplo Prático: Build para Raspberry Pi

```bash
# Raspberry Pi 3/4 (ARM64)
GOOS=linux GOARCH=arm64 go build -o myapp-rpi main.go

# Raspberry Pi Zero/1 (ARM)
GOOS=linux GOARCH=arm go build -o myapp-rpi-zero main.go
```

### Exemplo Prático: Build para Android

```bash
# Android ARM64
GOOS=android GOARCH=arm64 go build -o myapp-android main.go

# Android ARM
GOOS=android GOARCH=arm go build -o myapp-android-arm main.go
```

### Limitações da Cross-compilation

#### 1. CGO

Se seu código usa **CGO** (chama código C), a cross-compilation pode ser limitada ou requerer toolchains específicas:

```bash
# CGO desabilitado (mais fácil para cross-compilation)
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o myapp

# Com CGO, você precisa do toolchain da plataforma de destino
```

**Solução**: Se possível, evite CGO ou use build tags para código específico de plataforma.

#### 2. Dependências de Sistema

Algumas dependências podem ter código específico de plataforma. Teste sempre os binários na plataforma de destino.

#### 3. Testes

Testes que dependem de recursos específicos da plataforma podem falhar durante cross-compilation.

### Verificando Plataformas Disponíveis

```bash
# Ver todas as plataformas suportadas
go tool dist list
```

**Saída:**
```
android/386
android/amd64
android/arm
android/arm64
darwin/amd64
darwin/arm64
...
```

### Script para Build de Múltiplas Plataformas

```go
// build.go - Script em Go para builds multiplataforma
package main

import (
    "fmt"
    "os"
    "os/exec"
    "path/filepath"
)

var platforms = []struct {
    os   string
    arch string
}{
    {"linux", "amd64"},
    {"linux", "arm64"},
    {"darwin", "amd64"},
    {"darwin", "arm64"},
    {"windows", "amd64"},
}

func main() {
    appName := "myapp"
    if len(os.Args) > 1 {
        appName = os.Args[1]
    }

    os.MkdirAll("dist", 0755)

    for _, p := range platforms {
        output := fmt.Sprintf("dist/%s-%s-%s", appName, p.os, p.arch)
        if p.os == "windows" {
            output += ".exe"
        }

        cmd := exec.Command("go", "build",
            "-ldflags", "-s -w",
            "-o", output,
            "main.go",
        )
        cmd.Env = append(os.Environ(),
            "GOOS="+p.os,
            "GOARCH="+p.arch,
        )

        fmt.Printf("Building %s/%s...\n", p.os, p.arch)
        if err := cmd.Run(); err != nil {
            fmt.Printf("Error building %s/%s: %v\n", p.os, p.arch, err)
            continue
        }

        fmt.Printf("  -> %s\n", output)
    }

    fmt.Println("\nDone!")
}
```

**Uso:**
```bash
go run build.go
# ou
go run build.go minha-app
```

### Integração com CI/CD

#### GitHub Actions

```yaml
# .github/workflows/build.yml
name: Build

on:
  push:
    tags:
      - 'v*'

jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      
      - name: Set up Go
        uses: actions/setup-go@v4
        with:
          go-version: '1.21'
      
      - name: Build Linux
        run: GOOS=linux GOARCH=amd64 go build -o dist/myapp-linux-amd64
      
      - name: Build macOS
        run: |
          GOOS=darwin GOARCH=amd64 go build -o dist/myapp-darwin-amd64
          GOOS=darwin GOARCH=arm64 go build -o dist/myapp-darwin-arm64
      
      - name: Build Windows
        run: GOOS=windows GOARCH=amd64 go build -o dist/myapp-windows-amd64.exe
      
      - name: Upload artifacts
        uses: actions/upload-artifact@v3
        with:
          name: binaries
          path: dist/
```

### Boas Práticas

#### 1. Sempre Testar Binários Cross-compilados

```bash
# Compile
GOOS=linux GOARCH=amd64 go build -o myapp

# Teste em container Docker
docker run --rm -v $(pwd):/app -w /app alpine:latest ./myapp
```

#### 2. Usar CGO_ENABLED=0 Quando Possível

```bash
# Mais confiável para cross-compilation
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o myapp
```

#### 3. Versionar Binários

```bash
VERSION=$(git describe --tags)
GOOS=linux GOARCH=amd64 go build \
  -ldflags "-X main.Version=$VERSION" \
  -o myapp-v${VERSION}-linux-amd64
```

#### 4. Assinar Binários (Segurança)

Para distribuição pública, considere assinar binários:

```bash
# Build
go build -o myapp

# Assinar (macOS)
codesign --sign "Developer ID Application: Your Name" myapp

# Verificar
codesign --verify myapp
```

---

## Resumo dos Conceitos

| Conceito | Descrição | Quando Usar |
|----------|-----------|-------------|
| **`go build`** | Compila código Go em executável nativo | Sempre que precisar de binário |
| **`-o`** | Especifica nome do output | Para controlar nome do executável |
| **`-ldflags`** | Passa flags para o linker | Para injetar versão, build time, etc. |
| **`-s -w`** | Remove símbolos de debug | Para binários menores em produção |
| **`GOOS`** | Sistema operacional de destino | Cross-compilation |
| **`GOARCH`** | Arquitetura de destino | Cross-compilation |
| **`CGO_ENABLED=0`** | Desabilita CGO | Para cross-compilation mais fácil |

---

## Conclusão

Dominar `go build` e cross-compilation é essencial para:

1. **Criar aplicações de produção**: Gerar binários otimizados e prontos para deploy
2. **Distribuição multiplataforma**: Criar binários para todos os seus usuários
3. **CI/CD eficiente**: Automatizar builds para múltiplas plataformas
4. **Deployment simplificado**: Um único binário, sem dependências
5. **Performance**: Código compilado nativo é extremamente rápido

Essas ferramentas são fundamentais para qualquer desenvolvedor Go profissional. Integre-as no seu workflow e veja como o deployment se torna muito mais simples!

Na próxima parte desta aula, vamos simplificar esses conceitos com analogias do dia a dia para fixar melhor o aprendizado!

