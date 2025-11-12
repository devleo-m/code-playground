# 📚 Aula 3: Capacity and Growth (Capacidade e Crescimento)

## O que é Capacity?

**Capacity (capacidade)** é o número máximo de elementos que o array subjacente de um slice pode armazenar antes de precisar ser realocado. É diferente de **length** (comprimento), que indica quantos elementos estão atualmente no slice.

**Componentes de um Slice:**

    type slice struct {
        ponteiro  *array    // Aponta para o array subjacente
        length    int       // Quantos elementos existem
        capacity  int       // Quantos elementos CABEM
    }

**Visualização:**

    Array subjacente: [1][2][3][_][_][_][_][_]
    Slice:            [1][2][3]
                       ↑_____↑  length = 3
                       ↑___________________↑  capacity = 8

---

## 🔍 Diferença entre Length e Capacity

    s := make([]int, 3, 5)
    
    len(s)  // 3 - elementos que EXISTEM
    cap(s)  // 5 - elementos que PODEM existir sem realocação
    
    // Slice: [0 0 0]
    // Array subjacente: [0 0 0 _ _]

**Length:** Número de elementos acessíveis no slice agora

**Capacity:** Espaço total disponível no array subjacente a partir da posição inicial do slice

---

## 📈 Como Slices Crescem - O Algoritmo

Quando você faz `append()` e não há espaço (length == capacity), Go:

1. **Cria um novo array** com capacity maior
2. **Copia todos os elementos** do array antigo para o novo
3. **Adiciona o novo elemento**
4. **Atualiza o slice** para apontar para o novo array
5. **Descarta o array antigo** (garbage collector limpa)

---

## 🔢 Regras de Crescimento do Go

Go usa uma estratégia de **crescimento exponencial** para minimizar realocações:

**Para slices pequenos (capacity < 256):**

    Nova capacity = 2 × capacity atual

**Para slices grandes (capacity ≥ 256):**

    Nova capacity ≈ 1.25 × capacity atual
    
    Fórmula exata (Go 1.18+):
    newcap = oldcap + (oldcap + 3*256) / 4

**Exemplo de crescimento:**

    Capacity inicial: 0
    Após 1º append: 1
    Após 2º append: 2
    Após 3º append: 4
    Após 5º append: 8
    Após 9º append: 16
    Após 17º append: 32
    Após 33º append: 64
    Após 65º append: 128
    Após 129º append: 256
    Após 257º append: 512
    Após 513º append: 848  (não dobra mais!)
    Após 849º append: 1280
    ...

---

## 💻 Demonstração Prática - Observando Crescimento

    package main
    
    import "fmt"
    
    func main() {
        s := []int{}
        oldCap := cap(s)
        
        fmt.Printf("Início: len=%d cap=%d\n", len(s), cap(s))
        
        for i := 0; i < 30; i++ {
            s = append(s, i)
            
            if cap(s) != oldCap {
                fmt.Printf("Após adicionar %d: len=%d cap=%d (cresceu!)\n", 
                    i, len(s), cap(s))
                oldCap = cap(s)
            }
        }
    }

**Saída esperada:**

    Início: len=0 cap=0
    Após adicionar 0: len=1 cap=1 (cresceu!)
    Após adicionar 1: len=2 cap=2 (cresceu!)
    Após adicionar 2: len=3 cap=4 (cresceu!)
    Após adicionar 4: len=5 cap=8 (cresceu!)
    Após adicionar 8: len=9 cap=16 (cresceu!)
    Após adicionar 16: len=17 cap=32 (cresceu!)

---

## ⚡ Custo de Realocação

Cada realocação tem custo:

**Operações necessárias:**
1. Alocar novo array maior
2. Copiar TODOS os elementos do array antigo
3. Adicionar novo elemento
4. Liberar memória antiga

**Custo em Big O:**

    Append sem realocação: O(1) - constante
    Append com realocação: O(n) - linear (copia n elementos)

**Por isso pré-alocação é importante!**

---

## 🎯 Pré-alocação com make()

### Sintaxe do make() para Slices

    make([]tipo, length, capacity)
    make([]tipo, length)  // capacity = length

**Exemplos:**

    // Length=0, Capacity=0
    s1 := []int{}
    
    // Length=5, Capacity=5, valores zero
    s2 := make([]int, 5)
    
    // Length=0, Capacity=10 (ideal para append)
    s3 := make([]int, 0, 10)
    
    // Length=5, Capacity=20
    s4 := make([]int, 5, 20)

---

## 📊 Comparação: Com vs Sem Pré-alocação

### Sem Pré-alocação (RUIM)

    s := []int{}
    for i := 0; i < 1000; i++ {
        s = append(s, i)
    }

**O que acontece:**
- Realocações: ~10 vezes
- Cópias de elementos: ~2047 elementos copiados no total
- Alocações de memória: ~10

### Com Pré-alocação (BOM)

    s := make([]int, 0, 1000)
    for i := 0; i < 1000; i++ {
        s = append(s, i)
    }

**O que acontece:**
- Realocações: 0
- Cópias de elementos: 0
- Alocações de memória: 1

**Ganho:** 10-20x mais rápido!

---

## 🔬 Análise Matemática de Realocações

Vamos calcular quantas vezes um slice é realocado:

**Para adicionar N elementos começando de capacity=0:**

    Realocações ≈ log₂(N)

