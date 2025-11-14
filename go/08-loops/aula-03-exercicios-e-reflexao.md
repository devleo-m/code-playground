# Módulo 8: Loops em Go

## Aula 3 - Exercícios e Reflexão

Agora é hora de colocar a mão na massa! Complete os exercícios abaixo e reflita sobre as questões propostas.

---

## Exercícios Práticos

### Exercício 1: Contador e Tabela de Multiplicação

Crie um arquivo `exercicio1.go` e implemente o seguinte:

**Tarefa:**

1. Crie uma função `ContarAte(n int)` que imprime números de 1 até `n` usando um loop `for` clássico.

2. Crie uma função `TabelaMultiplicacao(n int)` que imprime a tabela de multiplicação de 1 até `n` usando loops aninhados.

3. No `main`, teste ambas as funções.

**Exemplo de uso esperado:**

```go
ContarAte(5)
// Saída:
// 1
// 2
// 3
// 4
// 5

TabelaMultiplicacao(3)
// Saída:
// 1 x 1 = 1
// 1 x 2 = 2
// 1 x 3 = 3
// 2 x 1 = 2
// 2 x 2 = 4
// 2 x 3 = 6
// 3 x 1 = 3
// 3 x 2 = 6
// 3 x 3 = 9
```

---

### Exercício 2: Processar Lista de Números

Crie um arquivo `exercicio2.go` e implemente:

**Tarefa:**

1. Crie uma função `SomarNumeros(numeros []int) int` que soma todos os números de um slice usando `for range`.

2. Crie uma função `EncontrarMaximo(numeros []int) int` que encontra o maior número do slice usando `for range`.

3. Crie uma função `FiltrarPares(numeros []int) []int` que retorna apenas os números pares usando `for range` e `append`.

4. Crie uma função `ContarOcorrencias(numeros []int, alvo int) int` que conta quantas vezes `alvo` aparece no slice.

5. No `main`, teste todas as funções com um slice de exemplo: `[]int{1, 2, 3, 4, 5, 2, 6, 2, 7, 8}`

**Exemplo de uso:**

```go
numeros := []int{1, 2, 3, 4, 5, 2, 6, 2, 7, 8}
fmt.Println(SomarNumeros(numeros))        // 40
fmt.Println(EncontrarMaximo(numeros))     // 8
fmt.Println(FiltrarPares(numeros))        // [2 4 6 2 8]
fmt.Println(ContarOcorrencias(numeros, 2)) // 3
```

---

### Exercício 3: Iterar sobre Map e String

Crie um arquivo `exercicio3.go` e implemente:

**Tarefa:**

1. Crie uma função `ListarCores(cores map[string]string)` que imprime todas as cores e seus códigos hex usando `for range`.

2. Crie uma função `SoletrarPalavra(palavra string)` que imprime cada caractere (rune) da palavra usando `for range`.

3. Crie uma função `ContarCaracteres(palavra string) int` que conta quantos caracteres (runes) a palavra tem. **Dica**: Use `for range` e não `len()` diretamente na string.

4. No `main`, teste com:
   - Map: `map[string]string{"vermelho": "#FF0000", "verde": "#00FF00", "azul": "#0000FF"}`
   - String: `"Olá, 世界! 🚀"`

**Exemplo de uso:**

```go
cores := map[string]string{"vermelho": "#FF0000", "verde": "#00FF00"}
ListarCores(cores)
// Saída (ordem pode variar):
// vermelho: #FF0000
// verde: #00FF00

SoletrarPalavra("Café")
// Saída:
// Posição 0: C
// Posição 1: a
// Posição 2: f
// Posição 3: é

fmt.Println(ContarCaracteres("Olá, 世界! 🚀")) // 10 (não 15 bytes!)
```

---

### Exercício 4: Buscar e Filtrar com break e continue

Crie um arquivo `exercicio4.go` e implemente:

**Tarefa:**

1. Crie uma função `BuscarNome(nomes []string, alvo string) (int, bool)` que:

   - Busca `alvo` no slice de nomes
   - Retorna o índice e `true` se encontrar
   - Retorna `-1` e `false` se não encontrar
   - Use `for range` e `break` para parar quando encontrar

2. Crie uma função `FiltrarPositivos(numeros []int) []int` que:

   - Retorna apenas números positivos (> 0)
   - Use `for range` e `continue` para pular números negativos ou zero

3. Crie uma função `ProcessarAtePrimeiroNegativo(numeros []int) int` que:

   - Soma números até encontrar o primeiro negativo
   - Para imediatamente quando encontrar negativo (use `break`)
   - Retorna a soma acumulada

4. No `main`, teste todas as funções.

**Exemplo de uso:**

```go
nomes := []string{"João", "Maria", "Pedro"}
indice, encontrado := BuscarNome(nomes, "Maria")
// indice = 1, encontrado = true

numeros := []int{-1, 2, 3, -4, 5}
positivos := FiltrarPositivos(numeros)
// positivos = [2 3 5]

soma := ProcessarAtePrimeiroNegativo([]int{1, 2, 3, -1, 4, 5})
// soma = 6 (para no -1)
```

---

### Exercício 5: Loops Aninhados e Labels

Crie um arquivo `exercicio5.go` e implemente:

**Tarefa:**

1. Crie uma função `BuscarEmMatriz(matriz [][]int, alvo int) (int, int, bool)` que:

   - Busca `alvo` em uma matriz 2D
   - Retorna linha, coluna e `true` se encontrar
   - Retorna `-1, -1, false` se não encontrar
   - Use loops aninhados e `break` com label para sair de ambos os loops quando encontrar

