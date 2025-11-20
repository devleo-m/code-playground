# Atividade 07: Event Sourcing Básico

## Objetivo
Implementar sistema de Event Sourcing para armazenar estado como sequência de eventos.

## Enunciado
Crie um sistema de Event Sourcing:
1. Defina interface `Event` com métodos: `Type() string`, `Timestamp() time.Time`
2. Defina interface `Aggregate` com métodos:
   - `Apply(event Event) error`
   - `GetEvents() []Event`
   - `GetVersion() int`
3. Implemente `EventStore` para persistir eventos
4. Implemente `EventBus` para publicar eventos
5. Crie aggregate exemplo: `ContaBancaria` com eventos: `ContaCriada`, `DepositoRealizado`, `SaqueRealizado`
6. Implemente snapshot (salva estado periódico para performance)
7. Implemente replay (reconstrói estado a partir de eventos)
8. Adicione projeções (read models derivados dos eventos)

## Exemplo de Uso
```go
conta := NewContaBancaria("123")
conta.Depositar(100)
conta.Sacar(30)
eventStore.Save(conta.GetEvents())
// Replay
conta2 := ReplayEvents(eventStore.GetEvents("123"))
```

## Dicas
- Eventos são imutáveis
- Estado é derivado dos eventos
- Version para otimistic locking
- Snapshot reduz número de eventos a reaplicar
- Projeções para queries eficientes

## Desafio Extra
Implemente CQRS (Command Query Responsibility Segregation) completo.



