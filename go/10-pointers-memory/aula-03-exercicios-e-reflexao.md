# Aula 3: Exercícios e Reflexão - Pointers e Memory Management

Olá! Agora é hora de colocar a mão na massa e praticar tudo que aprendemos sobre pointers e memory management. Vamos fazer alguns exercícios práticos e, mais importante, **refletir** sobre os conceitos.

---

## 📝 Exercícios Práticos

### Exercício 1: Pointer Básico

Crie um programa que:
1. Declare uma variável `numero` com valor `42`
2. Crie um pointer `ptr` que aponte para `numero`
3. Imprima o valor de `numero`, o endereço de `numero`, o valor de `ptr` (endereço) e o valor apontado por `ptr`
4. Modifique o valor através do pointer para `100`
5. Imprima novamente o valor de `numero` para confirmar a mudança

**Dica**: Use `fmt.Printf` com `%p` para imprimir endereços.

---

### Exercício 2: Função que Modifica Valor

Crie uma função chamada `dobrar` que recebe um pointer para `int` e dobra o valor apontado.

```go
func main() {
    x := 5
    fmt.Printf("Antes: %d\n", x)  // Deve imprimir: Antes: 5
    dobrar(&x)
    fmt.Printf("Depois: %d\n", x)  // Deve imprimir: Depois: 10
}
```

Implemente a função `dobrar`.

---

### Exercício 3: Struct com Pointer

Crie uma struct `ContaBancaria` com os campos:
- `Titular` (string)
- `Saldo` (float64)

Crie duas funções:
1. `depositar(conta *ContaBancaria, valor float64)` - adiciona valor ao saldo
2. `sacar(conta *ContaBancaria, valor float64) bool` - subtrai valor do saldo e retorna `true` se houver saldo suficiente, `false` caso contrário

Teste as funções criando uma conta com saldo inicial de 1000 e fazendo algumas operações.

---

### Exercício 4: Comparando Passagem por Valor vs Referência

Crie um programa que demonstre a diferença entre passagem por valor e por referência:

```go
type Produto struct {
    Nome  string
    Preco float64
}

func aumentarPrecoPorValor(p Produto) {
    // Implemente aqui
}

func aumentarPrecoPorReferencia(p *Produto) {
    // Implemente aqui
}

func main() {
    produto := Produto{Nome: "Notebook", Preco: 3000.0}
    
    // Teste ambas as funções e mostre a diferença
}
```

Implemente as duas funções que aumentam o preço em 10% e mostre por que uma funciona e a outra não.

---

## 🤔 Perguntas de Reflexão

### Reflexão 1: Por Que Pointers Existem?

Pense sobre esta situação:

Você tem uma função que precisa processar uma lista com **1 milhão de elementos**. Se você passar essa lista por valor, Go precisa **copiar** todos os 1 milhão de elementos. Se você passar por referência (pointer), Go só precisa copiar o **endereço** (8 bytes em sistemas 64-bit).

**Perguntas para refletir:**
1. Qual abordagem é mais eficiente em termos de memória? Por quê?
2. Em que situações você **não deveria** usar pointers mesmo que a struct seja grande?
3. Por que Go permite passar slices e maps sem pointer explícito, mas structs grandes precisam de pointer?

**Escreva suas reflexões aqui:**
```
[Seu espaço para escrever]
```

---

### Reflexão 2: O Dilema do Nil Pointer

Considere este código:

```go
func processarDados(dados *[]int) {
    if dados == nil {
        fmt.Println("Dados não fornecidos!")
        return
    }
    
    for _, valor := range *dados {
        fmt.Println(valor)
    }
}

func main() {
    var dados *[]int  // nil
    processarDados(dados)  // Funciona? Por quê?
    
    dadosVazios := &[]int{}  // Pointer para slice vazio
    processarDados(dadosVazios)  // Funciona? Por quê?
}
```

**Perguntas para refletir:**
1. Qual é a diferença entre um pointer `nil` e um pointer para um slice vazio?
2. Por que é importante verificar `nil` antes de usar um pointer?
3. Em sua opinião, qual é a melhor forma de lidar com pointers `nil` em funções? Você prefere retornar erro, usar valores padrão, ou outra abordagem?

