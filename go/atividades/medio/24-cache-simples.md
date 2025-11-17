# Atividade 24: Sistema de Cache Simples

## Objetivo
Implementar cache com TTL (Time To Live) usando maps e time.

## Enunciado
Crie um programa que:
1. Defina struct `CacheItem` com Valor, ExpiraEm (time.Time)
2. Defina struct `Cache` com map e mutex (para thread-safety básico)
3. Implemente métodos:
   - `Set(chave string, valor interface{}, ttl time.Duration)`
   - `Get(chave string) (interface{}, bool)` - retorna false se expirado
   - `Delete(chave string)`
   - `Clear()` - limpa cache
   - `Size() int`
4. Implemente limpeza automática de itens expirados
5. Teste com diferentes TTLs

## Exemplo de Saída
```
Cache.Set("chave1", "valor1", 5*time.Second)
Cache.Get("chave1"): valor1 (true)
Aguardando 6 segundos...
Cache.Get("chave1"): (false - expirado)
Cache limpo automaticamente
```

## Dicas
- Use `time.Now().Add(ttl)` para calcular expiração
- Verifique `time.Now().After(item.ExpiraEm)` para expiração
- Use goroutine para limpeza periódica (opcional)
- Map precisa de mutex se usado concorrentemente


