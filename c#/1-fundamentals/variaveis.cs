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

            const double test = 3.4567; //desta forma se cria uma constante

            //-----------------------------------------------------------------------------------

            //entrada de usuario;
            Console.WriteLine("digite seu nome:");
            string nomeUser = Console.ReadLine();

            Console.WriteLine($"Ola, {nomeUser}");


        }
    }
}

