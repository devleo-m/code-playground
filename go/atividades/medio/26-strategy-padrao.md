# Atividade 26: Padrão Strategy

## Objetivo
Implementar padrão Strategy usando interfaces.

## Enunciado
Crie um programa que:
1. Defina interface `EstrategiaDesconto` com método `CalcularDesconto(preco float64) float64`
2. Implemente estratégias: `DescontoFixo`, `DescontoPercentual`, `DescontoVIP`
3. Defina struct `Carrinho` com preco total e estratégia de desconto
4. Implemente método `AplicarDesconto()` que usa estratégia atual
5. Permita trocar estratégia dinamicamente
6. Calcule preço final com diferentes estratégias

## Exemplo de Saída
```
Carrinho: R$ 1000.00
Desconto fixo R$ 50: R$ 950.00
Desconto 10%: R$ 900.00
Desconto VIP 20%: R$ 800.00
```

## Dicas
- Interface define contrato comum
- Cada estratégia implementa interface
- Troque estratégia atribuindo nova implementação
- Útil para algoritmos intercambiáveis

