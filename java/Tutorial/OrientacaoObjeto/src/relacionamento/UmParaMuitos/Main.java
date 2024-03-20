package relacionamento.UmParaMuitos;

public class Main {
    public static void main(String[] args) {

        Cliente cliente1 = new Cliente("Leonardo", "002.001.004-99", "av fulano de tal - 40 - SP");

        Compra compra1 = new Compra();

        compra1.cliente = cliente1;

        compra1.itens.add(new Item("Mesa", 2, 265.88));
        compra1.itens.add(new Item("Caneta", 4, 2.99));
        compra1.itens.add(new Item("Caderno", 2, 37.78));

        System.out.println("R$"+compra1.getPrecoTotal());
        System.out.println(compra1);

    }
}
