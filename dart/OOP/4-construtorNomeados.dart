class Pessoa{
  String? nome;
  int? idade;
  String? cpf;
  double? altura;

//forma mais moderna de se criar um construtor
  Pessoa(this.nome, this.idade, this.cpf, this.altura);

  Pessoa.idade(this.idade);

  // Pessoa(String nome, int idade, String cpf, double altura){
  //   this.nome = nome;
  //   this.idade = idade;
  //   this.cpf = cpf;
  //   this.altura = altura;
  // }
}

class Pessoa2 extends Pessoa {
  Pessoa2({String nome = "Desconhecido", int idade = 1, String cpf = "000.000.000-00", double altura = 0.00}) : super(nome, idade, cpf, altura);
  Pessoa2.idade({int idade = 99});
}

main(){
  //normal
  Pessoa2 p1 = new Pessoa2(idade  : 27, altura: 1.70, cpf: "111.222.333-99", nome: "Leo");
  print(p1.nome);
  print(p1.cpf);

  //opcional
  Pessoa2 p2 = new Pessoa2();
  print(p2.nome);
  print(p2.cpf);

  //nomeado
  Pessoa2 p3 = new Pessoa2.idade();
  print(p3.idade);
}