//STRINGS

//Comprimento da String
//Uma string em C# é na verdade um objeto que contém propriedades e métodos que podem executar certas operações em strings.
//Por exemplo, o comprimento de uma string pode ser encontrado com a propriedade Length:

string txt = "ABCDEFGHIJKLMNOPQRSTUVWXYZ";
Console.WriteLine("The length of the txt string is: " + txt.Length);

//ToUpper() e ToLower(), que retornam uma cópia da string convertida em letras maiúsculas(ToUpper) ou minúsculas(ToLower):
string mensagem = "HelloC";
Console.WriteLine(mensagem.ToUpper());
Console.WriteLine(mensagem.ToLower());

//Interpolacao de strings
//forma mais correta de se usar strings e tecnica clean code
//Observe que você não precisa se preocupar com espaços, como acontece com a concatenação:
string nome = "Arolnode";
string cidade = "Recife";
Console.WriteLine($"{nome} esta morando em {cidade}");
//Observe também que você deve usar o cifrão ($) ao usar o método de interpolação de string.

//Strings de acesso
//Você pode acessar os caracteres de uma string consultando seu número de índice entre colchetes [].
//Este exemplo imprime o primeiro caractere em minhaString
string myString = "texto";
Console.WriteLine(myString[0]); //Vai ser printado apenas a letra T
Console.WriteLine(myString[1]);  //  "e"
//Você também pode encontrar a posição do índice de um caractere específico em uma string usando o método IndexOf():
Console.WriteLine(myString.IndexOf("x"));  // "1"
Console.WriteLine(myString.IndexOf("t"));  // "2"

//Outro método útil é Substring(), que extrai os caracteres de uma string, começando na posição/índice do caractere
//especificado e retorna uma nova string. Este método é frequentemente usado junto com IndexOf() para obter a posição
//específica do caractere:   Substring()
nome = "John Doe";
int charPos = nome.IndexOf("D");
string lastNome = nome.Substring(charPos);
Console.WriteLine(lastNome);





