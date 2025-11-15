# Módulo 3: Tipos de Dados em Go
## Aula 4 - Performance, Boas Práticas e O Que Deve/Não Deve Ser Utilizado

Agora que você entende os tipos de dados, vamos falar sobre **como usá-los de forma eficiente e profissional**. Esta é a parte que separa programadores iniciantes de programadores experientes.

---

## 1. Escolha de Tipos: Performance vs. Praticidade

### ❌ O Que NÃO Deve Ser Feito

#### 1.1. Usar Tipos Muito Específicos Sem Necessidade

```go
// ❌ RUIM: Usar int8 quando int seria suficiente
var contador int8 = 0
for i := int8(0); i < 100; i++ {
    contador++
    // Problema: Se o contador passar de 127, terá overflow!
}

// ✅ BOM: Usar int (mais seguro e idiomático)
var contador int = 0
for i := 0; i < 100; i++ {
    contador++
}
```

**Por quê?**: A menos que você tenha uma razão muito específica (economia de memória em arrays grandes, protocolos de rede, etc.), use `int` e `uint`. Eles são mais seguros e o código fica mais legível.

#### 1.2. Misturar Tipos de Inteiros Sem Conversão Explícita

```go
// ❌ RUIM: Tentar operar tipos diferentes sem conversão
var a int = 10
var b int64 = 20
soma := a + b  // ERRO DE COMPILAÇÃO!

// ✅ BOM: Converter explicitamente
var a int = 10
var b int64 = 20
soma := int64(a) + b  // Conversão explícita
```

**Por quê?**: Go força você a ser explícito para evitar bugs sutis. Isso é uma **feature**, não um bug!

#### 1.3. Usar float32 Quando float64 é o Padrão

```go
// ❌ EVITE (a menos que tenha razão específica)
var preco float32 = 19.99

// ✅ BOM: Use float64 (padrão do Go)
var preco float64 = 19.99
// ou simplesmente
var preco = 19.99  // Go infere como float64
```

**Por quê?**: `float64` é o padrão em Go. Usar `float32` só faz sentido quando você tem restrições de memória muito específicas (ex: processamento de imagens com milhões de pixels).

---

## 2. Performance: Quando Cada Tipo Importa

### 2.1. Arrays e Slices Grandes: Tipos Específicos Fazem Diferença

```go
// Cenário: Array com 10 milhões de elementos

// ❌ Usa 80 MB de memória (10.000.000 * 8 bytes)
var numerosGrandes []int64 = make([]int64, 10_000_000)

// ✅ Usa 10 MB de memória (10.000.000 * 1 byte)
var numerosPequenos []int8 = make([]int8, 10_000_000)
```

**Quando otimizar**: 
- ✅ Arrays/slices com **milhões de elementos**
- ✅ Dados que serão serializados (JSON, protocolos de rede)
- ✅ Estruturas de dados que serão replicadas muitas vezes

**Quando NÃO otimizar**:
- ❌ Variáveis simples ou pequenos arrays
- ❌ Código onde legibilidade é mais importante
- ❌ Quando a economia de memória é desprezível

### 2.2. Conversões de Tipo: Custo de Performance

```go
// ⚠️ Conversões têm um custo mínimo, mas em loops muito apertados podem importar

// Em um loop de 1 bilhão de iterações:
for i := 0; i < 1_000_000_000; i++ {
    valor := int64(i)  // Conversão a cada iteração
    // ...
}

// ✅ Melhor: Use o tipo correto desde o início
for i := int64(0); i < 1_000_000_000; i++ {
    // Sem conversão necessária
}
```

**Regra geral**: Em código normal, o custo de conversão é desprezível. Só se preocupe com isso em **hot paths** (código que roda bilhões de vezes).

---

## 3. Boas Práticas: O Que Deve Ser Utilizado

### 3.1. Use `int` e `uint` por Padrão

```go
// ✅ IDIOMÁTICO EM GO
var idade int = 25
var contador int = 0
var tamanho uint = 100

// Use tipos específicos (int8, int64, etc.) apenas quando:
// - Você tem restrições de memória específicas
// - Você está trabalhando com protocolos que exigem tamanhos específicos
// - Você precisa de ranges específicos e quer que o compilador valide
```

**Por quê?**: 
- Código mais legível
- Menos chance de overflow
- Portabilidade entre arquiteturas
- É o estilo idiomático de Go

### 3.2. Use `float64` por Padrão

```go
// ✅ PADRÃO
var altura float64 = 1.75
var peso = 70.5  // Inferido como float64

// Use float32 apenas quando:
// - Processamento de imagens/áudio (onde cada byte importa)
// - Protocolos que exigem precisão simples
// - Arrays enormes onde memória é crítica
```

### 3.3. Use Raw Strings para Regex, SQL, JSON, Caminhos

