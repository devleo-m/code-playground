# Atividade 44: Circuit Breaker Pattern

## Objetivo
Implementar padrão Circuit Breaker para resiliência.

## Enunciado
Crie um programa que:
1. Defina estados: Fechado, Aberto, MeioAberto
2. Defina struct `CircuitBreaker` com:
   - Limite de falhas
   - Timeout para tentar novamente
   - Contador de falhas
3. Implemente método `Executar(funcao func() error) error` que:
   - Se fechado: executa função, conta falhas
   - Se aberto: retorna erro imediatamente
   - Se meio-aberto: tenta executar, se sucesso volta para fechado
4. Simule chamadas a serviço externo que pode falhar
5. Demonstre transição entre estados

## Exemplo de Saída
```
Estado: Fechado
Executando... sucesso ✓
Executando... falha ✗ (1/3)
Executando... falha ✗ (2/3)
Executando... falha ✗ (3/3)
Estado: Aberto (circuito aberto)
Tentando após timeout...
Estado: MeioAberto
Executando... sucesso ✓
Estado: Fechado
```

## Dicas
- Estados: Fechado (normal), Aberto (bloqueado), MeioAberto (testando)
- Conte falhas consecutivas
- Use timer para timeout
- Útil para serviços externos instáveis

