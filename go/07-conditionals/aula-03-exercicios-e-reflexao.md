# Módulo 7: Conditionals em Go

## Aula 3 - Exercícios e Reflexão

Agora é hora de colocar a mão na massa! Complete os exercícios abaixo e reflita sobre as questões propostas.

---

## Exercícios Práticos

### Exercício 1: Sistema de Classificação de Notas

Crie um arquivo `exercicio1.go` e implemente o seguinte:

**Tarefa:**

1. Crie uma função `ClassificarNota(nota float64) string` que recebe uma nota e retorna a classificação:

   - 90-100: "A - Excelente"
   - 80-89: "B - Muito Bom"
   - 70-79: "C - Bom"
   - 60-69: "D - Regular"
   - 0-59: "F - Reprovado"
   - Nota inválida (< 0 ou > 100): "Nota inválida"

2. Use `if-else if-else` para implementar a lógica.

3. No `main`, teste a função com diferentes notas e exiba os resultados.

**Exemplo de uso esperado:**

```go
fmt.Println(ClassificarNota(95))  // "A - Excelente"
fmt.Println(ClassificarNota(75))  // "C - Bom"
fmt.Println(ClassificarNota(45))  // "F - Reprovado"
fmt.Println(ClassificarNota(150)) // "Nota inválida"
```

---

### Exercício 2: Calculadora com Validação

Crie um arquivo `exercicio2.go` e implemente uma calculadora com validação:

**Tarefa:**

1. Crie funções para as operações básicas:

   - `Somar(a, b float64) float64`
   - `Subtrair(a, b float64) float64`
   - `Multiplicar(a, b float64) float64`
   - `Dividir(a, b float64) (float64, error)` - retorna erro se divisão por zero

2. Crie uma função `Calcular(operacao string, a, b float64) (float64, error)` que:

   - Recebe o nome da operação ("soma", "subtracao", "multiplicacao", "divisao")
   - Usa `switch` para escolher a operação
   - Retorna o resultado ou erro

3. No `main`, teste todas as operações, incluindo casos de erro (divisão por zero, operação inválida).

**Exemplo de uso:**

```go
resultado, err := Calcular("soma", 10, 5)
if err != nil {
    fmt.Printf("Erro: %v\n", err)
} else {
    fmt.Printf("Resultado: %.2f\n", resultado)
}
```

---

### Exercício 3: Sistema de Estações do Ano

Crie um arquivo `exercicio3.go` e implemente um sistema que identifica a estação do ano:

**Tarefa:**

1. Crie uma função `ObterEstacao(mes string) string` que:

   - Recebe o nome do mês (em português, minúsculo)
   - Retorna a estação do ano correspondente
   - Use `switch` com múltiplos valores por `case`

2. Estações:

   - Verão: dezembro, janeiro, fevereiro
   - Outono: março, abril, maio
   - Inverno: junho, julho, agosto
   - Primavera: setembro, outubro, novembro

3. Crie uma função `ObterEstacaoPorNumero(numero int) string` que:

   - Recebe um número de 1 a 12 (representando os meses)
   - Retorna a estação
   - Use `switch` básico

4. No `main`, teste ambas as funções.

**Dica:** Use `switch` com múltiplos valores para tornar o código mais limpo.

---

### Exercício 4: Validador de Senha

Crie um arquivo `exercicio4.go` e implemente um validador de senha:

**Tarefa:**

1. Crie uma função `ValidarSenha(senha string) (bool, []string)` que:

   - Recebe uma senha
   - Retorna `true` se válida, `false` se inválida
   - Retorna uma lista de erros encontrados

