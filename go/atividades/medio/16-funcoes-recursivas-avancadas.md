# Atividade 16: Recursão Avançada

## Objetivo
Resolver problemas complexos usando recursão.

## Enunciado
Crie um programa que:
1. Implemente `torreHanoi(n int, origem, destino, auxiliar string) []string` que retorna sequência de movimentos
2. Implemente `permutacoes(elements []int) [][]int` que gera todas permutações
3. Implemente `buscaProfundidade(grafo map[string][]string, inicio, destino string) []string` (DFS)
4. Implemente `calcularExpressao(expr string) (float64, error)` que avalia expressão matemática simples (recursiva)
5. Compare performance recursiva vs iterativa para fatorial grande

## Exemplo de Saída
```
Torre de Hanoi (3 discos):
  Mover de A para C
  Mover de A para B
  Mover de C para B
  ...

Permutações de [1,2,3]:
  [[1,2,3], [1,3,2], [2,1,3], [2,3,1], [3,1,2], [3,2,1]]

Expressão "2+3*4": 14.00
```

## Dicas
- Recursão: função chama a si mesma com problema menor
- Caso base: condição de parada
- Recursão de cauda pode ser otimizada pelo compilador
- Cuidado com stack overflow em recursões profundas

