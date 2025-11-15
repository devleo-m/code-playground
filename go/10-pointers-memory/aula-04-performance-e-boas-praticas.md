# Aula 4: Performance e Boas Práticas - Pointers e Memory Management

Olá! Agora que você entende os conceitos fundamentais de pointers e memory management, é hora de aprender **como usar esses conceitos de forma eficiente** e **quais práticas evitar**. Esta aula vai te ajudar a escrever código Go mais performático e profissional.

---

## 1. Quando Usar Pointers: Diretrizes Práticas

### ✅ Use Pointers Quando:

#### 1.1. Você Precisa Modificar o Valor Original

```go
// ✅ CORRETO: Precisa modificar o valor
func incrementar(contador *int) {
    *contador++
}

// ❌ ERRADO: Não modifica nada
func incrementar(contador int) {
    contador++  // Modifica apenas a cópia
}
```

**Regra de ouro**: Se a função precisa modificar o valor, use pointer.

#### 1.2. A Struct é Grande (Evitar Cópia)

```go
type Configuracao struct {
    // 100 campos aqui...
    DatabaseURL string
    APIKey       string
    // ... muitos outros campos
}

// ✅ CORRETO: Evita copiar struct grande
func processarConfig(config *Configuracao) {
    // Trabalha com o pointer
}

// ❌ INEFICIENTE: Copia toda a struct
func processarConfig(config Configuracao) {
    // Cópia de 100 campos = muito trabalho!
}
```

**Regra prática**: Se a struct tem mais de ~5-10 campos ou contém slices/maps grandes, considere usar pointer.

#### 1.3. Method Receivers que Modificam Estado

```go
type Contador struct {
    valor int
}

// ✅ CORRETO: Modifica o estado
func (c *Contador) Incrementar() {
    c.valor++
}

// ❌ ERRADO: Não modifica nada
func (c Contador) Incrementar() {
    c.valor++  // Modifica apenas a cópia
}
```

**Convenção Go**: Se o método modifica o receiver, use pointer receiver.

#### 1.4. Representar "Opcional" ou "Pode Não Existir"

```go
// ✅ CORRETO: nil representa "não existe"
func encontrarUsuario(id int) *Usuario {
    // Se não encontrar, retorna nil
    return nil
}

// Uso:
usuario := encontrarUsuario(123)
if usuario != nil {
    fmt.Println(usuario.Nome)
}
```

Pointers podem representar valores opcionais (similar a `Option` em Rust ou `Optional` em Java).

---

### ❌ NÃO Use Pointers Quando:

#### 2.1. Tipos Primitivos Pequenos

```go
// ❌ DESNECESSÁRIO: int é pequeno (8 bytes)
func dobrar(valor *int) int {
    return *valor * 2
}

// ✅ MELHOR: Passar por valor
func dobrar(valor int) int {
    return valor * 2
}
```

**Por quê?** Copiar um `int` (8 bytes) é mais rápido que indirecionar um pointer (acessar memória).

#### 2.2. Quando Não Precisa Modificar

```go
// ❌ DESNECESSÁRIO: Não modifica nada
func imprimirNome(pessoa *Pessoa) {
    fmt.Println(pessoa.Nome)
}

// ✅ MELHOR: Passar por valor
func imprimirNome(pessoa Pessoa) {
    fmt.Println(pessoa.Nome)
}
```

**Regra**: Se você não precisa modificar, não use pointer. Código mais simples e seguro.

#### 2.3. Slices e Maps (Na Maioria dos Casos)

```go
// ❌ DESNECESSÁRIO: Slices já são reference types
func processarLista(lista *[]int) {
    // ...
}

// ✅ MELHOR: Passar slice diretamente
func processarLista(lista []int) {
    // Modificações nos elementos já afetam o original
}
```

**Lembre-se**: Slices e maps já são reference types. Só use `*[]int` se precisar **reatribuir** o slice inteiro.

---

## 2. Performance: Entendendo os Custos

### 2.1. Custo de Cópia vs Custo de Indireção

```go
// Custo de copiar uma struct pequena
type Ponto struct {
    X, Y int  // 16 bytes total
}

// Copiar: ~16 bytes
func porValor(p Ponto) { }

// Indireção: 8 bytes (pointer) + acesso à memória
func porReferencia(p *Ponto) { }
```

**Análise**:
- **Struct pequena (< 100 bytes)**: Geralmente mais rápido passar por valor
- **Struct grande (> 100 bytes)**: Geralmente mais rápido passar por pointer
- **Tipos primitivos**: Sempre mais rápido por valor

### 2.2. Cache Locality

```go
// ✅ BOM: Dados próximos na memória (cache-friendly)
pontos := []Ponto{
    {1, 2},
    {3, 4},
    {5, 6},
}

// ❌ PIOR: Pointers espalhados (cache misses)
pontos := []*Ponto{
    &Ponto{1, 2},
    &Ponto{3, 4},
    &Ponto{5, 6},
}
```

