#Ocultar informações é uma parte importante da programação: proteger senhas,
# estabelecer conexões seguras e proteger programas contra adulterações, tudo isso
# depende do controle do acesso às informações.
#
# Embora não possamos realmente nos aprofundar na ocultação de informações até
# cobrirmos os hashes em um curso posterior, podemos escrever um programa simples
# para alterar a entrada de um usuário com as ferramentas que temos agora: arrays e iteradores

puts "qual e o text?"
text = gets
puts "qual e o redacte?"
redact = gets.chomp

puts "Voce digitou: #{text} e #{redact}"

#O método .split
# Ruby tem um método interno para isso chamado .split; ele recebe uma string e retorna um array.
# Se passarmos um pedaço de texto entre parênteses, .splitdividiremos a string onde quer que veja
# esse pedaço de texto, chamado de delimitador

words = text.split(" ")
puts "Voce digitou: #{words}"

#Tudo bem! É hora de usar nossos iteradores para analisar o texto do usuário.

letras = ["a", "b", "c", "d", "e"]
letras.each do |letra|
  print "#{letra.upcase}"
end




