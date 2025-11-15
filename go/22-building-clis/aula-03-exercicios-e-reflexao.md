# Aula 22: Exercícios e Reflexão - Building CLIs

Olá! Agora é hora de praticar! Vamos criar alguns CLIs para consolidar o conhecimento.

---

## Exercício 1: CLI Simples com `flag`

Crie um CLI chamado `calculadora` que aceita dois números e uma operação.

**Requisitos:**
- Flags: `--a` (primeiro número), `--b` (segundo número), `--op` (operação: add, sub, mul, div)
- Validação: garantir que a operação seja válida
- Tratamento de erro: divisão por zero
- Mensagem de ajuda customizada

**Exemplo de uso:**
```bash
go run calculadora.go --a 10 --b 5 --op add
# Saída: 10 + 5 = 15
```

**Dica:**
```go
package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// Seu código aqui
}
```

---

## Exercício 2: CLI com Subcomandos (flag)

Crie um CLI chamado `arquivo` que tem dois subcomandos:
- `criar <nome>`: Cria um arquivo vazio
- `ler <nome>`: Lê e exibe o conteúdo de um arquivo

**Requisitos:**
- Usar `os.Args` para processar subcomandos
- Validação de argumentos
- Mensagens de erro claras
- Verificar se arquivo existe antes de ler

**Exemplo de uso:**
```bash
go run arquivo.go criar teste.txt
go run arquivo.go ler teste.txt
```

---

## Exercício 3: Todo CLI com Cobra

Crie um CLI de gerenciamento de tarefas usando Cobra.

**Requisitos:**
- Comando `add <tarefa>`: Adiciona uma nova tarefa
- Comando `list`: Lista todas as tarefas
- Comando `complete <id>`: Marca uma tarefa como completa
- Comando `delete <id>`: Remove uma tarefa
- Persistência: Salvar tarefas em um arquivo JSON

**Estrutura sugerida:**
```go
type Tarefa struct {
	ID        int    `json:"id"`
	Descricao string `json:"descricao"`
	Completa  bool   `json:"completa"`
}
```

**Exemplo de uso:**
```bash
go run todo.go add "Estudar Go"
go run todo.go list
go run todo.go complete 1
go run todo.go delete 1
```

**Dica:** Use `encoding/json` para salvar/carregar tarefas.

---

## Exercício 4: CLI de Conversão com urfave/cli

Crie um CLI chamado `converter` usando urfave/cli que converte entre diferentes unidades.

**Requisitos:**
- Comando `temp`: Converte temperatura (Celsius ↔ Fahrenheit ↔ Kelvin)
- Comando `dist`: Converte distância (km ↔ miles ↔ meters)
- Flags: `--from` e `--to` para especificar unidades
- Validação de unidades válidas

**Exemplo de uso:**
```bash
go run converter.go temp --from celsius --to fahrenheit --value 25
go run converter.go dist --from km --to miles --value 10
```

---

## Exercício 5: CLI Interativo com Bubble Tea

Crie um CLI interativo de seleção de cores usando Bubble Tea.

**Requisitos:**
- Lista de cores para selecionar
- Navegação com setas (↑↓)
- Seleção múltipla com espaço
- Exibir cores selecionadas ao final
- Sair com 'q' ou Ctrl+C

**Cores sugeridas:**
```go
cores := []string{
	"Vermelho", "Azul", "Verde", "Amarelo",
	"Roxo", "Laranja", "Rosa", "Preto",
}
```

**Dica:** Use o exemplo da aula como base e adicione seleção múltipla.

---

## Exercício 6: CLI Completo - Gerenciador de Contatos

Crie um CLI completo de gerenciamento de contatos usando Cobra.

**Requisitos:**
- `add <nome> <email> <telefone>`: Adiciona contato
- `list`: Lista todos os contatos
- `search <termo>`: Busca contatos por nome, email ou telefone
- `update <id> <campo> <valor>`: Atualiza um campo específico
- `delete <id>`: Remove um contato
- `export <formato>`: Exporta contatos (JSON ou CSV)
- Persistência em arquivo JSON
- Validação de email e telefone

**Estrutura sugerida:**
```go
type Contato struct {
	ID       int    `json:"id"`
	Nome     string `json:"nome"`
	Email    string `json:"email"`
	Telefone string `json:"telefone"`
}
```

**Exemplo de uso:**
```bash
go run contatos.go add "João Silva" "joao@email.com" "11999999999"
go run contatos.go list
go run contatos.go search "João"
go run contatos.go export json
```

