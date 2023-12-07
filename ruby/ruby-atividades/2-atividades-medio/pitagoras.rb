#Faca um programa de teorema de pitagoras
#c2 = a2 + b2

puts "Vamos fazer o calculo, primeiro me de o valor de a:"
a = gets.chomp.to_f
puts "Vamos fazer o calculo, primeiro me de o valor de b:"
b = gets.chomp.to_f
puts "Perfeito, agora vamos fazer o calculo"
a = a*a
b = b*b
puts "x = #{a} + #{b}"
calc = a + b
puts "x = #{calc}"
puts "x = raiz quadrada de #{calc}"
x = calc**0.5
puts "x = #{'%.2f' % x}"

#