**Escreva suas reflexões aqui:**
```
[Seu espaço para escrever]
```

---

### Reflexão 3: Slices são "Reference Types" - Mas Por Quê?

Você aprendeu que slices são reference types e que modificações dentro de funções afetam o original. Mas também aprendeu que **reatribuir** um slice dentro de uma função não afeta o original.

**Perguntas para refletir:**
1. Por que Go foi projetado dessa forma? Qual é a vantagem de slices serem reference types para elementos, mas não para reatribuição?
2. Em que situações você **precisaria** passar um pointer para um slice (`*[]int`) ao invés de apenas o slice (`[]int`)?
3. Se slices já são "reference types", por que ainda existem casos onde usar pointers com slices é necessário?

**Dica**: Pense sobre a diferença entre modificar **elementos** de um slice vs **substituir** o slice inteiro.

**Escreva suas reflexões aqui:**
```
[Seu espaço para escrever]
```

---

### Reflexão 4: Garbage Collection - Bênção ou Maldição?

O Garbage Collector do Go é automático e facilita muito a vida do programador. Mas ele também tem um "custo":

- O GC precisa rodar periodicamente
- Durante a coleta, pode haver pequenas pausas
- O GC consome recursos do sistema

**Perguntas para refletir:**
1. Por que linguagens como C e C++ não têm Garbage Collector? Quais são as vantagens e desvantagens?
2. Em que tipos de aplicações o GC pode ser um problema? (Dica: pense em sistemas em tempo real, jogos, etc.)
3. Se você estivesse projetando uma linguagem de programação, você incluiria GC automático? Por quê?
4. Como você pode escrever código Go que "ajuda" o GC a trabalhar melhor? (Dica: pense em reduzir alocações desnecessárias)

**Escreva suas reflexões aqui:**
```
[Seu espaço para escrever]
```

---

### Reflexão 5: Stack vs Heap - A Decisão do Compilador

O compilador Go decide automaticamente onde alocar cada variável (stack ou heap) através de escape analysis.

**Perguntas para refletir:**
1. Por que o compilador precisa fazer essa análise? Por que não colocar tudo no heap ou tudo na stack?
2. O que aconteceria se todas as variáveis fossem alocadas na stack? E se todas fossem no heap?
3. Em sua opinião, é melhor que o programador controle isso manualmente (como em C) ou que o compilador decida automaticamente (como em Go)? Por quê?
4. Você consegue pensar em uma situação onde você **gostaria** de forçar uma variável a ir para o heap ou para a stack? Por quê?

**Escreva suas reflexões aqui:**
```
[Seu espaço para escrever]
```

---

## ✅ Checklist de Aprendizado

Antes de prosseguir, certifique-se de que você consegue:

- [ ] Declarar e usar pointers básicos
- [ ] Entender a diferença entre `&` (address-of) e `*` (dereference)
- [ ] Explicar quando usar passagem por valor vs por referência
- [ ] Usar pointers com structs e entender o shorthand do Go
- [ ] Entender por que slices e maps não precisam de pointers explícitos na maioria dos casos
- [ ] Explicar a diferença entre stack e heap
- [ ] Entender o que é escape analysis
- [ ] Explicar o que é Garbage Collection e por que é útil
- [ ] Verificar `nil` antes de usar pointers

---

## 🎯 Desafio Extra (Opcional)

Crie uma função `trocar` que recebe dois pointers para `int` e troca os valores entre eles:

```go
func trocar(a, b *int) {
    // Implemente aqui
}

func main() {
    x := 10
    y := 20
    fmt.Printf("Antes: x=%d, y=%d\n", x, y)
    trocar(&x, &y)
    fmt.Printf("Depois: x=%d, y=%d\n", x, y)
    // Deve imprimir: Depois: x=20, y=10
}
```

**Dica**: Você vai precisar de uma variável temporária!

---

## 📚 Próximos Passos

Depois de completar os exercícios e reflexões, você estará pronto para a próxima aula sobre **Performance e Boas Práticas** com pointers e memory management!

**Lembre-se**: O objetivo não é apenas fazer os exercícios funcionarem, mas **entender o porquê** de cada conceito. As perguntas de reflexão são tão importantes quanto os exercícios práticos!

Boa sorte! 🚀

