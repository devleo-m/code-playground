#A sintaxe hash que você viu até agora (com o =>símbolo entre chaves e valores)
# às vezes é apelidada de estilo hash rocket.

numbers = {
  :one => 1,
  :two => "two",
  :three => 3,
}

#Isso ocorre porque =>parece um pequeno foguete!

# No entanto, a sintaxe hash mudou no Ruby 1.9. Bem quando você estava ficando confortável!

# A boa notícia é que a sintaxe alterada é mais fácil de digitar do que a antiga sintaxe
# hash rocket, e se você estiver acostumado com objetos JavaScript ou dicionários Python,
# ela parecerá muito familiar:

movies = {
  children: "Moana",
  scifi: "Mortal Kombat",
  history: "Lincoln"
}

#É importante observar que, embora essas chaves tenham dois pontos no final em vez de no
# início, elas ainda são símbolos!
puts new_hash
# => { :one => 1, :two => 2, :three => 3 }
#De agora em diante, usaremos a sintaxe hash 1.9 ao dar exemplos ou fornecer código padrão.
# Você vai querer estar familiarizado com o estilo hash rocket ao ler o código de outras pessoas,
# que pode ser mais antigo.


