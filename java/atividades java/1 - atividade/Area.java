import java.util.Scanner;
public class Area {
    public static void main(String[] args) {
        Scanner teclado = new Scanner(System.in);
        System.out.print("Digite a largura: ");
        double largura = teclado.nextDouble();
        System.out.print("Digite o comprimento: ");
        double comprimento = teclado.nextDouble();

        double area = largura * comprimento;

        System.out.printf("Valor da area eh: %.2f", area);
    }
}
