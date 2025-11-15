# Módulo 12: Interfaces em Go

## Aula 1: Interfaces - Contratos e Polimorfismo

Olá! Bem-vindo ao décimo segundo módulo. Até agora você aprendeu sobre structs, métodos e como adicionar comportamento aos seus tipos. Mas e se você quiser que diferentes tipos compartilhem o mesmo comportamento? Como você pode escrever código que funcione com vários tipos diferentes sem saber exatamente qual tipo está sendo usado?

Nesta aula, vamos mergulhar profundamente em **interfaces** (interfaces) em Go. Você vai aprender como definir contratos que especificam comportamentos, como tipos satisfazem interfaces automaticamente, e como usar interfaces para criar código flexível e polimórfico.

---

## 1. O Que São Interfaces?

**Interfaces** em Go são contratos que especificam um conjunto de métodos que um tipo deve implementar. Diferente de outras linguagens, em Go os tipos **satisfazem interfaces implicitamente** - você não precisa declarar explicitamente que um tipo implementa uma interface.

### Conceito Fundamental

Uma interface define **o que** um tipo pode fazer (os métodos), mas não **como** ele faz. Isso permite que você escreva código que funciona com qualquer tipo que tenha esses métodos, sem se preocupar com os detalhes de implementação.

### Sintaxe Básica

```go
// Interface define um contrato
type Forma interface {
    Area() float64
    Perimetro() float64
}

// Qualquer tipo que tenha esses métodos satisfaz a interface automaticamente
```

**Observação importante**: Em Go, interfaces são **implícitas**. Se um tipo tem os métodos necessários, ele automaticamente satisfaz a interface. Não há palavra-chave `implements` como em Java ou C#.

---

## 2. Interfaces Básicas

Vamos começar com um exemplo prático para entender como interfaces funcionam:

```go
package main

import (
    "fmt"
    "math"
)

// Interface que define o contrato para formas geométricas
type Forma interface {
    Area() float64
}

// Retangulo implementa Forma (implicitamente)
type Retangulo struct {
    Largura float64
    Altura  float64
}

func (r Retangulo) Area() float64 {
    return r.Largura * r.Altura
}

// Circulo implementa Forma (implicitamente)
type Circulo struct {
    Raio float64
}

func (c Circulo) Area() float64 {
    return math.Pi * c.Raio * c.Raio
}

// Função que aceita qualquer Forma
func ImprimirArea(f Forma) {
    fmt.Printf("Área: %.2f\n", f.Area())
}

func main() {
    ret := Retangulo{Largura: 10, Altura: 5}
    circ := Circulo{Raio: 3}
    
    // Ambos podem ser usados como Forma
    ImprimirArea(ret)  // Área: 50.00
    ImprimirArea(circ) // Área: 28.27
}
```

### Por Que Interfaces São Poderosas?

1. **Polimorfismo**: Você pode tratar diferentes tipos da mesma forma
2. **Desacoplamento**: Seu código depende de comportamentos, não de tipos concretos
3. **Testabilidade**: Fácil criar mocks e testar código
4. **Flexibilidade**: Adicione novos tipos sem modificar código existente

### Exemplo Prático: Sistema de Pagamento

