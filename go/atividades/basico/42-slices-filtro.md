# Atividade 42: Filtrar e Transformar Slices

## Objetivo
Praticar operações de filtro e transformação em slices.

## Enunciado
Crie um programa que:
1. Crie função `filtrarMaioresQue` que recebe `[]int` e valor mínimo, retorna novo slice
2. Crie função `dobrarValores` que recebe `[]int` e retorna novo slice com valores dobrados
3. Crie função `encontrarIndices` que recebe `[]int` e valor, retorna slice de índices onde valor aparece
4. No `main`, teste com slice [5, 10, 15, 20, 10, 25, 10]

## Exemplo de Saída
```
Slice original: [5 10 15 20 10 25 10]
Valores > 12: [15 20 25]
Valores dobrados: [10 20 30 40 20 50 20]
Índices onde 10 aparece: [1 4 6]
```

## Dicas
- Filtro: crie novo slice, adicione elementos que atendem condição
- Transformação: crie novo slice, aplique operação a cada elemento
- Use `append` para adicionar elementos
- Use `range` para iterar


