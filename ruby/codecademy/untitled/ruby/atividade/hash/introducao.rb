#Introdução aos hashes
# Sabemos que os arrays são indexados com números que começam com 0 e vão até o
# comprimento do array menos um. (Pense nisso: um array com quatro elementos tem
# os índices 0, 1, 2 e 3.)
#
# Mas e se quisermos usar índices numéricos que não vão de 0 até o final do array? E se não
# quisermos usar números como índices? Precisaremos de uma nova estrutura de array chamada hash.
#
# Hashes são como objetos JavaScript ou dicionários Python. Se você ainda não estudou essas
# linguagens, tudo que você precisa saber é que um hash é uma coleção de pares de valores-chave.
# A sintaxe hash é semelhante a esta:

hashes = {
  "key1" => "value1",
  "key2" => "value2",
  "key3" => "value3"
}

#Os valores são atribuídos às chaves usando =>. Você pode usar qualquer objeto
# Ruby para uma chave ou valor.

my_hash = {
  "name" => "Leo",
  "age" => 27,
  "married" => true
}

puts my_hash["name"]
puts my_hash["age"]
puts my_hash["married"]

#Usando Hash.new
# O que acabamos de mostrar foi a notação literal de hash . Chamamos assim porque você descreve
# literalmente o que deseja no hash: você dá um nome a ele e o define igual a um ou mais
# key => valuepares entre chaves.
#
# Você também pode criar um hash usando Hash.new, assim:

my_hash = Hash.new

#Definir uma variável igual a Hash.newcria um hash novo e vazio; é o mesmo que definir a
# variável igual a chaves vazias ({}).

#Adicionando a um Hash
# Podemos adicionar um hash de duas maneiras: se o criamos usando notação literal, podemos
# simplesmente adicionar um novo par de valores-chave diretamente entre as chaves. Se usarmos
# Hash.new, podemos adicionar ao hash usando a notação de colchetes:

pets = Hash.new
pets["Cachorro"] = "Simba"
# Adds the key "Stevie" with the
# value "cat" to the hash

#Você pode acessar valores em um hash como em um array.

pets2 = {
  "Stevie" => "cat",
  "Bowser" => "hamster",
  "Kevin Sorbo" => "fish"
}
puts pets2["Stevie"]
# will print "cat"

##########################################################################################
# ###################################################################################### #
# vamos da uma treinada em array e hash e como "iterar" sobre eles

frutas = ["Abacate", "Melancia", "Acerola", "Laranja", "Morango", "Goiaba", "Uva"]
tecnologias = {
  "backend" => "rails",
  "frontend" => "angular",
  "dados" => "postgres",
  "mobile" => "flutter",
  "devops" => "terraform"
}

#agora vamos acessar eles de forma diferente usando each

frutas.each do |value|
  puts "#{value} "
end

tecnologias.each do |key, value|
  puts "#{key.upcase}: #{value} "
end
