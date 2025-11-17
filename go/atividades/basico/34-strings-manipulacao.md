# Atividade 34: Manipulação de Strings

## Objetivo
Praticar manipulação avançada de strings.

## Enunciado
Crie um programa que:
1. Declare string "  Aprendendo Go Programming  "
2. Remova espaços do início e fim (use `strings.TrimSpace`)
3. Converta para minúsculas
4. Verifique se começa com "aprendendo"
5. Verifique se termina com "programming"
6. Conte quantas vezes "go" aparece (case-insensitive)
7. Substitua "go" por "Golang"

## Exemplo de Saída
```
Original: '  Aprendendo Go Programming  '
Após trim: 'Aprendendo Go Programming'
Minúsculas: 'aprendendo go programming'
Começa com 'aprendendo'? true
Termina com 'programming'? true
'go' aparece 1 vez
Após substituir: 'aprendendo Golang programming'
```

## Dicas
- `strings.TrimSpace(s)` remove espaços
- `strings.ToLower(s)` minúsculas
- `strings.HasPrefix(s, prefix)` verifica início
- `strings.HasSuffix(s, suffix)` verifica fim
- `strings.Count(s, substr)` conta ocorrências
- `strings.Replace(s, old, new, n)` substitui


