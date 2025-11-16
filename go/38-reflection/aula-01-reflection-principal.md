# Módulo 38: Reflection em Go

## Aula 1: Reflection - Inspeção e Manipulação Dinâmica de Tipos

Olá! Bem-vindo ao módulo 38 sobre **Reflection** (Reflexão) em Go. Até agora, você trabalhou com tipos conhecidos em **tempo de compilação**. Reflection permite inspecionar e manipular tipos e valores em **tempo de execução**, abrindo possibilidades poderosas mas que devem ser usadas com cuidado.

Reflection é usado em bibliotecas como JSON marshaling, ORMs, frameworks web, e ferramentas de validação. Entender reflection é essencial para trabalhar com essas ferramentas e criar suas próprias bibliotecas genéricas.

Nesta aula, vamos explorar:
1. **O que é Reflection**: Conceito e propósito
2. **Package reflect**: A biblioteca padrão de reflection
3. **Type Reflection**: Inspecionar tipos em runtime
4. **Value Reflection**: Inspecionar e modificar valores
5. **Casos de Uso**: Quando e como usar reflection
6. **Limitações e Cuidados**: Performance e armadilhas

---

## 1. O Que É Reflection?

### Definição

**Reflection** é a capacidade de um programa de **inspecionar e modificar** sua própria estrutura em tempo de execução. Em Go, isso significa poder:
- Descobrir o tipo de uma variável em runtime
- Examinar campos de structs
- Chamar métodos dinamicamente
- Criar novos valores de tipos conhecidos apenas em runtime
- Modificar valores de forma dinâmica

### Por Que Reflection Existe?

**Casos de uso:**
- ✅ **JSON/XML marshaling**: Converter structs para JSON sem conhecer campos em compile-time
- ✅ **ORM frameworks**: Mapear structs para tabelas de banco de dados
- ✅ **Validação**: Validar campos de structs dinamicamente
- ✅ **Frameworks web**: Binding de dados de formulários para structs
- ✅ **Code generation**: Ferramentas que geram código baseado em tipos

### Trade-offs

**Vantagens:**
- ✅ Flexibilidade: Trabalhar com tipos desconhecidos em compile-time
- ✅ Poder: Permite criar bibliotecas genéricas poderosas
- ✅ Conveniência: Facilita serialização e validação

**Desvantagens:**
- ⚠️ **Performance**: Reflection é mais lento que código estático
- ⚠️ **Segurança de tipos**: Perde verificação em compile-time
- ⚠️ **Complexidade**: Código mais difícil de entender e manter
- ⚠️ **Erros em runtime**: Erros aparecem apenas quando código executa

---

## 2. Package reflect - Fundamentos

### Importando

```go
import "reflect"
```

### Conceitos Fundamentais

O package `reflect` trabalha com dois tipos principais:

1. **`reflect.Type`**: Representa um tipo Go
2. **`reflect.Value`**: Representa um valor Go

### Obtendo Type e Value

```go
package main

import (
    "fmt"
    "reflect"
)

func main() {
    var x int = 42
    
    // Obter Type
    t := reflect.TypeOf(x)
    fmt.Println("Type:", t)  // int
    
    // Obter Value
    v := reflect.ValueOf(x)
    fmt.Println("Value:", v)  // 42
    fmt.Println("Kind:", v.Kind())  // int
}
```

### Type vs Kind

**Type**: O tipo específico (ex: `int`, `string`, `MyStruct`)
**Kind**: A categoria do tipo (ex: `Int`, `String`, `Struct`)

```go
type MyInt int

var x MyInt = 42

t := reflect.TypeOf(x)
fmt.Println("Type:", t)        // main.MyInt (tipo específico)
fmt.Println("Kind:", t.Kind())  // reflect.Int (categoria)
```

---

## 3. Type Reflection - Inspecionar Tipos

### Tipos Básicos

