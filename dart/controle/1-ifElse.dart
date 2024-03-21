import 'dart:math';

main(){
  String aluno = "Fulano";
  int nota = Random().nextInt(11);
  print("Aluno: $aluno, nota: $nota");

  if(nota >= 9){
    print("$aluno esta no quadro de honra - APROVADO");
  } else if(nota >= 7){
    print("APROVADO!");
  } else if(nota >= 5){
    print("RECUPERACAO!");
  } else if(nota >= 4){
    print("RECUPERACAO + TRABALHO");
  } else{
    print("REPROVADO!");
  }
}