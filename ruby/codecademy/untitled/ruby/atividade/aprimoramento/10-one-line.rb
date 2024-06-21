if 1 < 2
  puts "One is less than two!"
end

puts "One is less than two!" if 1 < 2

#Duas formas diferentes para resolver o mesmo problema, qual dos dois eh o melhor?
# O primeiro eh claro, ja que com ele fica mais facil de qualquer programador entender
# Codigo limpo nao eh apenas um codigo curto, e sim algo que todos possam entender :P

#O Operador Ternário
if 1 < 2
  puts "One is less than two!"
else
  puts "One is not less than two."
end

puts 1 < 2 ? "One is less than two!" : "One is not less than two."

#O ternario eh muito util e deixa muito mais simples
# fFora que muitas linguagens ja utilizam o ternario em suas sintaxes
# Eh uma otima boa pratica

#Em caso de muitas opções
puts "What's your favorite language?"
language = gets.chomp

if language == "Ruby"
  puts "Ruby is great for web apps!"
elsif language == "Python"
  puts "Python is great for science."
elsif language == "JavaScript"
  puts "JavaScript makes websites awesome."
elsif language == "HTML"
  puts "HTML is what websites are made of!"
elsif language == "CSS"
  puts "CSS makes websites pretty."
else
  puts "I don't know that language!"
end

puts "What's your favorite language?"
language2 = gets.chomp

case language2
when "Ruby"
  puts "Ruby is great for web apps!"
when "Python"
  puts "Python is great for science."
when "JavaScript"
  puts "JavaScript makes websites awesome."
when "HTML"
  puts "HTML is what websites are made of!"
when  "CSS"
  puts "CSS makes websites pretty."
else
  puts "I don't know that language!"
end

#Certamente utilizar case eh sem duvidas muito mais intuitivo e simples, fora que podemos
# deixar ele bem mais curto utilizando o then e chamando o print ou puts

#for
for i in (1..3)
  puts "I'm a refactoring master!"
end

3.times do
  puts "I'm a refactoring master!"
end

#Na minha opniao o times eh muito mais simples e util, mas infelizmente em muitas linguagens
# nao temos ele :(


