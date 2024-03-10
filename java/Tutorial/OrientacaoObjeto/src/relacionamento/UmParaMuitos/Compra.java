package relacionamento.UmParaMuitos;

import java.util.ArrayList;

public class Compra {

    Cliente cliente;
    ArrayList<Item> itens = new ArrayList<>();

    @Override
    public String toString() {
        return cliente +" "+itens;
    }
    public double getPrecoTotal(){
        double total = 0;
        for (Item key : itens){
            total += key.getQuantidade() * key.getPreco();
        }
        return total;
    }
}
