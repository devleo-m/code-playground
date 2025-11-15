# Aula 13 - Performance e Boas Práticas: Generics

Olá! Agora que você entende os conceitos de Generics, é crucial aprender **quando e como** usá-los de forma eficiente. Nesta aula, vamos explorar aspectos de performance, boas práticas, e os erros comuns que você deve evitar.

---

## 🚀 Performance: Como Generics Funcionam Internamente

### Compilação vs Runtime

**Ponto crucial:** Generics em Go são resolvidos em **tempo de compilação**, não em runtime. Isso significa:

✅ **Sem overhead de runtime** - O código gerado é tão eficiente quanto código não-genérico
✅ **Type safety em tempo de compilação** - Erros são detectados antes de executar
✅ **Zero-cost abstractions** - Você não paga pelo uso de generics em performance

### Como Funciona: Monomorphization

O compilador Go usa uma técnica chamada **monomorphization**:

1. Quando você escreve `Max[int](10, 20)`, o compilador cria uma versão específica para `int`
2. Quando você escreve `Max[string]("a", "b")`, o compilador cria uma versão específica para `string`
3. Cada uso com um tipo diferente gera código específico para aquele tipo

**Resultado:** O código final é idêntico ao que você escreveria manualmente para cada tipo!

```go
// Você escreve:
func Max[T constraints.Ordered](a, b T) T { ... }
Max(10, 20)
Max("a", "b")

// O compilador gera (conceitualmente):
func Max_int(a, b int) int { ... }
func Max_string(a, b string) string { ... }
```

### Impacto no Tempo de Compilação

⚠️ **Atenção:** Generics podem aumentar o tempo de compilação, especialmente em projetos grandes com muitos tipos genéricos.

**Por quê?**
- O compilador precisa gerar código para cada combinação de tipos usada
- Múltiplos type parameters multiplicam as combinações

**Boas práticas:**
- Use generics quando realmente necessário
- Evite criar muitas funções genéricas complexas se não forem reutilizadas

---

## ✅ Boas Práticas: Quando Usar Generics

### ✅ USE Generics Quando:

#### 1. Você Tem Duplicação de Código por Tipo

```go
// ❌ EVITE: Duplicação
func MaxInt(a, b int) int { ... }
func MaxFloat64(a, b float64) float64 { ... }
func MaxString(a, b string) string { ... }

// ✅ USE: Generics
func Max[T constraints.Ordered](a, b T) T { ... }
```

#### 2. Criando Estruturas de Dados Reutilizáveis

```go
// ✅ EXCELENTE uso de generics
type Stack[T any] struct {
    items []T
}

type Queue[T any] struct {
    items []T
}

type Set[T comparable] struct {
    items map[T]bool
}
```

#### 3. Funções Utilitárias que Funcionam com Múltiplos Tipos

```go
// ✅ BOM uso
func Find[T comparable](slice []T, value T) (int, bool) { ... }
func Map[T, U any](slice []T, fn func(T) U) []U { ... }
func Filter[T any](slice []T, fn func(T) bool) []T { ... }
```

#### 4. Algoritmos que Não Dependem de Tipos Específicos

```go
// ✅ BOM uso
func Sort[T constraints.Ordered](slice []T) { ... }
func Reverse[T any](slice []T) { ... }
```

---

## ❌ Evite Generics Quando:

### ❌ NÃO USE Generics Quando:

#### 1. A Lógica é Específica para um Tipo

```go
// ❌ EVITE: Muito específico para string
func ProcessString(s string) string {
    return strings.ToUpper(strings.TrimSpace(s))
}

// Não tente fazer genérico só porque pode!
// func Process[T string](s T) T { ... }  // ❌ Desnecessário
```

#### 2. `interface{}` é Suficiente e Simples

```go
// ✅ Para casos simples, interface{} pode ser suficiente
func Print(value interface{}) {
    fmt.Println(value)
}

// Não precisa de:
// func Print[T any](value T) { ... }  // Over-engineering
```

**Quando `interface{}` é OK:**
- Funções de logging/debugging simples
- Casos onde type safety não é crítico
- Quando a função é muito simples

#### 3. Você Está Criando Abstrações Desnecessárias

```go
// ❌ EVITE: Abstração desnecessária
func DoSomething[T any](value T) {
    // Lógica que não usa T de forma significativa
    fmt.Println("Doing something")
}

// ✅ MELHOR: Função simples
func DoSomething() {
    fmt.Println("Doing something")
}
```

#### 4. A Complexidade Não Justifica

Se criar uma versão genérica torna o código **significativamente mais complexo** sem benefício real, não faça.

---

## 🎯 Escolhendo Constraints: Guia Prático

### Hierarquia de Constraints (Do Mais Restritivo ao Menos)

