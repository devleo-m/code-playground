# Módulo 39: Unsafe Package em Go
## Aula 4: Performance e Boas Práticas

⚠️ **AVISO**: Este módulo trata de código perigoso. Siga todas as práticas de segurança!

---

## 1. Boas Práticas de Segurança

### ✅ Sempre Valide Antes de Usar

**❌ Ruim: Sem Validação**
```go
func dangerousAccess(arr []int, index int) int {
    ptr := unsafe.Pointer(&arr[0])
    // Sem validação - PERIGOSO!
    elemPtr := unsafe.Pointer(uintptr(ptr) + uintptr(index)*unsafe.Sizeof(int(0)))
    return *(*int)(elemPtr)
}
```

**✅ Bom: Com Validação Completa**
```go
func safeAccess(arr []int, index int) (int, error) {
    if arr == nil {
        return 0, fmt.Errorf("array is nil")
    }
    if index < 0 || index >= len(arr) {
        return 0, fmt.Errorf("index %d out of bounds [0,%d)", index, len(arr))
    }
    
    ptr := unsafe.Pointer(&arr[0])
    elemPtr := unsafe.Pointer(uintptr(ptr) + uintptr(index)*unsafe.Sizeof(int(0)))
    return *(*int)(elemPtr), nil
}
```

### ✅ Sempre Verifique Alinhamento

**❌ Ruim: Sem Verificação de Alinhamento**
```go
func bytesToStruct(b []byte) *MyStruct {
    return (*MyStruct)(unsafe.Pointer(&b[0]))  // Pode estar desalinhado!
}
```

**✅ Bom: Com Verificação**
```go
func bytesToStruct(b []byte) (*MyStruct, error) {
    if len(b) < int(unsafe.Sizeof(MyStruct{})) {
        return nil, fmt.Errorf("buffer too small")
    }
    
    ptr := unsafe.Pointer(&b[0])
    if uintptr(ptr)%unsafe.Alignof(MyStruct{}) != 0 {
        return nil, fmt.Errorf("misaligned buffer")
    }
    
    return (*MyStruct)(ptr), nil
}
```

### ✅ Documente Extensivamente

**❌ Ruim: Sem Documentação**
```go
func convert(x interface{}) []byte {
    // código unsafe...
}
```

**✅ Bom: Com Documentação Completa**
```go
// ⚠️ PERIGOSO: Esta função usa unsafe para conversão zero-copy.
//
// Requisitos:
//   - x deve ser uma struct ou tipo de tamanho fixo
//   - x deve permanecer válido enquanto []byte retornado é usado
//   - []byte retornado NÃO deve ser modificado
//   - Apenas use quando você tem controle total sobre a memória
//
// Riscos:
//   - Modificar []byte pode corromper x
//   - x sendo liberado enquanto []byte é usado causa use-after-free
//   - Comportamento indefinido se requisitos não forem atendidos
//
// Alternativas seguras:
//   - Use encoding/binary para serialização segura
//   - Use json.Marshal para serialização genérica
func convertUnsafe(x interface{}) []byte {
    // código unsafe...
}
```

---

## 2. Padrões de Segurança

### Padrão 1: Wrapper Seguro

```go
// Código unsafe isolado
func unsafeConvert(x *MyStruct) []byte {
    // unsafe aqui
}

// Interface segura para usuários
func Convert(x *MyStruct) ([]byte, error) {
    if x == nil {
        return nil, fmt.Errorf("struct is nil")
    }
    // Validações...
    return unsafeConvert(x), nil
}
```

### Padrão 2: Type-Safe Wrapper

```go
// Wrapper que garante tipos corretos
type SafeConverter struct {
    // campos internos
}

func NewSafeConverter() *SafeConverter {
    return &SafeConverter{}
}

func (sc *SafeConverter) Convert(x *MyStruct) ([]byte, error) {
    // Validações e conversão segura
}
```

---

## 3. Armadilhas Comuns

### ❌ Armadilha 1: Use-After-Free

```go
// ❌ ERRADO: Pointer aponta para memória liberada
func dangerous() unsafe.Pointer {
    x := 42
    return unsafe.Pointer(&x)  // x será liberado!
}

// ✅ CORRETO: Garantir que memória permanece válida
func safe() unsafe.Pointer {
    x := new(int)
    *x = 42
    return unsafe.Pointer(x)  // x permanece válido
}
```

### ❌ Armadilha 2: Buffer Overflow

