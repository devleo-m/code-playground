=begin
Definir um método é ótimo, mas não é muito útil para você, a menos que você o chame ou
diga ao seu programa para executá-lo. Por exemplo, se você chamar um método chamado
cartoon_fox, o programa começará a procurar o método com aquele nome e tentará executar
o código dentro dele.

Se o programa não encontrar um método chamado cartoon_fox, ele retornará um NameError.
Abordaremos erros em outra lição.

Você chama um método apenas digitando seu nome. Lembra quando você nos viu digitando
puts_1_to_10ou greetingdepois da definição do método nos dois últimos exercícios? Éramos
nós chamando nossos métodos!
=end

def hello
  puts "Hello!"
end

def array_of_10
  puts (1..10).to_a
end

# Parâmetros e argumentos
# Se um método aceita argumentos, dizemos que ele aceita ou espera esses argumentos.
# Poderíamos definir uma função, squareassim:

def square(n)
  puts n ** 2
end

square(12)
# ==> prints "144"

#DICA:
# Os parênteses geralmente são opcionais em Ruby, mas é uma boa ideia colocar seus
# parâmetros e argumentos entre parênteses para facilitar a leitura.