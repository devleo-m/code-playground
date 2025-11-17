# Atividade 03: Cache LRU (Least Recently Used)

## Objetivo
Implementar cache LRU thread-safe com TTL usando estruturas de dados eficientes.

## Enunciado
Crie um cache LRU que:
1. Defina struct `LRUCache` com capacidade máxima e TTL opcional
2. Use hash map + doubly linked list para O(1) get e put
3. Implemente métodos thread-safe:
   - `Get(key string) (interface{}, bool)` - move para frente se encontrado
   - `Put(key string, value interface{})` - adiciona ou atualiza
   - `Delete(key string) bool`
   - `Clear()`
   - `Size() int`
4. Implemente expiração automática baseada em TTL
5. Adicione métricas: hit rate, miss rate, evictions
6. Use `sync.RWMutex` para thread-safety
7. Teste com 10.000 operações concorrentes

## Exemplo de Saída
```
Cache LRU criado: capacidade 100, TTL 5min
Operações: 10000
Hit rate: 85.3%
Miss rate: 14.7%
Evictions: 150
Tempo médio Get: 0.05ms
Tempo médio Put: 0.08ms
```

## Dicas
- Hash map para lookup O(1)
- Doubly linked list para ordem LRU
- Move para frente ao acessar
- Remove do final quando cheio
- Goroutine para limpeza de expirados

## Desafio Extra
Implemente variantes: LFU (Least Frequently Used) e ARC (Adaptive Replacement Cache).


