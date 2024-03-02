package classe.aula1;
public class Main {
    public static void main(String[] args) {

        Produto produto1 = new Produto();

        Produto produto2 = new Produto("Notebook", 3699, 0.15);
        System.out.println(produto2.toString());
    }
}
