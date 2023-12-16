using System;

namespace variaveis
{
    class Program
    {
        static void Main(string[] args)
        {
            int idade = 40;
            double dinheiro = 150.50;
            bool verdadeiro = true; //or false
            char sala = 'b';
            string nome = "Fulano de tal da Silva";

            Console.WriteLine($"nome: {nome} idade: {idade} anos."); //Melhor forma de concatenar

            //constantes
            const double pi = 3.14159; //desta forma se cria uma constante
            //Se não quiser que outras pessoas (ou você mesmo) substituam os valores existentes, adicione a palavra-chave const na frente do tipo de variável.
            //Isso declarará a variável como "constante", o que significa imutável e somente leitura:
            //pi = 20; // error
            //-----------------------------------------------------------------------------------

            //entrada de usuario;
            Console.WriteLine("digite seu nome:");
            string nomeUser = Console.ReadLine();

            Console.WriteLine($"Ola, {nomeUser}");

            //Multiplas variveis
            int x, y, j, z = 10, 20, 30, 40;
            Console.WriteLine(y);

        }
    }
}

