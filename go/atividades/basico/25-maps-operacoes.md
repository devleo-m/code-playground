# Atividade 25: Operações com Maps

## Objetivo
Praticar operações avançadas com maps.

## Enunciado
Crie um programa que:
1. Crie um map de string para int representando estoque de produtos
2. Adicione produtos: "Maçã": 50, "Banana": 30, "Laranja": 40
3. Crie função `verificarEstoque` que recebe map, produto (string) e retorna (quantidade, bool)
4. Crie função `adicionarEstoque` que adiciona quantidade ao produto
5. Crie função `listarProdutos` que imprime todos os produtos e quantidades
6. Teste todas as funções

## Exemplo de Saída
```
Produto Maçã: 50 unidades (existe)
Produto Uva: 0 unidades (não existe)
Após adicionar 20 Maçãs:
Estoque:
  Maçã: 70
  Banana: 30
  Laranja: 40
```

## Dicas
- Verificar existência: `valor, ok := map[chave]`
- Adicionar/Modificar: `map[chave] = valor`
- Iterar: `for chave, valor := range map { }`
- Deletar: `delete(map, chave)`



