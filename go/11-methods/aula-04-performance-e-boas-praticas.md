# Aula 4: Performance e Boas Práticas - Methods vs Functions

Olá! Nesta aula, vamos mergulhar nas questões de **performance**, **boas práticas** e **decisões de design** relacionadas a methods e functions em Go. Esses conhecimentos são cruciais para escrever código Go profissional, eficiente e idiomático.

---

## 1. Performance: Value vs Pointer Receivers

### O Custo de Copiar

Quando você usa um **value receiver**, Go precisa **copiar** toda a struct. Para structs pequenas, isso é trivial. Mas para structs grandes, pode ser custoso:

```go
// Struct pequena: copiar é barato
type Ponto struct {
    X, Y float64  // 16 bytes
}

func (p Ponto) Distancia() float64 {  // Copia 16 bytes - OK
    return math.Sqrt(p.X*p.X + p.Y*p.Y)
}

// Struct grande: copiar é caro
type DadosCompletos struct {
    Nome     string
    Email    string
    Telefone string
    Endereco string
    Historico [1000]int  // 4000 bytes!
    // ... muitos outros campos
}

func (d DadosCompletos) Processar() {  // Copia 4000+ bytes - CUSTOSO!
    // processamento
}
```

### Quando a Performance Importa?

**Regra prática**: Se sua struct tem mais de **~100 bytes** ou contém slices/maps grandes, considere usar pointer receiver mesmo para métodos que não modificam.

```go
// ❌ EVITE: Struct grande com value receiver
type Usuario struct {
    ID       int
    Nome     string
    Email    string
    Historico []string  // Pode ser grande
    Configuracoes map[string]interface{}  // Pode ser grande
}

func (u Usuario) NomeCompleto() string {  // Copia tudo!
    return u.Nome
}

// ✅ PREFIRA: Pointer receiver para structs grandes
func (u *Usuario) NomeCompleto() string {  // Apenas copia o pointer (8 bytes)
    return u.Nome
}
```

### Benchmark: Value vs Pointer

Vamos ver a diferença na prática:

```go
type StructPequena struct {
    a, b, c, d int  // 32 bytes
}

type StructGrande struct {
    dados [1000]int  // 4000 bytes
}

// Value receiver
func (s StructPequena) MetodoValor() int {
    return s.a + s.b
}

// Pointer receiver
func (s *StructPequena) MetodoPointer() int {
    return s.a + s.b
}

// Benchmark mostra:
// StructPequena: diferença mínima (nanossegundos)
// StructGrande: pointer receiver é significativamente mais rápido
```

**Conclusão**: Para structs pequenas (< 100 bytes), a diferença é negligenciável. Para structs grandes, pointer receiver é mais eficiente.

---

## 2. Consistência: A Regra de Ouro

### Por Que Ser Consistente?

**Regra fundamental**: Se **qualquer** método de um tipo usa pointer receiver, **todos** os métodos desse tipo devem usar pointer receiver.

### Exemplo do Problema

```go
// ❌ EVITE: Mistura de receivers
type Conta struct {
    saldo float64
}

func (c Conta) Saldo() float64 {  // Value receiver
    return c.saldo
}

func (c *Conta) Depositar(valor float64) {  // Pointer receiver
    c.saldo += valor
}

// Problema: Agora você tem que lembrar qual método usa qual receiver
// Isso torna o código confuso e propenso a erros
```

### Solução: Consistência

```go
// ✅ PREFIRA: Todos os métodos com pointer receiver
type Conta struct {
    saldo float64
}

func (c *Conta) Saldo() float64 {  // Pointer receiver (consistente)
    return c.saldo
}

func (c *Conta) Depositar(valor float64) {  // Pointer receiver
    c.saldo += valor
}

func (c *Conta) Sacar(valor float64) bool {  // Pointer receiver
    if valor <= c.saldo {
        c.saldo -= valor
        return true
    }
    return false
}
```

**Benefícios da consistência**:
- Código mais previsível
- Mais fácil de manter
- Menos erros
- Mais legível

---

## 3. Methods vs Functions: Decisões de Design

### Quando Usar Methods?

Use methods quando:

1. **O comportamento pertence ao tipo**: A operação é uma ação que o tipo "faz"
2. **Você precisa implementar interfaces**: Methods são necessários para interfaces
3. **API mais limpa**: `obj.Metodo()` é mais legível que `Funcao(obj)`
4. **Encapsulamento**: Você quer controlar como o tipo é usado

