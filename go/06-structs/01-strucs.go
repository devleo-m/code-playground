# 🎓 Aula 6: Structs em Go

Bem-vindo à Aula 6! Hoje vamos explorar um dos conceitos mais fundamentais e poderosos do Go: **Structs**. Se você lembra das aulas anteriores sobre tipos de dados básicos e funções, agora vamos dar um passo gigante organizando informações complexas de forma elegante e profissional.

## 📋 Revisão Rápida da Aula Anterior

Antes de começarmos, vamos relembrar rapidamente:
- Você aprendeu sobre **tipos de dados básicos** (int, string, bool, etc.)
- Trabalhou com **funções** que recebem parâmetros e retornam valores
- Entendeu como organizar código de forma modular

Agora, imagine que você precise representar informações mais complexas, como dados de uma pessoa, um produto, ou um livro. É aqui que entram as **structs**!

---

## 🏗️ O Que São Structs?

Uma **struct** (estrutura) é um tipo de dado personalizado que agrupa campos relacionados sob um único nome. Pense nela como um "molde" ou "blueprint" para criar objetos com características específicas.

### Sintaxe Básica

    type NomeDaStruct struct {
        campo1 tipoDado1
        campo2 tipoDado2
        campo3 tipoDado3
    }

### Exemplo Prático: Representando uma Pessoa

    package main

    import "fmt"

    // Definindo a struct Pessoa
    type Pessoa struct {
        Nome  string
        Idade int
        Email string
    }

    func main() {
        // Criando uma instância da struct
        pessoa1 := Pessoa{
            Nome:  "Maria Silva",
            Idade: 28,
            Email: "maria@example.com",
        }
        
        fmt.Println(pessoa1)
        fmt.Println("Nome:", pessoa1.Nome)
        fmt.Println("Idade:", pessoa1.Idade)
    }

**Saída:**

    {Maria Silva 28 maria@example.com}
    Nome: Maria Silva
    Idade: 28

---

## 🎯 Formas de Criar Structs

### 1. Forma Literal Completa (Recomendada)

    pessoa := Pessoa{
        Nome:  "João",
        Idade: 30,
        Email: "joao@example.com",
    }

**Vantagem:** Clara, legível e explícita sobre quais campos estão sendo preenchidos.

### 2. Forma Literal Posicional

    pessoa := Pessoa{"João", 30, "joao@example.com"}

**Atenção:** Não recomendada! Se a ordem dos campos mudar na struct, seu código quebra silenciosamente.

### 3. Criação com Valores Zero

    var pessoa Pessoa
    pessoa.Nome = "Ana"
    pessoa.Idade = 25

**Campos não inicializados recebem valor zero:**
- `string` → `""`
- `int` → `0`
- `bool` → `false`

### 4. Usando Ponteiros para Structs

    pessoa := &Pessoa{
        Nome:  "Carlos",
        Idade: 35,
    }
    
    fmt.Println(pessoa.Nome) // Go automaticamente desreferencia

---

## 🔑 Acessando e Modificando Campos

Use a **notação de ponto** para acessar campos:

    package main

    import "fmt"

    type Livro struct {
        Titulo string
        Autor  string
        Paginas int
        Preco  float64
    }

    func main() {
        livro := Livro{
            Titulo: "Go em Ação",
            Autor: "William Kennedy",
            Paginas: 350,
            Preco: 89.90,
        }
        
        // Lendo valores
        fmt.Println("Título:", livro.Titulo)
        
        // Modificando valores
        livro.Preco = 79.90
        fmt.Println("Novo preço:", livro.Preco)
        
        // Incrementando
        livro.Paginas += 50
        fmt.Println("Novas páginas:", livro.Paginas)
    }

---

## 🧩 Structs Aninhadas (Nested Structs)

Você pode ter structs dentro de structs para representar relacionamentos complexos:

    package main

    import "fmt"

    type Endereco struct {
        Rua    string
        Numero int
        Cidade string
        Estado string
        CEP    string
    }

    type Pessoa struct {
        Nome     string
        Idade    int
        Endereco Endereco // Struct dentro de struct
    }

    func main() {
        pessoa := Pessoa{
            Nome:  "Roberto",
            Idade: 42,
            Endereco: Endereco{
                Rua:    "Av. Paulista",
                Numero: 1000,
                Cidade: "São Paulo",
                Estado: "SP",
                CEP:    "01310-100",
            },
        }
        
        fmt.Println("Nome:", pessoa.Nome)
        fmt.Println("Cidade:", pessoa.Endereco.Cidade)
        fmt.Println("CEP:", pessoa.Endereco.CEP)
    }

---

## 🎭 Métodos em Structs

Em Go, métodos são funções associadas a um tipo específico. Diferente de classes em outras linguagens, os métodos são definidos **separadamente** da struct:

    package main

    import "fmt"

    type Retangulo struct {
        Largura float64
        Altura  float64
    }

    // Método com receptor de valor
    func (r Retangulo) Area() float64 {
        return r.Largura * r.Altura
    }

    // Método com receptor de ponteiro
    func (r *Retangulo) Escalar(fator float64) {
        r.Largura *= fator
        r.Altura *= fator
    }

    func main() {
        ret := Retangulo{Largura: 10, Altura: 5}
        
        fmt.Println("Área:", ret.Area())
        
        ret.Escalar(2)
        fmt.Println("Nova área:", ret.Area())
    }

### 🔍 Receptor de Valor vs Receptor de Ponteiro

