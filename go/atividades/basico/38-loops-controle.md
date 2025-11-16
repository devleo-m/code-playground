# Atividade 38: Controle de Loops

## Objetivo
Praticar controle avançado de loops com break e continue.

## Enunciado
Crie um programa que:
1. Use loop para encontrar primeiro número divisível por 7 e 11 entre 1-100 (use `break`)
2. Use loop para imprimir números de 1-20, pulando múltiplos de 3 (use `continue`)
3. Use loop aninhado com label para sair do loop externo quando encontrar número específico

## Exemplo de Saída
```
Primeiro número divisível por 7 e 11: 77
Números de 1-20 (sem múltiplos de 3):
1 2 4 5 7 8 10 11 13 14 16 17 19 20
Encontrado 15 no loop aninhado!
```

## Dicas
- `break` sai do loop atual
- `continue` pula para próxima iteração
- Label: `loopExterno: for { for { break loopExterno } }`
- Use labels para sair de loops aninhados

