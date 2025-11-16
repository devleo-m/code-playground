# Módulo 29: Comandos Core do Go
## Aula 2 - Simplificada: Entendendo os Comandos Core do Go na Prática

Olá! Agora vamos entender esses comandos de uma forma muito mais simples, usando analogias do dia a dia. Imagine que você é um **chef de cozinha** e cada comando Go é uma ferramenta diferente na sua cozinha!

---

## 🍳 1. `go run` - O Microondas Rápido

### A Analogia da Cozinha

Imagine que você quer **testar uma receita rapidamente**. Você não quer sujar muitos pratos, não quer preparar tudo formalmente - só quer ver se funciona!

**`go run` é como usar o microondas:**
- ⚡ **Rápido**: Aquece (compila) e serve (executa) na hora
- 🧹 **Sem bagunça**: Não deixa pratos sujos (não cria arquivos binários)
- 🧪 **Para testar**: Perfeito para experimentar receitas novas
- ❌ **Não para servir**: Não é o que você serve aos convidados (produção)

### Exemplo do Dia a Dia

```bash
# Você escreveu um código rápido para testar uma ideia
go run teste.go
# Pronto! Executou e você viu o resultado
# Não deixou nenhum arquivo .exe no seu computador
```

**Quando usar?** Quando você está **experimentando** e quer ver o resultado rápido, sem se preocupar em criar um "prato final" (binário).

---

## 🏭 2. `go build` - A Fábrica de Embalagens

### A Analogia da Cozinha

Agora você quer **preparar o prato final** para servir aos clientes (usuários). Você precisa de uma embalagem bonita, durável e que funcione em qualquer lugar!

**`go build` é como uma fábrica de embalagens:**
- 📦 **Cria o produto final**: Gera um executável completo
- 🌍 **Funciona em qualquer lugar**: Pode criar para Windows, Linux, macOS
- 🎯 **Otimizado**: O produto é eficiente e rápido
- 💼 **Para distribuir**: É o que você entrega aos clientes

### Exemplo do Dia a Dia

```bash
# Você quer criar um programa para seus amigos usarem
go build -o meuapp main.go

# Agora você tem um arquivo "meuapp" que seus amigos podem executar
# Funciona mesmo se eles não tiverem Go instalado!
```

**Quando usar?** Quando você quer **criar o produto final** que outras pessoas vão usar, como um aplicativo ou ferramenta.

### Cross-Compilation: A Fábrica Multinacional

```bash
# Você está no Brasil (Linux) mas quer criar um app para seu amigo no Windows
GOOS=windows go build -o app.exe main.go

# É como uma fábrica que produz embalagens diferentes para diferentes países!
```

---

## 🛠️ 3. `go install` - A Loja de Ferramentas

### A Analogia da Cozinha

Você comprou uma **ferramenta nova** (como um descascador de legumes profissional). Você quer que ela fique **sempre disponível** na sua cozinha, não apenas em um projeto específico.

**`go install` é como instalar uma ferramenta na sua loja de ferramentas:**
- 🏪 **Disponível globalmente**: A ferramenta fica em um lugar especial (`$GOPATH/bin`)
- 🔧 **Para usar sempre**: Você pode usar de qualquer projeto
- 📦 **Versões**: Pode instalar versões específicas
- 🎯 **Ideal para CLIs**: Perfeito para ferramentas de linha de comando

### Exemplo do Dia a Dia

```bash
# Você quer instalar uma ferramenta útil (como um formatador de código)
go install golang.org/x/tools/cmd/godoc@latest

# Agora você pode usar "godoc" de qualquer lugar no seu computador!
# É como ter uma ferramenta sempre à mão na sua gaveta de ferramentas
```

**Diferença prática:**

```bash
# go build: Cria o executável AQUI (no diretório atual)
go build -o ferramenta .
# Resultado: ./ferramenta (no diretório atual)

# go install: Instala a ferramenta GLOBALMENTE
go install .
# Resultado: ~/go/bin/ferramenta (disponível em qualquer lugar)
```

