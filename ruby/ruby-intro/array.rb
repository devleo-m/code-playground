numeros = [1, 2, 3, 4, 5]

# Adicionando elementos
numeros.push(6)
puts numeros.inspect # Saída: [1, 2, 3, 4, 5, 6]

# Removendo elementos
numeros.delete(2)
puts numeros.inspect # Saída: [1, 3, 4, 5, 6]

# Acessando elementos
puts numeros[0] # Saída: 1

# Iterando sobre elementos
numeros.each do |numero|
  puts numero * 2
end
# Saída:
# 2
# 6
# 8
# 10
# 12
