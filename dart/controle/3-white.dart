import "dart:io";

main(){
  String sair = 'sair';
  String digitado = ''; 

  while(digitado != sair){
    print('Digite algo ou sair')
    digitado = stdin.readLineSync()!;
    print(digitado);
  }
  print("FIM!");
}