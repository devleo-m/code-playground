# 📚 Aula 2: Slices em Go

## O que são Slices?

**Slices** são estruturas de dados dinâmicas construídas sobre arrays. Ao contrário dos arrays, slices:

1. **Tamanho dinâmico**: Podem crescer e diminuir conforme necessário
2. **Tipo de referência**: Não copiam todos os dados quando atribuídos
3. **Mais flexíveis**: São o tipo de sequência mais usado em Go

Um slice é composto por três componentes internos:
- **Ponteiro**: Aponta para um array subjacente
- **Length (comprimento)**: Número de elementos no slice
- **Capacity (capacidade)**: Número máximo de elementos que o array subjacente pode conter

---

## 📝 Sintaxe e Declaração

### Forma 1: Declaração com make()

    // Criar slice com length e capacity
    numeros := make([]int, 5)       // length=5, capacity=5, valores zero
    numeros := make([]int, 5, 10)   // length=5, capacity=10
    
    // Slice vazio
    vazio := make([]string, 0)      // length=0, capacity=0

### Forma 2: Slice Literal

    // Similar a array, mas SEM tamanho especificado
    frutas := []string{"Maçã", "Banana", "Laranja"}
    numeros := []int{10, 20, 30, 40, 50}
    vazio := []int{}  // Slice vazio

### Forma 3: Declaração Nil

    var numeros []int  // Slice nil (não inicializado)
    // length = 0, capacity = 0, ponteiro = nil

**Diferença importante:**
- Slice nil: `var s []int` → nil
- Slice vazio: `s := []int{}` → não é nil, mas length=0

---

## 🔍 Length vs Capacity

    numeros := make([]int, 3, 5)
    
    fmt.Println(len(numeros))  // 3 (quantos elementos existem)
    fmt.Println(cap(numeros))  // 5 (quantos elementos PODEM existir)

**Visualização:**

    Array subjacente: [0, 0, 0, _, _]
                       ↑______↑   ↑___↑
                       length=3   capacity=5

---

## 🔍 Acessando e Modificando Elementos

Funciona exatamente como arrays:

    frutas := []string{"Maçã", "Banana", "Laranja"}
    
    fmt.Println(frutas[0])   // Maçã
    frutas[1] = "Morango"
    fmt.Println(frutas)      // [Maçã Morango Laranja]

---

## ➕ Adicionando Elementos com append()

A função `append()` é fundamental para slices:

    numeros := []int{1, 2, 3}
    numeros = append(numeros, 4)
    fmt.Println(numeros)  // [1 2 3 4]
    
    // Adicionar múltiplos elementos
    numeros = append(numeros, 5, 6, 7)
    fmt.Println(numeros)  // [1 2 3 4 5 6 7]
    
    // Adicionar outro slice
    outros := []int{8, 9, 10}
    numeros = append(numeros, outros...)
    fmt.Println(numeros)  // [1 2 3 4 5 6 7 8 9 10]

**IMPORTANTE:** `append()` retorna um novo slice. Sempre atribua o resultado!

    numeros := []int{1, 2, 3}
    append(numeros, 4)        // ERRADO! Resultado perdido
    numeros = append(numeros, 4)  // CORRETO!

---

## 📐 Slicing - Criando Sub-slices

Você pode criar slices a partir de slices (ou arrays):

    numeros := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
    
    // slice[inicio:fim] - fim é EXCLUSIVO
    parte1 := numeros[2:5]    // [2 3 4]
    parte2 := numeros[:4]     // [0 1 2 3] (do início até 4)
    parte3 := numeros[6:]     // [6 7 8 9] (de 6 até o fim)
    parte4 := numeros[:]      // [0 1 2 3 4 5 6 7 8 9] (cópia completa)

**Sintaxe completa:**

    slice[inicio:fim:capacidade_maxima]

Exemplo:

    numeros := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
    parte := numeros[2:5:7]
    // Elementos: [2 3 4]
    // Length: 3 (5-2)
    // Capacity: 5 (7-2)

---

## ⚠️ Slices Compartilham Memória!

**CRUCIAL:** Sub-slices apontam para o mesmo array subjacente!

    original := []int{1, 2, 3, 4, 5}
    parte := original[1:4]  // [2 3 4]
    
    parte[0] = 999
    
    fmt.Println(original)  // [1 999 3 4 5] - FOI MODIFICADO!
    fmt.Println(parte)     // [999 3 4]

