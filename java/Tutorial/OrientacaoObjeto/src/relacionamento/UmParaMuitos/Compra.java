package relacionamento.UmParaMuitos;

import java.util.ArrayList;

public class Compra {
    String nomeCliente;
    ArrayList<Item> itens = new ArrayList<>();
    public double valorTotal(){
        double precoTotal = 0;

        for (Item key : itens){
            precoTotal += key.quantidade * key.preco;
        }

        return precoTotal;
    }
}
