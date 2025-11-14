# Módulo 8: Loops em Go

## Aula 2 - Simplificada: Entendendo Loops

Agora vamos entender loops de forma mais simples, usando analogias do nosso dia a dia!

---

## 1. Loop: Repetir uma Tarefa

### Analogia: Contar de 1 até 10

Pense em um loop como **repetir uma tarefa** várias vezes:

**Mundo real:**

```
Você precisa contar de 1 até 10:
"1, 2, 3, 4, 5, 6, 7, 8, 9, 10"
```

**Loop em Go:**

```go
for i := 1; i <= 10; i++ {
    fmt.Printf("%d ", i)
}
```

**Por que funciona:**

- Você repete a mesma ação (contar)
- Para em um número específico (10)
- Incrementa de 1 em 1

---

## 2. for Clássico: A Receita com Passos

### Analogia: Receita de Bolo com Passos

Um `for` clássico é como seguir uma **receita com passos numerados**:

**Receita:**

```
Passo 1: Misturar farinha
Passo 2: Adicionar açúcar
Passo 3: Adicionar ovos
Passo 4: Misturar tudo
Passo 5: Assar
```

**Loop em Go:**

```go
for passo := 1; passo <= 5; passo++ {
    fmt.Printf("Passo %d: Executar ação\n", passo)
}
```

**Os três componentes:**

1. **Inicialização** (`passo := 1`): Começar no passo 1
2. **Condição** (`passo <= 5`): Continuar até o passo 5
3. **Pós-instrução** (`passo++`): Ir para o próximo passo

**Analogia:**
É como ter uma **lista de tarefas numeradas** que você segue em ordem!

---

## 3. for While-Style: Continuar Até Algo Acontecer

### Analogia: Esperar o Ônibus

Um `for` while-style é como **esperar até algo acontecer**:

**Mundo real:**

```
Você está na parada de ônibus:
- Enquanto o ônibus não chegar, continue esperando
- Quando o ônibus chegar, pare de esperar
```

**Loop em Go:**

```go
onibusChegou := false
for !onibusChegou {
    fmt.Println("Ainda esperando...")
    // ... alguma lógica que pode fazer onibusChegou = true
}
```

**Analogia:**
É como **esperar** até uma condição ser verdadeira. Você não sabe quanto tempo vai demorar, mas continua até acontecer!

---

## 4. for Infinito: Trabalho que Nunca Para

### Analogia: Atendente de Telefone

Um loop infinito é como um **atendente que fica sempre disponível**:

**Mundo real:**

```
Atendente de telemarketing:
- Fica sempre esperando ligações
- Atende quando recebe uma ligação
- Continua esperando para a próxima
- Só para quando alguém desliga o sistema
```

**Loop em Go:**

```go
for {
    // Esperar ligação
    ligacao := receberLigacao()
    if ligacao == "desligar" {
        break // Desligar o sistema
    }
    atender(ligacao)
}
```

**Analogia:**
É como um **serviço que roda 24 horas** - sempre disponível até ser desligado manualmente!

---

## 5. for range: Ver Cada Item de uma Lista

### Analogia: Verificar Itens da Lista de Compras

`for range` é como **verificar cada item** da sua lista de compras:

**Lista de compras:**

```
📝 Lista de Compras:
1. Arroz
2. Feijão
3. Açúcar
4. Café
```

**Loop em Go:**

```go
listaCompras := []string{"Arroz", "Feijão", "Açúcar", "Café"}

for indice, item := range listaCompras {
    fmt.Printf("%d. %s\n", indice+1, item)
}
```

**Analogia:**
É como **passar os olhos** por cada item da lista, um por um, até ver todos!

---

## 6. for range com Slices: Ver Cada Pessoa na Fila

### Analogia: Atender Pessoas na Fila

Iterar sobre um slice é como **atender cada pessoa** em uma fila:

**Fila do banco:**

```
Fila: [João] [Maria] [Pedro] [Ana]
       ↑
    Você atende uma por vez
```