**Quando usar?** Quando você quer **instalar ferramentas** que vai usar em vários projetos, como linters, formatadores, ou suas próprias ferramentas CLI.

---

## ✂️ 4. `go fmt` - O Organizador Automático

### A Analogia da Cozinha

Sua cozinha está **bagunçada**: facas em lugares errados, pratos desalinhados, temperos fora de ordem. Você precisa de alguém que **organize tudo automaticamente** seguindo um padrão perfeito!

**`go fmt` é como um organizador automático de cozinha:**
- 🤖 **Automático**: Organiza tudo sozinho
- 📏 **Padrão único**: Todo mundo usa o mesmo padrão (sem discussões!)
- ✨ **Limpo e bonito**: Deixa tudo alinhado e consistente
- ⚡ **Rápido**: Organiza em segundos

### Exemplo do Dia a Dia

**Antes (bagunçado):**
```go
package main
import "fmt"
func main(){
x:=10
y:=20
fmt.Println(x+y)
}
```

**Depois do `go fmt` (organizado):**
```go
package main

import "fmt"

func main() {
	x := 10
	y := 20
	fmt.Println(x + y)
}
```

**É como ter um assistente que:**
- ✅ Alinha todos os pratos
- ✅ Organiza as facas no lugar certo
- ✅ Deixa os espaços consistentes
- ✅ Segue sempre o mesmo padrão (sem você precisar pensar!)

**Quando usar?** **SEMPRE antes de mostrar seu código para outras pessoas** (commitar, fazer pull request, etc.). É como arrumar a cama antes de receber visitas!

---

## 📚 5. `go mod` - A Biblioteca de Receitas

### A Analogia da Cozinha

Você está cozinhando e precisa de **ingredientes especiais** (dependências). Você tem uma **biblioteca de receitas** (`go.mod`) que lista todos os ingredientes que você usa, e um **catálogo de fornecedores confiáveis** (`go.sum`) que garante que os ingredientes são autênticos.

**`go mod` é como gerenciar sua biblioteca de receitas:**
- 📖 **Lista de ingredientes**: `go.mod` lista todas as dependências
- 🔒 **Fornecedores confiáveis**: `go.sum` verifica que os ingredientes são autênticos
- 🛒 **Comprar ingredientes**: `go mod download` baixa as dependências
- 🧹 **Limpar receitas antigas**: `go mod tidy` remove ingredientes não usados

### Comandos na Prática

#### `go mod init` - Criar Nova Biblioteca

```bash
# Você está começando um novo projeto de culinária
go mod init minha-receita

# É como criar uma nova pasta para guardar suas receitas
# Agora você tem um "go.mod" vazio, pronto para adicionar ingredientes
```

#### `go mod tidy` - Organizar a Biblioteca

```bash
# Você usou alguns ingredientes, mas não anotou todos
# E tem ingredientes na lista que não está mais usando
go mod tidy

# É como organizar sua despensa:
# ✅ Adiciona ingredientes que você esqueceu de anotar
# ❌ Remove ingredientes que você não usa mais
# 📝 Atualiza a lista para ficar certinha
```

#### `go mod download` - Comprar Ingredientes

```bash
# Você quer garantir que tem todos os ingredientes em casa
go mod download

# É como fazer uma compra online de todos os ingredientes
# Eles ficam guardados no seu "armário" (cache) para usar depois
```

### Exemplo Completo

```bash
# 1. Começar novo projeto
mkdir meu-projeto
cd meu-projeto
go mod init github.com/eu/meu-projeto

# 2. Escrever código que usa uma biblioteca
# (Go adiciona automaticamente ao go.mod quando você faz go build/run)

# 3. Organizar tudo
go mod tidy

# Agora seu go.mod está limpo e organizado!
```

