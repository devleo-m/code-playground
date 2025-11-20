# Atividade 01: Worker Pool com Goroutines

## Objetivo
Implementar worker pool eficiente usando goroutines e channels para processamento paralelo.

## Enunciado
Crie um sistema de worker pool que:
1. Defina struct `WorkerPool` com número de workers, canal de jobs e canal de resultados
2. Implemente `CriarWorkerPool(numWorkers int) *WorkerPool`
3. Crie função `ProcessarJobs(pool *WorkerPool, jobs []Job) []Result` que:
   - Distribui jobs entre workers
   - Coleta resultados
   - Gerencia timeouts
4. Implemente `Job` interface com método `Executar() (Result, error)`
5. Adicione métricas: jobs processados, tempo médio, taxa de erro
6. Implemente graceful shutdown (aguarda workers terminarem)
7. Teste com 1000 jobs e compare performance sequencial vs paralelo

## Exemplo de Saída
```
Worker Pool: 4 workers
Processando 1000 jobs...
Jobs processados: 1000
Tempo total: 2.5s
Tempo médio por job: 2.5ms
Taxa de sucesso: 99.8%
Throughput: 400 jobs/segundo
```

## Dicas
- Use buffered channels para jobs
- Workers consomem do canal de jobs
- Results são enviados para canal de resultados
- Use `sync.WaitGroup` para sincronização
- Context para cancelamento e timeout

## Desafio Extra
Implemente priorização de jobs (priority queue) e load balancing entre workers.



