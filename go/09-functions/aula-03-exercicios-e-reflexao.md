# Módulo 9: Functions em Go

## Aula 3 - Exercícios e Reflexão

Agora é hora de colocar a mão na massa! Complete os exercícios abaixo e reflita sobre as questões propostas.

---

## Exercícios Práticos

### Exercício 1: Funções Básicas

Crie um arquivo `exercicio1.go` e implemente o seguinte:

**Tarefa:**

1. Crie uma função `calcularAreaRetangulo(largura, altura float64) float64` que calcula e retorna a área de um retângulo.

2. Crie uma função `calcularPerimetroRetangulo(largura, altura float64) float64` que calcula e retorna o perímetro.

3. Crie uma função `calcularIMC(peso, altura float64) float64` que calcula o Índice de Massa Corporal (IMC = peso / altura²).

4. No `main`, teste todas as funções com valores de exemplo e exiba os resultados.

**Exemplo de uso esperado:**

```go
area := calcularAreaRetangulo(5.0, 3.0)
perimetro := calcularPerimetroRetangulo(5.0, 3.0)
imc := calcularIMC(70.0, 1.75)

fmt.Printf("Área: %.2f\n", area)        // 15.00
fmt.Printf("Perímetro: %.2f\n", perimetro) // 16.00
fmt.Printf("IMC: %.2f\n", imc)         // ~22.86
```

---

### Exercício 2: Múltiplos Retornos e Tratamento de Erros

Crie um arquivo `exercicio2.go` e implemente:

**Tarefa:**

1. Crie uma função `dividir(a, b float64) (float64, error)` que:

   - Retorna o resultado da divisão e `nil` se `b != 0`
   - Retorna `0` e um erro se `b == 0`

2. Crie uma função `calcularRaizes(a, b, c float64) (x1, x2 float64, temSolucao bool)` que:

   - Calcula as raízes de uma equação quadrática (ax² + bx + c = 0)
   - Retorna as raízes e `true` se houver solução real
   - Retorna `0, 0, false` se não houver solução real (discriminante negativo)

3. Crie uma função `calcularEstatisticas(numeros []float64) (media, max, min float64)` que retorna a média, máximo e mínimo de um slice de números.

4. No `main`, teste todas as funções com diferentes casos (incluindo casos de erro).

**Exemplo de uso:**

```go
resultado, err := dividir(10, 2)
if err != nil {
    fmt.Printf("Erro: %v\n", err)
} else {
    fmt.Printf("Resultado: %.2f\n", resultado) // 5.00
}

x1, x2, ok := calcularRaizes(1, -5, 6)
if ok {
    fmt.Printf("Raízes: x1=%.2f, x2=%.2f\n", x1, x2) // x1=3.00, x2=2.00
}
```

---

### Exercício 3: Funções Variádicas

Crie um arquivo `exercicio3.go` e implemente:

**Tarefa:**

1. Crie uma função `somarNumeros(numeros ...int) int` que soma todos os números passados.

2. Crie uma função `encontrarMaximo(numeros ...int) int` que encontra o maior número.

3. Crie uma função `juntarStrings(separador string, textos ...string) string` que junta todas as strings com o separador.

4. Crie uma função `filtrarPares(numeros ...int) []int` que retorna apenas os números pares.

5. No `main`, teste todas as funções com diferentes números de argumentos.

**Exemplo de uso:**

```go
soma := somarNumeros(1, 2, 3, 4, 5)  // 15
max := encontrarMaximo(10, 5, 8, 20, 3)  // 20
texto := juntarStrings(", ", "a", "b", "c")  // "a, b, c"
pares := filtrarPares(1, 2, 3, 4, 5, 6)  // [2 4 6]
```

---

### Exercício 4: Funções Anônimas e Closures

Crie um arquivo `exercicio4.go` e implemente:

**Tarefa:**

1. Crie uma função `aplicarTransformacao(numeros []int, transformacao func(int) int) []int` que aplica uma função de transformação a cada elemento do slice.

