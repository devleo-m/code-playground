# Módulo 38: Reflection em Go
## Aula 4: Performance e Boas Práticas

Nesta aula, vamos focar em **otimizações práticas**, **boas práticas** e **armadilhas comuns** relacionadas a reflection em Go. Essas são lições aprendidas de projetos reais e podem fazer a diferença entre código eficiente e código lento.

---

## 1. Boas Práticas de Reflection

### ✅ Cache de Types e Values

**❌ Ruim: Recalcular Type Toda Vez**
```go
func process(x interface{}) {
    t := reflect.TypeOf(x)  // Recalcula toda vez
    // processar...
}

// Chamado milhões de vezes
for i := 0; i < 1000000; i++ {
    process(data)
}
```

**✅ Bom: Cache de Type**
```go
var cachedType reflect.Type

func init() {
    cachedType = reflect.TypeOf(MyStruct{})
}

func process(x interface{}) {
    // Usa type cached
    // processar...
}
```

**Benefício**: Evita recalcular type repetidamente.

### ✅ Evite Reflection em Hot Paths

**❌ Ruim: Reflection em Loop Crítico**
```go
func processItems(items []interface{}) {
    for _, item := range items {
        v := reflect.ValueOf(item)  // Reflection a cada iteração!
        // processar...
    }
}
```

**✅ Bom: Pré-processar ou Usar Tipos Concretos**
```go
// Opção 1: Pré-processar
type ProcessedItem struct {
    // campos já extraídos
}

func processItems(items []ProcessedItem) {
    for _, item := range items {
        // Sem reflection!
    }
}

// Opção 2: Usar tipos concretos
func processItems(items []MyType) {
    for _, item := range items {
        // Sem reflection!
    }
}
```

**Benefício**: Performance muito melhor em loops.

### ✅ Valide Antes de Usar

**❌ Ruim: Assumir que Existe**
```go
func callMethod(x interface{}, name string) {
    method := reflect.ValueOf(x).MethodByName(name)
    method.Call(nil)  // Panic se não existir!
}
```

**✅ Bom: Validar Primeiro**
```go
func callMethod(x interface{}, name string) error {
    v := reflect.ValueOf(x)
    method := v.MethodByName(name)
    
    if !method.IsValid() {
        return fmt.Errorf("method %s not found", name)
    }
    
    method.Call(nil)
    return nil
}
```

**Benefício**: Evita panics e permite tratamento de erros.

### ✅ Use Type Assertions Quando Possível

**❌ Ruim: Reflection para Tipos Conhecidos**
```go
func getString(x interface{}) string {
    v := reflect.ValueOf(x)
    return v.String()  // Reflection desnecessário
}
```

**✅ Bom: Type Assertion**
```go
func getString(x interface{}) string {
    if s, ok := x.(string); ok {
        return s  // Muito mais rápido!
    }
    return ""
}
```

**Benefício**: Type assertion é muito mais rápido que reflection.

---

## 2. Padrões de Otimização

### Padrão 1: Cache de Field Indexes

**❌ Ruim: Buscar Field por Nome Toda Vez**
```go
func getField(x interface{}, name string) reflect.Value {
    v := reflect.ValueOf(x)
    return v.FieldByName(name)  // Busca toda vez
}
```

**✅ Bom: Cache de Index**
```go
type fieldCache struct {
    index int
    field reflect.StructField
}

var cache = make(map[reflect.Type]map[string]fieldCache)

func getFieldCached(x interface{}, name string) reflect.Value {
    t := reflect.TypeOf(x)
    
    // Verificar cache
    if typeCache, ok := cache[t]; ok {
        if fieldInfo, ok := typeCache[name]; ok {
            return reflect.ValueOf(x).Field(fieldInfo.index)
        }
    }
    
    // Buscar e cachear
    field, _ := t.FieldByName(name)
    // ... cachear ...
    
    return reflect.ValueOf(x).Field(field.Index[0])
}
```