**Receptor de Valor** `(r Retangulo)`:
- Recebe uma **cópia** da struct
- Não modifica a struct original
- Usado para operações de leitura

**Receptor de Ponteiro** `(r *Retangulo)`:
- Recebe uma **referência** à struct
- Pode modificar a struct original
- Usado para operações de escrita
- Mais eficiente para structs grandes (evita cópias)

---

## 📦 Structs Vazias

Structs podem não ter campos (útil para marcadores ou sincronização):

    type Vazio struct{}
    
    var v Vazio
    // Ocupa 0 bytes de memória!

---

## 🚀 Passando Structs para Funções

### Por Valor (Cópia)

    func imprimirPessoa(p Pessoa) {
        fmt.Println(p.Nome)
    }
    
    // A função recebe uma cópia, modificações não afetam o original

### Por Referência (Ponteiro)

    func atualizarIdade(p *Pessoa, novaIdade int) {
        p.Idade = novaIdade
    }
    
    func main() {
        pessoa := Pessoa{Nome: "Ana", Idade: 25}
        atualizarIdade(&pessoa, 26)
        fmt.Println(pessoa.Idade) // 26
    }

**Regra de Ouro:** Use ponteiros quando precisar modificar a struct ou quando ela for muito grande (para evitar cópias custosas).

---

## 🎨 Comparação de Structs

Structs podem ser comparadas com `==` se todos os seus campos forem comparáveis:

    package main

    import "fmt"

    type Ponto struct {
        X int
        Y int
    }

    func main() {
        p1 := Ponto{X: 1, Y: 2}
        p2 := Ponto{X: 1, Y: 2}
        p3 := Ponto{X: 2, Y: 3}
        
        fmt.Println(p1 == p2) // true
        fmt.Println(p1 == p3) // false
    }

**Atenção:** Structs contendo slices, maps ou funções **não podem** ser comparadas com `==`.

---

## 🏭 Exemplo Completo: Sistema de Produtos

    package main

    import "fmt"

    type Produto struct {
        ID        int
        Nome      string
        Preco     float64
        Estoque   int
        Categoria string
    }

    // Método para verificar disponibilidade
    func (p Produto) EstaDisponivel() bool {
        return p.Estoque > 0
    }

    // Método para calcular valor total em estoque
    func (p Produto) ValorTotalEstoque() float64 {
        return p.Preco * float64(p.Estoque)
    }

    // Método para vender (modifica estoque)
    func (p *Produto) Vender(quantidade int) bool {
        if p.Estoque >= quantidade {
            p.Estoque -= quantidade
            return true
        }
        return false
    }

    // Função para exibir informações do produto
    func exibirProduto(p Produto) {
        fmt.Printf("ID: %d\n", p.ID)
        fmt.Printf("Nome: %s\n", p.Nome)
        fmt.Printf("Preço: R$ %.2f\n", p.Preco)
        fmt.Printf("Estoque: %d unidades\n", p.Estoque)
        fmt.Printf("Valor Total: R$ %.2f\n", p.ValorTotalEstoque())
        fmt.Printf("Disponível: %v\n", p.EstaDisponivel())
        fmt.Println("---")
    }

    func main() {
        produto1 := Produto{
            ID:        1,
            Nome:      "Notebook",
            Preco:     3500.00,
            Estoque:   10,
            Categoria: "Eletrônicos",
        }
        
        produto2 := Produto{
            ID:        2,
            Nome:      "Mouse Gamer",
            Preco:     150.00,
            Estoque:   0,
            Categoria: "Periféricos",
        }
        
        exibirProduto(produto1)
        exibirProduto(produto2)
        
        // Tentando vender
        if produto1.Vender(3) {
            fmt.Println("Venda realizada com sucesso!")
            exibirProduto(produto1)
        }
        
        if !produto2.Vender(1) {
            fmt.Println("Produto indisponível!")
        }
    }

---

## 🎯 Quando Usar Structs?

✅ **Use structs quando:**
- Precisar agrupar dados relacionados
- Quiser criar modelos de dados (usuário, produto, pedido)
- Necessitar organizar informações complexas
- Desenvolver APIs e manipular JSON
- Criar sistemas orientados a objetos

❌ **Evite structs quando:**
- Tiver apenas um ou dois campos simples (use tipos básicos)
- Os dados não têm relação lógica entre si
- Precisar de herança clássica (Go usa composição)

---

## 💡 Conceitos-Chave para Memorizar

1. **Structs agrupam campos relacionados** sob um nome único
2. **Métodos são definidos separadamente** da definição da struct
3. **Notação de ponto** é usada para acessar campos
4. **Receptores de ponteiro** modificam a struct original
5. **Receptores de valor** trabalham com cópias
6. **Structs são comparáveis** se todos os campos forem comparáveis
7. **Composição sobre herança** é a filosofia do Go

---

## 📚 Boas Práticas

1. **Nomeie campos com letra maiúscula** para exportá-los (torná-los públicos)
2. **Use ponteiros para structs grandes** para evitar cópias desnecessárias
3. **Prefira a forma literal completa** ao criar structs
4. **Use métodos de ponteiro** quando precisar modificar o estado
5. **Organize structs relacionadas** no mesmo arquivo ou pacote
6. **Documente structs complexas** com comentários

---

## 🎓 Próximos Passos

Na próxima parte desta aula, vamos explorar:
- **Struct Tags** para metadados e serialização JSON
- **Embedding** (composição de structs)
- Técnicas avançadas de organização de código

---

**Fim da Aula 6 - Parte 1: Structs Fundamentais**