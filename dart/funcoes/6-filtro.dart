main(){
  var lista = [3,64,7,12,9,3,54,74,53,2,76];
  var listaBoa = [];

  for(var key in lista){
    if(key >= 50){
      listaBoa.add(key);
    }
  }
  print(lista);
  print(listaBoa);

  //---------------------------------
  //WHERE

  var listaOtima = (int lista) => lista >= 60;
  var listaPessima = (int lista) => lista <= 10;

  print(lista.where(listaOtima));
  print(lista.where(listaPessima));
}