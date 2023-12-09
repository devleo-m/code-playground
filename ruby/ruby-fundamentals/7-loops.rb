#loops permitem executar um bloco de código várias vezes. 
#Existem principalmente três tipos de loops: while, for e each.

# 1 While
#executa um código enquanto uma condição específica for verdadeira.

contador = 1

while contador <= 5
    puts contador
    contador += 1 #Em ruby nao e possivel usar contador++ :(
end

puts "-----------------------------------------------------------"

# 2 For
#O for em Ruby é mais utilizado para percorrer elementos em uma faixa (range) 
#ou uma coleção.

for i in 0..3
    puts i
end

# 3 each
#ja estudamos sobre ele, mas vamos praticar mais
puts "-----------------------------------------------------------"

letras = ["a", "b", "c", "d"]

letras.each do |keys|
    puts "#{keys}"
end

# 4 Next e Break
#next é usado para pular a iteração atual e ir para a próxima.
#break é usado para sair do loop
puts "-----------------------------------------------------------"

numeros = 0
while true
    numeros += 1
    next if numeros % 2 == 0 # pula os numeros pares
    puts numeros
    break if numeros >=10 # finaliza ao chegar no numero 10
end




