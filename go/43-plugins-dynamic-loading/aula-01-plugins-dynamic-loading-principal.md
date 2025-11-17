# Módulo 43: Plugins & Dynamic Loading em Go

## Aula 1: Plugins & Dynamic Loading - Carregamento Dinâmico de Código

Olá! Bem-vindo ao módulo 43, a **última aula** sobre tópicos avançados em Go! Este módulo ensina como usar o sistema de plugins do Go para carregar código dinamicamente em tempo de execução.

**⚠️ IMPORTANTE**: O sistema de plugins do Go tem limitações significativas e não é amplamente usado. Entenda quando faz sentido e quando não faz.

Nesta aula, vamos explorar:
1. **O que são Plugins**: Conceito e propósito
2. **Package plugin**: A biblioteca padrão
3. **Como Funciona**: Build e carregamento
4. **Limitações**: O que não funciona
5. **Casos de Uso**: Quando usar plugins
6. **Alternativas**: Quando não usar plugins

---

## 1. O Que São Plugins?

### Definição

**Plugins** são bibliotecas compartilhadas (`.so` no Linux, `.dylib` no macOS) que podem ser carregadas dinamicamente em tempo de execução. Isso permite:
- ✅ **Extensibilidade**: Adicionar funcionalidades sem recompilar
- ✅ **Modularidade**: Separar funcionalidades em módulos
- ✅ **Hot-swapping**: Trocar implementações sem reiniciar

### Por Que Existe?

**Casos de uso:**
- ✅ **Sistemas extensíveis**: Aplicações que precisam de plugins de terceiros
- ✅ **Arquiteturas modulares**: Sistemas com múltiplos módulos
- ✅ **Hot-reloading**: Atualizar código sem reiniciar aplicação

### Limitações Importantes

**⚠️ Limitações do Go plugins:**
- ⚠️ **Unix-only**: Não funciona no Windows
- ⚠️ **Versão do Go**: Plugin e aplicação devem usar mesma versão do Go
- ⚠️ **Build mode**: Precisa compilar com `-buildmode=plugin`
- ⚠️ **Complexidade**: Pode ser complicado de manter
- ⚠️ **Pouco usado**: Não é amplamente adotado na comunidade

---

## 2. Package plugin

### Importar

```go
import "plugin"
```

### Funcionalidades Principais

O package `plugin` fornece:
- `plugin.Open()`: Abrir arquivo de plugin
- `plugin.Lookup()`: Buscar símbolo (função, variável) no plugin
- `plugin.Plugin`: Interface para plugin carregado

---

## 3. Criando um Plugin

### Código do Plugin

**Arquivo: `greeter.go`**
```go
package main

import "fmt"

// Função exportada (deve começar com maiúscula)
func Greet(name string) {
    fmt.Printf("Hello, %s!\n", name)
}

// Variável exportada
var Version = "1.0.0"

// Função init (executada ao carregar plugin)
func init() {
    fmt.Println("Plugin greeter loaded!")
}
```

**Compilar como plugin:**
```bash
go build -buildmode=plugin -o greeter.so greeter.go
```

**Resultado**: Arquivo `greeter.so` (shared object)

---

## 4. Carregando um Plugin

### Carregamento Básico

**Arquivo: `main.go`**
```go
package main

import (
    "fmt"
    "plugin"
)

func main() {
    // Carregar plugin
    p, err := plugin.Open("greeter.so")
    if err != nil {
        panic(err)
    }
    
    // Buscar função
    greetFunc, err := p.Lookup("Greet")
    if err != nil {
        panic(err)
    }
    
    // Converter para tipo correto e chamar
    greet := greetFunc.(func(string))
    greet("World")
    
    // Buscar variável
    version, err := p.Lookup("Version")
    if err != nil {
        panic(err)
    }
    
    versionStr := *version.(*string)
    fmt.Printf("Plugin version: %s\n", versionStr)
}
```

**Executar:**
```bash
go run main.go
```

**Output:**
```
Plugin greeter loaded!
Hello, World!
Plugin version: 1.0.0
```

---

## 5. Exemplo Completo: Sistema de Plugins

### Plugin 1: Calculadora

**Arquivo: `calculator.go`**
```go
package main

func Add(a, b int) int {
    return a + b
}

func Subtract(a, b int) int {
    return b - a
}

func Multiply(a, b int) int {
    return a * b
}

func Divide(a, b int) (int, error) {
    if b == 0 {
        return 0, fmt.Errorf("division by zero")
    }
    return a / b, nil
}
```

**Compilar:**
```bash
go build -buildmode=plugin -o calculator.so calculator.go
```

### Plugin 2: Formatador

