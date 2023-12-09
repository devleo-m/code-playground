# Vamos explorar as principais estruturas de controle em Ruby: 
# if, else, elsif, unless, e case.

# 1 Estrutura if e else: A estrutura if é usada para executar um bloco de 
# código se uma condição for verdadeira. O bloco else é executado se a 
# condição for falsa.

idade = 18
if idade >=18
    puts "Maior de idade"
else
    puts "Menor de idade"
end

# 2 Estrutura elsif:
# Você pode usar elsif para verificar múltiplas condições.

aluno = "Arolnode"
nota = 8
if nota >= 9
    puts "Nota incrivel, voce passou de ano #{aluno}"
elsif nota>=5
    puts "Voce passou de ano #{aluno}"
else
    puts "REPROVADO!! ate ano que vem #{aluno}"
end

#3 unless
=begin
    unless xecuta um bloco de código se uma condição 
    fornecida for falsa. Ou seja, ele faz algo somente se a condição
    for avaliada como false.
=end

valor_de_x = 7

unless valor_de_x <= 5
    puts "#{valor_de_x} nao eh maior que 5"
end

# case
=begin
    case é uma estrutura para fazer múltiplas comparações de 
    maneira mais legível do que encadeamentos de 
    if-elsif-else.
=end

folga = "domingo"

case folga
when "segunda", "terca", "quarta", "quinta", "sexta"
    puts "#{folga}-feira, dia para trabalhar 8hrs"
when "sabado"
    puts "#{folga}, dia para trabalhar 6hrs"
when "domingo"
    puts "#{folga}, dia para relaxar"
end



