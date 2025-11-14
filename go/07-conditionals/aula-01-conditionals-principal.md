# Módulo 7: Conditionals em Go

## Aula 1: Conditionals - Controle de Fluxo Condicional

Olá! Bem-vindo ao sétimo módulo. Até agora você aprendeu sobre tipos de dados, variáveis, structs e como organizar informações. Mas como fazer seu programa **tomar decisões**? Como fazer ele executar código diferente dependendo de situações diferentes?

É aqui que entram os **conditionals** (condicionais) - estruturas que permitem ao seu programa escolher qual caminho seguir baseado em condições.

Nesta aula, vamos mergulhar profundamente em `if`, `if-else` e `switch` - as ferramentas fundamentais para controlar o fluxo de execução do seu programa em Go.

---

## 1. O Que São Conditionals?

**Conditionals** (condicionais) são estruturas que permitem executar código **apenas se uma condição for verdadeira** ou escolher entre diferentes caminhos de execução baseado em condições.

### Por Que Precisamos de Conditionals?

**Sem conditionals**, seu programa sempre executaria as mesmas instruções na mesma ordem:

```go
func main() {
    fmt.Println("Sempre executa isso")
    fmt.Println("Sempre executa isso também")
    fmt.Println("Nunca muda")
}
```

**Com conditionals**, seu programa pode tomar decisões:

```go
func main() {
    idade := 20

    if idade >= 18 {
        fmt.Println("Você é maior de idade")
    } else {
        fmt.Println("Você é menor de idade")
    }
}
```

**Vantagens:**

- ✅ Código reativo - responde a diferentes situações
- ✅ Lógica de negócio - implementa regras do mundo real
- ✅ Validação - verifica dados antes de usar
- ✅ Controle de fluxo - decide qual código executar

---

## 2. A Estrutura `if` Básica

A estrutura `if` é a mais fundamental. Ela executa um bloco de código **apenas se** uma condição for verdadeira.

### Sintaxe Básica

```go
if condicao {
    // código executado se condição for verdadeira
}
```

### Características Importantes do `if` em Go

1. **Sem parênteses**: Go não requer parênteses ao redor da condição
2. **Chaves obrigatórias**: Mesmo para uma linha, as chaves `{}` são obrigatórias
3. **Condição booleana**: A condição deve resultar em `true` ou `false`

### Exemplo Básico

```go
idade := 20

if idade >= 18 {
    fmt.Println("Você é maior de idade")
}
```

**O que acontece:**

- Se `idade >= 18` for `true` → executa o código dentro das chaves
- Se `idade >= 18` for `false` → não executa nada, continua o programa

---

## 3. `if` com Inicialização Opcional

Go permite **declarar e inicializar uma variável** antes da condição do `if`. Essa variável existe **apenas dentro do escopo do `if`**.

### Sintaxe

```go
if inicializacao; condicao {
    // código
}
```

### Exemplo Prático

```go
// Declarar e usar variável no if
if num := 42; num > 40 {
    fmt.Printf("O número %d é maior que 40\n", num)
    // num está disponível aqui
}

// num não existe aqui fora do if
// fmt.Println(num) // ERRO: undefined: num
```

### Por Que Usar Inicialização no `if`?

**Vantagens:**

- ✅ **Escopo limitado**: Variável existe apenas onde é necessária
- ✅ **Código mais limpo**: Declaração próxima ao uso
- ✅ **Evita poluição**: Não cria variáveis desnecessárias no escopo externo

### Exemplo com Funções

```go
// Verificar erro diretamente
if resultado, err := dividir(10, 2); err != nil {
    fmt.Printf("Erro: %v\n", err)
} else {
    fmt.Printf("Resultado: %.2f\n", resultado)
}
```

---

## 4. A Estrutura `if-else`

Quando você precisa executar código **alternativo** se a condição for falsa, use `else`.

### Sintaxe

```go
if condicao {
    // código se condição for verdadeira
} else {
    // código se condição for falsa
}
```

### Exemplo

```go
temperatura := 25

if temperatura > 30 {
    fmt.Println("Está muito quente!")
} else {
    fmt.Println("Temperatura agradável")
}
```