Por quê? Ambos apontam para o mesmo array na memória.

---

## 📋 Função copy()

Para criar uma cópia independente:

    original := []int{1, 2, 3, 4, 5}
    copia := make([]int, len(original))
    
    copy(copia, original)
    
    copia[0] = 999
    
    fmt.Println(original)  // [1 2 3 4 5] - NÃO mudou
    fmt.Println(copia)     // [999 2 3 4 5]

**Sintaxe:** `copy(destino, origem)`

A função retorna o número de elementos copiados:

    origem := []int{1, 2, 3, 4, 5}
    destino := make([]int, 3)
    
    n := copy(destino, origem)
    fmt.Println(n)         // 3
    fmt.Println(destino)   // [1 2 3]

---

## 🔄 Iterando sobre Slices

Exatamente como arrays:

    numeros := []int{10, 20, 30, 40, 50}
    
    // Range
    for indice, valor := range numeros {
        fmt.Printf("%d: %d\n", indice, valor)
    }
    
    // For tradicional
    for i := 0; i < len(numeros); i++ {
        fmt.Println(numeros[i])
    }

---

## 🗑️ Removendo Elementos

Go não tem função built-in para remover. Use slicing:

**Remover primeiro elemento:**

    numeros := []int{1, 2, 3, 4, 5}
    numeros = numeros[1:]  // [2 3 4 5]

**Remover último elemento:**

    numeros = numeros[:len(numeros)-1]  // [2 3 4]

**Remover elemento no meio (índice i):**

    i := 2
    numeros = append(numeros[:i], numeros[i+1:]...)
    // Remove o elemento no índice 2

---

## 📊 Slices Multidimensionais

    // Matriz 3x3
    matriz := [][]int{
        {1, 2, 3},
        {4, 5, 6},
        {7, 8, 9},
    }
    
    fmt.Println(matriz[1][2])  // 6
    
    // Adicionar linha
    novaLinha := []int{10, 11, 12}
    matriz = append(matriz, novaLinha)

---

## 🎯 Exemplo Completo: Gerenciador de Tarefas

    package main
    
    import "fmt"
    
    func main() {
        // Lista de tarefas
        tarefas := []string{"Estudar Go", "Fazer exercícios"}
        
        // Adicionar tarefa
        tarefas = append(tarefas, "Revisar código")
        fmt.Println("Tarefas:", tarefas)
        
        // Remover primeira tarefa (completada)
        tarefas = tarefas[1:]
        fmt.Println("Após completar primeira:", tarefas)
        
        // Adicionar múltiplas tarefas
        novasTarefas := []string{"Ler documentação", "Praticar"}
        tarefas = append(tarefas, novasTarefas...)
        
        // Listar todas
        fmt.Println("\nTodas as tarefas:")
        for i, tarefa := range tarefas {
            fmt.Printf("%d. %s\n", i+1, tarefa)
        }
        
        // Total
        fmt.Printf("\nTotal de tarefas: %d\n", len(tarefas))
    }

**Saída:**

    Tarefas: [Estudar Go Fazer exercícios Revisar código]
    Após completar primeira: [Fazer exercícios Revisar código]
    
    Todas as tarefas:
    1. Fazer exercícios
    2. Revisar código
    3. Ler documentação
    4. Praticar
    
    Total de tarefas: 4

---

## 📌 Diferenças Arrays vs Slices

| Característica | Array | Slice |
|---------------|-------|-------|
| Tamanho | Fixo | Dinâmico |
| Tipo | `[5]int` | `[]int` |
| Valor/Referência | Valor (copia) | Referência |
| Pode crescer | Não | Sim (append) |
| Mais usado | Raramente | Sempre! |

---

## 📌 Resumo dos Conceitos-Chave

- **Slice**: Estrutura dinâmica sobre array
- **make()**: Cria slices com length e capacity
- **append()**: Adiciona elementos (sempre atribua resultado!)
- **len()**: Retorna número de elementos
- **cap()**: Retorna capacidade
- **copy()**: Copia elementos entre slices
- **Slicing**: `slice[inicio:fim]` cria sub-slices
- **Referência**: Sub-slices compartilham memória!

