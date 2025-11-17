# Atividade 33: Slices - Length e Capacity

## Objetivo
Entender e praticar length vs capacity em slices.

## Enunciado
Crie um programa que:
1. Crie slice com `make([]int, 3, 6)` - length 3, capacity 6
2. Preencha os 3 elementos com valores [10, 20, 30]
3. Imprima slice, length e capacity
4. Adicione elementos até capacity ser atingida
5. Adicione mais um elemento (capacity será expandida)
6. Imprima slice, length e capacity após cada operação

## Exemplo de Saída
```
Slice inicial: [10 20 30]
Length: 3, Capacity: 6
Após adicionar até capacity: [10 20 30 40 50 60]
Length: 6, Capacity: 6
Após exceder capacity: [10 20 30 40 50 60 70]
Length: 7, Capacity: 12 (dobrou!)
```

## Dicas
- `make([]tipo, length, capacity)`
- `len()` retorna length, `cap()` retorna capacity
- Quando length = capacity, `append` cria novo array (capacity dobra)
- Capacity sempre >= length


