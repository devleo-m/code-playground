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
            int x = 10, y = 20, j = 30, z = 40;
            Console.WriteLine((x + y + j + z)/4); //media

            /*
            As regras gerais para nomear variáveis ​​são:

            Os nomes podem conter letras, dígitos e o caractere sublinhado (_)
            Os nomes devem começar com uma letra ou sublinhado
            Os nomes devem começar com letra minúscula e não podem conter espaços em branco
            Os nomes diferenciam maiúsculas de minúsculas ("myVar" e "myvar" são variáveis ​​diferentes)
            Palavras reservadas (como palavras-chave C#, como int ou double) não podem ser usados ​​como nomes
             */

        }
    }
}

