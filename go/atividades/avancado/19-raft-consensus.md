# Atividade 19: Algoritmo Raft Simplificado

## Objetivo
Implementar algoritmo de consenso Raft básico para eleição de líder.

## Enunciado
Crie implementação simplificada de Raft:
1. Defina estados: Follower, Candidate, Leader
2. Implemente eleição de líder:
   - Timeout aleatório para iniciar eleição
   - RequestVote RPC
   - Maioria de votos para se tornar líder
3. Implemente heartbeat do líder
4. Log replication básico (simplificado)
5. Implemente 3-5 nodes (simulados)
6. Teste eleição quando líder falha
7. Adicione logs de transições de estado

## Exemplo de Saída
```
Raft Cluster: 5 nodes
Node 1: Follower
Node 2: Follower
Node 3: Leader (termo 1)
Node 4: Follower
Node 5: Follower

Leader 3 falhou...
Node 1: Candidate (termo 2)
Solicitando votos...
Node 1 recebeu 3 votos -> Leader
Novo líder: Node 1 (termo 2)
```

## Dicas
- Raft: algoritmo de consenso distribuído
- Eleição: candidate precisa maioria
- Heartbeat: líder mantém autoridade
- Termo: incrementa a cada eleição
- Simplificado: sem log replication completo

## Desafio Extra
Implemente log replication completo e snapshot.


