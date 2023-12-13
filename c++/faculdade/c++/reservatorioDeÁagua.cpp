#include <stdio.h>
#include <stdlib.h>
#include <locale.h>
int main() {
setlocale(LC_ALL, ""); // PARA FUNCIONAR ACENTUAÇÃO
int chuva[99], reserv, gasto[99];
int reservfull;
char mes[12][10] = { { 'J', 'a', 'n', 'e', 'i', 'r', 'o', '\0' }, { 'F',
'e', 'v', 'e', 'r', 'e' , 'i', 'r', 'o','\0' }, { 'M', 'a', 'r', 'ç', 'o','\0' },
{ 'A', 'b', 'r', 'i', 'l','\0' }, { 'M', 'a', 'i', 'o', '\0' },{ 'J', 'u', 'n',
'h', 'o', '\0' },{ 'J', 'u', 'l', 'h', 'o', '\0' }, { 'A', 'g', 'o', 's', 't',
'o', '\0' }, { 'S', 'e', 't', 'e', 'm', 'b', 'r', 'o', '\0'}, { 'O', 'u', 't',
'u', 'b', 'r', 'o', '\0' },{ 'N', 'o', 'v', 'e', 'm','b', 'r', 'o', '\0' }, {
'D', 'e', 'z', 'e', 'm','b', 'r', 'o', '\0' } };
// variável int para os cálculos e char para cada mês.
printf("Digite o valor máximo do reservatório: "); // Escrever o limite do reservatório
scanf_s("%d", &reservfull); // ler a variável do reservatório
for (int i = 0; i <= 11; i++) { // laço de repetição PARA registar a chuva de cada mês OBS SE EU COLOCAR I<=12 VAI APARECER UM 13 MES, ACHEI ESTRANHO E COLOQUEI 11, MESMO ASSIM FUNCIONOU :p
printf("Quantidade da chuva do mês %s: ", mes[i]); // aparece na tela p/ digitar o registro de cada chuva 12x de cada mês
scanf_s("%d", &chuva[i]); // ler a variável dos 12 meses do registro de chuva
}
for (int i = 0; i < 12; i++) { // novamente laço de repetição PARA registar
printf("Quantidade de água gasta de cada mês %s: ",
mes[i]);//Aparece na tela e p/ digitar a quantidade de agua gasta
scanf_s("%d", &gasto[i]);//ler a variável 12 X de cada mês
}
printf("_________________________________________\n");// um espaço que gosto da dar, nao leve a mal... lol
printf("Relatório do reservatorio de cada mês\n\n");
printf(" Mês Chuva Gasto Armazenado \n");// eu queria fazer uma tabela, mas ia ficar brega igual a atividade 1, então não coloquei aquelas linhas aqui
for (int i = 0; i < 12; i++) {// mais um laço de repetição PARA calcular o reservatório = chuva+ gasto de agua
printf("%.3s %d mm %d mm %d mm \n", mes[i], chuva[i], gasto[i], reserv); // eu dei esses espaço pra fazer uma tabela simples que de para o senhor entender reserv = (chuva[i] - gasto[i]) + reserv; // por esse motivo que criei a laço de rpetição PARA.
if (reserv >= reservfull) { // Se o reservatório for maior ou igual ao limite do reservatório ENTAO o reservatório esta cheio
printf("Reservatorio cheio\n %d mm\n", reserv);
}
if (reserv <= 0) {// Se o reservatório for menor ou igual a 0 ENTAO o reservatório esta VAZIO!! ^^
printf("Reservatório Vazio: %d mm\n", reserv);}
}
system("pause");
return 0;
}