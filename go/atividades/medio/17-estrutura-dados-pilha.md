# Atividade 17: Estrutura de Dados - Pilha

## Objetivo
Implementar estrutura de dados pilha (stack) usando slices.

## Enunciado
Crie um programa que:
1. Defina struct `Pilha` com slice genérico (use `[]interface{}` ou generics se Go 1.18+)
2. Implemente métodos:
   - `Push(valor interface{})` - adiciona no topo
   - `Pop() (interface{}, error)` - remove do topo
   - `Top() (interface{}, error)` - retorna topo sem remover
   - `IsEmpty() bool` - verifica se vazia
   - `Size() int` - retorna tamanho
3. Use pilha para verificar parênteses balanceados: `verificarBalanceamento(expr string) bool`
4. Use pilha para converter expressão infixa para pós-fixa (notação polonesa reversa)
5. Teste com diferentes expressões

## Exemplo de Saída
```
Pilha:
  Push(1), Push(2), Push(3)
  Top: 3
  Pop: 3
  Size: 2

Verificando "(a+b)*(c-d)": balanceado ✓
Verificando "(a+b*(c-d)": não balanceado ✗

Infixa: "a+b*c" -> Pós-fixa: "abc*+"
```

## Dicas
- Pilha: LIFO (Last In First Out)
- Use slice: `append` para push, `slice[len-1]` para top, `slice[:len-1]` para pop
- Parênteses: push '(' e pop quando encontrar ')'
- Pós-fixa: use pilha para operadores