```go
package main

import "fmt"

// Interface para métodos de pagamento
type Pagamento interface {
    Processar(valor float64) error
    Validar() bool
}

// CartaoCredito implementa Pagamento
type CartaoCredito struct {
    Numero string
    Valido bool
}

func (c CartaoCredito) Processar(valor float64) error {
    if !c.Validar() {
        return fmt.Errorf("cartão inválido")
    }
    fmt.Printf("Processando R$ %.2f no cartão %s\n", valor, c.Numero)
    return nil
}

func (c CartaoCredito) Validar() bool {
    return c.Valido && len(c.Numero) > 0
}

// Pix implementa Pagamento
type Pix struct {
    Chave string
}

func (p Pix) Processar(valor float64) error {
    fmt.Printf("Processando R$ %.2f via Pix para %s\n", valor, p.Chave)
    return nil
}

func (p Pix) Validar() bool {
    return len(p.Chave) > 0
}

// Função que aceita qualquer método de pagamento
func RealizarCompra(p Pagamento, valor float64) error {
    if !p.Validar() {
        return fmt.Errorf("método de pagamento inválido")
    }
    return p.Processar(valor)
}

func main() {
    cartao := CartaoCredito{Numero: "1234-5678", Valido: true}
    pix := Pix{Chave: "usuario@email.com"}
    
    // Mesma função funciona com diferentes tipos
    RealizarCompra(cartao, 100.50)  // Processando R$ 100.50 no cartão 1234-5678
    RealizarCompra(pix, 250.00)     // Processando R$ 250.00 via Pix para usuario@email.com
}
```

---

## 3. Empty Interface (interface{})

A **empty interface** (`interface{}`) é uma interface especial que não especifica nenhum método. Como todo tipo implementa pelo menos zero métodos, **qualquer tipo satisfaz a empty interface**.

### Quando Usar?

A empty interface é útil quando você precisa aceitar valores de **qualquer tipo**, especialmente antes do Go 1.18 (que introduziu generics). Hoje em dia, generics são preferíveis na maioria dos casos, mas ainda é importante entender `interface{}`.

### Sintaxe

```go
// Empty interface aceita qualquer tipo
var qualquerCoisa interface{}

qualquerCoisa = 42
qualquerCoisa = "hello"
qualquerCoisa = true
qualquerCoisa = []int{1, 2, 3}
```

### Exemplo Prático: Função que Aceita Qualquer Tipo

```go
package main

import "fmt"

// Função que aceita qualquer tipo
func Imprimir(valor interface{}) {
    fmt.Printf("Tipo: %T, Valor: %v\n", valor, valor)
}

func main() {
    Imprimir(42)                    // Tipo: int, Valor: 42
    Imprimir("Olá")                 // Tipo: string, Valor: Olá
    Imprimir(3.14)                  // Tipo: float64, Valor: 3.14
    Imprimir([]int{1, 2, 3})        // Tipo: []int, Valor: [1 2 3]
    Imprimir(map[string]int{"a": 1}) // Tipo: map[string]int, Valor: map[a:1]
}
```

### Limitação: Você Precisa de Type Assertion

Quando você usa `interface{}`, você perde informações sobre o tipo. Para acessar o valor original, você precisa usar **type assertion** (que veremos em detalhes mais adiante):

```go
func Processar(valor interface{}) {
    // Isso não funciona diretamente:
    // valor + 10  // Erro: interface{} não pode ser usado em operações
    
    // Você precisa verificar o tipo primeiro
    if num, ok := valor.(int); ok {
        fmt.Println(num + 10)
    }
}
```

**Nota**: A partir do Go 1.18, prefira usar **generics** em vez de `interface{}` quando possível, pois oferecem type safety em tempo de compilação.

---

## 4. Embedding Interfaces

Assim como você pode fazer embedding de structs, você pode fazer **embedding de interfaces** para criar interfaces maiores a partir de interfaces menores. Isso promove composição e reutilização.

### Sintaxe

```go
// Interface base
type Leitor interface {
    Ler() []byte
}

// Interface base
type Escritor interface {
    Escrever([]byte) error
}

// Interface composta (embedding)
type LeitorEscritor interface {
    Leitor  // Embedding: inclui todos os métodos de Leitor
    Escritor // Embedding: inclui todos os métodos de Escritor
}
```

### Exemplo Prático: Sistema de Arquivos

