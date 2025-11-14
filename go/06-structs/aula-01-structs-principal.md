# Módulo 6: Structs em Go

## Aula 1: Structs - Tipos de Dados Customizados

Olá! Bem-vindo ao sexto módulo. Até agora você trabalhou com tipos básicos (int, string, bool) e tipos compostos (arrays, slices, maps). Mas e quando você precisa agrupar dados relacionados sob um único nome? É aqui que entram as **structs**.

Nesta aula, vamos mergulhar no conceito de structs em Go - uma das ferramentas mais poderosas para organizar e estruturar dados em seus programas.

---

## 1. O Que São Structs?

**Structs** (estruturas) são tipos de dados customizados que agrupam campos relacionados sob um único nome. Pense em uma struct como um "formulário" ou "ficha" que contém várias informações sobre algo.

### Analogia Rápida

Imagine uma **ficha de cadastro** de uma pessoa:

```
Nome: João Silva
Idade: 30
Email: joao@email.com
Altura: 1.75
```

Em Go, isso seria uma struct com campos para cada informação.

---

## 2. Por Que Usar Structs?

### Problema: Dados Relacionados Separados

**Sem structs**, você teria variáveis separadas:

```go
var nome string = "João Silva"
var idade int = 30
var email string = "joao@email.com"
var altura float64 = 1.75
```

**Problemas:**

- Variáveis não estão relacionadas logicamente
- Difícil passar todas juntas para uma função
- Fácil perder a relação entre os dados
- Código desorganizado

### Solução: Agrupar em Structs

**Com structs**, você agrupa dados relacionados:

```go
type Pessoa struct {
    Nome   string
    Idade  int
    Email  string
    Altura float64
}
```

**Vantagens:**

- Dados relacionados ficam juntos
- Fácil passar para funções
- Código mais organizado e legível
- Modela entidades do mundo real

---

## 3. Declarando uma Struct

### Sintaxe Básica

```go
type NomeDaStruct struct {
    Campo1 Tipo1
    Campo2 Tipo2
    Campo3 Tipo3
}
```

### Exemplo Prático: Struct Pessoa

```go
type Pessoa struct {
    Nome   string
    Idade  int
    Email  string
    Altura float64
}
```

**Explicação:**

- `type Pessoa struct` - Define um novo tipo chamado `Pessoa`
- Dentro das chaves `{}` estão os **campos** da struct
- Cada campo tem um **nome** e um **tipo**

---

## 4. Criando Instâncias de Structs

Existem várias formas de criar uma struct. Vamos ver todas:

### Forma 1: Declaração com Valores Zero

```go
var pessoa Pessoa
// Todos os campos recebem valores zero:
// Nome: "" (string vazia)
// Idade: 0
// Email: "" (string vazia)
// Altura: 0.0
```

### Forma 2: Literal de Struct (Ordem dos Campos)

```go
pessoa := Pessoa{
    "João Silva",  // Nome
    30,            // Idade
    "joao@email.com", // Email
    1.75,          // Altura
}
```

**Importante:** Os valores devem estar na mesma ordem dos campos na declaração.

### Forma 3: Literal de Struct (Campos Nomeados) - RECOMENDADO

```go
pessoa := Pessoa{
    Nome:   "João Silva",
    Idade:  30,
    Email:  "joao@email.com",
    Altura: 1.75,
}
```

**Vantagens:**

- Mais legível
- Não precisa estar na ordem
- Mais seguro (erros de compilação se campo não existir)

### Forma 4: Inicialização Parcial

```go
pessoa := Pessoa{
    Nome:  "João Silva",
    Idade: 30,
    // Email e Altura receberão valores zero
}
```

### Forma 5: Usando Ponteiro com `new()`

```go
pessoa := new(Pessoa)
// Cria um ponteiro para Pessoa com valores zero
// pessoa é do tipo *Pessoa
```

### Forma 6: Ponteiro com Literal

```go
pessoa := &Pessoa{
    Nome:   "João Silva",
    Idade:  30,
    Email:  "joao@email.com",
    Altura: 1.75,
}
// pessoa é do tipo *Pessoa (ponteiro)
```

---

## 5. Acessando Campos de uma Struct

Você acessa campos usando a **notação de ponto** (`.`):

```go
pessoa := Pessoa{
    Nome:   "João Silva",
    Idade:  30,
    Email:  "joao@email.com",
    Altura: 1.75,
}

// Acessar campos
fmt.Println(pessoa.Nome)   // "João Silva"
fmt.Println(pessoa.Idade)  // 30
fmt.Println(pessoa.Email)  // "joao@email.com"
fmt.Println(pessoa.Altura) // 1.75

// Modificar campos
pessoa.Idade = 31
pessoa.Email = "joao.novo@email.com"
```

### Com Ponteiros

Quando você tem um ponteiro para struct, Go permite acessar campos diretamente (sem `*`):

