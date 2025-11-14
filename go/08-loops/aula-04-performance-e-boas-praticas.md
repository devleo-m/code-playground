# Módulo 8: Loops em Go

## Aula 4 - Performance, Boas Práticas e O Que Deve/Não Deve Ser Utilizado

Agora que você entende loops, vamos falar sobre **como usá-los de forma eficiente e profissional**. Esta é a parte que separa desenvolvedores iniciantes de desenvolvedores experientes.

---

## 1. Performance: `for range` vs Loop Tradicional

### ⚡ `for range` é Geralmente Mais Eficiente

**✅ BOM: Use `for range` para Slices**

```go
// Mais eficiente e idiomático
for i, valor := range numeros {
    // ...
}
```

**⚠️ ACEITÁVEL: Loop Tradicional (quando necessário)**

```go
// Use apenas quando precisar de controle especial
for i := 0; i < len(numeros); i++ {
    // Modificar tamanho do slice, etc.
}
```

**Por quê `for range` é melhor?**

1. **Menos chamadas de função**: `len()` é otimizado, mas `range` evita chamadas repetidas
2. **Menos propenso a erros**: Não precisa se preocupar com índices fora de range
3. **Mais legível**: Código mais claro e idiomático
4. **Otimizações do compilador**: Go pode otimizar melhor `for range`

### Performance Real

**Benchmark típico (1000 elementos):**

```go
// for range: ~150ns
for i, v := range numeros {
    _ = v
}

// Loop tradicional: ~180ns
for i := 0; i < len(numeros); i++ {
    _ = numeros[i]
}
```

**Diferença**: ~20% mais rápido com `for range` (em média)

---

## 2. Performance: Pré-alocar Slices em Loops

### ❌ Ruim: Append em Loop sem Pré-alocação

```go
// RUIM: Múltiplas realocações
resultado := []int{}
for _, num := range numeros {
    if num > 0 {
        resultado = append(resultado, num)
    }
}
```

**Problema**: Cada `append` pode causar realocação se capacity estiver cheia.

### ✅ Bom: Pré-alocar Capacity

```go
// BOM: Uma única alocação
resultado := make([]int, 0, len(numeros)) // Pré-aloca capacity
for _, num := range numeros {
    if num > 0 {
        resultado = append(resultado, num)
    }
}
```

**Ganho**: 5-10x mais rápido para slices grandes!

### Quando Pré-alocar?

- ✅ Você sabe o tamanho aproximado
- ✅ Loop processa muitos elementos (> 100)
- ✅ Performance é crítica

- ❌ Tamanho completamente desconhecido
- ❌ Loops pequenos (< 10 elementos)
- ❌ Código de inicialização (roda uma vez)

---

## 3. Performance: Evitar Cálculos Dentro do Loop

### ❌ Ruim: Calcular `len()` a Cada Iteração

```go
// RUIM: len() é chamado a cada iteração
for i := 0; i < len(numeros); i++ {
    // ...
}
```

**Problema**: Se `numeros` é uma função ou expressão complexa, `len()` é recalculado.

### ✅ Bom: Calcular Uma Vez

```go
// BOM: len() calculado uma vez
tamanho := len(numeros)
for i := 0; i < tamanho; i++ {
    // ...
}

// OU melhor ainda: use for range
for i := range numeros {
    // ...
}
```

### ❌ Ruim: Expressões Complexas no Loop

```go
// RUIM: processar() chamado a cada iteração
for i := 0; i < len(processar(dados)); i++ {
    // ...
}
```

### ✅ Bom: Calcular Fora do Loop

```go
// BOM: calcular uma vez
resultado := processar(dados)
for i := 0; i < len(resultado); i++ {
    // ...
}
```

---

## 4. Performance: Break Cedo Quando Possível

### ✅ Bom: Usar `break` para Parar Cedo

