# Módulo 9: Functions em Go

## Aula 2 - Simplificada: Entendendo Functions

Agora vamos entender funções de forma mais simples, usando analogias do nosso dia a dia!

---

## 1. Function: Uma Receita Reutilizável

### Analogia: Receita de Bolo

Pense em uma função como uma **receita de bolo**:

**Mundo real:**

```
Receita de Bolo:
1. Misturar farinha e açúcar
2. Adicionar ovos
3. Assar por 30 minutos
4. Resultado: Bolo pronto!
```

**Função em Go:**

```go
func fazerBolo() {
    fmt.Println("Misturar farinha e açúcar")
    fmt.Println("Adicionar ovos")
    fmt.Println("Assar por 30 minutos")
    fmt.Println("Bolo pronto!")
}
```

**Por que é útil:**

- Você escreve a receita **uma vez**
- Pode usar **quantas vezes quiser**
- Se precisar mudar, muda **só na receita**

---

## 2. Parâmetros: Ingredientes da Receita

### Analogia: Receita com Ingredientes Variáveis

Uma receita pode ter **ingredientes que mudam**:

**Mundo real:**

```
Receita de Suco:
- Ingrediente: Fruta (pode ser laranja, maçã, etc.)
- Resultado: Suco da fruta escolhida
```

**Função em Go:**

```go
func fazerSuco(fruta string) {
    fmt.Printf("Fazendo suco de %s\n", fruta)
}

fazerSuco("laranja")  // Faz suco de laranja
fazerSuco("maçã")     // Faz suco de maçã
```

**Parâmetros são como "espaços vazios"** na receita que você preenche quando usa!

---

## 3. Retorno: O Resultado da Receita

### Analogia: Receita que Produz Algo

Algumas receitas **produzem algo** que você pode usar depois:

**Mundo real:**

```
Receita de Torta:
- Ingredientes: massa, recheio
- Resultado: Torta pronta (que você pode comer)
```

**Função em Go:**

```go
func fazerTorta(massa, recheio string) string {
    return fmt.Sprintf("Torta de %s com %s", massa, recheio)
}

torta := fazerTorta("chocolate", "morango")
fmt.Println(torta)  // "Torta de chocolate com morango"
```

**Retorno é como o "produto final"** da sua receita!

---

## 4. First-Class Citizen: Receita como Objeto

### Analogia: Receita Escrita em Papel

Em Go, funções são como **receitas escritas em papel** que você pode:

- **Guardar** em uma gaveta (atribuir a variável)
- **Dar para alguém** (passar como argumento)
- **Receber de alguém** (retornar de outra função)

**Mundo real:**

```
Você tem uma receita de bolo escrita em um papel.
- Pode guardar na gaveta
- Pode dar para um amigo
- Pode receber receitas de outros
```

**Função em Go:**

```go
// Guardar receita na "gaveta" (variável)
receita := func() {
    fmt.Println("Fazer bolo")
}

// Dar receita para "amigo" (passar como argumento)
func usarReceita(r func()) {
    r()  // Usar a receita
}

usarReceita(receita)
```

**É como ter receitas que são "objetos" que você pode manipular!**

---

## 5. Variadic Functions: Receita com Ingredientes Flexíveis

### Analogia: Receita de Salada

Algumas receitas aceitam **quantos ingredientes você quiser**:

**Mundo real:**

```
Receita de Salada:
- Ingredientes: pode colocar alface, tomate, cebola, pepino...
- Não tem limite de ingredientes!
```

**Função em Go:**

```go
func fazerSalada(ingredientes ...string) {
    fmt.Println("Salada com:")
    for _, ing := range ingredientes {
        fmt.Printf("- %s\n", ing)
    }
}

fazerSalada("alface", "tomate")
fazerSalada("alface", "tomate", "cebola", "pepino")
```

**Os três pontos `...` significam "quantos você quiser"!**

---

## 6. Multiple Returns: Receita que Produz Várias Coisas

### Analogia: Receita que Faz Comida e Bebida