**Elementos copiados no total:**

    Total de cópias ≈ 2N

**Exemplo com N=1024:**

    Realocações: log₂(1024) = 10
    Cópias totais: ~2048 elementos

---

## 💡 Quando Usar Cada Forma de make()

### make([]T, length)

**Use quando:** Você sabe quantos elementos precisa E vai acessar por índice

    // Criar slice de 100 zeros
    numeros := make([]int, 100)
    
    for i := 0; i < 100; i++ {
        numeros[i] = i * 2  // Acesso direto por índice
    }

---

### make([]T, 0, capacity)

**Use quando:** Você sabe quantos elementos terá MAS vai usar append()

    // Vai adicionar ~1000 elementos
    numeros := make([]int, 0, 1000)
    
    for i := 0; i < 1000; i++ {
        numeros = append(numeros, i)  // Sem realocações!
    }

---

### []T{} ou var s []T

**Use quando:** Você NÃO sabe quantos elementos terá

    numeros := []int{}
    
    for _, item := range dados {
        if item.valido {
            numeros = append(numeros, item.valor)
        }
    }

---

## 🎯 Exemplo Completo: Processamento de Dados

    package main
    
    import (
        "fmt"
        "time"
    )
    
    func semPreAlocacao() {
        inicio := time.Now()
        
        s := []int{}
        for i := 0; i < 100000; i++ {
            s = append(s, i)
        }
        
        duracao := time.Since(inicio)
        fmt.Printf("Sem pré-alocação: %v\n", duracao)
    }
    
    func comPreAlocacao() {
        inicio := time.Now()
        
        s := make([]int, 0, 100000)
        for i := 0; i < 100000; i++ {
            s = append(s, i)
        }
        
        duracao := time.Since(inicio)
        fmt.Printf("Com pré-alocação: %v\n", duracao)
    }
    
    func main() {
        semPreAlocacao()
        comPreAlocacao()
    }

**Saída típica:**

    Sem pré-alocação: 3.5ms
    Com pré-alocação: 0.4ms

**Diferença: ~8-10x mais rápido!**

---

## 🔄 Capacity de Sub-slices

Quando você cria um sub-slice, a capacity é calculada do **início do sub-slice até o fim do array original**:

    original := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
    
    sub1 := original[2:5]
    fmt.Printf("len=%d cap=%d\n", len(sub1), cap(sub1))
    // len=3 cap=8 (de índice 2 até o fim: 10-2=8)
    
    sub2 := original[7:9]
    fmt.Printf("len=%d cap=%d\n", len(sub2), cap(sub2))
    // len=2 cap=3 (de índice 7 até o fim: 10-7=3)

---

## 🎯 Limitando Capacity com Full Slice Expression

Você pode limitar a capacity de um sub-slice usando a **expressão de slice completa**:

    slice[inicio:fim:capacidade_maxima]

**Exemplo:**

    original := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
    
    // Slice normal
    sub1 := original[2:5]
    fmt.Printf("len=%d cap=%d\n", len(sub1), cap(sub1))
    // len=3 cap=8
    
    // Limitando capacity
    sub2 := original[2:5:5]
    fmt.Printf("len=%d cap=%d\n", len(sub2), cap(sub2))
    // len=3 cap=3 (capacity limitada!)

**Por que fazer isso?**

Prevenir que append() no sub-slice modifique o slice original:

    original := []int{0, 1, 2, 3, 4, 5}
    
    // Sem limitar capacity
    sub1 := original[0:2]  // [0 1], cap=6
    sub1 = append(sub1, 99)
    fmt.Println(original)  // [0 1 99 3 4 5] - MODIFICOU!
    
    // Limitando capacity
    original = []int{0, 1, 2, 3, 4, 5}
    sub2 := original[0:2:2]  // [0 1], cap=2
    sub2 = append(sub2, 99)
    fmt.Println(original)  // [0 1 2 3 4 5] - NÃO modificou!
    fmt.Println(sub2)      // [0 1 99] - novo array criado

---

## 📐 Calculando Capacity Ideal

### Regra Prática 1: Conhece o tamanho exato

    elementos := make([]int, 0, tamanhoExato)

**Exemplo:**

    // Processar 1000 itens
    resultados := make([]int, 0, 1000)
    for i := 0; i < 1000; i++ {
        resultados = append(resultados, processar(i))
    }

---

### Regra Prática 2: Estima aproximada

    elementos := make([]int, 0, estimativa*1.1)

**Exemplo:**

    // Espera ~500-1000 elementos
    resultados := make([]int, 0, 1100)  // 10% de margem

---

### Regra Prática 3: Não sabe o tamanho

    elementos := []int{}  // Deixa Go gerenciar

**Exemplo:**

    // Não sei quantos elementos válidos existem
    validos := []int{}
    for _, item := range dados {
        if item.valido {
            validos = append(validos, item.valor)
        }
    }

---

## ⚠️ Armadilhas com Capacity

### Armadilha 1: make() com Length > 0

    // ERRADO para usar com append
    s := make([]int, 10)
    for i := 0; i < 5; i++ {
        s = append(s, i)
    }
    fmt.Println(s)  // [0 0 0 0 0 0 0 0 0 0 0 1 2 3 4]
    
    // CORRETO
    s := make([]int, 0, 10)
    for i := 0; i < 5; i++ {
        s = append(s, i)
    }
    fmt.Println(s)  // [0 1 2 3 4]

