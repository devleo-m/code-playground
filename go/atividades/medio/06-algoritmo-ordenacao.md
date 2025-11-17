# Atividade 06: Algoritmos de Ordenação

## Objetivo
Implementar algoritmos de ordenação e comparar performance.

## Enunciado
Crie um programa que:
1. Implemente `bubbleSort(slice []int) []int` (ordenação por bolha)
2. Implemente `selectionSort(slice []int) []int` (ordenação por seleção)
3. Implemente `quickSort(slice []int) []int` (ordenação rápida - recursiva)
4. Crie função que mede tempo de execução de cada algoritmo
5. Teste com slices de tamanhos diferentes (10, 100, 1000 elementos)
6. Compare resultados e performance

## Exemplo de Saída
```
Slice original: [64, 34, 25, 12, 22, 11, 90]
Bubble Sort: [11, 12, 22, 25, 34, 64, 90] (tempo: 2.5μs)
Selection Sort: [11, 12, 22, 25, 34, 64, 90] (tempo: 1.8μs)
Quick Sort: [11, 12, 22, 25, 34, 64, 90] (tempo: 0.5μs)

Com 1000 elementos:
  Bubble Sort: 15.2ms
  Selection Sort: 8.5ms
  Quick Sort: 1.2ms
```

## Dicas
- Bubble Sort: compare pares adjacentes, troque se necessário
- Selection Sort: encontre mínimo, coloque na posição correta
- Quick Sort: escolha pivô, particione, ordene recursivamente
- Use `time.Now()` e `time.Since()` para medir tempo


