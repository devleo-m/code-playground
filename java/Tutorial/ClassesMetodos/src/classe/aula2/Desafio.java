package classe.aula2;

public class Desafio {
    private String produto;
    private double preco;
    private static double desconto = 0.10;

    public Desafio(String produto, double preco) {
        this.produto = produto;
        this.preco = preco;
    }
    public Desafio(String produto, double preco, double desconto) {
        this.produto = produto;
        this.preco = preco;
        this.desconto = desconto;
    }

    @Override
    public String toString() {
        return "Desafio{" +
                "produto='" + produto + '\'' +
                ", preco=" + preco +
                ", desconto=" + preco * desconto +
                '}';
    }

    public void aplicarDesconto(){
        preco -= (preco * desconto);
    }

    public String getProduto() {
        return produto;
    }

    public void setProduto(String produto) {
        this.produto = produto;
    }

    public double getPreco() {
        return preco;
    }

    public void setPreco(double preco) {
        this.preco = preco;
    }

    public static double getDesconto() {
        return desconto;
    }

    public static void setDesconto(double desconto) {
        Desafio.desconto = desconto;
    }
}
