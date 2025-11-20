# Módulo 42: CGO Basics em Go

## Aula 1: CGO Basics - Integrando Go com C

Olá! Bem-vindo ao módulo 42 sobre **CGO Basics** em Go. CGO (C Go) permite que programas Go chamem código C e vice-versa, abrindo possibilidades de integração com bibliotecas C existentes.

**⚠️ IMPORTANTE**: CGO tem trade-offs significativos. Use apenas quando realmente necessário.

Nesta aula, vamos explorar:
1. **O que é CGO**: Conceito e propósito
2. **Como Funciona**: Sintaxe e mecanismos
3. **Chamando C de Go**: Integração básica
4. **Chamando Go de C**: Callbacks
5. **Limitações e Trade-offs**: Quando não usar
6. **Casos de Uso**: Quando CGO faz sentido

---

## 1. O Que É CGO?

### Definição

**CGO** é uma ferramenta que permite:
- ✅ Chamar funções C de código Go
- ✅ Chamar funções Go de código C
- ✅ Usar tipos C em Go
- ✅ Compartilhar memória entre Go e C

### Por Que Existe?

**Casos de uso:**
- ✅ **Bibliotecas C existentes**: Usar bibliotecas C sem reescrever
- ✅ **Performance crítica**: Código C pode ser mais rápido em casos específicos
- ✅ **Sistemas programming**: Acesso direto a APIs do sistema
- ✅ **Legacy code**: Integrar com código C antigo

### Trade-offs

**Desvantagens:**
- ⚠️ **Desabilita cross-compilation**: Precisa compilar na plataforma alvo
- ⚠️ **Performance overhead**: Chamadas CGO têm custo
- ⚠️ **Complexidade**: Código mais complexo de manter
- ⚠️ **Deploy complicado**: Precisa de toolchain C
- ⚠️ **Goroutines**: Código C não pode usar goroutines diretamente

---

## 2. Sintaxe Básica

### Import "C"

Para usar CGO, você precisa importar o package especial `"C"`:

```go
package main

/*
#include <stdio.h>
#include <stdlib.h>

void hello() {
    printf("Hello from C!\n");
}
*/
import "C"

func main() {
    C.hello()
}
```

**Importante**: 
- O comentário `/* */` antes de `import "C"` contém código C
- Deve estar imediatamente antes do import
- Pode incluir headers C, funções, etc.

### Compilar

```bash
go build main.go
```

CGO é habilitado automaticamente quando detecta `import "C"`.

---

## 3. Chamando Funções C

### Funções Simples

```go
package main

/*
#include <stdio.h>
#include <stdlib.h>

int add(int a, int b) {
    return a + b;
}
*/
import "C"
import "fmt"

func main() {
    result := C.add(10, 20)
    fmt.Printf("Result: %d\n", result)
}
```

### Usando Headers C

```go
package main

/*
#include <math.h>
#include <string.h>
*/
import "C"
import "fmt"

func main() {
    // Usar função da math.h
    result := C.sqrt(16.0)
    fmt.Printf("sqrt(16) = %f\n", result)
    
    // Usar função da string.h
    str1 := C.CString("Hello")
    str2 := C.CString("World")
    result2 := C.strcmp(str1, str2)
    fmt.Printf("strcmp result: %d\n", result2)
    
    // Liberar memória
    C.free(unsafe.Pointer(str1))
    C.free(unsafe.Pointer(str2))
}
```

**⚠️ Importante**: Strings C precisam ser liberadas com `C.free()`!

---

## 4. Tipos C em Go

### Conversão de Tipos

```go
package main

/*
#include <stdint.h>
*/
import "C"
import "fmt"

func main() {
    // Go int para C int
    goInt := 42
    cInt := C.int(goInt)
    fmt.Printf("C int: %d\n", cInt)
    
    // C int para Go int
    goInt2 := int(cInt)
    fmt.Printf("Go int: %d\n", goInt2)
}
```

### Tipos Comuns

```go
// Tipos C comuns
C.char
C.schar
C.uchar
C.short
C.ushort
C.int
C.uint
C.long
C.ulong
C.longlong
C.ulonglong
C.float
C.double
C.complexfloat
C.complexdouble
```

### Strings

```go
package main

/*
#include <stdlib.h>
#include <string.h>
*/
import "C"
import "unsafe"

func main() {
    // Go string para C string
    goStr := "Hello, World!"
    cStr := C.CString(goStr)
    defer C.free(unsafe.Pointer(cStr))  // Sempre liberar!
    
    // Usar C string
    length := C.strlen(cStr)
    fmt.Printf("Length: %d\n", length)
    
    // C string para Go string
    goStr2 := C.GoString(cStr)
    fmt.Printf("Go string: %s\n", goStr2)
}
```

**⚠️ CRÍTICO**: Sempre libere strings C com `C.free()`!

---

## 5. Usando Bibliotecas C Externas

### Linkar com Biblioteca C

```go
package main

/*
#cgo LDFLAGS: -lm
#include <math.h>
*/
import "C"
import "fmt"

func main() {
    result := C.sqrt(25.0)
    fmt.Printf("sqrt(25) = %f\n", result)
}
```

**Compilar:**
```bash
go build main.go
```

### Múltiplas Flags

```go
/*
#cgo CFLAGS: -I./include
#cgo LDFLAGS: -L./lib -lmylib
#include "mylib.h"
*/
import "C"
```

---

## 6. Passando Dados Complexos

### Structs

