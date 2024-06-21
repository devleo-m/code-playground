#Certamente podemos usar strings como chaves hash Ruby;
# como vimos, sempre há mais de uma maneira de fazer algo em Ruby.
# Contudo, a abordagem do Rubyista seria usar símbolos.

menagerie = {
  :foxes => 2,
  :giraffe => 1,
  :weezards => 17,
  :elves => 1,
  :canaries => 4,
  :ham => 1
}

#Você pode pensar em um símbolo Ruby como uma espécie de nome.
# É importante lembrar que símbolos não são strings:

"string" == :string # false

#Acima e além da sintaxe diferente, há um comportamento chave dos
# símbolos que os torna diferentes das strings . Embora possa haver
# várias strings diferentes com o mesmo valor, há apenas uma cópia de
# qualquer símbolo específico em um determinado momento.

#O .object_idmétodo obtém o ID de um objeto – é como Ruby sabe se
# dois objetos são exatamente o mesmo objeto. Execute o código no
# editor para ver se os dois "strings"são, na verdade, objetos
# diferentes, enquanto :symbolé o mesmo objeto listado duas vezes.

puts "string".object_id
puts "string".object_id

puts :symbol.object_id
puts :symbol.object_id

#Sintaxe do símbolo
# Os símbolos sempre começam com dois pontos ( :). Eles devem ser nomes de
# variáveis Ruby válidos, portanto o primeiro caractere após os dois
# pontos deve ser uma letra ou sublinhado ( _); depois disso, qualquer combinação
# de letras, números e sublinhados será permitida.
#
# Certifique-se de não colocar nenhum espaço no nome do seu símbolo – se você
# fizer isso, Ruby ficará confuso.

# :my symbol
# :my_symbol

# Crie uma variável antiga regular chamada my_first_symbole defina-a igual a
# qualquer símbolo válido que você desejar!

my_first_symbol = :my_symbol

#Para que servem os símbolos?
# Os símbolos aparecem em muitos lugares no Ruby, mas são usados principalmente
# como chaves hash ou para fazer referência a nomes de métodos.

sounds = {
  :cat => "meow",
  :dog => "woof",
  :computer => 10010110,
}

#Os símbolos são boas chaves hash por alguns motivos:
#
# Eles são imutáveis, o que significa que não podem ser alterados depois de criados;
# Existe apenas uma cópia de qualquer símbolo em um determinado momento, portanto
# eles economizam memória;
# Símbolos como chaves são mais rápidos do que strings como chaves pelos dois motivos acima.

#Convertendo entre símbolos e strings
# A conversão entre strings e símbolos é muito fácil.
:sasquatch.to_s
# ==> "sasquatch"

"sasquatch".to_sym
# ==> :sasquatch

#Os métodos .to_se .to_symsão o que você está procurando!
#  .
# Temos um conjunto de strings que gostaríamos de usar posteriormente como chaves
# hash, mas preferimos que sejam símbolos.
#
# Crie uma nova variável, symbolse armazene um array vazio nela.
# Use .eachpara iterar no stringsarray.
# Para cada sin strings, use .to_sympara converter sem um símbolo e use .pushpara
# adicionar esse novo símbolo a symbols.
# Imprima a symbolsmatriz.

strings = ["HTML", "CSS", "JavaScript", "Python", "Ruby"]

symbols = []

strings.each do |s|
  symbols.push(s.to_sym)
end
print symbols

#Muitos caminhos para o mesmo cume
# Lembre-se, sempre há muitas maneiras de realizar algo em Ruby.
# Converter strings em símbolos não é diferente!
# Além de usar .to_sym, você também pode usar .intern. Isso internalizará
# a string em um símbolo e funcionará exatamente como .to_sym:
puts ""
"hello".intern
# ==> :hello

#Ao examinar o código de outra pessoa, você poderá ver .to_symou .intern(ou ambos!)
# ao converter strings em símbolos.

strings2 = ["Nodejs", "Angular", "React", "Rails", "Spring"]
symbols2 = []

strings2.each do |s|
  symbols2.push(s.intern)
end
print symbols2



