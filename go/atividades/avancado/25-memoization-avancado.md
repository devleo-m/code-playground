# Atividade 25: Memoization Avançado

## Objetivo
Implementar sistema de memoization com TTL, invalidação e cache distribuído.

## Enunciado
Crie memoization avançado:
1. Defina função `Memoize(funcao func(...interface{}) interface{}, ttl time.Duration) func(...interface{}) interface{}`
2. Cache com chave baseada em argumentos
3. Suporte a TTL (expiração automática)
4. Implemente invalidação:
   - Por chave
   - Por padrão
   - Por tags
5. Adicione métricas: hit rate, miss rate, tamanho cache
6. Implemente cache distribuído (simulado)
7. Suporte a funções assíncronas
8. Teste com função custosa (Fibonacci, busca complexa)

## Exemplo de Saída
```
Memoization: TTL 5min
Chamadas: 1000
Cache hits: 850 (85%)
Cache misses: 150 (15%)
Tempo médio (hit): 0.01ms
Tempo médio (miss): 50ms
Economia de tempo: 42.5s
```

## Dicas
- Memoization: cache de resultados de função
- Chave: hash dos argumentos
- TTL: expira resultados antigos
- Invalidação: remove entradas específicas
- Tags: agrupa entradas relacionadas

## Desafio Extra
Implemente cache com LRU eviction e size limits.

