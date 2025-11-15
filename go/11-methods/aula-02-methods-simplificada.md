# Aula 2 - Simplificada: Entendendo Methods vs Functions

Olá! Vamos simplificar os conceitos de methods e functions usando analogias do dia a dia. Isso vai ajudar você a entender quando usar cada um e por quê.

---

## Methods vs Functions: A Analogia do Carro

Imagine que você tem um **carro**. O carro tem várias **ações** que ele pode fazer:

### Methods: Ações do Carro

Methods são como as **ações que o carro pode fazer por si mesmo**:

```go
type Carro struct {
    modelo string
    velocidade int
    tanque float64
}

// O carro "sabe" acelerar
func (c *Carro) Acelerar() {
    c.velocidade += 10
}

// O carro "sabe" mostrar sua velocidade
func (c *Carro) VelocidadeAtual() int {
    return c.velocidade
}
```

**Pense assim**: O carro **tem** a capacidade de acelerar. É uma ação que pertence ao carro.

### Functions: Operações Externas

Functions são como **operações que alguém de fora faz com o carro**:

```go
// Alguém de fora calcula a distância que o carro percorreu
func CalcularDistancia(velocidade int, tempo float64) float64 {
    return float64(velocidade) * tempo
}

// Alguém de fora verifica se o carro precisa de manutenção
func PrecisaManutencao(kmRodados int) bool {
    return kmRodados > 10000
}
```

**Pense assim**: Essas são operações que **não pertencem** ao carro, mas podem ser feitas **sobre** o carro.

### Quando Usar Cada Um?

- **Method**: Quando a ação é algo que o objeto "faz" ou "sabe fazer"
  - `carro.Acelerar()` - O carro acelera
  - `carro.Abastecer()` - O carro recebe combustível
  
- **Function**: Quando a operação é externa ou genérica
  - `CalcularDistancia()` - Cálculo matemático geral
  - `CompararCarros()` - Comparação entre dois carros

---

## Value Receivers: A Fotocópia

Imagine que você tem um **documento importante** e precisa fazer uma **fotocópia** para alguém trabalhar nele.

### A Analogia

```go
type Documento struct {
    texto string
    paginas int
}

// Value receiver: trabalha com uma CÓPIA
func (d Documento) Ler() string {
    return d.texto  // Lê a cópia
}

// Value receiver: não modifica o original
func (d Documento) ContarPalavras() int {
    // Conta palavras na cópia, original não muda
    return len(strings.Fields(d.texto))
}
```

**Pense assim**: 
- Você faz uma **fotocópia** do documento
- A pessoa trabalha na **cópia**
- O **original** permanece intacto
- Qualquer alteração na cópia **não afeta** o original

### Quando Usar Value Receiver?

Use quando você quer garantir que o **original não será modificado**:

- Ler informações
- Calcular valores
- Gerar relatórios
- Operações que não alteram o estado

**Exemplo do dia a dia**: Você empresta um livro para alguém ler. Eles podem ler, mas não podem escrever nele (value receiver). Se quiserem fazer anotações, precisam de uma cópia.

---

## Pointer Receivers: A Chave da Casa

Agora imagine que você tem uma **casa** e precisa dar a **chave** para alguém fazer reformas.

### A Analogia

```go
type Casa struct {
    cor string
    quartos int
    valor float64
}

// Pointer receiver: trabalha com a CASA REAL
func (c *Casa) Pintar(novaCor string) {
    c.cor = novaCor  // Modifica a casa REAL
}

// Pointer receiver: adiciona um quarto
func (c *Casa) ConstruirQuarto() {
    c.quartos++
    c.valor += 50000  // Aumenta o valor REAL
}
```

**Pense assim**:
- Você dá a **chave** da casa (pointer)
- A pessoa trabalha na **casa real**
- Qualquer mudança **afeta a casa original**
- Não há cópia, é a casa verdadeira

### Quando Usar Pointer Receiver?

Use quando você precisa **modificar o original**:

- Alterar valores
- Adicionar ou remover itens
- Atualizar estado
- Qualquer operação que muda o objeto

**Exemplo do dia a dia**: Você dá a chave do seu carro para um mecânico. Ele vai modificar o carro REAL, não uma cópia. Quando você pegar o carro de volta, ele estará modificado.

---

