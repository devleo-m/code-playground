# Módulo 39: Unsafe Package em Go

⚠️ **AVISO**: Este módulo trata de código perigoso. Use apenas quando absolutamente necessário!

Bem-vindo ao módulo sobre o **Unsafe Package** em Go. Este módulo ensina como usar `unsafe` para manipulação direta de memória, mas com extremo cuidado.

## 📚 Estrutura do Módulo

### Aula 1: Unsafe Package (Principal)
**Arquivo**: `aula-01-unsafe-package-principal.md`

Conteúdo sobre:
- O que é unsafe e por que existe
- Funcionalidades: Pointer arithmetic, conversões
- Casos de uso reais
- Riscos e cuidados
- Boas práticas

**Tempo estimado**: 2-3 horas

### Aula 2: Versão Simplificada
**Arquivo**: `aula-02-unsafe-package-simplificada.md`

Analogias e explicações simples.

### Aula 3: Exercícios e Reflexão
**Arquivo**: `aula-03-exercicios-e-reflexao.md`

Exercícios práticos (com cuidado!).

### Aula 4: Performance e Boas Práticas
**Arquivo**: `aula-04-performance-e-boas-praticas.md`

Segurança, validação, testes.

---

## ⚠️ Avisos Importantes

1. **Unsafe é perigoso**: Pode causar crashes, vulnerabilidades, comportamento indefinido
2. **Use apenas quando necessário**: Prefira sempre código seguro
3. **Valide sempre**: Nunca confie em inputs
4. **Documente extensivamente**: Explique riscos e requisitos
5. **Teste extensivamente**: Diferentes arquiteturas e casos

---

## 🎯 Quando Usar

✅ **Use apenas se:**
- Absolutamente necessário
- Performance crítica (e você mediu)
- Entende completamente
- Casos muito específicos (systems programming)

❌ **NÃO use se:**
- Há alternativa segura
- Performance não é crítica
- Não entende completamente
- Código de produção geral

---

## 📖 Conceitos Principais

- **unsafe.Pointer**: Chave universal para qualquer pointer
- **Pointer arithmetic**: Andar pela memória
- **Conversões**: Converter entre tipos incompatíveis
- **Riscos**: Crashes, vulnerabilidades, comportamento indefinido

---

**Bons estudos e use com extremo cuidado! ⚠️**