---

### Armadilha 2: Pré-alocar Demais

    // Desperdício de memória
    s := make([]int, 0, 10000000)  // 10 milhões!
    s = append(s, 1)
    s = append(s, 2)
    // Alocou 80MB mas usa apenas 16 bytes!

**Solução:** Aloque apenas o necessário ou use estimativas razoáveis.

---

### Armadilha 3: Esquecer que Capacity é do Array Subjacente

    original := make([]int, 5, 10)
    sub := original[3:5]
    
    fmt.Println(cap(sub))  // 7 (não 10!)
    // Capacity é de onde o slice começa até o fim do array

---

## 🎲 Growth Strategy - Por Que Dobrar?

**Por que Go dobra a capacity para slices pequenos?**

1. **Minimiza realocações:** log₂(N) realocações ao invés de N
2. **Amortiza o custo:** Cada elemento é copiado em média 2 vezes
3. **Previsível:** Comportamento consistente

**Exemplo:** Adicionar 1024 elementos

- **Estratégia "adicionar 1":** 1024 realocações, milhões de cópias
- **Estratégia "dobrar":** 10 realocações, ~2048 cópias

---

## 📊 Tabela de Crescimento Real

| Elementos | Realocações | Capacity Final | Cópias Totais |
|-----------|-------------|----------------|---------------|
| 10 | 4 | 16 | ~20 |
| 100 | 7 | 128 | ~200 |
| 1.000 | 10 | 1.024 | ~2.000 |
| 10.000 | 14 | 16.384 | ~20.000 |
| 100.000 | 17 | 131.072 | ~200.000 |
| 1.000.000 | 21 | 1.048.576 | ~2.000.000 |

**Padrão:** Cada elemento é copiado aproximadamente **2 vezes** em média.

---

## 🔬 Analisando Memória

    package main
    
    import (
        "fmt"
        "runtime"
    )
    
    func main() {
        var m1, m2 runtime.MemStats
        
        runtime.ReadMemStats(&m1)
        
        // Criar slice grande
        s := make([]int, 0, 1000000)
        for i := 0; i < 1000000; i++ {
            s = append(s, i)
        }
        
        runtime.ReadMemStats(&m2)
        
        fmt.Printf("Memória alocada: %d MB\n", 
            (m2.Alloc-m1.Alloc)/1024/1024)
    }

---

## 💡 Casos Especiais

### Caso 1: Slice de Structs Grandes

    type Pessoa struct {
        nome      string
        idade     int
        endereco  string
        telefone  string
        email     string
    }
    
    // Cada Pessoa ~100 bytes
    // 10.000 pessoas = ~1MB
    
    // Pré-alocar é CRÍTICO aqui!
    pessoas := make([]Pessoa, 0, 10000)

---

### Caso 2: Slice de Ponteiros

    // Slice de ponteiros é barato de realocar
    ponteiros := []*Objeto{}
    
    // Cada ponteiro = 8 bytes
    // Realocação copia apenas ponteiros, não objetos

---

### Caso 3: Slice Dentro de Struct

    type Cache struct {
        dados []int
    }
    
    func NovoCache(tamanho int) *Cache {
        return &Cache{
            dados: make([]int, 0, tamanho),
        }
    }

---

## 🎯 Exemplo Real: Processamento de Log

    package main
    
    import (
        "fmt"
        "strings"
    )
    
    type LogEntry struct {
        timestamp string
        level     string
        message   string
    }
    
    func processarLogs(linhas []string) []LogEntry {
        // Pré-alocar assumindo todas linhas são válidas
        logs := make([]LogEntry, 0, len(linhas))
        
        for _, linha := range linhas {
            if strings.Contains(linha, "ERROR") {
                partes := strings.Split(linha, "|")
                if len(partes) >= 3 {
                    log := LogEntry{
                        timestamp: partes[0],
                        level:     partes[1],
                        message:   partes[2],
                    }
                    logs = append(logs, log)
                }
            }
        }
        
        return logs
    }
    
    func main() {
        linhas := []string{
            "2024-01-01|INFO|Sistema iniciado",
            "2024-01-01|ERROR|Falha na conexão",
            "2024-01-01|WARN|Memória alta",
            "2024-01-01|ERROR|Timeout",
        }
        
        erros := processarLogs(linhas)
        fmt.Printf("Encontrados %d erros\n", len(erros))
        fmt.Printf("Capacity: %d\n", cap(erros))
    }

---

## 📌 Resumo dos Conceitos-Chave

- **Length:** Quantos elementos existem
- **Capacity:** Quanto espaço está alocado
- **Realocação:** Ocorre quando len == cap e você faz append
- **Crescimento:** Dobra para < 256, cresce ~25% para ≥ 256
- **Pré-alocação:** Use make([]T, 0, n) quando souber o tamanho
- **Custo:** Realocação é O(n), pré-alocação evita isso
- **Full slice expression:** slice[i:j:k] limita capacity

---

# 📚 Aula 3 - Simplificada: Entendendo Capacity and Growth

## 🎒 Analogia: Mochila com Bolsos Extras

Imagine que você tem uma mochila:

**Length (comprimento):** Quantos objetos você JÁ colocou na mochila

