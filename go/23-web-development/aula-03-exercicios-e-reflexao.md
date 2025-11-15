# Aula 23: Exercícios e Reflexão - Web Development com net/http

Olá! Agora é hora de praticar! Vamos criar alguns servidores web para consolidar o conhecimento sobre `net/http`.

---

## Exercício 1: Servidor HTTP Básico

Crie um servidor HTTP simples que responde em diferentes rotas.

**Requisitos:**
- Rota `/` que retorna "Bem-vindo!"
- Rota `/about` que retorna "Sobre nós"
- Rota `/contact` que retorna informações de contato em JSON
- Todas as rotas devem ter logging (middleware)

**Exemplo de saída:**
```bash
# GET /
Bem-vindo!

# GET /about
Sobre nós

# GET /contact
{"email":"contato@example.com","phone":"123-456-7890"}
```

**Dica:**
```go
func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%s %s\n", r.Method, r.URL.Path)
		next(w, r)
	}
}
```

---

## Exercício 2: API de Calculadora

Crie uma API REST para uma calculadora simples.

**Requisitos:**
- `GET /calc?op=add&a=10&b=5` - Retorna resultado da operação
- Operações suportadas: `add`, `sub`, `mul`, `div`
- Validação: garantir que `a` e `b` sejam números válidos
- Tratamento de erro: divisão por zero
- Retornar JSON: `{"result": 15}`

**Exemplo de uso:**
```bash
curl "http://localhost:8080/calc?op=add&a=10&b=5"
# {"result":15}

curl "http://localhost:8080/calc?op=div&a=10&b=0"
# {"error":"Divisão por zero"}
```

**Dica:**
```go
func calcHandler(w http.ResponseWriter, r *http.Request) {
	op := r.URL.Query().Get("op")
	a, _ := strconv.Atoi(r.URL.Query().Get("a"))
	b, _ := strconv.Atoi(r.URL.Query().Get("b"))
	// Implementar lógica
}
```

---

## Exercício 3: API de Tarefas (Todo List)

Crie uma API REST completa para gerenciar tarefas usando apenas `net/http`.

**Requisitos:**
- `GET /tasks` - Lista todas as tarefas
- `POST /tasks` - Cria uma nova tarefa (JSON body)
- `GET /tasks/:id` - Busca uma tarefa por ID
- `PUT /tasks/:id` - Atualiza uma tarefa
- `DELETE /tasks/:id` - Remove uma tarefa
- Persistência em memória (slice)
- Validação de dados de entrada

**Estrutura de Tarefa:**
```go
type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}
```

**Exemplo de uso:**
```bash
# Criar tarefa
curl -X POST http://localhost:8080/tasks \
  -H "Content-Type: application/json" \
  -d '{"title":"Estudar Go","description":"Aprender net/http"}'

# Listar tarefas
curl http://localhost:8080/tasks

# Atualizar tarefa
curl -X PUT http://localhost:8080/tasks/1 \
  -H "Content-Type: application/json" \
  -d '{"title":"Estudar Go","description":"Aprender net/http","completed":true}'
```

---

## Exercício 4: Servidor de Arquivos Estáticos

Crie um servidor que serve arquivos estáticos e tem uma API para upload.

**Requisitos:**
- Servir arquivos estáticos de `./static/` na rota `/static/`
- `POST /upload` - Upload de arquivo (multipart form)
- `GET /files` - Lista todos os arquivos enviados
- Salvar arquivos em `./uploads/`
- Validação: máximo 10MB por arquivo

**Exemplo de uso:**
```bash
# Upload
curl -X POST http://localhost:8080/upload \
  -F "file=@documento.pdf"

# Listar arquivos
curl http://localhost:8080/files
```

**Dica:**
```go
func uploadHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20) // 10MB
	file, handler, err := r.FormFile("file")
	// Salvar arquivo
}
```

---

## Exercício 5: Cliente HTTP para API Externa

Crie um cliente que consome uma API externa e processa os dados.

