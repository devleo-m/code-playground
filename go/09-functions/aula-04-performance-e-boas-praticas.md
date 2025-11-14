# Módulo 9: Functions em Go

## Aula 4 - Performance, Boas Práticas e O Que Deve/Não Deve Ser Utilizado

Agora que você entende funções, vamos falar sobre **como usá-las de forma eficiente e profissional**. Esta é a parte que separa desenvolvedores iniciantes de desenvolvedores experientes.

---

## 1. Performance: Passagem por Valor vs Ponteiro

### ⚡ Structs Grandes: Use Ponteiros

**❌ RUIM: Passar Struct Grande por Valor**

```go
type PessoaGrande struct {
    Nome      string
    Idade     int
    Endereco  string
    Telefone  string
    Email     string
    Historico [1000]string  // Array grande!
}

func processarPessoa(p PessoaGrande) {
    // Cópia completa de ~8KB a cada chamada!
}

func main() {
    pessoa := PessoaGrande{...}
    processarPessoa(pessoa)  // Copia 8KB!
}
```

**Problema**: Copia toda a struct (pode ser muito custoso para structs grandes).

**✅ BOM: Passar Ponteiro para Structs Grandes**

```go
func processarPessoa(p *PessoaGrande) {
    // Copia apenas 8 bytes (ponteiro)!
}

func main() {
    pessoa := PessoaGrande{...}
    processarPessoa(&pessoa)  // Copia apenas ponteiro
}
```

**Ganho**: Para structs > 100 bytes, ponteiros são mais eficientes.

### Regra de Ouro: Quando Usar Ponteiros?

**Use ponteiros quando:**

- ✅ Struct > 100 bytes
- ✅ Precisa modificar o valor original
- ✅ Struct será passada para muitas funções
- ✅ Performance é crítica

**Use valores quando:**

- ✅ Struct pequena (< 100 bytes)
- ✅ Não precisa modificar original
- ✅ Quer garantir imutabilidade
- ✅ Código mais simples

---

## 2. Performance: Evitar Cópias Desnecessárias

### ❌ Ruim: Passar Arrays Grandes por Valor

```go
// RUIM: Copia 1 milhão de ints (~8MB)!
func processar(dados [1000000]int) {
    // ...
}

func main() {
    dados := [1000000]int{}
    processar(dados)  // Copia 8MB!
}
```

**✅ BOM: Usar Slices ou Ponteiros**

```go
// BOM: Slice (copia apenas header ~24 bytes)
func processar(dados []int) {
    // ...
}

// OU ponteiro para array
func processar(dados *[1000000]int) {
    // ...
}
```

**Ganho**: De 8MB copiados para 24 bytes (slice) ou 8 bytes (ponteiro)!

---

## 3. Performance: Funções Inline vs Chamadas

### Custo de Chamada de Função

Cada chamada de função tem custo:

- Salvar contexto (stack frame)
- Passar argumentos
- Pular para código da função
- Retornar e restaurar contexto

**Para funções muito simples, o custo pode ser maior que a operação!**

### ✅ Bom: Funções Simples Podem Ser Inline

O compilador Go faz **inline** de funções pequenas automaticamente:

```go
// Função simples - compilador pode fazer inline
func somar(a, b int) int {
    return a + b
}

func main() {
    resultado := somar(5, 3)  // Pode ser inline, sem custo de chamada
}
```

**Regra**: Funções muito pequenas (< 10 linhas, sem loops) são candidatas a inline.

### ❌ Evite: Funções Muito Pequenas que Não Ajudam

```go
// RUIM: Função tão simples que não vale a pena
func getX(p Pessoa) string {
    return p.Nome
}

// BOM: Acesse diretamente
nome := pessoa.Nome
```

**Exceção**: Se a função adiciona valor (validação, transformação), mantenha-a.

---

## 4. Performance: Closures e Captura de Variáveis

### ⚠️ Custo de Closures

Closures têm custo adicional:

- Alocação no heap (se captura variáveis)
- Indireção para acessar variáveis capturadas

### ✅ Bom: Closure Simples

```go
// BOM: Closure simples, captura pequena
func criarMultiplicador(fator int) func(int) int {
    return func(n int) int {
        return n * fator  // Captura apenas int (8 bytes)
    }
}
```

### ❌ Ruim: Closure que Captura Struct Grande

```go
// RUIM: Captura struct grande no heap
type ConfigGrande struct {
    dados [1000]string
}

func criarProcessador(config ConfigGrande) func() {
    return func() {
        // Captura config (muito grande!)
    }
}
```

**Solução**: Capture apenas o necessário:

```go
func criarProcessador(config ConfigGrande) func() {
    apenasNecessario := config.campoNecessario  // Captura apenas o necessário
    return func() {
        // Usa apenasNecessario
    }
}
```

---

## 5. Performance: Variadic Functions

### Custo de Variadic

Variadic functions criam um slice temporário:

```go
func somar(numeros ...int) int {
    // Go cria []int temporário
}
```

**Custo**: Alocação de slice (mesmo que pequena).

