# Atividade 04: Message Queue Simples

## Objetivo
Implementar message queue com múltiplos produtores e consumidores, garantias de entrega.

## Enunciado
Crie um message queue que:
1. Defina struct `MessageQueue` com topics e subscribers
2. Implemente `Publish(topic string, message Message) error`
3. Implemente `Subscribe(topic string, handler MessageHandler) (Subscription, error)`
4. Suporte múltiplos subscribers por topic (pub/sub)
5. Implemente garantias:
   - At-least-once delivery
   - Ordering (FIFO por topic)
   - Durability (persistência opcional)
6. Adicione dead letter queue para mensagens que falham
7. Implemente acknowledgment (ack) para confirmação de processamento
8. Métricas: mensagens publicadas, consumidas, em fila, rejeitadas

## Exemplo de Saída
```
Message Queue iniciada
Topic 'pedidos' criado
3 subscribers registrados
Publicando 1000 mensagens...
Mensagens publicadas: 1000
Mensagens consumidas: 1000
Mensagens em fila: 0
Dead letter queue: 2 (falhas)
Throughput: 500 msg/segundo
```

## Dicas
- Channels para cada topic
- Map de subscribers por topic
- Goroutines para consumidores
- Context para cancelamento
- Persistência em arquivo (opcional)

## Desafio Extra
Implemente partições, replicação e consumer groups.


