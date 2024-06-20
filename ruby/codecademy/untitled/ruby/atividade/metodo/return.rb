# Às vezes, não queremos apenas que um método imprima algo no console, mas na
# verdade queremos que esse método nos devolva (ou outro método!) um valor.
# Para isso, usamos return.
def double(n)
  return n * 2
end

output = double(6)
output += 2
puts output

=begin
Defina um método chamado greeterque receba um único parâmetro de string,
namee retorne uma string contendo o nome dessa pessoa.

Certifique-se returnda string. Não use printou puts.
Defina um by_three?método que receba um único parâmetro inteiro, numbere
retorne truese esse número for divisível por três e falsese não for.
=end

def greeter(name)
  return name
end

def by_three? (number)
  if number % 3 == 0
    return true
  else
    return false
  end
end

