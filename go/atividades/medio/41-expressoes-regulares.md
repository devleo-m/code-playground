# Atividade 41: Trabalhando com Expressões Regulares

## Objetivo
Dominar uso de regex para validação e extração.

## Enunciado
Crie um programa que:
1. Valide CPF usando regex: `^\d{3}\.\d{3}\.\d{3}-\d{2}$`
2. Valide email: padrão básico de email
3. Extraia todos números de texto: `"Preço: R$ 100,50 e R$ 200,00"`
4. Extraia URLs de texto
5. Substitua padrões: mascarar emails (mostrar apenas primeiros 3 caracteres)
6. Valide senha forte: mínimo 8 chars, letra maiúscula, minúscula, número, especial

## Exemplo de Saída
```
CPF "123.456.789-00": válido ✓
Email "joao@email.com": válido ✓
Números encontrados: [100.50, 200.00]
URLs: [https://example.com, http://test.com]
Email mascarado: joa***@email.com
Senha forte: true
```

## Dicas
- Use `regexp` package
- `regexp.MustCompile` para compilar padrão
- `MatchString` para validar
- `FindAllString` para extrair todas ocorrências
- `ReplaceAllString` para substituir