---

# 📚 Aula 2 - Simplificada: Entendendo Slices

## 🎒 Analogia: Slices são como Mochilas Expansíveis

Imagine uma mochila mágica que:
- Começa com um tamanho (length)
- Tem um limite máximo atual (capacity)
- Quando enche, magicamente dobra de tamanho!
- Você pode ver através dela e pegar itens de qualquer posição

**Array era como:**
Um armário fixo com 5 gavetas. Não cresce, não diminui.

**Slice é como:**
Uma mochila que começa com 5 bolsos, mas quando você coloca o 6º item, ela automaticamente ganha mais bolsos!

---

## 📦 Os Três Componentes do Slice

Pense no slice como uma "janela" para ver um array:

    Slice tem 3 informações:
    1. Para onde está olhando (ponteiro)
    2. Quantos itens pode ver agora (length)
    3. Quantos itens PODE ver no máximo (capacity)

**Exemplo visual:**

    Array subjacente: [🍎 🍌 🍊 🍇 🍓 _ _ _]
    Slice "enxerga":  [🍎 🍌 🍊]
                       ↑_____↑ length=3
                       ↑_____________↑ capacity=5

---

## ➕ Append - Adicionar na Mochila

Quando você faz `append()`, é como colocar mais um item na mochila:

    mochila := []string{"Livro", "Caderno"}
    mochila = append(mochila, "Lápis")
    
    // mochila agora: ["Livro", "Caderno", "Lápis"]

**Se a mochila encher?**
Go automaticamente pega uma mochila maior e transfere tudo!

    mochila := make([]string, 2, 2)  // 2 itens, capacidade 2
    mochila = append(mochila, "Novo")
    // Go cria nova mochila com capacidade 4 e copia tudo!

---

## ✂️ Slicing - Cortar um Pedaço

Slicing é como tirar uma foto de parte da sua coleção:

    numeros := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
    meioDoArray := numeros[3:7]  // [3 4 5 6]

**Analogia da fita adesiva:**

Imagine uma fita com números de 0 a 9. Você coloca dois dedos:
- Dedo esquerdo no 3
- Dedo direito no 7

Você pegou o pedaço: [3, 4, 5, 6] (o 7 fica de fora!)

---

## ⚠️ Cuidado! Slices Compartilham Memória

**Analogia do caderno compartilhado:**

Você e seu amigo estão olhando para o MESMO caderno, mas páginas diferentes:

    caderno := []string{"Pág1", "Pág2", "Pág3", "Pág4", "Pág5"}
    minhasPaginas := caderno[0:3]    // [Pág1 Pág2 Pág3]
    paginasAmigo := caderno[2:5]     // [Pág3 Pág4 Pág5]
    
    minhasPaginas[2] = "ALTERADO"
    
    // Ambos veem a mudança, pois é o MESMO caderno!
    fmt.Println(caderno)          // [Pág1 Pág2 ALTERADO Pág4 Pág5]
    fmt.Println(paginasAmigo)     // [ALTERADO Pág4 Pág5]

**Solução:** Use `copy()` para fazer um caderno novo separado!

---

## 📋 Copy - Fazer uma Cópia Real

`copy()` é como fazer xerox do caderno:

    original := []int{1, 2, 3}
    xerox := make([]int, len(original))
    copy(xerox, original)
    
    xerox[0] = 999
    
    // original: [1 2 3] - não mudou
    // xerox: [999 2 3] - só a cópia mudou

---

## 🗑️ Removendo Itens - Analogia da Fila

**Remover da frente (primeira pessoa sai):**

    fila := []string{"João", "Maria", "Pedro"}
    fila = fila[1:]  // [Maria Pedro]

**Remover do final (última pessoa sai):**

    fila = fila[:len(fila)-1]  // [Maria]

**Remover do meio:**

É como juntar duas partes da fila, pulando uma pessoa:

    fila := []string{"João", "Maria", "Pedro", "Ana"}
    // Remover "Maria" (índice 1)
    fila = append(fila[:1], fila[2:]...)
    // [João Pedro Ana]

---

## 🎒 Length vs Capacity - Mochila Mágica

