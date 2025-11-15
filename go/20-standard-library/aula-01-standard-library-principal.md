# Módulo 20: Standard Library do Go
## Aula 1: Standard Library - A Caixa de Ferramentas do Go

Olá! Bem-vindo ao módulo 20. Você já aprendeu muito sobre Go: tipos, estruturas, concorrência, contextos e muito mais. Agora chegou a hora de conhecer uma das maiores forças do Go: sua **Standard Library** (Biblioteca Padrão).

A Standard Library do Go é uma coleção abrangente de pacotes que fornecem funcionalidades essenciais para praticamente qualquer tipo de aplicação. Ela é tão completa e bem projetada que muitos desenvolvedores conseguem construir aplicações complexas usando apenas a biblioteca padrão, sem precisar de dependências externas.

---

## O que é a Standard Library?

A Standard Library do Go é um conjunto de pacotes que vêm **prontos para uso** quando você instala o Go. Não é necessário instalar nada adicional via `go get`. Ela inclui:

- **I/O e manipulação de arquivos**: Ler e escrever arquivos, trabalhar com streams
- **Rede**: Cliente e servidor HTTP, TCP/UDP
- **Processamento de texto**: Strings, regex, encoding
- **Criptografia**: Hash, assinaturas, criptografia simétrica e assimétrica
- **Testes**: Framework completo de testes
- **JSON/XML/CSV**: Serialização e deserialização de dados
- **Tempo**: Operações com datas e horas
- **Argumentos de linha de comando**: Parsing de flags
- **Logging**: Sistema de logs estruturado
- **E muito mais!**

A filosofia do Go é: "Se é comum, deve estar na biblioteca padrão". Isso significa menos dependências, menos problemas de compatibilidade e código mais confiável.

---

## 1. I/O e File Handling

O Go fornece um sistema de I/O baseado em **interfaces**, o que torna o código flexível e testável. As interfaces principais são:

### Interfaces Fundamentais

```go
// io.Reader - lê dados de uma fonte
type Reader interface {
    Read(p []byte) (n int, err error)
}

// io.Writer - escreve dados em um destino
type Writer interface {
    Write(p []byte) (n int, err error)
}

// io.Closer - fecha recursos
type Closer interface {
    Close() error
}
```

### Trabalhando com Arquivos

O pacote `os` fornece operações de arquivo:

```go
package main

import (
    "fmt"
    "io"
    "os"
)

func main() {
    // Criar um arquivo
    arquivo, err := os.Create("exemplo.txt")
    if err != nil {
        fmt.Printf("Erro ao criar arquivo: %v\n", err)
        return
    }
    defer arquivo.Close() // Sempre feche o arquivo!

    // Escrever no arquivo
    conteudo := "Olá, mundo Go!\nEsta é uma linha de teste."
    bytesEscritos, err := arquivo.WriteString(conteudo)
    if err != nil {
        fmt.Printf("Erro ao escrever: %v\n", err)
        return
    }
    fmt.Printf("Escritos %d bytes\n", bytesEscritos)

    // Ler o arquivo
    arquivoLido, err := os.Open("exemplo.txt")
    if err != nil {
        fmt.Printf("Erro ao abrir arquivo: %v\n", err)
        return
    }
    defer arquivoLido.Close()

    dados, err := io.ReadAll(arquivoLido)
    if err != nil {
        fmt.Printf("Erro ao ler: %v\n", err)
        return
    }
    fmt.Printf("Conteúdo lido: %s\n", string(dados))
}
```

### Operações de Diretório

```go
package main

import (
    "fmt"
    "os"
    "path/filepath"
)

func main() {
    // Criar diretório
    err := os.Mkdir("meu_diretorio", 0755)
    if err != nil {
        fmt.Printf("Erro: %v\n", err)
    }

    // Criar diretórios aninhados
    err = os.MkdirAll("dir1/dir2/dir3", 0755)
    if err != nil {
        fmt.Printf("Erro: %v\n", err)
    }

    // Listar arquivos em um diretório
    entradas, err := os.ReadDir(".")
    if err != nil {
        fmt.Printf("Erro: %v\n", err)
        return
    }

    for _, entrada := range entradas {
        fmt.Printf("Nome: %s, É diretório: %v\n", entrada.Name(), entrada.IsDir())
    }

    // Caminhos com filepath
    caminho := filepath.Join("dir1", "dir2", "arquivo.txt")
    fmt.Printf("Caminho: %s\n", caminho)
}
```

