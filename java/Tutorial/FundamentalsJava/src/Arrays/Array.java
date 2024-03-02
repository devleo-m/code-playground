package Arrays;

public class Array {
    public static void main(String[] args) {
        // Declaração de um array de inteiros com tamanho 5
        int[] numero = new int[20];

        // Inicialização de um array de strings com valores
        String[] nomes = {"arolnode", "gin", "hin", "yen"};

        System.out.println(numero[3]);
        System.out.println(nomes[0]);

        //Percorrendo um Array:
        //Existem várias maneiras de percorrer um array. Um dos métodos comuns é
        //usar um loop for:

        int[] numeros = {1, 2, 3, 4, 5};
        for (int i : numeros) {
            System.out.println(i);
        }

        //ou
        System.out.println(" "); //apenas um quebra linha

        for (int i = 0; i < numeros.length; i++) {
            System.out.println(i);
        }


    }
}
