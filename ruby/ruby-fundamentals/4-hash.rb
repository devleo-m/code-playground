=begin
Em Ruby, um Hash é uma estrutura de dados que armazena pares chave-valor. 
Cada chave é única dentro de um Hash, e você pode recuperar valores usando 
essas chaves de forma eficiente. 
=end

#Você pode criar um Hash usando a sintaxe {} e especificando pares chave-valor

#1 Criando um Hash:

#uma hash vazia
hash_vazio = {}

#hash com valores
pessoa = {
    "nome" => "ArolNode",
    "idade" => 90,
    "altura" => 1.70
}

#2 Símbolos como Chaves:
#É comum usar símbolos como chaves em um Hash, pois eles consomem menos 
#memória e são mais eficientes.

cliente = {
    nome: "Fulano",
    cpf: "000455883-00",
    endereco: "av. general fulano de tal",
    telefone: "(48) 9843-3999"
}

#3 Acessando Valores:
#Você pode acessar valores usando a chave correspondente.

# Acessando valores do hash pessoa
puts pessoa["nome"]
puts pessoa["idade"]
# Acessando valores do hash cliente
puts cliente[:nome]
puts cliente[:cpf]

#4 Adicionando e Atualizando Pares Chave-Valor:
pessoa["profissao"] = "Developer"
cliente[:estado] = "Pará"

#printar as novas chaves atualizadas
puts pessoa["profissao"]
puts cliente[:estado]

# 5 Removendo pares chaves-valor
#Para remover um par chave-valor, você pode usar o método delete.

pessoa.delete("profissao")
cliente.delete(:estado)
#dados profissao e estado deletados das duas hashs

#6 Verificando a Existência de uma Chave:
#Você pode verificar se uma chave específica existe em um 
#Hash usando o método has_key? ou key?.

if pessoa.has_key?("profissao")
    puts "a chave profissao existe"
else
    puts "a chave profissao nao existe"
end

if cliente.key?(:cpf)
    puts "o cpf do cliente existe #{cliente[:cpf]}"
    else
        puts "o cpf do cliente nao existe"
end

#7 Iterando sobre um Hash
#Você pode percorrer um Hash usando métodos como each.

pessoa.each do |key, valor|
    puts "#{key} #{valor}"
end

# 8. Métodos Úteis de Hash:
# Ruby fornece vários métodos úteis para trabalhar com Hashes, como keys, 
# values, merge, entre outros.

# Obtendo todas as chaves do hash pessoa
todas_as_chaves = pessoa.keys

# Obtendo todos os valores do hash pessoa
todos_os_valores = pessoa.values

#cada variavel pega exatamente apenas as keys da hash ou as values

# 9 Hashes Aninhados:
# Você pode ter Hashes dentro de Hashes, criando uma estrutura de dados 
# mais complexa.

empresa = {
    "nome" => "Microwondas",
        "funcionarios" => {
            "gerente" => "Fulano",
            "atendente" => "Feluano",
            "caixa" => "Ciclano",
            "estagiario" => "Beltrano"
        }
}

#or

empresa2 = {
  nome: "Microwsufli",
  funcionarios: {
    gerente: "Beltrano",
    atendente: "Ciclano",
    estagiario: "Feluano"
  }
}

# acessando ao hashs
puts empresa["funcionarios"]["gerente"]
puts empresa2[:funcionarios][:estagiario]





