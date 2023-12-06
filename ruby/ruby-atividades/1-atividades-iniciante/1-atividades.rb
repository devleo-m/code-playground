#ATIVIDADE 1 Métodos com Parâmetros:
puts 'atividade 1'
=begin 
    1 -Crie um método chamado calcular_idade que recebe o ano de nascimento
    como parâmetro e retorna a idade atual.
    2 - Chame o método, passando o ano de nascimento que você definiu 
    anteriormente, e imprima a idade.
=end

#CODIGO:

def calcular_idade(nascimento)
    nascimento = 2023 - nascimento
    puts nascimento
end
calcular_idade(1996) #digite aqui o ano que voce nasceu

##########################################################################
#ATIVIDADE 2 Estruturas de Controle:
puts 'atividade 2'
=begin
    1- Crie uma variável chamada idade_usuario e atribua um valor à sua escolha.
    2- Use uma estrutura de controle (if, else) para imprimir "Maior de idade" 
    se a idade for maior ou igual a 18, e "Menor de idade" caso contrário.
=end

idade_usuario = 30
    if idade_usuario >= 18
        puts 'Maior de idade'
    else 
        puts 'Menor de idade'
    end
##########################################################################
#ATIVIDADE 3 Arrays e Iteração:
puts 'atividade 3'
=begin
    1- Crie um array chamado meses contendo os nomes dos meses do ano.
    2- Use um loop usando each para imprimir cada mês.

=end
meses = ['janeiro', 'fevereiro', 'marco', 'abril', 'maio', 'junho', 'julho', 'agosto', 'setembro', 'outubro', 'novembro', 'dezembro']

meses.each do |mes|
    puts mes
end

##########################################################################
#ATIVIDADE 4 Hashes:
puts 'atividade 4'
=begin
    1- Crie um hash chamado informacoes_pessoais com chaves como 
    "nome", "idade" e "profissao". Preencha com informações fictícias.
    2- Imprima as informações do hash de maneira formatada.
=end

informacoes_pessoais = {
    "nome" => "Arolnode",
    "idade" => 90,
    "profissao" => "Desenvolvedor"
}

puts informacoes_pessoais["nome"]
puts informacoes_pessoais["idade"]
puts informacoes_pessoais["profissao"]

puts informacoes_pessoais


##########################################################################
#ATIVIDADE 5 :
puts 'atividade 5'
=begin
    1- Crie uma classe chamada Carro com um método de inicialização que 
    define atributos como modelo e cor.
    2- Crie um objeto da classe Carro e imprima informações sobre esse carro.
=end

class Carro
    def initialize(modelo, cor)
        @modelo = modelo
        @cor = cor
    end

    def imprimir_informacoes
        puts "Carro: #{@modelo} da cor #{@cor}"
    end
end

carro = Carro.new("Volvo", "Cinza")
carro.imprimir_informacoes










    




