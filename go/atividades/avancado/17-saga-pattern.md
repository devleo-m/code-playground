# Atividade 17: Saga Pattern para Transações Distribuídas

## Objetivo
Implementar padrão Saga para gerenciar transações distribuídas com compensação.

## Enunciado
Crie sistema Saga:
1. Defina interface `Step` com métodos:
   - `Executar() error`
   - `Compensar() error` - ação de rollback
2. Defina struct `Saga` que orquestra steps
3. Implemente dois tipos de Saga:
   - Choreography (cada step publica eventos)
   - Orchestration (orquestrador central)
4. Implemente compensação automática em caso de falha
5. Adicione persistência de estado (para recovery)
6. Suporte a steps condicionais (if/else)
7. Implemente timeout por step
8. Teste com saga de 5 steps onde step 3 falha

## Exemplo de Saída
```
Saga: Criar Pedido
Step 1: Reservar Estoque... ✓
Step 2: Processar Pagamento... ✓
Step 3: Enviar Email... ✗ (falha)
Compensando:
  Step 2: Reverter Pagamento... ✓
  Step 1: Liberar Estoque... ✓
Saga compensada com sucesso
```

## Dicas
- Saga: sequência de transações locais
- Cada step tem compensação
- Falha: compensa steps anteriores
- Persistência: permite retry após crash
- Orchestration: mais controle, mais complexo

## Desafio Extra
Implemente saga paralela (steps executam em paralelo) e idempotência.



