# Módulo 38: Reflection em Go

Bem-vindo ao módulo sobre **Reflection** em Go! Este módulo ensina como usar o package `reflect` para inspecionar e manipular tipos e valores em tempo de execução.

## 📚 Estrutura do Módulo

Este módulo está dividido em 4 aulas principais:

### Aula 1: Reflection (Principal)
**Arquivo**: `aula-01-reflection-principal.md`

Conteúdo completo e detalhado sobre:
- O que é Reflection e por que existe
- Package reflect: Type e Value
- Type Reflection: Inspecionar tipos
- Value Reflection: Inspecionar e modificar valores
- Chamadas dinâmicas de métodos
- Casos de uso práticos
- Limitações e cuidados

**Tempo estimado**: 3-4 horas

---

### Aula 2: Versão Simplificada com Analogias
**Arquivo**: `aula-02-reflection-simplificada.md`

Explicações simplificadas com analogias do dia a dia:
- Reflection como "espelho mágico"
- Type vs Value: "o que é" vs "o que tem"
- Inspecionar tipos como "detetive investigando"
- Modificar valores como "abrir caixa com chave"
- Conceitos visuais e fáceis de entender

**Tempo estimado**: 1-2 horas

---

### Aula 3: Exercícios e Reflexão
**Arquivo**: `aula-03-exercicios-e-reflexao.md`

Exercícios práticos para fixar o aprendizado:
- 4 exercícios práticos progressivos
- 3 questões para reflexão profunda
- Desafios avançados
- Checklist de aprendizado

**Tempo estimado**: 2-3 horas

---

### Aula 4: Performance e Boas Práticas
**Arquivo**: `aula-04-performance-e-boas-praticas.md`

Otimizações e melhores práticas:
- Cache de Types e Values
- Evitar reflection em hot paths
- Padrões de otimização
- Armadilhas comuns
- Alternativas ao reflection
- Checklist de boas práticas

**Tempo estimado**: 1-2 horas

---

## 💻 Exemplos Práticos

**Arquivo**: `01-exemplos.go`

Contém 10 exemplos práticos que demonstram:
1. Inspecionar tipo básico
2. Inspecionar struct
3. Ler tags
4. Modificar valor
5. Modificar struct
6. Chamar método dinamicamente
7. Criar novo valor
8. Validador simples
9. JSON simples
10. Comparação genérica

**Como usar:**
```bash
# Modo interativo
go run 01-exemplos.go

# Executar exemplo específico
go run 01-exemplos.go 1  # Inspecionar tipo
go run 01-exemplos.go 2  # Inspecionar struct
go run 01-exemplos.go 8  # Validador
go run 01-exemplos.go 9  # JSON
```

---

## 🚀 Início Rápido

### 1. Inspecionar Tipo

```go
import "reflect"

var x int = 42
t := reflect.TypeOf(x)
fmt.Println("Type:", t)  // int
```

### 2. Inspecionar Valor

```go
v := reflect.ValueOf(x)
fmt.Println("Value:", v.Int())  // 42
```

### 3. Modificar Valor

```go
x := 42
v := reflect.ValueOf(&x).Elem()
v.SetInt(100)
fmt.Println(x)  // 100
```

### 4. Inspecionar Struct

```go
type Person struct {
    Name string
    Age  int
}

p := Person{Name: "John", Age: 30}
t := reflect.TypeOf(p)

for i := 0; i < t.NumField(); i++ {
    field := t.Field(i)
    fmt.Println("Field:", field.Name)
}
```

---

## 📖 Conceitos Principais

### Type vs Value

- **Type**: Representa um tipo Go (`reflect.Type`)
- **Value**: Representa um valor Go (`reflect.Value`)

### Type Reflection

- Inspecionar tipos básicos
- Inspecionar structs e campos
- Ler tags de structs
- Inspecionar slices, maps, functions

### Value Reflection

- Obter valores
- Modificar valores (precisa de pointer)
- Criar novos valores
- Chamar métodos dinamicamente

### Casos de Uso

- JSON/XML marshaling
- Validação genérica
- ORM frameworks
- Code generation tools

---

## 🎯 Objetivos de Aprendizado

Ao final deste módulo, você será capaz de:

- ✅ Entender o que é reflection
- ✅ Saber a diferença entre Type e Value
- ✅ Inspecionar tipos e valores
- ✅ Modificar valores usando reflection
- ✅ Chamar métodos dinamicamente
- ✅ Criar validadores genéricos
- ✅ Entender limitações de performance
- ✅ Saber quando usar reflection
- ✅ Aplicar boas práticas

---

## 📝 Checklist de Progresso

Marque conforme avança:

- [ ] Li a aula principal (aula-01)
- [ ] Li a aula simplificada (aula-02)
- [ ] Completei os exercícios (aula-03)
- [ ] Li sobre boas práticas (aula-04)
- [ ] Executei os exemplos práticos
- [ ] Sei inspecionar tipos básicos
- [ ] Sei inspecionar structs
- [ ] Sei modificar valores
- [ ] Sei chamar métodos dinamicamente
- [ ] Criei um validador genérico

---

## 🔗 Recursos Adicionais

### Documentação Oficial

- [reflect package](https://pkg.go.dev/reflect)
- [The Laws of Reflection](https://go.dev/blog/laws-of-reflection)

### Artigos Recomendados

- [Go Reflection Examples](https://golang.org/pkg/reflect/#pkg-examples)
- "Understanding Go Reflection" (vários artigos)

---

## 🐛 Troubleshooting

### Problema: Cannot set value

**Solução**: Passe um pointer e use `Elem()`:
```go
v := reflect.ValueOf(&x).Elem()
v.SetInt(100)
```

### Problema: Method not found

**Solução**: Sempre verifique com `IsValid()`:
```go
method := v.MethodByName("MethodName")
if !method.IsValid() {
    return fmt.Errorf("method not found")
}
```

### Problema: Performance lenta

**Solução**: Cache types, evite em hot paths, use type assertions quando possível.

---

## 📚 Próximos Módulos

Depois de dominar reflection, você pode avançar para:

- **Módulo 39**: Unsafe Package
- **Módulo 40**: Build Constraints & Tags
- **Módulo 41**: Compiler & Linker Flags
- **Módulo 42**: CGO Basics

---

## 💡 Dicas Finais

1. **Use com moderação**: Reflection é poderoso, mas tem custos
2. **Cache quando possível**: Types, fields, validators
3. **Evite em hot paths**: Use tipos concretos ou generics
4. **Valide sempre**: Verifique IsValid, CanSet antes de usar
5. **Considere alternativas**: Generics, code generation, interfaces

---

## 📞 Suporte

Se tiver dúvidas ou problemas:

1. Revise a aula simplificada (aula-02)
2. Consulte os exemplos práticos (01-exemplos.go)
3. Consulte a documentação oficial do reflect package
4. Leia "The Laws of Reflection"

---

**Bons estudos e happy reflecting! 🚀**



