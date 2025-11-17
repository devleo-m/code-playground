# Atividade 23: Actor Model

## Objetivo
Implementar modelo Actor para concorrência baseada em mensagens.

## Enunciado
Crie sistema Actor:
1. Defina interface `Actor` com método `Receber(mensagem interface{})`
2. Defina struct `ActorSystem` que gerencia actors
3. Cada actor tem mailbox (channel) e processa mensagens sequencialmente
4. Implemente `Spawn(handler ActorHandler) ActorRef`
5. Implemente `Send(actorRef ActorRef, mensagem interface{}) error`
6. Suporte a supervision (restart actors que falham)
7. Implemente actor hierarchy (parent-child)
8. Adicione métricas: mensagens processadas, actors ativos
9. Teste com 100 actors e 10.000 mensagens

## Exemplo de Uso
```go
system := NewActorSystem()
actor := system.Spawn(NewContadorActor())
actor.Send(Incrementar{})
actor.Send(Incrementar{})
result := actor.Send(ObterValor{})
```

## Dicas
- Actor: encapsula estado e comportamento
- Mailbox: fila de mensagens
- Processamento sequencial por actor
- Isolamento: actors não compartilham estado
- Supervision: gerencia falhas

## Desafio Extra
Implemente location transparency e remote actors.


