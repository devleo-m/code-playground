# Aula 2 - Simplificada: Entendendo Code Organization em Go

Olá! Na aula anterior, mergulhamos nos detalhes técnicos da organização de código em Go. Agora, vamos simplificar tudo isso usando analogias do dia a dia para que você realmente **entenda** e não apenas **decorar** os conceitos.

---

## 1. Go Modules: A Analogia da Biblioteca

Imagine que você está escrevendo um livro. Em vez de escrever tudo do zero, você quer referenciar outros livros que já foram escritos.

**Antes dos Go Modules (GOPATH):**
```
Você: "Preciso do livro 'Matemática Avançada'"
Bibliotecário: "Todos os livros devem estar na mesma estante gigante. 
                Você não pode ter sua própria estante."
```

**Com Go Modules:**
```
Você: "Preciso do livro 'Matemática Avançada'"
Bibliotecário: "Cada projeto tem sua própria estante. 
                Você pode ter versões diferentes do mesmo livro em projetos diferentes."
```

O arquivo `go.mod` é como o **índice do seu livro** - ele lista todos os outros livros (dependências) que você está usando e suas versões.

---

## 2. `go mod init`: A Analogia do Primeiro Passo

Criar um novo projeto Go é como começar a escrever um novo livro. O primeiro passo é criar a capa com o título.

**Analogia:**
```bash
# Você está começando um novo livro
go mod init github.com/seu-usuario/meu-livro
```

É como dizer: "Este é o título do meu livro e onde ele será publicado". O Go cria um arquivo `go.mod` (a capa) com essa informação.

**Por que usar URL do repositório?**
- É como colocar o endereço da editora na capa do livro
- Outros podem encontrar e usar seu código facilmente
- Mesmo que seja um projeto local, é uma boa prática

---

## 3. `go mod tidy`: A Analogia da Limpeza de Casa

Imagine que você tem uma lista de compras, mas com o tempo:
- Você comprou coisas que não estão na lista
- Você tem coisas na lista que nunca comprou

O `go mod tidy` é como **organizar sua lista de compras**:

**Antes:**
```
Lista de Compras:
- Leite ✓ (você comprou)
- Pão ✓ (você comprou)
- Queijo (está na lista, mas você nunca comprou)
- [Falta] Arroz (você comprou, mas não está na lista)
```

**Depois do `go mod tidy`:**
```
Lista de Compras (organizada):
- Leite ✓
- Pão ✓
- Arroz ✓
(Queijo foi removido porque você não usa)
```

**Quando fazer essa limpeza?**
- Antes de fazer commit (antes de sair de casa)
- Depois de adicionar ou remover código
- Regularmente para manter tudo organizado

---

## 4. `go mod vendor`: A Analogia da Mochila de Viagem

Imagine que você vai fazer uma viagem para um lugar sem internet. Você precisa levar **tudo** que vai precisar na sua mochila.

**Sem vendor:**
```
Você: "Preciso de um mapa"
Guia: "Baixe do aplicativo" (precisa de internet)
Você: "Mas não tenho internet aqui!"
```

**Com vendor:**
```
Você: "Preciso de um mapa"
Guia: "Está na sua mochila (vendor/)"
Você: "Perfeito! Tenho tudo que preciso!"
```

O `go mod vendor` cria uma "mochila" (`vendor/`) com **cópias** de todas as dependências. Assim, você pode compilar seu programa mesmo sem internet.

**Quando usar?**
- ✅ Quando você vai fazer deploy em servidor sem internet
- ✅ Quando você quer garantir versões exatas
- ✅ Em ambientes muito restritivos (air-gapped)

**Quando NÃO usar?**
- ❌ Para desenvolvimento normal (desnecessário)
- ❌ Quando você tem internet estável

---

## 5. Packages: A Analogia dos Departamentos de uma Empresa

Pense em uma empresa grande. Ela está organizada em **departamentos** (packages):

```
Empresa (Projeto)
├── Departamento de Vendas (package vendas)
├── Departamento de RH (package rh)
├── Departamento de TI (package ti)
└── Diretoria (package main)
```

Cada departamento:
- Tem suas próprias funções e responsabilidades
- Pode ter coisas públicas (qualquer um pode acessar) e privadas (só o departamento acessa)
- Trabalha de forma independente, mas pode colaborar

### Package `main`: A Diretoria

O package `main` é especial - é como a **diretoria** da empresa:
- É o ponto de entrada (onde tudo começa)
- Não pode ser "importado" por outros (outras empresas não podem ter acesso à sua diretoria)
- Deve ter uma função `main()` (o CEO que inicia tudo)

```go
package main  // Diretoria

func main() {  // CEO inicia a empresa
    // código aqui
}
```

### Outros Packages: Os Departamentos

```go
// package vendas (Departamento de Vendas)
package vendas

// Função pública - qualquer departamento pode usar
func CalcularComissao(valor float64) float64 {
    return valor * 0.1
}

// Função privada - só o departamento de vendas usa
func calcularDesconto(valor float64) float64 {
    return valor * 0.05
}
```

---

## 6. Exportação: A Analogia da Porta da Empresa

Em Go, se algo começa com **maiúscula**, é como ter uma **porta aberta** - qualquer um pode entrar. Se começa com **minúscula**, é como ter uma **porta fechada** - só quem está dentro pode acessar.

```go
package exemplo

// ✅ PORTA ABERTA (Maiúscula) - Qualquer um pode usar
func FuncaoPublica() {
    // Qualquer package pode chamar isso
}

// ❌ PORTA FECHADA (Minúscula) - Só este package pode usar
func funcaoPrivada() {
    // Só código dentro deste package pode chamar
}

// ✅ Público
var ConstantePublica = 42

// ❌ Privado
var constantePrivada = 100
```

