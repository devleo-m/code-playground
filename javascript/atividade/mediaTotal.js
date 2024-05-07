let notas = [10, 10, 10, 10];
console.log(calculaMedia(notas));

function calculaMedia(notas) {
  var totalNotas = 0;
  for (let i = 0; i < notas.length; i++) {
    totalNotas = totalNotas + notas[i];
  }

  let media = totalNotas / notas.length;
  return media;
}

function conceito(media) {
  switch (media) {
    case media >= 9:
      return "A";
      break;
    case media >= 8:
      return "B";
      break;
    case media >= 7:
      return "C";
      break;
    case media >= 6:
      return "D";
      break;
    case media >= 5:
      return "F";
      break;
    default:
      return "Valor invalido, tente novamente!";
      break;
  }
}

function relatorioNotas(notas) {}
