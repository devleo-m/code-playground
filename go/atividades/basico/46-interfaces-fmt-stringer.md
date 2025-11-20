# Atividade 46: Interface fmt.Stringer

## Objetivo
Praticar implementação da interface `fmt.Stringer`.

## Enunciado
Crie um programa que:
1. Defina struct `Pessoa` com Nome (string) e Idade (int)
2. Implemente método `String() string` para `Pessoa` (satisfaz `fmt.Stringer`)
3. Defina struct `Produto` com Nome (string) e Preco (float64)
4. Implemente `String() string` para `Produto`
5. No `main`, crie instâncias e use `fmt.Println` (chama `String()` automaticamente)

## Exemplo de Saída
```
Pessoa: João Silva (25 anos)
Produto: Notebook - R$ 2500.00
```

## Dicas
- `fmt.Stringer`: interface com método `String() string`
- `fmt.Println` chama `String()` automaticamente
- Use `fmt.Sprintf` para formatar string
- Implemente para tipos customizados que você quer imprimir



