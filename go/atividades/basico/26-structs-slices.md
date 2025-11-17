# Atividade 26: Structs e Slices

## Objetivo
Praticar uso de slices de structs.

## Enunciado
Crie um programa que:
1. Defina struct `Pessoa` com Nome (string) e Idade (int)
2. Crie um slice de `Pessoa` com pelo menos 3 pessoas
3. Crie função `encontrarMaisVelho` que recebe `[]Pessoa` e retorna a pessoa mais velha
4. Crie função `filtrarPorIdade` que recebe `[]Pessoa` e idade mínima, retornando pessoas que atendem
5. No `main`, teste as funções

## Exemplo de Saída
```
Pessoas:
  João, 25 anos
  Maria, 30 anos
  Pedro, 22 anos
Mais velho: Maria, 30 anos
Pessoas com 25+ anos:
  João, 25 anos
  Maria, 30 anos
```

## Dicas
- Slice de structs: `[]Pessoa{ Pessoa{...}, Pessoa{... }`
- Itere com `for range`
- Compare campos da struct
- Use `append` para criar novo slice filtrado


