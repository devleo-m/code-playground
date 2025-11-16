# Módulo 36: Memory Management em Profundidade
## Aula 3 - Exercícios e Reflexão

Agora é hora de colocar a mão na massa! Complete os exercícios abaixo e reflita sobre as questões propostas.

---

## Exercícios Práticos

### Exercício 1: Identificando Stack vs Heap

Analise os seguintes códigos e identifique quais variáveis vão para o **stack** e quais vão para o **heap**. Justifique sua resposta.

#### Código A:
```go
func exemploA() int {
    x := 42
    return x
}
```

#### Código B:
```go
func exemploB() *int {
    x := 42
    return &x
}
```

#### Código C:
```go
var global = 100

func exemploC() *int {
    return &global
}
```

#### Código D:
```go
func exemploD() []int {
    slice := make([]int, 1000)
    return slice
}
```

**Tarefa**: Para cada código, explique:
1. Onde cada variável é alocada (stack ou heap)
2. Por que foi alocada lá
3. Como você verificaria usando `go build -gcflags="-m"`

---

### Exercício 2: Otimizando Alocações

O código abaixo tem problemas de performance relacionados a alocações. Reescreva-o de forma mais eficiente.

#### Código Original (Ineficiente):
```go
package main

import "fmt"

func processData(items []string) []string {
    var result []string
    
    for _, item := range items {
        processed := "Processed: " + item
        result = append(result, processed)
    }
    
    return result
}

func main() {
    items := []string{"item1", "item2", "item3", "item4", "item5"}
    result := processData(items)
    fmt.Println(result)
}
```

**Tarefa**: 
1. Identifique os problemas de alocação
2. Reescreva o código de forma otimizada
3. Explique as melhorias que você fez

**Dicas**:
- Pense em pré-alocação
- Evite concatenação de strings desnecessária
- Considere usar `strings.Builder` para múltiplas concatenações

---

### Exercício 3: Implementando Memory Pool

Crie um sistema de pool de buffers usando `sync.Pool` para processar requisições HTTP simuladas.

**Requisitos**:
1. Crie um pool de `*bytes.Buffer`
2. Implemente uma função `processRequest(data string) string` que:
   - Obtém um buffer do pool
   - Escreve "Response: " + data no buffer
   - Retorna o conteúdo como string
   - Devolve o buffer ao pool (use `defer`)
3. Processe 100 requisições simuladas
4. Compare o número de alocações com e sem pool

**Código base**:
```go
package main

import (
    "bytes"
    "fmt"
    "sync"
)

// TODO: Criar sync.Pool para buffers

// TODO: Implementar processRequest

func main() {
    requests := []string{
        "GET /api/users",
        "POST /api/orders",
        "GET /api/products",
        // ... mais 97 requisições
    }
    
    // Processar com pool
    for _, req := range requests {
        response := processRequest(req)
        fmt.Println(response)
    }
}
```

**Tarefa**: 
1. Complete a implementação
2. Adicione código para medir alocações (use `runtime.ReadMemStats`)
3. Compare com versão sem pool

---

### Exercício 4: Analisando Escape Analysis

Crie um programa que demonstre diferentes cenários de escape analysis.

**Requisitos**:
1. Crie funções que demonstrem:
   - Variável que fica no stack
   - Variável que escapa para o heap
   - Slice que escapa
   - Struct grande que escapa
2. Compile com `go build -gcflags="-m"` e analise a saída
3. Documente suas descobertas

**Código base**:
```go
package main

import "fmt"

// TODO: Criar funções de exemplo

func main() {
    // Chamar suas funções e analisar escape
}
```

**Tarefa**:
1. Implemente pelo menos 5 cenários diferentes
2. Execute `go build -gcflags="-m -m"` (duplo -m para mais detalhes)
3. Crie uma tabela explicando cada caso:
   - Função
   - Variável
   - Onde aloca (stack/heap)
   - Razão do escape (se houver)

---

## Perguntas de Reflexão

### Reflexão 1: Trade-offs do Garbage Collector

