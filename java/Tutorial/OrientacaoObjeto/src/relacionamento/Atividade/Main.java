package relacionamento.Atividade;

import java.util.ArrayList;

public class Main {
    public static void main(String[] args) {
        Produto produto = new Produto("Cadeira", 2115.99);

        Item item = new Item(2, produto);

        Compra compra = new Compra();

        compra.adicionarItem(item);

        Cliente cliente = new Cliente("Gabriel");
        cliente.adicionarCompra(compra);

        System.out.println(cliente);

        System.out.println(cliente.valotTotal());
        
    }
}
