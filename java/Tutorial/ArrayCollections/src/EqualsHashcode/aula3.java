package EqualsHashcode;

import java.util.HashSet;
import java.util.Set;
//SET
public class aula3 {
    public static void main(String[] args) {
        Set<String> listasDeCandidatos = new HashSet();
        listasDeCandidatos.add("Jose");
        listasDeCandidatos.add("Ana");
        listasDeCandidatos.add("Bella");
        listasDeCandidatos.add("Lucas");
        listasDeCandidatos.add("Davi");

        for (String candidatos : listasDeCandidatos){
            System.out.println(candidatos);
        }
    }
}
