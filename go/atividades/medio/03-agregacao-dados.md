# Atividade 03: Agregação e Agrupamento de Dados

## Objetivo
Implementar operações de agregação (soma, média, máximo, mínimo) e agrupamento.

## Enunciado
Crie um programa que:
1. Defina struct `Venda` com Produto, Quantidade, Preco, Data
2. Crie função `agruparPorProduto(vendas []Venda) map[string][]Venda` que agrupa vendas por produto
3. Crie função `calcularTotalPorProduto(vendas []Venda) map[string]float64` que calcula total vendido por produto
4. Crie função `produtoMaisVendido(vendas []Venda) (string, int)` que retorna produto e quantidade total
5. Crie função `vendasPorPeriodo(vendas []Venda, dataInicio, dataFim string) []Venda` que filtra por período

## Exemplo de Saída
```
Vendas agrupadas por produto:
  Notebook: 3 vendas
  Mouse: 5 vendas
  Teclado: 2 vendas

Total por produto:
  Notebook: R$ 7500.00
  Mouse: R$ 125.00
  Teclado: R$ 90.00

Produto mais vendido: Mouse (5 unidades)
```

## Dicas
- Use map para agrupar: `map[string][]Venda`
- Itere sobre map: `for chave, valor := range map { }`
- Combine filtros e agregações
- Use structs para organizar dados



