# Módulo 9: Functions em Go

## Aula 1: Functions - Blocos de Código Reutilizáveis

Olá! Bem-vindo ao nono módulo. Até agora você aprendeu variáveis, tipos de dados, estruturas de controle e loops. Mas e se você precisar executar o mesmo código em **diferentes lugares**? E se precisar **organizar** seu código em partes lógicas e reutilizáveis?

É aqui que entram as **functions** (funções) - blocos de código nomeados que podem ser chamados de qualquer lugar do seu programa. Em Go, funções são **first-class citizens**, o que significa que podem ser tratadas como qualquer outro valor: atribuídas a variáveis, passadas como argumentos, retornadas de outras funções.

Nesta aula, vamos mergulhar profundamente em todas as formas de usar funções em Go, desde o básico até conceitos avançados como closures e passagem por valor.

---

## 1. O Que São Functions?

**Functions** (funções) são blocos de código nomeados que executam uma tarefa específica. Elas são essenciais para:

- **Reutilização de código**: Escreva uma vez, use muitas vezes
- **Organização**: Divida código complexo em partes menores e gerenciáveis
- **Modularidade**: Crie componentes independentes e testáveis
- **Abstração**: Esconda detalhes de implementação atrás de uma interface simples

### Por Que Precisamos de Functions?

**Sem funções**, você teria que repetir código:

```go
func main() {
    // Calcular área de retângulo 1
    largura1 := 5.0
    altura1 := 3.0
    area1 := largura1 * altura1
    fmt.Printf("Área 1: %.2f\n", area1)

    // Calcular área de retângulo 2
    largura2 := 7.0
    altura2 := 4.0
    area2 := largura2 * altura2
    fmt.Printf("Área 2: %.2f\n", area2)

    // Muito código repetido!
}
```

**Com funções**, você pode reutilizar:

```go
func calcularArea(largura, altura float64) float64 {
    return largura * altura
}

func main() {
    area1 := calcularArea(5.0, 3.0)
    area2 := calcularArea(7.0, 4.0)
    fmt.Printf("Área 1: %.2f\n", area1)
    fmt.Printf("Área 2: %.2f\n", area2)
}
```

**Vantagens:**

- ✅ Código mais limpo e organizado
- ✅ Fácil manutenção - mude uma vez, afeta todos os usos
- ✅ Testável - teste funções isoladamente
- ✅ Reutilizável - use em múltiplos lugares

---

## 2. Function Basics - Fundamentos

### Sintaxe Básica

A sintaxe para declarar uma função em Go é:

```go
func nomeDaFuncao(parametros) tipoRetorno {
    // corpo da função
    return valor
}
```

### Função Simples

```go
func cumprimentar() {
    fmt.Println("Olá, mundo!")
}
```

**Características:**

- `func`: Palavra-chave para declarar função
- `cumprimentar`: Nome da função
- `()`: Sem parâmetros
- Sem retorno (não especifica tipo de retorno)

### Função com Parâmetros

```go
func cumprimentarPessoa(nome string) {
    fmt.Printf("Olá, %s!\n", nome)
}
```

**Parâmetros:**

- `nome string`: Parâmetro chamado `nome` do tipo `string`
- Parâmetros são passados **por valor** (cópia)

### Função com Retorno

```go
func somar(a int, b int) int {
    return a + b
}
```

**Retorno:**

- `int`: Tipo do valor retornado
- `return`: Palavra-chave para retornar valor
- Função **deve** retornar valor do tipo especificado

### Função com Múltiplos Parâmetros do Mesmo Tipo

```go
func multiplicar(a, b, c int) int {
    return a * b * c
}
```

**Sintaxe simplificada:**

- `a, b, c int`: Todos são do tipo `int`
- Equivale a: `a int, b int, c int`

### Função com Múltiplos Parâmetros de Tipos Diferentes

```go
func calcularIMC(peso float64, altura float64) float64 {
    return peso / (altura * altura)
}
```

---

## 3. First-Class Citizens - Funções como Valores

Em Go, funções são **first-class citizens**, o que significa que podem ser:

### Atribuídas a Variáveis

```go
var somar func(int, int) int
somar = func(a, b int) int {
    return a + b
}

resultado := somar(5, 3)
```

