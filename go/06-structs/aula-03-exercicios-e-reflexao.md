# Módulo 6: Structs em Go

## Aula 3 - Exercícios e Reflexão

Agora é hora de colocar a mão na massa! Complete os exercícios abaixo e reflita sobre as questões propostas.

---

## Exercícios Práticos

### Exercício 1: Criar Struct de Produto

Crie um arquivo `exercicio1.go` e implemente o seguinte:

**Tarefa:**

1. Crie uma struct `Produto` com os seguintes campos:

   - `Nome` (string)
   - `Preco` (float64)
   - `Quantidade` (int)
   - `Disponivel` (bool)

2. Crie uma função `NovoProduto` que recebe nome, preço e quantidade e retorna um `*Produto` inicializado.

3. Crie métodos para a struct:

   - `CalcularTotal()` - retorna preço × quantidade
   - `Vender(quantidade int)` - reduz a quantidade disponível (use ponteiro)
   - `Repor(quantidade int)` - aumenta a quantidade disponível (use ponteiro)
   - `EstaDisponivel()` - retorna true se quantidade > 0

4. No `main`, crie alguns produtos, teste os métodos e exiba os resultados.

**Exemplo de uso esperado:**

```go
produto := NovoProduto("Notebook", 2500.00, 5)
fmt.Println(produto.CalcularTotal()) // 12500.00
produto.Vender(2)
fmt.Println(produto.Quantidade) // 3
fmt.Println(produto.EstaDisponivel()) // true
```

---

### Exercício 2: Sistema de Biblioteca

Crie um arquivo `exercicio2.go` e implemente um sistema de biblioteca:

**Tarefa:**

1. Crie uma struct `Livro` com:

   - `Titulo` (string)
   - `Autor` (string)
   - `AnoPublicacao` (int)
   - `Emprestado` (bool)

2. Crie uma struct `Biblioteca` com:

   - `Nome` (string)
   - `Livros` ([]Livro) - slice de livros

3. Implemente métodos para `Biblioteca`:

   - `AdicionarLivro(livro Livro)` - adiciona um livro à biblioteca
   - `BuscarPorTitulo(titulo string) *Livro` - busca um livro pelo título
   - `EmprestarLivro(titulo string) error` - marca um livro como emprestado
   - `DevolverLivro(titulo string) error` - marca um livro como disponível
   - `ListarLivrosDisponiveis() []Livro` - retorna apenas livros disponíveis

4. No `main`, crie uma biblioteca, adicione alguns livros, empreste e devolva livros, e liste os disponíveis.

**Dica:** Use ponteiros quando precisar modificar a struct.

---

### Exercício 3: Calculadora com Histórico

Crie um arquivo `exercicio3.go` e implemente uma calculadora que guarda histórico:

**Tarefa:**

1. Crie uma struct `Operacao` com:

   - `Tipo` (string) - "soma", "subtracao", etc.
   - `A` (float64)
   - `B` (float64)
   - `Resultado` (float64)

2. Crie uma struct `Calculadora` com:

   - `Historico` ([]Operacao) - slice de operações

3. Implemente métodos:

   - `Soma(a, b float64) float64` - soma e adiciona ao histórico
   - `Subtracao(a, b float64) float64` - subtrai e adiciona ao histórico
   - `Multiplicacao(a, b float64) float64` - multiplica e adiciona ao histórico
   - `Divisao(a, b float64) (float64, error)` - divide e adiciona ao histórico (retorna erro se divisão por zero)
   - `MostrarHistorico()` - exibe todas as operações
   - `LimparHistorico()` - limpa o histórico

4. No `main`, teste todas as operações e exiba o histórico.

**Exemplo:**

```go
calc := Calculadora{}
calc.Soma(10, 5)
calc.Multiplicacao(3, 4)
calc.MostrarHistorico()
```

---

### Exercício 4: Comparação de Structs

Crie um arquivo `exercicio4.go` e explore comparação de structs:

**Tarefa:**

1. Crie uma struct `Ponto` com campos `X` e `Y` (ambos int).

2. Crie uma struct `Retangulo` com campos que sejam comparáveis.

3. Crie uma struct `Circulo` que contenha um slice (isso a torna não comparável).

4. Demonstre:

   - Que `Ponto` pode ser comparado (==)
   - Que `Retangulo` pode ser comparado
   - Que `Circulo` NÃO pode ser comparado (comente o código que causa erro)

5. Crie uma função `SaoIguais(p1, p2 Ponto) bool` que compara dois pontos manualmente.

6. Teste todas as comparações no `main`.

**Dica:** Lembre-se que structs com slices, maps ou funções não são comparáveis.

---

## Perguntas de Reflexão

### Reflexão 1: Por Que Agrupar Dados?

Você aprendeu que structs agrupam dados relacionados.

**Pergunta:**

1. Pense em um programa que você já escreveu (ou imagine um). Quais dados poderiam ser agrupados em structs? Por quê?

2. Qual é a diferença prática entre ter variáveis separadas (`var nome string; var idade int`) e ter uma struct (`type Pessoa struct { Nome string; Idade int }`)? Quando cada abordagem é melhor?

3. Em que situações você NÃO deveria usar structs? Dê exemplos.

4. Como structs ajudam a tornar o código mais legível e manutenível?

**Sua resposta deve incluir:** Exemplos concretos, análise sobre quando usar structs, e pensamento crítico sobre organização de código.

---

### Reflexão 2: Métodos vs. Funções

Você aprendeu que métodos são funções associadas a structs.

**Pergunta:**

1. Qual é a diferença entre um método e uma função normal? Quando você escolheria usar um método em vez de uma função?

