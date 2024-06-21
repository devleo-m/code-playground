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
