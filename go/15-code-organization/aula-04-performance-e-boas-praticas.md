# Aula 4 - Performance e Boas Práticas: Code Organization em Go

Olá! Agora que você entende os conceitos de organização de código em Go, é crucial aprender **quando e como** organizar seu código de forma eficiente e seguir as melhores práticas da comunidade Go. Nesta aula, vamos explorar aspectos de performance, boas práticas, padrões comuns, e os erros que você deve evitar.

---

## 🚀 Performance: Impacto da Organização de Código

### Packages e Compilação

**Ponto crucial:** A organização em packages afeta o tempo de compilação e o tamanho dos binários.

✅ **Compilação Incremental** - Go compila apenas packages que mudaram
✅ **Paralelização** - Packages independentes são compilados em paralelo
✅ **Cache de Build** - Go cacheia resultados de compilação
✅ **Tree Shaking** - Código não usado é eliminado automaticamente

### Estrutura de Packages e Performance

**Estrutura com muitos packages pequenos:**
```
projeto/
├── pkg/
│   ├── math/
│   │   ├── soma.go
│   │   └── subtracao.go
│   ├── string/
│   │   └── utils.go
│   └── ...
```

**Vantagens:**
- ✅ Compilação mais rápida (apenas o que mudou)
- ✅ Testes mais rápidos (testa apenas packages modificados)
- ✅ Melhor paralelização

**Desvantagens:**
- ⚠️ Mais arquivos para gerenciar
- ⚠️ Pode ser over-engineering para projetos pequenos

**Estrutura com poucos packages grandes:**
```
projeto/
├── pkg/
│   └── utils.go  # Tudo em um arquivo
```

**Vantagens:**
- ✅ Simples para projetos pequenos
- ✅ Menos arquivos

**Desvantagens:**
- ❌ Recompila tudo mesmo com mudanças pequenas
- ❌ Menos paralelização
- ❌ Dificulta testes isolados

### Recomendação

**Para projetos pequenos (< 10 arquivos):**
- Use poucos packages ou até mesmo um único package
- Simplicidade > Otimização prematura

**Para projetos médios/grandes:**
- Organize em packages lógicos
- Balance entre granularidade e simplicidade
- Packages devem ter responsabilidades claras

---

## ✅ Boas Práticas: Organização de Packages

### ✅ SEMPRE: Use Nomes de Packages Descritivos

```go
// ❌ ERRADO: Nome genérico
package utils
package helpers
package common

// ✅ CORRETO: Nome específico
package stringutils
package httpclient
package database
```

**Por quê?**
- Nomes genéricos não dizem o que o package faz
- Quando você importa `utils`, não fica claro o que está disponível
- Nomes específicos são auto-documentados

### ✅ SEMPRE: Mantenha Packages Coesos

Um package deve ter uma **responsabilidade única e clara**:

```go
// ❌ ERRADO: Package faz muitas coisas
package utils

func Soma(a, b int) int { ... }
func ValidarEmail(email string) bool { ... }
func ConectarDB() *sql.DB { ... }
func EnviarEmail() { ... }

// ✅ CORRETO: Packages separados por responsabilidade
package math
func Soma(a, b int) int { ... }

package validation
func ValidarEmail(email string) bool { ... }

package database
func Conectar() *sql.DB { ... }

package email
func Enviar() { ... }
```

### ✅ SEMPRE: Evite Packages com Apenas um Arquivo

Se um package tem apenas um arquivo, considere se ele realmente precisa ser um package separado:

```go
// ❌ EVITE: Package com um único arquivo
// math/soma.go
package math
func Soma(a, b int) int { return a + b }

// ✅ PREFIRA: Combinar com package relacionado
// math/operacoes.go
package math
func Soma(a, b int) int { return a + b }
func Subtracao(a, b int) int { return a - b }
```

**Exceção:** Se o package é grande e você está organizando em múltiplos arquivos por funcionalidade.

### ✅ SEMPRE: Use `internal/` para Código Privado

```go
// ✅ CORRETO: Código que não deve ser usado externamente
internal/
└── database/
    └── connection.go  // Só acessível dentro do módulo
```