2. Use funções anônimas para criar transformações:

   - Dobrar cada número
   - Elevar ao quadrado
   - Adicionar 10

3. Crie uma função `criarContador() func() int` que retorna uma closure que mantém um contador interno.

4. Crie uma função `criarMultiplicador(fator int) func(int) int` que retorna uma closure que multiplica qualquer número pelo fator capturado.

5. No `main`, teste todas as funções e closures.

**Exemplo de uso:**

```go
numeros := []int{1, 2, 3, 4, 5}

// Dobrar usando função anônima
dobrados := aplicarTransformacao(numeros, func(n int) int {
    return n * 2
})
// Resultado: [2 4 6 8 10]

// Contador
contador := criarContador()
fmt.Println(contador()) // 1
fmt.Println(contador()) // 2

// Multiplicador
dobrar := criarMultiplicador(2)
fmt.Println(dobrar(5)) // 10
```

---

### Exercício 5: First-Class Functions

Crie um arquivo `exercicio5.go` e implemente:

**Tarefa:**

1. Crie funções de operação matemática:

   - `somar(a, b int) int`
   - `subtrair(a, b int) int`
   - `multiplicar(a, b int) int`

2. Crie uma função `aplicarOperacao(a, b int, operacao func(int, int) int) int` que recebe uma função de operação e a aplica.

3. Crie uma função `criarOperacao(tipo string) func(int, int) int` que retorna a função de operação baseada no tipo ("soma", "subtracao", "multiplicacao").

4. No `main`, demonstre:
   - Passar funções como argumentos
   - Atribuir funções a variáveis
   - Retornar funções de outras funções

**Exemplo de uso:**

```go
resultado1 := aplicarOperacao(10, 5, somar)  // 15
resultado2 := aplicarOperacao(10, 5, subtrair)  // 5

operacao := criarOperacao("multiplicacao")
resultado3 := operacao(10, 5)  // 50
```

---

### Exercício 6: Call by Value - Demonstração Prática

Crie um arquivo `exercicio6.go` e implemente:

**Tarefa:**

1. Crie uma struct `Pessoa` com campos `Nome` e `Idade`.

2. Crie uma função `modificarPessoa(p Pessoa)` que tenta modificar os campos da pessoa.

3. Crie uma função `modificarPessoaPonteiro(p *Pessoa)` que modifica usando ponteiro.

4. Crie uma função `modificarSlice(s []int)` que modifica elementos do slice.

5. No `main`, demonstre:
   - Passagem por valor de struct (não modifica original)
   - Passagem por ponteiro de struct (modifica original)
   - Passagem de slice (modifica elementos, mas append não afeta original)

**Exemplo de uso:**

```go
pessoa := Pessoa{Nome: "João", Idade: 30}
modificarPessoa(pessoa)
fmt.Println(pessoa)  // {João 30} - não mudou

modificarPessoaPonteiro(&pessoa)
fmt.Println(pessoa)  // {Maria 25} - mudou!

slice := []int{1, 2, 3}
modificarSlice(slice)
fmt.Println(slice)  // [999 2 3] - elementos modificados
```

---

## Perguntas de Reflexão

### Reflexão 1: Por Que Funções são First-Class Citizens?

Você aprendeu que funções em Go são first-class citizens (podem ser atribuídas, passadas, retornadas).

**Pergunta:**

1. Qual é a vantagem prática de funções serem first-class citizens?
2. Dê um exemplo real de quando você usaria uma função como argumento para outra função.
3. Como isso facilita a criação de código reutilizável e flexível?

Escreva sua resposta com suas próprias palavras (mínimo 4 linhas).

---

### Reflexão 2: Múltiplos Retornos e Tratamento de Erros

Go usa o padrão de retornar `(resultado, error)` em vez de lançar exceções.

**Pergunta:**

