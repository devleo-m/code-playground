main() {
  int n1 = 3;
  double n2 = (-5.67).abs();
  double n3 = double.parse("12.99");
  num n4 = 6; //Num é pai de int e double

  print(n1.abs() + n2 + n3 + n4);

  n4 = 6.7;
  print(n1.abs() + n2 + n3 + n4);

  String s1 = "Bom";
  String s2 = " dia";

  print(s1 + s2);

  //------------------------------------------------//

  List aprovados = ["Pedro", "Amanda", "Beatriz", "Rafael"];
  //print(aprovados is List);
  print(aprovados);
  print(aprovados.elementAt(2));
  print(aprovados[0]);
  print(aprovados.length);

  Map<String, String> telefone = {
    'Carlos': '+55 (48) 98543-9089',
    'José': '+55 (48) 98887-5478',
    'Rafaela': '+55 (48) 98643-5223',
    'Ricardo': '+55 (48) 98853-7432',
    'Dayane': '+55 (48) 98653-85432'
  };

  //print(telefone is Map);
  print(telefone);
  print(telefone['Dayane']);
  print(telefone.length);
  print(telefone.keys);
  print(telefone.values);
  print(telefone.entries);

  Set times = {'Flamengo', 'Vasco', 'São Paulo'};
  //print(times is Set);
  times.add('Vasco');
  times.add('Flamengo');
  times.add('Palmeiras');
  times.add('Palmeiras');
  times.add('Palmeiras');
  print(times.length);
  print(times.contains('Vasco'));
  print(times.first);
  print(times.last);
  print(times);
}
