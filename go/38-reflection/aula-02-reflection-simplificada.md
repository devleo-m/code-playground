# Módulo 38: Reflection em Go
## Aula 2 - Simplificada: Entendendo Reflection

Agora vamos entender esses conceitos de forma mais simples, usando analogias do nosso dia a dia!

---

## 1. O Que É Reflection? O Espelho Mágico

Imagine que você tem um **espelho mágico** que pode:

- **Ver o que está dentro** de uma caixa sem abrir
- **Descobrir o tipo** de coisa que está na caixa
- **Modificar** o que está dentro (se tiver permissão)
- **Chamar funções** sem saber o nome delas antes

**Reflection** é esse "espelho mágico" para código! Ele permite que você:
- Veja o tipo de uma variável em **tempo de execução** (não em tempo de compilação)
- Examine campos de structs sem conhecer os nomes antes
- Chame métodos dinamicamente
- Crie coisas novas baseado em tipos que você só descobre quando o programa roda

**Analogia**: É como ter um "detetive" que investiga seu código enquanto ele está rodando!

---

## 2. Type vs Value: O Que É vs O Que Tem

### Type: O Que É a Coisa

**Type** é como perguntar: "O que é isso?"

```go
var x int = 42

t := reflect.TypeOf(x)
// t diz: "Isso é um int"
```

**Analogia**: É como perguntar "O que é isso?" e receber a resposta: "É um carro", "É uma casa", "É um número inteiro".

### Value: O Que Tem Dentro

**Value** é como perguntar: "Qual é o valor disso?"

```go
var x int = 42

v := reflect.ValueOf(x)
// v diz: "O valor é 42"
```

**Analogia**: É como perguntar "Qual é o valor?" e receber: "O carro é vermelho", "A casa tem 3 quartos", "O número é 42".

### Exemplo Prático

```go
// Type: "Isso é um int"
// Value: "O valor é 42"

var x int = 42
t := reflect.TypeOf(x)   // Type: int
v := reflect.ValueOf(x)   // Value: 42
```

**Analogia**: É como ter uma caixa:
- **Type** diz: "É uma caixa de sapatos"
- **Value** diz: "Tem um par de tênis dentro"

---

## 3. Inspecionar Tipos: O Detetive Investigando

### Tipos Básicos

```go
func investigar(coisa interface{}) {
    t := reflect.TypeOf(coisa)
    fmt.Println("Tipo:", t)  // "É um int", "É um string", etc.
}
```

**Analogia**: É como um detetive olhando para algo e dizendo: "Isso é um carro", "Isso é uma casa", "Isso é um número".

### Structs: Caixas com Várias Coisas

```go
type Pessoa struct {
    Nome string
    Idade int
}

func investigarStruct(coisa interface{}) {
    t := reflect.TypeOf(coisa)
    
    // Quantos campos tem?
    fmt.Println("Tem", t.NumField(), "campos")
    
    // Quais são os campos?
    for i := 0; i < t.NumField(); i++ {
        campo := t.Field(i)
        fmt.Println("Campo:", campo.Name)  // "Nome", "Idade"
    }
}
```

**Analogia**: É como abrir uma caixa e listar tudo que tem dentro:
- "Esta caixa tem 2 coisas"
- "Coisa 1: Nome (é uma string)"
- "Coisa 2: Idade (é um int)"

### Tags: Etiquetas Especiais

```go
type Usuario struct {
    Nome  string `json:"name" validate:"required"`
    Email string `json:"email" validate:"required"`
}

func verEtiquetas(coisa interface{}) {
    t := reflect.TypeOf(coisa)
    
    for i := 0; i < t.NumField(); i++ {
        campo := t.Field(i)
        etiquetaJSON := campo.Tag.Get("json")
        fmt.Println(campo.Name, "tem etiqueta json:", etiquetaJSON)
    }
}
```

**Analogia**: É como ter etiquetas em uma mala:
- A etiqueta `json:"name"` diz: "Quando converter para JSON, use 'name'"
- A etiqueta `validate:"required"` diz: "Este campo é obrigatório"

---

## 4. Inspecionar Valores: Ver O Que Tem Dentro

### Valores Básicos

```go
func verValor(coisa interface{}) {
    v := reflect.ValueOf(coisa)
    
    if v.Kind() == reflect.Int {
        fmt.Println("Valor inteiro:", v.Int())
    } else if v.Kind() == reflect.String {
        fmt.Println("Valor string:", v.String())
    }
}
```

