package main

import (
	"fmt"
	"math/rand"
	"time"
)

// ============================================
// EXEMPLO 1: if Básico
// ============================================

func exemploIfBasico() {
	fmt.Println("=== Exemplo 1: if Básico ===")

	idade := 20

	// if simples - sem parênteses, mas chaves obrigatórias
	if idade >= 18 {
		fmt.Println("Você é maior de idade")
	}

	// if com condição falsa - nada acontece
	if idade < 18 {
		fmt.Println("Você é menor de idade")
	}
}

// ============================================
// EXEMPLO 2: if com Inicialização Opcional
// ============================================

func exemploIfComInicializacao() {
	fmt.Println("\n=== Exemplo 2: if com Inicialização ===")

	// Inicialização antes da condição
	if num := 42; num > 40 {
		fmt.Printf("O número %d é maior que 40\n", num)
		// num está disponível apenas dentro do bloco if
	}

	// Variável num não existe aqui fora do if
	// fmt.Println(num) // ERRO: undefined: num
}

// ============================================
// EXEMPLO 3: if-else
// ============================================

func exemploIfElse() {
	fmt.Println("\n=== Exemplo 3: if-else ===")

	temperatura := 25

	if temperatura > 30 {
		fmt.Println("Está muito quente!")
	} else {
		fmt.Println("Temperatura agradável")
	}
}

// ============================================
// EXEMPLO 4: if-else if-else (Múltiplas Condições)
// ============================================

func exemploIfElseIf() {
	fmt.Println("\n=== Exemplo 4: if-else if-else ===")

	nota := 85

	if nota >= 90 {
		fmt.Println("Nota A - Excelente!")
	} else if nota >= 80 {
		fmt.Println("Nota B - Muito bom!")
	} else if nota >= 70 {
		fmt.Println("Nota C - Bom")
	} else if nota >= 60 {
		fmt.Println("Nota D - Regular")
	} else {
		fmt.Println("Nota F - Reprovado")
	}
}

// ============================================
// EXEMPLO 5: if com Operadores Lógicos
// ============================================

func exemploIfComOperadores() {
	fmt.Println("\n=== Exemplo 5: if com Operadores Lógicos ===")

	salario := 3000
	horasTrabalhadas := 45

	// Operador && (E)
	if salario > 2500 && horasTrabalhadas > 40 {
		fmt.Println("Você recebeu horas extras!")
	}

	// Operador || (OU)
	if salario < 2000 || horasTrabalhadas < 30 {
		fmt.Println("Atenção: salário ou horas abaixo do esperado")
	}

	// Operador ! (NÃO)
	chovendo := false
	if !chovendo {
		fmt.Println("Pode sair sem guarda-chuva")
	}
}

// ============================================
// EXEMPLO 6: switch Básico
// ============================================

func exemploSwitchBasico() {
	fmt.Println("\n=== Exemplo 6: switch Básico ===")

	dia := "segunda"

	switch dia {
	case "segunda":
		fmt.Println("Início da semana")
	case "terça", "quarta", "quinta":
		fmt.Println("Meio da semana")
	case "sexta":
		fmt.Println("Final da semana")
	case "sábado", "domingo":
		fmt.Println("Fim de semana!")
	default:
		fmt.Println("Dia inválido")
	}
}

// ============================================
// EXEMPLO 7: switch sem Expressão (como if-else)
// ============================================

func exemploSwitchSemExpressao() {
	fmt.Println("\n=== Exemplo 7: switch sem Expressão ===")

	idade := 25

	switch {
	case idade < 18:
		fmt.Println("Menor de idade")
	case idade >= 18 && idade < 65:
		fmt.Println("Adulto")
	case idade >= 65:
		fmt.Println("Idoso")
	}
}

// ============================================
// EXEMPLO 8: switch com Inicialização
// ============================================

