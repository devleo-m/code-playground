# Módulo 8: Loops em Go

## Aula 1: Loops - Repetição e Iteração

Olá! Bem-vindo ao oitavo módulo. Até agora você aprendeu a declarar variáveis, criar estruturas de dados, tomar decisões com conditionals. Mas e se você precisar executar o mesmo código **múltiplas vezes**? E se precisar processar **todos os elementos** de uma lista?

É aqui que entram os **loops** (laços) - estruturas que permitem repetir código de forma controlada. Em Go, há apenas **um tipo de loop**: o `for`, mas ele é extremamente flexível e pode ser usado de várias formas diferentes.

Nesta aula, vamos mergulhar profundamente em todas as formas de usar `for` em Go, incluindo `for range` para iterar sobre coleções.

---

## 1. O Que São Loops?

**Loops** (laços) são estruturas que permitem executar um bloco de código **múltiplas vezes** de forma controlada. Eles são essenciais para:

- Processar listas de dados
- Repetir operações até uma condição ser atendida
- Iterar sobre arrays, slices, maps e strings
- Criar algoritmos que precisam de repetição

### Por Que Precisamos de Loops?

**Sem loops**, você teria que repetir código manualmente:

```go
func main() {
    fmt.Println("Número 1")
    fmt.Println("Número 2")
    fmt.Println("Número 3")
    fmt.Println("Número 4")
    fmt.Println("Número 5")
    // E assim por diante... muito trabalhoso!
}
```

**Com loops**, você pode fazer isso automaticamente:

```go
func main() {
    for i := 1; i <= 5; i++ {
        fmt.Printf("Número %d\n", i)
    }
}
```

**Vantagens:**

- ✅ Código conciso - escreva uma vez, execute muitas vezes
- ✅ Processamento de coleções - itere sobre arrays, slices, maps
- ✅ Algoritmos eficientes - implemente lógica complexa
- ✅ Flexibilidade - controle quando parar ou pular iterações

---

## 2. O `for` Loop Clássico

O `for` loop clássico é a forma mais tradicional e completa. Ele tem três componentes principais:

### Sintaxe Completa

```go
for inicializacao; condicao; posInstrucao {
    // código a ser repetido
}
```

**Componentes:**

1. **Inicialização**: Executada **uma vez** antes do loop começar
2. **Condição**: Verificada **antes de cada iteração**. Se `true`, executa o bloco
3. **Pós-instrução**: Executada **após cada iteração**, antes de verificar a condição novamente

### Exemplo Básico

```go
for i := 0; i < 5; i++ {
    fmt.Printf("Iteração %d\n", i)
}
```

**O que acontece:**

1. `i := 0` - Inicializa `i` com 0 (executa uma vez)
2. `i < 5` - Verifica se `i` é menor que 5
3. Se verdadeiro → executa `fmt.Printf`
4. `i++` - Incrementa `i` em 1
5. Volta para o passo 2
6. Quando `i >= 5`, o loop termina

**Saída:**

```
Iteração 0
Iteração 1
Iteração 2
Iteração 3
Iteração 4
```

### Características Importantes

- **Variável de loop**: A variável declarada na inicialização existe apenas dentro do escopo do loop
- **Condição**: Deve resultar em `bool` (true/false)
- **Pós-instrução**: Pode ser qualquer expressão (incremento, decremento, múltiplas operações)

---

## 3. `for` como `while` (Apenas Condição)

Go não tem uma palavra-chave `while` separada. Em vez disso, você usa `for` com apenas a condição, criando um loop estilo `while`.

### Sintaxe

```go
for condicao {
    // código
}
```

### Exemplo

```go
contador := 0
for contador < 5 {
    fmt.Printf("Contador: %d\n", contador)
    contador++
}
```

**Quando usar:**

- Quando você não sabe quantas iterações serão necessárias
- Quando a condição depende de algo que muda dentro do loop
- Quando você precisa de mais controle sobre quando parar

### Comparação: `for` Clássico vs `while`-style

```go
// for clássico - quando você sabe quantas iterações
for i := 0; i < 10; i++ {
    // ...
}

// for while-style - quando a condição é mais complexa
for condicaoComplexa() {
    // ...
}
```

---

## 4. `for` Loop Infinito

Um loop infinito executa **indefinidamente** até ser interrompido explicitamente com `break` ou `return`.

### Sintaxe

```go
for {
    // código que executa infinitamente
    // Precisa de break ou return para sair
}
```

### Exemplo com `break`

