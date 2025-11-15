# Módulo 13: Generics em Go

## Aula 1: Generics - Programação Genérica com Type Safety

Olá! Bem-vindo ao décimo terceiro módulo. Até agora você aprendeu sobre funções, métodos, interfaces e muitos outros conceitos fundamentais de Go. Mas você já se perguntou: **"E se eu quiser criar uma função que funcione com diferentes tipos, mas mantendo a segurança de tipos?"**

Antes do Go 1.18, essa era uma limitação real. Você tinha que escolher entre:
- Criar funções separadas para cada tipo (duplicação de código)
- Usar `interface{}` e perder a segurança de tipos
- Usar code generation (complexo e difícil de manter)

Com a introdução dos **Generics** no Go 1.18, essa limitação foi resolvida! Agora você pode escrever código reutilizável que funciona com múltiplos tipos, mantendo toda a segurança de tipos e performance que Go oferece.

Nesta aula, vamos mergulhar profundamente em Generics: entender o problema que eles resolvem, como usar type parameters, constraints, type inference, e como criar tipos e interfaces genéricas.

---

## 1. Por Que Generics?

### O Problema Antes do Go 1.18

Imagine que você precisa criar uma função que retorna o maior valor entre dois números. Sem generics, você teria que fazer algo assim:

```go
func MaxInt(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func MaxFloat64(a, b float64) float64 {
    if a > b {
        return a
    }
    return b
}

func MaxString(a, b string) string {
    if a > b {
        return a
    }
    return b
}
```

**Problemas dessa abordagem:**
1. **Duplicação de código**: A mesma lógica repetida para cada tipo
2. **Manutenção difícil**: Se precisar corrigir um bug, precisa corrigir em todos os lugares
3. **Escalabilidade ruim**: Para cada novo tipo, precisa criar uma nova função

### Alternativas Antigas (e seus problemas)

#### Opção 1: Usar `interface{}`
```go
func Max(a, b interface{}) interface{} {
    // Como comparar? Não sabemos o tipo!
    // Perdemos type safety
    // Precisamos de type assertions em todo lugar
}
```

**Problemas:**
- Perde type safety (sem verificação em tempo de compilação)
- Precisa de type assertions (pode causar panics)
- Sem autocomplete do IDE
- Performance pior (boxing/unboxing)

#### Opção 2: Code Generation
Usar ferramentas como `genny` ou `go generate` para gerar código para cada tipo.

**Problemas:**
- Complexo de configurar
- Código gerado é difícil de debugar
- Build time aumenta
- Não é idiomático em Go

### A Solução: Generics

Generics permite escrever código que funciona com múltiplos tipos, mantendo:
- ✅ **Type safety**: Verificação em tempo de compilação
- ✅ **Performance**: Sem overhead de runtime
- ✅ **Reutilização**: Um código para múltiplos tipos
- ✅ **Legibilidade**: Código limpo e expressivo

---

## 2. Type Parameters (Parâmetros de Tipo)

### Sintaxe Básica

Generics em Go usam **type parameters** (parâmetros de tipo) definidos entre colchetes `[]` antes dos parâmetros normais:

```go
func NomeFuncao[T TipoConstraint](parametro T) T {
    // corpo da função
}
```

Onde:
- `T` é o **type parameter** (nome da variável de tipo)
- `TipoConstraint` é a **constraint** (restrição do que T pode ser)

### Exemplo Simples

```go
// Função genérica que imprime qualquer tipo
func Print[T any](value T) {
    fmt.Println(value)
}

func main() {
    Print(42)           // T é inferido como int
    Print("hello")      // T é inferido como string
    Print(3.14)         // T é inferido como float64
}
```

### Múltiplos Type Parameters

Você pode ter múltiplos type parameters:

```go
func Swap[T, U any](a T, b U) (U, T) {
    return b, a
}

func main() {
    x, y := Swap(10, "hello")
    // x é string, y é int
}
```

---

## 3. Constraints (Restrições)

Constraints definem **quais tipos** podem ser usados como type parameters. Elas garantem que você só pode usar operações válidas para aquele tipo.

### Constraint `any`

A constraint mais permissiva é `any`, que permite **qualquer tipo**:

```go
func Print[T any](value T) {
    fmt.Println(value)  // Só pode fazer operações que funcionam com qualquer tipo
}
```

