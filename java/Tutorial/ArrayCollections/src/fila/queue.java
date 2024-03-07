package fila;

import java.util.Collection;
import java.util.Iterator;
import java.util.LinkedList;
import java.util.Queue;

public class queue {
    public static void main(String[] args) {
        Queue<String> fila = new LinkedList<>();
        fila.add("Ana");
        fila.offer("Bia");
        fila.add("Carlos");
        fila.offer("Daniel");
        fila.add("Gui");

        System.out.println(fila.peek()); //retorna ana
        //System.out.println(fila.peek());//retorna ana
        System.out.println(fila.element());//retorna ana
        //System.out.println(fila.element());//retorna ana

        System.out.println("=============================");
//        fila.size();
//        fila.clear();
//        fila.isEmpty();

        System.out.println(fila.poll());
        System.out.println(fila.poll());
        System.out.println(fila.poll());
        System.out.println(fila.poll());
        System.out.println(fila.poll());
        System.out.println(fila.poll());

        //poll pula uma linha e segue com os elementos por sequencia sem repetir

    }
}
