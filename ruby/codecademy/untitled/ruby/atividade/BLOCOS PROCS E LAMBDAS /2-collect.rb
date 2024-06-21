#O collectmétodo pega um bloco e aplica a expressão do bloco a cada elemento de um array.
my_nums = [1, 2, 3]

var_collection = my_nums.collect { |num| num ** 2 }

print "Collection #{var_collection}"
# ==> [1, 4, 9]

puts ""
#Se olharmos para o valor de my_nums, veremos que ele não mudou:
print my_nums
# ==> [1, 2, 3]

#Isso ocorre porque .collectretorna uma cópia de my_nums, mas não altera (ou altera ) o
# array original my_nums. Se quisermos fazer isso, podemos usar .collect!com um ponto
# de exclamação:

my_nums.collect! { |num| num ** 2 }
# ==> [1, 4, 9]
my_nums
# ==> [1, 4, 9]

#Lembre-se de que !em Ruby significa “este método pode fazer algo perigoso ou inesperado!”
# Nesse caso, ele altera o array original em vez de criar um novo.

