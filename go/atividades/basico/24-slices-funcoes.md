# Atividade 24: Slices e Funções

## Objetivo
Praticar passagem de slices para funções e operações.

## Enunciado
Crie um programa que:
1. Crie uma função `encontrarMaximo` que recebe `[]int` e retorna o maior número
2. Crie uma função `filtrarPares` que recebe `[]int` e retorna novo slice apenas com números pares
3. Crie uma função `somarTodos` que recebe `[]int` e retorna a soma
4. No `main`, crie um slice [3, 7, 2, 9, 4, 6] e chame todas as funções

## Exemplo de Saída
```
Slice: [3 7 2 9 4 6]
Máximo: 9
Apenas pares: [2 4 6]
Soma: 31
```

## Dicas
- Slice é passado por referência (modificações afetam original)
- Para não modificar, crie novo slice
- Use `range` para iterar
- Use `append` para adicionar a novo slice



