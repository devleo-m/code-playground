# Atividade 43: Rate Limiter

## Objetivo
Implementar limitador de taxa de requisições.

## Enunciado
Crie um programa que:
1. Defina struct `RateLimiter` com limite de requisições por período
2. Use map para rastrear requisições por chave (ex: IP, usuário)
3. Implemente método `Permitir(chave string) bool` que:
   - Retorna `true` se dentro do limite
   - Retorna `false` se excedeu limite
   - Limpa requisições antigas (fora do período)
4. Implemente algoritmo Token Bucket ou Sliding Window
5. Teste com diferentes limites e períodos

## Exemplo de Saída
```
Rate Limiter: 5 requisições por 10 segundos
Requisição 1 (IP: 192.168.1.1): permitida ✓
Requisição 2: permitida ✓
...
Requisição 6: negada ✗ (limite excedido)
Aguardando 10 segundos...
Requisição 7: permitida ✓
```

## Dicas
- Token Bucket: adiciona tokens periodicamente
- Sliding Window: mantém timestamps das requisições
- Limpe requisições antigas periodicamente
- Use `time` para gerenciar períodos



