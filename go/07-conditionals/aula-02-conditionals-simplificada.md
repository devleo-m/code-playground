# Módulo 7: Conditionals em Go

## Aula 2 - Simplificada: Entendendo Conditionals

Agora vamos entender conditionals de forma mais simples, usando analogias do nosso dia a dia!

---

## 1. Conditionals: As Decisões do Programa

### Analogia: Semáforo de Trânsito

Pense em conditionals como um **semáforo de trânsito**:

**Semáforo físico:**

```
🟢 VERDE → Pode passar
🟡 AMARELO → Atenção, prepare-se para parar
🔴 VERMELHO → Pare!
```

**Programa em Go:**

```go
cor := "verde"

if cor == "verde" {
    fmt.Println("Pode passar")
} else if cor == "amarelo" {
    fmt.Println("Atenção!")
} else {
    fmt.Println("Pare!")
}
```

**Por que funciona:**

- O semáforo **decide** se você pode passar ou não
- O conditional **decide** qual código executar
- Ambos dependem de uma **condição** (cor do semáforo)

---

## 2. `if`: A Pergunta Simples

### Analogia: Pergunta Simples

Pense em `if` como fazer uma **pergunta simples**:

**Pergunta do dia a dia:**

```
"Está chovendo?"
Se SIM → Pegar guarda-chuva
Se NÃO → Não precisa
```

**Em Go:**

```go
chovendo := true

if chovendo {
    fmt.Println("Pegar guarda-chuva")
}
```

**Analogia prática:**

É como perguntar: **"Se isso for verdade, faça aquilo"**

---

## 3. `if-else`: A Escolha entre Duas Opções

### Analogia: Porta com Duas Saídas

Pense em `if-else` como uma **porta com duas saídas**:

**Cenário físico:**

```
┌─────────────┐
│   PORTA     │
│             │
│  [SIM] → ───┼──→ Saída A
│             │
│  [NÃO] → ───┼──→ Saída B
└─────────────┘
```

**Em Go:**

```go
idade := 20

if idade >= 18 {
    fmt.Println("Maior de idade")  // Saída A
} else {
    fmt.Println("Menor de idade")   // Saída B
}
```

**Analogia:**

- Se a condição for **verdadeira** → vai pela **Saída A**
- Se a condição for **falsa** → vai pela **Saída B**
- **Sempre** escolhe uma das duas!

---

## 4. `if-else if`: Múltiplas Escolhas

### Analogia: Menu de Restaurante

Pense em `if-else if` como um **menu de restaurante**:

**Menu físico:**

```
┌─────────────────────┐
│   MENU              │
├─────────────────────┤
│ 1. Prato A → R$ 20  │
│ 2. Prato B → R$ 15  │
│ 3. Prato C → R$ 10  │
│ 4. Outro → R$ 5     │
└─────────────────────┘
```

**Em Go:**

```go
nota := 85

if nota >= 90 {
    fmt.Println("Nota A")      // Prato A
} else if nota >= 80 {
    fmt.Println("Nota B")      // Prato B
} else if nota >= 70 {
    fmt.Println("Nota C")      // Prato C
} else {
    fmt.Println("Reprovado")   // Outro
}
```

**Analogia:**

- Você escolhe o **primeiro prato** que pode pagar
- O programa escolhe o **primeiro bloco** cuja condição for verdadeira
- Se nenhum servir, escolhe o **padrão** (else)

---

## 5. `if` com Inicialização: Preparar Antes de Perguntar

### Analogia: Verificar o Saldo Antes de Comprar

Pense em `if` com inicialização como **verificar o saldo antes de comprar**:

**Cenário físico:**

```
1. Olhar saldo na conta
2. Se saldo >= preço → Comprar
3. Se saldo < preço → Não comprar
```

**Em Go:**

```go
// Verificar saldo e decidir na mesma "ação"
if saldo := verificarSaldo(); saldo >= 100 {
    fmt.Println("Pode comprar")
} else {
    fmt.Println("Saldo insuficiente")
}
```

**Analogia:**

- Você **prepara** (verifica saldo) e **decide** (pode comprar?) na mesma ação
- A variável `saldo` existe **apenas** durante essa decisão
- Depois, não precisa mais dela

