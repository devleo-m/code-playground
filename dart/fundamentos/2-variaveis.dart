main() {
  String nome = "Fulano de tal";
  int idade = 19;
  double dinheiro = 1.99;
  bool casado = true;

  if (idade >= 18) {
    print("$nome tem $idade anos. Maior de idade");
  } else {
    print("$nome tem $idade anos. Menor de idade");
  }

  var b = 2;
  var c = 1.99;
  print(b + c);

  print(b.runtimeType); //Ele print int/ informando que o b eh do tipo int
  print(c.runtimeType); //double, mesma coisa do de cima, so muda o tipo da variavel
}