```go
func inspectType(x interface{}) {
    t := reflect.TypeOf(x)
    fmt.Printf("Type: %s\n", t)
    fmt.Printf("Kind: %s\n", t.Kind())
    fmt.Printf("Name: %s\n", t.Name())
    fmt.Printf("Size: %d bytes\n", t.Size())
}

func main() {
    inspectType(42)           // int
    inspectType("hello")     // string
    inspectType(3.14)        // float64
    inspectType(true)        // bool
}
```

### Structs

```go
type Person struct {
    Name string
    Age  int
}

func inspectStruct(x interface{}) {
    t := reflect.TypeOf(x)
    
    fmt.Printf("Type: %s\n", t)
    fmt.Printf("Kind: %s\n", t.Kind())
    fmt.Printf("NumField: %d\n", t.NumField())
    
    // Iterar sobre campos
    for i := 0; i < t.NumField(); i++ {
        field := t.Field(i)
        fmt.Printf("  Field %d: %s (type: %s)\n", 
            i, field.Name, field.Type)
    }
}

func main() {
    p := Person{Name: "John", Age: 30}
    inspectStruct(p)
}
```

**Output:**
```
Type: main.Person
Kind: struct
NumField: 2
  Field 0: Name (type: string)
  Field 1: Age (type: int)
```

### Acessando Tags de Struct

```go
type User struct {
    Name  string `json:"name" validate:"required"`
    Email string `json:"email" validate:"required,email"`
    Age   int    `json:"age" validate:"min=18"`
}

func inspectTags(x interface{}) {
    t := reflect.TypeOf(x)
    
    for i := 0; i < t.NumField(); i++ {
        field := t.Field(i)
        jsonTag := field.Tag.Get("json")
        validateTag := field.Tag.Get("validate")
        
        fmt.Printf("%s: json=%s, validate=%s\n",
            field.Name, jsonTag, validateTag)
    }
}

func main() {
    u := User{}
    inspectTags(u)
}
```

**Output:**
```
Name: json=name, validate=required
Email: json=email, validate=required,email
Age: json=age, validate=min=18
```

### Slices e Arrays

```go
func inspectSlice(x interface{}) {
    t := reflect.TypeOf(x)
    
    fmt.Printf("Type: %s\n", t)
    fmt.Printf("Kind: %s\n", t.Kind())
    
    if t.Kind() == reflect.Slice {
        elemType := t.Elem()  // Tipo do elemento
        fmt.Printf("Element type: %s\n", elemType)
    }
}

func main() {
    inspectSlice([]int{1, 2, 3})
    inspectSlice([]string{"a", "b"})
}
```

### Maps

```go
func inspectMap(x interface{}) {
    t := reflect.TypeOf(x)
    
    fmt.Printf("Type: %s\n", t)
    fmt.Printf("Kind: %s\n", t.Kind())
    
    if t.Kind() == reflect.Map {
        keyType := t.Key()    // Tipo da chave
        elemType := t.Elem()  // Tipo do valor
        fmt.Printf("Key type: %s\n", keyType)
        fmt.Printf("Value type: %s\n", elemType)
    }
}

func main() {
    inspectMap(map[string]int{"a": 1, "b": 2})
}
```

### Functions

```go
func add(a, b int) int {
    return a + b
}

func inspectFunction(x interface{}) {
    t := reflect.TypeOf(x)
    
    fmt.Printf("Type: %s\n", t)
    fmt.Printf("Kind: %s\n", t.Kind())
    
    if t.Kind() == reflect.Func {
        fmt.Printf("NumIn: %d\n", t.NumIn())   // Número de parâmetros
        fmt.Printf("NumOut: %d\n", t.NumOut()) // Número de retornos
        
        // Tipos dos parâmetros
        for i := 0; i < t.NumIn(); i++ {
            paramType := t.In(i)
            fmt.Printf("  Param %d: %s\n", i, paramType)
        }
        
        // Tipos dos retornos
        for i := 0; i < t.NumOut(); i++ {
            returnType := t.Out(i)
            fmt.Printf("  Return %d: %s\n", i, returnType)
        }
    }
}

func main() {
    inspectFunction(add)
}
```

