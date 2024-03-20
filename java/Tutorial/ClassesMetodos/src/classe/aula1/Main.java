package classe.aula1;
public class Main {
    public static void main(String[] args) {

        Produto produto1 = new Produto();

        Data dia = new Data();
        dia.dia = 8;
        dia.mes = 10;
        dia.ano = 1996;

        System.out.println(dia.diaCompleto());

        Produto produto2 = new Produto("Notebook", 1000, 0.10);
        produto2.calcDesconto();
        System.out.println(produto2.toString());
    }
}