func exemploSwitchComInicializacao() {
	fmt.Println("\n=== Exemplo 8: switch com Inicialização ===")

	// Inicialização antes do switch
	switch hora := time.Now().Hour(); {
	case hora < 12:
		fmt.Println("Bom dia!")
	case hora < 18:
		fmt.Println("Boa tarde!")
	default:
		fmt.Println("Boa noite!")
	}
}

// ============================================
// EXEMPLO 9: switch com fallthrough
// ============================================

func exemploSwitchFallthrough() {
	fmt.Println("\n=== Exemplo 9: switch com fallthrough ===")

	numero := 2

	switch numero {
	case 1:
		fmt.Println("Um")
		fallthrough // Continua para o próximo case
	case 2:
		fmt.Println("Dois")
		fallthrough
	case 3:
		fmt.Println("Três")
		// Sem fallthrough, para aqui
	default:
		fmt.Println("Outro número")
	}
	// Saída: Dois, Três
}

// ============================================
// EXEMPLO 10: Type Switch
// ============================================

func exemploTypeSwitch() {
	fmt.Println("\n=== Exemplo 10: Type Switch ===")

	var valor interface{} = 42

	switch v := valor.(type) {
	case int:
		fmt.Printf("É um inteiro: %d\n", v)
	case string:
		fmt.Printf("É uma string: %s\n", v)
	case bool:
		fmt.Printf("É um booleano: %v\n", v)
	default:
		fmt.Printf("Tipo desconhecido: %T\n", v)
	}
}

// ============================================
// EXEMPLO 11: Comparação: if-else vs switch
// ============================================

func exemploComparacao() {
	fmt.Println("\n=== Exemplo 11: Comparação if-else vs switch ===")

	// Usando if-else
	cor := "vermelho"
	if cor == "vermelho" {
		fmt.Println("Pare!")
	} else if cor == "amarelo" {
		fmt.Println("Atenção!")
	} else if cor == "verde" {
		fmt.Println("Siga!")
	} else {
		fmt.Println("Cor desconhecida")
	}

	// Usando switch (mais limpo)
	switch cor {
	case "vermelho":
		fmt.Println("Pare!")
	case "amarelo":
		fmt.Println("Atenção!")
	case "verde":
		fmt.Println("Siga!")
	default:
		fmt.Println("Cor desconhecida")
	}
}

// ============================================
// EXEMPLO 12: if com Funções que Retornam Múltiplos Valores
// ============================================

func exemploIfComFuncoes() {
	fmt.Println("\n=== Exemplo 12: if com Funções ===")

	// Função que retorna valor e erro
	dividir := func(a, b float64) (float64, error) {
		if b == 0 {
			return 0, fmt.Errorf("divisão por zero")
		}
		return a / b, nil
	}

	// Verificar erro diretamente no if
	if resultado, err := dividir(10, 2); err != nil {
		fmt.Printf("Erro: %v\n", err)
	} else {
		fmt.Printf("Resultado: %.2f\n", resultado)
	}

	// Tentar dividir por zero
	if resultado, err := dividir(10, 0); err != nil {
		fmt.Printf("Erro: %v\n", err)
	} else {
		fmt.Printf("Resultado: %.2f\n", resultado)
	}
}

// ============================================
// EXEMPLO 13: switch com Expressões
// ============================================

func exemploSwitchComExpressoes() {
	fmt.Println("\n=== Exemplo 13: switch com Expressões ===")

	numero := 15

	switch {
	case numero%2 == 0:
		fmt.Println("Número par")
	case numero%2 == 1:
		fmt.Println("Número ímpar")
	}

	// switch com expressão e múltiplos casos
	switch numero {
	case 1, 3, 5, 7, 9, 11, 13, 15, 17, 19:
		fmt.Println("Número ímpar menor que 20")
	case 2, 4, 6, 8, 10, 12, 14, 16, 18, 20:
		fmt.Println("Número par menor que 20")
	default:
		fmt.Println("Número fora do range")
	}
}