**Fluxo:**

1. Avalia `temperatura > 30`
2. Se `true` → executa primeiro bloco
3. Se `false` → executa bloco do `else`

---

## 5. `if-else if-else` (Múltiplas Condições)

Quando você precisa verificar **múltiplas condições** em sequência, use `else if`.

### Sintaxe

```go
if condicao1 {
    // código se condicao1 for verdadeira
} else if condicao2 {
    // código se condicao2 for verdadeira
} else if condicao3 {
    // código se condicao3 for verdadeira
} else {
    // código se nenhuma condição for verdadeira
}
```

### Exemplo: Sistema de Notas

```go
nota := 85

if nota >= 90 {
    fmt.Println("Nota A - Excelente!")
} else if nota >= 80 {
    fmt.Println("Nota B - Muito bom!")
} else if nota >= 70 {
    fmt.Println("Nota C - Bom")
} else if nota >= 60 {
    fmt.Println("Nota D - Regular")
} else {
    fmt.Println("Nota F - Reprovado")
}
```

**Comportamento:**

- Avalia condições **de cima para baixo**
- Executa o **primeiro bloco** cuja condição for verdadeira
- Se nenhuma for verdadeira, executa o `else` (se existir)
- **Para na primeira condição verdadeira** - não avalia as seguintes

### Ordem Importante

```go
// ❌ ERRADO: ordem incorreta
if nota >= 60 {
    fmt.Println("Aprovado")  // Sempre executa para notas >= 60
} else if nota >= 70 {
    fmt.Println("Bom")  // Nunca executa!
}

// ✅ CORRETO: ordem correta (maior para menor)
if nota >= 90 {
    fmt.Println("Excelente")
} else if nota >= 80 {
    fmt.Println("Muito bom")
} else if nota >= 70 {
    fmt.Println("Bom")
}
```

---

## 6. `if` com Operadores Lógicos

Você pode combinar múltiplas condições usando operadores lógicos.

### Operadores Lógicos em Go

| Operador | Nome      | Descrição                                                |
| -------- | --------- | -------------------------------------------------------- |
| `&&`     | E (AND)   | Verdadeiro se **ambas** condições forem verdadeiras      |
| `\|\|`   | OU (OR)   | Verdadeiro se **pelo menos uma** condição for verdadeira |
| `!`      | NÃO (NOT) | Inverte o valor booleano                                 |

### Exemplo com `&&` (E)

```go
salario := 3000
horasTrabalhadas := 45

if salario > 2500 && horasTrabalhadas > 40 {
    fmt.Println("Você recebeu horas extras!")
}
```

**Lógica:** Executa apenas se **salário > 2500 E horas > 40** forem verdadeiros.

### Exemplo com `||` (OU)

```go
if salario < 2000 || horasTrabalhadas < 30 {
    fmt.Println("Atenção: salário ou horas abaixo do esperado")
}
```

**Lógica:** Executa se **salário < 2000 OU horas < 30** (ou ambos).

### Exemplo com `!` (NÃO)

```go
chovendo := false

if !chovendo {
    fmt.Println("Pode sair sem guarda-chuva")
}
```

**Lógica:** Executa se `chovendo` for `false` (ou seja, se **não** está chovendo).

### Short-Circuit Evaluation

Go usa **avaliação de curto-circuito**:

- `&&`: Se primeira condição for `false`, **não avalia** a segunda
- `||`: Se primeira condição for `true`, **não avalia** a segunda

```go
// Se valor > 10 for false, valor < 20 NÃO é avaliado
if valor > 10 && valor < 20 {
    fmt.Println("Entre 10 e 20")
}
```

---

## 7. `if` Aninhado

Você pode colocar um `if` dentro de outro `if` (aninhamento).

### Sintaxe

```go
if condicao1 {
    if condicao2 {
        // código se ambas condições forem verdadeiras
    }
}
```

