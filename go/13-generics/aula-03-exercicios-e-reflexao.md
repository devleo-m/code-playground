# Aula 13 - Exercícios e Reflexão: Generics

Olá! Agora é hora de colocar em prática tudo que você aprendeu sobre Generics. Vamos começar com exercícios práticos e depois refletir sobre os conceitos.

---

## 📝 Exercícios Práticos

### Exercício 1: Função Min Genérica

Crie uma função genérica `Min` que retorna o menor valor entre dois valores de qualquer tipo ordenado.

**Requisitos:**
- Use a constraint `constraints.Ordered`
- A função deve funcionar com `int`, `float64` e `string`
- Teste com pelo menos 3 tipos diferentes

**Exemplo de uso esperado:**
```go
minInt := Min(10, 20)        // 10
minFloat := Min(3.14, 2.71)  // 2.71
minString := Min("zebra", "apple")  // "apple"
```

**Dica:** Use a função `Max` que vimos na aula como referência, mas inverta a lógica.

---

### Exercício 2: Função Contains Genérica

Crie uma função genérica `Contains` que verifica se um valor existe em um slice de qualquer tipo comparável.

**Requisitos:**
- Use a constraint `comparable`
- Retorne `bool` (true se encontrar, false caso contrário)
- Teste com slices de `int` e `string`

**Exemplo de uso esperado:**
```go
numbers := []int{1, 2, 3, 4, 5}
fmt.Println(Contains(numbers, 3))  // true
fmt.Println(Contains(numbers, 10))  // false

names := []string{"Alice", "Bob", "Charlie"}
fmt.Println(Contains(names, "Bob"))    // true
fmt.Println(Contains(names, "David"))  // false
```

---

### Exercício 3: Stack Genérico com Método Peek

Estenda o exemplo de `Stack` que vimos na aula adicionando um método `Peek` que retorna o elemento do topo da pilha **sem removê-lo**.

**Requisitos:**
- Adicione o método `Peek() (T, bool)` ao tipo `Stack[T any]`
- O método deve retornar o elemento do topo e `true` se a pilha não estiver vazia
- Se a pilha estiver vazia, retorne o valor zero de `T` e `false`
- Teste criando uma stack de `int` e uma de `string`

**Exemplo de uso esperado:**
```go
stack := NewStack[int]()
stack.Push(10)
stack.Push(20)

top, ok := stack.Peek()
fmt.Println(top, ok)  // 20 true

// O elemento ainda está na pilha
val, _ := stack.Pop()
fmt.Println(val)  // 20
```

---

### Exercício 4: Função Reverse Genérica

Crie uma função genérica `Reverse` que inverte a ordem dos elementos em um slice.

**Requisitos:**
- Use a constraint `any` (qualquer tipo)
- A função deve modificar o slice original (não criar um novo)
- Teste com slices de diferentes tipos

**Exemplo de uso esperado:**
```go
numbers := []int{1, 2, 3, 4, 5}
Reverse(numbers)
fmt.Println(numbers)  // [5 4 3 2 1]

names := []string{"Alice", "Bob", "Charlie"}
Reverse(names)
fmt.Println(names)  // [Charlie Bob Alice]
```

**Dica:** Use dois índices, um no início e outro no fim, e vá trocando os elementos.

---

## 🤔 Perguntas de Reflexão

### Reflexão 1: Quando Usar Generics vs Interfaces?

Pense sobre a seguinte situação:

Você precisa criar uma função que processa dados. Você tem duas opções:

**Opção A - Usando Interface:**
```go
type Processor interface {
    Process() string
}

func DoWork(p Processor) {
    result := p.Process()
    fmt.Println(result)
}
```

**Opção B - Usando Generics:**
```go
func DoWork[T Processor](p T) {
    result := p.Process()
    fmt.Println(result)
}
```

**Perguntas para reflexão:**
1. Qual é a diferença prática entre essas duas abordagens?
2. Em que situações você escolheria a Opção A (interface)?
3. Em que situações você escolheria a Opção B (generics)?
4. Existe alguma diferença de performance entre elas? Por quê?

**Escreva suas reflexões aqui:**
```
[Seu espaço para escrever]
```

---

### Reflexão 2: O Custo da Abstração

Generics permitem criar código muito reutilizável e genérico. No entanto, há um debate na comunidade Go sobre quando usar generics.

**Cenário:**
Imagine que você está criando uma biblioteca de utilitários. Você pode criar:

