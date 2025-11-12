# 📘 CURSO DE PROGRAMAÇÃO GO - AULA 7

---

## **Aula 7: Strings (Cadeias de Caracteres)**

### 🎯 **Objetivos da Aula**
- Compreender a estrutura interna de strings em Go
- Entender a relação entre strings, bytes e runes
- Dominar operações comuns com strings
- Aprender sobre imutabilidade e suas implicações
- Trabalhar com Unicode e caracteres multibyte

---

### 📚 **1. Revisão Rápida das Aulas Anteriores**

Até agora aprendemos:
- **Arrays**: Estruturas de tamanho fixo
- **Slices**: Referências dinâmicas a arrays
- **Conversões**: Array ↔ Slice (referência vs cópia)

**Agora:** Strings são um tipo especial de slice imutável de bytes!

---

### 🔤 **2. O Que São Strings em Go?**

Em Go, uma **string** é uma sequência **imutável** de bytes. Internamente, uma string é estruturada como:

type StringHeader struct {
    Data uintptr  // Ponteiro para os dados (bytes)
    Len  int      // Comprimento em bytes
}

**Características fundamentais:**
1. **Imutáveis**: Uma vez criada, não pode ser modificada
2. **UTF-8**: Go usa UTF-8 por padrão para codificar caracteres
3. **Slice de bytes**: Conceitualmente similar a `[]byte`, mas read-only

---

### 💻 **3. Criando e Manipulando Strings**

#### **Exemplo 1: Declaração Básica**

package main

import "fmt"

func main() {
    // Diferentes formas de criar strings
    var str1 string = "Hello, World!"
    str2 := "Go é incrível"
    str3 := `String com
múltiplas linhas
usando backticks`
    
    fmt.Println(str1)
    fmt.Println(str2)
    fmt.Println(str3)
    
    // String vazia
    var vazia string
    fmt.Printf("String vazia: '%s' (len=%d)\n", vazia, len(vazia))
}

**Saída:**

Hello, World!
Go é incrível
String com
múltiplas linhas
usando backticks
String vazia: '' (len=0)

