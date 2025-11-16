# Atividade 49: Comparação de Structs

## Objetivo
Praticar comparação de structs e entender quando é possível.

## Enunciado
Crie um programa que:
1. Defina struct `Ponto` com X e Y (int) - comparável
2. Defina struct `Pessoa` com Nome (string) e Hobbies ([]string) - não comparável (tem slice)
3. Crie dois pontos iguais e compare com `==`
4. Crie dois pontos diferentes e compare
5. Tente comparar duas pessoas (deve dar erro de compilação - comente essa parte)
6. Crie função `pessoasIguais` que compara pessoas manualmente

## Exemplo de Saída
```
Ponto (1,2) == Ponto (1,2): true
Ponto (1,2) == Ponto (3,4): false
Pessoas com mesmo nome e hobbies: iguais (função customizada)
```

## Dicas
- Structs são comparáveis se todos campos forem comparáveis
- Tipos comparáveis: básicos, arrays, outras structs comparáveis
- Tipos não comparáveis: slices, maps, funções
- Use função customizada para comparar structs não comparáveis

