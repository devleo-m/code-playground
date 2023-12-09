#Criando arrays
numeros = [1, 2, 3, 4, 5, 6, 7, 8]
frutas = ['maca', 'pera', 'uva', 'banana', 'laranja']
#Um array é definido usando colchetes [].

puts numeros[0] #A indexação começa em 0
puts frutas[4] #laranja

#Modificando Elementos:
frutas[4] = 'melao' #troquei a laranja pelo melao
numeros[0] = 9 # troquei o numero 1 pelo 9

#Adicionando Elementos:
numeros.push(10) #adiciona o numero 10 no final da array
frutas << "goiaba" #tbm adiciona ao final. forma mais facil de se fazer

#Removendo Elementos:
numeros.pop #Remove o último elemento
frutas.delete_at(1) #remove a pera
#O método pop remove o último elemento do array.
#delete_at remove um elemento na posição específica.

numeros.each do |key|
    #puts key
end

frutas.each do |key|
    #puts key
end

#O método each itera sobre cada elemento do array.
#each_with_index permite acessar o índice durante a iteração.

frutas.each_with_index do |nome, indice|
    puts "#{indice + 1} = #{nome}"
end

#Métodos Úteis
#length ou size
puts "o tamanho da array eh: #{frutas.length}"
puts "o tamanho da array eh: #{frutas.size}"
#O método length (ou size) retorna o número de elementos presentes no array.
#Ambos podem ser usados, e a escolha entre eles é uma questão de preferência

#first e last
puts "O primeiro elemento é: #{frutas.first}"
puts "O ultimo elemento é: #{frutas.last}"
#Os métodos first e last retornam o primeiro e o último elemento da array, 
#Isso é útil quando você precisa acessar esses elementos de maneira rápida.

#include?
puts "O número 3 está presente no array? #{frutas.include?("banana")}"
#O método include? verifica se um determinado elemento está presente no array. 
#Retorna true se estiver e false se não estiver. 
#É uma maneira eficiente de verificar a presença de um valor no array.




