package classe.pratica2;

import java.util.ArrayList;
import java.util.Scanner;

public class Main {
    static ArrayList<Livro> livros = new ArrayList<>();
    static ArrayList<Cliente> clientes = new ArrayList<>();

    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);

        clientes.add(new Cliente("Fulano"));
        clientes.add(new Cliente("Beltrano"));
        clientes.add(new Cliente("Ciclano"));

        livros.add(new Livro("A Rapoza que come pão","Joaozinho", 1990));

        do {
            System.out.println("Bem vindo a biblioteca");
            System.out.println("Deseja (V)isualizar, (A)dicionar, (E)mprestar (R)emover ou (S)air?");
            String opcao = scanner.next();
            switch (opcao){
                case "v":
                    visualizarLivros();
                    break;
                case "a":
                    adicionarLivro();
                    break;
                case "e":
                    emprestarLivro();
                    break;
                case "r":
                    System.out.println("remover"); //deletar um livro especifico
                    break;
                case "s":
                    System.out.println("Sair do programa...");
                    return;
                default:
                    System.out.println("Valor invalido, tente novamente!");
            }

        }while (true);
    }

    public static void visualizarLivros(){
        System.out.println("Livros cadastrados: ");
        for (int i = 0; i < livros.size(); i++) {
            System.out.println(i + ": "+ livros.get(i));
        }
    }

    public static void adicionarLivro(){
        Scanner scanner = new Scanner(System.in);
        System.out.println("Digite o nome do titulo: ");
        String titulo = scanner.nextLine();
        System.out.println("Digite o nome do autor: ");
        String autor = scanner.nextLine();
        System.out.println("Digite o ano da publicacao: ");
        int ano = scanner.nextInt();
        livros.add(new Livro(titulo, autor, ano));
        System.out.println("Livro adicionado com sucesso!");
    }

    public static void emprestarLivro(){
        Scanner scanner = new Scanner(System.in);
        visualizarLivros();
        System.out.println("Voce deseja emprestar qual livro? ");
        int indiceLivro = scanner.nextInt();
        if (indiceLivro >=0 && indiceLivro < livros.size()){
            livros.get(indiceLivro);
            System.out.println("Digite o indice do cliente para emprestar a ele: ");
            int indiceCliente = scanner.nextInt();
            if (indiceCliente >= 0 && indiceCliente < clientes.size()){
                clientes.add(livros.get(indiceLivro));
            }
        }else {
            System.out.println("Livro nao existe ou nao encontrado!");
        }

    }
}
