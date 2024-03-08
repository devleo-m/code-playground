package relacionamento.UmParaUm;

import java.util.Scanner;

public class Carro {
    Motor motor;

    public Carro() {
        this.motor = new Motor(this);
    }

    public void ligarCarro(){
        if (!motor.ligado){
            motor.ligado = true;
            System.out.println("Carro ligado");
        } else {
            System.out.println("O carro ja esta ligado!!");
        }
    }
    public void deslicarCarro(){
        if (motor.ligado){
            motor.ligado = false;
            motor.velocidade = 0;
            System.out.println("Carro desligado");
        } else {
            System.out.println("O carro ja esta desligado!!");
        }
    }

    public void acelerar(){
        if (motor.ligado){
            Scanner scanner = new Scanner(System.in);
            int opcao;

            do {
                System.out.println("Digite 1 para acelerar 30km");
                System.out.println("--------------------------------");
                System.out.println("Digite 0 para freiar!!");
                System.out.println("--------------------------------");
                System.out.println("Digite 9 para desligar o carro! ");
                System.out.println("Velocidade: " + motor.velocidade);

                opcao = scanner.nextInt();

                switch (opcao) {
                    case 1:
                        motor.acelerarCarro();
                        break;
                    case 0:
                        motor.freiar();
                        break;
                    case 9:
                        deslicarCarro();
                        System.out.println("Velocidade: " + motor.velocidade);
                        break;
                    default:
                        System.out.println("Você digitou um valor inválido...");
                }
            } while (opcao != 9); // Continue enquanto a opção não for 9
        } else {
            System.out.println("O carro está desligado, ligue-o para acelerar...");
        }
    }

}
