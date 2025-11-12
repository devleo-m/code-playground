# 📚 Aula 1: Arrays em Go

## O que são Arrays?

Arrays são estruturas de dados que armazenam uma sequência de elementos do mesmo tipo em posições consecutivas na memória. Em Go, arrays têm duas características fundamentais:

1. Tamanho fixo: O tamanho é definido na declaração e não pode ser alterado
2. Tipo específico: O tamanho faz parte do tipo - [5]int e [10]int são tipos diferentes!

---

## 📝 Sintaxe e Declaração

**Forma 1: Declaração com inicialização automática (valores zero)**

    var numeros [5]int  
    var nomes [3]string 
    var flags [4]bool   

Quando você declara um array sem especificar valores, Go automaticamente inicializa cada posição com o valor zero do tipo:
- int → 0
- string → "" (string vazia)
- bool → false

**Forma 2: Declaração com valores iniciais**

    var idades [5]int = [5]int{18, 25, 30, 22, 45}
    notas := [4]float64{7.5, 8.0, 9.2, 6.8}
    dias := [...]string{"Segunda", "Terça", "Quarta"}

O operador ... permite que Go calcule automaticamente o tamanho do array baseado nos valores fornecidos.

**Forma 3: Inicialização parcial**

    numeros := [5]int{1, 2}
    valores := [5]int{0: 10, 2: 30, 4: 50}

Você pode inicializar apenas algumas posições. As posições não especificadas recebem o valor zero.

---

## 🔍 Acessando e Modificando Elementos

    package main
    
    import "fmt"
    
    func main() {
        frutas := [4]string{"Maçã", "Banana", "Laranja", "Uva"}
        
        fmt.Println(frutas[0])
        fmt.Println(frutas[3])
        
        frutas[1] = "Morango"
        fmt.Println(frutas)
        
        tamanho := len(frutas)
        fmt.Println(tamanho)
    }

Pontos importantes:
- A indexação começa em 0 (primeira posição é array[0])
- len(array) retorna o tamanho do array
- Você pode modificar valores usando array[indice] = novoValor

---

## 🔄 Iterando sobre Arrays

**Loop tradicional com índice**

    numeros := [5]int{10, 20, 30, 40, 50}
    
    for i := 0; i < len(numeros); i++ {
        fmt.Printf("Posição %d: %d\n", i, numeros[i])
    }

**Range (forma idiomática em Go)**

    numeros := [5]int{10, 20, 30, 40, 50}
    
    for indice, valor := range numeros {
        fmt.Printf("Posição %d: %d\n", indice, valor)
    }
    
    for _, valor := range numeros {
        fmt.Println(valor)
    }

O range retorna dois valores: índice e valor. Use _ (underscore) para ignorar o índice quando não precisar dele.

---

## 💡 Arrays são Tipos Valor (Value Types)

Conceito crucial: Arrays em Go são copiados quando atribuídos ou passados para funções.

    package main
    
    import "fmt"
    
    func modificarArray(arr [3]int) {
        arr[0] = 999
    }
    
    func main() {
        original := [3]int{1, 2, 3}
        
        modificarArray(original)
        
        fmt.Println(original)
    }

**Saída:** [1 2 3] - NÃO foi modificado!

Por quê? Quando você passa original para a função, Go cria uma cópia completa do array. A função trabalha na cópia, não no original.

---

## 📊 Arrays Multidimensionais

Arrays podem conter outros arrays:

    var matriz [3][3]int = [3][3]int{
        {1, 2, 3},
        {4, 5, 6},
        {7, 8, 9},
    }
    
    fmt.Println(matriz[1][2])
    
    for i := 0; i < len(matriz); i++ {
        for j := 0; j < len(matriz[i]); j++ {
            fmt.Printf("%d ", matriz[i][j])
        }
        fmt.Println()
    }

---

## ⚠️ Limitações dos Arrays