---

## Exercício 7: CLI de Monitoramento com Bubble Tea

Crie um dashboard interativo que simula monitoramento de sistema.

**Requisitos:**
- Exibir CPU, Memória e Disco (valores simulados)
- Atualização em tempo real (a cada segundo)
- Indicadores visuais (barras de progresso)
- Cores diferentes baseadas nos valores (verde/amarelo/vermelho)
- Atualizar com 'r' (refresh)
- Sair com 'q'

**Dica:** Use `time.Tick` para atualizações periódicas e componentes do Bubble Tea.

---

## Exercício 8: CLI Cross-Platform

Crie um CLI simples e compile para múltiplas plataformas.

**Requisitos:**
- CLI que exibe informações do sistema (OS, arquitetura)
- Script de build para Linux, Windows e macOS
- Testar em pelo menos 2 plataformas diferentes

**Script de build sugerido:**
```bash
#!/bin/bash
# build.sh

APP="meucli"
VERSION="1.0.0"

# Linux
GOOS=linux GOARCH=amd64 go build -o "dist/${APP}-${VERSION}-linux-amd64"

# Windows
GOOS=windows GOARCH=amd64 go build -o "dist/${APP}-${VERSION}-windows-amd64.exe"

# macOS
GOOS=darwin GOARCH=arm64 go build -o "dist/${APP}-${VERSION}-darwin-arm64"
```

---

## Exercício 9: CLI com Variáveis de Ambiente

Crie um CLI que lê configurações de variáveis de ambiente usando urfave/cli.

**Requisitos:**
- Flag `--porta` que também pode vir de `PORT`
- Flag `--host` que também pode vir de `HOST`
- Flag `--debug` que também pode vir de `DEBUG`
- Exibir valores finais usados (mostrar origem: flag ou env)

**Exemplo de uso:**
```bash
export PORT=3000
go run app.go --host localhost
# Deve usar PORT=3000 (do env) e host=localhost (da flag)
```

---

## Exercício 10: CLI com Testes

Crie um CLI simples e escreva testes para ele.

**Requisitos:**
- CLI que processa uma lista de números e retorna estatísticas (soma, média, máximo, mínimo)
- Testes unitários para a lógica de cálculo
- Testes de integração que executam o CLI completo
- Cobertura de pelo menos 80%

**Dica:**
```go
// main.go
func calcularEstatisticas(numeros []int) (soma, media, max, min int) {
	// Implementação
}

// main_test.go
func TestCalcularEstatisticas(t *testing.T) {
	// Testes aqui
}
```

---

## Reflexão e Perguntas

Após completar os exercícios, reflita sobre:

1. **Qual ferramenta você achou mais fácil de usar?** Por quê?
2. **Quando você escolheria `flag` sobre Cobra?** E vice-versa?
3. **Quais são as vantagens de usar urfave/cli?**
4. **Em que situações Bubble Tea seria útil?**
5. **Como você organizaria um CLI grande com muitos subcomandos?**
6. **Quais são os desafios de fazer cross-compilation?**
7. **Como você testaria um CLI complexo?**
8. **Qual é a importância de mensagens de erro claras em CLIs?**

---

## Desafios Extras (Opcional)

### Desafio 1: CLI de Git Simplificado
Crie um CLI que simula comandos básicos do Git (init, add, commit, status) usando Cobra.

### Desafio 2: CLI de Chat Terminal
Crie um chat interativo no terminal usando Bubble Tea com múltiplas "salas" e histórico de mensagens.

### Desafio 3: CLI de Build System
Crie um sistema de build que lê um arquivo de configuração (YAML/JSON) e executa tarefas definidas.

### Desafio 4: CLI com Plugins
Crie um sistema de CLI que permite carregar plugins externos (usando Go plugins ou subprocessos).

---

## Dicas Finais

1. **Comece simples**: Use `flag` primeiro, evolua conforme necessário
2. **Teste frequentemente**: Execute seu CLI enquanto desenvolve
3. **Valide entrada**: Sempre valide argumentos e flags
4. **Mensagens claras**: Erros devem ser informativos
5. **Documentação**: Adicione help útil e exemplos
6. **Cross-compile**: Teste em diferentes plataformas
7. **Versionamento**: Considere adicionar versão ao seu CLI (`--version`)

---

Boa sorte com os exercícios! Lembre-se: a prática é essencial para dominar o desenvolvimento de CLIs em Go.

Se tiver dúvidas, revise a aula principal ou a versão simplificada. E não se esqueça de experimentar e explorar!

