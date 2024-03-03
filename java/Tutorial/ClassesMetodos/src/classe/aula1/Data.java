package classe.aula1;

public class Data {
    int dia;
    int mes;
    int ano;

    public String diaCompleto(){
        String s = (dia + "/" + mes + "/" + ano);
        return s;
    }
}