```go
package main

import "fmt"

// Interfaces base
type Leitor interface {
    Ler() string
}

type Escritor interface {
    Escrever(conteudo string) error
}

type Fechador interface {
    Fechar() error
}

// Interface composta usando embedding
type Arquivo interface {
    Leitor
    Escritor
    Fechador
}

// Implementação concreta
type MeuArquivo struct {
    conteudo string
    fechado  bool
}

func (f *MeuArquivo) Ler() string {
    if f.fechado {
        return ""
    }
    return f.conteudo
}

func (f *MeuArquivo) Escrever(conteudo string) error {
    if f.fechado {
        return fmt.Errorf("arquivo fechado")
    }
    f.conteudo = conteudo
    return nil
}

func (f *MeuArquivo) Fechar() error {
    f.fechado = true
    return nil
}

// Função que aceita qualquer Arquivo
func ProcessarArquivo(a Arquivo) {
    fmt.Println("Lendo:", a.Ler())
    a.Escrever("Novo conteúdo")
    fmt.Println("Lendo após escrita:", a.Ler())
    a.Fechar()
}

func main() {
    arquivo := &MeuArquivo{}
    ProcessarArquivo(arquivo)
}
```

### Vantagens do Embedding de Interfaces

1. **Composição**: Crie interfaces complexas a partir de interfaces simples
2. **Reutilização**: Evite duplicação de código
3. **Flexibilidade**: Tipos podem implementar apenas as interfaces que precisam
4. **Hierarquia**: Crie hierarquias de interfaces de forma natural

### Exemplo: Interface com Métodos Adicionais

```go
// Você pode adicionar métodos além dos embedded
type LeitorEscritorAvancado interface {
    Leitor
    Escritor
    // Método adicional
    Limpar() error
}
```

---

## 5. Type Assertions

**Type assertions** permitem extrair o valor concreto de uma interface. Você "afirma" que o valor dentro da interface é de um tipo específico.

### Sintaxe Básica

```go
// Forma 1: Pode causar panic se o tipo estiver errado
valor := interfaceVar.(Tipo)

// Forma 2: Segura, retorna ok para verificar
valor, ok := interfaceVar.(Tipo)
```

### Exemplo Prático

```go
package main

import "fmt"

func Processar(valor interface{}) {
    // Type assertion segura
    if str, ok := valor.(string); ok {
        fmt.Printf("É uma string: %s\n", str)
    } else if num, ok := valor.(int); ok {
        fmt.Printf("É um int: %d\n", num)
    } else {
        fmt.Printf("Tipo desconhecido: %T\n", valor)
    }
}

func main() {
    Processar("hello")  // É uma string: hello
    Processar(42)       // É um int: 42
    Processar(3.14)     // Tipo desconhecido: float64
}
```

### Type Assertion Pode Causar Panic

```go
func ExemploPerigoso() {
    var i interface{} = "hello"
    
    // Isso funciona
    s := i.(string)
    fmt.Println(s) // hello
    
    // Isso causa PANIC!
    f := i.(float64) // panic: interface conversion: interface {} is string, not float64
    fmt.Println(f)
}

// Use sempre a forma segura:
func ExemploSeguro() {
    var i interface{} = "hello"
    
    if f, ok := i.(float64); ok {
        fmt.Println(f)
    } else {
        fmt.Println("Não é float64")
    }
}
```

### Exemplo Prático: Processamento de Dados

```go
package main

import "fmt"

type Processador interface {
    Processar()
}

type DadosTexto struct {
    Texto string
}

func (d DadosTexto) Processar() {
    fmt.Printf("Processando texto: %s\n", d.Texto)
}

type DadosNumericos struct {
    Numeros []int
}

func (d DadosNumericos) Processar() {
    soma := 0
    for _, n := range d.Numeros {
        soma += n
    }
    fmt.Printf("Soma dos números: %d\n", soma)
}

func ProcessarDados(p Processador) {
    p.Processar()
    
    // Type assertion para acessar campos específicos
    if texto, ok := p.(DadosTexto); ok {
        fmt.Printf("Tamanho do texto: %d caracteres\n", len(texto.Texto))
    }
    
    if numeros, ok := p.(DadosNumericos); ok {
        fmt.Printf("Quantidade de números: %d\n", len(numeros.Numeros))
    }
}

func main() {
    dados1 := DadosTexto{Texto: "Olá, mundo!"}
    dados2 := DadosNumericos{Numeros: []int{1, 2, 3, 4, 5}}
    
    ProcessarDados(dados1)
    ProcessarDados(dados2)
}
```