```go
// BOM: Para assim que encontrar
func buscar(nomes []string, alvo string) int {
    for i, nome := range nomes {
        if nome == alvo {
            return i // Para imediatamente
        }
    }
    return -1
}
```

**Ganho**: Evita processar elementos desnecessários.

### ❌ Ruim: Continuar Processando Desnecessariamente

```go
// RUIM: Processa todos mesmo depois de encontrar
func buscar(nomes []string, alvo string) int {
    indice := -1
    for i, nome := range nomes {
        if nome == alvo {
            indice = i
            // Continua processando mesmo depois de encontrar!
        }
    }
    return indice
}
```

**Regra**: Use `break` ou `return` assim que encontrar o que procura.

---

## 5. Performance: Evitar Alocações Dentro de Loops Quentes

### ❌ Ruim: Criar Slices Novos a Cada Iteração

```go
// RUIM: Nova alocação a cada iteração
for i := 0; i < 1000; i++ {
    resultado := []int{} // Nova alocação!
    // ...
}
```

### ✅ Bom: Reutilizar ou Pré-alocar

```go
// BOM: Reutilizar slice
resultado := make([]int, 0, 100)
for i := 0; i < 1000; i++ {
    resultado = resultado[:0] // Reseta length, mantém capacity
    // ...
}
```

### ❌ Ruim: Concatenar Strings em Loop

```go
// RUIM: Nova string a cada concatenação
resultado := ""
for _, palavra := range palavras {
    resultado += palavra // Cria nova string!
}
```

### ✅ Bom: Usar `strings.Builder`

```go
// BOM: Eficiente para construir strings
var builder strings.Builder
builder.Grow(len(palavras) * 10) // Pré-aloca
for _, palavra := range palavras {
    builder.WriteString(palavra)
}
resultado := builder.String()
```

**Ganho**: 100-1000x mais rápido para muitas concatenações!

---

## 6. Boas Práticas: Preferir `for range` para Coleções

### ✅ Sempre Use `for range` Quando Possível

**Para Slices:**

```go
// ✅ BOM
for i, valor := range numeros {
    // ...
}

// ⚠️ EVITE (a menos que necessário)
for i := 0; i < len(numeros); i++ {
    // ...
}
```

**Para Maps:**

```go
// ✅ BOM
for chave, valor := range mapa {
    // ...
}
```

**Para Strings:**

```go
// ✅ BOM (retorna runes)
for i, rune := range texto {
    // ...
}

// ❌ ERRADO (retorna bytes)
for i := 0; i < len(texto); i++ {
    byte := texto[i] // Byte, não rune!
}
```

---

## 7. Boas Práticas: Usar `break` e `continue` Apropriadamente

### ✅ Use `break` para Buscar

```go
// ✅ BOM: Para quando encontrar
func buscar(nomes []string, alvo string) bool {
    for _, nome := range nomes {
        if nome == alvo {
            return true // Para imediatamente
        }
    }
    return false
}
```

### ✅ Use `continue` para Filtrar

```go
// ✅ BOM: Pula elementos indesejados
func processarPositivos(numeros []int) {
    for _, num := range numeros {
        if num <= 0 {
            continue // Pula negativos e zero
        }
        processar(num)
    }
}
```

### ❌ Evite Aninhamento Desnecessário

```go
// ❌ RUIM: Aninhamento desnecessário
for _, num := range numeros {
    if num > 0 {
        if num < 100 {
            processar(num)
        }
    }
}

// ✅ BOM: Use continue
for _, num := range numeros {
    if num <= 0 || num >= 100 {
        continue
    }
    processar(num)
}
```

---

## 8. Boas Práticas: Cuidado com Modificações Durante Iteração

### ✅ Seguro: Modificar Elementos Existentes

```go
// ✅ SEGURO: Modificar valores
for i := range numeros {
    numeros[i] *= 2
}
```

### ❌ Perigoso: Adicionar/Remover Durante `range`

