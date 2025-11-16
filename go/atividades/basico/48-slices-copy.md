# Atividade 48: Copiar e Comparar Slices

## Objetivo
Praticar cópia de slices e comparação.

## Enunciado
Crie um programa que:
1. Crie slice original [1, 2, 3, 4, 5]
2. Crie cópia usando `copy` (crie slice destino com mesmo tamanho)
3. Modifique cópia e verifique que original não mudou
4. Crie função `slicesIguais` que compara dois slices (elemento por elemento)
5. Teste com slices iguais e diferentes

## Exemplo de Saída
```
Original: [1 2 3 4 5]
Cópia: [1 2 3 4 5]
Após modificar cópia:
Original: [1 2 3 4 5] (não mudou)
Cópia: [10 20 30 40 50]
Slices [1 2 3] e [1 2 3]: iguais
Slices [1 2 3] e [1 2 4]: diferentes
```

## Dicas
- `copy(destino, origem)` copia elementos
- Destino deve ter capacidade suficiente
- Retorna número de elementos copiados
- Comparação: itere e compare elemento por elemento (não use `==` diretamente)

