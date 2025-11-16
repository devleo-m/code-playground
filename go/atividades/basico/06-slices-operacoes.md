# Atividade 06: Operações com Slices

## Objetivo
Praticar operações avançadas com slices: slicing, length, capacity.

## Enunciado
Crie um programa que:
1. Crie um slice com `make([]int, 3, 5)` - length 3, capacity 5
2. Preencha os 3 primeiros elementos com valores [10, 20, 30]
3. Imprima o slice, seu length e capacity
4. Adicione mais 2 elementos usando `append`
5. Imprima novamente o slice, length e capacity
6. Crie um sub-slice dos primeiros 2 elementos
7. Imprima o sub-slice

## Exemplo de Saída
```
Slice inicial: [10 20 30]
Length: 3, Capacity: 5
Slice após append: [10 20 30 40 50]
Length: 5, Capacity: 5
Sub-slice: [10 20]
```

## Dicas
- `make([]tipo, length, capacity)` cria slice com tamanho específico
- `len()` retorna length, `cap()` retorna capacity
- Slicing: `slice[inicio:fim]` (fim não incluído)

