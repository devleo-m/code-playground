# Módulo 2: Variáveis e Constantes
## Aula 2: Valores Imutáveis com `const` e Sequências com `iota`

Olá! Na aula anterior, aprendemos sobre variáveis, que são como caixas com etiquetas onde o conteúdo pode mudar. Hoje, vamos aprender sobre suas primas mais rígidas: as **constantes**.

Constantes, como o nome sugere, guardam valores que **não podem ser alterados** depois de definidos. Elas são perfeitas para valores que devem ser fixos em todo o programa, como o valor de Pi, uma porta de servidor padrão ou uma chave de API.

---

### Declarando Constantes: `const`

A sintaxe é muito parecida com a do `var`, mas usamos a palavra-chave `const`.

```go
const Pi = 3.14159
const PortaServidor = 8080
const MensagemBoasVindas = "Bem-vindo ao sistema!"
```

**Regras importantes para constantes:**
1.  **Imutáveis:** Uma vez declarada, você não pode tentar atribuir um novo valor a uma constante. O código `Pi = 3.14` causaria um erro de compilação.
2.  **Definidas em tempo de compilação:** O valor de uma constante deve ser conhecido pelo compilador antes mesmo do programa rodar. Isso significa que você não pode atribuir o resultado de uma função a uma constante.
    ```go
    // Correto:
    const a = 10 + 5 // O compilador consegue calcular isso.

    // Incorreto (causa erro):
    // const b = MinhaFuncao() // O resultado da função só é conhecido em tempo de execução.
    ```
3.  **Podem ser "tipadas" ou "não tipadas":** Assim como com `var`, o tipo pode ser explícito ou inferido. Constantes "não tipadas" são mais flexíveis, pois podem ser usadas em contextos com diferentes tipos numéricos sem a necessidade de conversão explícita.
    ```go
    const MeuNumero = 42 // Constante não tipada (tipo inferido) - mais flexível
    const MeuNumeroTipado int = 42 // Constante tipada - mais rígida
    ```

---

### A Joia Escondida: `iota`

Go tem uma ferramenta especial para criar sequências de constantes numéricas: o `iota`. Ele é um identificador que o Go usa dentro de blocos `const` para simplificar a criação de números em sequência.

**Como funciona:**
-   `iota` começa em `0` no início de um bloco `const`.
-   A cada nova declaração de constante no mesmo bloco, `iota` é incrementado em `1`.

**Exemplo Simples:**
```go
const (
    Domingo = iota // Domingo = 0
    Segunda        // Segunda = 1 (iota é incrementado)
    Terca          // Terca   = 2
    Quarta         // Quarta  = 3
    Quinta         // Quinta  = 4
    Sexta          // Sexta   = 5
    Sabado         // Sabado  = 6
)
```
Isso é muito mais limpo e menos propenso a erros do que atribuir `0, 1, 2, 3...` manualmente. É a forma padrão em Go para criar enumerações (`enums`).

#### `iota` com Expressões

A mágica do `iota` não para por aí. Você pode usá-lo em expressões para criar padrões mais complexos.

**Exemplo: Potências de 2 (flags de bits)**
```go
const (
    // A expressão (1 << iota) é repetida para as constantes seguintes
    FlagLeitura   = 1 << iota // 1 << 0  (binário 01) -> 1
    FlagEscrita   = 1 << iota // 1 << 1  (binário 10) -> 2
    FlagExecucao  = 1 << iota // 1 << 2  (binário 100) -> 4
)
```

**Exemplo: Pular um valor**
Podemos usar o "identificador branco" (`_`) para pular um valor do `iota`.
```go
const (
    Valido = iota // 0
    _             // 1 (iota foi incrementado, mas o valor foi descartado)
    Invalido      // 2
)
```

---

Constantes e `iota` são ferramentas que tornam seu código mais seguro, legível e robusto. Ao usar constantes para valores que não devem mudar, você evita bugs acidentais e deixa a intenção do seu código mais clara para outros desenvolvedores.

Na próxima aula, vamos explorar o que acontece quando você declara uma variável, mas não dá um valor a ela: os **Valores Zero**.

Até a próxima!
