package classe.aula1;
public class Produto {
    private String nome;
    private double preco;
    private double desconto;
    public Produto(){

    }
    public Produto(String nome, double preco, double desconto) {
        this.nome = nome;
        this.preco = preco;
        this.desconto = desconto;
    }
    @Override
    public String toString() {
        return "Produto{" +
                "nome='" + nome + '\'' +
                ", preco=" + preco +
                ", desconto=" + desconto +
                '}';
    }
    public void calcDesconto(){
        desconto *= preco;
        preco -= desconto;
    }
}
