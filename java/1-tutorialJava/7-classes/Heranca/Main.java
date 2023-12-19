/*

a herança é um conceito fundamental na programação orientada a objetos que permite que uma
classe herde características (atributos e métodos) de outra classe. Em Java, você pode usar
a palavra-chave extends para implementar herança.

Vamos criar um exemplo simples para ilustrar isso. Suponha que temos a classe Veiculo e
queremos criar a classe Carro que herda de Veiculo:

 */
public class Main {
    public static void main(String[] args){
        Carro meuCarro = new Carro();
        meuCarro.Marca = "Fusca"; //Observe que a marca esta declarada em veiculo
        meuCarro.numPortas = 4; //numportas esta declarada em Carro e nao em veiculo

        meuCarro.ligar();
        meuCarro.abrirPortas();
    }
}
//Neste exemplo, a classe Veiculo contém os atributos marca e ano, além dos métodos ligar() e
//desligar(). A classe Carro herda esses atributos e métodos da classe Veiculo usando extends
// e adiciona um novo atributo numPortas e um método abrirPortas().

//OBS que ele adicionou tudo o que tem "dentro" de Veiculo e foi adicionado em carro
//e agora alem de tudo o que o veiculo tem, na classe carro tem mais novas coisas
//que no caso eh o numportas e o metodo abrirportas


