# Atividade 04: Structs com Embedding

## Objetivo
Praticar embedding de structs e composição.

## Enunciado
Crie um programa que:
1. Defina struct `Endereco` com Rua, Cidade, CEP
2. Defina struct `Contato` com Email, Telefone
3. Defina struct `Pessoa` que embede `Endereco` e `Contato`, e adiciona Nome, Idade
4. Crie struct `Funcionario` que embede `Pessoa` e adiciona Salario, Cargo
5. Implemente métodos para cada struct e demonstre acesso aos campos embedded
6. Crie função que aceita `Pessoa` e funciona tanto com `Pessoa` quanto `Funcionario`

## Exemplo de Saída
```
Pessoa: João Silva
  Endereço: Rua A, 123 - São Paulo
  Contato: joao@email.com, (11) 99999-9999

Funcionário: Maria Santos
  Endereço: Rua B, 456 - Rio de Janeiro
  Contato: maria@email.com, (21) 88888-8888
  Cargo: Desenvolvedora, Salário: R$ 8000.00
```

## Dicas
- Embedding: `type Pessoa struct { Endereco; Contato; Nome string }`
- Acesso direto: `pessoa.Rua` (sem `pessoa.Endereco.Rua`)
- Métodos de embedded são promovidos
- Use embedding para composição, não herança



