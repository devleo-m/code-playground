=begin
    Em Ruby, as variáveis são utilizadas para armazenar e representar 
    valores. Elas são fundamentais em qualquer linguagem de programação, 
    incluindo Ruby, pois permitem que você armazene informações e manipule 
    dados.
=end

#Em Ruby, você pode criar variáveis simplesmente atribuindo valores a nomes.

nome = 'Arolnode' # String
fruta_preferida = 'xxxx' # String
idade = 90 # Inteiro
salario = 1.99 # Float
verdadeiro = true # Boolean
falso = false #Boolean

#Ruby é uma linguagem dinâmica, o que significa que o tipo de uma variável 
#é determinado automaticamente.

#Utilize letras minúsculas para nomes de variáveis.
#Se o nome for composto, use sublinhados para separar palavras.

nome_do_usuario = "Fulano de tal da costa"

#Variáveis têm escopo local por padrão.
#local_variable é acessível apenas onde foi definida.

def exemplo
    $escopo_variavel = "Apenas dentro desse metodo(funcao)"
    return $escopo_variavel
end

puts exemplo
#puts escopo_variavel [Gera um erro, pois a variável não está no escopo fora do metodo]

#Constantes começam com letra maiúscula e têm escopo global.
PI = 3.14
SOBRENOME = 'run'
puts "valor de pi: #{PI}. seu nome e: Fulano de #{SOBRENOME}"

#Você pode atribuir vários valores de uma vez.
a, b, c = 1, 5, 10
puts b #Apenas o numero 5 imprimido, pois a variavel b recebe o valor 5 acima

#8. Variáveis de Instância e de Classe:
#Variáveis de instância começam com @ e são usadas em objetos.
#Variáveis de classe começam com @@ e são compartilhadas entre todas as instâncias de uma classe.

class Exemplo
    @variavel_de_instancia = "Sou uma variável de instância"
    @@variavel_de_classe = "Sou uma variável de classe"
end

#9. Variáveis Globais:
#Variáveis globais começam com $ e podem ser acessadas em qualquer lugar.
$variavel_global1 = "Sou uma variavel global1"

def teste_global
    $variavel_global2 = "Sou a segunda variavel global2"
end

#Chama a função para definir $variavel_global2
teste_global #se retirar a funcao, a variavel global 2 nao sera imprimida

#Agora imprime ambas as variáveis globais
puts $variavel_global1
puts $variavel_global2

#Ao atribuir nil, você está efetivamente removendo qualquer valor 
#previamente associado à variável.

idade = 40
idade = nil
puts idade

#O método undef não é usado para finalizar variáveis, mas sim para remover 
#métodos de uma classe.
#Em relação a variáveis, não há um método undef específico para isso.




