# Módulo 35: Deployment & Tooling em Go
## Aula 4: Performance e Boas Práticas - Otimizando Builds e Deployments

Olá! Agora vamos mergulhar em **boas práticas** e **otimizações** para tornar seus builds mais eficientes, seus binários mais rápidos e seus deployments mais confiáveis.

---

## 🚀 Otimizações de Build

### 1. Build Cache e Recompilação Inteligente

Go mantém um cache de builds para evitar recompilar código que não mudou.

**Boas Práticas:**

```bash
# Limpar cache (se necessário)
go clean -cache

# Verificar cache
go env GOCACHE

# Forçar recompilação completa (raramente necessário)
go build -a
```

**Quando usar `-a`:**
- Quando você suspeita que o cache está corrompido
- Quando mudou versão do Go e quer garantir compatibilidade
- Em CI/CD quando quer builds completamente limpos

**Quando NÃO usar `-a`:**
- Em desenvolvimento local (torna builds mais lentos)
- Quando você quer aproveitar o cache

### 2. Build Paralelo

Go compila pacotes em paralelo automaticamente. Você pode controlar:

```bash
# Usar todos os CPUs disponíveis (padrão)
go build

# Limitar número de processos
go build -p 4  # Usa 4 processos

# Ver quantos processos estão sendo usados
go build -x  # Mostra comandos, incluindo paralelismo
```

**Boas Práticas:**
- Deixe Go usar todos os CPUs (padrão é otimizado)
- Em CI/CD, considere limitar se houver problemas de memória
- Use `-p` apenas se tiver problemas específicos

### 3. Modo de Desenvolvimento vs Produção

**Desenvolvimento:**
```bash
# Build rápido, com informações de debug
go build

# Ou para testes rápidos
go run main.go
```

**Produção:**
```bash
# Build otimizado, sem debug, menor tamanho
go build -ldflags "-s -w" -trimpath -o myapp
```

**Criar variantes:**

```makefile
# Makefile
.PHONY: build build-dev build-prod

build-dev:
	go build -o bin/myapp-dev

build-prod:
	go build -ldflags "-s -w" -trimpath -o bin/myapp-prod
```

### 4. Reduzindo Tamanho do Binário

#### Técnicas Básicas

```bash
# 1. Remover símbolos de debug
go build -ldflags "-s -w" -o myapp

# 2. Desabilitar CGO (se não precisar)
CGO_ENABLED=0 go build -o myapp

# 3. Combinar ambas
CGO_ENABLED=0 go build -ldflags "-s -w" -o myapp
```

#### Comparação de Tamanhos

```bash
# Script para comparar tamanhos
#!/bin/bash

echo "Build normal:"
go build -o myapp-normal
ls -lh myapp-normal

echo -e "\nBuild sem debug:"
go build -ldflags "-s -w" -o myapp-no-debug
ls -lh myapp-no-debug

echo -e "\nBuild sem CGO:"
CGO_ENABLED=0 go build -o myapp-no-cgo
ls -lh myapp-no-cgo

echo -e "\nBuild otimizado completo:"
CGO_ENABLED=0 go build -ldflags "-s -w" -o myapp-optimized
ls -lh myapp-optimized
```

**Resultados típicos:**
- Normal: 10-15 MB
- Sem debug: 6-10 MB (30-40% menor)
- Sem CGO: 8-12 MB (depende do uso de CGO)
- Otimizado completo: 5-8 MB (40-50% menor)

#### Técnicas Avançadas

**UPX (compressão):**
```bash
# Instalar UPX
# Linux: sudo apt install upx
# Mac: brew install upx

# Comprimir binário
upx myapp

# Comprimir agressivamente
upx --best --lzma myapp
```

**⚠️ Atenção**: UPX pode causar problemas com antivírus e pode aumentar tempo de inicialização.

**TinyGo (para binários muito pequenos):**
```bash
# TinyGo é um compilador alternativo que gera binários menores
# Útil para embarcados e WebAssembly
tinygo build -o myapp-tiny
```

