public class Main {
    public static void main(String[] args){
        Carro meuCarro = new Carro();
        meuCarro.Modelo = "Celta";
        meuCarro.Cor = "Branco";
        meuCarro.ano = "2013";

        meuCarro.ligarCarro();
        meuCarro.acelerar(100);
    }
}



