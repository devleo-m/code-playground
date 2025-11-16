# Módulo 39: Unsafe Package em Go

## Aula 1: Unsafe Package - Manipulação Direta de Memória

Olá! Bem-vindo ao módulo 39 sobre o **Unsafe Package** em Go. Este é um dos tópicos mais avançados e perigosos da linguagem. O package `unsafe` permite contornar as garantias de segurança de tipos e memória do Go, dando acesso direto à memória.

**⚠️ AVISO IMPORTANTE**: O package `unsafe` é chamado assim por um motivo. Seu uso pode causar crashes, vulnerabilidades de segurança, e comportamento indefinido. Use apenas quando absolutamente necessário e com extremo cuidado.

Nesta aula, vamos explorar:
1. **O que é o Unsafe Package**: Conceito e propósito
2. **Por Que Existe**: Casos de uso legítimos
3. **Funcionalidades Principais**: Pointer arithmetic, conversões
4. **Casos de Uso Reais**: Quando e como usar
5. **Riscos e Cuidados**: O que pode dar errado
6. **Alternativas**: Quando não usar unsafe

---

## 1. O Que É o Unsafe Package?

### Definição

O package `unsafe` fornece funcionalidades que **contornam a segurança de tipos** do Go. Ele permite:
- **Pointer arithmetic**: Calcular endereços de memória
- **Conversões de tipos**: Converter entre tipos incompatíveis
- **Acesso direto à memória**: Manipular memória sem verificação de tipos

### Por Que É "Unsafe"?

**Riscos:**
- ⚠️ **Crashes**: Acesso a memória inválida pode causar panic ou segfault
- ⚠️ **Vulnerabilidades**: Buffer overflows, use-after-free, etc.
- ⚠️ **Comportamento indefinido**: Código pode funcionar em uma máquina e não em outra
- ⚠️ **Quebra de garantias**: Perde todas as garantias de segurança do Go

### Quando É Usado?

**Casos legítimos:**
- ✅ **Systems programming**: Interação com código C, drivers
- ✅ **Otimizações críticas**: Quando performance é absolutamente essencial
- ✅ **Estruturas de dados customizadas**: Implementações de baixo nível
- ✅ **Serialização eficiente**: Conversão direta de tipos

---

## 2. Funcionalidades Principais

### unsafe.Pointer

`unsafe.Pointer` é um tipo especial que pode conter o endereço de qualquer valor:

```go
import "unsafe"

var x int = 42
ptr := unsafe.Pointer(&x)  // Converte *int para unsafe.Pointer
```

**Características:**
- Pode ser convertido para/de qualquer tipo de pointer
- Permite pointer arithmetic
- Perde todas as verificações de tipo

### unsafe.Sizeof

Retorna o tamanho em bytes de um tipo ou valor:

```go
import "unsafe"

fmt.Println(unsafe.Sizeof(int(0)))      // 8 (em sistemas 64-bit)
fmt.Println(unsafe.Sizeof("hello"))    // 16 (string header)
fmt.Println(unsafe.Sizeof([]int{}))     // 24 (slice header)
```

### unsafe.Offsetof

Retorna o offset em bytes de um campo em uma struct:

```go
type Person struct {
    Name string  // offset 0
    Age  int     // offset 16 (após string header)
    City string  // offset 24
}

import "unsafe"

p := Person{}
nameOffset := unsafe.Offsetof(p.Name)  // 0
ageOffset := unsafe.Offsetof(p.Age)     // 16
cityOffset := unsafe.Offsetof(p.City)   // 24
```

### unsafe.Alignof

Retorna o alinhamento requerido de um tipo:

```go
import "unsafe"

fmt.Println(unsafe.Alignof(int(0)))     // 8
fmt.Println(unsafe.Alignof(int32(0)))   // 4
fmt.Println(unsafe.Alignof(bool))      // 1
```

---

## 3. Pointer Arithmetic

### Aritmética Básica

Go não permite pointer arithmetic diretamente, mas `unsafe` permite:

```go
import "unsafe"

func pointerArithmetic() {
    arr := [5]int{1, 2, 3, 4, 5}
    
    // Obter pointer para primeiro elemento
    ptr := unsafe.Pointer(&arr[0])
    
    // Avançar para próximo elemento (tamanho de int)
    intSize := unsafe.Sizeof(int(0))
    nextPtr := unsafe.Pointer(uintptr(ptr) + uintptr(intSize))
    
    // Converter de volta para *int
    nextInt := (*int)(nextPtr)
    fmt.Println(*nextInt)  // 2
}
```

### Iterando sobre Array

```go
func iterateArray(arr []int) {
    if len(arr) == 0 {
        return
    }
    
    ptr := unsafe.Pointer(&arr[0])
    intSize := unsafe.Sizeof(int(0))
    
    for i := 0; i < len(arr); i++ {
        // Calcular endereço do elemento i
        elemPtr := unsafe.Pointer(uintptr(ptr) + uintptr(i)*uintptr(intSize))
        elem := (*int)(elemPtr)
        fmt.Printf("arr[%d] = %d\n", i, *elem)
    }
}
```

