public class teste {
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
}
