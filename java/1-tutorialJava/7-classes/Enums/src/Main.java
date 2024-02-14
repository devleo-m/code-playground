public class Main {
    public static void main(String[] args) {
        //Declarar uma Variável Enum:
        DiaSemana hoje = DiaSemana.TERCA;

        //Comparar Enums:
        if (hoje == DiaSemana.SEGUNDA) {
            System.out.println("ATENÇÃO!!! Hoje é Segunda-feira.");
        }

        //Switch com Enums:
        switch (hoje) {
            case SEGUNDA:
                System.out.println("acabou o feriado! dia de trabalha");
                break;
            case TERCA:
                System.out.println("Terça-feira, dia para trabalhar muito");
                break;
            case QUARTA:
                System.out.println("quarta-feira, mais um dia p trabalhar mt");
                break;
            case QUINTA:
                System.out.println("quinta-feira, trabalhar mt");
                break;
            case SEXTA:
                System.out.println("SEXTOUUUUUU!!! UHUHUUUUU");
                break;
            case SABADO:
                System.out.println("SABADO, dia para relaxar");
                break;
            case DOMINGO:
                System.out.println("Domingo, dia para relaxar e se preparar para segunda");
                break;
            default:
                System.out.println("Dia invalido.");
        }

        //Iterar sobre Enums:
        for (DiaSemana dia: DiaSemana.values()){
            System.out.println("Iterar: "+dia);
        }

    }
}
