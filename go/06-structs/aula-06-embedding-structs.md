# Módulo 6: Structs em Go

## Aula 6: Embedding Structs - Composição em Go

Olá! Agora vamos aprender sobre **embedding** (incorporação) de structs - uma das características mais poderosas e únicas de Go. Embedding permite que você "incorpore" uma struct dentro de outra, tornando seus campos e métodos diretamente acessíveis.

---

## 1. O Que É Embedding?

**Embedding** é a incorporação de uma struct dentro de outra **sem dar um nome ao campo**. Isso torna os campos e métodos da struct incorporada diretamente acessíveis na struct que a incorpora.

### Analogia Rápida

Pense em embedding como **herdar funcionalidades** sem usar herança tradicional. É como ter uma "caixa de ferramentas" dentro de outra caixa, mas você pode acessar as ferramentas diretamente sem precisar abrir a caixa interna.

---

## 2. Embedding vs. Composição Normal

### Composição Normal (Com Nome)

```go
type Endereco struct {
    Rua    string
    Cidade string
}

type Pessoa struct {
    Nome     string
    Endereco Endereco  // Campo nomeado
}

func main() {
    pessoa := Pessoa{
        Nome: "João",
        Endereco: Endereco{
            Rua:    "Rua A",
            Cidade: "São Paulo",
        },
    }

    // Precisa acessar através do nome do campo
    fmt.Println(pessoa.Endereco.Rua)  // "Rua A"
}
```

### Embedding (Sem Nome)

```go
type Endereco struct {
    Rua    string
    Cidade string
}

type Pessoa struct {
    Nome string
    Endereco  // SEM nome do campo - isso é embedding!
}

func main() {
    pessoa := Pessoa{
        Nome: "João",
        Endereco: Endereco{  // Ainda inicializa normalmente
            Rua:    "Rua A",
            Cidade: "São Paulo",
        },
    }

    // Pode acessar diretamente!
    fmt.Println(pessoa.Rua)     // "Rua A" - acesso direto!
    fmt.Println(pessoa.Cidade)  // "São Paulo"

    // Também pode acessar pelo tipo (se necessário)
    fmt.Println(pessoa.Endereco.Rua)  // "Rua A" - ainda funciona
}
```

**Diferença chave:** Com embedding, você pode acessar campos diretamente sem precisar do nome do campo intermediário.

---

## 3. Embedding de Métodos

### Exemplo: Embedding com Métodos

```go
type Animal struct {
    Nome string
}

// Método da struct Animal
func (a Animal) FazerSom() {
    fmt.Printf("%s faz um som\n", a.Nome)
}

type Cachorro struct {
    Animal  // Embedding - incorpora Animal
    Raca    string
}

func main() {
    cachorro := Cachorro{
        Animal: Animal{Nome: "Rex"},
        Raca:   "Labrador",
    }

    // Pode chamar método diretamente!
    cachorro.FazerSom()  // "Rex faz um som"

    // Também funciona pelo tipo
    cachorro.Animal.FazerSom()  // "Rex faz um som"
}
```

---

## 4. Embedding Múltiplo

Você pode incorporar múltiplas structs:

```go
type Endereco struct {
    Rua    string
    Cidade string
}

type Contato struct {
    Email string
    Telefone string
}

type Pessoa struct {
    Nome string
    Endereco  // Embedding 1
    Contato   // Embedding 2
}

func main() {
    pessoa := Pessoa{
        Nome: "João",
        Endereco: Endereco{
            Rua:    "Rua A",
            Cidade: "São Paulo",
        },
        Contato: Contato{
            Email:    "joao@email.com",
            Telefone: "123-456",
        },
    }

    // Acesso direto a todos os campos!
    fmt.Println(pessoa.Nome)      // "João"
    fmt.Println(pessoa.Rua)       // "Rua A" (de Endereco)
    fmt.Println(pessoa.Email)      // "joao@email.com" (de Contato)
    fmt.Println(pessoa.Telefone)   // "123-456" (de Contato)
}
```

