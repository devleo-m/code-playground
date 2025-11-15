# Módulo 10: Pointers e Memory Management em Go

## Aula 1: Pointers e Memory Management - Controle de Memória e Eficiência

Olá! Bem-vindo ao décimo módulo. Até agora você aprendeu variáveis, tipos de dados, structs, funções e como trabalhar com slices e maps. Mas você já se perguntou **onde** essas variáveis são armazenadas na memória? E como Go gerencia toda essa memória automaticamente?

Nesta aula, vamos mergulhar profundamente em **pointers** (ponteiros) e **memory management** (gerenciamento de memória) em Go. Esses conceitos são fundamentais para entender como Go funciona internamente e como escrever código mais eficiente.

---

## 1. O Que São Pointers?

**Pointers** (ponteiros) são variáveis que armazenam **endereços de memória** de outras variáveis, ao invés de armazenar os valores diretamente. Eles permitem que você trabalhe com a localização exata de um dado na memória do computador.

### Por Que Precisamos de Pointers?

Pointers são essenciais para:

- **Eficiência de memória**: Evitar copiar grandes estruturas de dados
- **Modificar valores**: Permitir que funções modifiquem variáveis do chamador
- **Performance**: Trabalhar diretamente com endereços de memória
- **Estruturas de dados**: Construir listas, árvores e outras estruturas complexas

### Sintaxe Básica de Pointers

Em Go, você declara um pointer usando `*Type` e obtém o endereço de uma variável com `&`:

```go
var x int = 42
var p *int = &x  // p é um pointer para int, armazena o endereço de x

fmt.Printf("Valor de x: %d\n", x)           // 42
fmt.Printf("Endereço de x: %p\n", &x)        // 0xc0000140a0 (exemplo)
fmt.Printf("Valor de p (endereço): %p\n", p) // 0xc0000140a0
fmt.Printf("Valor apontado por p: %d\n", *p) // 42 (dereferencing)
```

**Operadores importantes:**
- `&` (address-of): Obtém o endereço de uma variável
- `*` (dereference): Acessa o valor apontado por um pointer
- `*Type`: Declara um pointer para um tipo específico

### Zero Value de Pointers

O zero value de um pointer é `nil`:

```go
var p *int
fmt.Println(p == nil)  // true
fmt.Printf("%p\n", p)   // 0x0

// Tentar usar *p causaria panic!
// *p = 10  // ERRO: panic: runtime error: invalid memory address
```

**Importante**: Sempre verifique se um pointer não é `nil` antes de usá-lo!

---

## 2. Passagem por Valor vs Passagem por Referência

### Passagem por Valor (Sem Pointer)

Quando você passa uma variável para uma função sem pointer, Go **copia** o valor:

```go
func main() {
    x := 10
    fmt.Printf("Antes: x = %d\n", x)  // 10
    naoModifica(x)
    fmt.Printf("Depois: x = %d\n", x)  // 10 (não mudou!)
}

func naoModifica(valor int) {
    valor = 20  // Modifica apenas a cópia local
}
```

A função recebe uma **cópia** da variável, então qualquer modificação não afeta o original.

### Passagem por Referência (Com Pointer)

Quando você passa um pointer, a função trabalha com o **mesmo endereço de memória**:

```go
func main() {
    x := 10
    fmt.Printf("Antes: x = %d\n", x)  // 10
    modifica(&x)  // Passa o endereço
    fmt.Printf("Depois: x = %d\n", x)  // 20 (mudou!)
}

func modifica(ptr *int) {
    *ptr = 20  // Modifica o valor original
}
```

Agora a função pode modificar o valor original porque trabalha diretamente com o endereço de memória.

### Quando Usar Cada Um?

- **Passagem por valor**: Use quando você **não precisa** modificar o valor original (mais seguro, evita efeitos colaterais)
- **Passagem por referência**: Use quando você **precisa** modificar o valor original ou quando a estrutura é muito grande (evita cópia)

---

## 3. Pointers com Structs

Pointers são especialmente úteis com structs, pois structs podem ser grandes e copiá-las pode ser custoso.

### Passando Struct por Valor