**Limitação**: Com `any`, você só pode fazer operações que funcionam com qualquer tipo (como passar como parâmetro, retornar, etc.). Não pode fazer comparações ou operações aritméticas.

### Constraint `comparable`

Para tipos que podem ser comparados com `==` e `!=`:

```go
func Equal[T comparable](a, b T) bool {
    return a == b  // Agora podemos comparar!
}

func main() {
    fmt.Println(Equal(10, 10))        // true
    fmt.Println(Equal("a", "b"))      // false
    fmt.Println(Equal(3.14, 3.14))    // true
}
```

**Tipos comparáveis em Go:**
- Tipos básicos (int, string, float64, etc.)
- Pointers
- Arrays (se os elementos forem comparáveis)
- Structs (se todos os campos forem comparáveis)
- Interfaces (se os valores subjacentes forem comparáveis)

### Constraint `constraints.Ordered`

Para tipos que podem ser comparados com `<`, `<=`, `>`, `>=`:

```go
import "golang.org/x/exp/constraints"

func Max[T constraints.Ordered](a, b T) T {
    if a > b {
        return a
    }
    return b
}

func main() {
    fmt.Println(Max(10, 20))              // 20
    fmt.Println(Max(3.14, 2.71))          // 3.14
    fmt.Println(Max("apple", "banana"))    // "banana"
}
```

**Tipos ordenados:**
- Todos os tipos numéricos (int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64)
- string

### Constraints Customizadas (Union Types)

Você pode criar suas próprias constraints usando **union types** (tipos união):

```go
// Constraint que aceita apenas tipos numéricos
type Numeric interface {
    int | int8 | int16 | int32 | int64 |
    uint | uint8 | uint16 | uint32 | uint64 |
    float32 | float64
}

func Soma[T Numeric](a, b T) T {
    return a + b  // Agora podemos somar!
}

func main() {
    fmt.Println(Soma(10, 20))      // 30
    fmt.Println(Soma(3.14, 2.71))  // 5.85
    // Soma("a", "b")  // ERRO: string não é Numeric
}
```

**Sintaxe de Union Types:**
- `T1 | T2 | T3` significa "T pode ser T1 OU T2 OU T3"
- Cada tipo na união deve ser um tipo concreto (não interface)

### Constraints com Métodos

Você também pode usar interfaces como constraints:

```go
// Interface como constraint
type Stringer interface {
    String() string
}

func PrintString[T Stringer](value T) {
    fmt.Println(value.String())
}

type Pessoa struct {
    Nome string
}

func (p Pessoa) String() string {
    return p.Nome
}

func main() {
    pessoa := Pessoa{Nome: "João"}
    PrintString(pessoa)  // Funciona porque Pessoa implementa Stringer
}
```

---

## 4. Generic Functions (Funções Genéricas)

Funções genéricas são funções que usam type parameters. Vamos ver exemplos práticos:

### Exemplo 1: Função Find

```go
// Encontra um valor em um slice
func Find[T comparable](slice []T, value T) (int, bool) {
    for i, v := range slice {
        if v == value {
            return i, true
        }
    }
    return -1, false
}

func main() {
    numbers := []int{1, 2, 3, 4, 5}
    idx, found := Find(numbers, 3)
    if found {
        fmt.Printf("Encontrado no índice %d\n", idx)
    }
    
    names := []string{"Alice", "Bob", "Charlie"}
    idx2, found2 := Find(names, "Bob")
    if found2 {
        fmt.Printf("Encontrado no índice %d\n", idx2)
    }
}
```

### Exemplo 2: Função Map

```go
// Aplica uma função a cada elemento do slice
func Map[T, U any](slice []T, fn func(T) U) []U {
    result := make([]U, len(slice))
    for i, v := range slice {
        result[i] = fn(v)
    }
    return result
}

func main() {
    numbers := []int{1, 2, 3, 4, 5}
    
    // Dobrar todos os números
    doubled := Map(numbers, func(n int) int {
        return n * 2
    })
    fmt.Println(doubled)  // [2 4 6 8 10]
    
    // Converter para string
    strings := Map(numbers, func(n int) string {
        return fmt.Sprintf("Número: %d", n)
    })
    fmt.Println(strings)  // [Número: 1 Número: 2 ...]
}
```

### Exemplo 3: Função Filter

