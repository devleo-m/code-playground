# Crie um codigo que recebe o nome do usuario, crie um if para captar se o nome do usuario
# tem um S no nome e troque o "s" por um "th"

print "qual seu nome?"
user_input = gets.upcase!
user_input.downcase!

if user_input.include? "s"
  user_input.gsub!(/s/, "th")
end

print user_input