# Atividade 03: Valores Zero

## Objetivo
Entender os valores zero de diferentes tipos em Go.

## Enunciado
Crie um programa que declare variáveis sem inicializar e imprima seus valores zero:
1. `var numero int`
2. `var texto string`
3. `var decimal float64`
4. `var logico bool`
5. `var lista []int` (slice)
6. `var mapa map[string]int` (map)

## Exemplo de Saída
```
numero: 0
texto: '' (string vazia)
decimal: 0
logico: false
lista: [] (nil)
mapa: map[] (nil)
```

## Dicas
- Valores zero são atribuídos automaticamente
- String vazia é `""`
- Slices e maps nil são `nil`
- Use `== nil` para verificar se slice/map é nil

