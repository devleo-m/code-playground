# Módulo 40: Build Constraints & Tags em Go

## Aula 1: Build Constraints & Tags - Código Condicional por Plataforma

Olá! Bem-vindo ao módulo 40 sobre **Build Constraints & Tags** em Go. Este módulo ensina como escrever código que se adapta a diferentes plataformas, ambientes e configurações usando comentários especiais que controlam quais arquivos são incluídos na compilação.

Build constraints permitem criar aplicações portáveis que funcionam em diferentes sistemas operacionais, arquiteturas, e configurações sem precisar manter múltiplos projetos separados.

Nesta aula, vamos explorar:
1. **O que são Build Constraints**: Conceito e propósito
2. **Sintaxe**: Como escrever constraints
3. **Build Tags**: Tags customizadas
4. **Casos de Uso**: Platform-specific code, feature toggles
5. **Boas Práticas**: Organização e manutenção

---

## 1. O Que São Build Constraints?

### Definição

**Build Constraints** (também chamados de "build tags") são comentários especiais que controlam quais arquivos são incluídos quando você compila um programa Go. Eles permitem:

- ✅ **Código específico de plataforma**: Diferentes implementações para Linux, Windows, macOS
- ✅ **Código específico de arquitetura**: Diferentes implementações para x86, ARM, etc.
- ✅ **Feature toggles**: Habilitar/desabilitar funcionalidades
- ✅ **Builds condicionais**: Debug vs Release, testes, etc.

### Por Que Usar?

**Vantagens:**
- ✅ **Um único projeto**: Não precisa manter projetos separados
- ✅ **Código limpo**: Cada plataforma tem seu próprio arquivo
- ✅ **Compilação eficiente**: Apenas código necessário é compilado
- ✅ **Manutenção fácil**: Mudanças em uma plataforma não afetam outras

---

## 2. Sintaxe: //go:build

### Formato Moderno (Go 1.17+)

```go
//go:build linux
// +build linux

package main

import "fmt"

func main() {
    fmt.Println("Running on Linux!")
}
```

**Nota**: Go 1.17+ usa `//go:build`, mas ainda aceita `// +build` para compatibilidade.

### Formato Antigo (Compatibilidade)

```go
// +build linux

package main
```

**Recomendação**: Use `//go:build` em código novo.

### Operadores Lógicos

```go
// AND: linux AND amd64
//go:build linux && amd64

// OR: linux OR darwin
//go:build linux || darwin

// NOT: não windows
//go:build !windows

// Combinações complexas
//go:build (linux || darwin) && amd64
```

---

## 3. Tags de Plataforma Comuns

### Sistemas Operacionais

```go
//go:build linux
//go:build windows
//go:build darwin      // macOS
//go:build freebsd
//go:build openbsd
//go:build netbsd
//go:build dragonfly
//go:build android
//go:build ios
//go:build js          // WebAssembly
```

### Arquiteturas

```go
//go:build amd64
//go:build 386         // x86 32-bit
//go:build arm
//go:build arm64
//go:build mips
//go:build mips64
//go:build ppc64
//go:build ppc64le
//go:build riscv64
//go:build s390x
//go:build wasm        // WebAssembly
```

### Combinações Comuns

```go
//go:build linux && amd64
//go:build windows && 386
//go:build darwin && arm64  // macOS Apple Silicon
```

---

## 4. Exemplos Práticos

### Exemplo 1: Código Específico de Plataforma

**Arquivo: `main_linux.go`**
```go
//go:build linux

package main

import "fmt"

func getOS() string {
    return "Linux"
}

func main() {
    fmt.Println("OS:", getOS())
}
```

**Arquivo: `main_windows.go`**
```go
//go:build windows

package main

import "fmt"

func getOS() string {
    return "Windows"
}

func main() {
    fmt.Println("OS:", getOS())
}
```

**Arquivo: `main_darwin.go`**
```go
//go:build darwin

package main

import "fmt"

func getOS() string {
    return "macOS"
}

func main() {
    fmt.Println("OS:", getOS())
}
```

**Arquivo: `main_default.go`** (sem constraint)
```go
package main

import "fmt"

func getOS() string {
    return "Unknown"
}

func main() {
    fmt.Println("OS:", getOS())
}
```

### Exemplo 2: Feature Toggles

**Arquivo: `feature_debug.go`**
```go
//go:build debug

package main

func init() {
    enableDebugMode()
}

func enableDebugMode() {
    // Código de debug
}
```

**Arquivo: `feature_release.go`** (sem constraint ou `!debug`)
```go
//go:build !debug

package main

func init() {
    // Sem código de debug
}
```

**Compilar:**
```bash
go build -tags debug    # Inclui código de debug
go build                # Sem código de debug
```

### Exemplo 3: Múltiplas Condições

**Arquivo: `performance_linux_amd64.go`**
```go
//go:build linux && amd64

package main

func optimizedFunction() {
    // Implementação otimizada para Linux x86_64
}
```

**Arquivo: `performance_default.go`**
```go
//go:build !(linux && amd64)

package main

func optimizedFunction() {
    // Implementação genérica
}
```

---

## 5. Build Tags Customizadas