**Length:** Quantas coisas você colocou na mochila AGORA
**Capacity:** Quantas coisas CABEM antes dela precisar crescer

    mochila := make([]string, 3, 5)
    
    len(mochila) = 3   // 3 itens dentro
    cap(mochila) = 5   // Cabe até 5 antes de expandir

Quando você faz `append()` no 6º item:
- Go pega uma mochila nova (geralmente dobra o tamanho → 10)
- Copia tudo da mochila velha
- Coloca o novo item
- Joga a mochila velha fora

---

## 🎯 Quando Usar Slices vs Arrays

**Use SLICES quando:**
- Não sabe quantos elementos terá ✅
- Precisa adicionar/remover elementos ✅
- A maioria dos casos! ✅

**Use ARRAYS quando:**
- Tamanho absolutamente fixo (dias da semana, coordenadas x/y)
- Casos muito raros!

**Regra de ouro:** 95% do tempo use slices!

---

## 💡 Dica Visual: Detectar Array vs Slice

    [5]int    → Array (tem número!)
    []int     → Slice (sem número!)

---

## 🎲 Exemplo do Mundo Real: Lista de Compras

    listaCompras := []string{}  // Começa vazia
    
    // Adicionar itens
    listaCompras = append(listaCompras, "Arroz")
    listaCompras = append(listaCompras, "Feijão")
    listaCompras = append(listaCompras, "Açúcar")
    
    // Comprou o primeiro item (remover)
    listaCompras = listaCompras[1:]
    
    // Adicionar mais itens
    listaCompras = append(listaCompras, "Café", "Leite")

Perfeito para slices! Tamanho muda constantemente.

---

## 📌 Resumo Visual

    Array:  [📦][📦][📦][📦][📦]  ← Fixo, não cresce
    
    Slice:  [📦][📦][📦][  ][  ]  ← Pode crescer infinitamente!
            ↑______↑ length
            ↑__________↑ capacity

---

# 📚 Aula 2 - Exercícios e Reflexão

## 🏋️ Exercício 1: Lista de Números Dinâmica

Crie um programa que:
1. Comece com um slice vazio de inteiros
2. Adicione os números: 5, 10, 15, 20, 25
3. Exiba o slice, seu length e capacity
4. Remova o primeiro elemento
5. Adicione os números 30, 35, 40
6. Exiba o slice final, length e capacity

---

## 🏋️ Exercício 2: Filtrar Números Pares

Crie um programa que:
1. Declare um slice: [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]
2. Crie um novo slice contendo apenas os números PARES
3. Use um loop e append para construir o novo slice
4. Exiba o slice de números pares

---

## 🏋️ Exercício 3: Compartilhamento de Memória

Crie um programa que:
1. Crie um slice original: [10, 20, 30, 40, 50]
2. Crie um sub-slice pegando os elementos do índice 1 ao 4: [20, 30, 40]
3. Modifique o primeiro elemento do sub-slice para 999
4. Exiba AMBOS os slices (original e sub-slice)
5. Explique no comentário do código: por que o original também mudou?

---

## 🏋️ Exercício 4: Copy vs Referência

Crie um programa que demonstre a diferença entre:
1. Criar um sub-slice (compartilha memória)
2. Usar copy() para criar cópia independente

Crie dois cenários lado a lado e mostre a diferença quando você modifica os valores.

---

## 🤔 Perguntas de Reflexão

### Pergunta 1: Por que Go Usa Slices Como Referência?

Você aprendeu que slices são "reference types" (não copiam todo o conteúdo).

Pense sobre:
- Qual é a vantagem de slices serem referências ao invés de cópias?
- Como isso impacta a performance quando você passa slices para funções?
- Quais cuidados você precisa ter por causa disso?

Escreva sua resposta com suas próprias palavras (mínimo 3 linhas).

---

### Pergunta 2: Capacity e Realocação

Quando você faz `append()` em um slice cheio, Go precisa realocar (criar array maior e copiar).

Reflita:
- Por que Go geralmente DOBRA a capacidade ao invés de aumentar apenas 1?
- Em que situação seria melhor usar `make([]int, 0, 1000)` ao invés de `[]int{}`?
- Como pré-alocar capacity pode melhorar performance?

Escreva sua resposta com suas próprias palavras (mínimo 4 linhas).

---

