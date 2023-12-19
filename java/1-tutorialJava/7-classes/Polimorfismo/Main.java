//o polimorfismo é um conceito fundamental na programação orientada a objetos e é
// amplamente utilizado em Java.

//O que é Polimorfismo?
//O polimorfismo é a capacidade de um objeto ser referenciado de várias formas.
// Em Java, isso se refere à habilidade de objetos de classes diferentes serem tratados
// através de uma mesma interface.
// Em outras palavras, diferentes classes podem ser tratadas de maneira uniforme, permitindo
// flexibilidade e reutilização de código.

//OBS:
//Existem dois tipos de polimorfismo em Java: polimorfismo de tempo de compilação (ou estático)
//e polimorfismo de tempo de execução (ou dinâmico).

//Polimorfismo de Tempo de Compilação
//O polimorfismo de tempo de compilação ocorre quando o tipo de referência de um objeto é
// decidido em tempo de compilação. Isso significa que, mesmo que o objeto seja de uma subclasse,
// se ele for referenciado pela classe pai, somente os métodos da classe pai serão acessíveis.
class Animal { //Criando a classe animal
    public void fazerSom() { //metodo fazer barulho
        System.out.println("Som do animal");
    }
}
class Cachorro extends Animal { //criando o objeto cachorro e herdando a class animal
    public void fazerSom() { //executando o metodo da classe herdada de animal
        System.out.println("Latido do cachorro");
    }
}

class Gato extends Animal{
    public void fazerSom(){
        System.out.println("Miado do gato");
    }
}

public class Main {
    public static void main(String[] args){
        
        Animal simba = new Cachorro(); //simba eh nome do meu dog
        simba.fazerSom(); //executando o metodo latir

        Animal aurora = new Gato(); //aurora nome da minha gata
        aurora.fazerSom();

        Animal cachorro = new Cachorro();
        cachorro.fazerSom(); //Resultado: "Latido do cachorro"
        cachorro = new Animal();
        cachorro.fazerSom(); //Resultado: "Som do animal"
    }
}