### Passadas como Argumentos

```go
func aplicarOperacao(a, b int, operacao func(int, int) int) int {
    return operacao(a, b)
}

func main() {
    resultado := aplicarOperacao(10, 5, somar)
    fmt.Println(resultado) // 15
}
```

### Retornadas de Outras Funções

```go
func criarOperacao(tipo string) func(int, int) int {
    switch tipo {
    case "soma":
        return func(a, b int) int { return a + b }
    case "multiplicacao":
        return func(a, b int) int { return a * b }
    default:
        return func(a, b int) int { return 0 }
    }
}
```

**Por que isso é útil?**

- Permite criar código mais flexível e reutilizável
- Habilita padrões de programação funcional
- Facilita criação de callbacks e handlers

---

## 4. Variadic Functions - Funções Variádicas

**Variadic functions** são funções que aceitam um **número variável de argumentos** do mesmo tipo.

### Sintaxe

```go
func nomeFuncao(args ...Tipo) TipoRetorno {
    // args é tratado como um slice do tipo Tipo
}
```

### Exemplo Básico

```go
func somarNumeros(numeros ...int) int {
    soma := 0
    for _, num := range numeros {
        soma += num
    }
    return soma
}

func main() {
    resultado1 := somarNumeros(1, 2, 3)           // 6
    resultado2 := somarNumeros(10, 20, 30, 40)    // 100
    resultado3 := somarNumeros()                   // 0 (nenhum argumento)
}
```

**Características:**

- `...int`: Indica que aceita zero ou mais `int`
- Dentro da função, `numeros` é um `[]int` (slice)
- Pode ser chamada com qualquer número de argumentos

### Variadic com Outros Parâmetros

O parâmetro variádico **deve ser o último**:

```go
func juntarStrings(separador string, textos ...string) string {
    return strings.Join(textos, separador)
}

func main() {
    resultado := juntarStrings(", ", "a", "b", "c")
    // "a, b, c"
}
```

### Passar Slice para Função Variádica

Use `...` para "desempacotar" um slice:

```go
numeros := []int{1, 2, 3, 4, 5}
soma := somarNumeros(numeros...)  // Equivale a somarNumeros(1, 2, 3, 4, 5)
```

**Funções padrão que usam variadic:**

- `fmt.Printf(format string, args ...interface{})`
- `append(slice []T, elems ...T) []T`

---

## 5. Multiple Return Values - Múltiplos Valores de Retorno

Go permite que funções retornem **múltiplos valores**. Isso é especialmente útil para o padrão idiomático de retornar resultado e erro.

### Sintaxe

```go
func nomeFuncao() (Tipo1, Tipo2) {
    return valor1, valor2
}
```

### Exemplo: Resultado e Erro

```go
func dividir(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("divisão por zero")
    }
    return a / b, nil
}

func main() {
    resultado, err := dividir(10, 2)
    if err != nil {
        fmt.Printf("Erro: %v\n", err)
        return
    }
    fmt.Printf("Resultado: %.2f\n", resultado)
}
```

**Padrão idiomático em Go:**

- Sempre retorne `(resultado, error)`
- Se sucesso: `return valor, nil`
- Se erro: `return zeroValue, erro`

### Múltiplos Valores Úteis

```go
func calcularEstatisticas(numeros []float64) (float64, float64, float64) {
    // Retorna: média, máximo, mínimo
    if len(numeros) == 0 {
        return 0, 0, 0
    }

    soma := 0.0
    max := numeros[0]
    min := numeros[0]

    for _, num := range numeros {
        soma += num
        if num > max {
            max = num
        }
        if num < min {
            min = num
        }
    }

    media := soma / float64(len(numeros))
    return media, max, min
}
```

### Ignorar Valores de Retorno

Use `_` (blank identifier) para ignorar valores:

```go
// Ignorar erro
resultado, _ := dividir(10, 2)

// Ignorar todos os valores (raro)
_, _, _ = calcularEstatisticas([]float64{1, 2, 3})
```

---

## 6. Named Return Values - Retornos Nomeados

Go permite **nomear** os valores de retorno. Eles são tratados como variáveis dentro da função e são inicializados com seus valores zero.

