import java.util.Scanner; //importando Scanner
public class Main {
    public static void main(String[] args) {
        Scanner teclado = new Scanner(System.in);

        System.out.println("Teorema de Pitagoras");
        System.out.println("x = a + b");

        System.out.print("Digite o valor de a: ");
        double a = teclado.nextInt();//Captando valor de A
        System.out.print("Digite o valor de b: ");
        double b = teclado.nextInt(); // Captando valor de b

        a = a * a; // multiplicando a x a
        b = b * b; //multiplicando b x b

        double x = a + b; // somando valor de a+b e jogando o valor para anova variavel x
        double raiz = Math.sqrt(x); //raiz quadrada de x e transformando o valor em int

        System.out.printf("O valor de x eh: %.2f", raiz);

    }
}
