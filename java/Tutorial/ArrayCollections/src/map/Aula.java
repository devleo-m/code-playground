package map;

import java.util.HashMap;
import java.util.Map;

public class Aula {
    public static void main(String[] args) {
        Map<Integer, String> usuarios = new HashMap<>();

        usuarios.put(1,"Rafaela");
        usuarios.put(2,"Rebeca");
        usuarios.put(3,"Gustavo");
        usuarios.put(4,"Luciano");
        usuarios.put(5,"Amanda");

        System.out.println(usuarios.size());   //tamanho do map, no caso sao 5 criados
        System.out.println(usuarios.isEmpty()); //Falso, pois ele tem valores, se estivesse vazio seria true

        System.out.println(usuarios.keySet()); //Tamanho do map, organizado sao 5
        System.out.println(usuarios.values()); //Mostra apenas o nome de todos da lista
        System.out.println(usuarios.entrySet()); //mostra todos os nomes e indice organizado

        System.out.println(usuarios.containsKey(3)); //Verifica se existe o indice 3 e retornar boolean
        System.out.println(usuarios.containsValue("Luciano")); //Faz o mesmo aqui, mas verifica se a string existe e retornar um boolean

        for (int keyIndice : usuarios.keySet()){
            System.out.println(keyIndice);
        }
        for (String keyNome : usuarios.values()){
            System.out.println(keyNome);
        }

        for (Map.Entry<Integer, String> registro : usuarios.entrySet()){
            System.out.println(registro.getKey());
            System.out.println(registro.getValue());
        }
    }


}
