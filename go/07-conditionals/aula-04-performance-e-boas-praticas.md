# Módulo 7: Conditionals em Go

## Aula 4 - Performance, Boas Práticas e O Que Deve/Não Deve Ser Utilizado

Agora que você entende conditionals, vamos falar sobre **como usá-los de forma eficiente e profissional**. Esta é a parte que separa desenvolvedores iniciantes de desenvolvedores experientes.

---

## 1. Performance: Ordem das Condições Importa

### ⚠️ Mito: "Ordem Não Importa"

**Verdade**: A ordem das condições em `if-else if` **importa muito** para performance.

### Ordem Eficiente: Condições Mais Prováveis Primeiro

**✅ BOM:**

```go
func Processar(dados interface{}) {
    // Condição mais provável primeiro
    if dados == nil {
        return
    }

    // Segunda mais provável
    if str, ok := dados.(string); ok {
        processarString(str)
        return
    }

    // Menos provável por último
    if num, ok := dados.(int); ok {
        processarNumero(num)
    }
}
```

**❌ RUIM:**

```go
func Processar(dados interface{}) {
    // Condição menos provável primeiro - avalia todas as outras desnecessariamente
    if num, ok := dados.(int); ok {
        processarNumero(num)
        return
    }

    // Mais provável por último - só chega aqui depois de verificar tudo
    if dados == nil {
        return
    }
}
```

**Por quê?**

- Go avalia condições **de cima para baixo**
- Para na **primeira condição verdadeira**
- Se condições mais comuns estão no final, avalia todas as anteriores **desnecessariamente**

### Regra de Ouro

**Coloque condições mais prováveis primeiro** para melhorar performance.

---

## 2. Performance: Short-Circuit Evaluation

Go usa **avaliação de curto-circuito** (short-circuit evaluation) nos operadores lógicos.

### Como Funciona

**Com `&&` (E):**

```go
if condicao1 && condicao2 {
    // Se condicao1 for false, condicao2 NÃO é avaliada
}
```

**Com `||` (OU):**

```go
if condicao1 || condicao2 {
    // Se condicao1 for true, condicao2 NÃO é avaliada
}
```

### Exemplo: Função Custosa

**✅ BOM:**

```go
// Verificar condição barata primeiro
if valor != nil && valor.Processar() {  // Processar() só é chamado se valor != nil
    // ...
}
```

**❌ RUIM:**

```go
// Se inverter, sempre chama Processar() mesmo quando valor é nil
if valor.Processar() && valor != nil {  // ERRO: pode causar panic!
    // ...
}
```

### Regra de Ouro

**Coloque condições mais baratas/seguras primeiro** para aproveitar short-circuit.

---

## 3. Performance: switch vs if-else if

### Quando São Equivalentes

Para **comparações de igualdade simples**, `switch` e `if-else if` têm performance similar:

```go
// Performance similar
if valor == 1 {
    // ...
} else if valor == 2 {
    // ...
}

switch valor {
case 1:
    // ...
case 2:
    // ...
}
```

### Quando switch Pode Ser Mais Rápido

**Com muitos casos**, `switch` pode ser otimizado pelo compilador:

```go
// switch pode ser otimizado para jump table
switch valor {
case 1:
    // ...
case 2:
    // ...
case 3:
    // ...
// ... muitos casos
}
```

**Por quê?**

- Compilador pode criar uma **tabela de saltos** (jump table)
- Acesso direto ao código correto
- Mais eficiente que múltiplas comparações sequenciais

### Regra de Ouro

**Use `switch` quando tiver muitos casos** (5+) - pode ser mais rápido e mais legível.

---

## 4. Boas Práticas: O Que Deve Ser Utilizado

### Prática 1: Sempre Valide Entradas

**✅ BOM:**

```go
func Dividir(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("divisão por zero")
    }
    return a / b, nil
}
```

**❌ RUIM:**

```go
func Dividir(a, b float64) float64 {
    return a / b  // Pode causar panic ou resultado infinito!
}
```

**Por quê?** Validação previne erros e torna o código mais robusto.

---

### Prática 2: Use Inicialização no `if` Quando Apropriado

**✅ BOM:**

```go
// Variável existe apenas onde é necessária
if resultado, err := processar(); err != nil {
    return err
}
// resultado não polui o escopo externo
```

**❌ RUIM:**

```go
// Variável no escopo externo desnecessariamente
resultado, err := processar()
if err != nil {
    return err
}
// resultado ainda existe aqui, mesmo que não seja usado
```