Algumas receitas **produzem mais de uma coisa**:

**Mundo real:**

```
Receita Completa:
- Produz: Prato principal E sobremesa
- Você recebe duas coisas!
```

**Função em Go:**

```go
func fazerRefeicao() (string, string) {
    prato := "Arroz com feijão"
    sobremesa := "Pudim"
    return prato, sobremesa
}

prato, sobremesa := fazerRefeicao()
fmt.Printf("Prato: %s, Sobremesa: %s\n", prato, sobremesa)
```

**É como receber uma "caixa" com várias coisas dentro!**

### Padrão Resultado + Erro

**Analogia: Pedido de Comida**

Quando você pede comida, pode receber:

- **Sucesso**: Comida + sem problemas
- **Erro**: Nada + problema (ex: restaurante fechado)

```go
func pedirComida() (string, error) {
    if restauranteFechado {
        return "", fmt.Errorf("restaurante fechado")
    }
    return "Hambúrguer", nil
}

comida, err := pedirComida()
if err != nil {
    fmt.Printf("Erro: %v\n", err)
} else {
    fmt.Printf("Comida: %s\n", comida)
}
```

---

## 7. Named Returns: Caixas com Etiquetas

### Analogia: Caixa de Comida com Etiquetas

Em vez de "caixa 1" e "caixa 2", você tem **caixas com etiquetas**:

**Mundo real:**

```
Caixa de Comida:
- Etiqueta "Prato Principal": Arroz
- Etiqueta "Sobremesa": Pudim
- Você sabe o que tem em cada caixa!
```

**Função em Go:**

```go
func fazerRefeicao() (pratoPrincipal string, sobremesa string) {
    pratoPrincipal = "Arroz com feijão"  // Coloca na caixa "pratoPrincipal"
    sobremesa = "Pudim"                  // Coloca na caixa "sobremesa"
    return  // Retorna as caixas com etiquetas
}

prato, sobremesa := fazerRefeicao()
```

**É como ter caixas com nomes ao invés de números!**

---

## 8. Anonymous Functions: Receita Sem Nome

### Analogia: Receita Rápida Sem Escrever

Às vezes você faz algo **rapidamente sem escrever a receita**:

**Mundo real:**

```
Você quer fazer um sanduíche rápido:
- Não escreve receita
- Faz na hora
- Usa uma vez e pronto
```

**Função em Go:**

```go
// Receita sem nome, feita na hora
func() {
    fmt.Println("Fazendo sanduíche rápido")
}()

// Ou guardar em variável
sanduiche := func() {
    fmt.Println("Fazendo sanduíche")
}
sanduiche()
```

**É como fazer algo rápido sem criar uma "receita formal"!**

---

## 9. Closures: Receita que Lembra

### Analogia: Receita com Memória

Uma closure é como uma **receita que lembra** o que você fez antes:

**Mundo real:**

```
Receita de Contador de Bolos:
- Lembra quantos bolos você já fez
- Cada vez que faz, adiciona 1
- Nunca esquece o total!
```

**Função em Go:**

```go
func criarContador() func() int {
    total := 0  // "Memória" da receita
    return func() int {
        total++  // Lembra e adiciona
        return total
    }
}

contador := criarContador()
fmt.Println(contador())  // 1 (lembrou que começou em 0)
fmt.Println(contador())  // 2 (lembrou que era 1)
fmt.Println(contador())  // 3 (lembrou que era 2)
```

**É como ter uma receita com uma "nota" que ela nunca perde!**

### Closure com Parâmetro Capturado

**Analogia: Receita Personalizada**

```go
func criarMultiplicador(fator int) func(int) int {
    return func(numero int) int {
        return numero * fator  // Lembra o "fator"
    }
}

dobrar := criarMultiplicador(2)  // Receita que sempre dobra
triplicar := criarMultiplicador(3)  // Receita que sempre triplica

fmt.Println(dobrar(5))      // 10 (lembrou: sempre dobrar)
fmt.Println(triplicar(5))    // 15 (lembrou: sempre triplicar)
```