**Capacity (capacidade):** Quantos bolsos a mochila TEM (mesmo que vazios)

**Exemplo:**

    Mochila com 10 bolsos (capacity = 10)
    Você colocou 3 livros (length = 3)
    Ainda tem 7 bolsos vazios disponíveis

---

## 📦 O Que Acontece Quando a Mochila Enche?

Você tem uma mochila com 4 bolsos, todos cheios:

    Bolsos: [📚][📚][📚][📚]
    Length: 4
    Capacity: 4

Você quer adicionar mais um livro 📕

**O que Go faz:**

1. 🎒 Compra uma mochila MAIOR (dobro de bolsos = 8)
2. 📦 Transfere TODOS os livros da mochila velha para a nova
3. 📕 Adiciona o livro novo
4. 🗑️ Joga a mochila velha fora

**Resultado:**

    Bolsos: [📚][📚][📚][📚][📕][  ][  ][  ]
    Length: 5
    Capacity: 8

---

## ⏱️ Por Que Isso é Lento?

Imagine transferir livros de uma mochila para outra:

- **1ª vez:** Transfere 4 livros
- **2ª vez:** Transfere 8 livros
- **3ª vez:** Transfere 16 livros
- **4ª vez:** Transfere 32 livros

Cada vez você transfere MAIS livros!

---

## 🚀 Solução: Comprar a Mochila Certa Desde o Início

**Cenário 1: Sem planejamento (RUIM)**

    Você vai à biblioteca pegar 100 livros
    Leva uma mochila com 1 bolso
    
    Resultado: Precisa trocar de mochila ~7 vezes!
    Transferências totais: ~200 livros movidos

**Cenário 2: Com planejamento (BOM)**

    Você sabe que vai pegar 100 livros
    Leva uma mochila com 100 bolsos desde o início
    
    Resultado: Nunca troca de mochila!
    Transferências totais: 0

---

## 🎯 Analogia do Prédio em Construção

**Length:** Quantos andares já foram construídos

**Capacity:** Quantos andares a fundação suporta

**Exemplo:**

    Fundação para 20 andares (capacity = 20)
    Construídos 5 andares (length = 5)
    Pode adicionar mais 15 andares sem refazer fundação

**Se ultrapassar 20:**

    Precisa demolir e fazer fundação maior!
    Muito caro e demorado! (realocação)

---

## 💰 Analogia da Conta Bancária

**Length:** Quanto dinheiro você TEM na conta

**Capacity:** Quanto dinheiro a conta PODE guardar antes de precisar mudar de tipo

**Exemplo:**

    Conta básica: limite R$ 10.000 (capacity)
    Saldo atual: R$ 3.000 (length)
    
    Se você depositar R$ 8.000:
    Total ficaria R$ 11.000 (ultrapassa capacity!)
    
    Banco cria conta premium automática:
    - Novo limite: R$ 20.000
    - Transfere os R$ 11.000
    - Fecha conta antiga

---

## 🎪 Por Que Dobrar ao Invés de Aumentar Pouco?

**Estratégia RUIM: Adicionar 1 bolso por vez**

    Mochila com 1 bolso → adiciona 1 livro
    Mochila com 2 bolsos → adiciona 1 livro
    Mochila com 3 bolsos → adiciona 1 livro
    ...
    Para 100 livros: troca de mochila 99 vezes!

**Estratégia BOA: Dobrar os bolsos**

    Mochila com 1 bolso → 2 → 4 → 8 → 16 → 32 → 64 → 128
    Para 100 livros: troca de mochila apenas 7 vezes!

---

## 🎯 Quando Vale a Pena Pré-alocar?

**Situação 1: Você sabe EXATAMENTE quanto precisa**

    Fazer lista de compras para 10 itens específicos
    → Pegue papel com espaço para 10 itens desde o início

**Situação 2: Você tem uma ESTIMATIVA boa**

    Vai ao mercado, geralmente compra ~30 itens
    → Pegue carrinho grande (40 espaços)

**Situação 3: Você NÃO FAZ IDEIA**

    Passeando no shopping, pode ou não comprar coisas
    → Leve bolsa pequena, se precisar compra sacola depois

---

## 🎲 Exemplo do Mundo Real: Playlist de Música

**Sem pré-alocar:**

    Você cria playlist vazia
    Adiciona música 1 → cria playlist para 1
    Adiciona música 2 → recria playlist para 2
    Adiciona música 3 → recria playlist para 4
    Adiciona música 5 → recria playlist para 8
    ...
    
    Muitas recriações de playlist!

**Com pré-alocação:**

    Você sabe que quer ~50 músicas
    Cria playlist preparada para 50 desde o início
    Adiciona todas as 50 músicas sem recriar!

---

## 🔍 Visualizando Length vs Capacity

**Estacionamento:**

    Capacity: 100 vagas no estacionamento
    Length: 30 carros estacionados
    
    Você pode estacionar mais 70 carros sem construir mais vagas!
    
    Se chegar o 101º carro:
    → Precisa construir novo estacionamento maior!

---

## 📊 Tabela Visual de Crescimento

    Adicionando livros na mochila:
    
    1º livro:  [📚]                    (capacity: 1)
    2º livro:  [📚][📚]                (dobrou! capacity: 2)
    3º livro:  [📚][📚][📚][  ]        (dobrou! capacity: 4)
    5º livro:  [📚][📚][📚][📚][📚][  ][  ][  ]  (dobrou! capacity: 8)
    9º livro:  16 bolsos totais
    17º livro: 32 bolsos totais
    33º livro: 64 bolsos totais

