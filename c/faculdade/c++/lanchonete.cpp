# include <stdio.h>
# include <stdlib.h>
# include <locale.h>
int main()
{
setlocale(LC_ALL, ""); // Comando para funcionar acentuação
int item, x = 0; // Variável inteiro
float quantd, valortotal = 0, pedido[10][7]; // Variável real e matriz
printf("---------------------------------------------------------\n");
printf("|Item | Produto | Código | Preço Unitário |\n");
printf("| 1 | CACHORRO-QUENTE | 100 | 5,00 |\n");
printf("| 2 | X-SALADA | 101 | 8,79 |\n");
printf("| 3 | X-BACON | 102 | 9,99 |\n");
printf("| 4 | MISTO | 103 | 6,89 |\n");
// Da linha 12 a 19 é o cardápio do restaurante
printf("| 5 | SALADA | 104 | 4,80 |\n");
printf("| 6 | ÁGUA | 105 | 3,49 |\n");
printf("| 7 | REFRIGERANTE | 106 | 4,99 |\n");
printf("---------------------------------------------------------\n");
do // Laço de repetição para repetir a quantidade do pedido 
{
 printf("\nEscolha um item: "); // Deve ser escrito o número do item 1 ao 7
 scanf_s("%d",&item); // Declara a variável do item
 if (item >= 1 && item <= 7) // se for verdadeiro
{
printf("Quantidade : "); // Quantidade desejada do item
scanf_s("%f", &quantd); // Variável da quantidade
if (quantd <= 0){ // Não existe comprei 0 saladaentão o pedido será cancelado
printf("Pedido cancelado!\n\n");
}
}
switch (item) // Montagem da lista de pedidos
{
case 1: // caso 1 se for cachorro-quente
pedido[x][0] = item; // item : cachorro-quente
pedido[x][2] = quantd; // quantidade do item
pedido[x][3] = 5.00; // valor do item em real
x++; // x = x+1
break; // fechando o caso 1
case 2:
pedido[x][0] = item; // item: X-salada
pedido[x][2] = quantd; // quantidade do item
pedido[x][3] = 8.79; // valor do item em real
x++; // x = x+1
break; // fechando o caso 1
case 3:
pedido[x][0] = item; //item : X-bacon
pedido[x][2] = quantd; // quantidade do item
pedido[x][3] = 9.99; // valor do item em real
x++; // x = x+1
break; // fechando o caso 1
case 4:
pedido[x][0] = item; // item : misto
pedido[x][2] = quantd; // quantidade do item
pedido[x][3] = 6.89; // valor do item em real
 x++; // x = x+1
break; // fechando o caso 1
case 5:
pedido[x][0] = item; // item: Salada
 pedido[x][2] = quantd; // quantidade do item
pedido[x][3] = 4.80; // valor do item em real
x++; // x = x+1
break; // fechando o caso 1
case 6:
pedido[x][0] = item; // item: Água
pedido[x][2] = quantd; // quantidade do item
pedido[x][3] = 3.49; // valor do item em real
x++; // x = x+1
break; // fechando o caso 1
case 7:
pedido[x][0] = item; // item: Refrigerante
pedido[x][2] = quantd; // quantidade do item
pedido[x][3] = 4.99; // valor do item em real
x++; // x = x+1
break; // fechando o caso 1
}
} while (item >= 1 && item <= 7); // o item tem que ser maior ou igual a 1 E item tem q ser menor ou igual a 7 senão o pedido será finalizado
printf("\n\n___________________________________________________________________________________\n\n\n"); // Apos finalizar o pedido eu coloquei essa barra pra dividir na tela o menu com o pedido e lista da compra
printf("\nLista dos itens comprados: \n\n");
printf("\t--------------------------------------");
printf("\n\t|ITENS|QUANTIDADE| PREÇO | VALOR |\n"); // coloquei esses ||| pra ficar legal pra dividir cada elemento e o \t pra dar um enter
for (int j = 0; j < x; j++) // laço de repetição PARA
{
printf("\t| %.0f | %.0f | %.2f | %.2f |\n", pedido[j][0],
pedido[j][2], pedido[j][3], (pedido[j][2] * pedido[j][3])); // lembrando que o pedido j 2 e pedido j 3 é a quantidade * item, calculando o valor de cada item
printf("\t--------------------------------------\n");
valortotal = valortotal + (pedido[j][2] * pedido[j][3]); // calculo geral de tudo para a variavel valortotal
}
printf("\n\nValor total a se pagar: %.2f reais\n", valortotal); // Valor total
printf("\n\n___________________________________________________________________________________\n\n\n");
system("pause");
return 0;
}