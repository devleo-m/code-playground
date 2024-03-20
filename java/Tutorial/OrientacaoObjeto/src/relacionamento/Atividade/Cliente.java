package relacionamento.Atividade;

import java.util.ArrayList;

public class Cliente {
    private String nome;
    Compra compra;
    ArrayList<Compra> compras = new ArrayList<>();
    public double valotTotal(){

        double total = 0;
        for (Compra compra : compras){
            for (Item item : compra.getItens()){
                total += item.getQuantidades() * item.getProdutos().getPreco();
            }
        }
        return total;
    }

    public void adicionarCompra(Compra compra){
        this.compras.add(compra);
    }

    public Cliente(String nome) {
        this.nome = nome;
    }

    @Override
    public String toString() {
        return "Cliente{" +
                "nome='" + nome + '\'' +
                ", compras=" + compras +
                '}';
    }

    public String getNome() {
        return nome;
    }

    public void setNome(String nome) {
        this.nome = nome;
    }

    public ArrayList<Compra> getCompras() {
        return compras;
    }

    public void setCompras(ArrayList<Compra> compras) {
        this.compras = compras;
    }
}