### Pergunta 3: Arrays vs Slices - Decisão Real

Imagine que você está desenvolvendo um sistema de controle de estoque.

Descreva:
- Você usaria arrays ou slices para armazenar os produtos? Por quê?
- Dê 3 razões específicas para sua escolha
- Existe alguma parte do sistema onde você usaria o outro tipo?

Escreva sua resposta com suas próprias palavras (mínimo 5 linhas).

---

# 📚 Aula 2 - Performance e Boas Práticas

## ⚡ Performance de Slices

### 1. Pré-alocação de Capacity - CRÍTICO!

**RUIM (causa múltiplas realocações):**

    numeros := []int{}
    for i := 0; i < 100000; i++ {
        numeros = append(numeros, i)
        // Realocação ocorre ~17 vezes!
    }

**BOM (uma única alocação):**

    numeros := make([]int, 0, 100000)
    for i := 0; i < 100000; i++ {
        numeros = append(numeros, i)
        // Nenhuma realocação!
    }

**Impacto de performance:**
- Ruim: ~17 alocações + cópias
- Bom: 1 alocação, 0 cópias
- Diferença: **10-50x mais rápido!**

---

### 2. Crescimento de Capacity - Como Go Funciona

Quando um slice precisa crescer:

**Para slices pequenos (< 1024 elementos):**
- Nova capacity = 2 × capacity atual

**Para slices grandes (≥ 1024 elementos):**
- Nova capacity = 1.25 × capacity atual

**Exemplo:**

    s := make([]int, 0, 4)
    
    // Capacidades ao adicionar elementos:
    // 4 → 8 → 16 → 32 → 64 → 128 → 256 → 512 → 1024
    // Depois de 1024:
    // 1024 → 1280 → 1600 → 2000 → ...

---

### 3. Slice Header - O Que Realmente é Copiado

Um slice é apenas uma estrutura pequena (24 bytes em 64-bit):

    type slice struct {
        ptr unsafe.Pointer  // 8 bytes - ponteiro para array
        len int             // 8 bytes - length
        cap int             // 8 bytes - capacity
    }

**Consequência:** Passar slice para função é BARATO!

    func processar(dados []int) {
        // Copia apenas 24 bytes (slice header)
        // NÃO copia os elementos!
    }

---

## ✅ Boas Práticas

### Prática 1: Sempre Pré-aloque Quando Souber o Tamanho

**BOM:**

    // Sei que terei ~1000 elementos
    resultados := make([]float64, 0, 1000)
    for _, item := range dados {
        resultados = append(resultados, processar(item))
    }

**RUIM:**

    resultados := []float64{}  // Vai realocar múltiplas vezes

---

### Prática 2: Use len(), Não Confie em Capacity

**BOM:**

    numeros := []int{1, 2, 3}
    if len(numeros) > 0 {
        primeiro := numeros[0]
    }

**RUIM:**

    if cap(numeros) > 0 {  // Capacity não garante elementos!
        primeiro := numeros[0]  // PODE DAR PANIC!
    }

---

### Prática 3: Sempre Atribua o Resultado de append()

**ERRADO:**

    numeros := []int{1, 2, 3}
    append(numeros, 4)  // Resultado perdido!
    fmt.Println(numeros)  // [1 2 3]

**CORRETO:**

    numeros := []int{1, 2, 3}
    numeros = append(numeros, 4)
    fmt.Println(numeros)  // [1 2 3 4]

---

### Prática 4: Cuidado com Sub-slices e Memória

**PROBLEMA:**

    func lerArquivo() []byte {
        dados := ioutil.ReadFile("huge_file.txt")  // 1GB
        return dados[0:100]  // Retorna só 100 bytes
    }
    
    // O arquivo inteiro de 1GB fica na memória!
    // Porque o sub-slice mantém referência ao array completo!

**SOLUÇÃO:**

    func lerArquivo() []byte {
        dados := ioutil.ReadFile("huge_file.txt")
        resultado := make([]byte, 100)
        copy(resultado, dados[0:100])
        return resultado  // Agora pode liberar os 1GB
    }

---

### Prática 5: Slice Nil vs Slice Vazio

