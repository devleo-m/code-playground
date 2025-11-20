# Atividade 45: Mecanismo de Retry

## Objetivo
Implementar sistema de retry com backoff exponencial.

## Enunciado
Crie um programa que:
1. Defina struct `RetryConfig` com: MaxTentativas, DelayInicial, Multiplicador
2. Crie função `ExecutarComRetry(funcao func() error, config RetryConfig) error`
3. Implemente backoff exponencial: delay aumenta a cada tentativa
4. Suporte backoff fixo também
5. Adicione jitter (variação aleatória) no delay
6. Retorne erro apenas após esgotar todas tentativas
7. Simule chamada a API que pode falhar temporariamente

## Exemplo de Saída
```
Tentativa 1: falha, aguardando 1s...
Tentativa 2: falha, aguardando 2s...
Tentativa 3: falha, aguardando 4s...
Tentativa 4: sucesso! ✓
Total de tentativas: 4
```

## Dicas
- Backoff exponencial: `delay = inicial * (multiplicador ^ tentativa)`
- Jitter: adicione variação aleatória ao delay
- Use `time.Sleep` para aguardar
- Limite máximo de tentativas para evitar loops infinitos



