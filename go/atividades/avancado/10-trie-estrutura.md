# Atividade 10: Estrutura de Dados Trie

## Objetivo
Implementar Trie (prefix tree) otimizada para busca de strings.

## Enunciado
Crie uma Trie completa:
1. Defina struct `NoTrie` com filhos (map[rune]*NoTrie), fimPalavra (bool), valor (interface{})
2. Implemente métodos:
   - `Inserir(palavra string, valor interface{})`
   - `Buscar(palavra string) (interface{}, bool)`
   - `Remover(palavra string) bool`
   - `BuscarPrefixo(prefixo string) []string` - retorna todas palavras com prefixo
   - `Autocompletar(prefixo string, limite int) []string`
   - `ContarPalavras() int`
3. Implemente compressão (Radix Tree) para reduzir memória
4. Adicione suporte a wildcards na busca
5. Implemente busca fuzzy (tolerância a erros)
6. Teste com dicionário de 10.000 palavras

## Exemplo de Saída
```
Trie criada
Inserindo 10000 palavras...
Memória: 2.5MB
Buscar "casa": encontrado (valor: "residência")
BuscarPrefixo "car": [carro, carta, carta] (3 palavras)
Autocompletar "prog": [programa, programar, programação]
Busca fuzzy "caza" (distância 1): encontrado "casa"
```

## Dicas
- Cada nó representa caractere
- Caminho da raiz forma palavra
- Marque fim de palavra
- Compressão: mescle nós com único filho
- Fuzzy: use Levenshtein distance

## Desafio Extra
Implemente suporte a Unicode completo e busca com regex.