### Exemplo

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
    fmt.Println("Menor de idade, não pode dirigir")
}
```

**Quando usar:**

- ✅ Quando condições são **dependentes** (segunda só faz sentido se primeira for verdadeira)
- ✅ Quando precisa verificar **múltiplos níveis** de condições

**Quando evitar:**

- ❌ Quando pode usar operadores lógicos (`&&`, `||`)
- ❌ Quando aninhamento fica muito profundo (> 2-3 níveis)

### Alternativa com Operadores Lógicos

```go
// Em vez de aninhar, pode usar &&
if idade >= 18 && temCarteira {
    fmt.Println("Pode dirigir")
} else if idade >= 18 {
    fmt.Println("Precisa tirar carteira")
} else {
    fmt.Println("Menor de idade")
}
```

---

## 8. A Estrutura `switch`

O `switch` é uma forma **mais limpa e legível** de lidar com múltiplas condições quando você está comparando uma variável contra vários valores possíveis.

### Sintaxe Básica

```go
switch variavel {
case valor1:
    // código se variavel == valor1
case valor2:
    // código se variavel == valor2
case valor3:
    // código se variavel == valor3
default:
    // código se nenhum case corresponder
}
```

### Exemplo Básico

```go
dia := "segunda"

switch dia {
case "segunda":
    fmt.Println("Início da semana")
case "terça", "quarta", "quinta":
    fmt.Println("Meio da semana")
case "sexta":
    fmt.Println("Final da semana")
case "sábado", "domingo":
    fmt.Println("Fim de semana!")
default:
    fmt.Println("Dia inválido")
}
```

### Características do `switch` em Go

1. **Sem `break`**: Go **não precisa** de `break` - não há fall-through automático
2. **Múltiplos valores**: Um `case` pode ter vários valores separados por vírgula
3. **`default` opcional**: Executado se nenhum `case` corresponder
4. **Comparação de igualdade**: Compara usando `==`

---

## 9. `switch` sem Expressão (como `if-else`)

Você pode usar `switch` **sem uma variável**, apenas com condições. Isso funciona como uma cadeia de `if-else if`.

### Sintaxe

```go
switch {
case condicao1:
    // código
case condicao2:
    // código
default:
    // código
}
```

### Exemplo

```go
idade := 25

switch {
case idade < 18:
    fmt.Println("Menor de idade")
case idade >= 18 && idade < 65:
    fmt.Println("Adulto")
case idade >= 65:
    fmt.Println("Idoso")
}
```

**Quando usar:**

- ✅ Quando condições são **complexas** (não apenas igualdade)
- ✅ Quando quer **legibilidade** melhor que `if-else if`
- ✅ Quando condições são **independentes**

---

## 10. `switch` com Inicialização

Assim como `if`, `switch` também suporta inicialização.

### Sintaxe

```go
switch inicializacao; variavel {
case valor1:
    // código
}
```

### Exemplo

```go
switch hora := time.Now().Hour(); {
case hora < 12:
    fmt.Println("Bom dia!")
case hora < 18:
    fmt.Println("Boa tarde!")
default:
    fmt.Println("Boa noite!")
}
```

---

## 11. `switch` com `fallthrough`

Por padrão, Go **não faz fall-through** (não continua para o próximo `case`). Mas você pode forçar isso com `fallthrough`.

### Sintaxe

```go
switch variavel {
case valor1:
    // código
    fallthrough  // Continua para o próximo case
case valor2:
    // código
}
```

### Exemplo

```go
numero := 2

switch numero {
case 1:
    fmt.Println("Um")
    fallthrough
case 2:
    fmt.Println("Dois")
    fallthrough
case 3:
    fmt.Println("Três")
    // Sem fallthrough, para aqui
default:
    fmt.Println("Outro número")
}
// Saída: Dois, Três
```

**Quando usar `fallthrough`:**

- ⚠️ **Raramente necessário** - a maioria dos casos não precisa
- ✅ Quando múltiplos casos devem executar o mesmo código
- ✅ Quando implementa lógica específica que requer continuidade

**Cuidado:** `fallthrough` pode tornar o código difícil de entender. Use com moderação.

---

## 12. Type Switch

Um **type switch** permite verificar o **tipo** de uma variável `interface{}`.

### Sintaxe

```go
switch v := variavel.(type) {
case Tipo1:
    // v é do tipo Tipo1
case Tipo2:
    // v é do tipo Tipo2
default:
    // tipo desconhecido
}
```

### Exemplo

```go
var valor interface{} = 42

