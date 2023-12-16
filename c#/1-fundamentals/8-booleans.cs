//Muitas vezes, na programação, você precisará de um tipo de dados que só possa ter um 
//de dois valores, como:
// SIM NÃO
// LIGADO DESLIGADO
// VERDADEIRO FALSO
// Para isso, C# possui um tipo de dados bool, que pode assumir os valores true ou false.

//Um tipo booleano é declarado com a palavra-chave bool e só pode receber os valores true ou false:
bool verdadeiro = true;
bool falso = false;
Console.WriteLine(verdadeiro);
Console.WriteLine(falso);

//Expressão Booleana
//Uma expressão booleana retorna um valor booleano: True ou False, comparando valores/variáveis.
//Isso é útil para construir lógica e encontrar respostas.
//Por exemplo, você pode usar um operador de comparação, como maior que>) para descobrir se uma expressão (ou variável) é verdadeira: (

int x = 10;
int y = 9;
Console.WriteLine(x > y); // Retorna verdadeiro 10 eh maior do que 9

//Exemplo da vida real
//Vamos pensar em um "real exemplo de vida" onde precisamos descobrir se uma pessoa tem idade suficiente para votar.
//No exemplo abaixo, usamos o operador de comparação >= para descobrir se a idade (25) é maior que OU igual a o limite
//de idade para votar, que é definido para 18:

Console.WriteLine("Digite a sua idade: ");
int idade = Convert.ToInt32(Console.ReadLine()); //Ja estudamos sobre isso na aula 5 - entrada do usuario
Console.WriteLine(18 > idade);
//esse exemplo ficaria melhor com if/else, mas vou falar sobre isso

if (idade >= 18) //Se tiver maior a idade maior que 18, sera true, entao o codigo dentro do if sera executado
{
    Console.WriteLine("Velho suficiente para votar!");
}
else // se a idade for menor que 18, entao sera executado esse codigo else
{
    Console.WriteLine("Não tem idade suficiente para votar.");
}