**Analogia**: É como abrir a caixa e ver o que tem dentro:
- "A caixa tem o número 42"
- "A caixa tem a palavra 'hello'"

### Structs: Ver Valores dos Campos

```go
type Pessoa struct {
    Nome  string
    Idade int
}

func verValoresStruct(coisa interface{}) {
    v := reflect.ValueOf(coisa)
    t := reflect.TypeOf(coisa)
    
    for i := 0; i < v.NumField(); i++ {
        campo := t.Field(i)
        valor := v.Field(i)
        fmt.Printf("%s = %v\n", campo.Name, valor.Interface())
    }
}

func main() {
    p := Pessoa{Nome: "João", Idade: 30}
    verValoresStruct(p)
    // Nome = João
    // Idade = 30
}
```

**Analogia**: É como abrir uma caixa e listar tudo que tem:
- "Nome = João"
- "Idade = 30"

---

## 5. Modificar Valores: Mudar O Que Tem Dentro

### ⚠️ Regra Importante: Precisa de "Chave"

Para modificar algo, você precisa de uma **"chave"** (pointer):

```go
// ❌ Não funciona: Sem "chave"
func modificarRuim(coisa interface{}) {
    v := reflect.ValueOf(coisa)
    v.SetInt(100)  // ERRO! Não tem "chave"
}

// ✅ Funciona: Com "chave" (pointer)
func modificarBom(coisa interface{}) {
    v := reflect.ValueOf(coisa)
    v = v.Elem()  // Usar a "chave" para abrir
    v.SetInt(100)  // Agora funciona!
}

func main() {
    x := 42
    modificarBom(&x)  // Passar "chave" (pointer)
    fmt.Println(x)  // 100
}
```

**Analogia**: É como modificar algo dentro de uma caixa:
- **Sem pointer**: Você só tem uma **cópia** da caixa, não pode modificar
- **Com pointer**: Você tem a **chave** da caixa, pode abrir e modificar

### Modificando Campos de Struct

```go
type Pessoa struct {
    Nome  string
    Idade int
}

func modificarStruct(coisa interface{}) {
    v := reflect.ValueOf(coisa)
    v = v.Elem()  // Usar "chave"
    
    // Modificar campo Nome
    campoNome := v.FieldByName("Nome")
    campoNome.SetString("Maria")
    
    // Modificar campo Idade
    campoIdade := v.FieldByName("Idade")
    campoIdade.SetInt(25)
}

func main() {
    p := Pessoa{Nome: "João", Idade: 30}
    fmt.Println("Antes:", p)  // {João 30}
    
    modificarStruct(&p)  // Passar "chave"
    fmt.Println("Depois:", p)  // {Maria 25}
}
```

**Analogia**: É como abrir uma caixa com a chave e trocar o que tem dentro:
- Antes: "Nome = João, Idade = 30"
- Depois: "Nome = Maria, Idade = 25"

---

## 6. Chamar Métodos Dinamicamente: O Assistente Mágico

### Chamar Método Sem Saber o Nome Antes

```go
type Calculadora struct {
    resultado int
}

func (c *Calculadora) Somar(x int) {
    c.resultado += x
}

func (c *Calculadora) Resultado() int {
    return c.resultado
}

func chamarMetodo(coisa interface{}, nomeMetodo string, argumentos ...interface{}) {
    v := reflect.ValueOf(coisa)
    
    // Procurar método pelo nome
    metodo := v.MethodByName(nomeMetodo)
    
    // Converter argumentos
    args := make([]reflect.Value, len(argumentos))
    for i, arg := range argumentos {
        args[i] = reflect.ValueOf(arg)
    }
    
    // Chamar método
    metodo.Call(args)
}

func main() {
    calc := &Calculadora{resultado: 10}
    
    // Chamar "Somar" sem saber o nome em compile-time!
    chamarMetodo(calc, "Somar", 5)
    fmt.Println(calc.Resultado())  // 15
}
```

**Analogia**: É como ter um assistente que:
- Você diz: "Chame o método 'Somar' com o número 5"
- O assistente encontra o método e chama
- Você não precisa saber o nome do método antes de escrever o código!

---

## 7. Casos de Uso: Quando Usar o Espelho Mágico

### Caso 1: Converter para JSON

```go
type Usuario struct {
    Nome  string `json:"name"`
    Email string `json:"email"`
}

// Reflection permite converter sem saber os campos antes!
func paraJSON(coisa interface{}) string {
    // Usa reflection para descobrir campos e valores
    // e criar JSON automaticamente
}
```

