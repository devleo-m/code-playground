# Criando um hash com símbolos como chaves
pessoa = {
  nome: "Maria",
  idade: 30,
  profissao: "Desenvolvedora"
}

# Acessando valores usando símbolos
puts pessoa[:nome] # Saída: Maria
puts pessoa[:idade] # Saída: 30

# Adicionando um novo par chave-valor
pessoa[:nacionalidade] = "Brasileira"
puts pessoa # Saída: {:nome=>"Maria", :idade=>30, :profissao=>"Desenvolvedora", :nacionalidade=>"Brasileira"}

# Iterando sobre um hash
pessoa.each do |chave, valor|
  puts "#{chave}: #{valor}"
end
# Saída:
# nome: Maria
# idade: 30
# profissao: Desenvolvedora
# nacionalidade: Brasileira