1. Tamanho fixo: Não pode crescer ou diminuir após declaração
2. Tipos diferentes: [5]int ≠ [10]int - não são compatíveis
3. Cópia custosa: Passar arrays grandes para funções copia todos os dados

Exemplo de erro de compilação:

    var a [5]int
    var b [10]int
    a = b  // ERRO: tipos incompatíveis

---

## 🎯 Exemplo Completo: Sistema de Notas

    package main
    
    import "fmt"
    
    func main() {
        notas := [5]float64{7.5, 8.0, 6.5, 9.0, 7.0}
        
        var soma float64 = 0
        for _, nota := range notas {
            soma += nota
        }
        media := soma / float64(len(notas))
        
        fmt.Printf("Média da turma: %.2f\n", media)
        
        maior := notas[0]
        for _, nota := range notas {
            if nota > maior {
                maior = nota
            }
        }
        
        fmt.Printf("Maior nota: %.2f\n", maior)
        
        aprovados := 0
        for _, nota := range notas {
            if nota >= 7.0 {
                aprovados++
            }
        }
        
        fmt.Printf("Alunos aprovados: %d de %d\n", aprovados, len(notas))
    }

**Saída:**

    Média da turma: 7.60
    Maior nota: 9.00
    Alunos aprovados: 4 de 5

---

## 📌 Resumo dos Conceitos-Chave

- Tamanho fixo: Definido na declaração, não muda
- Tipo incluindo tamanho: [5]int é diferente de [10]int
- Valor zero: Arrays são inicializados com valores zero do tipo
- Indexação: Começa em 0, acesso com array[indice]
- len(): Retorna o tamanho do array
- Value type: Arrays são copiados em atribuições
- range: Forma idiomática de iterar em Go

---

# 📚 Aula 1 - Simplificada: Entendendo Arrays

## 🏠 Analogia: Arrays são como Armários com Gavetas Fixas

Imagine um armário com gavetas numeradas. Cada gaveta:
- Tem um número fixo (índice)
- Só pode guardar um tipo específico de objeto
- O armário tem um número fixo de gavetas que não pode mudar

**Exemplo prático:**

Você tem um armário de sapatos com 5 gavetas numeradas de 0 a 4:

    Gaveta 0: Tênis
    Gaveta 1: Chinelo
    Gaveta 2: Bota
    Gaveta 3: Sandália
    Gaveta 4: Sapatênis

Em Go seria:

    sapatos := [5]string{"Tênis", "Chinelo", "Bota", "Sandália", "Sapatênis"}

---

## 🎯 Por que o índice começa em 0?

Pense assim: o índice representa "quantas gavetas você pulou para chegar naquela posição".

- Gaveta 0: você não pulou nenhuma (primeira gaveta)
- Gaveta 1: você pulou 1 gaveta
- Gaveta 2: você pulou 2 gavetas

---

## 🔄 Acessando e Modificando - Mundo Real

Imagine que você quer trocar o chinelo por uma sapatilha:

    sapatos[1] = "Sapatilha"

É como abrir a gaveta 1 e colocar outro item lá dentro.

Para ver o que tem na primeira gaveta:

    fmt.Println(sapatos[0])

---

## 📦 Arrays são Cópias - A Analogia da Fotocópia

Quando você passa um array para uma função, é como tirar uma fotocópia do armário inteiro. A função trabalha com a cópia, não com o armário original.

**Exemplo do mundo real:**

Você tem um caderno de receitas (array original). Seu amigo pede emprestado e você tira fotocópias de todas as páginas para ele. Ele pode riscar, adicionar notas nas cópias dele, mas seu caderno original continua intacto.

    func modificarReceitas(receitas [3]string) {
        receitas[0] = "Nova Receita"
    }
    
    func main() {
        minhasReceitas := [3]string{"Bolo", "Pizza", "Lasanha"}
        modificarReceitas(minhasReceitas)
        fmt.Println(minhasReceitas)
    }

