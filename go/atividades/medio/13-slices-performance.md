# Atividade 13: Otimização de Slices

## Objetivo
Entender e otimizar uso de slices para melhor performance.

## Enunciado
Crie um programa que:
1. Compare performance de diferentes formas de criar e preencher slices:
   - `make([]int, 0)` e depois `append`
   - `make([]int, tamanho)` e preencher índices
   - `make([]int, 0, capacidade)` com capacity pré-definida
2. Meça tempo para adicionar 10000 elementos em cada abordagem
3. Crie função `removerDuplicatas(slice []int) []int` que remove duplicatas eficientemente
4. Crie função `inverterSlice(slice []int) []int` que inverte ordem (in-place se possível)
5. Compare uso de memória (capacity) em cada abordagem

## Exemplo de Saída
```
Adicionando 10000 elementos:
  Sem capacity: 2.5ms, capacity final: 16384
  Com capacity: 0.8ms, capacity final: 10000
  Pre-alocado: 0.3ms, capacity final: 10000

Removendo duplicatas de [1,2,2,3,3,3]:
  Resultado: [1,2,3] (3 elementos)
```

## Dicas
- Pré-alocar capacity evita realocações
- `make([]tipo, length, capacity)` otimiza quando sabe tamanho
- Use map para remover duplicatas eficientemente
- Inverter in-place: troque elementos das extremidades

