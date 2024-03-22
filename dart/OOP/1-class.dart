class Data{
  int? dia;
  int? mes;
  int? ano;
}

main(){
  var dataAniversario = new Data();
  dataAniversario.dia = 8;
  dataAniversario.mes = 2;
  dataAniversario.ano = 1992;

  print("${dataAniversario.dia}/${dataAniversario.mes}/${dataAniversario.ano}");
}

//OBS:
//Com a nova atualizacao do dart nao precisava colocar ? no objeto sem valor como esta 

// class Data{
//   int dia;
//   int mes;
//   int ano;
// }