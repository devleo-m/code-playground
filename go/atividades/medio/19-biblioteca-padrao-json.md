# Atividade 19: Trabalhando com JSON

## Objetivo
Dominar serialização e deserialização JSON usando biblioteca padrão.

## Enunciado
Crie um programa que:
1. Defina struct `Usuario` com campos JSON tags apropriados
2. Crie função `serializarUsuario(u Usuario) ([]byte, error)` usando `json.Marshal`
3. Crie função `deserializarUsuario(dados []byte) (Usuario, error)` usando `json.Unmarshal`
4. Crie função `serializarIndentado(u Usuario) ([]byte, error)` com formatação
5. Trabalhe com JSON aninhado (struct dentro de struct)
6. Implemente `json.Marshaler` e `json.Unmarshaler` para tipo customizado (ex: data formatada)

## Exemplo de Saída
```
Usuário serializado:
{"nome":"João","idade":30,"email":"joao@email.com"}

Usuário indentado:
{
  "nome": "João",
  "idade": 30,
  "email": "joao@email.com"
}

Deserializado: João, 30 anos
```

## Dicas
- Use tags: `json:"nome_campo"`
- `json.Marshal` serializa, `json.Unmarshal` deserializa
- `json.MarshalIndent` para formatação
- Implemente interfaces para controle customizado
- Trate erros de parsing