### Definindo Tags Customizadas

Você pode definir suas próprias tags usando `-tags`:

```bash
go build -tags "feature1,feature2"
```

### Exemplo: Feature Flags

**Arquivo: `http_server.go`** (padrão)
```go
package main

func startServer() {
    // Servidor básico
}
```

**Arquivo: `http_server_tls.go`**
```go
//go:build tls

package main

func startServer() {
    // Servidor com TLS
}
```

**Arquivo: `http_server_metrics.go`**
```go
//go:build metrics

package main

func startServer() {
    // Servidor com métricas
}
```

**Compilar:**
```bash
go build                          # Servidor básico
go build -tags tls                # Servidor com TLS
go build -tags "tls,metrics"      # Servidor com TLS e métricas
```

---

## 6. Casos de Uso Reais

### Caso 1: Código Específico de OS

```go
// file_unix.go
//go:build !windows

package main

import "syscall"

func setNonBlocking(fd int) error {
    return syscall.SetNonBlocking(fd, true)
}
```

```go
// file_windows.go
//go:build windows

package main

import "golang.org/x/sys/windows"

func setNonBlocking(fd int) error {
    // Implementação Windows
}
```

### Caso 2: Debug vs Release

```go
// debug.go
//go:build debug

package main

var debugMode = true

func debugLog(msg string) {
    fmt.Println("[DEBUG]", msg)
}
```

```go
// release.go
//go:build !debug

package main

var debugMode = false

func debugLog(msg string) {
    // No-op em release
}
```

### Caso 3: Testes de Integração

```go
// integration_test.go
//go:build integration

package main

import "testing"

func TestIntegration(t *testing.T) {
    // Testes de integração
}
```

**Executar:**
```bash
go test -tags integration
```

---

## 7. Boas Práticas

### ✅ Organize por Plataforma

```
project/
├── main.go              # Código comum
├── main_linux.go        # Linux específico
├── main_windows.go      # Windows específico
├── main_darwin.go       # macOS específico
└── utils/
    ├── file_unix.go     # Unix comum
    └── file_windows.go   # Windows específico
```

### ✅ Use Nomes Descritivos

```go
// ❌ Ruim
//go:build x

// ✅ Bom
//go:build linux && amd64
```

### ✅ Documente Tags Customizadas

```go
//go:build tls
// Este arquivo é incluído apenas quando compilado com -tags tls
// Ele adiciona suporte a TLS ao servidor HTTP
```

### ✅ Teste em Todas as Plataformas

```bash
# Testar em diferentes plataformas
GOOS=linux go build
GOOS=windows go build
GOOS=darwin go build
```

---

## 8. Verificando Build Tags

### Listar Arquivos que Serão Compilados

```bash
go list -tags "debug,tls" ./...
```

### Verificar Tags de um Arquivo

```bash
grep "//go:build" *.go
```

### Testar Build

```bash
# Testar build com tags
go build -tags "debug,tls"

# Verificar se compila sem tags
go build
```

---

## 9. Armadilhas Comuns

### ❌ Armadilha 1: Esquecer Arquivo Padrão

```go
// ❌ Ruim: Apenas linux.go e windows.go
// Se compilar em macOS, não tem implementação!
```

**Solução**: Sempre tenha um arquivo padrão sem constraint.

### ❌ Armadilha 2: Constraints Conflitantes

```go
// ❌ Ruim: Dois arquivos com mesma constraint
// file_linux.go: //go:build linux
// file_unix.go:  //go:build linux
```

**Solução**: Use constraints mais específicas.

### ❌ Armadilha 3: Esquecer de Testar

```go
// ❌ Ruim: Compila, mas não testa em outras plataformas
```

**Solução**: Teste em todas as plataformas suportadas.

---

## 10. Exemplo Completo

```go
// main.go (código comum)
package main

import "fmt"

func main() {
    fmt.Println("OS:", getOS())
    fmt.Println("Arch:", getArch())
}

// os_linux.go
//go:build linux

package main

func getOS() string {
    return "Linux"
}

// os_windows.go
//go:build windows

package main

func getOS() string {
    return "Windows"
}

// os_darwin.go
//go:build darwin

package main

func getOS() string {
    return "macOS"
}

// arch_amd64.go
//go:build amd64

package main

func getArch() string {
    return "amd64"
}

// arch_arm64.go
//go:build arm64

package main

func getArch() string {
    return "arm64"
}
```

---

## Resumo

Nesta aula aprendemos:

1. **Build Constraints**: Controlam quais arquivos são compilados
2. **Sintaxe**: `//go:build` com operadores lógicos
3. **Tags de Plataforma**: OS, arquitetura, combinações
4. **Tags Customizadas**: Feature flags, debug, etc.
5. **Casos de Uso**: Platform-specific code, feature toggles
6. **Boas Práticas**: Organização, documentação, testes

**Lembre-se**: Build constraints são essenciais para criar aplicações portáveis. Use-os para manter código limpo e organizado!

---

**Referências:**
- [Build Constraints](https://pkg.go.dev/cmd/go#hdr-Build_constraints)
- [Build Tags](https://golang.org/pkg/go/build/#hdr-Build_Constraints)