```
constraints.Ordered  →  comparable  →  any
     (mais específico)              (mais genérico)
```

### Quando Usar Cada Uma:

#### `constraints.Ordered`
**Use quando:** Precisa comparar com `<`, `<=`, `>`, `>=`

```go
// ✅ BOM
func Max[T constraints.Ordered](a, b T) T { ... }
func Min[T constraints.Ordered](a, b T) T { ... }
func Sort[T constraints.Ordered](slice []T) { ... }
```

#### `comparable`
**Use quando:** Precisa comparar com `==` e `!=`

```go
// ✅ BOM
func Find[T comparable](slice []T, value T) (int, bool) { ... }
func Contains[T comparable](slice []T, value T) bool { ... }
func RemoveDuplicates[T comparable](slice []T) []T { ... }
```

#### `any`
**Use quando:** Não precisa de operações específicas, apenas passar/retornar valores

```go
// ✅ BOM
func Print[T any](value T) { ... }
func Identity[T any](value T) T { return value }
func Map[T, U any](slice []T, fn func(T) U) []U { ... }
```

**Regra de ouro:** Use a constraint **mais restritiva possível** que ainda permite o que você precisa fazer.

---

## 🔧 Boas Práticas de Implementação

### 1. Nomes de Type Parameters

**Convenção em Go:**
- Use nomes **curtos e descritivos**
- Comece com letra maiúscula
- Use nomes de uma letra para casos simples: `T`, `U`, `V`
- Use nomes mais descritivos quando há múltiplos: `Key`, `Value`, `Element`

```go
// ✅ BOM
func Max[T constraints.Ordered](a, b T) T { ... }
func Map[T, U any](slice []T, fn func(T) U) []U { ... }

// ✅ TAMBÉM BOM (quando há contexto)
func GetValue[Key comparable, Value any](m map[Key]Value, k Key) (Value, bool) { ... }
```

### 2. Constraints Customizadas: Quando Criar?

**Crie constraints customizadas quando:**
- Você usa a mesma união de tipos em múltiplos lugares
- A constraint tem significado semântico claro
- Melhora a legibilidade do código

```go
// ✅ BOM: Reutilizado em vários lugares
type Numeric interface {
    int | int8 | int16 | int32 | int64 |
    uint | uint8 | uint16 | uint32 | uint64 |
    float32 | float64
}

func Sum[T Numeric](a, b T) T { ... }
func Multiply[T Numeric](a, b T) T { ... }
func Average[T Numeric](slice []T) T { ... }
```

**Evite criar constraints customizadas quando:**
- Usado apenas uma vez
- Não adiciona clareza
- A constraint padrão (`any`, `comparable`, `Ordered`) é suficiente

### 3. Type Inference: Deixe o Compilador Trabalhar

**Sempre que possível, deixe o compilador inferir os tipos:**

```go
// ✅ BOM: Deixe inferir
max := Max(10, 20)

// ❌ DESNECESSÁRIO: Especificação explícita quando não precisa
max := Max[int](10, 20)
```

**Especifique explicitamente apenas quando:**
- A inferência falha
- Você quer deixar o tipo explícito para clareza
- Há ambiguidade

### 4. Evite Over-Engineering

```go
// ❌ EVITE: Muito genérico sem necessidade
func Process[Input any, Output any, Config any](
    input Input,
    config Config,
    fn func(Input, Config) Output,
) Output {
    return fn(input, config)
}

// ✅ MELHOR: Mais simples e direto
func Process(input string, config Config, fn func(string, Config) string) string {
    return fn(input, config)
}
```

**Regra:** Se você não está reutilizando o código genérico com diferentes tipos, provavelmente não precisa de generics.

---

## ⚠️ Erros Comuns e Como Evitá-los

### Erro 1: Usar `any` Quando Precisa de Operações Específicas

```go
// ❌ ERRADO: any não permite comparação
func Max[T any](a, b T) T {
    if a > b {  // ERRO: operador > não definido
        return a
    }
    return b
}

// ✅ CORRETO: Use constraints.Ordered
func Max[T constraints.Ordered](a, b T) T {
    if a > b {
        return a
    }
    return b
}
```

### Erro 2: Tentar Usar Generics em Métodos de Tipos Não-Genéricos

```go
type Pessoa struct {
    Nome string
}

// ❌ ERRADO: Não pode ter método genérico em tipo não-genérico
func (p Pessoa) Process[T any](value T) {
    // ...
}

// ✅ CORRETO: Use função genérica separada
func Process[T any](p Pessoa, value T) {
    // ...
}

// OU: Torne o tipo genérico
type Pessoa[T any] struct {
    Nome string
    Data T
}

func (p *Pessoa[T]) Process(value T) {
    // ...
}
```

### Erro 3: Não Entender Quando Type Inference Falha

