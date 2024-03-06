package arrays;

import java.util.Arrays;
import java.util.Scanner;

public class atividadeArray {
    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        System.out.println("Digite o tamanho da array: ");
        int tamanhoArray = scanner.nextInt();
        double[] notasAluno = new double[tamanhoArray];
        for (int i = 0; i < notasAluno.length; i++) {
            System.out.println("Digite a "+ (i + 1) +" nota: ");
            notasAluno[i] += scanner.nextDouble();
        }
        double total = 0;
        System.out.println("Notas do aluno: " + Arrays.toString(notasAluno));
        for (double nota : notasAluno) {
            total += nota;
        }
        System.out.printf("Nota total do aluno: %.1f", total / notasAluno.length);
        scanner.close();
    }
}
