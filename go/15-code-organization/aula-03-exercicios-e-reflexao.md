# Aula 3 - Exercícios e Reflexão: Code Organization em Go

Olá! Agora é hora de colocar a mão na massa e praticar tudo que aprendemos sobre organização de código em Go. Vamos começar com exercícios práticos e depois refletir sobre os conceitos.

---

## Exercício 1: Criando seu Primeiro Módulo

### Objetivo
Criar um módulo Go do zero e entender a estrutura básica.

### Tarefa
1. Crie um novo diretório chamado `meu-primeiro-modulo`
2. Inicialize um módulo Go com o caminho `github.com/seu-usuario/meu-primeiro-modulo`
3. Crie um arquivo `main.go` com um programa simples que imprime "Olá, Módulo!"
4. Execute `go mod tidy`
5. Compile e execute o programa

### Passos Detalhados

```bash
# 1. Criar diretório
mkdir meu-primeiro-modulo
cd meu-primeiro-modulo

# 2. Inicializar módulo
go mod init github.com/seu-usuario/meu-primeiro-modulo

# 3. Criar main.go (você escreve o código)
# 4. Executar go mod tidy
go mod tidy

# 5. Compilar e executar
go run main.go
```

### Código Sugerido para `main.go`

```go
package main

import "fmt"

func main() {
    fmt.Println("Olá, Módulo!")
    fmt.Println("Meu primeiro módulo Go está funcionando!")
}
```

### Reflexão
- O que o arquivo `go.mod` contém?
- Por que usamos uma URL do GitHub mesmo para um projeto local?
- O que acontece se você executar `go mod tidy` várias vezes?

---

## Exercício 2: Trabalhando com Packages

### Objetivo
Criar e usar packages personalizados em seu projeto.

### Tarefa
Crie uma estrutura de projeto com:
1. Um package `calculadora` com funções públicas e privadas
2. Um package `main` que usa o package `calculadora`

### Estrutura do Projeto

```
meu-projeto/
├── go.mod
├── main.go
└── calculadora/
    └── calculadora.go
```

### Código para `calculadora/calculadora.go`

```go
package calculadora

// Soma retorna a soma de dois números
func Soma(a, b int) int {
    return a + b
}

// Subtracao retorna a subtração de dois números
func Subtracao(a, b int) int {
    return a - b
}

// multiplicacao é privada (não exportada)
func multiplicacao(a, b int) int {
    return a * b
}

// Multiplicar usa a função privada multiplicacao
func Multiplicar(a, b int) int {
    return multiplicacao(a, b)
}
```

### Código para `main.go`

```go
package main

import (
    "fmt"
    "meu-projeto/calculadora"  // Ajuste o caminho conforme seu go.mod
)

func main() {
    resultado := calculadora.Soma(10, 5)
    fmt.Printf("10 + 5 = %d\n", resultado)
    
    resultado = calculadora.Subtracao(10, 5)
    fmt.Printf("10 - 5 = %d\n", resultado)
    
    resultado = calculadora.Multiplicar(10, 5)
    fmt.Printf("10 * 5 = %d\n", resultado)
    
    // Tente descomentar a linha abaixo - o que acontece?
    // resultado = calculadora.multiplicacao(10, 5)  // ERRO!
}
```

### Desafio Extra
1. Tente chamar `calculadora.multiplicacao()` diretamente. O que acontece?
2. Crie uma função privada `divisao()` e uma função pública `Dividir()` que a usa.
3. Adicione validação para evitar divisão por zero.

### Reflexão
- Por que algumas funções começam com maiúscula e outras com minúscula?
- O que acontece se você tentar usar uma função privada de outro package?
- Como você organizaria funções relacionadas em packages diferentes?

---

## Exercício 3: Adicionando Dependências Externas

### Objetivo
Aprender a adicionar e usar packages de terceiros.

### Tarefa
1. Crie um novo projeto
2. Adicione a biblioteca `github.com/fatih/color` para imprimir texto colorido
3. Use a biblioteca para imprimir mensagens coloridas
4. Execute `go mod tidy` e observe as mudanças no `go.mod`

### Passos

```bash
# 1. Criar projeto
mkdir projeto-colorido
cd projeto-colorido
go mod init github.com/seu-usuario/projeto-colorido

# 2. Adicionar dependência
go get github.com/fatih/color

# 3. Criar main.go (veja código abaixo)
# 4. Executar
go run main.go
```

### Código para `main.go`

```go
package main

import (
    "github.com/fatih/color"
)

func main() {
    color.Red("Esta mensagem está em vermelho!")
    color.Green("Esta mensagem está em verde!")
    color.Blue("Esta mensagem está em azul!")
    
    color.Cyan("Você pode combinar estilos:")
    color.New(color.FgYellow, color.Bold).Println("Texto amarelo e negrito!")
}
```

