# Aula 20 - Performance e Boas Práticas: Standard Library

Olá! Agora que você conhece os principais pacotes da Standard Library, é crucial entender **como usá-los de forma eficiente** e quais são as **armadilhas comuns** a evitar. Esta aula vai te preparar para escrever código profissional e performático.

---

## 🚀 Performance: O que Fazer e o que Evitar

### 1. I/O e File Handling

#### ✅ FAZER: Usar `bufio` para Operações Frequentes

**Problema:** Ler/escrever byte por byte faz muitas chamadas de sistema (syscalls), que são lentas.

```go
// ❌ LENTO: Muitas chamadas de sistema
arquivo, _ := os.Open("grande.txt")
defer arquivo.Close()
dados := make([]byte, 1)
for {
    n, err := arquivo.Read(dados)
    if err != nil {
        break
    }
    // Processa 1 byte por vez...
}
```

```go
// ✅ RÁPIDO: Buffer reduz chamadas de sistema
arquivo, _ := os.Open("grande.txt")
defer arquivo.Close()
scanner := bufio.NewScanner(arquivo)
for scanner.Scan() {
    linha := scanner.Text()
    // Processa linha inteira
}
```

**Regra de Ouro:** Se você vai ler/escrever mais de algumas vezes, use `bufio`.

#### ✅ FAZER: Sempre Fechar Recursos com `defer`

```go
// ✅ CORRETO
arquivo, err := os.Open("arquivo.txt")
if err != nil {
    return err
}
defer arquivo.Close() // Garante fechamento mesmo em caso de erro
// ... resto do código
```

**Por quê?** Arquivos abertos consomem recursos do sistema. Se não fechar, pode esgotar o limite de arquivos abertos.

#### ❌ EVITAR: Ler Arquivo Inteiro na Memória sem Necessidade

```go
// ❌ PROBLEMA: Se o arquivo for muito grande, pode esgotar memória
dados, _ := os.ReadFile("arquivo_gigante.txt") // Lê tudo na memória
```

```go
// ✅ MELHOR: Processar em chunks ou linha por linha
arquivo, _ := os.Open("arquivo_gigante.txt")
defer arquivo.Close()
scanner := bufio.NewScanner(arquivo)
for scanner.Scan() {
    // Processa uma linha por vez
}
```

**Regra:** Para arquivos grandes (>10MB), processe em chunks ou streaming.

---

### 2. JSON: Marshal/Unmarshal

#### ✅ FAZER: Reutilizar `json.Encoder`/`json.Decoder` para Streams

```go
// ❌ LENTO: Marshal múltiplas vezes
for _, pessoa := range pessoas {
    jsonBytes, _ := json.Marshal(pessoa)
    arquivo.Write(jsonBytes)
}
```

```go
// ✅ RÁPIDO: Encoder reutilizado
encoder := json.NewEncoder(arquivo)
for _, pessoa := range pessoas {
    encoder.Encode(pessoa) // Mais eficiente
}
```

#### ✅ FAZER: Usar Struct Tags Apropriadamente

```go
// ✅ BOM: Campos exportados + tags apropriadas
type Pessoa struct {
    Nome    string `json:"nome"`
    Email   string `json:"email,omitempty"` // Omitir se vazio
    Salario float64 `json:"-"`              // Nunca serializar
}
```

**Benefícios:**
- `omitempty`: Reduz tamanho do JSON quando campos estão vazios
- `-`: Evita serializar dados sensíveis

#### ❌ EVITAR: Marshal/Unmarshal Desnecessários

```go
// ❌ PROBLEMA: Marshal duas vezes
json1, _ := json.Marshal(pessoa)
json2, _ := json.MarshalIndent(pessoa, "", "  ") // Marshal novamente!
```

```go
// ✅ MELHOR: Marshal uma vez, formatar depois se necessário
jsonBytes, _ := json.Marshal(pessoa)
// Se precisar indentado, use uma biblioteca de formatação JSON
```

---

### 3. Regex: Compilar uma Vez, Usar Muitas

#### ❌ EVITAR: Compilar Regex a Cada Uso

