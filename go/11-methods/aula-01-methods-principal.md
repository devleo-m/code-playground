# Módulo 11: Methods vs Functions em Go

## Aula 1: Methods vs Functions - Comportamento Orientado a Tipos

Olá! Bem-vindo ao décimo primeiro módulo. Até agora você aprendeu sobre funções, structs e pointers. Mas você já se perguntou: **qual a diferença entre uma função e um método?** E quando devo usar cada um?

Nesta aula, vamos mergulhar profundamente em **methods** (métodos) e como eles se relacionam com **functions** (funções) em Go. Você vai aprender quando usar cada um, como definir métodos com diferentes tipos de receivers, e as melhores práticas para escrever código Go idiomático.

---

## 1. O Que São Methods?

**Methods** (métodos) são funções especiais que pertencem a um tipo específico. Eles são definidos com um **receiver** (receptor), que é o tipo ao qual o método pertence. Em outras palavras, métodos são funções que "pertencem" a um tipo.

### Diferença Fundamental: Methods vs Functions

A principal diferença é:

- **Function**: É independente, pode ser chamada de qualquer lugar
- **Method**: Pertence a um tipo específico, é chamado através de uma instância desse tipo

### Sintaxe de um Method

```go
// Method: tem um receiver (p Pessoa)
func (p Pessoa) NomeCompleto() string {
    return p.Nome + " " + p.Sobrenome
}

// Function: não tem receiver
func NomeCompleto(p Pessoa) string {
    return p.Nome + " " + p.Sobrenome
}
```

**Observação importante**: Em Go, métodos são definidos **fora** da declaração do tipo. Isso é diferente de linguagens como Java ou C++, onde métodos são definidos dentro da classe.

---

## 2. Por Que Usar Methods?

Methods permitem que você adicione **comportamento** aos seus tipos, criando uma interface mais natural e intuitiva:

```go
type Retangulo struct {
    Largura float64
    Altura  float64
}

// Method: mais natural e intuitivo
func (r Retangulo) Area() float64 {
    return r.Largura * r.Altura
}

// Function: funciona, mas menos idiomático
func AreaRetangulo(r Retangulo) float64 {
    return r.Largura * r.Altura
}

func main() {
    ret := Retangulo{Largura: 10, Altura: 5}
    
    // Com method: parece que o retângulo "tem" uma área
    fmt.Println(ret.Area())  // 50
    
    // Com function: parece uma operação externa
    fmt.Println(AreaRetangulo(ret))  // 50
}
```

### Vantagens dos Methods

1. **Sintaxe mais limpa**: `ret.Area()` é mais legível que `AreaRetangulo(ret)`
2. **Organização**: Métodos relacionados ficam agrupados conceitualmente
3. **Interface implícita**: Permite implementar interfaces automaticamente
4. **Encapsulamento**: Facilita o controle de acesso e modificação de estado

---

## 3. Value Receivers

**Value receivers** recebem uma **cópia** do valor do tipo. Eles são definidos sem o `*`:

```go
type Contador struct {
    valor int
}

// Value receiver: recebe uma cópia
func (c Contador) Incrementar() {
    c.valor++  // Modifica apenas a cópia
}

func (c Contador) Valor() int {
    return c.valor  // Retorna o valor da cópia
}

func main() {
    contador := Contador{valor: 0}
    contador.Incrementar()
    fmt.Println(contador.Valor())  // 0 (não mudou!)
}
```

### Quando Usar Value Receivers?

Use value receivers quando:

1. **O método não modifica o receiver**: Métodos que apenas leem ou calculam valores
2. **O tipo é pequeno**: Tipos primitivos (int, string, bool) ou structs pequenas
3. **Imutabilidade é desejada**: Quando você quer garantir que o original não será modificado

### Exemplo Prático: Value Receiver

