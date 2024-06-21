#Chamada e Resposta

# Lembra quando mencionamos que os símbolos são ótimos para fazer referência a
# nomes de métodos? Bem, .respond_to?pega um símbolo e retorna truese um objeto
# pode receber aquele método e falsecaso contrário. Por exemplo,

[1, 2, 3].respond_to?(:push)
#retornaria true, já que você pode chamar .pushum objeto array. No entanto,
[1, 2, 3].respond_to?(:to_sym)
# retornaria false, já que você não pode transformar um array em um símbolo.

#Em vez de verificar se nossa agevariável é um número inteiro, verifique se ela corresponde
# .respond_to?ao .nextmétodo. ( .nextretornará o número inteiro imediatamente após o número
# inteiro em que foi chamado, o que significa que 4.nextretornará 5.)

age = 26
age.respond_to?(:next)