```go
// ❌ ERRADO: Sem validação de tamanho
func dangerous(b []byte) *MyStruct {
    return (*MyStruct)(unsafe.Pointer(&b[0]))  // Pode acessar além do buffer!
}

// ✅ CORRETO: Validar tamanho
func safe(b []byte) (*MyStruct, error) {
    if len(b) < int(unsafe.Sizeof(MyStruct{})) {
        return nil, fmt.Errorf("buffer too small")
    }
    return (*MyStruct)(unsafe.Pointer(&b[0])), nil
}
```

### ❌ Armadilha 3: Alinhamento Incorreto

```go
// ❌ ERRADO: Pode estar desalinhado
func dangerous(b []byte) *int64 {
    return (*int64)(unsafe.Pointer(&b[1]))  // Offset ímpar!
}

// ✅ CORRETO: Verificar alinhamento
func safe(b []byte) (*int64, error) {
    if len(b) < 8 {
        return nil, fmt.Errorf("buffer too small")
    }
    ptr := unsafe.Pointer(&b[0])
    if uintptr(ptr)%8 != 0 {
        return nil, fmt.Errorf("misaligned")
    }
    return (*int64)(ptr), nil
}
```

---

## 4. Testes Extensivos

### Testar em Diferentes Arquiteturas

```go
func TestUnsafeConversion(t *testing.T) {
    // Testar em diferentes condições
    testCases := []struct {
        name string
        data []byte
        want error
    }{
        {"valid", validBytes, nil},
        {"too small", []byte{1}, errTooSmall},
        {"misaligned", misalignedBytes, errMisaligned},
    }
    
    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            _, err := bytesToStruct(tc.data)
            if err != tc.want {
                t.Errorf("got %v, want %v", err, tc.want)
            }
        })
    }
}
```

### Testar Casos Extremos

```go
func TestEdgeCases(t *testing.T) {
    // Testar com nil
    _, err := bytesToStruct(nil)
    if err == nil {
        t.Error("expected error for nil")
    }
    
    // Testar com buffer vazio
    _, err = bytesToStruct([]byte{})
    if err == nil {
        t.Error("expected error for empty buffer")
    }
    
    // Testar com buffer exatamente do tamanho
    exactSize := make([]byte, unsafe.Sizeof(MyStruct{}))
    _, err = bytesToStruct(exactSize)
    if err != nil {
        t.Errorf("unexpected error: %v", err)
    }
}
```

---

## 5. Checklist de Segurança

### Antes de Usar Unsafe

- [ ] Não há alternativa segura disponível
- [ ] Performance é realmente crítica e medimos que unsafe ajuda
- [ ] Entendemos completamente o que o código faz
- [ ] Validamos todos os inputs
- [ ] Verificamos alinhamento quando necessário
- [ ] Garantimos que memória permanece válida
- [ ] Documentamos extensivamente
- [ ] Isolamos código unsafe
- [ ] Testamos extensivamente
- [ ] Revisamos código com outros desenvolvedores

### Durante Desenvolvimento

- [ ] Código está bem comentado
- [ ] Validações estão em todos os lugares necessários
- [ ] Erros são tratados apropriadamente
- [ ] Testes cobrem casos extremos
- [ ] Documentação explica riscos

### Antes de Merge

- [ ] Code review por desenvolvedor experiente
- [ ] Todos os testes passam
- [ ] Documentação está completa
- [ ] Alternativas foram consideradas
- [ ] Justificativa está documentada

---

## 6. Alternativas ao Unsafe

### Use encoding/binary

```go
// ❌ Não use unsafe para serialização
func unsafeSerialize(x *MyStruct) []byte {
    // unsafe...
}

// ✅ Use encoding/binary
import "encoding/binary"

func safeSerialize(x *MyStruct) ([]byte, error) {
    buf := new(bytes.Buffer)
    err := binary.Write(buf, binary.LittleEndian, x)
    return buf.Bytes(), err
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

---

## Resumo Final

**Princípios fundamentais:**
1. **Valide sempre**: Nunca confie em inputs
2. **Verifique alinhamento**: Especialmente em conversões
3. **Documente extensivamente**: Explique riscos e requisitos
4. **Isole código unsafe**: Wrappers seguros
5. **Teste extensivamente**: Diferentes arquiteturas e casos
6. **Considere alternativas**: Sempre prefira código seguro

**Lembre-se**: Unsafe é uma ferramenta de último recurso. Use apenas quando absolutamente necessário e com extremo cuidado!

---

**Bons estudos e use com cuidado! ⚠️**