```go
contador := 0
for {
    contador++
    if contador >= 5 {
        break // Sair do loop
    }
    fmt.Printf("Iteração %d\n", contador)
}
```

**Quando usar:**

- Servidores que rodam continuamente
- Processamento de eventos
- Loops que dependem de condições internas complexas

**⚠️ Cuidado:** Sempre tenha uma forma de sair do loop, ou seu programa ficará travado!

---

## 5. `for range` - Iteração sobre Coleções

O `for range` é uma forma especial e idiomática de iterar sobre arrays, slices, maps, strings e channels em Go. É a forma **preferida** para iterar sobre coleções.

### Sintaxe Básica

```go
for indice, valor := range colecao {
    // código
}
```

**Retornos:**

- **Arrays/Slices**: `(índice, valor)`
- **Maps**: `(chave, valor)`
- **Strings**: `(índice, rune)` - ⚠️ Importante: retorna runes, não bytes!
- **Channels**: `(valor)` - apenas valor

### `for range` com Arrays e Slices

```go
numeros := []int{10, 20, 30, 40, 50}

// Com índice e valor
for indice, valor := range numeros {
    fmt.Printf("Índice %d: valor %d\n", indice, valor)
}

// Apenas valores (ignorar índice)
for _, valor := range numeros {
    fmt.Printf("Valor: %d\n", valor)
}

// Apenas índices (ignorar valor)
for indice := range numeros {
    fmt.Printf("Índice: %d\n", indice)
}
```

**Vantagens:**

- ✅ Mais legível que loop tradicional
- ✅ Menos propenso a erros (não precisa usar `len()`)
- ✅ Idiomático em Go
- ✅ Funciona com qualquer tamanho de coleção

### `for range` com Maps

```go
cores := map[string]string{
    "vermelho": "#FF0000",
    "verde":    "#00FF00",
    "azul":     "#0000FF",
}

// Iterar sobre map
for chave, valor := range cores {
    fmt.Printf("Cor: %s = %s\n", chave, valor)
}
```

**⚠️ Importante sobre Maps:**

1. **Ordem aleatória**: A iteração sobre maps é **intencionalmente aleatória** em Go. Isso é uma feature de segurança para prevenir bugs que dependem da ordem.
2. **Não pode modificar**: Você **não pode adicionar ou modificar** elementos do map durante a iteração `range`. Mas pode **deletar** com segurança.
3. **Apenas chaves ou valores**: Pode ignorar chave ou valor usando `_`

```go
// Apenas chaves
for chave := range cores {
    fmt.Println(chave)
}

// Apenas valores
for _, valor := range cores {
    fmt.Println(valor)
}
```

### `for range` com Strings

**CRUCIAL**: `for range` sobre strings retorna **runes** (pontos de código Unicode), não bytes!

```go
texto := "Olá, 世界! 🚀"

// for range retorna (índice, rune)
for indice, rune := range texto {
    fmt.Printf("Posição %d: %c (Unicode: %d)\n", indice, rune, rune)
}
```

**Por que isso importa?**

- Strings em Go são sequências de **bytes**, não caracteres
- Caracteres Unicode podem ocupar múltiplos bytes
- `for range` processa corretamente caracteres multibyte
- Indexação direta `str[i]` retorna **bytes**, não caracteres!

**Comparação:**

```go
texto := "Café"

// ERRADO: Indexação direta (bytes)
for i := 0; i < len(texto); i++ {
    fmt.Printf("texto[%d] = %d (byte)\n", i, texto[i])
}
// Saída: bytes individuais (pode quebrar caracteres multibyte)

// CORRETO: for range (runes)
for i, r := range texto {
    fmt.Printf("Posição %d: %c (rune)\n", i, r)
}
// Saída: caracteres completos
```

**Para acesso aleatório a caracteres:**

Se você precisa acessar caracteres por índice (não sequencialmente), converta para `[]rune` primeiro:

```go
texto := "Olá, 世界!"
runes := []rune(texto)

// Agora pode acessar por índice
fmt.Printf("Primeiro caractere: %c\n", runes[0])
fmt.Printf("Último caractere: %c\n", runes[len(runes)-1])
```

---

## 6. `break` - Sair do Loop

A palavra-chave `break` **sai imediatamente** do loop mais interno (ou do loop rotulado).

### Sintaxe

```go
for condicao {
    if algumaCondicao {
        break // Sai do loop
    }
    // código
}
```

### Exemplo

