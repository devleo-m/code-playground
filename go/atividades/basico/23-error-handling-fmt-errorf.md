# Atividade 23: Error Handling - fmt.Errorf

## Objetivo
Praticar criação de erros formatados com contexto.

## Enunciado
Crie um programa que:
1. Crie uma função `validarIdade` que recebe `int` e retorna `error`
2. Se idade < 0, retorne erro formatado: "idade inválida: %d (deve ser positiva)"
3. Se idade > 150, retorne erro formatado: "idade inválida: %d (muito alta)"
4. Caso contrário, retorne `nil`
5. No `main`, teste com idades: -5, 25, 200

## Exemplo de Saída
```
Idade -5: erro: idade inválida: -5 (deve ser positiva)
Idade 25: válida
Idade 200: erro: idade inválida: 200 (muito alta)
```

## Dicas
- `fmt.Errorf("formato %d", valor)` cria erro formatado
- Use verbos: `%d` (int), `%s` (string), `%v` (valor padrão)
- Sempre verifique erros antes de usar resultados


