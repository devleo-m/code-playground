# Projeto 12: Gerenciador de Senhas

## 📝 Descrição
Sistema seguro para armazenar e gerenciar senhas com criptografia básica e busca.

## 🎯 Requisitos

### Funcionalidades Obrigatórias
1. **Senhas**:
   - Adicionar senha (serviço, usuário, senha, notas)
   - Listar senhas
   - Buscar senha
   - Editar senha
   - Deletar senha

2. **Segurança**:
   - Criptografia básica (usar biblioteca crypto)
   - Senha mestra para acessar
   - Não exibir senhas em texto plano

3. **Geração**:
   - Gerar senha forte aleatória
   - Configurar tamanho e caracteres

4. **Categorias**: Organizar por categoria

5. **Persistência**: Arquivo criptografado

6. **Validação**: Força da senha mestra

## 📚 Conceitos Utilizados
- ✅ Structs
- ✅ Crypto package
- ✅ Slices e maps
- ✅ Error handling
- ✅ I/O
- ✅ Strings
- ✅ Organização de código

## 📁 Estrutura Sugerida
```
senhas/
├── main.go
├── senha.go
├── criptografia.go
├── gerador.go
├── storage.go
└── README.md
```

## 💡 Implementação Sugerida

### Structs
```go
type Credencial struct {
    ID       string
    Servico  string
    Usuario  string
    Senha    string // criptografada
    Categoria string
    Notas    string
    CriadoEm time.Time
}

type Gerenciador struct {
    senhaMestra string // hash
    credenciais []Credencial
}
```

### Segurança
- Hash da senha mestra (SHA256)
- Criptografia simétrica para senhas (AES)
- Não armazenar senha mestra em texto plano

## ✅ Critérios de Sucesso
- [ ] CRUD funciona
- [ ] Criptografia funciona
- [ ] Senha mestra protege acesso
- [ ] Gerador de senhas funciona
- [ ] Dados persistem criptografados
- [ ] Interface segura

## 🚀 Extras (Desafio)
- [ ] Exportar/importar
- [ ] Auditoria de senhas (força, idade)
- [ ] Compartilhamento seguro
- [ ] Backup automático
- [ ] Autofill (simulado)



