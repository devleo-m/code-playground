/*
* A abstração é um conceito fundamental na programação orientada a objetos.
* Ela se refere à capacidade de representar características essenciais sem incluir os detalhes
* de implementação.
*
* Em Java, a abstração é alcançada através de classes abstratas e interfaces.
*
* Classes Abstratas em Java
* Uma classe abstrata é uma classe que não pode ser instanciada diretamente, mas pode conter
* métodos concretos (com implementação) e métodos abstratos (sem implementação).
* Métodos abstratos são declarados sem corpo e devem ser implementados pelas classes que
* estendem a classe abstrata.
*/
abstract class Form{
    //metodo abstrato (sem implementacao)
    public abstract void desenhar();

    //metodo abstrato (com implementacao)
    public void redimensionar(){
        System.out.println("Redimensionando");
    }
}

/*
* Interfaces em Java
* Uma interface é semelhante a uma classe abstrata, mas todos os métodos em uma interface
* são por padrão abstratos e não possuem implementação. As classes podem implementar várias
* interfaces, fornecendo implementações para os métodos definidos na interface.
*/
//interfaces:
interface Animal{
    void fazerSom(); // Método abstrato (sem implementação)
    void mover(); // Método abstrato (sem implementação)
}

//Utilizando Abstração
//A abstração permite que você defina um conjunto de métodos (abstratos ou não) que representam
// o comportamento esperado de uma classe ou tipo, sem se preocupar com os detalhes específicos
// de implementação. Isso é útil para criar um contrato entre diferentes partes do código,
// facilitando a manutenção e extensão do sistema.

class Circulo extends Form{
    public void desenhar(){
        System.out.println("Desenhar um Circulo");
    }
}

class Cachorro implements Animal{
    public void fazerSom(){
        System.out.println("Latir");
    }
    public void mover(){
        System.out.println("Andar");
    }
}

public class Main {
    public static void main(String[] args){
        Animal meuDog = new Cachorro();
        meuDog.fazerSom();
        meuDog.mover();
        
        
    }
}
