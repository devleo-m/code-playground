main(){
  String nome = "Leonardo";
  String status = "Aprovado";
  int idade = 23;

  String frase1 = nome + " tem "+ idade + " anos e está: "+ status;
  String frase2 = "$nome tem $idade anos e está $status";

  print(frase1); 
  print(frase2);

  print("1 + 1 = ${1 + 1}");
  print("30 * 30 = ${30 * 30}");
}