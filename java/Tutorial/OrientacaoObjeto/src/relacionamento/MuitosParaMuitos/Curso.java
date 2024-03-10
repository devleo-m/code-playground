package relacionamento.MuitosParaMuitos;

import java.util.ArrayList;

public class Curso {
    private String nome;
    final ArrayList<Aluno> alunos = new ArrayList<>();
    @Override
    public String toString() {
        return "Curso{" +
                "nome='" + nome + '\'' +
                ", alunos=" + alunos;
    }

    public void adicionarAluno(Aluno aluno){
        this.alunos.add(aluno);
        aluno.cursos.add(this);
    }
    public void listarAlunos(){
        System.out.println(alunos.size()+" alunos cadastrados: ");
        for (Aluno key : alunos){
            System.out.println(key);
        }
    }

    public String getNome() {
        return nome;
    }

    public void setNome(String nome) {
        this.nome = nome;
    }
}