```go
type Ponto struct {
    X, Y float64
}

// Value receiver: não modifica, apenas calcula
func (p Ponto) DistanciaOrigem() float64 {
    return math.Sqrt(p.X*p.X + p.Y*p.Y)
}

// Value receiver: retorna um novo ponto (imutável)
func (p Ponto) Mover(dx, dy float64) Ponto {
    return Ponto{X: p.X + dx, Y: p.Y + dy}
}

func main() {
    p1 := Ponto{X: 3, Y: 4}
    fmt.Println(p1.DistanciaOrigem())  // 5
    
    p2 := p1.Mover(1, 1)
    fmt.Printf("p1: %+v\n", p1)  // {X:3 Y:4} (não mudou)
    fmt.Printf("p2: %+v\n", p2)  // {X:4 Y:5} (novo ponto)
}
```

### Comportamento Especial: Go Automaticamente Faz Conversão

Go é inteligente: você pode chamar um método com value receiver tanto em valores quanto em pointers:

```go
type Pessoa struct {
    Nome string
}

func (p Pessoa) Apresentar() {
    fmt.Printf("Olá, eu sou %s\n", p.Nome)
}

func main() {
    // Chamando em um valor
    p1 := Pessoa{Nome: "João"}
    p1.Apresentar()  // Funciona
    
    // Chamando em um pointer
    p2 := &Pessoa{Nome: "Maria"}
    p2.Apresentar()  // Também funciona! Go faz (*p2).Apresentar() automaticamente
}
```

---

## 4. Pointer Receivers

**Pointer receivers** recebem um **pointer** para o tipo. Eles são definidos com `*`:

```go
type Contador struct {
    valor int
}

// Pointer receiver: recebe um pointer
func (c *Contador) Incrementar() {
    c.valor++  // Modifica o original
}

func (c *Contador) Valor() int {
    return c.valor  // Retorna o valor original
}

func main() {
    contador := Contador{valor: 0}
    contador.Incrementar()
    fmt.Println(contador.Valor())  // 1 (mudou!)
}
```

### Quando Usar Pointer Receivers?

Use pointer receivers quando:

1. **O método modifica o receiver**: Precisa alterar o estado do objeto
2. **O tipo é grande**: Evita copiar structs grandes (mais eficiente)
3. **Consistência**: Se outros métodos do tipo usam pointer receivers, mantenha consistência
4. **Modificar estado interno**: Quando você precisa alterar campos da struct

### Exemplo Prático: Pointer Receiver

```go
type ContaBancaria struct {
    titular string
    saldo   float64
}

// Pointer receiver: modifica o estado
func (c *ContaBancaria) Depositar(valor float64) {
    c.saldo += valor
}

// Pointer receiver: modifica o estado
func (c *ContaBancaria) Sacar(valor float64) bool {
    if valor <= c.saldo {
        c.saldo -= valor
        return true
    }
    return false
}

// Pointer receiver: apenas lê, mas mantém consistência
func (c *ContaBancaria) Saldo() float64 {
    return c.saldo
}

func main() {
    conta := ContaBancaria{titular: "João", saldo: 1000}
    
    conta.Depositar(500)
    fmt.Println(conta.Saldo())  // 1500
    
    conta.Sacar(200)
    fmt.Println(conta.Saldo())  // 1300
}
```

### Comportamento Especial: Go Automaticamente Faz Conversão

Assim como com value receivers, Go automaticamente converte quando necessário:

```go
type Pessoa struct {
    Nome string
}

func (p *Pessoa) MudarNome(novoNome string) {
    p.Nome = novoNome
}

func main() {
    // Chamando em um valor
    p1 := Pessoa{Nome: "João"}
    p1.MudarNome("Pedro")  // Funciona! Go faz (&p1).MudarNome() automaticamente
    fmt.Println(p1.Nome)   // Pedro
    
    // Chamando em um pointer
    p2 := &Pessoa{Nome: "Maria"}
    p2.MudarNome("Ana")  // Funciona normalmente
    fmt.Println(p2.Nome)  // Ana
}
```

