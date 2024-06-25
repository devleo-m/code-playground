/*
Operadores
Os operadores aritméticos permitem realizar operações matemáticas básicas:

+: Adição
-: Subtração
*: Multiplicação
/: Divisão
%: Módulo (resto da divisão)
*/

let a = 10;
let b = 5;

console.log(a + b); // 15
console.log(a - b); // 5
console.log(a * b); // 50
console.log(a / b); // 2
console.log(a % b); // 0

/*
Atividade 4:
4.1 Crie duas variáveis x e y com valores numéricos.
4.2 Realize e exiba no console as operações de adição, subtração, multiplicação, divisão, potencia e módulo entre x e y.
4.2.1 Calcule tudo dentro de um unico console cada operacao dentro de "{}"
*/

let x = 8
let y = 8
console.log(`X=${x} & Y=${y} soma: ${x + y} subtraacao: ${x - y} divisao: ${x / y} multiplicacao: ${x * y} resto da divisao: ${x % y} potencia: ${x ** y} `)

/*
Operadores de Comparação
Os operadores de comparação são usados para comparar valores e retornam um valor booleano (true ou false):

== Igual a
=== Estritamente igual a (valor e tipo)
!= Diferente de
!== Estritamente diferente de (valor e tipo)
> Maior que
>= Maior ou igual a
< Menor que
<= Menor ou igual a
*/

let i = 15;
let j = 8;

console.log(i == j); // false
console.log(i === 15); // true
console.log(i != j); // true
console.log(i !== "15"); // true
console.log(i > j); // true
console.log(i >= 10); // true
console.log(i < j); // false
console.log(i <= 15); // true