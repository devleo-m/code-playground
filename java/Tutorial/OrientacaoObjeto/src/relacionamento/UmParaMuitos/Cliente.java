package relacionamento.UmParaMuitos;

public class Cliente {
    private String nomeCliente;
    private String cpf;
    private String endereco;

    public Cliente(String nomeCliente, String cpf, String endereco) {
        this.nomeCliente = nomeCliente;
        this.cpf = cpf;
        this.endereco = endereco;
    }

    @Override
    public String toString() {
        return "Cliente: "+ nomeCliente + '\'' +
                ", cpf: " + cpf + '\'' +
                ", endereco='" + endereco + '\'';
    }

}
