# Atividade 25: Padrão Observer

## Objetivo
Implementar padrão Observer usando interfaces e slices.

## Enunciado
Crie um programa que:
1. Defina interface `Observer` com método `Atualizar(evento string)`
2. Defina struct `Subject` com slice de observers
3. Implemente métodos:
   - `AdicionarObserver(obs Observer)`
   - `RemoverObserver(obs Observer)`
   - `Notificar(evento string)`
4. Crie structs `ObserverEmail`, `ObserverSMS` que implementam `Observer`
5. Simule sistema de notificações: quando evento ocorre, todos observers são notificados

## Exemplo de Saída
```
Adicionando observers...
Evento: 'pedido_criado'
  Email: Enviando email sobre pedido_criado
  SMS: Enviando SMS sobre pedido_criado
Evento: 'pagamento_aprovado'
  Email: Enviando email sobre pagamento_aprovado
  SMS: Enviando SMS sobre pagamento_aprovado
```

## Dicas
- Interface permite diferentes tipos de observers
- Slice armazena lista de observers
- Itere sobre slice para notificar todos
- Padrão útil para desacoplamento