```go
type Pessoa struct {
    Nome  string
    Idade int
}

func main() {
    p := Pessoa{Nome: "João", Idade: 30}
    modificaStructPorValor(p)
    fmt.Printf("%+v\n", p)  // {Nome:João Idade:30} (não mudou)
}

func modificaStructPorValor(p Pessoa) {
    p.Idade = 99  // Modifica apenas a cópia
}
```

### Passando Struct por Referência

```go
func main() {
    p := Pessoa{Nome: "João", Idade: 30}
    modificaStructPorReferencia(&p)
    fmt.Printf("%+v\n", p)  // {Nome:João Idade:99} (mudou!)
}

func modificaStructPorReferencia(p *Pessoa) {
    p.Idade = 99  // Modifica o original
}
```

### Shorthand de Go: Acesso Automático

Go tem uma conveniência: você pode acessar campos de uma struct através de um pointer **sem precisar usar parênteses**:

```go
p := &Pessoa{Nome: "Maria", Idade: 25}

// Ambas as formas são equivalentes:
fmt.Println((*p).Nome)  // Forma explícita
fmt.Println(p.Nome)     // Forma simplificada (Go faz o dereference automaticamente)

// Modificando também funciona:
p.Idade = 26  // Go automaticamente faz (*p).Idade = 26
```

Isso torna o código mais limpo e legível!

### Method Receivers com Pointers

Quando você define métodos para structs, pode escolher entre receiver por valor ou por pointer:

```go
type Pessoa struct {
    Nome  string
    Idade int
}

// Método com receiver por valor (não modifica)
func (p Pessoa) AniversarioPorValor() {
    p.Idade++  // Não modifica o original
}

// Método com receiver por pointer (modifica)
func (p *Pessoa) AniversarioPorReferencia() {
    p.Idade++  // Modifica o original
}

func main() {
    p := Pessoa{Nome: "Pedro", Idade: 20}
    
    p.AniversarioPorValor()
    fmt.Println(p.Idade)  // 20 (não mudou)
    
    p.AniversarioPorReferencia()
    fmt.Println(p.Idade)  // 21 (mudou!)
}
```

**Regra geral**: Use pointer receiver quando o método precisa modificar a struct ou quando a struct é grande (evita cópia).

---

## 4. Pointers com Slices e Maps

Aqui está um ponto **muito importante** sobre Go: **slices e maps são reference types**!

### Slices são Reference Types

Quando você passa um slice para uma função, você está passando uma **referência** ao array subjacente, não uma cópia:

```go
func main() {
    slice := []int{1, 2, 3}
    fmt.Printf("Antes: %v\n", slice)  // [1 2 3]
    
    modificaSlice(slice)
    fmt.Printf("Depois: %v\n", slice)  // [999 2 3] (modificado!)
}

func modificaSlice(s []int) {
    s[0] = 999  // Modifica o slice original
}
```

**Por quê?** Um slice é na verdade uma estrutura que contém:
- Um pointer para o array subjacente
- Um comprimento (length)
- Uma capacidade (capacity)

Quando você passa um slice, Go copia essa estrutura (que é pequena), mas o pointer dentro dela aponta para o **mesmo array**. Por isso, modificações nos elementos afetam o original.

**Mas atenção**: Reatribuir o slice dentro da função **não** afeta o original:

```go
func main() {
    slice := []int{1, 2, 3}
    reatribuiSlice(slice)
    fmt.Printf("Depois: %v\n", slice)  // [1 2 3] (não mudou!)
}

func reatribuiSlice(s []int) {
    s = []int{10, 20, 30}  // Reatribui apenas a cópia local
}
```

Se você **realmente** quiser reatribuir o slice original, precisa passar um pointer para o slice:

```go
func main() {
    slice := &[]int{1, 2, 3}
    reatribuiSlice(slice)
    fmt.Printf("Depois: %v\n", *slice)  // [10 20 30] (mudou!)
}

func reatribuiSlice(s *[]int) {
    *s = []int{10, 20, 30}  // Reatribui o slice original
}
```

### Maps são Reference Types

Maps funcionam da mesma forma:

```go
func main() {
    m := map[string]int{"a": 1, "b": 2}
    fmt.Printf("Antes: %v\n", m)  // map[a:1 b:2]
    
    modificaMap(m)
    fmt.Printf("Depois: %v\n", m)  // map[a:999 b:2 c:3] (modificado!)
}

func modificaMap(m map[string]int) {
    m["c"] = 3   // Adiciona novo elemento
    m["a"] = 999 // Modifica elemento existente
}
```

