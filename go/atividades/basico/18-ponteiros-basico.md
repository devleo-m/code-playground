# Atividade 18: Ponteiros Básico

## Objetivo
Praticar uso básico de ponteiros em Go.

## Enunciado
Crie um programa que:
1. Declare uma variável `numero` com valor 42
2. Crie um ponteiro `ptr` que aponte para `numero`
3. Imprima o valor de `numero`, o endereço de `numero`, o valor de `ptr` (endereço) e o valor apontado por `ptr`
4. Modifique o valor através do ponteiro
5. Imprima o novo valor de `numero`

## Exemplo de Saída
```
numero: 42
endereço de numero: 0xc0000140a0
ptr (endereço): 0xc0000140a0
valor apontado por ptr: 42
Após modificar via ponteiro:
numero: 100
```

## Dicas
- Ponteiro: `var ptr *int = &variavel`
- Obter endereço: `&variavel`
- Desreferenciar: `*ponteiro`
- Modificar: `*ponteiro = novoValor`