---

## 2. O Pacote `flag` - Argumentos de Linha de Comando

O pacote `flag` permite criar ferramentas de linha de comando profissionais com parsing automático de argumentos.

### Uso Básico

```go
package main

import (
    "flag"
    "fmt"
)

func main() {
    // Definir flags
    nome := flag.String("nome", "Visitante", "Seu nome")
    idade := flag.Int("idade", 0, "Sua idade")
    ativo := flag.Bool("ativo", false, "Está ativo?")
    
    // Parse dos argumentos
    flag.Parse()

    fmt.Printf("Olá, %s!\n", *nome)
    fmt.Printf("Idade: %d\n", *idade)
    fmt.Printf("Ativo: %v\n", *ativo)
    
    // Argumentos posicionais (não flags)
    fmt.Printf("Argumentos restantes: %v\n", flag.Args())
}
```

**Uso:**
```bash
go run main.go -nome="João" -idade=25 -ativo
```

### Flags com Duration

```go
package main

import (
    "flag"
    "fmt"
    "time"
)

func main() {
    duracao := flag.Duration("timeout", 5*time.Second, "Timeout da operação")
    flag.Parse()

    fmt.Printf("Timeout configurado: %v\n", *duracao)
}
```

**Uso:**
```bash
go run main.go -timeout=10s
go run main.go -timeout=2m30s
```

---

## 3. O Pacote `time` - Trabalhando com Tempo

O pacote `time` é essencial para qualquer aplicação que trabalhe com datas, horas, timers e timeouts.

### Tipos Principais

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    // Time - representa um momento específico
    agora := time.Now()
    fmt.Printf("Agora: %v\n", agora)
    fmt.Printf("Formato RFC3339: %s\n", agora.Format(time.RFC3339))

    // Duration - representa uma duração de tempo
    duracao := 2 * time.Hour + 30 * time.Minute + 45 * time.Second
    fmt.Printf("Duração: %v\n", duracao)
    fmt.Printf("Em segundos: %.0f\n", duracao.Seconds())

    // Criar um Time específico
    dataEspecifica := time.Date(2024, time.January, 15, 10, 30, 0, 0, time.UTC)
    fmt.Printf("Data específica: %v\n", dataEspecifica)

    // Parsing de strings
    tempoStr := "2024-01-15T10:30:00Z"
    tempoParsed, err := time.Parse(time.RFC3339, tempoStr)
    if err != nil {
        fmt.Printf("Erro: %v\n", err)
        return
    }
    fmt.Printf("Tempo parseado: %v\n", tempoParsed)

    // Operações aritméticas
    futuro := agora.Add(24 * time.Hour)
    passado := agora.Add(-7 * 24 * time.Hour)
    fmt.Printf("Amanhã: %v\n", futuro)
    fmt.Printf("Há uma semana: %v\n", passado)

    // Comparações
    if agora.Before(futuro) {
        fmt.Println("Agora é antes do futuro")
    }

    // Timers e Tickers
    timer := time.NewTimer(2 * time.Second)
    <-timer.C
    fmt.Println("Timer expirou!")

    ticker := time.NewTicker(1 * time.Second)
    go func() {
        for t := range ticker.C {
            fmt.Printf("Tick em %v\n", t)
        }
    }()
    time.Sleep(5 * time.Second)
    ticker.Stop()
    fmt.Println("Ticker parado")
}
```

### Timezones

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    // UTC
    utc := time.Now().UTC()
    fmt.Printf("UTC: %v\n", utc)

    // Timezone específico
    loc, err := time.LoadLocation("America/Sao_Paulo")
    if err != nil {
        fmt.Printf("Erro: %v\n", err)
        return
    }
    sp := time.Now().In(loc)
    fmt.Printf("São Paulo: %v\n", sp)

    // Formatação customizada
    layout := "02/01/2006 15:04:05"
    formatado := time.Now().Format(layout)
    fmt.Printf("Formatado: %s\n", formatado)
}
```

---

## 4. Encoding/JSON - Trabalhando com JSON

