# Projeto 15: API REST Simples

## 📝 Descrição
Crie uma API REST completa usando net/http para gerenciar recursos com CRUD completo.

## 🎯 Requisitos

### Funcionalidades Obrigatórias
1. **Endpoints REST**:
   - GET /recursos - listar todos
   - GET /recursos/:id - obter por ID
   - POST /recursos - criar
   - PUT /recursos/:id - atualizar
   - DELETE /recursos/:id - deletar

2. **Recurso**: Escolha um (produtos, usuários, tarefas, etc.)

3. **JSON**: Request e response em JSON

4. **Validação**: Validar dados de entrada

5. **Error Handling**: Status codes apropriados (200, 201, 404, 400, 500)

6. **Persistência**: JSON file ou in-memory

7. **Middleware**: Logging de requisições

## 📚 Conceitos Utilizados
- ✅ net/http package
- ✅ JSON encoding/decoding
- ✅ Error handling
- ✅ Structs
- ✅ Interfaces
- ✅ Middleware pattern
- ✅ REST principles

## 📁 Estrutura Sugerida
```
api/
├── main.go
├── handlers.go
├── models.go
├── repository.go
├── middleware.go
└── README.md
```

## 💡 Implementação Sugerida

### Handlers
```go
func ListarRecursos(w http.ResponseWriter, r *http.Request)
func ObterRecurso(w http.ResponseWriter, r *http.Request)
func CriarRecurso(w http.ResponseWriter, r *http.Request)
func AtualizarRecurso(w http.ResponseWriter, r *http.Request)
func DeletarRecurso(w http.ResponseWriter, r *http.Request)
```

### Middleware
- Logging middleware
- Error handling middleware
- CORS middleware (opcional)

## ✅ Critérios de Sucesso
- [ ] Todos endpoints funcionam
- [ ] JSON é válido
- [ ] Status codes corretos
- [ ] Validações funcionam
- [ ] Middleware funciona
- [ ] Código organizado

## 🚀 Extras (Desafio)
- [ ] Paginação
- [ ] Filtros e busca
- [ ] Autenticação básica
- [ ] Rate limiting
- [ ] Documentação (Swagger)
- [ ] Testes HTTP


