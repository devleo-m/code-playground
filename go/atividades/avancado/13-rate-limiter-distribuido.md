# Atividade 13: Rate Limiter Distribuído

## Objetivo
Implementar rate limiter que funciona em ambiente distribuído com múltiplas instâncias.

## Enunciado
Crie rate limiter distribuído:
1. Implemente algoritmo Token Bucket distribuído
2. Use backend compartilhado (simule com arquivo ou in-memory sync)
3. Suporte múltiplas estratégias:
   - Fixed window
   - Sliding window log
   - Sliding window counter
4. Implemente sincronização entre instâncias (use locks ou atomic operations)
5. Adicione métricas por cliente/IP
6. Implemente diferentes limites por tipo de cliente (VIP, normal, free)
7. Adicione burst allowance (permite picos temporários)
8. Teste com 10 instâncias concorrentes

## Exemplo de Saída
```
Rate Limiter Distribuído: 100 req/min por IP
Instâncias: 10
Cliente 192.168.1.1:
  Requisições: 95/100
  Window: 00:00-00:01
  Status: permitido
Cliente 192.168.1.2:
  Requisições: 101/100
  Status: rate limited (retry-after: 5s)
```

## Dicas
- Token bucket: adiciona tokens periodicamente
- Sincronização: use `sync.Mutex` ou atomic
- Backend compartilhado: arquivo, Redis (simulado), ou sync.Map
- Sliding window: mantém timestamps
- Burst: permite consumir tokens futuros

## Desafio Extra
Implemente com Redis real e suporte a cluster.



