# Projeto 14: Processador de Arquivos em Lote

## 📝 Descrição
Sistema para processar múltiplos arquivos em lote com operações como renomear, mover, copiar e converter.

## 🎯 Requisitos

### Funcionalidades Obrigatórias
1. **Operações**:
   - Renomear arquivos (padrão, sequencial)
   - Mover arquivos
   - Copiar arquivos
   - Deletar arquivos
   - Listar arquivos

2. **Filtros**:
   - Por extensão
   - Por tamanho
   - Por data de modificação
   - Por nome (regex)

3. **Conversão**:
   - Converter texto (encoding)
   - Normalizar nomes (remover acentos, espaços)

4. **Batch Processing**:
   - Processar múltiplos arquivos
   - Preview antes de executar
   - Log de operações

5. **Validação**: Verificar permissões e existência

## 📚 Conceitos Utilizados
- ✅ I/O avançado (os, path/filepath)
- ✅ Goroutines (processamento paralelo)
- ✅ Channels
- ✅ Error handling
- ✅ Strings e regex
- ✅ Slices
- ✅ Concorrência

## 📁 Estrutura Sugerida
```
processador/
├── main.go
├── operacoes.go
├── filtros.go
├── batch.go
└── README.md
```

## 💡 Implementação Sugerida

### Funcionalidades
- `ListarArquivos(diretorio string) ([]FileInfo, error)`
- `FiltrarArquivos(arquivos []FileInfo, filtro Filtro) []FileInfo`
- `RenomearLote(arquivos []string, padrao string) error`
- `ProcessarParalelo(arquivos []string, operacao Operacao) error`
- `GerarPreview(operacoes []Operacao) string`

## ✅ Critérios de Sucesso
- [ ] Operações funcionam
- [ ] Filtros são precisos
- [ ] Processamento paralelo funciona
- [ ] Validações impedem erros
- [ ] Logs são úteis
- [ ] Código seguro

## 🚀 Extras (Desafio)
- [ ] Undo/redo
- [ ] Processamento assíncrono
- [ ] Progress bar
- [ ] Suporte a diretórios recursivos
- [ ] Operações customizadas via plugins