### Sintaxe

```go
func nomeFuncao() (nome1 Tipo1, nome2 Tipo2) {
    // nome1 e nome2 são variáveis disponíveis
    nome1 = valor1
    nome2 = valor2
    return // Retorna nome1 e nome2 implicitamente
}
```

### Exemplo

```go
func dividirComNome(dividendo, divisor float64) (quociente float64, resto float64, err error) {
    if divisor == 0 {
        err = fmt.Errorf("divisão por zero")
        return // Retorna quociente=0, resto=0, err=erro
    }
    quociente = dividendo / divisor
    resto = math.Mod(dividendo, divisor)
    return // Retorna quociente, resto, err=nil
}
```

**Características:**

- Variáveis nomeadas são inicializadas com valores zero
- `return` sem argumentos retorna os valores atuais das variáveis nomeadas
- Melhora legibilidade em alguns casos
- Facilita refatoração

### Quando Usar Named Returns?

**✅ Use quando:**

- Melhora significativamente a legibilidade
- Retorna múltiplos valores relacionados
- Facilita documentação

**⚠️ Evite quando:**

- Função simples com um único retorno
- Pode confundir em funções curtas
- Pode mascarar erros (esquecer de inicializar)

---

## 7. Anonymous Functions - Funções Anônimas

**Anonymous functions** (funções anônimas) são funções declaradas **sem nome**. Também chamadas de **function literals** ou **lambdas**.

### Sintaxe

```go
func(parametros) tipoRetorno {
    // corpo
}
```

### Atribuir a Variável

```go
somar := func(a, b int) int {
    return a + b
}

resultado := somar(5, 3)
```

### Executar Imediatamente (IIFE)

```go
func() {
    fmt.Println("Executado imediatamente!")
}()

// Com parâmetros
func(nome string) {
    fmt.Printf("Olá, %s!\n", nome)
}("João")
```

### Passar como Argumento

```go
func processarNumeros(numeros []int, transformacao func(int) int) {
    for i, num := range numeros {
        numeros[i] = transformacao(num)
    }
}

func main() {
    numeros := []int{1, 2, 3, 4, 5}
    processarNumeros(numeros, func(n int) int {
        return n * 2
    })
    // numeros agora: [2, 4, 6, 8, 10]
}
```

**Casos de uso:**

- Callbacks e event handlers
- Transformações inline
- Goroutines (veremos depois)
- Programação funcional

---

## 8. Closures - Fechamentos

**Closures** são funções que **capturam variáveis** do escopo circundante. Elas "fecham sobre" (close over) variáveis externas, permitindo acesso mesmo após a função externa retornar.

### Exemplo Básico

```go
func criarContador() func() int {
    contador := 0
    return func() int {
        contador++  // Captura e modifica variável externa
        return contador
    }
}

func main() {
    contador := criarContador()
    fmt.Println(contador()) // 1
    fmt.Println(contador()) // 2
    fmt.Println(contador()) // 3
}
```

**O que acontece:**

1. `criarContador()` cria variável `contador`
2. Retorna função anônima que acessa `contador`
3. Função retornada "fecha sobre" `contador`
4. Cada chamada mantém estado entre execuções

### Closure com Parâmetros

```go
func criarMultiplicador(fator int) func(int) int {
    return func(numero int) int {
        return numero * fator  // Captura fator
    }
}

func main() {
    dobrar := criarMultiplicador(2)
    triplicar := criarMultiplicador(3)

    fmt.Println(dobrar(5))      // 10
    fmt.Println(triplicar(5))   // 15
}
```

### Closure com Estado Compartilhado

```go
func criarAcumulador() func(int) int {
    soma := 0
    return func(valor int) int {
        soma += valor
        return soma
    }
}

func main() {
    acumulador := criarAcumulador()
    fmt.Println(acumulador(10)) // 10
    fmt.Println(acumulador(20)) // 30
    fmt.Println(acumulador(30)) // 60
}
```

### ⚠️ Armadilha: Closure em Loops

**PROBLEMA:**