**Arquivo: `formatter.go`**
```go
package main

import "fmt"

func FormatNumber(n int) string {
    return fmt.Sprintf("Number: %d", n)
}

func FormatCurrency(amount float64) string {
    return fmt.Sprintf("$%.2f", amount)
}
```

**Compilar:**
```bash
go build -buildmode=plugin -o formatter.so formatter.go
```

### Aplicação Principal

**Arquivo: `main.go`**
```go
package main

import (
    "fmt"
    "plugin"
)

func main() {
    // Carregar plugin de calculadora
    calcPlugin, err := plugin.Open("calculator.so")
    if err != nil {
        panic(err)
    }
    
    // Buscar função Add
    addFunc, err := calcPlugin.Lookup("Add")
    if err != nil {
        panic(err)
    }
    add := addFunc.(func(int, int) int)
    
    result := add(10, 20)
    fmt.Printf("10 + 20 = %d\n", result)
    
    // Carregar plugin de formatador
    fmtPlugin, err := plugin.Open("formatter.so")
    if err != nil {
        panic(err)
    }
    
    // Buscar função FormatNumber
    formatFunc, err := fmtPlugin.Lookup("FormatNumber")
    if err != nil {
        panic(err)
    }
    format := formatFunc.(func(int) string)
    
    fmt.Println(format(result))
}
```

---

## 6. Interfaces e Plugins

### Definir Interface

**Arquivo: `interfaces.go`** (compartilhado)
```go
package main

// Interface que plugins devem implementar
type Processor interface {
    Process(data string) string
}
```

### Plugin que Implementa Interface

**Arquivo: `uppercase.go`**
```go
package main

import "strings"

type UppercaseProcessor struct{}

func (p *UppercaseProcessor) Process(data string) string {
    return strings.ToUpper(data)
}

// Função factory exportada
func NewProcessor() Processor {
    return &UppercaseProcessor{}
}
```

**Compilar:**
```bash
go build -buildmode=plugin -o uppercase.so uppercase.go
```

### Usar Plugin com Interface

**Arquivo: `main.go`**
```go
package main

import (
    "fmt"
    "plugin"
)

func main() {
    p, err := plugin.Open("uppercase.so")
    if err != nil {
        panic(err)
    }
    
    newFunc, err := p.Lookup("NewProcessor")
    if err != nil {
        panic(err)
    }
    
    newProcessor := newFunc.(func() Processor)
    processor := newProcessor()
    
    result := processor.Process("hello world")
    fmt.Println(result)  // HELLO WORLD
}
```

---

## 7. Limitações Importantes

### Limitação 1: Unix-only

```bash
# ❌ Não funciona no Windows
# ✅ Funciona apenas em Linux, macOS, etc.
```

**Solução**: Use apenas em sistemas Unix-like.

### Limitação 2: Versão do Go

```bash
# ❌ Plugin compilado com Go 1.18 não funciona com app Go 1.19
# ✅ Plugin e app devem usar MESMA versão do Go
```

**Solução**: Garanta mesma versão do Go.

### Limitação 3: Build Mode

```bash
# ❌ Não funciona sem -buildmode=plugin
go build plugin.go  # Não cria plugin!

# ✅ Precisa de build mode
go build -buildmode=plugin -o plugin.so plugin.go
```

### Limitação 4: Símbolos Exportados

```go
// ❌ Não funciona: função não exportada
func greet(name string) {  // minúscula = não exportada
    // ...
}

// ✅ Funciona: função exportada
func Greet(name string) {  // maiúscula = exportada
    // ...
}
```

### Limitação 5: Dependências

Plugins podem ter problemas com:
- Dependências conflitantes
- Versões diferentes de packages
- CGO

---

## 8. Casos de Uso

### ✅ Use Plugins Quando:

1. **Sistema extensível**: Aplicação que precisa de plugins de terceiros
2. **Hot-reloading**: Atualizar funcionalidades sem reiniciar
3. **Arquitetura modular**: Sistema com múltiplos módulos independentes
4. **Unix-only**: Aplicação que roda apenas em Unix

### ❌ NÃO Use Plugins Se:

1. **Windows necessário**: Plugins não funcionam no Windows
2. **Simplicidade importante**: Plugins adicionam complexidade
3. **Cross-platform**: Precisa funcionar em múltiplas plataformas
4. **Alternativas existem**: Interfaces, RPC, etc. podem ser melhores

---

## 9. Alternativas aos Plugins

### Alternativa 1: Interfaces e Injeção de Dependência

```go
// Em vez de plugins, use interfaces
type Processor interface {
    Process(data string) string
}

// Diferentes implementações
type UppercaseProcessor struct{}
type LowercaseProcessor struct{}

// Injetar implementação
func NewApp(processor Processor) *App {
    return &App{processor: processor}
}
```

**Vantagem**: Mais simples, funciona em todas as plataformas.

