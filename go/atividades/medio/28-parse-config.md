# Atividade 28: Parser de Configuração

## Objetivo
Criar parser de arquivo de configuração simples.

## Enunciado
Crie um programa que:
1. Defina formato de config simples: `chave=valor` (linha por linha)
2. Crie função `lerConfig(caminho string) (map[string]string, error)`
3. Suporte comentários (linhas começando com `#`)
4. Suporte seções: `[secao]` agrupa chaves
5. Retorne map aninhado: `map[string]map[string]string`
6. Crie função `obterConfig(config map[string]map[string]string, secao, chave string) (string, error)`

## Exemplo de Arquivo
```
# Configuração do sistema
[database]
host=localhost
port=5432
user=admin

[app]
name=MinhaApp
debug=true
```

## Dicas
- Leia linha por linha com `bufio.Scanner`
- Use `strings.Split` para separar chave=valor
- Detecte seções com `strings.HasPrefix(line, "[")`
- Trate espaços com `strings.TrimSpace`



