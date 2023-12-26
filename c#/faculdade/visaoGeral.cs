//Introducao c# pela faculdade:

//variaveis
//No C#, a declaração de variáveis está fortemente atrelada ao type dela, ou seja, uma vez que declaramos um tipo numérico,
//não podemos atribuir um valor booleano nele, por exemplo.

//Booleans:
using System.Runtime.Intrinsics.X86;
bool a = true;
bool b = false;

//int (numeros inteiros: 1, 2, 10, 30, 1000)
int num1 = 5;
int num2 = 22;
int calc = num1 * num2; // 5 * 22 = 110
uint uv3 = 10 + 10;
Console.WriteLine(calc);
Console.WriteLine(uv3);

//long
long l1 = 100;
long l2 = 200;
long resultL = (l1 + l2) * (9 / 3); // 900
Console.WriteLine(resultL);

//float (Numeros flutuantes: 1.99, 5.75, 9.99, 4.32)
float f1 = 1.99f; //note que sempre usamos um f no final do numero quando usamos float
float f2 = 5.75f;

//double (number do tipo float)
//mais comum em se usar no c#
double d1 = 9.99;
double d2 = 77.3;
//nao e necessario colocar f no final de um double

//calculando float com double:
double calcAll = (f1 * d2) + (d1 - f2);
Console.WriteLine(calcAll);
//note que misturamos float com double e no final tudo virou double sem problemas

//char
char sala1 = 'b';
char sala2 = 'c';
char sala3 = 'a';
char exec = (char)33;
//nao eh possivel somar dois caracteres, pois isso corresponde a uma string
//entao vamos inicializar uma string a partir de uma array de caracteres

string resultChar = new string(new char[] { sala1, sala2, sala3, exec });
// escrevendo no console a array char
Console.WriteLine(resultChar);

//objetos
//declaracao de referencia types, ou seja, objetos:
//string
string s1 = "Ola";

string s2 = new string('x', 3); // "xxx" sera repetido 3 vezes o x
string sAll = s1 + s2; // Ola xxx
Console.WriteLine(sAll);

//objetos
object obj1 = new object();
Console.WriteLine(obj1.ToString()); // System.Object //objeto vazio com um valor default
object obj2 = sAll;
Console.WriteLine(obj2); // ola xxx

//if else
//Blocos de decisão – IF / ELSE.
//No C#, podemos executar código condicional utilizando a construção if:

int num10 = 10;
int num20 = 20;
if (num20 >= 20)
{
    Console.WriteLine("maior q 20");
}
else
{
    Console.WriteLine("menor que 20");
}

int num30 = num10 + num20;

if (num30>=30)
{
    Console.WriteLine("Maior que 30");
}
else if (num30 > 10)
{
    Console.WriteLine("Maior que 10");
}
else
{
    Console.WriteLine("Menor que 10");
}

//Switch/Case.
//A sintaxe de switch/case do C# é igual à do Java. Variando apenas em opções diferentes que o bloco Switch pode receber para condicionar.
//No C# temos variações do Switch que aceitam strings, pattern matching, entre outras.
int teste = 10;
switch (teste)
{
    case 10:
        Console.WriteLine("10"); //A variavel teste eh 10, entao sera executado apenas essa casa
        break; //se a break for retirada, nesta linha, sera executada a case 9 tambem
    case 9:
        Console.WriteLine("9");
        break; // o break serve como uma string para finalizar o codigo nesta linha caso a case for executada
    case 8:
        Console.WriteLine("8");
        break;
    case 7:
        Console.WriteLine("7");
        break;
    default:
        Console.WriteLine("Valor invalido!"); //se a variavel teste for algum numero que nao esteja em uma case, sera executado esse default
        break;
}

//Blocos de loop – for.

for (int i = 0; i < 10; i++) //i++ eh uma abreviacao de i = i + 1
{ //i = 0, i sera executado ate chegar no valor 10
    Console.WriteLine(i); //sera executado de 0 a 10
}
for (int i = 10; i > 0; i--)//i-- eh uma abreviacao de i = i - 1
{
    Console.WriteLine(i);
}

//O bloco for no C# é igual às demais linguagens imperativas. Recebendo uma variável numérica, estabelecendo um ponto de
//checagem/parada e um ponto de incremento da variável numérica.

//foreach
string nome = "Arolnode storm +9 33aa";
foreach (var item in nome)
{
    Console.WriteLine(item);
}

//foreach eh mais comum sendo utilizado em arrays, vamos estudar sobre isso depois
//A sintaxe do bloco foreach no C# é igual às demais linguagens imperativas. No caso do C#, recebemos uma variável que implementa a
//interface IEnumerable. A partir disso, o código irá iterar sobre itens desta variável (vamos abordar isso de forma mais detalhada
//no próximo tema).

//No exemplo, o código iterou sobre os caracteres que compõem nossa string armazenada na variável stringForLoop.

//while.
int number9 = 9;
while (number9 <= 15)
{
    Console.WriteLine(number9);
    number9++; //sera executado o 9 ate ele chegar no valor 15
}

string cool = "c# eh legal"; // Inicializa a variável cool com a string "c# eh legal"
while (cool.Length > 0) // Enquanto o comprimento da string cool for maior que zero, o loop continuará
{
    Console.WriteLine(cool); // Imprime a string cool atual
    cool = cool.Substring(0, cool.Length - 1); // Atualiza a string cool para ser uma substring dela mesma, removendo o último caractere a cada iteração
    if (cool.EndsWith("#")) // Verifica se a string cool termina com o caractere '#'
    {
        cool = string.Empty; // Se a string terminar com '#', define cool como uma string vazia para sair do loop
    }
}

