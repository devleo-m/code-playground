# Módulo 6: Structs em Go

## Aula 4 - Performance, Boas Práticas e O Que Deve/Não Deve Ser Utilizado

Agora que você entende structs, vamos falar sobre **como usá-las de forma eficiente e profissional**. Esta é a parte que separa desenvolvedores iniciantes de desenvolvedores experientes.

---

## 1. Performance: Passar por Valor vs. Ponteiro

### ⚠️ Mito: "Sempre Use Ponteiros"

**Verdade**: Nem sempre ponteiros são mais rápidos. Depende do tamanho da struct.

### Quando Passar por Valor

**✅ BOM para structs pequenas:**

```go
type Ponto struct {
    X int
    Y int
}

// Struct pequena (2 ints = 16 bytes) - passar por valor é eficiente
func MoverPonto(p Ponto, dx, dy int) Ponto {
    p.X += dx
    p.Y += dy
    return p
}
```

**Por quê?**

- Structs pequenas cabem em registradores da CPU
- Evita indireção (dereferencing)
- Mais rápido para structs até ~100 bytes

### Quando Passar por Ponteiro

**✅ BOM para structs grandes ou quando precisa modificar:**

```go
type Pessoa struct {
    Nome      string
    Email     string
    Endereco  Endereco
    Telefone  string
    CPF       string
    // ... muitos outros campos
}

// Struct grande - passar por ponteiro é mais eficiente
func ModificarPessoa(p *Pessoa) {
    p.Nome = "Novo Nome"
}
```

**Por quê?**

- Evita copiar muitos bytes
- Ponteiro tem tamanho fixo (8 bytes em 64-bit)
- Mais eficiente para structs grandes

### Regra de Ouro

**Passe por valor quando:**

- Struct é pequena (< 100 bytes aproximadamente)
- Não precisa modificar
- Quer imutabilidade

**Passe por ponteiro quando:**

- Struct é grande (> 100 bytes)
- Precisa modificar
- Quer eficiência em memória

---

## 2. Performance: Métodos com Receptor por Valor vs. Ponteiro

### Receptor por Valor

```go
type Contador struct {
    valor int
}

// Receptor por valor - cria cópia
func (c Contador) Incrementar() {
    c.valor++  // Modifica apenas a cópia!
}

func main() {
    contador := Contador{valor: 0}
    contador.Incrementar()
    fmt.Println(contador.valor) // Ainda é 0!
}
```

**Problema:** Não modifica o original.

### Receptor por Ponteiro (CORRETO para Modificação)

```go
// Receptor por ponteiro - modifica original
func (c *Contador) Incrementar() {
    c.valor++  // Modifica o original!
}

func main() {
    contador := Contador{valor: 0}
    contador.Incrementar()  // Go permite chamar assim
    fmt.Println(contador.valor) // Agora é 1!
}
```

### Convenção Go

**Use ponteiro quando:**

- Método modifica a struct
- Struct é grande
- Quer consistência (todos os métodos usam ponteiro)

**Use valor quando:**

- Método apenas lê (mas ponteiro também funciona)
- Struct é muito pequena
- Quer imutabilidade

**Recomendação:** Para consistência, use ponteiro na maioria dos casos.

---

## 3. Boas Práticas: O Que Deve Ser Utilizado

### Prática 1: Sempre Inicialize Structs Corretamente

**✅ BOM:**

```go
// Forma explícita e clara
pessoa := Pessoa{
    Nome:  "João",
    Idade: 30,
    Email: "joao@email.com",
}
```

**❌ RUIM:**

```go
// Valores zero podem causar bugs
var pessoa Pessoa
// pessoa.Email é "" - pode causar problemas se não verificar
```

**Por quê?** Valores zero podem ser válidos ou inválidos dependendo do contexto. Seja explícito.

---

### Prática 2: Use Funções Construtoras

**✅ BOM:**

