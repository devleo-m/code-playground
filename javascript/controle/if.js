function apenasNotaBoa(nota) {
    if (nota >= 7) {
        console.log("Nota boa: ".concat(nota))
    }
}

function apenasNotaRuim(nota) {
    if (nota <= 5) {
        console.log(`Nota ruim: ${nota}`)
    }
}

let resultadoFinal = (nota) => {
    if (nota >= 6) {
        console.log("Voce passou, nota boa: "+ nota)
    } else{
        console.log("Reprovado, nota ruim: "+ nota)
    }
}

apenasNotaBoa(8.1)
apenasNotaBoa(4.2)

apenasNotaRuim(7.3)
apenasNotaRuim(2.7)

resultadoFinal(9)
resultadoFinal(1)

const informacaoVerdadeira = (inf) =>{
    if (inf) {
        console.log("Essa mensagem eh verdadeira")
    }
}

informacaoVerdadeira(true) //true
informacaoVerdadeira(false) //false
informacaoVerdadeira() //false
informacaoVerdadeira(null) //false
informacaoVerdadeira(undefined) //false
informacaoVerdadeira(NaN) //false
informacaoVerdadeira('') //false
informacaoVerdadeira(0) //false
informacaoVerdadeira(1) //true
informacaoVerdadeira(-1) //true
informacaoVerdadeira(' ') //true
informacaoVerdadeira(" ?") //true
informacaoVerdadeira([]) //true
informacaoVerdadeira([1,2]) //true
informacaoVerdadeira({}) //true
