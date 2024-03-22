/*
Encapsulamento: É o conceito de esconder os detalhes de implementação de um objeto e apenas 
expor uma interface pública para interagir com ele. Em Dart, você pode controlar o acesso 
aos campos e métodos de uma classe usando modificadores de acesso: public, private, e protected.

* Public: O acesso é permitido a partir de qualquer lugar. Em Dart, se um campo ou método não 
tiver um modificador de acesso, ele é considerado público por padrão.

* Private: O acesso é restrito à própria classe. Para tornar um campo ou método privado em 
Dart, você pode adicionar um _ antes do nome.

* Protected: Não há um modificador de acesso protegido em Dart, mas você pode simular o 
comportamento protegido prefixando um campo ou método com um _ e uma convenção de 
nomenclatura, indicando que é destinado apenas para uso interno ou de subclasses.
*/

class ContaBancaria{
  double _saldo = 0; // _ deixa o atributo privado

  void depositar(double dinheiro){
    this._saldo += dinheiro;
  }

  double get saldo => _saldo; // Metodo publico para acessar o saldo;
}

void main(){
  ContaBancaria conta = new ContaBancaria();
  conta.depositar(300);
  print(conta.saldo); //acessando o metodo public saldo
  print(conta._saldo); //acessando o metodo privado saldo
  //observe que mesmo sendo privado eh possivel acessar o valor diretamente
  //mas ele so eh possivel ser acessado SE criarmos o metodo main() no mesmo arquivo da class
  //em outros arquivos so vai ser possivel acessar o saldo pelo metodo publico

}