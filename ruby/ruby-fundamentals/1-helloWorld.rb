#MALDICAO
#Para aprender ruby voce primeiro deve seguir a tradição Hello World
#e tambem e claro, sair da maldicao hahaha

puts "Hello world";

#Desta forma voce aprendeu a da o famoso printl/ console.log ou escreval
########################################################################

#VARIAVEIS
name = 'Arolnode'
puts name + ", Ola" #Concatenacao usando operador +
puts "#{name}, tudo joia?" #Concatenacao usando interpolacao de strings
#Depois vamos aprofundar sobre as variaveis e tambem principalmente escopo

#COMENTAR
# eh so usar # que o seu codigo vai ser comentado em apenas uma linha

=begin
    Esta e outra forma de comentar
    em ruby, so que ele comenta
    em varias linhas
    :)
=end

#Saida de dados
#Existe dois tipos de saidas de dados o puts e print igual no python
print 25
puts 25
#Observe que o puts ele pulou uma linha e o print nao, formou 2525

#METODO //FUNCAO
=begin
    Em Ruby, métodos são blocos de código que realizam uma tarefa específica
    e podem ser invocados (chamados) a partir de outros lugares no código.
    Eles são semelhantes a funções em js. 
=end

def saudacao
    puts 'Ola mundo'
end

saudacao
#apenas escrevendo o nome do metodo ele ja e invocado em ruby

#Estruturas de Controle
=begin
    Ruby suporta estruturas de controle padrão, como; 
    if, else, elsif, unless, case, while e for. 
=end

idade = 90
if idade >= 70
    puts 'ja esta bem velinho'
elsif idade >= 18
    puts 'esta quase velho'
else 
    puts 'bebezin demais!'
end

#Arrays
#Arrays em Ruby podem conter qualquer tipo de objeto:
frutas = ['maca', 'uva', 'goiaba', 'laranja', 'abacate', 'melao', 'banana']
# forma simples de declarar uma array, mais afrente vamos estudar mais sobre

#HASHES
#Hashes sao estruturas de dados associativas:
aluno = {
    "nome" => "Arolnode",
    "idade" => 99,
    "peso" => 89.9,
    "profissao" => "Desenvolvedor"
}

#ITERACAO
# ruby oferece metodos como each, times, map, etc., para iterar sobre coleções.
# vamos estudar sobre isso depois

# OBJETOS / CLASSES
class Cliente
    def initialize(nome)
        @nome = nome
    end

    def saudacao 
        puts "Seja bem vindo #{@nome}, o que deseja?"
    end
end

maria = Cliente.new("Maria")
maria.saudacao

#AQUI E APENAS UMA INTRODUCAO DA LINGUAGEM RUBY
#APENAS O BASICO DO BASICO, VOU EXPLICAR MELHOR SOBRE CADA UM NAS OUTRAS PASTAS

