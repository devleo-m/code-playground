# Projeto 05: Blog Engine Simples

## 📝 Descrição
Crie um sistema de blog completo com posts, categorias, tags, comentários e busca.

## 🎯 Requisitos

### Funcionalidades Obrigatórias
1. **Posts**:
   - Criar post (título, conteúdo, autor, data)
   - Editar post
   - Deletar post
   - Listar posts (mais recentes primeiro)
   - Visualizar post completo

2. **Categorias e Tags**:
   - Associar categorias aos posts
   - Adicionar tags
   - Filtrar por categoria/tag

3. **Comentários**:
   - Adicionar comentário a post
   - Listar comentários de um post

4. **Busca**: Buscar posts por título ou conteúdo

5. **Persistência**: JSON

6. **Interface**: Menu interativo

## 📚 Conceitos Utilizados
- ✅ Structs e métodos
- ✅ Slices e maps
- ✅ Error handling
- ✅ JSON
- ✅ Strings (busca)
- ✅ Time package
- ✅ Organização de código

## 📁 Estrutura Sugerida
```
blog/
├── main.go
├── post.go
├── categoria.go
├── comentario.go
├── repositorio.go
├── busca.go
└── README.md
```

## 💡 Implementação Sugerida

### Structs
```go
type Post struct {
    ID         string
    Titulo     string
    Conteudo   string
    Autor      string
    Categoria string
    Tags      []string
    CriadoEm  time.Time
    AtualizadoEm time.Time
}

type Comentario struct {
    ID      string
    PostID  string
    Autor   string
    Conteudo string
    Data    time.Time
}
```

### Funcionalidades
- `CriarPost(post Post) error`
- `BuscarPost(id string) (Post, error)`
- `ListarPosts() []Post`
- `FiltrarPorCategoria(categoria string) []Post`
- `BuscarTexto(termo string) []Post`
- `AdicionarComentario(comentario Comentario) error`

## ✅ Critérios de Sucesso
- [ ] CRUD de posts funciona
- [ ] Categorias e tags funcionam
- [ ] Comentários são associados corretamente
- [ ] Busca retorna resultados relevantes
- [ ] Dados persistem
- [ ] Interface é intuitiva

## 🚀 Extras (Desafio)
- [ ] Markdown support
- [ ] Preview de posts
- [ ] Estatísticas (posts por mês, mais comentados)
- [ ] Exportar para HTML
- [ ] Sistema de likes
- [ ] RSS feed