```go
// ❌ PERIGOSO: Comportamento indefinido
for _, num := range numeros {
    if num > 0 {
        numeros = append(numeros, num*2) // Pode causar problemas!
    }
}
```

### ✅ Solução: Usar Loop Tradicional ou Criar Novo Slice

```go
// ✅ BOM: Loop tradicional para modificar tamanho
for i := 0; i < len(numeros); i++ {
    if numeros[i] < 0 {
        numeros = append(numeros[:i], numeros[i+1:]...)
        i-- // Ajustar índice
    }
}

// ✅ BOM: Criar novo slice
resultado := []int{}
for _, num := range numeros {
    if num > 0 {
        resultado = append(resultado, num)
    }
}
```

### ⚠️ Maps: Deletar é Seguro, Adicionar Não

```go
// ✅ SEGURO: Deletar durante iteração
for chave, valor := range mapa {
    if valor == 0 {
        delete(mapa, chave) // Seguro!
    }
}

// ❌ PERIGOSO: Adicionar/modificar durante iteração
for chave, valor := range mapa {
    mapa[chave] = valor * 2 // Comportamento indefinido!
}
```

---

## 9. Boas Práticas: Strings e Unicode

### ❌ ERRADO: Indexação Direta em Strings

```go
// ❌ ERRADO: Vê bytes, não caracteres
texto := "Café"
for i := 0; i < len(texto); i++ {
    fmt.Printf("%c", texto[i]) // Pode quebrar caracteres multibyte!
}
```

### ✅ CORRETO: `for range` para Strings

```go
// ✅ CORRETO: Vê runes (caracteres)
texto := "Café"
for i, rune := range texto {
    fmt.Printf("Posição %d: %c\n", i, rune)
}
```

### ✅ Para Acesso Aleatório: Converter para `[]rune`

```go
// ✅ CORRETO: Converter para acesso por índice
texto := "Olá, 世界!"
runes := []rune(texto)
primeiro := runes[0]  // 'O'
ultimo := runes[len(runes)-1]  // '!'
```

---

## 10. Boas Práticas: Evitar Labels Quando Possível

### ❌ Evite Labels em Código Simples

```go
// ❌ EVITE: Labels desnecessários
LoopExterno:
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            if j == 1 {
                break LoopExterno
            }
        }
    }
```

### ✅ Prefira Refatorar em Função

```go
// ✅ BOM: Extrair em função
func processar() bool {
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            if j == 1 {
                return true // Mais claro que break com label
            }
        }
    }
    return false
}
```

### ✅ Use Labels Apenas Quando Necessário

```go
// ✅ ACEITÁVEL: Labels para casos complexos
BuscarEmMatriz:
    for i := range matriz {
        for j := range matriz[i] {
            if matriz[i][j] == alvo {
                encontrado = true
                break BuscarEmMatriz // Labels úteis aqui
            }
        }
    }
```

---

## 11. Boas Práticas: Nomes de Variáveis Descritivos

### ❌ Ruim: Nomes Genéricos

```go
// ❌ RUIM: Nomes não descritivos
for i := 0; i < len(u); i++ {
    for j := 0; j < len(u[i]); j++ {
        // O que são 'u', 'i', 'j'?
    }
}
```

### ✅ Bom: Nomes Descritivos

```go
// ✅ BOM: Nomes claros
for linha := 0; linha < len(matriz); linha++ {
    for coluna := 0; coluna < len(matriz[linha]); coluna++ {
        // Muito mais claro!
    }
}

// ✅ BOM: for range com nomes descritivos
for indice, produto := range produtos {
    // ...
}
```

---

## 12. O Que NÃO Deve Ser Feito

### ❌ Erro 1: Modificar Slice Durante `range`

```go
// ❌ ERRADO
for _, item := range items {
    if condicao {
        items = append(items, novoItem) // Comportamento indefinido!
    }
}
```

### ❌ Erro 2: Usar `goto` Desnecessariamente

