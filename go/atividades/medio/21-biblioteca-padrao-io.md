# Atividade 21: Trabalhando com I/O

## Objetivo
Dominar leitura e escrita de arquivos usando `io` e `os` packages.

## Enunciado
Crie um programa que:
1. Crie função `lerArquivo(caminho string) (string, error)` que lê arquivo completo
2. Crie função `escreverArquivo(caminho string, conteudo string) error` que escreve arquivo
3. Crie função `copiarArquivo(origem, destino string) error` usando `io.Copy`
4. Crie função `contarLinhas(caminho string) (int, error)` que conta linhas do arquivo
5. Implemente leitura linha por linha usando `bufio.Scanner`
6. Crie função `buscarTexto(caminho, texto string) ([]int, error)` que retorna números das linhas onde texto aparece

## Exemplo de Saída
```
Arquivo lido: 150 bytes
Linhas no arquivo: 10
Texto 'erro' encontrado nas linhas: [3, 7, 12]
Arquivo copiado com sucesso
```

## Dicas
- `os.Open` para ler, `os.Create` para escrever
- `io.Copy` para copiar dados
- `bufio.Scanner` para leitura linha por linha
- Sempre feche arquivos com `defer file.Close()`

