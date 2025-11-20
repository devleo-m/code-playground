# Módulo 36: Memory Management em Profundidade

Bem-vindo ao módulo sobre **Memory Management em Profundidade** em Go! Este módulo ensina como o Go gerencia memória internamente, como otimizar alocações e como escrever código eficiente.

## 📚 Estrutura do Módulo

Este módulo está dividido em 4 aulas principais:

### Aula 1: Memory Management in Depth (Principal)
**Arquivo**: `aula-01-memory-management-principal.md`

Conteúdo completo e detalhado sobre:
- Stack vs Heap: onde variáveis são alocadas
- Garbage Collection: como Go limpa memória automaticamente
- Allocation Patterns: padrões de alocação e seus impactos
- Memory Pooling: técnicas para reduzir alocações
- Otimizações: como minimizar pressão no GC

**Tempo estimado**: 3-4 horas

---

### Aula 2: Versão Simplificada com Analogias
**Arquivo**: `aula-02-memory-management-simplificada.md`

Explicações simplificadas com analogias do dia a dia:
- Stack como "mesa de trabalho"
- Heap como "depósito"
- GC como "faxineiro automático"
- Memory pooling como "biblioteca de empréstimo"
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
- Boas práticas de alocação
- Quando usar sync.Pool
- Otimização de estruturas
- Reduzindo pressão no GC
- Monitoramento e profiling
- Armadilhas comuns
- Checklist de boas práticas

**Tempo estimado**: 1-2 horas

---

## 💻 Exemplos Práticos

**Arquivo**: `01-exemplos.go`

Contém 10 exemplos práticos que demonstram:
1. Stack allocation
2. Heap allocation (escape)
3. Comparação: sem vs com pré-alocação
4. sync.Pool demonstration
5. Reutilização de slices
6. Monitoramento de GC
7. Comparação de performance
8. Escape analysis info
9. Uso correto de sync.Pool
10. Executar todos os exemplos

**Como usar:**
```bash
# Modo interativo
go run 01-exemplos.go

# Executar exemplo específico
go run 01-exemplos.go 1  # Stack vs Heap
go run 01-exemplos.go 2  # Comparação pré-alocação
go run 01-exemplos.go 3  # sync.Pool
go run 01-exemplos.go 8  # Todos os exemplos
```

---

## 🚀 Início Rápido

### 1. Verificar Escape Analysis

```bash
# Ver decisões de escape do compilador
go build -gcflags="-m" main.go

# Mais detalhes
go build -gcflags="-m -m" main.go
```

### 2. Monitorar GC

```go
import (
    "runtime"
    "fmt"
)

var m runtime.MemStats
runtime.ReadMemStats(&m)
fmt.Printf("GC runs: %d\n", m.NumGC)
```

### 3. Usar sync.Pool

```go
var pool = sync.Pool{
    New: func() interface{} {
        return &bytes.Buffer{}
    },
}

func usePool() {
    buf := pool.Get().(*bytes.Buffer)
    defer pool.Put(buf)
    buf.Reset() // Sempre resetar!
    // usar buf...
}
```

### 4. Profiling de Memória

```go
import _ "net/http/pprof"

func main() {
    go func() {
        log.Println(http.ListenAndServe("localhost:6060", nil))
    }()
    // seu código...
}
```

Acesse `http://localhost:6060/debug/pprof/heap` para ver perfil.

---

## 📖 Conceitos Principais

### Stack vs Heap

- **Stack**: Rápido, automático, limitado. Para variáveis locais.
- **Heap**: Mais lento, gerenciado pelo GC, flexível. Para dados compartilhados.

### Garbage Collection

- **Concorrente**: Roda em paralelo com seu código
- **Baixa latência**: Pausas muito curtas (< 1ms)
- **Automático**: Você não precisa gerenciar manualmente

### Allocation Patterns

