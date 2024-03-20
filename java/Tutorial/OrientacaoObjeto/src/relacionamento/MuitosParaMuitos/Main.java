package relacionamento.MuitosParaMuitos;

public class Main {
    public static void main(String[] args) {

        Aluno aluno1 = new Aluno("Joao");
        Aluno aluno2 = new Aluno("Pedro");
        Aluno aluno3 = new Aluno("Ana");

        Curso curso1 = new Curso();
        curso1.setNome("Curso de Java profissional");

        curso1.adicionarAluno(aluno1);
        curso1.adicionarAluno(aluno2);
        curso1.adicionarAluno(aluno3);

        curso1.listarAlunos();

        System.out.println(" ");

        System.out.println(curso1);
//        System.out.println(curso1);

    }
}
