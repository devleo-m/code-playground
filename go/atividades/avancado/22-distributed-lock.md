# Atividade 22: Distributed Lock

## Objetivo
Implementar distributed lock com expiração e renovação automática.

## Enunciado
Crie distributed lock:
1. Defina interface `DistributedLock` com métodos:
   - `Lock(timeout time.Duration) error`
   - `Unlock() error`
   - `Renovar() error`
   - `TentarLock(timeout time.Duration) (bool, error)`
2. Implemente com backend (simule com arquivo ou sync.Map compartilhado)
3. Adicione expiração automática (TTL)
4. Implemente renovação automática (refresh)
5. Suporte a lock reentrante (mesmo owner pode lock múltiplas vezes)
6. Implemente wait com timeout
7. Adicione métricas: locks adquiridos, tempo médio de espera
8. Teste com 10 goroutines concorrentes

## Exemplo de Saída
```
Distributed Lock: recurso 'database'
Goroutine 1: tentando lock... adquirido ✓
Goroutine 2: tentando lock... aguardando...
Goroutine 3: tentando lock... aguardando...
Goroutine 1: liberando lock...
Goroutine 2: lock adquirido ✓
Tempo médio de espera: 150ms
Locks adquiridos: 10
```

## Dicas
- Lock: apenas um owner por vez
- TTL: expira automaticamente
- Renovação: estende TTL enquanto em uso
- Wait: bloqueia até lock disponível
- Backend: Redis, etcd, ou arquivo compartilhado

## Desafio Extra
Implemente com Redis real e suporte a lock hierarchy.


