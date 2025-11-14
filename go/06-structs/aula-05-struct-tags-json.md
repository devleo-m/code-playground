# Módulo 6: Structs em Go

## Aula 5: Struct Tags & JSON

Olá! Agora que você domina structs básicas, vamos aprender sobre **struct tags** - uma funcionalidade poderosa que permite adicionar metadados aos campos de structs. A aplicação mais comum é para **serialização JSON**, mas tags são usadas para muito mais.

---

## 1. O Que São Struct Tags?

**Struct tags** são strings literais que fornecem **metadados** sobre campos de structs. Elas são escritas entre crases (backticks) após o tipo do campo.

### Sintaxe Básica

```go
type Pessoa struct {
    Nome string `tag:"valor"`
}
```

### Exemplo Real: Tag JSON

```go
type Pessoa struct {
    Nome  string `json:"nome"`
    Idade int    `json:"idade"`
    Email string `json:"email,omitempty"`
}
```

**O que isso faz?**

- `json:"nome"` - Quando serializar para JSON, use "nome" como chave
- `json:"idade"` - Use "idade" como chave
- `json:"email,omitempty"` - Use "email" como chave, mas omita se estiver vazio

---

## 2. Por Que Usar Struct Tags?

### Problema: Nomes em Go vs. JSON

**Sem tags:**

```go
type Pessoa struct {
    NomeCompleto string  // Nome em Go (PascalCase)
    Idade        int
}

// JSON gerado:
// {
//   "NomeCompleto": "João",  // Nome em Go, não padrão JSON
//   "Idade": 30
// }
```

**Com tags:**

```go
type Pessoa struct {
    NomeCompleto string `json:"nome_completo"`  // Nome customizado para JSON
    Idade        int    `json:"idade"`
}

// JSON gerado:
// {
//   "nome_completo": "João",  // Nome customizado
//   "idade": 30
// }
```

**Vantagens:**

- Controla nomes de campos em JSON
- Pode omitir campos vazios
- Pode ignorar campos completamente
- Compatível com APIs externas

---

## 3. Tags JSON Básicas

### Tag `json:"nome"`

Define o nome do campo no JSON:

```go
type Pessoa struct {
    Nome string `json:"nome"`
}

pessoa := Pessoa{Nome: "João"}
jsonBytes, _ := json.Marshal(pessoa)
// {"nome":"João"}  ← "nome" em vez de "Nome"
```

### Tag `json:"-"`

Ignora o campo completamente:

```go
type Pessoa struct {
    Nome     string `json:"nome"`
    Senha    string `json:"-"`  // Nunca será serializado!
    Internal int    `json:"-"`  // Campo interno, não expor
}

pessoa := Pessoa{
    Nome:  "João",
    Senha: "secret123",
}
jsonBytes, _ := json.Marshal(pessoa)
// {"nome":"João"}  ← Senha não aparece!
```

### Tag `json:",omitempty"`

Omita o campo se estiver vazio (valor zero):

```go
type Pessoa struct {
    Nome  string `json:"nome"`
    Email string `json:"email,omitempty"`  // Omita se vazio
}

pessoa1 := Pessoa{Nome: "João", Email: "joao@email.com"}
json1, _ := json.Marshal(pessoa1)
// {"nome":"João","email":"joao@email.com"}

pessoa2 := Pessoa{Nome: "Maria"}  // Email vazio
json2, _ := json.Marshal(pessoa2)
// {"nome":"Maria"}  ← Email omitido!
```

### Múltiplas Opções

Você pode combinar opções:

```go
type Pessoa struct {
    Nome     string `json:"nome"`
    Email    string `json:"email,omitempty"`
    Telefone string `json:"telefone,omitempty"`
    Senha    string `json:"-"`  // Ignorar sempre
}
```

---

## 4. Serialização JSON: Marshal

### Convertendo Struct para JSON

```go
package main

import (
    "encoding/json"
    "fmt"
)

type Pessoa struct {
    Nome  string `json:"nome"`
    Idade int    `json:"idade"`
    Email string `json:"email,omitempty"`
}

func main() {
    pessoa := Pessoa{
        Nome:  "João",
        Idade: 30,
        Email: "joao@email.com",
    }

    // Marshal converte struct para JSON ([]byte)
    jsonBytes, err := json.Marshal(pessoa)
    if err != nil {
        fmt.Println("Erro:", err)
        return
    }

    fmt.Println(string(jsonBytes))
    // {"nome":"João","idade":30,"email":"joao@email.com"}
}
```

### Formatação Bonita (Indent)

```go
// MarshalIndent formata o JSON com indentação
jsonBytes, err := json.MarshalIndent(pessoa, "", "  ")
if err != nil {
    fmt.Println("Erro:", err)
    return
}

fmt.Println(string(jsonBytes))
// {
//   "nome": "João",
//   "idade": 30,
//   "email": "joao@email.com"
// }
```

---

## 5. Deserialização JSON: Unmarshal

