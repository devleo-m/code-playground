package Metodos;

public class Metodos {
    public int soma(int x, int y){
        return x+y;
    }
    public static void main(String[] args) {
        Metodos calc = new Metodos();
        int resultado = calc.soma(5,5);
        System.out.println("O resultado da soma eh: "+ resultado);
    }
}