---

## 4. Value Reflection - Inspecionar e Modificar Valores

### Obtendo Values

```go
func inspectValue(x interface{}) {
    v := reflect.ValueOf(x)
    
    fmt.Printf("Value: %v\n", v)
    fmt.Printf("Kind: %s\n", v.Kind())
    fmt.Printf("Type: %s\n", v.Type())
    fmt.Printf("CanSet: %v\n", v.CanSet())  // Pode modificar?
}

func main() {
    inspectValue(42)
    inspectValue("hello")
}
```

### Acessando Valores de Tipos Básicos

```go
func getValue(x interface{}) {
    v := reflect.ValueOf(x)
    
    switch v.Kind() {
    case reflect.Int:
        fmt.Printf("Int value: %d\n", v.Int())
    case reflect.String:
        fmt.Printf("String value: %s\n", v.String())
    case reflect.Bool:
        fmt.Printf("Bool value: %v\n", v.Bool())
    case reflect.Float64:
        fmt.Printf("Float value: %f\n", v.Float())
    }
}

func main() {
    getValue(42)
    getValue("hello")
    getValue(true)
    getValue(3.14)
}
```

### Modificando Valores

**⚠️ Importante**: Você só pode modificar valores se obtiver um pointer para o valor.

```go
func modifyValue(x interface{}) {
    v := reflect.ValueOf(x)
    
    // v.CanSet() retorna false se x não for pointer
    if !v.CanSet() {
        fmt.Println("Cannot set value directly")
        // Precisamos obter o pointer
        v = v.Elem()  // Obtém o valor apontado
    }
    
    if v.Kind() == reflect.Int && v.CanSet() {
        v.SetInt(100)
        fmt.Printf("New value: %d\n", v.Int())
    }
}

func main() {
    x := 42
    modifyValue(&x)  // Passar pointer!
    fmt.Println("x after modification:", x)  // 100
}
```

### Trabalhando com Structs

```go
type Person struct {
    Name string
    Age  int
}

func modifyStruct(x interface{}) {
    v := reflect.ValueOf(x)
    
    if v.Kind() == reflect.Ptr {
        v = v.Elem()  // Obtém o valor apontado
    }
    
    if v.Kind() == reflect.Struct {
        // Modificar campo Name
        nameField := v.FieldByName("Name")
        if nameField.CanSet() {
            nameField.SetString("Jane")
        }
        
        // Modificar campo Age
        ageField := v.FieldByName("Age")
        if ageField.CanSet() {
            ageField.SetInt(25)
        }
    }
}

func main() {
    p := Person{Name: "John", Age: 30}
    fmt.Println("Before:", p)
    
    modifyStruct(&p)  // Passar pointer!
    fmt.Println("After:", p)  // {Jane 25}
}
```

### Trabalhando com Slices

```go
func modifySlice(x interface{}) {
    v := reflect.ValueOf(x)
    
    if v.Kind() == reflect.Ptr {
        v = v.Elem()
    }
    
    if v.Kind() == reflect.Slice {
        // Adicionar elemento
        newValue := reflect.ValueOf(99)
        v.Set(reflect.Append(v, newValue))
    }
}

func main() {
    slice := []int{1, 2, 3}
    fmt.Println("Before:", slice)
    
    modifySlice(&slice)
    fmt.Println("After:", slice)  // [1 2 3 99]
}
```

### Criando Novos Valores