// ============================================
// EXEMPLO 14: if Aninhado
// ============================================

func exemploIfAninhado() {
	fmt.Println("\n=== Exemplo 14: if Aninhado ===")

	idade := 25
	temCarteira := true

	if idade >= 18 {
		if temCarteira {
			fmt.Println("Pode dirigir")
		} else {
			fmt.Println("Precisa tirar carteira")
		}
	} else {
		fmt.Println("Menor de idade, não pode dirigir")
	}
}

// ============================================
// EXEMPLO 15: switch com Múltiplos Valores por Case
// ============================================

func exemploSwitchMultiplosValores() {
	fmt.Println("\n=== Exemplo 15: switch com Múltiplos Valores ===")

	mes := "janeiro"

	switch mes {
	case "dezembro", "janeiro", "fevereiro":
		fmt.Println("Verão")
	case "março", "abril", "maio":
		fmt.Println("Outono")
	case "junho", "julho", "agosto":
		fmt.Println("Inverno")
	case "setembro", "outubro", "novembro":
		fmt.Println("Primavera")
	default:
		fmt.Println("Mês inválido")
	}
}

// ============================================
// EXEMPLO 16: Validação com if
// ============================================

func exemploValidacao() {
	fmt.Println("\n=== Exemplo 16: Validação com if ===")

	email := "usuario@email.com"
	senha := "senha123"

	// Validação múltipla
	if email == "" {
		fmt.Println("Email não pode ser vazio")
	} else if len(senha) < 6 {
		fmt.Println("Senha muito curta")
	} else if len(email) < 5 {
		fmt.Println("Email inválido")
	} else {
		fmt.Println("Dados válidos!")
	}
}

// ============================================
// EXEMPLO 17: switch com Type Assertion
// ============================================

func exemploSwitchTypeAssertion() {
	fmt.Println("\n=== Exemplo 17: switch com Type Assertion ===")

	processar := func(valor interface{}) {
		switch v := valor.(type) {
		case int:
			fmt.Printf("Inteiro: %d (dobro: %d)\n", v, v*2)
		case float64:
			fmt.Printf("Float: %.2f (quadrado: %.2f)\n", v, v*v)
		case string:
			fmt.Printf("String: %s (tamanho: %d)\n", v, len(v))
		case []int:
			fmt.Printf("Slice de ints: %v (tamanho: %d)\n", v, len(v))
		default:
			fmt.Printf("Tipo não tratado: %T\n", v)
		}
	}

	processar(42)
	processar(3.14)
	processar("Go")
	processar([]int{1, 2, 3})
	processar(true)
}

// ============================================
// EXEMPLO 18: if com Short-Circuit Evaluation
// ============================================

func exemploShortCircuit() {
	fmt.Println("\n=== Exemplo 18: Short-Circuit Evaluation ===")

	// Função que pode causar erro
	obterValor := func() (int, error) {
		return 10, nil
	}

	// Short-circuit: se primeira condição for falsa, segunda não é avaliada
	valor := 5
	if valor > 10 && valor < 20 {
		fmt.Println("Valor entre 10 e 20")
	}

	// Com função que retorna erro
	if v, err := obterValor(); err == nil && v > 5 {
		fmt.Printf("Valor válido: %d\n", v)
	}
}

// ============================================
// EXEMPLO 19: switch para Classificação
// ============================================

func exemploClassificacao() {
	fmt.Println("\n=== Exemplo 19: switch para Classificação ===")

	classificarIdade := func(idade int) string {
		switch {
		case idade < 0:
			return "Idade inválida"
		case idade < 13:
			return "Criança"
		case idade < 18:
			return "Adolescente"
		case idade < 65:
			return "Adulto"
		default:
			return "Idoso"
		}
	}

	fmt.Println(classificarIdade(5))   // Criança
	fmt.Println(classificarIdade(15))  // Adolescente
	fmt.Println(classificarIdade(30))  // Adulto
	fmt.Println(classificarIdade(70))  // Idoso
}

