
public class Main {

    public static void main(String[] args){
        Aluno mario = new Aluno();
        mario.nome = "Mario";
        mario.cpf = "033.333.393-99";
        mario.idade = 15;
        mario.matricula = "3489054";
        mario.turma = 'B';

        Aluno luigi = new Aluno();
        luigi.nome = "Luigi";
        luigi.cpf = "634.874.247-76";
        luigi.idade = 14;
        luigi.matricula = "17548264";
        luigi.turma = 'B';

        // Exibindo informações dos alunos
        System.out.println("Informações do aluno Mario:");
        System.out.println("Nome: " + mario.nome);
        System.out.println("CPF: " + mario.cpf);
        System.out.println("Idade: " + mario.idade);
        System.out.println("Matrícula: " + mario.matricula);
        System.out.println("Turma: " + mario.turma);

        System.out.println("Informações do aluno Luigi:");
        System.out.println("Nome: " + luigi.nome);
        System.out.println("CPF: " + luigi.cpf);
        System.out.println("Idade: " + luigi.idade);
        System.out.println("Matrícula: " + luigi.matricula);
        System.out.println("Turma: " + luigi.turma);
    }
}
