# Atividade 05: Sistema de Validação Complexo

## Objetivo
Criar sistema de validação com múltiplas regras e erros customizados.

## Enunciado
Crie um programa que:
1. Defina struct `ErroValidacao` com Campo, Mensagem, Valor
2. Defina struct `Usuario` com Nome, Email, Idade, Senha
3. Crie função `validarUsuario(u Usuario) []ErroValidacao` que retorna slice de erros:
   - Nome: não vazio, mínimo 3 caracteres
   - Email: formato válido (contém @ e .)
   - Idade: entre 18 e 120
   - Senha: mínimo 8 caracteres, contém letra e número
4. Crie função `imprimirErros(erros []ErroValidacao)` que formata erros
5. Teste com usuários válidos e inválidos

## Exemplo de Saída
```
Validando usuário:
Erros encontrados:
  - Campo 'Nome': muito curto (mínimo 3 caracteres) - Valor: 'Jo'
  - Campo 'Email': formato inválido - Valor: 'joaoemail.com'
  - Campo 'Senha': deve conter letra e número - Valor: '********'

Usuário válido!
```

## Dicas
- Retorne slice de erros para múltiplas validações
- Use `strings.Contains`, `len()` para validações
- Crie erros descritivos com contexto
- Itere sobre erros para imprimir todos

