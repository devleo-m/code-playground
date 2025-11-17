# Projeto 13: Sistema de Agenda e Compromissos

## 📝 Descrição
Sistema completo para gerenciar compromissos, eventos e lembretes com calendário.

## 🎯 Requisitos

### Funcionalidades Obrigatórias
1. **Compromissos**:
   - Criar compromisso (título, descrição, data/hora, duração)
   - Editar compromisso
   - Deletar compromisso
   - Listar compromissos

2. **Visualização**:
   - Ver por dia
   - Ver por semana
   - Ver por mês
   - Próximos compromissos

3. **Lembretes**:
   - Definir lembrete (X minutos antes)
   - Listar lembretes pendentes

4. **Filtros**:
   - Por data
   - Por período
   - Por status (passado, futuro, hoje)

5. **Conflitos**: Detectar compromissos sobrepostos

6. **Persistência**: JSON

## 📚 Conceitos Utilizados
- ✅ Structs e métodos
- ✅ Time package (datas complexas)
- ✅ Slices e maps
- ✅ Filtros e ordenação
- ✅ JSON
- ✅ Error handling
- ✅ Formatação de datas

## 📁 Estrutura Sugerida
```
agenda/
├── main.go
├── compromisso.go
├── calendario.go
├── lembrete.go
└── README.md
```

## 💡 Implementação Sugerida

### Structs
```go
type Compromisso struct {
    ID          string
    Titulo      string
    Descricao   string
    Inicio      time.Time
    Fim         time.Time
    Lembrete    time.Duration
    Repetir     bool
    Repeticao   string // "diario", "semanal", "mensal"
}

type Calendario struct {
    Compromissos []Compromisso
}
```

### Funcionalidades
- `CriarCompromisso(comp Compromisso) error`
- `ListarPorDia(data time.Time) []Compromisso`
- `ListarProximos() []Compromisso`
- `VerificarConflitos(comp Compromisso) []Compromisso`
- `VerificarLembretes() []Compromisso`

## ✅ Critérios de Sucesso
- [ ] CRUD funciona
- [ ] Visualizações são corretas
- [ ] Lembretes funcionam
- [ ] Conflitos são detectados
- [ ] Dados persistem
- [ ] Interface clara

## 🚀 Extras (Desafio)
- [ ] Compromissos recorrentes
- [ ] Sincronização com calendário
- [ ] Exportar para iCal
- [ ] Múltiplos calendários
- [ ] Compartilhamento


