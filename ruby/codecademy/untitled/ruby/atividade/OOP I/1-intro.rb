#Ruby é uma linguagem de programação orientada a objetos, o que significa que ela
# manipula construções de programação chamadas objetos . (Quase) tudo em Ruby é um
# objeto! Você os tem usado o tempo tod.o, então eles devem ser muito familiares.
# Objetos possuem métodos, que você já viu antes, e atributos , que são apenas dados.
# Por exemplo, em
#
"Matz".length
# ==> 4

#o "Matz"objeto é uma string com um .lengthmétodo e um atributo de comprimento 4.
# Aprenderemos como construir nossos próprios objetos com seus próprios métodos e
# variáveis internas nos próximos exercícios.
#
# Mas o que exatamente constitui "Matz"uma string? O fato de ser uma instância
# da String classe . Uma classe é apenas uma forma de organizar e produzir objetos
# com atributos e métodos semelhantes.

class Language
  def initialize(name, creator)
    @name = name
    @creator = creator
  end

  def description
    puts "I'm #{@name} and I was created by #{@creator}!"
  end
end

ruby = Language.new("Ruby", "Yukihiro Matsumoto")
python = Language.new("Python", "Guido van Rossum")
javascript = Language.new("JavaScript", "Brendan Eich")

ruby.description
python.description
javascript.description

#Sintaxe de classe
# Uma classe básica consiste apenas na classpalavra-chave e no nome da classe. Confira:
class NewClass
  # Class magic here
end

#Nosso NewClasstem a capacidade de criar novos objetos Ruby de classe NewClass
# (assim como "Hello!"is a Stringe 4is a Fixnum).
# Por convenção, os nomes das classes começam com letra maiúscula e usam CamelCase
# em vez de confiar_em_underscores.

#temos que ter certeza de que cada pessoa tem um arquivo @name.
#
# Em Ruby, usamos @antes de uma variável para indicar que é uma variável de instância .
# Isso significa que a variável está anexada à instância da classe.

class Person
  def initialize(nome)
    @nome = nome
  end
end

matz = Person.new("Yukihiro")

#Escopo
# Outro aspecto importante das classes Ruby é o escopo . O escopo de uma variável é o contexto no
# qual ela é visível para o programa.
# Você pode ficar surpreso ao saber que nem todas as variáveis são acessíveis a todas as
# partes de um programa Ruby o tempo tod.o. Ao lidar com classes, você pode ter variáveis
# que estão disponíveis em qualquer lugar ( variáveis globais ), algumas que estão disponíveis
# apenas dentro de determinados métodos ( variáveis locais ), outras que são membros de uma
# determinada classe ( variáveis de classe ) e variáveis que só estão
# disponíveis para instâncias específicas de uma classe ( variáveis de instância ).

#O mesmo vale para métodos: alguns estão disponíveis em qualquer lugar, alguns estão
# disponíveis apenas para membros de uma determinada classe e alguns estão disponíveis
# apenas para objetos de instância específicos.

class Computer
  $manufacturer = "Mango Computer, Inc."
  @@files = {hello: "Hello, world!"}

  def initialize(username, password)
    @username = username
    @password = password
  end

  def current_user
    @username
  end

  def self.display_files
    @@files
  end
end

# Make a new Computer instance:
hal = Computer.new("Dave", 12345)

puts "Current user: #{hal.current_user}"
# @username belongs to the hal instance.

puts "Manufacturer: #{$manufacturer}"
# $manufacturer is global! We can get it directly.

puts "Files: #{Computer.display_files}"
# @@files belongs to the Computer class.

#Nomeando suas variáveis
# Lembre-se de que as variáveis de instância começam com @. Esta não é apenas
# uma convenção Ruby – faz parte da sintaxe! Sempre inicie suas variáveis de
# instância com @.
#
# Variáveis de classe são como variáveis de instância, mas em vez de pertencerem
# a uma instância de uma classe, elas pertencem à própria classe. Variáveis de
# classe sempre começam com dois @s, assim: @@files.
#
# Variáveis globais podem ser declaradas de duas maneiras. A primeira já é familiar
# para você: basta definir a variável fora de qualquer método ou classe e pronto!
# É global. Se você quiser tornar uma variável global de dentro de um método ou classe,
# basta iniciá-la com a $, assim: $matz.
#
# Examinaremos variáveis de instância e classe com mais detalhes em um momento. Primeiro,
# vamos fazer uma rápida revisão do escopo local e global.

class Person2
  def initialize(name, age, profession)
    @name = name
    @age = age
    @profession = profession
  end
end

#Twice the @, Twice as Classy
# Podemos criar variáveis ​​de classe iniciando o nome de uma variável com dois @símbolos.
# Variáveis ​​de classe são anexadas a classes inteiras ,
# não apenas a instâncias de classes, assim:

class MyClass2
  @@class_variable
end

#Como há apenas uma cópia de uma variável de classe compartilhada por todas as instâncias
# de uma classe, podemos usá-las para realizar alguns truques interessantes do Ruby.
# Por exemplo, podemos usar uma variável de classe para controlar o número de instâncias
# daquela classe que criamos. Vamos fazer isso agora!

class Person3
  # Set your class variable to 0 on line 3
  @@people_count = 0

  def initialize(name)
    @name = name
    # Increment your class variable on line 8
    @@people_count += 1
  end

  def self.number_of_instances
    # Return your class variable on line 13
    return @@people_count
  end
end

matz = Person.new("Yukihiro")
dhh = Person.new("David")

puts "Number of Person instances: #{Person.number_of_instances}"





