# Projeto 16: Web Crawler Simples

## 📝 Descrição
Crie um crawler que navega por páginas web, extrai links e informações usando net/http e HTML parsing.

## 🎯 Requisitos

### Funcionalidades Obrigatórias
1. **Crawling**:
   - Iniciar de uma URL
   - Baixar página HTML
   - Extrair links
   - Seguir links (com limite de profundidade)
   - Evitar loops (não visitar mesma URL duas vezes)

2. **Extração**:
   - Extrair títulos
   - Extrair textos
   - Extrair imagens
   - Extrair metadados

3. **Controle**:
   - Limite de páginas
   - Limite de profundidade
   - Filtro de domínios

4. **Concorrência**: Processar múltiplas URLs em paralelo (worker pool)

5. **Persistência**: Salvar resultados em JSON

## 📚 Conceitos Utilizados
- ✅ net/http
- ✅ HTML parsing (goquery ou regex)
- ✅ Goroutines
- ✅ Channels
- ✅ Worker pool
- ✅ Slices e maps
- ✅ Error handling
- ✅ Concorrência avançada

## 📁 Estrutura Sugerida
```
crawler/
├── main.go
├── crawler.go
├── parser.go
├── worker.go
└── README.md
```

## 💡 Implementação Sugerida

### Structs
```go
type Pagina struct {
    URL     string
    Titulo  string
    Links   []string
    Texto   string
    Profundidade int
}

type Crawler struct {
    visitadas map[string]bool
    limite    int
    workers   int
}
```

### Concorrência
- Worker pool para processar URLs
- Channel para fila de URLs
- Mutex para map de visitadas

## ✅ Critérios de Sucesso
- [ ] Crawling funciona
- [ ] Extração é precisa
- [ ] Concorrência funciona
- [ ] Loops são evitados
- [ ] Limites são respeitados
- [ ] Dados são salvos

## 🚀 Extras (Desafio)
- [ ] Robots.txt respect
- [ ] Sitemap generation
- [ ] Exportar para diferentes formatos
- [ ] Análise de conteúdo
- [ ] Visualização de grafo de links

