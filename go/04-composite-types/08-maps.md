# 📘 CURSO DE PROGRAMAÇÃO GO - AULA 8

---

## **Aula 8: Maps (Mapas/Dicionários)**

### 🎯 **Objetivos da Aula**
- Compreender o que são maps e como funcionam internamente
- Dominar operações básicas: criar, adicionar, buscar, deletar
- Entender o comportamento de maps com diferentes tipos de chave
- Aprender sobre iteração e ordenação
- Identificar boas práticas e armadilhas comuns

---

### 📚 **1. Revisão Rápida das Aulas Anteriores**

Estruturas de dados que já conhecemos:
- **Arrays**: Tamanho fixo, acesso por índice numérico
- **Slices**: Tamanho dinâmico, acesso por índice numérico
- **Strings**: Sequência imutável de bytes/runes

**Limitação:** Todas acessam elementos por **posição numérica**

**Agora:** Maps permitem acessar valores usando **qualquer chave comparável**!

---

### 🗺️ **2. O Que São Maps?**

**Map** (mapa ou dicionário) é uma estrutura de dados que armazena pares **chave-valor**. É como uma tabela de busca onde você usa uma chave para encontrar um valor rapidamente.

**Analogia:** Dicionário de idiomas
- **Chave**: Palavra em português
- **Valor**: Tradução em inglês
- Exemplo: "olá" → "hello"

**Características:**
1. **Chaves únicas**: Cada chave aparece apenas uma vez
2. **Busca O(1)**: Acesso extremamente rápido (em média)
3. **Desordenados**: A ordem de iteração não é garantida
4. **Tipos flexíveis**: Chave e valor podem ser de tipos diferentes
5. **Referência**: Maps são tipos de referência (como slices)

---

### 💻 **3. Criando Maps**

#### **Exemplo 1: Declaração Básica**

package main

import "fmt"

func main() {
    // Método 1: var (inicializado como nil)
    var mapa1 map[string]int
    fmt.Println("mapa1:", mapa1)
    fmt.Println("mapa1 == nil?", mapa1 == nil)
    
    // Método 2: make (pronto para uso)
    mapa2 := make(map[string]int)
    fmt.Println("mapa2:", mapa2)
    fmt.Println("mapa2 == nil?", mapa2 == nil)
    
    // Método 3: Literal (com valores iniciais)
    mapa3 := map[string]int{
        "maçãs":   5,
        "bananas": 3,
        "laranjas": 7,
    }
    fmt.Println("mapa3:", mapa3)
    
    // Método 4: make com capacidade inicial
    mapa4 := make(map[string]int, 100)  // Pré-aloca para ~100 elementos
    fmt.Println("mapa4:", mapa4)
}

**Saída:**

mapa1: map[]
mapa1 == nil? true
mapa2: map[]
mapa2 == nil? false
mapa3: map[bananas:3 laranjas:7 maçãs:5]
mapa4: map[]

**IMPORTANTE:** 
- `var mapa map[K]V` cria um map **nil** (não pode ser usado!)
- `make(map[K]V)` cria um map **vazio** (pronto para uso)
- Map nil ≠ Map vazio!

---

#### **Exemplo 2: Map Nil vs Map Vazio**

package main

import "fmt"

func main() {
    var nilMap map[string]int
    emptyMap := make(map[string]int)
    
    // ✅ Leitura funciona em ambos
    fmt.Println("nilMap['chave']:", nilMap["chave"])      // 0 (valor zero)
    fmt.Println("emptyMap['chave']:", emptyMap["chave"]) // 0 (valor zero)
    
    // ❌ Escrita em nil map causa PANIC!
    // nilMap["chave"] = 10  // PANIC: assignment to entry in nil map
    
    // ✅ Escrita em empty map funciona
    emptyMap["chave"] = 10
    fmt.Println("emptyMap:", emptyMap)
}

**Saída:**

nilMap['chave']: 0
emptyMap['chave']: 0
emptyMap: map[chave:10]

---

### 🔧 **4. Operações Básicas**

#### **Exemplo 3: Adicionar, Buscar e Atualizar**

package main

import "fmt"

