#Mencionamos que a pesquisa de hash é mais rápida com chaves de símbolos do que com
# chaves de string. Aqui, vamos provar isso!
#
# O código no editor usa uma sintaxe nova, então não se preocupe em entender tudo ainda.
# Ele constrói dois hashes de alfabeto: um que emparelha letras de string com seu lugar
# no alfabeto ( “a” com 1, “b” com 2…) e outro que usa símbolos ( :acom 1, :bcom 2…).
# Procuraremos a letra “r” 100.000 vezes para ver qual processo é executado mais rápido!
#
# É bom ter em mente que os números que você verá estão separados por apenas frações de
# segundo, e fizemos a pesquisa de hash 100.000 vezes cada. Não é um grande aumento de
# desempenho usar símbolos neste caso, mas definitivamente existe!

#Os números não mentem. Clique em Executar para ver por si mesmo!

require 'benchmark'

string_AZ = Hash[("a".."z").to_a.zip((1..26).to_a)]
symbol_AZ = Hash[(:a..:z).to_a.zip((1..26).to_a)]

string_time = Benchmark.realtime do
  100_000.times { string_AZ["r"] }
end

symbol_time = Benchmark.realtime do
  100_000.times { symbol_AZ[:r] }
end

puts "String time: #{string_time} seconds."
puts "Symbol time: #{symbol_time} seconds."