---

## 5. Resolução de Conflitos de Nomes

### O Que Acontece com Campos Duplicados?

Se duas structs incorporadas têm campos com o mesmo nome, você **deve** usar o nome do tipo para acessar:

```go
type A struct {
    Nome string
}

type B struct {
    Nome string
}

type C struct {
    A  // Tem campo Nome
    B  // Também tem campo Nome
}

func main() {
    c := C{
        A: A{Nome: "Nome de A"},
        B: B{Nome: "Nome de B"},
    }

    // ERRO! Ambíguo - qual Nome?
    // fmt.Println(c.Nome)  // Erro de compilação!

    // Deve especificar qual tipo
    fmt.Println(c.A.Nome)  // "Nome de A"
    fmt.Println(c.B.Nome)    // "Nome de B"
}
```

**Regra:** Se houver conflito de nomes, você **deve** usar o nome do tipo para desambiguar.

---

## 6. Embedding de Ponteiros

Você também pode incorporar ponteiros para structs:

```go
type Endereco struct {
    Rua    string
    Cidade string
}

type Pessoa struct {
    Nome string
    *Endereco  // Embedding de ponteiro
}

func main() {
    pessoa := Pessoa{
        Nome: "João",
        Endereco: &Endereco{  // Ponteiro
            Rua:    "Rua A",
            Cidade: "São Paulo",
        },
    }

    // Acesso direto ainda funciona!
    fmt.Println(pessoa.Rua)  // "Rua A"
}
```

---

## 7. Exemplo Prático: Sistema de Logging

```go
package main

import (
    "fmt"
    "time"
)

// Logger fornece funcionalidade de logging
type Logger struct {
    Prefixo string
}

func (l Logger) Log(mensagem string) {
    timestamp := time.Now().Format("2006-01-02 15:04:05")
    fmt.Printf("[%s] %s: %s\n", timestamp, l.Prefixo, mensagem)
}

func (l Logger) LogErro(erro error) {
    l.Log(fmt.Sprintf("ERRO: %v", erro))
}

// Servico incorpora Logger
type Servico struct {
    Logger  // Embedding - incorpora funcionalidade de logging
    Nome    string
}

func (s Servico) Executar() {
    s.Log("Iniciando execução")  // Método de Logger acessível diretamente
    // ... lógica do serviço
    s.Log("Execução concluída")
}

func main() {
    servico := Servico{
        Logger: Logger{Prefixo: "SERVIÇO"},
        Nome:   "MeuServico",
    }

    servico.Executar()
    // [2024-01-01 10:00:00] SERVIÇO: Iniciando execução
    // [2024-01-01 10:00:01] SERVIÇO: Execução concluída
}
```

---

## 8. Embedding e Interfaces

Embedding é especialmente útil com interfaces:

```go
// Interface
type Leitor interface {
    Ler() string
}

type Escritor interface {
    Escrever(string)
}

// Interface composta via embedding
type LeitorEscritor interface {
    Leitor  // Incorpora Leitor
    Escritor  // Incorpora Escritor
}

// Implementação
type Arquivo struct {
    conteudo string
}

func (a Arquivo) Ler() string {
    return a.conteudo
}

func (a *Arquivo) Escrever(texto string) {
    a.conteudo = texto
}

// Arquivo implementa LeitorEscritor automaticamente!
```

---

## 9. Exemplo Completo: Sistema de Usuários