```go
// ✅ BOM: Method - comportamento do retângulo
type Retangulo struct {
    Largura, Altura float64
}

func (r Retangulo) Area() float64 {
    return r.Largura * r.Altura
}

// Uso: ret.Area() - natural e intuitivo
```

### Quando Usar Functions?

Use functions quando:

1. **Operação independente**: Não está relacionada a um tipo específico
2. **Utilitário geral**: Pode trabalhar com vários tipos
3. **Não há receiver natural**: A operação não "pertence" a nenhum tipo
4. **Flexibilidade**: Precisa trabalhar com diferentes tipos

```go
// ✅ BOM: Function - operação matemática geral
func CalcularAreaRetangulo(largura, altura float64) float64 {
    return largura * altura
}

// ✅ BOM: Function - comparação entre dois objetos
func CompararRetangulos(r1, r2 Retangulo) bool {
    return r1.Area() == r2.Area()
}

// ✅ BOM: Function - operação que não pertence a um tipo
func ValidarDimensoes(largura, altura float64) error {
    if largura <= 0 || altura <= 0 {
        return fmt.Errorf("dimensões inválidas")
    }
    return nil
}
```

### Casos Especiais

**Operações que modificam múltiplos objetos**: Geralmente são functions

```go
// ✅ Function: modifica dois objetos
func TrocarValores(a, b *int) {
    *a, *b = *b, *a
}

// Não faz sentido como method porque não pertence a um tipo específico
```

**Operações de construção**: Geralmente são functions

```go
// ✅ Function: construtor
func NovoRetangulo(largura, altura float64) *Retangulo {
    return &Retangulo{
        Largura: largura,
        Altura:  altura,
    }
}
```

---

## 4. Boas Práticas: Nomenclatura

### Nomes de Methods

Methods devem ter nomes que descrevam **o que o objeto faz**:

```go
// ✅ BOM
func (c *Conta) Depositar(valor float64)
func (c *Conta) Sacar(valor float64)
func (u *Usuario) AtualizarEmail(novoEmail string)

// ❌ EVITE
func (c *Conta) Deposito(valor float64)  // Nome de substantivo
func (c *Conta) FazerSaque(valor float64)  // "Fazer" é redundante
```

### Convenções Go

- Methods que retornam valores geralmente não têm prefixo "Get"
- Methods que modificam devem usar verbos de ação
- Methods booleanos geralmente começam com "Is", "Has", "Can"

```go
// ✅ BOM: Sem "Get"
func (u *Usuario) Nome() string
func (c *Conta) Saldo() float64

// ✅ BOM: Verbos de ação
func (u *Usuario) Atualizar(nome string)
func (c *Conta) Transferir(valor float64, destino *Conta)

// ✅ BOM: Booleanos
func (u *Usuario) IsAtivo() bool
func (c *Conta) PodeSacar(valor float64) bool
```

---

## 5. Methods em Tipos Não-Struct: Quando Usar

### Tipos Customizados

Methods em tipos customizados são úteis para:

1. **Adicionar comportamento a tipos básicos**
2. **Criar APIs mais expressivas**
3. **Implementar interfaces**

```go
// ✅ BOM: Adiciona comportamento útil
type Dinheiro int

func (d Dinheiro) Reais() string {
    return fmt.Sprintf("R$ %.2f", float64(d)/100)
}

// Uso: dinheiro.Reais() é mais expressivo que formatarDinheiro(dinheiro)
```

### Quando NÃO Usar

Não crie métodos em tipos customizados apenas por criar. Se o comportamento não adiciona valor real, use uma function:

```go
// ❌ EVITE: Method desnecessário
type Numero int

func (n Numero) Dobro() int {
    return int(n) * 2
}

// ✅ PREFIRA: Function é suficiente
func Dobro(n int) int {
    return n * 2
}
```

---

## 6. Performance: Métodos com Slices

### Slices como Receivers

Quando trabalhar com slices customizados, você geralmente precisa de pointer receiver para modificar:

```go
type ListaNumeros []int

// ❌ NÃO FUNCIONA: Value receiver não pode reatribuir slice
func (ln ListaNumeros) Adicionar(num int) {
    ln = append(ln, num)  // Não modifica o original!
}

// ✅ FUNCIONA: Pointer receiver pode reatribuir
func (ln *ListaNumeros) Adicionar(num int) {
    *ln = append(*ln, num)  // Modifica o original
}
```

### Performance de Slices

