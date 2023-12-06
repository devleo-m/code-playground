//Concatenacao
//Nesse guia voce vai conhecer todasa as formas de concatenar e qual a melhor

//Usando operador +
let primeironome = `Jhon`
let segundonome = `Doe`
let fullname = primeironome + ' ' + segundonome
console.log('Concatenacao com operador +: '+ fullname)

//usando o metodo concat()
let inf1 = 'Hello';
let inf2 = 'World';
let fullConcat = inf1.concat(' ', inf2);
console.log('Concatenacao com o metodo concat:', fullConcat);

//Melhor forma para concatenar:
//Usando template literals (interpolação de string)
let age = 39;
console.log(`Olá, meu nome é ${primeironome} ${segundonome} e eu tenho ${age} anos.`);
//Obs: usamos ``(crase) ao invez de "" ou ''
/* para colocar a variavel e necessario colocar ${} e dentro das 
chaves o nome da variavel */