```go
func createValue(t reflect.Type) reflect.Value {
    // Criar novo valor do tipo t
    return reflect.New(t).Elem()
}

func main() {
    // Criar novo int
    intType := reflect.TypeOf(0)
    intValue := createValue(intType)
    intValue.SetInt(42)
    fmt.Println("New int:", intValue.Int())
    
    // Criar novo string
    stringType := reflect.TypeOf("")
    stringValue := createValue(stringType)
    stringValue.SetString("hello")
    fmt.Println("New string:", stringValue.String())
    
    // Criar novo slice
    sliceType := reflect.TypeOf([]int{})
    sliceValue := reflect.MakeSlice(sliceType, 0, 10)
    sliceValue = reflect.Append(sliceValue, reflect.ValueOf(1))
    sliceValue = reflect.Append(sliceValue, reflect.ValueOf(2))
    fmt.Println("New slice:", sliceValue.Interface())
}
```

---

## 5. Chamando Métodos Dinamicamente

### Métodos de Struct

```go
type Calculator struct {
    result int
}

func (c *Calculator) Add(x int) {
    c.result += x
}

func (c *Calculator) GetResult() int {
    return c.result
}

func callMethod(x interface{}, methodName string, args ...interface{}) []reflect.Value {
    v := reflect.ValueOf(x)
    
    // Obter método
    method := v.MethodByName(methodName)
    if !method.IsValid() {
        panic(fmt.Sprintf("Method %s not found", methodName))
    }
    
    // Converter argumentos para reflect.Value
    argValues := make([]reflect.Value, len(args))
    for i, arg := range args {
        argValues[i] = reflect.ValueOf(arg)
    }
    
    // Chamar método
    return method.Call(argValues)
}

func main() {
    calc := &Calculator{result: 10}
    
    // Chamar Add dinamicamente
    callMethod(calc, "Add", 5)
    fmt.Println("Result:", calc.result)  // 15
    
    // Chamar GetResult dinamicamente
    results := callMethod(calc, "GetResult")
    if len(results) > 0 {
        fmt.Println("GetResult:", results[0].Int())  // 15
    }
}
```

### Funções

```go
func add(a, b int) int {
    return a + b
}

func callFunction(fn interface{}, args ...interface{}) []reflect.Value {
    v := reflect.ValueOf(fn)
    
    if v.Kind() != reflect.Func {
        panic("Not a function")
    }
    
    // Converter argumentos
    argValues := make([]reflect.Value, len(args))
    for i, arg := range args {
        argValues[i] = reflect.ValueOf(arg)
    }
    
    // Chamar função
    return v.Call(argValues)
}

func main() {
    results := callFunction(add, 10, 20)
    if len(results) > 0 {
        fmt.Println("Result:", results[0].Int())  // 30
    }
}
```

---

## 6. Casos de Uso Práticos

### Caso 1: JSON Marshal Customizado

```go
func toJSON(x interface{}) string {
    v := reflect.ValueOf(x)
    t := reflect.TypeOf(x)
    
    if v.Kind() == reflect.Ptr {
        v = v.Elem()
        t = t.Elem()
    }
    
    if v.Kind() != reflect.Struct {
        return "{}"
    }
    
    var result []string
    for i := 0; i < v.NumField(); i++ {
        field := t.Field(i)
        fieldValue := v.Field(i)
        
        // Obter tag JSON
        jsonTag := field.Tag.Get("json")
        if jsonTag == "" {
            jsonTag = field.Name
        }
        
        result = append(result, fmt.Sprintf(`"%s": "%v"`, 
            jsonTag, fieldValue.Interface()))
    }
    
    return "{" + strings.Join(result, ", ") + "}"
}

type User struct {
    Name  string `json:"name"`
    Email string `json:"email"`
    Age   int    `json:"age"`
}

func main() {
    user := User{Name: "John", Email: "john@example.com", Age: 30}
    fmt.Println(toJSON(user))
    // {"name": "John", "email": "john@example.com", "age": "30"}
}
```

### Caso 2: Validação de Structs

