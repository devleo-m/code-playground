# Módulo 42: CGO Basics em Go

Bem-vindo ao módulo sobre **CGO Basics**!

## 📚 Estrutura

- **Aula 1**: CGO Basics (Principal)
- **Aula 2**: Versão Simplificada
- **Aula 3**: Exercícios e Reflexão
- **Aula 4**: Performance e Boas Práticas

## 🚀 Início Rápido

```go
/*
#include <stdio.h>
void hello() {
    printf("Hello from C!\n");
}
*/
import "C"

func main() {
    C.hello()
}
```

## ⚠️ Avisos

- CGO desabilita cross-compilation
- Tem overhead de performance
- Use apenas quando necessário
- Sempre prefira Go puro

---

**Bons estudos! 🚀**



