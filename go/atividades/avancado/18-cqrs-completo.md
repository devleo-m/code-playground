# Atividade 18: CQRS Completo

## Objetivo
Implementar CQRS (Command Query Responsibility Segregation) completo com event sourcing.

## Enunciado
Crie sistema CQRS:
1. Separe Command (write) e Query (read) models
2. Command side:
   - `CommandBus` para rotear commands
   - `CommandHandler` interface
   - Validação de commands
   - Gera eventos após processamento
3. Query side:
   - `QueryBus` para rotear queries
   - `QueryHandler` interface
   - Read models otimizados para queries
4. Event handlers que atualizam read models
5. Implemente projeções (read models derivados de eventos)
6. Suporte a múltiplos read models para mesma query
7. Implemente eventual consistency
8. Teste com 1000 commands e verifique read models

## Exemplo de Uso
```go
// Write
command := CriarUsuarioCommand{Nome: "João"}
commandBus.Send(command)
// Event: UsuarioCriado gerado
// Read model atualizado assincronamente

// Read
query := BuscarUsuarioQuery{ID: "123"}
usuario := queryBus.Send(query)
```

## Dicas
- Commands mudam estado, Queries leem
- Event handlers atualizam read models
- Eventual consistency: read pode estar desatualizado
- Projeções: diferentes views dos mesmos dados
- Útil para alta performance e escalabilidade

## Desafio Extra
Implemente materialized views e sincronização de read models.