```go
pessoa := &Pessoa{
    Nome: "João Silva",
    Idade: 30,
}

// Go permite acessar diretamente (desreferencia automaticamente)
fmt.Println(pessoa.Nome)  // Funciona!
pessoa.Idade = 31         // Funciona!

// Equivalente a:
fmt.Println((*pessoa).Nome)  // Forma explícita
(*pessoa).Idade = 31         // Forma explícita
```

---

## 6. Structs Anônimas

Você pode criar structs sem nome (anônimas) para casos simples:

```go
// Struct anônima
pessoa := struct {
    Nome  string
    Idade int
}{
    Nome:  "João",
    Idade: 30,
}
```

**Quando usar:**

- Casos muito simples e locais
- Testes rápidos
- Evite em código de produção (prefira tipos nomeados)

---

## 7. Comparação de Structs

Structs podem ser comparadas se todos os campos forem comparáveis:

```go
p1 := Pessoa{Nome: "João", Idade: 30}
p2 := Pessoa{Nome: "João", Idade: 30}
p3 := Pessoa{Nome: "Maria", Idade: 25}

fmt.Println(p1 == p2) // true (todos os campos são iguais)
fmt.Println(p1 == p3) // false (campos diferentes)
```

**Campos comparáveis:**

- Tipos básicos (int, string, bool, etc.)
- Arrays (se elementos forem comparáveis)
- Outras structs (se campos forem comparáveis)

**Campos NÃO comparáveis:**

- Slices
- Maps
- Funções
- Channels

---

## 8. Structs como Parâmetros de Função

### Passando por Valor (Cópia)

```go
func Aniversario(p Pessoa) {
    p.Idade++  // Modifica apenas a cópia
    // A struct original não é alterada
}

func main() {
    pessoa := Pessoa{Nome: "João", Idade: 30}
    Aniversario(pessoa)
    fmt.Println(pessoa.Idade) // Ainda é 30!
}
```

### Passando por Referência (Ponteiro)

```go
func Aniversario(p *Pessoa) {
    p.Idade++  // Modifica a struct original
}

func main() {
    pessoa := Pessoa{Nome: "João", Idade: 30}
    Aniversario(&pessoa)
    fmt.Println(pessoa.Idade) // Agora é 31!
}
```

**Regra de Ouro:**

- Use ponteiros quando precisar modificar a struct
- Use valores quando apenas ler ou quando a struct for pequena

---

## 9. Métodos em Structs

Em Go, métodos são funções com um **receptor** (receiver). O receptor é a struct que "possui" o método.

### Sintaxe

```go
func (r Receptor) NomeDoMetodo(parametros) tipoRetorno {
    // código
}
```

### Exemplo: Método para Pessoa

```go
type Pessoa struct {
    Nome  string
    Idade int
}

// Método com receptor por valor
func (p Pessoa) Apresentar() string {
    return fmt.Sprintf("Olá, eu sou %s e tenho %d anos", p.Nome, p.Idade)
}

// Método com receptor por ponteiro (pode modificar)
func (p *Pessoa) FazerAniversario() {
    p.Idade++
}

func main() {
    pessoa := Pessoa{Nome: "João", Idade: 30}

    fmt.Println(pessoa.Apresentar()) // "Olá, eu sou João e tenho 30 anos"

    pessoa.FazerAniversario()
    fmt.Println(pessoa.Idade) // 31
}
```

### Receptor por Valor vs. Ponteiro

**Receptor por valor:**

- Recebe uma cópia da struct
- Não pode modificar a struct original
- Útil para métodos que apenas leem

**Receptor por ponteiro:**

- Recebe um ponteiro para a struct
- Pode modificar a struct original
- Útil para métodos que modificam
- Mais eficiente para structs grandes

**Convenção Go:**

- Se o método modifica a struct, use ponteiro
- Se o método apenas lê, pode usar valor (mas ponteiro também funciona)

---

## 10. Exemplo Completo: Sistema de Biblioteca

Vamos criar um exemplo completo usando structs:

```go
package main

import "fmt"

// Livro representa um livro na biblioteca
type Livro struct {
    Titulo     string
    Autor      string
    ISBN       string
    AnoPublicacao int
    Disponivel bool
}

// Usuario representa um usuário da biblioteca
type Usuario struct {
    Nome  string
    Email string
    LivrosEmprestados []string // ISBNs dos livros
}

// Emprestar marca um livro como emprestado
func (l *Livro) Emprestar() {
    l.Disponivel = false
}

// Devolver marca um livro como disponível
func (l *Livro) Devolver() {
    l.Disponivel = true
}

// EmprestarLivro empresta um livro para um usuário
func (u *Usuario) EmprestarLivro(livro *Livro) error {
    if !livro.Disponivel {
        return fmt.Errorf("livro %s não está disponível", livro.Titulo)
    }

    livro.Emprestar()
    u.LivrosEmprestados = append(u.LivrosEmprestados, livro.ISBN)
    return nil
}

// ListarLivros lista os livros emprestados pelo usuário
func (u Usuario) ListarLivros() {
    fmt.Printf("Livros emprestados por %s:\n", u.Nome)
    for _, isbn := range u.LivrosEmprestados {
        fmt.Printf("  - ISBN: %s\n", isbn)
    }
}

func main() {
    // Criar livros
    livro1 := Livro{
        Titulo:        "Aprendendo Go",
        Autor:         "Autor Exemplo",
        ISBN:          "123-456-789",
        AnoPublicacao: 2024,
        Disponivel:    true,
    }

    livro2 := Livro{
        Titulo:        "Go Avançado",
        Autor:         "Outro Autor",
        ISBN:          "987-654-321",
        AnoPublicacao: 2023,
        Disponivel:    true,
    }

    // Criar usuário
    usuario := Usuario{
        Nome:  "João Silva",
        Email: "joao@email.com",
    }

    // Emprestar livros
    usuario.EmprestarLivro(&livro1)
    usuario.EmprestarLivro(&livro2)

    // Listar livros emprestados
    usuario.ListarLivros()

    // Verificar disponibilidade
    fmt.Printf("Livro 1 disponível? %v\n", livro1.Disponivel) // false
    fmt.Printf("Livro 2 disponível? %v\n", livro2.Disponivel) // false
}
```

