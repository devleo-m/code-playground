/* Encapsulamento
*
* O encapsulamento é um princípio fundamental da programação orientada a objetos que se
* concentra em ocultar os detalhes internos de uma classe e permitir o acesso
* controlado aos seus componentes (métodos e atributos). Em Java, o encapsulamento é alcançado
* através do uso de modificadores de acesso e métodos getters e setters.
*
* # Modificadores de Acesso em Java
* Java oferece quatro tipos de modificadores de acesso:
* 1 - public: Acesso irrestrito. As classes, métodos ou atributos marcados como públicos podem
* ser acessados de qualquer lugar.
* 2 - private: Acesso restrito à própria classe. Métodos e atributos marcados como privados
* só podem ser acessados dentro da classe onde foram definidos.
* 3 - protected: Acesso restrito à classe e suas subclasses. Métodos e atributos marcados como
* protegidos podem ser acessados dentro da classe e por suas subclasses.
* 4 - default (ou package-private): Acesso restrito ao pacote. Se nenhuma modificação de acesso
* é especificada, o acesso é limitado ao próprio pacote.
*
*/
class Pessoa{
    private String name; //atributo privado

    //metodo para acessar o nome(getter)
    public String getName(){
        return name;
    }

    //Metodo para modificar o nome(setter)
    public void setName(String newName){
        this.name = newName;
    }
}

//obs:
/*
* Neste exemplo, nome é um atributo privado da classe Pessoa. Isso significa que não pode ser
* acessado diretamente de fora da classe. No entanto, a classe fornece métodos
* getNome() e setNome()
* para acessar e modificar o atributo nome de forma controlada.
*/

public class Main {
    public static void main(String[] args) {
        Pessoa alguem = new Pessoa();
        alguem.setName("Fulano de tal");
        System.out.println("O nome da pessoa eh: "+alguem.getName());
    }
}
