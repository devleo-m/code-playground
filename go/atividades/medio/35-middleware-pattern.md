# Atividade 35: Padrão Middleware

## Objetivo
Implementar padrão middleware usando funções de alta ordem.

## Enunciado
Crie um programa que:
1. Defina tipo `Handler func(string) string` (processa requisição)
2. Defina tipo `Middleware func(Handler) Handler` (envolve handler)
3. Implemente middlewares:
   - `LoggingMiddleware` - loga antes e depois
   - `AutenticacaoMiddleware` - verifica token
   - `TempoMiddleware` - mede tempo de execução
4. Crie função `AplicarMiddlewares(handler Handler, middlewares ...Middleware) Handler`
5. Demonstre encadeamento de middlewares
6. Simule processamento de requisição HTTP simples

## Exemplo de Saída
```
Aplicando middlewares: Logging -> Autenticação -> Tempo
[LOG] Iniciando requisição: /api/users
[LOG] Token válido
[LOG] Processando...
[LOG] Tempo: 15ms
[LOG] Requisição concluída: resposta
```

## Dicas
- Middleware recebe handler e retorna novo handler
- Encadeie: `m1(m2(m3(handler)))`
- Cada middleware pode modificar requisição/resposta
- Útil para cross-cutting concerns