Ambos têm `len() == 0`, mas são diferentes:

    var s1 []int        // nil slice
    s2 := []int{}       // empty slice
    s3 := make([]int, 0)  // empty slice
    
    fmt.Println(s1 == nil)  // true
    fmt.Println(s2 == nil)  // false
    fmt.Println(s3 == nil)  // false

**Quando usar cada um:**

- **Nil slice:** Valor padrão, economiza memória

      var usuarios []Usuario  // Sem usuários ainda

- **Empty slice:** Quando quer representar "lista vazia" explicitamente

      usuarios := []Usuario{}  // Lista vazia pronta para receber

**JSON encoding:**

    var s1 []int        // JSON: null
    s2 := []int{}       // JSON: []

---

## 🚫 O Que NÃO Fazer

### ❌ Erro 1: Reutilizar Slice Sem Redefinir Length

    numeros := make([]int, 5)
    // Length = 5, todos valores zero: [0 0 0 0 0]
    
    for i := 0; i < 3; i++ {
        numeros = append(numeros, i)
    }
    
    // Resultado: [0 0 0 0 0 0 1 2]
    // INESPERADO! Esqueceu dos 5 zeros iniciais

**CORRETO:**

    numeros := make([]int, 0, 5)  // Length=0, Capacity=5
    for i := 0; i < 3; i++ {
        numeros = append(numeros, i)
    }
    // Resultado: [0 1 2]

---

### ❌ Erro 2: Modificar Slice Durante Iteração

**PERIGOSO:**

    numeros := []int{1, 2, 3, 4, 5}
    
    for i, v := range numeros {
        if v%2 == 0 {
            numeros = append(numeros[:i], numeros[i+1:]...)
        }
    }
    // Comportamento indefinido!

**CORRETO:**

    numeros := []int{1, 2, 3, 4, 5}
    resultado := []int{}
    
    for _, v := range numeros {
        if v%2 != 0 {  // Manter ímpares
            resultado = append(resultado, v)
        }
    }
    numeros = resultado

---

### ❌ Erro 3: Confiar em Sub-slice Após Append

    original := []int{1, 2, 3, 4, 5}
    parte := original[1:3]  // [2 3]
    
    original = append(original, 6, 7, 8, 9, 10)
    
    // 'parte' pode estar INVÁLIDO agora!
    // Se append realocou, 'parte' aponta para memória antiga

---

## 🎯 Padrões Idiomáticos em Go

### Padrão 1: Filtrar Slice

    numeros := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    
    // Manter apenas pares
    pares := numeros[:0]  // Reutiliza capacity
    for _, n := range numeros {
        if n%2 == 0 {
            pares = append(pares, n)
        }
    }

---

### Padrão 2: Remover Elemento (Preservando Ordem)

    func remover(slice []int, i int) []int {
        return append(slice[:i], slice[i+1:]...)
    }
    
    numeros := []int{1, 2, 3, 4, 5}
    numeros = remover(numeros, 2)  // Remove índice 2

---

### Padrão 3: Remover Elemento (SEM Preservar Ordem - Mais Rápido)

    func removerRapido(slice []int, i int) []int {
        slice[i] = slice[len(slice)-1]
        return slice[:len(slice)-1]
    }
    
    numeros := []int{1, 2, 3, 4, 5}
    numeros = removerRapido(numeros, 2)  // [1 2 5 4]

---

### Padrão 4: Inserir no Meio

    func inserir(slice []int, i, value int) []int {
        slice = append(slice[:i+1], slice[i:]...)
        slice[i] = value
        return slice
    }

    numeros := []int{1, 2, 4, 5}
    numeros = inserir(numeros, 2, 3)  // [1 2 3 4 5]

---

### Padrão 5: Reverter Slice

    func reverter(slice []int) {
        for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
            slice[i], slice[j] = slice[j], slice[i]
        }
    }
    
    numeros := []int{1, 2, 3, 4, 5}
    reverter(numeros)
    // numeros agora: [5 4 3 2 1]

---

## 📊 Comparação de Performance

### Cenário 1: Adicionar 100.000 elementos

**Sem pré-alocação:**

    inicio := time.Now()
    s := []int{}
    for i := 0; i < 100000; i++ {
        s = append(s, i)
    }
    fmt.Println(time.Since(inicio))  // ~5ms

