# Aula 24: Exercícios e Reflexão - Frameworks Web - Gin

Olá! Agora é hora de praticar! Vamos criar algumas APIs usando Gin para consolidar o conhecimento.

---

## Exercício 1: API de Produtos

Crie uma API REST de produtos usando Gin.

**Requisitos:**
- `GET /products` - Lista todos os produtos
- `POST /products` - Cria um novo produto
- `GET /products/:id` - Busca produto por ID
- `PUT /products/:id` - Atualiza produto
- `DELETE /products/:id` - Remove produto
- Validação: nome (obrigatório, min 3), preço (obrigatório, > 0)

**Estrutura:**
```go
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name" binding:"required,min=3"`
	Description string  `json:"description"`
	Price       float64 `json:"price" binding:"required,gt=0"`
	Stock       int     `json:"stock" binding:"gte=0"`
}
```

**Exemplo de uso:**
```bash
curl -X POST http://localhost:8080/products \
  -H "Content-Type: application/json" \
  -d '{"name":"Notebook","price":2999.99,"stock":10}'
```

---

## Exercício 2: Sistema de Autenticação

Crie um sistema de autenticação básico com Gin.

**Requisitos:**
- `POST /register` - Registra novo usuário
- `POST /login` - Faz login (retorna token JWT simples)
- `GET /profile` - Retorna perfil do usuário (protegido)
- Middleware de autenticação
- Validação de email e senha

**Estrutura:**
```go
type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Name     string `json:"name" binding:"required"`
}
```

**Dica:** Use um token simples (string) armazenado em memória para este exercício.

---

## Exercício 3: API com Upload de Imagens

Crie uma API que permite upload de imagens.

**Requisitos:**
- `POST /upload` - Upload de imagem
- `GET /images` - Lista todas as imagens
- `GET /images/:id` - Retorna imagem específica
- Validação: apenas JPG, PNG, GIF (máx 5MB)
- Salvar em `./uploads/`
- Retornar URL da imagem

**Exemplo:**
```bash
curl -X POST http://localhost:8080/upload \
  -F "file=@imagem.jpg"
```

---

## Exercício 4: API com Paginação e Filtros

Crie uma API com paginação e filtros avançados.

**Requisitos:**
- `GET /users?page=1&limit=10&sort=name&order=asc`
- Paginação: `page` e `limit`
- Ordenação: `sort` (campo) e `order` (asc/desc)
- Filtros: `name`, `email`, `age`
- Retornar metadata: total, página atual, total de páginas

**Resposta esperada:**
```json
{
  "data": [...],
  "meta": {
    "total": 100,
    "page": 1,
    "limit": 10,
    "total_pages": 10
  }
}
```

---

## Exercício 5: Middleware Customizado

Crie middlewares customizados para diferentes funcionalidades.

**Requisitos:**
- Middleware de rate limiting (10 req/min por IP)
- Middleware de compressão (gzip)
- Middleware de cache (cachear respostas GET por 5 min)
- Middleware de validação de API key
- Aplicar em diferentes grupos de rotas

**Estrutura:**
```go
// Rotas públicas (sem rate limit)
GET /public

// Rotas com API key
GET /api/data (requer header X-API-Key)