```go
var funcoes []func() int
for i := 0; i < 3; i++ {
    funcoes = append(funcoes, func() int {
        return i  // Todas capturam a MESMA variável i!
    })
}

// Todas retornam 3 (valor final de i)
for _, f := range funcoes {
    fmt.Println(f()) // 3, 3, 3
}
```

**SOLUÇÃO 1: Criar nova variável**

```go
for i := 0; i < 3; i++ {
    i := i  // Nova variável em cada iteração
    funcoes = append(funcoes, func() int {
        return i
    })
}
```

**SOLUÇÃO 2: Passar como parâmetro**

```go
for i := 0; i < 3; i++ {
    funcoes = append(funcoes, func(valor int) func() int {
        return func() int { return valor }
    }(i))
}
```

**Casos de uso de closures:**

- Manter estado entre chamadas
- Criar funções especializadas
- Callbacks e event handlers
- Iteradores e geradores
- Programação funcional

---

## 9. Call by Value - Passagem por Valor

Em Go, **tudo é passado por valor**. Isso significa que quando você passa um argumento para uma função, Go cria uma **cópia** do valor.

### Tipos Primitivos

```go
func modificarInt(numero int) {
    numero = 999  // Modifica apenas a cópia
}

func main() {
    num := 42
    modificarInt(num)
    fmt.Println(num)  // 42 (não mudou!)
}
```

### Structs

```go
type Pessoa struct {
    Nome  string
    Idade int
}

func modificarPessoa(p Pessoa) {
    p.Nome = "Modificado"  // Modifica apenas a cópia
}

func main() {
    pessoa := Pessoa{Nome: "João", Idade: 30}
    modificarPessoa(pessoa)
    fmt.Println(pessoa)  // {João 30} (não mudou!)
}
```

### Arrays

```go
func modificarArray(arr [5]int) {
    arr[0] = 999  // Modifica apenas a cópia
}

func main() {
    arr := [5]int{1, 2, 3, 4, 5}
    modificarArray(arr)
    fmt.Println(arr)  // [1 2 3 4 5] (não mudou!)
}
```

**⚠️ Custo de Performance:**

- Arrays grandes são **copiados completamente**
- Structs grandes são **copiados completamente**
- Pode ser custoso para dados grandes!

### Slices, Maps e Channels

**IMPORTANTE:** Slices, maps e channels são passados por valor, mas o **valor é uma referência**!

```go
func modificarSlice(s []int) {
    if len(s) > 0 {
        s[0] = 999  // Modifica o array subjacente!
    }
    s = append(s, 100)  // Mas append não afeta o slice original
}

func main() {
    slice := []int{1, 2, 3}
    modificarSlice(slice)
    fmt.Println(slice)  // [999 2 3] (elementos modificados!)
    // Mas length não mudou (append não afeta original)
}
```

**Por quê?**

- Slice é um **header** (ponteiro + length + capacity)
- Header é copiado, mas aponta para o mesmo array
- Modificar elementos modifica o array compartilhado
- Mas `append` pode criar novo array (não afeta original)

### Como Modificar Valores Originais?

**Use ponteiros:**

```go
func modificarIntPonteiro(numero *int) {
    *numero = 999  // Modifica o valor original
}

func main() {
    num := 42
    modificarIntPonteiro(&num)
    fmt.Println(num)  // 999 (mudou!)
}
```

---

## 10. Resumo dos Conceitos-Chave

- **Functions**: Blocos de código reutilizáveis e organizados
- **First-class citizens**: Funções podem ser valores (atribuir, passar, retornar)
- **Variadic functions**: Aceitam número variável de argumentos (`...Tipo`)
- **Multiple returns**: Funções podem retornar múltiplos valores
- **Named returns**: Retornos podem ser nomeados e tratados como variáveis
- **Anonymous functions**: Funções sem nome, úteis para callbacks
- **Closures**: Funções que capturam variáveis do escopo externo
- **Call by value**: Tudo é passado por cópia (mas slices/maps são referências)

---

## Conclusão

Funções são a base da organização e reutilização de código em Go. Entender todos esses conceitos é fundamental para escrever código limpo, modular e eficiente.

Na próxima parte desta aula, vamos simplificar esses conceitos com analogias do dia a dia para fixar melhor o aprendizado!
