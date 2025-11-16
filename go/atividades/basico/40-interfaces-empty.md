# Atividade 40: Empty Interface

## Objetivo
Praticar uso básico de empty interface (interface{}).

## Enunciado
Crie um programa que:
1. Crie função `imprimirTipo` que recebe `interface{}` e imprime tipo e valor
2. Use type assertion para verificar se é int, string ou bool
3. No `main`, chame com diferentes tipos: int, string, bool, float64
4. Crie função `processar` que aceita `interface{}` e:
   - Se for int, multiplica por 2
   - Se for string, converte para maiúsculas
   - Caso contrário, imprime "tipo não suportado"

## Exemplo de Saída
```
Tipo: int, Valor: 42
Tipo: string, Valor: olá
Tipo: bool, Valor: true
Processando 21: 42
Processando 'mundo': MUNDO
Processando true: tipo não suportado
```

## Dicas
- `interface{}` aceita qualquer tipo
- Type assertion: `valor, ok := i.(tipo)`
- Use `%T` para imprimir tipo
- Prefira generics (Go 1.18+) quando possível

