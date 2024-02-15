//Em Java, o bloco try-catch é usado para lidar com exceções (erros) que podem ocorrer
// durante a execução do código. Ele permite que você escreva código que pode gerar
// exceções e fornece um mecanismo para lidar com essas exceções de maneira controlada.
public class tryCatch {
    //Aqui está a estrutura básica do bloco try-catch:
    public static void main(String[] args) {
        int num1 = 40;
        int num2 = 0;
        int total = num1/num2;
        try {
            System.out.println(total);
        }catch (ArithmeticException e){
            System.out.println("Error: "+ e);
        }finally {
            System.out.println("Executando Finally!!!");
        }
    }

    //try: Este é o bloco onde você coloca o código que pode gerar uma exceção.
    //
    //catch: Dentro deste bloco, você especifica o tipo de exceção que você está tentando
    // capturar. Se uma exceção do tipo especificado ocorrer no bloco try, o bloco catch
    // correspondente será executado.
    //
    //finally: Este bloco é opcional e é usado para colocar código que será executado
    // independentemente de uma exceção ser lançada ou não. Pode ser útil para ações que
    // precisam ser realizadas, como fechar recursos, independentemente do que aconteça
    // no bloco try.

}