func main() {
    // Criar map
    idades := make(map[string]int)
    
    // Adicionar elementos
    idades["Alice"] = 25
    idades["Bob"] = 30
    idades["Carlos"] = 28
    
    fmt.Println("Map inicial:", idades)
    
    // Buscar elemento
    idadeAlice := idades["Alice"]
    fmt.Println("Idade de Alice:", idadeAlice)
    
    // Buscar elemento inexistente
    idadeDesconhecido := idades["Desconhecido"]
    fmt.Println("Idade de Desconhecido:", idadeDesconhecido)  // 0 (valor zero)
    
    // Atualizar elemento existente
    idades["Alice"] = 26
    fmt.Println("Nova idade de Alice:", idades["Alice"])
    
    // Verificar se chave existe (IMPORTANTE!)
    idade, existe := idades["Bob"]
    if existe {
        fmt.Printf("Bob tem %d anos\n", idade)
    } else {
        fmt.Println("Bob não encontrado")
    }
    
    idade2, existe2 := idades["Desconhecido"]
    if existe2 {
        fmt.Printf("Desconhecido tem %d anos\n", idade2)
    } else {
        fmt.Println("Desconhecido não encontrado")
    }
}

**Saída:**

Map inicial: map[Alice:25 Bob:30 Carlos:28]
Idade de Alice: 25
Idade de Desconhecido: 0
Nova idade de Alice: 26
Bob tem 30 anos
Desconhecido não encontrado

**Padrão idiomático:**
- `valor := mapa[chave]` → Retorna valor ou zero se não existir
- `valor, ok := mapa[chave]` → ok é true se chave existir

---

#### **Exemplo 4: Deletar Elementos**

package main

import "fmt"

func main() {
    frutas := map[string]int{
        "maçã":    10,
        "banana":  5,
        "laranja": 8,
        "uva":     12,
    }
    
    fmt.Println("Antes:", frutas)
    
    // Deletar elemento existente
    delete(frutas, "banana")
    fmt.Println("Após deletar 'banana':", frutas)
    
    // Deletar elemento inexistente (não causa erro)
    delete(frutas, "melancia")
    fmt.Println("Após deletar 'melancia' (não existe):", frutas)
    
    // Verificar tamanho
    fmt.Println("Número de elementos:", len(frutas))
    
    // Limpar todo o map (deletar tudo)
    for chave := range frutas {
        delete(frutas, chave)
    }
    fmt.Println("Map limpo:", frutas)
    fmt.Println("Tamanho após limpar:", len(frutas))
}

**Saída:**

Antes: map[banana:5 laranja:8 maçã:10 uva:12]
Após deletar 'banana': map[laranja:8 maçã:10 uva:12]
Após deletar 'melancia' (não existe): map[laranja:8 maçã:10 uva:12]
Número de elementos: 3
Map limpo: map[]
Tamanho após limpar: 0

---

### 🔄 **5. Iteração em Maps**

#### **Exemplo 5: Iterando Chave-Valor**

package main

import "fmt"

func main() {
    notas := map[string]float64{
        "Matemática": 8.5,
        "Português":  7.0,
        "História":   9.0,
        "Física":     6.5,
    }
    
    // Iterar sobre chave e valor
    fmt.Println("=== CHAVE E VALOR ===")
    for disciplina, nota := range notas {
        fmt.Printf("%s: %.1f\n", disciplina, nota)
    }
    
    // Iterar apenas sobre chaves
    fmt.Println("\n=== APENAS CHAVES ===")
    for disciplina := range notas {
        fmt.Println(disciplina)
    }
    
    // Iterar apenas sobre valores (raro, mas possível)
    fmt.Println("\n=== APENAS VALORES ===")
    for _, nota := range notas {
        fmt.Printf("%.1f\n", nota)
    }
    
    // Calcular média
    total := 0.0
    for _, nota := range notas {
        total += nota
    }
    media := total / float64(len(notas))
    fmt.Printf("\nMédia: %.2f\n", media)
}

**Saída (ordem pode variar!):**

=== CHAVE E VALOR ===
Física: 6.5
Matemática: 8.5
Português: 7.0
História: 9.0

=== APENAS CHAVES ===
Matemática
Português
História
Física

=== APENAS VALORES ===
8.5
7.0
9.0
6.5

Média: 7.75

**⚠️ IMPORTANTE:** A ordem de iteração é **aleatória** e pode mudar entre execuções!

---

#### **Exemplo 6: Ordem de Iteração Não é Garantida**

