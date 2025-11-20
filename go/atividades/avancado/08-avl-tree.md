# Atividade 08: Árvore AVL (Auto-Balanceada)

## Objetivo
Implementar árvore AVL com rotações e balanceamento automático.

## Enunciado
Crie uma árvore AVL completa:
1. Defina struct `NoAVL` com Valor, Altura, Esquerda, Direita
2. Implemente rotações:
   - Rotação simples direita
   - Rotação simples esquerda
   - Rotação dupla direita-esquerda
   - Rotação dupla esquerda-direita
3. Implemente métodos:
   - `Inserir(valor int) *NoAVL` - insere e balanceia
   - `Remover(valor int) *NoAVL` - remove e balanceia
   - `Buscar(valor int) bool`
   - `Altura() int`
   - `Balanceamento() int` - retorna fator de balanceamento
4. Mantenha propriedade AVL: |altura_esq - altura_dir| <= 1
5. Implemente travessias: em-ordem, pré-ordem, pós-ordem
6. Teste com inserções e remoções que causam desbalanceamento

## Exemplo de Saída
```
Inserindo: 10, 20, 30, 40, 50, 25
Após inserções:
  Altura: 3
  Balanceamento: 0 (balanceada)
  Em ordem: [10, 20, 25, 30, 40, 50]

Removendo 40:
  Rotação necessária: simples esquerda
  Altura: 3
  Balanceamento: 0
```

## Dicas
- Fator de balanceamento: altura_esq - altura_dir
- Rotação quando |fator| > 1
- Atualize altura após cada operação
- Casos: LL, RR, LR, RL
- Complexidade: O(log n) para todas operações

## Desafio Extra
Implemente Red-Black Tree também e compare performance.



