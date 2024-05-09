function funcaoSegundoGrau(a, b, c, x) {
    return a * x * x + b * x + c;
}
  
function calcularRaizes(a, b, c) {
    const delta = b * b - 4 * a * c;
    if (delta < 0) {
      return 'Não existem raízes reais.';
    } else if (delta === 0) {
      const raiz = -b / (2 * a);
      return `Existe uma raiz real: ${raiz}`;
    } else {
      const raiz1 = (-b + Math.sqrt(delta)) / (2 * a);
      const raiz2 = (-b - Math.sqrt(delta)) / (2 * a);
      return `Existem duas raízes reais: ${raiz1} e ${raiz2}`;
    }
}
  
console.log(funcaoSegundoGrau(1, -5, 6, 3));
console.log(calcularRaizes(1, -5, 6)); 
  
console.log(funcaoSegundoGrau(1, 4, 4, -2)); 
console.log(calcularRaizes(1, 4, 4)); 
  