package main

import "fmt"

func main() {
    numeros := map[int]string{
        1: "um",
        2: "dois",
        3: "três",
        4: "quatro",
        5: "cinco",
    }
    
    fmt.Println("Execução 1:")
    for k, v := range numeros {
        fmt.Printf("%d: %s\n", k, v)
    }
    
    fmt.Println("\nExecute o programa várias vezes - a ordem muda!")
}

**Nota:** Go **intencionalmente** randomiza a ordem para evitar que programadores dependam dela!

---

### 📊 **6. Tipos de Chaves Permitidos**

#### **Exemplo 7: Diferentes Tipos de Chaves**

package main

import "fmt"

func main() {
    // ✅ Chaves string (mais comum)
    mapa1 := map[string]int{"a": 1, "b": 2}
    fmt.Println("String keys:", mapa1)
    
    // ✅ Chaves int
    mapa2 := map[int]string{1: "um", 2: "dois"}
    fmt.Println("Int keys:", mapa2)
    
    // ✅ Chaves float64
    mapa3 := map[float64]string{3.14: "pi", 2.71: "e"}
    fmt.Println("Float keys:", mapa3)
    
    // ✅ Chaves bool
    mapa4 := map[bool]string{true: "verdadeiro", false: "falso"}
    fmt.Println("Bool keys:", mapa4)
    
    // ✅ Chaves struct (se todos os campos são comparáveis)
    type Coordenada struct {
        X, Y int
    }
    mapa5 := map[Coordenada]string{
        {0, 0}: "origem",
        {1, 1}: "diagonal",
    }
    fmt.Println("Struct keys:", mapa5)
    
    // ✅ Chaves array (não slice!)
    mapa6 := map[[2]int]string{
        {1, 2}: "par",
        {3, 4}: "outro par",
    }
    fmt.Println("Array keys:", mapa6)
    
    // ❌ Chaves slice (ERRO DE COMPILAÇÃO!)
    // mapa7 := map[[]int]string{}  // ERRO: slice não é comparável
    
    // ❌ Chaves map (ERRO DE COMPILAÇÃO!)
    // mapa8 := map[map[string]int]string{}  // ERRO: map não é comparável
}

**Saída:**

String keys: map[a:1 b:2]
Int keys: map[1:um 2:dois]
Float keys: map[2.71:e 3.14:pi]
Bool keys: map[false:falso true:verdadeiro]
Struct keys: map[{0 0}:origem {1 1}:diagonal]
Array keys: map[[1 2]:par [3 4]:outro par]

**Regra:** Chaves devem ser **comparáveis** (suportar == e !=)

---

### 🎯 **7. Maps como Referências**

#### **Exemplo 8: Maps São Passados por Referência**

package main

import "fmt"

func adicionarItem(m map[string]int, chave string, valor int) {
    m[chave] = valor
}

func main() {
    estoque := map[string]int{
        "maçã":   10,
        "banana": 5,
    }
    
    fmt.Println("Antes:", estoque)
    
    // Passar map para função
    adicionarItem(estoque, "laranja", 8)
    
    fmt.Println("Depois:", estoque)  // Map foi modificado!
}

**Saída:**

Antes: map[banana:5 maçã:10]
Depois: map[banana:5 laranja:8 maçã:10]

**Análise:** Maps são **tipos de referência**, assim como slices. Modificações dentro de funções afetam o map original!

---

#### **Exemplo 9: Copiando Maps**

package main

import "fmt"

func main() {
    original := map[string]int{
        "a": 1,
        "b": 2,
        "c": 3,
    }
    
    // ❌ Atribuição simples NÃO copia!
    referencia := original
    referencia["a"] = 999
    fmt.Println("Original:", original)  // Modificado!
    fmt.Println("Referência:", referencia)
    
    // ✅ Para copiar, precisa iterar
    copia := make(map[string]int)
    for k, v := range original {
        copia[k] = v
    }
    
    copia["b"] = 888
    fmt.Println("\nOriginal:", original)  // Não modificado
    fmt.Println("Cópia:", copia)
}

**Saída:**

Original: map[a:999 b:2 c:3]
Referência: map[a:999 b:2 c:3]

Original: map[a:999 b:2 c:3]
Cópia: map[b:888 a:999 c:3]

---

