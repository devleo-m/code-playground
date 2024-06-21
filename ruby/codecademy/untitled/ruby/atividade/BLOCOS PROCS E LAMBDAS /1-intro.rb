# Aprendemos um pouco sobre blocos em Loops & Iterators e Methods, Blocks, & Sorting .
# Dissemos anteriormente que um bloco é como um método sem nome, mas isso não é
# bem verdade. (Veremos métodos sem nome reais, chamados lambdas , mais adiante nesta lição.)
#
# Um bloco Ruby é apenas um pedaço de código que pode ser executado. A sintaxe do bloco
# usa do.. endou chaves ( {}), assim:

[1, 2, 3].each do |num|
  puts num
end
# ==> Prints 1, 2, 3 on separate lines

[1, 2, 3].each { |num| puts num }
# ==> Prints 1, 2, 3 on separate lines

#Os blocos podem ser combinados com métodos como .eache .timespara executar uma instrução para
# cada elemento de uma coleção (como um hash ou array).


