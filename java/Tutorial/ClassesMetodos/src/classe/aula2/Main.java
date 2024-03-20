package classe.aula2;

public class Main {
    public static void main(String[] args) {
        Desafio desafio = new Desafio("Notebook", 1000, 0.5);
        desafio.aplicarDesconto();
        String allProduct = String.valueOf(desafio);
        System.out.println(allProduct);
    }
}
