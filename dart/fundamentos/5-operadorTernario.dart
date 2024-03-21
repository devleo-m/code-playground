import "dart:io";
main(){
  print("Esta chovendo? ");
  bool estaChovendo = stdin.readLineSync() == "s";

  print("Esta frio? ")
  bool estaFrio = stdin.readLineSync() == "s";

  String resultado = estaChovendo || estaFrio ? "Ficar em casa" : "Sair!!!";
  print(resultado);
}