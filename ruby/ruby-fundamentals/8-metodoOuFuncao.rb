# Estrutura de um Método
# Definição de um método

def sobrenome(nome)
    puts "#{nome} da silva"
end

# def: é a palavra-chave que inicia a definição de um método.
# saudacao: é o nome do método.
# nome: é um parâmetro do método.

# Chamada de um método:
sobrenome("fulano")

# saudacao é o nome do método sendo chamado.
# "fulano" é o argumento passado para o parâmetro nome.

# Parâmetros e Retorno
# Os métodos podem aceitar zero ou mais parâmetros.

def soma(x, y)
    puts x + y
end

# a e b são parâmetros do método soma.
# O valor retornado de um método é a última expressão avaliada.
soma(4,4)

# Escopo de Variáveis:
# Variáveis locais:

# Variáveis declaradas dentro de um método são locais para esse método.
def escopo_local(x, y)
    escopo = "valor locais de x:#{x} e y:#{y}"
    return escopo
end

puts escopo_local("escopo no","modulo")
#puts escopo /aqui e gerado um erro, pois a variavel esta apenas visivel
# no modulo escopo_local, apenas la dentro que eu posso mexer na variavel

# Variáveis globais:
# Variáveis globais podem ser acessadas em qualquer lugar do código.

$variavel_global = "estou em todo lugar" #variavel global criado
def mensagem #metodo
    puts $variavel_global #printando a variavel global
end

mensagem #executando o metodo

# Métodos e Blocos:
# Métodos podem receber blocos de código como argumentos usando yield.
def executando_bloco
    puts "1: antes do bloco"
    yield if block_given?
    puts "3: depois do bloco"
end

executando_bloco { puts "2: dentro do bloco"}
# yield: invoca o bloco passado como argumento para o método.

#ULTIMO
#Metodos de utilizacao
def calcula_media(numeros)
    total = numeros.reduce(0) { |sum, num| sum + num }
    media = total.to_f / numeros.length
    return media
end

notas = [8, 7, 9, 6, 4]
media_calculada = calcula_media(notas)
puts "Media eh: #{media_calculada}"


