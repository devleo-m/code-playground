package relacionamento.UmParaMuitos;

public class Item {
    private String nomeItem;
    private int quantidade;
    private double preco;

    public Item(String nomeItem, int quantidade, double preco){
        this.nomeItem = nomeItem;
        this.quantidade = quantidade;
        this.preco = preco;
    }

    @Override
    public String toString() {
        return "Produto: " + nomeItem +
                ", quantidade :" + quantidade +
                ", preco :" + preco;
    }
    public int getQuantidade() {
        return quantidade;
    }

    public double getPreco() {
        return preco;
    }

}
