# Atividade 31: Agregador de Dados

## Objetivo
Criar sistema de agregação e análise de dados.

## Enunciado
Crie um programa que:
1. Defina struct `Venda` com Produto, Quantidade, Preco, Data, Vendedor
2. Crie funções de agregação:
   - `totalVendas(vendas []Venda) float64`
   - `mediaVendas(vendas []Venda) float64`
   - `vendasPorVendedor(vendas []Venda) map[string]float64`
   - `produtoMaisVendido(vendas []Venda) (string, int)`
   - `vendasPorPeriodo(vendas []Venda, inicio, fim time.Time) []Venda`
3. Crie função `gerarRelatorio(vendas []Venda) string` que formata relatório completo
4. Implemente ordenação por diferentes critérios (preço, quantidade, data)

## Exemplo de Saída
```
=== Relatório de Vendas ===
Total: R$ 15.500,00
Média: R$ 1.550,00
Vendedor do mês: João (R$ 8.000,00)
Produto mais vendido: Notebook (15 unidades)
```

## Dicas
- Use `sort.Slice` para ordenação customizada
- Agrupe com maps
- Calcule estatísticas iterando sobre dados
- Formate saída com `fmt.Sprintf`