**Regra de Ouro:**
- 🟢 **Maiúscula** = Público = Exportado = Outros podem usar
- 🔴 **Minúscula** = Privado = Não exportado = Só uso interno

---

## 7. Imports: A Analogia de Pedir Ajuda a Outros Departamentos

Quando um departamento precisa de algo de outro departamento, ele "importa" essa ajuda:

```go
package vendas

import (
    "rh"        // Preciso do departamento de RH
    "ti"        // Preciso do departamento de TI
    "fmt"       // Preciso da biblioteca padrão (como serviços externos)
)
```

### Import Circular: O Problema

**Analogia:** Imagine que:
- O departamento de Vendas precisa do RH
- O RH precisa do departamento de Vendas

Isso cria um **ciclo** - nenhum dos dois pode funcionar porque cada um está esperando o outro!

```
Vendas: "Preciso do RH para funcionar"
RH: "Preciso do Vendas para funcionar"
Resultado: Ninguém funciona! ❌
```

**Solução:** Reorganize! Crie um departamento intermediário ou mova a funcionalidade compartilhada.

---

## 8. Packages de Terceiros: A Analogia de Contratar Serviços Externos

Às vezes, sua empresa precisa de serviços que outras empresas já fazem bem. Em vez de criar do zero, você **contrata** (importa) esses serviços.

**Analogia:**
```
Sua Empresa: "Preciso de um sistema de pagamento"
Você: "Vou criar do zero" (leva meses)
OU
Você: "Vou usar o Stripe" (já existe, testado, confiável)
```

```go
import "github.com/stripe/stripe-go"  // Contratando o serviço Stripe
```

### Como Escolher um Serviço (Package)?

**Checklist:**
1. ✅ **Está sendo mantido?** (Última atualização recente)
2. ✅ **Tem documentação?** (Manual de uso claro)
3. ✅ **É confiável?** (Muitos usam, poucos problemas)
4. ✅ **Licença compatível?** (Você pode usar legalmente)

É como escolher um fornecedor para sua empresa - você quer alguém confiável e bem documentado!

---

## 9. Publicar Módulos: A Analogia de Publicar um Livro

Você escreveu um livro incrível e quer que outras pessoas possam lê-lo. Como fazer?

### Passo 1: Preparar o Livro

```bash
go mod init github.com/seu-usuario/meu-livro
# Escrever o código...
go mod tidy  # Organizar tudo
```

### Passo 2: Publicar no Git (Editora)

```bash
git add .
git commit -m "Versão inicial"
git push
```

### Passo 3: Criar uma Edição (Tag de Versão)

```bash
git tag v1.0.0  # Primeira edição
git push origin v1.0.0
```

### Passo 4: O Mundo Descobre Automaticamente

O Go Proxy (como uma biblioteca nacional) automaticamente descobre seu livro e o disponibiliza. Agora qualquer pessoa pode fazer:

```bash
go get github.com/seu-usuario/meu-livro@v1.0.0
```

É como publicar um livro - você coloca na editora (Git), cria uma edição (tag), e as bibliotecas (Go Proxy) automaticamente o disponibilizam!

---

## 10. Estrutura de Projeto: A Analogia do Organizador de Arquivos

Imagine que você tem uma gaveta de arquivos. Você pode jogar tudo dentro de forma bagunçada, OU organizar em pastas:

**Bagunçado:**
```
Gaveta/
├── documento1.txt
├── documento2.txt
├── foto1.jpg
├── documento3.txt
├── foto2.jpg
└── ... (tudo misturado)
```

**Organizado:**
```
Gaveta/
├── Documentos/
│   ├── documento1.txt
│   ├── documento2.txt
│   └── documento3.txt
├── Fotos/
│   ├── foto1.jpg
│   └── foto2.jpg
└── ...
```

**Estrutura Recomendada em Go:**

```
meu-projeto/          (Gaveta principal)
├── cmd/              (Executáveis - como "Aplicativos Prontos")
│   └── server/       (Servidor web)
├── pkg/              (Código reutilizável - como "Biblioteca")
│   └── utils/        (Utilitários)
├── internal/         (Código privado - como "Arquivos Confidenciais")
│   └── database/
└── main.go            (Ponto de entrada)
```

**Analogia:**
- **`cmd/`**: Aplicativos prontos para usar (como programas instalados)
- **`pkg/`**: Ferramentas que outros podem usar (como uma biblioteca pública)
- **`internal/`**: Coisas privadas que ninguém de fora pode ver (como documentos confidenciais)

---

## Resumo com Analogias

1. **Go Modules** = Sistema de biblioteca onde cada projeto tem sua própria estante
2. **`go mod init`** = Criar a capa do livro com título
3. **`go mod tidy`** = Organizar lista de compras
4. **`go mod vendor`** = Mochila com tudo que você precisa para viagem sem internet
5. **Packages** = Departamentos de uma empresa
6. **Exportação** = Portas abertas (maiúscula) vs fechadas (minúscula)
7. **Imports** = Pedir ajuda de outros departamentos
8. **Packages de Terceiros** = Contratar serviços externos
9. **Publicar Módulos** = Publicar um livro
10. **Estrutura de Projeto** = Organizador de arquivos

---

## Dica Final: Comece Simples

Não precisa criar uma estrutura complexa desde o início! Comece simples:

```
meu-projeto/
├── go.mod
├── main.go
└── utils.go
```

Conforme seu projeto cresce, você pode reorganizar. É melhor ter código funcionando simples do que uma estrutura perfeita sem código!

**Lembre-se:** A organização vem com a prática. Quanto mais você codifica, mais natural fica organizar seu código.

Até a próxima aula, onde vamos colocar tudo isso em prática com exercícios!

