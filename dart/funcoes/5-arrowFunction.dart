main(){
  int somar1(int x, int y){
    return x + y;
  }
  //forma padrao de criar uma funcao
  print(somar1(4, 6));

  //agora vou criar a mesma funcao so que da forma como flecha(arrow)
  var somar2 = (int x, int y) => x + y;
  print(somar2(2,8));

  var subtrair = (double x, double y) => x - y;
  var multiplicar = (int x, int y) => x * y;
  var dividir = (int x, int y) => x / y;

  print(subtrair(10,2) + multiplicar(5,5) + dividir(50,2) / 2);
}