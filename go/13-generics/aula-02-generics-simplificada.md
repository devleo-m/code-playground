# Aula 13 - Simplificada: Entendendo Generics

Olá! Vamos simplificar o conceito de Generics usando analogias do dia a dia para que você entenda de forma mais intuitiva.

---

## 🎯 O Problema: Por Que Precisamos de Generics?

### Analogia: A Máquina de Café

Imagine que você tem uma **máquina de café** que só funciona com **café em pó**. Seu amigo tem uma máquina que só funciona com **cápsulas**. Outro amigo tem uma que só funciona com **grãos**.

Todas fazem a mesma coisa: **preparam café**. Mas cada uma só aceita um tipo específico de entrada.

**Sem Generics em Go**, é assim que funcionava:
- Você precisava criar uma função `MaxInt` para números inteiros
- Outra função `MaxFloat` para números decimais  
- Outra função `MaxString` para textos
- Todas fazem a mesma coisa (encontrar o maior), mas cada uma só aceita um tipo!

**Com Generics**, é como ter uma **máquina universal** que aceita café em pó, cápsulas E grãos! Uma única máquina que funciona com diferentes tipos, mas sempre fazendo a mesma coisa: preparar café.

---

## 🔧 O Que São Generics? (Versão Simples)

**Generics** são como uma **receita genérica** que funciona com diferentes ingredientes, mas mantendo a mesma lógica.

### Exemplo do Mundo Real: Receita de Bolo

Pense em uma receita de bolo que diz:
- "Pegue 2 xícaras de **qualquer tipo de farinha**"
- "Adicione 3 ovos"
- "Misture e asse"

Essa receita funciona com:
- Farinha de trigo ✅
- Farinha de amêndoa ✅
- Farinha de coco ✅

A **lógica** (misturar e assar) é a mesma, mas o **tipo de farinha** pode mudar!

**Em Go, é assim:**
```go
// A "receita genérica" - funciona com qualquer tipo T
func Max[T any](a, b T) T {
    if a > b {
        return a
    }
    return b
}
```

Essa função funciona com:
- `int` (números inteiros) ✅
- `float64` (números decimais) ✅
- `string` (textos) ✅

---

## 📦 Type Parameters: Os "Espaços em Branco"

### Analogia: Formulário Preenchível

Imagine um **formulário** onde você precisa preencher:
- Nome: _____________
- Idade: _____________
- Cidade: _____________

Os espaços em branco são como **type parameters** em Go. Eles são "lugares" que você preenche depois com valores específicos.

```go
// T é o "espaço em branco" - você preenche depois
func Print[T any](value T) {
    fmt.Println(value)
}

// Quando você usa:
Print(42)      // T vira "int"
Print("olá")   // T vira "string"
```

**Pense assim:**
- `T` = o espaço em branco no formulário
- Quando você chama `Print(42)`, o Go "preenche" o espaço com "int"
- Quando você chama `Print("olá")`, o Go "preenche" o espaço com "string"

---

## 🚦 Constraints: As "Regras do Jogo"

### Analogia: Restaurante com Código de Vestimenta

Imagine um restaurante que tem **regras** sobre o que você pode vestir:
- ✅ Aceita: camisa, calça, sapatos
- ❌ Não aceita: chinelos, regata, short

Essas regras são como **constraints** em Go. Elas dizem **quais tipos** podem ser usados.

### Exemplo 1: Constraint `any` (Qualquer Coisa)

É como um restaurante **sem código de vestimenta** - aceita qualquer coisa!

```go
func Print[T any](value T) {
    fmt.Println(value)
}
```

**Pode usar com:**
- Números ✅
- Textos ✅
- Structs ✅
- Qualquer coisa! ✅

### Exemplo 2: Constraint `comparable` (Pode Comparar)

É como um restaurante que só aceita pessoas que **podem ser comparadas** (ex: por altura, idade, etc.)

```go
func Equal[T comparable](a, b T) bool {
    return a == b
}
```

**Pode usar com:**
- Números (10 == 10) ✅
- Textos ("a" == "a") ✅
- **NÃO pode usar com:** slices, maps, functions ❌

### Exemplo 3: Constraint `constraints.Ordered` (Pode Ordenar)

É como um restaurante que só aceita coisas que podem ser **ordenadas** (ex: por tamanho, por preço)

```go
func Max[T constraints.Ordered](a, b T) T {
    if a > b {
        return a
    }
    return b
}
```

**Pode usar com:**
- Números (10 > 5) ✅
- Textos ("z" > "a") ✅
- **NÃO pode usar com:** structs, slices ❌

---

## 🏗️ Generic Types: Caixas Universais

### Analogia: Caixas de Armazenamento

Imagine que você tem **caixas de armazenamento** em casa:
- Uma caixa para livros
- Uma caixa para roupas
- Uma caixa para brinquedos

Todas são **caixas** (mesma estrutura), mas cada uma guarda um **tipo diferente** de coisa.