```go
// ❌ EVITE
Inicio:
    // código
    if condicao {
        goto Inicio // Use break/continue/funções ao invés
    }
```

### ❌ Erro 3: Loop Infinito sem Saída

```go
// ❌ PERIGOSO: Pode travar o programa
for {
    // Sem break ou return!
}
```

### ❌ Erro 4: Indexação Direta em Strings Unicode

```go
// ❌ ERRADO: Quebra caracteres multibyte
texto := "世界"
for i := 0; i < len(texto); i++ {
    fmt.Printf("%c", texto[i]) // Não funciona corretamente!
}
```

### ❌ Erro 5: Não Usar `break` Quando Encontra

```go
// ❌ INEFICIENTE: Continua processando desnecessariamente
for _, item := range items {
    if item == alvo {
        encontrado = true
        // Deveria ter break aqui!
    }
}
```

---

## 13. Padrões Idiomáticos em Go

### Padrão 1: Filtrar com `continue`

```go
func filtrarPares(numeros []int) []int {
    resultado := make([]int, 0, len(numeros)/2)
    for _, num := range numeros {
        if num%2 != 0 {
            continue // Pula ímpares
        }
        resultado = append(resultado, num)
    }
    return resultado
}
```

### Padrão 2: Buscar com `break`

```go
func buscar(nomes []string, alvo string) (int, bool) {
    for i, nome := range nomes {
        if nome == alvo {
            return i, true // Para imediatamente
        }
    }
    return -1, false
}
```

### Padrão 3: Processar até Condição

```go
func processarAteNegativo(numeros []int) int {
    soma := 0
    for _, num := range numeros {
        if num < 0 {
            break // Para no primeiro negativo
        }
        soma += num
    }
    return soma
}
```

### Padrão 4: Acumular Valores

```go
func somar(numeros []int) int {
    soma := 0
    for _, num := range numeros {
        soma += num
    }
    return soma
}
```

---

## 14. Checklist de Boas Práticas

Antes de considerar seu código "pronto", pergunte-se:

- [ ] Estou usando `for range` para coleções?
- [ ] Pré-aloquei slices quando sei o tamanho aproximado?
- [ ] Usei `break` para parar quando encontrei o que procurava?
- [ ] Usei `continue` para filtrar elementos?
- [ ] Evitei modificar tamanho de coleções durante `range`?
- [ ] Usei `for range` para strings (não indexação direta)?
- [ ] Evitei `goto`?
- [ ] Nomes de variáveis são descritivos?
- [ ] Evitei cálculos desnecessários dentro do loop?
- [ ] Loop infinito tem forma de sair?

---

## 15. Resumo: Regras de Ouro

1. **Prefira `for range`**: Mais idiomático, eficiente e seguro
2. **Pré-aloque quando possível**: Evita realocações custosas
3. **Break cedo**: Pare assim que encontrar o que procura
4. **Continue para filtrar**: Pule elementos indesejados
5. **Cuidado com modificações**: Não adicione/remova durante `range`
6. **Strings com `for range`**: Sempre use para Unicode correto
7. **Evite `goto`**: Use funções, `break`, `continue`
8. **Nomes descritivos**: Código mais legível
9. **Calcule fora do loop**: Evite recálculos desnecessários
10. **Teste casos extremos**: Slices vazios, valores nulos, etc.

---

## Conclusão

Loops são fundamentais, mas usá-los corretamente vai além da sintaxe:

- **Performance**: Pré-alocar, calcular fora, usar `for range`
- **Segurança**: Cuidado com modificações durante iteração
- **Legibilidade**: Nomes descritivos, evitar labels desnecessários
- **Idiomático**: Preferir `for range`, usar `break`/`continue` apropriadamente

Lembre-se: **Código é lido muito mais vezes do que escrito**. Priorize clareza, e otimize apenas quando necessário e medido.

---

Na próxima etapa, você completará os exercícios e eu analisarei seu desempenho. Boa sorte!