---

## 11. Structs com Campos de Diferentes Tipos

Structs podem ter campos de qualquer tipo:

```go
type Perfil struct {
    // Tipos básicos
    Nome      string
    Idade     int
    Salario   float64
    Ativo     bool

    // Tipos compostos
    Hobbies   []string           // slice
    Contatos  map[string]string  // map
    Endereco  Endereco           // outra struct

    // Ponteiros
    Gerente   *Pessoa            // ponteiro para Pessoa

    // Funções (menos comum)
    Callback  func() error
}

type Endereco struct {
    Rua    string
    Cidade string
    CEP    string
}
```

---

## 12. Valores Zero de Structs

Quando você declara uma struct sem inicializar, todos os campos recebem seus valores zero:

```go
var pessoa Pessoa
// pessoa.Nome = "" (string vazia)
// pessoa.Idade = 0
// pessoa.Email = "" (string vazia)
// pessoa.Altura = 0.0
```

Para structs aninhadas:

```go
type Endereco struct {
    Rua    string
    Cidade string
}

type Pessoa struct {
    Nome     string
    Endereco Endereco
}

var pessoa Pessoa
// pessoa.Nome = ""
// pessoa.Endereco.Rua = ""
// pessoa.Endereco.Cidade = ""
```

---

## 13. Structs e Slices

Você pode ter slices de structs:

```go
type Pessoa struct {
    Nome  string
    Idade int
}

func main() {
    pessoas := []Pessoa{
        {Nome: "João", Idade: 30},
        {Nome: "Maria", Idade: 25},
        {Nome: "Pedro", Idade: 35},
    }

    // Iterar sobre o slice
    for _, pessoa := range pessoas {
        fmt.Printf("%s tem %d anos\n", pessoa.Nome, pessoa.Idade)
    }
}
```

---

## 14. Structs e Maps

Você pode usar structs como chaves ou valores em maps:

```go
type Coordenada struct {
    X int
    Y int
}

func main() {
    // Map com struct como chave
    pontos := make(map[Coordenada]string)
    pontos[Coordenada{X: 0, Y: 0}] = "Origem"
    pontos[Coordenada{X: 1, Y: 1}] = "Ponto (1,1)"

    // Map com struct como valor
    pessoas := make(map[string]Pessoa)
    pessoas["joao"] = Pessoa{Nome: "João", Idade: 30}
}
```

**Importante:** Para usar struct como chave de map, todos os campos devem ser comparáveis.

---

## 15. Resumo dos Conceitos

| Conceito       | Descrição                                  |
| -------------- | ------------------------------------------ |
| **Declaração** | `type Nome struct { campos }`              |
| **Criação**    | `var x Nome` ou `x := Nome{...}`           |
| **Acesso**     | `x.Campo` (notação de ponto)               |
| **Métodos**    | `func (r Receptor) Metodo() {}`            |
| **Ponteiros**  | Use `*` para modificar ou `&` para criar   |
| **Comparação** | Possível se todos campos forem comparáveis |

---

## 16. Quando Usar Structs?

Use structs quando:

✅ **Agrupar dados relacionados** - Informações sobre uma entidade
✅ **Modelar entidades do mundo real** - Pessoa, Produto, Pedido, etc.
✅ **Passar múltiplos valores** - Em vez de muitos parâmetros
✅ **Organizar código** - Dados relacionados ficam juntos
✅ **Criar tipos customizados** - Comportamento específico com métodos

---

## Conclusão

Structs são fundamentais em Go para:

1. **Organizar dados** - Agrupar informações relacionadas
2. **Modelar domínios** - Representar entidades do mundo real
3. **Criar APIs claras** - Métodos associados aos dados
4. **Estruturar aplicações** - Base para design orientado a objetos em Go

Na próxima parte desta aula, vamos simplificar esses conceitos com analogias do dia a dia para fixar melhor o aprendizado!