### 🔍 **8. Padrões Comuns de Uso**

#### **Exemplo 10: Contador de Frequência**

package main

import "fmt"

func contarPalavras(texto string) map[string]int {
    palavras := strings.Fields(texto)
    contagem := make(map[string]int)
    
    for _, palavra := range palavras {
        contagem[palavra]++  // Incrementa (0 se não existir)
    }
    
    return contagem
}

func main() {
    texto := "go é legal go é rápido go é moderno"
    
    frequencia := contarPalavras(texto)
    
    for palavra, count := range frequencia {
        fmt.Printf("%s: %d vezes\n", palavra, count)
    }
}

**Saída:**

go: 3 vezes
é: 3 vezes
legal: 1 vezes
rápido: 1 vezes
moderno: 1 vezes

**Truque:** `map[chave]++` funciona mesmo se a chave não existir (começa do zero)!

---

#### **Exemplo 11: Agrupar Dados**

package main

import "fmt"

type Pessoa struct {
    Nome   string
    Cidade string
}

func agruparPorCidade(pessoas []Pessoa) map[string][]Pessoa {
    grupos := make(map[string][]Pessoa)
    
    for _, pessoa := range pessoas {
        grupos[pessoa.Cidade] = append(grupos[pessoa.Cidade], pessoa)
    }
    
    return grupos
}

func main() {
    pessoas := []Pessoa{
        {"Alice", "São Paulo"},
        {"Bob", "Rio de Janeiro"},
        {"Carlos", "São Paulo"},
        {"Diana", "Rio de Janeiro"},
        {"Eduardo", "São Paulo"},
    }
    
    porCidade := agruparPorCidade(pessoas)
    
    for cidade, grupo := range porCidade {
        fmt.Printf("\n%s (%d pessoas):\n", cidade, len(grupo))
        for _, pessoa := range grupo {
            fmt.Printf("  - %s\n", pessoa.Nome)
        }
    }
}

**Saída:**

São Paulo (3 pessoas):
  - Alice
  - Carlos
  - Eduardo

Rio de Janeiro (2 pessoas):
  - Bob
  - Diana

---

#### **Exemplo 12: Cache/Memoization**

package main

import (
    "fmt"
    "time"
)

// Função "cara" que queremos cachear
func calcularFibonacci(n int) int {
    if n <= 1 {
        return n
    }
    return calcularFibonacci(n-1) + calcularFibonacci(n-2)
}

// Versão com cache
var cache = make(map[int]int)

func fibonacciComCache(n int) int {
    // Verificar se já calculamos
    if valor, existe := cache[n]; existe {
        return valor
    }
    
    // Calcular
    var resultado int
    if n <= 1 {
        resultado = n
    } else {
        resultado = fibonacciComCache(n-1) + fibonacciComCache(n-2)
    }
    
    // Guardar no cache
    cache[n] = resultado
    return resultado
}

func main() {
    n := 40
    
    // Sem cache
    inicio := time.Now()
    resultado1 := calcularFibonacci(n)
    duracao1 := time.Since(inicio)
    
    // Com cache
    inicio = time.Now()
    resultado2 := fibonacciComCache(n)
    duracao2 := time.Since(inicio)
    
    fmt.Printf("Sem cache: %d em %v\n", resultado1, duracao1)
    fmt.Printf("Com cache: %d em %v\n", resultado2, duracao2)
    fmt.Printf("Speedup: %.0fx mais rápido\n", float64(duracao1)/float64(duracao2))
}

**Saída típica:**

Sem cache: 102334155 em 1.2s
Com cache: 102334155 em 50µs
Speedup: 24000x mais rápido

---

#### **Exemplo 13: Set (Conjunto)**

package main

import "fmt"

// Go não tem set nativo, mas map[T]bool simula um!
type Set map[string]bool

func NovoSet() Set {
    return make(Set)
}

func (s Set) Adicionar(item string) {
    s[item] = true
}

func (s Set) Remover(item string) {
    delete(s, item)
}

func (s Set) Contem(item string) bool {
    return s[item]
}

func (s Set) Tamanho() int {
    return len(s)
}

func (s Set) Lista() []string {
    lista := make([]string, 0, len(s))
    for item := range s {
        lista = append(lista, item)
    }
    return lista
}

