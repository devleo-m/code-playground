# Módulo 15: Code Organization em Go

## Aula 1: Code Organization - Organização de Código e Gerenciamento de Dependências

Olá! Bem-vindo ao décimo quinto módulo. Até agora você aprendeu sobre funções, métodos, interfaces, generics, tratamento de erros e muitos outros conceitos fundamentais de Go. Mas você já se perguntou: **"Como organizo meu código em projetos maiores? Como gerencio dependências? Como compartilho meu código com outros desenvolvedores?"**

Em projetos pequenos, você pode colocar tudo em um único arquivo. Mas quando seu projeto cresce, você precisa de uma estrutura clara, um sistema de gerenciamento de dependências robusto e uma forma de organizar seu código em módulos reutilizáveis.

Nesta aula, vamos mergulhar profundamente na organização de código em Go: entender o sistema de módulos (Go modules), como trabalhar com packages, como usar bibliotecas de terceiros e como publicar seus próprios módulos.

---

## 1. Go Modules: O Sistema de Gerenciamento de Dependências

### O que são Go Modules?

**Go modules** são o sistema oficial de gerenciamento de dependências do Go, introduzido no Go 1.11 e tornando-se padrão no Go 1.13. Eles resolvem o problema de como gerenciar dependências externas de forma confiável e reproduzível.

### Por que Go Modules?

Antes dos modules, Go usava o sistema `GOPATH`, que tinha várias limitações:
- ❌ Todos os projetos precisavam estar em um diretório específico
- ❌ Dificuldade em gerenciar versões de dependências
- ❌ Problemas com builds reproduzíveis
- ❌ Conflitos entre diferentes projetos

**Go modules resolvem isso:**
- ✅ Cada projeto é independente
- ✅ Versionamento semântico (semantic versioning)
- ✅ Builds reproduzíveis
- ✅ Sem necessidade de GOPATH

### Estrutura de um Module

Um módulo Go é definido por um arquivo `go.mod` na raiz do projeto:

```
meu-projeto/
├── go.mod          # Define o módulo e suas dependências
├── go.sum          # Checksums para segurança
├── main.go
└── pkg/
    └── utils.go
```

---

## 2. `go mod init`: Inicializando um Novo Módulo

### O que faz?

O comando `go mod init` cria um novo módulo Go inicializando um arquivo `go.mod` com o caminho do módulo especificado.

### Sintaxe

```bash
go mod init <module-path>
```

O `module-path` geralmente é a URL do repositório onde o código será hospedado (por exemplo, `github.com/usuario/projeto`), mas para projetos locais pode ser qualquer nome único.

### Exemplo Prático

```bash
# Criar um novo projeto
mkdir meu-projeto
cd meu-projeto

# Inicializar o módulo
go mod init github.com/meu-usuario/meu-projeto
```

Isso cria um arquivo `go.mod`:

```go
module github.com/meu-usuario/meu-projeto

go 1.21
```

### Quando Usar?

- ✅ **Sempre** ao criar um novo projeto Go
- ✅ Ao converter um projeto antigo para usar modules
- ✅ Ao começar qualquer aplicação Go moderna

### Boas Práticas

1. **Use URLs de repositório**: Mesmo para projetos privados, use um caminho que reflita onde o código estaria hospedado
2. **Seja consistente**: Use o mesmo padrão de nomenclatura em todos os seus projetos
3. **Não use caminhos genéricos**: Evite nomes como `projeto` ou `teste`

```bash
# ❌ EVITE
go mod init projeto

# ✅ PREFIRA
go mod init github.com/seu-usuario/projeto
```

---

## 3. `go mod tidy`: Mantendo Dependências Organizadas

### O que faz?

O comando `go mod tidy` garante que o arquivo `go.mod` reflita exatamente as dependências usadas no código-fonte. Ele:
- Adiciona dependências que estão sendo usadas mas não estão no `go.mod`
- Remove dependências que estão no `go.mod` mas não estão sendo usadas
- Atualiza o arquivo `go.sum` com checksums

### Quando Usar?

Execute `go mod tidy`:
- ✅ Antes de fazer commit do código
- ✅ Após adicionar ou remover imports
- ✅ Após atualizar dependências manualmente
- ✅ Regularmente para manter o projeto limpo

### Exemplo Prático

```bash
# Adicionar uma dependência
go get github.com/gin-gonic/gin

# Limpar dependências não usadas
go mod tidy
```

**Antes do `go mod tidy`:**
```go
module meu-projeto

go 1.21

require (
    github.com/gin-gonic/gin v1.9.1
    github.com/alguma-lib v1.0.0  // Não está sendo usada!
)
```

