########################### Classes e Objetos ###########################
class Carro
    def initialize(modelo, ano)
      @modelo = modelo
      @ano = ano
    end
  
    def exibir_detalhes
      puts "Modelo: #{@modelo}, Ano: #{@ano}"
    end
  end
  
  meu_carro = Carro.new('Tesla Model S', 2022)
  meu_carro.exibir_detalhes

############################ Herança e Polimorfismo ###########################
class Veiculo
    def mover
      puts 'Veículo está se movendo'
    end
  end
  
  class Carro < Veiculo
    def mover
      puts 'Carro está se movendo'
    end
  end
  
  class Moto < Veiculo
    def mover
      puts 'Moto está se movendo'
    end
  end
  
  veiculo = Veiculo.new
  veiculo.mover
  
  carro = Carro.new
  carro.mover
  
  moto = Moto.new
  moto.mover

########################### Módulos e Mixins ###########################
module Impressora
    def imprimir(texto)
      puts "Imprimindo: #{texto}"
    end
  end
  
  class Documento
    include Impressora
  end
  
  doc = Documento.new
  doc.imprimir('Olá, mundo!')

########################### Atributos e Métodos de Classe ###########################
class Computador
    @@quantidade = 0
  
    def self.quantidade
      @@quantidade
    end
  
    def initialize(marca)
      @marca = marca
      @@quantidade += 1
    end
  
    def self.exibir_info
      puts "Quantidade de computadores: #{@@quantidade}"
    end
  end
  
  Computador.new('Apple')
  Computador.new('Dell')
  puts Computador.quantidade # Saída: 2
  
  Computador.exibir_info # Saída: Quantidade de computadores: 2
  