**Importante**: Embora Go faça essa conversão automaticamente, é uma boa prática ser consistente. Se você tem métodos com pointer receiver, geralmente você deve trabalhar com pointers do tipo.

---

## 5. Comparação: Value vs Pointer Receivers

Vamos ver um exemplo completo comparando os dois:

```go
type Produto struct {
    Nome  string
    Preco float64
    Vendas int
}

// Value receiver: não modifica
func (p Produto) PrecoComDesconto(percentual float64) float64 {
    return p.Preco * (1 - percentual/100)
}

// Pointer receiver: modifica o estado
func (p *Produto) RegistrarVenda() {
    p.Vendas++
}

// Pointer receiver: modifica o estado
func (p *Produto) AtualizarPreco(novoPreco float64) {
    p.Preco = novoPreco
}

func main() {
    produto := Produto{Nome: "Notebook", Preco: 2000, Vendas: 0}
    
    // Value receiver: não modifica o original
    precoComDesconto := produto.PrecoComDesconto(10)
    fmt.Printf("Preço original: %.2f\n", produto.Preco)  // 2000.00
    fmt.Printf("Preço com desconto: %.2f\n", precoComDesconto)  // 1800.00
    
    // Pointer receiver: modifica o original
    produto.RegistrarVenda()
    fmt.Printf("Vendas: %d\n", produto.Vendas)  // 1
    
    produto.AtualizarPreco(1800)
    fmt.Printf("Novo preço: %.2f\n", produto.Preco)  // 1800.00
}
```

### Regra de Ouro

**Se você tem qualquer método com pointer receiver, todos os métodos do tipo devem usar pointer receivers para manter consistência.**

---

## 6. Methods com Parâmetros e Retornos

Methods podem ter parâmetros e retornar valores, assim como funções:

```go
type Calculadora struct {
    historico []float64
}

// Method com parâmetros e retorno
func (c *Calculadora) Somar(a, b float64) float64 {
    resultado := a + b
    c.historico = append(c.historico, resultado)
    return resultado
}

// Method com múltiplos retornos
func (c *Calculadora) Dividir(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("divisão por zero")
    }
    resultado := a / b
    c.historico = append(c.historico, resultado)
    return resultado, nil
}

// Method sem parâmetros, com retorno
func (c *Calculadora) Historico() []float64 {
    return c.historico
}

func main() {
    calc := Calculadora{}
    
    resultado1, _ := calc.Somar(10, 5)
    fmt.Println(resultado1)  // 15
    
    resultado2, err := calc.Dividir(20, 4)
    if err != nil {
        fmt.Println("Erro:", err)
    } else {
        fmt.Println(resultado2)  // 5
    }
    
    fmt.Println(calc.Historico())  // [15 5]
}
```

---

## 7. Methods em Tipos Não-Struct

Em Go, você pode definir methods para **qualquer tipo**, não apenas structs! Isso inclui tipos básicos e tipos definidos por você:

```go
// Definindo um novo tipo baseado em int
type Dinheiro int

// Method para o tipo Dinheiro
func (d Dinheiro) Reais() string {
    return fmt.Sprintf("R$ %.2f", float64(d)/100)
}

// Method para o tipo Dinheiro
func (d Dinheiro) Centavos() int {
    return int(d)
}

func main() {
    valor := Dinheiro(1250)  // 12,50 reais em centavos
    fmt.Println(valor.Reais())    // R$ 12.50
    fmt.Println(valor.Centavos()) // 1250
}
```

### Methods em Slices Customizados

