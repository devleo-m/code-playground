# "Subindo a escada para baixo"
# Se soubermos o intervalo de números que gostaríamos de incluir, podemos usar .uptoe
# .downto. Esta é uma solução muito mais Rubyista do que tentar usar um forloop que para
# quando uma variável do contador atinge um determinado valor.
#
# Podemos usar .uptopara imprimir um intervalo específico de valores:

95.upto(100) { |num| print num, " " }
# Prints 95 96 97 98 99 100
#
# e podemos usar .downtopara fazer a mesma coisa com valores decrescentes.
#
# Você pensa .uptoe .downtotrabalha no alfabeto? Só há uma maneira de descobrir!

"L".upto("P") { |letter| puts letter }