2. Crie uma função `ImprimirTabuleiro(tamanho int)` que:

   - Imprime um tabuleiro de xadrez (alternando caracteres)
   - Use loops aninhados
   - Exemplo para tamanho 3:
     ```
     # . #
     . # .
     # . #
     ```

3. No `main`, teste ambas as funções.

**Exemplo de uso:**

```go
matriz := [][]int{
    {1, 2, 3},
    {4, 5, 6},
    {7, 8, 9},
}
linha, coluna, encontrado := BuscarEmMatriz(matriz, 5)
// linha = 1, coluna = 1, encontrado = true

ImprimirTabuleiro(4)
// Imprime tabuleiro 4x4
```

---

## Perguntas de Reflexão

### Reflexão 1: Por Que Go Tem Apenas `for`?

Você aprendeu que Go tem apenas um tipo de loop (`for`), enquanto outras linguagens têm `while`, `do-while`, `foreach`, etc.

**Pergunta:**

1. Por que você acha que os criadores do Go escolheram ter apenas `for`?
2. Quais são as vantagens de ter apenas um tipo de loop?
3. Quais são as desvantagens (se houver)?
4. Como a flexibilidade do `for` em Go compensa a falta de outros tipos de loops?

**Sua resposta deve incluir**: Análise das vantagens e desvantagens, e sua opinião sobre se essa escolha de design foi boa ou não (mínimo 4 linhas).

---

### Reflexão 2: `for range` vs Loop Tradicional

Você aprendeu que `for range` é preferido para iterar sobre coleções em Go.

**Pergunta:**

1. Por que `for range` é considerado mais idiomático em Go?
2. Em que situações você ainda usaria um loop `for` tradicional ao invés de `for range`?
3. Qual é a diferença prática de performance entre `for range` e loop tradicional para slices?
4. Por que `for range` é mais seguro (menos propenso a erros)?

**Sua resposta deve demonstrar**: Compreensão profunda das diferenças e quando usar cada um (mínimo 5 linhas).

---

### Reflexão 3: Iteração sobre Strings e Unicode

Você aprendeu que `for range` sobre strings retorna runes, não bytes, e que indexação direta retorna bytes.

**Pergunta:**

1. Por que Go fez essa escolha de design (retornar runes no `for range`)?
2. Explique com suas próprias palavras por que indexação direta em strings pode causar problemas com caracteres Unicode.
3. Dê um exemplo real de quando isso seria um problema em um programa (ex: contar caracteres, buscar substring).
4. Quando você converteria uma string para `[]rune`? Dê um exemplo prático.

**Sua resposta deve incluir**: Explicação clara do problema Unicode e exemplos práticos (mínimo 5 linhas).

---

### Reflexão 4: Modificar Coleções Durante Iteração

Você aprendeu que modificar elementos durante `for range` é seguro, mas adicionar/remover pode causar problemas.

**Pergunta:**

1. Por que modificar elementos existentes é seguro, mas adicionar/remover não é?
2. O que acontece internamente quando você adiciona um elemento a um slice durante iteração `for range`?
3. Como você resolveria o problema de precisar adicionar/remover elementos durante iteração?
4. Por que deletar de um map durante iteração é seguro, mas adicionar/modificar não é?

**Sua resposta deve demonstrar**: Compreensão do comportamento interno e como evitar problemas (mínimo 5 linhas).

---

### Reflexão 5: `break`, `continue` e Labels

Você aprendeu sobre `break`, `continue` e labels para controlar loops aninhados.

**Pergunta:**

1. Quando você usaria `break` ao invés de `continue`? Dê um exemplo prático de cada um.
2. Em que situação labels são realmente necessários? Dê um exemplo onde labels são a melhor solução.
3. Por que alguns programadores evitam labels? Quando labels tornam código difícil de ler?
4. Como você refatoraria código que usa labels para torná-lo mais legível (sem usar labels)?

**Sua resposta deve incluir**: Exemplos práticos e análise de quando labels são apropriados (mínimo 5 linhas).

---

### Reflexão 6: Aplicação Real - Sistema de Processamento

Imagine que você está desenvolvendo um sistema que processa uma lista de transações bancárias.

**Pergunta:**

1. Como você usaria loops para processar todas as transações?
2. Você usaria `for range` ou loop tradicional? Por quê?
3. Como você implementaria busca de uma transação específica (usando `break`)?
4. Como você filtraria apenas transações acima de um valor (usando `continue`)?
5. Se precisasse processar transações em lotes (grupos de 10), como faria?

**Sua resposta deve incluir**: Código pseudocódigo ou Go mostrando suas escolhas e justificativas (mínimo 6 linhas).

---

## Como Entregar

Crie arquivos `.go` separados para cada exercício (ex: `exercicio1.go`, `exercicio2.go`, etc.) na pasta `go/08-loops/`. Para as perguntas de reflexão, você pode criar um arquivo `reflexoes.md` ou simplesmente responder diretamente aqui.

**Importante**:

- Compile e execute cada programa para garantir que funciona
- Comente seu código explicando o que cada parte faz
- Seja honesto nas reflexões - não há resposta "errada", mas há respostas que demonstram mais ou menos compreensão
- Teste casos extremos (slices vazios, valores negativos, etc.)

---

Após completar os exercícios e reflexões, envie suas respostas para que eu possa analisar seu desempenho e fornecer feedback construtivo!
