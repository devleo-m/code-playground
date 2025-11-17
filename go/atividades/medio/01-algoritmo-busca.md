# Atividade 01: Algoritmo de Busca em Slice

## Objetivo
Implementar algoritmos de busca (linear e binária) e comparar performance.

## Enunciado
Crie um programa que:
1. Implemente função `buscaLinear(slice []int, valor int) (int, bool)` que retorna índice e se encontrou
2. Implemente função `buscaBinaria(slice []int, valor int) (int, bool)` (slice deve estar ordenado)
3. Crie função `ordenarSlice(slice []int) []int` que retorna novo slice ordenado (use algoritmo simples ou `sort.Ints`)
4. Teste ambas as buscas com slice de 1000 elementos
5. Compare número de comparações necessárias (adicionar contador)

## Exemplo de Saída
```
Slice original: [42, 15, 8, 23, 16, ...]
Buscando 42:
  Busca linear: encontrado no índice 0 (comparações: 1)
  Busca binária: encontrado no índice 500 (comparações: 10)
Buscando 9999:
  Busca linear: não encontrado (comparações: 1000)
  Busca binária: não encontrado (comparações: 10)
```

## Dicas
- Busca linear: itere do início ao fim
- Busca binária: divida ao meio, compare, continue na metade apropriada
- Busca binária requer slice ordenado
- Adicione variável para contar comparações

## Desafio Extra
Implemente busca binária recursiva também.