**Loop em Go:**

```go
fila := []string{"João", "Maria", "Pedro", "Ana"}

for posicao, pessoa := range fila {
    fmt.Printf("Atendendo pessoa %d: %s\n", posicao+1, pessoa)
}
```

**Ignorar posição (só ver nomes):**

```go
for _, pessoa := range fila {
    fmt.Printf("Atendendo: %s\n", pessoa)
}
```

**Analogia:**
É como um **caixa de banco** que atende cada pessoa da fila, uma de cada vez!

---

## 7. for range com Maps: Ver Cada Item do Dicionário

### Analogia: Ler um Dicionário

Iterar sobre um map é como **ler um dicionário** palavra por palavra:

**Dicionário:**

```
Palavra: "casa" → Significado: "lugar onde se mora"
Palavra: "carro" → Significado: "veículo automotor"
Palavra: "livro" → Significado: "coleção de páginas"
```

**Loop em Go:**

```go
dicionario := map[string]string{
    "casa":  "lugar onde se mora",
    "carro": "veículo automotor",
    "livro": "coleção de páginas",
}

for palavra, significado := range dicionario {
    fmt.Printf("%s: %s\n", palavra, significado)
}
```

**⚠️ Importante: Ordem Aleatória!**

O dicionário pode ser lido em **qualquer ordem** (é aleatório!). É como embaralhar as páginas do dicionário antes de ler!

**Analogia:**
É como pegar um **dicionário embaralhado** - você lê todas as palavras, mas não sabe qual vem primeiro!

---

## 8. for range com Strings: Ler Cada Letra de uma Palavra

### Analogia: Soletrar uma Palavra

Iterar sobre uma string é como **soletrar** uma palavra, letra por letra:

**Soletrar "Café":**

```
C - a - f - é
```

**Loop em Go:**

```go
palavra := "Café"

for posicao, letra := range palavra {
    fmt.Printf("Posição %d: letra %c\n", posicao, letra)
}
```

**⚠️ CUIDADO: Letras vs Bytes!**

**ERRADO** (indexação direta - vê bytes, não letras):

```go
palavra := "Café"
for i := 0; i < len(palavra); i++ {
    fmt.Printf("Byte %d: %d\n", i, palavra[i])
}
// Isso vê BYTES, não letras! "é" tem 2 bytes!
```

**CORRETO** (for range - vê letras/runes):

```go
for i, letra := range palavra {
    fmt.Printf("Letra %d: %c\n", i, letra)
}
// Isso vê LETRAS completas!
```

**Analogia:**

- **Indexação direta**: Como olhar os **pixels** de uma foto (não vê a imagem completa)
- **for range**: Como olhar a **foto inteira** (vê cada elemento completo)

---

## 9. break: Parar Quando Encontrar

### Analogia: Procurar Chaves Perdidas

`break` é como **parar de procurar** quando você encontra o que procura:

**Mundo real:**

```
Você perdeu suas chaves e está procurando:
- Procurar no quarto... não está
- Procurar na sala... não está
- Procurar na cozinha... ENCONTROU! ✅
- Para de procurar (não precisa mais verificar outros lugares)
```

**Loop em Go:**

```go
lugares := []string{"quarto", "sala", "cozinha", "banheiro"}

for _, lugar := range lugares {
    if procurarChaves(lugar) {
        fmt.Printf("Encontrei as chaves no %s!\n", lugar)
        break // Para de procurar!
    }
}
```

**Analogia:**
É como **encontrar o que procura** e parar imediatamente - não precisa continuar procurando!

---

## 10. continue: Pular Algo e Continuar

### Analogia: Separar Frutas Boas das Ruins

`continue` é como **pular frutas ruins** e continuar verificando as outras:

**Mundo real:**

```
Você está separando frutas:
- Maçã 1: Boa ✅ (coloca na caixa)
- Maçã 2: Ruim ❌ (pula, não coloca)
- Maçã 3: Boa ✅ (coloca na caixa)
- Maçã 4: Ruim ❌ (pula, não coloca)
- Maçã 5: Boa ✅ (coloca na caixa)
```

