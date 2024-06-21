# Quando e então: a declaração do caso
# 5 minutos
# A instrução if/ elseé poderosa, mas podemos ficar atolados em ifs e elsifs se
# tivermos muitas condições para verificar. Felizmente, Ruby nos fornece uma
# alternativa concisa: a casedeclaração. A sintaxe é semelhante a esta:
language = "JS"
language2 = "Ruby"

case language
  when "JS"
    puts "Websites!"
  when "Python"
    puts "Science!"
  when "Ruby"
    puts "Web apps!"
  else
    puts "I don't know!"
end

# Mas você pode dobrá-lo assim:

case language2
  when "JS" then puts "Websites!"
  when "Python" then puts "Science!"
  when "Ruby" then puts "Web apps!"
  else puts "I don't know!"
end

#ATIVIDADE
# Adicionamos algum código para solicitar a entrada do usuário.
#
# Adicione uma case greetinginstrução (em qualquer estilo!) ao arquivo. Deve fazer o seguinte:
#
# Quando o usuário digitar “Inglês”, o programa deveráputs "Hello!"
# Quando o usuário digitar “Francês”, o programa deveráputs "Bonjour!"
# Quando o usuário digitar “Alemão”, o programa deveráputs "Guten Tag!"
# Quando o usuário digitar “Finlandês”, o programa deveráputs "Haloo!"
# Caso contrário, o programa deveria puts“Não conheço esse idioma!”
# Quando terminar, certifique-se de inserir alguma entrada no terminal depois de
# clicar no botão “Executar” para testar seu código.

puts "Hello there!"
greeting = gets.chomp

# Add your case statement below!
case greeting
  when "English" then puts "Hello!"
  when "French" then puts "Bonjour!"
  when "German" then puts "Guten Tag!!"
  when "Finnish" then puts "Haloo!"
  else puts "I don’t know that language!"
end