---

## 🎯 Regra de Ouro Simples

**Se você SABE quanto vai adicionar:**
→ Use `make([]int, 0, quantidade)` desde o início

**Se você NÃO SABE quanto vai adicionar:**
→ Use `[]int{}` e deixe Go gerenciar

---

## 💡 Exemplo: Festa de Aniversário

**Sem planejamento (ruim):**

    Você avisa 50 amigos para festa
    Aluga espaço para 5 pessoas
    
    Precisou realocar:
    5 → 10 → 20 → 40 → 80 pessoas
    
    Muito trabalho! Mudou de lugar 4 vezes!

**Com planejamento (bom):**

    Você avisa 50 amigos para festa
    Aluga espaço para 60 pessoas (margem de segurança)
    
    Resultado: Um lugar só, ninguém precisa mudar!

---

## 📌 Resumo Visual Simples

    make([]int, 0, 100) = Mochila vazia com 100 bolsos preparados
    
    make([]int, 5) = Mochila com 5 livros já dentro, 0 bolsos vazios
    
    []int{} = Bolsa pequena, cresce conforme precisa

---

# 📚 Aula 3 - Exercícios e Reflexão

## 🏋️ Exercício 1: Observando Crescimento

Crie um programa que:
1. Comece com um slice vazio: `s := []int{}`
2. Use um loop para adicionar números de 0 a 50
3. A CADA append, verifique se a capacity mudou
4. Quando mudar, imprima: "Cresceu! Length: X, Capacity: Y"
5. No final, mostre quantas vezes houve realocação

---

## 🏋️ Exercício 2: Comparação de Performance

Crie um programa com DUAS funções:

**Função 1: Sem pré-alocação**

    func semPreAlocacao() time.Duration {
        inicio := time.Now()
        s := []int{}
        for i := 0; i < 50000; i++ {
            s = append(s, i)
        }
        return time.Since(inicio)
    }

**Função 2: Com pré-alocação**

    func comPreAlocacao() time.Duration {
        inicio := time.Now()
        s := make([]int, 0, 50000)
        for i := 0; i < 50000; i++ {
            s = append(s, i)
        }
        return time.Since(inicio)
    }

Execute ambas e mostre a diferença de tempo.

---

## 🏋️ Exercício 3: Capacity de Sub-slices

Crie um programa que:
1. Crie um slice: `original := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}`
2. Crie sub-slices em diferentes posições:
   - `sub1 := original[0:3]`
   - `sub2 := original[5:8]`
   - `sub3 := original[8:10]`
3. Para cada sub-slice, exiba: length, capacity e os elementos
4. Explique nos comentários: por que as capacities são diferentes?

---

## 🏋️ Exercício 4: Full Slice Expression

Crie um programa que demonstre a diferença:

**Parte 1: Sem limitar capacity**

    original := []int{1, 2, 3, 4, 5}
    sub := original[0:2]
    sub = append(sub, 99)
    // Mostre original e sub

**Parte 2: Limitando capacity**

    original := []int{1, 2, 3, 4, 5}
    sub := original[0:2:2]
    sub = append(sub, 99)
    // Mostre original e sub

Compare os resultados e explique a diferença nos comentários.

---

## 🤔 Perguntas de Reflexão

### Pergunta 1: Custo de Realocação

Você aprendeu que cada realocação copia TODOS os elementos do slice para um novo array.

Reflita:
- Se você adicionar 10.000 elementos sem pré-alocar, aproximadamente quantas realocações ocorrerão?
- Por que copiar todos os elementos a cada realocação é custoso?
- Em que tipo de aplicação (sistema real) isso poderia causar problemas sérios de performance?

Escreva sua resposta com suas próprias palavras (mínimo 4 linhas).

---

### Pergunta 2: Trade-off de Pré-alocação

Pré-alocar capacity melhora performance, mas também usa mais memória logo no início.

Reflita:
- Quando vale a pena usar memória extra para ter mais performance?
- Quando NÃO vale a pena pré-alocar (seria desperdício)?
- Se você alocar `make([]int, 0, 1000000)` mas usar apenas 10 elementos, qual é o problema?

Escreva sua resposta com suas próprias palavras (mínimo 4 linhas).

---

### Pergunta 3: Estratégia de Crescimento

Go usa crescimento exponencial (dobra para slices pequenos, cresce ~25% para grandes).

Reflita:
- Por que você acha que Go muda a estratégia em 256 elementos?
- O que aconteceria se Go sempre dobrasse, mesmo para slices com milhões de elementos?
- O que aconteceria se Go sempre aumentasse apenas 1 elemento por vez?

Escreva sua resposta com suas próprias palavras (mínimo 5 linhas).

---

### Pergunta 4: Aplicação Real

Imagine que você está desenvolvendo um sistema de chat que armazena mensagens.

Descreva:
- Como você usaria slices para armazenar as mensagens?
- Você pré-alocaria capacity? Por quê ou por que não?
- Se decidir pré-alocar, quanto de capacity inicial escolheria e por quê?
- Como lidaria com o fato de não saber quantas mensagens o usuário enviará?

Escreva sua resposta com suas próprias palavras (mínimo 5 linhas).

