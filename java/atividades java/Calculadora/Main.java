import java.util.Scanner;

public class Main {
    public static void main(String[] args){
        Scanner scanner = new Scanner(System.in);
        int calc = 0;
        boolean calcSucesso = false;

        System.out.println("Primeiro numero: ");
        int number1 = scanner.nextInt();
        System.out.println("Digite o segundo numbero: ");
        int number2 = scanner.nextInt();

        System.out.println("Digite o operador que voce quer utilizar para calcular");
        System.out.println("Adicao: +");
        System.out.println("Subtracao: -");
        System.out.println("Multiplicacao: *");
        System.out.println("Divisao: /");
        System.out.println("Resto: %");

        String operador = scanner.next();

        switch (operador){
            case "+":
               calc = number1 + number2 ;
               calcSucesso = true;
               break;
            case "-":
                calc = number1 - number2 ;
                calcSucesso = true;
                break;
            case "*":
                calc = number1 * number2 ;
                calcSucesso = true;
                break;
            case "/":
                calc = number1 / number2 ;
                calcSucesso = true;
                break;
            case "%":
                calc = number1 % number2 ;
                calcSucesso = true;
                break;
            default:
                System.out.println("Operador invalido, tente novamente");
                System.out.println("Voce deve digitar + - * / %");
                System.out.println("Para fazer o calculo! :)");
                calcSucesso = false;
                break;
        }
        if (calcSucesso){
            System.out.println("Total = "+ calc);
        }else {
            System.out.println(":(");
        }

    }
}
