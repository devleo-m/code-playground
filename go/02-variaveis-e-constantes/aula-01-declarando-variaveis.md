# Módulo 2: Variáveis e Constantes
## Aula 1: Declarando Variáveis com `var` e `:=`

Olá! Bem-vindo(a) ao nosso segundo módulo. Se os programas são como receitas, as variáveis são os ingredientes. Elas são os potes onde guardamos os valores que nosso programa vai usar, como números, textos, etc.

Em Go, temos duas formas principais de "criar" esses potes: a forma explícita, com a palavra-chave `var`, e um atalho super conveniente, o operador `:=`. Vamos entender os dois.

---

### A Forma Clássica: `var`

A palavra-chave `var` é a maneira mais formal e explícita de declarar uma variável. A estrutura é:

`var [nome da variável] [tipo]`

**Exemplo:**
```go
// Declarando uma variável chamada 'idade' que vai guardar um número inteiro (int).
var idade int

// Agora podemos colocar um valor nela.
idade = 30
```

Podemos também declarar e inicializar na mesma linha:

```go
// Declarando e já atribuindo o valor.
var nome string = "João"
```

#### A Mágica da Inferência de Tipo

Go é uma linguagem inteligente. Se você já atribui um valor ao declarar com `var`, ela consegue "adivinhar" o tipo, e você não precisa escrevê-lo. Isso se chama **inferência de tipo**.

```go
// Go sabe que "Maria" é um texto (string), então `nome` será do tipo string.
var nome = "Maria"

// Go sabe que 1.75 é um número com casas decimais, então `altura` será float64.
var altura = 1.75
```
Esta forma é mais limpa e muito comum.

---

### O Atalho dos Gophers: `:=` (Declaração Curta)

Dentro de funções, há uma maneira ainda mais rápida de declarar e inicializar variáveis: o operador `:=`. Ele faz duas coisas ao mesmo tempo: **declara** a variável e **atribui** um valor a ela, usando a inferência de tipo automaticamente.

**Exemplo:**
```go
// Dentro de uma função...
cidade := "São Paulo" // Declara a variável 'cidade' e a inicializa como string.
populacao := 12000000  // Declara 'populacao' e a inicializa como int.
```
Esta é a forma **preferida e mais comum** de declarar variáveis dentro de funções em Go, por ser rápida, concisa e legível.

---

### `var` vs `:=`: Quando usar qual?

Essa é a pergunta de ouro! A regra é bem simples:

1.  **Fora de funções (no nível do pacote):** Você **só pode** usar `var`. A declaração curta `:=` não é permitida aqui.
    ```go
    package main

    // Correto:
    var versao = "1.0"

    // Incorreto (causa erro de compilação):
    // data := "2025-01-01"

    func main() {
        // ...
    }
    ```

2.  **Dentro de funções:** Use `:=` sempre que puder, ou seja, quando você está declarando e inicializando uma variável pela primeira vez.
    ```go
    func main() {
        nome := "Ana" // Use := para a primeira vez.
        idade := 25

        // Para mudar o valor de uma variável que JÁ EXISTE, use apenas o '='.
        idade = 26 // Não use := de novo!
    }
    ```

3.  **Use `var` dentro de uma função apenas quando precisar declarar uma variável sem inicializá-la imediatamente.** Quando você faz isso, a variável recebe o "valor zero" do seu tipo, um tópico que veremos em detalhes numa próxima aula!
    ```go
    func main() {
        var erro error // Declara 'erro' mas não atribui valor ainda.
                       // 'erro' começa com o valor 'nil' (zero para erros).
        // ... algum código que pode gerar um erro ...
    }
    ```

---

E é isso por hoje! Pratique declarando variáveis das duas formas. Entender essa distinção é crucial para escrever um código Go limpo e idiomático.

Na próxima aula, vamos falar sobre algo que **não muda**: as **Constantes** e o misterioso `iota`.

Até lá!