**Por quê?** Limita o escopo das variáveis, tornando o código mais limpo.

---

### Prática 3: Use `switch` para Múltiplas Comparações de Igualdade

**✅ BOM:**

```go
switch dia {
case "segunda", "terça", "quarta", "quinta", "sexta":
    fmt.Println("Dia útil")
case "sábado", "domingo":
    fmt.Println("Fim de semana")
}
```

**❌ RUIM:**

```go
if dia == "segunda" || dia == "terça" || dia == "quarta" || dia == "quinta" || dia == "sexta" {
    fmt.Println("Dia útil")
} else if dia == "sábado" || dia == "domingo" {
    fmt.Println("Fim de semana")
}
```

**Por quê?** `switch` é mais legível e mais fácil de manter.

---

### Prática 4: Sempre Trate Erros

**✅ BOM:**

```go
resultado, err := operacaoPerigosa()
if err != nil {
    return fmt.Errorf("falha na operação: %w", err)
}
// Usar resultado
```

**❌ RUIM:**

```go
resultado, err := operacaoPerigosa()
// Ignorar erro - pode causar bugs difíceis de encontrar
fmt.Println(resultado)
```

**Por quê?** Erros não tratados causam bugs silenciosos e difíceis de debugar.

---

### Prática 5: Use `default` em `switch` Quando Apropriado

**✅ BOM:**

```go
switch tipo {
case "admin":
    return "Administrador"
case "usuario":
    return "Usuário"
default:
    return "Tipo desconhecido"  // Trata casos inesperados
}
```

**❌ RUIM:**

```go
switch tipo {
case "admin":
    return "Administrador"
case "usuario":
    return "Usuário"
// Sem default - o que acontece se tipo for "moderador"?
}
```

**Por quê?** `default` garante que todos os casos sejam tratados.

---

### Prática 6: Evite Conditionals Aninhados Profundos

**✅ BOM:**

```go
// Usar operadores lógicos
if idade >= 18 && temCarteira && saldo > 0 {
    permitirAcesso()
}
```

**❌ RUIM:**

```go
// Aninhamento profundo - difícil de ler
if idade >= 18 {
    if temCarteira {
        if saldo > 0 {
            permitirAcesso()
        }
    }
}
```

**Por quê?** Aninhamento profundo torna o código difícil de ler e manter.

**Regra:** Evite mais de 2-3 níveis de aninhamento.

---

### Prática 7: Extraia Condições Complexas para Variáveis

**✅ BOM:**

```go
// Condição complexa em variável com nome descritivo
podeAcessar := idade >= 18 && temCarteira && !banido && saldo > 0

if podeAcessar {
    permitirAcesso()
}
```

**❌ RUIM:**

```go
// Condição complexa inline - difícil de entender
if idade >= 18 && temCarteira && !banido && saldo > 0 {
    permitirAcesso()
}
```

**Por quê?** Variáveis com nomes descritivos tornam o código auto-documentado.

---

### Prática 8: Use Early Returns

**✅ BOM:**

```go
func Processar(dados string) error {
    if dados == "" {
        return fmt.Errorf("dados vazios")
    }

    if len(dados) < 5 {
        return fmt.Errorf("dados muito curtos")
    }

    // Código principal aqui (sem aninhamento)
    processarDados(dados)
    return nil
}
```

**❌ RUIM:**

```go
func Processar(dados string) error {
    if dados != "" {
        if len(dados) >= 5 {
            // Código principal aqui (aninhado)
            processarDados(dados)
            return nil
        } else {
            return fmt.Errorf("dados muito curtos")
        }
    } else {
        return fmt.Errorf("dados vazios")
    }
}
```

**Por quê?** Early returns reduzem aninhamento e tornam o código mais linear e fácil de ler.

---

## 5. O Que NÃO Deve Ser Feito

### ❌ Erro 1: Comparações Desnecessárias

**RUIM:**

```go
if condicao == true {
    // ...
}

if condicao == false {
    // ...
}
```

**BOM:**

```go
if condicao {
    // ...
}

if !condicao {
    // ...
}
```

**Por quê?** `condicao` já é booleano - não precisa comparar com `true` ou `false`.

---

### ❌ Erro 2: Ignorar Erros

**RUIM:**

```go
resultado, _ := operacaoPerigosa()  // Ignora erro
fmt.Println(resultado)
```

**BOM:**

```go
resultado, err := operacaoPerigosa()
if err != nil {
    return fmt.Errorf("erro: %w", err)
}
fmt.Println(resultado)
```