---

## 6. Operadores Lógicos: Combinando Condições

### Analogia: Requisitos para Entrar em um Clube

Pense em operadores lógicos como **requisitos para entrar em um clube**:

**Clube físico:**

```
Para entrar você precisa:
✅ Ter 18 anos OU mais
✅ E ter carteira de sócio
✅ E não estar banido
```

**Em Go:**

```go
idade := 25
temCarteira := true
banido := false

// && (E) - TODAS as condições devem ser verdadeiras
if idade >= 18 && temCarteira && !banido {
    fmt.Println("Pode entrar")
}

// || (OU) - PELO MENOS UMA condição deve ser verdadeira
if idade < 18 || !temCarteira {
    fmt.Println("Não pode entrar")
}
```

**Analogia:**

- **`&&` (E)**: Como uma **lista de requisitos** - precisa de **TODOS**
- **`||` (OU)**: Como **alternativas** - precisa de **PELO MENOS UM**
- **`!` (NÃO)**: Como um **bloqueio** - inverte o valor

---

## 7. `switch`: O Seletor de Canais

### Analogia: Controle Remoto de TV

Pense em `switch` como um **controle remoto de TV**:

**Controle físico:**

```
┌─────────────────┐
│  CONTROLE TV    │
├─────────────────┤
│ [1] → Canal 1   │
│ [2] → Canal 2   │
│ [3] → Canal 3   │
│ [4] → Canal 4   │
│ [*] → Outro     │
└─────────────────┘
```

**Em Go:**

```go
canal := 2

switch canal {
case 1:
    fmt.Println("Canal 1")
case 2:
    fmt.Println("Canal 2")
case 3:
    fmt.Println("Canal 3")
case 4:
    fmt.Println("Canal 4")
default:
    fmt.Println("Canal desconhecido")
}
```

**Analogia:**

- Você **pressiona um botão** (escolhe um valor)
- A TV **muda para aquele canal** (executa aquele código)
- Se não houver botão, vai para o **padrão** (default)

---

## 8. `switch` sem Expressão: O Questionário

### Analogia: Questionário de Múltipla Escolha

Pense em `switch` sem expressão como um **questionário**:

**Questionário físico:**

```
┌─────────────────────────┐
│   QUESTIONÁRIO         │
├─────────────────────────┤
│ Qual sua idade?        │
│                        │
│ ( ) < 18 → Menor      │
│ ( ) 18-65 → Adulto    │
│ ( ) > 65 → Idoso      │
└─────────────────────────┘
```

**Em Go:**

```go
idade := 25

switch {
case idade < 18:
    fmt.Println("Menor")
case idade >= 18 && idade < 65:
    fmt.Println("Adulto")
case idade >= 65:
    fmt.Println("Idoso")
}
```

**Analogia:**

- Você **marca uma opção** (condição verdadeira)
- O programa **executa o código** daquela opção
- É como responder um **questionário** onde você escolhe a resposta que se aplica

---

## 9. `switch` com Múltiplos Valores: O Grupo de Botões

### Analogia: Painel de Controle com Grupos

Pense em `switch` com múltiplos valores como um **painel com grupos de botões**:

**Painel físico:**

```
┌─────────────────────┐
│   PAINEL            │
├─────────────────────┤
│ [Verão]             │
│  Dez, Jan, Fev      │
│                     │
│ [Outono]            │
│  Mar, Abr, Mai      │
│                     │
│ [Inverno]           │
│  Jun, Jul, Ago      │
└─────────────────────┘
```

**Em Go:**

```go
mes := "janeiro"

switch mes {
case "dezembro", "janeiro", "fevereiro":
    fmt.Println("Verão")
case "março", "abril", "maio":
    fmt.Println("Outono")
case "junho", "julho", "agosto":
    fmt.Println("Inverno")
}
```

**Analogia:**

- Vários **botões** fazem a **mesma coisa**
- Vários **valores** executam o **mesmo código**
- É como ter um **grupo de atalhos** para a mesma ação

---

## 10. `fallthrough`: A Corrente de Ações

### Analogia: Corrente de Dominós

