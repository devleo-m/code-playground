# Atividade 22: Error Handling Básico

## Objetivo
Praticar tratamento básico de erros em Go.

## Enunciado
Crie um programa que:
1. Crie uma função `dividir` que recebe dois `float64` e retorna `(float64, error)`
2. Se divisor for 0, retorne `0` e um erro criado com `errors.New("divisão por zero")`
3. Caso contrário, retorne resultado e `nil`
4. No `main`, chame `dividir` com (10, 2) e (10, 0)
5. Verifique erros e imprima mensagens apropriadas

## Exemplo de Saída
```
10 / 2 = 5.00 (sucesso)
10 / 0 = erro: divisão por zero
```

## Dicas
- Padrão Go: `(resultado, error)`
- Criar erro: `errors.New("mensagem")`
- Verificar: `if err != nil { }`
- Sucesso: retorne `nil` como erro

