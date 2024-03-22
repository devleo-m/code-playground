import "dart:math";
main(){
  var impar = () => print("Numero impar!");
  var par = () => print("Numero par!");

  parOuImpar(par(), impar());
}

  void parOuImpar(Function fnPar, Function fnImpar){
    var numeroSorteado = Random().nextInt(11);
    print("Valor do numero sorteado: $numeroSorteado");
    numeroSorteado % 2 == 0 ? fnPar() : fnImpar();
  }