Saída: [Bolo Pizza Lasanha] - o original não mudou!

---

## 🎲 Tamanho Fixo - Como uma Forma de Gelo

Pense em uma forma de gelo com 12 cubinhos. Você não pode adicionar mais 3 cubinhos porque a forma tem tamanho fixo. Se precisar de mais gelo, precisa de outra forma.

Arrays funcionam assim:

    cubos := [12]int{} // 12 posições fixas

Não dá para "adicionar mais uma posição" depois. O tamanho é permanente.

---

## 🔢 Arrays Multidimensionais - Prédio de Apartamentos

Um array multidimensional é como um prédio:

    predio := [3][4]string{
        {"Apto 101", "Apto 102", "Apto 103", "Apto 104"},
        {"Apto 201", "Apto 202", "Apto 203", "Apto 204"},
        {"Apto 301", "Apto 302", "Apto 303", "Apto 304"},
    }

- Primeiro índice: andar do prédio (0, 1, 2)
- Segundo índice: número do apartamento no andar (0, 1, 2, 3)

Para acessar o apartamento 203:

    predio[1][2]

Andar 1 (segundo andar), apartamento 2 (terceira posição).

---

## 💡 Quando Usar Arrays?

**Use arrays quando:**
- Você sabe exatamente quantos elementos precisa
- O tamanho nunca vai mudar
- Você quer garantir um tamanho específico

**Exemplo prático:**
- Dias da semana: sempre 7
- Meses do ano: sempre 12
- Coordenadas x,y: sempre 2 valores

---

## 📝 Resumo Visual

    [5]int → Armário com 5 gavetas de números
    [3]string → Armário com 3 gavetas de textos
    [2][3]int → Prédio com 2 andares, cada andar tem 3 apartamentos

---

# 📚 Aula 1 - Exercícios e Reflexão

## 🏋️ Exercício 1: Criar e Manipular um Array de Temperaturas

Crie um programa que:
1. Declare um array com as temperaturas dos últimos 7 dias: [23, 25, 22, 26, 24, 27, 25]
2. Calcule e exiba a temperatura média
3. Encontre e exiba a temperatura mais alta
4. Conte quantos dias tiveram temperatura acima de 24 graus

---

## 🏋️ Exercício 2: Array de Nomes

Crie um programa que:
1. Declare um array com 5 nomes de pessoas
2. Use um loop para exibir cada nome com sua posição (índice)
3. Modifique o terceiro nome para outro nome de sua escolha
4. Exiba o array modificado

---

## 🏋️ Exercício 3: Matriz (Array Multidimensional)

Crie um programa que:
1. Declare uma matriz 2x3 (2 linhas, 3 colunas) com números de sua escolha
2. Use loops aninhados para exibir todos os elementos da matriz
3. Calcule e exiba a soma de todos os elementos

---

## 🏋️ Exercício 4: Testar Cópia de Arrays

Crie um programa que:
1. Declare um array original: [10, 20, 30]
2. Crie uma função que recebe um array e modifica o primeiro elemento para 999
3. Chame a função passando o array original
4. Exiba o array original após a chamada da função
5. Observe se o valor foi modificado ou não

---

## 🤔 Perguntas de Reflexão

### Pergunta 1: Por que Arrays têm Tamanho Fixo?

Pense sobre: 
- Quais são as vantagens de ter um tamanho fixo?
- Em que situações do mundo real isso é útil?
- Quais são as desvantagens dessa limitação?

Escreva sua resposta com suas próprias palavras (mínimo 3 linhas).

---

### Pergunta 2: Arrays como Value Types

Você aprendeu que arrays são copiados quando passados para funções. 

Reflita:
- Por que Go fez essa escolha de design?
- Quando isso pode ser um problema de performance?
- Como você resolveria se precisasse modificar o array original dentro de uma função?