```go
// ❌ LENTO: Compila regex toda vez
func validarEmail(email string) bool {
    matched, _ := regexp.MatchString(`\b[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Z|a-z]{2,}\b`, email)
    return matched
}
```

```go
// ✅ RÁPIDO: Compila uma vez (package level)
var emailRegex = regexp.MustCompile(`\b[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Z|a-z]{2,}\b`)

func validarEmail(email string) bool {
    return emailRegex.MatchString(email)
}
```

**Regra:** Se você vai usar o mesmo padrão regex múltiplas vezes, compile uma vez e reutilize.

#### ⚠️ CUIDADO: Regex Pode Ser Lento para Textos Grandes

Para validações simples (email, telefone), regex é OK. Para processamento de texto muito grande, considere alternativas mais eficientes.

---

### 4. Time: Evitar Parsing Repetido

#### ❌ EVITAR: Parsing de Layouts Repetidamente

```go
// ❌ LENTO: Parse toda vez
func parsearData(str string) time.Time {
    t, _ := time.Parse("2006-01-02 15:04:05", str)
    return t
}
```

```go
// ✅ MELHOR: Definir layout como constante
const layoutData = "2006-01-02 15:04:05"

func parsearData(str string) time.Time {
    t, _ := time.Parse(layoutData, str)
    return t
}
```

#### ✅ FAZER: Usar `time.RFC3339` para APIs

```go
// ✅ PADRÃO: RFC3339 é universal e eficiente
timestamp := time.Now().Format(time.RFC3339)
// "2024-01-15T10:30:00Z"
```

**Por quê?** É o formato padrão da web, bem suportado e eficiente.

---

### 5. Flag: Validação e Help Automático

#### ✅ FAZER: Usar Flags com Valores Padrão Sensatos

```go
// ✅ BOM: Valores padrão úteis
porta := flag.Int("porta", 8080, "Porta do servidor")
timeout := flag.Duration("timeout", 30*time.Second, "Timeout da requisição")
```

#### ✅ FAZER: Validar Flags Após Parse

```go
porta := flag.Int("porta", 8080, "Porta do servidor")
flag.Parse()

// ✅ VALIDAÇÃO: Garantir que valores estão em range válido
if *porta < 1 || *porta > 65535 {
    log.Fatal("Porta deve estar entre 1 e 65535")
}
```

---

### 6. Slog: Configuração para Produção

#### ✅ FAZER: Usar JSON Handler em Produção

```go
// ✅ PRODUÇÃO: JSON é melhor para análise
opts := &slog.HandlerOptions{
    Level: slog.LevelInfo, // Não logar Debug em produção
}
logger := slog.New(slog.NewJSONHandler(os.Stdout, opts))
```

#### ❌ EVITAR: Logar em Excessos

```go
// ❌ PROBLEMA: Muitos logs podem degradar performance
for i := 0; i < 1000000; i++ {
    logger.Debug("Processando item", "item", i) // Muito lento!
}
```

```go
// ✅ MELHOR: Logar apenas o necessário
if logger.Enabled(context.Background(), slog.LevelDebug) {
    logger.Debug("Processando", "total", len(items))
}
```

**Regra:** Em produção, use nível `Info` ou superior. `Debug` apenas em desenvolvimento.

---

### 7. go:embed: Quando Usar e Quando Não

#### ✅ USAR `go:embed` para:
- Templates HTML/CSS que raramente mudam
- Arquivos de configuração padrão
- Assets estáticos (imagens pequenas, ícones)
- Dados de referência (listas, dicionários)

#### ❌ NÃO USAR `go:embed` para:
- Arquivos que mudam frequentemente
- Arquivos muito grandes (>10MB) - aumenta tamanho do binário
- Dados que precisam ser atualizados sem recompilar
- Arquivos de configuração que variam por ambiente

**Trade-off:** `go:embed` aumenta o tamanho do binário, mas elimina dependências externas.

---

## 🎯 Boas Práticas Gerais

### 1. Sempre Trate Erros

```go
// ❌ RUIM: Ignorar erros
dados, _ := os.ReadFile("arquivo.txt")

// ✅ BOM: Tratar erros apropriadamente
dados, err := os.ReadFile("arquivo.txt")
if err != nil {
    return fmt.Errorf("erro ao ler arquivo: %w", err)
}
```