func main() {
    frutas := NovoSet()
    
    frutas.Adicionar("maçã")
    frutas.Adicionar("banana")
    frutas.Adicionar("maçã")  // Duplicata ignorada
    
    fmt.Println("Contém 'maçã'?", frutas.Contem("maçã"))
    fmt.Println("Contém 'uva'?", frutas.Contem("uva"))
    fmt.Println("Tamanho:", frutas.Tamanho())
    fmt.Println("Itens:", frutas.Lista())
    
    frutas.Remover("banana")
    fmt.Println("Após remover banana:", frutas.Lista())
}

**Saída:**

Contém 'maçã'? true
Contém 'uva'? false
Tamanho: 2
Itens: [maçã banana]
Após remover banana: [maçã]

---

### 🎨 **9. Maps Aninhados (Nested Maps)**

#### **Exemplo 14: Map de Maps**

package main

import "fmt"

func main() {
    // Estrutura: país -> cidade -> população
    populacao := map[string]map[string]int{
        "Brasil": {
            "São Paulo":       12_000_000,
            "Rio de Janeiro":  6_700_000,
            "Brasília":        3_000_000,
        },
        "Argentina": {
            "Buenos Aires": 3_000_000,
            "Córdoba":      1_500_000,
        },
    }
    
    // Acessar dados
    fmt.Println("População de São Paulo:", populacao["Brasil"]["São Paulo"])
    
    // Adicionar nova cidade
    populacao["Brasil"]["Belo Horizonte"] = 2_500_000
    
    // Adicionar novo país
    populacao["Chile"] = make(map[string]int)
    populacao["Chile"]["Santiago"] = 5_600_000
    
    // Iterar
    for pais, cidades := range populacao {
        fmt.Printf("\n%s:\n", pais)
        for cidade, pop := range cidades {
            fmt.Printf("  %s: %d habitantes\n", cidade, pop)
        }
    }
}

**Saída:**

População de São Paulo: 12000000

Brasil:
  Belo Horizonte: 2500000
  São Paulo: 12000000
  Rio de Janeiro: 6700000
  Brasília: 3000000

Argentina:
  Buenos Aires: 3000000
  Córdoba: 1500000

Chile:
  Santiago: 5600000

---

### 📈 **10. Ordenando Maps**

#### **Exemplo 15: Ordenar por Chaves**

package main

import (
    "fmt"
    "sort"
)

func main() {
    notas := map[string]float64{
        "Maria":   8.5,
        "João":    7.0,
        "Ana":     9.5,
        "Pedro":   6.5,
        "Carla":   8.0,
    }
    
    // Maps não são ordenados, mas podemos ordenar as chaves!
    
    // 1. Extrair chaves
    nomes := make([]string, 0, len(notas))
    for nome := range notas {
        nomes = append(nomes, nome)
    }
    
    // 2. Ordenar chaves
    sort.Strings(nomes)
    
    // 3. Iterar em ordem
    fmt.Println("Notas em ordem alfabética:")
    for _, nome := range nomes {
        fmt.Printf("%s: %.1f\n", nome, notas[nome])
    }
}

**Saída:**

Notas em ordem alfabética:
Ana: 9.5
Carla: 8.0
João: 7.0
Maria: 8.5
Pedro: 6.5

---

#### **Exemplo 16: Ordenar por Valores**

package main

import (
    "fmt"
    "sort"
)

type Par struct {
    Chave string
    Valor float64
}

func main() {
    notas := map[string]float64{
        "Maria": 8.5,
        "João":  7.0,
        "Ana":   9.5,
        "Pedro": 6.5,
        "Carla": 8.0,
    }
    
    // 1. Converter map em slice de pares
    pares := make([]Par, 0, len(notas))
    for nome, nota := range notas {
        pares = append(pares, Par{nome, nota})
    }
    
    // 2. Ordenar slice por valor
    sort.Slice(pares, func(i, j int) bool {
        return pares[i].Valor > pares[j].Valor  // Ordem decrescente
    })
    
    // 3. Exibir
    fmt.Println("Ranking de notas:")
    for i, par := range pares {
        fmt.Printf("%d. %s: %.1f\n", i+1, par.Chave, par.Valor)
    }
}

**Saída:**

Ranking de notas:
1. Ana: 9.5
2. Maria: 8.5
3. Carla: 8.0
4. João: 7.0
5. Pedro: 6.5

