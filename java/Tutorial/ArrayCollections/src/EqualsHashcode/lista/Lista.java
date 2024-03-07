package EqualsHashcode.lista;

import java.util.ArrayList;

public class Lista {
    public static void main(String[] args) {
        ArrayList<Usuario> lista = new ArrayList<>();

        Usuario u1 = new Usuario("Ana");
        lista.add(u1);

        lista.add(new Usuario("Jose"));
        lista.add(new Usuario("Pedro"));
        lista.add(new Usuario("Lucas"));
        lista.add(new Usuario("Julia"));

        lista.remove(new Usuario("Ana")); //Observe que a ana foi removida da lista
        lista.remove(1); //Pedro foi removido da lista

        for (Usuario key : lista){
            System.out.println(key);
        }
    }
}
