# Atividade 23: Algoritmos em Grafos

## Objetivo
Implementar algoritmos básicos de grafos.

## Enunciado
Crie um programa que:
1. Represente grafo como `map[string][]string` (lista de adjacência)
2. Implemente DFS (busca em profundidade) recursiva
3. Implemente BFS (busca em largura) usando fila
4. Implemente função `caminhoExiste(grafo map[string][]string, inicio, destino string) bool`
5. Implemente função `componentesConectados(grafo map[string][]string) [][]string`
6. Encontre caminho mais curto entre dois nós (BFS)

## Exemplo de Saída
```
Grafo: A->[B,C], B->[D], C->[D], D->[]
DFS de A: A, B, D, C
BFS de A: A, B, C, D
Caminho A->D existe: true
Componentes: [[A,B,C,D], [E,F]]
```

## Dicas
- DFS: recursão ou pilha
- BFS: fila
- Marque nós visitados para evitar loops
- Use map para rastrear caminho


