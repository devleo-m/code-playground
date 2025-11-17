# Atividade 43: Maps Aninhados

## Objetivo
Praticar maps aninhados e estruturas complexas.

## Enunciado
Crie um programa que:
1. Crie map aninhado: `map[string]map[string]int` representando estoque por categoria
2. Estrutura: categoria -> produto -> quantidade
3. Adicione dados:
   - "Frutas" -> "Maçã": 50, "Banana": 30
   - "Verduras" -> "Alface": 20, "Tomate": 40
4. Crie função `obterEstoque` que recebe categoria e produto, retorna (quantidade, bool)
5. Crie função `listarCategoria` que imprime todos produtos de uma categoria

## Exemplo de Saída
```
Estoque de Frutas/Maçã: 50 (encontrado)
Estoque de Frutas/Uva: 0 (não encontrado)
Produtos em Frutas:
  Maçã: 50
  Banana: 30
```

## Dicas
- Map aninhado: `map[string]map[string]int`
- Acesse: `map[categoria][produto]`
- Verifique existência de categoria primeiro
- Inicialize map interno se não existir: `map[categoria] = make(map[string]int)`


