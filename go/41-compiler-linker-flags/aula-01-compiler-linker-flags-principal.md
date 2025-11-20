# Módulo 41: Compiler & Linker Flags em Go

## Aula 1: Compiler & Linker Flags - Controlando a Compilação

Olá! Bem-vindo ao módulo 41 sobre **Compiler & Linker Flags** em Go. Este módulo ensina como usar flags do compilador e linker para controlar como seu código é compilado, otimizado e linkado.

Flags permitem personalizar o processo de build, desde otimizações de performance até informações de debug, controle de tamanho de binário, e muito mais.

Nesta aula, vamos explorar:
1. **Compiler Flags (-gcflags)**: Controle do compilador
2. **Linker Flags (-ldflags)**: Controle do linker
3. **Build Tags (-tags)**: Tags customizadas
4. **Race Detector (-race)**: Detecção de race conditions
5. **Otimizações**: Reduzir tamanho, melhorar performance
6. **Casos de Uso**: Debug, produção, profiling

---

## 1. Compiler Flags (-gcflags)

### O Que São?

**Compiler flags** controlam como o compilador Go processa seu código. Eles afetam:
- Otimizações
- Informações de debug
- Análise de escape
- Inlining
- E muito mais

### Sintaxe Básica

```bash
go build -gcflags="flags" main.go
```

### Flags Comuns

#### -m: Mostrar Decisões de Otimização

```bash
# Ver escape analysis
go build -gcflags="-m" main.go

# Mais detalhes
go build -gcflags="-m -m" main.go

# Máximo detalhe
go build -gcflags="-m -m -m" main.go
```

**Output:**
```
./main.go:10:9: &x escapes to heap
./main.go:15:6: can inline example
```

#### -N: Desabilitar Otimizações

```bash
# Compilar sem otimizações (útil para debug)
go build -gcflags="-N" main.go
```

**Uso**: Facilita debugging, mas código fica mais lento.

#### -l: Desabilitar Inlining

```bash
# Desabilitar inlining
go build -gcflags="-l" main.go

# Desabilitar inlining agressivo
go build -gcflags="-l=4" main.go
```

**Uso**: Útil para debugging, profiling mais preciso.

#### -S: Gerar Assembly

```bash
# Ver código assembly gerado
go build -gcflags="-S" main.go
```

**Uso**: Entender como código é compilado, otimizações.

#### Combinando Flags

```bash
# Múltiplos flags
go build -gcflags="-m -N -l" main.go

# Flags diferentes para pacotes diferentes
go build -gcflags="all=-N" main.go  # -N para todos os pacotes
go build -gcflags="main=-N" main.go # -N apenas para main
```

---

## 2. Linker Flags (-ldflags)

### O Que São?

**Linker flags** controlam como o linker cria o executável final. Eles afetam:
- Tamanho do binário
- Informações de debug
- Variáveis embutidas
- Stack size
- E muito mais

### Sintaxe Básica

```bash
go build -ldflags="flags" main.go
```

### Flags Comuns

#### -s: Remover Símbolos de Debug

```bash
# Remover símbolos de debug (binário menor)
go build -ldflags="-s" main.go
```

**Benefício**: Reduz tamanho do binário significativamente.

#### -w: Remover DWARF

```bash
# Remover informações DWARF (binário menor)
go build -ldflags="-w" main.go

# Combinar -s e -w
go build -ldflags="-s -w" main.go
```

**Benefício**: Reduz ainda mais o tamanho, mas perde informações de debug.

#### -X: Definir Variáveis em Tempo de Build

```bash
# Definir variável em tempo de build
go build -ldflags="-X main.Version=1.0.0" main.go

# Múltiplas variáveis
go build -ldflags="-X main.Version=1.0.0 -X main.BuildTime=$(date)" main.go
```

**Exemplo:**
```go
package main

import "fmt"

var (
    Version   string
    BuildTime string
    GitCommit string
)

func main() {
    fmt.Printf("Version: %s\n", Version)
    fmt.Printf("BuildTime: %s\n", BuildTime)
    fmt.Printf("GitCommit: %s\n", GitCommit)
}
```

**Compilar:**
```bash
go build -ldflags="-X main.Version=1.0.0 -X main.BuildTime=$(date) -X main.GitCommit=$(git rev-parse --short HEAD)" main.go
```

#### -H: Mostrar Informações do Linker

```bash
# Ver informações do processo de linking
go build -ldflags="-H windowsgui" main.go  # Windows: sem console
```

---

## 3. Build Tags (-tags)

### Usar Tags Customizadas

```bash
# Compilar com tag específica
go build -tags "debug,tls" main.go

# Múltiplas tags
go build -tags "integration,metrics" main.go
```

### Tags em Código

```go
//go:build debug

package main

func init() {
    enableDebug()
}
```

**Compilar:**
```bash
go build -tags debug
```

---

## 4. Race Detector (-race)

### Detectar Race Conditions

```bash
# Compilar com race detector
go build -race main.go

# Testar com race detector
go test -race ./...
```

**⚠️ Importante**: 
- Aumenta uso de memória
- Reduz performance significativamente
- Use apenas para testes/debug

**Exemplo:**
```go
package main

import (
    "fmt"
    "sync"
)

var counter int
var mu sync.Mutex

func increment() {
    mu.Lock()
    counter++
    mu.Unlock()
}

func main() {
    var wg sync.WaitGroup
    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go func() {
            increment()
            wg.Done()
        }()
    }
    wg.Wait()
    fmt.Println(counter)
}
```

