# Atividade 37: Máquina de Estados

## Objetivo
Implementar máquina de estados finita.

## Enunciado
Crie um programa que:
1. Defina estados: Inicial, Processando, Concluido, Erro
2. Defina struct `StateMachine` com estado atual e transições permitidas
3. Defina map de transições: `map[Estado]map[Evento]Estado`
4. Implemente métodos:
   - `Transicionar(evento Evento) error` - muda estado se transição válida
   - `EstadoAtual() Estado`
   - `PodeTransicionar(evento Evento) bool`
5. Implemente callbacks para cada transição
6. Simule processo de pedido: criado -> processando -> concluido

## Exemplo de Saída
```
Estado inicial: Inicial
Evento 'iniciar': Inicial -> Processando
Estado atual: Processando
Evento 'concluir': Processando -> Concluido
Estado final: Concluido
Tentativa de transição inválida: erro
```

## Dicas
- Map de transições: `map[Estado]map[Evento]Estado`
- Verifique se transição é permitida antes de mudar
- Callbacks permitem ações durante transições
- Útil para fluxos de trabalho complexos



