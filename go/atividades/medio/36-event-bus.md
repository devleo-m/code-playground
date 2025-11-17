# Atividade 36: Event Bus Simples

## Objetivo
Implementar sistema de eventos (pub/sub).

## Enunciado
Crie um programa que:
1. Defina tipo `EventHandler func(evento string, dados interface{})`
2. Defina struct `EventBus` com map de handlers por tipo de evento
3. Implemente métodos:
   - `Subscribe(tipoEvento string, handler EventHandler)`
   - `Unsubscribe(tipoEvento string, handler EventHandler)`
   - `Publish(tipoEvento string, dados interface{})`
4. Permita múltiplos handlers para mesmo evento
5. Implemente handler assíncrono (usa goroutine)
6. Simule sistema de eventos: pedido_criado, pagamento_aprovado, etc.

## Exemplo de Saída
```
Subscrevendo handlers para 'pedido_criado'...
Publicando evento 'pedido_criado':
  Handler Email: processando pedido_criado
  Handler SMS: processando pedido_criado
  Handler Log: registrando pedido_criado
```

## Dicas
- Map: `map[string][]EventHandler`
- Append handlers ao slice
- Itere sobre handlers ao publicar
- Goroutine para execução assíncrona (opcional)