**Nota:** Backticks (`) preservam tudo literalmente, incluindo quebras de linha e caracteres especiais.

---

#### **Exemplo 2: Strings Raw vs Interpreted**

package main

import "fmt"

func main() {
    // String interpretada (processa escape sequences)
    interpretada := "Linha 1\nLinha 2\tTabulação"
    
    // String raw (literal, ignora escape sequences)
    raw := `Linha 1\nLinha 2\tTabulação`
    
    fmt.Println("=== INTERPRETADA ===")
    fmt.Println(interpretada)
    
    fmt.Println("\n=== RAW ===")
    fmt.Println(raw)
}

**Saída:**

=== INTERPRETADA ===
Linha 1
Linha 2	Tabulação

=== RAW ===
Linha 1\nLinha 2\tTabulação

---

### 📏 **4. Comprimento de Strings: len() vs RuneCount**

#### **Exemplo 3: O Problema do len()**

package main

import (
    "fmt"
    "unicode/utf8"
)

func main() {
    // String ASCII simples
    str1 := "Hello"
    fmt.Printf("'%s' -> len=%d\n", str1, len(str1))
    
    // String com caracteres especiais
    str2 := "Olá"
    fmt.Printf("'%s' -> len=%d\n", str2, len(str2))
    
    // String com emoji
    str3 := "Go 🚀"
    fmt.Printf("'%s' -> len=%d\n", str3, len(str3))
    
    // Contando runes (caracteres Unicode)
    fmt.Println("\n=== CONTAGEM DE RUNES ===")
    fmt.Printf("'%s' -> %d runes\n", str1, utf8.RuneCountInString(str1))
    fmt.Printf("'%s' -> %d runes\n", str2, utf8.RuneCountInString(str2))
    fmt.Printf("'%s' -> %d runes\n", str3, utf8.RuneCountInString(str3))
}

**Saída:**

'Hello' -> len=5
'Olá' -> len=4
'Go 🚀' -> len=7

=== CONTAGEM DE RUNES ===
'Hello' -> 5 runes
'Olá' -> 3 runes
'Go 🚀' -> 4 runes

**Por quê?**
- `len()` retorna o **número de bytes**
- "á" usa 2 bytes em UTF-8
- "🚀" usa 4 bytes em UTF-8
- `utf8.RuneCountInString()` conta **caracteres reais**

---

### 🔍 **5. Acessando Caracteres: Bytes vs Runes**

#### **Exemplo 4: Indexação e Iteração**

package main

import "fmt"

func main() {
    texto := "Café"
    
    fmt.Println("=== ACESSANDO POR ÍNDICE (BYTES) ===")
    for i := 0; i < len(texto); i++ {
        fmt.Printf("texto[%d] = %c (byte: %d)\n", i, texto[i], texto[i])
    }
    
    fmt.Println("\n=== ITERANDO COM RANGE (RUNES) ===")
    for indice, caractere := range texto {
        fmt.Printf("Posição %d: %c (rune: %d)\n", indice, caractere, caractere)
    }
}

**Saída:**

=== ACESSANDO POR ÍNDICE (BYTES) ===
texto[0] = C (byte: 67)
texto[1] = a (byte: 97)
texto[2] = f (byte: 102)
texto[3] = Ã (byte: 195)
texto[4] = © (byte: 169)

=== ITERANDO COM RANGE (RUNES) ===
Posição 0: C (rune: 67)
Posição 1: a (rune: 97)
Posição 2: f (rune: 102)
Posição 3: é (rune: 233)

**Análise Crítica:**
- Acessar por índice (`texto[i]`) retorna **bytes individuais**
- Iterar com `range` retorna **runes (caracteres Unicode completos)**
- "é" ocupa 2 bytes (índices 3 e 4), mas é 1 rune

---

### 🔄 **6. Conversão: String ↔ []byte ↔ []rune**

#### **Exemplo 5: Conversões Básicas**

package main

import "fmt"

func main() {
    original := "Go! 🎉"
    
    // String → []byte
    bytes := []byte(original)
    fmt.Println("Bytes:", bytes)
    
    // String → []rune
    runes := []rune(original)
    fmt.Println("Runes:", runes)
    
    // []byte → String
    deBytes := string(bytes)
    fmt.Println("De bytes:", deBytes)
    
    // []rune → String
    deRunes := string(runes)
    fmt.Println("De runes:", deRunes)
    
    // Comparando tamanhos
    fmt.Printf("\nTamanhos:\n")
    fmt.Printf("String: %d bytes\n", len(original))
    fmt.Printf("[]byte: %d elementos\n", len(bytes))
    fmt.Printf("[]rune: %d elementos\n", len(runes))
}

**Saída:**

Bytes: [71 111 33 32 240 159 142 137]
Runes: [71 111 33 32 127881]
De bytes: Go! 🎉
De runes: Go! 🎉

Tamanhos:
String: 8 bytes
[]byte: 8 elementos
[]rune: 5 elementos

**Observações:**
- `[]byte` preserva a representação UTF-8 (cada byte)
- `[]rune` converte para valores Unicode (cada caractere)
- Emoji 🎉 = 1 rune (127881) = 4 bytes (240, 159, 142, 137)

---

### ⚠️ **7. Imutabilidade: Por Que Strings Não Podem Ser Modificadas**

#### **Exemplo 6: Tentando Modificar uma String**

package main

import "fmt"

func main() {
    texto := "Hello"
    
    // ❌ ERRO DE COMPILAÇÃO!
    // texto[0] = 'J'
    
    // ✅ SOLUÇÃO 1: Converter para []byte
    bytes := []byte(texto)
    bytes[0] = 'J'
    texto = string(bytes)
    fmt.Println("Modificado via []byte:", texto)
    
    // ✅ SOLUÇÃO 2: Converter para []rune
    texto2 := "Café"
    runes := []rune(texto2)
    runes[3] = 'é'
    texto2 = string(runes)
    fmt.Println("Modificado via []rune:", texto2)
}

**Saída:**

Modificado via []byte: Jello
Modificado via []rune: Café

**Por Que Imutabilidade?**
1. **Segurança**: Strings podem ser compartilhadas sem medo de modificação
2. **Performance**: Strings podem ser copiadas apenas como referências
3. **Concorrência**: Strings são naturalmente thread-safe

---

### ✂️ **8. Operações Comuns com Strings**

#### **Exemplo 7: Concatenação**

package main

import (
    "fmt"
    "strings"
)

func main() {
    // Método 1: Operador +
    str1 := "Hello"
    str2 := "World"
    resultado := str1 + " " + str2
    fmt.Println("Operador +:", resultado)
    
    // Método 2: fmt.Sprintf
    nome := "João"
    idade := 25
    mensagem := fmt.Sprintf("%s tem %d anos", nome, idade)
    fmt.Println("fmt.Sprintf:", mensagem)
    
    // Método 3: strings.Join
    palavras := []string{"Go", "é", "incrível"}
    frase := strings.Join(palavras, " ")
    fmt.Println("strings.Join:", frase)
    
    // Método 4: strings.Builder (mais eficiente)
    var builder strings.Builder
    builder.WriteString("Construindo ")
    builder.WriteString("uma ")
    builder.WriteString("string ")
    builder.WriteString("eficientemente")
    fmt.Println("strings.Builder:", builder.String())
}

**Saída:**

Operador +: Hello World
fmt.Sprintf: João tem 25 anos
strings.Join: Go é incrível
strings.Builder: Construindo uma string eficientemente

---

#### **Exemplo 8: Substring (Fatiamento)**

package main

import "fmt"

func main() {
    texto := "Programação em Go"
    
    // Fatiamento básico (trabalha com bytes!)
    parte1 := texto[0:11]   // "Programação"
    parte2 := texto[12:]    // "em Go"
    
    fmt.Println("Parte 1:", parte1)
    fmt.Println("Parte 2:", parte2)
    
    // ⚠️ CUIDADO com caracteres multibyte!
    textoComAcento := "São Paulo"
    fmt.Println("Errado:", textoComAcento[0:3])  // Corta no meio do 'ã'
    
    // ✅ Correto: converter para []rune primeiro
    runes := []rune(textoComAcento)
    fmt.Println("Correto:", string(runes[0:3]))  // "São"
}

**Saída:**

Parte 1: Programação
Parte 2: em Go
Errado: S√
Correto: São

---

#### **Exemplo 9: Busca e Substituição**

package main

import (
    "fmt"
    "strings"
)

func main() {
    texto := "Go é rápido. Go é eficiente. Go é moderno."
    
    // Contém substring?
    fmt.Println("Contém 'rápido'?", strings.Contains(texto, "rápido"))
    
    // Encontrar índice
    indice := strings.Index(texto, "eficiente")
    fmt.Println("Índice de 'eficiente':", indice)
    
    // Contar ocorrências
    count := strings.Count(texto, "Go")
    fmt.Println("Quantas vezes 'Go' aparece:", count)
    
    // Substituir
    novo := strings.Replace(texto, "Go", "Golang", 2)  // Substitui 2 vezes
    fmt.Println("Substituindo:", novo)
    
    // Substituir todas
    novoTodas := strings.ReplaceAll(texto, "Go", "Golang")
    fmt.Println("Substituindo todas:", novoTodas)
    
    // Verificar prefixo/sufixo
    fmt.Println("Começa com 'Go'?", strings.HasPrefix(texto, "Go"))
    fmt.Println("Termina com 'moderno.'?", strings.HasSuffix(texto, "moderno."))
}

**Saída:**

Contém 'rápido'? true
Índice de 'eficiente': 17
Quantas vezes 'Go' aparece: 3
Substituindo: Golang é rápido. Golang é eficiente. Go é moderno.
Substituindo todas: Golang é rápido. Golang é eficiente. Golang é moderno.
Começa com 'Go'? true
Termina com 'moderno.'? true

---

#### **Exemplo 10: Transformações**

package main

import (
    "fmt"
    "strings"
)

func main() {
    texto := "  Olá, Mundo!  "
    
    // Maiúsculas/Minúsculas
    fmt.Println("Maiúsculas:", strings.ToUpper(texto))
    fmt.Println("Minúsculas:", strings.ToLower(texto))
    fmt.Println("Title Case:", strings.Title(texto))
    
    // Remover espaços
    fmt.Println("Trim:", strings.TrimSpace(texto))
    
    // Dividir em palavras
    palavras := strings.Fields(strings.TrimSpace(texto))
    fmt.Println("Palavras:", palavras)
    
    // Split customizado
    csv := "maçã,banana,laranja"
    frutas := strings.Split(csv, ",")
    fmt.Println("Frutas:", frutas)
    
    // Repetir
    fmt.Println("Repetido:", strings.Repeat("Go! ", 3))
}

**Saída:**

Maiúsculas:   OLÁ, MUNDO!  
Minúsculas:   olá, mundo!  
Title Case:   Olá, Mundo!  
Trim: Olá, Mundo!
Palavras: [Olá, Mundo!]
Frutas: [maçã banana laranja]
Repetido: Go! Go! Go! 

---

### 🏗️ **9. strings.Builder: Construção Eficiente**

#### **Exemplo 11: Por Que Usar Builder?**

package main

import (
    "fmt"
    "strings"
    "time"
)

func concatenacaoIneficiente(n int) string {
    resultado := ""
    for i := 0; i < n; i++ {
        resultado += "a"  // ❌ Cria nova string a cada iteração!
    }
    return resultado
}

func concatenacaoEficiente(n int) string {
    var builder strings.Builder
    builder.Grow(n)  // Pré-aloca memória
    for i := 0; i < n; i++ {
        builder.WriteString("a")  // ✅ Reutiliza buffer interno
    }
    return builder.String()
}

func main() {
    n := 10000
    
    // Método ineficiente
    inicio := time.Now()
    concatenacaoIneficiente(n)
    duracao1 := time.Since(inicio)
    
    // Método eficiente
    inicio = time.Now()
    concatenacaoEficiente(n)
    duracao2 := time.Since(inicio)
    
    fmt.Printf("Ineficiente: %v\n", duracao1)
    fmt.Printf("Eficiente: %v\n", duracao2)
    fmt.Printf("Speedup: %.2fx mais rápido\n", float64(duracao1)/float64(duracao2))
}

**Saída típica:**

Ineficiente: 15.2ms
Eficiente: 0.3ms
Speedup: 50.67x mais rápido

**Por quê?**
- Concatenação com `+` cria nova string a cada vez (O(n²))
- `strings.Builder` reutiliza buffer (O(n))

---

#### **Exemplo 12: Uso Prático do Builder**

package main

import (
    "fmt"
    "strings"
)

func gerarHTML(titulo string, itens []string) string {
    var builder strings.Builder
    
    builder.WriteString("<html>\n")
    builder.WriteString("  <head>\n")
    builder.WriteString("    <title>")
    builder.WriteString(titulo)
    builder.WriteString("</title>\n")
    builder.WriteString("  </head>\n")
    builder.WriteString("  <body>\n")
    builder.WriteString("    <ul>\n")
    
    for _, item := range itens {
        builder.WriteString("      <li>")
        builder.WriteString(item)
        builder.WriteString("</li>\n")
    }
    
    builder.WriteString("    </ul>\n")
    builder.WriteString("  </body>\n")
    builder.WriteString("</html>")
    
    return builder.String()
}

func main() {
    html := gerarHTML("Minha Lista", []string{"Item 1", "Item 2", "Item 3"})
    fmt.Println(html)
}

**Saída:**

<html>
  <head>
    <title>Minha Lista</title>
  </head>
  <body>
    <ul>
      <li>Item 1</li>
      <li>Item 2</li>
      <li>Item 3</li>
    </ul>
  </body>
</html>

---

### 🌍 **10. Trabalhando com Unicode**

#### **Exemplo 13: Runes e Caracteres Especiais**

package main

import (
    "fmt"
    "unicode"
)

func main() {
    texto := "Go 🚀 支持 中文 ñ"
    
    fmt.Println("Análise de cada rune:")
    for i, r := range texto {
        fmt.Printf("Posição %d: '%c' (U+%04X) - ", i, r, r)
        
        // Classificar o caractere
        switch {
        case unicode.IsLetter(r):
            fmt.Println("Letra")
        case unicode.IsDigit(r):
            fmt.Println("Dígito")
        case unicode.IsSpace(r):
            fmt.Println("Espaço")
        case unicode.IsSymbol(r):
            fmt.Println("Símbolo")
        default:
            fmt.Println("Outro")
        }
    }
}

**Saída:**

Análise de cada rune:
Posição 0: 'G' (U+0047) - Letra
Posição 1: 'o' (U+006F) - Letra
Posição 2: ' ' (U+0020) - Espaço
Posição 3: '🚀' (U+1F680) - Símbolo
Posição 7: ' ' (U+0020) - Espaço
Posição 8: '支' (U+652F) - Letra
Posição 11: '持' (U+6301) - Letra
Posição 14: ' ' (U+0020) - Espaço
Posição 15: '中' (U+4E2D) - Letra
Posição 18: '文' (U+6587) - Letra
Posição 21: ' ' (U+0020) - Espaço
Posição 22: 'ñ' (U+00F1) - Letra

---

#### **Exemplo 14: Normalização de Strings**

package main

import (
    "fmt"
    "strings"
    "unicode"
)

// Remover acentos (simplificado)
func removerAcentos(texto string) string {
    var builder strings.Builder
    
    for _, r := range texto {
        // Mapeamento simplificado
        switch r {
        case 'á', 'à', 'ã', 'â':
            builder.WriteRune('a')
        case 'é', 'ê':
            builder.WriteRune('e')
        case 'í':
            builder.WriteRune('i')
        case 'ó', 'ô', 'õ':
            builder.WriteRune('o')
        case 'ú':
            builder.WriteRune('u')
        case 'ç':
            builder.WriteRune('c')
        default:
            builder.WriteRune(unicode.ToLower(r))
        }
    }
    
    return builder.String()
}

func main() {
    textos := []string{
        "São Paulo",
        "Programação",
        "Café",
        "José",
    }
    
    for _, texto := range textos {
        normalizado := removerAcentos(texto)
        fmt.Printf("%s → %s\n", texto, normalizado)
    }
}

**Saída:**

São Paulo → sao paulo
Programação → programacao
Café → cafe
José → jose

---

### 📊 **11. Comparação de Strings**

#### **Exemplo 15: Diferentes Formas de Comparar**

package main

import (
    "fmt"
    "strings"
)

func main() {
    str1 := "Go"
    str2 := "go"
    str3 := "Go"
    
    // Comparação direta (case-sensitive)
    fmt.Println("str1 == str2:", str1 == str2)
    fmt.Println("str1 == str3:", str1 == str3)
    
    // Comparação case-insensitive
    fmt.Println("EqualFold:", strings.EqualFold(str1, str2))
    
    // Comparação lexicográfica
    fmt.Println("Compare Go vs go:", strings.Compare(str1, str2))
    fmt.Println("Compare Go vs Go:", strings.Compare(str1, str3))
    
    // Ordenando strings
    palavras := []string{"zebra", "maçã", "banana", "árvo re"}
    fmt.Println("Antes:", palavras)
    
    // Ordenação padrão (pode ser inesperada com acentos)
    // Para ordenação correta de Unicode, use golang.org/x/text/collate
}

**Saída:**

str1 == str2: false
str1 == str3: true
EqualFold: true
Compare Go vs go: -1
Compare Go vs Go: 0

---

### 🔬 **12. Strings e Memória**

#### **Exemplo 16: Compartilhamento de Dados**

package main

import "fmt"

func main() {
    original := "Hello, World! This is a long string."
    
    // Substring compartilha dados com original
    parte := original[0:5]
    
    fmt.Println("Original:", original)
    fmt.Println("Parte:", parte)
    
    // Mesmo que 'parte' seja pequena, ela mantém referência
    // ao array completo de 'original' (memory leak potencial)
    
    // ✅ SOLUÇÃO: Forçar cópia
    parteCopia := string([]byte(original[0:5]))
    fmt.Println("Parte (cópia):", parteCopia)
    
    // Agora 'parteCopia' é independente e 'original' pode ser coletado
}

**Saída:**

Original: Hello, World! This is a long string.
Parte: Hello
Parte (cópia): Hello

**Importante:** Substrings mantêm referência à string original completa. Em loops ou processamento de grandes volumes, isso pode causar memory leaks!

---

### 🎯 **13. Casos de Uso Práticos**

#### **Exemplo 17: Validação de Email (Simplificado)**

package main

import (
    "fmt"
    "strings"
)

func validarEmail(email string) bool {
    // Validação MUITO simplificada
    email = strings.TrimSpace(email)
    
    if len(email) == 0 {
        return false
    }
    
    if !strings.Contains(email, "@") {
        return false
    }
    
    partes := strings.Split(email, "@")
    if len(partes) != 2 {
        return false
    }
    
    if len(partes[0]) == 0 || len(partes[1]) == 0 {
        return false
    }
    
    if !strings.Contains(partes[1], ".") {
        return false
    }
    
    return true
}

func main() {
    emails := []string{
        "usuario@exemplo.com",
        "invalido@",
        "@exemplo.com",
        "sem-arroba.com",
        "valido@dominio.com.br",
    }
    
    for _, email := range emails {
        valido := validarEmail(email)
        fmt.Printf("%s -> %v\n", email, valido)
    }
}

**Saída:**

usuario@exemplo.com -> true
invalido@ -> false
@exemplo.com -> false
sem-arroba.com -> false
valido@dominio.com.br -> true

---

#### **Exemplo 18: Processamento de CSV**

package main

import (
    "fmt"
    "strings"
)

func parsearCSV(linha string) []string {
    // Parser simplificado (não trata aspas)
    campos := strings.Split(linha, ",")
    
    // Remover espaços de cada campo
    for i := range campos {
        campos[i] = strings.TrimSpace(campos[i])
    }
    
    return campos
}

func main() {
    csv := `nome, idade, cidade
João, 25, São Paulo
Maria, 30, Rio de Janeiro
Pedro, 28, Belo Horizonte`
    
    linhas := strings.Split(csv, "\n")
    
    // Cabeçalho
    cabecalho := parsearCSV(linhas[0])
    fmt.Println("Cabeçalho:", cabecalho)
    fmt.Println()
    
    // Dados
    for i := 1; i < len(linhas); i++ {
        dados := parsearCSV(linhas[i])
        fmt.Printf("Registro %d:\n", i)
        for j, valor := range dados {
            fmt.Printf("  %s: %s\n", cabecalho[j], valor)
        }
        fmt.Println()
    }
}

**Saída:**

Cabeçalho: [nome idade cidade]

Registro 1:
  nome: João
  idade: 25
  cidade: São Paulo

Registro 2:
  nome: Maria
  idade: 30
  cidade: Rio de Janeiro

Registro 3:
  nome: Pedro
  idade: 28
  cidade: Belo Horizonte

---

#### **Exemplo 19: Template Simples**

package main

import (
    "fmt"
    "strings"
)

func substituirVariaveis(template string, variaveis map[string]string) string {
    resultado := template
    
    for chave, valor := range variaveis {
        placeholder := "{{" + chave + "}}"
        resultado = strings.ReplaceAll(resultado, placeholder, valor)
    }
    
    return resultado
}

func main() {
    template := "Olá, {{nome}}! Você tem {{idade}} anos e mora em {{cidade}}."
    
    dados := map[string]string{
        "nome":   "Carlos",
        "idade":  "35",
        "cidade": "Florianópolis",
    }
    
    mensagem := substituirVariaveis(template, dados)
    fmt.Println(mensagem)
}

**Saída:**

Olá, Carlos! Você tem 35 anos e mora em Florianópolis.

---

### 📚 **14. Resumo dos Conceitos-Chave**

1. **Strings são imutáveis**: Não podem ser modificadas após criação
2. **UTF-8 por padrão**: Suporte nativo a Unicode
3. **len() conta bytes**: Use `utf8.RuneCountInString()` para caracteres
4. **range itera runes**: Iteração automática em caracteres Unicode
5. **[]byte vs []rune**: Bytes preservam UTF-8, runes são caracteres
6. **strings.Builder**: Forma eficiente de construir strings
7. **Substring compartilha memória**: Cuidado com memory leaks
8. **Package strings**: Funções utilitárias poderosas

---

## **Aula 7 - Simplificada: Entendendo Strings**

### 📝 **A Analogia do Livro**

Imagine que uma string é como um **livro já impresso**:

**Características do Livro (String):**
- Uma vez impresso, você NÃO pode mudar o texto
- Você pode LER qualquer página
- Você pode COPIAR partes para outro livro
- Você pode criar um NOVO livro com texto modificado

**Não é um caderno (slice mutável):** Você não pode usar borracha e reescrever!

---

### 🔤 **Bytes vs Caracteres: A Confusão**

**Analogia: Caixas e Objetos**

Imagine que você tem uma estante:

**Bytes = Caixas:**
- Cada caixa tem tamanho fixo
- Algumas coisas cabem em 1 caixa
- Outras precisam de 2, 3 ou 4 caixas

**Runes = Objetos Completos:**
- Um objeto pode ocupar múltiplas caixas
- Mas é sempre contado como 1 objeto

**Exemplo visual:**

texto := "Café"

// Caixas (bytes):
// [C] [a] [f] [é-parte1] [é-parte2]
//  1   1   1       2 caixas para 'é'
// Total: 5 bytes

// Objetos (runes):
// [C] [a] [f] [é]
//  1   1   1   1
// Total: 4 caracteres

---

### 🎯 **Regra Simples de Ouro**

**Use len() quando:**
- Você quer saber quantos BYTES (caixas)
- Está alocando memória
- Está trabalhando com arquivos/rede

**Use utf8.RuneCountInString() quando:**
- Você quer saber quantos CARACTERES (objetos)
- Está mostrando para o usuário
- Está validando tamanho de senha

---

### 🔧 **Por Que Strings São "Trancadas"?**

Imagine três situações:

**Situação 1: String Mutável (NÃO é assim em Go)**

var senha string = "minhasenha"
funcaoMalvada(senha)  // Essa função muda sua senha!
// Agora senha = "hackeado" ← PERIGO!

**Situação 2: String Imutável (Go real)**

var senha string = "minhasenha"
funcaoMalvada(senha)  // Tenta mudar, mas NÃO CONSEGUE
// senha continua "minhasenha" ← SEGURO!

**Moral:** Imutabilidade = Segurança e Previsibilidade

---

### 🔄 **Quando Você PRECISA Modificar**

**Processo em 3 passos:**

1. **Desmontar:** String → []byte ou []rune
2. **Modificar:** Mudar o que quiser
3. **Remontar:** []byte ou []rune → String

**Exemplo prático:**

// Quero trocar primeira letra
texto := "hello"

// Passo 1: Desmontar
bytes := []byte(texto)  // [h e l l o]

// Passo 2: Modificar
bytes[0] = 'H'  // [H e l l o]

// Passo 3: Remontar
novoTexto := string(bytes)  // "Hello"

---

### 🏗️ **strings.Builder: O Construtor de Lego**

**Analogia:**

**Método Ruim (Operador +):**
- Você tem uma torre de Lego
- Quer adicionar 1 peça?
- Desmonta TUDO
- Constrói tudo de novo com a peça nova
- Repete para cada peça → MUITO LENTO!

**Método Bom (strings.Builder):**
- Você tem uma base de Lego
- Quer adicionar 1 peça?
- Apenas ENCAIXA a peça nova
- Não precisa reconstruir nada → RÁPIDO!

**Código:**

// ❌ Lento
resultado := ""
for i := 0; i < 1000; i++ {
    resultado += "a"  // Reconstrói tudo 1000 vezes!
}

// ✅ Rápido
var builder strings.Builder
for i := 0; i < 1000; i++ {
    builder.WriteString("a")  // Apenas adiciona
}
resultado := builder.String()

---

### 🌍 **Unicode: O Mundo dos Caracteres**

**Analogia: Alfabetos Diferentes**

Imagine que o mundo tem vários alfabetos:
- Inglês: A, B, C (simples, 1 byte cada)
- Português: á, ç, õ (precisam de 2 bytes)
- Chinês: 中, 文 (precisam de 3 bytes)
- Emojis: 🚀, 😊 (precisam de 4 bytes)

**Go é inteligente:**
- Guarda tudo em UTF-8 (formato universal)
- Quando você itera com `range`, Go "traduz" automaticamente

**Exemplo:**

texto := "A é 中 🚀"

// Iterando (Go faz a mágica)
for _, char := range texto {
    fmt.Printf("%c ", char)
}
// Saída: A é 中 🚀 ← Caracteres corretos!

---

### ✂️ **Cortando Strings: A Tesoura Perigosa**

**PERIGO: Cortar no meio de um caractere!**

texto := "São"

// ❌ ERRADO: Cortar bytes
fmt.Println(texto[0:3])  // "S√" ← Cortou o 'ã' pela metade!

// ✅ CERTO: Converter para runes primeiro
runes := []rune(texto)
fmt.Println(string(runes[0:3]))  // "São" ← Perfeito!

**Regra prática:**
- Se o texto pode ter acentos/emojis: Use []rune
- Se é só ASCII (A-Z, 0-9): Pode usar bytes direto

---

### 🔍 **Buscando Coisas na String**

**Analogia: Procurando em um Livro**

livro := "Go é rápido. Go é eficiente."

// Tem essa palavra?
strings.Contains(livro, "rápido")  // true

// Em qual página (índice)?
strings.Index(livro, "eficiente")  // 17

// Quantas vezes aparece?
strings.Count(livro, "Go")  // 2

// Trocar uma palavra
strings.Replace(livro, "Go", "Golang", 1)  // Troca só a primeira

// Trocar todas
strings.ReplaceAll(livro, "Go", "Golang")  // Troca tudo

---

### 🎨 **Transformações Mágicas**

**Feitiços comuns:**

texto := "  OLÁ mundo  "

// Todas maiúsculas
strings.ToUpper(texto)  // "  OLÁ MUNDO  "

// Todas minúsculas
strings.ToLower(texto)  // "  olá mundo  "

// Remover espaços das pontas
strings.TrimSpace(texto)  // "OLÁ mundo"

// Separar em palavras
strings.Fields(strings.TrimSpace(texto))  // ["OLÁ", "mundo"]

// Juntar palavras
palavras := []string{"Go", "é", "legal"}
strings.Join(palavras, " ")  // "Go é legal"

---

### 📦 **Checklist Mental Rápido**

Antes de trabalhar com strings, pergunte:

1. **"Preciso modificar a string?"**
   - SIM: Converta para []byte ou []rune
   - NÃO: Trabalhe direto com a string

2. **"Tem caracteres especiais (acentos, emojis)?"**
   - SIM: Use []rune e range
   - NÃO: Pode usar []byte

3. **"Vou concatenar muitas vezes?"**
   - SIM: Use strings.Builder
   - NÃO: Operador + está OK

4. **"Preciso do tamanho?"**
   - Bytes: Use len()
   - Caracteres: Use utf8.RuneCountInString()

---

## **Aula 7 - Performance, Boas Práticas e Antipadrões**

### ⚡ **1. Performance: O Que Fazer e Não Fazer**

#### **Antipadrão 1: Concatenação em Loop**

**❌ MUITO RUIM:**

func construirGrande(n int) string {
    resultado := ""
    for i := 0; i < n; i++ {
        resultado += fmt.Sprintf("Item %d, ", i)  // Aloca n vezes!
    }
    return resultado
}

**Problema:** Para n=1000, cria 1000 strings temporárias!

**✅ SOLUÇÃO:**

func construirGrande(n int) string {
    var builder strings.Builder
    builder.Grow(n * 20)  // Pré-aloca memória estimada
    
    for i := 0; i < n; i++ {
        fmt.Fprintf(&builder, "Item %d, ", i)
    }
    return builder.String()
}

**Ganho:** 100x a 1000x mais rápido!

---

#### **Comparação de Performance**

package main

import (
    "fmt"
    "strings"
    "testing"
)

func BenchmarkConcatenacaoOperador(b *testing.B) {
    for i := 0; i < b.N; i++ {
        resultado := ""
        for j := 0; j < 100; j++ {
            resultado += "test"
        }
    }
}

func BenchmarkConcatenacaoBuilder(b *testing.B) {
    for i := 0; i < b.N; i++ {
        var builder strings.Builder
        builder.Grow(400)
        for j := 0; j < 100; j++ {
            builder.WriteString("test")
        }
        _ = builder.String()
    }
}

**Resultados típicos:**
- Operador +: ~50000 ns/op, 25000 B/op
- Builder: ~500 ns/op, 512 B/op
- **Speedup: 100x mais rápido, 50x menos memória!**

---

### ✅ **2. Boas Práticas**

#### **Prática 1: Sempre Use strings.Builder para Concatenações Dinâmicas**

**Quando usar:**
- Loops
- Construção de HTML/JSON/XML
- Formatação complexa
- Qualquer concatenação com mais de 3 operações

**Exemplo:**

func gerarRelatorio(dados []Item) string {
    var builder strings.Builder
    builder.Grow(len(dados) * 100)  // Estimativa
    
    builder.WriteString("=== RELATÓRIO ===\n")
    for _, item := range dados {
        fmt.Fprintf(&builder, "ID: %d | Nome: %s | Valor: %.2f\n", 
            item.ID, item.Nome, item.Valor)
    }
    
    return builder.String()
}

---

#### **Prática 2: Use []rune Para Manipulação de Caracteres Unicode**

**❌ ERRADO:**

func inverter(s string) string {
    bytes := []byte(s)
    // Inverte bytes - quebra caracteres multibyte!
    for i, j := 0, len(bytes)-1; i < j; i, j = i+1, j-1 {
        bytes[i], bytes[j] = bytes[j], bytes[i]
    }
    return string(bytes)
}

// "São" → "o√S" ← QUEBRADO!

**✅ CORRETO:**

func inverter(s string) string {
    runes := []rune(s)
    // Inverte caracteres completos
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}

// "São" → "oãS" ← CORRETO!

---

#### **Prática 3: Cuidado com Memory Leaks em Substrings**

**❌ PROBLEMA:**

func processarArquivoGrande() []string {
    conteudo := lerArquivoGigante()  // 1 GB
    
    linhas := strings.Split(conteudo, "\n")
    
    // Pegar apenas as 10 primeiras linhas
    resultado := linhas[:10]
    
    // PROBLEMA: 'resultado' mantém referência ao conteúdo completo de 1GB!
    return resultado
}

**✅ SOLUÇÃO:**

func processarArquivoGrande() []string {
    conteudo := lerArquivoGigante()  // 1 GB
    
    linhas := strings.Split(conteudo, "\n")
    
    // Criar cópias independentes
    resultado := make([]string, 10)
    for i := 0; i < 10; i++ {
        resultado[i] = string([]byte(linhas[i]))  // Força cópia
    }
    
    // Agora o conteúdo original pode ser liberado
    return resultado
}

---

#### **Prática 4: Valide e Sanitize Input do Usuário**

**Exemplo robusto:**

func validarUsername(username string) (string, error) {
    // 1. Remover espaços
    username = strings.TrimSpace(username)
    
    // 2. Validar tamanho
    runes := []rune(username)
    if len(runes) < 3 {
        return "", errors.New("username muito curto (mínimo 3 caracteres)")
    }
    if len(runes) > 20 {
        return "", errors.New("username muito longo (máximo 20 caracteres)")
    }
    
    // 3. Validar caracteres permitidos
    for _, r := range runes {
        if !unicode.IsLetter(r) && !unicode.IsDigit(r) && r != '_' {
            return "", fmt.Errorf("caractere inválido: %c", r)
        }
    }
    
    // 4. Normalizar (lowercase)
    username = strings.ToLower(username)
    
    return username, nil
}

func main() {
    testes := []string{
        "JoaoSilva",
        "ab",
        "usuario@invalido",
        "Nome_Válido123",
        "  EspacosNasPontas  ",
    }
    
    for _, teste := range testes {
        if valido, err := validarUsername(teste); err == nil {
            fmt.Printf("'%s' → '%s' ✅\n", teste, valido)
        } else {
            fmt.Printf("'%s' → ERRO: %v ❌\n", teste, err)
        }
    }
}

**Saída:**

'JoaoSilva' → 'joaosilva' ✅
'ab' → ERRO: username muito curto (mínimo 3 caracteres) ❌
'usuario@invalido' → ERRO: caractere inválido: @ ❌
'Nome_Válido123' → 'nome_válido123' ✅
'  EspacosNasPontas  ' → 'espacosnaspontas' ✅

---

#### **Prática 5: Use strings.EqualFold Para Comparação Case-Insensitive**

**❌ INEFICIENTE:**

func compararIgnorandoCase(s1, s2 string) bool {
    return strings.ToLower(s1) == strings.ToLower(s2)  // Aloca 2 strings novas!
}

**✅ EFICIENTE:**

func compararIgnorandoCase(s1, s2 string) bool {
    return strings.EqualFold(s1, s2)  // Não aloca memória
}

---

### ❌ **3. Antipadrões**

#### **Antipadrão 1: Usar + Para Construir Strings em Loop**

Já cobrimos isso, mas vale reforçar:

**❌ NUNCA FAÇA ISSO:**

func gerarSQL(campos []string) string {
    sql := "SELECT "
    for i, campo := range campos {
        if i > 0 {
            sql += ", "
        }
        sql += campo
    }
    sql += " FROM tabela"
    return sql
}

**✅ SEMPRE FAÇA ASSIM:**

func gerarSQL(campos []string) string {
    var builder strings.Builder
    builder.WriteString("SELECT ")
    builder.WriteString(strings.Join(campos, ", "))
    builder.WriteString(" FROM tabela")
    return builder.String()
}

---

#### **Antipadrão 2: Indexação Direta em Strings Unicode**

**❌ PERIGO:**

func primeiroCaractere(s string) string {
    if len(s) > 0 {
        return string(s[0])  // QUEBRA com caracteres multibyte!
    }
    return ""
}

// primeiroCaractere("São") → "S" ✅
// primeiroCaractere("中国") → "?" ❌ (byte inválido)

**✅ CORRETO:**

func primeiroCaractere(s string) string {
    runes := []rune(s)
    if len(runes) > 0 {
        return string(runes[0])
    }
    return ""
}

// primeiroCaractere("São") → "S" ✅
// primeiroCaractere("中国") → "中" ✅

---

#### **Antipadrão 3: Conversões Desnecessárias**

**❌ DESPERDÍCIO:**

func contarPalavras(texto string) int {
    bytes := []byte(texto)
    str := string(bytes)
    palavras := strings.Fields(str)
    return len(palavras)
}

// 2 conversões desnecessárias!

**✅ DIRETO:**

func contarPalavras(texto string) int {
    return len(strings.Fields(texto))
}

---

#### **Antipadrão 4: Não Usar strings.Builder.Grow()**

**❌ SUBÓTIMO:**

func construir(n int) string {
    var builder strings.Builder
    // Builder vai crescer várias vezes, realocando memória
    for i := 0; i < n; i++ {
        builder.WriteString("data")
    }
    return builder.String()
}

**✅ OTIMIZADO:**

func construir(n int) string {
    var builder strings.Builder
    builder.Grow(n * 4)  // Pré-aloca tamanho esperado
    for i := 0; i < n; i++ {
        builder.WriteString("data")
    }
    return builder.String()
}

**Diferença:** Evita realocações durante o crescimento

---

#### **Antipadrão 5: String Formatting Desnecessário**

**❌ LENTO:**

func construirPath(dir, file string) string {
    return fmt.Sprintf("%s/%s", dir, file)  // Usa reflection, lento
}

**✅ RÁPIDO:**

func construirPath(dir, file string) string {
    return dir + "/" + file  // Direto, mais rápido
}

**Ou ainda melhor para múltiplas concatenações:**

func construirPath(dir, file string) string {
    var builder strings.Builder
    builder.Grow(len(dir) + len(file) + 1)
    builder.WriteString(dir)
    builder.WriteByte('/')
    builder.WriteString(file)
    return builder.String()
}

---

### 🎯 **4. Quando Usar Cada Técnica**

#### **Tabela de Decisão**

| Operação | Melhor Escolha | Por Quê |
|---|---|---|
| Concatenar 2-3 strings | Operador + | Simples e legível |
| Concatenar em loop | strings.Builder | Performance |
| Formatar com variáveis | fmt.Sprintf | Legibilidade |
| Construir HTML/JSON | strings.Builder | Performance e controle |
| Manipular caracteres | []rune | Unicode-safe |
| Processar bytes brutos | []byte | Performance |
| Comparar case-insensitive | strings.EqualFold | Performance |
| Buscar/substituir | strings.Replace* | Otimizado |

---

### 🔬 **5. Casos Avançados**

#### **Caso 1: Parser de Expressões**

package main

import (
    "fmt"
    "strings"
)

type Token struct {
    Tipo  string
    Valor string
}

func tokenizar(expressao string) []Token {
    var tokens []Token
    var builder strings.Builder
    
    for _, r := range expressao {
        switch {
        case unicode.IsDigit(r):
            builder.WriteRune(r)
        case unicode.IsSpace(r):
            if builder.Len() > 0 {
                tokens = append(tokens, Token{"NUMBER", builder.String()})
                builder.Reset()
            }
        case r == '+' || r == '-' || r == '*' || r == '/':
            if builder.Len() > 0 {
                tokens = append(tokens, Token{"NUMBER", builder.String()})
                builder.Reset()
            }
            tokens = append(tokens, Token{"OPERATOR", string(r)})
        }
    }
    
    if builder.Len() > 0 {
        tokens = append(tokens, Token{"NUMBER", builder.String()})
    }
    
    return tokens
}

func main() {
    expressao := "10 + 20 * 3"
    tokens := tokenizar(expressao)
    
    for _, token := range tokens {
        fmt.Printf("%s: %s\n", token.Tipo, token.Valor)
    }
}

**Saída:**

NUMBER: 10
OPERATOR: +
NUMBER: 20
OPERATOR: *
NUMBER: 3

---

#### **Caso 2: Escape de HTML**

package main

import (
    "fmt"
    "strings"
)

func escaparHTML(texto string) string {
    var builder strings.Builder
    builder.Grow(len(texto) * 2)  // Estimativa conservadora
    
    for _, r := range texto {
        switch r {
        case '<':
            builder.WriteString("&lt;")
        case '>':
            builder.WriteString("&gt;")
        case '&':
            builder.WriteString("&amp;")
        case '"':
            builder.WriteString("&quot;")
        case '\'':
            builder.WriteString("&#39;")
        default:
            builder.WriteRune(r)
        }
    }
    
    return builder.String()
}

func main() {
    perigoso := `<script>alert("XSS")</script>`
    seguro := escaparHTML(perigoso)
    
    fmt.Println("Original:", perigoso)
    fmt.Println("Escapado:", seguro)
}

**Saída:**

Original: <script>alert("XSS")</script>
Escapado: &lt;script&gt;alert(&quot;XSS&quot;)&lt;/script&gt;

---

#### **Caso 3: Truncar String Preservando Palavras**

package main

import (
    "fmt"
    "strings"
    "unicode/utf8"
)

func truncar(texto string, maxLen int) string {
    runes := []rune(texto)
    
    // Se já é menor, retorna como está
    if len(runes) <= maxLen {
        return texto
    }
    
    // Trunca no tamanho máximo
    truncado := string(runes[:maxLen])
    
    // Procura o último espaço para não cortar palavra
    ultimoEspaco := strings.LastIndex(truncado, " ")
    if ultimoEspaco > 0 {
        truncado = truncado[:ultimoEspaco]
    }
    
    return truncado + "..."
}

func main() {
    texto := "Go é uma linguagem de programação incrível e poderosa"
    
    fmt.Println("Original:", texto)
    fmt.Println("20 chars:", truncar(texto, 20))
    fmt.Println("30 chars:", truncar(texto, 30))
    fmt.Println("50 chars:", truncar(texto, 50))
}

**Saída:**

Original: Go é uma linguagem de programação incrível e poderosa
20 chars: Go é uma linguagem...
30 chars: Go é uma linguagem de...
50 chars: Go é uma linguagem de programação incrível e...

---

### 📊 **6. Checklist Final de Performance**

Antes de finalizar código com strings:

- [ ] **Há concatenação em loop?**
  → Use strings.Builder

- [ ] **Está usando fmt.Sprintf desnecessariamente?**
  → Considere operador + ou Builder

- [ ] **Está manipulando caracteres individuais?**
  → Use []rune se Unicode, []byte se ASCII

- [ ] **Está criando substrings de strings grandes?**
  → Force cópia para evitar memory leak

- [ ] **Está pré-alocando Builder?**
  → Use Grow() com tamanho estimado

- [ ] **Está validando input do usuário?**
  → Sempre valide e sanitize

---

### 🎓 **7. Perguntas Frequentes Avançadas**

#### **Q1: Por que strings.Builder não tem método para obter tamanho atual?**

**R:** Tem sim! Use `builder.Len()`:

var builder strings.Builder
builder.WriteString("Hello")
fmt.Println(builder.Len())  // 5

---

#### **Q2: Posso resetar um strings.Builder para reutilizar?**

**R:** Sim! Use `builder.Reset()`:

var builder strings.Builder
builder.WriteString("primeira string")
resultado1 := builder.String()

builder.Reset()  // Limpa o buffer

builder.WriteString("segunda string")
resultado2 := builder.String()

---

#### **Q3: Como converter string para int e vice-versa?**

**R:** Use o package `strconv`:

import "strconv"

// String → Int
num, err := strconv.Atoi("42")

// Int → String
str := strconv.Itoa(42)

// Mais controle
num64, err := strconv.ParseInt("42", 10, 64)
str2 := strconv.FormatInt(42, 10)

---

#### **Q4: Como iterar sobre string sem range?**

**R:** Use utf8.DecodeRuneInString:

import "unicode/utf8"

texto := "Café"
for i := 0; i < len(texto); {
    r, tamanho := utf8.DecodeRuneInString(texto[i:])
    fmt.Printf("%c ", r)
    i += tamanho
}

---

### 🎯 **8. Resumo Final: As 7 Regras de Ouro**

1. **Strings são imutáveis** → Converta para []byte ou []rune para modificar
2. **len() ≠ número de caracteres** → Use utf8.RuneCountInString() para Unicode
3. **Use strings.Builder em loops** → 100x mais rápido que +
4. **range itera runes, não bytes** → Seguro para Unicode
5. **Substrings compartilham memória** → Force cópia se necessário
6. **Valide sempre input de usuário** → Segurança e robustez
7. **Prefira funções do package strings** → Otimizadas e testadas

---

### 📚 **9. Recursos Adicionais**

**Packages importantes:**
- `strings`: Operações gerais
- `strconv`: Conversões string ↔ números
- `unicode`: Classificação de caracteres
- `unicode/utf8`: Operações UTF-8
- `text/template`: Templates avançados
- `regexp`: Expressões regulares

**Documentação:**
- Go Blog: "Strings, bytes, runes and characters in Go"
- Effective Go: Strings section

---

## 📋 **Aula 7 Concluída!**

Você agora domina:
- ✅ Estrutura interna de strings
- ✅ Diferença entre bytes e runes
- ✅ Imutabilidade e suas implicações
- ✅ Operações eficientes (strings.Builder)
- ✅ Manipulação Unicode-safe
- ✅ Boas práticas e antipadrões

**Próxima aula:** **Aula 8: Maps (Mapas/Dicionários)** onde exploraremos estruturas de dados chave-valor, hash tables, e operações eficientes de busca.

**Pronto para continuar para a próxima aula ou prefere fazer uma pausa?**