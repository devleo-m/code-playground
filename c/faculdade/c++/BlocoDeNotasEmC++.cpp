#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <locale.h>
#include <iostream>

int main()
{
	setlocale(LC_ALL, ""); // acentuação PT-BR

	char arquivo[50], frase[50];
	int num[5] = { 1,2,3,4,5 };

	FILE* arq_texto; // Ponteiro
	errno_t err;

	printf_s("Nome do arquivo: ");
	gets_s(arquivo);
	strcat_s(arquivo, 50 - strlen(arquivo) - 1, ".txt"); //Concatena com o nome do arquivo com .txt
	printf_s("Digite a frase: ");

	for (int i = 0; i < 5; i++)
    
	{
		err = fopen_s(&arq_texto, arquivo, "a"); //aprendi isso no video aula do prof ao vivo kkkkk
		printf_s("Escreva a %i frase: \n", frase[i]);
		frase[i] = toupper(frase[i]); //NÃO CONSIGO COLOCAR ESSE COMANDO!!! Tentei colocar tudo em maiusculo... PQP >:#
		gets_s(frase); 
		fprintf_s(arq_texto, "%s", frase);
		fclose(arq_texto); // Fecha um arquivo aberto pela função fopen.
	}

	system(arquivo);
	system("pause");
	return 0;
}
