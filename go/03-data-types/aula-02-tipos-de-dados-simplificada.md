# Módulo 3: Tipos de Dados em Go
## Aula 2 - Simplificada: Entendendo Tipos de Dados

Agora vamos entender esses conceitos de forma mais simples, usando analogias do nosso dia a dia!

---

## 1. Integers: Os Números Inteiros (Como Caixas de Tamanhos Diferentes)

Imagine que você precisa guardar números em caixas. Go oferece caixas de vários tamanhos:

### Signed Integers (Caixas que Aceitam Positivos e Negativos)

Pense em uma **linha numérica** que vai do negativo ao positivo:

```
int8:   [-128 até 127]     → Uma caixinha pequena (8 bits)
int16:  [-32.768 até 32.767] → Uma caixa média (16 bits)
int32:  [Números muito grandes negativos até positivos] → Caixa grande (32 bits)
int64:  [Números ENORMES] → Caixa gigante (64 bits)
int:    [Tamanho depende do seu computador] → Caixa "inteligente" que se adapta
```

**Analogia**: É como escolher o tamanho de uma mala de viagem:
- `int8`: Mala de mão (só cabe o essencial)
- `int`: Mala padrão (serve para a maioria das viagens)
- `int64`: Container de navio (para coisas realmente grandes)

### Unsigned Integers (Caixas Só para Números Positivos)

Essas caixas **não aceitam números negativos**, mas podem guardar números positivos **maiores** que as signed equivalentes.

**Analogia**: É como uma balança de banheiro. Ela não mostra peso negativo (você não pode pesar -5kg), mas pode mostrar até 150kg. Se você sabe que nunca vai precisar de números negativos, use `uint` para ter mais espaço para números grandes positivos.

```go
// Idade nunca é negativa, então uint faz sentido
var idade uint8 = 25  // Vai de 0 até 255 (perfeito para idade!)

// Temperatura pode ser negativa, então int faz sentido
var temperatura int8 = -5  // Pode ser -128 até 127
```

---

## 2. Floating Points: Os Números com Vírgula (Como Medidas Precisas)

Pense em uma **régua** para medir coisas:

- **`float32`**: Uma régua comum, com precisão de **7 dígitos**. Serve para a maioria das medidas do dia a dia.
- **`float64`**: Uma régua **super precisa**, com precisão de **15-17 dígitos**. É como uma régua de engenheiro ou cientista.

**Analogia do Dia a Dia**:
```go
altura := 1.75        // float64 - Precisão suficiente para altura
peso := 70.5          // float64 - Precisão suficiente para peso
```

**⚠️ Cuidado com a Precisão!**

Imagine que você tem R$ 0,10 e quer somar com R$ 0,20. Você espera R$ 0,30, certo?

```go
total := 0.1 + 0.2
// Pode dar: 0.30000000000000004 (quase 0.30, mas não exatamente!)
```

**Analogia**: É como medir com uma régua que tem marcações muito pequenas. Às vezes, você não consegue medir exatamente, só "quase exato". Por isso, **nunca use float para dinheiro**! Use centavos como inteiros:

```go
// ❌ Ruim para dinheiro
preco := 19.99

// ✅ Bom para dinheiro
precoEmCentavos := 1999  // R$ 19,99 em centavos
```

---

## 3. Complex Numbers: Os Números "Imaginários" (Como Coordenadas no Plano)

Você se lembra de matemática? Números complexos têm uma parte "real" e uma parte "imaginária". Em Go, isso é nativo!

**Analogia**: Pense em um **mapa com coordenadas**:
- A parte real é a posição **horizontal** (leste/oeste)
- A parte imaginária é a posição **vertical** (norte/sul)

```go
ponto := 3 + 4i
// Está 3 unidades à direita e 4 unidades para cima
```

**Onde isso é útil?**: Processamento de imagens, sinais de rádio, gráficos 3D, física quântica. Para programação do dia a dia, você raramente vai usar, mas é bom saber que existe!

---

## 4. Boolean: O Interruptor (Ligado ou Desligado)

O tipo mais simples! É como um **interruptor de luz**:

- `true` = **Ligado** (a luz está acesa)
- `false` = **Desligado** (a luz está apagada)

