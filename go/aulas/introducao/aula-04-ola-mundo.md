# Aula 4: Seu Primeiro Programa: Olá, Mundo!

A tradição é clara: ao aprender uma nova linguagem de programação, o primeiro passo é fazer o computador dizer "Olá, Mundo!". E hoje, é exatamente o que faremos. Este simples programa nos ensinará muito sobre a estrutura básica de um arquivo Go.

## O Código

Vamos direto ao ponto. No diretório `ola-mundo` que criamos na aula anterior, crie um novo arquivo chamado `main.go` (ou, para esta aula, pode usar o arquivo `01-aula.go` que já existe) e coloque o seguinte conteúdo nele:

```go
package main

import "fmt"

func main() {
	fmt.Println("Olá, Mundo!")
}
```

## Desvendando o Código, Linha por Linha

Pode parecer pouco, mas cada linha tem um significado importante.

### `package main`

-   Todo arquivo Go começa com uma declaração de `package`.
-   Pacotes (`packages`) são a forma como o Go organiza e reutiliza código. Pense neles como bibliotecas ou módulos.
-   O pacote `main` é especial. Ele diz ao compilador do Go que este é um programa executável, e não uma biblioteca.

### `import "fmt"`

-   A palavra-chave `import` é usada para incluir pacotes externos em nosso arquivo.
-   Aqui, estamos importando o pacote `fmt` (uma abreviação de "format").
-   Este pacote, que faz parte da biblioteca padrão do Go, contém funções para entrada e saída de dados formatados, como imprimir texto no console.

### `func main() { ... }`

-   `func` é a palavra-chave para declarar uma nova função.
-   A função `main` é o coração do nosso programa executável. É o **ponto de entrada**. Quando executamos o programa, é a função `main` que é chamada primeiro.
-   O código que queremos executar fica dentro das chaves `{}`.

### `fmt.Println("Olá, Mundo!")`

-   Esta é a linha que faz a mágica acontecer.
-   `fmt.Println`: Estamos chamando a função `Println` (Print Line) que pertence ao pacote `fmt`.
-   A sintaxe `Pacote.Funcao()` é o padrão em Go.
-   Passamos o texto `"Olá, Mundo!"` como um argumento para a função, e ela se encarrega de imprimi-lo no terminal, adicionando uma nova linha no final.

## Executando o Programa

Agora, a parte divertida! Abra seu terminal e navegue até a pasta onde você salvou o arquivo (`ola-mundo`). Temos duas maneiras principais de executar o código:

### 1. Usando `go run`

Este comando compila e executa o programa de uma só vez. É perfeito para desenvolvimento rápido.

```bash
go run 01-aula.go
```
*(Se você nomeou o arquivo como `main.go`, use `go run main.go`)*

A saída no seu terminal será:
```
Olá, Mundo!
```

### 2. Usando `go build`

Este comando apenas compila o código, gerando um arquivo executável.

```bash
go build 01-aula.go
```
*(Ou `go build main.go`)*

Se você listar os arquivos na pasta agora (`ls` em Linux/macOS ou `dir` no Windows), verá um novo arquivo chamado `01-aula` (ou `ola-mundo` se você iniciou o módulo com esse nome, ou `main`). Este é o seu programa compilado!

Para executá-lo, basta chamar pelo nome:

**No Linux ou macOS:**
```bash
./01-aula
```

**No Windows:**
```bash
01-aula.exe
```

A saída será a mesma: `Olá, Mundo!`. A vantagem é que agora você tem um único arquivo que pode ser distribuído e executado em qualquer máquina com o mesmo sistema operacional, sem precisar do código-fonte ou do Go instalado.

---

Parabéns! Você escreveu, compilou e executou seu primeiro programa em Go. Você deu um passo fundamental para se tornar um Gopher!

Na nossa última aula introdutória, vamos explorar mais a fundo a ferramenta de linha de comando `go`, que é a sua principal aliada no dia a dia do desenvolvimento.

Vamos lá!
