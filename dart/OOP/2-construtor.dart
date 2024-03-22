class Pessoa{
  String? nome;
  int? idade;
  String? cpf;
  double? altura;

  Pessoa(String nome, int idade, String cpf, double altura){
    this.nome = nome;
    this.idade = idade;
    this.cpf = cpf;
    this.altura = altura;
  }

}

main(){
  //Com o construtor criado nao eh mais possivel instanciar uma classe desta forma
  // Pessoa p1 = new Pessoa();
  // p1.nome = "Leo";
  // p1.idade = 27;
  // p1.cpf = "111.222.333-99";
  // p1.altura = 1.70;

  // print(p1.nome);
  // print(p1);
  // print(p1.cpf);

  Pessoa p2 = new Pessoa("Leo", 27, "111.222.333-99", 1.70);

  print(p2.nome);
  print(p2);
  print(p2.cpf);

  //Criando um construtor fica mais facil de declarar os valores de uma class

}