```go
// Filtra elementos do slice baseado em uma condição
func Filter[T any](slice []T, fn func(T) bool) []T {
    result := make([]T, 0)
    for _, v := range slice {
        if fn(v) {
            result = append(result, v)
        }
    }
    return result
}

func main() {
    numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    
    // Apenas números pares
    evens := Filter(numbers, func(n int) bool {
        return n%2 == 0
    })
    fmt.Println(evens)  // [2 4 6 8 10]
    
    // Apenas números maiores que 5
    greaterThan5 := Filter(numbers, func(n int) bool {
        return n > 5
    })
    fmt.Println(greaterThan5)  // [6 7 8 9 10]
}
```

---

## 5. Generic Types (Tipos Genéricos)

Você também pode criar **tipos genéricos**: structs, slices, maps, etc. que trabalham com type parameters.

### Exemplo 1: Container Genérico

```go
// Container que pode armazenar qualquer tipo
type Container[T any] struct {
    value T
}

func (c *Container[T]) Set(value T) {
    c.value = value
}

func (c *Container[T]) Get() T {
    return c.value
}

func main() {
    // Container de string
    containerStr := Container[string]{}
    containerStr.Set("Hello")
    fmt.Println(containerStr.Get())  // Hello
    
    // Container de int
    containerInt := Container[int]{}
    containerInt.Set(42)
    fmt.Println(containerInt.Get())  // 42
}
```

### Exemplo 2: Stack Genérico

```go
// Stack (pilha) genérico
type Stack[T any] struct {
    items []T
}

func NewStack[T any]() *Stack[T] {
    return &Stack[T]{
        items: make([]T, 0),
    }
}

func (s *Stack[T]) Push(item T) {
    s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() (T, bool) {
    if len(s.items) == 0 {
        var zero T  // Valor zero do tipo T
        return zero, false
    }
    item := s.items[len(s.items)-1]
    s.items = s.items[:len(s.items)-1]
    return item, true
}

func (s *Stack[T]) IsEmpty() bool {
    return len(s.items) == 0
}

func main() {
    stack := NewStack[int]()
    stack.Push(1)
    stack.Push(2)
    stack.Push(3)
    
    for !stack.IsEmpty() {
        val, ok := stack.Pop()
        if ok {
            fmt.Println(val)  // 3, 2, 1
        }
    }
}
```

### Exemplo 3: Pair Genérico

```go
// Par de valores do mesmo tipo
type Pair[T any] struct {
    First  T
    Second T
}

func NewPair[T any](first, second T) *Pair[T] {
    return &Pair[T]{
        First:  first,
        Second: second,
    }
}

func (p *Pair[T]) Swap() {
    p.First, p.Second = p.Second, p.First
}

func main() {
    pair := NewPair(10, 20)
    fmt.Printf("Antes: First=%d, Second=%d\n", pair.First, pair.Second)
    pair.Swap()
    fmt.Printf("Depois: First=%d, Second=%d\n", pair.First, pair.Second)
}
```

---

## 6. Type Inference (Inferência de Tipo)

Uma das características mais poderosas dos generics em Go é a **type inference** (inferência de tipo). O compilador consegue **automaticamente** determinar qual tipo usar baseado nos argumentos da função.

### Inferência Automática

```go
func Max[T constraints.Ordered](a, b T) T {
    if a > b {
        return a
    }
    return b
}

func main() {
    // O compilador INFERE que T é int
    maxInt := Max(10, 20)
    
    // O compilador INFERE que T é string
    maxString := Max("apple", "banana")
    
    // O compilador INFERE que T é float64
    maxFloat := Max(3.14, 2.71)
}
```

Você **não precisa** especificar o tipo explicitamente na maioria dos casos!

### Especificação Explícita (Quando Necessário)

Às vezes você precisa especificar o tipo explicitamente:

```go
func main() {
    // Quando os argumentos não são suficientes para inferir
    var result int = Max[int](10, 20)  // Especificação explícita
    
    // Quando há ambiguidade
    var x interface{} = 10
    // Max(x, 20)  // ERRO: não consegue inferir
    Max[int](x.(int), 20)  // Precisa especificar
}
```

### Quando a Inferência Falha

A inferência pode falhar em alguns casos:

1. **Quando não há argumentos suficientes:**
```go
func Zero[T any]() T {
    var zero T
    return zero
}

func main() {
    // ERRO: não consegue inferir T
    // zero := Zero()
    
    // Solução: especificar explicitamente
    zero := Zero[int]()
}
```

