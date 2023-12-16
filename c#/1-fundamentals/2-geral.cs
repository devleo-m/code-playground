//variaveis

int numeroInteiro = 30;
double numeroFlutuante = 50.20;
string caractere = "arolnode";
bool boolean = true; //or false

//imprimir / console / print / sout

Console.WriteLine(numeroFlutuante);
Console.WriteLine("teste " + caractere);
Console.WriteLine($"teste {caractere} eh {boolean}");

//Entrada/ escreva / scanner /

Console.WriteLine("Escreva seu nome: ");
string nomeUser = Console.ReadLine();
Console.WriteLine($"Ola, {nomeUser}");

//Estrutura de controle / If Else

if (numeroInteiro > 20)
{
    Console.WriteLine($"{numeroInteiro} eh maior que 20");
}
else if (numeroInteiro == 20)
{
    Console.WriteLine($"{numeroInteiro} eh igual a 20");
}
else
{
    Console.WriteLine($"{numeroInteiro} eh menor que 20");
}

//Loops / while / do-while / for
    
int loop = 0; // Variavel padrao de todos os loops

//while
while (loop < 10)
{
    Console.WriteLine($"While {loop}");
    loop++;
}

//do while
do
{
    Console.WriteLine($"Do-While {loop}");
    loop++;
} while (loop < 20);

//for
for (int i = loop; i < 30; i++)
{
    Console.WriteLine($"for {i}");
}

//Arrays
int[] numeros = new int[5]; //cria um array de inteiros com 5 elementos
numeros[0] = 1;
