#Falando em enviar para arrays, Ruby tem alguns atalhos interessantes para
# nomes de métodos comuns. Por sorte, um é para .push!
#
# Em vez de digitar o .pushnome do método, você pode simplesmente usar <<,
# o operador de concatenação (também conhecido como “a pá”) para adicionar
# um elemento ao final de um array:

[1, 2, 3] << 4
# ==> [1, 2, 3, 4]

# Boas notícias: também funciona com strings! Você pode fazer:
"Yukihiro " << "Matsumoto"
# ==> "Yukihiro Matsumoto"

#Mais exemplos:
alphabet = ["a", "b", "c"] << "d"
caption = "A giraffe surrounded by " << "weezards!"

#Interpolação de strings
# Você sempre pode usar o antigo simples +ou <<adicionar um valor variável em uma string:
drink = "espresso"
"I love " + drink
# ==> I love espresso
"I love " << drink
# ==> I love espresso

#Mas se você quiser fazer isso para valores que não sejam string, você terá que usar
# .to_spara torná-lo uma string:
age = 26
"I am " + age.to_s + " years old."
# ==> "I am 26 years old."
"I am " << age.to_s << " years old."
# ==> "I am 26 years old."

#Isso é complicado, e complicado não é o jeito Ruby. A melhor maneira de fazer isso
# é com interpolação de strings . A sintaxe é semelhante a esta:

"I love #{drink}."
# ==> I love espresso.
"I am #{age} years old."
# ==> I am 26 years old.

#Tudo que você precisa fazer é colocar o nome da variável dentro #{}de uma string!



