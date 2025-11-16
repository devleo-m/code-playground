# Atividade 35: Structs como Valores em Maps

## Objetivo
Praticar uso de structs como valores em maps.

## Enunciado
Crie um programa que:
1. Defina struct `Produto` com Nome (string) e Preco (float64)
2. Crie map de string para `Produto` (chave: código do produto)
3. Adicione produtos: "A001": {Nome: "Notebook", Preco: 2500.00}
4. Adicione mais 2 produtos
5. Crie função `buscarProduto` que recebe map e código, retorna (Produto, bool)
6. Crie função `calcularTotal` que recebe map e retorna soma de todos os preços
7. No `main`, teste as funções

## Exemplo de Saída
```
Produto A001: Notebook - R$ 2500.00 (encontrado)
Produto X999: não encontrado
Total de todos os produtos: R$ 5500.00
```

## Dicas
- Map com struct: `map[string]Produto`
- Adicionar: `map[chave] = Produto{...}`
- Buscar: `produto, ok := map[chave]`
- Iterar: `for chave, valor := range map { }`

