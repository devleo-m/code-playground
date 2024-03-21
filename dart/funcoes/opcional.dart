import 'dart:math';
main(){

  var teste1 = numeroAleatorio();
  print("O numero eh: $teste1");

  print(dataPadrao(31));
  print(dataPadrao(8, 7));
  print(dataPadrao(20, 5, 1990));
  print(dataPadrao());

  print(saudacao1("Leonardo", 27));
  print(saudacao2());
}

int numeroAleatorio([int number = 11]){
  return Random().nextInt(number);
}

String dataPadrao([int dia = 1, int mes = 1, int ano = 1970]){
  return dia.toString() + "/" + mes.toString() + "/" + ano.toString();
}

//Observe que o [] dentro do parametro ele deixa a variavel como opcional caso o usuario nao coloque um valor
//vou colocar mais exemplos, um sem o opcional e depois um com

saudacao1(String nome, int idade){
  return "Ola $nome, nem parece que voce tem ${idade + 30}anos";
}

//Agora o mesmo metodo(funcao) com o opcional

saudacao2([String nome = 'Jose', int idade = 30]){
  return "Ola $nome, nem parece que voce tem ${idade + 10}anos";
}