**⚠️ Cuidado**: Isso é extremamente perigoso. Use apenas quando absolutamente necessário.

---

## 4. Conversões de Tipos

### Conversão Direta de Memória

`unsafe` permite converter entre tipos incompatíveis:

```go
import "unsafe"

func typeConversion() {
    // Converter []byte para []int (PERIGOSO!)
    bytes := []byte{1, 0, 0, 0, 2, 0, 0, 0}
    
    // ⚠️ EXTREMAMENTE PERIGOSO - apenas para demonstração
    ints := (*[2]int)(unsafe.Pointer(&bytes[0]))
    fmt.Println(ints[0], ints[1])  // 1, 2
}
```

**⚠️ Isso pode causar:**
- Crashes se tamanhos não corresponderem
- Comportamento indefinido
- Problemas de alinhamento

### Conversão String ↔ []byte

Um caso mais comum (mas ainda perigoso):

```go
import "unsafe"

func stringToBytes(s string) []byte {
    // ⚠️ PERIGOSO: Modificar o []byte resultante pode corromper a string
    return *(*[]byte)(unsafe.Pointer(&struct {
        data *byte
        len  int
        cap  int
    }{(*byte)(unsafe.Pointer((*reflect.StringHeader)(unsafe.Pointer(&s)).Data)), len(s), len(s)}))
}

func bytesToString(b []byte) string {
    // ⚠️ PERIGOSO: Modificar b após conversão pode corromper a string
    return *(*string)(unsafe.Pointer(&b))
}
```

**⚠️ Cuidado**: Essas conversões são perigosas. A biblioteca padrão tem funções seguras para isso.

---

## 5. Casos de Uso Reais

### Caso 1: Serialização Eficiente

```go
import "unsafe"

type Point struct {
    X, Y float64
}

// Converter struct para []byte sem cópia (PERIGOSO!)
func pointToBytes(p *Point) []byte {
    size := unsafe.Sizeof(*p)
    return (*[unsafe.Sizeof(*p)]byte)(unsafe.Pointer(p))[:size:size]
}

// Converter []byte para struct (PERIGOSO!)
func bytesToPoint(b []byte) *Point {
    return (*Point)(unsafe.Pointer(&b[0]))
}
```

**⚠️ Uso**: Apenas quando você tem controle total sobre a memória e sabe exatamente o que está fazendo.

### Caso 2: Acesso a Estruturas Internas

```go
import (
    "unsafe"
    "reflect"
)

// Acessar campos não exportados (HACK - não recomendado!)
func accessUnexported() {
    // Exemplo hipotético - não funciona com todas as structs
    // Apenas para demonstração do conceito
}
```

**⚠️ Não recomendado**: Quebra encapsulamento e pode quebrar com atualizações do Go.

### Caso 3: Zero-Copy Operations

```go
import "unsafe"

// Reinterpretar memória sem cópia (PERIGOSO!)
func reinterpretMemory() {
    // Apenas para casos extremos de performance
    // Quando você realmente precisa evitar cópias
}
```

---

## 6. Riscos e Cuidados

### Risco 1: Buffer Overflow

```go
// ❌ PERIGOSO: Pode acessar memória além do array
func dangerousAccess(arr []int, index int) int {
    ptr := unsafe.Pointer(&arr[0])
    intSize := unsafe.Sizeof(int(0))
    elemPtr := unsafe.Pointer(uintptr(ptr) + uintptr(index)*uintptr(intSize))
    return *(*int)(elemPtr)  // Pode acessar memória inválida!
}
```

**Solução**: Sempre valide índices e tamanhos.

### Risco 2: Use-After-Free

```go
// ❌ PERIGOSO: Pointer pode apontar para memória liberada
func dangerousPointer() {
    ptr := func() unsafe.Pointer {
        x := 42
        return unsafe.Pointer(&x)  // x será liberado quando função retornar!
    }()
    // Usar ptr aqui é PERIGOSO!
}
```

**Solução**: Garanta que a memória apontada permaneça válida.

### Risco 3: Alinhamento Incorreto

```go
// ❌ PERIGOSO: Pode causar crash em algumas arquiteturas
func misalignedAccess() {
    bytes := []byte{1, 2, 3, 4, 5, 6, 7, 8}
    // Tentar ler int64 a partir de offset ímpar pode crashar!
    int64Ptr := (*int64)(unsafe.Pointer(&bytes[1]))  // PERIGOSO!
}
```

**Solução**: Sempre respeite alinhamento de tipos.

### Risco 4: Comportamento Indefinido

Código com `unsafe` pode:
- Funcionar em uma máquina e não em outra
- Funcionar com uma versão do Go e não com outra
- Ter comportamento diferente em diferentes condições

---

## 7. Boas Práticas

### ✅ Sempre Valide