Pense em `fallthrough` como uma **corrente de dominós**:

**Corrente física:**

```
[Domino 1] → cai
    ↓
[Domino 2] → cai
    ↓
[Domino 3] → cai
    ↓
[Para aqui]
```

**Em Go:**

```go
numero := 2

switch numero {
case 1:
    fmt.Println("Um")
    fallthrough  // Continua caindo
case 2:
    fmt.Println("Dois")
    fallthrough  // Continua caindo
case 3:
    fmt.Println("Três")
    // Para aqui (sem fallthrough)
}
// Saída: Dois, Três
```

**Analogia:**

- Quando um **dominó cai**, ele **empurra o próximo**
- Quando um **case executa** com `fallthrough`, ele **continua para o próximo**
- É como uma **corrente** onde uma ação leva à próxima

---

## 11. Type Switch: O Identificador de Objetos

### Analogia: Identificar Objetos em uma Caixa

Pense em type switch como **identificar objetos em uma caixa**:

**Caixa física:**

```
┌─────────────┐
│   CAIXA     │
│             │
│  [Objeto?]  │
│             │
│ É um livro? │
│ É uma bola? │
│ É um lápis? │
└─────────────┘
```

**Em Go:**

```go
var objeto interface{} = 42

switch v := objeto.(type) {
case int:
    fmt.Printf("É um número: %d\n", v)
case string:
    fmt.Printf("É um texto: %s\n", v)
case bool:
    fmt.Printf("É verdadeiro/falso: %v\n", v)
default:
    fmt.Println("Tipo desconhecido")
}
```

**Analogia:**

- Você **pega um objeto** da caixa
- **Verifica o que é** (tipo)
- **Faz algo diferente** dependendo do tipo
- É como **classificar objetos** por categoria

---

## 12. Validação: O Guarda de Segurança

### Analogia: Guarda de Segurança

Pense em conditionals para validação como um **guarda de segurança**:

**Guarda físico:**

```
┌─────────────────────┐
│   ENTRADA           │
│                     │
│  [Verificar ID]     │
│                     │
│  ID válido?         │
│  SIM → Entrar       │
│  NÃO → Bloquear     │
└─────────────────────┘
```

**Em Go:**

```go
email := "usuario@email.com"
senha := "senha123"

if email == "" {
    fmt.Println("Email não pode ser vazio")
} else if len(senha) < 6 {
    fmt.Println("Senha muito curta")
} else {
    fmt.Println("Pode entrar")
}
```

**Analogia:**

- O guarda **verifica** se você pode entrar
- O programa **verifica** se os dados são válidos
- Ambos **bloqueiam** se algo estiver errado
- Ambos **permitem** se tudo estiver certo

---

## 13. Tratamento de Erros: O Detector de Problemas

### Analogia: Detector de Problemas

Pense em conditionals para erros como um **detector de problemas**:

**Detector físico:**

```
┌─────────────────────┐
│   OPERAÇÃO          │
│                     │
│  [Executar]         │
│                     │
│  Deu erro?          │
│  SIM → Avisar       │
│  NÃO → Continuar    │
└─────────────────────┘
```

**Em Go:**

```go
resultado, err := dividir(10, 2)

if err != nil {
    fmt.Printf("Erro: %v\n", err)
    return
}

fmt.Printf("Resultado: %.2f\n", resultado)
```

**Analogia:**

- O detector **verifica** se há problemas
- O programa **verifica** se há erros
- Se houver problema → **para e avisa**
- Se não houver → **continua normalmente**

---

## 14. `if` Aninhado: A Decisão Dentro de Outra Decisão

### Analogia: Portas com Chaves

Pense em `if` aninhado como **portas com chaves**:

**Cenário físico:**

```
┌─────────────┐
│ PORTA 1     │  ← Precisa ter 18 anos
│   [Abrir]   │
│      ↓      │
│ ┌─────────┐ │
│ │PORTA 2  │ │  ← Precisa ter carteira
│ │ [Abrir] │ │
│ └─────────┘ │
└─────────────┘
```

**Em Go:**