---

## 6. Type Switch

O **type switch** é uma forma especial de switch que permite verificar o tipo de uma interface de forma mais elegante que múltiplas type assertions.

### Sintaxe

```go
switch v := valor.(type) {
case Tipo1:
    // v é do tipo Tipo1
case Tipo2:
    // v é do tipo Tipo2
default:
    // tipo desconhecido
}
```

### Exemplo Básico

```go
package main

import "fmt"

func Descrever(valor interface{}) {
    switch v := valor.(type) {
    case int:
        fmt.Printf("É um inteiro: %d\n", v)
    case string:
        fmt.Printf("É uma string: %s\n", v)
    case bool:
        fmt.Printf("É um booleano: %v\n", v)
    case float64:
        fmt.Printf("É um float64: %.2f\n", v)
    default:
        fmt.Printf("Tipo desconhecido: %T\n", v)
    }
}

func main() {
    Descrever(42)        // É um inteiro: 42
    Descrever("hello")   // É uma string: hello
    Descrever(true)     // É um booleano: true
    Descrever(3.14)      // É um float64: 3.14
    Descrever([]int{1})  // Tipo desconhecido: []int
}
```

### Exemplo Prático: Calculadora Universal

```go
package main

import "fmt"

func Somar(a, b interface{}) interface{} {
    switch v1 := a.(type) {
    case int:
        switch v2 := b.(type) {
        case int:
            return v1 + v2
        case float64:
            return float64(v1) + v2
        }
    case float64:
        switch v2 := b.(type) {
        case int:
            return v1 + float64(v2)
        case float64:
            return v1 + v2
        }
    case string:
        if v2, ok := b.(string); ok {
            return v1 + v2
        }
    }
    return nil
}

func main() {
    fmt.Println(Somar(5, 3))           // 8
    fmt.Println(Somar(5.5, 2.5))       // 8
    fmt.Println(Somar(5, 2.5))         // 7.5
    fmt.Println(Somar("Olá, ", "mundo!")) // Olá, mundo!
}
```

### Type Switch com Interfaces

Você também pode usar type switch com interfaces customizadas:

```go
package main

import "fmt"

type Animal interface {
    FazerSom() string
}

type Cachorro struct {
    Nome string
}

func (c Cachorro) FazerSom() string {
    return "Au au!"
}

type Gato struct {
    Nome string
}

func (g Gato) FazerSom() string {
    return "Miau!"
}

func Interagir(animal Animal) {
    fmt.Println(animal.FazerSom())
    
    // Type switch para comportamento específico
    switch a := animal.(type) {
    case Cachorro:
        fmt.Printf("%s está abanando o rabo!\n", a.Nome)
    case Gato:
        fmt.Printf("%s está ronronando!\n", a.Nome)
    }
}

func main() {
    cachorro := Cachorro{Nome: "Rex"}
    gato := Gato{Nome: "Mimi"}
    
    Interagir(cachorro) // Au au! / Rex está abanando o rabo!
    Interagir(gato)     // Miau! / Mimi está ronronando!
}
```

### Múltiplos Tipos em um Case

Você pode agrupar múltiplos tipos em um único case:

```go
func Processar(valor interface{}) {
    switch v := valor.(type) {
    case int, int32, int64:
        fmt.Printf("É um inteiro: %v\n", v)
    case float32, float64:
        fmt.Printf("É um float: %v\n", v)
    case string:
        fmt.Printf("É uma string: %s\n", v)
    default:
        fmt.Printf("Tipo: %T\n", v)
    }
}
```

---

## 7. Interfaces Comuns na Biblioteca Padrão

