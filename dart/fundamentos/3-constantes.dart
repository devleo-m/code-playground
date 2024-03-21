import 'dart:io';

main() {
  const PI = 3.1415;

  //area da cicurferencia = PI* raio * raio
  stdout.write('Informe o valor do raio: '); //stdout.write é igual o print, a diferenca é que ele nao pula uma linha.
  var scanner = stdin.readLineSync()!;
  final raio = double.parse(scanner);

  print(raio.runtimeType);
  print("O valor do raio é: $raio");

  final area = PI * (raio * raio);
  print('O valor da area é: $area');
}
