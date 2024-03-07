package pilha;

import java.util.ArrayDeque;
import java.util.Deque;

public class Aula {
    public static void main(String[] args) {
        Deque<String> livros = new ArrayDeque<>();

        livros.push("O Pequeno Principe");
        livros.push("Don Quixote");
        livros.push("O Menino do Pijama Listrado");
        livros.add("O Hobbit");

        //System.out.println(livros.peek());
        System.out.println(livros.element()); //Mostra o ultimo colocado no livro

        System.out.println(livros.pop());
        System.out.println(livros.poll());
        System.out.println(livros.poll());
        System.out.println(livros.poll());
        System.out.println(livros.poll());
    }
}
