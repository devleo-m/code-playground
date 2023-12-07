#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <locale.h>
#define a 5
#define b 3
int main()

{
 setlocale(LC_ALL, ""); // PARA FUNCIONAR ACENTUAÇÃO
 char m[a][b], c;// VARIAVEL CHAR
 int n1; // VARIAVEL INTEIRO

 for (int i = 0; i < a; i++) { // LAÇO DE REPETIÇÃO PARA LINHA
 for (int j = 0; j < b; j++) // LAÇO DE REPETIÇÃO P COLUNA
 {

 printf("Digite o valor de m[%i][%i]\n",i,j); // ESCREVA
 scanf_s("%c", &m[i][j]); // LEIA A VARIAVEL
 while ((c = getchar()) != '\n' && c != EOF) {} //arruma o bag do teclado
 }

 }

 n1 = (int)m[0][0] * 2; // ESTA SENDO MULTIPLICADO POR 2 JA QUE A LETRA É MINUSCULA
 for (int i = 0; i < a; i++) { // LAÇO DE REPETIÇÃO PARA LINHA
 for (int j = 0; j < b; j++) // '' COLUNA

 {
 printf("_______________________"); // UM ESPAÇO PRA FICAR MANEIRO :P
 printf("\n");
 printf("%c", m[i][j]); // SO PRA APARECER A LETRAS RESPECTIVAMENTE
 printf("\n");
 }

 }
 printf("\n\n");

 printf("A multiplicação da caractere m[0][0] = %d\n\n", n1); // O RESULTADO FINAL DA MULTIPLICAÇÃO P TELA
 system("pause");

 return 0;
}