**Analogia**: É como ter um tradutor automático que:
- Olha para uma struct
- Descobre todos os campos
- Converte para JSON automaticamente
- Você não precisa escrever código para cada struct!

### Caso 2: Validar Formulários

```go
type Formulario struct {
    Nome  string `validate:"required"`
    Email string `validate:"required"`
}

// Reflection permite validar sem escrever código para cada campo!
func validar(coisa interface{}) []string {
    // Usa reflection para descobrir campos
    // Verifica se campos "required" estão preenchidos
    // Retorna lista de erros
}
```

**Analogia**: É como ter um fiscal que:
- Olha para um formulário
- Descobre quais campos são obrigatórios (pelas etiquetas)
- Verifica se estão preenchidos
- Você não precisa escrever código de validação para cada campo!

---

## 8. Limitações: O Que o Espelho Mágico NÃO Pode Fazer

### ⚠️ É Mais Lento

**Reflection é mais lento** que código normal:

```go
// Código normal (rápido)
func somar(a, b int) int {
    return a + b
}

// Código com reflection (lento)
func somarReflection(a, b interface{}) interface{} {
    // Precisa descobrir tipos, converter valores, etc.
    // Muito mais lento!
}
```

**Analogia**: É como a diferença entre:
- **Código normal**: Você sabe exatamente o que fazer, faz direto
- **Reflection**: Você precisa primeiro investigar o que é, depois fazer
- A investigação leva tempo!

### ⚠️ Erros Só Aparecem Quando Roda

```go
// Isso compila, mas pode dar erro quando roda!
func chamarMetodo(coisa interface{}, nome string) {
    metodo := coisa.MethodByName(nome)  // E se não existir?
    metodo.Call(nil)  // PANIC se método não existir!
}
```

**Analogia**: É como tentar abrir uma porta:
- **Código normal**: O compilador verifica se a chave existe antes
- **Reflection**: Você só descobre se a chave funciona quando tenta abrir
- Se a chave não funcionar, o programa quebra!

### ⚠️ Código Mais Difícil de Entender

Código com reflection é mais difícil de:
- Entender o que faz
- Manter
- Debugar
- Testar

**Analogia**: É como usar uma ferramenta muito complexa:
- Funciona, mas é difícil de usar
- Precisa de mais conhecimento
- Mais fácil de quebrar

---

## 9. Quando Usar Reflection?

### ✅ Use Quando:

1. **Não sabe os tipos antes**: Quando tipos só são conhecidos em runtime
2. **Bibliotecas genéricas**: Como JSON, ORM, validação
3. **Ferramentas**: Code generation, debugging tools

**Analogia**: Use quando você precisa do "espelho mágico" para descobrir coisas que não sabe antes.

### ❌ NÃO Use Quando:

1. **Sabe os tipos**: Quando você já sabe os tipos em compile-time
2. **Performance importante**: Quando código precisa ser muito rápido
3. **Código simples**: Quando código normal é suficiente

**Analogia**: Não use quando você já sabe o que tem na caixa e pode abrir diretamente!

---

## Resumo com Analogias

1. **Reflection**: É um "espelho mágico" que investiga código em runtime
2. **Type**: Diz "o que é" (int, string, struct)
3. **Value**: Diz "qual é o valor" (42, "hello", {Nome: "João"})
4. **Modificar**: Precisa de "chave" (pointer) para modificar
5. **Chamar métodos**: Pode chamar métodos sem saber o nome antes
6. **Casos de uso**: JSON, validação, ORM
7. **Limitações**: Mais lento, erros em runtime, mais complexo
8. **Quando usar**: Quando não sabe tipos antes, bibliotecas genéricas

---

## Perguntas para Pensar

1. **Por que reflection é mais lento?**
   - Pense: O que precisa acontecer que código normal não precisa?

2. **Quando faz sentido usar reflection?**
   - Pense: Em que situações você realmente não sabe os tipos antes?

3. **Por que precisa de pointer para modificar?**
   - Pense: Qual é a diferença entre ter uma cópia vs ter a "chave" original?

4. **Quando NÃO usar reflection?**
   - Pense: Em que situações código normal é melhor?

---

**Lembre-se**: Reflection é como um "espelho mágico" - poderoso, mas use com cuidado! Prefira código normal quando possível. 🪞✨