**Analogias do Dia a Dia**:
```go
estaChovendo := true      // Sim, está chovendo
temSol := false           // Não, não tem sol
maiorDeIdade := true      // Sim, é maior de idade
temCarteira := false      // Não, não tem carteira
```

É assim que você toma **decisões** no código: "Se está chovendo E tenho guarda-chuva, então saio. Senão, fico em casa."

---

## 5. Runes: Os "IDs" dos Caracteres (Como Códigos de Barras)

Cada caractere no computador tem um **número único**, como um código de barras. Esse número é chamado de "ponto de código Unicode".

**Analogia**: Pense em um **catálogo gigante** onde cada símbolo do mundo tem um número:
- 'A' = número 65
- '中' (chinês) = número 20013
- '🚀' (emoji) = número 128640

```go
letra := 'A'     // O "código de barras" do A é 65
chines := '中'   // O "código de barras" do caractere chinês é 20013
```

**Por que isso importa?**: Sem runes, seu programa não conseguiria entender português com acentos (á, é, ã), chinês, árabe, emojis, etc. Runes garantem que seu programa funcione **globalmente**!

---

## 6. Strings: Os Textos (Duas Formas de Escrever)

### Raw Strings: O "Modo Literal" (Como Copiar e Colar)

Quando você usa **backticks** (`` ` ``), é como se você estivesse **copiando e colando exatamente como está**, sem processar nada.

**Analogia**: É como escrever em um **bloco de notas** onde tudo aparece exatamente como você digita:

```go
caminho := `C:\Users\arquivo.txt`
// Go vê: C:\Users\arquivo.txt (exatamente assim, sem processar a \)
```

**Quando usar?**: Quando você quer que **tudo apareça literalmente**, como em:
- Caminhos de arquivo do Windows
- Códigos de regex complicados
- SQL queries longas
- Textos que têm muitas aspas ou barras

### Interpreted Strings: O "Modo Processado" (Como um Editor de Texto)

Quando você usa **aspas duplas** (`"`), Go **processa** sequências especiais.

**Analogia**: É como escrever em um **editor de texto** que entende comandos:
- `\n` = "pule uma linha"
- `\t` = "faça uma tabulação"
- `\"` = "coloque uma aspas aqui"

```go
mensagem := "Linha 1\nLinha 2"
// Go processa o \n e cria:
// Linha 1
// Linha 2
```

**Quando usar?**: Para textos normais que precisam de formatação (quebras de linha, tabs, etc.)

---

## 7. Type Conversion: As "Traduções" Entre Tipos

Go não "adivinha" conversões. Você precisa **explicitamente dizer** quando quer converter um tipo em outro.

**Analogia**: É como **traduzir entre idiomas**. Você não pode simplesmente falar português e esperar que alguém entenda inglês automaticamente. Precisa traduzir!

```go
// Tenho um número inteiro (42)
numeroInteiro := 42

// Quero um número decimal (42.0)
numeroDecimal := float64(numeroInteiro)  // "Traduzo" int para float64

// Quero uma string ("42")
numeroTexto := strconv.Itoa(numeroInteiro)  // "Traduzo" int para string
```

**⚠️ Cuidado com "Traduções" Perigosas**:

```go
numeroGrande := 1000
numeroPequeno := int8(numeroGrande)  // ⚠️ 1000 não cabe em int8 (max 127)!
// Resultado: overflow (o número "transborda" e vira outro valor)
```

É como tentar colocar um elefante em uma caixa de sapatos. Não cabe!

---

## Resumo Visual

Pense nos tipos como **ferramentas diferentes** para tarefas diferentes:

- **Inteiros**: Para contar coisas (1, 2, 3...)
- **Floats**: Para medir coisas (1.5kg, 2.3m...)
- **Boolean**: Para decisões (sim/não, ligado/desligado)
- **Strings**: Para textos e palavras
- **Runes**: Para caracteres individuais (especialmente internacionais)
- **Complex**: Para matemática avançada (raramente usado no dia a dia)

Cada ferramenta tem seu propósito. Use a ferramenta certa para a tarefa certa!

---

Agora que você entendeu os conceitos de forma simplificada, vamos praticar com exercícios na próxima parte!