```go
package main

import (
    "fmt"
    "time"
)

// Timestamps fornece campos de data/hora
type Timestamps struct {
    CriadoEm  time.Time
    AtualizadoEm time.Time
}

func (t *Timestamps) Atualizar() {
    t.AtualizadoEm = time.Now()
}

// Auditoria fornece campos de auditoria
type Auditoria struct {
    CriadoPor    string
    AtualizadoPor string
}

// Usuario incorpora Timestamps e Auditoria
type Usuario struct {
    ID       int
    Nome     string
    Email    string
    Timestamps  // Embedding
    Auditoria   // Embedding
}

func NovoUsuario(nome, email, criadoPor string) *Usuario {
    agora := time.Now()
    return &Usuario{
        ID:    1,
        Nome:  nome,
        Email: email,
        Timestamps: Timestamps{
            CriadoEm:     agora,
            AtualizadoEm: agora,
        },
        Auditoria: Auditoria{
            CriadoPor: criadoPor,
        },
    }
}

func (u *Usuario) Atualizar(nome, atualizadoPor string) {
    u.Nome = nome
    u.AtualizadoPor = atualizadoPor
    u.Atualizar()  // Método de Timestamps acessível diretamente
}

func main() {
    usuario := NovoUsuario("João", "joao@email.com", "admin")

    // Acesso direto a campos incorporados
    fmt.Println(usuario.CriadoEm)      // Timestamp de criação
    fmt.Println(usuario.CriadoPor)    // "admin"

    usuario.Atualizar("João Silva", "admin")
    fmt.Println(usuario.AtualizadoEm)  // Timestamp atualizado
    fmt.Println(usuario.AtualizadoPor)  // "admin"
}
```

---

## 10. Embedding vs. Herança

### Diferenças Fundamentais

**Herança (OOP tradicional):**

- "É um" (is-a)
- Hierarquia de classes
- Polimorfismo via herança
- Pode ser complexo

**Embedding (Go):**

- "Tem um" (has-a) com acesso direto
- Composição
- Polimorfismo via interfaces
- Mais simples e flexível

### Exemplo Comparativo

**Herança (conceitual, Go não tem):**

```java
// Java (exemplo conceitual)
class Animal { }
class Cachorro extends Animal { }  // Cachorro É UM Animal
```

**Embedding (Go):**

```go
type Animal struct { }
type Cachorro struct {
    Animal  // Cachorro TEM UM Animal (com acesso direto)
}
```

**Filosofia Go:** "Composição sobre herança"

---

## 11. Boas Práticas com Embedding

### ✅ BOM: Embedding para Funcionalidades Reutilizáveis

```go
// Funcionalidade reutilizável
type Timestamps struct {
    CriadoEm time.Time
    AtualizadoEm time.Time
}

// Múltiplas structs podem incorporar
type Usuario struct {
    Timestamps
    Nome string
}

type Produto struct {
    Timestamps
    Nome string
}
```

### ✅ BOM: Embedding para Interfaces

```go
type Leitor interface {
    Ler() string
}

type Escritor interface {
    Escrever(string)
}

type LeitorEscritor interface {
    Leitor
    Escritor
}
```

### ❌ EVITE: Embedding Excessivo

```go
// Muitos embeddings podem tornar o código confuso
type Pessoa struct {
    A
    B
    C
    D
    E  // Muitos embeddings!
}
```

**Problema:** Pode causar conflitos de nomes e tornar o código difícil de entender.

---

## 12. Resolução de Métodos

### Ordem de Resolução

Quando você chama um método, Go procura na seguinte ordem:

1. Métodos definidos diretamente na struct
2. Métodos de structs incorporadas (na ordem de declaração)
3. Métodos de structs incorporadas recursivamente

```go
type A struct{}

func (A) Metodo() { fmt.Println("A") }

type B struct{}

func (B) Metodo() { fmt.Println("B") }

type C struct {
    A
    B
}

func main() {
    c := C{}
    c.Metodo()  // ERRO! Ambíguo - A.Metodo() ou B.Metodo()?

    // Deve especificar
    c.A.Metodo()  // "A"
    c.B.Metodo()  // "B"
}
```

### Sobrescrever Métodos

Você pode "sobrescrever" métodos definindo na struct que incorpora:

```go
type Animal struct{}

func (Animal) FazerSom() {
    fmt.Println("Som genérico")
}

type Cachorro struct {
    Animal
}

// Sobrescreve o método de Animal
func (Cachorro) FazerSom() {
    fmt.Println("Au au!")
}

func main() {
    cachorro := Cachorro{}
    cachorro.FazerSom()  // "Au au!" - método de Cachorro
    cachorro.Animal.FazerSom()  // "Som genérico" - método de Animal
}
```

---

## 13. Exemplo Avançado: Middleware Pattern

```go
package main

import "fmt"

// Autenticacao fornece funcionalidade de autenticação
type Autenticacao struct {
    UsuarioID int
    Token     string
}

func (a Autenticacao) EstaAutenticado() bool {
    return a.UsuarioID > 0 && a.Token != ""
}

// Autorizacao fornece funcionalidade de autorização
type Autorizacao struct {
    Permissoes []string
}

func (a Autorizacao) TemPermissao(permissao string) bool {
    for _, p := range a.Permissoes {
        if p == permissao {
            return true
        }
    }
    return false
}

// Requisicao incorpora Autenticacao e Autorizacao
type Requisicao struct {
    Autenticacao  // Embedding
    Autorizacao   // Embedding
    Path          string
    Metodo        string
}

func (r Requisicao) PodeAcessar() bool {
    if !r.EstaAutenticado() {  // Método de Autenticacao
        return false
    }
    return r.TemPermissao(r.Path)  // Método de Autorizacao
}

func main() {
    req := Requisicao{
        Autenticacao: Autenticacao{
            UsuarioID: 1,
            Token:     "abc123",
        },
        Autorizacao: Autorizacao{
            Permissoes: []string{"/admin", "/users"},
        },
        Path:   "/admin",
        Metodo: "GET",
    }

    if req.PodeAcessar() {
        fmt.Println("Acesso permitido")
    }
}
```

---

## 14. Resumo: Embedding vs. Composição Normal

| Aspecto       | Composição Normal        | Embedding                             |
| ------------- | ------------------------ | ------------------------------------- |
| **Acesso**    | `obj.Campo.SubCampo`     | `obj.SubCampo` (direto)               |
| **Métodos**   | `obj.Campo.Metodo()`     | `obj.Metodo()` (direto)               |
| **Uso**       | Quando quer encapsular   | Quando quer reutilizar funcionalidade |
| **Conflitos** | Não há (campos nomeados) | Pode haver (deve especificar tipo)    |

---

## 15. Quando Usar Embedding?

Use embedding quando:

✅ **Reutilizar funcionalidade** - Timestamps, Logging, etc.
✅ **Compor interfaces** - Interfaces que incorporam outras
✅ **Acesso direto desejado** - Quer acessar campos/métodos diretamente
✅ **Evitar herança** - Prefere composição sobre herança

Evite embedding quando:

❌ **Muitos conflitos** - Muitos campos/métodos com mesmo nome
❌ **Encapsulamento necessário** - Quer esconder campos internos
❌ **Composição explícita** - Quer deixar claro que é uma composição

---

## Conclusão

Embedding é uma característica poderosa de Go que permite:

1. **Composição flexível** - Incorporar funcionalidades sem herança
2. **Acesso direto** - Campos e métodos acessíveis diretamente
3. **Reutilização** - Funcionalidades comuns podem ser incorporadas
4. **Simplicidade** - Mais simples que herança tradicional
5. **Filosofia Go** - "Composição sobre herança"

Lembre-se:

- Embedding não é herança - é composição com acesso direto
- Use para reutilizar funcionalidades comuns
- Cuidado com conflitos de nomes
- Prefira composição sobre herança

Parabéns! Você completou o módulo sobre Structs em Go. Agora você domina:

- Structs básicas
- Métodos
- Struct tags e JSON
- Embedding

Continue praticando e explorando!