1. **Funções específicas:**
```go
func MaxInt(a, b int) int { ... }
func MaxFloat64(a, b float64) float64 { ... }
func MaxString(a, b string) string { ... }
```

2. **Uma função genérica:**
```go
func Max[T constraints.Ordered](a, b T) T { ... }
```

**Perguntas para reflexão:**
1. Quais são as vantagens e desvantagens de cada abordagem?
2. Quando a duplicação de código é aceitável em Go?
3. O princípio "Don't Repeat Yourself" (DRY) sempre se aplica? Ou há situações onde um pouco de duplicação é melhor?
4. Como você decide entre simplicidade e reutilização?

**Escreva suas reflexões aqui:**
```
[Seu espaço para escrever]
```

---

### Reflexão 3: Type Safety e Flexibilidade

Generics em Go foram projetados para manter a type safety (segurança de tipos) enquanto oferecem flexibilidade.

**Cenário:**
Considere estas duas funções:

```go
// Versão 1: Usando interface{}
func Process(value interface{}) {
    // Precisa fazer type assertion
    if str, ok := value.(string); ok {
        fmt.Println("String:", str)
    }
}

// Versão 2: Usando generics
func Process[T any](value T) {
    fmt.Println(value)
}
```

**Perguntas para reflexão:**
1. Qual versão oferece mais type safety? Por quê?
2. Em que situações `interface{}` ainda pode ser a escolha correta?
3. Como você explica a diferença entre type safety em tempo de compilação vs runtime?
4. Por que type safety em tempo de compilação é geralmente preferível?

**Escreva suas reflexões aqui:**
```
[Seu espaço para escrever]
```

---

### Reflexão 4: Generics e a Filosofia de Go

Go é conhecida por sua simplicidade e filosofia de "menos é mais". A adição de generics no Go 1.18 foi uma decisão controversa na comunidade.

**Perguntas para reflexão:**
1. Você acha que generics tornam Go mais complexo? Por quê?
2. Como você equilibra o poder dos generics com a simplicidade que Go valoriza?
3. Quando você acha que é apropriado usar generics em um projeto Go?
4. Existe algum risco de "over-engineering" ao usar generics? Como você evita isso?

**Escreva suas reflexões aqui:**
```
[Seu espaço para escrever]
```

---

## 📋 Checklist de Aprendizado

Antes de prosseguir, verifique se você consegue:

- [ ] Explicar o problema que generics resolve em Go
- [ ] Criar uma função genérica com type parameters
- [ ] Diferenciar entre as constraints `any`, `comparable` e `constraints.Ordered`
- [ ] Criar um tipo genérico (struct, slice, etc.)
- [ ] Entender como type inference funciona
- [ ] Decidir quando usar generics vs interfaces
- [ ] Escrever código genérico que seja type-safe

---

## 🎯 Desafio Extra (Opcional)

Se você completou todos os exercícios e quer um desafio maior:

### Desafio: Implementar uma Queue (Fila) Genérica

Crie uma estrutura de dados `Queue` (fila) genérica com os seguintes métodos:

- `NewQueue[T any]() *Queue[T]` - Cria uma nova fila vazia
- `Enqueue(item T)` - Adiciona um elemento ao final da fila
- `Dequeue() (T, bool)` - Remove e retorna o elemento do início da fila
- `IsEmpty() bool` - Verifica se a fila está vazia
- `Size() int` - Retorna o número de elementos na fila
- `Peek() (T, bool)` - Retorna o elemento do início sem removê-lo

**Dica:** Use um slice interno para armazenar os elementos. Para `Dequeue`, você pode usar `slice[1:]` para remover o primeiro elemento, mas isso pode ser ineficiente. Pense em uma solução melhor!

---

## 📝 Instruções para Entrega

1. **Exercícios Práticos:**
   - Crie um arquivo `exercicios.go` na pasta `13-generics`
   - Implemente todos os 4 exercícios
   - Adicione comentários explicando sua solução
   - Teste cada função com diferentes tipos

2. **Reflexões:**
   - Responda todas as perguntas de reflexão
   - Seja honesto e detalhado em suas respostas
   - Não há resposta "certa" ou "errada" - o importante é pensar criticamente

3. **Entrega:**
   - Envie seu código e suas reflexões
   - Esteja preparado para discutir suas escolhas e pensamentos

---

**Boa sorte com os exercícios! Lembre-se: o objetivo não é apenas fazer funcionar, mas entender o "porquê" por trás de cada decisão.** 🚀

