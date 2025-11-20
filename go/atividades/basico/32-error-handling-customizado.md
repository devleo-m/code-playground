# Atividade 32: Erros Customizados

## Objetivo
Praticar criação de tipos de erro customizados.

## Enunciado
Crie um programa que:
1. Defina struct `ErroValidacao` com campos: Campo (string), Mensagem (string)
2. Implemente método `Error() string` para `ErroValidacao`
3. Crie função `validarEmail` que recebe string e retorna `error`
4. Se email vazio, retorne `ErroValidacao{Campo: "email", Mensagem: "não pode ser vazio"}`
5. Se não contém "@", retorne erro similar
6. No `main`, teste com emails válido e inválidos

## Exemplo de Saída
```
Email '': erro no campo 'email': não pode ser vazio
Email 'joao': erro no campo 'email': deve conter @
Email 'joao@email.com': válido
```

## Dicas
- Erro customizado: struct com método `Error() string`
- Retorne erro: `return ErroValidacao{...}`
- Verifique: `if err != nil { fmt.Println(err) }`
- `err.Error()` retorna string do erro



