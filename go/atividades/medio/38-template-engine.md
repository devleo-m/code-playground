# Atividade 38: Template Engine Simples

## Objetivo
Criar engine de templates básico.

## Enunciado
Crie um programa que:
1. Defina formato de template: `"Olá {{nome}}, você tem {{idade}} anos"`
2. Crie função `ProcessarTemplate(template string, dados map[string]interface{}) (string, error)`
3. Suporte variáveis: `{{variavel}}`
4. Suporte condicionais: `{{if condicao}}...{{else}}...{{endif}}`
5. Suporte loops: `{{for item in lista}}...{{endfor}}`
6. Escape caracteres especiais
7. Valide template (chaves balanceadas)

## Exemplo
```
Template: "Olá {{nome}}, você tem {{idade}} anos"
Dados: {"nome": "João", "idade": 30}
Resultado: "Olá João, você tem 30 anos"
```

## Dicas
- Use `strings.Replace` ou regex para substituir variáveis
- Parse condicionais e loops recursivamente
- Valide sintaxe antes de processar
- Escape: substitua `{{` por literal se necessário

