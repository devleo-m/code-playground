# Atividade 30: Sistema de Validação de Formulário

## Objetivo
Criar sistema flexível de validação de formulários.

## Enunciado
Crie um programa que:
1. Defina struct `Formulario` com campos: Nome, Email, Idade, Senha, ConfirmarSenha
2. Defina interface `Validador` com método `Validar(valor interface{}) error`
3. Implemente validadores: `ValidadorObrigatorio`, `ValidadorEmail`, `ValidadorMinimo`, `ValidadorMaximo`
4. Crie função `validarFormulario(form Formulario, regras map[string][]Validador) []ErroValidacao`
5. Suporte validação customizada (ex: senha deve ser igual a ConfirmarSenha)
6. Retorne todos os erros encontrados, não apenas o primeiro

## Exemplo de Saída
```
Validando formulário:
Erros:
  - Nome: campo obrigatório
  - Email: formato inválido
  - Idade: deve ser maior que 18
  - Senha: deve ter no mínimo 8 caracteres
  - ConfirmarSenha: senhas não coincidem
```

## Dicas
- Interface permite diferentes tipos de validadores
- Map de validadores por campo
- Colete todos os erros antes de retornar
- Use type assertion para validar tipos específicos

