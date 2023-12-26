#include <stdio.h>
#include <stdlib.h>
#include <locale.h>
int main()
{
    setlocale(LC_ALL, ""); // Acentuação PT-BR

    char nome[1][34] = {"Leonardo Madeira"}; // Vou ser sincero, não estou conseguindo declarar o nome digitando ele com printf e scanf.
    int ru = 3235609;
    float sal , IMP, total;
    
    printf("|--------------------------------------------|\n");
    printf("| Base de cálculo mensal em R$  | Aliquota % | |\n");
    printf("|--------------------------------------------|\n");
    printf("| Até 1.637,11                  |  -         |\n");
    printf("| De 1.637,12 até 2.453,50      | 7,5        |\n");
    printf("| De 2.453,51 até 3.271,38      | 15,0       |\n");
    printf("| De 3.271,39 até 4.087,65      | 22,5       |\n");
    printf("| Acima de 4.087,65             | 27,5       |\n");
    printf("|--------------------------------------------|\n");

    printf("\n\n");
    printf("Salário: ");
    scanf_s("%f", &sal); // Registrar o salário do usuario.

    IMP = (sal); // Jogando o salario na variavel IMP
    if (IMP <= 1637.11) // Caso o salario do usuario for menor que 1637, nao paga nada de imposto.
    {printf("0 de imposto");}
    else if (IMP >= 1637.12 && IMP <= 2453.50)// Caso o salario do usuario for maior que 1637 e menor que 2453.50, vai pagar 7,5 de Aliq imposto.
    {total = IMP * 0.075;} // Para calcular a porcentagem, eu prefiro usar a multiplicação, por ser mais pratico.
    else if (IMP >= 2453.51 && IMP <= 3271.38)
    {total = IMP * 0.15;}
    else if (IMP >= 3271.39 && IMP <= 4087.65)
    {total = IMP * 0.225;}
    else if (IMP >= 4087.66)
    {total = IMP * 0.275;}

    printf("\n\n");
    printf("Usuario: %s\n",nome);
    printf("Salário: %.2f\n", sal);
    printf("Imposto: %.2f reais\n\n",total);

    system("pause");
    return 0;
}
