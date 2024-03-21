//import "dart:io";
import 'dart:math';

main(){
  int numero1 = 2;
  int numero2 = 5;
  soma(numero1, numero2);

  //print("Qual eh seu nome?");
  //String nome = stdin.readLineSync()!;
  String nome1 = "Leonardo";
  sobrenomeMadeira(nome1);

  numeroAleatorio();
}

soma(int a, int b){
  print(a + b);
}

void sobrenomeMadeira(String nome){
  print("$nome Madeira");
}

void numeroAleatorio(){
  var x = Random().nextInt(11);
  var y = Random().nextInt(11);
  
  print("Valor de x: $x e y: $y");
  soma(x, y);
}