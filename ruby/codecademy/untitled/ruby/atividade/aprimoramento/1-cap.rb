# Um 'if' mais simples em Ruby
# Em Ruby, a instrução if é usada para executar um bloco de código somente se uma condição
# for verdadeira. A forma tradicional de usar if é a seguinte:

casado = true
puts "Eu sou solteiro" unless casado

condition = true

if condition
  # Faça alguma coisa!
end
# No entanto, se a "coisa a ser feita" for uma expressão curta e simples, você pode colocá-la
# em uma única linha. Isso torna o código mais conciso e fácil de ler.
#
#   Sintaxe de 'if' em uma linha
# A sintaxe para escrever uma instrução if em uma linha é:
#
# expressão if condição
# expressão: O que você quer que seja executado.
#   if: A palavra-chave que indica a condição.
#   condição: A condição que precisa ser verdadeira para a expressão ser executada.
#   Ordem Importante
# A ordem dos elementos é importante. O Ruby espera uma expressão seguida de if e então a condição. Vamos ver isso em ação:

puts "É verdade!" if true
# Neste exemplo, a expressão puts "É verdade!" será executada porque a condição true é verdadeira.

# Exemplo incorreto:
#if true puts "É verdade!"

# Esse exemplo é incorreto e resultará em um erro de sintaxe, pois a ordem correta não foi seguida.

#   Observações Importantes
# Sem 'end': Quando você escreve a instrução if em uma linha, você não precisa usar end.
#   Uso ideal: Este formato é ideal para expressões simples e curtas. Para blocos de código mais complexos, o uso tradicional com end é preferível por questões de clareza.
#   Exemplos Práticos
# Exemplo 1: Verificação de número positivo

numero = 10
puts "Número positivo" if numero > 0
# Se numero for maior que 0, a mensagem "Número positivo" será impressa.

  # Exemplo 2: Verificação de string vazia

# texto = ""
# puts "Texto está vazio" if texto.empty?
# Se texto estiver vazio, a mensagem "Texto está vazio" será impressa.
#
#   Conclusão
# A utilização de if em uma linha é uma maneira eficiente de escrever código Ruby mais conciso,
# especialmente quando a expressão a ser executada é curta e simples.
#   Lembre-se de seguir a ordem correta e use essa forma quando apropriado para melhorar a