switch v := valor.(type) {
case int:
    fmt.Printf("É um inteiro: %d\n", v)
case string:
    fmt.Printf("É uma string: %s\n", v)
case bool:
    fmt.Printf("É um booleano: %v\n", v)
default:
    fmt.Printf("Tipo desconhecido: %T\n", v)
}
```

**Quando usar:**

- ✅ Quando trabalha com `interface{}` e precisa verificar tipos
- ✅ Quando processa dados de tipos desconhecidos
- ✅ Quando implementa funções genéricas

---

## 13. `switch` com Múltiplos Valores por `case`

Um `case` pode corresponder a **múltiplos valores** separados por vírgula.

### Exemplo

```go
mes := "janeiro"

switch mes {
case "dezembro", "janeiro", "fevereiro":
    fmt.Println("Verão")
case "março", "abril", "maio":
    fmt.Println("Outono")
case "junho", "julho", "agosto":
    fmt.Println("Inverno")
case "setembro", "outubro", "novembro":
    fmt.Println("Primavera")
default:
    fmt.Println("Mês inválido")
}
```

**Vantagem:** Mais limpo que múltiplos `if-else if` com `||`.

---

## 14. Comparação: `if-else` vs `switch`

### Quando Usar `if-else`

✅ **Use `if-else` quando:**

- Condições são **complexas** (operadores lógicos, comparações variadas)
- Precisa de **intervalos** ou comparações não-igualdade
- Condições são **independentes** e não relacionadas a uma única variável

**Exemplo:**

```go
if idade >= 18 && temCarteira && saldo > 0 {
    // lógica complexa
}
```

### Quando Usar `switch`

✅ **Use `switch` quando:**

- Comparando uma **variável** contra **múltiplos valores específicos**
- Condições são principalmente de **igualdade** (`==`)
- Quer código **mais legível** e organizado
- Tem **muitos casos** para a mesma variável

**Exemplo:**

```go
switch dia {
case "segunda", "terça", "quarta", "quinta", "sexta":
    fmt.Println("Dia útil")
case "sábado", "domingo":
    fmt.Println("Fim de semana")
}
```

### Comparação Direta

**Com `if-else`:**

```go
cor := "vermelho"
if cor == "vermelho" {
    fmt.Println("Pare!")
} else if cor == "amarelo" {
    fmt.Println("Atenção!")
} else if cor == "verde" {
    fmt.Println("Siga!")
} else {
    fmt.Println("Cor desconhecida")
}
```

**Com `switch` (mais limpo):**

```go
switch cor {
case "vermelho":
    fmt.Println("Pare!")
case "amarelo":
    fmt.Println("Atenção!")
case "verde":
    fmt.Println("Siga!")
default:
    fmt.Println("Cor desconhecida")
}
```

---

## 15. Validação e Tratamento de Erros com Conditionals

Conditionals são fundamentais para **validação** e **tratamento de erros**.

### Validação de Dados

```go
email := "usuario@email.com"
senha := "senha123"

if email == "" {
    fmt.Println("Email não pode ser vazio")
} else if len(senha) < 6 {
    fmt.Println("Senha muito curta")
} else if len(email) < 5 {
    fmt.Println("Email inválido")
} else {
    fmt.Println("Dados válidos!")
}
```

### Tratamento de Erros

```go
// Função que retorna valor e erro
resultado, err := dividir(10, 2)

if err != nil {
    fmt.Printf("Erro: %v\n", err)
    return
}

fmt.Printf("Resultado: %.2f\n", resultado)
```

### Verificação de Map Lookup

```go
cores := map[string]string{
    "vermelho": "#FF0000",
    "verde":    "#00FF00",
}

cor := "vermelho"

