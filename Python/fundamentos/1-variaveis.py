#exemplo de variaveis

nome = "Leonardo"
idade = 27
altura = 1.69
casado = True

#Tipos de Dados: Python possui diversos tipos de dados básicos, incluindo:

# Inteiros (int): Números inteiros, por exemplo: idade = 25
# Ponto Flutuante (float): Números com casas decimais, por exemplo: altura = 1.75
# String (str): Sequências de caracteres, por exemplo: nome = "João"
# Lista (list): Coleção ordenada de elementos, por exemplo: numeros = [1, 2, 3, 4, 5]
# Tupla (tuple): Semelhante à lista, mas imutável, por exemplo: coordenadas = (10, 20)
# Dicionário (dict): Coleção de pares chave-valor, por exemplo: pessoa = {"nome": "João", "idade": 25}

if(idade>=18):
    print("Maior de idade")
else:
    print("Menor de idade")

if casado: #Se casado for true
    print(nome+" é casado")
else: #Senao
    print(nome+" não é casado")