### Alternativa 2: RPC/HTTP

```go
// Em vez de plugins, use serviços separados
// Plugin como serviço HTTP
func main() {
    http.HandleFunc("/process", handleProcess)
    http.ListenAndServe(":8080", nil)
}
```

**Vantagem**: Mais flexível, isolamento melhor.

### Alternativa 3: Scripting

```go
// Use linguagens de script (Lua, Python via CGO)
// Mais flexível que plugins Go
```

**Vantagem**: Mais dinâmico, hot-reloading mais fácil.

---

## 10. Exemplo Prático: Sistema de Processadores

### Interface Compartilhada

**Arquivo: `processor.go`**
```go
package main

type Processor interface {
    Process(input string) (string, error)
    Name() string
}
```

### Plugin 1: Uppercase

**Arquivo: `uppercase_plugin.go`**
```go
package main

import "strings"

type UppercaseProcessor struct{}

func (p *UppercaseProcessor) Process(input string) (string, error) {
    return strings.ToUpper(input), nil
}

func (p *UppercaseProcessor) Name() string {
    return "uppercase"
}

func NewProcessor() Processor {
    return &UppercaseProcessor{}
}
```

**Compilar:**
```bash
go build -buildmode=plugin -o processors/uppercase.so uppercase_plugin.go
```

### Plugin 2: Reverse

**Arquivo: `reverse_plugin.go`**
```go
package main

func Reverse(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}

type ReverseProcessor struct{}

func (p *ReverseProcessor) Process(input string) (string, error) {
    return Reverse(input), nil
}

func (p *ReverseProcessor) Name() string {
    return "reverse"
}

func NewProcessor() Processor {
    return &ReverseProcessor{}
}
```

**Compilar:**
```bash
go build -buildmode=plugin -o processors/reverse.so reverse_plugin.go
```

### Aplicação Principal

**Arquivo: `main.go`**
```go
package main

import (
    "fmt"
    "os"
    "plugin"
)

func main() {
    if len(os.Args) < 3 {
        fmt.Println("Usage: go run main.go <plugin.so> <input>")
        os.Exit(1)
    }
    
    pluginPath := os.Args[1]
    input := os.Args[2]
    
    // Carregar plugin
    p, err := plugin.Open(pluginPath)
    if err != nil {
        panic(err)
    }
    
    // Buscar função factory
    newFunc, err := p.Lookup("NewProcessor")
    if err != nil {
        panic(err)
    }
    
    // Criar processador
    processor := newFunc.(func() Processor)()
    
    // Processar
    result, err := processor.Process(input)
    if err != nil {
        panic(err)
    }
    
    fmt.Printf("Processor: %s\n", processor.Name())
    fmt.Printf("Input: %s\n", input)
    fmt.Printf("Output: %s\n", result)
}
```

**Usar:**
```bash
# Compilar plugins
go build -buildmode=plugin -o processors/uppercase.so uppercase_plugin.go
go build -buildmode=plugin -o processors/reverse.so reverse_plugin.go

# Usar
go run main.go processors/uppercase.so "hello world"
go run main.go processors/reverse.so "hello world"
```

---

## 11. Boas Práticas

### ✅ Use Interfaces

Sempre defina interfaces claras que plugins devem implementar.

### ✅ Documente Requisitos

Documente claramente:
- Versão do Go necessária
- Dependências requeridas
- Interface que deve ser implementada

### ✅ Trate Erros

Sempre trate erros ao carregar e usar plugins:
```go
p, err := plugin.Open("plugin.so")
if err != nil {
    // Tratar erro apropriadamente
    return err
}
```

### ✅ Valide Símbolos

Sempre verifique se símbolos existem:
```go
symbol, err := p.Lookup("FunctionName")
if err != nil {
    // Símbolo não encontrado
    return err
}
```

### ✅ Considere Alternativas

Sempre considere alternativas antes de usar plugins:
- Interfaces e DI
- RPC/HTTP
- Scripting

---

## 12. Resumo

Nesta aula aprendemos:

1. **Plugins**: Carregamento dinâmico de código
2. **Package plugin**: Biblioteca padrão para plugins
3. **Build mode**: `-buildmode=plugin` necessário
4. **Limitações**: Unix-only, versão do Go, complexidade
5. **Casos de uso**: Sistemas extensíveis, hot-reloading
6. **Alternativas**: Interfaces, RPC, scripting

**Lembre-se**: Plugins do Go têm limitações significativas. Use apenas quando realmente necessário e sempre considere alternativas primeiro!

---

**Referências:**
- [plugin package](https://pkg.go.dev/plugin)
- [Build Modes](https://pkg.go.dev/cmd/go#hdr-Build_modes)

---

**🎉 Parabéns! Você completou todas as aulas de tópicos avançados em Go!**