// Verificar se chave existe
if hex, existe := cores[cor]; existe {
    fmt.Printf("Cor %s: %s\n", cor, hex)
} else {
    fmt.Printf("Cor %s não encontrada\n", cor)
}
```

---

## 16. Exemplos Práticos Completos

### Exemplo 1: Sistema de Classificação de Idade

```go
func classificarIdade(idade int) string {
    switch {
    case idade < 0:
        return "Idade inválida"
    case idade < 13:
        return "Criança"
    case idade < 18:
        return "Adolescente"
    case idade < 65:
        return "Adulto"
    default:
        return "Idoso"
    }
}
```

### Exemplo 2: Calculadora com Validação

```go
func dividir(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("divisão por zero")
    }
    return a / b, nil
}

func main() {
    if resultado, err := dividir(10, 2); err != nil {
        fmt.Printf("Erro: %v\n", err)
    } else {
        fmt.Printf("Resultado: %.2f\n", resultado)
    }
}
```

### Exemplo 3: Sistema de Estações do Ano

```go
func obterEstacao(mes string) string {
    switch mes {
    case "dezembro", "janeiro", "fevereiro":
        return "Verão"
    case "março", "abril", "maio":
        return "Outono"
    case "junho", "julho", "agosto":
        return "Inverno"
    case "setembro", "outubro", "novembro":
        return "Primavera"
    default:
        return "Mês inválido"
    }
}
```

---

## 17. Resumo dos Conceitos

| Conceito                   | Sintaxe                          | Quando Usar                                |
| -------------------------- | -------------------------------- | ------------------------------------------ |
| **if básico**              | `if cond { }`                    | Executar código se condição for verdadeira |
| **if-else**                | `if cond { } else { }`           | Escolher entre dois caminhos               |
| **if-else if**             | `if cond1 { } else if cond2 { }` | Múltiplas condições em sequência           |
| **if com inicialização**   | `if init; cond { }`              | Declarar variável no escopo do if          |
| **switch básico**          | `switch var { case val: }`       | Comparar variável contra múltiplos valores |
| **switch sem expressão**   | `switch { case cond: }`          | Múltiplas condições complexas              |
| **switch com fallthrough** | `case val: fallthrough`          | Continuar para próximo case                |
| **type switch**            | `switch v := x.(type) { }`       | Verificar tipo de interface{}              |

---

## 18. Regras Importantes

### Regras do `if`

1. ✅ **Sem parênteses** ao redor da condição
2. ✅ **Chaves obrigatórias** mesmo para uma linha
3. ✅ **Condição deve ser booleana** (`true` ou `false`)
4. ✅ **Variável de inicialização** existe apenas no escopo do `if`

### Regras do `switch`

1. ✅ **Sem `break` necessário** - não há fall-through automático
2. ✅ **Múltiplos valores** por `case` separados por vírgula
3. ✅ **`default` é opcional** mas recomendado
4. ✅ **Comparação por igualdade** (`==`)
5. ✅ **`fallthrough`** continua para próximo `case`

---

## 19. Quando Usar Cada Estrutura

### Use `if` quando:

- ✅ Condições são **complexas** (operadores lógicos)
- ✅ Precisa de **intervalos** ou comparações não-igualdade
- ✅ Condições são **independentes**
- ✅ Validação e tratamento de erros

### Use `switch` quando:

- ✅ Comparando **uma variável** contra **múltiplos valores**
- ✅ Condições são principalmente de **igualdade**
- ✅ Quer código **mais legível** e organizado
- ✅ Tem **muitos casos** para a mesma variável

---

## Conclusão

Conditionals são fundamentais em programação porque permitem:

1. **Tomar decisões** - Executar código baseado em condições
2. **Validar dados** - Verificar antes de usar
3. **Tratar erros** - Responder a situações inesperadas
4. **Implementar lógica de negócio** - Regras do mundo real

**Conceitos-chave:**

- `if` para condições gerais e complexas
- `if-else` para escolhas binárias
- `if-else if` para múltiplas condições
- `switch` para comparações de igualdade múltiplas
- Type switch para verificação de tipos

Na próxima parte desta aula, vamos simplificar esses conceitos com analogias do dia a dia para fixar melhor o aprendizado!

