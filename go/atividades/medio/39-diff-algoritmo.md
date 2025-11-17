# Atividade 39: Algoritmo de Diff

## Objetivo
Implementar algoritmo simples de diferença entre textos.

## Enunciado
Crie um programa que:
1. Crie função `calcularDiff(texto1, texto2 string) []DiffItem`
2. Defina struct `DiffItem` com Tipo (Adicionado, Removido, Igual), Texto, Linha
3. Compare linha por linha
4. Identifique linhas adicionadas, removidas e iguais
5. Formate saída colorida (simulado com tags)
6. Calcule similaridade percentual

## Exemplo de Saída
```
Texto 1: 3 linhas
Texto 2: 4 linhas
Diff:
  [IGUAL] Linha 1
  [REMOVIDO] Linha 2 original
  [ADICIONADO] Linha 2 nova
  [IGUAL] Linha 3
Similaridade: 66.67%
```

## Dicas
- Divida textos em linhas: `strings.Split(texto, "\n")`
- Compare linha por linha
- Use algoritmo simples ou LCS (Longest Common Subsequence)
- Marque diferenças durante comparação


