# Atividade 08: Polimorfismo com Interfaces

## Objetivo
Implementar polimorfismo usando interfaces para diferentes tipos.

## Enunciado
Crie um programa que:
1. Defina interface `Calculavel` com método `Calcular() float64`
2. Defina structs `Retangulo`, `Circulo`, `Triangulo` que implementam `Calculavel` (calcular área)
3. Defina interface `Comparavel` com método `Comparar(outro Comparavel) int` (retorna -1, 0, 1)
4. Implemente `Comparavel` para as formas (compare por área)
5. Crie função `ordenarFormas(formas []Calculavel)` que ordena por área
6. Crie função `encontrarMaiorArea(formas []Calculavel) Calculavel`

## Exemplo de Saída
```
Formas:
  Retângulo 5x3: área 15.00
  Círculo r=4: área 50.27
  Triângulo 6x4: área 12.00

Ordenadas por área:
  1. Triângulo: 12.00
  2. Retângulo: 15.00
  3. Círculo: 50.27

Maior área: Círculo (50.27)
```

## Dicas
- Interface permite tratar tipos diferentes uniformemente
- Type assertion para acessar campos específicos: `ret, ok := forma.(Retangulo)`
- Use `sort.Slice` para ordenar ou implemente algoritmo próprio
- Polimorfismo: mesmo código funciona com diferentes tipos



