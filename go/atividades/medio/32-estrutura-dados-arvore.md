# Atividade 32: Estrutura de Dados - Árvore Binária

## Objetivo
Implementar árvore binária de busca.

## Enunciado
Crie um programa que:
1. Defina struct `No` com Valor (int), Esquerda, Direita (*No)
2. Defina struct `Arvore` com Raiz (*No)
3. Implemente métodos:
   - `Inserir(valor int)` - insere mantendo ordem
   - `Buscar(valor int) bool` - busca valor
   - `EmOrdem() []int` - percorre em ordem (esquerda, raiz, direita)
   - `PreOrdem() []int` - percorre pré-ordem (raiz, esquerda, direita)
   - `PosOrdem() []int` - percorre pós-ordem (esquerda, direita, raiz)
4. Implemente `Altura() int` que calcula altura da árvore
5. Teste com diferentes valores

## Exemplo de Saída
```
Inserindo: 5, 3, 7, 2, 4, 6, 8
Em ordem: [2, 3, 4, 5, 6, 7, 8]
Pré-ordem: [5, 3, 2, 4, 7, 6, 8]
Altura: 3
Buscar 4: true
Buscar 9: false
```

## Dicas
- Árvore binária: valores menores à esquerda, maiores à direita
- Use recursão para percorrer
- Insira comparando com nó atual
- Altura: máximo entre subárvores + 1

