# Módulo 2: Variáveis e Constantes
## Aula 4: Entendendo Escopo e Sombreamento (Shadowing)

Parabéns por chegar à nossa última aula deste módulo! Agora que dominamos como criar variáveis e constantes, precisamos entender **onde** elas podem ser acessadas. Esse conceito de "onde" é chamado de **escopo**.

---

### O que é Escopo?

Escopo é a região do seu código onde uma variável declarada é visível e pode ser utilizada. Se você tentar acessar uma variável fora do seu escopo, o compilador do Go vai reclamar.

Em Go, o escopo é determinado por blocos de código, que são delimitados por chaves `{}`.

#### Níveis de Escopo

1.  **Escopo de Universo (Universe Block):** É o escopo mais externo. Contém todas as funções e tipos pré-declarados da linguagem, como `int`, `string`, `true`, `false`, `make`, `len`, etc. Você tem acesso a eles em qualquer lugar do seu código.

2.  **Escopo de Pacote (Package Scope):** Qualquer variável, constante ou função declarada fora de uma função (no nível mais alto do seu arquivo) é visível para **todos os arquivos** dentro do mesmo pacote.
    ```go
    package main

    var versao = 1.1 // 'versao' tem escopo de pacote.

    func main() {
        imprimirVersao() // Pode ser acessada aqui.
    }

    // em outro arquivo no mesmo pacote 'main'...
    // func imprimirVersao() {
    //     fmt.Println(versao) // E aqui também!
    // }
    ```

3.  **Escopo de Função (Function Scope):** Parâmetros e variáveis declarados diretamente dentro de uma função são visíveis apenas dentro daquela função.

4.  **Escopo de Bloco (Block Scope):** Variáveis declaradas dentro de blocos como `if`, `for`, ou `switch` são visíveis apenas dentro daquele bloco (incluindo sub-blocos).
    ```go
    func main() {
        for i := 0; i < 5; i++ {
            // 'i' e 'mensagem' só existem dentro deste loop 'for'.
            mensagem := "Estou no loop"
            fmt.Println(i, mensagem)
        }
        // fmt.Println(i) // Erro de compilação: 'i' é indefinido aqui.
    }
    ```

---

### O Perigo do Sombreamento (Shadowing)

Sombreamento acontece quando você declara uma variável em um escopo interno com o **mesmo nome** de uma variável que já existe em um escopo mais externo. A variável interna "faz sombra" sobre a externa, tornando a externa inacessível dentro daquele bloco.

**Exemplo Clássico:**

```go
package main

import "fmt"

var numero = 100 // Escopo de pacote

func main() {
    fmt.Printf("Número (antes do if): %d\n", numero) // Imprime 100

    condicao := true
    if condicao {
        // Aqui, estamos criando uma NOVA variável 'numero'.
        // Ela "sombra" a variável 'numero' do pacote.
        numero := 200
        fmt.Printf("Número (dentro do if): %d\n", numero) // Imprime 200
    }

    fmt.Printf("Número (depois do if): %d\n", numero) // Imprime 100 de novo!
}
```
No exemplo acima, a variável `numero` de dentro do `if` é completamente diferente da variável `numero` de fora. A de fora não foi alterada.

#### Por que o Sombreamento pode ser perigoso?

O sombreamento é uma fonte comum de bugs sutis. Você pode acidentalmente criar uma nova variável quando sua intenção era apenas **modificar** a variável externa.

O operador `:=` é um culpado frequente. Se você o usa em uma variável que já foi declarada num escopo externo, você pode estar sombreando-a sem querer.

**Como evitar:**
-   **Atenção ao usar `:=`:** Tenha certeza de que não está redeclarando uma variável que você pretendia reutilizar. Para modificar uma variável existente, use sempre `=`.
-   **Nomes Descritivos:** Evite usar nomes de variáveis genéricos (`n`, `v`, `tmp`) que têm mais chance de colidir.
-   **Linters:** Ferramentas de análise de código (linters) como o `go vet` podem ser configuradas para avisar sobre variáveis sombreadas.

---

### Conclusão do Módulo

E com isso, você desvendou os segredos de como o Go gerencia dados!

**Recapitulando o Módulo 2:**
-   **Aula 1:** Aprendemos a declarar variáveis com `var` e `:=`.
-   **Aula 2:** Entendemos o poder dos valores imutáveis com `const` e `iota`.
-   **Aula 3:** Vimos a segurança dos `valores zero`.
-   **Aula 4:** Compreendemos onde as variáveis vivem com `escopo` e o cuidado a se ter com `sombreamento`.

Você agora possui o conhecimento fundamental para estruturar a lógica e os dados dos seus programas em Go.

Continue praticando esses conceitos. O próximo passo em nossa jornada será explorar os **tipos básicos** do Go em mais detalhes.

Parabéns pelo excelente progresso!