**Dica**: Se você precisa iterar sobre muitos elementos, arrays/slices de valores são geralmente mais rápidos que arrays/slices de pointers (melhor cache locality).

### 2.3. Escape Analysis e Alocações

```go
// ✅ BOM: Fica na stack (rápido)
func calcular() int {
    x := 10  // Stack
    y := 20  // Stack
    return x + y
}

// ⚠️ ATENÇÃO: Vai para heap (mais lento)
func criarPointer() *int {
    x := 10  // Heap (escape analysis)
    return &x
}
```

**Regra**: Variáveis que escapam para o heap são mais lentas porque:
1. Alocação no heap é mais lenta
2. Requer garbage collection
3. Pior cache locality

---

## 3. Boas Práticas com Pointers

### 3.1. Sempre Verifique Nil

```go
// ✅ CORRETO: Verifica nil
func processar(dados *[]int) error {
    if dados == nil {
        return fmt.Errorf("dados não podem ser nil")
    }
    // Processa...
    return nil
}

// ❌ PERIGOSO: Pode causar panic
func processar(dados *[]int) {
    for _, v := range *dados {  // PANIC se dados == nil!
        // ...
    }
}
```

**Regra**: Sempre verifique `nil` antes de usar um pointer, especialmente em funções públicas.

### 3.2. Use Zero Values Quando Possível

```go
// ✅ BOM: Usa zero value
type Config struct {
    Timeout time.Duration  // Zero value: 0
}

// ❌ DESNECESSÁRIO: Pointer para zero value
type Config struct {
    Timeout *time.Duration  // nil = não configurado?
}
```

**Dica**: Só use pointer quando `nil` tem significado semântico (opcional, não existe, etc.).

### 3.3. Evite Pointer para Pointer (Quando Possível)

```go
// ❌ CONFUSO: Pointer para pointer
func modificar(ptr **int) {
    **ptr = 10
}

// ✅ MAIS CLARO: Retornar novo valor
func modificar(valor int) int {
    return 10
}
```

**Regra**: Evite `**Type` a menos que seja absolutamente necessário. Geralmente há uma solução mais clara.

### 3.4. Documente Quando Pointers Podem Ser Nil

```go
// ✅ BOM: Documenta comportamento
// EncontrarUsuario busca um usuário pelo ID.
// Retorna nil se o usuário não for encontrado.
func EncontrarUsuario(id int) *Usuario {
    // ...
}

// ✅ MELHOR: Usa error para casos de erro
func EncontrarUsuario(id int) (*Usuario, error) {
    // Retorna erro se não encontrar
}
```

---

## 4. Memory Management: Otimizações

### 4.1. Reduza Alocações Desnecessárias

```go
// ❌ RUIM: Aloca novo slice a cada iteração
func processar(items []int) {
    for i := 0; i < 1000; i++ {
        resultado := []int{}  // Nova alocação!
        resultado = append(resultado, items[i])
    }
}

// ✅ MELHOR: Pré-aloca ou reutiliza
func processar(items []int) {
    resultado := make([]int, 0, len(items))  // Pré-aloca capacidade
    for i := 0; i < len(items); i++ {
        resultado = append(resultado, items[i])
    }
}
```

**Dica**: Use `make` com capacidade inicial quando souber o tamanho aproximado.

### 4.2. Reutilize Slices com `[:0]`

```go
// ✅ BOM: Reutiliza slice
var buffer []byte

func processar() {
    buffer = buffer[:0]  // "Limpa" mas mantém capacidade
    buffer = append(buffer, dados...)
}
```

**Vantagem**: Evita alocações repetidas mantendo a capacidade do slice.

### 4.3. Use `sync.Pool` para Objetos Temporários

```go
import "sync"

var pool = sync.Pool{
    New: func() interface{} {
        return make([]byte, 0, 1024)
    },
}

func processar() {
    buffer := pool.Get().([]byte)
    defer pool.Put(buffer[:0])  // Retorna ao pool
    
    // Usa buffer...
}
```

**Quando usar**: Quando você cria muitos objetos temporários de mesmo tipo em loops críticos.

### 4.4. Evite Memory Leaks com Goroutines

```go
// ❌ PERIGO: Goroutine pode vazar
func iniciar() {
    dados := make([]int, 1000000)
    go func() {
        // Usa dados...
        // Se essa goroutine nunca termina, dados nunca é coletado!
    }()
}

// ✅ MELHOR: Garanta que goroutine termina
func iniciar() {
    dados := make([]int, 1000000)
    done := make(chan bool)
    go func() {
        defer close(done)
        // Usa dados...
    }()
    <-done  // Espera terminar
}
```