**Depois do `go mod tidy`:**
```go
module meu-projeto

go 1.21

require (
    github.com/gin-gonic/gin v1.9.1
    // Dependência não usada foi removida automaticamente
)
```

### Importância do `go.sum`

O arquivo `go.sum` contém checksums criptográficos de todas as dependências. Ele garante:
- **Segurança**: Detecta se uma dependência foi modificada
- **Reproduzibilidade**: Garante que todos baixem as mesmas versões
- **Integridade**: Previne ataques de supply chain

**Nunca delete o `go.sum`!** Ele é essencial para builds seguros e reproduzíveis.

---

## 4. `go mod vendor`: Criando Vendor Directory

### O que faz?

O comando `go mod vendor` cria um diretório `vendor/` contendo cópias de todas as dependências do projeto. Isso permite builds sem acesso à internet.

### Quando Usar?

Use `go mod vendor` quando:
- ✅ Você precisa fazer builds em ambientes sem internet (air-gapped)
- ✅ Você quer controle total sobre as versões exatas das dependências
- ✅ Você está fazendo deploy e quer garantir que todas as dependências estejam incluídas
- ✅ Você quer fazer modificações locais em dependências

### Exemplo Prático

```bash
# Criar o diretório vendor
go mod vendor
```

Isso cria uma estrutura como:

```
meu-projeto/
├── go.mod
├── go.sum
├── main.go
└── vendor/
    ├── github.com/
    │   └── gin-gonic/
    │       └── gin/
    │           └── ...
    └── modules.txt
```

### Como Funciona?

Quando você usa `go mod vendor`, o Go:
1. Baixa todas as dependências listadas no `go.mod`
2. Copia os arquivos para o diretório `vendor/`
3. Cria um arquivo `modules.txt` com informações sobre as dependências

### Build com Vendor

Para usar o diretório `vendor` durante o build:

```bash
go build -mod=vendor
```

Ou defina a variável de ambiente:

```bash
export GOFLAGS=-mod=vendor
go build
```

### Quando NÃO Usar Vendor?

- ❌ Para projetos pequenos ou pessoais (desnecessário)
- ❌ Quando você tem acesso constante à internet
- ❌ Para desenvolvimento local normal (o Go baixa automaticamente)

**Nota**: O diretório `vendor/` geralmente é grande e não deve ser versionado no Git para projetos que não precisam dele. Use `.gitignore`:

```
vendor/
```

---

## 5. Packages: A Unidade Fundamental de Organização

### O que são Packages?

**Packages** (pacotes) são a unidade fundamental de organização de código em Go. Um package agrupa funções, tipos, variáveis e constantes relacionadas.

### Estrutura de um Package

Cada arquivo Go começa com uma declaração de package:

```go
package nome_do_package

// código do package
```

### Tipos de Packages

#### 1. Package `main`

O package `main` é especial: ele define um programa executável.

```go
package main

import "fmt"

func main() {
    fmt.Println("Olá, mundo!")
}
```

**Características:**
- Deve ter uma função `main()`
- Quando compilado, gera um executável
- Não pode ser importado por outros packages

#### 2. Packages Comuns

Qualquer outro nome de package cria uma biblioteca que pode ser importada:

```go
// arquivo: math/utils.go
package math

func Soma(a, b int) int {
    return a + b
}
```

### Exportação: Nomes Públicos vs Privados

Em Go, a exportação é determinada pela **primeira letra do nome**:

- **Maiúscula** = Exportado (público) - pode ser usado por outros packages
- **Minúscula** = Não exportado (privado) - só pode ser usado dentro do mesmo package

```go
package calculadora

// ✅ Exportado - pode ser usado por outros packages
func Soma(a, b int) int {
    return a + b
}

// ❌ Não exportado - só pode ser usado dentro deste package
func subtracao(a, b int) int {
    return a - b
}

// ✅ Exportado
var Pi = 3.14159

// ❌ Não exportado
var versao = "1.0.0"
```

### Importando Packages

Use a declaração `import` para usar packages:

```go
package main

import (
    "fmt"           // Package da biblioteca padrão
    "math"          // Package da biblioteca padrão
    "os"            // Package da biblioteca padrão
)

// Ou imports individuais
import "fmt"
import "math"
```

### Import com Alias

Você pode dar um alias a um package:

```go
import (
    f "fmt"         // Agora uso f.Println() ao invés de fmt.Println()
    m "math"
)

func main() {
    f.Println("Usando alias!")
}
```

### Import com Ponto

Importar com ponto permite usar funções sem o prefixo do package (use com cuidado!):

```go
import . "fmt"

func main() {
    Println("Sem prefixo fmt!")  // Ao invés de fmt.Println()
}
```

### Import Anônimo

Importar apenas para efeitos colaterais (como inicialização):