// Rotas com cache
GET /api/cacheable
```

---

## Exercício 6: API com Relacionamentos

Crie uma API que gerencia relacionamentos entre entidades.

**Requisitos:**
- `GET /users` - Lista usuários
- `POST /users/:id/posts` - Cria post para usuário
- `GET /users/:id/posts` - Lista posts do usuário
- `GET /posts/:id/comments` - Lista comentários do post
- `POST /posts/:id/comments` - Adiciona comentário

**Estrutura:**
```go
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Post struct {
	ID     int    `json:"id"`
	UserID int    `json:"user_id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

type Comment struct {
	ID     int    `json:"id"`
	PostID int    `json:"post_id"`
	Text   string `json:"text"`
}
```

---

## Exercício 7: Validação Customizada

Crie validações customizadas para casos específicos.

**Requisitos:**
- Validação de CPF (formato brasileiro)
- Validação de CEP (formato brasileiro)
- Validação de telefone (formato brasileiro)
- Validação de data (formato DD/MM/YYYY)
- Usar tags de binding customizadas

**Exemplo:**
```go
type User struct {
	CPF     string `json:"cpf" binding:"required,cpf"`
	CEP     string `json:"cep" binding:"required,cep"`
	Phone   string `json:"phone" binding:"required,phone"`
	BirthDate string `json:"birth_date" binding:"required,date_format"`
}
```

**Dica:** Use `validator.RegisterValidation()` do Gin.

---

## Exercício 8: API com Versionamento

Crie uma API com versionamento de endpoints.

**Requisitos:**
- `/api/v1/users` - Versão 1 (estrutura antiga)
- `/api/v2/users` - Versão 2 (estrutura nova)
- Manter compatibilidade entre versões
- Middleware que detecta versão e aplica transformações

**Estrutura v1:**
```go
type UserV1 struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
```

**Estrutura v2:**
```go
type UserV2 struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
}
```

---

## Exercício 9: API com WebSockets (Opcional)

Crie uma API que suporta WebSockets usando Gin.

**Requisitos:**
- Endpoint WebSocket `/ws`
- Broadcast de mensagens para todos os clientes conectados
- Sistema de chat simples
- Gerenciar conexões ativas

**Dica:** Use `github.com/gorilla/websocket` ou similar.

---

## Exercício 10: API Completa de E-commerce

Crie uma API completa de e-commerce com Gin.

**Requisitos:**
- **Produtos**: CRUD completo
- **Carrinho**: Adicionar/remover produtos, calcular total
- **Pedidos**: Criar pedido, listar pedidos do usuário
- **Categorias**: CRUD de categorias
- **Autenticação**: Login/registro
- **Validação**: Completa em todos os endpoints
- **Middleware**: Logging, CORS, Auth
- **Documentação**: Endpoint `/docs` com Swagger (opcional)

**Estrutura sugerida:**
```
/api/v1
├── /auth
│   ├── POST /register
│   └── POST /login
├── /products
│   ├── GET /
│   ├── POST /
│   ├── GET /:id
│   ├── PUT /:id
│   └── DELETE /:id
├── /categories
│   └── (CRUD completo)
├── /cart
│   ├── GET /
│   ├── POST /add
│   └── DELETE /remove/:id
└── /orders
    ├── POST /
    └── GET /
```

---

## Reflexão

Após completar os exercícios, reflita sobre:

1. **Gin vs net/http**
   - Quais são as principais vantagens do Gin?
   - Quando você escolheria Gin sobre net/http?

2. **Validação**
   - Como a validação automática do Gin ajuda no desenvolvimento?
   - Quais são as limitações das validações built-in?

3. **Middleware**
   - Como o sistema de middleware do Gin facilita o desenvolvimento?
   - Quais middlewares são essenciais em produção?

4. **Organização de Código**
   - Como você organizaria um projeto grande com Gin?
   - Quais padrões você seguiria?

5. **Performance**
   - Gin mantém a performance do Go?
   - Quais otimizações você pode fazer?

---

## Desafio Final

Crie uma API completa de rede social com:
- Usuários (perfis, seguir/deixar de seguir)
- Posts (criar, editar, deletar, curtir)
- Comentários (criar, editar, deletar)
- Feed (timeline personalizada)
- Busca (usuários e posts)
- Autenticação JWT completa
- Upload de imagens
- Notificações em tempo real (WebSocket)
- Rate limiting e cache
- Documentação completa

**Funcionalidades avançadas:**
- Paginação infinita
- Filtros e ordenação
- Relacionamentos complexos
- Validação customizada
- Middleware avançado

---

Boa sorte com os exercícios! Pratique bastante e explore todas as funcionalidades do Gin!

