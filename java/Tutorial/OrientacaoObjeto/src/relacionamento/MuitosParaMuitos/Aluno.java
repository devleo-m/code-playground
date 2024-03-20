package relacionamento.MuitosParaMuitos;

import java.util.ArrayList;

public class Aluno {
    private String nome;
    final ArrayList<Curso> cursos = new ArrayList<>();
    @Override
    public String toString() {
        return getNome();
    }

    public void adicionarCurso(Curso curso){
        this.cursos.add(curso);
        curso.alunos.add(this);
    }
    public Aluno(String nome) {
        this.nome = nome;
    }
    public String getNome() {
        return nome;
    }
    public void setNome(String nome) {
        this.nome = nome;
    }
}
