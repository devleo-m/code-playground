package Arrays;

public class Manipulacao {
    public static void main(String[] args) {
        //O método length() retorna o número de caracteres em uma string.
        String mensagem = "Ola, mundo!";
        int tamanho = mensagem.length(); //retorna 12
        System.out.println(tamanho);

        //Para comparar o conteúdo de duas strings, use equals()
        String msg1 = "oi";
        String msg2 = "hi";

        boolean saoIguais = msg1.equals(msg2); //retorna falso pois nao sao iguais
        System.out.println(saoIguais);

        //O método indexOf() retorna a posição da primeira ocorrência de um caractere ou substring.
        String frase = "O rato roeu a roupa.";
        int posicao = frase.indexOf("roeu"); // Retorna a posição onde "roeu" começa (4)
        System.out.println(posicao);

        //O método split() divide uma string em substrings com base em um delimitador.
        String dados = "nome:sobrenome:idade";
        String[] partes = dados.split(":"); // Divide a string em partes usando ":"
        System.out.println(partes);
        // partes[0] = "nome", partes[1] = "sobrenome", partes[2] = "idade"

        //O método replace() substitui todas as ocorrências de uma substring por outra.
        String frase2 = "Gatos são ótimos!";
        String novaFrase = frase2.replace("Gatos", "Cachorros"); // Substitui "Gatos" por "Cachorros"
        System.out.println(novaFrase);
        // novaFrase será "Cachorros são ótimos!"
    }
}
