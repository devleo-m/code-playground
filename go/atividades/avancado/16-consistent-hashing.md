# Atividade 16: Consistent Hashing

## Objetivo
Implementar consistent hashing para distribuição de carga em sistemas distribuídos.

## Enunciado
Crie sistema de consistent hashing:
1. Defina struct `ConsistentHash` com ring (map de hash para node)
2. Implemente `AdicionarNode(node string, peso int)` - adiciona ao ring
3. Implemente `RemoverNode(node string)` - remove do ring
4. Implemente `ObterNode(chave string) string` - retorna node responsável
5. Use hash function (SHA256 ou MD5) para posicionar nodes e keys
6. Implemente virtual nodes (replicas) para melhor distribuição
7. Adicione suporte a pesos (nodes com diferentes capacidades)
8. Calcule distribuição: verifique se carga está balanceada
9. Teste adicionando/removendo nodes dinamicamente

## Exemplo de Saída
```
Consistent Hash Ring: 3 nodes, 150 virtual nodes
Adicionando 1000 chaves...
Distribuição:
  Node1: 334 chaves (33.4%)
  Node2: 333 chaves (33.3%)
  Node3: 333 chaves (33.3%)
Removendo Node2...
Redistribuição:
  Node1: 500 chaves (50%)
  Node3: 500 chaves (50%)
Chaves movidas: 333 (33.3%)
```

## Dicas
- Ring circular: hash values de 0 a 2^32-1
- Virtual nodes: múltiplas posições por node físico
- Busca: primeiro node com hash >= hash da chave
- Pesos: mais virtual nodes para nodes mais poderosos
- Útil para sharding e load balancing

## Desafio Extra
Implemente data migration quando nodes são adicionados/removidos e monitoramento de hotspots.



