# Atividade 22: Testes Unitários

## Objetivo
Escrever testes unitários usando `testing` package.

## Enunciado
Crie um programa com:
1. Função `somar(a, b int) int` e teste unitário `TestSomar`
2. Função `dividir(a, b float64) (float64, error)` e teste que verifica erro de divisão por zero
3. Use tabela de testes para testar múltiplos casos de `validarEmail(email string) bool`
4. Crie testes com `t.Run` para agrupar testes relacionados
5. Use `t.Helper()` em funções auxiliares de teste
6. Implemente benchmark para função `ordenarSlice`

## Exemplo
```go
func TestSomar(t *testing.T) {
    resultado := somar(2, 3)
    if resultado != 5 {
        t.Errorf("Esperado 5, obtido %d", resultado)
    }
}
```

## Dicas
- Arquivo de teste: `*_test.go`
- Função de teste: `func TestNome(t *testing.T)`
- Tabela de testes: slice de structs com casos
- Benchmark: `func BenchmarkNome(b *testing.B)`