**Com pré-alocação:**

    inicio := time.Now()
    s := make([]int, 0, 100000)
    for i := 0; i < 100000; i++ {
        s = append(s, i)
    }
    fmt.Println(time.Since(inicio))  // ~0.5ms

**Ganho: 10x mais rápido!**

---

### Cenário 2: Copiar vs Sub-slice

**Sub-slice (compartilha memória):**

    original := make([]int, 1000000)
    parte := original[0:10]  // Instantâneo
    // Mas mantém 1 milhão de ints na memória!

**Copy (memória independente):**

    original := make([]int, 1000000)
    parte := make([]int, 10)
    copy(parte, original[0:10])  // ~100 nanosegundos
    // Apenas 10 ints na memória

---

## 💡 Otimizações Avançadas

### Otimização 1: Reutilizar Slice ao Invés de Criar Novo

**RUIM (cria garbage):**

    func processar(dados []int) []int {
        resultado := []int{}  // Nova alocação toda vez
        for _, d := range dados {
            if d > 0 {
                resultado = append(resultado, d)
            }
        }
        return resultado
    }

**BOM (reutiliza):**

    func processar(dados []int) []int {
        resultado := dados[:0]  // Reutiliza capacity
        for _, d := range dados {
            if d > 0 {
                resultado = append(resultado, d)
            }
        }
        return resultado
    }

---

### Otimização 2: Evitar Append em Loops Quentes

**RUIM:**

    for i := 0; i < 1000000; i++ {
        slice = append(slice, calcular(i))
    }

**BOM:**

    slice := make([]int, 1000000)
    for i := 0; i < 1000000; i++ {
        slice[i] = calcular(i)
    }

Se sabe o tamanho final, use indexação direta!

---

### Otimização 3: Batch Append

**RUIM:**

    for _, item := range items {
        resultado = append(resultado, item)
    }

**BOM:**

    resultado = append(resultado, items...)  // Uma operação

---

## 🔍 Debugging - Entendendo Realocações

Use este código para visualizar realocações:

    s := make([]int, 0, 2)
    oldCap := cap(s)
    
    for i := 0; i < 20; i++ {
        s = append(s, i)
        if cap(s) != oldCap {
            fmt.Printf("Realocou! Novo cap: %d\n", cap(s))
            oldCap = cap(s)
        }
    }

**Saída:**

    Realocou! Novo cap: 4
    Realocou! Novo cap: 8
    Realocou! Novo cap: 16
    Realocou! Novo cap: 32

---

## 📏 Tamanho de Slice na Memória

**Slice header:** 24 bytes (fixo)

**Elementos:** depende do tipo

    []int{1,2,3}
    - Header: 24 bytes
    - Elementos: 3 × 8 bytes = 24 bytes
    - Total: 48 bytes
    
    []byte{1,2,3}
    - Header: 24 bytes
    - Elementos: 3 × 1 byte = 3 bytes
    - Total: 27 bytes

---

## 🎯 Casos de Uso Específicos

### Caso 1: Processamento em Lote

    func processarLote(dados []int, tamanhoLote int) {
        for i := 0; i < len(dados); i += tamanhoLote {
            fim := i + tamanhoLote
            if fim > len(dados) {
                fim = len(dados)
            }
            
            lote := dados[i:fim]
            processar(lote)
        }
    }

---

### Caso 2: Ring Buffer (Buffer Circular)

    type RingBuffer struct {
        dados []int
        inicio int
        tamanho int
    }
    
    func (rb *RingBuffer) Adicionar(valor int) {
        pos := (rb.inicio + rb.tamanho) % len(rb.dados)
        rb.dados[pos] = valor
        if rb.tamanho < len(rb.dados) {
            rb.tamanho++
        } else {
            rb.inicio = (rb.inicio + 1) % len(rb.dados)
        }
    }

---

### Caso 3: Pool de Slices (Reduzir Garbage Collection)

    var bufferPool = sync.Pool{
        New: func() interface{} {
            return make([]byte, 4096)
        },
    }
    
    func processarDados() {
        buffer := bufferPool.Get().([]byte)
        defer bufferPool.Put(buffer)
        
        // Usar buffer...
    }

---

## 📌 Checklist de Boas Práticas