**Benefício**: Evita busca repetida de fields.

### Padrão 2: Usar reflect.DeepEqual com Cuidado

**❌ Ruim: DeepEqual em Hot Path**
```go
func compare(a, b interface{}) bool {
    return reflect.DeepEqual(a, b)  // Pode ser lento
}
```

**✅ Bom: Comparação Específica**
```go
func compare(a, b MyType) bool {
    return a.Field1 == b.Field1 && a.Field2 == b.Field2
}
```

**Benefício**: Comparação direta é muito mais rápida.

### Padrão 3: Evitar Criar Values Desnecessariamente

**❌ Ruim: Criar Value Múltiplas Vezes**
```go
func process(x interface{}) {
    v1 := reflect.ValueOf(x)  // Criação 1
    // ...
    v2 := reflect.ValueOf(x)  // Criação 2 (desnecessária!)
    // ...
}
```

**✅ Bom: Reutilizar Value**
```go
func process(x interface{}) {
    v := reflect.ValueOf(x)  // Uma única criação
    // usar v múltiplas vezes
}
```

**Benefício**: Reduz alocações.

---

## 3. Armadilhas Comuns

### ❌ Armadilha 1: Esquecer de Usar Elem() para Pointers

```go
// ❌ ERRADO
func modify(x interface{}) {
    v := reflect.ValueOf(x)
    v.SetInt(100)  // Erro! x é pointer, precisa de Elem()
}

// ✅ CORRETO
func modify(x interface{}) {
    v := reflect.ValueOf(x)
    v = v.Elem()  // Obter valor apontado
    v.SetInt(100)  // Agora funciona
}
```

### ❌ Armadilha 2: Não Verificar CanSet()

```go
// ❌ ERRADO
func modify(x interface{}) {
    v := reflect.ValueOf(x)
    v.SetInt(100)  // Panic se não puder setar!
}

// ✅ CORRETO
func modify(x interface{}) error {
    v := reflect.ValueOf(x)
    if !v.CanSet() {
        return fmt.Errorf("cannot set value")
    }
    v.SetInt(100)
    return nil
}
```

### ❌ Armadilha 3: Assumir Kind Sempre Corresponde

```go
// ⚠️ CUIDADO
func process(x interface{}) {
    v := reflect.ValueOf(x)
    if v.Kind() == reflect.Int {
        // Mas x pode ser int8, int16, int32, int64!
        // Todos têm Kind() == reflect.Int
    }
}
```

**Solução**: Use `Type()` quando precisar do tipo exato.

### ❌ Armadilha 4: Reflection em Generics

Com Go 1.18+, generics podem substituir reflection em muitos casos:

```go
// ❌ Reflection (lento)
func process(x interface{}) {
    v := reflect.ValueOf(x)
    // ...
}

// ✅ Generics (rápido)
func process[T any](x T) {
    // Sem reflection!
}
```

---

## 4. Quando NÃO Usar Reflection

### ❌ Não Use Reflection Se:

1. **Tipos são conhecidos**: Use tipos concretos ou generics
2. **Performance crítica**: Use código estático
3. **Alternativas existem**: Generics, interfaces, code generation
4. **Código simples**: Quando código normal é suficiente

### ✅ Use Reflection Quando:

1. **Tipos desconhecidos**: Serialização genérica (JSON, XML)
2. **Bibliotecas genéricas**: ORM, validação, frameworks
3. **Ferramentas**: Code generation, debugging tools
4. **Não há alternativa**: Quando realmente necessário

---

## 5. Exemplos Práticos de Otimização

### Exemplo 1: JSON Marshal Otimizado

**Antes:**
```go
func toJSON(x interface{}) string {
    v := reflect.ValueOf(x)  // Reflection toda vez
    t := reflect.TypeOf(x)
    // construir JSON...
}
```