// ============================================
// EXEMPLO 20: if com Map Lookup
// ============================================

func exemploMapLookup() {
	fmt.Println("\n=== Exemplo 20: if com Map Lookup ===")

	cores := map[string]string{
		"vermelho": "#FF0000",
		"verde":    "#00FF00",
		"azul":     "#0000FF",
	}

	cor := "vermelho"

	// Verificar se chave existe no map
	if hex, existe := cores[cor]; existe {
		fmt.Printf("Cor %s: %s\n", cor, hex)
	} else {
		fmt.Printf("Cor %s não encontrada\n", cor)
	}

	// Verificar se não existe
	if _, existe := cores["amarelo"]; !existe {
		fmt.Println("Amarelo não está no mapa")
	}
}

// ============================================
// EXEMPLO 21: switch com Constantes
// ============================================

func exemploSwitchConstantes() {
	fmt.Println("\n=== Exemplo 21: switch com Constantes ===")

	const (
		Segunda = iota
		Terca
		Quarta
		Quinta
		Sexta
		Sabado
		Domingo
	)

	dia := Terca

	switch dia {
	case Segunda, Terca, Quarta, Quinta, Sexta:
		fmt.Println("Dia útil")
	case Sabado, Domingo:
		fmt.Println("Fim de semana")
	default:
		fmt.Println("Dia inválido")
	}
}

// ============================================
// EXEMPLO 22: if com Slices
// ============================================

func exemploIfComSlices() {
	fmt.Println("\n=== Exemplo 22: if com Slices ===")

	numeros := []int{1, 2, 3, 4, 5}

	// Verificar se slice está vazio
	if len(numeros) == 0 {
		fmt.Println("Slice vazio")
	} else {
		fmt.Printf("Slice tem %d elementos\n", len(numeros))
	}

	// Verificar se tem elementos suficientes
	if len(numeros) >= 3 {
		fmt.Printf("Terceiro elemento: %d\n", numeros[2])
	}
}

// ============================================
// EXEMPLO 23: switch para Tratamento de Erros
// ============================================

func exemploTratamentoErros() {
	fmt.Println("\n=== Exemplo 23: switch para Tratamento de Erros ===")

	simularErro := func() error {
		erros := []error{
			fmt.Errorf("erro de conexão"),
			fmt.Errorf("erro de autenticação"),
			fmt.Errorf("erro de permissão"),
			nil,
		}
		return erros[rand.Intn(len(erros))]
	}

	err := simularErro()

	switch {
	case err == nil:
		fmt.Println("Operação bem-sucedida")
	case err.Error() == "erro de conexão":
		fmt.Println("Tente novamente mais tarde")
	case err.Error() == "erro de autenticação":
		fmt.Println("Verifique suas credenciais")
	case err.Error() == "erro de permissão":
		fmt.Println("Você não tem permissão")
	default:
		fmt.Printf("Erro desconhecido: %v\n", err)
	}
}

// ============================================
// MAIN
// ============================================

func main() {
	rand.Seed(time.Now().UnixNano())

	exemploIfBasico()
	exemploIfComInicializacao()
	exemploIfElse()
	exemploIfElseIf()
	exemploIfComOperadores()
	exemploSwitchBasico()
	exemploSwitchSemExpressao()
	exemploSwitchComInicializacao()
	exemploSwitchFallthrough()
	exemploTypeSwitch()
	exemploComparacao()
	exemploIfComFuncoes()
	exemploSwitchComExpressoes()
	exemploIfAninhado()
	exemploSwitchMultiplosValores()
	exemploValidacao()
	exemploSwitchTypeAssertion()
	exemploShortCircuit()
	exemploClassificacao()
	exemploMapLookup()
	exemploSwitchConstantes()
	exemploIfComSlices()
	exemploTratamentoErros()
}