```go
// ❌ ERRADO: Type inference falha aqui
func Zero[T any]() T {
    var zero T
    return zero
}

func main() {
    zero := Zero()  // ERRO: não consegue inferir T
}

// ✅ CORRETO: Especifique o tipo
func main() {
    zero := Zero[int]()  // OK
}
```

### Erro 4: Constraints Muito Restritivas

```go
// ❌ PROBLEMÁTICO: Muito restritivo
type OnlyInt interface {
    int
}

func Process[T OnlyInt](value T) { ... }

// Só funciona com int, então por que usar generics?
// ✅ MELHOR: Use int diretamente
func Process(value int) { ... }
```

### Erro 5: Ignorar Type Safety

```go
// ❌ EVITE: Perdendo type safety
func Process(value interface{}) {
    // Type assertions em todo lugar
    if str, ok := value.(string); ok {
        // ...
    }
}

// ✅ MELHOR: Use generics
func Process[T any](value T) {
    // Type-safe, sem assertions
}
```

---

## 🎯 Performance: Benchmarks e Comparações

### Generics vs Código Específico

**Resultado:** Performance idêntica! O código gerado é o mesmo.

```go
// Versão genérica
func Max[T constraints.Ordered](a, b T) T {
    if a > b {
        return a
    }
    return b
}

// Versão específica
func MaxInt(a, b int) int {
    if a > b {
        return a
    }
    return b
}

// Ambas têm a MESMA performance!
```

### Quando Há Diferença?

**Tempo de compilação:**
- Generics podem aumentar o tempo de compilação
- Especialmente com muitos tipos diferentes
- Geralmente não é um problema em projetos normais

**Tamanho do binário:**
- Pode aumentar ligeiramente devido à monomorphization
- Cada tipo usado gera código específico
- Geralmente insignificante

---

## 📚 Padrões Comuns com Generics

### Padrão 1: Container Genérico

```go
type Container[T any] struct {
    value T
}

func (c *Container[T]) Set(value T) { c.value = value }
func (c *Container[T]) Get() T { return c.value }
```

**Quando usar:** Quando você precisa de um wrapper simples para qualquer tipo.

### Padrão 2: Funções de Slice Utilitárias

```go
func Map[T, U any](slice []T, fn func(T) U) []U { ... }
func Filter[T any](slice []T, fn func(T) bool) []T { ... }
func Reduce[T, U any](slice []T, initial U, fn func(U, T) U) U { ... }
```

**Quando usar:** Para operações comuns em slices que você usa frequentemente.

### Padrão 3: Estruturas de Dados Genéricas

```go
type Stack[T any] struct { ... }
type Queue[T any] struct { ... }
type Set[T comparable] struct { ... }
```

**Quando usar:** Quando você precisa de estruturas de dados reutilizáveis.

### Padrão 4: Funções de Comparação

```go
func Max[T constraints.Ordered](a, b T) T { ... }
func Min[T constraints.Ordered](a, b T) T { ... }
func Clamp[T constraints.Ordered](value, min, max T) T { ... }
```

**Quando usar:** Para operações matemáticas e comparações.

---

## 🎓 Resumo: Decisões Rápidas

### Devo usar Generics?

**✅ SIM, se:**
- Você tem código duplicado para diferentes tipos
- A lógica é idêntica, apenas o tipo muda
- Você está criando estruturas de dados reutilizáveis
- Type safety é importante

**❌ NÃO, se:**
- A lógica é específica para um tipo
- `interface{}` é suficiente e simples
- Você está criando abstrações desnecessárias
- A complexidade não justifica

### Qual constraint usar?

**Use `constraints.Ordered` se:** Precisa de `<`, `>`, `<=`, `>=`
**Use `comparable` se:** Precisa de `==`, `!=`
**Use `any` se:** Não precisa de operações específicas

### Type inference ou especificação explícita?

**Deixe inferir sempre que possível!** Especifique apenas quando necessário.

---

## 🚀 Conclusão

Generics em Go são uma ferramenta poderosa que permite:
- ✅ Código reutilizável sem perder type safety
- ✅ Performance idêntica ao código não-genérico
- ✅ Melhor legibilidade e manutenibilidade

**Lembre-se:**
- Use generics quando faz sentido
- Não force o uso de generics
- Prefira simplicidade quando possível
- Escolha constraints apropriadas
- Deixe o compilador inferir tipos quando possível

**A chave é o equilíbrio:** Use generics para eliminar duplicação real e melhorar type safety, mas não sobrecarregue seu código com abstrações desnecessárias.

---

**Próximo passo:** Pratique! Crie suas próprias funções e tipos genéricos. Experimente diferentes constraints. Quanto mais você praticar, mais natural se tornará decidir quando e como usar generics.

Boa sorte! 🎯

