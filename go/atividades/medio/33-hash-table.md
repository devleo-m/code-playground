# Atividade 33: Tabela Hash Simples

## Objetivo
Implementar tabela hash básica.

## Enunciado
Crie um programa que:
1. Defina struct `HashTable` com slice de buckets (cada bucket é slice de pares chave-valor)
2. Implemente função hash simples: `hash(chave string, tamanho int) int`
3. Implemente métodos:
   - `Set(chave string, valor interface{})`
   - `Get(chave string) (interface{}, bool)`
   - `Delete(chave string) bool`
   - `Size() int`
4. Trate colisões usando encadeamento (cada bucket é slice)
5. Implemente redimensionamento quando carga > 0.75
6. Teste com diferentes tamanhos e chaves

## Exemplo de Saída
```
HashTable criada (tamanho: 10)
Set("chave1", "valor1")
Set("chave2", "valor2")
Get("chave1"): valor1 (true)
Tamanho: 2
Carga: 0.20
Redimensionando para 20 buckets...
```

## Dicas
- Função hash: some códigos ASCII e use módulo
- Colisões: mesmo hash, diferentes chaves
- Encadeamento: bucket contém slice de pares
- Redimensionar: recalculando hash de todos elementos

