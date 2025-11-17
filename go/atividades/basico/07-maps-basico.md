# Atividade 07: Maps Básico

## Objetivo
Praticar criação e manipulação básica de maps.

## Enunciado
Crie um programa que:
1. Crie um map de string para int representando idades de pessoas
2. Adicione as seguintes entradas:
   - "João": 25
   - "Maria": 30
   - "Pedro": 22
3. Imprima o map completo
4. Acesse e imprima a idade de "Maria"
5. Adicione uma nova entrada "Ana": 28
6. Verifique se "Carlos" existe no map
7. Imprima o map final

## Exemplo de Saída
```
Map completo: map[João:25 Maria:30 Pedro:22]
Idade de Maria: 30
Carlos existe? false
Map final: map[Ana:28 João:25 Maria:30 Pedro:22]
```

## Dicas
- Map: `make(map[tipoChave]tipoValor)` ou `map[tipoChave]tipoValor{}`
- Adicionar: `map[chave] = valor`
- Verificar existência: `valor, ok := map[chave]`