---

# 📚 Aula 3 - Performance e Boas Práticas

## ⚡ Performance: Análise Detalhada

### 1. Custo Real de Realocação

**Operações em uma realocação:**

    1. malloc() - Alocar novo array          ~100-500ns
    2. memmove() - Copiar elementos          ~0.5ns por byte
    3. free() - Liberar array antigo         ~50-200ns
    4. Atualizar ponteiro do slice           ~1ns

**Exemplo com 10.000 ints:**

    Tamanho: 10.000 × 8 bytes = 80KB
    Tempo de cópia: 80.000 × 0.5ns = 40.000ns = 40μs
    Tempo total: ~40-50μs por realocação

**Para adicionar 100.000 elementos sem pré-alocar:**

    Realocações: ~17
    Elementos copiados: ~200.000
    Tempo total em cópias: ~1-2ms
    
    Com pré-alocação: 0 realocações, 0 cópias, ~0.1ms

**Ganho: 10-20x mais rápido!**

---

### 2. Impacto no Garbage Collector

Cada realocação cria arrays órfãos que precisam ser coletados:

**Sem pré-alocação:**

    Adicionar 10.000 elementos:
    - Cria ~14 arrays temporários
    - GC precisa limpar todos
    - Pressão adicional no GC
    - Pausas de GC mais frequentes

**Com pré-alocação:**

    - Cria 1 array
    - GC trabalha menos
    - Pausas menores

---

### 3. Cache CPU e Locality

Arrays contíguos melhoram cache hit rate:

**Array realocado frequentemente:**

    - Cache misses frequentes
    - CPU precisa buscar na RAM
    - ~100x mais lento

**Array pré-alocado:**

    - Dados contíguos em memória
    - Cache hits altos
    - Acesso rápido

---

## ✅ Boas Práticas: Guia Definitivo

### Prática 1: Sempre Pré-aloque Para Tamanhos Conhecidos

**RUIM:**

    func processarArquivo(linhas []string) []Resultado {
        resultados := []Resultado{}
        for _, linha := range linhas {
            resultados = append(resultados, processar(linha))
        }
        return resultados
    }

**BOM:**

    func processarArquivo(linhas []string) []Resultado {
        resultados := make([]Resultado, 0, len(linhas))
        for _, linha := range linhas {
            resultados = append(resultados, processar(linha))
        }
        return resultados
    }

**Ganho:** Até 10x mais rápido, menos pressão no GC.

---

### Prática 2: Use Heurísticas Para Estimativas

Quando não sabe o tamanho exato, use estimativas inteligentes:

**Exemplo: Filtrar elementos**

    func filtrarPares(numeros []int) []int {
        // Estima que ~50% serão pares
        resultado := make([]int, 0, len(numeros)/2)
        for _, n := range numeros {
            if n%2 == 0 {
                resultado = append(resultado, n)
            }
        }
        return resultado
    }

**Melhor super-estimar que sub-estimar!**

    // Se espera 100-200 elementos, aloque 250
    slice := make([]T, 0, 250)

---

### Prática 3: Reutilize Slices em Loops

**RUIM (cria novo slice toda vez):**

    for i := 0; i < 1000; i++ {
        temp := []int{}
        for j := 0; j < 100; j++ {
            temp = append(temp, processar(i, j))
        }
        usar(temp)
    }
    // 1000 alocações!

**BOM (reutiliza):**

    temp := make([]int, 0, 100)
    for i := 0; i < 1000; i++ {
        temp = temp[:0]  // Reseta length, mantém capacity
        for j := 0; j < 100; j++ {
            temp = append(temp, processar(i, j))
        }
        usar(temp)
    }
    // 1 alocação!

---

### Prática 4: Use Full Slice Expression Para Segurança

Previna modificações acidentais via sub-slices:

**SEM proteção:**

    func pegarPrimeiros(s []int, n int) []int {
        return s[:n]  // Compartilha memória!
    }
    
    original := []int{1, 2, 3, 4, 5}
    sub := pegarPrimeiros(original, 3)
    sub = append(sub, 99)  // Pode modificar original!

**COM proteção:**

    func pegarPrimeiros(s []int, n int) []int {
        return s[:n:n]  // Limita capacity
    }
    
    original := []int{1, 2, 3, 4, 5}
    sub := pegarPrimeiros(original, 3)
    sub = append(sub, 99)  // Cria novo array, original seguro

---

### Prática 5: Monitore Capacity em Desenvolvimento

Use esse padrão durante desenvolvimento para detectar problemas:

    func processarComMonitoramento(dados []int) []int {
        resultado := []int{}
        realocacoes := 0
        oldCap := cap(resultado)
        
        for _, d := range dados {
            resultado = append(resultado, d*2)
            if cap(resultado) != oldCap {
                realocacoes++
                oldCap = cap(resultado)
            }
        }
        
        if realocacoes > 5 {
            log.Printf("ATENÇÃO: %d realocações! Considere pré-alocar", realocacoes)
        }
        
        return resultado
    }

---

## 🚫 O Que NÃO Fazer

### ❌ Erro 1: Pré-alocar com Length ao Invés de Capacity

**ERRADO:**

    s := make([]int, 100)  // Length=100, todos zeros
    for i := 0; i < 50; i++ {
        s = append(s, i)  // Adiciona DEPOIS dos 100 zeros!
    }
    // Resultado: [0 0 0...0 0 1 2...49] ❌

