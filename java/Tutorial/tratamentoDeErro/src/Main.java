
//Em Java, o bloco try-catch é usado para lidar com exceções (erros) que podem ocorrer
// durante a execução do código. Ele permite que você escreva código que pode gerar
// exceções e fornece um mecanismo para lidar com essas exceções de maneira controlada.
public class Main {
    public static void main(String[] args) {
        int num1 = 2;
        int num2 = 0;

        //Aqui está a estrutura básica do bloco try-catch:

        try{
            int resultado = num1 / num2;
            System.out.println(resultado);
        } catch (ArithmeticException e){
            System.err.println("Erro: Divisão por zero não é permitida.");
        } finally {
            System.out.println("Executando o finally");
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

    //OBS:Neste exemplo, a função divide tenta dividir dois números. Se o denominador for zero,
    // uma exceção ArithmeticException será lançada. O bloco catch correspondente trata essa
    // exceção, e o bloco finally é executado independentemente do que aconteça.
}
