package arrays;
// CRIANDO UM ARRAY
import java.util.Arrays;

public class aula1 {
    public static void main(String[] args) {
        double[] notasAlunosA = new double[4];

        //System.out.println(Arrays.toString(notasAlunosA));

        notasAlunosA[0] = 7.4;
        notasAlunosA[1] = 5.2;
        notasAlunosA[2] = 8.5;
        notasAlunosA[3] = 4.6;

        //System.out.println(Arrays.toString(notasAlunosA));
        //System.out.println(notasAlunosA[3]);

        double totalAlunoA = 0;
        for (int i = 0; i < notasAlunosA.length; i++) {
            totalAlunoA += notasAlunosA[i];
        }
        System.out.printf("Nota total: %.1f", totalAlunoA / notasAlunosA.length);

        //----------------------------------------------------------------

        double totalAlunoB = 0;
        double[] notasAlunoB = {5.3, 6.2, 3.7, 7.7};
        for (int i = 0; i < notasAlunoB.length; i++) {
            totalAlunoB += notasAlunoB[i];
        }
        System.out.printf("Nota total: %.1f", totalAlunoB / notasAlunoB.length);
    }
}