**Regra**: Garanta que goroutines terminem, especialmente se capturam variáveis grandes.

---

## 5. Ferramentas para Análise de Performance

### 5.1. Escape Analysis

```bash
# Mostra onde variáveis são alocadas
go build -gcflags="-m" seu_arquivo.go
```

**Exemplo de saída**:
```
./main.go:10:6: can inline exemplo
./main.go:15:6: moved to heap: x
```

### 5.2. Garbage Collection Trace

```bash
# Mostra informações do GC
GODEBUG=gctrace=1 go run seu_arquivo.go
```

**O que observar**:
- Frequência das coletas
- Tempo de pausa
- Quantidade de memória coletada

### 5.3. Memory Profiling

```go
import _ "net/http/pprof"
import "net/http"

func main() {
    go func() {
        log.Println(http.ListenAndServe("localhost:6060", nil))
    }()
    // Seu código...
}
```

Depois use `go tool pprof` para analisar o uso de memória.

### 5.4. Benchmarking

```go
func BenchmarkPorValor(b *testing.B) {
    p := Ponto{X: 1, Y: 2}
    for i := 0; i < b.N; i++ {
        processarPorValor(p)
    }
}

func BenchmarkPorReferencia(b *testing.B) {
    p := &Ponto{X: 1, Y: 2}
    for i := 0; i < b.N; i++ {
        processarPorReferencia(p)
    }
}
```

Execute com `go test -bench=.` para comparar performance.

---

## 6. Anti-padrões Comuns

### 6.1. Pointer para Tudo

```go
// ❌ RUIM: Pointer desnecessário
func calcular(a *int, b *int) *int {
    resultado := *a + *b
    return &resultado
}

// ✅ MELHOR: Valores simples
func calcular(a, b int) int {
    return a + b
}
```

**Problema**: Adiciona complexidade desnecessária e pode causar problemas de performance.

### 6.2. Retornar Pointer de Variável Local Sem Necessidade

```go
// ❌ DESNECESSÁRIO: Retorna pointer de valor pequeno
func criarInt() *int {
    x := 42
    return &x  // Força escape para heap
}

// ✅ MELHOR: Retornar valor
func criarInt() int {
    return 42  // Fica na stack
}
```

**Problema**: Força alocação no heap quando não é necessário.

### 6.3. Modificar Slices Sem Entender Referências

```go
// ⚠️ ATENÇÃO: Pode ter efeitos colaterais inesperados
func adicionarItem(lista []int, item int) {
    lista = append(lista, item)  // Não afeta o original se capacidade esgotar!
}

// ✅ CORRETO: Retornar novo slice
func adicionarItem(lista []int, item int) []int {
    return append(lista, item)
}
```

**Problema**: `append` pode realocar, criando novo slice. O original não é modificado.

---

## 7. Checklist de Boas Práticas

Antes de usar pointers, pergunte-se:

- [ ] **Preciso modificar o valor original?** → Se sim, use pointer
- [ ] **A struct é grande (>100 bytes)?** → Considere pointer para performance
- [ ] **É um tipo primitivo pequeno?** → Evite pointer
- [ ] **É um slice ou map?** → Geralmente não precisa de pointer
- [ ] **O pointer pode ser nil?** → Sempre verifique antes de usar
- [ ] **Estou retornando pointer de variável local?** → Avalie se é necessário
- [ ] **Documentei o comportamento do pointer?** → Especialmente se pode ser nil

---

## 8. Resumo: Regras de Ouro

1. **Use pointers quando precisar modificar valores ou evitar cópias grandes**
2. **Evite pointers para tipos primitivos pequenos**
3. **Slices e maps já são reference types - não precisa de `*` na maioria dos casos**
4. **Sempre verifique `nil` antes de usar pointers**
5. **Prefira retornar valores quando possível, use pointers quando necessário**
6. **Use ferramentas de profiling para medir, não adivinhar**
7. **Documente quando pointers podem ser `nil`**
8. **Reduza alocações desnecessárias para melhor performance**

---

## 9. Próximos Passos

Agora você tem uma base sólida em:
- ✅ Quando e como usar pointers
- ✅ Como Go gerencia memória
- ✅ Boas práticas e otimizações
- ✅ Ferramentas para análise de performance

**Lembre-se**: A otimização prematura é a raiz de todo mal. Primeiro escreva código claro e correto. Depois, se necessário, otimize com base em dados reais (benchmarks, profiling).

Na próxima aula, você continuará expandindo seus conhecimentos em Go. Continue praticando e experimentando!

---

**Dica Final**: Execute `go build -gcflags="-m"` nos seus programas para ver o que o compilador está fazendo. Isso ajuda muito a entender escape analysis e otimizações!

Bons estudos! 🚀