```go
idade := 25
temCarteira := true

if idade >= 18 {
    if temCarteira {
        fmt.Println("Pode dirigir")
    } else {
        fmt.Println("Precisa tirar carteira")
    }
} else {
    fmt.Println("Menor de idade")
}
```

**Analogia:**

- Você precisa **abrir a primeira porta** (idade >= 18)
- **Depois** pode tentar abrir a segunda (tem carteira)
- É como **chaves dentro de chaves** - precisa de todas

---

## 15. Comparação: `if-else` vs `switch`

### Analogia: Escada vs Elevador

Pense na diferença como **escalar uma escada vs usar um elevador**:

**Escada (if-else):**

```
┌─┐
│ │ Degrau 1 (if)
├─┤
│ │ Degrau 2 (else if)
├─┤
│ │ Degrau 3 (else if)
├─┤
│ │ Chão (else)
└─┘
```

**Elevador (switch):**

```
┌─────────────┐
│  ELEVADOR   │
├─────────────┤
│ [Andar 1]   │
│ [Andar 2]   │
│ [Andar 3]   │
│ [Térreo]    │
└─────────────┘
```

**Em Go:**

```go
// Escada (if-else) - passo a passo
if cor == "vermelho" {
    fmt.Println("Pare!")
} else if cor == "amarelo" {
    fmt.Println("Atenção!")
} else if cor == "verde" {
    fmt.Println("Siga!")
}

// Elevador (switch) - direto ao destino
switch cor {
case "vermelho":
    fmt.Println("Pare!")
case "amarelo":
    fmt.Println("Atenção!")
case "verde":
    fmt.Println("Siga!")
}
```

**Analogia:**

- **Escada (if-else)**: Você **sobe degrau por degrau**, verificando cada um
- **Elevador (switch)**: Você **pressiona o botão** e vai direto ao andar
- Ambos chegam ao mesmo lugar, mas de formas diferentes!

---

## 16. Resumo Visual: Conditionals como Decisões

Pense em conditionals como um **fluxograma de decisões**:

```
┌─────────────────┐
│   INÍCIO        │
└────────┬────────┘
         │
         ▼
    ┌─────────┐
    │ Condição│
    └────┬────┘
         │
    ┌────┴────┐
    │         │
   SIM       NÃO
    │         │
    ▼         ▼
┌──────┐  ┌──────┐
│ Ação │  │ Ação │
│  1   │  │  2   │
└──────┘  └──────┘
```

**Em código:**

```go
if condicao {
    // Ação 1 (SIM)
} else {
    // Ação 2 (NÃO)
}
```

---

## 17. Analogia Final: O GPS

Conditionals são como um **GPS**:

**GPS físico:**

```
┌─────────────────────┐
│   GPS               │
├─────────────────────┤
│ Você está em:      │
│ Rua A               │
│                     │
│ Se virar à direita: │
│ → Vai para Rua B    │
│                     │
│ Se virar à esquerda:│
│ → Vai para Rua C    │
│                     │
│ Se seguir reto:     │
│ → Vai para Rua D    │
└─────────────────────┘
```

**Em Go:**

```go
direcao := "direita"

switch direcao {
case "direita":
    fmt.Println("Vai para Rua B")
case "esquerda":
    fmt.Println("Vai para Rua C")
case "reto":
    fmt.Println("Vai para Rua D")
}
```

**Por que funciona:**

- O GPS **decide** qual rota seguir baseado na sua escolha
- O conditional **decide** qual código executar baseado na condição
- Ambos **guiam** você pelo caminho certo!

---

## 18. Regra de Ouro Simples

**Conditionals = Decisões do Programa**

- **`if`** = "Se isso, faça aquilo"
- **`if-else`** = "Se isso, faça aquilo; senão, faça outra coisa"
- **`if-else if`** = "Se isso, faça aquilo; senão se aquilo, faça isso; senão..."
- **`switch`** = "Dependendo do valor, faça isso ou aquilo"

**Quando usar:**

- Precisa **tomar uma decisão**? → Use conditional!
- Precisa **verificar algo**? → Use conditional!
- Precisa **escolher entre opções**? → Use conditional!

---

Agora que você entendeu os conceitos de forma simplificada, vamos praticar com exercícios na próxima parte!

