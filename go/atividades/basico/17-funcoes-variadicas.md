# Atividade 17: Funções Variádicas

## Objetivo
Praticar funções que aceitam número variável de argumentos.

## Enunciado
Crie um programa que:
1. Crie uma função `somarNumeros` que aceita números variáveis de `int` e retorna a soma
2. Crie uma função `juntarStrings` que aceita um separador (string) e strings variáveis, retornando a string juntada
3. No `main`, chame `somarNumeros` com diferentes quantidades de argumentos
4. Chame `juntarStrings` para criar uma frase

## Exemplo de Saída
```
Soma de 1, 2, 3: 6
Soma de 10, 20, 30, 40: 100
Soma de nenhum número: 0
Frase: Olá, mundo, Go!
```

## Dicas
- Variádica: `func nome(args ...tipo) tipoRetorno { }`
- Dentro da função, `args` é um slice
- Chame com: `funcao(1, 2, 3)` ou `funcao(slice...)`
- Parâmetro variádico deve ser o último



