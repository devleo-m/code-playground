# Atividade 14: Loops - range

## Objetivo
Praticar loops `for range` para iterar sobre coleções.

## Enunciado
Crie um programa que:
1. Crie um slice de strings: ["Maçã", "Banana", "Laranja"]
2. Use `for range` para iterar e imprimir cada fruta com seu índice
3. Crie um map: map[string]int{"João": 25, "Maria": 30, "Pedro": 22}
4. Use `for range` para iterar e imprimir cada nome e idade
5. Use `for range` em uma string para imprimir cada caractere

## Exemplo de Saída
```
Frutas:
0: Maçã
1: Banana
2: Laranja

Pessoas:
João: 25 anos
Maria: 30 anos
Pedro: 22 anos

Caracteres de "Go":
G
o
```

## Dicas
- `for indice, valor := range slice { }`
- `for chave, valor := range map { }`
- `for indice, rune := range string { }`
- Use `_` para ignorar valores não usados

