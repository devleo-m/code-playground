# Módulo 6: Structs em Go

## Aula 2 - Simplificada: Entendendo Structs

Agora vamos entender structs de forma mais simples, usando analogias do nosso dia a dia!

---

## 1. Struct: A Ficha de Cadastro

### Analogia: Ficha de Cadastro

Pense em uma struct como uma **ficha de cadastro** ou **formulário**:

**Ficha de cadastro física:**

```
┌─────────────────────────────┐
│   FICHA DE CADASTRO         │
├─────────────────────────────┤
│ Nome: ___________________   │
│ Idade: __________________   │
│ Email: __________________   │
│ Telefone: _______________   │
└─────────────────────────────┘
```

**Struct em Go:**

```go
type Pessoa struct {
    Nome     string
    Idade    int
    Email    string
    Telefone string
}
```

**Por que funciona:**

- A ficha agrupa todas as informações de uma pessoa
- A struct agrupa todos os dados relacionados
- É fácil encontrar tudo em um só lugar!

---

## 2. Campos: As Caixas do Formulário

### Analogia: Formulário com Caixas

Cada campo da struct é como uma **caixa** no formulário:

**Formulário:**

```
Nome: [___________]  ← Campo Nome
Idade: [____]        ← Campo Idade
Email: [___________] ← Campo Email
```

**Struct:**

```go
type Pessoa struct {
    Nome  string  // ← Campo Nome
    Idade int     // ← Campo Idade
    Email string  // ← Campo Email
}
```

**Acessando campos:**

```go
pessoa := Pessoa{Nome: "João", Idade: 30}

// É como preencher a caixa do formulário
pessoa.Nome = "João Silva"  // Preenche a caixa "Nome"
pessoa.Idade = 31           // Preenche a caixa "Idade"
```

---

## 3. Criar Struct: Preencher o Formulário

### Analogia: Preencher um Formulário

Criar uma struct é como **preencher um formulário**:

**Formulário físico:**

1. Você pega um formulário em branco
2. Preenche cada campo
3. Entrega o formulário preenchido

**Struct em Go:**

```go
// 1. Pegar formulário em branco (valores zero)
var pessoa Pessoa

// 2. Preencher cada campo
pessoa.Nome = "João"
pessoa.Idade = 30
pessoa.Email = "joao@email.com"

// OU preencher tudo de uma vez
pessoa := Pessoa{
    Nome:  "João",
    Idade: 30,
    Email: "joao@email.com",
}
```

**Analogia prática:**
É como ter um **formulário digital** que você preenche no computador!

---

## 4. Múltiplas Instâncias: Múltiplas Fichas

### Analogia: Arquivo com Múltiplas Fichas

Você pode ter **muitas fichas** (muitas structs):

**Arquivo físico:**

```
┌─────────┐  ┌─────────┐  ┌─────────┐
│ Ficha 1 │  │ Ficha 2 │  │ Ficha 3 │
│ João    │  │ Maria   │  │ Pedro   │
└─────────┘  └─────────┘  └─────────┘
```

**Em Go:**

```go
pessoa1 := Pessoa{Nome: "João", Idade: 30}
pessoa2 := Pessoa{Nome: "Maria", Idade: 25}
pessoa3 := Pessoa{Nome: "Pedro", Idade: 35}

// Ou em um slice (como uma pasta com fichas)
pessoas := []Pessoa{pessoa1, pessoa2, pessoa3}
```

**Analogia:**
É como ter uma **pasta com várias fichas de cadastro**!

---

## 5. Métodos: Ações que a Ficha Pode Fazer

### Analogia: Botões em um Formulário Digital

Métodos são como **botões** em um formulário digital que fazem ações:

**Formulário digital:**

```
┌─────────────────────────────┐
│ Nome: [João Silva]          │
│ Idade: [30]                 │
│                             │
│ [Calcular Idade em Dias] ← Botão (método)
│ [Fazer Aniversário]      ← Botão (método)
└─────────────────────────────┘
```

**Em Go:**

```go
type Pessoa struct {
    Nome  string
    Idade int
}

// Método: botão "Fazer Aniversário"
func (p *Pessoa) FazerAniversario() {
    p.Idade++  // Aumenta a idade
}

// Método: botão "Apresentar"
func (p Pessoa) Apresentar() string {
    return fmt.Sprintf("Olá, sou %s", p.Nome)
}

// Usar os "botões"
pessoa := Pessoa{Nome: "João", Idade: 30}
pessoa.FazerAniversario()  // Clica no botão
fmt.Println(pessoa.Apresentar())  // Clica no outro botão
```

**Analogia:**
Métodos são **ações** que a struct pode fazer, como botões em uma interface!

---

## 6. Passar para Funções: Enviar o Formulário

### Analogia: Enviar Formulário por Email

Passar uma struct para uma função é como **enviar um formulário**:

**Cenário 1: Enviar Cópia (por valor)**

