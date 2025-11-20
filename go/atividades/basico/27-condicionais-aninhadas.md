# Atividade 27: Condicionais Aninhadas

## Objetivo
Praticar condicionais complexas e aninhadas.

## Enunciado
Crie um programa que simule um sistema de desconto:
1. Declare variáveis: `preco` (float64) e `ehClienteVIP` (bool)
2. Aplique descontos:
   - Se preço >= 100 e é VIP: desconto 20%
   - Se preço >= 100 e não é VIP: desconto 10%
   - Se preço >= 50 e é VIP: desconto 15%
   - Se preço >= 50 e não é VIP: desconto 5%
   - Caso contrário: sem desconto
3. Calcule e imprima preço original, desconto aplicado e preço final

## Exemplo de Saída (preco=120, VIP=true)
```
Preço original: R$ 120.00
Cliente VIP: Sim
Desconto aplicado: 20%
Preço final: R$ 96.00
```

## Dicas
- Use `if-else if-else` para múltiplas condições
- Condições podem usar `&&` (E) e `||` (OU)
- Calcule desconto: `preco * (desconto / 100)`
- Use `fmt.Printf` com `%.2f` para formatar moeda



