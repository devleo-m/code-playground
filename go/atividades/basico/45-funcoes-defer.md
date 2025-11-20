# Atividade 45: Funções com defer

## Objetivo
Praticar uso de `defer` para garantir execução de código.

## Enunciado
Crie um programa que:
1. Crie função `processarArquivo` que simula abertura de arquivo
2. Use `defer` para garantir "fechamento" (imprimir "Arquivo fechado")
3. Crie função `contador` que usa `defer` para imprimir contador ao final
4. Demonstre que `defer` executa na ordem inversa (LIFO)
5. No `main`, chame as funções e observe a ordem de execução

## Exemplo de Saída
```
Abrindo arquivo...
Processando arquivo...
Arquivo fechado (defer)
Contador: 0
Contador: 1
Contador: 2
Final: 3
Defer 3
Defer 2
Defer 1
```

## Dicas
- `defer` executa ao final da função (antes do return)
- Ordem: último defer primeiro (LIFO - Last In First Out)
- Útil para cleanup, fechar recursos
- `defer` avalia argumentos imediatamente, executa depois



