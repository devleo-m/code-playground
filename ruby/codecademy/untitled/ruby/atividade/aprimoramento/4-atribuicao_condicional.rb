#Vimos que podemos usar o =operador para atribuir um valor a uma variável.
# Mas e se quisermos atribuir uma variável apenas se ela ainda não tiver sido atribuída?
#
# Para isso, podemos utilizar o operador de atribuição condicional: ||=.
# É composto pelo ||operador lógico ou ( ) e pelo =operador de atribuição normal.

#Execute o código no editor. Aqui está o que está acontecendo:

favorite_book = nil
puts favorite_book

favorite_book ||= "Cat's Cradle"
puts favorite_book

favorite_book ||= "Why's (Poignant) Guide to Ruby"
puts favorite_book

favorite_book = "Why's (Poignant) Guide to Ruby"
puts favorite_book

# favorite_bookestá definido como nil, que é Ruby para “nada”.
# Quando você tenta putsacessar a tela, obtém exatamente isso: nada!
#
# Agora nossa variável está condicionalmente definida como “Cat's Cradle”.
# Como o valor da variável não era nada antes, Ruby segue em frente e
# define-a, então você vê “Cat's Cradle” impresso.
#
# Tentamos a atribuição condicional novamente, desta vez com o “Guia (comovente)
# do Why para Ruby”. Mas espere! Nossa variável já tem um valor, “Cat's Cradle”,
# então ela permanece definida com esse valor e é isso que vemos impresso.
# 
# Finalmente, usamos a atribuição antiga regular para dizer ao Ruby para
# redefinir favorite_bookpara o “Guia (comovente) do Why para Ruby”,
# o que ele faz com prazer.