**Por quê?**
- Previne uso acidental de código interno
- Facilita refatoração sem quebrar dependências externas
- Documenta intenção: "isso é privado"

---

## ✅ Boas Práticas: Go Modules

### ✅ SEMPRE: Execute `go mod tidy` Antes de Commits

```bash
# ✅ BOM HÁBITO
git add .
go mod tidy  # Limpar dependências
git add go.mod go.sum
git commit -m "Adiciona nova funcionalidade"
```

**Por quê?**
- Mantém `go.mod` limpo e preciso
- Remove dependências não usadas
- Garante que `go.sum` está atualizado

### ✅ SEMPRE: Versionar `go.sum`

```bash
# ✅ CORRETO: Versionar ambos
git add go.mod go.sum
git commit -m "Atualiza dependências"
```

**Por quê?**
- `go.sum` garante builds reproduzíveis
- Previne ataques de supply chain
- É essencial para segurança

### ✅ SEMPRE: Use Versões Específicas em Produção

```go
// ❌ EVITE: Versão "latest" em produção
go get github.com/gin-gonic/gin@latest

// ✅ PREFIRA: Versão específica
go get github.com/gin-gonic/gin@v1.9.1
```

**Por quê?**
- `@latest` pode quebrar seu código com atualizações
- Versões específicas garantem builds reproduzíveis
- Facilita debugging (você sabe exatamente qual versão)

### ✅ SEMPRE: Documente Dependências Críticas

```go
// go.mod
module meu-projeto

go 1.21

require (
    // Core framework - versão estável
    github.com/gin-gonic/gin v1.9.1
    
    // Database driver - versão LTS
    github.com/lib/pq v1.10.9
)
```

Adicione comentários explicando por que dependências críticas estão em versões específicas.

---

## ✅ Boas Práticas: Estrutura de Projeto

### ✅ Estrutura Recomendada para Projetos Pequenos

```
projeto/
├── go.mod
├── go.sum
├── main.go
├── handlers.go
├── models.go
└── utils.go
```

**Quando usar:**
- Projetos pessoais pequenos
- Scripts e ferramentas simples
- Aprendizado e prototipagem

### ✅ Estrutura Recomendada para Projetos Médios

```
projeto/
├── go.mod
├── go.sum
├── cmd/
│   └── server/
│       └── main.go
├── pkg/
│   ├── handlers/
│   ├── models/
│   └── utils/
└── internal/
    └── config/
```

**Quando usar:**
- APIs e serviços
- Aplicações com múltiplos componentes
- Projetos que podem crescer

### ✅ Estrutura Recomendada para Projetos Grandes

```
projeto/
├── go.mod
├── go.sum
├── cmd/
│   ├── api/
│   ├── worker/
│   └── cli/
├── pkg/
│   ├── api/
│   ├── domain/
│   └── infrastructure/
├── internal/
│   ├── config/
│   └── database/
├── api/
│   └── proto/  # Se usar gRPC
└── deployments/
    └── k8s/
```

**Quando usar:**
- Microsserviços
- Sistemas distribuídos
- Projetos empresariais

### ✅ NÃO: Crie Estrutura Complexa Prematuramente

```go
// ❌ ERRADO: Over-engineering desde o início
projeto-pequeno/
├── cmd/
│   └── app/
│       └── main.go
├── pkg/
│   ├── domain/
│   │   ├── entities/
│   │   ├── repositories/
│   │   └── services/
│   └── infrastructure/
│       ├── persistence/
│       └── messaging/
└── ...

// ✅ CORRETO: Comece simples, evolua conforme necessário
projeto-pequeno/
├── go.mod
├── main.go
└── models.go
```

**Regra de ouro:** Comece simples. Reorganize quando a complexidade justificar.

---

## ✅ Boas Práticas: Imports

### ✅ SEMPRE: Organize Imports