O pacote `encoding/json` fornece serialização e deserialização eficiente de JSON.

### Marshal (Go → JSON)

```go
package main

import (
    "encoding/json"
    "fmt"
)

type Pessoa struct {
    Nome    string `json:"nome"`
    Idade   int    `json:"idade"`
    Email   string `json:"email,omitempty"` // omitempty omite se vazio
    Ativo   bool   `json:"ativo"`
    Salario float64 `json:"-"` // O "-" omite o campo completamente
}

func main() {
    pessoa := Pessoa{
        Nome:  "João Silva",
        Idade: 30,
        Email: "joao@example.com",
        Ativo: true,
        Salario: 5000.00,
    }

    // Marshal (struct → JSON)
    jsonBytes, err := json.Marshal(pessoa)
    if err != nil {
        fmt.Printf("Erro: %v\n", err)
        return
    }
    fmt.Printf("JSON: %s\n", string(jsonBytes))

    // Marshal com indentação (mais legível)
    jsonIndentado, err := json.MarshalIndent(pessoa, "", "  ")
    if err != nil {
        fmt.Printf("Erro: %v\n", err)
        return
    }
    fmt.Printf("JSON indentado:\n%s\n", string(jsonIndentado))
}
```

### Unmarshal (JSON → Go)

```go
package main

import (
    "encoding/json"
    "fmt"
)

type Pessoa struct {
    Nome  string `json:"nome"`
    Idade int    `json:"idade"`
    Email string `json:"email"`
}

func main() {
    jsonStr := `{"nome":"Maria","idade":25,"email":"maria@example.com"}`

    var pessoa Pessoa
    err := json.Unmarshal([]byte(jsonStr), &pessoa)
    if err != nil {
        fmt.Printf("Erro: %v\n", err)
        return
    }

    fmt.Printf("Nome: %s\n", pessoa.Nome)
    fmt.Printf("Idade: %d\n", pessoa.Idade)
    fmt.Printf("Email: %s\n", pessoa.Email)
}
```

### Decoder e Encoder (Streaming)

Para arquivos grandes ou streams:

```go
package main

import (
    "encoding/json"
    "fmt"
    "strings"
)

func main() {
    // Simulando um stream de dados
    jsonStream := `{"nome":"João","idade":30}
{"nome":"Maria","idade":25}
{"nome":"Pedro","idade":35}`

    decoder := json.NewDecoder(strings.NewReader(jsonStream))

    for {
        var pessoa map[string]interface{}
        if err := decoder.Decode(&pessoa); err != nil {
            break // Fim do stream
        }
        fmt.Printf("Pessoa: %v\n", pessoa)
    }
}
```

---

## 5. O Pacote `os` - Interface com o Sistema Operacional

O pacote `os` fornece uma interface para funcionalidades do sistema operacional.

### Variáveis de Ambiente

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    // Obter variável de ambiente
    path := os.Getenv("PATH")
    fmt.Printf("PATH: %s\n", path)

    // Definir variável de ambiente (apenas para este processo)
    os.Setenv("MINHA_VAR", "meu_valor")
    fmt.Printf("MINHA_VAR: %s\n", os.Getenv("MINHA_VAR"))

    // Listar todas as variáveis
    for _, env := range os.Environ() {
        fmt.Println(env)
    }
}
```

### Informações do Sistema

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    // Hostname
    hostname, err := os.Hostname()
    if err != nil {
        fmt.Printf("Erro: %v\n", err)
    } else {
        fmt.Printf("Hostname: %s\n", hostname)
    }

    // Diretório de trabalho atual
    dir, err := os.Getwd()
    if err != nil {
        fmt.Printf("Erro: %v\n", err)
    } else {
        fmt.Printf("Diretório atual: %s\n", dir)
    }

    // Mudar diretório
    err = os.Chdir("/tmp")
    if err != nil {
        fmt.Printf("Erro: %v\n", err)
    }

    // Informações do usuário
    fmt.Printf("UID: %d\n", os.Getuid())
    fmt.Printf("GID: %d\n", os.Getgid())
}
```

### Gerenciamento de Processos

