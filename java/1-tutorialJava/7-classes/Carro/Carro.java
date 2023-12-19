public class Carro {
    String Modelo;
    String Cor;
    String ano;
    boolean ligado = false;

    public void ligarCarro(){
        ligado = true;
        if (ligado){
            System.out.println("Ligando carro!");
            for (int i= 0; i <= 5; i++){
                System.out.println(i);
            }
            System.out.println("Carro Ligado!");
            System.out.println("Acelere o carro para poder andar! :)");
            System.out.println("acelerar()");
        } else {
            System.out.println("Carro desligado!");
        }
    }

    public void desligarCarro(){
        ligado = false;
        System.out.println("Carro Desligado!");

    }

    public void acelerar(int velocidade){
        if (ligado) {
            if (velocidade >= 360) {
                System.out.println("Velocidade: " + velocidade + "Km/h ! MULTA 3.000 reais!");
            } else if (velocidade >= 111) {
                System.out.println("Velocidade: " + velocidade + "Km/h ! MULTA 1.322 reais!");
            } else if (velocidade >= 55) {
                System.out.println("Velocidade: " + velocidade + "Km/h");
            } else if (velocidade >= 1) {
                System.out.println("Velocidade: " + velocidade + "Km/h. MULTA 1500 reais. lento demais");
            } else if (velocidade == 0) {
                System.out.println("Voce precisa acelerar, voce esta parado kkkk");
            } else if (velocidade < 0) {
                System.out.println("Velocidade negativa?");
                System.out.println("O que voce esta tentando e re?");
                System.out.println("kkkkkkk");
                System.out.println("Multado por Ze Ruelagem kkkkkk");
            }
        }else {
            System.out.println("Voce precisa ligar o carro primeiro. Ze ru");
            System.out.println("ela");
        }
    }
}