```
Você faz uma CÓPIA do formulário
Envia a cópia
O original não muda
```

```go
func Processar(p Pessoa) {
    p.Idade++  // Modifica apenas a CÓPIA
}

pessoa := Pessoa{Idade: 30}
Processar(pessoa)  // Envia cópia
// pessoa.Idade ainda é 30 (original não mudou)
```

**Cenário 2: Enviar Original (por ponteiro)**

```
Você envia o ORIGINAL
A pessoa recebe e modifica
O original muda
```

```go
func Processar(p *Pessoa) {
    p.Idade++  // Modifica o ORIGINAL
}

pessoa := Pessoa{Idade: 30}
Processar(&pessoa)  // Envia original
// pessoa.Idade agora é 31 (original mudou)
```

**Analogia:**

- **Por valor**: Como enviar uma **fotocópia** - original não muda
- **Por ponteiro**: Como enviar o **original** - pode ser modificado

---

## 7. Structs Aninhadas: Formulário dentro de Formulário

### Analogia: Formulário com Seção de Endereço

Structs podem conter outras structs, como um **formulário com seções**:

**Formulário físico:**

```
┌─────────────────────────────┐
│ DADOS PESSOAIS              │
│ Nome: [João]                │
│ Idade: [30]                 │
├─────────────────────────────┤
│ ENDEREÇO                    │
│ Rua: [Rua A]                │
│ Cidade: [São Paulo]         │
└─────────────────────────────┘
```

**Em Go:**

```go
type Endereco struct {
    Rua    string
    Cidade string
}

type Pessoa struct {
    Nome     string
    Idade    int
    Endereco Endereco  // Formulário dentro de formulário!
}

pessoa := Pessoa{
    Nome:  "João",
    Idade: 30,
    Endereco: Endereco{
        Rua:    "Rua A",
        Cidade: "São Paulo",
    },
}

// Acessar: como abrir a seção do formulário
fmt.Println(pessoa.Endereco.Rua)  // "Rua A"
```

**Analogia:**
É como ter um **formulário principal** com **subformulários** dentro!

---

## 8. Comparação: Comparar Duas Fichas

### Analogia: Comparar Duas Fichas de Cadastro

Comparar structs é como **comparar duas fichas** para ver se são iguais:

**Ficha 1:**

```
Nome: João
Idade: 30
```

**Ficha 2:**

```
Nome: João
Idade: 30
```

**São iguais?** Sim! Todos os campos são iguais.

**Em Go:**

```go
ficha1 := Pessoa{Nome: "João", Idade: 30}
ficha2 := Pessoa{Nome: "João", Idade: 30}

fmt.Println(ficha1 == ficha2)  // true - são iguais!
```

**Analogia:**
Go compara **campo por campo**, como você compararia duas fichas lado a lado!

---

## 9. Valores Zero: Formulário em Branco

### Analogia: Formulário Não Preenchido

Quando você declara uma struct sem valores, é como pegar um **formulário em branco**:

**Formulário em branco:**

```
Nome: [___________]  ← Vazio
Idade: [____]        ← Zero
Email: [___________] ← Vazio
```

**Em Go:**

```go
var pessoa Pessoa
// pessoa.Nome = "" (vazio)
// pessoa.Idade = 0 (zero)
// pessoa.Email = "" (vazio)
```

**Analogia:**
É como pegar um **formulário novo** - todas as caixas estão vazias (valores zero)!

---

## 10. Slice de Structs: Pasta com Múltiplas Fichas

### Analogia: Pasta com Várias Fichas

Um slice de structs é como uma **pasta** com várias fichas:

**Pasta física:**

```
┌─────────────────┐
│   PASTA         │
├─────────────────┤
│ 📄 Ficha 1      │
│ 📄 Ficha 2      │
│ 📄 Ficha 3      │
└─────────────────┘
```

**Em Go:**

```go
// Pasta com fichas
pasta := []Pessoa{
    {Nome: "João", Idade: 30},    // Ficha 1
    {Nome: "Maria", Idade: 25},  // Ficha 2
    {Nome: "Pedro", Idade: 35},  // Ficha 3
}

// Ver cada ficha
for _, ficha := range pasta {
    fmt.Println(ficha.Nome)
}
```

**Analogia:**
É como ter uma **pasta de arquivo** com várias fichas de cadastro dentro!

---

## 11. Map com Structs: Gaveta Organizada

### Analogia: Gaveta com Etiquetas

Um map com structs é como uma **gaveta organizada** com etiquetas:

**Gaveta física:**

```
┌─────────────────────┐
│ [joao] → Ficha João  │
│ [maria] → Ficha Maria│
│ [pedro] → Ficha Pedro│
└─────────────────────┘
```

**Em Go:**