### 5. Build Incremental

Go só recompila o que mudou. Para maximizar isso:

```bash
# Estrutura de projeto modular ajuda
projeto/
├── cmd/
│   └── server/
│       └── main.go
├── internal/
│   ├── pkg1/
│   └── pkg2/
└── go.mod

# Mudanças em pkg1 não recompilam pkg2
```

**Boas Práticas:**
- Organize código em pacotes independentes
- Evite dependências circulares
- Use `internal/` para código privado

---

## 🎯 Boas Práticas de Build

### 1. Versionamento de Binários

**Sempre inclua informações de versão:**

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
        fmt.Printf("Version: %s\n", Version)
        fmt.Printf("Build Time: %s\n", BuildTime)
        fmt.Printf("Git Commit: %s\n", GitCommit)
        os.Exit(0)
    }
    
    // Seu código aqui
}
```

**Build com versão:**
```bash
VERSION=$(git describe --tags --always --dirty)
BUILD_TIME=$(date -u +%Y-%m-%dT%H:%M:%SZ)
GIT_COMMIT=$(git rev-parse --short HEAD)

go build -ldflags \
  "-X main.Version=$VERSION \
   -X main.BuildTime=$BUILD_TIME \
   -X main.GitCommit=$GIT_COMMIT" \
  -o myapp
```

**Por que é importante:**
- Debugging em produção
- Rastreabilidade
- Rollback quando necessário
- Compliance e auditoria

### 2. Build Reproducible

**Builds reproduzíveis** geram binários idênticos quando compilados com as mesmas fontes.

```bash
# Habilitar build reproduzível
go build -trimpath -ldflags "-s -w -buildid=" -o myapp
```

**Benefícios:**
- Segurança (verificação de integridade)
- Debugging (mesmo binário = mesmo comportamento)
- CI/CD (detectar mudanças reais)

### 3. Estrutura de Diretórios

**Estrutura recomendada:**

```
projeto/
├── cmd/
│   ├── server/
│   │   └── main.go
│   └── cli/
│       └── main.go
├── internal/
│   ├── config/
│   ├── database/
│   └── handlers/
├── pkg/
│   └── public/
├── api/
│   └── openapi.yaml
├── scripts/
│   └── build.sh
├── Makefile
├── go.mod
└── README.md
```

**Makefile para múltiplos comandos:**

```makefile
.PHONY: build-server build-cli build-all

build-server:
	go build -o bin/server ./cmd/server

build-cli:
	go build -o bin/cli ./cmd/cli

build-all: build-server build-cli
```

### 4. Testes Antes do Build

**Sempre teste antes de buildar:**

```makefile
.PHONY: test build

test:
	go test ./...

build: test
	go build -o bin/myapp
```

**Ou em script:**

```bash
#!/bin/bash
set -e  # Parar em caso de erro

echo "Running tests..."
go test ./...

echo "Building..."
go build -o bin/myapp
```

### 5. Validação de Build

**Validar que o build funcionou:**

```bash
#!/bin/bash
go build -o myapp

# Verificar que o executável existe e é executável
if [ ! -f myapp ]; then
    echo "ERRO: Build falhou!"
    exit 1
fi

# Testar execução básica
./myapp --help || ./myapp version || true

echo "Build validado com sucesso!"
```

---

## 🌍 Boas Práticas de Cross-compilation

### 1. Sempre Teste Binários Cross-compilados

**Problema comum**: Binário compila mas não funciona na plataforma de destino.

**Solução**: Teste em ambiente similar:

```bash
# Compilar para Linux
GOOS=linux GOARCH=amd64 go build -o myapp-linux

# Testar em container Docker
docker run --rm -v $(pwd):/app -w /app alpine:latest ./myapp-linux
```

**Ou usar QEMU para emulação:**

```bash
# Instalar QEMU
# Linux: sudo apt install qemu-user-static
# Mac: brew install qemu

