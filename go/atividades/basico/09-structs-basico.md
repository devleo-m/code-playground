# Atividade 09: Structs Básico

## Objetivo
Praticar criação e uso de structs em Go.

## Enunciado
Crie um programa que:
1. Defina uma struct `Pessoa` com campos: Nome (string), Idade (int), Email (string)
2. Crie uma instância de `Pessoa` com seus dados
3. Imprima cada campo individualmente
4. Modifique a idade
5. Imprima a struct completa usando `fmt.Printf` com `%+v`

## Exemplo de Saída
```
Nome: João Silva
Idade: 25
Email: joao@email.com
Após modificar idade:
Pessoa: {Nome:João Silva Idade:26 Email:joao@email.com}
```

## Dicas
- Defina struct: `type Nome struct { Campos }`
- Crie instância: `Pessoa{Nome: "...", Idade: 25}`
- Acesse campos: `pessoa.Campo`
- Use `%+v` para imprimir struct com nomes dos campos