2. Pense em um exemplo: você tem uma struct `ContaBancaria` com saldo. Você criaria um método `Sacar(valor float64)` ou uma função `Sacar(conta *ContaBancaria, valor float64)`? Por quê?

3. Quando você usa receptor por valor `(p Pessoa)` vs. receptor por ponteiro `(p *Pessoa)`? Qual é o impacto na performance e no comportamento?

4. Em Go, métodos podem ser definidos em qualquer tipo, não apenas structs. Por que então structs são tão comuns para métodos?

**Sua resposta deve demonstrar:** Compreensão da diferença entre métodos e funções, entendimento de quando usar cada um, e conhecimento sobre receptores.

---

### Reflexão 3: Structs e Design de Software

Structs são fundamentais para modelar domínios em software.

**Pergunta:**

1. Pense em um sistema real (ex: e-commerce, rede social, banco). Quais structs você criaria para representar as entidades principais? Por exemplo, em um e-commerce: Produto, Cliente, Pedido, etc.

2. Como você decidiria quais campos colocar em cada struct? O que faz um campo "pertencer" a uma struct?

3. Structs em Go são diferentes de classes em linguagens orientadas a objetos. Quais são as principais diferenças? Quais são as vantagens e desvantagens da abordagem de Go?

4. Como structs ajudam a criar código mais modular e reutilizável?

**Sua resposta deve incluir:** Exemplos de modelagem de domínio, análise comparativa com OOP, e pensamento sobre design de software.

---

### Reflexão 4: Performance e Eficiência

Structs têm implicações de performance.

**Pergunta:**

1. Quando você passa uma struct por valor vs. por ponteiro? Qual é o impacto na memória e performance?

2. Para uma struct pequena (ex: `Ponto` com 2 campos int), faz diferença passar por valor ou ponteiro? E para uma struct grande (ex: 20 campos)?

3. Métodos com receptor por valor criam cópias. Quando isso é um problema? Quando é aceitável?

4. Como você decide se uma struct deve ser passada por valor ou ponteiro? Quais são os critérios?

**Sua resposta deve demonstrar:** Compreensão de performance, análise de trade-offs, e capacidade de tomar decisões informadas.

---

### Reflexão 5: Organização e Manutenibilidade

Structs ajudam a organizar código.

**Pergunta:**

1. Como structs tornam o código mais fácil de entender? Dê exemplos específicos.

2. Imagine que você precisa modificar um programa antigo que não usa structs (apenas variáveis separadas). Quais seriam os desafios? Como structs ajudariam?

3. Quando você trabalha em equipe, como structs facilitam a colaboração? Como outros desenvolvedores se beneficiam de structs bem definidas?

4. Qual é a relação entre structs bem definidas e bugs? Como structs podem prevenir erros?

**Sua resposta deve incluir:** Análise sobre manutenibilidade, exemplos de colaboração, e pensamento sobre qualidade de código.

---

## Como Entregar

Crie os seguintes arquivos na pasta `06-structs/`:

1. **`exercicio1.go`** - Struct de Produto com métodos
2. **`exercicio2.go`** - Sistema de Biblioteca
3. **`exercicio3.go`** - Calculadora com Histórico
4. **`exercicio4.go`** - Comparação de Structs
5. **`reflexoes.md`** - Respostas às 5 perguntas de reflexão

**Importante:**

- Compile e teste todos os códigos para garantir que funcionam
- Use `go run` para testar cada exercício
- Seja honesto nas reflexões - não há resposta "errada", mas há respostas que demonstram mais ou menos compreensão
- Use exemplos reais das suas experiências quando possível
- Pense criticamente sobre os conceitos

---

## Dicas para os Exercícios

### Dica para Exercício 1:

```go
// Lembre-se de usar ponteiro quando modificar
func (p *Produto) Vender(quantidade int) {
    p.Quantidade -= quantidade
    if p.Quantidade == 0 {
        p.Disponivel = false
    }
}
```

### Dica para Exercício 2:

```go
// Para buscar, itere sobre o slice
func (b *Biblioteca) BuscarPorTitulo(titulo string) *Livro {
    for i := range b.Livros {
        if b.Livros[i].Titulo == titulo {
            return &b.Livros[i]  // Retorna ponteiro
        }
    }
    return nil
}
```

### Dica para Exercício 3:

```go
// Crie a operação antes de calcular
func (c *Calculadora) Soma(a, b float64) float64 {
    resultado := a + b
    op := Operacao{
        Tipo:     "soma",
        A:        a,
        B:        b,
        Resultado: resultado,
    }
    c.Historico = append(c.Historico, op)
    return resultado
}
```

### Dica para Exercício 4:

```go
// Struct comparável
type Ponto struct {
    X int
    Y int
}

// Struct NÃO comparável (tem slice)
type Circulo struct {
    Raio float64
    Pontos []Ponto  // Slice torna não comparável!
}

// Isso causará erro de compilação:
// c1 := Circulo{Raio: 5.0}
// c2 := Circulo{Raio: 5.0}
// fmt.Println(c1 == c2)  // ERRO!
```

---

## Checklist de Revisão

Antes de entregar, verifique:

- [ ] Todos os códigos compilam sem erros
- [ ] Todos os métodos estão implementados
- [ ] Usei ponteiros quando necessário para modificar
- [ ] Testei todos os casos (valores normais, casos extremos)
- [ ] As reflexões são honestas e pensadas
- [ ] Incluí exemplos concretos nas reflexões
- [ ] O código está bem organizado e legível

---

Após completar os exercícios e reflexões, envie suas respostas para que eu possa analisar seu desempenho e fornecer feedback construtivo!
