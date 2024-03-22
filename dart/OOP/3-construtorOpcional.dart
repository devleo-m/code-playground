class Pessoa{
  String? nome;
  int? idade;
  String? cpf;
  double? altura;

  Pessoa([String nome = "Unknown", int idade = 1, String cpf = "000.000.000-00", double altura = 0.00]){
    this.nome = nome;
    this.idade = idade;
    this.cpf = cpf;
    this.altura = altura;
  }

}

main(){
  //normal
  Pessoa p1 = new Pessoa("Leo", 27, "111.222.333-99", 1.70);
  print(p1.nome);
  print(p1.cpf);

  //opcional
  Pessoa p2 = new Pessoa();
  print(p2.nome);
  print(p2.cpf);
}