```go
type Usuario struct {
    Nome  string
    Email string
    Ativo bool
}

// Função construtora com validação
func NovoUsuario(nome, email string) (*Usuario, error) {
    if email == "" {
        return nil, fmt.Errorf("email não pode ser vazio")
    }
    return &Usuario{
        Nome:  nome,
        Email: email,
        Ativo: true,  // Valor padrão
    }, nil
}
```

**❌ RUIM:**

```go
// Sem validação, pode criar estados inválidos
usuario := Usuario{
    Nome:  "João",
    Email: "",  // Email vazio - estado inválido!
}
```

**Por quê?** Construtores garantem que structs sejam criadas em estados válidos.

---

### Prática 3: Agrupe Campos Relacionados

**✅ BOM:**

```go
type Endereco struct {
    Rua    string
    Cidade string
    CEP    string
}

type Pessoa struct {
    Nome     string
    Idade    int
    Endereco Endereco  // Campos relacionados agrupados
}
```

**❌ RUIM:**

```go
type Pessoa struct {
    Nome      string
    Idade     int
    Rua       string  // Deveria estar em Endereco
    Cidade    string  // Deveria estar em Endereco
    CEP       string  // Deveria estar em Endereco
}
```

**Por quê?** Agrupar campos relacionados torna o código mais organizado e fácil de entender.

---

### Prática 4: Documente Structs Públicas

**✅ BOM:**

```go
// Usuario representa um usuário do sistema.
// Contém informações de autenticação e perfil.
type Usuario struct {
    // Nome é o nome completo do usuário.
    Nome string

    // Email é o endereço de email válido e único.
    Email string

    // Ativo indica se a conta está ativa.
    Ativo bool
}
```

**❌ RUIM:**

```go
// Sem documentação
type Usuario struct {
    Nome  string
    Email string
    Ativo bool
}
```

**Por quê?** Documentação ajuda outros desenvolvedores (e você no futuro) a entender o código.

---

### Prática 5: Use Métodos para Comportamento

**✅ BOM:**

```go
type ContaBancaria struct {
    Saldo float64
}

// Método encapsula a lógica
func (c *ContaBancaria) Sacar(valor float64) error {
    if valor > c.Saldo {
        return fmt.Errorf("saldo insuficiente")
    }
    c.Saldo -= valor
    return nil
}
```

**❌ RUIM:**

```go
// Lógica espalhada no código
func main() {
    conta := ContaBancaria{Saldo: 100}
    if 50 > conta.Saldo {  // Lógica fora da struct
        // erro
    }
    conta.Saldo -= 50
}
```

**Por quê?** Métodos encapsulam comportamento e tornam o código mais organizado.

---

## 4. O Que NÃO Deve Ser Feito

### ❌ Erro 1: Structs Muito Grandes

**RUIM:**

```go
// Struct com 50 campos - difícil de manter
type Usuario struct {
    Nome           string
    Email          string
    Telefone       string
    CPF            string
    RG             string
    DataNascimento time.Time
    EnderecoRua    string
    EnderecoNumero int
    EnderecoCidade string
    // ... mais 40 campos
}
```

**BOM:**

```go
// Dividir em structs menores
type Endereco struct {
    Rua    string
    Numero int
    Cidade string
}

type Documentos struct {
    CPF string
    RG  string
}

type Usuario struct {
    Nome           string
    Email          string
    Telefone       string
    DataNascimento time.Time
    Endereco       Endereco
    Documentos    Documentos
}
```

**Por quê?** Structs muito grandes são difíceis de manter e entender.

---

### ❌ Erro 2: Campos Públicos Quando Devem Ser Privados

**RUIM:**

```go
type ContaBancaria struct {
    Saldo float64  // Público - qualquer um pode modificar diretamente!
}

func main() {
    conta := ContaBancaria{Saldo: 100}
    conta.Saldo = -1000  // Pode fazer saldo negativo!
}
```

**BOM:**

