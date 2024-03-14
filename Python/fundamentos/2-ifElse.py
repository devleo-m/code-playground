
#if condição:
    # Bloco de código a ser executado se a condição for verdadeira
#else:
    # Bloco de código a ser executado se a condição for falsa

#Condicional Simples:
idade = 20
if idade >= 18:
    print("Você é maior de idade.")
else:
    print("Você é menor de idade.")


#Condicional com Elif:
nota = 75
if nota >= 90:
    print("Nota: A")
elif nota >= 80:
    print("Nota: B")
elif nota >= 70:
    print("Nota: C")
else:
    print("Nota: D")

# Operadores de Comparação em Condicionais:
# ==: Igual a
# !=: Diferente de
# <: Menor que
# >: Maior que
# <=: Menor ou igual a
# >=: Maior ou igual a

# O operador ternário em Python é uma forma concisa de escrever condicionais em uma única linha.
# Ele permite avaliar uma expressão e retornar um valor com base em uma condição.
# A sintaxe geral do operador ternário em Python é:

#valor_se_verdadeiro if condição else valor_se_falso

idade = 16
status = "Maior de idade" if idade >= 18 else "Menor de idade"
print(status)


