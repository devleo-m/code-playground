//Agora, uma classe pode decidir aderir a essa interface usando a palavra-chave implements e,
// em seguida, fornecer uma implementação para os métodos da interface.
public class MinhaClass implements MinhaInterface{
    public void metodo1(){
        System.out.println("Printando o metodo1");
    }

    public Byte metodo2() {
        System.out.println("Printando o metodo2");
        return 10;
    }
    //Aqui, MinhaClasse está aderindo à interface MinhaInterface e fornecendo implementações para os métodos metodo1 e metodo2.
}
