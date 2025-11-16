# Atividade 16: Funções - Múltiplos Retornos

## Objetivo
Praticar funções com múltiplos valores de retorno.

## Enunciado
Crie um programa que:
1. Crie uma função `dividir` que recebe dois `float64` e retorna (resultado, erro)
2. Se o divisor for 0, retorne erro; caso contrário, retorne resultado e nil
3. Crie uma função `calcularEstatisticas` que recebe um slice de `int` e retorna (soma, media, maximo, minimo)
4. No `main`, chame `dividir` e trate o erro
5. Chame `calcularEstatisticas` com [10, 20, 30, 40, 50] e imprima os resultados

## Exemplo de Saída
```
Divisão 10 / 2: 5.00 (sem erro)
Divisão 10 / 0: erro: divisão por zero
Estatísticas de [10 20 30 40 50]:
  Soma: 150
  Média: 30.00
  Máximo: 50
  Mínimo: 10
```

## Dicas
- Múltiplos retornos: `func nome() (tipo1, tipo2) { return val1, val2 }`
- Padrão Go: `(resultado, error)`
- Verifique erro: `if err != nil { }`
- Ignore valor: `_, err := funcao()`

