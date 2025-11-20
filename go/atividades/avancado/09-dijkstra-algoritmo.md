# Atividade 09: Algoritmo de Dijkstra

## Objetivo
Implementar algoritmo de Dijkstra para encontrar caminho mais curto em grafo ponderado.

## Enunciado
Crie implementação completa de Dijkstra:
1. Defina struct `Grafo` com lista de adjacência ponderada: `map[string]map[string]int`
2. Implemente `CaminhoMaisCurto(origem, destino string) ([]string, int, error)`:
   - Retorna caminho e distância total
   - Usa priority queue (min-heap)
   - Mantém distâncias e predecessores
3. Implemente min-heap para priority queue eficiente
4. Adicione suporte a grafo direcionado e não-direcionado
5. Retorne todos os caminhos mais curtos de um vértice (single-source)
6. Visualize grafo e caminho encontrado (formato texto)
7. Teste com grafos de diferentes tamanhos (10, 100, 1000 vértices)

## Exemplo de Saída
```
Grafo: A->B (peso 4), A->C (peso 2), B->D (peso 5), C->D (peso 1)
Caminho mais curto de A para D:
  Caminho: A -> C -> D
  Distância: 3
  Tempo: 0.05ms
```

## Dicas
- Priority queue: sempre processa vértice com menor distância
- Relaxação: atualiza distância se encontrar caminho menor
- Marque vértices visitados
- Predecessores para reconstruir caminho
- Complexidade: O((V+E) log V) com heap

## Desafio Extra
Implemente A* (A-star) com heurística e compare com Dijkstra.



