package relacionamento.Atividade;

import java.util.ArrayList;

public class Compra {
    ArrayList<Item> itens = new ArrayList<>();
    Item item;

    public void adicionarItem(Item item){
        this.itens.add(item);
    }
    @Override
    public String toString() {
        return "Compra{" +
                "itens=" + itens +
                '}';
    }

    public ArrayList<Item> getItens() {
        return itens;
    }

    public void setItens(ArrayList<Item> itens) {
        this.itens = itens;
    }
}
