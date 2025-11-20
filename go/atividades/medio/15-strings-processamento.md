# Atividade 15: Processamento Avançado de Strings

## Objetivo
Implementar processamento complexo de strings e parsing.

## Enunciado
Crie um programa que:
1. Crie função `extrairEmails(texto string) []string` que encontra todos os emails no texto
2. Crie função `contarPalavras(texto string) map[string]int` que conta frequência de palavras
3. Crie função `removerAcentos(texto string) string` que remove acentos (simplificado)
4. Crie função `formatarCPF(cpf string) (string, error)` que valida e formata CPF (XXX.XXX.XXX-XX)
5. Crie função `inverterPalavras(frase string) string` que inverte ordem das palavras

## Exemplo de Saída
```
Texto: "Entre em contato: joao@email.com ou maria@teste.com.br"
Emails encontrados: [joao@email.com, maria@teste.com.br]

Contagem de palavras:
  "o": 3
  "gato": 2
  "comeu": 1

CPF '12345678901': formatado: 123.456.789-01
Frase invertida: "mundo Olá" -> "Olá mundo"
```

## Dicas
- Use `strings` package: `Contains`, `Split`, `Replace`, `Trim`
- Regex para padrões complexos: `regexp` package
- Itere sobre runes para processar caracteres Unicode
- Use `strings.Builder` para construir strings eficientemente



