import java.time.LocalDate; // import the LocalDate class
import java.time.LocalTime; // import the LocalTime class
import java.time.LocalDateTime; // import the LocalDateTime class

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
    }
}