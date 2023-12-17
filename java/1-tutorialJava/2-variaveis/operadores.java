public class operadores {
    public static void main(String[] args) {
    //Aritméticos:
    //+, -, *, /, %: São operadores para adição, subtração, multiplicação, divisão 
    //e módulo (resto da divisão), respectivamente.
        int soma = 5+5;
        int subtracao = 5-5;
        int multiplicacao = 5*5;
        int divisao = 5/5;
        int resto = 5%5;

        System.out.println(soma);
        System.out.println(subtracao);
        System.out.println(multiplicacao);
        System.out.println(divisao);
        System.out.println(resto);

        //Incremento e Decremento:
        //++, --: São operadores para incrementar ou decrementar valores.

        int contador = 0;
        contador++; // Incrementa o valor de 'contador' por 1
        contador--; // Decrementa o valor de 'contador' por 1
        System.out.println(contador);

        //Concatenação:
        //O operador + também é usado para concatenar strings em Java.

        String nome = "João";
        String sobrenome = "Silva";
        System.out.println(nome +" "+sobrenome);

        //Java, a partir da versão 5, você pode utilizar a formatação de strings com
        // o operador % ou usar o método String.format() 
        System.out.println(String.format("Ola, %s %s", nome, sobrenome));

        //A partir do Java 15, foi introduzido um recurso experimental chamado "text blocks"

        String mensagem = """
                ola,
                seu nome eh %s
                e seu nomenome eh %s
                """;
        String mensagemFinal = String.format(mensagem, nome, sobrenome);
        System.out.println(mensagemFinal);


    }
}
