#Aprendendo a render
#Por que alguns métodos aceitam bloqueio e outros não? Isso ocorre porque os métodos
# que aceitam blocos têm uma maneira de transferir o controle do método de chamada
# para o bloco e vice-versa. Podemos incorporar isso nos métodos que definimos usando
# a yieldpalavra-chave.

def block_test
  puts "1 - We're in the method!"
  puts "2 - Yielding to the block..."
  yield
  puts "3 - We're back in the method!"
end

block_test do
  puts ">>> 4 - We're in the block!"
end

#Rendendo com parâmetros

# Você também pode passar parâmetros para yield! Confira o exemplo no editor.
#
# O yield_namemétodo é definido com um parâmetro, name.
#
# Na linha 9, chamamos o yield_namemétodo e fornecemos o argumento "Eric"para o
# nameparâmetro. Como yield_namepossui uma yieldinstrução, também precisaremos
# fornecer um bloco.
#
# Dentro do método, na linha 2, primeiro fazemos putsuma declaração introdutória.
#
# Então cedemos ao bloco e passamos "Kim".
#
# No bloco, nagora é igual "Kim"e saímos puts“Meu nome é Kim”.
#
# De volta ao método, descobrimos putsque estamos entre os rendimentos.
#
# Então cedemos ao bloco novamente. Desta vez, passamos o "Eric"que armazenamos no nameparâmetro.
#
# No bloco, nagora é igual a "Eric"e saímos puts“Meu nome é Eric”.
#
# Finalmente, apresentamos putsuma declaração final.

puts "################################################################"

def yield_name(name)
  puts "1 - In the method! Let's yield."
  yield("2 - Kim")
  puts "3 - In between the yields!"
  yield(name)
  puts "4 - Block complete! Back in the method."
end

yield_name("5 - Eric") { |n| puts "5 - My name is #{n}." }

# Now call the method with your name!
yield_name("6 - Jamie") { |n| puts "6 - My name is #{n}." }

