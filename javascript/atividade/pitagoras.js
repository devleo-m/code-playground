
console.log("Teorema de Pitagoras")
console.log(calc(5,3))

function calc(b,c) {
    a = (b*b) + (c*c)
    a = Math.sqrt(a)
    return a
}