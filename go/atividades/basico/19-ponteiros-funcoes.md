# Atividade 19: Ponteiros em Funções

## Objetivo
Praticar passagem de ponteiros para funções.

## Enunciado
Crie um programa que:
1. Crie uma função `dobrar` que recebe um ponteiro para `int` e dobra o valor
2. Crie uma função `trocar` que recebe dois ponteiros para `int` e troca seus valores
3. No `main`, declare duas variáveis `a = 10` e `b = 20`
4. Chame `dobrar` para dobrar `a`
5. Chame `trocar` para trocar `a` e `b`
6. Imprima os valores finais

## Exemplo de Saída
```
Valores iniciais: a=10, b=20
Após dobrar a: a=20, b=20
Após trocar: a=20, b=20
```

## Dicas
- Parâmetro ponteiro: `func nome(ptr *tipo) { }`
- Passar ponteiro: `funcao(&variavel)`
- Modificar: `*ptr = novoValor`
- Use ponteiros quando precisar modificar valores originais



