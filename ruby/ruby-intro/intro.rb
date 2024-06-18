#################################### Intro ###################################

# Strings
nome = "João"

# Integers
idade = 30

# Floats
altura = 1.75

# Arrays
cores = ["vermelho", "azul", "verde"]

# Hashes
pessoa = { "nome" => "Ana", "idade" => 25 }

#################################### Operadores ###################################

# Aritméticos
soma = 5 + 3
subtracao = 5 - 3
multiplicacao = 5 * 3
divisao = 5 / 3

# Comparação
maior_que = 5 > 3
menor_que = 5 < 3

# Lógicos
e = true && false
ou = true || false
nao = !true

#################################### Estrutura de controle ###################################

# if, else, elsif
if idade > 18
    puts "Você é maior de idade."
elsif idade == 18
    puts "Você acabou de se tornar um adulto."
else
    puts "Você é menor de idade."
end

# unless
unless idade >= 18
    puts "Você não pode votar."
else
    puts "Você pode votar."
end

# case
case idade
when 0..12 then puts "Criança"
when 13..18 then puts "Adolescente"
else puts "Adulto"
end

#################################### LOOPS ###################################

# while
contador = 0
while contador < 5 do
    puts contador
    contador += 1
end

# for
for numero in 1..5 do 
    puts numero 
end

# each (com array)
cores.each do |cor|
    puts cor 
end

# each (com hash)
pessoa.each do |chave, valor|
    puts "#{chave}: #{valor}"
end

################################### Metodos e Blocos ###################################

# Método simples que diz olá para uma pessoa.
def dizer_ola(nome)
    puts "Olá #{nome}!"
end

dizer_ola("Maria")

# Bloco que é passado para o método 'times'.
5.times do |numero|
    puts numero 
end

# Método que aceita um bloco.
def fazer_algo_importante 
    yield if block_given?
end

fazer_algo_importante { puts "Fazendo algo importante!" }
