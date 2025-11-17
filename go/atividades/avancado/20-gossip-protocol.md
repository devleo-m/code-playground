# Atividade 20: Gossip Protocol

## Objetivo
Implementar protocolo Gossip para disseminação de informações em cluster.

## Enunciado
Crie sistema Gossip:
1. Defina struct `GossipNode` com estado local e vizinhos
2. Implemente disseminação:
   - Push: envia estado para vizinho
   - Pull: solicita estado de vizinho
   - Push-Pull: combinação
3. Selecione vizinhos aleatoriamente a cada round
4. Implemente convergência: todos nodes têm mesmo estado
5. Adicione anti-entropy (corrige divergências)
6. Suporte a diferentes tipos de mensagens (rumor mongering)
7. Teste com 10 nodes e verifique tempo de convergência
8. Adicione métricas: rounds até convergência, mensagens trocadas

## Exemplo de Saída
```
Gossip Cluster: 10 nodes
Estado inicial: Node1 tem informação nova
Round 1: Node1 -> Node3, Node5
Round 2: Node3 -> Node7, Node1 -> Node2
Round 3: Node5 -> Node9, Node2 -> Node4
...
Convergência alcançada: Round 5
Todos nodes sincronizados
Mensagens totais: 23
```

## Dicas
- Gossip: disseminação epidêmica
- Seleção aleatória de pares
- Convergência eventual
- Anti-entropy: reconciliação periódica
- Útil para membership e state dissemination

## Desafio Extra
Implemente failure detection e membership management.


