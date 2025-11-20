# Módulo 38: Reflection em Go
## Aula 3 - Exercícios e Reflexão

Agora é hora de colocar a mão na massa! Complete os exercícios abaixo e reflita sobre as questões propostas.

---

## Exercícios Práticos

### Exercício 1: Inspecionar Tipos Básicos

Crie uma função que recebe qualquer valor e imprime informações detalhadas sobre seu tipo.

**Requisitos:**
1. A função deve aceitar `interface{}`
2. Deve imprimir:
   - Nome do tipo
   - Kind (categoria)
   - Tamanho em bytes
   - Se for numérico, mostrar se é signed ou unsigned
   - Se for slice/array, mostrar tipo do elemento
   - Se for map, mostrar tipos da chave e valor

**Código base:**
```go
package main

import (
    "fmt"
    "reflect"
)

func inspectType(x interface{}) {
    // TODO: Implementar inspeção de tipo
}

func main() {
    inspectType(42)
    inspectType("hello")
    inspectType(3.14)
    inspectType(true)
    inspectType([]int{1, 2, 3})
    inspectType(map[string]int{"a": 1})
}
```

**Tarefa**: Complete a função `inspectType` para imprimir todas as informações solicitadas.

---

### Exercício 2: Inspecionar e Modificar Structs

Crie uma função que:
1. Recebe uma struct (via pointer)
2. Lista todos os campos com seus valores atuais
3. Permite modificar campos por nome
4. Valida se o campo existe antes de modificar

**Requisitos:**
- Função `listFields(x interface{})` que lista campos
- Função `setField(x interface{}, fieldName string, value interface{}) error` que modifica campo
- Tratamento de erros apropriado

**Código base:**
```go
package main

import (
    "fmt"
    "reflect"
)

type Person struct {
    Name string
    Age  int
    City string
}

func listFields(x interface{}) {
    // TODO: Listar todos os campos e valores
}

func setField(x interface{}, fieldName string, value interface{}) error {
    // TODO: Modificar campo por nome
    // Retornar erro se campo não existir ou tipo incompatível
}

func main() {
    p := &Person{Name: "John", Age: 30, City: "NYC"}
    
    fmt.Println("Campos atuais:")
    listFields(p)
    
    fmt.Println("\nModificando campos...")
    setField(p, "Name", "Jane")
    setField(p, "Age", 25)
    
    fmt.Println("\nCampos após modificação:")
    listFields(p)
}
```

**Tarefa**: Implemente as funções `listFields` e `setField` com tratamento de erros adequado.

---

### Exercício 3: Validador Genérico

Crie um sistema de validação genérico usando reflection que:
1. Lê tags `validate` de campos de struct
2. Valida campos baseado nas regras:
   - `required`: Campo não pode ser vazio
   - `min=X`: Para números, valor mínimo
   - `max=X`: Para números, valor máximo
   - `email`: Para strings, deve ser email válido (formato básico)

**Requisitos:**
- Função `Validate(x interface{}) []string` que retorna lista de erros
- Suporte para as regras acima
- Mensagens de erro claras

**Código base:**
```go
package main

import (
    "fmt"
    "reflect"
    "strings"
)

type User struct {
    Name     string `validate:"required"`
    Email    string `validate:"required,email"`
    Age      int    `validate:"required,min=18,max=120"`
    Password string `validate:"required,min=8"`
}

func Validate(x interface{}) []string {
    var errors []string
    // TODO: Implementar validação
    return errors
}

func main() {
    user1 := User{
        Name:     "",
        Email:    "invalid-email",
        Age:      15,
        Password: "short",
    }
    
    errors := Validate(user1)
    for _, err := range errors {
        fmt.Println(err)
    }
}
```

**Tarefa**: Implemente a função `Validate` com suporte para todas as regras.

---

### Exercício 4: JSON Marshal Simples

Crie uma função que converte structs para JSON usando reflection (sem usar `encoding/json`).

**Requisitos:**
1. Função `ToJSON(x interface{}) string`
2. Deve respeitar tags `json` para nomes de campos
3. Deve suportar tipos básicos: int, string, bool, float
4. Deve suportar slices de tipos básicos
5. Formatação básica (não precisa ser perfeita)

**Código base:**
```go
package main

import (
    "fmt"
    "reflect"
    "strconv"
)

type Person struct {
    Name  string   `json:"name"`
    Age   int      `json:"age"`
    Hobby []string `json:"hobbies"`
}

func ToJSON(x interface{}) string {
    // TODO: Implementar conversão para JSON
    return ""
}

func main() {
    p := Person{
        Name:  "John",
        Age:   30,
        Hobby: []string{"reading", "coding"},
    }
    
    json := ToJSON(p)
    fmt.Println(json)
    // Esperado: {"name":"John","age":30,"hobbies":["reading","coding"]}
}
```

