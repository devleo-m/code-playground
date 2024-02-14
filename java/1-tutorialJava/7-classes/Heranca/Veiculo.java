public class Veiculo {
    String Marca;
    int ano;

    public void ligar(){
        System.out.println("Carro ligado!");
    }

    public void desligar(){
        System.out.println("Carro desligado!");
    }

}

class Carro extends Veiculo {
    public String Modelo;
    int numPortas;

    public void abrirPortas() {
        System.out.println("Abrir Portas do Carro");
    }
}