```go
// Tipo customizado baseado em slice
type ListaNumeros []int

// Method para adicionar número
func (ln *ListaNumeros) Adicionar(num int) {
    *ln = append(*ln, num)
}

// Method para calcular média
func (ln ListaNumeros) Media() float64 {
    if len(ln) == 0 {
        return 0
    }
    soma := 0
    for _, num := range ln {
        soma += num
    }
    return float64(soma) / float64(len(ln))
}

func main() {
    lista := ListaNumeros{}
    lista.Adicionar(10)
    lista.Adicionar(20)
    lista.Adicionar(30)
    
    fmt.Println(lista.Media())  // 20
}
```

**Importante**: Para modificar um slice customizado, você precisa usar pointer receiver, porque reatribuir um slice requer modificar o pointer.

---

## 8. Methods vs Functions: Quando Usar Cada Um?

### Use Methods Quando:

1. **O comportamento pertence ao tipo**: Se a operação é uma ação que o tipo "faz"
2. **Você precisa de polimorfismo**: Para implementar interfaces
3. **Encapsulamento**: Quando você quer controlar como o tipo é usado
4. **API mais limpa**: Quando `obj.Metodo()` é mais legível que `Funcao(obj)`

### Use Functions Quando:

1. **Operação independente**: A função não está relacionada a um tipo específico
2. **Utilitários gerais**: Funções auxiliares que podem trabalhar com vários tipos
3. **Não há estado**: Quando você não precisa de um receiver
4. **Flexibilidade**: Quando a função pode trabalhar com diferentes tipos

### Exemplo Comparativo

```go
type Retangulo struct {
    Largura, Altura float64
}

// Method: comportamento do retângulo
func (r Retangulo) Area() float64 {
    return r.Largura * r.Altura
}

// Function: operação matemática geral
func AreaRetangulo(largura, altura float64) float64 {
    return largura * altura
}

// Function: pode trabalhar com diferentes tipos
func CalcularArea(forma interface{}) float64 {
    switch v := forma.(type) {
    case Retangulo:
        return v.Area()
    case Circulo:
        return v.Area()
    default:
        return 0
    }
}
```

---

## 9. Métodos com Múltiplos Receivers (Não Suportado)

**Importante**: Em Go, um método pode ter apenas **um receiver**. Se você precisar de comportamento compartilhado entre tipos, use **interfaces** (que veremos em módulos futuros).

```go
// ❌ ISSO NÃO FUNCIONA EM GO
func (r Retangulo, c Circulo) Comparar() {
    // Erro: métodos não podem ter múltiplos receivers
}

// ✅ USE INTERFACES EM VEZ DISSO (veremos depois)
type Forma interface {
    Area() float64
}
```

---

## 10. Resumo e Próximos Passos

Nesta aula, você aprendeu:

✅ **Methods**: Funções com receivers que pertencem a tipos específicos  
✅ **Value Receivers**: Recebem cópia, usados quando não há modificação  
✅ **Pointer Receivers**: Recebem pointer, usados para modificar estado  
✅ **Conversão automática**: Go converte entre valores e pointers automaticamente  
✅ **Methods em qualquer tipo**: Não apenas structs, mas qualquer tipo definido  
✅ **Quando usar cada um**: Methods para comportamento do tipo, functions para operações gerais  

**Pontos importantes para lembrar:**

- Methods pertencem a tipos, functions são independentes
- Use value receivers para métodos que não modificam
- Use pointer receivers para métodos que modificam ou quando o tipo é grande
- Seja consistente: se um método usa pointer receiver, todos devem usar
- Go faz conversão automática entre valores e pointers ao chamar métodos

**Regra prática:**
- **Value receiver**: Métodos que apenas leem ou calculam
- **Pointer receiver**: Métodos que modificam estado ou tipos grandes
- **Consistência**: Todos os métodos de um tipo devem usar o mesmo tipo de receiver

Na próxima aula, vamos simplificar esses conceitos com analogias do dia a dia para fixar o aprendizado!

---

**Dica**: Execute o arquivo `01-exemplos.go` para ver todos esses conceitos em ação. Experimente modificar os exemplos e veja o que acontece!