## Comparação Prática: Agenda de Telefone

Vamos ver um exemplo completo usando uma **agenda de telefone**:

### Value Receiver: Consultar Contato

```go
type Agenda struct {
    contatos map[string]string
}

// Value receiver: apenas consulta, não modifica
func (a Agenda) BuscarTelefone(nome string) (string, bool) {
    telefone, existe := a.contatos[nome]
    return telefone, existe
}
```

**Analogia**: Você **olha** na agenda para encontrar um número. Olhar não modifica a agenda.

### Pointer Receiver: Adicionar Contato

```go
// Pointer receiver: modifica a agenda
func (a *Agenda) AdicionarContato(nome, telefone string) {
    if a.contatos == nil {
        a.contatos = make(map[string]string)
    }
    a.contatos[nome] = telefone  // Modifica a agenda REAL
}
```

**Analogia**: Você **escreve** um novo contato na agenda. Isso modifica a agenda real.

---

## A Mágica do Go: Conversão Automática

Go é muito inteligente! Ele automaticamente converte entre valores e pointers quando você chama métodos.

### Analogia: O Assistente Inteligente

Imagine que você tem um **assistente** que sempre sabe o que você quer:

```go
type Pessoa struct {
    Nome string
}

func (p *Pessoa) MudarNome(novoNome string) {
    p.Nome = novoNome
}

func main() {
    // Você tem um VALOR
    pessoa := Pessoa{Nome: "João"}
    
    // Você pede para mudar o nome
    pessoa.MudarNome("Pedro")  // Go automaticamente faz: (&pessoa).MudarNome()
    
    // O assistente (Go) entendeu que você quer trabalhar com a pessoa REAL
    // Então ele automaticamente pegou o endereço (pointer) para você
}
```

**Pense assim**: 
- Você pede algo ao assistente
- O assistente entende sua intenção
- Ele faz a conversão necessária automaticamente
- Você não precisa se preocupar com os detalhes técnicos

**Mas atenção**: Embora Go faça isso automaticamente, é melhor ser **consistente**. Se você tem métodos que modificam, trabalhe sempre com pointers para deixar claro sua intenção.

---

## Methods em Tipos Não-Struct: O Superpoder

Em Go, você pode adicionar métodos a **qualquer tipo**, não apenas structs!

### Analogia: Adicionar Funcionalidades a Coisas Existentes

Imagine que você tem um **número** (como 1000 centavos) e quer transformá-lo em "R$ 10,00":

```go
type Dinheiro int

// Agora o número "sabe" se apresentar como dinheiro
func (d Dinheiro) Reais() string {
    return fmt.Sprintf("R$ %.2f", float64(d)/100)
}
```

**Pense assim**: 
- Você pega algo simples (um número)
- Você adiciona uma "habilidade" a ele (saber se apresentar como dinheiro)
- Agora esse número tem um comportamento especial

**Exemplo do dia a dia**: É como adicionar uma etiqueta a um objeto. O objeto continua sendo o mesmo, mas agora ele "sabe" como se apresentar de uma forma especial.

---

## Resumo com Analogias

### Methods vs Functions

- **Method**: Ação que o objeto "sabe fazer" (carro acelerar)
- **Function**: Operação externa sobre o objeto (calcular distância)

### Value Receiver

- **Fotocópia**: Trabalha com uma cópia, original não muda
- **Ler livro**: Pode ler, mas não pode escrever no original
- **Consultar agenda**: Olhar não modifica

### Pointer Receiver

- **Chave da casa**: Trabalha com o objeto real
- **Mecânico no carro**: Modifica o carro real
- **Escrever na agenda**: Modifica a agenda real

### Conversão Automática

- **Assistente inteligente**: Go entende sua intenção e converte automaticamente

---

## Dica Final

Quando estiver em dúvida sobre qual usar, pergunte a si mesmo:

1. **Esta ação pertence ao objeto? → Method
2. **Preciso modificar o original? → Pointer Receiver
3. **Apenas preciso ler/calcular? → Value Receiver
4. **É uma operação genérica? → Function

Lembre-se: Go é uma linguagem prática. Com o tempo, você vai desenvolver uma "intuição" sobre quando usar cada um. O importante é começar e praticar!

---

Na próxima aula, vamos fazer exercícios práticos para fixar esses conceitos!

