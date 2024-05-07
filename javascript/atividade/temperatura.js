let convertCpF = (c) => {
    f = c * (9/5) + 32
    return f.toFixed(1)
}

console.log("Convertendo Celsius para Fahrenheit")
console.log(`F°: ${convertCpF(22.9)}`)