**Loop em Go:**

```go
frutas := []string{"maçã boa", "maçã ruim", "maçã boa", "maçã ruim", "maçã boa"}

for _, fruta := range frutas {
    if fruta == "maçã ruim" {
        continue // Pula esta fruta
    }
    fmt.Printf("Colocando na caixa: %s\n", fruta)
}
```

**Analogia:**
É como **ignorar** algo que não quer e continuar com o resto!

---

## 11. Loops Aninhados: Tabela de Multiplicação

### Analogia: Tabela de Multiplicação

Loops aninhados são como fazer uma **tabela de multiplicação**:

**Tabela:**

```
1 x 1 = 1
1 x 2 = 2
1 x 3 = 3
2 x 1 = 2
2 x 2 = 4
2 x 3 = 6
```

**Loop em Go:**

```go
for i := 1; i <= 2; i++ {
    for j := 1; j <= 3; j++ {
        fmt.Printf("%d x %d = %d\n", i, j, i*j)
    }
}
```

**Analogia:**
É como ter uma **grade** onde você percorre cada linha e, dentro de cada linha, percorre cada coluna!

**Visualização:**

```
Linha 1: [1x1] [1x2] [1x3]
Linha 2: [2x1] [2x2] [2x3]
```

---

## 12. break com Label: Sair de Múltiplos Loops

### Analogia: Procurar em Múltiplos Lugares

`break` com label é como **parar de procurar em TODOS os lugares** quando encontrar:

**Mundo real:**

```
Você está procurando suas chaves:
- Casa 1:
  - Quarto: não está
  - Sala: não está
  - Cozinha: ENCONTROU! ✅
  - Para de procurar em TODA a casa (não verifica banheiro, etc.)
```

**Loop em Go:**

```go
Casa1:
    for _, comodo := range []string{"quarto", "sala", "cozinha", "banheiro"} {
        if procurarChaves(comodo) {
            fmt.Println("Encontrei! Parando busca em toda a casa!")
            break Casa1 // Sai de TODOS os loops aninhados
        }
    }
```

**Analogia:**
É como **encontrar o que procura** e parar completamente - não precisa mais verificar nada!

---

## 13. continue com Label: Pular para Próxima Iteração Externa

### Analogia: Pular um Dia Inteiro

`continue` com label é como **pular um dia inteiro** e ir para o próximo:

**Mundo real:**

```
Semana de trabalho:
- Segunda: trabalhar normalmente
- Terça: FERIADO! Pula o dia inteiro
- Quarta: trabalhar normalmente
```

**Loop em Go:**

```go
Semana:
    for dia := 0; dia < 5; dia++ {
        if dia == 1 { // Terça
            continue Semana // Pula para próxima semana
        }
        trabalhar(dia)
    }
```

**Analogia:**
É como **pular uma iteração completa** do loop externo!

---

## 14. Modificar Durante Iteração: Atualizar Lista

### Analogia: Atualizar Preços na Prateleira

Modificar elementos durante iteração é como **atualizar preços** enquanto você caminha pela loja:

**Mundo real:**

```
Você é funcionário atualizando preços:
- Ver produto 1: atualizar preço ✅ (SEGURO)
- Ver produto 2: atualizar preço ✅ (SEGURO)
- Ver produto 3: atualizar preço ✅ (SEGURO)
```

**Loop em Go (SEGURO):**

```go
produtos := []Produto{...}

for i := range produtos {
    produtos[i].Preco *= 1.1 // Aumentar 10% - SEGURO!
}
```

**⚠️ CUIDADO: Adicionar/Remover**

**Mundo real:**

```
Você está contando produtos:
- Contar produto 1
- Adicionar produto novo enquanto conta
- Contar produto 2
- Você pode perder a conta! ❌
```

**Loop em Go (CUIDADO):**