### Desafio Extra
1. Explore outras funções da biblioteca `color`
2. Crie uma função que imprime mensagens de sucesso (verde) e erro (vermelho)
3. Adicione outra biblioteca de sua escolha e experimente

### Reflexão
- O que foi adicionado ao arquivo `go.mod`?
- O que é o arquivo `go.sum` e por que ele é importante?
- Como você sabe qual versão de uma biblioteca está sendo usada?

---

## Exercício 4: Organizando um Projeto Real

### Objetivo
Criar uma estrutura de projeto organizada seguindo as melhores práticas.

### Tarefa
Crie um projeto de "Gerenciador de Tarefas" com a seguinte estrutura:

```
gerenciador-tarefas/
├── go.mod
├── go.sum
├── README.md
├── cmd/
│   └── cli/
│       └── main.go          # Interface de linha de comando
├── pkg/
│   ├── models/
│   │   └── tarefa.go        # Modelo de Tarefa
│   ├── storage/
│   │   └── storage.go       # Armazenamento de tarefas
│   └── utils/
│       └── validacao.go     # Funções de validação
└── internal/
    └── config/
        └── config.go        # Configurações internas
```

### Implementação Sugerida

#### `pkg/models/tarefa.go`

```go
package models

// Tarefa representa uma tarefa no sistema
type Tarefa struct {
    ID          int
    Titulo      string
    Descricao   string
    Concluida   bool
}

// NovaTarefa cria uma nova tarefa
func NovaTarefa(id int, titulo, descricao string) *Tarefa {
    return &Tarefa{
        ID:        id,
        Titulo:    titulo,
        Descricao: descricao,
        Concluida: false,
    }
}

// Concluir marca a tarefa como concluída
func (t *Tarefa) Concluir() {
    t.Concluida = true
}
```

#### `pkg/storage/storage.go`

```go
package storage

import "gerenciador-tarefas/pkg/models"

// Storage gerencia o armazenamento de tarefas
type Storage struct {
    tarefas []*models.Tarefa
}

// NovoStorage cria um novo storage
func NovoStorage() *Storage {
    return &Storage{
        tarefas: make([]*models.Tarefa, 0),
    }
}

// Adicionar adiciona uma nova tarefa
func (s *Storage) Adicionar(tarefa *models.Tarefa) {
    s.tarefas = append(s.tarefas, tarefa)
}

// Listar retorna todas as tarefas
func (s *Storage) Listar() []*models.Tarefa {
    return s.tarefas
}
```

#### `cmd/cli/main.go`

```go
package main

import (
    "fmt"
    "gerenciador-tarefas/pkg/models"
    "gerenciador-tarefas/pkg/storage"
)

func main() {
    s := storage.NovoStorage()
    
    tarefa1 := models.NovaTarefa(1, "Aprender Go", "Estudar packages e modules")
    tarefa2 := models.NovaTarefa(2, "Fazer exercícios", "Praticar code organization")
    
    s.Adicionar(tarefa1)
    s.Adicionar(tarefa2)
    
    fmt.Println("Tarefas:")
    for _, tarefa := range s.Listar() {
        status := "Pendente"
        if tarefa.Concluida {
            status = "Concluída"
        }
        fmt.Printf("[%d] %s - %s (%s)\n", 
            tarefa.ID, tarefa.Titulo, tarefa.Descricao, status)
    }
}
```

### Desafio Extra
1. Adicione funcionalidade para marcar tarefas como concluídas
2. Adicione validação no package `utils`
3. Adicione persistência (salvar em arquivo)
4. Crie testes para os packages

### Reflexão
- Por que separamos o código em diferentes packages?
- Qual a diferença entre `pkg/` e `internal/`?
- Como a estrutura ajuda na manutenção do código?

---

## Exercício 5: Trabalhando com `go mod tidy` e `go mod vendor`

### Objetivo
Praticar o gerenciamento de dependências.

### Tarefa
1. Crie um projeto que usa várias dependências
2. Adicione uma dependência que não será usada
3. Use `go mod tidy` para limpar
4. Crie um diretório `vendor/` e explore seu conteúdo

### Passos

```bash
# 1. Criar projeto
mkdir projeto-deps
cd projeto-deps
go mod init github.com/seu-usuario/projeto-deps

# 2. Adicionar dependências (algumas serão usadas, outras não)
go get github.com/fatih/color
go get github.com/gin-gonic/gin
go get github.com/gorilla/mux  # Esta não será usada

# 3. Criar main.go usando apenas 'color'
# 4. Executar go mod tidy
go mod tidy

# 5. Verificar go.mod - gorilla/mux deve ter sido removido
# 6. Criar vendor
go mod vendor

# 7. Explorar o diretório vendor/
ls -R vendor/
```

