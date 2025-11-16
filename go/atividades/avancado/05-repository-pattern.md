# Atividade 05: Repository Pattern com Interfaces

## Objetivo
Implementar Repository Pattern completo com múltiplas implementações e testes.

## Enunciado
Crie um sistema usando Repository Pattern:
1. Defina interface `Repository[T Entity]` com métodos:
   - `Create(entity T) (T, error)`
   - `GetByID(id string) (T, error)`
   - `Update(id string, entity T) error`
   - `Delete(id string) error`
   - `FindAll(filters map[string]interface{}) ([]T, error)`
2. Implemente `InMemoryRepository` (map em memória)
3. Implemente `FileRepository` (persistência em JSON)
4. Crie `Service` layer que usa repository via interface
5. Implemente unit tests com mocks
6. Adicione validação de negócio no service
7. Implemente transações (simuladas para in-memory)

## Exemplo de Uso
```go
repo := NewInMemoryRepository[Usuario]()
service := NewUsuarioService(repo)
usuario, _ := service.CriarUsuario(Usuario{Nome: "João"})
usuarios, _ := service.BuscarTodos()
```

## Dicas
- Interface genérica permite diferentes implementações
- Service contém lógica de negócio
- Repository apenas acesso a dados
- Use generics (Go 1.18+) ou interface{} com type assertion
- Mocks para testes isolados

## Desafio Extra
Implemente cache layer entre service e repository, e suporte a múltiplos repositórios (read/write separation).