# Testar binário ARM em x86
qemu-arm-static ./myapp-arm
```

### 2. CGO e Cross-compilation

**Regra geral**: Desabilite CGO para cross-compilation mais fácil.

```bash
# Sempre use CGO_ENABLED=0 para cross-compilation
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o myapp
```

**Quando você PRECISA de CGO:**
- Você precisa de bibliotecas C específicas
- Você precisa usar código C existente

**Solução**: Use toolchains específicos ou compile na plataforma de destino.

### 3. Build Tags para Código Específico de Plataforma

**Use build tags para código específico:**

```go
// +build linux

package main

func init() {
    // Código específico para Linux
}
```

```go
// +build windows

package main

func init() {
    // Código específico para Windows
}
```

**Build condicional:**

```bash
# Build para Linux (inclui código Linux)
GOOS=linux go build

# Build para Windows (inclui código Windows)
GOOS=windows go build
```

### 4. Verificar Plataformas Disponíveis

**Antes de compilar, verifique suporte:**

```bash
# Listar todas as plataformas
go tool dist list

# Verificar se uma plataforma específica é suportada
go tool dist list | grep "linux/arm64"
```

### 5. Scripts de Build Multiplataforma

**Criar scripts robustos:**

```bash
#!/bin/bash
set -e  # Parar em caso de erro

APP_NAME="myapp"
VERSION=$(git describe --tags --always --dirty)
BUILD_TIME=$(date -u +%Y-%m-%dT%H:%M:%SZ)
GIT_COMMIT=$(git rev-parse --short HEAD)

LDFLAGS="-X main.Version=${VERSION} -X main.BuildTime=${BUILD_TIME} -X main.GitCommit=${GIT_COMMIT} -s -w"

mkdir -p dist

# Função para build com validação
build_platform() {
    local os=$1
    local arch=$2
    local ext=$3
    
    echo "Building for ${os}/${arch}..."
    
    CGO_ENABLED=0 GOOS=${os} GOARCH=${arch} \
        go build -ldflags "${LDFLAGS}" \
        -o "dist/${APP_NAME}-${os}-${arch}${ext}" \
        || {
            echo "ERRO: Build falhou para ${os}/${arch}"
            exit 1
        }
    
    echo "  ✓ ${APP_NAME}-${os}-${arch}${ext}"
}

# Build para todas as plataformas
build_platform linux amd64 ""
build_platform linux arm64 ""
build_platform darwin amd64 ""
build_platform darwin arm64 ""
build_platform windows amd64 ".exe"

echo -e "\nTodos os builds completos!"
echo "Arquivos em ./dist/"
ls -lh dist/
```

---

## 🔒 Segurança em Builds

### 1. Não Incluir Secrets no Binário

**❌ ERRADO:**
```go
const API_KEY = "sk-1234567890"  // NUNCA faça isso!
```

**✅ CORRETO:**
```go
apiKey := os.Getenv("API_KEY")
if apiKey == "" {
    log.Fatal("API_KEY não configurada")
}
```

### 2. Build Reproducible para Verificação

**Permite verificar integridade:**

```bash
# Build reproduzível
go build -trimpath -ldflags "-s -w -buildid=" -o myapp

# Gerar checksum
sha256sum myapp > myapp.sha256

# Verificar depois
sha256sum -c myapp.sha256
```

### 3. Assinatura de Binários

**Para distribuição pública:**

**macOS:**
```bash
# Assinar binário
codesign --sign "Developer ID Application: Your Name" myapp

# Verificar
codesign --verify myapp
```

**Windows:**
```powershell
# Assinar com certificado
signtool sign /f certificate.pfx /p password myapp.exe

# Verificar
signtool verify /pa myapp.exe
```

### 4. Verificação de Dependências

**Verificar vulnerabilidades:**

```bash
# Verificar dependências
go list -json -deps | nancy sleuth

# Ou usar govulncheck
go install golang.org/x/vuln/cmd/govulncheck@latest
govulncheck ./...
```

---

## 📊 Performance de Builds

### 1. Medir Tempo de Build

```bash
# Medir tempo
time go build

