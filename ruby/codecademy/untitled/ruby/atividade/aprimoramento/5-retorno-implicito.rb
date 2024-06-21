=begin
Sabemos que os métodos em Ruby podem retornar valores, e atribuímos um returnvalor a um
 método quando queremos usá-lo em outra parte do nosso programa. E se não colocarmos uma
 returndeclaração na definição do nosso método?

Por exemplo, se você não disser a uma função JavaScript exatamente o que fazer return,
ela retornará undefined. Para Python, o valor de retorno padrão é None. Mas para Ruby é
algo diferente: os métodos de Ruby retornarão o resultado da última expressão avaliada.

Isso significa que se você tiver um método Ruby como este:
=end

def add(a,b)
  return a + b
end

# Você pode simplesmente escrever:

def add2(a,b)
  a + b
end

#Observe que o proprio IDE RUBYMINE esta solicitando que a gente retire o return no metodo add