### 2. Use `defer` para Limpeza

```go
// ✅ SEMPRE: defer para recursos
arquivo, err := os.Open("arquivo.txt")
if err != nil {
    return err
}
defer arquivo.Close() // Garante fechamento
```

### 3. Prefira Interfaces da Standard Library

```go
// ✅ BOM: Função genérica que aceita qualquer Reader
func processarDados(r io.Reader) error {
    // Funciona com arquivo, rede, buffer, etc.
}

// ❌ RUIM: Função específica apenas para arquivos
func processarArquivo(arquivo *os.File) error {
    // Só funciona com arquivos
}
```

### 4. Documente Comportamentos Importantes

```go
// ProcessarArquivo processa um arquivo linha por linha.
// IMPORTANTE: Arquivos grandes (>1GB) devem ser processados
// em chunks para evitar esgotamento de memória.
func ProcessarArquivo(caminho string) error {
    // ...
}
```

---

## ⚡ Performance: Métricas e Benchmarks

### Como Medir Performance

```go
package main

import (
    "fmt"
    "os"
    "time"
)

func main() {
    inicio := time.Now()
    
    // Seu código aqui
    dados, _ := os.ReadFile("arquivo.txt")
    _ = dados
    
    duracao := time.Since(inicio)
    fmt.Printf("Tempo de execução: %v\n", duracao)
}
```

### Comparação: Com vs Sem Buffer

```go
// Sem buffer: ~500ms para arquivo de 10MB
// Com buffer: ~50ms para arquivo de 10MB
// Diferença: 10x mais rápido!
```

---

## 🚨 Armadilhas Comuns

### 1. Esquecer de Fechar Arquivos

```go
// ❌ PROBLEMA: Vazamento de recursos
arquivo, _ := os.Open("arquivo.txt")
// Esqueceu de fechar!
```

**Solução:** Sempre use `defer arquivo.Close()`.

### 2. Não Verificar Erros de I/O

```go
// ❌ PROBLEMA: Pode continuar com dados inválidos
dados, _ := os.ReadFile("arquivo.txt")
processar(dados) // dados pode estar vazio ou corrompido!
```

**Solução:** Sempre verifique erros.

### 3. Regex Compilado Múltiplas Vezes

```go
// ❌ PROBLEMA: Performance ruim
for _, texto := range textos {
    matched, _ := regexp.MatchString(pattern, texto) // Compila toda vez!
}
```

**Solução:** Compile uma vez, reutilize.

### 4. Marshal JSON Desnecessário

```go
// ❌ PROBLEMA: Marshal múltiplas vezes
json1, _ := json.Marshal(pessoa)
json2, _ := json.MarshalIndent(pessoa, "", "  ") // Marshal novamente!
```

**Solução:** Marshal uma vez, formate depois se necessário.

---

## 📊 Resumo: Checklist de Performance

Antes de considerar seu código otimizado, verifique:

- [ ] Uso `bufio` para operações de I/O frequentes?
- [ ] Fecho todos os recursos com `defer`?
- [ ] Compilo regex uma vez e reutilizo?
- [ ] Uso `json.Encoder`/`Decoder` para streams?
- [ ] Trato todos os erros apropriadamente?
- [ ] Uso `slog` com nível apropriado (Info+ em produção)?
- [ ] Evito ler arquivos grandes inteiros na memória?
- [ ] Uso interfaces (`io.Reader`/`Writer`) quando possível?
- [ ] Valido flags após parsing?
- [ ] Escolhi `go:embed` apenas quando apropriado?

---

## 🎓 Conclusão

A Standard Library do Go é poderosa, mas precisa ser usada corretamente. As principais lições são:

1. **Buffer é seu amigo** para I/O frequente
2. **Compile uma vez, use muitas vezes** (regex, layouts)
3. **Sempre feche recursos** com `defer`
4. **Trate erros** - não ignore
5. **Use interfaces** para flexibilidade e testabilidade
6. **Configure logging apropriadamente** para o ambiente

Com essas práticas, você estará escrevendo código Go profissional, eficiente e manutenível!