```go
for i := 0; i < 10; i++ {
    if i == 5 {
        fmt.Println("Encontrei 5! Saindo...")
        break
    }
    fmt.Printf("Valor: %d\n", i)
}
// Saída: 0, 1, 2, 3, 4, "Encontrei 5! Saindo..."
```

**Quando usar:**

- Buscar elemento em coleção (sair quando encontrar)
- Validar dados (sair quando encontrar erro)
- Processar até condição específica

---

## 7. `continue` - Pular Iteração

A palavra-chave `continue` **pula o resto da iteração atual** e vai para a próxima iteração do loop.

### Sintaxe

```go
for condicao {
    if algumaCondicao {
        continue // Pula para próxima iteração
    }
    // código que não executa se continue foi chamado
}
```

### Exemplo

```go
for i := 0; i < 10; i++ {
    if i%2 == 0 {
        continue // Pula números pares
    }
    fmt.Printf("Número ímpar: %d\n", i)
}
// Saída: 1, 3, 5, 7, 9
```

**Quando usar:**

- Filtrar elementos (pular elementos indesejados)
- Tratar casos especiais no início da iteração
- Evitar aninhamento de `if-else`

---

## 8. Labels e Loops Aninhados

Em loops aninhados, `break` e `continue` afetam apenas o **loop mais interno**. Para controlar loops externos, use **labels** (rótulos).

### Sintaxe de Label

```go
LabelNome:
    for condicao1 {
        for condicao2 {
            if algumaCondicao {
                break LabelNome // Sai do loop externo
            }
        }
    }
```

### Exemplo: `break` com Label

```go
LoopExterno:
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            if j == 1 {
                break LoopExterno // Sai do loop externo!
            }
            fmt.Printf("i=%d, j=%d\n", i, j)
        }
    }
```

### Exemplo: `continue` com Label

```go
LoopExterno:
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            if j == 1 {
                continue LoopExterno // Pula para próxima iteração do loop externo
            }
            fmt.Printf("i=%d, j=%d\n", i, j)
        }
    }
```

**Quando usar labels:**

- Loops aninhados complexos
- Quando precisa sair de múltiplos níveis
- Algoritmos de busca em estruturas 2D

**⚠️ Use com moderação:** Labels podem tornar código menos legível. Prefira refatorar em funções quando possível.

---

## 9. Loops Aninhados

Loops aninhados são loops dentro de outros loops. Úteis para processar estruturas multidimensionais.

### Exemplo: Tabela de Multiplicação

```go
for i := 1; i <= 3; i++ {
    for j := 1; j <= 3; j++ {
        fmt.Printf("%d x %d = %d\n", i, j, i*j)
    }
    fmt.Println()
}
```

**Complexidade:**

- Loop externo: `n` iterações
- Loop interno: `m` iterações
- Total: `n × m` iterações

**Quando usar:**

- Matrizes 2D
- Comparações entre elementos
- Algoritmos de busca/análise

---

## 10. Modificando Coleções Durante Iteração

### Slices e Arrays

**✅ SEGURO**: Modificar **elementos existentes** durante iteração `range`:

```go
numeros := []int{1, 2, 3, 4, 5}
for i := range numeros {
    numeros[i] *= 2 // Modificar elemento é seguro
}
```

**⚠️ CUIDADO**: Adicionar ou remover elementos durante `range` pode causar comportamento inesperado. Use loop tradicional com índice:

```go
numeros := []int{1, 2, 3, 4, 5}
for i := 0; i < len(numeros); i++ {
    if numeros[i]%2 == 0 {
        // Remover elemento
        numeros = append(numeros[:i], numeros[i+1:]...)
        i-- // Ajustar índice
    }
}
```

### Maps

**❌ NÃO PODE**: Adicionar ou modificar elementos durante iteração `range`:

```go
cores := map[string]int{"vermelho": 1, "verde": 2}

// ERRADO - comportamento indefinido!
for chave, valor := range cores {
    cores[chave] = valor * 2 // Pode causar problemas
}
```

**✅ PODE**: Deletar elementos durante iteração:

```go
cores := map[string]int{"vermelho": 1, "verde": 2, "azul": 3}

// CORRETO - deletar é seguro
for chave, valor := range cores {
    if valor == 2 {
        delete(cores, chave) // Deletar é seguro
    }
}
```

**Solução para modificar map**: Criar novo map:

```go
cores := map[string]int{"vermelho": 1, "verde": 2}
novasCores := make(map[string]int)

for chave, valor := range cores {
    novasCores[chave] = valor * 2
}
```

---

## 11. Padrões Comuns com Loops