---

### 🔒 **11. Concorrência e Maps**

#### **Exemplo 17: Maps NÃO São Thread-Safe**

package main

import (
    "fmt"
    "sync"
)

func main() {
    contador := make(map[string]int)
    var wg sync.WaitGroup
    
    // ❌ PERIGO: Acessos concorrentes causam panic!
    // for i := 0; i < 100; i++ {
    //     wg.Add(1)
    //     go func(n int) {
    //         defer wg.Done()
    //         contador["chave"]++  // RACE CONDITION!
    //     }(i)
    // }
    
    // ✅ SOLUÇÃO 1: Usar mutex
    var mu sync.Mutex
    for i := 0; i < 100; i++ {
        wg.Add(1)
        go func(n int) {
            defer wg.Done()
            mu.Lock()
            contador["chave"]++
            mu.Unlock()
        }(i)
    }
    
    wg.Wait()
    fmt.Println("Contador:", contador["chave"])
}

**Saída:**

Contador: 100

---

#### **Exemplo 18: sync.Map (Map Thread-Safe)**

package main

import (
    "fmt"
    "sync"
)

func main() {
    var m sync.Map
    var wg sync.WaitGroup
    
    // Escrever concorrentemente (seguro!)
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func(n int) {
            defer wg.Done()
            m.Store(fmt.Sprintf("chave%d", n), n*10)
        }(i)
    }
    
    wg.Wait()
    
    // Ler
    m.Range(func(key, value interface{}) bool {
        fmt.Printf("%s: %d\n", key.(string), value.(int))
        return true  // Continuar iteração
    })
    
    // Buscar
    if valor, ok := m.Load("chave5"); ok {
        fmt.Println("\nValor de chave5:", valor)
    }
}

**Saída (ordem pode variar):**

chave0: 0
chave1: 10
chave2: 20
chave3: 30
chave4: 40
chave5: 50
chave6: 60
chave7: 70
chave8: 80
chave9: 90

Valor de chave5: 50

---

### 📚 **12. Resumo dos Conceitos-Chave**

1. **Maps são tabelas hash**: Busca O(1) em média
2. **Chaves devem ser comparáveis**: String, int, struct, array (não slice!)
3. **Map nil vs vazio**: Nil não pode ser escrito, vazio pode
4. **Ordem não garantida**: Iteração é aleatória
5. **Tipo de referência**: Modificações em funções afetam original
6. **Verificar existência**: Use `valor, ok := mapa[chave]`
7. **delete() é seguro**: Não causa erro se chave não existe
8. **Não thread-safe**: Use mutex ou sync.Map para concorrência

---

## **Aula 8 - Simplificada: Entendendo Maps**

### 📖 **A Analogia da Agenda Telefônica**

Imagine uma **agenda telefônica** (daquelas de papel antigamente):

**Array/Slice = Lista Numerada:**
- Posição 0: João
- Posição 1: Maria
- Posição 2: Pedro
- **Problema:** Como encontrar o telefone de "Maria"? Precisa procurar página por página!

**Map = Agenda Alfabética:**
- "João" → 1234-5678
- "Maria" → 9876-5432
- "Pedro" → 5555-1111
- **Vantagem:** Vai direto na letra "M" e encontra Maria instantaneamente!

---

### 🔑 **Chaves e Valores: O Par Perfeito**

**Map** sempre tem dois componentes:

1. **Chave** (Key): O que você usa para buscar
2. **Valor** (Value): O que você quer encontrar

**Exemplos do dia a dia:**

// Tradução
ingles := map[string]string{
    "olá":  "hello",
    "tchau": "goodbye",
}

// Preços
produtos := map[string]float64{
    "maçã":   2.50,
    "banana": 1.80,
}

// Pontuações
jogadores := map[string]int{
    "Alice": 1500,
    "Bob":   1200,
}

---

### ⚠️ **Map Nil: A Armadilha do Iniciante**

**Situação perigosa:**

var mapa map[string]int  // Criou um map NIL!

// ✅ Ler funciona (retorna zero)
valor := mapa["chave"]  // 0

// ❌ Escrever causa PANIC!
mapa["chave"] = 10  // 💥 BOOM! Program crashed!

**Solução sempre segura:**

mapa := make(map[string]int)  //