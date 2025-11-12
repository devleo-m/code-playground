# Módulo 3: Tipos de Dados em Go
## Aula 1: Data Types - Tipos de Dados Fundamentais

Olá! Bem-vindo ao terceiro módulo. Agora que você já sabe declarar variáveis e constantes, é fundamental entender **o que** você pode guardar dentro delas. Go oferece um conjunto rico de tipos de dados embutidos, e entender cada um deles é crucial para escrever programas eficientes e confiáveis.

Go é uma linguagem **estaticamente tipada**, o que significa que os tipos são determinados em tempo de compilação. Isso permite detecção precoce de erros e melhor performance, mas exige que você entenda bem cada tipo disponível.

---

## 1. Integers (Inteiros)

### Signed Integers (Inteiros com Sinal)

Os inteiros com sinal podem representar números positivos e negativos. Go oferece quatro tamanhos:

```go
var a int8   = -128    // Vai de -128 até 127 (8 bits)
var b int16  = -32768  // Vai de -32.768 até 32.767 (16 bits)
var c int32  = -2147483648  // Vai de -2.147.483.648 até 2.147.483.647 (32 bits)
var d int64  = -9223372036854775808  // Vai de -9.223.372.036.854.775.808 até 9.223.372.036.854.775.807 (64 bits)
```

**O tipo `int`**: É um tipo especial que é **dependente da plataforma**. Em sistemas de 32 bits, `int` é equivalente a `int32`. Em sistemas de 64 bits (a maioria hoje em dia), `int` é equivalente a `int64`. Use `int` quando o tamanho exato não importa e você quer que o código seja portável.

```go
var numero int = 42  // O tamanho depende da arquitetura do sistema
```

### Unsigned Integers (Inteiros sem Sinal)

Os inteiros sem sinal **só podem representar números não-negativos** (zero e positivos), mas têm um **range maior** para números positivos do que os signed equivalentes:

```go
var a uint8   = 255   // Vai de 0 até 255 (8 bits)
var b uint16  = 65535 // Vai de 0 até 65.535 (16 bits)
var c uint32  = 4294967295  // Vai de 0 até 4.294.967.295 (32 bits)
var d uint64  = 18446744073709551615  // Vai de 0 até 18.446.744.073.709.551.615 (64 bits)
```

**O tipo `uint`**: Assim como `int`, é dependente da plataforma. Em sistemas de 64 bits, `uint` é equivalente a `uint64`.

### Quando Usar Cada Tipo?

- **`int`/`uint`**: Use quando o tamanho exato não importa (mais comum)
- **`int8`/`uint8`**: Para economizar memória quando você sabe que os valores cabem nesse range (ex: idade, notas de 0-100)
- **`int64`/`uint64`**: Para valores muito grandes (ex: timestamps, IDs de banco de dados)
- **Evite misturar**: Não misture diferentes tamanhos de inteiros sem conversão explícita

---

## 2. Floating Points (Números de Ponto Flutuante)

Go oferece dois tipos para representar números reais (com casas decimais):

### float32 (Precisão Simples)

Ocupa 32 bits e oferece aproximadamente 7 dígitos decimais de precisão.

```go
var preco float32 = 19.99
var temperatura float32 = -5.5
```

### float64 (Precisão Dupla - Padrão)

Ocupa 64 bits e oferece aproximadamente 15-17 dígitos decimais de precisão. **Este é o tipo padrão** quando você declara um número decimal sem especificar o tipo.

```go
var altura float64 = 1.75
var pi = 3.141592653589793  // Go infere como float64 automaticamente
```

### ⚠️ Importante: Erros de Precisão

Números de ponto flutuante seguem o padrão **IEEE 754** e podem introduzir erros de precisão. **NÃO são adequados para cálculos financeiros exatos**.

```go
var resultado float64 = 0.1 + 0.2
fmt.Println(resultado)  // Pode imprimir: 0.30000000000000004 (não exatamente 0.3!)
```

Para cálculos financeiros, use bibliotecas especializadas ou trabalhe com centavos usando inteiros.

---

## 3. Complex Numbers (Números Complexos)

Go tem suporte nativo para números complexos! Isso é raro em linguagens de programação.

### complex64 e complex128

```go
var z1 complex64  = 3 + 4i        // Parte real: 3, parte imaginária: 4
var z2 complex128 = 5.5 + 7.2i    // Precisão dupla

// Ou usando a função complex()
var z3 = complex(3.0, 4.0)  // Cria 3 + 4i
```

### Funções Úteis

```go
z := 3 + 4i

real(z)    // Retorna a parte real: 3
imag(z)    // Retorna a parte imaginária: 4
abs(z)     // Retorna o módulo (magnitude): 5 (raiz de 3² + 4²)
```

**Uso comum**: Processamento de sinais, computação científica, transformadas de Fourier, etc.

---

## 4. Boolean (Booleano)

O tipo `bool` representa valores lógicos: `true` (verdadeiro) ou `false` (falso).

```go
var estaChovendo bool = true
var temSol bool = false

// Valor zero de bool é false
var condicao bool  // Inicializa como false automaticamente
```

### Onde Surgem Valores Booleanos?

- **Comparações**: `==`, `!=`, `<`, `>`, `<=`, `>=`
- **Operações lógicas**: `&&` (E), `||` (OU), `!` (NÃO)

