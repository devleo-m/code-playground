public class MinhaClasse implements MinhaInterface{
    @Override
    public void metodo1() {

    }
    @Override
    public int metodo2(String parametro) {
        return parametro.length();
    }
}
//Para implementar uma interface em uma classe, você usa a palavra-chave implements.
// A classe então deve fornecer implementações para todos os métodos declarados na interface.
