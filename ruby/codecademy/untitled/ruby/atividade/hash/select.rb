#Tornando-se mais seletivo
# Sabemos como obter um valor específico de um hash especificando a chave associada, mas e se quisermos filtrar um hash
# para valores que atendam a determinados critérios? Para isso, podemos usar .select.
movie_ratings = {
  memento: 3,
  primer: 3.5,
  the_matrix: 5,
  truman_show: 4,
  red_dawn: 1.5,
  skyfall: 4,
  alex_cross: 2,
  uhf: 1,
  lion_king: 3.5
}

good_movies = movie_ratings.select { |name, rating| rating > 3 }

# Mais métodos, mais soluções
# Muitas vezes descobrimos que queremos apenas a chave ou o valor associado a um par chave/valor,
# e é meio chato colocar ambos em nosso bloco e trabalhar apenas com um. Podemos iterar apenas
# chaves ou apenas valores?

my_hash = { one: 1, two: 2, three: 3 }

my_hash.each_key { |k| print k, " " }
# ==> one two three

my_hash.each_value { |v| print v, " " }
# ==> 1 2 3

#Vamos encerrar nosso estudo de hashes e símbolos Ruby testando esses métodos.

movie_ratings = {
  memento: 3,
  primer: 3.5,
  the_matrix: 3,
  truman_show: 4,
  red_dawn: 1.5,
  skyfall: 4,
  alex_cross: 2,
  uhf: 1,
  lion_king: 3.5
}
# Add your code below!

movie_ratings.each_key { |title| puts title }

