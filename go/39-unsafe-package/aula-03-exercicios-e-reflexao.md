# Módulo 39: Unsafe Package em Go
## Aula 3 - Exercícios e Reflexão

⚠️ **AVISO**: Os exercícios abaixo são para aprendizado. Em produção, sempre prefira código seguro!

---

## Exercícios Práticos

### Exercício 1: Entender Sizeof, Offsetof, Alignof

Crie um programa que analisa uma struct e mostra:
1. Tamanho total da struct
2. Offset de cada campo
3. Alinhamento de cada campo
4. Padding (espaço desperdiçado)

**Código base:**
```go
package main

import (
    "fmt"
    "unsafe"
)

type Example struct {
    A bool
    B int64
    C bool
    D int32
}

func analyzeStruct(x interface{}) {
    // TODO: Analisar struct usando unsafe
}

func main() {
    analyzeStruct(Example{})
}
```

**Tarefa**: Complete a função para mostrar todas as informações solicitadas.

---

### Exercício 2: Pointer Arithmetic Controlado

Crie uma função que acessa elementos de um array usando pointer arithmetic, mas com **validação completa**.

**Requisitos:**
- Função `getElement(arr []int, index int) (int, error)`
- Deve validar índice antes de acessar
- Deve usar unsafe para acesso
- Deve retornar erro se índice inválido

**Código base:**
```go
package main

import (
    "fmt"
    "unsafe"
)

func getElement(arr []int, index int) (int, error) {
    // TODO: Implementar com validação completa
    return 0, nil
}

func main() {
    arr := []int{10, 20, 30, 40, 50}
    
    // Testar casos válidos
    val, err := getElement(arr, 2)
    if err != nil {
        fmt.Println("Erro:", err)
    } else {
        fmt.Printf("arr[2] = %d\n", val)
    }
    
    // Testar caso inválido
    _, err = getElement(arr, 10)
    if err != nil {
        fmt.Println("Erro esperado:", err)
    }
}
```

**Tarefa**: Implemente com validação completa e tratamento de erros.

---

### Exercício 3: Conversão Segura de Tipos

Crie funções que convertem entre tipos usando unsafe, mas com **validação de tamanho e alinhamento**.

**Requisitos:**
- `bytesToInt(b []byte) (int, error)` - Converter []byte para int
- Validar tamanho do buffer
- Validar alinhamento
- Retornar erro se inválido

**Código base:**
```go
package main

import (
    "fmt"
    "unsafe"
)

func bytesToInt(b []byte) (int, error) {
    // TODO: Implementar com validação
    return 0, nil
}

func main() {
    // Caso válido
    bytes := []byte{42, 0, 0, 0, 0, 0, 0, 0}
    val, err := bytesToInt(bytes)
    if err != nil {
        fmt.Println("Erro:", err)
    } else {
        fmt.Printf("Valor: %d\n", val)
    }
    
    // Caso inválido (buffer pequeno)
    smallBytes := []byte{42}
    _, err = bytesToInt(smallBytes)
    if err != nil {
        fmt.Println("Erro esperado:", err)
    }
}
```

**Tarefa**: Implemente com todas as validações necessárias.

---

## Perguntas de Reflexão

### Reflexão 1: Por Que Unsafe Existe?

O Go é uma linguagem focada em segurança, mas ainda fornece o package unsafe.

**Perguntas para refletir**:
1. **Por que uma linguagem segura** permite código inseguro? Quais são os trade-offs?
2. Em que situações o **custo de segurança** é maior que o benefício? Dê exemplos práticos.
3. Como o uso de unsafe afeta a **portabilidade** do código? O que acontece em diferentes arquiteturas?
4. Se você estivesse projetando uma linguagem, **incluiria** algo como unsafe? Por quê?

**Escreva suas reflexões** (mínimo 250 palavras):

---

### Reflexão 2: Quando Unsafe É Justificável?

O uso de unsafe raramente é necessário, mas às vezes é justificável.

**Perguntas para refletir**:
1. **Quais critérios objetivos** você usaria para decidir se unsafe é justificável? Liste pelo menos 5 critérios.
2. Como você **mediria** se unsafe realmente melhora performance o suficiente para justificar o risco?
3. Em um projeto de equipe, **quem deveria decidir** se unsafe pode ser usado? Que processo você estabeleceria?
4. Se você encontrasse código com unsafe em um projeto, **como você avaliaria** se é justificável ou não?

**Escreva suas reflexões** (mínimo 250 palavras):

---

### Reflexão 3: Segurança vs Performance

Unsafe oferece performance, mas sacrifica segurança.

**Perguntas para refletir**:
1. **Como você balancearia** segurança vs performance em diferentes contextos? (ex: sistema crítico, aplicação web, ferramenta interna)
2. Quais são os **custos reais** de usar unsafe? Pense em: manutenção, debugging, testes, riscos de segurança.
3. Em que situações você **aceitaria** o risco de unsafe? Em que situações você **nunca** aceitaria?
4. Como você **mitigaria** os riscos de unsafe se tivesse que usá-lo? Que práticas adotaria?

**Escreva suas reflexões** (mínimo 250 palavras):

---

## Checklist de Aprendizado

Marque conforme você completa:

- [ ] Entendi o que é unsafe package
- [ ] Sei usar unsafe.Sizeof, Offsetof, Alignof
- [ ] Entendo pointer arithmetic
- [ ] Sei fazer conversões de tipos (com cuidado)
- [ ] Entendo os riscos de unsafe
- [ ] Sei quando unsafe é justificável
- [ ] Sei validar antes de usar unsafe
- [ ] Entendo alternativas ao unsafe
- [ ] Sei documentar código com unsafe
- [ ] Entendo quando NÃO usar unsafe

---

## Desafio Extra (Opcional)

### Desafio: Serializador Zero-Copy

Crie um serializador que converte structs para []byte sem cópia, usando unsafe.

**Requisitos:**
- Função `serialize(x interface{}) ([]byte, error)`
- Função `deserialize(b []byte, x interface{}) error`
- Validação completa de tamanho e alinhamento
- Documentação extensiva dos riscos
- Testes em diferentes arquiteturas

**⚠️ AVISO**: Este é um exercício avançado. Certifique-se de entender completamente antes de implementar.

---

**Lembre-se**: Unsafe é uma ferramenta poderosa e perigosa. Use apenas quando absolutamente necessário! ⚠️

