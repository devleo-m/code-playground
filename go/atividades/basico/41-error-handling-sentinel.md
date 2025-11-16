# Atividade 41: Sentinel Errors

## Objetivo
Praticar uso de sentinel errors (erros pré-definidos).

## Enunciado
Crie um programa que:
1. Defina sentinel errors como variáveis de pacote:
   - `ErrUsuarioNaoEncontrado`
   - `ErrSenhaInvalida`
2. Crie função `autenticar` que recebe usuario e senha e retorna error
3. Retorne sentinel errors apropriados baseado na validação
4. No `main`, teste diferentes cenários e verifique erros usando `errors.Is`

## Exemplo de Saída
```
Autenticando 'joao' com senha '123':
  Erro: usuário não encontrado
Autenticando 'admin' com senha 'wrong':
  Erro: senha inválida
Autenticando 'admin' com senha 'admin123':
  Sucesso!
```

## Dicas
- Sentinel error: `var ErrNome = errors.New("mensagem")`
- Retorne: `return ErrUsuarioNaoEncontrado`
- Verifique: `if errors.Is(err, ErrUsuarioNaoEncontrado) { }`
- Use prefixo do pacote na mensagem