**Cada receita "lembra" seu próprio fator!**

### ⚠️ Armadilha: Closure em Loop

**Analogia: Receitas que Compartilham Ingrediente**

**PROBLEMA:**

```
Você quer criar 3 receitas diferentes:
- Receita 1: usar ingrediente i=0
- Receita 2: usar ingrediente i=1
- Receita 3: usar ingrediente i=2

Mas todas as receitas olham para o MESMO ingrediente!
Quando você termina, todas veem i=3 (valor final)!
```

**SOLUÇÃO:**

```
Criar uma CÓPIA do ingrediente para cada receita:
- Receita 1: tem sua própria cópia de i=0
- Receita 2: tem sua própria cópia de i=1
- Receita 3: tem sua própria cópia de i=2
```

**Código correto:**

```go
for i := 0; i < 3; i++ {
    i := i  // Criar cópia para cada receita
    receitas = append(receitas, func() int {
        return i  // Cada uma tem sua própria cópia
    })
}
```

---

## 10. Call by Value: Receita com Cópia dos Ingredientes

### Analogia: Receita que Usa Cópias

Quando você passa ingredientes para uma receita, Go faz **cópias**:

**Mundo real:**

```
Você tem 5 maçãs.
Dá as maçãs para fazer suco.
A receita recebe CÓPIAS das maçãs.
Suas maçãs originais continuam intactas!
```

**Função em Go:**

```go
func fazerSuco(quantidade int) {
    quantidade = 999  // Modifica apenas a CÓPIA
}

func main() {
    minhasMacas := 5
    fazerSuco(minhasMacas)
    fmt.Println(minhasMacas)  // Ainda 5! (original não mudou)
}
```

### Quando a Cópia é Grande?

**Analogia: Receita com Caixa Gigante**

**Arrays grandes:**

```
Você tem uma caixa GIGANTE com 1000 itens.
Para usar na receita, precisa COPIAR tudo!
Muito trabalho e demorado!
```

**Slices (referência):**

```
Você tem uma "etiqueta" que aponta para a caixa.
Passa a etiqueta (não a caixa inteira).
A receita pode ver e modificar a caixa original!
Mas se criar nova caixa, não afeta a original.
```

**Código:**

```go
// Array: copia tudo (lento para grandes)
func usarArray(arr [1000]int) {
    // Copiou 1000 números!
}

// Slice: copia apenas "etiqueta" (rápido)
func usarSlice(s []int) {
    s[0] = 999  // Modifica original!
    s = append(s, 100)  // Mas append não afeta original
}
```

---

## 11. Resumo Visual

**Function = Receita**

- Parâmetros = Ingredientes
- Retorno = Produto final
- First-class = Receita como objeto manipulável
- Variadic = Ingredientes flexíveis
- Multiple returns = Múltiplos produtos
- Named returns = Produtos com etiquetas
- Anonymous = Receita rápida sem nome
- Closure = Receita com memória
- Call by value = Usa cópias dos ingredientes

---

## 12. Analogia Final: Cozinha Completa

Imagine uma **cozinha completa**:

```go
// Receitas básicas (funções simples)
func fazerArroz() { ... }
func fazerFeijao() { ... }

// Receitas com ingredientes (parâmetros)
func temperar(comida string, tempero string) { ... }

// Receitas que produzem algo (retorno)
func cozinhar() string { return "Prato pronto" }

// Receitas flexíveis (variadic)
func fazerSalada(...ingredientes) { ... }

// Receitas que produzem várias coisas (multiple returns)
func fazerRefeicao() (prato, sobremesa string) { ... }

// Receitas que lembram (closures)
func criarReceitaPersonalizada() func() { ... }

// Receitas rápidas (anonymous)
func() { fazer algo rápido }()
```

**Todas trabalham juntas para criar um sistema de cozinha organizado!**

---

Agora que você entendeu os conceitos de forma simplificada, vamos praticar com exercícios na próxima parte!

