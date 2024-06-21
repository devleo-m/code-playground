#Lembra quando dissemos que tudo é um objeto em Ruby? Bem, nós meio que mentimos.
# Blocos não são objetos, e esta é uma das poucas exceções à regra “tudo é um objeto” em Ruby.
#
# Por causa disso, os blocos não podem ser salvos em variáveis ​​e não possuem
# todos os poderes e habilidades de um objeto real. Para isso precisaremos de… procs!
#
# Você pode pensar em um proc como um bloco “salvo”: assim como você pode dar um nome
# a um pedaço de código e transformá-lo em um método, você pode nomear um bloco e
# transformá-lo em um proc. Procs são ótimos para manter seu código DRY ,
# que significa D on't R epeat Yourself .
# Com blocos, você precisa escrever seu código sempre que precisar; com um proc, você
# escreve seu código uma vez e pode usá-lo muitas vezes!

multiples_of_3 = Proc.new do |n|
  n % 3 == 0
end

print (1..100).to_a.select(&multiples_of_3)
#Observe que estou pedindo para printar numero de 1 a 100
# mas colocamos o to_a.select e adicionamos o metodo procs

#Sintaxe de procedimento
#Os processos são fáceis de definir! É só ligar Proc.newe passar no bloco que deseja salvar. Veja
# como criaríamos um proc chamado cubeque eleva um número ao cubo (eleva-o à terceira potência):

cube = Proc.new { |x| x ** 3 }

#Podemos então passar o proc para um método que de outra forma receberia um bloco, e não precisamos
# reescrever o bloco repetidamente!

[1, 2, 3].collect!(&cube)
# ==> [1, 8, 27]
[4, 5, 6].map!(&cube)
# ==> [64, 125, 216]

# (Os .collect!métodos .map!e fazem exatamente a mesma coisa.)
# O &é usado para converter o cubeproc em um bloco (já que .collect!e .map!
# normalmente pega um bloco). Faremos isso sempre que passarmos um proc para um método que
# espera um bloco.

floats = [1.2, 3.45, 0.91, 7.727, 11.42, 482.911]
# Write your code below this line!

round_down = Proc.new { |x| x.floor }

# Write your code above this line!
ints = floats.collect(&round_down)
print ints

#Por que Procs?
# Por que se preocupar em salvar nossos blocos como procs? Existem duas vantagens principais:
#
# Procs são objetos completos, portanto possuem todos os poderes e habilidades dos objetos.
# (Os blocos não.)
# Ao contrário dos blocos, os procs podem ser chamados repetidamente sem reescrevê-los.
# Isso evita que você precise redigitar o conteúdo do seu bloco toda vez que precisar
# executar um determinado trecho de código.

# Here at the amusement park, you have to be four feet tall
# or taller to ride the roller coaster. Let's use .select on
# each group to get only the ones four feet tall or taller.

group_1 = [4.1, 5.5, 3.2, 3.3, 6.1, 3.9, 4.7]
group_2 = [7.0, 3.8, 6.2, 6.1, 4.4, 4.9, 3.0]
group_3 = [5.5, 5.1, 3.9, 4.3, 4.9, 3.2, 3.2]

# Complete this as a new Proc
over_4_feet = Proc.new { |height| height >= 4 }

# Change these three so that they use your new over_4_feet Proc
can_ride_1 = group_1.select(&over_4_feet)
can_ride_2 = group_2.select(&over_4_feet)
can_ride_3 = group_3.select(&over_4_feet)

puts can_ride_1
puts can_ride_2
puts can_ride_3

#Crie seu próprio!
# OK! Hora de tirar as rodinhas.
cube = Proc.new { |x| x ** 3 }
[1, 2, 3].map(&cube)
# [1, 8, 27]

#Você criará seu próprio método que chama seu próprio proc! Faremos isso em duas etapas.
# Use o exemplo acima como um lembrete de sintaxe.

def greeter
  yield
end

phrase = Proc.new { puts "Hello there!" }
greeter(&phrase)

#Chamar um proc com um método não é muito complicado. No entanto, existe uma maneira
# ainda mais fácil.
#
# Ao contrário dos blocos, podemos chamar procs diretamente usando .callo método Ruby. Confira!

test = Proc.new { # does something }
  # test.call
}

#Lembre-se: sempre há mais de uma maneira de fazer algo em Ruby.

hi = Proc.new { puts "Hello!" }
hi.call

#Símbolos, Conheça Procs
# Agora que você está aprendendo algumas das partes mais complexas da linguagem Ruby,
# você pode combiná-las para trabalhar algumas magias verdadeiramente misteriosas.
# Por exemplo, lembra quando dissemos que você poderia passar um nome de método Ruby
# com um símbolo ? Bem, você também pode converter símbolos em procs usando aquele
# pequeno arquivo &.

strings = ["1", "2", "3"]
nums = strings.map(&:to_i)
# ==> [1, 2, 3]

#Ao mapear &:to_icada elemento de strings, transformamos cada string em um número inteiro!

numbers_array = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]
strings_array = numbers_array.map(&:to_s)
puts strings_array