Escreva sua resposta com suas próprias palavras (mínimo 3 linhas).

---

### Pergunta 3: Aplicação Real

Pense em um problema do seu dia a dia ou de um sistema real que você conhece.

Descreva:
- Uma situação onde usar um array seria ideal
- Por que o tamanho fixo faria sentido nesse caso
- Que tipo de dados você armazenaria nesse array

Escreva sua resposta com suas próprias palavras (mínimo 4 linhas).

---

# 📚 Aula 1 - Performance e Boas Práticas

## ⚡ Performance de Arrays

### 1. Arrays são Armazenados em Memória Contígua

Arrays em Go são armazenados em posições consecutivas na memória. Isso significa:

**Vantagem:** Acesso extremamente rápido a qualquer elemento (O(1) - tempo constante)

    numeros := [1000]int{}
    valor := numeros[500]  // Acesso instantâneo, não importa a posição

**Por que é rápido?** O computador calcula: endereço_base + (índice × tamanho_do_tipo)

---

### 2. Cópia de Arrays - Custo de Performance

Quando você passa um array para uma função, Go copia todos os elementos.

**Exemplo ruim (array grande):**

    func processar(dados [1000000]int) {
        // Copia 1 milhão de inteiros toda vez!
    }
    
    func main() {
        meusDados := [1000000]int{}
        processar(meusDados)  // CÓPIA CUSTOSA!
    }

**Impacto:** 
- 1 milhão de ints = ~8MB copiados
- Tempo desperdiçado
- Memória duplicada

---

### 3. Solução: Use Ponteiros para Arrays Grandes

Se precisa passar arrays grandes, use ponteiros:

    func processar(dados *[1000000]int) {
        // Copia apenas o endereço de memória (8 bytes)
    }
    
    func main() {
        meusDados := [1000000]int{}
        processar(&meusDados)  // Passa referência
    }

**Quando usar ponteiros:**
- Arrays com mais de 100 elementos
- Quando você precisa modificar o array original
- Em funções chamadas frequentemente

---

## ✅ Boas Práticas

### Prática 1: Use Arrays Apenas Quando o Tamanho é Conhecido e Fixo

**BOM:**

    diasDaSemana := [7]string{"Dom", "Seg", "Ter", "Qua", "Qui", "Sex", "Sáb"}
    coordenadas := [2]float64{10.5, 20.3}
    rgb := [3]int{255, 128, 0}

**RUIM:**

    usuarios := [100]string{}  // E se precisar de 101 usuários?
    produtos := [50]Produto{}  // Tamanho arbitrário, use slice!

**Regra:** Se o tamanho pode mudar, NÃO use array!

---

### Prática 2: Prefira Range ao Invés de For Tradicional

**BOM (idiomático):**

    numeros := [5]int{10, 20, 30, 40, 50}
    
    for indice, valor := range numeros {
        fmt.Printf("%d: %d\n", indice, valor)
    }

**ACEITÁVEL (mas menos Go-like):**

    for i := 0; i < len(numeros); i++ {
        fmt.Printf("%d: %d\n", i, numeros[i])
    }

**Por que range é melhor?**
- Mais legível
- Menos propenso a erros (não precisa se preocupar com len())
- É a forma idiomática em Go

---

### Prática 3: Inicialize Arrays Corretamente

**BOM:**

    notas := [5]float64{7.5, 8.0, 9.0, 6.5, 7.8}

**BOM (valores zero intencionais):**

    contador := [10]int{}  // Todos zero intencionalmente

**RUIM:**

    var notas [5]float64
    notas[0] = 7.5
    notas[1] = 8.0
    notas[2] = 9.0
    // Esqueceu de inicializar [3] e [4]!

---

### Prática 4: Evite Arrays Multidimensionais Grandes

**RUIM:**

    matriz := [1000][1000]int{}  // 1 milhão de ints = ~8MB na stack!

