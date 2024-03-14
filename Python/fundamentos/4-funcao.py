# Em Python, uma função é definida usando a palavra-chave def, seguida pelo nome da função e,
# opcionalmente, parâmetros entre parênteses. A sintaxe geral para definir uma função é a seguinte:

#def nome_da_funcao(parametros):
    # Corpo da função
    # Bloco de código a ser executado quando a função for chamada


#Chamando uma Função:
def saudacao(nome):
    print("Ola "+ nome + "!")

saudacao("Fulano")

def saudacao2(nome):
    return "Ola "+ nome + "!"

print(saudacao2("Beltrano"))

# Parâmetros de Função:
# Os parâmetros são valores que podem ser passados para uma função quando ela é chamada.
# Uma função pode ter zero ou mais parâmetros. Por exemplo:

def soma(x,y):
    return x+y

print(soma(10,10))





