main(){
  // tipo nome do valor
  int Function(int, int) soma1 = somaFn;
  print(soma1(2,3));

  int Function(int, int) soma2 = (int x, int y){
    return x + y;
  };
    print(soma2(20,313));

  int soma3(int j, int k){
     return somaFn(j, k);
  }

  print(soma3(1, 1));

  int soma4(int q, int w){
    return q + w;
  }
  print(soma4(10,10));
}

  int somaFn(int a, int b){
    return a + b;
  }

