package relacionamento.UmParaUm;

import java.util.Scanner;

public class Motor {

    Carro carro;

    public Motor(Carro carro) {
        this.carro = carro;
    }
    double fatorInjecao = 1;
    double velocidade = 0;
    boolean ligado = false;

    public int giros(){
        if (!ligado){
            return 0;
        }else {
            return (int) Math.round(fatorInjecao * 3000) ;
        }
    }

    public void acelerarCarro(){
        if (velocidade <= 230 && velocidade>=0){
            velocidade += 30;
            System.out.println("Velocidade: " + velocidade);
        }else {
            System.out.println("Velocidade maxima ate de 2410km");
            System.out.println("Esse carro eh 1.0");
        }

    }
    public void freiar(){
        if (velocidade >= 0){
            velocidade -= 20;
            System.out.println("Velocidade: " + velocidade);
        } else if (velocidade <= 0){
            velocidade = 0;
            System.out.println("O carro ja esta parado!");
        }

    }

}