```go
// ✅ BOM: Raw string para regex
pattern := `^\d{3}-\d{2}-\d{4}$`

// ✅ BOM: Raw string para SQL
query := `SELECT nome, idade 
          FROM usuarios 
          WHERE idade > 18`

// ✅ BOM: Raw string para caminhos do Windows
caminho := `C:\Users\Documentos\arquivo.txt`

// ❌ EVITE: Interpreted string para esses casos
pattern := "^\\d{3}-\\d{2}-\\d{4}$"  // Difícil de ler!
```

**Por quê?**: Raw strings são mais legíveis e menos propensas a erros de escape.

### 3.4. Sempre Trate Erros de Conversão

```go
// ✅ BOM: Verificar erros
texto := "42"
numero, err := strconv.Atoi(texto)
if err != nil {
    log.Fatal("Erro ao converter:", err)
}

// ❌ RUIM: Ignorar erros (pode causar panics)
texto := "abc"
numero, _ := strconv.Atoi(texto)  // numero será 0, mas isso é um erro silencioso!
```

---

## 4. Armadilhas Comuns e Como Evitá-las

### 4.1. Overflow Silencioso

```go
// ⚠️ CUIDADO: Overflow não gera erro em tempo de execução!
var pequeno int8 = 127
pequeno++  // Agora pequeno vale -128 (overflow)!

// ✅ SOLUÇÃO: Use tipos maiores ou valide ranges
var numero int = 127
if numero < math.MaxInt8 {
    pequeno := int8(numero)
}
```

### 4.2. Conversão de String para Rune (Não é o que você pensa!)

```go
// ⚠️ CUIDADO: Isso NÃO converte número para string!
numero := 65
caractere := string(numero)  // Isso cria uma string com 1 byte de valor 65, não "65"!

// ✅ CORRETO: Use strconv
numero := 65
texto := strconv.Itoa(numero)  // "65"
```

### 4.3. Comparação de Floats (Erros de Precisão)

```go
// ⚠️ CUIDADO: Comparação direta pode falhar
var a float64 = 0.1 + 0.2
var b float64 = 0.3
if a == b {  // Pode ser false devido a erro de precisão!
    // ...
}

// ✅ CORRETO: Compare com tolerância
const epsilon = 1e-9
if math.Abs(a - b) < epsilon {
    // São "iguais" dentro da tolerância
}
```

---

## 5. Melhores Práticas para a Vida do Programador

### 5.1. Sempre Especifique Tipos em Interfaces Públicas

```go
// ✅ BOM: Tipo explícito em funções públicas
func CalcularArea(largura float64, altura float64) float64 {
    return largura * altura
}

// ⚠️ EVITE: Deixar inferência em APIs públicas (pode ser confuso)
func CalcularArea(largura, altura float64) float64 {
    // Mesmo resultado, mas menos explícito
}
```

### 5.2. Use Constantes Tipadas Quando Apropriado

```go
// ✅ BOM: Constante tipada evita conversões desnecessárias
const MaxUsers int = 1000

// ⚠️ EVITE: Constante não-tipada pode causar problemas
const MaxUsers = 1000  // Tipo será inferido no uso, pode ser problemático
```

### 5.3. Documente Escolhas de Tipo Não-Óbvias

```go
// ✅ BOM: Comente por que escolheu um tipo específico
// Usamos int8 para economizar memória, já que sabemos que
// a idade nunca excederá 127 (range válido: 0-120)
type Pessoa struct {
    Idade int8  // Otimização de memória para arrays grandes
}
```

---

## 6. Checklist de Revisão de Código

Antes de considerar seu código "pronto", pergunte-se:

- [ ] Estou usando `int`/`uint` por padrão, ou tenho uma razão específica para tipos menores?
- [ ] Estou usando `float64` por padrão, ou tenho uma razão específica para `float32`?
- [ ] Minhas strings raw estão sendo usadas para casos apropriados (regex, SQL, caminhos)?
- [ ] Estou tratando erros de conversão de tipos?
- [ ] Estou comparando floats com tolerância quando necessário?
- [ ] Meu código é legível e idiomático em Go?
- [ ] Documentei escolhas de tipo não-óbvias?

---

## 7. Resumo: Regras de Ouro

1. **Padrão**: Use `int`, `uint`, `float64` - são os tipos mais comuns e seguros
2. **Otimização**: Só otimize tipos quando há necessidade real (memória, performance medida)
3. **Conversão**: Sempre explícita, sempre com tratamento de erro quando aplicável
4. **Strings**: Raw para casos complexos, interpreted para casos normais
5. **Floats**: Nunca para dinheiro, sempre com tolerância em comparações
6. **Legibilidade**: Prefira código legível sobre micro-otimizações prematuras

---

## Conclusão

Entender tipos de dados vai além de saber a sintaxe. É sobre:

- **Escolher o tipo certo** para cada situação
- **Equilibrar performance e legibilidade**
- **Evitar armadilhas comuns** que causam bugs sutis
- **Escrever código idiomático** que outros programadores Go entenderão facilmente

Lembre-se: **Código é lido muito mais vezes do que escrito**. Priorize clareza, e otimize apenas quando necessário e medido.

---

Na próxima etapa, você completará os exercícios e eu analisarei seu desempenho. Boa sorte!