Slices já são "reference types" internamente, mas reatribuir requer pointer:

```go
// Para operações que modificam elementos: value receiver é OK
func (ln ListaNumeros) IncrementarTodos() {
    for i := range ln {
        ln[i]++  // Modifica elementos (funciona com value receiver)
    }
}

// Para operações que modificam o slice: pointer receiver é necessário
func (ln *ListaNumeros) Adicionar(num int) {
    *ln = append(*ln, num)  // Reatribui o slice (precisa pointer)
}
```

---

## 7. O Que NÃO Fazer: Anti-padrões

### ❌ Anti-padrão 1: Methods que não deveriam ser methods

```go
// ❌ EVITE: Operação genérica como method
type Retangulo struct {
    Largura, Altura float64
}

func (r Retangulo) Multiplicar(a, b float64) float64 {
    return a * b  // Isso não pertence ao retângulo!
}

// ✅ PREFIRA: Function genérica
func Multiplicar(a, b float64) float64 {
    return a * b
}
```

### ❌ Anti-padrão 2: Inconsistência de receivers

```go
// ❌ EVITE: Mistura desnecessária
type Usuario struct {
    nome string
}

func (u Usuario) Nome() string {  // Value
    return u.nome
}

func (u *Usuario) AtualizarNome(nome string) {  // Pointer
    u.nome = nome
}

// ✅ PREFIRA: Consistência
func (u *Usuario) Nome() string {
    return u.nome
}

func (u *Usuario) AtualizarNome(nome string) {
    u.nome = nome
}
```

### ❌ Anti-padrão 3: Methods muito complexos

```go
// ❌ EVITE: Method fazendo muitas coisas
func (u *Usuario) ProcessarTudo() {
    u.validar()
    u.salvar()
    u.enviarEmail()
    u.atualizarCache()
    u.logar()
    // ... muitas outras coisas
}

// ✅ PREFIRA: Methods menores e focados
func (u *Usuario) Validar() error { ... }
func (u *Usuario) Salvar() error { ... }
func (u *Usuario) EnviarEmail() error { ... }
```

---

## 8. Checklist de Boas Práticas

Use este checklist ao criar methods:

- [ ] O método realmente pertence ao tipo?
- [ ] Estou sendo consistente com os receivers (todos value ou todos pointer)?
- [ ] Para structs grandes, estou usando pointer receiver?
- [ ] O nome do método é claro e descritivo?
- [ ] O método tem uma responsabilidade única?
- [ ] Estou seguindo as convenções de nomenclatura do Go?
- [ ] O método não está fazendo coisas demais?

---

## 9. Resumo: Decisões de Design

### Value Receiver Use Quando:
- ✅ Struct é pequena (< 100 bytes)
- ✅ Método não modifica o receiver
- ✅ Você quer imutabilidade
- ✅ Todos os outros métodos também usam value receiver

### Pointer Receiver Use Quando:
- ✅ Struct é grande (> 100 bytes)
- ✅ Método modifica o receiver
- ✅ Qualquer outro método do tipo usa pointer receiver (consistência)
- ✅ Você precisa de eficiência máxima

### Method Use Quando:
- ✅ Comportamento pertence ao tipo
- ✅ Você precisa implementar interfaces
- ✅ API mais limpa e intuitiva

### Function Use Quando:
- ✅ Operação independente
- ✅ Utilitário geral
- ✅ Não há receiver natural
- ✅ Flexibilidade com múltiplos tipos

---

## 10. Conclusão

Nesta aula, você aprendeu:

✅ **Performance**: Pointer receivers são mais eficientes para structs grandes  
✅ **Consistência**: Mantenha todos os métodos de um tipo com o mesmo tipo de receiver  
✅ **Design**: Methods para comportamento do tipo, functions para operações gerais  
✅ **Nomenclatura**: Siga as convenções do Go  
✅ **Anti-padrões**: Evite métodos desnecessários e inconsistências  

**Lembre-se**: Go valoriza simplicidade, clareza e consistência. Quando em dúvida, escolha a opção mais simples e consistente com o resto do código.

**Próximos passos**: Continue praticando e observe como métodos e funções são usados em bibliotecas Go populares. Isso vai ajudar você a desenvolver uma intuição sobre o que é idiomático em Go.

---

**Dica Final**: Sempre que você criar um método, pergunte a si mesmo: "Isso realmente pertence a este tipo, ou seria melhor como uma function?" Essa simples pergunta vai melhorar muito a qualidade do seu código Go!