**Resumo**: Para slices e maps, você geralmente **não precisa** de pointers explícitos para modificar elementos, mas precisa se quiser reatribuir a variável inteira.

---

## 5. Memory Management em Go

Go gerencia memória de forma **automática** através de um processo chamado **garbage collection** (coleta de lixo). Mas entender como isso funciona ajuda a escrever código mais eficiente.

### Stack vs Heap

A memória do programa é dividida em duas áreas principais:

1. **Stack (Pilha)**: 
   - Memória rápida e organizada
   - Usada para variáveis locais de funções
   - Limpeza automática quando a função termina
   - Tamanho limitado

2. **Heap (Monte)**:
   - Memória mais lenta mas maior
   - Usada para dados que precisam viver além da função
   - Requer garbage collection para limpeza
   - Tamanho muito maior

### Escape Analysis

O compilador Go usa **escape analysis** para decidir onde alocar cada variável:

```go
func criaIntNoStack() int {
    valor := 42  // Provavelmente vai para stack
    return valor
}

func criaIntNoHeap() *int {
    valor := 42  // Escape analysis detecta que precisa viver além da função
    return &valor  // Retornar pointer força escape para heap
}
```

Quando você retorna um pointer de uma função, Go **deve** mover a variável para o heap, porque a stack será limpa quando a função terminar.

### Você Não Precisa Gerenciar Manualmente!

Diferente de linguagens como C ou C++, você **não precisa**:
- Alocar memória manualmente (`malloc`)
- Liberar memória manualmente (`free`)
- Se preocupar com memory leaks (na maioria dos casos)

O garbage collector do Go cuida disso automaticamente!

---

## 6. Garbage Collection (GC)

O **Garbage Collector** do Go é responsável por liberar memória que não está mais sendo usada.

### Como Funciona?

O GC do Go usa um algoritmo chamado **tri-color mark-and-sweep**:

1. **Mark (Marcação)**: Identifica toda memória que ainda está em uso
2. **Sweep (Varredura)**: Libera memória que não está mais em uso

### Características do GC do Go

- **Concorrente**: Roda em paralelo com seu programa (não pausa tudo)
- **Baixa latência**: Pausas muito curtas (geralmente < 1ms)
- **Automático**: Você não precisa fazer nada
- **Eficiente**: Otimizado para programas Go

### Quando o GC Roda?

O GC roda automaticamente quando:
- A memória heap atinge um certo limite
- O runtime detecta que há memória suficiente para coletar

Você pode ver estatísticas do GC com:

```bash
go run -gcflags="-m" seu_arquivo.go  # Mostra escape analysis
GODEBUG=gctrace=1 go run seu_arquivo.go  # Mostra informações do GC
```

### Boas Práticas para GC

1. **Evite alocações desnecessárias**: Reutilize slices quando possível
2. **Use sync.Pool** para objetos temporários frequentes
3. **Evite criar muitos pequenos objetos** em loops críticos
4. **Prefira passar por valor** para tipos pequenos (int, bool, etc.)

---

## 7. Resumo e Próximos Passos

Nesta aula, você aprendeu:

✅ **Pointers**: Variáveis que armazenam endereços de memória  
✅ **Operadores**: `&` para obter endereço, `*` para dereferenciar  
✅ **Passagem por valor vs referência**: Quando usar cada um  
✅ **Pointers com structs**: Eficiência e modificação de valores  
✅ **Slices e maps**: Reference types que não precisam de pointers explícitos  
✅ **Memory management**: Stack vs heap e escape analysis  
✅ **Garbage collection**: Como Go gerencia memória automaticamente  

**Pontos importantes para lembrar:**

- Pointers permitem eficiência e modificação de valores
- Slices e maps já são reference types (não precisa de `*` na maioria dos casos)
- Go gerencia memória automaticamente (você não precisa se preocupar)
- Use pointers quando precisar modificar valores ou evitar cópias grandes

Na próxima aula, vamos simplificar esses conceitos com analogias do dia a dia para fixar o aprendizado!

---

**Dica**: Execute o arquivo `01-exemplos.go` para ver todos esses conceitos em ação. Experimente modificar os exemplos e veja o que acontece!