2. Regras de validação:

   - Mínimo de 8 caracteres
   - Máximo de 20 caracteres
   - Deve conter pelo menos uma letra maiúscula
   - Deve conter pelo menos uma letra minúscula
   - Deve conter pelo menos um número
   - Deve conter pelo menos um caractere especial (!@#$%^&\*)

3. Use `if-else if` para verificar cada regra e adicionar erros à lista.

4. No `main`, teste com diferentes senhas e exiba os erros encontrados.

**Exemplo:**

```go
valida, erros := ValidarSenha("senha123")
if !valida {
    fmt.Println("Senha inválida. Erros:")
    for _, erro := range erros {
        fmt.Printf("  - %s\n", erro)
    }
}
```

**Dica:** Você precisará iterar sobre a string para verificar caracteres. Use loops (que aprenderá depois) ou funções como `strings.ContainsAny()`.

---

### Exercício 5: Sistema de Desconto

Crie um arquivo `exercicio5.go` e implemente um sistema de desconto:

**Tarefa:**

1. Crie uma função `CalcularDesconto(valor float64, tipoCliente string, temCupom bool) float64` que:

   - Recebe o valor da compra, tipo de cliente e se tem cupom
   - Retorna o valor do desconto

2. Regras de desconto:

   - Cliente "VIP": 15% de desconto
   - Cliente "Premium": 10% de desconto
   - Cliente "Regular": 5% de desconto
   - Se tiver cupom: +5% de desconto adicional (acumula com desconto do cliente)
   - Desconto máximo: 25%

3. Use `switch` para o tipo de cliente e `if` para o cupom.

4. No `main`, teste diferentes combinações.

**Exemplo:**

```go
desconto := CalcularDesconto(1000, "VIP", true)
// VIP: 15% + cupom: 5% = 20% de desconto
// Mas máximo é 25%, então 20% é aplicado
fmt.Printf("Desconto: R$ %.2f\n", desconto)
```

---

### Exercício 6: Type Switch para Processamento

Crie um arquivo `exercicio6.go` e implemente um processador genérico:

**Tarefa:**

1. Crie uma função `Processar(valor interface{}) string` que:

   - Recebe um valor de qualquer tipo (`interface{}`)
   - Usa `type switch` para identificar o tipo
   - Retorna uma string descrevendo o processamento

2. Trate os seguintes tipos:

   - `int`: retorna "Inteiro: [valor] (dobro: [valor*2])"
   - `float64`: retorna "Decimal: [valor] (quadrado: [valor²])"
   - `string`: retorna "Texto: [valor] (tamanho: [len])"
   - `bool`: retorna "Booleano: [valor] (invertido: [!valor])"
   - `[]int`: retorna "Slice: [valores] (soma: [soma dos valores])"
   - Outros: retorna "Tipo não suportado: [tipo]"

3. No `main`, teste com diferentes tipos.

**Exemplo:**

```go
fmt.Println(Processar(42))           // Inteiro: 42 (dobro: 84)
fmt.Println(Processar(3.14))       // Decimal: 3.14 (quadrado: 9.86)
fmt.Println(Processar("Go"))       // Texto: Go (tamanho: 2)
fmt.Println(Processar(true))       // Booleano: true (invertido: false)
fmt.Println(Processar([]int{1,2,3})) // Slice: [1 2 3] (soma: 6)
```

---

## Perguntas de Reflexão

### Reflexão 1: Por Que Conditionals São Fundamentais?

Você aprendeu que conditionals permitem ao programa tomar decisões.

**Pergunta:**

1. Pense em um programa que você usa no dia a dia (aplicativo, site, etc.). Quais decisões esse programa precisa tomar? Dê pelo menos 3 exemplos concretos.

2. Sem conditionals, um programa sempre executaria as mesmas instruções na mesma ordem. Por que isso seria um problema? Dê exemplos de situações onde isso tornaria o programa inútil.

3. Conditionals são a base da "lógica de negócio" em software. O que isso significa? Por que é importante separar a lógica de negócio do resto do código?

4. Em Go, `if` não precisa de parênteses mas precisa de chaves. Por que você acha que Go fez essa escolha de design? Qual é o impacto na legibilidade do código?

**Sua resposta deve incluir:** Exemplos concretos de programas reais, análise sobre a importância de conditionals, e pensamento crítico sobre design de linguagem.

---

### Reflexão 2: if-else vs switch - Quando Usar Cada Um?

Você aprendeu que tanto `if-else` quanto `switch` podem resolver problemas similares, mas cada um tem seu lugar.

**Pergunta:**

1. Quando você escolheria usar `if-else` em vez de `switch`? Dê exemplos concretos de situações onde `if-else` é mais apropriado.

2. Quando você escolheria usar `switch` em vez de `if-else`? Dê exemplos concretos de situações onde `switch` é mais apropriado.

3. Em Go, `switch` não precisa de `break` (não há fall-through automático). Por que isso é uma vantagem? Compare com outras linguagens que você conhece (se conhecer).

4. `switch` sem expressão (apenas com `case` e condições) é equivalente a `if-else if`. Quando você usaria cada um? Há alguma diferença prática além da legibilidade?

**Sua resposta deve demonstrar:** Compreensão das diferenças entre `if-else` e `switch`, capacidade de escolher a ferramenta certa para cada situação, e conhecimento sobre as características únicas de Go.

---

### Reflexão 3: Validação e Tratamento de Erros

Conditionals são fundamentais para validação e tratamento de erros.

**Pergunta:**

1. Por que validação é importante? O que acontece se um programa não validar dados de entrada? Dê exemplos de problemas reais que podem ocorrer.

2. Em Go, funções frequentemente retornam `(valor, error)`. Como conditionals são usados para tratar esses erros? Por que esse padrão é comum em Go?

3. Pense em um sistema de login. Quais validações você faria antes de permitir o login? Liste pelo menos 5 validações e explique por que cada uma é importante.

4. Validação e tratamento de erros tornam o código mais longo e complexo. Por que isso é aceitável (ou até desejável)? Qual é o trade-off?

**Sua resposta deve incluir:** Exemplos de validação em sistemas reais, análise sobre o padrão de erros em Go, e pensamento sobre trade-offs em design de software.

---

### Reflexão 4: Lógica de Negócio e Conditionals

Conditionals implementam a "lógica de negócio" - as regras que governam como o sistema funciona.

**Pergunta:**

1. O que é "lógica de negócio"? Dê exemplos de regras de negócio que você implementaria com conditionals (ex: "desconto de 10% para clientes VIP", "usuários menores de 18 anos não podem acessar conteúdo adulto", etc.).

2. Pense em um sistema de e-commerce. Quais regras de negócio você implementaria? Liste pelo menos 5 regras e explique como conditionals seriam usados.

3. Às vezes, regras de negócio mudam. Como você organizaria conditionals para facilitar mudanças futuras? O que torna o código "fácil de modificar"?

4. Conditionals podem se tornar muito complexos quando há muitas regras. Como você evitaria conditionals excessivamente complexos? Quais técnicas ou padrões você usaria?

**Sua resposta deve demonstrar:** Compreensão de lógica de negócio, capacidade de identificar regras em sistemas reais, e pensamento sobre manutenibilidade de código.

---

### Reflexão 5: Performance e Legibilidade

Conditionals têm implicações tanto em performance quanto em legibilidade.

**Pergunta:**

1. Quando você tem múltiplas condições em um `if-else if`, a ordem importa? Por quê? Dê um exemplo onde a ordem faz diferença na performance.

2. Short-circuit evaluation (`&&` e `||` param de avaliar quando possível) pode melhorar performance. Dê um exemplo onde isso é importante.

3. `switch` vs `if-else if` - há diferença de performance? Quando isso importa? Quando não importa?

4. Legibilidade vs Performance - às vezes código mais legível é um pouco mais lento. Quando você priorizaria legibilidade? Quando priorizaria performance? Dê exemplos.

5. Conditionals aninhados podem tornar o código difícil de ler. Qual é o limite? Quando você refatoraria conditionals aninhados? Como?

**Sua resposta deve demonstrar:** Compreensão de performance, análise de trade-offs, e capacidade de tomar decisões informadas sobre legibilidade vs performance.

---

## Como Entregar

Crie os seguintes arquivos na pasta `07-conditionals/`:

1. **`exercicio1.go`** - Sistema de Classificação de Notas
2. **`exercicio2.go`** - Calculadora com Validação
3. **`exercicio3.go`** - Sistema de Estações do Ano
4. **`exercicio4.go`** - Validador de Senha
5. **`exercicio5.go`** - Sistema de Desconto
6. **`exercicio6.go`** - Type Switch para Processamento
7. **`reflexoes.md`** - Respostas às 5 perguntas de reflexão

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
func ClassificarNota(nota float64) string {
    if nota < 0 || nota > 100 {
        return "Nota inválida"
    } else if nota >= 90 {
        return "A - Excelente"
    } else if nota >= 80 {
        return "B - Muito Bom"
    }
    // ... continue
}
```

### Dica para Exercício 2:

```go
func Calcular(operacao string, a, b float64) (float64, error) {
    switch operacao {
    case "soma":
        return Somar(a, b), nil
    case "subtracao":
        return Subtrair(a, b), nil
    case "multiplicacao":
        return Multiplicar(a, b), nil
    case "divisao":
        return Dividir(a, b)
    default:
        return 0, fmt.Errorf("operação inválida: %s", operacao)
    }
}
```

### Dica para Exercício 3:

```go
func ObterEstacao(mes string) string {
    switch mes {
    case "dezembro", "janeiro", "fevereiro":
        return "Verão"
    case "março", "abril", "maio":
        return "Outono"
    // ... continue
    default:
        return "Mês inválido"
    }
}
```

### Dica para Exercício 4:

```go
func ValidarSenha(senha string) (bool, []string) {
    var erros []string

    if len(senha) < 8 {
        erros = append(erros, "Senha deve ter no mínimo 8 caracteres")
    }
    if len(senha) > 20 {
        erros = append(erros, "Senha deve ter no máximo 20 caracteres")
    }
    // ... continue verificando outras regras

    return len(erros) == 0, erros
}
```

**Nota:** Para verificar letras maiúsculas, minúsculas, números e caracteres especiais, você pode usar loops (que aprenderá depois) ou funções como `strings.ContainsAny()` e `unicode.IsUpper()`.

### Dica para Exercício 5:

```go
func CalcularDesconto(valor float64, tipoCliente string, temCupom bool) float64 {
    var descontoPercentual float64

    switch tipoCliente {
    case "VIP":
        descontoPercentual = 15
    case "Premium":
        descontoPercentual = 10
    case "Regular":
        descontoPercentual = 5
    default:
        descontoPercentual = 0
    }

    if temCupom {
        descontoPercentual += 5
    }

    // Limitar a 25%
    if descontoPercentual > 25 {
        descontoPercentual = 25
    }

    return valor * (descontoPercentual / 100)
}
```

### Dica para Exercício 6:

```go
func Processar(valor interface{}) string {
    switch v := valor.(type) {
    case int:
        return fmt.Sprintf("Inteiro: %d (dobro: %d)", v, v*2)
    case float64:
        return fmt.Sprintf("Decimal: %.2f (quadrado: %.2f)", v, v*v)
    case string:
        return fmt.Sprintf("Texto: %s (tamanho: %d)", v, len(v))
    case bool:
        return fmt.Sprintf("Booleano: %v (invertido: %v)", v, !v)
    case []int:
        soma := 0
        for _, num := range v {
            soma += num
        }
        return fmt.Sprintf("Slice: %v (soma: %d)", v, soma)
    default:
        return fmt.Sprintf("Tipo não suportado: %T", v)
    }
}
```

**Nota:** Para somar o slice, você pode usar um loop simples (que aprenderá depois) ou uma função auxiliar.

---

## Checklist de Revisão

Antes de entregar, verifique:

- [ ] Todos os códigos compilam sem erros
- [ ] Todos os casos de teste foram executados
- [ ] Tratamento de erros está implementado onde necessário
- [ ] Validações estão corretas
- [ ] Usei `switch` quando apropriado
- [ ] Usei `if-else` quando apropriado
- [ ] As reflexões são honestas e pensadas
- [ ] Incluí exemplos concretos nas reflexões
- [ ] O código está bem organizado e legível

---

Após completar os exercícios e reflexões, envie suas respostas para que eu possa analisar seu desempenho e fornecer feedback construtivo!

