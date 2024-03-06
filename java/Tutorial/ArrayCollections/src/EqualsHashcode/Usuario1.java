package EqualsHashcode;

import java.util.Objects;

public class Usuario1 {
    private String nome;
    private String email;

    public Usuario1(String nome, String email) {
        this.nome = nome;
        this.email = email;
    }

    @Override
    public String toString() {
        return "Usuario1{" +
                "nome='" + nome + '\'' +
                ", email='" + email + '\'' +
                '}';
    }

    @Override
    public boolean equals(Object obj) {
        boolean nomeIgual = ((Usuario1) obj).nome.equals(this.nome);
        boolean nomeEmail = ((Usuario1) obj).email.equals(this.email);
        return nomeIgual && nomeEmail;
    }

}