**CORRETO:**

    s := make([]int, 0, 100)  // Length=0, Capacity=100
    for i := 0; i < 50; i++ {
        s = append(s, i)
    }
    // Resultado: [0 1 2...49] ✓

---

### ❌ Erro 2: Pré-alocar Demais (Memory Bloat)

**PROBLEMA:**

    // Sistema com 10.000 usuários
    type Usuario struct {
        mensagens []Mensagem
    }
    
    func novoUsuario() *Usuario {
        return &Usuario{
            mensagens: make([]Mensagem, 0, 10000),  // 10K mensagens!
        }
    }
    
    // 10.000 usuários × 10.000 mensagens pré-alocadas = 100 milhões!
    // Mas usuário médio tem apenas 10 mensagens!

**SOLUÇÃO:**

    func novoUsuario() *Usuario {
        return &Usuario{
            mensagens: make([]Mensagem, 0, 50),  // Estimativa razoável
        }
    }

---

### ❌ Erro 3: Não Considerar Tamanho do Elemento

**PROBLEMA:**

    type ImagemGrande struct {
        pixels [1000000]byte  // 1MB por struct!
    }
    
    // Pré-alocar 10.000 imagens
    imagens := make([]ImagemGrande, 0, 10000)
    // Alocou 10GB de memória! 😱

**SOLUÇÃO:**

    // Use ponteiros para structs grandes
    imagens := make([]*ImagemGrande, 0, 10000)
    // Alocou apenas 80KB (10.000 × 8 bytes)

---

### ❌ Erro 4: Esquecer que Sub-slice Compartilha Capacity

**PROBLEMA:**

    func lerLinhas(arquivo string) []string {
        conteudo := lerArquivoCompleto(arquivo)  // 1GB
        linhas := strings.Split(conteudo, "\n")
        
        // Retorna apenas primeira linha
        return linhas[0:1]
    }
    
    // BUG: O array de 1GB fica na memória!
    // Porque linhas[0:1] ainda referencia o array grande

**SOLUÇÃO:**

    func lerLinhas(arquivo string) []string {
        conteudo := lerArquivoCompleto(arquivo)
        linhas := strings.Split(conteudo, "\n")
        
        // Copia apenas o necessário
        resultado := make([]string, 1)
        resultado[0] = linhas[0]
        return resultado
    }

---

## 🎯 Padrões Avançados

### Padrão 1: Growing Buffer Pattern

Para processar streams de dados:

    type Buffer struct {
        dados []byte
        pos   int
    }
    
    func (b *Buffer) Write(p []byte) {
        needed := b.pos + len(p)
        if needed > cap(b.dados) {
            // Dobra capacity até caber
            newCap := cap(b.dados) * 2
            if newCap < needed {
                newCap = needed
            }
            novoDados := make([]byte, len(b.dados), newCap)
            copy(novoDados, b.dados)
            b.dados = novoDados
        }
        
        b.dados = b.dados[:needed]
        copy(b.dados[b.pos:], p)
        b.pos = needed
    }

---

### Padrão 2: Batch Allocation Pattern

Para reduzir fragmentação:

    type ObjectPool struct {
        objetos []Object
        livres  []int
    }
    
    func (p *ObjectPool) Alocar() *Object {
        if len(p.livres) > 0 {
            idx := p.livres[len(p.livres)-1]
            p.livres = p.livres[:len(p.livres)-1]
            return &p.objetos[idx]
        }
        
        // Aloca em batch de 100
        oldLen := len(p.objetos)
        p.objetos = append(p.objetos, make([]Object, 100)...)
        
        // Marca 99 como livres
        for i := oldLen + 1; i < len(p.objetos); i++ {
            p.livres = append(p.livres, i)
        }
        
        return &p.objetos[oldLen]
    }

---

### Padrão 3: Copy-on-Write Pattern

Para compartilhar slices com segurança:

    type CowSlice struct {
        dados []int
        refs  int
    }
    
    func (c *CowSlice) Get(i int) int {
        return c.dados[i]
    }
    
    func (c *CowSlice) Set(i, valor int) {
        if c.refs > 1 {
            // Copia se compartilhado
            novoDados := make([]int, len(c.dados))
            copy(novoDados, c.dados)
            c.dados = novoDados
            c.refs = 1
        }
        c.dados[i] = valor
    }

---

## 📊 Benchmarks Reais

### Benchmark 1: Tamanhos Diferentes

    func BenchmarkAppend1K(b *testing.B) {
        for i := 0; i < b.N; i++ {
            s := []int{}
            for j := 0; j < 1000; j++ {
                s = append(s, j)
            }
        }
    }
    
    func BenchmarkAppend1KPrealloc(b *testing.B) {
        for i := 0; i < b.N; i++ {
            s := make([]int, 0, 1000)
            for j := 0; j < 1000; j++ {
                s = append(s, j)
            }
        }
    }

**Resultados típicos:**

    BenchmarkAppend1K-8              50000    35000 ns/op    57344 B/op    10 allocs/op
    BenchmarkAppend1KPrealloc-8     200000     7000 ns/op     8192 B/op     1 allocs/op

**Análise:**
- 5x mais rápido
- 7x menos memória alocada
- 10x menos alocações

---

