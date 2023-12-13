#include <stdio.h>
#include <stdlib.h>
#include <math.h>

void calc_hexa(float L, float* area, float* perimetro) // exigido na questao
{

	*area = (3 * pow(L, 2) * sqrt(3)) / 2; // ponteiro da area e formula da questao p calcular a area
	*perimetro = L * 6;					 // ponteiro perimetro

}

int main() {

	float area, perimetro, L;
	int laco = 0;

	while (laco > 0 || laco < 10) // loop infinito.
	{

		printf("Entre com o lado do Hexagono:");  //Digitar o lado Hexagono
		scanf_s("%f", &L);  // leia do lado hexagono

		calc_hexa(L, &area, &perimetro);  // calculo hexagono com a formula *area e *perimetro da linha 7 e 8
		printf("Area: %f\nPerimeto: %f\n", area, perimetro); // imprimir o resultado das duas formulas da questão

	}

	return 0;
}
