package classe.pratica1;

public class Main {
    public static void main(String[] args) {
        Pessoa pessoa = new Pessoa("Dayane", 58);
        Comida food = new Comida("Podrao", 0.5);
        pessoa.comer(food);
        System.out.println(pessoa.toString());
    }
}
