# Atividade 15: Batch Processor com Backpressure

## Objetivo
Criar processador de batches que agrupa itens e processa em lotes com controle de backpressure.

## Enunciado
Crie batch processor:
1. Defina struct `BatchProcessor` com:
   - Tamanho máximo do batch
   - Intervalo máximo (timeout)
   - Buffer de itens pendentes
2. Implemente `Processar(item Item) error` que:
   - Adiciona ao buffer
   - Dispara batch quando cheio ou timeout
   - Bloqueia se buffer muito cheio (backpressure)
3. Processamento assíncrono de batches
4. Implemente diferentes estratégias de batching:
   - Por tamanho (N itens)
   - Por tempo (T segundos)
   - Por tamanho + tempo (qualquer condição)
5. Adicione métricas: batches processados, itens por batch, latência
6. Implemente graceful shutdown (processa batches pendentes)
7. Teste com 100.000 itens

## Exemplo de Saída
```
Batch Processor: tamanho 100, timeout 1s
Processando 100000 itens...
Batches criados: 1000
Itens por batch (média): 100
Latência média: 15ms
Throughput: 6666 itens/segundo
Backpressure ativado: 2 vezes
```

## Dicas
- Buffer com channel ou slice
- Timer para timeout
- Goroutine para processamento
- Select para múltiplas condições
- Backpressure: bloqueia quando buffer cheio
- Context para shutdown

## Desafio Extra
Implemente priorização de batches e múltiplos workers para processamento paralelo.

