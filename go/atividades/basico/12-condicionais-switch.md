# Atividade 12: Condicionais - switch

## Objetivo
Praticar estruturas condicionais com `switch`.

## Enunciado
Crie um programa que:
1. Declare uma variável `diaSemana` com valor de 1 a 7
2. Use `switch` para imprimir o nome do dia:
   - 1: "Domingo"
   - 2: "Segunda-feira"
   - 3: "Terça-feira"
   - 4: "Quarta-feira"
   - 5: "Quinta-feira"
   - 6: "Sexta-feira"
   - 7: "Sábado"
   - default: "Dia inválido"
3. Use `switch` sem expressão para verificar se um número é par ou ímpar

## Exemplo de Saída (diaSemana = 3)
```
Dia da semana: Terça-feira
```

## Dicas
- `switch variavel { case valor1: ... case valor2: ... default: ... }`
- `switch { case condicao1: ... }` (sem expressão)
- Use `break` não é necessário (automático)
- Use `fallthrough` se quiser continuar para próximo case


