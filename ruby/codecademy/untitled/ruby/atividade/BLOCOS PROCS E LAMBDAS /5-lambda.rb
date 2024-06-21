#Assim como procs, lambdas são objetos. As semelhanças não param por aí:
# com exceção de um pouco de sintaxe e algumas peculiaridades comportamentais,
# lambdas são idênticos a procs.
#
# Confira o código no editor. Veja a lambdaparte? Digitando

lambda { puts "Hello!" }

# É quase o mesmo que

Proc.new { puts "Hello!" }

#quando passamos o lambda para lambda_demo, o método chama o lambda e executa seu código.

def lambda_demo(a_lambda)
  puts "I'm the method!"
  a_lambda.call
end

lambda_demo(lambda { puts "I'm the lambda!" })

#Sintaxe Lambda
# Lambdas são definidos usando a seguinte sintaxe:

#lambda { |param| block }

# Lambdas são úteis nas mesmas situações em que você usaria um proc.
# Abordaremos as diferenças entre lambdas e procs no próximo exercício;
# enquanto isso, vamos praticar um pouco com a lambdasintaxe.

strings = ["leonardo", "donatello", "raphael", "michaelangelo"]

symbolize = lambda {|x| x.to_sym}

symbols = strings.collect(&symbolize)
print symbols

#Lambdas vs. procs

#Se você está pensando que procs e lambdas são super parecidos, é porque são!
# Existem apenas duas diferenças principais.
#
# Primeiro, um lambda verifica o número de argumentos passados a ele, enquanto um proc não.
# Isso significa que um lambda gerará um erro se você passar o número errado de argumentos,
# enquanto um proc ignorará argumentos inesperados e atribuirá nilaqueles que estiverem faltando.
#
# Segundo, quando um lambda retorna, ele passa o controle de volta para o método de chamada;
# quando um proc retorna, ele o faz imediatamente, sem voltar ao método de chamada.
#
# Para ver como isso funciona, dê uma olhada no código no editor.
# Nosso primeiro método chama um proc; o segundo chama um lambda.

def batman_ironman_proc
  victor = Proc.new { return "Batman will win!" }
  victor.call
  "Iron Man will win!"
end

puts batman_ironman_proc

def batman_ironman_lambda
  victor = lambda { return "Batman will win!" }
  victor.call
  "Iron Man will win!"
end

puts batman_ironman_lambda

#Você escreveu seu próprio lambda e viu como passá-lo para um método.
# Agora é hora de escrever um lambda e passá-lo para um método!
#
# Se você acha que isso será muito parecido com o que você já fez com procs, você está certo.
# Assim como com procs, precisaremos colocar &no início do nome do nosso lambda quando o
# passarmos para o método, pois isso converterá o lambda no bloco que o método espera.
#
# Esse symbolizelambda foi muito legal. Vamos abordar isso com um lambda que verifica se
# cada elemento em um array é um símbolo. Podemos fazer essa verificação com o .is_a?método,
# que retorna truese um objeto é do tipo de objeto nomeado e falsecaso contrário:

:hello.is_a? Symbol
# ==> true

#A palavra Symboldeve estar em maiúscula quando você estiver fazendo um .is_a?cheque!

my_array = ["raindrops", :kettles, "whiskers", :mittens, :packages]

# Add your code below!
symbol_filter = lambda {|x| x.is_a? Symbol}
symbols = my_array.select(&symbol_filter)

my_array = ["raindrops", :kettles, "whiskers", :mittens, :packages]

puts symbols

#Revisão rápida
#Toda essa conversa sobre blocos, procs e lambdas pode deixar sua cabeça girando.
# Vamos reservar um minuto para esclarecer exatamente o que cada um é:
#
# Um bloco é apenas um pedaço de código entre do.. endou {}. Não é um objeto por si só,
# mas pode ser passado para métodos como .eachou .select.
# Um proc é um bloco salvo que podemos usar continuamente.
# Um lambda é como um proc, só que se preocupa com o número de argumentos que obtém e
# retorna ao seu método de chamada em vez de retornar imediatamente.
# Obviamente, existem muitos casos em que blocos, procs e lambdas podem fazer um trabalho
# semelhante, mas as circunstâncias exatas do seu programa o ajudarão a decidir qual deles
# deseja usar.

odds_n_ends = [:weezard, 42, "Trady Blix", 3, true, 19, 12.345]

# Add your code below!

ints = odds_n_ends.select {|x|x.is_a? Integer}

puts ints

#########################

ages = [23, 101, 7, 104, 11, 94, 100, 121, 101, 70, 44]

# Add your code below!
under_100 = Proc.new { |x| x < 100 }
youngsters = ages.select(&under_100)

puts youngsters

#########################

crew = {
  captain: "Picard",
  first_officer: "Riker",
  lt_cdr: "Data",
  lt: "Worf",
  ensign: "Ro",
  counselor: "Troi",
  chief_engineer: "LaForge",
  doctor: "Crusher"
}
# Add your code below!
first_half = lambda { |x,y| y<"M"}
a_to_m = crew.select(&first_half)

puts a_to_m