**Depois:**
```go
type encoder struct {
    fields []fieldInfo
}

type fieldInfo struct {
    name  string
    index int
    typ   reflect.Type
}

var encoderCache = make(map[reflect.Type]*encoder)

func getEncoder(t reflect.Type) *encoder {
    if enc, ok := encoderCache[t]; ok {
        return enc  // Cache hit!
    }
    
    // Construir encoder uma vez
    enc := &encoder{}
    // ... processar campos ...
    encoderCache[t] = enc
    return enc
}

func toJSON(x interface{}) string {
    t := reflect.TypeOf(x)
    enc := getEncoder(t)  // Usa cache
    // construir JSON usando encoder cached...
}
```

**Benefício**: Processa campos apenas uma vez por tipo.

### Exemplo 2: Validador Otimizado

**Antes:**
```go
func validate(x interface{}) []string {
    t := reflect.TypeOf(x)  // Toda vez
    v := reflect.ValueOf(x)  // Toda vez
    // validar...
}
```

**Depois:**
```go
type validator struct {
    rules []rule
}

type rule struct {
    fieldIndex int
    validators []func(reflect.Value) error
}

var validatorCache = make(map[reflect.Type]*validator)

func getValidator(t reflect.Type) *validator {
    if val, ok := validatorCache[t]; ok {
        return val
    }
    
    // Construir validator uma vez
    val := &validator{}
    // ... processar tags e criar rules ...
    validatorCache[t] = val
    return val
}

func validate(x interface{}) []string {
    t := reflect.TypeOf(x)
    v := reflect.ValueOf(x)
    val := getValidator(t)  // Usa cache
    // validar usando rules cached...
}
```

**Benefício**: Processa tags apenas uma vez por tipo.

---

## 6. Checklist de Boas Práticas

### Performance
- [ ] Cache de Types quando possível
- [ ] Evite reflection em hot paths
- [ ] Use type assertions quando tipos são conhecidos
- [ ] Reutilize Values ao invés de criar novos
- [ ] Cache de field indexes e validators

### Segurança
- [ ] Sempre valide antes de usar (IsValid, CanSet)
- [ ] Trate erros apropriadamente
- [ ] Use Elem() para pointers quando necessário
- [ ] Verifique Kind() antes de operações

### Design
- [ ] Esconda reflection em APIs públicas
- [ ] Documente trade-offs de performance
- [ ] Considere alternativas (generics, code generation)
- [ ] Use reflection apenas quando necessário

### Manutenção
- [ ] Código bem documentado
- [ ] Testes abrangentes
- [ ] Exemplos claros
- [ ] Performance benchmarks

---

## 7. Alternativas ao Reflection

### Generics (Go 1.18+)

```go
// ❌ Reflection
func process(x interface{}) {
    v := reflect.ValueOf(x)
    // ...
}

// ✅ Generics
func process[T any](x T) {
    // Sem reflection!
}
```

### Code Generation

```go
//go:generate go run tools/generate.go

// Gera código estático baseado em structs
// Sem reflection em runtime!
```

### Interfaces

```go
// ❌ Reflection para descobrir tipo
func process(x interface{}) {
    v := reflect.ValueOf(x)
    // ...
}

// ✅ Interface específica
type Processor interface {
    Process()
}

func process(p Processor) {
    p.Process()  // Sem reflection!
}
```

---

## 8. Resumo Final

**Princípios fundamentais:**
1. **Cache quando possível**: Types, fields, validators
2. **Evite em hot paths**: Use tipos concretos ou generics
3. **Valide sempre**: Verifique IsValid, CanSet antes de usar
4. **Considere alternativas**: Generics, code generation, interfaces
5. **Documente trade-offs**: Performance, complexidade, manutenibilidade

**Lembre-se**: Reflection é poderoso, mas tem custos. Use apenas quando realmente necessário e otimize quando possível.

---

**Bons estudos e happy reflecting! 🚀**