```go
package main

import (
    "fmt"
    "os"
    "os/exec"
)

func main() {
    // PID do processo atual
    fmt.Printf("PID: %d\n", os.Getpid())

    // Executar comando externo
    cmd := exec.Command("ls", "-la")
    output, err := cmd.Output()
    if err != nil {
        fmt.Printf("Erro: %v\n", err)
        return
    }
    fmt.Printf("Saída:\n%s\n", string(output))

    // Sair do programa
    // os.Exit(0) // Sucesso
    // os.Exit(1) // Erro
}
```

---

## 6. O Pacote `bufio` - I/O com Buffer

O `bufio` fornece I/O com buffer, melhorando a performance ao reduzir chamadas de sistema.

### Scanner - Leitura Linha por Linha

```go
package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    // Ler do stdin linha por linha
    scanner := bufio.NewScanner(os.Stdin)
    fmt.Println("Digite linhas (Ctrl+D para terminar):")
    for scanner.Scan() {
        linha := scanner.Text()
        fmt.Printf("Você digitou: %s\n", linha)
    }

    if err := scanner.Err(); err != nil {
        fmt.Printf("Erro: %v\n", err)
    }

    // Ler de uma string
    texto := "linha 1\nlinha 2\nlinha 3"
    scanner2 := bufio.NewScanner(strings.NewReader(texto))
    for scanner2.Scan() {
        fmt.Printf("Linha: %s\n", scanner2.Text())
    }
}
```

### Reader e Writer com Buffer

```go
package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    // Reader com buffer
    texto := "Este é um texto de exemplo para leitura com buffer"
    reader := bufio.NewReader(strings.NewReader(texto))

    // Ler até encontrar um delimitador
    parte, err := reader.ReadString(' ')
    if err != nil {
        fmt.Printf("Erro: %v\n", err)
    } else {
        fmt.Printf("Primeira palavra: %s\n", parte)
    }

    // Writer com buffer
    arquivo, err := os.Create("buffer_exemplo.txt")
    if err != nil {
        fmt.Printf("Erro: %v\n", err)
        return
    }
    defer arquivo.Close()

    writer := bufio.NewWriter(arquivo)
    writer.WriteString("Linha 1\n")
    writer.WriteString("Linha 2\n")
    writer.WriteString("Linha 3\n")
    
    // IMPORTANTE: Flush para garantir que tudo foi escrito
    writer.Flush()
    fmt.Println("Dados escritos com buffer!")
}
```

---

## 7. O Pacote `slog` - Logging Estruturado

O `slog` (introduzido no Go 1.21) é o sistema moderno de logging estruturado do Go.

### Uso Básico

```go
package main

import (
    "log/slog"
    "os"
)

func main() {
    // Logger padrão (texto simples)
    logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
    logger.Info("Mensagem de informação")
    logger.Warn("Aviso!")
    logger.Error("Erro ocorreu", "codigo", 500)

    // Logger com JSON (melhor para produção)
    jsonLogger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
    jsonLogger.Info("Usuário logado",
        "usuario", "joao",
        "ip", "192.168.1.1",
        "timestamp", "2024-01-15T10:30:00Z")
}
```

### Logger com Níveis e Contexto

```go
package main

import (
    "context"
    "log/slog"
    "os"
)

func main() {
    // Configurar logger com nível mínimo
    opts := &slog.HandlerOptions{
        Level: slog.LevelDebug,
    }
    logger := slog.New(slog.NewJSONHandler(os.Stdout, opts))

    // Diferentes níveis
    logger.Debug("Debug: informação detalhada")
    logger.Info("Info: informação geral")
    logger.Warn("Warn: aviso")
    logger.Error("Error: erro ocorreu")

    // Com contexto
    ctx := context.Background()
    logger.InfoContext(ctx, "Operação concluída",
        "operacao", "criar_usuario",
        "sucesso", true)
}
```

---

## 8. O Pacote `regexp` - Expressões Regulares

O `regexp` implementa a sintaxe RE2, que é segura e eficiente.

### Match e Find