- **Pré-alocação**: `make([]T, 0, capacity)`
- **Reutilização**: Resetar slices ao invés de criar novos
- **sync.Pool**: Para objetos temporários e caros

### Escape Analysis

O compilador decide onde alocar baseado em:
- Se variável é retornada como pointer
- Se é compartilhada entre goroutines
- Se é muito grande para o stack
- Se precisa persistir além do escopo

---

## 🎯 Objetivos de Aprendizado

Ao final deste módulo, você será capaz de:

- ✅ Entender diferença entre stack e heap
- ✅ Saber como o Go decide onde alocar variáveis
- ✅ Entender como funciona o Garbage Collector
- ✅ Usar `go build -gcflags="-m"` para analisar escape
- ✅ Otimizar alocações em loops e hot paths
- ✅ Usar `sync.Pool` corretamente
- ✅ Monitorar e perfilar uso de memória
- ✅ Identificar problemas de alocação em código
- ✅ Aplicar boas práticas de memory management

---

## 📝 Checklist de Progresso

Marque conforme avança:

- [ ] Li a aula principal (aula-01)
- [ ] Li a aula simplificada (aula-02)
- [ ] Completei os exercícios (aula-03)
- [ ] Li sobre boas práticas (aula-04)
- [ ] Executei os exemplos práticos
- [ ] Usei `go build -gcflags="-m"` para analisar escape
- [ ] Implementei sync.Pool em um exemplo
- [ ] Fiz profiling de memória com pprof
- [ ] Entendi quando otimizar memória faz sentido
- [ ] Apliquei otimizações em código próprio

---

## 🔗 Recursos Adicionais

### Documentação Oficial

- [Go Memory Model](https://go.dev/ref/mem)
- [runtime package](https://pkg.go.dev/runtime)
- [sync.Pool](https://pkg.go.dev/sync#Pool)

### Ferramentas Úteis

- **pprof**: `go tool pprof` - Profiling de memória
- **trace**: `go tool trace` - Análise de execução
- **escape analysis**: `go build -gcflags="-m"` - Ver decisões de escape

### Artigos Recomendados

- "Understanding Go's Memory Allocator" (blog oficial)
- "Go GC: Prioritizing low latency" (blog oficial)
- "Escape Analysis in Go" (vários artigos)

---

## 🐛 Troubleshooting

### Problema: Muitas alocações em loop

**Solução**: Pré-aloque slices com `make([]T, 0, capacity)` ou reutilize slices.

### Problema: GC causando latência

**Solução**: Reduza alocações, use sync.Pool, monitore com pprof.

### Problema: Variável indo para heap quando não deveria

**Solução**: Verifique com `-gcflags="-m"`, evite retornar pointers desnecessários.

### Problema: sync.Pool não está funcionando

**Solução**: Certifique-se de resetar objetos antes de devolver ao pool. GC pode limpar objetos não usados.

---

## 📚 Próximos Módulos

Depois de dominar memory management, você pode avançar para:

- **Módulo 37**: Escape Analysis em Detalhes
- **Módulo 38**: Reflection
- **Módulo 39**: Unsafe Package
- **Módulo 40**: Build Constraints & Tags

---

## 💡 Dicas Finais

1. **Meça antes de otimizar**: Use profiling para identificar problemas reais
2. **Pré-aloque quando possível**: Reduz realocações e melhora performance
3. **Use sync.Pool com cuidado**: Apenas para objetos temporários e caros
4. **Monitore GC**: Entenda o comportamento do GC na sua aplicação
5. **Evite otimização prematura**: Legibilidade primeiro, otimize quando necessário

---

## 📞 Suporte

Se tiver dúvidas ou problemas:

1. Revise a aula simplificada (aula-02)
2. Consulte os exemplos práticos (01-exemplos.go)
3. Use `go build -gcflags="-m"` para analisar escape
4. Consulte a documentação oficial do Go

---

**Bons estudos e happy optimizing! 🚀**



