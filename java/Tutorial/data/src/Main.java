import java.time.LocalDate; // import the LocalDate class
import java.time.LocalTime; // import the LocalTime class
import java.time.LocalDateTime; // import the LocalDateTime class
import java.time.format.DateTimeFormatter; // Import the DateTimeFormatter class

public class Main {
    public static void main(String[] args) {
        //Datas Java
        //Java não possui uma classe Date integrada, mas podemos importar o java.time
        // pacote para trabalhar com a API de data e hora. O pacote inclui muitas classes
        // de data e hora. Por exemplo:

        //Class	               Description
        //LocalDate	           Represents a date (year, month, day (yyyy-MM-dd))
        //LocalTime	           Represents a time (hour, minute, second and nanoseconds (HH-mm-ss-ns))
        //LocalDateTime	       Represents both a date and a time (yyyy-MM-dd-HH-mm-ss-ns)
        //DateTimeFormatter	   Formatter for displaying and parsing date-time objects

        //Exibir data atual
        //Para exibir a data atual, importe a java.time.LocalDateclasse e use seu now()método:

        //import java.time.LocalDate; // import the LocalDate class
        LocalDate diaHoje = LocalDate.now(); // Create a date object
        System.out.println(diaHoje); // Display the current date
        //2024-02-20 // Exemplo da data exibida nesse dia que printei
        //OBS; EU BATI O CARRO NESSE DIA KKK

        //Exibir hora atual
        //Para exibir a hora atual (hora, minuto, segundo e nanossegundos), importe a
        // java.time.LocalTimeclasse e use seu now()método:
        //import java.time.LocalTime; // import the LocalTime class
        LocalTime horarioLocal = LocalTime.now();
        System.out.println(horarioLocal);
        //23:02:39.635744 //Horas que foi printada no momento

        //Exibir data e hora atuais
        //Para exibir a data e hora atuais, importe a java.time.LocalDateTimeclasse e
        // use seu now()método:
        LocalDateTime diaHorario = LocalDateTime.now();
        System.out.println(diaHorario);
        //2024-02-20T23:02:39.636025
        //Dia e hora printada no momento do metodo chamado

        LocalDateTime myDateObj = LocalDateTime.now();
        System.out.println("Before formatting: " + myDateObj);
        DateTimeFormatter myFormatObj = DateTimeFormatter.ofPattern("dd-MM-yyyy HH:mm:ss");
        DateTimeFormatter myFormatObj2 = DateTimeFormatter.ofPattern("dd-MM-yyyy");

        String formattedDate = myDateObj.format(myFormatObj);
        String formattedDate2 = myDateObj.format(myFormatObj2);

        System.out.println("After formatting: " + formattedDate);
        System.out.println("After formatting: " + formattedDate2);

        //OBS:
        //O ofPattern() método aceita todos os tipos de valores, se você quiser exibir a
        // data e a hora em um formato diferente. Por exemplo:
        //yyyy-MM-dd	   "1988-09-29"
        //dd/MM/yyyy	   "29/09/1988"
        //dd-MMM-yyyy	   "29-Sep-1988"
        //MMM dd yyyy      "Thu, Sep 29 1988"
    }
}