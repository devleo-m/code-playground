#A classificação de arrays é um problema muito comum na ciência da computação e
# é bem estudado por esse motivo. Existem muitos algoritmos — sequências de etapas
# bem definidas — cada um com suas próprias vantagens e desvantagens.
#
# Algoritmos de classificação podem ser uma ótima introdução à ciência da computação
# como disciplina teórica, mas por enquanto vamos nos concentrar em como usar a
# biblioteca de classificação integrada do Ruby
# (que implementa alguns desses algoritmos).

my_array = [3, 4, 8, 7, 1, 6, 5, 9, 2]

puts my_array.sort!

#Se lhe entregássemos cinco livros e pedíssemos que os organizasse, ordenados por título,
# numa estante, como você faria isso?
#
# A maioria dos algoritmos de classificação assume que estamos classificando uma matriz
# de itens, o que envolve comparar dois itens quaisquer na matriz e decidir qual deve vir primeiro.
#
# Para o nosso exemplo de livros, se para qualquer par sempre escolhêssemos o livro cujo
# título aparece primeiro no alfabeto, poderíamos elaborar uma estratégia de classificação.
# Essas “estratégias” são os algoritmos de classificação mencionados no exemplo anterior.
# Nosso trabalho é decidir como comparar dois itens da lista e deixar Ruby decidir qual
# estratégia usar.

books = ["Charlie and the Chocolate Factory", "War and Peace", "Utopia", "A Brief History of Time", "A Wrinkle in Time"]
puts books.sort!

#Deixando em ordem alfabete normal e ao contrario:
def alphabetize(arr, rev=false)
  if rev
    arr.sort { |item1, item2| item2 <=> item1 }
  else
    arr.sort { |item1, item2| item1 <=> item2 }
  end
end

books = ["Heart of Darkness", "Code Complete", "The Lorax", "The Prophet", "Absalom, Absalom!"]

puts "A-Z: #{alphabetize(books)}"
puts "Z-A: #{alphabetize(books, true)}"