### Convertendo JSON para Struct

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
    // JSON como string
    jsonString := `{"nome":"João","idade":30,"email":"joao@email.com"}`

    // Converter para []byte
    jsonBytes := []byte(jsonString)

    // Unmarshal converte JSON para struct
    var pessoa Pessoa
    err := json.Unmarshal(jsonBytes, &pessoa)
    if err != nil {
        fmt.Println("Erro:", err)
        return
    }

    fmt.Printf("Nome: %s, Idade: %d, Email: %s\n",
        pessoa.Nome, pessoa.Idade, pessoa.Email)
    // Nome: João, Idade: 30, Email: joao@email.com
}
```

### Campos Não Presentes no JSON

Se um campo não estiver no JSON, ele recebe valor zero:

```go
jsonString := `{"nome":"João"}`  // Sem idade e email

var pessoa Pessoa
json.Unmarshal([]byte(jsonString), &pessoa)

fmt.Println(pessoa.Idade)  // 0 (valor zero)
fmt.Println(pessoa.Email)   // "" (valor zero)
```

---

## 6. Tags JSON Avançadas

### Tipos Customizados

```go
type Status string

const (
    StatusAtivo   Status = "ativo"
    StatusInativo Status = "inativo"
)

type Usuario struct {
    Nome   string `json:"nome"`
    Status Status `json:"status"`
}
```

### Campos Aninhados

```go
type Endereco struct {
    Rua    string `json:"rua"`
    Cidade string `json:"cidade"`
}

type Pessoa struct {
    Nome     string   `json:"nome"`
    Endereco Endereco `json:"endereco"`
}

pessoa := Pessoa{
    Nome: "João",
    Endereco: Endereco{
        Rua:    "Rua A",
        Cidade: "São Paulo",
    },
}

jsonBytes, _ := json.Marshal(pessoa)
// {"nome":"João","endereco":{"rua":"Rua A","cidade":"São Paulo"}}
```

### Slices e Maps

```go
type Pessoa struct {
    Nome   string            `json:"nome"`
    Hobbies []string         `json:"hobbies"`
    Contatos map[string]string `json:"contatos"`
}

pessoa := Pessoa{
    Nome: "João",
    Hobbies: []string{"leitura", "natação"},
    Contatos: map[string]string{
        "email": "joao@email.com",
        "telefone": "123-456",
    },
}

jsonBytes, _ := json.Marshal(pessoa)
// {"nome":"João","hobbies":["leitura","natação"],"contatos":{"email":"joao@email.com","telefone":"123-456"}}
```

---

## 7. Outros Tipos de Tags

### Tag `xml`

Para serialização XML:

```go
type Pessoa struct {
    Nome  string `xml:"nome"`
    Idade int    `xml:"idade"`
}
```

### Tag `db`

Para mapeamento de banco de dados (usado por bibliotecas como GORM):

```go
type Usuario struct {
    ID    int    `db:"id" gorm:"primaryKey"`
    Nome  string `db:"nome" gorm:"not null"`
    Email string `db:"email" gorm:"unique"`
}
```

### Tag `validate`

Para validação (usado por bibliotecas como validator):

```go
type Usuario struct {
    Email string `validate:"required,email"`
    Idade int    `validate:"min=18,max=100"`
}
```

### Tag `yaml`

Para serialização YAML:

```go
type Config struct {
    Porta int    `yaml:"porta"`
    Host  string `yaml:"host"`
}
```

---

## 8. Múltiplas Tags

Você pode usar múltiplas tags separadas por espaço:

```go
type Usuario struct {
    Nome  string `json:"nome" xml:"nome" db:"nome" validate:"required"`
    Email string `json:"email" xml:"email" db:"email" validate:"required,email"`
}
```

---

## 9. Exemplo Completo: API REST

```go
package main

import (
    "encoding/json"
    "fmt"
    "net/http"
)

// Usuario representa um usuário da API
type Usuario struct {
    ID        int    `json:"id"`
    Nome      string `json:"nome"`
    Email     string `json:"email"`
    Senha     string `json:"-"`              // Nunca expor senha
    CriadoEm  string `json:"criado_em,omitempty"`
    Ativo     bool   `json:"ativo"`
}

// CriarUsuario cria um novo usuário
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
    var usuario Usuario

    // Decodificar JSON do request
    decoder := json.NewDecoder(r.Body)
    err := decoder.Decode(&usuario)
    if err != nil {
        http.Error(w, "JSON inválido", http.StatusBadRequest)
        return
    }

    // Validar e processar...
    usuario.ID = 1
    usuario.Ativo = true

    // Responder com JSON
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(usuario)
}

func main() {
    // Exemplo de uso
    usuario := Usuario{
        Nome:  "João",
        Email: "joao@email.com",
        Senha: "secret123",  // Não será serializado
    }

    jsonBytes, _ := json.Marshal(usuario)
    fmt.Println(string(jsonBytes))
    // {"id":0,"nome":"João","email":"joao@email.com","ativo":false}
    // Senha não aparece!
}
```

---

## 10. Tratamento de Erros

### Erros Comuns

```go
// Erro: JSON inválido
jsonString := `{"nome":"João"`  // JSON malformado
var pessoa Pessoa
err := json.Unmarshal([]byte(jsonString), &pessoa)
if err != nil {
    fmt.Println("Erro ao decodificar:", err)
    // Erro ao decodificar: unexpected end of JSON input
}