### Padrão 1: Buscar Elemento

```go
func buscar(nomes []string, alvo string) (int, bool) {
    for i, nome := range nomes {
        if nome == alvo {
            return i, true
        }
    }
    return -1, false
}
```

### Padrão 2: Filtrar Elementos

```go
func filtrarPares(numeros []int) []int {
    pares := []int{}
    for _, num := range numeros {
        if num%2 == 0 {
            pares = append(pares, num)
        }
    }
    return pares
}
```

### Padrão 3: Soma/Acumulação

```go
func somar(numeros []int) int {
    soma := 0
    for _, num := range numeros {
        soma += num
    }
    return soma
}
```

### Padrão 4: Contar Ocorrências

```go
func contar(numeros []int, alvo int) int {
    contador := 0
    for _, num := range numeros {
        if num == alvo {
            contador++
        }
    }
    return contador
}
```

### Padrão 5: Encontrar Máximo/Mínimo

```go
func encontrarMaximo(numeros []int) int {
    if len(numeros) == 0 {
        return 0
    }
    maximo := numeros[0]
    for _, num := range numeros[1:] {
        if num > maximo {
            maximo = num
        }
    }
    return maximo
}
```

---

## 12. `goto` (Desencorajado)

Go inclui a palavra-chave `goto`, mas seu uso é **fortemente desencorajado** pela comunidade Go.

### Sintaxe

```go
Label:
    // código
    goto Label
```

### Por Que é Desencorajado?

1. **Código não estruturado**: Dificulta leitura e manutenção
2. **Debugging difícil**: Fluxo de execução imprevisível
3. **Alternativas melhores**: `break`, `continue`, funções, `return`
4. **Não é idiomático**: Vai contra a filosofia Go de código simples e claro

### Quando (Raramente) Usar?

Apenas em casos muito específicos:

- Tratamento de erros complexo em código de baixo nível
- Otimizações de performance críticas
- Saída de loops profundamente aninhados (mas prefira refatorar)

**Recomendação**: **NUNCA** use `goto` a menos que seja absolutamente necessário e você entenda completamente as implicações.

---

## 13. Resumo dos Conceitos-Chave

| Conceito          | Descrição                      | Quando Usar                        |
| ----------------- | ------------------------------ | ---------------------------------- |
| `for` clássico    | `for init; cond; post {}`      | Quando sabe quantas iterações      |
| `for` while-style | `for cond {}`                  | Quando condição é complexa         |
| `for` infinito    | `for {}`                       | Servidores, processamento contínuo |
| `for range`       | `for i, v := range colecao {}` | Iterar sobre coleções (PREFERIDO)  |
| `break`           | Sair do loop                   | Quando encontrou o que procurava   |
| `continue`        | Pular iteração                 | Filtrar elementos                  |
| Labels            | `Label: for {}`                | Controlar loops aninhados          |
| `goto`            | `goto Label`                   | **EVITAR** - raramente necessário  |

---

## 14. Boas Práticas

### ✅ Use `for range` para Coleções

```go
// BOM
for i, valor := range numeros {
    // ...
}

// EVITE (a menos que precise de controle especial)
for i := 0; i < len(numeros); i++ {
    // ...
}
```

### ✅ Use `break` para Buscar

```go
// BOM
for i, item := range items {
    if item == alvo {
        encontrado = i
        break
    }
}
```

### ✅ Use `continue` para Filtrar

```go
// BOM
for _, num := range numeros {
    if num < 0 {
        continue // Pula negativos
    }
    processar(num)
}
```

### ❌ Evite Modificar Tamanho Durante `range`

```go
// EVITE
for _, item := range items {
    if condicao {
        items = append(items, novoItem) // Comportamento indefinido
    }
}
```

---

## Conclusão

Loops são fundamentais em programação. Em Go, o `for` é extremamente flexível e pode ser usado de várias formas:

- **`for` clássico**: Para iterações com contador
- **`for` while-style**: Para condições complexas
- **`for range`**: Para iterar sobre coleções (PREFERIDO)
- **`break`/`continue`**: Para controle de fluxo
- **Labels**: Para loops aninhados complexos

Lembre-se:

- ✅ Prefira `for range` para coleções
- ✅ Use `break` para sair quando encontrar
- ✅ Use `continue` para filtrar elementos
- ❌ Evite `goto`
- ⚠️ Cuidado ao modificar coleções durante iteração

Na próxima parte desta aula, vamos simplificar esses conceitos com analogias do dia a dia!
