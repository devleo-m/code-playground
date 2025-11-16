# Atividade 36: Funções Anônimas

## Objetivo
Praticar funções anônimas e closures básicas.

## Enunciado
Crie um programa que:
1. Atribua uma função anônima a variável `somar` que recebe dois int e retorna soma
2. Atribua função anônima a `multiplicar` que recebe dois int e retorna produto
3. Crie função `aplicarOperacao` que recebe dois int e uma função `func(int, int) int`
4. No `main`, chame `aplicarOperacao` com soma e multiplicação
5. Crie closure `contador` que retorna função que incrementa e retorna contador interno

## Exemplo de Saída
```
10 + 5 = 15
10 * 5 = 50
Contador: 1
Contador: 2
Contador: 3
```

## Dicas
- Função anônima: `var nome = func(params) tipoRetorno { }`
- Closure: função que captura variáveis do escopo externo
- Passar função: `funcao(operacao func(int, int) int)`
- Chamar: `operacao(a, b)`