**Testar:**
```bash
go run -race main.go
```

---

## 5. Otimizações de Build

### Reduzir Tamanho do Binário

```bash
# Build otimizado (menor tamanho)
CGO_ENABLED=0 go build -ldflags="-s -w" -trimpath main.go
```

**Flags:**
- `CGO_ENABLED=0`: Desabilita CGO (facilita cross-compilation)
- `-s`: Remove símbolos de debug
- `-w`: Remove DWARF
- `-trimpath`: Remove caminhos do sistema de arquivos

### Build para Produção

```bash
# Build de produção completo
CGO_ENABLED=0 \
go build \
  -ldflags="-s -w -X main.Version=$(git describe --tags) -X main.BuildTime=$(date -u +%Y-%m-%dT%H:%M:%SZ) -X main.GitCommit=$(git rev-parse --short HEAD)" \
  -trimpath \
  -o myapp \
  main.go
```

### Build com Informações de Debug

```bash
# Build com debug (maior, mas melhor para debugging)
go build -gcflags="all=-N -l" -ldflags="" main.go
```

---

## 6. Casos de Uso Práticos

### Caso 1: Build com Versão

```bash
#!/bin/bash
VERSION=$(git describe --tags --always --dirty)
BUILD_TIME=$(date -u +%Y-%m-%dT%H:%M:%SZ)
GIT_COMMIT=$(git rev-parse --short HEAD)

go build \
  -ldflags="-X main.Version=$VERSION -X main.BuildTime=$BUILD_TIME -X main.GitCommit=$GIT_COMMIT" \
  -o myapp \
  main.go
```

### Caso 2: Build Otimizado

```bash
# Build menor e mais rápido
CGO_ENABLED=0 go build -ldflags="-s -w" -trimpath -o myapp main.go
```

### Caso 3: Build para Debug

```bash
# Build com informações de debug
go build -gcflags="all=-N -l" -o myapp-debug main.go
```

### Caso 4: Análise de Escape

```bash
# Ver todas as decisões de escape
go build -gcflags="-m -m" main.go 2>&1 | grep escape
```

---

## 7. Flags Avançadas

### Compiler Flags Avançados

```bash
# -B: Desabilitar bounds checking (PERIGOSO!)
go build -gcflags="-B" main.go

# -C: Número de processadores para compilação paralela
go build -gcflags="-C 4" main.go

# -D: Definir variáveis de compilação
go build -gcflags="-D path/to/pkg=value" main.go
```

### Linker Flags Avançados

```bash
# -extldflags: Flags para linker externo
go build -ldflags="-extldflags '-static'" main.go

# -linkmode: Modo de linking
go build -ldflags="-linkmode external" main.go
```

---

## 8. Scripts de Build

### Makefile

```makefile
.PHONY: build build-release build-debug

VERSION ?= dev
BUILD_TIME := $(shell date -u +%Y-%m-%dT%H:%M:%SZ)
GIT_COMMIT := $(shell git rev-parse --short HEAD)

build:
	go build -o myapp main.go

build-release:
	CGO_ENABLED=0 go build \
		-ldflags="-s -w -X main.Version=$(VERSION) -X main.BuildTime=$(BUILD_TIME) -X main.GitCommit=$(GIT_COMMIT)" \
		-trimpath \
		-o myapp \
		main.go

build-debug:
	go build \
		-gcflags="all=-N -l" \
		-o myapp-debug \
		main.go

test-race:
	go test -race ./...
```

### Script Bash

```bash
#!/bin/bash
set -e

VERSION=${1:-dev}
BUILD_TIME=$(date -u +%Y-%m-%dT%H:%M:%SZ)
GIT_COMMIT=$(git rev-parse --short HEAD)

echo "Building version $VERSION..."

CGO_ENABLED=0 go build \
  -ldflags="-s -w -X main.Version=$VERSION -X main.BuildTime=$BUILD_TIME -X main.GitCommit=$GIT_COMMIT" \
  -trimpath \
  -o myapp \
  main.go

echo "Build complete: myapp"
```

---

## 9. Boas Práticas

### ✅ Use Flags Apropriados para Cada Ambiente

**Desenvolvimento:**
```bash
go build main.go  # Sem otimizações agressivas
```

**Produção:**
```bash
CGO_ENABLED=0 go build -ldflags="-s -w" -trimpath main.go
```

### ✅ Documente Flags Customizadas

```bash
# build.sh
# Flags usados:
# -s: Remove símbolos de debug
# -w: Remove DWARF
# -X: Define variáveis em build time
```

### ✅ Use Scripts para Builds Complexos

Não memorize comandos longos. Use scripts ou Makefiles.

### ✅ Teste com Race Detector

```bash
# Sempre teste com race detector antes de release
go test -race ./...
```

---

## 10. Resumo

Nesta aula aprendemos:

1. **Compiler Flags (-gcflags)**: Controlam compilação
2. **Linker Flags (-ldflags)**: Controlam linking
3. **Build Tags (-tags)**: Tags customizadas
4. **Race Detector (-race)**: Detecção de race conditions
5. **Otimizações**: Reduzir tamanho, melhorar performance
6. **Scripts**: Automatizar builds complexos

**Lembre-se**: Flags são ferramentas poderosas. Use-as apropriadamente para cada situação!

---

**Referências:**
- [go build](https://pkg.go.dev/cmd/go#hdr-Compile_packages_and_dependencies)
- [go test](https://pkg.go.dev/cmd/go#hdr-Test_packages)



