# Aula 20 - Exercícios e Reflexão: Standard Library

Olá! Agora é hora de colocar a mão na massa e praticar tudo que aprendemos sobre a Standard Library do Go. Vamos começar com exercícios práticos e depois refletir sobre os conceitos.

---

## 📝 Exercícios Práticos

### Exercício 1: Criador de Log de Atividades

Crie um programa que:
1. Aceite um nome de arquivo via flag `-arquivo` (padrão: "atividades.txt")
2. Aceite uma mensagem via flag `-mensagem` (obrigatório)
3. Registre a mensagem no arquivo com timestamp no formato: `[YYYY-MM-DD HH:MM:SS] mensagem`
4. Use `slog` para registrar no console também (nível Info)

**Exemplo de uso:**
```bash
go run main.go -arquivo="minhas_atividades.txt" -mensagem="Estudei Go hoje"
```

**Saída esperada no arquivo:**
```
[2024-01-15 14:30:00] Estudei Go hoje
[2024-01-15 14:35:00] Fiz exercícios de Standard Library
```

**Dica:** Use `time.Now().Format()` para o timestamp.

---

### Exercício 2: Validador de Dados Pessoais

Crie um programa que:
1. Leia um arquivo JSON com dados de pessoas (formato abaixo)
2. Valide cada pessoa usando regex:
   - Email deve ter formato válido
   - Telefone deve estar no formato `(XX) XXXXX-XXXX`
3. Crie um novo arquivo JSON apenas com pessoas válidas
4. Use `bufio.Scanner` para ler o arquivo linha por linha (caso tenha múltiplos objetos JSON)

**Formato do arquivo de entrada (`pessoas.json`):**
```json
{"nome":"João Silva","email":"joao@example.com","telefone":"(11) 98765-4321"}
{"nome":"Maria Santos","email":"maria.email.com","telefone":"(21) 91234-5678"}
{"nome":"Pedro Costa","email":"pedro@teste","telefone":"11987654321"}
```

**Dica:** Use `json.Decoder` para ler múltiplos objetos JSON de um stream.

---

### Exercício 3: Buscador de Arquivos

Crie um programa que:
1. Aceite um diretório via flag `-dir` (padrão: diretório atual)
2. Aceite um padrão regex via flag `-padrao` (obrigatório)
3. Procure em todos os arquivos do diretório (e subdiretórios) por linhas que correspondam ao padrão
4. Exiba: nome do arquivo, número da linha e conteúdo da linha
5. Use `slog` com nível Info para cada correspondência encontrada

**Exemplo de uso:**
```bash
go run main.go -dir="./meu_projeto" -padrao="func.*main"
```

**Saída esperada:**
```json
{"level":"info","msg":"Correspondência encontrada","arquivo":"main.go","linha":5,"conteudo":"func main() {"}
```

**Dica:** Use `filepath.Walk()` ou `os.WalkDir()` para percorrer diretórios recursivamente.

---

### Exercício 4: Servidor de Arquivos Estáticos com Embed

Crie um programa que:
1. Use `go:embed` para incluir uma pasta `static` com arquivos HTML/CSS
2. Crie um servidor HTTP simples que sirva esses arquivos
3. Quando acessar `/`, sirva `static/index.html`
4. Quando acessar `/estilo.css`, sirva `static/estilo.css`
5. Use `flag` para aceitar a porta via `-porta` (padrão: 8080)

**Estrutura de arquivos:**
```
static/
  index.html
  estilo.css
main.go
```

**Exemplo de uso:**
```bash
go run main.go -porta=3000
# Acesse http://localhost:3000
```

**Dica:** Use `embed.FS` e `http.FileServer()` ou leia os arquivos manualmente com `embed.FS.ReadFile()`.

---

## 🤔 Perguntas de Reflexão

### Reflexão 1: Por que Interfaces são Fundamentais no I/O do Go?

Na aula, vimos que Go usa interfaces (`io.Reader`, `io.Writer`) para I/O. Pense sobre:

1. **Flexibilidade**: Como as interfaces permitem que você escreva uma função que funcione tanto com arquivos quanto com conexões de rede?

2. **Testabilidade**: Como você poderia testar uma função que recebe um `io.Reader` sem precisar criar um arquivo real?

3. **Reutilização**: Se você criar uma função que processa dados de um `io.Reader`, quantos tipos diferentes de fontes de dados ela poderá processar? Liste pelo menos 5.

**Sua resposta deve incluir:**
- Uma explicação sobre o poder das interfaces no design do Go
- Um exemplo prático de como você usaria isso em um projeto real
- Uma comparação: como seria fazer isso sem interfaces (em outra linguagem que você conheça, se aplicável)

---

### Reflexão 2: Standard Library vs Bibliotecas Externas

A Standard Library do Go é muito completa, mas às vezes você pode precisar de bibliotecas externas. Pense sobre:

1. **Quando usar a Standard Library?**
   - Quais são as vantagens?
   - Em que situações ela é suficiente?

2. **Quando buscar bibliotecas externas?**
   - Quais são os riscos?
   - Como você decide se vale a pena adicionar uma dependência?

3. **Trade-offs**: 
   - Um projeto usando apenas Standard Library vs um projeto com muitas dependências externas
   - Quais são os prós e contras de cada abordagem?

**Sua resposta deve incluir:**
- Um cenário real onde você escolheria usar apenas Standard Library
- Um cenário real onde você escolheria uma biblioteca externa
- Sua opinião sobre a filosofia do Go de "bateria inclusa" (batteries included)

---

## 📋 Checklist de Aprendizado

Antes de enviar suas respostas, verifique se você:

- [ ] Consegue ler e escrever arquivos usando `os` e `io`
- [ ] Sabe usar `flag` para criar CLIs profissionais
- [ ] Compreende a diferença entre `time.Time` e `time.Duration`
- [ ] Consegue fazer Marshal/Unmarshal de JSON com struct tags
- [ ] Sabe usar `bufio.Scanner` para leitura eficiente
- [ ] Consegue configurar `slog` com diferentes níveis e formatos
- [ ] Sabe criar e usar expressões regulares com `regexp`
- [ ] Compreende como `go:embed` funciona e quando usá-lo
- [ ] Entendeu o conceito de interfaces no I/O do Go
- [ ] Refletiu sobre quando usar Standard Library vs bibliotecas externas

---

## 🎯 Dicas para os Exercícios

- **Exercício 1**: Lembre-se de usar `defer` para fechar arquivos e `os.O_APPEND` para adicionar ao final do arquivo
- **Exercício 2**: Teste seus regex antes de usar no código. Use sites como regex101.com para validar
- **Exercício 3**: Cuidado com arquivos binários! Você pode querer filtrar apenas arquivos de texto
- **Exercício 4**: O `embed.FS` implementa `fs.FS`, então você pode usar `http.FS()` para servir os arquivos

---

Boa sorte com os exercícios! Lembre-se: a prática é essencial para dominar a Standard Library. Quando terminar, envie suas soluções e respostas de reflexão para análise.

