# Atividade 24: Reactive Streams

## Objetivo
Implementar padrão Reactive Streams com backpressure.

## Enunciado
Crie sistema Reactive Streams:
1. Defina interfaces:
   - `Publisher` - publica itens
   - `Subscriber` - consome itens
   - `Subscription` - controla fluxo
2. Implemente `Publisher` que emite itens
3. Implemente `Subscriber` que processa itens
4. Implemente backpressure:
   - Subscriber solicita N itens
   - Publisher respeita limite
5. Suporte a operadores:
   - `Map`, `Filter`, `FlatMap`
   - `Take`, `Skip`, `Distinct`
6. Implemente múltiplos subscribers (multicast)
7. Teste com publisher rápido e subscriber lento

## Exemplo de Saída
```
Publisher: emitindo 10000 itens
Subscriber: solicitando 100 itens por vez
Backpressure ativo: publisher aguardando...
Itens processados: 10000
Throughput: 5000 itens/segundo
Solicitações: 100
```

## Dicas
- Reactive: programação assíncrona com streams
- Backpressure: subscriber controla velocidade
- Subscription: gerencia demanda
- Operadores: transformam streams
- Útil para I/O assíncrono

## Desafio Extra
Implemente hot vs cold publishers e error handling.