```go
type ContaBancaria struct {
    saldo float64  // Privado - só pode ser modificado via métodos
}

func (c *ContaBancaria) Saldo() float64 {
    return c.saldo
}

func (c *ContaBancaria) Sacar(valor float64) error {
    if valor > c.saldo {
        return fmt.Errorf("saldo insuficiente")
    }
    c.saldo -= valor
    return nil
}
```

**Por quê?** Encapsulamento previne estados inválidos.

---

### ❌ Erro 3: Métodos que Não Pertencem à Struct

**RUIM:**

```go
type Pessoa struct {
    Nome string
}

// Este método não usa nada da struct - deveria ser função
func (p Pessoa) CalcularIMC(peso, altura float64) float64 {
    return peso / (altura * altura)
}
```

**BOM:**

```go
// Função independente - não precisa de struct
func CalcularIMC(peso, altura float64) float64 {
    return peso / (altura * altura)
}

// Ou se realmente precisar da struct
func (p Pessoa) CalcularMeuIMC(peso, altura float64) float64 {
    // Usa dados da pessoa
    return peso / (altura * altura)
}
```

**Por quê?** Métodos devem estar relacionados aos dados da struct.

---

### ❌ Erro 4: Comparar Structs Não Comparáveis

**RUIM:**

```go
type Config struct {
    Porta int
    Hosts []string  // Slice - não comparável!
}

func main() {
    c1 := Config{Porta: 8080, Hosts: []string{"a", "b"}}
    c2 := Config{Porta: 8080, Hosts: []string{"a", "b"}}

    // ERRO DE COMPILAÇÃO!
    fmt.Println(c1 == c2)
}
```

**BOM:**

```go
// Opção 1: Remover campos não comparáveis
type Config struct {
    Porta int
    Hosts string  // String separada por vírgula
}

// Opção 2: Função de comparação manual
func (c1 Config) Equals(c2 Config) bool {
    if c1.Porta != c2.Porta {
        return false
    }
    if len(c1.Hosts) != len(c2.Hosts) {
        return false
    }
    for i := range c1.Hosts {
        if c1.Hosts[i] != c2.Hosts[i] {
            return false
        }
    }
    return true
}
```

**Por quê?** Go não permite comparar structs com slices, maps ou funções.

---

### ❌ Erro 5: Mutabilidade Inesperada

**RUIM:**

```go
type Pessoa struct {
    Nome   string
    Amigos []string  // Slice é mutável!
}

func main() {
    p1 := Pessoa{Nome: "João", Amigos: []string{"Maria"}}
    p2 := p1  // Cópia da struct, mas...
    p2.Amigos = append(p2.Amigos, "Pedro")

    fmt.Println(p1.Amigos)  // ["Maria", "Pedro"] - SURPRESA!
    // p1 também foi modificado porque slice é referência!
}
```

**BOM:**

```go
// Opção 1: Fazer cópia do slice
func (p *Pessoa) AdicionarAmigo(nome string) {
    novosAmigos := make([]string, len(p.Amigos))
    copy(novosAmigos, p.Amigos)
    novosAmigos = append(novosAmigos, nome)
    p.Amigos = novosAmigos
}

// Opção 2: Documentar que é mutável
// Amigos é um slice mutável - modificações afetam todas as referências
type Pessoa struct {
    Nome   string
    Amigos []string
}
```

**Por quê?** Slices são referências - modificar um afeta outros.

---

## 5. Padrões de Design com Structs

### Padrão 1: Builder Pattern

**Para structs complexas:**

```go
type Usuario struct {
    Nome      string
    Email     string
    Telefone  string
    Endereco  string
    Ativo     bool
}

type UsuarioBuilder struct {
    usuario Usuario
}

func NovoUsuarioBuilder() *UsuarioBuilder {
    return &UsuarioBuilder{}
}

func (b *UsuarioBuilder) ComNome(nome string) *UsuarioBuilder {
    b.usuario.Nome = nome
    return b
}

func (b *UsuarioBuilder) ComEmail(email string) *UsuarioBuilder {
    b.usuario.Email = email
    return b
}

func (b *UsuarioBuilder) Build() (*Usuario, error) {
    if b.usuario.Email == "" {
        return nil, fmt.Errorf("email é obrigatório")
    }
    return &b.usuario, nil
}

// Uso:
usuario, _ := NovoUsuarioBuilder().
    ComNome("João").
    ComEmail("joao@email.com").
    Build()
```

