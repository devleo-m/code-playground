#O case é uma estrutura de controle em Ruby que é usada para comparar um valor
# com diferentes condições, similar ao switch em outras linguagens de programação.
# Ele é especialmente útil quando você tem várias condições para verificar e quer
# evitar o uso excessivo de if e elsif.

# Princípios do case em Ruby
# Sintaxe básica: Um bloco case começa com a palavra-chave case seguida pelo valor a
# ser comparado. Cada condição é definida com a palavra-chave when, seguida pelo valor
# a ser comparado. O bloco termina com a palavra-chave end.
#
#   Execução: O Ruby compara o valor fornecido com cada condição, na ordem em que
# aparecem, e executa o código associado à primeira condição verdadeira. Se nenhuma
# condição for verdadeira, ele executa o código após a palavra-chave else, se presente.
#
# Expressões: As condições podem ser expressões complexas, não apenas valores literais.
#
# Exemplo básico
# Aqui está um exemplo simples usando case:

dia_da_semana = "segunda"

case dia_da_semana

when "segunda"
  puts "Hoje eh segunda-feira"
when "terca"
  puts "Hoje eh terca-feira"
when "quarta"
  puts "Hoje eh quarta-feira"
when "quinta"
  puts "Hoje eh quinta-feira"
when "sexta"
  puts "Hoje eh sexta-feira"
else
  puts "voce digitou um valor errado"
end

#Neste exemplo, a variável dia_da_semana é comparada com cada condição
# (when "segunda-feira", when "terça-feira", etc.).
# Quando a condição é verdadeira ("segunda-feira"),
# o código dentro daquele bloco é executado (puts "Hoje é segunda-feira.").
# Se nenhuma das condições for verdadeira, o código dentro do bloco else é executado.
#
# Exemplos práticos
# Exemplo 1: Avaliação de notas
# Vamos criar um exemplo que avalia a nota de um aluno e imprime a classificação correspondente.

nota_do_aluno = 85.6

case nota_do_aluno
when 90..100
  puts "Nota excelente!"
when 70..89.99
  puts "Nota Boa!"
when 50..69.99
  puts "Passou raspando!"
when 0..49.99
  puts "Reprovado!!"
else
  puts "Nota invalida!"
end

#Exemplo 2: Verificação de tipos de dados
# Aqui está um exemplo que verifica o tipo de uma variável.
variavel = [1, 2, 3]

case variavel
when String
  puts "É uma string"
when Array
  puts "É um array"
when Hash
  puts "É um hash"
else
  puts "Tipo desconhecido"
end




