# Atividade 40: Compressão Simples (RLE)

## Objetivo
Implementar algoritmo de compressão Run-Length Encoding.

## Enunciado
Crie um programa que:
1. Implemente `comprimirRLE(texto string) string` que comprime sequências repetidas
   - "AAAABBBCC" -> "4A3B2C"
2. Implemente `descomprimirRLE(comprimido string) (string, error)` que descomprime
3. Calcule taxa de compressão
4. Teste com diferentes tipos de dados (texto, números)
5. Implemente compressão para dados binários (slice de bytes)

## Exemplo de Saída
```
Texto original: "AAAABBBCCDDD" (12 bytes)
Comprimido: "4A3B2C3D" (8 bytes)
Taxa de compressão: 33.33%
Descomprimido: "AAAABBBCCDDD" ✓
```

## Dicas
- RLE: conta sequências consecutivas
- Itere sobre string contando repetições
- Formato: número seguido de caractere
- Parse números ao descomprimir
- Cuidado com números no texto original