2. **Quando há ambiguidade:**
```go
func main() {
    var x interface{} = 10
    // Max(x, 20)  // ERRO: x é interface{}, não int
    Max[int](x.(int), 20)  // OK
}
```

---

## 7. Generic Interfaces (Interfaces Genéricas)

Interfaces também podem ser genéricas! Isso permite criar contratos que trabalham com tipos específicos.

### Exemplo 1: Interface Adder

```go
// Interface genérica que define um tipo que pode somar
type Adder[T any] interface {
    Add(T) T
}

type Number int

func (n Number) Add(other Number) Number {
    return n + other
}

// Função que usa a interface genérica
func ProcessAdder[T Adder[T]](a, b T) T {
    return a.Add(b)
}

func main() {
    num1 := Number(10)
    num2 := Number(20)
    result := ProcessAdder(num1, num2)
    fmt.Println(result)  // 30
}
```

### Exemplo 2: Interface Comparable

```go
// Interface genérica para tipos comparáveis
type Comparable[T any] interface {
    Compare(T) int  // -1 se menor, 0 se igual, 1 se maior
}

type Pessoa struct {
    Idade int
}

func (p Pessoa) Compare(other Pessoa) int {
    if p.Idade < other.Idade {
        return -1
    }
    if p.Idade > other.Idade {
        return 1
    }
    return 0
}

func MaxComparable[T Comparable[T]](a, b T) T {
    if a.Compare(b) > 0 {
        return a
    }
    return b
}

func main() {
    p1 := Pessoa{Idade: 25}
    p2 := Pessoa{Idade: 30}
    maisVelho := MaxComparable(p1, p2)
    fmt.Printf("Mais velho tem %d anos\n", maisVelho.Idade)  // 30
}
```

---

## 8. Limitações e Considerações

### Limitações dos Generics em Go

1. **Não pode usar generics em métodos de tipos não-genéricos:**
```go
type Pessoa struct {
    Nome string
}

// ERRO: não pode ter método genérico em tipo não-genérico
func (p Pessoa) Process[T any](value T) {
    // ...
}
```

2. **Type parameters não podem ser usados em todos os contextos:**
   - Não pode usar em constantes
   - Algumas limitações em métodos de interfaces

3. **Performance:**
   - Generics são resolvidos em tempo de compilação
   - Não há overhead de runtime
   - Mas podem aumentar o tempo de compilação

### Quando Usar Generics?

**Use generics quando:**
- ✅ Você tem código duplicado para diferentes tipos
- ✅ A lógica é a mesma, apenas o tipo muda
- ✅ Você quer type safety sem perder performance
- ✅ Você está criando estruturas de dados reutilizáveis

**NÃO use generics quando:**
- ❌ A lógica é muito específica para um tipo
- ❌ `interface{}` é suficiente (casos raros)
- ❌ Você está criando código que só funciona com um tipo

---

## 9. Resumo dos Conceitos

### Type Parameters
- Definidos entre colchetes `[]` antes dos parâmetros
- Sintaxe: `func Nome[T Constraint](param T) T`

### Constraints
- `any`: qualquer tipo
- `comparable`: tipos comparáveis com `==` e `!=`
- `constraints.Ordered`: tipos ordenados (com `<`, `>`, etc.)
- Union types: `int | string | float64`
- Interfaces: tipos que implementam métodos específicos

### Generic Functions
- Funções que usam type parameters
- Permitem reutilização de código com type safety

### Generic Types
- Structs, slices, maps, etc. que usam type parameters
- Permitem estruturas de dados reutilizáveis

### Type Inference
- Compilador infere automaticamente os tipos
- Reduz verbosidade do código
- Funciona na maioria dos casos

### Generic Interfaces
- Interfaces que usam type parameters
- Permitem contratos genéricos

---

## 10. Próximos Passos

Agora que você entende os fundamentos de Generics em Go, você está pronto para:
- Criar suas próprias funções e tipos genéricos
- Entender quando usar generics vs interfaces
- Aplicar generics em projetos reais

Na próxima aula, vamos simplificar esses conceitos usando analogias e exemplos do cotidiano para fixar ainda mais o aprendizado!

---

**Dica de Estudo:**
Experimente criar suas próprias funções genéricas! Tente implementar:
- `Min[T constraints.Ordered](a, b T) T`
- `Contains[T comparable](slice []T, value T) bool`
- `Reverse[T any](slice []T) []T`

Quanto mais você praticar, mais natural os generics se tornarão!

