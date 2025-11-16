# Atividade 39: Métodos com Receptor Ponteiro

## Objetivo
Praticar diferença entre receptor por valor e por ponteiro.

## Enunciado
Crie um programa que:
1. Defina struct `ContaBancaria` com Saldo (float64)
2. Crie método `Depositar` com receptor por ponteiro que adiciona valor
3. Crie método `Sacar` com receptor por ponteiro que subtrai valor (verifique saldo)
4. Crie método `ObterSaldo` com receptor por valor que retorna saldo
5. No `main`, crie conta, deposite 100, saque 30, imprima saldo

## Exemplo de Saída
```
Saldo inicial: R$ 0.00
Após depositar R$ 100.00: R$ 100.00
Após sacar R$ 30.00: R$ 70.00
```

## Dicas
- Receptor ponteiro: `func (c *Conta) Metodo() { }` (modifica)
- Receptor valor: `func (c Conta) Metodo() tipo { }` (não modifica)
- Use ponteiro quando método modifica struct
- Acesse campo: `c.Saldo` (Go desreferencia automaticamente)

