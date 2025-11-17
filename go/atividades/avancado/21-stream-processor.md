# Atividade 21: Stream Processor

## Objetivo
Criar processador de streams de dados com janelas deslizantes e agregações.

## Enunciado
Crie stream processor:
1. Defina `Stream` interface com método `Next() (Item, error)`
2. Implemente operadores:
   - `Filtrar(stream Stream, predicado func(Item) bool) Stream`
   - `Mapear(stream Stream, transformacao func(Item) Item) Stream`
   - `AgruparPor(stream Stream, chave func(Item) string) Stream`
   - `JanelaDeslizante(stream Stream, tamanho time.Duration) Stream`
3. Implemente agregações em janelas:
   - Soma, média, máximo, mínimo
   - Contagem de eventos
   - Top N elementos
4. Suporte a múltiplos streams (join)
5. Adicione watermarking para eventos tardios
6. Teste com stream de 100.000 eventos

## Exemplo de Saída
```
Stream Processor iniciado
Processando stream de eventos...
Janela deslizante: 1 minuto
Agregações:
  Eventos: 1500
  Soma: 45000
  Média: 30.0
  Máximo: 100
  Top 3: [A: 500, B: 400, C: 300]
Throughput: 25000 eventos/segundo
```

## Dicas
- Stream: sequência infinita de eventos
- Janela: agrupa eventos por tempo
- Watermark: marca progresso do tempo
- Lazy evaluation: processa sob demanda
- Backpressure: controla velocidade

## Desafio Extra
Implemente exactly-once semantics e checkpointing.