### ✅ Bom: Variadic Quando Apropriado

```go
// BOM: Flexibilidade vale o custo
func juntar(separador string, textos ...string) string {
    return strings.Join(textos, separador)
}
```

### ⚠️ Evite: Variadic em Hot Paths

Se performance é crítica e você sabe o número de argumentos:

```go
// RUIM: Variadic em loop quente
for i := 0; i < 1000000; i++ {
    somar(1, 2, 3)  // Aloca slice a cada chamada!
}

// BOM: Slice pré-alocado
numeros := []int{1, 2, 3}
for i := 0; i < 1000000; i++ {
    somarSlice(numeros)  // Sem alocação
}
```

---

## 6. Boas Práticas: Nomenclatura

### ✅ BOM: Nomes Descritivos

```go
// BOM: Nome claro e descritivo
func calcularAreaRetangulo(largura, altura float64) float64 {
    return largura * altura
}

func buscarUsuarioPorID(id int) (*Usuario, error) {
    // ...
}
```

### ❌ RUIM: Nomes Genéricos

```go
// RUIM: Nome vago
func calc(a, b float64) float64 {
    // O que calcula?
}

func buscar(x int) (*Usuario, error) {
    // Busca o quê? Por quê?
}
```

### Convenções de Nomenclatura

- **Verbos**: `calcular`, `buscar`, `processar`, `validar`
- **Específicos**: `calcularAreaRetangulo` > `calcular`
- **Consistentes**: `buscarUsuario` e `buscarProduto` (não `getUser` e `buscarProduto`)

---

## 7. Boas Práticas: Tamanho de Funções

### ✅ BOM: Funções Pequenas e Focadas

```go
// BOM: Uma responsabilidade
func validarEmail(email string) error {
    if !strings.Contains(email, "@") {
        return fmt.Errorf("email inválido")
    }
    return nil
}
```

### ❌ RUIM: Funções Muito Grandes

```go
// RUIM: Faz muitas coisas
func processarUsuario(dados map[string]interface{}) error {
    // Valida email
    // Valida senha
    // Cria usuário
    // Envia email
    // Salva no banco
    // Gera token
    // ... 200 linhas
}
```

**Regra**: Funções devem fazer **uma coisa bem feita**.

**Tamanho ideal**: 10-50 linhas (depende da complexidade).

---

## 8. Boas Práticas: Tratamento de Erros

### ✅ BOM: Sempre Retornar Erro Quando Apropriado

```go
// BOM: Retorna erro para operações que podem falhar
func dividir(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("divisão por zero")
    }
    return a / b, nil
}
```

### ❌ RUIM: Ignorar Erros Silenciosamente

```go
// RUIM: Ignora erro
func dividir(a, b float64) float64 {
    if b == 0 {
        return 0  // Erro silencioso!
    }
    return a / b
}
```

### Padrão Idiomático: Resultado e Erro

**Sempre retorne `(resultado, error)` para operações que podem falhar:**

```go
func buscarUsuario(id int) (*Usuario, error)
func lerArquivo(nome string) ([]byte, error)
func conectarBanco() (*DB, error)
```

---

## 9. Boas Práticas: Named Returns

### ✅ BOM: Named Returns Quando Melhora Legibilidade

```go
// BOM: Nomes claros melhoram legibilidade
func calcularEstatisticas(numeros []float64) (media, max, min float64) {
    // Código claro sobre o que retorna
    media = calcularMedia(numeros)
    max = encontrarMaximo(numeros)
    min = encontrarMinimo(numeros)
    return
}
```

### ❌ RUIM: Named Returns Desnecessários

```go
// RUIM: Named return não adiciona valor
func somar(a, b int) (resultado int) {
    resultado = a + b
    return
}

// BOM: Mais simples sem named return
func somar(a, b int) int {
    return a + b
}
```

### ⚠️ Cuidado: Named Returns e Defer

Named returns podem causar confusão com `defer`:

```go
func exemplo() (resultado int) {
    defer func() {
        resultado++  // Modifica named return!
    }()
    return 5  // Retorna 6, não 5!
}
```

**Use com cuidado quando há `defer`!**

---

## 10. Boas Práticas: Documentação

### ✅ BOM: Documentar Funções Públicas

```go
// CalcularAreaRetangulo calcula a área de um retângulo
// dado sua largura e altura.
// Retorna a área em unidades quadradas.
func CalcularAreaRetangulo(largura, altura float64) float64 {
    return largura * altura
}
```

### Padrão de Documentação

```go
// NomeDaFuncao faz uma breve descrição do que a função faz.
//
// Parâmetros são descritos se não são óbvios.
// Retornos são descritos se não são óbvios.
// Exemplos são úteis para funções complexas.
func NomeDaFuncao(parametro Tipo) TipoRetorno {
    // ...
}
```

---

## 11. O Que NÃO Fazer

### ❌ Erro 1: Funções com Muitos Parâmetros

```go
// RUIM: Muitos parâmetros (difícil de usar)
func criarUsuario(nome, email, telefone, endereco, cidade, estado, cep string, idade int, ativo bool) error {
    // ...
}
```

