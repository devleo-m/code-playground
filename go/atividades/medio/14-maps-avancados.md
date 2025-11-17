# Atividade 14: Operações Avançadas com Maps

## Objetivo
Implementar operações complexas com maps e estruturas aninhadas.

## Enunciado
Crie um programa que:
1. Crie map aninhado: `map[string]map[string]int` representando vendas por vendedor e produto
2. Crie função `adicionarVenda(vendas map[string]map[string]int, vendedor, produto string, quantidade int)`
3. Crie função `totalPorVendedor(vendas map[string]map[string]int, vendedor string) int`
4. Crie função `produtoMaisVendido(vendas map[string]map[string]int) (string, int)`
5. Crie função `vendedoresOrdenadosPorTotal(vendas map[string]map[string]int) []string`
6. Implemente inicialização segura de maps aninhados (verificar se existe antes de usar)

## Exemplo de Saída
```
Vendas:
  João -> Notebook: 5, Mouse: 10
  Maria -> Notebook: 3, Teclado: 8

Total de João: 15 unidades
Produto mais vendido: Mouse (10 unidades)
Vendedores ordenados: [Maria: 11, João: 15]
```

## Dicas
- Inicialize map interno: `if map[chave] == nil { map[chave] = make(...) }`
- Itere sobre maps: `for chave, valor := range map { }`
- Use slice para ordenar chaves do map
- Combine maps aninhados com agregações