```go
func safeAccess(arr []int, index int) (int, error) {
    if index < 0 || index >= len(arr) {
        return 0, fmt.Errorf("index out of bounds")
    }
    
    ptr := unsafe.Pointer(&arr[0])
    intSize := unsafe.Sizeof(int(0))
    elemPtr := unsafe.Pointer(uintptr(ptr) + uintptr(index)*uintptr(intSize))
    return *(*int)(elemPtr), nil
}
```

### ✅ Documente Extensivamente

```go
// ⚠️ PERIGOSO: Esta função usa unsafe para conversão zero-copy.
// Requisitos:
// - b deve ter pelo menos sizeof(Point) bytes
// - b não deve ser modificado após conversão
// - Apenas use quando você tem controle total sobre a memória
func bytesToPoint(b []byte) *Point {
    return (*Point)(unsafe.Pointer(&b[0]))
}
```

### ✅ Isole em Funções Específicas

```go
// Encapsular unsafe em função bem definida
func convertPoint(p *Point) []byte {
    // Toda a lógica unsafe aqui
    // Interface limpa para o resto do código
}
```

### ✅ Teste Extensivamente

Teste em:
- Diferentes arquiteturas
- Diferentes versões do Go
- Diferentes condições de memória
- Casos extremos

---

## 8. Alternativas ao Unsafe

### Use Bibliotecas Padrão

```go
// ❌ Não faça isso com unsafe
func stringToBytes(s string) []byte {
    // unsafe...
}

// ✅ Use biblioteca padrão
func stringToBytes(s string) []byte {
    return []byte(s)  // Seguro e eficiente
}
```

### Use Generics (Go 1.18+)

```go
// ❌ Não use unsafe para "generics"
// ✅ Use generics reais
func process[T any](x T) {
    // Sem unsafe!
}
```

### Use Code Generation

```go
// ❌ Não use unsafe para evitar reflection
// ✅ Use code generation
//go:generate go run tools/generate.go
```

---

## 9. Quando NÃO Usar Unsafe

### ❌ NÃO Use Se:

1. **Há alternativa segura**: Sempre prefira código seguro
2. **Performance não é crítica**: Unsafe raramente é necessário
3. **Não entende completamente**: Se não entende 100%, não use
4. **Código de produção geral**: Unsafe é para casos muito específicos

### ✅ Use Apenas Se:

1. **Absolutamente necessário**: Não há alternativa
2. **Performance crítica**: E você mediu que unsafe ajuda
3. **Entende completamente**: Você sabe exatamente o que está fazendo
4. **Casos específicos**: Systems programming, drivers, etc.

---

## 10. Exemplo Completo: Serialização Eficiente

```go
package main

import (
    "fmt"
    "unsafe"
)

type Point struct {
    X, Y float64
}

// ⚠️ PERIGOSO: Converter struct para []byte sem cópia
// Requisitos:
// - p deve permanecer válido enquanto []byte é usado
// - []byte não deve ser modificado
func pointToBytes(p *Point) []byte {
    size := unsafe.Sizeof(*p)
    slice := (*[unsafe.Sizeof(*p)]byte)(unsafe.Pointer(p))[:size:size]
    return slice
}

// ⚠️ PERIGOSO: Converter []byte para struct
// Requisitos:
// - b deve ter pelo menos sizeof(Point) bytes
// - b deve estar alinhado corretamente
// - b não deve ser modificado após conversão
func bytesToPoint(b []byte) (*Point, error) {
    if len(b) < int(unsafe.Sizeof(Point{})) {
        return nil, fmt.Errorf("buffer too small")
    }
    
    // Verificar alinhamento
    if uintptr(unsafe.Pointer(&b[0]))%unsafe.Alignof(Point{}) != 0 {
        return nil, fmt.Errorf("misaligned buffer")
    }
    
    return (*Point)(unsafe.Pointer(&b[0])), nil
}

func main() {
    p := Point{X: 1.5, Y: 2.5}
    
    // Converter para bytes
    bytes := pointToBytes(&p)
    fmt.Printf("Bytes: %v\n", bytes)
    
    // Converter de volta
    p2, err := bytesToPoint(bytes)
    if err != nil {
        fmt.Println("Erro:", err)
        return
    }
    
    fmt.Printf("Point: %+v\n", p2)
}
```

---

## Resumo

Nesta aula aprendemos:

1. **Unsafe Package**: Contorna segurança de tipos do Go
2. **Funcionalidades**: Pointer arithmetic, conversões, acesso direto à memória
3. **Casos de Uso**: Systems programming, otimizações críticas
4. **Riscos**: Crashes, vulnerabilidades, comportamento indefinido
5. **Boas Práticas**: Validar, documentar, isolar, testar
6. **Alternativas**: Sempre prefira código seguro quando possível

**Lembre-se**: Unsafe é chamado assim por um motivo. Use apenas quando absolutamente necessário, com extremo cuidado, e sempre documente extensivamente.

---

**Referências:**
- [unsafe package](https://pkg.go.dev/unsafe)
- [Go Memory Model](https://go.dev/ref/mem)
- [unsafe.Pointer rules](https://pkg.go.dev/unsafe#Pointer)