**Por quê?** Erros ignorados causam bugs difíceis de encontrar.

---

### ❌ Erro 3: Conditionals Redundantes

**RUIM:**

```go
if x > 0 {
    return true
} else {
    return false
}
```

**BOM:**

```go
return x > 0
```

**Por quê?** A condição já retorna um booleano - não precisa de `if-else`.

---

### ❌ Erro 4: Usar `switch` Quando `if` é Mais Apropriado

**RUIM:**

```go
// switch para condição simples
switch {
case x > 10:
    fmt.Println("Maior que 10")
}
```

**BOM:**

```go
// if é mais simples e direto
if x > 10 {
    fmt.Println("Maior que 10")
}
```

**Por quê?** `if` é mais simples para condições únicas ou simples.

---

### ❌ Erro 5: `fallthrough` Desnecessário

**RUIM:**

```go
switch valor {
case 1:
    fmt.Println("Um")
    fallthrough  // Desnecessário - apenas um case
case 2:
    fmt.Println("Dois")
}
```

**BOM:**

```go
switch valor {
case 1:
    fmt.Println("Um")
case 2:
    fmt.Println("Dois")
}
```

**Por quê?** `fallthrough` raramente é necessário e pode confundir.

**Quando usar `fallthrough`:**

- ✅ Quando múltiplos casos devem executar o mesmo código
- ✅ Quando implementa lógica específica que requer continuidade
- ⚠️ Use com moderação e documente bem

---

### ❌ Erro 6: Condições Muito Complexas

**RUIM:**

```go
// Condição muito complexa - difícil de entender
if (x > 0 && y < 10) || (z == 5 && w != 3) && !(a == b || c > d) {
    // ...
}
```

**BOM:**

```go
// Extrair para variáveis com nomes descritivos
condicao1 := x > 0 && y < 10
condicao2 := z == 5 && w != 3
condicao3 := a == b || c > d

podeExecutar := (condicao1 || condicao2) && !condicao3

if podeExecutar {
    // ...
}
```

**Por quê?** Condições complexas são difíceis de entender e manter.

---

### ❌ Erro 7: Não Usar `default` Quando Deve

**RUIM:**

```go
switch tipo {
case "admin":
    return "Administrador"
case "usuario":
    return "Usuário"
// Sem default - o que acontece com tipos inesperados?
}
```

**BOM:**

```go
switch tipo {
case "admin":
    return "Administrador"
case "usuario":
    return "Usuário"
default:
    return "Tipo desconhecido"  // Trata casos inesperados
}
```

**Por quê?** Sem `default`, casos inesperados podem causar bugs silenciosos.

---

### ❌ Erro 8: Aninhamento Excessivo

**RUIM:**

```go
// 4 níveis de aninhamento - muito difícil de ler
if condicao1 {
    if condicao2 {
        if condicao3 {
            if condicao4 {
                fazerAlgo()
            }
        }
    }
}
```

**BOM:**

```go
// Usar operadores lógicos ou early returns
if !condicao1 || !condicao2 || !condicao3 || !condicao4 {
    return
}

fazerAlgo()
```

**Por quê?** Aninhamento profundo torna o código difícil de ler e manter.

---

## 6. Padrões de Design com Conditionals

### Padrão 1: Guard Clauses (Early Returns)

**Padrão comum em Go:**

```go
func Processar(dados string) error {
    // Guard clauses - retornar cedo se condições não forem atendidas
    if dados == "" {
        return fmt.Errorf("dados vazios")
    }

    if len(dados) < 5 {
        return fmt.Errorf("dados muito curtos")
    }

    // Código principal (sem aninhamento)
    processarDados(dados)
    return nil
}
```

**Vantagens:**

- ✅ Reduz aninhamento
- ✅ Código mais linear
- ✅ Mais fácil de ler

---

### Padrão 2: Strategy Pattern com switch

**Para diferentes estratégias baseadas em tipo:**

```go
type Processador interface {
    Processar() error
}

func ProcessarPorTipo(tipo string, dados interface{}) error {
    switch tipo {
    case "json":
        return processarJSON(dados)
    case "xml":
        return processarXML(dados)
    case "csv":
        return processarCSV(dados)
    default:
        return fmt.Errorf("tipo não suportado: %s", tipo)
    }
}
```

---

### Padrão 3: Validação em Camadas

**Validar em ordem de importância:**

