main() {
  //FOR
  print("i:");
  for (int i = 0; i < 10; i++) {
    print("i = $i");
  }

  print("\na: ");
  for (var a = 1; a <= 10; a++) {
    print("a = $a");
  }

  var j = 0;
  print("\n[FORA]j = $j");
  for (var j = 0; j < 20; j++) {
    print("[DENTRO] j = $j");
  }
  print("[FORA]j = $j");

  print("---------------------------");
  
  //double notas = [8.9, 7.4, 5.8, 3.7]; ERROR
  var notas = [8.9, 7.4, 5.8, 3.7];

  //for
  print("for: ");
  for (var i = 0; i < notas.length; i++) {
    print("nota: ${notas[i]}");
  }

  //foreach
  print("for in:");
  for (var key in notas) {
    print("nota: $key");
  }
}
