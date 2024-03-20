package relacionamento.Atividade;

public class Item {
    private int quantidades;
    Produto produtos;

    public Item(int quantidades, Produto produtos) {
        this.quantidades = quantidades;
        this.produtos = produtos;
    }

    @Override
    public String toString() {
        return "Item{" +
                "quantidades=" + quantidades +
                ", produtos=" + produtos +
                '}';
    }

    public int getQuantidades() {
        return quantidades;
    }

    public void setQuantidades(int quantidades) {
        this.quantidades = quantidades;
    }

    public Produto getProdutos() {
        return produtos;
    }

    public void setProdutos(Produto produtos) {
        this.produtos = produtos;
    }
}