### Código para `main.go`

```go
package main

import "github.com/fatih/color"

func main() {
    color.Green("Usando apenas a biblioteca color!")
    // Note que não estamos usando gin ou gorilla/mux
}
```

### Desafio Extra
1. Compare o tamanho do projeto antes e depois do `vendor`
2. Tente compilar com `-mod=vendor`
3. Remova o diretório `vendor/` e veja se o projeto ainda compila

### Reflexão
- Por que `go mod tidy` removeu a dependência não usada?
- Quando você usaria `go mod vendor` em um projeto real?
- Qual o tamanho do diretório `vendor/` e por que ele pode ser grande?

---

## Exercício 6: Entendendo Imports Circulares

### Objetivo
Entender o problema de imports circulares e como evitá-los.

### Tarefa
Tente criar um import circular e veja o erro. Depois, refatore para resolver.

### Código que Causa Erro

```go
// arquivo: a/a.go
package a

import "b"

func FuncaoA() {
    b.FuncaoB()
}

// arquivo: b/b.go
package b

import "a"  // ERRO! Import circular

func FuncaoB() {
    a.FuncaoA()
}
```

### Solução: Criar Package Intermediário

Crie um package `c` que contém a funcionalidade compartilhada:

```go
// arquivo: c/c.go
package c

func FuncaoCompartilhada() {
    // Lógica compartilhada
}

// arquivo: a/a.go
package a

import (
    "b"
    "c"
)

func FuncaoA() {
    c.FuncaoCompartilhada()
    b.FuncaoB()
}

// arquivo: b/b.go
package b

import "c"

func FuncaoB() {
    c.FuncaoCompartilhada()
}
```

### Reflexão
- Por que Go não permite imports circulares?
- Qual é a melhor estratégia para evitar esse problema?
- Como você identificaria imports circulares em um projeto grande?

---

## Exercício 7: Criando um Módulo Publicável

### Objetivo
Preparar um módulo para ser publicado e usado por outros.

### Tarefa
Crie um módulo simples de utilitários matemáticos que pode ser publicado:

1. Crie a estrutura do módulo
2. Escreva código bem documentado
3. Crie um README.md
4. Prepare para versionamento (mas não precisa publicar de verdade)

### Estrutura

```
math-utils/
├── go.mod
├── README.md
├── soma.go
├── subtracao.go
└── multiplicacao.go
```

### Código de Exemplo

#### `soma.go`

```go
// Package mathutils fornece funções matemáticas utilitárias.
package mathutils

// Soma retorna a soma de dois números inteiros.
//
// Exemplo:
//
//     resultado := mathutils.Soma(5, 3)
//     fmt.Println(resultado) // Output: 8
func Soma(a, b int) int {
    return a + b
}
```

#### `README.md`

```markdown
# Math Utils

Biblioteca de utilitários matemáticos para Go.

## Instalação

```bash
go get github.com/seu-usuario/math-utils
```

## Uso

```go
import "github.com/seu-usuario/math-utils"

func main() {
    resultado := mathutils.Soma(5, 3)
    fmt.Println(resultado)
}
```

## Funções Disponíveis

- `Soma(a, b int) int` - Soma dois números
- `Subtracao(a, b int) int` - Subtrai dois números
- `Multiplicacao(a, b int) int` - Multiplica dois números
```

### Desafio Extra
1. Adicione mais funções matemáticas
2. Crie testes para as funções
3. Adicione exemplos de uso
4. Documente todas as funções seguindo o padrão Go

### Reflexão
- O que torna um módulo "publicável"?
- Por que a documentação é importante?
- Como você versionaria este módulo (v1.0.0, v1.1.0, etc.)?

---

## Reflexão Final

Agora que você praticou, reflita sobre:

1. **Organização**: Como a organização em packages ajuda no desenvolvimento?
2. **Dependências**: Qual a importância de gerenciar dependências corretamente?
3. **Reutilização**: Como packages facilitam a reutilização de código?
4. **Manutenção**: Como uma boa estrutura facilita a manutenção?
5. **Colaboração**: Como a organização ajuda quando trabalhamos em equipe?

### Perguntas para Pensar

- Quando você criaria um novo package ao invés de adicionar ao existente?
- Como você decide se uma função deve ser pública ou privada?
- Qual a melhor estrutura para um projeto que você está começando?
- Como você balanceia organização com simplicidade?

---

## Próximos Passos

Agora que você praticou:
- ✅ Experimente criar projetos com diferentes estruturas
- ✅ Explore packages da biblioteca padrão do Go
- ✅ Tente usar diferentes bibliotecas de terceiros
- ✅ Pense em como organizar seus próprios projetos

Na próxima aula, vamos ver boas práticas e otimizações de performance relacionadas à organização de código!

Até lá, continue praticando! 🚀