**Quando usar?**
- ✅ **Sempre em projetos novos**: `go mod init`
- ✅ **Antes de commitar**: `go mod tidy`
- ✅ **Para garantir dependências**: `go mod download`

---

## 🧪 6. `go test` - O Laboratório de Qualidade

### A Analogia da Cozinha

Você criou uma receita nova, mas **como saber se ela funciona bem?** Você precisa testar! É como ter um **laboratório de qualidade** onde você prova cada receita antes de servir aos clientes.

**`go test` é como seu laboratório de testes:**
- 🔬 **Testa automaticamente**: Encontra e executa todos os testes
- 📊 **Relatórios detalhados**: Mostra o que passou e o que falhou
- ⚡ **Testa velocidade**: Pode medir performance (benchmarks)
- 📈 **Cobertura**: Mostra quanto do seu código foi testado

### Tipos de Testes

#### Testes Normais - Provar a Receita

```go
// Você quer testar se sua função Soma funciona
func TestSoma(t *testing.T) {
    resultado := Soma(2, 3)
    if resultado != 5 {
        t.Error("A receita não deu certo! Esperava 5, mas deu", resultado)
    }
}
```

**É como:** Provar se o bolo ficou doce o suficiente, se o sal está na medida certa, etc.

#### Benchmarks - Medir Velocidade

```go
// Você quer saber QUÃO RÁPIDO sua função é
func BenchmarkSoma(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Soma(2, 3)  // Executa milhões de vezes para medir velocidade
    }
}
```

**É como:** Cronometrar quanto tempo leva para fazer a receita.

### Exemplos Práticos

```bash
# Testar tudo
go test
# Output: ✅ Todos os testes passaram!

# Testar com detalhes
go test -v
# Output: Mostra cada teste executado e se passou ou falhou

# Medir velocidade
go test -bench=.
# Output: Mostra quantas operações por segundo sua função faz

# Ver cobertura (quanto do código foi testado)
go test -cover
# Output: cover: 85.3% of statements
```

**Quando usar?**
- ✅ **Sempre antes de commitar**: Garantir que nada quebrou
- ✅ **Durante desenvolvimento**: Testar enquanto escreve (TDD)
- ✅ **Para medir performance**: Ver se suas otimizações funcionaram

**É como:** Provar cada prato antes de servir aos clientes!

---

## 🧹 7. `go clean` - A Limpeza Geral

### A Analogia da Cozinha

Sua cozinha está **cheia de bagulhos**: pratos sujos de testes anteriores, ingredientes velhos no armário, ferramentas espalhadas. Você precisa fazer uma **limpeza geral**!

**`go clean` é como fazer uma limpeza geral na cozinha:**
- 🗑️ **Remove bagulhos**: Limpa arquivos temporários de compilação
- 📦 **Limpa armários**: Remove cache antigo (pode liberar muito espaço!)
- 🧹 **Deixa tudo limpo**: Garante que o próximo build seja "do zero"

### O Que Cada Flag Faz

```bash
# Limpeza básica (remove arquivos .o, binários locais)
go clean
# É como: Limpar a bancada e jogar pratos sujos fora

# Limpar cache de build
go clean -cache
# É como: Limpar o "armário de ingredientes preparados" (pode ter GB de coisas!)

# Limpar cache de módulos
go clean -modcache
# É como: Limpar o "depósito de ingredientes comprados" (pode ter dezenas de GB!)

# Limpar TUDO
go clean -cache -modcache -testcache
# É como: Uma limpeza geral completa da cozinha inteira!
```

### Quando Você Precisa Limpar?

**Problemas estranhos no build?**
```bash
go clean -cache
# Às vezes o cache fica corrompido e causa problemas
# Limpar resolve na maioria das vezes!
```

