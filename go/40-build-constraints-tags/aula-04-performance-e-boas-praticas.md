# Módulo 40: Build Constraints & Tags em Go
## Aula 4: Performance e Boas Práticas

---

## 1. Boas Práticas

### ✅ Organize por Plataforma

```
project/
├── main.go
├── platform/
│   ├── linux.go
│   ├── windows.go
│   └── darwin.go
```

### ✅ Sempre Tenha Arquivo Padrão

```go
// Sempre tenha implementação padrão
// Caso contrário pode não compilar!
```

### ✅ Documente Tags Customizadas

```go
//go:build tls
// Esta tag habilita suporte a TLS
// Use: go build -tags tls
```

### ✅ Teste em Todas as Plataformas

```bash
GOOS=linux go build
GOOS=windows go build
GOOS=darwin go build
```

---

## 2. Armadilhas

### ❌ Esquecer Arquivo Padrão

Sempre tenha implementação que funciona para todos.

### ❌ Constraints Conflitantes

Evite múltiplos arquivos com mesma constraint.

### ❌ Não Testar

Teste em todas as plataformas suportadas.

---

## 3. Checklist

- [ ] Código organizado por plataforma
- [ ] Arquivo padrão existe
- [ ] Tags documentadas
- [ ] Testado em todas as plataformas
- [ ] Constraints não conflitam

---

**Bons estudos! 🚀**