```go
package main

/*
typedef struct {
    int x;
    int y;
} Point;

Point create_point(int x, int y) {
    Point p;
    p.x = x;
    p.y = y;
    return p;
}
*/
import "C"
import "fmt"

func main() {
    p := C.create_point(10, 20)
    fmt.Printf("Point: (%d, %d)\n", p.x, p.y)
}
```

### Arrays e Slices

```go
package main

/*
#include <stdlib.h>

void process_array(int* arr, int len) {
    for (int i = 0; i < len; i++) {
        arr[i] *= 2;
    }
}
*/
import "C"
import "unsafe"

func main() {
    goSlice := []int32{1, 2, 3, 4, 5}
    
    // Converter slice para array C
    cArray := (*C.int)(unsafe.Pointer(&goSlice[0]))
    length := C.int(len(goSlice))
    
    // Chamar função C
    C.process_array(cArray, length)
    
    // Slice Go foi modificado!
    fmt.Println(goSlice)  // [2, 4, 6, 8, 10]
}
```

**⚠️ Cuidado**: Modificar arrays C pode modificar slices Go!

---

## 7. Callbacks: Chamando Go de C

### Exportar Função Go para C

```go
package main

/*
#include <stdio.h>

extern void goCallback();

void call_go() {
    printf("Calling Go from C...\n");
    goCallback();
}
*/
import "C"
import "fmt"

//export goCallback
func goCallback() {
    fmt.Println("Hello from Go callback!")
}

func main() {
    C.call_go()
}
```

**⚠️ Importante**: 
- Função deve ter comentário `//export`
- Nome da função em C deve ser igual ao nome em Go (lowercase)

---

## 8. Limitações e Trade-offs

### Limitação 1: Cross-compilation

```bash
# ❌ Não funciona com CGO
GOOS=linux GOARCH=arm64 go build main.go

# ✅ Precisa compilar na plataforma alvo
go build main.go  # Na máquina Linux ARM64
```

**Solução**: Use Docker ou compile na plataforma alvo.

### Limitação 2: Performance Overhead

Chamadas CGO têm overhead significativo:
- ~100-200ns por chamada
- Não pode ser inlined
- Precisa de context switch

**Solução**: Minimize chamadas CGO, faça batch quando possível.

### Limitação 3: Goroutines

Código C não pode usar goroutines diretamente:

```go
// ❌ Não funciona
go func() {
    C.some_c_function()  // Bloqueia goroutine
}()
```

**Solução**: Use channels ou callbacks.

### Limitação 4: Garbage Collector

Código C não é gerenciado pelo GC do Go:

```go
// ⚠️ Precisa gerenciar memória manualmente
cStr := C.CString("hello")
defer C.free(unsafe.Pointer(cStr))  // Sempre liberar!
```

---

## 9. Quando Usar CGO?

### ✅ Use CGO Quando:

1. **Biblioteca C existente**: Não há alternativa em Go puro
2. **Performance crítica**: E você mediu que CGO ajuda
3. **Sistemas programming**: Acesso direto a APIs do sistema
4. **Legacy code**: Integração com código C antigo

### ❌ NÃO Use CGO Se:

1. **Há alternativa em Go**: Sempre prefira Go puro
2. **Cross-compilation importante**: CGO complica muito
3. **Performance não é crítica**: Overhead não vale a pena
4. **Deploy simples**: CGO complica deployment

---

## 10. Alternativas ao CGO

### Use Go Puro Quando Possível

```go
// ❌ Não use CGO para coisas simples
// ✅ Use Go puro
import "math"

result := math.Sqrt(16.0)
```

### Use Wrappers Go

Muitas bibliotecas C têm wrappers Go:
- OpenSSL → `crypto/tls`
- SQLite → `database/sqlite3`
- etc.

### Use FFI (Future)

Go pode ter FFI no futuro, que seria mais eficiente que CGO.

---

## 11. Exemplo Completo

```go
package main

/*
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

typedef struct {
    char* name;
    int age;
} Person;

Person* create_person(char* name, int age) {
    Person* p = malloc(sizeof(Person));
    p->name = strdup(name);
    p->age = age;
    return p;
}

void print_person(Person* p) {
    printf("Name: %s, Age: %d\n", p->name, p->age);
}

void free_person(Person* p) {
    free(p->name);
    free(p);
}
*/
import "C"
import (
    "fmt"
    "unsafe"
)

func main() {
    // Criar pessoa em C
    name := C.CString("John")
    defer C.free(unsafe.Pointer(name))
    
    person := C.create_person(name, 30)
    defer C.free_person(person)
    
    // Usar pessoa
    C.print_person(person)
    
    // Acessar campos
    goName := C.GoString(person.name)
    goAge := int(person.age)
    fmt.Printf("Go: Name=%s, Age=%d\n", goName, goAge)
}
```

---

## Resumo

Nesta aula aprendemos:

1. **CGO**: Permite chamar C de Go e vice-versa
2. **Sintaxe**: `import "C"` com comentário C antes
3. **Funções C**: Chamar funções C de Go
4. **Tipos**: Conversão entre tipos Go e C
5. **Bibliotecas**: Linkar com bibliotecas C externas
6. **Callbacks**: Chamar Go de C
7. **Limitações**: Cross-compilation, performance, etc.
8. **Quando usar**: Apenas quando necessário

**Lembre-se**: CGO é poderoso, mas tem trade-offs significativos. Use apenas quando realmente necessário e sempre prefira Go puro quando possível!

---

**Referências:**
- [CGO Documentation](https://pkg.go.dev/cmd/cgo)
- [CGO Best Practices](https://github.com/golang/go/wiki/Cgo)