**Disco cheio?**
```bash
# Verificar tamanho
du -sh $(go env GOCACHE)      # Ver tamanho do cache
du -sh $(go env GOMODCACHE)   # Ver tamanho dos módulos

# Limpar se necessário
go clean -cache -modcache
# Pode liberar dezenas de GB!
```

**Build limpo para produção?**
```bash
go clean
go build -o app .
# Garante que está compilando "do zero", sem usar cache antigo
```

**É como:** Fazer uma limpeza geral antes de uma ocasião especial, ou quando a cozinha está muito bagunçada!

---

## 📖 8. `go doc` - O Manual de Instruções

### A Analogia da Cozinha

Você comprou uma **ferramenta nova** (como um processador de alimentos), mas não sabe como usar. Você precisa do **manual de instruções**!

**`go doc` é como ter acesso instantâneo ao manual de qualquer ferramenta:**
- 📚 **Manual completo**: Mostra como usar qualquer função, tipo, método
- 🔍 **Busca rápida**: Encontra o que você precisa na hora
- 💡 **Exemplos**: Mostra como usar na prática
- 🎯 **Preciso**: Vai direto ao que você precisa saber

### Exemplos Práticos

**Descobrir o que um pacote faz:**
```bash
go doc fmt
# É como: Ler a capa do manual - "O que este pacote faz?"
```

**Ver como usar uma função específica:**
```bash
go doc fmt.Println
# É como: Ir direto na página do manual que explica "Como usar Println"
```

**Ver o código-fonte (como funciona por dentro):**
```bash
go doc -src fmt.Println
# É como: Abrir a ferramenta e ver como ela funciona por dentro!
```

**Explorar tudo de um pacote:**
```bash
go doc -all fmt
# É como: Ler o manual inteiro do pacote fmt
```

### Quando Usar?

**Você está aprendendo:**
```bash
# "O que o pacote strings faz?"
go doc strings

# "Como usar strings.Contains?"
go doc strings.Contains
```

**Você esqueceu como usar:**
```bash
# "Como formatar uma data de novo?"
go doc time.Time.Format
```

**Você quer verificar sua própria documentação:**
```bash
# "Minha documentação está boa?"
go doc .MinhaFuncao
```

**É como:** Ter um assistente que sempre sabe o manual de qualquer ferramenta e te mostra na hora!

---

## 🔍 9. `go version` - A Etiqueta de Identificação

### A Analogia da Cozinha

Você tem várias **ferramentas** na cozinha, mas precisa saber **qual versão** de cada uma você tem. É como ver a **etiqueta de identificação** de cada ferramenta!

**`go version` é como ler a etiqueta de uma ferramenta:**
- 🏷️ **Identificação**: Mostra exatamente qual versão do Go você tem
- 💻 **Sistema**: Mostra para qual sistema operacional
- 🎯 **Arquitetura**: Mostra para qual tipo de processador
- ✅ **Verificação rápida**: Resolve dúvidas na hora

### Exemplos Práticos

**Ver sua versão:**
```bash
go version
# Output: go version go1.21.5 darwin/arm64

# É como ler a etiqueta:
# - Versão: 1.21.5
# - Sistema: macOS (darwin)
# - Processador: ARM 64-bit (Apple Silicon)
```

**Verificar um binário:**
```bash
go version ./meuapp
# Output: ./meuapp: go1.21.5

# É como: Ver com qual versão da ferramenta foi feito um produto
# "Este app foi compilado com Go 1.21.5"
```

### Quando Usar?

**Troubleshooting:**
```bash
# "Por que meu código não funciona?"
go version
# Talvez você precise atualizar o Go!
```

**Verificar compatibilidade:**
```bash
# "Meu código funciona na versão que tenho?"
go version
# Se você tem Go 1.21, mas o projeto precisa 1.22, vai dar problema!
```

**CI/CD:**
```bash
# Em pipelines, verificar a versão
go version
# Garantir que está usando a versão correta
```