```go
// ✅ CORRETO: Organizado por grupo
import (
    // Biblioteca padrão
    "fmt"
    "os"
    "time"
    
    // Dependências externas
    "github.com/gin-gonic/gin"
    "github.com/lib/pq"
    
    // Packages locais
    "meu-projeto/pkg/models"
    "meu-projeto/pkg/handlers"
)
```

**Dica:** Use `goimports` (ou configure seu editor) para organizar automaticamente:

```bash
go install golang.org/x/tools/cmd/goimports@latest
goimports -w .
```

### ✅ EVITE: Imports Não Usados

```go
// ❌ ERRADO: Import não usado
import (
    "fmt"
    "os"  // Não está sendo usado!
)

// ✅ CORRETO: Apenas o que é usado
import "fmt"
```

**Dica:** O compilador Go detecta isso, mas é melhor manter limpo desde o início.

### ✅ EVITE: Import com Ponto (`.`)

```go
// ❌ EVITE: Import com ponto
import . "fmt"

func main() {
    Println("Sem prefixo")  // Confuso - de onde vem?
}

// ✅ PREFIRA: Import normal
import "fmt"

func main() {
    fmt.Println("Claro de onde vem")
}
```

**Por quê?**
- Torna o código menos claro
- Pode causar conflitos de nomes
- Dificulta leitura e manutenção

**Exceção:** Às vezes usado em testes para facilitar, mas ainda assim é questionável.

---

## ✅ Boas Práticas: Exportação

### ✅ SEMPRE: Exporte Apenas o Necessário

```go
// ❌ ERRADO: Exportar tudo
package utils

func Funcao1() { ... }
func Funcao2() { ... }
func Funcao3() { ... }
func Funcao4() { ... }
func Funcao5() { ... }

// ✅ CORRETO: Exportar apenas a API pública
package utils

// API pública
func Processar() { ... }

// Implementação privada
func processarInterno() { ... }
func validar() { ... }
```

**Por quê?**
- API menor = mais fácil de usar
- Facilita refatoração (código privado pode mudar)
- Previne uso incorreto de funções internas

### ✅ SEMPRE: Documente APIs Exportadas

```go
// ❌ ERRADO: Sem documentação
func Processar(dados string) error {
    // ...
}

// ✅ CORRETO: Documentado
// Processar processa os dados fornecidos e retorna um erro
// se a validação falhar.
func Processar(dados string) error {
    // ...
}
```

**Dica:** Use `godoc` para gerar documentação:

```bash
godoc -http=:6060
```

---

## ⚠️ Erros Comuns e Como Evitá-los

### ❌ ERRO 1: Imports Circulares

```go
// ❌ ERRADO
// package a
import "b"

// package b
import "a"  // ERRO!
```

**Solução:**
1. Reorganize código compartilhado em um package separado
2. Use interfaces para quebrar dependências
3. Mova funcionalidade para um package comum

### ❌ ERRO 2: Package com Responsabilidades Múltiplas

```go
// ❌ ERRADO: Package faz muitas coisas
package utils

func ValidarEmail() { ... }
func ConectarDB() { ... }
func EnviarEmail() { ... }
func CalcularImposto() { ... }
```

**Solução:** Divida em packages com responsabilidades claras:

```go
package validation
func ValidarEmail() { ... }

package database
func Conectar() { ... }

package email
func Enviar() { ... }

package finance
func CalcularImposto() { ... }
```

### ❌ ERRO 3: Não Usar `go mod tidy`

```go
// ❌ ERRADO: go.mod desatualizado
// Dependências não usadas acumulam
// go.sum pode estar desatualizado
```

**Solução:** Execute `go mod tidy` regularmente, especialmente antes de commits.

### ❌ ERRO 4: Estrutura de Projeto Inconsistente

```go
// ❌ ERRADO: Estrutura confusa
projeto/
├── src/
│   ├── main.go
│   └── utils/
├── lib/
│   └── helpers.go
└── pkg/
    └── models.go
```

**Solução:** Siga uma estrutura consistente. Use a estrutura padrão do Go ou defina uma e mantenha.

### ❌ ERRO 5: Nomes de Packages Confusos

```go
// ❌ ERRADO
package pkg
package lib
package common
package util
```