**Sem Generics**, você precisaria criar:
- `CaixaDeLivros`
- `CaixaDeRoupas`
- `CaixaDeBrinquedos`

**Com Generics**, você cria uma **caixa universal**:

```go
// Uma "caixa universal" que pode guardar qualquer tipo T
type Container[T any] struct {
    value T
}

func (c *Container[T]) Set(value T) {
    c.value = value
}

func (c *Container[T]) Get() T {
    return c.value
}
```

Agora você pode ter:
- `Container[string]` - caixa de textos
- `Container[int]` - caixa de números
- `Container[Pessoa]` - caixa de pessoas

**Mesma estrutura, tipos diferentes!**

---

## 🧠 Type Inference: O "Assistente Inteligente"

### Analogia: Garçom que Adivinha

Imagine um **garçom muito esperto** em um restaurante. Você não precisa dizer "quero um café com leite", você só diz "quero um café" e ele **adivinha** que você quer com leite baseado no contexto (hora do dia, seu histórico, etc.)

**Type Inference** em Go é assim! O compilador é o "garçom esperto" que **adivinha** qual tipo você quer usar.

```go
func Max[T constraints.Ordered](a, b T) T {
    if a > b {
        return a
    }
    return b
}

func main() {
    // Você não precisa dizer "Max[int](10, 20)"
    // O Go "adivinha" que é int porque você passou números inteiros!
    max := Max(10, 20)  // Go sabe que T = int
    
    // Você não precisa dizer "Max[string]("a", "b")"
    // O Go "adivinha" que é string porque você passou textos!
    maxStr := Max("a", "b")  // Go sabe que T = string
}
```

**É como o garçom que adivinha seu pedido sem você precisar explicar tudo!**

---

## 🎨 Generic Functions: Ferramentas Universais

### Analogia: Chave Universal

Imagine uma **chave universal** que abre diferentes tipos de portas:
- Porta de casa ✅
- Porta do carro ✅
- Porta do escritório ✅

Uma única chave, múltiplos usos!

**Generic Functions** são como essa chave universal:

```go
// Uma "chave universal" que funciona com qualquer slice
func Find[T comparable](slice []T, value T) (int, bool) {
    for i, v := range slice {
        if v == value {
            return i, true
        }
    }
    return -1, false
}
```

Essa função funciona com:
- `Find([]int{1,2,3}, 2)` - encontrar em números ✅
- `Find([]string{"a","b","c"}, "b")` - encontrar em textos ✅
- `Find([]Pessoa{...}, pessoa)` - encontrar em pessoas ✅

**Uma função, múltiplos tipos!**

---

## 🔄 Comparação: Antes vs Depois

### Antes (Sem Generics): Múltiplas Ferramentas

É como ter uma **caixa de ferramentas** com:
- 🔨 Martelo para pregos pequenos
- 🔨 Martelo para pregos médios
- 🔨 Martelo para pregos grandes

Todos fazem a mesma coisa (martelar), mas cada um só funciona com um tipo de prego!

```go
func MaxInt(a, b int) int { ... }
func MaxFloat(a, b float64) float64 { ... }
func MaxString(a, b string) string { ... }
```

### Depois (Com Generics): Uma Ferramenta Universal

É como ter **um martelo universal** que funciona com qualquer tipo de prego!

```go
func Max[T constraints.Ordered](a, b T) T { ... }
```

**Uma função, todos os tipos!**

---

## 🎯 Resumo com Analogias

| Conceito | Analogia | Exemplo Real |
|----------|---------|--------------|
| **Generics** | Receita genérica | Receita de bolo que funciona com qualquer farinha |
| **Type Parameters** | Espaços em branco | Formulário com campos para preencher |
| **Constraints** | Regras do jogo | Código de vestimenta do restaurante |
| **Generic Types** | Caixas universais | Caixa que guarda qualquer tipo de coisa |
| **Type Inference** | Assistente inteligente | Garçom que adivinha seu pedido |
| **Generic Functions** | Ferramentas universais | Chave que abre qualquer porta |

---

## 💡 Dica Final: Pense em Templates

Generics são como **templates** ou **modelos**:
- Você cria um "molde" (a função genérica)
- Depois "preenche" o molde com tipos específicos
- O resultado é código reutilizável e type-safe!

**Exemplo:**
```go
// O "molde"
func Max[T constraints.Ordered](a, b T) T { ... }

// Preenchendo o molde com int
Max(10, 20)  // T = int

// Preenchendo o molde com string
Max("a", "b")  // T = string
```

---

## 🎓 Conclusão Simplificada

**Generics em Go = Código Reutilizável + Type Safety**

É como ter:
- ✅ Uma receita que funciona com diferentes ingredientes
- ✅ Uma ferramenta que funciona com diferentes materiais
- ✅ Uma caixa que guarda diferentes tipos de coisas

**Tudo isso mantendo a segurança de tipos e a performance!**

Na próxima aula, vamos praticar com exercícios para fixar esses conceitos! 🚀