**Solução**: Use struct:

```go
// BOM: Struct para muitos parâmetros
type DadosUsuario struct {
    Nome     string
    Email    string
    Telefone string
    // ...
}

func criarUsuario(dados DadosUsuario) error {
    // ...
}
```

### ❌ Erro 2: Retornar Erro Sem Verificar

```go
// RUIM: Ignora erro
resultado, _ := dividir(10, 0)
fmt.Println(resultado)  // 0 - mas deveria ser erro!
```

**Sempre verifique erros:**

```go
// BOM: Verifica erro
resultado, err := dividir(10, 0)
if err != nil {
    return err
}
```

### ❌ Erro 3: Closure em Loop sem Cuidado

```go
// RUIM: Todas capturam mesma variável
var funcoes []func() int
for i := 0; i < 3; i++ {
    funcoes = append(funcoes, func() int {
        return i  // Todas retornam 3!
    })
}
```

**Solução**: Criar nova variável:

```go
// BOM: Cada closure captura variável diferente
for i := 0; i < 3; i++ {
    i := i  // Nova variável
    funcoes = append(funcoes, func() int {
        return i
    })
}
```

---

## 12. Padrões Idiomáticos em Go

### Padrão 1: Função Construtora

```go
type Config struct {
    Host string
    Port int
}

// Função construtora com validação
func NovaConfig(host string, port int) (*Config, error) {
    if host == "" {
        return nil, fmt.Errorf("host não pode ser vazio")
    }
    if port <= 0 {
        return nil, fmt.Errorf("porta deve ser positiva")
    }
    return &Config{Host: host, Port: port}, nil
}
```

### Padrão 2: Função de Validação

```go
func validarEmail(email string) error {
    if !strings.Contains(email, "@") {
        return fmt.Errorf("email deve conter @")
    }
    if !strings.Contains(email, ".") {
        return fmt.Errorf("email deve conter .")
    }
    return nil
}
```

### Padrão 3: Função com Opções

```go
type Opcoes struct {
    Timeout time.Duration
    Retries int
}

func processarComOpcoes(opcoes Opcoes) {
    // Usa opcoes
}

// Com valores padrão
func processar() {
    processarComOpcoes(Opcoes{
        Timeout: 5 * time.Second,
        Retries: 3,
    })
}
```

---

## 13. Performance: Benchmarks Reais

### Benchmark: Passagem por Valor vs Ponteiro

```go
type Pessoa struct {
    Nome     string
    Idade    int
    Endereco string
    // ~50 bytes
}

func BenchmarkPorValor(b *testing.B) {
    p := Pessoa{Nome: "João", Idade: 30, Endereco: "Rua X"}
    for i := 0; i < b.N; i++ {
        processarValor(p)
    }
}

func BenchmarkPorPonteiro(b *testing.B) {
    p := &Pessoa{Nome: "João", Idade: 30, Endereco: "Rua X"}
    for i := 0; i < b.N; i++ {
        processarPonteiro(p)
    }
}
```

**Resultado típico:**

- Por valor: ~2ns por chamada
- Por ponteiro: ~1ns por chamada

**Para structs grandes (> 100 bytes), ponteiros são mais rápidos.**

---

## 14. Checklist de Boas Práticas

Antes de considerar sua função "pronta", pergunte-se:

- [ ] Nome da função é claro e descritivo?
- [ ] Função faz apenas uma coisa?
- [ ] Parâmetros são razoáveis em número (< 5)?
- [ ] Retorna erro quando apropriado?
- [ ] Documentada se é pública?
- [ ] Tamanho razoável (< 100 linhas)?
- [ ] Usa ponteiros para structs grandes?
- [ ] Evita cópias desnecessárias?
- [ ] Trata erros corretamente?
- [ ] Segue convenções de nomenclatura?

---

## 15. Resumo: Regras de Ouro

1. **Pequenas e focadas**: Uma função, uma responsabilidade
2. **Nomes claros**: Nome deve descrever o que faz
3. **Erros explícitos**: Sempre retorne erro quando apropriado
4. **Ponteiros para grandes**: Use ponteiros para structs > 100 bytes
5. **Evite cópias**: Use slices ao invés de arrays grandes
6. **Documente públicas**: Funções públicas devem ser documentadas
7. **Cuidado com closures**: Entenda o que está capturando
8. **Padrão idiomático**: `(resultado, error)` para operações que falham

---

## Conclusão

Funções são a base da organização de código. Usá-las corretamente não é apenas sobre sintaxe, mas sobre:

- **Design**: Como organizar código de forma clara
- **Performance**: Como evitar custos desnecessários
- **Manutenibilidade**: Como facilitar mudanças futuras
- **Idiomático**: Como escrever código "Go-like"

Lembre-se: **Código é lido muito mais vezes do que escrito**. Priorize clareza, e otimize apenas quando necessário e medido.

---

Na próxima etapa, você completará os exercícios e eu analisarei seu desempenho. Boa sorte!
