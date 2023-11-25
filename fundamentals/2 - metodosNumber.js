let height = 1.80;
let height2 = Number("1.60"); // The Number function converts a string to a number
let boolTrue = Number(true);
let boolFalse = Number(false);

// Displaying heights using template literals
console.log(`Height: ${height} and ${height2}\n`);
// This technique is called concatenation, and we'll learn more about it later.
// Just understand that with the Number method, I converted the string 'height2' to a number.

// Displaying boolean conversions
console.log(`true: ${boolTrue}\nfalse: ${boolFalse}`);
// Note that in numbers, true is always 1, and false is 0.

// Number.isNaN = checks if the variable is NaN and returns true or false
let result1 = Number.isNaN(10); // Returns false
let result2 = Number.isNaN(NaN);
console.log(`Is NaN: ${result2} and not ${result1}`);

// Number.isInteger(value) = Checks if the value is an integer.
let result3 = Number.isInteger(3.14); // Returns false

// Number.parseFloat(string) = Parses a string and returns a floating-point number.
let result4 = Number.parseFloat("3.14"); // Returns 3.14

// Number.parseInt(string, radix) = Parses a string and returns an integer.
let result5 = Number.parseInt("10", 2); // Returns 2 (in binary)

// Number.toFixed(digits) = Returns a string representing the number with a specific number of decimal places.
let num1 = 3.14159;
let result6 = num1.toFixed(2); // Returns "3.14"