```go
// EVITE adicionar durante range
for _, produto := range produtos {
    if condicao {
        produtos = append(produtos, novoProduto) // Pode causar problemas!
    }
}
```

**Analogia:**

- **Modificar elemento existente**: Como **atualizar** um preço na etiqueta (seguro)
- **Adicionar/remover**: Como **adicionar produtos** enquanto conta (pode confundir!)

---

## 15. Iterar String: Soletrar Corretamente

### Analogia: Soletrar Palavra com Acentos

Iterar sobre strings requer cuidado especial:

**Palavra: "Café"**

**ERRADO** (vê bytes):

```
C (byte 1)
a (byte 2)
f (byte 3)
é (byte 4 e 5 - DOIS bytes!)
```

**CORRETO** (vê letras/runes):

```
C (letra 1)
a (letra 2)
f (letra 3)
é (letra 4 - uma letra completa!)
```

**Loop em Go:**

```go
palavra := "Café"

// CORRETO: for range
for i, letra := range palavra {
    fmt.Printf("Letra %d: %c\n", i, letra)
}

// ERRADO: indexação direta
for i := 0; i < len(palavra); i++ {
    fmt.Printf("Byte %d: %d\n", i, palavra[i]) // Vê bytes, não letras!
}
```

**Analogia:**

- **for range**: Como **soletrar** palavra por palavra (vê cada letra completa)
- **Indexação direta**: Como olhar os **pixels** individuais (não vê a letra completa)

---

## 16. Padrões Comuns: Analogias do Dia a Dia

### Padrão 1: Buscar (Procurar Item)

**Analogia: Procurar Nome na Lista de Telefone**

```go
// Procurar "João" na lista
for i, nome := range nomes {
    if nome == "João" {
        fmt.Printf("Encontrei no índice %d!\n", i)
        break // Para de procurar
    }
}
```

### Padrão 2: Filtrar (Separar Itens)

**Analogia: Separar Frutas Boas**

```go
// Separar apenas frutas boas
frutasBoas := []string{}
for _, fruta := range frutas {
    if fruta == "boa" {
        frutasBoas = append(frutasBoas, fruta)
    }
}
```

### Padrão 3: Somar (Acumular Valores)

**Analogia: Somar Compras do Supermercado**

```go
// Somar preços
total := 0
for _, preco := range precos {
    total += preco
}
```

### Padrão 4: Contar (Contar Ocorrências)

**Analogia: Contar Quantas Maçãs Tem**

```go
// Contar maçãs
contador := 0
for _, fruta := range frutas {
    if fruta == "maçã" {
        contador++
    }
}
```

---

## 17. Resumo Visual: Analogias dos Loops

**for Clássico:**

```
📋 Lista de Tarefas Numeradas
1. Fazer X
2. Fazer Y
3. Fazer Z
```

**for While-Style:**

```
⏳ Esperar até algo acontecer
"Enquanto não acontecer, continue esperando"
```

**for Infinito:**

```
🔄 Serviço 24 horas
"Sempre disponível até ser desligado"
```

**for range:**

```
👀 Ver cada item de uma lista
"Passar os olhos por cada item"
```

**break:**

```
🛑 Parar quando encontrar
"Encontrei! Para de procurar!"
```

**continue:**

```
⏭️ Pular e continuar
"Ignorar isso e continuar com o resto"
```

---

## 18. Regra de Ouro Simples

**Loops = Repetir Tarefas**

- **for clássico**: Quando sabe quantas vezes repetir
- **for while**: Quando não sabe, mas tem condição
- **for range**: Para ver cada item de uma lista (PREFERIDO!)
- **break**: Parar quando encontrar o que procura
- **continue**: Pular algo e continuar

**Quando usar cada um:**

- Precisa repetir código? → Use `for`
- Precisa ver cada item de lista? → Use `for range`
- Precisa parar quando encontrar? → Use `break`
- Precisa pular alguns itens? → Use `continue`

---

Agora que você entendeu os conceitos de forma simplificada, vamos praticar com exercícios na próxima parte!
