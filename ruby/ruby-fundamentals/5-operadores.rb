# 1 Operadores Aritméticos:
# São utilizados para realizar operações matemáticas.

a = 5
b = 2

soma = a + b       # Adição
subtracao = a - b  # Subtração
multiplicacao = a * b  # Multiplicação
divisao = a / b    # Divisão
modulo = a % b     # Módulo (resto da divisão)

# 2 Operadores de Atribuição:
# Utilizados para atribuir valores a variáveis.

x = 10
y += 5  # Equivalente a y = y + 5
z -= 3  # Equivalente a z = z - 3

# 3 Operadores de Comparação:
# Comparam valores e retornam um valor booleano.

igual = a == b
diferente = a != b
maior_que = a > b
menor_que = a < b
maior_ou_igual = a >= b
menor_ou_igual = a <= b

# 4 Operadores Lógicos:
# Realizam operações lógicas em expressões booleanas.

e_logico = true && false
ou_logico = true || false
nao = !true

# 5 Operadores de Concatenação:
# Utilizados em strings para concatenar valores.

nome = "João"
sobrenome = "Silva"
nome_completo = nome + " " + sobrenome  # ou usando nome_completo = "#{nome} #{sobrenome}"

# 6 Operadores de Intervalo:
# Criam intervalos de valores.

intervalo = 1..5  # Inclui 1, 2, 3, 4, 5
intervalo_exclusivo = 1...5  # Inclui 1, 2, 3, 4 (exclui 5)

# 7 Operador Ternário:
# Uma forma concisa de expressar uma instrução condicional.

idade = 18
status = idade >= 18 ? "Adulto" : "Menor"
