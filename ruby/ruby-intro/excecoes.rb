#Entenda como lidar com erros usando **`begin`**, **`rescue`**, **`ensure`** e **`raise`**
########################### Tratamento de Exceções ###########################
def dividir_numeros1(num1, num2)
    begin
      resultado = num1 / num2
    rescue ZeroDivisionError => e
      puts "Erro ao dividir por zero: #{e}"
    rescue StandardError => e
      puts "Erro ao dividir os números: #{e}"
    ensure
      puts "Esta parte sempre será executada."
    end
    resultado
  end
  
  puts dividir_numeros1(10, 2) # Saída: 5
  puts dividir_numeros1(10, 0) # Saída: Erro ao dividir por zero: divided by 0

  ########################### Tratamento de Exceções ###########################

  def dividir_numeros(num1, num2)
    begin
      raise 'O denominador não pode ser zero.' if num2 == 0
      resultado = num1 / num2
    rescue ZeroDivisionError => e
      puts "Erro ao dividir por zero: #{e}"
    rescue StandardError => e
      puts "Erro ao dividir os números: #{e}"
    ensure
      puts "Esta parte sempre será executada."
    end
    resultado
  end
  
  puts dividir_numeros(10, 2) # Saída: 5
  puts dividir_numeros(10, 0) # Saída: Erro ao dividir por zero: O denominador não pode ser zero.
  
  