```go
// Gaveta organizada
gaveta := make(map[string]Pessoa)

// Colocar fichas na gaveta
gaveta["joao"] = Pessoa{Nome: "João", Idade: 30}
gaveta["maria"] = Pessoa{Nome: "Maria", Idade: 25}

// Pegar uma ficha
fichaJoao := gaveta["joao"]
fmt.Println(fichaJoao.Nome)  // "João"
```

**Analogia:**
É como uma **gaveta de arquivo** onde cada ficha tem uma **etiqueta** (chave) para encontrá-la rápido!

---

## 12. Métodos que Modificam: Botão que Altera o Formulário

### Analogia: Botão "Salvar Alterações"

Métodos com ponteiro são como **botões que modificam** o formulário:

**Formulário digital:**

```
┌─────────────────────────────┐
│ Nome: [João]                │
│ Idade: [30]                 │
│                             │
│ [Fazer Aniversário] ← Este botão MODIFICA
│   (aumenta a idade)         │
└─────────────────────────────┘
```

**Em Go:**

```go
type Pessoa struct {
    Nome  string
    Idade int
}

// Método que MODIFICA (usa ponteiro)
func (p *Pessoa) FazerAniversario() {
    p.Idade++  // Modifica o original!
}

pessoa := Pessoa{Nome: "João", Idade: 30}
pessoa.FazerAniversario()  // Clica no botão
// Agora pessoa.Idade é 31 (mudou!)
```

**Analogia:**
É como clicar em um **botão "Salvar"** que realmente **altera** os dados do formulário!

---

## 13. Métodos que Apenas Leem: Botão de Visualização

### Analogia: Botão "Ver Resumo"

Métodos sem ponteiro são como **botões que apenas mostram** informações:

**Formulário digital:**

```
┌─────────────────────────────┐
│ Nome: [João]                 │
│ Idade: [30]                  │
│                             │
│ [Ver Apresentação] ← Este botão APENAS LÊ
│   (não modifica nada)        │
└─────────────────────────────┘
```

**Em Go:**

```go
// Método que APENAS LÊ (não precisa de ponteiro)
func (p Pessoa) Apresentar() string {
    return fmt.Sprintf("Olá, sou %s", p.Nome)
}

pessoa := Pessoa{Nome: "João", Idade: 30}
mensagem := pessoa.Apresentar()  // Apenas lê, não modifica
// pessoa não mudou, apenas leu os dados
```

**Analogia:**
É como clicar em um **botão "Visualizar"** que **mostra** informações mas **não altera** nada!

---

## 14. Resumo Visual: Struct como Formulário

Pense em struct como um **formulário digital completo**:

```
┌─────────────────────────────────┐
│     FORMULÁRIO (STRUCT)          │
├─────────────────────────────────┤
│                                 │
│  Campos (Caixas):               │
│  ┌─────────────────────────┐    │
│  │ Nome: [___________]     │    │
│  │ Idade: [____]           │    │
│  │ Email: [___________]    │    │
│  └─────────────────────────┘    │
│                                 │
│  Métodos (Botões):              │
│  [Fazer Aniversário]            │
│  [Apresentar]                   │
│  [Calcular Idade em Dias]       │
│                                 │
└─────────────────────────────────┘
```

**Em código:**

```go
type Pessoa struct {
    // Campos (caixas do formulário)
    Nome  string
    Idade int
    Email string
}

// Métodos (botões do formulário)
func (p *Pessoa) FazerAniversario() { ... }
func (p Pessoa) Apresentar() string { ... }
```

---

## 15. Analogia Final: A Receita de Bolo

Structs são como uma **receita de bolo**:

**Receita física:**

```
┌─────────────────────────────┐
│   RECEITA DE BOLO           │
├─────────────────────────────┤
│ Ingredientes:               │
│ - Farinha: 2 xícaras        │
│ - Açúcar: 1 xícara          │
│ - Ovos: 3 unidades          │
│                             │
│ Método (Passos):            │
│ 1. Misturar ingredientes    │
│ 2. Assar por 30 min         │
└─────────────────────────────┘
```

**Struct em Go:**

```go
type Bolo struct {
    // Ingredientes (campos)
    Farinha float64
    Acucar  float64
    Ovos    int
}

// Método (passo da receita)
func (b *Bolo) Assar() {
    // Lógica para assar
}
```

**Por que funciona:**

- A receita **agrupa** ingredientes relacionados
- A struct **agrupa** dados relacionados
- A receita tem **passos** (métodos)
- A struct tem **métodos** (ações)

---

## 16. Regra de Ouro Simples

**Struct = Formulário Digital**

- **Campos** = Caixas do formulário
- **Valores** = O que você escreve nas caixas
- **Métodos** = Botões que fazem ações
- **Instâncias** = Múltiplas cópias do formulário

**Quando usar:**

- Precisa agrupar dados relacionados? → Use struct!
- Precisa representar uma "coisa" (pessoa, produto, etc.)? → Use struct!
- Precisa organizar informações? → Use struct!

---

Agora que você entendeu os conceitos de forma simplificada, vamos praticar com exercícios na próxima parte!
