# Atividade 10: Structs Aninhadas

## Objetivo
Praticar structs aninhadas e acesso a campos.

## Enunciado
Crie um programa que:
1. Defina uma struct `Endereco` com campos: Rua, Cidade, CEP (todos string)
2. Defina uma struct `Pessoa` com campos: Nome (string), Idade (int), Endereco (Endereco)
3. Crie uma instância de `Pessoa` com endereço completo
4. Acesse e imprima o endereço completo da pessoa
5. Modifique a cidade do endereço
6. Imprima a pessoa completa

## Exemplo de Saída
```
Pessoa: João Silva, 25 anos
Endereco: Rua das Flores, 123 - São Paulo - 01234-567
Cidade modificada: Rio de Janeiro
Pessoa completa: {Nome:João Silva Idade:25 Endereco:{Rua:Rua das Flores Cidade:Rio de Janeiro CEP:01234-567}}
```

## Dicas
- Struct aninhada: campo do tipo de outra struct
- Acesse campos aninhados: `pessoa.Endereco.Cidade`
- Modifique campos aninhados da mesma forma


