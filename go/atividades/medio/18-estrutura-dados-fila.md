# Atividade 18: Estrutura de Dados - Fila

## Objetivo
Implementar estrutura de dados fila (queue) e fila circular.

## Enunciado
Crie um programa que:
1. Defina struct `Fila` com slice e índices de início/fim
2. Implemente métodos:
   - `Enqueue(valor interface{})` - adiciona no fim
   - `Dequeue() (interface{}, error)` - remove do início
   - `Front() (interface{}, error)` - retorna início sem remover
   - `IsEmpty() bool`
   - `Size() int`
3. Implemente `FilaCircular` com tamanho fixo (usa array)
4. Use fila para implementar BFS (busca em largura) simples em grafo
5. Simule sistema de atendimento com fila de clientes

## Exemplo de Saída
```
Fila:
  Enqueue(1), Enqueue(2), Enqueue(3)
  Front: 1
  Dequeue: 1
  Size: 2

BFS no grafo:
  Visitando: A
  Visitando: B, C
  Visitando: D

Sistema de atendimento:
  Cliente 1 atendido
  Cliente 2 em espera...
  Cliente 3 em espera...
```

## Dicas
- Fila: FIFO (First In First Out)
- Use slice: `append` para enqueue, `slice[0]` e `slice[1:]` para dequeue
- Fila circular: use módulo para índices
- BFS: use fila para manter ordem de visitação