# Ou mais detalhado
go build -x 2>&1 | grep "real"
```

### 2. Build Paralelo

Go já usa paralelismo, mas você pode otimizar:

```bash
# Usar todos os CPUs (padrão)
go build -p $(nproc)  # Linux
go build -p $(sysctl -n hw.ncpu)  # Mac
```

### 3. Cache de Módulos

**Go mantém cache de módulos baixados:**

```bash
# Ver cache
go env GOMODCACHE

# Limpar cache (se necessário)
go clean -modcache
```

**Boas Práticas:**
- Em CI/CD, cache o `GOMODCACHE`
- Não limpe o cache desnecessariamente
- Use `go mod download` para pré-baixar dependências

### 4. Build Incremental

**Maximize builds incrementais:**

```makefile
# Makefile com dependências
bin/myapp: $(shell find . -name '*.go' -not -path './vendor/*')
	go build -o $@ .

# Só recompila se arquivos mudaram
```

---

## 🚢 Deploy e Distribuição

### 1. Estrutura de Release

```
releases/
├── v1.0.0/
│   ├── myapp-linux-amd64
│   ├── myapp-darwin-amd64
│   ├── myapp-darwin-arm64
│   ├── myapp-windows-amd64.exe
│   ├── checksums.txt
│   └── README.md
```

### 2. Checksums

**Sempre forneça checksums:**

```bash
#!/bin/bash
# Gerar checksums
cd dist/
sha256sum * > checksums.txt

# Ou para Mac
shasum -a 256 * > checksums.txt
```

### 3. GitHub Releases

**Automatizar releases:**

```yaml
# .github/workflows/release.yml
name: Release

on:
  push:
    tags:
      - 'v*'

jobs:
  release:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      - uses: actions/setup-go@v4
        with:
          go-version: '1.21'
      
      - name: Build
        run: |
          ./scripts/build-all.sh
          cd dist && sha256sum * > checksums.txt
      
      - name: Create Release
        uses: softprops/action-gh-release@v1
        with:
          files: dist/*
```

### 4. Containerização

**Dockerfile otimizado:**

```dockerfile
# Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-s -w" -o app .

# Runtime stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/

COPY --from=builder /build/app .

CMD ["./app"]
```

**Build:**
```bash
docker build -t myapp:latest .
```

---

## 📋 Checklist de Boas Práticas

Antes de fazer deploy, verifique:

- [ ] Binário inclui informações de versão
- [ ] Build é reproduzível (quando possível)
- [ ] Testes passaram antes do build
- [ ] Binário foi testado na plataforma de destino
- [ ] Checksums foram gerados
- [ ] Dependências foram verificadas (sem vulnerabilidades)
- [ ] Secrets não estão no binário
- [ ] Build está otimizado para produção (`-s -w`)
- [ ] CGO está desabilitado (se não necessário)
- [ ] Documentação de deploy está atualizada

---

## 🎯 Resumo

**Otimizações de Build:**
- Use cache do Go (padrão)
- Build paralelo (padrão)
- Diferencie builds de dev e produção
- Reduza tamanho com `-s -w` e `CGO_ENABLED=0`

**Boas Práticas:**
- Sempre inclua versão no binário
- Teste antes de buildar
- Valide builds
- Use estrutura de diretórios clara

**Cross-compilation:**
- Sempre teste binários cross-compilados
- Use `CGO_ENABLED=0` quando possível
- Use build tags para código específico de plataforma
- Crie scripts robustos

**Segurança:**
- Nunca inclua secrets no binário
- Use builds reproduzíveis
- Assine binários para distribuição
- Verifique vulnerabilidades em dependências

**Performance:**
- Aproveite cache e builds incrementais
- Meça tempo de build
- Otimize estrutura de projeto

Seguindo essas práticas, você terá builds mais rápidos, binários mais seguros e deployments mais confiáveis!