**Solução:** Use nomes descritivos que indiquem o propósito:

```go
package httpclient
package stringutils
package database
```

---

## 🎯 Padrões Avançados

### Padrão: Package de Configuração

```go
// internal/config/config.go
package config

import "os"

type Config struct {
    Port     string
    Database string
}

func Load() *Config {
    return &Config{
        Port:     getEnv("PORT", "8080"),
        Database: getEnv("DATABASE_URL", ""),
    }
}

func getEnv(key, defaultValue string) string {
    if value := os.Getenv(key); value != "" {
        return value
    }
    return defaultValue
}
```

**Uso:**
```go
config := config.Load()
// Usar config.Port, config.Database
```

### Padrão: Package de Erros

```go
// pkg/errors/errors.go
package errors

import "fmt"

var (
    ErrNotFound     = fmt.Errorf("não encontrado")
    ErrInvalidInput = fmt.Errorf("entrada inválida")
    ErrUnauthorized = fmt.Errorf("não autorizado")
)
```

**Uso:**
```go
if err != nil {
    if errors.Is(err, errors.ErrNotFound) {
        // Tratar especificamente
    }
}
```

### Padrão: Factory Functions

```go
// pkg/database/database.go
package database

type DB struct {
    // campos privados
}

// NewDB cria uma nova instância de DB
func NewDB(connectionString string) (*DB, error) {
    // lógica de inicialização
    return &DB{}, nil
}
```

**Por quê?**
- Força inicialização correta
- Pode retornar erros
- Esconde detalhes de implementação

---

## 📊 Métricas e Monitoramento

### Tamanho do Projeto

```bash
# Ver tamanho do projeto
du -sh .

# Ver tamanho de vendor (se usado)
du -sh vendor/

# Ver número de packages
find . -name "*.go" -not -path "./vendor/*" | xargs grep -l "^package " | wc -l
```

### Dependências

```bash
# Listar todas as dependências
go list -m all

# Ver dependências desatualizadas
go list -m -u all

# Ver por que uma dependência é necessária
go mod why github.com/alguma-lib
```

### Tempo de Compilação

```bash
# Medir tempo de compilação
time go build

# Compilar com informações detalhadas
go build -x
```

---

## 🔍 Ferramentas Úteis

### `goimports`

Organiza imports automaticamente:

```bash
go install golang.org/x/tools/cmd/goimports@latest
goimports -w .
```

### `golangci-lint`

Linter completo que verifica muitas coisas, incluindo organização:

```bash
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
golangci-lint run
```

### `modgraphviz`

Visualiza o grafo de dependências:

```bash
go install github.com/kisielk/godepgraph@latest
godepgraph . | dot -Tpng -o deps.png
```

---

## Resumo das Boas Práticas

### ✅ FAÇA

1. ✅ Use nomes de packages descritivos
2. ✅ Mantenha packages coesos (uma responsabilidade)
3. ✅ Execute `go mod tidy` regularmente
4. ✅ Versionar `go.sum`
5. ✅ Use `internal/` para código privado
6. ✅ Organize imports por grupos
7. ✅ Documente APIs exportadas
8. ✅ Comece simples, evolua conforme necessário

### ❌ EVITE

1. ❌ Nomes genéricos (`utils`, `common`)
2. ❌ Packages com múltiplas responsabilidades
3. ❌ Ignorar `go mod tidy`
4. ❌ Não versionar `go.sum`
5. ❌ Estrutura complexa prematura
6. ❌ Imports não usados
7. ❌ Exportar tudo
8. ❌ Imports circulares

---

## Conclusão

Organização de código não é apenas sobre estrutura - é sobre:
- **Manutenibilidade**: Código fácil de entender e modificar
- **Escalabilidade**: Estrutura que cresce com o projeto
- **Colaboração**: Código que outros desenvolvedores entendem
- **Performance**: Organização que facilita compilação rápida

Lembre-se: **Comece simples. Evolua conforme necessário. Mantenha consistente.**

Boa sorte com seus projetos Go! 🚀