- [ ] Pré-alocar capacity quando souber o tamanho aproximado
- [ ] Sempre atribuir resultado de append()
- [ ] Usar `make([]T, 0, n)` ao invés de `make([]T, n)` para append
- [ ] Cuidado com sub-slices de grandes arrays (vazamento de memória)
- [ ] Use copy() quando precisar de slice independente
- [ ] Evitar modificar slice durante iteração
- [ ] Preferir indexação direta quando souber tamanho final
- [ ] Reutilizar slices ao invés de criar novos em loops
- [ ] Entender diferença entre nil slice e empty slice

---

## ⚠️ Armadilhas Comuns

### Armadilha 1: Append com Slices Compartilhados

    a := []int{1, 2, 3}
    b := a[0:2]  // [1 2]
    c := a[0:2]  // [1 2]
    
    b = append(b, 99)
    c = append(c, 88)
    
    // O que acontece? Depende da capacity!

---

### Armadilha 2: Slice de Ponteiros

    type Item struct {
        valor int
    }
    
    items := []*Item{}
    
    for i := 0; i < 5; i++ {
        item := Item{valor: i}
        items = append(items, &item)
    }
    
    // TODOS os ponteiros apontam para o MESMO item!
    // item é reutilizado em cada iteração

**CORRETO:**

    for i := 0; i < 5; i++ {
        item := Item{valor: i}
        items = append(items, &item)  // & cria cópia
    }
    
    // OU melhor:
    for i := 0; i < 5; i++ {
        items = append(items, &Item{valor: i})
    }

---

### Armadilha 3: Range Copia Valores

    type Pessoa struct {
        nome string
        idade int
    }
    
    pessoas := []Pessoa{
        {"João", 30},
        {"Maria", 25},
    }
    
    for _, p := range pessoas {
        p.idade++  // Modifica CÓPIA, não original!
    }

**CORRETO:**

    for i := range pessoas {
        pessoas[i].idade++
    }

---

## 🎓 Comparação: Arrays vs Slices - Decisão Final

| Situação | Use |
|----------|-----|
| Tamanho fixo conhecido (dias semana, coordenadas) | Array |
| Tamanho variável | Slice |
| Passar para função | Slice |
| Performance crítica e tamanho pequeno | Array |
| Adicionar/remover elementos | Slice |
| API pública | Slice |
| Buffer temporário pequeno | Array |
| Coleções em geral | Slice |

**Regra geral:** Use slices em 95% dos casos!

---

## 📈 Benchmark Real

Código para testar performance:

    func BenchmarkSemPreAlocacao(b *testing.B) {
        for i := 0; i < b.N; i++ {
            s := []int{}
            for j := 0; j < 10000; j++ {
                s = append(s, j)
            }
        }
    }
    
    func BenchmarkComPreAlocacao(b *testing.B) {
        for i := 0; i < b.N; i++ {
            s := make([]int, 0, 10000)
            for j := 0; j < 10000; j++ {
                s = append(s, j)
            }
        }
    }

**Resultado típico:**

    BenchmarkSemPreAlocacao-8     5000    250000 ns/op
    BenchmarkComPreAlocacao-8    50000     25000 ns/op

**10x mais rápido com pré-alocação!**

---

## 📚 Recursos Adicionais

### Comandos Úteis para Debug

    s := make([]int, 3, 5)
    fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
    fmt.Printf("%p\n", s)  // Endereço do array subjacente

---

## 📌 Resumo Final - Slices

**O que você DEVE fazer:**
1. Sempre pré-alocar quando souber o tamanho
2. Atribuir resultado de append()
3. Usar copy() quando precisar independência
4. Cuidado com compartilhamento de memória em sub-slices

**O que você NÃO deve fazer:**
1. Confiar em capacity para acessar elementos
2. Modificar slice durante iteração range
3. Esquecer que sub-slices compartilham memória
4. Criar slices com length quando vai usar append

---

**Fim da Aula 2: Performance e Boas Práticas**

---

## 🎯 Próximo Passo

Agora que você completou TODAS as 4 etapas da Aula 2, aguardo suas respostas aos exercícios e perguntas de reflexão!

Quando terminar TODAS as aulas, enviarei suas respostas e farei a **Análise de Desempenho Completa** do seu aprendizado!

**Continuando para a Aula 3: Capacity and Growth** 🚀