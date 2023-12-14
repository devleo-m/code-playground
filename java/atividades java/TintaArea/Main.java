import java.util.Scanner;
public class Main {
    public static void main(String[] args){
        Scanner teclado = new Scanner(System.in);
        System.out.println("Digite seu nome: ");
        String cliente = teclado.next();

        System.out.println("Quantos metros quadrados tem a sua casa? ");
        double m2 = teclado.nextInt();

        double rendimentoTinta = 10;
        double tinta = m2/rendimentoTinta;
        tinta = Math.round(tinta);

        double galaoDeTinta = 3.6;
        double quantidadeDeGalao = tinta/galaoDeTinta;
        quantidadeDeGalao = Math.round(quantidadeDeGalao);

        int laco = 2;
        double marca = 0;
        int pintor = 250; //valor do servico do pintor

        do {
            System.out.println("Qual marca voce prefere para pintar sua casa? Numero 1: Marca1: 360 reais numero 2: Marca2: 682. Esse preco e por galao");
            int opcao = teclado.nextInt();

            if (opcao == 1) {
                System.out.println("Voce escolheu a marca1, cada galao custa 360 reais");
                marca = 360;
                laco = 3;
            } else if (opcao == 2) {
                System.out.println("Voce escolheu a marca1, cada galao custa 360 reais");
                marca = 682;
                laco = 3;
            } else {
                System.out.println("voce digitou um valor invalido, voce deve digitar apenas o numero 1 se quiser escolher a marca1 ou digitar o numero 2, para escolher a marca2");
                laco = 2;
            }
        }while (laco == 2);


        if (m2 >= 300){
            pintor = 400;
        } else if (m2 >= 250) {
            pintor = 250;
        } else {
            pintor = 200;
        }

        double servicoPintor = quantidadeDeGalao * pintor;
        servicoPintor = Math.round(servicoPintor);

        int totalTinta = (int) (quantidadeDeGalao * marca);
        double total = servicoPintor + totalTinta;

        System.out.println("---------------------------------------------------");
        System.out.printf("Caro cliente "+cliente+" seu orcamento: ");
        System.out.println("Marca da tinta: "+marca);
        System.out.println("quantidade necessaria para o servico "+quantidadeDeGalao);
        System.out.println("Preco total das tintas: "+totalTinta);
        System.out.println("Servico do pintor: "+servicoPintor);
        System.out.printf("TOTAL: %.2f", total);
        System.out.println("---------------------------------------------------");

        System.out.println("Voce deseja pagar de qual forma? PIX, DEBITO OU CARTAO?");
        System.out.println("Parcelamos em ate 12x com juros pelo cartao de credito");
        System.out.println("PIX, BOLETO OU DINHEIRO! voce ganha 10% de desconto!");
        System.out.println("Digite 1 para PIX, Dinheiro ou Boleto");
        System.out.println("Digite 2 para Cartao de credito");

        /*do {
            if (opcao = 1){
                if (opcao = 1){
                    System.out.println();
                }
            }
        } while (laco = 5);*/

    }
}
