# Atividade 02: Tipos de Dados

## Objetivo
Praticar diferentes tipos de dados em Go e suas conversões.

## Enunciado
Crie um programa que:
1. Declare uma variável `numero` do tipo `int` com valor 42
2. Declare uma variável `preco` do tipo `float64` com valor 19.99
3. Declare uma variável `ativo` do tipo `bool` com valor `true`
4. Converta `numero` para `float64` e armazene em `numeroFloat`
5. Converta `preco` para `int` (truncando) e armazene em `precoInt`
6. Imprima todos os valores e seus tipos usando `fmt.Printf` com `%T`

## Exemplo de Saída
```
numero: 42 (tipo: int)
preco: 19.99 (tipo: float64)
ativo: true (tipo: bool)
numeroFloat: 42 (tipo: float64)
precoInt: 19 (tipo: int)
```

## Dicas
- Use conversão de tipo: `Tipo(valor)`
- Use `%T` no `fmt.Printf` para imprimir o tipo
- Conversão de `float64` para `int` trunca a parte decimal



