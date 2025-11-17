# Atividade 09: Error Handling Avançado

## Objetivo
Implementar tratamento de erros robusto com wrapping e sentinel errors.

## Enunciado
Crie um programa que simula sistema de arquivos:
1. Defina sentinel errors: `ErrArquivoNaoEncontrado`, `ErrPermissaoNegada`, `ErrDiscoCheio`
2. Crie função `lerArquivo(caminho string) ([]byte, error)` que:
   - Retorna `ErrArquivoNaoEncontrado` se arquivo não existe
   - Retorna `ErrPermissaoNegada` se sem permissão
   - Adiciona contexto com `fmt.Errorf` e `%w`
3. Crie função `processarArquivo(caminho string) error` que chama `lerArquivo` e adiciona mais contexto
4. Crie função `salvarArquivo(caminho string, dados []byte) error` com tratamento similar
5. No `main`, teste diferentes cenários e use `errors.Is` e `errors.As` para verificar erros específicos

## Exemplo de Saída
```
Processando arquivo 'dados.txt':
  Erro: falha ao processar arquivo 'dados.txt': erro ao ler arquivo: arquivo não encontrado

Verificando tipo de erro:
  É ErrArquivoNaoEncontrado? true
  Ação: criar arquivo novo
```

## Dicas
- Use `fmt.Errorf("contexto: %w", err)` para wrapping
- `errors.Is(err, ErrEspecifico)` verifica na cadeia de erros
- `errors.As(err, &var)` extrai tipo específico de erro
- Adicione contexto em cada camada da aplicação


