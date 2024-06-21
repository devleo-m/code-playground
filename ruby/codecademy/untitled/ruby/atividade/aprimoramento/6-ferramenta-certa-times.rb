=begin
Mais cedo ou mais tarde, você precisará realizar uma tarefa repetitiva em seus programas.
Muitas linguagens de programação permitem que você faça isso com um forloop e, embora
Ruby inclua for loops , existem ferramentas melhores disponíveis para nós.

Se quisermos fazer algo um número específico de vezes, podemos usar o .timesmétodo, assim:
=end

5.times do
  puts "Odelay!"
end

# Se quisermos repetir uma ação para cada elemento de uma coleção, podemos usar .each:
array_de_frutas = ["Uva", "Goiaba", "Abacate", "Melao"]

array_de_frutas.each do |value|
  print "#{value} "
end

puts " "
#Agora aprenda a utilizar o each de outra forma

array_de_frutas.each do |value|
  puts value * 3
end


