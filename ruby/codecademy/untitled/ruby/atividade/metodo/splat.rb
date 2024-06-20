#Splat!
# Falando em não saber o que esperar…
#
# Os métodos não sabem quais argumentos receberão. Ocasionalmente, os métodos nem
# sabem quantos argumentos serão passados para eles.
#
# Digamos que você tenha um método, friendque é putso argumento que ele recebe do
# usuário. Pode ser algo assim:
#

def what_up(greeting, *friends)
  friends.each { |friend| puts "#{greeting}, #{friend}!" }
end

what_up("E ai? ", "Fulano", "Beltrano", "Fulano de tal", "Carinha")