#Herança é o processo pelo qual uma classe assume os atributos e métodos de outra e
# é usada para expressar um relacionamento é-um . Por exemplo, uma raposa de desenho
# animado é um mamífero de desenho animado, portanto, uma classe CartoonFox poderia
# herdar de uma classe CartoonMammal.
#
# No entanto, um Mago não é um Elfo, portanto não deve herdar da classe Elfo
# (mesmo que tenham muitos atributos e métodos mágicos em comum).
# Em vez disso, tanto o Wizard quanto o Elf poderiam herdar da mesma classe MagicalBeing.

class ApplicationError
  def display_error
    puts "Error! Error!"
  end
end

class SuperBadError < ApplicationError
end

err = SuperBadError.new
err.display_error

#Sintaxe de herança

class Application
  def initialize(name)
    @name = name
  end
end

# Add your code below!

class MyApp < Application
end

#Sobrepor
# Às vezes você desejará que uma classe que herde de outra não apenas assuma os métodos
# e atributos de seu pai, mas também substitua um ou mais deles.

#Por exemplo, você pode ter uma Emailclasse que herda de Message.
# Ambas as classes podem ter um sendmétodo que as envia, mas a versão de e-mail pode
# ter que identificar endereços de e-mail válidos e usar vários protocolos de e-mail
# sobre os quais Messagenada se sabe. Em vez de adicionar um send_emailmétodo à sua
# classe derivada e herdar um sendmétodo que você nunca usará, você pode simplesmente
# criar explicitamente um sendmétodo na Emailclasse e fazer com que ele faça tod.o o
# trabalho de envio de e-mail.

#Esta nova versão sendsubstituirá (ou seja, substituirá)
# a versão herdada de qualquer objeto que seja uma instância deEmail.

# class Creature
#   def initialize(name)
#     @name = name
#   end
#
#   def fight
#     return "Punch to the chops!"
#   end
# end

# Add your code below!

# class Dragon < Creature
#   def fight
#     return "Breathes fire!"
#   end
# end

# Por outro lado, às vezes você estará trabalhando com uma classe derivada
# (ou subclasse ) e perceberá que substituiu um método ou atributo definido na classe
# base dessa classe (também chamada de pai ou superclasse ) que você realmente precisa.
# Não tenha medo! Você pode acessar diretamente os atributos ou métodos de uma superclasse
# com a palavra-chave integrada do Ruby super.
  # A sintaxe é semelhante a esta:

# class DerivedClass < Base
#   def some_method
#     super(optional, args)
#     # Some other stuff
#   end
# end

#Quando você chama superde dentro de um método, isso diz ao Ruby para procurar na
# superclasse da classe atual e encontrar um método com o mesmo nome daquele a partir
# do qual superé chamado. Se encontrar, Ruby usará a versão do método da superclasse.

# class Creature
#   def initialize(name)
#     @name = name
#   end
#
#   def fight
#     return "Punch to the chops!"
#   end
# end

# Add your code below!

# class Dragon < Creature
#   def fight
#     puts "Instead of breathing fire. . . "
#     super
#   end
# end

#Confira o código no editor. Veja como estamos tentando Dragonherdar de Creaturee Person?
# Receberemos um superclass mismatch for class Dragonerro se tentarmos isso.
# Execute o código para ver por si mesmo!


# class Creature
#   def initialize(name)
#     @name = name
#   end
# end
#
# class Person
#   def initialize(name)
#     @name = name
#   end
# end
#
# class Dragon < Creature; end
# class Dragon < Person; end

#Nosso último tópico: a palavra-chave do Ruby super.
# ( Afinal, decidimos que gostamos Messagedo método.)initialize

class Message
  @@messages_sent = 0
  def initialize(from, to)
    @from = from
    @to = to
    @@messages_sent +=1
  end
end

class Email < Message
  def initialize(from, to)
    super
  end
end

my_message = Message.new("Ian", "Alex")

