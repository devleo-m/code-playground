package relacionamento.UmParaMuitos;

public class Main {
    public static void main(String[] args) {
        Compra compra1 = new Compra();
        compra1.nomeCliente = "Fulano de tal";
        compra1.itens.add(new Item("Caneta", 2, 5.89));
        compra1.itens.add(new Item("Borracha", 1, 2.68));
        compra1.itens.add(new Item("Caderno", 2, 27.79));

        System.out.println(compra1.valorTotal());
    }
}