```go
import _ "database/sql/driver"  // Registra o driver sem usar diretamente
```

---

## 6. Package Import Rules: Regras de Importação

### Regras Fundamentais

#### 1. Sem Imports Circulares

Go **não permite** imports circulares. Se o package A importa B, B não pode importar A.

```go
// ❌ ERRADO - Import circular
// package A
import "B"

// package B
import "A"  // ERRO!
```

**Solução**: Reorganize o código ou crie um package intermediário.

#### 2. Package `main` é Especial

- O package `main` é para executáveis
- Não pode ser importado por outros packages
- Deve ter uma função `main()`

#### 3. Nomes de Packages em Minúsculas

Por convenção, nomes de packages são sempre em **minúsculas**:

```go
// ✅ CORRETO
package calculadora
package utils
package httpclient

// ❌ EVITE
package Calculadora
package Utils
```

#### 4. Identificadores Exportados Começam com Maiúscula

```go
package exemplo

// ✅ Exportado
func FuncaoPublica() {}

// ❌ Não exportado
func funcaoPrivada() {}
```

#### 5. Caminhos de Import são Identificadores Únicos

O caminho de import identifica unicamente um package:

```go
import (
    "fmt"                           // Biblioteca padrão
    "github.com/gin-gonic/gin"     // Package externo
    "meu-projeto/utils"             // Package local
)
```

### Organização de Packages em um Projeto

Estrutura recomendada:

```
meu-projeto/
├── go.mod
├── main.go                 # package main
├── cmd/
│   └── server/
│       └── main.go         # package main (outro executável)
├── pkg/
│   ├── models/
│   │   └── user.go         # package models
│   ├── handlers/
│   │   └── user.go         # package handlers
│   └── utils/
│       └── helpers.go      # package utils
└── internal/
    └── database/
        └── db.go           # package database (só acessível pelo módulo)
```

### Package `internal`

O package `internal` é especial: packages dentro de `internal/` só podem ser importados por packages do mesmo módulo:

```go
// internal/database/db.go
package database

func Connect() {}  // Só pode ser usado dentro deste módulo
```

Isso é útil para código que você não quer expor publicamente.

---

## 7. Usando Packages de Terceiros

### Adicionando Dependências

Use `go get` para adicionar dependências:

```bash
# Adicionar uma dependência
go get github.com/gin-gonic/gin

# Adicionar versão específica
go get github.com/gin-gonic/gin@v1.9.1

# Adicionar versão mais recente
go get github.com/gin-gonic/gin@latest
```

### Exemplo Prático

```go
package main

import (
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()
    r.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{"message": "Olá!"})
    })
    r.Run()
}
```

Depois de adicionar o import, execute:

```bash
go mod tidy  # Atualiza go.mod e go.sum
```

### Escolhendo Packages de Terceiros

Ao escolher uma biblioteca, considere:

1. **Status de Manutenção**
   - ✅ Última atualização recente
   - ✅ Issues sendo resolvidas
   - ✅ Pull requests sendo revisados

2. **Documentação**
   - ✅ README claro
   - ✅ Exemplos de uso
   - ✅ Documentação da API

3. **Licença**
   - Verifique se a licença é compatível com seu projeto
   - MIT, Apache 2.0 são geralmente seguras

4. **Segurança**
   - Verifique vulnerabilidades conhecidas
   - Use `go list -m -u all` para ver atualizações disponíveis

5. **Popularidade**
   - Muitas estrelas no GitHub
   - Usado por projetos conhecidos
   - Comunidade ativa

### Atualizando Dependências

```bash
# Ver dependências desatualizadas
go list -m -u all

# Atualizar uma dependência específica
go get -u github.com/gin-gonic/gin

# Atualizar todas as dependências
go get -u ./...

# Atualizar para versão específica
go get github.com/gin-gonic/gin@v1.10.0
```

### Versionamento Semântico

Go modules usa **semantic versioning** (semver):

- **v1.2.3** = Major.Minor.Patch
- **v1.2.3-beta.1** = Versão pré-lançamento
- **v1.2.3+metadata** = Metadados adicionais

O Go escolhe automaticamente a versão mais recente compatível baseado no `go.mod`.

---

## 8. Publicando Módulos

### Como Publicar seu Módulo

Para publicar um módulo Go, você precisa:

1. **Ter um repositório Git**
2. **Usar tags de versão semântica**
3. **Ter um `go.mod` válido**

### Passo a Passo

#### 1. Preparar o Código

```bash
# Inicializar o módulo
go mod init github.com/seu-usuario/meu-modulo

# Desenvolver seu código
# ...

# Garantir que está tudo limpo
go mod tidy
```

#### 2. Fazer Commit e Push