// Erro: Tipo incompatível
jsonString := `{"idade":"trinta"}`  // String em vez de int
var pessoa Pessoa
err := json.Unmarshal([]byte(jsonString), &pessoa)
if err != nil {
    fmt.Println("Erro:", err)
    // Erro: json: cannot unmarshal string into struct field Pessoa.Idade of type int
}
```

### Validação Após Unmarshal

```go
func ValidarPessoa(p Pessoa) error {
    if p.Nome == "" {
        return fmt.Errorf("nome é obrigatório")
    }
    if p.Idade < 0 {
        return fmt.Errorf("idade deve ser positiva")
    }
    return nil
}

// Uso
var pessoa Pessoa
json.Unmarshal(jsonBytes, &pessoa)
if err := ValidarPessoa(pessoa); err != nil {
    fmt.Println("Validação falhou:", err)
}
```

---

## 11. Boas Práticas com Tags JSON

### ✅ BOM:

```go
type Usuario struct {
    ID        int    `json:"id"`
    Nome      string `json:"nome"`
    Email     string `json:"email,omitempty"`
    Senha     string `json:"-"`  // Nunca expor
    CriadoEm  time.Time `json:"criado_em"`
}
```

### ❌ RUIM:

```go
// Sem tags - nomes em Go aparecem no JSON
type Usuario struct {
    ID       int
    Nome     string
    Email    string
    Senha    string  // Senha exposta!
    CriadoEm time.Time
}
```

---

## 12. Resumo das Tags JSON

| Tag                     | Descrição                            |
| ----------------------- | ------------------------------------ |
| `json:"nome"`           | Define nome do campo no JSON         |
| `json:"-"`              | Ignora o campo completamente         |
| `json:",omitempty"`     | Omita se campo estiver vazio         |
| `json:"nome,omitempty"` | Combina nome customizado + omitempty |

---

## 13. Exemplo Prático Completo

```go
package main

import (
    "encoding/json"
    "fmt"
    "time"
)

type Endereco struct {
    Rua    string `json:"rua"`
    Cidade string `json:"cidade"`
    CEP    string `json:"cep,omitempty"`
}

type Usuario struct {
    ID        int       `json:"id"`
    Nome      string    `json:"nome"`
    Email     string    `json:"email"`
    Senha     string    `json:"-"`  // Nunca serializar
    Ativo     bool      `json:"ativo"`
    Endereco  Endereco  `json:"endereco,omitempty"`
    CriadoEm  time.Time `json:"criado_em"`
}

func main() {
    // Criar usuário
    usuario := Usuario{
        ID:       1,
        Nome:     "João Silva",
        Email:    "joao@email.com",
        Senha:    "senha123",  // Não será serializado
        Ativo:    true,
        Endereco: Endereco{
            Rua:    "Rua A",
            Cidade: "São Paulo",
            // CEP omitido (vazio)
        },
        CriadoEm: time.Now(),
    }

    // Serializar para JSON
    jsonBytes, err := json.MarshalIndent(usuario, "", "  ")
    if err != nil {
        fmt.Println("Erro:", err)
        return
    }

    fmt.Println("JSON serializado:")
    fmt.Println(string(jsonBytes))

    // Deserializar de JSON
    jsonString := `{
        "id": 2,
        "nome": "Maria",
        "email": "maria@email.com",
        "ativo": true,
        "criado_em": "2024-01-01T00:00:00Z"
    }`

    var usuario2 Usuario
    err = json.Unmarshal([]byte(jsonString), &usuario2)
    if err != nil {
        fmt.Println("Erro:", err)
        return
    }

    fmt.Printf("\nUsuário deserializado:\n")
    fmt.Printf("ID: %d\n", usuario2.ID)
    fmt.Printf("Nome: %s\n", usuario2.Nome)
    fmt.Printf("Email: %s\n", usuario2.Email)
    fmt.Printf("Senha: %s (não veio do JSON)\n", usuario2.Senha)
}
```

---

## Conclusão

Struct tags são essenciais para:

1. **Serialização JSON** - Controlar nomes e comportamento
2. **APIs REST** - Formatar respostas corretamente
3. **Segurança** - Omitir campos sensíveis (senhas, tokens)
4. **Integração** - Compatibilidade com APIs externas
5. **Flexibilidade** - Múltiplos formatos (JSON, XML, YAML)

Lembre-se:

- Use tags para controlar serialização
- Sempre omita campos sensíveis com `json:"-"`
- Use `omitempty` para campos opcionais
- Valide dados após deserialização

Na próxima aula, vamos aprender sobre **Embedding Structs** - uma forma poderosa de composição em Go!
