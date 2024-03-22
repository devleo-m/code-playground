main(){

  print(saudacao1("Leo", 50, true));
  print(saudacao2(idade : 90, casado : false, nome : "Beltrano"));

}

String saudacao1(String nome, int idade, bool casado){
  String casadoAtual = "";
  if(casado){
    casadoAtual = "Casado";
  }else{
    casadoAtual = "Solteiro";
  }
  return "Sr(a). $nome tem $idade anos e atualmente eh $casadoAtual";
}
//Note que isso eh uma funcao normal e temos que declarar elas normalmente pelo parametro por sequencia, correto?
//mas no dart podemos mudar a ordem delas, vou da um exemplo dentro de main

String saudacao2({String nome = '', int idade = 0, bool casado = false}) {
  //Lembrando que se eu quiser nomear em qualquer posicao as variaveis dos parametros eu preciso colocar um {} entre elas
  //Outra obs eh que se eu colocar o {} eu sou obrigado a sempre nomear cada parametro senao o codigo vai da erro.
  String casadoAtual = "";
  if (casado) {
    casadoAtual = "Casado";
  } else {
    casadoAtual = "Solteiro";
  }
  return "Sr(a). $nome tem $idade anos e atualmente é $casadoAtual";
}
