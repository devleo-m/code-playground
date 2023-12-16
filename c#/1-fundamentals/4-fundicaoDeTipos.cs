//Fundicao de tipos

//fundicao implicita:
int number = 50;
double number2 = number;

Console.WriteLine(number);
Console.WriteLine(number2);

//fundicao explicita:

double number3 = 1.99;
int number4 = (int)number3;

Console.WriteLine(number3);
Console.WriteLine(number4);

//Metodos de conversao de tipo
//tembem eh possivel converter tipos de dados explicitamente usando metodos integrados
//convert.toboolean/convert.todouble/convert.tostring/convert.toint32 (int) e convert.toint64 (long)

int number5 = 10;
double number6 = 5.25;
bool _true = true;

Console.WriteLine(Convert.ToDouble(number5));
Console.WriteLine(Convert.ToInt32(number6));
Console.WriteLine(Convert.ToString(_true));
