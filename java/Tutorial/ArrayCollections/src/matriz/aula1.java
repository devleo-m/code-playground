package matriz;
import java.util.Arrays;
import java.util.Scanner;
public class aula1 {
    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        System.out.println("Digite a quantidade de alunos: ");
        int qntAlunos = scanner.nextInt();
        System.out.println("Digite a quantidade de notas por aluno: ");
        int qntNotas = scanner.nextInt();
        double[][] notasTurma = new double[qntAlunos][qntNotas];
        double total = 0;
        for (int a = 0; a < notasTurma.length; a++) {
            for (int n = 0; n < notasTurma[a].length; n++) {
                System.out.println("Digite a nota do aluno: "+ (a + 1) +" nota: "+(n + 1));
                notasTurma[a][n] += scanner.nextDouble();
                System.out.println("Nota adicionada com sucesso!");
                total += notasTurma[a][n];
            }
            System.out.println("Notas do aluno "+ (a + 1) + " cadastrada com sucesso");
            System.out.println("---------------------------------------------------");
        }
        double mediaTurma = total / (qntAlunos * qntNotas);
        System.out.println("Nota de todos os alunos: ");
        for (double[] notasAluno : notasTurma){
            System.out.println(Arrays.toString(notasAluno));
        }
        System.out.printf("A media da turma eh: %.1f",mediaTurma);
        scanner.close();
    }
}