1. Por que você acha que Go escolheu esse padrão ao invés de exceções?
2. Quais são as vantagens e desvantagens desse padrão?
3. Em que situações você sempre deve verificar o erro retornado?
4. Quando é aceitável ignorar um erro usando `_`?

Escreva sua resposta com suas próprias palavras (mínimo 5 linhas).

---

### Reflexão 3: Closures e Estado

Closures permitem que funções "lembrem" variáveis do escopo externo.

**Pergunta:**

1. Explique com suas próprias palavras o que é uma closure e como ela funciona.
2. Dê um exemplo prático de quando você usaria uma closure no mundo real (não código, mas situação real).
3. Por que closures podem ser perigosas em loops? Explique o problema e a solução.
4. Quando você escolheria usar uma closure ao invés de uma struct com métodos?

Escreva sua resposta com suas próprias palavras (mínimo 6 linhas).

---

### Reflexão 4: Call by Value vs Referência

Go passa tudo por valor, mas slices, maps e channels se comportam como referências.

**Pergunta:**

1. Por que Go escolheu passar tudo por valor ao invés de por referência?
2. Explique por que modificar elementos de um slice dentro de uma função afeta o original, mas `append` não.
3. Quando você deve usar ponteiros (`*T`) ao invés de valores (`T`) para parâmetros?
4. Qual é o custo de performance de passar structs grandes por valor?

Escreva sua resposta com suas próprias palavras (mínimo 6 linhas).

---

### Reflexão 5: Named Return Values - Quando Usar?

Retornos nomeados podem melhorar legibilidade, mas também podem mascarar problemas.

**Pergunta:**

1. Em que situações named returns melhoram a legibilidade do código?
2. Quais são os perigos de usar named returns (quando pode causar problemas)?
3. Dê um exemplo de quando você usaria named returns e quando não usaria.
4. Por que alguns desenvolvedores Go evitam named returns?

Escreva sua resposta com suas próprias palavras (mínimo 5 linhas).

---

### Reflexão 6: Variadic Functions - Aplicação Prática

Funções variádicas permitem flexibilidade no número de argumentos.

**Pergunta:**

1. Dê três exemplos de funções da biblioteca padrão do Go que usam variadic functions.
2. Quando você criaria uma função variádica ao invés de receber um slice?
3. Qual é a diferença prática entre `func f(args ...int)` e `func f(args []int)`?
4. Em que situações variadic functions podem ser confusas ou problemáticas?

Escreva sua resposta com suas próprias palavras (mínimo 5 linhas).

---

## Como Entregar

Crie arquivos `.go` separados para cada exercício (ex: `exercicio1.go`, `exercicio2.go`, etc.) na pasta `09-functions/`. Para as perguntas de reflexão, você pode criar um arquivo `reflexoes.md` ou simplesmente responder diretamente aqui.

**Importante:**

- Compile e execute cada programa para garantir que funciona
- Comente seu código explicando o que cada parte faz
- Seja honesto nas reflexões - não há resposta "errada", mas há respostas que demonstram mais ou menos compreensão
- Teste casos extremos (divisão por zero, slices vazios, etc.)

---

## Dicas para os Exercícios

### Dica 1: Tratamento de Erros

Sempre verifique erros retornados:

```go
resultado, err := dividir(a, b)
if err != nil {
    // Tratar erro
    return
}
// Usar resultado
```

### Dica 2: Closures em Loops

Lembre-se da armadilha:

```go
// ERRADO
for i := 0; i < 3; i++ {
    funcoes = append(funcoes, func() int { return i })
}

// CORRETO
for i := 0; i < 3; i++ {
    i := i  // Nova variável
    funcoes = append(funcoes, func() int { return i })
}
```

### Dica 3: Variadic com Slice

Use `...` para passar slice:

```go
numeros := []int{1, 2, 3}
soma := somarNumeros(numeros...)  // Não esqueça do ...
```

---

Após completar os exercícios e reflexões, envie suas respostas para que eu possa analisar seu desempenho e fornecer feedback construtivo!

