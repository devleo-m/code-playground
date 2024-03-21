import 'dart:math';

main(){
  var nota = Random().nextInt(11);
  switch(nota){
    case 10:
     print("Quadro de honra, aluno nota 10!");
     break;
    case 9:
      print("Aluno nota 9, parabens!");
      break;
    case 8:
      print("Aluno nota 8, parabens!");
      break;
    case 7:
      print("Aluno nota 7, precisa melhorar mais!");
      break;
    case 6:
      print("Aluno nota 6, voce precisa se esforcar mais!");
      break;
    case 5:
      print("Aluno nota 5, recuperacao, estude mais!");
      break;
    case 4:
      print("Aluno nota 4, reprovado...!");
      break;
    case 3:
      print("Aluno nota 3, que coisa feia!");
      break;
    case 2:
      print("Aluno nota 2, como voce conseguiu essa nota?!");
      break;
    case 1:
      print("Aluno nota 1, misericordia!");
      break;
    case 0:
      print("Aluno nota 0, sem comentarios!");
      break;
    default:
    print("Numero invalido");
    }
}