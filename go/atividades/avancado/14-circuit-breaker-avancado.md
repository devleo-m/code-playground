# Atividade 14: Circuit Breaker Avançado

## Objetivo
Implementar Circuit Breaker robusto com métricas e múltiplas estratégias.

## Enunciado
Crie Circuit Breaker avançado:
1. Estados: Closed, Open, HalfOpen com transições automáticas
2. Múltiplas estratégias de abertura:
   - Por contagem de falhas consecutivas
   - Por taxa de erro (percentual)
   - Por latência (timeout)
3. Implemente métricas:
   - Total de requisições
   - Taxa de sucesso/falha
   - Latência média/p95/p99
   - Tempo em cada estado
4. Adicione callbacks: OnOpen, OnClose, OnHalfOpen
5. Implemente retry com exponential backoff quando half-open
6. Suporte a múltiplos circuit breakers (por serviço/endpoint)
7. Exporte métricas para monitoramento
8. Teste com simulação de serviço instável

## Exemplo de Saída
```
Circuit Breaker: api.exemplo.com
Estado: Open
Falhas consecutivas: 5/5
Taxa de erro: 100%
Latência p95: 5.2s
Última falha: 2s atrás
Timeout: 30s
Próxima tentativa em: 28s
```

## Dicas
- Estados com transições baseadas em condições
- Métricas para decisões inteligentes
- Timeout antes de tentar novamente
- Half-open para testar recuperação
- Callbacks para integração com observabilidade

## Desafio Extra
Implemente adaptive circuit breaker que ajusta thresholds automaticamente.