Go tem várias interfaces importantes na biblioteca padrão. Vamos ver algumas:

### io.Reader e io.Writer

```go
package main

import (
    "fmt"
    "io"
    "strings"
)

// io.Reader: lê dados
func LerDados(r io.Reader) {
    dados := make([]byte, 100)
    n, _ := r.Read(dados)
    fmt.Printf("Lidos %d bytes: %s\n", n, dados[:n])
}

// io.Writer: escreve dados
func EscreverDados(w io.Writer, dados string) {
    w.Write([]byte(dados))
}

func main() {
    // strings.Reader implementa io.Reader
    reader := strings.NewReader("Olá, mundo!")
    LerDados(reader)
    
    // strings.Builder implementa io.Writer
    var builder strings.Builder
    EscreverDados(&builder, "Hello, World!")
    fmt.Println(builder.String())
}
```

### fmt.Stringer

```go
package main

import "fmt"

type Pessoa struct {
    Nome  string
    Idade int
}

// Implementa fmt.Stringer
func (p Pessoa) String() string {
    return fmt.Sprintf("%s (%d anos)", p.Nome, p.Idade)
}

func main() {
    pessoa := Pessoa{Nome: "João", Idade: 30}
    fmt.Println(pessoa) // João (30 anos)
    // fmt.Println chama automaticamente o método String()
}
```

---

## 8. Nil Interfaces

Uma interface em Go pode ser `nil`, mas há uma distinção importante:

```go
package main

import "fmt"

type MinhaInterface interface {
    Metodo()
}

type MeuTipo struct{}

func (m MeuTipo) Metodo() {
    fmt.Println("Método chamado")
}

func main() {
    var i MinhaInterface
    
    // i é nil (não tem valor nem tipo)
    fmt.Println(i == nil) // true
    
    var t *MeuTipo = nil
    i = t
    
    // i NÃO é nil! (tem tipo, mas valor é nil)
    fmt.Println(i == nil) // false
    fmt.Println(t == nil) // true
    
    // Chamar método em interface nil causa panic
    // i.Metodo() // panic: runtime error: invalid memory address
}
```

**Regra importante**: Uma interface só é `nil` quando tanto o valor quanto o tipo são `nil`.

---

## 9. Resumo e Próximos Passos

Nesta aula, você aprendeu:

✅ **Interfaces**: Contratos que especificam métodos que tipos devem implementar  
✅ **Satisfação Implícita**: Tipos satisfazem interfaces automaticamente ao implementar os métodos  
✅ **Empty Interface**: `interface{}` aceita qualquer tipo  
✅ **Embedding**: Combine interfaces para criar interfaces maiores  
✅ **Type Assertions**: Extraia valores concretos de interfaces  
✅ **Type Switch**: Forma elegante de verificar tipos  
✅ **Polimorfismo**: Trate diferentes tipos da mesma forma  

**Pontos importantes para lembrar:**

- Interfaces definem **o que** um tipo pode fazer, não **como**
- Tipos satisfazem interfaces **implicitamente** em Go
- Use type assertions com a forma segura (`valor, ok := i.(Tipo)`)
- Type switch é mais elegante que múltiplas type assertions
- Interfaces permitem polimorfismo e código flexível
- Prefira interfaces pequenas e focadas (princípio da responsabilidade única)

**Regras práticas:**

- **Interfaces pequenas**: Melhor ter várias interfaces pequenas que uma grande
- **Dependa de interfaces, não de tipos concretos**: Facilita testes e flexibilidade
- **Use type assertions com cuidado**: Sempre use a forma segura quando possível
- **Type switch para múltiplos tipos**: Mais legível que várias ifs

Na próxima aula, vamos simplificar esses conceitos com analogias do dia a dia para fixar o aprendizado!

---

**Dica**: Execute o arquivo `01-exemplos.go` para ver todos esses conceitos em ação. Experimente modificar os exemplos e veja o que acontece!

