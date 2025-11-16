# Módulo 42: CGO Basics em Go
## Aula 4: Performance e Boas Práticas

---

## 1. Boas Práticas

### ✅ Sempre Libere Memória C

```go
cStr := C.CString("hello")
defer C.free(unsafe.Pointer(cStr))  // SEMPRE!
```

### ✅ Minimize Chamadas CGO

```go
// ❌ Ruim: Muitas chamadas
for i := 0; i < 1000; i++ {
    C.process(C.int(i))  // 1000 chamadas CGO!
}

// ✅ Bom: Batch
C.process_batch(array, 1000)  // Uma chamada
```

### ✅ Use Go Puro Quando Possível

Sempre prefira Go puro sobre CGO.

### ✅ Documente Por Que Usa CGO

Sempre documente a razão de usar CGO.

---

## 2. Armadilhas

### ❌ Esquecer de Liberar Memória

Sempre use `defer C.free()`.

### ❌ Muitas Chamadas CGO

Minimize chamadas, faça batch.

### ❌ Usar CGO Desnecessariamente

Sempre considere alternativas Go primeiro.

---

## 3. Checklist

- [ ] Memória C sempre liberada
- [ ] Chamadas CGO minimizadas
- [ ] Alternativas Go consideradas
- [ ] CGO documentado
- [ ] Performance medida

---

**Bons estudos! 🚀**