---

### Padrão 2: Factory Pattern

**Para criar structs com lógica complexa:**

```go
func NovoUsuario(tipo string) (*Usuario, error) {
    switch tipo {
    case "admin":
        return &Usuario{
            Nome:  "Admin",
            Email: "admin@email.com",
            Tipo:  "admin",
        }, nil
    case "cliente":
        return &Usuario{
            Nome:  "Cliente",
            Email: "cliente@email.com",
            Tipo:  "cliente",
        }, nil
    default:
        return nil, fmt.Errorf("tipo inválido")
    }
}
```

---

## 6. Performance: Alinhamento de Memória

### O Que É Alinhamento?

Go alinha campos na memória para eficiência. Campos são alinhados em múltiplos do seu tamanho.

**Exemplo:**

```go
// Struct ineficiente (24 bytes)
type Ineficiente struct {
    a bool    // 1 byte + 7 bytes de padding
    b int64   // 8 bytes
    c bool    // 1 byte + 7 bytes de padding
}

// Struct eficiente (16 bytes)
type Eficiente struct {
    b int64   // 8 bytes
    a bool    // 1 byte
    c bool    // 1 byte + 6 bytes de padding
}
```

**Regra:** Coloque campos maiores primeiro para reduzir padding.

---

## 7. Boas Práticas: Nomenclatura

### ✅ BOM:

```go
// Nomes descritivos
type Usuario struct {
    NomeCompleto string
    Email        string
    DataCriacao  time.Time
}

// Métodos com nomes claros
func (u *Usuario) Ativar() error { ... }
func (u *Usuario) Desativar() error { ... }
```

### ❌ RUIM:

```go
// Nomes genéricos
type U struct {
    n string
    e string
    d time.Time
}

// Métodos com nomes confusos
func (u *U) a() error { ... }
```

---

## 8. Resumo: Regras de Ouro

1. **Use ponteiros** para structs grandes ou quando precisa modificar
2. **Use valores** para structs pequenas quando não precisa modificar
3. **Sempre use ponteiro** em métodos que modificam
4. **Documente structs públicas** completamente
5. **Use funções construtoras** para validação
6. **Agrupe campos relacionados** em structs aninhadas
7. **Evite structs muito grandes** - divida em menores
8. **Use encapsulamento** (campos privados) quando apropriado
9. **Coloque campos grandes primeiro** para melhor alinhamento
10. **Métodos devem estar relacionados** aos dados da struct

---

## 9. Comparação: Com vs. Sem Boas Práticas

### Cenário: Sistema de Usuários

**SEM boas práticas:**

- Structs muito grandes
- Campos públicos sem validação
- Sem construtores
- Métodos que não pertencem
- Código difícil de manter

**COM boas práticas:**

- Structs bem organizadas
- Encapsulamento adequado
- Construtores com validação
- Métodos bem definidos
- Código fácil de manter e estender

---

## Conclusão

Structs são fundamentais em Go, mas usá-las corretamente requer:

1. **Entender performance** - quando usar valor vs. ponteiro
2. **Organizar bem** - agrupar campos relacionados
3. **Encapsular** - proteger dados quando necessário
4. **Documentar** - ajudar outros desenvolvedores
5. **Validar** - garantir estados válidos

Lembre-se: **Código é lido muito mais vezes do que escrito**. Structs bem projetadas tornam o código mais fácil de entender e manter.

---

Na próxima etapa, você completará os exercícios e eu analisarei seu desempenho. Boa sorte!
