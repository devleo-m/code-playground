# Atividade 12: Árvore de Merkle

## Objetivo
Implementar Merkle Tree para verificação de integridade de dados.

## Enunciado
Crie uma Merkle Tree completa:
1. Defina struct `MerkleNode` com Hash, Esquerda, Direita, Dados
2. Implemente construção da árvore:
   - Hash de cada folha (dados)
   - Hash de nós internos (hash(esq) + hash(dir))
   - Árvore binária completa
3. Métodos:
   - `CalcularRootHash() []byte`
   - `VerificarIntegridade() bool`
   - `GerarProvaMembro(dados []byte) ([]MerkleProof, error)` - Merkle proof
   - `VerificarProva(prova []MerkleProof, rootHash []byte, dados []byte) bool`
4. Implemente serialização da árvore
5. Adicione suporte a atualização de folhas
6. Teste com diferentes tamanhos de dados (10, 100, 1000 itens)

## Exemplo de Saída
```
Merkle Tree: 8 folhas
Root hash: a3f5b2c1d4e6f7a8b9c0d1e2f3a4b5c6d7e8f9
Verificando integridade: ✓
Gerando prova para item 3:
  Proof path: [hash1, hash23, hash4567]
Verificando prova: ✓ válida
```

## Dicas
- Hash cada folha
- Nós internos: hash dos filhos
- Merkle proof: caminho da folha à raiz
- Verificação: reconstrói hash até raiz
- Útil para blockchains e sistemas distribuídos

## Desafio Extra
Implemente Sparse Merkle Tree e otimizações para grandes datasets.