```go
func validate(x interface{}) []string {
    var errors []string
    v := reflect.ValueOf(x)
    t := reflect.TypeOf(x)
    
    if v.Kind() == reflect.Ptr {
        v = v.Elem()
        t = t.Elem()
    }
    
    if v.Kind() != reflect.Struct {
        return errors
    }
    
    for i := 0; i < v.NumField(); i++ {
        field := t.Field(i)
        fieldValue := v.Field(i)
        
        // Verificar tag "required"
        if required := field.Tag.Get("validate"); required == "required" {
            if fieldValue.Kind() == reflect.String && fieldValue.String() == "" {
                errors = append(errors, fmt.Sprintf("%s is required", field.Name))
            }
        }
    }
    
    return errors
}

type Form struct {
    Name  string `validate:"required"`
    Email string `validate:"required"`
    Age   int
}

func main() {
    form := Form{Name: "", Email: "test@example.com"}
    errors := validate(form)
    for _, err := range errors {
        fmt.Println(err)
    }
}
```

### Caso 3: Comparação Genérica

```go
func deepEqual(a, b interface{}) bool {
    return reflect.DeepEqual(a, b)
}

func main() {
    // Comparar slices
    fmt.Println(deepEqual([]int{1, 2, 3}, []int{1, 2, 3}))  // true
    
    // Comparar maps
    fmt.Println(deepEqual(
        map[string]int{"a": 1},
        map[string]int{"a": 1}))  // true
    
    // Comparar structs
    type Point struct{ X, Y int }
    fmt.Println(deepEqual(
        Point{X: 1, Y: 2},
        Point{X: 1, Y: 2}))  // true
}
```

---

## 7. Limitações e Cuidados

### Performance

Reflection é **significativamente mais lento** que código estático:

```go
// Código estático (rápido)
func addStatic(a, b int) int {
    return a + b
}

// Código com reflection (lento)
func addReflection(a, b interface{}) interface{} {
    va := reflect.ValueOf(a)
    vb := reflect.ValueOf(b)
    return va.Int() + vb.Int()
}
```

**Regra**: Use reflection apenas quando necessário. Evite em hot paths.

### Segurança de Tipos

Reflection perde verificação em compile-time:

```go
// Isso compila, mas pode causar panic em runtime
func unsafeCall(x interface{}) {
    v := reflect.ValueOf(x)
    method := v.MethodByName("NonExistentMethod")
    method.Call(nil)  // Panic se método não existir!
}
```

**Solução**: Sempre verifique se métodos/fields existem antes de usar.

### Complexidade

Código com reflection é mais difícil de:
- Entender
- Manter
- Debugar
- Testar

**Solução**: Documente bem e considere alternativas (generics, code generation).

---

## 8. Quando Usar Reflection?

### ✅ Use Reflection Quando:

1. **Serialização/Deserialização**: JSON, XML, etc.
2. **ORM/Frameworks**: Mapeamento automático
3. **Validação**: Validação genérica de structs
4. **Code Generation**: Ferramentas que geram código
5. **Debugging Tools**: Ferramentas de inspeção

### ❌ NÃO Use Reflection Quando:

1. **Hot Paths**: Código executado milhões de vezes
2. **Performance Crítica**: Quando performance é essencial
3. **Alternativas Existem**: Generics, interfaces, etc.
4. **Código Simples**: Quando código estático é suficiente

---

## Resumo

Nesta aula aprendemos:

1. **Reflection**: Inspeção e manipulação de tipos em runtime
2. **reflect.Type**: Representa tipos Go
3. **reflect.Value**: Representa valores Go
4. **Type Reflection**: Inspecionar tipos, campos, métodos
5. **Value Reflection**: Inspecionar e modificar valores
6. **Chamadas Dinâmicas**: Chamar métodos e funções dinamicamente
7. **Casos de Uso**: JSON, validação, comparação
8. **Limitações**: Performance, segurança de tipos, complexidade

**Lembre-se**: Reflection é poderoso, mas use com moderação. Prefira código estático quando possível.

---

**Referências:**
- [reflect package](https://pkg.go.dev/reflect)
- [The Laws of Reflection](https://go.dev/blog/laws-of-reflection)
- [Go Reflection Examples](https://golang.org/pkg/reflect/#pkg-examples)