```bash
git init
git add .
git commit -m "Versão inicial"
git remote add origin https://github.com/seu-usuario/meu-modulo.git
git push -u origin main
```

#### 3. Criar Tag de Versão

```bash
# Criar tag para v1.0.0
git tag v1.0.0
git push origin v1.0.0
```

#### 4. O Go Proxy Descobre Automaticamente

O Go proxy (proxy.golang.org) automaticamente descobre e serve seu módulo. Em alguns minutos, outros desenvolvedores poderão fazer:

```bash
go get github.com/seu-usuario/meu-modulo@v1.0.0
```

### Boas Práticas para Publicar

#### 1. Siga Convenções do Go

- ✅ Nomes de packages em minúsculas
- ✅ Documentação clara
- ✅ Exemplos de uso
- ✅ README completo

#### 2. Versionamento Semântico

- **MAJOR** (v1 → v2): Mudanças incompatíveis
- **MINOR** (v1.0 → v1.1): Novas funcionalidades compatíveis
- **PATCH** (v1.0.0 → v1.0.1): Correções de bugs

#### 3. Compatibilidade Retroativa

- ✅ Mantenha APIs estáveis
- ✅ Adicione novas funcionalidades sem quebrar as antigas
- ✅ Documente mudanças breaking

#### 4. Documentação

```go
// Package calculadora fornece funções matemáticas básicas.
//
// Exemplo de uso:
//
//     resultado := calculadora.Soma(2, 3)
//     fmt.Println(resultado) // Output: 5
package calculadora

// Soma retorna a soma de dois inteiros.
func Soma(a, b int) int {
    return a + b
}
```

#### 5. Módulos Privados

Para módulos privados (não públicos), você pode:
- Usar um proxy privado
- Configurar `GOPRIVATE`:

```bash
export GOPRIVATE=github.com/minha-empresa/*
```

### Módulos com Versão Major v2+

Quando você precisa fazer mudanças incompatíveis (v2+), você tem duas opções:

#### Opção 1: Caminho de Módulo com Versão

```
github.com/usuario/modulo/v2
```

#### Opção 2: Branch Separado

Manter v1 e v2 em branches diferentes.

---

## 9. Comandos Úteis do Go Modules

### `go mod download`

Baixa dependências sem compilar:

```bash
go mod download
```

### `go mod graph`

Mostra o grafo de dependências:

```bash
go mod graph
```

### `go mod verify`

Verifica a integridade das dependências:

```bash
go mod verify
```

### `go mod why`

Explica por que uma dependência é necessária:

```bash
go mod why github.com/gin-gonic/gin
```

### `go list -m all`

Lista todas as dependências (diretas e indiretas):

```bash
go list -m all
```

### `go list -m -versions`

Lista todas as versões disponíveis de um módulo:

```bash
go list -m -versions github.com/gin-gonic/gin
```

---

## 10. Estrutura de Projeto Recomendada

### Estrutura Básica

```
meu-projeto/
├── go.mod
├── go.sum
├── README.md
├── .gitignore
├── cmd/
│   ├── server/
│   │   └── main.go        # Executável do servidor
│   └── cli/
│       └── main.go        # Executável CLI
├── pkg/
│   ├── models/
│   │   └── user.go
│   ├── handlers/
│   │   └── user.go
│   └── utils/
│       └── helpers.go
├── internal/
│   └── database/
│       └── db.go
├── api/
│   └── routes.go
├── config/
│   └── config.go
└── tests/
    └── user_test.go
```

### Explicação dos Diretórios

- **`cmd/`**: Executáveis (cada subdiretório é um executável)
- **`pkg/`**: Código que pode ser usado por outros projetos
- **`internal/`**: Código privado do módulo
- **`api/`**: Definições de API
- **`config/`**: Configurações
- **`tests/`**: Testes de integração

---

## Resumo da Aula

Nesta aula, você aprendeu:

1. **Go Modules**: Sistema de gerenciamento de dependências do Go
2. **`go mod init`**: Inicializar um novo módulo
3. **`go mod tidy`**: Manter dependências organizadas
4. **`go mod vendor`**: Criar diretório vendor para builds offline
5. **Packages**: Unidade fundamental de organização de código
6. **Regras de Import**: Sem imports circulares, exportação por maiúsculas
7. **Packages de Terceiros**: Como adicionar e usar bibliotecas externas
8. **Publicar Módulos**: Como compartilhar seu código com a comunidade
9. **Estrutura de Projeto**: Organização recomendada para projetos Go

**Próximos passos:**
- Na próxima aula, vamos simplificar esses conceitos com analogias e exemplos do cotidiano
- Depois, exercícios práticos para fixar o aprendizado
- Por fim, boas práticas e otimizações de performance

Até a próxima!