**Tarefa**: Implemente a função `ToJSON` com suporte para os tipos mencionados.

---

## Perguntas de Reflexão

### Reflexão 1: Trade-offs do Reflection

Reflection oferece flexibilidade, mas tem custos significativos em performance e segurança de tipos.

**Perguntas para refletir**:
1. **Por que reflection é mais lento** que código estático? Quais são os passos adicionais que o runtime precisa fazer?
2. Em que situações o **custo de performance** de reflection é aceitável? Dê exemplos práticos.
3. Como a **perda de verificação em compile-time** afeta a segurança e manutenibilidade do código? Quais são os riscos?
4. Se você estivesse criando uma biblioteca que usa reflection, **como você documentaria** os trade-offs para os usuários?

**Escreva suas reflexões** (mínimo 250 palavras):

---

### Reflexão 2: Quando Reflection é Necessário?

Nem sempre reflection é a melhor solução. Existem alternativas como generics, interfaces, e code generation.

**Perguntas para refletir**:
1. **Quando reflection é realmente necessário** vs quando é apenas conveniente? Dê critérios objetivos.
2. Com a introdução de **generics no Go 1.18**, em que situações generics podem substituir reflection? Dê exemplos.
3. **Code generation** (como `go generate`) pode criar código estático que faz o mesmo que reflection. Quando faz sentido usar cada abordagem?
4. Em um projeto real, como você **decidiria** entre reflection, generics, interfaces, ou code generation? Quais fatores consideraria?

**Escreva suas reflexões** (mínimo 250 palavras):

---

### Reflexão 3: Reflection e Design de API

O uso de reflection em APIs públicas pode ter implicações significativas para usuários da biblioteca.

**Perguntas para refletir**:
1. Se você está criando uma **biblioteca pública** que usa reflection internamente, como você esconderia essa complexidade dos usuários?
2. Como o uso de reflection afeta a **compatibilidade** de uma API? O que acontece se tipos mudam?
3. **Tags de struct** (como `json:"name"`) são uma forma comum de usar reflection. Quais são as vantagens e desvantagens dessa abordagem?
4. Em que situações você **aceitaria** usar reflection em uma API pública vs quando você evitaria? Dê exemplos.

**Escreva suas reflexões** (mínimo 250 palavras):

---

## Checklist de Aprendizado

Marque conforme você completa:

- [ ] Entendi o que é reflection
- [ ] Sei a diferença entre Type e Value
- [ ] Sei inspecionar tipos básicos
- [ ] Sei inspecionar structs e campos
- [ ] Sei ler tags de structs
- [ ] Sei modificar valores usando reflection
- [ ] Sei chamar métodos dinamicamente
- [ ] Entendo as limitações de performance
- [ ] Sei quando usar reflection
- [ ] Posso criar validador genérico
- [ ] Posso criar serializador simples

---

## Desafio Extra (Opcional)

### Desafio: ORM Simples

Crie um ORM simples que:
1. Mapeia structs para tabelas SQL
2. Gera queries SQL baseado em structs
3. Usa tags para mapear campos para colunas
4. Suporta operações básicas: Insert, Select, Update

**Requisitos:**
- Struct com tags `db:"column_name"`
- Função `Insert(x interface{}) string` que gera SQL INSERT
- Função `Select(table string, where map[string]interface{}) string` que gera SQL SELECT
- Tratamento básico de tipos

**Código base:**
```go
type User struct {
    ID    int    `db:"id"`
    Name  string `db:"name"`
    Email string `db:"email"`
}

func Insert(x interface{}) string {
    // Gerar: INSERT INTO users (name, email) VALUES ('John', 'john@example.com')
}

func Select(table string, where map[string]interface{}) string {
    // Gerar: SELECT * FROM users WHERE name = 'John'
}
```

---

## Dicas para os Exercícios

1. **Exercício 1**: Use `reflect.TypeOf()` e `reflect.ValueOf()`. Verifique `Kind()` para diferentes casos.
2. **Exercício 2**: Lembre-se de usar `Elem()` quando receber pointer. Verifique `CanSet()` antes de modificar.
3. **Exercício 3**: Parse tags com `Tag.Get("validate")` e `strings.Split()`. Valide cada regra separadamente.
4. **Exercício 4**: Construa JSON manualmente como string. Use `strconv` para converter números.

---

## Recursos Adicionais

### Documentação
- [reflect package](https://pkg.go.dev/reflect)
- [The Laws of Reflection](https://go.dev/blog/laws-of-reflection)

### Exemplos
- [Go Reflection Examples](https://golang.org/pkg/reflect/#pkg-examples)

---

**Boa sorte com os exercícios! Lembre-se: reflection é poderoso, mas use com moderação.** 🚀



