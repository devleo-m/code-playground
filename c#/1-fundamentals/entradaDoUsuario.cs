//Entrada do usuário

//Você já aprendeu que Console.WriteLine() é usado para gerar (imprimir) valores.
//Agora usaremos Console.ReadLine() para obter a entrada do usuário.

Console.WriteLine("Escreva seu nome: ");

string nome = Console.ReadLine();

Console.WriteLine(nome);

//OBS:::: Entrada e números do usuário
//O método Console.ReadLine() retorna um string. Portanto, você não pode obter informações de outro tipo de dados,como int.
//O programa a seguir causará um erro:

/*
Console.WriteLine("Digite sua idade:");
int age = Console.ReadLine();
Console.WriteLine("voce tem "+age+" anos!");
*/

//para corrigir esse problema, eh so seguir com a mesma estrategia na aula passada. aula 4
Console.WriteLine("Digite a sua idade: ");
int age = Convert.ToInt32(Console.ReadLine());
Console.WriteLine(age);