//Em resumo, esse código começa com a string "c# eh legal" e, em cada iteração do loop while, ele imprime a string atual, remove o
//último caractere dela e verifica se o último caractere é #. Se for, a string cool se torna vazia, encerrando o loop.

//PROPRIEDADES, MODIFICADORES DE ACESSO, INTERFACES E OUTROS
//No C#, temos um “tipo de campo” chamado Propriedade (property), que é padrão da linguagem e fornece algumas características
//interessantes. O conceito de Propriedade em linguagens como o C# e Java é muito parecido, mudando pouco a forma de declaração
//e recursos das linguagens sobre as propriedades. Você notará que propriedades são a forma mais comum e prática de expor dados
//de suas classes (entre elas).

//Vamos conceituar primeiro o que são campos (fields), propriedades (properties) e métodos (method) no C#.
//A seguir, podemos ver a classe “Pessoa” com os três tipos:

class Pessoa // Declaração da classe Pessoa
{
    string nome; // Declaração de uma variável membro privada do tipo string chamada "nome"
    int Idade { get; set; } // Declaração de uma propriedade auto-implementada para a idade

    bool PodeAndar() // Declaração de um método privado que retorna um booleano chamado PodeAndar
    {
        return true; // Retorna sempre verdadeiro (true)
    }
}

//Propriedades no C# devem ser compostas por instruções de gets e/ou sets.Permitem criar um modelo para que a variável receba valores e/ou ceda seus valores.
//Uma propriedade precisa ter pelo menos um get ou um set. Eles possuem o seguinte comportamento:

//Get: é um trecho de código (similar a um método), que tem como objetivo retornar o valor da propriedade.
//Podendo fazer operações de checagem, atribuição de outros campos internos e inicialização de campos internos antes de retornar seu valor.
//Set: é um trecho de código (similar a um método), que tem como objetivo alterar o valor da propriedade.
//Geralmente modificando seu valor para o valor passado por parâmetro.

class Pessoa2 // Declaração da classe Pessoa2
{
    string Nome { get; set; } // Declaração de uma propriedade auto-implementada chamada Nome

    int idade; // Declaração de uma variável membro privada do tipo int chamada idade

    void SetIdade(int idadeDesejada) // Declaração de um método chamado SetIdade que recebe um parâmetro do tipo int
    {
        idade = idadeDesejada; // Define o valor da variável idade com base no parâmetro recebido
    }

    int GetIdade() // Declaração de um método chamado GetIdade que retorna um valor do tipo int
    {
        return idade; // Retorna o valor da variável idade
    }
}

//Podemos notar que o comportamento de alterar o valor do campo (field) idade depende do método “SetIdade” e “GetIdade”.
//Já na propriedade Nome, esse comportamento é automático.
//A linguagem gera (dentro do CIL) o get e set necessários para que a propriedade tenha o mesmo comportamento do campo idade,
//simplificando a necessidade de métodos para isso.

//O C# cria um campo e os métodos de forma totalmente transparente ao desenvolvedor e
//implementa o comportamento padrão para eles: sempre que se atribuir o valor na propriedade, ele atribui o mesmo
//valor/referência no campo, sempre que buscar (pegar) o valor da propriedade, ele retorna o mesmo valor/referência do campo.

//Podemos resumir do seguinte modo: as propriedades são abstrações para armazenar, recuperar e alterar o valor de um campo.

/*
private string _nomeCompleto = "de"; // uma variável privada para armazenar o valor

public string NomeCompleto
{
    get { return _nomeCompleto; } // retorna o valor da variável
    set { _nomeCompleto = value; } // define um novo valor para a variável
}
*/

//INTERFACES
//Uma interface é como uma classe, porém, apenas descreve comportamentos, ou seja, métodos e propriedades.
//Ela não pode implementar nenhum desses, apenas declará-los. Seu objetivo é descrever comportamentos comuns, para que classes
//que implementam a interface possam utilizar o conceito de abstração da orientação a objetos.
interface IComestivel
{
    void Mastigar(object comida);
    int MaximaComida { get; }
    string Descrever()
    {
        return "possibilita o comportamento de mastigar um alimento!";
    }
}

//É uma convenção no C# usarmos a leta “I” maiúscula na frente do nome da interface para diferenciá-la de uma classe.

//# MODIFICADORES DE ACESSO
//Para promover encapsulamento, um tipo (type) ou membro de um tipo pode limitar sua acessibilidade a outros tipos
//(classes, variáveis e métodos) adicionando um dos cinco modificadores de acesso existentes antes de sua declaração

//# public
//Totalmente acessível. Esse é o comportamento padrão de interfaces e enumeradores.
//Torna acessível uma classe ou membro desta a outras classes e membros.

//# internal
//Acessível somente dentro do mesmo “assembly” (dll) ou métodos da mesma classe. Esse é o tipo padrão de acessibilidade
//para “non-nested types” ou tipos não aninhados no c#.

//# private
//Acessível somente ao tipo ou classe onde está contido. Esse é o tipo padrão de membros em classes ou structs.

//# protected
//Acessível somente dentro do tipo/classe ou dentro de suas subclasses. Por exemplo, pode ser acessado dentro da
//classe mãe e dentro de todas as classes que herdam ela.

//# protected internal
//É a junção do protected com internal. Tornar o tipo acessível a todas as suas subclasses e a todos os outros tipos
//dentro do mesmo “assembly” (dll) ao mesmo tempo


//OBJECTS, STRCUTS E TYPES


