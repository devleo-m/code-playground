#Blocos são como métodos sem nome
# A maioria dos métodos com os quais você trabalhou tem nomes definidos que você ou
# outra pessoa deu a eles ( ou seja, [array].sort(), “string”.downcase() e assim por diante).
# Você pode pensar nos blocos como uma forma de criar métodos que não possuem nome.
# (São semelhantes às funções anônimas em JavaScript ou lambdas em Python.)

# Os blocos podem ser definidos com palavras-chave doe end/ou com chaves ({}).

1.times do
  puts "I'm a code block!"
end

1.times { puts "As am I!" }

# Como os blocos diferem dos métodos
# Existem algumas diferenças entre blocos e métodos, no entanto.
# Confira o código no editor. O capitalizemétodo coloca uma palavra em maiúscula e
# podemos invocá- capitalizelo continuamente pelo nome. Podemos capitalize("matz"),
# capitalize("eduardo")ou qualquer corda que quisermos, o quanto quisermos.

#No entanto, o bloco que definimos (a seguir .each) será chamado apenas uma vez e no
# contexto do array sobre o qual estamos iterando. Parece tempo suficiente para fazer
# algum trabalho eache depois desaparece na noite.

def capitalize(string)
  puts "#{string[0].upcase}#{string[1..-1]}"
end

capitalize("ryan") # prints "Ryan"
capitalize("jane") # prints "Jane"

["ryan", "jane"].each {|string| puts "#{string[0].upcase}#{string[1..-1]}"}