```go
package main

import (
    "fmt"
    "regexp"
)

func main() {
    texto := "Meu email é joao@example.com e meu telefone é (11) 98765-4321"

    // Verificar se há correspondência
    emailPattern := `\b[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Z|a-z]{2,}\b`
    matched, err := regexp.MatchString(emailPattern, texto)
    if err != nil {
        fmt.Printf("Erro: %v\n", err)
        return
    }
    fmt.Printf("Tem email? %v\n", matched)

    // Compilar regex (mais eficiente para uso múltiplo)
    re := regexp.MustCompile(emailPattern)
    
    // Encontrar primeira correspondência
    email := re.FindString(texto)
    fmt.Printf("Email encontrado: %s\n", email)

    // Encontrar todas as correspondências
    emails := re.FindAllString(texto, -1)
    fmt.Printf("Todos os emails: %v\n", emails)

    // Encontrar com grupos de captura
    telefonePattern := `\((\d{2})\)\s(\d{5})-(\d{4})`
    reTel := regexp.MustCompile(telefonePattern)
    matches := reTel.FindStringSubmatch(texto)
    if len(matches) > 0 {
        fmt.Printf("DDD: %s\n", matches[1])
        fmt.Printf("Número: %s-%s\n", matches[2], matches[3])
    }
}
```

### Replace

```go
package main

import (
    "fmt"
    "regexp"
)

func main() {
    texto := "Meus telefones são (11) 98765-4321 e (21) 91234-5678"

    // Substituir todas as ocorrências
    re := regexp.MustCompile(`\(\d{2}\)\s\d{5}-\d{4}`)
    substituido := re.ReplaceAllString(texto, "[TELEFONE OCULTO]")
    fmt.Printf("Texto substituído: %s\n", substituido)

    // Substituir com função
    resultado := re.ReplaceAllStringFunc(texto, func(s string) string {
        return "***" + s[len(s)-4:] // Mostra apenas últimos 4 dígitos
    })
    fmt.Printf("Texto mascarado: %s\n", resultado)
}
```

---

## 9. `go:embed` - Embedding de Arquivos

O `go:embed` permite incluir arquivos e diretórios diretamente no binário em tempo de compilação.

### Embedding de Arquivo Único

```go
package main

import (
    _ "embed"
    "fmt"
)

//go:embed config.txt
var configContent string

func main() {
    fmt.Printf("Conteúdo do arquivo:\n%s\n", configContent)
}
```

### Embedding de Múltiplos Arquivos

```go
package main

import (
    "embed"
    "fmt"
)

//go:embed templates/*.html
var templatesFS embed.FS

func main() {
    // Listar arquivos
    entries, err := templatesFS.ReadDir("templates")
    if err != nil {
        fmt.Printf("Erro: %v\n", err)
        return
    }

    for _, entry := range entries {
        fmt.Printf("Arquivo: %s\n", entry.Name())
        if !entry.IsDir() {
            conteudo, err := templatesFS.ReadFile("templates/" + entry.Name())
            if err != nil {
                fmt.Printf("Erro ao ler: %v\n", err)
                continue
            }
            fmt.Printf("Conteúdo: %s\n", string(conteudo))
        }
    }
}
```

### Embedding de Diretório Completo

```go
package main

import (
    "embed"
    "fmt"
    "io/fs"
    "path/filepath"
)

//go:embed static
var staticFiles embed.FS

func main() {
    // Caminhar por todos os arquivos
    fs.WalkDir(staticFiles, ".", func(path string, d fs.DirEntry, err error) error {
        if err != nil {
            return err
        }
        if !d.IsDir() {
            fmt.Printf("Arquivo embutido: %s\n", path)
        }
        return nil
    })
}
```

---

## Resumo

Nesta aula, você conheceu os principais pacotes da Standard Library do Go:

1. **I/O e File Handling**: Interfaces `Reader`/`Writer`, operações com arquivos e diretórios
2. **flag**: Parsing de argumentos de linha de comando
3. **time**: Operações com tempo, datas, timers e timezones
4. **encoding/json**: Serialização e deserialização de JSON
5. **os**: Interface com o sistema operacional
6. **bufio**: I/O com buffer para melhor performance
7. **slog**: Sistema moderno de logging estruturado
8. **regexp**: Expressões regulares seguras e eficientes
9. **go:embed**: Inclusão de arquivos no binário

A Standard Library do Go é uma das maiores forças da linguagem. Dominá-la permite construir aplicações robustas sem depender de bibliotecas externas, resultando em binários menores, mais seguros e mais fáceis de manter.

Na próxima aula, vamos simplificar esses conceitos com analogias do dia a dia para fixar o aprendizado!

