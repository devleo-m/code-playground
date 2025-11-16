# Atividade 21: Interfaces Básico

## Objetivo
Praticar criação e uso de interfaces básicas.

## Enunciado
Crie um programa que:
1. Defina uma interface `Forma` com método `Area() float64`
2. Defina structs `Circulo` (com Raio) e `Quadrado` (com Lado)
3. Implemente `Area()` para ambos (implicitamente satisfazem `Forma`)
4. Crie uma função `imprimirArea` que recebe uma `Forma` e imprime a área
5. No `main`, crie um círculo (raio=5) e um quadrado (lado=4) e chame `imprimirArea` para ambos

## Exemplo de Saída
```
Área do círculo (raio=5): 78.54
Área do quadrado (lado=4): 16.00
```

## Dicas
- Interface: `type Nome interface { Metodos() }`
- Tipos satisfazem interfaces automaticamente ao implementar métodos
- Use `math.Pi` para π
- Função pode receber interface como parâmetro

