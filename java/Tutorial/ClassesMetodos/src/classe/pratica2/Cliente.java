package classe.pratica2;

import java.lang.reflect.Array;
import java.util.ArrayList;

public class Cliente {
    private String nome;
    private ArrayList<Livro> livroEmprestadoCliente;

    public Cliente(String nome) {
        this.nome = nome;
    }

    public Cliente(ArrayList<Livro> livroEmprestadoCliente) {
        this.livroEmprestadoCliente = livroEmprestadoCliente;
    }

    @Override
    public String toString() {
        return "Cliente{" +
                "nome='" + nome + '\'' +
                ", livroEmprestadoCliente=" + livroEmprestadoCliente +
                '}';
    }

    public String getNome() {
        return nome;
    }

    public void setNome(String nome) {
        this.nome = nome;
    }

    public ArrayList<Livro> getLivroEmprestadoCliente() {
        return livroEmprestadoCliente;
    }

    public void setLivroEmprestadoCliente(ArrayList<Livro> livroEmprestadoCliente) {
        this.livroEmprestadoCliente = livroEmprestadoCliente;
    }
}
