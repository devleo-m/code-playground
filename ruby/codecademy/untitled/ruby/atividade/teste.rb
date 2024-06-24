# Clique em Executar para testar nosso programa de classificação de filmes e
# começar a criar o seu próprio.

movies = {
  Memento: 3,
  Primer: 4,
  Ishtar: 1
}

puts "What would you like to do?"
puts "-- Type 'add' to add a movie."
puts "-- Type 'update' to update a movie."
puts "-- Type 'display' to display all movies."
puts "-- Type 'delete' to delete a movie."

choice = gets.chomp.downcase
case choice
when 'add'
  puts "What movie do you want to add?"
  title = gets.chomp
  if movies[title.to_sym].nil?
    puts "What's the rating? (Type a number 0 to 4.)"
    rating = gets.chomp
    movies[title.to_sym] = rating.to_i
    puts "#{title} has been added with a rating of #{rating}."
  else
    puts "That movie already exists! Its rating is #{movies[title.to_sym]}."
  end
when 'update'
  puts "What movie do you want to update?"
  title = gets.chomp
  if movies[title.to_sym].nil?
    puts "Movie not found!"
  else
    puts "What's the new rating? (Type a number 0 to 4.)"
    rating = gets.chomp
    movies[title.to_sym] = rating.to_i
    puts "#{title} has been updated with new rating of #{rating}."
  end
when 'display'
  movies.each do |movie, rating|
    puts "#{movie}: #{rating}"
  end
when 'delete'
  puts "What movie do you want to delete?"
  title = gets.chomp
  if movies[title.to_sym].nil?
    puts "Movie not found!"
  else
    movies.delete(title.to_sym)
    puts "#{title} has been removed."
  end
else
  puts "Sorry, I didn't understand you."
end

#Comecemos pelo princípio: vamos criar um hash para armazenar nossos filmes e suas
# classificações. A seguir, vamos solicitar a entrada do usuário para que possamos
# eventualmente armazenar pares de filmes/classificações em nosso hash.


movies = {
  StarWars: 4.8,
  Divergent: 4.7
}

puts "What would you like to do? "
puts "-- Type 'add' to add a movie."
puts "-- Type 'update' to update a movie."
puts "-- Type 'display' to display all movies."
puts "-- Type 'delete' to delete a movie."

choice = gets.chomp

#Agora queremos criar o corpo principal do nosso programa: a caseinstrução,
# que decidirá quais ações tomar com base no que o usuário digitar.
#
# ife elsesão poderosos, mas podemos ficar atolados em ifs e elsifs se tivermos
# muitas condições para verificar. Felizmente, Ruby nos fornece uma alternativa
# concisa: a casedeclaração. A sintaxe é semelhante a esta:

movies = {
  StarWars: 4.8,
  Divergent: 4.7
}

puts "What would you like to do? "

choice = gets.chomp

case choice
when "add"
  puts "Added!"
when "update"
  puts "Updated!"
when "display"
  puts "Movies!"
when "delete"
  puts "Deleted!"
else
  puts "Error!"
end

#Vamos desenvolver cada parte do case, um passo de cada vez. Começaremos com o branch “add”.

movies = {
  StarWars: 4.8,
  Divergent: 4.7
}

puts "What would you like to do? "

choice = gets.chomp

case choice
when "add"
  puts "What movie would you like to add? "
  title = gets.chomp
  puts "What rating does the movie have? "
  rating = gets.chomp
  movies[title.to_s] = rating
when "update"
  puts "Updated!"
when "display"
  puts "Movies!"
when "delete"
  puts "Deleted!"
else
  puts "Error!"
end