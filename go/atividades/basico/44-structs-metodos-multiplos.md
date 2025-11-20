# Atividade 44: Structs com Múltiplos Métodos

## Objetivo
Praticar criação de structs com vários métodos relacionados.

## Enunciado
Crie um programa que:
1. Defina struct `ContaCorrente` com Saldo (float64) e Limite (float64)
2. Crie métodos:
   - `Depositar(valor float64)` - adiciona ao saldo
   - `Sacar(valor float64) error` - subtrai, retorna erro se sem saldo+limite
   - `Transferir(destino *ContaCorrente, valor float64) error` - transfere entre contas
   - `ObterSaldo() float64` - retorna saldo
3. No `main`, crie duas contas, deposite, transfira e teste saque além do limite

## Exemplo de Saída
```
Conta 1 - Saldo: R$ 1000.00
Conta 2 - Saldo: R$ 500.00
Após transferir R$ 300.00:
Conta 1 - Saldo: R$ 700.00
Conta 2 - Saldo: R$ 800.00
Tentativa de saque além do limite: erro
```

## Dicas
- Múltiplos métodos para mesma struct
- Use ponteiro para modificar: `func (c *Conta) Metodo() { }`
- Verifique condições antes de modificar
- Retorne erro quando operação não for possível