```go
func ValidarUsuario(usuario Usuario) error {
    // Validações mais críticas primeiro
    if usuario.Email == "" {
        return fmt.Errorf("email é obrigatório")
    }

    if !strings.Contains(usuario.Email, "@") {
        return fmt.Errorf("email inválido")
    }

    if len(usuario.Senha) < 8 {
        return fmt.Errorf("senha muito curta")
    }

    // Validações menos críticas depois
    if usuario.Nome == "" {
        return fmt.Errorf("nome é obrigatório")
    }

    return nil
}
```

---

## 7. Performance: Benchmarks e Otimizações

### Quando Otimizar

**Não otimize prematuramente!**

- ✅ Primeiro, escreva código **legível e correto**
- ✅ Depois, **meça** a performance
- ✅ Só então, **otimize** se necessário

### Exemplo: switch vs if-else if

Para **poucos casos** (< 5), diferença é negligenciável:

```go
// Performance similar
if valor == 1 {
    // ...
} else if valor == 2 {
    // ...
}

switch valor {
case 1:
    // ...
case 2:
    // ...
}
```

Para **muitos casos** (10+), `switch` pode ser mais rápido:

```go
// switch pode ser otimizado pelo compilador
switch valor {
case 1:
    // ...
case 2:
    // ...
// ... muitos casos
}
```

---

## 8. Boas Práticas: Nomenclatura

### ✅ BOM:

```go
// Variáveis booleanas com nomes descritivos
podeAcessar := idade >= 18 && temCarteira
estaAtivo := usuario.Status == "ativo"
temPermissao := verificarPermissao(usuario)
```

### ❌ RUIM:

```go
// Nomes genéricos ou confusos
flag := idade >= 18 && temCarteira
x := usuario.Status == "ativo"
temp := verificarPermissao(usuario)
```

**Por quê?** Nomes descritivos tornam o código auto-documentado.

---

## 9. Resumo: Regras de Ouro

1. **Valide sempre** - Previne erros e bugs
2. **Use early returns** - Reduz aninhamento
3. **Ordene condições** - Mais prováveis primeiro
4. **Use `switch` para igualdade** - Mais legível
5. **Use `if` para condições complexas** - Mais flexível
6. **Extraia condições complexas** - Variáveis com nomes descritivos
7. **Evite aninhamento profundo** - Máximo 2-3 níveis
8. **Sempre trate erros** - Nunca ignore
9. **Use `default` em `switch`** - Trata casos inesperados
10. **Aproveite short-circuit** - Condições baratas primeiro

---

## 10. Comparação: Com vs. Sem Boas Práticas

### Cenário: Validação de Usuário

**SEM boas práticas:**

```go
func CriarUsuario(usuario Usuario) error {
    if usuario.Email != "" {
        if strings.Contains(usuario.Email, "@") {
            if len(usuario.Senha) >= 8 {
                if usuario.Nome != "" {
                    // Criar usuário (código aninhado)
                    criarUsuario(usuario)
                    return nil
                } else {
                    return fmt.Errorf("nome é obrigatório")
                }
            } else {
                return fmt.Errorf("senha muito curta")
            }
        } else {
            return fmt.Errorf("email inválido")
        }
    } else {
        return fmt.Errorf("email é obrigatório")
    }
}
```

**COM boas práticas:**

```go
func CriarUsuario(usuario Usuario) error {
    // Early returns - sem aninhamento
    if usuario.Email == "" {
        return fmt.Errorf("email é obrigatório")
    }

    if !strings.Contains(usuario.Email, "@") {
        return fmt.Errorf("email inválido")
    }

    if len(usuario.Senha) < 8 {
        return fmt.Errorf("senha muito curta")
    }

    if usuario.Nome == "" {
        return fmt.Errorf("nome é obrigatório")
    }

    // Código principal (linear, sem aninhamento)
    criarUsuario(usuario)
    return nil
}
```

**Diferenças:**

- ✅ Código mais linear
- ✅ Mais fácil de ler
- ✅ Mais fácil de manter
- ✅ Menos propenso a erros

---

## Conclusão

Conditionals são fundamentais, mas usá-los corretamente requer:

1. **Entender performance** - Ordem das condições, short-circuit
2. **Validação adequada** - Sempre validar entradas
3. **Tratamento de erros** - Nunca ignorar erros
4. **Código limpo** - Evitar aninhamento, usar early returns
5. **Escolher a ferramenta certa** - `if` vs `switch`

Lembre-se: **Código é lido muito mais vezes do que escrito**. Conditionals bem escritos tornam o código mais fácil de entender, manter e debugar.

---

Na próxima etapa, você completará os exercícios e eu analisarei seu desempenho. Boa sorte!

