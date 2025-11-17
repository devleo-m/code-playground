# Atividade 02: Sistema de Filtros Avançado

## Objetivo
Criar sistema flexível de filtros usando funções como parâmetros.

## Enunciado
Crie um programa que:
1. Defina struct `Produto` com Nome, Preco, Categoria, Estoque
2. Crie função `filtrar` que recebe `[]Produto` e função `func(Produto) bool`, retorna slice filtrado
3. Crie funções de filtro:
   - `precoMenorQue(preco float64) func(Produto) bool`
   - `categoriaIgual(cat string) func(Produto) bool`
   - `estoqueMaiorQue(qtd int) func(Produto) bool`
4. Crie função `filtrarMultiplos` que aplica múltiplos filtros (AND lógico)
5. Teste combinando filtros diferentes

## Exemplo de Saída
```
Produtos com preço < 50.00:
  - Notebook: R$ 2500.00 (não)
  - Mouse: R$ 25.00 (sim)
  - Teclado: R$ 45.00 (sim)

Produtos: categoria='Eletrônicos' AND estoque > 10:
  - Mouse: R$ 25.00, Estoque: 50
```

## Dicas
- Funções como valores: `type Filtro func(Produto) bool`
- Closure: função retorna função que captura parâmetros
- Combine filtros: `produto := range produtos { if f1(produto) && f2(produto) { ... } }`


