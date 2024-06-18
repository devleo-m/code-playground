# Concatenação
saudacao = "Olá, " + "mundo!"
puts saudacao # Saída: Olá, mundo!

# Interpolação
nome = "João"
mensagem = "Bem-vindo, #{nome}!"
puts mensagem # Saída: Bem-vindo, João!

# Substituição
texto = "Olá, mundo!"
novo_texto = texto.gsub("mundo", "Ruby")
puts novo_texto # Saída: Olá, Ruby!

# Divisão
palavras = "Olá mundo Ruby".split
puts palavras # Saída: ["Olá", "mundo", "Ruby"]
