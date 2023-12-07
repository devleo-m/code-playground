
#include <stdio.h>
#include <stdlib.h>
#include <locale.h>

enum todas_pecas { peao, torre, cavalo, bispo, rei, rainha, vazio}; // enum p/ contabilizar todas as peças do tabuleiro;

int main()
{
	setlocale(LC_ALL, ""); // acentuação PT-BR

	int tabuleiro[8][8] = { {1, 3, 0, 5, 4, 0, 2, 1},{1, 0, 1, 0, 0, 1, 0, 0},{0, 0, 0, 0, 1, 0, 6, 0},{1, 0, 0, 1, 1, 0, 0, 1},{0, 1, 0, 4, 0, 0, 1, 0},{0, 0, 3, 1, 0, 0, 1, 1},{1, 0, 6, 6, 0, 0, 1, 0},{1, 0, 5, 0, 1, 1, 0, 6} };
	//  matrix 8x8 com todas as posições do xadrez

	printf("\nAluno: Leonardo Madeira Barbosa da Silva\n");
	printf("\nRU: 3235609\n");
	printf("\n*** TABULEIRO DE XADREZ ***\n");
	printf("\n\n");

	int lin, col; // laço de repetição for p/ linha e coluna do tabuleiro
    
	for (lin = 0; lin < 8; lin++) // linha
	{
		for (col = 0; col < 8; col++) //coluna
		{
			printf("  %d ", tabuleiro[lin][col]); //mostrar o tabuleiro
		}
		printf("\n\n");
	}

	int cont[7] = { 0 }; // Contador(vetor) com 7 opções para inserir nela o valor de cada peças
	int i, j;

	for (i = 0; i < 8; i++) { // laço de repetição com i e j para inserir as peças no tabuleiro 8x8
    
		for (j = 0; j < 8; j++) {

			switch (tabuleiro[i][j])  // switch no laço de repetição for, para guiar cada peça em varias situações "caso" escolher um valor
			{
			case vazio: // Casas vazias
				cont[0]++; // este laço ira contabilizar cada variavel das peças do tabuleiro
				break;
			case peao:
				cont[1]++;
				break;
			case torre:
				cont[2]++;
				break;
			case cavalo:
				cont[3]++;
				break;
			case bispo:
				cont[4]++;
				break;
			case rainha:
				cont[5]++;
				break;
			case rei:
				cont[6]++;
				break;
			default:
				break;
			}
		}
	}

	printf("\n");
	printf("|-----------------------------------------------------------------------------------|\n");
	printf("| 0 - Vazio | 1 - Peão | 2 - Torre | 3 - Cavalo | 4 - Bispo | 5 - Rainha | 6 - Rei  |\n");
	printf("| Vazias = %d                                                                        |\n", cont[0]);
	printf("| Peoes  = %d                                                                       |\n", cont[1]);
	printf("| Torres = %d                                                                       |\n", cont[2]);
	printf("| Cavalos= %d                                                                        |\n", cont[3]);
	printf("| Bispos = %d                                                                        |\n", cont[4]);
	printf("| Rainhas= %d                                                                        |\n", cont[5]);
	printf("| Reis   = %d                                                                        |\n", cont[6]);
	printf("|-----------------------------------------------------------------------------------|\n\n\n");
	system("pause");

	int tabuleiroreal[8][8] = { // tabuleiro real que vamos modificar ou tentar kkk

	{1, 3, 0, 5, 4, 0, 2, 1},
	{1, 0, 1, 0, 0, 1, 0, 0},
	{0, 0, 0, 0, 1, 0, 6, 0},
	{1, 0, 0, 1, 1, 0, 0, 1},
	{0, 1, 0, 4, 0, 0, 1, 0},
	{0, 0, 3, 1, 0, 0, 1, 1},
	{1, 0, 6, 6, 0, 0, 1, 0},
	{1, 0, 5, 0, 1, 1, 0, 6}
	};

	int linha, coluna, decida; // variaveis onde o jogador vai decidir onde vai colocar as peças do tabuleiro

	printf("\n\n");
	printf("tabuleiro do jogador!\n\n");

	// este é o tabuleiro onde podemos mudar com a decisão que escolhemos

	for (linha = 0; linha < 8; linha++) // laço de repetição for p mostrar o tabuleiro da linha
	{
		for (coluna = 0; coluna < 8; coluna++)
		{
			printf("%d ", tabuleiroreal[linha][coluna]);
		}
		printf("\n");
	}

	// peças do jogador.

	printf("\nescolha a posição da LINHA 1 a 8:\n"); // para modificar a LINHA  da tabela do tabuleiro do jogo

	scanf_s("%d", &linha);

	while ((linha < 1) || (linha > 8))
	{
		printf("Número inválido!!! escreva um número 1 a 8, zé ruela!\n"); // caso o jogador digite um numero que não seja de 1 a 8
		scanf_s("%d", &linha);
	}

	printf("escolha a posição da COLUNA 1 a 8:\n"); // Para modificar a COLUNA da tavela do tabuleiro do jogo
	scanf_s("%d", &coluna);

	while ((coluna < 1) || (coluna > 8))
	{
		printf("Número inválido!!! escreva um número 1 a 8, zé ruela!\n"); //caso o jogador digite um numero que não seja de 1 a 8
		scanf_s("%d", &coluna);
	}
	
	return 0;
}
