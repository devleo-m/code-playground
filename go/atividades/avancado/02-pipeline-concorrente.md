# Atividade 02: Pipeline Concorrente

## Objetivo
Criar pipeline de processamento de dados totalmente concorrente com múltiplos estágios.

## Enunciado
Crie um pipeline concorrente que:
1. Defina interface `Stage` com método `Processar(input <-chan Item) <-chan Item`
2. Implemente estágios:
   - `FiltroStage` - filtra itens baseado em critério
   - `TransformacaoStage` - transforma itens
   - `AgregacaoStage` - agrega itens
   - `ValidacaoStage` - valida itens
3. Crie função `CriarPipeline(stages ...Stage) Pipeline`
4. Pipeline processa dados em paralelo entre estágios
5. Implemente backpressure (limita produção se consumo está lento)
6. Adicione métricas: throughput por estágio, latência, buffer sizes
7. Teste com stream de 10.000 itens

## Exemplo de Saída
```
Pipeline: Filtro -> Transformação -> Agregação -> Validação
Processando stream de 10000 itens...
Estágio 1 (Filtro): 7500 itens processados (2500 filtrados)
Estágio 2 (Transformação): 7500 itens transformados
Estágio 3 (Agregação): 150 grupos criados
Estágio 4 (Validação): 148 válidos, 2 inválidos
Throughput: 5000 itens/segundo
Latência média: 2ms por item
```

## Dicas
- Cada estágio roda em goroutine separada
- Channels conectam estágios
- Buffered channels para pipeline
- Context para cancelamento
- Use `select` para multiplexar

## Desafio Extra
Implemente retry automático em estágios que falham e circuit breaker.