O Go usa um Garbage Collector automático, o que significa que você não precisa gerenciar memória manualmente. No entanto, isso tem custos.

**Perguntas para refletir**:
1. **Quais são as vantagens e desvantagens** de ter um GC automático vs gerenciamento manual (como em C/C++)?
2. Em que situações o GC pode se tornar um **gargalo de performance**? Dê exemplos práticos.
3. Por que Go escolheu um GC **concorrente** ao invés de um GC **stop-the-world**? Qual é o trade-off?
4. Se você pudesse controlar quando o GC roda, em que momentos você o executaria? Por quê?

**Escreva suas reflexões** (mínimo 200 palavras):

---

### Reflexão 2: Quando Otimizar Memória?

Nem sempre é necessário otimizar alocações de memória. Às vezes, a otimização prematura pode até piorar o código.

**Perguntas para refletir**:
1. **Quando faz sentido** investir tempo otimizando alocações de memória? Dê critérios objetivos.
2. Como você **identificaria** que um problema de performance é relacionado a memória e não a outro fator (CPU, I/O, etc.)?
3. Qual é o **custo** de otimizações como `sync.Pool`? Quando o overhead pode ser maior que o benefício?
4. Em um projeto real, como você **priorizaria** otimizações de memória vs outras melhorias (algoritmos, cache, etc.)?

**Escreva suas reflexões** (mínimo 200 palavras):

---

### Reflexão 3: Stack vs Heap na Prática

Entender quando algo vai para stack ou heap é crucial para performance, mas nem sempre é óbvio.

**Perguntas para refletir**:
1. Por que o Go **não permite** que você escolha explicitamente onde alocar (stack vs heap)? Quais seriam os problemas se isso fosse permitido?
2. Em que situações você **quer** que algo escape para o heap, mesmo sabendo que é mais lento? Dê exemplos.
3. Como a decisão stack vs heap afeta a **segurança** e **corretude** do programa? (Pense em pointers inválidos)
4. Se você descobrisse que uma variável está indo para o heap quando não deveria, **como você a otimizaria**? Quais técnicas usaria?

**Escreva suas reflexões** (mínimo 200 palavras):

---

## Checklist de Aprendizado

Marque conforme você completa:

- [ ] Entendi a diferença entre stack e heap
- [ ] Sei como o Go decide onde alocar variáveis
- [ ] Entendo como funciona o Garbage Collector do Go
- [ ] Sei usar `go build -gcflags="-m"` para analisar escape
- [ ] Entendo quando usar `sync.Pool`
- [ ] Sei otimizar alocações em loops
- [ ] Entendo os trade-offs de pointer vs value
- [ ] Sei medir e monitorar uso de memória
- [ ] Entendo quando otimizar memória faz sentido
- [ ] Posso identificar problemas de alocação em código

---

## Desafio Extra (Opcional)

Crie um benchmark comparando:

1. **Versão sem otimização**: Alocação repetida em loop
2. **Versão com pré-alocação**: Slice pré-alocado
3. **Versão com sync.Pool**: Reutilização de buffers

**Requisitos**:
- Use `go test -bench=. -benchmem` para medir
- Compare: tempo de execução, alocações, bytes alocados
- Documente os resultados e explique as diferenças

**Código base**:
```go
package main

import "testing"

func BenchmarkSemOtimizacao(b *testing.B) {
    // TODO
}

func BenchmarkPreAlocacao(b *testing.B) {
    // TODO
}

func BenchmarkComPool(b *testing.B) {
    // TODO
}
```

---

## Dicas para os Exercícios

1. **Exercício 1**: Use `go build -gcflags="-m"` para verificar suas hipóteses
2. **Exercício 2**: Use `strings.Builder` para concatenação eficiente
3. **Exercício 3**: Não esqueça de resetar o buffer antes de devolver ao pool
4. **Exercício 4**: O flag `-m -m` (duplo) mostra mais detalhes sobre escape

---

**Boa sorte com os exercícios! Lembre-se: a prática é essencial para dominar gerenciamento de memória.** 🚀

