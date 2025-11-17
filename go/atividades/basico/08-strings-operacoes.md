# Atividade 08: Operações com Strings

## Objetivo
Praticar operações comuns com strings em Go.

## Enunciado
Crie um programa que:
1. Declare uma string `texto` com valor "Aprendendo Go"
2. Imprima o tamanho da string
3. Imprima o primeiro caractere (rune)
4. Verifique se a string contém "Go"
5. Substitua "Go" por "Golang" (use `strings.Replace`)
6. Converta a string para maiúsculas
7. Divida a string em palavras (use `strings.Split`)

## Exemplo de Saída
```
Texto: Aprendendo Go
Tamanho: 14 caracteres
Primeiro caractere: A
Contém 'Go'? true
Texto substituído: Aprendendo Golang
Maiúsculas: APRENDENDO GO
Palavras: [Aprendendo Go]
```

## Dicas
- Use `len()` para tamanho (em bytes, não runes)
- Use `strings.Contains()` para verificar substring
- Use `strings.Replace()` para substituir
- Use `strings.ToUpper()` para maiúsculas
- Use `strings.Split()` para dividir


