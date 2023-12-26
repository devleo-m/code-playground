/*JavaScript Variables

Description:
In JavaScript, variables are used to store and manage data. 
They are containers for holding values, and their names are used to 
reference or manipulate these values in your code. Understanding how 
to declare, assign, and use variables is fundamental to writing effective 
JavaScript code. This guide covers the basics of JavaScript variables, 
including data types, declaration keywords (var, let, const), and best 
practices for naming conventions.
*/

var num1 = 5
console.log(num1);

let num2 = 10;
console.log(num2)

const num3 = 15
console.log(num3)

/*
In modern JavaScript development practices, it is advisable to avoid the 
use of var due to issues related to global scope and the concept of 
hoisting. Instead, it is preferable to use let and const, introduced in 
ECMAScript 6 (ES6), due to their more predictable and restrictive behaviors.

Here are some general guidelines for the use of let and const:

let: Use let when declaring variables that will need to be reassigned during the program's execution. The scope of a variable declared with let is limited to the block, statement, or expression in which the variable is declared.

*/

let count = 0;
for (let i = 0; i < 5; i++) {
  count += i;
}
console.log(count); // i is not accessible here

/*
const: Use const when declaring variables that will not be reassigned 
after their initialization. This helps ensure immutability whenever 
possible.
*/

const pi = 3.14;

/*
If you need a variable that won't be reassigned but may have its values 
changed (such as an object or array), you can use const for the variable 
reference and modify its attributes or elements.
*/

const person = {
  name: 'John',
  age: 30,
};
person.age = 31; // This is allowed


//Using let and const contributes to safer, more readable, and 
//easier-to-maintain code in JavaScript.


