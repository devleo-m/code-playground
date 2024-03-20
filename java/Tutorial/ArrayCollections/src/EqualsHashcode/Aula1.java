package EqualsHashcode;

public class Aula1 {
    public static void main(String[] args) {
        Usuario1 p1 = new Usuario1("Fulano","fulano@gmail.com");
        Usuario1 p2 = new Usuario1("Fulano","fulano@gmail.com");
        Usuario1 p3 = new Usuario1("Ciclano","ciclano@gmail.com");

        System.out.println(p1 == p2);
        System.out.println(p1.equals(p2));
    }
}
