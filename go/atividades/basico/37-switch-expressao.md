# Atividade 37: Switch com Expressão

## Objetivo
Praticar diferentes formas de usar switch.

## Enunciado
Crie um programa que:
1. Use `switch` com expressão para classificar nota (0-10):
   - 9-10: "Excelente"
   - 7-8: "Bom"
   - 5-6: "Regular"
   - 0-4: "Ruim"
2. Use `switch` sem expressão para verificar múltiplas condições de idade
3. Use `switch` com type assertion para processar diferentes tipos (int, string, bool)

## Exemplo de Saída
```
Nota 8.5: Bom
Idade 25: Adulto
Tipo int: 42
Tipo string: Olá
Tipo bool: true
```

## Dicas
- `switch variavel { case valor1: ... }`
- `switch { case condicao1: ... }` (sem expressão)
- `switch v := valor.(type) { case int: ... }` (type switch)
- Use `fallthrough` para continuar no próximo case

