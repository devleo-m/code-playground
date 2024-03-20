package EqualsHashcode;

import java.util.HashSet;

//hash set
public class aula2 {
    public static void main(String[] args) {
        HashSet conjunto = new HashSet();

        conjunto.add(1.3);
        conjunto.add(true);
        conjunto.add("Teste");
        conjunto.add(1);
        conjunto.add('X');
        System.out.println("Tamanho do conjunto: "+conjunto.size()); //resultado 5

        conjunto.add(1);
        System.out.println("Tamanho do conjunto: "+conjunto.size()); //resultado 5

        //Observe que mesmo eu colocando um novo conjunto nao mostrou que foi adicionado mais um number
        //isso ocorre pq o hashset nao aceita valores duplicados[repetidos], ja existe um number 1

        conjunto.add("teste");
        System.out.println("Tamanho do conjunto: "+conjunto.size()); //resultado 6
        //Observe que aqui ele foi adicionado mesmo ja tendo um string teste, a diferenca esta na
        //primeira letra maiuscula e nessa segunda esta minuscula, Ou seja sao valores diferentes

        conjunto.add("teste");
        System.out.println("Tamanho do conjunto: "+conjunto.size()); //resultado 6
        //observe que aqui ele nao aceitou e nao adicionou pelo fato da string ser repetida

        System.out.println(conjunto.remove(1.3));
        //se ele remover algum valor que existe na hashset, ele retorna um true, senao um false
        System.out.println(conjunto.remove(5));

        //printar tudo que tem em conjunto
        System.out.println(conjunto);

    }
}
