# Atividade 20: Métodos Básico

## Objetivo
Praticar criação de métodos para structs.

## Enunciado
Crie um programa que:
1. Defina uma struct `Retangulo` com campos `Largura` e `Altura` (float64)
2. Crie um método `Area()` que calcula e retorna a área do retângulo
3. Crie um método `Perimetro()` que calcula e retorna o perímetro
4. Crie um método `DobrarTamanho()` que dobra largura e altura (use ponteiro)
5. No `main`, crie um retângulo 5x3, imprima área e perímetro, dobre o tamanho e imprima novamente

## Exemplo de Saída
```
Retângulo: 5.0 x 3.0
Área: 15.00
Perímetro: 16.00
Após dobrar tamanho:
Retângulo: 10.0 x 6.0
Área: 60.00
Perímetro: 32.00
```

## Dicas
- Método: `func (r Receptor) Nome() tipoRetorno { }`
- Receptor por valor: `(r Tipo)`
- Receptor por ponteiro: `(r *Tipo)` (para modificar)
- Chame: `instancia.Metodo()`



