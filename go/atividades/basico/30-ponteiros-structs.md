# Atividade 30: Ponteiros e Structs

## Objetivo
Praticar uso de ponteiros com structs.

## Enunciado
Crie um programa que:
1. Defina struct `Contador` com campo `valor` (int)
2. Crie método `Incrementar()` com receptor por ponteiro que incrementa valor
3. Crie método `ObterValor()` com receptor por valor que retorna o valor
4. Crie função `Resetar` que recebe ponteiro para `Contador` e zera o valor
5. No `main`, crie contador, incremente 3 vezes, imprima valor, resete e imprima novamente

## Exemplo de Saída
```
Valor inicial: 0
Após 3 incrementos: 3
Após resetar: 0
```

## Dicas
- Receptor por ponteiro: `func (c *Contador) Metodo() { }`
- Modificar struct: use ponteiro
- Acessar campo: `c.valor` (Go desreferencia automaticamente)
- Passar ponteiro: `funcao(&instancia)`