**BOM:**

    // Use slice de slices (aprenderá na próxima aula)
    matriz := make([][]int, 1000)
    for i := range matriz {
        matriz[i] = make([]int, 1000)
    }

**Por quê?** Arrays grandes na stack podem causar stack overflow.

---

## 🚫 O Que NÃO Fazer

### ❌ Erro 1: Acessar Índice Fora dos Limites

    numeros := [3]int{10, 20, 30}
    valor := numeros[5]  // PANIC! Índice fora do range

**Como evitar:**

    if indice < len(numeros) {
        valor := numeros[indice]
    }

---

### ❌ Erro 2: Comparar Arrays de Tamanhos Diferentes

    a := [3]int{1, 2, 3}
    b := [5]int{1, 2, 3, 4, 5}
    
    if a == b {  // ERRO DE COMPILAÇÃO!
        // Tipos diferentes!
    }

---

### ❌ Erro 3: Tentar Adicionar Elementos

    numeros := [3]int{1, 2, 3}
    numeros = append(numeros, 4)  // ERRO! append não funciona em arrays

**Solução:** Use slices (próxima aula)!

---

## 🎯 Quando Usar Arrays vs Slices

### Use ARRAYS quando:
- Tamanho fixo e conhecido (dias da semana, meses, coordenadas)
- Pequeno número de elementos (< 100)
- Performance crítica e tamanho constante
- Quer garantias de tamanho em tempo de compilação

### Use SLICES quando:
- Tamanho pode mudar
- Não sabe quantos elementos terá
- Precisa adicionar/remover elementos
- Maioria dos casos! (Slices são mais comuns)

---

## 📊 Comparação de Performance

**Cenário 1: Array pequeno (10 elementos)**

    // Array - ÓTIMO
    func processar(dados [10]int) {
        // Copia 80 bytes - rápido
    }

**Cenário 2: Array médio (1000 elementos)**

    // Array - CONSIDERE PONTEIRO
    func processar(dados *[1000]int) {
        // Copia apenas 8 bytes (ponteiro)
    }

**Cenário 3: Array grande (1 milhão de elementos)**

    // Array - EVITE!
    // Use slice (próxima aula)

---

## 💡 Otimizações Práticas

### 1. Pré-calcule Tamanhos Conhecidos

**BOM:**

    const DIAS_NO_MES = 30
    vendas := [DIAS_NO_MES]float64{}

**RUIM:**

    vendas := [30]float64{}  // Magic number

---

### 2. Use Arrays para Lookup Tables

**ÓTIMO uso de arrays:**

    diasPorMes := [12]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
    diasEmMarco := diasPorMes[2]  // Acesso O(1)

---

### 3. Cache de Resultados com Arrays

    var cache [100]int  // Cache de resultados calculados
    
    func obterValor(indice int) int {
        if cache[indice] != 0 {
            return cache[indice]  // Já calculado
        }
        
        resultado := calcularValorCaro(indice)
        cache[indice] = resultado
        return resultado
    }

---

## 📌 Checklist de Boas Práticas

- [ ] Usar arrays apenas para tamanhos fixos e conhecidos
- [ ] Preferir range ao invés de for tradicional
- [ ] Usar ponteiros para arrays > 100 elementos
- [ ] Inicializar arrays explicitamente
- [ ] Evitar magic numbers (use constantes)
- [ ] Nunca acessar índices fora do range
- [ ] Considerar slices para dados dinâmicos
- [ ] Arrays pequenos < 1KB são seguros na stack

---

**Fim da Aula 1: Performance e Boas Práticas**

---

## 🎯 Próximo Passo

Agora que você completou todas as etapas da Aula 1, por favor responda aos exercícios e perguntas de reflexão.

Envie suas respostas e eu farei a **Análise de Desempenho** completa, identificando seus pontos fortes e áreas que precisam de mais atenção!

**Aguardando suas respostas para continuar...** 📝