**Requisitos:**
- Fazer requisição GET para `https://jsonplaceholder.typicode.com/posts`
- Processar a resposta JSON
- Criar endpoint `/posts` que retorna os posts formatados
- Adicionar cache em memória (TTL de 5 minutos)
- Tratamento de erros e timeouts

**Estrutura esperada:**
```go
type Post struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	UserID int    `json:"userId"`
}
```

**Dica:**
```go
func fetchPosts() ([]Post, error) {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	// Processar resposta
}
```

---

## Exercício 6: Middleware Chain Avançado

Crie um servidor com múltiplos middlewares encadeados.

**Requisitos:**
- Middleware de logging (registra todas as requisições)
- Middleware de CORS (permite requisições de qualquer origem)
- Middleware de autenticação (verifica header `Authorization`)
- Middleware de rate limiting (máximo 10 requisições por minuto por IP)
- Aplicar middlewares em diferentes grupos de rotas

**Estrutura:**
```go
// Rotas públicas (sem auth)
GET /public

// Rotas protegidas (com auth)
GET /api/users
POST /api/users
```

**Dica:**
```go
type Middleware func(http.HandlerFunc) http.HandlerFunc

func chainMiddleware(middlewares ...Middleware) Middleware {
	// Implementar encadeamento
}
```

---

## Exercício 7: API com Validação Completa

Crie uma API de usuários com validação robusta.

**Requisitos:**
- `POST /users` - Criar usuário com validação:
  - Nome: obrigatório, mínimo 3 caracteres
  - Email: obrigatório, formato válido
  - Idade: obrigatório, entre 18 e 120
- Retornar erros de validação detalhados
- Códigos de status HTTP apropriados

**Estrutura:**
```go
type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}
```

**Exemplo de resposta de erro:**
```json
{
  "errors": [
    {"field": "name", "message": "Nome deve ter no mínimo 3 caracteres"},
    {"field": "email", "message": "Email inválido"}
  ]
}
```

---

## Exercício 8: Servidor com Graceful Shutdown

Implemente um servidor que faz shutdown graceful.

**Requisitos:**
- Capturar sinais SIGINT e SIGTERM
- Aguardar requisições em andamento terminarem
- Timeout de 30 segundos para shutdown
- Log de shutdown

**Dica:**
```go
func main() {
	server := &http.Server{Addr: ":8080", Handler: mux}
	
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()
	
	// Capturar sinais
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	<-sigChan
	
	// Shutdown graceful
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	server.Shutdown(ctx)
}
```

---

## Reflexão

Após completar os exercícios, reflita sobre:

1. **Quando usar `net/http` vs frameworks?**
   - Quais são as vantagens de usar apenas a biblioteca padrão?
   - Em que situações você escolheria um framework?

2. **Middleware e Reutilização**
   - Como você organizaria middlewares em um projeto grande?
   - Quais middlewares são essenciais em produção?

3. **Tratamento de Erros**
   - Como você estruturaria mensagens de erro consistentes?
   - Qual a importância dos códigos de status HTTP?

4. **Performance**
   - Como `net/http` lida com concorrência?
   - Quais otimizações você pode fazer?

5. **Segurança**
   - Quais vulnerabilidades comuns em APIs HTTP?
   - Como proteger sua API?

---

## Desafio Final

Crie uma API REST completa de blog com:
- Posts (CRUD completo)
- Comentários (CRUD completo)
- Autenticação básica
- Validação completa
- Middleware de logging e CORS
- Tratamento de erros robusto
- Documentação básica (endpoint `/docs`)

**Estrutura sugerida:**
```
/blog
├── posts
│   ├── GET /posts
│   ├── POST /posts
│   ├── GET /posts/:id
│   ├── PUT /posts/:id
│   └── DELETE /posts/:id
└── comments
    ├── GET /posts/:id/comments
    ├── POST /posts/:id/comments
    └── DELETE /comments/:id
```

---

Boa sorte com os exercícios! Pratique bastante e experimente diferentes abordagens!