### Benchmark 2: Structs Pequenas vs Grandes

    type Pequena struct {
        id int
    }
    
    type Grande struct {
        dados [1000]byte
    }
    
    func BenchmarkPequena(b *testing.B) {
        for i := 0; i < b.N; i++ {
            s := []Pequena{}
            for j := 0; j < 1000; j++ {
                s = append(s, Pequena{id: j})
            }
        }
    }
    
    func BenchmarkGrande(b *testing.B) {
        for i := 0; i < b.N; i++ {
            s := []Grande{}
            for j := 0; j < 1000; j++ {
                s = append(s, Grande{})
            }
        }
    }

**Resultados:**

    BenchmarkPequena-8     50000    30000 ns/op
    BenchmarkGrande-8       5000   300000 ns/op

**Lição:** Structs grandes amplificam o custo de realocação!

---

## 🔍 Ferramentas de Análise

### Tool 1: Profiling de Memória

    import _ "net/http/pprof"
    
    go func() {
        log.Println(http.ListenAndServe("localhost:6060", nil))
    }()
    
    // Acesse: http://localhost:6060/debug/pprof/heap

---

### Tool 2: Escape Analysis

    go build -gcflags="-m" main.go
    
    // Mostra se slices escapam para heap

---

### Tool 3: Benchmark com Alocações

    go test -bench=. -benchmem

---

## 💡 Otimizações Específicas

### Otimização 1: String Building

**RUIM:**

    result := ""
    for i := 0; i < 1000; i++ {
        result += strconv.Itoa(i)  // Realoca a cada iteração!
    }

**BOM:**

    var builder strings.Builder
    builder.Grow(5000)  // Pré-aloca capacity
    for i := 0; i < 1000; i++ {
        builder.WriteString(strconv.Itoa(i))
    }
    result := builder.String()

---

### Otimização 2: Byte Slices para I/O

**RUIM:**

    func processarArquivo(nome string) {
        for i := 0; i < 1000; i++ {
            buffer := make([]byte, 4096)  // Aloca toda vez!
            // ler e processar
        }
    }

**BOM:**

    func processarArquivo(nome string) {
        buffer := make([]byte, 4096)  // Aloca uma vez
        for i := 0; i < 1000; i++ {
            // reutiliza buffer
        }
    }

---

### Otimização 3: Slices de Ponteiros para Grandes Structs

**RUIM:**

    type Registro struct {
        dados [10000]byte  // 10KB cada!
    }
    
    registros := make([]Registro, 1000)
    // Aloca 10MB de uma vez

**BOM:**

    registros := make([]*Registro, 0, 1000)
    // Aloca sob demanda
    for cada := range fonte {
        registros = append(registros, &Registro{...})
    }

---

## 📏 Regras de Capacity por Tipo

### Tipos Pequenos (≤ 16 bytes)

    []int, []float64, []bool, []*T
    
    Regra: Pré-aloque agressivamente
    Motivo: Custo de realocação é baixo, mas ainda vale a pena

---

### Tipos Médios (17-256 bytes)

    []struct pequenas
    
    Regra: Sempre pré-aloque se souber tamanho
    Motivo: Realocações começam a ficar caras

---

### Tipos Grandes (> 256 bytes)

    []struct grandes
    
    Regra: SEMPRE pré-aloque ou use ponteiros
    Motivo: Realocações são muito caras

---

## 🎯 Decisão: Quando Pré-alocar

**SEMPRE pré-aloque:**
- Tamanho conhecido ou facilmente estimável
- Structs grandes
- Loops com muitas iterações
- Código sensível a performance

**PODE não pré-alocar:**
- Tamanho completamente imprevisível
- Slices pequenos e temporários
- Código de inicialização (roda uma vez)
- Protótipos e testes

**NUNCA pré-aloque demais:**
- Não aloque milhões se usa milhares
- Não aloque para todos os usuários se só alguns usam
- Não aloque na stack se vai passar para heap

---

## 📌 Checklist Final

- [ ] Identificar slices em hot paths (código executado frequentemente)
- [ ] Medir tamanho típico ou estimar com margem de segurança
- [ ] Usar `make([]T, 0, capacity)` com capacity apropriada
- [ ] Considerar tamanho do tipo T (pequeno vs grande)
- [ ] Reutilizar slices em loops quando possível
- [ ] Usar full slice expression quando compartilhar é arriscado
- [ ] Monitorar alocações com benchmarks
- [ ] Profilear memória em produção

---

## 📊 Tabela de Decisão Rápida

| Situação | Ação | Exemplo |
|----------|------|---------|
| Tamanho exato conhecido | Pré-alocar exato | `make([]T, 0, n)` |
| Estimativa boa (±20%) | Pré-alocar com margem | `make([]T, 0, n*1.2)` |
| Estimativa vaga | Pré-alocar conservador | `make([]T, 0, 100)` |
| Totalmente desconhecido | Não pré-alocar | `[]T{}` |
| Struct > 256 bytes | Usar ponteiros | `[]*T` |
| Loop reutilizando | Alocar fora do loop | `s := make([]T, 0, n)` antes do loop |

---

**Fim da Aula 3: Performance e Boas Práticas**

---

## 🎯 Próximo Passo

Agora que você completou TODAS as 4 etapas da Aula 3, vamos continuar para a **Aula 4: make()** 🚀
    