**É como:** Sempre saber qual "modelo" de cada ferramenta você tem, para garantir compatibilidade!

---

## 🎯 Resumo com Analogias

| Comando | É Como... | Quando Usar |
|---------|-----------|-------------|
| `go run` | 🍳 Microondas rápido | Testar código rapidamente |
| `go build` | 🏭 Fábrica de embalagens | Criar produto final |
| `go install` | 🛠️ Loja de ferramentas | Instalar ferramentas globais |
| `go fmt` | ✂️ Organizador automático | Antes de commitar |
| `go mod` | 📚 Biblioteca de receitas | Gerenciar dependências |
| `go test` | 🧪 Laboratório de qualidade | Garantir que funciona |
| `go clean` | 🧹 Limpeza geral | Troubleshooting, liberar espaço |
| `go doc` | 📖 Manual de instruções | Aprender, verificar APIs |
| `go version` | 🏷️ Etiqueta de identificação | Verificar versão |

---

## 🎬 Cena do Dia a Dia: Workflow Completo

Imagine um **dia típico** de desenvolvimento:

### Manhã: Começando um Projeto Novo

```bash
# 1. Criar projeto
mkdir meu-projeto
cd meu-projeto
go mod init github.com/eu/meu-projeto
# 📚 Criou a "biblioteca de receitas" vazia

# 2. Verificar versão
go version
# 🏷️ "Tenho Go 1.21.5, perfeito!"
```

### Tarde: Desenvolvendo

```bash
# 3. Escrever código e testar rapidamente
go run main.go
# 🍳 "Funcionou! Vou continuar desenvolvendo"

# 4. Formatar código
go fmt ./...
# ✂️ "Deixei tudo organizado"

# 5. Testar oficialmente
go test -v
# 🧪 "Todos os testes passaram!"

# 6. Verificar documentação de uma função
go doc fmt.Printf
# 📖 "Ah, entendi como usar!"
```

### Noite: Finalizando

```bash
# 7. Organizar dependências
go mod tidy
# 📚 "Limpei a lista de ingredientes"

# 8. Limpar antes do build final
go clean
# 🧹 "Deixei tudo limpo"

# 9. Build para produção
go build -o meuapp .
# 🏭 "Criei o produto final!"

# 10. Verificar o binário
go version ./meuapp
# 🏷️ "Foi compilado com Go 1.21.5, perfeito!"
```

---

## 💡 Dicas Finais

### Comandos que Você Vai Usar TODO DIA

1. **`go run`** - Para testar código rapidamente
2. **`go fmt`** - Antes de commitar (sempre!)
3. **`go test`** - Para garantir qualidade

### Comandos para Situações Específicas

1. **`go build`** - Quando precisa do binário final
2. **`go install`** - Para instalar ferramentas
3. **`go mod tidy`** - Antes de commitar (limpar dependências)
4. **`go clean`** - Quando algo estranho acontece
5. **`go doc`** - Quando esqueceu como usar algo
6. **`go version`** - Para verificar/troubleshooting

---

## 🎓 Conclusão

Agora você entende cada comando como se fossem **ferramentas da sua cozinha de programação**! Cada uma tem um propósito específico e, juntas, elas formam seu **kit completo de desenvolvimento Go**.

Lembre-se:
- 🍳 **Microondas** (`go run`) para testes rápidos
- 🏭 **Fábrica** (`go build`) para produtos finais
- 🛠️ **Loja** (`go install`) para ferramentas
- ✂️ **Organizador** (`go fmt`) para código limpo
- 📚 **Biblioteca** (`go mod`) para dependências
- 🧪 **Laboratório** (`go test`) para qualidade
- 🧹 **Limpeza** (`go clean`) para troubleshooting
- 📖 **Manual** (`go doc`) para aprender
- 🏷️ **Etiqueta** (`go version`) para verificar

Na próxima parte, vamos colocar a mão na massa com exercícios práticos!

