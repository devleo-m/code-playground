#include <stdio.h>
#include <stdlib.h>
#include <math.h>
#include <locale.h>

int main()
{
	setlocale(LC_ALL, ""); // acentuação PT-BR

	int calculo; // variavel p/ 'calcular' a distancia entre os dois pontos
	int laco = 0; // variavel p/ o 'laço*' de repetição

	//as structs pedidas na questão A
	struct ponto2d
	{
		int x, y;
	};
	struct ponto_inicial
	{
	}; struct ponto2d comeco; //começo*

	struct ponto_final
	{
	}; struct ponto2d fim;
	// questao B


	while ( laco > 0 || laco < 10 ) // loop pra sempre iniciar o programa quando for finalizado. // não é possivel ultrapassar 10vzs

	{
		//questão C // menu*
		printf("*********************MENU***********************\n");
		printf("*[1] - Digitar os valores do primeiro ponto    *\n");
		printf("*[2] - Digitar os valores do segundo ponto     *\n");
		printf("*[3] - Mostrar a distancia entre os dois pontos*\n");
		printf("*[4] - Sair                                    *\n");
		printf("************************************************\n");
		scanf_s("%i", &laco);

		switch (laco)
		{
		case 1:
			printf("\nDigite o número do ponto X[1]: "); 
			scanf_s("%d", &comeco.x);
			printf("\nDigite o número do ponto Y[1]: ");
			scanf_s("%d", &comeco.y);
			break;
		case 2:
			printf("\nDigite o número do ponto X[2]: ");
			scanf_s("%d", &fim.x);
			printf("\nDigite o número do ponto Y[2]: ");
			scanf_s("%d", &fim.y);
			break;
		case 3:
			calculo = pow((comeco.x - fim.x), 2) + pow((comeco.y - fim.y), 2); //potência 
			calculo = sqrt(calculo); // raiz quadrada
			printf_s("\nA distância entre os dois pontos, é: %d\n", calculo);
			system("pause");
			break;
		case 4:
			printf("Sair\n\n");
			exit;
			break;
		default: printf_s("\nErro! Tente novamente...\n");
			system("pause");
		}
	}
	return 0;
}
