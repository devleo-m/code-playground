import java.util.Scanner;

public class Main {
    public static void main(String[] args) {
        Scanner teclado = new Scanner(System.in);

        //ATIVIDADEEE 1:::::::

        String [] nomes = {"Joao", "Renata", "Maria", "Dayane", "Elisabete", "Gustavo"};
        for (int i = 0; i < nomes.length; i++){
            System.out.println(i +" "+ nomes[i]);
        }

        String[] meses = {"Janeiro", "Fevereiro", "Marco", "Abril",
                "Maio", "Junho", "Julho", "Agosto",
                "Setembro", "Outubro", "Novembro", "Sezembro"
        };
        System.out.println("Fim atividade 1!!!");
        //ATIVIDADEEEE 2::::::::
        System.out.println("Digite o mes que voce nasceu: ");
        int opcao = teclado.nextInt();
        switch (opcao){
            case 1:
                System.out.println("Voce nasceu no mes de "+meses[0]);
                break;
            case 2:
                System.out.println("Voce nasceu no mes de "+meses[1]);
                break;
            case 3:
                System.out.println("Voce nasceu no mes de "+meses[2]);
                break;
            case 4:
                System.out.println("Voce nasceu no mes de "+meses[3]);
                break;
            case 5:
                System.out.println("Voce nasceu no mes de "+meses[4]);
                break;
            case 6:
                System.out.println("Voce nasceu no mes de "+meses[5]);
                break;
            case 7:
                System.out.println("Voce nasceu no mes de "+meses[6]);
                break;
            case 8:
                System.out.println("Voce nasceu no mes de "+meses[7]);
                break;
            case 9:
                System.out.println("Voce nasceu no mes de "+meses[8]);
                break;
            case 10:
                System.out.println("Voce nasceu no mes de "+meses[9]);
                break;
            case 11:
                System.out.println("Voce nasceu no mes de "+meses[10]);
                break;
            case 12:
                System.out.println("Voce nasceu no mes de "+meses[11]);
                break;
        }
        System.out.println("Fim atividade 2!!!");
        //ATIVIDADEEEEE 3::::::
        double[] nota = new double[4];
        double recuperacao = 0;
        double media = 0;

        System.out.println("Primeira nota: ");
        nota[0] = teclado.nextDouble();
        System.out.println("Segunda nota: ");
        nota[1] = teclado.nextDouble();
        System.out.println("Terceira nota: ");
        nota[2] = teclado.nextDouble();
        System.out.println("Quarta nota: ");
        nota[3] = teclado.nextDouble();

        media = (nota[0]+nota[1]+nota[2]+nota[3])/4;

        if (media > 7){
            System.out.println("Aprovado por média");
            System.out.println("nota: "+media);
        } else {
            System.out.println("nota: "+media);
            System.out.println("Qual a nota da recuperacao? ");
            recuperacao = teclado.nextDouble();

            media = (recuperacao + media)/2;
            if (media > 10){
                System.out.println("Aprovado na recuperação");
                System.out.println("nota: "+media);
            } else {
                System.out.println("Reprovado");
                System.out.println("nota: "+media);
            }
        }
        System.out.println("Fim atividade 3!!!");

        //ATIVIDADEEEEEEEE 4:::::
        System.out.println("Qual é o tempo que a bomba deve estourar? ");
        int timeBomb = teclado.nextInt();

        if (timeBomb >= 0) {
            for (int i = timeBomb; i >= 0; i--) {
                System.out.println(i);
                if (i == 0) {
                    System.out.println("Bum!");
                }
            }
        } else {
            System.out.println("Valor inválido!");
            System.out.println("tente novamente");
        }
        System.out.println("FIM ATIVIDADE 4");
    }
}
