import java.util.Scanner;
public class Main {
    public static void main(String[] args){
        Scanner teclado = new Scanner(System.in);
        System.out.print("Digite seu nome: ");
        String cliente = teclado.next();

        System.out.print("Quantos metros quadrados tem a sua casa? ");
        double m2 = teclado.nextInt();

        double rendimentoTinta = 10;
        double tinta = m2/rendimentoTinta;
        tinta = Math.round(tinta);

        double galaoDeTinta = 3.6;
        double quantidadeDeGalao = tinta/galaoDeTinta;
        quantidadeDeGalao = Math.round(quantidadeDeGalao);

        boolean temEmail = false; //Se cliente comprar em boleto vai ter email
        String email = "";

        int laco = 2;
        double marca = 0;
        int pintor; //valor do servico do pintor

        do {
            System.out.println("Qual marca voce prefere para pintar sua casa? ");
            System.out.println("Digite um numero de 1 a 5 para escolher ");

            System.out.println("1: ColorPrime 3.6L   R$ 89,99");
            System.out.println("2: RainbowCoat 3.6L  R$ 98,50");
            System.out.println("3: VibrantBlend 3.6L R$ 127,75");
            System.out.println("4: ChromaLux 3.6L    R$ 187,25");
            System.out.println("5: TonalTint 3.6L    R$ 312,90");

            int opcao = teclado.nextInt();

            switch (opcao){
                case 1:
                    System.out.println("Voce escolheu a ColorPrime, cada galao custa R$ 89,99");
                    marca = 89.99;
                    laco = 3;
                    break;
                case 2:
                    System.out.println("Voce escolheu a RainbowCoat, cada galao custa R$ 98,50");
                    marca = 98.50;
                    laco = 3;
                    break;
                case 3:
                    System.out.println("Voce escolheu a VibrantBlend, cada galao custa R$ 127,75");
                    marca = 127.75;
                    laco = 3;
                    break;
                case 4:
                    System.out.println("Voce escolheu a ChromaLux, cada galao custa R$ 187,25");
                    marca = 187.25;
                    laco = 3;
                    break;
                case 5:
                    System.out.println("Voce escolheu a ChromaLux, cada galao custa R$ 312,90");
                    marca = 312.90;
                    laco = 3;
                    break;
                default:
                    System.out.println("Opção inválida");
                    System.out.println("Voce deve digitar entre o numero 1 e 5 para escolher a tinta");
                    laco = 2;
            }

        }while (laco == 2);


        if (m2 >= 300){
            pintor = 400;
        } else if (m2 >= 250) {
            pintor = 250;
        } else if (m2 >= 150){
            pintor = 180;
        } else if (m2 >= 100) {
            pintor = 150;
        } else {
            pintor = 100;
        }

        double servicoPintor = quantidadeDeGalao * pintor;
        servicoPintor = Math.round(servicoPintor);

        int totalTinta = (int) (quantidadeDeGalao * marca);
        double total = servicoPintor + totalTinta;

        System.out.println("---------------------------------------------------");
        System.out.println("Preco total das tintas: "+totalTinta);
        System.out.println("Servico do pintor: "+servicoPintor);
        System.out.printf("TOTAL: %.2f", total);
        System.out.println(" ");
        System.out.println("---------------------------------------------------");

        double desconto = 0;
        int juros = 0;

        do {
            System.out.println("Forma de pagamento: PIX, Boleto, Dinheiro ou CARTAO");
            System.out.println("Parcelamos em ate 12x com juros pelo cartao de credito");
            System.out.println("PIX, BOLETO OU DINHEIRO! voce ganha em ate 15% de desconto!");
            System.out.println("Digite 1 para PIX, Dinheiro ou Boleto");
            System.out.println("Digite 2 para Cartao de credito");
            System.out.println("---------------------------------------------------");
            int opcao = teclado.nextInt();
            if (opcao == 1){
                System.out.println("Digite 1 para PIX");
                System.out.println("Digite 2 para Dinheiro");
                System.out.println("Digite 3 Boleto");
                int opcao2 = teclado.nextInt();

                if (opcao2 == 1) {
                    System.out.println("PAGAMENTO COM PIX!");
                    System.out.println("Pague com essa chave pix");
                    System.out.println("Codigo: xxxxxxxx-xxx");
                    //fazer o calculo
                    desconto = 15;
                    desconto = (total * desconto) / 100;
                    total = total - desconto;
                    laco = 5;

                } else if (opcao2 ==2) {
                    System.out.println("PAGAMENTO COM DINHEIRO");
                    desconto = 10;
                    desconto = (total * desconto) / 100;
                    total = total - desconto;
                    laco = 5;

                } else if (opcao2 == 3) {
                    System.out.println("PAGAMENTO EM BOLETO");
                    System.out.println("Sera enviado o boleto para o seu email");
                    System.out.println("Digite abaixo o endereco de seu email");
                    System.out.print("E-mail: ");
                    email = teclado.next();

                    desconto = 10;
                    desconto = (total * desconto) / 100;
                    total = total - desconto;
                    laco = 5;
                    temEmail = true;

                } else {
                    System.out.println("Algo deu errado, voce deve digitar entre");
                    System.out.println("1 para pix, 2 para dinheiro e 3 para boleto");
                    laco = 4;
                }
            } else if (opcao == 2) {
                System.out.println("Pagamento com Cartao de credito");
                System.out.println("Deseja parcelar de quantas vezes?");
                System.out.println("1x sem juros");
                System.out.println("2x com juros de 10%");
                System.out.println("3x com juros de 15%");
                System.out.println("4x com juros de 20%");
                System.out.println("5x com juros de 25%");
                System.out.println("6x com juros de 30%");
                System.out.println("7x com juros de 35%");
                System.out.println("8x com juros de 40%");
                System.out.println("9x com juros de 45%");
                System.out.println("10x com juros de 50%");
                System.out.println("11x com juros de 55%");
                System.out.println("12x com juros de 60%");
                System.out.println("--------------------");
                System.out.println("Digite quantas vezes voce deseja parcelar: ");
                juros = teclado.nextInt();

                switch (juros){
                    case 1:
                        System.out.println("sem juros");
                        juros = 0;
                        total = total * juros;
                        break;
                    case 2:
                        System.out.println("10% de juros");
                        juros = 10;
                        total = total * juros;
                        break;
                    case 3:
                        System.out.println("15% de juros");
                        juros = 15;
                        total = total * juros;
                        break;
                    case 4:
                        System.out.println("20% de juros");
                        juros = 20;
                        total = total * juros;
                        break;
                    case 5:
                        System.out.println("25% de juros");
                        juros = 25;
                        total = total * juros;
                        break;
                    case 6:
                        System.out.println("30% de juros");
                        juros = 30;
                        total = total * juros;
                        break;
                    case 7:
                        System.out.println("35% de juros");
                        juros = 35;
                        total = total * juros;
                        break;
                    case 8:
                        System.out.println("40% de juros");
                        juros = 40;
                        total = total * juros;
                        break;
                    case 9:
                        System.out.println("45% de juross");
                        juros = 45;
                        total = total * juros;
                        break;
                    case 10:
                        System.out.println("50% de juros");
                        juros = 50;
                        total = total * juros;
                        break;
                    case 11:
                        System.out.println("55% de juros");
                        juros = 55;
                        total = total * juros;
                        break;
                    case 12:
                        System.out.println("60% de juros");
                        juros = 60;
                        total = total * juros;
                        break;
                }
                laco = 5;
            } else {
                System.out.println("Numero invalido");
                laco = 4;
            }
        } while (laco == 4);

        System.out.println("-----------------------------------------------");
        System.out.println("Cliente "+cliente);
        System.out.println("Preco da tinta: R$"+marca);
        System.out.println("quantidade de tinta para servico: "+(int)quantidadeDeGalao);
        System.out.println("Preco das tintas: R$"+totalTinta);
        System.out.println("Servico do pintor: R$"+(int)servicoPintor);
        System.out.println("Desconto: R$"+desconto);
        System.out.println("Juros: "+juros+"%");
        if (temEmail == true){
            System.out.println("Seu Boleto sera enviado para seu email: "+email);
            System.out.println("OBS:. Voce tem 3 dias uteis para pagar!");
        }
        System.out.println("Total R$"+total);
        System.out.println("-----------------------------------------------");
    }
}