```go
idade := 18
maiorDeIdade := idade >= 18  // true
temCarteira := true
podeDirigir := maiorDeIdade && temCarteira  // true && true = true
```

---

## 5. Runes

Um `rune` é um **alias para `int32`** que representa um **ponto de código Unicode**. Isso permite que Go trabalhe corretamente com caracteres internacionais, emojis e qualquer caractere além do ASCII básico.

```go
var letra rune = 'A'        // Código Unicode 65
var chines rune = '中'      // Caractere chinês
var emoji rune = '🚀'       // Emoji de foguete

// Runes são números, então você pode fazer operações
var proximaLetra rune = 'A' + 1  // 'B'
```

### Por Que Runes São Importantes?

Sem runes, você não conseguiria processar corretamente textos em português (com acentos), chinês, árabe, emojis, etc. Eles garantem que seu programa funcione globalmente.

```go
texto := "Olá, 世界! 🎉"
for _, char := range texto {
    fmt.Printf("%c = %d\n", char, char)  // Imprime cada caractere e seu código Unicode
}
```

---

## 6. Strings (Cadeias de Caracteres)

Go oferece dois tipos de literais de string, cada um com seu propósito específico.

### Raw String Literals (Literais de String Brutos)

Envolvidos por **backticks** (`` ` ``) e interpretam **todos os caracteres literalmente**, sem processar sequências de escape.

```go
caminho := `C:\Users\Documentos\arquivo.txt`  // Não precisa escapar as barras invertidas
regex := `^\d{3}-\d{2}-\d{4}$`                // Regex fica mais limpo
sql := `SELECT * FROM usuarios 
        WHERE idade > 18`                      // Preserva quebras de linha
json := `{"nome": "João", "idade": 30}`       // JSON sem escapar aspas
```

**Ideal para**: Regex, caminhos de arquivo, SQL, JSON, templates, textos multi-linha onde escapar seria trabalhoso.

### Interpreted String Literals (Literais de String Interpretados)

Envolvidos por **aspas duplas** (`"`) e **processam sequências de escape**.

```go
mensagem := "Olá,\nMundo!"           // \n vira quebra de linha
caminho := "C:\\Users\\arquivo.txt"  // Precisa escapar a barra invertida
aspas := "Ele disse: \"Olá\""        // Precisa escapar as aspas
tab := "Nome\tIdade\tCidade"         // \t vira tabulação
```

**Sequências de escape comuns**:
- `\n` - Nova linha
- `\t` - Tabulação
- `\"` - Aspas duplas
- `\\` - Barra invertida
- `\uXXXX` - Caractere Unicode (4 dígitos hex)
- `\UXXXXXXXX` - Caractere Unicode (8 dígitos hex)

**Ideal para**: Textos normais que precisam de caracteres de controle (quebras de linha, tabs) ou formatação.

---

## 7. Type Conversion (Conversão de Tipos)

Go exige **conversão explícita** entre tipos, mesmo quando são relacionados. Não há conversão automática (coerção) como em algumas linguagens.

### Sintaxe: `Tipo(valor)`

```go
var x int = 42
var y int64 = int64(x)  // Converte int para int64
var z float64 = float64(x)  // Converte int para float64

var a float64 = 3.14
var b int = int(a)  // Converte float64 para int (trunca a parte decimal: 3)
```

### Conversões Comuns

```go
// Inteiros para strings
numero := 42
texto := string(numero)  // ⚠️ CUIDADO: Isso converte para o caractere Unicode 42, não "42"!
textoCorreto := strconv.Itoa(numero)  // Usa strconv para converter número para string

// Strings para inteiros
texto := "42"
numero, err := strconv.Atoi(texto)  // Retorna (int, error)

// Float para string
preco := 19.99
texto := strconv.FormatFloat(preco, 'f', 2, 64)  // "19.99"

// String para float
texto := "19.99"
preco, err := strconv.ParseFloat(texto, 64)  // Retorna (float64, error)
```

### ⚠️ Atenção: Conversões Podem Perder Dados

```go
var grande int64 = 9223372036854775807
var pequeno int8 = int8(grande)  // Overflow! O valor não cabe em int8
```

Sempre verifique se a conversão é segura antes de executá-la.

---

## Resumo dos Tipos e Seus Valores Zero

| Tipo | Valor Zero | Tamanho |
|------|------------|---------|
| `int`, `int8`, `int16`, `int32`, `int64` | `0` | Depende do tipo |
| `uint`, `uint8`, `uint16`, `uint32`, `uint64` | `0` | Depende do tipo |
| `float32`, `float64` | `0.0` | 32 ou 64 bits |
| `complex64`, `complex128` | `0+0i` | 64 ou 128 bits |
| `bool` | `false` | 1 bit |
| `rune` | `0` | 32 bits (int32) |
| `string` | `""` (string vazia) | Variável |

---

## Conclusão

Entender os tipos de dados em Go é fundamental porque:

1. **Detecção precoce de erros**: O compilador pega erros de tipo antes do programa rodar
2. **Performance**: Tipos explícitos permitem otimizações pelo compilador
3. **Clareza**: O código fica mais legível quando os tipos são explícitos
4. **Confiabilidade**: Evita bugs sutis causados por conversões automáticas inesperadas

Na próxima parte desta aula, vamos simplificar esses conceitos com analogias do dia a dia para fixar melhor o aprendizado!

